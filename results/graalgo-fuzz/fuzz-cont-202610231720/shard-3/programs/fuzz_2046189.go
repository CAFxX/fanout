package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32; F2 int }
type GenS2 struct { F0 bool; F1 int64; F2 int16 }
type GenS3 struct { F0 int64; F1 uint }
type GenS4 struct { F0 bool; F1 uint16 }
type GenS5 struct { F0 string; F1 int16 }
func genFn0(p0 int8, p1 uint16) uint64 {
    fmt.Println(p0, p1)
    var v6 uint32 = 2147483647
    var v7 bool = false
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 18446744073709551615
}

func main() {
    v1 := []uint32{65535, 2}
    v2 := []int{32767}
    v3 := make(map[int]int)
    var v4 bool = (3.14159 <= ((-2.25 / 1.0) + float32((-2.25 * 1e10))))
    v5 := make(map[string]string)
    v1 = append(v1, 899)
    v8 := GenS1{F0: 42, F1: 773, F2: 2147483647}
    var v9 int16 = 601
    _ = len(v3)
    v10 := GenS1{F0: 0, F1: -829, F2: 32767}
    v11 := GenS2{F0: (2147483647 == 0), F1: 42, F2: v9}
    if false {
        var v12 uint8 = 152
        if v4 {
            _ = len(v5)
        }
        _ = v12
    } else {
        for fi13 := 0; fi13 < 14; fi13++ {
            if v4 {
                v9 += 1000
                v1 = append(v1, 149)
            } else {
                if (!(255 > 44)) {
                    if false {
                        v9 = v9
                        if (v9 <= v9) {
                            for fi14 := 0; fi14 < 15; fi14++ {
                                v4 = false
                                switch "hello" {
                                case "hello":
                                    var v15 uint8 = 251
                                    _ = v15
                                case "\n\t":
                                    var v16 uint8 = 1
                                    _ = v16
                                }
                                fmt.Println(fi14)
                                _ = fi14
                            }
                            v17 := GenS3{F0: -9223372036854775808, F1: 2}
                            var v18 int = (fi13 ^ fi13)
                            _ = v17
                            _ = v18
                        } else {
                            if (!true) {
                                v9 = v9
                                v9 += v9
                                switch 127 {
                                case 709:
                                    v19 := GenS3{F0: 1, F1: 4294967295}
                                    _ = v19
                                case -500:
                                    delete(v3, fi13)
                                case -198:
                                    v20 := GenS4{F0: ((-2.25 <= (1e10 * (((-2.25 * 1e10) * 0.0) + -2.25))) && true), F1: 255}
                                    _ = v20
                                default:
                                    v21 := GenS1{F0: -2147483648, F1: 335, F2: (fi13 ^ 42)}
                                    _ = v21
                                }
                            }
                            delete(v3, 1)
                        }
                        v3[fi13] = (fi13 / ((1) | 1))
                    } else {
                        switch "foo bar" {
                        case "a":
                            fmt.Println(v9)
                        default:
                            _ = len(v3)
                        }
                    }
                    v9 -= 1
                }
                v4 = true
            }
            v22 := v5[("a" + "")]
            v9 += (v9 - v9)
            _ = fi13
            _ = v22
        }
        switch ("a" + ("\n\t" + ("hello" + "foo bar"))) {
        case "a":
            fmt.Println(v9, v9)
        default:
            v9 = 1
        }
    }
    v4 = ((3.14159 + 1.5) <= float64(3.14159))
    var v23 int64 = -375
    v23 += 0
    switch -795 {
    case -496:
        v4 = false
    case 0:
        for fi24 := 0; fi24 < 13; fi24++ {
            delete(v5, "foo bar")
            v25 := GenS1{F0: 0, F1: 2147483647, F2: (fi24 / ((1) | 1))}
            _ = fi24
            _ = v25
        }
    }
    v5[(("" + (("a" + "a") + "hello")) + ("" + "a"))] = ""
    v26 := GenS5{F0: "", F1: (v9 ^ v9)}
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    _ = len(v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v23
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
