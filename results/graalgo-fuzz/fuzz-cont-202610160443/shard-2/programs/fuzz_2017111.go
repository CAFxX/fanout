package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8 }
type GenS2 struct { F0 int8; F1 int32 }
type GenS3 struct { F0 bool; F1 int16 }
type GenS4 struct { F0 bool; F1 int; F2 string }
type GenS5 struct { F0 uint16; F1 int16 }
type GenS6 struct { F0 int8; F1 uint }
type GenS7 struct { F0 int8; F1 uint16 }
type GenS8 struct { F0 uint64; F1 uint16 }
type GenS9 struct { F0 uint32; F1 int64 }
type GenS10 struct { F0 int32; F1 uint64 }
type GenS11 struct { F0 int32; F1 int32; F2 int16 }
type GenS12 struct { F0 uint8 }
type GenS13 struct { F0 uint8 }
type GenS14 struct { F0 int16 }
type GenS15 struct { F0 uint8 }
type GenS16 struct { F0 uint16; F1 uint }
type GenS17 struct { F0 bool; F1 int }
type GenS18 struct { F0 int; F1 uint32; F2 uint16 }
type GenS19 struct { F0 int64; F1 string }
type GenS20 struct { F0 string; F1 uint64; F2 int64 }
type GenS21 struct { F0 uint }
type GenS22 struct { F0 int; F1 uint32; F2 uint }
type GenS23 struct { F0 uint8; F1 int; F2 uint32 }

