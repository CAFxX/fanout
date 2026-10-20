package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint }
type GenS2 struct { F0 int32; F1 int16; F2 uint16 }
func genFn0() string {
    var v6 float64 = ((0.0 - 1.5) - 0.0)
    v7 := []int{-9223372036854775808, 2147483647, -941, 1}
    _ = v6
    _ = v7
    return "hello"
}

func main() {
    v1 := make(map[uint32]string)
    var v2 uint32 = 4294967295
    var v3 int64 = 863
    var v4 uint16 = uint16(v2)
    var v5 float32 = -2.25
    v8 := GenS1{F0: 0, F1: 855}
    fmt.Println(v4, v4)
    var v9 bool = ((127 < 54) && false)
    var v10 bool = true
    var v11 int = -9223372036854775808
    var v12 bool = ((!(("foo bar" + "\n\t") != "foo bar")) || (v10 && false))
    v10 = v12
    if (836 <= uint64(v2)) {
        if v12 {
            var v13 int64 = 0
            _ = v13
        }
        v11 -= (^(v11))
        v5 += (-2.25 + float32((0.0 - 3.14159)))
    } else {
        v14 := GenS1{F0: 1, F1: 1}
        _ = v14
    }
    for fi15 := 0; fi15 < 9; fi15++ {
        v4 = (v4 | v4)
        switch "hello" {
        case "foo bar":
            var v16 bool = ((128 >= 42) && v10)
            _ = v16
        }
        _ = fi15
    }
    var v17 uint64 = 210
    v18 := []int{(v11 / ((v11) | 1)), v11, 128}
    if (v9 && (float64(v5) == float64(float32(1.5)))) {
        v19 := GenS1{F0: 4294967295, F1: 1}
        var v20 uint = 540
        _ = v19
        _ = v20
    }
    v3 = (v3 & v3)
    switch ("\n\t" + "hello") {
    case "hello":
        if len(v18) >= 2 {
            v18 = v18[1:len(v18)]
        }
    case "\n\t":
        switch "hello" {
        case "hello":
            v21 := GenS2{F0: -110, F1: 1000, F2: v4}
            _ = v21
        case "":
            var v22 uint = 255
            _ = v22
        }
    default:
        var v23 int64 = v3
        _ = v23
    }
    for fi24 := 0; fi24 < 12; fi24++ {
        var v25 uint16 = v4
        v4 += (v25 >> ((uint(v2)) & 63))
        _ = fi24
        _ = v25
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
