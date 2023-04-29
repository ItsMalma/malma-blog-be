package exception

import (
	"encoding/json"

	"github.com/ItsMalma/gomal"
)

type ValidatorErrors map[string][]string

func TransformValidationResults(results []gomal.ValidationResult) error {
	if results == nil || len(results) < 1 {
		return nil
	}
	err := ValidatorErrors{}
	for _, result := range results {
		err[result.Name] = result.Messages
	}
	return err
}

func (e ValidatorErrors) Error() string {
	enc, err := json.Marshal(e)
	if err != nil {
		panic(err)
	}

	return string(enc)
}

type ValidatorError string

func (e ValidatorError) Error() string {
	return string(e)
}
