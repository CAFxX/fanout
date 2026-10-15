package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32 }
type GenS2 struct { F0 uint16; F1 uint }
type GenS3 struct { F0 uint32 }
func genFn0(p0 bool) string {
    p0 = (!(1 == 1))
    p0 = false
    var v6 int8 = -5
    v7 := []int{9223372036854775807, 127, -1}
    fmt.Println(p0)
    _ = p0
    _ = v6
    _ = v7
    return ("foo bar" + "hello")
}

func main() {
    var v1 uint32 = 2147483648
    v2 := make(map[uint32]uint32)
    var v3 float32 = 1e10
    var v4 bool = (((v1 & (v1 - (v1 % ((816) | 1)))) < 2147483647) || ((false || true) || false))
    var v5 float32 = ((-2.25 * v3) - (float32(float64((0.0 - (v3 / v3)))) / 0.1))
    var v8 uint8 = 1
    var v9 uint = 442
    if false {
        var v10 string = "\n\t"
        var v11 int32 = 163
        fmt.Println(v11, v3, v5)
        _ = v10
        _ = v11
    }
    switch "" {
    case "\n\t":
        var v12 uint32 = uint32(v8)
        _ = v12
    default:
        for fi13 := 0; fi13 < 15; fi13++ {
            v14 := GenS1{F0: true, F1: 0}
            if (-2147483648 >= -1) {
                v15 := []int{-9223372036854775808, (^(fi13)), fi13}
                _ = v15
            } else {
                var v16 uint = v9
                fmt.Println(v16, fi13, v1)
                v5 = v5
                _ = v16
            }
            fmt.Println(fi13, v5, v9)
            _ = fi13
            _ = v14
        }
    }
    v17 := GenS2{F0: 0, F1: v9}
    v18 := []int{-428}
    if (!(v4 && ((!(-41 > 992)) || true))) {
        fmt.Println(v1, v9)
    }
    v19 := GenS3{F0: 0}
    if len(v18) >= 2 {
        v18 = v18[1:len(v18)]
    }
    v20 := GenS3{F0: 957}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
