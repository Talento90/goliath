package app

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewContext(t *testing.T) {
	parentCtx := context.Background()

	ctx := NewContext(parentCtx)

	assert.NotEmpty(t, ctx.TraceID())

	_, err := uuid.Parse(ctx.TraceID())

	assert.NoError(t, err)
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

	assert.Equal(t, expectedTraceID, ctx.TraceID())

	userID, checkUser := ctx.UserID()
	assert.Equal(t, expectedUserID, userID)
	assert.True(t, checkUser)

	tenantID, checkTenant := ctx.TenantID()
	assert.Equal(t, expectedTenantID, tenantID)
	assert.True(t, checkTenant)
}

func TestFromEmptyContext(t *testing.T) {
	parentCtx := context.Background()
	ctx := FromContext(parentCtx)

	assert.NotEmpty(t, ctx.TraceID())

	userID, checkUser := ctx.UserID()
	assert.Equal(t, "", userID)
	assert.False(t, checkUser)

	tenantID, checkTenant := ctx.TenantID()
	assert.Equal(t, "", tenantID)
	assert.False(t, checkTenant)
}

func TestContextCancellation(t *testing.T) {
	parentCtx, cancel := context.WithCancel(context.Background())

	ctx := NewContext(parentCtx)

	cancel()

	assert.Equal(t, context.Canceled, ctx.Err())
}
