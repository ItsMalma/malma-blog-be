package exception

import (
	"encoding/json"

	"github.com/ItsMalma/gomal"
)

type ValidatorError map[string][]string

func TransformValidationResults(results []gomal.ValidationResult) error {
	if results == nil || len(results) < 1 {
		return nil
	}
	err := ValidatorError{}
	for _, result := range results {
		err[result.Name] = result.Messages
	}
	return err
}

func (e ValidatorError) Error() string {
	enc, err := json.Marshal(e)
	if err != nil {
		panic(err)
	}

	return string(enc)
}
