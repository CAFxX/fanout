package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32 }
type GenS2 struct { F0 uint32; F1 bool; F2 uint64 }
func genFn0() string {
    if (708 > 0) {
        v6 := GenS1{F0: 2147483647, F1: 0}
        _ = v6
    } else {
        var v7 uint = 0
        _ = v7
    }
    var v8 uint = 291
    var v9 uint32 = 0
    v10 := GenS1{F0: -2147483648, F1: -502}
    v11 := GenS1{F0: -2147483648, F1: 0}
    _ = v10
    _ = v11
    _ = v8
    _ = v9
    return "hello"
}
func genFn1(p0 int8, p1 int8, p2 int) string {
    for fi12 := 0; fi12 < 18; fi12++ {
        for fi13 := 0; fi13 < 17; fi13++ {
            var v14 float64 = 3.14159
            p1 = p0
            _ = fi13
            _ = v14
        }
        v15 := GenS1{F0: 614, F1: 0}
        _ = fi12
        _ = v15
    }
    v16 := GenS2{F0: 2, F1: (!true), F2: 1}
    v17 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    _ = v17
    return ""
}

func main() {
    v1 := make(map[int]uint32)
    v2 := []string{"", "foo bar", "\n\t"}
    var v3 bool = false
    var v4 uint32 = 598
    v5 := make(map[int]string)
    var v18 uint32 = (^(v4))
    var v19 float32 = ((((1e10 * ((0.0 * -2.25) + (3.14159 / -2.25))) - 1.5) / ((3.14159 / (3.14159 / 1.5)) - float32(0.0))) + ((((1e10 * -2.25) * 0.0) / 1e10) + 0.0))
    fmt.Println(v18, v18, v19)
    v3 = (-2147483648 >= 4294967295)
    if false {
        for fi20 := 0; fi20 < 2; fi20++ {
            switch "hello" {
            case "foo bar":
                if len(v2) > 0 {
                    v2[0] = "\n\t"
                }
            case "":
                var v21 uint64 = 4294967295
                _ = v21
            }
            v4 -= v4
            v18 -= v4
            _ = fi20
        }
        var v22 uint64 = 1
        if (1 != 182) {
            var v23 string = "foo bar"
            fmt.Println(v4, v19)
            _ = v23
        }
        _ = v22
    } else {
        var v24 uint8 = 2
        _ = v24
    }
    if (18446744073709551615 >= 18446744073709551615) {
        var v25 bool = v3
        var v26 uint32 = v18
        _ = v25
        _ = v26
    }
    var v27 int16 = 1000
    v18 -= (v18 ^ v4)
    v3 = ((!(v19 < v19)) && (1 > -645))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
