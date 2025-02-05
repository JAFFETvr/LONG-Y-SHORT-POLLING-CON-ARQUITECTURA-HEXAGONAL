package entities

type Client struct {
	Id        int32
	Name      string
	Direccion string
}

func NewClient(name string, direccion string) *Client {
	return &Client{Id: 1, Name: name, Direccion: direccion}
}

func (c *Client) GetName() string {
	return c.Name
}

func (c *Client) SetName(name string) {
	c.Name = name
}

func (c *Client) GetDireccion() string {
	return c.Direccion
}

func (c *Client) SetDireccion(direccion string) {
	c.Direccion = direccion
}
