package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16 }
type GenS2 struct { F0 int16; F1 uint }
type GenS3 struct { F0 uint; F1 int32 }
type GenS4 struct { F0 uint16; F1 uint32; F2 uint32 }
type GenS5 struct { F0 uint32; F1 uint64 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 bool; F1 int32 }
type GenS8 struct { F0 string }
type GenS9 struct { F0 int }
type GenS10 struct { F0 int64; F1 uint; F2 string }

func main() {
    v1 := []string{"", "hello"}
    var v2 string = ("\n\t" + "\n\t")
    var v3 float64 = 1e10
    v4 := make(map[int]string)
    var v5 uint16 = 2
    var v6 bool = true
    for fi7 := 0; fi7 < 17; fi7++ {
        fmt.Println(v3, v3, fi7)
        switch ("hello" + "foo bar") {
        case "a":
            if v6 {
                fmt.Println(v5, v2, fi7)
                switch (^(fi7)) {
                case -1:
                    if v6 {
                        v2 = (("hello" + "hello") + ("hello" + ""))
                    } else {
                        if v6 {
                            fmt.Println(v5, fi7)
                        } else {
                            v8 := GenS1{F0: 42, F1: v5}
                            var v9 int8 = -43
                            _ = v8
                            _ = v9
                        }
                        var v10 int64 = 300
                        fmt.Println(v10)
                        _ = v10
                    }
                case 42:
                    v11 := GenS2{F0: -608, F1: 1}
                    _ = v11
                case 4294967295:
                    if (v3 != 3.14159) {
                        fmt.Println(v3, fi7)
                    } else {
                        var v12 uint64 = 42
                        _ = v12
                    }
                default:
                    var v13 float32 = 3.14159
                    _ = v13
                }
                v14 := GenS3{F0: 1, F1: 583}
                _ = v14
            } else {
                var v15 uint = 2
                if false {
                    v3 = (0.0 + (-2.25 - v3))
                }
                _ = v15
            }
        case "foo bar":
            var v16 float64 = 3.14159
            _ = v16
        case "hello":
            v17 := GenS4{F0: v5, F1: 42, F2: 4294967295}
            _ = v17
        default:
            if v6 {
                var v18 uint = 2
                _ = v18
            }
        }
        _ = fi7
    }
    for fi19 := 0; fi19 < 3; fi19++ {
        for fi20 := 0; fi20 < 16; fi20++ {
            switch (fi19 >> ((4294967295) & 63)) {
            case -667:
                fmt.Println(v2, fi19, fi19)
            case -693:
                if false {
                    switch fi19 {
                    case -2147483648:
                        v21 := GenS5{F0: 0, F1: 18446744073709551615}
                        _ = v21
                    case -862:
                        var v22 bool = (((!v6) && v6) && false)
                        _ = v22
                    case 1000:
                        var v23 uint64 = 42
                        _ = v23
                    default:
                        v5 -= 263
                    }
                    var v24 int64 = -1
                    _ = v24
                }
            default:
                var v25 uint8 = 229
                _ = v25
            }
            _ = len(v1)
            _ = fi20
        }
        _ = fi19
    }
    v2 = "\n\t"
    v5 = v5
    v2 = ""
    if v6 {
        switch 2147483647 {
        case 0:
            v6 = v6
        case 9223372036854775807:
            var v26 int = -2147483648
            _ = v26
        default:
            if len(v1) > 0 {
                v1[0] = "a"
            }
        }
        v27 := GenS6{F0: -122}
        _ = v27
    } else {
        if true {
            if (41 >= 2) {
                if (1 < -24) {
                    if v6 {
                        for fi28 := 0; fi28 < 11; fi28++ {
                            for fi29 := 0; fi29 < 13; fi29++ {
                                if (-465 > 42) {
                                    v2 = v2
                                }
                                fmt.Println(v6, fi28)
                                _ = fi29
                            }
                            v4[fi28] = v2
                            v30 := GenS7{F0: v6, F1: int32(fi28)}
                            _ = fi28
                            _ = v30
                        }
                        for fi31 := 0; fi31 < 2; fi31++ {
                            v5 = (v5 - v5)
                            v2 = ((v2 + (((v2 + (v2 + (v2 + ("\n\t" + v2)))) + "foo bar") + "hello")) + v2)
                            _ = fi31
                        }
                        if ((0 >= 65535) && (!((!(1000 > 0)) && (1 > 0)))) {
                            var v32 float32 = float32(((((v3 * ((((v3 + v3) / 0.1) + (v3 - v3)) / ((float64((1e10 + -2.25)) / 0.1) * v3))) - float64(1e10)) * ((0.1 / 1.5) - v3)) + -2.25))
                            v6 = true
                            for fi33 := 0; fi33 < 19; fi33++ {
                                fmt.Println(v6)
                                _ = fi33
                            }
                            _ = v32
                        }
                    } else {
                        fmt.Println(v5)
                    }
                    v6 = v6
                    v2 = (v2 + "hello")
                }
                v34 := GenS8{F0: v2}
                _ = v34
            } else {
                fmt.Println(v3, v3, v3)
                v1 = append(v1, ("hello" + v2))
                var v35 int16 = -32768
                _ = v35
            }
            var v36 uint32 = 2147483647
            _ = v36
        } else {
            for fi37 := 0; fi37 < 19; fi37++ {
                var v38 int8 = 127
                if ((fi37 > fi37) || ((fi37 >= 2147483647) && (false && ((v38 << ((976) & 63)) >= v38)))) {
                    v39 := GenS7{F0: (65535 <= 2), F1: -1}
                    if false {
                        v40 := GenS7{F0: v6, F1: 2147483647}
                        fmt.Println(fi37)
                        _ = v40
                    } else {
                        v2 = v2
                        for fi41 := 0; fi41 < 7; fi41++ {
                            var v42 string = v2
                            _ = fi41
                            _ = v42
                        }
                    }
                    _ = v39
                }
                if (-1 == -73) {
                    for fi43 := 0; fi43 < 10; fi43++ {
                        var v44 int64 = 42
                        _ = fi43
                        _ = v44
                    }
                }
                _ = fi37
                _ = v38
            }
        }
        switch (v2 + ((v2 + "foo bar") + v2)) {
        case "a":
            var v45 uint64 = 4294967295
            _ = v45
        case "hello":
            switch v2 {
            case "hello":
                fmt.Println(v5, v3)
            case "a":
                if v6 {
                    if (!((v5 * (v5 & 669)) > v5)) {
                        if (1 == 774) {
                            v46 := GenS4{F0: (^(v5)), F1: 444, F2: 2147483648}
                            var v47 uint8 = 130
                            fmt.Println(v3, v6)
                            _ = v46
                            _ = v47
                        }
                        if v6 {
                            if ((^(v5)) < v5) {
                                for fi48 := 0; fi48 < 8; fi48++ {
                                    var v49 int32 = -1
                                    switch ("hello" + (v2 + (((v2 + v2) + v2) + "a"))) {
                                    case "":
                                        if true {
                                            if len(v1) > 0 {
                                                v1[0] = "foo bar"
                                            }
                                        }
                                    default:
                                        switch (fi48 | (fi48 ^ (fi48 / ((fi48) | 1)))) {
                                        case 4294967295:
                                            var v50 string = "foo bar"
                                            _ = v50
                                        default:
                                            switch (v2 + "hello") {
                                            case "hello":
                                                for fi51 := 0; fi51 < 15; fi51++ {
                                                    switch -2147483648 {
                                                    case -377:
                                                        v2 = ("" + v2)
                                                    default:
                                                        _ = len(v4)
                                                    }
                                                    switch (fi48 >> ((1) & 63)) {
                                                    case 850:
                                                        switch v2 {
                                                        case "a":
                                                            switch fi48 {
                                                            case -9223372036854775808:
                                                                v52 := GenS4{F0: 32767, F1: 785, F2: uint32(v5)}
                                                                _ = v52
                                                            default:
                                                                fmt.Println(fi48, v2, v5)
                                                            }
                                                        case "foo bar":
                                                            if (132 <= 127) {
                                                                if v6 {
                                                                    var v53 float32 = (1e10 + (1.5 - (1e10 - float32(((1e10 - v3) / (v3 + v3))))))
                                                                    v5 = (v5 >> ((1) & 63))
                                                                    var v54 int = (fi51 - fi51)
                                                                    if (v6 && v6) {
                                                                        v55 := GenS5{F0: 2147483647, F1: 4294967295}
                                                                        v5 = (v5 + (^(v5)))
                                                                        _ = v55
                                                                    }
                                                                    _ = v53
                                                                    _ = v54
                                                                }
                                                                if (!("" <= v2)) {
                                                                    v49 -= (-(v49))
                                                                }
                                                            } else {
                                                                v6 = (0 > 100)
                                                                switch fi48 {
                                                                case 9223372036854775807:
                                                                    fmt.Println(fi51, v49, v2)
                                                                case 941:
                                                                    fmt.Println(fi48, v6)
                                                                case -1:
                                                                    if v6 {
                                                                        var v56 uint8 = 127
                                                                        _ = v56
                                                                    }
                                                                }
                                                                fmt.Println(fi51, v2, v49)
                                                            }
                                                        case "hello":
                                                            if v6 {
                                                                v6 = v6
                                                                v57 := GenS5{F0: 2, F1: 4294967295}
                                                                v2 = v2
                                                                _ = v57
                                                            } else {
                                                                _ = len(v1)
                                                                fmt.Println(v2)
                                                            }
                                                            v58 := GenS6{F0: -1}
                                                            _ = v58
                                                        default:
                                                            if len(v1) > 0 {
                                                                v1[0] = v2
                                                            }
                                                        }
                                                    case -1:
                                                        v3 -= 1.5
                                                    default:
                                                        if (v6 || false) {
                                                            fmt.Println(fi48, v49)
                                                        }
                                                    }
                                                    var v59 int8 = 112
                                                    v2 = "hello"
                                                    switch (fi51 - fi48) {
                                                    case 56:
                                                        if len(v1) > 0 {
                                                            v1[0] = (("a" + "\n\t") + v2)
                                                        }
                                                    case 127:
                                                        v59 += 127
                                                    case 1:
                                                        v60 := GenS4{F0: v5, F1: uint32(v5), F2: 2147483648}
                                                        switch fi51 {
                                                        case 127:
                                                            var v61 uint64 = 781
                                                            switch v2 {
                                                            case "a":
                                                                var v62 uint32 = 0
                                                                _ = v62
                                                            case "hello":
                                                                _ = len(v4)
                                                            case "\n\t":
                                                                if len(v1) >= 2 {
                                                                    v1 = v1[1:len(v1)]
                                                                }
                                                            default:
                                                                switch (fi48 % (((fi48 + fi51)) | 1)) {
                                                                case 42:
                                                                    _ = len(v1)
                                                                default:
                                                                    switch (fi51 * int(v49)) {
                                                                    case 0:
                                                                        var v63 uint64 = (v61 % ((306) | 1))
                                                                        _ = v63
                                                                    default:
                                                                        if (float32(1.5) >= 0.1) {
                                                                            if v6 {
                                                                                var v64 bool = ("a" != "")
                                                                                if v6 {
                                                                                    var v65 uint16 = (v5 | (^(v5)))
                                                                                    var v66 bool = v64
                                                                                    v67 := GenS3{F0: 65535, F1: (v49 | v49)}
                                                                                    switch v2 {
                                                                                    case "":
                                                                                        v59 += v59
                                                                                    default:
                                                                                        v49 = 0
                                                                                    }
                                                                                    _ = v65
                                                                                    _ = v66
                                                                                    _ = v67
                                                                                } else {
                                                                                    var v68 string = v2
                                                                                    fmt.Println(fi51, v49, v59)
                                                                                    _ = v68
                                                                                }
                                                                                _ = v64
                                                                            }
                                                                        } else {
                                                                            v5 = v5
                                                                        }
                                                                    }
                                                                }
                                                            }
                                                            v69 := GenS7{F0: false, F1: v49}
                                                            _ = v61
                                                            _ = v69
                                                        case -723:
                                                            if (!false) {
                                                                var v70 int32 = (v49 % ((v49) | 1))
                                                                _ = v70
                                                            } else {
                                                                v71 := GenS2{F0: 218, F1: 65535}
                                                                _ = v71
                                                            }
                                                        case -98:
                                                            if false {
                                                                var v72 uint = 18446744073709551615
                                                                var v73 uint16 = (v5 - (v5 << ((255) & 63)))
                                                                _ = v72
                                                                _ = v73
                                                            } else {
                                                                v3 = v3
                                                            }
                                                        }
                                                        _ = v60
                                                    default:
                                                        var v74 float32 = 3.14159
                                                        _ = v74
                                                    }
                                                    _ = fi51
                                                    _ = v59
                                                }
                                            case "foo bar":
                                                if (true || v6) {
                                                    if (fi48 > fi48) {
                                                        fmt.Println(v5, v6, fi48)
                                                    }
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                } else {
                                                    v3 += (v3 - v3)
                                                    switch v2 {
                                                    case "\n\t":
                                                        var v75 int64 = 0
                                                        _ = v75
                                                    default:
                                                        if v6 {
                                                            _ = len(v4)
                                                        }
                                                    }
                                                }
                                            }
                                        }
                                    }
                                    _ = fi48
                                    _ = v49
                                }
                                v5 = v5
                                v76 := GenS7{F0: v6, F1: 1}
                                _ = v76
                            } else {
                                if false {
                                    var v77 string = (v2 + v2)
                                    if v6 {
                                        if (v6 || (v6 && v6)) {
                                            v4[277] = "hello"
                                            v77 = (v2 + "\n\t")
                                            if (762 < -972) {
                                                fmt.Println(v77, v3, v6)
                                                if (88 == 177) {
                                                    var v78 float64 = v3
                                                    _ = v78
                                                }
                                            } else {
                                                v3 = v3
                                            }
                                        }
                                    } else {
                                        switch "\n\t" {
                                        case "hello":
                                            v79 := GenS5{F0: 975, F1: 1}
                                            _ = v79
                                        case "a":
                                            var v80 int = 4294967295
                                            _ = v80
                                        case "foo bar":
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                        }
                                        var v81 uint8 = 42
                                        if true {
                                            for fi82 := 0; fi82 < 2; fi82++ {
                                                fmt.Println(v2)
                                                v81 = (v81 << ((4294967295) & 63))
                                                if v6 {
                                                    fmt.Println(v3, v3, fi82)
                                                    v83 := GenS8{F0: ("foo bar" + "foo bar")}
                                                    _ = v83
                                                }
                                                _ = fi82
                                            }
                                            var v84 string = (v77 + v77)
                                            _ = v84
                                        } else {
                                            for fi85 := 0; fi85 < 15; fi85++ {
                                                var v86 uint32 = 462
                                                _ = fi85
                                                _ = v86
                                            }
                                        }
                                        _ = v81
                                    }
                                    if (!(!false)) {
                                        var v87 int8 = 42
                                        v87 -= (v87 ^ -16)
                                        _ = v87
                                    } else {
                                        switch (("" + "\n\t") + v77) {
                                        case "foo bar":
                                            v2 = v77
                                        case "":
                                            switch "\n\t" {
                                            case "a":
                                                var v88 uint32 = 42
                                                _ = v88
                                            case "hello":
                                                v89 := GenS7{F0: v6, F1: 1}
                                                _ = v89
                                            default:
                                                v6 = false
                                            }
                                        }
                                    }
                                    _ = v77
                                }
                            }
                            switch -794 {
                            case 904:
                                v5 = (^(v5))
                            }
                            v2 = "hello"
                        } else {
                            v5 = (v5 & (v5 / ((v5) | 1)))
                            v3 += (v3 + ((v3 + ((v3 - float64(1e10)) + ((-2.25 / (v3 + v3)) * v3))) / v3))
                        }
                    } else {
                        v90 := GenS3{F0: 65535, F1: 0}
                        if (-455 == 1) {
                            if v6 {
                                switch (v2 + v2) {
                                case "hello":
                                    for fi91 := 0; fi91 < 1; fi91++ {
                                        switch 42 {
                                        case 128:
                                            var v92 uint = 161
                                            _ = v92
                                        default:
                                            v6 = ((^(v5)) >= 32767)
                                        }
                                        fmt.Println(v6, fi91)
                                        fmt.Println(v2, fi91, v5)
                                        _ = fi91
                                    }
                                case "\n\t":
                                    v93 := GenS9{F0: 683}
                                    _ = v93
                                case "a":
                                    v2 = v2
                                }
                            }
                            fmt.Println(v5, v5)
                            v94 := GenS6{F0: 1}
                            _ = v94
                        }
                        if true {
                            var v95 int8 = -1
                            _ = v95
                        }
                        _ = v90
                    }
                }
            default:
                v96 := GenS10{F0: 9223372036854775807, F1: 153, F2: "foo bar"}
                _ = v96
            }
        case "":
            v3 = (0.0 - 1.5)
        default:
            if (47 >= 128) {
                var v97 uint32 = 2147483647
                if len(v1) > 0 {
                    v1[0] = "a"
                }
                _ = v97
            }
        }
    }
    for fi98 := 0; fi98 < 20; fi98++ {
        for fi99 := 0; fi99 < 13; fi99++ {
            var v100 int32 = -662
            v101 := GenS2{F0: 32767, F1: 18446744073709551615}
            _ = fi99
            _ = v100
            _ = v101
        }
        switch fi98 {
        case 32767:
            v6 = (127 <= (fi98 + fi98))
        }
        _ = fi98
    }
    var v102 string = ((((((v2 + v2) + v2) + v2) + ((v2 + (v2 + ("" + ((v2 + v2) + v2)))) + "")) + v2) + v2)
    v5 = 42
    if v6 {
        var v103 uint64 = 0
        if v6 {
            switch 32768 {
            case 4294967295:
                fmt.Println(v103, v3, v2)
            case 35:
                for fi104 := 0; fi104 < 12; fi104++ {
                    fmt.Println(v6, v6, v3)
                    fmt.Println(v2)
                    v2 = ""
                    _ = fi104
                }
            default:
                fmt.Println(v103, v5)
            }
            v3 = (3.14159 + 1e10)
            fmt.Println(v5, v6, v5)
        }
        var v105 int = -2147483648
        _ = v103
        _ = v105
    }
    v106 := GenS6{F0: 127}
    switch 0 {
    case 32767:
        v102 = "hello"
    default:
        var v107 int = -122
        _ = v107
    }
    if v6 {
        var v108 int = 802
        if (2147483647 > 956) {
            var v109 float64 = (v3 * ((((3.14159 / (v3 - v3)) / 1.0) - 1.5) - (float64(float32(1e10)) * v3)))
            var v110 int16 = -32768
            _ = len(v4)
            _ = v109
            _ = v110
        } else {
            for fi111 := 0; fi111 < 12; fi111++ {
                for fi112 := 0; fi112 < 8; fi112++ {
                    switch (fi112 & (fi112 | (fi112 >> ((439) & 63)))) {
                    case 706:
                        var v113 string = "foo bar"
                        _ = v113
                    default:
                        _ = len(v4)
                        v5 -= (v5 - v5)
                    }
                    if (4294967295 < -1) {
                        _ = len(v4)
                    }
                    _ = fi112
                }
                for fi114 := 0; fi114 < 17; fi114++ {
                    v108 = (fi111 / ((fi111) | 1))
                    v2 = v2
                    v6 = v6
                    _ = fi114
                }
                _ = fi111
            }
            var v115 int8 = 53
            _ = v115
        }
        v2 = "a"
        _ = v108
    }
    switch 9223372036854775807 {
    case -1:
        var v116 int16 = 1000
        _ = v116
    case 4294967295:
        v3 -= float64(0.0)
    case -916:
        switch "" {
        case "hello":
            v3 -= ((float64(3.14159) - (((0.1 + 1e10) / (float64(1.5) - 1e10)) + (-2.25 + 1e10))) / (1e10 / (float64(-2.25) + (0.0 - float64(float32(float64(1.5)))))))
        default:
            var v117 int16 = 1
            _ = v117
        }
    }
    v118 := GenS3{F0: 331, F1: 2147483647}
    v119 := GenS4{F0: 255, F1: 3, F2: 884}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v102:", v102)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v102
    _ = v106
    _ = v118
    _ = v119
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
