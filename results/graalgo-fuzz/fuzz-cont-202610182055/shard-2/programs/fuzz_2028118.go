package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool }
type GenS2 struct { F0 uint; F1 int64 }
func genFn0(p0 int, p1 string, p2 bool) uint8 {
    var v4 int16 = -1
    v5 := GenS1{F0: 61, F1: (371 >= 1)}
    p2 = (127 <= -128)
    var v6 uint8 = 127
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    return 2
}
func genFn1() int {
    var v7 int = 9223372036854775807
    fmt.Println(v7)
    var v8 float64 = (1e10 - 0.1)
    _ = v7
    _ = v8
    return v7
}

func main() {
    var v1 uint64 = 42
    v2 := make(map[uint32]bool)
    var v3 bool = true
    if v3 {
        if true {
            var v9 int16 = 624
            _ = v9
        } else {
            fmt.Println(v3, v1)
        }
    } else {
        v1 -= v1
    }
    v10 := []int{296}
    switch 9223372036854775807 {
    case 42:
        var v11 int64 = 9223372036854775807
        _ = v11
    case 689:
        fmt.Println(v3, v3)
    case -1:
        switch "\n\t" {
        case "":
            var v12 uint32 = 2147483647
            _ = v12
        case "a":
            var v13 float64 = (1e10 + 0.1)
            _ = v13
        default:
            v14 := GenS1{F0: -1, F1: (v3 && v3)}
            _ = v14
        }
    default:
        v3 = v3
    }
    fmt.Println(v1, v3)
    v10 = append(v10, 1)
    if (-84 > 127) {
        v15 := GenS1{F0: -128, F1: true}
        v16 := GenS1{F0: 1, F1: v3}
        var v17 uint16 = 65535
        _ = v15
        _ = v16
        _ = v17
    } else {
        _ = len(v2)
    }
    fmt.Println(v1)
    var v18 string = "\n\t"
    var v19 float32 = 0.0
    var v20 int = -2147483648
    _ = len(v10)
    for fi21 := 0; fi21 < 12; fi21++ {
        fmt.Println(v19, fi21, v20)
        v22 := GenS2{F0: 18446744073709551615, F1: 0}
        _ = fi21
        _ = v22
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
}
