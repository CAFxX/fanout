package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool }
type GenS2 struct { F0 int64; F1 uint16; F2 uint8 }
type GenS3 struct { F0 string; F1 uint64 }
type GenS4 struct { F0 uint64; F1 uint8 }
type GenS5 struct { F0 int64; F1 uint8; F2 int }
type GenS6 struct { F0 uint; F1 int16 }
type GenS7 struct { F0 uint8; F1 uint64 }
type GenS8 struct { F0 string; F1 uint }
type GenS9 struct { F0 int }
type GenS10 struct { F0 uint32; F1 int8; F2 int16 }
type GenS11 struct { F0 uint32; F1 int64; F2 int64 }
type GenS12 struct { F0 bool; F1 uint16; F2 uint8 }
func genFn0() int8 {
    for fi7 := 0; fi7 < 6; fi7++ {
        v8 := []int{32767, fi7}
        var v9 int16 = -610
        if (false || false) {
            v10 := GenS1{F0: ((((("a" + "\n\t") + (("" + ("hello" + "\n\t")) + "a")) + "") + "hello") + (("" + "\n\t") + "foo bar")), F1: (-2147483648 >= 495)}
            v8 = append(v8, (fi7 % ((fi7) | 1)))
            v11 := GenS2{F0: 0, F1: 2, F2: 42}
            _ = v10
            _ = v11
        } else {
            if false {
                for fi12 := 0; fi12 < 1; fi12++ {
                    if (fi7 == (fi7 + 0)) {
                        if true {
                            var v13 uint32 = 2147483648
                            switch 32767 {
                            case 543:
                                var v14 bool = true
                                _ = v14
                            case -672:
                                switch fi12 {
                                case -1:
                                    switch (("foo bar" + (("foo bar" + "a") + "a")) + ((("hello" + "hello") + "hello") + (("hello" + ((("hello" + "hello") + ("\n\t" + "hello")) + ((("foo bar" + "foo bar") + "a") + "a"))) + ""))) {
                                    case "a":
                                        v9 += v9
                                    default:
                                        v15 := GenS2{F0: 0, F1: 255, F2: 0}
                                        _ = v15
                                    }
                                default:
                                    v16 := GenS1{F0: "foo bar", F1: false}
                                    _ = v16
                                }
                            case 32767:
                                var v17 int16 = (v9 + v9)
                                _ = v17
                            default:
                                v9 = (v9 << ((65535) & 63))
                            }
                            _ = v13
                        } else {
                            v18 := GenS2{F0: -9223372036854775808, F1: 258, F2: 2}
                            switch "a" {
                            case "a":
                                v9 += v9
                                if len(v8) >= 2 {
                                    v8 = v8[1:len(v8)]
                                }
                                if (!((1e10 + ((1.5 + (3.14159 / 1e10)) / float64(3.14159))) > -2.25)) {
                                    fmt.Println(v9, fi12)
                                }
                            }
                            switch fi12 {
                            case 42:
                                switch "" {
                                case "foo bar":
                                    if false {
                                        var v19 int8 = -128
                                        _ = v19
                                    }
                                case "\n\t":
                                    fmt.Println(fi12)
                                case "a":
                                    v20 := GenS1{F0: ((((("foo bar" + "foo bar") + "hello") + ("\n\t" + "foo bar")) + "foo bar") + "hello"), F1: (!true)}
                                    _ = v20
                                default:
                                    var v21 int = -411
                                    _ = v21
                                }
                            case 1:
                                if (true || (65535 != 0)) {
                                    v9 = v9
                                    v9 += 593
                                    if true {
                                        var v22 int32 = -229
                                        switch (fi12 << ((637) & 63)) {
                                        case 2147483647:
                                            v8 = append(v8, (fi7 / ((fi7) | 1)))
                                        case 9223372036854775807:
                                            v22 = (v22 & (v22 / ((v22) | 1)))
                                        }
                                        _ = v22
                                    }
                                }
                            }
                            if true {
                                v9 -= (v9 & (v9 + (v9 << ((2) & 63))))
                                switch fi12 {
                                case 32768:
                                    v23 := GenS1{F0: "\n\t", F1: true}
                                    _ = v23
                                case 9223372036854775807:
                                    v9 = 0
                                default:
                                    if (!(true || false)) {
                                        var v24 int = fi7
                                        _ = v24
                                    }
                                }
                                if (false || (false && true)) {
                                    switch "hello" {
                                    case "":
                                        switch "a" {
                                        case "hello":
                                            if ("\n\t" < "foo bar") {
                                                v8 = append(v8, fi12)
                                                if (true && (661 < 374)) {
                                                    if (("" + "") < "a") {
                                                        v25 := GenS1{F0: ("a" + "\n\t"), F1: (!false)}
                                                        if len(v8) >= 2 {
                                                            v8 = v8[1:len(v8)]
                                                        }
                                                        v26 := GenS3{F0: "a", F1: 774}
                                                        _ = v25
                                                        _ = v26
                                                    }
                                                }
                                            } else {
                                                v9 -= 32767
                                            }
                                        case "\n\t":
                                            v27 := GenS4{F0: 18446744073709551615, F1: 1}
                                            _ = v27
                                        }
                                    case "foo bar":
                                        v28 := GenS1{F0: "a", F1: true}
                                        _ = v28
                                    case "\n\t":
                                        var v29 uint8 = 69
                                        _ = v29
                                    default:
                                        v8 = append(v8, int(v9))
                                    }
                                } else {
                                    v9 = v9
                                }
                            } else {
                                switch ("hello" + "\n\t") {
                                case "foo bar":
                                    v9 = (v9 | 0)
                                }
                                switch (("foo bar" + "a") + "") {
                                case "foo bar":
                                    v9 = (^(v9))
                                case "hello":
                                    v30 := GenS5{F0: -1, F1: 255, F2: fi12}
                                    _ = v30
                                case "a":
                                    var v31 uint32 = 889
                                    _ = v31
                                }
                                var v32 int16 = (v9 - int16(fi7))
                                _ = v32
                            }
                            v9 = (v9 & (-(v9)))
                            _ = v18
                        }
                        if len(v8) >= 2 {
                            v8 = v8[1:len(v8)]
                        }
                        v9 = 32767
                    } else {
                        v9 = v9
                        v9 -= v9
                    }
                    _ = fi12
                }
            }
            if len(v8) > 0 {
                v8[0] = fi7
            }
        }
        _ = fi7
        _ = v8
        _ = v9
    }
    var v33 bool = true
    for fi34 := 0; fi34 < 5; fi34++ {
        v35 := GenS2{F0: -9223372036854775808, F1: 428, F2: 1}
        _ = fi34
        _ = v35
    }
    if (2 > 19) {
        var v36 uint8 = 0
        var v37 int32 = -1
        v38 := GenS1{F0: "a", F1: (!false)}
        _ = v36
        _ = v37
        _ = v38
    }
    for fi39 := 0; fi39 < 14; fi39++ {
        v33 = true
        fmt.Println(v33, v33)
        _ = fi39
    }
    _ = v33
    return 64
}
func genFn1() int64 {
    var v40 bool = false
    v41 := []int{32768}
    v40 = (42 > 1)
    _ = v40
    _ = v41
    return 899
}

