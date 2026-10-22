package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint32; F1 uint32; F2 uint16 }
type GenS3 struct { F0 int8; F1 int64; F2 string }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int16; F1 bool }
type GenS6 struct { F0 int16; F1 int16 }
type GenS7 struct { F0 int8; F1 uint8; F2 uint32 }
type GenS8 struct { F0 uint64; F1 int64; F2 int8 }
type GenS9 struct { F0 uint64 }
type GenS10 struct { F0 uint8 }
type GenS11 struct { F0 uint8; F1 uint32; F2 uint64 }
type GenS12 struct { F0 int32 }
type GenS13 struct { F0 uint64 }
type GenS14 struct { F0 int8; F1 int32; F2 int32 }
type GenS15 struct { F0 int32 }
type GenS16 struct { F0 uint8 }
type GenS17 struct { F0 int64; F1 uint8; F2 int8 }
type GenS18 struct { F0 int64 }
type GenS19 struct { F0 int8; F1 string; F2 int8 }
type GenS20 struct { F0 string; F1 int16; F2 bool }
type GenS21 struct { F0 int64; F1 int16; F2 int }
type GenS22 struct { F0 string; F1 bool }
type GenS23 struct { F0 int8 }
type GenS24 struct { F0 uint64 }
type GenS25 struct { F0 uint8 }
type GenS26 struct { F0 uint8; F1 uint; F2 uint }

