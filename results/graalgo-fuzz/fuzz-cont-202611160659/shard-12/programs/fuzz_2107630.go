package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint; F1 uint64 }
type GenS3 struct { F0 bool; F1 int }
type GenS4 struct { F0 int16; F1 int64; F2 string }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 uint; F1 uint32; F2 uint64 }
type GenS7 struct { F0 bool; F1 int8; F2 uint16 }
type GenS8 struct { F0 string; F1 uint32 }
type GenS9 struct { F0 uint64 }
func genFn0(p0 int32) bool {
    v4 := []int{}
    var v5 bool = ("\n\t" > (("" + "hello") + "\n\t"))
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v6 uint = 469
    for fi7 := 0; fi7 < 16; fi7++ {
        v5 = false
        _ = fi7
    }
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    return v5
}
func genFn1() int32 {
    var v8 int16 = 32767
    v8 -= v8
    v9 := GenS1{F0: 2147483647}
    switch 9223372036854775807 {
    case -9223372036854775808:
        switch 2147483647 {
        case 4294967295:
            v10 := []int{128}
            _ = v10
        case 128:
            var v11 int8 = 15
            _ = v11
        case 127:
            switch "a" {
            case "":
                if false {
                    var v12 uint = 819
                    for fi13 := 0; fi13 < 4; fi13++ {
                        v8 = (-(v8))
                        switch "" {
                        case "\n\t":
                            for fi14 := 0; fi14 < 6; fi14++ {
                                if true {
                                    fmt.Println(fi13, v8, fi14)
                                }
                                v8 -= v8
                                _ = fi14
                            }
                        case "a":
                            v12 = (v12 ^ (v12 << ((v12) & 63)))
                        case "hello":
                            fmt.Println(fi13)
                        }
                        v12 = v12
                        _ = fi13
                    }
                    if (!true) {
                        switch "hello" {
                        case "a":
                            if (0.0 == (1.5 / 1.5)) {
                                v12 += v12
                                var v15 uint8 = 0
                                _ = v15
                            } else {
                                var v16 int8 = -1
                                _ = v16
                            }
                        case "foo bar":
                            v17 := []int{2147483647, 1, -910, 538}
                            _ = v17
                        case "\n\t":
                            v18 := []int{-9223372036854775808}
                            _ = v18
                        }
                        fmt.Println(v8, v12)
                    }
                    _ = v12
                } else {
                    v19 := GenS2{F0: 42, F1: 880}
                    if false {
                        v8 += v8
                        v20 := []int{32768, 2147483647, 128}
                        v8 -= (v8 % (((v8 - 42)) | 1))
                        _ = v20
                    }
                    switch 42 {
                    case 1:
                        fmt.Println(v8)
                    case 128:
                        fmt.Println(v8)
                    default:
                        var v21 int = 902
                        _ = v21
                    }
                    _ = v19
                }
            case "\n\t":
                v22 := []int{-214, 32768, 128}
                _ = v22
            case "foo bar":
                switch "\n\t" {
                case "a":
                    switch "hello" {
                    case "a":
                        v8 -= -976
                    case "hello":
                        switch "a" {
                        case "a":
                            fmt.Println(v8)
                        case "hello":
                            v23 := []int{-217, -945}
                            _ = v23
                        case "\n\t":
                            for fi24 := 0; fi24 < 7; fi24++ {
                                switch fi24 {
                                case 32767:
                                    var v25 int16 = (v8 + (^(v8)))
                                    _ = v25
                                default:
                                    var v26 string = "\n\t"
                                    _ = v26
                                }
                                var v27 int64 = -682
                                fmt.Println(fi24)
                                _ = fi24
                                _ = v27
                            }
                        }
                    case "":
                        fmt.Println(v8)
                    }
                case "hello":
                    v28 := GenS3{F0: true, F1: 9223372036854775807}
                    _ = v28
                case "":
                    for fi29 := 0; fi29 < 10; fi29++ {
                        var v30 int = fi29
                        v31 := GenS2{F0: 65535, F1: 1}
                        if false {
                            v8 -= 1
                        } else {
                            v30 = fi29
                        }
                        _ = fi29
                        _ = v30
                        _ = v31
                    }
                }
            }
        default:
            v8 -= v8
        }
    case 32768:
        var v32 string = "foo bar"
        _ = v32
    case 342:
        switch 106 {
        case 127:
            v8 += -1
        case -9223372036854775808:
            v33 := GenS2{F0: 255, F1: 18446744073709551615}
            _ = v33
        case 2147483647:
            v8 = -1
        default:
            for fi34 := 0; fi34 < 13; fi34++ {
                fmt.Println(v8, fi34)
                var v35 uint64 = 2
                v8 = v8
                _ = fi34
                _ = v35
            }
        }
    default:
        v8 = (v8 / ((v8) | 1))
    }
    _ = v8
    _ = v9
    return -71
}

