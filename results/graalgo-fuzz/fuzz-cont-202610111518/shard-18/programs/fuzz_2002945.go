package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int64; F1 int8 }
type GenS3 struct { F0 uint64; F1 uint8; F2 int16 }
type GenS4 struct { F0 bool; F1 int8 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 bool; F1 bool }
type GenS7 struct { F0 int32; F1 uint8; F2 uint32 }
type GenS8 struct { F0 int; F1 int32 }
type GenS9 struct { F0 int16; F1 int64; F2 uint32 }
type GenS10 struct { F0 bool; F1 int64; F2 uint32 }
type GenS11 struct { F0 uint16 }
type GenS12 struct { F0 uint32; F1 uint64 }
type GenS13 struct { F0 uint32; F1 int; F2 int }
type GenS14 struct { F0 bool }
type GenS15 struct { F0 uint64 }
type GenS16 struct { F0 uint16; F1 uint16 }
func genFn0(p0 uint8) int {
    v5 := GenS1{F0: 1}
    p0 += p0
    switch ("hello" + ("foo bar" + "hello")) {
    case "hello":
        p0 = (^(p0))
    default:
        if (!true) {
            if (("\n\t" + "foo bar") > "") {
                for fi6 := 0; fi6 < 10; fi6++ {
                    if (true && ("a" != ("" + "a"))) {
                        v7 := GenS1{F0: 1}
                        _ = v7
                    }
                    _ = fi6
                }
                if (0 > p0) {
                    v8 := []int{}
                    p0 -= (p0 << ((42) & 63))
                    var v9 int8 = 1
                    _ = v8
                    _ = v9
                } else {
                    v10 := GenS1{F0: 0}
                    _ = v10
                }
            }
            v11 := GenS2{F0: 42, F1: 53}
            p0 = (p0 * 42)
            _ = v11
        } else {
            switch 42 {
            case 4294967295:
                for fi12 := 0; fi12 < 1; fi12++ {
                    if true {
                        v13 := []int{127, 127, fi12, (fi12 | 0)}
                        v13 = append(v13, fi12)
                        var v14 string = (("\n\t" + "hello") + "hello")
                        _ = v13
                        _ = v14
                    } else {
                        if (!(false && (fi12 != (fi12 ^ -2147483648)))) {
                            if false {
                                v15 := []int{-500, 0, (fi12 << ((179) & 63)), -119}
                                p0 -= p0
                                v16 := GenS3{F0: uint64(p0), F1: (p0 - (p0 % ((144) | 1))), F2: 1}
                                _ = v15
                                _ = v16
                            } else {
                                for fi17 := 0; fi17 < 2; fi17++ {
                                    var v18 string = "\n\t"
                                    if (117 < 822) {
                                        switch fi12 {
                                        case 9223372036854775807:
                                            if (!false) {
                                                switch "foo bar" {
                                                case "hello":
                                                    if (true && true) {
                                                        v19 := GenS3{F0: 468, F1: p0, F2: 1000}
                                                        v20 := []int{(fi17 % (((fi12 << ((65535) & 63))) | 1)), fi12, fi17}
                                                        p0 = (^(p0))
                                                        _ = v19
                                                        _ = v20
                                                    } else {
                                                        var v21 float64 = 1e10
                                                        var v22 string = v18
                                                        _ = v21
                                                        _ = v22
                                                    }
                                                case "a":
                                                    var v23 int64 = 1
                                                    _ = v23
                                                case "":
                                                    v18 = v18
                                                default:
                                                    var v24 int = fi12
                                                    var v25 float32 = -2.25
                                                    _ = v24
                                                    _ = v25
                                                }
                                                v26 := []int{-2147483648}
                                                p0 = 153
                                                _ = v26
                                            } else {
                                                var v27 int32 = 715
                                                v28 := []int{(fi17 - (^(fi17))), (-(fi12)), fi12}
                                                if true {
                                                    fmt.Println(p0, fi17, p0)
                                                    v28 = append(v28, (fi12 & (-(fi12))))
                                                    fmt.Println(p0, fi17)
                                                    switch 1 {
                                                    case 102:
                                                        fmt.Println(v27, p0)
                                                    case 0:
                                                        var v29 uint = 0
                                                        _ = v29
                                                    default:
                                                        v30 := GenS2{F0: 371, F1: 1}
                                                        _ = v30
                                                    }
                                                }
                                                _ = v27
                                                _ = v28
                                            }
                                        case 1:
                                            v18 = v18
                                        case -2147483648:
                                            p0 += 42
                                        }
                                    }
                                    var v31 uint = 802
                                    _ = fi17
                                    _ = v18
                                    _ = v31
                                }
                                for fi32 := 0; fi32 < 15; fi32++ {
                                    if (!(4294967295 <= 377)) {
                                        v33 := GenS4{F0: (((73 >= 2147483648) && false) || true), F1: 0}
                                        var v34 uint8 = p0
                                        v34 += p0
                                        _ = v33
                                        _ = v34
                                    }
                                    v35 := GenS3{F0: 2, F1: (p0 + 128), F2: 0}
                                    switch "\n\t" {
                                    case "a":
                                        fmt.Println(fi32, p0)
                                    }
                                    _ = fi32
                                    _ = v35
                                }
                            }
                        } else {
                            p0 = (p0 - (p0 - p0))
                            v36 := []int{}
                            fmt.Println(fi12)
                            _ = v36
                        }
                        var v37 uint64 = 4294967295
                        _ = v37
                    }
                    _ = fi12
                }
            }
        }
    }
    _ = p0
    _ = v5
    return 1
}
func genFn1(p0 int8, p1 int64) uint16 {
    fmt.Println(p0, p1)
    v38 := GenS3{F0: 2, F1: 91, F2: 347}
    _ = p0
    _ = p1
    _ = v38
    return 255
}

