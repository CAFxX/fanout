package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 string; F1 int16; F2 int16 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 int; F1 bool }
func genFn0(p0 uint8, p1 int32, p2 uint32) uint32 {
    v6 := GenS1{F0: 1}
    var v7 float64 = 0.0
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 948
}

func main() {
    var v1 float64 = 0.1
    var v2 int = -9223372036854775808
    var v3 int8 = 100
    var v4 int8 = (^(v3))
    var v5 uint64 = 42
    v8 := GenS1{F0: 4294967295}
    if (!((-2.25 / -2.25) >= float32(float64(((3.14159 * (0.0 / 1.5)) / (float32(v1) + float32(v1))))))) {
        for fi9 := 0; fi9 < 5; fi9++ {
            if (42 != 42) {
                v4 += v3
                v3 = v4
            } else {
                v10 := []int{}
                var v11 int16 = 1
                if len(v10) > 0 {
                    v10[0] = (fi9 | 42)
                }
                _ = v10
                _ = v11
            }
            for fi12 := 0; fi12 < 13; fi12++ {
                v13 := GenS2{F0: v5}
                v14 := GenS2{F0: v5}
                var v15 string = "\n\t"
                _ = fi12
                _ = v13
                _ = v14
                _ = v15
            }
            _ = fi9
        }
    } else {
        var v16 float64 = v1
        v17 := GenS3{F0: "foo bar", F1: -32768, F2: 1000}
        var v18 uint = 648
        _ = v16
        _ = v17
        _ = v18
    }
    v4 += v4
    fmt.Println(v4, v2)
    v1 = float64((-2.25 / 3.14159))
    v19 := []int{v2, (-(v2)), v2}
    switch ("hello" + "") {
    case "foo bar":
        var v20 string = ""
        _ = v20
    case "":
        var v21 float64 = v1
        _ = v21
    case "\n\t":
        v1 -= v1
    }
    v22 := GenS4{F0: 266}
    for fi23 := 0; fi23 < 14; fi23++ {
        switch fi23 {
        case 0:
            _ = len(v19)
        default:
            var v24 float32 = (((0.1 - 0.0) * (1.5 + 3.14159)) - (0.1 + float32(-2.25)))
            _ = v24
        }
        v4 -= (v4 ^ (v3 / ((v3) | 1)))
        _ = fi23
    }
    v25 := GenS4{F0: 0}
    v1 -= 0.1
    if (!true) {
        v3 = (v4 + v3)
    } else {
        v3 += -35
        v3 += v3
        if (0.0 == 1e10) {
            v26 := GenS5{F0: 127}
            for fi27 := 0; fi27 < 18; fi27++ {
                var v28 int16 = 1000
                _ = fi27
                _ = v28
            }
            v4 = v4
            _ = v26
        }
    }
    if (!false) {
        v29 := GenS2{F0: (v5 & (v5 / ((v5) | 1)))}
        _ = v29
    } else {
        if (v1 > -2.25) {
            var v30 int8 = v3
            v31 := GenS6{F0: (v2 % ((v2) | 1)), F1: false}
            _ = v30
            _ = v31
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v22
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
