// Code generated by counterfeiter. DO NOT EDIT.
package authfakes

import (
	"context"
	"sync"

	v1a "k8s.io/api/core/v1"
	v1c "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
	v1b "k8s.io/client-go/applyconfigurations/core/v1"
	v1 "k8s.io/client-go/kubernetes/typed/core/v1"
)

type FakeSecretInterface struct {
	ApplyStub        func(context.Context, *v1b.SecretApplyConfiguration, v1c.ApplyOptions) (*v1a.Secret, error)
	applyMutex       sync.RWMutex
	applyArgsForCall []struct {
		arg1 context.Context
		arg2 *v1b.SecretApplyConfiguration
		arg3 v1c.ApplyOptions
	}
	applyReturns struct {
		result1 *v1a.Secret
		result2 error
	}
	applyReturnsOnCall map[int]struct {
		result1 *v1a.Secret
		result2 error
	}
	CreateStub        func(context.Context, *v1a.Secret, v1c.CreateOptions) (*v1a.Secret, error)
	createMutex       sync.RWMutex
	createArgsForCall []struct {
		arg1 context.Context
		arg2 *v1a.Secret
		arg3 v1c.CreateOptions
	}
	createReturns struct {
		result1 *v1a.Secret
		result2 error
	}
	createReturnsOnCall map[int]struct {
		result1 *v1a.Secret
		result2 error
	}
	DeleteStub        func(context.Context, string, v1c.DeleteOptions) error
	deleteMutex       sync.RWMutex
	deleteArgsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 v1c.DeleteOptions
	}
	deleteReturns struct {
		result1 error
	}
	deleteReturnsOnCall map[int]struct {
		result1 error
	}
	DeleteCollectionStub        func(context.Context, v1c.DeleteOptions, v1c.ListOptions) error
	deleteCollectionMutex       sync.RWMutex
	deleteCollectionArgsForCall []struct {
		arg1 context.Context
		arg2 v1c.DeleteOptions
		arg3 v1c.ListOptions
	}
	deleteCollectionReturns struct {
		result1 error
	}
	deleteCollectionReturnsOnCall map[int]struct {
		result1 error
	}
	GetStub        func(context.Context, string, v1c.GetOptions) (*v1a.Secret, error)
	getMutex       sync.RWMutex
	getArgsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 v1c.GetOptions
	}
	getReturns struct {
		result1 *v1a.Secret
		result2 error
	}
	getReturnsOnCall map[int]struct {
		result1 *v1a.Secret
		result2 error
	}
	ListStub        func(context.Context, v1c.ListOptions) (*v1a.SecretList, error)
	listMutex       sync.RWMutex
	listArgsForCall []struct {
		arg1 context.Context
		arg2 v1c.ListOptions
	}
	listReturns struct {
		result1 *v1a.SecretList
		result2 error
	}
	listReturnsOnCall map[int]struct {
		result1 *v1a.SecretList
		result2 error
	}
	PatchStub        func(context.Context, string, types.PatchType, []byte, v1c.PatchOptions, ...string) (*v1a.Secret, error)
	patchMutex       sync.RWMutex
	patchArgsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 types.PatchType
		arg4 []byte
		arg5 v1c.PatchOptions
		arg6 []string
	}
	patchReturns struct {
		result1 *v1a.Secret
		result2 error
	}
	patchReturnsOnCall map[int]struct {
		result1 *v1a.Secret
		result2 error
	}
	UpdateStub        func(context.Context, *v1a.Secret, v1c.UpdateOptions) (*v1a.Secret, error)
	updateMutex       sync.RWMutex
	updateArgsForCall []struct {
		arg1 context.Context
		arg2 *v1a.Secret
		arg3 v1c.UpdateOptions
	}
	updateReturns struct {
		result1 *v1a.Secret
		result2 error
	}
	updateReturnsOnCall map[int]struct {
		result1 *v1a.Secret
		result2 error
	}
	WatchStub        func(context.Context, v1c.ListOptions) (watch.Interface, error)
	watchMutex       sync.RWMutex
	watchArgsForCall []struct {
		arg1 context.Context
		arg2 v1c.ListOptions
	}
	watchReturns struct {
		result1 watch.Interface
		result2 error
	}
	watchReturnsOnCall map[int]struct {
		result1 watch.Interface
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeSecretInterface) Apply(arg1 context.Context, arg2 *v1b.SecretApplyConfiguration, arg3 v1c.ApplyOptions) (*v1a.Secret, error) {
	fake.applyMutex.Lock()
	ret, specificReturn := fake.applyReturnsOnCall[len(fake.applyArgsForCall)]
	fake.applyArgsForCall = append(fake.applyArgsForCall, struct {
		arg1 context.Context
		arg2 *v1b.SecretApplyConfiguration
		arg3 v1c.ApplyOptions
	}{arg1, arg2, arg3})
	stub := fake.ApplyStub
	fakeReturns := fake.applyReturns
	fake.recordInvocation("Apply", []interface{}{arg1, arg2, arg3})
	fake.applyMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeSecretInterface) ApplyCallCount() int {
	fake.applyMutex.RLock()
	defer fake.applyMutex.RUnlock()
	return len(fake.applyArgsForCall)
}

