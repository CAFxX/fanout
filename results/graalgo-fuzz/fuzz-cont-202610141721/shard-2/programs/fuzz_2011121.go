package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint64; F1 int32 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 bool; F1 uint32 }
type GenS6 struct { F0 uint64; F1 int; F2 int64 }
type GenS7 struct { F0 uint16; F1 uint }
type GenS8 struct { F0 uint32; F1 int32 }
type GenS9 struct { F0 int32; F1 int16; F2 int32 }
type GenS10 struct { F0 uint16 }
type GenS11 struct { F0 uint64; F1 uint; F2 uint16 }
type GenS12 struct { F0 string; F1 int8; F2 uint8 }
type GenS13 struct { F0 int }
type GenS14 struct { F0 uint16 }
type GenS15 struct { F0 uint8; F1 uint32; F2 uint16 }
type GenS16 struct { F0 string; F1 uint64 }
type GenS17 struct { F0 bool; F1 string }
type GenS18 struct { F0 bool; F1 int32 }
func genFn0(p0 string) int64 {
    var v6 bool = (99 > 4294967295)
    p0 = p0
    fmt.Println(p0)
    for fi7 := 0; fi7 < 17; fi7++ {
        v8 := GenS1{F0: ((!(0 >= 1)) || false)}
        fmt.Println(fi7, v6)
        var v9 uint64 = 22
        _ = fi7
        _ = v8
        _ = v9
    }
    _ = p0
    _ = v6
    return -1
}
func genFn1(p0 int) string {
    var v10 string = (("foo bar" + "\n\t") + "")
    v11 := GenS2{F0: 1}
    if true {
        v12 := GenS3{F0: 811, F1: 1}
        var v13 int = (p0 * 680)
        _ = v12
        _ = v13
    }
    switch (p0 >> ((42) & 63)) {
    case 32767:
        if true {
            v10 = v10
            switch (p0 ^ p0) {
            case 0:
                v10 = "a"
            case 435:
                v14 := []int{p0, p0}
                _ = v14
            case 2147483647:
                v15 := GenS4{F0: p0}
                _ = v15
            default:
                v16 := GenS4{F0: (-(p0))}
                _ = v16
            }
            if (1 == 100) {
                if (42 >= -2147483648) {
                    p0 = (p0 / (((p0 << ((65535) & 63))) | 1))
                    if (505 < 0) {
                        var v17 float32 = (1.5 * -2.25)
                        switch v10 {
                        case "":
                            switch 4294967295 {
                            case -2147483648:
                                switch p0 {
                                case 2147483647:
                                    v10 = (v10 + (v10 + (v10 + v10)))
                                case 32767:
                                    if false {
                                        var v18 uint8 = 2
                                        v18 = 1
                                        _ = v18
                                    } else {
                                        v17 -= 0.1
                                    }
                                case -1:
                                    if true {
                                        var v19 int32 = -2147483648
                                        switch "foo bar" {
                                        case "\n\t":
                                            for fi20 := 0; fi20 < 1; fi20++ {
                                                if true {
                                                    switch 145 {
                                                    case 127:
                                                        var v21 uint32 = 2147483647
                                                        _ = v21
                                                    default:
                                                        var v22 int8 = -119
                                                        _ = v22
                                                    }
                                                    switch fi20 {
                                                    case 127:
                                                        var v23 int32 = 1
                                                        _ = v23
                                                    case 32768:
                                                        v17 = (v17 - v17)
                                                    }
                                                } else {
                                                    switch (fi20 - 42) {
                                                    case -811:
                                                        switch v10 {
                                                        case "":
                                                            var v24 int64 = -9223372036854775808
                                                            _ = v24
                                                        case "hello":
                                                            fmt.Println(fi20, v10, v19)
                                                        }
                                                    case 42:
                                                        if true {
                                                            v10 = ("\n\t" + ("" + v10))
                                                            for fi25 := 0; fi25 < 6; fi25++ {
                                                                var v26 string = ("\n\t" + (v10 + ("a" + (v10 + "hello"))))
                                                                var v27 int64 = -1
                                                                v28 := []int{809, (fi20 + (p0 % (((fi25 - (fi20 ^ p0))) | 1)))}
                                                                _ = fi25
                                                                _ = v26
                                                                _ = v27
                                                                _ = v28
                                                            }
                                                        }
                                                    case 230:
                                                        fmt.Println(v17)
                                                    }
                                                    var v29 float64 = (0.0 / 3.14159)
                                                    var v30 int32 = -642
                                                    _ = v29
                                                    _ = v30
                                                }
                                                v19 = (v19 | (v19 - v19))
                                                _ = fi20
                                            }
                                        }
                                        v31 := GenS1{F0: true}
                                        _ = v19
                                        _ = v31
                                    }
                                default:
                                    for fi32 := 0; fi32 < 14; fi32++ {
                                        switch (fi32 ^ (^(p0))) {
                                        case 0:
                                            v33 := []int{p0}
                                            _ = v33
                                        case 1:
                                            if false {
                                                v10 = "foo bar"
                                                v34 := []int{fi32, 4294967295, p0}
                                                _ = v34
                                            }
                                        case -2147483648:
                                            fmt.Println(p0, fi32)
                                        default:
                                            v10 = v10
                                        }
                                        p0 = fi32
                                        _ = fi32
                                    }
                                }
                            case -504:
                                var v35 int8 = 0
                                _ = v35
                            }
                        case "a":
                            if true {
                                if (!false) {
                                    fmt.Println(v10)
                                }
                                v36 := []int{p0, -72, 1, p0}
                                if len(v36) >= 2 {
                                    v36 = v36[1:len(v36)]
                                }
                                _ = v36
                            } else {
                                switch -186 {
                                case 0:
                                    p0 = (p0 | p0)
                                case 32767:
                                    switch p0 {
                                    case 4294967295:
                                        var v37 string = "hello"
                                        _ = v37
                                    default:
                                        p0 -= p0
                                    }
                                case -621:
                                    v17 = (((v17 * (v17 * float32(0.1))) + 0.1) - ((v17 + ((float32(1e10) * 1e10) + v17)) - 1.5))
                                default:
                                    var v38 int8 = 100
                                    _ = v38
                                }
                                v10 = (v10 + v10)
                            }
                        }
                        fmt.Println(v10, v10)
                        _ = v17
                    } else {
                        v39 := GenS2{F0: 255}
                        _ = v39
                    }
                } else {
                    var v40 uint = 65535
                    for fi41 := 0; fi41 < 20; fi41++ {
                        var v42 int16 = 0
                        _ = fi41
                        _ = v42
                    }
                    _ = v40
                }
                if true {
                    switch -605 {
                    case 32768:
                        var v43 int8 = 127
                        _ = v43
                    }
                    if false {
                        if (968 < 732) {
                            v44 := []int{p0}
                            _ = v44
                        } else {
                            if false {
                                var v45 int8 = -1
                                _ = v45
                            }
                            p0 = p0
                            var v46 uint64 = 18446744073709551615
                            _ = v46
                        }
                    }
                    var v47 float32 = (1e10 + 1e10)
                    _ = v47
                } else {
                    switch (p0 * (p0 ^ p0)) {
                    case 1:
                        v48 := []int{}
                        _ = v48
                    default:
                        v10 = (v10 + (v10 + (v10 + v10)))
                    }
                    for fi49 := 0; fi49 < 3; fi49++ {
                        v50 := []int{(p0 * p0)}
                        if false {
                            if len(v50) > 0 {
                                v50[0] = 1
                            }
                        } else {
                            _ = len(v50)
                            var v51 int8 = 1
                            _ = v51
                        }
                        if len(v50) >= 2 {
                            v50 = v50[1:len(v50)]
                        }
                        _ = fi49
                        _ = v50
                    }
                    switch (-(p0)) {
                    case 80:
                        var v52 uint = 42
                        _ = v52
                    case 1:
                        switch (p0 ^ (p0 - (p0 * -1))) {
                        case 2147483647:
                            fmt.Println(v10, p0)
                        case 128:
                            v10 = (v10 + (v10 + v10))
                        }
                    case 9223372036854775807:
                        if (1 <= 127) {
                            fmt.Println(p0)
                            if ((1.5 - 0.0) > 3.14159) {
                                v10 = v10
                            }
                            for fi53 := 0; fi53 < 4; fi53++ {
                                var v54 int64 = genFn0(v10)
                                if (2 < 1) {
                                    var v55 float32 = (0.1 * 3.14159)
                                    fmt.Println(p0)
                                    _ = v55
                                }
                                _ = fi53
                                _ = v54
                            }
                        }
                    }
                }
                fmt.Println(v10, v10)
            }
        }
    case 4294967295:
        v56 := GenS2{F0: 109}
        _ = v56
    case 751:
        var v57 int8 = 1
        _ = v57
    default:
        if false {
            var v58 string = v10
            for fi59 := 0; fi59 < 13; fi59++ {
                fmt.Println(p0, p0, p0)
                _ = fi59
            }
            _ = v58
        } else {
            var v60 uint8 = 127
            var v61 uint64 = 42
            _ = v60
            _ = v61
        }
    }
    _ = p0
    _ = v10
    _ = v11
    return (v10 + v10)
}

