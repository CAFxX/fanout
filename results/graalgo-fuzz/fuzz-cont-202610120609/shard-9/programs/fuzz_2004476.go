package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int8; F1 int32; F2 uint64 }
type GenS3 struct { F0 int16; F1 int16; F2 uint64 }
type GenS4 struct { F0 string; F1 uint; F2 uint }
type GenS5 struct { F0 string; F1 int; F2 uint16 }
type GenS6 struct { F0 uint; F1 uint32; F2 bool }
type GenS7 struct { F0 int8; F1 string }
type GenS8 struct { F0 int16 }
type GenS9 struct { F0 string; F1 int }
type GenS10 struct { F0 bool }
type GenS11 struct { F0 uint64 }
type GenS12 struct { F0 int; F1 uint64 }
type GenS13 struct { F0 uint64; F1 int16 }
type GenS14 struct { F0 string; F1 int32 }
type GenS15 struct { F0 uint32 }
type GenS16 struct { F0 string; F1 uint64; F2 uint8 }
type GenS17 struct { F0 int; F1 int }
type GenS18 struct { F0 uint16; F1 uint32 }
type GenS19 struct { F0 uint; F1 uint64 }
type GenS20 struct { F0 int64; F1 int16; F2 uint }
func genFn0(p0 uint16, p1 uint32) uint {
    for fi7 := 0; fi7 < 8; fi7++ {
        for fi8 := 0; fi8 < 7; fi8++ {
            p1 -= 486
            switch 32767 {
            case -9223372036854775808:
                var v9 uint16 = p0
                _ = v9
            case -1:
                var v10 int8 = 100
                _ = v10
            default:
                switch "" {
                case "":
                    var v11 int = fi7
                    p1 = 2147483647
                    _ = v11
                case "foo bar":
                    p1 += 1
                    v12 := GenS1{F0: 2}
                    if false {
                        v13 := []int{}
                        if len(v13) > 0 {
                            v13[0] = 0
                        }
                        _ = v13
                    }
                    _ = v12
                }
            }
            if true {
                v14 := GenS2{F0: 42, F1: 42, F2: 4294967295}
                if ((1.5 * (0.1 + 0.1)) >= 0.1) {
                    p0 = (p0 - p0)
                    v15 := GenS3{F0: 42, F1: -777, F2: 0}
                    v16 := []int{(fi8 % (((fi7 << ((954) & 63))) | 1))}
                    v17 := GenS4{F0: (("foo bar" + "\n\t") + "\n\t"), F1: uint(p1), F2: 65535}
                    var v18 uint16 = (p0 % ((411) | 1))
                    _ = v15
                    _ = v16
                    _ = v17
                    _ = v18
                }
                _ = v14
            }
            _ = fi8
        }
        _ = fi7
    }
    p1 += uint32(p0)
    fmt.Println(p0)
    var v19 uint32 = p1
    var v20 uint8 = 0
    _ = p0
    _ = p1
    _ = v19
    _ = v20
    return 0
}
func genFn1(p0 uint8) int16 {
    p0 = p0
    if (45 > 128) {
        fmt.Println(p0)
    } else {
        v21 := GenS1{F0: p0}
        v22 := []int{865, 32767}
        _ = v21
        _ = v22
    }
    v23 := GenS1{F0: p0}
    if ((2 != 1) && false) {
        switch -672 {
        case 9223372036854775807:
            v24 := []int{}
            _ = v24
        case -1:
            v25 := []int{32768, 4294967295, 32767}
            _ = v25
        case 1:
            if true {
                v26 := []int{0, 127, 925, 128}
                _ = v26
            } else {
                v27 := []int{-128, -9223372036854775808, 2147483647, -570}
                p0 -= p0
                _ = v27
            }
        }
    }
    switch "foo bar" {
    case "foo bar":
        if (!(77 < 65535)) {
            v28 := []int{-9223372036854775808, -154}
            var v29 uint = 629
            var v30 int16 = -32768
            _ = v28
            _ = v29
            _ = v30
        } else {
            for fi31 := 0; fi31 < 3; fi31++ {
                var v32 float32 = 1.5
                switch fi31 {
                case -956:
                    fmt.Println(p0, v32, fi31)
                case 42:
                    var v33 uint16 = 255
                    _ = v33
                default:
                    var v34 float32 = (v32 * (v32 / v32))
                    _ = v34
                }
                _ = fi31
                _ = v32
            }
            var v35 uint64 = 542
            if (true && true) {
                switch (("foo bar" + "") + "hello") {
                case "hello":
                    var v36 int16 = 0
                    _ = v36
                default:
                    p0 = (p0 * p0)
                }
                fmt.Println(p0, v35)
                for fi37 := 0; fi37 < 15; fi37++ {
                    var v38 bool = (false && false)
                    p0 = (p0 + (^(p0)))
                    fmt.Println(v38, v38)
                    _ = fi37
                    _ = v38
                }
            } else {
                v35 = (v35 / (((v35 % (((v35 << ((1) & 63))) | 1))) | 1))
                v35 -= 4294967295
            }
            _ = v35
        }
    case "hello":
        var v39 int = -574
        _ = v39
    default:
        for fi40 := 0; fi40 < 4; fi40++ {
            var v41 bool = true
            _ = fi40
            _ = v41
        }
    }
    _ = p0
    _ = v23
    return 1000
}

