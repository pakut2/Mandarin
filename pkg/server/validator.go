package server

import (
	"slices"

	"github.com/go-playground/validator/v10"
	"github.com/pakut2/mandarin/pkg/constants"
	"github.com/pakut2/mandarin/pkg/types"
)

var Validate *validator.Validate

func initValidator() {
	Validate = validator.New()

	Validate.RegisterValidation("supportedProvider", supportedProvider)
}

func supportedProvider(field validator.FieldLevel) bool {
	return slices.Contains(constants.SCHEDULE_PROVIDERS, types.ScheduleProvider(field.Field().String()))
}
