package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int32; F1 int16; F2 int }
type GenS5 struct { F0 int16; F1 uint32 }
type GenS6 struct { F0 int64; F1 int64; F2 int8 }
type GenS7 struct { F0 uint32; F1 int }
type GenS8 struct { F0 uint32 }
type GenS9 struct { F0 uint64; F1 uint; F2 int32 }
type GenS10 struct { F0 int8; F1 uint8; F2 uint32 }
type GenS11 struct { F0 uint; F1 uint8; F2 int16 }
type GenS12 struct { F0 int16; F1 uint8 }
type GenS13 struct { F0 uint16 }
type GenS14 struct { F0 uint8 }
type GenS15 struct { F0 bool }
type GenS16 struct { F0 int8 }
type GenS17 struct { F0 uint16; F1 int16; F2 bool }

func main() {
    v1 := make(map[string]bool)
    var v2 float32 = 0.0
    v3 := []int{0}
    var v4 int64 = 9223372036854775807
    var v5 float64 = (0.0 * -2.25)
    v6 := []uint32{1, 4294967295, 673, 995}
    switch 42 {
    case 233:
        v7 := GenS1{F0: 255}
        _ = v7
    case -2147483648:
        var v8 uint64 = 493
        _ = v8
    }
    for fi9 := 0; fi9 < 3; fi9++ {
        var v10 float64 = v5
        _ = fi9
        _ = v10
    }
    fmt.Println(v2)
    fmt.Println(v4, v2)
    for fi11 := 0; fi11 < 15; fi11++ {
        var v12 int16 = -922
        v4 += v4
        _ = fi11
        _ = v12
    }
    var v13 uint32 = 65535
    v13 += 2147483648
    v4 -= v4
    if false {
        if (!false) {
            if len(v6) > 0 {
                v6[0] = v13
            }
            v4 -= -9223372036854775808
            if (844 == 32768) {
                var v14 uint32 = (^(v13))
                _ = v14
            }
        }
    } else {
        for fi15 := 0; fi15 < 13; fi15++ {
            switch "hello" {
            case "a":
                fmt.Println(v2, v13)
            case "\n\t":
                var v16 string = "hello"
                _ = v16
            default:
                switch -541 {
                case 1:
                    if true {
                        if false {
                            for fi17 := 0; fi17 < 6; fi17++ {
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                switch ((((("" + ("\n\t" + "foo bar")) + ("" + ((("a" + ((("hello" + ("a" + "")) + ("\n\t" + "\n\t")) + "hello")) + "hello") + "hello"))) + "foo bar") + "") + ("hello" + ("" + (("a" + "hello") + ("hello" + ("hello" + "hello")))))) {
                                case "":
                                    fmt.Println(fi17, v5)
                                case "foo bar":
                                    v18 := GenS2{F0: -1}
                                    _ = v18
                                }
                                if false {
                                    if true {
                                        v2 += v2
                                        v19 := GenS1{F0: 128}
                                        if (v4 <= 750) {
                                            if ((^(v13)) > v13) {
                                                var v20 float64 = v5
                                                _ = v20
                                            } else {
                                                if (18446744073709551615 == 173) {
                                                    v5 += (0.0 - ((3.14159 / 1.0) / v5))
                                                } else {
                                                    var v21 int64 = v4
                                                    v4 += v4
                                                    _ = v21
                                                }
                                                v22 := GenS3{F0: 949}
                                                _ = v22
                                            }
                                        }
                                        _ = v19
                                    } else {
                                        fmt.Println(fi17)
                                        v1["a"] = (2 < 18446744073709551615)
                                        var v23 uint16 = 315
                                        _ = v23
                                    }
                                    if true {
                                        var v24 uint = 18446744073709551615
                                        var v25 uint = 1
                                        _ = v24
                                        _ = v25
                                    } else {
                                        var v26 float32 = v2
                                        if false {
                                            v4 = -9223372036854775808
                                        }
                                        if (0 <= 2) {
                                            switch (fi17 - (fi17 / (((fi15 - (-(fi17)))) | 1))) {
                                            case 699:
                                                v26 = 1e10
                                            case 127:
                                                if ((false || (!(((^(v4)) <= v4) || (0 < -2147483648)))) || true) {
                                                    if true {
                                                        _ = len(v1)
                                                        v4 = v4
                                                        var v27 uint8 = 128
                                                        _ = v27
                                                    } else {
                                                        var v28 float64 = v5
                                                        if false {
                                                            if true {
                                                                switch ((("" + "hello") + ("foo bar" + "")) + ("a" + "")) {
                                                                case "\n\t":
                                                                    if ("foo bar" >= "a") {
                                                                        var v29 uint8 = 128
                                                                        fmt.Println(v28, v26)
                                                                        var v30 int = (fi17 << ((0) & 63))
                                                                        _ = v29
                                                                        _ = v30
                                                                    } else {
                                                                        var v31 float32 = -2.25
                                                                        _ = v31
                                                                    }
                                                                case "":
                                                                    switch "" {
                                                                    case "hello":
                                                                        if true {
                                                                            if false {
                                                                                if true {
                                                                                    v6 = append(v6, v13)
                                                                                    if true {
                                                                                        var v32 uint16 = 728
                                                                                        v33 := GenS1{F0: 42}
                                                                                        _ = v32
                                                                                        _ = v33
                                                                                    }
                                                                                } else {
                                                                                    v5 -= v5
                                                                                    v28 += v28
                                                                                    fmt.Println(v4)
                                                                                }
                                                                                if true {
                                                                                    if true {
                                                                                        fmt.Println(v28, v13, fi17)
                                                                                        fmt.Println(v2, fi17)
                                                                                        v2 += (-2.25 * 1e10)
                                                                                    }
                                                                                    var v34 int32 = 0
                                                                                    _ = v34
                                                                                } else {
                                                                                    var v35 string = "foo bar"
                                                                                    _ = v35
                                                                                }
                                                                            }
                                                                            _ = len(v1)
                                                                            fmt.Println(v4)
                                                                        } else {
                                                                            v2 = v2
                                                                            var v36 uint = 348
                                                                            if (255 > 77) {
                                                                                v37 := GenS4{F0: -386, F1: -32768, F2: 4294967295}
                                                                                _ = v37
                                                                            } else {
                                                                                if (!(true && (!true))) {
                                                                                    v38 := v1["\n\t"]
                                                                                    _ = v38
                                                                                } else {
                                                                                    fmt.Println(v28, v13, v36)
                                                                                }
                                                                                v39 := GenS5{F0: 32767, F1: (v13 + (v13 + (v13 / ((v13) | 1))))}
                                                                                _ = v39
                                                                            }
                                                                            _ = v36
                                                                        }
                                                                    case "\n\t":
                                                                        switch "\n\t" {
                                                                        case "hello":
                                                                            if false {
                                                                                v40 := GenS2{F0: int16(fi17)}
                                                                                if false {
                                                                                    var v41 int8 = -36
                                                                                    if false {
                                                                                        var v42 uint16 = 0
                                                                                        _ = v42
                                                                                    }
                                                                                    _ = v41
                                                                                } else {
                                                                                    if false {
                                                                                        switch fi15 {
                                                                                        case -1:
                                                                                            if false {
                                                                                                switch "\n\t" {
                                                                                                case "a":
                                                                                                    fmt.Println(fi15, fi17, v5)
                                                                                                case "hello":
                                                                                                    v43 := GenS4{F0: 976, F1: 32767, F2: (^(fi15))}
                                                                                                    _ = v43
                                                                                                case "":
                                                                                                    v2 = v26
                                                                                                }
                                                                                                v5 += 3.14159
                                                                                            } else {
                                                                                                var v44 int32 = 42
                                                                                                v28 += v28
                                                                                                var v45 uint8 = 128
                                                                                                _ = v44
                                                                                                _ = v45
                                                                                            }
                                                                                        case -315:
                                                                                            switch -9223372036854775808 {
                                                                                            case -1:
                                                                                                var v46 uint16 = 309
                                                                                                _ = v46
                                                                                            case 128:
                                                                                                if false {
                                                                                                    v3 = append(v3, (fi17 + fi17))
                                                                                                    v47 := GenS2{F0: 596}
                                                                                                    if false {
                                                                                                        v5 -= (3.14159 + (3.14159 / (((1e10 / (0.1 / float64(0.1))) - (float64(v26) / v28)) - float64(float32(v28)))))
                                                                                                    }
                                                                                                    _ = v47
                                                                                                }
                                                                                            case 4294967295:
                                                                                                var v48 int64 = v4
                                                                                                _ = v48
                                                                                            default:
                                                                                                v49 := v1["a"]
                                                                                                _ = v49
                                                                                            }
                                                                                        case 1:
                                                                                            if ((42 > 1) && false) {
                                                                                                switch "\n\t" {
                                                                                                case "\n\t":
                                                                                                    v26 = v2
                                                                                                case "a":
                                                                                                    v13 = 55
                                                                                                }
                                                                                                var v50 float32 = (v2 * (float32(v5) + v2))
                                                                                                switch ((("a" + "") + "") + "hello") {
                                                                                                case "foo bar":
                                                                                                    v26 = float32(v5)
                                                                                                case "\n\t":
                                                                                                    switch "foo bar" {
                                                                                                    case "":
                                                                                                        if true {
                                                                                                            if false {
                                                                                                                _ = len(v1)
                                                                                                                v6 = append(v6, 1)
                                                                                                            } else {
                                                                                                                fmt.Println(v2, v5, v2)
                                                                                                            }
                                                                                                        }
                                                                                                    case "hello":
                                                                                                        v1[("a" + (("a" + "\n\t") + ("" + (((("a" + ("a" + ((("" + "a") + "") + "a"))) + "\n\t") + ("\n\t" + "hello")) + ""))))] = false
                                                                                                    case "a":
                                                                                                        _ = len(v3)
                                                                                                        v28 = v5
                                                                                                        v13 = v13
                                                                                                    default:
                                                                                                        v3 = append(v3, fi17)
                                                                                                    }
                                                                                                case "":
                                                                                                    var v51 uint32 = (v13 >> ((1) & 63))
                                                                                                    _ = v51
                                                                                                }
                                                                                                _ = v50
                                                                                            } else {
                                                                                                v28 += 1e10
                                                                                                v52 := GenS3{F0: 755}
                                                                                                _ = v52
                                                                                            }
                                                                                        default:
                                                                                            var v53 int32 = -1
                                                                                            _ = v53
                                                                                        }
                                                                                        _ = len(v6)
                                                                                        var v54 int64 = (-(v4))
                                                                                        _ = v54
                                                                                    } else {
                                                                                        var v55 float32 = v26
                                                                                        v28 = float64(((float32((v28 - v28)) - 0.0) * v55))
                                                                                        _ = v55
                                                                                    }
                                                                                    v2 = ((((float32(3.14159) - v2) / 0.1) / ((v2 / v2) * float32(v28))) + (v2 / v26))
                                                                                }
                                                                                v13 = (v13 | 2147483648)
                                                                                switch fi15 {
                                                                                case 4294967295:
                                                                                    fmt.Println(v5)
                                                                                case -1:
                                                                                    v56 := GenS6{F0: (-(v4)), F1: -1, F2: -27}
                                                                                    _ = v56
                                                                                default:
                                                                                    var v57 int64 = (v4 - 42)
                                                                                    _ = v57
                                                                                }
                                                                                if false {
                                                                                    var v58 string = ("hello" + "a")
                                                                                    if len(v3) > 0 {
                                                                                        v3[0] = fi15
                                                                                    }
                                                                                    v58 = v58
                                                                                    var v59 bool = (true || true)
                                                                                    _ = v58
                                                                                    _ = v59
                                                                                } else {
                                                                                    fmt.Println(v4, v5)
                                                                                }
                                                                                _ = v40
                                                                            } else {
                                                                                var v60 int64 = v4
                                                                                var v61 float32 = v26
                                                                                _ = v60
                                                                                _ = v61
                                                                            }
                                                                        case "foo bar":
                                                                            _ = len(v1)
                                                                        case "\n\t":
                                                                            if (417 != 0) {
                                                                                fmt.Println(v4, fi15)
                                                                                v62 := GenS7{F0: (v13 * v13), F1: (fi17 - (fi17 & (fi15 + (^(fi17)))))}
                                                                                fmt.Println(v5, v26, v28)
                                                                                var v63 uint16 = 32767
                                                                                _ = v62
                                                                                _ = v63
                                                                            } else {
                                                                                var v64 uint32 = 2147483648
                                                                                _ = v64
                                                                            }
                                                                        default:
                                                                            var v65 bool = (false || (-85 < 0))
                                                                            _ = v65
                                                                        }
                                                                    default:
                                                                        v66 := GenS8{F0: 0}
                                                                        _ = v66
                                                                    }
                                                                case "hello":
                                                                    fmt.Println(fi15, v26, fi15)
                                                                default:
                                                                    var v67 float32 = 1.5
                                                                    _ = v67
                                                                }
                                                                var v68 uint32 = (v13 | (v13 >> ((255) & 63)))
                                                                _ = v68
                                                            } else {
                                                                var v69 int32 = 1
                                                                switch fi17 {
                                                                case -743:
                                                                    switch (fi17 - (fi15 << ((898) & 63))) {
                                                                    case 9223372036854775807:
                                                                        v4 -= 1
                                                                    case -9223372036854775808:
                                                                        fmt.Println(fi17)
                                                                    default:
                                                                        v2 = 1e10
                                                                        v69 = v69
                                                                    }
                                                                case 32767:
                                                                    v3 = append(v3, fi15)
                                                                    switch "foo bar" {
                                                                    case "hello":
                                                                        fmt.Println(fi17, v69, fi15)
                                                                    case "a":
                                                                        fmt.Println(fi15, v28)
                                                                    case "foo bar":
                                                                        v1[(((("\n\t" + "foo bar") + ("a" + "foo bar")) + "") + "a")] = true
                                                                    }
                                                                case 1:
                                                                    v5 = (v5 - v5)
                                                                }
                                                                var v70 uint16 = 32767
                                                                _ = v69
                                                                _ = v70
                                                            }
                                                            var v71 uint = uint(v13)
                                                            switch fi17 {
                                                            case 651:
                                                                if (!(false && false)) {
                                                                    v71 = (v71 | v71)
                                                                    v5 -= v5
                                                                    var v72 bool = (true && (!true))
                                                                    v73 := GenS9{F0: 42, F1: (v71 | (^(v71))), F2: 940}
                                                                    _ = v72
                                                                    _ = v73
                                                                }
                                                            case 32768:
                                                                v74 := GenS7{F0: v13, F1: (fi15 + 128)}
                                                                _ = v74
                                                            case -1:
                                                                var v75 uint = (v71 | (v71 ^ 1))
                                                                _ = v75
                                                            }
                                                            _ = v71
                                                        }
                                                        var v76 uint16 = 701
                                                        _ = v28
                                                        _ = v76
                                                    }
                                                    _ = len(v3)
                                                } else {
                                                    var v77 int = fi15
                                                    fmt.Println(v5, v2, v5)
                                                    _ = v77
                                                }
                                            case -9223372036854775808:
                                                if false {
                                                    v78 := GenS10{F0: 100, F1: 255, F2: (v13 * 862)}
                                                    _ = v78
                                                } else {
                                                    var v79 string = ("hello" + ("foo bar" + "\n\t"))
                                                    _ = v79
                                                }
                                            default:
                                                if false {
                                                    var v80 float64 = float64((((((v26 / ((((-2.25 / v26) - (-2.25 * 1.5)) + 3.14159) + (float32((3.14159 + 0.0)) - v2))) / float32(((float64((float32(((0.0 * float64(v2)) / v5)) * float32(0.0))) / float64((-2.25 - 0.0))) + float64((v2 - v2))))) + v26) * 1.5) * v26))
                                                    v80 += ((((v5 - float64(v2)) - (float64(float32(0.0)) - v80)) / v80) - 0.0)
                                                    var v81 uint32 = (v13 + 4294967295)
                                                    _ = v80
                                                    _ = v81
                                                }
                                            }
                                            var v82 string = "hello"
                                            _ = v82
                                        }
                                        _ = v26
                                    }
                                }
                                fmt.Println(v2, fi15)
                                v83 := GenS11{F0: 42, F1: 42, F2: 1}
                                fmt.Println(fi15, v13, fi17)
                                v5 = -2.25
                                v84 := GenS2{F0: -1}
                                _ = fi17
                                _ = v83
                                _ = v84
                            }
                        } else {
                            v5 = (v5 - 0.0)
                            delete(v1, ("hello" + ("hello" + "hello")))
                        }
                        v13 = 42
                    }
                case 128:
                    v1["\n\t"] = false
                default:
                    if ((fi15 >> ((65535) & 63)) > fi15) {
                        v85 := GenS7{F0: 4294967295, F1: fi15}
                        var v86 uint16 = 255
                        var v87 int16 = -610
                        _ = v85
                        _ = v86
                        _ = v87
                    }
                }
            }
            var v88 int16 = 32767
            v5 = float64(v2)
            _ = fi15
            _ = v88
        }
        if true {
            v2 += v2
            switch 753 {
            case -852:
                v89 := GenS12{F0: 42, F1: 0}
                _ = v89
            case -1:
                switch "foo bar" {
                case "hello":
                    var v90 string = "hello"
                    _ = v90
                default:
                    v91 := GenS13{F0: uint16(v13)}
                    _ = v91
                }
            case 9223372036854775807:
                if false {
                    fmt.Println(v13, v13)
                    v5 -= v5
                } else {
                    v92 := GenS9{F0: 42, F1: 4294967295, F2: 0}
                    if true {
                        v3 = append(v3, -9223372036854775808)
                    } else {
                        var v93 uint64 = 648
                        _ = v93
                    }
                    _ = v92
                }
            }
            fmt.Println(v2, v2, v4)
        }
    }
    v94 := GenS14{F0: 1}
    var v95 string = ""
    v96 := v1[(v95 + ((v95 + v95) + (v95 + "hello")))]
    fmt.Println(v96, v2)
    v1[v95] = (199 <= 255)
    switch v95 {
    case "foo bar":
        if (255 < 18446744073709551615) {
            fmt.Println(v13)
            switch "" {
            case "a":
                var v97 bool = (((v95 + ((v95 + "") + "foo bar")) <= v95) || (236 < 255))
                _ = v97
            }
            switch "foo bar" {
            case "foo bar":
                var v98 float64 = v5
                _ = v98
            case "\n\t":
                if true {
                    switch v95 {
                    case "hello":
                        var v99 uint32 = 2
                        _ = v99
                    case "":
                        v100 := v1[("hello" + ((((v95 + v95) + "hello") + v95) + v95))]
                        _ = v100
                    }
                    for fi101 := 0; fi101 < 18; fi101++ {
                        var v102 int64 = (-(v4))
                        if v96 {
                            for fi103 := 0; fi103 < 20; fi103++ {
                                switch (fi101 | (fi103 + fi101)) {
                                case -9223372036854775808:
                                    var v104 uint64 = 4294967295
                                    _ = v104
                                case 4294967295:
                                    v105 := GenS4{F0: -2147483648, F1: 768, F2: 216}
                                    _ = v105
                                default:
                                    v106 := GenS15{F0: v96}
                                    _ = v106
                                }
                                _ = fi103
                            }
                        } else {
                            for fi107 := 0; fi107 < 3; fi107++ {
                                delete(v1, v95)
                                fmt.Println(v95)
                                v108 := GenS16{F0: -32}
                                _ = fi107
                                _ = v108
                            }
                        }
                        _ = fi101
                        _ = v102
                    }
                } else {
                    for fi109 := 0; fi109 < 7; fi109++ {
                        switch (-(fi109)) {
                        case 128:
                            v110 := v1[(v95 + "hello")]
                            _ = v110
                        }
                        var v111 float64 = v5
                        v3 = append(v3, 17)
                        _ = fi109
                        _ = v111
                    }
                    var v112 string = "hello"
                    if v96 {
                        _ = len(v1)
                    } else {
                        v113 := GenS17{F0: 65535, F1: -1, F2: false}
                        var v114 uint64 = 18446744073709551615
                        if v96 {
                            v95 = v112
                            v112 = (("" + (v95 + ("hello" + (v95 + (((((v95 + ("foo bar" + "")) + "foo bar") + (v112 + v95)) + v95) + v112))))) + (("\n\t" + ((("\n\t" + (v95 + v95)) + "hello") + (((((("\n\t" + ("hello" + v112)) + v95) + "\n\t") + v112) + v95) + ("foo bar" + v112)))) + v112))
                        }
                        _ = v113
                        _ = v114
                    }
                    _ = v112
                }
            }
        }
    case "hello":
        if v96 {
            switch -163 {
            case -496:
                v2 = v2
            case 42:
                v115 := GenS3{F0: uint(v13)}
                _ = v115
            default:
                if len(v6) > 0 {
                    v6[0] = v13
                }
            }
            v96 = v96
        } else {
            if (false && (v96 || v96)) {
                var v116 int = 128
                v13 += v13
                var v117 uint32 = v13
                _ = v116
                _ = v117
            } else {
                for fi118 := 0; fi118 < 10; fi118++ {
                    for fi119 := 0; fi119 < 7; fi119++ {
                        if v96 {
                            fmt.Println(v4)
                            if len(v6) > 0 {
                                v6[0] = 591
                            }
                            if (v96 && v96) {
                                var v120 int8 = 127
                                _ = v120
                            } else {
                                v121 := GenS8{F0: 278}
                                v122 := GenS3{F0: 65535}
                                _ = len(v1)
                                fmt.Println(fi118, v5, v95)
                                var v123 float32 = 0.0
                                _ = v121
                                _ = v122
                                _ = v123
                            }
                        } else {
                            var v124 int = -2147483648
                            switch (((("\n\t" + v95) + (v95 + "hello")) + v95) + v95) {
                            case "":
                                v125 := v1[v95]
                                _ = v125
                            }
                            switch "\n\t" {
                            case "a":
                                v2 -= (float32((1.5 * (-2.25 / v5))) * (v2 / 0.1))
                            case "foo bar":
                                if false {
                                    v96 = (v5 <= float64(v2))
                                    switch v95 {
                                    case "foo bar":
                                        v3 = append(v3, 127)
                                    }
                                } else {
                                    if len(v6) >= 2 {
                                        v6 = v6[1:len(v6)]
                                    }
                                    fmt.Println(v2, v2, fi119)
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                }
                            case "":
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                            }
                            fmt.Println(v96, v13)
                            var v126 int16 = 1000
                            _ = v124
                            _ = v126
                        }
                        switch "hello" {
                        case "foo bar":
                            v5 -= (((v5 * v5) * 0.0) - float64(float32(v5)))
                        case "a":
                            fmt.Println(v95)
                        case "\n\t":
                            fmt.Println(v13)
                        }
                        var v127 string = "\n\t"
                        _ = fi119
                        _ = v127
                    }
                    _ = fi118
                }
                v2 += 0.1
            }
        }
    }
    var v128 int64 = (v4 * (v4 + -9223372036854775808))
    v5 = 0.0
    for fi129 := 0; fi129 < 3; fi129++ {
        v130 := GenS7{F0: (^(v13)), F1: fi129}
        v4 = v128
        _ = fi129
        _ = v130
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v128:", v128)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v95:", v95)
    fmt.Println("v96:", v96)
    _ = v1
    _ = v128
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v94
    _ = v95
    _ = v96
}
