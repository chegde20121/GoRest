package handlers

import (
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/chegde20121/GoRest/src/data"
)

type Products struct {
	logger *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{logger: l}

}

// ServeHTTP handles the GET, POST, PUT, DELETE operations for the Products
func (p *Products) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		p.getProducts(rw, req)
		return
	}
	if req.Method == http.MethodPost {
		p.addProducts(rw, req)
		return
	}
	if req.Method == http.MethodPut {
		p.logger.Println("Validating product update request")
		// expect the id in the URI
		g, isValid := validateId(req, p, rw)
		if g == nil || !isValid {
			p.logger.Println("Failed to covert id to number")
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}
		idString := g[0][1]
		id, err := strconv.Atoi(idString)
		if err != nil {
			p.logger.Println("Failed to covert id to number")
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}
		p.updateProduct(id, rw, req)
		return
	}
	if req.Method == http.MethodDelete {
		g, isValid := validateId(req, p, rw)
		if g == nil || !isValid {
			p.logger.Println("Failed to covert id to number")
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}
		idString := g[0][1]
		id, err := strconv.Atoi(idString)
		if err != nil {
			p.logger.Println("Failed to covert id to number")
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}
		p.DeleteProduct(id, rw, req)
		return
	}
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

// validate Id , validates the id passed in uri
func validateId(req *http.Request, p *Products, rw http.ResponseWriter) ([][]string, bool) {
	reg := regexp.MustCompile(`products/([0-9]+)`)
	g := reg.FindAllStringSubmatch(req.URL.Path, -1)
	if len(g) != 1 {
		p.logger.Println("Invalid URI more than one id")
		http.Error(rw, "Invalid URI", http.StatusBadRequest)
		return nil, false
	}
	if len(g[0]) != 2 {
		p.logger.Println("Invalid URI more than one id")
		http.Error(rw, "Invalid URI", http.StatusBadRequest)
		return nil, false
	}
	return g, true
}

// getProducts returns list of available products
func (p *Products) getProducts(rw http.ResponseWriter, req *http.Request) {
	p.logger.Println("Fetching Products")
	products := data.GetProducts()
	err := products.ToJSON(rw)
	if err != nil {
		p.logger.Println("Failed to fetch products")
		http.Error(rw, "Unable to process request", http.StatusInternalServerError)
		return
	}
	p.logger.Println("Products Fetched Successfully :)")
}

// addProducts adds new product to product list
func (p *Products) addProducts(rw http.ResponseWriter, req *http.Request) {
	p.logger.Println("Adding New Product...")
	product := &data.Product{}
	err := product.FormJson(req.Body)
	if err != nil {
		p.logger.Println("Failed to add new product")
		http.Error(rw, "Failed to add new product", http.StatusBadRequest)
		return
	}
	data.AddProduct(product)
}

func (p *Products) updateProduct(id int, rw http.ResponseWriter, req *http.Request) {
	p.logger.Printf("Updating Product %d...", id)
	prod := &data.Product{}
	err := prod.FormJson(req.Body)
	if err != nil {
		p.logger.Println("Failed to parse request body")
		http.Error(rw, "Error in Parsing Payload", http.StatusBadRequest)
		return
	}
	err = data.UpdateProduct(id, prod)
	if err != nil {
		p.logger.Println("Product not founc")
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}
	p.logger.Printf("Successfully updated product %d", id)
}

// DeleteProduct deletes the product by id,
// if product is not found returns error
func (p *Products) DeleteProduct(id int, rw http.ResponseWriter, req *http.Request) {
	p.logger.Printf("Deleting Product %d...", id)
	err := data.DeleteProduct(id)
	if err != nil {
		p.logger.Println("Product not founc")
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}
	rw.WriteHeader(http.StatusAccepted)
}
