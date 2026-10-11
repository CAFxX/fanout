package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool }
func genFn0(p0 uint32) uint {
    p0 -= (p0 ^ p0)
    if (1 > 112) {
        v7 := []int{}
        _ = len(v7)
        _ = v7
    } else {
        p0 -= 175
    }
    var v8 uint64 = 42
    _ = p0
    _ = v8
    return 153
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int16 = 1000
    var v3 uint32 = 159
    v4 := []uint32{(v3 & (v3 - (^(v3)))), 65535}
    var v5 int32 = 1
    v6 := make(map[uint32]uint32)
    for fi9 := 0; fi9 < 6; fi9++ {
        var v10 float32 = 0.1
        _ = fi9
        _ = v10
    }
    switch 676 {
    case 252:
        v3 = 4294967295
    }
    var v11 int8 = 48
    _ = len(v4)
    fmt.Println(v3, v2)
    if ((255 > 123) && false) {
        for fi12 := 0; fi12 < 7; fi12++ {
            var v13 bool = false
            for fi14 := 0; fi14 < 15; fi14++ {
                var v15 int32 = (^(v5))
                _ = fi14
                _ = v15
            }
            _ = fi12
            _ = v13
        }
    } else {
        fmt.Println(v3, v5)
        v5 = v5
    }
    fmt.Println(v11)
    var v16 bool = true
    var v17 int64 = 0
    fmt.Println(v17)
    fmt.Println(v11, v16)
    v5 = v5
    switch "a" {
    case "\n\t":
        for fi18 := 0; fi18 < 6; fi18++ {
            fmt.Println(v2, v17)
            switch -2147483648 {
            case 42:
                v11 -= v11
            case 127:
                if v16 {
                    var v19 uint32 = (^(v3))
                    v4 = append(v4, v19)
                    _ = len(v1)
                    _ = v19
                } else {
                    fmt.Println(fi18, v2, fi18)
                    v20 := GenS1{F0: 260, F1: (((^(v2)) != 1) && false)}
                    _ = v20
                }
            }
            _ = fi18
        }
    default:
        var v21 int32 = (v5 * v5)
        _ = v21
    }
    v3 = (v3 - (v3 >> ((255) & 63)))
    v2 -= -565
    v6[v3] = 2147483647
    v4 = append(v4, (^(v3)))
    fmt.Println(v3, v16)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v11
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
