// Code generated by counterfeiter. DO NOT EDIT.
package execfakes

import (
	"os"
	"sync"

	"github.com/concourse/atc/exec"
)

type FakeStep struct {
	RunStub        func(signals <-chan os.Signal, ready chan<- struct{}) error
	runMutex       sync.RWMutex
	runArgsForCall []struct {
		signals <-chan os.Signal
		ready   chan<- struct{}
	}
	runReturns struct {
		result1 error
	}
	runReturnsOnCall map[int]struct {
		result1 error
	}
	ResultStub        func(interface{}) bool
	resultMutex       sync.RWMutex
	resultArgsForCall []struct {
		arg1 interface{}
	}
	resultReturns struct {
		result1 bool
	}
	resultReturnsOnCall map[int]struct {
		result1 bool
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeStep) Run(signals <-chan os.Signal, ready chan<- struct{}) error {
	fake.runMutex.Lock()
	ret, specificReturn := fake.runReturnsOnCall[len(fake.runArgsForCall)]
	fake.runArgsForCall = append(fake.runArgsForCall, struct {
		signals <-chan os.Signal
		ready   chan<- struct{}
	}{signals, ready})
	fake.recordInvocation("Run", []interface{}{signals, ready})
	fake.runMutex.Unlock()
	if fake.RunStub != nil {
		return fake.RunStub(signals, ready)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.runReturns.result1
}

func (fake *FakeStep) RunCallCount() int {
	fake.runMutex.RLock()
	defer fake.runMutex.RUnlock()
	return len(fake.runArgsForCall)
}

func (fake *FakeStep) RunArgsForCall(i int) (<-chan os.Signal, chan<- struct{}) {
	fake.runMutex.RLock()
	defer fake.runMutex.RUnlock()
	return fake.runArgsForCall[i].signals, fake.runArgsForCall[i].ready
}

func (fake *FakeStep) RunReturns(result1 error) {
	fake.RunStub = nil
	fake.runReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeStep) RunReturnsOnCall(i int, result1 error) {
	fake.RunStub = nil
	if fake.runReturnsOnCall == nil {
		fake.runReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.runReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeStep) Result(arg1 interface{}) bool {
	fake.resultMutex.Lock()
	ret, specificReturn := fake.resultReturnsOnCall[len(fake.resultArgsForCall)]
	fake.resultArgsForCall = append(fake.resultArgsForCall, struct {
		arg1 interface{}
	}{arg1})
	fake.recordInvocation("Result", []interface{}{arg1})
	fake.resultMutex.Unlock()
	if fake.ResultStub != nil {
		return fake.ResultStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.resultReturns.result1
}

func (fake *FakeStep) ResultCallCount() int {
	fake.resultMutex.RLock()
	defer fake.resultMutex.RUnlock()
	return len(fake.resultArgsForCall)
}

func (fake *FakeStep) ResultArgsForCall(i int) interface{} {
	fake.resultMutex.RLock()
	defer fake.resultMutex.RUnlock()
	return fake.resultArgsForCall[i].arg1
}

func (fake *FakeStep) ResultReturns(result1 bool) {
	fake.ResultStub = nil
	fake.resultReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeStep) ResultReturnsOnCall(i int, result1 bool) {
	fake.ResultStub = nil
	if fake.resultReturnsOnCall == nil {
		fake.resultReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.resultReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *FakeStep) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.runMutex.RLock()
	defer fake.runMutex.RUnlock()
	fake.resultMutex.RLock()
	defer fake.resultMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeStep) recordInvocation(key string, args []interface{}) {
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

var _ exec.Step = new(FakeStep)
