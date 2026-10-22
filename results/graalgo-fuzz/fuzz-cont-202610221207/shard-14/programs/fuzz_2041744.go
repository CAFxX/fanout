package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int; F1 uint8 }
type GenS3 struct { F0 string; F1 int }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int; F1 uint64 }
type GenS6 struct { F0 int; F1 bool; F2 uint16 }
type GenS7 struct { F0 int16; F1 int32 }
type GenS8 struct { F0 uint8; F1 int64; F2 int32 }
type GenS9 struct { F0 int16; F1 int32; F2 uint64 }
type GenS10 struct { F0 uint64 }
type GenS11 struct { F0 uint8; F1 uint64; F2 int8 }
type GenS12 struct { F0 uint }
type GenS13 struct { F0 uint32 }
type GenS14 struct { F0 uint64; F1 int32 }

func main() {
    var v1 int32 = 1
    var v2 int16 = 0
    var v3 uint = 2
    v4 := make(map[uint32]string)
    var v5 float64 = 0.1
    var v6 int = 9223372036854775807
    v7 := GenS1{F0: ("\n\t" + "a")}
    for fi8 := 0; fi8 < 1; fi8++ {
        for fi9 := 0; fi9 < 6; fi9++ {
            v10 := []int{fi9}
            var v11 float64 = float64(1e10)
            _ = fi9
            _ = v10
            _ = v11
        }
        fmt.Println(v5, v6)
        switch (("" + "hello") + ((("" + "\n\t") + (("" + "a") + "a")) + ("a" + "a"))) {
        case "hello":
            for fi12 := 0; fi12 < 6; fi12++ {
                if (255 == 179) {
                    v13 := GenS2{F0: fi12, F1: uint8(v3)}
                    _ = v13
                } else {
                    switch (("foo bar" + "") + "\n\t") {
                    case "\n\t":
                        if false {
                            if (382 != 65535) {
                                if ((-2.25 * ((((float32(float64(float32((float64((1.5 / ((0.0 - 1e10) - 1e10))) / 3.14159)))) * 3.14159) * (0.0 + 3.14159)) + (1.5 + (1.5 - -2.25))) - 1e10)) < (((0.0 - (float32(v5) + float32(v5))) + -2.25) / float32(v5))) {
                                    v6 = 280
                                    if true {
                                        if true {
                                            var v14 uint32 = 0
                                            switch fi8 {
                                            case -738:
                                                switch "" {
                                                case "foo bar":
                                                    v6 = fi12
                                                case "hello":
                                                    if false {
                                                        switch "a" {
                                                        case "\n\t":
                                                            v2 = 0
                                                        case "":
                                                            switch -216 {
                                                            case 128:
                                                                v15 := GenS3{F0: "hello", F1: (fi12 & fi12)}
                                                                _ = v15
                                                            case -2147483648:
                                                                if true {
                                                                    _ = len(v4)
                                                                    var v16 uint = v3
                                                                    if (false || (((v5 * (((v5 - v5) / v5) - 1e10)) / (((v5 - 1e10) + (v5 / float64(1e10))) + -2.25)) > 1.5)) {
                                                                        var v17 uint16 = 940
                                                                        _ = v17
                                                                    } else {
                                                                        var v18 int16 = -647
                                                                        v16 = v16
                                                                        v2 -= (v2 * v2)
                                                                        v19 := GenS4{F0: v14}
                                                                        var v20 bool = (true || true)
                                                                        _ = v18
                                                                        _ = v19
                                                                        _ = v20
                                                                    }
                                                                    _ = v16
                                                                } else {
                                                                    v21 := GenS5{F0: fi12, F1: 1}
                                                                    _ = v21
                                                                }
                                                                v22 := GenS1{F0: "hello"}
                                                                var v23 uint64 = 4294967295
                                                                _ = v22
                                                                _ = v23
                                                            case 1:
                                                                var v24 uint32 = v14
                                                                _ = v24
                                                            default:
                                                                v14 = (v14 * 0)
                                                            }
                                                        }
                                                        v5 += 3.14159
                                                        switch "\n\t" {
                                                        case "hello":
                                                            delete(v4, v14)
                                                        default:
                                                            v2 += (v2 % ((42) | 1))
                                                        }
                                                    }
                                                }
                                            case 952:
                                                v25 := GenS5{F0: 2147483647, F1: 4294967295}
                                                _ = v25
                                            case 32767:
                                                var v26 uint = (v3 % ((910) | 1))
                                                var v27 int8 = int8(v1)
                                                fmt.Println(v27)
                                                _ = v26
                                                _ = v27
                                            }
                                            if true {
                                                v28 := []int{}
                                                _ = v28
                                            } else {
                                                var v29 uint16 = 65535
                                                v30 := GenS6{F0: (fi12 / ((v6) | 1)), F1: true, F2: v29}
                                                v29 += uint16(v3)
                                                _ = v29
                                                _ = v30
                                            }
                                            _ = v14
                                        }
                                        fmt.Println(fi12)
                                        v5 = ((v5 * (1.5 / float64(float32((0.1 - 1.5))))) / -2.25)
                                    }
                                    v1 += v1
                                } else {
                                    var v31 uint8 = 175
                                    if ((^(fi12)) == (fi8 % (((fi8 + -2147483648)) | 1))) {
                                        switch "hello" {
                                        case "":
                                            switch "\n\t" {
                                            case "hello":
                                                switch "" {
                                                case "hello":
                                                    if false {
                                                        fmt.Println(v5, v6)
                                                        v32 := GenS7{F0: (v2 - (v2 % (((v2 * v2)) | 1))), F1: (v1 * v1)}
                                                        _ = v32
                                                    } else {
                                                        v1 -= v1
                                                        var v33 uint8 = (^(v31))
                                                        _ = v33
                                                    }
                                                default:
                                                    switch (("a" + "\n\t") + ("hello" + ("\n\t" + ("hello" + (((((("a" + "") + "") + "a") + ("hello" + (((("\n\t" + "hello") + "foo bar") + ("hello" + "hello")) + (("hello" + "\n\t") + "foo bar")))) + ("" + ("\n\t" + ""))) + "hello"))))) {
                                                    case "a":
                                                        fmt.Println(fi12, v3)
                                                    }
                                                }
                                            case "a":
                                                v31 = v31
                                            case "foo bar":
                                                v6 = fi12
                                            }
                                            v34 := []int{4294967295}
                                            v35 := GenS8{F0: (^(v31)), F1: 253, F2: v1}
                                            _ = v34
                                            _ = v35
                                        case "\n\t":
                                            var v36 string = "hello"
                                            _ = v36
                                        case "foo bar":
                                            if false {
                                                var v37 float64 = v5
                                                _ = v37
                                            }
                                        default:
                                            switch (fi8 * int(v1)) {
                                            case -2147483648:
                                                if false {
                                                    if false {
                                                        v2 -= v2
                                                        var v38 uint16 = 2
                                                        var v39 float64 = float64(1e10)
                                                        _ = v38
                                                        _ = v39
                                                    }
                                                    var v40 uint32 = 633
                                                    _ = v40
                                                } else {
                                                    var v41 uint64 = 0
                                                    v42 := GenS7{F0: v2, F1: int32(v3)}
                                                    _ = v41
                                                    _ = v42
                                                }
                                            case 2147483647:
                                                v43 := v4[1]
                                                _ = v43
                                            case -443:
                                                switch "\n\t" {
                                                case "a":
                                                    var v44 uint = (v3 * v3)
                                                    _ = v44
                                                default:
                                                    v45 := GenS9{F0: (v2 / (((v2 - v2)) | 1)), F1: v1, F2: 18446744073709551615}
                                                    _ = v45
                                                }
                                            }
                                        }
                                        switch "hello" {
                                        case "hello":
                                            v5 += ((float64((1e10 - ((-2.25 - -2.25) * 3.14159))) / v5) * v5)
                                        case "\n\t":
                                            v46 := GenS10{F0: 4294967295}
                                            _ = v46
                                        default:
                                            switch (fi8 >> ((v3) & 63)) {
                                            case -465:
                                                if true {
                                                    if false {
                                                        switch v6 {
                                                        case 752:
                                                            var v47 uint16 = 2
                                                            _ = v47
                                                        case 1:
                                                            v48 := []int{(fi8 / ((fi8) | 1)), (fi12 & v6)}
                                                            _ = v48
                                                        case -2147483648:
                                                            var v49 int32 = (-(v1))
                                                            _ = v49
                                                        default:
                                                            switch -667 {
                                                            case -938:
                                                                var v50 uint16 = 2
                                                                _ = v50
                                                            case 0:
                                                                v3 += v3
                                                            default:
                                                                v5 = 0.0
                                                            }
                                                        }
                                                        v51 := GenS1{F0: (("hello" + "a") + "foo bar")}
                                                        switch ("hello" + "hello") {
                                                        case "foo bar":
                                                            var v52 int16 = 32767
                                                            _ = v52
                                                        case "hello":
                                                            var v53 uint32 = 65535
                                                            _ = v53
                                                        default:
                                                            switch "foo bar" {
                                                            case "":
                                                                var v54 float64 = (float64(0.0) + -2.25)
                                                                v2 = v2
                                                                _ = v54
                                                            case "a":
                                                                v55 := GenS7{F0: v2, F1: (v1 << ((488) & 63))}
                                                                _ = v55
                                                            }
                                                        }
                                                        _ = v51
                                                    }
                                                    var v56 int32 = (^(v1))
                                                    _ = v56
                                                } else {
                                                    if true {
                                                        fmt.Println(v1)
                                                    }
                                                    var v57 float64 = v5
                                                    v1 -= (v1 - -361)
                                                    var v58 int32 = (v1 >> ((1) & 63))
                                                    fmt.Println(v1, v31)
                                                    var v59 uint = (v3 << ((v3) & 63))
                                                    if true {
                                                        var v60 int32 = -195
                                                        _ = v60
                                                    } else {
                                                        var v61 string = "\n\t"
                                                        if (347 == 65535) {
                                                            var v62 int64 = 396
                                                            v63 := GenS8{F0: v31, F1: 9223372036854775807, F2: -1}
                                                            fmt.Println(fi12)
                                                            _ = v62
                                                            _ = v63
                                                        }
                                                        v2 = v2
                                                        _ = v61
                                                    }
                                                    _ = v57
                                                    _ = v58
                                                    _ = v59
                                                }
                                            case 128:
                                                if false {
                                                    fmt.Println(v6)
                                                }
                                            default:
                                                var v64 string = (("a" + "a") + ("" + "foo bar"))
                                                _ = v64
                                            }
                                        }
                                    }
                                    var v65 float32 = (1e10 - 0.0)
                                    _ = v31
                                    _ = v65
                                }
                                v66 := []int{(fi12 / ((v6) | 1)), 652}
                                var v67 uint16 = 413
                                var v68 int32 = (v1 - v1)
                                switch fi8 {
                                case 9223372036854775807:
                                    if true {
                                        var v69 bool = (fi8 <= int(v2))
                                        v1 -= v68
                                        if len(v66) > 0 {
                                            v66[0] = (fi12 % ((4294967295) | 1))
                                        }
                                        _ = v69
                                    } else {
                                        v70 := GenS11{F0: 42, F1: 700, F2: 42}
                                        if ((("foo bar" + ("" + "")) + "foo bar") == ("hello" + ((((("hello" + "\n\t") + "hello") + "hello") + "hello") + ""))) {
                                            var v71 int = (^(fi8))
                                            v72 := GenS1{F0: "\n\t"}
                                            if false {
                                                if false {
                                                    if true {
                                                        switch "\n\t" {
                                                        case "a":
                                                            if ((true || (0.1 >= 1.5)) || ((v3 + (v3 >> ((v3) & 63))) > v3)) {
                                                                v73 := GenS12{F0: 0}
                                                                v74 := GenS12{F0: 2}
                                                                fmt.Println(v3, v68, v67)
                                                                if (!false) {
                                                                    var v75 int = (v6 * 750)
                                                                    _ = v75
                                                                }
                                                                _ = v73
                                                                _ = v74
                                                            } else {
                                                                var v76 int16 = int16(v3)
                                                                var v77 int64 = -1
                                                                _ = v76
                                                                _ = v77
                                                            }
                                                        case "hello":
                                                            var v78 int8 = -128
                                                            _ = v78
                                                        default:
                                                            var v79 uint64 = 1
                                                            _ = v79
                                                        }
                                                    } else {
                                                        var v80 float64 = v5
                                                        _ = v80
                                                    }
                                                    var v81 uint16 = v67
                                                    switch ("\n\t" + "hello") {
                                                    case "a":
                                                        var v82 int = 2147483647
                                                        _ = v82
                                                    case "foo bar":
                                                        fmt.Println(v2)
                                                    default:
                                                        if (v2 <= v2) {
                                                            v83 := GenS3{F0: "foo bar", F1: (v71 / ((1) | 1))}
                                                            v84 := GenS9{F0: v2, F1: 1, F2: 671}
                                                            _ = v83
                                                            _ = v84
                                                        }
                                                    }
                                                    _ = v81
                                                } else {
                                                    if false {
                                                        v85 := GenS4{F0: 2}
                                                        v68 = (v68 - (v68 >> ((521) & 63)))
                                                        _ = v85
                                                    }
                                                }
                                                v71 -= (v71 * fi12)
                                            }
                                            _ = v71
                                            _ = v72
                                        } else {
                                            v4[2147483648] = ((("\n\t" + "a") + "hello") + "a")
                                        }
                                        v86 := GenS10{F0: 1}
                                        _ = v70
                                        _ = v86
                                    }
                                    v3 = (v3 | 42)
                                    v68 += v68
                                }
                                _ = v66
                                _ = v67
                                _ = v68
                            } else {
                                switch v6 {
                                case -185:
                                    var v87 uint16 = 1
                                    _ = v87
                                case 1:
                                    if (18446744073709551615 == 4294967295) {
                                        fmt.Println(fi8, v3, v2)
                                        fmt.Println(fi12)
                                    }
                                }
                            }
                            v88 := v4[837]
                            v89 := []int{-649, (fi8 / (((fi8 / ((fi12) | 1))) | 1)), v6}
                            _ = v88
                            _ = v89
                        }
                    default:
                        var v90 int32 = -124
                        _ = v90
                    }
                }
                switch (fi12 << ((v3) & 63)) {
                case -2147483648:
                    fmt.Println(v2)
                case 4294967295:
                    switch "" {
                    case "hello":
                        if ((v1 ^ (v1 >> (((v3 / ((2) | 1))) & 63))) >= 217) {
                            v3 = (^(v3))
                            delete(v4, 1)
                        }
                    default:
                        fmt.Println(fi8, fi12)
                    }
                default:
                    v2 = (^(v2))
                }
                v4[2147483647] = "\n\t"
                _ = fi12
            }
        default:
            for fi91 := 0; fi91 < 13; fi91++ {
                v2 += (v2 << ((v3) & 63))
                fmt.Println(v6)
                var v92 uint64 = 4294967295
                _ = fi91
                _ = v92
            }
        }
        _ = fi8
    }
    v1 = 2147483647
    v6 += -9223372036854775808
    v2 += (v2 * (v2 << ((v3) & 63)))
    var v93 int16 = (v2 / (((v2 - v2)) | 1))
    if (128 == 128) {
        for fi94 := 0; fi94 < 4; fi94++ {
            var v95 string = ""
            v96 := []int{}
            for fi97 := 0; fi97 < 9; fi97++ {
                v96 = append(v96, fi97)
                _ = fi97
            }
            _ = fi94
            _ = v95
            _ = v96
        }
        if (true || true) {
            fmt.Println(v1, v2)
            if false {
                v98 := GenS13{F0: 1}
                _ = v98
            }
        } else {
            var v99 bool = ((true || (false && false)) || false)
            var v100 uint = (^(v3))
            _ = v100
            _ = v99
        }
    } else {
        v5 = (v5 / -2.25)
    }
    v4[1] = (((("foo bar" + "hello") + ("foo bar" + "\n\t")) + ("\n\t" + "hello")) + ("" + ""))
    var v101 float64 = (0.1 / float64(0.1))
    if (v2 >= v2) {
        var v102 uint16 = 255
        if (-2147483648 >= (v1 / (((v1 | -154)) | 1))) {
            v3 += v3
        } else {
            var v103 float32 = (1.5 * (3.14159 - 0.1))
            v3 += 255
            if false {
                v104 := GenS12{F0: v3}
                switch (v6 % ((678) | 1)) {
                case -9223372036854775808:
                    var v105 int16 = -98
                    _ = v105
                case 0:
                    switch "" {
                    case "foo bar":
                        var v106 bool = true
                        _ = v106
                    }
                case 2147483647:
                    var v107 float64 = 0.0
                    _ = v107
                }
                _ = v104
            } else {
                var v108 float64 = (-2.25 - (1.5 * v101))
                switch (v6 % (((v6 ^ 127)) | 1)) {
                case 0:
                    v109 := GenS1{F0: "\n\t"}
                    _ = v109
                case 4294967295:
                    fmt.Println(v103)
                default:
                    for fi110 := 0; fi110 < 6; fi110++ {
                        _ = len(v4)
                        var v111 int32 = (v1 ^ v1)
                        _ = fi110
                        _ = v111
                    }
                }
                if true {
                    fmt.Println(v103, v3)
                    v112 := GenS10{F0: 18446744073709551615}
                    _ = v112
                }
                _ = v108
            }
            _ = v103
        }
        _ = v102
    } else {
        fmt.Println(v1, v3, v3)
        v93 = -660
        var v113 int16 = v93
        _ = v113
    }
    v114 := GenS7{F0: v93, F1: (v1 | 1)}
    for fi115 := 0; fi115 < 5; fi115++ {
        v116 := GenS14{F0: 0, F1: v1}
        _ = fi115
        _ = v116
    }
    var v117 int16 = v2
    if (!false) {
        switch (v6 + (v6 - v6)) {
        case -9223372036854775808:
            v118 := v4[0]
            _ = v118
        case 2147483647:
            for fi119 := 0; fi119 < 15; fi119++ {
                if true {
                    v120 := []int{(-(v6)), (fi119 / ((fi119) | 1)), 75}
                    v101 = v101
                    fmt.Println(v3, fi119, fi119)
                    _ = v120
                }
                var v121 int8 = -9
                _ = fi119
                _ = v121
            }
        default:
            v122 := GenS2{F0: v6, F1: 213}
            _ = v122
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v101:", v101)
    fmt.Println("v117:", v117)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v93:", v93)
    _ = v1
    _ = v101
    _ = v114
    _ = v117
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v93
}
