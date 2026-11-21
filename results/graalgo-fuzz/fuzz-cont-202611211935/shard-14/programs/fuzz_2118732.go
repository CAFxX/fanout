package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32 }
type GenS2 struct { F0 bool; F1 int16 }
type GenS3 struct { F0 uint16; F1 int64 }
type GenS4 struct { F0 uint }
func genFn0(p0 int8, p1 bool, p2 int16) uint8 {
    v7 := []int{}
    var v8 int = 79
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 128
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[string]string)
    var v3 uint64 = 42
    v4 := []uint32{2147483648, 0, 26, 2}
    var v5 float32 = 1e10
    var v6 int64 = -9223372036854775808
    var v9 int64 = (v6 & 511)
    v10 := GenS1{F0: (v6 * 42), F1: -2147483648}
    var v11 uint8 = 1
    var v12 uint16 = 223
    v4 = append(v4, 237)
    v6 = v9
    if false {
        for fi13 := 0; fi13 < 13; fi13++ {
            var v14 float32 = (v5 * float32((1e10 - 1e10)))
            switch -963 {
            case 844:
                var v15 float32 = v14
                _ = v15
            case 0:
                if (genFn0(100, (genFn0(1, ((v11 << ((471) & 63)) <= v11), 1) <= (v11 - genFn0(100, (false && false), 0))), 42) < (v11 << ((0) & 63))) {
                    switch (-(fi13)) {
                    case 42:
                        var v16 int16 = int16(v6)
                        _ = v16
                    default:
                        for fi17 := 0; fi17 < 19; fi17++ {
                            v11 += genFn0(0, true, 42)
                            v14 += (float32(((0.0 + 1.5) + ((3.14159 * -2.25) - 0.1))) / (0.0 / v5))
                            _ = fi17
                        }
                    }
                } else {
                    for fi18 := 0; fi18 < 14; fi18++ {
                        v3 = (v3 * 1)
                        switch "" {
                        case "a":
                            fmt.Println(v3)
                        case "\n\t":
                            var v19 bool = (-128 <= -128)
                            _ = v19
                        case "foo bar":
                            fmt.Println(v5, v9)
                        default:
                            delete(v2, "a")
                        }
                        _ = fi18
                    }
                }
            case -608:
                fmt.Println(v6, v11, v3)
            default:
                v20 := v1[fi13]
                _ = v20
            }
            v21 := GenS1{F0: (v9 ^ v6), F1: 685}
            _ = fi13
            _ = v14
            _ = v21
        }
        v22 := GenS2{F0: false, F1: 32767}
        _ = v22
    }
    switch "" {
    case "a":
        _ = len(v1)
    case "":
        v12 -= (v12 | (v12 ^ v12))
    }
    fmt.Println(v6, v5)
    v23 := GenS3{F0: (v12 >> ((4294967295) & 63)), F1: 1}
    v11 = genFn0(-1, false, 1)
    for fi24 := 0; fi24 < 4; fi24++ {
        v25 := GenS4{F0: 42}
        var v26 int16 = 32767
        var v27 int32 = -1
        _ = fi24
        _ = v25
        _ = v26
        _ = v27
    }
    fmt.Println(v5)
    v28 := GenS3{F0: v12, F1: v6}
    var v29 int32 = -379
    v4 = append(v4, 0)
    v30 := GenS2{F0: true, F1: 1}
    v4 = append(v4, uint32(v3))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v23
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
