package formater

import (
	"strings"
	"tonible14012002/ascenda-test-cli/core/domain"
	"tonible14012002/ascenda-test-cli/core/ultils/mergeutils"
)

type DescFormatter struct {
	ApplySources []string
}

func NewDescFormatter(names []string) *DescFormatter {
	return &DescFormatter{
		ApplySources: names,
	}
}

func (f *DescFormatter) IsApplicable(s string) bool {
	for _, v := range f.ApplySources {
		if v == s {
			return true
		}
	}
	return false
}

func (f *DescFormatter) FormatField(h domain.Hotel) domain.Hotel {
	// FIXME: This is a temporary solution to remove redundant spaces
	h.Description = mergeutils.RemoveRedundantSpaces(h.Description)
	h.Description = strings.ReplaceAll(h.Description, "===", "\n")

	return h
}
