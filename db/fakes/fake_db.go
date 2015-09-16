// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/cloudfoundry-incubator/routing-api/db"
	"github.com/cloudfoundry/storeadapter"
)

type FakeDB struct {
	ReadRoutesStub        func() ([]db.Route, error)
	readRoutesMutex       sync.RWMutex
	readRoutesArgsForCall []struct{}
	readRoutesReturns struct {
		result1 []db.Route
		result2 error
	}
	SaveRouteStub        func(route db.Route) error
	saveRouteMutex       sync.RWMutex
	saveRouteArgsForCall []struct {
		route db.Route
	}
	saveRouteReturns struct {
		result1 error
	}
	DeleteRouteStub        func(route db.Route) error
	deleteRouteMutex       sync.RWMutex
	deleteRouteArgsForCall []struct {
		route db.Route
	}
	deleteRouteReturns struct {
		result1 error
	}
	ReadTcpMappingsStub        func() ([]db.TcpMapping, error)
	readTcpMappingsMutex       sync.RWMutex
	readTcpMappingsArgsForCall []struct{}
	readTcpMappingsReturns struct {
		result1 []db.TcpMapping
		result2 error
	}
	SaveTcpMappingStub        func(tcpMapping db.TcpMapping) error
	saveTcpMappingMutex       sync.RWMutex
	saveTcpMappingArgsForCall []struct {
		tcpMapping db.TcpMapping
	}
	saveTcpMappingReturns struct {
		result1 error
	}
	ConnectStub        func() error
	connectMutex       sync.RWMutex
	connectArgsForCall []struct{}
	connectReturns struct {
		result1 error
	}
	DisconnectStub        func() error
	disconnectMutex       sync.RWMutex
	disconnectArgsForCall []struct{}
	disconnectReturns struct {
		result1 error
	}
	WatchRouteChangesStub        func() (<-chan storeadapter.WatchEvent, chan<- bool, <-chan error)
	watchRouteChangesMutex       sync.RWMutex
	watchRouteChangesArgsForCall []struct{}
	watchRouteChangesReturns struct {
		result1 <-chan storeadapter.WatchEvent
		result2 chan<- bool
		result3 <-chan error
	}
}

func (fake *FakeDB) ReadRoutes() ([]db.Route, error) {
	fake.readRoutesMutex.Lock()
	fake.readRoutesArgsForCall = append(fake.readRoutesArgsForCall, struct{}{})
	fake.readRoutesMutex.Unlock()
	if fake.ReadRoutesStub != nil {
		return fake.ReadRoutesStub()
	} else {
		return fake.readRoutesReturns.result1, fake.readRoutesReturns.result2
	}
}

func (fake *FakeDB) ReadRoutesCallCount() int {
	fake.readRoutesMutex.RLock()
	defer fake.readRoutesMutex.RUnlock()
	return len(fake.readRoutesArgsForCall)
}

func (fake *FakeDB) ReadRoutesReturns(result1 []db.Route, result2 error) {
	fake.ReadRoutesStub = nil
	fake.readRoutesReturns = struct {
		result1 []db.Route
		result2 error
	}{result1, result2}
}

func (fake *FakeDB) SaveRoute(route db.Route) error {
	fake.saveRouteMutex.Lock()
	fake.saveRouteArgsForCall = append(fake.saveRouteArgsForCall, struct {
		route db.Route
	}{route})
	fake.saveRouteMutex.Unlock()
	if fake.SaveRouteStub != nil {
		return fake.SaveRouteStub(route)
	} else {
		return fake.saveRouteReturns.result1
	}
}

func (fake *FakeDB) SaveRouteCallCount() int {
	fake.saveRouteMutex.RLock()
	defer fake.saveRouteMutex.RUnlock()
	return len(fake.saveRouteArgsForCall)
}

func (fake *FakeDB) SaveRouteArgsForCall(i int) db.Route {
	fake.saveRouteMutex.RLock()
	defer fake.saveRouteMutex.RUnlock()
	return fake.saveRouteArgsForCall[i].route
}

func (fake *FakeDB) SaveRouteReturns(result1 error) {
	fake.SaveRouteStub = nil
	fake.saveRouteReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeDB) DeleteRoute(route db.Route) error {
	fake.deleteRouteMutex.Lock()
	fake.deleteRouteArgsForCall = append(fake.deleteRouteArgsForCall, struct {
		route db.Route
	}{route})
	fake.deleteRouteMutex.Unlock()
	if fake.DeleteRouteStub != nil {
		return fake.DeleteRouteStub(route)
	} else {
		return fake.deleteRouteReturns.result1
	}
}

func (fake *FakeDB) DeleteRouteCallCount() int {
	fake.deleteRouteMutex.RLock()
	defer fake.deleteRouteMutex.RUnlock()
	return len(fake.deleteRouteArgsForCall)
}

func (fake *FakeDB) DeleteRouteArgsForCall(i int) db.Route {
	fake.deleteRouteMutex.RLock()
	defer fake.deleteRouteMutex.RUnlock()
	return fake.deleteRouteArgsForCall[i].route
}

