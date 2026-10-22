package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint }
type GenS2 struct { F0 uint8; F1 uint; F2 int8 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int; F1 int8; F2 uint }
type GenS5 struct { F0 uint32; F1 uint32 }
type GenS6 struct { F0 uint8; F1 uint16; F2 uint16 }
type GenS7 struct { F0 uint16; F1 int8 }
type GenS8 struct { F0 uint32; F1 string; F2 uint16 }
type GenS9 struct { F0 uint8; F1 int8 }
type GenS10 struct { F0 uint; F1 uint8 }
type GenS11 struct { F0 int32; F1 int64 }
type GenS12 struct { F0 uint64; F1 uint32; F2 int }
type GenS13 struct { F0 bool }
type GenS14 struct { F0 uint16 }
func genFn0(p0 uint32) int8 {
    fmt.Println(p0)
    var v4 uint = 65535
    if false {
        v5 := GenS1{F0: v4, F1: v4}
        _ = v5
    }
    if (((v4 + v4) > 2) || (!(0 >= 4294967295))) {
        for fi6 := 0; fi6 < 17; fi6++ {
            var v7 int16 = -171
            var v8 int64 = 9223372036854775807
            switch "foo bar" {
            case "a":
                for fi9 := 0; fi9 < 6; fi9++ {
                    v10 := GenS1{F0: v4, F1: v4}
                    var v11 uint8 = 92
                    var v12 uint64 = 2
                    _ = fi9
                    _ = v10
                    _ = v11
                    _ = v12
                }
            case "":
                var v13 bool = (0 <= 128)
                _ = v13
            case "hello":
                if ((((v8 & (v8 | v8)) != int64(v4)) || (false && (-2.25 != 1.5))) || false) {
                    for fi14 := 0; fi14 < 13; fi14++ {
                        if true {
                            v8 = v8
                            v7 -= 32767
                            v15 := GenS2{F0: 2, F1: (^(v4)), F2: int8(v4)}
                            _ = v15
                        } else {
                            switch (("\n\t" + ("\n\t" + (("a" + "\n\t") + "foo bar"))) + "a") {
                            case "a":
                                if (false && (false && true)) {
                                    v16 := []int{fi6, fi6, fi14}
                                    var v17 int16 = (v7 % ((v7) | 1))
                                    _ = v16
                                    _ = v17
                                }
                            case "\n\t":
                                var v18 uint64 = 42
                                _ = v18
                            case "foo bar":
                                if (("" == "foo bar") && false) {
                                    var v19 uint32 = p0
                                    var v20 int = 1
                                    _ = v19
                                    _ = v20
                                } else {
                                    v7 = (v7 >> ((v4) & 63))
                                }
                                fmt.Println(fi14, v7, fi14)
                                v8 = (v8 ^ v8)
                                switch ("foo bar" + ("a" + "a")) {
                                case "a":
                                    v21 := GenS1{F0: v4, F1: 0}
                                    _ = v21
                                case "hello":
                                    switch "a" {
                                    case "hello":
                                        v8 += (v8 - (v8 << (((v4 % (((^(v4))) | 1))) & 63)))
                                    case "\n\t":
                                        v4 += v4
                                    default:
                                        v8 += v8
                                    }
                                }
                                if true {
                                    fmt.Println(fi14, v8, v7)
                                    var v22 uint8 = 1
                                    v23 := []int{}
                                    if ("\n\t" != "foo bar") {
                                        var v24 int8 = -128
                                        fmt.Println(fi6)
                                        _ = v24
                                    }
                                    _ = v22
                                    _ = v23
                                } else {
                                    if ((!(288 <= p0)) || true) {
                                        v8 = v8
                                        if false {
                                            var v25 uint = v4
                                            v26 := GenS1{F0: 4294967295, F1: v4}
                                            _ = v25
                                            _ = v26
                                        } else {
                                            var v27 uint32 = (p0 | 2147483648)
                                            _ = v27
                                        }
                                    }
                                    if false {
                                        v28 := GenS1{F0: 2, F1: (^(v4))}
                                        _ = v28
                                    } else {
                                        fmt.Println(fi6, fi14, v8)
                                        v4 = v4
                                    }
                                    fmt.Println(fi14, v4, fi6)
                                }
                            }
                        }
                        var v29 float32 = 0.0
                        switch 127 {
                        case 1:
                            v30 := GenS2{F0: 2, F1: (v4 / ((0) | 1)), F2: -128}
                            _ = v30
                        case -282:
                            v31 := GenS3{F0: false}
                            _ = v31
                        default:
                            v8 = v8
                        }
                        _ = fi14
                        _ = v29
                    }
                    switch -56 {
                    case 1:
                        v4 = (v4 % (((v4 >> ((v4) & 63))) | 1))
                    }
                }
            default:
                for fi32 := 0; fi32 < 12; fi32++ {
                    v8 -= (-(v8))
                    switch ((((("hello" + "\n\t") + "a") + "\n\t") + "hello") + "") {
                    case "":
                        fmt.Println(p0, fi6, v7)
                    case "a":
                        v33 := GenS4{F0: fi32, F1: 127, F2: (v4 - v4)}
                        _ = v33
                    }
                    _ = fi32
                }
            }
            _ = fi6
            _ = v7
            _ = v8
        }
        v34 := []int{214}
        var v35 uint64 = 1
        _ = v34
        _ = v35
    }
    var v36 float32 = -2.25
    _ = p0
    _ = v36
    _ = v4
    return 127
}
func genFn1(p0 uint16) int {
    p0 = (^(p0))
    if (-74 != -128) {
        var v37 uint32 = 2
        _ = v37
    } else {
        var v38 int64 = 995
        switch 0 {
        case 128:
            if (!(42 < 4294967295)) {
                v39 := GenS2{F0: 255, F1: 2, F2: 100}
                var v40 uint64 = 4294967295
                _ = v39
                _ = v40
            }
        case 310:
            v41 := []int{9223372036854775807, 2147483647}
            _ = v41
        case 724:
            p0 = p0
        default:
            v38 = (v38 % (((v38 ^ (v38 + (v38 ^ v38)))) | 1))
        }
        _ = v38
    }
    _ = p0
    return 2147483647
}

