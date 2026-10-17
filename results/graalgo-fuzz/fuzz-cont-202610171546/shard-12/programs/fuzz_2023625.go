package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int64; F1 uint16 }
type GenS5 struct { F0 uint; F1 uint }
type GenS6 struct { F0 uint16; F1 uint8 }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 uint32; F1 uint8 }
type GenS9 struct { F0 string; F1 uint64; F2 int32 }
type GenS10 struct { F0 string }

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[uint32]string)
    v3 := make(map[uint32]int)
    v4 := make(map[int]int)
    v5 := []string{"a", ""}
    var v6 string = "a"
    v6 = "a"
    for fi7 := 0; fi7 < 13; fi7++ {
        for fi8 := 0; fi8 < 16; fi8++ {
            fmt.Println(fi8, fi7, fi8)
            _ = len(v1)
            _ = fi8
        }
        for fi9 := 0; fi9 < 1; fi9++ {
            v6 = ("foo bar" + v6)
            v6 = v6
            _ = fi9
        }
        _ = fi7
    }
    for fi10 := 0; fi10 < 19; fi10++ {
        var v11 int16 = -1
        switch v6 {
        case "\n\t":
            v5 = append(v5, "\n\t")
        case "hello":
            for fi12 := 0; fi12 < 13; fi12++ {
                fmt.Println(v11)
                if true {
                    v11 = -753
                    v6 = v6
                }
                _ = fi12
            }
        case "":
            var v13 int16 = (v11 << ((42) & 63))
            _ = v13
        }
        _ = fi10
        _ = v11
    }
    v6 = ("hello" + v6)
    switch "a" {
    case "":
        if (true || false) {
            if false {
                var v14 string = v6
                v6 = v6
                _ = v14
            } else {
                if (35 >= 255) {
                    fmt.Println(v6)
                    switch (v6 + v6) {
                    case "":
                        var v15 bool = false
                        _ = v15
                    case "hello":
                        if ((42 > 42) || (!true)) {
                            switch -813 {
                            case 2147483647:
                                var v16 float64 = ((3.14159 - 3.14159) + -2.25)
                                _ = v16
                            case 32768:
                                v17 := GenS1{F0: -272}
                                _ = v17
                            }
                            for fi18 := 0; fi18 < 1; fi18++ {
                                var v19 int64 = 379
                                for fi20 := 0; fi20 < 3; fi20++ {
                                    var v21 int16 = 32767
                                    if false {
                                        _ = len(v5)
                                        if (v19 > (v19 / (((v19 % ((0) | 1))) | 1))) {
                                            v6 = v6
                                        }
                                        var v22 string = v6
                                        _ = v22
                                    } else {
                                        if true {
                                            v21 = (v21 & 32767)
                                            var v23 int32 = 0
                                            _ = v23
                                        } else {
                                            v19 += -579
                                            var v24 uint64 = 2
                                            var v25 int8 = 42
                                            _ = v24
                                            _ = v25
                                        }
                                        if (42 < 18446744073709551615) {
                                            v21 -= (v21 ^ 32767)
                                            v26 := GenS1{F0: (v19 - (^(v19)))}
                                            _ = v26
                                        } else {
                                            v19 -= (v19 - (-(v19)))
                                        }
                                    }
                                    _ = fi20
                                    _ = v21
                                }
                                if (0 == 1) {
                                    if (true || (1000 != 32767)) {
                                        switch v6 {
                                        case "hello":
                                            v19 = (v19 ^ v19)
                                        case "\n\t":
                                            var v27 int8 = 42
                                            _ = v27
                                        }
                                    }
                                } else {
                                    v6 = ("a" + "hello")
                                    for fi28 := 0; fi28 < 4; fi28++ {
                                        v19 -= 513
                                        var v29 int8 = 42
                                        delete(v4, fi18)
                                        fmt.Println(v29, fi28, v29)
                                        _ = fi28
                                        _ = v29
                                    }
                                }
                                _ = fi18
                                _ = v19
                            }
                            var v30 uint = 255
                            _ = v30
                        }
                    }
                } else {
                    for fi31 := 0; fi31 < 1; fi31++ {
                        if (false || true) {
                            v6 = v6
                        } else {
                            var v32 float64 = -2.25
                            _ = v32
                        }
                        _ = fi31
                    }
                }
                var v33 int32 = -2147483648
                _ = len(v4)
                _ = v33
            }
        } else {
            if len(v5) > 0 {
                v5[0] = v6
            }
            var v34 uint16 = 65535
            for fi35 := 0; fi35 < 8; fi35++ {
                if true {
                    var v36 float64 = (((0.0 - (1.5 / ((0.1 - -2.25) + 3.14159))) / (1e10 - 3.14159)) * (float64(1.5) + (0.1 - 1e10)))
                    switch (fi35 ^ (-(fi35))) {
                    case 1:
                        var v37 uint64 = 1
                        _ = v37
                    case 0:
                        for fi38 := 0; fi38 < 20; fi38++ {
                            switch v6 {
                            case "foo bar":
                                v39 := GenS1{F0: -1}
                                var v40 uint8 = 42
                                _ = v39
                                _ = v40
                            case "a":
                                v41 := GenS1{F0: 9223372036854775807}
                                _ = v41
                            case "":
                                if (-2.25 < -2.25) {
                                    v6 = v6
                                    switch fi35 {
                                    case 2147483647:
                                        var v42 float32 = 0.0
                                        var v43 uint64 = 18446744073709551615
                                        _ = v42
                                        _ = v43
                                    case 32767:
                                        v36 = float64(-2.25)
                                    default:
                                        if (false && ((v6 + v6) >= v6)) {
                                            switch fi38 {
                                            case 1:
                                                var v44 int64 = -386
                                                _ = v44
                                            case -2147483648:
                                                if len(v5) >= 2 {
                                                    v5 = v5[1:len(v5)]
                                                }
                                            }
                                            delete(v3, 477)
                                        }
                                    }
                                    if (((-171 != 932) && true) && false) {
                                        v36 = float64((1e10 - (float32(float64(0.1)) - (1.5 / 1e10))))
                                        if true {
                                            fmt.Println(v34, fi38, fi35)
                                        } else {
                                            if true {
                                                fmt.Println(v6, v36, v34)
                                                v45 := GenS2{F0: 2}
                                                _ = v45
                                            } else {
                                                v46 := GenS1{F0: 345}
                                                _ = v46
                                            }
                                            v47 := GenS2{F0: 2}
                                            _ = v47
                                        }
                                        switch (fi38 % ((fi35) | 1)) {
                                        case 879:
                                            v6 = (v6 + ("" + (((((v6 + (v6 + (v6 + "foo bar"))) + "foo bar") + ("hello" + "")) + v6) + ("foo bar" + ((v6 + (v6 + v6)) + "foo bar")))))
                                        }
                                    }
                                } else {
                                    fmt.Println(v6, fi35, fi38)
                                }
                            default:
                                v48 := GenS2{F0: 1}
                                _ = v48
                            }
                            v6 = v6
                            _ = fi38
                        }
                    case 9223372036854775807:
                        fmt.Println(v34)
                    default:
                        v36 += 3.14159
                    }
                    _ = v36
                } else {
                    var v49 uint64 = 42
                    _ = v49
                }
                _ = fi35
            }
            _ = v34
        }
    case "\n\t":
        v50 := GenS2{F0: 0}
        _ = v50
    case "a":
        if false {
            if false {
                var v51 uint = 0
                for fi52 := 0; fi52 < 12; fi52++ {
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    v51 = (v51 << (((v51 + 453)) & 63))
                    v53 := GenS2{F0: 2}
                    _ = fi52
                    _ = v53
                }
                delete(v2, 2147483648)
                _ = v51
            }
            for fi54 := 0; fi54 < 7; fi54++ {
                var v55 bool = (159 <= 42)
                var v56 float32 = ((((3.14159 * (((1.5 + 1e10) * 0.0) + 0.0)) * -2.25) * ((-2.25 / 1.0) * 1.5)) * (((0.0 - 3.14159) + 3.14159) - float32(0.1)))
                _ = fi54
                _ = v55
                _ = v56
            }
            switch 127 {
            case -1:
                v6 = ((v6 + v6) + "\n\t")
            case 4294967295:
                var v57 bool = false
                _ = v57
            case 32767:
                var v58 int32 = 1
                _ = v58
            default:
                if true {
                    for fi59 := 0; fi59 < 12; fi59++ {
                        var v60 int8 = 127
                        if false {
                            var v61 float64 = (1.5 * (-2.25 - (0.0 + 1.5)))
                            _ = v61
                        } else {
                            if true {
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                            }
                            var v62 uint32 = 65535
                            _ = v62
                        }
                        _ = fi59
                        _ = v60
                    }
                    v6 = v6
                }
            }
        }
    }
    var v63 float64 = (3.14159 + 0.1)
    v6 = v6
    fmt.Println(v63, v63)
    switch v6 {
    case "hello":
        if true {
            var v64 int16 = 0
            _ = v64
        } else {
            if (v6 > "foo bar") {
                v65 := GenS3{F0: 1}
                v6 = ""
                _ = v65
            }
        }
    case "\n\t":
        if false {
            if false {
                for fi66 := 0; fi66 < 15; fi66++ {
                    if (87 == 128) {
                        var v67 uint32 = 2147483648
                        v67 += (^(v67))
                        _ = v67
                    }
                    v6 = (v6 + (v6 + "a"))
                    _ = fi66
                }
            }
            if true {
                if (145 < 2147483648) {
                    v63 += v63
                } else {
                    for fi68 := 0; fi68 < 10; fi68++ {
                        _ = len(v2)
                        switch (fi68 & 44) {
                        case 9223372036854775807:
                            fmt.Println(fi68, fi68, v6)
                        case -1:
                            switch (^(fi68)) {
                            case 127:
                                _ = len(v5)
                            }
                        case -851:
                            fmt.Println(v6)
                        }
                        _ = fi68
                    }
                    for fi69 := 0; fi69 < 18; fi69++ {
                        var v70 uint64 = 4294967295
                        if (false || (1e10 >= 1.5)) {
                            if false {
                                v71 := GenS4{F0: -373, F1: 32767}
                                _ = len(v5)
                                _ = v71
                            } else {
                                fmt.Println(v63, v6, v6)
                            }
                        }
                        _ = fi69
                        _ = v70
                    }
                    var v72 int64 = 1
                    _ = v72
                }
            } else {
                for fi73 := 0; fi73 < 10; fi73++ {
                    if false {
                        for fi74 := 0; fi74 < 1; fi74++ {
                            fmt.Println(fi73, fi74, v6)
                            v63 = v63
                            _ = fi74
                        }
                    } else {
                        var v75 int64 = 0
                        var v76 string = v6
                        _ = v75
                        _ = v76
                    }
                    v77 := GenS4{F0: 0, F1: 1}
                    var v78 int64 = 0
                    _ = fi73
                    _ = v77
                    _ = v78
                }
                var v79 uint16 = 432
                for fi80 := 0; fi80 < 19; fi80++ {
                    v79 -= (v79 << ((4294967295) & 63))
                    for fi81 := 0; fi81 < 20; fi81++ {
                        var v82 uint = 65535
                        if ((true && (true && ((((v6 + v6) + ("" + "a")) > "hello") || true))) || false) {
                            var v83 string = v6
                            v84 := GenS2{F0: 21}
                            v79 -= 255
                            _ = v83
                            _ = v84
                        }
                        if true {
                            v85 := v3[2147483648]
                            v79 = (v79 + (v79 - v79))
                            _ = v85
                        } else {
                            var v86 int8 = 0
                            if true {
                                v87 := GenS1{F0: -1}
                                switch (((v6 + v6) + (((("foo bar" + ("\n\t" + v6)) + (v6 + (v6 + ((v6 + ("a" + v6)) + v6)))) + v6) + "\n\t")) + "foo bar") {
                                case "hello":
                                    var v88 string = (("\n\t" + (v6 + v6)) + v6)
                                    _ = v88
                                case "a":
                                    v82 += v82
                                    var v89 uint64 = 42
                                    _ = v89
                                case "foo bar":
                                    v82 = v82
                                }
                                _ = v87
                            } else {
                                switch (("" + v6) + "foo bar") {
                                case "a":
                                    if len(v5) > 0 {
                                        v5[0] = ""
                                    }
                                    var v90 uint = 18446744073709551615
                                    _ = v90
                                default:
                                    fmt.Println(v63, v82, v79)
                                }
                                var v91 int8 = v86
                                _ = v91
                            }
                            v79 = 0
                            _ = v86
                        }
                        _ = fi81
                        _ = v82
                    }
                    v92 := GenS5{F0: 1, F1: 0}
                    _ = fi80
                    _ = v92
                }
                _ = v79
            }
            var v93 int32 = 42
            _ = v93
        } else {
            for fi94 := 0; fi94 < 9; fi94++ {
                if (58 == 255) {
                    v95 := GenS5{F0: 255, F1: 154}
                    v96 := GenS4{F0: 42, F1: 42}
                    _ = v95
                    _ = v96
                }
                _ = fi94
            }
        }
    case "":
        var v97 uint64 = 18446744073709551615
        _ = v97
    }
    if true {
        for fi98 := 0; fi98 < 3; fi98++ {
            v5 = append(v5, v6)
            _ = fi98
        }
    }
    var v99 float64 = (0.0 / (-2.25 / v63))
    v100 := GenS1{F0: 0}
    if ((42 == -32768) || false) {
        var v101 int16 = 1000
        for fi102 := 0; fi102 < 3; fi102++ {
            fmt.Println(v99, fi102, fi102)
            switch v6 {
            case "hello":
                var v103 bool = false
                _ = v103
            case "a":
                delete(v3, 4294967295)
            default:
                v63 -= v99
            }
            _ = fi102
        }
        for fi104 := 0; fi104 < 15; fi104++ {
            v3[0] = fi104
            if (false && true) {
                var v105 int64 = 42
                if true {
                    var v106 uint8 = 0
                    fmt.Println(v101)
                    v107 := GenS4{F0: (v105 >> ((947) & 63)), F1: 2}
                    _ = v106
                    _ = v107
                }
                var v108 int = 42
                _ = v105
                _ = v108
            }
            if false {
                switch "\n\t" {
                case "hello":
                    if ((128 != 164) && false) {
                        delete(v4, fi104)
                    }
                }
                if true {
                    v109 := GenS6{F0: 42, F1: 255}
                    for fi110 := 0; fi110 < 6; fi110++ {
                        var v111 uint64 = 18446744073709551615
                        _ = fi110
                        _ = v111
                    }
                    switch fi104 {
                    case 2147483647:
                        var v112 uint64 = 4294967295
                        _ = v112
                    case 32767:
                        v2[365] = (v6 + (v6 + (v6 + v6)))
                    }
                    _ = v109
                } else {
                    v113 := GenS7{F0: 255}
                    for fi114 := 0; fi114 < 4; fi114++ {
                        delete(v4, fi104)
                        v115 := GenS8{F0: 4294967295, F1: 238}
                        if (!(v6 <= v6)) {
                            fmt.Println(fi104, v6)
                            var v116 uint8 = 218
                            _ = v116
                        } else {
                            var v117 int16 = -675
                            if (fi114 > -9223372036854775808) {
                                var v118 uint64 = 42
                                _ = v118
                            } else {
                                v101 -= v101
                                fmt.Println(v101, fi104, v99)
                                if (631 == -1) {
                                    fmt.Println(v101, v63)
                                    v101 -= 80
                                } else {
                                    v119 := v3[42]
                                    fmt.Println(v63, fi104, fi114)
                                    _ = v119
                                }
                            }
                            v117 += v101
                            _ = v117
                        }
                        _ = fi114
                        _ = v115
                    }
                    var v120 string = v6
                    _ = v113
                    _ = v120
                }
            } else {
                v99 = float64(3.14159)
                for fi121 := 0; fi121 < 1; fi121++ {
                    var v122 int8 = -125
                    _ = fi121
                    _ = v122
                }
            }
            _ = fi104
        }
        _ = v101
    }
    var v123 int16 = 0
    if false {
        var v124 int64 = -1
        for fi125 := 0; fi125 < 3; fi125++ {
            switch 42 {
            case -416:
                v4[fi125] = fi125
            }
            for fi126 := 0; fi126 < 5; fi126++ {
                v127 := GenS9{F0: v6, F1: 666, F2: -138}
                _ = fi126
                _ = v127
            }
            _ = fi125
        }
        _ = v124
    } else {
        if (true || (42 != 42)) {
            if (1 == 18446744073709551615) {
                for fi128 := 0; fi128 < 12; fi128++ {
                    v63 = 3.14159
                    _ = fi128
                }
                for fi129 := 0; fi129 < 13; fi129++ {
                    v6 = v6
                    _ = fi129
                }
            }
        } else {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            var v130 int32 = 2147483647
            _ = len(v4)
            _ = v130
        }
        v99 += v63
        v63 = v99
    }
    switch v6 {
    case "hello":
        v6 = v6
    case "foo bar":
        if ("hello" < v6) {
            if true {
                fmt.Println(v99, v99, v123)
                for fi131 := 0; fi131 < 16; fi131++ {
                    var v132 uint16 = 2
                    _ = fi131
                    _ = v132
                }
            }
        }
    case "\n\t":
        var v133 uint16 = 255
        _ = v133
    }
    if (!false) {
        switch -147 {
        case -2147483648:
            v134 := GenS4{F0: 0, F1: 42}
            _ = v134
        case 238:
            v123 -= 0
        case 945:
            var v135 float32 = 0.1
            _ = v135
        }
        switch "\n\t" {
        case "":
            fmt.Println(v123, v99)
        }
    } else {
        var v136 bool = (((v6 + "\n\t") + (v6 + ("" + v6))) < "a")
        if (!false) {
            var v137 uint8 = 42
            v138 := GenS4{F0: int64(v137), F1: 0}
            v5 = append(v5, ((v6 + (v6 + v6)) + ""))
            _ = v137
            _ = v138
        } else {
            for fi139 := 0; fi139 < 6; fi139++ {
                v140 := GenS10{F0: (((v6 + (v6 + (v6 + ((v6 + v6) + v6)))) + (v6 + v6)) + (v6 + v6))}
                _ = fi139
                _ = v140
            }
            v63 = v63
        }
        _ = v136
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v123:", v123)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v63:", v63)
    fmt.Println("v99:", v99)
    _ = v1
    _ = v100
    _ = v123
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v63
    _ = v99
}
