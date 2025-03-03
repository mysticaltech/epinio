package acceptance_test

import (
	"github.com/epinio/epinio/acceptance/helpers/catalog"
	"github.com/epinio/epinio/acceptance/helpers/proc"

	. "github.com/epinio/epinio/acceptance/helpers/matchers"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Login", func() {
	var tmpSettingsPath string

	BeforeEach(func() {
		tmpSettingsPath = catalog.NewTmpName("tmpEpinio") + `.yaml`
	})

	AfterEach(func() {
		// Remove transient settings
		out, err := proc.Run("", false, "rm", "-f", tmpSettingsPath)
		Expect(err).ToNot(HaveOccurred(), out)
	})

	It("succeed with a valid user", func() {
		// check that the initial settings are empty
		settings, err := env.Epinio("", "settings", "show", "--settings-file", tmpSettingsPath)
		Expect(err).ToNot(HaveOccurred(), settings)
		Expect(settings).To(
			HaveATable(
				WithHeaders("KEY", "VALUE"),
				WithRow("API User Name", ""),
				WithRow("API Password", ""),
				WithRow("Certificates", "None defined"),
			),
		)

		// login with a different user
		out, err := env.Epinio("", "login", "-u", "epinio", "-p", env.EpinioPassword, "--trust-ca", "--settings-file", tmpSettingsPath, serverURL)
		Expect(err).ToNot(HaveOccurred())
		Expect(out).To(ContainSubstring(`Login to your Epinio cluster`))
		Expect(out).To(ContainSubstring(`Trusting certificate...`))
		Expect(out).To(ContainSubstring(`Login successful`))

		// check that the settings are now updated
		settings, err = env.Epinio("", "settings", "show", "--settings-file", tmpSettingsPath)
		Expect(err).ToNot(HaveOccurred(), settings)
		Expect(settings).To(
			HaveATable(
				WithHeaders("KEY", "VALUE"),
				WithRow("API User Name", "epinio"),
				WithRow("API Password", "[*]+"),
				WithRow("Certificates", "Present"),
			),
		)
	})

	It("fails with a non existing user", func() {
		// check that the initial settings are empty
		settings, err := env.Epinio("", "settings", "show", "--settings-file", tmpSettingsPath)
		Expect(err).ToNot(HaveOccurred(), settings)
		Expect(settings).To(
			HaveATable(
				WithHeaders("KEY", "VALUE"),
				WithRow("API User Name", ""),
				WithRow("API Password", ""),
				WithRow("Certificates", "None defined"),
			),
		)

		// login with a non existing user
		out, err := env.Epinio("", "login", "-u", "unknown", "-p", env.EpinioPassword, "--trust-ca", "--settings-file", tmpSettingsPath, serverURL)
		Expect(err).To(HaveOccurred(), out)
		Expect(out).To(ContainSubstring(`error verifying credentials`))

		// check that the initial settings are still empty
		settings, err = env.Epinio("", "settings", "show", "--settings-file", tmpSettingsPath)
		Expect(err).ToNot(HaveOccurred(), settings)
		Expect(settings).To(
			HaveATable(
				WithHeaders("KEY", "VALUE"),
				WithRow("API User Name", ""),
				WithRow("API Password", ""),
				WithRow("Certificates", "None defined"),
			),
		)
	})
})
