package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32 }

func main() {
    var v1 float32 = -2.25
    v2 := make(map[int]int)
    v3 := []int{}
    var v4 uint64 = 18446744073709551615
    v4 += (v4 >> ((255) & 63))
    var v5 int64 = 0
    var v6 int64 = (^(v5))
    _ = len(v3)
    for fi7 := 0; fi7 < 13; fi7++ {
        _ = len(v3)
        v5 += 1
        switch "hello" {
        case "\n\t":
            for fi8 := 0; fi8 < 6; fi8++ {
                if len(v3) > 0 {
                    v3[0] = 42
                }
                fmt.Println(fi7)
                _ = fi8
            }
        default:
            if false {
                delete(v2, 4294967295)
                var v9 int32 = -573
                _ = v9
            }
        }
        _ = fi7
    }
    v10 := GenS1{F0: 4294967295, F1: -2147483648}
    var v11 int32 = 2147483647
    var v12 bool = true
    var v13 int64 = (-(v5))
    for fi14 := 0; fi14 < 8; fi14++ {
        v2[fi14] = fi14
        _ = fi14
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
