package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int; F2 uint8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint16; F1 int32 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint64; F1 uint; F2 int32 }
type GenS6 struct { F0 uint32; F1 int64; F2 string }
type GenS7 struct { F0 uint64; F1 int16; F2 int64 }
type GenS8 struct { F0 int16; F1 uint32; F2 uint64 }
type GenS9 struct { F0 uint32; F1 string; F2 uint16 }
type GenS10 struct { F0 uint; F1 uint }
type GenS11 struct { F0 int32; F1 int8 }
type GenS12 struct { F0 uint64; F1 bool }
type GenS13 struct { F0 uint16 }
func genFn0() int16 {
    var v6 int16 = 1
    var v7 int = 9223372036854775807
    if false {
        v8 := GenS1{F0: (1 < (v7 & (^(v7)))), F1: (v7 << ((209) & 63)), F2: 42}
        _ = v8
    }
    for fi9 := 0; fi9 < 14; fi9++ {
        v7 -= (v7 | v7)
        _ = fi9
    }
    _ = v6
    _ = v7
    return 42
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 int32 = -455
    v3 := []uint32{827, 2147483648, 1, 0}
    var v4 uint64 = 2
    var v5 uint16 = 599
    if true {
        v4 = 0
        switch "hello" {
        case "a":
            switch "\n\t" {
            case "foo bar":
                v10 := GenS2{F0: ("hello" + "foo bar")}
                _ = v10
            default:
                v11 := GenS3{F0: (v5 | 61), F1: v2}
                _ = v11
            }
        case "hello":
            for fi12 := 0; fi12 < 7; fi12++ {
                var v13 int32 = v2
                _ = fi12
                _ = v13
            }
        default:
            if (128 > 127) {
                v4 += (^(v4))
                var v14 uint8 = 127
                _ = v14
            } else {
                v4 += v4
                switch "hello" {
                case "foo bar":
                    for fi15 := 0; fi15 < 6; fi15++ {
                        var v16 uint32 = 598
                        v17 := GenS2{F0: ("\n\t" + "a")}
                        _ = fi15
                        _ = v16
                        _ = v17
                    }
                default:
                    var v18 float64 = ((-2.25 * float64(0.1)) / 1e10)
                    _ = v18
                }
            }
        }
        if (2 <= 0) {
            if len(v3) > 0 {
                v3[0] = 349
            }
        } else {
            var v19 int8 = 100
            _ = v19
        }
    } else {
        if false {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            for fi20 := 0; fi20 < 9; fi20++ {
                v21 := GenS1{F0: (!(0 <= -9223372036854775808)), F1: (fi20 >> ((18446744073709551615) & 63)), F2: 2}
                for fi22 := 0; fi22 < 7; fi22++ {
                    fmt.Println(v2, v2)
                    var v23 int64 = -9223372036854775808
                    if true {
                        var v24 int32 = v2
                        switch "hello" {
                        case "\n\t":
                            var v25 float64 = 1.5
                            _ = v25
                        case "a":
                            var v26 string = ""
                            _ = v26
                        case "hello":
                            v24 = v2
                        }
                        var v27 int32 = v24
                        _ = v24
                        _ = v27
                    }
                    switch "foo bar" {
                    case "a":
                        _ = len(v3)
                    case "":
                        v28 := GenS4{F0: 38}
                        _ = v28
                    case "\n\t":
                        fmt.Println(fi22)
                    default:
                        if len(v3) > 0 {
                            v3[0] = 117
                        }
                    }
                    if false {
                        if (!false) {
                            v29 := GenS5{F0: 2, F1: 2, F2: v2}
                            var v30 bool = ((fi20 ^ (fi20 * fi22)) == (fi20 | 4294967295))
                            v5 += v5
                            _ = v29
                            _ = v30
                        }
                        v23 = -1
                    } else {
                        var v31 float32 = 1e10
                        v2 -= 2147483647
                        _ = v31
                    }
                    var v32 int = fi20
                    _ = fi22
                    _ = v23
                    _ = v32
                }
                v4 -= v4
                _ = fi20
                _ = v21
            }
        }
        for fi33 := 0; fi33 < 15; fi33++ {
            if false {
                var v34 int = fi33
                _ = v34
            }
            if (uint8(v5) > 42) {
                v35 := GenS6{F0: 65535, F1: 42, F2: (("\n\t" + (("foo bar" + ((("a" + "a") + ("foo bar" + "")) + "\n\t")) + "a")) + "foo bar")}
                _ = v35
            }
            var v36 int64 = 0
            _ = fi33
            _ = v36
        }
        for fi37 := 0; fi37 < 17; fi37++ {
            v5 = 65535
            if (42 <= 79) {
                if true {
                    if len(v3) > 0 {
                        v3[0] = 449
                    }
                    if true {
                        fmt.Println(v2)
                    }
                    if true {
                        fmt.Println(v2)
                        switch fi37 {
                        case 9223372036854775807:
                            if true {
                                for fi38 := 0; fi38 < 20; fi38++ {
                                    v39 := GenS7{F0: v4, F1: 38, F2: 0}
                                    v5 = v5
                                    _ = fi38
                                    _ = v39
                                }
                                v2 = v2
                            }
                        case 4294967295:
                            v2 = v2
                        default:
                            var v40 uint32 = 55
                            _ = v40
                        }
                    } else {
                        for fi41 := 0; fi41 < 18; fi41++ {
                            var v42 string = "\n\t"
                            var v43 uint = 65535
                            v44 := GenS8{F0: 1, F1: 366, F2: 2}
                            _ = fi41
                            _ = v42
                            _ = v43
                            _ = v44
                        }
                        for fi45 := 0; fi45 < 16; fi45++ {
                            if true {
                                var v46 int64 = -1
                                switch (fi45 << ((2) & 63)) {
                                case 937:
                                    var v47 uint8 = 128
                                    _ = v47
                                }
                                _ = v46
                            }
                            _ = len(v1)
                            v5 = 0
                            _ = fi45
                        }
                    }
                } else {
                    if len(v3) > 0 {
                        v3[0] = 4294967295
                    }
                    if true {
                        var v48 int8 = 42
                        v49 := GenS9{F0: 4294967295, F1: ("a" + "\n\t"), F2: (^(v5))}
                        _ = v48
                        _ = v49
                    } else {
                        v4 = uint64(v5)
                        switch "a" {
                        case "\n\t":
                            for fi50 := 0; fi50 < 9; fi50++ {
                                v51 := v1[2147483647]
                                if (v5 == v5) {
                                    if v51 {
                                        v2 = (v2 | 2147483647)
                                    }
                                    if (-2.25 != (((3.14159 / 3.14159) * (3.14159 * float64(0.0))) - 0.1)) {
                                        if v51 {
                                            v52 := GenS9{F0: 4294967295, F1: "foo bar", F2: (^(v5))}
                                            var v53 int = (fi50 % ((fi50) | 1))
                                            v54 := v1[uint32(v5)]
                                            _ = v52
                                            _ = v53
                                            _ = v54
                                        }
                                        fmt.Println(v51, fi50, v51)
                                    } else {
                                        v4 = v4
                                    }
                                } else {
                                    v51 = v51
                                }
                                _ = fi50
                                _ = v51
                            }
                        case "":
                            if false {
                                v3 = append(v3, 2)
                                v55 := GenS9{F0: 2, F1: (("\n\t" + ((("hello" + "") + "hello") + "hello")) + (("foo bar" + ("" + "hello")) + "foo bar")), F2: v5}
                                _ = v55
                            } else {
                                var v56 int16 = -32768
                                v1[1] = (-2.25 < (1e10 - (3.14159 * ((((3.14159 - ((1.5 - 1e10) * 1.5)) * float32(-2.25)) * (((1e10 * (0.1 + (3.14159 / float32((0.0 - 0.1))))) - (1e10 / 3.14159)) * ((-2.25 - (3.14159 / -2.25)) + 0.0))) / 3.14159))))
                                if (!false) {
                                    for fi57 := 0; fi57 < 14; fi57++ {
                                        var v58 float32 = 0.1
                                        var v59 int8 = 0
                                        _ = fi57
                                        _ = v58
                                        _ = v59
                                    }
                                    var v60 int32 = v2
                                    switch -9223372036854775808 {
                                    case 127:
                                        for fi61 := 0; fi61 < 4; fi61++ {
                                            v56 -= v56
                                            _ = fi61
                                        }
                                    case 2147483647:
                                        v4 = v4
                                    case 9223372036854775807:
                                        fmt.Println(v5)
                                    default:
                                        for fi62 := 0; fi62 < 18; fi62++ {
                                            var v63 uint16 = v5
                                            fmt.Println(v63, fi62)
                                            _ = fi62
                                            _ = v63
                                        }
                                    }
                                    _ = v60
                                }
                                _ = v56
                            }
                        case "foo bar":
                            v64 := GenS3{F0: (v5 >> ((2) & 63)), F1: 1}
                            _ = v64
                        }
                        if (!true) {
                            switch fi37 {
                            case 32767:
                                v4 += v4
                            }
                            if len(v3) > 0 {
                                v3[0] = 2147483647
                            }
                            v65 := GenS7{F0: 4294967295, F1: 610, F2: -9223372036854775808}
                            _ = v65
                        } else {
                            v66 := GenS10{F0: 1, F1: 733}
                            if len(v3) > 0 {
                                v3[0] = 4294967295
                            }
                            _ = v66
                        }
                    }
                }
            } else {
                var v67 bool = false
                if ((v2 * v2) <= 2147483647) {
                    fmt.Println(v4, fi37)
                    v4 -= (v4 ^ v4)
                    fmt.Println(fi37, fi37, v67)
                } else {
                    for fi68 := 0; fi68 < 14; fi68++ {
                        v4 = v4
                        _ = fi68
                    }
                    if v67 {
                        v69 := GenS11{F0: (v2 * (-(v2))), F1: 100}
                        if true {
                            v3 = append(v3, 519)
                            var v70 float32 = (1.5 * (0.0 * 1.5))
                            _ = v70
                        } else {
                            if len(v3) > 0 {
                                v3[0] = 2147483647
                            }
                            delete(v1, 42)
                            for fi71 := 0; fi71 < 2; fi71++ {
                                v5 += (^(v5))
                                switch "foo bar" {
                                case "foo bar":
                                    if false {
                                        v67 = v67
                                        v2 += v2
                                        if (!v67) {
                                            var v72 uint32 = 65535
                                            fmt.Println(fi71, v67, v2)
                                            _ = v72
                                        } else {
                                            v73 := GenS12{F0: 0, F1: (!true)}
                                            if true {
                                                switch fi37 {
                                                case 0:
                                                    if false {
                                                        v1[941] = false
                                                    }
                                                    v74 := GenS13{F0: v5}
                                                    var v75 string = "\n\t"
                                                    _ = v74
                                                    _ = v75
                                                case 1:
                                                    switch 647 {
                                                    case 0:
                                                        if (((true && v67) || v67) && v67) {
                                                            var v76 float64 = -2.25
                                                            v5 = v5
                                                            var v77 float64 = float64(float32((v76 - 1.5)))
                                                            _ = v76
                                                            _ = v77
                                                        }
                                                    case -805:
                                                        delete(v1, 751)
                                                    default:
                                                        var v78 uint8 = 255
                                                        _ = v78
                                                    }
                                                case 42:
                                                    v4 -= (v4 & (v4 << ((132) & 63)))
                                                default:
                                                    var v79 int16 = 32767
                                                    _ = v79
                                                }
                                                var v80 int8 = -1
                                                v67 = ((true || (false || (65535 <= uint32(v4)))) && v67)
                                                _ = v80
                                            } else {
                                                v81 := GenS10{F0: 4294967295, F1: 18446744073709551615}
                                                _ = v81
                                            }
                                            _ = v73
                                        }
                                    }
                                }
                                fmt.Println(fi71, v4, v4)
                                _ = fi71
                            }
                        }
                        _ = v69
                    }
                }
                var v82 string = (("" + ("" + ("\n\t" + ((("" + ("foo bar" + "")) + (("hello" + "foo bar") + ("foo bar" + "\n\t"))) + "foo bar")))) + "\n\t")
                _ = v67
                _ = v82
            }
            var v83 float64 = (float64(-2.25) - 3.14159)
            _ = fi37
            _ = v83
        }
    }
    fmt.Println(v2)
    var v84 float64 = (0.1 + 0.1)
    var v85 uint32 = 965
    v86 := GenS2{F0: (("" + (("foo bar" + ("\n\t" + "a")) + "\n\t")) + "a")}
    var v87 int8 = -91
    var v88 uint8 = 214
    v89 := v1[(v85 & (^(v85)))]
    fmt.Println(v88, v85, v85)
    fmt.Println(v2, v88)
    switch ("\n\t" + ("foo bar" + (("\n\t" + "") + ""))) {
    case "\n\t":
        fmt.Println(v4)
    }
    var v90 float32 = (float32((v84 / -2.25)) + -2.25)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v84:", v84)
    fmt.Println("v85:", v85)
    fmt.Println("v87:", v87)
    fmt.Println("v88:", v88)
    fmt.Println("v89:", v89)
    fmt.Println("v90:", v90)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v84
    _ = v85
    _ = v86
    _ = v87
    _ = v88
    _ = v89
    _ = v90
}
