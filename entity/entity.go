package entity

import "fmt"

type Product struct {
	ID    int
	Name  string
	Stock int
}

func (p Product) StockStatus() string {
	var status string
	if p.Stock < 3 {
		status = "Stock Hampir Habis"
	} else if p.Stock < 10 {
		status = "Stock Terbatas"
	} else {
		status = fmt.Sprintf("%d", p.Stock)
	}
	return status
}