func main() {
    var v1 uint32 = 42
    var v2 uint = 378
    var v3 int16 = 431
    if (!false) {
        var v36 int = -2147483648
        v37 := GenS3{F0: ((-(v36)) > (v36 | -999)), F1: (v36 % (((v36 * (v36 * 4294967295))) | 1))}
        _ = v36
        _ = v37
    } else {
        v38 := []int{4294967295, -9223372036854775808, 127, 0}
        v1 = v1
        _ = v38
    }
    for fi39 := 0; fi39 < 14; fi39++ {
        if true {
            switch "\n\t" {
            case "foo bar":
                var v40 bool = ((^(v3)) > v3)
                _ = v40
            case "":
                v1 += v1
            }
            v3 -= (v3 >> (((v2 ^ v2)) & 63))
            v3 = v3
        } else {
            fmt.Println(v2, fi39, v3)
            var v41 int64 = 42
            for fi42 := 0; fi42 < 14; fi42++ {
                fmt.Println(v2, fi39, v41)
                v43 := GenS4{F0: (^(v3)), F1: 1, F2: ("\n\t" + ("a" + "\n\t"))}
                v44 := GenS5{F0: 42}
                _ = fi42
                _ = v43
                _ = v44
            }
            _ = v41
        }
        if true {
            v1 = (^(v1))
            var v45 float64 = (1e10 / ((((0.1 / (-2.25 + 0.1)) + (1e10 / -2.25)) * (-2.25 + 1.5)) - 0.1))
            var v46 int = fi39
            _ = v45
            _ = v46
        } else {
            v47 := GenS6{F0: v2, F1: 1, F2: 42}
            _ = v47
        }
        v48 := GenS6{F0: uint(v1), F1: v1, F2: 4294967295}
        _ = fi39
        _ = v48
    }
    var v49 int8 = -89
    if true {
        v3 += (^(v3))
        var v50 float32 = float32(-2.25)
        fmt.Println(v50)
        _ = v50
    } else {
        for fi51 := 0; fi51 < 5; fi51++ {
            v2 = v2
            _ = fi51
        }
        fmt.Println(v1, v3, v1)
    }
    v52 := GenS7{F0: false, F1: (^(v49)), F2: 65535}
    v53 := []int{-674, 688, -602, -985}
    for fi54 := 0; fi54 < 7; fi54++ {
        _ = len(v53)
        if true {
            v55 := GenS2{F0: v2, F1: 311}
            _ = v55
        } else {
            var v56 uint = (v2 ^ (v2 - 18446744073709551615))
            _ = v56
        }
        _ = fi54
    }
    _ = len(v53)
    for fi57 := 0; fi57 < 20; fi57++ {
        v1 += v1
        fmt.Println(v1, v49)
        _ = fi57
    }
    v2 = (^(v2))
    for fi58 := 0; fi58 < 17; fi58++ {
        v2 = 62
        _ = fi58
    }
    fmt.Println(v1, v2, v2)
    for fi59 := 0; fi59 < 4; fi59++ {
        switch "a" {
        case "":
            var v60 uint32 = 0
            _ = v60
        case "foo bar":
            v49 = 53
        case "\n\t":
            if ((("a" + "") + ("\n\t" + "hello")) >= "") {
                if true {
                    v61 := GenS8{F0: ("hello" + "foo bar"), F1: (v1 | 4294967295)}
                    v1 += 755
                    _ = v61
                }
            } else {
                v62 := GenS2{F0: v2, F1: 1}
                _ = v62
            }
        }
        v49 = (v49 ^ 127)
        if len(v53) > 0 {
            v53[0] = fi59
        }
        _ = fi59
    }
    if (!(!("a" >= "hello"))) {
        var v63 int8 = (v49 * (v49 & -4))
        v64 := GenS9{F0: 1}
        _ = v63
        _ = v64
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v49:", v49)
    fmt.Println("slicelen v53:", len(v53), v53)
    _ = v1
    _ = v2
    _ = v3
    _ = v49
    _ = v52
    _ = v53
}
