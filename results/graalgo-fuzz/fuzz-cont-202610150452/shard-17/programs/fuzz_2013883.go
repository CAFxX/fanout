package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16; F2 int }
func genFn0(p0 int8, p1 uint16, p2 int) uint8 {
    var v7 float32 = (float32(0.0) / (0.1 + -2.25))
    for fi8 := 0; fi8 < 9; fi8++ {
        p1 += (p1 - p1)
        _ = fi8
    }
    for fi9 := 0; fi9 < 17; fi9++ {
        var v10 uint32 = 42
        _ = fi9
        _ = v10
    }
    p2 += (p2 & p2)
    v11 := []int{p2, p2, p2}
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v7
    return 186
}

func main() {
    v1 := make(map[string]int)
    var v2 string = ((("a" + (("\n\t" + "hello") + ("" + "foo bar"))) + "\n\t") + "foo bar")
    var v3 float32 = (3.14159 - 0.1)
    var v4 uint8 = 114
    var v5 bool = false
    v6 := make(map[string]bool)
    for fi12 := 0; fi12 < 18; fi12++ {
        v2 = v2
        _ = fi12
    }
    var v13 float64 = ((0.0 + (float64(0.0) + 1e10)) - (1e10 - (0.1 + 1.5)))
    v4 = v4
    v14 := GenS1{F0: 216, F1: 278, F2: 42}
    switch 128 {
    case 591:
        var v15 int16 = 0
        _ = v15
    case 1:
        v4 -= (v4 * (v4 + v4))
    }
    var v16 uint16 = 255
    v4 -= (^(v4))
    v17 := []int{4294967295, 127}
    v16 = v16
    var v18 int16 = 32767
    if v5 {
        v19 := GenS1{F0: v18, F1: (v16 % ((v16) | 1)), F2: 9223372036854775807}
        if (813 != 2147483647) {
            fmt.Println(v5)
        }
        _ = v19
    }
    _ = len(v17)
    if (1 < 923) {
        for fi20 := 0; fi20 < 9; fi20++ {
            var v21 uint16 = v16
            _ = fi20
            _ = v21
        }
        if false {
            var v22 uint8 = genFn0(-128, 0, -476)
            v5 = (v3 == v3)
            _ = v22
        }
    }
    var v23 int = 32768
    var v24 int8 = -128
    if len(v17) >= 2 {
        v17 = v17[1:len(v17)]
    }
    var v25 uint16 = v16
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v13
    _ = v14
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
