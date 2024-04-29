package app

import (
	"context"

	"github.com/google/uuid"
)

type ContextKey string

const (
	TraceIDKey  ContextKey = "trace_id"
	UserIDKey   ContextKey = "user_id"
	TenantIDKey ContextKey = "tenant_id"
)

// Context carries the context of the current execution.
type Context struct {
	// original context
	context.Context
}

// UserID returns the user id
func (sc *Context) UserID() (string, bool) {
	userID := sc.Context.Value(UserIDKey)

	if id, ok := userID.(string); ok {
		return id, ok
	}

	return "", false
}

// SetUserID sets the user id
func (sc *Context) SetUserID(userID string) *Context {
	sc.Context = context.WithValue(sc.Context, UserIDKey, userID)

	return sc
}

// TenantID returns the tenant id
func (sc *Context) TenantID() (string, bool) {
	tenantIDKey := sc.Context.Value(TenantIDKey)

	if id, ok := tenantIDKey.(string); ok {
		return id, true
	}

	return "", false
}

// SetTenantID sets the user id
func (sc *Context) SetTenantID(tenantID string) *Context {
	sc.Context = context.WithValue(sc.Context, TenantIDKey, tenantID)

	return sc
}

// SetTraceID sets the trace id
func (sc *Context) SetTraceID(traceID string) *Context {
	sc.Context = context.WithValue(sc.Context, TraceIDKey, traceID)

	return sc
}

// TraceID returns the trace identifier for the current flow
func (sc *Context) TraceID() string {
	return sc.Context.Value(TraceIDKey).(string)
}

// FromContext returns a new Context from a context.Context
func FromContext(ctx context.Context) Context {
	appCtx := NewContext(ctx)

	if traceID, ok := ctx.Value(TraceIDKey).(string); ok {
		appCtx.SetTraceID(traceID)
	}

	if userID, ok := ctx.Value(UserIDKey).(string); ok {
		appCtx.SetUserID(userID)
	}

	if tenantID, ok := ctx.Value(TenantIDKey).(string); ok {
		appCtx.SetTenantID(tenantID)
	}

	return appCtx
}

// NewContext returns a new Context
func NewContext(ctx context.Context) Context {
	ctx = context.WithValue(ctx, TraceIDKey, uuid.NewString())
	return Context{Context: ctx}
}
