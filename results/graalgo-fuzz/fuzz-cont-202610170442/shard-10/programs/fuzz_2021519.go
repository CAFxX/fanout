package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int8; F1 bool }
type GenS3 struct { F0 int32; F1 uint16 }
type GenS4 struct { F0 uint; F1 int16 }
func genFn0(p0 int32) uint64 {
    p0 = p0
    if false {
        v6 := GenS1{F0: "\n\t"}
        _ = v6
    } else {
        fmt.Println(p0)
    }
    _ = p0
    return 42
}
func genFn1(p0 int64) uint64 {
    var v7 uint = 18446744073709551615
    var v8 int32 = -2147483648
    _ = p0
    _ = v7
    _ = v8
    return 324
}

func main() {
    var v1 uint = 1
    var v2 uint = (v1 & 18446744073709551615)
    v3 := []int{-461, -9223372036854775808, 127}
    var v4 uint64 = 358
    var v5 int64 = 42
    if (-708 < 2147483647) {
        var v9 int32 = -1
        v10 := GenS2{F0: 127, F1: true}
        _ = v10
        _ = v9
    }
    var v11 int32 = -93
    var v12 float32 = (1.5 + (0.1 * 0.0))
    var v13 int64 = v5
    switch 999 {
    case 32768:
        if true {
            v14 := GenS3{F0: (^(v11)), F1: 445}
            _ = v14
        }
    default:
        var v15 uint8 = 128
        _ = v15
    }
    v1 -= 2
    v16 := GenS3{F0: -2147483648, F1: 42}
    var v17 float64 = float64((0.1 - v12))
    var v18 int8 = 127
    v18 += (v18 | (v18 / ((v18) | 1)))
    switch 425 {
    case -9223372036854775808:
        v13 = v5
    }
    v19 := GenS3{F0: -1, F1: 242}
    switch ("" + "\n\t") {
    case "\n\t":
        v1 = v1
    case "":
        v1 = v1
    }
    var v20 int16 = 441
    v21 := GenS4{F0: (^(v1)), F1: (v20 - -485)}
    v22 := GenS4{F0: 145, F1: (v20 - v20)}
    var v23 uint32 = 2147483647
    for fi24 := 0; fi24 < 8; fi24++ {
        var v25 int8 = -128
        _ = len(v3)
        _ = fi24
        _ = v25
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
