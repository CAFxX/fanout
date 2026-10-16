package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 bool; F1 int32; F2 int16 }
type GenS4 struct { F0 bool; F1 bool }
type GenS5 struct { F0 uint64; F1 uint }
type GenS6 struct { F0 string }
type GenS7 struct { F0 uint; F1 uint; F2 int32 }
type GenS8 struct { F0 int32; F1 uint32 }
type GenS9 struct { F0 int; F1 int }
type GenS10 struct { F0 uint; F1 bool; F2 int32 }
type GenS11 struct { F0 uint8 }
type GenS12 struct { F0 uint; F1 int32; F2 uint64 }
type GenS13 struct { F0 int16; F1 int }
type GenS14 struct { F0 uint8 }
type GenS15 struct { F0 uint16; F1 string }
type GenS16 struct { F0 uint64; F1 uint }
type GenS17 struct { F0 string; F1 int16 }
type GenS18 struct { F0 bool; F1 uint16 }
type GenS19 struct { F0 int8 }
type GenS20 struct { F0 uint8; F1 int32 }

func main() {
    var v1 uint8 = 42
    v2 := []uint32{65535}
    v3 := []int{}
    switch "foo bar" {
    case "hello":
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    case "foo bar":
        fmt.Println(v1)
    }
    if (-89 < -36) {
        if true {
            for fi4 := 0; fi4 < 10; fi4++ {
                var v5 int64 = -9223372036854775808
                if true {
                    for fi6 := 0; fi6 < 7; fi6++ {
                        var v7 int8 = -1
                        fmt.Println(fi4)
                        _ = fi6
                        _ = v7
                    }
                    var v8 int = (fi4 & fi4)
                    switch "hello" {
                    case "hello":
                        for fi9 := 0; fi9 < 16; fi9++ {
                            var v10 int = fi4
                            var v11 int16 = 15
                            v12 := GenS1{F0: 2, F1: "foo bar"}
                            _ = fi9
                            _ = v10
                            _ = v11
                            _ = v12
                        }
                    case "a":
                        v8 = v8
                    default:
                        for fi13 := 0; fi13 < 5; fi13++ {
                            if false {
                                v14 := GenS1{F0: 2, F1: "a"}
                                _ = v14
                            }
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            var v15 uint16 = 1
                            _ = fi13
                            _ = v15
                        }
                    }
                    _ = v8
                }
                _ = fi4
                _ = v5
            }
            if (!(!((2147483647 < 4294967295) || (2147483647 != 42)))) {
                for fi16 := 0; fi16 < 6; fi16++ {
                    for fi17 := 0; fi17 < 2; fi17++ {
                        fmt.Println(fi16, fi16, fi16)
                        v1 = (v1 >> ((493) & 63))
                        _ = fi17
                    }
                    for fi18 := 0; fi18 < 3; fi18++ {
                        v1 = (v1 | (v1 + (v1 - (v1 | 128))))
                        _ = fi18
                    }
                    for fi19 := 0; fi19 < 6; fi19++ {
                        if len(v3) > 0 {
                            v3[0] = 0
                        }
                        _ = fi19
                    }
                    _ = fi16
                }
                if len(v3) > 0 {
                    v3[0] = -2147483648
                }
                _ = len(v3)
            }
        } else {
            v1 += 42
            v20 := GenS2{F0: 1}
            v1 += v1
            _ = v20
        }
    }
    if true {
        v1 += 208
        var v21 int16 = -32768
        v22 := GenS1{F0: 18446744073709551615, F1: ""}
        _ = v21
        _ = v22
    } else {
        _ = len(v3)
        for fi23 := 0; fi23 < 11; fi23++ {
            v1 -= v1
            for fi24 := 0; fi24 < 2; fi24++ {
                if (true || (!true)) {
                    if len(v3) > 0 {
                        v3[0] = 950
                    }
                } else {
                    v25 := GenS3{F0: true, F1: -2147483648, F2: -32768}
                    _ = v25
                }
                _ = len(v3)
                _ = fi24
            }
            _ = fi23
        }
        if true {
            for fi26 := 0; fi26 < 13; fi26++ {
                switch -2147483648 {
                case 0:
                    v3 = append(v3, fi26)
                default:
                    v27 := GenS4{F0: (2 > 42), F1: false}
                    _ = v27
                }
                var v28 uint64 = 464
                _ = fi26
                _ = v28
            }
            if true {
                v1 = (v1 & 1)
                v29 := GenS4{F0: false, F1: true}
                _ = v29
            } else {
                if false {
                    if (32767 > 781) {
                        var v30 float32 = 1e10
                        for fi31 := 0; fi31 < 17; fi31++ {
                            switch "\n\t" {
                            case "":
                                var v32 uint32 = 42
                                _ = v32
                            case "hello":
                                var v33 uint8 = 127
                                _ = v33
                            case "\n\t":
                                switch -908 {
                                case -1:
                                    var v34 uint8 = (v1 | (v1 / (((v1 & 210)) | 1)))
                                    _ = v34
                                case 9223372036854775807:
                                    switch ("foo bar" + (("\n\t" + "a") + "a")) {
                                    case "":
                                        for fi35 := 0; fi35 < 2; fi35++ {
                                            fmt.Println(fi31)
                                            var v36 int32 = -689
                                            fmt.Println(fi35, v1, fi35)
                                            _ = fi35
                                            _ = v36
                                        }
                                    case "\n\t":
                                        if (42 > -1) {
                                            v37 := GenS5{F0: 2, F1: 1}
                                            var v38 bool = (false || (46 > 4294967295))
                                            _ = v37
                                            _ = v38
                                        } else {
                                            v30 += v30
                                            switch fi31 {
                                            case -774:
                                                var v39 int8 = 0
                                                _ = v39
                                            case 128:
                                                var v40 int64 = -937
                                                _ = v40
                                            case 0:
                                                v1 += v1
                                            }
                                            if (-565 != 42) {
                                                var v41 float64 = float64(((v30 - v30) - (3.14159 / float32(0.1))))
                                                switch "\n\t" {
                                                case "hello":
                                                    switch (fi31 << ((255) & 63)) {
                                                    case -9223372036854775808:
                                                        fmt.Println(v41, v30)
                                                    case 128:
                                                        fmt.Println(v41)
                                                    case 42:
                                                        if true {
                                                            if false {
                                                                if false {
                                                                    if false {
                                                                        for fi42 := 0; fi42 < 15; fi42++ {
                                                                            var v43 int8 = 42
                                                                            if (584 == 18446744073709551615) {
                                                                                v44 := GenS4{F0: true, F1: false}
                                                                                switch ("hello" + "a") {
                                                                                case "foo bar":
                                                                                    var v45 int = fi42
                                                                                    _ = v45
                                                                                }
                                                                                fmt.Println(fi42)
                                                                                v46 := GenS5{F0: 1, F1: 65535}
                                                                                v47 := GenS6{F0: "foo bar"}
                                                                                if true {
                                                                                    var v48 uint = 2
                                                                                    if (0 <= 0) {
                                                                                        var v49 uint16 = 1
                                                                                        _ = v49
                                                                                    } else {
                                                                                        var v50 uint16 = 255
                                                                                        fmt.Println(v41, v43, v50)
                                                                                        _ = v50
                                                                                    }
                                                                                    _ = v48
                                                                                }
                                                                                var v51 float64 = 3.14159
                                                                                _ = v44
                                                                                _ = v46
                                                                                _ = v47
                                                                                _ = v51
                                                                            } else {
                                                                                switch (fi31 / ((4294967295) | 1)) {
                                                                                case 127:
                                                                                    fmt.Println(v30, fi31)
                                                                                case 128:
                                                                                    if ((0.1 / v41) > v41) {
                                                                                        v1 += 90
                                                                                    } else {
                                                                                        v2 = append(v2, 2)
                                                                                    }
                                                                                case 2147483647:
                                                                                    v52 := GenS4{F0: false, F1: true}
                                                                                    _ = v52
                                                                                }
                                                                                if len(v3) > 0 {
                                                                                    v3[0] = (fi42 * -2147483648)
                                                                                }
                                                                            }
                                                                            v41 = 0.1
                                                                            var v53 string = "foo bar"
                                                                            if (0 >= -1) {
                                                                                switch (fi31 << ((18446744073709551615) & 63)) {
                                                                                case 32768:
                                                                                    v43 += (v43 % ((v43) | 1))
                                                                                case -1:
                                                                                    _ = len(v2)
                                                                                case 0:
                                                                                    v41 = float64(v30)
                                                                                }
                                                                                v30 -= ((float32(float64(float32((float64(v30) - 3.14159)))) + 3.14159) * v30)
                                                                                var v54 int32 = -116
                                                                                var v55 int8 = (v43 + 73)
                                                                                _ = v54
                                                                                _ = v55
                                                                            }
                                                                            if ((v30 * float32(float64(1e10))) >= 0.0) {
                                                                                v56 := GenS2{F0: 484}
                                                                                var v57 bool = false
                                                                                fmt.Println(v41, v1, v1)
                                                                                _ = v56
                                                                                _ = v57
                                                                            } else {
                                                                                var v58 int8 = (v43 / ((v43) | 1))
                                                                                v2 = append(v2, 65535)
                                                                                switch fi31 {
                                                                                case -2147483648:
                                                                                    if (true && false) {
                                                                                        _ = len(v2)
                                                                                        v59 := GenS7{F0: 818, F1: 18446744073709551615, F2: 0}
                                                                                        _ = v59
                                                                                    } else {
                                                                                        var v60 uint = 18446744073709551615
                                                                                        v1 += v1
                                                                                        var v61 int8 = (v43 & v43)
                                                                                        _ = v60
                                                                                        _ = v61
                                                                                    }
                                                                                case 57:
                                                                                    if (!(false && true)) {
                                                                                        var v62 int = fi42
                                                                                        var v63 float64 = v41
                                                                                        var v64 bool = false
                                                                                        _ = v62
                                                                                        _ = v63
                                                                                        _ = v64
                                                                                    } else {
                                                                                        var v65 uint8 = (^(v1))
                                                                                        _ = v65
                                                                                    }
                                                                                case 826:
                                                                                    v53 = v53
                                                                                }
                                                                                _ = v58
                                                                            }
                                                                            _ = len(v3)
                                                                            _ = fi42
                                                                            _ = v43
                                                                            _ = v53
                                                                        }
                                                                    }
                                                                } else {
                                                                    if len(v2) >= 2 {
                                                                        v2 = v2[1:len(v2)]
                                                                    }
                                                                    if (719 < 42) {
                                                                        if true {
                                                                            var v66 int = 135
                                                                            var v67 int8 = 100
                                                                            _ = v66
                                                                            _ = v67
                                                                        }
                                                                        fmt.Println(v30, fi31)
                                                                    } else {
                                                                        _ = len(v2)
                                                                        if (true && true) {
                                                                            v68 := GenS1{F0: 4294967295, F1: ""}
                                                                            var v69 string = "\n\t"
                                                                            _ = v68
                                                                            _ = v69
                                                                        }
                                                                    }
                                                                    if true {
                                                                        var v70 int16 = -336
                                                                        if ((v1 ^ (^(v1))) < (v1 ^ v1)) {
                                                                            v71 := GenS4{F0: false, F1: (true && true)}
                                                                            _ = v71
                                                                        }
                                                                        _ = v70
                                                                    }
                                                                }
                                                                if false {
                                                                    var v72 int16 = 1
                                                                    var v73 string = "hello"
                                                                    _ = v72
                                                                    _ = v73
                                                                } else {
                                                                    fmt.Println(fi31)
                                                                    var v74 uint = 344
                                                                    _ = len(v2)
                                                                    _ = v74
                                                                }
                                                                v75 := GenS8{F0: 84, F1: 2147483647}
                                                                _ = v75
                                                            } else {
                                                                v76 := GenS2{F0: 32767}
                                                                _ = v76
                                                            }
                                                            var v77 int8 = 127
                                                            if false {
                                                                v1 -= (v1 / ((v1) | 1))
                                                                v78 := GenS9{F0: fi31, F1: (fi31 & fi31)}
                                                                v79 := GenS10{F0: 787, F1: (-485 > 0), F2: 487}
                                                                _ = v78
                                                                _ = v79
                                                            }
                                                            _ = v77
                                                        }
                                                    default:
                                                        for fi80 := 0; fi80 < 17; fi80++ {
                                                            v81 := GenS11{F0: (^(v1))}
                                                            v30 -= v30
                                                            v82 := GenS3{F0: false, F1: 42, F2: -114}
                                                            _ = fi80
                                                            _ = v81
                                                            _ = v82
                                                        }
                                                    }
                                                case "foo bar":
                                                    fmt.Println(v1)
                                                case "":
                                                    switch ("foo bar" + "foo bar") {
                                                    case "hello":
                                                        var v83 uint16 = 355
                                                        _ = v83
                                                    case "foo bar":
                                                        v3 = append(v3, (fi31 - (fi31 & fi31)))
                                                    }
                                                }
                                                v84 := GenS11{F0: 0}
                                                _ = v41
                                                _ = v84
                                            }
                                        }
                                    case "hello":
                                        for fi85 := 0; fi85 < 17; fi85++ {
                                            v86 := GenS5{F0: 876, F1: 65535}
                                            fmt.Println(fi85)
                                            fmt.Println(v30)
                                            _ = fi85
                                            _ = v86
                                        }
                                    default:
                                        if (float32(((-2.25 - (0.1 * 1e10)) + ((-2.25 - float64((-2.25 * float32(float64(v30))))) / 1.5))) >= ((v30 / (((v30 / (v30 * (v30 - v30))) + 0.0) - float32(0.0))) * float32(1e10))) {
                                            for fi87 := 0; fi87 < 5; fi87++ {
                                                v88 := GenS7{F0: 42, F1: 0, F2: -1}
                                                _ = fi87
                                                _ = v88
                                            }
                                        }
                                    }
                                }
                            }
                            v1 = v1
                            v1 = v1
                            _ = fi31
                        }
                        for fi89 := 0; fi89 < 9; fi89++ {
                            v30 += 0.0
                            _ = fi89
                        }
                        _ = v30
                    } else {
                        if (0 > 42) {
                            switch -411 {
                            case 32767:
                                if ((!(false || (42 >= 0))) || false) {
                                    v90 := GenS12{F0: 361, F1: 208, F2: 42}
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    _ = v90
                                }
                            case 128:
                                var v91 int64 = 0
                                _ = v91
                            }
                        } else {
                            if ((1.5 / float64((((0.1 / 1.0) - ((float32(float64(float32((3.14159 / (((0.1 * 0.1) + 1e10) - ((float64((-2.25 * 3.14159)) - float64(0.1)) - (3.14159 + ((3.14159 - -2.25) / 3.14159)))))))) + 3.14159) - 1e10)) + (float32(-2.25) / (1e10 + 0.0))))) >= -2.25) {
                                for fi92 := 0; fi92 < 17; fi92++ {
                                    fmt.Println(fi92)
                                    if false {
                                        fmt.Println(fi92, v1)
                                        v1 = 2
                                        v93 := GenS13{F0: -400, F1: fi92}
                                        _ = v93
                                    }
                                    fmt.Println(v1, fi92)
                                    _ = fi92
                                }
                            }
                            var v94 int = -9223372036854775808
                            _ = v94
                        }
                        v1 = (v1 | (v1 / (((^(v1))) | 1)))
                    }
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    for fi95 := 0; fi95 < 10; fi95++ {
                        if (127 < 1) {
                            switch fi95 {
                            case 767:
                                v1 += v1
                            }
                            var v96 int8 = 1
                            v97 := GenS2{F0: 32767}
                            _ = v96
                            _ = v97
                        } else {
                            switch 32767 {
                            case 32768:
                                v98 := GenS14{F0: 1}
                                _ = v98
                            }
                        }
                        v99 := GenS8{F0: -2147483648, F1: 65535}
                        v1 -= 48
                        _ = fi95
                        _ = v99
                    }
                }
                fmt.Println(v1)
            }
            for fi100 := 0; fi100 < 15; fi100++ {
                var v101 int = -145
                v2 = append(v2, 820)
                v101 = v101
                _ = fi100
                _ = v101
            }
        } else {
            fmt.Println(v1)
        }
    }
    for fi102 := 0; fi102 < 13; fi102++ {
        if len(v3) > 0 {
            v3[0] = (fi102 | fi102)
        }
        if (32767 == (-(fi102))) {
            v103 := GenS6{F0: ("hello" + "hello")}
            if false {
                for fi104 := 0; fi104 < 7; fi104++ {
                    var v105 string = "hello"
                    v105 = ((v105 + "a") + "hello")
                    _ = fi104
                    _ = v105
                }
            } else {
                switch ("" + "\n\t") {
                case "a":
                    if (0 >= 42) {
                        v106 := GenS7{F0: 4294967295, F1: 1, F2: -800}
                        var v107 string = "hello"
                        for fi108 := 0; fi108 < 10; fi108++ {
                            var v109 float64 = (1.5 - (-2.25 * 0.0))
                            if false {
                                v2 = append(v2, 4294967295)
                                if (1e10 >= 1e10) {
                                    var v110 uint8 = v1
                                    if (2 != 2) {
                                        var v111 int32 = 0
                                        if len(v3) > 0 {
                                            v3[0] = (fi108 & 42)
                                        }
                                        _ = v111
                                    }
                                    _ = v110
                                } else {
                                    switch fi102 {
                                    case -9223372036854775808:
                                        v112 := GenS8{F0: -83, F1: 4294967295}
                                        _ = v112
                                    case -1:
                                        if (496 != 42) {
                                            var v113 string = v107
                                            switch ((v113 + v113) + v107) {
                                            case "\n\t":
                                                var v114 int = (fi108 | int(v1))
                                                var v115 uint8 = v1
                                                _ = v114
                                                _ = v115
                                            case "":
                                                v113 = (v107 + "a")
                                            default:
                                                switch "hello" {
                                                case "a":
                                                    switch v113 {
                                                    case "":
                                                        var v116 float32 = (-2.25 - (1.5 - (0.0 / 0.1)))
                                                        _ = v116
                                                    case "\n\t":
                                                        switch "hello" {
                                                        case "\n\t":
                                                            _ = len(v3)
                                                        case "foo bar":
                                                            var v117 uint16 = 2
                                                            _ = v117
                                                        case "":
                                                            switch fi108 {
                                                            case 128:
                                                                fmt.Println(v113, fi108)
                                                            case 695:
                                                                v107 = (v113 + v113)
                                                            case 2147483647:
                                                                _ = len(v2)
                                                            default:
                                                                var v118 uint16 = 781
                                                                _ = v118
                                                            }
                                                            var v119 float64 = 1.5
                                                            v120 := GenS15{F0: 684, F1: (((v113 + v113) + (v107 + "a")) + v107)}
                                                            _ = v119
                                                            _ = v120
                                                        default:
                                                            fmt.Println(fi108)
                                                        }
                                                    default:
                                                        v109 -= v109
                                                    }
                                                case "foo bar":
                                                    v121 := GenS16{F0: 1, F1: 226}
                                                    _ = v121
                                                }
                                            }
                                            v1 += 128
                                            _ = v113
                                        } else {
                                            if false {
                                                v122 := GenS11{F0: v1}
                                                _ = len(v3)
                                                switch fi108 {
                                                case 32768:
                                                    _ = len(v3)
                                                case -650:
                                                    switch (v107 + ((("foo bar" + (v107 + v107)) + ("a" + v107)) + ((v107 + (v107 + v107)) + v107))) {
                                                    case "a":
                                                        v107 = ((v107 + ((v107 + "\n\t") + v107)) + (v107 + (v107 + ("" + (v107 + (v107 + "hello"))))))
                                                    case "\n\t":
                                                        fmt.Println(fi102)
                                                    }
                                                default:
                                                    var v123 uint8 = (v1 - v1)
                                                    _ = v123
                                                }
                                                _ = v122
                                            }
                                            var v124 float32 = ((float32(v109) / -2.25) + (0.0 / 3.14159))
                                            v107 = v107
                                            v124 += v124
                                            v125 := GenS9{F0: 127, F1: fi108}
                                            _ = v124
                                            _ = v125
                                        }
                                    default:
                                        if (!false) {
                                            v1 = (v1 | v1)
                                            v1 = v1
                                        }
                                    }
                                }
                                if len(v2) > 0 {
                                    v2[0] = 0
                                }
                            }
                            fmt.Println(fi108, fi108)
                            fmt.Println(v109)
                            _ = fi108
                            _ = v109
                        }
                        _ = v106
                        _ = v107
                    }
                case "foo bar":
                    var v126 uint = 4294967295
                    _ = v126
                default:
                    var v127 bool = true
                    _ = v127
                }
                v128 := GenS3{F0: (908 != 65535), F1: -101, F2: 916}
                v1 -= (v1 - (^(v1)))
                _ = v128
            }
            var v129 float64 = (3.14159 / 1e10)
            _ = v103
            _ = v129
        }
        _ = fi102
    }
    fmt.Println(v1)
    switch (("hello" + ("foo bar" + "\n\t")) + "hello") {
    case "":
        v1 += v1
    }
    fmt.Println(v1)
    if true {
        var v130 float64 = (((float64((-2.25 / (((1e10 * (float32((0.0 / 1e10)) + 1e10)) + 0.1) + 3.14159))) + (0.0 * 3.14159)) - (1.5 - 0.0)) / 1.0)
        var v131 uint32 = 2147483647
        _ = v130
        _ = v131
    } else {
        var v132 int8 = 42
        for fi133 := 0; fi133 < 15; fi133++ {
            if (false || (!(!(false || ("hello" < "a"))))) {
                if ((fi133 + 800) != fi133) {
                    v134 := GenS17{F0: "\n\t", F1: 42}
                    if len(v2) > 0 {
                        v2[0] = 1
                    }
                    fmt.Println(fi133, v1, v132)
                    _ = v134
                } else {
                    var v135 float32 = 0.0
                    var v136 int32 = -1
                    fmt.Println(v1, v1, v136)
                    _ = v135
                    _ = v136
                }
                var v137 float32 = 3.14159
                if len(v2) > 0 {
                    v2[0] = 520
                }
                _ = v137
            }
            _ = fi133
        }
        _ = v132
    }
    if true {
        var v138 bool = true
        v138 = (!false)
        _ = v138
    } else {
        for fi139 := 0; fi139 < 6; fi139++ {
            for fi140 := 0; fi140 < 7; fi140++ {
                if false {
                    v1 = (v1 | (v1 - v1))
                    v1 += v1
                    if true {
                        v1 = (v1 << ((868) & 63))
                        if len(v3) > 0 {
                            v3[0] = (^(fi139))
                        }
                        fmt.Println(fi140, fi140, fi140)
                    } else {
                        var v141 uint = 693
                        var v142 int8 = 100
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        v142 -= int8(v1)
                        _ = v141
                        _ = v142
                    }
                }
                _ = fi140
            }
            v143 := GenS7{F0: 4294967295, F1: 1, F2: -2147483648}
            _ = fi139
            _ = v143
        }
        v144 := GenS2{F0: 2}
        _ = v144
    }
    var v145 int16 = -32768
    var v146 int = 469
    v145 -= v145
    switch "" {
    case "foo bar":
        for fi147 := 0; fi147 < 2; fi147++ {
            var v148 int8 = 100
            _ = fi147
            _ = v148
        }
    case "a":
        fmt.Println(v1)
    default:
        switch "a" {
        case "hello":
            if (18446744073709551615 <= 457) {
                v149 := GenS15{F0: 32767, F1: "a"}
                if (!false) {
                    if (v146 > v146) {
                        var v150 bool = (0 <= 42)
                        _ = v150
                    } else {
                        for fi151 := 0; fi151 < 16; fi151++ {
                            var v152 int16 = v145
                            fmt.Println(v146)
                            _ = fi151
                            _ = v152
                        }
                        switch "foo bar" {
                        case "":
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        case "\n\t":
                            v153 := GenS18{F0: true, F1: 736}
                            _ = v153
                        default:
                            if true {
                                var v154 uint8 = (v1 + v1)
                                if (true || true) {
                                    v145 = -1
                                }
                                var v155 uint64 = 334
                                _ = v154
                                _ = v155
                            }
                        }
                        fmt.Println(v146, v1, v1)
                    }
                    fmt.Println(v145, v145)
                }
                v1 -= 38
                _ = v149
            }
        case "a":
            if len(v2) > 0 {
                v2[0] = 0
            }
        }
    }
    var v156 uint16 = 0
    if len(v3) > 0 {
        v3[0] = 488
    }
    switch 4294967295 {
    case 320:
        if true {
            v157 := GenS18{F0: false, F1: (v156 | 0)}
            v158 := GenS1{F0: 0, F1: "a"}
            v145 -= (v145 | -32768)
            _ = v157
            _ = v158
        } else {
            if true {
                for fi159 := 0; fi159 < 20; fi159++ {
                    fmt.Println(v1)
                    _ = fi159
                }
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                for fi160 := 0; fi160 < 15; fi160++ {
                    if (!((0.1 / 1.0) >= 0.1)) {
                        _ = len(v3)
                    } else {
                        if (!(v1 >= v1)) {
                            if (18446744073709551615 == 1) {
                                fmt.Println(v156, fi160, v1)
                            } else {
                                _ = len(v2)
                            }
                        }
                        var v161 bool = false
                        _ = v161
                    }
                    fmt.Println(fi160, v146)
                    _ = fi160
                }
            } else {
                v162 := GenS19{F0: 101}
                _ = v162
            }
            if false {
                var v163 uint16 = v156
                var v164 int8 = 127
                _ = v163
                _ = v164
            } else {
                _ = len(v2)
            }
            fmt.Println(v146, v1, v145)
        }
    case -556:
        if ((v1 - 2) == 167) {
            var v165 uint16 = (^(v156))
            for fi166 := 0; fi166 < 9; fi166++ {
                switch (v146 / ((fi166) | 1)) {
                case -2147483648:
                    if false {
                        var v167 float32 = 1.5
                        v165 = (v156 >> ((772) & 63))
                        _ = v167
                    }
                case -9223372036854775808:
                    if ((int8(v145) >= 42) && false) {
                        v168 := GenS20{F0: v1, F1: 2147483647}
                        var v169 uint32 = 2147483647
                        if true {
                            var v170 bool = true
                            _ = v170
                        } else {
                            fmt.Println(v145)
                        }
                        _ = v168
                        _ = v169
                    }
                default:
                    for fi171 := 0; fi171 < 2; fi171++ {
                        v165 -= 42
                        if true {
                            var v172 string = "\n\t"
                            var v173 uint32 = 299
                            _ = v172
                            _ = v173
                        }
                        var v174 int = (^(fi171))
                        if (!(2 >= 751)) {
                            v174 = v174
                            var v175 int64 = 1
                            _ = v175
                        } else {
                            v174 -= (v174 >> ((2) & 63))
                            if (592 != 547) {
                                var v176 int32 = -290
                                _ = v176
                            } else {
                                var v177 int16 = v145
                                if false {
                                    var v178 uint16 = v156
                                    if true {
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                    }
                                    _ = v178
                                }
                                if (-529 < fi171) {
                                    if (-1 >= 100) {
                                        var v179 int64 = 735
                                        _ = v179
                                    } else {
                                        var v180 int = fi166
                                        _ = v180
                                    }
                                } else {
                                    var v181 float64 = (0.0 * float64(0.0))
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    _ = v181
                                }
                                _ = v177
                            }
                            if true {
                                var v182 int = (^(fi166))
                                _ = v182
                            } else {
                                v156 -= 455
                            }
                        }
                        _ = fi171
                        _ = v174
                    }
                }
                _ = fi166
            }
            _ = v165
        }
    default:
        var v183 uint64 = 330
        _ = v183
    }
    fmt.Println("v1:", v1)
    fmt.Println("v145:", v145)
    fmt.Println("v146:", v146)
    fmt.Println("v156:", v156)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v145
    _ = v146
    _ = v156
    _ = v2
    _ = v3
}
