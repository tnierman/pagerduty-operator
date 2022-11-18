// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/pagerduty/service.go

// Package pagerduty is a generated GoMock package.
package pagerduty

import (
	reflect "reflect"

	go_pagerduty "github.com/PagerDuty/go-pagerduty"
	gomock "github.com/golang/mock/gomock"
)

// MockClient is a mock of Client interface.
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient.
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance.
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// ApplyServiceOrchestrationRule mocks base method.
func (m *MockClient) ApplyServiceOrchestrationRule(data *Data) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ApplyServiceOrchestrationRule", data)
	ret0, _ := ret[0].(error)
	return ret0
}

// ApplyServiceOrchestrationRule indicates an expected call of ApplyServiceOrchestrationRule.
func (mr *MockClientMockRecorder) ApplyServiceOrchestrationRule(data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApplyServiceOrchestrationRule", reflect.TypeOf((*MockClient)(nil).ApplyServiceOrchestrationRule), data)
}

// CreateService mocks base method.
func (m *MockClient) CreateService(data *Data) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateService", data)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateService indicates an expected call of CreateService.
func (mr *MockClientMockRecorder) CreateService(data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateService", reflect.TypeOf((*MockClient)(nil).CreateService), data)
}

// DeleteService mocks base method.
func (m *MockClient) DeleteService(data *Data) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteService", data)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteService indicates an expected call of DeleteService.
func (mr *MockClientMockRecorder) DeleteService(data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteService", reflect.TypeOf((*MockClient)(nil).DeleteService), data)
}

// DisableService mocks base method.
func (m *MockClient) DisableService(data *Data) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DisableService", data)
	ret0, _ := ret[0].(error)
	return ret0
}

// DisableService indicates an expected call of DisableService.
func (mr *MockClientMockRecorder) DisableService(data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DisableService", reflect.TypeOf((*MockClient)(nil).DisableService), data)
}

// EnableService mocks base method.
func (m *MockClient) EnableService(data *Data) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnableService", data)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnableService indicates an expected call of EnableService.
func (mr *MockClientMockRecorder) EnableService(data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnableService", reflect.TypeOf((*MockClient)(nil).EnableService), data)
}

// GetIntegrationKey mocks base method.
func (m *MockClient) GetIntegrationKey(data *Data) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIntegrationKey", data)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetIntegrationKey indicates an expected call of GetIntegrationKey.
func (mr *MockClientMockRecorder) GetIntegrationKey(data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIntegrationKey", reflect.TypeOf((*MockClient)(nil).GetIntegrationKey), data)
}

