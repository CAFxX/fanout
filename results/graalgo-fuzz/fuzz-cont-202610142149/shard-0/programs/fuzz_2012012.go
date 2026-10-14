package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32 }
type GenS2 struct { F0 int; F1 uint }
func genFn0() uint64 {
    v4 := []int{959, 1, 127, -9223372036854775808}
    var v5 float64 = 0.0
    _ = v4
    _ = v5
    return 0
}
func genFn1() int32 {
    if (!true) {
        for fi6 := 0; fi6 < 15; fi6++ {
            switch "\n\t" {
            case "foo bar":
                var v7 int = fi6
                _ = v7
            case "\n\t":
                var v8 int16 = 42
                _ = v8
            }
            var v9 uint32 = 975
            _ = fi6
            _ = v9
        }
    }
    v10 := GenS1{F0: 65535, F1: 0}
    var v11 int8 = -1
    _ = v10
    _ = v11
    return 2147483647
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := []uint32{712}
    var v3 int = 2147483647
    v12 := GenS2{F0: -2147483648, F1: 2}
    var v13 int32 = -1
    if true {
        var v14 int = 679
        fmt.Println(v13, v3, v14)
        _ = v14
    }
    for fi15 := 0; fi15 < 3; fi15++ {
        v2 = append(v2, 999)
        v16 := v1[2147483648]
        _ = fi15
        _ = v16
    }
    v13 -= (v13 & v13)
    v13 += 267
    var v17 bool = false
    v17 = (0 <= 527)
    var v18 uint64 = 42
    v13 = v13
    v19 := GenS1{F0: 255, F1: 2147483647}
    v20 := GenS2{F0: v3, F1: 42}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v13
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
}
