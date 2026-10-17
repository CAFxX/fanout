package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8 }

func main() {
    v1 := make(map[uint32]int)
    var v2 float32 = float32((-2.25 * 1e10))
    var v3 float64 = (0.0 + (0.0 * 1.5))
    v4 := make(map[string]bool)
    var v5 int64 = -99
    v6 := GenS1{F0: -9223372036854775808, F1: 127}
    fmt.Println(v5, v5)
    var v7 bool = false
    var v8 int8 = -1
    for fi9 := 0; fi9 < 18; fi9++ {
        if v7 {
            if ((!((fi9 ^ fi9) > 5)) || (2147483647 == -2147483648)) {
                fmt.Println(fi9, v5, v5)
                if false {
                    v2 -= v2
                    v10 := []int{(fi9 ^ (fi9 - (fi9 << ((0) & 63)))), fi9}
                    _ = v10
                }
            }
            v11 := GenS1{F0: v5, F1: v8}
            _ = v11
        } else {
            if v7 {
                v12 := []int{128, fi9, fi9, 32767}
                var v13 uint16 = 260
                for fi14 := 0; fi14 < 4; fi14++ {
                    v8 += (v8 | 1)
                    v12 = append(v12, 989)
                    _ = fi14
                }
                _ = v12
                _ = v13
            }
            switch 128 {
            case -963:
                v15 := GenS1{F0: -9223372036854775808, F1: 1}
                _ = v15
            case -911:
                v3 += (v3 + 1.5)
            case 123:
                v7 = false
            }
        }
        _ = fi9
    }
    for fi16 := 0; fi16 < 9; fi16++ {
        v17 := []int{(fi16 * (fi16 / ((-358) | 1)))}
        _ = fi16
        _ = v17
    }
    v18 := GenS1{F0: -1, F1: 0}
    fmt.Println(v8, v2)
    fmt.Println(v2, v5, v2)
    v5 += 42
    v19 := []int{}
    switch "foo bar" {
    case "":
        v2 = (v2 - (v2 + v2))
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
