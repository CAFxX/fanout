package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64 }
type GenS2 struct { F0 bool; F1 int16 }
func genFn0() uint8 {
    if (false || ((127 <= -62) || false)) {
        v7 := GenS1{F0: 42, F1: 4294967295}
        _ = v7
    } else {
        v8 := []int{0, 127}
        _ = v8
    }
    var v9 uint32 = 2
    _ = v9
    return 0
}

func main() {
    var v1 uint8 = 167
    v2 := make(map[string]int)
    var v3 bool = true
    var v4 int = 119
    var v5 int16 = 42
    var v6 float32 = 1.5
    v1 += (v1 / ((v1) | 1))
    var v10 int = (v4 % (((^(v4))) | 1))
    if (v6 <= ((v6 * float32(1e10)) - v6)) {
        for fi11 := 0; fi11 < 2; fi11++ {
            var v12 uint8 = (v1 / (((v1 / ((v1) | 1))) | 1))
            _ = fi11
            _ = v12
        }
        fmt.Println(v1)
        var v13 float64 = (1.5 / float64((-2.25 * (0.1 * v6))))
        _ = v13
    }
    _ = len(v2)
    if (134 <= 0) {
        var v14 int = v10
        var v15 int64 = 0
        var v16 float32 = 1.5
        _ = v14
        _ = v15
        _ = v16
    } else {
        if ((((v3 || (((0.1 * (0.1 / (-2.25 - (-2.25 - (3.14159 - 1e10))))) + (3.14159 * 1e10)) <= (float64(v6) / float64(((((v6 - v6) + v6) - (0.0 / v6)) - (v6 / (1e10 + v6))))))) && ((^(v1)) > v1)) && true) && v3) {
            var v17 uint8 = (v1 % ((v1) | 1))
            delete(v2, "")
            _ = v17
        }
        v18 := GenS1{F0: 1, F1: 1}
        _ = v18
    }
    fmt.Println(v4, v5)
    for fi19 := 0; fi19 < 6; fi19++ {
        fmt.Println(v1, v10, v3)
        _ = fi19
    }
    v20 := GenS2{F0: true, F1: (v5 % ((v5) | 1))}
    if true {
        v21 := GenS1{F0: 128, F1: 1}
        _ = v21
    } else {
        v22 := []int{-1}
        var v23 uint64 = 4294967295
        _ = v22
        _ = v23
    }
    var v24 float64 = float64(v6)
    var v25 int32 = 992
    if false {
        v10 += v4
    } else {
        v3 = ((-(v5)) == v5)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v20
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
