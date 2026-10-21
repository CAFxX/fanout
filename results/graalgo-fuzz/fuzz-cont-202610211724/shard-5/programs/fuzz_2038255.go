package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64; F2 int16 }
type GenS2 struct { F0 uint32; F1 string; F2 uint32 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 string; F1 bool }
type GenS5 struct { F0 uint; F1 uint }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 int16; F1 int; F2 uint64 }
type GenS8 struct { F0 int; F1 string; F2 bool }
type GenS9 struct { F0 int }
type GenS10 struct { F0 uint32; F1 int32; F2 uint32 }
type GenS11 struct { F0 int }

func main() {
    v1 := []int{9223372036854775807}
    var v2 bool = (false || false)
    v3 := []string{"a", "", ("a" + "\n\t")}
    switch ("a" + "a") {
    case "":
        var v4 int = 0
        _ = v4
    case "\n\t":
        if true {
            for fi5 := 0; fi5 < 20; fi5++ {
                for fi6 := 0; fi6 < 1; fi6++ {
                    v2 = (4294967295 >= 255)
                    fmt.Println(fi5, fi6, v2)
                    fmt.Println(v2)
                    _ = fi6
                }
                _ = fi5
            }
        } else {
            if v2 {
                var v7 int64 = 1
                if true {
                    for fi8 := 0; fi8 < 8; fi8++ {
                        switch (-(fi8)) {
                        case 32768:
                            v7 = (v7 / ((v7) | 1))
                        case 32767:
                            v9 := GenS1{F0: -2147483648, F1: 4294967295, F2: 344}
                            _ = v9
                        case 0:
                            if len(v3) > 0 {
                                v3[0] = "\n\t"
                            }
                        default:
                            if (2 != 2) {
                                if v2 {
                                    for fi10 := 0; fi10 < 2; fi10++ {
                                        v7 = v7
                                        var v11 uint16 = 42
                                        if len(v3) > 0 {
                                            v3[0] = "\n\t"
                                        }
                                        _ = fi10
                                        _ = v11
                                    }
                                    v12 := GenS1{F0: 42, F1: 1, F2: 1000}
                                    _ = v12
                                }
                                fmt.Println(fi8)
                                for fi13 := 0; fi13 < 3; fi13++ {
                                    var v14 int64 = 842
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    _ = fi13
                                    _ = v14
                                }
                            }
                        }
                        var v15 int32 = 0
                        _ = fi8
                        _ = v15
                    }
                    _ = len(v1)
                    v16 := GenS1{F0: 79, F1: 18446744073709551615, F2: 0}
                    _ = v16
                } else {
                    if (-1 >= 100) {
                        _ = len(v3)
                        v3 = append(v3, "\n\t")
                        v17 := GenS1{F0: 552, F1: 240, F2: 1000}
                        _ = v17
                    }
                    for fi18 := 0; fi18 < 8; fi18++ {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        v7 -= (v7 & v7)
                        fmt.Println(v7, v2, fi18)
                        _ = fi18
                    }
                    var v19 uint = 471
                    _ = v19
                }
                _ = v7
            } else {
                fmt.Println(v2)
            }
            v2 = (!(false && (((!(("foo bar" + "a") == "\n\t")) && (0 <= 42)) || ((3.14159 <= 1e10) && (1 <= 1)))))
        }
    case "a":
        var v20 float64 = 1e10
        _ = v20
    }
    if v2 {
        v21 := GenS2{F0: 284, F1: "hello", F2: 609}
        _ = v21
    } else {
        for fi22 := 0; fi22 < 11; fi22++ {
            fmt.Println(fi22)
            for fi23 := 0; fi23 < 3; fi23++ {
                var v24 float32 = -2.25
                if len(v3) > 0 {
                    v3[0] = ("hello" + ("a" + "hello"))
                }
                var v25 uint64 = 18446744073709551615
                _ = fi23
                _ = v24
                _ = v25
            }
            _ = fi22
        }
        if false {
            for fi26 := 0; fi26 < 6; fi26++ {
                fmt.Println(v2, v2)
                if len(v1) > 0 {
                    v1[0] = fi26
                }
                _ = fi26
            }
            for fi27 := 0; fi27 < 11; fi27++ {
                if true {
                    v2 = (!v2)
                    var v28 uint16 = 164
                    _ = v28
                } else {
                    if (v2 && false) {
                        v3 = append(v3, (("foo bar" + (("foo bar" + "a") + "foo bar")) + "a"))
                    }
                    for fi29 := 0; fi29 < 17; fi29++ {
                        switch 42 {
                        case 9223372036854775807:
                            v30 := GenS1{F0: 2147483647, F1: 2, F2: 42}
                            _ = v30
                        default:
                            v2 = false
                        }
                        fmt.Println(v2)
                        var v31 int = 128
                        fmt.Println(fi27, v31)
                        v31 = fi27
                        _ = fi29
                        _ = v31
                    }
                    v2 = v2
                }
                var v32 float32 = 0.0
                v33 := GenS3{F0: 4294967295}
                _ = fi27
                _ = v32
                _ = v33
            }
            if v2 {
                v34 := GenS1{F0: 1, F1: 4294967295, F2: 42}
                var v35 uint = 446
                _ = v34
                _ = v35
            }
        } else {
            fmt.Println(v2)
            v36 := GenS4{F0: ("a" + "a"), F1: true}
            for fi37 := 0; fi37 < 12; fi37++ {
                v3 = append(v3, "a")
                if (-2.25 >= (((0.1 - 1e10) / (1.5 * (1.5 / 1e10))) + -2.25)) {
                    fmt.Println(v2, v2)
                    var v38 uint8 = 128
                    for fi39 := 0; fi39 < 19; fi39++ {
                        if false {
                            v3 = append(v3, "hello")
                            var v40 int64 = 1
                            _ = v40
                        } else {
                            v38 = v38
                            if len(v3) > 0 {
                                v3[0] = "hello"
                            }
                            var v41 float64 = ((0.1 / 1.5) / 1.0)
                            _ = v41
                        }
                        var v42 uint8 = v38
                        v2 = (false || v2)
                        v43 := GenS4{F0: "foo bar", F1: v2}
                        if true {
                            _ = len(v1)
                            var v44 int64 = 953
                            v3 = append(v3, "foo bar")
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            var v45 float64 = (0.0 * -2.25)
                            _ = v44
                            _ = v45
                        } else {
                            var v46 int16 = 32767
                            v47 := GenS5{F0: 255, F1: 65535}
                            var v48 float64 = 1e10
                            _ = v46
                            _ = v47
                            _ = v48
                        }
                        switch ("a" + "a") {
                        case "a":
                            switch ("foo bar" + "hello") {
                            case "foo bar":
                                fmt.Println(v42, fi39, v2)
                            default:
                                if false {
                                    v3 = append(v3, "")
                                    if len(v3) > 0 {
                                        v3[0] = ("hello" + "\n\t")
                                    }
                                    if ((!(658 <= 4294967295)) && (!(42 >= -128))) {
                                        if len(v1) > 0 {
                                            v1[0] = (fi39 >> ((0) & 63))
                                        }
                                        var v49 string = ""
                                        _ = v49
                                    } else {
                                        v2 = (!true)
                                        if false {
                                            var v50 int64 = -428
                                            v42 = (v38 & (v42 | (v42 >> ((403) & 63))))
                                            _ = v50
                                        }
                                        if v2 {
                                            var v51 string = "hello"
                                            if v2 {
                                                v51 = "hello"
                                            }
                                            _ = v51
                                        } else {
                                            v52 := GenS3{F0: 1}
                                            var v53 int8 = 127
                                            var v54 float64 = 3.14159
                                            v55 := GenS3{F0: 42}
                                            var v56 int32 = 1
                                            switch (fi37 >> ((uint(v38)) & 63)) {
                                            case -313:
                                                if false {
                                                    if (18446744073709551615 <= 2) {
                                                        if ((v2 || (2147483647 < fi39)) && ((-(v53)) > (v53 << ((65535) & 63)))) {
                                                            if v2 {
                                                                var v57 uint32 = 65535
                                                                v58 := GenS4{F0: ("hello" + ("foo bar" + "hello")), F1: true}
                                                                switch ("a" + "hello") {
                                                                case "":
                                                                    v59 := GenS2{F0: (v57 * (^(v57))), F1: "\n\t", F2: v57}
                                                                    _ = v59
                                                                default:
                                                                    fmt.Println(v42)
                                                                }
                                                                _ = v57
                                                                _ = v58
                                                            }
                                                            v56 -= v56
                                                            if false {
                                                                _ = len(v1)
                                                                var v60 int8 = (-(v53))
                                                                v61 := GenS3{F0: 139}
                                                                _ = v60
                                                                _ = v61
                                                            } else {
                                                                v62 := GenS3{F0: 2}
                                                                v54 += -2.25
                                                                v38 = v38
                                                                _ = v62
                                                            }
                                                            fmt.Println(fi37)
                                                        }
                                                    }
                                                    v1 = append(v1, (fi39 ^ (fi37 << ((18446744073709551615) & 63))))
                                                    v63 := GenS6{F0: 1}
                                                    _ = v63
                                                } else {
                                                    v38 = (^(v42))
                                                }
                                            case 32768:
                                                var v64 int = fi39
                                                _ = v64
                                            default:
                                                if false {
                                                    if ((false || (518 != 4294967295)) || true) {
                                                        switch "\n\t" {
                                                        case "\n\t":
                                                            if (!v2) {
                                                                switch "foo bar" {
                                                                case "hello":
                                                                    var v65 int16 = 1
                                                                    _ = v65
                                                                case "a":
                                                                    _ = len(v3)
                                                                }
                                                                v2 = v2
                                                                var v66 int64 = 42
                                                                _ = v66
                                                            } else {
                                                                v2 = v2
                                                                var v67 uint16 = 1
                                                                fmt.Println(fi37, v2, v56)
                                                                fmt.Println(v42, fi37, v53)
                                                                if false {
                                                                    if true {
                                                                        var v68 float32 = (3.14159 / 1e10)
                                                                        v67 -= uint16(v38)
                                                                        if len(v1) >= 2 {
                                                                            v1 = v1[1:len(v1)]
                                                                        }
                                                                        _ = v68
                                                                    }
                                                                    v69 := GenS5{F0: 626, F1: 18446744073709551615}
                                                                    _ = v69
                                                                }
                                                                _ = v67
                                                            }
                                                        default:
                                                            if false {
                                                                _ = len(v1)
                                                                fmt.Println(v54)
                                                                v70 := GenS5{F0: 2, F1: 987}
                                                                _ = v70
                                                            }
                                                        }
                                                        var v71 string = "hello"
                                                        if true {
                                                            v72 := GenS1{F0: v56, F1: 0, F2: 0}
                                                            if len(v3) > 0 {
                                                                v3[0] = v71
                                                            }
                                                            v73 := GenS7{F0: 1, F1: (fi37 | -1), F2: 0}
                                                            _ = v72
                                                            _ = v73
                                                        }
                                                        _ = v71
                                                    }
                                                    v74 := GenS1{F0: v56, F1: 0, F2: 1}
                                                    _ = v74
                                                }
                                            }
                                            _ = v52
                                            _ = v53
                                            _ = v54
                                            _ = v55
                                            _ = v56
                                        }
                                        fmt.Println(fi39, v2)
                                        v42 += 42
                                        var v75 uint64 = 64
                                        v76 := GenS3{F0: 331}
                                        var v77 int = fi37
                                        _ = v75
                                        _ = v76
                                        _ = v77
                                    }
                                } else {
                                    if false {
                                        fmt.Println(v42)
                                        var v78 uint16 = 707
                                        switch ("\n\t" + "\n\t") {
                                        case "foo bar":
                                            var v79 float32 = 1e10
                                            _ = v79
                                        }
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                        if false {
                                            switch (fi37 - (fi37 | fi37)) {
                                            case -815:
                                                switch ("hello" + ("hello" + "foo bar")) {
                                                case "a":
                                                    if (!v2) {
                                                        var v80 uint16 = (v78 - (v78 - v78))
                                                        _ = v80
                                                    }
                                                case "\n\t":
                                                    fmt.Println(v2)
                                                case "foo bar":
                                                    switch 32767 {
                                                    case -9223372036854775808:
                                                        v81 := GenS5{F0: 405, F1: 13}
                                                        _ = v81
                                                    case 9223372036854775807:
                                                        var v82 float32 = 3.14159
                                                        var v83 float64 = float64(float32((3.14159 + -2.25)))
                                                        _ = v82
                                                        _ = v83
                                                    case -1:
                                                        fmt.Println(v78)
                                                    default:
                                                        var v84 uint8 = (v38 & (^(v38)))
                                                        _ = v84
                                                    }
                                                default:
                                                    _ = len(v1)
                                                }
                                            case 128:
                                                var v85 uint16 = v78
                                                _ = v85
                                            case 127:
                                                var v86 uint = 569
                                                _ = v86
                                            default:
                                                v87 := GenS8{F0: (fi39 % (((fi39 | (fi39 - fi39))) | 1)), F1: ("foo bar" + ""), F2: (42 != -113)}
                                                _ = v87
                                            }
                                        } else {
                                            var v88 uint32 = 433
                                            _ = v88
                                        }
                                        switch (-(fi37)) {
                                        case 2147483647:
                                            switch fi37 {
                                            case 466:
                                                fmt.Println(v78)
                                            case 32768:
                                                switch (fi39 >> ((191) & 63)) {
                                                case 13:
                                                    if len(v3) >= 2 {
                                                        v3 = v3[1:len(v3)]
                                                    }
                                                case 0:
                                                    switch (fi37 / ((-968) | 1)) {
                                                    case -9223372036854775808:
                                                        v89 := GenS8{F0: -2147483648, F1: ((("\n\t" + ("foo bar" + "hello")) + ("" + "foo bar")) + "a"), F2: ((float32((3.14159 / 1.0)) * 0.0) < (1e10 - (3.14159 + ((3.14159 / (3.14159 + 0.0)) - 1.5))))}
                                                        _ = v89
                                                    default:
                                                        if (!(v2 && (0 == 2))) {
                                                            v78 += v78
                                                            v38 = 128
                                                            v90 := GenS6{F0: 97}
                                                            _ = v90
                                                        } else {
                                                            var v91 uint64 = 0
                                                            _ = v91
                                                        }
                                                    }
                                                    fmt.Println(fi39)
                                                case 453:
                                                    v92 := GenS9{F0: -831}
                                                    _ = v92
                                                }
                                            case 128:
                                                v93 := GenS3{F0: 617}
                                                _ = v93
                                            }
                                        case 892:
                                            if true {
                                                switch 128 {
                                                case 9223372036854775807:
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                default:
                                                    switch "\n\t" {
                                                    case "":
                                                        var v94 int8 = 42
                                                        _ = v94
                                                    case "\n\t":
                                                        fmt.Println(v78, fi39, v2)
                                                    }
                                                }
                                                v38 -= v42
                                                _ = len(v1)
                                            } else {
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                                v2 = (18446744073709551615 < 2)
                                                v1 = append(v1, fi37)
                                            }
                                        default:
                                            v2 = v2
                                        }
                                        var v95 int64 = 869
                                        _ = v78
                                        _ = v95
                                    } else {
                                        v38 = v38
                                        switch fi39 {
                                        case -355:
                                            if (4294967295 > 350) {
                                                var v96 uint8 = v42
                                                var v97 int64 = 0
                                                _ = v96
                                                _ = v97
                                            }
                                        case 1:
                                            v98 := GenS9{F0: (^(fi37))}
                                            _ = v98
                                        }
                                    }
                                    if v2 {
                                        if v2 {
                                            fmt.Println(fi37, v38)
                                            if v2 {
                                                if len(v1) > 0 {
                                                    v1[0] = (fi37 & fi37)
                                                }
                                            }
                                        } else {
                                            var v99 string = (("a" + "") + "")
                                            _ = v99
                                        }
                                    }
                                    v2 = (0.0 >= 0.1)
                                }
                            }
                        case "foo bar":
                            v3 = append(v3, "a")
                        case "hello":
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        default:
                            switch "a" {
                            case "hello":
                                var v100 uint = 39
                                _ = v100
                            case "foo bar":
                                v38 = v38
                            }
                        }
                        _ = fi39
                        _ = v42
                        _ = v43
                    }
                    _ = v38
                }
                fmt.Println(fi37, v2)
                _ = fi37
            }
            _ = v36
        }
    }
    var v101 float32 = 1e10
    for fi102 := 0; fi102 < 6; fi102++ {
        v101 += v101
        fmt.Println(v2, v101)
        _ = fi102
    }
    for fi103 := 0; fi103 < 15; fi103++ {
        _ = len(v1)
        if len(v1) > 0 {
            v1[0] = (fi103 / (((fi103 / ((fi103) | 1))) | 1))
        }
        _ = fi103
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if (42 >= 128) {
        if (77 <= 791) {
            var v104 int16 = -181
            switch ("foo bar" + ((("hello" + "foo bar") + "a") + (("foo bar" + "foo bar") + "hello"))) {
            case "\n\t":
                if ((3.14159 - 3.14159) != ((1e10 + -2.25) * (1.5 * 1e10))) {
                    for fi105 := 0; fi105 < 7; fi105++ {
                        if false {
                            var v106 int16 = -715
                            _ = v106
                        } else {
                            switch 42 {
                            case 127:
                                v2 = true
                            case -946:
                                if v2 {
                                    for fi107 := 0; fi107 < 6; fi107++ {
                                        var v108 int16 = 0
                                        var v109 uint16 = 255
                                        _ = fi107
                                        _ = v108
                                        _ = v109
                                    }
                                }
                            }
                            _ = len(v3)
                            v101 = v101
                        }
                        _ = fi105
                    }
                    var v110 uint64 = 2
                    _ = v110
                }
            case "":
                if (v2 && false) {
                    v101 += v101
                    var v111 float32 = float32(0.1)
                    for fi112 := 0; fi112 < 8; fi112++ {
                        for fi113 := 0; fi113 < 16; fi113++ {
                            v114 := GenS7{F0: 85, F1: fi112, F2: 0}
                            _ = fi113
                            _ = v114
                        }
                        var v115 float64 = 1.5
                        if true {
                            v2 = (0 >= 2147483648)
                            var v116 int = (fi112 << ((0) & 63))
                            _ = v116
                        } else {
                            for fi117 := 0; fi117 < 14; fi117++ {
                                switch "" {
                                case "hello":
                                    if true {
                                        v118 := GenS4{F0: ("hello" + ""), F1: (198 != 239)}
                                        v101 = (3.14159 * (v111 / 1.0))
                                        switch (-(fi117)) {
                                        case 42:
                                            _ = len(v1)
                                        }
                                        if (("hello" + ("\n\t" + "foo bar")) < (("hello" + "hello") + "hello")) {
                                            if false {
                                                fmt.Println(v104)
                                                if len(v3) > 0 {
                                                    v3[0] = ""
                                                }
                                            }
                                        }
                                        _ = v118
                                    }
                                case "foo bar":
                                    v115 -= v115
                                    var v119 int32 = -56
                                    if len(v1) > 0 {
                                        v1[0] = fi117
                                    }
                                    _ = v119
                                case "\n\t":
                                    v104 = 1000
                                default:
                                    var v120 uint32 = 65535
                                    _ = v120
                                }
                                v121 := GenS4{F0: "\n\t", F1: true}
                                _ = fi117
                                _ = v121
                            }
                        }
                        _ = fi112
                        _ = v115
                    }
                    _ = v111
                }
            case "hello":
                v122 := GenS9{F0: 127}
                _ = v122
            }
            var v123 int32 = -2147483648
            _ = v104
            _ = v123
        } else {
            v101 -= ((v101 - v101) + v101)
        }
        var v124 uint16 = 42
        _ = v124
    }
    for fi125 := 0; fi125 < 6; fi125++ {
        var v126 int16 = -32768
        _ = fi125
        _ = v126
    }
    if (2 > 110) {
        if v2 {
            var v127 string = ""
            for fi128 := 0; fi128 < 19; fi128++ {
                if len(v1) > 0 {
                    v1[0] = (fi128 >> ((1) & 63))
                }
                v101 = (v101 * (((v101 / v101) * v101) / v101))
                fmt.Println(v101, v101)
                _ = fi128
            }
            v129 := GenS6{F0: 856}
            _ = v127
            _ = v129
        }
        switch ("hello" + "\n\t") {
        case "hello":
            var v130 int16 = 32767
            _ = v130
        case "":
            switch 32768 {
            case 42:
                v101 += (v101 + 0.0)
            case -9223372036854775808:
                v1 = append(v1, 4294967295)
            }
        default:
            v131 := GenS2{F0: 2, F1: "\n\t", F2: 1}
            _ = v131
        }
    } else {
        var v132 float32 = v101
        if v2 {
            v101 = v132
        } else {
            v133 := GenS10{F0: 0, F1: 2147483647, F2: 251}
            _ = v133
        }
        _ = v132
    }
    var v134 uint8 = 2
    if v2 {
        fmt.Println(v134, v134, v2)
        var v135 uint = 0
        _ = v135
    } else {
        var v136 int32 = 572
        _ = v136
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v137 := GenS11{F0: 4294967295}
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v101:", v101)
    fmt.Println("v134:", v134)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v101
    _ = v134
    _ = v137
    _ = v2
    _ = v3
}
