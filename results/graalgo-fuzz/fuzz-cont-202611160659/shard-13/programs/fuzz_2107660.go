package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int; F2 uint64 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 string; F1 int32; F2 uint64 }
type GenS4 struct { F0 uint16; F1 uint32; F2 int32 }
type GenS5 struct { F0 int16; F1 int32; F2 string }
type GenS6 struct { F0 int32; F1 uint8 }
type GenS7 struct { F0 uint; F1 string }
type GenS8 struct { F0 uint16; F1 uint32; F2 int }
type GenS9 struct { F0 int8; F1 bool }
type GenS10 struct { F0 uint32; F1 uint64; F2 uint32 }
type GenS11 struct { F0 uint16 }
type GenS12 struct { F0 int16 }
type GenS13 struct { F0 string; F1 int; F2 int16 }
type GenS14 struct { F0 int64 }
type GenS15 struct { F0 uint32; F1 int64; F2 uint }
type GenS16 struct { F0 bool }
type GenS17 struct { F0 int32 }
type GenS18 struct { F0 int32; F1 uint16; F2 int64 }
type GenS19 struct { F0 uint64 }
type GenS20 struct { F0 bool }
type GenS21 struct { F0 int8; F1 uint32 }
type GenS22 struct { F0 int; F1 int }
type GenS23 struct { F0 int8; F1 uint16 }
type GenS24 struct { F0 int8 }
type GenS25 struct { F0 bool }
type GenS26 struct { F0 uint8; F1 int16; F2 uint32 }
type GenS27 struct { F0 uint16; F1 uint32 }
type GenS28 struct { F0 bool; F1 int8; F2 int8 }
type GenS29 struct { F0 int64; F1 uint }
type GenS30 struct { F0 uint16 }
type GenS31 struct { F0 int32; F1 bool }
type GenS32 struct { F0 uint64; F1 string; F2 uint32 }
type GenS33 struct { F0 uint16; F1 int; F2 int32 }
type GenS34 struct { F0 uint64; F1 int16 }
type GenS35 struct { F0 uint16; F1 uint32 }
type GenS36 struct { F0 int; F1 string; F2 int64 }
type GenS37 struct { F0 bool }
type GenS38 struct { F0 uint32; F1 int16; F2 int64 }
type GenS39 struct { F0 uint }
type GenS40 struct { F0 int8 }
type GenS41 struct { F0 int16; F1 string; F2 string }
type GenS42 struct { F0 uint64 }
type GenS43 struct { F0 int8; F1 uint32 }
type GenS44 struct { F0 int; F1 uint32; F2 string }
type GenS45 struct { F0 uint16 }

