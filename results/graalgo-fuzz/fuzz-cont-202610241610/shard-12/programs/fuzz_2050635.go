package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int64; F1 uint64; F2 int64 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 uint64; F1 bool }
type GenS7 struct { F0 int; F1 uint32 }
type GenS8 struct { F0 uint16; F1 int16; F2 uint64 }
type GenS9 struct { F0 string; F1 bool; F2 int64 }
type GenS10 struct { F0 uint16; F1 int64 }
type GenS11 struct { F0 int64; F1 uint16 }
type GenS12 struct { F0 int; F1 uint8 }
type GenS13 struct { F0 uint64 }
type GenS14 struct { F0 int32; F1 int64; F2 int }
func genFn0(p0 bool) int32 {
    var v5 int32 = -1
    fmt.Println(p0)
    if p0 {
        fmt.Println(p0, p0)
        fmt.Println(p0)
    } else {
        p0 = (65535 < 2)
        v6 := []int{32767}
        if ((p0 || (!true)) && (!p0)) {
            switch "foo bar" {
            case "":
                var v7 int64 = 9223372036854775807
                _ = v7
            case "a":
                switch -2147483648 {
                case 396:
                    for fi8 := 0; fi8 < 12; fi8++ {
                        p0 = p0
                        _ = fi8
                    }
                case 1:
                    for fi9 := 0; fi9 < 12; fi9++ {
                        var v10 uint = 165
                        switch fi9 {
                        case 0:
                            switch fi9 {
                            case 128:
                                for fi11 := 0; fi11 < 13; fi11++ {
                                    p0 = p0
                                    v12 := GenS1{F0: 217, F1: 42}
                                    fmt.Println(p0, v10)
                                    _ = fi11
                                    _ = v12
                                }
                            default:
                                v5 = -1
                            }
                        case -194:
                            v10 = v10
                        }
                        if len(v6) > 0 {
                            v6[0] = fi9
                        }
                        _ = fi9
                        _ = v10
                    }
                }
            }
        } else {
            p0 = (0.0 <= float32((0.1 * float64((1e10 + (1.5 / 0.1))))))
            var v13 uint64 = 281
            var v14 uint = 0
            _ = v13
            _ = v14
        }
        _ = v6
    }
    _ = p0
    _ = v5
    return v5
}

