// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package handlers

import (
	"context"
	"github.com/mikenai/gowork/internal/models"
	"sync"
)

// Ensure, that UsersServiceMock does implement UsersService.
// If this is not the case, regenerate this file with moq.
var _ UsersService = &UsersServiceMock{}

// UsersServiceMock is a mock implementation of UsersService.
//
//	func TestSomethingThatUsesUsersService(t *testing.T) {
//
//		// make and configure a mocked UsersService
//		mockedUsersService := &UsersServiceMock{
//			CreateFunc: func(ctx context.Context, name string) (models.User, error) {
//				panic("mock out the Create method")
//			},
//			GetOneFunc: func(ctx context.Context, id string) (models.User, error) {
//				panic("mock out the GetOne method")
//			},
//		}
//
//		// use mockedUsersService in code that requires UsersService
//		// and then make assertions.
//
//	}
type UsersServiceMock struct {
	// CreateFunc mocks the Create method.
	CreateFunc func(ctx context.Context, name string) (models.User, error)

	// GetOneFunc mocks the GetOne method.
	GetOneFunc func(ctx context.Context, id string) (models.User, error)

	// DeleteFunc mocks the Delete method.
	DeleteFunc func(ctx context.Context, id string) error

	// calls tracks calls to the methods.
	calls struct {
		// Create holds details about calls to the Create method.
		Create []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Name is the name argument value.
			Name string
		}
		// GetOne holds details about calls to the GetOne method.
		GetOne []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ID is the id argument value.
			ID string
		}
		Delete []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ID is the id argument value.
			ID string
		}
	}
	lockCreate sync.RWMutex
	lockGetOne sync.RWMutex
	lockDelete sync.RWMutex
}

// Create calls CreateFunc.
func (mock *UsersServiceMock) Create(ctx context.Context, name string) (models.User, error) {
	if mock.CreateFunc == nil {
		panic("UsersServiceMock.CreateFunc: method is nil but UsersService.Create was just called")
	}
	callInfo := struct {
		Ctx  context.Context
		Name string
	}{
		Ctx:  ctx,
		Name: name,
	}
	mock.lockCreate.Lock()
	mock.calls.Create = append(mock.calls.Create, callInfo)
	mock.lockCreate.Unlock()
	return mock.CreateFunc(ctx, name)
}

// CreateCalls gets all the calls that were made to Create.
// Check the length with:
//
//	len(mockedUsersService.CreateCalls())
func (mock *UsersServiceMock) CreateCalls() []struct {
	Ctx  context.Context
	Name string
} {
	var calls []struct {
		Ctx  context.Context
		Name string
	}
	mock.lockCreate.RLock()
	calls = mock.calls.Create
	mock.lockCreate.RUnlock()
	return calls
}

// GetOne calls GetOneFunc.
func (mock *UsersServiceMock) GetOne(ctx context.Context, id string) (models.User, error) {
	if mock.GetOneFunc == nil {
		panic("UsersServiceMock.GetOneFunc: method is nil but UsersService.GetOne was just called")
	}
	callInfo := struct {
		Ctx context.Context
		ID  string
	}{
		Ctx: ctx,
		ID:  id,
	}
	mock.lockGetOne.Lock()
	mock.calls.GetOne = append(mock.calls.GetOne, callInfo)
	mock.lockGetOne.Unlock()
	return mock.GetOneFunc(ctx, id)
}

// GetOneCalls gets all the calls that were made to GetOne.
// Check the length with:
//
//	len(mockedUsersService.GetOneCalls())
func (mock *UsersServiceMock) GetOneCalls() []struct {
	Ctx context.Context
	ID  string
} {
	var calls []struct {
		Ctx context.Context
		ID  string
	}
	mock.lockGetOne.RLock()
	calls = mock.calls.GetOne
	mock.lockGetOne.RUnlock()
	return calls
}

// Delete calls DeleteFunc.
func (mock *UsersServiceMock) Delete(ctx context.Context, id string) error {
	callInfo := struct {
		Ctx context.Context
		ID  string
	}{
		Ctx: ctx,
		ID:  id,
	}
	mock.lockDelete.Lock()
	mock.calls.Delete = append(mock.calls.Delete, callInfo)
	mock.lockDelete.Unlock()
	return mock.DeleteFunc(ctx, id)
}

// DeleteCalls gets all the calls that were made to Delete.
// Check the length with:
//
//	len(mockedUsersService.DeleteCalls())
func (mock *UsersServiceMock) DeleteCalls() []struct {
	Ctx context.Context
	ID  string
} {
	var calls []struct {
		Ctx context.Context
		ID  string
	}
	mock.lockDelete.RLock()
	calls = mock.calls.Delete
	mock.lockDelete.RUnlock()
	return calls
}