func main() {
    var v1 int16 = -987
    v2 := []uint32{0}
    v3 := make(map[string]bool)
    v4 := make(map[uint32]uint32)
    v5 := make(map[uint32]int)
    if ((127 > 32767) && (2 <= 1)) {
        v62 := GenS3{F0: 937, F1: 448}
        var v63 float64 = (3.14159 * float64((3.14159 / (0.1 / 3.14159))))
        _ = v62
        _ = v63
    } else {
        if false {
            v64 := GenS5{F0: false, F1: 1}
            if (!true) {
                v1 = v1
                v1 = (v1 * (v1 << ((880) & 63)))
            } else {
                for fi65 := 0; fi65 < 8; fi65++ {
                    fmt.Println(v1)
                    switch "" {
                    case "hello":
                        v1 = 780
                    default:
                        if false {
                            if false {
                                v1 += (v1 | v1)
                                v1 = v1
                            } else {
                                if false {
                                    v66 := GenS6{F0: 18446744073709551615, F1: (fi65 ^ (-(fi65))), F2: 527}
                                    v67 := GenS3{F0: 910, F1: -913}
                                    v68 := GenS7{F0: 2, F1: 2}
                                    _ = v66
                                    _ = v67
                                    _ = v68
                                } else {
                                    fmt.Println(v1)
                                }
                            }
                        } else {
                            fmt.Println(fi65)
                        }
                    }
                    v1 = (^(v1))
                    _ = fi65
                }
                switch ("\n\t" + "\n\t") {
                case "\n\t":
                    if true {
                        v1 -= v1
                        v1 -= v1
                    } else {
                        fmt.Println(v1)
                    }
                case "hello":
                    fmt.Println(v1)
                default:
                    switch "hello" {
                    case "a":
                        if (835 == (v1 & (v1 >> ((42) & 63)))) {
                            if (!false) {
                                if (1 <= 2147483647) {
                                    var v69 uint16 = 606
                                    _ = v69
                                }
                            }
                        } else {
                            v3[("" + "foo bar")] = false
                            if (true && (!true)) {
                                for fi70 := 0; fi70 < 18; fi70++ {
                                    v71 := GenS3{F0: 1, F1: 1}
                                    for fi72 := 0; fi72 < 16; fi72++ {
                                        v1 = v1
                                        _ = fi72
                                    }
                                    _ = fi70
                                    _ = v71
                                }
                                if true {
                                    fmt.Println(v1)
                                } else {
                                    v1 -= (v1 * v1)
                                    for fi73 := 0; fi73 < 13; fi73++ {
                                        var v74 int32 = -2147483648
                                        var v75 float32 = 1.5
                                        _ = fi73
                                        _ = v74
                                        _ = v75
                                    }
                                    v1 = v1
                                }
                            } else {
                                if false {
                                    for fi76 := 0; fi76 < 1; fi76++ {
                                        fmt.Println(v1, v1)
                                        v77 := GenS6{F0: 1, F1: fi76, F2: 0}
                                        _ = fi76
                                        _ = v77
                                    }
                                    if true {
                                        v1 = (^(v1))
                                        if (9223372036854775807 > 127) {
                                            var v78 uint8 = 2
                                            if (2 != 524) {
                                                v79 := GenS7{F0: 1, F1: 42}
                                                _ = v79
                                            }
                                            _ = v78
                                        }
                                    } else {
                                        if false {
                                            switch "" {
                                            case "":
                                                switch "a" {
                                                case "\n\t":
                                                    if false {
                                                        switch 32767 {
                                                        case 2147483647:
                                                            fmt.Println(v1)
                                                        case -1:
                                                            fmt.Println(v1)
                                                        default:
                                                            switch -878 {
                                                            case 127:
                                                                fmt.Println(v1)
                                                            }
                                                        }
                                                    } else {
                                                        for fi80 := 0; fi80 < 3; fi80++ {
                                                            if (2147483648 != 4294967295) {
                                                                v1 = (v1 | v1)
                                                                var v81 float64 = 1.5
                                                                _ = v81
                                                            } else {
                                                                if false {
                                                                    v3[("a" + "a")] = (29 > -128)
                                                                    var v82 uint64 = 48
                                                                    v1 -= (v1 ^ 1)
                                                                    _ = v82
                                                                }
                                                                for fi83 := 0; fi83 < 12; fi83++ {
                                                                    if (false || true) {
                                                                        v84 := v5[4294967295]
                                                                        if (!(506 <= 0)) {
                                                                            v85 := GenS6{F0: 0, F1: (fi83 % ((v84) | 1)), F2: -1}
                                                                            v86 := GenS8{F0: 1, F1: -947}
                                                                            var v87 string = ""
                                                                            _ = v85
                                                                            _ = v86
                                                                            _ = v87
                                                                        }
                                                                        _ = v84
                                                                    } else {
                                                                        var v88 uint16 = 65535
                                                                        var v89 bool = false
                                                                        if (-2 != -2147483648) {
                                                                            v89 = (418 >= 255)
                                                                        } else {
                                                                            fmt.Println(v88, fi80)
                                                                        }
                                                                        _ = v88
                                                                        _ = v89
                                                                    }
                                                                    switch fi83 {
                                                                    case 2147483647:
                                                                        var v90 int16 = (v1 + 1000)
                                                                        v91 := GenS8{F0: 548, F1: -1}
                                                                        _ = v90
                                                                        _ = v91
                                                                    case 128:
                                                                        var v92 int16 = v1
                                                                        _ = v92
                                                                    default:
                                                                        var v93 uint8 = 127
                                                                        _ = v93
                                                                    }
                                                                    if (genFn0("\n\t") > 9223372036854775807) {
                                                                        v94 := GenS9{F0: -48, F1: (v1 >> ((262) & 63)), F2: 1}
                                                                        switch -1 {
                                                                        case 32768:
                                                                            var v95 int8 = 42
                                                                            _ = v95
                                                                        default:
                                                                            if (65535 == 42) {
                                                                                v96 := GenS8{F0: 930, F1: -265}
                                                                                _ = v96
                                                                            }
                                                                        }
                                                                        fmt.Println(v1)
                                                                        _ = v94
                                                                    }
                                                                    _ = fi83
                                                                }
                                                                fmt.Println(fi80)
                                                            }
                                                            for fi97 := 0; fi97 < 4; fi97++ {
                                                                v1 += v1
                                                                _ = fi97
                                                            }
                                                            _ = fi80
                                                        }
                                                        var v98 int64 = -9223372036854775808
                                                        _ = v98
                                                    }
                                                default:
                                                    v1 += v1
                                                }
                                            case "foo bar":
                                                fmt.Println(v1)
                                            case "a":
                                                switch "foo bar" {
                                                case "":
                                                    switch 9223372036854775807 {
                                                    case 1:
                                                        switch "\n\t" {
                                                        case "\n\t":
                                                            v1 = v1
                                                        default:
                                                            switch 2147483647 {
                                                            case 127:
                                                                switch "\n\t" {
                                                                case "a":
                                                                    if (!(65535 < 65535)) {
                                                                        fmt.Println(v1)
                                                                        for fi99 := 0; fi99 < 2; fi99++ {
                                                                            if (2 == 522) {
                                                                                v1 -= 32767
                                                                                for fi100 := 0; fi100 < 9; fi100++ {
                                                                                    var v101 int64 = 265
                                                                                    var v102 float32 = 1e10
                                                                                    v101 += (v101 << ((1) & 63))
                                                                                    if true {
                                                                                        v103 := GenS7{F0: 523, F1: 42}
                                                                                        v1 = (v1 | 42)
                                                                                        v104 := GenS9{F0: 0, F1: (v1 >> ((2) & 63)), F2: -733}
                                                                                        _ = v103
                                                                                        _ = v104
                                                                                    }
                                                                                    if false {
                                                                                        if ((fi100 > 32768) && ("a" <= "foo bar")) {
                                                                                            var v105 float32 = -2.25
                                                                                            fmt.Println(v1)
                                                                                            var v106 float32 = 1.5
                                                                                            _ = v105
                                                                                            _ = v106
                                                                                        }
                                                                                    } else {
                                                                                        _ = len(v4)
                                                                                        v3["\n\t"] = (128 == (-(fi100)))
                                                                                        v101 -= 42
                                                                                    }
                                                                                    _ = fi100
                                                                                    _ = v101
                                                                                    _ = v102
                                                                                }
                                                                            }
                                                                            _ = fi99
                                                                        }
                                                                        if (-1 > 92) {
                                                                            v107 := GenS2{F0: 733}
                                                                            delete(v4, 197)
                                                                            _ = v107
                                                                        }
                                                                    }
                                                                default:
                                                                    if len(v2) >= 2 {
                                                                        v2 = v2[1:len(v2)]
                                                                    }
                                                                }
                                                            case 512:
                                                                v5[42] = -538
                                                            }
                                                        }
                                                    case 9223372036854775807:
                                                        fmt.Println(v1)
                                                    case 32767:
                                                        if true {
                                                            if len(v2) > 0 {
                                                                v2[0] = 929
                                                            }
                                                            v108 := GenS5{F0: false, F1: 1}
                                                            var v109 uint64 = 2
                                                            _ = v108
                                                            _ = v109
                                                        }
                                                    }
                                                case "hello":
                                                    var v110 uint16 = 2
                                                    _ = v110
                                                }
                                            }
                                            var v111 uint8 = 255
                                            _ = v111
                                        } else {
                                            var v112 float32 = (((float32(3.14159) / (0.1 + (float32(1.5) - ((1.5 * (((-2.25 * (((0.1 * 0.1) / 0.1) + float32(0.0))) - 3.14159) * 0.1)) + 1.5)))) + (float32(-2.25) - -2.25)) - (float32((float64((((0.1 + 3.14159) / (3.14159 * float32(float64(3.14159)))) * float32(float64(((1e10 - 0.1) / 1.5))))) + 1e10)) * -2.25))
                                            var v113 int32 = 1
                                            v112 += (0.1 * (v112 * v112))
                                            _ = v112
                                            _ = v113
                                        }
                                        v114 := GenS10{F0: 1}
                                        var v115 int = 32768
                                        _ = v114
                                        _ = v115
                                    }
                                } else {
                                    if false {
                                        _ = len(v5)
                                        v5[652] = 716
                                    }
                                    var v116 uint32 = 617
                                    _ = v116
                                }
                                var v117 bool = true
                                v1 = v1
                                _ = v117
                            }
                        }
                    }
                }
            }
            _ = v64
        } else {
            v1 = (-(v1))
            if false {
                if true {
                    for fi118 := 0; fi118 < 16; fi118++ {
                        var v119 int = fi118
                        _ = fi118
                        _ = v119
                    }
                    v120 := GenS4{F0: 0}
                    _ = v120
                } else {
                    var v121 uint32 = 65535
                    if false {
                        if false {
                            for fi122 := 0; fi122 < 7; fi122++ {
                                if false {
                                    for fi123 := 0; fi123 < 16; fi123++ {
                                        if true {
                                            v121 += 4294967295
                                        } else {
                                            var v124 uint = 1
                                            var v125 float32 = 0.1
                                            _ = v124
                                            _ = v125
                                        }
                                        _ = fi123
                                    }
                                    v121 = 628
                                    if false {
                                        fmt.Println(fi122, v121)
                                        v1 -= (v1 | v1)
                                    }
                                }
                                _ = fi122
                            }
                            switch "foo bar" {
                            case "hello":
                                switch 32767 {
                                case 127:
                                    v126 := GenS11{F0: 314, F1: 42, F2: 463}
                                    _ = v126
                                }
                            case "\n\t":
                                for fi127 := 0; fi127 < 16; fi127++ {
                                    switch "\n\t" {
                                    case "a":
                                        v121 -= 518
                                    case "hello":
                                        var v128 int8 = 1
                                        _ = v128
                                    }
                                    _ = fi127
                                }
                            case "a":
                                if true {
                                    var v129 int64 = 227
                                    for fi130 := 0; fi130 < 9; fi130++ {
                                        v131 := GenS12{F0: "foo bar", F1: 100, F2: 128}
                                        _ = fi130
                                        _ = v131
                                    }
                                    _ = v129
                                } else {
                                    _ = len(v2)
                                    for fi132 := 0; fi132 < 20; fi132++ {
                                        v133 := GenS12{F0: "foo bar", F1: 35, F2: 189}
                                        _ = fi132
                                        _ = v133
                                    }
                                }
                            }
                            var v134 int16 = 1
                            _ = v134
                        }
                        v135 := GenS7{F0: 803, F1: 633}
                        var v136 float32 = 1.5
                        _ = v135
                        _ = v136
                    } else {
                        if ((true || true) || (v121 == (v121 * 2147483648))) {
                            var v137 bool = (2 <= 2)
                            switch ("\n\t" + (("\n\t" + "\n\t") + ("" + ("\n\t" + "hello")))) {
                            case "":
                                v137 = true
                            case "hello":
                                var v138 uint8 = 0
                                _ = v138
                            }
                            _ = v137
                        }
                        fmt.Println(v1, v1)
                    }
                    var v139 float32 = 3.14159
                    _ = v121
                    _ = v139
                }
                v140 := v4[4294967295]
                _ = v140
            } else {
                if (!true) {
                    v141 := GenS8{F0: 700, F1: 42}
                    fmt.Println(v1)
                    v1 = v1
                    _ = v141
                } else {
                    if false {
                        fmt.Println(v1)
                    }
                    if (!false) {
                        var v142 int = 396
                        v1 -= 513
                        v143 := GenS13{F0: v142}
                        _ = v142
                        _ = v143
                    }
                    switch 0 {
                    case 4294967295:
                        v1 = (v1 << ((4294967295) & 63))
                    case -1:
                        v144 := GenS14{F0: 42}
                        _ = v144
                    }
                }
                fmt.Println(v1)
                var v145 uint = 806
                _ = v145
            }
            if (1.5 == 1.5) {
                if false {
                    switch 9223372036854775807 {
                    case 2147483647:
                        _ = len(v4)
                    case 223:
                        if (9223372036854775807 <= 0) {
                            var v146 int = -9223372036854775808
                            v1 = v1
                            if true {
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                switch -1 {
                                case 32767:
                                    _ = len(v3)
                                default:
                                    var v147 float64 = 3.14159
                                    _ = v147
                                }
                                v148 := v5[747]
                                _ = v148
                            } else {
                                var v149 uint = 2
                                var v150 bool = (!((42 == 1) || false))
                                v151 := GenS15{F0: 64, F1: 65535, F2: 65535}
                                _ = v149
                                _ = v150
                                _ = v151
                            }
                            _ = v146
                        } else {
                            if (v1 >= (v1 ^ v1)) {
                                v1 += (v1 * v1)
                                for fi152 := 0; fi152 < 1; fi152++ {
                                    var v153 int32 = -2147483648
                                    v154 := GenS9{F0: 833, F1: (v1 ^ (v1 << ((4294967295) & 63))), F2: (^(v153))}
                                    for fi155 := 0; fi155 < 18; fi155++ {
                                        v156 := GenS4{F0: (fi155 + (-(fi152)))}
                                        if false {
                                            switch fi152 {
                                            case 0:
                                                if ((fi155 | fi152) >= int(v153)) {
                                                    var v157 uint64 = 245
                                                    if false {
                                                        v153 += v153
                                                        if false {
                                                            v1 = (v1 & (-(v1)))
                                                            v158 := GenS16{F0: ("" + "a"), F1: (v157 - (v157 >> ((18446744073709551615) & 63)))}
                                                            _ = v158
                                                        }
                                                    }
                                                    _ = v157
                                                }
                                            default:
                                                v1 = 32767
                                            }
                                            fmt.Println(fi152)
                                            v1 = v1
                                            v159 := GenS17{F0: (true && true), F1: ("\n\t" + "\n\t")}
                                            _ = v159
                                        } else {
                                            var v160 uint64 = 18446744073709551615
                                            var v161 uint32 = 65535
                                            _ = v160
                                            _ = v161
                                        }
                                        _ = fi155
                                        _ = v156
                                    }
                                    _ = fi152
                                    _ = v153
                                    _ = v154
                                }
                                v1 = (v1 / (((v1 / ((483) | 1))) | 1))
                            }
                            v1 -= v1
                            v1 = (v1 / (((v1 * (-(v1)))) | 1))
                        }
                    case 32767:
                        var v162 uint32 = 1
                        _ = v162
                    default:
                        v163 := GenS18{F0: false, F1: -434}
                        _ = v163
                    }
                    v2 = append(v2, 65535)
                }
            }
        }
        switch 525 {
        case 833:
            if (false && false) {
                _ = len(v5)
                fmt.Println(v1)
                v2 = append(v2, 1)
            }
        default:
            fmt.Println(v1)
        }
        var v164 int64 = 1
        _ = v164
    }
    fmt.Println(v1)
    for fi165 := 0; fi165 < 18; fi165++ {
        if true {
            var v166 int8 = 127
            fmt.Println(v166, v1, v166)
            v166 += (v166 | (^(v166)))
            _ = v166
        }
        for fi167 := 0; fi167 < 13; fi167++ {
            var v168 uint32 = 2
            v1 = v1
            if len(v2) > 0 {
                v2[0] = 0
            }
            _ = fi167
            _ = v168
        }
        _ = fi165
    }
    for fi169 := 0; fi169 < 11; fi169++ {
        var v170 string = ""
        _ = fi169
        _ = v170
    }
    var v171 int8 = 1
    var v172 uint = 65535
    var v173 int16 = -1
    var v174 string = ("\n\t" + "hello")
    for fi175 := 0; fi175 < 18; fi175++ {
        var v176 uint16 = 501
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = fi175
        _ = v176
    }
    v171 -= v171
    var v177 string = (v174 + v174)
    v1 -= v1
    v178 := GenS18{F0: ((!false) || false), F1: int32(v171)}
    switch "hello" {
    case "foo bar":
        v179 := GenS15{F0: 0, F1: 42, F2: 599}
        _ = v179
    case "\n\t":
        fmt.Println(v177, v173, v1)
    }
    v171 = -51
    fmt.Println("v1:", v1)
    fmt.Println("v171:", v171)
    fmt.Println("v172:", v172)
    fmt.Println("v173:", v173)
    fmt.Println("v174:", v174)
    fmt.Println("v177:", v177)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v171
    _ = v172
    _ = v173
    _ = v174
    _ = v177
    _ = v178
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
