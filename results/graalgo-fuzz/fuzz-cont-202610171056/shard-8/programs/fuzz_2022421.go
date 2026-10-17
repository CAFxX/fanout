package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint64; F1 uint32; F2 int64 }
type GenS3 struct { F0 uint8; F1 int64; F2 uint }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int64; F1 int8 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 uint64; F1 uint }
type GenS8 struct { F0 int8 }
type GenS9 struct { F0 uint64 }
type GenS10 struct { F0 string; F1 string }
type GenS11 struct { F0 uint32; F1 uint }
type GenS12 struct { F0 string; F1 int; F2 string }
type GenS13 struct { F0 int16 }
type GenS14 struct { F0 string; F1 int }
type GenS15 struct { F0 string }
type GenS16 struct { F0 uint64; F1 int16 }
type GenS17 struct { F0 uint8; F1 int64 }
type GenS18 struct { F0 int }
func genFn0(p0 int) uint32 {
    for fi6 := 0; fi6 < 16; fi6++ {
        var v7 uint8 = 255
        v8 := []int{fi6, (p0 | (p0 >> ((4294967295) & 63))), fi6, p0}
        if (42 <= 0) {
            for fi9 := 0; fi9 < 4; fi9++ {
                if false {
                    if ((true || (-877 == -930)) || false) {
                        p0 += -1
                    } else {
                        v7 += v7
                    }
                } else {
                    if (2147483647 != int32(v7)) {
                        if (309 >= -364) {
                            p0 -= p0
                            var v10 uint32 = 2
                            _ = v10
                        }
                        v8 = append(v8, -9223372036854775808)
                        fmt.Println(p0, fi9)
                    } else {
                        if true {
                            p0 = (p0 - (fi6 << ((65535) & 63)))
                            var v11 uint32 = 2147483648
                            _ = v11
                        }
                    }
                    v7 = 42
                    _ = len(v8)
                }
                if len(v8) >= 2 {
                    v8 = v8[1:len(v8)]
                }
                _ = fi9
            }
            if (1 < 255) {
                p0 = fi6
                for fi12 := 0; fi12 < 1; fi12++ {
                    v8 = append(v8, fi6)
                    var v13 uint16 = 32767
                    _ = fi12
                    _ = v13
                }
            }
        } else {
            var v14 uint64 = 18446744073709551615
            v15 := GenS1{F0: 321}
            switch ("\n\t" + "hello") {
            case "hello":
                if len(v8) > 0 {
                    v8[0] = (^(fi6))
                }
            case "\n\t":
                var v16 uint64 = (v14 ^ (v14 + v14))
                _ = v16
            default:
                var v17 int64 = 326
                _ = v17
            }
            _ = v14
            _ = v15
        }
        _ = fi6
        _ = v7
        _ = v8
    }
    v18 := GenS2{F0: 18446744073709551615, F1: 1, F2: 0}
    _ = p0
    _ = v18
    return 230
}

