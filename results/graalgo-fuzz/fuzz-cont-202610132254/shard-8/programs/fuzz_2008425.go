package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8; F2 int32 }
type GenS2 struct { F0 bool; F1 uint16 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint8; F1 uint; F2 uint16 }
type GenS5 struct { F0 string; F1 uint8; F2 uint32 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 int8; F1 string; F2 int16 }
type GenS8 struct { F0 int8; F1 uint; F2 int64 }
type GenS9 struct { F0 uint32; F1 uint16 }
func genFn0(p0 int16) bool {
    for fi7 := 0; fi7 < 18; fi7++ {
        var v8 uint64 = 18446744073709551615
        _ = fi7
        _ = v8
    }
    for fi9 := 0; fi9 < 8; fi9++ {
        v10 := GenS1{F0: ("foo bar" + "\n\t"), F1: 255, F2: 465}
        if (!false) {
            fmt.Println(fi9, p0)
        }
        fmt.Println(fi9)
        _ = fi9
        _ = v10
    }
    var v11 int16 = 1000
    var v12 uint = 798
    v13 := []int{0, 32768, 32767, 32767}
    _ = p0
    _ = v11
    _ = v12
    _ = v13
    return false
}

func main() {
    var v1 bool = false
    var v2 uint = 735
    var v3 int = -1
    v4 := []int{32767, v3, -2147483648}
    var v5 int16 = 1
    v6 := make(map[uint32]int)
    v14 := GenS1{F0: "hello", F1: 128, F2: 2147483647}
    fmt.Println(v2, v1, v3)
    fmt.Println(v1, v2)
    var v15 float64 = (((0.0 - 3.14159) * 0.0) / 1e10)
    var v16 int8 = -100
    if (v1 || (!false)) {
        for fi17 := 0; fi17 < 12; fi17++ {
            delete(v6, 4294967295)
            v5 += (^(v5))
            _ = fi17
        }
    } else {
        var v18 float32 = 0.0
        _ = v18
    }
    v5 = 32767
    var v19 uint = v2
    fmt.Println(v19)
    for fi20 := 0; fi20 < 2; fi20++ {
        var v21 uint32 = 662
        delete(v6, v21)
        _ = fi20
        _ = v21
    }
    if ((0 > 1) || false) {
        v22 := GenS1{F0: ("a" + ("a" + "\n\t")), F1: 1, F2: 1}
        switch v3 {
        case -1:
            switch v3 {
            case 1:
                v6[8] = (v3 - (v3 << ((v19) & 63)))
            }
        }
        _ = v22
    }
    switch ("\n\t" + "a") {
    case "a":
        delete(v6, 2147483647)
    case "":
        v3 += (v3 - int(v16))
    case "hello":
        v3 -= v3
    default:
        if false {
            for fi23 := 0; fi23 < 9; fi23++ {
                if v1 {
                    for fi24 := 0; fi24 < 5; fi24++ {
                        v3 -= 0
                        _ = fi24
                    }
                }
                v15 -= 1e10
                var v25 int = (v3 % ((fi23) | 1))
                _ = fi23
                _ = v25
            }
        }
    }
    var v26 bool = ((v3 | v3) >= v3)
    _ = len(v6)
    for fi27 := 0; fi27 < 10; fi27++ {
        switch ("" + (("hello" + "foo bar") + (("foo bar" + "foo bar") + ""))) {
        case "\n\t":
            var v28 uint16 = 1
            _ = v28
        case "foo bar":
            switch (("" + "hello") + "foo bar") {
            case "":
                v5 += v5
            case "\n\t":
                v29 := GenS1{F0: (((("foo bar" + "") + "hello") + "hello") + ""), F1: 128, F2: 0}
                _ = v29
            case "hello":
                v26 = v26
            }
        case "a":
            if (v1 || (!false)) {
                var v30 int16 = v5
                v6[42] = (v3 >> ((v2) & 63))
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                _ = v30
            }
        }
        switch "foo bar" {
        case "a":
            if true {
                var v31 uint64 = 42
                _ = v31
            }
        case "hello":
            if v26 {
                for fi32 := 0; fi32 < 15; fi32++ {
                    v5 = v5
                    _ = fi32
                }
                switch fi27 {
                case -1:
                    var v33 uint16 = 255
                    _ = v33
                case 32767:
                    var v34 int8 = (v16 << ((v19) & 63))
                    _ = v34
                case 32768:
                    if (42 == 4294967295) {
                        switch "foo bar" {
                        case "\n\t":
                            switch fi27 {
                            case -2147483648:
                                var v35 float64 = 0.1
                                _ = v35
                            case 127:
                                for fi36 := 0; fi36 < 7; fi36++ {
                                    v15 += v15
                                    var v37 float32 = 0.0
                                    fmt.Println(v1, v19)
                                    _ = fi36
                                    _ = v37
                                }
                            default:
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                            }
                        default:
                            for fi38 := 0; fi38 < 4; fi38++ {
                                switch 127 {
                                case -973:
                                    switch (("a" + ("foo bar" + "hello")) + "hello") {
                                    case "foo bar":
                                        var v39 uint64 = 42
                                        _ = v39
                                    default:
                                        if (!v26) {
                                            if v26 {
                                                if len(v4) >= 2 {
                                                    v4 = v4[1:len(v4)]
                                                }
                                                v3 -= 1
                                            } else {
                                                if v1 {
                                                    v26 = (0.1 < float64(((0.0 * 0.0) * ((3.14159 - 3.14159) * (1e10 + (-2.25 + 1e10))))))
                                                    if v1 {
                                                        v40 := GenS2{F0: (v3 <= (fi27 - (fi38 & int(v19)))), F1: 1}
                                                        _ = v40
                                                    } else {
                                                        if (false || v26) {
                                                            var v41 float64 = v15
                                                            var v42 int16 = 32767
                                                            var v43 uint16 = 32767
                                                            _ = v41
                                                            _ = v42
                                                            _ = v43
                                                        }
                                                        fmt.Println(v16)
                                                        if v26 {
                                                            var v44 int8 = 0
                                                            _ = v44
                                                        } else {
                                                            v1 = (!v1)
                                                        }
                                                    }
                                                } else {
                                                    switch v3 {
                                                    case 32768:
                                                        if (!v1) {
                                                            var v45 int32 = -50
                                                            v46 := v6[437]
                                                            _ = v45
                                                            _ = v46
                                                        } else {
                                                            if (32767 == 42) {
                                                                v47 := GenS1{F0: "\n\t", F1: 243, F2: -2147483648}
                                                                _ = v47
                                                            } else {
                                                                switch 231 {
                                                                case 826:
                                                                    v5 -= (v5 << (((v2 - v19)) & 63))
                                                                case -369:
                                                                    v48 := GenS1{F0: "hello", F1: 122, F2: -1}
                                                                    _ = v48
                                                                }
                                                                var v49 float32 = 1.5
                                                                _ = v49
                                                            }
                                                            if v1 {
                                                                if ((1 > -569) || (v2 >= v19)) {
                                                                    v1 = v26
                                                                    v2 = (v19 / (((v2 & (v19 + (v2 * v2)))) | 1))
                                                                    fmt.Println(v15)
                                                                    var v50 uint8 = 128
                                                                    _ = v50
                                                                } else {
                                                                    v19 = v2
                                                                }
                                                                var v51 uint8 = 120
                                                                fmt.Println(v2, v5)
                                                                _ = v51
                                                            } else {
                                                                v52 := GenS1{F0: "", F1: 236, F2: 1}
                                                                v53 := v6[175]
                                                                switch (("a" + ("foo bar" + (((("foo bar" + (("\n\t" + ("a" + "hello")) + ("" + "a"))) + "hello") + "hello") + ("a" + "foo bar")))) + (("foo bar" + "\n\t") + ("a" + "hello"))) {
                                                                case "":
                                                                    var v54 uint32 = 1
                                                                    _ = v54
                                                                case "foo bar":
                                                                    switch v53 {
                                                                    case 127:
                                                                        var v55 int8 = (^(v16))
                                                                        _ = v55
                                                                    default:
                                                                        v56 := GenS3{F0: 65535}
                                                                        _ = v56
                                                                    }
                                                                }
                                                                _ = v52
                                                                _ = v53
                                                            }
                                                        }
                                                        var v57 int16 = 73
                                                        _ = v57
                                                    case 4294967295:
                                                        v16 -= -3
                                                    default:
                                                        v5 = v5
                                                    }
                                                    v16 -= 0
                                                    var v58 uint32 = 278
                                                    v59 := GenS4{F0: 241, F1: v2, F2: 255}
                                                    _ = v58
                                                    _ = v59
                                                }
                                            }
                                        }
                                    }
                                case 128:
                                    var v60 string = "\n\t"
                                    _ = v60
                                }
                                if len(v4) > 0 {
                                    v4[0] = -938
                                }
                                _ = fi38
                            }
                        }
                        for fi61 := 0; fi61 < 3; fi61++ {
                            var v62 int32 = 715
                            var v63 uint32 = 954
                            _ = fi61
                            _ = v62
                            _ = v63
                        }
                        if true {
                            fmt.Println(v15, v3)
                            for fi64 := 0; fi64 < 13; fi64++ {
                                v5 += int16(v16)
                                v15 -= v15
                                _ = fi64
                            }
                            fmt.Println(v5, fi27)
                        } else {
                            if true {
                                switch "hello" {
                                case "a":
                                    switch v3 {
                                    case -992:
                                        var v65 uint64 = 1
                                        _ = v65
                                    case -811:
                                        if (v1 || false) {
                                            v5 -= (v5 - v5)
                                            fmt.Println(v19)
                                        } else {
                                            fmt.Println(v1, v26)
                                            if len(v4) > 0 {
                                                v4[0] = (-(v3))
                                            }
                                            var v66 bool = (true || v1)
                                            _ = v66
                                        }
                                    case 1:
                                        v67 := GenS5{F0: "a", F1: 0, F2: 4294967295}
                                        _ = v67
                                    }
                                }
                            }
                            if true {
                                switch -9223372036854775808 {
                                case 127:
                                    fmt.Println(v26, v15, v26)
                                case -1:
                                    v2 -= (v2 / ((65535) | 1))
                                default:
                                    fmt.Println(v26)
                                }
                                if (uint8(v19) != 62) {
                                    fmt.Println(fi27, v5)
                                    var v68 int8 = v16
                                    var v69 string = "foo bar"
                                    _ = v68
                                    _ = v69
                                }
                                v2 -= (v2 & v2)
                            }
                            v16 -= -1
                        }
                    }
                }
            }
        default:
            for fi70 := 0; fi70 < 7; fi70++ {
                if v26 {
                    var v71 uint32 = 513
                    if ((!v26) || true) {
                        if true {
                            v72 := GenS6{F0: (v19 >> ((255) & 63))}
                            if ((float64(float32(v15)) - 1e10) >= v15) {
                                switch v3 {
                                case -9223372036854775808:
                                    v73 := GenS1{F0: ("" + ("" + ("\n\t" + "hello"))), F1: 42, F2: 2147483647}
                                    _ = v73
                                }
                                if v1 {
                                    var v74 int16 = v5
                                    v75 := GenS4{F0: 0, F1: (v2 & 499), F2: 65535}
                                    _ = v74
                                    _ = v75
                                }
                                fmt.Println(fi27, fi70)
                            } else {
                                switch "\n\t" {
                                case "a":
                                    switch ("" + "hello") {
                                    case "":
                                        v2 = 255
                                    case "a":
                                        var v76 uint16 = 1
                                        _ = v76
                                    case "hello":
                                        var v77 string = ""
                                        _ = v77
                                    }
                                case "":
                                    v6[(v71 * v71)] = (^(fi70))
                                case "hello":
                                    if (!v26) {
                                        var v78 int = (fi27 / (((v3 >> ((v2) & 63))) | 1))
                                        _ = v78
                                    } else {
                                        v79 := GenS1{F0: "foo bar", F1: 0, F2: 0}
                                        var v80 string = "foo bar"
                                        if (2 <= 18446744073709551615) {
                                            v81 := GenS1{F0: v80, F1: 255, F2: 636}
                                            _ = v81
                                        }
                                        _ = v79
                                        _ = v80
                                    }
                                default:
                                    v2 -= (^(v2))
                                    var v82 uint64 = uint64(v71)
                                    _ = v82
                                }
                            }
                            v83 := GenS5{F0: ("" + "a"), F1: 255, F2: 1}
                            _ = v72
                            _ = v83
                        } else {
                            switch ("foo bar" + "a") {
                            case "a":
                                var v84 uint8 = 42
                                _ = v84
                            case "":
                                fmt.Println(v3, fi27, v26)
                            }
                            var v85 float64 = v15
                            v86 := v6[v71]
                            _ = v85
                            _ = v86
                        }
                        v71 = (^(v71))
                    } else {
                        switch ("\n\t" + ((("a" + "\n\t") + "a") + "a")) {
                        case "hello":
                            switch (fi70 * (fi70 ^ fi70)) {
                            case 42:
                                v87 := v6[(v71 | v71)]
                                _ = v87
                            case 127:
                                v71 = 455
                            case 32768:
                                if true {
                                    v5 = (v5 % (((v5 << ((v19) & 63))) | 1))
                                }
                            }
                        default:
                            if true {
                                if v1 {
                                    v88 := GenS7{F0: v16, F1: "foo bar", F2: 474}
                                    var v89 string = ""
                                    v26 = (v26 || v26)
                                    _ = v88
                                    _ = v89
                                }
                            } else {
                                v90 := GenS8{F0: (-(v16)), F1: v19, F2: 0}
                                if true {
                                    if v1 {
                                        var v91 int32 = -217
                                        fmt.Println(v3)
                                        _ = v91
                                    } else {
                                        if v1 {
                                            fmt.Println(v15, v15)
                                            v2 += (v2 >> ((v2) & 63))
                                            var v92 uint64 = 18446744073709551615
                                            _ = v92
                                        } else {
                                            fmt.Println(fi27, v5, v16)
                                        }
                                        if v26 {
                                            fmt.Println(v5, v1, v16)
                                            v15 = 1.5
                                            v93 := GenS7{F0: (v16 | (v16 & (v16 + (v16 & int8(fi27))))), F1: "hello", F2: 1}
                                            _ = v93
                                        }
                                        _ = len(v4)
                                        var v94 float32 = 3.14159
                                        var v95 uint32 = (v71 % ((2147483647) | 1))
                                        _ = v94
                                        _ = v95
                                    }
                                    if (!v1) {
                                        _ = len(v4)
                                        v26 = (((606 != 583) || (4294967295 == 18446744073709551615)) || v1)
                                        v2 -= (v19 - v2)
                                        switch fi27 {
                                        case -555:
                                            var v96 uint32 = 2147483647
                                            _ = v96
                                        case 810:
                                            if (-429 > v5) {
                                                fmt.Println(v5)
                                            } else {
                                                if v1 {
                                                    v97 := GenS5{F0: "a", F1: 1, F2: 0}
                                                    v98 := GenS6{F0: 1}
                                                    v99 := v6[v71]
                                                    var v100 string = "\n\t"
                                                    _ = v100
                                                    _ = v97
                                                    _ = v98
                                                    _ = v99
                                                } else {
                                                    var v101 string = "\n\t"
                                                    v102 := GenS1{F0: v101, F1: 191, F2: 918}
                                                    _ = v101
                                                    _ = v102
                                                }
                                                var v103 uint = v19
                                                switch "a" {
                                                case "\n\t":
                                                    fmt.Println(v2, fi27, v3)
                                                case "":
                                                    v26 = ((v71 + (^(v71))) != 1)
                                                }
                                                _ = v103
                                            }
                                        case -283:
                                            v15 = float64(1.5)
                                        }
                                        v3 = (fi27 / ((fi27) | 1))
                                        v19 -= (v19 | (v19 >> (((^(v19))) & 63)))
                                    }
                                    _ = len(v6)
                                }
                                _ = v90
                            }
                        }
                        fmt.Println(v1)
                        if len(v4) > 0 {
                            v4[0] = fi27
                        }
                    }
                    v15 = 0.0
                    _ = v71
                } else {
                    var v104 string = ("\n\t" + "")
                    _ = v104
                }
                switch "" {
                case "":
                    var v105 uint64 = 1
                    _ = v105
                default:
                    if (-2.25 >= 0.1) {
                        v106 := v6[2]
                        if v26 {
                            v5 += (^(v5))
                        }
                        fmt.Println(v3)
                        _ = v106
                    }
                }
                if ("" < ("" + "hello")) {
                    if ("\n\t" == "foo bar") {
                        var v107 int8 = 1
                        var v108 int8 = (v107 ^ int8(v19))
                        var v109 uint64 = 714
                        _ = v107
                        _ = v108
                        _ = v109
                    }
                    var v110 float64 = -2.25
                    if false {
                        var v111 int32 = -2147483648
                        _ = v111
                    } else {
                        var v112 int = v3
                        _ = v112
                    }
                    var v113 int = v3
                    switch int(v16) {
                    case 801:
                        v113 -= (fi27 >> ((34) & 63))
                    case 1:
                        switch ("foo bar" + "foo bar") {
                        case "foo bar":
                            var v114 int = fi70
                            _ = v114
                        default:
                            v115 := GenS9{F0: 4294967295, F1: 835}
                            _ = v115
                        }
                    }
                    _ = v110
                    _ = v113
                } else {
                    if v26 {
                        v116 := GenS5{F0: "a", F1: 0, F2: 352}
                        v1 = v26
                        _ = v116
                    } else {
                        switch "" {
                        case "":
                            var v117 float64 = ((v15 - v15) * -2.25)
                            _ = v117
                        }
                    }
                    v2 -= (v19 >> ((2) & 63))
                }
                v19 = v19
                _ = fi70
            }
        }
        _ = fi27
    }
    var v118 int8 = v16
    fmt.Println(v1, v1, v26)
    var v119 float32 = (-2.25 * 3.14159)
    fmt.Println("v1:", v1)
    fmt.Println("v118:", v118)
    fmt.Println("v119:", v119)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v118
    _ = v119
    _ = v14
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
