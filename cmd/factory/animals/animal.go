package animal

import "errors"

type Animal interface {
	Say() string
}

func Create(x int) Animal {
	switch x {
	case 1:
		return Dog{}
	case 2:
		return Cat{}
	}

	panic(errors.New("指定した数値に対応する動物がいませんでした"))
}
