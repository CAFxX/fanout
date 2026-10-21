package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32 }
type GenS2 struct { F0 int16; F1 bool }
type GenS3 struct { F0 int8; F1 bool }
type GenS4 struct { F0 int32; F1 bool }
type GenS5 struct { F0 uint16; F1 uint64 }
type GenS6 struct { F0 uint32; F1 uint32 }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 int16 }
type GenS9 struct { F0 int64 }
type GenS10 struct { F0 string; F1 uint16 }
type GenS11 struct { F0 int16; F1 uint8; F2 uint16 }
type GenS12 struct { F0 string; F1 uint16 }
type GenS13 struct { F0 int64; F1 string }
type GenS14 struct { F0 int64 }
type GenS15 struct { F0 bool; F1 int64; F2 int8 }
type GenS16 struct { F0 int8; F1 string }
type GenS17 struct { F0 uint64 }
type GenS18 struct { F0 uint8; F1 int32; F2 uint32 }
type GenS19 struct { F0 bool; F1 bool; F2 uint8 }
type GenS20 struct { F0 string; F1 int8; F2 int64 }
type GenS21 struct { F0 bool }
type GenS22 struct { F0 string; F1 uint }
type GenS23 struct { F0 int; F1 string }
type GenS24 struct { F0 int64; F1 int16; F2 uint32 }
type GenS25 struct { F0 int32; F1 uint16 }
type GenS26 struct { F0 uint64 }
type GenS27 struct { F0 uint8 }
type GenS28 struct { F0 bool; F1 uint }
type GenS29 struct { F0 int64; F1 int }
type GenS30 struct { F0 bool; F1 int32 }
type GenS31 struct { F0 bool; F1 uint16 }
type GenS32 struct { F0 int; F1 int32; F2 int32 }

