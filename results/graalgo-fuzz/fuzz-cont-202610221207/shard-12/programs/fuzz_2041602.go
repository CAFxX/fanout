package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8 }
type GenS2 struct { F0 int64; F1 string; F2 uint16 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int; F1 uint32 }
type GenS6 struct { F0 int; F1 bool; F2 int8 }
type GenS7 struct { F0 uint16; F1 int32 }
type GenS8 struct { F0 uint; F1 uint64 }
type GenS9 struct { F0 int32; F1 uint16; F2 uint32 }
type GenS10 struct { F0 int; F1 string; F2 uint16 }
type GenS11 struct { F0 uint64; F1 int32 }

func main() {
    var v1 int16 = 290
    var v2 uint8 = 1
    v3 := []uint32{2147483647, 105, 981}
    var v4 int16 = -915
    var v5 uint32 = 0
    var v6 uint32 = 42
    if true {
        fmt.Println(v5, v6)
        v4 = v4
        switch "foo bar" {
        case "\n\t":
            fmt.Println(v4, v2)
        default:
            fmt.Println(v5, v1)
        }
    }
    v5 += (v5 + v6)
    if false {
        switch "hello" {
        case "":
            if ((v5 < 4294967295) && false) {
                var v7 int = 1
                v6 = (v5 + v6)
                _ = v7
            }
        case "a":
            v1 -= v1
        }
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    if true {
        if true {
            var v8 uint8 = (v2 & v2)
            if (false || true) {
                _ = len(v3)
            } else {
                var v9 uint = 18446744073709551615
                v10 := GenS1{F0: false, F1: 0}
                fmt.Println(v9, v5)
                _ = v10
                _ = v9
            }
            _ = v8
        } else {
            var v11 int = -1
            for fi12 := 0; fi12 < 5; fi12++ {
                v4 = 1
                v13 := GenS1{F0: false, F1: int8(v6)}
                var v14 string = "hello"
                _ = fi12
                _ = v13
                _ = v14
            }
            _ = v11
        }
        switch -657 {
        case 2147483647:
            switch 2147483647 {
            case -9223372036854775808:
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            case 128:
                var v15 uint64 = 2
                _ = v15
            default:
                var v16 int8 = 92
                _ = v16
            }
        }
    } else {
        v5 = v5
    }
    v2 -= (v2 >> ((1) & 63))
    for fi17 := 0; fi17 < 5; fi17++ {
        fmt.Println(v5, v4, v5)
        for fi18 := 0; fi18 < 16; fi18++ {
            var v19 uint8 = (v2 + (v2 ^ 245))
            _ = fi18
            _ = v19
        }
        for fi20 := 0; fi20 < 10; fi20++ {
            switch (fi20 - fi20) {
            case 32767:
                if true {
                    if (255 <= v2) {
                        var v21 string = "foo bar"
                        _ = v21
                    } else {
                        if (v1 != v4) {
                            var v22 int32 = 0
                            v23 := GenS1{F0: (42 == 0), F1: -11}
                            var v24 uint = 2
                            _ = v22
                            _ = v23
                            _ = v24
                        }
                        var v25 uint16 = 255
                        switch "" {
                        case "\n\t":
                            switch 1 {
                            case 332:
                                var v26 int64 = -9223372036854775808
                                _ = v26
                            }
                        case "a":
                            v1 += v1
                            var v27 uint8 = (v2 & v2)
                            if ((130 < uint(v25)) || true) {
                                var v28 bool = false
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                _ = v28
                            } else {
                                v3 = append(v3, v5)
                                v29 := GenS1{F0: (int16(v25) != v1), F1: 127}
                                switch "hello" {
                                case "foo bar":
                                    var v30 uint16 = 42
                                    _ = v30
                                case "":
                                    switch (("" + ("foo bar" + "foo bar")) + ("" + ("\n\t" + ""))) {
                                    case "\n\t":
                                        fmt.Println(fi20, v5, v6)
                                    case "a":
                                        if true {
                                            switch ("a" + ("hello" + "foo bar")) {
                                            case "foo bar":
                                                var v31 uint8 = v2
                                                _ = v31
                                            case "\n\t":
                                                if true {
                                                    var v32 int8 = -1
                                                    var v33 uint32 = v5
                                                    if len(v3) > 0 {
                                                        v3[0] = v5
                                                    }
                                                    _ = v32
                                                    _ = v33
                                                } else {
                                                    if false {
                                                        v1 = int16(v5)
                                                        switch fi17 {
                                                        case -904:
                                                            if ((fi17 <= (fi20 ^ (fi20 >> ((1) & 63)))) || (!false)) {
                                                                if len(v3) >= 2 {
                                                                    v3 = v3[1:len(v3)]
                                                                }
                                                                fmt.Println(v2, v4, v25)
                                                            } else {
                                                                var v34 uint = 0
                                                                v1 = -1
                                                                v2 += v27
                                                                var v35 int64 = -1
                                                                switch fi17 {
                                                                case 208:
                                                                    var v36 int16 = 243
                                                                    _ = v36
                                                                case 42:
                                                                    var v37 uint64 = 18446744073709551615
                                                                    _ = v37
                                                                default:
                                                                    v4 = (v4 << ((42) & 63))
                                                                }
                                                                _ = v34
                                                                _ = v35
                                                            }
                                                        case 32768:
                                                            if (false || false) {
                                                                v38 := GenS1{F0: false, F1: 100}
                                                                _ = v38
                                                            }
                                                        case 57:
                                                            _ = len(v3)
                                                            switch (-(fi20)) {
                                                            case -1:
                                                                v2 = v2
                                                            case 1:
                                                                v2 += (v27 ^ v27)
                                                            case 42:
                                                                if len(v3) > 0 {
                                                                    v3[0] = (v5 * uint32(v27))
                                                                }
                                                            }
                                                        default:
                                                            switch "\n\t" {
                                                            case "\n\t":
                                                                v27 += 0
                                                            case "a":
                                                                var v39 int32 = 1
                                                                switch "" {
                                                                case "hello":
                                                                    var v40 int64 = -9223372036854775808
                                                                    _ = v40
                                                                case "foo bar":
                                                                    var v41 float32 = 0.0
                                                                    _ = v41
                                                                case "\n\t":
                                                                    var v42 string = "foo bar"
                                                                    _ = v42
                                                                }
                                                                _ = v39
                                                            }
                                                        }
                                                        fmt.Println(v1, v6, v1)
                                                    }
                                                    switch fi17 {
                                                    case -980:
                                                        var v43 uint32 = v6
                                                        _ = v43
                                                    case 127:
                                                        var v44 float32 = ((1.5 + (0.1 - 0.0)) + 1e10)
                                                        _ = v44
                                                    case 32767:
                                                        var v45 int16 = (v4 / (((-(v1))) | 1))
                                                        _ = v45
                                                    default:
                                                        v27 += 128
                                                    }
                                                }
                                            default:
                                                v4 += (v1 | (v4 * (v4 * (v1 & v1))))
                                            }
                                        } else {
                                            var v46 string = "\n\t"
                                            var v47 int64 = -9223372036854775808
                                            _ = v46
                                            _ = v47
                                        }
                                    default:
                                        v1 = -1
                                    }
                                default:
                                    if ("foo bar" > "foo bar") {
                                        v25 -= uint16(v6)
                                    } else {
                                        var v48 int8 = 1
                                        v27 -= (v27 % (((v2 * v2)) | 1))
                                        fmt.Println(v25)
                                        _ = v48
                                    }
                                }
                                _ = v29
                            }
                            _ = v27
                        default:
                            var v49 int = fi17
                            _ = v49
                        }
                        _ = len(v3)
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        _ = v25
                    }
                } else {
                    var v50 string = ("foo bar" + "")
                    _ = v50
                }
            case 1:
                v4 -= v4
            }
            v51 := GenS1{F0: (true && false), F1: 1}
            var v52 int64 = 1
            var v53 float32 = 3.14159
            _ = fi20
            _ = v51
            _ = v52
            _ = v53
        }
        _ = fi17
    }
    v3 = append(v3, v5)
    for fi54 := 0; fi54 < 15; fi54++ {
        var v55 int8 = 0
        for fi56 := 0; fi56 < 16; fi56++ {
            v1 = v1
            _ = fi56
        }
        if false {
            for fi57 := 0; fi57 < 10; fi57++ {
                v58 := GenS1{F0: false, F1: 106}
                _ = fi57
                _ = v58
            }
            if false {
                for fi59 := 0; fi59 < 8; fi59++ {
                    v1 += (v1 & (v1 ^ (v1 % ((42) | 1))))
                    if true {
                        v6 = 546
                        v5 += v6
                        v5 -= (v5 + 2147483648)
                        v60 := GenS1{F0: (v2 < v2), F1: (v55 << ((255) & 63))}
                        if (-193 == 85) {
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            if (false || true) {
                                fmt.Println(v1)
                                if ((65535 == 2) && true) {
                                    fmt.Println(v1, v6, v5)
                                } else {
                                    var v61 int32 = 42
                                    var v62 uint8 = v2
                                    v63 := GenS1{F0: false, F1: (-(v55))}
                                    var v64 string = ("hello" + "")
                                    if false {
                                        v65 := GenS1{F0: (true || false), F1: -74}
                                        _ = v65
                                    }
                                    _ = v61
                                    _ = v62
                                    _ = v63
                                    _ = v64
                                }
                                switch "foo bar" {
                                case "\n\t":
                                    var v66 int32 = 0
                                    _ = v66
                                case "hello":
                                    v67 := GenS1{F0: (0 == uint16(v2)), F1: v55}
                                    _ = v67
                                case "":
                                    v6 -= (v5 & (v5 + 0))
                                default:
                                    if false {
                                        if (!(false || (true || (!((1.5 / ((((0.1 * ((1e10 + 3.14159) / 1e10)) + -2.25) * 3.14159) - float64((-2.25 * float32(0.1))))) != float64(3.14159)))))) {
                                            v5 += 42
                                            switch fi54 {
                                            case 1:
                                                var v68 uint = 408
                                                _ = v68
                                            default:
                                                switch -38 {
                                                case -1:
                                                    var v69 uint32 = v6
                                                    _ = v69
                                                case 42:
                                                    if len(v3) > 0 {
                                                        v3[0] = v6
                                                    }
                                                case 412:
                                                    v70 := GenS2{F0: 42, F1: "\n\t", F2: 450}
                                                    _ = v70
                                                default:
                                                    var v71 float32 = 1e10
                                                    _ = v71
                                                }
                                            }
                                        } else {
                                            v2 = v2
                                        }
                                        var v72 int64 = 0
                                        _ = v72
                                    }
                                }
                            }
                            v73 := GenS3{F0: (v55 << ((65535) & 63))}
                            _ = v73
                        }
                        switch (fi59 * fi54) {
                        case -2147483648:
                            v74 := GenS1{F0: (!false), F1: v55}
                            _ = v74
                        case 805:
                            v75 := GenS4{F0: 4294967295}
                            _ = v75
                        case -546:
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        }
                        _ = v60
                    } else {
                        v76 := GenS3{F0: (v55 ^ 57)}
                        _ = v76
                    }
                    v4 -= v1
                    var v77 int = int(v6)
                    _ = fi59
                    _ = v77
                }
                for fi78 := 0; fi78 < 4; fi78++ {
                    v55 = (v55 >> ((65535) & 63))
                    switch ("a" + (("a" + "hello") + "")) {
                    case "":
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    case "a":
                        var v79 uint = 740
                        fmt.Println(v79, v55, v55)
                        if false {
                            var v80 int8 = 8
                            _ = v80
                        }
                        _ = v79
                    default:
                        v1 = v1
                    }
                    _ = fi78
                }
            } else {
                var v81 uint32 = (v6 << ((42) & 63))
                _ = v81
            }
            for fi82 := 0; fi82 < 2; fi82++ {
                if true {
                    v83 := GenS3{F0: 117}
                    v84 := GenS5{F0: (fi54 | (fi82 >> ((194) & 63))), F1: v6}
                    if true {
                        v85 := GenS6{F0: fi82, F1: (558 == 693), F2: -120}
                        if true {
                            v1 -= v1
                            v86 := GenS7{F0: 438, F1: -1}
                            _ = v86
                        }
                        _ = v85
                    } else {
                        if (true || false) {
                            if ((782 < 794) && ((192 == -726) && false)) {
                                switch 32768 {
                                case 127:
                                    v3 = append(v3, v5)
                                }
                                v6 += (v6 % (((v5 / ((65535) | 1))) | 1))
                                fmt.Println(fi82, fi54, fi82)
                            } else {
                                v2 += (v2 - 0)
                                fmt.Println(v55)
                                v87 := GenS8{F0: 255, F1: 42}
                                _ = v87
                            }
                        }
                        fmt.Println(fi82, fi54)
                    }
                    _ = v83
                    _ = v84
                } else {
                    v1 = -320
                }
                if true {
                    if (!(v2 >= (v2 + (v2 | 255)))) {
                        v6 = v5
                        var v88 uint8 = 128
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        _ = v88
                    }
                    if len(v3) > 0 {
                        v3[0] = v6
                    }
                    v5 = (^(v6))
                    switch fi54 {
                    case 32768:
                        if ((true && (128 <= 255)) && false) {
                            switch (fi54 - 128) {
                            case 42:
                                switch fi82 {
                                case -1:
                                    _ = len(v3)
                                case 2147483647:
                                    var v89 int32 = 1
                                    _ = v89
                                case 4294967295:
                                    switch 0 {
                                    case 1:
                                        if true {
                                            var v90 int64 = -1
                                            fmt.Println(v5)
                                            fmt.Println(v4, v4, v1)
                                            _ = v90
                                        } else {
                                            fmt.Println(v6, fi54, v1)
                                            if true {
                                                var v91 uint32 = v5
                                                v92 := GenS9{F0: -1, F1: 255, F2: v91}
                                                v5 = v6
                                                _ = v91
                                                _ = v92
                                            } else {
                                                v2 = 127
                                            }
                                        }
                                    case 0:
                                        if true {
                                            v93 := GenS3{F0: v55}
                                            var v94 uint16 = 255
                                            if false {
                                                switch ("a" + "\n\t") {
                                                case "\n\t":
                                                    var v95 int16 = (v4 >> ((18446744073709551615) & 63))
                                                    _ = v95
                                                case "":
                                                    v1 -= (v4 & v1)
                                                case "foo bar":
                                                    var v96 int64 = 9223372036854775807
                                                    _ = v96
                                                default:
                                                    var v97 uint16 = (v94 >> ((65535) & 63))
                                                    _ = v97
                                                }
                                            } else {
                                                if (!true) {
                                                    v6 = v5
                                                }
                                                if true {
                                                    var v98 uint8 = 37
                                                    _ = len(v3)
                                                    _ = v98
                                                }
                                                var v99 int8 = v55
                                                if false {
                                                    v2 += v2
                                                    v2 = 1
                                                    var v100 int8 = v55
                                                    _ = v100
                                                }
                                                _ = v99
                                            }
                                            _ = v93
                                            _ = v94
                                        } else {
                                            var v101 uint16 = 2
                                            var v102 float64 = 0.1
                                            _ = v101
                                            _ = v102
                                        }
                                    case 57:
                                        v4 += (v4 ^ (v1 / ((0) | 1)))
                                    }
                                    switch "" {
                                    case "foo bar":
                                        fmt.Println(v55, v5)
                                    }
                                    fmt.Println(v6)
                                }
                            case -830:
                                v1 += (v1 | 144)
                            case 472:
                                v6 = 4294967295
                                v4 = 0
                                var v103 float32 = (0.0 / float32(float64(float32(float64(-2.25)))))
                                if (!((v6 << ((4294967295) & 63)) > (^(v6)))) {
                                    if true {
                                        var v104 uint8 = 255
                                        v5 -= v5
                                        _ = v104
                                    }
                                    if len(v3) > 0 {
                                        v3[0] = (^(v5))
                                    }
                                } else {
                                    var v105 int32 = -1
                                    var v106 int = (fi82 << ((255) & 63))
                                    v107 := GenS10{F0: fi82, F1: "hello", F2: 42}
                                    v55 = (v55 ^ 127)
                                    _ = v105
                                    _ = v106
                                    _ = v107
                                }
                                _ = v103
                            }
                            var v108 uint32 = (v5 & v6)
                            _ = v108
                        }
                    case 128:
                        fmt.Println(v1, fi82)
                    }
                } else {
                    v109 := GenS1{F0: false, F1: -128}
                    fmt.Println(v2)
                    v110 := GenS10{F0: (fi54 * fi82), F1: "hello", F2: 32767}
                    _ = v109
                    _ = v110
                }
                _ = fi82
            }
        }
        _ = fi54
        _ = v55
    }
    for fi111 := 0; fi111 < 12; fi111++ {
        var v112 uint16 = 65535
        for fi113 := 0; fi113 < 10; fi113++ {
            switch fi111 {
            case -1:
                var v114 string = ""
                _ = v114
            case 4294967295:
                v2 -= 93
            }
            var v115 int16 = -1
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            if (!(0.1 > (-2.25 + 3.14159))) {
                if false {
                    v3 = append(v3, (v5 / (((v6 & (v5 * v6))) | 1)))
                    v116 := GenS3{F0: 0}
                    if true {
                        var v117 int8 = -1
                        v112 += (^(v112))
                        _ = v117
                    } else {
                        _ = len(v3)
                    }
                    _ = v116
                } else {
                    if (255 < 255) {
                        var v118 int32 = 1
                        fmt.Println(v115, v5)
                        _ = v118
                    } else {
                        v115 += v1
                        if true {
                            var v119 uint16 = v112
                            var v120 int16 = -625
                            fmt.Println(v112, v119)
                            _ = v119
                            _ = v120
                        } else {
                            v121 := GenS11{F0: 383, F1: -1}
                            var v122 int16 = (v115 >> ((1) & 63))
                            _ = v121
                            _ = v122
                        }
                        _ = len(v3)
                        switch (("\n\t" + "foo bar") + "") {
                        case "a":
                            v1 = (v4 << ((uint(v112)) & 63))
                        case "hello":
                            v5 += (v6 / ((v6) | 1))
                        }
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    }
                    var v123 uint16 = v112
                    v3 = append(v3, v5)
                    _ = v123
                }
                var v124 int64 = -1
                _ = v124
            } else {
                fmt.Println(fi113, fi111, v115)
                v125 := GenS11{F0: 287, F1: 0}
                _ = v125
            }
            _ = fi113
            _ = v115
        }
        _ = fi111
        _ = v112
    }
    v126 := GenS11{F0: 18446744073709551615, F1: 42}
    v5 = v6
    v6 = (v6 | 928)
    for fi127 := 0; fi127 < 11; fi127++ {
        var v128 float32 = 3.14159
        _ = fi127
        _ = v128
    }
    v5 = v5
    var v129 float32 = (1.5 + -2.25)
    v6 = (v5 % ((v5) | 1))
    _ = len(v3)
    for fi130 := 0; fi130 < 20; fi130++ {
        for fi131 := 0; fi131 < 10; fi131++ {
            var v132 string = ("foo bar" + (("\n\t" + "") + (("a" + ("" + "\n\t")) + "hello")))
            fmt.Println(v6)
            _ = fi131
            _ = v132
        }
        _ = fi130
    }
    fmt.Println("v1:", v1)
    fmt.Println("v129:", v129)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v126
    _ = v129
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
