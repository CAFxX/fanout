package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string }
type GenS2 struct { F0 int16; F1 uint }

func main() {
    var v1 uint32 = 4294967295
    var v2 string = "hello"
    var v3 float64 = float64(float32(-2.25))
    v4 := make(map[int]bool)
    var v5 int16 = 1000
    v6 := make(map[string]bool)
    for fi7 := 0; fi7 < 9; fi7++ {
        var v8 uint64 = 42
        for fi9 := 0; fi9 < 4; fi9++ {
            v1 -= 2147483647
            v1 = 4294967295
            _ = fi9
        }
        v3 = v3
        _ = fi7
        _ = v8
    }
    v10 := GenS1{F0: false, F1: v2}
    var v11 int8 = 20
    var v12 uint16 = 42
    if false {
        var v13 uint = 439
        v11 = 16
        v14 := GenS2{F0: 328, F1: v13}
        _ = v13
        _ = v14
    } else {
        var v15 uint16 = (v12 | (v12 << ((2) & 63)))
        var v16 uint8 = 42
        v17 := []int{}
        _ = v15
        _ = v16
        _ = v17
    }
    for fi18 := 0; fi18 < 18; fi18++ {
        v1 += 978
        v3 = 0.1
        fmt.Println(v11, v3, fi18)
        _ = fi18
    }
    switch -9223372036854775808 {
    case 42:
        for fi19 := 0; fi19 < 3; fi19++ {
            var v20 int16 = (v5 & (v5 | v5))
            _ = fi19
            _ = v20
        }
    default:
        v21 := GenS1{F0: true, F1: ((v2 + v2) + v2)}
        _ = v21
    }
    var v22 int = 2147483647
    for fi23 := 0; fi23 < 7; fi23++ {
        switch -9223372036854775808 {
        case 9223372036854775807:
            fmt.Println(v2, v2, v12)
        case -1:
            var v24 string = "a"
            _ = v24
        }
        var v25 uint8 = 1
        _ = fi23
        _ = v25
    }
    var v26 int = v22
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v22
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
