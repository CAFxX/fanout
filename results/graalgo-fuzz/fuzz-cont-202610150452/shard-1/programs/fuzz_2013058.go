package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16; F2 int16 }
type GenS2 struct { F0 int32; F1 int }
type GenS3 struct { F0 uint16; F1 uint }
type GenS4 struct { F0 int64; F1 uint; F2 string }
type GenS5 struct { F0 uint; F1 uint64; F2 uint32 }
type GenS6 struct { F0 int32; F1 int32 }
type GenS7 struct { F0 string; F1 string }
type GenS8 struct { F0 int64; F1 int8; F2 uint32 }
type GenS9 struct { F0 string; F1 uint16; F2 int64 }
type GenS10 struct { F0 uint64; F1 int32; F2 uint32 }
type GenS11 struct { F0 uint8; F1 uint; F2 uint32 }
type GenS12 struct { F0 bool; F1 uint64; F2 uint64 }
type GenS13 struct { F0 int64; F1 string; F2 string }
type GenS14 struct { F0 uint64 }
type GenS15 struct { F0 int32; F1 uint }
type GenS16 struct { F0 bool; F1 uint8 }
type GenS17 struct { F0 uint64; F1 uint64; F2 bool }
type GenS18 struct { F0 int64; F1 int32 }
type GenS19 struct { F0 int16; F1 int16; F2 uint32 }
type GenS20 struct { F0 uint32; F1 string; F2 bool }
type GenS21 struct { F0 int16 }
type GenS22 struct { F0 int32; F1 string; F2 int64 }
type GenS23 struct { F0 uint8; F1 int64 }
type GenS24 struct { F0 int; F1 uint }
type GenS25 struct { F0 bool }
type GenS26 struct { F0 int32; F1 int16; F2 int8 }
type GenS27 struct { F0 string }
func genFn0(p0 uint64) uint64 {
    var v6 uint32 = 4294967295
    var v7 int = 32768
    if (-32768 <= 32767) {
        v6 += v6
    }
    for fi8 := 0; fi8 < 5; fi8++ {
        v6 = 2147483647
        _ = fi8
    }
    if (-369 < 1) {
        for fi9 := 0; fi9 < 14; fi9++ {
            switch "" {
            case "hello":
                fmt.Println(v7, v6)
            case "\n\t":
                if false {
                    v10 := GenS1{F0: 2, F1: 407, F2: -713}
                    _ = v10
                }
            default:
                v11 := []int{128}
                _ = v11
            }
            _ = fi9
        }
        var v12 float64 = -2.25
        _ = v12
    }
    _ = p0
    _ = v6
    _ = v7
    return (p0 >> ((709) & 63))
}
func genFn1(p0 int8, p1 uint32, p2 bool) uint8 {
    if false {
        if true {
            v13 := GenS1{F0: 255, F1: 1000, F2: -32768}
            _ = v13
        } else {
            var v14 uint32 = 65535
            if (!(98 > 42)) {
                fmt.Println(p2)
                var v15 int = -368
                for fi16 := 0; fi16 < 20; fi16++ {
                    p1 -= (p1 % (((v14 % (((p1 << ((4294967295) & 63))) | 1))) | 1))
                    var v17 int64 = -9223372036854775808
                    switch fi16 {
                    case -61:
                        switch ("\n\t" + "a") {
                        case "":
                            var v18 float64 = (float64(1.5) / 1.0)
                            _ = v18
                        default:
                            if (!p2) {
                                switch "" {
                                case "a":
                                    fmt.Println(p0, p2, p2)
                                }
                            }
                        }
                    case -9223372036854775808:
                        for fi19 := 0; fi19 < 14; fi19++ {
                            var v20 uint32 = 0
                            switch fi16 {
                            case -928:
                                v17 += v17
                            default:
                                v21 := GenS1{F0: 4294967295, F1: 0, F2: 215}
                                _ = v21
                            }
                            _ = fi19
                            _ = v20
                        }
                    case 42:
                        v15 += v15
                    default:
                        for fi22 := 0; fi22 < 3; fi22++ {
                            if false {
                                v23 := GenS1{F0: 1, F1: -32768, F2: 42}
                                v24 := GenS2{F0: 2147483647, F1: 2147483647}
                                v25 := []int{fi22, 925}
                                _ = v23
                                _ = v24
                                _ = v25
                            }
                            if false {
                                switch (fi22 << ((458) & 63)) {
                                case -9223372036854775808:
                                    v15 -= (v15 | (-(fi16)))
                                }
                                var v26 uint16 = 65535
                                switch int(v26) {
                                case 128:
                                    switch "foo bar" {
                                    case "a":
                                        v27 := []int{(fi16 % ((fi22) | 1)), v15, fi22, v15}
                                        _ = v27
                                    case "":
                                        if p2 {
                                            var v28 int32 = 742
                                            fmt.Println(p1)
                                            fmt.Println(fi22, fi22, v26)
                                            switch (fi22 * 128) {
                                            case 128:
                                                fmt.Println(v28, v28)
                                            case 458:
                                                v29 := GenS1{F0: 2, F1: 42, F2: -32768}
                                                _ = v29
                                            default:
                                                v30 := GenS3{F0: (v26 / ((65535) | 1)), F1: 921}
                                                fmt.Println(p0, p2, fi16)
                                                _ = v30
                                            }
                                            _ = v28
                                        }
                                    case "\n\t":
                                        v26 = 65535
                                    default:
                                        v31 := []int{fi22}
                                        _ = v31
                                    }
                                case -96:
                                    if (p2 || p2) {
                                        v32 := GenS4{F0: (-(v17)), F1: 18446744073709551615, F2: "\n\t"}
                                        _ = v32
                                    }
                                    var v33 int64 = v17
                                    v34 := GenS5{F0: 64, F1: 988, F2: (p1 >> ((2) & 63))}
                                    _ = v33
                                    _ = v34
                                case 606:
                                    if ((p2 || p2) && true) {
                                        if false {
                                            v35 := []int{v15, (fi16 >> ((0) & 63))}
                                            _ = v35
                                        }
                                        var v36 float64 = 1e10
                                        switch ("hello" + "hello") {
                                        case "a":
                                            var v37 float32 = (0.0 + -2.25)
                                            _ = v37
                                        case "\n\t":
                                            switch "a" {
                                            case "a":
                                                v38 := []int{}
                                                _ = v38
                                            default:
                                                p1 = 2147483648
                                            }
                                            var v39 uint8 = 1
                                            if (0 < 0) {
                                                var v40 int8 = -128
                                                v41 := GenS5{F0: 65535, F1: 0, F2: (p1 >> ((65535) & 63))}
                                                _ = v40
                                                _ = v41
                                            }
                                            _ = v39
                                        case "":
                                            var v42 uint32 = (v14 % (((v14 & (v14 << ((0) & 63)))) | 1))
                                            _ = v42
                                        }
                                        _ = v36
                                    }
                                default:
                                    v43 := []int{fi22, v15}
                                    _ = v43
                                }
                                _ = v26
                            }
                            p0 += (^(p0))
                            if p2 {
                                fmt.Println(fi16)
                                switch v15 {
                                case 42:
                                    v44 := []int{fi16}
                                    _ = v44
                                case 32767:
                                    v45 := GenS2{F0: int32(fi16), F1: -936}
                                    _ = v45
                                default:
                                    var v46 int64 = 758
                                    v47 := GenS6{F0: 0, F1: -142}
                                    v15 -= int(v14)
                                    _ = v46
                                    _ = v47
                                }
                                v48 := GenS7{F0: "foo bar", F1: "foo bar"}
                                var v49 string = "a"
                                _ = v48
                                _ = v49
                            }
                            _ = fi22
                        }
                    }
                    _ = fi16
                    _ = v17
                }
                _ = v15
            }
            if true {
                v50 := GenS8{F0: -1, F1: -41, F2: 0}
                var v51 int32 = 0
                fmt.Println(p1, p2, v51)
                _ = v50
                _ = v51
            }
            _ = v14
        }
        if p2 {
            v52 := []int{4294967295, -627, 127, 42}
            v53 := GenS8{F0: 42, F1: p0, F2: (p1 ^ p1)}
            _ = v52
            _ = v53
        }
    }
    p0 = (p0 - p0)
    p1 = (^(p1))
    var v54 int64 = -814
    v55 := GenS9{F0: "\n\t", F1: 65535, F2: v54}
    _ = p0
    _ = p1
    _ = p2
    _ = v54
    _ = v55
    return 128
}

