package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32; F2 int8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint32; F1 uint64; F2 int16 }
func genFn0(p0 uint32, p1 uint) int64 {
    p1 = (p1 >> (((p1 << ((65535) & 63))) & 63))
    fmt.Println(p0)
    v7 := GenS1{F0: "foo bar", F1: 2147483647, F2: -128}
    p1 = 2
    p0 = (p0 % ((p0) | 1))
    _ = p0
    _ = p1
    _ = v7
    return 1
}
func genFn1(p0 int8) bool {
    p0 += p0
    v8 := []int{127, -2147483648}
    if (1.5 == 3.14159) {
        if len(v8) > 0 {
            v8[0] = 675
        }
        p0 += p0
        var v9 uint = 896
        _ = v9
    } else {
        if true {
            _ = len(v8)
            v10 := GenS1{F0: "a", F1: 2147483647, F2: -128}
            _ = len(v8)
            _ = v10
        }
        var v11 float64 = 0.1
        _ = v11
    }
    var v12 float32 = 0.1
    if ((v12 + (v12 * (-2.25 + (-2.25 * 0.0)))) < float32((1.5 - -2.25))) {
        switch "foo bar" {
        case "hello":
            v12 = v12
        case "a":
            v13 := GenS2{F0: "foo bar"}
            _ = v13
        case "\n\t":
            for fi14 := 0; fi14 < 7; fi14++ {
                if (384 == 1) {
                    v15 := GenS2{F0: "a"}
                    if false {
                        p0 -= 127
                        var v16 int64 = -870
                        if (298 <= 898) {
                            v17 := GenS1{F0: "hello", F1: -2147483648, F2: 104}
                            fmt.Println(p0)
                            p0 = p0
                            _ = v17
                        }
                        _ = v16
                    }
                    _ = v15
                }
                _ = fi14
            }
        default:
            if len(v8) > 0 {
                v8[0] = 128
            }
        }
    }
    _ = p0
    _ = v12
    _ = v8
    return true
}

func main() {
    var v1 string = (("foo bar" + "foo bar") + ("" + ("" + "")))
    var v2 uint8 = 127
    v3 := []string{("a" + v1)}
    var v4 int16 = 0
    v5 := make(map[string]bool)
    var v6 bool = (false || true)
    v18 := v5[v1]
    switch 635 {
    case 32768:
        var v19 uint64 = 4294967295
        _ = v19
    default:
        v20 := GenS3{F0: 0}
        _ = v20
    }
    if ((1e10 * (1e10 * ((3.14159 + 1.5) * (0.0 * 3.14159)))) <= 1e10) {
        v1 = (v1 + (((v1 + "hello") + v1) + (v1 + v1)))
        v4 = v4
        v21 := GenS3{F0: 42}
        _ = v21
    } else {
        var v22 uint = 0
        v23 := GenS3{F0: 255}
        var v24 uint = (v22 & (v22 ^ 372))
        _ = v22
        _ = v23
        _ = v24
    }
    if v18 {
        v5["\n\t"] = true
        v4 += v4
    }
    for fi25 := 0; fi25 < 15; fi25++ {
        v18 = v18
        var v26 float64 = float64(((((3.14159 / 0.1) / 1.0) + -2.25) + float32((3.14159 / 1.0))))
        _ = fi25
        _ = v26
    }
    var v27 float64 = 3.14159
    v28 := GenS4{F0: 0, F1: 653, F2: v4}
    fmt.Println(v18, v18)
    v6 = v6
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
