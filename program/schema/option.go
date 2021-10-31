package schema

import "fmt"

//CarExtOption defines the action for ext option
type CarExtOption interface {
	apply(*CarExt)
}

//EmptyExtOption implents CarExtOption interface
type EmptyExtOption struct{}

func (EmptyExtOption) apply(*CarExt) {}

//FuncExtOption implents CarExtOption interface
type FuncExtOption struct {
	f func(*CarExt)
}

func (this *FuncExtOption) apply(cext *CarExt) {
	this.f(cext)
}

func newFuncExtOption(f func(*CarExt)) *FuncExtOption {
	return &FuncExtOption{
		f: f,
	}
}

func WithPriceExtOption(price float64) CarExtOption {
	f := func(carExt *CarExt) {
		carExt.Price = price
	}
	return newFuncExtOption(f)
}

func WithYearExtOption(year int64) CarExtOption {
	f := func(carExt *CarExt) {
		carExt.Year = year
	}

	return newFuncExtOption(f)
}

//CarExt is ext Option for Car
type CarExt struct {
	Price float64
	Year  int64
}

func defaultCarExt() CarExt {
	return CarExt{
		Price: 0,
		Year:  1970,
	}
}

//Car is main subject
type Car struct {
	Name string
	Size int64
	Ext  CarExt
}

//NewCar is the construct func for Car
func NewCar(name string, size int64, opts ...CarExtOption) *Car {
	carInstance := &Car{
		Name: name,
		Size: size,
		Ext:  defaultCarExt(),
	}

	for _, opt := range opts {
		opt.apply(&carInstance.Ext)
	}

	return carInstance
}

func DemoMain() {
	// opts := []CarExtOption{
	// 	WithPriceExtOption(350000),
	// 	WithYearExtOption(2050),
	// }
	car := NewCar("bmw", 4, WithPriceExtOption(300000), WithYearExtOption(2020))
	fmt.Println(car)
}
