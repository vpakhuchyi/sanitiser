package formatter

import (
	"fmt"
	"reflect"

	"github.com/vpakhuchyi/censor/internal/models"
)

// Float formats a value as a float.
// The value is formatted with up to 7 significant figures for float32 and up to 15 significant figures for float64.
// Note: this method panics if the provided value is not a float.
func (f *Formatter) Float(v models.Value) string {
	if v.Kind != reflect.Float32 && v.Kind != reflect.Float64 {
		panic("provided value is not a float")
	}

	if v.Kind == reflect.Float32 {
		return fmt.Sprintf(`%.7g`, v.Value)
	}

	return fmt.Sprintf(`%.15g`, v.Value)
}