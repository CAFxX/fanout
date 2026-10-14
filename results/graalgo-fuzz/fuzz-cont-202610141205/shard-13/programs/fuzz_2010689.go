package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint }
type GenS2 struct { F0 int8; F1 int8 }
type GenS3 struct { F0 bool }
func genFn0(p0 uint8, p1 int8) int16 {
    var v5 uint8 = 56
    for fi6 := 0; fi6 < 5; fi6++ {
        v5 += v5
        fmt.Println(p0, fi6)
        var v7 float32 = (1.5 - 0.1)
        _ = fi6
        _ = v7
    }
    v8 := GenS1{F0: -616, F1: 4294967295}
    _ = p0
    _ = p1
    _ = v5
    _ = v8
    return -867
}
func genFn1(p0 uint8, p1 uint8, p2 int) uint64 {
    v9 := GenS1{F0: 4294967295, F1: 1}
    var v10 int64 = 9223372036854775807
    var v11 uint32 = 758
    var v12 int32 = -709
    var v13 float32 = 1e10
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v9
    return 1
}

func main() {
    v1 := make(map[string]int)
    var v2 float64 = 0.1
    v3 := []string{"hello", "hello", ((("foo bar" + "") + "") + "\n\t")}
    v4 := make(map[uint32]uint32)
    var v14 int8 = 18
    v14 += 0
    v15 := GenS1{F0: -1, F1: 668}
    var v16 int16 = -32768
    if (2 > 444) {
        switch 42 {
        case 127:
            if true {
                v17 := GenS2{F0: (v14 + (v14 - v14)), F1: (v14 % ((127) | 1))}
                v2 += (((((v2 * v2) + 0.1) - (0.0 / (v2 + (v2 * v2)))) * v2) / (float64(1e10) * (v2 + v2)))
                _ = v17
            } else {
                v3 = append(v3, "")
                v2 = v2
                var v18 int32 = 42
                _ = v18
            }
        case 811:
            var v19 float64 = 3.14159
            _ = v19
        default:
            fmt.Println(v16)
        }
    }
    if false {
        switch 1 {
        case 2147483647:
            var v20 float64 = v2
            _ = v20
        case 32767:
            v21 := GenS2{F0: 89, F1: v14}
            _ = v21
        }
    } else {
        var v22 uint16 = 255
        _ = v22
    }
    for fi23 := 0; fi23 < 16; fi23++ {
        v24 := GenS1{F0: (fi23 << ((42) & 63)), F1: 1}
        _ = fi23
        _ = v24
    }
    v25 := GenS1{F0: -9223372036854775808, F1: 255}
    v26 := GenS3{F0: true}
    if len(v3) > 0 {
        v3[0] = "\n\t"
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
}
