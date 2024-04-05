package cli

import (
	"fmt"

	"github.com/matheus-santos-souza/go-hexagonal-architecture/application"
)

func Run(
	productService application.ProductServiceInterface,
	action string,
	productId string,
	productName string,
	price float64,
) (string, error) {
	var result = ""

	switch action {
	case "create":
		product, err := productService.Create(productName, price)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf(
			"Product ID %s, white the name %s has been created with the price %f and status %s",
			product.GetID(), product.GetName(), product.GetPrice(), product.GetStatus(),
		)
	case "enable":
		product, err := productService.Get(productId)
		if err != nil {
			return result, err
		}
		res, err := productService.Enable(product)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf(
			"Product %s has been enabled",
			res.GetName(),
		)
	case "disable":
		product, err := productService.Get(productId)
		if err != nil {
			return result, err
		}
		res, err := productService.Disable(product)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf(
			"Product %s has been disabled",
			res.GetName(),
		)
	default:
		product, err := productService.Get(productId)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf(
			"Product ID %s, white the name %s has been get with the price %f and status %s",
			product.GetID(), product.GetName(), product.GetPrice(), product.GetStatus(),
		)
	}
	return result, nil
}
