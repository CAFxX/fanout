package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint }
type GenS2 struct { F0 uint; F1 uint64 }
type GenS3 struct { F0 int64; F1 uint }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int64; F1 uint8 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 int32; F1 uint32; F2 int16 }
type GenS8 struct { F0 uint16 }
type GenS9 struct { F0 int; F1 int8; F2 int32 }
type GenS10 struct { F0 int8; F1 int8 }
type GenS11 struct { F0 int8; F1 uint16; F2 uint }
type GenS12 struct { F0 int64; F1 int8; F2 uint8 }
type GenS13 struct { F0 uint8; F1 uint32; F2 string }
type GenS14 struct { F0 uint }
type GenS15 struct { F0 uint8; F1 uint16 }
type GenS16 struct { F0 int64; F1 uint64 }
func genFn0(p0 int, p1 uint, p2 int64) uint64 {
    for fi6 := 0; fi6 < 2; fi6++ {
        var v7 float32 = (3.14159 / 0.1)
        _ = fi6
        _ = v7
    }
    v8 := GenS1{F0: 65535, F1: (p1 & 18446744073709551615)}
    for fi9 := 0; fi9 < 3; fi9++ {
        p0 += -1
        for fi10 := 0; fi10 < 8; fi10++ {
            p1 = p1
            var v11 bool = true
            _ = fi10
            _ = v11
        }
        _ = fi9
    }
    var v12 uint16 = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v8
    return 587
}
func genFn1() uint64 {
    v13 := []int{}
    switch 1 {
    case -2147483648:
        v14 := GenS2{F0: 0, F1: 4294967295}
        _ = v14
    case 9223372036854775807:
        for fi15 := 0; fi15 < 20; fi15++ {
            v16 := GenS1{F0: 0, F1: 65535}
            _ = fi15
            _ = v16
        }
    }
    var v17 uint64 = 2
    fmt.Println(v17)
    if (255 >= 42) {
        switch "a" {
        case "":
            if len(v13) > 0 {
                v13[0] = -948
            }
        case "hello":
            for fi18 := 0; fi18 < 13; fi18++ {
                v17 -= v17
                if len(v13) >= 2 {
                    v13 = v13[1:len(v13)]
                }
                _ = fi18
            }
        case "\n\t":
            fmt.Println(v17)
        }
        fmt.Println(v17)
    }
    _ = v13
    _ = v17
    return 0
}

