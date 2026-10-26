package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool }
type GenS2 struct { F0 int64; F1 uint32; F2 int }
type GenS3 struct { F0 int16; F1 bool; F2 int8 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint8; F1 int8; F2 uint8 }
type GenS6 struct { F0 int8; F1 int32; F2 uint }
type GenS7 struct { F0 uint64; F1 int8; F2 int32 }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 uint8; F1 uint64 }
type GenS10 struct { F0 int16; F1 int8 }
type GenS11 struct { F0 uint16; F1 uint16; F2 int32 }
type GenS12 struct { F0 uint16; F1 uint; F2 string }
func genFn0(p0 bool) int {
    if ((!false) || p0) {
        var v4 uint16 = 2
        _ = v4
    }
    for fi5 := 0; fi5 < 11; fi5++ {
        var v6 uint16 = 1
        p0 = true
        _ = fi5
        _ = v6
    }
    if false {
        var v7 int = 1
        _ = v7
    } else {
        v8 := GenS1{F0: 0, F1: true}
        v9 := []int{32768, 9223372036854775807, -2147483648}
        for fi10 := 0; fi10 < 6; fi10++ {
            p0 = p0
            _ = fi10
        }
        _ = v8
        _ = v9
    }
    v11 := GenS1{F0: 0, F1: (!(1 < -475))}
    p0 = (false && (p0 && false))
    _ = p0
    _ = v11
    return -9223372036854775808
}
func genFn1(p0 bool) uint8 {
    fmt.Println(p0)
    p0 = p0
    for fi12 := 0; fi12 < 11; fi12++ {
        p0 = (!false)
        _ = fi12
    }
    if (576 != 4294967295) {
        v13 := []int{}
        _ = v13
    } else {
        v14 := GenS1{F0: -31, F1: true}
        if p0 {
            for fi15 := 0; fi15 < 6; fi15++ {
                switch fi15 {
                case -9223372036854775808:
                    fmt.Println(p0)
                default:
                    var v16 int32 = 205
                    _ = v16
                }
                v17 := []int{(fi15 * fi15), (fi15 << ((2) & 63)), (-(fi15))}
                for fi18 := 0; fi18 < 6; fi18++ {
                    if (102 != 32767) {
                        if (!((("" + "") + "a") < "\n\t")) {
                            if p0 {
                                if p0 {
                                    v19 := GenS2{F0: 9223372036854775807, F1: 1, F2: -2147483648}
                                    _ = v19
                                } else {
                                    var v20 uint = 490
                                    _ = v20
                                }
                                if true {
                                    switch "foo bar" {
                                    case "a":
                                        if true {
                                            p0 = (922 >= 32767)
                                            v21 := GenS2{F0: 9223372036854775807, F1: 461, F2: 903}
                                            _ = v21
                                        }
                                    case "":
                                        var v22 string = ""
                                        _ = v22
                                    case "foo bar":
                                        p0 = ((1e10 != 0.1) && false)
                                    default:
                                        switch "foo bar" {
                                        case "a":
                                            var v23 int = fi15
                                            if (!p0) {
                                                var v24 uint32 = 1
                                                if p0 {
                                                    v24 += (^(v24))
                                                    fmt.Println(fi18)
                                                    if len(v17) >= 2 {
                                                        v17 = v17[1:len(v17)]
                                                    }
                                                    switch fi15 {
                                                    case -2147483648:
                                                        switch fi15 {
                                                        case 128:
                                                            if len(v17) > 0 {
                                                                v17[0] = v23
                                                            }
                                                        default:
                                                            var v25 bool = (!p0)
                                                            _ = v25
                                                        }
                                                    }
                                                    v26 := GenS1{F0: -1, F1: p0}
                                                    _ = v26
                                                } else {
                                                    v23 -= v23
                                                    if len(v17) >= 2 {
                                                        v17 = v17[1:len(v17)]
                                                    }
                                                    if (0 >= 0) {
                                                        fmt.Println(fi15)
                                                        if p0 {
                                                            v17 = append(v17, 32768)
                                                            fmt.Println(v23, v23, fi18)
                                                        }
                                                        if p0 {
                                                            switch (fi15 / (((fi15 << ((65535) & 63))) | 1)) {
                                                            case 1:
                                                                if (110 != 1) {
                                                                    if (32767 < 0) {
                                                                        var v27 uint64 = 18446744073709551615
                                                                        v28 := GenS2{F0: -255, F1: (v24 + (v24 * 2147483647)), F2: (^(fi18))}
                                                                        var v29 int = v23
                                                                        _ = v27
                                                                        _ = v28
                                                                        _ = v29
                                                                    }
                                                                    fmt.Println(p0)
                                                                    p0 = (!(42 == 1))
                                                                } else {
                                                                    switch v23 {
                                                                    case 1:
                                                                        fmt.Println(v24)
                                                                    }
                                                                    if len(v17) >= 2 {
                                                                        v17 = v17[1:len(v17)]
                                                                    }
                                                                }
                                                            case 4294967295:
                                                                v24 = (^(v24))
                                                            default:
                                                                if true {
                                                                    switch (fi15 >> ((466) & 63)) {
                                                                    case -689:
                                                                        fmt.Println(fi18, fi18, v23)
                                                                    case 2147483647:
                                                                        v23 = (fi18 % ((fi15) | 1))
                                                                    default:
                                                                        fmt.Println(fi18, v24)
                                                                    }
                                                                } else {
                                                                    v23 -= (v23 * fi15)
                                                                    var v30 bool = (!p0)
                                                                    _ = v30
                                                                }
                                                            }
                                                        }
                                                        fmt.Println(p0, p0)
                                                        v23 -= v23
                                                        fmt.Println(fi15, fi15)
                                                    } else {
                                                        var v31 uint = 0
                                                        _ = v31
                                                    }
                                                }
                                                _ = len(v17)
                                                _ = v24
                                            } else {
                                                if len(v17) > 0 {
                                                    v17[0] = (fi18 << ((65535) & 63))
                                                }
                                            }
                                            p0 = (0 >= 42)
                                            var v32 uint8 = 127
                                            switch 4294967295 {
                                            case -2147483648:
                                                v33 := GenS1{F0: -1, F1: p0}
                                                _ = v33
                                            case 699:
                                                v32 -= (v32 & 28)
                                            case 4294967295:
                                                v23 += 127
                                            }
                                            if len(v17) > 0 {
                                                v17[0] = genFn0((-28 <= 100))
                                            }
                                            _ = v23
                                            _ = v32
                                        case "\n\t":
                                            if p0 {
                                                p0 = true
                                                p0 = true
                                            } else {
                                                switch 9223372036854775807 {
                                                case 936:
                                                    v17 = append(v17, fi15)
                                                case 32768:
                                                    fmt.Println(p0)
                                                case 324:
                                                    if p0 {
                                                        v17 = append(v17, (fi15 / (((-(fi15))) | 1)))
                                                        _ = len(v17)
                                                        v34 := GenS1{F0: 1, F1: p0}
                                                        _ = v34
                                                    }
                                                }
                                                p0 = (p0 || true)
                                                p0 = p0
                                            }
                                        case "foo bar":
                                            if (!(65535 >= 2147483648)) {
                                                p0 = false
                                                var v35 int64 = 42
                                                _ = v35
                                            } else {
                                                var v36 int16 = -306
                                                var v37 int8 = 127
                                                _ = v36
                                                _ = v37
                                            }
                                        default:
                                            p0 = true
                                        }
                                    }
                                } else {
                                    var v38 float64 = 0.1
                                    fmt.Println(p0, v38)
                                    _ = v38
                                }
                                switch ("hello" + "hello") {
                                case "foo bar":
                                    var v39 string = ""
                                    _ = v39
                                case "a":
                                    var v40 bool = p0
                                    _ = v40
                                case "":
                                    if p0 {
                                        if (0 > -128) {
                                            p0 = (-1 < -1)
                                            p0 = true
                                            if (255 < 221) {
                                                fmt.Println(fi18, fi15, fi18)
                                                fmt.Println(p0)
                                                var v41 uint8 = 165
                                                _ = v41
                                            }
                                        }
                                        var v42 uint64 = 2
                                        _ = len(v17)
                                        _ = v42
                                    }
                                default:
                                    v43 := GenS1{F0: -2147483648, F1: (p0 && p0)}
                                    _ = v43
                                }
                            } else {
                                fmt.Println(fi18, p0, fi18)
                                switch "\n\t" {
                                case "a":
                                    var v44 float32 = 0.0
                                    _ = v44
                                case "":
                                    v45 := GenS2{F0: 9223372036854775807, F1: 2147483648, F2: genFn0(((1.5 * float64(0.0)) >= (((float64(3.14159) / 0.1) * -2.25) * 1e10)))}
                                    _ = v45
                                case "\n\t":
                                    if len(v17) > 0 {
                                        v17[0] = -97
                                    }
                                default:
                                    switch (fi15 << ((42) & 63)) {
                                    case 42:
                                        var v46 int8 = -128
                                        _ = v46
                                    default:
                                        v17 = append(v17, fi15)
                                    }
                                }
                            }
                            p0 = false
                        } else {
                            var v47 string = (("a" + (("a" + "\n\t") + ("hello" + "hello"))) + "a")
                            fmt.Println(fi15, fi18, fi18)
                            _ = v47
                        }
                    }
                    fmt.Println(fi15, fi18)
                    v17 = append(v17, 9223372036854775807)
                    _ = fi18
                }
                _ = fi15
                _ = v17
            }
            p0 = p0
            var v48 uint32 = 324
            _ = v48
        }
        switch "hello" {
        case "\n\t":
            for fi49 := 0; fi49 < 3; fi49++ {
                switch "a" {
                case "foo bar":
                    for fi50 := 0; fi50 < 5; fi50++ {
                        switch (fi50 % ((fi50) | 1)) {
                        case -353:
                            if (!((2 == 4294967295) && false)) {
                                v51 := []int{(fi49 - fi49), (fi50 & genFn0(p0)), fi49, 32768}
                                _ = v51
                            }
                        }
                        p0 = (0 == 131)
                        _ = fi50
                    }
                case "\n\t":
                    for fi52 := 0; fi52 < 20; fi52++ {
                        v53 := GenS1{F0: -620, F1: true}
                        _ = fi52
                        _ = v53
                    }
                default:
                    if p0 {
                        for fi54 := 0; fi54 < 18; fi54++ {
                            var v55 bool = (-62 >= -128)
                            v55 = ((fi54 & fi49) >= fi54)
                            p0 = (p0 && (true || v55))
                            _ = fi54
                            _ = v55
                        }
                        v56 := []int{}
                        for fi57 := 0; fi57 < 18; fi57++ {
                            p0 = (42 == 0)
                            p0 = p0
                            if true {
                                p0 = false
                                switch fi49 {
                                case 127:
                                    fmt.Println(p0)
                                }
                                p0 = (((-2.25 * 0.0) <= float32(0.0)) && p0)
                            } else {
                                if (0.1 < -2.25) {
                                    if false {
                                        var v58 bool = p0
                                        var v59 int16 = 860
                                        _ = v58
                                        _ = v59
                                    }
                                    if (1.5 < 0.1) {
                                        p0 = p0
                                        if true {
                                            p0 = ((-15 <= -1) && p0)
                                            if (p0 || (3.14159 < 1.5)) {
                                                p0 = (225 > 127)
                                            } else {
                                                if p0 {
                                                    v60 := GenS3{F0: 341, F1: p0, F2: -128}
                                                    if (-943 <= -1) {
                                                        var v61 int8 = 57
                                                        _ = v61
                                                    } else {
                                                        switch genFn0((!(!(0 > 607)))) {
                                                        case 620:
                                                            v62 := GenS4{F0: ("" + "")}
                                                            _ = v62
                                                        case 32768:
                                                            v56 = append(v56, 4294967295)
                                                        case 2147483647:
                                                            if (42 > 0) {
                                                                if p0 {
                                                                    var v63 bool = (42 == 346)
                                                                    var v64 uint8 = 0
                                                                    v65 := GenS3{F0: -32768, F1: (42 > fi57), F2: 1}
                                                                    _ = v63
                                                                    _ = v64
                                                                    _ = v65
                                                                } else {
                                                                    var v66 string = ""
                                                                    p0 = (88 >= 0)
                                                                    var v67 int = (fi49 + 42)
                                                                    v67 = fi49
                                                                    _ = v66
                                                                    _ = v67
                                                                }
                                                                fmt.Println(p0)
                                                            } else {
                                                                var v68 int32 = 854
                                                                var v69 int64 = -328
                                                                _ = v68
                                                                _ = v69
                                                            }
                                                        }
                                                        var v70 int = (fi57 * fi57)
                                                        _ = len(v56)
                                                        _ = v70
                                                    }
                                                    var v71 string = "foo bar"
                                                    _ = v60
                                                    _ = v71
                                                } else {
                                                    p0 = p0
                                                }
                                                var v72 uint = 259
                                                var v73 int16 = -32768
                                                var v74 uint64 = 2
                                                _ = v72
                                                _ = v73
                                                _ = v74
                                            }
                                            var v75 bool = p0
                                            var v76 int16 = -44
                                            _ = v75
                                            _ = v76
                                        } else {
                                            p0 = p0
                                        }
                                    }
                                    fmt.Println(fi49, fi57)
                                    if (p0 || p0) {
                                        var v77 uint32 = 2147483647
                                        _ = v77
                                    } else {
                                        var v78 uint = 534
                                        v78 -= v78
                                        if false {
                                            switch fi49 {
                                            case 936:
                                                var v79 uint = 699
                                                _ = v79
                                            case 127:
                                                if (((819 != 634) && (float32((3.14159 / (1.5 * 3.14159))) >= (1.5 * 1e10))) || (!(941 != 752))) {
                                                    switch genFn0((!(-9223372036854775808 < 0))) {
                                                    case -1:
                                                        p0 = (!p0)
                                                    case 623:
                                                        v80 := GenS5{F0: 1, F1: 127, F2: 127}
                                                        _ = v80
                                                    default:
                                                        switch "a" {
                                                        case "":
                                                            v81 := GenS6{F0: -44, F1: -917, F2: 987}
                                                            fmt.Println(fi57, fi57, v78)
                                                            _ = v81
                                                        case "foo bar":
                                                            if p0 {
                                                                v82 := GenS5{F0: 0, F1: 100, F2: 7}
                                                                _ = v82
                                                            }
                                                        }
                                                    }
                                                    switch fi49 {
                                                    case 127:
                                                        fmt.Println(fi49, fi57)
                                                    }
                                                    v83 := GenS7{F0: 0, F1: 42, F2: 1}
                                                    _ = v83
                                                }
                                            case -185:
                                                var v84 uint64 = 42
                                                _ = v84
                                            }
                                            if p0 {
                                                v78 = 2
                                                p0 = p0
                                                var v85 uint32 = 4294967295
                                                _ = v85
                                            } else {
                                                fmt.Println(v78, fi57)
                                                switch (fi49 & int(v78)) {
                                                case 32768:
                                                    v86 := GenS2{F0: 1, F1: 503, F2: fi57}
                                                    _ = v86
                                                }
                                            }
                                            v87 := GenS5{F0: 128, F1: 127, F2: 255}
                                            _ = v87
                                        } else {
                                            if p0 {
                                                fmt.Println(fi49)
                                            }
                                            if (806 < (fi49 - fi49)) {
                                                switch (-(fi49)) {
                                                case 1:
                                                    switch fi49 {
                                                    case -163:
                                                        var v88 int8 = 0
                                                        _ = v88
                                                    case -9223372036854775808:
                                                        var v89 float64 = 0.0
                                                        if true {
                                                            v56 = append(v56, (fi49 * fi49))
                                                            var v90 uint16 = 42
                                                            _ = v90
                                                        }
                                                        v91 := GenS8{F0: true}
                                                        _ = v89
                                                        _ = v91
                                                    default:
                                                        _ = len(v56)
                                                    }
                                                case 127:
                                                    if len(v56) > 0 {
                                                        v56[0] = fi49
                                                    }
                                                case -2147483648:
                                                    fmt.Println(p0, v78)
                                                default:
                                                    var v92 string = ("a" + "foo bar")
                                                    _ = v92
                                                }
                                            }
                                        }
                                        v93 := GenS1{F0: 1, F1: (true && p0)}
                                        switch (fi49 ^ (fi57 - -458)) {
                                        case -9223372036854775808:
                                            if true {
                                                fmt.Println(p0, fi57, p0)
                                                switch "hello" {
                                                case "foo bar":
                                                    var v94 int64 = -825
                                                    _ = v94
                                                case "":
                                                    v78 = (v78 + 573)
                                                case "a":
                                                    _ = len(v56)
                                                default:
                                                    p0 = false
                                                }
                                                switch "\n\t" {
                                                case "\n\t":
                                                    if len(v56) > 0 {
                                                        v56[0] = (fi57 | 1)
                                                    }
                                                case "foo bar":
                                                    v56 = append(v56, 32767)
                                                }
                                            } else {
                                                fmt.Println(v78, v78)
                                                var v95 string = "\n\t"
                                                fmt.Println(p0)
                                                _ = v95
                                            }
                                        case -1:
                                            p0 = p0
                                        }
                                        _ = v78
                                        _ = v93
                                    }
                                    if len(v56) > 0 {
                                        v56[0] = (fi57 & fi57)
                                    }
                                    if p0 {
                                        var v96 string = (((("" + ("" + "hello")) + ("foo bar" + (((("a" + (("\n\t" + "foo bar") + "a")) + ("" + "\n\t")) + "a") + "hello"))) + ("hello" + "foo bar")) + ("a" + ("" + ("" + "foo bar"))))
                                        _ = v96
                                    }
                                    p0 = false
                                    if len(v56) >= 2 {
                                        v56 = v56[1:len(v56)]
                                    }
                                    p0 = true
                                } else {
                                    if len(v56) > 0 {
                                        v56[0] = fi57
                                    }
                                }
                            }
                            _ = fi57
                        }
                        _ = v56
                    }
                }
                _ = fi49
            }
        case "hello":
            if p0 {
                if p0 {
                    fmt.Println(p0)
                    p0 = false
                    v97 := []int{-1, genFn0(((630 > 809) && p0))}
                    _ = v97
                } else {
                    if (-689 == 583) {
                        v98 := []int{4294967295, 32767, 42}
                        for fi99 := 0; fi99 < 5; fi99++ {
                            for fi100 := 0; fi100 < 8; fi100++ {
                                v101 := GenS7{F0: 2, F1: -1, F2: 42}
                                if (-681 > 0) {
                                    var v102 uint16 = 128
                                    var v103 string = "\n\t"
                                    _ = v102
                                    _ = v103
                                }
                                v104 := GenS4{F0: (("" + (("\n\t" + ("\n\t" + ("a" + (("foo bar" + "hello") + ("foo bar" + "a"))))) + ("hello" + "\n\t"))) + "")}
                                _ = fi100
                                _ = v101
                                _ = v104
                            }
                            _ = fi99
                        }
                        v105 := GenS9{F0: 255, F1: 4294967295}
                        _ = v105
                        _ = v98
                    } else {
                        var v106 uint32 = 65535
                        _ = v106
                    }
                }
                fmt.Println(p0)
                switch "a" {
                case "foo bar":
                    if ((p0 && (1 < 4294967295)) && (1 > 1)) {
                        if (1.5 == 3.14159) {
                            v107 := GenS2{F0: 686, F1: 4294967295, F2: 2147483647}
                            var v108 bool = p0
                            _ = v107
                            _ = v108
                        }
                        switch "hello" {
                        case "foo bar":
                            var v109 int16 = 1
                            _ = v109
                        case "":
                            if true {
                                if p0 {
                                    v110 := []int{}
                                    var v111 bool = false
                                    v112 := GenS2{F0: 107, F1: 2147483647, F2: 143}
                                    _ = v110
                                    _ = v111
                                    _ = v112
                                }
                                for fi113 := 0; fi113 < 3; fi113++ {
                                    var v114 uint8 = 31
                                    _ = fi113
                                    _ = v114
                                }
                            } else {
                                switch "foo bar" {
                                case "\n\t":
                                    if (127 == -1) {
                                        v115 := GenS7{F0: 851, F1: 42, F2: -1}
                                        var v116 uint = 68
                                        v116 = 1
                                        _ = v115
                                        _ = v116
                                    } else {
                                        fmt.Println(p0)
                                    }
                                case "hello":
                                    if (0.1 != 0.1) {
                                        p0 = (("\n\t" == "hello") || p0)
                                    }
                                }
                                p0 = (-1 > 50)
                            }
                        case "\n\t":
                            switch 2147483647 {
                            case 2147483647:
                                v117 := GenS10{F0: -32768, F1: 127}
                                _ = v117
                            case -9223372036854775808:
                                var v118 float64 = (0.1 + ((-2.25 / 1.0) * -2.25))
                                _ = v118
                            case 9223372036854775807:
                                v119 := GenS9{F0: 0, F1: 42}
                                _ = v119
                            default:
                                v120 := []int{genFn0(p0), 1, 0}
                                _ = v120
                            }
                        default:
                            var v121 string = ""
                            _ = v121
                        }
                    } else {
                        for fi122 := 0; fi122 < 4; fi122++ {
                            if p0 {
                                for fi123 := 0; fi123 < 16; fi123++ {
                                    p0 = (!(p0 || p0))
                                    _ = fi123
                                }
                                var v124 int = fi122
                                var v125 int32 = 732
                                _ = v124
                                _ = v125
                            }
                            switch (fi122 ^ fi122) {
                            case 949:
                                var v126 uint64 = 4294967295
                                _ = v126
                            }
                            _ = fi122
                        }
                        if true {
                            v127 := GenS8{F0: (0 == -128)}
                            p0 = p0
                            var v128 float64 = 1e10
                            _ = v127
                            _ = v128
                        } else {
                            var v129 uint32 = 42
                            _ = v129
                        }
                        var v130 int8 = -128
                        _ = v130
                    }
                case "\n\t":
                    var v131 uint16 = 350
                    _ = v131
                case "a":
                    var v132 int = 32768
                    _ = v132
                default:
                    fmt.Println(p0)
                }
            }
        default:
            for fi133 := 0; fi133 < 5; fi133++ {
                if false {
                    if (((fi133 ^ fi133) > (fi133 | fi133)) || p0) {
                        for fi134 := 0; fi134 < 10; fi134++ {
                            if (0 != 127) {
                                var v135 bool = ((-495 > 42) && (true && (2 == 4294967295)))
                                _ = v135
                            }
                            p0 = p0
                            _ = fi134
                        }
                        if ("" != "foo bar") {
                            v136 := []int{(-(fi133)), fi133, fi133, fi133}
                            p0 = (!(!p0))
                            var v137 int64 = 813
                            _ = v136
                            _ = v137
                        } else {
                            var v138 uint = 997
                            _ = v138
                        }
                    }
                    var v139 int16 = 463
                    _ = v139
                }
                _ = fi133
            }
        }
        _ = v14
    }
    p0 = true
    _ = p0
    return 5
}

