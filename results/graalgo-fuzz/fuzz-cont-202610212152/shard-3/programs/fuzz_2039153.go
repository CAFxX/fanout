package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool; F2 int16 }
type GenS2 struct { F0 uint8; F1 uint8; F2 int8 }
type GenS3 struct { F0 int64; F1 uint8; F2 uint16 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 int64; F1 int32; F2 uint }
type GenS7 struct { F0 uint; F1 uint16 }
type GenS8 struct { F0 uint16; F1 int64 }
func genFn0() int64 {
    v6 := GenS1{F0: -2147483648, F1: false, F2: 1000}
    var v7 int16 = 599
    v8 := []int{-2147483648, 128, 127, 9223372036854775807}
    if true {
        v7 = 42
        var v9 string = ("hello" + ((("" + "\n\t") + "") + ("a" + ("a" + (("hello" + "") + "a")))))
        _ = v9
    }
    switch "hello" {
    case "":
        v7 -= 450
    case "foo bar":
        fmt.Println(v7)
    case "a":
        fmt.Println(v7)
    default:
        v10 := GenS1{F0: 42, F1: (!false), F2: v7}
        _ = v10
    }
    _ = v6
    _ = v7
    _ = v8
    return 42
}
func genFn1(p0 uint, p1 int, p2 uint32) uint8 {
    for fi11 := 0; fi11 < 5; fi11++ {
        var v12 float64 = (0.0 - -2.25)
        _ = fi11
        _ = v12
    }
    p2 += p2
    fmt.Println(p0, p1)
    p2 = p2
    v13 := []int{p1, (p1 >> (((p0 % ((p0) | 1))) & 63))}
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    return 127
}

