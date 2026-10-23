package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32; F2 uint16 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint64; F1 int }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int32; F1 int16; F2 uint16 }
type GenS6 struct { F0 int16; F1 string; F2 int16 }
func genFn0() int16 {
    v6 := []int{232, 9223372036854775807, 128, -235}
    switch 32768 {
    case -9223372036854775808:
        var v7 int64 = 42
        _ = v7
    case 2147483647:
        v8 := GenS1{F0: ("a" + ("hello" + "hello")), F1: 1, F2: 255}
        _ = v8
    default:
        var v9 int8 = 100
        _ = v9
    }
    var v10 float64 = (0.1 / 0.1)
    fmt.Println(v10)
    _ = v10
    _ = v6
    return -742
}
func genFn1(p0 bool, p1 int, p2 int16) int8 {
    p0 = (p0 && (((((true || (0.1 < ((float64((0.1 / (1e10 - (3.14159 + ((((1.5 * (0.1 + (1e10 + 0.1))) / (3.14159 * 3.14159)) * 3.14159) + (0.0 * 1e10)))))) - 1e10) / -2.25))) && (255 > 42)) && p0) || (true || ((p2 / ((p2) | 1)) > (^(p2))))) && p0))
    for fi11 := 0; fi11 < 9; fi11++ {
        p1 = (-(p1))
        p0 = ((0 != 33) || p0)
        _ = fi11
    }
    p0 = ((0.1 >= 0.0) && true)
    _ = p0
    _ = p1
    _ = p2
    return -128
}

func main() {
    var v1 int16 = -93
    var v2 uint64 = 4294967295
    var v3 int8 = 98
    var v4 bool = (((-128 <= v3) || false) || false)
    v5 := []string{("" + "a"), ("\n\t" + (((("" + "\n\t") + ((("foo bar" + "") + "\n\t") + ("a" + "foo bar"))) + "\n\t") + "\n\t")), "foo bar", "a"}
    fmt.Println(v4)
    if v4 {
        var v12 uint = 42
        _ = len(v5)
        if (false || ((9223372036854775807 < 0) && (v12 > (v12 & (v12 | v12))))) {
            v13 := GenS2{F0: 2147483647}
            var v14 int8 = 127
            switch 32768 {
            case 127:
                v15 := GenS1{F0: "foo bar", F1: 469, F2: 65535}
                _ = v15
            case 152:
                v2 = (^(v2))
            case 0:
                fmt.Println(v1, v4, v3)
            default:
                fmt.Println(v3)
            }
            _ = v13
            _ = v14
        }
        _ = v12
    } else {
        if v4 {
            var v16 float32 = (0.1 + (1e10 / (3.14159 / 1e10)))
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            switch ("hello" + "foo bar") {
            case "foo bar":
                if len(v5) > 0 {
                    v5[0] = "hello"
                }
            case "hello":
                if v4 {
                    fmt.Println(v3, v16)
                    for fi17 := 0; fi17 < 13; fi17++ {
                        switch "\n\t" {
                        case "":
                            v16 -= v16
                        case "\n\t":
                            v18 := GenS2{F0: 0}
                            _ = v18
                        }
                        _ = fi17
                    }
                }
            default:
                v16 = v16
            }
            _ = v16
        } else {
            if ((^(v3)) > v3) {
                v19 := GenS3{F0: v2, F1: 32767}
                if (-1 != v3) {
                    for fi20 := 0; fi20 < 11; fi20++ {
                        if v4 {
                            if (0 >= 0) {
                                v2 = 4294967295
                                v1 = (v1 ^ (^(v1)))
                            } else {
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                                fmt.Println(v2)
                            }
                        } else {
                            v1 = (v1 * genFn0())
                            switch "foo bar" {
                            case "a":
                                for fi21 := 0; fi21 < 1; fi21++ {
                                    v2 -= v2
                                    _ = fi21
                                }
                            case "hello":
                                var v22 int64 = -9223372036854775808
                                _ = v22
                            case "":
                                _ = len(v5)
                            default:
                                _ = len(v5)
                            }
                            if v4 {
                                v23 := GenS2{F0: 42}
                                var v24 uint32 = 249
                                _ = v23
                                _ = v24
                            } else {
                                if (fi20 != 1) {
                                    v3 += -1
                                    var v25 float64 = 1.5
                                    _ = v25
                                }
                                v26 := GenS3{F0: (v2 ^ (v2 | v2)), F1: fi20}
                                v27 := GenS4{F0: 1}
                                _ = v26
                                _ = v27
                            }
                        }
                        _ = fi20
                    }
                } else {
                    v28 := GenS5{F0: 1, F1: genFn0(), F2: 255}
                    var v29 int16 = (v1 | (v1 * (v1 / (((^(v1))) | 1))))
                    v4 = v4
                    _ = v28
                    _ = v29
                }
                var v30 bool = false
                _ = v19
                _ = v30
            } else {
                var v31 int = int(v2)
                _ = v31
            }
            var v32 uint64 = (^(v2))
            v1 = -1
            _ = v32
        }
    }
    v4 = ((v4 && v4) || v4)
    v3 = v3
    var v33 uint = 65535
    if (256 > -187) {
        var v34 float64 = -2.25
        var v35 int16 = (v1 % ((v1) | 1))
        _ = v34
        _ = v35
    }
    if len(v5) > 0 {
        v5[0] = "a"
    }
    for fi36 := 0; fi36 < 16; fi36++ {
        if (!(v4 && v4)) {
            switch fi36 {
            case 9223372036854775807:
                _ = len(v5)
            }
            v1 = genFn0()
            _ = len(v5)
        } else {
            fmt.Println(v2, v2, v33)
            fmt.Println(v33, v3, v4)
            v37 := GenS6{F0: genFn0(), F1: "\n\t", F2: v1}
            _ = v37
        }
        _ = fi36
    }
    _ = len(v5)
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    switch "hello" {
    case "foo bar":
        v38 := GenS5{F0: -1, F1: (v1 >> ((v33) & 63)), F2: 42}
        _ = v38
    case "":
        v2 = v2
    default:
        fmt.Println(v4, v33, v33)
    }
    var v39 string = "\n\t"
    if ((v2 & (v2 | (v2 - v2))) > (v2 + (v2 >> ((v33) & 63)))) {
        v4 = true
    } else {
        fmt.Println(v3, v4, v39)
        v39 = (v39 + "")
        v40 := GenS2{F0: 929}
        _ = v40
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v39
    _ = v4
    _ = v5
}
