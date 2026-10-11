package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint; F1 uint8; F2 uint8 }
type GenS5 struct { F0 int8; F1 uint32; F2 uint32 }
type GenS6 struct { F0 uint64; F1 int16; F2 bool }
type GenS7 struct { F0 bool; F1 uint }
type GenS8 struct { F0 uint64 }
type GenS9 struct { F0 uint32 }
type GenS10 struct { F0 uint8 }
type GenS11 struct { F0 int64; F1 uint16 }
type GenS12 struct { F0 uint32; F1 int; F2 int }
type GenS13 struct { F0 uint; F1 uint16; F2 uint64 }
type GenS14 struct { F0 uint64 }
type GenS15 struct { F0 uint32; F1 int32; F2 uint }
type GenS16 struct { F0 string; F1 int }
type GenS17 struct { F0 int16; F1 int64; F2 int8 }
type GenS18 struct { F0 uint16; F1 int64 }
type GenS19 struct { F0 int16; F1 int8 }
func genFn0() bool {
    v6 := GenS1{F0: 0}
    var v7 uint8 = 0
    switch "foo bar" {
    case "foo bar":
        v8 := GenS2{F0: 2}
        _ = v8
    case "":
        if true {
            fmt.Println(v7)
            if true {
                if ((127 != 127) && true) {
                    v7 += v7
                    v7 -= 29
                } else {
                    fmt.Println(v7)
                    var v9 float64 = 0.0
                    _ = v9
                }
                fmt.Println(v7)
            }
            for fi10 := 0; fi10 < 7; fi10++ {
                v11 := GenS3{F0: 65535}
                v12 := []int{(^(fi10)), fi10, 42}
                if false {
                    for fi13 := 0; fi13 < 13; fi13++ {
                        if false {
                            fmt.Println(fi10)
                        } else {
                            switch "" {
                            case "foo bar":
                                fmt.Println(fi10)
                            case "\n\t":
                                switch "\n\t" {
                                case "\n\t":
                                    v14 := GenS2{F0: 32767}
                                    _ = v14
                                case "hello":
                                    v7 -= 225
                                default:
                                    fmt.Println(fi13, fi10, fi13)
                                }
                            case "a":
                                v15 := GenS2{F0: 851}
                                _ = v15
                            default:
                                if true {
                                    v16 := GenS2{F0: 81}
                                    v12 = append(v12, (fi13 / (((fi10 + (-(fi13)))) | 1)))
                                    fmt.Println(v7, v7, fi10)
                                    _ = v16
                                } else {
                                    v7 += v7
                                    _ = len(v12)
                                    fmt.Println(v7, fi10, v7)
                                    var v17 bool = (!true)
                                    var v18 int16 = 0
                                    _ = v17
                                    _ = v18
                                }
                            }
                            var v19 int8 = 1
                            var v20 uint64 = 66
                            _ = v19
                            _ = v20
                        }
                        v7 = 0
                        switch "" {
                        case "\n\t":
                            v7 = (v7 ^ 127)
                        case "foo bar":
                            switch fi13 {
                            case -934:
                                v7 += 0
                            }
                        }
                        _ = fi13
                    }
                    if true {
                        if (!true) {
                            v12 = append(v12, fi10)
                        } else {
                            if (-9223372036854775808 != 9223372036854775807) {
                                v7 = 128
                                v7 = 127
                                if false {
                                    switch (fi10 % ((fi10) | 1)) {
                                    case 1:
                                        var v21 int64 = -826
                                        _ = v21
                                    case -9223372036854775808:
                                        if (0 == -9223372036854775808) {
                                            if true {
                                                v22 := GenS4{F0: 65535, F1: (v7 / (((v7 / ((v7) | 1))) | 1)), F2: (v7 & 2)}
                                                v7 += v7
                                                _ = len(v12)
                                                _ = v22
                                            }
                                            v7 = v7
                                        }
                                    case -450:
                                        if len(v12) >= 2 {
                                            v12 = v12[1:len(v12)]
                                        }
                                    }
                                    for fi23 := 0; fi23 < 17; fi23++ {
                                        var v24 int16 = -32768
                                        _ = fi23
                                        _ = v24
                                    }
                                    fmt.Println(fi10, fi10)
                                }
                            }
                        }
                        if false {
                            if true {
                                v7 -= v7
                            } else {
                                switch "" {
                                case "\n\t":
                                    if false {
                                        for fi25 := 0; fi25 < 18; fi25++ {
                                            if len(v12) > 0 {
                                                v12[0] = fi25
                                            }
                                            _ = fi25
                                        }
                                    }
                                case "a":
                                    v7 = (v7 + v7)
                                case "":
                                    for fi26 := 0; fi26 < 9; fi26++ {
                                        var v27 float32 = (3.14159 / 0.1)
                                        _ = fi26
                                        _ = v27
                                    }
                                }
                            }
                            if false {
                                switch fi10 {
                                case 32767:
                                    for fi28 := 0; fi28 < 10; fi28++ {
                                        switch fi10 {
                                        case 4294967295:
                                            v7 += (v7 | (v7 | v7))
                                        case 0:
                                            var v29 uint8 = (v7 / ((v7) | 1))
                                            _ = v29
                                        case 815:
                                            if len(v12) > 0 {
                                                v12[0] = fi10
                                            }
                                            v30 := GenS1{F0: 656}
                                            fmt.Println(v7)
                                            _ = v30
                                        default:
                                            v7 = (v7 >> ((328) & 63))
                                        }
                                        if true {
                                            v31 := GenS5{F0: -1, F1: 1, F2: 994}
                                            var v32 int64 = 1
                                            _ = v31
                                            _ = v32
                                        }
                                        v7 += v7
                                        _ = fi28
                                    }
                                case -329:
                                    var v33 bool = true
                                    _ = v33
                                case 217:
                                    fmt.Println(v7, v7)
                                }
                                if true {
                                    _ = len(v12)
                                    for fi34 := 0; fi34 < 18; fi34++ {
                                        _ = len(v12)
                                        var v35 string = ("foo bar" + "foo bar")
                                        if (1e10 >= 1e10) {
                                            var v36 uint32 = 440
                                            v7 += v7
                                            var v37 string = v35
                                            _ = v36
                                            _ = v37
                                        } else {
                                            var v38 bool = false
                                            _ = v38
                                        }
                                        _ = fi34
                                        _ = v35
                                    }
                                    fmt.Println(fi10)
                                }
                            } else {
                                if (!(42 <= 2147483647)) {
                                    var v39 int16 = -1
                                    var v40 uint64 = 473
                                    for fi41 := 0; fi41 < 12; fi41++ {
                                        v7 = v7
                                        _ = fi41
                                    }
                                    _ = v39
                                    _ = v40
                                } else {
                                    if false {
                                        v42 := GenS6{F0: 533, F1: 42, F2: (true && false)}
                                        v7 = v7
                                        _ = v42
                                    }
                                    v43 := GenS2{F0: 65535}
                                    _ = v43
                                }
                            }
                        }
                    } else {
                        if len(v12) > 0 {
                            v12[0] = (fi10 & (fi10 % (((^(fi10))) | 1)))
                        }
                    }
                }
                _ = fi10
                _ = v11
                _ = v12
            }
        }
    case "hello":
        v7 = (v7 / ((v7) | 1))
    default:
        var v44 int = int(v7)
        _ = v44
    }
    fmt.Println(v7)
    _ = v6
    _ = v7
    return ((v7 & v7) == 1)
}