func (fake *FakeSecretInterface) ApplyCalls(stub func(context.Context, *v1b.SecretApplyConfiguration, v1c.ApplyOptions) (*v1a.Secret, error)) {
	fake.applyMutex.Lock()
	defer fake.applyMutex.Unlock()
	fake.ApplyStub = stub
}

func (fake *FakeSecretInterface) ApplyArgsForCall(i int) (context.Context, *v1b.SecretApplyConfiguration, v1c.ApplyOptions) {
	fake.applyMutex.RLock()
	defer fake.applyMutex.RUnlock()
	argsForCall := fake.applyArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeSecretInterface) ApplyReturns(result1 *v1a.Secret, result2 error) {
	fake.applyMutex.Lock()
	defer fake.applyMutex.Unlock()
	fake.ApplyStub = nil
	fake.applyReturns = struct {
		result1 *v1a.Secret
		result2 error
	}{result1, result2}
}

func (fake *FakeSecretInterface) ApplyReturnsOnCall(i int, result1 *v1a.Secret, result2 error) {
	fake.applyMutex.Lock()
	defer fake.applyMutex.Unlock()
	fake.ApplyStub = nil
	if fake.applyReturnsOnCall == nil {
		fake.applyReturnsOnCall = make(map[int]struct {
			result1 *v1a.Secret
			result2 error
		})
	}
	fake.applyReturnsOnCall[i] = struct {
		result1 *v1a.Secret
		result2 error
	}{result1, result2}
}

func (fake *FakeSecretInterface) Create(arg1 context.Context, arg2 *v1a.Secret, arg3 v1c.CreateOptions) (*v1a.Secret, error) {
	fake.createMutex.Lock()
	ret, specificReturn := fake.createReturnsOnCall[len(fake.createArgsForCall)]
	fake.createArgsForCall = append(fake.createArgsForCall, struct {
		arg1 context.Context
		arg2 *v1a.Secret
		arg3 v1c.CreateOptions
	}{arg1, arg2, arg3})
	stub := fake.CreateStub
	fakeReturns := fake.createReturns
	fake.recordInvocation("Create", []interface{}{arg1, arg2, arg3})
	fake.createMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeSecretInterface) CreateCallCount() int {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return len(fake.createArgsForCall)
}

func (fake *FakeSecretInterface) CreateCalls(stub func(context.Context, *v1a.Secret, v1c.CreateOptions) (*v1a.Secret, error)) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = stub
}

