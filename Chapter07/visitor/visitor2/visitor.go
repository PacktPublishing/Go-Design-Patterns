package main

import "fmt"

// INTERFACES --------------------------------------------------------------

type ProductInfoRetriever interface {
	GetPrice() float32
	GetName() string
}

type Visitor interface {
	Visit(ProductInfoRetriever)
}

type Visitable interface {
	Accept(Visitor)
}

// PRODUCT -----------------------------------------------------------------

type Product struct {
	Price float32
	Name  string
}

func (p *Product) GetPrice() float32 {
	return p.Price
}

func (p *Product) Accept(v Visitor) {
	v.Visit(p)
}

func (p *Product) GetName() string {
	return p.Name
}

// PRODUCTS ----------------------------------------------------------------

type Rice struct {
	Product
}

type Pasta struct {
	Product
}


type Fridge struct {
	Product
}

//GetPrice overrides GetPrice method of Product type
func (f *Fridge) GetPrice() float32 {
	return f.Product.Price + 20
}

//Accept overrides "Accept" method from Product and implements the Visitable
//interface
func (f *Fridge) Accept(v Visitor) {
	v.Visit(f)
}

// VISITOR -----------------------------------------------------------------

type PriceVisitor struct {
	Sum float32
}

func (pv *PriceVisitor) Visit(p ProductInfoRetriever) {
	pv.Sum += p.GetPrice()
}

type NamePrinter struct {
	ProductList string
}

func (n *NamePrinter) Visit(p ProductInfoRetriever) {
	n.ProductList = fmt.Sprintf("%s\n%s", p.GetName(), n.ProductList)
}

// MAIN -------------------------------------------------------------------

func main() {
	products := make([]Visitable, 3)
	products[0] = &Rice{
		Product: Product{
			Price: 32.0,
			Name:  "Some rice",
		},
	}
	products[1] = &Pasta{
		Product: Product{
			Price: 40.0,
			Name:  "Some pasta",
		},
	}
	products[2] = &Fridge{
		Product: Product{
			Price: 50,
			Name:  "A fridge",
		},
	}

	priceVisitor := &PriceVisitor{}

	for _, p := range products {
		p.Accept(priceVisitor)
	}

	fmt.Printf("Total: %f\n", priceVisitor.Sum)


	nameVisitor := &NamePrinter{}

	for _, p := range products {
		p.Accept(nameVisitor)
	}

	fmt.Printf("\nProduct list:\n-------------\n%s", nameVisitor.ProductList)
}
