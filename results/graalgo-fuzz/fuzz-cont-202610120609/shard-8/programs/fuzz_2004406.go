package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64 }
type GenS2 struct { F0 uint16; F1 int64; F2 uint }

func main() {
    var v1 int = -925
    var v2 int8 = 0
    v3 := make(map[uint32]int)
    delete(v3, 42)
    switch -9223372036854775808 {
    case 128:
        var v4 uint64 = 1
        _ = v4
    case -9223372036854775808:
        switch "\n\t" {
        case "a":
            if (!(!true)) {
                v2 += (v2 >> ((42) & 63))
            } else {
                var v5 float64 = float64((3.14159 * ((3.14159 / 3.14159) / (3.14159 + 1e10))))
                v5 = 1e10
                _ = v5
            }
        case "foo bar":
            for fi6 := 0; fi6 < 3; fi6++ {
                var v7 string = ""
                delete(v3, 2147483648)
                _ = fi6
                _ = v7
            }
        }
    }
    if (!((1 == 4294967295) && true)) {
        for fi8 := 0; fi8 < 16; fi8++ {
            v2 = v2
            v9 := []int{32768, (fi8 | fi8), (v1 % ((-911) | 1))}
            _ = fi8
            _ = v9
        }
    } else {
        v1 = (v1 & 4294967295)
        var v10 int16 = 32767
        if (0 >= 9223372036854775807) {
            var v11 uint8 = 178
            v2 += v2
            var v12 float32 = (-2.25 * (0.1 - ((-2.25 * 0.1) / 1.0)))
            _ = v11
            _ = v12
        } else {
            v13 := []int{(v1 + 9223372036854775807), (v1 ^ (v1 << ((719) & 63))), int(v10), -185}
            _ = v13
        }
        _ = v10
    }
    v14 := GenS1{F0: 1, F1: 9223372036854775807}
    var v15 int64 = 1
    if false {
        v1 -= (v1 + v1)
        var v16 uint8 = 128
        v2 = v2
        _ = v16
    } else {
        var v17 uint64 = 2
        _ = v17
    }
    _ = len(v3)
    v18 := []int{(^(v1)), (v1 << ((659) & 63)), v1}
    fmt.Println(v2, v2)
    v2 = v2
    var v19 int16 = -576
    for fi20 := 0; fi20 < 9; fi20++ {
        var v21 int = fi20
        fmt.Println(v2, v2, v1)
        _ = fi20
        _ = v21
    }
    for fi22 := 0; fi22 < 13; fi22++ {
        v15 += (^(v15))
        v23 := GenS2{F0: 32767, F1: v15, F2: 1}
        v2 = (v2 - v2)
        _ = fi22
        _ = v23
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v14
    _ = v15
    _ = v18
    _ = v19
    _ = v2
    _ = v3
}
