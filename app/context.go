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
	id, ok := userID.(string)
	return id, ok
}

// SetUserID sets the user id
func (sc *Context) SetUserID(userID string) *Context {
	sc.Context = context.WithValue(sc.Context, UserIDKey, userID)

	return sc
}

// TenantID returns the tenant id
func (sc *Context) TenantID() (string, bool) {
	tenantIDKey := sc.Context.Value(TenantIDKey)
	id, ok := tenantIDKey.(string)
	return id, ok
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
	traceID := sc.Context.Value(TraceIDKey)
	id, ok := traceID.(string)

	if !ok {
		return ""
	}

	return id
}

// NewContext returns a new Context from a context.Context
func NewContext(ctx context.Context) Context {
	appCtx := Context{Context: ctx}

	if _, ok := ctx.Value(TraceIDKey).(string); !ok {
		appCtx.SetTraceID(uuid.NewString())
	}

	return appCtx
}
