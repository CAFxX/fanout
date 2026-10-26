package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string; F2 uint }
type GenS2 struct { F0 string; F1 bool; F2 uint16 }
type GenS3 struct { F0 uint64; F1 uint64 }
type GenS4 struct { F0 int64; F1 int64; F2 int8 }
type GenS5 struct { F0 int8; F1 string }
type GenS6 struct { F0 int32; F1 int8 }
type GenS7 struct { F0 string; F1 uint64; F2 bool }
type GenS8 struct { F0 int; F1 int }
type GenS9 struct { F0 uint64; F1 string }
type GenS10 struct { F0 string; F1 int64; F2 uint64 }
type GenS11 struct { F0 int64; F1 int16 }
type GenS12 struct { F0 uint64 }
func genFn0(p0 int8) uint16 {
    p0 = (p0 % (((p0 << ((42) & 63))) | 1))
    var v5 int64 = -374
    fmt.Println(v5, p0)
    var v6 int = 32768
    _ = p0
    _ = v5
    _ = v6
    return 0
}

func main() {
    v1 := make(map[int]uint32)
    var v2 uint16 = 255
    v3 := []uint32{1, 1}
    var v4 float64 = (1e10 * 0.0)
    for fi7 := 0; fi7 < 18; fi7++ {
        var v8 int8 = 1
        v9 := GenS1{F0: 255, F1: ("\n\t" + "\n\t"), F2: 0}
        if true {
            v10 := GenS1{F0: uint(v2), F1: ("foo bar" + ("hello" + ("" + ("\n\t" + "")))), F2: 971}
            switch "" {
            case "hello":
                switch "hello" {
                case "hello":
                    var v11 string = "\n\t"
                    _ = v11
                case "\n\t":
                    var v12 uint = 4294967295
                    _ = v12
                case "":
                    switch 32768 {
                    case 9223372036854775807:
                        if (864 != 681) {
                            fmt.Println(v2, v2)
                        }
                    case 0:
                        v13 := GenS1{F0: 65535, F1: "hello", F2: 0}
                        _ = v13
                    default:
                        if false {
                            if (2147483648 < 1) {
                                var v14 bool = (("foo bar" + "\n\t") >= ("" + ("a" + "hello")))
                                var v15 int = 915
                                _ = v14
                                _ = v15
                            } else {
                                if false {
                                    for fi16 := 0; fi16 < 7; fi16++ {
                                        v4 += v4
                                        _ = len(v3)
                                        var v17 int8 = v8
                                        _ = fi16
                                        _ = v17
                                    }
                                    fmt.Println(v4, v4, v2)
                                } else {
                                    v4 -= 3.14159
                                    switch "foo bar" {
                                    case "a":
                                        if true {
                                            delete(v1, fi7)
                                        } else {
                                            v18 := GenS1{F0: 1, F1: (("foo bar" + ("\n\t" + "")) + "\n\t"), F2: 400}
                                            _ = v18
                                        }
                                    default:
                                        var v19 uint8 = 1
                                        _ = v19
                                    }
                                }
                            }
                            switch "" {
                            case "":
                                v8 -= -128
                            }
                        }
                    }
                default:
                    v8 += (v8 * v8)
                }
            case "\n\t":
                var v20 uint32 = 1
                _ = v20
            case "a":
                if true {
                    var v21 uint64 = 42
                    var v22 int64 = -1
                    _ = v21
                    _ = v22
                } else {
                    v23 := GenS2{F0: "\n\t", F1: (!false), F2: v2}
                    var v24 uint16 = v2
                    v8 -= (v8 - (-(v8)))
                    _ = v23
                    _ = v24
                }
            default:
                v25 := v1[int(v8)]
                _ = v25
            }
            _ = v10
        } else {
            for fi26 := 0; fi26 < 5; fi26++ {
                var v27 uint8 = 255
                var v28 uint8 = v27
                if false {
                    v29 := GenS2{F0: (("\n\t" + "hello") + "foo bar"), F1: true, F2: 42}
                    switch "a" {
                    case "\n\t":
                        var v30 bool = false
                        _ = v30
                    }
                    v31 := GenS3{F0: 246, F1: 2}
                    _ = v29
                    _ = v31
                }
                v8 = -1
                _ = fi26
                _ = v27
                _ = v28
            }
            for fi32 := 0; fi32 < 5; fi32++ {
                var v33 float64 = v4
                if false {
                    if true {
                        fmt.Println(fi32)
                        v34 := GenS4{F0: 9223372036854775807, F1: 536, F2: v8}
                        _ = v34
                    } else {
                        v8 -= (v8 << ((4294967295) & 63))
                        v35 := GenS5{F0: 42, F1: ("\n\t" + ("" + "a"))}
                        _ = v35
                    }
                    v36 := GenS3{F0: 18446744073709551615, F1: 716}
                    _ = v36
                } else {
                    v37 := GenS2{F0: "", F1: (!(255 >= 0)), F2: (v2 & (v2 / ((v2) | 1)))}
                    _ = v37
                }
                var v38 float32 = -2.25
                _ = fi32
                _ = v33
                _ = v38
            }
        }
        _ = fi7
        _ = v8
        _ = v9
    }
    v2 = v2
    for fi39 := 0; fi39 < 4; fi39++ {
        v2 = genFn0(-1)
        _ = fi39
    }
    switch 127 {
    case 615:
        var v40 float64 = 3.14159
        _ = v40
    case -601:
        var v41 int32 = -2147483648
        _ = v41
    }
    if (126 == 255) {
        var v42 string = ("a" + "\n\t")
        _ = v42
    }
    for fi43 := 0; fi43 < 12; fi43++ {
        if true {
            v2 = 0
            for fi44 := 0; fi44 < 15; fi44++ {
                if (((fi44 * fi44) < (fi43 | 0)) || false) {
                    switch "a" {
                    case "\n\t":
                        v4 -= v4
                    case "hello":
                        v4 += v4
                        var v45 uint8 = 128
                        if (237 < (^(v45))) {
                            var v46 uint64 = 0
                            var v47 float32 = (((1e10 / -2.25) + 0.0) / 1.5)
                            _ = v46
                            _ = v47
                        }
                        _ = v45
                    case "foo bar":
                        switch ("\n\t" + "") {
                        case "hello":
                            if true {
                                _ = len(v1)
                            }
                        default:
                            v2 -= (^(v2))
                        }
                    }
                    switch fi44 {
                    case 127:
                        var v48 string = ""
                        _ = v48
                    case 9223372036854775807:
                        var v49 float32 = (1.5 * 3.14159)
                        _ = v49
                    case -532:
                        v50 := GenS3{F0: 1, F1: 0}
                        _ = v50
                    }
                    var v51 bool = ((fi43 * fi43) >= (fi43 * (fi43 ^ fi43)))
                    var v52 int32 = 44
                    switch fi43 {
                    case -22:
                        switch ("\n\t" + "\n\t") {
                        case "\n\t":
                            v51 = true
                        case "":
                            if (!true) {
                                var v53 int8 = 42
                                if ((!(true || true)) || ((v51 || false) && (((v53 / (((v53 % ((v53) | 1))) | 1)) > -128) || false))) {
                                    v54 := GenS5{F0: v53, F1: ("a" + "\n\t")}
                                    v51 = true
                                    _ = v54
                                } else {
                                    v52 += 2147483647
                                }
                                v51 = v51
                                _ = v53
                            } else {
                                var v55 int32 = (v52 ^ v52)
                                _ = v55
                            }
                        }
                    case -1:
                        switch (fi44 >> ((844) & 63)) {
                        case 2147483647:
                            delete(v1, fi44)
                        }
                        var v56 int16 = -32768
                        var v57 uint8 = 58
                        var v58 string = ((("foo bar" + "") + "a") + "a")
                        _ = v56
                        _ = v57
                        _ = v58
                    default:
                        if v51 {
                            var v59 int32 = (v52 / ((42) | 1))
                            if (v4 < ((float64(((3.14159 - (((3.14159 - -2.25) - ((float32(v4) + (0.0 * float32(0.0))) + float32(((-2.25 - -2.25) / ((v4 * 1.5) + (v4 + ((v4 * v4) / 3.14159))))))) + -2.25)) / (0.1 - 1e10))) - ((v4 + (((3.14159 - v4) + (v4 + v4)) * (((1.5 * ((v4 - float64(float32(v4))) - 1e10)) - float64(0.1)) - ((v4 + float64(0.1)) * v4)))) + 3.14159)) - ((((((float64(0.0) / 3.14159) - v4) * v4) + (0.0 * ((v4 - float64(1.5)) / v4))) / 1.0) / v4))) {
                                v52 = (v59 + v52)
                                if false {
                                    fmt.Println(v51, v52)
                                    v60 := GenS6{F0: -2147483648, F1: 4}
                                    _ = v60
                                }
                            }
                            v61 := v1[fi44]
                            _ = v59
                            _ = v61
                        }
                    }
                    _ = v51
                    _ = v52
                } else {
                    v62 := GenS4{F0: 42, F1: -1, F2: -50}
                    _ = len(v3)
                    v2 = 2
                    _ = v62
                }
                v63 := GenS3{F0: 898, F1: 0}
                _ = fi44
                _ = v63
            }
        } else {
            if false {
                if (fi43 >= (^(fi43))) {
                    for fi64 := 0; fi64 < 16; fi64++ {
                        v4 += 1.5
                        v65 := GenS7{F0: "\n\t", F1: 42, F2: (868 < 42)}
                        fmt.Println(v4)
                        _ = fi64
                        _ = v65
                    }
                    fmt.Println(fi43, v2, v2)
                    for fi66 := 0; fi66 < 6; fi66++ {
                        if (212 > 189) {
                            if false {
                                if (!true) {
                                    _ = len(v1)
                                }
                                if ((fi66 & 548) > 32767) {
                                    v67 := GenS1{F0: 4294967295, F1: "", F2: 0}
                                    if (fi43 == fi43) {
                                        v4 -= -2.25
                                        var v68 int32 = -397
                                        _ = v68
                                    } else {
                                        var v69 float32 = 3.14159
                                        v4 = v4
                                        v69 = v69
                                        _ = v69
                                    }
                                    v2 = (v2 & genFn0(42))
                                    var v70 int64 = 9223372036854775807
                                    _ = v67
                                    _ = v70
                                } else {
                                    var v71 string = ("\n\t" + "hello")
                                    var v72 float32 = 1.5
                                    if false {
                                        var v73 uint32 = 4294967295
                                        v71 = (("\n\t" + "") + ((v71 + "a") + "foo bar"))
                                        _ = v73
                                    } else {
                                        switch (fi43 / ((fi66) | 1)) {
                                        case -2147483648:
                                            v74 := GenS8{F0: (fi43 & fi66), F1: fi66}
                                            _ = v74
                                        case 274:
                                            var v75 int64 = 9223372036854775807
                                            _ = v75
                                        case 128:
                                            v76 := GenS9{F0: 1, F1: "\n\t"}
                                            _ = v76
                                        }
                                        v72 = float32(float64(3.14159))
                                        var v77 uint32 = 2
                                        _ = v77
                                    }
                                    _ = v71
                                    _ = v72
                                }
                                var v78 float32 = 0.1
                                _ = v78
                            }
                            v79 := v1[(fi43 | (-(fi43)))]
                            fmt.Println(v79, v4, v4)
                            _ = v79
                        }
                        _ = fi66
                    }
                }
                switch (("hello" + "a") + "foo bar") {
                case "":
                    v80 := GenS4{F0: 833, F1: -9223372036854775808, F2: -1}
                    _ = v80
                case "\n\t":
                    v4 += float64(float32(0.0))
                case "a":
                    var v81 uint32 = 197
                    _ = v81
                default:
                    if (!(!(fi43 <= (fi43 ^ fi43)))) {
                        for fi82 := 0; fi82 < 16; fi82++ {
                            var v83 uint8 = 198
                            if true {
                                v84 := GenS10{F0: "\n\t", F1: 1, F2: 4294967295}
                                _ = v84
                            } else {
                                v85 := GenS1{F0: 659, F1: ((("\n\t" + "foo bar") + "") + "\n\t"), F2: 4294967295}
                                var v86 int64 = 0
                                _ = len(v1)
                                delete(v1, fi43)
                                v87 := GenS6{F0: 0, F1: 9}
                                _ = v85
                                _ = v86
                                _ = v87
                            }
                            v83 = (v83 / (((v83 << ((65535) & 63))) | 1))
                            _ = fi82
                            _ = v83
                        }
                        v88 := GenS6{F0: -2147483648, F1: -50}
                        v89 := GenS9{F0: 0, F1: ""}
                        _ = v88
                        _ = v89
                    }
                }
                var v90 int64 = -261
                _ = v90
            } else {
                v4 = 0.1
                var v91 int8 = 44
                _ = v91
            }
            if (2 > 18446744073709551615) {
                for fi92 := 0; fi92 < 11; fi92++ {
                    if (((!(false && (0 < 865))) || (18446744073709551615 == 42)) || true) {
                        _ = len(v1)
                    }
                    _ = fi92
                }
                fmt.Println(v2, v4)
            } else {
                v2 += (v2 - genFn0(1))
                if true {
                    switch ("a" + "a") {
                    case "\n\t":
                        for fi93 := 0; fi93 < 2; fi93++ {
                            if true {
                                var v94 uint = 847
                                v94 = v94
                                var v95 uint = v94
                                _ = v94
                                _ = v95
                            }
                            fmt.Println(fi93, fi93, fi93)
                            switch 32767 {
                            case 32768:
                                v2 = genFn0(68)
                            case 798:
                                v4 += v4
                            }
                            _ = fi93
                        }
                    case "a":
                        var v96 float64 = 0.1
                        _ = v96
                    }
                    var v97 uint8 = 132
                    var v98 float64 = v4
                    _ = v97
                    _ = v98
                }
            }
        }
        switch (fi43 | (fi43 - fi43)) {
        case 4294967295:
            v99 := GenS11{F0: 1, F1: -433}
            _ = v99
        }
        _ = fi43
    }
    var v100 int32 = 270
    for fi101 := 0; fi101 < 3; fi101++ {
        var v102 string = ("a" + "hello")
        var v103 int = -243
        _ = fi101
        _ = v102
        _ = v103
    }
    v104 := GenS11{F0: 0, F1: -720}
    if (1 < 504) {
        fmt.Println(v100)
    }
    if true {
        v100 -= v100
    }
    for fi105 := 0; fi105 < 19; fi105++ {
        v106 := GenS10{F0: "\n\t", F1: 1, F2: 18446744073709551615}
        _ = fi105
        _ = v106
    }
    for fi107 := 0; fi107 < 1; fi107++ {
        fmt.Println(v100, v2)
        _ = fi107
    }
    switch 4294967295 {
    case 0:
        var v108 uint8 = 255
        _ = v108
    default:
        fmt.Println(v4, v4, v4)
    }
    if (genFn0(100) == (v2 - 42)) {
        for fi109 := 0; fi109 < 10; fi109++ {
            v110 := GenS12{F0: 18446744073709551615}
            v111 := GenS5{F0: 100, F1: ((("\n\t" + "hello") + "a") + "foo bar")}
            _ = fi109
            _ = v110
            _ = v111
        }
        for fi112 := 0; fi112 < 20; fi112++ {
            if true {
                var v113 int16 = -407
                var v114 uint8 = 1
                if (-9223372036854775808 == 42) {
                    switch "\n\t" {
                    case "a":
                        v113 = 259
                    case "\n\t":
                        v2 += (v2 * 255)
                    }
                    for fi115 := 0; fi115 < 6; fi115++ {
                        v4 = float64(-2.25)
                        _ = fi115
                    }
                    var v116 int = fi112
                    _ = v116
                }
                _ = v113
                _ = v114
            } else {
                v100 += v100
                var v117 int64 = -685
                v100 = v100
                _ = v117
            }
            _ = fi112
        }
        v100 += v100
    }
    delete(v1, 32768)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v100:", v100)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v100
    _ = v104
    _ = v2
    _ = v3
    _ = v4
}
