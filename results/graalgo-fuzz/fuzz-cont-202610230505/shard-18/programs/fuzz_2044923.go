package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64; F2 uint64 }
type GenS2 struct { F0 uint32; F1 string }
type GenS3 struct { F0 int16; F1 int64; F2 int32 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 bool; F1 int64 }
type GenS7 struct { F0 int64; F1 int64 }
type GenS8 struct { F0 int8; F1 int32; F2 int16 }
type GenS9 struct { F0 int32; F1 bool }
type GenS10 struct { F0 uint8 }
type GenS11 struct { F0 int16 }
type GenS12 struct { F0 string }
type GenS13 struct { F0 uint8; F1 int64; F2 bool }
type GenS14 struct { F0 uint8 }
type GenS15 struct { F0 uint64; F1 uint16 }
func genFn0(p0 uint8) bool {
    if (-32768 != 42) {
        var v4 string = "hello"
        v5 := GenS1{F0: 100, F1: 9223372036854775807, F2: 18446744073709551615}
        v6 := GenS1{F0: -69, F1: 1, F2: 4294967295}
        _ = v4
        _ = v5
        _ = v6
    } else {
        switch 0 {
        case -2147483648:
            var v7 float32 = -2.25
            _ = v7
        case 0:
            p0 += p0
        case -568:
            v8 := []int{42}
            _ = v8
        }
        if (2147483647 <= 0) {
            for fi9 := 0; fi9 < 8; fi9++ {
                if true {
                    if (220 > -931) {
                        var v10 bool = false
                        fmt.Println(p0, fi9, v10)
                        _ = v10
                    }
                    v11 := GenS2{F0: 274, F1: "hello"}
                    var v12 int8 = 101
                    _ = v11
                    _ = v12
                } else {
                    p0 = (p0 / (((p0 % ((p0) | 1))) | 1))
                    var v13 uint32 = 2147483647
                    var v14 uint16 = 65535
                    _ = v13
                    _ = v14
                }
                v15 := []int{fi9}
                _ = fi9
                _ = v15
            }
        } else {
            if true {
                v16 := []int{32768, -177}
                _ = v16
            }
            var v17 uint8 = p0
            _ = v17
        }
    }
    if false {
        var v18 int8 = -106
        fmt.Println(v18, p0)
        _ = v18
    } else {
        v19 := []int{169, 128, -9223372036854775808, -9223372036854775808}
        var v20 int = 127
        _ = v19
        _ = v20
    }
    _ = p0
    return true
}
func genFn1() string {
    var v21 uint8 = 48
    if (42 == 4294967295) {
        switch 0 {
        case 9223372036854775807:
            v21 += 255
        case 1:
            for fi22 := 0; fi22 < 15; fi22++ {
                switch (fi22 / ((128) | 1)) {
                case 2147483647:
                    if (32767 == -32768) {
                        switch fi22 {
                        case -9223372036854775808:
                            v23 := []int{fi22, fi22}
                            _ = v23
                        default:
                            v24 := GenS2{F0: 143, F1: "a"}
                            _ = v24
                        }
                    } else {
                        switch "a" {
                        case "foo bar":
                            fmt.Println(fi22, v21)
                        }
                        v21 = v21
                    }
                default:
                    var v25 uint = 631
                    _ = v25
                }
                _ = fi22
            }
        }
        for fi26 := 0; fi26 < 16; fi26++ {
            if (2 == 532) {
                for fi27 := 0; fi27 < 15; fi27++ {
                    v21 += v21
                    var v28 float32 = (1e10 / (3.14159 / float32((((0.1 - 1.5) * float64(0.0)) + (0.0 + 1.5)))))
                    v29 := GenS1{F0: 42, F1: 42, F2: 4294967295}
                    _ = fi27
                    _ = v28
                    _ = v29
                }
                for fi30 := 0; fi30 < 4; fi30++ {
                    var v31 uint8 = (v21 & (^(v21)))
                    v32 := GenS3{F0: 1000, F1: 458, F2: -1}
                    switch (^(fi30)) {
                    case -355:
                        var v33 uint32 = 839
                        _ = v33
                    case -556:
                        v34 := []int{39, fi26, (fi30 ^ (-(fi30))), int(v31)}
                        _ = v34
                    case 9223372036854775807:
                        v31 = v21
                    default:
                        if (-1 <= 1000) {
                            v31 -= 241
                            var v35 uint = 707
                            var v36 uint16 = 32767
                            _ = v35
                            _ = v36
                        } else {
                            v31 += v31
                            var v37 int8 = -125
                            _ = v37
                        }
                    }
                    _ = fi30
                    _ = v31
                    _ = v32
                }
                v21 = (^(v21))
            } else {
                v21 -= (v21 - 42)
            }
            fmt.Println(v21, fi26)
            var v38 int64 = -248
            _ = fi26
            _ = v38
        }
        v39 := GenS2{F0: 1, F1: ("foo bar" + ("foo bar" + ""))}
        _ = v39
    } else {
        var v40 int = 571
        _ = v40
    }
    _ = v21
    return "foo bar"
}

