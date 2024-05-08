package app

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestNewContext(t *testing.T) {
	parentCtx := context.Background()

	ctx := NewContext(parentCtx)

	require.NotEmpty(t, ctx.TraceID())

	_, err := uuid.Parse(ctx.TraceID())

	require.NoError(t, err)
}

func TestFromContext(t *testing.T) {
	expectedTraceID := "asd-asd-123-asd"
	expectedUserID := "c6d7dc51-c2a5-4aed-91fc-6f151342f9e2"
	expectedTenantID := "c6ad12dc51-c2a5-asd-91fc-6f151342f9e2"
	parentCtx := context.Background()
	parentCtx = context.WithValue(parentCtx, TraceIDKey, expectedTraceID)
	parentCtx = context.WithValue(parentCtx, UserIDKey, expectedUserID)
	parentCtx = context.WithValue(parentCtx, TenantIDKey, expectedTenantID)

	ctx := FromContext(parentCtx)

	require.Equal(t, expectedTraceID, ctx.TraceID())

	userID, checkUser := ctx.UserID()
	require.Equal(t, expectedUserID, userID)
	require.True(t, checkUser)

	tenantID, checkTenant := ctx.TenantID()
	require.Equal(t, expectedTenantID, tenantID)
	require.True(t, checkTenant)
}

func TestFromEmptyContext(t *testing.T) {
	parentCtx := context.Background()
	ctx := FromContext(parentCtx)

	require.NotEmpty(t, ctx.TraceID())

	userID, checkUser := ctx.UserID()
	require.Equal(t, "", userID)
	require.False(t, checkUser)

	tenantID, checkTenant := ctx.TenantID()
	require.Equal(t, "", tenantID)
	require.False(t, checkTenant)
}

func TestContextCancellation(t *testing.T) {
	parentCtx, cancel := context.WithCancel(context.Background())

	ctx := NewContext(parentCtx)

	cancel()

	require.Equal(t, context.Canceled, ctx.Err())
}
