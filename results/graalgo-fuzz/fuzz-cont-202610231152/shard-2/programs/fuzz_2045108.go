package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint; F2 int64 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint; F1 uint }

func main() {
    var v1 int32 = 1
    var v2 float32 = (3.14159 / 0.1)
    var v3 string = (("a" + "foo bar") + "foo bar")
    var v4 float64 = ((1e10 / 0.1) - (((-2.25 + -2.25) + ((float64((v2 + v2)) / -2.25) * -2.25)) * (float64((3.14159 / (float32((((1e10 + 1e10) - (1e10 * 3.14159)) / 0.1)) * 3.14159))) / 1.5)))
    v5 := make(map[int]bool)
    v6 := make(map[string]int)
    var v7 uint16 = 42
    v4 = 3.14159
    v8 := []int{}
    for fi9 := 0; fi9 < 1; fi9++ {
        v10 := GenS1{F0: 1, F1: 18446744073709551615, F2: -9223372036854775808}
        var v11 float32 = v2
        var v12 float32 = 1.5
        _ = fi9
        _ = v10
        _ = v11
        _ = v12
    }
    var v13 int32 = (v1 - v1)
    var v14 int8 = 0
    v7 = 32767
    v15 := GenS1{F0: v14, F1: 2, F2: -819}
    for fi16 := 0; fi16 < 14; fi16++ {
        delete(v5, 9223372036854775807)
        switch -2147483648 {
        case -248:
            if false {
                var v17 uint64 = 407
                v3 = v3
                _ = v17
            } else {
                var v18 uint32 = 845
                _ = v18
            }
        }
        _ = fi16
    }
    v1 -= -1
    v19 := GenS2{F0: v7}
    v20 := GenS3{F0: 65535, F1: 4294967295}
    if true {
        v13 = (v13 * v13)
        v13 = 1
        v21 := GenS2{F0: 42}
        _ = v21
    }
    var v22 uint64 = 984
    for fi23 := 0; fi23 < 2; fi23++ {
        v1 = (-(v1))
        if (9223372036854775807 < -1) {
            fmt.Println(v7)
        }
        _ = fi23
    }
    if false {
        fmt.Println(v3, v3)
        var v24 int64 = -9223372036854775808
        if true {
            fmt.Println(v3, v3, v22)
        }
        _ = v24
    }
    switch "hello" {
    case "foo bar":
        var v25 uint16 = (v7 % ((v7) | 1))
        _ = v25
    default:
        for fi26 := 0; fi26 < 17; fi26++ {
            switch "" {
            case "hello":
                _ = len(v5)
            default:
                v14 += (v14 ^ v14)
            }
            if false {
                fmt.Println(v14)
            }
            _ = fi26
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v19
    _ = v2
    _ = v20
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
