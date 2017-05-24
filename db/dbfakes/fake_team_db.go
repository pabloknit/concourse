// Code generated by counterfeiter. DO NOT EDIT.
package dbfakes

import (
	"sync"

	"github.com/concourse/atc/db"
)

type FakeTeamDB struct {
	GetTeamStub        func() (db.SavedTeam, bool, error)
	getTeamMutex       sync.RWMutex
	getTeamArgsForCall []struct{}
	getTeamReturns     struct {
		result1 db.SavedTeam
		result2 bool
		result3 error
	}
	getTeamReturnsOnCall map[int]struct {
		result1 db.SavedTeam
		result2 bool
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeTeamDB) GetTeam() (db.SavedTeam, bool, error) {
	fake.getTeamMutex.Lock()
	ret, specificReturn := fake.getTeamReturnsOnCall[len(fake.getTeamArgsForCall)]
	fake.getTeamArgsForCall = append(fake.getTeamArgsForCall, struct{}{})
	fake.recordInvocation("GetTeam", []interface{}{})
	fake.getTeamMutex.Unlock()
	if fake.GetTeamStub != nil {
		return fake.GetTeamStub()
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.getTeamReturns.result1, fake.getTeamReturns.result2, fake.getTeamReturns.result3
}

func (fake *FakeTeamDB) GetTeamCallCount() int {
	fake.getTeamMutex.RLock()
	defer fake.getTeamMutex.RUnlock()
	return len(fake.getTeamArgsForCall)
}

func (fake *FakeTeamDB) GetTeamReturns(result1 db.SavedTeam, result2 bool, result3 error) {
	fake.GetTeamStub = nil
	fake.getTeamReturns = struct {
		result1 db.SavedTeam
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeTeamDB) GetTeamReturnsOnCall(i int, result1 db.SavedTeam, result2 bool, result3 error) {
	fake.GetTeamStub = nil
	if fake.getTeamReturnsOnCall == nil {
		fake.getTeamReturnsOnCall = make(map[int]struct {
			result1 db.SavedTeam
			result2 bool
			result3 error
		})
	}
	fake.getTeamReturnsOnCall[i] = struct {
		result1 db.SavedTeam
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeTeamDB) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getTeamMutex.RLock()
	defer fake.getTeamMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeTeamDB) recordInvocation(key string, args []interface{}) {
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

var _ db.TeamDB = new(FakeTeamDB)
