package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string }
type GenS2 struct { F0 uint64; F1 uint16; F2 uint32 }
type GenS3 struct { F0 uint32; F1 int }
type GenS4 struct { F0 int; F1 bool; F2 int64 }
type GenS5 struct { F0 uint; F1 uint8; F2 uint8 }
type GenS6 struct { F0 int64; F1 uint32 }
type GenS7 struct { F0 int }
type GenS8 struct { F0 uint64 }
type GenS9 struct { F0 uint16; F1 uint8; F2 uint64 }
type GenS10 struct { F0 uint8; F1 int }
type GenS11 struct { F0 string; F1 int16 }
type GenS12 struct { F0 uint32; F1 int32 }
type GenS13 struct { F0 int16; F1 int8 }
type GenS14 struct { F0 bool }
type GenS15 struct { F0 int16 }
type GenS16 struct { F0 int; F1 bool; F2 string }
type GenS17 struct { F0 int64; F1 int }
func genFn0() int32 {
    var v7 int32 = 2147483647
    for fi8 := 0; fi8 < 5; fi8++ {
        v7 += (v7 ^ (v7 | v7))
        var v9 uint8 = 42
        if true {
            var v10 uint8 = v9
            v9 = (v9 & v10)
            if false {
                var v11 uint8 = v10
                for fi12 := 0; fi12 < 1; fi12++ {
                    v9 = 212
                    switch ((("a" + "foo bar") + (("a" + ("" + ((("a" + "a") + ("" + "\n\t")) + "\n\t"))) + "\n\t")) + "hello") {
                    case "hello":
                        v13 := GenS1{F0: 2, F1: "\n\t"}
                        _ = v13
                    default:
                        v11 -= v9
                    }
                    _ = fi12
                }
                _ = v11
            }
            _ = v10
        } else {
            fmt.Println(fi8, fi8)
            for fi14 := 0; fi14 < 13; fi14++ {
                v15 := []int{(^(fi14)), fi8, fi14, 1}
                if len(v15) >= 2 {
                    v15 = v15[1:len(v15)]
                }
                _ = fi14
                _ = v15
            }
        }
        _ = fi8
        _ = v9
    }
    v7 += (v7 % (((v7 << ((4294967295) & 63))) | 1))
    _ = v7
    return v7
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := []string{("hello" + ("hello" + "foo bar"))}
    var v3 float64 = 3.14159
    v4 := make(map[uint32]int)
    v5 := make(map[string]uint32)
    var v6 uint16 = 682
    var v16 uint8 = uint8(v6)
    if true {
        for fi17 := 0; fi17 < 2; fi17++ {
            if len(v2) > 0 {
                v2[0] = "foo bar"
            }
            _ = fi17
        }
        var v18 uint64 = 18446744073709551615
        v19 := GenS1{F0: 776, F1: "foo bar"}
        _ = v18
        _ = v19
    }
    for fi20 := 0; fi20 < 19; fi20++ {
        for fi21 := 0; fi21 < 18; fi21++ {
            if false {
                switch "hello" {
                case "":
                    if false {
                        v22 := GenS2{F0: 26, F1: v6, F2: 2147483647}
                        v3 -= 0.1
                        switch 2147483647 {
                        case 2147483647:
                            if false {
                                fmt.Println(v16)
                            }
                        default:
                            v16 += (v16 / ((127) | 1))
                        }
                        _ = v22
                    }
                }
            } else {
                v23 := v5["foo bar"]
                var v24 string = "a"
                var v25 float64 = ((v3 - v3) - v3)
                _ = v23
                _ = v24
                _ = v25
            }
            if false {
                v3 = 1.5
                if (true || false) {
                    v26 := GenS3{F0: 4294967295, F1: (fi21 & 32768)}
                    _ = v26
                }
                switch fi21 {
                case 98:
                    v16 += v16
                case 1:
                    var v27 string = ""
                    _ = v27
                }
                if true {
                    if len(v2) > 0 {
                        v2[0] = ("foo bar" + (("hello" + "hello") + (("" + "\n\t") + "")))
                    }
                } else {
                    var v28 int16 = 32767
                    _ = v28
                }
            } else {
                v3 = float64(-2.25)
                switch "\n\t" {
                case "":
                    v6 -= v6
                case "hello":
                    v29 := GenS2{F0: 750, F1: 32767, F2: 23}
                    _ = v29
                case "a":
                    if false {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    }
                default:
                    v16 -= v16
                }
            }
            v6 = v6
            _ = fi21
        }
        for fi30 := 0; fi30 < 2; fi30++ {
            var v31 int8 = int8(fi30)
            _ = fi30
            _ = v31
        }
        fmt.Println(v3, v16)
        _ = fi20
    }
    var v32 uint8 = 128
    v3 += v3
    for fi33 := 0; fi33 < 8; fi33++ {
        v6 = (v6 ^ 0)
        v16 = 0
        switch "foo bar" {
        case "a":
            for fi34 := 0; fi34 < 15; fi34++ {
                v3 = (1.5 + v3)
                _ = fi34
            }
        default:
            if (false && ((^(v6)) != v6)) {
                switch "\n\t" {
                case "foo bar":
                    fmt.Println(v32, v32)
                case "hello":
                    v16 += v16
                default:
                    if true {
                        switch "a" {
                        case "\n\t":
                            fmt.Println(v6, v6)
                        case "hello":
                            fmt.Println(fi33, v16)
                        case "foo bar":
                            var v35 uint64 = 18446744073709551615
                            _ = v35
                        default:
                            v6 += 2
                        }
                        _ = len(v2)
                        switch 31 {
                        case 128:
                            var v36 int64 = -9223372036854775808
                            _ = v36
                        case -1:
                            v37 := GenS1{F0: uint64(v6), F1: ("" + "hello")}
                            _ = v37
                        }
                    } else {
                        v6 = 543
                    }
                }
                v16 = (v16 % ((2) | 1))
                if false {
                    v3 += ((v3 * float64(0.0)) / ((((v3 + (1.5 * (v3 - (((float64(0.1) + (float64(0.0) - 1.5)) - (v3 + v3)) * v3)))) + 1e10) * -2.25) / v3))
                } else {
                    var v38 int8 = 74
                    if false {
                        for fi39 := 0; fi39 < 10; fi39++ {
                            v3 = ((v3 + v3) - v3)
                            var v40 int32 = genFn0()
                            v38 += v38
                            switch (fi39 | fi39) {
                            case -122:
                                var v41 uint = uint(v6)
                                _ = v41
                            case 32768:
                                switch "foo bar" {
                                case "hello":
                                    fmt.Println(v16, v3, v38)
                                case "\n\t":
                                    var v42 int32 = v40
                                    _ = v42
                                case "":
                                    v43 := GenS3{F0: 230, F1: (^(fi39))}
                                    _ = v43
                                default:
                                    if false {
                                        switch "foo bar" {
                                        case "":
                                            _ = len(v1)
                                        }
                                        if true {
                                            if (3.14159 != (((1.5 - 0.0) * 0.0) / ((float32(((v3 / v3) / v3)) * (((1.5 * (((1e10 * (1e10 / float32((v3 / v3)))) / (float32((float64(0.0) + (v3 - v3))) + 0.1)) - 1e10)) * float32(v3)) * (0.0 - 0.1))) * 1e10))) {
                                                var v44 bool = true
                                                _ = v44
                                            } else {
                                                v40 -= int32(v6)
                                            }
                                        } else {
                                            switch ("" + "a") {
                                            case "":
                                                if (1.5 != float32(0.1)) {
                                                    switch ("" + "") {
                                                    case "":
                                                        v40 -= genFn0()
                                                    default:
                                                        var v45 float64 = (1e10 - v3)
                                                        _ = v45
                                                    }
                                                    switch (fi33 >> ((0) & 63)) {
                                                    case 9223372036854775807:
                                                        var v46 int16 = 42
                                                        _ = v46
                                                    default:
                                                        v47 := GenS4{F0: fi39, F1: true, F2: int64(v32)}
                                                        _ = v47
                                                    }
                                                    switch "\n\t" {
                                                    case "":
                                                        if true {
                                                            var v48 float32 = -2.25
                                                            if (711 == 18446744073709551615) {
                                                                if (!(true && true)) {
                                                                    if true {
                                                                        switch 1 {
                                                                        case 9223372036854775807:
                                                                            fmt.Println(fi39)
                                                                        case -9223372036854775808:
                                                                            delete(v5, (("a" + "hello") + ""))
                                                                        }
                                                                        delete(v4, 91)
                                                                        v16 -= (v32 | v32)
                                                                    }
                                                                    fmt.Println(fi39, v38)
                                                                } else {
                                                                    var v49 uint16 = 714
                                                                    _ = len(v5)
                                                                    var v50 int32 = (^(v40))
                                                                    _ = v49
                                                                    _ = v50
                                                                }
                                                                var v51 bool = (true && false)
                                                                if (fi33 >= 979) {
                                                                    if (!((239 < 12) && v51)) {
                                                                        switch "hello" {
                                                                        case "\n\t":
                                                                            var v52 uint32 = uint32(v6)
                                                                            _ = v52
                                                                        case "hello":
                                                                            v40 += 371
                                                                        }
                                                                    } else {
                                                                        if (!true) {
                                                                            if ((!v51) || (0.1 <= 1e10)) {
                                                                                switch fi39 {
                                                                                case 9223372036854775807:
                                                                                    v3 += (v3 * v3)
                                                                                    var v53 int = (-(fi39))
                                                                                    _ = v53
                                                                                case 414:
                                                                                    switch ("a" + ("" + (("foo bar" + ("\n\t" + ("a" + ""))) + ""))) {
                                                                                    case "foo bar":
                                                                                        v54 := GenS5{F0: 389, F1: (^(v16)), F2: v32}
                                                                                        _ = v54
                                                                                    case "hello":
                                                                                        switch fi33 {
                                                                                        case 0:
                                                                                            if ((((1.5 / ((1.5 / v48) / ((0.0 - float32(v3)) - v48))) + 0.1) + v48) >= -2.25) {
                                                                                                var v55 int = (fi33 >> ((4294967295) & 63))
                                                                                                var v56 bool = (((v3 / 1e10) - 3.14159) < v3)
                                                                                                fmt.Println(v38)
                                                                                                _ = v55
                                                                                                _ = v56
                                                                                            } else {
                                                                                                var v57 float64 = (float64((v48 / v48)) - 1.5)
                                                                                                _ = v57
                                                                                            }
                                                                                        case 128:
                                                                                            if (4294967295 >= 0) {
                                                                                                v58 := GenS6{F0: 405, F1: 4294967295}
                                                                                                var v59 uint32 = 65535
                                                                                                var v60 int64 = 687
                                                                                                _ = v58
                                                                                                _ = v59
                                                                                                _ = v60
                                                                                            }
                                                                                        default:
                                                                                            var v61 string = "a"
                                                                                            _ = v61
                                                                                        }
                                                                                    }
                                                                                }
                                                                                switch 2147483647 {
                                                                                case 88:
                                                                                    var v62 string = ("" + "a")
                                                                                    _ = v62
                                                                                case -2147483648:
                                                                                    _ = len(v5)
                                                                                default:
                                                                                    if false {
                                                                                        switch "\n\t" {
                                                                                        case "":
                                                                                            if v51 {
                                                                                                delete(v5, ("\n\t" + ("foo bar" + "hello")))
                                                                                            } else {
                                                                                                if v51 {
                                                                                                    v40 = v40
                                                                                                    var v63 int32 = v40
                                                                                                    v64 := GenS7{F0: fi39}
                                                                                                    fmt.Println(v38)
                                                                                                    var v65 string = ((("a" + (("foo bar" + ("a" + ("hello" + ("hello" + "foo bar")))) + "a")) + "\n\t") + ("\n\t" + "a"))
                                                                                                    _ = v63
                                                                                                    _ = v64
                                                                                                    _ = v65
                                                                                                } else {
                                                                                                    v3 = v3
                                                                                                    if v51 {
                                                                                                        if (v32 <= v32) {
                                                                                                            if v51 {
                                                                                                                v32 = v32
                                                                                                                v6 += v6
                                                                                                                fmt.Println(v16, v16, fi33)
                                                                                                                if (42 >= 942) {
                                                                                                                    v32 += 0
                                                                                                                    _ = len(v4)
                                                                                                                    var v66 int16 = -553
                                                                                                                    _ = v66
                                                                                                                }
                                                                                                                if v51 {
                                                                                                                    v67 := GenS8{F0: 18446744073709551615}
                                                                                                                    var v68 int = fi33
                                                                                                                    _ = v67
                                                                                                                    _ = v68
                                                                                                                }
                                                                                                            } else {
                                                                                                                var v69 int8 = v38
                                                                                                                _ = v69
                                                                                                            }
                                                                                                        } else {
                                                                                                            v40 += -415
                                                                                                        }
                                                                                                        v48 = v48
                                                                                                        fmt.Println(v16)
                                                                                                    }
                                                                                                    switch "foo bar" {
                                                                                                    case "hello":
                                                                                                        var v70 float64 = v3
                                                                                                        _ = v70
                                                                                                    case "\n\t":
                                                                                                        _ = len(v4)
                                                                                                    default:
                                                                                                        var v71 int = fi39
                                                                                                        _ = v71
                                                                                                    }
                                                                                                    switch 0 {
                                                                                                    case 0:
                                                                                                        var v72 float64 = v3
                                                                                                        _ = v72
                                                                                                    default:
                                                                                                        if v51 {
                                                                                                            var v73 int64 = -9223372036854775808
                                                                                                            if true {
                                                                                                                switch int(v38) {
                                                                                                                case 410:
                                                                                                                    if v51 {
                                                                                                                        v74 := GenS9{F0: v6, F1: v32, F2: 230}
                                                                                                                        _ = v74
                                                                                                                    } else {
                                                                                                                        if v51 {
                                                                                                                            v75 := GenS8{F0: 317}
                                                                                                                            v38 -= v38
                                                                                                                            var v76 int16 = -919
                                                                                                                            _ = v75
                                                                                                                            _ = v76
                                                                                                                        }
                                                                                                                        if len(v2) > 0 {
                                                                                                                            v2[0] = "foo bar"
                                                                                                                        }
                                                                                                                    }
                                                                                                                }
                                                                                                                if false {
                                                                                                                    v40 += v40
                                                                                                                    v77 := v4[2147483648]
                                                                                                                    var v78 uint32 = 2147483648
                                                                                                                    v79 := v4[65535]
                                                                                                                    _ = v77
                                                                                                                    _ = v78
                                                                                                                    _ = v79
                                                                                                                } else {
                                                                                                                    v80 := GenS10{F0: v16, F1: fi33}
                                                                                                                    var v81 float32 = (((0.0 - v48) - v48) + v48)
                                                                                                                    var v82 uint = 4294967295
                                                                                                                    v83 := GenS2{F0: 622, F1: (v6 | uint16(v16)), F2: 4294967295}
                                                                                                                    _ = v80
                                                                                                                    _ = v81
                                                                                                                    _ = v82
                                                                                                                    _ = v83
                                                                                                                }
                                                                                                            } else {
                                                                                                                v84 := GenS9{F0: 255, F1: v32, F2: 220}
                                                                                                                v51 = true
                                                                                                                v40 = (v40 / (((v40 >> ((644) & 63))) | 1))
                                                                                                                var v85 uint16 = (v6 * v6)
                                                                                                                v51 = ((v40 % ((int32(v38)) | 1)) >= (-(v40)))
                                                                                                                _ = v84
                                                                                                                _ = v85
                                                                                                            }
                                                                                                            fmt.Println(v38, v40, v3)
                                                                                                            _ = v73
                                                                                                        }
                                                                                                    }
                                                                                                }
                                                                                            }
                                                                                        }
                                                                                        v1[42] = (v51 || v51)
                                                                                        var v86 uint32 = 518
                                                                                        if false {
                                                                                            var v87 int32 = (v40 + -2147483648)
                                                                                            _ = v87
                                                                                        } else {
                                                                                            v40 -= genFn0()
                                                                                            if v51 {
                                                                                                v86 += (v86 << ((1) & 63))
                                                                                            } else {
                                                                                                fmt.Println(fi39)
                                                                                                if ((v51 && v51) && ((!(!(127 <= v38))) || (0 > (-(v38))))) {
                                                                                                    fmt.Println(v86, v16, v32)
                                                                                                    var v88 float64 = (((v3 / v3) / ((1.5 * (v3 + v3)) - (float64(float32(((0.0 + v3) / v3))) / 1e10))) * float64((v48 * 3.14159)))
                                                                                                    _ = v88
                                                                                                }
                                                                                            }
                                                                                            var v89 float32 = (float32(v3) * (0.0 / v48))
                                                                                            if v51 {
                                                                                                v90 := GenS6{F0: 9223372036854775807, F1: 212}
                                                                                                _ = v90
                                                                                            } else {
                                                                                                var v91 bool = (v51 && ((v86 & (v86 / ((v86) | 1))) <= v86))
                                                                                                _ = v91
                                                                                            }
                                                                                            _ = v89
                                                                                        }
                                                                                        _ = v86
                                                                                    } else {
                                                                                        switch fi33 {
                                                                                        case -271:
                                                                                            if false {
                                                                                                _ = len(v2)
                                                                                            } else {
                                                                                                v92 := GenS3{F0: 65535, F1: (fi39 ^ fi33)}
                                                                                                v6 -= v6
                                                                                                _ = v92
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                }
                                                                                if len(v2) >= 2 {
                                                                                    v2 = v2[1:len(v2)]
                                                                                }
                                                                            } else {
                                                                                if (0 != 42) {
                                                                                    if (2 <= 15) {
                                                                                        var v93 uint32 = 664
                                                                                        _ = v93
                                                                                    } else {
                                                                                        var v94 uint64 = 4294967295
                                                                                        v40 += (^(v40))
                                                                                        v95 := GenS4{F0: (fi39 * fi39), F1: (!(v6 > v6)), F2: 1}
                                                                                        delete(v1, 4294967295)
                                                                                        var v96 int64 = 1
                                                                                        if v51 {
                                                                                            var v97 uint32 = 2147483648
                                                                                            v3 -= v3
                                                                                            v2 = append(v2, "foo bar")
                                                                                            v98 := GenS1{F0: v94, F1: ""}
                                                                                            _ = v97
                                                                                            _ = v98
                                                                                        }
                                                                                        _ = v94
                                                                                        _ = v95
                                                                                        _ = v96
                                                                                    }
                                                                                    switch "" {
                                                                                    case "foo bar":
                                                                                        v99 := GenS11{F0: ("\n\t" + "a"), F1: 32767}
                                                                                        v100 := GenS12{F0: 2, F1: (v40 - (v40 >> ((761) & 63)))}
                                                                                        v101 := GenS6{F0: -9223372036854775808, F1: 399}
                                                                                        _ = v100
                                                                                        _ = v101
                                                                                        _ = v99
                                                                                    case "a":
                                                                                        delete(v5, "hello")
                                                                                    case "\n\t":
                                                                                        if (false && true) {
                                                                                            if len(v2) >= 2 {
                                                                                                v2 = v2[1:len(v2)]
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                }
                                                                            }
                                                                            var v102 uint16 = 0
                                                                            _ = v102
                                                                        }
                                                                        if ("hello" != "\n\t") {
                                                                            var v103 int64 = 9223372036854775807
                                                                            var v104 int16 = 1
                                                                            v16 = v32
                                                                            _ = v103
                                                                            _ = v104
                                                                        } else {
                                                                            v16 -= v16
                                                                            v3 += float64((((0.0 - (v48 + v48)) * v48) + v48))
                                                                        }
                                                                        v105 := GenS6{F0: 42, F1: 175}
                                                                        _ = v105
                                                                    }
                                                                    if len(v2) > 0 {
                                                                        v2[0] = ""
                                                                    }
                                                                    v3 = v3
                                                                    v51 = true
                                                                }
                                                                _ = v51
                                                            } else {
                                                                v106 := GenS9{F0: (v6 >> ((1) & 63)), F1: 127, F2: 1}
                                                                v2 = append(v2, "a")
                                                                v107 := GenS10{F0: v16, F1: fi33}
                                                                _ = v106
                                                                _ = v107
                                                            }
                                                            fmt.Println(fi39, v32, v3)
                                                            v32 = 2
                                                            var v108 int = (fi39 + fi39)
                                                            var v109 int = -888
                                                            if len(v2) > 0 {
                                                                v2[0] = ""
                                                            }
                                                            _ = v108
                                                            _ = v109
                                                            _ = v48
                                                        } else {
                                                            fmt.Println(v16, v40)
                                                        }
                                                    }
                                                }
                                            }
                                            var v110 int8 = 0
                                            _ = len(v5)
                                            _ = v110
                                        }
                                        var v111 int16 = 0
                                        var v112 int8 = int8(v6)
                                        _ = v111
                                        _ = v112
                                    }
                                    var v113 float32 = 1.5
                                    _ = v113
                                }
                            }
                            v40 = v40
                            if true {
                                v114 := GenS13{F0: 42, F1: 127}
                                v40 = (v40 | (-(v40)))
                                fmt.Println(v3, v6)
                                _ = v114
                            } else {
                                var v115 uint8 = v16
                                _ = v115
                            }
                            var v116 uint8 = (v16 % ((v32) | 1))
                            v1[65535] = (-32768 > 1000)
                            v117 := GenS14{F0: false}
                            _ = fi39
                            _ = v116
                            _ = v117
                            _ = v40
                        }
                        var v118 uint = 18446744073709551615
                        v32 = (v16 ^ 42)
                        _ = v118
                    }
                    v119 := GenS15{F0: 32767}
                    _ = v119
                    _ = v38
                }
            } else {
                if (false && (2 == 4294967295)) {
                    var v120 int32 = 42
                    v32 = 0
                    v32 -= (v16 + (v32 % ((1) | 1)))
                    _ = v120
                }
                var v121 int16 = 962
                _ = v121
            }
        }
        _ = fi33
    }
    v16 -= 0
    for fi122 := 0; fi122 < 3; fi122++ {
        var v123 int8 = 40
        for fi124 := 0; fi124 < 18; fi124++ {
            fmt.Println(v3, v123)
            _ = fi124
        }
        switch 2147483647 {
        case -9223372036854775808:
            var v125 float64 = (v3 * (((0.0 / float64(float32((float64((float32(float64(3.14159)) + 1e10)) + v3)))) + (3.14159 + (((v3 * v3) * (v3 - 1e10)) - (1.5 - ((0.1 + (0.1 - v3)) / (v3 * float64((float32(0.1) - 0.1)))))))) * (v3 + (v3 + 1.5))))
            _ = v125
        case -2147483648:
            var v126 int = (fi122 & 2147483647)
            _ = v126
        default:
            var v127 int8 = v123
            _ = v127
        }
        _ = fi122
        _ = v123
    }
    v6 = 164
    fmt.Println(v32, v32)
    fmt.Println(v32, v32)
    v128 := GenS2{F0: 467, F1: v6, F2: 381}
    var v129 uint64 = 53
    delete(v4, 135)
    for fi130 := 0; fi130 < 18; fi130++ {
        if false {
            for fi131 := 0; fi131 < 11; fi131++ {
                var v132 int64 = -1
                if (0 > v132) {
                    _ = len(v1)
                    switch (fi130 + (fi131 >> ((65535) & 63))) {
                    case -9223372036854775808:
                        v6 -= v6
                    case -74:
                        var v133 uint8 = v16
                        _ = v133
                    case 2147483647:
                        switch "hello" {
                        case "a":
                            var v134 int64 = -1
                            _ = v134
                        case "":
                            v6 = 42
                        case "foo bar":
                            if (!true) {
                                var v135 uint8 = 42
                                v16 += 127
                                v136 := GenS1{F0: 42, F1: ""}
                                _ = v135
                                _ = v136
                            } else {
                                fmt.Println(v32)
                                v2 = append(v2, ("foo bar" + ("" + ("foo bar" + ""))))
                                v32 -= v32
                            }
                        }
                        var v137 uint64 = (^(v129))
                        v138 := GenS16{F0: fi131, F1: (!((2147483647 < (fi130 >> ((uint(v6)) & 63))) || (127 != int8(v32)))), F2: "hello"}
                        _ = v137
                        _ = v138
                    default:
                        fmt.Println(v3)
                    }
                } else {
                    var v139 int8 = 42
                    v140 := GenS11{F0: "hello", F1: -859}
                    _ = v139
                    _ = v140
                }
                v16 += v32
                _ = fi131
                _ = v132
            }
            for fi141 := 0; fi141 < 3; fi141++ {
                _ = len(v1)
                switch "a" {
                case "\n\t":
                    fmt.Println(v16, v32, fi130)
                default:
                    var v142 int64 = -9223372036854775808
                    _ = v142
                }
                var v143 float64 = v3
                _ = fi141
                _ = v143
            }
        }
        var v144 int16 = 0
        v3 -= ((((float64(3.14159) - v3) * -2.25) - 0.1) * 0.1)
        _ = fi130
        _ = v144
    }
    if (((v129 >> ((42) & 63)) < v129) || (float64(3.14159) != (v3 + ((3.14159 / ((0.1 / (v3 / 1.0)) - v3)) - ((v3 * v3) - v3))))) {
        fmt.Println(v16, v129)
    }
    if (-139 > 42) {
        delete(v4, uint32(v129))
        switch "hello" {
        case "hello":
            for fi145 := 0; fi145 < 11; fi145++ {
                _ = len(v2)
                for fi146 := 0; fi146 < 14; fi146++ {
                    v2 = append(v2, "")
                    var v147 bool = (!false)
                    v148 := GenS17{F0: 152, F1: fi146}
                    _ = fi146
                    _ = v147
                    _ = v148
                }
                _ = fi145
            }
        case "foo bar":
            for fi149 := 0; fi149 < 11; fi149++ {
                for fi150 := 0; fi150 < 18; fi150++ {
                    v5[("" + "a")] = 42
                    _ = fi150
                }
                _ = fi149
            }
        case "a":
            switch -9223372036854775808 {
            case -2147483648:
                var v151 bool = false
                _ = v151
            case 402:
                if (true && ((^(v6)) <= v6)) {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                }
            }
        default:
            switch ("a" + (("" + "a") + "\n\t")) {
            case "a":
                if false {
                    var v152 uint32 = uint32(v32)
                    _ = v152
                }
            case "":
                var v153 uint32 = uint32(v16)
                _ = v153
            }
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v129:", v129)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v128
    _ = v129
    _ = v16
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
