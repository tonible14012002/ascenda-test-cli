package formater

import (
	"tonible14012002/ascenda-test-cli/core/domain"
	"tonible14012002/ascenda-test-cli/core/ultils/mergeutils"
)

type CapitalizeFormatter struct {
	ApplySources []string
}

func NewCapitalizeInfoFormatter(names []string) *CapitalizeFormatter {
	return &CapitalizeFormatter{
		ApplySources: names,
	}
}

func (f *CapitalizeFormatter) IsApplicable(s string) bool {
	for _, v := range f.ApplySources {
		if v == s {
			return true
		}
	}
	return false
}

func (f *CapitalizeFormatter) FormatField(h domain.Hotel) domain.Hotel {
	h.Address = mergeutils.CapitalizeFirstLetters(h.Address)
	h.City = mergeutils.CapitalizeFirstLetters(h.City)
	h.Name = mergeutils.CapitalizeFirstLetters(h.Name)
	h.Country = mergeutils.CapitalizeFirstLetters(h.Country)
	return h
}
