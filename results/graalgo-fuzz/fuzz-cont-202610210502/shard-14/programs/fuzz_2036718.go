package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint; F2 int32 }
type GenS2 struct { F0 int16; F1 uint16 }
type GenS3 struct { F0 uint; F1 uint32; F2 int16 }
type GenS4 struct { F0 int; F1 int16 }
type GenS5 struct { F0 int; F1 uint64; F2 int64 }
type GenS6 struct { F0 uint64; F1 int32; F2 uint64 }
type GenS7 struct { F0 string }
type GenS8 struct { F0 string; F1 int; F2 uint16 }
type GenS9 struct { F0 int16; F1 uint8; F2 uint }
type GenS10 struct { F0 uint32; F1 int8 }
type GenS11 struct { F0 uint16; F1 uint16; F2 int32 }
type GenS12 struct { F0 uint16; F1 int; F2 int8 }
type GenS13 struct { F0 uint8; F1 uint32; F2 int32 }
type GenS14 struct { F0 string }
type GenS15 struct { F0 string; F1 int; F2 int }
type GenS16 struct { F0 bool }
type GenS17 struct { F0 uint; F1 string }
type GenS18 struct { F0 int64; F1 bool; F2 int }
type GenS19 struct { F0 int64; F1 int64; F2 string }
type GenS20 struct { F0 string; F1 uint32 }
type GenS21 struct { F0 uint8; F1 string; F2 uint32 }
func genFn0(p0 int) int16 {
    if (1 < 255) {
        var v7 float32 = 0.1
        _ = v7
    }
    p0 += (p0 + (p0 + 127))
    _ = p0
    return -712
}
func genFn1(p0 int32, p1 uint8, p2 int8) uint16 {
    for fi8 := 0; fi8 < 10; fi8++ {
        var v9 uint = 580
        _ = fi8
        _ = v9
    }
    v10 := GenS1{F0: 255, F1: 38, F2: p0}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return 255
}

