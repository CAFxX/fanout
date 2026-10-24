package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint }
type GenS2 struct { F0 int16 }
func genFn0() uint32 {
    var v4 bool = false
    v5 := []int{-2147483648, -438, 32767}
    var v6 uint64 = 18446744073709551615
    _ = v4
    _ = v5
    _ = v6
    return 1
}
func genFn1(p0 uint8) string {
    fmt.Println(p0)
    var v7 int8 = 124
    v8 := GenS1{F0: -1, F1: 18446744073709551615}
    switch -552 {
    case 32768:
        v7 = (v7 ^ (v7 * (v7 << ((0) & 63))))
    case 1:
        if (!(882 != 18446744073709551615)) {
            v9 := []int{-2147483648, 32767}
            _ = v9
        }
    case 128:
        if (0 > 42) {
            v7 = (v7 + (v7 - 0))
            for fi10 := 0; fi10 < 1; fi10++ {
                var v11 string = "foo bar"
                _ = fi10
                _ = v11
            }
        } else {
            for fi12 := 0; fi12 < 18; fi12++ {
                var v13 uint8 = p0
                for fi14 := 0; fi14 < 14; fi14++ {
                    p0 = (p0 + (v13 % ((0) | 1)))
                    var v15 float64 = 1.5
                    _ = fi14
                    _ = v15
                }
                _ = fi12
                _ = v13
            }
            v16 := GenS2{F0: -32768}
            fmt.Println(p0, v7)
            _ = v16
        }
    default:
        var v17 uint32 = genFn0()
        _ = v17
    }
    _ = p0
    _ = v7
    _ = v8
    return "\n\t"
}

func main() {
    v1 := make(map[string]string)
    v2 := []string{"foo bar", "\n\t", "hello", ""}
    var v3 float64 = (((float64((0.1 / ((3.14159 * 0.1) - (-2.25 / 3.14159)))) / 1.5) + (0.0 / 1.5)) - 1.5)
    v3 = v3
    if (780 == 255) {
        var v18 float64 = float64(float32((v3 - float64(-2.25))))
        var v19 int = 4294967295
        fmt.Println(v19, v3)
        _ = v18
        _ = v19
    } else {
        v3 -= (float64(-2.25) * (v3 + (v3 / (v3 + 0.0))))
    }
    v3 -= float64(0.0)
    var v20 float64 = v3
    for fi21 := 0; fi21 < 4; fi21++ {
        var v22 int16 = -829
        v23 := v1[("" + "foo bar")]
        _ = fi21
        _ = v22
        _ = v23
    }
    var v24 float64 = (((3.14159 - v20) + ((1.5 / (1e10 / v20)) + 1.5)) * v3)
    var v25 int16 = 322
    fmt.Println(v24, v24)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v20
    _ = v24
    _ = v25
    _ = v3
}
