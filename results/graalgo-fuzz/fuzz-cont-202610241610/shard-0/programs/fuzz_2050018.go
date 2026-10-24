package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int64; F1 uint; F2 int }
type GenS4 struct { F0 uint16; F1 bool; F2 uint }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 string; F1 uint16 }
type GenS8 struct { F0 int16 }
type GenS9 struct { F0 uint16; F1 string }
func genFn0() uint32 {
    var v5 int16 = 466
    if false {
        v6 := []int{-553, 127, 9223372036854775807}
        v7 := GenS1{F0: 182}
        for fi8 := 0; fi8 < 2; fi8++ {
            var v9 uint = 42
            switch (fi8 << (((v9 + 777)) & 63)) {
            case 212:
                v10 := GenS2{F0: 2147483648}
                _ = v10
            case 2147483647:
                var v11 float64 = 3.14159
                _ = v11
            }
            _ = fi8
            _ = v9
        }
        _ = v6
        _ = v7
    } else {
        for fi12 := 0; fi12 < 2; fi12++ {
            if (!false) {
                v13 := []int{2147483647}
                for fi14 := 0; fi14 < 2; fi14++ {
                    _ = len(v13)
                    _ = fi14
                }
                var v15 uint16 = 714
                _ = v13
                _ = v15
            } else {
                for fi16 := 0; fi16 < 4; fi16++ {
                    if true {
                        fmt.Println(v5)
                    }
                    if (!true) {
                        v5 -= v5
                        v17 := GenS3{F0: 1, F1: 255, F2: (fi12 & (fi12 % (((fi16 - 142)) | 1)))}
                        _ = v17
                    } else {
                        if true {
                            var v18 float32 = ((1.5 + (-2.25 / float32(3.14159))) * (0.0 / 1.0))
                            var v19 uint8 = 58
                            var v20 float64 = 1e10
                            _ = v18
                            _ = v19
                            _ = v20
                        }
                    }
                    v5 = 32767
                    var v21 uint32 = 65535
                    var v22 bool = (!(18 != 100))
                    _ = fi16
                    _ = v21
                    _ = v22
                }
                for fi23 := 0; fi23 < 3; fi23++ {
                    var v24 string = "foo bar"
                    var v25 float64 = 0.0
                    switch "a" {
                    case "":
                        switch v24 {
                        case "a":
                            v26 := []int{(-(fi12))}
                            _ = v26
                        case "hello":
                            v5 += v5
                        }
                    case "a":
                        if true {
                            switch v24 {
                            case "":
                                switch fi23 {
                                case -1:
                                    var v27 uint64 = 1
                                    _ = v27
                                case 401:
                                    if (1e10 >= 1e10) {
                                        switch v24 {
                                        case "hello":
                                            v28 := []int{}
                                            _ = v28
                                        default:
                                            fmt.Println(fi23)
                                        }
                                        v25 = 0.0
                                        v29 := GenS2{F0: 74}
                                        _ = v29
                                    }
                                    switch -434 {
                                    case -884:
                                        var v30 uint8 = 110
                                        _ = v30
                                    case 0:
                                        var v31 int32 = int32(v5)
                                        _ = v31
                                    default:
                                        v32 := GenS3{F0: -1, F1: 255, F2: (fi12 + fi12)}
                                        _ = v32
                                    }
                                default:
                                    v24 = v24
                                }
                            default:
                                var v33 int16 = 0
                                _ = v33
                            }
                        }
                    }
                    _ = fi23
                    _ = v24
                    _ = v25
                }
            }
            switch ("" + "foo bar") {
            case "foo bar":
                for fi34 := 0; fi34 < 2; fi34++ {
                    if false {
                        if false {
                            v35 := []int{}
                            _ = v35
                        } else {
                            v5 += v5
                            var v36 string = ("a" + "\n\t")
                            v5 = v5
                            _ = v36
                        }
                    }
                    var v37 uint16 = 2
                    _ = fi34
                    _ = v37
                }
            case "\n\t":
                for fi38 := 0; fi38 < 15; fi38++ {
                    var v39 uint16 = 275
                    if (!true) {
                        var v40 float32 = 0.1
                        _ = v40
                    }
                    var v41 uint64 = 666
                    _ = fi38
                    _ = v39
                    _ = v41
                }
            case "":
                v5 -= -32768
            }
            v42 := GenS3{F0: -159, F1: 1, F2: fi12}
            _ = fi12
            _ = v42
        }
    }
    switch 1 {
    case -9223372036854775808:
        fmt.Println(v5)
    case 2147483647:
        v43 := []int{-174, -1}
        _ = v43
    case 9223372036854775807:
        v44 := []int{32768, -952, 4294967295, 305}
        _ = v44
    }
    switch "a" {
    case "hello":
        var v45 float64 = float64(0.0)
        _ = v45
    case "":
        var v46 int = 32767
        _ = v46
    case "a":
        if true {
            v5 = (-(v5))
            for fi47 := 0; fi47 < 9; fi47++ {
                for fi48 := 0; fi48 < 2; fi48++ {
                    switch "hello" {
                    case "foo bar":
                        v49 := GenS4{F0: 32767, F1: false, F2: 158}
                        _ = v49
                    case "a":
                        var v50 uint8 = 102
                        _ = v50
                    case "":
                        if (42 > -1) {
                            var v51 int = fi47
                            _ = v51
                        }
                    }
                    _ = fi48
                }
                if true {
                    for fi52 := 0; fi52 < 10; fi52++ {
                        v53 := GenS3{F0: 9223372036854775807, F1: 65535, F2: (fi47 ^ (fi52 + (fi52 & fi52)))}
                        _ = fi52
                        _ = v53
                    }
                } else {
                    if false {
                        fmt.Println(v5)
                        switch 9223372036854775807 {
                        case -584:
                            var v54 int32 = 449
                            _ = v54
                        case -2147483648:
                            for fi55 := 0; fi55 < 20; fi55++ {
                                var v56 uint8 = 103
                                _ = fi55
                                _ = v56
                            }
                        case 32768:
                            v5 -= (v5 | 494)
                        }
                        switch 4294967295 {
                        case 4294967295:
                            switch "hello" {
                            case "":
                                v57 := []int{128, (^(fi47)), fi47, -1}
                                _ = v57
                            case "hello":
                                var v58 int = fi47
                                _ = v58
                            default:
                                var v59 int32 = 0
                                _ = v59
                            }
                        case -280:
                            var v60 uint32 = 334
                            _ = v60
                        default:
                            v61 := []int{(fi47 ^ -382), (fi47 << ((18446744073709551615) & 63))}
                            _ = v61
                        }
                    }
                    for fi62 := 0; fi62 < 7; fi62++ {
                        var v63 int64 = -9223372036854775808
                        _ = fi62
                        _ = v63
                    }
                    v5 += (-(v5))
                }
                fmt.Println(fi47, v5)
                _ = fi47
            }
            var v64 string = ("foo bar" + "a")
            _ = v64
        }
    default:
        var v65 uint32 = 291
        _ = v65
    }
    v66 := []int{0}
    _ = v5
    _ = v66
    return 2
}

