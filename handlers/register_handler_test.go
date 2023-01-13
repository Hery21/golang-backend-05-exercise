package handlers_test

import (
	"encoding/json"
	"dto"
	"httperror"
	" hery-ciaputra/assignment-05-golang-backend/mocks"
	" hery-ciaputra/assignment-05-golang-backend/server"
	" hery-ciaputra/assignment-05-golang-backend/testutils"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var (
	registerReq = dto.RegisterReq{
		Email:    "user1@mail.com",
		Name:     "user1",
		Password: "",
	}

	registerRes = dto.RegisterRes{
		ID:    1,
		Name:  "user1",
		Email: "user1@mail.com",
	}
)

func TestHandler_Register(t *testing.T) {
	t.Run("Should return status 200 and registered user info", func(t *testing.T) {
		mocksService := new(mocks.RegisterService)
		config := server.RouterConfig{RegisterService: mocksService}
		payload := registerReq

		mocksService.On("Register", &payload).Return(&registerRes, nil)
		expectedBody, _ := json.Marshal(registerRes)

		requestBody := testutils.MakeRequestBody(payload)
		req, _ := http.NewRequest(http.MethodPost, "/register", requestBody)
		_, w := testutils.ServeReq(&config, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, string(expectedBody), w.Body.String())
	})

	t.Run("Should return status 500 when failed to fetch", func(t *testing.T) {
		mocksService := new(mocks.RegisterService)
		config := server.RouterConfig{RegisterService: mocksService}
		payload := registerReq

		mocksService.On("Register", &payload).Return(nil, httperror.InternalServerError("Internal Server Error"))
		expectedErr := testutils.ErrorToString(httperror.InternalServerError("Internal Server Error"))

		requestBody := testutils.MakeRequestBody(payload)
		req, _ := http.NewRequest(http.MethodPost, "/register", requestBody)
		_, w := testutils.ServeReq(&config, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Equal(t, expectedErr, w.Body.String())
	})
}
