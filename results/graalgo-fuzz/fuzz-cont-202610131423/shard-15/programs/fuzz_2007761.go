package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8; F2 int16 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int8; F1 int64 }
type GenS4 struct { F0 bool; F1 int16 }
type GenS5 struct { F0 int64; F1 uint32 }
type GenS6 struct { F0 uint64; F1 int32; F2 uint16 }
type GenS7 struct { F0 uint8 }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 float32 = (-2.25 / (((0.0 - 3.14159) / (1.5 / float32((1e10 - float64((0.0 + 1.5)))))) + -2.25))
    v3 := make(map[int]string)
    var v4 bool = (false && false)
    var v5 uint16 = 1
    var v6 int16 = -32768
    switch ("\n\t" + (((("a" + "foo bar") + "foo bar") + "foo bar") + "a")) {
    case "\n\t":
        v6 = 1000
    case "":
        var v7 int64 = 1
        _ = v7
    case "a":
        fmt.Println(v4, v4)
    }
    if (!v4) {
        if v4 {
            switch 127 {
            case -2147483648:
                v8 := []int{9223372036854775807, 121, -9223372036854775808}
                _ = v8
            case 1:
                var v9 int16 = (v6 ^ v6)
                _ = v9
            default:
                v10 := GenS1{F0: (^(v5)), F1: 2, F2: v6}
                _ = v10
            }
            delete(v1, 65535)
            v2 = 1.5
        }
        switch ("a" + "\n\t") {
        case "foo bar":
            var v11 float32 = v2
            _ = v11
        }
    }
    v12 := v1[669]
    v2 -= 1.5
    v13 := []int{}
    var v14 int = -673
    if (!((!(1 != (^(v14)))) || v4)) {
        for fi15 := 0; fi15 < 13; fi15++ {
            for fi16 := 0; fi16 < 2; fi16++ {
                v14 = fi15
                if false {
                    v5 = 769
                    var v17 uint8 = 127
                    v18 := GenS2{F0: -1}
                    _ = v17
                    _ = v18
                } else {
                    if (!(127 > 42)) {
                        var v19 int64 = 816
                        v20 := GenS2{F0: 2147483647}
                        v14 += fi15
                        _ = v19
                        _ = v20
                    }
                }
                _ = fi16
            }
            _ = fi15
        }
        v6 -= (v6 * v6)
    } else {
        var v21 float32 = 0.0
        v12 = (v12 - v12)
        _ = v21
    }
    for fi22 := 0; fi22 < 19; fi22++ {
        _ = len(v13)
        _ = fi22
    }
    if (!false) {
        switch v14 {
        case 4294967295:
            v12 = 2
        case 32768:
            for fi23 := 0; fi23 < 3; fi23++ {
                v24 := GenS1{F0: 0, F1: 179, F2: -32768}
                var v25 bool = v4
                _ = fi23
                _ = v24
                _ = v25
            }
        case -2147483648:
            fmt.Println(v12, v2)
        default:
            if v4 {
                var v26 uint64 = 468
                v12 -= (v12 - v12)
                v27 := v3[(v14 + v14)]
                _ = v26
                _ = v27
            }
        }
    } else {
        if v4 {
            var v28 int = (v14 ^ (v14 >> ((853) & 63)))
            v29 := GenS3{F0: 1, F1: 621}
            _ = v28
            _ = v29
        } else {
            var v30 string = "hello"
            v31 := GenS2{F0: 1}
            _ = v30
            _ = v31
        }
        fmt.Println(v5, v12)
        var v32 string = "foo bar"
        _ = v32
    }
    if false {
        var v33 float32 = ((v2 / ((v2 + v2) / (v2 - 0.1))) / float32(3.14159))
        _ = v33
    }
    if (!((615 < -653) || (127 >= 0))) {
        for fi34 := 0; fi34 < 19; fi34++ {
            var v35 int32 = 2147483647
            if len(v13) > 0 {
                v13[0] = 42
            }
            v35 = int32(fi34)
            _ = fi34
            _ = v35
        }
        v36 := GenS1{F0: 0, F1: 128, F2: 0}
        _ = v36
    } else {
        for fi37 := 0; fi37 < 5; fi37++ {
            v38 := GenS4{F0: false, F1: v6}
            fmt.Println(fi37)
            _ = fi37
            _ = v38
        }
    }
    _ = len(v13)
    if v4 {
        var v39 uint = 255
        _ = v39
    } else {
        if (((-2.25 + (0.1 + 1.5)) + float64(v2)) == 0.0) {
            switch (("" + ("a" + "hello")) + "") {
            case "a":
                if v4 {
                    if (736 < -537) {
                        if v4 {
                            v4 = false
                            var v40 uint32 = (v12 + v12)
                            switch ("hello" + "foo bar") {
                            case "a":
                                v4 = v4
                            case "\n\t":
                                if false {
                                    delete(v1, 42)
                                    var v41 string = "\n\t"
                                    v42 := GenS5{F0: 0, F1: v12}
                                    _ = v41
                                    _ = v42
                                } else {
                                    v43 := GenS2{F0: int32(v5)}
                                    for fi44 := 0; fi44 < 14; fi44++ {
                                        v45 := v1[v12]
                                        for fi46 := 0; fi46 < 18; fi46++ {
                                            v47 := GenS4{F0: true, F1: 0}
                                            if (v5 <= (v5 << ((18446744073709551615) & 63))) {
                                                v48 := GenS6{F0: 42, F1: -1, F2: v5}
                                                if v4 {
                                                    v12 += v40
                                                    switch fi46 {
                                                    case 0:
                                                        if false {
                                                            var v49 int32 = 42
                                                            v50 := GenS4{F0: (v4 && (0.1 >= 1.5)), F1: v6}
                                                            _ = v49
                                                            _ = v50
                                                        } else {
                                                            v51 := v1[v45]
                                                            v52 := GenS4{F0: (!(128 > 128)), F1: (v6 << ((941) & 63))}
                                                            _ = v51
                                                            _ = v52
                                                        }
                                                    }
                                                    v12 = uint32(v5)
                                                } else {
                                                    if true {
                                                        if (44 == -116) {
                                                            var v53 uint32 = 2147483648
                                                            fmt.Println(v5, v14, v12)
                                                            v4 = v4
                                                            _ = v53
                                                        }
                                                        if v4 {
                                                            switch (fi46 | v14) {
                                                            case 250:
                                                                v54 := GenS3{F0: 100, F1: 42}
                                                                _ = v54
                                                            case -296:
                                                                delete(v3, 895)
                                                            }
                                                            fmt.Println(v4)
                                                        } else {
                                                            fmt.Println(v2, fi44)
                                                            if ((127 < 16) && true) {
                                                                v40 += (v45 >> ((4294967295) & 63))
                                                                var v55 int32 = -125
                                                                _ = v55
                                                            }
                                                        }
                                                        v56 := GenS5{F0: 909, F1: (v45 & (v12 + (v12 % ((2147483647) | 1))))}
                                                        _ = v56
                                                    }
                                                    switch "a" {
                                                    case "a":
                                                        fmt.Println(v14)
                                                    case "\n\t":
                                                        switch ("foo bar" + "foo bar") {
                                                        case "":
                                                            v40 += v12
                                                        default:
                                                            v1[v40] = (v40 - v40)
                                                        }
                                                    default:
                                                        fmt.Println(v5, fi46, v40)
                                                    }
                                                }
                                                _ = v48
                                            } else {
                                                switch "\n\t" {
                                                case "\n\t":
                                                    if len(v13) > 0 {
                                                        v13[0] = (fi46 % ((v14) | 1))
                                                    }
                                                case "":
                                                    fmt.Println(v5, v2)
                                                case "foo bar":
                                                    fmt.Println(v2, v5)
                                                }
                                            }
                                            if (-1 != 936) {
                                                v57 := GenS4{F0: true, F1: (v6 + 1)}
                                                v13 = append(v13, -45)
                                                _ = len(v3)
                                                _ = v57
                                            } else {
                                                v58 := v1[65535]
                                                if (!v4) {
                                                    fmt.Println(v45, v4)
                                                    v59 := GenS6{F0: 409, F1: 42, F2: 256}
                                                    _ = v59
                                                } else {
                                                    var v60 int32 = -124
                                                    v2 += v2
                                                    _ = v60
                                                }
                                                var v61 uint64 = 18446744073709551615
                                                _ = v58
                                                _ = v61
                                            }
                                            _ = fi46
                                            _ = v47
                                        }
                                        _ = fi44
                                        _ = v45
                                    }
                                    var v62 uint = 700
                                    _ = v43
                                    _ = v62
                                }
                            }
                            _ = v40
                        }
                        switch v14 {
                        case 315:
                            switch "hello" {
                            case "foo bar":
                                v6 = 42
                            case "hello":
                                var v63 int16 = v6
                                _ = v63
                            default:
                                var v64 int32 = int32(v12)
                                _ = v64
                            }
                        default:
                            v4 = v4
                        }
                    } else {
                        if len(v13) >= 2 {
                            v13 = v13[1:len(v13)]
                        }
                        if v4 {
                            switch "" {
                            case "\n\t":
                                for fi65 := 0; fi65 < 2; fi65++ {
                                    v4 = (!v4)
                                    if (0 != 108) {
                                        v1[(v12 % ((2147483647) | 1))] = v12
                                    }
                                    _ = len(v1)
                                    _ = fi65
                                }
                            case "a":
                                v4 = false
                            case "hello":
                                if (!(!(!(42 != 9223372036854775807)))) {
                                    for fi66 := 0; fi66 < 15; fi66++ {
                                        var v67 uint16 = v5
                                        fmt.Println(v4, v4)
                                        _ = fi66
                                        _ = v67
                                    }
                                    switch (v14 - 339) {
                                    case 32768:
                                        for fi68 := 0; fi68 < 19; fi68++ {
                                            var v69 int = 9223372036854775807
                                            _ = fi68
                                            _ = v69
                                        }
                                    case -1:
                                        var v70 int = v14
                                        _ = v70
                                    case 128:
                                        switch "" {
                                        case "a":
                                            var v71 int8 = 42
                                            _ = v71
                                        case "":
                                            var v72 float64 = 0.1
                                            _ = v72
                                        case "\n\t":
                                            for fi73 := 0; fi73 < 10; fi73++ {
                                                if v4 {
                                                    switch fi73 {
                                                    case -1:
                                                        for fi74 := 0; fi74 < 16; fi74++ {
                                                            v3[2147483647] = "\n\t"
                                                            _ = fi74
                                                        }
                                                    default:
                                                        v75 := GenS5{F0: 1, F1: v12}
                                                        _ = v75
                                                    }
                                                } else {
                                                    for fi76 := 0; fi76 < 3; fi76++ {
                                                        var v77 int8 = 0
                                                        v4 = (2 <= 255)
                                                        fmt.Println(v14, v14)
                                                        _ = fi76
                                                        _ = v77
                                                    }
                                                }
                                                _ = fi73
                                            }
                                        }
                                    }
                                    v78 := v1[(v12 ^ (^(v12)))]
                                    _ = v78
                                }
                            }
                            switch 42 {
                            case -9223372036854775808:
                                v79 := GenS2{F0: 51}
                                _ = v79
                            case 32768:
                                switch v14 {
                                case 787:
                                    for fi80 := 0; fi80 < 14; fi80++ {
                                        if (!(v4 && true)) {
                                            v4 = true
                                        } else {
                                            var v81 uint64 = 42
                                            _ = v81
                                        }
                                        var v82 float64 = 3.14159
                                        var v83 bool = (!(-208 > int64(v14)))
                                        _ = fi80
                                        _ = v82
                                        _ = v83
                                    }
                                case -2147483648:
                                    if ((v12 / ((0) | 1)) > v12) {
                                        var v84 string = ""
                                        v85 := GenS5{F0: -526, F1: (v12 / (((^(v12))) | 1))}
                                        fmt.Println(v14, v12, v84)
                                        _ = v84
                                        _ = v85
                                    }
                                case 2147483647:
                                    if true {
                                        for fi86 := 0; fi86 < 17; fi86++ {
                                            var v87 int32 = 630
                                            for fi88 := 0; fi88 < 17; fi88++ {
                                                var v89 int = (fi88 / (((v14 - v14)) | 1))
                                                if true {
                                                    v14 = (-(fi88))
                                                    v90 := GenS5{F0: -1, F1: 1}
                                                    _ = v90
                                                } else {
                                                    var v91 uint16 = (v5 & (v5 % ((v5) | 1)))
                                                    var v92 string = "hello"
                                                    v93 := GenS2{F0: 42}
                                                    switch v92 {
                                                    case "\n\t":
                                                        switch v89 {
                                                        case -51:
                                                            fmt.Println(v6, v91)
                                                        case 255:
                                                            v2 = float32(float64((float32(0.1) / v2)))
                                                            if (!v4) {
                                                                v94 := GenS1{F0: 65535, F1: 148, F2: (v6 - (^(v6)))}
                                                                _ = v94
                                                            } else {
                                                                var v95 float32 = (0.0 + v2)
                                                                _ = v95
                                                            }
                                                            fmt.Println(v12, v4)
                                                        case -581:
                                                            v6 += (v6 * v6)
                                                        }
                                                    default:
                                                        v4 = false
                                                    }
                                                    _ = v91
                                                    _ = v92
                                                    _ = v93
                                                }
                                                fmt.Println(v14, v5)
                                                _ = fi88
                                                _ = v89
                                            }
                                            v96 := GenS3{F0: 98, F1: -1}
                                            _ = fi86
                                            _ = v87
                                            _ = v96
                                        }
                                        if (v4 || false) {
                                            if v4 {
                                                fmt.Println(v12)
                                                for fi97 := 0; fi97 < 15; fi97++ {
                                                    fmt.Println(fi97, fi97)
                                                    fmt.Println(fi97)
                                                    _ = fi97
                                                }
                                            } else {
                                                v5 = v5
                                            }
                                            v98 := GenS6{F0: 42, F1: 1, F2: v5}
                                            _ = v98
                                        } else {
                                            var v99 int = 9223372036854775807
                                            _ = v99
                                        }
                                        fmt.Println(v5, v5, v2)
                                    } else {
                                        v100 := GenS2{F0: 1}
                                        for fi101 := 0; fi101 < 6; fi101++ {
                                            v12 = 942
                                            _ = fi101
                                        }
                                        _ = v100
                                    }
                                }
                            case 1:
                                v4 = true
                            }
                            if v4 {
                                v102 := v1[(^(v12))]
                                var v103 int32 = 2147483647
                                fmt.Println(v4)
                                _ = v102
                                _ = v103
                            }
                        } else {
                            if v4 {
                                switch (("\n\t" + "a") + "foo bar") {
                                case "\n\t":
                                    for fi104 := 0; fi104 < 8; fi104++ {
                                        switch "\n\t" {
                                        case "":
                                            for fi105 := 0; fi105 < 8; fi105++ {
                                                fmt.Println(fi105)
                                                if len(v13) > 0 {
                                                    v13[0] = (fi105 + v14)
                                                }
                                                _ = fi105
                                            }
                                        }
                                        v5 -= 795
                                        _ = fi104
                                    }
                                }
                                switch (((("foo bar" + "a") + "a") + ("foo bar" + (("" + "a") + ""))) + "foo bar") {
                                case "a":
                                    v12 = (^(v12))
                                case "foo bar":
                                    v2 = 1.5
                                }
                            }
                        }
                        var v106 uint32 = 2
                        _ = v106
                    }
                    var v107 float32 = float32((0.1 * 0.1))
                    for fi108 := 0; fi108 < 17; fi108++ {
                        v12 = (v12 - (v12 - v12))
                        v14 = -9223372036854775808
                        _ = fi108
                    }
                    _ = v107
                }
            default:
                var v109 uint8 = 0
                _ = v109
            }
            switch (v14 | v14) {
            case -2147483648:
                for fi110 := 0; fi110 < 9; fi110++ {
                    var v111 uint8 = 42
                    var v112 int8 = 105
                    fmt.Println(v112)
                    _ = fi110
                    _ = v111
                    _ = v112
                }
            default:
                v12 = (^(v12))
            }
            switch "foo bar" {
            case "\n\t":
                if ((v5 | v5) == (v5 * (v5 ^ v5))) {
                    delete(v3, (v14 + int(v5)))
                }
            case "hello":
                v2 = v2
            case "":
                switch "foo bar" {
                case "\n\t":
                    if false {
                        v3[v14] = ("a" + "\n\t")
                    } else {
                        var v113 uint32 = v12
                        _ = v113
                    }
                default:
                    var v114 int8 = int8(v5)
                    _ = v114
                }
            }
        }
    }
    _ = len(v3)
    if (false || ((true && (42 <= 0)) && v4)) {
        fmt.Println(v4, v5, v12)
        if (1 > 42) {
            for fi115 := 0; fi115 < 18; fi115++ {
                var v116 bool = v4
                v117 := GenS7{F0: 128}
                if len(v13) > 0 {
                    v13[0] = -2147483648
                }
                _ = fi115
                _ = v116
                _ = v117
            }
            if v4 {
                delete(v1, (v12 & (^(v12))))
                v118 := GenS3{F0: -24, F1: -1}
                _ = v118
            }
            v2 += v2
        } else {
            v5 += 542
            for fi119 := 0; fi119 < 17; fi119++ {
                var v120 string = ("hello" + "")
                fmt.Println(fi119, v12, v4)
                _ = fi119
                _ = v120
            }
            switch 128 {
            case 909:
                v2 -= (((v2 / v2) * (v2 - 1e10)) * (v2 - v2))
            case -881:
                v1[v12] = 65535
            case -9:
                v14 = int(v12)
            default:
                switch (v14 | v14) {
                case -300:
                    _ = len(v3)
                case 42:
                    v14 -= -1
                }
            }
        }
        fmt.Println(v6)
    }
    var v121 uint64 = 42
    var v122 uint32 = 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v121:", v121)
    fmt.Println("v122:", v122)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v121
    _ = v122
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
