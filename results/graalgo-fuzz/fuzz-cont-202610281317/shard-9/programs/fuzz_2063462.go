package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int; F2 int8 }
type GenS2 struct { F0 int8; F1 uint16; F2 uint8 }
type GenS3 struct { F0 uint32; F1 int8; F2 uint }
func genFn0(p0 uint8) uint32 {
    var v6 uint16 = 90
    v7 := GenS1{F0: "", F1: -916, F2: 1}
    _ = p0
    _ = v6
    _ = v7
    return 2147483647
}

func main() {
    v1 := []int{127, -423, 140, -9223372036854775808}
    v2 := []uint32{763, 2147483648, 192, 2147483647}
    var v3 uint8 = 210
    var v4 uint8 = (v3 & 2)
    var v5 uint16 = 32767
    v8 := GenS2{F0: 42, F1: 255, F2: 2}
    v3 = (v3 & (v3 | 255))
    var v9 int = 32768
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v10 uint8 = 128
    if len(v1) > 0 {
        v1[0] = 2147483647
    }
    var v11 int8 = 93
    v9 -= v9
    if true {
        var v12 string = "foo bar"
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = v12
    }
    fmt.Println(v4)
    if (18446744073709551615 <= 2) {
        fmt.Println(v11, v4, v4)
    } else {
        var v13 uint = 65535
        _ = v13
    }
    var v14 int8 = (v11 + (^(v11)))
    v4 -= v3
    if ((0 <= 32767) || (0 <= 1)) {
        v15 := GenS3{F0: 2147483647, F1: 66, F2: 42}
        _ = v15
    } else {
        var v16 float32 = 0.1
        _ = v16
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