func main() {
    var v1 int = 32767
    var v2 uint16 = 42
    v3 := []string{"", "hello", "hello"}
    v2 += 32767
    v3 = append(v3, ("\n\t" + "foo bar"))
    v42 := GenS5{F0: 2147483647, F1: 2147483648}
    switch v1 {
    case 128:
        v43 := GenS6{F0: 42, F1: 2, F2: (v2 >> ((18446744073709551615) & 63))}
        _ = v43
    }
    var v44 string = "\n\t"
    for fi45 := 0; fi45 < 2; fi45++ {
        switch ((("a" + ("\n\t" + ("foo bar" + ("" + v44)))) + ("\n\t" + v44)) + v44) {
        case "a":
            v3 = append(v3, "a")
        case "":
            v2 -= v2
        case "hello":
            for fi46 := 0; fi46 < 2; fi46++ {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                v1 = (v1 - (-(fi45)))
                v1 -= (fi46 ^ (fi46 | (fi46 | (fi45 * fi45))))
                var v47 int = (v1 / (((fi45 & 32768)) | 1))
                _ = fi46
                _ = v47
            }
        default:
            var v48 uint = 18446744073709551615
            _ = v48
        }
        for fi49 := 0; fi49 < 7; fi49++ {
            v1 -= genFn1((^(v2)))
            var v50 string = ""
            v51 := GenS5{F0: 42, F1: 65535}
            v1 += fi45
            _ = fi49
            _ = v50
            _ = v51
        }
        _ = fi45
    }
    v1 = genFn1(v2)
    for fi52 := 0; fi52 < 14; fi52++ {
        if false {
            if (!true) {
                if false {
                    for fi53 := 0; fi53 < 18; fi53++ {
                        switch v44 {
                        case "a":
                            var v54 uint32 = 0
                            _ = v54
                        }
                        var v55 int64 = -958
                        v44 = "a"
                        _ = fi53
                        _ = v55
                    }
                    fmt.Println(v1, v2, fi52)
                }
                for fi56 := 0; fi56 < 3; fi56++ {
                    v44 = "foo bar"
                    _ = len(v3)
                    if true {
                        v2 = (^(v2))
                        if len(v3) > 0 {
                            v3[0] = "hello"
                        }
                        if len(v3) > 0 {
                            v3[0] = ((((v44 + v44) + "a") + v44) + v44)
                        }
                    } else {
                        v2 = 2
                        if false {
                            switch fi56 {
                            case 128:
                                var v57 float64 = 1.5
                                _ = v57
                            }
                            var v58 bool = true
                            var v59 int16 = -32768
                            _ = v58
                            _ = v59
                        }
                    }
                    _ = fi56
                }
            } else {
                _ = len(v3)
            }
            v3 = append(v3, v44)
            v3 = append(v3, v44)
        }
        _ = fi52
    }
    var v60 string = ("foo bar" + v44)
    var v61 uint64 = 4294967295
    v62 := GenS7{F0: v2, F1: genFn0(974)}
    _ = len(v3)
    switch 127 {
    case 969:
        switch (v1 % ((v1) | 1)) {
        case 0:
            switch ("\n\t" + v44) {
            case "foo bar":
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            case "a":
                if (!(!(!false))) {
                    if (1 != 1) {
                        fmt.Println(v44, v61)
                        if len(v3) > 0 {
                            v3[0] = (v60 + ((v60 + "") + (v44 + "")))
                        }
                        if (343 >= 1) {
                            fmt.Println(v61, v61, v44)
                            v60 = v60
                            for fi63 := 0; fi63 < 4; fi63++ {
                                var v64 int64 = int64(v2)
                                for fi65 := 0; fi65 < 11; fi65++ {
                                    _ = len(v3)
                                    fmt.Println(v44)
                                    fmt.Println(v61, v44, v44)
                                    var v66 uint8 = uint8(v61)
                                    v3 = append(v3, v44)
                                    var v67 uint8 = 198
                                    var v68 uint = 878
                                    _ = fi65
                                    _ = v66
                                    _ = v67
                                    _ = v68
                                }
                                _ = fi63
                                _ = v64
                            }
                        }
                    }
                    v69 := GenS8{F0: 65535, F1: "hello", F2: v2}
                    _ = v69
                } else {
                    var v70 uint16 = 255
                    _ = v70
                }
            case "\n\t":
                var v71 uint16 = v2
                _ = v71
            }
        case -507:
            v72 := GenS9{F0: 255, F1: 42}
            _ = v72
        default:
            var v73 int32 = 2147483647
            _ = v73
        }
    case 128:
        switch (v1 % (((v1 / ((32768) | 1))) | 1)) {
        case 32767:
            var v74 uint32 = 2147483648
            _ = v74
        default:
            var v75 string = v60
            _ = v75
        }
    case -2147483648:
        if false {
            v44 = v60
        }
    default:
        for fi76 := 0; fi76 < 3; fi76++ {
            switch (-(v1)) {
            case 42:
                for fi77 := 0; fi77 < 4; fi77++ {
                    fmt.Println(fi77, v60, v1)
                    _ = fi77
                }
            case -293:
                for fi78 := 0; fi78 < 17; fi78++ {
                    if false {
                        var v79 uint16 = v2
                        var v80 string = (v44 + v60)
                        _ = v79
                        _ = v80
                    }
                    var v81 uint32 = 4294967295
                    fmt.Println(v61)
                    v81 += v81
                    _ = fi78
                    _ = v81
                }
            default:
                if (!(0 <= 756)) {
                    switch (fi76 ^ v1) {
                    case -651:
                        var v82 uint16 = 0
                        _ = v82
                    case 9223372036854775807:
                        fmt.Println(fi76, v60, fi76)
                    case 2147483647:
                        if false {
                            var v83 int16 = -32768
                            _ = v83
                        } else {
                            if false {
                                if true {
                                    if false {
                                        if (1e10 <= 1e10) {
                                            if false {
                                                v61 = v61
                                                v84 := GenS5{F0: 2147483647, F1: 65535}
                                                _ = v84
                                            }
                                            var v85 uint64 = (v61 + (v61 * v61))
                                            _ = v85
                                        } else {
                                            v86 := GenS10{F0: 2, F1: 0}
                                            _ = v86
                                        }
                                    } else {
                                        for fi87 := 0; fi87 < 4; fi87++ {
                                            v1 = -9223372036854775808
                                            _ = fi87
                                        }
                                        switch "\n\t" {
                                        case "":
                                            for fi88 := 0; fi88 < 4; fi88++ {
                                                if (!(true && (-9223372036854775808 >= -336))) {
                                                    fmt.Println(fi88, v1, v44)
                                                    var v89 int16 = -32768
                                                    var v90 int64 = 1
                                                    if len(v3) > 0 {
                                                        v3[0] = v44
                                                    }
                                                    if false {
                                                        v1 -= 209
                                                    } else {
                                                        switch (((("" + "foo bar") + (v44 + ("hello" + v60))) + v44) + "foo bar") {
                                                        case "\n\t":
                                                            var v91 float64 = 0.0
                                                            _ = v91
                                                        case "hello":
                                                            if true {
                                                                var v92 int = (-(fi76))
                                                                switch (v1 - genFn1((v2 ^ v2))) {
                                                                case 81:
                                                                    v1 = -2147483648
                                                                case 32767:
                                                                    switch 1 {
                                                                    case 1:
                                                                        v3 = append(v3, v44)
                                                                    }
                                                                    switch (v60 + v44) {
                                                                    case "hello":
                                                                        v89 += 32767
                                                                    case "\n\t":
                                                                        if false {
                                                                            v60 = "\n\t"
                                                                            var v93 int32 = -587
                                                                            _ = v93
                                                                        }
                                                                        var v94 string = (v44 + "foo bar")
                                                                        _ = v94
                                                                    case "foo bar":
                                                                        if len(v3) >= 2 {
                                                                            v3 = v3[1:len(v3)]
                                                                        }
                                                                    default:
                                                                        if (!false) {
                                                                            fmt.Println(v44)
                                                                            var v95 uint = 769
                                                                            _ = v95
                                                                        }
                                                                    }
                                                                    switch v60 {
                                                                    case "hello":
                                                                        fmt.Println(fi88, v2, v44)
                                                                    case "\n\t":
                                                                        v96 := GenS11{F0: 911, F1: (v90 >> ((42) & 63))}
                                                                        _ = v96
                                                                    case "foo bar":
                                                                        v97 := GenS12{F0: 799, F1: 0, F2: genFn1(255)}
                                                                        _ = v97
                                                                    default:
                                                                        var v98 uint8 = 128
                                                                        _ = v98
                                                                    }
                                                                    v44 = "a"
                                                                }
                                                                _ = v92
                                                            }
                                                        default:
                                                            v99 := GenS6{F0: 1, F1: (v2 & (v2 & (v2 + (v2 * (^(v2)))))), F2: (v2 >> ((0) & 63))}
                                                            _ = v99
                                                        }
                                                        if ((v1 <= genFn1((v2 ^ (^(v2))))) && true) {
                                                            if len(v3) > 0 {
                                                                v3[0] = v44
                                                            }
                                                            if false {
                                                                if (false || false) {
                                                                    if len(v3) > 0 {
                                                                        v3[0] = ((v60 + v44) + (((v44 + "a") + "hello") + v44))
                                                                    }
                                                                    switch -1 {
                                                                    case 4294967295:
                                                                        v100 := GenS8{F0: 65535, F1: ((((v60 + v44) + (v60 + v60)) + v60) + (v60 + (v60 + v60))), F2: 4}
                                                                        _ = v100
                                                                    }
                                                                    _ = len(v3)
                                                                } else {
                                                                    v101 := GenS10{F0: 2, F1: 229}
                                                                    fmt.Println(v1)
                                                                    _ = len(v3)
                                                                    var v102 int8 = 127
                                                                    var v103 uint64 = v61
                                                                    var v104 bool = (916 == 42)
                                                                    if v104 {
                                                                        v105 := GenS1{F0: 18446744073709551615, F1: 2}
                                                                        _ = v105
                                                                    } else {
                                                                        v1 -= fi76
                                                                        v106 := GenS6{F0: 83, F1: v2, F2: (v2 + 0)}
                                                                        _ = v106
                                                                    }
                                                                    _ = v101
                                                                    _ = v102
                                                                    _ = v103
                                                                    _ = v104
                                                                }
                                                                _ = len(v3)
                                                                _ = len(v3)
                                                                _ = len(v3)
                                                            }
                                                        }
                                                        v107 := GenS9{F0: uint8(v61), F1: 42}
                                                        _ = v107
                                                    }
                                                    _ = v89
                                                    _ = v90
                                                }
                                                v60 = v44
                                                _ = fi88
                                            }
                                        case "a":
                                            v1 = -2147483648
                                        default:
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                        }
                                    }
                                    var v108 float32 = ((3.14159 + 3.14159) * ((((((1.5 - (1e10 - 0.0)) / 3.14159) * (float32(1e10) * 0.1)) - 3.14159) * (((-2.25 + 1e10) * ((-2.25 * ((0.1 / (-2.25 + (3.14159 / 3.14159))) + 1.5)) / 1.0)) * (1.5 - 0.1))) - 1.5))
                                    v1 = fi76
                                    _ = v108
                                } else {
                                    if (!true) {
                                        for fi109 := 0; fi109 < 10; fi109++ {
                                            var v110 float32 = 0.1
                                            _ = fi109
                                            _ = v110
                                        }
                                        v3 = append(v3, ("foo bar" + v44))
                                        for fi111 := 0; fi111 < 20; fi111++ {
                                            v44 = (v60 + v60)
                                            var v112 uint16 = (v2 - v2)
                                            _ = fi111
                                            _ = v112
                                        }
                                    } else {
                                        for fi113 := 0; fi113 < 12; fi113++ {
                                            v44 = ("foo bar" + ("hello" + "\n\t"))
                                            if true {
                                                v114 := GenS6{F0: 2, F1: 65535, F2: (v2 ^ v2)}
                                                _ = v114
                                            } else {
                                                if (9223372036854775807 < -1) {
                                                    v115 := GenS7{F0: v2, F1: 42}
                                                    if true {
                                                        if false {
                                                            if true {
                                                                var v116 int8 = 68
                                                                _ = v116
                                                            }
                                                        }
                                                    }
                                                    var v117 int16 = 0
                                                    var v118 int8 = 100
                                                    fmt.Println(v44, fi113, v1)
                                                    _ = v115
                                                    _ = v117
                                                    _ = v118
                                                } else {
                                                    var v119 uint = 42
                                                    _ = v119
                                                }
                                                v2 -= (v2 >> ((412) & 63))
                                            }
                                            fmt.Println(v60, v61, v61)
                                            _ = fi113
                                        }
                                    }
                                }
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                            }
                            if false {
                                for fi120 := 0; fi120 < 14; fi120++ {
                                    var v121 int = (^(fi120))
                                    v2 -= v2
                                    _ = fi120
                                    _ = v121
                                }
                                for fi122 := 0; fi122 < 13; fi122++ {
                                    v123 := GenS1{F0: 346, F1: 96}
                                    v124 := GenS7{F0: (v2 & v2), F1: 100}
                                    _ = fi122
                                    _ = v123
                                    _ = v124
                                }
                                var v125 uint8 = 1
                                _ = v125
                            } else {
                                fmt.Println(v1, v61, v61)
                            }
                            for fi126 := 0; fi126 < 12; fi126++ {
                                v127 := GenS13{F0: true}
                                fmt.Println(v61, v44)
                                if len(v3) > 0 {
                                    v3[0] = ("hello" + v60)
                                }
                                _ = fi126
                                _ = v127
                            }
                        }
                    default:
                        v128 := GenS11{F0: -4, F1: 1}
                        _ = v128
                    }
                    _ = len(v3)
                }
            }
            _ = fi76
        }
    }
    switch v1 {
    case -1:
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    case -905:
        if (244 != 30) {
            var v129 int32 = 1
            switch ("a" + v60) {
            case "foo bar":
                v129 -= (v129 + v129)
            case "a":
                for fi130 := 0; fi130 < 20; fi130++ {
                    switch (v1 / (((v1 * v1)) | 1)) {
                    case 32768:
                        if (false && ((false && (!false)) || (625 < uint32(v61)))) {
                            if true {
                                switch 4294967295 {
                                case 128:
                                    if (v44 <= "foo bar") {
                                        fmt.Println(v60)
                                        if (!false) {
                                            v129 = (^(v129))
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                            var v131 uint64 = v61
                                            _ = v131
                                        }
                                    } else {
                                        if (812 == -9223372036854775808) {
                                            var v132 uint32 = 4294967295
                                            v132 -= (v132 ^ (^(v132)))
                                            _ = v132
                                        } else {
                                            var v133 uint8 = 0
                                            v60 = "foo bar"
                                            _ = v133
                                        }
                                        switch "" {
                                        case "":
                                            for fi134 := 0; fi134 < 15; fi134++ {
                                                v135 := GenS14{F0: (^(v2))}
                                                if len(v3) > 0 {
                                                    v3[0] = "foo bar"
                                                }
                                                var v136 float64 = ((1.5 - 3.14159) + 3.14159)
                                                _ = fi134
                                                _ = v135
                                                _ = v136
                                            }
                                        }
                                        fmt.Println(v2, v60, v60)
                                    }
                                default:
                                    fmt.Println(v60, v44, v129)
                                }
                            } else {
                                v3 = append(v3, v44)
                            }
                        }
                    case 2147483647:
                        if (false && true) {
                            fmt.Println(v44, fi130)
                            var v137 string = (v60 + (v44 + ""))
                            _ = v137
                        } else {
                            v138 := GenS3{F0: false}
                            _ = v138
                        }
                    case 1:
                        v44 = v44
                    }
                    switch "\n\t" {
                    case "a":
                        var v139 int8 = -1
                        _ = v139
                    default:
                        v44 = ((((v60 + ((v44 + v60) + v44)) + "a") + "foo bar") + (v44 + v44))
                    }
                    var v140 int = (fi130 - v1)
                    _ = fi130
                    _ = v140
                }
            case "\n\t":
                if ((!false) || (42 <= -1)) {
                    v61 -= 0
                    var v141 int16 = 32767
                    if false {
                        v2 = v2
                    }
                    _ = v141
                } else {
                    if len(v3) > 0 {
                        v3[0] = "hello"
                    }
                }
            }
            _ = v129
        }
    default:
        var v142 uint = 27
        _ = v142
    }
    v1 += genFn1((v2 | 65535))
    var v143 bool = true
    fmt.Println("v1:", v1)
    fmt.Println("v143:", v143)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v44:", v44)
    fmt.Println("v60:", v60)
    fmt.Println("v61:", v61)
    _ = v1
    _ = v143
    _ = v2
    _ = v3
    _ = v42
    _ = v44
    _ = v60
    _ = v61
    _ = v62
}
