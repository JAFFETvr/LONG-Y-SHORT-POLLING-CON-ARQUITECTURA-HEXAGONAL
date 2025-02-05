package entities

type Product struct {
	Id    int32
	Name  string
	Price float32
}

func NewProduct(name string, price float32) *Product {
	return &Product{Id: 1, Name: name, Price: price}
}

func (p *Product) GetName() string {
	return p.Name
}

func (p *Product) SetName(name string) {
	p.Name = name
}

func (p *Product) SetPrice(price float32) {
	p.Price = price
}