func main() {
    v1 := make(map[string]bool)
    v2 := []uint32{65535, 65535, 651}
    var v3 int32 = 42
    var v4 uint32 = 2
    _ = len(v2)
    if (!true) {
        v3 = v3
        var v15 uint = 65535
        if true {
            if false {
                var v16 int32 = (v3 + v3)
                for fi17 := 0; fi17 < 2; fi17++ {
                    switch ("\n\t" + "foo bar") {
                    case "foo bar":
                        if true {
                            if false {
                                fmt.Println(v15, fi17)
                            }
                        } else {
                            var v18 bool = true
                            v19 := GenS1{F0: 0, F1: 42}
                            fmt.Println(v15, v15, v3)
                            _ = v18
                            _ = v19
                        }
                    }
                    v20 := GenS2{F0: 924}
                    if true {
                        v4 = v4
                        v16 = (v3 - int32(v4))
                    } else {
                        v16 -= (v16 / ((v16) | 1))
                        switch "hello" {
                        case "":
                            if (false || true) {
                                if len(v2) > 0 {
                                    v2[0] = 42
                                }
                                if true {
                                    for fi21 := 0; fi21 < 20; fi21++ {
                                        fmt.Println(fi21, v16, v16)
                                        _ = fi21
                                    }
                                }
                                switch "" {
                                case "foo bar":
                                    for fi22 := 0; fi22 < 11; fi22++ {
                                        v15 += (v15 << (((v15 << (((v15 ^ (v15 * (^(v15))))) & 63))) & 63))
                                        var v23 string = "hello"
                                        _ = len(v1)
                                        var v24 uint32 = v4
                                        var v25 uint8 = 1
                                        _ = fi22
                                        _ = v23
                                        _ = v24
                                        _ = v25
                                    }
                                case "":
                                    for fi26 := 0; fi26 < 18; fi26++ {
                                        fmt.Println(v4, v4, v4)
                                        _ = fi26
                                    }
                                case "hello":
                                    switch ((("foo bar" + "a") + "foo bar") + ("hello" + "hello")) {
                                    case "\n\t":
                                        var v27 float64 = 1.5
                                        _ = v27
                                    case "foo bar":
                                        v28 := GenS3{F0: -1, F1: 622, F2: 766}
                                        _ = v28
                                    default:
                                        v29 := GenS3{F0: 0, F1: 4294967295, F2: -9223372036854775808}
                                        _ = v29
                                    }
                                }
                            }
                        case "hello":
                            var v30 int64 = int64(v15)
                            _ = v30
                        case "foo bar":
                            v31 := GenS4{F0: 128}
                            _ = v31
                        }
                    }
                    _ = fi17
                    _ = v20
                }
                switch "" {
                case "a":
                    switch ((("a" + ((("\n\t" + ("\n\t" + "foo bar")) + ("" + "hello")) + "\n\t")) + "hello") + "\n\t") {
                    case "hello":
                        var v32 bool = true
                        _ = v32
                    case "":
                        _ = len(v1)
                    case "foo bar":
                        var v33 int32 = (^(v16))
                        _ = v33
                    }
                case "\n\t":
                    var v34 string = (("a" + ("\n\t" + ("a" + ""))) + ((("foo bar" + "a") + ("\n\t" + ("hello" + ("\n\t" + ("\n\t" + "a"))))) + "a"))
                    _ = v34
                case "hello":
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                }
                _ = v16
            } else {
                v3 += (-(v3))
                switch "a" {
                case "a":
                    var v35 uint64 = 4294967295
                    _ = v35
                }
                v36 := GenS5{F0: -9223372036854775808}
                _ = v36
            }
            for fi37 := 0; fi37 < 14; fi37++ {
                fmt.Println(v15)
                fmt.Println(v4)
                _ = fi37
            }
            var v38 bool = true
            _ = v38
        } else {
            switch "\n\t" {
            case "":
                if true {
                    if true {
                        fmt.Println(v15, v4)
                        var v39 uint8 = 120
                        _ = v39
                    }
                    v40 := GenS6{F0: 0, F1: ((v4 <= (v4 + v4)) || false)}
                    _ = v40
                } else {
                    v15 = (v15 >> ((255) & 63))
                    var v41 int8 = 127
                    _ = v41
                }
            case "hello":
                if false {
                    fmt.Println(v4)
                } else {
                    v4 = (v4 & v4)
                    fmt.Println(v4, v3)
                    for fi42 := 0; fi42 < 6; fi42++ {
                        if (!true) {
                            for fi43 := 0; fi43 < 6; fi43++ {
                                if false {
                                    delete(v1, "a")
                                    v15 += 234
                                } else {
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    var v44 uint = (v15 * (v15 | (v15 / (((^(v15))) | 1))))
                                    var v45 int32 = 0
                                    _ = v44
                                    _ = v45
                                }
                                _ = fi43
                            }
                            var v46 uint = v15
                            v4 += v4
                            _ = v46
                        } else {
                            var v47 uint32 = v4
                            if ((v3 << (((v15 >> ((v15) & 63))) & 63)) <= (v3 ^ -1)) {
                                for fi48 := 0; fi48 < 19; fi48++ {
                                    if false {
                                        fmt.Println(v4)
                                        var v49 uint16 = 2
                                        _ = v49
                                    }
                                    if (((!true) || true) && false) {
                                        delete(v1, "hello")
                                        v15 = 2
                                    } else {
                                        fmt.Println(v3, v4)
                                        v50 := GenS3{F0: 615, F1: 79, F2: -1}
                                        _ = v50
                                    }
                                    v51 := GenS2{F0: 386}
                                    _ = fi48
                                    _ = v51
                                }
                                v3 = 1
                            } else {
                                if (-18 >= 102) {
                                    for fi52 := 0; fi52 < 10; fi52++ {
                                        var v53 string = "a"
                                        _ = fi52
                                        _ = v53
                                    }
                                    for fi54 := 0; fi54 < 20; fi54++ {
                                        fmt.Println(fi42, v3)
                                        v3 += -2147483648
                                        v4 -= (^(v4))
                                        _ = fi54
                                    }
                                } else {
                                    if (!(true || true)) {
                                        for fi55 := 0; fi55 < 13; fi55++ {
                                            if false {
                                                if false {
                                                    if true {
                                                        var v56 uint8 = 0
                                                        v47 += (v4 * v4)
                                                        _ = v56
                                                    } else {
                                                        v57 := GenS4{F0: 65}
                                                        var v58 uint16 = uint16(v15)
                                                        v47 = (v4 << (((v15 | v15)) & 63))
                                                        _ = v57
                                                        _ = v58
                                                    }
                                                    var v59 int16 = 1
                                                    v3 = (-(v3))
                                                    _ = v59
                                                }
                                                if (9223372036854775807 < 1) {
                                                    v4 = 42
                                                }
                                                if (0 == 1) {
                                                    var v60 string = ""
                                                    _ = v60
                                                } else {
                                                    var v61 uint8 = 42
                                                    _ = v61
                                                }
                                                var v62 uint16 = 415
                                                var v63 int8 = 42
                                                _ = v62
                                                _ = v63
                                            }
                                            _ = fi55
                                        }
                                        if (false || (((v3 % ((2147483647) | 1)) >= genFn0((-81 != -1))) && false)) {
                                            v64 := GenS2{F0: 42}
                                            if ((true && false) || true) {
                                                v3 += (v3 << ((v15) & 63))
                                                _ = len(v2)
                                            }
                                            v65 := GenS1{F0: 42, F1: 42}
                                            _ = v64
                                            _ = v65
                                        } else {
                                            var v66 uint32 = v47
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                            var v67 string = (("a" + "") + "")
                                            _ = v66
                                            _ = v67
                                        }
                                    }
                                    v2 = append(v2, (v4 ^ uint32(v15)))
                                }
                            }
                            _ = v47
                        }
                        _ = fi42
                    }
                }
            case "foo bar":
                v4 -= uint32(v15)
            }
            switch "hello" {
            case "a":
                var v68 uint = (v15 + (v15 ^ v15))
                _ = v68
            case "foo bar":
                v3 = v3
            default:
                if (!true) {
                    v15 = (v15 * (v15 >> (((v15 ^ 4294967295)) & 63)))
                    if true {
                        for fi69 := 0; fi69 < 12; fi69++ {
                            if (128 <= 0) {
                                fmt.Println(v15, v4, v3)
                                switch 0 {
                                case 42:
                                    var v70 uint64 = 0
                                    _ = v70
                                case -9223372036854775808:
                                    for fi71 := 0; fi71 < 7; fi71++ {
                                        var v72 uint8 = 221
                                        fmt.Println(v72, v15)
                                        _ = fi71
                                        _ = v72
                                    }
                                case 128:
                                    var v73 float32 = float32((1e10 * 1e10))
                                    _ = v73
                                }
                                fmt.Println(v4)
                            } else {
                                var v74 uint = v15
                                if false {
                                    v15 -= (v74 * v15)
                                    v15 = (v74 * v74)
                                    var v75 bool = (1 <= 127)
                                    _ = v75
                                }
                                _ = v74
                            }
                            var v76 int32 = (v3 & v3)
                            if (false && false) {
                                if ("a" == "a") {
                                    v77 := GenS4{F0: 0}
                                    _ = v77
                                }
                            } else {
                                for fi78 := 0; fi78 < 18; fi78++ {
                                    delete(v1, "\n\t")
                                    switch int(v15) {
                                    case 476:
                                        _ = len(v2)
                                    case 127:
                                        v76 = (v76 & v3)
                                    }
                                    v79 := GenS3{F0: 42, F1: 106, F2: 9223372036854775807}
                                    _ = fi78
                                    _ = v79
                                }
                                v80 := GenS7{F0: (fi69 >> ((2) & 63)), F1: 179}
                                fmt.Println(v15)
                                _ = v80
                            }
                            _ = fi69
                            _ = v76
                        }
                        _ = len(v2)
                        v4 = 0
                    } else {
                        v4 -= 65535
                        v81 := v1[""]
                        _ = v81
                    }
                    var v82 uint = (v15 | v15)
                    _ = v82
                } else {
                    switch -125 {
                    case 127:
                        v4 -= v4
                    case 2147483647:
                        var v83 uint = (v15 + v15)
                        _ = v83
                    case -409:
                        var v84 int8 = 5
                        _ = v84
                    }
                    switch "\n\t" {
                    case "":
                        switch "" {
                        case "hello":
                            v4 = v4
                        case "\n\t":
                            var v85 string = (("a" + ("" + (("foo bar" + "hello") + "hello"))) + "\n\t")
                            _ = v85
                        }
                    case "hello":
                        v86 := GenS1{F0: 9223372036854775807, F1: -744}
                        _ = v86
                    case "\n\t":
                        if false {
                            v4 = (v4 - 2147483647)
                            if (-41 <= 1) {
                                v15 += v15
                            } else {
                                switch "a" {
                                case "foo bar":
                                    v15 = 18446744073709551615
                                case "a":
                                    v1["hello"] = (((true && false) && false) || ((!(((true && ((3.14159 * ((-2.25 + -2.25) / 0.1)) < float32(0.1))) || (false || (!(v15 > 2)))) && (true && false))) || false))
                                default:
                                    var v87 int = 795
                                    _ = v87
                                }
                                var v88 float64 = 1.5
                                switch -9223372036854775808 {
                                case -228:
                                    switch -1 {
                                    case 301:
                                        v89 := GenS8{F0: uint16(v15), F1: 0, F2: 0}
                                        _ = v89
                                    case 9223372036854775807:
                                        if false {
                                            v90 := GenS9{F0: ("hello" + "a"), F1: (!false), F2: 1}
                                            var v91 float64 = (v88 * (0.0 / 1.0))
                                            var v92 string = ((("hello" + "") + "\n\t") + "foo bar")
                                            _ = v90
                                            _ = v91
                                            _ = v92
                                        } else {
                                            var v93 uint64 = 192
                                            if ((295 >= 794) || true) {
                                                for fi94 := 0; fi94 < 8; fi94++ {
                                                    for fi95 := 0; fi95 < 2; fi95++ {
                                                        v96 := GenS6{F0: v93, F1: true}
                                                        if len(v2) >= 2 {
                                                            v2 = v2[1:len(v2)]
                                                        }
                                                        if (249 >= 127) {
                                                            var v97 int = -64
                                                            switch v97 {
                                                            case -1:
                                                                v97 += fi95
                                                            case 9223372036854775807:
                                                                v98 := GenS2{F0: v93}
                                                                _ = v98
                                                            }
                                                            v99 := GenS10{F0: 0, F1: 42}
                                                            _ = v97
                                                            _ = v99
                                                        }
                                                        var v100 int8 = -47
                                                        if len(v2) >= 2 {
                                                            v2 = v2[1:len(v2)]
                                                        }
                                                        v3 += genFn0(true)
                                                        _ = fi95
                                                        _ = v100
                                                        _ = v96
                                                    }
                                                    v93 = (v93 * (^(v93)))
                                                    fmt.Println(v4)
                                                    _ = fi94
                                                }
                                            }
                                            _ = v93
                                        }
                                    default:
                                        v3 = v3
                                    }
                                }
                                _ = v88
                            }
                        } else {
                            v101 := GenS11{F0: 42, F1: 255}
                            _ = v101
                        }
                    default:
                        var v102 float64 = 1.5
                        _ = v102
                    }
                }
            }
            switch 4294967295 {
            case 128:
                v103 := GenS12{F0: -737, F1: 42}
                _ = v103
            case 32768:
                if ("hello" == "hello") {
                    switch "a" {
                    case "":
                        v3 = 2147483647
                    default:
                        v15 -= 1
                    }
                    if true {
                        v104 := GenS10{F0: 0, F1: -9223372036854775808}
                        for fi105 := 0; fi105 < 18; fi105++ {
                            if false {
                                switch (fi105 + fi105) {
                                case 776:
                                    _ = len(v1)
                                case -300:
                                    v4 -= v4
                                case 9223372036854775807:
                                    var v106 uint8 = 0
                                    _ = v106
                                default:
                                    for fi107 := 0; fi107 < 4; fi107++ {
                                        if false {
                                            var v108 int = 0
                                            _ = v108
                                        }
                                        v109 := GenS3{F0: 0, F1: 4294967295, F2: -1}
                                        _ = fi107
                                        _ = v109
                                    }
                                }
                                v15 = (v15 + 983)
                            }
                            fmt.Println(v3, v15)
                            _ = fi105
                        }
                        _ = v104
                    }
                }
            }
        }
        _ = v15
    } else {
        var v110 int16 = 0
        _ = v110
    }
    fmt.Println(v4, v3)
    if (true || true) {
        for fi111 := 0; fi111 < 2; fi111++ {
            for fi112 := 0; fi112 < 4; fi112++ {
                v3 -= v3
                v113 := GenS13{F0: 786}
                v3 -= (v3 | 1)
                _ = fi112
                _ = v113
            }
            _ = fi111
        }
        switch ("a" + ("foo bar" + "\n\t")) {
        case "":
            v4 = (v4 ^ (v4 << ((881) & 63)))
        case "a":
            var v114 int32 = (^(v3))
            _ = v114
        default:
            for fi115 := 0; fi115 < 11; fi115++ {
                v1["foo bar"] = (9223372036854775807 == -9223372036854775808)
                _ = fi115
            }
        }
    } else {
        if (!(true || true)) {
            var v116 float64 = 1e10
            _ = v116
        }
    }
    var v117 int16 = -1
    for fi118 := 0; fi118 < 9; fi118++ {
        switch fi118 {
        case -2147483648:
            switch fi118 {
            case 128:
                if len(v2) > 0 {
                    v2[0] = v4
                }
            case 32767:
                v119 := GenS5{F0: fi118}
                _ = v119
            }
        case -388:
            for fi120 := 0; fi120 < 20; fi120++ {
                if (!false) {
                    v2 = append(v2, v4)
                    if len(v2) > 0 {
                        v2[0] = v4
                    }
                    v117 = -1
                }
                v117 = (-(v117))
                v1["a"] = false
                var v121 uint64 = 42
                _ = fi120
                _ = v121
            }
        case 42:
            v117 = int16(v3)
        default:
            if (-32768 <= v117) {
                fmt.Println(v4, v3)
                var v122 uint8 = 127
                _ = v122
            } else {
                if true {
                    switch ("" + "a") {
                    case "":
                        var v123 int8 = -86
                        _ = v123
                    case "a":
                        if (9223372036854775807 < -9223372036854775808) {
                            if len(v2) > 0 {
                                v2[0] = v4
                            }
                        }
                    case "foo bar":
                        if true {
                            switch (fi118 / ((-1) | 1)) {
                            case 9223372036854775807:
                                var v124 bool = true
                                _ = v124
                            case -985:
                                var v125 uint32 = v4
                                _ = v125
                            }
                            var v126 uint32 = v4
                            _ = v126
                        }
                    default:
                        for fi127 := 0; fi127 < 9; fi127++ {
                            if true {
                                fmt.Println(v3)
                                if (!(9223372036854775807 <= 0)) {
                                    switch "foo bar" {
                                    case "foo bar":
                                        v4 = (v4 - v4)
                                    case "":
                                        if false {
                                            v128 := GenS5{F0: fi118}
                                            var v129 int = 4294967295
                                            _ = v128
                                            _ = v129
                                        } else {
                                            v3 -= (v3 ^ 42)
                                            var v130 uint8 = 42
                                            _ = v130
                                        }
                                    }
                                    switch 127 {
                                    case 2147483647:
                                        v3 = 631
                                    }
                                    var v131 bool = (!false)
                                    _ = v131
                                }
                            }
                            var v132 int64 = 199
                            _ = fi127
                            _ = v132
                        }
                    }
                } else {
                    var v133 bool = (0 <= 32767)
                    for fi134 := 0; fi134 < 4; fi134++ {
                        v135 := GenS3{F0: 1, F1: 2, F2: -238}
                        if v133 {
                            v136 := GenS14{F0: 2147483647, F1: int64(v4), F2: fi134}
                            _ = v136
                        } else {
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            if ((!v133) && v133) {
                                v3 -= v3
                            }
                            var v137 int8 = 55
                            _ = v137
                        }
                        _ = fi134
                        _ = v135
                    }
                    _ = v133
                }
                v3 = -1
            }
        }
        if ((float32((((3.14159 + 3.14159) * 1e10) - 0.1)) * 0.1) == (-2.25 + 0.0)) {
            var v138 float32 = -2.25
            delete(v1, "\n\t")
            _ = v138
        }
        v3 = (v3 | v3)
        _ = fi118
    }
    var v139 int = -414
    _ = len(v1)
    v4 = (v4 << ((2) & 63))
    v4 -= 4294967295
    var v140 float32 = float32((0.1 / (((1.5 * 1e10) / -2.25) - 1e10)))
    v140 -= v140
    if true {
        v3 += 378
        var v141 uint8 = 127
        _ = v141
    }
    switch v139 {
    case -9223372036854775808:
        var v142 float64 = 0.1
        _ = v142
    }
    fmt.Println(v4, v139, v3)
    var v143 uint64 = 18446744073709551615
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v117:", v117)
    fmt.Println("v139:", v139)
    fmt.Println("v140:", v140)
    fmt.Println("v143:", v143)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v117
    _ = v139
    _ = v140
    _ = v143
    _ = v2
    _ = v3
    _ = v4
}
