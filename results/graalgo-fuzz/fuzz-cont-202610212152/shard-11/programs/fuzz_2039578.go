package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int8 }
func genFn0() uint32 {
    var v4 string = "hello"
    v4 = v4
    v5 := GenS1{F0: v4, F1: true}
    for fi6 := 0; fi6 < 13; fi6++ {
        var v7 int32 = 262
        var v8 int8 = 91
        _ = fi6
        _ = v7
        _ = v8
    }
    _ = v4
    _ = v5
    return 575
}
func genFn1(p0 uint8, p1 int64) int64 {
    p0 -= (p0 ^ p0)
    p1 = (p1 ^ 1)
    var v9 float32 = ((0.1 + 0.1) * -2.25)
    _ = p0
    _ = p1
    _ = v9
    return p1
}

func main() {
    var v1 uint = 0
    var v2 float32 = -2.25
    var v3 float32 = v2
    for fi10 := 0; fi10 < 6; fi10++ {
        var v11 float64 = -2.25
        _ = fi10
        _ = v11
    }
    v1 -= v1
    if (4294967295 != 1) {
        v1 = v1
        v12 := []int{-168, 128, 1}
        v13 := GenS1{F0: "", F1: (1 <= 93)}
        _ = v12
        _ = v13
    }
    v1 -= (v1 + 1)
    for fi14 := 0; fi14 < 3; fi14++ {
        v1 = 1
        _ = fi14
    }
    if true {
        v15 := GenS1{F0: "foo bar", F1: ((genFn1(42, -9223372036854775808) > 477) || true)}
        v16 := []int{9223372036854775807, 9223372036854775807, -27}
        _ = v15
        _ = v16
    } else {
        for fi17 := 0; fi17 < 17; fi17++ {
            for fi18 := 0; fi18 < 1; fi18++ {
                var v19 uint8 = 1
                v20 := []int{fi18, fi18, int(v19), (-(fi17))}
                _ = fi18
                _ = v19
                _ = v20
            }
            v21 := []int{(fi17 - -352), (fi17 ^ fi17), -9223372036854775808}
            if true {
                fmt.Println(v1)
            }
            _ = fi17
            _ = v21
        }
    }
    v1 += 390
    switch "" {
    case "hello":
        v22 := []int{}
        _ = v22
    default:
        v23 := GenS2{F0: 0}
        _ = v23
    }
    v24 := GenS3{F0: 66}
    if (100 > 19) {
        var v25 uint8 = 2
        v3 = 1e10
        _ = v25
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
}
