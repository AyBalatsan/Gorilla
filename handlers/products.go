package handlers

import (
	"log"
	"net/http"
	"restmicro/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		p.getProducts(rw, r)
		return
	case http.MethodPost:
		p.addProduct(rw, r)
		return
	case http.MethodPut:
		p.updateProduct(rw, r)
		return
	default:
		// catch all other methods
		rw.WriteHeader(http.StatusMethodNotAllowed)
	}

}

func (p *Products) getProducts(rw http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
}
func (p *Products) addProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("handler Post Product")

	prod := &data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
	}

	p.l.Println("Product: ", prod)
	data.AddProduct(prod)
}
func (p *Products) updateProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("handler Put Product")

	prod := &data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
	}

	p.l.Println("Product: ", prod)
	pl, err := data.UpdateProduct(prod.ID, prod)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
	err = pl.ToJSON(rw)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	p.l.Println("Products list:  ", prod)
}
