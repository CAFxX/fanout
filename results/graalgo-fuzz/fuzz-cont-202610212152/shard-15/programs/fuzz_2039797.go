package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64 }
type GenS2 struct { F0 int16; F1 string }
type GenS3 struct { F0 uint; F1 int; F2 uint }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 bool; F1 string; F2 int16 }
type GenS6 struct { F0 int32; F1 uint64; F2 int16 }
type GenS7 struct { F0 string }
type GenS8 struct { F0 int; F1 int }
type GenS9 struct { F0 int8; F1 int16; F2 int16 }
type GenS10 struct { F0 int32 }
type GenS11 struct { F0 uint8 }
type GenS12 struct { F0 int }
type GenS13 struct { F0 int; F1 uint16 }
type GenS14 struct { F0 bool }
type GenS15 struct { F0 int8; F1 uint64 }
type GenS16 struct { F0 int16 }
type GenS17 struct { F0 uint; F1 bool; F2 uint64 }
type GenS18 struct { F0 uint8; F1 uint16; F2 bool }
type GenS19 struct { F0 int16; F1 bool }
type GenS20 struct { F0 int32; F1 uint8; F2 bool }
type GenS21 struct { F0 uint8; F1 uint64; F2 int64 }
type GenS22 struct { F0 bool; F1 uint16; F2 uint }
type GenS23 struct { F0 int32; F1 uint }
type GenS24 struct { F0 uint8 }
type GenS25 struct { F0 string }