func main() {
    v1 := []string{"", ("foo bar" + ("\n\t" + "")), "\n\t"}
    var v2 uint = 65535
    var v3 uint8 = 42
    v4 := make(map[string]bool)
    var v5 int8 = 42
    v6 := make(map[string]bool)
    v11 := GenS2{F0: 1000, F1: genFn1(-81, (v3 ^ (v3 ^ 128)), (^(v5)))}
    var v12 uint64 = 42
    v13 := GenS2{F0: -32768, F1: 2}
    var v14 float64 = float64(((0.0 + (-2.25 / 1e10)) - 1.5))
    for fi15 := 0; fi15 < 16; fi15++ {
        if (!(!(true || false))) {
            if ("\n\t" <= ((("a" + "foo bar") + "a") + "\n\t")) {
                fmt.Println(v5)
                if false {
                    v1 = append(v1, "")
                    if (true && false) {
                        for fi16 := 0; fi16 < 4; fi16++ {
                            switch "foo bar" {
                            case "\n\t":
                                v17 := GenS1{F0: (^(v2)), F1: (v2 / ((1) | 1)), F2: 2147483647}
                                _ = v17
                            case "":
                                var v18 uint16 = 887
                                _ = v18
                            case "hello":
                                fmt.Println(fi15, v5, v14)
                            }
                            if (!(true || false)) {
                                var v19 int64 = -341
                                v20 := v6[("\n\t" + "")]
                                var v21 uint16 = 42
                                _ = v19
                                _ = v20
                                _ = v21
                            } else {
                                if len(v1) > 0 {
                                    v1[0] = ("a" + "hello")
                                }
                                var v22 float32 = (((3.14159 * 0.1) - (-2.25 * (-2.25 / -2.25))) / 1.0)
                                v5 -= -122
                                _ = v22
                            }
                            v2 += v2
                            _ = fi16
                        }
                        if (true && (true && true)) {
                            v23 := GenS3{F0: 42, F1: 561, F2: 420}
                            _ = v23
                        } else {
                            _ = len(v6)
                            var v24 uint16 = 255
                            v25 := GenS4{F0: -726, F1: 287}
                            _ = v24
                            _ = v25
                        }
                    }
                    var v26 int64 = -1
                    _ = v26
                }
                for fi27 := 0; fi27 < 19; fi27++ {
                    if true {
                        v3 += uint8(v12)
                        v2 -= v2
                    } else {
                        if (-1 > 0) {
                            fmt.Println(v12, v3)
                            fmt.Println(v5, v12, v5)
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            v14 -= float64(-2.25)
                            v28 := GenS1{F0: v2, F1: v2, F2: -264}
                            if (-1 < -881) {
                                v29 := GenS3{F0: (v2 * 42), F1: 364, F2: -32768}
                                _ = v29
                            }
                            if true {
                                switch ("foo bar" + "\n\t") {
                                case "":
                                    v30 := GenS2{F0: 0, F1: 628}
                                    if true {
                                        var v31 bool = false
                                        var v32 float64 = 0.0
                                        var v33 uint64 = 1
                                        _ = v31
                                        _ = v32
                                        _ = v33
                                    }
                                    if ((v5 & v5) == v5) {
                                        v34 := GenS3{F0: uint(v12), F1: 2147483648, F2: 42}
                                        if true {
                                            v14 = float64(1e10)
                                        }
                                        _ = v34
                                    } else {
                                        v35 := v6["hello"]
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        _ = v35
                                    }
                                    _ = v30
                                }
                            } else {
                                v2 = (v2 | (v2 * (v2 << ((v2) & 63))))
                                v36 := GenS5{F0: (fi15 % ((fi27) | 1)), F1: v12, F2: -1}
                                fmt.Println(v3, fi27)
                                _ = v36
                            }
                            v37 := GenS6{F0: 0, F1: 968, F2: (v12 & (v12 % ((v12) | 1)))}
                            var v38 string = ""
                            _ = v28
                            _ = v37
                            _ = v38
                        }
                        if true {
                            if (985 > 1) {
                                var v39 int8 = (v5 ^ (v5 & v5))
                                _ = len(v6)
                                var v40 bool = (0.0 >= ((-2.25 * 1.5) / 1e10))
                                v5 = 0
                                _ = v39
                                _ = v40
                            } else {
                                var v41 int64 = -9223372036854775808
                                _ = v41
                            }
                            v3 = (^(v3))
                            fmt.Println(v3, v2)
                            switch ("foo bar" + (("hello" + "foo bar") + "\n\t")) {
                            case "hello":
                                if (false && true) {
                                    var v42 uint16 = 0
                                    var v43 uint64 = 465
                                    v44 := GenS3{F0: (v2 & 108), F1: 42, F2: 154}
                                    _ = v42
                                    _ = v43
                                    _ = v44
                                } else {
                                    if (false && (!false)) {
                                        v14 = (1.5 + v14)
                                    } else {
                                        v45 := GenS5{F0: int(v12), F1: (v12 % ((v12) | 1)), F2: -1}
                                        switch (fi15 >> ((v2) & 63)) {
                                        case -998:
                                            var v46 float64 = 1e10
                                            _ = v46
                                        case -732:
                                            v12 = v12
                                            v47 := GenS7{F0: ("a" + "hello")}
                                            v48 := GenS4{F0: fi15, F1: 42}
                                            _ = v47
                                            _ = v48
                                        case 127:
                                            v2 = uint(v3)
                                        }
                                        v3 -= v3
                                        v3 += 178
                                        v49 := GenS4{F0: 128, F1: -32768}
                                        if (!(false || true)) {
                                            var v50 uint = 1
                                            v14 = v14
                                            v14 -= 0.1
                                            var v51 int32 = 1
                                            v2 = uint(v12)
                                            _ = v50
                                            _ = v51
                                        } else {
                                            v52 := GenS1{F0: v2, F1: v2, F2: 2147483647}
                                            if (v2 != uint(v3)) {
                                                switch "" {
                                                case "":
                                                    var v53 int64 = 0
                                                    _ = v53
                                                case "foo bar":
                                                    var v54 int = (fi27 ^ (fi27 | (-(fi15))))
                                                    _ = v54
                                                default:
                                                    delete(v4, "hello")
                                                }
                                                v55 := GenS8{F0: "foo bar", F1: fi15, F2: uint16(v2)}
                                                var v56 uint8 = (v3 * uint8(v12))
                                                v57 := GenS5{F0: (^(fi15)), F1: v12, F2: 1}
                                                _ = len(v1)
                                                if ((v56 >> ((v2) & 63)) < 1) {
                                                    if (!true) {
                                                        v3 = v56
                                                        if (!true) {
                                                            var v58 int64 = 41
                                                            v2 += (^(v2))
                                                            var v59 uint8 = v56
                                                            _ = len(v1)
                                                            _ = v58
                                                            _ = v59
                                                        } else {
                                                            v56 -= v56
                                                        }
                                                    }
                                                    v60 := GenS5{F0: fi15, F1: (v12 | v12), F2: -9223372036854775808}
                                                    _ = v60
                                                } else {
                                                    if (168 <= v3) {
                                                        var v61 int = fi15
                                                        v3 += v56
                                                        _ = v61
                                                    }
                                                    v5 += (v5 << ((42) & 63))
                                                    var v62 int64 = -789
                                                    _ = v62
                                                }
                                                _ = v55
                                                _ = v56
                                                _ = v57
                                            } else {
                                                v5 = v5
                                            }
                                            v14 -= (1.5 * v14)
                                            _ = v52
                                        }
                                        v12 += 827
                                        v12 -= v12
                                        _ = v45
                                        _ = v49
                                    }
                                    v14 -= v14
                                }
                            case "":
                                switch ("hello" + "a") {
                                case "hello":
                                    _ = len(v1)
                                    _ = len(v1)
                                case "foo bar":
                                    switch (fi27 & fi27) {
                                    case 32768:
                                        v63 := GenS7{F0: ""}
                                        var v64 int = 813
                                        _ = v63
                                        _ = v64
                                    case 1:
                                        _ = len(v6)
                                    case -1:
                                        v65 := GenS9{F0: 560, F1: uint8(v12), F2: (v2 & (^(v2)))}
                                        _ = v65
                                    }
                                    if false {
                                        v66 := GenS10{F0: uint32(v2), F1: (^(v5))}
                                        _ = v66
                                    } else {
                                        v67 := GenS11{F0: 0, F1: 65535, F2: -515}
                                        var v68 int8 = (v5 & v5)
                                        _ = v67
                                        _ = v68
                                    }
                                case "":
                                    if true {
                                        switch 32768 {
                                        case 372:
                                            v12 = v12
                                        case -931:
                                            var v69 int16 = genFn0((-(fi27)))
                                            _ = v69
                                        default:
                                            var v70 float64 = v14
                                            _ = v70
                                        }
                                        if false {
                                            switch ("\n\t" + ("foo bar" + (((("" + "a") + "a") + "") + "foo bar"))) {
                                            case "\n\t":
                                                switch (fi15 << (((v2 << ((v2) & 63))) & 63)) {
                                                case 282:
                                                    if true {
                                                        var v71 string = "hello"
                                                        var v72 int32 = -403
                                                        _ = v71
                                                        _ = v72
                                                    } else {
                                                        v73 := GenS12{F0: 784, F1: 847, F2: v5}
                                                        if (65535 != 2147483648) {
                                                            v74 := GenS13{F0: uint8(v12), F1: 4294967295, F2: -96}
                                                            var v75 uint32 = 2
                                                            _ = v74
                                                            _ = v75
                                                        } else {
                                                            fmt.Println(fi27, fi27)
                                                            var v76 float32 = (float32(v14) / (0.1 / float32((v14 * 3.14159))))
                                                            _ = v76
                                                        }
                                                        _ = v73
                                                    }
                                                case 268:
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                case -9223372036854775808:
                                                    fmt.Println(v5, v2)
                                                default:
                                                    var v77 int16 = 1
                                                    _ = v77
                                                }
                                            case "hello":
                                                v78 := GenS13{F0: v3, F1: 674, F2: 1}
                                                _ = v78
                                            }
                                        } else {
                                            fmt.Println(v5, fi15)
                                            var v79 bool = false
                                            _ = v79
                                        }
                                    }
                                default:
                                    fmt.Println(v5)
                                }
                            default:
                                if false {
                                    fmt.Println(v3, v2, v5)
                                    v4["foo bar"] = (((-744 <= -2147483648) || true) && (false || true))
                                    var v80 int64 = -1
                                    if true {
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                    } else {
                                        v81 := GenS14{F0: "\n\t"}
                                        _ = v81
                                    }
                                    if (v14 >= 0.1) {
                                        switch "" {
                                        case "a":
                                            v82 := GenS2{F0: -525, F1: 12}
                                            _ = v82
                                        default:
                                            if ((255 == 1) && false) {
                                                var v83 int = fi27
                                                v12 = (v12 >> ((v2) & 63))
                                                if (!(-32768 <= 42)) {
                                                    v3 = (v3 / ((v3) | 1))
                                                    if len(v1) > 0 {
                                                        v1[0] = "hello"
                                                    }
                                                    switch "a" {
                                                    case "hello":
                                                        var v84 int16 = -32768
                                                        _ = v84
                                                    case "a":
                                                        var v85 bool = true
                                                        _ = v85
                                                    default:
                                                        v86 := GenS15{F0: ("\n\t" + "hello"), F1: (fi15 >> (((v2 * v2)) & 63)), F2: 1}
                                                        _ = v86
                                                    }
                                                }
                                                v87 := GenS15{F0: "hello", F1: fi15, F2: int(v2)}
                                                v88 := GenS16{F0: true}
                                                _ = v83
                                                _ = v87
                                                _ = v88
                                            }
                                        }
                                        v14 = float64((-2.25 / -2.25))
                                    }
                                    _ = v80
                                } else {
                                    switch (-(fi15)) {
                                    case 32767:
                                        v12 = v12
                                    default:
                                        switch "foo bar" {
                                        case "hello":
                                            var v89 uint = (^(v2))
                                            _ = v89
                                        }
                                    }
                                    var v90 int32 = 1
                                    _ = v90
                                }
                            }
                        }
                        var v91 bool = (1e10 != (v14 + v14))
                        _ = v91
                    }
                    _ = fi27
                }
            } else {
                fmt.Println(v2, v14)
                v2 += v2
            }
        }
        switch (fi15 / ((2147483647) | 1)) {
        case 32767:
            v14 += v14
        case 0:
            for fi92 := 0; fi92 < 15; fi92++ {
                var v93 float32 = ((1.5 - (0.0 * float32(1e10))) * -2.25)
                _ = fi92
                _ = v93
            }
        default:
            if true {
                var v94 int32 = -1
                if (-2.25 > 1.5) {
                    var v95 uint64 = 0
                    if true {
                        v96 := GenS1{F0: 0, F1: (v2 + (v2 * v2)), F2: (v94 - (v94 << ((1) & 63)))}
                        switch (fi15 + (fi15 * (fi15 % (((fi15 & fi15)) | 1)))) {
                        case 272:
                            var v97 uint = (v2 ^ (v2 ^ (v2 << (((v2 * v2)) & 63))))
                            _ = v97
                        case -2147483648:
                            v95 += (^(v12))
                        }
                        _ = v96
                    } else {
                        var v98 int32 = (v94 & v94)
                        _ = v98
                    }
                    fmt.Println(v3, v5, fi15)
                    _ = v95
                } else {
                    switch fi15 {
                    case -1:
                        v99 := GenS13{F0: v3, F1: 2147483647, F2: 1}
                        _ = v99
                    case 4294967295:
                        if false {
                            fmt.Println(fi15, v3)
                            if false {
                                var v100 float64 = (1e10 + v14)
                                fmt.Println(v14, v94, v100)
                                _ = v100
                            } else {
                                for fi101 := 0; fi101 < 8; fi101++ {
                                    var v102 bool = false
                                    v102 = (496 < 42)
                                    _ = fi101
                                    _ = v102
                                }
                                var v103 int32 = (v94 * 0)
                                v104 := GenS6{F0: v12, F1: (v94 >> (((v2 * 155)) & 63)), F2: v12}
                                _ = v103
                                _ = v104
                            }
                            if (-1 <= 42) {
                                v105 := GenS11{F0: 2, F1: 0, F2: int32(fi15)}
                                _ = v105
                            }
                        }
                    default:
                        v12 -= 524
                    }
                    var v106 float64 = float64((-2.25 / 1.5))
                    _ = v106
                }
                _ = v94
            } else {
                if len(v1) > 0 {
                    v1[0] = "hello"
                }
                switch fi15 {
                case 127:
                    var v107 float64 = 3.14159
                    _ = v107
                case 349:
                    switch (fi15 / ((fi15) | 1)) {
                    case -2147483648:
                        var v108 int32 = -978
                        _ = v108
                    default:
                        fmt.Println(v2, fi15)
                    }
                default:
                    switch "a" {
                    case "hello":
                        fmt.Println(fi15, v14)
                    case "foo bar":
                        switch ("hello" + "hello") {
                        case "\n\t":
                            for fi109 := 0; fi109 < 12; fi109++ {
                                v12 = 1
                                var v110 int8 = (^(v5))
                                if true {
                                    switch fi109 {
                                    case 4294967295:
                                        v3 = (^(v3))
                                    default:
                                        fmt.Println(fi15)
                                    }
                                    var v111 int64 = 935
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    _ = v111
                                } else {
                                    var v112 uint16 = 255
                                    _ = v112
                                }
                                _ = fi109
                                _ = v110
                            }
                        }
                    case "\n\t":
                        for fi113 := 0; fi113 < 2; fi113++ {
                            v114 := GenS5{F0: fi113, F1: (v12 / ((uint64(v2)) | 1)), F2: -9223372036854775808}
                            _ = fi113
                            _ = v114
                        }
                    }
                }
                switch "foo bar" {
                case "a":
                    if (-1 == 0) {
                        fmt.Println(fi15)
                        for fi115 := 0; fi115 < 16; fi115++ {
                            v116 := GenS17{F0: (v2 >> ((v2) & 63)), F1: "foo bar"}
                            var v117 int64 = 925
                            v118 := v4[""]
                            _ = fi115
                            _ = v116
                            _ = v117
                            _ = v118
                        }
                    } else {
                        fmt.Println(v3)
                        switch (-(fi15)) {
                        case -54:
                            switch "hello" {
                            case "hello":
                                for fi119 := 0; fi119 < 12; fi119++ {
                                    var v120 int8 = (v5 + (^(v5)))
                                    _ = fi119
                                    _ = v120
                                }
                            case "":
                                var v121 string = (("" + ("foo bar" + "foo bar")) + "hello")
                                _ = v121
                            case "\n\t":
                                var v122 uint64 = (^(v12))
                                _ = v122
                            default:
                                var v123 int = (fi15 ^ (fi15 & fi15))
                                _ = v123
                            }
                        case 0:
                            var v124 int8 = (v5 << ((42) & 63))
                            _ = v124
                        case 2147483647:
                            switch "foo bar" {
                            case "":
                                var v125 string = (("a" + (("" + "") + "\n\t")) + "")
                                _ = v125
                            default:
                                for fi126 := 0; fi126 < 8; fi126++ {
                                    delete(v6, ("foo bar" + "hello"))
                                    switch ("a" + ("a" + "")) {
                                    case "foo bar":
                                        v3 -= v3
                                    }
                                    switch ("foo bar" + "") {
                                    case "hello":
                                        switch 128 {
                                        case 0:
                                            var v127 int16 = 0
                                            _ = v127
                                        case 446:
                                            var v128 uint16 = 32767
                                            _ = v128
                                        case 127:
                                            if ((float64((0.1 / 3.14159)) * (v14 - (float64(0.1) - float64(0.0)))) == 0.1) {
                                                v3 += 0
                                                var v129 int16 = genFn0(fi126)
                                                _ = v129
                                            }
                                        }
                                    default:
                                        switch "\n\t" {
                                        case "\n\t":
                                            if (uint64(v3) >= v12) {
                                                var v130 uint32 = 979
                                                v131 := GenS9{F0: 446, F1: (^(v3)), F2: v2}
                                                _ = v130
                                                _ = v131
                                            }
                                        case "foo bar":
                                            v14 -= v14
                                        case "hello":
                                            v132 := GenS11{F0: 255, F1: 246, F2: -2147483648}
                                            _ = v132
                                        }
                                    }
                                    _ = fi126
                                }
                            }
                        default:
                            _ = len(v6)
                        }
                        if ((v3 | (v3 & v3)) == v3) {
                            v133 := GenS2{F0: 0, F1: 42}
                            var v134 uint32 = 859
                            v135 := GenS3{F0: (v2 - (v2 - (v2 | v2))), F1: (^(v134)), F2: genFn0((fi15 % ((fi15) | 1)))}
                            _ = v133
                            _ = v134
                            _ = v135
                        }
                    }
                case "hello":
                    var v136 int = fi15
                    _ = v136
                default:
                    var v137 int32 = 1
                    _ = v137
                }
            }
        }
        var v138 float32 = (1.5 - (float32((((v14 - 1e10) + (v14 + v14)) - v14)) - 1.5))
        _ = fi15
        _ = v138
    }
    if true {
        v3 = v3
    } else {
        var v139 uint64 = v12
        v140 := v4[("hello" + "")]
        _ = v139
        _ = v140
    }
    v6[("foo bar" + ("a" + "foo bar"))] = true
    switch 2147483647 {
    case -689:
        v141 := v6["hello"]
        _ = v141
    case 32767:
        v142 := GenS9{F0: 1, F1: (v3 + (^(v3))), F2: (v2 + v2)}
        _ = v142
    default:
        v5 += (v5 << ((4294967295) & 63))
    }
    var v143 float64 = (v14 - float64((float32((v14 * v14)) - ((0.0 - float32(3.14159)) + 0.0))))
    if (!(616 <= 437)) {
        var v144 uint64 = (v12 | v12)
        for fi145 := 0; fi145 < 13; fi145++ {
            for fi146 := 0; fi146 < 6; fi146++ {
                v147 := GenS18{F0: 42, F1: (1e10 == ((0.1 - float64(float32((((v14 - v14) - v14) * (((((float64(-2.25) + (v143 / v143)) - v143) * -2.25) / 3.14159) / (0.0 / (v14 + 3.14159))))))) * 3.14159)), F2: fi145}
                if (fi146 >= (fi145 & (fi145 % ((fi145) | 1)))) {
                    switch "" {
                    case "a":
                        v12 -= 2
                    case "":
                        if (("foo bar" < ("foo bar" + "")) || true) {
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            switch "foo bar" {
                            case "hello":
                                if ((true || true) && true) {
                                    var v148 uint16 = 255
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    switch 4294967295 {
                                    case 267:
                                        v149 := GenS19{F0: 147, F1: 1, F2: "\n\t"}
                                        _ = v149
                                    case -91:
                                        if true {
                                            switch (fi146 % ((32768) | 1)) {
                                            case 1:
                                                var v150 uint64 = 0
                                                _ = v150
                                            }
                                            v143 -= v143
                                            var v151 uint8 = uint8(v148)
                                            _ = v151
                                        } else {
                                            var v152 uint = v2
                                            _ = v152
                                        }
                                    case 617:
                                        v153 := GenS20{F0: ("\n\t" + ("a" + "foo bar")), F1: 42}
                                        _ = v153
                                    default:
                                        switch (("foo bar" + "a") + "\n\t") {
                                        case "foo bar":
                                            if (false && (false && true)) {
                                                var v154 uint16 = (v148 >> ((v2) & 63))
                                                if (!(!(!true))) {
                                                    var v155 bool = (((v143 - v143) * (((1e10 * 3.14159) + v14) + float64(3.14159))) < v14)
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                    var v156 int = (fi146 << ((65535) & 63))
                                                    _ = v155
                                                    _ = v156
                                                }
                                                _ = v154
                                            }
                                        case "hello":
                                            v4["\n\t"] = true
                                        case "":
                                            delete(v6, "")
                                        }
                                    }
                                    v157 := GenS21{F0: (v3 - v3), F1: "a", F2: 842}
                                    _ = v148
                                    _ = v157
                                }
                            }
                        }
                    }
                }
                _ = fi146
                _ = v147
            }
            _ = fi145
        }
        v158 := GenS8{F0: (("a" + "hello") + ("" + ("a" + (("a" + ("foo bar" + "a")) + "")))), F1: 4294967295, F2: 799}
        _ = v144
        _ = v158
    } else {
        for fi159 := 0; fi159 < 19; fi159++ {
            if (-2.25 > ((((3.14159 / 0.1) * (-2.25 / 3.14159)) - ((0.0 + (1.5 + 1.5)) * 1e10)) * (((-2.25 / 1.0) / 1.5) * (float32((-2.25 * (3.14159 + float64(1e10)))) * 1e10)))) {
                var v160 int = 4294967295
                v2 -= uint(v12)
                _ = v160
            } else {
                switch "hello" {
                case "":
                    for fi161 := 0; fi161 < 15; fi161++ {
                        var v162 int8 = (^(v5))
                        _ = fi161
                        _ = v162
                    }
                }
            }
            switch fi159 {
            case 32767:
                if (69 > (v3 / ((v3) | 1))) {
                    if len(v1) > 0 {
                        v1[0] = ("a" + "a")
                    }
                }
            case -2147483648:
                if false {
                    for fi163 := 0; fi163 < 11; fi163++ {
                        var v164 uint64 = 42
                        v3 = uint8(v2)
                        fmt.Println(fi159, v2, v2)
                        _ = fi163
                        _ = v164
                    }
                    fmt.Println(fi159, fi159)
                    v165 := GenS14{F0: ("\n\t" + "foo bar")}
                    _ = v165
                } else {
                    fmt.Println(fi159, v5)
                    v2 -= v2
                    var v166 float64 = v14
                    _ = v166
                }
            case 4294967295:
                fmt.Println(fi159)
            }
            switch (^(fi159)) {
            case -11:
                if (-20 > int64(v5)) {
                    v3 = 0
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                } else {
                    switch -1 {
                    case 1:
                        fmt.Println(fi159, v5)
                    case 32767:
                        switch -663 {
                        case 1:
                            v167 := GenS6{F0: (v12 >> ((v2) & 63)), F1: -999, F2: v12}
                            _ = v167
                        case -9223372036854775808:
                            for fi168 := 0; fi168 < 11; fi168++ {
                                switch 9223372036854775807 {
                                case 42:
                                    v1 = append(v1, "a")
                                }
                                var v169 uint16 = 879
                                _ = fi168
                                _ = v169
                            }
                        }
                    }
                }
            case 128:
                if len(v1) > 0 {
                    v1[0] = "\n\t"
                }
            }
            _ = fi159
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v143:", v143)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v143
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