func main() {
    v1 := make(map[string]bool)
    var v2 bool = (false && true)
    var v3 int64 = -1
    v4 := []string{(("" + "") + (("\n\t" + (("foo bar" + "\n\t") + "foo bar")) + "a")), "a"}
    v5 := make(map[int]bool)
    var v6 uint64 = 2
    for fi42 := 0; fi42 < 2; fi42++ {
        var v43 int8 = -92
        v2 = true
        _ = fi42
        _ = v43
    }
    if (false || v2) {
        var v44 uint = 985
        var v45 uint16 = 819
        _ = v44
        _ = v45
    }
    v6 = (v6 * (v6 << ((65535) & 63)))
    for fi46 := 0; fi46 < 4; fi46++ {
        v4 = append(v4, "")
        _ = fi46
    }
    var v47 int32 = 0
    if false {
        for fi48 := 0; fi48 < 16; fi48++ {
            v47 = 1
            _ = fi48
        }
        v6 = v6
        delete(v1, "\n\t")
    }
    v49 := GenS1{F0: "\n\t", F1: (float64(0.0) != 0.1)}
    switch 128 {
    case 273:
        if v2 {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            if ((v3 >> ((65535) & 63)) >= 165) {
                if false {
                    for fi50 := 0; fi50 < 11; fi50++ {
                        fmt.Println(fi50)
                        _ = fi50
                    }
                    v3 -= (v3 ^ 9223372036854775807)
                    fmt.Println(v3)
                } else {
                    var v51 int = 32768
                    _ = v51
                }
                if (!true) {
                    v2 = (1 == -1)
                    v47 = (v47 | (v47 * (v47 | v47)))
                    switch (("hello" + "") + ("" + ((("a" + "a") + "foo bar") + ("a" + "\n\t")))) {
                    case "a":
                        v52 := GenS1{F0: "hello", F1: v2}
                        _ = v52
                    case "foo bar":
                        _ = len(v5)
                    }
                } else {
                    var v53 uint32 = 2
                    switch 1 {
                    case -2147483648:
                        v54 := GenS2{F0: genFn1(), F1: 65535, F2: uint8(v6)}
                        _ = v54
                    default:
                        for fi55 := 0; fi55 < 20; fi55++ {
                            var v56 int16 = 1000
                            _ = fi55
                            _ = v56
                        }
                    }
                    _ = v53
                }
            } else {
                if len(v4) > 0 {
                    v4[0] = ""
                }
            }
            var v57 uint32 = 0
            _ = v57
        }
    default:
        v6 += v6
    }
    for fi58 := 0; fi58 < 10; fi58++ {
        if (!(genFn1() != -1)) {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        } else {
            for fi59 := 0; fi59 < 13; fi59++ {
                var v60 int16 = 243
                if v2 {
                    if v2 {
                        v60 = -32768
                        v2 = false
                    } else {
                        v6 = v6
                        if true {
                            v6 += (v6 | v6)
                            if false {
                                var v61 int8 = -1
                                v1[("a" + "a")] = ((v3 - genFn1()) < 9223372036854775807)
                                _ = v61
                            } else {
                                var v62 string = ""
                                _ = v62
                            }
                            fmt.Println(v2, v47, v47)
                            v63 := GenS3{F0: ("a" + "a"), F1: v6}
                            v6 = 0
                            switch "" {
                            case "":
                                switch ("" + ((((("foo bar" + (("" + (("foo bar" + "foo bar") + ("hello" + "hello"))) + "")) + ("foo bar" + "\n\t")) + "") + "hello") + "")) {
                                case "\n\t":
                                    var v64 int32 = -2147483648
                                    _ = v64
                                case "hello":
                                    if ((-2.25 * 1e10) < (-2.25 / 0.1)) {
                                        switch "\n\t" {
                                        case "hello":
                                            if (!true) {
                                                var v65 float32 = 3.14159
                                                var v66 bool = v2
                                                if (!(2147483648 != 721)) {
                                                    v3 = v3
                                                    v66 = v66
                                                }
                                                v67 := GenS3{F0: "", F1: (v6 << ((42) & 63))}
                                                _ = v65
                                                _ = v66
                                                _ = v67
                                            } else {
                                                switch fi58 {
                                                case 954:
                                                    if v2 {
                                                        if v2 {
                                                            switch (-(fi58)) {
                                                            case -2147483648:
                                                                v68 := v1["hello"]
                                                                _ = v68
                                                            case 726:
                                                                if v2 {
                                                                    v69 := v5[(fi58 + 42)]
                                                                    v70 := GenS6{F0: 4294967295, F1: v60}
                                                                    switch "\n\t" {
                                                                    case "hello":
                                                                        v71 := GenS2{F0: (v3 + v3), F1: 0, F2: 252}
                                                                        _ = v71
                                                                    case "a":
                                                                        if v2 {
                                                                            var v72 uint16 = 1
                                                                            v69 = v2
                                                                            var v73 float32 = 1e10
                                                                            if v2 {
                                                                                v73 = (1e10 + v73)
                                                                                v2 = true
                                                                                v60 -= (v60 | v60)
                                                                            } else {
                                                                                v1[("" + "foo bar")] = false
                                                                                if len(v4) > 0 {
                                                                                    v4[0] = ("" + "")
                                                                                }
                                                                            }
                                                                            _ = v72
                                                                            _ = v73
                                                                        } else {
                                                                            v1[("" + "")] = true
                                                                            var v74 uint8 = 1
                                                                            v2 = ((65 != -118) || v2)
                                                                            v1["foo bar"] = v2
                                                                            _ = v74
                                                                        }
                                                                    case "\n\t":
                                                                        v75 := GenS2{F0: v3, F1: 0, F2: 1}
                                                                        _ = v75
                                                                    default:
                                                                        var v76 int32 = 220
                                                                        _ = v76
                                                                    }
                                                                    fmt.Println(v69, v47)
                                                                    var v77 uint64 = v6
                                                                    _ = v69
                                                                    _ = v70
                                                                    _ = v77
                                                                }
                                                            case 505:
                                                                var v78 int8 = 47
                                                                _ = v78
                                                            }
                                                            var v79 string = ("foo bar" + "\n\t")
                                                            v6 -= v6
                                                            _ = v79
                                                        }
                                                    }
                                                case 32767:
                                                    var v80 uint64 = 18446744073709551615
                                                    _ = v80
                                                case 32768:
                                                    var v81 uint16 = 42
                                                    _ = v81
                                                default:
                                                    if false {
                                                        var v82 int64 = genFn1()
                                                        var v83 float32 = (3.14159 * ((0.0 - float32(0.0)) * ((1e10 / (1e10 / ((0.0 / 1e10) + -2.25))) / 3.14159)))
                                                        _ = v82
                                                        _ = v83
                                                    } else {
                                                        v6 = (^(v6))
                                                        _ = len(v4)
                                                        var v84 int64 = (v3 / ((1) | 1))
                                                        var v85 uint64 = 42
                                                        fmt.Println(v6)
                                                        _ = v84
                                                        _ = v85
                                                    }
                                                }
                                                fmt.Println(v2, v2)
                                            }
                                        case "a":
                                            if (v2 || v2) {
                                                var v86 float64 = -2.25
                                                var v87 uint = 4294967295
                                                v1["hello"] = v2
                                                if false {
                                                    v2 = v2
                                                    v88 := GenS2{F0: v3, F1: 1, F2: 169}
                                                    _ = v88
                                                }
                                                if true {
                                                    if v2 {
                                                        v2 = v2
                                                    }
                                                    var v89 uint64 = v6
                                                    _ = v89
                                                }
                                                switch (fi59 + fi59) {
                                                case 1:
                                                    var v90 uint32 = 2147483647
                                                    _ = v90
                                                case 32767:
                                                    v91 := GenS4{F0: v6, F1: 42}
                                                    _ = v91
                                                }
                                                _ = v86
                                                _ = v87
                                            }
                                        default:
                                            v3 += genFn1()
                                        }
                                        v92 := GenS2{F0: (v3 ^ 555), F1: 899, F2: 255}
                                        _ = v92
                                    } else {
                                        if (false && v2) {
                                            fmt.Println(fi58)
                                            if (32767 > 1) {
                                                if (v2 && false) {
                                                    v93 := GenS1{F0: "foo bar", F1: false}
                                                    if v2 {
                                                        v94 := GenS7{F0: 127, F1: 572}
                                                        v6 += 0
                                                        _ = v94
                                                    } else {
                                                        var v95 float64 = (0.1 / (0.0 + (3.14159 / 1.0)))
                                                        _ = v95
                                                    }
                                                    var v96 int32 = (^(v47))
                                                    _ = v93
                                                    _ = v96
                                                }
                                                _ = len(v4)
                                                v6 = v6
                                                v97 := GenS5{F0: -959, F1: 255, F2: fi58}
                                                _ = v97
                                            }
                                            _ = len(v4)
                                        } else {
                                            v60 += v60
                                            v2 = false
                                            fmt.Println(v3)
                                        }
                                    }
                                }
                            case "\n\t":
                                v4 = append(v4, "hello")
                            default:
                                v2 = v2
                            }
                            v5[-2147483648] = (v3 == v3)
                            _ = v63
                        } else {
                            v4 = append(v4, "\n\t")
                            v60 = (v60 & v60)
                            var v98 int = fi58
                            _ = v98
                        }
                        v99 := GenS8{F0: "", F1: 2}
                        v100 := GenS5{F0: int64(v60), F1: 128, F2: (^(fi59))}
                        _ = v100
                        _ = v99
                    }
                    var v101 uint8 = 148
                    var v102 float32 = ((float32((-2.25 / 0.1)) * (1.5 + -2.25)) / 1.5)
                    fmt.Println(v102, v3)
                    fmt.Println(v47)
                    var v103 uint8 = 59
                    if v2 {
                        var v104 int64 = -9223372036854775808
                        var v105 int16 = v60
                        v60 -= v60
                        _ = v104
                        _ = v105
                    } else {
                        var v106 uint16 = uint16(v101)
                        if true {
                            var v107 uint16 = v106
                            v4 = append(v4, "")
                            _ = v107
                        } else {
                            var v108 int16 = (v60 | -1)
                            if (!true) {
                                var v109 float32 = 1e10
                                v47 = (v47 >> ((2) & 63))
                                _ = v109
                            } else {
                                v110 := GenS9{F0: fi59}
                                if (!(2 < 266)) {
                                    if v2 {
                                        switch ("\n\t" + "") {
                                        case "foo bar":
                                            if v2 {
                                                v1["foo bar"] = true
                                                var v111 bool = ((fi58 / (((fi59 + (fi59 & fi59))) | 1)) == int(v106))
                                                v102 += (v102 + (v102 * 3.14159))
                                                if (v111 || false) {
                                                    var v112 int = (fi59 >> ((4294967295) & 63))
                                                    v102 = 1.5
                                                    _ = v112
                                                }
                                                var v113 bool = (fi59 == fi59)
                                                _ = v111
                                                _ = v113
                                            } else {
                                                var v114 int16 = (-(v108))
                                                v115 := GenS2{F0: (^(v3)), F1: v106, F2: (v103 + 252)}
                                                _ = v114
                                                _ = v115
                                            }
                                        default:
                                            var v116 float64 = ((1.5 / (((0.0 - -2.25) - float64((float32(3.14159) / 1.5))) * 1.5)) / (float64(v102) - float64(float32(3.14159))))
                                            _ = v116
                                        }
                                        v6 += (^(v6))
                                    } else {
                                        if (v108 > v60) {
                                            v60 += (v108 % ((v108) | 1))
                                            v117 := GenS2{F0: v3, F1: v106, F2: 1}
                                            _ = v117
                                        } else {
                                            if v2 {
                                                v103 -= (v101 - (v101 / (((v101 << ((65535) & 63))) | 1)))
                                                fmt.Println(v3, v108)
                                                if v2 {
                                                    v118 := GenS5{F0: 42, F1: (v101 & v103), F2: 4294967295}
                                                    v102 = v102
                                                    switch "\n\t" {
                                                    case "\n\t":
                                                        v101 -= (v101 << ((uint(v106)) & 63))
                                                    case "hello":
                                                        v119 := GenS10{F0: 155, F1: genFn0(), F2: int16(fi59)}
                                                        _ = v119
                                                    case "":
                                                        v108 -= v60
                                                    default:
                                                        fmt.Println(v6)
                                                    }
                                                    v106 = v106
                                                    _ = v118
                                                } else {
                                                    delete(v5, fi58)
                                                    v120 := GenS6{F0: 15, F1: v60}
                                                    _ = v120
                                                }
                                            }
                                        }
                                    }
                                    v121 := GenS10{F0: 0, F1: -1, F2: v108}
                                    v103 -= (v103 % (((v103 >> ((65535) & 63))) | 1))
                                    _ = v121
                                } else {
                                    if (!v2) {
                                        if v2 {
                                            var v122 float64 = 1.5
                                            if len(v4) > 0 {
                                                v4[0] = ("a" + ("" + ""))
                                            }
                                            _ = v122
                                        }
                                        v4 = append(v4, "foo bar")
                                    } else {
                                        v123 := GenS11{F0: 42, F1: genFn1(), F2: (v3 % (((v3 - int64(fi59))) | 1))}
                                        _ = v123
                                    }
                                    var v124 float64 = 3.14159
                                    v106 = v106
                                    _ = v124
                                }
                                _ = v110
                            }
                            v6 = v6
                            _ = v108
                        }
                        _ = v106
                    }
                    _ = v101
                    _ = v102
                    _ = v103
                }
                _ = fi59
                _ = v60
            }
            fmt.Println(v47, v3)
            var v125 float32 = float32((0.0 * 0.0))
            _ = v125
        }
        v4 = append(v4, "hello")
        switch "" {
        case "hello":
            for fi126 := 0; fi126 < 6; fi126++ {
                if (float64(1e10) >= (((float64(float32(3.14159)) - (1e10 - 0.1)) - 1.5) - ((-2.25 - (1e10 * 0.0)) - (-2.25 + 1.5)))) {
                    var v127 int8 = -1
                    var v128 uint64 = 1
                    _ = v127
                    _ = v128
                }
                var v129 int64 = (v3 << ((42) & 63))
                v3 -= (v3 ^ int64(v6))
                var v130 uint32 = 2147483647
                v131 := GenS12{F0: v2, F1: 2, F2: 127}
                _ = fi126
                _ = v129
                _ = v130
                _ = v131
            }
        }
        _ = fi58
    }
    if (4294967295 > 116) {
        var v132 int16 = -1
        _ = v132
    }
    _ = len(v1)
    if v2 {
        v47 = (v47 * v47)
    } else {
        for fi133 := 0; fi133 < 12; fi133++ {
            v134 := GenS11{F0: 42, F1: genFn1(), F2: (-(v3))}
            var v135 int32 = (^(v47))
            _ = fi133
            _ = v134
            _ = v135
        }
        for fi136 := 0; fi136 < 9; fi136++ {
            v137 := GenS4{F0: 2, F1: 1}
            for fi138 := 0; fi138 < 5; fi138++ {
                v139 := GenS12{F0: v2, F1: 2, F2: 128}
                fmt.Println(v2, v47, v47)
                _ = fi138
                _ = v139
            }
            _ = fi136
            _ = v137
        }
        if v2 {
            var v140 float64 = 1e10
            if (!(0.1 <= (1e10 / (v140 + float64((3.14159 * 1.5)))))) {
                v140 -= v140
                if ((v47 | (-(v47))) == (v47 & 0)) {
                    fmt.Println(v3)
                    var v141 uint = 0
                    v142 := GenS3{F0: "foo bar", F1: (v6 * v6)}
                    _ = v141
                    _ = v142
                }
                v143 := GenS3{F0: "\n\t", F1: (v6 - v6)}
                _ = v143
            } else {
                v144 := GenS11{F0: 2, F1: v3, F2: v3}
                var v145 uint = 419
                _ = len(v4)
                _ = v144
                _ = v145
            }
            _ = v140
        } else {
            if v2 {
                v3 = -9223372036854775808
            } else {
                var v146 float64 = 1e10
                _ = len(v4)
                _ = v146
            }
            fmt.Println(v47, v6)
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v47:", v47)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v49
    _ = v5
    _ = v6
}
