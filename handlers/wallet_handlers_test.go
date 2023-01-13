package handlers_test

//import (
//	"encoding/json"
//	" hery-ciaputra/assignment-05-golang-backend/mocks"
//	" hery-ciaputra/assignment-05-golang-backend/models"
//	" hery-ciaputra/assignment-05-golang-backend/server"
//	"github.com/stretchr/testify/assert"
//	"net/http"
//	"net/http/httptest"
//	"strings"
//	"testing"
//	"time"
//)
//
//var (
//	transaction = models.Transaction{
//		Amount:          100000,
//		TransactionType: "DEBIT",
//		WalletID:        100001,
//		TargetWalletID:  100002,
//		Description:     "Bayar makan",
//		CreatedAt:       time.Now(),
//	}
//	transactions = []*models.Transaction{&transaction}
//)
//
//func TestHandler_GetWallet(t *testing.T) {
//	t.Run("Should return status 200 and get wallet info", func(t *testing.T) {
//		mocksService := new(mocks.WalletService)
//		mocksAuthService := new(mocks.AuthService)
//		config := server.RouterConfig{WalletService: mocksService, AuthService: mocksAuthService}
//		router := server.NewRouter(&config)
//		mocksService.On("GetWallet").Return(transactions, nil)
//		expectedBody, _ := json.Marshal(map[string]any{"data": transactions, "status": "OK", "statusCode": http.StatusOK})
//
//		w := httptest.NewRecorder()
//		req, _ := http.NewRequest(http.MethodGet, "/wallets", nil)
//		router.ServeHTTP(w, req)
//
//		assert.Equal(t, http.StatusOK, w.Code)
//		assert.Equal(t, string(expectedBody), w.Body.String())
//	})
//}
//
//func TestHandler_TopUpWallet(t *testing.T) {
//
//}
//
//func TestHandler_TransferWallet(t *testing.T) {
//
//}
//
//func TestHandler_GetTransactions(t *testing.T) {
//	t.Run("Should return status 200 and get transactions info", func(t *testing.T) {
//		mocksService := new(mocks.RegisterService)
//		config := server.RouterConfig{RegisterService: mocksService}
//		router := server.NewRouter(&config)
//		mocksService.On("GetTransactions").Return(transactions, nil)
//		expectedBody, _ := json.Marshal(registerRes)
//		jsonBody, _ := json.Marshal(user)
//
//		w := httptest.NewRecorder()
//		req, _ := http.NewRequest(http.MethodPost, "/register", strings.NewReader(string(jsonBody)))
//		router.ServeHTTP(w, req)
//
//		assert.Equal(t, http.StatusOK, w.Code)
//		assert.Equal(t, string(expectedBody), w.Body.String())
//	})
//}
