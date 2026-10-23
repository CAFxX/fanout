package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int32; F1 string; F2 int32 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int32; F1 string }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 int64; F1 string }
type GenS7 struct { F0 int16; F1 uint32 }
type GenS8 struct { F0 int16; F1 string; F2 bool }
type GenS9 struct { F0 uint }
type GenS10 struct { F0 int }
type GenS11 struct { F0 uint32; F1 uint16 }
type GenS12 struct { F0 uint32; F1 int32 }
type GenS13 struct { F0 uint32; F1 uint; F2 int64 }
type GenS14 struct { F0 string; F1 int8 }
type GenS15 struct { F0 int64 }
type GenS16 struct { F0 uint32; F1 uint; F2 int32 }
type GenS17 struct { F0 uint8; F1 bool }
type GenS18 struct { F0 uint64; F1 int }
type GenS19 struct { F0 string; F1 uint16; F2 uint64 }
type GenS20 struct { F0 int16; F1 uint32 }
type GenS21 struct { F0 uint64; F1 int8; F2 uint64 }
type GenS22 struct { F0 int8; F1 uint32; F2 bool }
type GenS23 struct { F0 int; F1 int64; F2 int8 }
func genFn0(p0 uint8) uint32 {
    v7 := GenS1{F0: true}
    if (-128 <= 32) {
        for fi8 := 0; fi8 < 9; fi8++ {
            var v9 int = fi8
            switch v9 {
            case 0:
                v10 := GenS1{F0: (true || true)}
                _ = v10
            case -1:
                v9 += v9
            }
            p0 += (p0 * p0)
            _ = fi8
            _ = v9
        }
    } else {
        p0 -= 42
    }
    var v11 string = "\n\t"
    v12 := []int{-724, -644, 32767, 127}
    _ = p0
    _ = v11
    _ = v12
    _ = v7
    return 0
}

