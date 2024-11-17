package formater

import (
	"regexp"
	"tonible14012002/ascenda-test-cli/core/domain"
	"tonible14012002/ascenda-test-cli/core/ultils/mergeutils"
)

type ConditionBookingFormatter struct {
	ApplySources []string
}

func NewConditionBookingFormatter(names []string) *ConditionBookingFormatter {
	return &ConditionBookingFormatter{
		ApplySources: names,
	}
}

func (f *ConditionBookingFormatter) IsApplicable(s string) bool {
	for _, v := range f.ApplySources {
		if v == s {
			return true
		}
	}
	return false
}

func (f *ConditionBookingFormatter) FormatField(h domain.Hotel) domain.Hotel {
	// Regular expression to match "---" and any surrounding white space
	re := regexp.MustCompile(`\s*===\s*`)

	// Replace "---" and surrounding spaces with "\n"

	for i, a := range h.Condition {
		h.Condition[i] = re.ReplaceAllString(mergeutils.RemoveRedundantSpaces(a), "\n")
	}
	return h
}
