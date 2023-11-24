// Code generated by MockGen. DO NOT EDIT.
// Source: service.go
//
// Generated by this command:
//
//	mockgen -source=service.go -destination=service_mock.go -package=services
//
// Package services is a generated GoMock package.
package services

import (
	context "context"
	models "job-portal-api/internal/models"
	reflect "reflect"

	jwt "github.com/golang-jwt/jwt/v5"
	gomock "go.uber.org/mock/gomock"
)

// MockUserService is a mock of UserService interface.
type MockUserService struct {
	ctrl     *gomock.Controller
	recorder *MockUserServiceMockRecorder
}

// MockUserServiceMockRecorder is the mock recorder for MockUserService.
type MockUserServiceMockRecorder struct {
	mock *MockUserService
}

// NewMockUserService creates a new mock instance.
func NewMockUserService(ctrl *gomock.Controller) *MockUserService {
	mock := &MockUserService{ctrl: ctrl}
	mock.recorder = &MockUserServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserService) EXPECT() *MockUserServiceMockRecorder {
	return m.recorder
}

// AddCompanyDetails mocks base method.
func (m *MockUserService) AddCompanyDetails(ctx context.Context, companyData models.Company) (models.Company, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddCompanyDetails", ctx, companyData)
	ret0, _ := ret[0].(models.Company)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddCompanyDetails indicates an expected call of AddCompanyDetails.
func (mr *MockUserServiceMockRecorder) AddCompanyDetails(ctx, companyData any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddCompanyDetails", reflect.TypeOf((*MockUserService)(nil).AddCompanyDetails), ctx, companyData)
}

// AddJobDetails mocks base method.
func (m *MockUserService) AddJobDetails(ctx context.Context, jobData models.NewJobRequest, cid uint64) (models.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddJobDetails", ctx, jobData, cid)
	ret0, _ := ret[0].(models.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddJobDetails indicates an expected call of AddJobDetails.
func (mr *MockUserServiceMockRecorder) AddJobDetails(ctx, jobData, cid any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddJobDetails", reflect.TypeOf((*MockUserService)(nil).AddJobDetails), ctx, jobData, cid)
}

// Login mocks base method.
func (m *MockUserService) Login(ctx context.Context, email, password string) (jwt.RegisteredClaims, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", ctx, email, password)
	ret0, _ := ret[0].(jwt.RegisteredClaims)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Login indicates an expected call of Login.
func (mr *MockUserServiceMockRecorder) Login(ctx, email, password any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockUserService)(nil).Login), ctx, email, password)
}

// ProcessJobApplications mocks base method.
func (m *MockUserService) ProcessJobApplications(ctappData []models.NewUserApplication) ([]models.NewUserApplication, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProcessJobApplications", ctappData)
	ret0, _ := ret[0].([]models.NewUserApplication)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProcessJobApplications indicates an expected call of ProcessJobApplications.
func (mr *MockUserServiceMockRecorder) ProcessJobApplications(ctappData any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessJobApplications", reflect.TypeOf((*MockUserService)(nil).ProcessJobApplications), ctappData)
}

// Signup mocks base method.
func (m *MockUserService) Signup(ctx context.Context, userData models.NewUser) (models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Signup", ctx, userData)
	ret0, _ := ret[0].(models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Signup indicates an expected call of Signup.
func (mr *MockUserServiceMockRecorder) Signup(ctx, userData any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Signup", reflect.TypeOf((*MockUserService)(nil).Signup), ctx, userData)
}

// ViewAllCompanies mocks base method.
func (m *MockUserService) ViewAllCompanies(ctx context.Context) ([]models.Company, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ViewAllCompanies", ctx)
	ret0, _ := ret[0].([]models.Company)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ViewAllCompanies indicates an expected call of ViewAllCompanies.
func (mr *MockUserServiceMockRecorder) ViewAllCompanies(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ViewAllCompanies", reflect.TypeOf((*MockUserService)(nil).ViewAllCompanies), ctx)
}

// ViewAllJobs mocks base method.
func (m *MockUserService) ViewAllJobs(ctx context.Context) ([]models.Job, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ViewAllJobs", ctx)
	ret0, _ := ret[0].([]models.Job)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ViewAllJobs indicates an expected call of ViewAllJobs.
func (mr *MockUserServiceMockRecorder) ViewAllJobs(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ViewAllJobs", reflect.TypeOf((*MockUserService)(nil).ViewAllJobs), ctx)
}

// ViewCompanyDetails mocks base method.
func (m *MockUserService) ViewCompanyDetails(ctx context.Context, id uint64) (models.Company, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ViewCompanyDetails", ctx, id)
	ret0, _ := ret[0].(models.Company)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ViewCompanyDetails indicates an expected call of ViewCompanyDetails.
func (mr *MockUserServiceMockRecorder) ViewCompanyDetails(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ViewCompanyDetails", reflect.TypeOf((*MockUserService)(nil).ViewCompanyDetails), ctx, id)
}

// ViewJobById mocks base method.
func (m *MockUserService) ViewJobById(ctx context.Context, jid uint64) ([]models.Job, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ViewJobById", ctx, jid)
	ret0, _ := ret[0].([]models.Job)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ViewJobById indicates an expected call of ViewJobById.
func (mr *MockUserServiceMockRecorder) ViewJobById(ctx, jid any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ViewJobById", reflect.TypeOf((*MockUserService)(nil).ViewJobById), ctx, jid)
}

// ViewJobFromCompany mocks base method.
func (m *MockUserService) ViewJobFromCompany(cid uint64) ([]models.Job, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ViewJobFromCompany", cid)
	ret0, _ := ret[0].([]models.Job)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ViewJobFromCompany indicates an expected call of ViewJobFromCompany.
func (mr *MockUserServiceMockRecorder) ViewJobFromCompany(cid any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ViewJobFromCompany", reflect.TypeOf((*MockUserService)(nil).ViewJobFromCompany), cid)
}
