package factorymode

import "fmt"

// 定义服装产品接口
type IClothes interface {
	setName(name string)
	setSize(size int)
	GetName() string
	GetSize() int
}

// 定义服装产品类
type clothes struct {
	name string
	size int
}

func (c *clothes) setName(name string) {
	c.name = name
}
func (c *clothes) setSize(size int) {
	c.size = size
}
func (c *clothes) GetName() string {
	return c.name
}
func (c *clothes) GetSize() int {
	return c.size
}

type PEAK struct {
	clothes
}

func newPEAK() IClothes {
	return &PEAK{
		clothes: clothes{
			name: "PEAK clothes",
			size: 1,
		},
	}
}

type ANTA struct {
	clothes
}

func newANTA() IClothes {
	return &PEAK{
		clothes: clothes{
			name: "ANTA clothes",
			size: 1,
		},
	}
}

func MakeClothes(clothesType string) (IClothes, error) {
	switch clothesType {
	case "ANTA":
		return newANTA(), nil
	case "PEAK":
		return newPEAK(), nil
	}
	return nil, fmt.Errorf("Wrong clothes type passed")
}
