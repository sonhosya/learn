package factorymode

import (
	"fmt"
	"testing"
)

func TestMakeClothes(t *testing.T) {
	anta, err := MakeClothes("ANTA")
	if err != nil {
		panic(fmt.Sprintf("make clothes ANTA failed, err: %s", err))
	}
	peak, err := MakeClothes("PEAK")
	if err != nil {
		panic(fmt.Sprintf("make clothes PEAK failed, err: %s", err))
	}
	fmt.Printf("ANTA: name: %s, size:%d \n", anta.GetName(), anta.GetSize())
	fmt.Printf("PEAK: name: %s, size:%d \n", peak.GetName(), peak.GetSize())
}