func main() {
    var v1 uint = 18446744073709551615
    v2 := make(map[uint32]int)
    var v3 uint16 = 0
    for fi41 := 0; fi41 < 16; fi41++ {
        var v42 int8 = 33
        _ = fi41
        _ = v42
    }
    switch -319 {
    case 1:
        v2[1] = -1
    default:
        v3 += (v3 / ((v3) | 1))
    }
    switch 9223372036854775807 {
    case -840:
        switch ("\n\t" + "") {
        case "":
            fmt.Println(v1, v3)
        default:
            var v43 float32 = 1.5
            _ = v43
        }
    default:
        v1 -= v1
    }
    switch 127 {
    case -284:
        if (true || (uint(v3) > 39)) {
            if true {
                var v44 int = 16
                v2[2147483648] = 32767
                if true {
                    v45 := v2[2]
                    v3 = v3
                    for fi46 := 0; fi46 < 15; fi46++ {
                        fmt.Println(v1, v1, v3)
                        var v47 int64 = -490
                        fmt.Println(v1, v45)
                        _ = fi46
                        _ = v47
                    }
                    _ = v45
                } else {
                    var v48 float32 = float32((0.1 + 1e10))
                    fmt.Println(v3, v1, v44)
                    _ = v48
                }
                _ = v44
            } else {
                v3 += v3
            }
            for fi49 := 0; fi49 < 5; fi49++ {
                var v50 uint64 = 991
                fmt.Println(v1)
                var v51 int8 = int8(fi49)
                _ = fi49
                _ = v50
                _ = v51
            }
        } else {
            v3 -= v3
            v52 := []int{42, 42}
            _ = v52
        }
    default:
        fmt.Println(v3, v1)
    }
    for fi53 := 0; fi53 < 16; fi53++ {
        v3 += (v3 / ((766) | 1))
        if ((fi53 == (fi53 % ((-2147483648) | 1))) && (true || false)) {
            v3 = 65535
        } else {
            _ = len(v2)
            v3 += (v3 / ((v3) | 1))
        }
        _ = fi53
    }
    v54 := []int{390}
    var v55 uint = (v1 >> ((v1) & 63))
    fmt.Println(v3, v1, v55)
    if false {
        var v56 int = 32767
        _ = v56
    }
    v57 := GenS4{F0: 1000}
    v58 := GenS5{F0: -243}
    if true {
        switch "\n\t" {
        case "hello":
            v3 -= (v3 - v3)
        default:
            v55 = v55
        }
        for fi59 := 0; fi59 < 14; fi59++ {
            v60 := GenS4{F0: 1000}
            v61 := GenS3{F0: -209, F1: 1, F2: 601}
            for fi62 := 0; fi62 < 12; fi62++ {
                fmt.Println(v3, v55, v55)
                switch fi59 {
                case 32768:
                    v55 += (v55 / (((v1 + 465)) | 1))
                case 722:
                    switch (^(fi59)) {
                    case -2147483648:
                        fmt.Println(fi59, v3)
                    case 32768:
                        if len(v54) > 0 {
                            v54[0] = (fi62 >> ((v55) & 63))
                        }
                    }
                case 128:
                    var v63 int = 573
                    _ = v63
                default:
                    if (1 < 2147483647) {
                        var v64 int = 9223372036854775807
                        v1 -= v1
                        _ = v64
                    } else {
                        v2[28] = -1
                        v54 = append(v54, (fi62 ^ (fi62 | 32768)))
                        var v65 uint16 = v3
                        _ = v65
                    }
                }
                v3 = (v3 << ((65535) & 63))
                _ = fi62
            }
            _ = fi59
            _ = v60
            _ = v61
        }
        if false {
            switch ("hello" + "foo bar") {
            case "\n\t":
                switch 127 {
                case 314:
                    fmt.Println(v55, v1, v1)
                case 5:
                    v1 = v1
                case 42:
                    for fi66 := 0; fi66 < 18; fi66++ {
                        fmt.Println(v1, v3, fi66)
                        _ = fi66
                    }
                default:
                    for fi67 := 0; fi67 < 5; fi67++ {
                        for fi68 := 0; fi68 < 14; fi68++ {
                            if len(v54) >= 2 {
                                v54 = v54[1:len(v54)]
                            }
                            var v69 float32 = 1e10
                            fmt.Println(v1)
                            _ = fi68
                            _ = v69
                        }
                        for fi70 := 0; fi70 < 13; fi70++ {
                            switch ("a" + ("hello" + "\n\t")) {
                            case "":
                                v3 = (^(v3))
                            case "a":
                                if (-1 <= 9223372036854775807) {
                                    v71 := GenS3{F0: 32767, F1: -701, F2: 1}
                                    _ = v71
                                }
                            case "foo bar":
                                v55 = 4294967295
                            }
                            var v72 uint64 = 0
                            _ = fi70
                            _ = v72
                        }
                        _ = fi67
                    }
                }
            case "foo bar":
                var v73 float64 = 1.5
                _ = v73
            }
            fmt.Println(v1, v1, v1)
            var v74 bool = ("" > "hello")
            _ = v74
        } else {
            if (255 < v1) {
                switch "a" {
                case "a":
                    if ((!(-1 == 237)) && ((9223372036854775807 != 42) || false)) {
                        v55 = (v1 & v1)
                        fmt.Println(v55, v1, v55)
                        switch -1 {
                        case 9223372036854775807:
                            var v75 int16 = 897
                            _ = v75
                        case -555:
                            if (281 != 245) {
                                v3 += v3
                            }
                        }
                    } else {
                        if (!(false || true)) {
                            var v76 uint16 = (v3 ^ 255)
                            _ = v76
                        }
                        if false {
                            switch 4294967295 {
                            case 9223372036854775807:
                                for fi77 := 0; fi77 < 14; fi77++ {
                                    var v78 int8 = 0
                                    _ = fi77
                                    _ = v78
                                }
                            default:
                                v79 := GenS2{F0: 499, F1: ("" + "")}
                                _ = v79
                            }
                            switch "" {
                            case "hello":
                                if (-128 != 127) {
                                    switch -2147483648 {
                                    case 1:
                                        v80 := GenS2{F0: 65535, F1: ""}
                                        _ = v80
                                    default:
                                        _ = len(v2)
                                    }
                                    if (-926 > 339) {
                                        var v81 int8 = 127
                                        v82 := GenS6{F0: (4294967295 >= 4294967295), F1: 843}
                                        var v83 int8 = v81
                                        _ = v81
                                        _ = v82
                                        _ = v83
                                    }
                                    v84 := v2[292]
                                    _ = v84
                                }
                            case "a":
                                for fi85 := 0; fi85 < 18; fi85++ {
                                    switch (fi85 << ((2) & 63)) {
                                    case -2147483648:
                                        for fi86 := 0; fi86 < 4; fi86++ {
                                            v3 = v3
                                            var v87 int32 = 2147483647
                                            var v88 uint8 = 42
                                            _ = fi86
                                            _ = v87
                                            _ = v88
                                        }
                                    case 4294967295:
                                        var v89 uint32 = 4294967295
                                        _ = v89
                                    case 32768:
                                        if ("hello" != "hello") {
                                            v54 = append(v54, -2147483648)
                                            v90 := v2[182]
                                            _ = v90
                                        } else {
                                            v1 -= v55
                                            fmt.Println(v3, v3)
                                        }
                                    default:
                                        if (-32768 == 1000) {
                                            for fi91 := 0; fi91 < 17; fi91++ {
                                                v1 = (v1 ^ (v55 & v1))
                                                _ = fi91
                                            }
                                        }
                                    }
                                    _ = fi85
                                }
                            default:
                                var v92 int32 = 42
                                _ = v92
                            }
                            if ((2147483647 == 0) || true) {
                                var v93 int8 = -10
                                fmt.Println(v55)
                                switch "hello" {
                                case "hello":
                                    for fi94 := 0; fi94 < 8; fi94++ {
                                        var v95 bool = false
                                        _ = fi94
                                        _ = v95
                                    }
                                }
                                _ = v93
                            } else {
                                switch "a" {
                                case "foo bar":
                                    if true {
                                        var v96 string = ""
                                        var v97 uint16 = v3
                                        var v98 uint = v1
                                        _ = v96
                                        _ = v97
                                        _ = v98
                                    }
                                case "hello":
                                    for fi99 := 0; fi99 < 2; fi99++ {
                                        v3 += v3
                                        var v100 float32 = (1e10 + (float32(-2.25) / 3.14159))
                                        var v101 uint64 = 42
                                        _ = fi99
                                        _ = v100
                                        _ = v101
                                    }
                                case "":
                                    var v102 bool = (220 > 2147483648)
                                    _ = v102
                                default:
                                    var v103 int = 954
                                    _ = v103
                                }
                            }
                        } else {
                            if (-369 >= 4294967295) {
                                for fi104 := 0; fi104 < 10; fi104++ {
                                    var v105 uint8 = uint8(v3)
                                    v55 = 4294967295
                                    _ = fi104
                                    _ = v105
                                }
                                var v106 uint = (v55 - (^(v55)))
                                _ = v106
                            }
                        }
                        v107 := GenS5{F0: 0}
                        _ = v107
                    }
                case "hello":
                    var v108 uint64 = 430
                    _ = v108
                }
                var v109 int16 = 1000
                _ = v109
            } else {
                if false {
                    switch "" {
                    case "hello":
                        if ((v3 & (^(v3))) != 1000) {
                            var v110 int16 = 1000
                            fmt.Println(v1, v110, v55)
                            v55 += 18446744073709551615
                            _ = v110
                        }
                    case "\n\t":
                        for fi111 := 0; fi111 < 15; fi111++ {
                            for fi112 := 0; fi112 < 4; fi112++ {
                                var v113 string = "a"
                                if true {
                                    if true {
                                        if true {
                                            var v114 int = -9223372036854775808
                                            if false {
                                                var v115 int64 = 1
                                                if (true || false) {
                                                    var v116 int8 = -55
                                                    v113 = v113
                                                    _ = v116
                                                }
                                                _ = v115
                                            }
                                            switch (-(fi111)) {
                                            case -1:
                                                fmt.Println(v114, v1, v114)
                                            case -2147483648:
                                                if (3.14159 <= (0.0 * 0.0)) {
                                                    var v117 float64 = 0.0
                                                    _ = v117
                                                } else {
                                                    if ((true && false) || true) {
                                                        var v118 string = ((v113 + (v113 + ((v113 + v113) + v113))) + v113)
                                                        _ = v118
                                                    }
                                                    var v119 uint = v1
                                                    if false {
                                                        if true {
                                                            switch (fi111 * (^(v114))) {
                                                            case 127:
                                                                fmt.Println(v113)
                                                            }
                                                            v1 -= (v1 % ((v119) | 1))
                                                        } else {
                                                            v3 -= (v3 << (((^(v119))) & 63))
                                                            v55 -= 2
                                                        }
                                                    } else {
                                                        v55 = 1
                                                        v113 = (v113 + "hello")
                                                        var v120 uint64 = 1
                                                        _ = v120
                                                    }
                                                    if len(v54) >= 2 {
                                                        v54 = v54[1:len(v54)]
                                                    }
                                                    var v121 uint32 = 745
                                                    _ = v119
                                                    _ = v121
                                                }
                                            }
                                            _ = v114
                                        } else {
                                            v122 := GenS6{F0: (int32(fi111) >= 0), F1: 42}
                                            _ = v122
                                        }
                                        if (0.0 < 0.0) {
                                            if true {
                                                var v123 int8 = -1
                                                v113 = (("\n\t" + "foo bar") + (((v113 + v113) + v113) + v113))
                                                if (1.5 == -2.25) {
                                                    v113 = ""
                                                } else {
                                                    v3 = 0
                                                }
                                                _ = v123
                                            }
                                            if false {
                                                fmt.Println(v1, v55, v55)
                                                delete(v2, 38)
                                                var v124 bool = true
                                                _ = v124
                                            } else {
                                                v55 -= 65535
                                                v55 -= (^(v55))
                                            }
                                            if false {
                                                var v125 int = 32768
                                                v3 += (v3 << (((v1 ^ (^(v55)))) & 63))
                                                _ = v125
                                            } else {
                                                v126 := GenS7{F0: -1, F1: 42}
                                                _ = v126
                                            }
                                        } else {
                                            var v127 int32 = -368
                                            _ = v127
                                        }
                                        switch v113 {
                                        case "\n\t":
                                            v128 := GenS4{F0: 0}
                                            _ = v128
                                        case "a":
                                            if true {
                                                if false {
                                                    switch 4294967295 {
                                                    case -886:
                                                        v129 := GenS6{F0: ((0.1 * 0.1) >= ((-2.25 + (((float64(float32(float64(0.1))) + ((float64(0.0) / 0.1) / 1.0)) - 1.5) * 1e10)) / 1e10)), F1: -9223372036854775808}
                                                        _ = v129
                                                    }
                                                    var v130 uint16 = v3
                                                    _ = v130
                                                } else {
                                                    var v131 int64 = 0
                                                    _ = v131
                                                }
                                                if (18446744073709551615 != 4294967295) {
                                                    if true {
                                                        v3 = v3
                                                        if false {
                                                            if (128 <= 151) {
                                                                v113 = v113
                                                                v132 := GenS6{F0: false, F1: -9223372036854775808}
                                                                v1 += v55
                                                                var v133 string = ""
                                                                _ = v132
                                                                _ = v133
                                                            }
                                                            v134 := GenS1{F0: -1, F1: -9223372036854775808, F2: uint64(v3)}
                                                            _ = v134
                                                        } else {
                                                            var v135 string = v113
                                                            if false {
                                                                var v136 float32 = 0.0
                                                                var v137 uint8 = 74
                                                                switch -702 {
                                                                case 128:
                                                                    var v138 string = v113
                                                                    _ = v138
                                                                case 499:
                                                                    if true {
                                                                        switch fi112 {
                                                                        case 32768:
                                                                            v3 = (v3 + (v3 % (((^(v3))) | 1)))
                                                                        case 437:
                                                                            var v139 uint8 = (v137 / ((uint8(v1)) | 1))
                                                                            _ = v139
                                                                        case 2147483647:
                                                                            if false {
                                                                                var v140 int32 = -1
                                                                                _ = v140
                                                                            } else {
                                                                                var v141 uint = v1
                                                                                _ = v141
                                                                            }
                                                                        default:
                                                                            fmt.Println(v113, fi111, v3)
                                                                            if ((int64(fi112) == -9223372036854775808) && false) {
                                                                                if ((v3 | v3) == (v3 % ((65535) | 1))) {
                                                                                    var v142 int64 = 288
                                                                                    _ = len(v54)
                                                                                    _ = len(v2)
                                                                                    _ = v142
                                                                                }
                                                                                switch fi112 {
                                                                                case 127:
                                                                                    v143 := GenS8{F0: -113, F1: 2147483647, F2: -32768}
                                                                                    _ = v143
                                                                                case 32767:
                                                                                    v3 = 65535
                                                                                case -2147483648:
                                                                                    if true {
                                                                                        if ((v1 <= (v55 / ((1) | 1))) || (!false)) {
                                                                                            if ((!true) && true) {
                                                                                                v1 = 962
                                                                                                if true {
                                                                                                    if true {
                                                                                                        v3 += v3
                                                                                                    } else {
                                                                                                        var v144 float64 = 1.5
                                                                                                        _ = v144
                                                                                                    }
                                                                                                    var v145 uint32 = 1
                                                                                                    if true {
                                                                                                        var v146 uint16 = (v3 + (v3 ^ v3))
                                                                                                        _ = v146
                                                                                                    }
                                                                                                    _ = v145
                                                                                                }
                                                                                                if false {
                                                                                                    if true {
                                                                                                        switch -414 {
                                                                                                        case 127:
                                                                                                            v147 := GenS9{F0: 425, F1: (("foo bar" + v113) <= ((v113 + "") + (v135 + (v113 + (v113 + "")))))}
                                                                                                            _ = v147
                                                                                                        default:
                                                                                                            var v148 uint64 = 4294967295
                                                                                                            _ = v148
                                                                                                        }
                                                                                                        if (!(v3 < 0)) {
                                                                                                            v137 = (v137 % ((uint8(v55)) | 1))
                                                                                                        } else {
                                                                                                            v136 += v136
                                                                                                            v135 = v113
                                                                                                        }
                                                                                                    }
                                                                                                }
                                                                                            } else {
                                                                                                if (false || (32767 < 1)) {
                                                                                                    v113 = ("hello" + ("hello" + "\n\t"))
                                                                                                    v149 := GenS5{F0: -1}
                                                                                                    var v150 int32 = -2147483648
                                                                                                    var v151 int = 42
                                                                                                    _ = v149
                                                                                                    _ = v150
                                                                                                    _ = v151
                                                                                                } else {
                                                                                                    if true {
                                                                                                        v137 = (^(v137))
                                                                                                    }
                                                                                                    fmt.Println(v135)
                                                                                                    fmt.Println(v113, v135, v136)
                                                                                                }
                                                                                            }
                                                                                            v137 = v137
                                                                                            v3 -= (^(v3))
                                                                                        }
                                                                                        var v152 float32 = v136
                                                                                        if true {
                                                                                            v153 := GenS10{F0: 128}
                                                                                            v154 := GenS6{F0: false, F1: 0}
                                                                                            v137 = 128
                                                                                            _ = v153
                                                                                            _ = v154
                                                                                        } else {
                                                                                            v155 := GenS11{F0: -850}
                                                                                            _ = v155
                                                                                        }
                                                                                        _ = v152
                                                                                    } else {
                                                                                        switch 42 {
                                                                                        case 739:
                                                                                            switch "foo bar" {
                                                                                            case "":
                                                                                                v156 := GenS1{F0: -1, F1: 0, F2: 761}
                                                                                                _ = v156
                                                                                            case "hello":
                                                                                                v55 = v1
                                                                                            default:
                                                                                                v135 = v113
                                                                                            }
                                                                                        case -713:
                                                                                            v3 = 0
                                                                                        default:
                                                                                            if false {
                                                                                                if true {
                                                                                                    delete(v2, 495)
                                                                                                    fmt.Println(v55, v113, fi111)
                                                                                                }
                                                                                            } else {
                                                                                                var v157 uint8 = (^(v137))
                                                                                                fmt.Println(v113)
                                                                                                var v158 int = 128
                                                                                                _ = v157
                                                                                                _ = v158
                                                                                            }
                                                                                        }
                                                                                        v55 = 535
                                                                                    }
                                                                                default:
                                                                                    if false {
                                                                                        switch 127 {
                                                                                        case -2147483648:
                                                                                            switch (v135 + v113) {
                                                                                            case "foo bar":
                                                                                                var v159 uint32 = 0
                                                                                                _ = v159
                                                                                            case "\n\t":
                                                                                                if (false && ((v3 ^ v3) == v3)) {
                                                                                                    v160 := GenS3{F0: 32767, F1: -9223372036854775808, F2: -726}
                                                                                                    fmt.Println(v136, fi112)
                                                                                                    _ = v160
                                                                                                }
                                                                                            }
                                                                                        case 223:
                                                                                            _ = len(v2)
                                                                                        }
                                                                                        var v161 int32 = 42
                                                                                        _ = v161
                                                                                    } else {
                                                                                        switch ("\n\t" + (v113 + (("foo bar" + ((v113 + (((("" + ("\n\t" + (v113 + ((v113 + ("" + "foo bar")) + (v135 + (((v135 + "hello") + v113) + "")))))) + "a") + "foo bar") + ((((v113 + (("hello" + (v113 + (v135 + ("" + "")))) + v135)) + v135) + v113) + ""))) + (v113 + v113))) + v135))) {
                                                                                        case "\n\t":
                                                                                            if (true || (false || false)) {
                                                                                                switch v113 {
                                                                                                case "a":
                                                                                                    v54 = append(v54, 128)
                                                                                                    if ((fi111 ^ fi112) >= fi112) {
                                                                                                        v162 := GenS9{F0: -29, F1: (0 == 42)}
                                                                                                        _ = v162
                                                                                                    }
                                                                                                    if true {
                                                                                                        delete(v2, 4294967295)
                                                                                                    } else {
                                                                                                        var v163 int64 = -9223372036854775808
                                                                                                        _ = v163
                                                                                                    }
                                                                                                case "hello":
                                                                                                    fmt.Println(fi111, v113, fi111)
                                                                                                case "\n\t":
                                                                                                    v135 = "hello"
                                                                                                }
                                                                                                fmt.Println(v55, v136)
                                                                                                v1 = 255
                                                                                                v1 = (v55 * 1)
                                                                                            }
                                                                                        default:
                                                                                            v1 = 452
                                                                                        }
                                                                                    }
                                                                                }
                                                                                var v164 int64 = 9223372036854775807
                                                                                _ = v164
                                                                            }
                                                                            if false {
                                                                                if (-2147483648 == -1) {
                                                                                    v1 -= (v1 | 18446744073709551615)
                                                                                    var v165 int8 = -108
                                                                                    v3 -= v3
                                                                                    _ = v165
                                                                                } else {
                                                                                    var v166 uint = v1
                                                                                    v167 := GenS12{F0: v135}
                                                                                    var v168 string = (("\n\t" + v135) + v135)
                                                                                    _ = v166
                                                                                    _ = v167
                                                                                    _ = v168
                                                                                }
                                                                                if false {
                                                                                    v2[4294967295] = (^(fi112))
                                                                                    if ((((true || (v1 != (v55 - v55))) || false) && (((v137 + (v137 + (v137 | uint8(v3)))) <= 128) && (!true))) || (0 >= (v3 << (((v55 - 18446744073709551615)) & 63)))) {
                                                                                        v3 += 1
                                                                                        if true {
                                                                                            v55 += (v55 * v55)
                                                                                        }
                                                                                        v169 := GenS8{F0: -128, F1: -2147483648, F2: 42}
                                                                                        _ = v169
                                                                                    }
                                                                                } else {
                                                                                    switch -2147483648 {
                                                                                    case 884:
                                                                                        v170 := GenS8{F0: -128, F1: 310, F2: int16(v3)}
                                                                                        _ = v170
                                                                                    default:
                                                                                        v113 = v113
                                                                                    }
                                                                                    var v171 uint = 1
                                                                                    _ = v171
                                                                                }
                                                                            } else {
                                                                                v137 = v137
                                                                                v1 = 65535
                                                                                v172 := GenS4{F0: -1}
                                                                                v1 = v55
                                                                                _ = v172
                                                                            }
                                                                        }
                                                                        v113 = ""
                                                                        var v173 uint16 = v3
                                                                        var v174 string = "a"
                                                                        _ = v173
                                                                        _ = v174
                                                                    }
                                                                }
                                                                _ = v136
                                                                _ = v137
                                                            }
                                                            v135 = v113
                                                            _ = v135
                                                        }
                                                        v1 = v55
                                                    } else {
                                                        v1 = 929
                                                        fmt.Println(v3)
                                                        if len(v54) >= 2 {
                                                            v54 = v54[1:len(v54)]
                                                        }
                                                        if ((fi112 >> (((v1 << ((879) & 63))) & 63)) != fi111) {
                                                            v175 := GenS12{F0: "\n\t"}
                                                            v55 -= (v1 % ((544) | 1))
                                                            _ = v175
                                                        } else {
                                                            v1 -= 846
                                                            var v176 uint = (v55 & v1)
                                                            var v177 float64 = 1.5
                                                            _ = v176
                                                            _ = v177
                                                        }
                                                    }
                                                } else {
                                                    var v178 uint64 = uint64(v3)
                                                    _ = v178
                                                }
                                            }
                                            switch fi111 {
                                            case 4294967295:
                                                if ((false || false) || (2147483647 >= -2147483648)) {
                                                    fmt.Println(v55, fi111)
                                                } else {
                                                    delete(v2, uint32(v3))
                                                    v54 = append(v54, (^(fi112)))
                                                }
                                            }
                                        }
                                    }
                                    v179 := GenS10{F0: 255}
                                    _ = v179
                                } else {
                                    if (v1 < v1) {
                                        var v180 int16 = -453
                                        delete(v2, 2147483647)
                                        _ = v180
                                    } else {
                                        var v181 string = "\n\t"
                                        _ = len(v2)
                                        var v182 uint16 = 42
                                        _ = v181
                                        _ = v182
                                    }
                                    var v183 bool = false
                                    v184 := v2[0]
                                    _ = v183
                                    _ = v184
                                }
                                v185 := GenS8{F0: 1, F1: 0, F2: -270}
                                _ = fi112
                                _ = v113
                                _ = v185
                            }
                            var v186 uint8 = 42
                            _ = fi111
                            _ = v186
                        }
                    default:
                        delete(v2, 2)
                    }
                } else {
                    v1 = (v55 ^ (v1 ^ v55))
                }
                if (((float64(-2.25) / 1.0) - (float64((0.0 + float32(((1e10 / -2.25) / 1.0)))) / 1e10)) >= 1e10) {
                    for fi187 := 0; fi187 < 6; fi187++ {
                        v1 = (^(v55))
                        if (-32768 > 32767) {
                            if (("foo bar" + "hello") != ("" + ("foo bar" + "hello"))) {
                                var v188 string = ("a" + "\n\t")
                                _ = v188
                            } else {
                                v3 = 32767
                                var v189 int = fi187
                                if (-1 >= 42) {
                                    v1 -= v55
                                }
                                _ = v189
                            }
                            for fi190 := 0; fi190 < 16; fi190++ {
                                v191 := GenS11{F0: 32767}
                                _ = fi190
                                _ = v191
                            }
                            if (738 == -1) {
                                v192 := GenS2{F0: 273, F1: "foo bar"}
                                _ = v192
                            } else {
                                if false {
                                    v193 := GenS13{F0: 128, F1: -9223372036854775808, F2: false}
                                    if (18446744073709551615 != 773) {
                                        if (!false) {
                                            v55 += v55
                                            v2[42] = (fi187 - (-(fi187)))
                                            v194 := GenS14{F0: 92}
                                            _ = v194
                                        }
                                        if ("a" >= (("a" + ("foo bar" + "a")) + "foo bar")) {
                                            v3 = (v3 * v3)
                                            if ((0.0 * 1.5) == 3.14159) {
                                                if (true || true) {
                                                    switch "hello" {
                                                    case "\n\t":
                                                        var v195 uint16 = v3
                                                        _ = v195
                                                    }
                                                    for fi196 := 0; fi196 < 10; fi196++ {
                                                        v55 = v1
                                                        v3 = (v3 >> ((1) & 63))
                                                        _ = fi196
                                                    }
                                                }
                                                var v197 float64 = 0.1
                                                _ = v197
                                            }
                                        } else {
                                            v55 += v55
                                        }
                                        var v198 uint32 = 383
                                        _ = v198
                                    }
                                    for fi199 := 0; fi199 < 11; fi199++ {
                                        switch (fi187 | (fi199 - -810)) {
                                        case 128:
                                            var v200 uint16 = v3
                                            _ = v200
                                        default:
                                            v201 := GenS7{F0: -630, F1: 9223372036854775807}
                                            _ = v201
                                        }
                                        if len(v54) >= 2 {
                                            v54 = v54[1:len(v54)]
                                        }
                                        if false {
                                            var v202 float32 = (0.0 - 1e10)
                                            _ = v202
                                        } else {
                                            v1 += (v55 % ((v1) | 1))
                                        }
                                        _ = fi199
                                    }
                                    _ = v193
                                }
                                if true {
                                    for fi203 := 0; fi203 < 20; fi203++ {
                                        v204 := GenS10{F0: 128}
                                        _ = fi203
                                        _ = v204
                                    }
                                }
                                v205 := GenS5{F0: 463}
                                _ = v205
                            }
                        } else {
                            v3 = 32767
                        }
                        fmt.Println(fi187, v1, v55)
                        _ = fi187
                    }
                } else {
                    switch "foo bar" {
                    case "\n\t":
                        if true {
                            for fi206 := 0; fi206 < 14; fi206++ {
                                fmt.Println(fi206, fi206, v1)
                                _ = fi206
                            }
                        }
                    case "a":
                        for fi207 := 0; fi207 < 3; fi207++ {
                            if true {
                                if len(v54) > 0 {
                                    v54[0] = fi207
                                }
                            } else {
                                v208 := GenS8{F0: 68, F1: int32(v3), F2: 32767}
                                for fi209 := 0; fi209 < 6; fi209++ {
                                    v210 := GenS15{F0: 958, F1: 42}
                                    switch (fi209 * int(v1)) {
                                    case -267:
                                        var v211 uint = v55
                                        _ = v211
                                    case 127:
                                        v3 -= v3
                                    default:
                                        var v212 int8 = 42
                                        _ = v212
                                    }
                                    v54 = append(v54, (fi207 ^ fi209))
                                    v3 = (v3 ^ (^(v3)))
                                    delete(v2, 624)
                                    _ = fi209
                                    _ = v210
                                }
                                for fi213 := 0; fi213 < 17; fi213++ {
                                    v1 = v1
                                    var v214 uint16 = 65535
                                    var v215 uint64 = 1
                                    _ = fi213
                                    _ = v214
                                    _ = v215
                                }
                                _ = v208
                            }
                            var v216 uint = v1
                            _ = fi207
                            _ = v216
                        }
                    }
                }
            }
            for fi217 := 0; fi217 < 18; fi217++ {
                var v218 bool = true
                for fi219 := 0; fi219 < 18; fi219++ {
                    var v220 int64 = 9223372036854775807
                    var v221 string = "a"
                    var v222 bool = (v55 <= 307)
                    _ = fi219
                    _ = v220
                    _ = v221
                    _ = v222
                }
                _ = fi217
                _ = v218
            }
            var v223 int32 = -97
            _ = v223
        }
    } else {
        v55 = (^(v55))
    }
    var v224 float64 = (0.0 / 1.0)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v224:", v224)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v54:", len(v54), v54)
    fmt.Println("v55:", v55)
    _ = v1
    _ = v2
    _ = v224
    _ = v3
    _ = v54
    _ = v55
    _ = v57
    _ = v58
}