func main() {
    var v1 int = 0
    var v2 uint32 = 674
    v3 := make(map[int]string)
    var v4 int64 = 0
    v5 := make(map[uint32]int)
    v6 := make(map[int]bool)
    switch "hello" {
    case "foo bar":
        for fi7 := 0; fi7 < 1; fi7++ {
            if true {
                if true {
                    v1 = fi7
                    switch fi7 {
                    case 32768:
                        delete(v3, (fi7 ^ fi7))
                    }
                    fmt.Println(fi7, v1, fi7)
                } else {
                    var v8 string = "a"
                    _ = v8
                }
                var v9 uint = 2
                var v10 uint8 = 127
                _ = v10
                _ = v9
            } else {
                for fi11 := 0; fi11 < 15; fi11++ {
                    fmt.Println(v4, fi11)
                    var v12 int16 = -330
                    v12 += (-(v12))
                    _ = fi11
                    _ = v12
                }
                var v13 uint8 = 0
                switch 787 {
                case 688:
                    v1 = (^(fi7))
                case 4294967295:
                    switch (v1 % (((fi7 | 2147483647)) | 1)) {
                    case 371:
                        switch (v1 + v1) {
                        case -1:
                            fmt.Println(v1, v1, v1)
                        case 320:
                            v2 -= v2
                        }
                    case 129:
                        if false {
                            var v14 int32 = -1
                            var v15 float32 = 1.5
                            if (42 < 725) {
                                for fi16 := 0; fi16 < 19; fi16++ {
                                    v17 := GenS1{F0: (!((v14 >> ((36) & 63)) != (v14 / (((v14 ^ -2147483648)) | 1)))), F1: 9223372036854775807}
                                    _ = fi16
                                    _ = v17
                                }
                                switch (v1 % ((fi7) | 1)) {
                                case -761:
                                    v4 = (v4 + (v4 + (-(v4))))
                                default:
                                    switch "\n\t" {
                                    case "hello":
                                        for fi18 := 0; fi18 < 4; fi18++ {
                                            var v19 uint64 = 18446744073709551615
                                            if true {
                                                fmt.Println(fi18, v19, v2)
                                            } else {
                                                v4 -= (^(v4))
                                                if true {
                                                    var v20 float32 = v15
                                                    fmt.Println(v4, v2, v14)
                                                    var v21 uint8 = v13
                                                    _ = v20
                                                    _ = v21
                                                } else {
                                                    switch "a" {
                                                    case "a":
                                                        v22 := GenS2{F0: 0, F1: ("" + "foo bar")}
                                                        _ = v22
                                                    }
                                                    fmt.Println(v2, v1)
                                                    v23 := GenS3{F0: 1, F1: (fi7 % (((fi18 % (((^(v1))) | 1))) | 1)), F2: 916}
                                                    _ = v23
                                                }
                                            }
                                            _ = fi18
                                            _ = v19
                                        }
                                    case "\n\t":
                                        fmt.Println(v2, v15)
                                    case "a":
                                        var v24 uint = 16
                                        _ = v24
                                    }
                                }
                            } else {
                                v4 = v4
                                if (0 < -638) {
                                    var v25 uint64 = 0
                                    _ = v25
                                }
                            }
                            _ = v14
                            _ = v15
                        }
                    case 1:
                        if (!false) {
                            for fi26 := 0; fi26 < 1; fi26++ {
                                if true {
                                    var v27 uint16 = 654
                                    var v28 int64 = v4
                                    v29 := GenS4{F0: 0}
                                    _ = v27
                                    _ = v28
                                    _ = v29
                                }
                                v30 := []int{}
                                _ = fi26
                                _ = v30
                            }
                            v31 := []int{v1, v1, (fi7 >> ((309) & 63))}
                            _ = v31
                        }
                    }
                case 9223372036854775807:
                    v4 += (v4 ^ v4)
                default:
                    v3[v1] = ("foo bar" + (("" + "\n\t") + (("hello" + "foo bar") + "\n\t")))
                }
                _ = v13
            }
            _ = fi7
        }
    default:
        fmt.Println(v4, v1)
    }
    var v32 uint64 = 18446744073709551615
    var v33 uint32 = (v2 * v2)
    var v34 bool = true
    for fi35 := 0; fi35 < 6; fi35++ {
        fmt.Println(v34, v33, v1)
        v36 := GenS3{F0: 1, F1: v1, F2: 2}
        _ = fi35
        _ = v36
    }
    v33 += v2
    if v34 {
        if v34 {
            var v37 uint8 = 1
            _ = v37
        }
    } else {
        v38 := GenS2{F0: -297, F1: "foo bar"}
        fmt.Println(v33, v34, v2)
        _ = v38
    }
    var v39 float32 = -2.25
    switch -1 {
    case 1:
        for fi40 := 0; fi40 < 10; fi40++ {
            v41 := GenS1{F0: (127 != 1), F1: v4}
            _ = fi40
            _ = v41
        }
    default:
        v32 = v32
    }
    v42 := []int{v1, v1}
    if v34 {
        var v43 string = "a"
        if (((-9223372036854775808 < (v1 & v1)) || false) && (!v34)) {
            fmt.Println(v33, v34, v39)
            for fi44 := 0; fi44 < 14; fi44++ {
                switch v1 {
                case 2147483647:
                    fmt.Println(v33, fi44)
                case -2147483648:
                    switch (((v43 + v43) + (v43 + (v43 + v43))) + v43) {
                    case "":
                        var v45 uint32 = 956
                        _ = v45
                    default:
                        var v46 uint8 = 92
                        _ = v46
                    }
                case 9223372036854775807:
                    switch (v43 + "") {
                    case "\n\t":
                        var v47 float32 = 0.1
                        _ = v47
                    default:
                        if v34 {
                            for fi48 := 0; fi48 < 1; fi48++ {
                                v1 -= (v1 % ((fi44) | 1))
                                var v49 uint32 = 4294967295
                                v1 += (v1 * (fi44 | 9223372036854775807))
                                _ = fi48
                                _ = v49
                            }
                            for fi50 := 0; fi50 < 10; fi50++ {
                                if v34 {
                                    v51 := GenS5{F0: true, F1: v43, F2: 32767}
                                    fmt.Println(v43, fi44)
                                    _ = v51
                                } else {
                                    var v52 uint32 = 42
                                    _ = v52
                                }
                                v53 := GenS5{F0: v34, F1: "", F2: 1}
                                var v54 uint32 = 1
                                _ = fi50
                                _ = v53
                                _ = v54
                            }
                            if len(v42) > 0 {
                                v42[0] = 4294967295
                            }
                        }
                    }
                default:
                    _ = len(v42)
                }
                _ = fi44
            }
        } else {
            v34 = v34
        }
        for fi55 := 0; fi55 < 15; fi55++ {
            delete(v3, (fi55 - v1))
            _ = fi55
        }
        _ = v43
    } else {
        if (!v34) {
            var v56 int64 = v4
            v57 := GenS1{F0: ((v39 * 3.14159) <= v39), F1: 290}
            _ = v56
            _ = v57
        } else {
            if (true && v34) {
                switch v1 {
                case 528:
                    fmt.Println(v4)
                case 9223372036854775807:
                    for fi58 := 0; fi58 < 7; fi58++ {
                        fmt.Println(v39)
                        v6[280] = ((v2 + v33) == v2)
                        var v59 float32 = 1e10
                        _ = fi58
                        _ = v59
                    }
                case 2147483647:
                    fmt.Println(v32, v33)
                default:
                    v2 -= v2
                }
            }
            v60 := GenS5{F0: true, F1: "\n\t", F2: 0}
            switch "" {
            case "\n\t":
                var v61 uint32 = uint32(v32)
                _ = v61
            case "hello":
                v1 = v1
            }
            _ = v60
        }
        v32 += (v32 / (((v32 * 1)) | 1))
    }
    v1 = v1
    v32 = 2
    for fi62 := 0; fi62 < 17; fi62++ {
        if true {
            for fi63 := 0; fi63 < 10; fi63++ {
                v64 := GenS3{F0: 2, F1: 32767, F2: 0}
                _ = fi63
                _ = v64
            }
            var v65 int8 = 0
            fmt.Println(v39)
            _ = v65
        } else {
            if false {
                v34 = (v34 && ((v2 << ((18446744073709551615) & 63)) != 2147483648))
            } else {
                v66 := GenS5{F0: (v34 && true), F1: ((("a" + "hello") + "\n\t") + ((((("a" + "a") + "\n\t") + (("hello" + "hello") + "a")) + "a") + "")), F2: 32767}
                if v34 {
                    v4 = v4
                    fmt.Println(v33)
                } else {
                    var v67 uint16 = 255
                    v68 := GenS6{F0: 42, F1: v32, F2: 1000}
                    switch 127 {
                    case 9223372036854775807:
                        switch "foo bar" {
                        case "\n\t":
                            var v69 int = fi62
                            _ = v69
                        case "hello":
                            var v70 int8 = 127
                            _ = v70
                        default:
                            for fi71 := 0; fi71 < 3; fi71++ {
                                v72 := GenS4{F0: 32767}
                                v73 := GenS7{F0: "a"}
                                var v74 uint16 = 938
                                _ = fi71
                                _ = v72
                                _ = v73
                                _ = v74
                            }
                        }
                    case 4294967295:
                        for fi75 := 0; fi75 < 20; fi75++ {
                            v76 := GenS8{F0: fi75, F1: fi75}
                            var v77 bool = (51 == 915)
                            _ = fi75
                            _ = v76
                            _ = v77
                        }
                    default:
                        v3[v1] = "hello"
                    }
                    _ = v67
                    _ = v68
                }
                _ = v66
            }
            fmt.Println(fi62, v4)
            if v34 {
                for fi78 := 0; fi78 < 18; fi78++ {
                    switch fi78 {
                    case 396:
                        if len(v42) > 0 {
                            v42[0] = (fi62 ^ 1)
                        }
                    case 32768:
                        v79 := GenS7{F0: ("" + "\n\t")}
                        _ = v79
                    }
                    _ = fi78
                }
            } else {
                switch fi62 {
                case 1:
                    var v80 int16 = 42
                    _ = v80
                case 32768:
                    v33 += v2
                case -349:
                    switch (v1 << ((255) & 63)) {
                    case -196:
                        switch ("\n\t" + ("a" + "a")) {
                        case "hello":
                            v5[(v2 & (^(v33)))] = (fi62 >> ((2) & 63))
                        case "foo bar":
                            for fi81 := 0; fi81 < 7; fi81++ {
                                switch "\n\t" {
                                case "hello":
                                    v32 += (v32 / (((v32 / ((506) | 1))) | 1))
                                case "a":
                                    if (!((!((-393 <= -1) || (v34 && (!(!v34))))) || v34)) {
                                        v2 = 2147483648
                                        if v34 {
                                            v4 = v4
                                            var v82 uint16 = 42
                                            v83 := GenS9{F0: 0, F1: int16(v32), F2: -643}
                                            v1 = 574
                                            if len(v42) >= 2 {
                                                v42 = v42[1:len(v42)]
                                            }
                                            _ = v82
                                            _ = v83
                                        }
                                        var v84 string = "foo bar"
                                        _ = v84
                                    } else {
                                        if true {
                                            v33 -= 2
                                            if (-44 <= 127) {
                                                fmt.Println(fi81)
                                                switch (("\n\t" + ("a" + "foo bar")) + "") {
                                                case "a":
                                                    var v85 uint64 = (v32 * v32)
                                                    _ = v85
                                                case "\n\t":
                                                    if (!((1.5 + ((1e10 * (((-2.25 * 1.5) * 1e10) / 1.0)) + 3.14159)) < 0.1)) {
                                                        if v34 {
                                                            switch "a" {
                                                            case "hello":
                                                                v39 = ((((3.14159 * 3.14159) + v39) + v39) - v39)
                                                            case "a":
                                                                if ("hello" <= "a") {
                                                                    if len(v42) >= 2 {
                                                                        v42 = v42[1:len(v42)]
                                                                    }
                                                                    v2 += 2
                                                                    var v86 float64 = (0.1 - 0.0)
                                                                    _ = v86
                                                                } else {
                                                                    v87 := GenS5{F0: false, F1: "a", F2: 291}
                                                                    var v88 bool = v34
                                                                    _ = v87
                                                                    _ = v88
                                                                }
                                                            default:
                                                                if false {
                                                                    v34 = (!v34)
                                                                    if ((43 <= 1) || (v39 <= ((v39 / (((v39 * v39) * (v39 * -2.25)) - v39)) * (v39 - 1e10)))) {
                                                                        v42 = append(v42, 624)
                                                                    } else {
                                                                        if ((0.1 == 3.14159) && v34) {
                                                                            var v89 float32 = 1e10
                                                                            if ((v34 || v34) || true) {
                                                                                v90 := GenS7{F0: ("foo bar" + (("foo bar" + "a") + "hello"))}
                                                                                switch (fi81 % ((v1) | 1)) {
                                                                                case 9223372036854775807:
                                                                                    v89 -= (1.5 / 1e10)
                                                                                case 333:
                                                                                    if false {
                                                                                        fmt.Println(v33, v39)
                                                                                    }
                                                                                case -1:
                                                                                    var v91 int32 = -95
                                                                                    _ = v91
                                                                                default:
                                                                                    v92 := v3[(^(v1))]
                                                                                    _ = v92
                                                                                }
                                                                                _ = v90
                                                                            }
                                                                            if ((^(v2)) >= 42) {
                                                                                var v93 uint64 = v32
                                                                                _ = v93
                                                                            }
                                                                            _ = v89
                                                                        }
                                                                        if len(v42) > 0 {
                                                                            v42[0] = fi62
                                                                        }
                                                                    }
                                                                } else {
                                                                    v94 := v3[fi81]
                                                                    v95 := GenS4{F0: -648}
                                                                    var v96 bool = false
                                                                    if (255 <= 0) {
                                                                        v97 := GenS7{F0: ("" + (v94 + v94))}
                                                                        _ = v97
                                                                    }
                                                                    if true {
                                                                        v32 = v32
                                                                    }
                                                                    if v34 {
                                                                        if (v34 && false) {
                                                                            switch v94 {
                                                                            case "\n\t":
                                                                                v33 = v33
                                                                            case "foo bar":
                                                                                var v98 uint16 = 0
                                                                                _ = v98
                                                                            }
                                                                            var v99 int64 = v4
                                                                            if (v34 || ((4294967295 > 18446744073709551615) || true)) {
                                                                                if true {
                                                                                    v39 = 1e10
                                                                                    if v34 {
                                                                                        fmt.Println(v34, v33)
                                                                                    } else {
                                                                                        v32 += (v32 - v32)
                                                                                        var v100 string = v94
                                                                                        switch 127 {
                                                                                        case 0:
                                                                                            var v101 uint32 = v2
                                                                                            _ = v101
                                                                                        }
                                                                                        _ = v100
                                                                                    }
                                                                                }
                                                                                switch 32767 {
                                                                                case -288:
                                                                                    v96 = (253 >= 255)
                                                                                case 9223372036854775807:
                                                                                    v3[-720] = "a"
                                                                                }
                                                                                if ((false && false) && v34) {
                                                                                    var v102 bool = false
                                                                                    if len(v42) > 0 {
                                                                                        v42[0] = -2147483648
                                                                                    }
                                                                                    fmt.Println(v39, v2, fi81)
                                                                                    _ = v102
                                                                                }
                                                                            } else {
                                                                                switch int(v2) {
                                                                                case -261:
                                                                                    var v103 int32 = -1
                                                                                    _ = v103
                                                                                default:
                                                                                    switch v94 {
                                                                                    case "":
                                                                                        v34 = v96
                                                                                        v96 = v34
                                                                                        var v104 int = v1
                                                                                        _ = v104
                                                                                    default:
                                                                                        var v105 int = -9223372036854775808
                                                                                        _ = v105
                                                                                    }
                                                                                }
                                                                                v1 = -238
                                                                                switch v94 {
                                                                                case "a":
                                                                                    v2 = 2147483647
                                                                                default:
                                                                                    var v106 uint = 4294967295
                                                                                    _ = v106
                                                                                }
                                                                            }
                                                                            v107 := GenS10{F0: -1}
                                                                            v108 := GenS7{F0: (v94 + (v94 + (v94 + v94)))}
                                                                            _ = v107
                                                                            _ = v108
                                                                            _ = v99
                                                                        } else {
                                                                            v109 := GenS2{F0: 1, F1: ((("\n\t" + v94) + ((((v94 + (v94 + (v94 + (v94 + "hello")))) + ("\n\t" + ("foo bar" + ("" + "a")))) + (v94 + v94)) + "")) + v94)}
                                                                            _ = v109
                                                                        }
                                                                    }
                                                                    switch "" {
                                                                    case "foo bar":
                                                                        if (false && (1 == 1)) {
                                                                            var v110 uint = 65535
                                                                            _ = v110
                                                                        }
                                                                    case "":
                                                                        v34 = v96
                                                                    default:
                                                                        _ = len(v5)
                                                                    }
                                                                    v32 = v32
                                                                    if v34 {
                                                                        switch v94 {
                                                                        case "a":
                                                                            _ = len(v42)
                                                                        case "hello":
                                                                            _ = len(v6)
                                                                        }
                                                                        switch (^(fi81)) {
                                                                        case -129:
                                                                            v111 := v5[v33]
                                                                            _ = v111
                                                                        case 0:
                                                                            v4 -= v4
                                                                        }
                                                                        var v112 int64 = (v4 << ((0) & 63))
                                                                        _ = v112
                                                                    } else {
                                                                        var v113 uint32 = (v33 << ((166) & 63))
                                                                        v5[v2] = (-(fi81))
                                                                        v114 := v3[v1]
                                                                        var v115 int8 = 120
                                                                        _ = v113
                                                                        _ = v114
                                                                        _ = v115
                                                                    }
                                                                    if (true && (897 != 42)) {
                                                                        var v116 int8 = 15
                                                                        _ = v116
                                                                    } else {
                                                                        var v117 bool = (v2 >= v2)
                                                                        if (-332 == 1) {
                                                                            switch fi62 {
                                                                            case 42:
                                                                                var v118 string = v94
                                                                                _ = v118
                                                                            default:
                                                                                if true {
                                                                                    if v96 {
                                                                                        fmt.Println(fi81, v1)
                                                                                        v2 += 2147483647
                                                                                    } else {
                                                                                        var v119 int64 = 42
                                                                                        var v120 bool = (0.1 >= (3.14159 * 1.5))
                                                                                        _ = v119
                                                                                        _ = v120
                                                                                    }
                                                                                    if ((v1 | fi62) <= -326) {
                                                                                        v121 := GenS7{F0: "a"}
                                                                                        _ = v121
                                                                                    } else {
                                                                                        v1 += fi62
                                                                                    }
                                                                                    var v122 uint8 = 42
                                                                                    _ = v122
                                                                                } else {
                                                                                    v34 = v34
                                                                                    var v123 int8 = 1
                                                                                    _ = v123
                                                                                }
                                                                            }
                                                                            v124 := GenS4{F0: -32768}
                                                                            if ((v2 >> ((0) & 63)) <= v33) {
                                                                                v2 = v2
                                                                                v125 := GenS11{F0: 26}
                                                                                v32 = 42
                                                                                _ = v125
                                                                            }
                                                                            v117 = v117
                                                                            if (42 > (v32 / (((v32 * v32)) | 1))) {
                                                                                switch v94 {
                                                                                case "hello":
                                                                                    v1 = (v1 * 9223372036854775807)
                                                                                case "":
                                                                                    fmt.Println(v1)
                                                                                default:
                                                                                    v1 += v1
                                                                                }
                                                                                fmt.Println(fi62)
                                                                            } else {
                                                                                v39 = (v39 + v39)
                                                                            }
                                                                            var v126 uint8 = 127
                                                                            fmt.Println(v2)
                                                                            v33 = 2147483648
                                                                            v126 = 128
                                                                            _ = len(v42)
                                                                            v127 := GenS12{F0: 1}
                                                                            _ = len(v3)
                                                                            _ = v124
                                                                            _ = v126
                                                                            _ = v127
                                                                        }
                                                                        _ = v117
                                                                    }
                                                                    _ = v94
                                                                    _ = v95
                                                                    _ = v96
                                                                }
                                                            }
                                                            switch "a" {
                                                            case "a":
                                                                v4 = v4
                                                            case "hello":
                                                                v32 = 0
                                                            default:
                                                                if v34 {
                                                                    v128 := GenS13{F0: (fi81 % ((0) | 1)), F1: 1}
                                                                    fmt.Println(v39, v34, fi62)
                                                                    v2 = (v2 ^ (v2 & v2))
                                                                    _ = v128
                                                                }
                                                                v6[(fi81 ^ (^(fi81)))] = true
                                                                switch "a" {
                                                                case "\n\t":
                                                                    switch -414 {
                                                                    case 683:
                                                                        if len(v42) >= 2 {
                                                                            v42 = v42[1:len(v42)]
                                                                        }
                                                                    case 32768:
                                                                        switch "foo bar" {
                                                                        case "\n\t":
                                                                            if (v39 > v39) {
                                                                                v129 := v3[0]
                                                                                _ = v129
                                                                            } else {
                                                                                v1 += fi81
                                                                                fmt.Println(v4, v32, v4)
                                                                            }
                                                                        case "":
                                                                            var v130 int = (fi81 - -289)
                                                                            _ = v130
                                                                        }
                                                                    }
                                                                default:
                                                                    var v131 float32 = ((3.14159 * v39) * (v39 + -2.25))
                                                                    _ = v131
                                                                }
                                                                v1 -= 0
                                                                v132 := GenS2{F0: 1000, F1: "\n\t"}
                                                                if (!((0 > 255) || true)) {
                                                                    switch 0 {
                                                                    case 128:
                                                                        v39 = 0.1
                                                                    case -2147483648:
                                                                        switch "a" {
                                                                        case "\n\t":
                                                                            switch "foo bar" {
                                                                            case "hello":
                                                                                v133 := GenS5{F0: v34, F1: "foo bar", F2: 0}
                                                                                _ = v133
                                                                            case "":
                                                                                if (fi81 != (^(v1))) {
                                                                                    switch "" {
                                                                                    case "\n\t":
                                                                                        var v134 bool = false
                                                                                        _ = v134
                                                                                    default:
                                                                                        if false {
                                                                                            switch fi62 {
                                                                                            case 4294967295:
                                                                                                v135 := v6[(fi62 / ((128) | 1))]
                                                                                                _ = v135
                                                                                            case -2147483648:
                                                                                                _ = len(v5)
                                                                                                _ = len(v42)
                                                                                                _ = len(v5)
                                                                                            case 1:
                                                                                                v33 = v33
                                                                                            }
                                                                                            switch ("\n\t" + "foo bar") {
                                                                                            case "\n\t":
                                                                                                switch (fi62 * (-(fi81))) {
                                                                                                case 371:
                                                                                                    switch "hello" {
                                                                                                    case "":
                                                                                                        switch 384 {
                                                                                                        case 0:
                                                                                                            var v136 bool = true
                                                                                                            _ = v136
                                                                                                        }
                                                                                                    case "a":
                                                                                                        var v137 bool = false
                                                                                                        _ = v137
                                                                                                    case "hello":
                                                                                                        if (871 <= 32767) {
                                                                                                            v39 -= 3.14159
                                                                                                            if len(v42) > 0 {
                                                                                                                v42[0] = (fi62 - (v1 >> ((214) & 63)))
                                                                                                            }
                                                                                                            switch 518 {
                                                                                                            case 32767:
                                                                                                                v2 = v2
                                                                                                            case -635:
                                                                                                                if len(v42) > 0 {
                                                                                                                    v42[0] = fi81
                                                                                                                }
                                                                                                                _ = len(v42)
                                                                                                            case 0:
                                                                                                                delete(v5, 1)
                                                                                                            }
                                                                                                        } else {
                                                                                                            v138 := GenS5{F0: (3.14159 <= 3.14159), F1: "\n\t", F2: 0}
                                                                                                            _ = v138
                                                                                                        }
                                                                                                        v39 -= ((v39 / (v39 + (((0.1 + v39) + (v39 / v39)) + ((v39 - ((v39 + v39) - 0.0)) + float32((((3.14159 - (-2.25 / 3.14159)) * (((((float64(0.1) - 3.14159) - (3.14159 - 1.5)) + 0.1) / 1.5) * 1.5)) * 0.1)))))) + (v39 + (1.5 - v39)))
                                                                                                    }
                                                                                                case -268:
                                                                                                    v139 := GenS11{F0: 0}
                                                                                                    _ = v139
                                                                                                }
                                                                                            case "":
                                                                                                if (v34 || ((v4 / ((1) | 1)) == (v4 & (v4 * v4)))) {
                                                                                                    v3[(-(fi62))] = "hello"
                                                                                                    if false {
                                                                                                        v1 = v1
                                                                                                        v2 = 2
                                                                                                        var v140 int64 = v4
                                                                                                        _ = v140
                                                                                                    }
                                                                                                    var v141 string = ""
                                                                                                    _ = v141
                                                                                                } else {
                                                                                                    var v142 string = "hello"
                                                                                                    _ = v142
                                                                                                }
                                                                                            case "foo bar":
                                                                                                var v143 float64 = ((((0.1 + 1e10) / float64((1.5 * ((v39 - v39) * (v39 / v39))))) * (1e10 + 0.0)) * 1e10)
                                                                                                _ = v143
                                                                                            }
                                                                                            v144 := GenS14{F0: (v34 || v34)}
                                                                                            if v34 {
                                                                                                v4 = v4
                                                                                            }
                                                                                            if (613 == 1) {
                                                                                                if v34 {
                                                                                                    switch "a" {
                                                                                                    case "a":
                                                                                                        if len(v42) > 0 {
                                                                                                            v42[0] = (fi81 - -1)
                                                                                                        }
                                                                                                    }
                                                                                                }
                                                                                                if len(v42) >= 2 {
                                                                                                    v42 = v42[1:len(v42)]
                                                                                                }
                                                                                            } else {
                                                                                                v33 = (v33 ^ v33)
                                                                                                switch (fi81 & (v1 + (fi81 ^ (v1 | (^(fi81)))))) {
                                                                                                case -627:
                                                                                                    if true {
                                                                                                        var v145 int32 = 579
                                                                                                        var v146 float32 = (1e10 / (1e10 + (v39 * (v39 + float32(float64(1.5))))))
                                                                                                        _ = v145
                                                                                                        _ = v146
                                                                                                    }
                                                                                                default:
                                                                                                    var v147 int16 = -32768
                                                                                                    _ = v147
                                                                                                }
                                                                                            }
                                                                                            _ = v144
                                                                                        }
                                                                                    }
                                                                                    if v34 {
                                                                                        fmt.Println(v34, v32)
                                                                                        _ = len(v42)
                                                                                    }
                                                                                    v33 = v33
                                                                                } else {
                                                                                    v148 := GenS3{F0: 2, F1: 127, F2: 698}
                                                                                    _ = v148
                                                                                }
                                                                            }
                                                                        case "hello":
                                                                            var v149 float64 = 1e10
                                                                            _ = v149
                                                                        }
                                                                    default:
                                                                        var v150 uint8 = 128
                                                                        _ = v150
                                                                    }
                                                                }
                                                                _ = v132
                                                            }
                                                        }
                                                        var v151 uint64 = (v32 << ((18446744073709551615) & 63))
                                                        _ = v151
                                                    }
                                                default:
                                                    v152 := GenS15{F0: 100, F1: v32}
                                                    _ = v152
                                                }
                                                switch (v1 / ((fi62) | 1)) {
                                                case 9223372036854775807:
                                                    v1 = v1
                                                case 32768:
                                                    if v34 {
                                                        if v34 {
                                                            if ((v34 || (!v34)) || (!true)) {
                                                                v153 := GenS16{F0: -1}
                                                                v2 -= v2
                                                                _ = v153
                                                            }
                                                            var v154 string = "\n\t"
                                                            _ = v154
                                                        } else {
                                                            fmt.Println(fi62, fi81, v2)
                                                            if (true || (!v34)) {
                                                                var v155 bool = v34
                                                                _ = v155
                                                            } else {
                                                                v32 = v32
                                                                v156 := GenS17{F0: 255, F1: ((128 != 209) && v34), F2: 0}
                                                                _ = v156
                                                            }
                                                        }
                                                        var v157 float64 = 3.14159
                                                        var v158 uint64 = 18446744073709551615
                                                        v159 := GenS18{F0: 127, F1: 604, F2: (!(v34 || (-1 != 127)))}
                                                        fmt.Println(v2, v34, fi81)
                                                        switch ((("a" + (("hello" + "a") + ((("foo bar" + (("" + (("hello" + "") + "hello")) + "hello")) + "") + "a"))) + "a") + "foo bar") {
                                                        case "":
                                                            switch "a" {
                                                            case "a":
                                                                var v160 int16 = -32768
                                                                v161 := v6[(v1 / (((v1 + 127)) | 1))]
                                                                _ = v160
                                                                _ = v161
                                                            case "":
                                                                var v162 int = (v1 ^ fi62)
                                                                _ = v162
                                                            case "foo bar":
                                                                var v163 uint32 = (v2 / ((2147483647) | 1))
                                                                _ = v163
                                                            }
                                                        default:
                                                            switch (fi62 + int(v2)) {
                                                            case 32767:
                                                                if false {
                                                                    v164 := GenS9{F0: 113, F1: -32768, F2: 32767}
                                                                    _ = v164
                                                                }
                                                            }
                                                        }
                                                        _ = v157
                                                        _ = v158
                                                        _ = v159
                                                    } else {
                                                        v4 += int64(v2)
                                                        var v165 string = ""
                                                        _ = v165
                                                    }
                                                case -1:
                                                    v39 += v39
                                                default:
                                                    if v34 {
                                                        v34 = false
                                                        var v166 int = fi81
                                                        v1 -= v1
                                                        var v167 uint = 656
                                                        _ = v166
                                                        _ = v167
                                                    }
                                                    v2 = (^(v33))
                                                    v34 = (!((((-2.25 / ((3.14159 * 3.14159) - 1e10)) * 1.5) <= 0.0) || (312 == 2147483648)))
                                                }
                                                v168 := GenS19{F0: 32767, F1: v34}
                                                var v169 float32 = v39
                                                v4 -= (v4 ^ v4)
                                                v170 := v5[0]
                                                fmt.Println(v32, v32, v1)
                                                v1 += fi81
                                                _ = v168
                                                _ = v169
                                                _ = v170
                                            } else {
                                                if (0 == 4294967295) {
                                                    fmt.Println(v4, v32)
                                                    fmt.Println(v33, v32)
                                                } else {
                                                    var v171 uint = 303
                                                    _ = v171
                                                }
                                                fmt.Println(v1, fi81)
                                                fmt.Println(v33)
                                            }
                                        } else {
                                            fmt.Println(v32, v33, v33)
                                            v5[(v2 + v2)] = -724
                                            v172 := GenS20{F0: -432, F1: 14, F2: v34}
                                            v34 = v34
                                            _ = v172
                                        }
                                        var v173 uint64 = 18446744073709551615
                                        var v174 uint64 = (v173 - 834)
                                        var v175 uint16 = 42
                                        if (int64(v175) != v4) {
                                            v174 += (^(v173))
                                            v176 := GenS21{F0: 198, F1: v173, F2: 0}
                                            switch "hello" {
                                            case "a":
                                                var v177 float64 = ((-2.25 - 0.1) * ((3.14159 - (0.0 / -2.25)) - (1.5 * 1e10)))
                                                _ = v177
                                            case "foo bar":
                                                var v178 uint16 = v175
                                                _ = v178
                                            case "\n\t":
                                                v179 := GenS19{F0: -701, F1: v34}
                                                _ = v179
                                            }
                                            _ = v176
                                        } else {
                                            switch "foo bar" {
                                            case "foo bar":
                                                switch "\n\t" {
                                                case "hello":
                                                    if (v175 < (v175 ^ (^(v175)))) {
                                                        v32 = (v173 & (v173 % ((530) | 1)))
                                                        v180 := GenS22{F0: true, F1: v175, F2: 1}
                                                        v181 := GenS9{F0: -69, F1: -32768, F2: 1}
                                                        _ = v180
                                                        _ = v181
                                                    }
                                                case "foo bar":
                                                    if (false && v34) {
                                                        var v182 float32 = ((v39 * 1e10) * ((-2.25 - 0.1) / (float32(0.1) - (v39 + 0.0))))
                                                        _ = v182
                                                    }
                                                case "a":
                                                    if (false && v34) {
                                                        if len(v42) >= 2 {
                                                            v42 = v42[1:len(v42)]
                                                        }
                                                        switch v1 {
                                                        case 519:
                                                            if v34 {
                                                                v1 += v1
                                                            }
                                                        default:
                                                            v183 := GenS23{F0: -1, F1: 4294967295}
                                                            _ = v183
                                                        }
                                                        fmt.Println(v173)
                                                    } else {
                                                        var v184 bool = v34
                                                        _ = v184
                                                    }
                                                default:
                                                    if ((v4 * 9223372036854775807) == v4) {
                                                        v39 = float32(-2.25)
                                                        v33 = (v2 - 1)
                                                        v2 = (v33 << ((574) & 63))
                                                    } else {
                                                        var v185 int32 = -2147483648
                                                        _ = len(v42)
                                                        _ = v185
                                                    }
                                                }
                                            }
                                            v186 := GenS24{F0: 255}
                                            switch (v1 * (fi81 + (v1 ^ -849))) {
                                            case 588:
                                                var v187 uint = 255
                                                _ = v187
                                            }
                                            _ = v186
                                        }
                                        var v188 string = "hello"
                                        _ = v173
                                        _ = v174
                                        _ = v175
                                        _ = v188
                                    }
                                case "\n\t":
                                    switch (("hello" + "hello") + "hello") {
                                    case "\n\t":
                                        var v189 float64 = 0.1
                                        _ = v189
                                    default:
                                        var v190 uint16 = 42
                                        _ = v190
                                    }
                                }
                                var v191 int = v1
                                if true {
                                    v42 = append(v42, (fi81 - fi62))
                                }
                                _ = fi81
                                _ = v191
                            }
                        }
                    case -380:
                        v39 -= v39
                    default:
                        v192 := GenS1{F0: (!true), F1: -9223372036854775808}
                        _ = v192
                    }
                }
            }
        }
        var v193 uint64 = (v32 / ((2) | 1))
        v39 = v39
        _ = fi62
        _ = v193
    }
    fmt.Println(v2, v33, v1)
    fmt.Println(v4, v33, v32)
    for fi194 := 0; fi194 < 8; fi194++ {
        fmt.Println(v4, v4)
        _ = fi194
    }
    for fi195 := 0; fi195 < 18; fi195++ {
        var v196 string = "hello"
        v197 := GenS25{F0: v196}
        _ = fi195
        _ = v196
        _ = v197
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v42:", len(v42), v42)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v39
    _ = v4
    _ = v42
    _ = v5
    _ = v6
}
