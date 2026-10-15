package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 string }

func main() {
    var v1 uint64 = 18446744073709551615
    v2 := make(map[uint32]bool)
    var v3 uint8 = 1
    v4 := []int{0, 42, 4294967295}
    var v5 float64 = ((1e10 * (float64(3.14159) + (-2.25 / (0.1 * (float64(3.14159) / ((3.14159 * 0.1) + (3.14159 * (float64(-2.25) / -2.25)))))))) - (1.5 / float64(1.5)))
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v3 += 246
    v3 += (v3 ^ v3)
    for fi6 := 0; fi6 < 20; fi6++ {
        if (0 > 1000) {
            fmt.Println(v1, fi6, v5)
            v1 = (^(v1))
        } else {
            v1 = v1
        }
        v2[2147483648] = false
        _ = fi6
    }
    v7 := GenS1{F0: 10}
    for fi8 := 0; fi8 < 14; fi8++ {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        if false {
            switch (fi8 << ((0) & 63)) {
            case 2147483647:
                v1 += v1
            case -9223372036854775808:
                v1 += v1
            case 42:
                var v9 float64 = (v5 * v5)
                _ = v9
            default:
                v10 := GenS2{F0: "a"}
                _ = v10
            }
        }
        _ = fi8
    }
    if (!false) {
        switch 32767 {
        case 954:
            v5 = (1.5 / v5)
        }
        var v11 int64 = -851
        fmt.Println(v1, v3)
        _ = v11
    } else {
        v1 = v1
    }
    fmt.Println(v3, v3)
    v4 = append(v4, -9223372036854775808)
    fmt.Println(v3, v5, v5)
    var v12 int = 346
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