func main() {
    var v1 float32 = (1.5 / (((0.0 + float32(3.14159)) * (1.5 - 0.0)) * 0.1))
    var v2 int8 = 0
    var v3 int64 = 1
    v140 := []int{-2147483648, genFn0(false), -453, 32767}
    v2 += (v2 + v2)
    var v141 int32 = 2147483647
    if (!false) {
        var v142 int32 = v141
        fmt.Println(v1, v3, v141)
        _ = v142
    } else {
        fmt.Println(v141)
    }
    switch "hello" {
    case "\n\t":
        var v143 int = 42
        _ = v143
    case "foo bar":
        fmt.Println(v1, v3)
    default:
        _ = len(v140)
    }
    v144 := GenS11{F0: 720, F1: 65535, F2: 1}
    fmt.Println(v3, v3)
    v3 += (v3 ^ v3)
    for fi145 := 0; fi145 < 16; fi145++ {
        var v146 uint16 = 0
        _ = fi145
        _ = v146
    }
    if ((false || (!true)) || true) {
        v141 = -403
        var v147 string = ("foo bar" + "foo bar")
        _ = v147
    } else {
        if len(v140) > 0 {
            v140[0] = 4294967295
        }
    }
    for fi148 := 0; fi148 < 3; fi148++ {
        if ((v141 * v141) != (-(v141))) {
            fmt.Println(fi148, v2)
            if true {
                if true {
                    switch ((("a" + "\n\t") + "") + ("hello" + "")) {
                    case "foo bar":
                        v3 -= (v3 >> ((797) & 63))
                    default:
                        v2 = v2
                    }
                    v1 = (v1 * v1)
                }
                if ("\n\t" < "a") {
                    switch "\n\t" {
                    case "foo bar":
                        v141 = 161
                    case "hello":
                        v3 = (v3 * (v3 ^ (v3 / (((^(v3))) | 1))))
                    case "a":
                        if len(v140) >= 2 {
                            v140 = v140[1:len(v140)]
                        }
                    }
                }
            }
            v149 := GenS3{F0: 1000, F1: ((2 <= 375) && false), F2: (v2 / (((v2 ^ 78)) | 1))}
            _ = v149
        }
        if len(v140) > 0 {
            v140[0] = (-(fi148))
        }
        var v150 int16 = -32768
        _ = fi148
        _ = v150
    }
    var v151 bool = true
    if (!v151) {
        switch "" {
        case "\n\t":
            var v152 int8 = v2
            _ = v152
        case "a":
            for fi153 := 0; fi153 < 9; fi153++ {
                v141 -= v141
                v151 = true
                _ = fi153
            }
        }
        if (-475 >= 32767) {
            fmt.Println(v141)
            if ((!v151) || (!false)) {
                fmt.Println(v141)
            } else {
                fmt.Println(v141, v141)
                fmt.Println(v2, v141, v1)
            }
        } else {
            if v151 {
                for fi154 := 0; fi154 < 11; fi154++ {
                    if (!((-(v141)) <= v141)) {
                        switch "\n\t" {
                        case "foo bar":
                            if true {
                                if v151 {
                                    switch "a" {
                                    case "foo bar":
                                        switch (fi154 / ((fi154) | 1)) {
                                        case 9223372036854775807:
                                            v1 -= v1
                                        case -1:
                                            for fi155 := 0; fi155 < 19; fi155++ {
                                                v1 = v1
                                                if (!v151) {
                                                    if v151 {
                                                        v156 := GenS1{F0: 0, F1: v151}
                                                        fmt.Println(fi154, v151)
                                                        _ = v156
                                                    }
                                                } else {
                                                    v2 = (v2 % (((v2 % (((v2 / (((v2 & v2)) | 1))) | 1))) | 1))
                                                    var v157 string = ("a" + (("\n\t" + ("a" + "")) + "hello"))
                                                    v158 := GenS12{F0: 255, F1: 610, F2: v157}
                                                    _ = v157
                                                    _ = v158
                                                }
                                                _ = fi155
                                            }
                                        case 1:
                                            if (v2 > -31) {
                                                v140 = append(v140, (fi154 + (fi154 | fi154)))
                                                for fi159 := 0; fi159 < 7; fi159++ {
                                                    _ = len(v140)
                                                    v2 -= (v2 & -1)
                                                    _ = fi159
                                                }
                                            } else {
                                                var v160 float64 = 1e10
                                                if len(v140) > 0 {
                                                    v140[0] = fi154
                                                }
                                                _ = v160
                                            }
                                        default:
                                            fmt.Println(v2, v2, v3)
                                        }
                                    case "\n\t":
                                        v141 = (v141 + int32(v2))
                                    default:
                                        switch genFn0(v151) {
                                        case 32768:
                                            v2 -= (v2 & v2)
                                        }
                                    }
                                    switch (fi154 ^ fi154) {
                                    case 42:
                                        fmt.Println(fi154, fi154, v141)
                                    case 9223372036854775807:
                                        v140 = append(v140, (^(fi154)))
                                    default:
                                        v151 = v151
                                    }
                                    var v161 int32 = 0
                                    _ = v161
                                } else {
                                    v141 = v141
                                    fmt.Println(v141)
                                    v141 -= v141
                                }
                            } else {
                                v3 += (v3 + -906)
                            }
                        case "\n\t":
                            var v162 float64 = (float64(v1) - ((1.5 * 0.1) + -2.25))
                            _ = v162
                        default:
                            var v163 string = ("" + ("a" + (("a" + ("" + "a")) + (("a" + "") + "a"))))
                            _ = v163
                        }
                        v141 = v141
                    }
                    for fi164 := 0; fi164 < 8; fi164++ {
                        var v165 float64 = 0.0
                        _ = fi164
                        _ = v165
                    }
                    fmt.Println(v151, v141)
                    _ = fi154
                }
            } else {
                var v166 int32 = v141
                var v167 int = -2147483648
                v168 := GenS7{F0: 4294967295, F1: (v2 & 1), F2: (v141 ^ 2147483647)}
                _ = v166
                _ = v167
                _ = v168
            }
        }
        v169 := GenS4{F0: ("\n\t" + "foo bar")}
        _ = v169
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v140:", len(v140), v140)
    fmt.Println("v141:", v141)
    fmt.Println("v151:", v151)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v140
    _ = v141
    _ = v144
    _ = v151
    _ = v2
    _ = v3
}
