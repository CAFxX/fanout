package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint8; F1 string; F2 bool }
type GenS3 struct { F0 uint; F1 uint32 }
func genFn0() uint {
    var v5 float64 = 1.5
    switch "" {
    case "a":
        var v6 uint64 = 4294967295
        _ = v6
    case "hello":
        v7 := []int{}
        _ = v7
    }
    _ = v5
    return 25
}
func genFn1(p0 uint, p1 int32) int8 {
    for fi8 := 0; fi8 < 5; fi8++ {
        v9 := GenS1{F0: false}
        var v10 uint32 = 4294967295
        v11 := []int{}
        _ = fi8
        _ = v10
        _ = v11
        _ = v9
    }
    if (0 < 66) {
        p0 += 485
        var v12 uint32 = 2147483647
        _ = v12
    }
    p0 -= p0
    for fi13 := 0; fi13 < 14; fi13++ {
        v14 := GenS1{F0: true}
        var v15 float32 = float32((0.0 / 3.14159))
        v16 := []int{127}
        _ = fi13
        _ = v14
        _ = v15
        _ = v16
    }
    v17 := []int{42, -362}
    _ = p0
    _ = p1
    _ = v17
    return 100
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := []int{0, -9223372036854775808, 710}
    var v3 uint16 = 65535
    v4 := make(map[string]string)
    v18 := GenS1{F0: true}
    if (!(!(-1 <= -738))) {
        for fi19 := 0; fi19 < 2; fi19++ {
            var v20 int8 = -35
            v3 += (v3 + (v3 | v3))
            fmt.Println(v20, fi19)
            _ = fi19
            _ = v20
        }
    }
    if ((255 >= 2) && false) {
        var v21 int8 = 127
        if len(v2) > 0 {
            v2[0] = 0
        }
        v3 -= v3
        _ = v21
    }
    v22 := GenS2{F0: 26, F1: "", F2: (false && false)}
    v23 := GenS3{F0: 509, F1: 928}
    v24 := GenS3{F0: genFn0(), F1: 0}
    var v25 uint16 = (v3 / ((v3) | 1))
    v3 = (v3 >> ((0) & 63))
    switch (("a" + "a") + "foo bar") {
    case "hello":
        var v26 float32 = (1e10 / 0.1)
        _ = v26
    case "\n\t":
        var v27 uint = 542
        _ = v27
    default:
        v3 = v3
    }
    v25 -= (v3 % ((255) | 1))
    if ((v3 + v3) == (^(v3))) {
        v25 += v25
    }
    if false {
        var v28 uint16 = (v3 % (((v3 | v25)) | 1))
        _ = v28
    }
    for fi29 := 0; fi29 < 8; fi29++ {
        v3 += v3
        var v30 int8 = -128
        _ = fi29
        _ = v30
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v18
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
}