func (fake *FakeDB) DeleteRouteReturns(result1 error) {
	fake.DeleteRouteStub = nil
	fake.deleteRouteReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeDB) ReadTcpMappings() ([]db.TcpMapping, error) {
	fake.readTcpMappingsMutex.Lock()
	fake.readTcpMappingsArgsForCall = append(fake.readTcpMappingsArgsForCall, struct{}{})
	fake.readTcpMappingsMutex.Unlock()
	if fake.ReadTcpMappingsStub != nil {
		return fake.ReadTcpMappingsStub()
	} else {
		return fake.readTcpMappingsReturns.result1, fake.readTcpMappingsReturns.result2
	}
}

func (fake *FakeDB) ReadTcpMappingsCallCount() int {
	fake.readTcpMappingsMutex.RLock()
	defer fake.readTcpMappingsMutex.RUnlock()
	return len(fake.readTcpMappingsArgsForCall)
}

func (fake *FakeDB) ReadTcpMappingsReturns(result1 []db.TcpMapping, result2 error) {
	fake.ReadTcpMappingsStub = nil
	fake.readTcpMappingsReturns = struct {
		result1 []db.TcpMapping
		result2 error
	}{result1, result2}
}

func (fake *FakeDB) SaveTcpMapping(tcpMapping db.TcpMapping) error {
	fake.saveTcpMappingMutex.Lock()
	fake.saveTcpMappingArgsForCall = append(fake.saveTcpMappingArgsForCall, struct {
		tcpMapping db.TcpMapping
	}{tcpMapping})
	fake.saveTcpMappingMutex.Unlock()
	if fake.SaveTcpMappingStub != nil {
		return fake.SaveTcpMappingStub(tcpMapping)
	} else {
		return fake.saveTcpMappingReturns.result1
	}
}

func (fake *FakeDB) SaveTcpMappingCallCount() int {
	fake.saveTcpMappingMutex.RLock()
	defer fake.saveTcpMappingMutex.RUnlock()
	return len(fake.saveTcpMappingArgsForCall)
}

func (fake *FakeDB) SaveTcpMappingArgsForCall(i int) db.TcpMapping {
	fake.saveTcpMappingMutex.RLock()
	defer fake.saveTcpMappingMutex.RUnlock()
	return fake.saveTcpMappingArgsForCall[i].tcpMapping
}

func (fake *FakeDB) SaveTcpMappingReturns(result1 error) {
	fake.SaveTcpMappingStub = nil
	fake.saveTcpMappingReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeDB) Connect() error {
	fake.connectMutex.Lock()
	fake.connectArgsForCall = append(fake.connectArgsForCall, struct{}{})
	fake.connectMutex.Unlock()
	if fake.ConnectStub != nil {
		return fake.ConnectStub()
	} else {
		return fake.connectReturns.result1
	}
}

func (fake *FakeDB) ConnectCallCount() int {
	fake.connectMutex.RLock()
	defer fake.connectMutex.RUnlock()
	return len(fake.connectArgsForCall)
}

func (fake *FakeDB) ConnectReturns(result1 error) {
	fake.ConnectStub = nil
	fake.connectReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeDB) Disconnect() error {
	fake.disconnectMutex.Lock()
	fake.disconnectArgsForCall = append(fake.disconnectArgsForCall, struct{}{})
	fake.disconnectMutex.Unlock()
	if fake.DisconnectStub != nil {
		return fake.DisconnectStub()
	} else {
		return fake.disconnectReturns.result1
	}
}

func (fake *FakeDB) DisconnectCallCount() int {
	fake.disconnectMutex.RLock()
	defer fake.disconnectMutex.RUnlock()
	return len(fake.disconnectArgsForCall)
}

func (fake *FakeDB) DisconnectReturns(result1 error) {
	fake.DisconnectStub = nil
	fake.disconnectReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeDB) WatchRouteChanges() (<-chan storeadapter.WatchEvent, chan<- bool, <-chan error) {
	fake.watchRouteChangesMutex.Lock()
	fake.watchRouteChangesArgsForCall = append(fake.watchRouteChangesArgsForCall, struct{}{})
	fake.watchRouteChangesMutex.Unlock()
	if fake.WatchRouteChangesStub != nil {
		return fake.WatchRouteChangesStub()
	} else {
		return fake.watchRouteChangesReturns.result1, fake.watchRouteChangesReturns.result2, fake.watchRouteChangesReturns.result3
	}
}

func (fake *FakeDB) WatchRouteChangesCallCount() int {
	fake.watchRouteChangesMutex.RLock()
	defer fake.watchRouteChangesMutex.RUnlock()
	return len(fake.watchRouteChangesArgsForCall)
}

func (fake *FakeDB) WatchRouteChangesReturns(result1 <-chan storeadapter.WatchEvent, result2 chan<- bool, result3 <-chan error) {
	fake.WatchRouteChangesStub = nil
	fake.watchRouteChangesReturns = struct {
		result1 <-chan storeadapter.WatchEvent
		result2 chan<- bool
		result3 <-chan error
	}{result1, result2, result3}
}

var _ db.DB = new(FakeDB)
