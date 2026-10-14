package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint32; F1 int64; F2 int32 }
type GenS4 struct { F0 uint16; F1 uint8 }
type GenS5 struct { F0 int16; F1 int8; F2 uint8 }
type GenS6 struct { F0 int64; F1 bool; F2 uint }
type GenS7 struct { F0 uint64; F1 uint32 }
type GenS8 struct { F0 int; F1 bool }
type GenS9 struct { F0 int32; F1 uint8 }
type GenS10 struct { F0 uint32; F1 uint16 }
type GenS11 struct { F0 int64 }
type GenS12 struct { F0 int32; F1 uint8; F2 int32 }

func main() {
    v1 := make(map[uint32]int)
    var v2 int8 = 42
    v3 := []int{4294967295, 4294967295}
    var v4 float64 = 3.14159
    var v5 string = (("foo bar" + ("" + "")) + "foo bar")
    var v6 string = ""
    if false {
        delete(v1, 2147483648)
        for fi7 := 0; fi7 < 14; fi7++ {
            v8 := GenS1{F0: 255}
            switch 20 {
            case -902:
                var v9 int = (fi7 << ((0) & 63))
                _ = v9
            case 32768:
                if true {
                    fmt.Println(v5)
                    v4 -= v4
                }
            case 4294967295:
                var v10 int = (fi7 | fi7)
                _ = v10
            default:
                v4 = (0.0 / (v4 * -2.25))
            }
            _ = fi7
            _ = v8
        }
    }
    if false {
        for fi11 := 0; fi11 < 14; fi11++ {
            _ = len(v3)
            v12 := GenS1{F0: 32767}
            _ = fi11
            _ = v12
        }
        if (42 != 121) {
            for fi13 := 0; fi13 < 3; fi13++ {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                _ = len(v1)
                v14 := GenS1{F0: 0}
                _ = fi13
                _ = v14
            }
            var v15 uint8 = 162
            _ = v15
        } else {
            var v16 float64 = (v4 * (((-2.25 - float64(3.14159)) / float64(1e10)) / (-2.25 - 0.0)))
            v6 = v6
            switch v6 {
            case "a":
                v16 = (0.0 / 1e10)
            case "hello":
                for fi17 := 0; fi17 < 1; fi17++ {
                    v18 := GenS2{F0: false}
                    _ = fi17
                    _ = v18
                }
            case "":
                fmt.Println(v16, v16, v2)
            }
            _ = v16
        }
        var v19 int8 = v2
        _ = v19
    }
    var v20 int = 1
    for fi21 := 0; fi21 < 5; fi21++ {
        for fi22 := 0; fi22 < 5; fi22++ {
            delete(v1, 118)
            v20 = 9223372036854775807
            _ = fi22
        }
        _ = fi21
    }
    switch (v6 + v6) {
    case "\n\t":
        fmt.Println(v6)
    case "":
        v5 = "hello"
    case "hello":
        for fi23 := 0; fi23 < 13; fi23++ {
            v20 = (fi23 - 241)
            switch (v20 * fi23) {
            case 9223372036854775807:
                var v24 int8 = v2
                _ = v24
            case -9223372036854775808:
                switch v5 {
                case "hello":
                    v20 = -324
                case "foo bar":
                    if ((92 <= -503) && (163 < 228)) {
                        v1[2] = fi23
                        var v25 uint = 656
                        var v26 float32 = float32(v4)
                        _ = v25
                        _ = v26
                    }
                }
            default:
                if false {
                    switch "foo bar" {
                    case "":
                        fmt.Println(v2)
                    case "foo bar":
                        switch "" {
                        case "a":
                            v1[2] = -1
                        }
                    default:
                        for fi27 := 0; fi27 < 4; fi27++ {
                            v4 -= v4
                            if true {
                                v6 = v6
                            } else {
                                var v28 uint16 = 778
                                v20 = -217
                                _ = v28
                            }
                            _ = fi27
                        }
                    }
                    v29 := GenS2{F0: (18446744073709551615 > 4294967295)}
                    _ = v29
                } else {
                    switch (fi23 * v20) {
                    case -468:
                        var v30 int32 = -1
                        _ = v30
                    case 127:
                        var v31 uint64 = 2
                        _ = v31
                    }
                    v32 := GenS3{F0: 2, F1: -901, F2: 42}
                    switch v20 {
                    case 4294967295:
                        if false {
                            if (42 != 800) {
                                switch fi23 {
                                case 2147483647:
                                    v33 := GenS1{F0: 399}
                                    _ = v33
                                case 0:
                                    _ = len(v3)
                                default:
                                    v5 = "\n\t"
                                }
                                for fi34 := 0; fi34 < 18; fi34++ {
                                    v20 = v20
                                    var v35 uint32 = 42
                                    _ = fi34
                                    _ = v35
                                }
                                v2 += (v2 ^ 42)
                            } else {
                                var v36 uint = 4294967295
                                _ = v36
                            }
                            var v37 int8 = 127
                            _ = v37
                        }
                    case 32768:
                        v38 := GenS4{F0: 158, F1: 1}
                        _ = v38
                    case 32767:
                        fmt.Println(fi23)
                    default:
                        var v39 int64 = 42
                        _ = v39
                    }
                    _ = v32
                }
            }
            if (!(!false)) {
                if (0 <= 0) {
                    v20 += (v20 >> ((786) & 63))
                    if false {
                        _ = len(v3)
                    } else {
                        v40 := GenS5{F0: -976, F1: v2, F2: 1}
                        _ = v40
                    }
                    var v41 bool = false
                    _ = v41
                } else {
                    if true {
                        var v42 int32 = -1
                        _ = v42
                    } else {
                        var v43 int16 = 1
                        if false {
                            var v44 int64 = 42
                            _ = v44
                        }
                        _ = v43
                    }
                    if (2147483647 >= (v20 & 0)) {
                        v45 := v1[4294967295]
                        if false {
                            delete(v1, 65535)
                        } else {
                            for fi46 := 0; fi46 < 5; fi46++ {
                                v47 := GenS6{F0: 1, F1: true, F2: 636}
                                _ = fi46
                                _ = v47
                            }
                        }
                        _ = v45
                    }
                }
            }
            _ = fi23
        }
    }
    if (-704 != -1) {
        fmt.Println(v20, v6)
        if (42 < 2) {
            if true {
                var v48 uint8 = 255
                for fi49 := 0; fi49 < 15; fi49++ {
                    var v50 float64 = v4
                    for fi51 := 0; fi51 < 7; fi51++ {
                        fmt.Println(v50, v50, v2)
                        var v52 uint8 = 36
                        _ = fi51
                        _ = v52
                    }
                    _ = fi49
                    _ = v50
                }
                if (-1 == -1) {
                    v53 := GenS3{F0: 2147483647, F1: 9223372036854775807, F2: -2147483648}
                    fmt.Println(v6, v20, v6)
                    var v54 uint32 = 2147483648
                    _ = v53
                    _ = v54
                }
                _ = v48
            }
            if false {
                var v55 uint8 = 238
                switch v20 {
                case 0:
                    var v56 uint16 = 255
                    _ = v56
                case 32768:
                    var v57 uint = 65535
                    _ = v57
                }
                _ = v55
            }
        } else {
            if (!false) {
                if ("\n\t" == "foo bar") {
                    v58 := GenS7{F0: 18446744073709551615, F1: 65535}
                    _ = v58
                } else {
                    v59 := GenS3{F0: 882, F1: -9223372036854775808, F2: 655}
                    _ = v59
                }
            } else {
                for fi60 := 0; fi60 < 4; fi60++ {
                    for fi61 := 0; fi61 < 11; fi61++ {
                        v5 = v6
                        v1[2147483648] = v20
                        _ = fi61
                    }
                    switch v20 {
                    case 9223372036854775807:
                        _ = len(v1)
                    case 2147483647:
                        for fi62 := 0; fi62 < 4; fi62++ {
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            _ = fi62
                        }
                    case -452:
                        for fi63 := 0; fi63 < 2; fi63++ {
                            switch (v20 << ((65535) & 63)) {
                            case -2147483648:
                                var v64 float64 = (v4 * v4)
                                _ = v64
                            default:
                                v20 = fi63
                            }
                            var v65 int16 = 1
                            _ = fi63
                            _ = v65
                        }
                    default:
                        if ((3.14159 / 1.5) > 0.1) {
                            v66 := GenS1{F0: 42}
                            for fi67 := 0; fi67 < 16; fi67++ {
                                v4 -= v4
                                _ = fi67
                            }
                            _ = len(v1)
                            _ = v66
                        } else {
                            v2 -= (^(v2))
                            if true {
                                for fi68 := 0; fi68 < 1; fi68++ {
                                    switch (("" + v5) + ("hello" + v6)) {
                                    case "foo bar":
                                        var v69 uint32 = 645
                                        _ = v69
                                    default:
                                        fmt.Println(v20)
                                    }
                                    if true {
                                        v6 = v6
                                    }
                                    switch v6 {
                                    case "":
                                        v6 = (v5 + "\n\t")
                                        var v70 float64 = (v4 * v4)
                                        _ = v70
                                    case "a":
                                        fmt.Println(fi68, v6)
                                    case "foo bar":
                                        if (!(127 <= 97)) {
                                            v71 := GenS8{F0: 42, F1: (42 > -2147483648)}
                                            var v72 float64 = v4
                                            _ = v71
                                            _ = v72
                                        } else {
                                            fmt.Println(v4)
                                        }
                                    default:
                                        fmt.Println(fi68, v2)
                                    }
                                    switch "hello" {
                                    case "\n\t":
                                        fmt.Println(v4, v4, v5)
                                    default:
                                        fmt.Println(fi68, v20, fi68)
                                    }
                                    var v73 uint8 = 69
                                    if false {
                                        if (317 >= 4294967295) {
                                            v6 = v5
                                            var v74 uint16 = 65535
                                            _ = v74
                                        }
                                        v3 = append(v3, 9223372036854775807)
                                    } else {
                                        v4 -= 0.1
                                    }
                                    var v75 bool = true
                                    var v76 uint16 = 0
                                    var v77 int64 = -9223372036854775808
                                    _ = fi68
                                    _ = v73
                                    _ = v75
                                    _ = v76
                                    _ = v77
                                }
                                v1[4294967295] = (^(v20))
                                for fi78 := 0; fi78 < 12; fi78++ {
                                    switch v6 {
                                    case "hello":
                                        var v79 string = (v6 + ((v5 + v6) + ("foo bar" + v5)))
                                        _ = v79
                                    default:
                                        v5 = v6
                                    }
                                    fmt.Println(v4, v5, fi78)
                                    var v80 bool = false
                                    _ = fi78
                                    _ = v80
                                }
                            }
                            switch "\n\t" {
                            case "a":
                                if len(v3) > 0 {
                                    v3[0] = fi60
                                }
                            case "hello":
                                var v81 string = v6
                                _ = v81
                            case "foo bar":
                                fmt.Println(v4, v2)
                            default:
                                for fi82 := 0; fi82 < 15; fi82++ {
                                    v83 := v1[2]
                                    if (952 > 255) {
                                        var v84 string = "hello"
                                        v5 = ((v5 + "hello") + v84)
                                        switch (^(v20)) {
                                        case -1:
                                            var v85 int16 = -32768
                                            _ = v85
                                        case 127:
                                            var v86 uint64 = 673
                                            _ = v86
                                        case 42:
                                            fmt.Println(v6)
                                        }
                                        _ = v84
                                    } else {
                                        if ("a" >= "\n\t") {
                                            v87 := GenS9{F0: 119, F1: 255}
                                            v4 += v4
                                            v1[2] = v20
                                            _ = v87
                                        } else {
                                            v88 := GenS9{F0: -1, F1: 2}
                                            if true {
                                                switch (v20 & v20) {
                                                case 1:
                                                    var v89 uint64 = 42
                                                    _ = v89
                                                case 32768:
                                                    switch v83 {
                                                    case 9223372036854775807:
                                                        v20 += v20
                                                    case 128:
                                                        v20 -= (fi60 & v20)
                                                    }
                                                case -2147483648:
                                                    fmt.Println(fi82, v4, fi60)
                                                }
                                            }
                                            _ = v88
                                        }
                                        if len(v3) > 0 {
                                            v3[0] = (fi60 << ((255) & 63))
                                        }
                                        v20 += 32768
                                    }
                                    _ = fi82
                                    _ = v83
                                }
                            }
                        }
                    }
                    for fi90 := 0; fi90 < 1; fi90++ {
                        if (2147483647 != 128) {
                            var v91 float32 = float32(((1e10 * v4) + 1e10))
                            switch v5 {
                            case "hello":
                                var v92 uint16 = 0
                                _ = v92
                            case "foo bar":
                                switch (v6 + "hello") {
                                case "a":
                                    v6 = ("\n\t" + v6)
                                case "foo bar":
                                    fmt.Println(v20, v20)
                                default:
                                    var v93 float32 = 1.5
                                    _ = v93
                                }
                            default:
                                v4 -= v4
                                switch (v5 + v6) {
                                case "foo bar":
                                    var v94 float32 = 3.14159
                                    _ = v94
                                case "hello":
                                    var v95 int8 = v2
                                    _ = v95
                                }
                                _ = len(v3)
                            }
                            _ = len(v1)
                            var v96 uint = 255
                            v97 := GenS10{F0: 1, F1: 42}
                            _ = v91
                            _ = v96
                            _ = v97
                        }
                        v98 := v1[481]
                        if true {
                            switch v98 {
                            case -871:
                                v20 += (v20 >> ((2) & 63))
                            case 712:
                                var v99 uint32 = 0
                                switch (v20 & (v20 ^ fi60)) {
                                case 4294967295:
                                    var v100 uint32 = v99
                                    _ = v100
                                case 32767:
                                    var v101 uint8 = 255
                                    _ = v101
                                }
                                _ = v99
                            case 42:
                                var v102 int64 = 42
                                _ = v102
                            }
                            var v103 int64 = -1
                            _ = v103
                        }
                        var v104 int8 = v2
                        _ = fi90
                        _ = v104
                        _ = v98
                    }
                    _ = fi60
                }
                switch v6 {
                case "a":
                    v105 := GenS1{F0: 660}
                    _ = v105
                case "":
                    v6 = "\n\t"
                }
                for fi106 := 0; fi106 < 15; fi106++ {
                    var v107 int = v20
                    var v108 uint = 1
                    fmt.Println(v5)
                    _ = fi106
                    _ = v107
                    _ = v108
                }
            }
            if false {
                v109 := GenS1{F0: 2}
                _ = v109
            } else {
                if len(v3) > 0 {
                    v3[0] = -2147483648
                }
            }
        }
        v20 = (v20 ^ 504)
    }
    _ = len(v1)
    v110 := GenS11{F0: -766}
    v111 := GenS11{F0: 9223372036854775807}
    for fi112 := 0; fi112 < 13; fi112++ {
        v6 = v6
        var v113 uint8 = 2
        for fi114 := 0; fi114 < 15; fi114++ {
            var v115 int8 = v2
            v1[1] = (v20 >> ((2) & 63))
            _ = fi114
            _ = v115
        }
        _ = fi112
        _ = v113
    }
    if true {
        fmt.Println(v5, v20)
    } else {
        if len(v3) > 0 {
            v3[0] = v20
        }
        var v116 float64 = float64(((0.1 - (((3.14159 / 1.0) - 0.1) * 3.14159)) * (1e10 - float32((v4 * (3.14159 - 0.1))))))
        var v117 string = v5
        _ = v116
        _ = v117
    }
    v3 = append(v3, 127)
    if false {
        for fi118 := 0; fi118 < 14; fi118++ {
            for fi119 := 0; fi119 < 6; fi119++ {
                var v120 float32 = float32((((1.5 * ((v4 * v4) * float64(((1e10 + float32(float64((0.0 - 1e10)))) / 1e10)))) * 0.0) * 1e10))
                _ = fi119
                _ = v120
            }
            var v121 bool = false
            _ = fi118
            _ = v121
        }
        var v122 int32 = 2147483647
        _ = v122
    } else {
        v2 = v2
        for fi123 := 0; fi123 < 19; fi123++ {
            if len(v3) > 0 {
                v3[0] = (v20 << ((42) & 63))
            }
            _ = fi123
        }
        fmt.Println(v4, v5)
    }
    for fi124 := 0; fi124 < 7; fi124++ {
        fmt.Println(v6)
        v125 := GenS4{F0: 385, F1: 42}
        _ = fi124
        _ = v125
    }
    v126 := GenS12{F0: -1, F1: 1, F2: 651}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v110
    _ = v111
    _ = v126
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
