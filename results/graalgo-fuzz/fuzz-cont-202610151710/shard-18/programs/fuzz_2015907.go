package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
func genFn0(p0 string) int16 {
    v7 := []int{940}
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    v7 = append(v7, 127)
    p0 = (p0 + p0)
    p0 = "\n\t"
    _ = p0
    _ = v7
    return 0
}
func genFn1() uint64 {
    var v8 uint8 = 107
    v8 -= v8
    _ = v8
    return 4294967295
}

func main() {
    var v1 uint64 = 42
    v2 := []int{}
    var v3 float32 = (((float32(float64(1.5)) * 1.5) / 3.14159) / 1.0)
    v4 := make(map[string]uint32)
    v5 := make(map[uint32]int)
    var v6 bool = false
    v9 := GenS1{F0: 42}
    v1 -= v1
    var v10 uint32 = 2147483648
    v11 := v5[4294967295]
    for fi12 := 0; fi12 < 13; fi12++ {
        var v13 int8 = 42
        if (!true) {
            v1 = genFn1()
        }
        _ = fi12
        _ = v13
    }
    v14 := v5[411]
    _ = len(v2)
    v2 = append(v2, v11)
    var v15 int8 = 127
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