func main() {
    var v1 float32 = 1e10
    v2 := []string{"\n\t"}
    v3 := []int{42, -557, -141}
    var v4 uint64 = 4294967295
    var v67 string = ((("a" + "hello") + ("" + "foo bar")) + ("hello" + "\n\t"))
    var v68 bool = false
    for fi69 := 0; fi69 < 6; fi69++ {
        for fi70 := 0; fi70 < 18; fi70++ {
            fmt.Println(fi69, fi69, v67)
            fmt.Println(v1)
            _ = fi70
        }
        _ = fi69
    }
    var v71 uint32 = genFn0()
    if (32767 >= 1) {
        switch ("\n\t" + ("foo bar" + ((("" + v67) + (v67 + "\n\t")) + v67))) {
        case "a":
            var v72 string = v67
            _ = v72
        default:
            v73 := GenS5{F0: 305}
            _ = v73
        }
    } else {
        v74 := GenS6{F0: 32767}
        _ = v74
    }
    if v68 {
        v71 = genFn0()
        for fi75 := 0; fi75 < 5; fi75++ {
            if (!v68) {
                if v68 {
                    v76 := GenS6{F0: 1}
                    v71 = v71
                    switch v67 {
                    case "a":
                        fmt.Println(v68, fi75)
                    }
                    _ = v76
                } else {
                    fmt.Println(fi75)
                    v67 = (("\n\t" + v67) + ("foo bar" + ("" + v67)))
                }
            } else {
                for fi77 := 0; fi77 < 14; fi77++ {
                    if (!v68) {
                        if (-91 <= 48) {
                            v68 = true
                        }
                        if (!v68) {
                            v4 = v4
                            var v78 uint = 255
                            _ = v78
                        }
                        fmt.Println(fi77, fi77)
                    }
                    var v79 int16 = 0
                    var v80 int = fi75
                    switch 32768 {
                    case 128:
                        v81 := GenS7{F0: "foo bar", F1: 32767}
                        _ = v81
                    case -1:
                        if true {
                            fmt.Println(v4, v71)
                            var v82 float32 = v1
                            _ = v82
                        }
                    case 127:
                        if true {
                            var v83 string = (v67 + v67)
                            _ = v83
                        }
                    default:
                        v79 = v79
                    }
                    _ = fi77
                    _ = v79
                    _ = v80
                }
                for fi84 := 0; fi84 < 6; fi84++ {
                    fmt.Println(v71, v4, v67)
                    var v85 float32 = (((v1 - v1) - (v1 + ((-2.25 / (float32(3.14159) * v1)) - v1))) - (v1 - v1))
                    switch (v67 + ("foo bar" + (((("a" + ("a" + (v67 + v67))) + v67) + ("foo bar" + "foo bar")) + v67))) {
                    case "a":
                        if (!v68) {
                            var v86 int32 = 0
                            if (174 != 273) {
                                v1 = v1
                                switch "a" {
                                case "":
                                    var v87 uint = uint(v71)
                                    if (!true) {
                                        switch ((v67 + v67) + v67) {
                                        case "":
                                            if v68 {
                                                var v88 uint64 = 4294967295
                                                var v89 uint8 = 42
                                                v4 += (^(v4))
                                                _ = v88
                                                _ = v89
                                            }
                                        case "foo bar":
                                            v71 -= v71
                                        }
                                    } else {
                                        var v90 float32 = 0.0
                                        _ = v90
                                    }
                                    fmt.Println(fi75)
                                    _ = v87
                                default:
                                    if (!(!true)) {
                                        v1 = 3.14159
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                    } else {
                                        var v91 uint8 = 11
                                        v67 = v67
                                        v91 += 1
                                        v67 = v67
                                        if v68 {
                                            v67 = v67
                                        }
                                        var v92 uint64 = 42
                                        v68 = (!(0 <= v71))
                                        _ = v91
                                        _ = v92
                                    }
                                }
                            }
                            _ = v86
                        }
                    default:
                        var v93 uint16 = 255
                        _ = v93
                    }
                    _ = fi84
                    _ = v85
                }
                v94 := GenS4{F0: 42, F1: true, F2: 2}
                _ = v94
            }
            v1 = 3.14159
            if ((-2.25 > v1) && v68) {
                v71 += v71
            } else {
                v95 := GenS5{F0: 752}
                var v96 float64 = (0.0 * -2.25)
                v68 = v68
                _ = v95
                _ = v96
            }
            _ = fi75
        }
    }
    var v97 int = 2147483647
    switch (v97 << ((255) & 63)) {
    case -9223372036854775808:
        v1 = 0.0
    case 127:
        if (!((!(1 < 0)) && true)) {
            for fi98 := 0; fi98 < 14; fi98++ {
                var v99 float64 = 3.14159
                switch v97 {
                case -682:
                    for fi100 := 0; fi100 < 13; fi100++ {
                        if (!((v4 >= (v4 - v4)) || (!v68))) {
                            switch (("hello" + "foo bar") + v67) {
                            case "foo bar":
                                if (185 <= -767) {
                                    v1 += v1
                                } else {
                                    var v101 uint64 = v4
                                    var v102 uint16 = 32767
                                    _ = v101
                                    _ = v102
                                }
                            case "hello":
                                var v103 int16 = 1000
                                _ = v103
                            default:
                                if v68 {
                                    var v104 int = (fi98 + fi100)
                                    switch -798 {
                                    case 823:
                                        switch fi98 {
                                        case 314:
                                            fmt.Println(v97, v99)
                                        case 42:
                                            v99 = (float64((-2.25 / 0.1)) + (((v99 * ((v99 + v99) - float64(0.1))) / 1.5) - (v99 * v99)))
                                        case 0:
                                            v68 = (!true)
                                        }
                                    case 2147483647:
                                        v2 = append(v2, ("\n\t" + ("\n\t" + v67)))
                                    case -9223372036854775808:
                                        var v105 uint = 0
                                        _ = v105
                                    }
                                    _ = v104
                                } else {
                                    var v106 uint8 = 2
                                    _ = v106
                                }
                            }
                            var v107 string = v67
                            _ = v107
                        } else {
                            if ((v71 * 2) < v71) {
                                v68 = ((0 != 1) || v68)
                                v1 = 0.0
                                v108 := GenS4{F0: 210, F1: true, F2: 255}
                                _ = v108
                            }
                            v97 = (fi100 - (fi100 >> ((168) & 63)))
                            if false {
                                fmt.Println(v97, v1)
                            } else {
                                fmt.Println(v68, v97, v71)
                                fmt.Println(v71, v67, fi100)
                                var v109 int32 = -1
                                _ = v109
                            }
                        }
                        _ = fi100
                    }
                case 1:
                    v99 += v99
                case 127:
                    _ = len(v2)
                default:
                    for fi110 := 0; fi110 < 4; fi110++ {
                        switch v67 {
                        case "foo bar":
                            fmt.Println(v68, v71)
                        case "":
                            v4 = (v4 ^ 4294967295)
                        default:
                            fmt.Println(v71, v1)
                        }
                        _ = fi110
                    }
                }
                _ = fi98
                _ = v99
            }
            _ = len(v2)
        } else {
            for fi111 := 0; fi111 < 15; fi111++ {
                if (((v97 >> ((uint(v4)) & 63)) >= v97) || false) {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    v4 = (v4 & (v4 << ((4294967295) & 63)))
                }
                switch v97 {
                case 527:
                    if len(v2) > 0 {
                        v2[0] = "foo bar"
                    }
                }
                v112 := GenS1{F0: 1}
                _ = fi111
                _ = v112
            }
            if (1 < 42) {
                var v113 bool = v68
                _ = v113
            } else {
                switch v97 {
                case 4294967295:
                    v1 -= (float32(((0.1 - -2.25) + 0.0)) - (v1 * v1))
                case 42:
                    v2 = append(v2, v67)
                default:
                    var v114 int32 = 2147483647
                    _ = v114
                }
                if true {
                    v68 = true
                }
                if v68 {
                    for fi115 := 0; fi115 < 18; fi115++ {
                        switch (fi115 | (fi115 * -799)) {
                        case 246:
                            if v68 {
                                v2 = append(v2, v67)
                                switch (((v67 + v67) + "hello") + v67) {
                                case "\n\t":
                                    if v68 {
                                        v4 = (^(v4))
                                        switch (v67 + (v67 + ("" + v67))) {
                                        case "\n\t":
                                            v116 := GenS8{F0: 626}
                                            _ = v116
                                        case "":
                                            v2 = append(v2, "\n\t")
                                        case "hello":
                                            if (v71 == genFn0()) {
                                                v117 := GenS8{F0: 81}
                                                _ = v117
                                            } else {
                                                var v118 float64 = 0.0
                                                var v119 int8 = -66
                                                _ = v118
                                                _ = v119
                                            }
                                        }
                                    }
                                case "hello":
                                    var v120 float64 = -2.25
                                    _ = v120
                                default:
                                    v121 := GenS5{F0: 18446744073709551615}
                                    _ = v121
                                }
                                for fi122 := 0; fi122 < 16; fi122++ {
                                    v71 -= (v71 >> ((4294967295) & 63))
                                    _ = fi122
                                }
                            }
                        case 110:
                            var v123 bool = (1 > 0)
                            _ = v123
                        default:
                            v1 -= v1
                        }
                        _ = fi115
                    }
                } else {
                    var v124 uint16 = 72
                    switch -9223372036854775808 {
                    case -1:
                        if ((false && (127 < 1)) && v68) {
                            fmt.Println(v124)
                            var v125 uint8 = 128
                            var v126 uint = uint(v4)
                            _ = v125
                            _ = v126
                        }
                    case 1:
                        switch v67 {
                        case "":
                            var v127 int32 = 709
                            _ = v127
                        case "\n\t":
                            _ = len(v3)
                        default:
                            fmt.Println(v71)
                        }
                    case -2147483648:
                        var v128 int16 = -1
                        _ = v128
                    }
                    _ = v124
                }
            }
            if (-97 > -38) {
                fmt.Println(v68)
                var v129 int64 = 989
                _ = v129
            } else {
                v67 = v67
                for fi130 := 0; fi130 < 11; fi130++ {
                    v4 -= v4
                    fmt.Println(v68, v71)
                    _ = fi130
                }
                for fi131 := 0; fi131 < 4; fi131++ {
                    var v132 int = 1
                    switch "a" {
                    case "foo bar":
                        if v68 {
                            var v133 uint8 = uint8(v4)
                            _ = v133
                        } else {
                            var v134 uint = 65535
                            _ = v134
                        }
                    case "hello":
                        switch "foo bar" {
                        case "":
                            var v135 string = v67
                            _ = v135
                        }
                    case "":
                        fmt.Println(v4, v132)
                    }
                    _ = fi131
                    _ = v132
                }
            }
        }
    default:
        if false {
            v68 = v68
            fmt.Println(v4)
        }
    }
    if ((v97 & (v97 / (((v97 | (v97 + 4294967295))) | 1))) < v97) {
        v4 = (v4 % ((2) | 1))
        fmt.Println(v97)
    } else {
        switch v67 {
        case "\n\t":
            if v68 {
                var v136 int16 = -605
                if (v97 >= 42) {
                    v97 -= int(v136)
                    var v137 int8 = 0
                    v2 = append(v2, ((v67 + (v67 + "hello")) + "hello"))
                    _ = v137
                } else {
                    v4 += 18446744073709551615
                    for fi138 := 0; fi138 < 11; fi138++ {
                        _ = len(v3)
                        for fi139 := 0; fi139 < 4; fi139++ {
                            var v140 float32 = v1
                            v1 -= v1
                            if len(v2) > 0 {
                                v2[0] = (v67 + v67)
                            }
                            var v141 string = "\n\t"
                            v136 = v136
                            v142 := GenS3{F0: -242, F1: 42, F2: 127}
                            _ = fi139
                            _ = v140
                            _ = v141
                            _ = v142
                        }
                        _ = fi138
                    }
                }
                _ = v136
            }
        case "":
            switch 75 {
            case -1000:
                var v143 int8 = 100
                _ = v143
            case 128:
                for fi144 := 0; fi144 < 3; fi144++ {
                    switch fi144 {
                    case -282:
                        v97 += (v97 % ((fi144) | 1))
                    default:
                        v4 -= 2
                    }
                    _ = fi144
                }
            case -361:
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            }
        case "a":
            v145 := GenS7{F0: (v67 + v67), F1: 226}
            _ = v145
        }
        v146 := GenS2{F0: (v71 + 765)}
        if (-128 >= 0) {
            for fi147 := 0; fi147 < 9; fi147++ {
                if v68 {
                    v67 = v67
                }
                if len(v3) > 0 {
                    v3[0] = v97
                }
                fmt.Println(fi147, v71, v67)
                _ = fi147
            }
            v4 -= 2
            fmt.Println(v67, v1, v71)
        }
        _ = v146
    }
    var v148 int32 = 0
    fmt.Println(v67)
    fmt.Println(v97)
    var v149 float64 = 0.1
    fmt.Println(v1)
    v150 := GenS9{F0: 42, F1: "\n\t"}
    fmt.Println(v68)
    fmt.Println("v1:", v1)
    fmt.Println("v148:", v148)
    fmt.Println("v149:", v149)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v67:", v67)
    fmt.Println("v68:", v68)
    fmt.Println("v71:", v71)
    fmt.Println("v97:", v97)
    _ = v1
    _ = v148
    _ = v149
    _ = v150
    _ = v2
    _ = v3
    _ = v4
    _ = v67
    _ = v68
    _ = v71
    _ = v97
}
