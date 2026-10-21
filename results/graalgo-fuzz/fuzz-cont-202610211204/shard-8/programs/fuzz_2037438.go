package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint }
type GenS2 struct { F0 int64; F1 int; F2 uint16 }

func main() {
    v1 := make(map[string]string)
    v2 := make(map[uint32]uint32)
    v3 := make(map[string]uint32)
    var v4 float64 = -2.25
    v5 := make(map[int]bool)
    v4 -= 1e10
    switch 4294967295 {
    case 0:
        var v6 int16 = 560
        _ = v6
    default:
        var v7 uint8 = 128
        _ = v7
    }
    v4 = ((float64(-2.25) * (v4 - float64((0.1 / 1e10)))) * v4)
    if true {
        fmt.Println(v4)
        v4 -= (3.14159 + 1.5)
        v8 := []int{}
        _ = v8
    }
    for fi9 := 0; fi9 < 18; fi9++ {
        var v10 bool = true
        var v11 float64 = v4
        v12 := GenS1{F0: (42 <= 889), F1: 255}
        _ = fi9
        _ = v10
        _ = v11
        _ = v12
    }
    v13 := GenS2{F0: -9223372036854775808, F1: 9223372036854775807, F2: 170}
    fmt.Println(v4)
    v4 += 0.1
    v4 += v4
    var v14 uint64 = 226
    if false {
        var v15 float32 = 0.1
        if true {
            fmt.Println(v14)
            var v16 int16 = 1
            _ = v16
        }
        fmt.Println(v4)
        _ = v15
    }
    fmt.Println(v14, v14)
    var v17 int64 = 9223372036854775807
    v17 -= v17
    fmt.Println(v14, v17, v4)
    delete(v5, -733)
    v17 -= -392
    for fi18 := 0; fi18 < 6; fi18++ {
        var v19 uint64 = v14
        fmt.Println(v19, v17)
        v20 := GenS1{F0: (!((fi18 | 4294967295) != fi18)), F1: 192}
        _ = fi18
        _ = v19
        _ = v20
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v13
    _ = v14
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
