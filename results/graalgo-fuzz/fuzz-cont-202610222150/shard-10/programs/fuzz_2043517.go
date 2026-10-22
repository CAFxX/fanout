package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int8; F1 int }
type GenS3 struct { F0 uint8; F1 bool; F2 int8 }
type GenS4 struct { F0 uint8; F1 int16 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 int8; F1 uint; F2 uint64 }
type GenS7 struct { F0 uint32; F1 uint16; F2 int16 }
type GenS8 struct { F0 string; F1 uint32; F2 int32 }
type GenS9 struct { F0 uint }
type GenS10 struct { F0 string; F1 int }
type GenS11 struct { F0 int64; F1 int16; F2 uint32 }
type GenS12 struct { F0 uint32 }
type GenS13 struct { F0 uint64; F1 uint8; F2 int64 }

func main() {
    v1 := make(map[string]bool)
    var v2 uint64 = 453
    var v3 int = 4294967295
    var v4 bool = (true || (true && true))
    var v5 float64 = (1e10 - -2.25)
    var v6 int32 = 1
    if ((127 < 36) && v4) {
        fmt.Println(v4, v6, v3)
        switch "hello" {
        case "a":
            v4 = v4
        case "":
            if v4 {
                switch v3 {
                case 1:
                    fmt.Println(v6)
                }
                for fi7 := 0; fi7 < 19; fi7++ {
                    delete(v1, ("" + (("a" + "hello") + "")))
                    fmt.Println(v3)
                    fmt.Println(v3, v5)
                    _ = fi7
                }
            } else {
                switch "" {
                case "foo bar":
                    if v4 {
                        var v8 int16 = -1
                        _ = v8
                    }
                case "\n\t":
                    var v9 string = ("\n\t" + "foo bar")
                    _ = v9
                case "":
                    var v10 int64 = -9223372036854775808
                    _ = v10
                default:
                    v4 = (255 == 0)
                }
                var v11 int16 = -1
                _ = v11
            }
        }
    }
    if v4 {
        if (!(99 >= 75)) {
            fmt.Println(v4, v6, v5)
            v2 += v2
        } else {
            var v12 float32 = 1.5
            v3 += 2147483647
            _ = v12
        }
        switch "\n\t" {
        case "hello":
            v2 = v2
        case "\n\t":
            for fi13 := 0; fi13 < 15; fi13++ {
                switch ((("\n\t" + (("a" + "foo bar") + (("\n\t" + "") + "\n\t"))) + "a") + "a") {
                case "":
                    var v14 uint8 = 42
                    _ = v14
                }
                for fi15 := 0; fi15 < 9; fi15++ {
                    switch "hello" {
                    case "hello":
                        v16 := GenS1{F0: 42}
                        _ = v16
                    case "":
                        switch "" {
                        case "hello":
                            var v17 int32 = -547
                            _ = v17
                        }
                    }
                    _ = fi15
                }
                _ = fi13
            }
        case "a":
            v5 = (0.1 / -2.25)
        }
        v4 = (0 == 4294967295)
    }
    var v18 uint = 2
    fmt.Println(v18, v5)
    for fi19 := 0; fi19 < 11; fi19++ {
        _ = len(v1)
        v18 += (^(v18))
        _ = fi19
    }
    switch -1 {
    case 127:
        v6 = v6
    default:
        fmt.Println(v4, v6, v2)
    }
    if true {
        var v20 int64 = 42
        v6 = v6
        if ((v4 && (!((v20 & v20) != v20))) || v4) {
            switch (("hello" + "\n\t") + ("\n\t" + "hello")) {
            case "a":
                switch "a" {
                case "hello":
                    v2 -= v2
                }
            case "\n\t":
                if false {
                    v4 = ((!v4) && ("\n\t" >= ("a" + (("foo bar" + ((("foo bar" + "\n\t") + "hello") + "hello")) + ((((("a" + "\n\t") + ((("" + "foo bar") + "hello") + "\n\t")) + ("a" + "hello")) + "foo bar") + "hello")))))
                    switch 127 {
                    case 32767:
                        if (int(v18) >= v3) {
                            v21 := []int{int(v6), v3, v3}
                            for fi22 := 0; fi22 < 15; fi22++ {
                                switch "hello" {
                                case "hello":
                                    v1["a"] = (237 >= 1)
                                default:
                                    if (false && (((-2.25 * ((1.5 / (3.14159 + (-2.25 - 0.1))) - 1e10)) * float32(0.0)) > (1.5 / 1.5))) {
                                        if v4 {
                                            var v23 float32 = 1.5
                                            var v24 int32 = v6
                                            var v25 int8 = -1
                                            _ = v23
                                            _ = v24
                                            _ = v25
                                        }
                                        var v26 int16 = -32768
                                        _ = v26
                                    } else {
                                        if len(v21) >= 2 {
                                            v21 = v21[1:len(v21)]
                                        }
                                        var v27 uint32 = 4294967295
                                        var v28 int16 = 1
                                        _ = v27
                                        _ = v28
                                    }
                                }
                                for fi29 := 0; fi29 < 1; fi29++ {
                                    v30 := GenS2{F0: 42, F1: fi29}
                                    var v31 uint16 = 42
                                    fmt.Println(v18, v2)
                                    _ = fi29
                                    _ = v30
                                    _ = v31
                                }
                                _ = fi22
                            }
                            fmt.Println(v18, v20, v4)
                            _ = v21
                        } else {
                            v3 -= int(v6)
                            if v4 {
                                v20 = (v20 >> (((v18 % (((v18 + uint(v2))) | 1))) & 63))
                                var v32 uint = (v18 ^ (v18 ^ uint(v2)))
                                _ = v32
                            }
                            for fi33 := 0; fi33 < 15; fi33++ {
                                fmt.Println(v6)
                                _ = fi33
                            }
                        }
                    case 42:
                        var v34 int8 = 100
                        _ = v34
                    default:
                        v35 := []int{v3}
                        _ = v35
                    }
                    _ = len(v1)
                } else {
                    fmt.Println(v4, v3, v18)
                    if (790 > 313) {
                        v36 := []int{v3}
                        if true {
                            var v37 uint64 = 915
                            _ = v37
                        } else {
                            switch "\n\t" {
                            case "\n\t":
                                v38 := GenS1{F0: 18446744073709551615}
                                _ = v38
                            case "hello":
                                for fi39 := 0; fi39 < 9; fi39++ {
                                    v1[(("\n\t" + "a") + ("\n\t" + "\n\t"))] = (!(0 >= 127))
                                    var v40 int = int(v18)
                                    _ = fi39
                                    _ = v40
                                }
                            }
                            if true {
                                v3 -= 4294967295
                            }
                        }
                        _ = v36
                    } else {
                        v3 -= v3
                    }
                    switch (("" + "") + "\n\t") {
                    case "\n\t":
                        var v41 uint = 186
                        _ = v41
                    case "foo bar":
                        if ((!false) && false) {
                            fmt.Println(v5)
                            var v42 bool = (float64(float32(v5)) >= (v5 * v5))
                            switch ("foo bar" + "") {
                            case "hello":
                                v43 := []int{(v3 * v3), 436, int(v6)}
                                _ = v43
                            default:
                                var v44 int16 = int16(v6)
                                _ = v44
                            }
                            _ = v42
                        }
                    case "hello":
                        if v4 {
                            var v45 int32 = int32(v3)
                            var v46 int8 = 42
                            _ = v45
                            _ = v46
                        } else {
                            fmt.Println(v6, v3)
                            if v4 {
                                v47 := GenS2{F0: 123, F1: (v3 ^ v3)}
                                v5 -= v5
                                _ = v47
                            }
                        }
                    }
                }
            }
            if (-32768 <= 251) {
                switch ("a" + "foo bar") {
                case "a":
                    var v48 int = v3
                    _ = v48
                case "\n\t":
                    v6 = 2147483647
                case "":
                    var v49 int32 = (v6 | (v6 % ((-1) | 1)))
                    _ = v49
                }
                for fi50 := 0; fi50 < 11; fi50++ {
                    for fi51 := 0; fi51 < 5; fi51++ {
                        v52 := GenS2{F0: 123, F1: (fi50 * v3)}
                        v53 := GenS3{F0: 1, F1: (float32(v5) == 1e10), F2: 100}
                        var v54 uint = (v18 % (((v18 ^ 1)) | 1))
                        _ = fi51
                        _ = v52
                        _ = v53
                        _ = v54
                    }
                    _ = fi50
                }
            } else {
                for fi55 := 0; fi55 < 13; fi55++ {
                    var v56 bool = (799 >= uint32(v18))
                    var v57 uint8 = 27
                    fmt.Println(v4, fi55, v6)
                    _ = fi55
                    _ = v56
                    _ = v57
                }
            }
        }
        _ = v20
    } else {
        var v58 bool = ((!(((v18 != v18) || false) || ((float32(float64(0.0)) - 0.1) < 1e10))) || (35 == 32767))
        _ = v58
    }
    switch "foo bar" {
    case "":
        switch "foo bar" {
        case "foo bar":
            var v59 uint8 = 237
            _ = v59
        case "\n\t":
            var v60 int32 = -2147483648
            _ = v60
        case "":
            if v4 {
                if true {
                    for fi61 := 0; fi61 < 3; fi61++ {
                        for fi62 := 0; fi62 < 4; fi62++ {
                            v6 += -1
                            switch v3 {
                            case -811:
                                var v63 int8 = 127
                                v3 = (v3 << ((1) & 63))
                                fmt.Println(v63, v3, v3)
                                _ = v63
                            }
                            v3 = 32767
                            _ = fi62
                        }
                        v64 := GenS4{F0: 127, F1: -213}
                        v4 = v4
                        _ = fi61
                        _ = v64
                    }
                    v65 := GenS4{F0: 1, F1: 42}
                    if v4 {
                        if (v4 || false) {
                            var v66 uint = (v18 + uint(v2))
                            fmt.Println(v18, v5)
                            _ = v66
                        } else {
                            if ((v18 | (v18 << ((v18) & 63))) != v18) {
                                v3 = v3
                                for fi67 := 0; fi67 < 16; fi67++ {
                                    var v68 string = ("" + "")
                                    var v69 uint8 = 2
                                    var v70 bool = true
                                    _ = fi67
                                    _ = v68
                                    _ = v69
                                    _ = v70
                                }
                                for fi71 := 0; fi71 < 19; fi71++ {
                                    v72 := GenS5{F0: 9223372036854775807}
                                    _ = fi71
                                    _ = v72
                                }
                            } else {
                                switch ("\n\t" + "foo bar") {
                                case "\n\t":
                                    if false {
                                        var v73 int64 = 9223372036854775807
                                        var v74 bool = true
                                        _ = v73
                                        _ = v74
                                    } else {
                                        var v75 uint8 = 109
                                        fmt.Println(v2, v4, v6)
                                        _ = v75
                                    }
                                default:
                                    fmt.Println(v6, v6, v18)
                                }
                                for fi76 := 0; fi76 < 4; fi76++ {
                                    var v77 bool = (32767 < -456)
                                    v78 := GenS3{F0: 2, F1: (v18 >= 797), F2: -1}
                                    _ = fi76
                                    _ = v77
                                    _ = v78
                                }
                                v18 = 1
                            }
                            var v79 int = 4294967295
                            _ = v79
                        }
                    } else {
                        for fi80 := 0; fi80 < 19; fi80++ {
                            var v81 float32 = 0.0
                            _ = fi80
                            _ = v81
                        }
                        for fi82 := 0; fi82 < 15; fi82++ {
                            if (42 == 1) {
                                var v83 uint = (^(v18))
                                v84 := GenS6{F0: 14, F1: v83, F2: v2}
                                _ = v83
                                _ = v84
                            } else {
                                fmt.Println(v2)
                                for fi85 := 0; fi85 < 5; fi85++ {
                                    _ = len(v1)
                                    if (!((v18 >= (^(v18))) || (-1 >= 1))) {
                                        switch "a" {
                                        case "\n\t":
                                            v86 := GenS5{F0: -9223372036854775808}
                                            _ = v86
                                        default:
                                            delete(v1, "foo bar")
                                        }
                                        if v4 {
                                            v5 = (v5 * (v5 - -2.25))
                                        }
                                    } else {
                                        v5 -= v5
                                    }
                                    v87 := GenS2{F0: 1, F1: -46}
                                    _ = fi85
                                    _ = v87
                                }
                            }
                            var v88 int8 = -128
                            _ = fi82
                            _ = v88
                        }
                    }
                    _ = v65
                } else {
                    for fi89 := 0; fi89 < 19; fi89++ {
                        v4 = v4
                        for fi90 := 0; fi90 < 12; fi90++ {
                            switch v3 {
                            case 32:
                                v91 := GenS5{F0: 1}
                                _ = v91
                            }
                            v2 = (v2 & (v2 << (((v18 / (((v18 / (((v18 + (v18 + v18))) | 1))) | 1))) & 63)))
                            v6 -= v6
                            _ = fi90
                        }
                        v92 := GenS3{F0: 57, F1: (false || v4), F2: 100}
                        _ = fi89
                        _ = v92
                    }
                    var v93 string = ("foo bar" + "hello")
                    _ = v93
                }
                for fi94 := 0; fi94 < 6; fi94++ {
                    var v95 uint16 = 32767
                    _ = fi94
                    _ = v95
                }
                v96 := GenS7{F0: 42, F1: 251, F2: 0}
                _ = v96
            }
        }
    case "hello":
        v6 = -2147483648
    default:
        v97 := v1[("" + ("foo bar" + "foo bar"))]
        _ = v97
    }
    fmt.Println(v18)
    for fi98 := 0; fi98 < 15; fi98++ {
        for fi99 := 0; fi99 < 9; fi99++ {
            v100 := []int{v3}
            v6 -= v6
            if (127 < 95) {
                if false {
                    var v101 int8 = int8(v2)
                    var v102 uint8 = 38
                    _ = v101
                    _ = v102
                } else {
                    v4 = (!(1e10 > float64(0.0)))
                    if (v4 && (((v4 || (v5 == v5)) && v4) || v4)) {
                        var v103 bool = false
                        _ = v103
                    }
                }
                _ = len(v100)
            }
            var v104 string = "\n\t"
            switch 42 {
            case -2147483648:
                v18 += (v18 + 2)
            case 385:
                switch (fi99 << (((v18 + (v18 % ((v18) | 1)))) & 63)) {
                case 1:
                    var v105 string = "hello"
                    _ = v105
                default:
                    _ = len(v100)
                }
            case 42:
                v106 := GenS8{F0: "\n\t", F1: 498, F2: v6}
                _ = v106
            }
            _ = len(v100)
            v2 -= (v2 ^ (v2 << ((v18) & 63)))
            v2 += 18446744073709551615
            v107 := GenS9{F0: v18}
            var v108 int = v3
            _ = fi99
            _ = v100
            _ = v104
            _ = v107
            _ = v108
        }
        v109 := GenS10{F0: ("" + ("hello" + "")), F1: 128}
        _ = fi98
        _ = v109
    }
    switch "a" {
    case "\n\t":
        v4 = false
    }
    v3 += (v3 * v3)
    if v4 {
        var v110 uint16 = 65535
        _ = v110
    }
    switch "" {
    case "":
        for fi111 := 0; fi111 < 18; fi111++ {
            v6 += -1
            v112 := []int{-2147483648}
            v4 = (255 == 42)
            _ = fi111
            _ = v112
        }
    }
    switch (v3 ^ -985) {
    case 32768:
        var v113 float32 = (1.5 / ((((0.1 + 1e10) * ((0.0 / float32(v5)) * 1.5)) - float32(v5)) * ((float32(float64(0.0)) + float32(-2.25)) / 1.5)))
        _ = v113
    case -1:
        v114 := GenS11{F0: -306, F1: 999, F2: 2}
        _ = v114
    }
    if true {
        if v4 {
            v115 := GenS12{F0: 2147483647}
            v6 -= (-(v6))
            v116 := GenS13{F0: v2, F1: 72, F2: 286}
            _ = v115
            _ = v116
        } else {
            v117 := []int{614, 2147483647, -1}
            var v118 uint16 = 2
            _ = v117
            _ = v118
        }
    } else {
        for fi119 := 0; fi119 < 17; fi119++ {
            if (v4 || true) {
                var v120 int32 = (v6 * v6)
                for fi121 := 0; fi121 < 19; fi121++ {
                    var v122 int64 = 1
                    v123 := GenS10{F0: "hello", F1: v3}
                    v5 = 3.14159
                    switch v3 {
                    case 127:
                        var v124 string = "a"
                        _ = v124
                    case -2147483648:
                        var v125 int = v3
                        _ = v125
                    default:
                        if (1e10 > -2.25) {
                            var v126 int16 = 32767
                            v127 := []int{(-(fi119)), fi119, (fi121 % ((2147483647) | 1)), (v3 * fi121)}
                            _ = v126
                            _ = v127
                        }
                    }
                    _ = fi121
                    _ = v122
                    _ = v123
                }
                _ = v120
            }
            fmt.Println(v6)
            for fi128 := 0; fi128 < 20; fi128++ {
                switch ("hello" + "hello") {
                case "a":
                    v18 -= v18
                case "\n\t":
                    if v4 {
                        v129 := GenS13{F0: v2, F1: 255, F2: 1}
                        _ = v129
                    }
                }
                _ = fi128
            }
            _ = fi119
        }
        if (!false) {
            v6 += 1
            for fi130 := 0; fi130 < 18; fi130++ {
                var v131 uint = v18
                _ = fi130
                _ = v131
            }
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
