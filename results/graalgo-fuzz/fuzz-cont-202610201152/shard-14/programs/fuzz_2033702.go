package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint8; F1 int64; F2 int }
type GenS3 struct { F0 int; F1 uint }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint64; F1 int8; F2 int8 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 uint16; F1 int64; F2 bool }
type GenS8 struct { F0 uint8; F1 int; F2 int64 }
type GenS9 struct { F0 int16; F1 uint; F2 int8 }
type GenS10 struct { F0 string }
type GenS11 struct { F0 string; F1 uint8 }
type GenS12 struct { F0 int64; F1 uint16 }
type GenS13 struct { F0 uint16; F1 uint16 }
type GenS14 struct { F0 int16 }
type GenS15 struct { F0 int64; F1 int8 }
type GenS16 struct { F0 int32; F1 int64 }
type GenS17 struct { F0 int8; F1 string }
type GenS18 struct { F0 int16 }
type GenS19 struct { F0 string; F1 int32; F2 uint16 }
type GenS20 struct { F0 int16; F1 string }
type GenS21 struct { F0 int16 }
type GenS22 struct { F0 int64; F1 int }
type GenS23 struct { F0 int64; F1 int; F2 bool }

func main() {
    v1 := make(map[string]int)
    v2 := make(map[string]int)
    var v3 uint16 = 0
    v4 := make(map[int]int)
    var v5 float64 = -2.25
    switch 128 {
    case 32767:
        for fi6 := 0; fi6 < 2; fi6++ {
            if (2 != 301) {
                v7 := []int{(fi6 % ((fi6) | 1))}
                _ = v7
            } else {
                var v8 int16 = -386
                _ = v8
            }
            _ = fi6
        }
    case 128:
        var v9 int8 = 1
        _ = v9
    case 838:
        switch "foo bar" {
        case "\n\t":
            if (!(0 < 42)) {
                var v10 uint64 = 1
                if false {
                    if (1 >= -1) {
                        v11 := []int{128, 994}
                        if len(v11) >= 2 {
                            v11 = v11[1:len(v11)]
                        }
                        for fi12 := 0; fi12 < 12; fi12++ {
                            v5 = ((v5 * v5) - float64(0.0))
                            _ = fi12
                        }
                        _ = v11
                    } else {
                        v4[-2147483648] = 9223372036854775807
                    }
                }
                _ = v10
            }
        default:
            v13 := GenS1{F0: 32767}
            _ = v13
        }
    }
    v3 = (v3 & (^(v3)))
    v3 = v3
    v5 = v5
    var v14 string = ""
    if (1 >= 248) {
        v3 += (^(v3))
        if false {
            fmt.Println(v14, v14, v14)
            fmt.Println(v14, v5, v5)
            for fi15 := 0; fi15 < 18; fi15++ {
                v16 := []int{fi15, fi15, 32768}
                switch "a" {
                case "a":
                    if (68 > uint32(v3)) {
                        v17 := GenS1{F0: -1}
                        v5 = (v5 + (((v5 + ((v5 - (((float64(-2.25) - (v5 / (((3.14159 * 0.0) + v5) * ((v5 + v5) * 0.1)))) + (v5 - v5)) - (v5 * 1.5))) - (float64(0.1) - -2.25))) * float64(0.1)) * ((-2.25 / (v5 / float64(1.5))) / (v5 + v5))))
                        _ = v17
                    }
                case "hello":
                    fmt.Println(fi15)
                case "\n\t":
                    for fi18 := 0; fi18 < 6; fi18++ {
                        var v19 uint8 = 0
                        var v20 float64 = (float64(-2.25) * (v5 * (3.14159 / 3.14159)))
                        if len(v16) > 0 {
                            v16[0] = (fi18 & (fi15 * fi15))
                        }
                        _ = fi18
                        _ = v19
                        _ = v20
                    }
                }
                _ = fi15
                _ = v16
            }
        }
        var v21 uint64 = 296
        _ = v21
    } else {
        if false {
            v22 := GenS1{F0: 1000}
            if false {
                for fi23 := 0; fi23 < 16; fi23++ {
                    if true {
                        v5 += v5
                        fmt.Println(v5, fi23)
                        v14 = ""
                    } else {
                        switch fi23 {
                        case 2147483647:
                            v24 := GenS1{F0: 1000}
                            _ = v24
                        default:
                            var v25 uint8 = 127
                            _ = v25
                        }
                        v2[v14] = 4294967295
                    }
                    fmt.Println(v14, v3, v5)
                    fmt.Println(v3)
                    _ = fi23
                }
                switch -1 {
                case 1:
                    var v26 string = v14
                    _ = v26
                default:
                    _ = len(v4)
                }
            }
            _ = v22
        }
        var v27 int64 = -9223372036854775808
        fmt.Println(v14, v3, v3)
        _ = v27
    }
    v28 := []int{1, 0, -383, 1}
    var v29 uint8 = 22
    if ((32767 != 907) || true) {
        var v30 bool = (!false)
        if (v29 <= v29) {
            for fi31 := 0; fi31 < 2; fi31++ {
                if v30 {
                    v32 := GenS1{F0: 42}
                    if (127 < v29) {
                        var v33 int64 = 0
                        v34 := GenS2{F0: 238, F1: v33, F2: fi31}
                        fmt.Println(v33, v3)
                        _ = v33
                        _ = v34
                    } else {
                        v3 = 42
                        if false {
                            if (639 == 0) {
                                fmt.Println(fi31)
                                if v30 {
                                    for fi35 := 0; fi35 < 2; fi35++ {
                                        var v36 uint16 = (v3 >> ((238) & 63))
                                        v1[("" + v14)] = int(v36)
                                        fmt.Println(v5, v30)
                                        _ = fi35
                                        _ = v36
                                    }
                                    var v37 int8 = -128
                                    _ = v37
                                } else {
                                    for fi38 := 0; fi38 < 10; fi38++ {
                                        switch (v14 + v14) {
                                        case "a":
                                            if v30 {
                                                fmt.Println(v29, v3, v3)
                                            } else {
                                                var v39 int8 = 100
                                                v39 = 46
                                                _ = v39
                                            }
                                        case "foo bar":
                                            _ = len(v28)
                                        case "":
                                            v40 := GenS3{F0: fi31, F1: 255}
                                            _ = v40
                                        default:
                                            v41 := GenS4{F0: 2}
                                            _ = v41
                                        }
                                        var v42 int32 = -2147483648
                                        v29 -= (v29 - uint8(v3))
                                        _ = fi38
                                        _ = v42
                                    }
                                    v30 = false
                                    for fi43 := 0; fi43 < 19; fi43++ {
                                        fmt.Println(v14, v29)
                                        var v44 uint16 = 551
                                        _ = fi43
                                        _ = v44
                                    }
                                }
                                if (-37 < -83) {
                                    var v45 int8 = 127
                                    _ = v45
                                } else {
                                    var v46 uint32 = 0
                                    switch v14 {
                                    case "":
                                        var v47 uint16 = (v3 << ((65535) & 63))
                                        _ = v47
                                    default:
                                        var v48 int = (^(fi31))
                                        _ = v48
                                    }
                                    _ = v46
                                }
                            } else {
                                for fi49 := 0; fi49 < 20; fi49++ {
                                    delete(v4, (^(fi49)))
                                    _ = fi49
                                }
                            }
                        } else {
                            var v50 int = (fi31 & fi31)
                            if (v30 || ((!v30) && false)) {
                                switch v50 {
                                case 128:
                                    fmt.Println(v29, v5, v29)
                                default:
                                    var v51 uint = 599
                                    _ = v51
                                }
                            } else {
                                var v52 int16 = -32768
                                _ = v52
                            }
                            if (!v30) {
                                if (((42 < -32768) || false) || (42 < 651)) {
                                    _ = len(v1)
                                    v53 := GenS5{F0: 4294967295, F1: 1, F2: 42}
                                    for fi54 := 0; fi54 < 6; fi54++ {
                                        if v30 {
                                            fmt.Println(v50)
                                            if ((!false) && (42 != 339)) {
                                                var v55 uint64 = 2
                                                fmt.Println(v55, v55, v14)
                                                _ = v55
                                            }
                                            v56 := GenS6{F0: int32(v3)}
                                            _ = v56
                                        }
                                        v5 = 1e10
                                        _ = len(v28)
                                        _ = fi54
                                    }
                                    _ = v53
                                } else {
                                    v14 = (v14 + ((((("a" + "") + ("hello" + v14)) + ("" + v14)) + "") + v14))
                                    for fi57 := 0; fi57 < 7; fi57++ {
                                        v58 := GenS3{F0: fi31, F1: 684}
                                        var v59 float32 = float32(1e10)
                                        v14 = v14
                                        var v60 float64 = (float64(v59) - ((0.0 + (v5 + v5)) + (v5 - (v5 * v5))))
                                        _ = fi57
                                        _ = v58
                                        _ = v59
                                        _ = v60
                                    }
                                    v61 := GenS5{F0: 18446744073709551615, F1: 0, F2: 100}
                                    _ = v61
                                }
                                v62 := v1["foo bar"]
                                if len(v28) > 0 {
                                    v28[0] = -9223372036854775808
                                }
                                _ = v62
                            } else {
                                fmt.Println(v14, v30, v3)
                                if false {
                                    var v63 int64 = 328
                                    var v64 uint32 = 538
                                    v64 = 472
                                    _ = v63
                                    _ = v64
                                } else {
                                    var v65 int32 = 2147483647
                                    _ = len(v2)
                                    _ = v65
                                }
                            }
                            _ = v50
                        }
                    }
                    _ = v32
                } else {
                    var v66 float32 = -2.25
                    switch -1 {
                    case 771:
                        var v67 uint32 = 4294967295
                        _ = v67
                    default:
                        switch (fi31 % (((fi31 ^ (fi31 - (fi31 >> ((135) & 63))))) | 1)) {
                        case 0:
                            for fi68 := 0; fi68 < 12; fi68++ {
                                switch (fi31 + -823) {
                                case -2147483648:
                                    _ = len(v2)
                                }
                                _ = fi68
                            }
                        }
                    }
                    switch 2147483647 {
                    case 127:
                        for fi69 := 0; fi69 < 8; fi69++ {
                            if len(v28) > 0 {
                                v28[0] = fi31
                            }
                            v70 := GenS1{F0: -1}
                            v29 -= v29
                            v30 = true
                            _ = fi69
                            _ = v70
                        }
                    case 989:
                        v29 -= v29
                    }
                    _ = v66
                }
                _ = fi31
            }
            if len(v28) >= 2 {
                v28 = v28[1:len(v28)]
            }
            if false {
                v71 := v4[-708]
                v30 = v30
                _ = v71
            }
        }
        if v30 {
            var v72 uint16 = (v3 % (((^(v3))) | 1))
            _ = v72
        } else {
            if ((1e10 * 0.1) <= 0.1) {
                switch v14 {
                case "\n\t":
                    if false {
                        if len(v28) > 0 {
                            v28[0] = 9223372036854775807
                        }
                        _ = len(v2)
                        if ((v14 + (v14 + v14)) >= v14) {
                            if true {
                                for fi73 := 0; fi73 < 3; fi73++ {
                                    if ((0.0 < (0.1 + 0.1)) || ((float32(((v5 * v5) + v5)) / float32(3.14159)) == ((((1e10 / 1.0) + (1.5 / 1e10)) / (-2.25 - (3.14159 * ((float32((v5 - 0.0)) / 1.0) / (-2.25 + 0.1))))) / 1e10))) {
                                        delete(v1, "foo bar")
                                    } else {
                                        v29 += (^(v29))
                                        v29 -= 42
                                        v74 := v4[42]
                                        _ = v74
                                    }
                                    if (18446744073709551615 == 2) {
                                        if v30 {
                                            if false {
                                                v75 := GenS2{F0: v29, F1: 186, F2: fi73}
                                                var v76 bool = (2 >= 730)
                                                _ = v75
                                                _ = v76
                                            } else {
                                                if v30 {
                                                    if (-9223372036854775808 >= -525) {
                                                        delete(v1, v14)
                                                        v30 = false
                                                    } else {
                                                        if (!(0.1 != (0.0 + 0.0))) {
                                                            if v30 {
                                                                if v30 {
                                                                    var v77 uint = 255
                                                                    _ = v77
                                                                }
                                                            } else {
                                                                _ = len(v1)
                                                                if len(v28) >= 2 {
                                                                    v28 = v28[1:len(v28)]
                                                                }
                                                                var v78 uint = 930
                                                                _ = v78
                                                            }
                                                            var v79 uint = 4294967295
                                                            for fi80 := 0; fi80 < 19; fi80++ {
                                                                if len(v28) >= 2 {
                                                                    v28 = v28[1:len(v28)]
                                                                }
                                                                _ = fi80
                                                            }
                                                            _ = v79
                                                        }
                                                    }
                                                    switch 128 {
                                                    case 32768:
                                                        var v81 float64 = (-2.25 + (v5 * (0.1 / float64((0.1 + float32(((v5 / -2.25) - v5)))))))
                                                        _ = v81
                                                    default:
                                                        for fi82 := 0; fi82 < 14; fi82++ {
                                                            var v83 int8 = 1
                                                            v84 := GenS7{F0: (^(v3)), F1: 1, F2: true}
                                                            var v85 int16 = 42
                                                            _ = fi82
                                                            _ = v83
                                                            _ = v84
                                                            _ = v85
                                                        }
                                                    }
                                                    if len(v28) > 0 {
                                                        v28[0] = fi73
                                                    }
                                                } else {
                                                    var v86 float32 = float32((-2.25 - (v5 - v5)))
                                                    _ = v86
                                                }
                                                switch (fi73 + 1) {
                                                case -2147483648:
                                                    switch v14 {
                                                    case "\n\t":
                                                        if v30 {
                                                            var v87 float32 = 1e10
                                                            var v88 uint = 0
                                                            switch v14 {
                                                            case "":
                                                                v88 -= (v88 | 4294967295)
                                                            case "\n\t":
                                                                switch (fi73 - (fi73 << (((v88 & v88)) & 63))) {
                                                                case 32767:
                                                                    switch (((((v14 + (v14 + v14)) + (v14 + v14)) + "") + v14) + "\n\t") {
                                                                    case "":
                                                                        for fi89 := 0; fi89 < 19; fi89++ {
                                                                            var v90 int16 = 1000
                                                                            if v30 {
                                                                                delete(v4, (fi73 >> (((v88 - (v88 | v88))) & 63)))
                                                                                switch (fi73 - (^(fi89))) {
                                                                                case 32768:
                                                                                    var v91 int = (fi89 | fi89)
                                                                                    _ = v91
                                                                                default:
                                                                                    if v30 {
                                                                                        _ = len(v28)
                                                                                    } else {
                                                                                        var v92 int8 = -1
                                                                                        v14 = (v14 + v14)
                                                                                        if (9223372036854775807 != 9223372036854775807) {
                                                                                            v29 = v29
                                                                                        } else {
                                                                                            v3 += (v3 / (((v3 * v3)) | 1))
                                                                                        }
                                                                                        _ = v92
                                                                                    }
                                                                                }
                                                                                if v30 {
                                                                                    if (true || false) {
                                                                                        if false {
                                                                                            if ((-1 != -140) && true) {
                                                                                                var v93 float64 = v5
                                                                                                _ = v93
                                                                                            } else {
                                                                                                v14 = "hello"
                                                                                                switch 4294967295 {
                                                                                                case 849:
                                                                                                    fmt.Println(v90)
                                                                                                case 1:
                                                                                                    var v94 int8 = 40
                                                                                                    _ = v94
                                                                                                case 128:
                                                                                                    var v95 int64 = 9223372036854775807
                                                                                                    _ = v95
                                                                                                default:
                                                                                                    if len(v28) >= 2 {
                                                                                                        v28 = v28[1:len(v28)]
                                                                                                    }
                                                                                                }
                                                                                                v96 := GenS7{F0: 42, F1: 42, F2: false}
                                                                                                _ = v96
                                                                                            }
                                                                                            var v97 bool = (0 != 4294967295)
                                                                                            var v98 int64 = 9223372036854775807
                                                                                            _ = v97
                                                                                            _ = v98
                                                                                        } else {
                                                                                            switch (fi89 | (-(fi89))) {
                                                                                            case 32768:
                                                                                                var v99 int8 = 1
                                                                                                _ = v99
                                                                                            case -2147483648:
                                                                                                fmt.Println(v90, v88, v5)
                                                                                            case 1:
                                                                                                if v30 {
                                                                                                    if (v5 != (v5 - v5)) {
                                                                                                        v3 -= (v3 >> ((v88) & 63))
                                                                                                        switch v14 {
                                                                                                        case "foo bar":
                                                                                                            var v100 uint16 = uint16(v88)
                                                                                                            _ = v100
                                                                                                        default:
                                                                                                            var v101 int64 = -910
                                                                                                            var v102 uint32 = 4294967295
                                                                                                            _ = v101
                                                                                                            _ = v102
                                                                                                        }
                                                                                                    } else {
                                                                                                        v30 = (!v30)
                                                                                                    }
                                                                                                    fmt.Println(v29)
                                                                                                    v14 = v14
                                                                                                    switch (fi89 / (((fi73 & -611)) | 1)) {
                                                                                                    case 42:
                                                                                                        v103 := GenS7{F0: (v3 ^ v3), F1: -280, F2: (v30 || (v90 >= (-(v90))))}
                                                                                                        _ = v103
                                                                                                    default:
                                                                                                        v104 := GenS4{F0: v3}
                                                                                                        _ = v104
                                                                                                    }
                                                                                                    v90 = 1
                                                                                                } else {
                                                                                                    switch "\n\t" {
                                                                                                    case "foo bar":
                                                                                                        var v105 uint = (v88 >> ((v88) & 63))
                                                                                                        _ = v105
                                                                                                    case "":
                                                                                                        v106 := GenS8{F0: (v29 - (^(v29))), F1: -422, F2: -1}
                                                                                                        _ = v106
                                                                                                    default:
                                                                                                        var v107 int32 = 351
                                                                                                        _ = v107
                                                                                                    }
                                                                                                    var v108 float64 = v5
                                                                                                    _ = v108
                                                                                                }
                                                                                            default:
                                                                                                switch (v14 + v14) {
                                                                                                case "\n\t":
                                                                                                    var v109 int = fi89
                                                                                                    _ = v109
                                                                                                default:
                                                                                                    if v30 {
                                                                                                        switch ("foo bar" + v14) {
                                                                                                        case "foo bar":
                                                                                                            fmt.Println(v88, v87, v29)
                                                                                                        case "hello":
                                                                                                            var v110 uint16 = (v3 * 65535)
                                                                                                            _ = v110
                                                                                                        case "\n\t":
                                                                                                            if v30 {
                                                                                                                _ = len(v4)
                                                                                                                switch fi89 {
                                                                                                                case 127:
                                                                                                                    var v111 int8 = 100
                                                                                                                    _ = v111
                                                                                                                case -9223372036854775808:
                                                                                                                    if v30 {
                                                                                                                        switch (-(fi89)) {
                                                                                                                        case 110:
                                                                                                                            switch (v14 + (v14 + "\n\t")) {
                                                                                                                            case "hello":
                                                                                                                                fmt.Println(fi73, fi89)
                                                                                                                            case "foo bar":
                                                                                                                                v30 = true
                                                                                                                            case "":
                                                                                                                                var v112 float32 = v87
                                                                                                                                _ = v112
                                                                                                                            }
                                                                                                                        }
                                                                                                                        if len(v28) > 0 {
                                                                                                                            v28[0] = (fi73 + int(v90))
                                                                                                                        }
                                                                                                                    } else {
                                                                                                                        if true {
                                                                                                                            var v113 uint = (v88 - v88)
                                                                                                                            _ = v113
                                                                                                                        }
                                                                                                                    }
                                                                                                                    if false {
                                                                                                                        _ = len(v28)
                                                                                                                        v114 := GenS1{F0: v90}
                                                                                                                        switch v14 {
                                                                                                                        case "a":
                                                                                                                            fmt.Println(fi89)
                                                                                                                        default:
                                                                                                                            v5 -= v5
                                                                                                                        }
                                                                                                                        _ = v114
                                                                                                                    }
                                                                                                                default:
                                                                                                                    fmt.Println(v88, fi89, fi89)
                                                                                                                }
                                                                                                            }
                                                                                                        }
                                                                                                        switch (fi73 & (-(fi73))) {
                                                                                                        case 32767:
                                                                                                            v115 := v2["\n\t"]
                                                                                                            _ = v115
                                                                                                        }
                                                                                                        switch 687 {
                                                                                                        case 757:
                                                                                                            delete(v4, -872)
                                                                                                        case -886:
                                                                                                            fmt.Println(v3, v29, v29)
                                                                                                        default:
                                                                                                            v30 = (((-1 < 667) || (float64(v87) >= 3.14159)) && v30)
                                                                                                        }
                                                                                                    }
                                                                                                }
                                                                                            }
                                                                                            switch ("foo bar" + "foo bar") {
                                                                                            case "\n\t":
                                                                                                if (((594 <= 0) || v30) || true) {
                                                                                                    fmt.Println(v90)
                                                                                                } else {
                                                                                                    if (v30 || true) {
                                                                                                        v5 = float64((v87 / 3.14159))
                                                                                                    }
                                                                                                    if (-1 > -33) {
                                                                                                        v29 = (v29 >> ((v88) & 63))
                                                                                                        if len(v28) >= 2 {
                                                                                                            v28 = v28[1:len(v28)]
                                                                                                        }
                                                                                                        v88 = (v88 | (v88 - (v88 % ((42) | 1))))
                                                                                                    } else {
                                                                                                        var v116 int = (fi73 + 127)
                                                                                                        switch ((v14 + v14) + (v14 + v14)) {
                                                                                                        case "foo bar":
                                                                                                            v88 = v88
                                                                                                        default:
                                                                                                            if false {
                                                                                                                v117 := GenS8{F0: (v29 << ((v88) & 63)), F1: (-(fi89)), F2: -195}
                                                                                                                _ = v117
                                                                                                            } else {
                                                                                                                v118 := GenS1{F0: v90}
                                                                                                                if v30 {
                                                                                                                    var v119 uint32 = 65535
                                                                                                                    var v120 int8 = 42
                                                                                                                    _ = v119
                                                                                                                    _ = v120
                                                                                                                } else {
                                                                                                                    if (!v30) {
                                                                                                                        if v30 {
                                                                                                                            switch v14 {
                                                                                                                            case "a":
                                                                                                                                v2[(((("hello" + (v14 + ("\n\t" + ("foo bar" + (v14 + (v14 + v14)))))) + "") + ((v14 + v14) + v14)) + v14)] = (^(v116))
                                                                                                                            case "hello":
                                                                                                                                fmt.Println(fi89, v5)
                                                                                                                            }
                                                                                                                        }
                                                                                                                        if len(v28) > 0 {
                                                                                                                            v28[0] = (fi73 * fi73)
                                                                                                                        }
                                                                                                                        var v121 float32 = 1.5
                                                                                                                        _ = v121
                                                                                                                    } else {
                                                                                                                        if (!((((!false) || ((668 == 1) || ((true && (v30 && v30)) && (0.0 != (v5 / v5))))) && ((fi89 ^ (fi89 & (^(v116)))) >= (^(fi73)))) && (!v30))) {
                                                                                                                            _ = len(v28)
                                                                                                                            v122 := GenS1{F0: v90}
                                                                                                                            switch (^(v116)) {
                                                                                                                            case -2147483648:
                                                                                                                                fmt.Println(v90, v87)
                                                                                                                            case 551:
                                                                                                                                if ((v3 * (v3 - (v3 ^ (v3 ^ (v3 >> (((v88 & 42)) & 63)))))) != (v3 / (((v3 - v3)) | 1))) {
                                                                                                                                    var v123 int16 = (v90 & (v90 + v90))
                                                                                                                                    _ = v123
                                                                                                                                } else {
                                                                                                                                    var v124 int16 = v90
                                                                                                                                    var v125 uint16 = (v3 - 0)
                                                                                                                                    _ = v124
                                                                                                                                    _ = v125
                                                                                                                                }
                                                                                                                            default:
                                                                                                                                v126 := GenS3{F0: v116, F1: (v88 + (^(v88)))}
                                                                                                                                _ = v126
                                                                                                                            }
                                                                                                                            _ = v122
                                                                                                                        } else {
                                                                                                                            fmt.Println(v14, fi89)
                                                                                                                            if (v5 == float64(((v87 * ((((v87 * (v87 * v87)) - v87) + (float32((v5 * v5)) / (v87 * (v87 + float32(v5))))) * v87)) + 1.5))) {
                                                                                                                                v127 := GenS9{F0: v90, F1: (v88 + v88), F2: 42}
                                                                                                                                if (!(-446 > 0)) {
                                                                                                                                    switch ((v14 + v14) + v14) {
                                                                                                                                    case "\n\t":
                                                                                                                                        if true {
                                                                                                                                            _ = len(v28)
                                                                                                                                            v90 += (v90 - v90)
                                                                                                                                        }
                                                                                                                                    case "":
                                                                                                                                        v128 := v2[v14]
                                                                                                                                        _ = v128
                                                                                                                                    default:
                                                                                                                                        if (v30 && (0 != 661)) {
                                                                                                                                            v116 -= v116
                                                                                                                                            v129 := GenS6{F0: 42}
                                                                                                                                            switch 32767 {
                                                                                                                                            case 32768:
                                                                                                                                                v130 := GenS4{F0: v3}
                                                                                                                                                _ = v130
                                                                                                                                            default:
                                                                                                                                                v3 = (v3 | (v3 | (v3 % ((v3) | 1))))
                                                                                                                                            }
                                                                                                                                            _ = v129
                                                                                                                                        } else {
                                                                                                                                            fmt.Println(fi89, v14)
                                                                                                                                            v87 = 0.1
                                                                                                                                        }
                                                                                                                                    }
                                                                                                                                    v14 = "hello"
                                                                                                                                    var v131 bool = (!(18446744073709551615 == uint64(v29)))
                                                                                                                                    _ = v131
                                                                                                                                }
                                                                                                                                _ = v127
                                                                                                                            } else {
                                                                                                                                switch v14 {
                                                                                                                                case "foo bar":
                                                                                                                                    var v132 float64 = v5
                                                                                                                                    _ = v132
                                                                                                                                }
                                                                                                                            }
                                                                                                                        }
                                                                                                                        v133 := GenS10{F0: v14}
                                                                                                                        _ = v133
                                                                                                                    }
                                                                                                                    var v134 int = v116
                                                                                                                    _ = v134
                                                                                                                }
                                                                                                                if (!(uint16(v29) > (v3 ^ (v3 >> ((v88) & 63))))) {
                                                                                                                    if v30 {
                                                                                                                        v135 := GenS10{F0: (v14 + ("hello" + v14))}
                                                                                                                        v5 -= ((((v5 - (v5 / v5)) / -2.25) - float64(v87)) - (0.1 / ((1e10 - v5) + (0.1 - v5))))
                                                                                                                        _ = v135
                                                                                                                    } else {
                                                                                                                        if false {
                                                                                                                            v136 := GenS11{F0: "a", F1: (v29 - v29)}
                                                                                                                            _ = v136
                                                                                                                        } else {
                                                                                                                            if (0 > -1) {
                                                                                                                                var v137 bool = true
                                                                                                                                v88 = 1
                                                                                                                                _ = v137
                                                                                                                            }
                                                                                                                        }
                                                                                                                        if (!v30) {
                                                                                                                            v3 -= 255
                                                                                                                            v138 := GenS12{F0: -9223372036854775808, F1: (v3 >> (((v88 ^ v88)) & 63))}
                                                                                                                            var v139 int = v116
                                                                                                                            fmt.Println(v87, v139, v30)
                                                                                                                            _ = len(v28)
                                                                                                                            if v30 {
                                                                                                                                if (v3 == 32767) {
                                                                                                                                    v140 := GenS13{F0: (v3 | (v3 ^ uint16(v88))), F1: 338}
                                                                                                                                    var v141 uint = (v88 % ((18446744073709551615) | 1))
                                                                                                                                    _ = v140
                                                                                                                                    _ = v141
                                                                                                                                } else {
                                                                                                                                    v90 = v90
                                                                                                                                    var v142 bool = ((v14 <= v14) || v30)
                                                                                                                                    if true {
                                                                                                                                        var v143 uint16 = (^(v3))
                                                                                                                                        _ = v143
                                                                                                                                    } else {
                                                                                                                                        var v144 uint8 = v29
                                                                                                                                        _ = v144
                                                                                                                                    }
                                                                                                                                    _ = v142
                                                                                                                                }
                                                                                                                            } else {
                                                                                                                                fmt.Println(v90)
                                                                                                                                switch "" {
                                                                                                                                case "hello":
                                                                                                                                    var v145 uint64 = 18446744073709551615
                                                                                                                                    _ = v145
                                                                                                                                case "foo bar":
                                                                                                                                    v5 = v5
                                                                                                                                case "\n\t":
                                                                                                                                    var v146 int8 = int8(v88)
                                                                                                                                    _ = v146
                                                                                                                                default:
                                                                                                                                    v28 = append(v28, fi73)
                                                                                                                                }
                                                                                                                            }
                                                                                                                            _ = v138
                                                                                                                            _ = v139
                                                                                                                        }
                                                                                                                        v90 = v90
                                                                                                                        v87 -= v87
                                                                                                                    }
                                                                                                                    var v147 int32 = 0
                                                                                                                    v87 = -2.25
                                                                                                                    _ = v147
                                                                                                                } else {
                                                                                                                    v28 = append(v28, (-(fi73)))
                                                                                                                    if ((!(114 != 4294967295)) && true) {
                                                                                                                        var v148 int16 = -32768
                                                                                                                        v149 := GenS14{F0: -688}
                                                                                                                        var v150 bool = false
                                                                                                                        var v151 int16 = v148
                                                                                                                        v90 -= (v148 / (((v151 % ((211) | 1))) | 1))
                                                                                                                        _ = v148
                                                                                                                        _ = v149
                                                                                                                        _ = v150
                                                                                                                        _ = v151
                                                                                                                    } else {
                                                                                                                        if v30 {
                                                                                                                            if v30 {
                                                                                                                                if false {
                                                                                                                                    v152 := v4[(fi73 << ((v88) & 63))]
                                                                                                                                    var v153 float32 = (float32(1.5) + ((v87 * (-2.25 * (((0.0 - (v87 + v87)) + -2.25) * (v87 - float32(v5))))) + (((float32(-2.25) / 1.0) - 3.14159) + 1e10)))
                                                                                                                                    _ = len(v28)
                                                                                                                                    _ = v152
                                                                                                                                    _ = v153
                                                                                                                                }
                                                                                                                                if len(v28) >= 2 {
                                                                                                                                    v28 = v28[1:len(v28)]
                                                                                                                                }
                                                                                                                            }
                                                                                                                            if (-1 != (fi89 >> (((^(v88))) & 63))) {
                                                                                                                                switch "" {
                                                                                                                                case "a":
                                                                                                                                    v154 := GenS2{F0: v29, F1: -741, F2: (v116 << ((v88) & 63))}
                                                                                                                                    _ = v154
                                                                                                                                case "hello":
                                                                                                                                    v155 := GenS4{F0: 32767}
                                                                                                                                    _ = v155
                                                                                                                                case "\n\t":
                                                                                                                                    var v156 uint16 = (v3 | 0)
                                                                                                                                    _ = v156
                                                                                                                                }
                                                                                                                                var v157 uint64 = 41
                                                                                                                                _ = v157
                                                                                                                            } else {
                                                                                                                                var v158 string = v14
                                                                                                                                switch fi89 {
                                                                                                                                case 4294967295:
                                                                                                                                    var v159 uint = v88
                                                                                                                                    _ = v159
                                                                                                                                case 27:
                                                                                                                                    if len(v28) > 0 {
                                                                                                                                        v28[0] = fi73
                                                                                                                                    }
                                                                                                                                    v3 = 386
                                                                                                                                    var v160 int8 = -1
                                                                                                                                    _ = v160
                                                                                                                                }
                                                                                                                                v158 = (v158 + ("" + "foo bar"))
                                                                                                                                _ = v158
                                                                                                                            }
                                                                                                                            if true {
                                                                                                                                fmt.Println(v3, v87)
                                                                                                                                v90 = 32767
                                                                                                                                fmt.Println(v90)
                                                                                                                            }
                                                                                                                        } else {
                                                                                                                            v161 := GenS15{F0: -9223372036854775808, F1: -39}
                                                                                                                            v162 := GenS2{F0: 249, F1: 42, F2: fi73}
                                                                                                                            var v163 uint32 = 0
                                                                                                                            _ = v161
                                                                                                                            _ = v162
                                                                                                                            _ = v163
                                                                                                                        }
                                                                                                                        v88 = (v88 >> ((4294967295) & 63))
                                                                                                                        switch (fi89 * fi73) {
                                                                                                                        case 4294967295:
                                                                                                                            fmt.Println(v3, v30)
                                                                                                                        case 276:
                                                                                                                            v28 = append(v28, 32768)
                                                                                                                        case -2147483648:
                                                                                                                            delete(v2, v14)
                                                                                                                        default:
                                                                                                                            var v164 int32 = 787
                                                                                                                            _ = v164
                                                                                                                        }
                                                                                                                    }
                                                                                                                    switch "hello" {
                                                                                                                    case "hello":
                                                                                                                        v3 = 42
                                                                                                                    }
                                                                                                                }
                                                                                                                v165 := GenS16{F0: -892, F1: 881}
                                                                                                                var v166 float32 = ((3.14159 + float32(0.1)) * (0.1 - float32((v5 * v5))))
                                                                                                                var v167 uint8 = (^(v29))
                                                                                                                v87 = v87
                                                                                                                if (!(!v30)) {
                                                                                                                    var v168 float32 = 3.14159
                                                                                                                    _ = v168
                                                                                                                } else {
                                                                                                                    var v169 uint8 = 128
                                                                                                                    fmt.Println(fi73, v30)
                                                                                                                    _ = v169
                                                                                                                }
                                                                                                                _ = v118
                                                                                                                _ = v165
                                                                                                                _ = v166
                                                                                                                _ = v167
                                                                                                            }
                                                                                                        }
                                                                                                        var v170 int = fi73
                                                                                                        _ = v116
                                                                                                        _ = v170
                                                                                                    }
                                                                                                }
                                                                                            case "foo bar":
                                                                                                var v171 float64 = 0.1
                                                                                                _ = v171
                                                                                            default:
                                                                                                v172 := GenS17{F0: -34, F1: v14}
                                                                                                _ = v172
                                                                                            }
                                                                                        }
                                                                                        v14 = (("a" + ("a" + "a")) + v14)
                                                                                    } else {
                                                                                        var v173 uint = (v88 % (((v88 | v88)) | 1))
                                                                                        var v174 bool = ((937 > 2147483647) || (!false))
                                                                                        _ = v173
                                                                                        _ = v174
                                                                                    }
                                                                                    var v175 uint64 = 734
                                                                                    fmt.Println(v90, v90)
                                                                                    _ = v175
                                                                                } else {
                                                                                    if (-128 <= 1) {
                                                                                        v176 := GenS18{F0: v90}
                                                                                        v177 := GenS16{F0: 1, F1: 474}
                                                                                        v178 := v4[(fi89 << (((v88 * (v88 * v88))) & 63))]
                                                                                        if false {
                                                                                            v178 += 9223372036854775807
                                                                                            v87 = ((v87 - v87) / (1.5 * 0.1))
                                                                                            v179 := GenS19{F0: ((v14 + v14) + "hello"), F1: 2147483647, F2: (v3 - uint16(v88))}
                                                                                            _ = v179
                                                                                        } else {
                                                                                            var v180 uint16 = uint16(v29)
                                                                                            var v181 int64 = 42
                                                                                            _ = v180
                                                                                            _ = v181
                                                                                        }
                                                                                        if ((!v30) || v30) {
                                                                                            if len(v28) > 0 {
                                                                                                v28[0] = -1
                                                                                            }
                                                                                            var v182 uint8 = (^(v29))
                                                                                            _ = v182
                                                                                        } else {
                                                                                            var v183 int8 = -1
                                                                                            _ = v183
                                                                                        }
                                                                                        _ = v176
                                                                                        _ = v177
                                                                                        _ = v178
                                                                                    } else {
                                                                                        v14 = "a"
                                                                                    }
                                                                                    if false {
                                                                                        v4[fi89] = (fi73 << ((v88) & 63))
                                                                                        fmt.Println(v3, v88, fi73)
                                                                                    } else {
                                                                                        v14 = "hello"
                                                                                        v3 = (^(v3))
                                                                                        fmt.Println(v87, v5)
                                                                                    }
                                                                                    fmt.Println(v14)
                                                                                    var v184 float32 = v87
                                                                                    _ = v184
                                                                                }
                                                                            }
                                                                            _ = fi89
                                                                            _ = v90
                                                                        }
                                                                    }
                                                                case 0:
                                                                    v185 := GenS4{F0: v3}
                                                                    _ = v185
                                                                default:
                                                                    v5 += -2.25
                                                                }
                                                            default:
                                                                v30 = false
                                                            }
                                                            _ = v87
                                                            _ = v88
                                                        }
                                                    case "":
                                                        v186 := GenS20{F0: 1, F1: ((((v14 + v14) + v14) + (v14 + "")) + v14)}
                                                        _ = v186
                                                    default:
                                                        if len(v28) > 0 {
                                                            v28[0] = (fi73 & (fi73 << ((1) & 63)))
                                                        }
                                                    }
                                                case -906:
                                                    v5 = (float64(float32(1.5)) / v5)
                                                case -1:
                                                    v187 := GenS11{F0: v14, F1: v29}
                                                    _ = v187
                                                }
                                                if v30 {
                                                    fmt.Println(v14, v3, v3)
                                                    var v188 bool = v30
                                                    v14 = (v14 + "\n\t")
                                                    _ = v188
                                                }
                                            }
                                            for fi189 := 0; fi189 < 15; fi189++ {
                                                v190 := GenS3{F0: fi189, F1: 2}
                                                fmt.Println(v3)
                                                var v191 uint32 = 4294967295
                                                _ = fi189
                                                _ = v190
                                                _ = v191
                                            }
                                            fmt.Println(v3)
                                        } else {
                                            v5 -= ((3.14159 * 0.0) + v5)
                                        }
                                        v14 = ""
                                    } else {
                                        v28 = append(v28, 1)
                                        v192 := GenS21{F0: -32768}
                                        var v193 uint16 = v3
                                        _ = v192
                                        _ = v193
                                    }
                                    if (255 > 2) {
                                        v30 = (v14 != (((v14 + ("\n\t" + "")) + v14) + v14))
                                        for fi194 := 0; fi194 < 1; fi194++ {
                                            if true {
                                                var v195 float32 = (-2.25 * 3.14159)
                                                fmt.Println(v3, v5, fi73)
                                                var v196 uint64 = 0
                                                _ = v195
                                                _ = v196
                                            }
                                            _ = fi194
                                        }
                                        switch fi73 {
                                        case 9223372036854775807:
                                            fmt.Println(v30, v29, v29)
                                        }
                                    }
                                    _ = fi73
                                }
                            } else {
                                if len(v28) >= 2 {
                                    v28 = v28[1:len(v28)]
                                }
                                if len(v28) >= 2 {
                                    v28 = v28[1:len(v28)]
                                }
                                if (1 == -1) {
                                    v29 = (^(v29))
                                    if len(v28) >= 2 {
                                        v28 = v28[1:len(v28)]
                                    }
                                } else {
                                    if len(v28) > 0 {
                                        v28[0] = 42
                                    }
                                }
                            }
                            fmt.Println(v3)
                        }
                    }
                default:
                    for fi197 := 0; fi197 < 17; fi197++ {
                        var v198 string = (((v14 + v14) + v14) + v14)
                        _ = fi197
                        _ = v198
                    }
                }
                v29 = uint8(v3)
                var v199 bool = true
                _ = v199
            } else {
                switch -172 {
                case 128:
                    if (false && (!(1 != uint64(v29)))) {
                        var v200 uint = 694
                        v201 := GenS17{F0: 27, F1: "\n\t"}
                        _ = v200
                        _ = v201
                    } else {
                        v202 := GenS6{F0: 433}
                        if len(v28) > 0 {
                            v28[0] = 665
                        }
                        _ = v202
                    }
                case 1:
                    var v203 uint = 255
                    _ = v203
                case -1:
                    if v30 {
                        for fi204 := 0; fi204 < 16; fi204++ {
                            if false {
                                fmt.Println(v3, v14, v5)
                            }
                            v3 -= v3
                            _ = fi204
                        }
                        for fi205 := 0; fi205 < 6; fi205++ {
                            delete(v2, (v14 + "\n\t"))
                            var v206 bool = true
                            for fi207 := 0; fi207 < 18; fi207++ {
                                var v208 bool = (!(966 <= -1))
                                switch int(v29) {
                                case -9223372036854775808:
                                    if v208 {
                                        v30 = v30
                                        v208 = ((v5 / float64(0.1)) > v5)
                                    } else {
                                        v5 = 1e10
                                        v4[9223372036854775807] = (-(fi207))
                                    }
                                case 42:
                                    switch v14 {
                                    case "hello":
                                        if v206 {
                                            v209 := GenS4{F0: v3}
                                            _ = v209
                                        } else {
                                            var v210 float64 = v5
                                            v14 = ("" + "\n\t")
                                            _ = v210
                                        }
                                    default:
                                        v28 = append(v28, (fi207 << ((18446744073709551615) & 63)))
                                        var v211 string = v14
                                        var v212 string = ("\n\t" + "\n\t")
                                        _ = v211
                                        _ = v212
                                    }
                                case 4294967295:
                                    var v213 bool = true
                                    _ = v213
                                default:
                                    _ = len(v28)
                                }
                                _ = fi207
                                _ = v208
                            }
                            _ = fi205
                            _ = v206
                        }
                        fmt.Println(v5, v30)
                    } else {
                        v3 -= (v3 >> ((42) & 63))
                    }
                }
            }
            var v214 int64 = -9223372036854775808
            var v215 int8 = 100
            _ = v214
            _ = v215
        }
        _ = v30
    } else {
        switch ((((v14 + v14) + v14) + ("\n\t" + (("" + ("foo bar" + "a")) + v14))) + ((v14 + v14) + (v14 + v14))) {
        case "hello":
            if false {
                v14 = v14
            }
        case "\n\t":
            for fi216 := 0; fi216 < 12; fi216++ {
                if (!(((v5 * v5) / v5) >= 0.0)) {
                    var v217 int8 = 0
                    v217 -= v217
                    _ = v217
                }
                fmt.Println(fi216)
                for fi218 := 0; fi218 < 14; fi218++ {
                    var v219 int64 = int64(v29)
                    _ = fi218
                    _ = v219
                }
                _ = fi216
            }
        default:
            if false {
                for fi220 := 0; fi220 < 15; fi220++ {
                    v29 += (^(v29))
                    if ((((!(-94 < 0)) && (true && (1e10 > (0.1 / 0.1)))) || true) || ((fi220 & 128) == fi220)) {
                        v14 = "hello"
                        v221 := GenS16{F0: 686, F1: -1}
                        var v222 int16 = -32768
                        _ = v221
                        _ = v222
                    } else {
                        switch v14 {
                        case "foo bar":
                            v28 = append(v28, -1)
                        }
                        var v223 uint8 = v29
                        _ = v223
                    }
                    var v224 int8 = -128
                    _ = fi220
                    _ = v224
                }
            } else {
                for fi225 := 0; fi225 < 14; fi225++ {
                    switch (fi225 % ((fi225) | 1)) {
                    case 32767:
                        for fi226 := 0; fi226 < 9; fi226++ {
                            v3 = (v3 + 31)
                            _ = fi226
                        }
                    default:
                        if false {
                            var v227 uint8 = v29
                            var v228 int16 = -32768
                            var v229 int32 = 52
                            _ = v227
                            _ = v228
                            _ = v229
                        } else {
                            v230 := GenS22{F0: 0, F1: -455}
                            for fi231 := 0; fi231 < 2; fi231++ {
                                v29 += (v29 << ((143) & 63))
                                _ = fi231
                            }
                            _ = v230
                        }
                    }
                    _ = fi225
                }
            }
        }
        fmt.Println(v29, v5, v14)
    }
    if len(v28) >= 2 {
        v28 = v28[1:len(v28)]
    }
    var v232 uint8 = 127
    var v233 uint16 = 0
    for fi234 := 0; fi234 < 9; fi234++ {
        var v235 int = 127
        var v236 bool = ((18446744073709551615 != 42) || false)
        _ = fi234
        _ = v235
        _ = v236
    }
    v237 := GenS23{F0: -1, F1: 802, F2: (uint(v3) != 255)}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v232:", v232)
    fmt.Println("v233:", v233)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v2
    _ = v232
    _ = v233
    _ = v237
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
