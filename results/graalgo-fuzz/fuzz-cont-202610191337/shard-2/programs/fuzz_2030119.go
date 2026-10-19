package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16; F2 int64 }
func genFn0() uint32 {
    var v7 int32 = -424
    for fi8 := 0; fi8 < 18; fi8++ {
        if false {
            fmt.Println(v7, v7)
        } else {
            switch "" {
            case "\n\t":
                var v9 int8 = -128
                _ = v9
            case "foo bar":
                var v10 int64 = 42
                _ = v10
            case "a":
                var v11 float64 = (((0.0 * float64((1.5 * (1e10 - 1e10)))) / 1.0) + 3.14159)
                _ = v11
            default:
                var v12 int8 = 42
                _ = v12
            }
            v7 = (v7 >> ((42) & 63))
        }
        _ = fi8
    }
    v13 := []int{}
    _ = v13
    _ = v7
    return 2
}

func main() {
    v1 := []uint32{312, 0, 2147483648, 714}
    v2 := make(map[int]uint32)
    var v3 int8 = 1
    v4 := []uint32{}
    var v5 bool = true
    var v6 int64 = 9223372036854775807
    v3 = v3
    v2[-1] = 282
    for fi14 := 0; fi14 < 18; fi14++ {
        switch "hello" {
        case "hello":
            for fi15 := 0; fi15 < 18; fi15++ {
                var v16 float32 = (1.5 + -2.25)
                _ = fi15
                _ = v16
            }
        case "\n\t":
            fmt.Println(v6)
        }
        _ = fi14
    }
    if true {
        var v17 bool = v5
        _ = v17
    } else {
        v5 = (2147483648 == 530)
    }
    v5 = false
    var v18 int = 9223372036854775807
    v19 := GenS1{F0: 1, F1: 0, F2: -1}
    fmt.Println(v18, v5)
    fmt.Println(v5, v18, v3)
    var v20 int64 = v6
    v21 := GenS1{F0: 255, F1: -906, F2: v20}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
