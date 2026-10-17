package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16 }
type GenS2 struct { F0 int16; F1 uint; F2 uint32 }
type GenS3 struct { F0 int16; F1 uint64; F2 uint8 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 uint64; F1 uint; F2 uint }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 uint; F1 int }
type GenS9 struct { F0 int64; F1 int32 }
type GenS10 struct { F0 int }
type GenS11 struct { F0 uint64; F1 uint64; F2 int32 }
type GenS12 struct { F0 int8; F1 uint8 }
type GenS13 struct { F0 uint64; F1 string; F2 uint32 }
type GenS14 struct { F0 uint16 }

func main() {
    v1 := make(map[int]uint32)
    var v2 int8 = 1
    v3 := []string{("\n\t" + "a"), "\n\t", ("foo bar" + ("" + ("hello" + "a"))), "hello"}
    switch 309 {
    case -195:
        var v4 uint8 = 42
        _ = v4
    default:
        fmt.Println(v2)
    }
    v5 := GenS1{F0: 0, F1: 255}
    var v6 uint16 = 65535
    for fi7 := 0; fi7 < 15; fi7++ {
        for fi8 := 0; fi8 < 6; fi8++ {
            v9 := GenS1{F0: 1, F1: v6}
            if false {
                switch (fi8 ^ fi8) {
                case -28:
                    v2 += v2
                case 128:
                    var v10 int16 = 32767
                    _ = v10
                case -761:
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                default:
                    var v11 uint32 = 841
                    _ = v11
                }
                if true {
                    v12 := v1[-127]
                    var v13 bool = true
                    if v13 {
                        if v13 {
                            switch (("a" + "foo bar") + "hello") {
                            case "":
                                if true {
                                    v6 -= v6
                                    if v13 {
                                        v12 = (v12 & (v12 | 42))
                                    }
                                }
                            case "\n\t":
                                fmt.Println(v6)
                            case "a":
                                v6 = (v6 * v6)
                            default:
                                v6 = 1
                            }
                        } else {
                            if ((!v13) && v13) {
                                v1[655] = v12
                                var v14 int16 = 1000
                                switch "hello" {
                                case "hello":
                                    if v13 {
                                        v15 := GenS1{F0: 1, F1: (v6 >> ((4294967295) & 63))}
                                        if v13 {
                                            v13 = true
                                            switch -9223372036854775808 {
                                            case -762:
                                                switch fi7 {
                                                case -1:
                                                    var v16 int16 = v14
                                                    _ = v16
                                                case -737:
                                                    var v17 string = ((((("a" + "foo bar") + "hello") + "\n\t") + "foo bar") + "\n\t")
                                                    _ = v17
                                                case 9223372036854775807:
                                                    var v18 uint = 4294967295
                                                    _ = v18
                                                }
                                            }
                                        } else {
                                            v13 = (true || ((v6 | v6) != (v6 + 160)))
                                            var v19 string = "foo bar"
                                            _ = v19
                                        }
                                        if (v13 && v13) {
                                            if (-1 < 889) {
                                                v6 -= (v6 ^ v6)
                                            } else {
                                                if ((((float32(float64((0.0 * 0.1))) / -2.25) - float32((0.0 * 1.5))) - 3.14159) == (-2.25 + float32(3.14159))) {
                                                    v3 = append(v3, "hello")
                                                    v20 := GenS2{F0: v14, F1: 4294967295, F2: v12}
                                                    v21 := GenS2{F0: v14, F1: 4294967295, F2: v12}
                                                    _ = v20
                                                    _ = v21
                                                }
                                                fmt.Println(fi8)
                                                var v22 int = fi7
                                                _ = v22
                                            }
                                        } else {
                                            v6 += (v6 | v6)
                                        }
                                        _ = v15
                                    }
                                case "\n\t":
                                    v13 = ((^(v2)) <= v2)
                                case "":
                                    if v13 {
                                        var v23 int8 = (v2 % ((v2) | 1))
                                        var v24 uint64 = 1
                                        switch (-(fi8)) {
                                        case 4294967295:
                                            _ = len(v1)
                                        case 991:
                                            if v13 {
                                                v2 -= (v23 ^ v2)
                                                v24 += (v24 << ((2) & 63))
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                            } else {
                                                v2 -= v23
                                                switch 9223372036854775807 {
                                                case 127:
                                                    var v25 int8 = 100
                                                    _ = v25
                                                }
                                                if len(v3) > 0 {
                                                    v3[0] = (("a" + "hello") + "")
                                                }
                                            }
                                        case 127:
                                            var v26 bool = true
                                            _ = v26
                                        default:
                                            var v27 int16 = v14
                                            _ = v27
                                        }
                                        _ = v23
                                        _ = v24
                                    }
                                }
                                _ = v14
                            } else {
                                switch "" {
                                case "":
                                    if false {
                                        var v28 uint = 4294967295
                                        var v29 uint = v28
                                        v1[32768] = v12
                                        var v30 int = (fi8 >> ((v29) & 63))
                                        _ = v28
                                        _ = v29
                                        _ = v30
                                    }
                                case "a":
                                    v6 = v6
                                case "foo bar":
                                    v31 := GenS2{F0: 32767, F1: 18446744073709551615, F2: 2147483648}
                                    _ = v31
                                }
                            }
                        }
                        if len(v3) > 0 {
                            v3[0] = "hello"
                        }
                        var v32 uint = 65535
                        _ = v32
                    } else {
                        switch ("hello" + "") {
                        case "a":
                            fmt.Println(v13, fi7)
                        case "hello":
                            if v13 {
                                var v33 float64 = (1e10 - -2.25)
                                _ = v33
                            }
                        case "\n\t":
                            v6 -= (v6 % ((v6) | 1))
                            var v34 float64 = 0.1
                            v1[fi8] = (v12 & v12)
                            _ = v34
                        default:
                            var v35 int8 = -1
                            _ = v35
                        }
                        if (2147483647 == 20) {
                            v2 = v2
                            switch (fi8 % ((32768) | 1)) {
                            case -9223372036854775808:
                                var v36 int64 = -9223372036854775808
                                _ = v36
                            case 287:
                                if (((((3.14159 - ((3.14159 - 0.1) * 1.5)) / -2.25) + 0.0) + float32(-2.25)) <= (float32((1e10 + float64(1e10))) - 1.5)) {
                                    if true {
                                        switch (("hello" + ("foo bar" + ("" + ""))) + "\n\t") {
                                        case "a":
                                            var v37 uint8 = 42
                                            _ = v37
                                        case "foo bar":
                                            v38 := GenS1{F0: 255, F1: 255}
                                            _ = v38
                                        }
                                        var v39 string = ("foo bar" + "")
                                        if (int16(v6) > 168) {
                                            var v40 uint8 = 37
                                            _ = v40
                                        }
                                        _ = v39
                                    }
                                    if false {
                                        v2 += v2
                                        if len(v3) > 0 {
                                            v3[0] = "a"
                                        }
                                    } else {
                                        v41 := GenS1{F0: 127, F1: (v6 | 255)}
                                        if (1 == fi8) {
                                            fmt.Println(v2, v6)
                                        }
                                        var v42 uint = 77
                                        _ = v41
                                        _ = v42
                                    }
                                    var v43 int = (-(fi7))
                                    _ = v43
                                } else {
                                    v13 = true
                                    fmt.Println(fi7, v13)
                                }
                                var v44 bool = (-1 != 42)
                                _ = v44
                            case 0:
                                v12 = 65535
                            }
                        } else {
                            v2 = (v2 / ((-1) | 1))
                            if (float64(1.5) > 3.14159) {
                                var v45 int8 = v2
                                v46 := GenS3{F0: 0, F1: 212, F2: 115}
                                _ = v45
                                _ = v46
                            }
                        }
                        switch fi7 {
                        case -1:
                            v6 -= (v6 - (v6 - (v6 - (v6 % (((v6 << ((334) & 63))) | 1)))))
                            var v47 uint64 = 0
                            var v48 int8 = v2
                            _ = v47
                            _ = v48
                        case -558:
                            v13 = v13
                        }
                    }
                    _ = v12
                    _ = v13
                } else {
                    v6 += (^(v6))
                    if true {
                        fmt.Println(fi7)
                        var v49 int = fi8
                        _ = v49
                    } else {
                        switch "hello" {
                        case "":
                            var v50 uint8 = 169
                            _ = v50
                        default:
                            var v51 uint32 = 287
                            switch "" {
                            case "":
                                var v52 int8 = (v2 / ((127) | 1))
                                var v53 bool = ("foo bar" == ("\n\t" + ("a" + "foo bar")))
                                v6 += 0
                                v54 := GenS4{F0: (^(v51))}
                                var v55 uint32 = v51
                                _ = v52
                                _ = v53
                                _ = v54
                                _ = v55
                            case "\n\t":
                                var v56 uint16 = (v6 & (v6 % (((v6 | (v6 * v6))) | 1)))
                                _ = v56
                            }
                            _ = v51
                        }
                    }
                }
            } else {
                var v57 uint16 = (^(v6))
                v3 = append(v3, "")
                v2 -= v2
                var v58 bool = ((((((("hello" + "a") + "") + (("\n\t" + "a") + "\n\t")) + "hello") >= ("\n\t" + "foo bar")) && ((0.0 + 0.0) >= 1.5)) && (-2.25 == ((0.1 / 3.14159) + 0.1)))
                var v59 int16 = 1
                _ = v57
                _ = v58
                _ = v59
            }
            var v60 uint16 = (v6 ^ v6)
            _ = fi8
            _ = v60
            _ = v9
        }
        _ = fi7
    }
    v61 := GenS3{F0: 32767, F1: 42, F2: 1}
    v62 := GenS2{F0: -611, F1: 65535, F2: 479}
    switch "hello" {
    case "hello":
        for fi63 := 0; fi63 < 14; fi63++ {
            if (!(true && false)) {
                v2 = int8(fi63)
                if true {
                    v2 = -1
                    v64 := GenS1{F0: 2, F1: v6}
                    _ = v64
                }
            } else {
                v2 = (v2 % ((int8(fi63)) | 1))
            }
            _ = fi63
        }
    case "foo bar":
        var v65 int16 = 0
        _ = v65
    case "a":
        v3 = append(v3, ("hello" + ((("" + "a") + "hello") + "\n\t")))
    }
    for fi66 := 0; fi66 < 19; fi66++ {
        v67 := GenS1{F0: 127, F1: v6}
        if false {
            v2 = -1
            if false {
                fmt.Println(fi66)
                if (0.1 > (-2.25 - 1.5)) {
                    var v68 bool = false
                    _ = v68
                }
            } else {
                v6 = v6
                fmt.Println(v6, v6)
            }
        } else {
            v6 = (v6 / ((2) | 1))
            for fi69 := 0; fi69 < 1; fi69++ {
                v2 = 57
                switch "hello" {
                case "":
                    if (3.14159 > -2.25) {
                        var v70 int8 = 73
                        _ = v70
                    } else {
                        switch "\n\t" {
                        case "\n\t":
                            fmt.Println(fi66, v6, v6)
                        case "a":
                            switch (fi69 - fi69) {
                            case 4294967295:
                                var v71 int64 = 245
                                _ = v71
                            case 1:
                                if (!(float32(0.1) > (1e10 - (0.1 - (float32(float64(((float32((0.0 / (-2.25 - 1e10))) - ((1.5 - 1e10) * (((float32((0.1 * float64((1.5 + 1e10)))) / 1e10) + float32(3.14159)) * 1.5))) / -2.25))) / 3.14159))))) {
                                    if (0 <= 0) {
                                        v72 := GenS4{F0: 65535}
                                        _ = v72
                                    } else {
                                        fmt.Println(v2)
                                        if (2147483647 <= -1) {
                                            switch fi66 {
                                            case -1:
                                                var v73 uint8 = 127
                                                _ = v73
                                            case 4294967295:
                                                if false {
                                                    if (2147483647 > 65535) {
                                                        v3 = append(v3, "")
                                                        fmt.Println(v6, v2)
                                                        var v74 int32 = 1
                                                        _ = v74
                                                    }
                                                    v75 := GenS5{F0: 569, F1: 4294967295, F2: 255}
                                                    v76 := GenS6{F0: true}
                                                    v6 += (v6 + (v6 ^ v6))
                                                    v3 = append(v3, ((("hello" + ("" + ("\n\t" + (("\n\t" + "") + "\n\t")))) + "hello") + "foo bar"))
                                                    _ = v75
                                                    _ = v76
                                                }
                                            case -9223372036854775808:
                                                v77 := GenS6{F0: (true || (901 <= uint32(v6)))}
                                                _ = v77
                                            }
                                            v2 -= v2
                                        } else {
                                            fmt.Println(v6)
                                            if false {
                                                v2 = 0
                                                var v78 int32 = 6
                                                v79 := GenS4{F0: 21}
                                                _ = v78
                                                _ = v79
                                            }
                                        }
                                        if true {
                                            v80 := GenS6{F0: (v6 < (^(v6)))}
                                            if false {
                                                v2 = (v2 & v2)
                                            }
                                            delete(v1, (fi69 % ((fi69) | 1)))
                                            _ = v80
                                        } else {
                                            v2 = (v2 ^ v2)
                                            switch "" {
                                            case "a":
                                                if true {
                                                    fmt.Println(fi66, fi69, fi69)
                                                    if true {
                                                        if ((fi66 + fi69) <= fi66) {
                                                            var v81 int64 = 42
                                                            delete(v1, -2147483648)
                                                            _ = v81
                                                        }
                                                        if len(v3) > 0 {
                                                            v3[0] = ("foo bar" + ("" + (("a" + "a") + "foo bar")))
                                                        }
                                                        var v82 int16 = -1
                                                        _ = v82
                                                    } else {
                                                        if (!(!(195 >= 65535))) {
                                                            var v83 int8 = v2
                                                            switch 4294967295 {
                                                            case 128:
                                                                v2 -= (^(v2))
                                                            case 529:
                                                                v84 := GenS1{F0: 1, F1: 204}
                                                                _ = v84
                                                            default:
                                                                var v85 bool = true
                                                                _ = v85
                                                            }
                                                            _ = v83
                                                        }
                                                    }
                                                    switch ("\n\t" + "") {
                                                    case "":
                                                        v2 += v2
                                                    default:
                                                        fmt.Println(v2)
                                                    }
                                                } else {
                                                    var v86 int64 = -721
                                                    _ = v86
                                                }
                                            default:
                                                var v87 int64 = -9223372036854775808
                                                _ = v87
                                            }
                                        }
                                    }
                                }
                            case -718:
                                v88 := GenS3{F0: 42, F1: 18446744073709551615, F2: 42}
                                v3 = append(v3, "")
                                switch ((("hello" + ("a" + "foo bar")) + "") + ("hello" + ("hello" + "hello"))) {
                                case "foo bar":
                                    v89 := GenS2{F0: 682, F1: 1, F2: 356}
                                    _ = v89
                                default:
                                    if true {
                                        v2 = -82
                                    } else {
                                        var v90 float64 = (((float64((0.1 - (1.5 * -2.25))) + 0.1) + float64(1.5)) + float64(1e10))
                                        _ = v90
                                    }
                                }
                                var v91 uint64 = 601
                                _ = v88
                                _ = v91
                            default:
                                v3 = append(v3, "")
                            }
                        default:
                            if ((fi66 - (fi69 / ((fi69) | 1))) <= fi69) {
                                var v92 uint8 = 191
                                var v93 float32 = (1e10 * 0.0)
                                var v94 uint32 = 709
                                _ = v92
                                _ = v93
                                _ = v94
                            } else {
                                switch ("\n\t" + ("\n\t" + "a")) {
                                case "\n\t":
                                    switch 969 {
                                    case 42:
                                        v6 += v6
                                    default:
                                        if ((1.5 - -2.25) == -2.25) {
                                            switch -893 {
                                            case 4294967295:
                                                var v95 int8 = v2
                                                _ = v95
                                            case 32768:
                                                switch (fi69 * fi69) {
                                                case -104:
                                                    fmt.Println(v6, fi69)
                                                }
                                                v2 -= (v2 + (v2 / ((v2) | 1)))
                                                v3 = append(v3, ("hello" + "hello"))
                                            case 585:
                                                fmt.Println(v2, v2)
                                                _ = len(v3)
                                                _ = len(v3)
                                                v96 := GenS7{F0: 1}
                                                _ = v96
                                            default:
                                                fmt.Println(v6)
                                                v6 = 298
                                                if (true && false) {
                                                    fmt.Println(fi66, fi66)
                                                } else {
                                                    if len(v3) >= 2 {
                                                        v3 = v3[1:len(v3)]
                                                    }
                                                }
                                            }
                                        }
                                    }
                                }
                                if true {
                                    fmt.Println(v6)
                                    if false {
                                        var v97 float32 = 1e10
                                        var v98 int16 = 32767
                                        if false {
                                            var v99 uint8 = 106
                                            switch ("a" + "\n\t") {
                                            case "a":
                                                if true {
                                                    if false {
                                                        if len(v3) >= 2 {
                                                            v3 = v3[1:len(v3)]
                                                        }
                                                        if false {
                                                            v100 := v1[fi66]
                                                            v101 := GenS8{F0: 65535, F1: fi66}
                                                            var v102 float64 = ((1e10 - -2.25) + (-2.25 * (3.14159 - 0.1)))
                                                            if true {
                                                                v97 -= (((v97 - 1e10) - v97) / 0.1)
                                                                var v103 int16 = v98
                                                                _ = v103
                                                            } else {
                                                                v97 += (v97 + v97)
                                                                v99 = uint8(v6)
                                                            }
                                                            var v104 string = ("\n\t" + "")
                                                            v105 := GenS9{F0: 42, F1: -475}
                                                            _ = v100
                                                            _ = v101
                                                            _ = v102
                                                            _ = v104
                                                            _ = v105
                                                        }
                                                    } else {
                                                        if ((v98 << ((uint(v6)) & 63)) == (v98 ^ v98)) {
                                                            switch "foo bar" {
                                                            case "\n\t":
                                                                var v106 int32 = 0
                                                                v3 = append(v3, "foo bar")
                                                                _ = v106
                                                            case "hello":
                                                                v98 = 0
                                                            case "a":
                                                                if len(v3) > 0 {
                                                                    v3[0] = "\n\t"
                                                                }
                                                            }
                                                        }
                                                    }
                                                    switch "\n\t" {
                                                    case "\n\t":
                                                        v99 = (v99 ^ (v99 - v99))
                                                    default:
                                                        v107 := GenS10{F0: fi69}
                                                        _ = v107
                                                    }
                                                    var v108 uint16 = (v6 >> ((0) & 63))
                                                    v109 := GenS9{F0: -838, F1: 728}
                                                    _ = v108
                                                    _ = v109
                                                }
                                            }
                                            _ = v99
                                        } else {
                                            if len(v3) > 0 {
                                                v3[0] = "hello"
                                            }
                                            var v110 float64 = ((0.0 + (0.1 - ((1.5 - 0.0) + (-2.25 - 0.0)))) - 0.1)
                                            switch "foo bar" {
                                            case "a":
                                                v2 += v2
                                            }
                                            switch "hello" {
                                            case "a":
                                                if true {
                                                    v111 := GenS11{F0: 2, F1: 2, F2: 1}
                                                    _ = v111
                                                } else {
                                                    fmt.Println(fi66, v6)
                                                    switch "\n\t" {
                                                    case "foo bar":
                                                        var v112 uint = 18446744073709551615
                                                        _ = v112
                                                    }
                                                }
                                            }
                                            v6 = 255
                                            _ = v110
                                        }
                                        _ = v97
                                        _ = v98
                                    }
                                    var v113 uint8 = 172
                                    if false {
                                        if ((0.1 + -2.25) <= (-2.25 + 1.5)) {
                                            fmt.Println(fi69)
                                        }
                                        fmt.Println(v2)
                                    }
                                    v6 += (v6 - (v6 % ((v6) | 1)))
                                    fmt.Println(v113, v2)
                                    _ = len(v1)
                                    _ = v113
                                }
                                var v114 int16 = 32767
                                _ = v114
                            }
                            delete(v1, (fi69 + -9223372036854775808))
                        }
                    }
                default:
                    v115 := GenS12{F0: v2, F1: 255}
                    _ = v115
                }
                _ = fi69
            }
            var v116 float32 = 1.5
            _ = v116
        }
        _ = fi66
        _ = v67
    }
    v117 := GenS13{F0: 4294967295, F1: "foo bar", F2: 1}
    v118 := GenS14{F0: v6}
    var v119 int32 = 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v119:", v119)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v117
    _ = v118
    _ = v119
    _ = v2
    _ = v3
    _ = v5
    _ = v6
    _ = v61
    _ = v62
}
