package install_test

import (
	"encoding/json"
	"fmt"
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/epinio/epinio/acceptance/helpers/catalog"
	"github.com/epinio/epinio/acceptance/helpers/epinio"
	"github.com/epinio/epinio/acceptance/helpers/proc"
	"github.com/epinio/epinio/acceptance/testenv"
)

var _ = Describe("<Scenario3> RKE, Private CA, Service, on External Registry", func() {
	var (
		flags            []string
		epinioHelper     epinio.Epinio
		serviceName      = catalog.NewServiceName()
		appName          string
		loadbalancer     string
		metallbURL       string
		localpathURL     string
		registryUsername string
		registryPassword string
		// testenv.New is not needed for VerifyAppServiceBound helper :shrug:
		env      testenv.EpinioEnv
		domainIP = "192.168.1.240" // Set it to an arbitrary private IP
	)

	BeforeEach(func() {
		epinioHelper = epinio.NewEpinioHelper(testenv.EpinioBinaryPath())

		metallbURL = "https://raw.githubusercontent.com/google/metallb/v0.10.3/manifests/metallb.yaml"
		localpathURL = "https://raw.githubusercontent.com/rancher/local-path-provisioner/v0.0.20/deploy/local-path-storage.yaml"
		appName = "externalregtest"

		registryUsername = os.Getenv("REGISTRY_USERNAME")
		Expect(registryUsername).ToNot(BeEmpty())

		registryPassword = os.Getenv("REGISTRY_PASSWORD")
		Expect(registryPassword).ToNot(BeEmpty())
		flags = []string{
			"--set", "skipCertManager=true",
			"--set", "domain=" + fmt.Sprintf("%s.omg.howdoi.website", domainIP),
			"--set", "tlsIssuer=private-ca",
			"--set", "externalRegistryURL=registry.hub.docker.com",
			"--set", "externalRegistryUsername=" + registryUsername,
			"--set", "externalRegistryPassword=" + registryPassword,
			"--set", "externalRegistryNamespace=splatform",
		}

	})

	AfterEach(func() {
		out, err := epinioHelper.Uninstall()
		Expect(err).NotTo(HaveOccurred(), out)
	})

	It("installs with private CA and pushes an app with service", func() {
		By("Installing MetalLB", func() {
			out, err := proc.RunW("kubectl", "create", "namespace", "metallb-system")
			Expect(err).NotTo(HaveOccurred(), out)

			out, err = proc.RunW("kubectl", "apply", "-f", metallbURL)
			Expect(err).NotTo(HaveOccurred(), out)

			out, err = proc.RunW("sed", "-i", fmt.Sprintf("s/myip/%s/g", domainIP), testenv.TestAssetPath("config-metallb-rke.yml"))
			Expect(err).NotTo(HaveOccurred(), out)

			out, err = proc.RunW("kubectl", "apply", "-f", testenv.TestAssetPath("config-metallb-rke.yml"))
			Expect(err).NotTo(HaveOccurred(), out)
		})

		By("Configuring local-path storage", func() {
			out, err := proc.RunW("kubectl", "apply", "-f", localpathURL)
			Expect(err).NotTo(HaveOccurred(), out)

			value := `{"metadata": {"annotations":{"storageclass.kubernetes.io/is-default-class":"true"}}}`
			out, err = proc.RunW("kubectl", "patch", "storageclass", "local-path", "-p", value)
			Expect(err).NotTo(HaveOccurred(), out)
		})

		By("Installing CertManager", func() {
			out, err := proc.RunW("helm", "repo", "add", "jetstack", "https://charts.jetstack.io")
			Expect(err).NotTo(HaveOccurred(), out)
			out, err = proc.RunW("helm", "repo", "update")
			Expect(err).NotTo(HaveOccurred(), out)
			out, err = proc.RunW("helm", "upgrade", "--install", "cert-manager", "jetstack/cert-manager",
				"-n", "cert-manager",
				"--create-namespace",
				"--set", "installCRDs=true",
				"--set", "extraArgs[0]=--enable-certificate-owner-ref=true",
			)
			Expect(err).NotTo(HaveOccurred(), out)

			// Create certificate secret and cluster_issuer
			out, err = proc.RunW("kubectl", "apply", "-f", testenv.TestAssetPath("cluster-issuer-private-ca.yml"))
			Expect(err).NotTo(HaveOccurred(), out)
		})

		By("Installing Epinio", func() {
			out, err := epinioHelper.Install(flags...)
			Expect(err).NotTo(HaveOccurred(), out)
			Expect(out).To(ContainSubstring("STATUS: deployed"))

			out, err = testenv.PatchEpinio()
			Expect(err).ToNot(HaveOccurred(), out)
		})

		By("Updating Epinio config", func() {
			out, err := epinioHelper.Run("config", "update")
			Expect(err).NotTo(HaveOccurred(), out)
			Expect(out).To(ContainSubstring("Ok"))
		})

		By("Checking Loadbalancer IP", func() {
			out, err := proc.RunW("kubectl", "get", "service", "-n", "traefik", "traefik", "-o", "json")
			Expect(err).NotTo(HaveOccurred(), out)

			status := &testenv.LoadBalancerHostname{}
			err = json.Unmarshal([]byte(out), status)
			Expect(err).NotTo(HaveOccurred())
			Expect(status.Status.LoadBalancer.Ingress).To(HaveLen(1))
			loadbalancer = status.Status.LoadBalancer.Ingress[0].IP
			Expect(loadbalancer).ToNot(BeEmpty())
		})

		By("Creating a service and pushing an app", func() {
			out, err := epinioHelper.Run("service", "create", serviceName, "mariadb", "10-3-22")
			Expect(err).NotTo(HaveOccurred(), out)

			out, err = epinioHelper.Run("push",
				"--name", appName,
				"--path", testenv.AssetPath("sample-app"),
				"--bind", serviceName)
			Expect(err).NotTo(HaveOccurred(), out)

			env.VerifyAppServiceBound(appName, serviceName, testenv.DefaultWorkspace, 1)

			// Verify cluster_issuer is used
			out, err = proc.RunW("kubectl", "get", "certificate",
				"-n", testenv.DefaultWorkspace,
				"--selector", "app.kubernetes.io/name="+appName,
				"-o", "jsonpath='{.items[*].spec.issuerRef.name}'")
			Expect(err).NotTo(HaveOccurred(), out)
			Expect(out).To(Equal("'private-ca'"))
		})
	})
})
