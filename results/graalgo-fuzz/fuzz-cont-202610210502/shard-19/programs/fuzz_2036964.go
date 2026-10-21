package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string; F2 uint8 }
type GenS2 struct { F0 uint8; F1 uint; F2 bool }
type GenS3 struct { F0 string; F1 int16; F2 uint64 }
type GenS4 struct { F0 uint; F1 string }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 int; F1 uint32; F2 int8 }
type GenS9 struct { F0 int; F1 int16 }
type GenS10 struct { F0 uint; F1 uint }
type GenS11 struct { F0 int16; F1 int16 }
type GenS12 struct { F0 int32; F1 uint8 }

func main() {
    var v1 uint64 = 0
    var v2 uint8 = 1
    v3 := []string{"a", "a"}
    var v4 bool = (42 <= 2147483647)
    v5 := []int{4294967295, 9223372036854775807}
    var v6 uint32 = 2147483647
    fmt.Println(v2, v1, v4)
    for fi7 := 0; fi7 < 11; fi7++ {
        v8 := GenS1{F0: (^(v6)), F1: ("" + "a"), F2: (v2 % (((v2 - 128)) | 1))}
        var v9 int64 = -640
        _ = fi7
        _ = v8
        _ = v9
    }
    var v10 string = "foo bar"
    var v11 int8 = 42
    v12 := GenS2{F0: v2, F1: uint(v2), F2: v4}
    var v13 bool = v4
    fmt.Println(v1, v4)
    v11 += (v11 << ((1) & 63))
    for fi14 := 0; fi14 < 9; fi14++ {
        v1 = uint64(v2)
        fmt.Println(fi14, v11, v11)
        _ = fi14
    }
    if v13 {
        fmt.Println(v1, v6, v11)
    }
    v15 := GenS2{F0: (v2 / (((v2 + 42)) | 1)), F1: 398, F2: v4}
    switch 872 {
    case 2147483647:
        v16 := GenS3{F0: ("\n\t" + ""), F1: 32767, F2: v1}
        _ = v16
    case 4294967295:
        if true {
            switch (v10 + v10) {
            case "\n\t":
                for fi17 := 0; fi17 < 12; fi17++ {
                    if (65535 > uint16(v6)) {
                        if v13 {
                            v10 = "a"
                            v18 := GenS2{F0: uint8(v1), F1: 2, F2: v13}
                            switch (fi17 << ((uint(v1)) & 63)) {
                            case 317:
                                switch "a" {
                                case "":
                                    if len(v5) > 0 {
                                        v5[0] = fi17
                                    }
                                case "a":
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                case "\n\t":
                                    for fi19 := 0; fi19 < 6; fi19++ {
                                        var v20 uint8 = (v2 & v2)
                                        _ = len(v5)
                                        _ = fi19
                                        _ = v20
                                    }
                                default:
                                    v21 := GenS2{F0: (v2 ^ 2), F1: 1, F2: ((false || true) || (0.0 != ((0.0 + -2.25) + 3.14159)))}
                                    _ = v21
                                }
                            case -691:
                                fmt.Println(v13)
                            default:
                                var v22 uint64 = (^(v1))
                                _ = v22
                            }
                            _ = v18
                        }
                        var v23 int8 = (-(v11))
                        fmt.Println(v6, v1, v4)
                        _ = v23
                    } else {
                        if false {
                            for fi24 := 0; fi24 < 18; fi24++ {
                                v1 -= v1
                                switch (fi24 >> ((42) & 63)) {
                                case 9223372036854775807:
                                    if (0.0 < float64((float32(0.1) - -2.25))) {
                                        switch v10 {
                                        case "hello":
                                            if v4 {
                                                v25 := GenS4{F0: 2, F1: ((("" + v10) + v10) + "foo bar")}
                                                v26 := GenS3{F0: "hello", F1: 1, F2: 0}
                                                _ = v25
                                                _ = v26
                                            } else {
                                                v2 = (v2 / (((v2 | (^(v2)))) | 1))
                                                fmt.Println(fi24, v10, v1)
                                            }
                                        case "\n\t":
                                            if false {
                                                v1 = v1
                                            }
                                            v27 := GenS1{F0: uint32(v1), F1: "foo bar", F2: v2}
                                            _ = v27
                                        default:
                                            fmt.Println(fi17)
                                        }
                                        var v28 uint = 867
                                        var v29 int32 = 2147483647
                                        _ = v28
                                        _ = v29
                                    }
                                case 42:
                                    if len(v5) > 0 {
                                        v5[0] = (fi17 | (fi24 | (fi24 - -874)))
                                    }
                                case 127:
                                    if len(v3) > 0 {
                                        v3[0] = ""
                                    }
                                }
                                fmt.Println(fi24)
                                _ = fi24
                            }
                        } else {
                            if true {
                                var v30 int32 = 0
                                _ = v30
                            }
                            v31 := GenS3{F0: v10, F1: 1, F2: (v1 - 389)}
                            if false {
                                if false {
                                    v2 -= 255
                                } else {
                                    var v32 int8 = v11
                                    if len(v5) > 0 {
                                        v5[0] = fi17
                                    }
                                    _ = v32
                                }
                            } else {
                                fmt.Println(fi17)
                                if len(v5) > 0 {
                                    v5[0] = (fi17 ^ 0)
                                }
                                var v33 bool = ((1e10 < (-2.25 + (3.14159 - float64(0.0)))) && true)
                                _ = v33
                            }
                            _ = v31
                        }
                        v13 = (0.1 == float64(0.1))
                        for fi34 := 0; fi34 < 12; fi34++ {
                            if (0 != 32767) {
                                fmt.Println(v11, fi34, v11)
                            } else {
                                switch -120 {
                                case 9223372036854775807:
                                    _ = len(v5)
                                case 32768:
                                    if (fi17 > 32767) {
                                        var v35 string = "a"
                                        if (249 > 2) {
                                            _ = len(v5)
                                            var v36 float32 = 3.14159
                                            v37 := GenS2{F0: 0, F1: 42, F2: true}
                                            var v38 float64 = -2.25
                                            v39 := GenS5{F0: 0}
                                            if (!(((v38 - v38) * v38) > (((v38 - (1.5 + v38)) * (((v38 - 0.1) + float64(3.14159)) * float64((((v36 - (v36 - v36)) - 0.0) / (v36 - (((v36 * v36) - v36) + v36)))))) * v38))) {
                                                var v40 uint = 624
                                                _ = v40
                                            }
                                            _ = v36
                                            _ = v37
                                            _ = v38
                                            _ = v39
                                        } else {
                                            var v41 uint8 = 127
                                            _ = v41
                                        }
                                        _ = v35
                                    }
                                case -1:
                                    var v42 uint16 = 611
                                    _ = v42
                                }
                                _ = len(v5)
                                v43 := GenS2{F0: v2, F1: 143, F2: true}
                                _ = v43
                            }
                            v44 := GenS5{F0: -32768}
                            v45 := GenS1{F0: (v6 ^ v6), F1: (v10 + ((v10 + v10) + (v10 + (v10 + "")))), F2: (v2 >> ((65535) & 63))}
                            v1 += (v1 / ((v1) | 1))
                            _ = fi34
                            _ = v44
                            _ = v45
                        }
                    }
                    if ((uint(v6) > 1) || v4) {
                        for fi46 := 0; fi46 < 16; fi46++ {
                            v3 = append(v3, v10)
                            var v47 string = "foo bar"
                            _ = fi46
                            _ = v47
                        }
                        v48 := GenS5{F0: -1}
                        v6 += v6
                        _ = v48
                    }
                    fmt.Println(fi17)
                    _ = fi17
                }
            }
            var v49 float64 = (0.1 + float64(1.5))
            _ = v49
        } else {
            v50 := GenS3{F0: "\n\t", F1: -32768, F2: v1}
            switch 42 {
            case 0:
                var v51 uint32 = v6
                _ = v51
            case 127:
                var v52 string = v10
                _ = v52
            case 32767:
                v1 -= (v1 / (((v1 % ((v1) | 1))) | 1))
            default:
                v53 := GenS3{F0: (v10 + v10), F1: 0, F2: 4294967295}
                _ = v53
            }
            _ = v50
        }
    default:
        if ((v6 != 2) && (3.14159 > 0.0)) {
            if false {
                var v54 int = -102
                for fi55 := 0; fi55 < 4; fi55++ {
                    for fi56 := 0; fi56 < 9; fi56++ {
                        v3 = append(v3, (v10 + "foo bar"))
                        _ = fi56
                    }
                    v57 := GenS1{F0: v6, F1: (v10 + "\n\t"), F2: (^(v2))}
                    _ = fi55
                    _ = v57
                }
                _ = v54
            }
            v58 := GenS4{F0: 65535, F1: (v10 + v10)}
            _ = v58
        }
    }
    if false {
        var v59 int16 = 32767
        switch (v10 + v10) {
        case "foo bar":
            v60 := GenS6{F0: 41}
            _ = v60
        case "\n\t":
            v13 = (!false)
        case "hello":
            var v61 int8 = 27
            _ = v61
        default:
            v1 = (^(v1))
        }
        _ = v59
    } else {
        fmt.Println(v13, v10, v4)
        switch (("foo bar" + v10) + v10) {
        case "":
            for fi62 := 0; fi62 < 19; fi62++ {
                fmt.Println(fi62, fi62, v11)
                switch "\n\t" {
                case "hello":
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                case "\n\t":
                    if (!false) {
                        if len(v3) > 0 {
                            v3[0] = "a"
                        }
                        for fi63 := 0; fi63 < 17; fi63++ {
                            switch fi62 {
                            case 627:
                                var v64 float32 = 1.5
                                _ = v64
                            default:
                                var v65 uint16 = 981
                                v2 = 2
                                _ = v65
                            }
                            v5 = append(v5, -1)
                            _ = fi63
                        }
                    } else {
                        for fi66 := 0; fi66 < 5; fi66++ {
                            var v67 uint64 = v1
                            v4 = v4
                            switch v10 {
                            case "hello":
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                            case "foo bar":
                                switch v10 {
                                case "hello":
                                    if v13 {
                                        var v68 int = (fi62 | fi62)
                                        v2 -= 128
                                        fmt.Println(v4)
                                        _ = v68
                                    }
                                    if (!(828 >= 1000)) {
                                        var v69 uint8 = (v2 ^ v2)
                                        _ = v69
                                    } else {
                                        var v70 int64 = 0
                                        if v4 {
                                            var v71 uint8 = (v2 * uint8(v67))
                                            if (0.0 >= (0.0 / ((((3.14159 + (float64(3.14159) - 1e10)) + (0.0 - 1e10)) - 0.1) * -2.25))) {
                                                var v72 float64 = 0.0
                                                var v73 float32 = float32(v72)
                                                _ = v72
                                                _ = v73
                                            } else {
                                                var v74 uint64 = v1
                                                v5 = append(v5, fi66)
                                                _ = v74
                                            }
                                            v10 = "a"
                                            _ = v71
                                        } else {
                                            if len(v3) > 0 {
                                                v3[0] = (v10 + v10)
                                            }
                                        }
                                        v2 = (^(v2))
                                        _ = v70
                                    }
                                case "foo bar":
                                    v5 = append(v5, 32767)
                                }
                            case "\n\t":
                                var v75 uint64 = uint64(v6)
                                v76 := GenS7{F0: -139}
                                var v77 uint32 = v6
                                _ = v75
                                _ = v76
                                _ = v77
                            }
                            _ = fi66
                            _ = v67
                        }
                    }
                case "foo bar":
                    if (3.14159 < (((0.1 / float64(1.5)) / float64(-2.25)) + 0.0)) {
                        v78 := GenS8{F0: (fi62 | fi62), F1: (v6 | v6), F2: (v11 % ((0) | 1))}
                        if v4 {
                            var v79 int16 = 951
                            var v80 uint32 = (^(v6))
                            _ = v79
                            _ = v80
                        } else {
                            if true {
                                for fi81 := 0; fi81 < 16; fi81++ {
                                    switch ("\n\t" + "\n\t") {
                                    case "hello":
                                        v2 += (^(v2))
                                    case "a":
                                        v82 := GenS5{F0: -902}
                                        _ = v82
                                    default:
                                        switch fi81 {
                                        case -125:
                                            if len(v5) > 0 {
                                                v5[0] = fi81
                                            }
                                        case -130:
                                            v2 = (v2 % ((0) | 1))
                                        case -2147483648:
                                            v83 := GenS7{F0: int32(fi81)}
                                            _ = v83
                                        default:
                                            var v84 int8 = v11
                                            _ = v84
                                        }
                                    }
                                    if (v2 < (v2 ^ v2)) {
                                        var v85 uint16 = 2
                                        _ = v85
                                    } else {
                                        switch 32768 {
                                        case -1:
                                            switch int(v2) {
                                            case 0:
                                                var v86 int = fi81
                                                _ = v86
                                            case 2147483647:
                                                v87 := GenS9{F0: 957, F1: -32768}
                                                _ = v87
                                            default:
                                                var v88 uint32 = (^(v6))
                                                _ = v88
                                            }
                                        case 1:
                                            switch (fi62 & (-(fi81))) {
                                            case 4294967295:
                                                v89 := GenS10{F0: 255, F1: 18446744073709551615}
                                                _ = v89
                                            }
                                        case -234:
                                            if (!v13) {
                                                v6 -= (v6 & v6)
                                                fmt.Println(v2, fi81, v13)
                                                v90 := GenS5{F0: 0}
                                                _ = v90
                                            } else {
                                                switch (v10 + v10) {
                                                case "a":
                                                    if (!(((((1.5 + (((float32(0.0) * 0.0) - 1.5) * 1e10)) - float32(0.1)) * (1e10 + 1.5)) >= (1e10 - -2.25)) && (!(65535 == uint(v1))))) {
                                                        if v4 {
                                                            var v91 int64 = 1
                                                            _ = v91
                                                        } else {
                                                            var v92 bool = (!true)
                                                            if (236 <= 1) {
                                                                v2 -= (v2 | (v2 - uint8(v6)))
                                                                if (!true) {
                                                                    var v93 int8 = v11
                                                                    _ = v93
                                                                }
                                                                var v94 int16 = 644
                                                                switch "" {
                                                                case "":
                                                                    v95 := GenS8{F0: (fi81 & (fi81 - fi62)), F1: (v6 | (v6 | 2147483647)), F2: (^(v11))}
                                                                    _ = v95
                                                                default:
                                                                    var v96 uint16 = uint16(v6)
                                                                    _ = v96
                                                                }
                                                                _ = v94
                                                            }
                                                            if (v11 <= (v11 + 116)) {
                                                                var v97 uint32 = (v6 << ((158) & 63))
                                                                _ = v97
                                                            }
                                                            _ = v92
                                                        }
                                                        var v98 uint64 = uint64(v6)
                                                        switch (fi62 << ((255) & 63)) {
                                                        case 9223372036854775807:
                                                            v11 = 0
                                                        case 32767:
                                                            v99 := GenS11{F0: 0, F1: -32768}
                                                            _ = v99
                                                        default:
                                                            if v4 {
                                                                v100 := GenS10{F0: 2, F1: 0}
                                                                _ = v100
                                                            } else {
                                                                switch v10 {
                                                                case "a":
                                                                    var v101 uint32 = v6
                                                                    _ = v101
                                                                case "hello":
                                                                    switch 32767 {
                                                                    case 127:
                                                                        if (!(false || v13)) {
                                                                            if len(v5) >= 2 {
                                                                                v5 = v5[1:len(v5)]
                                                                            }
                                                                            if v4 {
                                                                                v4 = ((v11 % (((v11 ^ (^(v11)))) | 1)) < 0)
                                                                                var v102 int8 = -128
                                                                                var v103 float64 = float64(1.5)
                                                                                _ = v102
                                                                                _ = v103
                                                                            } else {
                                                                                v104 := GenS2{F0: 1, F1: 2, F2: (893 > 1)}
                                                                                if (v4 && v13) {
                                                                                    if v4 {
                                                                                        var v105 bool = (1 == 792)
                                                                                        fmt.Println(v13)
                                                                                        v3 = append(v3, (v10 + v10))
                                                                                        _ = v105
                                                                                    }
                                                                                    v3 = append(v3, v10)
                                                                                }
                                                                                v106 := GenS4{F0: 945, F1: ("a" + "")}
                                                                                _ = v104
                                                                                _ = v106
                                                                            }
                                                                            v107 := GenS6{F0: 32767}
                                                                            _ = v107
                                                                        } else {
                                                                            fmt.Println(fi62, v98)
                                                                            var v108 int16 = 32767
                                                                            v5 = append(v5, fi81)
                                                                            _ = v108
                                                                        }
                                                                    case 137:
                                                                        switch v10 {
                                                                        case "hello":
                                                                            if len(v3) >= 2 {
                                                                                v3 = v3[1:len(v3)]
                                                                            }
                                                                            if true {
                                                                                fmt.Println(v2, v98, v11)
                                                                            }
                                                                        case "\n\t":
                                                                            switch (^(fi81)) {
                                                                            case 1:
                                                                                fmt.Println(v4, v1, v11)
                                                                            case -414:
                                                                                if len(v5) >= 2 {
                                                                                    v5 = v5[1:len(v5)]
                                                                                }
                                                                            case -2147483648:
                                                                                if v13 {
                                                                                    var v109 bool = ((v13 || (-857 > 42)) || true)
                                                                                    if len(v5) > 0 {
                                                                                        v5[0] = fi62
                                                                                    }
                                                                                    _ = v109
                                                                                } else {
                                                                                    v4 = v13
                                                                                    var v110 uint = 42
                                                                                    fmt.Println(v4, v4, fi62)
                                                                                    if v4 {
                                                                                        _ = len(v3)
                                                                                        if len(v3) > 0 {
                                                                                            v3[0] = v10
                                                                                        }
                                                                                    } else {
                                                                                        switch "hello" {
                                                                                        case "hello":
                                                                                            var v111 int8 = (v11 - (v11 / (((v11 - v11)) | 1)))
                                                                                            _ = v111
                                                                                        case "":
                                                                                            var v112 int8 = (v11 - 127)
                                                                                            _ = v112
                                                                                        case "foo bar":
                                                                                            v98 -= uint64(v110)
                                                                                        }
                                                                                        v4 = false
                                                                                        if len(v3) > 0 {
                                                                                            v3[0] = v10
                                                                                        }
                                                                                        v5 = append(v5, (fi62 << ((v110) & 63)))
                                                                                    }
                                                                                    switch ("\n\t" + ("" + ((v10 + (v10 + v10)) + (v10 + v10)))) {
                                                                                    case "":
                                                                                        var v113 uint16 = uint16(v6)
                                                                                        _ = v113
                                                                                    case "foo bar":
                                                                                        _ = len(v3)
                                                                                    }
                                                                                    _ = v110
                                                                                }
                                                                            }
                                                                        case "":
                                                                            fmt.Println(v1)
                                                                        default:
                                                                            if len(v3) >= 2 {
                                                                                v3 = v3[1:len(v3)]
                                                                            }
                                                                        }
                                                                    }
                                                                    if v13 {
                                                                        if true {
                                                                            v1 = 4294967295
                                                                            fmt.Println(v4, fi62, v11)
                                                                            var v114 uint8 = v2
                                                                            _ = v114
                                                                        }
                                                                        if false {
                                                                            v98 -= (v1 >> ((0) & 63))
                                                                            v115 := GenS6{F0: 2}
                                                                            _ = v115
                                                                        } else {
                                                                            var v116 uint8 = (v2 * (v2 & (v2 ^ (^(v2)))))
                                                                            v117 := GenS9{F0: (fi62 / (((fi81 << ((1) & 63))) | 1)), F1: 1000}
                                                                            switch (v10 + v10) {
                                                                            case "\n\t":
                                                                                var v118 uint64 = v1
                                                                                _ = v118
                                                                            default:
                                                                                if len(v5) > 0 {
                                                                                    v5[0] = fi81
                                                                                }
                                                                            }
                                                                            _ = v116
                                                                            _ = v117
                                                                        }
                                                                    }
                                                                    var v119 uint8 = 170
                                                                    _ = v119
                                                                case "":
                                                                    v6 += (v6 - v6)
                                                                default:
                                                                    var v120 int32 = 42
                                                                    _ = v120
                                                                }
                                                                var v121 uint32 = 2147483648
                                                                v13 = (1 <= 829)
                                                                _ = v121
                                                            }
                                                        }
                                                        _ = v98
                                                    }
                                                case "hello":
                                                    switch (fi62 & fi62) {
                                                    case -399:
                                                        var v122 uint32 = v6
                                                        _ = v122
                                                    case 847:
                                                        switch v10 {
                                                        case "foo bar":
                                                            _ = len(v3)
                                                        case "\n\t":
                                                            switch "hello" {
                                                            case "a":
                                                                var v123 uint = 65535
                                                                _ = v123
                                                            case "foo bar":
                                                                var v124 string = v10
                                                                _ = v124
                                                            case "\n\t":
                                                                var v125 int = (^(fi81))
                                                                _ = v125
                                                            }
                                                        case "hello":
                                                            v13 = true
                                                        default:
                                                            var v126 string = v10
                                                            _ = v126
                                                        }
                                                    case 4294967295:
                                                        switch v10 {
                                                        case "":
                                                            v3 = append(v3, "a")
                                                        case "hello":
                                                            fmt.Println(v13, v2, v11)
                                                        }
                                                    default:
                                                        var v127 string = "hello"
                                                        _ = v127
                                                    }
                                                default:
                                                    v6 -= v6
                                                }
                                                _ = len(v3)
                                            }
                                        default:
                                            var v128 int16 = int16(v2)
                                            _ = v128
                                        }
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                        v4 = (fi81 <= fi62)
                                        var v129 string = (v10 + (((v10 + (v10 + v10)) + v10) + "foo bar"))
                                        _ = v129
                                    }
                                    if (0 != 551) {
                                        switch "a" {
                                        case "hello":
                                            var v130 uint32 = v6
                                            _ = v130
                                        case "a":
                                            fmt.Println(v10, v10)
                                        case "foo bar":
                                            var v131 uint32 = (v6 / (((^(v6))) | 1))
                                            _ = v131
                                        }
                                    } else {
                                        var v132 uint = 65535
                                        var v133 float32 = ((0.0 + 1e10) / (1.5 + 0.1))
                                        _ = v132
                                        _ = v133
                                    }
                                    _ = fi81
                                }
                                for fi134 := 0; fi134 < 17; fi134++ {
                                    var v135 string = (v10 + "foo bar")
                                    v136 := GenS12{F0: 1, F1: (^(v2))}
                                    _ = fi134
                                    _ = v135
                                    _ = v136
                                }
                            } else {
                                if ((v11 - (v11 << ((4294967295) & 63))) < (v11 * (v11 >> ((801) & 63)))) {
                                    var v137 uint8 = (v2 << ((33) & 63))
                                    v138 := GenS3{F0: v10, F1: 1, F2: (v1 + v1)}
                                    var v139 int16 = 0
                                    _ = v137
                                    _ = v138
                                    _ = v139
                                } else {
                                    switch fi62 {
                                    case 32767:
                                        v1 = v1
                                    case 4294967295:
                                        var v140 uint64 = v1
                                        _ = v140
                                    case 0:
                                        fmt.Println(v13, v4)
                                    }
                                    if len(v5) > 0 {
                                        v5[0] = (fi62 / (((fi62 << ((uint(v6)) & 63))) | 1))
                                    }
                                }
                            }
                            var v141 int16 = 1000
                            v10 = (("hello" + v10) + v10)
                            _ = v141
                        }
                        _ = v78
                    }
                default:
                    var v142 int64 = 9223372036854775807
                    _ = v142
                }
                if (v4 && true) {
                    v143 := GenS6{F0: 32767}
                    fmt.Println(fi62, v6, v4)
                    _ = v143
                } else {
                    var v144 uint16 = 255
                    _ = v144
                }
                _ = fi62
            }
        case "\n\t":
            if len(v5) > 0 {
                v5[0] = -2147483648
            }
        }
    }
    _ = len(v3)
    if false {
        var v145 uint32 = (v6 << ((18446744073709551615) & 63))
        v11 -= 127
        _ = v145
    }
    var v146 int8 = v11
    var v147 int64 = -9223372036854775808
    switch 4294967295 {
    case 4294967295:
        v10 = "a"
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v146:", v146)
    fmt.Println("v147:", v147)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v146
    _ = v147
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
