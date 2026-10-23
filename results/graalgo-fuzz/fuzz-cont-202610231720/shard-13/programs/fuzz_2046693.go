package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16 }
type GenS2 struct { F0 int; F1 int8 }
type GenS3 struct { F0 int8; F1 int8; F2 uint32 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 string; F1 int64 }
type GenS7 struct { F0 int32; F1 uint16 }
type GenS8 struct { F0 int64 }
type GenS9 struct { F0 string }
type GenS10 struct { F0 string; F1 uint; F2 uint8 }
type GenS11 struct { F0 uint; F1 int; F2 uint8 }
type GenS12 struct { F0 int }
type GenS13 struct { F0 int8; F1 int64; F2 uint }
type GenS14 struct { F0 uint64; F1 int8; F2 uint8 }
type GenS15 struct { F0 int8; F1 string; F2 int8 }
type GenS16 struct { F0 uint16; F1 int16 }
type GenS17 struct { F0 string; F1 uint32 }
type GenS18 struct { F0 int32 }
type GenS19 struct { F0 uint8; F1 uint8 }
type GenS20 struct { F0 uint64; F1 uint32 }
type GenS21 struct { F0 string; F1 uint64; F2 int32 }
type GenS22 struct { F0 int8; F1 int; F2 string }
type GenS23 struct { F0 int32; F1 int }
type GenS24 struct { F0 int64; F1 int32; F2 bool }
type GenS25 struct { F0 int8; F1 int32; F2 uint16 }
func genFn0() int64 {
    if false {
        v5 := []int{-9223372036854775808, 2147483647, -2147483648}
        _ = v5
    }
    if (1.5 != ((3.14159 * -2.25) * 1.5)) {
        v6 := []int{0, 4294967295, 9223372036854775807}
        _ = v6
    } else {
        var v7 uint8 = 1
        var v8 uint32 = 65535
        if true {
            var v9 uint64 = 21
            var v10 uint16 = 330
            fmt.Println(v9)
            _ = v10
            _ = v9
        } else {
            v7 += 128
            for fi11 := 0; fi11 < 10; fi11++ {
                v12 := GenS1{F0: true, F1: 32767}
                switch "hello" {
                case "hello":
                    v8 = 0
                case "foo bar":
                    for fi13 := 0; fi13 < 9; fi13++ {
                        v7 = (v7 & v7)
                        _ = fi13
                    }
                case "":
                    switch "a" {
                    case "foo bar":
                        var v14 int16 = 1
                        _ = v14
                    }
                }
                v15 := GenS2{F0: 4294967295, F1: 42}
                _ = fi11
                _ = v12
                _ = v15
            }
            fmt.Println(v7, v8)
        }
        _ = v7
        _ = v8
    }
    var v16 bool = false
    _ = v16
    return 984
}
func genFn1() int {
    if false {
        switch 0 {
        case 42:
            if (!(127 <= -84)) {
                if (255 > 255) {
                    v17 := []int{-9223372036854775808, 128, -506, 9223372036854775807}
                    switch 102 {
                    case 32767:
                        for fi18 := 0; fi18 < 16; fi18++ {
                            var v19 string = ("\n\t" + ("foo bar" + "\n\t"))
                            v20 := GenS1{F0: false, F1: 42}
                            for fi21 := 0; fi21 < 20; fi21++ {
                                v19 = ""
                                v17 = append(v17, 4294967295)
                                if len(v17) > 0 {
                                    v17[0] = fi21
                                }
                                _ = fi21
                            }
                            _ = fi18
                            _ = v19
                            _ = v20
                        }
                    case -2147483648:
                        var v22 int32 = 2147483647
                        _ = v22
                    case 1:
                        if (!true) {
                            switch 128 {
                            case 741:
                                v23 := GenS3{F0: 23, F1: -1, F2: 0}
                                _ = v23
                            case -1:
                                var v24 int8 = -64
                                _ = v24
                            }
                        } else {
                            for fi25 := 0; fi25 < 15; fi25++ {
                                var v26 uint32 = 975
                                if len(v17) > 0 {
                                    v17[0] = 128
                                }
                                v26 = (v26 / ((691) | 1))
                                _ = fi25
                                _ = v26
                            }
                            var v27 float32 = 0.1
                            v27 -= (float32(1e10) * 3.14159)
                            _ = v27
                        }
                    }
                    _ = v17
                } else {
                    var v28 float32 = 1.5
                    var v29 uint = 303
                    fmt.Println(v28, v28)
                    _ = v28
                    _ = v29
                }
                var v30 uint16 = 2
                _ = v30
            }
        }
        for fi31 := 0; fi31 < 5; fi31++ {
            v32 := GenS4{F0: 2}
            _ = fi31
            _ = v32
        }
    } else {
        var v33 float32 = float32(0.1)
        for fi34 := 0; fi34 < 11; fi34++ {
            var v35 int = fi34
            var v36 bool = true
            _ = fi34
            _ = v35
            _ = v36
        }
        _ = v33
    }
    for fi37 := 0; fi37 < 2; fi37++ {
        var v38 uint8 = 127
        for fi39 := 0; fi39 < 11; fi39++ {
            fmt.Println(fi37)
            _ = fi39
        }
        _ = fi37
        _ = v38
    }
    if true {
        var v40 bool = true
        _ = v40
    } else {
        switch -2147483648 {
        case 42:
            var v41 float64 = 0.0
            _ = v41
        }
        var v42 int = 42
        _ = v42
    }
    var v43 int16 = 32767
    if (-69 >= -1) {
        v44 := GenS2{F0: -2147483648, F1: -1}
        if false {
            if (18446744073709551615 > 4294967295) {
                for fi45 := 0; fi45 < 9; fi45++ {
                    for fi46 := 0; fi46 < 17; fi46++ {
                        v47 := []int{}
                        v48 := GenS5{F0: true}
                        _ = fi46
                        _ = v47
                        _ = v48
                    }
                    _ = fi45
                }
                if ((false || (-589 >= 127)) && true) {
                    if true {
                        if true {
                            v49 := []int{}
                            var v50 uint8 = 254
                            _ = v49
                            _ = v50
                        } else {
                            switch 9223372036854775807 {
                            case -955:
                                var v51 string = "foo bar"
                                _ = v51
                            case -966:
                                v52 := GenS6{F0: "\n\t", F1: 911}
                                _ = v52
                            case -895:
                                if (!true) {
                                    v43 = -1
                                } else {
                                    if false {
                                        switch ("hello" + "a") {
                                        case "foo bar":
                                            v43 = (-(v43))
                                        case "a":
                                            if false {
                                                var v53 int8 = 100
                                                v54 := GenS4{F0: 0}
                                                for fi55 := 0; fi55 < 16; fi55++ {
                                                    v53 = (v53 ^ v53)
                                                    fmt.Println(fi55)
                                                    _ = fi55
                                                }
                                                _ = v53
                                                _ = v54
                                            } else {
                                                fmt.Println(v43)
                                                v43 = (v43 ^ (-(v43)))
                                                if (406 > 32767) {
                                                    v56 := GenS4{F0: 269}
                                                    v57 := GenS3{F0: 1, F1: 127, F2: 2147483647}
                                                    _ = v56
                                                    _ = v57
                                                }
                                            }
                                        case "":
                                            if true {
                                                var v58 bool = false
                                                switch ("" + "a") {
                                                case "hello":
                                                    v59 := []int{-9223372036854775808, 443, -2147483648}
                                                    _ = v59
                                                default:
                                                    var v60 int = 127
                                                    _ = v60
                                                }
                                                _ = v58
                                            } else {
                                                v61 := []int{0, -9223372036854775808, 32768, 4294967295}
                                                v61 = append(v61, 2147483647)
                                                _ = v61
                                            }
                                        default:
                                            var v62 int16 = (v43 << ((255) & 63))
                                            _ = v62
                                        }
                                        var v63 int64 = 9223372036854775807
                                        var v64 uint64 = 805
                                        _ = v63
                                        _ = v64
                                    } else {
                                        var v65 int16 = (-(v43))
                                        var v66 int16 = 1000
                                        _ = v65
                                        _ = v66
                                    }
                                }
                            }
                        }
                        switch 925 {
                        case 32768:
                            fmt.Println(v43)
                        default:
                            var v67 uint64 = 226
                            _ = v67
                        }
                        v43 += (v43 % (((v43 & -32768)) | 1))
                    } else {
                        var v68 float32 = 1.5
                        _ = v68
                    }
                    if ("\n\t" == "\n\t") {
                        if false {
                            v43 += (v43 << ((265) & 63))
                        }
                        if true {
                            switch 9223372036854775807 {
                            case -9223372036854775808:
                                var v69 float32 = 0.1
                                _ = v69
                            case 2147483647:
                                var v70 int64 = -9223372036854775808
                                _ = v70
                            }
                            for fi71 := 0; fi71 < 15; fi71++ {
                                for fi72 := 0; fi72 < 9; fi72++ {
                                    if false {
                                        var v73 int32 = 2147483647
                                        _ = v73
                                    } else {
                                        v74 := []int{(fi71 << ((746) & 63)), fi71, -160}
                                        if (!(-2147483648 != -2147483648)) {
                                            if false {
                                                v75 := GenS7{F0: -1, F1: 65535}
                                                if true {
                                                    fmt.Println(fi71, v43)
                                                } else {
                                                    switch "" {
                                                    case "\n\t":
                                                        v43 = v43
                                                    case "a":
                                                        var v76 int16 = 1000
                                                        _ = v76
                                                    case "hello":
                                                        switch "\n\t" {
                                                        case "foo bar":
                                                            if true {
                                                                v43 -= v43
                                                                var v77 float32 = 1e10
                                                                v78 := GenS8{F0: 9223372036854775807}
                                                                _ = v77
                                                                _ = v78
                                                            }
                                                        default:
                                                            fmt.Println(v43)
                                                        }
                                                    default:
                                                        v79 := GenS9{F0: (("foo bar" + "") + ("hello" + "a"))}
                                                        var v80 int64 = -275
                                                        v43 = (v43 / (((-(v43))) | 1))
                                                        _ = v79
                                                        _ = v80
                                                    }
                                                    fmt.Println(fi72, fi72, v43)
                                                    _ = len(v74)
                                                }
                                                v43 += (v43 % ((v43) | 1))
                                                _ = v75
                                            }
                                            fmt.Println(fi72, fi72, v43)
                                        }
                                        _ = v74
                                    }
                                    var v81 uint16 = 255
                                    v81 = v81
                                    v82 := []int{(-(fi71))}
                                    _ = fi72
                                    _ = v81
                                    _ = v82
                                }
                                switch (fi71 ^ (-(fi71))) {
                                case 9223372036854775807:
                                    var v83 uint32 = 562
                                    _ = v83
                                case -9223372036854775808:
                                    v84 := GenS8{F0: 0}
                                    _ = v84
                                case 846:
                                    v85 := GenS7{F0: -1, F1: 42}
                                    _ = v85
                                default:
                                    fmt.Println(v43, v43)
                                }
                                _ = fi71
                            }
                            fmt.Println(v43)
                        }
                        for fi86 := 0; fi86 < 9; fi86++ {
                            var v87 float64 = (1.5 - float64(float32(1.5)))
                            var v88 string = "a"
                            _ = fi86
                            _ = v87
                            _ = v88
                        }
                    }
                }
                if false {
                    if false {
                        if false {
                            if false {
                                switch 638 {
                                case 2147483647:
                                    if false {
                                        v43 = (v43 ^ (v43 & v43))
                                    } else {
                                        v89 := GenS3{F0: -128, F1: 0, F2: 2}
                                        var v90 int64 = genFn0()
                                        v90 += 9223372036854775807
                                        _ = v89
                                        _ = v90
                                    }
                                case -581:
                                    if false {
                                        for fi91 := 0; fi91 < 16; fi91++ {
                                            fmt.Println(fi91, fi91)
                                            v43 = (v43 << ((883) & 63))
                                            _ = fi91
                                        }
                                    } else {
                                        var v92 int32 = 70
                                        _ = v92
                                    }
                                default:
                                    var v93 uint8 = 0
                                    _ = v93
                                }
                            } else {
                                if (!false) {
                                    for fi94 := 0; fi94 < 2; fi94++ {
                                        var v95 int = (-(fi94))
                                        _ = fi94
                                        _ = v95
                                    }
                                    if (-2.25 != (1e10 * float64(3.14159))) {
                                        v96 := []int{0, 128, 9223372036854775807, 4294967295}
                                        v43 = (v43 >> ((1) & 63))
                                        v43 -= (^(v43))
                                        _ = v96
                                    }
                                    var v97 uint32 = 42
                                    _ = v97
                                } else {
                                    if (((42 > 388) && (0.0 == float64((0.0 / (-2.25 + (3.14159 + -2.25)))))) || true) {
                                        for fi98 := 0; fi98 < 9; fi98++ {
                                            fmt.Println(fi98, v43)
                                            for fi99 := 0; fi99 < 18; fi99++ {
                                                v43 = (^(v43))
                                                v43 += v43
                                                _ = fi99
                                            }
                                            if (v43 >= (v43 + (v43 - 820))) {
                                                switch fi98 {
                                                case -9223372036854775808:
                                                    switch (("a" + "") + ("a" + "foo bar")) {
                                                    case "a":
                                                        for fi100 := 0; fi100 < 11; fi100++ {
                                                            v101 := GenS10{F0: "foo bar", F1: 42, F2: 185}
                                                            v43 -= (v43 >> ((42) & 63))
                                                            _ = fi100
                                                            _ = v101
                                                        }
                                                    case "foo bar":
                                                        v102 := GenS9{F0: "\n\t"}
                                                        _ = v102
                                                    default:
                                                        v43 += v43
                                                    }
                                                case 4294967295:
                                                    if true {
                                                        v103 := GenS8{F0: -768}
                                                        _ = v103
                                                    } else {
                                                        switch "foo bar" {
                                                        case "":
                                                            switch "foo bar" {
                                                            case "":
                                                                v104 := []int{32768, -2147483648, 9223372036854775807, fi98}
                                                                _ = v104
                                                            case "\n\t":
                                                                for fi105 := 0; fi105 < 10; fi105++ {
                                                                    v106 := GenS9{F0: ("foo bar" + "foo bar")}
                                                                    _ = fi105
                                                                    _ = v106
                                                                }
                                                            default:
                                                                var v107 uint = 2
                                                                _ = v107
                                                            }
                                                        default:
                                                            switch "foo bar" {
                                                            case "a":
                                                                if (-53 != 1) {
                                                                    switch "a" {
                                                                    case "a":
                                                                        switch "a" {
                                                                        case "":
                                                                            v108 := GenS9{F0: ("" + ("foo bar" + ((("hello" + ("foo bar" + "")) + "foo bar") + "a")))}
                                                                            _ = v108
                                                                        case "\n\t":
                                                                            if true {
                                                                                for fi109 := 0; fi109 < 8; fi109++ {
                                                                                    v110 := []int{fi98, 183}
                                                                                    switch "foo bar" {
                                                                                    case "foo bar":
                                                                                        fmt.Println(v43)
                                                                                    case "a":
                                                                                        v43 = (v43 | v43)
                                                                                    }
                                                                                    v111 := GenS4{F0: 0}
                                                                                    switch (fi109 + (fi109 / ((fi109) | 1))) {
                                                                                    case -1:
                                                                                        switch 196 {
                                                                                        case 9223372036854775807:
                                                                                            var v112 int64 = 1
                                                                                            _ = v112
                                                                                        case 2147483647:
                                                                                            v43 += v43
                                                                                        case 1:
                                                                                            var v113 string = ("foo bar" + ("foo bar" + "a"))
                                                                                            _ = v113
                                                                                        default:
                                                                                            switch (fi109 << ((255) & 63)) {
                                                                                            case 42:
                                                                                                if len(v110) > 0 {
                                                                                                    v110[0] = (fi109 << ((2) & 63))
                                                                                                }
                                                                                            }
                                                                                        }
                                                                                        fmt.Println(v43)
                                                                                    case 4294967295:
                                                                                        v43 = v43
                                                                                    case -731:
                                                                                        var v114 int = (fi109 / (((fi109 & (fi98 & (fi98 & fi109)))) | 1))
                                                                                        _ = v114
                                                                                    default:
                                                                                        var v115 string = "\n\t"
                                                                                        _ = v115
                                                                                    }
                                                                                    if false {
                                                                                        var v116 uint16 = 65535
                                                                                        fmt.Println(fi109)
                                                                                        if (!(((97 <= 127) || (((-2.25 / (float64((1e10 + (0.0 / -2.25))) * 1e10)) - 1e10) != (-2.25 * 0.0))) || ((v43 - v43) > int16(v116)))) {
                                                                                            v116 = v116
                                                                                        } else {
                                                                                            v116 -= 65535
                                                                                            v117 := GenS11{F0: 42, F1: -9223372036854775808, F2: 2}
                                                                                            if true {
                                                                                                var v118 uint64 = 4294967295
                                                                                                _ = len(v110)
                                                                                                _ = v118
                                                                                            }
                                                                                            _ = v117
                                                                                        }
                                                                                        switch "" {
                                                                                        case "a":
                                                                                            var v119 uint8 = 1
                                                                                            v43 = -32768
                                                                                            var v120 uint32 = 1
                                                                                            _ = v119
                                                                                            _ = v120
                                                                                        default:
                                                                                            v121 := GenS5{F0: (uint32(v116) < uint32(v116))}
                                                                                            _ = v121
                                                                                        }
                                                                                        _ = v116
                                                                                    } else {
                                                                                        if (false && (42 != 65535)) {
                                                                                            if (255 == 42) {
                                                                                                v43 = (v43 >> ((65535) & 63))
                                                                                                var v122 float32 = (1.5 + (3.14159 - 1e10))
                                                                                                _ = v122
                                                                                            }
                                                                                            v43 += -32768
                                                                                            switch (-(fi98)) {
                                                                                            case 32768:
                                                                                                if (true && (0 < 714)) {
                                                                                                    var v123 uint16 = 1
                                                                                                    _ = v123
                                                                                                } else {
                                                                                                    switch (fi109 * (^(fi109))) {
                                                                                                    case 1:
                                                                                                        var v124 float32 = (-2.25 - -2.25)
                                                                                                        _ = v124
                                                                                                    case 0:
                                                                                                        fmt.Println(v43, fi98)
                                                                                                    }
                                                                                                    if false {
                                                                                                        switch fi98 {
                                                                                                        case 127:
                                                                                                            var v125 float64 = 0.1
                                                                                                            _ = v125
                                                                                                        }
                                                                                                        fmt.Println(fi98, v43, fi109)
                                                                                                        switch "\n\t" {
                                                                                                        case "a":
                                                                                                            var v126 bool = (5 <= -128)
                                                                                                            _ = v126
                                                                                                        case "foo bar":
                                                                                                            v43 += v43
                                                                                                            fmt.Println(v43)
                                                                                                        default:
                                                                                                            var v127 float64 = (float64(1.5) + float64(0.0))
                                                                                                            var v128 uint64 = 4294967295
                                                                                                            var v129 uint16 = 65535
                                                                                                            _ = v127
                                                                                                            _ = v128
                                                                                                            _ = v129
                                                                                                        }
                                                                                                    } else {
                                                                                                        var v130 bool = (!(104 >= 42))
                                                                                                        v131 := GenS5{F0: v130}
                                                                                                        switch "\n\t" {
                                                                                                        case "":
                                                                                                            if v130 {
                                                                                                                fmt.Println(v43)
                                                                                                                v132 := GenS12{F0: (fi109 % ((fi109) | 1))}
                                                                                                                var v133 int64 = genFn0()
                                                                                                                _ = v132
                                                                                                                _ = v133
                                                                                                            }
                                                                                                        default:
                                                                                                            switch ("foo bar" + (("\n\t" + "hello") + "hello")) {
                                                                                                            case "foo bar":
                                                                                                                if v130 {
                                                                                                                    switch fi109 {
                                                                                                                    case 32768:
                                                                                                                        if v130 {
                                                                                                                            v43 -= v43
                                                                                                                            v130 = ((2147483647 > 42) || (127 <= 1))
                                                                                                                            var v134 uint64 = 882
                                                                                                                            _ = v134
                                                                                                                        }
                                                                                                                    case 0:
                                                                                                                        v130 = v130
                                                                                                                    default:
                                                                                                                        v135 := GenS13{F0: -1, F1: 512, F2: 65535}
                                                                                                                        var v136 int16 = 516
                                                                                                                        v137 := GenS2{F0: (-(fi98)), F1: -1}
                                                                                                                        _ = v135
                                                                                                                        _ = v136
                                                                                                                        _ = v137
                                                                                                                    }
                                                                                                                    var v138 int8 = 1
                                                                                                                    _ = v138
                                                                                                                } else {
                                                                                                                    fmt.Println(fi98)
                                                                                                                    v43 += (v43 % (((v43 ^ (v43 & (-(v43))))) | 1))
                                                                                                                    if len(v110) > 0 {
                                                                                                                        v110[0] = fi109
                                                                                                                    }
                                                                                                                }
                                                                                                                v43 = (v43 / ((v43) | 1))
                                                                                                            case "\n\t":
                                                                                                                if true {
                                                                                                                    v43 -= (v43 >> ((2) & 63))
                                                                                                                    var v139 uint8 = 0
                                                                                                                    _ = v139
                                                                                                                } else {
                                                                                                                    var v140 int64 = 42
                                                                                                                    _ = v140
                                                                                                                }
                                                                                                            }
                                                                                                        }
                                                                                                        _ = v130
                                                                                                        _ = v131
                                                                                                    }
                                                                                                    if (0 >= 42) {
                                                                                                        fmt.Println(fi98, fi109, fi98)
                                                                                                        v141 := GenS10{F0: ("\n\t" + ("foo bar" + ("foo bar" + "foo bar"))), F1: 402, F2: 159}
                                                                                                        _ = v141
                                                                                                    } else {
                                                                                                        v110 = append(v110, fi98)
                                                                                                        switch ("hello" + ("a" + ("" + ("hello" + "a")))) {
                                                                                                        case "":
                                                                                                            v142 := GenS14{F0: 2, F1: -12, F2: 78}
                                                                                                            v143 := GenS8{F0: -9223372036854775808}
                                                                                                            v144 := GenS13{F0: 1, F1: 9223372036854775807, F2: 18446744073709551615}
                                                                                                            if (!((-(v43)) == v43)) {
                                                                                                                if false {
                                                                                                                    if true {
                                                                                                                        switch "foo bar" {
                                                                                                                        case "a":
                                                                                                                            var v145 uint8 = 255
                                                                                                                            var v146 uint = 18446744073709551615
                                                                                                                            fmt.Println(v145)
                                                                                                                            v145 -= uint8(v146)
                                                                                                                            _ = v145
                                                                                                                            _ = v146
                                                                                                                        default:
                                                                                                                            v43 = (v43 + (v43 << ((4294967295) & 63)))
                                                                                                                        }
                                                                                                                        var v147 int8 = 1
                                                                                                                        v147 += v147
                                                                                                                        _ = v147
                                                                                                                    } else {
                                                                                                                        v43 = v43
                                                                                                                    }
                                                                                                                    if (true || true) {
                                                                                                                        switch "foo bar" {
                                                                                                                        case "\n\t":
                                                                                                                            if true {
                                                                                                                                var v148 string = ((("" + "a") + "foo bar") + "\n\t")
                                                                                                                                v43 += 1000
                                                                                                                                _ = v148
                                                                                                                            }
                                                                                                                        case "":
                                                                                                                            if (-870 <= -2147483648) {
                                                                                                                                fmt.Println(fi109, fi98)
                                                                                                                                switch ("hello" + "\n\t") {
                                                                                                                                case "a":
                                                                                                                                    v43 += (v43 | v43)
                                                                                                                                case "\n\t":
                                                                                                                                    fmt.Println(fi98, v43)
                                                                                                                                case "":
                                                                                                                                    var v149 float64 = -2.25
                                                                                                                                    if false {
                                                                                                                                        v149 = v149
                                                                                                                                        v43 = (v43 - v43)
                                                                                                                                        v149 = float64(0.1)
                                                                                                                                        v43 = v43
                                                                                                                                    }
                                                                                                                                    v149 = (1e10 / (v149 - v149))
                                                                                                                                    _ = v149
                                                                                                                                }
                                                                                                                                switch fi98 {
                                                                                                                                case -600:
                                                                                                                                    if false {
                                                                                                                                        v150 := GenS15{F0: 42, F1: ("" + ("" + ((((("foo bar" + (("a" + ("" + "\n\t")) + "foo bar")) + "") + "\n\t") + "") + ""))), F2: 100}
                                                                                                                                        var v151 int64 = -9223372036854775808
                                                                                                                                        v152 := GenS9{F0: ("\n\t" + ((("\n\t" + "a") + ("foo bar" + ("\n\t" + ""))) + "hello"))}
                                                                                                                                        _ = v150
                                                                                                                                        _ = v151
                                                                                                                                        _ = v152
                                                                                                                                    }
                                                                                                                                }
                                                                                                                            } else {
                                                                                                                                if (!false) {
                                                                                                                                    v43 = (v43 / ((v43) | 1))
                                                                                                                                    v43 = (v43 >> ((18446744073709551615) & 63))
                                                                                                                                    _ = len(v110)
                                                                                                                                } else {
                                                                                                                                    fmt.Println(fi109)
                                                                                                                                    v153 := GenS16{F0: 32767, F1: 42}
                                                                                                                                    _ = v153
                                                                                                                                }
                                                                                                                            }
                                                                                                                        }
                                                                                                                        var v154 int32 = 1
                                                                                                                        _ = v154
                                                                                                                    } else {
                                                                                                                        if len(v110) >= 2 {
                                                                                                                            v110 = v110[1:len(v110)]
                                                                                                                        }
                                                                                                                        fmt.Println(fi98, fi109)
                                                                                                                        v43 = v43
                                                                                                                    }
                                                                                                                    var v155 uint64 = 42
                                                                                                                    _ = v155
                                                                                                                }
                                                                                                                v156 := GenS17{F0: (("" + "a") + (("foo bar" + "") + "a")), F1: 2147483648}
                                                                                                                _ = v156
                                                                                                            }
                                                                                                            _ = v142
                                                                                                            _ = v143
                                                                                                            _ = v144
                                                                                                        case "hello":
                                                                                                            v157 := GenS15{F0: -1, F1: "", F2: 100}
                                                                                                            _ = v157
                                                                                                        case "\n\t":
                                                                                                            v110 = append(v110, (fi98 ^ (-(fi109))))
                                                                                                        }
                                                                                                        if len(v110) > 0 {
                                                                                                            v110[0] = fi109
                                                                                                        }
                                                                                                    }
                                                                                                }
                                                                                            default:
                                                                                                var v158 uint = 240
                                                                                                _ = v158
                                                                                            }
                                                                                        } else {
                                                                                            if len(v110) > 0 {
                                                                                                v110[0] = (fi98 >> ((42) & 63))
                                                                                            }
                                                                                            if ("hello" != (("" + (("" + "a") + "")) + "")) {
                                                                                                switch "" {
                                                                                                case "foo bar":
                                                                                                    v43 = v43
                                                                                                case "":
                                                                                                    v159 := GenS9{F0: (("foo bar" + ((("a" + "hello") + "\n\t") + ("foo bar" + ("\n\t" + "hello")))) + "")}
                                                                                                    _ = len(v110)
                                                                                                    _ = v159
                                                                                                }
                                                                                                if false {
                                                                                                    var v160 uint64 = 0
                                                                                                    var v161 float64 = (0.1 - 1e10)
                                                                                                    _ = v160
                                                                                                    _ = v161
                                                                                                } else {
                                                                                                    var v162 int = (fi98 * fi109)
                                                                                                    var v163 uint8 = 192
                                                                                                    var v164 uint64 = 815
                                                                                                    _ = v162
                                                                                                    _ = v163
                                                                                                    _ = v164
                                                                                                }
                                                                                                var v165 int = -10
                                                                                                _ = v165
                                                                                            }
                                                                                            if ((0.0 / 1.0) < 1.5) {
                                                                                                v43 += 32767
                                                                                                var v166 int32 = -1
                                                                                                if ((fi109 < -830) && ("foo bar" <= ((("\n\t" + ("foo bar" + "a")) + "foo bar") + ""))) {
                                                                                                    v43 -= v43
                                                                                                    v43 -= (v43 << ((4294967295) & 63))
                                                                                                    fmt.Println(fi109)
                                                                                                }
                                                                                                _ = v166
                                                                                            }
                                                                                            if (127 < 127) {
                                                                                                v43 = (v43 / ((v43) | 1))
                                                                                                var v167 bool = false
                                                                                                fmt.Println(v167)
                                                                                                _ = v167
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                    switch "" {
                                                                                    case "\n\t":
                                                                                        if (false || (128 < (fi109 % (((fi98 & fi98)) | 1)))) {
                                                                                            v110 = append(v110, 0)
                                                                                            var v168 float64 = ((1.5 * 3.14159) * 3.14159)
                                                                                            _ = v168
                                                                                        }
                                                                                    default:
                                                                                        switch fi98 {
                                                                                        case 32767:
                                                                                            var v169 int64 = 112
                                                                                            _ = v169
                                                                                        default:
                                                                                            v43 = v43
                                                                                        }
                                                                                    }
                                                                                    if false {
                                                                                        if true {
                                                                                            v170 := GenS18{F0: 2147483647}
                                                                                            if len(v110) >= 2 {
                                                                                                v110 = v110[1:len(v110)]
                                                                                            }
                                                                                            _ = v170
                                                                                        }
                                                                                        var v171 int32 = 0
                                                                                        _ = v171
                                                                                    }
                                                                                    _ = len(v110)
                                                                                    var v172 int32 = -223
                                                                                    _ = fi109
                                                                                    _ = v110
                                                                                    _ = v111
                                                                                    _ = v172
                                                                                }
                                                                                var v173 uint16 = 496
                                                                                if true {
                                                                                    for fi174 := 0; fi174 < 13; fi174++ {
                                                                                        v175 := GenS19{F0: 42, F1: 255}
                                                                                        if (((-822 > 994) && false) || (!(26 >= 2))) {
                                                                                            var v176 uint32 = 42
                                                                                            fmt.Println(v173)
                                                                                            _ = v176
                                                                                        }
                                                                                        _ = fi174
                                                                                        _ = v175
                                                                                    }
                                                                                    for fi177 := 0; fi177 < 18; fi177++ {
                                                                                        switch "a" {
                                                                                        case "hello":
                                                                                            fmt.Println(fi177)
                                                                                        case "foo bar":
                                                                                            v43 += -987
                                                                                        case "":
                                                                                            if false {
                                                                                                if (1e10 == float64(0.1)) {
                                                                                                    var v178 uint64 = 18446744073709551615
                                                                                                    _ = v178
                                                                                                }
                                                                                            } else {
                                                                                                var v179 uint16 = v173
                                                                                                var v180 int64 = 9223372036854775807
                                                                                                v181 := GenS15{F0: 54, F1: "", F2: 42}
                                                                                                _ = v179
                                                                                                _ = v180
                                                                                                _ = v181
                                                                                            }
                                                                                        }
                                                                                        switch fi177 {
                                                                                        case -1:
                                                                                            var v182 bool = false
                                                                                            _ = v182
                                                                                        case 32768:
                                                                                            fmt.Println(v173, fi98, v43)
                                                                                        }
                                                                                        var v183 int16 = 1
                                                                                        if (false && false) {
                                                                                            switch "foo bar" {
                                                                                            case "a":
                                                                                                v184 := []int{}
                                                                                                _ = v184
                                                                                            case "\n\t":
                                                                                                v183 = int16(v173)
                                                                                            default:
                                                                                                v173 = v173
                                                                                            }
                                                                                            v43 -= v183
                                                                                            switch "\n\t" {
                                                                                            case "\n\t":
                                                                                                if true {
                                                                                                    var v185 uint16 = (v173 ^ (^(v173)))
                                                                                                    if (!false) {
                                                                                                        if ((("a" + "") + "a") <= "") {
                                                                                                            if false {
                                                                                                                v186 := []int{(^(fi98)), (fi98 % ((672) | 1))}
                                                                                                                fmt.Println(v173, fi177)
                                                                                                                var v187 int64 = 0
                                                                                                                fmt.Println(fi177)
                                                                                                                if false {
                                                                                                                    var v188 uint8 = 42
                                                                                                                    v173 += (v173 << ((42) & 63))
                                                                                                                    _ = v188
                                                                                                                }
                                                                                                                _ = v186
                                                                                                                _ = v187
                                                                                                            } else {
                                                                                                                v189 := []int{(fi98 >> ((2) & 63)), (fi98 % ((-259) | 1)), fi177}
                                                                                                                switch ("foo bar" + "a") {
                                                                                                                case "\n\t":
                                                                                                                    _ = len(v189)
                                                                                                                }
                                                                                                                _ = v189
                                                                                                            }
                                                                                                            if true {
                                                                                                                var v190 uint32 = 226
                                                                                                                var v191 float64 = 0.0
                                                                                                                _ = v190
                                                                                                                _ = v191
                                                                                                            }
                                                                                                            v183 += (v43 % ((v43) | 1))
                                                                                                        }
                                                                                                        fmt.Println(v183)
                                                                                                        v192 := []int{(fi177 * (fi98 ^ 128))}
                                                                                                        var v193 uint16 = (^(v185))
                                                                                                        _ = v192
                                                                                                        _ = v193
                                                                                                    }
                                                                                                    var v194 int8 = 100
                                                                                                    _ = v185
                                                                                                    _ = v194
                                                                                                }
                                                                                            case "hello":
                                                                                                if (!false) {
                                                                                                    v43 = 1
                                                                                                }
                                                                                            case "a":
                                                                                                v195 := []int{}
                                                                                                _ = v195
                                                                                            }
                                                                                        } else {
                                                                                            v183 += 42
                                                                                            fmt.Println(fi98, v183)
                                                                                        }
                                                                                        v43 = v183
                                                                                        _ = fi177
                                                                                        _ = v183
                                                                                    }
                                                                                    var v196 int32 = -978
                                                                                    _ = v196
                                                                                }
                                                                                _ = v173
                                                                            }
                                                                        }
                                                                    case "\n\t":
                                                                        for fi197 := 0; fi197 < 16; fi197++ {
                                                                            v198 := GenS19{F0: 2, F1: 47}
                                                                            switch (fi98 % ((108) | 1)) {
                                                                            case 127:
                                                                                if false {
                                                                                    v43 = 1000
                                                                                    v199 := GenS2{F0: (-(fi98)), F1: 75}
                                                                                    _ = v199
                                                                                } else {
                                                                                    v43 -= (v43 << ((42) & 63))
                                                                                    v43 -= -695
                                                                                }
                                                                            case 32767:
                                                                                v43 -= v43
                                                                            }
                                                                            fmt.Println(fi98, fi197)
                                                                            _ = fi197
                                                                            _ = v198
                                                                        }
                                                                    }
                                                                    var v200 uint = 2
                                                                    v201 := []int{(fi98 & (-(fi98)))}
                                                                    _ = v200
                                                                    _ = v201
                                                                } else {
                                                                    var v202 string = ""
                                                                    v43 = v43
                                                                    _ = v202
                                                                }
                                                            case "":
                                                                for fi203 := 0; fi203 < 12; fi203++ {
                                                                    var v204 uint = 129
                                                                    _ = fi203
                                                                    _ = v204
                                                                }
                                                            case "hello":
                                                                var v205 uint64 = 2
                                                                _ = v205
                                                            }
                                                        }
                                                        v43 -= (v43 * v43)
                                                    }
                                                case 375:
                                                    switch 128 {
                                                    case 2147483647:
                                                        switch (("" + "hello") + "foo bar") {
                                                        case "\n\t":
                                                            var v206 uint64 = 413
                                                            _ = v206
                                                        case "hello":
                                                            v207 := GenS9{F0: "a"}
                                                            _ = v207
                                                        case "":
                                                            v43 = (-(v43))
                                                        }
                                                    case 4294967295:
                                                        var v208 int = fi98
                                                        _ = v208
                                                    case -2147483648:
                                                        v209 := []int{(fi98 & fi98), fi98, (fi98 * (fi98 + (fi98 ^ fi98))), 898}
                                                        _ = v209
                                                    }
                                                default:
                                                    var v210 float64 = ((0.0 * 0.0) + (1e10 + 3.14159))
                                                    _ = v210
                                                }
                                            } else {
                                                v43 -= (v43 * (v43 >> ((255) & 63)))
                                                v43 = 355
                                                v211 := []int{(fi98 - (-(fi98))), -398, (-(fi98))}
                                                _ = v211
                                            }
                                            _ = fi98
                                        }
                                        v212 := []int{}
                                        _ = v212
                                    }
                                    v213 := []int{}
                                    switch "foo bar" {
                                    case "foo bar":
                                        var v214 uint16 = 2
                                        _ = v214
                                    case "":
                                        var v215 uint8 = 1
                                        _ = v215
                                    default:
                                        v43 = (^(v43))
                                    }
                                    _ = v213
                                }
                            }
                            v43 -= (^(v43))
                            if true {
                                v43 = (v43 << ((4294967295) & 63))
                                for fi216 := 0; fi216 < 12; fi216++ {
                                    var v217 uint16 = 981
                                    v218 := []int{(fi216 ^ fi216), 4294967295, (fi216 << ((0) & 63))}
                                    v218 = append(v218, (fi216 & fi216))
                                    _ = fi216
                                    _ = v217
                                    _ = v218
                                }
                                fmt.Println(v43)
                            } else {
                                for fi219 := 0; fi219 < 15; fi219++ {
                                    var v220 uint = 236
                                    v43 = (v43 | v43)
                                    _ = fi219
                                    _ = v220
                                }
                            }
                        } else {
                            if (true || ((false || true) || ((false && ((3.14159 >= (1e10 * -2.25)) || false)) || false))) {
                                for fi221 := 0; fi221 < 17; fi221++ {
                                    v222 := GenS20{F0: 2, F1: 2147483648}
                                    _ = fi221
                                    _ = v222
                                }
                                v223 := GenS21{F0: "foo bar", F1: 18446744073709551615, F2: 2147483647}
                                fmt.Println(v43)
                                _ = v223
                            } else {
                                v224 := GenS22{F0: 42, F1: 32767, F2: ("" + (("foo bar" + (("a" + ("hello" + ("foo bar" + ""))) + "")) + "hello"))}
                                if false {
                                    v225 := []int{9223372036854775807, 4294967295, -2147483648}
                                    v43 = v43
                                    var v226 uint32 = 4294967295
                                    _ = v225
                                    _ = v226
                                } else {
                                    fmt.Println(v43)
                                    var v227 int16 = (v43 - v43)
                                    v227 += v227
                                    _ = v227
                                }
                                _ = v224
                            }
                        }
                        v43 = v43
                    }
                    v43 = (v43 * v43)
                    v228 := []int{0, 9223372036854775807, 42, 2147483647}
                    _ = v228
                } else {
                    v229 := GenS23{F0: 42, F1: 9223372036854775807}
                    var v230 string = "\n\t"
                    _ = v229
                    _ = v230
                }
            }
            if false {
                for fi231 := 0; fi231 < 19; fi231++ {
                    if (572 <= 42) {
                        var v232 int32 = -1
                        if false {
                            v233 := GenS9{F0: ""}
                            for fi234 := 0; fi234 < 8; fi234++ {
                                v235 := []int{(fi234 ^ fi231), (-(fi231)), (^(fi234))}
                                _ = fi234
                                _ = v235
                            }
                            switch (fi231 >> ((668) & 63)) {
                            case 32768:
                                v232 += (v232 + (v232 & -2147483648))
                            }
                            _ = v233
                        }
                        switch "hello" {
                        case "":
                            if false {
                                switch fi231 {
                                case 42:
                                    switch "a" {
                                    case "":
                                        switch "\n\t" {
                                        case "hello":
                                            var v236 uint = 1
                                            _ = v236
                                        case "foo bar":
                                            if (!(((25 >= 75) && (0 != (^(v232)))) && (!true))) {
                                                var v237 int8 = 42
                                                var v238 int64 = 9223372036854775807
                                                var v239 uint = 692
                                                _ = v237
                                                _ = v238
                                                _ = v239
                                            } else {
                                                for fi240 := 0; fi240 < 18; fi240++ {
                                                    var v241 int32 = (v232 / ((v232) | 1))
                                                    _ = fi240
                                                    _ = v241
                                                }
                                            }
                                        case "a":
                                            if (4294967295 > 868) {
                                                var v242 bool = (!(!(2 > 957)))
                                                for fi243 := 0; fi243 < 6; fi243++ {
                                                    var v244 uint8 = 127
                                                    v244 = v244
                                                    v232 = (v232 ^ -1)
                                                    _ = fi243
                                                    _ = v244
                                                }
                                                v245 := []int{fi231, (^(fi231)), (fi231 & fi231), 128}
                                                _ = v242
                                                _ = v245
                                            } else {
                                                v43 = (v43 << ((243) & 63))
                                                for fi246 := 0; fi246 < 15; fi246++ {
                                                    v43 += v43
                                                    v247 := []int{(fi231 | fi231), fi246, (-(fi231)), 507}
                                                    v232 -= (v232 << ((0) & 63))
                                                    var v248 int16 = 42
                                                    v249 := GenS4{F0: 18446744073709551615}
                                                    if len(v247) >= 2 {
                                                        v247 = v247[1:len(v247)]
                                                    }
                                                    _ = fi246
                                                    _ = v247
                                                    _ = v248
                                                    _ = v249
                                                }
                                                v250 := []int{fi231, fi231, -2147483648, 32768}
                                                _ = v250
                                            }
                                        }
                                    }
                                case -384:
                                    v251 := GenS20{F0: 277, F1: 65535}
                                    _ = v251
                                case 32767:
                                    var v252 uint16 = 546
                                    _ = v252
                                default:
                                    fmt.Println(v43)
                                }
                            }
                        case "\n\t":
                            fmt.Println(fi231)
                        }
                        _ = v232
                    }
                    if false {
                        v43 = (-(v43))
                        for fi253 := 0; fi253 < 8; fi253++ {
                            v43 = (v43 % ((v43) | 1))
                            v43 -= v43
                            _ = fi253
                        }
                        switch (fi231 * (^(fi231))) {
                        case 690:
                            v254 := []int{-191, fi231, fi231}
                            _ = v254
                        default:
                            v255 := GenS24{F0: 1, F1: -899, F2: (true || (-68 < -118))}
                            _ = v255
                        }
                    }
                    for fi256 := 0; fi256 < 1; fi256++ {
                        if true {
                            v257 := GenS5{F0: ((!(true && ((0.1 != (-2.25 - (3.14159 - 0.0))) || false))) || false)}
                            v43 -= v43
                            var v258 uint8 = 42
                            _ = v257
                            _ = v258
                        } else {
                            v259 := []int{fi231, fi231}
                            _ = v259
                        }
                        _ = fi256
                    }
                    _ = fi231
                }
                if true {
                    v260 := GenS25{F0: int8(v43), F1: 2147483647, F2: 32767}
                    var v261 int = 708
                    for fi262 := 0; fi262 < 15; fi262++ {
                        if true {
                            var v263 uint = 653
                            _ = v263
                        } else {
                            fmt.Println(v43)
                            var v264 uint = 255
                            v265 := GenS7{F0: 980, F1: 134}
                            _ = v264
                            _ = v265
                        }
                        v261 = (fi262 >> ((2) & 63))
                        v266 := []int{}
                        _ = fi262
                        _ = v266
                    }
                    _ = v260
                    _ = v261
                } else {
                    v43 -= v43
                    var v267 uint8 = 2
                    var v268 uint16 = 484
                    _ = v267
                    _ = v268
                }
            } else {
                fmt.Println(v43)
                var v269 int64 = -9223372036854775808
                var v270 int8 = 42
                _ = v269
                _ = v270
            }
        }
        switch -9223372036854775808 {
        case 42:
            var v271 float64 = 1.5
            _ = v271
        }
        _ = v44
    }
    _ = v43
    return 2147483647
}

func main() {
    var v1 int = 32768
    v2 := []string{}
    var v3 uint8 = 0
    v4 := []int{-830}
    var v272 string = "\n\t"
    var v273 int8 = 127
    if len(v2) > 0 {
        v2[0] = v272
    }
    v273 = v273
    if (genFn0() > 372) {
        switch genFn1() {
        case -9223372036854775808:
            for fi274 := 0; fi274 < 9; fi274++ {
                for fi275 := 0; fi275 < 11; fi275++ {
                    v2 = append(v2, v272)
                    _ = fi275
                }
                v273 = v273
                v3 = v3
                _ = fi274
            }
        }
        if false {
            var v276 int16 = -1
            v277 := GenS21{F0: (v272 + "\n\t"), F1: 18446744073709551615, F2: 1}
            v278 := GenS5{F0: true}
            _ = v276
            _ = v277
            _ = v278
        }
        fmt.Println(v1)
    }
    fmt.Println(v272, v1)
    v279 := GenS15{F0: 127, F1: v272, F2: (-(v273))}
    fmt.Println(v272, v272)
    switch (v1 % (((v1 / ((genFn1()) | 1))) | 1)) {
    case -929:
        var v280 bool = (true && ((v1 / (((v1 % (((^(v1))) | 1))) | 1)) < v1))
        _ = v280
    case 4294967295:
        if (488 == 679) {
            switch ("hello" + v272) {
            case "hello":
                if ("" != "\n\t") {
                    v272 = "\n\t"
                }
            case "foo bar":
                if false {
                    if false {
                        v1 = (v1 | v1)
                    }
                    var v281 bool = true
                    _ = v281
                } else {
                    var v282 float32 = float32(1.5)
                    v1 = (-(v1))
                    _ = v282
                }
            }
        }
    case -190:
        if (1 < 690) {
            var v283 uint64 = 311
            v4 = append(v4, v1)
            for fi284 := 0; fi284 < 8; fi284++ {
                if (!((v3 <= v3) && false)) {
                    var v285 float64 = 0.1
                    _ = v285
                }
                if len(v4) > 0 {
                    v4[0] = fi284
                }
                v286 := GenS21{F0: v272, F1: (v283 | (v283 | v283)), F2: -17}
                _ = fi284
                _ = v286
            }
            _ = v283
        } else {
            fmt.Println(v273, v1, v3)
            v1 = (v1 + v1)
        }
    default:
        for fi287 := 0; fi287 < 11; fi287++ {
            switch (fi287 * (fi287 >> ((395) & 63))) {
            case 42:
                if true {
                    v3 = (^(v3))
                }
            }
            _ = fi287
        }
    }
    switch (("a" + ("" + v272)) + v272) {
    case "\n\t":
        var v288 int = v1
        _ = v288
    case "foo bar":
        var v289 uint16 = 372
        _ = v289
    case "a":
        v290 := GenS3{F0: (v273 & v273), F1: (v273 >> ((1) & 63)), F2: 42}
        _ = v290
    }
    v272 = v272
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v272:", v272)
    fmt.Println("v273:", v273)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v272
    _ = v273
    _ = v279
    _ = v3
    _ = v4
}
