package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint32; F1 int16 }

func main() {
    var v1 uint32 = 106
    var v2 int16 = -425
    var v3 uint64 = 42
    v4 := make(map[string]bool)
    var v5 int8 = 40
    v6 := make(map[uint32]int)
    var v7 string = "\n\t"
    v8 := GenS1{F0: 1}
    v9 := GenS2{F0: 632}
    var v10 bool = true
    fmt.Println(v2)
    v10 = (0.1 <= 0.0)
    switch v7 {
    case "":
        if v10 {
            v1 += v1
        }
    default:
        v11 := GenS3{F0: v7}
        _ = v11
    }
    var v12 int64 = 9223372036854775807
    v13 := []int{-2147483648}
    for fi14 := 0; fi14 < 15; fi14++ {
        v15 := v6[v1]
        fmt.Println(v5, v12)
        _ = fi14
        _ = v15
    }
    if (!(v10 && (65535 != 743))) {
        v3 = v3
        if (0.0 >= 0.0) {
            for fi16 := 0; fi16 < 5; fi16++ {
                if ((-1 <= 859) || true) {
                    var v17 int16 = (v2 % ((379) | 1))
                    _ = v17
                }
                v18 := GenS3{F0: v7}
                fmt.Println(fi16)
                _ = fi16
                _ = v18
            }
            _ = len(v4)
            v2 = v2
        }
        if (-2.25 == (0.0 - -2.25)) {
            if (!(v10 || (128 != 0))) {
                var v19 int16 = -1
                v19 = v19
                _ = v19
            } else {
                var v20 uint = 853
                v1 -= uint32(v20)
                v7 = v7
                _ = v20
            }
        }
    } else {
        var v21 uint8 = 1
        _ = v21
    }
    for fi22 := 0; fi22 < 15; fi22++ {
        if (100 == (v5 - 100)) {
            v23 := GenS4{F0: 996, F1: -32768}
            _ = v23
        } else {
            var v24 int64 = (v12 >> ((18446744073709551615) & 63))
            _ = v24
        }
        _ = fi22
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