// GetService mocks base method.
func (m *MockClient) GetService(data *Data) (*go_pagerduty.Service, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetService", data)
	ret0, _ := ret[0].(*go_pagerduty.Service)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetService indicates an expected call of GetService.
func (mr *MockClientMockRecorder) GetService(data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetService", reflect.TypeOf((*MockClient)(nil).GetService), data)
}

// ToggleServiceOrchestration mocks base method.
func (m *MockClient) ToggleServiceOrchestration(data *Data, active bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ToggleServiceOrchestration", data, active)
	ret0, _ := ret[0].(error)
	return ret0
}

// ToggleServiceOrchestration indicates an expected call of ToggleServiceOrchestration.
func (mr *MockClientMockRecorder) ToggleServiceOrchestration(data, active interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ToggleServiceOrchestration", reflect.TypeOf((*MockClient)(nil).ToggleServiceOrchestration), data, active)
}

// UpdateEscalationPolicy mocks base method.
func (m *MockClient) UpdateEscalationPolicy(data *Data) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateEscalationPolicy", data)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateEscalationPolicy indicates an expected call of UpdateEscalationPolicy.
func (mr *MockClientMockRecorder) UpdateEscalationPolicy(data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateEscalationPolicy", reflect.TypeOf((*MockClient)(nil).UpdateEscalationPolicy), data)
}

// MockPdClient is a mock of PdClient interface.
type MockPdClient struct {
	ctrl     *gomock.Controller
	recorder *MockPdClientMockRecorder
}

// MockPdClientMockRecorder is the mock recorder for MockPdClient.
type MockPdClientMockRecorder struct {
	mock *MockPdClient
}

// NewMockPdClient creates a new mock instance.
func NewMockPdClient(ctrl *gomock.Controller) *MockPdClient {
	mock := &MockPdClient{ctrl: ctrl}
	mock.recorder = &MockPdClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPdClient) EXPECT() *MockPdClientMockRecorder {
	return m.recorder
}

// CreateIntegration mocks base method.
func (m *MockPdClient) CreateIntegration(serviceID string, integration go_pagerduty.Integration) (*go_pagerduty.Integration, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateIntegration", serviceID, integration)
	ret0, _ := ret[0].(*go_pagerduty.Integration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateIntegration indicates an expected call of CreateIntegration.
func (mr *MockPdClientMockRecorder) CreateIntegration(serviceID, integration interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateIntegration", reflect.TypeOf((*MockPdClient)(nil).CreateIntegration), serviceID, integration)
}

// CreateService mocks base method.
func (m *MockPdClient) CreateService(service go_pagerduty.Service) (*go_pagerduty.Service, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateService", service)
	ret0, _ := ret[0].(*go_pagerduty.Service)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateService indicates an expected call of CreateService.
func (mr *MockPdClientMockRecorder) CreateService(service interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateService", reflect.TypeOf((*MockPdClient)(nil).CreateService), service)
}

// DeleteService mocks base method.
func (m *MockPdClient) DeleteService(id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteService", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteService indicates an expected call of DeleteService.
func (mr *MockPdClientMockRecorder) DeleteService(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteService", reflect.TypeOf((*MockPdClient)(nil).DeleteService), id)
}

// GetEscalationPolicy mocks base method.
func (m *MockPdClient) GetEscalationPolicy(arg0 string, arg1 *go_pagerduty.GetEscalationPolicyOptions) (*go_pagerduty.EscalationPolicy, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEscalationPolicy", arg0, arg1)
	ret0, _ := ret[0].(*go_pagerduty.EscalationPolicy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEscalationPolicy indicates an expected call of GetEscalationPolicy.
func (mr *MockPdClientMockRecorder) GetEscalationPolicy(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEscalationPolicy", reflect.TypeOf((*MockPdClient)(nil).GetEscalationPolicy), arg0, arg1)
}

// GetIntegration mocks base method.
func (m *MockPdClient) GetIntegration(arg0, arg1 string, arg2 go_pagerduty.GetIntegrationOptions) (*go_pagerduty.Integration, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIntegration", arg0, arg1, arg2)
	ret0, _ := ret[0].(*go_pagerduty.Integration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetIntegration indicates an expected call of GetIntegration.
func (mr *MockPdClientMockRecorder) GetIntegration(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIntegration", reflect.TypeOf((*MockPdClient)(nil).GetIntegration), arg0, arg1, arg2)
}

// GetService mocks base method.
func (m *MockPdClient) GetService(arg0 string, arg1 *go_pagerduty.GetServiceOptions) (*go_pagerduty.Service, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetService", arg0, arg1)
	ret0, _ := ret[0].(*go_pagerduty.Service)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetService indicates an expected call of GetService.
func (mr *MockPdClientMockRecorder) GetService(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetService", reflect.TypeOf((*MockPdClient)(nil).GetService), arg0, arg1)
}

// ListIncidentAlertsWithOpts mocks base method.
func (m *MockPdClient) ListIncidentAlertsWithOpts(incidentId string, o go_pagerduty.ListIncidentAlertsOptions) (*go_pagerduty.ListAlertsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListIncidentAlertsWithOpts", incidentId, o)
	ret0, _ := ret[0].(*go_pagerduty.ListAlertsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListIncidentAlertsWithOpts indicates an expected call of ListIncidentAlertsWithOpts.
func (mr *MockPdClientMockRecorder) ListIncidentAlertsWithOpts(incidentId, o interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListIncidentAlertsWithOpts", reflect.TypeOf((*MockPdClient)(nil).ListIncidentAlertsWithOpts), incidentId, o)
}

// ListIncidents mocks base method.
func (m *MockPdClient) ListIncidents(arg0 go_pagerduty.ListIncidentsOptions) (*go_pagerduty.ListIncidentsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListIncidents", arg0)
	ret0, _ := ret[0].(*go_pagerduty.ListIncidentsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListIncidents indicates an expected call of ListIncidents.
func (mr *MockPdClientMockRecorder) ListIncidents(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListIncidents", reflect.TypeOf((*MockPdClient)(nil).ListIncidents), arg0)
}

// ListServices mocks base method.
func (m *MockPdClient) ListServices(arg0 go_pagerduty.ListServiceOptions) (*go_pagerduty.ListServiceResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListServices", arg0)
	ret0, _ := ret[0].(*go_pagerduty.ListServiceResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListServices indicates an expected call of ListServices.
func (mr *MockPdClientMockRecorder) ListServices(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListServices", reflect.TypeOf((*MockPdClient)(nil).ListServices), arg0)
}

// ManageEvent mocks base method.
func (m *MockPdClient) ManageEvent(e *go_pagerduty.V2Event) (*go_pagerduty.V2EventResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ManageEvent", e)
	ret0, _ := ret[0].(*go_pagerduty.V2EventResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ManageEvent indicates an expected call of ManageEvent.
func (mr *MockPdClientMockRecorder) ManageEvent(e interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ManageEvent", reflect.TypeOf((*MockPdClient)(nil).ManageEvent), e)
}

// UpdateService mocks base method.
func (m *MockPdClient) UpdateService(service go_pagerduty.Service) (*go_pagerduty.Service, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateService", service)
	ret0, _ := ret[0].(*go_pagerduty.Service)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateService indicates an expected call of UpdateService.
func (mr *MockPdClientMockRecorder) UpdateService(service interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateService", reflect.TypeOf((*MockPdClient)(nil).UpdateService), service)
}
