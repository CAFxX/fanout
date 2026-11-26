package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint }
type GenS2 struct { F0 uint32; F1 uint32; F2 uint }
type GenS3 struct { F0 uint16; F1 uint64 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 int }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 int }
type GenS9 struct { F0 uint }
type GenS10 struct { F0 int8 }
type GenS11 struct { F0 uint8 }
type GenS12 struct { F0 int; F1 uint32; F2 int }
type GenS13 struct { F0 int }
type GenS14 struct { F0 uint; F1 bool; F2 int64 }
type GenS15 struct { F0 uint64; F1 int }
type GenS16 struct { F0 int16; F1 uint8; F2 uint16 }
type GenS17 struct { F0 int32; F1 bool; F2 int64 }
type GenS18 struct { F0 bool; F1 int8; F2 int16 }
type GenS19 struct { F0 uint; F1 uint16 }
type GenS20 struct { F0 int16; F1 int32; F2 int16 }
type GenS21 struct { F0 uint32; F1 uint64 }
func genFn0(p0 string, p1 uint32) int64 {
    var v4 int = 32767
    var v5 uint8 = 203
    var v6 bool = (1 >= 175)
    for fi7 := 0; fi7 < 13; fi7++ {
        v8 := GenS1{F0: p0, F1: 1}
        v9 := GenS1{F0: ((p0 + p0) + p0), F1: 65535}
        var v10 int8 = -128
        _ = fi7
        _ = v10
        _ = v8
        _ = v9
    }
    for fi11 := 0; fi11 < 9; fi11++ {
        switch p0 {
        case "":
            v12 := []int{v4, (v4 / (((fi11 >> ((42) & 63))) | 1))}
            _ = v12
        }
        _ = fi11
    }
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    return 9223372036854775807
}
func genFn1(p0 int64) int8 {
    if false {
        p0 = genFn0("\n\t", 2147483647)
    } else {
        p0 -= 226
    }
    v13 := []int{0, 32768, 127, 9223372036854775807}
    _ = p0
    _ = v13
    return 127
}

