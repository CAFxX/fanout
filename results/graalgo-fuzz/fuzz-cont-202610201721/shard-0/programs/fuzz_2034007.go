package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string; F2 int32 }
type GenS2 struct { F0 uint64; F1 uint; F2 uint8 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint8; F1 string; F2 uint64 }
type GenS5 struct { F0 uint32; F1 int16 }
type GenS6 struct { F0 int; F1 uint32 }
type GenS7 struct { F0 bool; F1 uint64 }
type GenS8 struct { F0 uint16; F1 uint8; F2 uint64 }
type GenS9 struct { F0 uint8 }
type GenS10 struct { F0 int64 }
type GenS11 struct { F0 int16 }
type GenS12 struct { F0 uint16; F1 uint8; F2 int32 }
func genFn0(p0 uint64, p1 uint64, p2 int32) bool {
    p1 = p1
    v4 := []int{1}
    for fi5 := 0; fi5 < 19; fi5++ {
        fmt.Println(p2, fi5, p0)
        _ = fi5
    }
    for fi6 := 0; fi6 < 12; fi6++ {
        p0 = (p0 / ((p0) | 1))
        var v7 int64 = -473
        _ = fi6
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return (2 <= 128)
}
func genFn1(p0 uint) uint16 {
    p0 -= p0
    for fi8 := 0; fi8 < 7; fi8++ {
        v9 := GenS1{F0: -128, F1: ("" + ("" + "a")), F2: 42}
        switch "" {
        case "foo bar":
            if true {
                if (244 >= 18446744073709551615) {
                    p0 -= (p0 % ((898) | 1))
                    var v10 int32 = -574
                    _ = v10
                }
                if (284 == 2) {
                    var v11 int32 = 232
                    v12 := []int{fi8, fi8, (-(fi8)), (fi8 ^ (fi8 | 42))}
                    _ = v11
                    _ = v12
                } else {
                    var v13 uint16 = 213
                    v13 -= v13
                    v14 := []int{(-(fi8)), (fi8 | fi8), -317, (fi8 + (-(fi8)))}
                    _ = v13
                    _ = v14
                }
            }
        case "\n\t":
            p0 = p0
        default:
            var v15 string = ("" + "")
            _ = v15
        }
        for fi16 := 0; fi16 < 19; fi16++ {
            p0 = (p0 >> ((239) & 63))
            var v17 uint = p0
            if (!false) {
                if true {
                    if true {
                        v18 := []int{(fi8 & fi8), fi16}
                        switch "a" {
                        case "\n\t":
                            v19 := GenS2{F0: 18446744073709551615, F1: v17, F2: 76}
                            _ = v19
                        case "a":
                            switch fi8 {
                            case 127:
                                switch (("\n\t" + "a") + "") {
                                case "foo bar":
                                    if len(v18) > 0 {
                                        v18[0] = -108
                                    }
                                    v17 += p0
                                default:
                                    if len(v18) > 0 {
                                        v18[0] = (fi16 >> ((p0) & 63))
                                    }
                                }
                            case -438:
                                var v20 int32 = 42
                                v20 = (v20 << ((p0) & 63))
                                v18 = append(v18, (fi16 * fi8))
                                _ = v20
                            }
                        default:
                            p0 += (^(v17))
                        }
                        _ = v18
                    } else {
                        v21 := []int{(-(fi8)), fi8}
                        _ = v21
                    }
                    v22 := GenS3{F0: false}
                    v23 := []int{-2147483648, (-(fi8)), (^(fi16)), (^(fi8))}
                    _ = v22
                    _ = v23
                }
                var v24 float64 = 3.14159
                _ = v24
            }
            _ = fi16
            _ = v17
        }
        _ = fi8
        _ = v9
    }
    _ = p0
    return 722
}

func main() {
    v1 := []int{-963, -2147483648, 609, -1}
    var v2 uint8 = 127
    v3 := make(map[string]uint32)
    v1 = append(v1, 0)
    switch 2147483647 {
    case -2147483648:
        switch 0 {
        case 343:
            for fi25 := 0; fi25 < 20; fi25++ {
                fmt.Println(fi25, fi25)
                _ = fi25
            }
        case 492:
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        case 128:
            if (!true) {
                fmt.Println(v2)
                fmt.Println(v2)
            } else {
                var v26 uint = 1
                _ = v26
            }
        default:
            if len(v1) > 0 {
                v1[0] = 38
            }
        }
    case -195:
        var v27 uint16 = genFn1(255)
        _ = v27
    case 385:
        if (-2147483648 > 2147483647) {
            var v28 uint32 = 712
            _ = v28
        }
    default:
        if (-666 <= 96) {
            v29 := GenS1{F0: 3, F1: "a", F2: 0}
            _ = v29
        } else {
            v30 := GenS3{F0: true}
            var v31 int64 = -810
            _ = v30
            _ = v31
        }
    }
    switch 2147483647 {
    case 9223372036854775807:
        if false {
            switch -448 {
            case 32767:
                switch "foo bar" {
                case "\n\t":
                    v32 := GenS3{F0: (((v2 > 175) || true) || false)}
                    _ = v32
                }
            default:
                v33 := GenS4{F0: (v2 & v2), F1: "\n\t", F2: 18446744073709551615}
                _ = v33
            }
            fmt.Println(v2)
        } else {
            for fi34 := 0; fi34 < 4; fi34++ {
                if (1.5 == (0.1 / 3.14159)) {
                    switch "a" {
                    case "a":
                        _ = len(v1)
                    case "foo bar":
                        v35 := GenS5{F0: 42, F1: 623}
                        _ = v35
                    }
                    v36 := v3["a"]
                    switch "\n\t" {
                    case "foo bar":
                        if (v2 > 255) {
                            var v37 float32 = 3.14159
                            _ = v37
                        } else {
                            v38 := GenS2{F0: uint64(v2), F1: 832, F2: (v2 ^ v2)}
                            v39 := GenS6{F0: (fi34 * (^(fi34))), F1: (^(v36))}
                            fmt.Println(fi34, fi34)
                            _ = v38
                            _ = v39
                        }
                    case "":
                        var v40 int8 = 39
                        _ = v40
                    default:
                        fmt.Println(fi34)
                    }
                    _ = v36
                } else {
                    var v41 uint32 = 2
                    var v42 int = fi34
                    v43 := GenS1{F0: 127, F1: ("foo bar" + "\n\t"), F2: -1}
                    _ = v41
                    _ = v42
                    _ = v43
                }
                _ = fi34
            }
            var v44 float64 = 1e10
            if true {
                var v45 float32 = 0.1
                if (!(false && ((1000 == 1) && (972 != 626)))) {
                    var v46 uint32 = 4294967295
                    v2 = (v2 - 255)
                    v2 += (v2 ^ (^(v2)))
                    _ = v46
                } else {
                    v44 = (float64(v45) * v44)
                }
                v47 := GenS7{F0: ((true && false) && false), F1: 4294967295}
                _ = v45
                _ = v47
            } else {
                v2 = (v2 & v2)
                v48 := GenS8{F0: 2, F1: (v2 << ((194) & 63)), F2: 18}
                _ = v48
            }
            _ = v44
        }
    case -703:
        fmt.Println(v2)
    default:
        switch 0 {
        case 2147483647:
            if ((!false) || ((207 > 18446744073709551615) && (("" + "\n\t") != "hello"))) {
                if true {
                    var v49 uint32 = 2
                    if (false && true) {
                        _ = len(v3)
                    } else {
                        switch "a" {
                        case "hello":
                            v50 := GenS8{F0: genFn1(255), F1: v2, F2: 4294967295}
                            _ = v50
                        case "foo bar":
                            v49 = v49
                        default:
                            if false {
                                fmt.Println(v2, v49)
                                fmt.Println(v2)
                            }
                        }
                        var v51 bool = (128 != -9223372036854775808)
                        _ = v51
                    }
                    _ = v49
                } else {
                    switch "" {
                    case "":
                        if (!("\n\t" < ("foo bar" + "foo bar"))) {
                            var v52 uint16 = genFn1(804)
                            switch "hello" {
                            case "":
                                for fi53 := 0; fi53 < 20; fi53++ {
                                    v2 -= (^(v2))
                                    v54 := GenS9{F0: (^(v2))}
                                    _ = fi53
                                    _ = v54
                                }
                            case "hello":
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                            default:
                                switch 37 {
                                case 115:
                                    for fi55 := 0; fi55 < 12; fi55++ {
                                        for fi56 := 0; fi56 < 6; fi56++ {
                                            switch "a" {
                                            case "hello":
                                                v57 := GenS6{F0: (fi55 / ((fi56) | 1)), F1: 0}
                                                _ = v57
                                            case "\n\t":
                                                v2 = (v2 ^ 44)
                                            }
                                            fmt.Println(v52, fi56, fi55)
                                            if true {
                                                v2 = (v2 - (v2 + v2))
                                            }
                                            _ = fi56
                                        }
                                        switch "foo bar" {
                                        case "":
                                            _ = len(v1)
                                        case "hello":
                                            var v58 uint = 255
                                            _ = v58
                                        }
                                        fmt.Println(fi55)
                                        _ = fi55
                                    }
                                default:
                                    var v59 float32 = -2.25
                                    _ = v59
                                }
                            }
                            _ = v52
                        }
                    }
                }
            } else {
                for fi60 := 0; fi60 < 19; fi60++ {
                    for fi61 := 0; fi61 < 6; fi61++ {
                        v2 = (v2 >> ((2) & 63))
                        v62 := GenS10{F0: -874}
                        v2 -= (v2 ^ (v2 * (v2 + (v2 | v2))))
                        var v63 uint64 = 2
                        v3["foo bar"] = 0
                        _ = fi61
                        _ = v62
                        _ = v63
                    }
                    v2 -= v2
                    _ = fi60
                }
                switch (("\n\t" + "") + "") {
                case "":
                    var v64 uint16 = 42
                    _ = v64
                case "hello":
                    v65 := GenS1{F0: 100, F1: ("" + "hello"), F2: 497}
                    _ = v65
                case "foo bar":
                    v2 -= 128
                default:
                    for fi66 := 0; fi66 < 8; fi66++ {
                        var v67 uint16 = 9
                        if (((!false) || (true || ((4294967295 < 2147483647) && (((false || true) && (((uint32(v2) < 42) && ((!(!((1.5 != 1.5) && true))) && (!(!((((float32(-2.25) + 0.1) * 1.5) / 1.0) >= 0.0))))) && true)) && (((-2.25 / (1e10 / 1.0)) / 0.1) != 1e10))))) || false) {
                            v68 := GenS6{F0: (fi66 | 128), F1: 65535}
                            _ = v68
                        }
                        if true {
                            v2 -= v2
                            var v69 bool = false
                            var v70 uint64 = 2
                            _ = v69
                            _ = v70
                        }
                        _ = fi66
                        _ = v67
                    }
                }
                v2 -= v2
            }
        case 42:
            for fi71 := 0; fi71 < 12; fi71++ {
                var v72 int = (^(fi71))
                for fi73 := 0; fi73 < 9; fi73++ {
                    var v74 uint64 = 1
                    _ = fi73
                    _ = v74
                }
                for fi75 := 0; fi75 < 19; fi75++ {
                    _ = len(v3)
                    delete(v3, "")
                    _ = fi75
                }
                _ = fi71
                _ = v72
            }
        case 32768:
            var v76 uint32 = 4294967295
            _ = v76
        }
    }
    var v77 uint16 = 2
    switch "" {
    case "foo bar":
        var v78 uint8 = v2
        _ = v78
    case "hello":
        if true {
            var v79 float64 = ((-2.25 + 0.1) / float64(-2.25))
            v2 -= v2
            _ = v79
        } else {
            v80 := GenS4{F0: v2, F1: ("\n\t" + "a"), F2: 4294967295}
            var v81 uint32 = 520
            _ = v80
            _ = v81
        }
    default:
        var v82 uint = 4294967295
        _ = v82
    }
    var v83 string = "a"
    fmt.Println(v83, v77)
    v84 := GenS11{F0: -91}
    fmt.Println(v83, v77, v2)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v85 int16 = 678
    v86 := GenS9{F0: v2}
    for fi87 := 0; fi87 < 17; fi87++ {
        v85 = v85
        _ = fi87
    }
    var v88 uint32 = 65535
    var v89 int8 = 127
    v90 := GenS6{F0: -601, F1: (v88 << ((42) & 63))}
    for fi91 := 0; fi91 < 2; fi91++ {
        delete(v3, ("" + ""))
        v92 := GenS12{F0: (v77 >> ((0) & 63)), F1: (v2 & uint8(v88)), F2: 2147483647}
        _ = fi91
        _ = v92
    }
    var v93 uint32 = v88
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v77:", v77)
    fmt.Println("v83:", v83)
    fmt.Println("v85:", v85)
    fmt.Println("v88:", v88)
    fmt.Println("v89:", v89)
    fmt.Println("v93:", v93)
    _ = v1
    _ = v2
    _ = v3
    _ = v77
    _ = v83
    _ = v84
    _ = v85
    _ = v86
    _ = v88
    _ = v89
    _ = v90
    _ = v93
}
