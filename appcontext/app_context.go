package appcontext

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

// AppContext carries the context of the current execution.
type AppContext struct {
	// Context for cancellation
	ctx context.Context
}

// Context returns the context.Context
func (sc *AppContext) Context() context.Context {
	return sc.ctx
}

// UserID returns the user id
func (sc *AppContext) UserID() (string, bool) {
	userID := sc.ctx.Value(UserIDKey)

	if id, ok := userID.(string); ok {
		return id, ok
	}

	return "", false
}

// SetUserID sets the user id
func (sc *AppContext) SetUserID(userID string) AppContext {
	sc.ctx = context.WithValue(sc.ctx, UserIDKey, userID)

	return *sc
}

// TenantID returns the tenant id
func (sc *AppContext) TenantID() (string, bool) {
	tenantIDKey := sc.ctx.Value(TenantIDKey)

	if id, ok := tenantIDKey.(string); ok {
		return id, true
	}

	return "", false
}

// SetTenantID sets the user id
func (sc *AppContext) SetTenantID(tenantID string) AppContext {
	sc.ctx = context.WithValue(sc.ctx, TenantIDKey, tenantID)

	return *sc
}

// SetTraceID sets the trace id
func (sc *AppContext) SetTraceID(traceID string) AppContext {
	sc.ctx = context.WithValue(sc.ctx, TraceIDKey, traceID)

	return *sc
}

// TraceID returns the trace identifier for the current flow
func (sc *AppContext) TraceID() string {
	return sc.ctx.Value(TraceIDKey).(string)
}

// FromContext returns a new AppContext from a context.Context
func FromContext(ctx context.Context) AppContext {
	appCtx := NewAppContext(ctx)

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

// NewContext returns a new AppContext
func NewAppContext(ctx context.Context) AppContext {
	ctx = context.WithValue(ctx, TraceIDKey, uuid.NewString())
	return AppContext{ctx: ctx}
}
