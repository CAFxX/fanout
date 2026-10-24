package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8; F2 uint }
type GenS2 struct { F0 uint8; F1 uint64 }

func main() {
    var v1 int16 = -32768
    v2 := make(map[string]uint32)
    v3 := []int{1, 127}
    v4 := []int{42, 98}
    var v5 float32 = 0.0
    var v6 float32 = ((v5 / -2.25) * v5)
    var v7 uint32 = 587
    v2["a"] = (^(v7))
    var v8 uint16 = 250
    if false {
        delete(v2, ("" + (("" + "a") + "\n\t")))
    }
    for fi9 := 0; fi9 < 20; fi9++ {
        v10 := GenS1{F0: 1, F1: 42, F2: 2}
        var v11 int64 = -9223372036854775808
        _ = fi9
        _ = v10
        _ = v11
    }
    v8 = (v8 - (v8 - 42))
    v6 -= v5
    if true {
        if (42 > 2) {
            fmt.Println(v6)
            if (128 > 128) {
                var v12 bool = ((v7 % (((v7 / (((v7 & uint32(v8))) | 1))) | 1)) == v7)
                _ = v12
            } else {
                v13 := GenS1{F0: 18446744073709551615, F1: 0, F2: 255}
                var v14 uint64 = 4294967295
                _ = v13
                _ = v14
            }
        } else {
            _ = len(v2)
            var v15 int64 = -765
            _ = v15
        }
        fmt.Println(v8, v7, v5)
    }
    switch -250 {
    case 2147483647:
        _ = len(v2)
    case -340:
        var v16 uint16 = v8
        _ = v16
    case 902:
        v5 += (v6 + v6)
    default:
        var v17 uint8 = 23
        _ = v17
    }
    var v18 uint = 204
    var v19 int8 = 42
    switch 0 {
    case -764:
        var v20 bool = false
        _ = v20
    case 32768:
        var v21 int32 = -276
        _ = v21
    }
    for fi22 := 0; fi22 < 6; fi22++ {
        if len(v3) > 0 {
            v3[0] = (fi22 - fi22)
        }
        v23 := GenS2{F0: 128, F1: 0}
        var v24 uint = (v18 + 65535)
        _ = fi22
        _ = v23
        _ = v24
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
