package analytics

type Sale struct {
	Product          string
	Category         string
	Quantity         int
	UnitPriceKopecks int64
	Cancelled        bool
}

type Report struct {
	RevenueKopecks    int64
	UnitsSold         int
	RevenueByCategory map[string]int64
	TopProduct        string
	InvalidSales      int
}

func AnalyzeSales(sales []Sale) Report {
	res := Report{
		0,
		0,
		make(map[string]int64),
		"",
		0,
	}

	totalSold := make(map[string]int)

	for _, sale := range sales {
		if sale.Cancelled {
			continue
		}

		if sale.Quantity <= 0 ||
			sale.UnitPriceKopecks <= 0 ||
			sale.Product == "" ||
			sale.Category == "" {
			res.InvalidSales++
			continue
		}

		totalSold[sale.Product] += sale.Quantity

		revenue := sale.UnitPriceKopecks * int64(sale.Quantity)

		res.RevenueKopecks += revenue

		res.RevenueByCategory[sale.Category] += revenue

		res.UnitsSold += sale.Quantity

	}

	for product, sold := range totalSold {
		if sold > totalSold[res.TopProduct] ||
			sold == totalSold[res.TopProduct] && res.TopProduct > product {
			res.TopProduct = product
		}
	}

	return res
}
