package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16 }
type GenS2 struct { F0 uint8; F1 int64 }
func genFn0() int8 {
    v6 := []int{-804, -998}
    v7 := GenS1{F0: 4294967295, F1: 65535}
    var v8 float32 = float32(-2.25)
    switch -9223372036854775808 {
    case -9223372036854775808:
        v9 := GenS1{F0: 65535, F1: 255}
        _ = v9
    default:
        v8 += (1e10 + float32(((-2.25 + ((1e10 + float64(0.1)) + 1.5)) * float64(((0.0 + v8) + ((v8 * float32(3.14159)) / v8))))))
    }
    _ = v6
    _ = v7
    _ = v8
    return 115
}

func main() {
    var v1 int64 = 9223372036854775807
    var v2 uint = 44
    var v3 uint16 = 669
    v4 := []uint32{1}
    v5 := make(map[string]int)
    var v10 bool = true
    v2 += 359
    v11 := GenS1{F0: 628, F1: v3}
    var v12 float64 = 1.5
    if (!v10) {
        if v10 {
            var v13 int64 = -1
            fmt.Println(v12)
            delete(v5, ("\n\t" + "a"))
            _ = v13
        } else {
            var v14 int16 = 90
            switch "a" {
            case "hello":
                v2 -= (v2 << ((0) & 63))
            case "\n\t":
                var v15 float64 = 3.14159
                _ = v15
            case "a":
                v16 := GenS2{F0: 128, F1: 42}
                _ = v16
            default:
                v10 = false
            }
            _ = v14
        }
    }
    if v10 {
        switch ("" + "a") {
        case "a":
            var v17 string = "\n\t"
            _ = v17
        default:
            v12 = (float64(0.0) * v12)
        }
    } else {
        v2 -= v2
        if false {
            var v18 int64 = v1
            fmt.Println(v3)
            _ = v18
        } else {
            v1 = (v1 | (v1 & (^(v1))))
            switch "a" {
            case "hello":
                var v19 bool = (3.14159 == 0.0)
                _ = v19
            case "\n\t":
                var v20 uint32 = 0
                _ = v20
            case "a":
                var v21 uint8 = 2
                _ = v21
            default:
                v22 := GenS1{F0: 1, F1: 255}
                _ = v22
            }
            var v23 uint = (^(v2))
            _ = v23
        }
    }
    var v24 uint64 = 730
    var v25 bool = v10
    v10 = true
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
