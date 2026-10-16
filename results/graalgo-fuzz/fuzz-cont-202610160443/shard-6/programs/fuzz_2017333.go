package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 string; F1 int8; F2 string }
type GenS4 struct { F0 int; F1 int }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 string; F1 string }
type GenS7 struct { F0 uint; F1 int }

func main() {
    var v1 uint = 2
    var v2 int64 = -1
    var v3 string = ""
    var v4 int16 = 486
    var v5 bool = (true || (1e10 != -2.25))
    if (42 >= 975) {
        fmt.Println(v3, v5)
        if false {
            switch 32767 {
            case 42:
                if (1 != 1) {
                    var v6 string = ""
                    switch 42 {
                    case -629:
                        var v7 uint8 = 1
                        _ = v7
                    default:
                        v8 := GenS1{F0: -145}
                        _ = v8
                    }
                    _ = v6
                }
            case 473:
                var v9 int16 = (v4 + -787)
                _ = v9
            case -57:
                var v10 bool = (v2 == v2)
                _ = v10
            }
            v11 := GenS2{F0: (v4 % ((v4) | 1))}
            _ = v11
        } else {
            if (-97 <= -128) {
                v4 = v4
                for fi12 := 0; fi12 < 1; fi12++ {
                    if (-35 >= 0) {
                        fmt.Println(fi12)
                    } else {
                        var v13 uint64 = 1
                        _ = v13
                    }
                    _ = fi12
                }
            } else {
                if (!v5) {
                    switch v3 {
                    case "foo bar":
                        var v14 int8 = 100
                        _ = v14
                    case "hello":
                        v15 := GenS3{F0: v3, F1: 39, F2: v3}
                        _ = v15
                    case "a":
                        var v16 int8 = 57
                        _ = v16
                    default:
                        var v17 int64 = 42
                        _ = v17
                    }
                } else {
                    switch "hello" {
                    case "a":
                        v18 := []int{515}
                        _ = v18
                    default:
                        v19 := []int{387, 9223372036854775807}
                        _ = v19
                    }
                    var v20 string = (("foo bar" + v3) + "hello")
                    if ((false && (v5 || (!false))) && (!v5)) {
                        v21 := []int{127, 4294967295, 127, -967}
                        _ = v21
                    }
                    _ = v20
                }
                v2 += (v2 >> ((18446744073709551615) & 63))
            }
            if false {
                var v22 int8 = 100
                _ = v22
            }
            if v5 {
                v2 = 601
                var v23 int = 2147483647
                for fi24 := 0; fi24 < 19; fi24++ {
                    if true {
                        var v25 uint64 = 0
                        if v5 {
                            switch (((v3 + ("\n\t" + "hello")) + v3) + v3) {
                            case "foo bar":
                                v4 = v4
                            default:
                                for fi26 := 0; fi26 < 14; fi26++ {
                                    fmt.Println(v5, v4)
                                    fmt.Println(v3)
                                    _ = fi26
                                }
                            }
                            v27 := []int{(v23 * v23), v23, (fi24 | v23)}
                            var v28 uint16 = 65535
                            _ = v27
                            _ = v28
                        } else {
                            var v29 uint32 = 42
                            _ = v29
                        }
                        v30 := GenS4{F0: 776, F1: fi24}
                        _ = v25
                        _ = v30
                    }
                    var v31 string = (v3 + v3)
                    v32 := []int{-2147483648, 32767}
                    _ = fi24
                    _ = v31
                    _ = v32
                }
                _ = v23
            } else {
                var v33 uint64 = 18446744073709551615
                _ = v33
            }
        }
    }
    var v34 uint16 = 715
    switch ((v3 + "") + (v3 + (v3 + v3))) {
    case "hello":
        switch v3 {
        case "a":
            for fi35 := 0; fi35 < 10; fi35++ {
                v34 = 211
                _ = fi35
            }
        }
    case "foo bar":
        switch 9223372036854775807 {
        case -129:
            v36 := []int{-87, 454}
            _ = v36
        case 306:
            if ((v2 ^ -15) >= (v2 << ((v1) & 63))) {
                if (!v5) {
                    switch v3 {
                    case "a":
                        for fi37 := 0; fi37 < 15; fi37++ {
                            switch (v3 + (v3 + v3)) {
                            case "a":
                                if v5 {
                                    if true {
                                        v5 = (!false)
                                        if (4 == -1) {
                                            if (2147483648 >= 42) {
                                                if (!(!v5)) {
                                                    var v38 int16 = -32768
                                                    for fi39 := 0; fi39 < 3; fi39++ {
                                                        var v40 float32 = (-2.25 + 3.14159)
                                                        v1 = (^(v1))
                                                        v41 := []int{}
                                                        _ = fi39
                                                        _ = v40
                                                        _ = v41
                                                    }
                                                    for fi42 := 0; fi42 < 13; fi42++ {
                                                        v43 := []int{2147483647}
                                                        _ = fi42
                                                        _ = v43
                                                    }
                                                    _ = v38
                                                }
                                                for fi44 := 0; fi44 < 18; fi44++ {
                                                    v45 := []int{-1}
                                                    v3 = ""
                                                    if v5 {
                                                        v5 = v5
                                                        v1 = uint(v34)
                                                    }
                                                    _ = fi44
                                                    _ = v45
                                                }
                                                fmt.Println(v3, fi37)
                                            }
                                            var v46 float64 = -2.25
                                            v47 := []int{fi37, (-(fi37))}
                                            _ = v46
                                            _ = v47
                                        }
                                        if (true || v5) {
                                            var v48 int16 = v4
                                            _ = v48
                                        }
                                    }
                                    var v49 int8 = -128
                                    v50 := []int{fi37, (fi37 * fi37)}
                                    _ = v49
                                    _ = v50
                                } else {
                                    v51 := GenS4{F0: fi37, F1: fi37}
                                    _ = v51
                                }
                            default:
                                v52 := []int{fi37, -9223372036854775808, fi37, fi37}
                                _ = v52
                            }
                            _ = fi37
                        }
                    }
                    switch v3 {
                    case "foo bar":
                        v4 -= 0
                    }
                    if v5 {
                        var v53 int = -9223372036854775808
                        if true {
                            for fi54 := 0; fi54 < 12; fi54++ {
                                var v55 int64 = (v2 - (v2 + v2))
                                _ = fi54
                                _ = v55
                            }
                            v3 = v3
                        }
                        v34 -= (v34 >> (((v1 >> (((v1 ^ 42)) & 63))) & 63))
                        _ = v53
                    } else {
                        switch ("\n\t" + v3) {
                        case "a":
                            var v56 int = 32768
                            _ = v56
                        case "\n\t":
                            var v57 int8 = -41
                            _ = v57
                        default:
                            v58 := GenS4{F0: 42, F1: 128}
                            _ = v58
                        }
                        switch 49 {
                        case 1:
                            var v59 int8 = -128
                            _ = v59
                        default:
                            v2 = v2
                        }
                        if true {
                            v1 += v1
                        } else {
                            v60 := GenS5{F0: ((float32(1.5) * ((float32(float64(0.0)) + -2.25) / (1e10 + ((0.1 - 1.5) + (0.0 * 1.5))))) > 1e10)}
                            _ = v60
                        }
                    }
                }
            } else {
                switch (("foo bar" + v3) + v3) {
                case "a":
                    switch 32768 {
                    case -720:
                        for fi61 := 0; fi61 < 17; fi61++ {
                            v62 := []int{(fi61 ^ fi61), fi61, 4294967295}
                            if (-2.25 > 0.0) {
                                if len(v62) >= 2 {
                                    v62 = v62[1:len(v62)]
                                }
                                fmt.Println(v4, v34)
                            }
                            _ = fi61
                            _ = v62
                        }
                    case 128:
                        switch v3 {
                        case "foo bar":
                            var v63 uint32 = 4294967295
                            _ = v63
                        case "":
                            v64 := []int{42, int(v4), -2147483648}
                            _ = v64
                        case "a":
                            v65 := []int{32768, -9223372036854775808, 9223372036854775807, 817}
                            _ = v65
                        }
                    }
                case "hello":
                    var v66 uint64 = 234
                    _ = v66
                default:
                    v3 = "hello"
                }
            }
        }
    default:
        var v67 float32 = 0.1
        _ = v67
    }
    for fi68 := 0; fi68 < 8; fi68++ {
        fmt.Println(v34, v34, v3)
        var v69 float64 = (0.0 - 1e10)
        switch (v3 + v3) {
        case "a":
            fmt.Println(v69, v4)
        case "foo bar":
            var v70 float32 = (-2.25 / 1.5)
            _ = v70
        case "hello":
            var v71 int32 = 1
            _ = v71
        }
        _ = fi68
        _ = v69
    }
    v2 += v2
    v72 := GenS2{F0: v4}
    v73 := GenS2{F0: (v4 / ((v4) | 1))}
    for fi74 := 0; fi74 < 19; fi74++ {
        switch fi74 {
        case 0:
            v75 := GenS2{F0: -32768}
            _ = v75
        case -572:
            var v76 int32 = 1
            _ = v76
        default:
            switch v3 {
            case "hello":
                if v5 {
                    for fi77 := 0; fi77 < 15; fi77++ {
                        v78 := []int{9223372036854775807}
                        var v79 uint = v1
                        var v80 uint32 = 1
                        _ = fi77
                        _ = v78
                        _ = v79
                        _ = v80
                    }
                    switch "hello" {
                    case "":
                        var v81 string = (((v3 + "foo bar") + (((v3 + "a") + "a") + v3)) + (("a" + v3) + ("" + v3)))
                        _ = v81
                    case "a":
                        if (1e10 > ((1e10 / ((0.1 + 1e10) + 0.0)) + 3.14159)) {
                            fmt.Println(fi74, v4, fi74)
                            if ((!v5) || (!(-2.25 != (-2.25 * 3.14159)))) {
                                v82 := GenS2{F0: v4}
                                if (v5 || false) {
                                    v3 = ("hello" + v3)
                                    switch v3 {
                                    case "":
                                        v2 = 80
                                    default:
                                        v2 -= v2
                                    }
                                } else {
                                    var v83 uint8 = 127
                                    _ = v83
                                }
                                v1 = 255
                                _ = v82
                            } else {
                                v34 = v34
                                fmt.Println(fi74, v1, v3)
                                v1 += 4294967295
                            }
                            v3 = (v3 + (v3 + v3))
                        } else {
                            v84 := []int{fi74, (fi74 / (((fi74 * fi74)) | 1)), (fi74 * fi74), (-(fi74))}
                            if (-680 >= 1000) {
                                if true {
                                    v84 = append(v84, 128)
                                    v85 := GenS5{F0: true}
                                    v86 := GenS3{F0: "", F1: -13, F2: ("a" + (v3 + v3))}
                                    _ = v85
                                    _ = v86
                                }
                            } else {
                                var v87 float64 = 0.0
                                _ = v87
                            }
                            _ = v84
                        }
                    }
                } else {
                    v4 -= (v4 * (v4 & -32768))
                    v88 := []int{fi74, 2147483647, (fi74 % ((-159) | 1))}
                    var v89 float64 = ((1e10 / 1.0) - ((0.1 - 1.5) / -2.25))
                    _ = v88
                    _ = v89
                }
            default:
                if ((uint16(v1) == 237) || true) {
                    switch fi74 {
                    case 161:
                        switch (fi74 * -811) {
                        case 0:
                            v4 += (v4 - v4)
                        case 42:
                            var v90 bool = false
                            _ = v90
                        }
                    case 1:
                        var v91 uint16 = v34
                        _ = v91
                    case 128:
                        var v92 int16 = (v4 + 1000)
                        _ = v92
                    }
                }
            }
        }
        v93 := GenS2{F0: (v4 ^ v4)}
        _ = fi74
        _ = v93
    }
    if v5 {
        if ((((v3 + "hello") + "") >= v3) && (18446744073709551615 < v1)) {
            var v94 uint32 = 249
            _ = v94
        } else {
            for fi95 := 0; fi95 < 11; fi95++ {
                if v5 {
                    var v96 int16 = v4
                    _ = v96
                }
                for fi97 := 0; fi97 < 5; fi97++ {
                    v3 = "hello"
                    fmt.Println(fi97)
                    var v98 uint16 = 255
                    _ = fi97
                    _ = v98
                }
                _ = fi95
            }
            v5 = (0.1 >= 0.1)
            v99 := GenS2{F0: 1}
            _ = v99
        }
    }
    v100 := GenS2{F0: 1}
    switch 1 {
    case -1:
        v101 := []int{}
        _ = v101
    case 118:
        if ((v4 & (v4 - 42)) <= (v4 - v4)) {
            if v5 {
                var v102 int32 = -527
                _ = v102
            }
        } else {
            switch "hello" {
            case "hello":
                v103 := GenS6{F0: ("\n\t" + "\n\t"), F1: ((("foo bar" + ("foo bar" + ((v3 + v3) + ""))) + v3) + "hello")}
                _ = v103
            }
            var v104 uint64 = 2
            _ = v104
        }
    case 15:
        var v105 uint8 = 2
        _ = v105
    default:
        v1 = (v1 - v1)
    }
    v34 -= 46
    v34 = 2
    v106 := []int{}
    v107 := GenS7{F0: (v1 << ((255) & 63)), F1: 0}
    for fi108 := 0; fi108 < 9; fi108++ {
        var v109 uint16 = 2
        var v110 bool = (true || true)
        if len(v106) > 0 {
            v106[0] = (-(fi108))
        }
        _ = fi108
        _ = v109
        _ = v110
    }
    for fi111 := 0; fi111 < 16; fi111++ {
        for fi112 := 0; fi112 < 19; fi112++ {
            var v113 uint32 = uint32(v34)
            _ = fi112
            _ = v113
        }
        _ = fi111
    }
    if (4294967295 <= 4294967295) {
        if len(v106) >= 2 {
            v106 = v106[1:len(v106)]
        }
        var v114 uint16 = 0
        _ = v114
    } else {
        v115 := GenS4{F0: 32768, F1: 42}
        if ((!v5) && v5) {
            var v116 float32 = -2.25
            _ = v116
        } else {
            if (!false) {
                fmt.Println(v3, v3, v1)
            } else {
                v117 := GenS6{F0: v3, F1: (v3 + ("" + v3))}
                v3 = v3
                var v118 uint = (v1 | v1)
                _ = v117
                _ = v118
            }
            v119 := GenS1{F0: 0}
            _ = v119
        }
        _ = v115
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v106:", len(v106), v106)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v100
    _ = v106
    _ = v107
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v5
    _ = v72
    _ = v73
}
