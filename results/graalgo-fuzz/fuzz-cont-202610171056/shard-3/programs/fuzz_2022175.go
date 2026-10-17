package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int64; F1 uint }
type GenS4 struct { F0 int }
func genFn0() string {
    var v4 float32 = 3.14159
    v5 := GenS1{F0: 2, F1: 28}
    _ = v4
    _ = v5
    return "\n\t"
}
func genFn1(p0 uint64, p1 uint64, p2 int16) int8 {
    switch int(p1) {
    case 42:
        p0 = (p0 | 568)
    case 32767:
        p1 = (p1 << ((255) & 63))
    }
    var v6 uint32 = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 127
}

func main() {
    v1 := []int{-795, 127, 1}
    v2 := []uint32{1, 177, 42}
    v3 := make(map[string]uint32)
    var v7 int8 = -1
    var v8 string = ("foo bar" + "\n\t")
    v9 := v3[v8]
    _ = len(v3)
    fmt.Println(v8, v8, v7)
    if true {
        _ = len(v3)
        v8 = v8
        v10 := GenS2{F0: 557}
        _ = v10
    }
    v11 := GenS1{F0: 255, F1: 42}
    for fi12 := 0; fi12 < 16; fi12++ {
        if ((-(fi12)) != (fi12 / (((fi12 + fi12)) | 1))) {
            for fi13 := 0; fi13 < 7; fi13++ {
                v9 = (^(v9))
                v1 = append(v1, fi13)
                _ = fi13
            }
        }
        v7 += (v7 << ((313) & 63))
        _ = fi12
    }
    v7 = genFn1(459, 18446744073709551615, 0)
    if false {
        switch 9223372036854775807 {
        case 128:
            v9 = (^(v9))
        case 1:
            if (815 <= 446) {
                v14 := GenS3{F0: 0, F1: 65535}
                var v15 float32 = ((3.14159 * float32(0.0)) - 1e10)
                _ = v14
                _ = v15
            } else {
                v7 += v7
            }
        default:
            if len(v2) > 0 {
                v2[0] = v9
            }
        }
        var v16 uint64 = 0
        var v17 int = 9223372036854775807
        _ = v16
        _ = v17
    } else {
        v8 = "\n\t"
        fmt.Println(v7, v7)
        var v18 uint64 = 744
        _ = v18
    }
    if false {
        v19 := GenS4{F0: 1}
        var v20 float32 = 0.0
        v9 = 65535
        _ = v19
        _ = v20
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v7
    _ = v8
    _ = v9
}
