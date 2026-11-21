package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool }

func main() {
    var v1 int8 = -1
    var v2 int64 = 42
    v3 := make(map[string]uint32)
    var v4 bool = (1000 < -32768)
    var v5 float64 = (1.5 - 1.5)
    var v6 int32 = 0
    v7 := []int{4294967295}
    v2 = 483
    var v8 string = ""
    v4 = false
    _ = len(v7)
    var v9 bool = v4
    var v10 int32 = -72
    v6 = (v10 * (v10 >> ((42) & 63)))
    v5 = float64(-2.25)
    var v11 uint64 = 2
    var v12 string = ""
    v10 = v10
    v13 := GenS1{F0: v2, F1: (1e10 != v5)}
    for fi14 := 0; fi14 < 2; fi14++ {
        if v9 {
            for fi15 := 0; fi15 < 14; fi15++ {
                if len(v7) > 0 {
                    v7[0] = (fi15 / ((fi14) | 1))
                }
                _ = fi15
            }
            if false {
                fmt.Println(v12)
                for fi16 := 0; fi16 < 11; fi16++ {
                    fmt.Println(v6, v11)
                    _ = fi16
                }
                var v17 uint8 = 231
                _ = v17
            }
            v7 = append(v7, (^(fi14)))
        }
        _ = fi14
    }
    var v18 uint64 = 42
    if v4 {
        v18 -= v18
        var v19 uint8 = 1
        _ = v19
    } else {
        v20 := GenS1{F0: -1, F1: v4}
        _ = v20
    }
    switch int(v2) {
    case 365:
        v5 -= v5
    default:
        var v21 bool = v4
        _ = v21
    }
    for fi22 := 0; fi22 < 3; fi22++ {
        v18 = 18446744073709551615
        var v23 int64 = -9223372036854775808
        _ = fi22
        _ = v23
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
