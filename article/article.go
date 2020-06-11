package article

type Item struct {
	Title	string `json:"title"`
	Body	string `json:"body"`
}

type Articles struct {
	 Items []Item
}

func New() *Articles {
	return &Articles{}
}

func (r *Articles) Add(item Item) {
	r.Items = append(r.Items, item)
}

// func (r *Articles) GetAll() []Item {
	
// }