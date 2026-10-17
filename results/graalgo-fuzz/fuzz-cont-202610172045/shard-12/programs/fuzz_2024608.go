package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int; F2 bool }
type GenS2 struct { F0 int8; F1 int64 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 string; F1 uint16 }
type GenS5 struct { F0 uint16; F1 uint8 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 uint64; F1 bool; F2 uint8 }
type GenS9 struct { F0 int }
type GenS10 struct { F0 uint; F1 int8; F2 uint32 }
type GenS11 struct { F0 uint64 }
type GenS12 struct { F0 int16; F1 int8; F2 bool }
type GenS13 struct { F0 int32 }
type GenS14 struct { F0 uint16; F1 bool; F2 int64 }
type GenS15 struct { F0 int; F1 int64; F2 uint8 }

func main() {
    var v1 string = ("a" + ("" + "hello"))
    v2 := []int{-470, 128}
    v3 := make(map[int]string)
    var v4 string = ((v1 + v1) + ("hello" + v1))
    for fi5 := 0; fi5 < 10; fi5++ {
        if (255 <= 42) {
            var v6 uint16 = 42
            _ = v6
        } else {
            v7 := GenS1{F0: 32767, F1: 127, F2: true}
            v8 := GenS1{F0: 255, F1: (fi5 >> ((42) & 63)), F2: ((42 > 973) && (true || (false && false)))}
            _ = v7
            _ = v8
        }
        v9 := GenS2{F0: 42, F1: -1}
        _ = fi5
        _ = v9
    }
    v1 = "a"
    if true {
        var v10 float32 = float32(1.5)
        if len(v2) > 0 {
            v2[0] = 32768
        }
        if true {
            v3[9223372036854775807] = "hello"
            v11 := GenS3{F0: false}
            _ = v11
        }
        _ = v10
    }
    v12 := GenS3{F0: ((!false) && true)}
    for fi13 := 0; fi13 < 20; fi13++ {
        v1 = "\n\t"
        for fi14 := 0; fi14 < 13; fi14++ {
            if false {
                v4 = v4
                var v15 uint32 = 65535
                _ = v15
            }
            if (false && true) {
                var v16 float32 = ((float32(3.14159) + 0.1) / 1.0)
                switch ((v1 + "foo bar") + ((v1 + (v1 + (("hello" + v4) + (v4 + ("a" + v1))))) + "")) {
                case "hello":
                    if len(v2) > 0 {
                        v2[0] = (-(fi14))
                    }
                }
                _ = v16
            } else {
                var v17 uint = 1
                _ = v17
            }
            if (154 >= -9223372036854775808) {
                switch "\n\t" {
                case "a":
                    v18 := GenS2{F0: 99, F1: 1}
                    if (true || ((!false) || (((-32768 <= 42) && (!true)) && false))) {
                        v4 = (v1 + (v1 + v4))
                        var v19 int8 = 100
                        if true {
                            var v20 float32 = (1e10 * (1.5 - 0.1))
                            switch fi13 {
                            case 127:
                                var v21 int8 = v19
                                _ = v21
                            }
                            _ = v20
                        } else {
                            v22 := GenS1{F0: 0, F1: fi14, F2: (!false)}
                            v3[-550] = v4
                            v19 = (v19 & v19)
                            _ = v22
                        }
                        _ = v19
                    } else {
                        fmt.Println(fi13, v1)
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        var v23 int64 = 9223372036854775807
                        switch (v4 + "foo bar") {
                        case "\n\t":
                            fmt.Println(v1, v4)
                            v4 = ("hello" + v1)
                            if false {
                                var v24 bool = false
                                fmt.Println(fi14, v4)
                                _ = v24
                            } else {
                                if true {
                                    switch "" {
                                    case "hello":
                                        switch ("" + v4) {
                                        case "\n\t":
                                            var v25 uint16 = 32767
                                            _ = v25
                                        default:
                                            v4 = v1
                                        }
                                    default:
                                        var v26 uint64 = 1
                                        _ = v26
                                    }
                                    v4 = "a"
                                    v27 := GenS2{F0: -128, F1: v23}
                                    v28 := GenS2{F0: 62, F1: (v23 / ((0) | 1))}
                                    _ = v27
                                    _ = v28
                                }
                            }
                        case "foo bar":
                            fmt.Println(v23, fi13, v23)
                        case "a":
                            v29 := GenS4{F0: v4, F1: 2}
                            _ = v29
                        }
                        if len(v2) > 0 {
                            v2[0] = fi14
                        }
                        fmt.Println(v1, v1, fi14)
                        v23 += (v23 >> ((4294967295) & 63))
                        v1 = "a"
                        if (!false) {
                            v4 = "foo bar"
                        }
                        var v30 uint = 4294967295
                        if (1.5 > (-2.25 * 1.5)) {
                            v31 := GenS5{F0: 750, F1: 1}
                            var v32 uint8 = 42
                            _ = v31
                            _ = v32
                        } else {
                            fmt.Println(v23, v30)
                            v33 := GenS4{F0: "", F1: 285}
                            switch "hello" {
                            case "a":
                                delete(v3, fi13)
                            }
                            _ = v33
                        }
                        _ = v23
                        _ = v30
                    }
                    v4 = v4
                    _ = v18
                }
                fmt.Println(fi14)
            } else {
                switch -875 {
                case 315:
                    v34 := GenS3{F0: true}
                    _ = v34
                default:
                    if ("hello" <= (v4 + (v4 + v4))) {
                        if true {
                            v1 = v1
                            fmt.Println(fi13)
                        }
                        if true {
                            fmt.Println(fi14)
                        }
                        var v35 int32 = 2147483647
                        var v36 int = fi14
                        _ = v35
                        _ = v36
                    } else {
                        var v37 uint8 = 255
                        _ = v37
                    }
                }
                var v38 float64 = (-2.25 - ((0.0 - 1e10) * 0.1))
                _ = v38
            }
            _ = fi14
        }
        var v39 uint32 = 0
        _ = fi13
        _ = v39
    }
    switch (v1 + (("\n\t" + (v1 + v1)) + (("" + v4) + "a"))) {
    case "hello":
        fmt.Println(v1, v1)
    case "\n\t":
        v4 = "a"
    }
    v1 = v4
    v40 := GenS4{F0: v4, F1: 255}
    if true {
        fmt.Println(v4)
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    v41 := v3[999]
    var v42 int = 9223372036854775807
    for fi43 := 0; fi43 < 13; fi43++ {
        v41 = ("" + v1)
        var v44 uint32 = 0
        switch v42 {
        case 603:
            for fi45 := 0; fi45 < 3; fi45++ {
                switch fi43 {
                case -9223372036854775808:
                    switch (v1 + (v1 + "\n\t")) {
                    case "":
                        v1 = v1
                    case "foo bar":
                        v42 -= fi45
                    case "a":
                        var v46 uint32 = v44
                        _ = v46
                    }
                }
                _ = fi45
            }
        case 32767:
            if (128 <= 42) {
                v44 = 0
            }
        }
        _ = fi43
        _ = v44
    }
    switch v42 {
    case 127:
        v47 := GenS4{F0: v41, F1: 823}
        _ = v47
    }
    if ((("a" + "") + ((v1 + (v41 + v41)) + ("\n\t" + ("foo bar" + v1)))) < (("foo bar" + v1) + v41)) {
        for fi48 := 0; fi48 < 18; fi48++ {
            fmt.Println(v4, fi48)
            if (127 < -1) {
                if (!true) {
                    var v49 float64 = (((1e10 / float64((0.1 - (0.0 / ((0.0 + 1e10) / 1e10))))) - -2.25) / (-2.25 / (1.5 / -2.25)))
                    fmt.Println(v4)
                    _ = v49
                } else {
                    fmt.Println(v1, v42)
                    if false {
                        if true {
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            v1 = "hello"
                            for fi50 := 0; fi50 < 18; fi50++ {
                                v1 = ("foo bar" + v4)
                                v1 = (v1 + v1)
                                v51 := GenS6{F0: 231}
                                _ = fi50
                                _ = v51
                            }
                        }
                        switch v42 {
                        case 1:
                            var v52 float32 = ((3.14159 * 3.14159) + float32(0.1))
                            _ = v52
                        case 42:
                            fmt.Println(v42, v41, v41)
                        }
                        for fi53 := 0; fi53 < 15; fi53++ {
                            if true {
                                if true {
                                    v41 = ("\n\t" + "foo bar")
                                    v54 := GenS3{F0: false}
                                    if false {
                                        var v55 int16 = 547
                                        fmt.Println(v41)
                                        if false {
                                            fmt.Println(v1)
                                            v56 := GenS6{F0: 42}
                                            _ = v56
                                        }
                                        _ = v55
                                    }
                                    var v57 int8 = 127
                                    _ = v54
                                    _ = v57
                                } else {
                                    delete(v3, (fi48 << ((0) & 63)))
                                    switch 32768 {
                                    case -1:
                                        if true {
                                            if false {
                                                v58 := GenS1{F0: 65535, F1: (^(fi53)), F2: (0 <= 1)}
                                                v42 = v42
                                                v59 := v3[-2147483648]
                                                _ = v58
                                                _ = v59
                                            }
                                        }
                                    case 364:
                                        var v60 string = v41
                                        _ = v60
                                    }
                                }
                                fmt.Println(v1)
                            }
                            v4 = v1
                            v4 = "foo bar"
                            _ = fi53
                        }
                    } else {
                        delete(v3, v42)
                        v61 := GenS5{F0: 65535, F1: 255}
                        _ = v61
                    }
                }
            } else {
                for fi62 := 0; fi62 < 18; fi62++ {
                    v63 := GenS7{F0: 65535}
                    _ = fi62
                    _ = v63
                }
                switch ((("\n\t" + "") + v4) + v1) {
                case "":
                    if (-42 < 840) {
                        if false {
                            for fi64 := 0; fi64 < 4; fi64++ {
                                if false {
                                    v4 = "\n\t"
                                    var v65 uint = 42
                                    v41 = ("" + v4)
                                    _ = v65
                                } else {
                                    v3[(fi48 & fi64)] = "a"
                                }
                                _ = fi64
                            }
                            for fi66 := 0; fi66 < 10; fi66++ {
                                _ = len(v2)
                                var v67 uint16 = 571
                                var v68 string = v41
                                _ = fi66
                                _ = v67
                                _ = v68
                            }
                        } else {
                            for fi69 := 0; fi69 < 17; fi69++ {
                                fmt.Println(v41, v41, v42)
                                switch ((v41 + "hello") + v1) {
                                case "":
                                    var v70 float32 = 0.0
                                    _ = v70
                                case "hello":
                                    v42 += 0
                                case "foo bar":
                                    v71 := GenS8{F0: 18446744073709551615, F1: (!false), F2: 1}
                                    _ = v71
                                }
                                _ = fi69
                            }
                            fmt.Println(v1)
                        }
                        if (!true) {
                            switch (v42 - (fi48 % ((-1) | 1))) {
                            case 9223372036854775807:
                                var v72 int16 = -1
                                _ = v72
                            case 0:
                                v73 := GenS5{F0: 255, F1: 2}
                                _ = v73
                            }
                            if true {
                                switch v1 {
                                case "\n\t":
                                    var v74 uint = 284
                                    _ = v74
                                case "":
                                    var v75 float32 = (0.0 * 0.0)
                                    _ = v75
                                case "a":
                                    v4 = v4
                                }
                            }
                            var v76 uint64 = 896
                            _ = v76
                        }
                    }
                case "foo bar":
                    for fi77 := 0; fi77 < 19; fi77++ {
                        if true {
                            v4 = (v41 + ("" + v1))
                        } else {
                            v41 = v4
                            v2 = append(v2, -119)
                            switch (v1 + v41) {
                            case "\n\t":
                                v42 = fi77
                            case "":
                                v78 := GenS9{F0: (fi48 ^ fi48)}
                                _ = v78
                            case "a":
                                v1 = v4
                            }
                        }
                        v79 := v3[fi77]
                        v80 := GenS10{F0: 0, F1: 100, F2: 2}
                        v1 = "hello"
                        switch "a" {
                        case "hello":
                            switch (v42 * (fi77 | fi48)) {
                            case -689:
                                var v81 float64 = (0.1 / 3.14159)
                                v82 := GenS5{F0: 1, F1: 42}
                                v83 := v3[fi77]
                                _ = v81
                                _ = v82
                                _ = v83
                            default:
                                v84 := GenS3{F0: false}
                                _ = v84
                            }
                        case "a":
                            if false {
                                var v85 uint32 = 4294967295
                                _ = v85
                            } else {
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                fmt.Println(v42, v4)
                            }
                        default:
                            var v86 string = v79
                            _ = v86
                        }
                        _ = fi77
                        _ = v79
                        _ = v80
                    }
                case "\n\t":
                    fmt.Println(v4)
                }
                switch v1 {
                case "\n\t":
                    fmt.Println(v4)
                case "foo bar":
                    if len(v2) > 0 {
                        v2[0] = (fi48 ^ (^(v42)))
                    }
                }
            }
            switch "\n\t" {
            case "":
                if true {
                    var v87 int32 = -1
                    _ = v87
                } else {
                    var v88 int64 = -1
                    if (668 > 548) {
                        if (2147483647 >= fi48) {
                            fmt.Println(fi48)
                            var v89 uint8 = 255
                            _ = v89
                        } else {
                            v4 = "foo bar"
                            if (true && true) {
                                v90 := GenS11{F0: 42}
                                var v91 uint64 = 4294967295
                                _ = v90
                                _ = v91
                            } else {
                                v41 = ("" + ("\n\t" + "a"))
                                if (!true) {
                                    for fi92 := 0; fi92 < 11; fi92++ {
                                        fmt.Println(fi48)
                                        var v93 int8 = 100
                                        switch ("" + v4) {
                                        case "hello":
                                            v41 = v4
                                        case "\n\t":
                                            if (-2.25 >= (1e10 * ((1.5 - 1.5) * 0.1))) {
                                                if (-2.25 == (3.14159 - (0.1 * ((((-2.25 * -2.25) / ((0.0 + (-2.25 * 3.14159)) / -2.25)) - 0.1) / 3.14159)))) {
                                                    v94 := GenS5{F0: 32767, F1: 2}
                                                    v41 = v4
                                                    _ = v94
                                                } else {
                                                    fmt.Println(v88, fi48, v93)
                                                    if (2 == 0) {
                                                        switch 4294967295 {
                                                        case 2147483647:
                                                            var v95 int64 = (v88 * v88)
                                                            _ = v95
                                                        }
                                                        if false {
                                                            if true {
                                                                delete(v3, (v42 - v42))
                                                                var v96 string = v4
                                                                _ = v96
                                                            } else {
                                                                if ((fi48 | fi48) != (-(v42))) {
                                                                    var v97 bool = (30 == 1000)
                                                                    fmt.Println(fi48, v42, v93)
                                                                    _ = v97
                                                                }
                                                                var v98 bool = false
                                                                _ = v98
                                                            }
                                                        } else {
                                                            fmt.Println(v93, v88)
                                                            fmt.Println(v93, v93)
                                                            switch "hello" {
                                                            case "a":
                                                                var v99 int64 = v88
                                                                var v100 int64 = (v88 / (((v99 / ((v88) | 1))) | 1))
                                                                v88 -= v99
                                                                _ = v100
                                                                _ = v99
                                                            case "":
                                                                v101 := GenS12{F0: 1000, F1: -128, F2: false}
                                                                v102 := GenS11{F0: 4294967295}
                                                                _ = v101
                                                                _ = v102
                                                            }
                                                            delete(v3, -2147483648)
                                                            var v103 uint = 65535
                                                            _ = v103
                                                        }
                                                    } else {
                                                        var v104 uint32 = 2147483648
                                                        _ = v104
                                                    }
                                                }
                                                var v105 float64 = float64((3.14159 - (0.0 / 1.0)))
                                                v1 = "hello"
                                                switch (fi48 / ((fi48) | 1)) {
                                                case -197:
                                                    var v106 uint64 = 400
                                                    _ = v106
                                                case -2147483648:
                                                    v88 = v88
                                                default:
                                                    switch v41 {
                                                    case "":
                                                        var v107 uint32 = 346
                                                        _ = v107
                                                    case "\n\t":
                                                        if ((fi92 * (fi92 << ((255) & 63))) > fi92) {
                                                            if false {
                                                                if true {
                                                                    v1 = "hello"
                                                                    fmt.Println(fi92)
                                                                } else {
                                                                    v108 := GenS9{F0: fi48}
                                                                    v93 += v93
                                                                    v2 = append(v2, -736)
                                                                    if true {
                                                                        var v109 int8 = (v93 | v93)
                                                                        _ = v109
                                                                    } else {
                                                                        switch fi92 {
                                                                        case -2147483648:
                                                                            v1 = v4
                                                                        case -9223372036854775808:
                                                                            var v110 int8 = 98
                                                                            _ = v110
                                                                        case 864:
                                                                            v111 := GenS13{F0: -126}
                                                                            _ = v111
                                                                        default:
                                                                            var v112 uint32 = 1
                                                                            _ = v112
                                                                        }
                                                                        fmt.Println(v93, v4, fi92)
                                                                    }
                                                                    _ = v108
                                                                }
                                                                fmt.Println(v4)
                                                            } else {
                                                                if len(v2) >= 2 {
                                                                    v2 = v2[1:len(v2)]
                                                                }
                                                                switch v41 {
                                                                case "a":
                                                                    v113 := v3[(fi48 + (fi92 - (fi92 % ((-9223372036854775808) | 1))))]
                                                                    _ = v113
                                                                case "foo bar":
                                                                    var v114 uint = 255
                                                                    _ = v114
                                                                default:
                                                                    v115 := GenS14{F0: 208, F1: (2 > 589), F2: (v88 >> ((766) & 63))}
                                                                    _ = v115
                                                                }
                                                                fmt.Println(v93, v41)
                                                            }
                                                        } else {
                                                            v116 := GenS13{F0: -1}
                                                            var v117 int16 = 42
                                                            _ = v116
                                                            _ = v117
                                                        }
                                                    case "hello":
                                                        var v118 int = v42
                                                        _ = v118
                                                    }
                                                }
                                                v88 -= v88
                                                _ = v105
                                            } else {
                                                delete(v3, fi48)
                                                if true {
                                                    v3[128] = ("hello" + v4)
                                                    if (true || (0 < 255)) {
                                                        v119 := GenS10{F0: 18446744073709551615, F1: (v93 + (v93 * (v93 + (v93 & v93)))), F2: 879}
                                                        var v120 float64 = (1.5 + (0.0 - (-2.25 + float64(1e10))))
                                                        _ = v119
                                                        _ = v120
                                                    }
                                                }
                                                v4 = v1
                                            }
                                        case "":
                                            var v121 string = v4
                                            _ = v121
                                        default:
                                            v41 = (((v1 + v41) + v1) + v1)
                                        }
                                        fmt.Println(fi92, v4, v42)
                                        _ = fi92
                                        _ = v93
                                    }
                                    var v122 float64 = float64(((0.1 / 3.14159) + ((1e10 * ((1.5 + 1.5) + 0.1)) * -2.25)))
                                    _ = v122
                                }
                            }
                        }
                        v123 := GenS10{F0: 0, F1: 1, F2: 4294967295}
                        if (((!false) || (2 <= 191)) || false) {
                            var v124 string = v1
                            _ = v124
                        }
                        _ = v123
                    } else {
                        for fi125 := 0; fi125 < 8; fi125++ {
                            var v126 int16 = 331
                            _ = fi125
                            _ = v126
                        }
                    }
                    _ = v88
                }
            case "\n\t":
                v4 = ("\n\t" + v1)
            case "hello":
                v127 := GenS15{F0: v42, F1: 0, F2: 1}
                _ = v127
            default:
                if true {
                    v4 = v1
                    v128 := GenS4{F0: v4, F1: 315}
                    switch "foo bar" {
                    case "":
                        v41 = v41
                    }
                    _ = v128
                }
            }
            _ = fi48
        }
    } else {
        var v129 uint = 319
        _ = v129
    }
    _ = len(v3)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v42
}