func main() {
    v1 := []string{("" + ("foo bar" + "a"))}
    v2 := []int{-343, -9223372036854775808, 32767, 32767}
    v3 := []string{"foo bar", ("\n\t" + ((("" + "hello") + "foo bar") + ""))}
    var v4 int64 = 42
    v5 := []int{-615, 4294967295, 96, 1}
    switch -1 {
    case -360:
        v4 = v4
    case -1:
        var v45 int32 = -1
        _ = v45
    }
    if ((466 != -32768) || true) {
        switch "a" {
        case "":
            var v46 int64 = v4
            _ = v46
        case "foo bar":
            var v47 uint16 = 0
            _ = v47
        case "\n\t":
            v4 = v4
        default:
            v4 -= (v4 - 1)
        }
        if (98 != 100) {
            for fi48 := 0; fi48 < 7; fi48++ {
                for fi49 := 0; fi49 < 10; fi49++ {
                    if true {
                        var v50 int64 = v4
                        fmt.Println(fi49)
                        v3 = append(v3, ("hello" + "\n\t"))
                        _ = v50
                    }
                    v4 -= (v4 ^ v4)
                    _ = fi49
                }
                _ = fi48
            }
            fmt.Println(v4)
            if len(v5) > 0 {
                v5[0] = 2147483647
            }
        }
    }
    if (false && false) {
        v5 = append(v5, 4294967295)
        v1 = append(v1, ("\n\t" + "hello"))
    } else {
        fmt.Println(v4)
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    if (false || true) {
        fmt.Println(v4)
        for fi51 := 0; fi51 < 5; fi51++ {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            switch "" {
            case "":
                var v52 uint32 = 42
                _ = v52
            case "hello":
                var v53 uint64 = 42
                _ = v53
            }
            _ = len(v2)
            _ = fi51
        }
        v54 := GenS7{F0: true, F1: 783}
        _ = v54
    }
    for fi55 := 0; fi55 < 6; fi55++ {
        for fi56 := 0; fi56 < 7; fi56++ {
            v4 = v4
            if false {
                if (-853 == 1) {
                    v4 = 9223372036854775807
                } else {
                    switch 128 {
                    case 4294967295:
                        if true {
                            switch "hello" {
                            case "hello":
                                v57 := GenS2{F0: 631}
                                _ = v57
                            case "a":
                                if true {
                                    if (v4 == (v4 - (v4 >> ((42) & 63)))) {
                                        v58 := GenS8{F0: 2}
                                        _ = v58
                                    }
                                    v2 = append(v2, fi56)
                                    if true {
                                        var v59 uint32 = 787
                                        var v60 int8 = -1
                                        if true {
                                            switch (("a" + "") + (("foo bar" + "foo bar") + "a")) {
                                            case "a":
                                                var v61 uint64 = 0
                                                _ = v61
                                            default:
                                                if true {
                                                    if false {
                                                        var v62 string = "\n\t"
                                                        v60 = int8(fi56)
                                                        _ = v62
                                                    }
                                                }
                                            }
                                            v2 = append(v2, -345)
                                        } else {
                                            v59 -= v59
                                            if (!((!(-2147483648 >= 389)) && ((v60 >= int8(v4)) || (!false)))) {
                                                _ = len(v3)
                                                var v63 string = ("\n\t" + "hello")
                                                _ = v63
                                            }
                                            var v64 int = (fi56 & fi56)
                                            _ = v64
                                        }
                                        _ = v59
                                        _ = v60
                                    } else {
                                        var v65 float32 = (3.14159 - (float32(0.0) + 1.5))
                                        v66 := GenS9{F0: 2147483647}
                                        _ = v65
                                        _ = v66
                                    }
                                } else {
                                    fmt.Println(v4, fi55)
                                    var v67 uint = 912
                                    _ = v67
                                }
                            case "foo bar":
                                v2 = append(v2, (fi56 + (fi55 * fi56)))
                            }
                        }
                    default:
                        v68 := GenS10{F0: 128}
                        fmt.Println(fi56, fi56)
                        v4 -= (v4 % (((v4 ^ v4)) | 1))
                        _ = v68
                    }
                    var v69 float32 = 0.0
                    fmt.Println(v69, v4, fi56)
                    _ = v69
                }
                v70 := GenS1{F0: 371}
                _ = v70
            } else {
                var v71 uint32 = 4294967295
                switch fi56 {
                case 0:
                    if true {
                        if (!(-2.25 >= (float32(1e10) + (float32(0.1) - 3.14159)))) {
                            _ = len(v5)
                            var v72 string = "\n\t"
                            v73 := GenS1{F0: 660}
                            _ = v72
                            _ = v73
                        } else {
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                        }
                    } else {
                        if false {
                            if ((0 >= 18446744073709551615) || true) {
                                v71 = v71
                            } else {
                                var v74 uint16 = 42
                                _ = v74
                            }
                            var v75 int = 42
                            v71 = (v71 ^ (v71 + v71))
                            _ = v75
                        }
                        var v76 uint32 = (v71 % (((^(v71))) | 1))
                        _ = v76
                    }
                case -9223372036854775808:
                    fmt.Println(fi55, fi55, fi56)
                case 416:
                    fmt.Println(v4, v4)
                    switch fi56 {
                    case -2147483648:
                        if len(v5) > 0 {
                            v5[0] = (fi56 << ((2) & 63))
                        }
                    case 32768:
                        v77 := GenS11{F0: (v4 | v4), F1: 32767}
                        _ = v77
                    }
                default:
                    var v78 int16 = 32767
                    _ = v78
                }
                var v79 int = -2147483648
                _ = v71
                _ = v79
            }
            switch fi56 {
            case 9223372036854775807:
                fmt.Println(fi56, fi55, fi56)
                var v80 bool = (!false)
                if (!v80) {
                    switch (fi56 >> ((284) & 63)) {
                    case 9223372036854775807:
                        var v81 bool = ((0.0 * float64(3.14159)) >= -2.25)
                        _ = v81
                    default:
                        if true {
                            if true {
                                v82 := GenS12{F0: 0, F1: fi56, F2: fi56}
                                _ = v82
                            } else {
                                var v83 float32 = (((1.5 * (3.14159 * 3.14159)) / ((0.1 - 0.0) / -2.25)) - float32((0.0 * 1.5)))
                                _ = v83
                            }
                            v80 = (!(!false))
                        } else {
                            var v84 float32 = float32(1.5)
                            _ = v84
                        }
                    }
                    v4 = (v4 / ((-460) | 1))
                    switch -187 {
                    case 42:
                        v85 := GenS3{F0: 65535}
                        _ = v85
                    }
                    var v86 uint16 = 32767
                    _ = v86
                }
                _ = v80
            case 127:
                v87 := GenS13{F0: 4294967295, F1: 42, F2: 18446744073709551615}
                _ = v87
            }
            v4 = v4
            if len(v2) > 0 {
                v2[0] = 32768
            }
            var v88 uint64 = 0
            _ = fi56
            _ = v88
        }
        fmt.Println(v4, fi55)
        _ = fi55
    }
    var v89 int64 = v4
    for fi90 := 0; fi90 < 12; fi90++ {
        var v91 int64 = (v89 / (((v4 / (((-(v89))) | 1))) | 1))
        fmt.Println(v4)
        _ = fi90
        _ = v91
    }
    var v92 float64 = float64(((0.0 * (((0.0 - 0.1) - (3.14159 / (1.5 + ((-2.25 - -2.25) * (-2.25 - 1.5))))) / 1e10)) * (0.1 / ((0.1 * (0.0 - ((1e10 * 1e10) - (float32(1.5) - ((3.14159 / 1.5) * float32(1e10)))))) + 0.0))))
    v5 = append(v5, 0)
    for fi93 := 0; fi93 < 11; fi93++ {
        fmt.Println(v89)
        var v94 uint16 = 470
        _ = fi93
        _ = v94
    }
    var v95 uint8 = 255
    v89 -= (^(v4))
    for fi96 := 0; fi96 < 19; fi96++ {
        v95 = (v95 - (^(v95)))
        for fi97 := 0; fi97 < 8; fi97++ {
            if len(v2) > 0 {
                v2[0] = fi97
            }
            _ = fi97
        }
        if ((42 < 42) || true) {
            for fi98 := 0; fi98 < 19; fi98++ {
                if true {
                    var v99 int = 127
                    switch "" {
                    case "hello":
                        var v100 string = "hello"
                        _ = v100
                    case "":
                        if (42 != 0) {
                            v92 = (v92 / (v92 + 1e10))
                            v3 = append(v3, "hello")
                            fmt.Println(v4, fi96, v92)
                            var v101 uint8 = (v95 << ((18446744073709551615) & 63))
                            _ = v101
                        } else {
                            v102 := GenS11{F0: v4, F1: 721}
                            _ = v102
                        }
                    case "\n\t":
                        v89 += (-(v89))
                    }
                    var v103 uint64 = 4294967295
                    v89 = 1
                    if (((v95 + (v95 << ((18446744073709551615) & 63))) <= uint8(v103)) && true) {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        var v104 uint = 65535
                        var v105 string = ""
                        _ = v104
                        _ = v105
                    } else {
                        v4 = v89
                        var v106 uint64 = (v103 | v103)
                        _ = v106
                    }
                    _ = v103
                    _ = v99
                } else {
                    switch 4294967295 {
                    case 1:
                        var v107 bool = (!true)
                        _ = v107
                    case 959:
                        if len(v1) > 0 {
                            v1[0] = "foo bar"
                        }
                    case 127:
                        var v108 int64 = (v4 * v4)
                        _ = v108
                    }
                    var v109 string = ("foo bar" + "a")
                    var v110 int = 0
                    _ = v109
                    _ = v110
                }
                _ = fi98
            }
            v111 := GenS11{F0: v4, F1: 255}
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = v111
        }
        _ = fi96
    }
    if ((255 >= 1) || (42 < 1)) {
        v92 -= 0.0
        var v112 int64 = 1
        switch -2 {
        case 32767:
            var v113 int = -525
            _ = v113
        case -418:
            if true {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                if false {
                    fmt.Println(v4)
                    switch -522 {
                    case 2147483647:
                        if ((v4 << ((990) & 63)) > 0) {
                            if false {
                                v112 = (v112 << ((65535) & 63))
                                fmt.Println(v89, v95, v112)
                            }
                            var v114 int64 = (v112 + (v89 / ((v89) | 1)))
                            if (1 >= 65535) {
                                if true {
                                    if true {
                                        fmt.Println(v114, v92)
                                    }
                                }
                                if ((v95 % (((v95 * 127)) | 1)) <= (v95 & (v95 % (((^(v95))) | 1)))) {
                                    _ = len(v2)
                                    if (!false) {
                                        fmt.Println(v112, v89, v4)
                                    } else {
                                        v92 += (((v92 * 1e10) + float64(1e10)) * v92)
                                        v115 := GenS7{F0: (!((0 >= 255) || false)), F1: 18446744073709551615}
                                        _ = v115
                                    }
                                    for fi116 := 0; fi116 < 16; fi116++ {
                                        _ = len(v2)
                                        var v117 int32 = 42
                                        switch (-(fi116)) {
                                        case -1:
                                            switch "" {
                                            case "hello":
                                                if false {
                                                    for fi118 := 0; fi118 < 16; fi118++ {
                                                        v114 = v112
                                                        switch int(v95) {
                                                        case 2147483647:
                                                            fmt.Println(v112, fi118, fi118)
                                                        default:
                                                            v112 -= v4
                                                        }
                                                        _ = fi118
                                                    }
                                                }
                                            case "":
                                                var v119 uint64 = 0
                                                _ = v119
                                            case "\n\t":
                                                for fi120 := 0; fi120 < 1; fi120++ {
                                                    if false {
                                                        fmt.Println(v114)
                                                    } else {
                                                        var v121 int8 = 1
                                                        fmt.Println(v117, v89)
                                                        switch ("a" + "hello") {
                                                        case "a":
                                                            v121 += v121
                                                        }
                                                        var v122 uint = 0
                                                        switch "a" {
                                                        case "foo bar":
                                                            if (4294967295 > v122) {
                                                                v121 = (^(v121))
                                                                switch "\n\t" {
                                                                case "\n\t":
                                                                    var v123 int8 = v121
                                                                    _ = v123
                                                                case "foo bar":
                                                                    v95 += 255
                                                                case "hello":
                                                                    var v124 int = 127
                                                                    _ = v124
                                                                default:
                                                                    var v125 float32 = (-2.25 - -2.25)
                                                                    _ = v125
                                                                }
                                                            } else {
                                                                var v126 uint32 = 88
                                                                v127 := GenS1{F0: 42}
                                                                _ = v126
                                                                _ = v127
                                                            }
                                                        case "":
                                                            var v128 int32 = v117
                                                            _ = v128
                                                        }
                                                        _ = v121
                                                        _ = v122
                                                    }
                                                    fmt.Println(fi120, v89, v89)
                                                    _ = fi120
                                                }
                                            }
                                        case 366:
                                            v114 -= v114
                                        case -9223372036854775808:
                                            fmt.Println(v114)
                                        }
                                        _ = fi116
                                        _ = v117
                                    }
                                }
                                v129 := GenS6{F0: 42, F1: 376, F2: ((v112 < v112) || false)}
                                _ = v129
                            } else {
                                for fi130 := 0; fi130 < 10; fi130++ {
                                    switch "a" {
                                    case "foo bar":
                                        var v131 string = "foo bar"
                                        _ = v131
                                    case "\n\t":
                                        fmt.Println(v112, v114)
                                    case "hello":
                                        v112 = v112
                                    }
                                    _ = fi130
                                }
                                v89 = (v89 % (((v89 - v4)) | 1))
                            }
                            _ = v114
                        } else {
                            if (true && true) {
                                v132 := GenS14{F0: 4294967295}
                                switch "hello" {
                                case "foo bar":
                                    var v133 string = ""
                                    _ = v133
                                case "a":
                                    v3 = append(v3, "hello")
                                case "hello":
                                    _ = len(v2)
                                }
                                _ = v132
                            } else {
                                fmt.Println(v92, v4)
                                for fi134 := 0; fi134 < 4; fi134++ {
                                    for fi135 := 0; fi135 < 13; fi135++ {
                                        switch 32767 {
                                        case 128:
                                            if (2147483647 != -2147483648) {
                                                fmt.Println(v92)
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                            }
                                            v5 = append(v5, fi134)
                                        }
                                        _ = fi135
                                    }
                                    _ = fi134
                                }
                            }
                            v112 -= (v89 >> ((65535) & 63))
                            var v136 uint8 = (v95 % ((v95) | 1))
                            _ = v136
                        }
                    case 4294967295:
                        if false {
                            if (!true) {
                                switch "" {
                                case "\n\t":
                                    v137 := GenS3{F0: 0}
                                    _ = v137
                                case "a":
                                    switch "foo bar" {
                                    case "foo bar":
                                        if true {
                                            v2 = append(v2, 324)
                                            v89 -= -1
                                            var v138 int8 = -54
                                            _ = v138
                                        } else {
                                            if (!(false && ((-1 >= -2147483648) && (false && true)))) {
                                                switch 32767 {
                                                case 42:
                                                    var v139 int16 = 42
                                                    _ = v139
                                                case 0:
                                                    if ("" <= "") {
                                                        var v140 int = -1
                                                        if len(v3) >= 2 {
                                                            v3 = v3[1:len(v3)]
                                                        }
                                                        fmt.Println(v112, v140)
                                                        _ = v140
                                                    }
                                                default:
                                                    for fi141 := 0; fi141 < 7; fi141++ {
                                                        if true {
                                                            fmt.Println(fi141, v112)
                                                            for fi142 := 0; fi142 < 11; fi142++ {
                                                                _ = len(v1)
                                                                _ = fi142
                                                            }
                                                        }
                                                        var v143 string = ""
                                                        fmt.Println(v92, v95)
                                                        _ = fi141
                                                        _ = v143
                                                    }
                                                }
                                                v92 = float64(float32(v92))
                                            } else {
                                                v89 -= (v89 / (((v112 * v112)) | 1))
                                            }
                                            if ("" <= "") {
                                                var v144 string = "foo bar"
                                                for fi145 := 0; fi145 < 6; fi145++ {
                                                    var v146 int16 = int16(fi145)
                                                    if true {
                                                        v89 += (^(v89))
                                                    } else {
                                                        var v147 string = "a"
                                                        var v148 uint16 = 255
                                                        _ = v147
                                                        _ = v148
                                                    }
                                                    v112 = (v112 | v4)
                                                    _ = fi145
                                                    _ = v146
                                                }
                                                _ = v144
                                            } else {
                                                v149 := GenS8{F0: 18446744073709551615}
                                                v150 := GenS7{F0: (!false), F1: 65535}
                                                v151 := GenS4{F0: 255, F1: v95, F2: 255}
                                                _ = v149
                                                _ = v150
                                                _ = v151
                                            }
                                            if false {
                                                if len(v2) > 0 {
                                                    v2[0] = 42
                                                }
                                                v4 = (-(v112))
                                                for fi152 := 0; fi152 < 5; fi152++ {
                                                    v153 := GenS12{F0: 1, F1: fi152, F2: (fi152 * (fi152 * (fi152 | 0)))}
                                                    v92 = v92
                                                    v154 := GenS12{F0: 65535, F1: fi152, F2: -522}
                                                    _ = fi152
                                                    _ = v153
                                                    _ = v154
                                                }
                                            }
                                        }
                                    case "":
                                        if (false || (!false)) {
                                            if (float64(0.0) <= (0.0 + 0.0)) {
                                                v89 += v112
                                            }
                                            v92 -= v92
                                        }
                                    }
                                default:
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                }
                                if (false && false) {
                                    switch 32768 {
                                    case 16:
                                        v155 := GenS2{F0: 0}
                                        _ = v155
                                    case -9223372036854775808:
                                        if (127 > 2) {
                                            if true {
                                                for fi156 := 0; fi156 < 1; fi156++ {
                                                    var v157 uint16 = 65535
                                                    switch (^(fi156)) {
                                                    case 0:
                                                        if (0.1 >= (1.5 / (v92 + (-2.25 + (v92 / v92))))) {
                                                            switch "foo bar" {
                                                            case "a":
                                                                switch ("\n\t" + "") {
                                                                case "foo bar":
                                                                    v89 = (v112 - v112)
                                                                default:
                                                                    var v158 int8 = 100
                                                                    _ = v158
                                                                }
                                                            case "foo bar":
                                                                if true {
                                                                    fmt.Println(v157)
                                                                }
                                                            case "\n\t":
                                                                var v159 int16 = 42
                                                                _ = v159
                                                            }
                                                            if len(v1) >= 2 {
                                                                v1 = v1[1:len(v1)]
                                                            }
                                                            for fi160 := 0; fi160 < 10; fi160++ {
                                                                var v161 int64 = -495
                                                                switch fi156 {
                                                                case 4294967295:
                                                                    var v162 uint16 = v157
                                                                    _ = v162
                                                                case 1:
                                                                    var v163 float64 = -2.25
                                                                    _ = v163
                                                                }
                                                                var v164 int32 = -1
                                                                _ = fi160
                                                                _ = v161
                                                                _ = v164
                                                            }
                                                        }
                                                    case -1:
                                                        var v165 int64 = (^(v4))
                                                        _ = v165
                                                    case 32768:
                                                        v166 := GenS6{F0: 42, F1: 230, F2: (v95 > (^(v95)))}
                                                        _ = v166
                                                    }
                                                    v167 := GenS6{F0: 18446744073709551615, F1: 42, F2: (!((65535 != uint(v157)) && false))}
                                                    _ = fi156
                                                    _ = v157
                                                    _ = v167
                                                }
                                            }
                                            if false {
                                                for fi168 := 0; fi168 < 13; fi168++ {
                                                    v169 := GenS12{F0: 2147483648, F1: fi168, F2: (fi168 * fi168)}
                                                    for fi170 := 0; fi170 < 7; fi170++ {
                                                        var v171 int16 = 32767
                                                        var v172 int8 = 49
                                                        v173 := GenS15{F0: 2, F1: 2147483647, F2: 692}
                                                        _ = fi170
                                                        _ = v171
                                                        _ = v172
                                                        _ = v173
                                                    }
                                                    var v174 uint32 = 391
                                                    _ = fi168
                                                    _ = v169
                                                    _ = v174
                                                }
                                            } else {
                                                var v175 int64 = (v4 << ((65535) & 63))
                                                var v176 string = "hello"
                                                v177 := GenS2{F0: 789}
                                                _ = v175
                                                _ = v176
                                                _ = v177
                                            }
                                        } else {
                                            if len(v5) >= 2 {
                                                v5 = v5[1:len(v5)]
                                            }
                                        }
                                    }
                                } else {
                                    _ = len(v3)
                                }
                                switch (("" + ("" + ("" + ((("\n\t" + ("" + "hello")) + (("hello" + "foo bar") + ((("" + ("\n\t" + "foo bar")) + "hello") + (("hello" + "hello") + ("hello" + ((("foo bar" + "foo bar") + (((("foo bar" + "foo bar") + "a") + "\n\t") + ("" + ((((("hello" + "a") + "a") + ("foo bar" + "hello")) + "a") + ("foo bar" + (("hello" + (("\n\t" + (("foo bar" + (("\n\t" + ("hello" + "a")) + "\n\t")) + "foo bar")) + "\n\t")) + "")))))) + ("\n\t" + ""))))))) + "hello")))) + "\n\t") {
                                case "hello":
                                    if (!(1 != -1)) {
                                        switch 32767 {
                                        case 139:
                                            if true {
                                                v1 = append(v1, "\n\t")
                                            }
                                        case 2147483647:
                                            v178 := GenS12{F0: 0, F1: -9223372036854775808, F2: 128}
                                            _ = v178
                                        case 35:
                                            v89 = (v4 % (((v4 / ((42) | 1))) | 1))
                                        }
                                        for fi179 := 0; fi179 < 7; fi179++ {
                                            var v180 bool = false
                                            v95 = v95
                                            v181 := GenS7{F0: true, F1: 4294967295}
                                            _ = fi179
                                            _ = v180
                                            _ = v181
                                        }
                                        for fi182 := 0; fi182 < 9; fi182++ {
                                            switch "hello" {
                                            case "a":
                                                var v183 int64 = v89
                                                _ = v183
                                            case "\n\t":
                                                if false {
                                                    var v184 int = fi182
                                                    v89 = -9223372036854775808
                                                    v185 := GenS11{F0: v112, F1: 32767}
                                                    _ = v184
                                                    _ = v185
                                                }
                                            case "foo bar":
                                                if true {
                                                    if true {
                                                        v89 += v112
                                                        v89 += (v112 % (((v4 + (v112 + v112))) | 1))
                                                    }
                                                    v186 := GenS9{F0: 4294967295}
                                                    _ = v186
                                                }
                                            default:
                                                var v187 int8 = -85
                                                _ = v187
                                            }
                                            _ = fi182
                                        }
                                    }
                                case "a":
                                    if true {
                                        fmt.Println(v92, v112, v92)
                                        var v188 int8 = 1
                                        v112 -= (v112 - v4)
                                        _ = v188
                                    }
                                }
                            } else {
                                if (false && ("hello" == "a")) {
                                    for fi189 := 0; fi189 < 12; fi189++ {
                                        v190 := GenS10{F0: v95}
                                        for fi191 := 0; fi191 < 11; fi191++ {
                                            _ = len(v5)
                                            var v192 uint = 42
                                            v89 = v4
                                            _ = fi191
                                            _ = v192
                                        }
                                        _ = fi189
                                        _ = v190
                                    }
                                    for fi193 := 0; fi193 < 17; fi193++ {
                                        for fi194 := 0; fi194 < 11; fi194++ {
                                            if false {
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                            } else {
                                                v195 := GenS16{F0: "", F1: (fi193 % (((fi193 >> ((65535) & 63))) | 1))}
                                                _ = v195
                                            }
                                            v196 := GenS4{F0: 42, F1: (v95 & v95), F2: 9}
                                            if len(v2) > 0 {
                                                v2[0] = fi193
                                            }
                                            _ = fi194
                                            _ = v196
                                        }
                                        switch (fi193 | 127) {
                                        case 128:
                                            if (2147483647 > -2147483648) {
                                                v197 := GenS17{F0: 1000, F1: 1, F2: -1}
                                                _ = v197
                                            } else {
                                                v4 = 0
                                                if false {
                                                    fmt.Println(fi193, v95)
                                                }
                                                if (!false) {
                                                    for fi198 := 0; fi198 < 6; fi198++ {
                                                        if (40 == 127) {
                                                            v89 += 315
                                                            if len(v2) > 0 {
                                                                v2[0] = 128
                                                            }
                                                            var v199 string = ""
                                                            _ = v199
                                                        }
                                                        _ = fi198
                                                    }
                                                    v200 := GenS13{F0: 870, F1: 1, F2: 4294967295}
                                                    _ = v200
                                                }
                                            }
                                        case 1:
                                            if ((-199 >= fi193) && (-158 > 777)) {
                                                if false {
                                                    v201 := GenS1{F0: 472}
                                                    var v202 uint16 = 0
                                                    if (v92 < 1e10) {
                                                        fmt.Println(v92, v4)
                                                    } else {
                                                        v203 := GenS15{F0: 592, F1: 42, F2: 0}
                                                        var v204 float32 = (0.1 / 3.14159)
                                                        _ = v203
                                                        _ = v204
                                                    }
                                                    _ = v201
                                                    _ = v202
                                                }
                                                v205 := GenS12{F0: 4294967295, F1: fi193, F2: fi193}
                                                _ = v205
                                            } else {
                                                if (427 != 0) {
                                                    var v206 float64 = (0.1 * (((v92 - (1e10 + (v92 * 0.0))) / 1e10) / v92))
                                                    v206 = (0.1 + ((v206 * ((v206 * v92) / v206)) * 0.0))
                                                    _ = v206
                                                }
                                                switch "" {
                                                case "":
                                                    _ = len(v1)
                                                case "foo bar":
                                                    switch 275 {
                                                    case -2147483648:
                                                        v89 = (v112 * v112)
                                                    case 4294967295:
                                                        switch "\n\t" {
                                                        case "a":
                                                            var v207 float64 = v92
                                                            _ = v207
                                                        }
                                                    case 300:
                                                        _ = len(v5)
                                                    default:
                                                        var v208 uint = 759
                                                        _ = v208
                                                    }
                                                case "hello":
                                                    switch ("hello" + ("foo bar" + "hello")) {
                                                    case "":
                                                        var v209 uint32 = 2147483648
                                                        _ = v209
                                                    }
                                                default:
                                                    var v210 uint = 4294967295
                                                    _ = v210
                                                }
                                                var v211 uint = 0
                                                _ = v211
                                            }
                                        case -1:
                                            var v212 float64 = 3.14159
                                            _ = v212
                                        default:
                                            v112 -= (v4 >> ((118) & 63))
                                        }
                                        _ = len(v3)
                                        _ = fi193
                                    }
                                } else {
                                    v213 := GenS1{F0: 42}
                                    var v214 uint32 = 42
                                    fmt.Println(v95, v4, v92)
                                    _ = v213
                                    _ = v214
                                }
                            }
                            _ = len(v1)
                            v4 -= 1
                        } else {
                            for fi215 := 0; fi215 < 6; fi215++ {
                                for fi216 := 0; fi216 < 6; fi216++ {
                                    v95 += 84
                                    fmt.Println(v95, v92, fi215)
                                    v4 = (v112 + -1)
                                    var v217 uint32 = 0
                                    _ = fi216
                                    _ = v217
                                }
                                _ = fi215
                            }
                            var v218 string = ("" + ("" + ""))
                            v4 += 9223372036854775807
                            _ = v218
                        }
                    case 0:
                        v89 = (-(v4))
                    default:
                        var v219 int = 355
                        _ = v219
                    }
                } else {
                    v4 = (v112 / (((v112 + -1)) | 1))
                    switch 42 {
                    case 9223372036854775807:
                        var v220 uint = 255
                        _ = v220
                    default:
                        v1 = append(v1, "foo bar")
                    }
                }
                var v221 uint16 = 1
                _ = v221
            } else {
                var v222 uint = 0
                _ = v222
            }
        default:
            if true {
                v223 := GenS18{F0: 1, F1: (^(v112))}
                v112 += (v112 % (((-(v112))) | 1))
                _ = v223
            }
        }
        _ = v112
    } else {
        v224 := GenS19{F0: -600, F1: 42}
        _ = v224
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v89:", v89)
    fmt.Println("v92:", v92)
    fmt.Println("v95:", v95)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v89
    _ = v92
    _ = v95
}
