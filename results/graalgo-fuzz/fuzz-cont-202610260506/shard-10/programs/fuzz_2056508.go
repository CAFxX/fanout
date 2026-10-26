package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64 }

func main() {
    var v1 int32 = -2147483648
    var v2 int = 128
    var v3 uint32 = 809
    var v4 int8 = -93
    if (!false) {
        v4 = -39
    } else {
        var v5 uint32 = 1
        _ = v5
    }
    var v6 float32 = (1e10 * (0.0 + (3.14159 - ((((((3.14159 * float32(1e10)) / 1.0) - 0.0) + (0.0 - ((1.5 * 3.14159) * 1.5))) * 0.0) - 1.5))))
    if false {
        v2 += (v2 & v2)
        for fi7 := 0; fi7 < 9; fi7++ {
            v6 -= (0.0 - v6)
            _ = fi7
        }
    } else {
        v8 := []int{427}
        _ = v8
    }
    var v9 int8 = (v4 % ((0) | 1))
    fmt.Println(v1, v1)
    var v10 float32 = (float32(1e10) - (0.0 - 1.5))
    if true {
        v4 = v9
        v6 -= v10
        var v11 int16 = -917
        _ = v11
    } else {
        var v12 int64 = -1
        v13 := []int{}
        _ = v12
        _ = v13
    }
    v14 := GenS1{F0: "a", F1: 42}
    v15 := []int{v2}
    v2 += v2
    var v16 uint16 = 2
    v9 = (-(v4))
    if true {
        for fi17 := 0; fi17 < 4; fi17++ {
            v10 += (v10 / 1e10)
            var v18 uint16 = 42
            _ = fi17
            _ = v18
        }
        v9 += -128
    } else {
        var v19 float32 = (-2.25 - v6)
        _ = v19
    }
    var v20 string = "\n\t"
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v6
    _ = v9
}
