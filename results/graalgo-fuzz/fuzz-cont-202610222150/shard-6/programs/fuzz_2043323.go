package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16 }
type GenS2 struct { F0 bool; F1 int32 }
type GenS3 struct { F0 uint64; F1 uint16 }
type GenS4 struct { F0 uint16; F1 int8; F2 int }
type GenS5 struct { F0 uint32; F1 string; F2 int32 }
type GenS6 struct { F0 bool; F1 uint8 }
type GenS7 struct { F0 int8; F1 uint64 }
type GenS8 struct { F0 uint64 }

func main() {
    v1 := []int{128, 4294967295, -9223372036854775808, 0}
    v2 := []int{-1, -824, 32768}
    v3 := []string{("\n\t" + ("hello" + ("" + "\n\t"))), "foo bar", (("" + ("" + ("hello" + "a"))) + "a"), "a"}
    var v4 int16 = 1000
    if true {
        for fi5 := 0; fi5 < 14; fi5++ {
            fmt.Println(v4, v4)
            if ((-1 > 28) || true) {
                v4 = (v4 * (^(v4)))
            }
            var v6 int64 = 127
            _ = fi5
            _ = v6
        }
        v4 += (v4 / (((v4 * (v4 & (v4 & v4)))) | 1))
    } else {
        v4 = v4
        _ = len(v3)
    }
    v4 = (v4 % (((v4 / ((v4) | 1))) | 1))
    v4 -= -273
    if (!true) {
        if false {
            fmt.Println(v4)
            switch "" {
            case "":
                v7 := GenS1{F0: 2, F1: 670}
                _ = v7
            case "a":
                for fi8 := 0; fi8 < 12; fi8++ {
                    v4 += (^(v4))
                    _ = fi8
                }
            default:
                fmt.Println(v4)
            }
        } else {
            if true {
                for fi9 := 0; fi9 < 10; fi9++ {
                    var v10 string = ("\n\t" + "\n\t")
                    _ = fi9
                    _ = v10
                }
                switch "a" {
                case "\n\t":
                    _ = len(v1)
                }
                v4 += (v4 >> ((745) & 63))
            }
            var v11 int64 = -872
            _ = v11
        }
    } else {
        if (0 == 32767) {
            if false {
                v4 += (v4 % ((-1) | 1))
            } else {
                var v12 float32 = 0.0
                if (2 < 65535) {
                    if (4294967295 <= -125) {
                        var v13 uint64 = 677
                        var v14 int = 9223372036854775807
                        var v15 int32 = -930
                        _ = v13
                        _ = v14
                        _ = v15
                    }
                    for fi16 := 0; fi16 < 14; fi16++ {
                        v4 = (v4 % (((-(v4))) | 1))
                        _ = fi16
                    }
                    v17 := GenS1{F0: 2, F1: 255}
                    _ = v17
                } else {
                    v4 = (v4 * v4)
                    for fi18 := 0; fi18 < 15; fi18++ {
                        v2 = append(v2, (fi18 | (fi18 >> ((18446744073709551615) & 63))))
                        _ = fi18
                    }
                    v4 += (v4 * v4)
                }
                if false {
                    var v19 bool = ((1 <= 1) && true)
                    _ = v19
                }
                _ = v12
            }
            if len(v3) > 0 {
                v3[0] = "a"
            }
        } else {
            var v20 uint32 = 2147483648
            if (-1 <= 42) {
                if (1e10 > (3.14159 + -2.25)) {
                    fmt.Println(v4)
                    switch "hello" {
                    case "hello":
                        if false {
                            var v21 uint64 = 0
                            _ = len(v3)
                            _ = v21
                        } else {
                            if (!true) {
                                for fi22 := 0; fi22 < 18; fi22++ {
                                    fmt.Println(v20, v20)
                                    var v23 int32 = 557
                                    if (!(-2.25 <= (-2.25 - float64((1e10 - 0.0))))) {
                                        var v24 int16 = (v4 * -32768)
                                        _ = v24
                                    }
                                    _ = fi22
                                    _ = v23
                                }
                            }
                        }
                    default:
                        v25 := GenS2{F0: (658 > 29), F1: -2147483648}
                        _ = v25
                    }
                    var v26 uint = 4294967295
                    _ = v26
                } else {
                    v27 := GenS1{F0: 42, F1: 1}
                    var v28 int = -722
                    var v29 float32 = 0.1
                    _ = v27
                    _ = v28
                    _ = v29
                }
                if (0 > 32768) {
                    var v30 uint8 = 39
                    if (("a" >= "a") && (false && ((!false) || (true && false)))) {
                        for fi31 := 0; fi31 < 19; fi31++ {
                            var v32 uint16 = 65535
                            _ = fi31
                            _ = v32
                        }
                        v20 = (v20 + 375)
                    } else {
                        v33 := GenS2{F0: true, F1: 2147483647}
                        _ = v33
                    }
                    _ = v30
                } else {
                    v34 := GenS3{F0: 287, F1: 42}
                    _ = v34
                }
            }
            var v35 int16 = 619
            _ = v20
            _ = v35
        }
        for fi36 := 0; fi36 < 6; fi36++ {
            if (v4 < 79) {
                var v37 float32 = 3.14159
                v4 += 174
                v4 -= v4
                _ = v37
            } else {
                for fi38 := 0; fi38 < 4; fi38++ {
                    v4 = (v4 | v4)
                    _ = fi38
                }
                fmt.Println(v4, fi36)
                if true {
                    var v39 bool = (-87 == -1)
                    _ = v39
                } else {
                    v40 := GenS2{F0: true, F1: 1}
                    _ = v40
                }
            }
            _ = fi36
        }
        v4 = (-(v4))
    }
    _ = len(v3)
    if len(v1) > 0 {
        v1[0] = -47
    }
    var v41 int8 = 1
    switch "a" {
    case "foo bar":
        for fi42 := 0; fi42 < 19; fi42++ {
            if false {
                var v43 uint8 = 255
                _ = v43
            }
            _ = fi42
        }
    case "\n\t":
        if ((3.14159 + (0.0 - 1e10)) <= 0.1) {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            v4 += v4
            v44 := GenS4{F0: 2, F1: (v41 - -1), F2: 282}
            _ = v44
        } else {
            if true {
                switch -1 {
                case 32768:
                    var v45 int8 = -128
                    _ = v45
                default:
                    switch ("" + (("hello" + "a") + (((((((("a" + "\n\t") + "foo bar") + "a") + ("foo bar" + "foo bar")) + ("a" + ("" + ("a" + "foo bar")))) + "a") + "") + "foo bar"))) {
                    case "foo bar":
                        _ = len(v2)
                    default:
                        switch 4294967295 {
                        case -575:
                            v41 = (v41 % ((v41) | 1))
                        case 1:
                            for fi46 := 0; fi46 < 15; fi46++ {
                                switch ("\n\t" + ("\n\t" + "")) {
                                case "\n\t":
                                    for fi47 := 0; fi47 < 7; fi47++ {
                                        var v48 string = ("\n\t" + (("a" + "hello") + "\n\t"))
                                        _ = fi47
                                        _ = v48
                                    }
                                case "hello":
                                    if false {
                                        fmt.Println(v4)
                                        var v49 float64 = float64(1.5)
                                        _ = v49
                                    }
                                case "foo bar":
                                    v41 = -128
                                default:
                                    if false {
                                        var v50 float32 = 0.1
                                        _ = v50
                                    }
                                }
                                _ = fi46
                            }
                        }
                    }
                }
            } else {
                v51 := GenS1{F0: 579, F1: 65535}
                if ((true && (!false)) && (1e10 >= 0.1)) {
                    if false {
                        if (false || true) {
                            switch -952 {
                            case -288:
                                var v52 string = ("" + "\n\t")
                                _ = v52
                            case 0:
                                if true {
                                    fmt.Println(v4, v4)
                                    if ((float64((0.1 - 1e10)) + 0.1) == 1.5) {
                                        v4 -= (v4 & (v4 - v4))
                                        if len(v1) > 0 {
                                            v1[0] = 42
                                        }
                                    } else {
                                        var v53 uint32 = 2
                                        fmt.Println(v41)
                                        _ = v53
                                    }
                                    fmt.Println(v4, v41)
                                }
                            case 42:
                                if false {
                                    v4 -= (v4 * v4)
                                } else {
                                    if (1.5 == 0.0) {
                                        var v54 uint16 = 255
                                        for fi55 := 0; fi55 < 17; fi55++ {
                                            var v56 int = (fi55 + (fi55 ^ int(v4)))
                                            v57 := GenS5{F0: 2, F1: "foo bar", F2: 0}
                                            switch "foo bar" {
                                            case "\n\t":
                                                v58 := GenS6{F0: (false || (!(102 != 128))), F1: 135}
                                                _ = v58
                                            case "foo bar":
                                                switch "foo bar" {
                                                case "a":
                                                    v56 = fi55
                                                default:
                                                    v4 = v4
                                                }
                                            case "":
                                                var v59 float64 = ((3.14159 + -2.25) + float64(0.1))
                                                _ = v59
                                            }
                                            _ = fi55
                                            _ = v56
                                            _ = v57
                                        }
                                        if false {
                                            if len(v2) > 0 {
                                                v2[0] = 127
                                            }
                                        } else {
                                            fmt.Println(v54, v54, v41)
                                            var v60 int = 0
                                            v61 := GenS7{F0: v41, F1: 18446744073709551615}
                                            _ = v60
                                            _ = v61
                                        }
                                        _ = v54
                                    }
                                    var v62 int8 = v41
                                    v63 := GenS4{F0: 1, F1: v62, F2: 32767}
                                    _ = v62
                                    _ = v63
                                }
                            }
                            v41 -= (v41 ^ (v41 - v41))
                        } else {
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            if (true || ((v4 & v4) < v4)) {
                                for fi64 := 0; fi64 < 19; fi64++ {
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    if false {
                                        for fi65 := 0; fi65 < 4; fi65++ {
                                            v66 := GenS7{F0: (v41 / ((-6) | 1)), F1: 2}
                                            if true {
                                                switch "\n\t" {
                                                case "foo bar":
                                                    var v67 int8 = (v41 / (((-(v41))) | 1))
                                                    _ = v67
                                                default:
                                                    v68 := GenS8{F0: 680}
                                                    _ = v68
                                                }
                                            } else {
                                                var v69 int8 = (v41 - (v41 | v41))
                                                _ = v69
                                            }
                                            v4 = -32768
                                            var v70 int64 = -1
                                            _ = fi65
                                            _ = v66
                                            _ = v70
                                        }
                                    }
                                    if false {
                                        v4 += (v4 + -289)
                                        fmt.Println(v41)
                                    } else {
                                        v41 += (-(v41))
                                        v71 := GenS4{F0: 0, F1: v41, F2: (fi64 + 32767)}
                                        _ = v71
                                    }
                                    _ = fi64
                                }
                            } else {
                                v2 = append(v2, 127)
                                fmt.Println(v41, v41)
                            }
                            for fi72 := 0; fi72 < 20; fi72++ {
                                switch "foo bar" {
                                case "":
                                    var v73 float32 = 0.0
                                    _ = v73
                                case "\n\t":
                                    switch fi72 {
                                    case 676:
                                        var v74 uint = 65535
                                        _ = v74
                                    case 2147483647:
                                        switch "a" {
                                        case "":
                                            fmt.Println(v4, v41, fi72)
                                        case "\n\t":
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                        case "a":
                                            v75 := GenS1{F0: 515, F1: 2}
                                            _ = v75
                                        default:
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                        }
                                    case 4294967295:
                                        v41 += v41
                                    default:
                                        for fi76 := 0; fi76 < 2; fi76++ {
                                            v41 -= (v41 % ((v41) | 1))
                                            v4 -= v4
                                            switch fi76 {
                                            case -666:
                                                fmt.Println(fi76, v41, v4)
                                            case 1:
                                                var v77 float64 = -2.25
                                                _ = v77
                                            default:
                                                v78 := GenS3{F0: 2, F1: 357}
                                                _ = v78
                                            }
                                            _ = fi76
                                        }
                                    }
                                case "hello":
                                    if (9223372036854775807 == 279) {
                                        v2 = append(v2, fi72)
                                        if true {
                                            v4 -= (^(v4))
                                            var v79 float64 = 0.1
                                            if (2 < 2) {
                                                v2 = append(v2, (-(fi72)))
                                            }
                                            _ = v79
                                        } else {
                                            var v80 int32 = 1
                                            var v81 uint = 6
                                            _ = v80
                                            _ = v81
                                        }
                                    }
                                default:
                                    _ = len(v1)
                                }
                                for fi82 := 0; fi82 < 2; fi82++ {
                                    var v83 int = fi72
                                    v84 := GenS3{F0: 4294967295, F1: 1}
                                    _ = fi82
                                    _ = v83
                                    _ = v84
                                }
                                v41 = (v41 ^ (v41 + (^(v41))))
                                _ = fi72
                            }
                        }
                    }
                }
                _ = v51
            }
            var v85 uint = 72
            _ = v85
        }
    }
    var v86 int64 = -1
    if (v86 > (v86 % (((^(v86))) | 1))) {
        fmt.Println(v41, v86, v41)
        var v87 float64 = 0.0
        _ = v87
    }
    v88 := GenS5{F0: 4294967295, F1: ((("foo bar" + ("hello" + "hello")) + ("foo bar" + ("a" + ("" + "")))) + "\n\t"), F2: 1}
    for fi89 := 0; fi89 < 8; fi89++ {
        for fi90 := 0; fi90 < 17; fi90++ {
            v4 += -190
            _ = fi90
        }
        v91 := GenS1{F0: 4294967295, F1: 881}
        _ = fi89
        _ = v91
    }
    fmt.Println(v41)
    var v92 uint32 = 122
    v3 = append(v3, "foo bar")
    fmt.Println(v86)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v86:", v86)
    fmt.Println("v92:", v92)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v86
    _ = v88
    _ = v92
}
