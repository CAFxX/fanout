package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
func genFn0(p0 uint64, p1 int16, p2 uint8) int {
    v5 := []int{}
    if false {
        switch -346 {
        case -534:
            v6 := GenS1{F0: "a"}
            _ = v6
        default:
            for fi7 := 0; fi7 < 13; fi7++ {
                var v8 int32 = 708
                _ = fi7
                _ = v8
            }
        }
        if true {
            p2 += 245
        }
        var v9 uint32 = 799
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 127
}

func main() {
    var v1 float64 = 0.0
    v2 := make(map[int]uint32)
    v3 := make(map[uint32]int)
    v4 := make(map[uint32]bool)
    var v10 uint16 = 1
    v11 := []int{-2147483648, 4294967295, 262, 127}
    for fi12 := 0; fi12 < 10; fi12++ {
        if false {
            var v13 int64 = 9223372036854775807
            v13 = (v13 >> ((255) & 63))
            _ = v13
        }
        fmt.Println(fi12, v10, v10)
        _ = fi12
    }
    v1 = (((1e10 / 1e10) - v1) + ((v1 / ((float64((float32(1e10) + (1.5 + 0.1))) + (v1 / 1e10)) + 0.1)) + v1))
    v10 -= (v10 & (v10 & v10))
    v1 = ((v1 * 0.1) / ((float64((1.5 - 0.1)) - (v1 * 1.5)) - v1))
    var v14 uint16 = 0
    var v15 int = 965
    _ = len(v4)
    fmt.Println(v1)
    var v16 int16 = -445
    if false {
        v17 := v2[v15]
        _ = v17
    }
    switch (^(v15)) {
    case 32768:
        if (1e10 <= 3.14159) {
            var v18 int32 = 1
            if (1 < v16) {
                var v19 int32 = (v18 % ((v18) | 1))
                _ = v19
            }
            _ = v18
        }
    default:
        fmt.Println(v15, v1)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
