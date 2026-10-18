package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint; F2 string }
type GenS2 struct { F0 int8; F1 uint }
type GenS3 struct { F0 int16; F1 int64 }
type GenS4 struct { F0 int; F1 bool; F2 uint64 }
type GenS5 struct { F0 int32; F1 int32 }
type GenS6 struct { F0 int64; F1 string }
type GenS7 struct { F0 string }
type GenS8 struct { F0 int8 }
type GenS9 struct { F0 uint32 }
type GenS10 struct { F0 int64 }
type GenS11 struct { F0 uint64; F1 int64; F2 int }
type GenS12 struct { F0 int8; F1 int64; F2 uint }

func main() {
    var v1 uint32 = 2147483648
    v2 := make(map[uint32]string)
    var v3 int16 = 32767
    for fi4 := 0; fi4 < 12; fi4++ {
        var v5 string = ""
        fmt.Println(fi4, v3, fi4)
        v6 := GenS1{F0: v1, F1: 65535, F2: (v5 + "a")}
        _ = fi4
        _ = v5
        _ = v6
    }
    for fi7 := 0; fi7 < 3; fi7++ {
        var v8 uint32 = (v1 + v1)
        v9 := GenS1{F0: v8, F1: uint(v1), F2: "a"}
        _ = fi7
        _ = v8
        _ = v9
    }
    v10 := GenS1{F0: 1, F1: 898, F2: "foo bar"}
    v1 += v1
    v11 := GenS2{F0: -30, F1: 1}
    for fi12 := 0; fi12 < 15; fi12++ {
        switch "a" {
        case "foo bar":
            v13 := []int{fi12, fi12, (fi12 + fi12)}
            _ = v13
        case "a":
            if (236 >= 2) {
                var v14 int32 = 42
                if ((("" + "\n\t") + "a") < (("a" + "foo bar") + "a")) {
                    v15 := GenS1{F0: v1, F1: 255, F2: ("foo bar" + (("hello" + "a") + ("a" + "a")))}
                    v16 := []int{fi12, 42}
                    for fi17 := 0; fi17 < 18; fi17++ {
                        switch (("\n\t" + (("a" + ("a" + "hello")) + "")) + "hello") {
                        case "\n\t":
                            if false {
                                var v18 uint = 65535
                                v19 := GenS1{F0: 4294967295, F1: v18, F2: ("\n\t" + "foo bar")}
                                _ = v18
                                _ = v19
                            } else {
                                fmt.Println(v14, v1)
                            }
                        default:
                            if false {
                                if (!false) {
                                    var v20 int8 = 127
                                    v16 = append(v16, fi12)
                                    _ = v20
                                } else {
                                    var v21 float32 = ((0.0 / float32(3.14159)) - (0.1 + (((0.0 / 0.1) + 0.1) + 3.14159)))
                                    _ = v21
                                }
                            }
                        }
                        v1 = 751
                        _ = fi17
                    }
                    _ = v15
                    _ = v16
                }
                _ = v14
            } else {
                v22 := []int{(fi12 >> ((693) & 63)), (^(fi12)), (-(fi12))}
                v1 -= v1
                _ = v22
            }
        case "\n\t":
            if (0.0 < (((0.0 * 0.1) * (float32((float64(-2.25) * 0.1)) - 0.0)) - 1.5)) {
                for fi23 := 0; fi23 < 11; fi23++ {
                    v1 += (^(v1))
                    fmt.Println(v3, fi23, v3)
                    _ = fi23
                }
            }
        }
        switch fi12 {
        case 488:
            if true {
                v24 := GenS3{F0: -691, F1: -9223372036854775808}
                v1 = v1
                v1 = (v1 << ((361) & 63))
                _ = v24
            } else {
                switch (("hello" + (((("\n\t" + "foo bar") + ("foo bar" + ("\n\t" + "hello"))) + "hello") + ("hello" + "\n\t"))) + "a") {
                case "":
                    var v25 int64 = -1
                    _ = v25
                case "a":
                    fmt.Println(v3, v1, fi12)
                default:
                    fmt.Println(v1)
                }
                var v26 float32 = 1.5
                _ = v26
            }
        case 32767:
            v27 := []int{fi12}
            _ = v27
        case 921:
            for fi28 := 0; fi28 < 19; fi28++ {
                if (true || true) {
                    v3 = 1
                } else {
                    v29 := GenS4{F0: 1, F1: (24 != v3), F2: 317}
                    _ = v29
                }
                v3 = (v3 * (v3 >> ((2) & 63)))
                v3 = (v3 % ((v3) | 1))
                v30 := []int{}
                _ = fi28
                _ = v30
            }
        default:
            for fi31 := 0; fi31 < 4; fi31++ {
                delete(v2, 1)
                var v32 int16 = -977
                v33 := []int{(fi12 | (fi31 % ((fi31) | 1)))}
                _ = fi31
                _ = v32
                _ = v33
            }
        }
        _ = fi12
    }
    var v34 uint8 = 127
    fmt.Println(v3, v34, v34)
    var v35 float64 = -2.25
    switch "hello" {
    case "foo bar":
        v3 = v3
    case "a":
        var v36 int64 = -405
        _ = v36
    case "":
        v34 -= v34
    default:
        if true {
            fmt.Println(v34)
            v37 := GenS3{F0: (v3 * (-(v3))), F1: -9223372036854775808}
            v35 = (v35 * float64(-2.25))
            _ = v37
        } else {
            v3 = (v3 + (v3 / ((v3) | 1)))
        }
    }
    if (0.1 >= 1e10) {
        v38 := GenS1{F0: (v1 + v1), F1: 2, F2: ("\n\t" + ("a" + ((("a" + "\n\t") + "\n\t") + "\n\t")))}
        _ = v38
    } else {
        for fi39 := 0; fi39 < 1; fi39++ {
            switch (((("a" + ("a" + "\n\t")) + "hello") + "foo bar") + ("a" + "\n\t")) {
            case "\n\t":
                var v40 int16 = (^(v3))
                _ = v40
            default:
                switch fi39 {
                case -621:
                    switch "" {
                    case "\n\t":
                        if true {
                            v3 -= v3
                            if true {
                                v3 = v3
                                v3 = (-(v3))
                                fmt.Println(fi39)
                            }
                        }
                    case "":
                        var v41 int64 = -9223372036854775808
                        _ = v41
                    case "foo bar":
                        for fi42 := 0; fi42 < 11; fi42++ {
                            var v43 uint8 = 2
                            var v44 uint64 = 2
                            v43 = (v34 / (((v34 % ((255) | 1))) | 1))
                            _ = fi42
                            _ = v43
                            _ = v44
                        }
                    }
                }
            }
            for fi45 := 0; fi45 < 8; fi45++ {
                var v46 int16 = v3
                var v47 string = "foo bar"
                _ = fi45
                _ = v46
                _ = v47
            }
            _ = fi39
        }
        for fi48 := 0; fi48 < 1; fi48++ {
            switch (fi48 % ((-1) | 1)) {
            case -9223372036854775808:
                v34 = v34
            case 194:
                for fi49 := 0; fi49 < 5; fi49++ {
                    switch "foo bar" {
                    case "a":
                        fmt.Println(v3, fi49)
                    case "foo bar":
                        v50 := GenS4{F0: (fi49 * (fi48 / (((-(fi48))) | 1))), F1: ((2147483647 >= 661) || false), F2: 42}
                        _ = v50
                    }
                    _ = fi49
                }
            case 673:
                for fi51 := 0; fi51 < 14; fi51++ {
                    v35 = float64((0.0 + 0.1))
                    v52 := []int{32768, fi48, fi48}
                    _ = fi51
                    _ = v52
                }
            default:
                var v53 uint16 = 0
                _ = v53
            }
            var v54 int = fi48
            switch (((("hello" + "\n\t") + ("hello" + "")) + "\n\t") + ("" + "a")) {
            case "foo bar":
                fmt.Println(v34, v34)
            case "":
                switch v54 {
                case -1:
                    for fi55 := 0; fi55 < 10; fi55++ {
                        v54 += fi55
                        switch v54 {
                        case 32768:
                            var v56 uint32 = v1
                            _ = v56
                        case 1:
                            fmt.Println(v35, v3)
                        case -350:
                            if true {
                                switch "" {
                                case "hello":
                                    var v57 uint32 = v1
                                    _ = v57
                                case "\n\t":
                                    v34 -= uint8(v1)
                                    v58 := GenS2{F0: -128, F1: 1}
                                    _ = v58
                                }
                                fmt.Println(fi48, v54)
                            }
                        }
                        if true {
                            fmt.Println(v54, fi48, fi48)
                        } else {
                            if false {
                                if ((!false) || true) {
                                    v1 -= 927
                                    if false {
                                        v59 := GenS5{F0: -2147483648, F1: 42}
                                        _ = v59
                                    } else {
                                        switch fi55 {
                                        case 32767:
                                            if false {
                                                var v60 uint32 = v1
                                                _ = v60
                                            } else {
                                                fmt.Println(v3, v1, v34)
                                                var v61 string = ("\n\t" + (("foo bar" + "hello") + ((("\n\t" + ("\n\t" + "\n\t")) + "hello") + "a")))
                                                _ = v61
                                            }
                                        case 9223372036854775807:
                                            var v62 float64 = (v35 * v35)
                                            _ = v62
                                        default:
                                            var v63 float32 = (3.14159 - 1.5)
                                            _ = v63
                                        }
                                        if true {
                                            v35 -= v35
                                        } else {
                                            if false {
                                                var v64 int64 = 0
                                                var v65 uint16 = 620
                                                var v66 int16 = -1
                                                _ = v64
                                                _ = v65
                                                _ = v66
                                            } else {
                                                if (v34 < v34) {
                                                    v3 -= v3
                                                }
                                                var v67 string = (((("\n\t" + ("\n\t" + "\n\t")) + "hello") + "foo bar") + "a")
                                                switch v54 {
                                                case 2147483647:
                                                    v68 := GenS6{F0: -1, F1: v67}
                                                    _ = v68
                                                case 407:
                                                    fmt.Println(v35, v1)
                                                    v67 = v67
                                                case 133:
                                                    switch 4294967295 {
                                                    case 1:
                                                        var v69 uint = 2
                                                        _ = v69
                                                    }
                                                default:
                                                    v67 = "\n\t"
                                                }
                                                v70 := GenS5{F0: 2147483647, F1: 895}
                                                var v71 uint64 = 616
                                                _ = v67
                                                _ = v70
                                                _ = v71
                                            }
                                            v1 -= 65535
                                        }
                                        var v72 int16 = -1
                                        v73 := GenS7{F0: ("a" + "")}
                                        v34 = v34
                                        v35 -= 3.14159
                                        var v74 uint32 = 823
                                        _ = v72
                                        _ = v73
                                        _ = v74
                                    }
                                    v75 := GenS6{F0: 1, F1: (("hello" + ("\n\t" + "\n\t")) + "")}
                                    _ = v75
                                }
                                if (!true) {
                                    switch "foo bar" {
                                    case "foo bar":
                                        switch v54 {
                                        case 42:
                                            fmt.Println(v1)
                                        case -9223372036854775808:
                                            switch (v54 - fi48) {
                                            case 989:
                                                v1 = (v1 % ((v1) | 1))
                                            case 2147483647:
                                                _ = len(v2)
                                            case -406:
                                                v35 -= v35
                                            default:
                                                var v76 uint16 = 164
                                                _ = v76
                                            }
                                        case 4294967295:
                                            delete(v2, (v1 | (^(v1))))
                                        }
                                    }
                                    if (float64(-2.25) == 3.14159) {
                                        if (uint32(v34) < 4294967295) {
                                            v2[968] = ("hello" + ("a" + "\n\t"))
                                            switch ("foo bar" + "a") {
                                            case "":
                                                v3 = v3
                                            }
                                            v34 += v34
                                        }
                                        v2[(v1 % (((v1 & v1)) | 1))] = "\n\t"
                                        if true {
                                            var v77 uint16 = 2
                                            v78 := GenS5{F0: -1, F1: 42}
                                            if (65535 <= 877) {
                                                var v79 bool = (false && (v35 < (float64(0.1) + 0.1)))
                                                _ = v79
                                            }
                                            _ = v77
                                            _ = v78
                                        } else {
                                            if (128 > v34) {
                                                v80 := GenS2{F0: -1, F1: 18446744073709551615}
                                                v81 := v2[(v1 >> ((2) & 63))]
                                                v82 := GenS6{F0: -9223372036854775808, F1: "\n\t"}
                                                _ = v80
                                                _ = v81
                                                _ = v82
                                            } else {
                                                switch -2147483648 {
                                                case 4294967295:
                                                    var v83 uint64 = 384
                                                    _ = v83
                                                default:
                                                    v3 = v3
                                                }
                                                var v84 float32 = 1.5
                                                v85 := []int{fi48, v54}
                                                _ = v84
                                                _ = v85
                                            }
                                            v86 := GenS4{F0: fi55, F1: false, F2: uint64(v1)}
                                            if ((!false) && (v1 > (v1 + 149))) {
                                                v35 -= v35
                                                delete(v2, v1)
                                                if false {
                                                    if (!true) {
                                                        var v87 int32 = 795
                                                        switch "" {
                                                        case "a":
                                                            v54 -= (v54 - int(v87))
                                                            if (0 <= -1) {
                                                                v88 := GenS8{F0: -28}
                                                                if (!((4294967295 >= 400) && false)) {
                                                                    v1 = v1
                                                                    _ = len(v2)
                                                                    fmt.Println(v3)
                                                                }
                                                                var v89 int16 = (v3 | v3)
                                                                _ = v88
                                                                _ = v89
                                                            } else {
                                                                fmt.Println(v87)
                                                                v90 := GenS9{F0: v1}
                                                                v91 := []int{(fi55 + (^(fi48)))}
                                                                var v92 bool = (false || ((v35 * v35) == v35))
                                                                var v93 int = -130
                                                                _ = v90
                                                                _ = v91
                                                                _ = v92
                                                                _ = v93
                                                            }
                                                        case "foo bar":
                                                            delete(v2, 302)
                                                        default:
                                                            switch (^(v54)) {
                                                            case 42:
                                                                if ((v3 ^ (^(v3))) != (-(v3))) {
                                                                    if false {
                                                                        v35 = float64(-2.25)
                                                                    } else {
                                                                        fmt.Println(v34)
                                                                    }
                                                                }
                                                            case 1:
                                                                if (!false) {
                                                                    var v94 int16 = (v3 % (((v3 % ((0) | 1))) | 1))
                                                                    v95 := []int{32768}
                                                                    _ = v94
                                                                    _ = v95
                                                                }
                                                                switch (fi55 / ((fi48) | 1)) {
                                                                case -21:
                                                                    v96 := []int{}
                                                                    _ = v96
                                                                case 0:
                                                                    v3 += v3
                                                                case -451:
                                                                    if (!(!(1.5 <= float32(v35)))) {
                                                                        v97 := GenS10{F0: 42}
                                                                        v98 := GenS8{F0: 100}
                                                                        v99 := GenS10{F0: 9223372036854775807}
                                                                        _ = v97
                                                                        _ = v98
                                                                        _ = v99
                                                                    } else {
                                                                        var v100 int32 = v87
                                                                        fmt.Println(v3, v100)
                                                                        var v101 float32 = 1e10
                                                                        _ = v100
                                                                        _ = v101
                                                                    }
                                                                default:
                                                                    var v102 float32 = (3.14159 - 0.1)
                                                                    v35 -= float64(v102)
                                                                    _ = v102
                                                                }
                                                            }
                                                        }
                                                        v103 := GenS5{F0: (^(v87)), F1: -1}
                                                        _ = v103
                                                        _ = v87
                                                    } else {
                                                        var v104 int16 = v3
                                                        _ = v104
                                                    }
                                                    switch fi48 {
                                                    case 104:
                                                        if (false || (768 <= 793)) {
                                                            switch "foo bar" {
                                                            case "":
                                                                var v105 int64 = 42
                                                                _ = v105
                                                            case "hello":
                                                                v35 += v35
                                                            default:
                                                                v106 := []int{32768, 4294967295, (fi48 / ((v54) | 1))}
                                                                _ = v106
                                                            }
                                                        }
                                                    case 32767:
                                                        v54 += v54
                                                    case -1:
                                                        v34 -= 2
                                                    }
                                                    fmt.Println(v34)
                                                } else {
                                                    v54 = (v54 << ((255) & 63))
                                                    var v107 int = fi48
                                                    v108 := GenS5{F0: -262, F1: 226}
                                                    _ = v107
                                                    _ = v108
                                                }
                                                v34 -= 42
                                                var v109 int8 = 100
                                                _ = v109
                                            } else {
                                                var v110 int16 = int16(v1)
                                                switch (-(fi55)) {
                                                case 32768:
                                                    var v111 int16 = v3
                                                    _ = v111
                                                }
                                                if false {
                                                    var v112 uint8 = (v34 >> ((949) & 63))
                                                    var v113 int16 = (^(v3))
                                                    if ("" != ("a" + "a")) {
                                                        v34 -= 255
                                                        if (!true) {
                                                            v110 -= -85
                                                            if true {
                                                                var v114 float32 = 0.0
                                                                _ = v114
                                                            } else {
                                                                fmt.Println(v3, fi55, v113)
                                                                var v115 int8 = 118
                                                                _ = v115
                                                            }
                                                            fmt.Println(v110, fi55, v35)
                                                            var v116 int32 = 0
                                                            v117 := v2[65535]
                                                            _ = v116
                                                            _ = v117
                                                        }
                                                        if (37 <= 255) {
                                                            v118 := []int{}
                                                            var v119 int = (v54 ^ (fi55 / (((-(fi55))) | 1)))
                                                            var v120 float64 = ((v35 * v35) + 0.1)
                                                            _ = v118
                                                            _ = v119
                                                            _ = v120
                                                        } else {
                                                            v121 := GenS6{F0: 702, F1: ""}
                                                            _ = v121
                                                        }
                                                        var v122 bool = (0 >= (^(v1)))
                                                        _ = v122
                                                    }
                                                    _ = v112
                                                    _ = v113
                                                }
                                                v123 := GenS7{F0: ""}
                                                v54 = fi48
                                                _ = v110
                                                _ = v123
                                            }
                                            _ = v86
                                        }
                                    }
                                    v124 := []int{fi55}
                                    _ = v124
                                }
                                v125 := GenS11{F0: 679, F1: 216, F2: fi48}
                                switch "a" {
                                case "":
                                    if false {
                                        delete(v2, 20)
                                        var v126 string = ""
                                        _ = v126
                                    } else {
                                        var v127 int64 = 126
                                        var v128 uint = 4294967295
                                        _ = v127
                                        _ = v128
                                    }
                                default:
                                    v129 := GenS10{F0: 482}
                                    _ = v129
                                }
                                _ = v125
                            } else {
                                v1 = (v1 - v1)
                            }
                            delete(v2, v1)
                            fmt.Println(v34, v54, v35)
                            var v130 int32 = 321
                            v1 -= 0
                            _ = v130
                        }
                        _ = fi55
                    }
                case 4294967295:
                    fmt.Println(v1, v3)
                case 9223372036854775807:
                    var v131 uint64 = 427
                    _ = v131
                default:
                    if false {
                        var v132 string = "hello"
                        _ = v132
                    } else {
                        v133 := []int{(v54 ^ -321), 155}
                        _ = v133
                    }
                }
            }
            _ = fi48
            _ = v54
        }
        switch ("" + "\n\t") {
        case "hello":
            if false {
                v3 = v3
            } else {
                for fi134 := 0; fi134 < 9; fi134++ {
                    var v135 bool = true
                    var v136 string = "a"
                    _ = len(v2)
                    _ = fi134
                    _ = v135
                    _ = v136
                }
                var v137 uint8 = (v34 & v34)
                if (((v1 + (v1 << ((431) & 63))) != 0) || (-1 != int(v1))) {
                    v1 -= v1
                    var v138 float32 = 0.0
                    _ = v138
                } else {
                    var v139 uint = 2
                    _ = v139
                }
                _ = v137
            }
        case "":
            v140 := GenS12{F0: 109, F1: 1, F2: 255}
            _ = v140
        }
    }
    for fi141 := 0; fi141 < 1; fi141++ {
        var v142 uint8 = (v34 - 42)
        v3 -= 32767
        v35 = -2.25
        _ = fi141
        _ = v142
    }
    for fi143 := 0; fi143 < 13; fi143++ {
        var v144 int16 = (v3 | v3)
        _ = fi143
        _ = v144
    }
    if true {
        var v145 int16 = -32768
        v34 = 42
        v34 = v34
        _ = v145
    }
    v146 := []int{743, 32767}
    for fi147 := 0; fi147 < 10; fi147++ {
        if false {
            v3 = (v3 % ((v3) | 1))
        }
        _ = fi147
    }
    var v148 bool = true
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v146:", len(v146), v146)
    fmt.Println("v148:", v148)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    _ = v1
    _ = v10
    _ = v11
    _ = v146
    _ = v148
    _ = v2
    _ = v3
    _ = v34
    _ = v35
}
