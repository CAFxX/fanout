package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool }
type GenS2 struct { F0 int8; F1 int16 }
type GenS3 struct { F0 uint64 }

func main() {
    var v1 float64 = (float64(1e10) * 3.14159)
    var v2 float64 = v1
    var v3 int16 = -483
    var v4 uint = 436
    v5 := make(map[uint32]uint32)
    v6 := []int{4294967295, -841, 4294967295, 0}
    if false {
        v7 := GenS1{F0: (2147483647 > 2147483647), F1: true}
        var v8 int32 = -1
        v9 := GenS1{F0: false, F1: (false && false)}
        _ = v7
        _ = v8
        _ = v9
    }
    var v10 string = "foo bar"
    var v11 uint64 = 2
    fmt.Println(v11, v11)
    var v12 int = 32767
    v2 = ((float64(-2.25) / v1) * 0.1)
    for fi13 := 0; fi13 < 1; fi13++ {
        fmt.Println(v1, v1)
        for fi14 := 0; fi14 < 19; fi14++ {
            v15 := GenS1{F0: (42 >= 228), F1: false}
            v11 = 42
            switch v10 {
            case "foo bar":
                _ = len(v5)
            default:
                if true {
                    delete(v5, 720)
                } else {
                    v16 := GenS2{F0: 100, F1: v3}
                    _ = v16
                }
            }
            _ = fi14
            _ = v15
        }
        v4 = 2
        _ = fi13
    }
    v6 = append(v6, (v12 | 127))
    fmt.Println(v2)
    fmt.Println(v4, v3)
    for fi17 := 0; fi17 < 3; fi17++ {
        v18 := GenS2{F0: 42, F1: (-(v3))}
        v19 := GenS3{F0: (v11 << (((v4 & (v4 >> (((v4 % (((v4 << (((v4 * 18446744073709551615)) & 63))) | 1))) & 63)))) & 63))}
        delete(v5, uint32(v4))
        _ = fi17
        _ = v18
        _ = v19
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
