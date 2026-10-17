package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool }
type GenS2 struct { F0 bool; F1 uint8 }
func genFn0() uint8 {
    var v5 uint8 = 62
    v6 := []int{112, 4294967295, -842}
    for fi7 := 0; fi7 < 12; fi7++ {
        v5 = (^(v5))
        _ = fi7
    }
    for fi8 := 0; fi8 < 15; fi8++ {
        v5 = v5
        _ = fi8
    }
    _ = v5
    _ = v6
    return (v5 ^ (v5 | 255))
}

func main() {
    v1 := make(map[int]int)
    v2 := []uint32{4294967295}
    v3 := make(map[string]bool)
    var v4 int32 = -2147483648
    var v9 int16 = 42
    var v10 float32 = 3.14159
    fmt.Println(v9, v10, v10)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    for fi11 := 0; fi11 < 1; fi11++ {
        switch "" {
        case "a":
            fmt.Println(fi11, v9, v9)
        case "":
            v4 = -1
        default:
            v4 += v4
        }
        var v12 float64 = 1e10
        delete(v3, "hello")
        _ = fi11
        _ = v12
    }
    v9 = 1000
    v13 := GenS1{F0: 9223372036854775807, F1: (202 <= 2)}
    v14 := GenS2{F0: (32767 == 1), F1: 2}
    var v15 bool = (false && (!false))
    var v16 int32 = -208
    fmt.Println(v4, v10)
    v16 -= v16
    v15 = (!((false || false) || v15))
    var v17 int = 861
    fmt.Println(v16)
    v15 = true
    delete(v1, 127)
    fmt.Println(v9, v16, v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
