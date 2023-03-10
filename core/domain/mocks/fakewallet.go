// Code generated by MockGen. DO NOT EDIT.
// Source: core/domain/wallet.go

// Package mocks is a generated GoMock package.
package mocks

import (
	http "net/http"
	reflect "reflect"

	dto "github.com/EliasSantiago/app-bank-transactions/core/dto"
	gomock "github.com/golang/mock/gomock"
)

// MockWalletService is a mock of WalletService interface.
type MockWalletService struct {
	ctrl     *gomock.Controller
	recorder *MockWalletServiceMockRecorder
}

// MockWalletServiceMockRecorder is the mock recorder for MockWalletService.
type MockWalletServiceMockRecorder struct {
	mock *MockWalletService
}

// NewMockWalletService creates a new mock instance.
func NewMockWalletService(ctrl *gomock.Controller) *MockWalletService {
	mock := &MockWalletService{ctrl: ctrl}
	mock.recorder = &MockWalletServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWalletService) EXPECT() *MockWalletServiceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockWalletService) Create(response http.ResponseWriter, request *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Create", response, request)
}

// Create indicates an expected call of Create.
func (mr *MockWalletServiceMockRecorder) Create(response, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockWalletService)(nil).Create), response, request)
}

// MockWalletUseCase is a mock of WalletUseCase interface.
type MockWalletUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockWalletUseCaseMockRecorder
}

// MockWalletUseCaseMockRecorder is the mock recorder for MockWalletUseCase.
type MockWalletUseCaseMockRecorder struct {
	mock *MockWalletUseCase
}

// NewMockWalletUseCase creates a new mock instance.
func NewMockWalletUseCase(ctrl *gomock.Controller) *MockWalletUseCase {
	mock := &MockWalletUseCase{ctrl: ctrl}
	mock.recorder = &MockWalletUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWalletUseCase) EXPECT() *MockWalletUseCaseMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockWalletUseCase) Create(walletRequest *dto.CreateWalletRequest) (*dto.CreateWalletResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", walletRequest)
	ret0, _ := ret[0].(*dto.CreateWalletResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockWalletUseCaseMockRecorder) Create(walletRequest interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockWalletUseCase)(nil).Create), walletRequest)
}

// MockWalletRepository is a mock of WalletRepository interface.
type MockWalletRepository struct {
	ctrl     *gomock.Controller
	recorder *MockWalletRepositoryMockRecorder
}

// MockWalletRepositoryMockRecorder is the mock recorder for MockWalletRepository.
type MockWalletRepositoryMockRecorder struct {
	mock *MockWalletRepository
}

// NewMockWalletRepository creates a new mock instance.
func NewMockWalletRepository(ctrl *gomock.Controller) *MockWalletRepository {
	mock := &MockWalletRepository{ctrl: ctrl}
	mock.recorder = &MockWalletRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWalletRepository) EXPECT() *MockWalletRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockWalletRepository) Create(walletRequest *dto.CreateWalletStore) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", walletRequest)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockWalletRepositoryMockRecorder) Create(walletRequest interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockWalletRepository)(nil).Create), walletRequest)
}
