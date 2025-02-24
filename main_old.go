// go:build ignore
package main

import "fmt"

type Product struct {
	Name, Category string
	Price          float64
}

type Supplier struct {
	Name, City string
}

type ProductList []Product
type ProductListP []*Product

func (products ProductList) printNames() {
	for _, p := range products {
		fmt.Println("Name:", p.Name)
	}
}

func (products *ProductListP) printNames() {
	for _, p := range *products {
		fmt.Println("Name:", p.Name)
	}
}

func (p Product) test() {
	fmt.Println("Name:", p.Name, "Category:", p.Category, "Price:", p.Price)
}

func (s *Supplier) printDetails() {
	fmt.Println("Supplier:", s.Name, "City:", s.City)
}

func printDetails(product *Product) {
	fmt.Println("Name:", product.Name, "Category:", product.Category, "Price", product.Price)
}

func newProduct(name, category string, price float64) *Product {
	return &Product{name, category, price}
}

func (p *Product) printDetails() {
	fmt.Println("Name:", p.Name, "Category:", p.Category, "Price", p.calcTax(0.2, 10))
}

func (p *Product) calcTax(rate, threshold float64) float64 {
	if p.Price > threshold {
		return p.Price + (p.Price * rate)
	}
	return p.Price
}

func main() {
	products := []*Product{
		newProduct("Kayak", "Watersports", 275),
		newProduct("Lifejacket", "Watersports", 48.95),
		newProduct("Soccer Ball", "Soccer", 19.50),
	}

	printDetails(&Product{Name: "Foo", Category: "Bar", Price: 123})

	for _, p := range products {
		p.printDetails()
	}

	suppliers := []*Supplier{
		&Supplier{Name: "Acme Co", City: "New York"},
		&Supplier{Name: "Suppliers R Us", City: "Paris"},
		&Supplier{Name: "Our Motorcycles", City: "New Delhi"},
	}

	for _, s := range suppliers {
		s.printDetails()
	}

	// chamando via o type e passando o receiver explicitamente
	(*Supplier).printDetails(&Supplier{Name: "Acme Co", City: "New York"})
	Product.test(Product{Name: "Kayak", Category: "Watersports", Price: 275})

	productList := ProductList{
		Product{Name: "Name is Kayak", Category: "Watersports", Price: 275},
		Product{Name: "Name is Lifejacket", Category: "Watersports", Price: 48.95},
		Product{Name: "Name is Soccer Ball", Category: "Soccer", Price: 19.50},
	}

	productListP := ProductListP{
		&productList[0], &productList[1], &productList[2],
	}

	productList.printNames()
	productListP.printNames()

}
