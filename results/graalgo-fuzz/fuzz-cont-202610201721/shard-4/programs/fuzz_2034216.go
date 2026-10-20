package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8 }
func genFn0(p0 uint64, p1 uint64) int32 {
    v5 := []int{0, -357}
    for fi6 := 0; fi6 < 9; fi6++ {
        var v7 bool = (!false)
        v8 := GenS1{F0: (p1 | 0), F1: 42}
        switch "foo bar" {
        case "":
            if len(v5) > 0 {
                v5[0] = -807
            }
        case "hello":
            fmt.Println(p0)
        case "a":
            var v9 int = fi6
            _ = v9
        }
        _ = fi6
        _ = v7
        _ = v8
    }
    switch ("a" + ("" + ("" + "foo bar"))) {
    case "hello":
        p1 += p1
    case "":
        var v10 int32 = 0
        _ = v10
    case "a":
        for fi11 := 0; fi11 < 4; fi11++ {
            var v12 uint8 = 42
            if (!false) {
                fmt.Println(p1, fi11)
                p0 = (p1 + (^(p0)))
                p1 = 18446744073709551615
            } else {
                if (false && true) {
                    if true {
                        v13 := GenS1{F0: p0, F1: v12}
                        _ = v13
                    } else {
                        if (-265 > 1) {
                            fmt.Println(p0, v12, v12)
                            switch (fi11 * fi11) {
                            case -9223372036854775808:
                                v12 = v12
                            case 1:
                                v14 := GenS1{F0: (p1 / (((p1 << ((4294967295) & 63))) | 1)), F1: v12}
                                _ = v14
                            case 0:
                                v12 -= 1
                            default:
                                if (3.14159 != -2.25) {
                                    if (-128 <= 127) {
                                        var v15 uint8 = (v12 / ((v12) | 1))
                                        _ = v15
                                    }
                                    v5 = append(v5, (fi11 + 9223372036854775807))
                                }
                            }
                        }
                    }
                } else {
                    switch (fi11 | 9223372036854775807) {
                    case 4:
                        var v16 int8 = 127
                        _ = v16
                    default:
                        for fi17 := 0; fi17 < 9; fi17++ {
                            switch "\n\t" {
                            case "":
                                v18 := GenS1{F0: p1, F1: 43}
                                _ = v18
                            case "a":
                                v12 -= 42
                            default:
                                switch ("a" + "foo bar") {
                                case "\n\t":
                                    if ((fi11 - fi17) <= fi11) {
                                        switch "\n\t" {
                                        case "":
                                            v12 = v12
                                        default:
                                            var v19 int32 = 0
                                            _ = v19
                                        }
                                        p1 += p0
                                        if len(v5) > 0 {
                                            v5[0] = 9223372036854775807
                                        }
                                    }
                                case "foo bar":
                                    v20 := GenS1{F0: p1, F1: (v12 / ((v12) | 1))}
                                    _ = v20
                                }
                            }
                            p0 += p0
                            p0 += p1
                            _ = fi17
                        }
                    }
                }
            }
            var v21 uint16 = 895
            _ = fi11
            _ = v12
            _ = v21
        }
    default:
        _ = len(v5)
    }
    var v22 bool = false
    _ = p0
    _ = p1
    _ = v22
    _ = v5
    return 0
}

func main() {
    var v1 float64 = (-2.25 + (((1.5 - 1e10) - (0.0 - (1e10 - (((float64(3.14159) / 1.0) + 1.5) / 1.5)))) / 1.5))
    var v2 float32 = 1.5
    var v3 uint64 = 150
    v4 := make(map[string]uint32)
    v23 := []int{}
    fmt.Println(v3, v3, v2)
    for fi24 := 0; fi24 < 16; fi24++ {
        var v25 uint = 936
        _ = fi24
        _ = v25
    }
    v3 += (^(v3))
    for fi26 := 0; fi26 < 4; fi26++ {
        if false {
            v2 -= float32((v1 + -2.25))
        } else {
            for fi27 := 0; fi27 < 19; fi27++ {
                if len(v23) >= 2 {
                    v23 = v23[1:len(v23)]
                }
                v2 -= 1e10
                v1 = 3.14159
                _ = fi27
            }
        }
        _ = fi26
    }
    _ = len(v4)
    for fi28 := 0; fi28 < 13; fi28++ {
        _ = len(v23)
        _ = fi28
    }
    var v29 int32 = 653
    fmt.Println(v3)
    v29 = genFn0(1, (^(v3)))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v23
    _ = v29
    _ = v3
    _ = v4
}
