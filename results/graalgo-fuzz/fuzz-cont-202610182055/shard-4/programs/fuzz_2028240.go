package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int }

func main() {
    var v1 bool = (42 >= 2)
    var v2 int64 = -122
    var v3 uint8 = 2
    var v4 float64 = 1e10
    var v5 uint32 = uint32(v3)
    var v6 float64 = (1.5 / v4)
    v7 := []int{128}
    for fi8 := 0; fi8 < 17; fi8++ {
        var v9 int32 = 696
        fmt.Println(v6, v6, v4)
        _ = len(v7)
        _ = fi8
        _ = v9
    }
    v1 = false
    v2 -= v2
    var v10 uint = uint(v5)
    v2 -= (v2 + (v2 & (v2 - (v2 * (v2 * (v2 ^ (v2 * (v2 * (v2 << ((300) & 63))))))))))
    if v1 {
        switch "" {
        case "foo bar":
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
        case "a":
            switch 2147483647 {
            case 1:
                if len(v7) >= 2 {
                    v7 = v7[1:len(v7)]
                }
            case 4294967295:
                var v11 uint8 = v3
                _ = v11
            default:
                v12 := GenS1{F0: v2}
                _ = v12
            }
        case "":
            if (v2 != (v2 / (((-(v2))) | 1))) {
                var v13 uint32 = 2147483647
                _ = v13
            }
        default:
            for fi14 := 0; fi14 < 20; fi14++ {
                v4 = (v6 - v6)
                v5 = (v5 * 758)
                v3 = v3
                _ = fi14
            }
        }
        if (2 < 2) {
            var v15 uint = (v10 >> ((v10) & 63))
            v2 = -632
            _ = v15
        }
        var v16 uint = (v10 + (v10 * v10))
        _ = v16
    } else {
        var v17 int64 = -1
        _ = v17
    }
    v18 := GenS2{F0: 725}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v10
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
