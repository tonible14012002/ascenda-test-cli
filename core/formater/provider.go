package formater

import (
	"tonible14012002/ascenda-test-cli/core/domain"
	"tonible14012002/ascenda-test-cli/core/port"
)

type FormatProvider struct {
	formatter []port.Formatter
}

func NewFormatProvider(formatter []port.Formatter) *FormatProvider {
	return &FormatProvider{
		formatter: formatter,
	}
}

func (f *FormatProvider) Format(hotel domain.Hotel) domain.Hotel {
	for _, formatter := range f.formatter {
		if formatter.IsApplicable(hotel.GetSource()) {
			hotel = formatter.FormatField(hotel)
		}
	}
	return hotel

}
