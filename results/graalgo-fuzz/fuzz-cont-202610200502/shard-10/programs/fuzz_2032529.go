package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8 }
type GenS2 struct { F0 uint16; F1 uint8; F2 int8 }
type GenS3 struct { F0 uint16; F1 int8 }

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[int]int)
    v3 := make(map[string]bool)
    v4 := []int{1, -2, 4294967295}
    v5 := make(map[string]string)
    v6 := make(map[int]bool)
    var v7 float32 = (float32((((float64(float32(3.14159)) - 1.5) + 0.0) * 3.14159)) + float32((-2.25 / -2.25)))
    for fi8 := 0; fi8 < 10; fi8++ {
        for fi9 := 0; fi9 < 11; fi9++ {
            if true {
                v7 -= -2.25
                v7 = float32(1e10)
                fmt.Println(v7, fi9)
            } else {
                v10 := v3["a"]
                v10 = true
                var v11 float32 = 3.14159
                _ = v10
                _ = v11
            }
            v7 = 1.5
            _ = fi9
        }
        fmt.Println(v7, fi8)
        _ = fi8
    }
    var v12 uint64 = 4294967295
    v13 := GenS1{F0: 18446744073709551615, F1: 128}
    switch 32767 {
    case 9223372036854775807:
        v7 += v7
    case 127:
        if true {
            var v14 int32 = -1
            _ = v14
        }
    default:
        if ((v7 > (v7 / v7)) && (false || false)) {
            var v15 int = -9223372036854775808
            _ = v15
        }
    }
    v7 -= ((v7 - 3.14159) - 1.5)
    v16 := GenS2{F0: 663, F1: 235, F2: 100}
    if (-2147483648 > -328) {
        v17 := GenS3{F0: 985, F1: 62}
        _ = v17
    } else {
        v18 := GenS3{F0: 652, F1: 100}
        _ = v18
    }
    if len(v4) > 0 {
        v4[0] = -171
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v13
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
