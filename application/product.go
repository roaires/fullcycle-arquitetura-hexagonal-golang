package application

import (
	"errors"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

type ProductInterface interface {
	IsValid() (bool, error)
	Enable() error
	Disable() error
	GetID() string
	GetName() string
	GetStatus() string
	GetPrice() float64
}

type ProductServiceInterface interface {
	Get(id string) (ProductInterface, error)

	Create(name string, price float64) (ProductInterface, error)
	Enable(product ProductInterface) (ProductInterface, error)
	Disable(product ProductInterface) (ProductInterface, error)
}

type ProductReader interface {
	Get(id string) (ProductInterface, error)
}

type ProductWriter interface {
	Save(product ProductInterface) (ProductInterface, error)
}

type ProductPersistenceInterface interface {
	ProductReader
	ProductWriter
}

const (
	DISABLED             = "Disable"
	ENABLED              = "Enable"
	ERROR_PRICE_ENABLE   = "O preço precisa ser maior que zero para habilitar o produto."
	ERROR_PRICE_DISABLE  = "O preço precisa ser igual a zero para desabilitar o produto."
	ERROR_PRICE_NEVATIVE = "O preço precisa ser maior que zero."
	ERROR_STATUS         = "O status ser igual a enable ou disable."
)

type Product struct {
	ID     string  `valid:"uuidv4"`
	Name   string  `valid:"required"`
	Price  float64 `valid:"float, optional"`
	Status string  `valid:"required"`
}

func NewProduct() *Product {
	product := Product{
		ID:     uuid.NewV4().String(),
		Status: DISABLED,
	}
	return &product
}

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

func (p *Product) IsValid() (bool, error) {
	if p.Status == "" {
		p.Status = DISABLED
	}

	if p.Status != ENABLED && p.Status != DISABLED {
		return false, errors.New(ERROR_STATUS)
	}

	if p.Price < 0 {
		return false, errors.New(ERROR_PRICE_NEVATIVE)
	}

	_, err := govalidator.ValidateStruct(p)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (p *Product) Enable() error {
	if p.Price > 0 {
		p.Status = ENABLED
		return nil
	}

	return errors.New(ERROR_PRICE_ENABLE)
}

func (p *Product) Disable() error {
	if p.Price == 0 {
		p.Status = DISABLED
		return nil
	}

	return errors.New(ERROR_PRICE_DISABLE)
}

func (p *Product) GetID() string {
	return p.ID
}

func (p *Product) GetName() string {
	return p.Name
}

func (p *Product) GetStatus() string {
	return p.Status
}

func (p *Product) GetPrice() float64 {
	return p.Price
}
