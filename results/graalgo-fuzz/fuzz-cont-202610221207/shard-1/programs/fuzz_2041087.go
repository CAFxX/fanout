package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int; F2 int64 }
type GenS2 struct { F0 string; F1 bool; F2 int16 }
type GenS3 struct { F0 uint32; F1 bool; F2 uint8 }
type GenS4 struct { F0 bool; F1 uint16 }
type GenS5 struct { F0 uint64; F1 uint64; F2 bool }
type GenS6 struct { F0 bool; F1 uint32 }
type GenS7 struct { F0 bool; F1 bool }
type GenS8 struct { F0 uint8 }
type GenS9 struct { F0 int16; F1 int16; F2 uint }
type GenS10 struct { F0 string; F1 uint }
type GenS11 struct { F0 int16; F1 uint64 }
type GenS12 struct { F0 uint64; F1 bool; F2 int8 }
type GenS13 struct { F0 int64 }
type GenS14 struct { F0 string; F1 int }

func main() {
    var v1 uint32 = 42
    v2 := []uint32{v1, 65535, 0}
    var v3 uint = 0
    var v4 int64 = -9223372036854775808
    v5 := make(map[string]uint32)
    v3 = (v3 - 42)
    fmt.Println(v4, v1, v1)
    switch -280 {
    case -213:
        var v6 int64 = v4
        _ = v6
    case 32767:
        v7 := GenS1{F0: (0 == -1), F1: -893, F2: v4}
        _ = v7
    default:
        for fi8 := 0; fi8 < 6; fi8++ {
            var v9 float64 = 0.1
            if (128 <= 128) {
                for fi10 := 0; fi10 < 10; fi10++ {
                    switch "foo bar" {
                    case "hello":
                        var v11 float32 = 3.14159
                        _ = v11
                    case "\n\t":
                        if ("\n\t" < "foo bar") {
                            var v12 float64 = v9
                            var v13 bool = true
                            _ = v12
                            _ = v13
                        }
                    case "a":
                        v3 = v3
                    }
                    if (uint16(v3) > 42) {
                        switch -1 {
                        case -2147483648:
                            v14 := GenS1{F0: (false || (fi10 == fi8)), F1: (fi8 - fi8), F2: (v4 - 0)}
                            _ = v14
                        case -196:
                            var v15 string = "a"
                            _ = v15
                        case 32768:
                            switch ((("foo bar" + "") + "a") + ("\n\t" + "\n\t")) {
                            case "a":
                                if (true && true) {
                                    fmt.Println(v4, v3, fi8)
                                    if (1 < uint64(v3)) {
                                        if (!(!false)) {
                                            v9 -= ((v9 * v9) + ((1.5 * v9) * (v9 / 1.0)))
                                            delete(v5, "a")
                                            switch fi10 {
                                            case -1:
                                                var v16 bool = true
                                                _ = v16
                                            case 706:
                                                v17 := GenS2{F0: (("foo bar" + "") + "a"), F1: false, F2: -1}
                                                _ = v17
                                            default:
                                                switch fi10 {
                                                case -9223372036854775808:
                                                    if (127 != 0) {
                                                        v18 := v5[("foo bar" + "")]
                                                        _ = v18
                                                    } else {
                                                        var v19 int16 = -32768
                                                        var v20 bool = false
                                                        var v21 int64 = v4
                                                        switch ((("a" + ("foo bar" + "hello")) + "\n\t") + "a") {
                                                        case "a":
                                                            _ = len(v2)
                                                        }
                                                        _ = v19
                                                        _ = v20
                                                        _ = v21
                                                    }
                                                case 903:
                                                    v22 := GenS1{F0: (!((v3 % ((65535) | 1)) <= (^(v3)))), F1: (fi10 - (^(fi10))), F2: v4}
                                                    _ = v22
                                                default:
                                                    v23 := GenS3{F0: (v1 >> ((v3) & 63)), F1: true, F2: 2}
                                                    _ = v23
                                                }
                                            }
                                        } else {
                                            fmt.Println(fi8, v1)
                                            switch 322 {
                                            case -128:
                                                v1 = 1
                                            case 127:
                                                var v24 string = ("" + "\n\t")
                                                _ = v24
                                            case -1:
                                                if (!(!false)) {
                                                    fmt.Println(fi8, v4, v4)
                                                    fmt.Println(fi10, fi10)
                                                    if true {
                                                        var v25 int = 2147483647
                                                        v25 += v25
                                                        _ = v25
                                                    } else {
                                                        if (128 != 42) {
                                                            var v26 uint32 = (^(v1))
                                                            _ = v26
                                                        }
                                                    }
                                                } else {
                                                    if false {
                                                        if (true || true) {
                                                            v27 := GenS4{F0: false, F1: 899}
                                                            var v28 uint8 = 128
                                                            _ = v27
                                                            _ = v28
                                                        }
                                                    }
                                                    switch (fi8 >> ((v3) & 63)) {
                                                    case -486:
                                                        v29 := GenS5{F0: 1, F1: 4294967295, F2: (true && (v3 >= v3))}
                                                        _ = v29
                                                    default:
                                                        fmt.Println(v1)
                                                    }
                                                    switch ("" + (("hello" + (("\n\t" + ("foo bar" + "a")) + "foo bar")) + "\n\t")) {
                                                    case "\n\t":
                                                        fmt.Println(fi8)
                                                    default:
                                                        v30 := GenS1{F0: false, F1: (^(fi10)), F2: v4}
                                                        _ = v30
                                                    }
                                                }
                                            }
                                            v31 := GenS1{F0: ((float64(-2.25) != ((v9 + (((((float64(0.1) * (v9 * float64(float32(v9)))) / 1e10) + v9) - v9) / v9)) / (-2.25 / 1.0))) && (true || (1 != (v4 | 0)))), F1: fi10, F2: v4}
                                            v9 -= v9
                                            _ = v31
                                        }
                                        var v32 float32 = (1.5 / 3.14159)
                                        _ = v32
                                    }
                                } else {
                                    fmt.Println(v1, v3)
                                    if (1000 != -394) {
                                        v1 = v1
                                    } else {
                                        var v33 int = 1
                                        var v34 int = (fi10 * (-(v33)))
                                        _ = v33
                                        _ = v34
                                    }
                                }
                            case "\n\t":
                                v35 := GenS3{F0: uint32(v3), F1: ((fi8 << ((91) & 63)) != fi10), F2: 25}
                                _ = v35
                            }
                        }
                        if false {
                            var v36 uint64 = 979
                            if (!(!true)) {
                                var v37 uint32 = v1
                                _ = v37
                            }
                            _ = v36
                        }
                        var v38 uint8 = 220
                        _ = v38
                    }
                    v39 := GenS6{F0: (fi10 < (fi8 + (^(fi8)))), F1: (v1 & (v1 & (v1 % (((v1 * v1)) | 1))))}
                    _ = fi10
                    _ = v39
                }
                var v40 float64 = float64((1.5 / 3.14159))
                v40 -= float64(0.0)
                _ = v40
            } else {
                if (!((!true) && (!true))) {
                    _ = len(v5)
                } else {
                    v41 := GenS7{F0: false, F1: false}
                    _ = v41
                }
                if len(v2) > 0 {
                    v2[0] = uint32(v3)
                }
                v42 := GenS6{F0: false, F1: v1}
                _ = v42
            }
            _ = fi8
            _ = v9
        }
    }
    if (42 >= 2) {
        v43 := GenS8{F0: 223}
        if true {
            switch ("foo bar" + ("hello" + "foo bar")) {
            case "a":
                for fi44 := 0; fi44 < 1; fi44++ {
                    switch ("\n\t" + ("foo bar" + "foo bar")) {
                    case "\n\t":
                        var v45 int64 = (-(v4))
                        _ = v45
                    case "a":
                        if true {
                            var v46 float32 = 1.5
                            v47 := GenS9{F0: -1, F1: 42, F2: v3}
                            _ = v46
                            _ = v47
                        }
                    }
                    _ = fi44
                }
            }
            var v48 string = "a"
            var v49 int8 = -1
            _ = v48
            _ = v49
        }
        if false {
            v3 = v3
            v3 = (v3 * v3)
        } else {
            if true {
                var v50 int8 = -36
                v3 -= (v3 + (v3 & (v3 & (v3 | (v3 << ((905) & 63))))))
                switch ("" + "hello") {
                case "foo bar":
                    for fi51 := 0; fi51 < 2; fi51++ {
                        for fi52 := 0; fi52 < 15; fi52++ {
                            var v53 int64 = (v4 | (v4 * v4))
                            _ = fi52
                            _ = v53
                        }
                        var v54 uint = v3
                        _ = fi51
                        _ = v54
                    }
                case "a":
                    fmt.Println(v1, v1, v50)
                case "\n\t":
                    var v55 bool = true
                    _ = v55
                }
                _ = v50
            } else {
                if ((v1 - uint32(v3)) > 2147483647) {
                    if (255 < 32767) {
                        if true {
                            v1 = (v1 ^ v1)
                        } else {
                            if false {
                                _ = len(v2)
                            }
                            if true {
                                var v56 int64 = (v4 - 42)
                                if ((v56 ^ v56) == (-(v4))) {
                                    var v57 int = 1
                                    var v58 int64 = -615
                                    _ = v57
                                    _ = v58
                                }
                                _ = v56
                            }
                            v59 := v5["hello"]
                            _ = v59
                        }
                        var v60 int8 = 107
                        var v61 int = 42
                        _ = v60
                        _ = v61
                    }
                    if true {
                        switch 42 {
                        case 626:
                            var v62 int16 = -1
                            _ = v62
                        case -109:
                            for fi63 := 0; fi63 < 9; fi63++ {
                                v1 = (v1 + (^(v1)))
                                fmt.Println(fi63, v4, fi63)
                                _ = fi63
                            }
                        }
                        v3 = (v3 ^ (v3 * 1))
                        var v64 bool = (2 < 18446744073709551615)
                        _ = v64
                    } else {
                        v65 := GenS5{F0: 42, F1: 16, F2: true}
                        v3 -= v3
                        _ = v65
                    }
                    _ = len(v5)
                } else {
                    for fi66 := 0; fi66 < 17; fi66++ {
                        var v67 int16 = 1000
                        fmt.Println(v4, v67, v67)
                        v4 = -9223372036854775808
                        _ = fi66
                        _ = v67
                    }
                    for fi68 := 0; fi68 < 2; fi68++ {
                        var v69 int16 = 655
                        v70 := GenS3{F0: 881, F1: ((fi68 * (fi68 - fi68)) >= fi68), F2: 42}
                        _ = fi68
                        _ = v69
                        _ = v70
                    }
                    var v71 float64 = ((0.1 * 1.5) / (1.5 - 3.14159))
                    _ = v71
                }
                switch "a" {
                case "":
                    if true {
                        if false {
                            for fi72 := 0; fi72 < 5; fi72++ {
                                v73 := GenS2{F0: "", F1: true, F2: 42}
                                var v74 int32 = 989
                                _ = fi72
                                _ = v73
                                _ = v74
                            }
                            if true {
                                if false {
                                    switch "hello" {
                                    case "foo bar":
                                        if (!(!(v4 <= (-(v4))))) {
                                            fmt.Println(v1, v1)
                                            fmt.Println(v4, v1)
                                            fmt.Println(v1, v3, v4)
                                        } else {
                                            v75 := GenS5{F0: 0, F1: 775, F2: (-2.25 >= (0.0 + float64(((0.0 + ((0.1 / 0.1) / 1.0)) - 0.1))))}
                                            v5["foo bar"] = v1
                                            _ = v75
                                        }
                                    case "\n\t":
                                        _ = len(v5)
                                    case "hello":
                                        v3 = 65535
                                    }
                                }
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                fmt.Println(v1, v3)
                            } else {
                                fmt.Println(v3)
                                v4 = v4
                                var v76 bool = false
                                _ = v76
                            }
                            if (2147483647 <= v1) {
                                v77 := GenS10{F0: ("foo bar" + "hello"), F1: v3}
                                if false {
                                    v78 := GenS10{F0: "foo bar", F1: (v3 << (((^(v3))) & 63))}
                                    var v79 uint = (v3 % ((800) | 1))
                                    for fi80 := 0; fi80 < 17; fi80++ {
                                        if len(v2) > 0 {
                                            v2[0] = 65535
                                        }
                                        v3 -= v3
                                        for fi81 := 0; fi81 < 4; fi81++ {
                                            _ = len(v5)
                                            switch "\n\t" {
                                            case "foo bar":
                                                v4 += v4
                                            case "hello":
                                                v1 = 1
                                                v1 = v1
                                                var v82 uint = (^(v79))
                                                _ = v82
                                            case "a":
                                                fmt.Println(v79, v1, v3)
                                            }
                                            v5["a"] = 1
                                            _ = fi81
                                        }
                                        _ = fi80
                                    }
                                    _ = v78
                                    _ = v79
                                } else {
                                    v2 = append(v2, (^(v1)))
                                    var v83 int64 = (-(v4))
                                    v84 := GenS9{F0: -577, F1: -1, F2: 1}
                                    _ = v83
                                    _ = v84
                                }
                                for fi85 := 0; fi85 < 11; fi85++ {
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    var v86 uint = 0
                                    _ = fi85
                                    _ = v86
                                }
                                _ = v77
                            } else {
                                if (int(v1) >= 9223372036854775807) {
                                    switch -1 {
                                    case 2147483647:
                                        var v87 string = "foo bar"
                                        _ = v87
                                    }
                                    var v88 int8 = 37
                                    fmt.Println(v4)
                                    _ = v88
                                }
                                for fi89 := 0; fi89 < 4; fi89++ {
                                    fmt.Println(v4)
                                    v90 := GenS10{F0: "hello", F1: (v3 >> ((v3) & 63))}
                                    _ = fi89
                                    _ = v90
                                }
                                for fi91 := 0; fi91 < 2; fi91++ {
                                    var v92 uint = (v3 + (v3 << (((^(v3))) & 63)))
                                    v92 += (v92 + v3)
                                    switch (fi91 >> ((65535) & 63)) {
                                    case 32767:
                                        v93 := v5[("foo bar" + "foo bar")]
                                        _ = v93
                                    case 419:
                                        _ = len(v2)
                                    case 0:
                                        var v94 bool = (((((1e10 / (float64(float32((3.14159 * 1e10))) / (3.14159 + (1e10 / -2.25)))) + (((1.5 + float64(0.0)) - 1e10) - ((0.0 * 0.0) - (-2.25 + -2.25)))) / (-2.25 / 1e10)) + (1.5 / 1.0)) < (1.5 * 0.0))
                                        _ = v94
                                    default:
                                        var v95 uint32 = (v1 + (v1 / ((2147483647) | 1)))
                                        _ = v95
                                    }
                                    _ = fi91
                                    _ = v92
                                }
                            }
                        }
                        var v96 uint = (v3 >> ((v3) & 63))
                        var v97 float64 = 0.0
                        _ = v96
                        _ = v97
                    }
                case "a":
                    fmt.Println(v3)
                case "\n\t":
                    if (!true) {
                        if len(v2) > 0 {
                            v2[0] = 329
                        }
                        var v98 uint8 = 128
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        _ = v98
                    } else {
                        v5["\n\t"] = (v1 & 65535)
                        for fi99 := 0; fi99 < 15; fi99++ {
                            var v100 uint64 = 436
                            var v101 float64 = (0.1 / 3.14159)
                            _ = fi99
                            _ = v100
                            _ = v101
                        }
                        if false {
                            switch "foo bar" {
                            case "foo bar":
                                v3 = 997
                            case "\n\t":
                                for fi102 := 0; fi102 < 18; fi102++ {
                                    if true {
                                        if ((fi102 & fi102) >= (fi102 - (fi102 | 32767))) {
                                            var v103 int = (fi102 >> ((179) & 63))
                                            for fi104 := 0; fi104 < 8; fi104++ {
                                                v105 := GenS2{F0: "a", F1: false, F2: 1000}
                                                fmt.Println(v103, v3, fi104)
                                                _ = fi104
                                                _ = v105
                                            }
                                            _ = v103
                                        } else {
                                            for fi106 := 0; fi106 < 14; fi106++ {
                                                var v107 uint16 = 1
                                                _ = fi106
                                                _ = v107
                                            }
                                        }
                                        for fi108 := 0; fi108 < 19; fi108++ {
                                            v4 += v4
                                            _ = fi108
                                        }
                                        v109 := GenS11{F0: 1000, F1: 844}
                                        _ = v109
                                    } else {
                                        v2 = append(v2, 0)
                                        v2 = append(v2, v1)
                                    }
                                    _ = fi102
                                }
                            case "a":
                                v3 = v3
                            default:
                                if true {
                                    v1 = uint32(v3)
                                }
                            }
                        }
                    }
                default:
                    var v110 string = ("hello" + (("foo bar" + "foo bar") + ("hello" + (("" + ("hello" + "hello")) + "\n\t"))))
                    _ = v110
                }
                v4 -= v4
            }
        }
        _ = v43
    }
    v3 = v3
    var v111 uint16 = 684
    v112 := GenS8{F0: uint8(v3)}
    for fi113 := 0; fi113 < 13; fi113++ {
        var v114 int64 = v4
        var v115 int16 = 1
        fmt.Println(v1, v3, v3)
        _ = fi113
        _ = v114
        _ = v115
    }
    switch "hello" {
    case "\n\t":
        var v116 uint32 = (^(v1))
        _ = v116
    default:
        var v117 float32 = -2.25
        _ = v117
    }
    var v118 int16 = -32768
    if false {
        var v119 bool = (128 == 4294967295)
        var v120 bool = (!v119)
        v121 := GenS12{F0: 0, F1: false, F2: -60}
        _ = v119
        _ = v120
        _ = v121
    }
    var v122 int8 = 1
    fmt.Println(v111, v122)
    _ = len(v2)
    if true {
        switch 127 {
        case 128:
            v118 -= (v118 & v118)
        case -451:
            var v123 int8 = -128
            _ = v123
        case 2147483647:
            var v124 int8 = 78
            _ = v124
        }
        fmt.Println(v1, v4)
        fmt.Println(v118)
    } else {
        for fi125 := 0; fi125 < 19; fi125++ {
            v4 -= int64(v111)
            if true {
                if (false || false) {
                    var v126 string = "hello"
                    for fi127 := 0; fi127 < 17; fi127++ {
                        if true {
                            _ = len(v5)
                            switch ("foo bar" + v126) {
                            case "hello":
                                var v128 uint32 = uint32(v111)
                                _ = v128
                            case "foo bar":
                                v5[v126] = 685
                            case "\n\t":
                                fmt.Println(v122, v1)
                            default:
                                v126 = ("\n\t" + (v126 + v126))
                            }
                            fmt.Println(v126, v126, v122)
                        }
                        var v129 bool = true
                        _ = len(v5)
                        switch fi125 {
                        case 548:
                            v130 := GenS2{F0: ("hello" + "hello"), F1: v129, F2: -1}
                            _ = v130
                        case 32767:
                            if (int16(v1) <= v118) {
                                v131 := v5[((("\n\t" + "\n\t") + "hello") + ("a" + v126))]
                                switch v126 {
                                case "":
                                    fmt.Println(fi127, v129)
                                case "\n\t":
                                    if v129 {
                                        var v132 int64 = 1
                                        switch (fi125 ^ (^(fi127))) {
                                        case 1:
                                            if (v4 <= (v132 << ((v3) & 63))) {
                                                v133 := GenS8{F0: 128}
                                                var v134 float64 = (float64(1.5) / -2.25)
                                                _ = v133
                                                _ = v134
                                            }
                                        case 9223372036854775807:
                                            switch v126 {
                                            case "foo bar":
                                                var v135 int32 = 1
                                                _ = v135
                                            default:
                                                if v129 {
                                                    v122 = 100
                                                    if false {
                                                        var v136 int32 = 0
                                                        _ = v136
                                                    } else {
                                                        fmt.Println(v4)
                                                        fmt.Println(v4)
                                                    }
                                                }
                                            }
                                        default:
                                            if (2 > 4294967295) {
                                                v137 := GenS13{F0: (v132 | v132)}
                                                if len(v2) > 0 {
                                                    v2[0] = v131
                                                }
                                                _ = v137
                                            } else {
                                                v138 := GenS2{F0: (("" + (v126 + "\n\t")) + "foo bar"), F1: v129, F2: (v118 & (v118 << ((255) & 63)))}
                                                switch "foo bar" {
                                                case "":
                                                    v139 := v5[v126]
                                                    _ = v139
                                                }
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                                if (!((v111 ^ v111) >= uint16(v3))) {
                                                    v131 -= 65535
                                                }
                                                v122 = v122
                                                _ = v138
                                            }
                                        }
                                        _ = v132
                                    }
                                case "foo bar":
                                    v111 -= 2
                                }
                                v122 += v122
                                _ = v131
                            }
                        default:
                            if v129 {
                                var v140 string = ("\n\t" + v126)
                                var v141 int16 = int16(v4)
                                _ = v140
                                _ = v141
                            } else {
                                v142 := GenS5{F0: uint64(v3), F1: 0, F2: false}
                                if v129 {
                                    v143 := GenS5{F0: 175, F1: 184, F2: (v129 || (!(!v129)))}
                                    switch int(v111) {
                                    case -9223372036854775808:
                                        var v144 string = ""
                                        _ = v144
                                    case 32767:
                                        v1 = (v1 >> (((^(v3))) & 63))
                                    case 127:
                                        if (v129 && v129) {
                                            v4 -= v4
                                            if false {
                                                var v145 uint16 = v111
                                                v146 := GenS1{F0: ((v122 + (-(v122))) > (v122 / ((72) | 1))), F1: int(v1), F2: (v4 - (v4 & int64(v122)))}
                                                if ((v1 % ((v1) | 1)) >= v1) {
                                                    if len(v2) > 0 {
                                                        v2[0] = 65535
                                                    }
                                                    switch 32768 {
                                                    case 1:
                                                        v4 = v4
                                                    case -715:
                                                        if len(v2) > 0 {
                                                            v2[0] = 0
                                                        }
                                                    case 832:
                                                        fmt.Println(v145, v4)
                                                    default:
                                                        var v147 int32 = -1
                                                        _ = v147
                                                    }
                                                    v5["hello"] = (v1 ^ v1)
                                                } else {
                                                    if false {
                                                        v118 += 1
                                                        v3 = (v3 | v3)
                                                        if v129 {
                                                            var v148 float32 = 3.14159
                                                            var v149 int16 = int16(v122)
                                                            v3 -= v3
                                                            _ = v148
                                                            _ = v149
                                                        } else {
                                                            var v150 uint = 214
                                                            _ = v150
                                                        }
                                                    }
                                                }
                                                _ = v145
                                                _ = v146
                                            } else {
                                                v1 -= 2147483648
                                                if true {
                                                    var v151 int16 = (v118 - (v118 % (((v118 / (((v118 ^ (v118 >> (((^(v3))) & 63)))) | 1))) | 1)))
                                                    _ = v151
                                                } else {
                                                    var v152 uint = v3
                                                    v126 = v126
                                                    v2 = append(v2, (v1 + v1))
                                                    _ = v152
                                                }
                                            }
                                            v153 := GenS11{F0: v118, F1: 1}
                                            v154 := GenS2{F0: "foo bar", F1: v129, F2: v118}
                                            _ = v153
                                            _ = v154
                                        }
                                    }
                                    _ = len(v2)
                                    _ = v143
                                } else {
                                    switch 0 {
                                    case -9223372036854775808:
                                        var v155 float32 = 3.14159
                                        _ = v155
                                    case 127:
                                        switch ("hello" + (("a" + "foo bar") + (v126 + ((("hello" + (("a" + "foo bar") + "")) + v126) + v126)))) {
                                        case "a":
                                            v4 = v4
                                        case "hello":
                                            if false {
                                                if len(v2) > 0 {
                                                    v2[0] = v1
                                                }
                                            } else {
                                                var v156 int = fi125
                                                switch 1 {
                                                case 0:
                                                    switch v126 {
                                                    case "hello":
                                                        if false {
                                                            var v157 uint = 272
                                                            _ = len(v2)
                                                            _ = v157
                                                        } else {
                                                            delete(v5, v126)
                                                            delete(v5, ((((v126 + "") + (v126 + ("a" + v126))) + "a") + ""))
                                                            var v158 int16 = (v118 * 1)
                                                            var v159 int16 = (-(v118))
                                                            v158 += (-(v158))
                                                            _ = v158
                                                            _ = v159
                                                        }
                                                    }
                                                    if v129 {
                                                        v160 := v5[(v126 + ("hello" + "a"))]
                                                        _ = v160
                                                    }
                                                    fmt.Println(v3)
                                                default:
                                                    switch fi127 {
                                                    case -1:
                                                        v161 := GenS10{F0: v126, F1: (v3 | (v3 / ((793) | 1)))}
                                                        _ = v161
                                                    default:
                                                        var v162 string = v126
                                                        _ = v162
                                                    }
                                                }
                                                _ = len(v5)
                                                _ = v156
                                            }
                                        case "\n\t":
                                            var v163 int = (fi127 * fi125)
                                            _ = v163
                                        }
                                    case -754:
                                        v111 = 442
                                    default:
                                        fmt.Println(fi125)
                                    }
                                }
                                _ = v142
                            }
                        }
                        _ = fi127
                        _ = v129
                    }
                    _ = v126
                }
            }
            _ = fi125
        }
        v118 = v118
    }
    for fi164 := 0; fi164 < 9; fi164++ {
        if false {
            if (true && false) {
                if (!true) {
                    if len(v2) > 0 {
                        v2[0] = v1
                    }
                    var v165 bool = false
                    _ = v165
                }
            } else {
                v166 := GenS14{F0: ("foo bar" + "foo bar"), F1: (-(fi164))}
                var v167 int8 = (v122 % (((^(v122))) | 1))
                _ = v166
                _ = v167
            }
            v118 = (v118 ^ 1000)
            if (255 == (v3 << ((v3) & 63))) {
                fmt.Println(fi164, fi164)
            } else {
                for fi168 := 0; fi168 < 9; fi168++ {
                    v169 := GenS6{F0: ((^(v111)) == 255), F1: 497}
                    var v170 bool = (((3.14159 - 3.14159) >= 1.5) && ((false || true) && (true || true)))
                    var v171 string = ""
                    _ = fi168
                    _ = v169
                    _ = v170
                    _ = v171
                }
                fmt.Println(v4, v1, v4)
                var v172 int64 = (-(v4))
                _ = v172
            }
        } else {
            fmt.Println(fi164, v4, fi164)
            if true {
                if (false && false) {
                    var v173 uint32 = (v1 << ((v3) & 63))
                    _ = v173
                } else {
                    v1 = (v1 % ((91) | 1))
                }
                if true {
                    v1 += (v1 / (((v1 << (((v3 ^ (v3 / ((v3) | 1)))) & 63))) | 1))
                }
            } else {
                if true {
                    v1 += (v1 & 589)
                    var v174 int32 = -985
                    v2 = append(v2, (v1 * v1))
                    _ = v174
                }
                v1 += 251
                var v175 bool = (0 < v4)
                _ = v175
            }
            _ = len(v5)
        }
        switch 1 {
        case 32767:
            fmt.Println(v4, fi164, v1)
        case 42:
            var v176 int8 = 119
            _ = v176
        case 4294967295:
            var v177 uint16 = v111
            _ = v177
        default:
            v2 = append(v2, (v1 ^ 1))
        }
        for fi178 := 0; fi178 < 20; fi178++ {
            var v179 uint = v3
            var v180 int64 = (v4 << (((v3 + v179)) & 63))
            _ = fi178
            _ = v179
            _ = v180
        }
        _ = fi164
    }
    v181 := GenS13{F0: -9223372036854775808}
    fmt.Println(v111, v3, v111)
    fmt.Println("v1:", v1)
    fmt.Println("v111:", v111)
    fmt.Println("v118:", v118)
    fmt.Println("v122:", v122)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v111
    _ = v112
    _ = v118
    _ = v122
    _ = v181
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