func main() {
    var v1 uint8 = 42
    var v2 int = 1
    var v3 uint = 65535
    v4 := []uint32{42, 66, 65535, 25}
    var v5 bool = true
    v14 := GenS1{F0: -1, F1: v5, F2: 1}
    v4 = append(v4, 2147483648)
    v15 := GenS1{F0: -1, F1: v5, F2: 32767}
    if v5 {
        switch "\n\t" {
        case "foo bar":
            if ("foo bar" <= ("\n\t" + "a")) {
                v16 := GenS2{F0: (v1 ^ (v1 >> (((^(v3))) & 63))), F1: (v1 / ((genFn1(4294967295, (v2 / ((v2) | 1)), 316)) | 1)), F2: 100}
                for fi17 := 0; fi17 < 6; fi17++ {
                    for fi18 := 0; fi18 < 16; fi18++ {
                        v19 := GenS3{F0: -9223372036854775808, F1: (^(v1)), F2: 42}
                        _ = fi18
                        _ = v19
                    }
                    if v5 {
                        var v20 uint8 = genFn1(v3, v2, 472)
                        if v5 {
                            var v21 uint8 = genFn1((v3 + (v3 / (((^(v3))) | 1))), -2147483648, 310)
                            _ = v21
                        } else {
                            var v22 uint32 = 137
                            var v23 float32 = (3.14159 * 1e10)
                            _ = v22
                            _ = v23
                        }
                        _ = v20
                    } else {
                        v3 = 307
                    }
                    _ = fi17
                }
                switch "\n\t" {
                case "foo bar":
                    v2 = v2
                case "":
                    var v24 int8 = -128
                    _ = v24
                case "a":
                    v25 := GenS1{F0: 0, F1: (v5 && v5), F2: 888}
                    _ = v25
                default:
                    var v26 float32 = -2.25
                    _ = v26
                }
                _ = v16
            }
        case "":
            if false {
                switch "foo bar" {
                case "a":
                    if v5 {
                        if (1 <= -32768) {
                            var v27 string = "a"
                            v3 = 0
                            _ = v27
                        }
                        v2 -= v2
                    }
                }
            }
        case "\n\t":
            for fi28 := 0; fi28 < 8; fi28++ {
                switch (v2 - v2) {
                case 4294967295:
                    fmt.Println(v3)
                case 128:
                    fmt.Println(fi28)
                case 127:
                    switch "\n\t" {
                    case "hello":
                        for fi29 := 0; fi29 < 15; fi29++ {
                            var v30 float32 = (1.5 + (-2.25 + ((0.0 * -2.25) + 3.14159)))
                            _ = len(v4)
                            _ = fi29
                            _ = v30
                        }
                    case "":
                        fmt.Println(v2, fi28)
                    }
                }
                v5 = true
                switch (("foo bar" + ("hello" + "a")) + ("foo bar" + "foo bar")) {
                case "":
                    v3 = (v3 >> ((v3) & 63))
                case "hello":
                    var v31 uint = 577
                    _ = v31
                }
                _ = fi28
            }
        default:
            var v32 string = "\n\t"
            _ = v32
        }
        var v33 string = "\n\t"
        fmt.Println(v1)
        _ = v33
    } else {
        _ = len(v4)
        if (v3 != (^(v3))) {
            if len(v4) > 0 {
                v4[0] = 792
            }
        } else {
            v34 := GenS1{F0: -2147483648, F1: (395 >= v2), F2: 802}
            var v35 bool = (v5 && (-36 <= 1))
            if len(v4) > 0 {
                v4[0] = 1
            }
            _ = v34
            _ = v35
        }
    }
    fmt.Println(v2, v2, v5)
    v1 += v1
    if ((127 > -82) || (-109 == 1)) {
        switch (("" + ("" + (("\n\t" + "") + "hello"))) + "hello") {
        case "\n\t":
            if (32767 <= 32767) {
                v4 = append(v4, 390)
            } else {
                for fi36 := 0; fi36 < 11; fi36++ {
                    if v5 {
                        v4 = append(v4, 2)
                        var v37 int16 = 1000
                        var v38 int16 = v37
                        _ = v37
                        _ = v38
                    } else {
                        fmt.Println(fi36)
                        if len(v4) > 0 {
                            v4[0] = 4294967295
                        }
                        switch (v2 / ((-2147483648) | 1)) {
                        case 42:
                            for fi39 := 0; fi39 < 11; fi39++ {
                                var v40 float32 = float32(((0.1 + 0.1) + (3.14159 - (((1e10 - (0.0 / float64(0.1))) - ((3.14159 / 1.0) + -2.25)) / 1.0))))
                                v41 := GenS2{F0: (v1 % (((v1 * v1)) | 1)), F1: 127, F2: -31}
                                v42 := GenS4{F0: -91}
                                v5 = v5
                                if true {
                                    v43 := GenS4{F0: 66}
                                    var v44 int = 91
                                    v45 := GenS1{F0: 417, F1: v5, F2: 421}
                                    _ = v43
                                    _ = v44
                                    _ = v45
                                }
                                v1 -= uint8(v3)
                                _ = fi39
                                _ = v40
                                _ = v41
                                _ = v42
                            }
                        }
                    }
                    v4 = append(v4, 2147483647)
                    v46 := GenS4{F0: int8(v3)}
                    _ = fi36
                    _ = v46
                }
                v47 := GenS1{F0: -1, F1: false, F2: -32768}
                switch "" {
                case "\n\t":
                    v48 := GenS5{F0: 65535}
                    _ = v48
                case "":
                    var v49 uint64 = 886
                    _ = v49
                case "hello":
                    fmt.Println(v1, v3, v5)
                default:
                    v1 = 255
                }
                _ = v47
            }
        case "hello":
            _ = len(v4)
        case "":
            v50 := GenS6{F0: 42, F1: 912, F2: uint(v1)}
            _ = v50
        default:
            v51 := GenS7{F0: v3, F1: 2}
            _ = v51
        }
        if (!true) {
            var v52 uint8 = (v1 / (((v1 / (((v1 + 37)) | 1))) | 1))
            if (-1 == 77) {
                if (65535 < 2) {
                    var v53 int = (v2 % (((v2 << ((2) & 63))) | 1))
                    v54 := GenS7{F0: v3, F1: 195}
                    for fi55 := 0; fi55 < 9; fi55++ {
                        for fi56 := 0; fi56 < 10; fi56++ {
                            v1 -= (^(v52))
                            _ = fi56
                        }
                        v57 := GenS8{F0: 42, F1: 0}
                        _ = fi55
                        _ = v57
                    }
                    _ = v53
                    _ = v54
                }
                v5 = ((v2 == v2) && true)
            }
            for fi58 := 0; fi58 < 19; fi58++ {
                var v59 uint16 = 1
                switch (^(fi58)) {
                case -1:
                    v5 = v5
                case 42:
                    var v60 int32 = -1
                    _ = v60
                }
                for fi61 := 0; fi61 < 19; fi61++ {
                    fmt.Println(fi61, v1, v2)
                    _ = fi61
                }
                _ = fi58
                _ = v59
            }
            _ = v52
        } else {
            v62 := GenS1{F0: 2147483647, F1: (v5 && (271 < uint16(v3))), F2: -242}
            for fi63 := 0; fi63 < 7; fi63++ {
                fmt.Println(v3, v1)
                fmt.Println(v5)
                _ = fi63
            }
            _ = v62
        }
    }
    v3 += v3
    var v64 uint8 = (^(v1))
    var v65 uint64 = 72
    var v66 uint = (^(v3))
    fmt.Println(v3, v65, v5)
    var v67 uint8 = 90
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v64:", v64)
    fmt.Println("v65:", v65)
    fmt.Println("v66:", v66)
    fmt.Println("v67:", v67)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v64
    _ = v65
    _ = v66
    _ = v67
}
