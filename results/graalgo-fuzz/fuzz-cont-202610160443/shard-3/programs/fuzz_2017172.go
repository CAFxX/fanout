package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string; F2 int64 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint; F1 int16 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int64; F1 bool; F2 int }
type GenS6 struct { F0 int64; F1 int8 }
type GenS7 struct { F0 uint; F1 uint8; F2 uint32 }
type GenS8 struct { F0 uint32 }
type GenS9 struct { F0 int16; F1 string }
type GenS10 struct { F0 uint32; F1 int64 }
type GenS11 struct { F0 uint64; F1 uint16 }
type GenS12 struct { F0 string; F1 uint8; F2 uint }
type GenS13 struct { F0 int32 }
type GenS14 struct { F0 string }
type GenS15 struct { F0 int32; F1 bool; F2 uint8 }
type GenS16 struct { F0 uint8 }
type GenS17 struct { F0 int64; F1 int8; F2 bool }
type GenS18 struct { F0 uint8; F1 bool; F2 int16 }
type GenS19 struct { F0 int8; F1 int64; F2 uint16 }
type GenS20 struct { F0 int8 }
type GenS21 struct { F0 string; F1 int8 }
type GenS22 struct { F0 int16; F1 string; F2 int32 }
type GenS23 struct { F0 uint }
type GenS24 struct { F0 string; F1 string; F2 uint32 }
func genFn0(p0 int16, p1 uint64) string {
    p0 = p0
    switch ("a" + ("a" + "foo bar")) {
    case "foo bar":
        switch "\n\t" {
        case "hello":
            for fi4 := 0; fi4 < 7; fi4++ {
                var v5 uint64 = 18446744073709551615
                var v6 float64 = 0.0
                v6 = (v6 / float64(0.1))
                _ = fi4
                _ = v5
                _ = v6
            }
        case "foo bar":
            for fi7 := 0; fi7 < 15; fi7++ {
                p0 = p0
                v8 := GenS1{F0: (fi7 << ((18446744073709551615) & 63)), F1: "hello", F2: 392}
                _ = fi7
                _ = v8
            }
        case "a":
            if false {
                v9 := []int{}
                var v10 uint64 = (p1 + p1)
                if (-2.25 != ((float32((1e10 * 1e10)) / (-2.25 / 3.14159)) * 0.0)) {
                    p0 = p0
                    var v11 string = ("foo bar" + "")
                    if len(v9) >= 2 {
                        v9 = v9[1:len(v9)]
                    }
                    _ = v11
                }
                _ = v10
                _ = v9
            } else {
                v12 := []int{}
                for fi13 := 0; fi13 < 3; fi13++ {
                    var v14 int64 = -117
                    p0 -= -32768
                    v15 := GenS1{F0: (fi13 ^ (fi13 & 32767)), F1: "a", F2: v14}
                    _ = fi13
                    _ = v14
                    _ = v15
                }
                var v16 bool = (!(196 > 179))
                _ = v12
                _ = v16
            }
        default:
            var v17 float64 = 1.5
            _ = v17
        }
    }
    if true {
        fmt.Println(p0)
        if ((false || (p0 == (p0 | p0))) && ((8 != 0) || false)) {
            p1 = 0
            if (false && ((p0 >= p0) && false)) {
                if true {
                    p1 = p1
                    var v18 uint32 = 2147483648
                    _ = v18
                } else {
                    var v19 int32 = -366
                    _ = v19
                }
                v20 := []int{}
                _ = len(v20)
                _ = v20
            }
        } else {
            fmt.Println(p0, p1)
            for fi21 := 0; fi21 < 9; fi21++ {
                fmt.Println(fi21)
                switch (fi21 / ((fi21) | 1)) {
                case -1:
                    if (true && ((1e10 + (1e10 / 1.0)) < 3.14159)) {
                        var v22 float32 = 1e10
                        var v23 float64 = 3.14159
                        _ = v22
                        _ = v23
                    }
                case 283:
                    var v24 bool = true
                    _ = v24
                case 32767:
                    p1 += p1
                default:
                    p0 += (p0 - p0)
                }
                if false {
                    p1 -= (^(p1))
                    switch (fi21 % (((fi21 ^ (fi21 % (((fi21 * (fi21 ^ fi21))) | 1)))) | 1)) {
                    case -693:
                        p1 += p1
                    case 127:
                        v25 := GenS2{F0: 4294967295}
                        _ = v25
                    default:
                        var v26 uint = 500
                        _ = v26
                    }
                } else {
                    var v27 float64 = (((0.1 + 1e10) + (0.0 * 3.14159)) - 0.0)
                    _ = v27
                }
                _ = fi21
            }
        }
    } else {
        p0 -= -1
    }
    _ = p0
    _ = p1
    return "foo bar"
}

