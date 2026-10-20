package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16; F2 uint16 }
type GenS2 struct { F0 uint; F1 uint }
type GenS3 struct { F0 uint; F1 uint }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int32; F1 int64 }
type GenS6 struct { F0 int16; F1 uint64 }
type GenS7 struct { F0 uint8; F1 int }
type GenS8 struct { F0 int; F1 uint8 }
type GenS9 struct { F0 int32; F1 uint16 }
type GenS10 struct { F0 int }
type GenS11 struct { F0 int16; F1 int32 }
type GenS12 struct { F0 bool; F1 uint64; F2 uint8 }
type GenS13 struct { F0 uint16; F1 uint32 }
type GenS14 struct { F0 int16 }
type GenS15 struct { F0 uint64; F1 int64; F2 int64 }
type GenS16 struct { F0 bool; F1 int16; F2 uint }
type GenS17 struct { F0 bool }
type GenS18 struct { F0 bool }
type GenS19 struct { F0 int8 }
type GenS20 struct { F0 uint; F1 uint64 }
type GenS21 struct { F0 int64; F1 bool; F2 uint32 }
type GenS22 struct { F0 uint8 }
func genFn0(p0 uint32) uint64 {
    p0 -= 4294967295
    v7 := []int{}
    var v8 int64 = -1
    _ = p0
    _ = v7
    _ = v8
    return 18446744073709551615
}

