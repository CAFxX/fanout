package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8; F2 uint64 }
type GenS2 struct { F0 int32; F1 int64; F2 uint }
func genFn0(p0 int8, p1 int64) uint64 {
    p1 += -121
    p0 = (p0 - p0)
    switch 42 {
    case 0:
        var v7 uint32 = 564
        _ = v7
    default:
        if (p0 != (p0 - -123)) {
            var v8 uint32 = 243
            _ = v8
        } else {
            v9 := GenS1{F0: false, F1: (p0 & (p0 - (p0 * (p0 & p0)))), F2: 42}
            _ = v9
        }
    }
    for fi10 := 0; fi10 < 1; fi10++ {
        p1 += (p1 / ((0) | 1))
        _ = fi10
    }
    p0 = (p0 % (((p0 | p0)) | 1))
    _ = p0
    _ = p1
    return 2
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 int32 = 1
    var v3 float32 = float32(((0.1 + (1.5 * (3.14159 + 1.5))) - ((0.1 - (0.1 - 1e10)) + (0.0 / (((0.1 + (float64(3.14159) * 3.14159)) / -2.25) + 1.5)))))
    v4 := make(map[uint32]bool)
    v5 := make(map[string]string)
    v6 := make(map[int]string)
    switch "hello" {
    case "":
        _ = len(v6)
    }
    for fi11 := 0; fi11 < 14; fi11++ {
        var v12 uint16 = 1
        for fi13 := 0; fi13 < 14; fi13++ {
            v14 := GenS1{F0: false, F1: -1, F2: genFn0(int8(v12), 9223372036854775807)}
            v15 := v5[(("\n\t" + "") + "")]
            v15 = v15
            _ = fi13
            _ = v14
            _ = v15
        }
        _ = fi11
        _ = v12
    }
    switch 629 {
    case 32768:
        if ("foo bar" == (("" + "hello") + "\n\t")) {
            _ = len(v1)
            v2 += v2
            var v16 string = "foo bar"
            _ = v16
        }
    }
    if (0 > 65535) {
        if true {
            var v17 bool = (32767 != 32767)
            var v18 uint = 845
            _ = v17
            _ = v18
        } else {
            v2 = v2
        }
        v3 = float32(3.14159)
    }
    var v19 uint16 = 0
    var v20 float64 = float64((v3 + (v3 * 3.14159)))
    for fi21 := 0; fi21 < 12; fi21++ {
        var v22 uint8 = 212
        _ = fi21
        _ = v22
    }
    if (0 != 1) {
        v3 -= ((float32((0.1 - v20)) * (1.5 + v3)) * v3)
        v19 -= (^(v19))
        v23 := GenS1{F0: (!(false || true)), F1: -83, F2: genFn0(4, 74)}
        _ = v23
    }
    v2 += v2
    if (!true) {
        switch "\n\t" {
        case "foo bar":
            for fi24 := 0; fi24 < 5; fi24++ {
                v25 := GenS1{F0: ((false || (42 > 254)) && true), F1: 127, F2: 0}
                _ = fi24
                _ = v25
            }
        case "\n\t":
            var v26 float32 = v3
            _ = v26
        case "":
            var v27 int8 = 42
            _ = v27
        }
        fmt.Println(v20, v19, v3)
        if true {
            fmt.Println(v2)
            v19 -= (v19 | v19)
            var v28 string = "foo bar"
            _ = v28
        } else {
            for fi29 := 0; fi29 < 16; fi29++ {
                if ((v19 - v19) != v19) {
                    v19 = v19
                    v30 := []int{}
                    for fi31 := 0; fi31 < 6; fi31++ {
                        if (!true) {
                            switch ((((("foo bar" + "\n\t") + "foo bar") + "hello") + "foo bar") + "\n\t") {
                            case "a":
                                fmt.Println(v20, v20, v2)
                            case "hello":
                                if len(v30) > 0 {
                                    v30[0] = fi31
                                }
                            default:
                                v19 = (v19 * v19)
                            }
                        }
                        _ = fi31
                    }
                    _ = v30
                }
                _ = fi29
            }
        }
    }
    v32 := []int{32768, 32767, 1, 133}
    v6[954] = "\n\t"
    var v33 int = -9223372036854775808
    v3 += 1e10
    v34 := GenS2{F0: (v2 ^ 0), F1: -9223372036854775808, F2: 1}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v32:", len(v32), v32)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
