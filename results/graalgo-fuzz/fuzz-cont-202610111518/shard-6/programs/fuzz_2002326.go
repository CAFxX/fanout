package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string; F2 int }
type GenS2 struct { F0 uint8; F1 string }
type GenS3 struct { F0 uint16; F1 int8 }
type GenS4 struct { F0 int64; F1 int32 }
type GenS5 struct { F0 uint; F1 string; F2 string }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 string }
type GenS9 struct { F0 uint32; F1 int8 }
type GenS10 struct { F0 int32; F1 int64 }
func genFn0(p0 uint, p1 uint32, p2 uint16) uint8 {
    var v4 float64 = 1e10
    v5 := GenS1{F0: 9223372036854775807, F1: "hello", F2: -2147483648}
    var v6 int32 = 773
    var v7 int = 32768
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return 139
}
func genFn1(p0 int16, p1 uint, p2 uint64) int {
    for fi8 := 0; fi8 < 15; fi8++ {
        v9 := GenS1{F0: 9223372036854775807, F1: "", F2: (^(fi8))}
        _ = fi8
        _ = v9
    }
    p1 -= (p1 + p1)
    _ = p0
    _ = p1
    _ = p2
    return 32768
}

func main() {
    var v1 uint = 255
    var v2 int32 = -421
    v3 := make(map[int]uint32)
    fmt.Println(v2, v1)
    var v10 string = "a"
    switch 127 {
    case -920:
        v11 := []int{-202}
        _ = v11
    case 32767:
        var v12 string = "foo bar"
        _ = v12
    }
    fmt.Println(v10)
    fmt.Println(v2, v10, v2)
    if (!((-2.25 > 1e10) || (1e10 <= 1.5))) {
        var v13 uint8 = 102
        v14 := GenS1{F0: 42, F1: "foo bar", F2: 4294967295}
        v15 := GenS2{F0: (^(v13)), F1: v10}
        _ = v13
        _ = v14
        _ = v15
    } else {
        for fi16 := 0; fi16 < 15; fi16++ {
            switch (fi16 - fi16) {
            case 4294967295:
                v3[fi16] = 2
            case -2147483648:
                v2 -= v2
            default:
                for fi17 := 0; fi17 < 10; fi17++ {
                    v3[fi17] = 42
                    delete(v3, fi17)
                    v18 := GenS3{F0: 1, F1: 127}
                    delete(v3, (fi17 << ((v1) & 63)))
                    _ = fi17
                    _ = v18
                }
            }
            _ = fi16
        }
        for fi19 := 0; fi19 < 13; fi19++ {
            var v20 float64 = float64(0.1)
            _ = fi19
            _ = v20
        }
    }
    v21 := []int{1, genFn1(1000, 4294967295, 4294967295), 127, -705}
    var v22 string = "hello"
    v22 = ((("\n\t" + v10) + (("\n\t" + v10) + "a")) + v22)
    if true {
        for fi23 := 0; fi23 < 6; fi23++ {
            v22 = v10
            _ = fi23
        }
        v22 = v10
        for fi24 := 0; fi24 < 11; fi24++ {
            if (919 < 1) {
                v1 -= 18446744073709551615
                v3[9223372036854775807] = 4294967295
                v25 := GenS2{F0: 189, F1: ("" + ("a" + v10))}
                _ = v25
            } else {
                var v26 int = fi24
                for fi27 := 0; fi27 < 5; fi27++ {
                    switch "" {
                    case "foo bar":
                        v2 = v2
                    default:
                        var v28 int16 = 706
                        var v29 int8 = -29
                        _ = v28
                        _ = v29
                    }
                    fmt.Println(v26)
                    var v30 uint64 = 4294967295
                    _ = fi27
                    _ = v30
                }
                if (v1 != v1) {
                    fmt.Println(v10, v26, v1)
                    var v31 uint = (v1 - v1)
                    _ = v31
                } else {
                    v26 -= v26
                    if ((v1 - v1) > v1) {
                        var v32 int64 = -1
                        switch fi24 {
                        case 4294967295:
                            v33 := v3[(fi24 + (v26 / (((fi24 * v26)) | 1)))]
                            _ = v33
                        case 901:
                            for fi34 := 0; fi34 < 20; fi34++ {
                                switch v22 {
                                case "\n\t":
                                    v35 := GenS4{F0: v32, F1: 42}
                                    _ = v35
                                case "foo bar":
                                    v1 = v1
                                }
                                v10 = "hello"
                                v10 = ("foo bar" + ((v22 + v10) + v10))
                                _ = fi34
                            }
                        }
                        _ = v32
                    } else {
                        if true {
                            v36 := v3[0]
                            if true {
                                v10 = ("a" + "\n\t")
                                var v37 uint32 = v36
                                if true {
                                    if true {
                                        v38 := GenS1{F0: 42, F1: v10, F2: (fi24 * v26)}
                                        _ = v38
                                    }
                                    var v39 float32 = 0.0
                                    _ = v39
                                } else {
                                    var v40 uint32 = (v37 + v37)
                                    _ = v40
                                }
                                _ = v37
                            }
                            v41 := GenS5{F0: (v1 - (v1 | v1)), F1: v22, F2: v10}
                            _ = v36
                            _ = v41
                        } else {
                            var v42 uint8 = 2
                            var v43 string = "hello"
                            if false {
                                if false {
                                    switch 1 {
                                    case -373:
                                        var v44 int32 = v2
                                        _ = v44
                                    case -505:
                                        for fi45 := 0; fi45 < 16; fi45++ {
                                            v42 = v42
                                            _ = fi45
                                        }
                                    case 1:
                                        fmt.Println(fi24, fi24)
                                    }
                                    v46 := GenS6{F0: 507}
                                    switch v10 {
                                    case "":
                                        v47 := GenS2{F0: uint8(v1), F1: "a"}
                                        _ = v47
                                    default:
                                        switch ((v22 + ("\n\t" + (v22 + (v22 + v22)))) + v10) {
                                        case "hello":
                                            if ((1.5 - 0.0) < (((1.5 / (-2.25 * (1e10 * 1.5))) - 1.5) / 3.14159)) {
                                                switch v43 {
                                                case "foo bar":
                                                    fmt.Println(v26)
                                                case "":
                                                    for fi48 := 0; fi48 < 3; fi48++ {
                                                        v10 = v43
                                                        _ = fi48
                                                    }
                                                case "a":
                                                    v49 := GenS1{F0: -9223372036854775808, F1: (("foo bar" + ("foo bar" + (v10 + v43))) + ""), F2: (-(fi24))}
                                                    _ = v49
                                                default:
                                                    if (!false) {
                                                        var v50 int32 = -994
                                                        var v51 int8 = 16
                                                        if ((4294967295 != 4294967295) && false) {
                                                            v50 -= 0
                                                            for fi52 := 0; fi52 < 16; fi52++ {
                                                                var v53 int64 = 85
                                                                _ = fi52
                                                                _ = v53
                                                            }
                                                        }
                                                        _ = v50
                                                        _ = v51
                                                    }
                                                }
                                                if (!false) {
                                                    if true {
                                                        if true {
                                                            v43 = v43
                                                            for fi54 := 0; fi54 < 18; fi54++ {
                                                                switch (fi54 / (((v26 ^ 128)) | 1)) {
                                                                case 369:
                                                                    var v55 uint8 = 48
                                                                    _ = v55
                                                                case 32768:
                                                                    _ = len(v3)
                                                                case 127:
                                                                    var v56 uint32 = 2
                                                                    _ = v56
                                                                default:
                                                                    var v57 int = v26
                                                                    _ = v57
                                                                }
                                                                _ = fi54
                                                            }
                                                        } else {
                                                            if false {
                                                                for fi58 := 0; fi58 < 19; fi58++ {
                                                                    if ((uint(v42) <= (v1 * 42)) && (true && ((!false) || false))) {
                                                                        if len(v21) > 0 {
                                                                            v21[0] = 9223372036854775807
                                                                        }
                                                                    } else {
                                                                        var v59 int64 = -1
                                                                        v10 = "foo bar"
                                                                        switch (v26 + fi58) {
                                                                        case 2147483647:
                                                                            v60 := GenS1{F0: (v59 % ((v59) | 1)), F1: v22, F2: (fi24 - (fi58 ^ genFn1(-544, v1, 1)))}
                                                                            _ = v60
                                                                        case 128:
                                                                            if ((1.5 < 1e10) && (((v59 * 1) >= -109) || false)) {
                                                                                var v61 uint64 = 42
                                                                                _ = v61
                                                                            }
                                                                        }
                                                                        if (112 < 75) {
                                                                            var v62 float32 = (1e10 - 1.5)
                                                                            if false {
                                                                                if false {
                                                                                    if (((fi58 ^ int(v42)) < (v26 % ((42) | 1))) && (!false)) {
                                                                                        var v63 uint = v1
                                                                                        _ = v63
                                                                                    } else {
                                                                                        var v64 string = "foo bar"
                                                                                        v59 -= v59
                                                                                        _ = v64
                                                                                    }
                                                                                } else {
                                                                                    var v65 float32 = 0.1
                                                                                    _ = v65
                                                                                }
                                                                            }
                                                                            _ = v62
                                                                        } else {
                                                                            v66 := GenS2{F0: (v42 & v42), F1: (v10 + (v43 + ("" + "foo bar")))}
                                                                            _ = v66
                                                                        }
                                                                        _ = v59
                                                                    }
                                                                    fmt.Println(v10)
                                                                    var v67 float64 = 0.0
                                                                    _ = fi58
                                                                    _ = v67
                                                                }
                                                            }
                                                        }
                                                        for fi68 := 0; fi68 < 8; fi68++ {
                                                            switch v10 {
                                                            case "foo bar":
                                                                v69 := GenS6{F0: 2}
                                                                _ = v69
                                                            default:
                                                                if false {
                                                                    var v70 int16 = 633
                                                                    _ = v70
                                                                } else {
                                                                    v43 = v43
                                                                    v2 = (v2 + 42)
                                                                    fmt.Println(fi68)
                                                                    v71 := GenS6{F0: 13}
                                                                    _ = v71
                                                                }
                                                            }
                                                            _ = fi68
                                                        }
                                                    } else {
                                                        var v72 float64 = float64((3.14159 / -2.25))
                                                        switch fi24 {
                                                        case -2147483648:
                                                            var v73 bool = false
                                                            _ = v73
                                                        case 1:
                                                            var v74 int16 = 707
                                                            _ = v74
                                                        case 32768:
                                                            v42 = v42
                                                        default:
                                                            _ = len(v21)
                                                        }
                                                        fmt.Println(v72, v42)
                                                        _ = v72
                                                    }
                                                    v42 += v42
                                                    for fi75 := 0; fi75 < 2; fi75++ {
                                                        var v76 int16 = 42
                                                        _ = fi75
                                                        _ = v76
                                                    }
                                                } else {
                                                    if (((-2.25 / (((0.1 + float32(1e10)) / -2.25) + (1e10 * (1.5 * 3.14159)))) == (0.1 + 0.1)) || (("" + (v10 + "foo bar")) < (v10 + "\n\t"))) {
                                                        v10 = ((v10 + v22) + "foo bar")
                                                        v10 = "\n\t"
                                                    } else {
                                                        var v77 int = 4294967295
                                                        _ = v77
                                                    }
                                                    switch (v26 * (v26 ^ v26)) {
                                                    case 296:
                                                        var v78 int = genFn1(-275, 0, 2)
                                                        _ = v78
                                                    case 42:
                                                        for fi79 := 0; fi79 < 17; fi79++ {
                                                            v22 = v10
                                                            v43 = "hello"
                                                            var v80 uint = 515
                                                            _ = fi79
                                                            _ = v80
                                                        }
                                                    case -1:
                                                        if (!false) {
                                                            v81 := GenS1{F0: -770, F1: (("" + ((v43 + "a") + v22)) + "hello"), F2: v26}
                                                            v82 := GenS7{F0: 323}
                                                            v83 := GenS8{F0: ("foo bar" + ((((("hello" + v10) + "") + (v43 + ((v22 + "") + v22))) + (v43 + v22)) + (v22 + "\n\t")))}
                                                            _ = v81
                                                            _ = v82
                                                            _ = v83
                                                        }
                                                    default:
                                                        v84 := GenS3{F0: 1, F1: 100}
                                                        _ = v84
                                                    }
                                                    switch v43 {
                                                    case "\n\t":
                                                        v43 = v10
                                                    }
                                                }
                                            }
                                        case "a":
                                            for fi85 := 0; fi85 < 4; fi85++ {
                                                if (true && false) {
                                                    _ = len(v3)
                                                } else {
                                                    v86 := GenS9{F0: 65535, F1: 100}
                                                    var v87 uint = (v1 % (((v1 * v1)) | 1))
                                                    switch "foo bar" {
                                                    case "foo bar":
                                                        v87 -= 4294967295
                                                    case "":
                                                        v88 := v3[fi85]
                                                        var v89 uint64 = 575
                                                        _ = v88
                                                        _ = v89
                                                    case "\n\t":
                                                        if (((v22 + "\n\t") + ((("" + "foo bar") + "\n\t") + v10)) == "") {
                                                            v10 = ""
                                                        } else {
                                                            var v90 float64 = 3.14159
                                                            delete(v3, (fi24 ^ 127))
                                                            _ = v90
                                                        }
                                                        if (1 <= 0) {
                                                            v43 = (v22 + (v43 + v43))
                                                        }
                                                        if true {
                                                            var v91 uint = v1
                                                            _ = len(v3)
                                                            _ = v91
                                                        }
                                                    }
                                                    _ = v86
                                                    _ = v87
                                                }
                                                fmt.Println(fi85, fi85)
                                                var v92 int64 = -1
                                                _ = fi85
                                                _ = v92
                                            }
                                        case "":
                                            fmt.Println(v42)
                                        }
                                    }
                                    _ = v46
                                } else {
                                    var v93 uint8 = (v42 >> ((v1) & 63))
                                    var v94 uint32 = 2
                                    v21 = append(v21, genFn1(-32768, 845, 1))
                                    _ = v93
                                    _ = v94
                                }
                            }
                            _ = v42
                            _ = v43
                        }
                    }
                    fmt.Println(fi24, fi24)
                }
                _ = v26
            }
            v95 := v3[(fi24 * fi24)]
            var v96 uint16 = 65535
            _ = fi24
            _ = v95
            _ = v96
        }
    } else {
        v97 := GenS7{F0: 502}
        fmt.Println(v10)
        var v98 bool = false
        _ = v97
        _ = v98
    }
    var v99 int32 = v2
    var v100 uint64 = 557
    for fi101 := 0; fi101 < 6; fi101++ {
        fmt.Println(v2, v100)
        var v102 uint32 = 2147483648
        if true {
            switch ((("" + "foo bar") + v10) + "hello") {
            case "hello":
                fmt.Println(fi101, v100, v1)
            }
        } else {
            switch v10 {
            case "foo bar":
                v103 := GenS10{F0: int32(v100), F1: int64(fi101)}
                _ = v103
            default:
                if false {
                    var v104 uint64 = 171
                    v104 = v100
                    switch -2147483648 {
                    case 128:
                        if ("" == v22) {
                            for fi105 := 0; fi105 < 3; fi105++ {
                                v10 = v22
                                _ = fi105
                            }
                        } else {
                            switch "hello" {
                            case "hello":
                                switch v22 {
                                case "\n\t":
                                    switch (fi101 - (fi101 >> (((v1 << ((v1) & 63))) & 63))) {
                                    case 1:
                                        var v106 int = -2147483648
                                        _ = v106
                                    default:
                                        fmt.Println(v104)
                                    }
                                default:
                                    v2 = -2147483648
                                }
                            case "a":
                                v107 := v3[fi101]
                                _ = v107
                            case "":
                                var v108 float64 = (3.14159 * float64(1.5))
                                _ = v108
                            default:
                                if true {
                                    fmt.Println(v1)
                                }
                            }
                            switch (v10 + v10) {
                            case "\n\t":
                                var v109 uint = v1
                                _ = v109
                            case "":
                                if false {
                                    v100 -= v100
                                } else {
                                    v110 := GenS10{F0: v2, F1: -909}
                                    fmt.Println(v99)
                                    _ = v110
                                }
                            case "foo bar":
                                v100 = (^(v104))
                            default:
                                v104 += v104
                            }
                        }
                    default:
                        switch fi101 {
                        case -929:
                            if false {
                                v1 += v1
                                if (32767 != 1) {
                                    switch -762 {
                                    case 32768:
                                        v104 = 4294967295
                                    }
                                    fmt.Println(fi101, v104)
                                    for fi111 := 0; fi111 < 18; fi111++ {
                                        switch fi101 {
                                        case 362:
                                            var v112 uint = 0
                                            _ = v112
                                        }
                                        v10 = v10
                                        var v113 int16 = 0
                                        _ = fi111
                                        _ = v113
                                    }
                                } else {
                                    if true {
                                        if ("a" != ((v22 + "hello") + "a")) {
                                            var v114 int = genFn1(191, v1, (v100 | v100))
                                            v115 := GenS7{F0: v104}
                                            if len(v21) >= 2 {
                                                v21 = v21[1:len(v21)]
                                            }
                                            _ = v114
                                            _ = v115
                                        }
                                        fmt.Println(v100, fi101)
                                        _ = len(v21)
                                    }
                                    if false {
                                        var v116 int = genFn1(-132, v1, (v104 - uint64(v1)))
                                        fmt.Println(v99, v10)
                                        var v117 uint64 = (^(v100))
                                        _ = v116
                                        _ = v117
                                    }
                                    var v118 uint16 = 528
                                    _ = v118
                                }
                            }
                        }
                    }
                    _ = v104
                }
            }
            fmt.Println(v2, v10)
            for fi119 := 0; fi119 < 15; fi119++ {
                v99 -= v99
                var v120 uint32 = uint32(v100)
                v120 += v102
                _ = fi119
                _ = v120
            }
        }
        _ = fi101
        _ = v102
    }
    v22 = (("foo bar" + v10) + "hello")
    for fi121 := 0; fi121 < 15; fi121++ {
        var v122 int64 = -371
        _ = fi121
        _ = v122
    }
    v10 = "a"
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v100:", v100)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v99:", v99)
    _ = v1
    _ = v10
    _ = v100
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v99
}
