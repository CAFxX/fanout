package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint }
type GenS2 struct { F0 int64; F1 int32 }
type GenS3 struct { F0 int; F1 int }
type GenS4 struct { F0 int64; F1 bool }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 int32; F1 int64; F2 uint64 }
type GenS7 struct { F0 uint32; F1 int16; F2 uint16 }
type GenS8 struct { F0 int; F1 uint16; F2 uint8 }
type GenS9 struct { F0 bool; F1 string }
type GenS10 struct { F0 uint8 }
type GenS11 struct { F0 string }
type GenS12 struct { F0 int16; F1 string; F2 int8 }
type GenS13 struct { F0 string; F1 int16; F2 uint64 }
type GenS14 struct { F0 uint32; F1 bool }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[string]uint32)
    var v3 int16 = 32767
    v4 := make(map[uint32]uint32)
    fmt.Println(v3)
    var v5 uint32 = 218
    for fi6 := 0; fi6 < 15; fi6++ {
        if false {
            if (int16(v5) >= v3) {
                var v7 uint8 = uint8(v5)
                _ = v7
            }
            delete(v4, (v5 & (v5 & (v5 + (v5 & 0)))))
            if false {
                v8 := []int{(^(fi6))}
                _ = v8
            } else {
                v9 := []int{}
                v10 := GenS1{F0: 0, F1: 18446744073709551615}
                _ = v10
                _ = v9
            }
        } else {
            var v11 uint16 = 8
            switch 460 {
            case 0:
                if (!true) {
                    for fi12 := 0; fi12 < 8; fi12++ {
                        v5 -= 375
                        switch "foo bar" {
                        case "foo bar":
                            var v13 uint8 = 120
                            _ = v13
                        default:
                            var v14 uint8 = 2
                            _ = v14
                        }
                        if true {
                            v4[v5] = v5
                            switch "" {
                            case "a":
                                fmt.Println(fi6)
                            case "":
                                v11 = 1
                            default:
                                var v15 float64 = float64(float32(1.5))
                                _ = v15
                            }
                        }
                        _ = fi12
                    }
                    fmt.Println(v5, fi6, v3)
                    fmt.Println(v5)
                }
            }
            if false {
                switch fi6 {
                case 42:
                    switch ("\n\t" + (("a" + "\n\t") + "foo bar")) {
                    case "":
                        switch "a" {
                        case "foo bar":
                            var v16 int = fi6
                            _ = v16
                        default:
                            v17 := GenS1{F0: int64(v11), F1: 4294967295}
                            _ = v17
                        }
                    case "foo bar":
                        _ = len(v2)
                    }
                case -9223372036854775808:
                    for fi18 := 0; fi18 < 1; fi18++ {
                        v11 += 935
                        if (!false) {
                            v19 := v4[(v5 * (v5 / ((785) | 1)))]
                            v3 -= (v3 >> ((821) & 63))
                            _ = v19
                        } else {
                            fmt.Println(fi6, fi6)
                            if ((v5 >> ((324) & 63)) >= 1) {
                                v11 = (^(v11))
                                v5 = 65535
                                v5 += (v5 << ((4294967295) & 63))
                            } else {
                                v11 = v11
                                v20 := GenS1{F0: 1, F1: 42}
                                switch (fi6 << ((uint(v5)) & 63)) {
                                case 42:
                                    v21 := GenS2{F0: -346, F1: 913}
                                    _ = v21
                                case 127:
                                    if true {
                                        v22 := GenS1{F0: 1, F1: 0}
                                        v23 := GenS2{F0: 9223372036854775807, F1: 2147483647}
                                        var v24 uint8 = 204
                                        _ = v22
                                        _ = v23
                                        _ = v24
                                    }
                                default:
                                    v11 -= v11
                                }
                                _ = v20
                            }
                        }
                        switch (fi18 + fi18) {
                        case 0:
                            if (1 == 255) {
                                if false {
                                    var v25 uint64 = 671
                                    v26 := GenS2{F0: -1, F1: 0}
                                    v27 := GenS1{F0: 0, F1: 255}
                                    _ = v25
                                    _ = v26
                                    _ = v27
                                } else {
                                    v28 := []int{32768, int(v3), -712, fi18}
                                    _ = v28
                                }
                                delete(v1, (v5 ^ 2147483647))
                            } else {
                                v29 := GenS1{F0: 0, F1: 65535}
                                var v30 string = ""
                                switch (-(fi18)) {
                                case 42:
                                    var v31 float64 = (0.1 / -2.25)
                                    _ = v31
                                case 1:
                                    v4[v5] = (v5 ^ v5)
                                default:
                                    var v32 int8 = -73
                                    _ = v32
                                }
                                _ = v29
                                _ = v30
                            }
                        case 4294967295:
                            v33 := GenS3{F0: (fi18 << ((uint(v11)) & 63)), F1: 42}
                            _ = v33
                        }
                        _ = fi18
                    }
                case 128:
                    fmt.Println(v5, v5, fi6)
                default:
                    v34 := v2["hello"]
                    _ = v34
                }
            } else {
                fmt.Println(v3, v11)
            }
            _ = v11
        }
        delete(v4, 4294967295)
        _ = fi6
    }
    fmt.Println(v5, v3)
    var v35 uint64 = 1
    fmt.Println(v35, v5)
    for fi36 := 0; fi36 < 16; fi36++ {
        if true {
            var v37 float32 = 0.1
            for fi38 := 0; fi38 < 10; fi38++ {
                var v39 uint = 0
                _ = fi38
                _ = v39
            }
            for fi40 := 0; fi40 < 12; fi40++ {
                v41 := GenS1{F0: 42, F1: 710}
                if ((2147483648 > v5) || (((false || (-85 <= 42)) || true) || true)) {
                    var v42 bool = (false && false)
                    _ = v42
                } else {
                    var v43 string = "\n\t"
                    var v44 uint = 58
                    v45 := GenS4{F0: 0, F1: false}
                    _ = v43
                    _ = v44
                    _ = v45
                }
                v3 = -97
                _ = fi40
                _ = v41
            }
            _ = v37
        } else {
            if true {
                v46 := GenS2{F0: 1, F1: 1}
                _ = v46
            } else {
                if ("hello" >= (("foo bar" + ("hello" + "hello")) + "")) {
                    _ = len(v2)
                    if ((1e10 + float32(float64(-2.25))) == 1.5) {
                        fmt.Println(v35, fi36)
                        for fi47 := 0; fi47 < 13; fi47++ {
                            v35 = 4294967295
                            switch "hello" {
                            case "a":
                                switch (-(fi36)) {
                                case -2147483648:
                                    fmt.Println(v3, fi36, v3)
                                case 32767:
                                    switch (^(fi47)) {
                                    case -9223372036854775808:
                                        var v48 string = "foo bar"
                                        _ = v48
                                    }
                                case 9223372036854775807:
                                    v3 += v3
                                }
                            case "foo bar":
                                switch (((("a" + ("a" + "a")) + "foo bar") + "foo bar") + "hello") {
                                case "foo bar":
                                    if (uint16(v5) > 65535) {
                                        if (-49 > -9223372036854775808) {
                                            switch ("" + "a") {
                                            case "\n\t":
                                                var v49 int16 = v3
                                                _ = v49
                                            case "":
                                                delete(v4, 1)
                                                fmt.Println(v3)
                                                v50 := []int{2147483647, (fi36 << ((65535) & 63))}
                                                _ = v50
                                            }
                                        }
                                        fmt.Println(v35, v35, v3)
                                        var v51 uint = uint(v35)
                                        _ = v51
                                    }
                                case "hello":
                                    var v52 int64 = -1
                                    _ = v52
                                case "a":
                                    v53 := GenS5{F0: v5}
                                    _ = v53
                                }
                            default:
                                if ("" <= ("\n\t" + "")) {
                                    var v54 bool = ((^(v5)) <= v5)
                                    fmt.Println(fi36)
                                    v55 := GenS6{F0: int32(v5), F1: 335, F2: (v35 | v35)}
                                    v3 = (v3 - -187)
                                    v56 := GenS7{F0: uint32(v35), F1: 1000, F2: 255}
                                    _ = v54
                                    _ = v55
                                    _ = v56
                                }
                            }
                            _ = fi47
                        }
                        var v57 int = fi36
                        _ = v57
                    }
                } else {
                    v5 = (v5 ^ v5)
                }
                var v58 int8 = -38
                _ = v58
            }
            var v59 int = fi36
            _ = v59
        }
        var v60 uint8 = 128
        v35 += v35
        _ = fi36
        _ = v60
    }
    for fi61 := 0; fi61 < 20; fi61++ {
        v5 = uint32(v35)
        var v62 uint32 = v5
        v35 -= (v35 << ((0) & 63))
        _ = fi61
        _ = v62
    }
    if true {
        switch 128 {
        case 128:
            v63 := GenS1{F0: 344, F1: 962}
            _ = v63
        case 9223372036854775807:
            var v64 int8 = -20
            _ = v64
        }
        switch -9223372036854775808 {
        case -2147483648:
            v35 -= 320
        case -388:
            v3 -= (v3 * (-(v3)))
        }
        v65 := GenS7{F0: v5, F1: -1, F2: 255}
        _ = v65
    } else {
        fmt.Println(v35, v3, v3)
    }
    switch -267 {
    case 128:
        if true {
            for fi66 := 0; fi66 < 13; fi66++ {
                switch (-(fi66)) {
                case 771:
                    if true {
                        if false {
                            v67 := []int{int(v3)}
                            for fi68 := 0; fi68 < 10; fi68++ {
                                v3 += -32768
                                var v69 bool = (!true)
                                _ = fi68
                                _ = v69
                            }
                            _ = v67
                        }
                        v70 := []int{fi66}
                        var v71 int8 = 127
                        _ = v70
                        _ = v71
                    }
                case 519:
                    fmt.Println(v35)
                }
                _ = fi66
            }
            if (false && (v5 >= 65535)) {
                switch 127 {
                case -9223372036854775808:
                    v35 -= v35
                case 2147483647:
                    for fi72 := 0; fi72 < 17; fi72++ {
                        var v73 uint32 = 270
                        if false {
                            var v74 uint8 = 2
                            for fi75 := 0; fi75 < 17; fi75++ {
                                var v76 float64 = (1.5 - 0.0)
                                if (int16(v35) >= v3) {
                                    var v77 int = (^(fi75))
                                    fmt.Println(v35, fi72, v76)
                                    delete(v4, v5)
                                    if (!(int32(v3) >= -2147483648)) {
                                        v74 -= (^(v74))
                                    } else {
                                        delete(v2, ("" + "hello"))
                                        if false {
                                            fmt.Println(fi72, v35, fi75)
                                            delete(v2, "hello")
                                            v78 := GenS6{F0: 0, F1: 1, F2: v35}
                                            if (!((((0.1 * (-2.25 / 1.0)) < -2.25) || (405 > 32767)) && false)) {
                                                v79 := []int{fi72, (fi72 ^ (-(fi72)))}
                                                var v80 int = (fi72 & (fi75 | (fi72 >> ((2) & 63))))
                                                _ = v79
                                                _ = v80
                                            }
                                            if false {
                                                v81 := GenS8{F0: (fi72 & v77), F1: 65535, F2: (v74 % ((v74) | 1))}
                                                v82 := []int{}
                                                v83 := GenS8{F0: (fi72 + (-(v77))), F1: 2, F2: v74}
                                                v3 += 1000
                                                v74 += 0
                                                if true {
                                                    fmt.Println(v35, v3)
                                                    if false {
                                                        var v84 bool = (0.1 >= v76)
                                                        fmt.Println(v5, v5, fi75)
                                                        var v85 uint32 = (v73 / ((uint32(v35)) | 1))
                                                        var v86 bool = v84
                                                        _ = v84
                                                        _ = v85
                                                        _ = v86
                                                    } else {
                                                        v87 := GenS8{F0: -9223372036854775808, F1: 65535, F2: 0}
                                                        fmt.Println(v73)
                                                        _ = v87
                                                    }
                                                    var v88 uint32 = v5
                                                    _ = v88
                                                }
                                                var v89 string = "a"
                                                var v90 int16 = 422
                                                v91 := GenS9{F0: (!((32767 >= 127) && ((1.5 * (float64(0.1) - (v76 + v76))) != (((v76 * v76) - (1.5 * 0.1)) - 1e10)))), F1: ((v89 + v89) + v89)}
                                                v92 := GenS10{F0: (^(v74))}
                                                _ = v81
                                                _ = v82
                                                _ = v83
                                                _ = v89
                                                _ = v90
                                                _ = v91
                                                _ = v92
                                            }
                                            _ = v78
                                        }
                                        fmt.Println(v3, v73, v74)
                                    }
                                    _ = v77
                                }
                                _ = fi75
                                _ = v76
                            }
                            _ = v74
                        } else {
                            v73 = (v5 - v73)
                            var v93 float32 = -2.25
                            for fi94 := 0; fi94 < 11; fi94++ {
                                delete(v1, v73)
                                switch (fi94 & fi94) {
                                case 127:
                                    var v95 uint16 = uint16(v5)
                                    _ = v95
                                case 9223372036854775807:
                                    fmt.Println(v73, v3, fi72)
                                    v96 := GenS1{F0: -527, F1: 1}
                                    _ = v96
                                case -9223372036854775808:
                                    var v97 float64 = ((0.0 * 0.1) * 0.0)
                                    _ = v97
                                }
                                _ = fi94
                            }
                            _ = v93
                        }
                        v35 += (v35 - v35)
                        _ = fi72
                        _ = v73
                    }
                case 4294967295:
                    v98 := []int{42, -329, -2147483648, 9223372036854775807}
                    _ = v98
                default:
                    v99 := GenS11{F0: (("foo bar" + ("\n\t" + "foo bar")) + (("a" + "hello") + ("a" + "a")))}
                    _ = v99
                }
            }
        }
    case -1:
        var v100 uint64 = v35
        _ = v100
    case -673:
        v5 += v5
    default:
        if ("" < "") {
            switch 0 {
            case 128:
                v35 = (v35 << ((0) & 63))
            case 32767:
                switch 2147483647 {
                case -821:
                    v101 := GenS12{F0: 42, F1: "a", F2: 85}
                    _ = v101
                case 128:
                    for fi102 := 0; fi102 < 3; fi102++ {
                        v103 := GenS13{F0: ("a" + "a"), F1: v3, F2: v35}
                        _ = fi102
                        _ = v103
                    }
                case -35:
                    for fi104 := 0; fi104 < 8; fi104++ {
                        for fi105 := 0; fi105 < 5; fi105++ {
                            delete(v1, 2147483648)
                            fmt.Println(v5, v35, fi105)
                            v35 = v35
                            _ = fi105
                        }
                        var v106 int8 = 1
                        _ = fi104
                        _ = v106
                    }
                default:
                    v5 = 928
                }
            case 0:
                v35 -= (v35 / ((v35) | 1))
            default:
                for fi107 := 0; fi107 < 4; fi107++ {
                    var v108 int16 = v3
                    _ = fi107
                    _ = v108
                }
            }
            v5 = v5
        }
    }
    fmt.Println(v3, v5, v35)
    if false {
        if (!(v5 < v5)) {
            for fi109 := 0; fi109 < 2; fi109++ {
                v5 = 0
                if true {
                    v35 += (v35 - 1)
                    v35 -= (v35 - (v35 << ((21) & 63)))
                    var v110 uint = 1
                    _ = v110
                } else {
                    delete(v2, ("foo bar" + (((("a" + "foo bar") + "") + "foo bar") + "")))
                }
                _ = fi109
            }
            v3 += (v3 << ((659) & 63))
        } else {
            v3 = 1
        }
        v111 := GenS9{F0: (127 > 42), F1: ("\n\t" + "hello")}
        _ = v111
    } else {
        if true {
            v5 += (^(v5))
        }
        if false {
            var v112 string = (("foo bar" + "a") + "foo bar")
            _ = v112
        }
    }
    var v113 bool = false
    if ((-2.25 * 3.14159) >= 0.0) {
        v113 = v113
        if ((3.14159 / 0.1) < ((0.1 * (1e10 - (1.5 + 3.14159))) - 3.14159)) {
            v114 := []int{-9223372036854775808, -9223372036854775808, -9223372036854775808, -616}
            _ = v114
        }
        if v113 {
            switch "hello" {
            case "\n\t":
                if (1 == 1) {
                    if v113 {
                        v115 := GenS8{F0: 334, F1: 0, F2: 47}
                        _ = v115
                    }
                    for fi116 := 0; fi116 < 7; fi116++ {
                        switch fi116 {
                        case 4294967295:
                            if ((0 != 0) || (!(42 == 274))) {
                                var v117 string = ("a" + "\n\t")
                                v35 -= 42
                                _ = v117
                            }
                        case -213:
                            v118 := GenS10{F0: 255}
                            _ = v118
                        case 1:
                            var v119 bool = (true && (0 <= 42))
                            _ = v119
                        default:
                            var v120 uint32 = (v5 - (v5 % (((^(v5))) | 1)))
                            _ = v120
                        }
                        fmt.Println(v113, v35, v113)
                        _ = fi116
                    }
                    _ = len(v2)
                }
            case "":
                for fi121 := 0; fi121 < 7; fi121++ {
                    switch (fi121 / ((0) | 1)) {
                    case -1:
                        var v122 uint16 = 1
                        _ = v122
                    case 870:
                        v123 := GenS1{F0: -1, F1: 42}
                        _ = v123
                    default:
                        var v124 int = fi121
                        _ = v124
                    }
                    if v113 {
                        for fi125 := 0; fi125 < 7; fi125++ {
                            _ = len(v2)
                            _ = fi125
                        }
                        v5 += 2147483647
                        v126 := GenS1{F0: 775, F1: 574}
                        _ = v126
                    }
                    _ = fi121
                }
            case "a":
                v35 = uint64(v5)
            default:
                v127 := GenS14{F0: v5, F1: (!true)}
                _ = v127
            }
            for fi128 := 0; fi128 < 17; fi128++ {
                v35 += (^(v35))
                _ = fi128
            }
            switch 4294967295 {
            case 2147483647:
                var v129 float64 = ((-2.25 - 1.5) - -2.25)
                _ = v129
            case 621:
                for fi130 := 0; fi130 < 11; fi130++ {
                    switch "\n\t" {
                    case "hello":
                        fmt.Println(v35, fi130, v3)
                    }
                    if v113 {
                        fmt.Println(v113)
                        v35 -= 42
                    } else {
                        var v131 int = (^(fi130))
                        v113 = true
                        _ = v131
                    }
                    fmt.Println(fi130, v5, v35)
                    _ = fi130
                }
            case 9223372036854775807:
                if v113 {
                    var v132 uint = 255
                    _ = v132
                }
            default:
                fmt.Println(v35)
            }
        }
    } else {
        v133 := []int{}
        _ = v133
    }
    var v134 float32 = (((0.0 + 0.1) * 3.14159) - 0.1)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v113:", v113)
    fmt.Println("v134:", v134)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v113
    _ = v134
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v5
}