func main() {
    var v1 int = 9223372036854775807
    var v2 float64 = (3.14159 / 0.1)
    var v3 bool = false
    var v4 string = ("hello" + ("\n\t" + (("hello" + "a") + "a")))
    for fi39 := 0; fi39 < 2; fi39++ {
        var v40 int = v1
        _ = fi39
        _ = v40
    }
    switch ("a" + ("\n\t" + v4)) {
    case "hello":
        for fi41 := 0; fi41 < 1; fi41++ {
            if ("" < v4) {
                if v3 {
                    var v42 int16 = 32767
                    switch fi41 {
                    case 128:
                        var v43 uint8 = 185
                        _ = v43
                    case 32767:
                        v44 := []int{(-(v1)), genFn0(127), fi41, (v1 - (fi41 - (-(fi41))))}
                        _ = v44
                    case 42:
                        fmt.Println(v4, v3, fi41)
                    }
                    _ = v42
                }
            }
            if (42 <= -17) {
                v45 := GenS5{F0: "a"}
                var v46 int8 = 127
                var v47 int8 = (v46 * 99)
                _ = v45
                _ = v46
                _ = v47
            } else {
                switch fi41 {
                case -18:
                    if true {
                        v48 := []int{0, fi41, 32767, 9223372036854775807}
                        _ = v48
                    }
                default:
                    var v49 int16 = 304
                    _ = v49
                }
                fmt.Println(v4, fi41)
            }
            _ = fi41
        }
    case "":
        switch (v4 + (((("a" + v4) + v4) + ("foo bar" + (v4 + ""))) + "hello")) {
        case "\n\t":
            v3 = (839 >= 42)
        case "":
            switch genFn0(104) {
            case 2147483647:
                for fi50 := 0; fi50 < 16; fi50++ {
                    if v3 {
                        var v51 int16 = 42
                        switch v4 {
                        case "hello":
                            v3 = (65535 != 42)
                        case "a":
                            switch v4 {
                            case "foo bar":
                                switch "foo bar" {
                                case "hello":
                                    if v3 {
                                        for fi52 := 0; fi52 < 11; fi52++ {
                                            if (-1 > 1) {
                                                fmt.Println(fi52, v51)
                                                var v53 uint64 = 0
                                                fmt.Println(v1)
                                                _ = v53
                                            }
                                            switch v4 {
                                            case "hello":
                                                var v54 int16 = v51
                                                _ = v54
                                            case "":
                                                v55 := GenS4{F0: false, F1: 0}
                                                _ = v55
                                            }
                                            v56 := GenS6{F0: (v3 && true), F1: v3}
                                            _ = fi52
                                            _ = v56
                                        }
                                    } else {
                                        v57 := GenS7{F0: 2147483647, F1: 236, F2: 984}
                                        if true {
                                            var v58 int8 = 0
                                            switch v4 {
                                            case "":
                                                v59 := GenS5{F0: "a"}
                                                _ = v59
                                            case "foo bar":
                                                v1 = (fi50 + (fi50 - v1))
                                            case "\n\t":
                                                if false {
                                                    switch fi50 {
                                                    case 1:
                                                        if (v3 || (false && true)) {
                                                            for fi60 := 0; fi60 < 2; fi60++ {
                                                                switch v4 {
                                                                case "foo bar":
                                                                    var v61 uint8 = 143
                                                                    _ = v61
                                                                case "\n\t":
                                                                    fmt.Println(v2, v4, v1)
                                                                case "hello":
                                                                    switch v4 {
                                                                    case "a":
                                                                        if ((v51 + (v51 << ((255) & 63))) == 314) {
                                                                            if v3 {
                                                                                var v62 int64 = -9223372036854775808
                                                                                _ = v62
                                                                            } else {
                                                                                v63 := GenS4{F0: (90 != 32767), F1: (v58 / ((v58) | 1))}
                                                                                if true {
                                                                                    v64 := GenS7{F0: 0, F1: 80, F2: 1}
                                                                                    if v3 {
                                                                                        var v65 uint16 = 255
                                                                                        _ = v65
                                                                                    }
                                                                                    v3 = v3
                                                                                    _ = v64
                                                                                } else {
                                                                                    v66 := []int{-2147483648, v1, v1}
                                                                                    if v3 {
                                                                                        v67 := GenS3{F0: 697, F1: 255, F2: v51}
                                                                                        _ = len(v66)
                                                                                        switch v4 {
                                                                                        case "foo bar":
                                                                                            fmt.Println(v4, fi60, fi50)
                                                                                        }
                                                                                        _ = v67
                                                                                    }
                                                                                    _ = v66
                                                                                }
                                                                                v68 := GenS2{F0: 1, F1: (^(v58))}
                                                                                _ = v63
                                                                                _ = v68
                                                                            }
                                                                            switch v4 {
                                                                            case "\n\t":
                                                                                v1 = (fi60 << ((4294967295) & 63))
                                                                            case "foo bar":
                                                                                var v69 int32 = -769
                                                                                _ = v69
                                                                            default:
                                                                                v70 := []int{}
                                                                                _ = v70
                                                                            }
                                                                            v3 = true
                                                                        }
                                                                    }
                                                                }
                                                                _ = fi60
                                                            }
                                                            v71 := GenS7{F0: -253, F1: 128, F2: 775}
                                                            v72 := GenS1{F0: 513}
                                                            _ = v71
                                                            _ = v72
                                                        }
                                                    case 2147483647:
                                                        v73 := GenS8{F0: genFn0(127), F1: -2147483648}
                                                        _ = v73
                                                    case -118:
                                                        switch (v4 + v4) {
                                                        case "":
                                                            var v74 uint8 = 0
                                                            _ = v74
                                                        case "hello":
                                                            for fi75 := 0; fi75 < 7; fi75++ {
                                                                var v76 string = ("\n\t" + (((v4 + v4) + (("hello" + "foo bar") + "hello")) + v4))
                                                                if (-25 > v58) {
                                                                    switch "a" {
                                                                    case "hello":
                                                                        var v77 string = v4
                                                                        var v78 int = fi75
                                                                        switch (^(fi75)) {
                                                                        case 32768:
                                                                            fmt.Println(v58, v51)
                                                                        case 254:
                                                                            v78 -= fi75
                                                                        case 1:
                                                                            v76 = v4
                                                                        default:
                                                                            v78 += fi50
                                                                        }
                                                                        var v79 uint16 = 42
                                                                        switch (fi50 * (^(v78))) {
                                                                        case 149:
                                                                            v76 = "\n\t"
                                                                        default:
                                                                            v80 := []int{(v78 % ((-701) | 1)), fi50}
                                                                            _ = v80
                                                                        }
                                                                        switch genFn0(127) {
                                                                        case -2147483648:
                                                                            v81 := []int{}
                                                                            _ = v81
                                                                        }
                                                                        fmt.Println(fi50, v79)
                                                                        var v82 int32 = -711
                                                                        _ = v77
                                                                        _ = v78
                                                                        _ = v79
                                                                        _ = v82
                                                                    case "":
                                                                        switch (v1 * fi75) {
                                                                        case -9223372036854775808:
                                                                            var v83 uint = 255
                                                                            _ = v83
                                                                        case 128:
                                                                            var v84 float32 = (3.14159 * 0.0)
                                                                            _ = v84
                                                                        default:
                                                                            var v85 bool = (!(v3 && (2147483647 < 0)))
                                                                            _ = v85
                                                                        }
                                                                    default:
                                                                        v2 -= (v2 * v2)
                                                                    }
                                                                } else {
                                                                    if (0.0 == float32((v2 - (0.0 * 0.0)))) {
                                                                        var v86 uint = 2
                                                                        v87 := GenS9{F0: v51, F1: 9223372036854775807, F2: uint32(v86)}
                                                                        v76 = v4
                                                                        var v88 float32 = 1.5
                                                                        _ = v86
                                                                        _ = v87
                                                                        _ = v88
                                                                    } else {
                                                                        fmt.Println(v3)
                                                                    }
                                                                    switch (fi50 / ((genFn0(0)) | 1)) {
                                                                    case 32768:
                                                                        switch fi75 {
                                                                        case 4294967295:
                                                                            switch (^(v1)) {
                                                                            case 128:
                                                                                v3 = (!(4294967295 >= 4294967295))
                                                                            }
                                                                        case 128:
                                                                            var v89 uint8 = 82
                                                                            _ = v89
                                                                        case 792:
                                                                            switch (^(v1)) {
                                                                            case 0:
                                                                                var v90 float32 = 0.0
                                                                                _ = v90
                                                                            case 32768:
                                                                                var v91 string = (v76 + v76)
                                                                                _ = v91
                                                                            case -991:
                                                                                var v92 float32 = (3.14159 + ((float32(1.5) - ((((0.1 / 3.14159) + ((0.1 - ((float32(3.14159) - 1.5) - float32(((v2 + float64((-2.25 + 0.1))) + 3.14159)))) - -2.25)) * float32((0.1 / v2))) - (0.0 * (3.14159 - (0.1 - 0.1))))) / 1.0))
                                                                                _ = v92
                                                                            }
                                                                        default:
                                                                            v93 := []int{(fi50 - v1), (v1 + genFn0(127)), -2147483648, 127}
                                                                            _ = v93
                                                                        }
                                                                    case 1:
                                                                        switch 0 {
                                                                        case -1:
                                                                            v94 := []int{v1}
                                                                            var v95 string = (v76 + "")
                                                                            var v96 float32 = 0.1
                                                                            _ = v94
                                                                            _ = v95
                                                                            _ = v96
                                                                        case 32767:
                                                                            v3 = true
                                                                        case -526:
                                                                            v58 += v58
                                                                            v58 = (v58 * (v58 + 100))
                                                                            var v97 uint64 = 334
                                                                            if v3 {
                                                                                v51 -= (v51 >> ((255) & 63))
                                                                                v51 = v51
                                                                                switch int(v51) {
                                                                                case 32768:
                                                                                    v98 := GenS10{F0: v3, F1: -698, F2: 1}
                                                                                    _ = v98
                                                                                case 9223372036854775807:
                                                                                    v4 = v76
                                                                                }
                                                                                v99 := GenS6{F0: v3, F1: v3}
                                                                                _ = v99
                                                                            }
                                                                            _ = v97
                                                                        }
                                                                    case -9223372036854775808:
                                                                        var v100 int = genFn0(195)
                                                                        if (v51 > 42) {
                                                                            var v101 float32 = 0.0
                                                                            switch (v100 & v100) {
                                                                            case 624:
                                                                                v102 := []int{(v100 << ((18446744073709551615) & 63)), (^(v1))}
                                                                                fmt.Println(v1)
                                                                                v3 = false
                                                                                var v103 bool = (-32768 < 32767)
                                                                                _ = v102
                                                                                _ = v103
                                                                            case -1:
                                                                                v101 = float32(v2)
                                                                            case 0:
                                                                                switch v100 {
                                                                                case 9223372036854775807:
                                                                                    v51 = v51
                                                                                case 32767:
                                                                                    v104 := GenS11{F0: 596}
                                                                                    _ = v104
                                                                                }
                                                                            default:
                                                                                var v105 float64 = (3.14159 / 3.14159)
                                                                                v105 -= 1e10
                                                                                if (v105 != -2.25) {
                                                                                    v106 := []int{(v100 % ((genFn0(109)) | 1)), fi75}
                                                                                    v107 := GenS7{F0: 1, F1: 255, F2: 4294967295}
                                                                                    if false {
                                                                                        var v108 float64 = v2
                                                                                        _ = v108
                                                                                    }
                                                                                    _ = v106
                                                                                    _ = v107
                                                                                } else {
                                                                                    var v109 int16 = (v51 * int16(v58))
                                                                                    if true {
                                                                                        fmt.Println(v76, v105)
                                                                                        var v110 int8 = 1
                                                                                        v101 += (1.5 - float32((v105 * 1e10)))
                                                                                        v111 := []int{int(v109)}
                                                                                        switch v4 {
                                                                                        case "foo bar":
                                                                                            if v3 {
                                                                                                fmt.Println(v51)
                                                                                                v112 := GenS5{F0: "hello"}
                                                                                                _ = v112
                                                                                            } else {
                                                                                                if true {
                                                                                                    fmt.Println(v101, v109)
                                                                                                    v76 = (v76 + v76)
                                                                                                }
                                                                                                v113 := GenS9{F0: (v51 % (((-(v51))) | 1)), F1: -107, F2: 67}
                                                                                                _ = v113
                                                                                            }
                                                                                        case "a":
                                                                                            v4 = ""
                                                                                        case "\n\t":
                                                                                            v110 = v58
                                                                                        }
                                                                                        _ = v110
                                                                                        _ = v111
                                                                                    }
                                                                                    switch v76 {
                                                                                    case "":
                                                                                        var v114 float32 = (((-2.25 * (float32(v105) / (0.1 + 1.5))) * v101) / v101)
                                                                                        _ = v114
                                                                                    case "\n\t":
                                                                                        var v115 bool = (v3 || true)
                                                                                        _ = v115
                                                                                    }
                                                                                    _ = v109
                                                                                }
                                                                                _ = v105
                                                                            }
                                                                            _ = v101
                                                                        } else {
                                                                            fmt.Println(v3, v4, v2)
                                                                        }
                                                                        _ = v100
                                                                    default:
                                                                        if v3 {
                                                                            fmt.Println(v58, fi75, v3)
                                                                            v3 = (true && v3)
                                                                        } else {
                                                                            var v116 int = (-(fi50))
                                                                            var v117 int32 = 42
                                                                            v118 := GenS11{F0: 794}
                                                                            _ = v116
                                                                            _ = v117
                                                                            _ = v118
                                                                        }
                                                                        fmt.Println(v1, v2, v1)
                                                                        v119 := []int{}
                                                                        _ = v119
                                                                    }
                                                                    v1 = fi50
                                                                }
                                                                _ = fi75
                                                                _ = v76
                                                            }
                                                        case "foo bar":
                                                            v120 := GenS3{F0: 740, F1: 128, F2: -1}
                                                            _ = v120
                                                        default:
                                                            switch "foo bar" {
                                                            case "\n\t":
                                                                for fi121 := 0; fi121 < 7; fi121++ {
                                                                    var v122 uint16 = 65535
                                                                    if (v3 || v3) {
                                                                        v123 := GenS1{F0: 2}
                                                                        if v3 {
                                                                            var v124 float64 = 1e10
                                                                            _ = v124
                                                                        } else {
                                                                            v125 := []int{128, fi50, genFn0(42), fi121}
                                                                            v126 := GenS10{F0: (genFn1((v58 | (v58 ^ 100)), 948) < (v122 & (v122 ^ v122))), F1: 9223372036854775807, F2: 1}
                                                                            _ = v125
                                                                            _ = v126
                                                                        }
                                                                        _ = v123
                                                                    }
                                                                    _ = fi121
                                                                    _ = v122
                                                                }
                                                            default:
                                                                if false {
                                                                    fmt.Println(v3)
                                                                } else {
                                                                    v4 = (v4 + ((v4 + v4) + v4))
                                                                    if (!(2147483648 < 1)) {
                                                                        var v127 int = fi50
                                                                        v127 = 9223372036854775807
                                                                        var v128 int = v1
                                                                        _ = v127
                                                                        _ = v128
                                                                    }
                                                                }
                                                            }
                                                        }
                                                    }
                                                    switch fi50 {
                                                    case -364:
                                                        var v129 string = v4
                                                        _ = v129
                                                    case -9223372036854775808:
                                                        var v130 int16 = v51
                                                        _ = v130
                                                    case 893:
                                                        v131 := GenS2{F0: 9223372036854775807, F1: int8(fi50)}
                                                        _ = v131
                                                    }
                                                    var v132 float64 = ((1e10 / (-2.25 - v2)) + (v2 + (((v2 + float64(float32(3.14159))) * 1.5) / (1.5 * (float64(0.0) + (float64((0.0 * (1.5 / 1.0))) + v2))))))
                                                    _ = v132
                                                }
                                            default:
                                                if (18446744073709551615 != 2) {
                                                    v2 -= v2
                                                    for fi133 := 0; fi133 < 5; fi133++ {
                                                        v134 := GenS12{F0: 63, F1: 277}
                                                        _ = fi133
                                                        _ = v134
                                                    }
                                                }
                                            }
                                            v2 = float64(0.1)
                                            _ = v58
                                        }
                                        _ = v57
                                    }
                                case "":
                                    var v135 uint64 = 1
                                    _ = v135
                                case "\n\t":
                                    v136 := []int{(v1 >> ((2) & 63)), (fi50 >> ((4294967295) & 63))}
                                    _ = v136
                                }
                            case "\n\t":
                                v51 += v51
                            }
                        default:
                            switch 0 {
                            case 32767:
                                v137 := []int{(-(v1)), (fi50 % ((fi50) | 1))}
                                _ = v137
                            case 127:
                                v2 = v2
                            default:
                                v138 := []int{(-(v1))}
                                _ = v138
                            }
                        }
                        var v139 string = v4
                        _ = v139
                        _ = v51
                    } else {
                        var v140 float64 = ((1.5 + v2) / v2)
                        _ = v140
                    }
                    var v141 bool = (1 != -124)
                    var v142 int8 = 29
                    _ = fi50
                    _ = v141
                    _ = v142
                }
            }
        case "hello":
            v1 += genFn0(42)
        }
    case "foo bar":
        v3 = v3
    }
    for fi143 := 0; fi143 < 19; fi143++ {
        var v144 int = 509
        var v145 int16 = 752
        for fi146 := 0; fi146 < 5; fi146++ {
            if v3 {
                switch (-(v144)) {
                case 9223372036854775807:
                    var v147 float64 = 0.1
                    _ = v147
                default:
                    v145 = v145
                }
                v148 := []int{-2147483648, genFn0(255), (fi143 >> ((4294967295) & 63))}
                v148 = append(v148, (v1 | (v144 - (v144 << ((300) & 63)))))
                var v149 int32 = -1
                switch (((((v4 + v4) + ("" + v4)) + ("a" + v4)) + v4) + v4) {
                case "foo bar":
                    _ = len(v148)
                case "\n\t":
                    var v150 int64 = 200
                    v151 := GenS11{F0: 165}
                    fmt.Println(v149, v3)
                    var v152 string = (v4 + "hello")
                    fmt.Println(v4, v144)
                    v153 := GenS1{F0: 42}
                    _ = v150
                    _ = v151
                    _ = v152
                    _ = v153
                case "hello":
                    if (2147483647 < (v149 ^ v149)) {
                        v1 += fi146
                        var v154 uint64 = 1
                        _ = v154
                    } else {
                        if v3 {
                            switch "a" {
                            case "hello":
                                var v155 int64 = 9223372036854775807
                                if len(v148) > 0 {
                                    v148[0] = 2147483647
                                }
                                if (((3.14159 / (3.14159 / 1.0)) - (v2 - v2)) == (v2 + v2)) {
                                    v156 := GenS12{F0: 0, F1: 119}
                                    _ = v156
                                }
                                _ = v155
                            case "\n\t":
                                if false {
                                    v1 += 2147483647
                                    v157 := GenS10{F0: (v3 || (!false)), F1: 1, F2: 298}
                                    v158 := GenS13{F0: 2, F1: fi146, F2: (fi143 % (((fi143 * genFn0(153))) | 1))}
                                    _ = v157
                                    _ = v158
                                } else {
                                    switch (((v4 + v4) + "foo bar") + v4) {
                                    case "foo bar":
                                        var v159 string = "\n\t"
                                        _ = v159
                                    case "a":
                                        v145 = (v145 / ((42) | 1))
                                    case "hello":
                                        var v160 float64 = 1.5
                                        _ = v160
                                    }
                                    fmt.Println(v144, v145, v4)
                                    var v161 int16 = (-(v145))
                                    v149 += (v149 & v149)
                                    var v162 bool = v3
                                    _ = v161
                                    _ = v162
                                }
                            default:
                                v1 += (v1 * genFn0(128))
                            }
                            switch "foo bar" {
                            case "\n\t":
                                if (42 == 1) {
                                    fmt.Println(v4)
                                } else {
                                    if len(v148) >= 2 {
                                        v148 = v148[1:len(v148)]
                                    }
                                }
                            }
                            if v3 {
                                var v163 string = v4
                                _ = v163
                            }
                            var v164 int8 = 1
                            v144 -= (v1 / ((v144) | 1))
                            v164 -= v164
                            switch fi143 {
                            case 855:
                                if false {
                                    v145 = (v145 - v145)
                                    switch "" {
                                    case "foo bar":
                                        if v3 {
                                            var v165 float32 = 3.14159
                                            _ = v165
                                        } else {
                                            switch genFn0(127) {
                                            case -9223372036854775808:
                                                switch v4 {
                                                case "hello":
                                                    var v166 bool = (0 > 255)
                                                    _ = v166
                                                case "":
                                                    v167 := GenS14{F0: true}
                                                    _ = v167
                                                default:
                                                    if true {
                                                        _ = len(v148)
                                                        v2 = v2
                                                    } else {
                                                        v168 := GenS11{F0: 42}
                                                        _ = v168
                                                    }
                                                }
                                            case -728:
                                                var v169 uint8 = 76
                                                _ = v169
                                            case 224:
                                                var v170 int8 = (v164 + v164)
                                                v171 := GenS9{F0: (v145 << ((2) & 63)), F1: -96, F2: 2}
                                                _ = v170
                                                _ = v171
                                            default:
                                                v149 = v149
                                            }
                                        }
                                    }
                                    var v172 uint32 = 2147483648
                                    _ = v172
                                }
                            case -923:
                                v173 := GenS8{F0: (v144 % (((fi146 ^ (-(v1)))) | 1)), F1: (v149 / ((-887) | 1))}
                                _ = v173
                            }
                            v174 := GenS12{F0: 2147483647, F1: 18446744073709551615}
                            var v175 uint = 18446744073709551615
                            v2 += -2.25
                            var v176 uint8 = 1
                            _ = v164
                            _ = v174
                            _ = v175
                            _ = v176
                        } else {
                            switch v4 {
                            case "foo bar":
                                switch genFn0(127) {
                                case 878:
                                    v177 := GenS10{F0: (v145 < v145), F1: 1, F2: 983}
                                    _ = v177
                                case -1:
                                    switch (fi143 - fi146) {
                                    case 4294967295:
                                        v148 = append(v148, (-(v1)))
                                    default:
                                        switch (v144 / (((fi143 % ((fi143) | 1))) | 1)) {
                                        case 127:
                                            fmt.Println(v4, v144, v144)
                                        case -211:
                                            if len(v148) > 0 {
                                                v148[0] = v1
                                            }
                                        case 9223372036854775807:
                                            if len(v148) >= 2 {
                                                v148 = v148[1:len(v148)]
                                            }
                                        default:
                                            if v3 {
                                                if len(v148) >= 2 {
                                                    v148 = v148[1:len(v148)]
                                                }
                                                v178 := GenS3{F0: 485, F1: 1, F2: v145}
                                                v179 := GenS3{F0: 0, F1: 42, F2: 42}
                                                _ = v178
                                                _ = v179
                                            }
                                        }
                                    }
                                }
                            case "\n\t":
                                v3 = v3
                            case "":
                                switch v1 {
                                case 2147483647:
                                    switch (fi143 | 42) {
                                    case 32768:
                                        var v180 int = genFn0(128)
                                        _ = v180
                                    case 4294967295:
                                        switch v144 {
                                        case -2147483648:
                                            var v181 int32 = (v149 - v149)
                                            _ = v181
                                        case 42:
                                            v182 := GenS15{F0: 1}
                                            _ = v182
                                        default:
                                            if (!(v145 >= (v145 / ((800) | 1)))) {
                                                if len(v148) >= 2 {
                                                    v148 = v148[1:len(v148)]
                                                }
                                                var v183 float32 = 0.0
                                                v184 := GenS10{F0: (502 != 42), F1: 752, F2: 2147483647}
                                                _ = v183
                                                _ = v184
                                            }
                                        }
                                    }
                                case 128:
                                    var v185 uint = 255
                                    _ = v185
                                case -2147483648:
                                    var v186 bool = v3
                                    _ = v186
                                }
                                v1 += v144
                            default:
                                v145 -= v145
                            }
                            if false {
                                fmt.Println(v3, v149, fi146)
                            }
                            v187 := GenS16{F0: 154, F1: 1}
                            _ = v187
                        }
                        var v188 uint16 = 0
                        _ = v188
                    }
                default:
                    var v189 bool = (255 <= 141)
                    _ = v189
                }
                _ = v148
                _ = v149
            }
            v1 = (fi146 + (v144 & 335))
            _ = fi146
        }
        _ = fi143
        _ = v144
        _ = v145
    }
    if v3 {
        var v190 uint8 = 54
        var v191 uint32 = 2147483647
        _ = v190
        _ = v191
    }
    v192 := GenS11{F0: 255}
    var v193 uint = 707
    for fi194 := 0; fi194 < 15; fi194++ {
        v193 += v193
        v195 := GenS2{F0: 979, F1: 100}
        v196 := []int{}
        _ = fi194
        _ = v195
        _ = v196
    }
    if true {
        v197 := []int{(^(v1)), genFn0(131), genFn0(255)}
        var v198 int = (v1 | v1)
        _ = v197
        _ = v198
    }
    v199 := GenS1{F0: 2}
    v200 := GenS5{F0: (((("foo bar" + v4) + v4) + "") + v4)}
    v201 := GenS14{F0: ((602 == 2147483647) && v3)}
    if (v3 && (v4 < v4)) {
        for fi202 := 0; fi202 < 18; fi202++ {
            var v203 float64 = (((((0.1 - (float64(((0.1 * 0.1) / 1.0)) * (v2 * 0.0))) * v2) / float64(1.5)) * 1.5) * v2)
            var v204 int8 = 127
            switch v1 {
            case -575:
                v205 := GenS14{F0: (-2147483648 != -1)}
                _ = v205
            case 1:
                v206 := GenS5{F0: ""}
                _ = v206
            case 32768:
                fmt.Println(v203, v4, v4)
            default:
                var v207 uint32 = uint32(v193)
                _ = v207
            }
            _ = fi202
            _ = v203
            _ = v204
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v193:", v193)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v192
    _ = v193
    _ = v199
    _ = v2
    _ = v200
    _ = v201
    _ = v3
    _ = v4
}