func main() {
    var v1 int32 = 0
    v2 := make(map[uint32]int)
    v3 := make(map[int]int)
    v4 := []string{"hello", ("\n\t" + (("hello" + "") + "\n\t")), ("" + "a"), "foo bar"}
    var v5 float64 = float64((0.0 + 0.1))
    if false {
        var v6 uint8 = 0
        v7 := GenS1{F0: (v1 & 947), F1: v6}
        var v8 int32 = -1
        _ = v6
        _ = v7
        _ = v8
    }
    switch ("foo bar" + "a") {
    case "foo bar":
        switch (("a" + ("" + "a")) + "") {
        case "":
            delete(v3, 42)
        case "a":
            var v9 uint64 = 42
            _ = v9
        }
    case "a":
        for fi10 := 0; fi10 < 4; fi10++ {
            v5 += v5
            var v11 int8 = 0
            var v12 uint = 1
            _ = fi10
            _ = v11
            _ = v12
        }
    case "\n\t":
        switch -423 {
        case -733:
            if len(v4) > 0 {
                v4[0] = "a"
            }
        case 976:
            fmt.Println(v1)
        case 127:
            var v13 uint8 = 42
            _ = v13
        }
    }
    for fi14 := 0; fi14 < 10; fi14++ {
        v15 := GenS2{F0: 39, F1: 42}
        v1 = (v1 & 2147483647)
        _ = fi14
        _ = v15
    }
    for fi16 := 0; fi16 < 8; fi16++ {
        for fi17 := 0; fi17 < 16; fi17++ {
            v1 = (v1 | 1)
            v1 += 2147483647
            var v18 int64 = 9223372036854775807
            var v19 int8 = -38
            _ = fi17
            _ = v18
            _ = v19
        }
        _ = fi16
    }
    var v20 uint64 = 204
    fmt.Println(v1, v20)
    v21 := GenS1{F0: -310, F1: 255}
    for fi22 := 0; fi22 < 6; fi22++ {
        fmt.Println(v1, v1, v5)
        switch (fi22 + 1) {
        case -9223372036854775808:
            delete(v2, 113)
        }
        switch "hello" {
        case "\n\t":
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        case "foo bar":
            if ((v20 % (((v20 + (v20 % ((498) | 1)))) | 1)) < v20) {
                for fi23 := 0; fi23 < 10; fi23++ {
                    v20 = v20
                    _ = fi23
                }
            } else {
                switch fi22 {
                case -276:
                    var v24 string = ""
                    _ = v24
                case -1:
                    var v25 uint = 639
                    _ = v25
                case 128:
                    for fi26 := 0; fi26 < 6; fi26++ {
                        if (!true) {
                            if true {
                                v27 := GenS1{F0: (v1 % ((v1) | 1)), F1: 42}
                                _ = v27
                            }
                            if false {
                                var v28 int = (fi26 % ((fi22) | 1))
                                _ = v28
                            } else {
                                var v29 uint64 = 789
                                switch (^(fi22)) {
                                case -688:
                                    var v30 int64 = 42
                                    _ = v30
                                }
                                var v31 uint = 65535
                                _ = v29
                                _ = v31
                            }
                            v32 := GenS3{F0: (42 == 1), F1: 0}
                            _ = v32
                        }
                        v5 = 1.5
                        if (!true) {
                            switch ("hello" + ((((("foo bar" + "hello") + "a") + "") + "hello") + "\n\t")) {
                            case "\n\t":
                                v1 = v1
                            case "":
                                var v33 string = (("foo bar" + (("\n\t" + "") + "\n\t")) + ("\n\t" + (("\n\t" + "foo bar") + "hello")))
                                _ = v33
                            case "foo bar":
                                v5 = v5
                            default:
                                v1 = (v1 & v1)
                            }
                            v34 := GenS4{F0: (!((!true) || false)), F1: (fi26 | (fi22 >> ((2) & 63))), F2: (("" + "a") + "\n\t")}
                            v4 = append(v4, "foo bar")
                            _ = v34
                        } else {
                            v20 += (v20 & v20)
                            v35 := GenS3{F0: (false || false), F1: -1}
                            _ = v35
                        }
                        _ = fi26
                    }
                default:
                    v20 += v20
                }
                for fi36 := 0; fi36 < 13; fi36++ {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    switch ("" + ("\n\t" + "\n\t")) {
                    case "\n\t":
                        v37 := GenS1{F0: (-(v1)), F1: 42}
                        _ = v37
                    case "a":
                        switch (fi36 << ((65535) & 63)) {
                        case 239:
                            switch ("\n\t" + ("" + "\n\t")) {
                            case "a":
                                if (false && false) {
                                    _ = len(v3)
                                }
                                var v38 uint64 = v20
                                var v39 bool = (((3.14159 * 1e10) * float32((float64((1.5 + float32((((v5 + 0.1) / v5) + 1.5)))) * (v5 + 1e10)))) >= 1e10)
                                switch (-(fi36)) {
                                case 32768:
                                    v40 := GenS4{F0: (!false), F1: (fi22 / (((fi22 - 32768)) | 1)), F2: "hello"}
                                    _ = v40
                                case 4294967295:
                                    if (1.5 > ((1e10 - v5) * v5)) {
                                        switch (fi22 >> ((65535) & 63)) {
                                        case -1:
                                            fmt.Println(v20, v38)
                                        case 127:
                                            if false {
                                                switch ("a" + ("\n\t" + (("foo bar" + "\n\t") + "hello"))) {
                                                case "a":
                                                    var v41 uint8 = 255
                                                    _ = v41
                                                case "\n\t":
                                                    var v42 int64 = 74
                                                    _ = v42
                                                case "":
                                                    v43 := GenS5{F0: 130, F1: 0}
                                                    _ = v43
                                                default:
                                                    v38 -= v38
                                                }
                                                v1 = v1
                                                v5 -= (v5 - float64((-2.25 - (((0.0 / 1e10) * 1e10) + ((1e10 + 1e10) - 0.1)))))
                                            } else {
                                                var v44 float32 = float32(3.14159)
                                                if (v1 > v1) {
                                                    v5 += (v5 - float64(float32(0.1)))
                                                    v38 = v38
                                                    if v39 {
                                                        v5 += ((((v5 * v5) * (float64(v44) / -2.25)) - v5) - 1.5)
                                                        switch -9223372036854775808 {
                                                        case 127:
                                                            switch fi36 {
                                                            case 127:
                                                                v45 := GenS1{F0: v1, F1: 255}
                                                                _ = v45
                                                            case 4294967295:
                                                                v46 := GenS6{F0: 100, F1: 255}
                                                                _ = v46
                                                            }
                                                        case 4:
                                                            switch -2147483648 {
                                                            case -1:
                                                                var v47 string = "\n\t"
                                                                _ = v47
                                                            default:
                                                                fmt.Println(v5, v39)
                                                            }
                                                        case 971:
                                                            delete(v3, (fi22 * fi22))
                                                        default:
                                                            _ = len(v2)
                                                        }
                                                    }
                                                } else {
                                                    switch "\n\t" {
                                                    case "":
                                                        if false {
                                                            if v39 {
                                                                fmt.Println(v1)
                                                                var v48 uint64 = (v38 / ((v20) | 1))
                                                                _ = v48
                                                            }
                                                            if (!v39) {
                                                                v49 := GenS4{F0: ((true && (!v39)) && (32767 > 1)), F1: fi36, F2: ("" + "\n\t")}
                                                                _ = v49
                                                            } else {
                                                                if (2 < 1) {
                                                                    var v50 uint = 2
                                                                    fmt.Println(fi22)
                                                                    _ = v50
                                                                } else {
                                                                    v51 := GenS5{F0: uint16(v20), F1: 0}
                                                                    v52 := GenS3{F0: (v38 == v38), F1: 1}
                                                                    if true {
                                                                        var v53 float64 = (v5 - v5)
                                                                        if v39 {
                                                                            v54 := GenS7{F0: 127, F1: 1}
                                                                            v55 := GenS1{F0: v1, F1: 253}
                                                                            v53 += v5
                                                                            _ = v54
                                                                            _ = v55
                                                                        }
                                                                        var v56 uint8 = 1
                                                                        v57 := GenS8{F0: v20, F1: 202}
                                                                        _ = v53
                                                                        _ = v56
                                                                        _ = v57
                                                                    } else {
                                                                        var v58 float32 = v44
                                                                        v59 := GenS7{F0: 42, F1: 1}
                                                                        v60 := GenS9{F0: 65535, F1: 0}
                                                                        var v61 int16 = -32768
                                                                        var v62 float32 = 0.0
                                                                        _ = v58
                                                                        _ = v59
                                                                        _ = v60
                                                                        _ = v61
                                                                        _ = v62
                                                                    }
                                                                    _ = v51
                                                                    _ = v52
                                                                }
                                                                var v63 uint16 = 0
                                                                var v64 uint16 = 32767
                                                                _ = v63
                                                                _ = v64
                                                            }
                                                        } else {
                                                            var v65 bool = (v1 >= -1)
                                                            if false {
                                                                var v66 uint8 = 42
                                                                _ = v66
                                                            } else {
                                                                var v67 uint = 161
                                                                _ = v67
                                                            }
                                                            var v68 uint8 = 127
                                                            _ = v65
                                                            _ = v68
                                                        }
                                                    case "foo bar":
                                                        v20 = 2
                                                    default:
                                                        delete(v3, 4294967295)
                                                    }
                                                    if v39 {
                                                        v20 -= 2
                                                    }
                                                    v20 += (v38 - (v20 & 4294967295))
                                                }
                                                _ = v44
                                            }
                                        default:
                                            v69 := GenS4{F0: v39, F1: (fi36 & (-(fi22))), F2: "\n\t"}
                                            _ = v69
                                        }
                                    } else {
                                        v70 := GenS10{F0: 1, F1: v20}
                                        if v39 {
                                            var v71 uint8 = 110
                                            if (!v39) {
                                                switch "a" {
                                                case "":
                                                    if v39 {
                                                        v5 += 3.14159
                                                        if true {
                                                            if v39 {
                                                                v39 = false
                                                                var v72 bool = false
                                                                v73 := GenS11{F0: v1, F1: -589, F2: -32768}
                                                                delete(v2, 511)
                                                                v74 := GenS4{F0: (!v72), F1: fi36, F2: ("\n\t" + "a")}
                                                                v75 := GenS12{F0: 124}
                                                                v76 := GenS13{F0: v71}
                                                                _ = v72
                                                                _ = v73
                                                                _ = v74
                                                                _ = v75
                                                                _ = v76
                                                            }
                                                            switch fi36 {
                                                            case 127:
                                                                switch (fi36 | -9223372036854775808) {
                                                                case 1:
                                                                    fmt.Println(v39)
                                                                case -2147483648:
                                                                    v77 := GenS14{F0: 42}
                                                                    _ = v77
                                                                default:
                                                                    var v78 string = (("\n\t" + ("" + "hello")) + "a")
                                                                    _ = v78
                                                                }
                                                                v79 := GenS2{F0: -1, F1: v1}
                                                                _ = v79
                                                            case -1:
                                                                if false {
                                                                    v71 += v71
                                                                    v80 := GenS9{F0: 214, F1: 9223372036854775807}
                                                                    _ = v80
                                                                }
                                                            }
                                                            if (!v39) {
                                                                switch 4294967295 {
                                                                case 9223372036854775807:
                                                                    v81 := GenS15{F0: 127}
                                                                    _ = v81
                                                                default:
                                                                    var v82 uint16 = 0
                                                                    _ = v82
                                                                }
                                                                var v83 int8 = 42
                                                                v5 += (v5 / ((((v5 / float64(1.5)) / v5) + (1.5 * v5)) - 0.0))
                                                                _ = v83
                                                            }
                                                        }
                                                    } else {
                                                        if v39 {
                                                            switch "a" {
                                                            case "hello":
                                                                v20 += (v38 % ((v38) | 1))
                                                            case "":
                                                                delete(v2, 2)
                                                            case "foo bar":
                                                                v4 = append(v4, "")
                                                            }
                                                        }
                                                        v5 -= v5
                                                        var v84 uint16 = 32767
                                                        _ = v84
                                                    }
                                                    var v85 int32 = 0
                                                    var v86 int8 = -112
                                                    _ = v85
                                                    _ = v86
                                                default:
                                                    if (("hello" + "") < "hello") {
                                                        _ = len(v4)
                                                        v20 -= v20
                                                        switch (fi22 & -2147483648) {
                                                        case -827:
                                                            if len(v4) >= 2 {
                                                                v4 = v4[1:len(v4)]
                                                            }
                                                        default:
                                                            fmt.Println(v20, v71, v71)
                                                        }
                                                        v87 := GenS14{F0: 32767}
                                                        _ = v87
                                                    }
                                                }
                                            } else {
                                                fmt.Println(fi22, fi22, v20)
                                                v88 := GenS8{F0: (v20 + 660), F1: 255}
                                                var v89 float64 = (v5 + (v5 / 1.5))
                                                _ = v88
                                                _ = v89
                                            }
                                            _ = v71
                                        } else {
                                            if v39 {
                                                fmt.Println(v1, v5)
                                            } else {
                                                if len(v4) > 0 {
                                                    v4[0] = "\n\t"
                                                }
                                                v39 = v39
                                                v90 := GenS12{F0: 2}
                                                _ = v90
                                            }
                                            var v91 int32 = v1
                                            v92 := GenS11{F0: 42, F1: v1, F2: 474}
                                            _ = v91
                                            _ = v92
                                        }
                                        v39 = (v39 || v39)
                                        _ = v70
                                    }
                                default:
                                    fmt.Println(v20)
                                    var v93 uint64 = v20
                                    var v94 string = ""
                                    _ = v93
                                    _ = v94
                                }
                                _ = len(v4)
                                _ = v38
                                _ = v39
                            default:
                                var v95 uint64 = 42
                                _ = v95
                            }
                            if false {
                                if (v5 <= ((float64(0.0) - v5) + v5)) {
                                    if (v5 <= v5) {
                                        fmt.Println(v5)
                                        switch "" {
                                        case "":
                                            v96 := GenS14{F0: 32767}
                                            _ = v96
                                        case "\n\t":
                                            var v97 int = (fi22 | (fi22 ^ -720))
                                            v5 = v5
                                            _ = v97
                                        default:
                                            v98 := v3[fi22]
                                            _ = v98
                                        }
                                        switch ((("" + ("\n\t" + ("" + ""))) + "a") + "") {
                                        case "a":
                                            switch "" {
                                            case "":
                                                fmt.Println(v5)
                                            default:
                                                v1 += 104
                                                v99 := GenS16{F0: 32767, F1: 1}
                                                _ = v99
                                            }
                                            if true {
                                                switch (fi22 | (^(fi22))) {
                                                case 4294967295:
                                                    var v100 int = (fi22 - 586)
                                                    _ = v100
                                                case 128:
                                                    fmt.Println(fi36)
                                                default:
                                                    fmt.Println(fi22)
                                                }
                                            } else {
                                                v1 += v1
                                                switch ("" + ("\n\t" + ("\n\t" + "hello"))) {
                                                case "hello":
                                                    v5 = 1e10
                                                case "a":
                                                    if true {
                                                        v101 := GenS17{F0: ((^(v20)) != (v20 - v20)), F1: (fi22 >> ((42) & 63))}
                                                        switch (fi36 - fi36) {
                                                        case 4294967295:
                                                            fmt.Println(v20)
                                                        default:
                                                            v1 += 0
                                                        }
                                                        fmt.Println(fi36, v5)
                                                        if false {
                                                            v102 := GenS18{F0: 1, F1: 2, F2: 2}
                                                            v1 += v1
                                                            delete(v2, 0)
                                                            _ = v102
                                                        } else {
                                                            v1 -= 42
                                                            switch (fi22 + fi22) {
                                                            case -2147483648:
                                                                v5 = v5
                                                            case 476:
                                                                v4 = append(v4, ("" + "a"))
                                                            case -914:
                                                                var v103 int32 = v1
                                                                _ = v103
                                                            default:
                                                                v5 = v5
                                                            }
                                                            var v104 int = 4294967295
                                                            _ = v104
                                                        }
                                                        _ = v101
                                                    }
                                                default:
                                                    if true {
                                                        v20 += v20
                                                        if true {
                                                            var v105 string = "\n\t"
                                                            if (!true) {
                                                                var v106 bool = (v105 >= "a")
                                                                fmt.Println(v1)
                                                                _ = v106
                                                            }
                                                            fmt.Println(fi22, v105, fi22)
                                                            _ = v105
                                                        }
                                                        if (-9223372036854775808 < -481) {
                                                            v107 := GenS12{F0: 1}
                                                            var v108 int = (fi36 % (((fi36 / ((32767) | 1))) | 1))
                                                            var v109 int = 649
                                                            _ = v107
                                                            _ = v108
                                                            _ = v109
                                                        } else {
                                                            var v110 string = ""
                                                            switch "foo bar" {
                                                            case "\n\t":
                                                                v111 := GenS10{F0: v1, F1: v20}
                                                                _ = v111
                                                            case "a":
                                                                var v112 uint = 1
                                                                _ = v112
                                                            }
                                                            _ = v110
                                                        }
                                                    }
                                                }
                                            }
                                            if false {
                                                v20 = (v20 >> ((2) & 63))
                                                fmt.Println(v20, v20)
                                                v113 := GenS14{F0: -726}
                                                _ = v113
                                            }
                                        default:
                                            var v114 float32 = (1.5 * -2.25)
                                            var v115 int = (^(fi36))
                                            _ = v114
                                            _ = v115
                                        }
                                    }
                                    fmt.Println(fi22, v1)
                                    v116 := GenS14{F0: -214}
                                    _ = v116
                                } else {
                                    var v117 int64 = -1
                                    _ = v117
                                }
                                v5 = ((v5 / 1.0) * v5)
                                var v118 uint8 = 0
                                if len(v4) > 0 {
                                    v4[0] = ("a" + ("" + ("a" + "foo bar")))
                                }
                                v119 := GenS2{F0: int8(v20), F1: int32(v118)}
                                var v120 string = "foo bar"
                                _ = v118
                                _ = v119
                                _ = v120
                            } else {
                                v20 -= v20
                                switch "\n\t" {
                                case "\n\t":
                                    v121 := GenS19{F0: -1, F1: "\n\t"}
                                    _ = v121
                                }
                            }
                        case 128:
                            _ = len(v3)
                        case 32767:
                            var v122 int16 = 887
                            _ = v122
                        default:
                            switch (fi22 & 127) {
                            case 2147483647:
                                if false {
                                    switch "foo bar" {
                                    case "foo bar":
                                        v123 := GenS20{F0: "\n\t", F1: (v20 | v20), F2: 0}
                                        _ = v123
                                    case "":
                                        v1 += -1
                                    default:
                                        if (!true) {
                                            var v124 int32 = (v1 * (v1 + -945))
                                            switch 0 {
                                            case 0:
                                                var v125 uint32 = 2147483647
                                                _ = v125
                                            case -1:
                                                v124 = (v1 ^ v124)
                                            case -9223372036854775808:
                                                v126 := v3[fi22]
                                                _ = v126
                                            }
                                            switch "a" {
                                            case "\n\t":
                                                var v127 uint8 = 1
                                                _ = v127
                                            case "foo bar":
                                                if len(v4) >= 2 {
                                                    v4 = v4[1:len(v4)]
                                                }
                                            case "a":
                                                switch 9223372036854775807 {
                                                case 127:
                                                    if true {
                                                        v128 := GenS21{F0: 400}
                                                        _ = v128
                                                    } else {
                                                        delete(v3, -9223372036854775808)
                                                        if (74 < 128) {
                                                            switch "a" {
                                                            case "hello":
                                                                switch ("hello" + ("foo bar" + "\n\t")) {
                                                                case "hello":
                                                                    v20 = (v20 | v20)
                                                                }
                                                            case "\n\t":
                                                                v129 := GenS7{F0: 100, F1: 32767}
                                                                _ = v129
                                                            default:
                                                                v130 := v2[65535]
                                                                _ = v130
                                                            }
                                                            var v131 float32 = float32(v5)
                                                            _ = v131
                                                        }
                                                    }
                                                case -641:
                                                    v124 = v124
                                                case -2147483648:
                                                    v1 = (v1 & 2147483647)
                                                default:
                                                    v20 = (v20 << ((2) & 63))
                                                    v20 = (v20 / ((v20) | 1))
                                                }
                                            }
                                            _ = v124
                                        }
                                    }
                                    switch (fi22 % ((int(v1)) | 1)) {
                                    case 126:
                                        _ = len(v2)
                                    case 32768:
                                        var v132 uint = 42
                                        v133 := GenS2{F0: 42, F1: v1}
                                        _ = v132
                                        _ = v133
                                    case 1:
                                        v134 := GenS6{F0: 109, F1: 255}
                                        _ = v134
                                    }
                                } else {
                                    v135 := GenS14{F0: 42}
                                    fmt.Println(fi22, v1)
                                    v5 -= ((v5 * 0.1) * 0.1)
                                    _ = v135
                                }
                            }
                            switch -2147483648 {
                            case -1:
                                v20 = (v20 >> ((255) & 63))
                            case 128:
                                v136 := GenS12{F0: 34}
                                v20 += (v20 % ((v20) | 1))
                                if true {
                                    if false {
                                        v137 := GenS9{F0: 65535, F1: -9223372036854775808}
                                        var v138 int64 = 9223372036854775807
                                        _ = v137
                                        _ = v138
                                    } else {
                                        var v139 int64 = -9223372036854775808
                                        if (((v139 >> ((2) & 63)) == (^(v139))) && ((1e10 * (0.1 / 1.0)) > (float32(1.5) - -2.25))) {
                                            var v140 uint64 = v20
                                            _ = v140
                                        }
                                        if len(v4) > 0 {
                                            v4[0] = "foo bar"
                                        }
                                        _ = v139
                                    }
                                    if false {
                                        var v141 int = (fi36 | (-(fi36)))
                                        if true {
                                            v5 += float64(((0.1 * 1e10) + 0.1))
                                        }
                                        _ = v141
                                    } else {
                                        _ = len(v3)
                                        switch "" {
                                        case "\n\t":
                                            v1 = (-(v1))
                                            var v142 uint = 0
                                            _ = v142
                                        case "":
                                            v143 := GenS12{F0: 81}
                                            _ = v143
                                        default:
                                            if true {
                                                v5 = 0.0
                                                if (false && (-128 >= -40)) {
                                                    if len(v4) > 0 {
                                                        v4[0] = (("a" + "\n\t") + "hello")
                                                    }
                                                    v144 := GenS3{F0: (!true), F1: -1}
                                                    _ = v144
                                                } else {
                                                    if (0 < 0) {
                                                        if len(v4) >= 2 {
                                                            v4 = v4[1:len(v4)]
                                                        }
                                                        var v145 uint32 = 390
                                                        _ = v145
                                                    }
                                                    switch "\n\t" {
                                                    case "a":
                                                        fmt.Println(fi22, v5)
                                                    }
                                                }
                                            } else {
                                                switch "hello" {
                                                case "\n\t":
                                                    var v146 float32 = 1e10
                                                    _ = v146
                                                case "hello":
                                                    var v147 float32 = (float32(v5) / 3.14159)
                                                    _ = v147
                                                }
                                                v148 := GenS22{F0: 127, F1: 60, F2: 65535}
                                                if false {
                                                    v20 -= v20
                                                    switch "hello" {
                                                    case "\n\t":
                                                        v149 := GenS15{F0: 128}
                                                        v150 := GenS6{F0: 0, F1: 4294967295}
                                                        _ = v149
                                                        _ = v150
                                                    case "":
                                                        v5 -= ((float64(0.0) / (v5 * float64(1.5))) / (v5 - (float64(3.14159) + float64((((((-2.25 / 3.14159) + 1.5) * 0.1) - 0.0) / (3.14159 + 1e10))))))
                                                    default:
                                                        fmt.Println(fi22)
                                                    }
                                                } else {
                                                    if ((47 != 0) && false) {
                                                        if false {
                                                            v5 = (-2.25 - ((float64(0.1) - v5) * v5))
                                                        } else {
                                                            switch fi22 {
                                                            case -344:
                                                                if (false || true) {
                                                                    v5 = 1e10
                                                                }
                                                            case 513:
                                                                v5 = float64(0.0)
                                                            default:
                                                                v1 += (v1 - (-(v1)))
                                                            }
                                                            v5 = v5
                                                            fmt.Println(v5)
                                                        }
                                                    }
                                                }
                                                if false {
                                                    var v151 int = (fi22 % ((fi36) | 1))
                                                    if true {
                                                        v20 -= 628
                                                        var v152 string = ("\n\t" + ("a" + "foo bar"))
                                                        v151 -= 2147483647
                                                        _ = v152
                                                    } else {
                                                        var v153 uint8 = 2
                                                        _ = v153
                                                    }
                                                    v20 += v20
                                                    v5 = v5
                                                    _ = v151
                                                }
                                                _ = v148
                                            }
                                        }
                                    }
                                }
                                _ = v136
                            default:
                                v1 += 325
                            }
                            fmt.Println(v20, v20)
                        }
                    case "":
                        v1 -= 1
                    }
                    switch 641 {
                    case -798:
                        v1 = v1
                    case 701:
                        var v154 float32 = -2.25
                        _ = v154
                    default:
                        if len(v4) > 0 {
                            v4[0] = "foo bar"
                        }
                    }
                    if false {
                        var v155 int16 = 761
                        switch "\n\t" {
                        case "":
                            var v156 uint32 = 42
                            _ = v156
                        case "hello":
                            v4 = append(v4, "a")
                        case "foo bar":
                            switch -409 {
                            case -1:
                                v157 := GenS1{F0: (v1 * (v1 << ((255) & 63))), F1: uint8(v20)}
                                _ = v157
                            case -2147483648:
                                fmt.Println(v20, v155)
                            }
                        }
                        if ((1 > 0) && (((v20 % (((v20 ^ (v20 + (v20 | v20)))) | 1)) > v20) || (2 < 65535))) {
                            if (!(905 < 732)) {
                                v158 := GenS16{F0: 255, F1: 1}
                                _ = v158
                            } else {
                                fmt.Println(v20)
                            }
                            v20 -= v20
                        }
                        _ = v155
                    } else {
                        if true {
                            fmt.Println(fi36, fi36, v20)
                        } else {
                            var v159 int8 = 127
                            v160 := v3[fi22]
                            _ = v159
                            _ = v160
                        }
                    }
                    if (true && ((((-2.25 * ((((-2.25 * (-2.25 * (0.0 * (1.5 + 3.14159)))) + (-2.25 - (((0.1 + float32(float64(((0.0 * 1e10) + (1e10 / (0.1 - float32(v5))))))) * ((-2.25 + (1.5 + ((((3.14159 / 3.14159) + 3.14159) - (1e10 + ((0.1 + 3.14159) + 0.1))) * (((0.1 / (3.14159 - 1.5)) / 1.0) * 0.1)))) / (-2.25 + float32(v5)))) - 1e10))) + -2.25) + float32(0.1))) + 1e10) != 1e10) && true)) {
                        delete(v3, (fi36 & (fi22 - (fi22 + (fi36 / ((fi36) | 1))))))
                        var v161 int = 4294967295
                        var v162 bool = (1000 != 1)
                        _ = v161
                        _ = v162
                    }
                    _ = fi36
                }
                switch -508 {
                case -1:
                    if (!(true || false)) {
                        var v163 int8 = 100
                        if (!true) {
                            v1 = (v1 << ((895) & 63))
                            var v164 int = fi22
                            v165 := GenS23{F0: 0, F1: fi22, F2: 1}
                            _ = v164
                            _ = v165
                        } else {
                            if len(v4) > 0 {
                                v4[0] = "a"
                            }
                            v1 = (v1 - v1)
                        }
                        _ = len(v3)
                        _ = v163
                    }
                default:
                    fmt.Println(v1, v1)
                }
            }
        case "a":
            var v166 uint32 = 2147483647
            _ = v166
        default:
            v167 := v3[(fi22 - 4294967295)]
            _ = v167
        }
        _ = fi22
    }
    fmt.Println(v20, v20, v5)
    switch 127 {
    case 4294967295:
        for fi168 := 0; fi168 < 2; fi168++ {
            fmt.Println(v20, v1, v20)
            _ = fi168
        }
    }
    v20 = 643
    v20 = (v20 ^ 1)
    switch "\n\t" {
    case "a":
        var v169 string = "\n\t"
        _ = v169
    case "":
        switch ("" + (("hello" + ("foo bar" + "")) + "hello")) {
        case "a":
            v20 -= 4294967295
        }
    case "\n\t":
        v170 := GenS3{F0: true, F1: -32768}
        _ = v170
    default:
        v171 := GenS23{F0: 128, F1: 4294967295, F2: 1}
        _ = v171
    }
    for fi172 := 0; fi172 < 14; fi172++ {
        var v173 uint16 = 255
        switch "\n\t" {
        case "\n\t":
            var v174 int64 = 899
            _ = v174
        case "":
            fmt.Println(v1)
        default:
            if (4294967295 != (^(fi172))) {
                var v175 float32 = 0.1
                v176 := v2[0]
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                _ = v175
                _ = v176
            }
        }
        _ = fi172
        _ = v173
    }
    v177 := GenS2{F0: 127, F1: (v1 << ((255) & 63))}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v177
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
