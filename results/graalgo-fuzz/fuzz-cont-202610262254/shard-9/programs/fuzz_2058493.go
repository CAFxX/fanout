package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint16; F1 uint32 }
type GenS3 struct { F0 int8; F1 int16; F2 string }
func genFn0(p0 int16) uint64 {
    p0 = (p0 * p0)
    p0 = (p0 + (p0 | (p0 >> ((255) & 63))))
    fmt.Println(p0)
    v4 := GenS1{F0: 0}
    _ = p0
    _ = v4
    return 404
}
func genFn1(p0 int8, p1 int, p2 int32) uint64 {
    v5 := GenS1{F0: (p2 >> ((42) & 63))}
    p1 = (p1 ^ -423)
    v6 := []int{p1}
    if false {
        var v7 uint32 = 277
        _ = v7
    }
    var v8 uint16 = 65535
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v8
    return 4294967295
}

func main() {
    var v1 bool = true
    v2 := make(map[int]uint32)
    v3 := []uint32{417}
    v9 := GenS1{F0: 2147483647}
    v10 := GenS1{F0: 2147483647}
    v1 = true
    v1 = true
    v1 = true
    v11 := GenS1{F0: -234}
    switch 127 {
    case 9223372036854775807:
        if len(v3) > 0 {
            v3[0] = 1
        }
    default:
        v1 = v1
    }
    v1 = (false && (-2147483648 != 1))
    v1 = ((0 != -56) && v1)
    _ = len(v3)
    for fi12 := 0; fi12 < 5; fi12++ {
        v13 := GenS2{F0: 32767, F1: 2147483648}
        _ = fi12
        _ = v13
    }
    v1 = v1
    delete(v2, -9223372036854775808)
    v14 := GenS3{F0: -128, F1: 950, F2: "foo bar"}
    var v15 int64 = -459
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v9
}
