package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int32; F1 uint16; F2 int64 }
type GenS4 struct { F0 int64; F1 uint64; F2 int8 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 int64; F1 int8; F2 int64 }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 int64 }
type GenS9 struct { F0 uint8; F1 int32; F2 uint8 }
type GenS10 struct { F0 string }
type GenS11 struct { F0 uint8; F1 int16 }
type GenS12 struct { F0 int64; F1 uint64 }
type GenS13 struct { F0 uint64 }
type GenS14 struct { F0 int32; F1 uint32; F2 uint16 }

func main() {
    var v1 float32 = 1e10
    var v2 int32 = -2147483648
    v3 := []uint32{4294967295}
    var v4 string = ("foo bar" + "foo bar")
    var v5 int64 = 300
    var v6 float32 = (3.14159 / 1e10)
    if false {
        var v7 uint = 4294967295
        v8 := GenS1{F0: 274, F1: ((!(4294967295 == 669)) && true)}
        var v9 int8 = 23
        _ = v7
        _ = v8
        _ = v9
    }
    if (2 > 2) {
        for fi10 := 0; fi10 < 9; fi10++ {
            if ((894 < 42) && (2147483647 <= 65535)) {
                for fi11 := 0; fi11 < 20; fi11++ {
                    var v12 uint32 = 194
                    _ = fi11
                    _ = v12
                }
                v2 -= 1
            } else {
                for fi13 := 0; fi13 < 8; fi13++ {
                    fmt.Println(fi13)
                    fmt.Println(v4, v1, v2)
                    v2 = (v2 % ((0) | 1))
                    _ = fi13
                }
            }
            v6 = 1.5
            var v14 bool = (1.5 == 0.0)
            _ = fi10
            _ = v14
        }
        var v15 string = ("\n\t" + v4)
        _ = v15
    }
    v16 := GenS2{F0: 22}
    v17 := GenS3{F0: -84, F1: 474, F2: (^(v5))}
    var v18 float64 = float64(v6)
    fmt.Println(v18, v18)
    for fi19 := 0; fi19 < 1; fi19++ {
        var v20 string = v4
        fmt.Println(v1, v18, v4)
        _ = fi19
        _ = v20
    }
    if true {
        if ((false && false) || (!false)) {
            v5 = v5
        } else {
            var v21 uint8 = 42
            var v22 float64 = v18
            for fi23 := 0; fi23 < 14; fi23++ {
                v2 = v2
                var v24 uint8 = (v21 * 63)
                v25 := GenS4{F0: v5, F1: 42, F2: -1}
                _ = fi23
                _ = v24
                _ = v25
            }
            _ = v21
            _ = v22
        }
        if (true && (1 <= 2)) {
            fmt.Println(v2)
            v3 = append(v3, 1)
        } else {
            v4 = "a"
            v1 += (v1 * v1)
            v3 = append(v3, 2147483648)
        }
        _ = len(v3)
    }
    if true {
        switch v4 {
        case "foo bar":
            for fi26 := 0; fi26 < 9; fi26++ {
                var v27 string = "a"
                for fi28 := 0; fi28 < 17; fi28++ {
                    if (((-(v2)) > (v2 << ((4294967295) & 63))) && (false && (0 >= 0))) {
                        v6 += v6
                        var v29 int8 = int8(v5)
                        _ = v29
                    }
                    _ = fi28
                }
                if (2 > 2) {
                    v1 = 1e10
                    if ((^(v2)) != v2) {
                        for fi30 := 0; fi30 < 6; fi30++ {
                            v6 -= v6
                            switch "" {
                            case "a":
                                var v31 float32 = v1
                                v27 = (v27 + "a")
                                v32 := GenS5{F0: false}
                                _ = v31
                                _ = v32
                            case "":
                                var v33 int32 = v2
                                _ = v33
                            case "hello":
                                var v34 float64 = 0.0
                                _ = v34
                            default:
                                var v35 int16 = 730
                                _ = v35
                            }
                            v4 = ("\n\t" + v27)
                            v5 = -9223372036854775808
                            _ = fi30
                        }
                        var v36 uint64 = 2
                        _ = v36
                    } else {
                        for fi37 := 0; fi37 < 18; fi37++ {
                            switch (v27 + "foo bar") {
                            case "hello":
                                v4 = v27
                            case "a":
                                if false {
                                    var v38 float32 = 1e10
                                    _ = v38
                                } else {
                                    _ = len(v3)
                                }
                            case "foo bar":
                                switch fi26 {
                                case 42:
                                    if (2 == 128) {
                                        v6 -= (v1 / v6)
                                    } else {
                                        v18 -= 1e10
                                        v6 += (v6 * v1)
                                        v18 = (v18 + (-2.25 * float64(v1)))
                                    }
                                case 65:
                                    v39 := GenS5{F0: (false && true)}
                                    if true {
                                        v27 = (((v27 + ((v27 + v4) + v27)) + v27) + v27)
                                        fmt.Println(v4, v27)
                                        switch v4 {
                                        case "\n\t":
                                            if true {
                                                fmt.Println(fi37, v2)
                                                switch fi26 {
                                                case 0:
                                                    if len(v3) > 0 {
                                                        v3[0] = 2
                                                    }
                                                case -25:
                                                    v4 = v27
                                                case -9223372036854775808:
                                                    fmt.Println(fi26, fi37, fi26)
                                                }
                                                switch fi37 {
                                                case 9223372036854775807:
                                                    if len(v3) > 0 {
                                                        v3[0] = 2
                                                    }
                                                case 2147483647:
                                                    var v40 float32 = float32((((v18 * (0.0 + ((float64(v6) - v18) - ((1e10 - v18) - 1e10)))) - v18) * 1.5))
                                                    _ = v40
                                                }
                                                if false {
                                                    v41 := GenS6{F0: v5, F1: 16, F2: 0}
                                                    _ = v41
                                                } else {
                                                    var v42 int16 = int16(v5)
                                                    _ = v42
                                                }
                                                v43 := GenS5{F0: true}
                                                _ = v43
                                            }
                                        case "":
                                            v44 := GenS2{F0: 255}
                                            _ = v44
                                        case "foo bar":
                                            switch ("a" + v4) {
                                            case "":
                                                if false {
                                                    v45 := GenS5{F0: (255 > 32767)}
                                                    if true {
                                                        if len(v3) > 0 {
                                                            v3[0] = 4294967295
                                                        }
                                                        v1 -= v1
                                                    }
                                                    if true {
                                                        switch "foo bar" {
                                                        case "":
                                                            if len(v3) >= 2 {
                                                                v3 = v3[1:len(v3)]
                                                            }
                                                        case "foo bar":
                                                            v3 = append(v3, 2147483648)
                                                            var v46 float64 = float64((((-2.25 * (float32((v18 + (v18 - -2.25))) + 0.1)) / -2.25) / ((((v6 / float32((0.1 + v18))) / ((v1 + (((0.0 + v1) / (3.14159 - v6)) - v1)) + 1e10)) + (float32((0.0 + 3.14159)) - v1)) - v6)))
                                                            if len(v3) > 0 {
                                                                v3[0] = 2
                                                            }
                                                            _ = v46
                                                        default:
                                                            v1 -= v6
                                                        }
                                                    } else {
                                                        if len(v3) > 0 {
                                                            v3[0] = 65535
                                                        }
                                                    }
                                                    _ = v45
                                                }
                                            case "hello":
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                            case "a":
                                                if true {
                                                    v2 -= -764
                                                    switch (fi26 + fi26) {
                                                    case 32767:
                                                        var v47 uint64 = 653
                                                        _ = v47
                                                    case 32768:
                                                        if len(v3) > 0 {
                                                            v3[0] = 2
                                                        }
                                                    case -424:
                                                        if len(v3) > 0 {
                                                            v3[0] = 2147483648
                                                        }
                                                    default:
                                                        fmt.Println(v6, v18)
                                                    }
                                                }
                                            default:
                                                switch (fi26 - fi26) {
                                                case 32768:
                                                    v27 = ((("a" + v27) + v27) + v4)
                                                    var v48 int32 = v2
                                                    v18 = v18
                                                    _ = v48
                                                case -2147483648:
                                                    var v49 float64 = (v18 - 1.5)
                                                    _ = v49
                                                case -871:
                                                    fmt.Println(v27, v5)
                                                }
                                            }
                                        default:
                                            fmt.Println(fi26, v5, v4)
                                        }
                                    }
                                    _ = v39
                                }
                            }
                            v50 := GenS5{F0: ((0 == 1) || false)}
                            _ = fi37
                            _ = v50
                        }
                        for fi51 := 0; fi51 < 2; fi51++ {
                            var v52 uint64 = 143
                            switch (fi26 >> ((0) & 63)) {
                            case 9223372036854775807:
                                var v53 int16 = 1000
                                _ = v53
                            }
                            v52 = (v52 % (((v52 % (((v52 - 1)) | 1))) | 1))
                            _ = fi51
                            _ = v52
                        }
                        var v54 string = (v27 + "a")
                        _ = v54
                    }
                    var v55 int16 = 444
                    _ = v55
                }
                _ = fi26
                _ = v27
            }
        case "hello":
            if ((v5 | (v5 - 331)) < v5) {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                var v56 int64 = -457
                _ = v56
            }
        case "":
            v57 := GenS4{F0: (v5 ^ (v5 % (((v5 & (^(v5)))) | 1))), F1: 2, F2: -25}
            _ = v57
        }
        switch v4 {
        case "a":
            v18 = (v18 * float64(v1))
        case "foo bar":
            if len(v3) > 0 {
                v3[0] = 0
            }
        default:
            if true {
                switch v4 {
                case "foo bar":
                    if (32767 > 0) {
                        switch "a" {
                        case "a":
                            var v58 int16 = 0
                            _ = v58
                        case "hello":
                            v59 := GenS7{F0: 0}
                            _ = v59
                        default:
                            var v60 int8 = -128
                            _ = v60
                        }
                        v61 := GenS6{F0: (v5 / ((v5) | 1)), F1: 58, F2: v5}
                        _ = v61
                    } else {
                        v62 := GenS8{F0: (v5 | 9223372036854775807)}
                        for fi63 := 0; fi63 < 20; fi63++ {
                            switch (fi63 >> ((18446744073709551615) & 63)) {
                            case 32768:
                                v2 = v2
                            case 21:
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                            case -1:
                                var v64 float32 = v1
                                _ = v64
                            default:
                                var v65 uint32 = 0
                                _ = v65
                            }
                            _ = len(v3)
                            if (254 <= 1) {
                                fmt.Println(v6, v2)
                                if ((v1 + (v6 / v1)) <= (v6 - 0.1)) {
                                    var v66 int16 = int16(fi63)
                                    v18 += ((v18 / 3.14159) * v18)
                                    v4 = (v4 + "")
                                    _ = v66
                                } else {
                                    if false {
                                        switch (fi63 / ((fi63) | 1)) {
                                        case 9223372036854775807:
                                            v67 := GenS9{F0: 128, F1: (v2 / ((v2) | 1)), F2: 0}
                                            _ = v67
                                        default:
                                            switch (v4 + v4) {
                                            case "\n\t":
                                                var v68 float32 = 1.5
                                                _ = v68
                                            default:
                                                v1 = float32((v18 - v18))
                                            }
                                        }
                                        switch 128 {
                                        case 32768:
                                            var v69 int = (fi63 >> ((271) & 63))
                                            _ = v69
                                        case 9223372036854775807:
                                            v6 = (1e10 * (v1 * ((-2.25 + v1) - ((1.5 - (0.1 - v1)) / (v1 / (-2.25 + v1))))))
                                        }
                                        v18 = v18
                                    } else {
                                        if false {
                                            var v70 float64 = v18
                                            _ = v70
                                        }
                                    }
                                    fmt.Println(v5)
                                    if false {
                                        v71 := GenS10{F0: v4}
                                        if (2 >= 2147483647) {
                                            var v72 uint64 = 4294967295
                                            if (!(234 < 42)) {
                                                v72 -= v72
                                                fmt.Println(v4)
                                                var v73 bool = (1 < 4294967295)
                                                _ = v73
                                            } else {
                                                var v74 int8 = -116
                                                for fi75 := 0; fi75 < 5; fi75++ {
                                                    switch v4 {
                                                    case "foo bar":
                                                        v1 += v1
                                                    case "hello":
                                                        v3 = append(v3, 1)
                                                    case "\n\t":
                                                        v1 = float32(-2.25)
                                                    default:
                                                        if (0 >= -1) {
                                                            fmt.Println(v74)
                                                            if ((2 >= 718) && false) {
                                                                if ((v74 + (v74 - 100)) <= (v74 >> ((255) & 63))) {
                                                                    if ((((-(v74)) > 42) && (v74 == (-(v74)))) && true) {
                                                                        if (!true) {
                                                                            var v76 uint = 792
                                                                            if false {
                                                                                var v77 int = fi63
                                                                                var v78 string = v4
                                                                                _ = v77
                                                                                _ = v78
                                                                            } else {
                                                                                v3 = append(v3, 0)
                                                                                v79 := GenS11{F0: 1, F1: 0}
                                                                                _ = v79
                                                                            }
                                                                            _ = v76
                                                                        }
                                                                    }
                                                                } else {
                                                                    if len(v3) > 0 {
                                                                        v3[0] = 65535
                                                                    }
                                                                    if true {
                                                                        switch v4 {
                                                                        case "":
                                                                            switch v4 {
                                                                            case "foo bar":
                                                                                if (!(42 > 255)) {
                                                                                    v3 = append(v3, 2147483648)
                                                                                } else {
                                                                                    v72 = 4294967295
                                                                                    v3 = append(v3, 65535)
                                                                                }
                                                                                v80 := GenS7{F0: 1000}
                                                                                var v81 int16 = 42
                                                                                _ = v80
                                                                                _ = v81
                                                                            default:
                                                                                v72 += v72
                                                                            }
                                                                        }
                                                                        var v82 uint = 1
                                                                        v72 += v72
                                                                        _ = v82
                                                                    } else {
                                                                        if (!(false && false)) {
                                                                            var v83 int32 = 0
                                                                            v3 = append(v3, 2)
                                                                            _ = v83
                                                                        }
                                                                        var v84 bool = true
                                                                        var v85 int8 = v74
                                                                        _ = v84
                                                                        _ = v85
                                                                    }
                                                                }
                                                                v5 -= (v5 * (v5 >> ((185) & 63)))
                                                            } else {
                                                                v1 -= 3.14159
                                                            }
                                                        }
                                                    }
                                                    _ = fi75
                                                }
                                                if (!false) {
                                                    for fi86 := 0; fi86 < 1; fi86++ {
                                                        if len(v3) > 0 {
                                                            v3[0] = 1
                                                        }
                                                        _ = fi86
                                                    }
                                                    fmt.Println(v5)
                                                }
                                                _ = v74
                                            }
                                            _ = v72
                                        }
                                        _ = v71
                                    } else {
                                        v6 -= (v1 - (v1 / 1.0))
                                        var v87 uint32 = 864
                                        if false {
                                            var v88 bool = (!(1 != 193))
                                            _ = v88
                                        } else {
                                            for fi89 := 0; fi89 < 8; fi89++ {
                                                var v90 bool = true
                                                v5 -= v5
                                                _ = fi89
                                                _ = v90
                                            }
                                        }
                                        _ = v87
                                    }
                                }
                                for fi91 := 0; fi91 < 20; fi91++ {
                                    var v92 bool = false
                                    switch v4 {
                                    case "a":
                                        var v93 uint32 = 2
                                        _ = v93
                                    default:
                                        v4 = v4
                                    }
                                    fmt.Println(v2, v92)
                                    v4 = (v4 + v4)
                                    _ = fi91
                                    _ = v92
                                }
                            }
                            _ = fi63
                        }
                        v6 -= v6
                        _ = v62
                    }
                case "hello":
                    v94 := GenS1{F0: 0, F1: true}
                    _ = v94
                default:
                    if true {
                        if false {
                            switch v4 {
                            case "":
                                fmt.Println(v5)
                            case "a":
                                if (!(2 == 0)) {
                                    v4 = v4
                                    v2 = (v2 % ((0) | 1))
                                    fmt.Println(v5, v1, v5)
                                } else {
                                    var v95 uint = 2
                                    for fi96 := 0; fi96 < 1; fi96++ {
                                        v95 += (v95 | (v95 >> ((v95) & 63)))
                                        for fi97 := 0; fi97 < 20; fi97++ {
                                            v1 += v6
                                            var v98 uint = 1
                                            _ = len(v3)
                                            _ = fi97
                                            _ = v98
                                        }
                                        _ = fi96
                                    }
                                    v99 := GenS12{F0: v5, F1: 4294967295}
                                    _ = v95
                                    _ = v99
                                }
                            default:
                                if true {
                                    if true {
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                        var v100 uint8 = 132
                                        _ = v100
                                    }
                                    for fi101 := 0; fi101 < 17; fi101++ {
                                        var v102 int = (fi101 / ((fi101) | 1))
                                        v3 = append(v3, 532)
                                        v102 -= (v102 * 32768)
                                        _ = fi101
                                        _ = v102
                                    }
                                } else {
                                    for fi103 := 0; fi103 < 15; fi103++ {
                                        v5 = (^(v5))
                                        _ = fi103
                                    }
                                    if false {
                                        if ((false || (true || true)) && (255 == 65535)) {
                                            for fi104 := 0; fi104 < 10; fi104++ {
                                                switch ("foo bar" + v4) {
                                                case "foo bar":
                                                    if len(v3) > 0 {
                                                        v3[0] = 2
                                                    }
                                                case "hello":
                                                    if false {
                                                        v1 = -2.25
                                                    } else {
                                                        var v105 string = ("a" + v4)
                                                        _ = v105
                                                    }
                                                case "a":
                                                    _ = len(v3)
                                                }
                                                fmt.Println(v18)
                                                _ = fi104
                                            }
                                        } else {
                                            var v106 uint32 = 2
                                            switch "foo bar" {
                                            case "a":
                                                for fi107 := 0; fi107 < 1; fi107++ {
                                                    switch fi107 {
                                                    case 32768:
                                                        var v108 int16 = 42
                                                        _ = v108
                                                    case 895:
                                                        v5 = v5
                                                    }
                                                    var v109 int8 = 1
                                                    v4 = ""
                                                    _ = fi107
                                                    _ = v109
                                                }
                                            case "hello":
                                                switch ("\n\t" + v4) {
                                                case "a":
                                                    if len(v3) >= 2 {
                                                        v3 = v3[1:len(v3)]
                                                    }
                                                case "\n\t":
                                                    v110 := GenS13{F0: 1}
                                                    _ = v110
                                                default:
                                                    fmt.Println(v1)
                                                }
                                            case "\n\t":
                                                fmt.Println(v6, v4)
                                            }
                                            _ = v106
                                        }
                                        var v111 int16 = 32767
                                        _ = v111
                                    } else {
                                        var v112 uint64 = 991
                                        switch v4 {
                                        case "":
                                            var v113 int = 128
                                            _ = v113
                                        case "foo bar":
                                            for fi114 := 0; fi114 < 14; fi114++ {
                                                v2 = (^(v2))
                                                if (!false) {
                                                    var v115 int = (fi114 * fi114)
                                                    _ = v115
                                                } else {
                                                    fmt.Println(v4, v2, fi114)
                                                    for fi116 := 0; fi116 < 11; fi116++ {
                                                        fmt.Println(fi116, v18)
                                                        if (3.14159 <= v18) {
                                                            fmt.Println(v5, v18)
                                                        } else {
                                                            var v117 uint64 = 126
                                                            _ = v117
                                                        }
                                                        if false {
                                                            v2 = (-(v2))
                                                            var v118 int64 = int64(v2)
                                                            _ = v118
                                                        } else {
                                                            var v119 uint8 = uint8(v112)
                                                            if ((1e10 - v18) < v18) {
                                                                v120 := GenS9{F0: v119, F1: v2, F2: (v119 % (((v119 | (v119 >> ((65535) & 63)))) | 1))}
                                                                fmt.Println(fi116, v112)
                                                                _ = len(v3)
                                                                _ = v120
                                                            }
                                                            _ = v119
                                                        }
                                                        _ = fi116
                                                    }
                                                }
                                                _ = fi114
                                            }
                                        case "a":
                                            var v121 uint8 = 2
                                            _ = v121
                                        }
                                        _ = v112
                                    }
                                }
                            }
                            var v122 string = (v4 + "\n\t")
                            if len(v3) > 0 {
                                v3[0] = 4294967295
                            }
                            _ = v122
                        }
                    } else {
                        if len(v3) > 0 {
                            v3[0] = 0
                        }
                    }
                }
                v5 -= -9223372036854775808
            } else {
                var v123 int32 = v2
                v1 += v6
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                _ = v123
            }
        }
        fmt.Println(v2, v5)
    } else {
        v3 = append(v3, 65535)
    }
    var v124 uint64 = 4294967295
    var v125 uint = 2
    for fi126 := 0; fi126 < 9; fi126++ {
        var v127 int = -844
        v125 = (v125 ^ (^(v125)))
        v127 -= -174
        _ = fi126
        _ = v127
    }
    var v128 uint64 = (v124 % ((378) | 1))
    for fi129 := 0; fi129 < 19; fi129++ {
        for fi130 := 0; fi130 < 12; fi130++ {
            if (v6 >= v1) {
                var v131 int16 = 32767
                _ = v131
            }
            var v132 int64 = v5
            var v133 uint64 = (v124 * 42)
            _ = fi130
            _ = v132
            _ = v133
        }
        v18 = float64(v6)
        _ = fi129
    }
    if true {
        v125 = v125
    }
    switch 0 {
    case -661:
        switch v4 {
        case "a":
            if false {
                v5 = int64(v128)
                v124 += (v124 * (v128 >> ((10) & 63)))
                fmt.Println(v6, v6, v18)
            }
        default:
            v124 = v124
        }
    }
    v134 := GenS14{F0: -464, F1: 332, F2: 0}
    var v135 string = ("hello" + v4)
    fmt.Println("v1:", v1)
    fmt.Println("v124:", v124)
    fmt.Println("v125:", v125)
    fmt.Println("v128:", v128)
    fmt.Println("v135:", v135)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v124
    _ = v125
    _ = v128
    _ = v134
    _ = v135
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
