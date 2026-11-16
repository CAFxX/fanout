package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint32; F1 int8 }
type GenS3 struct { F0 int; F1 int64 }
type GenS4 struct { F0 int; F1 uint16 }
type GenS5 struct { F0 uint16 }

func main() {
    var v1 int16 = 1000
    var v2 int8 = 0
    v3 := make(map[int]bool)
    v4 := GenS1{F0: 243}
    if false {
        v1 += 0
    } else {
        var v5 int64 = -135
        _ = v5
    }
    v6 := GenS2{F0: 224, F1: v2}
    v7 := GenS3{F0: 9223372036854775807, F1: 42}
    v2 -= v2
    v8 := []int{32767, 928}
    v2 = v2
    v2 = v2
    var v9 float64 = 0.0
    fmt.Println(v1)
    v1 += (v1 & 1000)
    var v10 bool = (((71 > 2) && false) || true)
    fmt.Println(v2, v9)
    v2 = (v2 - v2)
    v11 := GenS4{F0: 32767, F1: 255}
    for fi12 := 0; fi12 < 6; fi12++ {
        v3[fi12] = (!(-9223372036854775808 != -1))
        if v10 {
            switch (fi12 % ((fi12) | 1)) {
            case -45:
                var v13 uint8 = 42
                _ = v13
            case 24:
                v14 := GenS1{F0: 0}
                _ = v14
            default:
                var v15 uint8 = 255
                _ = v15
            }
            v10 = (v1 <= 32767)
        } else {
            switch "hello" {
            case "hello":
                v9 = (v9 + v9)
            }
            v2 -= v2
            if ((579 < 65535) || v10) {
                v10 = false
                v16 := v3[fi12]
                v17 := GenS5{F0: 100}
                _ = v16
                _ = v17
            }
        }
        _ = fi12
    }
    v18 := GenS1{F0: -404}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
