package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64; F2 uint16 }

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[int]string)
    v3 := make(map[int]string)
    var v4 int64 = 241
    if ((true || (42 >= 131)) && (4294967295 != 771)) {
        v4 = 0
        v1[""] = 42
        v4 += 663
    } else {
        v5 := []int{}
        _ = v5
    }
    if false {
        if (2 == 255) {
            switch 4294967295 {
            case 2147483647:
                v6 := []int{}
                _ = v6
            case -9223372036854775808:
                fmt.Println(v4)
            case 587:
                delete(v2, 4294967295)
            }
            v4 -= v4
        } else {
            for fi7 := 0; fi7 < 11; fi7++ {
                v4 = (^(v4))
                v4 -= (-(v4))
                _ = fi7
            }
            var v8 float64 = (float64((3.14159 - 0.0)) * 3.14159)
            _ = v8
        }
        var v9 uint32 = 1
        var v10 int16 = -29
        _ = v10
        _ = v9
    }
    v11 := GenS1{F0: true, F1: 4294967295, F2: 0}
    for fi12 := 0; fi12 < 5; fi12++ {
        var v13 int = (fi12 + (fi12 / ((-2147483648) | 1)))
        var v14 uint16 = 42
        v15 := []int{}
        _ = fi12
        _ = v13
        _ = v14
        _ = v15
    }
    fmt.Println(v4)
    v4 = v4
    switch 127 {
    case 127:
        for fi16 := 0; fi16 < 7; fi16++ {
            var v17 bool = true
            v4 -= (-(v4))
            var v18 int64 = (v4 * v4)
            _ = fi16
            _ = v17
            _ = v18
        }
    case 730:
        var v19 string = "a"
        _ = v19
    default:
        v4 = 42
    }
    if true {
        var v20 uint8 = 81
        _ = v20
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
}
