package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint; F2 string }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint64; F1 string; F2 int64 }
func genFn0(p0 bool) uint {
    p0 = p0
    fmt.Println(p0)
    p0 = p0
    var v5 int16 = -108
    _ = p0
    _ = v5
    return 4294967295
}

func main() {
    v1 := make(map[string]int)
    v2 := []int{-680, 127, -688, -1}
    v3 := make(map[string]int)
    var v4 int8 = -128
    v6 := GenS1{F0: 370, F1: 11, F2: ("foo bar" + ("" + ("" + "hello")))}
    var v7 uint32 = 2147483647
    v8 := GenS2{F0: 1}
    switch 0 {
    case -264:
        var v9 int16 = 0
        _ = v9
    default:
        var v10 uint8 = 127
        _ = v10
    }
    v11 := v1[(("\n\t" + "a") + "a")]
    switch "foo bar" {
    case "foo bar":
        var v12 int64 = int64(v11)
        _ = v12
    default:
        v4 = (^(v4))
    }
    v4 -= (-(v4))
    if len(v2) > 0 {
        v2[0] = v11
    }
    v13 := GenS1{F0: 32767, F1: 2, F2: "foo bar"}
    v14 := GenS1{F0: 42, F1: 692, F2: "foo bar"}
    if true {
        v7 += 637
        var v15 uint8 = 221
        _ = v15
    }
    if ((0.0 + (float64(((0.0 + 0.1) - 3.14159)) / 1.0)) > (1e10 / 1.0)) {
        v7 = 4294967295
        v16 := GenS3{F0: 42, F1: "\n\t", F2: 9223372036854775807}
        _ = v16
    }
    for fi17 := 0; fi17 < 18; fi17++ {
        _ = len(v2)
        fmt.Println(fi17, v11, v11)
        _ = fi17
    }
    v11 -= v11
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
    _ = v8
}
