package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int }
type GenS2 struct { F0 int16; F1 int16 }
type GenS3 struct { F0 int; F1 int32; F2 uint64 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 uint64; F1 string; F2 int }
type GenS8 struct { F0 uint8; F1 uint32 }
type GenS9 struct { F0 uint32; F1 string; F2 uint8 }
type GenS10 struct { F0 uint64; F1 uint32 }
type GenS11 struct { F0 int16 }
type GenS12 struct { F0 string }
type GenS13 struct { F0 uint16; F1 uint32; F2 uint64 }
type GenS14 struct { F0 int8; F1 uint }
type GenS15 struct { F0 uint8 }
type GenS16 struct { F0 int16 }
func genFn0(p0 uint16, p1 uint, p2 uint16) int8 {
    var v6 float64 = 1.5
    var v7 int16 = 902
    switch 654 {
    case 42:
        v8 := GenS1{F0: 0, F1: -327}
        _ = v8
    case -645:
        v9 := []int{128, 32767}
        _ = v9
    default:
        for fi10 := 0; fi10 < 15; fi10++ {
            v11 := GenS1{F0: 2, F1: (fi10 & fi10)}
            var v12 float32 = (float32(1.5) / (float32((v6 - float64(((-2.25 * ((float32(((((v6 * (v6 * v6)) + v6) + ((v6 * v6) * float64(((0.1 * -2.25) - 1e10)))) + v6)) + 0.0) * (1.5 / 0.1))) / 1.5)))) / (-2.25 / float32(-2.25))))
            var v13 uint64 = uint64(p1)
            _ = fi10
            _ = v11
            _ = v12
            _ = v13
        }
    }
    fmt.Println(p1, p2)
    for fi14 := 0; fi14 < 8; fi14++ {
        var v15 uint64 = 0
        v16 := []int{1}
        _ = fi14
        _ = v15
        _ = v16
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 100
}
func genFn1(p0 uint8) uint {
    var v17 string = ("foo bar" + "hello")
    v18 := []int{9223372036854775807}
    fmt.Println(v17)
    _ = p0
    _ = v17
    _ = v18
    return 42
}

func main() {
    var v1 string = (("a" + "\n\t") + ("hello" + (("foo bar" + ("foo bar" + (("foo bar" + ("a" + "foo bar")) + (((("" + "hello") + "\n\t") + "a") + "hello")))) + "\n\t")))
    v2 := make(map[string]bool)
    var v3 int32 = 0
    var v4 string = "a"
    v5 := make(map[string]uint32)
    var v19 float64 = (0.1 / (((1.5 * (-2.25 * ((1.5 / 0.1) + 0.0))) * 0.0) + 1e10))
    switch 886 {
    case 42:
        switch ("" + v4) {
        case "a":
            if (!(true || true)) {
                v20 := GenS1{F0: 2, F1: -553}
                var v21 int16 = 32767
                _ = v20
                _ = v21
            }
        case "\n\t":
            for fi22 := 0; fi22 < 3; fi22++ {
                var v23 uint = 18446744073709551615
                _ = fi22
                _ = v23
            }
        }
    case -9223372036854775808:
        for fi24 := 0; fi24 < 20; fi24++ {
            if (3.14159 == (v19 - float64((((0.0 * -2.25) + 1.5) + 3.14159)))) {
                fmt.Println(v1)
                if ((((3.14159 - float32((0.1 - v19))) - float32((v19 + v19))) + float32((-2.25 - v19))) <= (1.5 - 0.1)) {
                    for fi25 := 0; fi25 < 5; fi25++ {
                        v26 := GenS1{F0: 42, F1: 405}
                        v3 = v3
                        if (228 > 108) {
                            v27 := v2[(v1 + (v4 + ("\n\t" + (v1 + "\n\t"))))]
                            v28 := []int{}
                            _ = v27
                            _ = v28
                        }
                        _ = fi25
                        _ = v26
                    }
                    v3 = (v3 | v3)
                    fmt.Println(fi24, fi24)
                } else {
                    var v29 string = (v4 + "")
                    v30 := GenS1{F0: 2, F1: (fi24 & fi24)}
                    _ = v29
                    _ = v30
                }
                v31 := GenS2{F0: 1, F1: 1}
                _ = v31
            } else {
                for fi32 := 0; fi32 < 13; fi32++ {
                    if (88 != -5) {
                        if true {
                            v33 := GenS3{F0: -312, F1: v3, F2: 2}
                            _ = v33
                        } else {
                            v34 := []int{}
                            _ = v34
                        }
                        var v35 float64 = ((v19 + (1e10 + 0.1)) / float64(1.5))
                        switch 32768 {
                        case -341:
                            v4 = v4
                        case 32767:
                            var v36 uint16 = 0
                            v37 := []int{fi32}
                            fmt.Println(v4)
                            _ = v36
                            _ = v37
                        }
                        _ = v35
                    } else {
                        var v38 uint64 = 2
                        _ = v38
                    }
                    _ = fi32
                }
            }
            fmt.Println(fi24, v3)
            var v39 int64 = 0
            _ = fi24
            _ = v39
        }
    default:
        var v40 uint32 = 583
        _ = v40
    }
    var v41 string = v4
    v19 = -2.25
    if ((v1 + v41) == v41) {
        switch ("" + (v1 + v1)) {
        case "a":
            for fi42 := 0; fi42 < 9; fi42++ {
                var v43 float32 = (3.14159 * float32(float64(((3.14159 * -2.25) * 1.5))))
                _ = fi42
                _ = v43
            }
        case "foo bar":
            for fi44 := 0; fi44 < 16; fi44++ {
                switch "foo bar" {
                case "a":
                    var v45 string = ((v4 + v1) + "\n\t")
                    _ = v45
                case "foo bar":
                    v41 = v4
                case "hello":
                    var v46 uint8 = 130
                    _ = v46
                }
                _ = fi44
            }
        case "hello":
            if (true && false) {
                if (-559 >= 101) {
                    switch 439 {
                    case -2147483648:
                        var v47 uint32 = 583
                        _ = v47
                    case 982:
                        switch "a" {
                        case "\n\t":
                            v41 = ((v4 + "a") + ("foo bar" + v1))
                        case "":
                            if false {
                                v41 = v4
                                var v48 uint64 = 4294967295
                                _ = v48
                            }
                        case "foo bar":
                            var v49 uint = 58
                            _ = v49
                        default:
                            if false {
                                fmt.Println(v4, v19, v19)
                                if (!(-32768 != -32768)) {
                                    for fi50 := 0; fi50 < 6; fi50++ {
                                        v51 := []int{(fi50 / ((-2147483648) | 1)), fi50, (-(fi50))}
                                        _ = fi50
                                        _ = v51
                                    }
                                    for fi52 := 0; fi52 < 9; fi52++ {
                                        var v53 float32 = 3.14159
                                        if true {
                                            var v54 float64 = v19
                                            for fi55 := 0; fi55 < 11; fi55++ {
                                                if (!true) {
                                                    if true {
                                                        fmt.Println(v41, v54)
                                                        v56 := []int{(fi55 - (fi52 - fi52)), (fi52 | fi52), (fi52 + (fi55 | -844)), fi52}
                                                        v3 = (v3 - 42)
                                                        _ = v56
                                                    } else {
                                                        switch (fi55 % (((fi52 >> ((422) & 63))) | 1)) {
                                                        case 42:
                                                            _ = len(v5)
                                                            v3 = (v3 % ((v3) | 1))
                                                            if false {
                                                                v57 := []int{-991}
                                                                v58 := v5[v4]
                                                                var v59 int = 32767
                                                                _ = v57
                                                                _ = v58
                                                                _ = v59
                                                            }
                                                        case -2147483648:
                                                            v53 = v53
                                                        }
                                                    }
                                                    v41 = "foo bar"
                                                } else {
                                                    var v60 uint = 0
                                                    v61 := []int{}
                                                    fmt.Println(v19)
                                                    _ = v60
                                                    _ = v61
                                                }
                                                var v62 uint64 = 528
                                                var v63 bool = (!(127 != -128))
                                                fmt.Println(v53, v41, v63)
                                                v5[v4] = 42
                                                switch (fi52 / ((fi52) | 1)) {
                                                case -9223372036854775808:
                                                    v19 += v19
                                                case -384:
                                                    v64 := GenS3{F0: fi52, F1: (v3 ^ (v3 + (^(v3)))), F2: (v62 << ((2) & 63))}
                                                    _ = v64
                                                case 127:
                                                    switch ((v1 + (("foo bar" + (v41 + v4)) + v4)) + (v1 + ("" + v41))) {
                                                    case "":
                                                        var v65 uint16 = 255
                                                        _ = v65
                                                    }
                                                }
                                                v66 := GenS4{F0: 127}
                                                _ = fi55
                                                _ = v62
                                                _ = v63
                                                _ = v66
                                            }
                                            v67 := GenS3{F0: (fi52 + 9223372036854775807), F1: -1, F2: 0}
                                            _ = v54
                                            _ = v67
                                        } else {
                                            fmt.Println(v4)
                                            v53 = v53
                                            fmt.Println(v4, fi52, v3)
                                        }
                                        v19 += float64(v53)
                                        _ = fi52
                                        _ = v53
                                    }
                                    fmt.Println(v4, v4)
                                } else {
                                    var v68 uint = 878
                                    for fi69 := 0; fi69 < 16; fi69++ {
                                        if false {
                                            v70 := GenS3{F0: (fi69 * fi69), F1: (v3 * 458), F2: 4294967295}
                                            _ = v70
                                        }
                                        if (!false) {
                                            v71 := GenS1{F0: 533, F1: (fi69 << ((348) & 63))}
                                            v3 = v3
                                            _ = v71
                                        }
                                        v41 = ""
                                        _ = fi69
                                    }
                                    _ = v68
                                }
                                var v72 int = 1
                                _ = v72
                            }
                        }
                    case -9223372036854775808:
                        var v73 uint16 = 65535
                        _ = v73
                    default:
                        var v74 int8 = 1
                        _ = v74
                    }
                    var v75 int = -434
                    v1 = v41
                    _ = v75
                } else {
                    v76 := GenS5{F0: v41}
                    _ = v76
                }
                if (0 == 127) {
                    switch 4294967295 {
                    case 254:
                        switch 32767 {
                        case 1:
                            switch (v4 + (("foo bar" + (v1 + "a")) + "foo bar")) {
                            case "\n\t":
                                var v77 float64 = v19
                                _ = v77
                            case "a":
                                var v78 int64 = -935
                                _ = v78
                            case "hello":
                                v3 -= (^(v3))
                            default:
                                v19 += 0.0
                            }
                        case 32768:
                            v79 := GenS6{F0: 2}
                            _ = v79
                        case 954:
                            var v80 float64 = float64(3.14159)
                            _ = v80
                        }
                    case -352:
                        v81 := GenS3{F0: -8, F1: -261, F2: 0}
                        _ = v81
                    case -2147483648:
                        v41 = v1
                    }
                    v82 := []int{-2147483648, 453, 32767, 128}
                    switch v4 {
                    case "foo bar":
                        v19 = v19
                    case "\n\t":
                        fmt.Println(v4)
                    case "":
                        var v83 uint64 = 4294967295
                        _ = v83
                    default:
                        v84 := GenS7{F0: 42, F1: ("foo bar" + (("foo bar" + v1) + v1)), F2: -622}
                        _ = v84
                    }
                    _ = v82
                } else {
                    fmt.Println(v1)
                    for fi85 := 0; fi85 < 16; fi85++ {
                        var v86 uint = 2
                        switch (v4 + v41) {
                        case "foo bar":
                            switch fi85 {
                            case -334:
                                var v87 int16 = 42
                                _ = v87
                            }
                        default:
                            for fi88 := 0; fi88 < 10; fi88++ {
                                switch 128 {
                                case -338:
                                    if ((1.5 < (float32((float64((3.14159 - -2.25)) - 0.0)) + float32(0.1))) || true) {
                                        var v89 int = 2147483647
                                        fmt.Println(v41, v19, v1)
                                        var v90 uint32 = 59
                                        _ = v89
                                        _ = v90
                                    } else {
                                        switch v1 {
                                        case "foo bar":
                                            v1 = (v4 + v4)
                                        case "":
                                            if (uint64(v86) >= 553) {
                                                var v91 float32 = (-2.25 - 3.14159)
                                                var v92 uint16 = 1
                                                _ = v91
                                                _ = v92
                                            } else {
                                                fmt.Println(v41, v86)
                                                fmt.Println(v19, v3)
                                                if (false || false) {
                                                    if true {
                                                        if ((true && (v19 > 1e10)) || (1000 < 958)) {
                                                            v86 -= v86
                                                            v93 := GenS8{F0: 89, F1: 184}
                                                            _ = v93
                                                        } else {
                                                            fmt.Println(v4, v4)
                                                            switch (((("\n\t" + "") + "a") + (v1 + (v4 + v1))) + ("hello" + "\n\t")) {
                                                            case "hello":
                                                                var v94 bool = true
                                                                _ = v94
                                                            default:
                                                                if false {
                                                                    if false {
                                                                        var v95 float64 = v19
                                                                        switch v1 {
                                                                        case "hello":
                                                                            fmt.Println(v41, v19, v3)
                                                                        case "":
                                                                            if false {
                                                                                v19 += v95
                                                                                v19 = v95
                                                                                switch fi85 {
                                                                                case 9223372036854775807:
                                                                                    if (!false) {
                                                                                        fmt.Println(fi88, fi88)
                                                                                        delete(v2, "hello")
                                                                                        var v96 uint8 = 169
                                                                                        _ = v96
                                                                                    } else {
                                                                                        delete(v5, (v4 + v1))
                                                                                        v97 := GenS1{F0: 0, F1: 127}
                                                                                        _ = v97
                                                                                    }
                                                                                }
                                                                                if true {
                                                                                    v86 -= (v86 - v86)
                                                                                }
                                                                            } else {
                                                                                fmt.Println(v19, v41, v19)
                                                                                var v98 float64 = v19
                                                                                _ = v98
                                                                            }
                                                                        }
                                                                        v99 := GenS7{F0: 42, F1: v41, F2: (fi88 / ((32767) | 1))}
                                                                        _ = v95
                                                                        _ = v99
                                                                    } else {
                                                                        v100 := []int{fi88}
                                                                        if ((3.14159 / float32(1.5)) >= 1.5) {
                                                                            v4 = v1
                                                                            v41 = "a"
                                                                        }
                                                                        var v101 float32 = -2.25
                                                                        var v102 float64 = ((0.0 - -2.25) - v19)
                                                                        _ = v100
                                                                        _ = v101
                                                                        _ = v102
                                                                    }
                                                                    v103 := v5["foo bar"]
                                                                    _ = v103
                                                                } else {
                                                                    v104 := GenS9{F0: 0, F1: v41, F2: 234}
                                                                    _ = v104
                                                                }
                                                            }
                                                        }
                                                        v105 := GenS10{F0: 1, F1: 2147483648}
                                                        _ = v105
                                                    }
                                                    fmt.Println(v1, v1, v1)
                                                } else {
                                                    v19 = v19
                                                    if ((v3 ^ (v3 % ((v3) | 1))) != -1) {
                                                        if (!(!(true || true))) {
                                                            v41 = ""
                                                            delete(v2, "\n\t")
                                                            var v106 uint16 = 731
                                                            _ = v106
                                                        } else {
                                                            switch 69 {
                                                            case 42:
                                                                v107 := GenS5{F0: v4}
                                                                _ = v107
                                                            case 127:
                                                                v19 -= 0.1
                                                                v108 := GenS11{F0: 0}
                                                                _ = v108
                                                            case -760:
                                                                if true {
                                                                    var v109 uint = (v86 / (((v86 & v86)) | 1))
                                                                    v110 := GenS3{F0: (fi88 >> (((v86 % (((v86 ^ v86)) | 1))) & 63)), F1: 1, F2: 29}
                                                                    if false {
                                                                        v111 := v2[(("a" + v41) + "")]
                                                                        var v112 int32 = (v3 << ((0) & 63))
                                                                        v3 -= v3
                                                                        _ = v111
                                                                        _ = v112
                                                                    }
                                                                    _ = v109
                                                                    _ = v110
                                                                }
                                                            }
                                                            delete(v5, v1)
                                                            var v113 int = fi88
                                                            _ = v113
                                                        }
                                                        v41 = "\n\t"
                                                    } else {
                                                        var v114 float64 = v19
                                                        _ = v114
                                                    }
                                                    switch v41 {
                                                    case "hello":
                                                        if (-1 == 32767) {
                                                            v1 = ("hello" + v4)
                                                            var v115 int32 = -50
                                                            switch (v1 + v41) {
                                                            case "foo bar":
                                                                fmt.Println(fi88)
                                                            case "\n\t":
                                                                var v116 int16 = 1
                                                                _ = v116
                                                            }
                                                            v5[v1] = 0
                                                            v1 = ((v1 + v4) + ("" + v1))
                                                            if (9223372036854775807 >= 425) {
                                                                switch "a" {
                                                                case "":
                                                                    v4 = ((v4 + v1) + "a")
                                                                case "foo bar":
                                                                    if ((v3 >> ((genFn1(127)) & 63)) >= (v3 / (((v115 & v115)) | 1))) {
                                                                        switch fi85 {
                                                                        case 2147483647:
                                                                            v41 = "hello"
                                                                        default:
                                                                            fmt.Println(fi88, v115)
                                                                        }
                                                                        if ((int8(fi88) > -1) && false) {
                                                                            v117 := []int{-791, (^(fi88))}
                                                                            if ((-(fi85)) >= fi88) {
                                                                                var v118 int = fi85
                                                                                var v119 string = ""
                                                                                var v120 uint8 = 0
                                                                                _ = len(v117)
                                                                                var v121 int8 = -128
                                                                                _ = v118
                                                                                _ = v119
                                                                                _ = v120
                                                                                _ = v121
                                                                            }
                                                                            if (!(!false)) {
                                                                                v122 := GenS12{F0: "foo bar"}
                                                                                v19 = (0.1 - 3.14159)
                                                                                _ = v122
                                                                            } else {
                                                                                v19 -= v19
                                                                                v123 := GenS8{F0: 255, F1: 42}
                                                                                _ = v123
                                                                            }
                                                                            _ = v117
                                                                        }
                                                                        var v124 uint32 = 2147483648
                                                                        v125 := GenS13{F0: 1, F1: 2147483647, F2: 276}
                                                                        _ = v124
                                                                        _ = v125
                                                                    }
                                                                case "hello":
                                                                    var v126 uint16 = 255
                                                                    _ = v126
                                                                default:
                                                                    fmt.Println(v19, fi88, v19)
                                                                }
                                                                if ((0.1 / 3.14159) <= float32(v19)) {
                                                                    v1 = v4
                                                                    var v127 uint32 = 152
                                                                    _ = v127
                                                                } else {
                                                                    v128 := GenS3{F0: (fi85 % ((-9223372036854775808) | 1)), F1: 1, F2: 554}
                                                                    v129 := v2[v41]
                                                                    v129 = (-9223372036854775808 == 0)
                                                                    _ = v128
                                                                    _ = v129
                                                                }
                                                                v41 = v1
                                                            }
                                                            v130 := v5["hello"]
                                                            var v131 float32 = 1e10
                                                            switch 188 {
                                                            case 40:
                                                                v4 = (v41 + v4)
                                                            }
                                                            if false {
                                                                if (genFn0(1, (v86 >> ((556) & 63)), 0) >= int8(v115)) {
                                                                    v132 := GenS11{F0: 1}
                                                                    var v133 float64 = v19
                                                                    if true {
                                                                        if false {
                                                                            v133 -= float64((v131 - v131))
                                                                            v131 = v131
                                                                        } else {
                                                                            fmt.Println(fi85, v86)
                                                                        }
                                                                        var v134 uint32 = (^(v130))
                                                                        v135 := []int{}
                                                                        _ = v134
                                                                        _ = v135
                                                                    }
                                                                    v136 := []int{(-(fi85))}
                                                                    var v137 int32 = v115
                                                                    _ = v132
                                                                    _ = v133
                                                                    _ = v136
                                                                    _ = v137
                                                                }
                                                            }
                                                            if ((((v131 * v131) * -2.25) + ((float32(0.0) - float32((1.5 * v19))) / 1.0)) == v131) {
                                                                v131 -= (3.14159 + v131)
                                                                switch "hello" {
                                                                case "foo bar":
                                                                    delete(v5, "foo bar")
                                                                case "\n\t":
                                                                    var v138 uint64 = 539
                                                                    _ = v138
                                                                }
                                                                if ((int64(v86) <= -515) || true) {
                                                                    v41 = v41
                                                                }
                                                            }
                                                            _ = v115
                                                            _ = v130
                                                            _ = v131
                                                        }
                                                        delete(v2, ("foo bar" + v1))
                                                    }
                                                }
                                                fmt.Println(v4, v3)
                                            }
                                        case "hello":
                                            v19 -= 3.14159
                                            var v139 float64 = (v19 / 1e10)
                                            var v140 bool = (true || true)
                                            _ = v139
                                            _ = v140
                                        }
                                    }
                                case 32768:
                                    switch (fi88 % ((fi88) | 1)) {
                                    case 128:
                                        if true {
                                            v141 := []int{}
                                            _ = v141
                                        }
                                    case 127:
                                        var v142 int64 = -1
                                        v142 = 116
                                        fmt.Println(fi85, fi88, v3)
                                        switch v1 {
                                        case "a":
                                            fmt.Println(fi88, v19, v86)
                                            v143 := GenS7{F0: 4294967295, F1: "\n\t", F2: fi88}
                                            v4 = "hello"
                                            v144 := GenS14{F0: -128, F1: (v86 << (((v86 * (v86 - (v86 << (((^(v86))) & 63))))) & 63))}
                                            fmt.Println(v3, fi88, v3)
                                            var v145 uint32 = 0
                                            _ = v143
                                            _ = v144
                                            _ = v145
                                        case "":
                                            switch fi85 {
                                            case 32767:
                                                v4 = "a"
                                            case 42:
                                                fmt.Println(v4, fi88, v142)
                                            default:
                                                var v146 int8 = 81
                                                _ = v146
                                            }
                                        case "hello":
                                            if false {
                                                v147 := GenS4{F0: 0}
                                                v148 := GenS10{F0: 0, F1: 0}
                                                var v149 uint = v86
                                                _ = v147
                                                _ = v148
                                                _ = v149
                                            }
                                        default:
                                            var v150 string = ((("a" + v41) + v1) + (v41 + (((v4 + v1) + v4) + ("\n\t" + (("a" + ("foo bar" + "")) + v1)))))
                                            var v151 uint16 = 820
                                            _ = v150
                                            _ = v151
                                        }
                                        if ((677 <= 932) && false) {
                                            v152 := []int{(fi88 ^ fi85), (fi85 << (((v86 * genFn1(180))) & 63))}
                                            if false {
                                                v152 = append(v152, fi85)
                                                if (v142 >= (v142 ^ -1)) {
                                                    v153 := GenS15{F0: 42}
                                                    var v154 bool = (true && (fi88 < fi85))
                                                    if v154 {
                                                        delete(v5, v1)
                                                        var v155 uint64 = 42
                                                        _ = v155
                                                    } else {
                                                        if len(v152) >= 2 {
                                                            v152 = v152[1:len(v152)]
                                                        }
                                                        var v156 bool = v154
                                                        _ = v156
                                                    }
                                                    _ = v153
                                                    _ = v154
                                                } else {
                                                    v157 := v2["a"]
                                                    v158 := v5[((v1 + v4) + (v41 + ((v41 + ("" + v4)) + (("a" + v1) + (((("hello" + v4) + "a") + v1) + ((v41 + v4) + v4))))))]
                                                    _ = v157
                                                    _ = v158
                                                }
                                                v3 = v3
                                            } else {
                                                v159 := GenS6{F0: 42}
                                                if true {
                                                    var v160 uint32 = uint32(v86)
                                                    _ = v160
                                                } else {
                                                    var v161 int16 = 1000
                                                    var v162 float32 = -2.25
                                                    _ = v161
                                                    _ = v162
                                                }
                                                _ = v159
                                            }
                                            _ = v152
                                        } else {
                                            v163 := []int{(fi88 << ((v86) & 63)), 1, fi85, (fi85 - 0)}
                                            switch ("\n\t" + ((v1 + ("\n\t" + v4)) + v4)) {
                                            case "":
                                                if true {
                                                    var v164 int8 = 1
                                                    if (1e10 != ((((1.5 * 0.1) * ((1.5 + float32(v19)) - 0.0)) * 1e10) / 3.14159)) {
                                                        var v165 uint = 18446744073709551615
                                                        switch (-(fi85)) {
                                                        case 0:
                                                            var v166 float64 = v19
                                                            var v167 uint16 = 32767
                                                            _ = v166
                                                            _ = v167
                                                        default:
                                                            v168 := GenS16{F0: 32767}
                                                            _ = v168
                                                        }
                                                        var v169 int8 = (-(v164))
                                                        _ = v165
                                                        _ = v169
                                                    } else {
                                                        if (false && false) {
                                                            v41 = v1
                                                            if false {
                                                                v164 = genFn0(1, v86, 32767)
                                                                fmt.Println(v3, fi88, v41)
                                                                if len(v163) > 0 {
                                                                    v163[0] = fi88
                                                                }
                                                            }
                                                            var v170 uint8 = 42
                                                            fmt.Println(v19)
                                                            _ = v170
                                                        }
                                                        v163 = append(v163, (fi88 << ((genFn1(255)) & 63)))
                                                    }
                                                    _ = v164
                                                } else {
                                                    fmt.Println(fi88)
                                                }
                                            default:
                                                v1 = "foo bar"
                                            }
                                            _ = v163
                                        }
                                        v171 := []int{fi85, fi85, 0, fi85}
                                        _ = v142
                                        _ = v171
                                    case 32767:
                                        delete(v5, v4)
                                        switch fi88 {
                                        case -482:
                                            v172 := []int{}
                                            _ = v172
                                        }
                                    }
                                case 209:
                                    v173 := []int{(fi88 | (fi85 << ((v86) & 63))), (^(fi85))}
                                    _ = v173
                                default:
                                    v1 = "foo bar"
                                }
                                v174 := []int{}
                                _ = fi88
                                _ = v174
                            }
                        }
                        v175 := GenS3{F0: fi85, F1: v3, F2: 18446744073709551615}
                        _ = fi85
                        _ = v175
                        _ = v86
                    }
                }
                fmt.Println(v4, v3)
            } else {
                var v176 int16 = 707
                v177 := []int{-292, 232, 0, 1}
                switch -1 {
                case 1:
                    for fi178 := 0; fi178 < 18; fi178++ {
                        v179 := GenS13{F0: 2, F1: 2147483647, F2: 331}
                        if len(v177) > 0 {
                            v177[0] = 32767
                        }
                        delete(v2, ((v4 + v4) + v41))
                        _ = fi178
                        _ = v179
                    }
                case -488:
                    for fi180 := 0; fi180 < 6; fi180++ {
                        v181 := GenS1{F0: 771, F1: fi180}
                        _ = fi180
                        _ = v181
                    }
                case 9223372036854775807:
                    fmt.Println(v4, v19, v41)
                }
                _ = v176
                _ = v177
            }
        }
        v2[v1] = false
    }
    for fi182 := 0; fi182 < 15; fi182++ {
        v3 -= v3
        v41 = ""
        _ = fi182
    }
    v183 := GenS6{F0: 32767}
    var v184 float64 = v19
    var v185 int32 = (v3 >> ((18446744073709551615) & 63))
    fmt.Println(v3, v3)
    if false {
        v185 = 901
        var v186 float64 = float64(((3.14159 * -2.25) * (-2.25 - -2.25)))
        _ = v186
    }
    var v187 bool = false
    if (false && (!v187)) {
        fmt.Println(v19)
        var v188 uint8 = 42
        v41 = "hello"
        _ = v188
    } else {
        if false {
            switch "\n\t" {
            case "hello":
                v189 := []int{32768}
                _ = v189
            }
            fmt.Println(v19, v41)
        }
    }
    v3 -= (v185 | v185)
    for fi190 := 0; fi190 < 9; fi190++ {
        switch (fi190 ^ fi190) {
        case 1:
            switch -9223372036854775808 {
            case -1:
                switch "hello" {
                case "foo bar":
                    for fi191 := 0; fi191 < 19; fi191++ {
                        fmt.Println(fi191)
                        var v192 uint = 280
                        v185 = v185
                        _ = fi191
                        _ = v192
                    }
                default:
                    v193 := []int{fi190, 4294967295}
                    _ = v193
                }
            default:
                fmt.Println(v1)
            }
        }
        var v194 uint8 = 128
        _ = fi190
        _ = v194
    }
    v187 = (0 != -104)
    fmt.Println(v187, v185)
    fmt.Println("v1:", v1)
    fmt.Println("v184:", v184)
    fmt.Println("v185:", v185)
    fmt.Println("v187:", v187)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v183
    _ = v184
    _ = v185
    _ = v187
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v5
}
