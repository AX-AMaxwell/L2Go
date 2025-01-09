package main

import (
	"fmt"
	"math"
	"reflect"
	"strconv"
)

func main() {
	var (
		i1 int = 0b00000001
		// 0 0 0 0 0 0 0 0 | 0 0 0 0 0 0 0 0 | 0 0 0 0 0 0 0 0 | 0 0 0 0 0 0 0 0
		// 0 0 0 0 0 0 0 0 | 0 0 0 0 0 0 0 0 | 0 0 0 0 0 0 0 0 | 0 0 0 0 0 0 0 1
		i2 int8 = 1
		// 0 0 0 0 0 0 0 1
		i3 int16 = 1
		// 0 0 0 0 0 0 0 0 | 0 0 0 0 0 0 0 1
		i4 int32 = 1
		// 0 0 0 0 0 0 0 0 | 0 0 0 0 0 0 0 0 | 0 0 0 0 0 0 0 0 | 0 0 0 0 0 0 0 1
		i5 int64 = 1
		// 0 0 0 0 0 0 0 0 | 0 0 0 0 0 0 0 0 | 0 0 0 0 0 0 0 0 | 0 0 0 0 0 0 0 0
		// 0 0 0 0 0 0 0 0 | 0 0 0 0 0 0 0 0 | 0 0 0 0 0 0 0 0 | 0 0 0 0 0 0 0 1
		ints = []any{i1, i2, i3, i4, i5}

		u1    uint   = 1
		u2    uint8  = 1
		u3    uint16 = 1
		u4    uint32 = 1
		u5    uint64 = 1
		uints        = []any{u1, u2, u3, u4, u5}
	)

	slice1 := []string{"a"}
	fmt.Printf("['%p']=>['%#v']\n", &slice1, slice1)
	slice1 = append(slice1, "b")
	fmt.Printf("['%p']=>['%#v']\n", &slice1, slice1)
	slice1 = modSlice(slice1)
	fmt.Printf("['%p']=>['%#v']\n", &slice1, slice1)

	var base2 = strconv.FormatInt(int64(70), 16)
	fmt.Printf("Base2: %v\n", base2)

	var x1 int = math.MaxUint8
	fmt.Println("X1:", reflect.ValueOf(x1).Type().Size())
	var x2 int = math.MaxUint8
	fmt.Println("X2:", reflect.ValueOf(x2).Type().Size())
	var x3 int = x1 + x2
	fmt.Println("X3:", reflect.ValueOf(x3).Type().Size())

	for _, v := range ints {
		r := reflect.ValueOf(v)
		fmt.Printf("['%T']=>[%d byte(s)]\n", v, r.Type().Size())
	}

	for _, v := range uints {
		r := reflect.ValueOf(v)
		fmt.Printf("['%T']=>[%d byte(s)]\n", v, r.Type().Size())
	}
}

func modSlice(v []string) []string {
	fmt.Printf("[ModSlice]: ['%p']=>['%#v']\n", &v, v)
	return append(v, "c")
}

/*

Decimal: 70

        70
	0 1 0 0 1 0 0 0

IP Address (4-bytes):
	192.168.255.1

			 192               168                255                1
	1 1 0 0 0 0 0 0 . 1 0 1 0 1 0 0 0 . 1 1 1 1 1 1 1 1 . 0 0 0 0 0 0 0 1

As uint8 (1-byte):
	0 1 0 0 0 0 0 0

As int8 (1-byte):
	0 1 0 0 0 0 0 0

*/