func main() {
    v1 := make(map[int]uint32)
    var v2 int32 = -2147483648
    v3 := []uint32{210, 2147483648, 0}
    v4 := make(map[uint32]bool)
    v5 := make(map[string]string)
    v6 := []int{}
    v1[-2147483648] = 538
    v13 := GenS1{F0: (("foo bar" + "\n\t") <= "a")}
    var v14 uint32 = 1
    v2 -= -101
    var v15 int = 164
    if len(v3) > 0 {
        v3[0] = genFn0(76)
    }
    for fi16 := 0; fi16 < 17; fi16++ {
        fmt.Println(v2)
        _ = fi16
    }
    if false {
        v17 := GenS1{F0: false}
        fmt.Println(v15, v14, v14)
        _ = v17
    } else {
        v14 = 42
    }
    v15 += 9223372036854775807
    fmt.Println(v2, v15)
    var v18 string = ("" + (("" + ("" + "")) + ("a" + "a")))
    v19 := GenS1{F0: false}
    if true {
        if false {
            if (2 < 255) {
                v15 -= (v15 / ((v15) | 1))
                _ = len(v6)
                if (2 >= 2) {
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                    v14 += v14
                }
            }
        } else {
            if true {
                v14 -= v14
                _ = len(v5)
                v14 += (v14 / ((v14) | 1))
            } else {
                var v20 uint64 = 4294967295
                if false {
                    var v21 uint64 = v20
                    for fi22 := 0; fi22 < 5; fi22++ {
                        fmt.Println(v21, v18)
                        if true {
                            v23 := GenS2{F0: v2, F1: "", F2: 2147483647}
                            switch v18 {
                            case "a":
                                delete(v4, (v14 | v14))
                            default:
                                fmt.Println(v20, fi22, v18)
                            }
                            _ = v23
                        }
                        for fi24 := 0; fi24 < 3; fi24++ {
                            var v25 uint32 = v14
                            fmt.Println(v21)
                            _ = fi24
                            _ = v25
                        }
                        _ = fi22
                    }
                    _ = v21
                } else {
                    var v26 float32 = (3.14159 - (-2.25 + -2.25))
                    var v27 uint16 = 1
                    _ = v26
                    _ = v27
                }
                _ = v20
            }
            var v28 uint = 4294967295
            v18 = v18
            _ = v28
        }
        v2 += (v2 - v2)
        var v29 int32 = v2
        _ = v29
    } else {
        v30 := GenS1{F0: false}
        var v31 float64 = 0.1
        _ = v30
        _ = v31
    }
    if true {
        v18 = "hello"
    } else {
        switch (("hello" + ((("a" + v18) + (v18 + v18)) + "a")) + ((("a" + "\n\t") + v18) + "\n\t")) {
        case "a":
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
        case "foo bar":
            if false {
                v32 := GenS1{F0: false}
                if (608 >= 1) {
                    v15 += (-(v15))
                    v33 := GenS2{F0: v2, F1: v18, F2: v2}
                    for fi34 := 0; fi34 < 14; fi34++ {
                        switch v15 {
                        case 32767:
                            if (!(!true)) {
                                var v35 int32 = v2
                                v2 -= (v35 & 811)
                                for fi36 := 0; fi36 < 14; fi36++ {
                                    if len(v3) > 0 {
                                        v3[0] = 65535
                                    }
                                    _ = fi36
                                }
                                _ = v35
                            }
                        case 6:
                            for fi37 := 0; fi37 < 17; fi37++ {
                                switch (v18 + v18) {
                                case "\n\t":
                                    if (false && (!false)) {
                                        var v38 int8 = 0
                                        v39 := GenS2{F0: v2, F1: "", F2: v2}
                                        v40 := GenS3{F0: (-(v38))}
                                        _ = v38
                                        _ = v39
                                        _ = v40
                                    }
                                case "a":
                                    v18 = ((v18 + v18) + v18)
                                case "foo bar":
                                    v18 = ("" + v18)
                                }
                                _ = fi37
                            }
                        case 633:
                            v14 -= 829
                        default:
                            switch v15 {
                            case 1:
                                v18 = v18
                            case 32768:
                                fmt.Println(v15, fi34)
                            }
                        }
                        if (42 != uint8(v14)) {
                            if len(v3) > 0 {
                                v3[0] = (v14 * 174)
                            }
                            v41 := GenS2{F0: v2, F1: ((("foo bar" + v18) + (v18 + v18)) + ""), F2: (-(v2))}
                            _ = v41
                        } else {
                            v14 = (v14 / ((v14) | 1))
                            if false {
                                v14 -= v14
                                fmt.Println(v14, v2)
                            }
                        }
                        _ = fi34
                    }
                    _ = v33
                } else {
                    switch (v15 + (v15 / ((102) | 1))) {
                    case 0:
                        for fi42 := 0; fi42 < 20; fi42++ {
                            if (float32(-2.25) >= 1e10) {
                                v15 += fi42
                                var v43 int8 = -50
                                _ = v43
                            }
                            v44 := GenS4{F0: (-(v2)), F1: "a"}
                            v45 := GenS1{F0: false}
                            _ = fi42
                            _ = v44
                            _ = v45
                        }
                    case 2147483647:
                        if (((0.1 + (1e10 / -2.25)) * ((float32(0.0) + (3.14159 / (float32(-2.25) - 0.0))) + (((0.0 / 1.0) + 1.5) - ((3.14159 + 1e10) - 1e10)))) < 3.14159) {
                            var v46 int8 = int8(v2)
                            _ = v46
                        } else {
                            var v47 uint = 18446744073709551615
                            if false {
                                switch v18 {
                                case "a":
                                    v48 := GenS5{F0: 418}
                                    _ = v48
                                default:
                                    for fi49 := 0; fi49 < 4; fi49++ {
                                        switch (v15 | 4294967295) {
                                        case -2147483648:
                                            v47 = (v47 - v47)
                                        case 924:
                                            switch fi49 {
                                            case 42:
                                                v14 -= 860
                                            case 128:
                                                var v50 int32 = v2
                                                _ = v50
                                            case -2147483648:
                                                if len(v6) >= 2 {
                                                    v6 = v6[1:len(v6)]
                                                }
                                            default:
                                                v51 := GenS6{F0: -1, F1: v18}
                                                _ = v51
                                            }
                                        case 127:
                                            v52 := v1[517]
                                            _ = v52
                                        default:
                                            if ((genFn0(255) == genFn0(1)) && false) {
                                                v1[fi49] = 2147483648
                                                delete(v5, (v18 + v18))
                                            }
                                        }
                                        _ = fi49
                                    }
                                }
                                for fi53 := 0; fi53 < 15; fi53++ {
                                    var v54 int8 = -128
                                    if (1e10 > (0.1 / (((0.0 - 1.5) - 0.1) + (((0.1 - (1.5 / 1.5)) - -2.25) / 3.14159)))) {
                                        if true {
                                            v2 = 42
                                            fmt.Println(v54, v18)
                                            for fi55 := 0; fi55 < 8; fi55++ {
                                                var v56 int8 = (v54 + v54)
                                                _ = fi55
                                                _ = v56
                                            }
                                        }
                                        v2 = (v2 * (v2 & (^(v2))))
                                        if (42 <= 1000) {
                                            for fi57 := 0; fi57 < 19; fi57++ {
                                                v54 = v54
                                                fmt.Println(v18)
                                                var v58 uint = (v47 | v47)
                                                var v59 bool = (false && ((!((0.0 / -2.25) < 1.5)) || (false && false)))
                                                _ = fi57
                                                _ = v58
                                                _ = v59
                                            }
                                        } else {
                                            fmt.Println(v15, v54)
                                        }
                                    } else {
                                        fmt.Println(v15)
                                    }
                                    if (1.5 < (0.1 + 0.1)) {
                                        for fi60 := 0; fi60 < 6; fi60++ {
                                            switch v18 {
                                            case "foo bar":
                                                var v61 int = v15
                                                _ = v61
                                            case "a":
                                                fmt.Println(v15)
                                            case "":
                                                fmt.Println(fi53, v54, v54)
                                            default:
                                                switch v18 {
                                                case "foo bar":
                                                    switch fi60 {
                                                    case -9223372036854775808:
                                                        var v62 int32 = -904
                                                        _ = v62
                                                    case 0:
                                                        if false {
                                                            v63 := GenS7{F0: 436, F1: 4294967295}
                                                            v1[1] = 65535
                                                            _ = v63
                                                        } else {
                                                            var v64 float32 = -2.25
                                                            switch (fi53 * v15) {
                                                            case -406:
                                                                if (-893 != -886) {
                                                                    var v65 int = fi60
                                                                    v5[v18] = ("" + v18)
                                                                    var v66 uint8 = 77
                                                                    _ = v65
                                                                    _ = v66
                                                                } else {
                                                                    v67 := GenS8{F0: 1, F1: "\n\t", F2: true}
                                                                    v54 = 0
                                                                    if (127 < 127) {
                                                                        v68 := v1[32767]
                                                                        _ = v68
                                                                    }
                                                                    _ = v67
                                                                }
                                                            case 0:
                                                                v54 -= v54
                                                            case 386:
                                                                v4[v14] = false
                                                            default:
                                                                switch v18 {
                                                                case "":
                                                                    fmt.Println(v14, v15, fi60)
                                                                case "a":
                                                                    if ((!(!false)) || false) {
                                                                        if len(v6) >= 2 {
                                                                            v6 = v6[1:len(v6)]
                                                                        }
                                                                        v69 := v5["hello"]
                                                                        _ = v69
                                                                    } else {
                                                                        if false {
                                                                            var v70 float64 = -2.25
                                                                            _ = v70
                                                                        }
                                                                        v5[v18] = ("" + v18)
                                                                        var v71 int16 = 32767
                                                                        _ = v71
                                                                    }
                                                                }
                                                            }
                                                            v18 = "hello"
                                                            _ = v64
                                                        }
                                                    }
                                                }
                                            }
                                            switch 0 {
                                            case 32768:
                                                if false {
                                                    switch v18 {
                                                    case "a":
                                                        var v72 uint = v47
                                                        if (398 <= 534) {
                                                            var v73 uint = (v47 % ((v47) | 1))
                                                            _ = v73
                                                        }
                                                        var v74 int32 = (v2 - (v2 << ((42) & 63)))
                                                        _ = v72
                                                        _ = v74
                                                    }
                                                    v75 := GenS1{F0: (true || false)}
                                                    var v76 string = v18
                                                    _ = v75
                                                    _ = v76
                                                } else {
                                                    _ = len(v4)
                                                    var v77 uint16 = 47
                                                    _ = v77
                                                }
                                            case 1:
                                                var v78 float64 = (0.0 * (1.5 + (0.1 / ((((1e10 * 1e10) - float64(1.5)) - (-2.25 - -2.25)) / (0.1 * 0.1)))))
                                                _ = v78
                                            default:
                                                if (false && (!(!(233 > 1)))) {
                                                    v47 = v47
                                                    v79 := GenS9{F0: 18446744073709551615}
                                                    v14 -= (v14 % (((^(v14))) | 1))
                                                    _ = v79
                                                }
                                            }
                                            _ = fi60
                                        }
                                        fmt.Println(fi53, v2, v14)
                                    }
                                    _ = fi53
                                    _ = v54
                                }
                                var v80 uint16 = 32767
                                _ = v80
                            }
                            var v81 float32 = float32(float64(1e10))
                            _ = v47
                            _ = v81
                        }
                    case 4294967295:
                        v14 = v14
                    }
                    switch (v15 + v15) {
                    case -2147483648:
                        v82 := v4[2147483648]
                        _ = v82
                    case 819:
                        v14 = genFn0(1)
                    case 902:
                        var v83 uint64 = 4294967295
                        _ = v83
                    }
                    if false {
                        for fi84 := 0; fi84 < 7; fi84++ {
                            v15 += v15
                            v18 = v18
                            var v85 uint16 = 759
                            _ = fi84
                            _ = v85
                        }
                    }
                }
                switch "" {
                case "a":
                    if ((v15 % ((v15) | 1)) <= v15) {
                        for fi86 := 0; fi86 < 8; fi86++ {
                            var v87 float64 = ((0.0 * 3.14159) - ((1.5 / (((0.0 * (-2.25 * -2.25)) + ((3.14159 / (1e10 - 1.5)) / 1.0)) - (3.14159 - 0.0))) * 1.5))
                            var v88 int32 = (^(v2))
                            var v89 uint32 = (v14 + (v14 & (v14 * v14)))
                            _ = fi86
                            _ = v87
                            _ = v88
                            _ = v89
                        }
                        if false {
                            v90 := GenS10{F0: v15}
                            if true {
                                if (-1 <= (v2 << ((4294967295) & 63))) {
                                    var v91 int64 = -737
                                    if true {
                                        for fi92 := 0; fi92 < 18; fi92++ {
                                            switch ((("foo bar" + v18) + v18) + v18) {
                                            case "foo bar":
                                                fmt.Println(v15)
                                            case "":
                                                if (true && false) {
                                                    for fi93 := 0; fi93 < 17; fi93++ {
                                                        if (!((!false) && (false || false))) {
                                                            var v94 int16 = -776
                                                            _ = v94
                                                        } else {
                                                            var v95 uint16 = 0
                                                            if (628 > 65535) {
                                                                switch "\n\t" {
                                                                case "\n\t":
                                                                    v96 := GenS11{F0: v14, F1: 32767}
                                                                    _ = v96
                                                                case "hello":
                                                                    v2 += 561
                                                                case "":
                                                                    if len(v6) > 0 {
                                                                        v6[0] = 0
                                                                    }
                                                                default:
                                                                    v91 += -1
                                                                }
                                                            }
                                                            var v97 uint = 834
                                                            _ = v95
                                                            _ = v97
                                                        }
                                                        var v98 int = v15
                                                        if false {
                                                            v2 = (v2 >> ((251) & 63))
                                                            v2 -= (v2 / (((v2 - v2)) | 1))
                                                            if len(v3) >= 2 {
                                                                v3 = v3[1:len(v3)]
                                                            }
                                                        }
                                                        _ = fi93
                                                        _ = v98
                                                    }
                                                    switch v15 {
                                                    case 127:
                                                        v99 := GenS12{F0: 2, F1: (v2 & (v2 & v2))}
                                                        _ = v99
                                                    }
                                                }
                                            default:
                                                _ = len(v3)
                                            }
                                            var v100 int32 = 2147483647
                                            var v101 int32 = 554
                                            _ = fi92
                                            _ = v100
                                            _ = v101
                                        }
                                        v18 = ("foo bar" + v18)
                                    }
                                    _ = v91
                                }
                                for fi102 := 0; fi102 < 13; fi102++ {
                                    var v103 uint32 = (^(v14))
                                    switch v18 {
                                    case "":
                                        v5[(v18 + v18)] = (((v18 + v18) + (v18 + v18)) + v18)
                                    default:
                                        v14 = v103
                                    }
                                    var v104 uint = 18446744073709551615
                                    _ = fi102
                                    _ = v103
                                    _ = v104
                                }
                            } else {
                                _ = len(v5)
                                switch v15 {
                                case -408:
                                    v2 = v2
                                case -2147483648:
                                    if len(v6) >= 2 {
                                        v6 = v6[1:len(v6)]
                                    }
                                }
                                v105 := GenS13{F0: 2, F1: 65535, F2: 0}
                                _ = v105
                            }
                            for fi106 := 0; fi106 < 19; fi106++ {
                                if (255 < 65535) {
                                    if false {
                                        v107 := GenS7{F0: 1, F1: (v14 % (((v14 ^ (^(v14)))) | 1))}
                                        v108 := v4[(v14 | v14)]
                                        _ = v107
                                        _ = v108
                                    }
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                } else {
                                    for fi109 := 0; fi109 < 7; fi109++ {
                                        if (276 <= 927) {
                                            if false {
                                                var v110 uint32 = 817
                                                _ = v110
                                            } else {
                                                var v111 uint = 42
                                                if false {
                                                    v2 -= 2147483647
                                                } else {
                                                    switch v15 {
                                                    case 42:
                                                        v18 = v18
                                                    case -160:
                                                        var v112 bool = false
                                                        _ = v112
                                                    }
                                                    var v113 uint16 = 65535
                                                    _ = v113
                                                }
                                                _ = v111
                                            }
                                            switch -9223372036854775808 {
                                            case 2147483647:
                                                if (false || false) {
                                                    v3 = append(v3, (^(v14)))
                                                    switch -611 {
                                                    case 127:
                                                        v114 := GenS7{F0: 42, F1: v14}
                                                        _ = v114
                                                    case 728:
                                                        v15 -= fi106
                                                    }
                                                }
                                            case 127:
                                                v115 := GenS3{F0: 101}
                                                _ = v115
                                            case 4294967295:
                                                fmt.Println(fi109)
                                            }
                                        }
                                        _ = fi109
                                    }
                                    _ = len(v3)
                                    for fi116 := 0; fi116 < 13; fi116++ {
                                        if false {
                                            var v117 uint = 42
                                            var v118 int32 = (v2 - (v2 - (v2 - 42)))
                                            _ = v117
                                            _ = v118
                                        }
                                        v15 -= (v15 % ((fi106) | 1))
                                        var v119 int64 = -9223372036854775808
                                        v119 = -9223372036854775808
                                        v120 := GenS14{F0: (v18 + v18), F1: -128}
                                        fmt.Println(fi116)
                                        if false {
                                            v14 += v14
                                            _ = len(v4)
                                            var v121 uint64 = 2
                                            if true {
                                                if (fi116 >= v15) {
                                                    if false {
                                                        v15 += v15
                                                        v122 := GenS15{F0: (-(v119))}
                                                        _ = v122
                                                    } else {
                                                        var v123 uint8 = 128
                                                        v123 += (v123 >> ((1) & 63))
                                                        _ = v123
                                                    }
                                                    switch 26 {
                                                    case 9223372036854775807:
                                                        var v124 float32 = 1.5
                                                        _ = v124
                                                    case 942:
                                                        switch 127 {
                                                        case 32767:
                                                            switch (fi106 / (((fi106 >> ((255) & 63))) | 1)) {
                                                            case 9223372036854775807:
                                                                if true {
                                                                    v15 = (-(fi116))
                                                                }
                                                            case -9223372036854775808:
                                                                var v125 uint8 = 42
                                                                v126 := GenS3{F0: 100}
                                                                _ = v125
                                                                _ = v126
                                                            }
                                                        case -742:
                                                            if (true && (0.1 < 3.14159)) {
                                                                var v127 float64 = ((1e10 - (1e10 - 1e10)) * (float64(float32(-2.25)) / (0.0 + 1.5)))
                                                                v127 = 0.1
                                                                var v128 int64 = 9223372036854775807
                                                                switch int(v2) {
                                                                case -2147483648:
                                                                    if false {
                                                                        v129 := GenS16{F0: 4294967295, F1: 18446744073709551615, F2: (-(v2))}
                                                                        _ = v129
                                                                    }
                                                                default:
                                                                    if (-32768 < -32768) {
                                                                        fmt.Println(fi116)
                                                                        if true {
                                                                            fmt.Println(v127, v128, v121)
                                                                        } else {
                                                                            _ = len(v5)
                                                                            if true {
                                                                                v130 := GenS3{F0: 1}
                                                                                v2 = v2
                                                                                v131 := v5[(v18 + v18)]
                                                                                v2 = v2
                                                                                _ = v130
                                                                                _ = v131
                                                                            }
                                                                            if true {
                                                                                fmt.Println(v121, fi116, v15)
                                                                            } else {
                                                                                if true {
                                                                                    var v132 float64 = float64(3.14159)
                                                                                    v133 := GenS17{F0: 52, F1: false}
                                                                                    fmt.Println(v18)
                                                                                    v134 := GenS1{F0: (1.5 == -2.25)}
                                                                                    var v135 int = (v15 - -633)
                                                                                    _ = v132
                                                                                    _ = v133
                                                                                    _ = v134
                                                                                    _ = v135
                                                                                } else {
                                                                                    var v136 uint32 = (v14 * (v14 - genFn0(1)))
                                                                                    switch 826 {
                                                                                    case -9223372036854775808:
                                                                                        if ((!(false || false)) && (v18 <= v18)) {
                                                                                            var v137 int32 = (v2 >> ((1) & 63))
                                                                                            v137 = v137
                                                                                            switch ((v18 + ("a" + v18)) + ((v18 + v18) + "\n\t")) {
                                                                                            case "foo bar":
                                                                                                var v138 uint16 = uint16(v121)
                                                                                                _ = v138
                                                                                            case "hello":
                                                                                                var v139 uint16 = 255
                                                                                                _ = v139
                                                                                            case "\n\t":
                                                                                                v137 -= (v2 + (^(v137)))
                                                                                            default:
                                                                                                v14 -= genFn0(1)
                                                                                            }
                                                                                            var v140 float64 = (3.14159 - v127)
                                                                                            v15 = 128
                                                                                            fmt.Println(v119, v128, v15)
                                                                                            if (v18 <= v18) {
                                                                                                fmt.Println(v127, v128, v119)
                                                                                                var v141 int8 = 109
                                                                                                v2 -= v2
                                                                                                _ = v141
                                                                                            }
                                                                                            _ = v137
                                                                                            _ = v140
                                                                                        } else {
                                                                                            var v142 uint16 = 32767
                                                                                            v143 := GenS18{F0: v121, F1: v15}
                                                                                            _ = v142
                                                                                            _ = v143
                                                                                        }
                                                                                    default:
                                                                                        if ((false && true) && (0 >= 255)) {
                                                                                            v2 += v2
                                                                                            fmt.Println(v18, v119)
                                                                                        } else {
                                                                                            v144 := GenS19{F0: v18, F1: 66, F2: (v121 + (v121 - 18446744073709551615))}
                                                                                            if ((230 == 3) || ((v136 << ((589) & 63)) >= (v136 ^ v14))) {
                                                                                                v128 -= v119
                                                                                            } else {
                                                                                                fmt.Println(v14, v128, v119)
                                                                                                fmt.Println(v119, v128)
                                                                                                var v145 int32 = v2
                                                                                                _ = v145
                                                                                            }
                                                                                            var v146 float32 = 0.0
                                                                                            _ = v144
                                                                                            _ = v146
                                                                                        }
                                                                                    }
                                                                                    _ = v136
                                                                                }
                                                                                _ = len(v4)
                                                                                v147 := GenS19{F0: v18, F1: 433, F2: 4294967295}
                                                                                _ = v147
                                                                            }
                                                                        }
                                                                        v15 = v15
                                                                    }
                                                                }
                                                                _ = v127
                                                                _ = v128
                                                            } else {
                                                                var v148 uint32 = (v14 >> ((42) & 63))
                                                                var v149 uint8 = 42
                                                                _ = v148
                                                                _ = v149
                                                            }
                                                        default:
                                                            v150 := GenS18{F0: (v121 * (^(v121))), F1: (v15 * (fi116 << ((65535) & 63)))}
                                                            _ = v150
                                                        }
                                                    }
                                                    v2 -= v2
                                                }
                                                v151 := GenS20{F0: int16(v14), F1: 42}
                                                _ = v151
                                            }
                                            _ = v121
                                        } else {
                                            var v152 int32 = (-(v2))
                                            _ = v152
                                        }
                                        _ = fi116
                                        _ = v119
                                        _ = v120
                                    }
                                }
                                _ = fi106
                            }
                            _ = v90
                        } else {
                            v15 = (^(v15))
                            delete(v4, v14)
                            v2 = (v2 | 230)
                        }
                    } else {
                        if (255 != 42) {
                            var v153 uint32 = 2147483647
                            v18 = v18
                            fmt.Println(v18, v15, v14)
                            _ = v153
                        } else {
                            v15 += 0
                            v154 := GenS15{F0: -9223372036854775808}
                            _ = v154
                        }
                        for fi155 := 0; fi155 < 18; fi155++ {
                            v18 = v18
                            v156 := GenS1{F0: false}
                            for fi157 := 0; fi157 < 8; fi157++ {
                                fmt.Println(v2, v14)
                                if (!true) {
                                    v158 := GenS21{F0: 2, F1: 42, F2: 42}
                                    _ = v158
                                } else {
                                    switch v18 {
                                    case "a":
                                        v159 := GenS15{F0: 42}
                                        _ = v159
                                    case "":
                                        fmt.Println(v2, fi157)
                                    }
                                }
                                v2 = v2
                                _ = fi157
                            }
                            _ = fi155
                            _ = v156
                        }
                    }
                case "hello":
                    for fi160 := 0; fi160 < 9; fi160++ {
                        fmt.Println(v15, v18)
                        var v161 int8 = int8(v15)
                        for fi162 := 0; fi162 < 19; fi162++ {
                            v15 -= fi160
                            v2 -= (v2 % ((v2) | 1))
                            if false {
                                var v163 uint16 = 0
                                var v164 float64 = 1.5
                                _ = v163
                                _ = v164
                            }
                            v165 := GenS7{F0: 1, F1: (v14 >> ((65535) & 63))}
                            if (true && (!true)) {
                                v18 = v18
                                v161 = (^(v161))
                            }
                            v166 := GenS22{F0: int8(v14), F1: genFn0(0), F2: false}
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                            _ = fi162
                            _ = v165
                            _ = v166
                        }
                        _ = fi160
                        _ = v161
                    }
                }
                _ = v32
            } else {
                if false {
                    fmt.Println(v2)
                    var v167 uint32 = v14
                    _ = v167
                } else {
                    var v168 uint32 = genFn0(148)
                    v169 := GenS23{F0: v15, F1: 9223372036854775807, F2: 127}
                    _ = v168
                    _ = v169
                }
            }
        case "hello":
            for fi170 := 0; fi170 < 17; fi170++ {
                v3 = append(v3, v14)
                _ = fi170
            }
        }
        v6 = append(v6, 0)
        var v171 float32 = 0.0
        _ = v171
    }
    v18 = ("a" + v18)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
