package main

import (
    "fmt"
)

type GenS1 struct { F0 int }

func main() {
    v1 := []uint32{}
    v2 := make(map[int]int)
    v3 := []int{-9223372036854775808, 264, -381, -2147483648}
    var v4 uint8 = 249
    var v5 int64 = 42
    v4 -= (v4 + 152)
    v6 := v2[385]
    if true {
        v7 := v2[v6]
        _ = v7
    } else {
        v4 = (v4 - 232)
        for fi8 := 0; fi8 < 8; fi8++ {
            var v9 int = 122
            _ = fi8
            _ = v9
        }
        if (!false) {
            v5 = v5
        }
    }
    v10 := GenS1{F0: v6}
    v11 := GenS1{F0: v6}
    v4 += v4
    var v12 float32 = 0.1
    v12 += v12
    switch 4294967295 {
    case 2147483647:
        var v13 int32 = 2147483647
        _ = v13
    case -1:
        if (188 == 42) {
            fmt.Println(v6)
            for fi14 := 0; fi14 < 19; fi14++ {
                delete(v2, 644)
                _ = fi14
            }
            var v15 int32 = -2147483648
            _ = v15
        }
    }
    var v16 int16 = 1
    v17 := GenS1{F0: (v6 & 127)}
    if len(v3) > 0 {
        v3[0] = v6
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
