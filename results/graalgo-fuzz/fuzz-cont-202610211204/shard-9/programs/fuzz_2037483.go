package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string; F2 int64 }
type GenS2 struct { F0 uint; F1 uint32; F2 uint64 }
func genFn0(p0 int32, p1 uint32, p2 bool) int64 {
    v5 := []int{}
    if (true && (42 >= 18446744073709551615)) {
        if (!(((1.5 - 0.1) <= float32(-2.25)) || true)) {
            fmt.Println(p0, p2)
        } else {
            for fi6 := 0; fi6 < 16; fi6++ {
                var v7 int32 = (p0 | p0)
                _ = fi6
                _ = v7
            }
            p2 = ((1 < 1) && p2)
        }
        p0 -= p0
        if false {
            p1 -= p1
        } else {
            fmt.Println(p1, p0)
            p2 = p2
            fmt.Println(p0)
        }
    } else {
        for fi8 := 0; fi8 < 12; fi8++ {
            p0 += (-(p0))
            _ = fi8
        }
        var v9 uint16 = 180
        _ = v9
    }
    if (93 <= 0) {
        v10 := GenS1{F0: -182, F1: (((("a" + "foo bar") + "foo bar") + ("\n\t" + ("foo bar" + "a"))) + "foo bar"), F2: 1}
        fmt.Println(p0, p2)
        var v11 float32 = 0.0
        _ = v10
        _ = v11
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 0
}

func main() {
    var v1 uint32 = 394
    var v2 string = "a"
    v3 := []uint32{}
    v4 := make(map[int]int)
    var v12 int64 = 535
    if (2 == 18446744073709551615) {
        var v13 float64 = 3.14159
        v1 = v1
        fmt.Println(v12, v12, v1)
        _ = v13
    }
    fmt.Println(v1, v1, v1)
    for fi14 := 0; fi14 < 6; fi14++ {
        fmt.Println(v12)
        _ = fi14
    }
    if false {
        if len(v3) > 0 {
            v3[0] = (v1 + (v1 | 1))
        }
        if true {
            fmt.Println(v1, v1, v1)
        } else {
            switch v2 {
            case "foo bar":
                v15 := GenS1{F0: 32767, F1: v2, F2: (v12 >> ((0) & 63))}
                _ = v15
            case "":
                var v16 uint8 = 42
                _ = v16
            case "\n\t":
                v12 += -9223372036854775808
            default:
                if (-2147483648 != 44) {
                    switch "hello" {
                    case "":
                        if ((^(v1)) <= (v1 / ((42) | 1))) {
                            var v17 string = v2
                            fmt.Println(v2)
                            var v18 float64 = (3.14159 / 1.5)
                            _ = v17
                            _ = v18
                        }
                    default:
                        v12 += (v12 - (v12 ^ (v12 | v12)))
                    }
                } else {
                    for fi19 := 0; fi19 < 19; fi19++ {
                        v20 := GenS2{F0: 65535, F1: (v1 & v1), F2: 2}
                        v12 = (v12 | v12)
                        v1 += v1
                        _ = fi19
                        _ = v20
                    }
                    var v21 int16 = 1000
                    _ = v21
                }
            }
            var v22 uint32 = (v1 + v1)
            _ = v22
        }
    }
    v23 := GenS1{F0: 4, F1: (((v2 + ("foo bar" + ("hello" + (("foo bar" + "hello") + (("a" + ("a" + (((v2 + ("foo bar" + v2)) + (v2 + v2)) + (((((v2 + (v2 + (v2 + ("foo bar" + "foo bar")))) + (v2 + "foo bar")) + (v2 + "\n\t")) + v2) + ("\n\t" + (v2 + ((v2 + v2) + "foo bar"))))))) + "\n\t"))))) + v2) + v2), F2: v12}
    var v24 int16 = 1
    var v25 uint16 = 42
    v3 = append(v3, v1)
    var v26 bool = (0 != 100)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v12
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
}
