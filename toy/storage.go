package toy

type Toy struct {
	Store map[string]string
}

func New() (*Toy) {
	return &Toy{
		Store: make(map[string]string),
	}
}