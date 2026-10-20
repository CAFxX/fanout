package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int; F1 int16; F2 int16 }
func genFn0(p0 uint32) uint {
    v4 := []int{127, 9223372036854775807, int(p0), -1}
    if len(v4) > 0 {
        v4[0] = 42
    }
    _ = p0
    _ = v4
    return 255
}

func main() {
    var v1 int8 = -1
    var v2 int64 = 1
    var v3 float32 = ((((1e10 + (0.1 - (3.14159 * 1.5))) + 0.1) + ((-2.25 - float32(3.14159)) / (((1e10 / 1.0) - (-2.25 + -2.25)) + (3.14159 - 1.5)))) / 0.1)
    v5 := []int{4294967295, 128, -2147483648}
    v6 := GenS1{F0: 0, F1: 128}
    var v7 int = 406
    v1 = v1
    if (!true) {
        if len(v5) > 0 {
            v5[0] = v7
        }
    } else {
        fmt.Println(v1, v1, v1)
        v8 := GenS1{F0: 2, F1: 206}
        var v9 uint8 = 252
        _ = v8
        _ = v9
    }
    v10 := GenS2{F0: 1}
    v11 := GenS2{F0: 186}
    v2 += v2
    if len(v5) > 0 {
        v5[0] = v7
    }
    if true {
        var v12 uint16 = 1
        _ = v12
    } else {
        var v13 float32 = float32(3.14159)
        _ = v13
    }
    v7 = -2147483648
    var v14 int16 = 414
    fmt.Println(v3, v7)
    v14 -= 546
    v15 := GenS3{F0: -151, F1: 0, F2: v14}
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v5
    _ = v6
    _ = v7
}