func main() {
    v1 := []string{"\n\t"}
    var v2 uint8 = 127
    var v3 string = "hello"
    var v4 string = "\n\t"
    var v5 float32 = float32(0.1)
    if false {
        var v19 string = (v3 + "a")
        if false {
            v1 = append(v1, ("hello" + "foo bar"))
        } else {
            v3 = "\n\t"
            for fi20 := 0; fi20 < 2; fi20++ {
                fmt.Println(fi20, v3)
                fmt.Println(v2)
                switch "foo bar" {
                case "foo bar":
                    var v21 uint64 = 913
                    _ = v21
                case "hello":
                    var v22 float64 = 0.0
                    _ = v22
                }
                _ = fi20
            }
        }
        if true {
            fmt.Println(v3, v4)
        } else {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            switch 127 {
            case -2147483648:
                v19 = "a"
            default:
                for fi23 := 0; fi23 < 12; fi23++ {
                    v19 = v4
                    var v24 uint16 = 298
                    switch "" {
                    case "hello":
                        v4 = v19
                    case "a":
                        var v25 int16 = 32767
                        _ = v25
                    case "\n\t":
                        if false {
                            v26 := GenS3{F0: (v2 << ((1) & 63)), F1: -441, F2: 219}
                            var v27 uint8 = (v2 / (((v2 + v2)) | 1))
                            for fi28 := 0; fi28 < 18; fi28++ {
                                var v29 uint8 = v27
                                _ = fi28
                                _ = v29
                            }
                            _ = v26
                            _ = v27
                        }
                    default:
                        for fi30 := 0; fi30 < 12; fi30++ {
                            var v31 uint8 = (v2 - v2)
                            var v32 uint16 = v24
                            if len(v1) > 0 {
                                v1[0] = v4
                            }
                            v1 = append(v1, "foo bar")
                            v33 := GenS4{F0: 841}
                            _ = len(v1)
                            _ = fi30
                            _ = v31
                            _ = v32
                            _ = v33
                        }
                    }
                    _ = fi23
                    _ = v24
                }
            }
        }
        _ = v19
    }
    v34 := GenS4{F0: 42}
    var v35 int32 = 318
    var v36 int = 1
    for fi37 := 0; fi37 < 16; fi37++ {
        if (!(103 == 100)) {
            if len(v1) > 0 {
                v1[0] = v3
            }
            v38 := GenS3{F0: (^(v2)), F1: 1, F2: 18446744073709551615}
            _ = v38
        } else {
            for fi39 := 0; fi39 < 12; fi39++ {
                fmt.Println(fi37, v5, v36)
                _ = len(v1)
                _ = fi39
            }
            if false {
                if len(v1) > 0 {
                    v1[0] = (("\n\t" + "") + v3)
                }
                if ((fi37 >> ((4294967295) & 63)) <= v36) {
                    v5 += 0.0
                    var v40 uint32 = 971
                    _ = v40
                } else {
                    switch v3 {
                    case "\n\t":
                        for fi41 := 0; fi41 < 15; fi41++ {
                            if (127 == -53) {
                                var v42 int64 = 42
                                _ = v42
                            } else {
                                var v43 int32 = (^(v35))
                                v44 := GenS1{F0: 65535}
                                _ = v43
                                _ = v44
                            }
                            if true {
                                v45 := GenS5{F0: 42, F1: 85}
                                _ = v45
                            } else {
                                _ = len(v1)
                                fmt.Println(fi37)
                            }
                            _ = fi41
                        }
                    case "foo bar":
                        v5 += v5
                    case "":
                        fmt.Println(v35, v36)
                    default:
                        for fi46 := 0; fi46 < 16; fi46++ {
                            var v47 uint32 = 844
                            var v48 uint64 = 4294967295
                            _ = fi46
                            _ = v47
                            _ = v48
                        }
                    }
                    v49 := GenS5{F0: 336, F1: -128}
                    if (true && (true || ((^(v2)) <= 128))) {
                        switch v36 {
                        case 200:
                            for fi50 := 0; fi50 < 6; fi50++ {
                                v51 := GenS6{F0: false}
                                if false {
                                    switch (v36 - 0) {
                                    case 32768:
                                        var v52 int = 1
                                        _ = v52
                                    case 9223372036854775807:
                                        if (!(true || false)) {
                                            var v53 bool = false
                                            switch ((v3 + (("" + v3) + v4)) + ((v4 + "\n\t") + v3)) {
                                            case "\n\t":
                                                v2 = 2
                                            case "":
                                                v35 = (v35 % ((v35) | 1))
                                            }
                                            if (0 > 396) {
                                                v53 = v53
                                                v54 := GenS3{F0: 128, F1: 1, F2: 1}
                                                _ = v54
                                            }
                                            _ = v53
                                        } else {
                                            fmt.Println(v2, v4, fi50)
                                        }
                                    case 0:
                                        var v55 float32 = 0.1
                                        if false {
                                            v3 = ((v4 + v4) + (v4 + ((v4 + (((v4 + "foo bar") + (v3 + v4)) + ((v4 + "a") + (v4 + v3)))) + ("\n\t" + v3))))
                                            var v56 bool = true
                                            _ = v56
                                        }
                                        if (((^(v2)) == v2) || true) {
                                            var v57 float64 = ((0.0 + float64(1.5)) + ((1e10 + 1.5) * (float64(1.5) + 0.0)))
                                            var v58 int64 = -1
                                            v36 = fi37
                                            _ = v57
                                            _ = v58
                                        }
                                        _ = v55
                                    default:
                                        _ = len(v1)
                                        var v59 int16 = 32767
                                        _ = v59
                                    }
                                    fmt.Println(v35)
                                    v60 := GenS7{F0: 1, F1: 0}
                                    _ = v60
                                } else {
                                    var v61 float32 = (((1e10 - (0.1 + 0.1)) - float32(0.0)) / (v5 / float32(-2.25)))
                                    if (((!(int64(fi50) <= -356)) && (1 == -128)) || false) {
                                        fmt.Println(fi37)
                                        v2 -= 42
                                        v62 := GenS8{F0: 127}
                                        var v63 uint64 = 264
                                        _ = v62
                                        _ = v63
                                    } else {
                                        fmt.Println(v5, v36, v2)
                                        switch v4 {
                                        case "":
                                            if false {
                                                if true {
                                                    var v64 uint8 = (v2 * (^(v2)))
                                                    switch fi37 {
                                                    case 4294967295:
                                                        _ = len(v1)
                                                    default:
                                                        var v65 int = fi50
                                                        _ = v65
                                                    }
                                                    _ = v64
                                                } else {
                                                    if (0 > (v2 & 2)) {
                                                        var v66 uint64 = 0
                                                        if ((-(fi37)) >= (fi50 + (^(fi37)))) {
                                                            var v67 uint32 = 2147483647
                                                            var v68 int16 = 275
                                                            var v69 int64 = 42
                                                            switch (v36 + (fi50 / ((v36) | 1))) {
                                                            case 42:
                                                                if (!("a" < v3)) {
                                                                    v67 = v67
                                                                }
                                                            case 9223372036854775807:
                                                                if false {
                                                                    v66 -= v66
                                                                    var v70 string = v3
                                                                    fmt.Println(v36)
                                                                    _ = v70
                                                                }
                                                            }
                                                            _ = v67
                                                            _ = v68
                                                            _ = v69
                                                        } else {
                                                            switch (("a" + "a") + v4) {
                                                            case "\n\t":
                                                                var v71 float32 = (3.14159 + v5)
                                                                _ = v71
                                                            case "":
                                                                _ = len(v1)
                                                                v36 = v36
                                                            case "a":
                                                                if false {
                                                                    if false {
                                                                        var v72 bool = (false && false)
                                                                        var v73 uint32 = 87
                                                                        _ = v72
                                                                        _ = v73
                                                                    } else {
                                                                        v2 += v2
                                                                        if (!false) {
                                                                            v35 += 42
                                                                        }
                                                                    }
                                                                    v74 := GenS1{F0: 1}
                                                                    _ = v74
                                                                }
                                                            default:
                                                                var v75 float32 = (v5 * (v5 / v5))
                                                                _ = v75
                                                            }
                                                            var v76 uint8 = v2
                                                            _ = v76
                                                        }
                                                        v4 = "hello"
                                                        _ = v66
                                                    } else {
                                                        if (false || (false && false)) {
                                                            var v77 uint8 = v2
                                                            _ = v77
                                                        }
                                                        if len(v1) >= 2 {
                                                            v1 = v1[1:len(v1)]
                                                        }
                                                        if (-1 != -49) {
                                                            if false {
                                                                if (1 < 1) {
                                                                    var v78 uint64 = 1
                                                                    _ = v78
                                                                } else {
                                                                    fmt.Println(v3)
                                                                    if (!true) {
                                                                        _ = len(v1)
                                                                        if true {
                                                                            fmt.Println(v61, fi37, v61)
                                                                            v5 = float32(float64((v61 * v61)))
                                                                            v5 -= v5
                                                                            var v79 float64 = 0.1
                                                                            v80 := GenS3{F0: v2, F1: 1, F2: 1}
                                                                            _ = v79
                                                                            _ = v80
                                                                        }
                                                                    } else {
                                                                        if true {
                                                                            if (25 >= 65535) {
                                                                                var v81 int16 = 1
                                                                                _ = len(v1)
                                                                                v82 := GenS9{F0: 905}
                                                                                _ = v81
                                                                                _ = v82
                                                                            }
                                                                        } else {
                                                                            var v83 uint64 = 4294967295
                                                                            var v84 uint = 434
                                                                            _ = v83
                                                                            _ = v84
                                                                        }
                                                                        if (-32768 != 0) {
                                                                            var v85 float32 = v61
                                                                            _ = v85
                                                                        } else {
                                                                            if len(v1) > 0 {
                                                                                v1[0] = ("hello" + (("hello" + (v4 + ("hello" + (((v3 + v4) + v3) + v4)))) + v3))
                                                                            }
                                                                            if false {
                                                                                switch "\n\t" {
                                                                                case "a":
                                                                                    switch fi37 {
                                                                                    case -781:
                                                                                        if true {
                                                                                            switch 32768 {
                                                                                            case -808:
                                                                                                var v86 float64 = 3.14159
                                                                                                _ = v86
                                                                                            }
                                                                                            _ = len(v1)
                                                                                        }
                                                                                    case 42:
                                                                                        if false {
                                                                                            v87 := GenS10{F0: "a", F1: (("\n\t" + (("foo bar" + v4) + v4)) + v4)}
                                                                                            _ = len(v1)
                                                                                            _ = v87
                                                                                        }
                                                                                        switch (fi50 ^ (fi37 & 2147483647)) {
                                                                                        case 2147483647:
                                                                                            switch v36 {
                                                                                            case 128:
                                                                                                if len(v1) > 0 {
                                                                                                    v1[0] = ("" + v4)
                                                                                                }
                                                                                            case 1:
                                                                                                _ = len(v1)
                                                                                            case 42:
                                                                                                var v88 int8 = int8(v35)
                                                                                                _ = v88
                                                                                            }
                                                                                        case -1:
                                                                                            v35 = -2147483648
                                                                                        case 32768:
                                                                                            var v89 int = (fi50 + (fi37 | (fi50 * fi37)))
                                                                                            _ = v89
                                                                                        }
                                                                                    }
                                                                                case "hello":
                                                                                    fmt.Println(v2)
                                                                                case "\n\t":
                                                                                    var v90 int64 = -1
                                                                                    _ = v90
                                                                                }
                                                                                v91 := GenS2{F0: 808, F1: 2147483647, F2: 1}
                                                                                _ = v91
                                                                            }
                                                                            v61 -= -2.25
                                                                        }
                                                                        v5 = v61
                                                                    }
                                                                    var v92 uint16 = 2
                                                                    _ = v92
                                                                }
                                                                v93 := GenS11{F0: genFn0(fi50), F1: 490}
                                                                _ = v93
                                                            } else {
                                                                if len(v1) > 0 {
                                                                    v1[0] = (((((v3 + "\n\t") + "") + "") + (("hello" + v3) + "\n\t")) + v3)
                                                                }
                                                                var v94 uint8 = (v2 >> ((2) & 63))
                                                                _ = v94
                                                            }
                                                            _ = len(v1)
                                                        } else {
                                                            if len(v1) > 0 {
                                                                v1[0] = v3
                                                            }
                                                            v95 := GenS3{F0: (v2 % ((2) | 1)), F1: 42, F2: 17}
                                                            var v96 int32 = (-(v35))
                                                            _ = v95
                                                            _ = v96
                                                        }
                                                    }
                                                }
                                                if (true && false) {
                                                    var v97 uint16 = 42
                                                    switch 32767 {
                                                    case 2147483647:
                                                        var v98 int8 = 1
                                                        v61 += 3.14159
                                                        _ = v98
                                                    }
                                                    var v99 int32 = v35
                                                    fmt.Println(v4)
                                                    _ = v97
                                                    _ = v99
                                                }
                                            } else {
                                                v35 += v35
                                            }
                                        case "foo bar":
                                            var v100 int16 = 1000
                                            _ = v100
                                        default:
                                            if true {
                                                var v101 uint64 = 4294967295
                                                var v102 int = 32767
                                                var v103 float32 = (float32(1e10) + v61)
                                                if false {
                                                    if len(v1) > 0 {
                                                        v1[0] = "foo bar"
                                                    }
                                                    if ((v2 >> ((255) & 63)) == (v2 * v2)) {
                                                        var v104 float64 = (-2.25 - (0.1 - 3.14159))
                                                        var v105 int8 = -94
                                                        if len(v1) >= 2 {
                                                            v1 = v1[1:len(v1)]
                                                        }
                                                        _ = v104
                                                        _ = v105
                                                    } else {
                                                        if (7 <= 22) {
                                                            fmt.Println(fi37, v3, v101)
                                                            v36 = 9223372036854775807
                                                        } else {
                                                            v106 := GenS1{F0: 32767}
                                                            _ = v106
                                                        }
                                                        _ = len(v1)
                                                    }
                                                    v1 = append(v1, "\n\t")
                                                }
                                                _ = v101
                                                _ = v102
                                                _ = v103
                                            }
                                        }
                                    }
                                    v107 := GenS8{F0: 42}
                                    _ = v107
                                    _ = v61
                                }
                                v108 := GenS6{F0: true}
                                fmt.Println(fi50, v2, v2)
                                _ = fi50
                                _ = v108
                                _ = v51
                            }
                        default:
                            for fi109 := 0; fi109 < 6; fi109++ {
                                v3 = v3
                                _ = fi109
                            }
                        }
                        for fi110 := 0; fi110 < 17; fi110++ {
                            var v111 int = (v36 & (fi110 % (((v36 - v36)) | 1)))
                            if len(v1) > 0 {
                                v1[0] = ""
                            }
                            _ = fi110
                            _ = v111
                        }
                        var v112 uint32 = 2147483648
                        _ = v112
                    }
                    _ = v49
                }
            }
            for fi113 := 0; fi113 < 1; fi113++ {
                if true {
                    fmt.Println(fi37)
                    v35 -= (v35 >> ((649) & 63))
                } else {
                    v114 := GenS7{F0: 2, F1: 2}
                    v36 -= -2147483648
                    var v115 uint8 = 127
                    _ = v114
                    _ = v115
                }
                var v116 int64 = 506
                var v117 bool = true
                var v118 int16 = 720
                _ = fi113
                _ = v116
                _ = v117
                _ = v118
            }
        }
        for fi119 := 0; fi119 < 1; fi119++ {
            v5 = (1e10 + 1.5)
            var v120 int = 1
            var v121 uint32 = 0
            _ = fi119
            _ = v120
            _ = v121
        }
        _ = fi37
    }
    v1 = append(v1, v4)
    v35 = (v35 >> ((132) & 63))
    if (!(!((1000 != 0) || false))) {
        if (v35 != (v35 << ((42) & 63))) {
            if (42 <= 0) {
                if (1 < 100) {
                    if false {
                        if true {
                            v36 = (v36 | v36)
                            var v122 int8 = 42
                            _ = v122
                        } else {
                            v123 := GenS12{F0: (("" + v4) + ""), F1: (v36 | v36), F2: "a"}
                            _ = v123
                        }
                        v124 := GenS10{F0: (v3 + ((v4 + "") + v4)), F1: (("foo bar" + "\n\t") + (v4 + "foo bar"))}
                        _ = v124
                    }
                    var v125 bool = ((v35 << ((2) & 63)) < 1)
                    _ = v125
                } else {
                    switch v36 {
                    case -432:
                        v126 := GenS4{F0: 610}
                        _ = v126
                    case -32:
                        v5 = v5
                    }
                    var v127 uint8 = 1
                    _ = v127
                }
                fmt.Println(v2, v36)
            }
        } else {
            switch v36 {
            case -344:
                v2 = v2
            case 0:
                for fi128 := 0; fi128 < 16; fi128++ {
                    var v129 float32 = (1.5 + (v5 * v5))
                    var v130 float32 = v5
                    if (false && ((-727 < 42) && false)) {
                        for fi131 := 0; fi131 < 17; fi131++ {
                            if (true || true) {
                                switch fi131 {
                                case 42:
                                    v132 := GenS8{F0: 127}
                                    _ = v132
                                default:
                                    var v133 uint8 = (v2 - v2)
                                    _ = v133
                                }
                            } else {
                                if len(v1) > 0 {
                                    v1[0] = v3
                                }
                            }
                            v130 = v130
                            switch fi128 {
                            case -9223372036854775808:
                                v134 := GenS13{F0: 1}
                                switch fi131 {
                                case 32768:
                                    v135 := GenS7{F0: 335, F1: 255}
                                    _ = v135
                                case -1:
                                    var v136 uint32 = 2
                                    var v137 int64 = -241
                                    fmt.Println(fi131, v2)
                                    v138 := GenS14{F0: v4, F1: 127}
                                    if true {
                                        var v139 uint32 = v136
                                        v140 := GenS1{F0: 1}
                                        var v141 uint8 = (v2 / ((255) | 1))
                                        _ = v139
                                        _ = v140
                                        _ = v141
                                    } else {
                                        if false {
                                            if (v137 > v137) {
                                                v137 = (v137 * v137)
                                                var v142 string = v3
                                                _ = v142
                                            }
                                            fmt.Println(v130, v4, v2)
                                            v35 -= v35
                                        }
                                        v137 += (-(v137))
                                    }
                                    if (-325 != 4294967295) {
                                        v143 := GenS15{F0: ("hello" + v4)}
                                        _ = v143
                                    } else {
                                        v36 -= fi131
                                    }
                                    _ = v136
                                    _ = v137
                                    _ = v138
                                default:
                                    if false {
                                        var v144 int = 9223372036854775807
                                        var v145 int16 = 72
                                        v146 := GenS16{F0: 806, F1: 42}
                                        _ = v144
                                        _ = v145
                                        _ = v146
                                    }
                                }
                                v2 += v2
                                _ = v134
                            default:
                                v147 := GenS17{F0: 42, F1: -9223372036854775808}
                                _ = v147
                            }
                            _ = fi131
                        }
                        var v148 float64 = 3.14159
                        _ = v148
                    }
                    _ = fi128
                    _ = v129
                    _ = v130
                }
            case 4294967295:
                v35 = v35
            default:
                v149 := GenS2{F0: 991, F1: 2147483647, F2: -1}
                _ = v149
            }
            if false {
                var v150 int32 = v35
                var v151 int32 = -1
                var v152 uint8 = (v2 << ((42) & 63))
                _ = v150
                _ = v151
                _ = v152
            } else {
                fmt.Println(v4, v4, v2)
                v1 = append(v1, "hello")
                _ = len(v1)
            }
        }
    }
    if (!(!(3.14159 <= float64(1.5)))) {
        var v153 uint16 = 335
        v5 = 3.14159
        if false {
            v153 += v153
            var v154 uint = 42
            v155 := GenS18{F0: (-(v36))}
            _ = v154
            _ = v155
        } else {
            fmt.Println(v3, v2)
        }
        _ = v153
    }
    if (1000 == 1000) {
        fmt.Println(v4, v5)
        var v156 float64 = 0.1
        _ = v156
    }
    var v157 int = v36
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v157:", v157)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v157
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v5
}