func main() {
    v1 := make(map[string]bool)
    v2 := []uint32{4294967295, 4294967295}
    var v3 uint8 = 255
    var v4 float64 = (-2.25 * 0.0)
    v5 := make(map[int]bool)
    var v6 int8 = 37
    v4 += v4
    var v42 uint64 = 869
    fmt.Println(v3, v4, v42)
    v4 = v4
    switch 931 {
    case 9223372036854775807:
        switch "a" {
        case "hello":
            if (-9223372036854775808 < 9223372036854775807) {
                if true {
                    fmt.Println(v4)
                    v3 = 1
                } else {
                    var v43 int = 2147483647
                    _ = v43
                }
                v44 := GenS5{F0: ("a" + ""), F1: 9223372036854775807, F2: 904}
                _ = v44
            } else {
                for fi45 := 0; fi45 < 15; fi45++ {
                    if false {
                        if (!(!true)) {
                            var v46 int8 = v6
                            var v47 uint16 = 65535
                            _ = v46
                            _ = v47
                        } else {
                            var v48 int = (fi45 + (fi45 | (fi45 + (fi45 / (((fi45 & (fi45 - 673))) | 1)))))
                            _ = v48
                        }
                        v49 := v1["hello"]
                        v50 := GenS2{F0: (v6 + (v6 * -110)), F1: -605, F2: (v42 / ((v42) | 1))}
                        _ = v49
                        _ = v50
                    }
                    if (false && true) {
                        var v51 int16 = 904
                        _ = v51
                    } else {
                        delete(v1, ("foo bar" + "hello"))
                        switch (("\n\t" + "hello") + ("a" + "\n\t")) {
                        case "":
                            var v52 uint64 = v42
                            _ = v52
                        case "\n\t":
                            fmt.Println(fi45, v3, v4)
                        case "foo bar":
                            for fi53 := 0; fi53 < 7; fi53++ {
                                var v54 int = fi45
                                v42 = v42
                                switch fi45 {
                                case 0:
                                    delete(v1, ("\n\t" + "foo bar"))
                                case -1:
                                    if true {
                                        v3 = (v3 % ((255) | 1))
                                    }
                                case 32768:
                                    var v55 uint32 = 798
                                    _ = v55
                                default:
                                    v3 = v3
                                }
                                v54 -= v54
                                _ = fi53
                                _ = v54
                            }
                        default:
                            v56 := GenS6{F0: 18446744073709551615, F1: 474, F2: true}
                            _ = v56
                        }
                        switch (fi45 % ((674) | 1)) {
                        case 9223372036854775807:
                            v57 := GenS1{F0: v3}
                            _ = v57
                        }
                    }
                    var v58 uint64 = 4294967295
                    _ = fi45
                    _ = v58
                }
                var v59 uint8 = 128
                if (!(true && ((v42 + v42) >= 1))) {
                    if false {
                        if true {
                            v60 := GenS7{F0: (v6 * -128), F1: ""}
                            _ = v60
                        }
                        v61 := GenS6{F0: 0, F1: 16, F2: (false || true)}
                        v62 := GenS8{F0: -32768}
                        _ = v61
                        _ = v62
                    } else {
                        v4 += -2.25
                        switch ("hello" + "foo bar") {
                        case "foo bar":
                            v59 += 233
                        case "\n\t":
                            for fi63 := 0; fi63 < 14; fi63++ {
                                v42 += v42
                                v42 = (v42 - uint64(v3))
                                _ = fi63
                            }
                        default:
                            fmt.Println(v42)
                        }
                    }
                    if false {
                        var v64 float32 = 1e10
                        v65 := GenS7{F0: 42, F1: "\n\t"}
                        _ = v64
                        _ = v65
                    }
                } else {
                    switch "\n\t" {
                    case "hello":
                        v66 := GenS4{F0: "", F1: 42, F2: 2}
                        _ = v66
                    case "foo bar":
                        switch -9 {
                        case 169:
                            if ((v3 << ((42) & 63)) <= (v59 / (((^(v59))) | 1))) {
                                var v67 float32 = ((float32((1e10 + (float64(1e10) + (float64(1.5) * (0.1 / (v4 * ((1e10 + ((v4 + v4) - (-2.25 - float64(-2.25)))) / v4))))))) * -2.25) - -2.25)
                                var v68 uint64 = v42
                                _ = v67
                                _ = v68
                            }
                        case 42:
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        case 2147483647:
                            for fi69 := 0; fi69 < 8; fi69++ {
                                v4 = v4
                                v4 += v4
                                v3 -= v3
                                _ = fi69
                            }
                        default:
                            var v70 float64 = v4
                            _ = v70
                        }
                    default:
                        v42 -= (^(v42))
                    }
                    v42 = 1
                }
                _ = v59
            }
        case "a":
            if true {
                var v71 float64 = -2.25
                for fi72 := 0; fi72 < 20; fi72++ {
                    v73 := GenS7{F0: 42, F1: "\n\t"}
                    _ = fi72
                    _ = v73
                }
                fmt.Println(v42)
                _ = v71
            }
        case "":
            for fi74 := 0; fi74 < 4; fi74++ {
                for fi75 := 0; fi75 < 8; fi75++ {
                    var v76 uint64 = 0
                    _ = fi75
                    _ = v76
                }
                _ = fi74
            }
        }
    case -2147483648:
        switch "hello" {
        case "hello":
            for fi77 := 0; fi77 < 9; fi77++ {
                for fi78 := 0; fi78 < 7; fi78++ {
                    var v79 int = (fi77 * fi78)
                    v80 := GenS3{F0: -377, F1: -32768, F2: 937}
                    if false {
                        if (1.5 != float32(1e10)) {
                            fmt.Println(v4, v6)
                            var v81 int64 = 0
                            fmt.Println(v81, fi77, v3)
                            var v82 int64 = -9223372036854775808
                            _ = v81
                            _ = v82
                        } else {
                            v83 := GenS9{F0: "\n\t", F1: (fi77 ^ fi77)}
                            if true {
                                _ = len(v2)
                                if (v79 > (-(fi78))) {
                                    v4 = float64((3.14159 * 0.0))
                                    v84 := GenS1{F0: (v3 << ((0) & 63))}
                                    switch 42 {
                                    case 9223372036854775807:
                                        if (!false) {
                                            v85 := GenS3{F0: 32767, F1: -645, F2: 399}
                                            _ = v85
                                        }
                                    case 821:
                                        v86 := GenS8{F0: -32768}
                                        _ = v86
                                    }
                                    _ = v84
                                }
                            }
                            v3 -= v3
                            _ = v83
                        }
                        switch "foo bar" {
                        case "foo bar":
                            var v87 int16 = 1000
                            _ = v87
                        case "hello":
                            switch (fi78 << ((2) & 63)) {
                            case 4294967295:
                                fmt.Println(fi77)
                            }
                            v6 = v6
                            v3 -= 255
                        default:
                            v42 -= 0
                        }
                        v6 += 100
                        if (false || (true || false)) {
                            delete(v1, ("hello" + ("a" + ("foo bar" + (("a" + "a") + ("\n\t" + "a"))))))
                        }
                    }
                    _ = fi78
                    _ = v79
                    _ = v80
                }
                for fi88 := 0; fi88 < 12; fi88++ {
                    var v89 string = ""
                    v90 := GenS4{F0: v89, F1: 65535, F2: genFn0(280, uint32(v42))}
                    _ = fi88
                    _ = v89
                    _ = v90
                }
                _ = fi77
            }
        case "\n\t":
            if false {
                for fi91 := 0; fi91 < 18; fi91++ {
                    v42 = v42
                    v92 := GenS10{F0: false}
                    var v93 uint16 = 1
                    _ = fi91
                    _ = v92
                    _ = v93
                }
            }
        case "a":
            if false {
                if false {
                    v94 := GenS4{F0: "foo bar", F1: 130, F2: 255}
                    if false {
                        v42 = v42
                        var v95 uint64 = 2
                        v4 = v4
                        _ = v95
                    }
                    _ = v94
                }
                if len(v2) > 0 {
                    v2[0] = 2147483648
                }
                var v96 bool = (false && (!((v42 % ((725) | 1)) >= (v42 * v42))))
                _ = v96
            }
        default:
            v6 = (v6 | (^(v6)))
        }
    case -736:
        var v97 string = (((("hello" + "\n\t") + "a") + "") + ("" + "foo bar"))
        _ = v97
    }
    var v98 bool = true
    if len(v2) > 0 {
        v2[0] = 256
    }
    if (!(((0 > 677) || (1e10 != 0.1)) && ((v42 % ((v42) | 1)) >= v42))) {
        if (("foo bar" + "") == "hello") {
            if (65535 != uint(v42)) {
                _ = len(v2)
                v99 := GenS10{F0: (172 <= 255)}
                _ = v99
            } else {
                _ = len(v5)
                v100 := GenS8{F0: 42}
                for fi101 := 0; fi101 < 11; fi101++ {
                    fmt.Println(v4, v6, v3)
                    _ = fi101
                }
                _ = v100
            }
            var v102 int8 = (v6 >> ((genFn0(390, 42)) & 63))
            for fi103 := 0; fi103 < 13; fi103++ {
                var v104 int = (fi103 - fi103)
                switch (v104 & fi103) {
                case 128:
                    v105 := GenS11{F0: v42}
                    _ = v105
                case -9223372036854775808:
                    v106 := GenS11{F0: 1}
                    _ = v106
                default:
                    var v107 int64 = -286
                    _ = v107
                }
                _ = fi103
                _ = v104
            }
            _ = v102
        }
        var v108 uint16 = 42
        for fi109 := 0; fi109 < 20; fi109++ {
            var v110 int = (fi109 + fi109)
            fmt.Println(v3, v4, v3)
            v2 = append(v2, 1)
            _ = fi109
            _ = v110
        }
        _ = v108
    } else {
        var v111 float32 = 3.14159
        _ = v111
    }
    for fi112 := 0; fi112 < 8; fi112++ {
        switch 32768 {
        case -995:
            if true {
                v1[("foo bar" + ("hello" + ("" + ((("a" + "hello") + "a") + "\n\t"))))] = (!true)
            } else {
                for fi113 := 0; fi113 < 4; fi113++ {
                    v114 := GenS8{F0: 42}
                    var v115 uint8 = v3
                    v98 = v98
                    _ = fi113
                    _ = v114
                    _ = v115
                }
            }
        case -9223372036854775808:
            v42 += v42
        case -731:
            v42 = (v42 << ((499) & 63))
        }
        v98 = v98
        fmt.Println(v4, v42)
        _ = fi112
    }
    switch "foo bar" {
    case "":
        switch ("a" + "") {
        case "\n\t":
            if (790 > 128) {
                var v116 int8 = v6
                _ = v116
            }
        case "foo bar":
            if len(v2) > 0 {
                v2[0] = 65535
            }
        case "":
            switch ("hello" + ("foo bar" + "foo bar")) {
            case "\n\t":
                var v117 uint = 722
                _ = v117
            case "foo bar":
                v118 := GenS1{F0: v3}
                _ = v118
            default:
                if v98 {
                    fmt.Println(v3, v3, v4)
                } else {
                    for fi119 := 0; fi119 < 12; fi119++ {
                        v120 := GenS6{F0: 1, F1: 628, F2: false}
                        v4 = 1e10
                        if len(v2) > 0 {
                            v2[0] = 625
                        }
                        _ = fi119
                        _ = v120
                    }
                    if (!false) {
                        var v121 uint = 264
                        v122 := GenS1{F0: 1}
                        _ = v121
                        _ = v122
                    }
                }
            }
        default:
            fmt.Println(v6)
        }
    case "a":
        if v98 {
            if len(v2) > 0 {
                v2[0] = 762
            }
        }
    case "hello":
        var v123 int16 = 1000
        _ = v123
    default:
        if (uint16(v42) >= 1) {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            v124 := GenS5{F0: "foo bar", F1: 771, F2: 88}
            v42 += v42
            _ = v124
        }
    }
    fmt.Println(v3)
    var v125 uint16 = 65535
    fmt.Println(v125)
    if v98 {
        for fi126 := 0; fi126 < 1; fi126++ {
            fmt.Println(v3, v3, v98)
            var v127 int8 = (-(v6))
            _ = fi126
            _ = v127
        }
    }
    v125 += v125
    switch "hello" {
    case "foo bar":
        if true {
            if (127 == 255) {
                v125 += (^(v125))
                for fi128 := 0; fi128 < 6; fi128++ {
                    for fi129 := 0; fi129 < 10; fi129++ {
                        if v98 {
                            switch -9223372036854775808 {
                            case 32767:
                                v4 = 3.14159
                            case 42:
                                if v98 {
                                    switch "foo bar" {
                                    case "\n\t":
                                        v6 = -128
                                    case "":
                                        v130 := GenS12{F0: (-(fi129)), F1: 18446744073709551615}
                                        _ = v130
                                    case "hello":
                                        v131 := GenS13{F0: v42, F1: 1000}
                                        _ = v131
                                    default:
                                        if ((fi128 & -2147483648) > (fi128 & fi129)) {
                                            v125 = v125
                                            var v132 int = fi128
                                            _ = v132
                                        } else {
                                            var v133 int8 = (v6 >> ((uint(v125)) & 63))
                                            if v98 {
                                                if len(v2) > 0 {
                                                    v2[0] = 695
                                                }
                                                var v134 int64 = 578
                                                _ = v134
                                            } else {
                                                switch "a" {
                                                case "a":
                                                    if ((-106 > 1) && true) {
                                                        switch fi128 {
                                                        case 127:
                                                            v6 += (v133 >> ((127) & 63))
                                                        case -928:
                                                            v133 -= (-(v6))
                                                        case -834:
                                                            _ = len(v5)
                                                        }
                                                    } else {
                                                        fmt.Println(v42)
                                                        fmt.Println(v4, fi128, fi129)
                                                        v98 = (1 > 845)
                                                    }
                                                default:
                                                    fmt.Println(fi128, v133, v133)
                                                }
                                                switch (fi129 >> ((4294967295) & 63)) {
                                                case -1:
                                                    if v98 {
                                                        var v135 uint16 = (v125 * 42)
                                                        _ = v135
                                                    }
                                                default:
                                                    var v136 int = fi128
                                                    _ = v136
                                                }
                                                var v137 uint8 = (v3 & 127)
                                                _ = v137
                                            }
                                            if (v3 >= (v3 & v3)) {
                                                v138 := GenS14{F0: "\n\t", F1: -2147483648}
                                                var v139 uint16 = (v125 ^ (^(v125)))
                                                v140 := GenS1{F0: v3}
                                                _ = v138
                                                _ = v139
                                                _ = v140
                                            }
                                            _ = v133
                                        }
                                    }
                                    fmt.Println(v4, fi129, v42)
                                    v125 = (v125 << ((255) & 63))
                                    fmt.Println(v3, fi128, v3)
                                    v141 := v1["\n\t"]
                                    _ = v141
                                } else {
                                    v142 := GenS13{F0: (v42 % (((^(v42))) | 1)), F1: 257}
                                    _ = v142
                                }
                            case -1:
                                switch (-(fi128)) {
                                case 0:
                                    fmt.Println(v125, v125, v98)
                                case 372:
                                    v125 = (v125 * v125)
                                case 966:
                                    var v143 int = 42
                                    _ = v143
                                default:
                                    switch ("\n\t" + (("hello" + ("hello" + "\n\t")) + "foo bar")) {
                                    case "\n\t":
                                        v6 -= v6
                                    default:
                                        v144 := v1["hello"]
                                        _ = v144
                                    }
                                }
                            }
                            switch -9223372036854775808 {
                            case -2147483648:
                                v145 := GenS9{F0: ("\n\t" + "a"), F1: (fi129 | (fi128 & 127))}
                                _ = v145
                            case -1:
                                if true {
                                    v4 = v4
                                    v146 := GenS2{F0: -1, F1: 900, F2: 2}
                                    fmt.Println(fi129, v42)
                                    switch "\n\t" {
                                    case "":
                                        v125 = v125
                                    case "foo bar":
                                        v98 = v98
                                    case "a":
                                        switch (fi129 % ((32768) | 1)) {
                                        case 688:
                                            v147 := GenS6{F0: uint(v42), F1: uint32(v125), F2: (0 < 251)}
                                            _ = v147
                                        }
                                    }
                                    _ = len(v2)
                                    if (0 != 32767) {
                                        var v148 float32 = -2.25
                                        var v149 int32 = -2147483648
                                        v42 = v42
                                        v98 = false
                                        switch -749 {
                                        case 42:
                                            var v150 uint32 = 2
                                            _ = v150
                                        }
                                        var v151 bool = (-400 >= v149)
                                        v152 := GenS1{F0: 128}
                                        _ = v148
                                        _ = v149
                                        _ = v151
                                        _ = v152
                                    } else {
                                        v153 := GenS9{F0: "foo bar", F1: (fi128 + fi129)}
                                        fmt.Println(fi129, fi129, v6)
                                        _ = v153
                                    }
                                    _ = v146
                                } else {
                                    fmt.Println(v3)
                                    if ((-9223372036854775808 >= -1) && (0 >= 958)) {
                                        if (-2.25 >= 0.1) {
                                            if v98 {
                                                v1["foo bar"] = v98
                                                v125 = uint16(v3)
                                                if (((v125 / ((2) | 1)) <= 32767) || v98) {
                                                    var v154 int8 = 1
                                                    _ = v154
                                                }
                                                var v155 uint64 = v42
                                                v156 := GenS7{F0: (-(v6)), F1: "a"}
                                                var v157 int = fi128
                                                v42 -= v155
                                                _ = v155
                                                _ = v156
                                                _ = v157
                                            } else {
                                                switch 128 {
                                                case 253:
                                                    v158 := GenS15{F0: 982}
                                                    fmt.Println(v42, fi128)
                                                    v2 = append(v2, 65535)
                                                    _ = v158
                                                case 691:
                                                    var v159 int = fi129
                                                    _ = v159
                                                }
                                            }
                                            var v160 int32 = -750
                                            switch (fi128 % ((fi129) | 1)) {
                                            case -2147483648:
                                                switch (fi129 | -574) {
                                                case 0:
                                                    var v161 uint64 = 1
                                                    _ = v161
                                                default:
                                                    switch -9223372036854775808 {
                                                    case 2147483647:
                                                        v4 -= (((float64(1.5) + (v4 - -2.25)) + float64(1.5)) / v4)
                                                    case 291:
                                                        fmt.Println(fi129)
                                                    case -482:
                                                        switch -778 {
                                                        case -293:
                                                            v6 += v6
                                                        case 127:
                                                            var v162 float32 = 0.0
                                                            _ = v162
                                                        default:
                                                            var v163 int32 = v160
                                                            _ = v163
                                                        }
                                                    }
                                                }
                                            case -9223372036854775808:
                                                v164 := GenS15{F0: uint32(v125)}
                                                _ = v164
                                            case 127:
                                                delete(v1, "hello")
                                            }
                                            _ = v160
                                        } else {
                                            if v98 {
                                                v165 := GenS16{F0: ("a" + "foo bar"), F1: 1, F2: 24}
                                                var v166 int = fi129
                                                _ = v165
                                                _ = v166
                                            } else {
                                                v167 := GenS6{F0: 42, F1: 949, F2: v98}
                                                v3 = 255
                                                _ = v167
                                            }
                                        }
                                        v125 = (v125 | (v125 % (((^(v125))) | 1)))
                                        if v98 {
                                            v168 := GenS17{F0: (fi129 | fi129), F1: (^(fi128))}
                                            v169 := GenS17{F0: (^(fi128)), F1: fi129}
                                            _ = v168
                                            _ = v169
                                        }
                                    } else {
                                        v2 = append(v2, 2147483648)
                                    }
                                    var v170 uint32 = 107
                                    _ = v170
                                }
                            case -9223372036854775808:
                                v171 := v5[(fi128 + (fi128 & (fi129 - (fi129 & 128))))]
                                _ = v171
                            }
                            var v172 int32 = -590
                            _ = v172
                        }
                        v3 += v3
                        _ = fi129
                    }
                    if ((1e10 / 0.1) < (0.1 - 1e10)) {
                        v6 = v6
                        if v98 {
                            var v173 uint = 42
                            fmt.Println(v3)
                            _ = v173
                        } else {
                            for fi174 := 0; fi174 < 16; fi174++ {
                                fmt.Println(v6)
                                if (!(0 != genFn1(v3))) {
                                    fmt.Println(v3, v125)
                                } else {
                                    var v175 int8 = 127
                                    _ = v175
                                }
                                _ = fi174
                            }
                            v98 = false
                            if true {
                                v2 = append(v2, 497)
                                if v98 {
                                    if (!true) {
                                        if len(v2) > 0 {
                                            v2[0] = 42
                                        }
                                        var v176 int64 = 74
                                        _ = v176
                                    }
                                    fmt.Println(v3, v6, v125)
                                }
                            } else {
                                if true {
                                    v98 = (-2147483648 > 2147483647)
                                }
                                v4 += -2.25
                                var v177 string = (("" + ("\n\t" + "a")) + "\n\t")
                                _ = v177
                            }
                        }
                    } else {
                        v178 := GenS10{F0: v98}
                        _ = v178
                    }
                    v179 := GenS18{F0: (v125 - (v125 + (v125 ^ (^(v125))))), F1: 2147483647}
                    _ = fi128
                    _ = v179
                }
            } else {
                var v180 float32 = 0.0
                if (v42 >= 93) {
                    for fi181 := 0; fi181 < 19; fi181++ {
                        v98 = (32767 >= 740)
                        _ = fi181
                    }
                } else {
                    v42 -= v42
                    _ = len(v5)
                    if v98 {
                        fmt.Println(v3, v4)
                    } else {
                        v2 = append(v2, 493)
                        switch ("" + ("" + ("hello" + "hello"))) {
                        case "foo bar":
                            fmt.Println(v125, v3, v3)
                        case "a":
                            v125 = 0
                        default:
                            var v182 uint8 = 42
                            _ = v182
                        }
                    }
                }
                v183 := GenS19{F0: 956, F1: (^(v42))}
                _ = v180
                _ = v183
            }
        } else {
            _ = len(v1)
            v4 += (1.5 - float64((3.14159 - ((((0.1 - (float32(1.5) / (1e10 + 0.0))) / ((1.5 - 0.0) - ((1e10 * ((1e10 - (float32(v4) * 3.14159)) - ((0.0 / 1.5) - (1e10 / 1.5)))) / 1.0))) - 0.0) / 1.0))))
        }
    case "hello":
        if v98 {
            v184 := GenS20{F0: -9223372036854775808, F1: 1000, F2: 340}
            switch (("a" + ("foo bar" + "\n\t")) + ("" + "\n\t")) {
            case "hello":
                var v185 uint8 = 127
                _ = v185
            case "\n\t":
                fmt.Println(v125, v6, v42)
            default:
                var v186 bool = v98
                _ = v186
            }
            _ = v184
        }
    case "a":
        var v187 uint32 = 65535
        _ = v187
    }
    v4 -= -2.25
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v125:", v125)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v98:", v98)
    _ = v1
    _ = v125
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v5
    _ = v6
    _ = v98
}
