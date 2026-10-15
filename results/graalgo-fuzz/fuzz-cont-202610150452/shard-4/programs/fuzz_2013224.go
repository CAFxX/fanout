package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint32; F1 int16; F2 uint }
type GenS4 struct { F0 int32; F1 int64 }
type GenS5 struct { F0 uint32; F1 bool; F2 uint16 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 int8 }
func genFn0() uint8 {
    var v6 uint32 = 2147483648
    v7 := GenS1{F0: 9223372036854775807, F1: 2}
    v8 := GenS2{F0: -128}
    var v9 uint16 = 255
    var v10 uint = 4294967295
    _ = v10
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return 73
}
func genFn1(p0 int64, p1 uint) uint16 {
    p1 = (p1 ^ (p1 * 459))
    var v11 float32 = (-2.25 / -2.25)
    v12 := []int{}
    _ = p0
    _ = p1
    _ = v11
    _ = v12
    return 653
}

func main() {
    var v1 uint16 = 42
    var v2 int32 = 0
    v3 := []int{}
    var v4 uint = 321
    var v5 uint16 = (^(v1))
    if false {
        switch 1 {
        case -9223372036854775808:
            v2 = v2
        case 502:
            var v13 uint16 = (v1 * (v1 / ((v5) | 1)))
            _ = v13
        default:
            if false {
                fmt.Println(v4, v4)
                var v14 int16 = 1
                if true {
                    var v15 int = 130
                    _ = v15
                }
                _ = v14
            }
        }
        for fi16 := 0; fi16 < 10; fi16++ {
            if (1000 > -32768) {
                for fi17 := 0; fi17 < 4; fi17++ {
                    v5 = (v5 & 648)
                    _ = fi17
                }
                v5 = genFn1(0, v4)
                if true {
                    for fi18 := 0; fi18 < 13; fi18++ {
                        if true {
                            var v19 float64 = (3.14159 - 0.1)
                            var v20 int64 = -9223372036854775808
                            v5 -= uint16(v4)
                            if ((^(v2)) == 0) {
                                v21 := GenS1{F0: 1, F1: v5}
                                v2 = 2147483647
                                _ = v21
                            }
                            _ = v19
                            _ = v20
                        }
                        _ = fi18
                    }
                }
            } else {
                v3 = append(v3, 32768)
                switch -2147483648 {
                case 32768:
                    fmt.Println(v4, v5)
                default:
                    fmt.Println(v1, v2)
                }
                v22 := GenS3{F0: 65535, F1: int16(v5), F2: 767}
                _ = v22
            }
            v5 += 65535
            fmt.Println(v5)
            _ = fi16
        }
    }
    var v23 uint = (v4 & 42)
    v24 := GenS4{F0: 1, F1: 1}
    if (genFn1(9223372036854775807, 415) < v5) {
        v2 += (^(v2))
        for fi25 := 0; fi25 < 6; fi25++ {
            if (466 <= 29) {
                var v26 int = (fi25 >> ((v4) & 63))
                _ = v26
            } else {
                v27 := GenS5{F0: 980, F1: true, F2: (v1 << (((^(v4))) & 63))}
                _ = len(v3)
                v4 = 65535
                _ = v27
            }
            if (!(128 >= 42)) {
                switch "\n\t" {
                case "hello":
                    if (("hello" + "foo bar") > ("" + "hello")) {
                        v2 -= v2
                        var v28 int = fi25
                        if true {
                            if (32767 != -32768) {
                                v23 = uint(v5)
                                for fi29 := 0; fi29 < 6; fi29++ {
                                    _ = len(v3)
                                    _ = fi29
                                }
                            } else {
                                fmt.Println(fi25, v1, v5)
                                fmt.Println(v23, v2, v1)
                            }
                        } else {
                            v30 := GenS6{F0: 1}
                            _ = v30
                        }
                        _ = v28
                    }
                }
                if (0.0 >= float64(0.1)) {
                    var v31 int64 = 9223372036854775807
                    if (!((!true) || (!false))) {
                        var v32 float32 = 1.5
                        v5 = (v1 / (((v1 / ((genFn1(v31, (v4 - v4))) | 1))) | 1))
                        _ = v32
                    }
                    var v33 int8 = 1
                    _ = v31
                    _ = v33
                } else {
                    if true {
                        var v34 int = (fi25 | 127)
                        var v35 uint = v4
                        switch ("foo bar" + (((("a" + ("foo bar" + (("a" + ("a" + "foo bar")) + ""))) + "hello") + "") + "foo bar")) {
                        case "":
                            if (128 <= 42) {
                                if true {
                                    v34 += (^(v34))
                                } else {
                                    v4 -= v4
                                }
                            }
                        }
                        _ = v34
                        _ = v35
                    } else {
                        var v36 int16 = 0
                        var v37 int8 = 1
                        for fi38 := 0; fi38 < 2; fi38++ {
                            v1 -= genFn1(-349, v4)
                            _ = fi38
                        }
                        _ = v36
                        _ = v37
                    }
                }
            }
            _ = fi25
        }
        v39 := GenS5{F0: 65535, F1: false, F2: 2}
        _ = v39
    } else {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        fmt.Println(v2, v1, v23)
    }
    switch "foo bar" {
    case "a":
        var v40 int16 = -619
        _ = v40
    case "":
        v23 = v23
    case "foo bar":
        var v41 int64 = -200
        _ = v41
    }
    var v42 uint16 = (^(v5))
    v43 := GenS6{F0: 9223372036854775807}
    for fi44 := 0; fi44 < 14; fi44++ {
        switch "foo bar" {
        case "a":
            var v45 int64 = 9223372036854775807
            _ = v45
        case "hello":
            fmt.Println(fi44)
        }
        v46 := GenS7{F0: 0}
        _ = fi44
        _ = v46
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v5
}