func (fake *FakeSecretInterface) CreateArgsForCall(i int) (context.Context, *v1a.Secret, v1c.CreateOptions) {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	argsForCall := fake.createArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeSecretInterface) CreateReturns(result1 *v1a.Secret, result2 error) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = nil
	fake.createReturns = struct {
		result1 *v1a.Secret
		result2 error
	}{result1, result2}
}

func (fake *FakeSecretInterface) CreateReturnsOnCall(i int, result1 *v1a.Secret, result2 error) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = nil
	if fake.createReturnsOnCall == nil {
		fake.createReturnsOnCall = make(map[int]struct {
			result1 *v1a.Secret
			result2 error
		})
	}
	fake.createReturnsOnCall[i] = struct {
		result1 *v1a.Secret
		result2 error
	}{result1, result2}
}

func (fake *FakeSecretInterface) Delete(arg1 context.Context, arg2 string, arg3 v1c.DeleteOptions) error {
	fake.deleteMutex.Lock()
	ret, specificReturn := fake.deleteReturnsOnCall[len(fake.deleteArgsForCall)]
	fake.deleteArgsForCall = append(fake.deleteArgsForCall, struct {
		arg1 context.Context
		arg2 string
		arg3 v1c.DeleteOptions
	}{arg1, arg2, arg3})
	stub := fake.DeleteStub
	fakeReturns := fake.deleteReturns
	fake.recordInvocation("Delete", []interface{}{arg1, arg2, arg3})
	fake.deleteMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeSecretInterface) DeleteCallCount() int {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return len(fake.deleteArgsForCall)
}

func (fake *FakeSecretInterface) DeleteCalls(stub func(context.Context, string, v1c.DeleteOptions) error) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = stub
}

func (fake *FakeSecretInterface) DeleteArgsForCall(i int) (context.Context, string, v1c.DeleteOptions) {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	argsForCall := fake.deleteArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeSecretInterface) DeleteReturns(result1 error) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = nil
	fake.deleteReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSecretInterface) DeleteReturnsOnCall(i int, result1 error) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = nil
	if fake.deleteReturnsOnCall == nil {
		fake.deleteReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeSecretInterface) DeleteCollection(arg1 context.Context, arg2 v1c.DeleteOptions, arg3 v1c.ListOptions) error {
	fake.deleteCollectionMutex.Lock()
	ret, specificReturn := fake.deleteCollectionReturnsOnCall[len(fake.deleteCollectionArgsForCall)]
	fake.deleteCollectionArgsForCall = append(fake.deleteCollectionArgsForCall, struct {
		arg1 context.Context
		arg2 v1c.DeleteOptions
		arg3 v1c.ListOptions
	}{arg1, arg2, arg3})
	stub := fake.DeleteCollectionStub
	fakeReturns := fake.deleteCollectionReturns
	fake.recordInvocation("DeleteCollection", []interface{}{arg1, arg2, arg3})
	fake.deleteCollectionMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeSecretInterface) DeleteCollectionCallCount() int {
	fake.deleteCollectionMutex.RLock()
	defer fake.deleteCollectionMutex.RUnlock()
	return len(fake.deleteCollectionArgsForCall)
}

func (fake *FakeSecretInterface) DeleteCollectionCalls(stub func(context.Context, v1c.DeleteOptions, v1c.ListOptions) error) {
	fake.deleteCollectionMutex.Lock()
	defer fake.deleteCollectionMutex.Unlock()
	fake.DeleteCollectionStub = stub
}

