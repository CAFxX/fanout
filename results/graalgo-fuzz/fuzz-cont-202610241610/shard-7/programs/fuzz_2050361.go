package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int; F2 int }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint8; F1 uint8; F2 int64 }
type GenS4 struct { F0 string; F1 uint64; F2 uint8 }
func genFn0(p0 uint8) int32 {
    var v4 int32 = -906
    var v5 string = ""
    var v6 bool = false
    var v7 uint32 = 2147483647
    for fi8 := 0; fi8 < 7; fi8++ {
        v4 += -2147483648
        if (!v6) {
            var v9 uint64 = 42
            var v10 int64 = -9223372036854775808
            var v11 float64 = 0.1
            _ = v10
            _ = v11
            _ = v9
        }
        fmt.Println(v4, v4, fi8)
        _ = fi8
    }
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return 42
}
func genFn1() int {
    for fi12 := 0; fi12 < 13; fi12++ {
        v13 := GenS1{F0: 255, F1: fi12, F2: (fi12 - fi12)}
        _ = fi12
        _ = v13
    }
    var v14 int16 = 32767
    _ = v14
    return 9223372036854775807
}

func main() {
    var v1 int64 = 9223372036854775807
    var v2 int16 = -37
    v3 := make(map[uint32]int)
    fmt.Println(v1)
    if (true && (!true)) {
        fmt.Println(v2, v1)
    } else {
        v15 := GenS1{F0: 65535, F1: -1, F2: -1}
        _ = v15
    }
    v1 = v1
    switch "a" {
    case "hello":
        for fi16 := 0; fi16 < 13; fi16++ {
            v17 := GenS2{F0: (-(v2))}
            fmt.Println(fi16, fi16, fi16)
            _ = fi16
            _ = v17
        }
    default:
        v18 := GenS3{F0: 1, F1: 42, F2: (v1 | v1)}
        _ = v18
    }
    v2 = v2
    v3[4294967295] = -1
    v19 := GenS4{F0: "", F1: 18446744073709551615, F2: 255}
    _ = len(v3)
    v20 := GenS3{F0: 128, F1: 128, F2: v1}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
}