func main() {
    v1 := make(map[string]uint32)
    var v2 float64 = 3.14159
    v3 := make(map[int]string)
    v14 := GenS1{F0: ("hello" + "\n\t"), F1: 87}
    if true {
        var v15 float64 = (v2 * v2)
        for fi16 := 0; fi16 < 6; fi16++ {
            v15 = (v15 * v15)
            if (205 < 290) {
                v15 = (float64(-2.25) * (v2 / (float64((3.14159 - 3.14159)) - 0.1)))
                var v17 uint64 = 615
                var v18 int8 = 100
                _ = v17
                _ = v18
            } else {
                var v19 int64 = -81
                _ = v19
            }
            v15 = 3.14159
            _ = fi16
        }
        for fi20 := 0; fi20 < 12; fi20++ {
            switch (-(fi20)) {
            case 128:
                v21 := []int{0, 42, (fi20 << ((95) & 63)), 128}
                _ = v21
            case 32767:
                switch "" {
                case "\n\t":
                    if (!(34 > 17)) {
                        v22 := GenS2{F0: 1, F1: 65535, F2: 18446744073709551615}
                        switch (("a" + "foo bar") + ("hello" + "a")) {
                        case "\n\t":
                            if true {
                                var v23 uint32 = 946
                                v23 -= (v23 & (v23 | 2147483647))
                                _ = v23
                            }
                        }
                        v24 := GenS3{F0: 42, F1: 2}
                        _ = v22
                        _ = v24
                    } else {
                        if false {
                            v15 -= float64(1.5)
                            v25 := GenS2{F0: 42, F1: 292, F2: 963}
                            _ = v25
                        }
                        v15 -= 0.1
                    }
                case "hello":
                    v26 := []int{}
                    _ = v26
                case "a":
                    if ((fi20 >> ((1) & 63)) != fi20) {
                        for fi27 := 0; fi27 < 4; fi27++ {
                            var v28 float64 = (3.14159 / -2.25)
                            _ = fi27
                            _ = v28
                        }
                        v29 := []int{(fi20 >> ((4294967295) & 63))}
                        _ = v29
                    } else {
                        var v30 int64 = 9223372036854775807
                        for fi31 := 0; fi31 < 11; fi31++ {
                            var v32 int64 = (v30 * v30)
                            if false {
                                fmt.Println(fi20)
                                v15 = ((0.1 / (v15 * float64(-2.25))) - v2)
                                var v33 int8 = 100
                                var v34 int16 = 3
                                var v35 float64 = float64(-2.25)
                                if true {
                                    v30 = -133
                                    v32 -= genFn0("a", 1)
                                    v35 -= -2.25
                                    var v36 int8 = (v33 >> ((4294967295) & 63))
                                    _ = v36
                                } else {
                                    v3[fi31] = "foo bar"
                                }
                                v35 = 0.1
                                if (v33 > (v33 & (v33 | v33))) {
                                    if false {
                                        var v37 float64 = float64(1e10)
                                        fmt.Println(v35)
                                        if (663 >= (v34 / ((-887) | 1))) {
                                            v30 = -753
                                        }
                                        fmt.Println(v30, v32)
                                        _ = len(v3)
                                        _ = v37
                                    } else {
                                        v15 -= ((v35 * v35) / (v2 + (v35 + (float64(((float32((3.14159 * 1.5)) * 1e10) * (-2.25 - 0.0))) - v2))))
                                        switch "foo bar" {
                                        case "":
                                            v38 := GenS4{F0: (v34 - 23)}
                                            _ = v38
                                        }
                                    }
                                    v39 := []int{(fi31 | -9223372036854775808), fi20}
                                    var v40 uint32 = 2147483648
                                    _ = v39
                                    _ = v40
                                } else {
                                    var v41 int16 = (v34 % (((v34 | v34)) | 1))
                                    var v42 uint64 = 659
                                    _ = v41
                                    _ = v42
                                }
                                _ = v33
                                _ = v34
                                _ = v35
                            }
                            fmt.Println(fi31, v30)
                            if (!true) {
                                if (true && true) {
                                    if true {
                                        fmt.Println(fi31)
                                        switch fi20 {
                                        case 966:
                                            if false {
                                                fmt.Println(v15, v2, v15)
                                                var v43 uint8 = 0
                                                v43 = v43
                                                var v44 uint64 = 18446744073709551615
                                                fmt.Println(v44)
                                                _ = v43
                                                _ = v44
                                            }
                                        case 9223372036854775807:
                                            switch ("a" + "hello") {
                                            case "a":
                                                fmt.Println(fi20, v30, v32)
                                            case "\n\t":
                                                switch (fi20 << ((990) & 63)) {
                                                case 42:
                                                    v45 := GenS3{F0: 174, F1: 201}
                                                    _ = v45
                                                case 26:
                                                    if true {
                                                        fmt.Println(v2, fi20)
                                                        v2 = 0.1
                                                        if false {
                                                            v46 := GenS5{F0: 0}
                                                            if true {
                                                                fmt.Println(v15, v30)
                                                                var v47 uint64 = 406
                                                                fmt.Println(v2, fi20)
                                                                _ = v47
                                                            } else {
                                                                v48 := GenS4{F0: 1}
                                                                var v49 float64 = 3.14159
                                                                _ = v48
                                                                _ = v49
                                                            }
                                                            _ = v46
                                                        }
                                                        if (1 == 777) {
                                                            var v50 int16 = -674
                                                            v51 := []int{fi20, fi20, 42, -9223372036854775808}
                                                            var v52 int = (fi31 ^ fi20)
                                                            _ = v50
                                                            _ = v51
                                                            _ = v52
                                                        } else {
                                                            v53 := GenS1{F0: "\n\t", F1: 1}
                                                            switch "" {
                                                            case "foo bar":
                                                                v32 -= 9223372036854775807
                                                            default:
                                                                v54 := GenS1{F0: "a", F1: 960}
                                                                _ = v54
                                                            }
                                                            switch "hello" {
                                                            case "foo bar":
                                                                if (18446744073709551615 < 2) {
                                                                    fmt.Println(v32, v32)
                                                                    v55 := GenS5{F0: 32767}
                                                                    var v56 int = (^(fi20))
                                                                    _ = v55
                                                                    _ = v56
                                                                }
                                                            case "\n\t":
                                                                if false {
                                                                    v57 := GenS2{F0: 42, F1: 268, F2: 4294967295}
                                                                    _ = v57
                                                                }
                                                            case "a":
                                                                fmt.Println(fi20)
                                                            default:
                                                                v32 = v30
                                                            }
                                                            _ = v53
                                                        }
                                                        v58 := GenS4{F0: -32768}
                                                        v59 := []int{(fi31 << ((4294967295) & 63)), (fi20 ^ fi20), (fi20 + fi31)}
                                                        if false {
                                                            var v60 string = "\n\t"
                                                            if (genFn0("a", 1) != genFn0(((v60 + v60) + "foo bar"), 799)) {
                                                                var v61 int16 = -32768
                                                                _ = v61
                                                            }
                                                            _ = v60
                                                        }
                                                        _ = v58
                                                        _ = v59
                                                    } else {
                                                        var v62 int64 = (v32 & -9223372036854775808)
                                                        var v63 string = "a"
                                                        v64 := GenS3{F0: 2, F1: 4294967295}
                                                        switch (fi20 + 4294967295) {
                                                        case 284:
                                                            v65 := v1[""]
                                                            _ = v65
                                                        case -2147483648:
                                                            v66 := GenS1{F0: "foo bar", F1: 65535}
                                                            _ = v66
                                                        default:
                                                            if true {
                                                                switch "foo bar" {
                                                                case "foo bar":
                                                                    if true {
                                                                        var v67 int = fi20
                                                                        var v68 int8 = 42
                                                                        fmt.Println(v15)
                                                                        _ = v67
                                                                        _ = v68
                                                                    } else {
                                                                        if (1 >= 38) {
                                                                            switch "\n\t" {
                                                                            case "foo bar":
                                                                                var v69 bool = ((float64((1.5 + 0.1)) + 1e10) < (0.0 * (1.5 - 1e10)))
                                                                                _ = v69
                                                                            }
                                                                            v70 := GenS6{F0: (fi20 % (((fi31 / ((-1) | 1))) | 1))}
                                                                            var v71 int16 = -32768
                                                                            _ = v70
                                                                            _ = v71
                                                                        } else {
                                                                            v2 -= (0.0 * (((v15 - 1e10) - -2.25) + 1e10))
                                                                            v62 = v62
                                                                            if false {
                                                                                v15 += v2
                                                                                if (true || false) {
                                                                                    v15 = 0.1
                                                                                    v72 := GenS7{F0: 0}
                                                                                    _ = v72
                                                                                } else {
                                                                                    var v73 string = (v63 + v63)
                                                                                    if (genFn1(genFn0("a", 261)) >= 93) {
                                                                                        var v74 bool = (4294967295 > 2147483648)
                                                                                        v30 = (v62 * (-(v62)))
                                                                                        _ = v74
                                                                                    } else {
                                                                                        var v75 float32 = float32(v2)
                                                                                        _ = v75
                                                                                    }
                                                                                    _ = v73
                                                                                }
                                                                            } else {
                                                                                if false {
                                                                                    v76 := GenS2{F0: 232, F1: 65535, F2: 233}
                                                                                    _ = v76
                                                                                }
                                                                            }
                                                                        }
                                                                        v15 -= v15
                                                                    }
                                                                }
                                                                v30 -= (v62 * int64(fi20))
                                                                v77 := GenS4{F0: 1}
                                                                if ((fi31 >= fi20) && true) {
                                                                    v15 = v15
                                                                    switch "\n\t" {
                                                                    case "":
                                                                        if true {
                                                                            var v78 uint32 = 2147483647
                                                                            v79 := v1[v63]
                                                                            _ = v78
                                                                            _ = v79
                                                                        } else {
                                                                            if false {
                                                                                v30 = -117
                                                                            }
                                                                            var v80 uint16 = 2
                                                                            _ = v80
                                                                        }
                                                                    case "\n\t":
                                                                        v81 := []int{fi31, 9223372036854775807, (fi20 ^ 4294967295)}
                                                                        _ = v81
                                                                    default:
                                                                        delete(v3, (-(fi31)))
                                                                    }
                                                                    var v82 uint64 = 18446744073709551615
                                                                    _ = v82
                                                                } else {
                                                                    var v83 uint64 = 0
                                                                    var v84 uint64 = (v83 ^ 502)
                                                                    v63 = (v63 + (((v63 + v63) + (((v63 + "hello") + v63) + v63)) + v63))
                                                                    _ = v83
                                                                    _ = v84
                                                                }
                                                                _ = v77
                                                            }
                                                        }
                                                        var v85 int64 = (v32 * v62)
                                                        v86 := GenS1{F0: v63, F1: 4294967295}
                                                        _ = v62
                                                        _ = v63
                                                        _ = v64
                                                        _ = v85
                                                        _ = v86
                                                    }
                                                }
                                            case "hello":
                                                v30 += genFn0(((("a" + ("a" + ("a" + ""))) + "a") + "foo bar"), 2147483647)
                                            default:
                                                var v87 uint64 = 1
                                                _ = v87
                                            }
                                        }
                                        if true {
                                            switch (-(fi20)) {
                                            case -242:
                                                v15 = v2
                                            case 233:
                                                if (true && ("a" < ("" + "a"))) {
                                                    var v88 int32 = 557
                                                    switch ("hello" + "hello") {
                                                    case "\n\t":
                                                        v89 := []int{(fi20 | 9223372036854775807), fi20, (fi31 % ((fi20) | 1)), fi20}
                                                        _ = v89
                                                    case "":
                                                        v15 += v2
                                                    default:
                                                        v15 -= 3.14159
                                                    }
                                                    _ = v88
                                                }
                                            default:
                                                v90 := v1["hello"]
                                                _ = v90
                                            }
                                            v91 := []int{fi31, 386, -681}
                                            v92 := GenS3{F0: 1, F1: 4294967295}
                                            _ = v91
                                            _ = v92
                                        }
                                    }
                                    switch (("\n\t" + "") + "") {
                                    case "hello":
                                        var v93 bool = (!((true || false) || (!true)))
                                        _ = v93
                                    }
                                }
                                switch fi31 {
                                case 9223372036854775807:
                                    v15 = (float64((0.1 * (0.1 * 1e10))) / (3.14159 + (v15 - v2)))
                                case -973:
                                    if ((v32 >> ((388) & 63)) != -9223372036854775808) {
                                        if (((18446744073709551615 < 18446744073709551615) || ((float32(float64(0.0)) + 3.14159) > 1.5)) || false) {
                                            if (!false) {
                                                switch (fi20 >> ((4294967295) & 63)) {
                                                case 2147483647:
                                                    if true {
                                                        var v94 int = (fi31 | fi31)
                                                        var v95 uint16 = 0
                                                        v30 = v32
                                                        if (true || (true && (false || (18446744073709551615 >= 0)))) {
                                                            _ = len(v1)
                                                            v96 := []int{fi20, fi20}
                                                            _ = v96
                                                        }
                                                        switch "a" {
                                                        case "a":
                                                            if false {
                                                                v30 -= v32
                                                            }
                                                        }
                                                        _ = v94
                                                        _ = v95
                                                    }
                                                }
                                            } else {
                                                v15 = v15
                                                v97 := GenS8{F0: fi20}
                                                fmt.Println(v2, v15, fi20)
                                                _ = v97
                                            }
                                            v98 := []int{(fi20 + fi20)}
                                            _ = v98
                                        } else {
                                            if false {
                                                v99 := GenS7{F0: 112}
                                                v32 -= 1
                                                _ = v99
                                            }
                                        }
                                        if (((1e10 / 1e10) + (((((0.0 - v2) - 0.0) - v15) * v2) * (((v2 + v2) * 3.14159) * 0.0))) >= v15) {
                                            fmt.Println(v15, v2, fi20)
                                            var v100 int32 = 916
                                            if true {
                                                v15 -= v2
                                                if false {
                                                    switch "hello" {
                                                    case "\n\t":
                                                        if (42 <= 4294967295) {
                                                            v101 := GenS3{F0: 0, F1: 303}
                                                            v2 -= v2
                                                            v102 := GenS9{F0: 255}
                                                            v2 -= v2
                                                            _ = v101
                                                            _ = v102
                                                        } else {
                                                            v32 = v30
                                                            var v103 uint16 = 700
                                                            if (!((v2 * -2.25) == v15)) {
                                                                v104 := GenS4{F0: 1}
                                                                v105 := []int{fi20, (fi31 * 0), (fi20 / (((fi20 & fi20)) | 1))}
                                                                v103 = (v103 + (v103 - v103))
                                                                var v106 uint = 255
                                                                var v107 float64 = (v15 - float64(1.5))
                                                                _ = v104
                                                                _ = v105
                                                                _ = v106
                                                                _ = v107
                                                            }
                                                            _ = v103
                                                        }
                                                    case "foo bar":
                                                        fmt.Println(v30)
                                                    case "hello":
                                                        if (383 == -32768) {
                                                            switch fi20 {
                                                            case -9223372036854775808:
                                                                v108 := []int{(fi20 * (fi20 % (((-(fi31))) | 1))), (-(fi31))}
                                                                _ = v108
                                                            case 9223372036854775807:
                                                                var v109 uint32 = 4294967295
                                                                _ = v109
                                                            case 42:
                                                                if (42 >= 32767) {
                                                                    var v110 float64 = (1.5 - v15)
                                                                    _ = v110
                                                                } else {
                                                                    v100 = v100
                                                                }
                                                            }
                                                            v15 = 0.0
                                                        }
                                                        v3[-1] = "\n\t"
                                                    default:
                                                        v3[(fi20 >> ((65535) & 63))] = "foo bar"
                                                    }
                                                } else {
                                                    fmt.Println(v15, v30)
                                                    v111 := GenS2{F0: 372, F1: 42, F2: 105}
                                                    v3[(fi31 & fi20)] = "\n\t"
                                                    v15 -= 0.1
                                                    if false {
                                                        var v112 uint32 = 65535
                                                        v100 = (v100 | (v100 ^ (v100 << ((0) & 63))))
                                                        fmt.Println(v112, fi31)
                                                        _ = v112
                                                    }
                                                    _ = v111
                                                }
                                            }
                                            fmt.Println(fi20, v100, fi31)
                                            v113 := GenS10{F0: 42}
                                            _ = v100
                                            _ = v113
                                        }
                                    } else {
                                        if false {
                                            v114 := GenS4{F0: -326}
                                            _ = v114
                                        }
                                        switch fi20 {
                                        case -9223372036854775808:
                                            v115 := GenS11{F0: 118}
                                            _ = v115
                                        case 128:
                                            fmt.Println(v32)
                                            v2 -= v2
                                        }
                                    }
                                case 4294967295:
                                    if false {
                                        var v116 bool = (1 < (fi20 << ((0) & 63)))
                                        if false {
                                            var v117 uint16 = 32767
                                            var v118 int32 = 2147483647
                                            _ = v117
                                            _ = v118
                                        } else {
                                            v119 := GenS11{F0: 0}
                                            _ = v119
                                        }
                                        switch "hello" {
                                        case "a":
                                            v120 := v3[fi20]
                                            _ = v120
                                        case "":
                                            v15 = 1e10
                                        }
                                        v15 -= 0.0
                                        _ = v116
                                    } else {
                                        v30 = v30
                                        v121 := []int{}
                                        var v122 uint = 217
                                        _ = v121
                                        _ = v122
                                    }
                                }
                            }
                            if false {
                                v32 -= (^(v30))
                                var v123 uint8 = 255
                                var v124 uint = 1
                                var v125 float64 = (3.14159 * v15)
                                if (!(("" + ("" + "a")) <= (("hello" + "") + ("foo bar" + "foo bar")))) {
                                    v126 := GenS12{F0: 42, F1: 306, F2: -1}
                                    switch "foo bar" {
                                    case "a":
                                        v32 -= genFn0("", 0)
                                    default:
                                        var v127 int64 = (-(v32))
                                        _ = v127
                                    }
                                    switch fi20 {
                                    case -908:
                                        if true {
                                            if (((3.14159 < 0.1) || false) && (true || true)) {
                                                v128 := GenS3{F0: 42, F1: 741}
                                                delete(v1, ("hello" + (("hello" + (("" + ("hello" + "foo bar")) + "a")) + ("a" + "foo bar"))))
                                                if false {
                                                    v15 -= (3.14159 + 1.5)
                                                    v129 := GenS11{F0: v123}
                                                    _ = v129
                                                }
                                                _ = v128
                                            } else {
                                                v125 += v2
                                            }
                                        } else {
                                            v124 -= v124
                                        }
                                    case 42:
                                        fmt.Println(v125)
                                    }
                                    _ = v126
                                }
                                _ = v123
                                _ = v124
                                _ = v125
                            } else {
                                var v130 uint16 = 551
                                var v131 int16 = 32767
                                _ = v130
                                _ = v131
                            }
                            v132 := []int{(fi31 * (fi31 << ((279) & 63)))}
                            v133 := GenS2{F0: 65535, F1: 2147483647, F2: 0}
                            fmt.Println(v32, v15)
                            if true {
                                v134 := GenS1{F0: "\n\t", F1: 1}
                                _ = v134
                            } else {
                                v15 = ((3.14159 * 0.0) * v15)
                                var v135 uint32 = 0
                                _ = v135
                            }
                            _ = fi31
                            _ = v132
                            _ = v133
                            _ = v32
                        }
                        _ = v30
                    }
                default:
                    var v136 int16 = 960
                    _ = v136
                }
            default:
                fmt.Println(v2)
            }
            switch (^(fi20)) {
            case 127:
                v15 -= float64(-2.25)
            default:
                var v137 uint = 135
                _ = v137
            }
            _ = fi20
        }
        _ = v15
    } else {
        if (false || (!false)) {
            if (558 > genFn0("", 135)) {
                if false {
                    for fi138 := 0; fi138 < 20; fi138++ {
                        var v139 int8 = genFn1(608)
                        _ = fi138
                        _ = v139
                    }
                    _ = len(v3)
                    if (!(v2 > (3.14159 + ((-2.25 + (float64(float32(1.5)) - float64(0.0))) * float64(((float32(v2) * 0.0) * float32(v2))))))) {
                        fmt.Println(v2)
                    } else {
                        for fi140 := 0; fi140 < 10; fi140++ {
                            if (-1 <= -981) {
                                fmt.Println(fi140, v2)
                            } else {
                                v2 = 1e10
                                if false {
                                    v141 := GenS1{F0: "a", F1: 255}
                                    _ = v141
                                }
                                v2 = (1.5 - v2)
                            }
                            _ = fi140
                        }
                        fmt.Println(v2)
                        if ((true || true) || true) {
                            for fi142 := 0; fi142 < 5; fi142++ {
                                v3[fi142] = "hello"
                                _ = fi142
                            }
                            switch "hello" {
                            case "foo bar":
                                var v143 uint = 0
                                _ = v143
                            case "\n\t":
                                var v144 uint = 1
                                _ = v144
                            case "":
                                var v145 int8 = -77
                                _ = v145
                            }
                            if true {
                                if (!(-32768 != -842)) {
                                    fmt.Println(v2)
                                    switch "\n\t" {
                                    case "hello":
                                        _ = len(v1)
                                    case "a":
                                        v146 := []int{-224, 32768, -341, 2147483647}
                                        _ = v146
                                    default:
                                        v2 += 0.1
                                    }
                                }
                                v2 = v2
                                for fi147 := 0; fi147 < 1; fi147++ {
                                    for fi148 := 0; fi148 < 1; fi148++ {
                                        fmt.Println(fi147)
                                        v149 := []int{fi148, fi147, 689}
                                        v150 := GenS1{F0: ("foo bar" + ("\n\t" + "")), F1: 65535}
                                        _ = fi148
                                        _ = v149
                                        _ = v150
                                    }
                                    var v151 int = fi147
                                    switch ((("" + ("\n\t" + "foo bar")) + ("foo bar" + (((("hello" + ("a" + "\n\t")) + ("foo bar" + "")) + ((("" + "foo bar") + ("a" + ("hello" + "a"))) + "\n\t")) + ("foo bar" + ("" + "a"))))) + "hello") {
                                    case "hello":
                                        if true {
                                            switch "foo bar" {
                                            case "":
                                                v152 := GenS9{F0: 18446744073709551615}
                                                _ = v152
                                            case "a":
                                                if (true && (1 <= 255)) {
                                                    v151 -= (fi147 - (v151 % (((v151 / ((732) | 1))) | 1)))
                                                } else {
                                                    var v153 string = "\n\t"
                                                    if (false && true) {
                                                        fmt.Println(v151, v2, v151)
                                                        for fi154 := 0; fi154 < 15; fi154++ {
                                                            if (false && false) {
                                                                var v155 int = fi154
                                                                switch (v153 + (("" + v153) + "")) {
                                                                case "":
                                                                    switch fi147 {
                                                                    case 1:
                                                                        if (!false) {
                                                                            if true {
                                                                                v156 := GenS13{F0: (-(fi147))}
                                                                                v157 := []int{0, (v155 & (v151 % ((612) | 1)))}
                                                                                if true {
                                                                                    v158 := GenS14{F0: 897, F1: true, F2: 9223372036854775807}
                                                                                    var v159 int32 = 986
                                                                                    _ = v158
                                                                                    _ = v159
                                                                                } else {
                                                                                    var v160 string = (v153 + ("" + v153))
                                                                                    _ = v160
                                                                                }
                                                                                _ = v156
                                                                                _ = v157
                                                                            } else {
                                                                                fmt.Println(v2, v153, fi147)
                                                                                var v161 int = 32768
                                                                                v162 := []int{v161, v161, (v161 / (((v161 + v155)) | 1)), (-(fi147))}
                                                                                _ = v161
                                                                                _ = v162
                                                                            }
                                                                            v163 := []int{}
                                                                            _ = v163
                                                                        }
                                                                    case -2147483648:
                                                                        if (false || true) {
                                                                            v151 -= v155
                                                                            v2 -= (3.14159 * (v2 + float64(0.1)))
                                                                            switch v153 {
                                                                            case "":
                                                                                _ = len(v1)
                                                                            case "hello":
                                                                                v151 += (fi147 ^ (fi154 ^ 315))
                                                                            case "\n\t":
                                                                                if (1 >= 18446744073709551615) {
                                                                                    v164 := []int{392, (v151 + 537), v151, 2147483647}
                                                                                    v3[(v151 | fi154)] = ("foo bar" + v153)
                                                                                    v151 += 442
                                                                                    _ = v164
                                                                                }
                                                                            }
                                                                        }
                                                                    case 316:
                                                                        var v165 bool = false
                                                                        _ = v165
                                                                    }
                                                                case "foo bar":
                                                                    v166 := GenS12{F0: (fi154 | v151), F1: 65535, F2: -2147483648}
                                                                    _ = v166
                                                                default:
                                                                    fmt.Println(v155, v151)
                                                                }
                                                                if (!false) {
                                                                    v153 = ((v153 + (v153 + "\n\t")) + v153)
                                                                    var v167 uint8 = 77
                                                                    _ = v167
                                                                }
                                                                _ = v155
                                                            }
                                                            if false {
                                                                v168 := GenS10{F0: genFn1(-9223372036854775808)}
                                                                v2 += float64(((((1.5 - (3.14159 - 1e10)) / 3.14159) * float32(v2)) - 0.1))
                                                                _ = v168
                                                            } else {
                                                                if (false && true) {
                                                                    v169 := GenS9{F0: 291}
                                                                    v170 := GenS8{F0: 9223372036854775807}
                                                                    _ = v169
                                                                    _ = v170
                                                                } else {
                                                                    var v171 int64 = -9223372036854775808
                                                                    var v172 int8 = 96
                                                                    if (2 != 0) {
                                                                        switch (fi147 / (((fi147 + 396)) | 1)) {
                                                                        case 76:
                                                                            var v173 float64 = 0.0
                                                                            _ = v173
                                                                        case -1:
                                                                            v151 = (fi147 + fi154)
                                                                        }
                                                                    }
                                                                    v171 = (-(v171))
                                                                    if ((!(2 != 0)) || ((-(v172)) < genFn1(genFn0("\n\t", 709)))) {
                                                                        v174 := []int{}
                                                                        var v175 string = v153
                                                                        var v176 int8 = (v172 * 1)
                                                                        _ = v174
                                                                        _ = v175
                                                                        _ = v176
                                                                    } else {
                                                                        v177 := v1["a"]
                                                                        v172 = (v172 / ((genFn1((v171 & v171))) | 1))
                                                                        v153 = v153
                                                                        _ = v177
                                                                    }
                                                                    var v178 int8 = v172
                                                                    if true {
                                                                        var v179 int16 = -247
                                                                        if false {
                                                                            var v180 uint16 = 1
                                                                            var v181 int8 = (v178 & 48)
                                                                            _ = v180
                                                                            _ = v181
                                                                        } else {
                                                                            if false {
                                                                                var v182 bool = true
                                                                                v179 = (v179 - (v179 | v179))
                                                                                _ = v182
                                                                            }
                                                                            var v183 uint64 = 4294967295
                                                                            _ = v183
                                                                        }
                                                                        v184 := GenS15{F0: 2, F1: (fi154 ^ (fi147 * (-(fi154))))}
                                                                        _ = v179
                                                                        _ = v184
                                                                    }
                                                                    _ = v171
                                                                    _ = v172
                                                                    _ = v178
                                                                }
                                                            }
                                                            _ = fi154
                                                        }
                                                    }
                                                    _ = v153
                                                }
                                            case "foo bar":
                                                switch 377 {
                                                case -2147483648:
                                                    fmt.Println(fi147)
                                                }
                                            }
                                            v185 := GenS15{F0: 4294967295, F1: fi147}
                                            for fi186 := 0; fi186 < 3; fi186++ {
                                                v1[("hello" + (((("hello" + ("foo bar" + ("" + ("hello" + ("" + "hello"))))) + "a") + "hello") + "foo bar"))] = 42
                                                switch "\n\t" {
                                                case "\n\t":
                                                    fmt.Println(fi186, v2)
                                                default:
                                                    if true {
                                                        fmt.Println(fi147)
                                                        if false {
                                                            var v187 int8 = 52
                                                            if false {
                                                                var v188 uint16 = 679
                                                                _ = v188
                                                            } else {
                                                                v189 := GenS16{F0: -32768, F1: 255, F2: 42}
                                                                var v190 uint32 = 4294967295
                                                                if true {
                                                                    v151 += -1
                                                                    var v191 float64 = -2.25
                                                                    _ = v191
                                                                } else {
                                                                    v151 += (v151 - v151)
                                                                }
                                                                _ = v189
                                                                _ = v190
                                                            }
                                                            _ = v187
                                                        } else {
                                                            var v192 float64 = (-2.25 / (3.14159 * v2))
                                                            v1[("" + "\n\t")] = 4294967295
                                                            _ = v192
                                                        }
                                                        switch (("\n\t" + "a") + (("foo bar" + "") + (("\n\t" + ("" + "")) + "hello"))) {
                                                        case "a":
                                                            v151 = v151
                                                        default:
                                                            var v193 uint8 = 75
                                                            _ = v193
                                                        }
                                                        if (!false) {
                                                            var v194 string = "a"
                                                            v195 := []int{v151}
                                                            _ = v194
                                                            _ = v195
                                                        }
                                                    }
                                                }
                                                _ = fi186
                                            }
                                            _ = v185
                                        }
                                    case "":
                                        v151 = 768
                                    }
                                    _ = fi147
                                    _ = v151
                                }
                            }
                        }
                    }
                }
                v2 = 1.5
            } else {
                v2 += v2
                v196 := GenS2{F0: 4294967295, F1: 61, F2: 1}
                _ = v196
            }
            for fi197 := 0; fi197 < 1; fi197++ {
                v2 = v2
                _ = fi197
            }
        } else {
            v198 := v3[-34]
            _ = v198
        }
    }
    fmt.Println(v2)
    var v199 int32 = -684
    for fi200 := 0; fi200 < 10; fi200++ {
        v201 := GenS14{F0: 18446744073709551615, F1: (224 >= -1), F2: genFn0("\n\t", 2147483648)}
        v202 := GenS12{F0: (fi200 - (fi200 / ((249) | 1))), F1: 2147483648, F2: (fi200 / (((fi200 << ((65535) & 63))) | 1))}
        v203 := GenS17{F0: v199, F1: (1000 >= 42), F2: 1}
        _ = fi200
        _ = v201
        _ = v202
        _ = v203
    }
    var v204 bool = true
    v205 := GenS18{F0: v204, F1: -61, F2: 32767}
    v206 := []int{}
    for fi207 := 0; fi207 < 18; fi207++ {
        if true {
            if v204 {
                var v208 uint16 = 42
                _ = v208
            } else {
                v209 := GenS19{F0: 65535, F1: 843}
                switch "hello" {
                case "a":
                    v210 := v3[326]
                    _ = v210
                case "hello":
                    var v211 uint32 = 2
                    _ = v211
                case "":
                    v212 := GenS20{F0: 435, F1: v199, F2: -1}
                    _ = v212
                }
                _ = v209
            }
            if (float64(0.0) == 3.14159) {
                var v213 uint = 200
                switch "" {
                case "a":
                    fmt.Println(v2, v199, fi207)
                default:
                    switch (fi207 % (((fi207 - (fi207 ^ (^(fi207))))) | 1)) {
                    case 128:
                        for fi214 := 0; fi214 < 17; fi214++ {
                            if (("foo bar" >= ((("hello" + "a") + "a") + "foo bar")) && (1 != 0)) {
                                var v215 uint64 = 425
                                v213 += (v213 & v213)
                                var v216 int64 = 42
                                _ = v215
                                _ = v216
                            }
                            _ = fi214
                        }
                    case -1:
                        v3[(fi207 | fi207)] = ""
                    }
                }
                switch fi207 {
                case -9223372036854775808:
                    var v217 int64 = 42
                    _ = v217
                case 25:
                    var v218 uint32 = 820
                    _ = v218
                case 2147483647:
                    v219 := GenS21{F0: 4294967295, F1: 2}
                    _ = v219
                }
                _ = v213
            } else {
                if len(v206) >= 2 {
                    v206 = v206[1:len(v206)]
                }
            }
        }
        var v220 uint = 255
        _ = fi207
        _ = v220
    }
    for fi221 := 0; fi221 < 4; fi221++ {
        var v222 uint16 = 32767
        _ = fi221
        _ = v222
    }
    switch 42 {
    case 599:
        delete(v1, "hello")
    }
    var v223 uint = 362
    v224 := GenS7{F0: 94}
    for fi225 := 0; fi225 < 10; fi225++ {
        var v226 bool = ((!v204) || v204)
        _ = fi225
        _ = v226
    }
    if v204 {
        if false {
            switch "\n\t" {
            case "":
                var v227 int8 = -91
                _ = v227
            case "hello":
                var v228 int16 = 1
                _ = v228
            case "a":
                var v229 int8 = 1
                _ = v229
            }
            for fi230 := 0; fi230 < 4; fi230++ {
                var v231 float64 = v2
                var v232 int32 = v199
                _ = fi230
                _ = v231
                _ = v232
            }
        }
        v233 := GenS14{F0: 16, F1: v204, F2: 42}
        v234 := GenS13{F0: 1}
        _ = v233
        _ = v234
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v199:", v199)
    fmt.Println("v2:", v2)
    fmt.Println("v204:", v204)
    fmt.Println("slicelen v206:", len(v206), v206)
    fmt.Println("v223:", v223)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v14
    _ = v199
    _ = v2
    _ = v204
    _ = v205
    _ = v206
    _ = v223
    _ = v224
    _ = v3
}