func main() {
    v1 := []int{32767, 32767, -654}
    v2 := []int{-151}
    var v3 int32 = 0
    var v4 uint16 = 255
    v5 := []string{}
    v1 = append(v1, -172)
    if true {
        fmt.Println(v4)
    } else {
        if len(v2) > 0 {
            v2[0] = 2147483647
        }
        v4 -= 1
    }
    switch ((("a" + ("foo bar" + "foo bar")) + "\n\t") + ("a" + "\n\t")) {
    case "foo bar":
        switch -9223372036854775808 {
        case 9223372036854775807:
            v3 = (v3 / ((v3) | 1))
        case 42:
            _ = len(v5)
        default:
            if (-1 != 9223372036854775807) {
                v1 = append(v1, -2147483648)
                switch 32768 {
                case -161:
                    _ = len(v5)
                case 338:
                    v4 = (v4 ^ 32767)
                default:
                    v4 += 255
                }
            } else {
                var v56 uint32 = 2147483647
                var v57 int16 = 32767
                _ = v56
                _ = v57
            }
        }
    case "hello":
        var v58 int64 = 1
        _ = v58
    default:
        v59 := GenS2{F0: -1, F1: -516}
        _ = v59
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println(v4, v3)
    if false {
        if false {
            for fi60 := 0; fi60 < 3; fi60++ {
                v4 += (v4 & (^(v4)))
                _ = fi60
            }
            var v61 float32 = 0.1
            v62 := GenS2{F0: -905, F1: 4294967295}
            _ = v61
            _ = v62
        } else {
            var v63 int8 = 86
            switch "" {
            case "\n\t":
                switch -2147483648 {
                case 9223372036854775807:
                    var v64 bool = false
                    _ = v64
                case 128:
                    switch 4294967295 {
                    case -2147483648:
                        if false {
                            if len(v2) > 0 {
                                v2[0] = 9223372036854775807
                            }
                        }
                    }
                case 32768:
                    for fi65 := 0; fi65 < 7; fi65++ {
                        v66 := GenS10{F0: 42, F1: (v3 % ((v3) | 1)), F2: 485}
                        v67 := GenS11{F0: 90, F1: 1, F2: 961}
                        switch fi65 {
                        case 1:
                            for fi68 := 0; fi68 < 19; fi68++ {
                                switch fi68 {
                                case 4294967295:
                                    v4 += v4
                                default:
                                    var v69 bool = (255 == 42)
                                    _ = v69
                                }
                                _ = fi68
                            }
                        case 4294967295:
                            var v70 string = (("a" + "hello") + "foo bar")
                            _ = v70
                        case 9223372036854775807:
                            switch (fi65 >> ((18446744073709551615) & 63)) {
                            case 42:
                                _ = len(v5)
                            case -1:
                                for fi71 := 0; fi71 < 10; fi71++ {
                                    v72 := GenS5{F0: 0, F1: 4294967295, F2: 2147483647}
                                    var v73 uint32 = 0
                                    _ = fi71
                                    _ = v72
                                    _ = v73
                                }
                            case -2147483648:
                                if (42 == 127) {
                                    var v74 bool = false
                                    if (42 == 1000) {
                                        if (!(true || false)) {
                                            for fi75 := 0; fi75 < 20; fi75++ {
                                                var v76 int8 = (v63 / ((v63) | 1))
                                                v1 = append(v1, fi65)
                                                v5 = append(v5, ("" + (("" + "hello") + ("a" + "\n\t"))))
                                                v77 := GenS12{F0: (genFn1((^(v63)), 2, true) >= uint8(v4)), F1: 18446744073709551615, F2: 674}
                                                _ = fi75
                                                _ = v76
                                                _ = v77
                                            }
                                            var v78 float32 = 1.5
                                            var v79 int64 = 42
                                            _ = v78
                                            _ = v79
                                        } else {
                                            fmt.Println(v63, v74, v74)
                                        }
                                        fmt.Println(fi65, v74, fi65)
                                    } else {
                                        if v74 {
                                            _ = len(v5)
                                        }
                                        v3 = -1
                                        for fi80 := 0; fi80 < 17; fi80++ {
                                            v2 = append(v2, (fi80 | 127))
                                            var v81 uint = 2
                                            if (0 == 0) {
                                                if (3.14159 <= 3.14159) {
                                                    v74 = (2 >= (v81 % ((865) | 1)))
                                                    v63 = v63
                                                }
                                            }
                                            _ = fi80
                                            _ = v81
                                        }
                                    }
                                    _ = v74
                                }
                            default:
                                for fi82 := 0; fi82 < 1; fi82++ {
                                    if ((!(v3 != v3)) || false) {
                                        if ((false || (!true)) || false) {
                                            if len(v1) > 0 {
                                                v1[0] = int(v4)
                                            }
                                        } else {
                                            var v83 uint16 = (v4 << ((406) & 63))
                                            fmt.Println(fi82)
                                            var v84 int8 = int8(fi65)
                                            _ = v83
                                            _ = v84
                                        }
                                    } else {
                                        var v85 uint8 = 240
                                        _ = len(v1)
                                        _ = v85
                                    }
                                    _ = len(v1)
                                    _ = fi82
                                }
                            }
                        default:
                            if (uint(v4) >= 1) {
                                v86 := GenS6{F0: v3, F1: v3}
                                fmt.Println(v4, fi65, fi65)
                                _ = v86
                            } else {
                                fmt.Println(v3, v4)
                                v4 = v4
                            }
                        }
                        _ = fi65
                        _ = v66
                        _ = v67
                    }
                }
            case "foo bar":
                _ = len(v2)
            default:
                if (255 == 2) {
                    if true {
                        if false {
                            if false {
                                var v87 int8 = 100
                                if ((false || (!((false && (!true)) || (((3.14159 - 0.0) >= ((0.0 / (3.14159 / 0.1)) - -2.25)) && ((!(false && (false || true))) && false))))) || true) {
                                    v88 := GenS13{F0: -9223372036854775808, F1: "", F2: "hello"}
                                    if true {
                                        switch ("foo bar" + "\n\t") {
                                        case "\n\t":
                                            v4 += v4
                                        case "foo bar":
                                            var v89 bool = false
                                            _ = v89
                                        }
                                        v87 = (v63 / (((^(v63))) | 1))
                                        switch 354 {
                                        case 0:
                                            fmt.Println(v87, v87, v87)
                                        case -2147483648:
                                            for fi90 := 0; fi90 < 19; fi90++ {
                                                v63 = 10
                                                fmt.Println(v63, fi90, v4)
                                                _ = fi90
                                            }
                                        case -1:
                                            if false {
                                                if (948 <= -9223372036854775808) {
                                                    _ = len(v5)
                                                }
                                                var v91 uint = 520
                                                _ = v91
                                            }
                                        default:
                                            switch -9223372036854775808 {
                                            case -474:
                                                for fi92 := 0; fi92 < 18; fi92++ {
                                                    switch fi92 {
                                                    case 2147483647:
                                                        for fi93 := 0; fi93 < 13; fi93++ {
                                                            var v94 bool = true
                                                            if v94 {
                                                                v95 := GenS13{F0: -1, F1: (("hello" + "") + ("hello" + "foo bar")), F2: "\n\t"}
                                                                var v96 float64 = ((float64((-2.25 - (-2.25 / float32(((-2.25 / 3.14159) * ((0.1 * ((3.14159 - (((3.14159 / ((-2.25 + 3.14159) + 0.0)) * -2.25) - 0.0)) / (1.5 - 1e10))) * (float64((3.14159 * 3.14159)) * 1.5))))))) * (float64(0.1) / 1e10)) + float64(0.0))
                                                                _ = v95
                                                                _ = v96
                                                            } else {
                                                                fmt.Println(v94)
                                                            }
                                                            _ = fi93
                                                            _ = v94
                                                        }
                                                    }
                                                    var v97 uint32 = 42
                                                    _ = fi92
                                                    _ = v97
                                                }
                                            case 128:
                                                var v98 int16 = -39
                                                _ = v98
                                            case 32768:
                                                var v99 float64 = -2.25
                                                _ = v99
                                            default:
                                                switch "hello" {
                                                case "hello":
                                                    v87 = 100
                                                case "\n\t":
                                                    var v100 int32 = (v3 + -2147483648)
                                                    _ = v100
                                                }
                                            }
                                        }
                                    }
                                    for fi101 := 0; fi101 < 4; fi101++ {
                                        if true {
                                            for fi102 := 0; fi102 < 7; fi102++ {
                                                v3 = -1
                                                _ = fi102
                                            }
                                        }
                                        for fi103 := 0; fi103 < 16; fi103++ {
                                            v3 += v3
                                            _ = fi103
                                        }
                                        _ = fi101
                                    }
                                    _ = v88
                                }
                                _ = v87
                            }
                            var v104 uint16 = (^(v4))
                            _ = v104
                        } else {
                            v105 := GenS13{F0: -87, F1: ("foo bar" + ("a" + "foo bar")), F2: "a"}
                            v106 := GenS14{F0: 865}
                            _ = v105
                            _ = v106
                        }
                        var v107 uint32 = 0
                        v108 := GenS15{F0: v3, F1: 2}
                        _ = v107
                        _ = v108
                    } else {
                        switch "\n\t" {
                        case "hello":
                            fmt.Println(v4, v4)
                        case "":
                            v63 = (v63 - (v63 % (((v63 | (-(v63)))) | 1)))
                        case "foo bar":
                            if true {
                                v3 += v3
                                switch 127 {
                                case -1:
                                    var v109 uint64 = genFn0(4294967295)
                                    _ = v109
                                case 891:
                                    fmt.Println(v4, v63, v3)
                                }
                            }
                        default:
                            switch "hello" {
                            case "\n\t":
                                for fi110 := 0; fi110 < 11; fi110++ {
                                    var v111 uint16 = v4
                                    _ = fi110
                                    _ = v111
                                }
                            case "":
                                var v112 float32 = 3.14159
                                _ = v112
                            case "foo bar":
                                var v113 int16 = -1
                                _ = v113
                            }
                        }
                        v2 = append(v2, -9223372036854775808)
                        v4 = (v4 << ((614) & 63))
                    }
                    v3 = v3
                    if (0.0 > (3.14159 * -2.25)) {
                        switch 47 {
                        case 9223372036854775807:
                            v63 = v63
                        case 725:
                            switch ("hello" + ("hello" + "a")) {
                            case "a":
                                var v114 int64 = 42
                                _ = v114
                            }
                        case 128:
                            for fi115 := 0; fi115 < 17; fi115++ {
                                v63 = (v63 ^ (v63 & (v63 - v63)))
                                v1 = append(v1, fi115)
                                _ = fi115
                            }
                        }
                        var v116 string = "\n\t"
                        var v117 float32 = (1e10 / 1.5)
                        _ = v116
                        _ = v117
                    }
                } else {
                    var v118 string = ("foo bar" + (((("\n\t" + "hello") + "") + "") + "a"))
                    if (!(false || true)) {
                        for fi119 := 0; fi119 < 3; fi119++ {
                            var v120 int32 = v3
                            if false {
                                if ((!true) && (303 > fi119)) {
                                    if false {
                                        fmt.Println(v118)
                                        v118 = "\n\t"
                                    } else {
                                        v4 = (v4 ^ (v4 * 0))
                                        if (v63 < (v63 ^ -7)) {
                                            v121 := GenS16{F0: ((2 < 506) && false), F1: 0}
                                            switch v118 {
                                            case "a":
                                                var v122 uint32 = 350
                                                _ = v122
                                            case "":
                                                if len(v2) > 0 {
                                                    v2[0] = (fi119 << ((65535) & 63))
                                                }
                                            default:
                                                v118 = v118
                                            }
                                            switch ((v118 + v118) + ("foo bar" + "")) {
                                            case "":
                                                v123 := GenS17{F0: 660, F1: 18446744073709551615, F2: ((!true) && false)}
                                                _ = v123
                                            default:
                                                v63 += v63
                                            }
                                            _ = v121
                                        } else {
                                            fmt.Println(v118, v63, v4)
                                        }
                                    }
                                }
                                for fi124 := 0; fi124 < 8; fi124++ {
                                    _ = len(v1)
                                    v125 := GenS18{F0: -1, F1: (v3 ^ (v120 % ((v120) | 1)))}
                                    _ = fi124
                                    _ = v125
                                }
                            } else {
                                if false {
                                    switch fi119 {
                                    case 2147483647:
                                        var v126 float64 = -2.25
                                        _ = v126
                                    default:
                                        var v127 uint64 = 0
                                        _ = v127
                                    }
                                }
                            }
                            _ = fi119
                            _ = v120
                        }
                    }
                    v63 = 1
                    _ = v118
                }
            }
            _ = v63
        }
        var v128 string = ("hello" + "\n\t")
        if (true || (!false)) {
            if (373 > v4) {
                v128 = "a"
            }
            if true {
                var v129 int = 42
                v130 := GenS4{F0: -1, F1: 42, F2: "foo bar"}
                v128 = v128
                _ = v129
                _ = v130
            }
            for fi131 := 0; fi131 < 11; fi131++ {
                var v132 uint8 = 0
                _ = fi131
                _ = v132
            }
        } else {
            v4 += (^(v4))
            if ((3.14159 + 1.5) >= (-2.25 - 1e10)) {
                var v133 uint32 = 2147483647
                var v134 int = -667
                _ = v133
                _ = v134
            } else {
                var v135 float64 = 1e10
                v136 := GenS19{F0: 400, F1: 1000, F2: 42}
                _ = v135
                _ = v136
            }
        }
        _ = v128
    }
    var v137 bool = ((("hello" + (((("foo bar" + "foo bar") + ("foo bar" + "foo bar")) + "hello") + "foo bar")) + "\n\t") == (("hello" + (("foo bar" + ("a" + "a")) + ("a" + "hello"))) + "a"))
    var v138 float64 = 3.14159
    var v139 float64 = ((1.5 - 0.1) + (3.14159 / (v138 + 0.1)))
    if (genFn1(100, 1, v137) >= genFn1(-1, 2, (9223372036854775807 == -9223372036854775808))) {
        for fi140 := 0; fi140 < 19; fi140++ {
            var v141 float32 = 1e10
            _ = len(v1)
            v141 = v141
            _ = fi140
            _ = v141
        }
        v3 += v3
    } else {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        if (!false) {
            if true {
                switch "a" {
                case "foo bar":
                    for fi142 := 0; fi142 < 20; fi142++ {
                        if (true || v137) {
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        }
                        var v143 uint64 = 0
                        _ = fi142
                        _ = v143
                    }
                case "a":
                    _ = len(v5)
                }
            } else {
                if len(v2) > 0 {
                    v2[0] = 56
                }
                for fi144 := 0; fi144 < 7; fi144++ {
                    v145 := GenS14{F0: 42}
                    _ = fi144
                    _ = v145
                }
                switch 32768 {
                case -9223372036854775808:
                    _ = len(v2)
                default:
                    var v146 int16 = 0
                    _ = v146
                }
            }
        } else {
            if (4294967295 != 478) {
                v1 = append(v1, -9223372036854775808)
                switch "foo bar" {
                case "foo bar":
                    if len(v2) > 0 {
                        v2[0] = -9223372036854775808
                    }
                case "hello":
                    for fi147 := 0; fi147 < 15; fi147++ {
                        var v148 float32 = (3.14159 + -2.25)
                        if (uint8(v4) >= 1) {
                            v149 := GenS1{F0: 255, F1: 1000, F2: 331}
                            fmt.Println(v138)
                            _ = v149
                        } else {
                            v139 = (((v139 * 0.1) + ((((float64(v148) - v139) * v138) / (float64(v148) / (float64(0.1) / (((((v138 + 1.5) - float64(v148)) + -2.25) - -2.25) / v138)))) - (float64(v148) + (v138 * 3.14159)))) + (float64((v148 - v148)) + v139))
                            var v150 uint = 255
                            fmt.Println(v3)
                            _ = v150
                        }
                        _ = fi147
                        _ = v148
                    }
                case "":
                    if false {
                        v151 := GenS8{F0: -1, F1: -44, F2: 0}
                        _ = v151
                    }
                }
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            } else {
                for fi152 := 0; fi152 < 18; fi152++ {
                    v139 += v139
                    _ = fi152
                }
            }
            v3 -= v3
        }
        _ = len(v2)
    }
    if v137 {
        var v153 uint = 42
        switch "\n\t" {
        case "a":
            for fi154 := 0; fi154 < 2; fi154++ {
                fmt.Println(v3)
                for fi155 := 0; fi155 < 3; fi155++ {
                    v156 := GenS6{F0: (v3 << ((v153) & 63)), F1: (v3 ^ 1)}
                    if ((!(((3.14159 / ((-2.25 + ((v139 * v139) + 0.1)) / 1.5)) * v139) <= v138)) && v137) {
                        v138 += float64(1e10)
                        switch (fi154 >> ((v153) & 63)) {
                        case -459:
                            v153 = v153
                        case 128:
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        case 1:
                            switch -818 {
                            case -9223372036854775808:
                                switch (-(fi154)) {
                                case 128:
                                    if ((fi154 * fi154) != (fi155 & 142)) {
                                        var v157 int = 9223372036854775807
                                        v137 = v137
                                        if len(v1) > 0 {
                                            v1[0] = (-(fi155))
                                        }
                                        _ = v157
                                    } else {
                                        var v158 float64 = v138
                                        _ = len(v2)
                                        _ = v158
                                    }
                                case 0:
                                    v159 := GenS8{F0: 42, F1: 74, F2: 58}
                                    _ = v159
                                }
                            case 108:
                                if (!(-1 == -13)) {
                                    var v160 int16 = 0
                                    switch "" {
                                    case "\n\t":
                                        v139 = v139
                                    case "hello":
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                    case "foo bar":
                                        var v161 float32 = 0.1
                                        _ = v161
                                    default:
                                        v153 = 24
                                    }
                                    v138 = (1.5 / 3.14159)
                                    _ = v160
                                } else {
                                    v162 := GenS20{F0: 4294967295, F1: "a", F2: v137}
                                    if v137 {
                                        if v137 {
                                            fmt.Println(v139, v138, v137)
                                            v163 := GenS21{F0: 42}
                                            v164 := GenS7{F0: "foo bar", F1: "\n\t"}
                                            _ = v163
                                            _ = v164
                                        }
                                        if true {
                                            var v165 uint64 = genFn0(2)
                                            var v166 float32 = 1.5
                                            _ = v165
                                            _ = v166
                                        } else {
                                            var v167 int = fi155
                                            var v168 int64 = 0
                                            _ = v167
                                            _ = v168
                                        }
                                        v169 := GenS22{F0: (^(v3)), F1: "hello", F2: int64(fi155)}
                                        _ = v169
                                    }
                                    var v170 uint32 = 42
                                    _ = v162
                                    _ = v170
                                }
                            case -1:
                                var v171 float32 = 0.0
                                _ = v171
                            }
                        default:
                            var v172 uint = v153
                            _ = v172
                        }
                    }
                    var v173 uint8 = 153
                    _ = fi155
                    _ = v156
                    _ = v173
                }
                _ = fi154
            }
        case "hello":
            v174 := GenS23{F0: 128, F1: 0}
            _ = v174
        case "\n\t":
            v4 = (^(v4))
        default:
            switch "\n\t" {
            case "\n\t":
                v175 := GenS13{F0: -1, F1: "foo bar", F2: ""}
                _ = v175
            case "":
                v176 := GenS18{F0: -315, F1: -1}
                _ = v176
            case "foo bar":
                var v177 uint16 = (v4 + v4)
                _ = v177
            default:
                v138 -= 0.1
            }
        }
        _ = v153
    } else {
        _ = len(v5)
        v138 += float64((1.5 * 1e10))
    }
    if v137 {
        for fi178 := 0; fi178 < 4; fi178++ {
            fmt.Println(v3, v3)
            _ = fi178
        }
        var v179 uint64 = 4294967295
        _ = v179
    } else {
        var v180 float64 = 1e10
        for fi181 := 0; fi181 < 13; fi181++ {
            v138 += (v138 / 1.0)
            switch fi181 {
            case -679:
                v138 -= (float64(-2.25) * v139)
            case 128:
                v139 += -2.25
            case 2147483647:
                var v182 uint16 = v4
                _ = v182
            default:
                fmt.Println(fi181, v3)
            }
            v183 := GenS24{F0: (fi181 & fi181), F1: 65535}
            _ = fi181
            _ = v183
        }
        _ = v180
    }
    if v137 {
        v139 -= (-2.25 + v138)
    }
    switch 127 {
    case 88:
        if false {
            var v184 string = ("a" + "foo bar")
            for fi185 := 0; fi185 < 20; fi185++ {
                switch "" {
                case "hello":
                    if false {
                        v186 := GenS25{F0: v137}
                        v187 := GenS7{F0: v184, F1: (((v184 + v184) + (v184 + (v184 + v184))) + v184)}
                        v139 = (v139 / v138)
                        _ = v186
                        _ = v187
                    }
                case "foo bar":
                    v188 := GenS22{F0: v3, F1: v184, F2: -9223372036854775808}
                    _ = v188
                case "\n\t":
                    _ = len(v1)
                default:
                    fmt.Println(v184)
                }
                _ = fi185
            }
            _ = v184
        } else {
            v3 += (v3 | -611)
        }
    case -803:
        for fi189 := 0; fi189 < 15; fi189++ {
            var v190 int16 = 0
            v191 := GenS11{F0: 2, F1: 84, F2: 2147483647}
            _ = fi189
            _ = v190
            _ = v191
        }
    default:
        for fi192 := 0; fi192 < 12; fi192++ {
            for fi193 := 0; fi193 < 12; fi193++ {
                if len(v2) > 0 {
                    v2[0] = (fi193 - fi192)
                }
                _ = fi193
            }
            var v194 int8 = 42
            _ = fi192
            _ = v194
        }
    }
    for fi195 := 0; fi195 < 3; fi195++ {
        switch "" {
        case "hello":
            if v137 {
                switch ((("foo bar" + ("a" + "hello")) + ("hello" + ("a" + "foo bar"))) + ("\n\t" + "foo bar")) {
                case "foo bar":
                    if v137 {
                        v4 = v4
                        v139 = v139
                        var v196 uint32 = 78
                        _ = v196
                    }
                case "a":
                    if len(v1) > 0 {
                        v1[0] = fi195
                    }
                }
                switch fi195 {
                case -227:
                    var v197 int16 = 42
                    _ = v197
                case -1:
                    var v198 float64 = v138
                    _ = v198
                }
            }
        default:
            for fi199 := 0; fi199 < 3; fi199++ {
                switch ("\n\t" + ((("" + "hello") + "foo bar") + "\n\t")) {
                case "foo bar":
                    v200 := GenS26{F0: v3, F1: 0, F2: -127}
                    _ = v200
                case "a":
                    switch fi199 {
                    case 128:
                        v201 := GenS27{F0: "\n\t"}
                        _ = v201
                    }
                }
                _ = fi199
            }
        }
        var v202 int64 = 42
        _ = fi195
        _ = v202
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v137:", v137)
    fmt.Println("v138:", v138)
    fmt.Println("v139:", v139)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v137
    _ = v138
    _ = v139
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
