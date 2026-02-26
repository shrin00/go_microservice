package handler

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/shrin00/go_microservice/data"
)

const (
	KEYPRODUCT = "key_product"
)

type Product struct {
	l *log.Logger
}

func NewProduct(l *log.Logger) *Product {
	return &Product{l}
}

// func (p *Product) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	switch r.Method {
// 	case http.MethodGet:
// 		p.getProducts(w, r)
// 	case http.MethodPost:
// 		// add a new product
// 		p.addProduct(w, r)
// 	case http.MethodPut:
// 		// update an existing product
// 		re := regexp.MustCompile(`/([0-9]+)`)
// 		g := re.FindAllStringSubmatch(r.URL.Path, -1)
// 		p.l.Printf("G: %#v", g)
// 		if len(g) != 1 || len(g[0]) != 2 {
// 			http.Error(w, "Invalid URI", http.StatusBadRequest)
// 			return
// 		}
// 		// id := g[0][1]
// 		id, err := strconv.Atoi(g[0][1])
// 		if err != nil {
// 			http.Error(w, "Invalid Product ID", http.StatusBadRequest)
// 			return
// 		}
// 		p.l.Printf("ID: %d", id)
// 		p.updateProduct(id, w, r)
// 	default:
// 		w.WriteHeader(http.StatusMethodNotAllowed)
// 	}
// }

func (p *Product) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "expected a integer", http.StatusBadRequest)
	}
	p.l.Println("Handle PUT Product")
	prod := r.Context().Value(KEYPRODUCT).(*data.Product)
	// prod := &data.Product{}
	// err = prod.FromJSON(r.Body)
	// if err != nil {
	// 	http.Error(w, "Unable to unmarshal json", http.StatusBadRequest)
	// }
	p.l.Printf("Prod: %#v", *prod)
	err = data.UpdateProduct(id, prod)
	if err == data.ErrProductNotFound {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, "Unable to update product", http.StatusInternalServerError)
		return
	}

}

func (p *Product) AddProduct(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Product")
	prod := r.Context().Value(KEYPRODUCT).(*data.Product)
	p.l.Printf("Prod: %#v", *prod)
	data.AddProduct(prod)
}

func (p *Product) GetProducts(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Products")
	pl := data.GetProducts()
	err := pl.ToJSON(w)
	if err != nil {
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
	}
}

// type KeyProduct struct{} // prefered way to use key in context, can use string as key

func (p *Product) ValidateProductMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		prod := &data.Product{}
		err := prod.FromJSON(r.Body)
		if err != nil {
			http.Error(w, "Unable to unmarshal json", http.StatusBadRequest)
			return
		}
		ctx := context.WithValue(r.Context(), KEYPRODUCT, prod)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
