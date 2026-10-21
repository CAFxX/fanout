package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int; F1 int8 }
type GenS3 struct { F0 string; F1 int64 }
type GenS4 struct { F0 uint32; F1 uint8; F2 string }
type GenS5 struct { F0 int }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 uint; F1 uint64 }
type GenS8 struct { F0 uint64; F1 bool; F2 uint32 }
type GenS9 struct { F0 int8; F1 uint16; F2 int }
type GenS10 struct { F0 string; F1 int8 }
type GenS11 struct { F0 uint32; F1 uint8 }
type GenS12 struct { F0 uint; F1 string }
type GenS13 struct { F0 int32; F1 int; F2 int32 }
type GenS14 struct { F0 int; F1 int; F2 uint32 }
type GenS15 struct { F0 uint16; F1 int64; F2 uint64 }
type GenS16 struct { F0 int8; F1 int }
type GenS17 struct { F0 uint64; F1 uint32; F2 uint32 }

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[int]bool)
    v3 := make(map[uint32]uint32)
    var v4 int32 = -494
    delete(v2, 992)
    v5 := GenS1{F0: "hello"}
    v6 := []int{-159, -293, 9223372036854775807}
    v7 := GenS2{F0: 9223372036854775807, F1: 0}
    for fi8 := 0; fi8 < 1; fi8++ {
        fmt.Println(v4)
        v4 -= v4
        var v9 string = "hello"
        _ = fi8
        _ = v9
    }
    for fi10 := 0; fi10 < 5; fi10++ {
        v11 := GenS3{F0: ("\n\t" + ("\n\t" + ((("\n\t" + ("\n\t" + "a")) + "") + "foo bar"))), F1: 1}
        for fi12 := 0; fi12 < 9; fi12++ {
            v13 := v3[1]
            v4 += 2147483647
            if len(v6) > 0 {
                v6[0] = fi12
            }
            _ = fi12
            _ = v13
        }
        if true {
            v14 := GenS4{F0: 42, F1: 0, F2: "\n\t"}
            _ = v14
        } else {
            v4 = 1
            v15 := GenS5{F0: (^(fi10))}
            _ = v15
        }
        _ = fi10
        _ = v11
    }
    var v16 uint16 = 446
    v16 += 494
    fmt.Println(v16, v4)
    _ = len(v2)
    if false {
        if (1 < 866) {
            v17 := GenS6{F0: v16}
            _ = v17
        } else {
            var v18 uint16 = v16
            _ = v18
        }
        if len(v6) > 0 {
            v6[0] = -318
        }
        for fi19 := 0; fi19 < 1; fi19++ {
            if len(v6) > 0 {
                v6[0] = fi19
            }
            _ = fi19
        }
    } else {
        v4 -= -146
        _ = len(v6)
        v4 = (v4 - v4)
    }
    v20 := v1[42]
    if (-9223372036854775808 < 0) {
        for fi21 := 0; fi21 < 13; fi21++ {
            for fi22 := 0; fi22 < 20; fi22++ {
                v4 = (^(v4))
                v4 = (-(v4))
                _ = fi22
            }
            v23 := v1[4294967295]
            var v24 int64 = 401
            _ = fi21
            _ = v23
            _ = v24
        }
    } else {
        switch 42 {
        case 904:
            var v25 string = "a"
            _ = v25
        case 32767:
            _ = len(v6)
        }
        if true {
            fmt.Println(v4)
        }
        if (-1 > 57) {
            if (!v20) {
                switch 1 {
                case -2147483648:
                    for fi26 := 0; fi26 < 16; fi26++ {
                        switch ("hello" + "a") {
                        case "a":
                            _ = len(v3)
                        case "foo bar":
                            _ = len(v2)
                        case "\n\t":
                            var v27 float32 = (-2.25 - 3.14159)
                            _ = v27
                        default:
                            v16 = v16
                        }
                        _ = fi26
                    }
                default:
                    v16 -= (^(v16))
                }
                var v28 int64 = -1
                if false {
                    if v20 {
                        fmt.Println(v4)
                        _ = len(v2)
                    } else {
                        for fi29 := 0; fi29 < 14; fi29++ {
                            var v30 uint8 = 144
                            var v31 uint = 771
                            _ = fi29
                            _ = v30
                            _ = v31
                        }
                    }
                } else {
                    if false {
                        v32 := GenS7{F0: 717, F1: 18446744073709551615}
                        switch "a" {
                        case "":
                            fmt.Println(v20)
                        case "\n\t":
                            if v20 {
                                v33 := v1[127]
                                if len(v6) > 0 {
                                    v6[0] = -9223372036854775808
                                }
                                _ = v33
                            } else {
                                if len(v6) >= 2 {
                                    v6 = v6[1:len(v6)]
                                }
                                for fi34 := 0; fi34 < 12; fi34++ {
                                    var v35 float64 = float64(-2.25)
                                    switch "hello" {
                                    case "foo bar":
                                        var v36 int32 = v4
                                        _ = v36
                                    case "\n\t":
                                        fmt.Println(v28, v16)
                                    }
                                    _ = fi34
                                    _ = v35
                                }
                            }
                        default:
                            var v37 int32 = (v4 ^ (v4 | v4))
                            _ = v37
                        }
                        _ = v32
                    } else {
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                        v38 := GenS7{F0: 18446744073709551615, F1: 4294967295}
                        _ = v38
                    }
                    v16 = (v16 & (^(v16)))
                }
                _ = v28
            } else {
                for fi39 := 0; fi39 < 14; fi39++ {
                    for fi40 := 0; fi40 < 5; fi40++ {
                        v4 = v4
                        _ = fi40
                    }
                    v20 = v20
                    _ = fi39
                }
                var v41 float64 = -2.25
                for fi42 := 0; fi42 < 13; fi42++ {
                    switch 9223372036854775807 {
                    case 4294967295:
                        fmt.Println(v20, v20)
                    case 127:
                        switch 128 {
                        case 662:
                            fmt.Println(fi42, fi42)
                        case 794:
                            fmt.Println(v20)
                        case -555:
                            _ = len(v1)
                        }
                    case 712:
                        v4 -= v4
                    default:
                        fmt.Println(fi42)
                    }
                    v43 := GenS8{F0: 1, F1: v20, F2: 4294967295}
                    v41 = 0.0
                    _ = fi42
                    _ = v43
                }
                _ = v41
            }
        }
    }
    v44 := GenS9{F0: 0, F1: v16, F2: 2147483647}
    var v45 uint32 = 1
    for fi46 := 0; fi46 < 4; fi46++ {
        for fi47 := 0; fi47 < 6; fi47++ {
            delete(v1, (fi46 + (fi47 / (((fi47 ^ fi47)) | 1))))
            var v48 int16 = -819
            _ = fi47
            _ = v48
        }
        _ = fi46
    }
    if (false || (!v20)) {
        for fi49 := 0; fi49 < 5; fi49++ {
            var v50 uint16 = (v16 >> ((2) & 63))
            for fi51 := 0; fi51 < 19; fi51++ {
                if v20 {
                    var v52 int64 = -1
                    var v53 float32 = ((1.5 + 3.14159) * 1.5)
                    _ = v52
                    _ = v53
                }
                if (((v20 && ((0 != 127) && v20)) && v20) || (!false)) {
                    var v54 int64 = 0
                    fmt.Println(fi51, v50, v54)
                    if ((v20 || true) || false) {
                        var v55 float64 = (-2.25 + 0.1)
                        v4 = int32(v54)
                        var v56 string = "a"
                        fmt.Println(fi49)
                        _ = v55
                        _ = v56
                    } else {
                        if len(v6) > 0 {
                            v6[0] = (fi51 ^ 748)
                        }
                        fmt.Println(v16, v20)
                        if v20 {
                            switch 0 {
                            case 2147483647:
                                v57 := GenS10{F0: "hello", F1: 127}
                                _ = v57
                            case 127:
                                switch "hello" {
                                case "foo bar":
                                    v4 -= v4
                                case "a":
                                    fmt.Println(v50)
                                default:
                                    v16 += v50
                                }
                            case 32768:
                                if v20 {
                                    v58 := GenS2{F0: (-(fi51)), F1: 0}
                                    if true {
                                        if (2 <= 318) {
                                            v50 -= (^(v16))
                                            v59 := v3[(v45 - (v45 * (v45 << ((1) & 63))))]
                                            _ = v59
                                        }
                                        v6 = append(v6, fi51)
                                    } else {
                                        v45 -= 42
                                        var v60 int = int(v45)
                                        v61 := GenS6{F0: (v16 | (v16 << ((18446744073709551615) & 63)))}
                                        _ = v60
                                        _ = v61
                                    }
                                    fmt.Println(v16)
                                    _ = v58
                                } else {
                                    if ((!true) && true) {
                                        v62 := GenS6{F0: v50}
                                        if (!(3.14159 != ((1e10 / -2.25) - (((-2.25 / 1e10) * 1.5) + (3.14159 + 1.5))))) {
                                            v20 = v20
                                        } else {
                                            v45 = v45
                                            if ((v20 && v20) && (1.5 != 3.14159)) {
                                                var v63 uint16 = 0
                                                _ = v63
                                            }
                                            var v64 uint16 = v16
                                            _ = v64
                                        }
                                        var v65 string = "hello"
                                        _ = v62
                                        _ = v65
                                    }
                                    v54 += (v54 & (^(v54)))
                                    v20 = false
                                    var v66 uint16 = v16
                                    _ = v66
                                }
                            }
                            var v67 int32 = v4
                            v68 := GenS7{F0: 881, F1: 0}
                            if len(v6) > 0 {
                                v6[0] = fi49
                            }
                            switch "foo bar" {
                            case "foo bar":
                                _ = len(v2)
                            case "hello":
                                v6 = append(v6, (fi51 % ((fi49) | 1)))
                            default:
                                v67 = v4
                            }
                            _ = v67
                            _ = v68
                        } else {
                            var v69 int64 = v54
                            _ = len(v6)
                            _ = v69
                        }
                        switch fi49 {
                        case 59:
                            if (true && (!v20)) {
                                v70 := GenS11{F0: (v45 ^ v45), F1: uint8(v45)}
                                _ = v70
                            } else {
                                if v20 {
                                    v16 = v16
                                } else {
                                    var v71 bool = (0 < (v16 ^ (v50 / (((v50 & v16)) | 1))))
                                    switch fi51 {
                                    case 32767:
                                        switch "" {
                                        case "hello":
                                            v16 -= (v16 * v16)
                                        default:
                                            switch 4294967295 {
                                            case 32768:
                                                fmt.Println(v50, fi51)
                                            case -1:
                                                fmt.Println(fi51, v20, fi51)
                                            }
                                        }
                                    case 1:
                                        if ((1 > 0) && ((v54 >> ((360) & 63)) != v54)) {
                                            fmt.Println(v54, fi51)
                                        } else {
                                            if false {
                                                fmt.Println(v4, v16, fi51)
                                            } else {
                                                v45 = (v45 | (v45 / (((v45 / (((v45 % ((v45) | 1))) | 1))) | 1)))
                                            }
                                            var v72 int = fi49
                                            if v71 {
                                                if true {
                                                    if ((((("" + "foo bar") + "foo bar") + ("\n\t" + "a")) + "hello") < "") {
                                                        switch (fi51 % ((1) | 1)) {
                                                        case 1:
                                                            if v20 {
                                                                var v73 float32 = 0.1
                                                                _ = v73
                                                            }
                                                            v6 = append(v6, fi51)
                                                        default:
                                                            if (!v71) {
                                                                var v74 uint32 = (v45 % ((v45) | 1))
                                                                v75 := GenS6{F0: v50}
                                                                fmt.Println(v74, v71, v50)
                                                                var v76 int8 = 0
                                                                _ = v74
                                                                _ = v75
                                                                _ = v76
                                                            }
                                                        }
                                                    }
                                                } else {
                                                    v4 = v4
                                                    v50 = (v16 * (v16 ^ 255))
                                                }
                                                fmt.Println(v50, v54)
                                                if ((1 != 2) || (1e10 != -2.25)) {
                                                    fmt.Println(v4)
                                                    fmt.Println(v16)
                                                }
                                            }
                                            if v20 {
                                                var v77 int = (-(fi49))
                                                v71 = v71
                                                v78 := GenS2{F0: (v77 * 4294967295), F1: 42}
                                                _ = v77
                                                _ = v78
                                            } else {
                                                v54 = (v54 << ((0) & 63))
                                            }
                                            switch "hello" {
                                            case "\n\t":
                                                switch ("" + "a") {
                                                case "hello":
                                                    v79 := v3[(v45 / ((2) | 1))]
                                                    _ = v79
                                                case "a":
                                                    var v80 int8 = 42
                                                    _ = v80
                                                case "":
                                                    if v20 {
                                                        v81 := GenS1{F0: ((("foo bar" + ((("hello" + ("hello" + "hello")) + (("\n\t" + "hello") + "hello")) + (("\n\t" + "hello") + ""))) + ("a" + "\n\t")) + (("a" + "hello") + "hello"))}
                                                        delete(v1, v72)
                                                        _ = v81
                                                    } else {
                                                        if false {
                                                            if len(v6) > 0 {
                                                                v6[0] = 128
                                                            }
                                                            v20 = (true && v71)
                                                            v82 := GenS9{F0: -128, F1: (v16 / ((v50) | 1)), F2: -1}
                                                            _ = v82
                                                        }
                                                        switch ("" + (("foo bar" + ("" + "a")) + ("\n\t" + "foo bar"))) {
                                                        case "hello":
                                                            var v83 int64 = int64(v16)
                                                            _ = v83
                                                        case "\n\t":
                                                            if false {
                                                                var v84 int64 = (v54 + -801)
                                                                var v85 string = ("" + ("a" + (("foo bar" + (("\n\t" + "") + "hello")) + "hello")))
                                                                if true {
                                                                    v50 += v16
                                                                    var v86 int8 = 127
                                                                    v16 = v50
                                                                    _ = v86
                                                                } else {
                                                                    fmt.Println(v54, v54, v50)
                                                                    var v87 string = (v85 + "\n\t")
                                                                    _ = v87
                                                                }
                                                                v84 -= (v84 - 1)
                                                                _ = v84
                                                                _ = v85
                                                            }
                                                        }
                                                    }
                                                    fmt.Println(fi49, v71, v16)
                                                    v88 := GenS12{F0: 4294967295, F1: "foo bar"}
                                                    if true {
                                                        if len(v6) >= 2 {
                                                            v6 = v6[1:len(v6)]
                                                        }
                                                        switch 127 {
                                                        case -396:
                                                            v1[0] = v71
                                                        case -689:
                                                            switch "" {
                                                            case "\n\t":
                                                                v1[32767] = true
                                                            default:
                                                                fmt.Println(v72, v54)
                                                            }
                                                        case 32767:
                                                            if v71 {
                                                                switch "hello" {
                                                                case "a":
                                                                    v50 = 0
                                                                case "":
                                                                    var v89 uint64 = uint64(v45)
                                                                    _ = v89
                                                                default:
                                                                    switch "\n\t" {
                                                                    case "foo bar":
                                                                        v71 = ((!((v54 >> ((4294967295) & 63)) != v54)) || false)
                                                                    case "hello":
                                                                        switch v72 {
                                                                        case 32768:
                                                                            if true {
                                                                                fmt.Println(v45, v45, v71)
                                                                                var v90 int32 = (-(v4))
                                                                                fmt.Println(v90, fi51)
                                                                                _ = v90
                                                                            }
                                                                        case 127:
                                                                            switch (fi49 % (((-(fi49))) | 1)) {
                                                                            case 128:
                                                                                v54 += (^(v54))
                                                                            default:
                                                                                v6 = append(v6, (fi51 << ((896) & 63)))
                                                                            }
                                                                        }
                                                                    case "":
                                                                        if (208 != 32767) {
                                                                            v2[(^(fi51))] = v71
                                                                            if v20 {
                                                                                if (v71 && false) {
                                                                                    var v91 int16 = -1
                                                                                    switch "foo bar" {
                                                                                    case "\n\t":
                                                                                        if v71 {
                                                                                            var v92 string = ("hello" + "\n\t")
                                                                                            var v93 uint64 = 4294967295
                                                                                            _ = v92
                                                                                            _ = v93
                                                                                        } else {
                                                                                            var v94 float64 = 0.1
                                                                                            _ = len(v6)
                                                                                            _ = v94
                                                                                        }
                                                                                        v95 := GenS13{F0: (v4 / (((v4 << ((253) & 63))) | 1)), F1: fi51, F2: v4}
                                                                                        _ = v95
                                                                                    case "hello":
                                                                                        if (127 > -84) {
                                                                                            var v96 int32 = (v4 % ((v4) | 1))
                                                                                            var v97 int = (^(fi49))
                                                                                            v4 = 751
                                                                                            _ = v96
                                                                                            _ = v97
                                                                                        }
                                                                                    }
                                                                                    v71 = (v91 != int16(v16))
                                                                                    _ = v91
                                                                                }
                                                                            } else {
                                                                                fmt.Println(v20, v4)
                                                                            }
                                                                        }
                                                                    }
                                                                }
                                                                v98 := GenS3{F0: "foo bar", F1: (^(v54))}
                                                                _ = v98
                                                            }
                                                        }
                                                        if ((fi51 * (-(fi49))) <= v72) {
                                                            v6 = append(v6, (fi51 | (v72 + fi49)))
                                                        }
                                                    }
                                                    switch ("a" + "foo bar") {
                                                    case "\n\t":
                                                        switch "foo bar" {
                                                        case "hello":
                                                            var v99 int32 = v4
                                                            _ = v99
                                                        case "\n\t":
                                                            v16 += 1
                                                        case "":
                                                            switch "\n\t" {
                                                            case "\n\t":
                                                                if ((v45 % ((v45) | 1)) == (v45 - v45)) {
                                                                    var v100 uint8 = 128
                                                                    _ = v100
                                                                } else {
                                                                    if len(v6) > 0 {
                                                                        v6[0] = 4294967295
                                                                    }
                                                                }
                                                            default:
                                                                delete(v3, v45)
                                                            }
                                                        }
                                                    }
                                                    _ = v88
                                                default:
                                                    var v101 int64 = -9223372036854775808
                                                    _ = v101
                                                }
                                            case "a":
                                                v102 := GenS6{F0: v50}
                                                _ = v102
                                            case "foo bar":
                                                v72 = (fi51 / ((v72) | 1))
                                            }
                                            var v103 float32 = 1.5
                                            v16 -= 32767
                                            _ = v103
                                            _ = v72
                                        }
                                    }
                                    v104 := GenS4{F0: (v45 + v45), F1: 128, F2: "hello"}
                                    _ = v104
                                    _ = v71
                                }
                                v20 = (100 != -128)
                                _ = len(v1)
                            }
                        default:
                            if v20 {
                                delete(v1, 32767)
                            }
                            var v105 int32 = v4
                            _ = v105
                        }
                    }
                    _ = v54
                }
                _ = fi51
            }
            if (312 > (v4 + v4)) {
                switch int(v4) {
                case 2147483647:
                    for fi106 := 0; fi106 < 3; fi106++ {
                        if true {
                            if v20 {
                                switch "" {
                                case "foo bar":
                                    var v107 float64 = 3.14159
                                    _ = v107
                                case "hello":
                                    v108 := GenS14{F0: fi106, F1: fi49, F2: (v45 / ((1) | 1))}
                                    if (false || (false || true)) {
                                        _ = len(v6)
                                    }
                                    _ = v108
                                default:
                                    fmt.Println(v16)
                                }
                            }
                            v1[(fi106 << ((4294967295) & 63))] = (!v20)
                        }
                        v4 -= -1
                        switch (fi49 >> ((134) & 63)) {
                        case 127:
                            var v109 uint64 = 18446744073709551615
                            _ = v109
                        case -9223372036854775808:
                            v110 := GenS6{F0: (v16 >> ((829) & 63))}
                            _ = v110
                        case 1:
                            fmt.Println(v4, fi106)
                        }
                        v50 = (^(v16))
                        _ = fi106
                    }
                case 128:
                    switch (fi49 / (((fi49 % ((fi49) | 1))) | 1)) {
                    case -9223372036854775808:
                        switch "foo bar" {
                        case "":
                            if true {
                                v4 -= v4
                                for fi111 := 0; fi111 < 10; fi111++ {
                                    fmt.Println(fi49, v50)
                                    if (!((((!(464 >= 0)) || v20) || false) && (v20 || (v20 || (437 == 260))))) {
                                        delete(v3, v45)
                                        var v112 float64 = 3.14159
                                        v113 := GenS15{F0: (v50 << ((uint(v50)) & 63)), F1: 1, F2: 18446744073709551615}
                                        var v114 uint = 0
                                        v50 -= 255
                                        _ = v112
                                        _ = v113
                                        _ = v114
                                    }
                                    _ = fi111
                                }
                            }
                        case "\n\t":
                            switch (fi49 ^ (-(fi49))) {
                            case 4294967295:
                                var v115 int32 = (v4 % ((v4) | 1))
                                _ = v115
                            }
                        case "a":
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                        }
                    case 32767:
                        if v20 {
                            for fi116 := 0; fi116 < 6; fi116++ {
                                var v117 uint16 = (v16 ^ v50)
                                v4 = 42
                                v118 := GenS16{F0: -128, F1: fi116}
                                _ = fi116
                                _ = v117
                                _ = v118
                            }
                        }
                    case -779:
                        v119 := GenS8{F0: 18446744073709551615, F1: v20, F2: v45}
                        _ = v119
                    default:
                        for fi120 := 0; fi120 < 8; fi120++ {
                            v121 := GenS13{F0: v4, F1: -2147483648, F2: 1}
                            v50 = (^(v50))
                            _ = fi120
                            _ = v121
                        }
                    }
                case -2147483648:
                    fmt.Println(v16, fi49, fi49)
                default:
                    _ = len(v3)
                }
            } else {
                v122 := GenS17{F0: 4294967295, F1: (v45 | (^(v45))), F2: 42}
                v123 := GenS8{F0: 0, F1: ((!(0.1 >= (3.14159 / (0.1 / 3.14159)))) || (!v20)), F2: v45}
                _ = v122
                _ = v123
            }
            _ = fi49
            _ = v50
        }
    }
    var v124 uint8 = 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v124:", v124)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v124
    _ = v16
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v44
    _ = v45
    _ = v5
    _ = v6
    _ = v7
}
