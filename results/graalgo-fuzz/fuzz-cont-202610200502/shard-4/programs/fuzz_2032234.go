package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool; F2 uint32 }
type GenS2 struct { F0 bool; F1 uint8; F2 uint8 }
type GenS3 struct { F0 bool; F1 int8 }
type GenS4 struct { F0 int16; F1 uint; F2 uint32 }
type GenS5 struct { F0 int64 }
func genFn0(p0 uint64) int {
    fmt.Println(p0)
    var v6 int16 = 1
    for fi7 := 0; fi7 < 19; fi7++ {
        v8 := []int{}
        var v9 float32 = 1e10
        _ = fi7
        _ = v8
        _ = v9
    }
    _ = p0
    _ = v6
    return -113
}
func genFn1(p0 int64, p1 uint16) int {
    v10 := []int{954, 1}
    if (p1 != (p1 | (p1 - 65535))) {
        var v11 int8 = -11
        v11 = (v11 | (v11 ^ (v11 >> ((42) & 63))))
        _ = v11
    }
    p0 -= (-(p0))
    _ = len(v10)
    _ = p0
    _ = p1
    _ = v10
    return genFn0(467)
}

func main() {
    var v1 bool = true
    var v2 bool = v1
    var v3 float64 = float64((0.0 + (3.14159 + (float32(((1.5 / ((1e10 / ((0.1 / (float64((0.1 * 0.1)) / (float64(1e10) + 0.1))) + 0.1)) - ((1.5 / (-2.25 * 1.5)) - float64((0.0 * (0.0 * 3.14159)))))) * (((1.5 + ((1.5 / 1.0) + 0.0)) * (0.1 + 3.14159)) + 1e10))) - ((1.5 / 1.0) * ((((0.1 / 1.0) / ((-2.25 - 1e10) / (1e10 - float32(0.1)))) * 1e10) / 1e10))))))
    var v4 uint = 255
    v5 := make(map[uint32]int)
    if ((!("a" < "\n\t")) && ("" >= "hello")) {
        v12 := GenS1{F0: 127, F1: v1, F2: 65535}
        fmt.Println(v2)
        fmt.Println(v3, v3)
        _ = v12
    }
    var v13 uint32 = 484
    v14 := GenS1{F0: 178, F1: (("foo bar" + "foo bar") < "foo bar"), F2: (v13 % (((v13 << ((595) & 63))) | 1))}
    var v15 uint32 = (v13 | 2)
    for fi16 := 0; fi16 < 3; fi16++ {
        if true {
            v5[(v13 ^ (^(v13)))] = 646
            switch (-(fi16)) {
            case 32768:
                v2 = (32767 == 42)
            default:
                v17 := GenS2{F0: (v3 == v3), F1: 29, F2: 1}
                _ = v17
            }
        } else {
            var v18 int64 = 0
            v19 := []int{genFn0(616), (^(fi16)), (fi16 << ((v4) & 63))}
            v3 = v3
            _ = v18
            _ = v19
        }
        var v20 int16 = 1000
        for fi21 := 0; fi21 < 10; fi21++ {
            var v22 int32 = 2147483647
            v23 := GenS3{F0: (!(v2 || ((!(-9223372036854775808 >= 42)) || (v2 && v1)))), F1: 0}
            _ = fi21
            _ = v22
            _ = v23
        }
        _ = fi16
        _ = v20
    }
    v15 += v15
    v15 = 1
    v24 := GenS3{F0: ("" < "foo bar"), F1: 1}
    var v25 uint64 = 1
    if (v25 < (v25 - (v25 >> (((v4 >> (((v4 >> ((v4) & 63))) & 63))) & 63)))) {
        var v26 uint = uint(v25)
        _ = v26
    }
    v27 := GenS4{F0: 1000, F1: uint(v25), F2: v15}
    v28 := GenS4{F0: -494, F1: v4, F2: v15}
    switch -403 {
    case -9223372036854775808:
        v29 := GenS5{F0: 437}
        _ = v29
    case 1:
        var v30 int32 = -141
        _ = v30
    case -2147483648:
        v31 := []int{42, -9223372036854775808, -272, 32767}
        _ = v31
    default:
        var v32 uint32 = v15
        _ = v32
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v24
    _ = v25
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
