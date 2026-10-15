package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int64; F1 uint16; F2 int64 }
type GenS3 struct { F0 uint8; F1 int8; F2 bool }
type GenS4 struct { F0 string; F1 string }
type GenS5 struct { F0 string; F1 uint16; F2 uint64 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 string; F1 uint64 }
type GenS9 struct { F0 int32; F1 uint8 }
type GenS10 struct { F0 string }
type GenS11 struct { F0 uint16 }
type GenS12 struct { F0 uint8; F1 string }
type GenS13 struct { F0 uint }
type GenS14 struct { F0 int; F1 int8; F2 uint32 }
type GenS15 struct { F0 uint16; F1 int8 }
type GenS16 struct { F0 uint; F1 uint64; F2 int8 }
type GenS17 struct { F0 int64; F1 uint32; F2 uint8 }
type GenS18 struct { F0 int8; F1 uint; F2 int8 }
type GenS19 struct { F0 uint64; F1 int; F2 int }
type GenS20 struct { F0 int; F1 uint }
type GenS21 struct { F0 int16; F1 int64; F2 bool }

func main() {
    var v1 float64 = 0.1
    var v2 int8 = 100
    var v3 int16 = -220
    v4 := []string{(("a" + "\n\t") + "hello"), "a", ""}
    v5 := GenS1{F0: 535}
    fmt.Println(v2, v1, v1)
    if true {
        v2 = (v2 >> ((2) & 63))
        var v6 uint = 2
        var v7 uint8 = 238
        _ = v6
        _ = v7
    }
    if (-13 >= v2) {
        v3 -= v3
        for fi8 := 0; fi8 < 14; fi8++ {
            v9 := GenS1{F0: 9223372036854775807}
            var v10 uint8 = 25
            _ = fi8
            _ = v10
            _ = v9
        }
    }
    v2 += v2
    for fi11 := 0; fi11 < 8; fi11++ {
        fmt.Println(v1, fi11, v2)
        if (-9223372036854775808 != 42) {
            fmt.Println(v1, v3, v2)
        } else {
            if len(v4) > 0 {
                v4[0] = "\n\t"
            }
            for fi12 := 0; fi12 < 5; fi12++ {
                v4 = append(v4, (("" + "foo bar") + "\n\t"))
                fmt.Println(v1, fi12)
                _ = len(v4)
                _ = fi12
            }
        }
        for fi13 := 0; fi13 < 19; fi13++ {
            _ = len(v4)
            _ = fi13
        }
        _ = fi11
    }
    switch -2147483648 {
    case 32767:
        v2 = 1
    case -9223372036854775808:
        for fi14 := 0; fi14 < 17; fi14++ {
            v15 := GenS1{F0: -578}
            if (("\n\t" <= "\n\t") && false) {
                var v16 string = "a"
                _ = v16
            } else {
                for fi17 := 0; fi17 < 4; fi17++ {
                    if (!((!false) && (81 >= 1))) {
                        var v18 int64 = -1
                        _ = v18
                    }
                    v19 := GenS2{F0: -1, F1: 0, F2: -803}
                    if false {
                        if (!(!(!false))) {
                            if (!false) {
                                var v20 string = "hello"
                                var v21 uint32 = 2147483648
                                if len(v4) > 0 {
                                    v4[0] = v20
                                }
                                var v22 bool = true
                                _ = v20
                                _ = v21
                                _ = v22
                            }
                            var v23 int16 = 1000
                            _ = v23
                        } else {
                            fmt.Println(v2, v3, fi17)
                        }
                        _ = len(v4)
                    } else {
                        fmt.Println(v1)
                        var v24 int32 = 1
                        _ = v24
                    }
                    _ = fi17
                    _ = v19
                }
                switch ("\n\t" + "") {
                case "a":
                    for fi25 := 0; fi25 < 13; fi25++ {
                        if (!false) {
                            switch (fi25 ^ fi25) {
                            case 32768:
                                v26 := GenS2{F0: -537, F1: 0, F2: 9223372036854775807}
                                _ = v26
                            case 9223372036854775807:
                                v4 = append(v4, (((("\n\t" + "hello") + "hello") + "a") + ""))
                            default:
                                _ = len(v4)
                            }
                            _ = len(v4)
                        } else {
                            fmt.Println(v3, fi14, v2)
                        }
                        _ = fi25
                    }
                }
                switch fi14 {
                case -842:
                    fmt.Println(fi14, v3, v3)
                case -2147483648:
                    var v27 uint = 2
                    _ = v27
                default:
                    for fi28 := 0; fi28 < 2; fi28++ {
                        v2 = (v2 / (((-(v2))) | 1))
                        _ = fi28
                    }
                }
            }
            fmt.Println(v1, v1)
            _ = fi14
            _ = v15
        }
    case 2147483647:
        v2 = (v2 & (v2 | (v2 % (((v2 - 100)) | 1))))
    }
    v1 = v1
    v29 := GenS1{F0: 0}
    switch ("hello" + "") {
    case "a":
        v2 = 1
    case "\n\t":
        v4 = append(v4, "hello")
    case "hello":
        for fi30 := 0; fi30 < 6; fi30++ {
            v31 := GenS3{F0: 136, F1: (v2 % ((v2) | 1)), F2: true}
            _ = len(v4)
            _ = fi30
            _ = v31
        }
    }
    if true {
        v1 += 3.14159
        if (false || false) {
            _ = len(v4)
            if false {
                var v32 int16 = v3
                switch "a" {
                case "foo bar":
                    for fi33 := 0; fi33 < 15; fi33++ {
                        if len(v4) > 0 {
                            v4[0] = ("a" + ("a" + ("" + ("" + "foo bar"))))
                        }
                        _ = fi33
                    }
                case "a":
                    var v34 int64 = -343
                    _ = v34
                default:
                    switch "foo bar" {
                    case "a":
                        v35 := GenS2{F0: 9223372036854775807, F1: 2, F2: 42}
                        _ = v35
                    case "hello":
                        if false {
                            var v36 string = ""
                            if (0.1 >= 0.0) {
                                for fi37 := 0; fi37 < 9; fi37++ {
                                    v38 := GenS1{F0: -1}
                                    if (-203 <= 0) {
                                        switch ("\n\t" + "foo bar") {
                                        case "foo bar":
                                            if false {
                                                fmt.Println(v1)
                                                for fi39 := 0; fi39 < 2; fi39++ {
                                                    v1 = 0.1
                                                    v32 -= -968
                                                    _ = fi39
                                                }
                                            } else {
                                                for fi40 := 0; fi40 < 11; fi40++ {
                                                    var v41 bool = (!(18446744073709551615 > 2))
                                                    _ = fi40
                                                    _ = v41
                                                }
                                                v3 = -32768
                                                v4 = append(v4, v36)
                                            }
                                        case "\n\t":
                                            for fi42 := 0; fi42 < 12; fi42++ {
                                                if false {
                                                    v43 := GenS3{F0: 0, F1: -128, F2: true}
                                                    var v44 int = 555
                                                    if true {
                                                        switch (fi42 >> ((65535) & 63)) {
                                                        case 445:
                                                            v4 = append(v4, v36)
                                                            fmt.Println(v1)
                                                        }
                                                    }
                                                    _ = v43
                                                    _ = v44
                                                } else {
                                                    switch ((v36 + v36) + "hello") {
                                                    case "hello":
                                                        if (!(true || false)) {
                                                            var v45 float32 = -2.25
                                                            v46 := GenS4{F0: v36, F1: ((v36 + ((v36 + v36) + v36)) + v36)}
                                                            _ = v45
                                                            _ = v46
                                                        }
                                                        fmt.Println(fi42, fi37, v2)
                                                        v3 += 180
                                                    default:
                                                        v2 = (v2 << ((18446744073709551615) & 63))
                                                    }
                                                    var v47 int32 = 0
                                                    _ = v47
                                                }
                                                v1 = 1.5
                                                v48 := GenS4{F0: v36, F1: (v36 + "hello")}
                                                _ = fi42
                                                _ = v48
                                            }
                                        }
                                        v49 := GenS5{F0: v36, F1: 0, F2: 2}
                                        _ = v49
                                    } else {
                                        for fi50 := 0; fi50 < 12; fi50++ {
                                            v51 := GenS2{F0: 987, F1: 65535, F2: -9223372036854775808}
                                            var v52 bool = (0 == 173)
                                            _ = fi50
                                            _ = v51
                                            _ = v52
                                        }
                                        for fi53 := 0; fi53 < 15; fi53++ {
                                            v54 := GenS4{F0: v36, F1: "foo bar"}
                                            if (false || false) {
                                                v1 -= v1
                                            } else {
                                                if len(v4) >= 2 {
                                                    v4 = v4[1:len(v4)]
                                                }
                                                v55 := GenS6{F0: 0}
                                                v32 = 12
                                                _ = v55
                                            }
                                            v36 = v36
                                            if (!(-103 <= v2)) {
                                                var v56 float32 = ((0.0 * (0.0 - (-2.25 + 1e10))) * 0.0)
                                                _ = v56
                                            } else {
                                                v3 -= v32
                                                _ = len(v4)
                                            }
                                            _ = fi53
                                            _ = v54
                                        }
                                    }
                                    var v57 uint32 = 1
                                    _ = fi37
                                    _ = v38
                                    _ = v57
                                }
                                var v58 int16 = (v3 % ((-572) | 1))
                                for fi59 := 0; fi59 < 17; fi59++ {
                                    v32 = (v32 + v3)
                                    var v60 uint64 = 42
                                    if true {
                                        v36 = (v36 + "a")
                                    } else {
                                        v2 = (-(v2))
                                        for fi61 := 0; fi61 < 9; fi61++ {
                                            fmt.Println(fi61, fi61, fi59)
                                            var v62 uint64 = v60
                                            var v63 uint32 = 42
                                            _ = fi61
                                            _ = v62
                                            _ = v63
                                        }
                                        v60 = v60
                                    }
                                    _ = fi59
                                    _ = v60
                                }
                                _ = v58
                            } else {
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                                var v64 float64 = (v1 * v1)
                                v1 = (1e10 / (float64(-2.25) + float64(((float32(((float64((1e10 * 1.5)) * 1e10) - v1)) * 1e10) + -2.25))))
                                _ = v64
                            }
                            if (v2 > v2) {
                                v36 = (("hello" + "a") + v36)
                                if (100 >= v2) {
                                    v65 := GenS7{F0: 255}
                                    var v66 int8 = v2
                                    _ = v65
                                    _ = v66
                                }
                                v3 -= 357
                            } else {
                                v67 := GenS6{F0: 1}
                                if true {
                                    for fi68 := 0; fi68 < 10; fi68++ {
                                        if (false && true) {
                                            v36 = v36
                                        } else {
                                            v2 -= 100
                                        }
                                        v69 := GenS3{F0: 48, F1: (v2 + 127), F2: ((^(v32)) == -879)}
                                        _ = fi68
                                        _ = v69
                                    }
                                    for fi70 := 0; fi70 < 12; fi70++ {
                                        fmt.Println(v36, v1, v2)
                                        _ = fi70
                                    }
                                    fmt.Println(v32, v32, v36)
                                } else {
                                    if true {
                                        for fi71 := 0; fi71 < 4; fi71++ {
                                            _ = len(v4)
                                            _ = fi71
                                        }
                                        var v72 float32 = 1.5
                                        _ = v72
                                    }
                                }
                                _ = v67
                            }
                            _ = v36
                        } else {
                            fmt.Println(v3)
                            var v73 float32 = 1.5
                            if true {
                                v32 = (v32 >> ((763) & 63))
                                for fi74 := 0; fi74 < 7; fi74++ {
                                    var v75 bool = false
                                    var v76 float32 = v73
                                    _ = fi74
                                    _ = v75
                                    _ = v76
                                }
                                switch ("" + "a") {
                                case "a":
                                    var v77 uint16 = 255
                                    _ = v77
                                default:
                                    switch 42 {
                                    case 831:
                                        v78 := GenS8{F0: "a", F1: 615}
                                        _ = v78
                                    }
                                }
                            } else {
                                var v79 uint = 255
                                fmt.Println(v73, v3, v1)
                                _ = v79
                            }
                            _ = v73
                        }
                    case "foo bar":
                        v2 = 1
                    default:
                        v2 -= v2
                    }
                }
                _ = v32
            } else {
                switch 0 {
                case 128:
                    var v80 bool = (375 != (-(v3)))
                    _ = v80
                case 127:
                    if (v2 >= (v2 | -27)) {
                        var v81 bool = (!(1 < 531))
                        var v82 int32 = 42
                        var v83 bool = v81
                        _ = v81
                        _ = v82
                        _ = v83
                    }
                }
                v84 := GenS7{F0: 18446744073709551615}
                _ = v84
            }
            if false {
                var v85 int16 = (v3 >> ((764) & 63))
                fmt.Println(v1, v85)
                v85 = (v85 - (^(v85)))
                _ = v85
            } else {
                _ = len(v4)
            }
        } else {
            v86 := GenS7{F0: 580}
            if true {
                var v87 int64 = 1
                fmt.Println(v2, v87)
                if true {
                    fmt.Println(v3, v1)
                    if true {
                        if len(v4) > 0 {
                            v4[0] = ("a" + "hello")
                        }
                        if (true || (float32(float64(0.1)) < float32(float64(float32(1.5))))) {
                            switch 127 {
                            case -200:
                                if true {
                                    var v88 uint8 = 168
                                    switch 32768 {
                                    case -386:
                                        for fi89 := 0; fi89 < 13; fi89++ {
                                            if (false || (true && false)) {
                                                for fi90 := 0; fi90 < 5; fi90++ {
                                                    switch "a" {
                                                    case "hello":
                                                        _ = len(v4)
                                                        v3 += -1
                                                        v91 := GenS9{F0: -776, F1: (^(v88))}
                                                        var v92 int8 = -128
                                                        _ = v91
                                                        _ = v92
                                                    case "\n\t":
                                                        v93 := GenS1{F0: (v87 * (^(v87)))}
                                                        _ = v93
                                                    default:
                                                        fmt.Println(v87)
                                                    }
                                                    v94 := GenS8{F0: "", F1: 4294967295}
                                                    _ = fi90
                                                    _ = v94
                                                }
                                                v1 = -2.25
                                                for fi95 := 0; fi95 < 12; fi95++ {
                                                    fmt.Println(v1, v87)
                                                    if ((360 >= 471) || true) {
                                                        v96 := GenS10{F0: ("a" + "\n\t")}
                                                        if (true && true) {
                                                            if true {
                                                                switch (("\n\t" + ("" + "")) + "\n\t") {
                                                                case "hello":
                                                                    v88 -= (v88 - v88)
                                                                case "foo bar":
                                                                    var v97 uint = 255
                                                                    _ = v97
                                                                }
                                                            } else {
                                                                var v98 uint64 = 1
                                                                v98 = 202
                                                                var v99 int8 = v2
                                                                _ = v98
                                                                _ = v99
                                                            }
                                                            var v100 uint8 = v88
                                                            _ = v100
                                                        } else {
                                                            v101 := GenS11{F0: 0}
                                                            _ = v101
                                                        }
                                                        if false {
                                                            v102 := GenS12{F0: (v88 % ((v88) | 1)), F1: "\n\t"}
                                                            v4 = append(v4, ("\n\t" + "a"))
                                                            _ = v102
                                                        }
                                                        _ = v96
                                                    }
                                                    v103 := GenS1{F0: int64(v2)}
                                                    _ = fi95
                                                    _ = v103
                                                }
                                            } else {
                                                fmt.Println(v88, v87, v87)
                                                for fi104 := 0; fi104 < 12; fi104++ {
                                                    v2 = (v2 + (v2 | v2))
                                                    _ = fi104
                                                }
                                                for fi105 := 0; fi105 < 20; fi105++ {
                                                    v106 := GenS13{F0: 2}
                                                    _ = fi105
                                                    _ = v106
                                                }
                                            }
                                            switch ((("\n\t" + "\n\t") + ("foo bar" + "a")) + "foo bar") {
                                            case "\n\t":
                                                var v107 int32 = 777
                                                _ = v107
                                            case "":
                                                v108 := GenS9{F0: -286, F1: 255}
                                                _ = v108
                                            case "foo bar":
                                                fmt.Println(v2, v87, v3)
                                            default:
                                                v2 = int8(v88)
                                            }
                                            switch 598 {
                                            case 1:
                                                var v109 bool = (-2.25 == (0.1 * 1e10))
                                                _ = v109
                                            case -9223372036854775808:
                                                switch (-(fi89)) {
                                                case -457:
                                                    for fi110 := 0; fi110 < 15; fi110++ {
                                                        if true {
                                                            v111 := GenS14{F0: fi110, F1: v2, F2: 0}
                                                            _ = v111
                                                        }
                                                        v87 += v87
                                                        if len(v4) > 0 {
                                                            v4[0] = ("hello" + "")
                                                        }
                                                        _ = fi110
                                                    }
                                                case 721:
                                                    var v112 int16 = (v3 << ((42) & 63))
                                                    _ = v112
                                                default:
                                                    _ = len(v4)
                                                }
                                            default:
                                                if (!(true || true)) {
                                                    if len(v4) > 0 {
                                                        v4[0] = (((("foo bar" + "foo bar") + "a") + "hello") + "hello")
                                                    }
                                                    if (2 != 2) {
                                                        var v113 float64 = (((3.14159 / 1.0) - v1) / 0.1)
                                                        v113 = 3.14159
                                                        v2 += -128
                                                        _ = v113
                                                    } else {
                                                        var v114 int32 = 1
                                                        v3 = (v3 & 32767)
                                                        _ = v114
                                                    }
                                                    if (float32((3.14159 / v1)) <= float32(v1)) {
                                                        v88 = (^(v88))
                                                        _ = len(v4)
                                                        v88 += v88
                                                    }
                                                } else {
                                                    if (!true) {
                                                        var v115 uint32 = 1
                                                        _ = v115
                                                    } else {
                                                        var v116 uint8 = v88
                                                        v117 := GenS13{F0: 760}
                                                        _ = v116
                                                        _ = v117
                                                    }
                                                    var v118 float64 = v1
                                                    var v119 float32 = 1e10
                                                    _ = v118
                                                    _ = v119
                                                }
                                            }
                                            _ = fi89
                                        }
                                    case 127:
                                        if true {
                                            v87 -= (v87 % (((v87 / (((v87 ^ v87)) | 1))) | 1))
                                            for fi120 := 0; fi120 < 4; fi120++ {
                                                if (0 == 213) {
                                                    var v121 int16 = v3
                                                    if true {
                                                        v4 = append(v4, "")
                                                    } else {
                                                        v2 += (v2 + v2)
                                                        v1 += v1
                                                    }
                                                    _ = v121
                                                } else {
                                                    fmt.Println(fi120)
                                                }
                                                for fi122 := 0; fi122 < 15; fi122++ {
                                                    if true {
                                                        var v123 uint32 = 42
                                                        var v124 float64 = v1
                                                        _ = v123
                                                        _ = v124
                                                    } else {
                                                        if false {
                                                            v125 := GenS12{F0: v88, F1: "hello"}
                                                            v87 -= (v87 & (-(v87)))
                                                            _ = v125
                                                        } else {
                                                            fmt.Println(fi122, v2, v88)
                                                        }
                                                        v88 = v88
                                                    }
                                                    _ = len(v4)
                                                    _ = fi122
                                                }
                                                _ = fi120
                                            }
                                        }
                                    case 128:
                                        var v126 uint32 = 1
                                        _ = v126
                                    default:
                                        v88 += (v88 >> ((0) & 63))
                                    }
                                    _ = v88
                                } else {
                                    var v127 int = 9223372036854775807
                                    v87 -= (v87 % ((v87) | 1))
                                    _ = v127
                                }
                            case 202:
                                if (false && (!(193 < 2))) {
                                    var v128 float32 = 0.1
                                    _ = v128
                                }
                            }
                            v129 := GenS13{F0: 4294967295}
                            v2 = (^(v2))
                            _ = v129
                        } else {
                            v87 -= v87
                            v1 = (float64(3.14159) * 1.5)
                            if (18446744073709551615 >= 4294967295) {
                                switch 32767 {
                                case -9223372036854775808:
                                    var v130 int16 = (v3 + v3)
                                    _ = v130
                                case 643:
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                default:
                                    if true {
                                        v131 := GenS13{F0: 1}
                                        switch 128 {
                                        case 128:
                                            for fi132 := 0; fi132 < 13; fi132++ {
                                                if len(v4) >= 2 {
                                                    v4 = v4[1:len(v4)]
                                                }
                                                switch -2147483648 {
                                                case 128:
                                                    v4 = append(v4, "foo bar")
                                                case -2147483648:
                                                    v133 := GenS8{F0: ((("foo bar" + "") + "a") + "hello"), F1: 2}
                                                    _ = v133
                                                }
                                                _ = fi132
                                            }
                                        case -79:
                                            for fi134 := 0; fi134 < 20; fi134++ {
                                                for fi135 := 0; fi135 < 18; fi135++ {
                                                    var v136 int8 = 1
                                                    fmt.Println(v87, v3, v2)
                                                    switch "" {
                                                    case "hello":
                                                        v137 := GenS10{F0: "\n\t"}
                                                        var v138 uint32 = 0
                                                        _ = v137
                                                        _ = v138
                                                    case "foo bar":
                                                        fmt.Println(v3, v87, fi135)
                                                    default:
                                                        if true {
                                                            fmt.Println(v136, v87)
                                                            v87 -= (^(v87))
                                                        }
                                                    }
                                                    _ = fi135
                                                    _ = v136
                                                }
                                                if (48 >= 4294967295) {
                                                    v1 = ((3.14159 * 3.14159) - ((((v1 - -2.25) * v1) + (v1 / 1.0)) - v1))
                                                    v87 += v87
                                                    v139 := GenS15{F0: 0, F1: (^(v2))}
                                                    _ = v139
                                                } else {
                                                    var v140 uint64 = 18446744073709551615
                                                    _ = v140
                                                }
                                                _ = fi134
                                            }
                                        default:
                                            v2 -= v2
                                        }
                                        _ = v131
                                    }
                                }
                                if (42 == (^(v87))) {
                                    v141 := GenS16{F0: 545, F1: 18446744073709551615, F2: (v2 << ((18446744073709551615) & 63))}
                                    _ = v141
                                } else {
                                    if (true || (0 != 42)) {
                                        var v142 float32 = 1.5
                                        if false {
                                            switch 798 {
                                            case 4294967295:
                                                fmt.Println(v87)
                                            default:
                                                if ((252 != 244) && (((!(0 > 1)) && (("\n\t" + "foo bar") < "a")) || true)) {
                                                    if len(v4) > 0 {
                                                        v4[0] = (("hello" + (("foo bar" + "foo bar") + "hello")) + "\n\t")
                                                    }
                                                    switch 42 {
                                                    case -853:
                                                        v87 -= v87
                                                    case -1:
                                                        if (!((v3 * v3) == (v3 >> ((187) & 63)))) {
                                                            if len(v4) >= 2 {
                                                                v4 = v4[1:len(v4)]
                                                            }
                                                            var v143 string = "a"
                                                            switch 4294967295 {
                                                            case 9223372036854775807:
                                                                if (v2 != v2) {
                                                                    var v144 int64 = (v87 << ((42) & 63))
                                                                    _ = v144
                                                                } else {
                                                                    v143 = (v143 + "hello")
                                                                    if false {
                                                                        fmt.Println(v3)
                                                                        var v145 int8 = v2
                                                                        v146 := GenS10{F0: "a"}
                                                                        _ = v145
                                                                        _ = v146
                                                                    } else {
                                                                        switch (v143 + "hello") {
                                                                        case "foo bar":
                                                                            if (v2 < (v2 + v2)) {
                                                                                var v147 uint8 = 0
                                                                                v87 = v87
                                                                                v143 = ""
                                                                                _ = v147
                                                                            }
                                                                        case "\n\t":
                                                                            v2 = v2
                                                                        case "":
                                                                            var v148 uint16 = 42
                                                                            _ = v148
                                                                        default:
                                                                            if true {
                                                                                var v149 uint32 = 2147483647
                                                                                _ = v149
                                                                            } else {
                                                                                v143 = v143
                                                                            }
                                                                        }
                                                                    }
                                                                }
                                                            default:
                                                                if len(v4) >= 2 {
                                                                    v4 = v4[1:len(v4)]
                                                                }
                                                            }
                                                            _ = v143
                                                        } else {
                                                            if len(v4) > 0 {
                                                                v4[0] = (("" + "a") + "")
                                                            }
                                                            for fi150 := 0; fi150 < 14; fi150++ {
                                                                switch -839 {
                                                                case 42:
                                                                    v2 = (^(v2))
                                                                case 127:
                                                                    for fi151 := 0; fi151 < 10; fi151++ {
                                                                        v152 := GenS1{F0: v87}
                                                                        v142 += v142
                                                                        v153 := GenS16{F0: 42, F1: 575, F2: (v2 + 100)}
                                                                        _ = fi151
                                                                        _ = v152
                                                                        _ = v153
                                                                    }
                                                                case 430:
                                                                    var v154 uint64 = 147
                                                                    _ = v154
                                                                }
                                                                fmt.Println(fi150, v3, v87)
                                                                var v155 int8 = 42
                                                                _ = fi150
                                                                _ = v155
                                                            }
                                                        }
                                                    case 9223372036854775807:
                                                        v156 := GenS17{F0: 897, F1: 2147483648, F2: 42}
                                                        _ = v156
                                                    }
                                                }
                                            }
                                            if len(v4) > 0 {
                                                v4[0] = "a"
                                            }
                                            fmt.Println(v1, v3)
                                        }
                                        _ = v142
                                    }
                                    if false {
                                        if ((true || true) || (v87 == 962)) {
                                            var v157 uint32 = 2147483648
                                            _ = v157
                                        } else {
                                            v87 = (v87 & 42)
                                            fmt.Println(v1, v87, v3)
                                            var v158 int64 = 42
                                            _ = v158
                                        }
                                        var v159 int64 = -1
                                        _ = v159
                                    } else {
                                        var v160 uint64 = 831
                                        v3 += (v3 + v3)
                                        switch ("" + ("a" + "foo bar")) {
                                        case "\n\t":
                                            var v161 uint16 = 1
                                            _ = v161
                                        case "a":
                                            v162 := GenS5{F0: "", F1: 777, F2: v160}
                                            _ = v162
                                        default:
                                            if (0.1 >= (-2.25 - 0.0)) {
                                                v4 = append(v4, "foo bar")
                                            }
                                        }
                                        _ = v160
                                    }
                                }
                                var v163 int16 = -62
                                _ = v163
                            }
                        }
                        var v164 float64 = ((v1 / 0.1) * (0.1 / -2.25))
                        _ = v164
                    } else {
                        switch -9223372036854775808 {
                        case 127:
                            fmt.Println(v87)
                        }
                        fmt.Println(v2)
                    }
                    for fi165 := 0; fi165 < 12; fi165++ {
                        v4 = append(v4, ("\n\t" + ""))
                        _ = fi165
                    }
                } else {
                    for fi166 := 0; fi166 < 19; fi166++ {
                        if false {
                            if len(v4) > 0 {
                                v4[0] = (("" + ("\n\t" + "foo bar")) + "a")
                            }
                            if ((v2 / ((v2) | 1)) != 42) {
                                if true {
                                    if false {
                                        fmt.Println(v87)
                                    } else {
                                        for fi167 := 0; fi167 < 17; fi167++ {
                                            v2 = v2
                                            v2 += 127
                                            _ = fi167
                                        }
                                    }
                                    if false {
                                        var v168 uint = 4294967295
                                        _ = v168
                                    } else {
                                        v87 = 1
                                        switch fi166 {
                                        case -1:
                                            var v169 uint64 = 42
                                            _ = v169
                                        case -2147483648:
                                            if (!(-634 <= (^(fi166)))) {
                                                var v170 uint32 = 2
                                                v171 := GenS18{F0: (^(v2)), F1: 1, F2: v2}
                                                var v172 int8 = v2
                                                _ = v170
                                                _ = v171
                                                _ = v172
                                            } else {
                                                var v173 uint16 = 1
                                                v174 := GenS18{F0: (v2 - 100), F1: 4294967295, F2: (v2 ^ v2)}
                                                switch -1 {
                                                case -1:
                                                    v2 -= (v2 & 0)
                                                }
                                                _ = v173
                                                _ = v174
                                            }
                                        case 367:
                                            switch "\n\t" {
                                            case "\n\t":
                                                for fi175 := 0; fi175 < 16; fi175++ {
                                                    var v176 uint16 = 822
                                                    if true {
                                                        var v177 uint64 = 4294967295
                                                        _ = v177
                                                    }
                                                    v176 += 932
                                                    _ = fi175
                                                    _ = v176
                                                }
                                            case "foo bar":
                                                v87 += v87
                                            case "":
                                                if true {
                                                    var v178 uint8 = 127
                                                    v3 += 1000
                                                    switch ("foo bar" + "a") {
                                                    case "hello":
                                                        v179 := GenS19{F0: 1, F1: 127, F2: fi166}
                                                        _ = v179
                                                    case "foo bar":
                                                        var v180 string = "hello"
                                                        _ = v180
                                                    case "a":
                                                        var v181 bool = false
                                                        _ = v181
                                                    default:
                                                        var v182 string = "foo bar"
                                                        _ = v182
                                                    }
                                                    _ = v178
                                                } else {
                                                    if true {
                                                        for fi183 := 0; fi183 < 16; fi183++ {
                                                            var v184 float64 = v1
                                                            _ = fi183
                                                            _ = v184
                                                        }
                                                        for fi185 := 0; fi185 < 13; fi185++ {
                                                            if true {
                                                                if (false && false) {
                                                                    fmt.Println(fi185, v2)
                                                                    if (1 <= 0) {
                                                                        var v186 int16 = (^(v3))
                                                                        var v187 float64 = (((v1 - v1) / (3.14159 + v1)) / (v1 - v1))
                                                                        _ = v186
                                                                        _ = v187
                                                                    }
                                                                } else {
                                                                    v2 -= 116
                                                                }
                                                            } else {
                                                                if len(v4) > 0 {
                                                                    v4[0] = "hello"
                                                                }
                                                            }
                                                            if false {
                                                                if len(v4) >= 2 {
                                                                    v4 = v4[1:len(v4)]
                                                                }
                                                                v188 := GenS11{F0: 663}
                                                                _ = v188
                                                            } else {
                                                                var v189 int8 = int8(fi185)
                                                                _ = v189
                                                            }
                                                            v190 := GenS20{F0: fi185, F1: 18446744073709551615}
                                                            _ = fi185
                                                            _ = v190
                                                        }
                                                        var v191 int = fi166
                                                        _ = v191
                                                    } else {
                                                        fmt.Println(v1)
                                                        var v192 uint64 = 286
                                                        _ = v192
                                                    }
                                                }
                                            default:
                                                var v193 float32 = float32(v1)
                                                _ = v193
                                            }
                                        default:
                                            v2 = (-(v2))
                                        }
                                    }
                                }
                                v1 += v1
                            }
                            for fi194 := 0; fi194 < 5; fi194++ {
                                switch fi166 {
                                case 32767:
                                    var v195 uint32 = 933
                                    _ = v195
                                }
                                var v196 string = ("\n\t" + "")
                                v87 += (-(v87))
                                _ = fi194
                                _ = v196
                            }
                        } else {
                            fmt.Println(v2)
                            var v197 int16 = 1000
                            v3 = v197
                            _ = v197
                        }
                        fmt.Println(v87, v3, v2)
                        _ = fi166
                    }
                }
                _ = v87
            }
            switch "foo bar" {
            case "hello":
                for fi198 := 0; fi198 < 11; fi198++ {
                    v2 += v2
                    _ = fi198
                }
            case "":
                for fi199 := 0; fi199 < 5; fi199++ {
                    v200 := GenS21{F0: v3, F1: -36, F2: ("" > (("foo bar" + "hello") + "\n\t"))}
                    _ = fi199
                    _ = v200
                }
            case "\n\t":
                switch "a" {
                case "foo bar":
                    fmt.Println(v2, v2, v2)
                }
            }
            _ = v86
        }
        fmt.Println(v3, v1, v3)
    } else {
        v2 -= -1
    }
    fmt.Println(v1)
    var v201 uint64 = 0
    var v202 int = 128
    var v203 uint64 = v201
    v2 += v2
    v4 = append(v4, ("foo bar" + ("hello" + "\n\t")))
    _ = len(v4)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v201:", v201)
    fmt.Println("v202:", v202)
    fmt.Println("v203:", v203)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v201
    _ = v202
    _ = v203
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
