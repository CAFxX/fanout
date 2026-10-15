package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8 }
type GenS2 struct { F0 int8; F1 int8 }
type GenS3 struct { F0 uint16; F1 uint }
type GenS4 struct { F0 int32; F1 int16 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 uint; F1 uint32; F2 int64 }
type GenS7 struct { F0 uint8; F1 string; F2 int }
type GenS8 struct { F0 uint32; F1 int64 }
type GenS9 struct { F0 int8; F1 string; F2 int64 }
type GenS10 struct { F0 int; F1 uint16 }
type GenS11 struct { F0 uint16 }
type GenS12 struct { F0 int; F1 uint; F2 string }
type GenS13 struct { F0 int16; F1 int8; F2 int16 }
type GenS14 struct { F0 uint32; F1 uint16 }
type GenS15 struct { F0 int }
type GenS16 struct { F0 string }
func genFn0(p0 int16) uint32 {
    p0 += p0
    var v5 bool = ((42 != 1) && (true || (!true)))
    v6 := []int{491, 9223372036854775807}
    v5 = true
    v5 = (2147483647 <= 2)
    _ = p0
    _ = v5
    _ = v6
    return 65535
}

func main() {
    var v1 int8 = -1
    var v2 uint64 = 940
    var v3 uint64 = (v2 >> ((2) & 63))
    v4 := []string{}
    switch "a" {
    case "hello":
        _ = len(v4)
    default:
        v3 = (v3 - (v3 | v3))
    }
    var v7 int8 = (v1 % ((v1) | 1))
    v7 = v7
    for fi8 := 0; fi8 < 11; fi8++ {
        if true {
            v2 -= v2
            if false {
                if (!(32767 == 0)) {
                    if false {
                        switch "" {
                        case "\n\t":
                            fmt.Println(v2)
                        default:
                            v4 = append(v4, "")
                        }
                    } else {
                        v7 = (v7 ^ 100)
                        _ = len(v4)
                    }
                    v4 = append(v4, ((((("foo bar" + ("hello" + "foo bar")) + "") + "hello") + "a") + "foo bar"))
                } else {
                    var v9 uint32 = genFn0(32767)
                    _ = v9
                }
            }
        } else {
            if ((-307 <= 9223372036854775807) || (!(false && false))) {
                for fi10 := 0; fi10 < 13; fi10++ {
                    if (-1 > -1) {
                        v1 = v7
                        if (-2.25 >= 0.1) {
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            switch "hello" {
                            case "":
                                var v11 uint32 = genFn0(0)
                                _ = v11
                            case "\n\t":
                                v7 = (v7 - 1)
                            }
                        }
                    }
                    v1 += v1
                    _ = fi10
                }
                v12 := GenS1{F0: (v3 / (((v2 / ((v2) | 1))) | 1)), F1: (v7 - v1)}
                _ = v12
            } else {
                var v13 string = "foo bar"
                if (v13 != v13) {
                    if true {
                        for fi14 := 0; fi14 < 20; fi14++ {
                            var v15 string = v13
                            _ = fi14
                            _ = v15
                        }
                    } else {
                        if false {
                            var v16 uint32 = 2
                            if (824 > 1) {
                                if false {
                                    v17 := GenS1{F0: v2, F1: v1}
                                    if true {
                                        v7 = 0
                                        v16 -= (^(v16))
                                        for fi18 := 0; fi18 < 2; fi18++ {
                                            switch (v13 + v13) {
                                            case "foo bar":
                                                if false {
                                                    v19 := GenS1{F0: (v3 * 942), F1: (v7 ^ (v7 + v1))}
                                                    switch -2147483648 {
                                                    case 843:
                                                        if false {
                                                            if len(v4) >= 2 {
                                                                v4 = v4[1:len(v4)]
                                                            }
                                                            var v20 int32 = 387
                                                            _ = v20
                                                        }
                                                    case 624:
                                                        v4 = append(v4, (("foo bar" + (v13 + (v13 + ("hello" + (v13 + v13))))) + v13))
                                                    }
                                                    switch (fi18 & (fi8 << ((65535) & 63))) {
                                                    case 127:
                                                        switch fi18 {
                                                        case -9223372036854775808:
                                                            v21 := GenS1{F0: (v3 ^ 1), F1: v1}
                                                            _ = v21
                                                        case 32768:
                                                            fmt.Println(v2, v2)
                                                        case 128:
                                                            fmt.Println(v7, fi8)
                                                        }
                                                    default:
                                                        var v22 float64 = ((((1e10 + 0.1) / (float64(float32(float64((-2.25 + (float32(3.14159) + 1e10))))) - 3.14159)) + 3.14159) / 0.1)
                                                        _ = v22
                                                    }
                                                    _ = v19
                                                } else {
                                                    if ((-1 == -2147483648) && true) {
                                                        v23 := GenS2{F0: v1, F1: (v7 % ((v1) | 1))}
                                                        switch "a" {
                                                        case "":
                                                            fmt.Println(fi18, v2)
                                                        default:
                                                            v24 := GenS2{F0: (v1 ^ v7), F1: (v1 * (^(v1)))}
                                                            _ = v24
                                                        }
                                                        _ = v23
                                                    } else {
                                                        if (1.5 <= 0.0) {
                                                            switch (^(fi8)) {
                                                            case -1:
                                                                v16 -= v16
                                                            default:
                                                                fmt.Println(v13, v16)
                                                            }
                                                            if len(v4) > 0 {
                                                                v4[0] = "a"
                                                            }
                                                            if (-1 >= 136) {
                                                                var v25 uint32 = genFn0(-369)
                                                                v26 := GenS2{F0: 1, F1: v7}
                                                                _ = v25
                                                                _ = v26
                                                            }
                                                        } else {
                                                            var v27 float64 = float64((3.14159 / 0.1))
                                                            _ = v27
                                                        }
                                                    }
                                                    fmt.Println(v13, v2)
                                                }
                                            case "\n\t":
                                                var v28 uint64 = (v3 ^ 638)
                                                _ = v28
                                            case "a":
                                                var v29 uint32 = v16
                                                _ = v29
                                            }
                                            v30 := GenS1{F0: (v3 ^ (v3 / (((v2 - 2)) | 1))), F1: (v7 - v1)}
                                            if false {
                                                var v31 float32 = 0.0
                                                if false {
                                                    if len(v4) > 0 {
                                                        v4[0] = ((v13 + ("a" + v13)) + v13)
                                                    }
                                                    v31 = v31
                                                } else {
                                                    switch (fi18 / ((42) | 1)) {
                                                    case 973:
                                                        if true {
                                                            if (581 != 9223372036854775807) {
                                                                fmt.Println(v3, v13)
                                                                v7 = v1
                                                            }
                                                            var v32 int = fi8
                                                            switch fi8 {
                                                            case 32768:
                                                                switch (v32 >> ((uint(v16)) & 63)) {
                                                                case 128:
                                                                    switch (v32 >> ((65535) & 63)) {
                                                                    case 127:
                                                                        var v33 bool = false
                                                                        _ = v33
                                                                    case -244:
                                                                        if len(v4) > 0 {
                                                                            v4[0] = "foo bar"
                                                                        }
                                                                    case -1:
                                                                        var v34 bool = true
                                                                        _ = v34
                                                                    default:
                                                                        var v35 float32 = v31
                                                                        _ = v35
                                                                    }
                                                                case 127:
                                                                    fmt.Println(v1)
                                                                case -411:
                                                                    if (false || false) {
                                                                        if true {
                                                                            switch (^(fi18)) {
                                                                            case 484:
                                                                                switch (-(fi8)) {
                                                                                case 42:
                                                                                    fmt.Println(v32, v7)
                                                                                }
                                                                                fmt.Println(v3, v3, v32)
                                                                            }
                                                                            if false {
                                                                                if true {
                                                                                    fmt.Println(v16)
                                                                                    switch "hello" {
                                                                                    case "":
                                                                                        var v36 uint = 65535
                                                                                        _ = v36
                                                                                    case "hello":
                                                                                        v1 += v1
                                                                                    case "\n\t":
                                                                                        switch fi18 {
                                                                                        case 35:
                                                                                            v7 = v1
                                                                                        default:
                                                                                            v1 = v7
                                                                                            v37 := GenS3{F0: 2, F1: 4294967295}
                                                                                            _ = v37
                                                                                        }
                                                                                    }
                                                                                }
                                                                                var v38 float32 = (0.1 * 0.0)
                                                                                _ = len(v4)
                                                                                _ = v38
                                                                            } else {
                                                                                v16 += uint32(v2)
                                                                            }
                                                                            if ((0 >= 27) && (v31 == v31)) {
                                                                                if true {
                                                                                    var v39 uint16 = 42
                                                                                    var v40 int16 = -39
                                                                                    _ = v39
                                                                                    _ = v40
                                                                                }
                                                                            }
                                                                        } else {
                                                                            v41 := GenS4{F0: -66, F1: 0}
                                                                            _ = v41
                                                                        }
                                                                    }
                                                                default:
                                                                    fmt.Println(v16)
                                                                }
                                                                fmt.Println(fi18)
                                                                fmt.Println(v16)
                                                            case 4294967295:
                                                                if (!(998 == (v3 / ((1) | 1)))) {
                                                                    var v42 uint64 = v2
                                                                    if ((18446744073709551615 < (v3 >> ((543) & 63))) && ((fi18 / ((2147483647) | 1)) != (v32 / ((fi18) | 1)))) {
                                                                        switch v32 {
                                                                        case 325:
                                                                            var v43 int = (fi8 * fi8)
                                                                            _ = v43
                                                                        case -982:
                                                                            if len(v4) >= 2 {
                                                                                v4 = v4[1:len(v4)]
                                                                            }
                                                                        case 32767:
                                                                            fmt.Println(v42)
                                                                        default:
                                                                            _ = len(v4)
                                                                        }
                                                                        _ = len(v4)
                                                                        if (false || (v31 != 3.14159)) {
                                                                            var v44 int8 = v1
                                                                            _ = v44
                                                                        } else {
                                                                            v45 := GenS2{F0: (v1 | (v1 * v7)), F1: 42}
                                                                            if false {
                                                                                v46 := GenS5{F0: -587}
                                                                                var v47 uint8 = 2
                                                                                if len(v4) > 0 {
                                                                                    v4[0] = (v13 + v13)
                                                                                }
                                                                                _ = v46
                                                                                _ = v47
                                                                            } else {
                                                                                switch v13 {
                                                                                case "hello":
                                                                                    fmt.Println(v31, v2, v2)
                                                                                case "a":
                                                                                    fmt.Println(v3, v2)
                                                                                }
                                                                                var v48 uint64 = v3
                                                                                _ = len(v4)
                                                                                _ = v48
                                                                            }
                                                                            _ = v45
                                                                        }
                                                                    }
                                                                    v42 = 42
                                                                    v49 := GenS2{F0: (v7 / ((127) | 1)), F1: (^(v7))}
                                                                    if (true || (65535 <= 883)) {
                                                                        switch "\n\t" {
                                                                        case "foo bar":
                                                                            _ = len(v4)
                                                                        case "hello":
                                                                            var v50 uint8 = 112
                                                                            _ = v50
                                                                        default:
                                                                            v31 -= v31
                                                                        }
                                                                        if len(v4) > 0 {
                                                                            v4[0] = v13
                                                                        }
                                                                    } else {
                                                                        fmt.Println(v3)
                                                                    }
                                                                    v4 = append(v4, ("" + (((v13 + v13) + v13) + (v13 + (("\n\t" + (v13 + "\n\t")) + "hello")))))
                                                                    v1 -= v1
                                                                    v32 += 360
                                                                    _ = v42
                                                                    _ = v49
                                                                }
                                                                v51 := GenS1{F0: v2, F1: 0}
                                                                _ = v51
                                                            default:
                                                                if true {
                                                                    v7 = v7
                                                                    fmt.Println(v32, fi18)
                                                                }
                                                            }
                                                            _ = v32
                                                        } else {
                                                            var v52 int32 = -1
                                                            v53 := GenS4{F0: (^(v52)), F1: 1000}
                                                            v54 := GenS4{F0: v52, F1: 0}
                                                            _ = v52
                                                            _ = v53
                                                            _ = v54
                                                        }
                                                        v55 := GenS3{F0: 65535, F1: 4294967295}
                                                        _ = v55
                                                    default:
                                                        var v56 uint32 = (v16 << ((65535) & 63))
                                                        _ = v56
                                                    }
                                                    switch fi18 {
                                                    case -639:
                                                        if ((true || false) && true) {
                                                            fmt.Println(fi18)
                                                        }
                                                    case 9223372036854775807:
                                                        v16 = 1
                                                    case 32768:
                                                        fmt.Println(fi18, v2)
                                                    }
                                                    if (true || false) {
                                                        switch v13 {
                                                        case "a":
                                                            v57 := GenS6{F0: 0, F1: v16, F2: 980}
                                                            _ = v57
                                                        }
                                                        v16 = v16
                                                        var v58 uint8 = 177
                                                        _ = v58
                                                    }
                                                    if (!true) {
                                                        switch ((v13 + (("foo bar" + ("hello" + v13)) + (v13 + v13))) + v13) {
                                                        case "foo bar":
                                                            if (!false) {
                                                                if false {
                                                                    v4 = append(v4, ((v13 + v13) + "\n\t"))
                                                                } else {
                                                                    v59 := GenS7{F0: 255, F1: v13, F2: -588}
                                                                    if true {
                                                                        v16 = genFn0(-550)
                                                                        _ = len(v4)
                                                                        v60 := GenS4{F0: 1, F1: -779}
                                                                        _ = v60
                                                                    } else {
                                                                        if len(v4) >= 2 {
                                                                            v4 = v4[1:len(v4)]
                                                                        }
                                                                    }
                                                                    _ = v59
                                                                }
                                                                var v61 int32 = -2147483648
                                                                _ = v61
                                                            }
                                                            var v62 float64 = 1e10
                                                            _ = v62
                                                        case "":
                                                            fmt.Println(v13)
                                                            var v63 uint = 1
                                                            switch v13 {
                                                            case "\n\t":
                                                                v64 := GenS8{F0: 2, F1: 42}
                                                                var v65 uint8 = 128
                                                                if false {
                                                                    var v66 int16 = int16(v63)
                                                                    _ = v66
                                                                }
                                                                _ = v64
                                                                _ = v65
                                                            case "":
                                                                fmt.Println(v16, v2, fi18)
                                                            case "a":
                                                                if true {
                                                                    v67 := GenS3{F0: 1, F1: 1}
                                                                    v1 = v7
                                                                    v68 := GenS5{F0: 9223372036854775807}
                                                                    _ = v67
                                                                    _ = v68
                                                                }
                                                            default:
                                                                v69 := GenS4{F0: 777, F1: -930}
                                                                _ = v69
                                                            }
                                                            v2 += 1
                                                            v70 := GenS9{F0: v7, F1: "a", F2: -1}
                                                            _ = v63
                                                            _ = v70
                                                        case "hello":
                                                            var v71 uint16 = 298
                                                            _ = v71
                                                        }
                                                    } else {
                                                        v72 := GenS5{F0: 884}
                                                        var v73 int64 = 838
                                                        _ = v72
                                                        _ = v73
                                                    }
                                                }
                                                var v74 float32 = (v31 / v31)
                                                _ = v31
                                                _ = v74
                                            } else {
                                                var v75 string = "foo bar"
                                                _ = v75
                                            }
                                            _ = fi18
                                            _ = v30
                                        }
                                    }
                                    _ = v17
                                } else {
                                    for fi76 := 0; fi76 < 7; fi76++ {
                                        switch v13 {
                                        case "hello":
                                            v77 := GenS10{F0: fi76, F1: 47}
                                            _ = v77
                                        default:
                                            if false {
                                                var v78 float32 = 1.5
                                                v79 := GenS10{F0: fi8, F1: 350}
                                                _ = v78
                                                _ = v79
                                            }
                                        }
                                        v80 := GenS11{F0: 2}
                                        _ = fi76
                                        _ = v80
                                    }
                                    for fi81 := 0; fi81 < 1; fi81++ {
                                        var v82 string = ""
                                        v2 = 1
                                        _ = fi81
                                        _ = v82
                                    }
                                }
                            } else {
                                v83 := GenS8{F0: v16, F1: -980}
                                for fi84 := 0; fi84 < 12; fi84++ {
                                    fmt.Println(v16)
                                    if (32767 < 1) {
                                        switch fi84 {
                                        case 59:
                                            if len(v4) > 0 {
                                                v4[0] = (("\n\t" + (v13 + v13)) + (v13 + "hello"))
                                            }
                                            switch (fi84 ^ fi84) {
                                            case 2147483647:
                                                switch fi8 {
                                                case 32768:
                                                    var v85 float32 = 1e10
                                                    _ = v85
                                                case -862:
                                                    v13 = v13
                                                    fmt.Println(v16, v16)
                                                    v7 = 12
                                                case 42:
                                                    v86 := GenS6{F0: 18446744073709551615, F1: (^(v16)), F2: 42}
                                                    _ = v86
                                                default:
                                                    v1 += v1
                                                }
                                            case 9223372036854775807:
                                                var v87 uint32 = v16
                                                _ = v87
                                            case 128:
                                                switch v13 {
                                                case "a":
                                                    v13 = v13
                                                }
                                            }
                                            if (1000 <= 975) {
                                                var v88 int32 = 2147483647
                                                var v89 float64 = (1.5 - (0.1 - 0.0))
                                                _ = v88
                                                _ = v89
                                            }
                                        case -2147483648:
                                            v2 += 878
                                        case 4294967295:
                                            if len(v4) > 0 {
                                                v4[0] = ("" + v13)
                                            }
                                        default:
                                            v90 := GenS7{F0: 0, F1: (v13 + v13), F2: (fi84 % (((fi84 * 127)) | 1))}
                                            v1 = v7
                                            _ = v90
                                        }
                                        v3 += (v3 >> ((18446744073709551615) & 63))
                                    } else {
                                        var v91 uint64 = (v2 + 42)
                                        switch (fi8 * int(v1)) {
                                        case 32768:
                                            if true {
                                                v13 = "foo bar"
                                                var v92 uint8 = 2
                                                var v93 float64 = 1.5
                                                _ = v92
                                                _ = v93
                                            } else {
                                                var v94 uint16 = 1
                                                v95 := GenS12{F0: fi84, F1: 454, F2: ""}
                                                fmt.Println(v94, v13)
                                                if (255 == 248) {
                                                    var v96 int32 = 2147483647
                                                    fmt.Println(fi84, v16)
                                                    _ = len(v4)
                                                    _ = v96
                                                }
                                                var v97 uint8 = 127
                                                v98 := GenS12{F0: (fi8 + (fi8 | (fi8 << ((889) & 63)))), F1: 255, F2: v13}
                                                _ = v94
                                                _ = v95
                                                _ = v97
                                                _ = v98
                                            }
                                        case 930:
                                            if (true || (!false)) {
                                                fmt.Println(v16, v16, v7)
                                                var v99 bool = false
                                                var v100 uint64 = 18446744073709551615
                                                _ = v100
                                                _ = v99
                                            }
                                        }
                                        _ = v91
                                    }
                                    v4 = append(v4, v13)
                                    var v101 float32 = -2.25
                                    _ = fi84
                                    _ = v101
                                }
                                if (!false) {
                                    v102 := GenS2{F0: -128, F1: (v1 * (v1 << ((662) & 63)))}
                                    for fi103 := 0; fi103 < 17; fi103++ {
                                        fmt.Println(v13)
                                        _ = len(v4)
                                        _ = fi103
                                    }
                                    for fi104 := 0; fi104 < 2; fi104++ {
                                        v105 := GenS13{F0: 1000, F1: -116, F2: 0}
                                        var v106 float32 = 0.1
                                        _ = fi104
                                        _ = v105
                                        _ = v106
                                    }
                                    _ = v102
                                }
                                _ = v83
                            }
                            var v107 bool = (1.5 <= 0.1)
                            _ = v107
                            _ = v16
                        }
                    }
                }
                var v108 uint8 = 100
                _ = v108
                _ = v13
            }
        }
        fmt.Println(v2, v3)
        _ = fi8
    }
    var v109 int = -972
    for fi110 := 0; fi110 < 19; fi110++ {
        v4 = append(v4, "a")
        _ = fi110
    }
    if (32768 <= v109) {
        if true {
            v111 := GenS1{F0: v3, F1: int8(v2)}
            var v112 uint8 = 74
            _ = v111
            _ = v112
        }
        fmt.Println(v3, v3)
    } else {
        var v113 int32 = int32(v2)
        for fi114 := 0; fi114 < 20; fi114++ {
            var v115 int64 = 826
            v116 := GenS14{F0: 2147483647, F1: 2}
            for fi117 := 0; fi117 < 16; fi117++ {
                if (-1 > -411) {
                    var v118 int = 32767
                    v1 -= v7
                    if (702 < uint(v3)) {
                        if len(v4) > 0 {
                            v4[0] = "a"
                        }
                        v119 := GenS13{F0: -1, F1: v1, F2: 42}
                        var v120 uint64 = (v2 / ((v3) | 1))
                        _ = v119
                        _ = v120
                    } else {
                        v4 = append(v4, "\n\t")
                        v121 := GenS15{F0: v109}
                        _ = v121
                    }
                    _ = v118
                }
                var v122 uint32 = 2147483647
                v3 -= (^(v3))
                _ = fi117
                _ = v122
            }
            _ = fi114
            _ = v115
            _ = v116
        }
        v4 = append(v4, ("a" + "\n\t"))
        _ = v113
    }
    v109 -= (v109 % (((v109 / ((v109) | 1))) | 1))
    v123 := GenS16{F0: ""}
    var v124 float64 = 3.14159
    v109 += 42
    v2 -= (v3 & (^(v3)))
    _ = len(v4)
    switch (v109 + v109) {
    case -9223372036854775808:
        v125 := GenS6{F0: 4294967295, F1: 777, F2: 1}
        _ = v125
    default:
        var v126 int64 = 1
        _ = v126
    }
    fmt.Println("v1:", v1)
    fmt.Println("v109:", v109)
    fmt.Println("v124:", v124)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v109
    _ = v123
    _ = v124
    _ = v2
    _ = v3
    _ = v4
    _ = v7
}
