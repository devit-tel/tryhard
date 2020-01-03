package tryhard

//go:generate mockery -name=Store -output ./internal/mocks
type Store interface {
	AddCounter(id string)
	AddCounterWithTTL(id string, second int)
	GetCounter(id string) int
	ClearCounter(id string)
}
