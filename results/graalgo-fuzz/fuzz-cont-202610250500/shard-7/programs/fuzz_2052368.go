package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint64; F1 int32 }
func genFn0() uint8 {
    if true {
        v4 := GenS1{F0: 42}
        v5 := []int{9223372036854775807, 32768, 127}
        if (0 == 667) {
            var v6 string = "\n\t"
            _ = v6
        }
        _ = v4
        _ = v5
    } else {
        for fi7 := 0; fi7 < 12; fi7++ {
            var v8 uint32 = 471
            v8 = (v8 | v8)
            if (int32(fi7) >= 67) {
                fmt.Println(fi7, fi7)
                v8 = v8
            }
            _ = fi7
            _ = v8
        }
    }
    switch 32767 {
    case 32767:
        v9 := []int{-1, -2147483648, 42}
        _ = v9
    case -441:
        var v10 float64 = (1e10 + (float64(3.14159) / float64(float32(((1.5 * float64(3.14159)) * (1e10 + 0.0))))))
        _ = v10
    default:
        v11 := GenS1{F0: 0}
        _ = v11
    }
    var v12 int64 = 926
    if (25 != 51) {
        var v13 uint32 = 2147483647
        _ = v13
    }
    v14 := GenS2{F0: -1}
    _ = v12
    _ = v14
    return 127
}

func main() {
    v1 := []uint32{2, 541, 2147483647}
    var v2 uint8 = 255
    var v3 uint32 = 2
    fmt.Println(v2, v3)
    var v15 uint = 4294967295
    var v16 uint64 = 922
    fmt.Println(v3, v16)
    if true {
        fmt.Println(v2, v16)
        v3 = 1
        for fi17 := 0; fi17 < 11; fi17++ {
            switch "\n\t" {
            case "a":
                v1 = append(v1, v3)
            }
            for fi18 := 0; fi18 < 4; fi18++ {
                fmt.Println(v15, v16)
                fmt.Println(fi18, v16, fi17)
                _ = fi18
            }
            _ = fi17
        }
    } else {
        var v19 uint = 4294967295
        var v20 uint32 = 900
        if true {
            fmt.Println(v20, v2, v16)
        }
        _ = v19
        _ = v20
    }
    v21 := GenS3{F0: 396, F1: int32(v15)}
    v16 = v16
    var v22 uint = (v15 % ((327) | 1))
    var v23 uint8 = 0
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
}
