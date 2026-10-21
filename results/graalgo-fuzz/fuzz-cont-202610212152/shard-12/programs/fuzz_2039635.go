package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool }
type GenS2 struct { F0 int8; F1 int; F2 string }
type GenS3 struct { F0 int64; F1 int64 }
type GenS4 struct { F0 int; F1 uint }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 int64; F1 uint32; F2 int8 }
type GenS7 struct { F0 uint; F1 uint64 }
type GenS8 struct { F0 int64; F1 int16; F2 int8 }
type GenS9 struct { F0 int8; F1 int8 }
type GenS10 struct { F0 uint; F1 uint; F2 int32 }
type GenS11 struct { F0 uint16 }
type GenS12 struct { F0 int16; F1 uint16 }
type GenS13 struct { F0 int32; F1 bool; F2 uint }
type GenS14 struct { F0 string; F1 uint }
type GenS15 struct { F0 int }
type GenS16 struct { F0 string; F1 uint16 }
func genFn0(p0 uint8, p1 bool, p2 int8) int64 {
    for fi6 := 0; fi6 < 6; fi6++ {
        var v7 int16 = 42
        for fi8 := 0; fi8 < 2; fi8++ {
            v9 := []int{(fi6 % ((fi6) | 1)), (fi6 * fi8), (fi8 ^ fi6)}
            if p1 {
                v9 = append(v9, fi8)
            }
            _ = fi8
            _ = v9
        }
        switch int(p0) {
        case -446:
            p2 = p2
        case -995:
            var v10 uint16 = 38
            _ = v10
        case 127:
            for fi11 := 0; fi11 < 4; fi11++ {
                switch (-(fi11)) {
                case -9223372036854775808:
                    var v12 int32 = 61
                    _ = v12
                case -858:
                    v13 := GenS1{F0: 9223372036854775807, F1: true}
                    switch (fi11 >> ((909) & 63)) {
                    case -9223372036854775808:
                        p2 += 100
                    }
                    if true {
                        var v14 float32 = ((0.0 + (0.1 - (-2.25 + float32(float64(3.14159))))) - 3.14159)
                        if (p1 && p1) {
                            switch (("a" + (("a" + "foo bar") + "a")) + ("hello" + "a")) {
                            case "a":
                                if ((p0 != (p0 >> ((4294967295) & 63))) || (4294967295 == 255)) {
                                    if (((v14 - (v14 - 0.1)) * v14) <= 1e10) {
                                        p0 -= p0
                                        fmt.Println(p0)
                                    } else {
                                        fmt.Println(p2, fi11)
                                        var v15 uint8 = (p0 >> ((255) & 63))
                                        _ = v15
                                    }
                                    p0 = (p0 / (((^(p0))) | 1))
                                    switch (("\n\t" + (("a" + "hello") + ("" + ("hello" + ("hello" + ""))))) + ("hello" + "hello")) {
                                    case "\n\t":
                                        if ((!((!(32768 > (-(fi11)))) || p1)) || p1) {
                                            v16 := []int{}
                                            fmt.Println(p1)
                                            switch -267 {
                                            case 1:
                                                switch "a" {
                                                case "foo bar":
                                                    if true {
                                                        if p1 {
                                                            p0 = (p0 | 42)
                                                            switch ("" + "hello") {
                                                            case "a":
                                                                var v17 int = 128
                                                                _ = v17
                                                            case "hello":
                                                                p0 -= 0
                                                            default:
                                                                if (42 <= 1) {
                                                                    _ = len(v16)
                                                                    p0 += p0
                                                                    fmt.Println(fi6)
                                                                }
                                                            }
                                                            if ((!true) || true) {
                                                                if len(v16) > 0 {
                                                                    v16[0] = fi11
                                                                }
                                                                if (!(0 >= 450)) {
                                                                    v18 := GenS1{F0: -847, F1: (v14 < v14)}
                                                                    _ = v18
                                                                }
                                                                switch (fi6 >> ((837) & 63)) {
                                                                case 0:
                                                                    v7 = 1
                                                                case -935:
                                                                    v7 += (-(v7))
                                                                case 127:
                                                                    v19 := GenS2{F0: (p2 + -97), F1: -73, F2: "\n\t"}
                                                                    _ = v19
                                                                }
                                                                var v20 int8 = p2
                                                                if len(v16) >= 2 {
                                                                    v16 = v16[1:len(v16)]
                                                                }
                                                                v21 := GenS2{F0: p2, F1: 42, F2: "\n\t"}
                                                                _ = v20
                                                                _ = v21
                                                            } else {
                                                                var v22 uint32 = 497
                                                                v16 = append(v16, (^(fi6)))
                                                                _ = v22
                                                            }
                                                            v23 := GenS2{F0: 42, F1: fi11, F2: ("\n\t" + ("hello" + ""))}
                                                            p1 = p1
                                                            v14 = v14
                                                            var v24 uint = 4294967295
                                                            _ = v23
                                                            _ = v24
                                                        }
                                                        if p1 {
                                                            p0 = p0
                                                        }
                                                        p2 -= p2
                                                    } else {
                                                        var v25 float64 = 3.14159
                                                        if p1 {
                                                            v26 := GenS1{F0: int64(v7), F1: (!(-2147483648 != 0))}
                                                            _ = v26
                                                        } else {
                                                            var v27 uint32 = 2
                                                            _ = v27
                                                        }
                                                        var v28 uint = 0
                                                        _ = v25
                                                        _ = v28
                                                    }
                                                case "hello":
                                                    if (!p1) {
                                                        v7 = (v7 / ((v7) | 1))
                                                        v29 := GenS1{F0: 754, F1: (int32(v7) != -1)}
                                                        _ = v29
                                                    } else {
                                                        v30 := GenS3{F0: 1, F1: -763}
                                                        switch (fi11 & int(p0)) {
                                                        case 4294967295:
                                                            if (!(("\n\t" + "\n\t") >= ("hello" + "foo bar"))) {
                                                                var v31 uint64 = 746
                                                                var v32 int16 = v7
                                                                var v33 int = (fi6 & fi6)
                                                                v34 := GenS4{F0: v33, F1: 1}
                                                                _ = v31
                                                                _ = v32
                                                                _ = v33
                                                                _ = v34
                                                            } else {
                                                                var v35 int8 = (p2 % ((100) | 1))
                                                                _ = v35
                                                            }
                                                        case -1:
                                                            switch "" {
                                                            case "a":
                                                                var v36 int64 = int64(p2)
                                                                _ = v36
                                                            default:
                                                                fmt.Println(p0)
                                                            }
                                                        default:
                                                            if len(v16) > 0 {
                                                                v16[0] = fi11
                                                            }
                                                        }
                                                        var v37 string = "a"
                                                        _ = v30
                                                        _ = v37
                                                    }
                                                    var v38 uint8 = p0
                                                    fmt.Println(fi11)
                                                    _ = v38
                                                }
                                            case -1:
                                                v7 += (^(v7))
                                            case 2147483647:
                                                if true {
                                                    p1 = p1
                                                } else {
                                                    if len(v16) > 0 {
                                                        v16[0] = -287
                                                    }
                                                }
                                            }
                                            if p1 {
                                                if false {
                                                    if (("\n\t" >= ("foo bar" + "foo bar")) && (p2 <= p2)) {
                                                        if len(v16) >= 2 {
                                                            v16 = v16[1:len(v16)]
                                                        }
                                                        v39 := GenS1{F0: -9223372036854775808, F1: (false && false)}
                                                        _ = v39
                                                    } else {
                                                        var v40 float32 = ((0.0 - v14) + float32(1.5))
                                                        _ = v40
                                                    }
                                                }
                                                p1 = true
                                                v16 = append(v16, fi11)
                                            } else {
                                                v41 := GenS5{F0: 42}
                                                _ = v41
                                            }
                                            _ = v16
                                        }
                                    default:
                                        fmt.Println(p2)
                                    }
                                } else {
                                    switch (fi11 << ((4294967295) & 63)) {
                                    case 32768:
                                        v42 := []int{127, fi6, (fi11 * fi6)}
                                        _ = v42
                                    }
                                }
                            case "foo bar":
                                if true {
                                    var v43 string = "\n\t"
                                    var v44 int64 = 0
                                    _ = v43
                                    _ = v44
                                }
                            default:
                                if p1 {
                                    v45 := GenS6{F0: 9223372036854775807, F1: 344, F2: (p2 | -1)}
                                    if p1 {
                                        fmt.Println(p2, fi11, v7)
                                        v7 -= v7
                                    } else {
                                        var v46 bool = p1
                                        fmt.Println(p2)
                                        v46 = false
                                        _ = v46
                                    }
                                    _ = v45
                                }
                            }
                            switch 127 {
                            case 199:
                                if p1 {
                                    if (1 >= 42) {
                                        p0 = (p0 % (((^(p0))) | 1))
                                    }
                                    fmt.Println(p2, fi11, fi6)
                                }
                            case 64:
                                switch 42 {
                                case 4294967295:
                                    v47 := GenS1{F0: 0, F1: (v14 >= -2.25)}
                                    _ = v47
                                default:
                                    v48 := GenS7{F0: 219, F1: 18446744073709551615}
                                    _ = v48
                                }
                                fmt.Println(v14, v7, p0)
                                if (2 <= 450) {
                                    var v49 int16 = -1
                                    _ = v49
                                } else {
                                    v50 := []int{}
                                    _ = v50
                                }
                            default:
                                switch -2147483648 {
                                case 584:
                                    if (1 <= 42) {
                                        var v51 uint64 = 169
                                        switch (fi11 / ((-9223372036854775808) | 1)) {
                                        case 72:
                                            var v52 bool = false
                                            _ = v52
                                        case 792:
                                            v53 := GenS8{F0: -9223372036854775808, F1: v7, F2: 50}
                                            _ = v53
                                        case 1:
                                            if true {
                                                v54 := GenS9{F0: p2, F1: p2}
                                                _ = v54
                                            } else {
                                                fmt.Println(p2, v51, v51)
                                                var v55 uint8 = (p0 & (p0 >> ((52) & 63)))
                                                _ = v55
                                            }
                                        }
                                        v51 = 2
                                        _ = v51
                                    } else {
                                        switch ((("foo bar" + ("a" + "hello")) + "a") + "hello") {
                                        case "a":
                                            var v56 uint = 2
                                            _ = v56
                                        case "foo bar":
                                            v57 := GenS6{F0: 1, F1: 4294967295, F2: -61}
                                            _ = v57
                                        }
                                        switch (fi11 >> ((0) & 63)) {
                                        case 32767:
                                            p1 = ("" != "hello")
                                        default:
                                            switch "a" {
                                            case "\n\t":
                                                if p1 {
                                                    if (-7 <= (-(p2))) {
                                                        v7 += v7
                                                        p0 += p0
                                                        v58 := GenS10{F0: 364, F1: 372, F2: -2147483648}
                                                        _ = v58
                                                    } else {
                                                        p1 = p1
                                                    }
                                                    if (-23 < (v7 & v7)) {
                                                        v59 := GenS11{F0: 32767}
                                                        _ = v59
                                                    } else {
                                                        var v60 int = (fi11 << ((4294967295) & 63))
                                                        _ = v60
                                                    }
                                                    v61 := []int{}
                                                    var v62 float32 = ((v14 - 3.14159) / (v14 + float32(1e10)))
                                                    fmt.Println(fi11)
                                                    if (false && (!(false || p1))) {
                                                        fmt.Println(p0, fi11)
                                                    }
                                                    _ = v61
                                                    _ = v62
                                                }
                                            case "foo bar":
                                                switch "foo bar" {
                                                case "":
                                                    var v63 int32 = -2147483648
                                                    _ = v63
                                                default:
                                                    switch ("\n\t" + "") {
                                                    case "a":
                                                        if false {
                                                            p0 += (p0 ^ 190)
                                                            if p1 {
                                                                var v64 int8 = 42
                                                                v7 = -32768
                                                                v65 := []int{-2147483648, fi6}
                                                                _ = v64
                                                                _ = v65
                                                            }
                                                            if ((^(p2)) <= (-(p2))) {
                                                                var v66 int8 = -1
                                                                if true {
                                                                    v67 := GenS10{F0: 549, F1: 449, F2: -1}
                                                                    v66 = p2
                                                                    _ = v67
                                                                }
                                                                if ((fi6 >> ((2) & 63)) >= fi6) {
                                                                    v7 = (v7 | (v7 / ((v7) | 1)))
                                                                    p0 -= 255
                                                                }
                                                                _ = v66
                                                            }
                                                        }
                                                    case "hello":
                                                        v68 := GenS9{F0: int8(p0), F1: (p2 & 0)}
                                                        _ = v68
                                                    case "":
                                                        v69 := GenS12{F0: (v7 >> ((0) & 63)), F1: 42}
                                                        _ = v69
                                                    }
                                                }
                                            case "a":
                                                v14 = v14
                                            default:
                                                var v70 float32 = (float32((((-2.25 - (((float64(float32(0.0)) * 3.14159) / 1.0) - 0.1)) * -2.25) + (1.5 / 1e10))) * v14)
                                                _ = v70
                                            }
                                        }
                                        var v71 int32 = 0
                                        switch -572 {
                                        case 381:
                                            if p1 {
                                                if (!(330 != 255)) {
                                                    v72 := GenS9{F0: 0, F1: (^(p2))}
                                                    _ = v72
                                                }
                                            } else {
                                                v73 := GenS2{F0: p2, F1: fi11, F2: ("\n\t" + "")}
                                                p0 = p0
                                                v74 := GenS13{F0: (^(v71)), F1: p1, F2: 55}
                                                var v75 int16 = v7
                                                _ = v73
                                                _ = v74
                                                _ = v75
                                            }
                                        case 2147483647:
                                            v76 := []int{fi6, fi6, fi6}
                                            _ = v76
                                        }
                                        v77 := GenS14{F0: (("hello" + "foo bar") + "hello"), F1: 65535}
                                        _ = v71
                                        _ = v77
                                    }
                                case 32768:
                                    fmt.Println(p1, v7)
                                    if (v7 <= v7) {
                                        switch "hello" {
                                        case "\n\t":
                                            if p1 {
                                                switch fi11 {
                                                case -88:
                                                    var v78 int64 = 1
                                                    _ = v78
                                                case -2147483648:
                                                    var v79 uint64 = 480
                                                    _ = v79
                                                case 1:
                                                    var v80 int64 = -9223372036854775808
                                                    _ = v80
                                                default:
                                                    v81 := []int{fi6, (-(fi6)), (fi6 | (^(fi6))), 4294967295}
                                                    _ = v81
                                                }
                                                var v82 uint64 = 18446744073709551615
                                                switch 128 {
                                                case 1:
                                                    v83 := []int{}
                                                    _ = v83
                                                case 32767:
                                                    if (42 != 618) {
                                                        v84 := []int{2147483647, fi6, (^(fi6))}
                                                        switch fi11 {
                                                        case 32768:
                                                            var v85 uint64 = (v82 - (v82 - 0))
                                                            _ = v85
                                                        case 0:
                                                            var v86 bool = p1
                                                            _ = v86
                                                        case 17:
                                                            switch ("" + "") {
                                                            case "\n\t":
                                                                switch (("a" + "a") + "hello") {
                                                                case "\n\t":
                                                                    var v87 int16 = v7
                                                                    _ = v87
                                                                case "":
                                                                    if (p2 < (^(p2))) {
                                                                        v84 = append(v84, fi11)
                                                                    }
                                                                }
                                                            case "hello":
                                                                var v88 int16 = v7
                                                                _ = v88
                                                            case "foo bar":
                                                                if len(v84) > 0 {
                                                                    v84[0] = fi6
                                                                }
                                                            default:
                                                                var v89 float32 = 3.14159
                                                                v90 := GenS13{F0: 42, F1: p1, F2: uint(p0)}
                                                                if (789 != 79) {
                                                                    v91 := GenS7{F0: 269, F1: (v82 >> ((42) & 63))}
                                                                    if (-119 == (p2 + (p2 / (((p2 + -1)) | 1)))) {
                                                                        switch (fi11 >> ((4294967295) & 63)) {
                                                                        case 0:
                                                                            p0 += p0
                                                                        case 127:
                                                                            var v92 int32 = int32(p0)
                                                                            _ = v92
                                                                        case 1:
                                                                            v93 := GenS5{F0: v82}
                                                                            _ = v93
                                                                        }
                                                                        var v94 float32 = ((((1e10 + v89) / ((v14 / float32(3.14159)) - ((-2.25 - -2.25) * float32(float64(float32(0.0)))))) - v14) + float32(1.5))
                                                                        _ = v94
                                                                    }
                                                                    v14 += v89
                                                                    _ = v91
                                                                }
                                                                if (65535 == 65535) {
                                                                    fmt.Println(p0, p0, v14)
                                                                } else {
                                                                    v14 -= 0.0
                                                                    v95 := GenS7{F0: 0, F1: 18446744073709551615}
                                                                    _ = v95
                                                                }
                                                                if len(v84) >= 2 {
                                                                    v84 = v84[1:len(v84)]
                                                                }
                                                                _ = v89
                                                                _ = v90
                                                            }
                                                        }
                                                        _ = v84
                                                    }
                                                case 684:
                                                    v82 = v82
                                                }
                                                _ = v82
                                            } else {
                                                v96 := []int{}
                                                _ = v96
                                            }
                                        }
                                        p2 = 1
                                    }
                                case -2147483648:
                                    fmt.Println(p0, p0, fi6)
                                default:
                                    p2 -= 127
                                }
                            }
                        }
                        v7 = (v7 >> ((1) & 63))
                        v97 := []int{}
                        _ = v14
                        _ = v97
                    } else {
                        v7 -= (v7 >> ((255) & 63))
                        var v98 uint64 = 42
                        _ = v98
                    }
                    _ = v13
                default:
                    var v99 uint64 = 18446744073709551615
                    _ = v99
                }
                fmt.Println(p0)
                _ = fi11
            }
        }
        _ = fi6
        _ = v7
    }
    if (!(p2 <= -128)) {
        for fi100 := 0; fi100 < 6; fi100++ {
            p1 = (529 > 2)
            for fi101 := 0; fi101 < 5; fi101++ {
                if (!(3.14159 > 3.14159)) {
                    if false {
                        v102 := []int{(fi100 * (fi100 * fi100)), fi101}
                        var v103 string = "\n\t"
                        v103 = v103
                        _ = v102
                        _ = v103
                    }
                    if (false || (float64(1e10) > 3.14159)) {
                        v104 := GenS15{F0: (fi100 >> ((18446744073709551615) & 63))}
                        p0 = (p0 | (p0 | 0))
                        p1 = ((true || (((((-2.25 + (1.5 + 0.0)) - 0.1) + ((-2.25 * 0.1) + 0.0)) >= 3.14159) && ("foo bar" == ""))) && (!false))
                        fmt.Println(p2)
                        _ = v104
                    } else {
                        var v105 uint32 = 176
                        _ = v105
                    }
                }
                _ = fi101
            }
            for fi106 := 0; fi106 < 13; fi106++ {
                v107 := []int{32768, 768, (fi106 / (((fi106 * fi106)) | 1))}
                _ = len(v107)
                v108 := GenS5{F0: 4294967295}
                p1 = p1
                var v109 int = (fi106 & 42)
                var v110 uint64 = 318
                v107 = append(v107, (v109 - v109))
                _ = fi106
                _ = v107
                _ = v108
                _ = v109
                _ = v110
            }
            _ = fi100
        }
        switch -392 {
        case 42:
            for fi111 := 0; fi111 < 17; fi111++ {
                var v112 uint32 = 674
                var v113 bool = p1
                _ = fi111
                _ = v112
                _ = v113
            }
        case 88:
            v114 := GenS16{F0: "a", F1: 42}
            _ = v114
        case -764:
            var v115 int = -268
            _ = v115
        default:
            p1 = false
        }
    } else {
        var v116 float32 = 1e10
        _ = v116
    }
    var v117 uint16 = 32767
    for fi118 := 0; fi118 < 13; fi118++ {
        v119 := []int{fi118, (fi118 * fi118), fi118, (fi118 << ((18446744073709551615) & 63))}
        var v120 uint = 672
        p0 += (p0 / ((1) | 1))
        _ = fi118
        _ = v119
        _ = v120
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v117
    return -348
}

func main() {
    v1 := []string{"foo bar", "hello", "foo bar", "hello"}
    var v2 float32 = (1.5 * (float32(0.0) * (3.14159 * float32(0.1))))
    var v3 float32 = (v2 + float32(1.5))
    v4 := make(map[string]string)
    var v5 uint64 = 0
    fmt.Println(v2, v2, v5)
    if (!true) {
        for fi121 := 0; fi121 < 13; fi121++ {
            v3 = v2
            _ = fi121
        }
    } else {
        if len(v1) > 0 {
            v1[0] = ("hello" + "foo bar")
        }
        for fi122 := 0; fi122 < 3; fi122++ {
            v5 = (v5 >> ((0) & 63))
            _ = fi122
        }
        var v123 uint64 = (v5 * v5)
        _ = v123
    }
    v3 -= ((v2 + float32(1e10)) * float32(1.5))
    var v124 uint64 = v5
    var v125 float64 = 1e10
    var v126 float64 = float64(-2.25)
    fmt.Println(v5, v126, v126)
    fmt.Println(v3)
    for fi127 := 0; fi127 < 15; fi127++ {
        fmt.Println(fi127, v124, v5)
        _ = len(v1)
        _ = fi127
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v124:", v124)
    fmt.Println("v125:", v125)
    fmt.Println("v126:", v126)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v124
    _ = v125
    _ = v126
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
