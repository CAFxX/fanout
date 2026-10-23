package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int64; F1 string }
type GenS3 struct { F0 uint8; F1 int8; F2 uint64 }
type GenS4 struct { F0 uint64; F1 int8 }
type GenS5 struct { F0 int8; F1 int64 }
type GenS6 struct { F0 int64; F1 bool }
type GenS7 struct { F0 int64; F1 string }
type GenS8 struct { F0 bool; F1 string; F2 string }
type GenS9 struct { F0 uint16 }
type GenS10 struct { F0 uint }
type GenS11 struct { F0 string; F1 int16 }
type GenS12 struct { F0 int16; F1 int; F2 uint32 }
type GenS13 struct { F0 int16; F1 uint8; F2 uint16 }
type GenS14 struct { F0 uint32 }
type GenS15 struct { F0 int; F1 string; F2 uint32 }
type GenS16 struct { F0 int8; F1 uint64 }
type GenS17 struct { F0 uint16; F1 uint }
type GenS18 struct { F0 uint8; F1 int }
type GenS19 struct { F0 uint32; F1 string }
type GenS20 struct { F0 uint; F1 uint; F2 string }
type GenS21 struct { F0 int16; F1 int32; F2 int16 }
type GenS22 struct { F0 int }

func main() {
    var v1 string = "a"
    v2 := make(map[int]bool)
    var v3 uint = 2
    var v4 uint32 = 524
    v5 := []uint32{}
    for fi6 := 0; fi6 < 13; fi6++ {
        v1 = v1
        v3 = v3
        switch fi6 {
        case 42:
            fmt.Println(fi6, fi6, fi6)
        case -313:
            if false {
                v3 = (v3 << ((v3) & 63))
                for fi7 := 0; fi7 < 6; fi7++ {
                    var v8 int = fi7
                    _ = fi7
                    _ = v8
                }
            }
        }
        _ = fi6
    }
    for fi9 := 0; fi9 < 17; fi9++ {
        for fi10 := 0; fi10 < 7; fi10++ {
            v11 := GenS1{F0: int32(fi10)}
            fmt.Println(fi9, v3)
            v1 = ("a" + "\n\t")
            if ((!false) && true) {
                v4 = 42
            }
            if ((!(((v4 & (v4 * v4)) >= (v4 ^ v4)) || (true && (-2.25 <= 1e10)))) || (true && true)) {
                v3 -= (^(v3))
                switch ((v1 + v1) + (("\n\t" + "") + "\n\t")) {
                case "":
                    var v12 float64 = (1.5 * (((float64(((0.0 / (0.1 - 3.14159)) * (1.5 + ((((((-2.25 * 1.5) - 0.0) + 1e10) + 0.0) + 1e10) * (0.0 - 0.1))))) * 3.14159) / float64(3.14159)) / (0.1 + float64(3.14159))))
                    _ = v12
                case "hello":
                    v3 += uint(v4)
                }
                v2[(fi9 * (fi10 ^ (fi10 - fi9)))] = (!false)
            }
            v1 = (v1 + (v1 + "a"))
            v1 = (v1 + "")
            v13 := GenS2{F0: -9223372036854775808, F1: "hello"}
            _ = fi10
            _ = v11
            _ = v13
        }
        fmt.Println(v1, fi9, v4)
        fmt.Println(v1, v3, v4)
        _ = fi9
    }
    if (!(1 >= 105)) {
        var v14 int16 = -323
        if (false && false) {
            switch -2147483648 {
            case 32768:
                switch v1 {
                case "a":
                    switch v1 {
                    case "a":
                        switch 42 {
                        case -1:
                            v3 += v3
                        case 2147483647:
                            var v15 bool = false
                            _ = v15
                        default:
                            for fi16 := 0; fi16 < 10; fi16++ {
                                v3 -= (v3 * v3)
                                fmt.Println(v14, v1, v1)
                                v3 += v3
                                _ = fi16
                            }
                        }
                    case "foo bar":
                        for fi17 := 0; fi17 < 19; fi17++ {
                            switch "a" {
                            case "\n\t":
                                v18 := GenS1{F0: -144}
                                _ = v18
                            case "foo bar":
                                v4 += (v4 * 65535)
                            case "":
                                switch (fi17 | (fi17 >> ((v3) & 63))) {
                                case -1:
                                    if (-634 > 679) {
                                        if false {
                                            fmt.Println(v4, v1, v1)
                                        } else {
                                            v19 := GenS1{F0: -1}
                                            v1 = v1
                                            _ = v19
                                        }
                                    } else {
                                        v20 := GenS1{F0: -935}
                                        _ = v20
                                    }
                                }
                            default:
                                var v21 uint64 = 0
                                _ = v21
                            }
                            v14 = -234
                            _ = fi17
                        }
                    case "\n\t":
                        var v22 string = v1
                        _ = v22
                    }
                case "\n\t":
                    var v23 string = "foo bar"
                    _ = v23
                case "":
                    var v24 uint16 = 255
                    _ = v24
                }
            case -390:
                if true {
                    if ((^(v4)) < (v4 + 548)) {
                        var v25 int8 = 100
                        if true {
                            fmt.Println(v4)
                            for fi26 := 0; fi26 < 12; fi26++ {
                                if true {
                                    v27 := GenS3{F0: 128, F1: 42, F2: 111}
                                    _ = v27
                                }
                                v3 = v3
                                _ = fi26
                            }
                        }
                        v2[128] = true
                        _ = v25
                    }
                    var v28 float32 = 1.5
                    v3 += v3
                    _ = v28
                } else {
                    if (!(false || false)) {
                        var v29 uint32 = (^(v4))
                        var v30 float64 = -2.25
                        if (42 >= 1) {
                            var v31 float32 = (0.0 + (1.5 + 1e10))
                            _ = v31
                        }
                        _ = v29
                        _ = v30
                    } else {
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        var v32 uint = (v3 % ((v3) | 1))
                        _ = v32
                    }
                }
            default:
                if (v1 == v1) {
                    if true {
                        v33 := GenS4{F0: 834, F1: -83}
                        _ = v33
                    } else {
                        var v34 uint64 = 1
                        for fi35 := 0; fi35 < 6; fi35++ {
                            var v36 int16 = v14
                            _ = fi35
                            _ = v36
                        }
                        for fi37 := 0; fi37 < 17; fi37++ {
                            v38 := GenS2{F0: 42, F1: v1}
                            _ = len(v2)
                            _ = fi37
                            _ = v38
                        }
                        _ = v34
                    }
                    v4 = v4
                }
            }
            if len(v5) > 0 {
                v5[0] = 429
            }
        } else {
            switch "foo bar" {
            case "":
                for fi39 := 0; fi39 < 12; fi39++ {
                    if true {
                        if false {
                            if true {
                                v40 := GenS5{F0: 0, F1: 42}
                                var v41 uint = v3
                                v3 -= (^(v41))
                                _ = v40
                                _ = v41
                            } else {
                                v2[fi39] = true
                                if (!false) {
                                    delete(v2, fi39)
                                    v14 -= v14
                                }
                            }
                            if false {
                                for fi42 := 0; fi42 < 2; fi42++ {
                                    if (v4 == 2) {
                                        var v43 int16 = (v14 * v14)
                                        _ = v43
                                    }
                                    _ = fi42
                                }
                                _ = len(v2)
                                var v44 uint16 = 65535
                                _ = v44
                            } else {
                                v45 := v2[(fi39 & fi39)]
                                var v46 float64 = 0.0
                                _ = v45
                                _ = v46
                            }
                        } else {
                            if (-2.25 == 1e10) {
                                if ((^(v4)) == 2147483648) {
                                    if true {
                                        v47 := GenS6{F0: -1, F1: (!false)}
                                        v4 -= v4
                                        _ = v47
                                    } else {
                                        v3 = (v3 % (((^(v3))) | 1))
                                        v48 := GenS2{F0: -712, F1: v1}
                                        _ = v48
                                    }
                                    if ((fi39 * (fi39 * fi39)) != (fi39 & 32767)) {
                                        if false {
                                            for fi49 := 0; fi49 < 16; fi49++ {
                                                var v50 int16 = v14
                                                v14 -= v14
                                                _ = fi49
                                                _ = v50
                                            }
                                        }
                                    } else {
                                        v1 = "hello"
                                        fmt.Println(fi39, v14)
                                    }
                                    for fi51 := 0; fi51 < 12; fi51++ {
                                        var v52 string = ""
                                        v14 = v14
                                        _ = fi51
                                        _ = v52
                                    }
                                } else {
                                    switch fi39 {
                                    case 719:
                                        switch v1 {
                                        case "a":
                                            switch fi39 {
                                            case -9223372036854775808:
                                                for fi53 := 0; fi53 < 1; fi53++ {
                                                    v54 := GenS7{F0: -28, F1: v1}
                                                    v14 -= v14
                                                    v14 = 32767
                                                    var v55 int8 = -128
                                                    _ = fi53
                                                    _ = v54
                                                    _ = v55
                                                }
                                            case 9223372036854775807:
                                                v3 = v3
                                            case 2147483647:
                                                v56 := GenS6{F0: -9223372036854775808, F1: false}
                                                _ = v56
                                            }
                                        case "\n\t":
                                            var v57 int16 = v14
                                            _ = v57
                                        case "hello":
                                            v14 = v14
                                        default:
                                            v3 += v3
                                        }
                                    case -9223372036854775808:
                                        var v58 uint8 = 127
                                        _ = v58
                                    }
                                }
                                v1 = v1
                                if (127 != 128) {
                                    switch v1 {
                                    case "foo bar":
                                        _ = len(v5)
                                    case "hello":
                                        if (!false) {
                                            if true {
                                                if true {
                                                    v59 := GenS7{F0: 1, F1: ""}
                                                    switch ("hello" + v1) {
                                                    case "":
                                                        fmt.Println(v1, v1, fi39)
                                                    case "hello":
                                                        for fi60 := 0; fi60 < 16; fi60++ {
                                                            v1 = "\n\t"
                                                            v14 -= v14
                                                            _ = fi60
                                                        }
                                                    case "\n\t":
                                                        var v61 uint32 = (v4 ^ v4)
                                                        _ = v61
                                                    default:
                                                        if (0.0 <= (0.1 + float64(3.14159))) {
                                                            var v62 uint = v3
                                                            v63 := GenS4{F0: 2, F1: 127}
                                                            var v64 int8 = 31
                                                            _ = v62
                                                            _ = v63
                                                            _ = v64
                                                        } else {
                                                            var v65 uint = v3
                                                            if len(v5) > 0 {
                                                                v5[0] = (v4 | v4)
                                                            }
                                                            _ = v65
                                                        }
                                                    }
                                                    v66 := GenS8{F0: (0 > 1), F1: v1, F2: ("\n\t" + v1)}
                                                    _ = v59
                                                    _ = v66
                                                }
                                                switch (fi39 % (((fi39 >> (((v3 ^ (v3 << (((v3 << ((912) & 63))) & 63)))) & 63))) | 1)) {
                                                case 818:
                                                    var v67 float64 = (1e10 / 3.14159)
                                                    _ = v67
                                                case 0:
                                                    var v68 uint16 = 159
                                                    _ = v68
                                                default:
                                                    v3 = 1
                                                }
                                                switch v1 {
                                                case "a":
                                                    var v69 uint32 = (v4 | v4)
                                                    _ = v69
                                                case "\n\t":
                                                    v70 := v2[(fi39 - 32767)]
                                                    _ = v70
                                                case "foo bar":
                                                    _ = len(v5)
                                                }
                                            } else {
                                                v71 := GenS2{F0: -1, F1: v1}
                                                _ = v71
                                            }
                                            v3 += v3
                                        }
                                    }
                                    var v72 uint = (v3 / ((2) | 1))
                                    switch (fi39 << (((v3 * (v3 >> ((uint(v4)) & 63)))) & 63)) {
                                    case -1:
                                        fmt.Println(v4, fi39, v14)
                                    case 2147483647:
                                        for fi73 := 0; fi73 < 19; fi73++ {
                                            var v74 int64 = 9223372036854775807
                                            var v75 uint8 = 2
                                            _ = fi73
                                            _ = v74
                                            _ = v75
                                        }
                                    case 127:
                                        for fi76 := 0; fi76 < 7; fi76++ {
                                            if (100 > 36) {
                                                v2[(fi76 ^ 32767)] = (v14 != (^(v14)))
                                            }
                                            switch fi76 {
                                            case 1:
                                                v3 += 0
                                            case 153:
                                                switch v1 {
                                                case "a":
                                                    var v77 uint8 = 42
                                                    _ = v77
                                                }
                                            case 909:
                                                var v78 float64 = 1.5
                                                if (v4 >= v4) {
                                                    switch v1 {
                                                    case "\n\t":
                                                        delete(v2, (-(fi76)))
                                                    default:
                                                        var v79 int = (fi39 - fi76)
                                                        _ = v79
                                                    }
                                                }
                                                if len(v5) > 0 {
                                                    v5[0] = (v4 % (((v4 & (v4 / (((v4 >> (((v72 % ((255) | 1))) & 63))) | 1)))) | 1))
                                                }
                                                var v80 int8 = 100
                                                _ = v78
                                                _ = v80
                                            default:
                                                v3 = v3
                                            }
                                            var v81 uint = 39
                                            _ = fi76
                                            _ = v81
                                        }
                                    }
                                    _ = v72
                                }
                            }
                            v4 += (v4 + v4)
                            v4 = 0
                        }
                    } else {
                        v82 := GenS3{F0: 127, F1: -101, F2: 659}
                        fmt.Println(v4, v4, v1)
                        _ = v82
                    }
                    var v83 int8 = 42
                    _ = fi39
                    _ = v83
                }
            case "hello":
                switch ("foo bar" + (v1 + "\n\t")) {
                case "\n\t":
                    v1 = v1
                case "a":
                    v1 = (v1 + "foo bar")
                case "":
                    switch v1 {
                    case "a":
                        switch v1 {
                        case "foo bar":
                            var v84 float32 = ((-2.25 - ((-2.25 / 1.0) * ((-2.25 + ((3.14159 * ((float32(((0.0 * float64((0.1 / float32(0.1)))) / (-2.25 + 0.1))) + (0.0 - (-2.25 + (0.1 + 0.0)))) * -2.25)) * -2.25)) - 1.5))) / 1e10)
                            _ = v84
                        case "":
                            if (float32(3.14159) > -2.25) {
                                if true {
                                    v85 := GenS1{F0: 0}
                                    for fi86 := 0; fi86 < 1; fi86++ {
                                        var v87 bool = true
                                        var v88 string = (v1 + v1)
                                        var v89 int = (fi86 >> (((^(v3))) & 63))
                                        _ = fi86
                                        _ = v87
                                        _ = v88
                                        _ = v89
                                    }
                                    _ = v85
                                } else {
                                    var v90 int = 128
                                    var v91 uint32 = v4
                                    _ = v90
                                    _ = v91
                                }
                                v1 = "foo bar"
                                for fi92 := 0; fi92 < 3; fi92++ {
                                    switch fi92 {
                                    case 2147483647:
                                        fmt.Println(fi92, v3)
                                    case 1:
                                        v14 = v14
                                    default:
                                        v4 += (^(v4))
                                    }
                                    v93 := GenS9{F0: 115}
                                    _ = fi92
                                    _ = v93
                                }
                            } else {
                                if ((v3 << (((v3 & (^(v3)))) & 63)) != (v3 - v3)) {
                                    v94 := GenS4{F0: 0, F1: 100}
                                    if true {
                                        v4 = (v4 * v4)
                                        if (32767 == 1) {
                                            var v95 uint = v3
                                            v95 = (v3 | (v3 % ((v3) | 1)))
                                            _ = v95
                                        } else {
                                            if false {
                                                if false {
                                                    for fi96 := 0; fi96 < 12; fi96++ {
                                                        if true {
                                                            var v97 int8 = -1
                                                            v1 = v1
                                                            _ = v97
                                                        }
                                                        var v98 uint32 = v4
                                                        _ = fi96
                                                        _ = v98
                                                    }
                                                    for fi99 := 0; fi99 < 18; fi99++ {
                                                        if len(v5) >= 2 {
                                                            v5 = v5[1:len(v5)]
                                                        }
                                                        _ = fi99
                                                    }
                                                    v4 += 65535
                                                } else {
                                                    switch -905 {
                                                    case 9223372036854775807:
                                                        for fi100 := 0; fi100 < 12; fi100++ {
                                                            for fi101 := 0; fi101 < 10; fi101++ {
                                                                if true {
                                                                    switch v1 {
                                                                    case "\n\t":
                                                                        _ = len(v2)
                                                                    case "a":
                                                                        if len(v5) > 0 {
                                                                            v5[0] = (v4 | 2147483648)
                                                                        }
                                                                    default:
                                                                        if (564 > (v14 >> ((65535) & 63))) {
                                                                            if false {
                                                                                fmt.Println(v3)
                                                                                var v102 int32 = 0
                                                                                v3 += v3
                                                                                v103 := GenS10{F0: v3}
                                                                                if ((-90 >= 40) && true) {
                                                                                    if false {
                                                                                        if (0.0 >= ((0.0 / 1.0) + (1.5 - 0.0))) {
                                                                                            v3 -= v3
                                                                                            fmt.Println(v1, fi100, v14)
                                                                                        }
                                                                                        v104 := GenS9{F0: 742}
                                                                                        _ = v104
                                                                                    }
                                                                                    if len(v5) > 0 {
                                                                                        v5[0] = 42
                                                                                    }
                                                                                } else {
                                                                                    v105 := GenS5{F0: 1, F1: 1}
                                                                                    if true {
                                                                                        v106 := GenS7{F0: -446, F1: "hello"}
                                                                                        var v107 int8 = int8(v4)
                                                                                        var v108 uint = 0
                                                                                        if true {
                                                                                            v107 = (v107 * (v107 ^ 100))
                                                                                        }
                                                                                        if len(v5) >= 2 {
                                                                                            v5 = v5[1:len(v5)]
                                                                                        }
                                                                                        if len(v5) > 0 {
                                                                                            v5[0] = (v4 & 2147483647)
                                                                                        }
                                                                                        fmt.Println(v14, v108)
                                                                                        v14 = 0
                                                                                        var v109 int16 = 42
                                                                                        _ = len(v5)
                                                                                        _ = v106
                                                                                        _ = v107
                                                                                        _ = v108
                                                                                        _ = v109
                                                                                    } else {
                                                                                        if len(v5) >= 2 {
                                                                                            v5 = v5[1:len(v5)]
                                                                                        }
                                                                                    }
                                                                                    v14 -= (v14 >> (((v3 - 720)) & 63))
                                                                                    var v110 float64 = 0.0
                                                                                    _ = v105
                                                                                    _ = v110
                                                                                }
                                                                                if len(v5) > 0 {
                                                                                    v5[0] = 65535
                                                                                }
                                                                                _ = v102
                                                                                _ = v103
                                                                            } else {
                                                                                if (float32(0.1) < -2.25) {
                                                                                    fmt.Println(v4, v1)
                                                                                }
                                                                                v4 = (v4 ^ 2147483647)
                                                                                var v111 float64 = 1e10
                                                                                _ = v111
                                                                            }
                                                                            var v112 bool = ((-1 >= (fi100 + 1)) && false)
                                                                            var v113 int8 = 1
                                                                            _ = v112
                                                                            _ = v113
                                                                        } else {
                                                                            if false {
                                                                                switch fi101 {
                                                                                case 42:
                                                                                    if false {
                                                                                        v1 = "\n\t"
                                                                                        if (2147483647 < 2147483647) {
                                                                                            var v114 int16 = (v14 - (v14 + (^(v14))))
                                                                                            _ = v114
                                                                                        } else {
                                                                                            v115 := GenS11{F0: v1, F1: (v14 << ((v3) & 63))}
                                                                                            var v116 uint16 = 542
                                                                                            _ = v115
                                                                                            _ = v116
                                                                                        }
                                                                                    }
                                                                                case -1:
                                                                                    if len(v5) >= 2 {
                                                                                        v5 = v5[1:len(v5)]
                                                                                    }
                                                                                    var v117 uint64 = 0
                                                                                    if true {
                                                                                        var v118 int8 = -128
                                                                                        _ = len(v2)
                                                                                        _ = v118
                                                                                    } else {
                                                                                        var v119 int64 = 9223372036854775807
                                                                                        var v120 int = fi100
                                                                                        fmt.Println(v14, v1, fi100)
                                                                                        _ = v119
                                                                                        _ = v120
                                                                                    }
                                                                                    _ = v117
                                                                                case -804:
                                                                                    if len(v5) >= 2 {
                                                                                        v5 = v5[1:len(v5)]
                                                                                    }
                                                                                }
                                                                                if true {
                                                                                    if (42 >= 679) {
                                                                                        switch "hello" {
                                                                                        case "":
                                                                                            if (true && true) {
                                                                                                if true {
                                                                                                    delete(v2, (fi100 - (fi100 + -535)))
                                                                                                    if (209 != 868) {
                                                                                                        v1 = v1
                                                                                                    } else {
                                                                                                        if false {
                                                                                                            var v121 float32 = ((3.14159 + (((float32(0.1) * 1e10) - (0.0 / float32((0.1 + 1.5)))) / float32(((-2.25 - float64(0.1)) * 1.5)))) - 0.0)
                                                                                                            if (!(100 >= 127)) {
                                                                                                                v14 = v14
                                                                                                                fmt.Println(v121)
                                                                                                                var v122 int64 = 272
                                                                                                                _ = len(v5)
                                                                                                                v123 := GenS12{F0: 1000, F1: fi101, F2: v4}
                                                                                                                if (-2147483648 > -881) {
                                                                                                                    v3 += 1
                                                                                                                }
                                                                                                                _ = v122
                                                                                                                _ = v123
                                                                                                            } else {
                                                                                                                v14 += (v14 | 3)
                                                                                                                if true {
                                                                                                                    v1 = v1
                                                                                                                    fmt.Println(v3, v121, v121)
                                                                                                                } else {
                                                                                                                    var v124 string = ("\n\t" + "a")
                                                                                                                    if true {
                                                                                                                        var v125 float32 = float32(-2.25)
                                                                                                                        _ = v125
                                                                                                                    }
                                                                                                                    _ = v124
                                                                                                                }
                                                                                                            }
                                                                                                            v5 = append(v5, v4)
                                                                                                            _ = v121
                                                                                                        } else {
                                                                                                            switch fi101 {
                                                                                                            case 248:
                                                                                                                if true {
                                                                                                                    v126 := v2[fi100]
                                                                                                                    v14 += 42
                                                                                                                    _ = v126
                                                                                                                }
                                                                                                            default:
                                                                                                                var v127 uint64 = 4294967295
                                                                                                                _ = v127
                                                                                                            }
                                                                                                        }
                                                                                                        if len(v5) > 0 {
                                                                                                            v5[0] = 1
                                                                                                        }
                                                                                                        v14 = v14
                                                                                                    }
                                                                                                    v14 = 1
                                                                                                } else {
                                                                                                    if ((fi101 == fi100) && true) {
                                                                                                        v128 := GenS13{F0: (-(v14)), F1: 2, F2: 32767}
                                                                                                        _ = v128
                                                                                                    } else {
                                                                                                        var v129 uint = 42
                                                                                                        _ = v129
                                                                                                    }
                                                                                                }
                                                                                                var v130 uint8 = 127
                                                                                                if len(v5) >= 2 {
                                                                                                    v5 = v5[1:len(v5)]
                                                                                                }
                                                                                                _ = v130
                                                                                            }
                                                                                        case "a":
                                                                                            switch v1 {
                                                                                            case "hello":
                                                                                                v3 -= 181
                                                                                                var v131 uint64 = 2
                                                                                                v131 = 18446744073709551615
                                                                                                _ = v131
                                                                                            case "\n\t":
                                                                                                if len(v5) > 0 {
                                                                                                    v5[0] = (v4 * v4)
                                                                                                }
                                                                                            default:
                                                                                                v3 -= v3
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                    v132 := GenS11{F0: (("foo bar" + v1) + ((v1 + (v1 + "foo bar")) + "a")), F1: -1}
                                                                                    _ = v132
                                                                                } else {
                                                                                    v133 := GenS9{F0: 65535}
                                                                                    fmt.Println(v3, v3)
                                                                                    _ = v133
                                                                                }
                                                                            } else {
                                                                                v134 := GenS12{F0: 32767, F1: (fi101 * fi100), F2: (v4 >> ((v3) & 63))}
                                                                                _ = v134
                                                                            }
                                                                            delete(v2, fi101)
                                                                        }
                                                                        fmt.Println(v1, fi100, fi101)
                                                                    }
                                                                    _ = len(v5)
                                                                    v14 += 206
                                                                }
                                                                if false {
                                                                    v135 := GenS13{F0: 580, F1: 2, F2: 65535}
                                                                    v136 := GenS14{F0: 2147483647}
                                                                    v3 = (v3 >> ((v3) & 63))
                                                                    _ = v135
                                                                    _ = v136
                                                                } else {
                                                                    v1 = (("a" + "") + v1)
                                                                }
                                                                _ = fi101
                                                            }
                                                            if true {
                                                                var v137 uint8 = 112
                                                                _ = v137
                                                            }
                                                            for fi138 := 0; fi138 < 12; fi138++ {
                                                                fmt.Println(fi138, fi100)
                                                                var v139 uint32 = (v4 + v4)
                                                                v140 := GenS9{F0: 65535}
                                                                var v141 int = fi138
                                                                var v142 uint32 = (v4 | 970)
                                                                _ = fi138
                                                                _ = v139
                                                                _ = v140
                                                                _ = v141
                                                                _ = v142
                                                            }
                                                            _ = fi100
                                                        }
                                                    case 127:
                                                        v5 = append(v5, v4)
                                                    case -441:
                                                        v143 := GenS7{F0: -1, F1: ((v1 + v1) + v1)}
                                                        _ = v143
                                                    default:
                                                        v144 := v2[489]
                                                        _ = v144
                                                    }
                                                    fmt.Println(v3, v14, v3)
                                                    var v145 int64 = -1
                                                    _ = v145
                                                }
                                                var v146 int8 = 0
                                                _ = v146
                                            } else {
                                                if true {
                                                    var v147 bool = false
                                                    v3 = (v3 & (v3 / ((v3) | 1)))
                                                    v148 := GenS4{F0: 18446744073709551615, F1: -14}
                                                    _ = v147
                                                    _ = v148
                                                } else {
                                                    for fi149 := 0; fi149 < 3; fi149++ {
                                                        v3 += (v3 | (^(v3)))
                                                        if (1 >= 2147483647) {
                                                            for fi150 := 0; fi150 < 15; fi150++ {
                                                                if (0.1 >= float64((((0.1 + (1e10 + 1e10)) / 1.0) / 1e10))) {
                                                                    switch (^(fi149)) {
                                                                    case -129:
                                                                        var v151 uint16 = 686
                                                                        _ = v151
                                                                    case 127:
                                                                        var v152 int32 = -855
                                                                        _ = v152
                                                                    case 0:
                                                                        fmt.Println(v4, fi149)
                                                                    default:
                                                                        var v153 int64 = 9223372036854775807
                                                                        v3 = v3
                                                                        _ = v153
                                                                    }
                                                                    v14 = v14
                                                                } else {
                                                                    if true {
                                                                        var v154 int8 = 127
                                                                        v14 += (v14 * v14)
                                                                        _ = v154
                                                                    } else {
                                                                        if true {
                                                                            switch fi150 {
                                                                            case 4294967295:
                                                                                v4 = v4
                                                                                switch fi150 {
                                                                                case 717:
                                                                                    fmt.Println(v1, v14)
                                                                                case -9223372036854775808:
                                                                                    fmt.Println(v1, v1)
                                                                                case 32767:
                                                                                    v4 -= v4
                                                                                }
                                                                            case -556:
                                                                                fmt.Println(v4)
                                                                            case 9223372036854775807:
                                                                                v155 := GenS2{F0: 880, F1: (v1 + v1)}
                                                                                _ = v155
                                                                            }
                                                                            v156 := GenS15{F0: fi150, F1: v1, F2: v4}
                                                                            v157 := v2[fi150]
                                                                            _ = v156
                                                                            _ = v157
                                                                        }
                                                                        fmt.Println(fi150, fi149)
                                                                        if (2147483647 > 896) {
                                                                            switch fi149 {
                                                                            case 9223372036854775807:
                                                                                fmt.Println(fi149, v4, v3)
                                                                            case 1:
                                                                                if true {
                                                                                    var v158 string = v1
                                                                                    _ = v158
                                                                                } else {
                                                                                    var v159 uint8 = 2
                                                                                    var v160 uint16 = 42
                                                                                    _ = v159
                                                                                    _ = v160
                                                                                }
                                                                                var v161 int32 = 0
                                                                                _ = v161
                                                                            default:
                                                                                v162 := GenS4{F0: 327, F1: 100}
                                                                                _ = v162
                                                                            }
                                                                        } else {
                                                                            if true {
                                                                                switch v1 {
                                                                                case "foo bar":
                                                                                    v1 = v1
                                                                                case "hello":
                                                                                    v1 = "foo bar"
                                                                                default:
                                                                                    var v163 int8 = 103
                                                                                    _ = v163
                                                                                }
                                                                                var v164 uint = 2
                                                                                _ = v164
                                                                            }
                                                                            if (128 != 42) {
                                                                                var v165 bool = false
                                                                                if false {
                                                                                    if (!v165) {
                                                                                        var v166 uint64 = 200
                                                                                        var v167 string = ("foo bar" + v1)
                                                                                        _ = len(v2)
                                                                                        _ = v166
                                                                                        _ = v167
                                                                                    }
                                                                                }
                                                                                if len(v5) >= 2 {
                                                                                    v5 = v5[1:len(v5)]
                                                                                }
                                                                                _ = v165
                                                                            } else {
                                                                                v1 = "foo bar"
                                                                                if false {
                                                                                    var v168 bool = ((705 < -9223372036854775808) || true)
                                                                                    _ = v168
                                                                                } else {
                                                                                    v1 = v1
                                                                                    var v169 int8 = -46
                                                                                    _ = v169
                                                                                }
                                                                            }
                                                                            v170 := GenS13{F0: (^(v14)), F1: 128, F2: 32767}
                                                                            _ = v170
                                                                        }
                                                                    }
                                                                    switch ("foo bar" + v1) {
                                                                    case "\n\t":
                                                                        fmt.Println(v1)
                                                                    case "foo bar":
                                                                        v171 := GenS10{F0: v3}
                                                                        var v172 bool = ((-110 < -145) && false)
                                                                        fmt.Println(v4, v172, v4)
                                                                        _ = v171
                                                                        _ = v172
                                                                    case "a":
                                                                        if true {
                                                                            fmt.Println(fi149, fi149)
                                                                            v4 = v4
                                                                            v1 = v1
                                                                            var v173 float32 = float32(-2.25)
                                                                            var v174 int32 = 884
                                                                            var v175 float32 = v173
                                                                            v173 = 1e10
                                                                            _ = v173
                                                                            _ = v174
                                                                            _ = v175
                                                                        } else {
                                                                            fmt.Println(fi150, v1)
                                                                            v2[(-(fi149))] = false
                                                                            var v176 uint16 = 65535
                                                                            _ = v176
                                                                        }
                                                                    }
                                                                }
                                                                var v177 string = v1
                                                                var v178 int32 = 1
                                                                _ = fi150
                                                                _ = v177
                                                                _ = v178
                                                            }
                                                        } else {
                                                            var v179 uint8 = 127
                                                            v3 -= v3
                                                            _ = v179
                                                        }
                                                        _ = fi149
                                                    }
                                                    for fi180 := 0; fi180 < 19; fi180++ {
                                                        for fi181 := 0; fi181 < 18; fi181++ {
                                                            if len(v5) > 0 {
                                                                v5[0] = v4
                                                            }
                                                            delete(v2, (fi180 % ((fi181) | 1)))
                                                            _ = fi181
                                                        }
                                                        var v182 int8 = 91
                                                        _ = fi180
                                                        _ = v182
                                                    }
                                                }
                                                if (v14 == 576) {
                                                    for fi183 := 0; fi183 < 7; fi183++ {
                                                        switch (v1 + (v1 + ((v1 + (((("hello" + (v1 + (("hello" + "hello") + v1))) + "\n\t") + "a") + v1)) + (v1 + (v1 + ((("" + v1) + v1) + "hello")))))) {
                                                        case "foo bar":
                                                            v1 = v1
                                                        }
                                                        _ = fi183
                                                    }
                                                }
                                                if len(v5) > 0 {
                                                    v5[0] = v4
                                                }
                                            }
                                            _ = len(v5)
                                            for fi184 := 0; fi184 < 15; fi184++ {
                                                if false {
                                                    switch fi184 {
                                                    case 702:
                                                        var v185 uint16 = 1
                                                        _ = v185
                                                    case 2147483647:
                                                        v4 = (v4 & v4)
                                                    default:
                                                        v186 := GenS14{F0: v4}
                                                        _ = v186
                                                    }
                                                    if (0 <= 42) {
                                                        v14 = v14
                                                    }
                                                    for fi187 := 0; fi187 < 10; fi187++ {
                                                        if (!true) {
                                                            var v188 uint16 = 0
                                                            if true {
                                                                var v189 float32 = (float32(1e10) + 3.14159)
                                                                _ = v189
                                                            } else {
                                                                fmt.Println(v4, v14, v1)
                                                                var v190 uint = 856
                                                                v3 -= v190
                                                                _ = v190
                                                            }
                                                            if (true || (1.5 >= float64((float32(3.14159) - (-2.25 - float32((1e10 - (0.1 / 1.0)))))))) {
                                                                v191 := GenS16{F0: 42, F1: uint64(v188)}
                                                                v192 := GenS17{F0: v188, F1: v3}
                                                                _ = v191
                                                                _ = v192
                                                            } else {
                                                                v1 = v1
                                                                fmt.Println(v1, v1, v4)
                                                                v1 = v1
                                                                var v193 uint16 = (^(v188))
                                                                _ = v193
                                                            }
                                                            _ = v188
                                                        }
                                                        _ = fi187
                                                    }
                                                } else {
                                                    var v194 int32 = -2147483648
                                                    fmt.Println(v1, v1)
                                                    _ = v194
                                                }
                                                _ = fi184
                                            }
                                        }
                                    } else {
                                        v4 += v4
                                        var v195 bool = false
                                        var v196 bool = (1 != 65535)
                                        _ = v195
                                        _ = v196
                                    }
                                    _ = v94
                                } else {
                                    var v197 string = ""
                                    v198 := GenS3{F0: 1, F1: -1, F2: 42}
                                    v199 := GenS7{F0: 443, F1: ("\n\t" + (v197 + "\n\t"))}
                                    _ = v197
                                    _ = v198
                                    _ = v199
                                }
                            }
                        case "a":
                            v1 = "hello"
                        default:
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                        }
                    case "hello":
                        var v200 uint = v3
                        _ = v200
                    }
                }
            }
            for fi201 := 0; fi201 < 3; fi201++ {
                v14 -= v14
                v3 += 65535
                _ = fi201
            }
            if (0 != 32767) {
                for fi202 := 0; fi202 < 3; fi202++ {
                    for fi203 := 0; fi203 < 15; fi203++ {
                        if false {
                            if (1e10 <= float64(0.0)) {
                                v204 := GenS18{F0: 128, F1: 9223372036854775807}
                                v5 = append(v5, 65535)
                                _ = v204
                            }
                            switch (-(fi203)) {
                            case 127:
                                v205 := GenS19{F0: 2, F1: v1}
                                _ = v205
                            case -9223372036854775808:
                                v4 = v4
                            }
                        }
                        _ = fi203
                    }
                    v5 = append(v5, (v4 * (v4 | v4)))
                    var v206 uint = (^(v3))
                    _ = fi202
                    _ = v206
                }
                switch -2147483648 {
                case 32768:
                    v5 = append(v5, (v4 ^ (^(v4))))
                default:
                    var v207 uint8 = 1
                    _ = v207
                }
            }
        }
        if false {
            v1 = v1
            for fi208 := 0; fi208 < 7; fi208++ {
                switch fi208 {
                case 4294967295:
                    v209 := v2[fi208]
                    _ = v209
                case 32768:
                    fmt.Println(fi208, v1)
                case 616:
                    var v210 int32 = -2147483648
                    _ = v210
                default:
                    v2[(fi208 << ((v3) & 63))] = true
                }
                for fi211 := 0; fi211 < 15; fi211++ {
                    if (!("" < v1)) {
                        switch (fi211 << ((v3) & 63)) {
                        case -749:
                            if false {
                                var v212 int = (-(fi208))
                                _ = v212
                            }
                        case 1:
                            if true {
                                switch v1 {
                                case "a":
                                    var v213 uint64 = 2
                                    _ = v213
                                case "\n\t":
                                    if len(v5) > 0 {
                                        v5[0] = v4
                                    }
                                case "foo bar":
                                    fmt.Println(fi208, fi208, fi211)
                                }
                            } else {
                                var v214 string = ("\n\t" + (v1 + (v1 + ("foo bar" + ((((("foo bar" + v1) + (("foo bar" + v1) + ((((v1 + v1) + ("a" + v1)) + v1) + v1))) + (("" + v1) + v1)) + v1) + "\n\t")))))
                                fmt.Println(fi208, v214)
                                v1 = v214
                                _ = v214
                            }
                        case 627:
                            v14 += 1
                        }
                        if (!(!true)) {
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                            v1 = v1
                            if true {
                                var v215 int16 = (v14 % ((42) | 1))
                                var v216 string = v1
                                v217 := GenS18{F0: 128, F1: -2147483648}
                                _ = v215
                                _ = v216
                                _ = v217
                            }
                        } else {
                            var v218 uint16 = 2
                            var v219 int = fi208
                            delete(v2, (v219 << (((v3 * uint(v218))) & 63)))
                            _ = v218
                            _ = v219
                        }
                        if true {
                            v2[(fi211 >> ((0) & 63))] = (127 > 0)
                            var v220 uint64 = 4294967295
                            _ = v220
                        } else {
                            fmt.Println(fi211, v14)
                            var v221 int = (fi211 | 42)
                            if (131 <= -9223372036854775808) {
                                fmt.Println(v14, fi211)
                            } else {
                                v222 := GenS20{F0: 4294967295, F1: (^(v3)), F2: "\n\t"}
                                var v223 uint16 = 31
                                _ = v222
                                _ = v223
                            }
                            _ = v221
                        }
                    }
                    v4 = (v4 / ((2147483647) | 1))
                    var v224 int8 = 127
                    _ = fi211
                    _ = v224
                }
                v4 += 2
                _ = fi208
            }
            if true {
                v225 := GenS5{F0: 0, F1: int64(v3)}
                _ = v225
            }
        } else {
            for fi226 := 0; fi226 < 15; fi226++ {
                v227 := v2[(fi226 & (fi226 - fi226))]
                v1 = "foo bar"
                if (!v227) {
                    switch "hello" {
                    case "hello":
                        var v228 uint16 = 42
                        _ = v228
                    case "foo bar":
                        _ = len(v5)
                    }
                    _ = len(v5)
                }
                _ = fi226
                _ = v227
            }
        }
        _ = v14
    } else {
        if (128 == 155) {
            for fi229 := 0; fi229 < 13; fi229++ {
                v4 = v4
                _ = fi229
            }
        } else {
            for fi230 := 0; fi230 < 16; fi230++ {
                v4 += (v4 & v4)
                fmt.Println(v4, v1, v3)
                if len(v5) > 0 {
                    v5[0] = (^(v4))
                }
                _ = fi230
            }
            v231 := GenS5{F0: -58, F1: -1}
            _ = v231
        }
        fmt.Println(v4)
        v232 := GenS21{F0: -101, F1: 161, F2: 42}
        _ = v232
    }
    var v233 int16 = -1
    var v234 uint8 = 175
    fmt.Println(v234, v4)
    v3 -= v3
    v5 = append(v5, v4)
    for fi235 := 0; fi235 < 11; fi235++ {
        var v236 int64 = -9223372036854775808
        if ((v234 * (^(v234))) >= v234) {
            _ = len(v2)
        }
        fmt.Println(v4)
        _ = fi235
        _ = v236
    }
    var v237 uint64 = 574
    v238 := GenS9{F0: 410}
    var v239 int64 = -783
    var v240 int32 = 0
    fmt.Println(v4, v234, v3)
    v5 = append(v5, v4)
    if false {
        if true {
            v241 := GenS13{F0: -32768, F1: (v234 - (^(v234))), F2: 0}
            if true {
                if (!false) {
                    if true {
                        var v242 uint32 = (v4 & (v4 | (v4 % ((v4) | 1))))
                        _ = v242
                    } else {
                        if true {
                            fmt.Println(v240, v3, v240)
                            var v243 float32 = float32((-2.25 + 1e10))
                            var v244 uint8 = 0
                            _ = v243
                            _ = v244
                        }
                        v239 = v239
                        fmt.Println(v234, v234, v3)
                    }
                } else {
                    for fi245 := 0; fi245 < 9; fi245++ {
                        if true {
                            var v246 uint32 = (v4 + (v4 << ((18446744073709551615) & 63)))
                            _ = v246
                        }
                        _ = fi245
                    }
                }
                v247 := GenS22{F0: 9223372036854775807}
                _ = v247
            }
            _ = v241
        }
        if false {
            for fi248 := 0; fi248 < 2; fi248++ {
                v4 = uint32(v237)
                var v249 uint = v3
                var v250 uint32 = 994
                _ = fi248
                _ = v249
                _ = v250
            }
            _ = len(v2)
            v3 = v3
        } else {
            v233 -= (^(v233))
            var v251 int = 2147483647
            _ = v251
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v233:", v233)
    fmt.Println("v234:", v234)
    fmt.Println("v237:", v237)
    fmt.Println("v239:", v239)
    fmt.Println("v240:", v240)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v233
    _ = v234
    _ = v237
    _ = v238
    _ = v239
    _ = v240
    _ = v3
    _ = v4
    _ = v5
}
