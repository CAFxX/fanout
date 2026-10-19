package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32; F2 bool }
type GenS2 struct { F0 uint; F1 int8; F2 int32 }

func main() {
    var v1 float64 = (-2.25 - 3.14159)
    v2 := make(map[string]uint32)
    var v3 int8 = 100
    v4 := make(map[int]bool)
    if (!false) {
        v1 = (1e10 * (3.14159 / float64((3.14159 / 1.5))))
        v3 = (v3 << ((18446744073709551615) & 63))
        fmt.Println(v3, v3)
    }
    var v5 int32 = -1
    if true {
        for fi6 := 0; fi6 < 20; fi6++ {
            if true {
                v5 += (v5 & v5)
            } else {
                fmt.Println(v5, v1)
            }
            for fi7 := 0; fi7 < 7; fi7++ {
                v8 := GenS1{F0: v5, F1: 65535, F2: true}
                if (131 > 255) {
                    v9 := []int{0}
                    _ = v9
                }
                _ = fi7
                _ = v8
            }
            fmt.Println(v3)
            _ = fi6
        }
    }
    v10 := GenS2{F0: 786, F1: v3, F2: v5}
    var v11 int = 32768
    fmt.Println(v11)
    var v12 uint64 = 215
    switch (v11 + v11) {
    case 2147483647:
        if false {
            v13 := []int{v11, 32767, (v11 + v11)}
            _ = v13
        }
    default:
        v14 := GenS2{F0: 2, F1: v3, F2: v5}
        _ = v14
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
