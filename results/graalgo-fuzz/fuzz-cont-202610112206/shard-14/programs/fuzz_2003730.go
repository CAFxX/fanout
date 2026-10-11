package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8; F2 uint8 }
type GenS2 struct { F0 int32; F1 int }
type GenS3 struct { F0 int; F1 string; F2 string }
type GenS4 struct { F0 uint; F1 uint32; F2 uint64 }

func main() {
    v1 := []int{-610}
    var v2 uint32 = 2147483647
    v3 := []int{-809, 9223372036854775807, int(v2)}
    if true {
        if (0.1 < 0.0) {
            if (!(1 == 1)) {
                var v4 uint64 = 18446744073709551615
                v5 := GenS1{F0: (^(v2)), F1: 127, F2: 2}
                _ = v4
                _ = v5
            }
            v2 -= 29
            v2 -= (v2 % ((233) | 1))
        } else {
            v2 -= 0
            v2 -= 449
            v2 -= (v2 ^ 289)
        }
    }
    if true {
        v6 := GenS1{F0: (v2 | 4294967295), F1: 127, F2: 0}
        switch 127 {
        case -38:
            fmt.Println(v2)
        }
        _ = v6
    }
    v2 = v2
    if false {
        v2 -= v2
    } else {
        v2 += (^(v2))
    }
    v7 := GenS2{F0: -2147483648, F1: -9223372036854775808}
    fmt.Println(v2)
    if len(v1) > 0 {
        v1[0] = 32767
    }
    for fi8 := 0; fi8 < 11; fi8++ {
        for fi9 := 0; fi9 < 20; fi9++ {
            v2 = v2
            fmt.Println(v2, fi8)
            v2 = 2147483648
            _ = fi9
        }
        _ = fi8
    }
    if true {
        if (false && (4294967295 != 42)) {
            fmt.Println(v2)
            if (0.0 < 1.5) {
                var v10 uint16 = 2
                v11 := GenS3{F0: 9223372036854775807, F1: "\n\t", F2: "hello"}
                _ = v10
                _ = v11
            } else {
                v12 := GenS4{F0: 18446744073709551615, F1: v2, F2: 831}
                _ = v12
            }
        }
    } else {
        var v13 int16 = -1
        var v14 float32 = float32(1e10)
        _ = v13
        _ = v14
    }
    v2 = (v2 >> ((42) & 63))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v2
    _ = v3
    _ = v7
}
