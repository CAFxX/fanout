package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8; F2 uint }
func genFn0(p0 int8, p1 uint) int32 {
    switch 42 {
    case -930:
        for fi7 := 0; fi7 < 7; fi7++ {
            var v8 bool = true
            _ = fi7
            _ = v8
        }
    default:
        var v9 uint32 = 1
        _ = v9
    }
    var v10 int8 = (p0 * p0)
    for fi11 := 0; fi11 < 17; fi11++ {
        v10 -= p0
        _ = fi11
    }
    _ = p0
    _ = p1
    _ = v10
    return -2147483648
}

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[string]bool)
    var v3 float64 = 0.0
    v4 := make(map[string]bool)
    v5 := make(map[uint32]string)
    var v6 int64 = -1
    v12 := []int{}
    v6 = (-(v6))
    fmt.Println(v6, v6)
    var v13 int64 = (^(v6))
    v3 = v3
    for fi14 := 0; fi14 < 3; fi14++ {
        v3 += v3
        v6 = -268
        _ = fi14
    }
    var v15 float64 = 0.0
    v6 = (v13 & v13)
    v16 := GenS1{F0: 32767, F1: 2, F2: 684}
    var v17 bool = false
    if (false || (!false)) {
        v15 -= ((((0.0 * ((v15 / (v15 + float64((3.14159 + -2.25)))) + float64((0.1 / 3.14159)))) + v15) * 1e10) / (float64(0.1) / -2.25))
    } else {
        var v18 int = -2147483648
        var v19 int32 = genFn0(-40, 4294967295)
        var v20 int = v18
        _ = v18
        _ = v19
        _ = v20
    }
    v17 = true
    v17 = v17
    v17 = v17
    v21 := v5[405]
    var v22 float32 = 1e10
    v21 = ((v21 + (v21 + v21)) + "a")
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
