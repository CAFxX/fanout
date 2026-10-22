package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64; F2 int }
func genFn0(p0 uint8, p1 int) int16 {
    for fi5 := 0; fi5 < 10; fi5++ {
        var v6 int = 2147483647
        _ = fi5
        _ = v6
    }
    for fi7 := 0; fi7 < 1; fi7++ {
        p1 -= (fi7 + (p1 + 207))
        _ = fi7
    }
    p1 -= p1
    fmt.Println(p0, p1)
    _ = p0
    _ = p1
    return 32767
}
func genFn1() uint64 {
    if true {
        var v8 string = "\n\t"
        v9 := []int{-788, 4294967295, 385, -121}
        _ = v8
        _ = v9
    } else {
        var v10 bool = (0 >= 206)
        v11 := []int{}
        _ = v10
        _ = v11
    }
    var v12 int8 = 1
    _ = v12
    return 872
}

func main() {
    var v1 int16 = 32767
    var v2 int8 = -18
    var v3 int64 = 9223372036854775807
    var v4 int16 = v1
    var v13 int8 = (^(v2))
    for fi14 := 0; fi14 < 6; fi14++ {
        v1 = v4
        for fi15 := 0; fi15 < 5; fi15++ {
            fmt.Println(fi14, v4)
            var v16 int16 = (v1 * v1)
            v17 := []int{(fi14 + (fi14 / ((42) | 1)))}
            _ = fi15
            _ = v16
            _ = v17
        }
        _ = fi14
    }
    v18 := []int{-789, 128, 127}
    for fi19 := 0; fi19 < 9; fi19++ {
        fmt.Println(v2, v2, v3)
        _ = fi19
    }
    for fi20 := 0; fi20 < 8; fi20++ {
        var v21 uint32 = 4294967295
        if len(v18) > 0 {
            v18[0] = 128
        }
        _ = fi20
        _ = v21
    }
    v22 := GenS1{F0: "", F1: genFn1(), F2: 1}
    var v23 uint16 = 256
    v1 += v1
    _ = len(v18)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v18
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
}
