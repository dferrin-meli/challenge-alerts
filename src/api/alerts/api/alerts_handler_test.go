package api

import (
	"challenge/alerts/src/api/alerts/domain"
	"challenge/alerts/src/api/application/common"
	"challenge/alerts/src/api/mocks"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"gopkg.in/go-playground/assert.v1"
)

var validate = common.CreateValidator()

func Test_GetAllAlerts(t *testing.T) {
	alerts := []domain.AlertsDTO{
		{
			Type:        "Red",
			Description: "Incremental data access",
			CreatedAt:   "2023-04-12T00:00:00Z",
			Country:     "Colombia",
		},
	}
	testCases := []struct {
		name              string
		alertsServiceMock func(*gin.Context) func(m *mock.Mock)
		expectedError     common.ApiError
		expectedCode      int
		expectedResponse  string
	}{
		{
			name: "Error get alerts",
			alertsServiceMock: func(ctx *gin.Context) func(m *mock.Mock) {
				return func(m *mock.Mock) {
					m.On("GetAll", ctx).
						Return(nil, errors.New("error get alerts")).Once()
				}
			},
			expectedError: common.NewInternalServerApiError(
				"Error getting alerts",
				errors.New("error get alerts"),
			),
		},
		{
			name: "Get alerts successfully",
			alertsServiceMock: func(ctx *gin.Context) func(m *mock.Mock) {
				return func(m *mock.Mock) {
					m.On("GetAll", ctx).
						Return(alerts, nil).Once()
				}
			},
			expectedCode:     http.StatusOK,
			expectedResponse: "[{\"Type\":\"Red\",\"Description\":\"Incremental data access\",\"CreatedAt\":\"2023-04-12T00:00:00Z\",\"Country\":\"Colombia\"}]",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			ginTestContext, responseRecorder := common.CreateGinTestContext()
			ginTestContext.Request = httptest.NewRequest("GET", "/alerts", strings.NewReader(string("")))

			alertsServiceMock := new(mocks.AlertsService)
			testCase.alertsServiceMock(ginTestContext)(&alertsServiceMock.Mock)

			handler := NewAlertsHandler(
				validate,
				alertsServiceMock,
			)
			err := handler.GetAll(ginTestContext)

			assert.Equal(t, testCase.expectedError, err)
			alertsServiceMock.AssertExpectations(t)

			responseAsBytes, _ := io.ReadAll(responseRecorder.Body)
			responseAsJSON := string(responseAsBytes)
			assert.Equal(t, testCase.expectedResponse, responseAsJSON)

			if testCase.expectedError == nil {
				assert.Equal(t, testCase.expectedCode, ginTestContext.Writer.Status())
			}
		})
	}
}
