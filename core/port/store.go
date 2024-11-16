package port

import "tonible14012002/ascenda-test-cli/core/domain"

type Store interface {
	Save([]domain.Hotel) *domain.Error
	List(Ids []string, DesIds []int) []domain.Hotel
}
