// Code generated by counterfeiter. DO NOT EDIT.
package mock

import (
	"sync"

	"github.com/hyperledger-labs/fabric-token-sdk/token/services/vault/translator"
)

type IssueAction struct {
	GetIssuerStub        func() []byte
	getIssuerMutex       sync.RWMutex
	getIssuerArgsForCall []struct {
	}
	getIssuerReturns struct {
		result1 []byte
	}
	getIssuerReturnsOnCall map[int]struct {
		result1 []byte
	}
	GetSerializedOutputsStub        func() ([][]byte, error)
	getSerializedOutputsMutex       sync.RWMutex
	getSerializedOutputsArgsForCall []struct {
	}
	getSerializedOutputsReturns struct {
		result1 [][]byte
		result2 error
	}
	getSerializedOutputsReturnsOnCall map[int]struct {
		result1 [][]byte
		result2 error
	}
	IsAnonymousStub        func() bool
	isAnonymousMutex       sync.RWMutex
	isAnonymousArgsForCall []struct {
	}
	isAnonymousReturns struct {
		result1 bool
	}
	isAnonymousReturnsOnCall map[int]struct {
		result1 bool
	}
	NumOutputsStub        func() int
	numOutputsMutex       sync.RWMutex
	numOutputsArgsForCall []struct {
	}
	numOutputsReturns struct {
		result1 int
	}
	numOutputsReturnsOnCall map[int]struct {
		result1 int
	}
	SerializeStub        func() ([]byte, error)
	serializeMutex       sync.RWMutex
	serializeArgsForCall []struct {
	}
	serializeReturns struct {
		result1 []byte
		result2 error
	}
	serializeReturnsOnCall map[int]struct {
		result1 []byte
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *IssueAction) GetIssuer() []byte {
	fake.getIssuerMutex.Lock()
	ret, specificReturn := fake.getIssuerReturnsOnCall[len(fake.getIssuerArgsForCall)]
	fake.getIssuerArgsForCall = append(fake.getIssuerArgsForCall, struct {
	}{})
	fake.recordInvocation("GetIssuer", []interface{}{})
	fake.getIssuerMutex.Unlock()
	if fake.GetIssuerStub != nil {
		return fake.GetIssuerStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.getIssuerReturns
	return fakeReturns.result1
}

func (fake *IssueAction) GetIssuerCallCount() int {
	fake.getIssuerMutex.RLock()
	defer fake.getIssuerMutex.RUnlock()
	return len(fake.getIssuerArgsForCall)
}

func (fake *IssueAction) GetIssuerCalls(stub func() []byte) {
	fake.getIssuerMutex.Lock()
	defer fake.getIssuerMutex.Unlock()
	fake.GetIssuerStub = stub
}

func (fake *IssueAction) GetIssuerReturns(result1 []byte) {
	fake.getIssuerMutex.Lock()
	defer fake.getIssuerMutex.Unlock()
	fake.GetIssuerStub = nil
	fake.getIssuerReturns = struct {
		result1 []byte
	}{result1}
}

func (fake *IssueAction) GetIssuerReturnsOnCall(i int, result1 []byte) {
	fake.getIssuerMutex.Lock()
	defer fake.getIssuerMutex.Unlock()
	fake.GetIssuerStub = nil
	if fake.getIssuerReturnsOnCall == nil {
		fake.getIssuerReturnsOnCall = make(map[int]struct {
			result1 []byte
		})
	}
	fake.getIssuerReturnsOnCall[i] = struct {
		result1 []byte
	}{result1}
}

func (fake *IssueAction) GetSerializedOutputs() ([][]byte, error) {
	fake.getSerializedOutputsMutex.Lock()
	ret, specificReturn := fake.getSerializedOutputsReturnsOnCall[len(fake.getSerializedOutputsArgsForCall)]
	fake.getSerializedOutputsArgsForCall = append(fake.getSerializedOutputsArgsForCall, struct {
	}{})
	fake.recordInvocation("GetSerializedOutputs", []interface{}{})
	fake.getSerializedOutputsMutex.Unlock()
	if fake.GetSerializedOutputsStub != nil {
		return fake.GetSerializedOutputsStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getSerializedOutputsReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *IssueAction) GetSerializedOutputsCallCount() int {
	fake.getSerializedOutputsMutex.RLock()
	defer fake.getSerializedOutputsMutex.RUnlock()
	return len(fake.getSerializedOutputsArgsForCall)
}

func (fake *IssueAction) GetSerializedOutputsCalls(stub func() ([][]byte, error)) {
	fake.getSerializedOutputsMutex.Lock()
	defer fake.getSerializedOutputsMutex.Unlock()
	fake.GetSerializedOutputsStub = stub
}

func (fake *IssueAction) GetSerializedOutputsReturns(result1 [][]byte, result2 error) {
	fake.getSerializedOutputsMutex.Lock()
	defer fake.getSerializedOutputsMutex.Unlock()
	fake.GetSerializedOutputsStub = nil
	fake.getSerializedOutputsReturns = struct {
		result1 [][]byte
		result2 error
	}{result1, result2}
}

func (fake *IssueAction) GetSerializedOutputsReturnsOnCall(i int, result1 [][]byte, result2 error) {
	fake.getSerializedOutputsMutex.Lock()
	defer fake.getSerializedOutputsMutex.Unlock()
	fake.GetSerializedOutputsStub = nil
	if fake.getSerializedOutputsReturnsOnCall == nil {
		fake.getSerializedOutputsReturnsOnCall = make(map[int]struct {
			result1 [][]byte
			result2 error
		})
	}
	fake.getSerializedOutputsReturnsOnCall[i] = struct {
		result1 [][]byte
		result2 error
	}{result1, result2}
}

func (fake *IssueAction) IsAnonymous() bool {
	fake.isAnonymousMutex.Lock()
	ret, specificReturn := fake.isAnonymousReturnsOnCall[len(fake.isAnonymousArgsForCall)]
	fake.isAnonymousArgsForCall = append(fake.isAnonymousArgsForCall, struct {
	}{})
	fake.recordInvocation("IsAnonymous", []interface{}{})
	fake.isAnonymousMutex.Unlock()
	if fake.IsAnonymousStub != nil {
		return fake.IsAnonymousStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.isAnonymousReturns
	return fakeReturns.result1
}

func (fake *IssueAction) IsAnonymousCallCount() int {
	fake.isAnonymousMutex.RLock()
	defer fake.isAnonymousMutex.RUnlock()
	return len(fake.isAnonymousArgsForCall)
}

func (fake *IssueAction) IsAnonymousCalls(stub func() bool) {
	fake.isAnonymousMutex.Lock()
	defer fake.isAnonymousMutex.Unlock()
	fake.IsAnonymousStub = stub
}

func (fake *IssueAction) IsAnonymousReturns(result1 bool) {
	fake.isAnonymousMutex.Lock()
	defer fake.isAnonymousMutex.Unlock()
	fake.IsAnonymousStub = nil
	fake.isAnonymousReturns = struct {
		result1 bool
	}{result1}
}

func (fake *IssueAction) IsAnonymousReturnsOnCall(i int, result1 bool) {
	fake.isAnonymousMutex.Lock()
	defer fake.isAnonymousMutex.Unlock()
	fake.IsAnonymousStub = nil
	if fake.isAnonymousReturnsOnCall == nil {
		fake.isAnonymousReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.isAnonymousReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *IssueAction) NumOutputs() int {
	fake.numOutputsMutex.Lock()
	ret, specificReturn := fake.numOutputsReturnsOnCall[len(fake.numOutputsArgsForCall)]
	fake.numOutputsArgsForCall = append(fake.numOutputsArgsForCall, struct {
	}{})
	fake.recordInvocation("NumOutputs", []interface{}{})
	fake.numOutputsMutex.Unlock()
	if fake.NumOutputsStub != nil {
		return fake.NumOutputsStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.numOutputsReturns
	return fakeReturns.result1
}

func (fake *IssueAction) NumOutputsCallCount() int {
	fake.numOutputsMutex.RLock()
	defer fake.numOutputsMutex.RUnlock()
	return len(fake.numOutputsArgsForCall)
}

func (fake *IssueAction) NumOutputsCalls(stub func() int) {
	fake.numOutputsMutex.Lock()
	defer fake.numOutputsMutex.Unlock()
	fake.NumOutputsStub = stub
}

func (fake *IssueAction) NumOutputsReturns(result1 int) {
	fake.numOutputsMutex.Lock()
	defer fake.numOutputsMutex.Unlock()
	fake.NumOutputsStub = nil
	fake.numOutputsReturns = struct {
		result1 int
	}{result1}
}

func (fake *IssueAction) NumOutputsReturnsOnCall(i int, result1 int) {
	fake.numOutputsMutex.Lock()
	defer fake.numOutputsMutex.Unlock()
	fake.NumOutputsStub = nil
	if fake.numOutputsReturnsOnCall == nil {
		fake.numOutputsReturnsOnCall = make(map[int]struct {
			result1 int
		})
	}
	fake.numOutputsReturnsOnCall[i] = struct {
		result1 int
	}{result1}
}

func (fake *IssueAction) Serialize() ([]byte, error) {
	fake.serializeMutex.Lock()
	ret, specificReturn := fake.serializeReturnsOnCall[len(fake.serializeArgsForCall)]
	fake.serializeArgsForCall = append(fake.serializeArgsForCall, struct {
	}{})
	fake.recordInvocation("Serialize", []interface{}{})
	fake.serializeMutex.Unlock()
	if fake.SerializeStub != nil {
		return fake.SerializeStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.serializeReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *IssueAction) SerializeCallCount() int {
	fake.serializeMutex.RLock()
	defer fake.serializeMutex.RUnlock()
	return len(fake.serializeArgsForCall)
}

func (fake *IssueAction) SerializeCalls(stub func() ([]byte, error)) {
	fake.serializeMutex.Lock()
	defer fake.serializeMutex.Unlock()
	fake.SerializeStub = stub
}

func (fake *IssueAction) SerializeReturns(result1 []byte, result2 error) {
	fake.serializeMutex.Lock()
	defer fake.serializeMutex.Unlock()
	fake.SerializeStub = nil
	fake.serializeReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *IssueAction) SerializeReturnsOnCall(i int, result1 []byte, result2 error) {
	fake.serializeMutex.Lock()
	defer fake.serializeMutex.Unlock()
	fake.SerializeStub = nil
	if fake.serializeReturnsOnCall == nil {
		fake.serializeReturnsOnCall = make(map[int]struct {
			result1 []byte
			result2 error
		})
	}
	fake.serializeReturnsOnCall[i] = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *IssueAction) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getIssuerMutex.RLock()
	defer fake.getIssuerMutex.RUnlock()
	fake.getSerializedOutputsMutex.RLock()
	defer fake.getSerializedOutputsMutex.RUnlock()
	fake.isAnonymousMutex.RLock()
	defer fake.isAnonymousMutex.RUnlock()
	fake.numOutputsMutex.RLock()
	defer fake.numOutputsMutex.RUnlock()
	fake.serializeMutex.RLock()
	defer fake.serializeMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *IssueAction) recordInvocation(key string, args []interface{}) {
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

var _ translator.IssueAction = new(IssueAction)