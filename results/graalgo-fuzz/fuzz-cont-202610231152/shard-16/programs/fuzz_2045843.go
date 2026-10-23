package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint64; F1 int64; F2 uint64 }
func genFn0(p0 uint16, p1 uint32, p2 int64) int16 {
    p2 = -119
    v7 := []int{}
    fmt.Println(p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 1000
}
func genFn1(p0 int16, p1 bool, p2 int32) bool {
    p2 -= p2
    switch ("\n\t" + "foo bar") {
    case "\n\t":
        p2 -= 0
    case "hello":
        v8 := []int{-18, 32767}
        _ = v8
    default:
        v9 := GenS1{F0: -128}
        _ = v9
    }
    if p1 {
        v10 := []int{9223372036854775807, 592}
        _ = v10
    } else {
        p1 = p1
        if (65535 >= 42) {
            var v11 int64 = -9223372036854775808
            var v12 int = -1
            _ = v11
            _ = v12
        } else {
            var v13 int32 = 2147483647
            _ = v13
        }
    }
    switch 42 {
    case -1:
        fmt.Println(p1)
    case 2147483647:
        var v14 uint8 = 128
        _ = v14
    default:
        v15 := GenS1{F0: -9}
        _ = v15
    }
    _ = p0
    _ = p1
    _ = p2
    return (!true)
}

func main() {
    var v1 int = -2147483648
    v2 := []string{"", "hello", (("\n\t" + "hello") + "hello"), ""}
    var v3 uint = 2
    v4 := make(map[string]string)
    v5 := []string{"\n\t", "a", ""}
    v6 := []string{"", "\n\t", "a", "a"}
    if (((731 != 9223372036854775807) || false) && (0 != 220)) {
        var v16 uint8 = 127
        v1 += v1
        _ = v16
    }
    fmt.Println(v3)
    switch 9223372036854775807 {
    case -249:
        v3 = v3
    case 2147483647:
        var v17 bool = false
        _ = v17
    }
    var v18 int8 = 0
    if (true && false) {
        if (1.5 >= (0.0 + 1e10)) {
            var v19 bool = true
            _ = v19
        }
    }
    v5 = append(v5, "hello")
    if ((!false) || true) {
        var v20 uint64 = 0
        _ = v20
    } else {
        switch v1 {
        case 42:
            var v21 int16 = 228
            _ = v21
        default:
            for fi22 := 0; fi22 < 4; fi22++ {
                fmt.Println(v3, v18, v18)
                if (!false) {
                    fmt.Println(v3)
                } else {
                    if (fi22 <= 4294967295) {
                        if true {
                            for fi23 := 0; fi23 < 19; fi23++ {
                                v1 = fi22
                                if (1.5 == (float32(float64(((((1e10 + 0.1) - (((float32(1e10) + 1e10) * 1.5) + 0.1)) * (0.0 * float32(((0.0 - (-2.25 + float64(1e10))) + 1e10)))) - 1e10))) * ((1e10 - float32(float64((0.1 - (0.1 / (0.1 * 1e10)))))) * ((0.1 / -2.25) - 0.1)))) {
                                    v3 = (v3 - v3)
                                    fmt.Println(v18, v3)
                                }
                                _ = fi23
                            }
                            v24 := v4["\n\t"]
                            _ = v24
                        }
                        v25 := GenS2{F0: 415, F1: -576, F2: 4294967295}
                        var v26 int8 = (v18 * (v18 & (v18 ^ v18)))
                        _ = v25
                        _ = v26
                    } else {
                        v3 -= 547
                        var v27 float32 = 0.0
                        _ = len(v6)
                        _ = v27
                    }
                }
                v18 = v18
                _ = fi22
            }
        }
        v28 := GenS2{F0: 470, F1: 1, F2: 18446744073709551615}
        for fi29 := 0; fi29 < 17; fi29++ {
            fmt.Println(fi29)
            v30 := v4[""]
            fmt.Println(v30)
            _ = fi29
            _ = v30
        }
        _ = v28
    }
    if false {
        var v31 uint16 = 0
        _ = v31
    } else {
        v3 = (v3 ^ 914)
        _ = len(v4)
    }
    var v32 uint32 = 4294967295
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    var v33 uint16 = 2
    v34 := GenS2{F0: 18446744073709551615, F1: 0, F2: 2}
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
