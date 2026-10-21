package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int }
type GenS2 struct { F0 int32 }
func genFn0(p0 uint64) int64 {
    switch 556 {
    case -646:
        if true {
            for fi7 := 0; fi7 < 10; fi7++ {
                for fi8 := 0; fi8 < 2; fi8++ {
                    p0 = p0
                    _ = fi8
                }
                v9 := GenS1{F0: 252, F1: 32767}
                _ = fi7
                _ = v9
            }
            v10 := []int{-435, 0, 2147483647}
            _ = v10
        }
    case 1:
        fmt.Println(p0)
    case 128:
        var v11 int8 = 77
        _ = v11
    }
    p0 -= p0
    fmt.Println(p0)
    _ = p0
    return 0
}

func main() {
    var v1 uint = 876
    var v2 uint16 = 751
    var v3 float32 = 0.0
    var v4 bool = false
    v5 := make(map[uint32]int)
    var v6 int = 626
    var v12 uint16 = v2
    fmt.Println(v6, v3, v12)
    var v13 uint64 = 532
    v1 += 0
    v14 := GenS2{F0: 42}
    var v15 uint = v1
    v16 := GenS2{F0: 793}
    if ((int64(v13) <= 42) && (v4 && (v4 && (v4 && (-2.25 > 1.5))))) {
        var v17 bool = ((0 <= int32(v1)) || v4)
        fmt.Println(v17, v3, v1)
        _ = v17
    }
    switch "hello" {
    case "\n\t":
        if (!(v4 || ((float64((-2.25 * v3)) - (3.14159 + -2.25)) > 3.14159))) {
            if v4 {
                var v18 uint16 = (v12 - 2)
                fmt.Println(v3, v6)
                _ = v18
            }
        }
    case "a":
        var v19 int = (v6 / (((v6 - (v6 + (v6 >> (((v15 >> ((v1) & 63))) & 63))))) | 1))
        _ = v19
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
