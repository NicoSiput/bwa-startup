package campaign

type Service interface {
	GetAll() ([]Campaign, error)
}
