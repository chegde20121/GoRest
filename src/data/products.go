package data

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

// Product defines the structure for an API product
type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	SKU         string  `json:"sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

// Products is a collection of Product
type Products []*Product

// ToJSON serializes the contents of the collection to JSON
// NewEncoder provides better performance than json.Unmarshal as it does not
// have to buffer the output into an in memory slice of bytes
// this reduces allocations and the overheads of the service
//
// https://golang.org/pkg/encoding/json/#NewEncoder
func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

// GetProducts returns a list of products
func GetProducts() Products {
	return productList
}

// FromJson decodes json data
func (p *Product) FormJson(reader io.Reader) error {
	e := json.NewDecoder(reader)
	return e.Decode(p)
}

// AddProducts adds new Procuct to list
func AddProduct(product *Product) {
	product.ID = getNextID()
	productList = append(productList, product)
}

// getNextID returns next ID for the product
func getNextID() int {
	if len(productList) == 0 {
		return 1
	}
	lastProd := productList[len(productList)-1]
	return lastProd.ID + 1
}

var ErrProductNotFound = fmt.Errorf("Product not found")

// findProduct returns product by ID
// If product is not found returns error
func findProductById(id int) (*Product, int, error) {
	for i, p := range productList {
		if p.ID == id {
			return p, i, nil
		}
	}
	return nil, -1, ErrProductNotFound
}

// UpdateProduct updates the product by ID,
// If product is not found returns error
func UpdateProduct(id int, product *Product) error {
	_, pos, err := findProductById(id)
	if err != nil {
		return err
	}
	product.ID = id
	productList[pos] = product
	return nil
}

// DeleteProduct deletes the product by ID,
// If product is not found returns error
func DeleteProduct(id int) error {
	_, pos, err := findProductById(id)
	if err != nil {
		return err
	}
	productList = append(productList[:pos], productList[pos+1:]...)
	return nil
}

// productList is a hard coded list of products for this
// dummy data source
var productList = []*Product{
	&Product{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy milky coffee",
		Price:       2.45,
		SKU:         "abc323",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          2,
		Name:        "Espresso",
		Description: "Short and strong coffee without milk",
		Price:       1.99,
		SKU:         "fjd34",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
