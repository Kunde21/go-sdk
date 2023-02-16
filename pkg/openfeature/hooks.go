package openfeature

import "context"

// Hook allows application developers to add arbitrary behavior to the flag evaluation lifecycle.
// They operate similarly to middleware in many web frameworks.
// https://github.com/open-feature/spec/blob/main/specification/hooks.md
type Hook interface {
	Before(hookContext HookContext, hookHints HookHints) (*EvaluationContext, error)
	After(hookContext HookContext, flagEvaluationDetails InterfaceEvaluationDetails, hookHints HookHints) (*EvaluationContext, error)
	Error(hookContext HookContext, err error, hookHints HookHints) *EvaluationContext
	Finally(hookContext HookContext, hookHints HookHints) *EvaluationContext
}

// HookHints contains a map of hints for hooks
type HookHints struct {
	mapOfHints map[string]interface{}
}

// NewHookHints constructs HookHints
func NewHookHints(mapOfHints map[string]interface{}) HookHints {
	return HookHints{mapOfHints: mapOfHints}
}

// Value returns the value at the given key in the underlying map.
// Maintains immutability of the map.
func (h HookHints) Value(key string) interface{} {
	return h.mapOfHints[key]
}

// HookContext defines the base level fields of a hook context
type HookContext struct {
	flagKey           string
	flagType          Type
	defaultValue      interface{}
	clientMetadata    ClientMetadata
	providerMetadata  Metadata
	evaluationContext EvaluationContext
}

// FlagKey returns the hook context's flag key
func (h HookContext) FlagKey() string {
	return h.flagKey
}

// FlagType returns the hook context's flag type
func (h HookContext) FlagType() Type {
	return h.flagType
}

// DefaultValue returns the hook context's default value
func (h HookContext) DefaultValue() interface{} {
	return h.defaultValue
}

// ClientMetadata returns the client's metadata
func (h HookContext) ClientMetadata() ClientMetadata {
	return h.clientMetadata
}

// ProviderMetadata returns the provider's metadata
func (h HookContext) ProviderMetadata() Metadata {
	return h.providerMetadata
}

// EvaluationContext returns the hook context's EvaluationContext
func (h HookContext) EvaluationContext() EvaluationContext {
	return h.evaluationContext
}

// Context returns the hook's go context.
func (h HookContext) Context() context.Context {
	return h.EvaluationContext().Context()
}

// NewHookContext constructs HookContext
// Allows for simplified hook test cases while maintaining immutability
func NewHookContext(
	flagKey string,
	flagType Type,
	defaultValue interface{},
	clientMetadata ClientMetadata,
	providerMetadata Metadata,
	evaluationContext EvaluationContext,
) HookContext {
	return HookContext{
		flagKey:           flagKey,
		flagType:          flagType,
		defaultValue:      defaultValue,
		clientMetadata:    clientMetadata,
		providerMetadata:  providerMetadata,
		evaluationContext: evaluationContext,
	}
}

// check at compile time that UnimplementedHook implements the Hook interface
var _ Hook = UnimplementedHook{}

// UnimplementedHook implements all hook methods with empty functions
// Include UnimplementedHook in your hook struct to avoid defining empty functions
// e.g.
//
//	type MyHook struct {
//	  UnimplementedHook
//	}
type UnimplementedHook struct{}

func (UnimplementedHook) After(HookContext, InterfaceEvaluationDetails, HookHints) (*EvaluationContext, error) {
	return nil, nil
}

func (UnimplementedHook) Before(HookContext, HookHints) (*EvaluationContext, error) { return nil, nil }
func (UnimplementedHook) Error(HookContext, error, HookHints) *EvaluationContext    { return nil }
func (UnimplementedHook) Finally(HookContext, HookHints) *EvaluationContext         { return nil }
