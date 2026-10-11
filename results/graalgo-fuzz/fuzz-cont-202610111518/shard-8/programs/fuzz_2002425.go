package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int16; F1 string }

func main() {
    var v1 int = -96
    v2 := make(map[int]string)
    v3 := []uint32{4294967295, 320}
    var v4 float32 = (-2.25 + (1.5 + ((0.0 / ((3.14159 - 3.14159) - 0.1)) - 0.0)))
    fmt.Println(v1, v1)
    v4 = v4
    if true {
        if false {
            if len(v3) > 0 {
                v3[0] = 2
            }
            v1 -= v1
        }
        v3 = append(v3, 2147483648)
    }
    var v5 float64 = float64((0.1 * float32(1e10)))
    for fi6 := 0; fi6 < 10; fi6++ {
        fmt.Println(fi6)
        fmt.Println(fi6, v4, v5)
        _ = fi6
    }
    if true {
        v1 += v1
        v7 := GenS1{F0: 4294967295}
        switch v1 {
        case 945:
            switch v1 {
            case 127:
                v1 += (v1 * 128)
            case -1:
                if true {
                    var v8 uint64 = 4294967295
                    _ = v8
                }
            }
        case 629:
            v4 = 0.0
        default:
            for fi9 := 0; fi9 < 20; fi9++ {
                var v10 uint16 = 371
                if false {
                    var v11 uint32 = 4294967295
                    _ = v11
                }
                v12 := GenS1{F0: 0}
                _ = fi9
                _ = v10
                _ = v12
            }
        }
        _ = v7
    } else {
        var v13 uint = 313
        var v14 int8 = -58
        var v15 int64 = 9223372036854775807
        _ = v13
        _ = v14
        _ = v15
    }
    if false {
        var v16 int32 = 0
        _ = v16
    }
    v17 := GenS2{F0: 0, F1: ("" + "foo bar")}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