func main() {
    var v1 int16 = -908
    var v2 int64 = 188
    v3 := []uint32{126, 42}
    var v4 int16 = -1
    var v5 int = 9223372036854775807
    var v6 int64 = v2
    for fi9 := 0; fi9 < 15; fi9++ {
        if len(v3) > 0 {
            v3[0] = 50
        }
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        for fi10 := 0; fi10 < 18; fi10++ {
            _ = len(v3)
            if (1e10 > 3.14159) {
                v11 := GenS1{F0: (v2 - -1), F1: 255, F2: 42}
                if false {
                    switch ("" + ("a" + ("" + "foo bar"))) {
                    case "\n\t":
                        if len(v3) > 0 {
                            v3[0] = 2147483647
                        }
                    }
                    var v12 float64 = (0.0 + (0.1 / 1.5))
                    if len(v3) > 0 {
                        v3[0] = 2
                    }
                    v13 := GenS1{F0: (v2 - v6), F1: 0, F2: 0}
                    _ = v12
                    _ = v13
                }
                if len(v3) > 0 {
                    v3[0] = 2147483647
                }
                _ = v11
            } else {
                v6 += int64(fi10)
                _ = len(v3)
                if (!(46 > 42)) {
                    if len(v3) > 0 {
                        v3[0] = 4294967295
                    }
                } else {
                    var v14 int = -749
                    v15 := GenS1{F0: (-(v6)), F1: 255, F2: 0}
                    fmt.Println(fi9)
                    _ = v14
                    _ = v15
                }
                switch (fi10 * fi9) {
                case 555:
                    if true {
                        v16 := GenS1{F0: (v2 % ((int64(v1)) | 1)), F1: 65535, F2: 255}
                        var v17 uint16 = 847
                        _ = v16
                        _ = v17
                    }
                case -568:
                    if (!true) {
                        var v18 int16 = v4
                        if true {
                            var v19 int8 = 1
                            v18 = (v18 << ((1) & 63))
                            v18 += (v4 & int16(v2))
                            _ = v19
                        }
                        _ = v18
                    } else {
                        v20 := GenS2{F0: 18446744073709551615, F1: 0}
                        if false {
                            var v21 float64 = 0.1
                            if len(v3) > 0 {
                                v3[0] = 0
                            }
                            var v22 uint = 281
                            _ = v21
                            _ = v22
                        }
                        if true {
                            var v23 bool = (true || (2147483647 < 1))
                            v5 += v5
                            _ = v23
                        }
                        _ = v20
                    }
                }
            }
            v6 -= v6
            _ = fi10
        }
        _ = fi9
    }
    switch (v5 | (v5 + (v5 << ((257) & 63)))) {
    case -9223372036854775808:
        switch ("hello" + "") {
        case "":
            if true {
                var v24 bool = (-1 != 1)
                for fi25 := 0; fi25 < 12; fi25++ {
                    for fi26 := 0; fi26 < 12; fi26++ {
                        if (438 != 42) {
                            var v27 float32 = 1e10
                            v1 = v1
                            fmt.Println(fi25, v1)
                            _ = v27
                        } else {
                            var v28 uint8 = 42
                            _ = v28
                        }
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        _ = fi26
                    }
                    v29 := GenS3{F0: 368, F1: 255}
                    _ = fi25
                    _ = v29
                }
                fmt.Println(v24)
                _ = v24
            } else {
                for fi30 := 0; fi30 < 16; fi30++ {
                    var v31 int = v5
                    if true {
                        switch v31 {
                        case -2147483648:
                            v32 := GenS2{F0: 18446744073709551615, F1: 2}
                            _ = v32
                        }
                        if true {
                            if ((!(0.1 > -2.25)) && true) {
                                for fi33 := 0; fi33 < 16; fi33++ {
                                    fmt.Println(fi30)
                                    _ = fi33
                                }
                                v2 += v2
                                switch 127 {
                                case 734:
                                    if true {
                                        v1 += 664
                                        for fi34 := 0; fi34 < 18; fi34++ {
                                            var v35 uint64 = 42
                                            _ = fi34
                                            _ = v35
                                        }
                                    } else {
                                        var v36 uint16 = 255
                                        _ = v36
                                    }
                                case 337:
                                    switch fi30 {
                                    case 32767:
                                        fmt.Println(v4)
                                    case -601:
                                        if (v2 == (v6 ^ v6)) {
                                            _ = len(v3)
                                            v37 := GenS4{F0: 42}
                                            _ = v37
                                        } else {
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                            if len(v3) > 0 {
                                                v3[0] = 65535
                                            }
                                        }
                                    }
                                default:
                                    v38 := GenS2{F0: 255, F1: 70}
                                    _ = v38
                                }
                            } else {
                                if false {
                                    if len(v3) > 0 {
                                        v3[0] = 65535
                                    }
                                    switch "\n\t" {
                                    case "\n\t":
                                        switch "foo bar" {
                                        case "\n\t":
                                            var v39 int16 = v4
                                            _ = v39
                                        default:
                                            v40 := GenS2{F0: 4294967295, F1: 664}
                                            _ = v40
                                        }
                                    default:
                                        if len(v3) > 0 {
                                            v3[0] = 4294967295
                                        }
                                    }
                                }
                            }
                            fmt.Println(v5, v31, v1)
                            var v41 int64 = (v6 * 1)
                            _ = v41
                        } else {
                            v42 := GenS5{F0: 235, F1: v2}
                            var v43 bool = false
                            if ((v4 << ((176) & 63)) > 1000) {
                                v44 := GenS6{F0: v4, F1: 2}
                                v45 := GenS2{F0: 18446744073709551615, F1: 42}
                                _ = v44
                                _ = v45
                            }
                            _ = v42
                            _ = v43
                        }
                    }
                    var v46 int32 = -2147483648
                    _ = fi30
                    _ = v31
                    _ = v46
                }
                if false {
                    for fi47 := 0; fi47 < 15; fi47++ {
                        if (!(65535 >= 596)) {
                            v48 := GenS6{F0: v4, F1: 42}
                            if true {
                                v6 = v2
                                _ = len(v3)
                            }
                            _ = v48
                        }
                        for fi49 := 0; fi49 < 4; fi49++ {
                            if (45 >= 18446744073709551615) {
                                v3 = append(v3, 169)
                                v50 := GenS7{F0: 42, F1: (^(fi49))}
                                var v51 int32 = 863
                                v6 = v2
                                _ = v50
                                _ = v51
                            }
                            _ = fi49
                        }
                        _ = fi47
                    }
                }
                v52 := GenS8{F0: (v5 << ((2) & 63)), F1: 2}
                _ = v52
            }
        case "foo bar":
            fmt.Println(v1)
        case "\n\t":
            for fi53 := 0; fi53 < 5; fi53++ {
                var v54 int16 = (v1 << ((1) & 63))
                _ = fi53
                _ = v54
            }
        }
    case -43:
        if false {
            if (2147483647 >= 42) {
                v5 += 1
                var v55 float32 = float32(((float64(-2.25) - (0.0 - (((1.5 - 1.5) / float64((-2.25 * 1e10))) / 1.0))) / (1.5 - (3.14159 * (1.5 + (float64((-2.25 + 1.5)) / ((1.5 - -2.25) - 3.14159)))))))
                _ = v55
            } else {
                v56 := GenS9{F0: -2147483648, F1: 322}
                _ = v56
            }
            for fi57 := 0; fi57 < 6; fi57++ {
                v1 = v4
                _ = fi57
            }
        } else {
            v2 += (v6 & (v2 + 42))
            switch "hello" {
            case "hello":
                v3 = append(v3, 0)
            default:
                v1 = (-(v4))
            }
        }
    case 42:
        for fi58 := 0; fi58 < 18; fi58++ {
            if (1 <= 255) {
                for fi59 := 0; fi59 < 2; fi59++ {
                    fmt.Println(fi59)
                    switch 42 {
                    case -840:
                        _ = len(v3)
                    case -754:
                        var v60 int16 = (^(v4))
                        _ = v60
                    default:
                        v2 += 1
                    }
                    if (104 > 100) {
                        var v61 int = (^(fi59))
                        _ = v61
                    } else {
                        v62 := GenS7{F0: 128, F1: v5}
                        var v63 int16 = 0
                        var v64 int8 = 0
                        _ = v62
                        _ = v63
                        _ = v64
                    }
                    _ = fi59
                }
            } else {
                var v65 uint32 = 572
                _ = v65
            }
            for fi66 := 0; fi66 < 1; fi66++ {
                var v67 int8 = 42
                switch (("hello" + ("" + "\n\t")) + "foo bar") {
                case "\n\t":
                    v4 -= (v4 - v1)
                default:
                    var v68 uint8 = 2
                    _ = v68
                }
                if len(v3) > 0 {
                    v3[0] = 4294967295
                }
                v69 := GenS9{F0: 42, F1: 32767}
                fmt.Println(v1, v1, v6)
                v6 = (v2 << ((2) & 63))
                _ = fi66
                _ = v67
                _ = v69
            }
            _ = fi58
        }
    default:
        var v70 string = ""
        _ = v70
    }
    for fi71 := 0; fi71 < 1; fi71++ {
        var v72 uint64 = genFn0(2147483648)
        for fi73 := 0; fi73 < 18; fi73++ {
            switch "" {
            case "foo bar":
                var v74 uint32 = 2147483647
                _ = v74
            }
            if len(v3) > 0 {
                v3[0] = 2147483648
            }
            switch (fi73 - 127) {
            case -353:
                v75 := GenS10{F0: 32768}
                _ = v75
            }
            _ = fi73
        }
        if len(v3) > 0 {
            v3[0] = 1
        }
        _ = fi71
        _ = v72
    }
    var v76 uint16 = 65535
    if (!(uint(v76) != 65535)) {
        if ((v6 & v6) >= (v2 | (^(v6)))) {
            v77 := GenS11{F0: v4, F1: -527}
            for fi78 := 0; fi78 < 9; fi78++ {
                var v79 float32 = (-2.25 + 1e10)
                _ = fi78
                _ = v79
            }
            _ = v77
        }
        v6 -= (v6 - v6)
    } else {
        if (v2 == 396) {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
        v80 := GenS12{F0: true, F1: 1, F2: 1}
        _ = len(v3)
        _ = v80
    }
    v1 = (-(v4))
    v4 = v1
    if false {
        switch 1 {
        case 32768:
            if (!true) {
                for fi81 := 0; fi81 < 4; fi81++ {
                    fmt.Println(v4, v76)
                    for fi82 := 0; fi82 < 9; fi82++ {
                        var v83 uint8 = 153
                        v83 = (v83 + 255)
                        _ = fi82
                        _ = v83
                    }
                    fmt.Println(v6, v1)
                    _ = fi81
                }
                fmt.Println(v1, v2)
                v6 -= (v6 ^ 0)
            } else {
                if true {
                    v3 = append(v3, 423)
                } else {
                    var v84 int32 = 42
                    switch v5 {
                    case 1:
                        var v85 int8 = 0
                        _ = v85
                    case -9223372036854775808:
                        v86 := GenS13{F0: (v76 - 676), F1: 2}
                        _ = v86
                    default:
                        var v87 uint32 = 42
                        _ = v87
                    }
                    _ = v84
                }
                if true {
                    v1 += (v1 & (v4 | (v4 / (((v4 + (v1 >> ((18446744073709551615) & 63)))) | 1))))
                    v1 = (-(v4))
                    switch (("foo bar" + (("foo bar" + "\n\t") + (("" + "\n\t") + ("hello" + ("" + "foo bar"))))) + "hello") {
                    case "\n\t":
                        v88 := GenS14{F0: v1}
                        _ = v88
                    default:
                        v1 = 1
                    }
                } else {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                }
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            }
        case 4294967295:
            v2 -= 42
        case 9223372036854775807:
            var v89 uint64 = 18446744073709551615
            _ = v89
        }
        v90 := GenS13{F0: v76, F1: uint32(v76)}
        if true {
            if true {
                if true {
                    v91 := GenS10{F0: v5}
                    for fi92 := 0; fi92 < 10; fi92++ {
                        if (!true) {
                            v4 -= (v1 % ((32767) | 1))
                            if false {
                                fmt.Println(v6)
                                switch 42 {
                                case -1:
                                    for fi93 := 0; fi93 < 20; fi93++ {
                                        if (742 >= 2147483647) {
                                            var v94 int8 = -128
                                            _ = v94
                                        } else {
                                            if (0.1 != float64(3.14159)) {
                                                v3 = append(v3, 42)
                                            }
                                            v95 := GenS10{F0: (fi92 + -9223372036854775808)}
                                            switch "foo bar" {
                                            case "foo bar":
                                                var v96 float32 = (1e10 / 1.0)
                                                _ = v96
                                            case "a":
                                                fmt.Println(fi92, v1, fi92)
                                            default:
                                                if false {
                                                    var v97 int = (v5 + fi93)
                                                    _ = v97
                                                }
                                            }
                                            v1 -= (v1 / ((v4) | 1))
                                            v76 += v76
                                            _ = v95
                                        }
                                        v98 := GenS15{F0: 319, F1: -3, F2: v2}
                                        var v99 int = 127
                                        _ = fi93
                                        _ = v98
                                        _ = v99
                                    }
                                case -842:
                                    for fi100 := 0; fi100 < 9; fi100++ {
                                        fmt.Println(v76, v6)
                                        var v101 string = ("a" + ("" + "\n\t"))
                                        _ = fi100
                                        _ = v101
                                    }
                                }
                            }
                        } else {
                            v4 = 1000
                            var v102 uint = 4294967295
                            _ = v102
                        }
                        var v103 uint8 = 2
                        switch "foo bar" {
                        case "":
                            switch "hello" {
                            case "a":
                                v6 = (v6 - v6)
                            }
                        default:
                            for fi104 := 0; fi104 < 10; fi104++ {
                                if (488 != 2147483647) {
                                    v105 := GenS16{F0: true, F1: 216, F2: 0}
                                    v106 := GenS17{F0: true}
                                    _ = v105
                                    _ = v106
                                }
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                var v107 uint = 65535
                                v107 = 65535
                                var v108 uint8 = v103
                                if true {
                                    switch -1 {
                                    case 128:
                                        switch fi104 {
                                        case 1:
                                            var v109 int8 = 100
                                            _ = v109
                                        }
                                    case 32768:
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                    case -807:
                                        var v110 int32 = 1
                                        _ = v110
                                    default:
                                        v6 -= (v6 + 407)
                                    }
                                } else {
                                    switch 1 {
                                    case 9223372036854775807:
                                        v111 := GenS18{F0: (fi92 >= fi104)}
                                        _ = v111
                                    case 122:
                                        v103 = 2
                                    case 127:
                                        if ((v108 >> ((v107) & 63)) > 2) {
                                            v6 = (v6 << (((v107 & (^(v107)))) & 63))
                                        }
                                    default:
                                        fmt.Println(v6)
                                    }
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                }
                                v2 = 42
                                if ((0.1 * 0.1) != -2.25) {
                                    switch "" {
                                    case "a":
                                        var v112 uint64 = uint64(v103)
                                        _ = v112
                                    case "":
                                        v113 := GenS19{F0: 100}
                                        v76 = v76
                                        v108 = uint8(v107)
                                        fmt.Println(v6, v76, v6)
                                        v107 -= 65535
                                        _ = v113
                                    case "\n\t":
                                        var v114 uint32 = 613
                                        _ = v114
                                    default:
                                        v115 := GenS20{F0: (v107 / ((255) | 1)), F1: 760}
                                        _ = v115
                                    }
                                    if false {
                                        switch 42 {
                                        case 9223372036854775807:
                                            v103 = (v108 >> ((v107) & 63))
                                        }
                                        var v116 int32 = 42
                                        var v117 uint = (v107 / (((v107 % ((42) | 1))) | 1))
                                        var v118 int8 = 5
                                        fmt.Println(v5, v6, v108)
                                        _ = v116
                                        _ = v117
                                        _ = v118
                                    } else {
                                        if len(v3) > 0 {
                                            v3[0] = 0
                                        }
                                        var v119 uint16 = 0
                                        v5 -= v5
                                        if (true || (false && (v107 >= v107))) {
                                            if false {
                                                v120 := GenS12{F0: (1e10 < 1e10), F1: 1, F2: 2}
                                                _ = v120
                                            } else {
                                                if false {
                                                    v108 = v103
                                                    v4 -= -1
                                                } else {
                                                    var v121 uint32 = 2
                                                    v5 += 1
                                                    _ = len(v3)
                                                    _ = v121
                                                }
                                                switch 2147483647 {
                                                case 32768:
                                                    v108 = 2
                                                case 1:
                                                    if ((!(42 != 100)) && false) {
                                                        _ = len(v3)
                                                        v119 = 1
                                                        switch (v5 / (((^(fi92))) | 1)) {
                                                        case -688:
                                                            if false {
                                                                v122 := GenS14{F0: v1}
                                                                fmt.Println(v1)
                                                                _ = v122
                                                            }
                                                        case 421:
                                                            if (int8(v108) < -39) {
                                                                v1 += v1
                                                                v2 = v6
                                                            }
                                                            v3 = append(v3, 0)
                                                            fmt.Println(v1, v4)
                                                        }
                                                    } else {
                                                        var v123 string = "a"
                                                        _ = v123
                                                    }
                                                    var v124 uint16 = v119
                                                    var v125 int16 = v1
                                                    v5 += v5
                                                    _ = v124
                                                    _ = v125
                                                default:
                                                    if (v4 != v4) {
                                                        var v126 string = (("\n\t" + "\n\t") + "a")
                                                        var v127 uint16 = 0
                                                        _ = v126
                                                        _ = v127
                                                    } else {
                                                        fmt.Println(v107)
                                                        fmt.Println(fi92, v2, fi92)
                                                    }
                                                }
                                            }
                                        }
                                        _ = v119
                                    }
                                    switch (fi104 & (fi104 % ((-1) | 1))) {
                                    case -479:
                                        v4 -= v1
                                        _ = len(v3)
                                        var v128 float64 = (float64(float32(-2.25)) * (0.1 / ((0.1 * -2.25) + (1e10 / (3.14159 / -2.25)))))
                                        v103 = 1
                                        v129 := GenS7{F0: (v108 ^ (v103 - v108)), F1: v5}
                                        _ = v128
                                        _ = v129
                                    case -379:
                                        fmt.Println(v107, v2)
                                    case -9223372036854775808:
                                        v130 := GenS5{F0: -2147483648, F1: 42}
                                        _ = v130
                                    }
                                }
                                var v131 int64 = 42
                                var v132 uint32 = 0
                                v133 := GenS1{F0: v131, F1: v76, F2: v76}
                                var v134 uint8 = (v103 & 140)
                                _ = fi104
                                _ = v107
                                _ = v108
                                _ = v131
                                _ = v132
                                _ = v133
                                _ = v134
                            }
                        }
                        _ = fi92
                        _ = v103
                    }
                    v2 = v2
                    _ = v91
                } else {
                    fmt.Println(v2, v5, v76)
                }
            } else {
                if (2 == 65535) {
                    v6 = (-(v2))
                    for fi135 := 0; fi135 < 6; fi135++ {
                        v1 += v4
                        var v136 int64 = (v6 << ((762) & 63))
                        if false {
                            switch "foo bar" {
                            case "":
                                fmt.Println(v4)
                            }
                            var v137 float64 = (1e10 + 0.1)
                            for fi138 := 0; fi138 < 17; fi138++ {
                                var v139 int = (fi135 ^ fi135)
                                v2 += (v136 % (((v136 & v2)) | 1))
                                if true {
                                    fmt.Println(v1, v2, v4)
                                    if ((false && true) && (!(0 == 65535))) {
                                        var v140 float64 = v137
                                        var v141 int64 = v136
                                        _ = v140
                                        _ = v141
                                    }
                                }
                                v76 -= v76
                                _ = fi138
                                _ = v139
                            }
                            _ = v137
                        } else {
                            v2 = v136
                        }
                        _ = fi135
                        _ = v136
                    }
                } else {
                    var v142 float64 = 0.0
                    _ = v142
                }
                var v143 uint16 = 65535
                _ = v143
            }
            if false {
                switch v5 {
                case 4294967295:
                    if false {
                        _ = len(v3)
                        v3 = append(v3, 0)
                    } else {
                        for fi144 := 0; fi144 < 4; fi144++ {
                            if false {
                                v145 := GenS2{F0: 4294967295, F1: 65535}
                                for fi146 := 0; fi146 < 12; fi146++ {
                                    v1 -= (v4 - v1)
                                    v147 := GenS19{F0: 127}
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    v148 := GenS9{F0: 42, F1: (^(v76))}
                                    _ = fi146
                                    _ = v147
                                    _ = v148
                                }
                                _ = v145
                            } else {
                                if true {
                                    var v149 uint32 = 42
                                    _ = v149
                                } else {
                                    for fi150 := 0; fi150 < 17; fi150++ {
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                        if true {
                                            var v151 int32 = -2147483648
                                            _ = v151
                                        }
                                        var v152 string = ""
                                        var v153 float32 = float32(0.0)
                                        if (!((int8(fi150) == int8(fi150)) && true)) {
                                            if (true || (!(v76 == v76))) {
                                                var v154 uint8 = 223
                                                _ = v154
                                            }
                                            var v155 float64 = (float64(v153) / (1.5 + (-2.25 * (float64(v153) / 0.1))))
                                            if false {
                                                v153 += v153
                                                v3 = append(v3, 643)
                                                var v156 int32 = -2147483648
                                                _ = v156
                                            }
                                            var v157 uint64 = 4294967295
                                            _ = v155
                                            _ = v157
                                        }
                                        var v158 uint64 = 0
                                        _ = fi150
                                        _ = v152
                                        _ = v153
                                        _ = v158
                                    }
                                    for fi159 := 0; fi159 < 18; fi159++ {
                                        v2 -= v2
                                        if ((v6 << ((2) & 63)) != (v6 >> ((0) & 63))) {
                                            _ = len(v3)
                                        } else {
                                            fmt.Println(fi144)
                                            v1 = 42
                                            v160 := GenS21{F0: (^(v6)), F1: (true || true), F2: 4294967295}
                                            fmt.Println(v5, v6)
                                            _ = v160
                                        }
                                        _ = fi159
                                    }
                                    v161 := GenS20{F0: 842, F1: 156}
                                    _ = v161
                                }
                                v3 = append(v3, 142)
                                var v162 uint = 2
                                _ = v162
                            }
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            _ = fi144
                        }
                        switch "" {
                        case "hello":
                            for fi163 := 0; fi163 < 1; fi163++ {
                                if (genFn0(177) >= 18446744073709551615) {
                                    v6 += v2
                                    v76 = v76
                                    fmt.Println(v2, fi163, v5)
                                } else {
                                    var v164 uint32 = 42
                                    _ = v164
                                }
                                for fi165 := 0; fi165 < 2; fi165++ {
                                    var v166 uint = 0
                                    var v167 int8 = -1
                                    _ = fi165
                                    _ = v166
                                    _ = v167
                                }
                                _ = fi163
                            }
                        case "a":
                            var v168 float64 = (float64((3.14159 - (-2.25 / (3.14159 - 1e10)))) / (0.1 * 0.1))
                            _ = v168
                        case "foo bar":
                            v169 := GenS21{F0: -819, F1: true, F2: 2147483648}
                            _ = v169
                        default:
                            v76 = (^(v76))
                        }
                    }
                case 42:
                    _ = len(v3)
                }
            } else {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            }
        } else {
            v76 = (^(v76))
        }
        _ = v90
    }
    fmt.Println(v2, v4)
    v76 = (v76 | v76)
    if len(v3) > 0 {
        v3[0] = 144
    }
    for fi170 := 0; fi170 < 11; fi170++ {
        if true {
            v171 := GenS12{F0: (1.5 < -2.25), F1: 2, F2: 1}
            v1 = v4
            _ = v171
        }
        fmt.Println(v76)
        _ = fi170
    }
    v3 = append(v3, 0)
    if false {
        v172 := GenS22{F0: 0}
        switch (v5 >> ((42) & 63)) {
        case 9223372036854775807:
            v173 := GenS6{F0: v4, F1: 4294967295}
            _ = v173
        }
        _ = v172
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println(v5, v76)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v76:", v76)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v76
}
