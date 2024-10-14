package repository

type repository struct {
	url string
}

func (r *repository) GetName() string {
	//do sql zapros
	return "name"
}
