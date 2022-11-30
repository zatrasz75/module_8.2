package describe

type UnitType string

const (
	Inch UnitType = "inch"
	CM   UnitType = "cm"
)

type Unit struct {
	Value float64
	T     UnitType
}

func (u Unit) Get(t UnitType) float64 {
	value := u.Value

	if t != u.T {
		// сконвертировать value в заданный в параметре UnitType
		if t == CM {
			value = value * 2.54
		}
		if t == Inch {
			value = value / 2.54
		}
	}
	return value
}

type Dimensions interface {
	Length() Unit //Длина
	Width() Unit  //Ширина
	Height() Unit //Высота
}

type Auto interface {
	Brand() string          //Брэнд
	Model() string          //Модель
	Dimensions() Dimensions //Размеры
	MaxSpeed() int          //Максимальная скорость
	EnginePower() int       //Мощность двигателя
}
type dimensions struct {
	unitType UnitType
	length   float64 //Длина
	width    float64 //Ширина
	height   float64 //Высота
}
type auto struct {
	dimensions  dimensions //Размеры
	brand       string     //Брэнд
	model       string     //Модель
	maxSpeed    int        //Максимальная скорость
	enginePower int        //Мощность двигателя
}

//-------Конструкторы методов

func NewAuto(unitType UnitType, brand, model string, length, width, height float64, maxSpeed, enginePower int) *auto {
	return &auto{
		dimensions{
			unitType,
			length,
			width,
			height,
		},
		brand, model, maxSpeed, enginePower,
	}
}
func (a auto) Brand() string          { return a.brand }
func (a auto) Model() string          { return a.model }
func (a auto) Dimensions() Dimensions { return a.dimensions }
func (a auto) MaxSpeed() int          { return a.maxSpeed }
func (a auto) EnginePower() int       { return a.enginePower }

func (a dimensions) Length() Unit {
	return Unit{
		Value: a.length,
		T:     a.unitType,
	}
}
func (a dimensions) Width() Unit {
	return Unit{
		Value: a.width,
		T:     a.unitType,
	}
}
func (a dimensions) Height() Unit {
	return Unit{
		Value: a.height,
		T:     a.unitType,
	}
}