func main() {
    var v1 int64 = 0
    v2 := []string{"\n\t", ("hello" + "a"), "\n\t", "a"}
    var v3 uint32 = 42
    var v4 bool = (!true)
    for fi5 := 0; fi5 < 13; fi5++ {
        var v6 uint16 = 218
        _ = fi5
        _ = v6
    }
    v7 := GenS1{F0: v3, F1: -885, F2: 4294967295}
    fmt.Println(v1)
    var v8 float64 = ((1e10 / -2.25) / (-2.25 * 1.5))
    for fi9 := 0; fi9 < 19; fi9++ {
        v8 -= (float64(1.5) + (v8 / v8))
        v3 -= v3
        _ = fi9
    }
    _ = len(v2)
    if (-529 > -1) {
        switch -667 {
        case -362:
            if len(v2) > 0 {
                v2[0] = ""
            }
        case 32767:
            if false {
                if (90 > 0) {
                    var v10 int16 = -557
                    for fi11 := 0; fi11 < 4; fi11++ {
                        for fi12 := 0; fi12 < 18; fi12++ {
                            fmt.Println(v1, v3)
                            _ = fi12
                        }
                        _ = fi11
                    }
                    v13 := GenS1{F0: (v3 - v3), F1: 4294967295, F2: 18446744073709551615}
                    _ = v10
                    _ = v13
                }
                if true {
                    if (174 == 42) {
                        fmt.Println(v4, v8)
                        var v14 int64 = -1
                        v15 := GenS2{F0: 507}
                        _ = v14
                        _ = v15
                    }
                    if len(v2) > 0 {
                        v2[0] = (("\n\t" + "hello") + "hello")
                    }
                    var v16 int8 = 127
                    _ = v16
                } else {
                    var v17 bool = false
                    v18 := GenS3{F0: ("foo bar" + "foo bar"), F1: 42, F2: 14}
                    _ = v17
                    _ = v18
                }
                var v19 bool = (v4 && v4)
                _ = v19
            }
        case 2147483647:
            for fi20 := 0; fi20 < 7; fi20++ {
                for fi21 := 0; fi21 < 20; fi21++ {
                    switch "" {
                    case "\n\t":
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    case "foo bar":
                        if (!v4) {
                            var v22 uint32 = v3
                            _ = v22
                        }
                    case "hello":
                        v23 := GenS2{F0: 0}
                        _ = v23
                    default:
                        fmt.Println(v1, fi20)
                    }
                    if ((197 >= 127) && (45 >= 16)) {
                        if len(v2) > 0 {
                            v2[0] = "\n\t"
                        }
                        var v24 bool = (2 >= 525)
                        v3 += (v3 << ((4294967295) & 63))
                        _ = v24
                    }
                    var v25 int = (-(fi20))
                    _ = fi21
                    _ = v25
                }
                v8 += v8
                _ = fi20
            }
        default:
            var v26 int8 = 1
            _ = v26
        }
        if true {
            v8 += 3.14159
        }
    }
    if (!((0 > 65535) || v4)) {
        if (v4 || true) {
            if false {
                fmt.Println(v4, v3)
            } else {
                for fi27 := 0; fi27 < 5; fi27++ {
                    var v28 uint8 = 1
                    fmt.Println(v4, v3)
                    var v29 float64 = v8
                    _ = fi27
                    _ = v28
                    _ = v29
                }
                if v4 {
                    var v30 uint = 1
                    v4 = (-217 < -787)
                    _ = len(v2)
                    _ = v30
                }
            }
            v2 = append(v2, "hello")
        } else {
            for fi31 := 0; fi31 < 20; fi31++ {
                for fi32 := 0; fi32 < 12; fi32++ {
                    var v33 float32 = ((3.14159 - 0.0) / (1.5 / ((3.14159 / (0.0 + float32((v8 / (0.1 + float64(3.14159)))))) / 0.1)))
                    if len(v2) > 0 {
                        v2[0] = (("a" + ("foo bar" + ("foo bar" + ("hello" + "\n\t")))) + (("foo bar" + "foo bar") + "a"))
                    }
                    _ = fi32
                    _ = v33
                }
                v2 = append(v2, "\n\t")
                _ = fi31
            }
        }
        for fi34 := 0; fi34 < 18; fi34++ {
            var v35 uint64 = 895
            _ = fi34
            _ = v35
        }
        var v36 uint32 = (v3 % ((2147483648) | 1))
        _ = v36
    } else {
        fmt.Println(v1)
        var v37 uint64 = 46
        _ = v37
    }
    if (1e10 != (float32(v8) / 1.0)) {
        if true {
            fmt.Println(v1, v3)
            fmt.Println(v4)
        } else {
            v8 = v8
            switch (("hello" + "") + ("foo bar" + "foo bar")) {
            case "\n\t":
                switch "foo bar" {
                case "a":
                    for fi38 := 0; fi38 < 4; fi38++ {
                        if true {
                            _ = len(v2)
                            if true {
                                if ((3.14159 / (0.1 + 0.0)) > (1.5 * (0.1 / (1e10 * float32(v8))))) {
                                    v3 -= v3
                                    for fi39 := 0; fi39 < 3; fi39++ {
                                        v1 = -603
                                        _ = fi39
                                    }
                                    if v4 {
                                        for fi40 := 0; fi40 < 10; fi40++ {
                                            var v41 float64 = v8
                                            v3 -= (v3 / ((v3) | 1))
                                            _ = fi40
                                            _ = v41
                                        }
                                    } else {
                                        var v42 uint8 = 125
                                        _ = v42
                                    }
                                } else {
                                    var v43 int32 = 42
                                    v44 := GenS2{F0: 1}
                                    if v4 {
                                        v2 = append(v2, "a")
                                        var v45 int64 = v1
                                        switch (("foo bar" + ("" + (("hello" + "foo bar") + ((("foo bar" + "a") + ("\n\t" + (("foo bar" + "hello") + "hello"))) + "foo bar")))) + "\n\t") {
                                        case "":
                                            v46 := GenS4{F0: 0, F1: v3, F2: int32(v1)}
                                            _ = v46
                                        case "hello":
                                            v4 = ((((!false) || false) && (true || true)) || (!(true && v4)))
                                        }
                                        _ = v45
                                    } else {
                                        if ((-600 > 1000) || (v4 && (1 > 65535))) {
                                            if true {
                                                switch (fi38 >> ((564) & 63)) {
                                                case 42:
                                                    fmt.Println(v43)
                                                case 4294967295:
                                                    v43 += 1
                                                case 783:
                                                    v47 := GenS5{F0: -1, F1: (v43 << ((1) & 63)), F2: "hello"}
                                                    _ = v47
                                                default:
                                                    var v48 int16 = -32768
                                                    _ = v48
                                                }
                                                for fi49 := 0; fi49 < 15; fi49++ {
                                                    if v4 {
                                                        switch fi49 {
                                                        case 9223372036854775807:
                                                            var v50 int8 = 100
                                                            _ = v50
                                                        case -2147483648:
                                                            if (float64(0.0) == ((((3.14159 / (v8 / v8)) + (v8 - float64(1.5))) - v8) - v8)) {
                                                                v1 = (v1 / (((^(v1))) | 1))
                                                                var v51 string = ""
                                                                _ = v51
                                                            } else {
                                                                v52 := GenS6{F0: (v43 << ((518) & 63)), F1: 127}
                                                                fmt.Println(fi49, v3, fi49)
                                                                if (76 != 42) {
                                                                    v43 = (v43 << ((410) & 63))
                                                                } else {
                                                                    fmt.Println(v4, v4)
                                                                }
                                                                _ = v52
                                                            }
                                                        }
                                                        if ((v4 || true) || v4) {
                                                            v53 := GenS7{F0: 18446744073709551615, F1: (("a" + "hello") + "hello")}
                                                            switch "hello" {
                                                            case "foo bar":
                                                                if (v4 || (1 >= 0)) {
                                                                    var v54 uint = 427
                                                                    var v55 uint8 = 255
                                                                    var v56 uint16 = 0
                                                                    _ = v54
                                                                    _ = v55
                                                                    _ = v56
                                                                }
                                                            default:
                                                                if v4 {
                                                                    v1 += 0
                                                                }
                                                            }
                                                            _ = v53
                                                        }
                                                    }
                                                    v3 = (v3 / ((v3) | 1))
                                                    v1 = (v1 << ((4294967295) & 63))
                                                    _ = fi49
                                                }
                                                for fi57 := 0; fi57 < 3; fi57++ {
                                                    if (v4 && (95 <= 94)) {
                                                        if (935 > 947) {
                                                            v8 -= (float64(0.1) * v8)
                                                            v4 = v4
                                                        } else {
                                                            v58 := GenS1{F0: v3, F1: 32768, F2: uint64(v3)}
                                                            _ = v58
                                                        }
                                                        fmt.Println(v4, fi57)
                                                        switch (((("hello" + ("a" + "")) + "\n\t") + "hello") + "hello") {
                                                        case "hello":
                                                            var v59 float32 = (float32(float64(1.5)) + 0.0)
                                                            _ = v59
                                                        case "a":
                                                            fmt.Println(v4, v4, fi57)
                                                            var v60 string = "\n\t"
                                                            fmt.Println(v43, fi57, fi57)
                                                            _ = v60
                                                        case "\n\t":
                                                            fmt.Println(v43, v43)
                                                        default:
                                                            v43 -= v43
                                                        }
                                                    } else {
                                                        _ = len(v2)
                                                    }
                                                    if (-1 >= 42) {
                                                        v61 := GenS8{F0: 65535, F1: v3, F2: fi57}
                                                        if (!false) {
                                                            var v62 float32 = ((float32(((v8 * v8) * (v8 - (1.5 - v8)))) - 1.5) / 1.0)
                                                            var v63 float64 = v8
                                                            if (true && v4) {
                                                                v2 = append(v2, "\n\t")
                                                                if (128 > 128) {
                                                                    switch fi57 {
                                                                    case 0:
                                                                        if (2 != 2) {
                                                                            var v64 uint64 = 0
                                                                            _ = v64
                                                                        }
                                                                    case -9223372036854775808:
                                                                        v63 += v63
                                                                    default:
                                                                        var v65 int64 = 1
                                                                        _ = v65
                                                                    }
                                                                    v8 -= 0.1
                                                                    if len(v2) > 0 {
                                                                        v2[0] = ((((("foo bar" + "foo bar") + ("hello" + ("foo bar" + "foo bar"))) + ((("\n\t" + ("hello" + "")) + "foo bar") + "foo bar")) + ("hello" + "foo bar")) + "foo bar")
                                                                    }
                                                                }
                                                                fmt.Println(v63)
                                                            } else {
                                                                v1 -= v1
                                                            }
                                                            _ = v62
                                                            _ = v63
                                                        }
                                                        _ = v61
                                                    }
                                                    if v4 {
                                                        if len(v2) >= 2 {
                                                            v2 = v2[1:len(v2)]
                                                        }
                                                        v66 := GenS8{F0: 495, F1: (v3 >> ((724) & 63)), F2: fi57}
                                                        v4 = (("hello" <= ("a" + "\n\t")) && (47 != 101))
                                                        _ = v66
                                                    }
                                                    _ = fi57
                                                }
                                            } else {
                                                _ = len(v2)
                                                if ((128 != 2) && ((!(v4 || (fi38 < 0))) && false)) {
                                                    if len(v2) > 0 {
                                                        v2[0] = "a"
                                                    }
                                                } else {
                                                    var v67 float32 = 1.5
                                                    for fi68 := 0; fi68 < 12; fi68++ {
                                                        fmt.Println(v43, v3, v1)
                                                        var v69 int64 = v1
                                                        var v70 float32 = float32(v8)
                                                        _ = fi68
                                                        _ = v69
                                                        _ = v70
                                                    }
                                                    _ = v67
                                                }
                                            }
                                            v71 := GenS5{F0: -823, F1: (v43 - v43), F2: "a"}
                                            v72 := GenS1{F0: (v3 * 42), F1: fi38, F2: 0}
                                            _ = v71
                                            _ = v72
                                        } else {
                                            v8 += v8
                                            v3 = v3
                                            v3 = v3
                                        }
                                        switch (fi38 / (((fi38 % (((fi38 / (((fi38 / (((fi38 % ((fi38) | 1))) | 1))) | 1))) | 1))) | 1)) {
                                        case -1:
                                            for fi73 := 0; fi73 < 5; fi73++ {
                                                if (0 < 719) {
                                                    _ = len(v2)
                                                }
                                                fmt.Println(v4, v8, fi73)
                                                var v74 int32 = 333
                                                v1 -= 1
                                                v75 := GenS9{F0: 1, F1: (1000 == 0)}
                                                _ = fi73
                                                _ = v74
                                                _ = v75
                                            }
                                        case 127:
                                            for fi76 := 0; fi76 < 5; fi76++ {
                                                fmt.Println(v8, v3)
                                                fmt.Println(fi38)
                                                _ = fi76
                                            }
                                        case -2147483648:
                                            fmt.Println(v1, v1, v3)
                                        }
                                        fmt.Println(fi38)
                                    }
                                    _ = v43
                                    _ = v44
                                }
                            }
                            fmt.Println(fi38, v1, v4)
                        }
                        _ = fi38
                    }
                case "":
                    v2 = append(v2, "hello")
                }
            case "":
                v4 = ((v1 * (^(v1))) < 0)
            case "foo bar":
                v77 := GenS6{F0: 2147483647, F1: 0}
                _ = v77
            }
            for fi78 := 0; fi78 < 4; fi78++ {
                if v4 {
                    for fi79 := 0; fi79 < 15; fi79++ {
                        v80 := GenS10{F0: 547, F1: 418, F2: v3}
                        _ = fi79
                        _ = v80
                    }
                    _ = len(v2)
                    for fi81 := 0; fi81 < 17; fi81++ {
                        var v82 int8 = 42
                        var v83 float32 = float32(v8)
                        if false {
                            if len(v2) > 0 {
                                v2[0] = ("a" + "hello")
                            }
                            if ("\n\t" > "") {
                                var v84 int64 = int64(v82)
                                if false {
                                    var v85 int64 = v1
                                    if false {
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                        var v86 int64 = v84
                                        var v87 float32 = 0.0
                                        if (65535 <= 1) {
                                            _ = len(v2)
                                            v8 = v8
                                            v3 -= v3
                                        } else {
                                            v88 := GenS5{F0: 1, F1: -2147483648, F2: ""}
                                            v87 += v87
                                            if v4 {
                                                v89 := GenS8{F0: 42, F1: 0, F2: (fi81 % (((fi81 * (-(fi78)))) | 1))}
                                                v90 := GenS9{F0: int8(fi81), F1: false}
                                                switch ("a" + "hello") {
                                                case "":
                                                    v91 := GenS7{F0: 255, F1: ("foo bar" + "hello")}
                                                    _ = v91
                                                case "\n\t":
                                                    var v92 uint64 = 1
                                                    v1 -= (v86 >> ((2) & 63))
                                                    if v4 {
                                                        if len(v2) >= 2 {
                                                            v2 = v2[1:len(v2)]
                                                        }
                                                        var v93 uint16 = 255
                                                        if len(v2) >= 2 {
                                                            v2 = v2[1:len(v2)]
                                                        }
                                                        _ = v93
                                                    } else {
                                                        fmt.Println(v85)
                                                        v94 := GenS8{F0: 203, F1: v3, F2: (-(fi78))}
                                                        _ = v94
                                                    }
                                                    _ = v92
                                                case "foo bar":
                                                    var v95 int32 = 0
                                                    _ = v95
                                                default:
                                                    var v96 int8 = int8(fi81)
                                                    _ = v96
                                                }
                                                v8 += (v8 * (0.1 * float64(((((v83 / 1e10) / (((3.14159 + (float32(3.14159) * 1e10)) - (0.0 - (v83 - 0.1))) / float32(-2.25))) / v83) * 3.14159))))
                                                fmt.Println(fi78, v86)
                                                _ = v89
                                                _ = v90
                                            }
                                            if (!(2 >= 2)) {
                                                fmt.Println(v87, v8)
                                                switch fi81 {
                                                case -79:
                                                    var v97 int32 = -1
                                                    fmt.Println(v85, v8, v84)
                                                    _ = v97
                                                case 32768:
                                                    var v98 int8 = (v82 % ((69) | 1))
                                                    _ = v98
                                                case 2147483647:
                                                    switch (^(fi81)) {
                                                    case -1:
                                                        var v99 float64 = ((0.1 * float64((-2.25 + v87))) + ((v8 / 0.1) * 0.1))
                                                        _ = v99
                                                    default:
                                                        if len(v2) >= 2 {
                                                            v2 = v2[1:len(v2)]
                                                        }
                                                    }
                                                }
                                            } else {
                                                _ = len(v2)
                                            }
                                            switch ("a" + "hello") {
                                            case "":
                                                v83 = 1e10
                                            case "a":
                                                _ = len(v2)
                                            }
                                            v100 := GenS11{F0: 1}
                                            var v101 float64 = v8
                                            _ = v100
                                            _ = v101
                                            _ = v88
                                        }
                                        v102 := GenS7{F0: 255, F1: "\n\t"}
                                        if (true && (0 < 1000)) {
                                            switch "foo bar" {
                                            case "hello":
                                                var v103 float64 = float64(((v87 - v83) + v87))
                                                if true {
                                                    v103 -= v8
                                                    if ((^(v82)) < v82) {
                                                        var v104 int32 = -16
                                                        if v4 {
                                                            switch ("foo bar" + ("hello" + ("" + "\n\t"))) {
                                                            case "hello":
                                                                if len(v2) > 0 {
                                                                    v2[0] = ""
                                                                }
                                                            case "foo bar":
                                                                var v105 uint64 = 613
                                                                _ = v105
                                                            }
                                                            v2 = append(v2, "foo bar")
                                                        } else {
                                                            v2 = append(v2, "foo bar")
                                                        }
                                                        _ = v104
                                                    }
                                                }
                                                fmt.Println(v103)
                                                _ = v103
                                            case "foo bar":
                                                v106 := GenS6{F0: -644, F1: 127}
                                                _ = v106
                                            case "":
                                                v107 := GenS12{F0: -1}
                                                _ = v107
                                            default:
                                                v86 -= v84
                                                v1 += -307
                                            }
                                            fmt.Println(v82)
                                        }
                                        _ = v102
                                        _ = v86
                                        _ = v87
                                    } else {
                                        var v108 string = (("hello" + ("\n\t" + "a")) + ("" + "hello"))
                                        _ = len(v2)
                                        var v109 float32 = 1.5
                                        _ = v108
                                        _ = v109
                                    }
                                    _ = v85
                                }
                                _ = v84
                            } else {
                                v83 += v83
                                if v4 {
                                    var v110 uint32 = v3
                                    v8 += v8
                                    _ = v110
                                } else {
                                    var v111 int64 = v1
                                    _ = v111
                                }
                                v83 = v83
                            }
                        }
                        switch (("" + ("foo bar" + "\n\t")) + ("" + ("hello" + "a"))) {
                        case "foo bar":
                            v82 = 42
                        }
                        _ = fi81
                        _ = v82
                        _ = v83
                    }
                } else {
                    var v112 float64 = 1.5
                    v4 = true
                    _ = v112
                }
                _ = fi78
            }
        }
        switch -1 {
        case 42:
            v113 := GenS13{F0: "a", F1: -1, F2: -1}
            _ = v113
        case 841:
            for fi114 := 0; fi114 < 18; fi114++ {
                _ = len(v2)
                var v115 string = "foo bar"
                v116 := GenS9{F0: 100, F1: false}
                _ = fi114
                _ = v115
                _ = v116
            }
        default:
            var v117 string = "foo bar"
            _ = v117
        }
    } else {
        if v4 {
            switch 32767 {
            case -890:
                if len(v2) > 0 {
                    v2[0] = "a"
                }
            case 32767:
                if (!v4) {
                    var v118 float32 = -2.25
                    var v119 int32 = -2147483648
                    switch "foo bar" {
                    case "":
                        v8 = ((1.5 + float64(-2.25)) / (((3.14159 / 3.14159) - float64(0.1)) * v8))
                    case "\n\t":
                        v2 = append(v2, "hello")
                    case "hello":
                        v4 = (113 > 42)
                    default:
                        if v4 {
                            fmt.Println(v3, v8, v4)
                            v2 = append(v2, "foo bar")
                            for fi120 := 0; fi120 < 9; fi120++ {
                                for fi121 := 0; fi121 < 14; fi121++ {
                                    if ((v4 && (4294967295 < 0)) || false) {
                                        v118 = -2.25
                                        if true {
                                            v118 = -2.25
                                        } else {
                                            v122 := GenS2{F0: 705}
                                            var v123 bool = false
                                            _ = v122
                                            _ = v123
                                        }
                                        v1 = (v1 << ((18446744073709551615) & 63))
                                    }
                                    if ((-32768 > int16(v1)) || (2 != 96)) {
                                        _ = len(v2)
                                        v124 := GenS14{F0: (v1 & (^(v1)))}
                                        _ = v124
                                    } else {
                                        switch fi120 {
                                        case -513:
                                            var v125 uint = 42
                                            _ = v125
                                        case 0:
                                            var v126 int16 = 42
                                            _ = v126
                                        case -9223372036854775808:
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                        default:
                                            if true {
                                                if v4 {
                                                    v127 := GenS6{F0: v119, F1: 2}
                                                    v3 += 65535
                                                    if (int(v1) <= fi120) {
                                                        v1 -= v1
                                                        var v128 bool = (1 != 42)
                                                        var v129 int8 = int8(v119)
                                                        _ = v128
                                                        _ = v129
                                                    }
                                                    _ = v127
                                                }
                                                if ((!(32767 == 0)) || (0 == 251)) {
                                                    v2 = append(v2, "\n\t")
                                                    v8 -= v8
                                                } else {
                                                    v119 = (v119 >> ((42) & 63))
                                                    v1 = (v1 / ((v1) | 1))
                                                    var v130 uint = 0
                                                    _ = v130
                                                }
                                                v2 = append(v2, "hello")
                                            } else {
                                                var v131 int = (fi121 >> ((939) & 63))
                                                v132 := GenS14{F0: v1}
                                                _ = v131
                                                _ = v132
                                            }
                                        }
                                        v8 = -2.25
                                        var v133 uint16 = 255
                                        _ = v133
                                    }
                                    _ = fi121
                                }
                                v134 := GenS3{F0: "", F1: (v119 | 505), F2: 1}
                                for fi135 := 0; fi135 < 17; fi135++ {
                                    if v4 {
                                        v4 = true
                                        if v4 {
                                            v3 -= (v3 << ((23) & 63))
                                        }
                                        switch 127 {
                                        case 32768:
                                            fmt.Println(v118, v4)
                                        default:
                                            if (float32(1e10) != 3.14159) {
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                                v136 := GenS11{F0: 0}
                                                switch (^(fi135)) {
                                                case 1:
                                                    if (!((((v8 / (v8 + v8)) < v8) && (-9223372036854775808 >= (v1 % ((-9223372036854775808) | 1)))) || (!(228 <= 0)))) {
                                                        v4 = v4
                                                        if v4 {
                                                            v137 := GenS7{F0: uint(v3), F1: "foo bar"}
                                                            fmt.Println(fi135)
                                                            v118 = 1e10
                                                            _ = v137
                                                        } else {
                                                            if (!v4) {
                                                                var v138 int32 = (v119 * v119)
                                                                _ = v138
                                                            }
                                                        }
                                                        switch ("" + ("foo bar" + "")) {
                                                        case "a":
                                                            fmt.Println(fi120, v119, v1)
                                                        case "":
                                                            v4 = v4
                                                        case "foo bar":
                                                            v3 += v3
                                                        default:
                                                            switch fi135 {
                                                            case 2147483647:
                                                                if v4 {
                                                                    var v139 float64 = ((v8 * v8) + 3.14159)
                                                                    v4 = (!("foo bar" <= "hello"))
                                                                    _ = v139
                                                                }
                                                            default:
                                                                var v140 uint32 = v3
                                                                _ = v140
                                                            }
                                                        }
                                                    }
                                                case 42:
                                                    switch "hello" {
                                                    case "\n\t":
                                                        fmt.Println(v118, v118, v1)
                                                    }
                                                }
                                                _ = v136
                                            } else {
                                                var v141 int16 = -577
                                                fmt.Println(v3)
                                                _ = v141
                                            }
                                        }
                                    }
                                    var v142 uint32 = v3
                                    _ = fi135
                                    _ = v142
                                }
                                _ = fi120
                                _ = v134
                            }
                        } else {
                            if v4 {
                                var v143 uint8 = 128
                                var v144 int32 = v119
                                _ = len(v2)
                                _ = v143
                                _ = v144
                            }
                            v1 = v1
                        }
                    }
                    _ = v118
                    _ = v119
                }
            default:
                for fi145 := 0; fi145 < 18; fi145++ {
                    for fi146 := 0; fi146 < 6; fi146++ {
                        v8 += 1e10
                        v147 := GenS7{F0: 932, F1: "hello"}
                        switch (fi145 & 9223372036854775807) {
                        case -2147483648:
                            v2 = append(v2, ("" + "a"))
                        case -708:
                            v148 := GenS11{F0: 237}
                            v149 := GenS11{F0: 0}
                            v150 := GenS15{F0: 819, F1: (v1 & (v1 / ((604) | 1))), F2: 46}
                            if false {
                                switch 42 {
                                case 2147483647:
                                    fmt.Println(v1, fi146, fi146)
                                case 0:
                                    fmt.Println(v3, v1)
                                case -1:
                                    fmt.Println(v1)
                                }
                                v2 = append(v2, "\n\t")
                                v3 -= (v3 & (v3 & (v3 >> ((164) & 63))))
                            } else {
                                if false {
                                    _ = len(v2)
                                    var v151 int32 = 2147483647
                                    _ = len(v2)
                                    _ = v151
                                } else {
                                    v3 = v3
                                }
                                if (v4 || (!v4)) {
                                    v1 = 1
                                } else {
                                    switch "hello" {
                                    case "foo bar":
                                        v4 = true
                                    case "":
                                        var v152 int8 = -1
                                        _ = v152
                                    case "a":
                                        v2 = append(v2, "\n\t")
                                    default:
                                        if len(v2) > 0 {
                                            v2[0] = ("foo bar" + "hello")
                                        }
                                    }
                                    if (378 <= 2) {
                                        switch fi146 {
                                        case -1:
                                            fmt.Println(v4, v1)
                                        case 0:
                                            v1 = 9223372036854775807
                                        case -9223372036854775808:
                                            v153 := GenS2{F0: uint(v3)}
                                            _ = v153
                                        }
                                        _ = len(v2)
                                        var v154 uint64 = 0
                                        _ = v154
                                    } else {
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                    }
                                }
                            }
                            _ = v148
                            _ = v149
                            _ = v150
                        default:
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        }
                        _ = fi146
                        _ = v147
                    }
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    if v4 {
                        for fi155 := 0; fi155 < 2; fi155++ {
                            var v156 int = fi145
                            _ = len(v2)
                            var v157 int32 = 42
                            v157 -= -622
                            v158 := GenS2{F0: 918}
                            switch (^(fi155)) {
                            case -2147483648:
                                switch ("" + ("a" + ("\n\t" + (("" + "hello") + "")))) {
                                case "hello":
                                    if len(v2) > 0 {
                                        v2[0] = "hello"
                                    }
                                default:
                                    v159 := GenS16{F0: true}
                                    _ = v159
                                }
                            case 42:
                                v157 -= -233
                            case 128:
                                if false {
                                    _ = len(v2)
                                    v4 = v4
                                }
                            default:
                                v156 = (^(fi155))
                            }
                            _ = fi155
                            _ = v156
                            _ = v157
                            _ = v158
                        }
                        if true {
                            switch fi145 {
                            case 128:
                                var v160 string = "\n\t"
                                _ = v160
                            case -877:
                                switch (("a" + ("" + "")) + ("\n\t" + "foo bar")) {
                                case "foo bar":
                                    switch fi145 {
                                    case -1:
                                        var v161 string = "\n\t"
                                        _ = v161
                                    }
                                }
                            default:
                                var v162 string = ("hello" + "")
                                _ = v162
                            }
                            v1 += v1
                            v8 += 1.5
                        }
                    }
                    _ = fi145
                }
            }
            v163 := GenS17{F0: 42}
            for fi164 := 0; fi164 < 14; fi164++ {
                var v165 uint64 = 2
                if true {
                    if len(v2) > 0 {
                        v2[0] = "foo bar"
                    }
                }
                v165 -= v165
                _ = fi164
                _ = v165
            }
            _ = v163
        } else {
            v166 := GenS18{F0: 0, F1: 2, F2: v1}
            var v167 string = (("a" + "a") + (("foo bar" + ("hello" + "hello")) + "a"))
            _ = v166
            _ = v167
        }
        var v168 int32 = 1
        _ = v168
    }
    switch "foo bar" {
    case "hello":
        for fi169 := 0; fi169 < 14; fi169++ {
            var v170 uint = 2
            if true {
                v171 := GenS8{F0: 65535, F1: (v3 % ((v3) | 1)), F2: 128}
                for fi172 := 0; fi172 < 15; fi172++ {
                    if ((!((v3 | 2) > 666)) || (false && v4)) {
                        fmt.Println(v1)
                        v170 -= (v170 >> (((v170 + v170)) & 63))
                        var v173 float32 = 1e10
                        v170 = (v170 << (((v170 / ((v170) | 1))) & 63))
                        _ = v173
                    } else {
                        if v4 {
                            var v174 int = (fi169 + (fi172 / ((42) | 1)))
                            _ = v174
                        }
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    }
                    _ = fi172
                }
                _ = v171
            } else {
                v175 := GenS19{F0: 547}
                v1 -= (v1 & (v1 | v1))
                _ = v175
            }
            if v4 {
                v176 := GenS20{F0: (!(250 == (v3 - 65535)))}
                v4 = false
                _ = v176
            } else {
                if true {
                    for fi177 := 0; fi177 < 6; fi177++ {
                        var v178 int64 = (-(v1))
                        if v4 {
                            if false {
                                switch "" {
                                case "hello":
                                    v1 = (^(v1))
                                }
                                v178 = v178
                                if false {
                                    v179 := GenS14{F0: v178}
                                    var v180 int = fi177
                                    v178 = (v178 / ((v178) | 1))
                                    v181 := GenS21{F0: 42, F1: 1}
                                    _ = v179
                                    _ = v180
                                    _ = v181
                                } else {
                                    if (!((0.1 + ((1e10 / float32(v8)) + (float32(1.5) / 1.5))) != 0.0)) {
                                        v182 := GenS22{F0: -548, F1: (fi169 * -174)}
                                        v3 += 526
                                        v8 = float64((-2.25 / -2.25))
                                        _ = v182
                                    } else {
                                        var v183 uint8 = 2
                                        _ = v183
                                    }
                                }
                            } else {
                                if false {
                                    if (((1e10 + 0.1) < v8) && v4) {
                                        if (!(2 <= 1)) {
                                            v3 = (^(v3))
                                            v178 = v1
                                        } else {
                                            fmt.Println(fi169)
                                            if len(v2) > 0 {
                                                v2[0] = "hello"
                                            }
                                        }
                                        if v4 {
                                            v184 := GenS23{F0: 42, F1: 42}
                                            var v185 int64 = (-(v178))
                                            _ = v184
                                            _ = v185
                                        } else {
                                            v4 = v4
                                            v1 = 9223372036854775807
                                        }
                                        v178 = (v178 % ((635) | 1))
                                    } else {
                                        var v186 float32 = (1e10 + 1.5)
                                        _ = v186
                                    }
                                }
                                if v4 {
                                    var v187 int16 = 32767
                                    if v4 {
                                        var v188 uint64 = 42
                                        if (-32768 <= v187) {
                                            if ((v187 | -32768) < 42) {
                                                var v189 uint = (^(v170))
                                                _ = v189
                                            }
                                            fmt.Println(v3, fi177)
                                            if len(v2) > 0 {
                                                v2[0] = "a"
                                            }
                                        }
                                        v1 = (v1 ^ v178)
                                        _ = v188
                                    }
                                    _ = v187
                                }
                                if len(v2) > 0 {
                                    v2[0] = "\n\t"
                                }
                                v178 = (^(v1))
                                fmt.Println(v1, v3)
                            }
                            var v190 uint64 = 934
                            if (583 <= 171) {
                                v2 = append(v2, ("a" + "foo bar"))
                                switch (fi177 << (((^(v170))) & 63)) {
                                case 2147483647:
                                    if v4 {
                                        fmt.Println(v4, fi177)
                                        v4 = true
                                    }
                                case 32768:
                                    fmt.Println(v8, v4, v1)
                                }
                                var v191 float32 = -2.25
                                _ = v191
                            } else {
                                var v192 int64 = (v178 * (v178 % (((v178 | (-(v1)))) | 1)))
                                if v4 {
                                    var v193 uint = (v170 * 42)
                                    v4 = v4
                                    if len(v2) > 0 {
                                        v2[0] = ""
                                    }
                                    _ = v193
                                }
                                if v4 {
                                    v2 = append(v2, "foo bar")
                                } else {
                                    v194 := GenS18{F0: -787, F1: 200, F2: (v1 % (((v192 / ((v1) | 1))) | 1))}
                                    switch ("" + "") {
                                    case "\n\t":
                                        if (1 <= 128) {
                                            if (fi177 == fi169) {
                                                v195 := GenS24{F0: -103}
                                                _ = v195
                                            } else {
                                                _ = len(v2)
                                            }
                                        }
                                    case "foo bar":
                                        if false {
                                            switch "a" {
                                            case "foo bar":
                                                var v196 float64 = ((3.14159 / -2.25) + ((0.0 + v8) + (1e10 * v8)))
                                                _ = v196
                                            case "\n\t":
                                                v197 := GenS7{F0: v170, F1: ("hello" + "a")}
                                                if (v4 && (-219 >= (fi169 % (((fi177 ^ int(v192))) | 1)))) {
                                                    switch ("foo bar" + "foo bar") {
                                                    case "hello":
                                                        if true {
                                                            if (!(!(11 < 0))) {
                                                                v1 = (v178 % ((187) | 1))
                                                            } else {
                                                                v192 += v178
                                                                v192 += v1
                                                                fmt.Println(v170)
                                                                fmt.Println(fi169, v178)
                                                            }
                                                        }
                                                    case "":
                                                        v170 = (v170 ^ 972)
                                                    default:
                                                        if v4 {
                                                            if v4 {
                                                                v198 := GenS25{F0: v4}
                                                                _ = v198
                                                            }
                                                            var v199 uint8 = 0
                                                            var v200 uint16 = 65535
                                                            _ = v199
                                                            _ = v200
                                                        }
                                                    }
                                                    if ((^(v170)) >= (v170 ^ uint(v190))) {
                                                        var v201 int16 = 1000
                                                        v202 := GenS26{F0: 1, F1: v201, F2: 65535}
                                                        fmt.Println(v190, v8, v201)
                                                        _ = v201
                                                        _ = v202
                                                    } else {
                                                        var v203 int32 = -2147483648
                                                        var v204 uint8 = 196
                                                        switch "\n\t" {
                                                        case "a":
                                                            switch "\n\t" {
                                                            case "a":
                                                                if v4 {
                                                                    _ = len(v2)
                                                                    if false {
                                                                        v203 += v203
                                                                        switch "a" {
                                                                        case "a":
                                                                            var v205 int16 = -617
                                                                            _ = v205
                                                                        case "hello":
                                                                            if v4 {
                                                                                v2 = append(v2, ("hello" + "a"))
                                                                                v206 := GenS13{F0: "hello", F1: int(v190), F2: 1000}
                                                                                v207 := GenS12{F0: 1000}
                                                                                _ = v206
                                                                                _ = v207
                                                                            } else {
                                                                                if false {
                                                                                    var v208 float64 = v8
                                                                                    var v209 int8 = 100
                                                                                    _ = v208
                                                                                    _ = v209
                                                                                } else {
                                                                                    _ = len(v2)
                                                                                }
                                                                                _ = len(v2)
                                                                                var v210 float64 = 1e10
                                                                                _ = v210
                                                                            }
                                                                        default:
                                                                            fmt.Println(v8)
                                                                        }
                                                                        v211 := GenS5{F0: 0, F1: (v203 << ((v170) & 63)), F2: "\n\t"}
                                                                        _ = v211
                                                                    }
                                                                    if false {
                                                                        v212 := GenS27{F0: uint16(v190), F1: (^(v3))}
                                                                        fmt.Println(v1)
                                                                        v190 = (v190 | v190)
                                                                        _ = v212
                                                                    }
                                                                } else {
                                                                    var v213 bool = (false || (!(true || true)))
                                                                    v190 = 890
                                                                    _ = v213
                                                                }
                                                            case "foo bar":
                                                                switch "foo bar" {
                                                                case "":
                                                                    switch fi169 {
                                                                    case -1:
                                                                        v214 := GenS28{F0: v4, F1: -128, F2: 110}
                                                                        _ = v214
                                                                    case -656:
                                                                        var v215 uint = (v170 >> ((v170) & 63))
                                                                        _ = v215
                                                                    default:
                                                                        if v4 {
                                                                            v216 := GenS29{F0: v192, F1: (v170 + v170)}
                                                                            var v217 uint32 = v3
                                                                            _ = v216
                                                                            _ = v217
                                                                        }
                                                                    }
                                                                case "a":
                                                                    if (-2.25 == v8) {
                                                                        if (v4 || ((v203 >> ((570) & 63)) < (v203 | -1))) {
                                                                            var v218 int32 = v203
                                                                            _ = v218
                                                                        } else {
                                                                            var v219 int8 = 42
                                                                            switch (^(fi169)) {
                                                                            case 32768:
                                                                                fmt.Println(v8, v219)
                                                                            case -1:
                                                                                fmt.Println(v204)
                                                                            default:
                                                                                if false {
                                                                                    var v220 int8 = v219
                                                                                    _ = v220
                                                                                } else {
                                                                                    var v221 uint64 = v190
                                                                                    _ = v221
                                                                                }
                                                                            }
                                                                            var v222 uint16 = 32767
                                                                            if v4 {
                                                                                v223 := GenS30{F0: 0}
                                                                                v224 := GenS23{F0: (-(v219)), F1: 667}
                                                                                _ = v223
                                                                                _ = v224
                                                                            } else {
                                                                                v225 := GenS31{F0: 606, F1: ((1e10 * v8) < (float64(0.0) - v8))}
                                                                                switch "hello" {
                                                                                case "a":
                                                                                    v1 += (^(v1))
                                                                                case "\n\t":
                                                                                    fmt.Println(fi177)
                                                                                case "":
                                                                                    var v226 int16 = -32768
                                                                                    _ = v226
                                                                                default:
                                                                                    if ((v8 + (1e10 - 3.14159)) < v8) {
                                                                                        var v227 uint64 = (v190 * (v190 | v190))
                                                                                        _ = v227
                                                                                    }
                                                                                }
                                                                                var v228 uint8 = (v204 | (v204 % ((v204) | 1)))
                                                                                _ = v225
                                                                                _ = v228
                                                                            }
                                                                            _ = v219
                                                                            _ = v222
                                                                        }
                                                                        if len(v2) >= 2 {
                                                                            v2 = v2[1:len(v2)]
                                                                        }
                                                                        switch 65 {
                                                                        case 4294967295:
                                                                            _ = len(v2)
                                                                        }
                                                                    } else {
                                                                        v229 := GenS17{F0: int32(v170)}
                                                                        v170 = 4294967295
                                                                        v8 -= float64(1.5)
                                                                        _ = v229
                                                                    }
                                                                default:
                                                                    if (int8(v170) <= -128) {
                                                                        var v230 uint16 = 575
                                                                        _ = v230
                                                                    }
                                                                }
                                                            }
                                                        case "foo bar":
                                                            if false {
                                                                v190 -= v190
                                                                if (!v4) {
                                                                    fmt.Println(v3, fi169, v203)
                                                                    switch ("" + "hello") {
                                                                    case "foo bar":
                                                                        _ = len(v2)
                                                                    default:
                                                                        switch "a" {
                                                                        case "foo bar":
                                                                            v231 := GenS17{F0: 1}
                                                                            _ = v231
                                                                        case "":
                                                                            switch (fi169 >> (((v170 + v170)) & 63)) {
                                                                            case -9223372036854775808:
                                                                                if len(v2) > 0 {
                                                                                    v2[0] = ("hello" + (("\n\t" + "a") + ("" + "a")))
                                                                                }
                                                                            default:
                                                                                if v4 {
                                                                                    v2 = append(v2, "")
                                                                                    fmt.Println(v4, v204, v1)
                                                                                    v232 := GenS32{F0: 18446744073709551615, F1: ("a" + (((((("\n\t" + "\n\t") + "\n\t") + "\n\t") + "a") + "") + ((("\n\t" + "a") + ("\n\t" + ("foo bar" + "\n\t"))) + "foo bar"))), F2: (v3 / ((v3) | 1))}
                                                                                    _ = v232
                                                                                } else {
                                                                                    if ((v170 * (v170 - (v170 % (((^(v170))) | 1)))) < v170) {
                                                                                        v170 += v170
                                                                                        v178 += (v192 / ((v1) | 1))
                                                                                    } else {
                                                                                        v233 := GenS33{F0: 1, F1: fi177, F2: v203}
                                                                                        _ = v233
                                                                                    }
                                                                                    v190 = 1
                                                                                }
                                                                            }
                                                                            fmt.Println(v1)
                                                                            v8 += 1.5
                                                                            var v234 uint8 = v204
                                                                            if len(v2) >= 2 {
                                                                                v2 = v2[1:len(v2)]
                                                                            }
                                                                            _ = v234
                                                                        }
                                                                    }
                                                                    if v4 {
                                                                        if (!true) {
                                                                            var v235 uint16 = 65535
                                                                            _ = v235
                                                                        } else {
                                                                            _ = len(v2)
                                                                        }
                                                                        v203 = (v203 * (v203 * v203))
                                                                        var v236 int64 = (v192 + -9223372036854775808)
                                                                        _ = v236
                                                                    } else {
                                                                        switch (-(fi177)) {
                                                                        case -9223372036854775808:
                                                                            switch "\n\t" {
                                                                            case "\n\t":
                                                                                v3 += 2147483648
                                                                            case "hello":
                                                                                v237 := GenS31{F0: -1, F1: true}
                                                                                _ = v237
                                                                            }
                                                                        default:
                                                                            fmt.Println(v170)
                                                                        }
                                                                        var v238 uint16 = 1
                                                                        var v239 bool = v4
                                                                        _ = v238
                                                                        _ = v239
                                                                    }
                                                                } else {
                                                                    v4 = (1 <= v192)
                                                                    var v240 uint16 = 42
                                                                    _ = v240
                                                                }
                                                                var v241 uint8 = 9
                                                                _ = v241
                                                            } else {
                                                                var v242 int = (fi169 + (fi169 << (((v170 | v170)) & 63)))
                                                                var v243 int = (fi169 ^ int(v190))
                                                                var v244 int8 = 42
                                                                _ = v242
                                                                _ = v243
                                                                _ = v244
                                                            }
                                                        }
                                                        var v245 uint32 = v3
                                                        if v4 {
                                                            v192 += v192
                                                            var v246 uint = v170
                                                            switch (fi177 ^ fi169) {
                                                            case 1:
                                                                var v247 int32 = v203
                                                                _ = v247
                                                            case 128:
                                                                v248 := GenS34{F0: (v190 | v190), F1: -32768}
                                                                var v249 uint8 = (v204 / (((v204 >> ((18446744073709551615) & 63))) | 1))
                                                                v4 = false
                                                                _ = v248
                                                                _ = v249
                                                            }
                                                            _ = v246
                                                        }
                                                        v4 = v4
                                                        if v4 {
                                                            v178 = -9223372036854775808
                                                            v203 = -430
                                                            if (v4 && v4) {
                                                                v245 += v3
                                                                fmt.Println(v203, v8)
                                                            }
                                                        } else {
                                                            var v250 float32 = ((0.0 + (3.14159 / -2.25)) / 3.14159)
                                                            v251 := GenS35{F0: 32767, F1: (^(v245))}
                                                            switch 128 {
                                                            case -9223372036854775808:
                                                                fmt.Println(v8, v203)
                                                            case 42:
                                                                if ((((v250 + float32((v8 * ((v8 / (v8 - 3.14159)) / 0.1)))) <= v250) || ((v8 - 0.0) < float64(float32(v8)))) && v4) {
                                                                    switch 32768 {
                                                                    case 96:
                                                                        var v252 uint = 1
                                                                        _ = v252
                                                                    case 1:
                                                                        v253 := GenS16{F0: v4}
                                                                        _ = v253
                                                                    case 2147483647:
                                                                        var v254 int8 = 122
                                                                        _ = v254
                                                                    }
                                                                    switch ("" + "a") {
                                                                    case "hello":
                                                                        if (!v4) {
                                                                            fmt.Println(fi177, fi177)
                                                                            v178 += (v1 - v178)
                                                                            v204 += (v204 - 1)
                                                                        }
                                                                    case "foo bar":
                                                                        _ = len(v2)
                                                                    }
                                                                    v1 -= (-(v192))
                                                                } else {
                                                                    var v255 int = 32768
                                                                    var v256 float32 = v250
                                                                    _ = v255
                                                                    _ = v256
                                                                }
                                                            case -1:
                                                                v257 := GenS36{F0: int(v190), F1: "", F2: v192}
                                                                _ = v257
                                                            default:
                                                                switch "" {
                                                                case "a":
                                                                    v3 = (v245 & v245)
                                                                case "\n\t":
                                                                    switch (fi177 * -182) {
                                                                    case 682:
                                                                        var v258 int8 = -18
                                                                        _ = v258
                                                                    case 108:
                                                                        var v259 int = fi177
                                                                        _ = v259
                                                                    case 32768:
                                                                        var v260 uint16 = 32767
                                                                        _ = v260
                                                                    }
                                                                    var v261 int32 = v203
                                                                    _ = v261
                                                                default:
                                                                    if len(v2) > 0 {
                                                                        v2[0] = "a"
                                                                    }
                                                                }
                                                            }
                                                            v262 := GenS26{F0: v204, F1: 1000, F2: v3}
                                                            _ = v250
                                                            _ = v251
                                                            _ = v262
                                                        }
                                                        _ = v203
                                                        _ = v204
                                                        _ = v245
                                                    }
                                                    fmt.Println(v4, v192, v178)
                                                } else {
                                                    var v263 uint16 = 42
                                                    _ = v263
                                                }
                                                if false {
                                                    fmt.Println(fi177)
                                                } else {
                                                    _ = len(v2)
                                                    var v264 float32 = (1e10 * (0.0 - 0.0))
                                                    v265 := GenS37{F0: ((v264 + 1.5) != v264)}
                                                    _ = v264
                                                    _ = v265
                                                }
                                                _ = v197
                                            case "hello":
                                                var v266 int32 = -417
                                                _ = v266
                                            default:
                                                _ = len(v2)
                                            }
                                        } else {
                                            v170 = (v170 & (v170 % (((v170 ^ (v170 + v170))) | 1)))
                                            if true {
                                                v3 -= (^(v3))
                                                v267 := GenS38{F0: v3, F1: 540, F2: v178}
                                                _ = v267
                                            }
                                            if ((!((^(v170)) <= v170)) && v4) {
                                                fmt.Println(fi169, fi169, v4)
                                                var v268 int16 = 1000
                                                v4 = (42 != 1)
                                                _ = v268
                                            } else {
                                                v8 = 3.14159
                                                v3 += uint32(v190)
                                                switch ("\n\t" + ("foo bar" + ("\n\t" + ""))) {
                                                case "a":
                                                    fmt.Println(v3, fi177, v3)
                                                case "\n\t":
                                                    var v269 uint64 = (v190 % ((853) | 1))
                                                    _ = v269
                                                default:
                                                    fmt.Println(fi169, v190, v178)
                                                }
                                            }
                                            v270 := GenS4{F0: uint16(v3), F1: (v3 - v3), F2: 1}
                                            if (false || (int(v190) > int(v1))) {
                                                v4 = v4
                                            } else {
                                                switch 1 {
                                                case -689:
                                                    v271 := GenS39{F0: v170}
                                                    _ = v271
                                                default:
                                                    v190 = v190
                                                }
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                            }
                                            _ = v270
                                        }
                                    default:
                                        v272 := GenS40{F0: 0}
                                        v2 = append(v2, "foo bar")
                                        v3 -= (v3 << ((327) & 63))
                                        fmt.Println(v178, v3)
                                        _ = v272
                                    }
                                    _ = v194
                                }
                                _ = v192
                            }
                            _ = v190
                        } else {
                            v178 = v178
                            v170 += 18446744073709551615
                            var v273 int = fi177
                            v4 = (1 <= 32767)
                            v274 := GenS9{F0: 127, F1: v4}
                            v4 = (v4 || (183 != -782))
                            fmt.Println(v1)
                            v2 = append(v2, (("" + ("foo bar" + "foo bar")) + (("" + ("" + ("hello" + "hello"))) + "")))
                            _ = v273
                            _ = v274
                        }
                        _ = fi177
                        _ = v178
                    }
                    for fi275 := 0; fi275 < 17; fi275++ {
                        _ = len(v2)
                        _ = fi275
                    }
                } else {
                    if v4 {
                        v276 := GenS34{F0: 2, F1: 32767}
                        _ = v276
                    }
                }
                if v4 {
                    v277 := GenS14{F0: (v1 & -1)}
                    _ = v277
                } else {
                    for fi278 := 0; fi278 < 10; fi278++ {
                        var v279 float32 = -2.25
                        v3 += v3
                        v2 = append(v2, "\n\t")
                        _ = fi278
                        _ = v279
                    }
                    v3 = 0
                    if true {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        for fi280 := 0; fi280 < 10; fi280++ {
                            v8 += 3.14159
                            v281 := GenS41{F0: -405, F1: "\n\t", F2: "\n\t"}
                            var v282 string = "a"
                            if ((v4 && true) || false) {
                                fmt.Println(v8, v4)
                                v8 -= v8
                                if true {
                                    var v283 int8 = 127
                                    v284 := GenS42{F0: uint64(v170)}
                                    switch v282 {
                                    case "a":
                                        if v4 {
                                            fmt.Println(v8, fi169)
                                            var v285 uint64 = uint64(v170)
                                            var v286 float32 = (0.1 * 1e10)
                                            _ = v285
                                            _ = v286
                                        }
                                    case "foo bar":
                                        if v4 {
                                            switch (fi169 / (((fi169 << ((0) & 63))) | 1)) {
                                            case -877:
                                                v287 := GenS32{F0: 524, F1: v282, F2: 4294967295}
                                                _ = v287
                                            default:
                                                v288 := GenS43{F0: (v283 >> (((v170 + v170)) & 63)), F1: v3}
                                                _ = v288
                                            }
                                            var v289 uint64 = 18446744073709551615
                                            _ = v289
                                        }
                                    default:
                                        v290 := GenS11{F0: 1}
                                        _ = len(v2)
                                        _ = v290
                                    }
                                    _ = v283
                                    _ = v284
                                } else {
                                    var v291 int = fi280
                                    var v292 uint = 178
                                    var v293 bool = true
                                    _ = v291
                                    _ = v292
                                    _ = v293
                                }
                            } else {
                                if (141 > -2147483648) {
                                    switch v282 {
                                    case "":
                                        v4 = true
                                    default:
                                        var v294 uint64 = 18446744073709551615
                                        _ = v294
                                    }
                                } else {
                                    var v295 int = (fi280 & (fi280 / ((fi280) | 1)))
                                    _ = v295
                                }
                            }
                            switch v282 {
                            case "a":
                                v296 := GenS44{F0: 2147483647, F1: 65535, F2: v282}
                                _ = v296
                            case "":
                                v297 := GenS45{F0: 383}
                                _ = v297
                            case "foo bar":
                                var v298 int32 = 274
                                _ = v298
                            default:
                                v3 += v3
                            }
                            v282 = v282
                            _ = fi280
                            _ = v281
                            _ = v282
                        }
                    } else {
                        v8 -= 0.0
                    }
                }
            }
            _ = fi169
            _ = v170
        }
    default:
        if len(v2) > 0 {
            v2[0] = ""
        }
    }
    if v4 {
        if len(v2) > 0 {
            v2[0] = (("" + ("a" + ("foo bar" + ("" + ("\n\t" + "hello"))))) + ("" + "hello"))
        }
        var v299 uint32 = (v3 & (v3 * v3))
        var v300 uint32 = 2
        _ = v299
        _ = v300
    } else {
        var v301 uint = 350
        _ = v301
    }
    var v302 uint = 895
    var v303 int16 = -129
    var v304 int = 60
    switch v304 {
    case 1:
        v305 := GenS12{F0: -32768}
        _ = v305
    case -873:
        fmt.Println(v3, v8)
    case 32767:
        for fi306 := 0; fi306 < 18; fi306++ {
            v307 := GenS19{F0: 1}
            _ = fi306
            _ = v307
        }
    }
    v2 = append(v2, "a")
    if len(v2) > 0 {
        v2[0] = "a"
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v302:", v302)
    fmt.Println("v303:", v303)
    fmt.Println("v304:", v304)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v302
    _ = v303
    _ = v304
    _ = v4
    _ = v7
    _ = v8
}
