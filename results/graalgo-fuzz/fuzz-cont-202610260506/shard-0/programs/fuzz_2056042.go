package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16 }
type GenS2 struct { F0 int32; F1 uint64; F2 uint }
type GenS3 struct { F0 uint64; F1 int32; F2 int16 }

func main() {
    v1 := make(map[string]int)
    var v2 int64 = 1
    var v3 int64 = 0
    var v4 int64 = 0
    v2 = v4
    if (!false) {
        v5 := []int{}
        if ((933 <= 0) || (false || (2 != 511))) {
            fmt.Println(v3)
            v6 := GenS1{F0: 2147483647, F1: 85}
            _ = v6
        } else {
            var v7 bool = true
            if len(v5) > 0 {
                v5[0] = 4294967295
            }
            fmt.Println(v4)
            _ = v7
        }
        v8 := GenS2{F0: -215, F1: 42, F2: 365}
        _ = v5
        _ = v8
    }
    fmt.Println(v4, v2, v3)
    v2 += v2
    if (!true) {
        v4 += (v2 | v2)
    }
    fmt.Println(v3)
    v2 -= v4
    switch 32767 {
    case 32768:
        if (-350 == (v3 / ((v3) | 1))) {
            v2 += v2
        } else {
            var v9 uint32 = 4294967295
            v10 := []int{0, 662, 42, -250}
            if true {
                v11 := GenS3{F0: 2, F1: 0, F2: -889}
                _ = v11
            } else {
                for fi12 := 0; fi12 < 6; fi12++ {
                    v2 = (v3 << ((18446744073709551615) & 63))
                    _ = fi12
                }
            }
            _ = v10
            _ = v9
        }
    case 4294967295:
        v13 := GenS1{F0: -2147483648, F1: 65535}
        _ = v13
    default:
        v14 := []int{1, 128, -520, 4294967295}
        _ = v14
    }
    switch "" {
    case "hello":
        fmt.Println(v4, v4)
    case "":
        var v15 int = -195
        _ = v15
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
}
