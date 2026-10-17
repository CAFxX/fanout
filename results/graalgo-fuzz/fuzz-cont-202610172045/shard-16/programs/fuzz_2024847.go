package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint8; F1 int; F2 bool }
type GenS3 struct { F0 bool; F1 uint32 }
func genFn0(p0 uint, p1 int32, p2 int64) bool {
    v5 := []int{9223372036854775807, 305}
    if len(v5) > 0 {
        v5[0] = -702
    }
    p1 = (p1 | (p1 << (((p0 << ((p0) & 63))) & 63)))
    v6 := GenS1{F0: p0}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return true
}

func main() {
    var v1 float64 = (0.0 * 1.5)
    var v2 int8 = -83
    var v3 float64 = v1
    v4 := make(map[string]string)
    var v7 float32 = 3.14159
    v8 := GenS2{F0: 128, F1: 0, F2: (2 < 0)}
    if (true || (594 < 32767)) {
        v9 := v4["\n\t"]
        var v10 uint32 = 746
        for fi11 := 0; fi11 < 7; fi11++ {
            fmt.Println(fi11, v7, v1)
            if (v7 == (float32(v3) * 0.0)) {
                fmt.Println(v9)
                v3 -= ((v1 / (((1.5 * v1) + 1e10) / ((((1e10 + v1) + 0.1) + (v3 / v3)) - (v1 + v3)))) - (v3 / 1.5))
            }
            _ = fi11
        }
        _ = v10
        _ = v9
    }
    for fi12 := 0; fi12 < 8; fi12++ {
        v13 := []int{(^(fi12)), 32767, 0}
        var v14 int64 = -24
        var v15 uint64 = 0
        _ = fi12
        _ = v13
        _ = v14
        _ = v15
    }
    v7 = v7
    v16 := GenS3{F0: (0 < 273), F1: 0}
    v1 += float64(0.0)
    for fi17 := 0; fi17 < 1; fi17++ {
        switch (("a" + (("\n\t" + "") + ("a" + ("\n\t" + "a")))) + "a") {
        case "foo bar":
            v3 = ((float64(((float32(float64(-2.25)) + 1e10) - v7)) * v1) / float64(1e10))
        default:
            v2 = 100
        }
        _ = fi17
    }
    v7 -= ((((((-2.25 - ((3.14159 - (v7 + v7)) + (v7 / v7))) * float32(v1)) - float32(0.1)) / (v7 * ((float32(v3) * v7) + v7))) + float32(1.5)) + v7)
    v18 := []int{}
    switch 4294967295 {
    case -9223372036854775808:
        v2 += 100
    case 210:
        var v19 uint64 = 18446744073709551615
        _ = v19
    }
    _ = len(v4)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v16
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
