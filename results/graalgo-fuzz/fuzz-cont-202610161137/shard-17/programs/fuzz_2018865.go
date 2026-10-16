package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool; F2 int32 }
type GenS2 struct { F0 bool }
func genFn0(p0 uint, p1 int16, p2 int32) int32 {
    fmt.Println(p2, p2, p0)
    v6 := []int{0, -2147483648, 9223372036854775807}
    for fi7 := 0; fi7 < 15; fi7++ {
        var v8 uint32 = 2147483648
        switch fi7 {
        case 127:
            fmt.Println(v8)
        default:
            var v9 int = fi7
            _ = v9
        }
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        _ = fi7
        _ = v8
    }
    p0 += (p0 >> (((p0 % (((p0 >> ((p0) & 63))) | 1))) & 63))
    fmt.Println(p0, p2, p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return (^(p2))
}
func genFn1(p0 uint16, p1 int8) uint32 {
    var v10 uint = 403
    p1 += -128
    _ = p0
    _ = p1
    _ = v10
    return 360
}

func main() {
    var v1 int8 = 0
    var v2 uint32 = 42
    v3 := []uint32{(v2 - 42), v2}
    var v4 int8 = v1
    var v5 int8 = v4
    var v11 int32 = -1
    for fi12 := 0; fi12 < 10; fi12++ {
        switch "" {
        case "":
            switch ("" + "a") {
            case "a":
                v4 += v4
            default:
                for fi13 := 0; fi13 < 14; fi13++ {
                    if (-377 < -9223372036854775808) {
                        if len(v3) > 0 {
                            v3[0] = (v2 << ((1) & 63))
                        }
                        fmt.Println(v2, fi13, v5)
                        var v14 bool = true
                        _ = v14
                    } else {
                        v4 -= (v1 - 52)
                        v2 += 0
                        v4 += -128
                    }
                    if ((!((fi13 * fi13) >= 0)) || (0.0 > 0.0)) {
                        var v15 uint = 42
                        if false {
                            switch "hello" {
                            case "hello":
                                var v16 uint32 = v2
                                _ = v16
                            case "foo bar":
                                if len(v3) > 0 {
                                    v3[0] = v2
                                }
                            default:
                                v3 = append(v3, (^(v2)))
                            }
                        } else {
                            switch 4294967295 {
                            case -872:
                                _ = len(v3)
                            }
                            switch "hello" {
                            case "foo bar":
                                var v17 int32 = (v11 | v11)
                                _ = v17
                            case "a":
                                v11 += (^(v11))
                            default:
                                v18 := GenS1{F0: -630, F1: (!true), F2: 404}
                                _ = v18
                            }
                            if (((1e10 + (((1e10 / (float64(1e10) * 0.1)) + float64(3.14159)) + 1e10)) == 0.0) || true) {
                                v11 -= (v11 / (((^(v11))) | 1))
                                var v19 int32 = 2147483647
                                _ = v19
                            } else {
                                v1 -= 116
                                v20 := GenS2{F0: true}
                                _ = v20
                            }
                        }
                        _ = v15
                    } else {
                        if (4294967295 != 533) {
                            var v21 int = (fi12 >> ((42) & 63))
                            _ = v21
                        }
                    }
                    _ = fi13
                }
            }
        }
        _ = fi12
    }
    var v22 float64 = 1.5
    switch ("a" + "foo bar") {
    case "a":
        for fi23 := 0; fi23 < 15; fi23++ {
            v5 -= 0
            v5 = (-(v4))
            if true {
                v5 -= (^(v1))
            } else {
                v11 += -1
                v24 := GenS2{F0: (v11 < (v11 << ((0) & 63)))}
                _ = v24
            }
            _ = fi23
        }
    case "":
        _ = len(v3)
    case "foo bar":
        var v25 int8 = v5
        _ = v25
    default:
        var v26 int16 = -32768
        _ = v26
    }
    var v27 int16 = 481
    v22 += v22
    var v28 float64 = v22
    v29 := GenS2{F0: (-2147483648 != genFn0(988, (v27 - (v27 * -613)), (v11 % (((^(v11))) | 1))))}
    if (!("\n\t" != "\n\t")) {
        v28 += 1e10
    } else {
        if ((0 != 2) || (810 < 32767)) {
            fmt.Println(v1, v22, v11)
            v4 += 42
        }
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v2
    _ = v22
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
