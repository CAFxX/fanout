package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint; F1 int16 }
type GenS4 struct { F0 int32; F1 int; F2 uint32 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 int32 }
type GenS9 struct { F0 uint; F1 uint64 }
type GenS10 struct { F0 uint64; F1 uint32; F2 int8 }
type GenS11 struct { F0 int32; F1 bool; F2 string }
type GenS12 struct { F0 uint; F1 uint8; F2 uint32 }
type GenS13 struct { F0 uint16; F1 int32; F2 int16 }
type GenS14 struct { F0 int16; F1 uint32 }
type GenS15 struct { F0 uint8; F1 uint64; F2 uint }
type GenS16 struct { F0 int8 }
func genFn0(p0 uint64) uint8 {
    v5 := []int{-40, 0, 0, 2147483647}
    p0 = p0
    p0 += 18446744073709551615
    _ = p0
    _ = v5
    return 127
}

func main() {
    var v1 uint = 0
    v2 := make(map[int]int)
    var v3 uint8 = 75
    var v4 bool = (false || false)
    if (false || (!((!false) || (((0 < 757) && (19 == 9223372036854775807)) || true)))) {
        for fi6 := 0; fi6 < 11; fi6++ {
            v2[fi6] = fi6
            _ = fi6
        }
    } else {
        var v7 int = -564
        _ = v7
    }
    fmt.Println(v1)
    fmt.Println(v1, v1)
    if (42 < 406) {
        v3 = 1
        if false {
            var v8 int64 = 0
            if (!v4) {
                var v9 bool = false
                for fi10 := 0; fi10 < 8; fi10++ {
                    var v11 int32 = -64
                    var v12 uint16 = 0
                    _ = fi10
                    _ = v11
                    _ = v12
                }
                _ = v9
            } else {
                v4 = v4
                if v4 {
                    if true {
                        for fi13 := 0; fi13 < 3; fi13++ {
                            fmt.Println(v3, v1)
                            _ = fi13
                        }
                    }
                }
            }
            v3 += v3
            _ = v8
        }
    }
    if ((float32(((0.1 / float64(3.14159)) * 0.0)) * (float32(0.1) * float32(float64(((1e10 + (float32(0.0) + 0.1)) - 0.0))))) >= float32(0.0)) {
        for fi14 := 0; fi14 < 2; fi14++ {
            var v15 uint = (v1 * v1)
            if (-2.25 != 1.5) {
                v3 = (v3 << (((v1 - 141)) & 63))
                for fi16 := 0; fi16 < 9; fi16++ {
                    if v4 {
                        var v17 int64 = 9223372036854775807
                        _ = v17
                    }
                    switch ("hello" + "hello") {
                    case "foo bar":
                        var v18 int8 = 0
                        if (!v4) {
                            fmt.Println(fi14)
                            v3 -= v3
                        } else {
                            var v19 float32 = float32(float64(1.5))
                            _ = v19
                        }
                        _ = v18
                    case "\n\t":
                        var v20 uint16 = 255
                        _ = v20
                    case "a":
                        v1 = v15
                    }
                    v21 := []int{fi14, fi16, (^(fi16)), (-(fi16))}
                    _ = len(v21)
                    var v22 uint8 = (v3 ^ v3)
                    v1 += 0
                    v3 -= v22
                    _ = fi16
                    _ = v21
                    _ = v22
                }
            } else {
                v3 += (v3 << ((v15) & 63))
                v3 -= 127
                if true {
                    v23 := []int{(^(fi14)), (-(fi14))}
                    for fi24 := 0; fi24 < 4; fi24++ {
                        if v4 {
                            if false {
                                var v25 uint64 = 0
                                v23 = append(v23, int(v25))
                                _ = v25
                            }
                            if v4 {
                                if (1e10 < 1.5) {
                                    switch "a" {
                                    case "hello":
                                        var v26 uint8 = v3
                                        delete(v2, 0)
                                        _ = v26
                                    case "\n\t":
                                        var v27 uint8 = (v3 + (v3 | v3))
                                        var v28 int64 = 1
                                        switch "hello" {
                                        case "\n\t":
                                            switch (fi24 % (((^(fi14))) | 1)) {
                                            case 0:
                                                var v29 int32 = 0
                                                _ = v29
                                            case -812:
                                                v30 := GenS1{F0: 18446744073709551615, F1: 2}
                                                _ = v30
                                            }
                                        case "":
                                            v3 = 57
                                        }
                                        _ = v27
                                        _ = v28
                                    default:
                                        v3 = (v3 + (v3 << ((v1) & 63)))
                                    }
                                } else {
                                    fmt.Println(v15, fi14, v3)
                                    v4 = (42 != 2147483647)
                                }
                                var v31 uint = (v1 - 65535)
                                v23 = append(v23, (fi14 * 4294967295))
                                _ = v31
                            } else {
                                if v4 {
                                    switch "" {
                                    case "foo bar":
                                        if v4 {
                                            v15 = (v1 & v15)
                                            v15 += (v15 * (v15 & (v15 >> ((255) & 63))))
                                            if len(v23) > 0 {
                                                v23[0] = 572
                                            }
                                        } else {
                                            var v32 int8 = 1
                                            _ = v32
                                        }
                                    case "hello":
                                        var v33 uint = 4294967295
                                        _ = v33
                                    case "":
                                        if v4 {
                                            v34 := GenS1{F0: 779, F1: 448}
                                            var v35 uint64 = 4294967295
                                            _ = v34
                                            _ = v35
                                        } else {
                                            if (2147483647 <= -29) {
                                                var v36 string = ("" + "a")
                                                var v37 uint8 = (v3 >> ((1) & 63))
                                                v1 = v1
                                                _ = v36
                                                _ = v37
                                            } else {
                                                v4 = (179 == 0)
                                            }
                                            if v4 {
                                                var v38 int8 = -76
                                                v4 = v4
                                                _ = v38
                                            }
                                            if v4 {
                                                switch "\n\t" {
                                                case "foo bar":
                                                    if len(v23) >= 2 {
                                                        v23 = v23[1:len(v23)]
                                                    }
                                                case "hello":
                                                    v15 = (v1 >> (((v15 | (v1 >> ((255) & 63)))) & 63))
                                                }
                                                v39 := GenS1{F0: 42, F1: 2}
                                                _ = v39
                                            } else {
                                                if (v4 && (v4 && v4)) {
                                                    v40 := GenS1{F0: 567, F1: 42}
                                                    _ = v40
                                                } else {
                                                    var v41 int8 = 1
                                                    _ = v41
                                                }
                                                if false {
                                                    if (v4 || v4) {
                                                        v15 += (v1 % (((v1 / ((42) | 1))) | 1))
                                                        var v42 int64 = 1
                                                        _ = v42
                                                    } else {
                                                        var v43 int32 = -1
                                                        var v44 int16 = 1000
                                                        v45 := GenS1{F0: 4294967295, F1: 622}
                                                        _ = v43
                                                        _ = v44
                                                        _ = v45
                                                    }
                                                }
                                            }
                                        }
                                    }
                                    var v46 int32 = 2147483647
                                    v23 = append(v23, 32767)
                                    _ = v46
                                } else {
                                    v47 := GenS1{F0: 865, F1: 994}
                                    v48 := GenS2{F0: (fi14 * (fi14 << (((v1 + v1)) & 63)))}
                                    v1 -= v1
                                    switch fi24 {
                                    case 128:
                                        if false {
                                            var v49 uint16 = 42
                                            _ = v49
                                        }
                                    case -1:
                                        v50 := GenS1{F0: 699, F1: 1}
                                        _ = v50
                                    case -67:
                                        v51 := GenS2{F0: fi24}
                                        _ = v51
                                    }
                                    _ = v47
                                    _ = v48
                                }
                                switch fi24 {
                                case -197:
                                    var v52 uint32 = 2
                                    _ = v52
                                case -1:
                                    fmt.Println(v15)
                                case 4294967295:
                                    delete(v2, (fi14 % ((fi24) | 1)))
                                    if len(v23) >= 2 {
                                        v23 = v23[1:len(v23)]
                                    }
                                    if v4 {
                                        switch (fi14 & (^(fi14))) {
                                        case -562:
                                            v53 := GenS1{F0: 1, F1: 1}
                                            _ = v53
                                        default:
                                            v54 := GenS3{F0: v1, F1: 0}
                                            _ = v54
                                        }
                                        switch "a" {
                                        case "hello":
                                            v1 = (v15 << (((v15 + (v15 << ((65535) & 63)))) & 63))
                                        }
                                        switch ("foo bar" + "") {
                                        case "a":
                                            if (-119 == int8(v3)) {
                                                v2[0] = fi24
                                            } else {
                                                v3 += v3
                                                var v55 float32 = (float32(0.1) - 1e10)
                                                v56 := GenS4{F0: -1, F1: (fi14 >> ((v15) & 63)), F2: 4294967295}
                                                _ = v55
                                                _ = v56
                                            }
                                        case "":
                                            var v57 int32 = -1
                                            _ = v57
                                        default:
                                            v58 := GenS5{F0: -128}
                                            _ = v58
                                        }
                                    } else {
                                        switch "foo bar" {
                                        case "foo bar":
                                            v15 = v15
                                        case "a":
                                            fmt.Println(v1, v3, v1)
                                        }
                                        v59 := GenS4{F0: 0, F1: -2147483648, F2: 2147483648}
                                        _ = v59
                                    }
                                default:
                                    v1 = (v1 & v1)
                                }
                            }
                            if (v4 && false) {
                                var v60 uint64 = 974
                                _ = v60
                            }
                            fmt.Println(v4, v1)
                            v1 = 2
                            v3 += uint8(v15)
                            if v4 {
                                if (float32(3.14159) != 0.1) {
                                    delete(v2, 0)
                                    v15 += v1
                                    _ = len(v2)
                                }
                                delete(v2, 128)
                                if (-32768 == -1) {
                                    if (((^(v3)) < (v3 / (((^(v3))) | 1))) && (!v4)) {
                                        var v61 uint64 = uint64(v3)
                                        v2[fi24] = int(v3)
                                        _ = v61
                                    }
                                    fmt.Println(fi24, v3)
                                }
                                var v62 string = "\n\t"
                                _ = v62
                            }
                        }
                        _ = fi24
                    }
                    _ = v23
                }
            }
            v63 := []int{(fi14 & (fi14 ^ fi14)), fi14, -509}
            _ = fi14
            _ = v15
            _ = v63
        }
        v64 := []int{574, 1, -2147483648, 128}
        if v4 {
            if v4 {
                switch "hello" {
                case "\n\t":
                    v2[42] = 2147483647
                default:
                    var v65 int32 = 877
                    _ = v65
                }
            }
            var v66 uint8 = (^(v3))
            for fi67 := 0; fi67 < 4; fi67++ {
                v3 += 95
                _ = fi67
            }
            _ = v66
        }
        _ = v64
    }
    for fi68 := 0; fi68 < 5; fi68++ {
        var v69 bool = v4
        if true {
            v69 = v69
            if (v69 || v69) {
                if (32767 > 0) {
                    fmt.Println(v3, v3)
                    switch "hello" {
                    case "hello":
                        v70 := GenS2{F0: fi68}
                        _ = v70
                    case "\n\t":
                        if v69 {
                            _ = len(v2)
                            for fi71 := 0; fi71 < 19; fi71++ {
                                _ = len(v2)
                                _ = fi71
                            }
                            if ("a" > ("a" + "foo bar")) {
                                v4 = (!v69)
                                if v4 {
                                    v1 += (v1 + (v1 << ((997) & 63)))
                                } else {
                                    for fi72 := 0; fi72 < 18; fi72++ {
                                        var v73 uint64 = 1
                                        var v74 uint64 = uint64(v1)
                                        v75 := GenS3{F0: (v1 | (^(v1))), F1: int16(fi68)}
                                        _ = fi72
                                        _ = v73
                                        _ = v74
                                        _ = v75
                                    }
                                }
                            } else {
                                v76 := v2[fi68]
                                delete(v2, (fi68 * (fi68 % ((fi68) | 1))))
                                _ = v76
                            }
                        }
                    }
                    switch "a" {
                    case "foo bar":
                        v1 = v1
                    case "":
                        for fi77 := 0; fi77 < 14; fi77++ {
                            fmt.Println(v1, fi68)
                            if (false && false) {
                                var v78 int32 = -605
                                var v79 uint = (v1 * v1)
                                _ = v78
                                _ = v79
                            }
                            fmt.Println(v3, v4, v4)
                            _ = fi77
                        }
                    }
                }
                v80 := GenS4{F0: 253, F1: fi68, F2: 534}
                _ = v80
            }
        } else {
            v69 = v4
            var v81 int = fi68
            _ = v81
        }
        fmt.Println(v3, v1)
        _ = fi68
        _ = v69
    }
    for fi82 := 0; fi82 < 11; fi82++ {
        fmt.Println(v1, v4, v3)
        if v4 {
            v1 = (v1 / ((uint(v3)) | 1))
            fmt.Println(v1)
        } else {
            switch (^(fi82)) {
            case 32768:
                fmt.Println(v4)
            default:
                for fi83 := 0; fi83 < 12; fi83++ {
                    v3 -= 255
                    _ = fi83
                }
            }
            fmt.Println(v4, v1)
        }
        _ = fi82
    }
    switch -910 {
    case 1:
        if v4 {
            fmt.Println(v1)
            v4 = (-1 > -2147483648)
            var v84 uint = v1
            _ = v84
        } else {
            v3 = v3
            var v85 float32 = 0.0
            if (449 <= -2147483648) {
                v3 = v3
                var v86 string = "hello"
                _ = v86
            } else {
                if false {
                    for fi87 := 0; fi87 < 16; fi87++ {
                        for fi88 := 0; fi88 < 16; fi88++ {
                            var v89 int64 = 1
                            v90 := GenS2{F0: (^(fi87))}
                            v3 -= 95
                            v91 := GenS2{F0: (fi87 / ((32767) | 1))}
                            if (602 < 589) {
                                v1 += v1
                                v3 = (^(v3))
                                v2[fi88] = (^(fi87))
                                v1 = 0
                                var v92 uint32 = 275
                                if ((v89 / (((v89 ^ (-(v89)))) | 1)) != 1) {
                                    v93 := GenS2{F0: (fi87 & (fi87 / ((fi88) | 1)))}
                                    switch "a" {
                                    case "hello":
                                        v94 := GenS2{F0: fi88}
                                        _ = v94
                                    case "\n\t":
                                        v85 -= 0.1
                                    }
                                    var v95 bool = false
                                    if false {
                                        var v96 string = ""
                                        v1 = v1
                                        _ = v96
                                    }
                                    switch (fi87 - (fi88 >> ((v1) & 63))) {
                                    case 1:
                                        if v95 {
                                            v89 -= (v89 - (-(v89)))
                                            var v97 uint16 = 255
                                            v98 := GenS6{F0: 0}
                                            _ = v97
                                            _ = v98
                                        }
                                    case 4294967295:
                                        switch "a" {
                                        case "foo bar":
                                            v99 := GenS6{F0: (-(v89))}
                                            _ = v99
                                        case "a":
                                            if v95 {
                                                if ((255 >= 107) && true) {
                                                    switch "a" {
                                                    case "a":
                                                        switch fi87 {
                                                        case 128:
                                                            v100 := []int{int(v92), (fi87 | fi87), (fi88 ^ fi87), fi87}
                                                            _ = v100
                                                        default:
                                                            var v101 uint = v1
                                                            _ = v101
                                                        }
                                                    case "\n\t":
                                                        switch "\n\t" {
                                                        case "hello":
                                                            v1 += (v1 + (^(v1)))
                                                        case "\n\t":
                                                            v102 := []int{(-(fi88))}
                                                            var v103 uint64 = 13
                                                            if false {
                                                                var v104 float32 = v85
                                                                _ = v104
                                                            }
                                                            _ = v102
                                                            _ = v103
                                                        case "a":
                                                            if (1.5 == 1.5) {
                                                                var v105 uint64 = 669
                                                                var v106 uint16 = 255
                                                                var v107 int32 = 2147483647
                                                                v108 := []int{}
                                                                _ = v105
                                                                _ = v106
                                                                _ = v107
                                                                _ = v108
                                                            } else {
                                                                var v109 int64 = v89
                                                                v92 = 42
                                                                v110 := []int{}
                                                                if true {
                                                                    switch (fi88 * (fi88 | (fi88 >> (((v1 << ((v1) & 63))) & 63)))) {
                                                                    case 42:
                                                                        v109 = (^(v109))
                                                                    case 649:
                                                                        var v111 int32 = 506
                                                                        _ = v111
                                                                    case 32768:
                                                                        v112 := GenS6{F0: (v109 - (v89 ^ v89))}
                                                                        _ = v112
                                                                    }
                                                                    switch "a" {
                                                                    case "":
                                                                        var v113 uint = (^(v1))
                                                                        _ = v113
                                                                    case "foo bar":
                                                                        v85 = -2.25
                                                                    case "a":
                                                                        if v4 {
                                                                            if (32767 == 1) {
                                                                                switch fi88 {
                                                                                case -143:
                                                                                    var v114 uint32 = v92
                                                                                    _ = v114
                                                                                }
                                                                            } else {
                                                                                if ((!v4) && (((("hello" + ("hello" + "a")) + ("hello" + (("hello" + "foo bar") + "a"))) + "\n\t") == "\n\t")) {
                                                                                    delete(v2, (fi88 * 127))
                                                                                    if (2 == 42) {
                                                                                        v115 := GenS7{F0: 32767}
                                                                                        v116 := v2[9223372036854775807]
                                                                                        v117 := GenS8{F0: 0}
                                                                                        switch (v116 | fi88) {
                                                                                        case 0:
                                                                                            var v118 uint32 = v92
                                                                                            _ = v118
                                                                                        }
                                                                                        _ = v115
                                                                                        _ = v116
                                                                                        _ = v117
                                                                                    } else {
                                                                                        if true {
                                                                                            v119 := v2[2147483647]
                                                                                            v120 := GenS3{F0: (v1 | 42), F1: -912}
                                                                                            v121 := GenS6{F0: 9223372036854775807}
                                                                                            v2[-1] = (fi88 ^ (fi88 << ((2) & 63)))
                                                                                            _ = v119
                                                                                            _ = v120
                                                                                            _ = v121
                                                                                        } else {
                                                                                            switch 42 {
                                                                                            case 4294967295:
                                                                                                if (!v4) {
                                                                                                    v122 := GenS2{F0: fi87}
                                                                                                    switch fi87 {
                                                                                                    case 32767:
                                                                                                        if (v92 >= 2) {
                                                                                                            if ("a" >= ("" + ("\n\t" + (("foo bar" + ("" + "hello")) + "\n\t")))) {
                                                                                                                fmt.Println(v1, v85)
                                                                                                                v89 -= v89
                                                                                                                var v123 int = (fi88 >> (((v1 * 650)) & 63))
                                                                                                                _ = v123
                                                                                                            }
                                                                                                        } else {
                                                                                                            switch -1 {
                                                                                                            case -308:
                                                                                                                switch "hello" {
                                                                                                                case "a":
                                                                                                                    v95 = (v3 < (^(v3)))
                                                                                                                case "hello":
                                                                                                                    fmt.Println(v4, v92)
                                                                                                                default:
                                                                                                                    var v124 int64 = v89
                                                                                                                    _ = v124
                                                                                                                }
                                                                                                            case -9223372036854775808:
                                                                                                                var v125 uint8 = (v3 + (v3 - v3))
                                                                                                                _ = v125
                                                                                                            }
                                                                                                        }
                                                                                                    case 128:
                                                                                                        if (0.1 > (1e10 * (3.14159 + 3.14159))) {
                                                                                                            fmt.Println(fi88, v1, fi88)
                                                                                                            delete(v2, fi88)
                                                                                                            v2[fi87] = -9223372036854775808
                                                                                                            v89 -= (v89 << (((v1 | (v1 + (v1 | (v1 << ((42) & 63)))))) & 63))
                                                                                                        } else {
                                                                                                            v95 = ((v4 && v4) && true)
                                                                                                            v3 -= v3
                                                                                                        }
                                                                                                    case 0:
                                                                                                        v126 := GenS9{F0: (^(v1)), F1: 980}
                                                                                                        _ = v126
                                                                                                    default:
                                                                                                        switch ("" + ("" + "a")) {
                                                                                                        case "hello":
                                                                                                            v127 := GenS10{F0: 239, F1: 335, F2: 9}
                                                                                                            _ = v127
                                                                                                        case "":
                                                                                                            v128 := GenS9{F0: v1, F1: 4294967295}
                                                                                                            _ = v128
                                                                                                        default:
                                                                                                            v85 += (v85 - (1e10 + 1e10))
                                                                                                        }
                                                                                                    }
                                                                                                    fmt.Println(v85)
                                                                                                    var v129 uint = v1
                                                                                                    if ((!((float32(1e10) / float32(0.1)) < (v85 + 1e10))) || ((^(v1)) > (v129 >> ((v129) & 63)))) {
                                                                                                        v130 := GenS3{F0: v1, F1: -437}
                                                                                                        _ = v130
                                                                                                    } else {
                                                                                                        if (42 != -780) {
                                                                                                            v92 -= 65535
                                                                                                        }
                                                                                                    }
                                                                                                    _ = v122
                                                                                                    _ = v129
                                                                                                }
                                                                                            case 128:
                                                                                                var v131 float32 = (-2.25 + 3.14159)
                                                                                                _ = v131
                                                                                            case -594:
                                                                                                v3 = v3
                                                                                            default:
                                                                                                var v132 int = (fi88 + 128)
                                                                                                var v133 float32 = v85
                                                                                                _ = v132
                                                                                                _ = v133
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                    v92 -= v92
                                                                                } else {
                                                                                    v1 = v1
                                                                                }
                                                                            }
                                                                        } else {
                                                                            v95 = ((v95 || (false || false)) || true)
                                                                            v89 = (v109 / ((-1) | 1))
                                                                        }
                                                                    }
                                                                }
                                                                _ = v109
                                                                _ = v110
                                                            }
                                                        default:
                                                            if v4 {
                                                                if v4 {
                                                                    fmt.Println(v95)
                                                                    fmt.Println(v1, fi87, v3)
                                                                    v2[fi88] = (^(fi88))
                                                                }
                                                                v134 := GenS6{F0: -907}
                                                                v3 += 22
                                                                _ = v134
                                                            } else {
                                                                var v135 uint64 = 18446744073709551615
                                                                switch fi87 {
                                                                case 1:
                                                                    v136 := GenS4{F0: -2147483648, F1: int(v92), F2: v92}
                                                                    var v137 string = "\n\t"
                                                                    switch "hello" {
                                                                    case "a":
                                                                        v137 = v137
                                                                    case "":
                                                                        if v4 {
                                                                            var v138 int16 = -32768
                                                                            fmt.Println(v138, fi87, v1)
                                                                            v139 := GenS2{F0: -1}
                                                                            _ = v138
                                                                            _ = v139
                                                                        } else {
                                                                            fmt.Println(fi88)
                                                                            if (52 == 255) {
                                                                                v2[534] = fi88
                                                                                v140 := GenS9{F0: 932, F1: (v135 | (v135 ^ 2))}
                                                                                v141 := []int{}
                                                                                var v142 int64 = (^(v89))
                                                                                switch (fi87 & fi88) {
                                                                                case 32767:
                                                                                    if v4 {
                                                                                        fmt.Println(v92, v3)
                                                                                        fmt.Println(v142, fi88)
                                                                                        if ((v95 || v95) || false) {
                                                                                            var v143 string = v137
                                                                                            _ = v143
                                                                                        } else {
                                                                                            fmt.Println(v137, v137)
                                                                                            v144 := GenS2{F0: int(v92)}
                                                                                            _ = v144
                                                                                        }
                                                                                        var v145 string = v137
                                                                                        _ = v145
                                                                                    }
                                                                                }
                                                                                if ((float32(0.1) * (v85 / (float32(float64(0.1)) * 0.1))) == v85) {
                                                                                    fmt.Println(fi87, v4, fi88)
                                                                                    _ = len(v141)
                                                                                    v146 := GenS11{F0: 1, F1: (false && v95), F2: v137}
                                                                                    _ = v146
                                                                                }
                                                                                v85 += 0.0
                                                                                _ = v140
                                                                                _ = v141
                                                                                _ = v142
                                                                            } else {
                                                                                var v147 uint16 = 248
                                                                                if (fi88 <= (fi87 + (fi87 & fi88))) {
                                                                                    v148 := GenS2{F0: (-(fi87))}
                                                                                    _ = v148
                                                                                }
                                                                                _ = v147
                                                                            }
                                                                            v4 = (127 != fi88)
                                                                            v92 += (^(v92))
                                                                        }
                                                                    case "\n\t":
                                                                        fmt.Println(v4)
                                                                    default:
                                                                        if ((v137 + "foo bar") != ("a" + ("" + ((v137 + (v137 + (v137 + v137))) + v137)))) {
                                                                            fmt.Println(v135, v4, fi87)
                                                                        }
                                                                    }
                                                                    _ = v136
                                                                    _ = v137
                                                                }
                                                                _ = v135
                                                            }
                                                        }
                                                    case "":
                                                        delete(v2, -448)
                                                    }
                                                    if false {
                                                        delete(v2, fi88)
                                                        if (!false) {
                                                            v149 := GenS4{F0: 42, F1: (fi87 << ((v1) & 63)), F2: v92}
                                                            _ = v149
                                                        } else {
                                                            v150 := GenS12{F0: 455, F1: (v3 << (((v1 >> ((18446744073709551615) & 63))) & 63)), F2: v92}
                                                            _ = v150
                                                        }
                                                        fmt.Println(fi87, v4, v3)
                                                    } else {
                                                        v89 += v89
                                                        var v151 uint = v1
                                                        v85 = v85
                                                        _ = v151
                                                    }
                                                } else {
                                                    v152 := GenS13{F0: 6, F1: 2147483647, F2: -32768}
                                                    _ = v152
                                                }
                                            } else {
                                                var v153 float64 = 0.1
                                                v95 = true
                                                _ = v153
                                            }
                                        default:
                                            var v154 int = fi88
                                            _ = v154
                                        }
                                    }
                                    fmt.Println(fi87)
                                    _ = v93
                                    _ = v95
                                }
                                var v155 string = ""
                                _ = v155
                                _ = v92
                            } else {
                                var v156 bool = v4
                                _ = v156
                            }
                            var v157 int8 = 7
                            v158 := GenS14{F0: 527, F1: 2147483648}
                            v159 := []int{-9223372036854775808, fi87, (fi88 | fi87), (fi88 << (((v1 + v1)) & 63))}
                            _ = fi88
                            _ = v157
                            _ = v158
                            _ = v159
                            _ = v89
                            _ = v90
                            _ = v91
                        }
                        v160 := GenS15{F0: (v3 / ((v3) | 1)), F1: 0, F2: v1}
                        _ = fi87
                        _ = v160
                    }
                    var v161 bool = true
                    fmt.Println(v85, v85)
                    _ = v161
                } else {
                    if (!false) {
                        var v162 uint = 2
                        v163 := GenS16{F0: 1}
                        v164 := []int{830, 4294967295, 329}
                        _ = v162
                        _ = v163
                        _ = v164
                    }
                }
                v2[9223372036854775807] = -595
            }
            _ = v85
        }
    case 2147483647:
        v1 += 255
    case -2147483648:
        var v165 float32 = float32(1e10)
        _ = v165
    }
    for fi166 := 0; fi166 < 4; fi166++ {
        var v167 float32 = (0.1 / 1.0)
        fmt.Println(v167, v1)
        for fi168 := 0; fi168 < 5; fi168++ {
            var v169 int8 = -8
            _ = fi168
            _ = v169
        }
        _ = fi166
        _ = v167
    }
    if v4 {
        var v170 int16 = -122
        var v171 uint8 = v3
        switch -445 {
        case -302:
            var v172 uint16 = 65535
            _ = v172
        case 127:
            v171 -= (^(v3))
        case 167:
            v3 -= genFn0(1)
        default:
            var v173 int32 = -469
            _ = v173
        }
        _ = v170
        _ = v171
    } else {
        fmt.Println(v4, v1, v1)
        var v174 int = 32767
        if v4 {
            if true {
                fmt.Println(v1, v1)
            } else {
                for fi175 := 0; fi175 < 20; fi175++ {
                    if v4 {
                        var v176 int64 = 1
                        var v177 string = ""
                        if v4 {
                            var v178 float32 = float32(float64(1.5))
                            var v179 int = v174
                            fmt.Println(v1, v3)
                            _ = v178
                            _ = v179
                        }
                        _ = v176
                        _ = v177
                    }
                    var v180 float64 = 3.14159
                    _ = fi175
                    _ = v180
                }
                if (!v4) {
                    v181 := GenS5{F0: 42}
                    v1 = (^(v1))
                    _ = v181
                }
                var v182 bool = false
                _ = v182
            }
        }
        _ = v174
    }
    if v4 {
        var v183 float32 = 1.5
        var v184 uint32 = uint32(v1)
        v185 := []int{4294967295, 128}
        _ = v183
        _ = v184
        _ = v185
    } else {
        v1 -= v1
    }
    v1 = (v1 / (((v1 >> ((786) & 63))) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
}