func main() {
    var v1 int64 = -1
    v2 := make(map[uint32]uint32)
    v3 := make(map[int]string)
    v4 := []int{9223372036854775807, -26, 922, 348}
    for fi5 := 0; fi5 < 11; fi5++ {
        var v6 string = "hello"
        v7 := v2[733]
        _ = fi5
        _ = v6
        _ = v7
    }
    v8 := GenS1{F0: 2}
    v9 := GenS1{F0: 59}
    for fi10 := 0; fi10 < 13; fi10++ {
        var v11 string = ("a" + "\n\t")
        var v12 int = 796
        _ = fi10
        _ = v11
        _ = v12
    }
    switch "foo bar" {
    case "\n\t":
        v13 := GenS1{F0: 42}
        _ = v13
    case "":
        fmt.Println(v1)
    case "foo bar":
        var v14 int = -209
        _ = v14
    }
    if (true && ("" > "foo bar")) {
        if false {
            v15 := GenS1{F0: 0}
            v1 = v1
            fmt.Println(v1)
            _ = v15
        } else {
            switch "\n\t" {
            case "":
                v1 -= (-(v1))
            case "\n\t":
                for fi16 := 0; fi16 < 15; fi16++ {
                    var v17 int = -921
                    _ = fi16
                    _ = v17
                }
            default:
                for fi18 := 0; fi18 < 17; fi18++ {
                    v19 := GenS1{F0: 42}
                    for fi20 := 0; fi20 < 6; fi20++ {
                        var v21 uint = 65535
                        fmt.Println(fi20, v1, fi20)
                        var v22 int32 = 0
                        if false {
                            var v23 int16 = 0
                            if (false || false) {
                                v24 := v3[fi20]
                                _ = v24
                            }
                            _ = v23
                        } else {
                            if (uint16(v21) > 32767) {
                                var v25 uint8 = 127
                                v1 = v1
                                if (false || (1 != -1)) {
                                    var v26 bool = true
                                    v27 := GenS1{F0: v25}
                                    _ = v26
                                    _ = v27
                                }
                                _ = v25
                            } else {
                                if (!false) {
                                    v21 = (v21 / ((v21) | 1))
                                    fmt.Println(v21, v22, fi20)
                                    var v28 uint = v21
                                    v22 = (v22 % (((^(v22))) | 1))
                                    v29 := GenS1{F0: 147}
                                    v21 = v28
                                    fmt.Println(fi18, v1)
                                    switch (fi18 | (fi18 - (fi20 - fi20))) {
                                    case 4294967295:
                                        if (127 == -47) {
                                            v2[2147483647] = 2147483647
                                        } else {
                                            if (!(!false)) {
                                                v21 = v21
                                                _ = len(v4)
                                                v30 := GenS1{F0: 2}
                                                _ = v30
                                            }
                                            v31 := GenS2{F0: 42, F1: 0, F2: 255}
                                            _ = v31
                                        }
                                    case 9223372036854775807:
                                        if true {
                                            v32 := GenS2{F0: 65535, F1: 2147483647, F2: 691}
                                            v22 = (v22 & v22)
                                            _ = v32
                                        } else {
                                            v33 := GenS2{F0: 0, F1: 42, F2: 32767}
                                            v28 = 358
                                            switch "\n\t" {
                                            case "hello":
                                                switch "hello" {
                                                case "hello":
                                                    switch 2147483647 {
                                                    case 1:
                                                        if (true && (!false)) {
                                                            if (2 != 806) {
                                                                var v34 int64 = v1
                                                                switch 42 {
                                                                case 727:
                                                                    if true {
                                                                        switch 32767 {
                                                                        case 1:
                                                                            var v35 int8 = -1
                                                                            _ = v35
                                                                        }
                                                                    } else {
                                                                        var v36 uint = (v28 - (v21 - v21))
                                                                        _ = v36
                                                                    }
                                                                case -9223372036854775808:
                                                                    v37 := GenS3{F0: 0, F1: (v1 - (v1 - 186)), F2: "\n\t"}
                                                                    delete(v3, (^(fi20)))
                                                                    delete(v2, 1)
                                                                    _ = v37
                                                                default:
                                                                    var v38 int16 = 1000
                                                                    if len(v4) >= 2 {
                                                                        v4 = v4[1:len(v4)]
                                                                    }
                                                                    v22 += (-(v22))
                                                                    _ = v38
                                                                }
                                                                _ = v34
                                                            }
                                                        } else {
                                                            v21 = 4294967295
                                                            switch fi20 {
                                                            case -2147483648:
                                                                v39 := GenS4{F0: (fi20 | (fi18 >> ((v28) & 63)))}
                                                                _ = v39
                                                            default:
                                                                v1 -= v1
                                                            }
                                                        }
                                                    case 0:
                                                        switch 42 {
                                                        case 0:
                                                            var v40 uint32 = 602
                                                            _ = v40
                                                        case -217:
                                                            fmt.Println(v1, v21, v21)
                                                        case 32768:
                                                            v41 := GenS5{F0: 42, F1: (false || true)}
                                                            _ = v41
                                                        }
                                                    case -2147483648:
                                                        fmt.Println(fi18)
                                                    }
                                                }
                                            }
                                            fmt.Println(v1, v28)
                                            v42 := GenS6{F0: -128, F1: int16(fi20)}
                                            var v43 int8 = -128
                                            _ = v33
                                            _ = v42
                                            _ = v43
                                        }
                                    case -9223372036854775808:
                                        switch "" {
                                        case "a":
                                            if ("a" == ("a" + ("\n\t" + ""))) {
                                                if true {
                                                    if (!true) {
                                                        v28 = v28
                                                        var v44 uint8 = 128
                                                        _ = v44
                                                    }
                                                    v1 = v1
                                                    if true {
                                                        v45 := v2[4294967295]
                                                        fmt.Println(fi18, fi20, v1)
                                                        _ = v45
                                                    }
                                                } else {
                                                    v21 = (v28 << (((v21 >> ((v21) & 63))) & 63))
                                                    var v46 int64 = 42
                                                    _ = v46
                                                }
                                            } else {
                                                var v47 float32 = 1.5
                                                v48 := GenS7{F0: -89, F1: 225, F2: 42}
                                                _ = v47
                                                _ = v48
                                            }
                                        case "":
                                            if (!false) {
                                                var v49 int8 = 0
                                                v22 = v22
                                                if false {
                                                    var v50 uint16 = 2
                                                    switch (fi20 | (fi18 ^ int(v22))) {
                                                    case 9223372036854775807:
                                                        v50 = (v50 * 42)
                                                    default:
                                                        if len(v4) > 0 {
                                                            v4[0] = 0
                                                        }
                                                    }
                                                    v50 = v50
                                                    _ = v50
                                                } else {
                                                    v51 := GenS8{F0: 1, F1: (v1 / (((v1 % (((v1 & (v1 | (-(v1))))) | 1))) | 1)), F2: v49}
                                                    if (1e10 > -2.25) {
                                                        var v52 int = (^(fi18))
                                                        _ = v52
                                                    }
                                                    switch "a" {
                                                    case "\n\t":
                                                        v21 = v28
                                                    default:
                                                        var v53 string = "\n\t"
                                                        _ = v53
                                                    }
                                                    delete(v3, fi20)
                                                    _ = v51
                                                }
                                                _ = v49
                                            }
                                        default:
                                            v54 := GenS9{F0: 4294967295}
                                            _ = v54
                                        }
                                        if true {
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                            v55 := GenS5{F0: 0, F1: false}
                                            var v56 uint = v21
                                            v56 += 49
                                            if false {
                                                v57 := v3[-2147483648]
                                                v28 -= v28
                                                _ = v57
                                            } else {
                                                var v58 int64 = (v1 | (-(v1)))
                                                var v59 uint64 = 655
                                                v60 := GenS10{F0: uint8(v59)}
                                                fmt.Println(fi18, v56)
                                                _ = len(v3)
                                                v58 += (^(v1))
                                                _ = v58
                                                _ = v59
                                                _ = v60
                                            }
                                            _ = v55
                                            _ = v56
                                        } else {
                                            v61 := GenS11{F0: 42, F1: 0, F2: 1}
                                            v3[fi20] = "hello"
                                            var v62 uint8 = 1
                                            _ = v61
                                            _ = v62
                                        }
                                    default:
                                        var v63 int64 = (v1 * 9223372036854775807)
                                        _ = v63
                                    }
                                    v21 += (v21 + (^(v21)))
                                    if (!((!true) || (true && ((("a" + "foo bar") + ((((("foo bar" + ("" + "hello")) + "") + "foo bar") + "\n\t") + "hello")) == "hello")))) {
                                        if (((v22 % (((v22 * v22)) | 1)) >= (v22 + 119)) || false) {
                                            v28 = (^(v28))
                                        } else {
                                            v21 += (^(v28))
                                        }
                                    }
                                    _ = v28
                                    _ = v29
                                }
                                fmt.Println(v22)
                                if (155 == 38) {
                                    fmt.Println(v1, fi18, v1)
                                    delete(v3, 4294967295)
                                    switch "" {
                                    case "hello":
                                        var v64 uint64 = 588
                                        _ = v64
                                    case "":
                                        if false {
                                            fmt.Println(fi20)
                                            var v65 uint8 = 255
                                            _ = v65
                                        }
                                    }
                                    v21 += (v21 | 730)
                                    var v66 int16 = 42
                                    v67 := GenS8{F0: 42, F1: (v1 / (((v1 + v1)) | 1)), F2: 42}
                                    var v68 uint8 = 1
                                    _ = v66
                                    _ = v67
                                    _ = v68
                                } else {
                                    delete(v2, 590)
                                    switch "a" {
                                    case "foo bar":
                                        var v69 string = ("\n\t" + "\n\t")
                                        _ = v69
                                    }
                                    fmt.Println(fi18, fi18, v21)
                                }
                                var v70 bool = (true || (v1 == -1))
                                _ = v70
                            }
                            if true {
                                v71 := GenS7{F0: 100, F1: 0, F2: 2}
                                v72 := GenS12{F0: v22}
                                _ = v71
                                _ = v72
                            } else {
                                var v73 int16 = -1
                                var v74 string = "a"
                                var v75 uint64 = 1
                                v22 = (v22 << (((^(v21))) & 63))
                                if len(v4) > 0 {
                                    v4[0] = (^(fi18))
                                }
                                _ = v73
                                _ = v74
                                _ = v75
                            }
                            v76 := GenS6{F0: -698, F1: 432}
                            _ = v76
                        }
                        switch fi18 {
                        case 792:
                            switch ("hello" + "") {
                            case "\n\t":
                                v22 -= v22
                            case "a":
                                switch "foo bar" {
                                case "a":
                                    v77 := GenS13{F0: 4294967295}
                                    _ = v77
                                case "foo bar":
                                    delete(v2, 525)
                                case "\n\t":
                                    v78 := GenS14{F0: 0, F1: (v22 ^ (v22 * v22)), F2: v22}
                                    _ = v78
                                }
                            case "foo bar":
                                v1 -= v1
                            default:
                                if (!(798 > 251)) {
                                    var v79 uint64 = 697
                                    _ = v79
                                }
                            }
                        case 128:
                            var v80 int32 = (v22 / ((v22) | 1))
                            _ = v80
                        default:
                            if true {
                                var v81 float32 = -2.25
                                v22 -= -751
                                _ = v81
                            } else {
                                if (v22 == 1) {
                                    var v82 uint64 = 4294967295
                                    v83 := GenS11{F0: 248, F1: 965, F2: v82}
                                    var v84 uint32 = 2
                                    _ = v82
                                    _ = v83
                                    _ = v84
                                } else {
                                    if ((v22 << (((v21 & (v21 ^ 12))) & 63)) != v22) {
                                        fmt.Println(v21, v21)
                                        if len(v4) > 0 {
                                            v4[0] = (^(fi18))
                                        }
                                        var v85 int32 = (v22 ^ (-(v22)))
                                        v86 := GenS15{F0: v22}
                                        v85 += (v22 - 42)
                                        if false {
                                            if true {
                                                if (-1 < -64) {
                                                    if (2 == (^(v21))) {
                                                        if true {
                                                            v87 := GenS7{F0: 1, F1: 0, F2: 2}
                                                            fmt.Println(v22, fi18)
                                                            var v88 int32 = (^(v22))
                                                            _ = v87
                                                            _ = v88
                                                        }
                                                        delete(v2, 2147483648)
                                                        v89 := GenS16{F0: 42}
                                                        _ = v89
                                                    } else {
                                                        fmt.Println(v22, v22)
                                                    }
                                                    if len(v4) >= 2 {
                                                        v4 = v4[1:len(v4)]
                                                    }
                                                    switch (fi18 - fi18) {
                                                    case -722:
                                                        v90 := GenS17{F0: -1, F1: 136, F2: 68}
                                                        _ = v90
                                                    case 959:
                                                        _ = len(v4)
                                                    }
                                                }
                                                v91 := GenS18{F0: (v1 | v1)}
                                                v92 := GenS17{F0: v1, F1: 128, F2: 75}
                                                _ = v91
                                                _ = v92
                                            }
                                            v85 = (v85 - -1)
                                        }
                                        _ = v85
                                        _ = v86
                                    }
                                    if (v21 != 1) {
                                        var v93 int = fi20
                                        v94 := GenS19{F0: -48, F1: "hello", F2: 0}
                                        if (false || true) {
                                            switch (v93 % (((fi20 / (((fi20 / (((^(fi18))) | 1))) | 1))) | 1)) {
                                            case -272:
                                                var v95 int8 = 1
                                                _ = v95
                                            case 1:
                                                fmt.Println(v93, v22, v21)
                                                var v96 uint8 = 1
                                                fmt.Println(v22)
                                                _ = v96
                                            case -108:
                                                v21 = v21
                                            }
                                            v97 := GenS9{F0: 114}
                                            var v98 uint64 = 0
                                            var v99 float32 = (float32(-2.25) - ((((3.14159 + 0.1) + (float32(3.14159) * float32((((1e10 + 0.0) / float64(1e10)) * 1.5)))) * -2.25) + 0.0))
                                            v100 := GenS12{F0: (v22 - 0)}
                                            if false {
                                                v98 -= (v98 % (((v98 >> (((v21 >> ((v21) & 63))) & 63))) | 1))
                                                var v101 uint = (v21 / ((v21) | 1))
                                                _ = v101
                                            } else {
                                                if true {
                                                    v22 = v22
                                                    if (!(!false)) {
                                                        if (!((!(104 >= 0)) || ((v99 * (v99 / ((0.1 - ((v99 + ((v99 / -2.25) * ((v99 / 1.0) / v99))) * ((1.5 * float32(float64(float32(-2.25)))) * ((((v99 / 1.0) * 1.5) - 0.0) + -2.25)))) + v99))) >= (-2.25 - (v99 * (-2.25 + (0.0 - v99))))))) {
                                                            var v102 float64 = 1.5
                                                            switch "\n\t" {
                                                            case "":
                                                                if false {
                                                                    var v103 uint32 = uint32(v21)
                                                                    v104 := GenS20{F0: "hello", F1: -32768, F2: ("foo bar" < "a")}
                                                                    _ = v103
                                                                    _ = v104
                                                                }
                                                            case "\n\t":
                                                                v4 = append(v4, fi18)
                                                            case "a":
                                                                var v105 uint32 = 0
                                                                _ = v105
                                                            }
                                                            if (100 < 1) {
                                                                if (-128 <= -1) {
                                                                    v106 := GenS11{F0: 127, F1: 4294967295, F2: 1}
                                                                    _ = v106
                                                                }
                                                                if (4294967295 == 41) {
                                                                    v22 = (^(v22))
                                                                    var v107 int16 = 42
                                                                    _ = v107
                                                                } else {
                                                                    switch "foo bar" {
                                                                    case "a":
                                                                        v102 += ((-2.25 + (1.5 / 1.5)) / 0.1)
                                                                    case "":
                                                                        var v108 uint16 = 65535
                                                                        _ = v108
                                                                    case "\n\t":
                                                                        if false {
                                                                            v4 = append(v4, 1)
                                                                            switch (fi20 | v93) {
                                                                            case -1:
                                                                                var v109 bool = (v98 != v98)
                                                                                _ = v109
                                                                            case 1:
                                                                                var v110 uint16 = 997
                                                                                _ = v110
                                                                            }
                                                                        }
                                                                    }
                                                                }
                                                                var v111 int8 = -100
                                                                _ = v111
                                                            } else {
                                                                v99 = float32(v102)
                                                                var v112 uint32 = 42
                                                                _ = v112
                                                            }
                                                            v113 := GenS1{F0: 158}
                                                            _ = v102
                                                            _ = v113
                                                        } else {
                                                            if true {
                                                                v114 := GenS16{F0: 191}
                                                                _ = v114
                                                            } else {
                                                                v22 -= -1
                                                                switch "a" {
                                                                case "\n\t":
                                                                    fmt.Println(fi18)
                                                                }
                                                            }
                                                        }
                                                        v21 = (v21 / (((v21 & v21)) | 1))
                                                        var v115 uint64 = 18446744073709551615
                                                        _ = v115
                                                    } else {
                                                        switch "hello" {
                                                        case "":
                                                            switch 32768 {
                                                            case 128:
                                                                v116 := GenS5{F0: -859, F1: (v1 > 0)}
                                                                _ = v116
                                                            case 42:
                                                                var v117 string = "\n\t"
                                                                _ = v117
                                                            case 996:
                                                                switch (("\n\t" + (("\n\t" + "foo bar") + "\n\t")) + (("a" + "foo bar") + "foo bar")) {
                                                                case "\n\t":
                                                                    var v118 int16 = 42
                                                                    _ = v118
                                                                case "":
                                                                    var v119 int8 = 0
                                                                    _ = v119
                                                                case "hello":
                                                                    switch "a" {
                                                                    case "foo bar":
                                                                        v120 := GenS21{F0: v1, F1: int16(v22), F2: fi20}
                                                                        _ = v120
                                                                    case "a":
                                                                        if false {
                                                                            v99 += v99
                                                                        }
                                                                        fmt.Println(fi18, fi18)
                                                                    case "hello":
                                                                        v22 -= -575
                                                                    }
                                                                }
                                                            default:
                                                                fmt.Println(v98, v99, v21)
                                                            }
                                                        case "\n\t":
                                                            var v121 uint16 = 0
                                                            v21 += 1
                                                            if (v1 == 0) {
                                                                fmt.Println(v99)
                                                                switch (("hello" + ("hello" + "a")) + "foo bar") {
                                                                case "a":
                                                                    v98 = v98
                                                                default:
                                                                    if (3.14159 < 1e10) {
                                                                        var v122 float32 = v99
                                                                        var v123 uint16 = (v121 ^ (v121 ^ (v121 / (((v121 - (v121 & uint16(v98)))) | 1))))
                                                                        _ = v122
                                                                        _ = v123
                                                                    }
                                                                }
                                                                v98 = v98
                                                            }
                                                            _ = v121
                                                        case "hello":
                                                            var v124 bool = true
                                                            _ = v124
                                                        }
                                                    }
                                                } else {
                                                    v98 -= v98
                                                    switch (fi18 * (fi20 ^ 9223372036854775807)) {
                                                    case 1:
                                                        fmt.Println(v99, v98, v21)
                                                    default:
                                                        var v125 uint64 = 4294967295
                                                        _ = v125
                                                    }
                                                    v98 += 1
                                                    if (32767 > 1) {
                                                        var v126 uint16 = uint16(v21)
                                                        v127 := GenS22{F0: "hello", F1: (32767 == 1000)}
                                                        delete(v3, 32767)
                                                        if ((v21 % ((660) | 1)) != 2) {
                                                            switch ("" + "foo bar") {
                                                            case "hello":
                                                                fmt.Println(v22)
                                                                var v128 int16 = -32768
                                                                _ = v128
                                                            case "":
                                                                if (!(true && (1e10 < (v99 + v99)))) {
                                                                    v2[0] = 2147483648
                                                                    v1 += 308
                                                                }
                                                            default:
                                                                switch fi18 {
                                                                case -2147483648:
                                                                    var v129 int64 = (v1 >> (((^(v21))) & 63))
                                                                    fmt.Println(v129)
                                                                    _ = v129
                                                                case 1:
                                                                    v130 := GenS16{F0: 87}
                                                                    _ = v130
                                                                default:
                                                                    v21 -= 235
                                                                }
                                                                if true {
                                                                    fmt.Println(fi18)
                                                                } else {
                                                                    fmt.Println(v93)
                                                                    if true {
                                                                        var v131 string = "hello"
                                                                        v132 := GenS23{F0: -128}
                                                                        var v133 bool = (true && (100 != 100))
                                                                        fmt.Println(v131)
                                                                        _ = v131
                                                                        _ = v132
                                                                        _ = v133
                                                                    } else {
                                                                        if ((((3.14159 - -2.25) < (3.14159 - 1.5)) && false) && false) {
                                                                            v4 = append(v4, 32767)
                                                                            if ((v22 % (((v22 / (((v22 * v22)) | 1))) | 1)) != v22) {
                                                                                v22 = v22
                                                                                var v134 int64 = v1
                                                                                _ = v134
                                                                            }
                                                                            if false {
                                                                                v135 := v3[v93]
                                                                                _ = len(v4)
                                                                                _ = v135
                                                                            } else {
                                                                                var v136 uint8 = 1
                                                                                v137 := GenS24{F0: v98}
                                                                                var v138 uint32 = 396
                                                                                _ = v136
                                                                                _ = v137
                                                                                _ = v138
                                                                            }
                                                                        }
                                                                    }
                                                                    if true {
                                                                        _ = len(v2)
                                                                        var v139 uint16 = 1
                                                                        var v140 int16 = 32767
                                                                        _ = v139
                                                                        _ = v140
                                                                    }
                                                                }
                                                                if false {
                                                                    _ = len(v2)
                                                                    switch "\n\t" {
                                                                    case "\n\t":
                                                                        var v141 int8 = 42
                                                                        _ = v141
                                                                    default:
                                                                        switch v93 {
                                                                        case 32768:
                                                                            var v142 uint8 = 42
                                                                            _ = v142
                                                                        default:
                                                                            _ = len(v4)
                                                                        }
                                                                    }
                                                                }
                                                            }
                                                        } else {
                                                            var v143 int8 = 0
                                                            _ = v143
                                                        }
                                                        _ = v126
                                                        _ = v127
                                                    } else {
                                                        _ = len(v3)
                                                    }
                                                    switch 462 {
                                                    case 1:
                                                        var v144 uint32 = 2
                                                        v22 -= (-(v22))
                                                        var v145 uint32 = (^(v144))
                                                        v98 = 42
                                                        v1 = (v1 + int64(v144))
                                                        _ = v144
                                                        _ = v145
                                                    }
                                                    var v146 int = (^(fi20))
                                                    var v147 uint64 = v98
                                                    v148 := GenS2{F0: 0, F1: 2, F2: 115}
                                                    switch "\n\t" {
                                                    case "hello":
                                                        var v149 uint32 = 942
                                                        _ = v149
                                                    default:
                                                        v2[628] = 0
                                                    }
                                                    _ = v146
                                                    _ = v147
                                                    _ = v148
                                                }
                                                var v150 int64 = v1
                                                v151 := GenS12{F0: v22}
                                                _ = v150
                                                _ = v151
                                            }
                                            _ = v100
                                            _ = v97
                                            _ = v98
                                            _ = v99
                                        } else {
                                            if false {
                                                v93 += (fi20 << ((2) & 63))
                                                fmt.Println(v93, fi20)
                                                if (false || (v21 != 65535)) {
                                                    var v152 int32 = v22
                                                    _ = v152
                                                }
                                            }
                                            var v153 int64 = v1
                                            _ = v153
                                        }
                                        var v154 string = "a"
                                        v154 = (v154 + ("\n\t" + v154))
                                        v22 = (-(v22))
                                        if len(v4) > 0 {
                                            v4[0] = (-(v93))
                                        }
                                        _ = v154
                                        _ = v93
                                        _ = v94
                                    } else {
                                        v22 = (v22 % (((v22 >> (((v21 ^ (v21 - (v21 - v21)))) & 63))) | 1))
                                    }
                                    v1 = (v1 / ((0) | 1))
                                }
                                switch "hello" {
                                case "\n\t":
                                    fmt.Println(v1)
                                }
                                switch (fi18 % (((fi18 | fi20)) | 1)) {
                                case -968:
                                    var v155 uint32 = uint32(v21)
                                    _ = v155
                                case -9223372036854775808:
                                    if (!(64 != 2147483647)) {
                                        _ = len(v4)
                                        v21 = v21
                                        if false {
                                            v22 += (v22 * v22)
                                            _ = len(v3)
                                            switch "hello" {
                                            case "a":
                                                if false {
                                                    fmt.Println(v1)
                                                    fmt.Println(fi18, fi18)
                                                    fmt.Println(v1, fi20)
                                                    switch "a" {
                                                    case "hello":
                                                        if len(v4) > 0 {
                                                            v4[0] = (fi20 + 127)
                                                        }
                                                    default:
                                                        fmt.Println(v1, v1, v21)
                                                    }
                                                    var v156 uint8 = 0
                                                    _ = v156
                                                } else {
                                                    switch ("hello" + "") {
                                                    case "hello":
                                                        switch fi20 {
                                                        case 2147483647:
                                                            v157 := GenS23{F0: 100}
                                                            _ = v157
                                                        case 294:
                                                            v21 -= (v21 ^ (v21 - (^(v21))))
                                                        }
                                                    default:
                                                        switch 1 {
                                                        case -1:
                                                            switch ("foo bar" + "foo bar") {
                                                            case "\n\t":
                                                                fmt.Println(fi20)
                                                            case "a":
                                                                switch ("" + "hello") {
                                                                case "a":
                                                                    v22 = v22
                                                                case "foo bar":
                                                                    var v158 int8 = 0
                                                                    _ = v158
                                                                default:
                                                                    v4 = append(v4, fi18)
                                                                }
                                                            case "hello":
                                                                switch fi20 {
                                                                case 1:
                                                                    var v159 uint16 = 255
                                                                    _ = v159
                                                                default:
                                                                    var v160 int32 = v22
                                                                    _ = v160
                                                                }
                                                            default:
                                                                if (false && (fi18 > fi20)) {
                                                                    v1 -= v1
                                                                    v161 := GenS25{F0: 42}
                                                                    _ = v161
                                                                } else {
                                                                    v22 = (^(v22))
                                                                    if len(v4) > 0 {
                                                                        v4[0] = (fi18 + (fi18 / (((-(fi20))) | 1)))
                                                                    }
                                                                    _ = len(v4)
                                                                    switch (fi20 % ((fi18) | 1)) {
                                                                    case -538:
                                                                        var v162 int32 = 727
                                                                        _ = v162
                                                                    case 957:
                                                                        fmt.Println(v21, v22, v21)
                                                                    case 313:
                                                                        var v163 float64 = 3.14159
                                                                        _ = v163
                                                                    default:
                                                                        v21 -= (v21 << ((0) & 63))
                                                                    }
                                                                    v22 -= (v22 ^ v22)
                                                                }
                                                            }
                                                        case -919:
                                                            _ = len(v4)
                                                        }
                                                    }
                                                }
                                            default:
                                                v164 := GenS15{F0: -314}
                                                _ = v164
                                            }
                                        }
                                        if (true || false) {
                                            var v165 uint = (v21 - (^(v21)))
                                            v22 += (v22 | (v22 >> ((v165) & 63)))
                                            v166 := GenS21{F0: v1, F1: 473, F2: int(v165)}
                                            _ = v165
                                            _ = v166
                                        }
                                        var v167 string = ("a" + "hello")
                                        v168 := GenS16{F0: 1}
                                        _ = v167
                                        _ = v168
                                    } else {
                                        fmt.Println(v1)
                                        v21 = 724
                                        v1 = -9223372036854775808
                                        if true {
                                            var v169 int8 = 105
                                            _ = v169
                                        }
                                    }
                                case -1:
                                    v21 = v21
                                }
                            }
                        }
                        if (false && true) {
                            switch (fi20 % ((fi20) | 1)) {
                            case 2147483647:
                                v170 := GenS21{F0: (v1 >> ((65535) & 63)), F1: 307, F2: fi20}
                                _ = v170
                            case 978:
                                if false {
                                    v1 = v1
                                    v1 = -784
                                }
                            }
                        } else {
                            fmt.Println(v21, fi20)
                            var v171 uint16 = 1
                            v22 += v22
                            _ = v171
                        }
                        _ = fi20
                        _ = v21
                        _ = v22
                    }
                    fmt.Println(v1, fi18)
                    _ = fi18
                    _ = v19
                }
            }
            v1 -= -1
        }
    }
    v172 := GenS10{F0: 0}
    _ = len(v4)
    v4 = append(v4, 42)
    for fi173 := 0; fi173 < 5; fi173++ {
        switch "a" {
        case "a":
            fmt.Println(v1)
        case "hello":
            for fi174 := 0; fi174 < 2; fi174++ {
                switch 9223372036854775807 {
                case -968:
                    var v175 int16 = 1000
                    delete(v2, 65535)
                    v176 := GenS5{F0: v175, F1: ((65535 == 65535) && true)}
                    _ = v175
                    _ = v176
                case 42:
                    var v177 uint8 = 2
                    _ = v177
                default:
                    if (0.1 < (3.14159 - float64(-2.25))) {
                        if (255 > 127) {
                            v178 := GenS26{F0: 84, F1: 830, F2: 65535}
                            var v179 int16 = 0
                            fmt.Println(v179, fi174, fi173)
                            switch -9223372036854775808 {
                            case -9223372036854775808:
                                var v180 bool = ("foo bar" >= "\n\t")
                                _ = v180
                            case -939:
                                v179 -= v179
                            case 1:
                                v1 = -9223372036854775808
                                var v181 uint32 = 2147483648
                                _ = v181
                            }
                            v182 := v2[2147483648]
                            fmt.Println(v182)
                            v183 := v3[(^(fi173))]
                            _ = v178
                            _ = v179
                            _ = v182
                            _ = v183
                        } else {
                            delete(v3, -1)
                            v1 -= v1
                            var v184 string = "hello"
                            _ = v184
                        }
                        v1 = v1
                    } else {
                        v1 -= (v1 ^ v1)
                    }
                }
                _ = fi174
            }
        }
        _ = fi173
    }
    switch "a" {
    case "\n\t":
        v3[32768] = "a"
    }
    if false {
        var v185 int32 = 207
        var v186 float64 = 1.5
        _ = v185
        _ = v186
    }
    var v187 int64 = v1
    switch ("hello" + "foo bar") {
    case "foo bar":
        fmt.Println(v187, v187)
    case "a":
        v1 = v187
    default:
        fmt.Println(v187)
    }
    for fi188 := 0; fi188 < 9; fi188++ {
        v1 = v187
        delete(v2, 396)
        _ = fi188
    }
    v189 := GenS1{F0: 2}
    fmt.Println("v1:", v1)
    fmt.Println("v187:", v187)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v172
    _ = v187
    _ = v189
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