func main() {
    var v1 int16 = -32768
    v2 := []string{(("" + "\n\t") + "hello"), "a", "foo bar"}
    var v3 int = 451
    v4 := make(map[uint32]int)
    v5 := make(map[uint32]bool)
    var v19 int16 = (v1 >> ((2) & 63))
    v1 += v19
    var v20 int16 = v19
    var v21 int32 = -1
    if false {
        var v22 float64 = -2.25
        _ = v22
    }
    switch (v3 + v3) {
    case -570:
        v23 := GenS2{F0: 821, F1: 4294967295}
        _ = v23
    case 1:
        if ("hello" < (("hello" + ((("\n\t" + ("foo bar" + "foo bar")) + "hello") + "\n\t")) + "")) {
            fmt.Println(v3, v1)
            v19 = (v20 - 32767)
            switch 4294967295 {
            case -9223372036854775808:
                for fi24 := 0; fi24 < 8; fi24++ {
                    switch "\n\t" {
                    case "a":
                        var v25 bool = true
                        _ = v25
                    case "hello":
                        var v26 float64 = (float64(1.5) / 1e10)
                        _ = v26
                    case "foo bar":
                        for fi27 := 0; fi27 < 5; fi27++ {
                            v1 -= v1
                            var v28 int16 = -641
                            switch fi24 {
                            case -2147483648:
                                switch fi27 {
                                case -1:
                                    var v29 uint = 42
                                    _ = v29
                                case -2147483648:
                                    v3 = fi27
                                case -242:
                                    switch (v3 >> ((18446744073709551615) & 63)) {
                                    case 0:
                                        if (false || false) {
                                            fmt.Println(v21)
                                        }
                                    case 32768:
                                        v1 += (v19 << ((202) & 63))
                                    case -182:
                                        v30 := GenS2{F0: 201, F1: genFn1()}
                                        _ = v30
                                    default:
                                        var v31 int8 = -60
                                        _ = v31
                                    }
                                default:
                                    var v32 int = (fi24 >> ((2) & 63))
                                    _ = v32
                                }
                            case 128:
                                if (v21 >= (v21 + v21)) {
                                    fmt.Println(v3)
                                    v20 = v1
                                    var v33 float32 = (1e10 / 1.0)
                                    _ = v33
                                } else {
                                    v20 += (-(v20))
                                    switch "foo bar" {
                                    case "":
                                        v3 = 32768
                                    case "\n\t":
                                        switch ("foo bar" + "a") {
                                        case "foo bar":
                                            var v34 int64 = -1
                                            _ = v34
                                        case "hello":
                                            var v35 float32 = 3.14159
                                            _ = v35
                                        case "\n\t":
                                            var v36 bool = ((0 == -1) && true)
                                            _ = v36
                                        default:
                                            v37 := GenS3{F0: 250, F1: 4294967295}
                                            v19 -= 42
                                            _ = v37
                                        }
                                    }
                                }
                            }
                            _ = fi27
                            _ = v28
                        }
                    default:
                        if (9223372036854775807 <= -9223372036854775808) {
                            v1 = v1
                            var v38 int32 = -796
                            _ = v38
                        } else {
                            if (false || (false || false)) {
                                if (!true) {
                                    v39 := GenS3{F0: -9223372036854775808, F1: 65535}
                                    _ = v39
                                }
                            } else {
                                if (!(-2.25 < ((1.5 + 0.1) / (float64(-2.25) / (-2.25 * 1.5))))) {
                                    switch v3 {
                                    case 2147483647:
                                        if (!false) {
                                            v20 = -749
                                            v20 = (v20 & v20)
                                            switch "" {
                                            case "foo bar":
                                                for fi40 := 0; fi40 < 8; fi40++ {
                                                    var v41 int8 = -128
                                                    if (true && true) {
                                                        switch 738 {
                                                        case 783:
                                                            if true {
                                                                if false {
                                                                    switch "a" {
                                                                    case "\n\t":
                                                                        _ = len(v2)
                                                                    }
                                                                    switch "\n\t" {
                                                                    case "hello":
                                                                        v42 := GenS1{F0: 65535, F1: 465}
                                                                        _ = v42
                                                                    case "":
                                                                        var v43 int32 = 42
                                                                        _ = v43
                                                                    case "foo bar":
                                                                        if (0 > 1) {
                                                                            var v44 int64 = -9223372036854775808
                                                                            _ = v44
                                                                        }
                                                                        v45 := GenS3{F0: -237, F1: 81}
                                                                        _ = v45
                                                                    default:
                                                                        v46 := GenS4{F0: false}
                                                                        _ = v46
                                                                    }
                                                                }
                                                                v20 = 149
                                                                v41 = 0
                                                            }
                                                            v2 = append(v2, "a")
                                                            if (759 < 465) {
                                                                v47 := GenS1{F0: 391, F1: 2}
                                                                if (42 <= 42) {
                                                                    var v48 uint = 65535
                                                                    var v49 uint32 = 527
                                                                    v49 += 2
                                                                    _ = v48
                                                                    _ = v49
                                                                } else {
                                                                    if (-506 < -1) {
                                                                        if (-2.25 == ((1.5 + (float64(0.0) - 0.1)) / 1.0)) {
                                                                            v50 := GenS2{F0: 551, F1: genFn1()}
                                                                            var v51 uint32 = 49
                                                                            var v52 uint64 = 834
                                                                            _ = v50
                                                                            _ = v51
                                                                            _ = v52
                                                                        } else {
                                                                            if true {
                                                                                var v53 float64 = -2.25
                                                                                switch "\n\t" {
                                                                                case "hello":
                                                                                    switch (fi24 / (((fi40 >> ((2) & 63))) | 1)) {
                                                                                    case -314:
                                                                                        fmt.Println(fi24, v3, v21)
                                                                                    }
                                                                                default:
                                                                                    v2 = append(v2, "hello")
                                                                                }
                                                                                _ = v53
                                                                            } else {
                                                                                _ = len(v2)
                                                                                if false {
                                                                                    if (42 < 65535) {
                                                                                        switch 9223372036854775807 {
                                                                                        case 32767:
                                                                                            v54 := GenS1{F0: 1, F1: 1}
                                                                                            _ = v54
                                                                                        case 42:
                                                                                            if len(v2) > 0 {
                                                                                                v2[0] = "a"
                                                                                            }
                                                                                        case 679:
                                                                                            if true {
                                                                                                fmt.Println(fi40)
                                                                                                if (true && (!(v20 == v20))) {
                                                                                                    var v55 int = (fi24 - fi40)
                                                                                                    var v56 int = 32768
                                                                                                    _ = v55
                                                                                                    _ = v56
                                                                                                } else {
                                                                                                    if true {
                                                                                                        if (!((1 >= (-(v20))) || (42 != 26))) {
                                                                                                            var v57 int32 = v21
                                                                                                            _ = v57
                                                                                                        } else {
                                                                                                            v58 := GenS5{F0: 887, F1: 42}
                                                                                                            v59 := GenS6{F0: 42}
                                                                                                            _ = v58
                                                                                                            _ = v59
                                                                                                        }
                                                                                                        var v60 int8 = (v41 * -128)
                                                                                                        var v61 uint = 18446744073709551615
                                                                                                        _ = v60
                                                                                                        _ = v61
                                                                                                    } else {
                                                                                                        v19 = (v1 ^ v19)
                                                                                                        switch "a" {
                                                                                                        case "":
                                                                                                            switch "" {
                                                                                                            case "hello":
                                                                                                                var v62 bool = true
                                                                                                                _ = v62
                                                                                                            case "foo bar":
                                                                                                                switch (fi24 >> ((18446744073709551615) & 63)) {
                                                                                                                case 777:
                                                                                                                    var v63 int8 = int8(fi40)
                                                                                                                    _ = v63
                                                                                                                case 1:
                                                                                                                    fmt.Println(v19, v20)
                                                                                                                default:
                                                                                                                    var v64 uint8 = 80
                                                                                                                    _ = v64
                                                                                                                }
                                                                                                            case "a":
                                                                                                                v65 := GenS7{F0: v21, F1: 1, F2: (^(v1))}
                                                                                                                _ = v65
                                                                                                            default:
                                                                                                                var v66 int32 = v21
                                                                                                                _ = v66
                                                                                                            }
                                                                                                        }
                                                                                                    }
                                                                                                }
                                                                                                v21 += -2147483648
                                                                                            }
                                                                                        default:
                                                                                            fmt.Println(v41)
                                                                                        }
                                                                                        v67 := GenS8{F0: 98}
                                                                                        v21 = v21
                                                                                        _ = v67
                                                                                    } else {
                                                                                        fmt.Println(v3, v20)
                                                                                    }
                                                                                }
                                                                                switch "" {
                                                                                case "foo bar":
                                                                                    switch "" {
                                                                                    case "foo bar":
                                                                                        _ = len(v5)
                                                                                    case "hello":
                                                                                        _ = len(v2)
                                                                                    default:
                                                                                        if true {
                                                                                            v1 -= v19
                                                                                        }
                                                                                    }
                                                                                case "a":
                                                                                    switch "" {
                                                                                    case "a":
                                                                                        v68 := GenS5{F0: 1, F1: 1}
                                                                                        _ = v68
                                                                                    case "":
                                                                                        var v69 int = 1
                                                                                        _ = v69
                                                                                    case "\n\t":
                                                                                        v21 -= v21
                                                                                    }
                                                                                    fmt.Println(fi40, v19)
                                                                                default:
                                                                                    var v70 int16 = v20
                                                                                    _ = v70
                                                                                }
                                                                                v1 = (v1 ^ (v1 + (-(v20))))
                                                                            }
                                                                            delete(v4, 2147483648)
                                                                            if true {
                                                                                var v71 uint64 = 904
                                                                                _ = v71
                                                                            }
                                                                            if false {
                                                                                v21 = v21
                                                                            }
                                                                            if len(v2) >= 2 {
                                                                                v2 = v2[1:len(v2)]
                                                                            }
                                                                        }
                                                                        var v72 int = -833
                                                                        v73 := GenS1{F0: 1, F1: 2}
                                                                        _ = v72
                                                                        _ = v73
                                                                    }
                                                                    v74 := GenS6{F0: 0}
                                                                    _ = v74
                                                                }
                                                                _ = v47
                                                            }
                                                        case 32767:
                                                            if (42 < -882) {
                                                                var v75 int16 = (v1 << ((2) & 63))
                                                                v76 := GenS5{F0: 9223372036854775807, F1: 128}
                                                                _ = v75
                                                                _ = v76
                                                            } else {
                                                                fmt.Println(v1)
                                                                v41 -= v41
                                                                delete(v4, 2147483647)
                                                            }
                                                        case 42:
                                                            var v77 uint16 = 52
                                                            _ = v77
                                                        }
                                                        var v78 int32 = v21
                                                        v79 := GenS9{F0: (v3 * (fi40 - (fi24 / ((int(v20)) | 1)))), F1: v41, F2: (v21 | (^(v21)))}
                                                        _ = v78
                                                        _ = v79
                                                    } else {
                                                        if (18446744073709551615 >= 398) {
                                                            var v80 uint16 = 1
                                                            v2 = append(v2, "")
                                                            var v81 int = (^(fi40))
                                                            var v82 int8 = v41
                                                            _ = v80
                                                            _ = v81
                                                            _ = v82
                                                        }
                                                    }
                                                    switch ("hello" + ("hello" + "foo bar")) {
                                                    case "hello":
                                                        var v83 uint16 = 1
                                                        _ = v83
                                                    default:
                                                        v3 = 0
                                                    }
                                                    _ = fi40
                                                    _ = v41
                                                }
                                            case "a":
                                                var v84 int16 = (-(v19))
                                                _ = v84
                                            default:
                                                v85 := GenS8{F0: 662}
                                                _ = v85
                                            }
                                        } else {
                                            var v86 float32 = -2.25
                                            _ = v86
                                        }
                                    default:
                                        var v87 uint16 = 492
                                        _ = v87
                                    }
                                    var v88 int8 = 46
                                    _ = v88
                                } else {
                                    var v89 uint32 = 4294967295
                                    v21 -= -1
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    _ = v89
                                }
                                v90 := GenS5{F0: int64(v20), F1: 128}
                                switch "" {
                                case "":
                                    v91 := v5[1]
                                    _ = v91
                                case "hello":
                                    switch "" {
                                    case "a":
                                        if true {
                                            for fi92 := 0; fi92 < 19; fi92++ {
                                                fmt.Println(v3)
                                                v19 = (v19 + v19)
                                                if true {
                                                    _ = len(v5)
                                                    v20 = v20
                                                    var v93 int64 = 9223372036854775807
                                                    v2 = append(v2, "a")
                                                    var v94 uint8 = 108
                                                    _ = v93
                                                    _ = v94
                                                }
                                                _ = fi92
                                            }
                                            switch (fi24 / ((v3) | 1)) {
                                            case 826:
                                                if false {
                                                    fmt.Println(v1, v21)
                                                    var v95 float32 = 3.14159
                                                    _ = v95
                                                } else {
                                                    v19 = (v19 << ((866) & 63))
                                                    v3 -= (^(fi24))
                                                    var v96 int64 = 9223372036854775807
                                                    _ = v96
                                                }
                                            default:
                                                v97 := GenS1{F0: 2, F1: 255}
                                                _ = v97
                                            }
                                            if false {
                                                v2 = append(v2, "foo bar")
                                                v2 = append(v2, "")
                                                v2 = append(v2, "hello")
                                            } else {
                                                if true {
                                                    if false {
                                                        for fi98 := 0; fi98 < 10; fi98++ {
                                                            switch "" {
                                                            case "":
                                                                v99 := v4[1]
                                                                _ = v99
                                                            case "a":
                                                                var v100 uint32 = 4294967295
                                                                _ = v100
                                                            case "\n\t":
                                                                if len(v2) >= 2 {
                                                                    v2 = v2[1:len(v2)]
                                                                }
                                                            default:
                                                                v5[0] = (float64((0.1 / 1.0)) != 0.1)
                                                            }
                                                            v2 = append(v2, (((("a" + (("" + "a") + "foo bar")) + (("hello" + "\n\t") + "\n\t")) + "a") + "a"))
                                                            switch (("a" + "hello") + "foo bar") {
                                                            case "foo bar":
                                                                switch "foo bar" {
                                                                case "a":
                                                                    v3 -= v3
                                                                default:
                                                                    v3 += fi98
                                                                }
                                                            case "":
                                                                v21 = (^(v21))
                                                            case "a":
                                                                v101 := v4[65535]
                                                                _ = v101
                                                            default:
                                                                var v102 bool = true
                                                                _ = v102
                                                            }
                                                            _ = fi98
                                                        }
                                                    } else {
                                                        v103 := GenS9{F0: 42, F1: -103, F2: 0}
                                                        v104 := GenS10{F0: -128, F1: 42}
                                                        _ = v103
                                                        _ = v104
                                                    }
                                                    for fi105 := 0; fi105 < 12; fi105++ {
                                                        v106 := v4[954]
                                                        var v107 float32 = (1e10 * ((1e10 * 1.5) / -2.25))
                                                        v108 := GenS9{F0: (v3 + (fi105 * 4294967295)), F1: 1, F2: -1}
                                                        _ = fi105
                                                        _ = v106
                                                        _ = v107
                                                        _ = v108
                                                    }
                                                } else {
                                                    delete(v4, 0)
                                                }
                                                _ = len(v4)
                                                var v109 string = (("a" + (("a" + "") + ("\n\t" + "hello"))) + "a")
                                                _ = v109
                                            }
                                        } else {
                                            fmt.Println(v1)
                                            fmt.Println(v20, fi24)
                                            var v110 int32 = (v21 * (v21 / ((v21) | 1)))
                                            _ = v110
                                        }
                                    case "":
                                        for fi111 := 0; fi111 < 4; fi111++ {
                                            var v112 int64 = 9223372036854775807
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                            _ = fi111
                                            _ = v112
                                        }
                                    default:
                                        for fi113 := 0; fi113 < 20; fi113++ {
                                            if false {
                                                fmt.Println(v3, v21, v1)
                                                v1 = (-(v20))
                                                if true {
                                                    v114 := GenS10{F0: 0, F1: -1}
                                                    v115 := GenS11{F0: -24, F1: 1, F2: 18446744073709551615}
                                                    if true {
                                                        v116 := GenS12{F0: 1, F1: 100, F2: 1}
                                                        _ = v116
                                                    }
                                                    _ = v114
                                                    _ = v115
                                                }
                                            }
                                            delete(v4, 0)
                                            v117 := GenS3{F0: 42, F1: 65535}
                                            _ = len(v2)
                                            _ = len(v2)
                                            v1 = (v1 << ((526) & 63))
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                            fmt.Println(v1, v3)
                                            v118 := v5[21]
                                            v118 = (-9223372036854775808 > -8)
                                            var v119 uint8 = 0
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                            var v120 uint64 = 4294967295
                                            _ = fi113
                                            _ = v117
                                            _ = v118
                                            _ = v119
                                            _ = v120
                                        }
                                    }
                                case "a":
                                    v5[65535] = false
                                default:
                                    v1 = v19
                                }
                                _ = v90
                            }
                            v19 -= (v19 >> ((1) & 63))
                            var v121 uint8 = 2
                            _ = v121
                        }
                    }
                    var v122 int64 = 9223372036854775807
                    _ = fi24
                    _ = v122
                }
            case 42:
                var v123 uint16 = 1
                _ = v123
            case 242:
                v19 += -1
            default:
                fmt.Println(v3, v20)
            }
        } else {
            if true {
                if false {
                    v21 = v21
                }
                v124 := GenS10{F0: -1, F1: -128}
                _ = v124
            } else {
                for fi125 := 0; fi125 < 15; fi125++ {
                    var v126 int = v3
                    _ = fi125
                    _ = v126
                }
                v127 := GenS6{F0: -9223372036854775808}
                fmt.Println(v21)
                _ = v127
            }
            for fi128 := 0; fi128 < 3; fi128++ {
                _ = len(v2)
                _ = fi128
            }
            v1 -= (v1 + v1)
        }
    case 0:
        switch "\n\t" {
        case "":
            switch (v3 + (^(v3))) {
            case 9223372036854775807:
                v19 -= v1
            case 6:
                v20 = -783
            }
        case "foo bar":
            var v129 float32 = (0.1 + float32(3.14159))
            _ = v129
        default:
            v130 := v5[0]
            _ = v130
        }
    default:
        switch "" {
        case "":
            var v131 int32 = v21
            _ = v131
        case "foo bar":
            switch "hello" {
            case "\n\t":
                for fi132 := 0; fi132 < 14; fi132++ {
                    var v133 int = fi132
                    for fi134 := 0; fi134 < 4; fi134++ {
                        var v135 float64 = 1.5
                        fmt.Println(v20)
                        if (9223372036854775807 < 1) {
                            if len(v2) > 0 {
                                v2[0] = ""
                            }
                        }
                        v136 := GenS13{F0: 128, F1: 65535, F2: "hello"}
                        var v137 float64 = 0.1
                        _ = fi134
                        _ = v135
                        _ = v136
                        _ = v137
                    }
                    v138 := GenS14{F0: 65535}
                    _ = fi132
                    _ = v133
                    _ = v138
                }
            case "hello":
                v3 = (v3 & v3)
            }
        default:
            v139 := GenS12{F0: -9223372036854775808, F1: -128, F2: 197}
            _ = v139
        }
    }
    if false {
        v140 := GenS7{F0: v21, F1: 304, F2: int16(v3)}
        _ = v140
    } else {
        delete(v4, 2147483648)
        v20 += 1
        v141 := GenS15{F0: 218, F1: 42}
        _ = v141
    }
    v142 := GenS12{F0: -1, F1: 0, F2: 96}
    _ = len(v2)
    if ((v3 / (((v3 + 32768)) | 1)) == v3) {
        fmt.Println(v3, v1, v3)
        if true {
            v19 += (v19 + v1)
        } else {
            if (v21 != 2147483647) {
                v3 -= v3
                switch ("foo bar" + "") {
                case "foo bar":
                    v21 = -187
                default:
                    var v143 bool = (42 <= 127)
                    _ = v143
                }
                v144 := GenS16{F0: 450, F1: 1}
                _ = v144
            } else {
                if (!(-260 >= -535)) {
                    var v145 string = ""
                    fmt.Println(v19, v145)
                    v145 = "hello"
                    _ = v145
                }
            }
            for fi146 := 0; fi146 < 1; fi146++ {
                var v147 bool = false
                _ = fi146
                _ = v147
            }
            v20 += v20
        }
    } else {
        v4[775] = v3
    }
    if ((v3 | v3) != (v3 & (v3 + -696))) {
        if true {
            v20 = 32767
            v20 -= 32767
        } else {
            var v148 uint8 = 0
            _ = v148
        }
        if false {
            for fi149 := 0; fi149 < 1; fi149++ {
                for fi150 := 0; fi150 < 8; fi150++ {
                    v19 = (v19 >> ((4294967295) & 63))
                    var v151 int = v3
                    v19 = v20
                    _ = fi150
                    _ = v151
                }
                for fi152 := 0; fi152 < 11; fi152++ {
                    fmt.Println(v19)
                    _ = fi152
                }
                _ = fi149
            }
            v153 := GenS10{F0: 42, F1: 1}
            _ = v153
        }
    } else {
        v3 = (v3 + v3)
    }
    var v154 float32 = ((float32(-2.25) + -2.25) + (1e10 + 1e10))
    v2 = append(v2, ("" + (("\n\t" + ("hello" + (("hello" + "a") + ""))) + "a")))
    v154 = (((v154 + v154) - (1e10 * 1.5)) * (v154 * v154))
    v21 = 42
    var v155 uint8 = 160
    fmt.Println("v1:", v1)
    fmt.Println("v154:", v154)
    fmt.Println("v155:", v155)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v142
    _ = v154
    _ = v155
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
