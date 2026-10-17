package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint32; F1 string; F2 int16 }
type GenS3 struct { F0 int16; F1 int8; F2 int64 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int64; F1 uint8 }
type GenS6 struct { F0 uint16; F1 uint64 }
type GenS7 struct { F0 int32; F1 string; F2 int16 }
type GenS8 struct { F0 int64; F1 uint16; F2 uint64 }
type GenS9 struct { F0 int64 }
type GenS10 struct { F0 int16 }
type GenS11 struct { F0 int8 }

func main() {
    var v1 uint = 65535
    var v2 uint = (^(v1))
    v3 := make(map[uint32]string)
    v4 := GenS1{F0: 2}
    var v5 int16 = -688
    if (127 <= 0) {
        v1 = v2
        fmt.Println(v2, v2, v2)
    } else {
        if true {
            var v6 uint32 = 689
            _ = v6
        } else {
            fmt.Println(v5, v2, v1)
            v7 := v3[979]
            for fi8 := 0; fi8 < 2; fi8++ {
                var v9 float64 = 3.14159
                _ = fi8
                _ = v9
            }
            _ = v7
        }
        for fi10 := 0; fi10 < 18; fi10++ {
            v11 := GenS1{F0: 0}
            _ = fi10
            _ = v11
        }
        v12 := v3[308]
        _ = v12
    }
    for fi13 := 0; fi13 < 3; fi13++ {
        var v14 int8 = 7
        _ = fi13
        _ = v14
    }
    if false {
        for fi15 := 0; fi15 < 3; fi15++ {
            switch (("a" + "\n\t") + "hello") {
            case "a":
                if (((0.1 / 1.0) + 1.5) == float32(3.14159)) {
                    var v16 string = ""
                    var v17 float32 = float32(-2.25)
                    _ = v16
                    _ = v17
                }
            case "\n\t":
                for fi18 := 0; fi18 < 7; fi18++ {
                    fmt.Println(fi18, v1, fi18)
                    _ = fi18
                }
            case "":
                _ = len(v3)
            }
            for fi19 := 0; fi19 < 4; fi19++ {
                v20 := GenS2{F0: 4294967295, F1: ("foo bar" + "foo bar"), F2: 32767}
                v2 = 495
                _ = fi19
                _ = v20
            }
            _ = fi15
        }
    }
    if false {
        v1 -= (v1 + v1)
        if (!false) {
            for fi21 := 0; fi21 < 4; fi21++ {
                v2 += v2
                v22 := GenS2{F0: 311, F1: ("\n\t" + ""), F2: -177}
                v5 += -662
                _ = fi21
                _ = v22
            }
            for fi23 := 0; fi23 < 18; fi23++ {
                if (!((("hello" + ("hello" + "")) + "") < "a")) {
                    if false {
                        fmt.Println(v5, v5)
                        v3[2147483647] = ("" + "\n\t")
                        var v24 string = "foo bar"
                        _ = v24
                    }
                    if true {
                        switch ("foo bar" + "foo bar") {
                        case "\n\t":
                            if false {
                                switch (fi23 + fi23) {
                                case -661:
                                    v2 = (v2 & v2)
                                case -979:
                                    for fi25 := 0; fi25 < 7; fi25++ {
                                        var v26 int8 = -128
                                        _ = fi25
                                        _ = v26
                                    }
                                }
                                if true {
                                    v27 := []int{(fi23 + (-(fi23))), fi23}
                                    _ = v27
                                }
                            }
                        }
                    }
                    v5 += v5
                }
                v2 -= (^(v1))
                v5 = (v5 / (((v5 | (v5 & (v5 % (((v5 | (v5 % ((0) | 1)))) | 1))))) | 1))
                _ = fi23
            }
            switch ((((("a" + "hello") + ("foo bar" + "hello")) + "\n\t") + (("\n\t" + "") + "\n\t")) + "\n\t") {
            case "":
                var v28 float32 = (((0.0 / 1.5) + (1.5 * (-2.25 / -2.25))) - float32((-2.25 + 3.14159)))
                _ = v28
            case "foo bar":
                for fi29 := 0; fi29 < 3; fi29++ {
                    var v30 uint16 = 42
                    v30 = (v30 % ((v30) | 1))
                    v1 = (^(v2))
                    _ = fi29
                    _ = v30
                }
            }
        } else {
            for fi31 := 0; fi31 < 8; fi31++ {
                var v32 int16 = v5
                if ((-2.25 * float32((0.0 * 0.0))) < (((3.14159 - ((1.5 / -2.25) / -2.25)) * (3.14159 / (3.14159 + 0.0))) * (0.0 / 1e10))) {
                    for fi33 := 0; fi33 < 5; fi33++ {
                        if ((-2.25 - (((((3.14159 + float32(float64(1e10))) * 0.1) - (((((0.1 + ((1.5 - ((float32((3.14159 * (float64((-2.25 / float32(1e10))) - 0.0))) * (3.14159 + (1.5 + 0.0))) / 0.1)) / ((((-2.25 / 1.0) - 3.14159) - 1e10) - -2.25))) - -2.25) / ((-2.25 / -2.25) / 3.14159)) + 1e10) / 3.14159)) * ((float32(-2.25) - 1.5) / 1.0)) + 0.0)) == 1e10) {
                            if true {
                                v5 = -497
                            } else {
                                if true {
                                    if (-27 < 100) {
                                        var v34 int32 = -2147483648
                                        var v35 float32 = 0.1
                                        v36 := GenS2{F0: 1, F1: "foo bar", F2: v5}
                                        v1 -= (v1 % ((1) | 1))
                                        _ = v34
                                        _ = v35
                                        _ = v36
                                    }
                                } else {
                                    v37 := GenS1{F0: 0}
                                    switch ("hello" + "hello") {
                                    case "":
                                        v5 = (v32 ^ 272)
                                        if (466 > 677) {
                                            fmt.Println(v5)
                                            if false {
                                                v38 := v3[2]
                                                if false {
                                                    var v39 int = (fi33 & fi31)
                                                    var v40 uint = v1
                                                    _ = v39
                                                    _ = v40
                                                }
                                                _ = v38
                                            } else {
                                                if true {
                                                    _ = len(v3)
                                                } else {
                                                    v41 := []int{2147483647}
                                                    v42 := GenS3{F0: 927, F1: 127, F2: 42}
                                                    _ = v41
                                                    _ = v42
                                                }
                                                delete(v3, 42)
                                                fmt.Println(v32, v5)
                                            }
                                        }
                                        if true {
                                            v2 = 42
                                            v43 := v3[481]
                                            _ = v43
                                        }
                                    case "a":
                                        v2 -= v2
                                    default:
                                        v2 = v1
                                    }
                                    _ = v37
                                }
                                if false {
                                    v1 = (v1 & v2)
                                    var v44 float32 = 0.1
                                    _ = v44
                                }
                            }
                            v1 -= (v2 * (v1 >> ((18446744073709551615) & 63)))
                            fmt.Println(v1, v32)
                        } else {
                            var v45 uint8 = 127
                            _ = v45
                        }
                        v1 -= (v1 / (((v2 / (((v2 + v2)) | 1))) | 1))
                        if (!false) {
                            fmt.Println(v2, v32, v1)
                            if ((fi33 % (((-(fi31))) | 1)) != fi33) {
                                if true {
                                    switch "\n\t" {
                                    case "\n\t":
                                        if false {
                                            v46 := []int{(fi31 - (fi33 * -9223372036854775808)), fi33}
                                            switch fi33 {
                                            case -658:
                                                v2 -= 18446744073709551615
                                            default:
                                                var v47 int16 = (v5 | 1)
                                                _ = v47
                                            }
                                            delete(v3, uint32(v2))
                                            _ = v46
                                        }
                                    case "":
                                        if (true || true) {
                                            v32 = (v5 / ((-1) | 1))
                                        }
                                    case "foo bar":
                                        v48 := v3[2147483648]
                                        _ = v48
                                    default:
                                        fmt.Println(v2)
                                    }
                                    var v49 uint32 = 2147483648
                                    var v50 float64 = float64(float32(0.0))
                                    var v51 int32 = -2147483648
                                    _ = v49
                                    _ = v50
                                    _ = v51
                                } else {
                                    v5 -= v32
                                    v52 := GenS1{F0: 16}
                                    if true {
                                        switch "\n\t" {
                                        case "\n\t":
                                            if false {
                                                v32 += v32
                                                if (42 < 182) {
                                                    v53 := []int{fi31, fi31}
                                                    var v54 int = (fi33 >> ((v2) & 63))
                                                    _ = v53
                                                    _ = v54
                                                } else {
                                                    v55 := GenS4{F0: -1}
                                                    _ = v55
                                                }
                                            } else {
                                                if false {
                                                    v56 := GenS3{F0: v5, F1: 0, F2: -9223372036854775808}
                                                    v32 = v5
                                                    v5 += v5
                                                    v57 := v3[439]
                                                    _ = v56
                                                    _ = v57
                                                } else {
                                                    v3[2147483647] = ""
                                                    if (0 != -90) {
                                                        v32 = (v5 - -215)
                                                    } else {
                                                        var v58 int16 = v32
                                                        if (!true) {
                                                            fmt.Println(fi33, v32)
                                                            if (454 >= 0) {
                                                                var v59 int32 = int32(v2)
                                                                if ((v1 & 18446744073709551615) != v1) {
                                                                    var v60 uint64 = 15
                                                                    if (v60 < (v60 | v60)) {
                                                                        v3[42] = "foo bar"
                                                                    }
                                                                    _ = v60
                                                                } else {
                                                                    v5 = v32
                                                                    v2 = 542
                                                                }
                                                                if ((^(v2)) >= v2) {
                                                                    var v61 int64 = 453
                                                                    switch "a" {
                                                                    case "a":
                                                                        switch ("hello" + ("hello" + ("a" + "\n\t"))) {
                                                                        case "foo bar":
                                                                            delete(v3, 42)
                                                                        default:
                                                                            var v62 float32 = (0.1 + -2.25)
                                                                            _ = v62
                                                                        }
                                                                    case "":
                                                                        fmt.Println(v32)
                                                                    case "hello":
                                                                        v59 = 2147483647
                                                                    default:
                                                                        fmt.Println(fi31, fi31)
                                                                    }
                                                                    v59 = -578
                                                                    _ = v61
                                                                }
                                                                v58 = int16(v2)
                                                                _ = v59
                                                            } else {
                                                                var v63 uint8 = 118
                                                                switch "\n\t" {
                                                                case "hello":
                                                                    v2 += 255
                                                                default:
                                                                    var v64 int16 = v5
                                                                    v65 := GenS5{F0: 42, F1: (v63 % (((v63 & (v63 / (((v63 + 2)) | 1)))) | 1))}
                                                                    var v66 uint64 = 42
                                                                    _ = v64
                                                                    _ = v65
                                                                    _ = v66
                                                                }
                                                                _ = v63
                                                            }
                                                            v32 = v32
                                                        }
                                                        _ = v58
                                                    }
                                                }
                                                if true {
                                                    v5 = v32
                                                    fmt.Println(fi33, v5, v2)
                                                    var v67 bool = true
                                                    v68 := GenS6{F0: 42, F1: 442}
                                                    v69 := GenS7{F0: 1, F1: ("\n\t" + "\n\t"), F2: 32767}
                                                    var v70 uint8 = 42
                                                    var v71 int32 = 2147483647
                                                    switch "hello" {
                                                    case "a":
                                                        v72 := []int{fi31, -978, (fi31 * fi33)}
                                                        _ = v72
                                                    case "\n\t":
                                                        v73 := v3[65535]
                                                        _ = v73
                                                    default:
                                                        fmt.Println(fi33)
                                                    }
                                                    _ = v67
                                                    _ = v68
                                                    _ = v69
                                                    _ = v70
                                                    _ = v71
                                                }
                                                v74 := GenS2{F0: 65535, F1: ("\n\t" + ""), F2: (v5 >> (((v1 + 2)) & 63))}
                                                _ = v74
                                            }
                                        case "foo bar":
                                            if false {
                                                v1 -= v1
                                            } else {
                                                var v75 int = 275
                                                var v76 string = ("" + "foo bar")
                                                _ = v75
                                                _ = v76
                                            }
                                        case "a":
                                            var v77 int = fi31
                                            _ = v77
                                        }
                                    }
                                    if ("foo bar" == "\n\t") {
                                        v32 += (v32 % ((v32) | 1))
                                        var v78 int = (fi33 ^ -9223372036854775808)
                                        _ = v78
                                    }
                                    v3[369] = "hello"
                                    v79 := GenS4{F0: -1}
                                    switch "\n\t" {
                                    case "foo bar":
                                        if (65535 <= 320) {
                                            switch (fi33 << (((v2 & 42)) & 63)) {
                                            case 42:
                                                var v80 int64 = 9223372036854775807
                                                _ = v80
                                            case 2147483647:
                                                var v81 int16 = v5
                                                _ = v81
                                            }
                                            var v82 float32 = (float32(0.0) + 0.1)
                                            _ = v82
                                        }
                                    }
                                    _ = v52
                                    _ = v79
                                }
                                if true {
                                    if (true || (((127 < -8) && false) && (!true))) {
                                        v5 -= (v32 << ((v2) & 63))
                                        v32 -= (v5 | (v32 % ((0) | 1)))
                                        v3[930] = ""
                                    }
                                } else {
                                    v83 := []int{}
                                    var v84 uint8 = 255
                                    _ = v83
                                    _ = v84
                                }
                                delete(v3, 42)
                                if true {
                                    var v85 uint64 = 4294967295
                                    _ = v85
                                } else {
                                    if true {
                                        var v86 string = "foo bar"
                                        v87 := GenS5{F0: 9223372036854775807, F1: 127}
                                        v3[11] = ("hello" + (("" + v86) + "\n\t"))
                                        _ = len(v3)
                                        _ = v86
                                        _ = v87
                                    }
                                    if false {
                                        v5 -= (v32 << (((v1 << (((v1 | (^(v2)))) & 63))) & 63))
                                        var v88 uint8 = 42
                                        v1 += v2
                                        _ = v88
                                    }
                                }
                                var v89 uint64 = 18446744073709551615
                                v5 += v5
                                _ = v89
                            } else {
                                v1 += (^(v1))
                            }
                        }
                        v32 += (v32 >> ((v2) & 63))
                        _ = fi33
                    }
                    if true {
                        if false {
                            v90 := []int{(^(fi31)), (fi31 ^ (fi31 + fi31)), fi31}
                            _ = v90
                        } else {
                            if (890 < 1) {
                                var v91 uint = v2
                                _ = len(v3)
                                var v92 int32 = -1
                                _ = v91
                                _ = v92
                            } else {
                                v2 -= (v2 / ((190) | 1))
                            }
                            v5 = 0
                            v2 = v1
                        }
                        v93 := GenS8{F0: -9223372036854775808, F1: 1, F2: 18446744073709551615}
                        _ = v93
                    } else {
                        switch -9223372036854775808 {
                        case 513:
                            if (((0.0 + (float32(3.14159) + (1.5 - 0.0))) / (0.1 + 0.0)) < ((3.14159 * (1.5 - ((0.0 / (0.1 + (0.0 - (-2.25 - 3.14159)))) + (1.5 - -2.25)))) * -2.25)) {
                                if true {
                                    v94 := []int{fi31, fi31, fi31, fi31}
                                    v32 += 0
                                    v2 = v2
                                    _ = v94
                                }
                                switch ("a" + "foo bar") {
                                case "hello":
                                    var v95 int8 = -48
                                    _ = v95
                                case "":
                                    v3[0] = ""
                                }
                            }
                        }
                        if ((^(v1)) != 4294967295) {
                            var v96 int32 = -2147483648
                            v2 = v1
                            for fi97 := 0; fi97 < 11; fi97++ {
                                v32 = v5
                                _ = fi97
                            }
                            _ = v96
                        } else {
                            fmt.Println(fi31)
                            switch ("a" + "foo bar") {
                            case "foo bar":
                                var v98 int = (fi31 / (((fi31 & 119)) | 1))
                                _ = v98
                            case "a":
                                switch (("\n\t" + "") + "foo bar") {
                                case "\n\t":
                                    delete(v3, 4294967295)
                                }
                            case "hello":
                                switch fi31 {
                                case -51:
                                    v2 = v1
                                case 4294967295:
                                    v99 := GenS4{F0: -1}
                                    _ = v99
                                default:
                                    switch (("" + "foo bar") + ("" + "hello")) {
                                    case "":
                                        v100 := v3[4294967295]
                                        _ = v100
                                    case "foo bar":
                                        delete(v3, 2147483647)
                                    }
                                }
                            }
                            if true {
                                v5 = v32
                                fmt.Println(v32, v5)
                            } else {
                                v101 := GenS9{F0: -220}
                                if true {
                                    v2 += (v2 * v1)
                                }
                                v32 += v5
                                _ = v101
                            }
                        }
                    }
                }
                v102 := GenS9{F0: 1}
                _ = fi31
                _ = v102
                _ = v32
            }
            for fi103 := 0; fi103 < 5; fi103++ {
                fmt.Println(v5, fi103)
                var v104 int16 = v5
                v3[2] = "a"
                _ = fi103
                _ = v104
            }
            if false {
                v105 := []int{128, -9223372036854775808}
                v106 := GenS7{F0: -2147483648, F1: "", F2: (v5 ^ 1000)}
                switch 1 {
                case 42:
                    var v107 string = ("a" + ("foo bar" + "hello"))
                    _ = v107
                case 0:
                    var v108 string = "foo bar"
                    _ = v108
                case 9223372036854775807:
                    if true {
                        for fi109 := 0; fi109 < 18; fi109++ {
                            fmt.Println(v2)
                            var v110 float64 = 0.1
                            switch 32767 {
                            case 138:
                                var v111 uint64 = 218
                                _ = v111
                            case 9223372036854775807:
                                if true {
                                    var v112 uint = v1
                                    _ = v112
                                } else {
                                    v3[0] = "foo bar"
                                    if (42 == 4294967295) {
                                        v113 := GenS8{F0: -9223372036854775808, F1: 1, F2: 928}
                                        v5 = (v5 >> ((v1) & 63))
                                        if true {
                                            v2 = (v2 ^ (v2 * 0))
                                        }
                                        _ = v113
                                    }
                                    var v114 uint16 = 0
                                    _ = v114
                                }
                            }
                            _ = fi109
                            _ = v110
                        }
                        delete(v3, 65535)
                        switch -9223372036854775808 {
                        case -531:
                            delete(v3, 164)
                        case 904:
                            if (2 != 304) {
                                _ = len(v105)
                            }
                        }
                    }
                }
                _ = v105
                _ = v106
            } else {
                var v115 float64 = ((((1e10 * (1.5 + ((float64(float32((0.1 + ((0.1 * ((0.1 * (0.0 + -2.25)) / 1.0)) + 0.0)))) + (0.1 - (3.14159 + 1e10))) + (1.5 / (-2.25 * float64(float32((((0.0 + -2.25) - ((1e10 + 0.0) - 0.1)) - -2.25)))))))) + (1.5 / 0.1)) - 1e10) / 3.14159)
                for fi116 := 0; fi116 < 17; fi116++ {
                    switch fi116 {
                    case 102:
                        v117 := GenS1{F0: 128}
                        _ = v117
                    case -1:
                        if false {
                            delete(v3, 65535)
                            v5 -= (v5 % ((32767) | 1))
                            v2 = (v2 + 457)
                        }
                    case 32767:
                        v118 := GenS3{F0: v5, F1: 123, F2: -68}
                        _ = v118
                    default:
                        v119 := []int{(^(fi116)), fi116, 882}
                        _ = v119
                    }
                    var v120 int = fi116
                    _ = fi116
                    _ = v120
                }
                _ = v115
            }
        }
        if (!true) {
            var v121 string = ""
            _ = v121
        } else {
            v5 = v5
            switch "foo bar" {
            case "foo bar":
                for fi122 := 0; fi122 < 2; fi122++ {
                    fmt.Println(fi122)
                    _ = fi122
                }
            default:
                var v123 uint8 = 161
                _ = v123
            }
            v124 := GenS4{F0: 0}
            _ = v124
        }
    }
    switch int(v5) {
    case -194:
        var v125 float32 = float32(0.1)
        _ = v125
    default:
        for fi126 := 0; fi126 < 3; fi126++ {
            delete(v3, 340)
            switch (fi126 >> ((v2) & 63)) {
            case -993:
                v1 = v2
            case 4294967295:
                if (float64((0.0 / (0.0 - 0.1))) > (1.5 + 3.14159)) {
                    v2 = v2
                } else {
                    var v127 uint8 = 42
                    _ = v127
                }
            default:
                v128 := GenS10{F0: v5}
                _ = v128
            }
            _ = fi126
        }
    }
    v129 := GenS10{F0: (v5 / ((v5) | 1))}
    if false {
        v1 -= (v1 >> ((2) & 63))
    } else {
        switch -2147483648 {
        case -21:
            var v130 float32 = float32(0.0)
            _ = v130
        default:
            v131 := GenS11{F0: -72}
            _ = v131
        }
        v132 := []int{32767, 4294967295, 231, 0}
        fmt.Println(v5, v1, v2)
        _ = v132
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v129
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
