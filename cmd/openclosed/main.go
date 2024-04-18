package main

import "fmt"

// OCP
// open for extension, closed for modification

type Size int
type Color int

const (
	red Color = iota
	green
	blue
)

const (
	small Size = iota
	medium
	large
)

type Product struct {
	name  string
	color Color
	size  Size
}

type Filter struct {
}

func (f *Filter) FilterByColor(products []Product, color Color) []*Product {
	result := make([]*Product, 0)
	for i, v := range products {
		if v.color == color {
			result = append(result, &products[i])
		}
	}
	return result
}

func (f *Filter) FilterBySize(products []Product, size Size) []*Product {
	result := make([]*Product, 0)
	for i, v := range products {
		if v.size == size {
			result = append(result, &products[i])
		}
	}
	return result
}

func (f *Filter) FilterBySizeAndColor(products []Product, size Size, color Color) []*Product {
	result := make([]*Product, 0)
	for i, v := range products {
		if v.size == size && v.color == color {
			result = append(result, &products[i])
		}
	}
	return result
}

type Specification interface {
	IsSatisfied(p *Product) bool
}

type ColorSpecification struct {
	color Color
}

func (c ColorSpecification) IsSatisfied(p *Product) bool {
	return p.color == c.color
}

type SizeSpecification struct {
	size Size
}

func (c SizeSpecification) IsSatisfied(p *Product) bool {
	return p.size == c.size
}

type BetterFilter struct{}

func (f *BetterFilter) Filter(products []Product, spec Specification) []*Product {
	result := make([]*Product, 0)
	for i, v := range products {
		if spec.IsSatisfied(&v) {
			result = append(result, &products[i])
		}
	}
	return result
}

type AndSpecification struct {
	first, second Specification
}

func (a AndSpecification) IsSatisfied(p *Product) bool {
	return a.first.IsSatisfied(p) && a.second.IsSatisfied(p)
}

func main() {
	apple := Product{"Apple", green, small}
	tree := Product{"Tree", green, large}
	house := Product{"House", red, small}

	products := []Product{apple, tree, house}
	fmt.Println("Green products (old):\n")
	f := Filter{}
	for _, v := range f.FilterByColor(products, green) {
		fmt.Println(" - ", v.name)
	}

	fmt.Printf("Green products (new):\n")
	greenSpec := ColorSpecification{green}
	sizeSpec := SizeSpecification{small}

	bf := BetterFilter{}
	for _, v := range bf.Filter(products, sizeSpec) {
		//for _, v := range bf.Filter(products, greenSpec) {
		fmt.Println(" - ", v.name)
	}

	largeSpec := SizeSpecification{large}

	lgSpec := AndSpecification{largeSpec, greenSpec}

	fmt.Println("Large Green products (new):\n")
	for _, v := range bf.Filter(products, lgSpec) {
		fmt.Println(" -  %s is large and green\n", v.name)
	}

}
