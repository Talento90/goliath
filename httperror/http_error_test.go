package httperror

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/Talento90/goliath/app"
)

func TestNewProblemDetailWithapp(t *testing.T) {
	appCtx := app.NewContext(context.Background())
	err := app.NewInternalError("insuffient_funds", "No funds available").SetDetail("The account does not have enough funds to execute the transaction.")

	httpErr := New(appCtx, err, "/payments")
	require.Equal(t, err.Code(), httpErr.Type)
	require.Equal(t, err.Error(), httpErr.Title)
	require.Equal(t, err.Detail(), httpErr.Detail)
	require.Equal(t, "/payments", httpErr.Instance)
	require.Equal(t, 500, httpErr.Status)
	require.Equal(t, appCtx.TraceID(), httpErr.TraceID)

	require.Equal(t, "insuffient_funds: No funds available", httpErr.Error())
}

func TestNewProblemDetailStatusCodeMapping(t *testing.T) {
	tt := []struct {
		name               string
		errType            app.ErrorType
		expectedStatusCode int
	}{
		{
			name:               "map to internal error",
			errType:            app.Internal,
			expectedStatusCode: http.StatusInternalServerError,
		},
		{
			name:               "map to cancelled",
			errType:            app.Cancelled,
			expectedStatusCode: http.StatusAccepted,
		},
		{
			name:               "map to conflict",
			errType:            app.Conflict,
			expectedStatusCode: http.StatusConflict,
		},
		{
			name:               "map to not found",
			errType:            app.NotFound,
			expectedStatusCode: http.StatusNotFound,
		},
		{
			name:               "map to forbidden",
			errType:            app.Permission,
			expectedStatusCode: http.StatusForbidden,
		},
		{
			name:               "map to timeout",
			errType:            app.Timeout,
			expectedStatusCode: http.StatusRequestTimeout,
		},
		{
			name:               "map to unauthorised",
			errType:            app.Unauthorised,
			expectedStatusCode: http.StatusUnauthorized,
		},
		{
			name:               "map to bad request",
			errType:            app.Validation,
			expectedStatusCode: http.StatusBadRequest,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			appCtx := app.NewContext(context.Background())
			err := app.NewError("insuffient_funds", tc.errType, app.Low, "No funds available")

			httpErr := New(appCtx, err, "/payments")
			require.Equal(t, tc.expectedStatusCode, httpErr.Status)
		})
	}
}

func TestNewProblemDetailWithappAndValidationErrors(t *testing.T) {
	ctx := context.WithValue(context.Background(), app.TraceIDKey, "9b1b4579-b455-4eed-ac80-923668593dcc")
	appCtx := app.FromContext(ctx)
	err := app.NewValidationError("invalid_payment_data", "The payment request is invalid")
	err.AddValidationError(app.NewFieldValidationError("amount", "Amount needs to be positive"))
	err.AddValidationError(app.NewFieldValidationError("currency", "currency is required"))

	httpErr := New(appCtx, err, "/payments")

	require.Equal(t, err.Code(), httpErr.Type)
	require.Equal(t, err.Error(), httpErr.Title)
	require.Equal(t, "/payments", httpErr.Instance)
	require.Equal(t, 400, httpErr.Status)
	require.Equal(t, appCtx.TraceID(), httpErr.TraceID)
	require.Equal(t, err.ValidationErrors(), httpErr.Errors)

	errJSON, jsonErr := json.Marshal(httpErr)

	require.NoError(t, jsonErr)

	expectedJSON := "{\"type\":\"invalid_payment_data\",\"title\":\"The payment request is invalid\",\"status\":400,\"instance\":\"/payments\",\"traceId\":\"9b1b4579-b455-4eed-ac80-923668593dcc\",\"errors\":{\"amount\":[\"Amount needs to be positive\"],\"currency\":[\"currency is required\"]}}"

	require.Equal(t, expectedJSON, string(errJSON))
}

func TestNewProblemDetailWithGenericError(t *testing.T) {
	appCtx := app.NewContext(context.Background())
	err := errors.New("No funds available")

	httpErr := New(appCtx, err, "/payments")

	require.Equal(t, UnknownErrorType, httpErr.Type)
	require.Equal(t, "An error occurred, please contact support.", httpErr.Title)
	require.Equal(t, "", httpErr.Detail)
	require.Equal(t, "/payments", httpErr.Instance)
	require.Equal(t, 500, httpErr.Status)
	require.Equal(t, appCtx.TraceID(), httpErr.TraceID)
}
