package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64; F2 int8 }
func genFn0(p0 int64) uint16 {
    p0 += p0
    var v4 string = "\n\t"
    if true {
        p0 = (^(p0))
    }
    v5 := GenS1{F0: 0, F1: 4294967295, F2: -1}
    v4 = (v4 + "a")
    _ = p0
    _ = v4
    _ = v5
    return 255
}
func genFn1() uint32 {
    var v6 string = "foo bar"
    switch v6 {
    case "foo bar":
        var v7 int8 = -128
        _ = v7
    }
    v6 = v6
    for fi8 := 0; fi8 < 14; fi8++ {
        v6 = ("" + v6)
        var v9 uint32 = 530
        _ = fi8
        _ = v9
    }
    _ = v6
    return 42
}

func main() {
    var v1 int64 = 990
    v2 := make(map[int]int)
    var v3 float32 = float32((1e10 * ((((((((-2.25 + float64(float32(0.1))) + 3.14159) - 1e10) * float64(1.5)) - (1.5 - 0.1)) - (-2.25 / 1.0)) * 0.0) * 3.14159)))
    var v10 string = "foo bar"
    var v11 int32 = -2147483648
    var v12 uint = 851
    v13 := []int{42, 32768, -95, -2147483648}
    fmt.Println(v10, v11)
    switch 1 {
    case 2147483647:
        v14 := GenS1{F0: 65535, F1: 2, F2: 114}
        _ = v14
    }
    v12 = (v12 | (^(v12)))
    for fi15 := 0; fi15 < 19; fi15++ {
        switch (^(fi15)) {
        case 0:
            v16 := GenS1{F0: 255, F1: 726, F2: 13}
            _ = v16
        default:
            var v17 uint = v12
            _ = v17
        }
        fmt.Println(v11)
        _ = fi15
    }
    var v18 string = "a"
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v18
    _ = v2
    _ = v3
}
