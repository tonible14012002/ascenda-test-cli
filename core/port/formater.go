package port

import "tonible14012002/ascenda-test-cli/core/domain"

type FormatProvider struct {
	Formatter []Formatter
}
type Formatter interface {
	IsApplicable(string) bool
	FormatField(domain.Hotel) domain.Hotel
}

func NewFormatProvider(formatter []Formatter) *FormatProvider {
	return &FormatProvider{
		Formatter: formatter,
	}
}

func (f *FormatProvider) Format(hotel domain.Hotel) domain.Hotel {
	for _, formatter := range f.Formatter {
		if formatter.IsApplicable(hotel.GetSource()) {
			hotel = formatter.FormatField(hotel)
		}
	}
	return hotel

}
