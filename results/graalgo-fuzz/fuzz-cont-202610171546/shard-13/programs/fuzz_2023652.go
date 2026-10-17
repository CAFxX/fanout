package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string; F2 int32 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint; F1 uint8; F2 string }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int64; F1 int8 }
func genFn0(p0 int, p1 uint16, p2 int) uint16 {
    v7 := []int{32767, 0, p0}
    if len(v7) > 0 {
        v7[0] = p0
    }
    p2 += 32768
    p1 = (^(p1))
    switch "foo bar" {
    case "foo bar":
        switch p0 {
        case 4294967295:
            switch p2 {
            case 128:
                switch -9223372036854775808 {
                case 127:
                    var v8 uint8 = 17
                    _ = v8
                case 9223372036854775807:
                    v7 = append(v7, (p0 + p2))
                }
            case 127:
                var v9 int16 = 395
                _ = v9
            case 2147483647:
                _ = len(v7)
            default:
                p1 = p1
            }
        case 669:
            if (2147483647 > -2147483648) {
                var v10 uint64 = 42
                fmt.Println(p1, v10, p0)
                var v11 int16 = 124
                _ = v10
                _ = v11
            }
        default:
            p1 = p1
        }
    case "\n\t":
        p0 = p2
    case "hello":
        p1 = (^(p1))
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return (p1 * (p1 + (p1 << ((65535) & 63))))
}

func main() {
    var v1 uint32 = 444
    var v2 uint16 = 98
    v3 := make(map[uint32]bool)
    var v4 uint8 = 185
    var v5 uint32 = (v1 / (((^(v1))) | 1))
    v6 := []string{("a" + ("hello" + ("hello" + "a"))), "a"}
    if false {
        fmt.Println(v2)
        _ = len(v6)
    }
    var v12 int = 0
    v1 -= uint32(v4)
    v13 := GenS1{F0: v4, F1: "\n\t", F2: 1}
    v1 -= uint32(v4)
    v2 = uint16(v4)
    var v14 float32 = 0.0
    switch -9223372036854775808 {
    case 2147483647:
        v6 = append(v6, "a")
    case 1:
        var v15 bool = false
        _ = v15
    }
    if false {
        v2 -= (v2 | (^(v2)))
        v16 := GenS1{F0: uint8(v5), F1: "hello", F2: 2147483647}
        _ = v16
    }
    switch v12 {
    case 128:
        delete(v3, v1)
    case 127:
        v17 := GenS1{F0: 127, F1: "\n\t", F2: -2147483648}
        _ = v17
    }
    v2 -= genFn0((v12 >> ((664) & 63)), 2, 128)
    v2 -= v2
    switch "foo bar" {
    case "hello":
        v1 = uint32(v4)
    case "a":
        var v18 bool = ((65535 < 4294967295) && true)
        _ = v18
    case "\n\t":
        var v19 float64 = 3.14159
        _ = v19
    }
    switch v12 {
    case -911:
        fmt.Println(v5, v2)
    }
    var v20 int16 = -1
    for fi21 := 0; fi21 < 12; fi21++ {
        for fi22 := 0; fi22 < 19; fi22++ {
            if (!true) {
                _ = len(v3)
                v12 = int(v1)
                if true {
                    fmt.Println(v12, v14)
                } else {
                    v12 = v12
                    var v23 int8 = -37
                    _ = v23
                }
            }
            _ = fi22
        }
        if len(v6) > 0 {
            v6[0] = "a"
        }
        switch -916 {
        case -922:
            for fi24 := 0; fi24 < 1; fi24++ {
                v25 := GenS2{F0: 42}
                v4 = (^(v4))
                var v26 uint = 77
                _ = fi24
                _ = v25
                _ = v26
            }
        case 32768:
            var v27 uint64 = 357
            _ = v27
        }
        _ = fi21
    }
    for fi28 := 0; fi28 < 18; fi28++ {
        if false {
            var v29 uint8 = (v4 ^ (v4 & (v4 | 128)))
            switch "" {
            case "\n\t":
                for fi30 := 0; fi30 < 4; fi30++ {
                    fmt.Println(v2, fi28)
                    if true {
                        v1 = 2147483648
                    } else {
                        v20 += (v20 ^ (-(v20)))
                        var v31 uint8 = (v29 + (v4 ^ (v29 >> ((445) & 63))))
                        v32 := GenS3{F0: 0, F1: v29, F2: ("" + ("\n\t" + ("a" + ("" + "\n\t"))))}
                        _ = v31
                        _ = v32
                    }
                    switch (fi30 - -9223372036854775808) {
                    case -1:
                        v33 := GenS2{F0: 18446744073709551615}
                        _ = v33
                    case 32768:
                        v2 = (v2 * 255)
                    case 32767:
                        var v34 string = "foo bar"
                        _ = v34
                    default:
                        var v35 float32 = float32(((0.0 * 0.0) / 1.5))
                        _ = v35
                    }
                    _ = fi30
                }
            case "foo bar":
                v14 = v14
            case "a":
                v36 := v3[682]
                _ = v36
            default:
                v37 := GenS3{F0: 907, F1: 127, F2: (("" + "") + "foo bar")}
                _ = v37
            }
            _ = v29
        } else {
            if true {
                for fi38 := 0; fi38 < 3; fi38++ {
                    if true {
                        var v39 float32 = 0.0
                        if (1 < -9223372036854775808) {
                            var v40 int16 = (v20 | (v20 / ((42) | 1)))
                            v6 = append(v6, "hello")
                            v3[v1] = ((v4 < (v4 >> ((848) & 63))) || (42 >= 190))
                            _ = v40
                        } else {
                            if (104 == 0) {
                                fmt.Println(fi38)
                                v3[65535] = (-9223372036854775808 == 0)
                                fmt.Println(v5, v12)
                                if true {
                                    if (false || false) {
                                        v3[2] = ((444 == 724) && ((0.0 == (-2.25 - (1e10 - float64(((((float32(-2.25) * 1e10) * v14) - float32((1.5 / ((0.0 + 1e10) - (1.5 + float64((v14 + v14))))))) + 1.5))))) || true))
                                        v41 := GenS4{F0: (^(v5))}
                                        _ = v41
                                    } else {
                                        var v42 uint32 = uint32(v2)
                                        var v43 int32 = 860
                                        _ = v42
                                        _ = v43
                                    }
                                    v1 -= 2147483648
                                } else {
                                    switch fi28 {
                                    case 0:
                                        var v44 int16 = v20
                                        _ = v44
                                    default:
                                        if (2 < 2) {
                                            var v45 int32 = 537
                                            v14 = v39
                                            v39 -= v14
                                            var v46 int = v12
                                            var v47 float32 = (v14 * v39)
                                            v1 += 2147483648
                                            _ = v45
                                            _ = v46
                                            _ = v47
                                        }
                                    }
                                }
                                v48 := GenS5{F0: 9223372036854775807, F1: int8(v5)}
                                _ = v48
                            } else {
                                var v49 int32 = 42
                                var v50 string = "a"
                                _ = v49
                                _ = v50
                            }
                            v51 := GenS2{F0: 0}
                            _ = v51
                        }
                        _ = v39
                    }
                    _ = fi38
                }
            } else {
                var v52 string = ((("foo bar" + ("a" + "a")) + "foo bar") + "")
                fmt.Println(fi28)
                _ = v52
            }
        }
        switch v12 {
        case -1:
            fmt.Println(v12, v4, v14)
        case 0:
            v53 := v3[(v5 + 240)]
            _ = v53
        case 2147483647:
            delete(v3, 853)
        }
        for fi54 := 0; fi54 < 18; fi54++ {
            fmt.Println(fi54)
            switch "a" {
            case "foo bar":
                v20 -= v20
            }
            _ = fi54
        }
        _ = fi28
    }
    v2 -= (^(v2))
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
