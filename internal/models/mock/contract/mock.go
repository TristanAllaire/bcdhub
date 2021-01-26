// Code generated by MockGen. DO NOT EDIT.
// Source: contract/repository.go

// Package mock_contract is a generated GoMock package.
package contract

import (
	contractModel "github.com/baking-bad/bcdhub/internal/models/contract"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockRepository is a mock of Repository interface
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// Get mocks base method
func (m *MockRepository) Get(arg0 map[string]interface{}) (contractModel.Contract, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].(contractModel.Contract)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockRepositoryMockRecorder) Get(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockRepository)(nil).Get), arg0)
}

// GetMany mocks base method
func (m *MockRepository) GetMany(arg0 map[string]interface{}) ([]contractModel.Contract, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMany", arg0)
	ret0, _ := ret[0].([]contractModel.Contract)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMany indicates an expected call of GetMany
func (mr *MockRepositoryMockRecorder) GetMany(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMany", reflect.TypeOf((*MockRepository)(nil).GetMany), arg0)
}

// GetRandom mocks base method
func (m *MockRepository) GetRandom() (contractModel.Contract, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRandom")
	ret0, _ := ret[0].(contractModel.Contract)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRandom indicates an expected call of GetRandom
func (mr *MockRepositoryMockRecorder) GetRandom() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRandom", reflect.TypeOf((*MockRepository)(nil).GetRandom))
}

// GetAddressesByNetworkAndLevel mocks base method
func (m *MockRepository) GetAddressesByNetworkAndLevel(arg0 string, arg1 int64) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAddressesByNetworkAndLevel", arg0, arg1)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAddressesByNetworkAndLevel indicates an expected call of GetAddressesByNetworkAndLevel
func (mr *MockRepositoryMockRecorder) GetAddressesByNetworkAndLevel(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAddressesByNetworkAndLevel", reflect.TypeOf((*MockRepository)(nil).GetAddressesByNetworkAndLevel), arg0, arg1)
}

// GetIDsByAddresses mocks base method
func (m *MockRepository) GetIDsByAddresses(arg0 []string, arg1 string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIDsByAddresses", arg0, arg1)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetIDsByAddresses indicates an expected call of GetIDsByAddresses
func (mr *MockRepositoryMockRecorder) GetIDsByAddresses(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIDsByAddresses", reflect.TypeOf((*MockRepository)(nil).GetIDsByAddresses), arg0, arg1)
}

// IsFA mocks base method
func (m *MockRepository) IsFA(arg0, arg1 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsFA", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsFA indicates an expected call of IsFA
func (mr *MockRepositoryMockRecorder) IsFA(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsFA", reflect.TypeOf((*MockRepository)(nil).IsFA), arg0, arg1)
}

// UpdateMigrationsCount mocks base method
func (m *MockRepository) UpdateMigrationsCount(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateMigrationsCount", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateMigrationsCount indicates an expected call of UpdateMigrationsCount
func (mr *MockRepositoryMockRecorder) UpdateMigrationsCount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateMigrationsCount", reflect.TypeOf((*MockRepository)(nil).UpdateMigrationsCount), arg0, arg1)
}

// GetByAddresses mocks base method
func (m *MockRepository) GetByAddresses(addresses []contractModel.Address) ([]contractModel.Contract, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByAddresses", addresses)
	ret0, _ := ret[0].([]contractModel.Contract)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByAddresses indicates an expected call of GetByAddresses
func (mr *MockRepositoryMockRecorder) GetByAddresses(addresses interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByAddresses", reflect.TypeOf((*MockRepository)(nil).GetByAddresses), addresses)
}

// GetTokens mocks base method
func (m *MockRepository) GetTokens(arg0, arg1 string, arg2, arg3 int64) ([]contractModel.Contract, int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTokens", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([]contractModel.Contract)
	ret1, _ := ret[1].(int64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetTokens indicates an expected call of GetTokens
func (mr *MockRepositoryMockRecorder) GetTokens(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTokens", reflect.TypeOf((*MockRepository)(nil).GetTokens), arg0, arg1, arg2, arg3)
}

// GetProjectsLastContract mocks base method
func (m *MockRepository) GetProjectsLastContract(contract *contractModel.Contract) ([]contractModel.Contract, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProjectsLastContract", contract)
	ret0, _ := ret[0].([]contractModel.Contract)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProjectsLastContract indicates an expected call of GetProjectsLastContract
func (mr *MockRepositoryMockRecorder) GetProjectsLastContract(contract interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProjectsLastContract", reflect.TypeOf((*MockRepository)(nil).GetProjectsLastContract), contract)
}

// GetSameContracts mocks base method
func (m *MockRepository) GetSameContracts(arg0 contractModel.Contract, arg1, arg2 int64) (contractModel.SameResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSameContracts", arg0, arg1, arg2)
	ret0, _ := ret[0].(contractModel.SameResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSameContracts indicates an expected call of GetSameContracts
func (mr *MockRepositoryMockRecorder) GetSameContracts(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSameContracts", reflect.TypeOf((*MockRepository)(nil).GetSameContracts), arg0, arg1, arg2)
}

// GetSimilarContracts mocks base method
func (m *MockRepository) GetSimilarContracts(arg0 contractModel.Contract, arg1, arg2 int64) ([]contractModel.Similar, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSimilarContracts", arg0, arg1, arg2)
	ret0, _ := ret[0].([]contractModel.Similar)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetSimilarContracts indicates an expected call of GetSimilarContracts
func (mr *MockRepositoryMockRecorder) GetSimilarContracts(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSimilarContracts", reflect.TypeOf((*MockRepository)(nil).GetSimilarContracts), arg0, arg1, arg2)
}

// GetDiffTasks mocks base method
func (m *MockRepository) GetDiffTasks() ([]contractModel.DiffTask, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDiffTasks")
	ret0, _ := ret[0].([]contractModel.DiffTask)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDiffTasks indicates an expected call of GetDiffTasks
func (mr *MockRepositoryMockRecorder) GetDiffTasks() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDiffTasks", reflect.TypeOf((*MockRepository)(nil).GetDiffTasks))
}

// UpdateField mocks base method
func (m *MockRepository) UpdateField(where []contractModel.Contract, fields ...string) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{where}
	for _, a := range fields {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateField", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateField indicates an expected call of UpdateField
func (mr *MockRepositoryMockRecorder) UpdateField(where interface{}, fields ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{where}, fields...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateField", reflect.TypeOf((*MockRepository)(nil).UpdateField), varargs...)
}