func (fake *FakeSecretInterface) DeleteCollectionArgsForCall(i int) (context.Context, v1c.DeleteOptions, v1c.ListOptions) {
	fake.deleteCollectionMutex.RLock()
	defer fake.deleteCollectionMutex.RUnlock()
	argsForCall := fake.deleteCollectionArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeSecretInterface) DeleteCollectionReturns(result1 error) {
	fake.deleteCollectionMutex.Lock()
	defer fake.deleteCollectionMutex.Unlock()
	fake.DeleteCollectionStub = nil
	fake.deleteCollectionReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSecretInterface) DeleteCollectionReturnsOnCall(i int, result1 error) {
	fake.deleteCollectionMutex.Lock()
	defer fake.deleteCollectionMutex.Unlock()
	fake.DeleteCollectionStub = nil
	if fake.deleteCollectionReturnsOnCall == nil {
		fake.deleteCollectionReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteCollectionReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeSecretInterface) Get(arg1 context.Context, arg2 string, arg3 v1c.GetOptions) (*v1a.Secret, error) {
	fake.getMutex.Lock()
	ret, specificReturn := fake.getReturnsOnCall[len(fake.getArgsForCall)]
	fake.getArgsForCall = append(fake.getArgsForCall, struct {
		arg1 context.Context
		arg2 string
		arg3 v1c.GetOptions
	}{arg1, arg2, arg3})
	stub := fake.GetStub
	fakeReturns := fake.getReturns
	fake.recordInvocation("Get", []interface{}{arg1, arg2, arg3})
	fake.getMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeSecretInterface) GetCallCount() int {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	return len(fake.getArgsForCall)
}

func (fake *FakeSecretInterface) GetCalls(stub func(context.Context, string, v1c.GetOptions) (*v1a.Secret, error)) {
	fake.getMutex.Lock()
	defer fake.getMutex.Unlock()
	fake.GetStub = stub
}

func (fake *FakeSecretInterface) GetArgsForCall(i int) (context.Context, string, v1c.GetOptions) {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	argsForCall := fake.getArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeSecretInterface) GetReturns(result1 *v1a.Secret, result2 error) {
	fake.getMutex.Lock()
	defer fake.getMutex.Unlock()
	fake.GetStub = nil
	fake.getReturns = struct {
		result1 *v1a.Secret
		result2 error
	}{result1, result2}
}

func (fake *FakeSecretInterface) GetReturnsOnCall(i int, result1 *v1a.Secret, result2 error) {
	fake.getMutex.Lock()
	defer fake.getMutex.Unlock()
	fake.GetStub = nil
	if fake.getReturnsOnCall == nil {
		fake.getReturnsOnCall = make(map[int]struct {
			result1 *v1a.Secret
			result2 error
		})
	}
	fake.getReturnsOnCall[i] = struct {
		result1 *v1a.Secret
		result2 error
	}{result1, result2}
}

func (fake *FakeSecretInterface) List(arg1 context.Context, arg2 v1c.ListOptions) (*v1a.SecretList, error) {
	fake.listMutex.Lock()
	ret, specificReturn := fake.listReturnsOnCall[len(fake.listArgsForCall)]
	fake.listArgsForCall = append(fake.listArgsForCall, struct {
		arg1 context.Context
		arg2 v1c.ListOptions
	}{arg1, arg2})
	stub := fake.ListStub
	fakeReturns := fake.listReturns
	fake.recordInvocation("List", []interface{}{arg1, arg2})
	fake.listMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeSecretInterface) ListCallCount() int {
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	return len(fake.listArgsForCall)
}

func (fake *FakeSecretInterface) ListCalls(stub func(context.Context, v1c.ListOptions) (*v1a.SecretList, error)) {
	fake.listMutex.Lock()
	defer fake.listMutex.Unlock()
	fake.ListStub = stub
}

func (fake *FakeSecretInterface) ListArgsForCall(i int) (context.Context, v1c.ListOptions) {
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	argsForCall := fake.listArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeSecretInterface) ListReturns(result1 *v1a.SecretList, result2 error) {
	fake.listMutex.Lock()
	defer fake.listMutex.Unlock()
	fake.ListStub = nil
	fake.listReturns = struct {
		result1 *v1a.SecretList
		result2 error
	}{result1, result2}
}

func (fake *FakeSecretInterface) ListReturnsOnCall(i int, result1 *v1a.SecretList, result2 error) {
	fake.listMutex.Lock()
	defer fake.listMutex.Unlock()
	fake.ListStub = nil
	if fake.listReturnsOnCall == nil {
		fake.listReturnsOnCall = make(map[int]struct {
			result1 *v1a.SecretList
			result2 error
		})
	}
	fake.listReturnsOnCall[i] = struct {
		result1 *v1a.SecretList
		result2 error
	}{result1, result2}
}

func (fake *FakeSecretInterface) Patch(arg1 context.Context, arg2 string, arg3 types.PatchType, arg4 []byte, arg5 v1c.PatchOptions, arg6 ...string) (*v1a.Secret, error) {
	var arg4Copy []byte
	if arg4 != nil {
		arg4Copy = make([]byte, len(arg4))
		copy(arg4Copy, arg4)
	}
	fake.patchMutex.Lock()
	ret, specificReturn := fake.patchReturnsOnCall[len(fake.patchArgsForCall)]
	fake.patchArgsForCall = append(fake.patchArgsForCall, struct {
		arg1 context.Context
		arg2 string
		arg3 types.PatchType
		arg4 []byte
		arg5 v1c.PatchOptions
		arg6 []string
	}{arg1, arg2, arg3, arg4Copy, arg5, arg6})
	stub := fake.PatchStub
	fakeReturns := fake.patchReturns
	fake.recordInvocation("Patch", []interface{}{arg1, arg2, arg3, arg4Copy, arg5, arg6})
	fake.patchMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4, arg5, arg6...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeSecretInterface) PatchCallCount() int {
	fake.patchMutex.RLock()
	defer fake.patchMutex.RUnlock()
	return len(fake.patchArgsForCall)
}

func (fake *FakeSecretInterface) PatchCalls(stub func(context.Context, string, types.PatchType, []byte, v1c.PatchOptions, ...string) (*v1a.Secret, error)) {
	fake.patchMutex.Lock()
	defer fake.patchMutex.Unlock()
	fake.PatchStub = stub
}

func (fake *FakeSecretInterface) PatchArgsForCall(i int) (context.Context, string, types.PatchType, []byte, v1c.PatchOptions, []string) {
	fake.patchMutex.RLock()
	defer fake.patchMutex.RUnlock()
	argsForCall := fake.patchArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5, argsForCall.arg6
}

func (fake *FakeSecretInterface) PatchReturns(result1 *v1a.Secret, result2 error) {
	fake.patchMutex.Lock()
	defer fake.patchMutex.Unlock()
	fake.PatchStub = nil
	fake.patchReturns = struct {
		result1 *v1a.Secret
		result2 error
	}{result1, result2}
}

func (fake *FakeSecretInterface) PatchReturnsOnCall(i int, result1 *v1a.Secret, result2 error) {
	fake.patchMutex.Lock()
	defer fake.patchMutex.Unlock()
	fake.PatchStub = nil
	if fake.patchReturnsOnCall == nil {
		fake.patchReturnsOnCall = make(map[int]struct {
			result1 *v1a.Secret
			result2 error
		})
	}
	fake.patchReturnsOnCall[i] = struct {
		result1 *v1a.Secret
		result2 error
	}{result1, result2}
}

func (fake *FakeSecretInterface) Update(arg1 context.Context, arg2 *v1a.Secret, arg3 v1c.UpdateOptions) (*v1a.Secret, error) {
	fake.updateMutex.Lock()
	ret, specificReturn := fake.updateReturnsOnCall[len(fake.updateArgsForCall)]
	fake.updateArgsForCall = append(fake.updateArgsForCall, struct {
		arg1 context.Context
		arg2 *v1a.Secret
		arg3 v1c.UpdateOptions
	}{arg1, arg2, arg3})
	stub := fake.UpdateStub
	fakeReturns := fake.updateReturns
	fake.recordInvocation("Update", []interface{}{arg1, arg2, arg3})
	fake.updateMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeSecretInterface) UpdateCallCount() int {
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	return len(fake.updateArgsForCall)
}

func (fake *FakeSecretInterface) UpdateCalls(stub func(context.Context, *v1a.Secret, v1c.UpdateOptions) (*v1a.Secret, error)) {
	fake.updateMutex.Lock()
	defer fake.updateMutex.Unlock()
	fake.UpdateStub = stub
}

func (fake *FakeSecretInterface) UpdateArgsForCall(i int) (context.Context, *v1a.Secret, v1c.UpdateOptions) {
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	argsForCall := fake.updateArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeSecretInterface) UpdateReturns(result1 *v1a.Secret, result2 error) {
	fake.updateMutex.Lock()
	defer fake.updateMutex.Unlock()
	fake.UpdateStub = nil
	fake.updateReturns = struct {
		result1 *v1a.Secret
		result2 error
	}{result1, result2}
}

func (fake *FakeSecretInterface) UpdateReturnsOnCall(i int, result1 *v1a.Secret, result2 error) {
	fake.updateMutex.Lock()
	defer fake.updateMutex.Unlock()
	fake.UpdateStub = nil
	if fake.updateReturnsOnCall == nil {
		fake.updateReturnsOnCall = make(map[int]struct {
			result1 *v1a.Secret
			result2 error
		})
	}
	fake.updateReturnsOnCall[i] = struct {
		result1 *v1a.Secret
		result2 error
	}{result1, result2}
}

func (fake *FakeSecretInterface) Watch(arg1 context.Context, arg2 v1c.ListOptions) (watch.Interface, error) {
	fake.watchMutex.Lock()
	ret, specificReturn := fake.watchReturnsOnCall[len(fake.watchArgsForCall)]
	fake.watchArgsForCall = append(fake.watchArgsForCall, struct {
		arg1 context.Context
		arg2 v1c.ListOptions
	}{arg1, arg2})
	stub := fake.WatchStub
	fakeReturns := fake.watchReturns
	fake.recordInvocation("Watch", []interface{}{arg1, arg2})
	fake.watchMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeSecretInterface) WatchCallCount() int {
	fake.watchMutex.RLock()
	defer fake.watchMutex.RUnlock()
	return len(fake.watchArgsForCall)
}

func (fake *FakeSecretInterface) WatchCalls(stub func(context.Context, v1c.ListOptions) (watch.Interface, error)) {
	fake.watchMutex.Lock()
	defer fake.watchMutex.Unlock()
	fake.WatchStub = stub
}

func (fake *FakeSecretInterface) WatchArgsForCall(i int) (context.Context, v1c.ListOptions) {
	fake.watchMutex.RLock()
	defer fake.watchMutex.RUnlock()
	argsForCall := fake.watchArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeSecretInterface) WatchReturns(result1 watch.Interface, result2 error) {
	fake.watchMutex.Lock()
	defer fake.watchMutex.Unlock()
	fake.WatchStub = nil
	fake.watchReturns = struct {
		result1 watch.Interface
		result2 error
	}{result1, result2}
}

func (fake *FakeSecretInterface) WatchReturnsOnCall(i int, result1 watch.Interface, result2 error) {
	fake.watchMutex.Lock()
	defer fake.watchMutex.Unlock()
	fake.WatchStub = nil
	if fake.watchReturnsOnCall == nil {
		fake.watchReturnsOnCall = make(map[int]struct {
			result1 watch.Interface
			result2 error
		})
	}
	fake.watchReturnsOnCall[i] = struct {
		result1 watch.Interface
		result2 error
	}{result1, result2}
}

func (fake *FakeSecretInterface) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.applyMutex.RLock()
	defer fake.applyMutex.RUnlock()
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	fake.deleteCollectionMutex.RLock()
	defer fake.deleteCollectionMutex.RUnlock()
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	fake.patchMutex.RLock()
	defer fake.patchMutex.RUnlock()
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	fake.watchMutex.RLock()
	defer fake.watchMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeSecretInterface) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ v1.SecretInterface = new(FakeSecretInterface)