func main() {
    var v1 uint64 = 1
    v2 := make(map[int]string)
    v3 := make(map[string]string)
    v4 := make(map[string]string)
    v1 = (v1 + (v1 >> ((65535) & 63)))
    var v5 uint8 = 1
    v5 -= (v5 + v5)
    v6 := GenS1{F0: -1, F1: 4294967295}
    for fi7 := 0; fi7 < 19; fi7++ {
        v1 = 2
        _ = fi7
    }
    for fi8 := 0; fi8 < 2; fi8++ {
        v1 = 2
        _ = fi8
    }
    if (((1e10 + ((-2.25 / 1.0) / 1.0)) != 3.14159) && false) {
        for fi9 := 0; fi9 < 10; fi9++ {
            var v10 uint32 = 4294967295
            fmt.Println(v1)
            _ = fi9
            _ = v10
        }
    }
    v11 := GenS2{F0: -32768, F1: (true && (28 != 41))}
    _ = len(v3)
    switch "a" {
    case "foo bar":
        var v12 uint = 1
        _ = v12
    case "":
        var v13 int16 = -1
        _ = v13
    case "\n\t":
        for fi14 := 0; fi14 < 15; fi14++ {
            if true {
                v15 := GenS3{F0: 1, F1: false}
                if true {
                    if true {
                        for fi16 := 0; fi16 < 20; fi16++ {
                            v17 := GenS4{F0: 0, F1: true}
                            _ = fi16
                            _ = v17
                        }
                        if (2147483647 >= 0) {
                            v18 := GenS3{F0: -1, F1: true}
                            v19 := GenS5{F0: 107, F1: v1}
                            for fi20 := 0; fi20 < 12; fi20++ {
                                v1 = v1
                                _ = fi20
                            }
                            _ = v18
                            _ = v19
                        } else {
                            var v21 uint64 = v1
                            var v22 string = ""
                            v21 += 1
                            _ = v21
                            _ = v22
                        }
                    } else {
                        v23 := GenS4{F0: 2147483647, F1: (-29 == 42)}
                        _ = v23
                    }
                    fmt.Println(v1, v5, v1)
                    v4["\n\t"] = ("hello" + ("a" + "hello"))
                } else {
                    v24 := []int{(fi14 << ((4294967295) & 63)), fi14, (fi14 / ((fi14) | 1)), fi14}
                    switch "hello" {
                    case "\n\t":
                        if true {
                            if false {
                                switch fi14 {
                                case 911:
                                    fmt.Println(fi14, v1)
                                case 9223372036854775807:
                                    v5 = v5
                                case 4294967295:
                                    v25 := GenS4{F0: 2147483647, F1: false}
                                    _ = v25
                                default:
                                    if false {
                                        for fi26 := 0; fi26 < 11; fi26++ {
                                            if (!true) {
                                                fmt.Println(v5, fi26, fi26)
                                                fmt.Println(fi26, v1)
                                                fmt.Println(v5, v1, fi14)
                                                fmt.Println(fi14)
                                                var v27 uint8 = (^(v5))
                                                v5 -= (^(v5))
                                                switch fi26 {
                                                case 4294967295:
                                                    switch (fi26 / ((-1) | 1)) {
                                                    case 2147483647:
                                                        var v28 int32 = 0
                                                        _ = v28
                                                    }
                                                case -492:
                                                    if (true && (100 >= 0)) {
                                                        switch fi14 {
                                                        case 32767:
                                                            fmt.Println(v27)
                                                        case -9223372036854775808:
                                                            if false {
                                                                var v29 string = "\n\t"
                                                                v30 := GenS6{F0: 0, F1: 1}
                                                                _ = v29
                                                                _ = v30
                                                            }
                                                        default:
                                                            var v31 uint8 = v27
                                                            _ = v31
                                                        }
                                                        v27 = v27
                                                    } else {
                                                        if false {
                                                            v32 := GenS1{F0: (fi14 & 766), F1: 744}
                                                            _ = v32
                                                        }
                                                        v27 += 42
                                                    }
                                                case 32768:
                                                    var v33 int16 = 546
                                                    _ = v33
                                                default:
                                                    v5 = v27
                                                }
                                                var v34 int64 = 9223372036854775807
                                                _ = v27
                                                _ = v34
                                            }
                                            if true {
                                                if len(v24) >= 2 {
                                                    v24 = v24[1:len(v24)]
                                                }
                                                if false {
                                                    var v35 int16 = -1
                                                    v36 := GenS3{F0: 1, F1: false}
                                                    _ = v35
                                                    _ = v36
                                                } else {
                                                    v37 := GenS1{F0: fi14, F1: 65535}
                                                    _ = v37
                                                }
                                                var v38 float64 = 3.14159
                                                _ = v38
                                            } else {
                                                switch -536 {
                                                case 22:
                                                    switch 0 {
                                                    case 624:
                                                        switch ("\n\t" + (("a" + (("hello" + ("foo bar" + "hello")) + ((("\n\t" + (("foo bar" + "foo bar") + ("" + (("a" + "hello") + "foo bar")))) + "foo bar") + "foo bar"))) + "")) {
                                                        case "a":
                                                            v5 = v5
                                                        case "":
                                                            v39 := GenS5{F0: 255, F1: v1}
                                                            _ = v39
                                                        case "foo bar":
                                                            var v40 uint64 = (v1 << ((4294967295) & 63))
                                                            _ = v40
                                                        }
                                                    }
                                                case 128:
                                                    fmt.Println(fi26, fi14, fi14)
                                                    var v41 float32 = 0.0
                                                    var v42 int32 = 2147483647
                                                    _ = v41
                                                    _ = v42
                                                case 42:
                                                    v43 := GenS7{F0: 18446744073709551615}
                                                    _ = v43
                                                default:
                                                    v5 = v5
                                                }
                                            }
                                            _ = fi26
                                        }
                                        v5 += v5
                                        var v44 int64 = -985
                                        _ = v44
                                    }
                                }
                                var v45 uint64 = (v1 / ((929) | 1))
                                _ = v45
                            } else {
                                if (((v5 - (v5 >> ((42) & 63))) > v5) || false) {
                                    for fi46 := 0; fi46 < 15; fi46++ {
                                        var v47 int = (fi14 ^ (fi14 ^ 2147483647))
                                        _ = fi46
                                        _ = v47
                                    }
                                    if (49 >= uint32(v5)) {
                                        switch -134 {
                                        case 724:
                                            v1 = v1
                                        case -1:
                                            switch "\n\t" {
                                            case "hello":
                                                var v48 uint8 = v5
                                                _ = v48
                                            case "foo bar":
                                                var v49 int16 = 0
                                                _ = v49
                                            default:
                                                for fi50 := 0; fi50 < 13; fi50++ {
                                                    if false {
                                                        fmt.Println(fi14, fi14, fi50)
                                                        v51 := GenS8{F0: 1000}
                                                        _ = v51
                                                    } else {
                                                        if (-9223372036854775808 < 149) {
                                                            v1 = v1
                                                        }
                                                        switch "" {
                                                        case "\n\t":
                                                            v52 := GenS9{F0: -1}
                                                            _ = v52
                                                        case "hello":
                                                            v1 = v1
                                                        default:
                                                            if (!((fi14 ^ fi50) == fi14)) {
                                                                var v53 uint16 = 1
                                                                _ = v53
                                                            }
                                                        }
                                                    }
                                                    switch "\n\t" {
                                                    case "\n\t":
                                                        if (1.5 < ((1.5 * 1.5) - 3.14159)) {
                                                            _ = len(v3)
                                                            fmt.Println(fi50, v1)
                                                        }
                                                    default:
                                                        fmt.Println(v1)
                                                    }
                                                    _ = fi50
                                                }
                                            }
                                        }
                                        var v54 int = (fi14 + 42)
                                        fmt.Println(fi14)
                                        _ = v54
                                    } else {
                                        v1 += v1
                                        if true {
                                            v5 = 75
                                            if true {
                                                var v55 int64 = -1
                                                var v56 int8 = 127
                                                for fi57 := 0; fi57 < 8; fi57++ {
                                                    if true {
                                                        v55 = -9223372036854775808
                                                        switch fi14 {
                                                        case -9223372036854775808:
                                                            v4["foo bar"] = "\n\t"
                                                        default:
                                                            v58 := GenS1{F0: (fi14 >> ((65535) & 63)), F1: 233}
                                                            _ = v58
                                                        }
                                                    }
                                                    v1 = v1
                                                    var v59 uint8 = 2
                                                    _ = fi57
                                                    _ = v59
                                                }
                                                _ = v55
                                                _ = v56
                                            }
                                        }
                                    }
                                    v60 := GenS8{F0: 1000}
                                    _ = v60
                                }
                            }
                            v5 = v5
                            if len(v24) > 0 {
                                v24[0] = fi14
                            }
                        } else {
                            var v61 bool = true
                            if (v61 && (((!v61) || v61) || (241 <= 423))) {
                                _ = len(v24)
                                for fi62 := 0; fi62 < 12; fi62++ {
                                    fmt.Println(fi62, v5)
                                    _ = fi62
                                }
                                var v63 int16 = 42
                                _ = v63
                            } else {
                                fmt.Println(v61, v1, v5)
                                if (!false) {
                                    switch "\n\t" {
                                    case "hello":
                                        for fi64 := 0; fi64 < 18; fi64++ {
                                            v4["foo bar"] = "foo bar"
                                            _ = fi64
                                        }
                                    default:
                                        fmt.Println(v1, v1)
                                    }
                                    for fi65 := 0; fi65 < 16; fi65++ {
                                        v66 := GenS10{F0: "a", F1: 255}
                                        if (!(0 == int8(v5))) {
                                            var v67 uint64 = 333
                                            _ = v67
                                        } else {
                                            v4[""] = "hello"
                                            v1 -= v1
                                            v1 = (v1 % (((v1 * 1)) | 1))
                                            v68 := GenS11{F0: -23, F1: (v5 / ((243) | 1)), F2: 928}
                                            fmt.Println(v1, v61)
                                            _ = v68
                                        }
                                        _ = fi65
                                        _ = v66
                                    }
                                }
                                if ((9223372036854775807 == 890) && (32767 > -1)) {
                                    switch "\n\t" {
                                    case "":
                                        var v69 float64 = -2.25
                                        _ = v69
                                    case "foo bar":
                                        for fi70 := 0; fi70 < 7; fi70++ {
                                            if v61 {
                                                if (!v61) {
                                                    v61 = (986 != 255)
                                                    v61 = v61
                                                } else {
                                                    fmt.Println(v61)
                                                    var v71 uint32 = 2147483648
                                                    switch "" {
                                                    case "":
                                                        v72 := GenS1{F0: fi14, F1: (v71 % (((v71 | 0)) | 1))}
                                                        v73 := GenS6{F0: v71, F1: 42}
                                                        var v74 uint16 = uint16(v71)
                                                        _ = v72
                                                        _ = v73
                                                        _ = v74
                                                    case "a":
                                                        if ((387 <= 2) || false) {
                                                            switch 0 {
                                                            case 0:
                                                                var v75 string = "\n\t"
                                                                v76 := GenS5{F0: 42, F1: v1}
                                                                v77 := GenS12{F0: v75, F1: uint16(v5)}
                                                                _ = v75
                                                                _ = v76
                                                                _ = v77
                                                            case 34:
                                                                v3[(("foo bar" + "\n\t") + (("a" + "hello") + "foo bar"))] = "foo bar"
                                                            default:
                                                                v78 := GenS13{F0: 1, F1: (("" + "hello") + ((("foo bar" + (("foo bar" + "foo bar") + "foo bar")) + "hello") + ""))}
                                                                _ = v78
                                                            }
                                                            v79 := GenS14{F0: -6}
                                                            _ = v79
                                                        } else {
                                                            var v80 int32 = 42
                                                            switch "hello" {
                                                            case "":
                                                                if v61 {
                                                                    v1 += (v1 | (v1 ^ v1))
                                                                    var v81 float64 = 1.5
                                                                    _ = v81
                                                                } else {
                                                                    fmt.Println(v71)
                                                                }
                                                            default:
                                                                switch 2147483647 {
                                                                case 127:
                                                                    v61 = ((^(fi70)) <= 523)
                                                                case 1:
                                                                    if (42 == -50) {
                                                                        v82 := GenS15{F0: v61, F1: 0, F2: 100}
                                                                        if (!v61) {
                                                                            v83 := GenS16{F0: 99, F1: "foo bar"}
                                                                            _ = v83
                                                                        } else {
                                                                            switch fi70 {
                                                                            case -2147483648:
                                                                                switch "" {
                                                                                case "\n\t":
                                                                                    fmt.Println(fi70, v61, v71)
                                                                                    var v84 uint16 = 981
                                                                                    fmt.Println(v61)
                                                                                    _ = v84
                                                                                case "foo bar":
                                                                                    if (2 == 390) {
                                                                                        v85 := GenS2{F0: int16(v1), F1: (!(int16(v5) < 42))}
                                                                                        switch fi14 {
                                                                                        case 32768:
                                                                                            v61 = v61
                                                                                            v71 = (v71 * 451)
                                                                                            v61 = true
                                                                                            v61 = v61
                                                                                        case 2147483647:
                                                                                            var v86 float64 = 1.5
                                                                                            _ = v86
                                                                                        default:
                                                                                            v1 = 42
                                                                                        }
                                                                                        switch (fi14 % (((fi70 * (^(fi14)))) | 1)) {
                                                                                        case 438:
                                                                                            if (!(0 <= v80)) {
                                                                                                var v87 bool = (!false)
                                                                                                fmt.Println(v5)
                                                                                                switch (("\n\t" + ((("\n\t" + "foo bar") + "a") + "\n\t")) + (("hello" + "") + ((("a" + "hello") + "hello") + ""))) {
                                                                                                case "\n\t":
                                                                                                    v88 := GenS17{F0: v1}
                                                                                                    _ = v88
                                                                                                case "":
                                                                                                    v89 := GenS6{F0: (^(v71)), F1: 0}
                                                                                                    _ = v89
                                                                                                }
                                                                                                fmt.Println(v80, fi14)
                                                                                                _ = v87
                                                                                            } else {
                                                                                                if v61 {
                                                                                                    v90 := GenS15{F0: (3.14159 <= (-2.25 - 0.0)), F1: 9223372036854775807, F2: int8(v80)}
                                                                                                    _ = v90
                                                                                                }
                                                                                                var v91 int64 = 348
                                                                                                v92 := GenS3{F0: 100, F1: v61}
                                                                                                _ = v91
                                                                                                _ = v92
                                                                                            }
                                                                                        case -839:
                                                                                            var v93 uint16 = 32767
                                                                                            _ = v93
                                                                                        default:
                                                                                            if len(v24) > 0 {
                                                                                                v24[0] = (-(fi70))
                                                                                            }
                                                                                        }
                                                                                        v5 = 127
                                                                                        if len(v24) > 0 {
                                                                                            v24[0] = fi14
                                                                                        }
                                                                                        _ = v85
                                                                                    } else {
                                                                                        if len(v24) >= 2 {
                                                                                            v24 = v24[1:len(v24)]
                                                                                        }
                                                                                    }
                                                                                case "a":
                                                                                    v71 = v71
                                                                                default:
                                                                                    fmt.Println(v61, v5)
                                                                                }
                                                                            case 280:
                                                                                _ = len(v24)
                                                                            }
                                                                            fmt.Println(v80, v5)
                                                                            if v61 {
                                                                                v80 = (^(v80))
                                                                                var v94 uint64 = 18446744073709551615
                                                                                var v95 float32 = 0.0
                                                                                v96 := GenS18{F0: v5, F1: 2147483647, F2: (v71 % (((v71 - 778)) | 1))}
                                                                                v1 -= 18446744073709551615
                                                                                _ = v94
                                                                                _ = v95
                                                                                _ = v96
                                                                            }
                                                                        }
                                                                        _ = v82
                                                                    }
                                                                    var v97 int = fi70
                                                                    if len(v24) > 0 {
                                                                        v24[0] = int(v1)
                                                                    }
                                                                    _ = v97
                                                                case 2147483647:
                                                                    v80 = (v80 + -692)
                                                                default:
                                                                    v98 := GenS19{F0: false, F1: (v71 != v71), F2: v5}
                                                                    _ = v98
                                                                }
                                                            }
                                                            _ = v80
                                                        }
                                                    default:
                                                        v99 := GenS20{F0: "a", F1: -96, F2: 42}
                                                        _ = v99
                                                    }
                                                    _ = v71
                                                }
                                                v100 := GenS6{F0: 975, F1: uint32(v5)}
                                                v61 = v61
                                                v101 := GenS21{F0: true}
                                                v102 := GenS22{F0: "", F1: 1}
                                                _ = v100
                                                _ = v101
                                                _ = v102
                                            } else {
                                                if len(v24) >= 2 {
                                                    v24 = v24[1:len(v24)]
                                                }
                                                if ((42 != 1) && false) {
                                                    v24 = append(v24, fi70)
                                                    var v103 uint32 = 93
                                                    v5 -= 1
                                                    v1 -= 0
                                                    v61 = (v61 && v61)
                                                    v104 := GenS14{F0: 0}
                                                    _ = v103
                                                    _ = v104
                                                }
                                                var v105 int8 = -128
                                                _ = v105
                                            }
                                            _ = fi70
                                        }
                                    case "\n\t":
                                        var v106 string = "\n\t"
                                        _ = v106
                                    default:
                                        fmt.Println(v61, fi14)
                                    }
                                    if false {
                                        v1 -= 42
                                        if (!v61) {
                                            for fi107 := 0; fi107 < 7; fi107++ {
                                                v5 += v5
                                                var v108 int16 = 1000
                                                if v61 {
                                                    switch "\n\t" {
                                                    case "a":
                                                        if v61 {
                                                            var v109 uint = 2
                                                            var v110 int16 = 246
                                                            v1 += 201
                                                            var v111 uint64 = (v1 + uint64(v5))
                                                            _ = v109
                                                            _ = v110
                                                            _ = v111
                                                        }
                                                    case "\n\t":
                                                        var v112 uint64 = 16
                                                        _ = v112
                                                    default:
                                                        if (32768 != (-(fi107))) {
                                                            v61 = (((v61 && (0.0 >= (((1.5 / (3.14159 / 3.14159)) * float32(1.5)) * (-2.25 * (-2.25 * (0.0 * (0.1 / 1.5))))))) && (42 <= 100)) || true)
                                                            v113 := GenS23{F0: 2147483647, F1: ("hello" + "")}
                                                            var v114 float32 = (3.14159 - ((0.0 * ((1e10 + 0.1) + -2.25)) * 1e10))
                                                            _ = v113
                                                            _ = v114
                                                        }
                                                    }
                                                    if v61 {
                                                        var v115 uint32 = 172
                                                        v61 = (!(1 != 510))
                                                        v1 += v1
                                                        _ = v115
                                                    }
                                                    if ((0 == (v5 ^ v5)) || ((v1 % ((v1) | 1)) >= (v1 / ((4294967295) | 1)))) {
                                                        v116 := GenS7{F0: 65535}
                                                        v108 = (v108 % ((v108) | 1))
                                                        v117 := v4[""]
                                                        if (128 <= fi14) {
                                                            var v118 uint32 = 2
                                                            _ = v118
                                                        } else {
                                                            v61 = true
                                                            if len(v24) >= 2 {
                                                                v24 = v24[1:len(v24)]
                                                            }
                                                            v119 := GenS24{F0: -406, F1: v108, F2: 0}
                                                            _ = v119
                                                        }
                                                        _ = v116
                                                        _ = v117
                                                    }
                                                } else {
                                                    switch 32767 {
                                                    case 42:
                                                        v1 -= v1
                                                    }
                                                }
                                                if v61 {
                                                    v108 += (v108 & v108)
                                                    var v120 uint8 = (v5 & 128)
                                                    _ = v120
                                                } else {
                                                    v108 -= v108
                                                    v5 -= v5
                                                }
                                                if len(v24) >= 2 {
                                                    v24 = v24[1:len(v24)]
                                                }
                                                v1 = v1
                                                v108 = v108
                                                var v121 float32 = 1e10
                                                _ = len(v4)
                                                var v122 uint = uint(v1)
                                                _ = fi107
                                                _ = v108
                                                _ = v121
                                                _ = v122
                                            }
                                            switch -729 {
                                            case -847:
                                                if (!v61) {
                                                    v61 = v61
                                                    var v123 uint = 551
                                                    v123 = (^(v123))
                                                    _ = v123
                                                } else {
                                                    fmt.Println(v1, v1, v5)
                                                }
                                            case -894:
                                                if v61 {
                                                    if v61 {
                                                        for fi124 := 0; fi124 < 10; fi124++ {
                                                            var v125 int16 = 179
                                                            fmt.Println(v61)
                                                            if (float64((1.5 + -2.25)) < 1e10) {
                                                                if true {
                                                                    v61 = v61
                                                                    if (((v125 | v125) <= (v125 & -32768)) && v61) {
                                                                        v1 -= 18446744073709551615
                                                                    } else {
                                                                        var v126 int = fi14
                                                                        v127 := v4[("foo bar" + (("hello" + ("\n\t" + "foo bar")) + ("" + (("\n\t" + "a") + "\n\t"))))]
                                                                        _ = v126
                                                                        _ = v127
                                                                    }
                                                                }
                                                                var v128 int8 = 1
                                                                v129 := v4["\n\t"]
                                                                _ = v128
                                                                _ = v129
                                                            } else {
                                                                var v130 int64 = 395
                                                                switch (fi14 / ((127) | 1)) {
                                                                case -2147483648:
                                                                    v61 = (false && (656 != 1))
                                                                }
                                                                fmt.Println(fi124, v5)
                                                                if v61 {
                                                                    v1 += v1
                                                                    var v131 uint8 = (v5 ^ v5)
                                                                    _ = v131
                                                                }
                                                                _ = v130
                                                            }
                                                            _ = fi124
                                                            _ = v125
                                                        }
                                                        for fi132 := 0; fi132 < 17; fi132++ {
                                                            if (!(-240 >= 0)) {
                                                                var v133 string = ("\n\t" + "")
                                                                _ = v133
                                                            }
                                                            fmt.Println(fi14)
                                                            var v134 uint16 = 255
                                                            _ = fi132
                                                            _ = v134
                                                        }
                                                    } else {
                                                        v1 = (v1 * (v1 % ((18446744073709551615) | 1)))
                                                        v5 += v5
                                                    }
                                                    var v135 uint32 = 2147483647
                                                    var v136 int = fi14
                                                    _ = v135
                                                    _ = v136
                                                }
                                            }
                                            fmt.Println(fi14, v5, v61)
                                        }
                                    } else {
                                        for fi137 := 0; fi137 < 4; fi137++ {
                                            if (((32767 == 594) || false) && (1 == 0)) {
                                                if (!v61) {
                                                    var v138 float32 = 0.1
                                                    v24 = append(v24, -797)
                                                    switch ("" + "") {
                                                    case "hello":
                                                        if v61 {
                                                            if (((-2.25 - 3.14159) - 0.1) != 3.14159) {
                                                                var v139 int8 = 1
                                                                var v140 string = ("" + "\n\t")
                                                                var v141 int = 0
                                                                v4[(v140 + v140)] = (v140 + (("foo bar" + v140) + v140))
                                                                v5 -= (v5 + (v5 * v5))
                                                                if len(v24) >= 2 {
                                                                    v24 = v24[1:len(v24)]
                                                                }
                                                                v2[fi14] = (((v140 + (v140 + "foo bar")) + v140) + v140)
                                                                v5 += v5
                                                                _ = v139
                                                                _ = v140
                                                                _ = v141
                                                            }
                                                            v1 = uint64(v5)
                                                        }
                                                    default:
                                                        switch (fi14 / (((fi14 / (((fi14 | 860)) | 1))) | 1)) {
                                                        case 32768:
                                                            switch "a" {
                                                            case "\n\t":
                                                                if len(v24) > 0 {
                                                                    v24[0] = fi137
                                                                }
                                                            case "":
                                                                v1 = v1
                                                            case "a":
                                                                delete(v2, (fi137 * 32768))
                                                            }
                                                        case -2147483648:
                                                            if (4294967295 <= 2) {
                                                                if (((fi137 | 4294967295) <= fi14) && false) {
                                                                    if true {
                                                                        var v142 uint64 = (^(v1))
                                                                        if (v61 && true) {
                                                                            var v143 int64 = -330
                                                                            _ = v143
                                                                        } else {
                                                                            v1 = (^(v1))
                                                                            if (v138 >= ((1.5 * 3.14159) - ((1.5 - (((3.14159 * 3.14159) - v138) - 0.0)) + 0.1))) {
                                                                                v144 := GenS15{F0: (v1 < v142), F1: -9223372036854775808, F2: -53}
                                                                                switch "a" {
                                                                                case "":
                                                                                    v1 = (v1 | v142)
                                                                                }
                                                                                v61 = v61
                                                                                v145 := GenS4{F0: -60, F1: (("a" == "hello") || (!false))}
                                                                                _ = v144
                                                                                _ = v145
                                                                            } else {
                                                                                if false {
                                                                                    var v146 float32 = ((((float32((0.0 - (0.1 * 0.0))) / (v138 * float32(0.0))) * (((v138 + ((v138 + v138) - v138)) / (v138 + v138)) - 1.5)) / v138) - v138)
                                                                                    delete(v3, "")
                                                                                    var v147 bool = v61
                                                                                    _ = v146
                                                                                    _ = v147
                                                                                }
                                                                            }
                                                                            _ = len(v24)
                                                                            if (!((("\n\t" + "") + "") > "foo bar")) {
                                                                                v24 = append(v24, fi14)
                                                                            }
                                                                            var v148 int8 = -117
                                                                            _ = v148
                                                                        }
                                                                        _ = v142
                                                                    }
                                                                    switch "hello" {
                                                                    case "hello":
                                                                        v1 = v1
                                                                    default:
                                                                        if true {
                                                                            var v149 int64 = -209
                                                                            var v150 bool = false
                                                                            if (695 == 2) {
                                                                                v151 := GenS25{F0: 42, F1: 984}
                                                                                switch ("foo bar" + "foo bar") {
                                                                                case "":
                                                                                    v150 = v150
                                                                                case "hello":
                                                                                    v4[(("foo bar" + "a") + (("a" + ("" + "hello")) + ("\n\t" + "hello")))] = "\n\t"
                                                                                case "\n\t":
                                                                                    v152 := v2[(fi14 & fi137)]
                                                                                    _ = v152
                                                                                }
                                                                                _ = v151
                                                                            } else {
                                                                                var v153 uint64 = 0
                                                                                v149 = v149
                                                                                fmt.Println(v61, v61)
                                                                                fmt.Println(fi137)
                                                                                _ = v153
                                                                            }
                                                                            if (42 < 217) {
                                                                                _ = len(v4)
                                                                            }
                                                                            switch ("foo bar" + "a") {
                                                                            case "a":
                                                                                switch "hello" {
                                                                                case "":
                                                                                    var v154 uint8 = v5
                                                                                    _ = v154
                                                                                default:
                                                                                    v155 := GenS11{F0: 649, F1: v5, F2: 42}
                                                                                    _ = v155
                                                                                }
                                                                            case "hello":
                                                                                v149 -= v149
                                                                            case "":
                                                                                if (4294967295 >= (v1 * (v1 - (v1 % (((v1 | (v1 ^ (v1 << ((255) & 63))))) | 1))))) {
                                                                                    var v156 uint32 = 706
                                                                                    _ = v156
                                                                                }
                                                                            }
                                                                            if ((2 != 665) || (-1 != 1)) {
                                                                                var v157 int64 = (v149 + (v149 | 1))
                                                                                _ = v157
                                                                            }
                                                                            fmt.Println(fi14, v149)
                                                                            _ = v149
                                                                            _ = v150
                                                                        }
                                                                    }
                                                                }
                                                            } else {
                                                                v1 = v1
                                                                switch (fi137 / ((fi137) | 1)) {
                                                                case 127:
                                                                    v158 := v3["a"]
                                                                    _ = v158
                                                                default:
                                                                    v159 := v3["a"]
                                                                    v159 = v159
                                                                    _ = v159
                                                                }
                                                                v160 := GenS2{F0: 0, F1: v61}
                                                                _ = v160
                                                            }
                                                        case -9223372036854775808:
                                                            v138 += (-2.25 + v138)
                                                        default:
                                                            v5 = (v5 | v5)
                                                        }
                                                        v1 = (v1 % ((v1) | 1))
                                                    }
                                                    var v161 int16 = 42
                                                    v138 = v138
                                                    _ = v138
                                                    _ = v161
                                                }
                                                if ((-(fi137)) >= fi137) {
                                                    delete(v2, (fi14 & (fi14 + fi137)))
                                                    delete(v3, "hello")
                                                }
                                            } else {
                                                fmt.Println(fi14, v5, v1)
                                                v5 = 42
                                                if v61 {
                                                    var v162 int32 = 1
                                                    _ = len(v3)
                                                    var v163 int = -9223372036854775808
                                                    _ = v162
                                                    _ = v163
                                                } else {
                                                    var v164 int16 = -553
                                                    _ = v164
                                                }
                                            }
                                            switch "hello" {
                                            case "a":
                                                var v165 uint64 = v1
                                                _ = v165
                                            }
                                            v166 := GenS24{F0: -1, F1: 1, F2: 0}
                                            _ = fi137
                                            _ = v166
                                        }
                                        fmt.Println(fi14, fi14)
                                        if (446 < -832) {
                                            v1 = v1
                                            v61 = (!(!v61))
                                        } else {
                                            v167 := GenS26{F0: (^(v1))}
                                            switch "\n\t" {
                                            case "hello":
                                                var v168 uint64 = (v1 % (((v1 % (((v1 % ((462) | 1))) | 1))) | 1))
                                                _ = v168
                                            default:
                                                if (0 >= 42) {
                                                    delete(v3, ("" + ((("" + "\n\t") + "hello") + "")))
                                                    if v61 {
                                                        v61 = true
                                                    }
                                                }
                                            }
                                            if v61 {
                                                switch (^(fi14)) {
                                                case -409:
                                                    fmt.Println(v61)
                                                case 4294967295:
                                                    var v169 int16 = 0
                                                    _ = v169
                                                case 136:
                                                    _ = len(v3)
                                                }
                                                switch (fi14 + fi14) {
                                                case 32767:
                                                    var v170 int32 = -270
                                                    _ = v170
                                                case 1:
                                                    v5 = uint8(v1)
                                                case -1:
                                                    var v171 int16 = 552
                                                    _ = v171
                                                }
                                            }
                                            _ = v167
                                        }
                                    }
                                    switch "\n\t" {
                                    case "a":
                                        var v172 float64 = (float64(float32((((0.0 + 3.14159) + 3.14159) + 0.1))) / -2.25)
                                        _ = v172
                                    default:
                                        v1 = v1
                                    }
                                }
                            }
                            _ = v61
                        }
                    }
                    fmt.Println(v1, fi14)
                    _ = v24
                }
                var v173 float32 = 0.0
                _ = v15
                _ = v173
            } else {
                if false {
                    switch 9223372036854775807 {
                    case -1:
                        v1 = (^(v1))
                    default:
                        for fi174 := 0; fi174 < 14; fi174++ {
                            switch fi14 {
                            case -2147483648:
                                v1 -= 18446744073709551615
                            case 9223372036854775807:
                                v1 = uint64(v5)
                            }
                            fmt.Println(fi14, v1, fi14)
                            v5 += (^(v5))
                            _ = fi174
                        }
                    }
                    v1 = (^(v1))
                    v175 := GenS24{F0: 1, F1: 32767, F2: 209}
                    _ = v175
                } else {
                    v176 := []int{fi14, (fi14 ^ 42), (fi14 % ((fi14) | 1))}
                    switch ("\n\t" + "\n\t") {
                    case "\n\t":
                        if false {
                            var v177 int32 = 974
                            fmt.Println(fi14, v177)
                            v4["foo bar"] = ((("\n\t" + "hello") + "a") + "")
                            _ = v177
                        }
                    case "hello":
                        if len(v176) > 0 {
                            v176[0] = (-(fi14))
                        }
                    }
                    _ = v176
                }
            }
            v178 := GenS17{F0: 0}
            _ = fi14
            _ = v178
        }
    }
    switch "" {
    case "hello":
        var v179 float64 = ((-2.25 + -2.25) * (1e10 + 1.5))
        _ = v179
    }
    for fi180 := 0; fi180 < 18; fi180++ {
        v181 := GenS17{F0: v1}
        var v182 float32 = 1e10
        v5 += (v5 % (((v5 & v5)) | 1))
        _ = fi180
        _ = v181
        _ = v182
    }
    if false {
        switch -9223372036854775808 {
        case -2147483648:
            v5 = (v5 + v5)
        default:
            if (32767 <= 626) {
                switch -158 {
                case 32767:
                    if (!true) {
                        var v183 string = "foo bar"
                        if true {
                            for fi184 := 0; fi184 < 17; fi184++ {
                                v1 = (v1 * 0)
                                fmt.Println(v1, v1)
                                var v185 float32 = 0.1
                                _ = fi184
                                _ = v185
                            }
                            var v186 int32 = 679
                            if (65535 > 32767) {
                                if ((true || true) || (v1 != 2)) {
                                    v186 = 42
                                    v187 := GenS27{F0: 0}
                                    _ = v187
                                }
                                fmt.Println(v5, v1)
                            }
                            _ = v186
                        } else {
                            if true {
                                v183 = v183
                            } else {
                                v1 -= v1
                            }
                        }
                        var v188 bool = (!false)
                        _ = v183
                        _ = v188
                    }
                case -1:
                    var v189 bool = true
                    _ = v189
                case 128:
                    switch 22 {
                    case 500:
                        if false {
                            var v190 uint64 = 42
                            for fi191 := 0; fi191 < 7; fi191++ {
                                switch "" {
                                case "\n\t":
                                    v192 := []int{fi191}
                                    _ = v192
                                case "a":
                                    var v193 bool = (false || false)
                                    _ = v193
                                }
                                _ = fi191
                            }
                            _ = v190
                        }
                    case -9223372036854775808:
                        fmt.Println(v5)
                    default:
                        if (0 <= -155) {
                            var v194 bool = false
                            v194 = true
                            v1 = v1
                            _ = v194
                        } else {
                            fmt.Println(v5)
                            if false {
                                v195 := GenS20{F0: "\n\t", F1: -1, F2: 9223372036854775807}
                                _ = v195
                            }
                            v196 := GenS21{F0: (true && true)}
                            _ = v196
                        }
                    }
                }
                var v197 int8 = -128
                var v198 bool = false
                _ = v197
                _ = v198
            } else {
                for fi199 := 0; fi199 < 5; fi199++ {
                    var v200 uint16 = 0
                    _ = fi199
                    _ = v200
                }
                v201 := []int{-2147483648, 4294967295}
                _ = v201
            }
        }
        for fi202 := 0; fi202 < 4; fi202++ {
            for fi203 := 0; fi203 < 10; fi203++ {
                if true {
                    var v204 string = "foo bar"
                    _ = v204
                } else {
                    v1 = 4294967295
                    switch 4294967295 {
                    case 128:
                        _ = len(v2)
                    case -2147483648:
                        v205 := []int{-583}
                        _ = v205
                    default:
                        switch "\n\t" {
                        case "foo bar":
                            v206 := []int{(fi202 % (((fi202 / ((308) | 1))) | 1))}
                            _ = v206
                        case "\n\t":
                            var v207 int8 = 0
                            _ = v207
                        case "a":
                            var v208 uint = 263
                            _ = v208
                        default:
                            if (true && (!(false && true))) {
                                v209 := GenS28{F0: (!(("hello" + ((("\n\t" + "\n\t") + "hello") + "\n\t")) < "")), F1: 255}
                                _ = v209
                            }
                        }
                    }
                    v1 -= v1
                }
                if false {
                    v210 := GenS29{F0: 9223372036854775807, F1: fi203}
                    v5 = (v5 >> ((83) & 63))
                    _ = v210
                }
                v211 := GenS30{F0: ((float64(3.14159) + float64(-2.25)) < 1.5), F1: 2147483647}
                _ = fi203
                _ = v211
            }
            for fi212 := 0; fi212 < 6; fi212++ {
                var v213 uint8 = 156
                var v214 int32 = 20
                if false {
                    if (0.0 < ((float64(-2.25) - float64(1.5)) + 3.14159)) {
                        switch fi212 {
                        case 1:
                            v213 = v5
                        }
                        switch 889 {
                        case 128:
                            var v215 bool = false
                            v5 = (^(v5))
                            _ = v215
                        case 4294967295:
                            v216 := GenS31{F0: true, F1: 1}
                            _ = v216
                        case -816:
                            var v217 float32 = -2.25
                            _ = v217
                        default:
                            v213 = v5
                        }
                    }
                    v218 := GenS32{F0: 362, F1: v214, F2: int32(v5)}
                    _ = v218
                }
                _ = fi212
                _ = v213
                _ = v214
            }
            _ = fi202
        }
        var v219 string = ("\n\t" + ("hello" + ("a" + "")))
        _ = v219
    } else {
        v220 := GenS9{F0: 0}
        _ = v220
    }
    fmt.Println(v1, v1)
    if false {
        v1 += 42
        v221 := GenS29{F0: 42, F1: 127}
        var v222 int16 = 0
        _ = v221
        _ = v222
    }
    if false {
        switch "\n\t" {
        case "foo bar":
            var v223 float32 = 0.1
            _ = v223
        case "\n\t":
            var v224 bool = true
            _ = v224
        }
    }
    fmt.Println(v1, v1)
    if (false && (1e10 == 0.0)) {
        var v225 int16 = -554
        _ = v225
    } else {
        if true {
            v1 -= (v1 & (v1 + 85))
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
