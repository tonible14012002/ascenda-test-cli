package port

import "tonible14012002/ascenda-test-cli/core/domain"

type Formatter interface {
	IsApplicable(string) bool
	FormatField(domain.Hotel) domain.Hotel
}
