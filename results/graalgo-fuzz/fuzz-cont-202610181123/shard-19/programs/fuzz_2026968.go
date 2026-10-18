package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int; F2 int64 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint64; F1 string; F2 string }
type GenS5 struct { F0 uint32; F1 uint }
type GenS6 struct { F0 int8 }

func main() {
    v1 := make(map[string]string)
    v2 := make(map[uint32]int)
    var v3 uint32 = 40
    var v4 int16 = -460
    var v5 uint8 = 1
    for fi6 := 0; fi6 < 11; fi6++ {
        v7 := GenS1{F0: (true && false), F1: (-(fi6)), F2: 1}
        for fi8 := 0; fi8 < 17; fi8++ {
            v4 = (-(v4))
            delete(v2, 412)
            v9 := GenS2{F0: (v5 - v5)}
            _ = fi8
            _ = v9
        }
        _ = fi6
        _ = v7
    }
    var v10 int = 964
    v11 := GenS3{F0: -43}
    v3 += 2147483648
    var v12 int8 = 11
    fmt.Println(v10, v4)
    v13 := GenS3{F0: v12}
    fmt.Println(v10)
    fmt.Println(v5, v4, v3)
    if ((true && true) && (3.14159 <= float64(((1e10 + 0.0) - (0.1 * 1e10))))) {
        v5 = (v5 << ((4294967295) & 63))
        if (32767 < int16(v5)) {
            v14 := GenS4{F0: 4294967295, F1: "", F2: "foo bar"}
            v15 := []int{v10, (v10 - -877), (v10 - v10)}
            v16 := GenS2{F0: (v5 & (^(v5)))}
            _ = v14
            _ = v15
            _ = v16
        }
        v17 := GenS5{F0: v3, F1: 65535}
        _ = v17
    }
    v18 := GenS6{F0: v12}
    v19 := []int{42, v10, (v10 + (v10 * (v10 / ((v10) | 1)))), v10}
    _ = len(v19)
    if false {
        fmt.Println(v12, v4, v4)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
