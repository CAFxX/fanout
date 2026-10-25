package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int; F2 int32 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int16; F1 int8 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 string; F1 int64; F2 uint32 }
type GenS6 struct { F0 int; F1 uint64; F2 string }
type GenS7 struct { F0 uint64; F1 int32; F2 uint }
type GenS8 struct { F0 int; F1 int8; F2 int64 }
type GenS9 struct { F0 string; F1 int8; F2 string }
type GenS10 struct { F0 bool }

func main() {
    var v1 uint64 = 115
    var v2 int64 = -9223372036854775808
    var v3 bool = (-9223372036854775808 <= -75)
    v4 := make(map[uint32]bool)
    var v5 uint8 = 255
    var v6 uint8 = (^(v5))
    fmt.Println(v2)
    switch -9223372036854775808 {
    case 127:
        if false {
            fmt.Println(v1, v1, v1)
            var v7 uint32 = 2147483648
            _ = v7
        } else {
            v3 = (float32(float64(-2.25)) != 1.5)
            v8 := []int{}
            v9 := GenS1{F0: -1, F1: -4, F2: -2147483648}
            _ = v8
            _ = v9
        }
    case -828:
        v6 = (v5 ^ v5)
    case 42:
        var v10 uint16 = 1
        _ = v10
    }
    v2 += -9223372036854775808
    switch 0 {
    case 32767:
        var v11 bool = (0 <= 318)
        _ = v11
    case 315:
        if v3 {
            v3 = (-1 <= -2147483648)
            if v3 {
                v12 := GenS1{F0: 0, F1: 2147483647, F2: -783}
                v13 := GenS1{F0: -64, F1: 9223372036854775807, F2: 0}
                _ = v12
                _ = v13
            }
            var v14 uint8 = v6
            _ = v14
        }
    default:
        if ((!false) || (58 <= 65535)) {
            switch 631 {
            case 4294967295:
                switch "foo bar" {
                case "a":
                    v2 = (v2 ^ (v2 % (((v2 | v2)) | 1)))
                case "":
                    if v3 {
                        for fi15 := 0; fi15 < 7; fi15++ {
                            var v16 uint = 255
                            var v17 int16 = 819
                            for fi18 := 0; fi18 < 18; fi18++ {
                                if v3 {
                                    var v19 uint = v16
                                    var v20 bool = (!((128 == 128) || v3))
                                    _ = v19
                                    _ = v20
                                }
                                if false {
                                    v5 -= 0
                                    var v21 int64 = int64(v5)
                                    v4[194] = (!v3)
                                    _ = v21
                                }
                                v22 := GenS2{F0: v17}
                                _ = fi18
                                _ = v22
                            }
                            _ = fi15
                            _ = v16
                            _ = v17
                        }
                        var v23 int32 = 324
                        _ = v23
                    } else {
                        if (!(1000 != 9)) {
                            v1 -= (v1 << ((288) & 63))
                            v1 = 42
                            v5 = (v5 << ((18446744073709551615) & 63))
                        } else {
                            if (556 < 65535) {
                                _ = len(v4)
                                if v3 {
                                    switch "hello" {
                                    case "":
                                        if (!(!(v1 >= 42))) {
                                            for fi24 := 0; fi24 < 2; fi24++ {
                                                for fi25 := 0; fi25 < 12; fi25++ {
                                                    fmt.Println(v3, v1, fi25)
                                                    if (float64(0.0) > 3.14159) {
                                                        fmt.Println(fi25, v2)
                                                    } else {
                                                        switch (fi24 / (((fi24 * (fi24 ^ fi25))) | 1)) {
                                                        case 0:
                                                            v26 := []int{fi25}
                                                            _ = v26
                                                        case 1:
                                                            if (v1 == 68) {
                                                                if v3 {
                                                                    fmt.Println(v3, v1)
                                                                }
                                                                fmt.Println(v2, v2)
                                                                v27 := GenS1{F0: -19, F1: fi25, F2: 1}
                                                                _ = v27
                                                            } else {
                                                                fmt.Println(v1, v1)
                                                                var v28 string = "a"
                                                                if v3 {
                                                                    v29 := []int{}
                                                                    if ((v6 << ((18446744073709551615) & 63)) < (v6 * (v6 << ((18446744073709551615) & 63)))) {
                                                                        _ = len(v29)
                                                                        v3 = ((fi25 * 0) < fi24)
                                                                    } else {
                                                                        if (v3 || (v3 || v3)) {
                                                                            var v30 int16 = 1000
                                                                            _ = v30
                                                                        } else {
                                                                            v31 := GenS1{F0: 127, F1: fi24, F2: 2147483647}
                                                                            v32 := GenS1{F0: 42, F1: fi24, F2: 0}
                                                                            _ = v31
                                                                            _ = v32
                                                                        }
                                                                    }
                                                                    _ = v29
                                                                } else {
                                                                    if v3 {
                                                                        switch 128 {
                                                                        case 1:
                                                                            var v33 float32 = (1e10 * -2.25)
                                                                            _ = v33
                                                                        default:
                                                                            var v34 float64 = 1e10
                                                                            _ = v34
                                                                        }
                                                                        var v35 uint32 = 90
                                                                        _ = v35
                                                                    } else {
                                                                        if (42 != 49) {
                                                                            if (1 != 2147483648) {
                                                                                if v3 {
                                                                                    var v36 int64 = v2
                                                                                    if (!(65535 != 2)) {
                                                                                        v3 = v3
                                                                                        v36 += v2
                                                                                        var v37 float64 = -2.25
                                                                                        _ = v37
                                                                                    }
                                                                                    _ = v36
                                                                                } else {
                                                                                    v38 := []int{-2147483648, -1, fi24}
                                                                                    v2 = (v2 >> ((65535) & 63))
                                                                                    if v3 {
                                                                                        var v39 int8 = -1
                                                                                        fmt.Println(v1)
                                                                                        _ = v39
                                                                                    }
                                                                                    _ = v38
                                                                                }
                                                                                var v40 int8 = 3
                                                                                _ = v40
                                                                            }
                                                                            var v41 float32 = float32(((-2.25 + 0.1) - 1.5))
                                                                            _ = v41
                                                                        }
                                                                        switch 147 {
                                                                        case 623:
                                                                            v42 := GenS2{F0: 42}
                                                                            _ = v42
                                                                        case -650:
                                                                            var v43 uint64 = (v1 & 42)
                                                                            _ = v43
                                                                        case 128:
                                                                            switch 9223372036854775807 {
                                                                            case 9223372036854775807:
                                                                                v2 = -1
                                                                                v4[4294967295] = false
                                                                                v44 := GenS2{F0: 896}
                                                                                v45 := GenS3{F0: 1, F1: 42}
                                                                                _ = v44
                                                                                _ = v45
                                                                            default:
                                                                                var v46 int = fi25
                                                                                _ = v46
                                                                            }
                                                                        }
                                                                        switch (fi25 * 1) {
                                                                        case 741:
                                                                            if (false && v3) {
                                                                                var v47 float32 = (-2.25 * ((1.5 / -2.25) / (3.14159 - ((((0.1 + 0.1) + 1e10) + ((1e10 - (0.1 + 0.0)) + ((0.0 * (1.5 + 3.14159)) * -2.25))) / 1.5))))
                                                                                if (!v3) {
                                                                                    switch ((("" + "\n\t") + ("a" + v28)) + "foo bar") {
                                                                                    case "hello":
                                                                                        fmt.Println(v2)
                                                                                    case "\n\t":
                                                                                        fmt.Println(v6, fi25, v47)
                                                                                    case "a":
                                                                                        fmt.Println(v2, fi24, v47)
                                                                                    }
                                                                                }
                                                                                _ = v47
                                                                            } else {
                                                                                v48 := []int{127, fi24}
                                                                                var v49 float32 = 0.1
                                                                                _ = v48
                                                                                _ = v49
                                                                            }
                                                                        case -2147483648:
                                                                            v2 += int64(v5)
                                                                        }
                                                                    }
                                                                }
                                                                _ = v28
                                                            }
                                                        case -9223372036854775808:
                                                            if true {
                                                                v5 += (v6 - v5)
                                                                if (!("" == "\n\t")) {
                                                                    v1 = v1
                                                                    switch (^(fi24)) {
                                                                    case 127:
                                                                        var v50 uint16 = 2
                                                                        _ = v50
                                                                    case 9223372036854775807:
                                                                        var v51 int8 = 1
                                                                        _ = v51
                                                                    case -1:
                                                                        var v52 string = "hello"
                                                                        _ = v52
                                                                    default:
                                                                        switch fi24 {
                                                                        case -9223372036854775808:
                                                                            switch fi24 {
                                                                            case -878:
                                                                                v6 = v5
                                                                            case -969:
                                                                                var v53 string = "a"
                                                                                _ = v53
                                                                            case 0:
                                                                                switch 2147483647 {
                                                                                case 0:
                                                                                    v54 := GenS2{F0: -875}
                                                                                    var v55 bool = ("" != ("hello" + ""))
                                                                                    _ = v54
                                                                                    _ = v55
                                                                                }
                                                                            default:
                                                                                if (!(!(1e10 <= float64(1e10)))) {
                                                                                    v1 = (v1 * 336)
                                                                                }
                                                                            }
                                                                        case -756:
                                                                            if (28 <= 1) {
                                                                                fmt.Println(fi24, fi25)
                                                                                switch (fi24 ^ 629) {
                                                                                case 0:
                                                                                    v56 := GenS1{F0: 1, F1: 9223372036854775807, F2: 2147483647}
                                                                                    _ = v56
                                                                                default:
                                                                                    if v3 {
                                                                                        v57 := GenS1{F0: 1, F1: fi25, F2: 2147483647}
                                                                                        _ = len(v4)
                                                                                        v58 := []int{}
                                                                                        _ = v57
                                                                                        _ = v58
                                                                                    } else {
                                                                                        v6 += v6
                                                                                        if false {
                                                                                            switch (^(fi25)) {
                                                                                            case 127:
                                                                                                var v59 bool = (1 > 829)
                                                                                                _ = v59
                                                                                            }
                                                                                        } else {
                                                                                            v60 := GenS4{F0: "\n\t"}
                                                                                            var v61 float64 = 3.14159
                                                                                            v62 := v4[0]
                                                                                            var v63 uint16 = 42
                                                                                            fmt.Println(v62, v63)
                                                                                            switch (fi25 >> ((197) & 63)) {
                                                                                            case -9223372036854775808:
                                                                                                var v64 uint8 = (v5 - v6)
                                                                                                _ = v64
                                                                                            case 339:
                                                                                                fmt.Println(v5)
                                                                                            }
                                                                                            switch fi24 {
                                                                                            case 9223372036854775807:
                                                                                                fmt.Println(fi25, v5)
                                                                                            case 42:
                                                                                                fmt.Println(v63)
                                                                                            case -9223372036854775808:
                                                                                                v65 := GenS2{F0: -1}
                                                                                                _ = v65
                                                                                            default:
                                                                                                v2 -= 200
                                                                                            }
                                                                                            if v62 {
                                                                                                switch fi24 {
                                                                                                case 2147483647:
                                                                                                    var v66 int8 = -57
                                                                                                    _ = v66
                                                                                                case 4294967295:
                                                                                                    fmt.Println(v3)
                                                                                                default:
                                                                                                    fmt.Println(v3)
                                                                                                }
                                                                                                fmt.Println(v62)
                                                                                            }
                                                                                            var v67 int64 = (v2 - (-(v2)))
                                                                                            _ = v60
                                                                                            _ = v61
                                                                                            _ = v62
                                                                                            _ = v63
                                                                                            _ = v67
                                                                                        }
                                                                                        var v68 uint64 = 42
                                                                                        _ = v68
                                                                                    }
                                                                                }
                                                                                v1 = (v1 / (((^(v1))) | 1))
                                                                            }
                                                                        }
                                                                    }
                                                                    fmt.Println(fi25)
                                                                }
                                                                if v3 {
                                                                    v69 := GenS5{F0: "", F1: v2, F2: 65535}
                                                                    _ = v69
                                                                }
                                                            }
                                                        default:
                                                            v6 = v5
                                                        }
                                                        switch (^(fi24)) {
                                                        case -9223372036854775808:
                                                            fmt.Println(v6, fi25)
                                                        }
                                                        v70 := GenS4{F0: "\n\t"}
                                                        _ = v70
                                                    }
                                                    var v71 int64 = v2
                                                    _ = fi25
                                                    _ = v71
                                                }
                                                if ((4294967295 == 2) && v3) {
                                                    var v72 int32 = 42
                                                    delete(v4, 879)
                                                    for fi73 := 0; fi73 < 16; fi73++ {
                                                        v1 += (v1 - (v1 | 42))
                                                        v6 = 52
                                                        v74 := GenS3{F0: -32768, F1: 54}
                                                        _ = fi73
                                                        _ = v74
                                                    }
                                                    _ = v72
                                                } else {
                                                    v5 = (v6 ^ v6)
                                                    v1 += v1
                                                    v2 = v2
                                                }
                                                _ = fi24
                                            }
                                        } else {
                                            var v75 int8 = 1
                                            for fi76 := 0; fi76 < 14; fi76++ {
                                                v77 := GenS4{F0: ("a" + "foo bar")}
                                                fmt.Println(fi76, v3)
                                                _ = fi76
                                                _ = v77
                                            }
                                            for fi78 := 0; fi78 < 17; fi78++ {
                                                if true {
                                                    switch "" {
                                                    case "hello":
                                                        var v79 int16 = 876
                                                        _ = v79
                                                    case "foo bar":
                                                        var v80 int16 = int16(v75)
                                                        _ = v80
                                                    default:
                                                        if (18446744073709551615 >= 423) {
                                                            v81 := []int{(fi78 << ((656) & 63)), (fi78 + fi78)}
                                                            switch (fi78 + fi78) {
                                                            case 0:
                                                                switch fi78 {
                                                                case 32767:
                                                                    if len(v81) > 0 {
                                                                        v81[0] = (fi78 * fi78)
                                                                    }
                                                                case 128:
                                                                    if v3 {
                                                                        var v82 float64 = float64(-2.25)
                                                                        _ = v82
                                                                    }
                                                                case 573:
                                                                    var v83 int = fi78
                                                                    _ = v83
                                                                default:
                                                                    v84 := GenS3{F0: 1, F1: (v75 | (v75 / ((v75) | 1)))}
                                                                    _ = v84
                                                                }
                                                            case 32767:
                                                                v6 = v5
                                                            case 127:
                                                                v85 := GenS2{F0: 42}
                                                                _ = v85
                                                            default:
                                                                var v86 float64 = -2.25
                                                                _ = v86
                                                            }
                                                            v75 -= (v75 % ((42) | 1))
                                                            _ = v81
                                                        }
                                                    }
                                                }
                                                var v87 bool = ((v6 << ((67) & 63)) <= (v6 * v6))
                                                if true {
                                                    if (0 <= -32768) {
                                                        for fi88 := 0; fi88 < 10; fi88++ {
                                                            v89 := GenS6{F0: fi78, F1: v1, F2: "\n\t"}
                                                            switch fi88 {
                                                            case -140:
                                                                v90 := GenS2{F0: 961}
                                                                _ = v90
                                                            case 613:
                                                                var v91 int = 32767
                                                                _ = len(v4)
                                                                _ = v91
                                                            case 128:
                                                                var v92 int16 = -1
                                                                _ = v92
                                                            }
                                                            _ = fi88
                                                            _ = v89
                                                        }
                                                        var v93 uint16 = 42
                                                        _ = v93
                                                    }
                                                    switch "foo bar" {
                                                    case "hello":
                                                        v1 += 0
                                                    }
                                                    for fi94 := 0; fi94 < 6; fi94++ {
                                                        if v87 {
                                                            if v87 {
                                                                if v87 {
                                                                    fmt.Println(v3, v6)
                                                                    v95 := GenS7{F0: v1, F1: -1, F2: 255}
                                                                    _ = v95
                                                                } else {
                                                                    fmt.Println(v75, v3)
                                                                }
                                                                v96 := GenS5{F0: (("foo bar" + "foo bar") + ("" + "hello")), F1: -1, F2: 42}
                                                                _ = v96
                                                            }
                                                            var v97 int = (^(fi94))
                                                            _ = v97
                                                        }
                                                        var v98 uint8 = v5
                                                        v3 = false
                                                        _ = fi94
                                                        _ = v98
                                                    }
                                                } else {
                                                    v99 := v4[278]
                                                    _ = v99
                                                }
                                                _ = fi78
                                                _ = v87
                                            }
                                            _ = v75
                                        }
                                    case "\n\t":
                                        v100 := GenS8{F0: 127, F1: -29, F2: (v2 << ((603) & 63))}
                                        _ = v100
                                    }
                                    if v3 {
                                        if false {
                                            if v3 {
                                                fmt.Println(v1, v5, v1)
                                                _ = len(v4)
                                            }
                                        }
                                    } else {
                                        v101 := GenS4{F0: (("foo bar" + "foo bar") + ("" + (("\n\t" + "hello") + "hello")))}
                                        v3 = (1 >= 0)
                                        v1 = (v1 | v1)
                                        _ = v101
                                    }
                                    for fi102 := 0; fi102 < 15; fi102++ {
                                        var v103 string = "hello"
                                        for fi104 := 0; fi104 < 17; fi104++ {
                                            v1 = 0
                                            v105 := GenS8{F0: fi104, F1: 127, F2: 0}
                                            switch fi104 {
                                            case 32768:
                                                delete(v4, 0)
                                            case 395:
                                                v103 = "\n\t"
                                            }
                                            _ = fi104
                                            _ = v105
                                        }
                                        _ = fi102
                                        _ = v103
                                    }
                                } else {
                                    if (v6 > (^(v5))) {
                                        var v106 int8 = 1
                                        _ = v106
                                    } else {
                                        for fi107 := 0; fi107 < 3; fi107++ {
                                            var v108 uint16 = 42
                                            _ = fi107
                                            _ = v108
                                        }
                                    }
                                    var v109 int8 = -54
                                    _ = v109
                                }
                                v2 = v2
                            }
                        }
                        v3 = true
                        v110 := []int{int(v2), -60}
                        _ = v110
                    }
                case "\n\t":
                    if v3 {
                        v2 = (v2 << ((322) & 63))
                        var v111 int64 = -9223372036854775808
                        if true {
                            var v112 int8 = 127
                            var v113 uint8 = 42
                            v114 := GenS4{F0: (("hello" + "") + "foo bar")}
                            _ = v112
                            _ = v113
                            _ = v114
                        } else {
                            var v115 float32 = (3.14159 + (3.14159 / ((-2.25 - (1e10 * 0.0)) / 1.5)))
                            v2 -= (v2 % ((1) | 1))
                            v116 := GenS9{F0: "a", F1: 127, F2: ""}
                            _ = v115
                            _ = v116
                        }
                        _ = v111
                    } else {
                        v4[4294967295] = false
                    }
                default:
                    if (453 != 0) {
                        var v117 int64 = 0
                        v117 = v117
                        _ = v117
                    }
                }
            default:
                var v118 int16 = 32767
                _ = v118
            }
            if (!(v3 || ((v5 ^ 49) != 146))) {
                v4[uint32(v5)] = v3
                if (!(!((!false) && true))) {
                    for fi119 := 0; fi119 < 13; fi119++ {
                        var v120 uint8 = v6
                        v3 = ((1000 <= -1) && true)
                        _ = fi119
                        _ = v120
                    }
                } else {
                    v1 -= 4294967295
                }
                switch -1 {
                case -926:
                    var v121 uint16 = 0
                    _ = v121
                }
            }
        }
    }
    for fi122 := 0; fi122 < 18; fi122++ {
        v123 := GenS2{F0: -939}
        v124 := v4[878]
        _ = fi122
        _ = v123
        _ = v124
    }
    for fi125 := 0; fi125 < 10; fi125++ {
        v126 := GenS10{F0: false}
        switch "hello" {
        case "\n\t":
            var v127 int = fi125
            _ = v127
        default:
            if v3 {
                if v3 {
                    var v128 int64 = (v2 >> ((4294967295) & 63))
                    if (2147483648 == 2) {
                        v129 := v4[0]
                        _ = v129
                    }
                    var v130 int = (fi125 + fi125)
                    _ = v128
                    _ = v130
                } else {
                    v131 := GenS5{F0: "\n\t", F1: 841, F2: 2147483647}
                    if v3 {
                        switch "foo bar" {
                        case "foo bar":
                            v132 := GenS2{F0: 32767}
                            _ = v132
                        }
                        v3 = (v3 || v3)
                    }
                    _ = v131
                }
            } else {
                v133 := GenS7{F0: v1, F1: -120, F2: 255}
                if true {
                    fmt.Println(v2, fi125)
                    v2 = (v2 | (v2 << ((289) & 63)))
                    v134 := []int{int(v1), int(v1), (fi125 * (fi125 ^ (fi125 / ((fi125) | 1))))}
                    _ = v134
                }
                if v3 {
                    if (!true) {
                        var v135 int8 = -128
                        var v136 uint64 = v1
                        _ = v135
                        _ = v136
                    } else {
                        for fi137 := 0; fi137 < 11; fi137++ {
                            var v138 uint = 138
                            var v139 int32 = 42
                            _ = fi137
                            _ = v138
                            _ = v139
                        }
                        var v140 int16 = 1000
                        for fi141 := 0; fi141 < 13; fi141++ {
                            switch "\n\t" {
                            case "hello":
                                switch fi125 {
                                case 0:
                                    v142 := GenS4{F0: ""}
                                    _ = v142
                                case 370:
                                    var v143 int8 = -128
                                    _ = v143
                                case 128:
                                    switch ("\n\t" + "foo bar") {
                                    case "foo bar":
                                        v144 := GenS4{F0: ""}
                                        _ = v144
                                    case "hello":
                                        var v145 int64 = (v2 % ((v2) | 1))
                                        _ = v145
                                    }
                                default:
                                    v146 := GenS8{F0: -183, F1: 42, F2: 9223372036854775807}
                                    _ = v146
                                }
                            case "a":
                                fmt.Println(v6, v2)
                            default:
                                delete(v4, 2147483647)
                            }
                            switch fi125 {
                            case 1:
                                var v147 int64 = v2
                                _ = v147
                            case 551:
                                v148 := []int{fi141, (fi141 + (^(fi141)))}
                                _ = v148
                            case -460:
                                var v149 int16 = 0
                                _ = v149
                            default:
                                if false {
                                    if false {
                                        v6 -= uint8(v1)
                                        fmt.Println(v6, v3, v140)
                                    }
                                    var v150 int32 = -846
                                    _ = v150
                                }
                            }
                            if (847 >= 920) {
                                if (465 <= (v140 << ((487) & 63))) {
                                    v5 = uint8(v1)
                                    var v151 uint = 2
                                    var v152 uint32 = 2
                                    _ = v151
                                    _ = v152
                                }
                                v1 = (v1 / (((v1 / ((v1) | 1))) | 1))
                            } else {
                                var v153 string = (((("hello" + ("foo bar" + "a")) + "\n\t") + (("\n\t" + "hello") + "a")) + ("foo bar" + ("a" + "hello")))
                                _ = v153
                            }
                            _ = fi141
                        }
                        _ = v140
                    }
                }
                _ = v133
            }
        }
        _ = fi125
        _ = v126
    }
    v4[371] = true
    var v154 float32 = 3.14159
    var v155 uint = 268
    var v156 int8 = -128
    fmt.Println("v1:", v1)
    fmt.Println("v154:", v154)
    fmt.Println("v155:", v155)
    fmt.Println("v156:", v156)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v154
    _ = v155
    _ = v156
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
