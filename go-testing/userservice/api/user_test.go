package api_test

import (
	"context"
	"encoding/json"
	"github.com/EmilGeorgiev/training/go-testing/userservice/api"
	"github.com/go-chi/chi/v5"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/EmilGeorgiev/training/go-testing/userservice"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateUser_WhenUserIsProvided_ThenItIsCreated(t *testing.T) {
	// Arrange
	user := userservice.User{ID: "id", Name: "Ivan", IdentificationNumber: "971010101496"}
	mockRepo := new(MockUserRepository)
	mockRepo.On("Create", userservice.NewUser{Name: "Ivan", IdentificationNumber: "971010101496"}).Return(user, nil)
	h := api.UserHandler{UserRepo: mockRepo}
	w := httptest.NewRecorder()

	// Act
	payload := `{"Name":"Ivan", "IdentificationNumber":"971010101496"}`
	req := httptest.NewRequest("POST", "/users", strings.NewReader(payload))
	h.Create(w, req)

	// Assert
	var got userservice.User
	_ = json.NewDecoder(w.Body).Decode(&got)

	assert.Equal(t, user, got)
	assert.EqualValues(t, http.StatusCreated, w.Code)
	mockRepo.AssertExpectations(t)
}

func TestCreateUser_WhenBadRequest_ThenNotCreated(t *testing.T) {

} //TODO

func TestGetUser_WhenRequested_ThenProvided(t *testing.T)  {
	// Arrange
	mockRepo := new(MockUserRepository)
	user := userservice.User{ID: "2", Name: "Ivan", IdentificationNumber: "971010101496"}
	h := api.UserHandler{UserRepo: mockRepo}
	mockRepo.On("Get", "2").Return(user, nil)

	request := httptest.NewRequest("GET", "/users/2", nil)
	w := httptest.NewRecorder()

	// Act
	//rctx := chi.NewRouteContext()
	//req := request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, rctx))
	//rctx.URLParams.Add("id", "2")
	//h.Get(w, req)

	router := chi.NewRouter()
	router.Get("/users/{id}", h.Get)
	router.ServeHTTP(w, request)

	// Assert
	var got userservice.User
	_ = json.NewDecoder(w.Body).Decode(&got)

	assert.Equal(t, user, got)
	assert.EqualValues(t, http.StatusOK, w.Code)
	mockRepo.AssertExpectations(t)
}

func TestGetUser_WhenNotExistingID_ThenUserNotFound(t *testing.T) {
	mockRepo := new(MockUserRepository)
	mockRepo.On("Get", "2").Return(userservice.User{ID: "id"}, nil)
	h := api.UserHandler{UserRepo: mockRepo}

	request := httptest.NewRequest("GET", "/users/2", nil)
	w := httptest.NewRecorder()
	// Act
	rctx := chi.NewRouteContext()
	req := request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, rctx))
	rctx.URLParams.Add("id", "2")
	h.Get(w, req)

	//router := chi.NewRouter()
	//router.Get("/users/{id}", h.Get)
	//router.ServeHTTP(w, request)

	// Assert
	var got userservice.User
	_ = json.NewDecoder(w.Body).Decode(&got)

	//assert.Equal(t, user, got)
	assert.EqualValues(t, http.StatusNotFound, w.Code)
	//mockRepo.AssertExpectations(t)
} //TODO

func TestUpdateUser_WhenRequested_ThenChanged(t *testing.T) {

} //TODO

func TestUpdateUser_WhenWrongRequest_ThenNotUpdated(t *testing.T) {

} //TODO

func TestDeleteUser_WhenRequested_ThenDeleted(t *testing.T) {

} // TODO

func TestDeleteUser_WhenWrongRequested_ThenNotDeleted(t *testing.T) {

} // TODO

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) Create(ctx context.Context, user userservice.NewUser) (userservice.User, error) {
	args := m.Called(user)
	return args.Get(0).(userservice.User), args.Error(1)
}

func (m *MockUserRepository) Get(ctx context.Context, id string) (userservice.User, error) {
	args := m.Called(id)
	return args.Get(0).(userservice.User), args.Error(1)
}

func (m *MockUserRepository) Update(ctx context.Context, id string, user userservice.NewUser) (userservice.User, error) {
	args := m.Called(id, user)
	return args.Get(0).(userservice.User), args.Error(1)
}

func (m *MockUserRepository) Delete(ctx context.Context, id string) error {
	args := m.Called(id)
	return args.Error(0)
}