func main() {
    var v1 uint = 4294967295
    var v2 int8 = 42
    var v3 uint = (v1 >> (((v1 | v1)) & 63))
    fmt.Println(v3, v2)
    fmt.Println(v3, v2, v3)
    fmt.Println(v2, v3)
    for fi28 := 0; fi28 < 17; fi28++ {
        var v29 float64 = 3.14159
        for fi30 := 0; fi30 < 19; fi30++ {
            var v31 int8 = v2
            if (0 == 1000) {
                var v32 float32 = 0.1
                var v33 int32 = 0
                v34 := []int{617, 127, fi30, 127}
                _ = v32
                _ = v33
                _ = v34
            }
            var v35 int8 = v31
            v1 = (v3 ^ (v1 << ((v1) & 63)))
            var v36 int32 = 42
            _ = fi30
            _ = v31
            _ = v35
            _ = v36
        }
        _ = fi28
        _ = v29
    }
    v37 := []int{}
    var v38 int32 = 1
    for fi39 := 0; fi39 < 11; fi39++ {
        fmt.Println(fi39)
        _ = fi39
    }
    if (false && false) {
        if len(v37) > 0 {
            v37[0] = 32767
        }
        var v40 uint16 = 255
        _ = v40
    } else {
        if false {
            if true {
                v1 -= (v3 % (((v1 + v1)) | 1))
                v38 = v38
                if (-859 == -476) {
                    var v41 float64 = 1e10
                    var v42 int8 = -104
                    _ = v41
                    _ = v42
                } else {
                    fmt.Println(v38, v3)
                }
            }
        } else {
            fmt.Println(v3, v3)
            v38 += 2147483647
            if (!false) {
                fmt.Println(v38, v2, v38)
            }
        }
        var v43 int16 = 42
        _ = v43
    }
    var v44 int8 = (v2 * (^(v2)))
    for fi45 := 0; fi45 < 11; fi45++ {
        if false {
            v1 = v3
            _ = len(v37)
            v46 := GenS3{F0: v1, F1: 32767}
            _ = v46
        } else {
            var v47 int16 = 289
            var v48 uint = 65535
            if (!(4294967295 > 775)) {
                v44 = (v44 ^ (v2 << ((1) & 63)))
                for fi49 := 0; fi49 < 8; fi49++ {
                    v1 += (^(v48))
                    switch fi45 {
                    case 2147483647:
                        v47 = (v47 - 42)
                    case 9223372036854775807:
                        fmt.Println(fi45, fi45)
                    case 42:
                        if len(v37) >= 2 {
                            v37 = v37[1:len(v37)]
                        }
                    default:
                        fmt.Println(v38)
                    }
                    _ = fi49
                }
            } else {
                v44 -= v44
            }
            _ = v47
            _ = v48
        }
        _ = fi45
    }
    var v50 int64 = -1
    fmt.Println(v3)
    fmt.Println(v50, v1, v50)
    v51 := GenS3{F0: (v1 / ((287) | 1)), F1: 32767}
    v3 = v3
    for fi52 := 0; fi52 < 9; fi52++ {
        var v53 int = (fi52 + (fi52 - fi52))
        fmt.Println(v3)
        _ = fi52
        _ = v53
    }
    if len(v37) > 0 {
        v37[0] = 1
    }
    if false {
        for fi54 := 0; fi54 < 4; fi54++ {
            for fi55 := 0; fi55 < 14; fi55++ {
                var v56 float32 = float32((1e10 + 1.5))
                v57 := GenS2{F0: 993}
                _ = fi55
                _ = v56
                _ = v57
            }
            v2 = (v44 << (((^(v1))) & 63))
            if true {
                var v58 float32 = (((3.14159 * 1e10) + ((0.0 / float32(-2.25)) + 0.1)) / (-2.25 * 0.1))
                switch ("a" + "") {
                case "":
                    switch (fi54 - fi54) {
                    case 42:
                        for fi59 := 0; fi59 < 15; fi59++ {
                            if true {
                                if true {
                                    v44 = -128
                                    if true {
                                        v37 = append(v37, fi54)
                                        var v60 int32 = v38
                                        _ = v60
                                    }
                                    var v61 uint = v3
                                    _ = v61
                                }
                                fmt.Println(v44, v44, v2)
                            } else {
                                v37 = append(v37, fi59)
                                if ((v44 * v2) != v44) {
                                    v62 := GenS4{F0: (4294967295 == 262)}
                                    var v63 int32 = (v38 + 2147483647)
                                    fmt.Println(v50)
                                    _ = v62
                                    _ = v63
                                } else {
                                    v64 := GenS5{F0: v50, F1: (true && (true && ((-(v44)) <= (v44 + (v44 << (((^(v1))) & 63)))))), F2: (fi59 + fi54)}
                                    fmt.Println(v3, v2, v1)
                                    var v65 float64 = -2.25
                                    _ = v64
                                    _ = v65
                                }
                                fmt.Println(v44)
                                switch fi54 {
                                case 32767:
                                    v66 := GenS1{F0: 0, F1: ("hello" + ("a" + "foo bar")), F2: 1}
                                    _ = v66
                                case 0:
                                    switch "\n\t" {
                                    case "foo bar":
                                        v58 -= -2.25
                                    case "a":
                                        v1 -= (v3 / (((^(v1))) | 1))
                                    case "hello":
                                        switch "a" {
                                        case "":
                                            v2 -= v44
                                        default:
                                            v67 := GenS1{F0: -694, F1: ((("\n\t" + "hello") + "") + ""), F2: v50}
                                            _ = v67
                                        }
                                    }
                                case 42:
                                    if true {
                                        switch (fi59 + (fi54 % ((-1) | 1))) {
                                        case 128:
                                            var v68 uint = 2
                                            _ = v68
                                        case -478:
                                            if true {
                                                var v69 uint = 65535
                                                if true {
                                                    switch (fi59 + 4294967295) {
                                                    case 4294967295:
                                                        var v70 int = fi54
                                                        _ = v70
                                                    case -2147483648:
                                                        switch 1 {
                                                        case -1:
                                                            var v71 float32 = (v58 / (v58 - -2.25))
                                                            _ = v71
                                                        }
                                                    }
                                                } else {
                                                    v38 -= (v38 - 2147483647)
                                                    fmt.Println(v38)
                                                    if true {
                                                        switch fi59 {
                                                        case -9223372036854775808:
                                                            var v72 int64 = (v50 ^ v50)
                                                            if len(v37) > 0 {
                                                                v37[0] = -9223372036854775808
                                                            }
                                                            switch "foo bar" {
                                                            case "":
                                                                fmt.Println(v72, v3)
                                                            case "hello":
                                                                _ = len(v37)
                                                            case "\n\t":
                                                                v2 = -1
                                                            }
                                                            _ = v72
                                                        default:
                                                            v38 += 2147483647
                                                        }
                                                        switch (fi59 >> ((v3) & 63)) {
                                                        case -1:
                                                            v44 = v44
                                                        }
                                                        v73 := GenS1{F0: fi59, F1: ("" + ""), F2: (-(v50))}
                                                        _ = v73
                                                    }
                                                }
                                                switch fi54 {
                                                case 342:
                                                    if true {
                                                        fmt.Println(fi54, v50, v1)
                                                        var v74 int32 = (v38 | (v38 & (v38 >> (((v3 << ((4294967295) & 63))) & 63))))
                                                        switch (fi54 + 1) {
                                                        case 9223372036854775807:
                                                            v75 := GenS6{F0: (^(v50)), F1: (v2 ^ v2)}
                                                            _ = v75
                                                        }
                                                        _ = v74
                                                    }
                                                case -9223372036854775808:
                                                    _ = len(v37)
                                                case -2147483648:
                                                    if true {
                                                        v3 += (v1 | v69)
                                                        v76 := GenS7{F0: v1, F1: 42, F2: 2147483648}
                                                        var v77 uint64 = 1
                                                        _ = v76
                                                        _ = v77
                                                    } else {
                                                        var v78 uint16 = 65535
                                                        _ = v78
                                                    }
                                                }
                                                _ = v69
                                            } else {
                                                v50 = v50
                                                var v79 uint64 = 320
                                                switch (fi59 - (fi59 / ((-224) | 1))) {
                                                case 0:
                                                    if (-1 > v50) {
                                                        fmt.Println(fi54, v1, fi59)
                                                    }
                                                }
                                                _ = v79
                                            }
                                        case -2147483648:
                                            if len(v37) >= 2 {
                                                v37 = v37[1:len(v37)]
                                            }
                                            if true {
                                                if (!false) {
                                                    var v80 int = (fi59 & -1)
                                                    _ = v80
                                                } else {
                                                    var v81 uint = v3
                                                    if true {
                                                        if (1e10 > 0.1) {
                                                            if false {
                                                                if len(v37) >= 2 {
                                                                    v37 = v37[1:len(v37)]
                                                                }
                                                                var v82 float64 = float64(v58)
                                                                _ = v82
                                                            }
                                                            switch "hello" {
                                                            case "foo bar":
                                                                v83 := GenS7{F0: v81, F1: 34, F2: 65535}
                                                                _ = v83
                                                            case "hello":
                                                                v84 := GenS8{F0: 2}
                                                                _ = v84
                                                            case "":
                                                                var v85 uint = v3
                                                                _ = v85
                                                            default:
                                                                if (v1 >= 255) {
                                                                    v37 = append(v37, fi54)
                                                                    var v86 float64 = (((0.0 / 0.1) + 3.14159) / 1.0)
                                                                    if ((v58 - (3.14159 / v58)) > ((v58 / 1.0) * (v58 * ((v58 * (0.0 - -2.25)) - (0.0 + v58))))) {
                                                                        fmt.Println(v50, v86, v81)
                                                                        v87 := GenS9{F0: int16(v38), F1: "\n\t"}
                                                                        _ = v87
                                                                    }
                                                                    _ = v86
                                                                } else {
                                                                    switch fi59 {
                                                                    case -179:
                                                                        if (255 < 411) {
                                                                            var v88 uint32 = 0
                                                                            var v89 uint = v1
                                                                            v90 := GenS10{F0: v88, F1: (v50 >> (((v1 % ((v3) | 1))) & 63))}
                                                                            _ = v88
                                                                            _ = v89
                                                                            _ = v90
                                                                        }
                                                                    }
                                                                    _ = len(v37)
                                                                }
                                                            }
                                                        } else {
                                                            v91 := GenS1{F0: (fi54 & fi59), F1: "hello", F2: v50}
                                                            v58 = (1e10 * float32(-2.25))
                                                            v92 := GenS11{F0: 280, F1: 32767}
                                                            _ = v91
                                                            _ = v92
                                                        }
                                                        v93 := GenS10{F0: 42, F1: (v50 << ((v1) & 63))}
                                                        v44 = (v44 | (v2 - 0))
                                                        _ = v93
                                                    }
                                                    v2 = v44
                                                    _ = v81
                                                }
                                                if false {
                                                    var v94 uint = v1
                                                    _ = v94
                                                }
                                                if false {
                                                    fmt.Println(v58, fi59)
                                                } else {
                                                    v1 -= v3
                                                    var v95 int16 = int16(v2)
                                                    if (2 <= 42) {
                                                        var v96 float32 = v58
                                                        if true {
                                                            if len(v37) >= 2 {
                                                                v37 = v37[1:len(v37)]
                                                            }
                                                            v96 += v96
                                                            fmt.Println(v3)
                                                        } else {
                                                            var v97 int32 = v38
                                                            _ = v97
                                                        }
                                                        if (false && (255 <= 2)) {
                                                            if len(v37) >= 2 {
                                                                v37 = v37[1:len(v37)]
                                                            }
                                                            v98 := GenS12{F0: "hello", F1: 178, F2: 1}
                                                            _ = v98
                                                        }
                                                        _ = v96
                                                    } else {
                                                        var v99 uint8 = 0
                                                        var v100 bool = (true && (v44 > v2))
                                                        if (2147483648 < 863) {
                                                            v1 = v3
                                                            fmt.Println(v100, v2, v95)
                                                            fmt.Println(v95, fi54)
                                                        }
                                                        switch (("foo bar" + (("hello" + ("" + (("foo bar" + "foo bar") + "a"))) + "")) + (("foo bar" + "foo bar") + ("hello" + ("a" + "hello")))) {
                                                        case "":
                                                            v50 = v50
                                                        case "foo bar":
                                                            var v101 int8 = (v44 * 42)
                                                            _ = v101
                                                        case "hello":
                                                            var v102 int16 = int16(v50)
                                                            _ = v102
                                                        default:
                                                            v3 += 65535
                                                        }
                                                        v99 = (v99 - v99)
                                                        var v103 string = "\n\t"
                                                        _ = v100
                                                        _ = v103
                                                        _ = v99
                                                    }
                                                    if len(v37) >= 2 {
                                                        v37 = v37[1:len(v37)]
                                                    }
                                                    v104 := GenS6{F0: int64(v38), F1: 0}
                                                    _ = v104
                                                    _ = v95
                                                }
                                            } else {
                                                v38 = v38
                                                fmt.Println(v3)
                                                switch ("hello" + ((("hello" + "\n\t") + "foo bar") + "hello")) {
                                                case "":
                                                    var v105 float32 = v58
                                                    _ = v105
                                                case "foo bar":
                                                    v58 += (0.1 - (v58 - v58))
                                                }
                                                fmt.Println(fi54, fi54)
                                            }
                                            switch (fi54 & fi59) {
                                            case 42:
                                                v106 := GenS5{F0: (-(v50)), F1: false, F2: fi54}
                                                var v107 uint8 = 2
                                                _ = v106
                                                _ = v107
                                            case 32768:
                                                var v108 int32 = (-(v38))
                                                _ = v108
                                            case 836:
                                                var v109 int32 = (v38 >> ((42) & 63))
                                                _ = v109
                                            default:
                                                var v110 bool = (0 <= 0)
                                                _ = v110
                                            }
                                        }
                                    } else {
                                        v50 += (v50 ^ v50)
                                    }
                                default:
                                    if true {
                                        v111 := GenS13{F0: (v38 + v38)}
                                        if false {
                                            if len(v37) > 0 {
                                                v37[0] = fi59
                                            }
                                            if len(v37) >= 2 {
                                                v37 = v37[1:len(v37)]
                                            }
                                            _ = len(v37)
                                        } else {
                                            v112 := GenS10{F0: 42, F1: v50}
                                            _ = v112
                                        }
                                        v113 := GenS14{F0: "hello"}
                                        _ = v111
                                        _ = v113
                                    } else {
                                        var v114 uint64 = 2
                                        v115 := GenS2{F0: 4294967295}
                                        _ = v114
                                        _ = v115
                                    }
                                }
                                fmt.Println(fi54, v58, v50)
                                fmt.Println(v50, v50, v2)
                            }
                            _ = fi59
                        }
                    default:
                        fmt.Println(fi54, v3)
                    }
                default:
                    for fi116 := 0; fi116 < 14; fi116++ {
                        v117 := GenS15{F0: v38, F1: false, F2: 1}
                        v2 += v44
                        _ = fi116
                        _ = v117
                    }
                }
                if (true || true) {
                    if ((((222 == 127) || (!true)) && (true || true)) || true) {
                        if len(v37) >= 2 {
                            v37 = v37[1:len(v37)]
                        }
                        var v118 uint16 = 943
                        _ = v118
                    }
                    var v119 uint32 = 357
                    if true {
                        switch (fi54 << (((v1 / ((738) | 1))) & 63)) {
                        case 32768:
                            fmt.Println(v1, v38)
                        case 32767:
                            if len(v37) >= 2 {
                                v37 = v37[1:len(v37)]
                            }
                        case 147:
                            if false {
                                var v120 int16 = 1000
                                if true {
                                    if true {
                                        if (v58 > (0.1 * ((float32(3.14159) * 1e10) + v58))) {
                                            v2 -= (v2 + -71)
                                            for fi121 := 0; fi121 < 9; fi121++ {
                                                switch "a" {
                                                case "\n\t":
                                                    switch (fi121 - (fi121 | fi54)) {
                                                    case 32768:
                                                        if true {
                                                            if (!(!true)) {
                                                                v122 := GenS8{F0: (^(v119))}
                                                                _ = v122
                                                            } else {
                                                                var v123 int64 = v50
                                                                if len(v37) > 0 {
                                                                    v37[0] = (fi54 + 1)
                                                                }
                                                                v124 := GenS8{F0: (v119 >> (((v3 + v1)) & 63))}
                                                                var v125 uint8 = 1
                                                                if (3.14159 >= -2.25) {
                                                                    v58 = (0.1 + -2.25)
                                                                    if len(v37) >= 2 {
                                                                        v37 = v37[1:len(v37)]
                                                                    }
                                                                    v58 = v58
                                                                    var v126 string = ""
                                                                    v127 := GenS16{F0: 128}
                                                                    _ = v126
                                                                    _ = v127
                                                                } else {
                                                                    v37 = append(v37, (fi121 ^ fi54))
                                                                    v1 = v1
                                                                    var v128 uint32 = v119
                                                                    _ = v128
                                                                }
                                                                _ = v123
                                                                _ = v124
                                                                _ = v125
                                                            }
                                                            v129 := GenS17{F0: v50, F1: (-(v2)), F2: true}
                                                            v130 := GenS18{F0: 42, F1: ((v119 - uint32(v1)) == (^(v119))), F2: (v120 & v120)}
                                                            v50 += v50
                                                            _ = v129
                                                            _ = v130
                                                        } else {
                                                            fmt.Println(v50, v58, v44)
                                                            if len(v37) >= 2 {
                                                                v37 = v37[1:len(v37)]
                                                            }
                                                            var v131 int16 = (v120 / (((v120 / (((^(v120))) | 1))) | 1))
                                                            _ = v131
                                                        }
                                                    default:
                                                        var v132 float32 = float32(0.1)
                                                        _ = v132
                                                    }
                                                default:
                                                    var v133 string = "a"
                                                    _ = v133
                                                }
                                                fmt.Println(v3)
                                                if false {
                                                    v44 = 100
                                                } else {
                                                    var v134 int32 = v38
                                                    v135 := GenS19{F0: (v2 | -108), F1: 243, F2: 0}
                                                    v136 := GenS20{F0: v44}
                                                    _ = v134
                                                    _ = v135
                                                    _ = v136
                                                }
                                                fmt.Println(v50, v2)
                                                _ = fi121
                                            }
                                        }
                                        v137 := GenS21{F0: ("a" + "\n\t"), F1: v2}
                                        for fi138 := 0; fi138 < 9; fi138++ {
                                            if true {
                                                if true {
                                                    if len(v37) > 0 {
                                                        v37[0] = fi54
                                                    }
                                                    v139 := GenS9{F0: 645, F1: ("foo bar" + "foo bar")}
                                                    _ = v139
                                                }
                                                if false {
                                                    var v140 uint16 = 0
                                                    v141 := GenS6{F0: (v50 - v50), F1: v2}
                                                    _ = v140
                                                    _ = v141
                                                }
                                                var v142 bool = (!true)
                                                _ = v142
                                            }
                                            switch fi138 {
                                            case -14:
                                                v143 := GenS19{F0: v2, F1: 1, F2: 65535}
                                                _ = v143
                                            case -2147483648:
                                                var v144 uint8 = 30
                                                _ = v144
                                            case 1:
                                                if false {
                                                    fmt.Println(v50, v44)
                                                    _ = len(v37)
                                                } else {
                                                    switch (((("hello" + "") + "\n\t") + "a") + "") {
                                                    case "a":
                                                        if len(v37) > 0 {
                                                            v37[0] = (fi54 << (((v3 % ((594) | 1))) & 63))
                                                        }
                                                    default:
                                                        fmt.Println(fi138, v38)
                                                    }
                                                    switch (fi138 + -65) {
                                                    case 128:
                                                        _ = len(v37)
                                                    case -9223372036854775808:
                                                        var v145 uint16 = 42
                                                        _ = v145
                                                    }
                                                    v3 = 383
                                                }
                                            }
                                            fmt.Println(v58)
                                            _ = fi138
                                        }
                                        _ = v137
                                    }
                                    v120 += (v120 | 32767)
                                    switch (("foo bar" + "foo bar") + "\n\t") {
                                    case "a":
                                        v44 = (v44 << ((v3) & 63))
                                    case "":
                                        v37 = append(v37, (fi54 << (((v1 ^ v1)) & 63)))
                                    case "hello":
                                        for fi146 := 0; fi146 < 20; fi146++ {
                                            var v147 bool = (!((("\n\t" + "") + "hello") <= "foo bar"))
                                            if v147 {
                                                _ = len(v37)
                                                v148 := GenS22{F0: (v120 & -1), F1: ("" + "\n\t"), F2: 1}
                                                if (42 != 4294967295) {
                                                    v44 = v2
                                                } else {
                                                    v149 := GenS23{F0: (v1 % (((v3 + v3)) | 1))}
                                                    var v150 float64 = 0.1
                                                    _ = v149
                                                    _ = v150
                                                }
                                                _ = v148
                                            }
                                            v151 := GenS19{F0: (v44 - (v44 | v44)), F1: (v50 ^ (^(v50))), F2: 2}
                                            _ = fi146
                                            _ = v147
                                            _ = v151
                                        }
                                    default:
                                        fmt.Println(fi54, fi54, fi54)
                                    }
                                } else {
                                    if false {
                                        fmt.Println(v58, v58, fi54)
                                    }
                                }
                                if true {
                                    if len(v37) > 0 {
                                        v37[0] = (fi54 ^ fi54)
                                    }
                                    switch (fi54 ^ 127) {
                                    case 896:
                                        var v152 int32 = v38
                                        _ = v152
                                    }
                                    if len(v37) > 0 {
                                        v37[0] = fi54
                                    }
                                }
                                _ = v120
                            } else {
                                v38 = v38
                                fmt.Println(v58, v44)
                                v50 = v50
                            }
                        }
                    }
                    _ = v119
                }
                _ = v58
            } else {
                fmt.Println(v2, v50, v1)
            }
            _ = fi54
        }
        var v153 string = "\n\t"
        _ = v153
    } else {
        v154 := GenS24{F0: ("\n\t" + "foo bar"), F1: "\n\t", F2: 826}
        v38 -= v38
        _ = v154
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v37:", len(v37), v37)
    fmt.Println("v38:", v38)
    fmt.Println("v44:", v44)
    fmt.Println("v50:", v50)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v44
    _ = v50
    _ = v51
}
