package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 string; F1 int }
func genFn0(p0 uint64, p1 bool, p2 bool) uint8 {
    v4 := GenS1{F0: 32767}
    var v5 string = ""
    v6 := []int{1, 137}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    return 2
}
func genFn1(p0 int8, p1 uint, p2 int32) uint8 {
    if (42 < -964) {
        var v7 int32 = 2147483647
        fmt.Println(p0)
        _ = v7
    }
    v8 := GenS1{F0: 1000}
    if (127 >= 1) {
        switch "\n\t" {
        case "\n\t":
            v9 := []int{}
            _ = v9
        default:
            var v10 uint16 = 32767
            _ = v10
        }
        v11 := []int{-1}
        fmt.Println(p2, p0, p2)
        _ = v11
    } else {
        v12 := []int{9223372036854775807, 0, 202, 706}
        _ = v12
    }
    v13 := []int{32767}
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v8
    return genFn0(4294967295, false, (!(2147483647 >= 1)))
}

func main() {
    var v1 uint8 = 127
    var v2 int16 = 32767
    var v3 float64 = float64((1.5 / 3.14159))
    if false {
        var v14 int = 42
        if ((3.14159 / (v3 / ((float64(1e10) + -2.25) / 3.14159))) != (0.0 - float64((0.1 / float32(v3))))) {
            fmt.Println(v1)
            v1 += (v1 << ((18446744073709551615) & 63))
        } else {
            v15 := []int{}
            if (false || ((v14 << ((1) & 63)) != v14)) {
                var v16 float64 = float64(-2.25)
                _ = v16
            } else {
                var v17 bool = true
                switch "foo bar" {
                case "foo bar":
                    for fi18 := 0; fi18 < 3; fi18++ {
                        fmt.Println(fi18, v1)
                        _ = fi18
                    }
                case "a":
                    if len(v15) > 0 {
                        v15[0] = 127
                    }
                case "\n\t":
                    v2 = (v2 | (v2 % ((323) | 1)))
                default:
                    v14 -= v14
                }
                switch (v14 / ((v14) | 1)) {
                case -2147483648:
                    var v19 uint = 18446744073709551615
                    _ = v19
                case -1:
                    v3 += v3
                case -484:
                    v3 = ((v3 * v3) / (v3 + 1.5))
                }
                _ = v17
            }
            _ = v15
        }
        _ = v14
    }
    v2 += v2
    v20 := []int{128, -509}
    v2 = v2
    var v21 int32 = -494
    v20 = append(v20, -9223372036854775808)
    var v22 uint16 = 519
    v23 := GenS2{F0: "", F1: 32767}
    v21 += 2147483647
    if (!false) {
        v1 += 244
        var v24 uint64 = 1
        if ((v24 == (v24 + 18446744073709551615)) || (!true)) {
            if false {
                for fi25 := 0; fi25 < 11; fi25++ {
                    fmt.Println(v22, v1)
                    var v26 float32 = (-2.25 - 0.0)
                    v27 := GenS2{F0: "\n\t", F1: 240}
                    _ = fi25
                    _ = v26
                    _ = v27
                }
                if (-1 >= -1) {
                    if len(v20) > 0 {
                        v20[0] = 42
                    }
                    v28 := GenS2{F0: ("hello" + "a"), F1: -802}
                    _ = v28
                } else {
                    v3 = (float64(-2.25) / (v3 + 0.0))
                    var v29 uint16 = (v22 | 1)
                    _ = v29
                }
            }
        }
        _ = v24
    } else {
        for fi30 := 0; fi30 < 2; fi30++ {
            var v31 bool = true
            v32 := GenS1{F0: v2}
            _ = fi30
            _ = v31
            _ = v32
        }
        if len(v20) >= 2 {
            v20 = v20[1:len(v20)]
        }
    }
    if len(v20) > 0 {
        v20[0] = -317
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
}
