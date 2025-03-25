package validator

import "testing"

func TestValidator_IsValid(t *testing.T) {
	tests := []struct {
		name      string
		validator *Validator
		want      bool
	}{
		{"no errors returns true", &Validator{Errors: make(map[string]string)}, true},
		{"with errors returns false", &Validator{Errors: map[string]string{"field": "error"}}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.validator.IsValid(); got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidator_AddError(t *testing.T) {
	tests := []struct {
		name      string
		validator *Validator
		key       string
		message   string
		want      map[string]string
	}{
		{"add new error", New(), "field", "error", map[string]string{"field": "error"}},
		{"add existing error", &Validator{Errors: map[string]string{"field": "existing error"}}, "field", "new error", map[string]string{"field": "existing error"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.validator.AddError(tt.key, tt.message)
			if got := tt.validator.Errors; !equal(got, tt.want) {
				t.Errorf("Validator.AddError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidator_In(t *testing.T) {
	tests := []struct {
		name  string
		value string
		list  []string
		want  bool
	}{
		{"value in list", "apple", []string{"apple", "banana", "cherry"}, true},
		{"value not in list", "grape", []string{"apple", "banana", "cherry"}, false},
		{"empty list", "apple", []string{}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := New()
			if got := v.In(tt.value, tt.list...); got != tt.want {
				t.Errorf("Validator.In() = %v, want %v", got, tt.want)
			}
		})
	}
}

func equal(a, b map[string]string) bool {
	if len(a) != len(b) {
		return false
	}
	for k, v := range a {
		if b[k] != v {
			return false
		}
	}
	return true
}
