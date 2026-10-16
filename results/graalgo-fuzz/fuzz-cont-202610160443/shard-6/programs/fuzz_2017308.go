package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint32; F1 uint8 }
type GenS3 struct { F0 int64; F1 uint32 }
func genFn0(p0 bool, p1 int32, p2 uint8) int8 {
    var v4 float32 = 1.5
    p0 = (518 > 450)
    for fi5 := 0; fi5 < 4; fi5++ {
        v4 -= (v4 - 1.5)
        switch ("" + "foo bar") {
        case "a":
            p2 = (p2 >> ((255) & 63))
        }
        var v6 int16 = -923
        _ = fi5
        _ = v6
    }
    for fi7 := 0; fi7 < 1; fi7++ {
        if (p2 <= (p2 * p2)) {
            for fi8 := 0; fi8 < 4; fi8++ {
                p0 = p0
                var v9 int16 = -652
                fmt.Println(v4)
                p2 -= 128
                _ = fi8
                _ = v9
            }
        } else {
            fmt.Println(v4)
        }
        var v10 int32 = (p1 << ((277) & 63))
        _ = fi7
        _ = v10
    }
    v11 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v4
    return 38
}
func genFn1(p0 int) uint16 {
    for fi12 := 0; fi12 < 17; fi12++ {
        switch ("hello" + "foo bar") {
        case "foo bar":
            if false {
                v13 := []int{}
                _ = v13
            }
        case "hello":
            v14 := GenS1{F0: true}
            _ = v14
        }
        _ = fi12
    }
    v15 := GenS2{F0: 1, F1: 215}
    var v16 bool = false
    fmt.Println(v16)
    for fi17 := 0; fi17 < 9; fi17++ {
        var v18 uint32 = 234
        var v19 uint32 = (v18 >> ((255) & 63))
        _ = fi17
        _ = v18
        _ = v19
    }
    _ = p0
    _ = v15
    _ = v16
    return 1
}

func main() {
    var v1 uint16 = 65535
    v2 := make(map[int]uint32)
    v3 := make(map[string]bool)
    v20 := []int{9223372036854775807, -478}
    v20 = append(v20, 4294967295)
    v21 := GenS2{F0: 65535, F1: 42}
    var v22 bool = true
    if len(v20) > 0 {
        v20[0] = 32768
    }
    fmt.Println(v1)
    v1 -= 0
    v22 = (23 < 0)
    if ((!("foo bar" >= "hello")) || v22) {
        var v23 int8 = 62
        _ = v23
    }
    v22 = v22
    if v22 {
        v1 -= v1
        switch -492 {
        case 539:
            var v24 float32 = (0.1 * 3.14159)
            _ = v24
        }
        var v25 int16 = 42
        _ = v25
    }
    v22 = true
    if (!(754 == 371)) {
        for fi26 := 0; fi26 < 10; fi26++ {
            if (2 > 31) {
                v27 := GenS2{F0: 977, F1: 128}
                _ = v27
            } else {
                if v22 {
                    fmt.Println(fi26)
                    v28 := GenS3{F0: 42, F1: 1}
                    _ = v28
                } else {
                    v29 := v3[("\n\t" + "hello")]
                    v22 = v22
                    var v30 uint32 = 37
                    _ = v29
                    _ = v30
                }
            }
            _ = fi26
        }
        v22 = v22
        if (!(genFn1(-196) == (v1 + 1))) {
            v20 = append(v20, 2147483647)
        } else {
            var v31 float32 = (0.1 / float32(1e10))
            _ = v31
        }
    } else {
        var v32 int8 = 1
        _ = v32
    }
    var v33 float64 = float64(0.0)
    v1 -= v1
    delete(v2, 83)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v33
}
