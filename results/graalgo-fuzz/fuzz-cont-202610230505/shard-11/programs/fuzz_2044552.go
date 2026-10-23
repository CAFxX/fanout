package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint }
type GenS2 struct { F0 string; F1 uint32 }
type GenS3 struct { F0 int16; F1 string }
func genFn0() uint16 {
    for fi5 := 0; fi5 < 8; fi5++ {
        var v6 int8 = -93
        _ = fi5
        _ = v6
    }
    var v7 uint8 = 1
    var v8 string = ""
    _ = v7
    _ = v8
    return 699
}
func genFn1(p0 int, p1 bool) uint32 {
    v9 := []int{(^(p0)), p0, (p0 & -253)}
    v10 := GenS1{F0: genFn0(), F1: 559}
    if (1000 >= 1) {
        var v11 int64 = 0
        _ = v11
    }
    fmt.Println(p0, p1)
    _ = p0
    _ = p1
    _ = v10
    _ = v9
    return 259
}

func main() {
    var v1 int64 = -9223372036854775808
    var v2 uint = 1
    v3 := []string{"foo bar", ""}
    v4 := make(map[int]bool)
    _ = len(v4)
    v1 -= 807
    v2 -= (v2 * (v2 / ((v2) | 1)))
    switch "a" {
    case "\n\t":
        var v12 int64 = v1
        _ = v12
    }
    v2 -= v2
    for fi13 := 0; fi13 < 10; fi13++ {
        switch 127 {
        case 42:
            var v14 bool = false
            _ = v14
        case 127:
            delete(v4, (fi13 - fi13))
        default:
            v2 = 162
        }
        v4[(fi13 >> ((v2) & 63))] = ((0.1 <= float64((((float32((-2.25 + (((0.1 + (0.1 * 0.0)) / 1.0) + 1.5))) - ((1.5 + 1e10) - (0.1 * 0.0))) * 1e10) - (0.0 + (float32((-2.25 / 1.0)) * 0.1))))) || false)
        for fi15 := 0; fi15 < 10; fi15++ {
            switch 4294967295 {
            case -887:
                if true {
                    if true {
                        var v16 bool = false
                        _ = v16
                    } else {
                        switch 1 {
                        case -1:
                            v4[fi15] = true
                        default:
                            if len(v3) > 0 {
                                v3[0] = "hello"
                            }
                        }
                        delete(v4, (fi15 - fi13))
                    }
                    var v17 int16 = -505
                    v18 := GenS1{F0: genFn0(), F1: v2}
                    _ = v17
                    _ = v18
                }
            }
            v1 = (v1 | 42)
            fmt.Println(fi13, fi15)
            _ = fi15
        }
        _ = fi13
    }
    v2 = v2
    delete(v4, 127)
    v19 := GenS2{F0: "a", F1: 65535}
    if (false && false) {
        for fi20 := 0; fi20 < 1; fi20++ {
            for fi21 := 0; fi21 < 9; fi21++ {
                switch (fi21 >> (((v2 + v2)) & 63)) {
                case 127:
                    var v22 int64 = 1
                    _ = v22
                default:
                    v2 -= v2
                }
                _ = fi21
            }
            var v23 uint16 = 749
            if true {
                fmt.Println(v1)
            }
            _ = fi20
            _ = v23
        }
        var v24 int32 = -2147483648
        v24 = (-(v24))
        _ = v24
    }
    fmt.Println(v2, v1)
    if len(v3) > 0 {
        v3[0] = "a"
    }
    switch 42 {
    case -1:
        var v25 float64 = -2.25
        _ = v25
    }
    var v26 int8 = -1
    switch "" {
    case "foo bar":
        if true {
            for fi27 := 0; fi27 < 14; fi27++ {
                var v28 int32 = -305
                if (18446744073709551615 > 42) {
                    v29 := GenS1{F0: genFn0(), F1: 1}
                    _ = v29
                } else {
                    var v30 uint8 = 48
                    var v31 string = "hello"
                    var v32 uint16 = 836
                    _ = v30
                    _ = v31
                    _ = v32
                }
                _ = fi27
                _ = v28
            }
            v3 = append(v3, "foo bar")
            v26 = v26
        } else {
            v2 += v2
        }
    case "\n\t":
        v33 := GenS3{F0: -32768, F1: "foo bar"}
        _ = v33
    case "":
        v34 := GenS3{F0: 1, F1: "a"}
        _ = v34
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v19
    _ = v2
    _ = v26
    _ = v3
    _ = v4
}
