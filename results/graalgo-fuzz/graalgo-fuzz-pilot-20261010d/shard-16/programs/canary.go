package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8; F2 uint8 }
type GenS2 struct { F0 string; F1 int8 }
type GenS3 struct { F0 string; F1 int16; F2 bool }
type GenS4 struct { F0 uint16; F1 uint32 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 uint8; F1 int16 }
type GenS7 struct { F0 uint64; F1 int64 }
type GenS8 struct { F0 string }
type GenS9 struct { F0 int8; F1 int8 }
func genFn0(p0 int32, p1 uint64, p2 bool) uint8 {
    for fi5 := 0; fi5 < 10; fi5++ {
        p0 -= p0
        _ = fi5
    }
    switch "\n\t" {
    case "\n\t":
        v6 := GenS1{F0: 42, F1: 255, F2: 2}
        _ = v6
    }
    var v7 uint64 = p1
    v8 := GenS1{F0: -102, F1: 127, F2: 42}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 255
}
func genFn1(p0 string, p1 uint) uint32 {
    var v9 uint32 = 0
    p1 = p1
    _ = p0
    _ = p1
    _ = v9
    return (^(v9))
}

func main() {
    v1 := make(map[uint32]string)
    var v2 string = "foo bar"
    var v3 uint8 = 255
    var v4 uint = 12
    fmt.Println(v2)
    v2 = (("hello" + "foo bar") + ("" + "a"))
    if true {
        switch "foo bar" {
        case "\n\t":
            v10 := []int{128, -491, 128, -9223372036854775808}
            _ = v10
        case "hello":
            for fi11 := 0; fi11 < 19; fi11++ {
                if ((((float32(0.0) + float32(3.14159)) < 1.5) && false) && true) {
                    if false {
                        switch "hello" {
                        case "foo bar":
                            v2 = v2
                        default:
                            switch v2 {
                            case "hello":
                                v12 := GenS2{F0: (("\n\t" + v2) + v2), F1: -128}
                                _ = v12
                            case "a":
                                v13 := GenS3{F0: "foo bar", F1: -590, F2: true}
                                _ = v13
                            case "foo bar":
                                switch (v2 + (v2 + ("\n\t" + v2))) {
                                case "\n\t":
                                    var v14 uint8 = (v3 - (v3 ^ (v3 % (((v3 | (v3 + v3))) | 1))))
                                    _ = v14
                                }
                            }
                        }
                        v4 = v4
                    }
                    v4 = (v4 | v4)
                } else {
                    fmt.Println(v2, v2, v3)
                    for fi15 := 0; fi15 < 9; fi15++ {
                        if true {
                            v16 := GenS4{F0: 255, F1: 2147483648}
                            for fi17 := 0; fi17 < 10; fi17++ {
                                v18 := GenS5{F0: 0}
                                if false {
                                    v19 := GenS5{F0: 65535}
                                    _ = v19
                                }
                                if (!false) {
                                    v20 := GenS5{F0: 42}
                                    _ = v20
                                }
                                _ = fi17
                                _ = v18
                            }
                            var v21 int32 = 0
                            _ = v16
                            _ = v21
                        } else {
                            var v22 float64 = (0.0 - ((float64(0.0) / (0.0 + float64(1.5))) / (1e10 / (float64((1.5 + (0.0 / ((-2.25 / 1.5) / 0.1)))) / 1.0))))
                            v4 += v4
                            var v23 uint = 255
                            _ = v22
                            _ = v23
                        }
                        _ = fi15
                    }
                }
                var v24 uint16 = 42
                if (60 > -1) {
                    v1[1] = ((v2 + v2) + (v2 + "\n\t"))
                    var v25 uint8 = (v3 << (((v4 - 278)) & 63))
                    switch v2 {
                    case "":
                        v1[65535] = v2
                    case "foo bar":
                        if true {
                            fmt.Println(v4, v3)
                        } else {
                            var v26 bool = (!true)
                            if (v26 && v26) {
                                switch ((v2 + "\n\t") + ((v2 + v2) + v2)) {
                                case "hello":
                                    if (184 >= 2) {
                                        fmt.Println(v3, v24)
                                        var v27 int8 = 42
                                        _ = v27
                                    }
                                case "a":
                                    fmt.Println(v26, v3)
                                }
                                var v28 uint64 = 0
                                var v29 uint = v4
                                _ = v28
                                _ = v29
                            } else {
                                var v30 int = 32768
                                if v26 {
                                    switch -1 {
                                    case 32768:
                                        if (65535 == 21) {
                                            if (340 > 22) {
                                                switch "foo bar" {
                                                case "":
                                                    fmt.Println(v2, v3)
                                                case "hello":
                                                    if (-1 <= 747) {
                                                        if ((v26 || true) && (!(v2 >= v2))) {
                                                            for fi31 := 0; fi31 < 2; fi31++ {
                                                                fmt.Println(fi11, v26, v25)
                                                                var v32 int32 = -2147483648
                                                                v26 = true
                                                                _ = fi31
                                                                _ = v32
                                                            }
                                                            switch (v30 & 187) {
                                                            case 9223372036854775807:
                                                                var v33 int = (v30 | (v30 ^ -606))
                                                                _ = v33
                                                            case 1:
                                                                var v34 string = v2
                                                                _ = v34
                                                            case 0:
                                                                v26 = v26
                                                            default:
                                                                switch (((("a" + v2) + v2) + v2) + v2) {
                                                                case "hello":
                                                                    v24 = (v24 / (((v24 ^ v24)) | 1))
                                                                default:
                                                                    fmt.Println(v3, v30, fi11)
                                                                }
                                                            }
                                                            var v35 uint64 = 1
                                                            _ = v35
                                                        }
                                                        var v36 int32 = -846
                                                        _ = v36
                                                    }
                                                default:
                                                    switch 32767 {
                                                    case 0:
                                                        v4 = (v4 + v4)
                                                    case 4294967295:
                                                        var v37 uint64 = 0
                                                        _ = v37
                                                    }
                                                }
                                            } else {
                                                switch (((v2 + (v2 + v2)) + v2) + (v2 + v2)) {
                                                case "":
                                                    for fi38 := 0; fi38 < 10; fi38++ {
                                                        for fi39 := 0; fi39 < 6; fi39++ {
                                                            for fi40 := 0; fi40 < 11; fi40++ {
                                                                v4 = v4
                                                                _ = fi40
                                                            }
                                                            if v26 {
                                                                delete(v1, 802)
                                                                var v41 bool = (true || true)
                                                                fmt.Println(fi11, v4, fi39)
                                                                _ = v41
                                                            } else {
                                                                fmt.Println(v26)
                                                            }
                                                            _ = fi39
                                                        }
                                                        _ = fi38
                                                    }
                                                }
                                            }
                                            var v42 uint8 = (v3 << ((838) & 63))
                                            _ = v42
                                        }
                                    }
                                    v43 := []int{}
                                    _ = v43
                                }
                                if (3.14159 <= 1e10) {
                                    if true {
                                        for fi44 := 0; fi44 < 14; fi44++ {
                                            for fi45 := 0; fi45 < 11; fi45++ {
                                                for fi46 := 0; fi46 < 14; fi46++ {
                                                    if (-2147483648 > -7) {
                                                        v47 := GenS1{F0: -1, F1: uint8(v24), F2: genFn0(2147483647, 42, ((fi46 & 32768) < fi11))}
                                                        _ = v47
                                                    }
                                                    v48 := []int{}
                                                    _ = fi46
                                                    _ = v48
                                                }
                                                switch ((v2 + (v2 + "\n\t")) + "a") {
                                                case "a":
                                                    fmt.Println(v2, fi45)
                                                case "foo bar":
                                                    v2 = v2
                                                case "":
                                                    fmt.Println(v4)
                                                default:
                                                    v49 := GenS2{F0: v2, F1: 5}
                                                    _ = v49
                                                }
                                                _ = fi45
                                            }
                                            _ = fi44
                                        }
                                        v50 := GenS1{F0: -29, F1: (v3 & (^(v3))), F2: v25}
                                        v51 := []int{fi11, (fi11 | (v30 << ((v4) & 63))), 128, fi11}
                                        _ = v50
                                        _ = v51
                                    } else {
                                        if v26 {
                                            switch fi11 {
                                            case -19:
                                                v30 = v30
                                            case 9223372036854775807:
                                                v26 = v26
                                            }
                                            var v52 uint = v4
                                            _ = v52
                                        } else {
                                            v53 := v1[65535]
                                            for fi54 := 0; fi54 < 9; fi54++ {
                                                for fi55 := 0; fi55 < 4; fi55++ {
                                                    if (fi11 < (v30 << (((v4 & 1)) & 63))) {
                                                        var v56 uint16 = uint16(v4)
                                                        switch (fi54 * v30) {
                                                        case 128:
                                                            var v57 uint = v4
                                                            _ = v57
                                                        case 32768:
                                                            fmt.Println(fi55)
                                                        case 4294967295:
                                                            for fi58 := 0; fi58 < 4; fi58++ {
                                                                v59 := v1[4294967295]
                                                                var v60 uint16 = 402
                                                                for fi61 := 0; fi61 < 16; fi61++ {
                                                                    for fi62 := 0; fi62 < 8; fi62++ {
                                                                        var v63 bool = (!v26)
                                                                        v59 = (v2 + ("hello" + v53))
                                                                        _ = fi62
                                                                        _ = v63
                                                                    }
                                                                    var v64 int16 = -1
                                                                    _ = fi61
                                                                    _ = v64
                                                                }
                                                                _ = fi58
                                                                _ = v59
                                                                _ = v60
                                                            }
                                                        default:
                                                            if true {
                                                                for fi65 := 0; fi65 < 14; fi65++ {
                                                                    v66 := []int{(fi54 / ((fi65) | 1))}
                                                                    for fi67 := 0; fi67 < 10; fi67++ {
                                                                        for fi68 := 0; fi68 < 1; fi68++ {
                                                                            switch (v30 % ((-9223372036854775808) | 1)) {
                                                                            case 0:
                                                                                for fi69 := 0; fi69 < 6; fi69++ {
                                                                                    switch fi67 {
                                                                                    case 128:
                                                                                        var v70 int8 = 100
                                                                                        _ = v70
                                                                                    case -1:
                                                                                        switch v2 {
                                                                                        case "a":
                                                                                            var v71 uint16 = 65535
                                                                                            _ = v71
                                                                                        default:
                                                                                            var v72 float64 = 1e10
                                                                                            _ = v72
                                                                                        }
                                                                                    case 2147483647:
                                                                                        v26 = v26
                                                                                    }
                                                                                    if v26 {
                                                                                        var v73 string = v2
                                                                                        v56 = uint16(v4)
                                                                                        switch v2 {
                                                                                        case "a":
                                                                                            if len(v66) >= 2 {
                                                                                                v66 = v66[1:len(v66)]
                                                                                            }
                                                                                        case "foo bar":
                                                                                            var v74 string = (v73 + v53)
                                                                                            _ = v74
                                                                                        }
                                                                                        _ = v73
                                                                                    }
                                                                                    v75 := GenS1{F0: 1, F1: (v25 & (v25 >> ((0) & 63))), F2: v3}
                                                                                    _ = fi69
                                                                                    _ = v75
                                                                                }
                                                                            case 9223372036854775807:
                                                                                switch (v30 ^ fi67) {
                                                                                case -9223372036854775808:
                                                                                    switch v2 {
                                                                                    case "":
                                                                                        v76 := GenS1{F0: 1, F1: (v3 << ((v4) & 63)), F2: v25}
                                                                                        _ = v76
                                                                                    case "foo bar":
                                                                                        var v77 string = (((("foo bar" + ("hello" + v2)) + "a") + (v2 + (v53 + v2))) + v2)
                                                                                        _ = v77
                                                                                    case "a":
                                                                                        var v78 int = fi65
                                                                                        _ = v78
                                                                                    default:
                                                                                        var v79 int16 = 1000
                                                                                        _ = v79
                                                                                    }
                                                                                case -972:
                                                                                    if ((0 >= 9223372036854775807) || v26) {
                                                                                        v30 -= int(v24)
                                                                                        var v80 uint16 = (v56 + (v24 + (v24 << ((uint(v25)) & 63))))
                                                                                        switch fi67 {
                                                                                        case 127:
                                                                                            if (1 >= 0) {
                                                                                                for fi81 := 0; fi81 < 19; fi81++ {
                                                                                                    switch (fi67 * (fi11 + (-(v30)))) {
                                                                                                    case 32768:
                                                                                                        v53 = v53
                                                                                                    case 4294967295:
                                                                                                        fmt.Println(v2, v56, fi11)
                                                                                                    default:
                                                                                                        v82 := GenS6{F0: uint8(v80), F1: 42}
                                                                                                        _ = v82
                                                                                                    }
                                                                                                    _ = fi81
                                                                                                }
                                                                                                switch "hello" {
                                                                                                case "":
                                                                                                    _ = len(v1)
                                                                                                }
                                                                                            } else {
                                                                                                v2 = (v2 + v53)
                                                                                                fmt.Println(v3)
                                                                                            }
                                                                                        case 32768:
                                                                                            _ = len(v66)
                                                                                        }
                                                                                        _ = v80
                                                                                    }
                                                                                case 32768:
                                                                                    v56 -= 0
                                                                                }
                                                                            default:
                                                                                if len(v66) >= 2 {
                                                                                    v66 = v66[1:len(v66)]
                                                                                }
                                                                            }
                                                                            var v83 uint64 = 2
                                                                            if ((^(v24)) != (v24 >> ((uint(v83)) & 63))) {
                                                                                v84 := v1[570]
                                                                                v85 := GenS5{F0: uint32(v83)}
                                                                                v1[genFn1(v53, v4)] = v84
                                                                                _ = v84
                                                                                _ = v85
                                                                            }
                                                                            _ = fi68
                                                                            _ = v83
                                                                        }
                                                                        _ = len(v1)
                                                                        if (523 <= 1) {
                                                                            v86 := GenS4{F0: (v56 * 65535), F1: 571}
                                                                            _ = v86
                                                                        }
                                                                        _ = fi67
                                                                    }
                                                                    _ = fi65
                                                                    _ = v66
                                                                }
                                                                if v26 {
                                                                    for fi87 := 0; fi87 < 18; fi87++ {
                                                                        if v26 {
                                                                            var v88 uint8 = genFn0(881, 0, ((fi55 | fi54) >= fi55))
                                                                            _ = v88
                                                                        } else {
                                                                            v89 := GenS7{F0: 750, F1: -9223372036854775808}
                                                                            _ = v89
                                                                        }
                                                                        _ = fi87
                                                                    }
                                                                }
                                                                v4 -= v4
                                                            }
                                                        }
                                                        _ = v56
                                                    }
                                                    v53 = (v53 + v53)
                                                    _ = fi55
                                                }
                                                fmt.Println(v26)
                                                _ = fi54
                                            }
                                            for fi90 := 0; fi90 < 18; fi90++ {
                                                fmt.Println(v25)
                                                var v91 uint64 = 2
                                                _ = fi90
                                                _ = v91
                                            }
                                            _ = v53
                                        }
                                        if false {
                                            switch (v30 ^ fi11) {
                                            case 1:
                                                v4 = (v4 + (v4 & (^(v4))))
                                            }
                                            v25 += (v3 | (v25 ^ (v3 / ((v3) | 1))))
                                        } else {
                                            v92 := []int{42, (fi11 - v30), v30, fi11}
                                            var v93 uint64 = 0
                                            fmt.Println(v30, v24)
                                            _ = v92
                                            _ = v93
                                        }
                                        switch v2 {
                                        case "\n\t":
                                            var v94 int64 = -9223372036854775808
                                            _ = v94
                                        case "":
                                            if (v24 >= uint16(v25)) {
                                                var v95 bool = (9223372036854775807 >= v30)
                                                v96 := []int{fi11, 32767, int(v3), (v30 & (-(fi11)))}
                                                _ = v95
                                                _ = v96
                                            }
                                        case "hello":
                                            v97 := GenS2{F0: v2, F1: 42}
                                            _ = v97
                                        }
                                    }
                                    v26 = (v26 && (2147483648 >= 1))
                                    var v98 uint32 = 4294967295
                                    _ = v98
                                } else {
                                    _ = len(v1)
                                }
                                _ = v30
                            }
                            _ = v26
                        }
                    default:
                        for fi99 := 0; fi99 < 20; fi99++ {
                            fmt.Println(v2, fi11, v4)
                            var v100 int16 = -954
                            if (!(int16(v25) > 0)) {
                                var v101 float32 = (3.14159 + 3.14159)
                                _ = v101
                            }
                            _ = fi99
                            _ = v100
                        }
                    }
                    _ = v25
                } else {
                    v102 := GenS8{F0: v2}
                    _ = v102
                }
                _ = fi11
                _ = v24
            }
        case "":
            var v103 uint8 = 247
            _ = v103
        default:
            delete(v1, 2)
        }
        var v104 float32 = float32((-2.25 * 3.14159))
        _ = v104
    }
    v105 := []int{}
    if (true || (v3 > 1)) {
        fmt.Println(v3, v2, v4)
        var v106 int = 4294967295
        _ = v106
    } else {
        fmt.Println(v2, v3)
    }
    v107 := GenS9{F0: 1, F1: -81}
    _ = len(v1)
    fmt.Println(v2, v2, v3)
    var v108 float64 = 0.0
    for fi109 := 0; fi109 < 14; fi109++ {
        switch (("a" + v2) + (v2 + ("foo bar" + ((("foo bar" + v2) + "") + v2)))) {
        case "\n\t":
            var v110 int16 = 42
            _ = v110
        case "hello":
            fmt.Println(v4, v2)
        case "foo bar":
            v111 := v1[genFn1((v2 + "foo bar"), v4)]
            _ = v111
        default:
            for fi112 := 0; fi112 < 13; fi112++ {
                var v113 int16 = -293
                var v114 uint32 = 2147483647
                if len(v105) > 0 {
                    v105[0] = fi112
                }
                _ = fi112
                _ = v113
                _ = v114
            }
        }
        var v115 uint8 = genFn0(-2147483648, 106, (3.14159 != v108))
        if false {
            var v116 uint32 = 2147483648
            var v117 int8 = -65
            _ = v116
            _ = v117
        } else {
            switch v2 {
            case "foo bar":
                for fi118 := 0; fi118 < 7; fi118++ {
                    if (!(42 == 949)) {
                        for fi119 := 0; fi119 < 17; fi119++ {
                            if (!true) {
                                if false {
                                    _ = len(v1)
                                } else {
                                    switch (fi109 ^ 127) {
                                    case 304:
                                        v2 = ""
                                    case 2147483647:
                                        var v120 int64 = 353
                                        _ = v120
                                    case 127:
                                        if (fi109 >= fi118) {
                                            v4 -= (v4 * (v4 >> (((v4 >> ((v4) & 63))) & 63)))
                                            var v121 uint16 = 42
                                            v108 -= v108
                                            _ = v121
                                        }
                                    }
                                }
                                var v122 uint = v4
                                _ = v122
                            } else {
                                var v123 uint64 = 927
                                _ = v123
                            }
                            var v124 int64 = 42
                            _ = len(v1)
                            _ = fi119
                            _ = v124
                        }
                        if true {
                            var v125 int32 = 1
                            switch fi109 {
                            case 4294967295:
                                v1[2147483647] = (v2 + v2)
                            }
                            switch "\n\t" {
                            case "\n\t":
                                for fi126 := 0; fi126 < 13; fi126++ {
                                    v108 -= 0.1
                                    _ = fi126
                                }
                            }
                            _ = v125
                        }
                    } else {
                        v3 += 42
                        var v127 int8 = -1
                        _ = v127
                    }
                    v128 := GenS2{F0: "foo bar", F1: 0}
                    _ = fi118
                    _ = v128
                }
            case "a":
                v129 := GenS6{F0: (v115 & 99), F1: 1}
                _ = v129
            case "":
                for fi130 := 0; fi130 < 19; fi130++ {
                    switch (fi109 * fi130) {
                    case -694:
                        var v131 int16 = 0
                        _ = v131
                    }
                    _ = fi130
                }
            }
            var v132 uint64 = 218
            v2 = v2
            _ = v132
        }
        _ = fi109
        _ = v115
    }
    switch (v2 + "foo bar") {
    case "hello":
        if (false && (!true)) {
            v2 = "hello"
            fmt.Println(v4, v4, v4)
            fmt.Println(v4)
        }
    default:
        if true {
            if false {
                var v133 uint16 = 32767
                var v134 string = v2
                var v135 int = -230
                _ = v133
                _ = v134
                _ = v135
            }
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v105:", len(v105), v105)
    fmt.Println("v108:", v108)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v105
    _ = v107
    _ = v108
    _ = v2
    _ = v3
    _ = v4
}
