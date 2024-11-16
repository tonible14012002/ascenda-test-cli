package formater

import (
	"tonible14012002/ascenda-test-cli/core/domain"
	"tonible14012002/ascenda-test-cli/core/ultils/mergeutils"
)

type PascalToSentenceFormatter struct {
	ApplySources []string
}

func NewPascalToSentenceFormatter(Names []string) *PascalToSentenceFormatter {
	return &PascalToSentenceFormatter{
		ApplySources: Names,
	}
}

func (f *PascalToSentenceFormatter) IsApplicable(SourceName string) bool {
	for _, v := range f.ApplySources {
		if v == SourceName {
			return true
		}
	}
	return false
}

func (f *PascalToSentenceFormatter) FormatField(h domain.Hotel) domain.Hotel {
	for i, a := range h.Amenities.General {
		if a == "WiFi" {
			h.Amenities.General[i] = "wifi"
			continue
		}
		h.Amenities.General[i] = mergeutils.ToLower(mergeutils.PascalToSentence(mergeutils.RemoveRedundantSpaces(a)))
	}
	for i, a := range h.Amenities.Room {
		if a == "WiFi" {
			h.Amenities.Room[i] = "wifi"
			continue
		}
		h.Amenities.Room[i] = mergeutils.ToLower(mergeutils.PascalToSentence(mergeutils.RemoveRedundantSpaces(a)))
	}
	return h
}
