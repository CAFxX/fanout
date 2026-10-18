package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint16; F1 int32 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 bool }
func genFn0(p0 int16) uint32 {
    var v4 uint8 = 128
    for fi5 := 0; fi5 < 14; fi5++ {
        p0 = int16(v4)
        v4 = v4
        _ = fi5
    }
    var v6 string = ("a" + "a")
    var v7 uint64 = 42
    _ = p0
    _ = v4
    _ = v6
    _ = v7
    return 836
}
func genFn1(p0 int) bool {
    var v8 int64 = 606
    v8 = v8
    v9 := []int{-9223372036854775808, 32767, -2147483648, 845}
    switch "" {
    case "hello":
        for fi10 := 0; fi10 < 8; fi10++ {
            var v11 int8 = 125
            var v12 int = (-(fi10))
            _ = fi10
            _ = v11
            _ = v12
        }
    case "\n\t":
        var v13 uint32 = genFn0(-222)
        _ = v13
    default:
        v14 := GenS1{F0: 255}
        _ = v14
    }
    switch p0 {
    case 42:
        v15 := GenS2{F0: 0, F1: -176}
        _ = v15
    case 127:
        var v16 int = 32767
        _ = v16
    case 32768:
        if false {
            var v17 int16 = -32768
            if len(v9) >= 2 {
                v9 = v9[1:len(v9)]
            }
            _ = v17
        }
    }
    _ = p0
    _ = v8
    _ = v9
    return (false && (-2.25 == ((1e10 * 1.5) - 0.1)))
}

func main() {
    var v1 int64 = -431
    v2 := make(map[string]int)
    var v3 int = 935
    var v18 int64 = v1
    v19 := GenS3{F0: v1}
    v3 += v3
    v20 := []int{9223372036854775807, 32768, v3}
    var v21 int64 = v1
    if ((-2147483648 != 42) && (128 == 51)) {
        v22 := GenS3{F0: v1}
        var v23 float32 = -2.25
        if true {
            var v24 float64 = 3.14159
            var v25 int = (v3 - v3)
            _ = v24
            _ = v25
        } else {
            if false {
                fmt.Println(v1)
            } else {
                v21 -= v1
                _ = len(v2)
            }
            v26 := GenS4{F0: (true || false)}
            _ = v26
        }
        _ = v22
        _ = v23
    } else {
        var v27 uint = 42
        _ = v27
    }
    if len(v20) >= 2 {
        v20 = v20[1:len(v20)]
    }
    if (795 == 0) {
        v28 := GenS2{F0: 657, F1: 2147483647}
        _ = v28
    } else {
        v3 = v3
        for fi29 := 0; fi29 < 2; fi29++ {
            var v30 int8 = 42
            v21 = int64(v3)
            v21 = (^(v1))
            _ = fi29
            _ = v30
        }
    }
    var v31 uint = 457
    _ = len(v2)
    v21 += v18
    if (((v3 / ((v3) | 1)) == v3) || false) {
        var v32 int16 = 218
        _ = v32
    }
    v20 = append(v20, v3)
    var v33 float32 = 1e10
    var v34 uint32 = 1
    v2[("\n\t" + "hello")] = v3
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v31
    _ = v33
    _ = v34
}
