package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int8; F1 int }
type GenS4 struct { F0 uint; F1 string }
type GenS5 struct { F0 uint; F1 int8; F2 uint }
type GenS6 struct { F0 int32; F1 uint }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 uint32 }
type GenS9 struct { F0 int; F1 bool }
type GenS10 struct { F0 int; F1 int8 }
type GenS11 struct { F0 bool; F1 int32 }
type GenS12 struct { F0 uint16; F1 string; F2 uint8 }
type GenS13 struct { F0 uint32; F1 bool }
type GenS14 struct { F0 int }
type GenS15 struct { F0 string; F1 int16 }
type GenS16 struct { F0 uint64; F1 uint }
type GenS17 struct { F0 uint16; F1 uint16; F2 int8 }

func main() {
    var v1 bool = false
    var v2 int8 = 103
    var v3 int = -122
    var v4 uint16 = 1
    v5 := []int{(v3 * 0), (^(v3))}
    v6 := make(map[int]bool)
    switch v3 {
    case 2147483647:
        var v7 int64 = 9223372036854775807
        _ = v7
    default:
        var v8 int8 = 90
        _ = v8
    }
    var v9 int8 = (v2 ^ v2)
    v9 = v2
    switch -1 {
    case 9223372036854775807:
        v6[(v3 - v3)] = v1
    case 128:
        if true {
            var v10 string = ("foo bar" + "\n\t")
            _ = len(v5)
            _ = v10
        }
    case 4294967295:
        for fi11 := 0; fi11 < 19; fi11++ {
            v1 = (v1 || v1)
            switch (-(fi11)) {
            case -1:
                var v12 bool = v1
                _ = v12
            case -610:
                var v13 int16 = 32767
                _ = v13
            }
            if false {
                delete(v6, (fi11 | v3))
            } else {
                switch 4294967295 {
                case -1:
                    if v1 {
                        if (110 >= 65535) {
                            var v14 int32 = 2147483647
                            var v15 float64 = -2.25
                            _ = v14
                            _ = v15
                        } else {
                            v16 := GenS1{F0: 0}
                            fmt.Println(v9, v3)
                            switch "foo bar" {
                            case "hello":
                                v9 += (v9 * (v2 ^ v9))
                            }
                            _ = v16
                        }
                    } else {
                        v9 = v2
                        for fi17 := 0; fi17 < 13; fi17++ {
                            v3 = 9223372036854775807
                            if (3.14159 != 3.14159) {
                                fmt.Println(v1, fi11, v2)
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                            }
                            _ = fi17
                        }
                        for fi18 := 0; fi18 < 2; fi18++ {
                            v1 = v1
                            _ = fi18
                        }
                    }
                case 42:
                    for fi19 := 0; fi19 < 2; fi19++ {
                        v20 := GenS2{F0: 0}
                        v21 := GenS3{F0: 42, F1: v3}
                        if v1 {
                            v1 = true
                            v4 += v4
                        } else {
                            switch ("\n\t" + "foo bar") {
                            case "\n\t":
                                var v22 uint32 = 42
                                _ = v22
                            }
                            var v23 string = (("" + "") + ("foo bar" + "a"))
                            _ = v23
                        }
                        _ = fi19
                        _ = v20
                        _ = v21
                    }
                }
                var v24 int32 = -110
                _ = len(v6)
                _ = v24
            }
            _ = fi11
        }
    default:
        var v25 bool = v1
        _ = v25
    }
    fmt.Println(v2, v1, v3)
    var v26 uint32 = 611
    for fi27 := 0; fi27 < 16; fi27++ {
        var v28 bool = (!false)
        v29 := GenS4{F0: 18446744073709551615, F1: ("\n\t" + ("\n\t" + "foo bar"))}
        _ = fi27
        _ = v28
        _ = v29
    }
    if v1 {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        v30 := GenS3{F0: v9, F1: 680}
        var v31 uint8 = 6
        _ = v30
        _ = v31
    }
    switch "\n\t" {
    case "a":
        switch "hello" {
        case "a":
            if v1 {
                for fi32 := 0; fi32 < 7; fi32++ {
                    v33 := GenS4{F0: 603, F1: ""}
                    if v1 {
                        v4 += (v4 % (((v4 - v4)) | 1))
                    }
                    _ = fi32
                    _ = v33
                }
                fmt.Println(v1)
                for fi34 := 0; fi34 < 11; fi34++ {
                    var v35 float64 = (1.5 + 3.14159)
                    v36 := GenS1{F0: v26}
                    _ = len(v5)
                    _ = fi34
                    _ = v35
                    _ = v36
                }
            }
        case "foo bar":
            switch (v3 - v3) {
            case -235:
                v5 = append(v5, (v3 | 128))
            case -2:
                _ = len(v6)
            default:
                var v37 bool = v1
                _ = v37
            }
        default:
            if ((0.0 >= 1.5) || false) {
                _ = len(v6)
                switch (v3 << ((0) & 63)) {
                case 0:
                    switch v3 {
                    case 32767:
                        v38 := v6[2147483647]
                        _ = v38
                    case -628:
                        var v39 uint8 = 244
                        _ = v39
                    default:
                        if (!(42 > 42)) {
                            if (("a" + "") != "") {
                                fmt.Println(v1, v26, v4)
                                for fi40 := 0; fi40 < 1; fi40++ {
                                    v3 += -398
                                    var v41 uint32 = (v26 / (((v26 << ((4294967295) & 63))) | 1))
                                    v42 := v6[9223372036854775807]
                                    _ = fi40
                                    _ = v41
                                    _ = v42
                                }
                            } else {
                                var v43 string = "foo bar"
                                var v44 float32 = 0.1
                                _ = v43
                                _ = v44
                            }
                            fmt.Println(v26, v9)
                        } else {
                            fmt.Println(v1)
                        }
                    }
                default:
                    if (!(!(2147483647 > 1))) {
                        var v45 uint64 = 1
                        _ = v45
                    } else {
                        var v46 int32 = 827
                        fmt.Println(v3, v46)
                        var v47 float64 = float64(1e10)
                        _ = v46
                        _ = v47
                    }
                }
            }
        }
    case "foo bar":
        switch "" {
        case "\n\t":
            var v48 uint32 = (v26 & v26)
            _ = v48
        case "a":
            if (!(v1 && false)) {
                v49 := GenS5{F0: uint(v26), F1: v2, F2: 947}
                _ = v49
            }
        case "foo bar":
            var v50 int16 = 1000
            _ = v50
        default:
            for fi51 := 0; fi51 < 3; fi51++ {
                var v52 float32 = 0.0
                var v53 uint64 = uint64(v26)
                v54 := v6[(fi51 - 0)]
                _ = fi51
                _ = v52
                _ = v53
                _ = v54
            }
        }
    }
    if v1 {
        if (((0.1 + float32(1.5)) > 3.14159) || (v1 && ((float32(1e10) / 1.0) >= (3.14159 / ((((-2.25 / 0.1) - 1e10) * (float32(0.1) * (1.5 * (1e10 * (-2.25 / 0.1))))) / (-2.25 * (float32(float64((1e10 * -2.25))) * 0.1))))))) {
            v55 := GenS6{F0: 189, F1: 0}
            _ = v55
        } else {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            var v56 uint16 = (^(v4))
            fmt.Println(v3)
            _ = v56
        }
        if (!(-1 <= -1)) {
            switch (^(v3)) {
            case 4294967295:
                var v57 bool = false
                _ = v57
            case 2147483647:
                fmt.Println(v1)
            case -1:
                _ = len(v5)
            default:
                v1 = ((0 != 4294967295) || true)
            }
            v58 := GenS4{F0: 18446744073709551615, F1: "hello"}
            _ = v58
        }
        v4 += v4
    }
    if v1 {
        switch "a" {
        case "\n\t":
            if (!(!(v1 || v1))) {
                if (65535 < v4) {
                    for fi59 := 0; fi59 < 18; fi59++ {
                        v2 += (v2 * 42)
                        v60 := GenS4{F0: 2, F1: (((((("" + (("hello" + ("hello" + "foo bar")) + "\n\t")) + "a") + "\n\t") + "hello") + (("foo bar" + "") + "a")) + "foo bar")}
                        var v61 float64 = (float64(1e10) + (float64(0.0) * 0.1))
                        _ = fi59
                        _ = v60
                        _ = v61
                    }
                    var v62 int32 = 2147483647
                    _ = v62
                }
                var v63 int16 = 1
                v64 := GenS7{F0: -1}
                _ = v63
                _ = v64
            } else {
                for fi65 := 0; fi65 < 10; fi65++ {
                    if v1 {
                        delete(v6, (v3 - (v3 % (((fi65 + (fi65 % (((v3 << ((0) & 63))) | 1)))) | 1))))
                        for fi66 := 0; fi66 < 9; fi66++ {
                            var v67 uint32 = v26
                            fmt.Println(fi66, fi66)
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                            v3 -= v3
                            var v68 uint64 = 1
                            _ = fi66
                            _ = v67
                            _ = v68
                        }
                    } else {
                        switch "a" {
                        case "":
                            for fi69 := 0; fi69 < 13; fi69++ {
                                fmt.Println(v2, fi69, v26)
                                _ = fi69
                            }
                        case "foo bar":
                            if (true && (!((false && (true && ((v4 % (((^(v4))) | 1)) >= (v4 * v4)))) && false))) {
                                fmt.Println(fi65, v26, v9)
                            }
                        case "\n\t":
                            v3 = v3
                        default:
                            v2 += v9
                        }
                    }
                    _ = fi65
                }
                if (255 >= 60) {
                    if true {
                        v4 = (v4 * (v4 ^ v4))
                        switch v3 {
                        case 629:
                            _ = len(v6)
                        default:
                            if ("a" == "hello") {
                                v2 -= (v9 ^ (v9 / (((v9 * 0)) | 1)))
                                v70 := GenS8{F0: v26}
                                _ = v70
                            }
                        }
                        var v71 int = (v3 << ((18446744073709551615) & 63))
                        _ = v71
                    }
                    v72 := v6[-344]
                    _ = v72
                }
            }
        case "foo bar":
            switch int(v26) {
            case 811:
                v73 := GenS8{F0: (v26 | (v26 / ((uint32(v4)) | 1)))}
                _ = v73
            case 389:
                var v74 int = (v3 - (v3 >> ((4294967295) & 63)))
                _ = v74
            default:
                v2 += v2
            }
        case "a":
            switch "" {
            case "":
                if v1 {
                    var v75 uint32 = 24
                    _ = v75
                } else {
                    v3 = v3
                    for fi76 := 0; fi76 < 16; fi76++ {
                        if true {
                            for fi77 := 0; fi77 < 16; fi77++ {
                                v2 += 127
                                _ = fi77
                            }
                            if (((v9 <= v9) || ((-619 > int16(v3)) || v1)) || (!("foo bar" <= "foo bar"))) {
                                v2 -= (v9 << ((479) & 63))
                                var v78 int16 = 0
                                if (v1 || (float32(((((0.0 / 1e10) * (3.14159 / ((((-2.25 / 1.5) / -2.25) + 0.0) / ((-2.25 + (float64(((3.14159 / 0.1) / (-2.25 - 3.14159))) - ((1.5 * 0.0) - 3.14159))) - 1e10)))) / 1.5) * (0.1 / 1.5))) <= ((-2.25 * 3.14159) - 3.14159))) {
                                    var v79 int8 = (v9 >> ((763) & 63))
                                    for fi80 := 0; fi80 < 4; fi80++ {
                                        if v1 {
                                            v81 := GenS4{F0: 0, F1: ""}
                                            var v82 float32 = 0.1
                                            _ = v81
                                            _ = v82
                                        } else {
                                            if false {
                                                v79 = 127
                                                v83 := GenS9{F0: int(v78), F1: (4294967295 <= 4294967295)}
                                                v9 -= (v2 & v79)
                                                var v84 float32 = (-2.25 + (1.5 - -2.25))
                                                _ = v83
                                                _ = v84
                                            }
                                            if len(v5) >= 2 {
                                                v5 = v5[1:len(v5)]
                                            }
                                        }
                                        var v85 int = v3
                                        var v86 int16 = (-(v78))
                                        _ = fi80
                                        _ = v85
                                        _ = v86
                                    }
                                    _ = v79
                                }
                                _ = v78
                            }
                        }
                        _ = fi76
                    }
                    switch "a" {
                    case "hello":
                        switch "hello" {
                        case "\n\t":
                            for fi87 := 0; fi87 < 3; fi87++ {
                                for fi88 := 0; fi88 < 9; fi88++ {
                                    switch "a" {
                                    case "foo bar":
                                        if false {
                                            v89 := GenS10{F0: (fi88 + (fi87 << ((0) & 63))), F1: (-(v9))}
                                            if (-2147483648 >= 2147483647) {
                                                v4 -= v4
                                                fmt.Println(v3, v3)
                                                if (false || v1) {
                                                    var v90 uint64 = 18446744073709551615
                                                    if v1 {
                                                        if (!(!(!v1))) {
                                                            if ((!((v1 || (false || (false || v1))) && (1e10 != 1e10))) && false) {
                                                                var v91 uint64 = uint64(v26)
                                                                var v92 uint8 = 1
                                                                v93 := v6[fi87]
                                                                _ = v91
                                                                _ = v92
                                                                _ = v93
                                                            }
                                                        }
                                                        var v94 int8 = v9
                                                        fmt.Println(v2, v2)
                                                        _ = v94
                                                    }
                                                    _ = v90
                                                } else {
                                                    v95 := GenS11{F0: v1, F1: 766}
                                                    fmt.Println(fi87, v3, v1)
                                                    switch 492 {
                                                    case -9223372036854775808:
                                                        fmt.Println(v26, v4, v4)
                                                        v26 = (v26 | v26)
                                                        v26 = 2
                                                    case 32768:
                                                        if true {
                                                            v1 = (1e10 == ((1e10 * (0.1 * (((3.14159 / (1.5 - 1e10)) / 1e10) / 0.1))) + 3.14159))
                                                            v6[fi87] = (((0.1 * -2.25) >= (float32((float64((1.5 * 1.5)) + ((-2.25 * (float64(-2.25) - -2.25)) - float64(3.14159)))) - float32(0.0))) && v1)
                                                        } else {
                                                            fmt.Println(v4)
                                                        }
                                                    default:
                                                        var v96 uint = 0
                                                        _ = v96
                                                    }
                                                    _ = v95
                                                }
                                            } else {
                                                var v97 uint = 255
                                                switch (fi88 & fi87) {
                                                case -157:
                                                    v98 := GenS12{F0: uint16(v97), F1: "hello", F2: 0}
                                                    _ = v98
                                                case 305:
                                                    switch (fi87 % ((fi87) | 1)) {
                                                    case -2147483648:
                                                        _ = len(v6)
                                                    case 127:
                                                        var v99 uint8 = 0
                                                        _ = v99
                                                    case 32767:
                                                        if v1 {
                                                            var v100 uint32 = (^(v26))
                                                            var v101 string = "foo bar"
                                                            fmt.Println(v9)
                                                            v102 := GenS11{F0: false, F1: -744}
                                                            _ = v100
                                                            _ = v101
                                                            _ = v102
                                                        }
                                                    default:
                                                        v103 := v6[fi87]
                                                        _ = v103
                                                    }
                                                case 883:
                                                    var v104 bool = (1 == v97)
                                                    _ = v104
                                                default:
                                                    if (!(0.1 >= 0.1)) {
                                                        v97 = (^(v97))
                                                        var v105 int8 = (v9 >> ((v97) & 63))
                                                        v5 = append(v5, (fi88 % (((fi87 ^ (fi88 << ((v97) & 63)))) | 1)))
                                                        var v106 uint = (^(v97))
                                                        var v107 int16 = -1
                                                        _ = v105
                                                        _ = v106
                                                        _ = v107
                                                    }
                                                }
                                                var v108 uint16 = 0
                                                _ = v108
                                                _ = v97
                                            }
                                            _ = v89
                                        } else {
                                            v109 := GenS2{F0: -1}
                                            if false {
                                                v1 = (!true)
                                                if false {
                                                    v5 = append(v5, fi87)
                                                    if (v4 <= (^(v4))) {
                                                        if len(v5) > 0 {
                                                            v5[0] = (v3 / ((fi87) | 1))
                                                        }
                                                    }
                                                } else {
                                                    v4 += v4
                                                    _ = len(v5)
                                                }
                                                v26 = 2
                                                v110 := v6[(-(fi87))]
                                                if v110 {
                                                    if (int8(v4) <= 42) {
                                                        v111 := GenS2{F0: 443}
                                                        fmt.Println(v2, fi88, fi88)
                                                        if v1 {
                                                            var v112 uint64 = 4294967295
                                                            fmt.Println(fi87, fi87)
                                                            _ = v112
                                                        }
                                                        switch fi88 {
                                                        case 4294967295:
                                                            var v113 int = (fi87 % ((0) | 1))
                                                            _ = v113
                                                        case 769:
                                                            var v114 int64 = -588
                                                            _ = v114
                                                        }
                                                        _ = len(v6)
                                                        var v115 uint32 = (v26 - v26)
                                                        _ = v111
                                                        _ = v115
                                                    }
                                                } else {
                                                    if (!((v2 - 100) >= int8(v26))) {
                                                        if (94 == uint(v26)) {
                                                            if (917 == 42) {
                                                                var v116 int64 = 0
                                                                v1 = ((v116 / ((-9223372036854775808) | 1)) >= -1)
                                                                var v117 int8 = -24
                                                                var v118 string = ""
                                                                _ = v116
                                                                _ = v117
                                                                _ = v118
                                                            }
                                                        } else {
                                                            var v119 uint32 = (v26 & (^(v26)))
                                                            _ = v119
                                                        }
                                                        fmt.Println(v3, v110)
                                                        fmt.Println(v26, fi87, v9)
                                                    }
                                                }
                                                _ = v110
                                            }
                                            _ = v109
                                        }
                                    case "":
                                        if true {
                                            fmt.Println(v1)
                                            var v120 int = (^(fi87))
                                            v1 = (-30 == 0)
                                            var v121 int = fi88
                                            _ = v120
                                            _ = v121
                                        }
                                    case "\n\t":
                                        v122 := GenS12{F0: (v4 * 0), F1: "", F2: 128}
                                        _ = v122
                                    default:
                                        switch (("hello" + "\n\t") + "a") {
                                        case "a":
                                            var v123 float32 = (1.5 / ((3.14159 / -2.25) - (0.1 / 0.1)))
                                            _ = v123
                                        case "foo bar":
                                            v124 := GenS2{F0: -863}
                                            _ = v124
                                        }
                                    }
                                    fmt.Println(v4, fi87, v26)
                                    var v125 uint = uint(v4)
                                    if ((v2 / (((^(v2))) | 1)) != (^(v9))) {
                                        switch (fi87 * fi88) {
                                        case 32767:
                                            var v126 int32 = 678
                                            _ = v126
                                        case -2147483648:
                                            v127 := GenS13{F0: v26, F1: v1}
                                            _ = v127
                                        case 244:
                                            if true {
                                                fmt.Println(v3)
                                                v128 := GenS13{F0: v26, F1: (v3 != (fi87 * -1))}
                                                if (1.5 == 0.1) {
                                                    switch (fi88 >> (((^(v125))) & 63)) {
                                                    case 42:
                                                        switch ("\n\t" + "a") {
                                                        case "a":
                                                            v1 = (v9 > 103)
                                                        case "\n\t":
                                                            if v1 {
                                                                v129 := GenS7{F0: -130}
                                                                _ = len(v5)
                                                                _ = v129
                                                            } else {
                                                                if v1 {
                                                                    if v1 {
                                                                        delete(v6, 819)
                                                                        switch int(v125) {
                                                                        case 198:
                                                                            var v130 uint64 = 0
                                                                            _ = v130
                                                                        case 946:
                                                                            var v131 int16 = 1
                                                                            _ = v131
                                                                        case 128:
                                                                            var v132 uint32 = 870
                                                                            _ = v132
                                                                        default:
                                                                            v125 -= 255
                                                                        }
                                                                        v125 = 888
                                                                    } else {
                                                                        var v133 uint64 = 836
                                                                        var v134 int32 = 860
                                                                        _ = v133
                                                                        _ = v134
                                                                    }
                                                                    var v135 float64 = 3.14159
                                                                    var v136 uint64 = 4294967295
                                                                    _ = v135
                                                                    _ = v136
                                                                } else {
                                                                    if v1 {
                                                                        if true {
                                                                            v2 = (v2 / ((127) | 1))
                                                                            var v137 uint = v125
                                                                            _ = v137
                                                                        }
                                                                    }
                                                                }
                                                                v2 = -64
                                                                v138 := GenS12{F0: (v4 + 65535), F1: "", F2: 127}
                                                                _ = v138
                                                            }
                                                        }
                                                    case 108:
                                                        switch "foo bar" {
                                                        case "":
                                                            fmt.Println(fi88, v9, v26)
                                                        case "foo bar":
                                                            if true {
                                                                switch ("foo bar" + "") {
                                                                case "hello":
                                                                    if ((v1 || v1) && (1 < 2147483647)) {
                                                                        v139 := GenS14{F0: fi87}
                                                                        v26 = (v26 | 2147483647)
                                                                        _ = v139
                                                                    } else {
                                                                        var v140 bool = (v26 <= v26)
                                                                        _ = v140
                                                                    }
                                                                    var v141 int64 = -103
                                                                    v142 := GenS6{F0: 375, F1: v125}
                                                                    fmt.Println(fi87, v9, fi88)
                                                                    _ = v141
                                                                    _ = v142
                                                                case "foo bar":
                                                                    fmt.Println(v9, fi88, v2)
                                                                case "\n\t":
                                                                    v4 = v4
                                                                }
                                                                var v143 float64 = 1e10
                                                                v144 := GenS7{F0: 0}
                                                                _ = v143
                                                                _ = v144
                                                            }
                                                            v26 = (v26 - (v26 % (((v26 ^ v26)) | 1)))
                                                            _ = len(v5)
                                                        }
                                                    case 9223372036854775807:
                                                        if v1 {
                                                            var v145 uint64 = 0
                                                            if len(v5) >= 2 {
                                                                v5 = v5[1:len(v5)]
                                                            }
                                                            _ = v145
                                                        }
                                                    }
                                                    v4 -= (v4 & v4)
                                                }
                                                _ = v128
                                            } else {
                                                v26 += v26
                                            }
                                            v3 -= (v3 | fi88)
                                            var v146 int8 = (v9 | v9)
                                            _ = v146
                                        }
                                        var v147 int64 = 9223372036854775807
                                        delete(v6, 826)
                                        _ = v147
                                    } else {
                                        if true {
                                            var v148 uint16 = (v4 >> ((454) & 63))
                                            if (!(!v1)) {
                                                v26 = (v26 ^ v26)
                                            }
                                            v3 += fi87
                                            _ = v148
                                        }
                                        v149 := GenS2{F0: 42}
                                        v150 := GenS15{F0: "a", F1: -221}
                                        _ = v149
                                        _ = v150
                                    }
                                    if true {
                                        v9 += (v9 >> ((255) & 63))
                                        fmt.Println(v1, v26)
                                        v151 := GenS8{F0: uint32(v4)}
                                        fmt.Println(v1)
                                        var v152 int16 = 1000
                                        _ = v151
                                        _ = v152
                                    }
                                    _ = fi88
                                    _ = v125
                                }
                                _ = fi87
                            }
                        case "a":
                            v153 := GenS16{F0: 42, F1: 723}
                            _ = v153
                        case "foo bar":
                            _ = len(v6)
                        default:
                            v154 := GenS15{F0: "\n\t", F1: 1}
                            _ = v154
                        }
                    case "\n\t":
                        v155 := GenS7{F0: -1}
                        _ = v155
                    }
                }
            case "hello":
                var v156 uint8 = 1
                _ = v156
            default:
                v9 = v9
            }
        }
    }
    var v157 int32 = 0
    v158 := GenS17{F0: v4, F1: 213, F2: (v9 + v2)}
    v5 = append(v5, (v3 % (((v3 + 0)) | 1)))
    var v159 uint8 = 255
    v2 += v9
    fmt.Println("v1:", v1)
    fmt.Println("v157:", v157)
    fmt.Println("v159:", v159)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v157
    _ = v158
    _ = v159
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
