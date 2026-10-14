package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16 }
type GenS2 struct { F0 int16; F1 uint; F2 int64 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 bool; F1 int32; F2 uint8 }
type GenS5 struct { F0 string; F1 bool; F2 bool }
type GenS6 struct { F0 uint; F1 int8 }
type GenS7 struct { F0 int16; F1 uint64 }
type GenS8 struct { F0 int8 }
type GenS9 struct { F0 uint32; F1 int }
type GenS10 struct { F0 uint64; F1 int16 }
type GenS11 struct { F0 string; F1 int16; F2 uint16 }
type GenS12 struct { F0 string }
type GenS13 struct { F0 uint16 }
func genFn0(p0 int8, p1 uint16, p2 int16) uint64 {
    var v7 uint16 = 1
    v7 += p1
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 42
}

func main() {
    v1 := []int{-14, 165, 42}
    v2 := make(map[uint32]int)
    v3 := []string{"\n\t", "\n\t"}
    var v4 float32 = -2.25
    v5 := []int{463, 32767}
    var v6 int = 0
    if true {
        v8 := GenS1{F0: 18446744073709551615, F1: 0}
        _ = v8
    }
    v9 := GenS2{F0: 143, F1: 65535, F2: 42}
    if (!(9223372036854775807 > -9223372036854775808)) {
        var v10 int32 = 2147483647
        for fi11 := 0; fi11 < 7; fi11++ {
            if false {
                v10 += v10
                var v12 bool = true
                switch v6 {
                case -2147483648:
                    v10 -= v10
                }
                _ = v12
            }
            for fi13 := 0; fi13 < 19; fi13++ {
                v6 += (fi13 % ((v6) | 1))
                var v14 uint32 = 65535
                var v15 int = -547
                _ = fi13
                _ = v14
                _ = v15
            }
            _ = fi11
        }
        _ = v10
    }
    fmt.Println(v6, v4)
    for fi16 := 0; fi16 < 20; fi16++ {
        for fi17 := 0; fi17 < 14; fi17++ {
            switch (fi16 | fi16) {
            case -9223372036854775808:
                var v18 uint8 = 128
                _ = v18
            case 862:
                switch "a" {
                case "foo bar":
                    v6 = 42
                case "\n\t":
                    var v19 float64 = (0.1 + (float64(v4) + 0.0))
                    v6 -= fi17
                    var v20 int64 = -9223372036854775808
                    _ = v19
                    _ = v20
                case "a":
                    switch "\n\t" {
                    case "foo bar":
                        var v21 uint16 = 32767
                        _ = v21
                    case "\n\t":
                        v22 := GenS2{F0: -990, F1: 1, F2: -9223372036854775808}
                        _ = v22
                    case "":
                        v23 := GenS1{F0: 42, F1: 1}
                        _ = v23
                    }
                }
            }
            var v24 int32 = 1
            fmt.Println(v24, fi16)
            fmt.Println(fi17)
            v25 := GenS3{F0: 42}
            _ = fi17
            _ = v24
            _ = v25
        }
        switch ("\n\t" + ("foo bar" + "hello")) {
        case "foo bar":
            for fi26 := 0; fi26 < 18; fi26++ {
                fmt.Println(fi26, v6, fi26)
                var v27 uint32 = 991
                if (v4 >= v4) {
                    switch "a" {
                    case "a":
                        if (false || (true || ((false && false) || true))) {
                            fmt.Println(fi16, v6)
                            switch -912 {
                            case 219:
                                if ((!true) && (251 != 127)) {
                                    v27 -= v27
                                    v6 = (fi16 | (^(v6)))
                                    switch fi26 {
                                    case 42:
                                        v4 = (v4 - float32(float64(((v4 + v4) - (float32((-2.25 + (float64(float32((0.0 / 1.0))) * (((1.5 / 1.5) - 0.0) + 1.5)))) + (float32(0.1) / (v4 + -2.25)))))))
                                    case 970:
                                        var v28 int64 = -284
                                        _ = v28
                                    case -1:
                                        fmt.Println(v4, v27)
                                    }
                                } else {
                                    var v29 uint32 = v27
                                    _ = len(v3)
                                    _ = v29
                                }
                            case 374:
                                v30 := GenS4{F0: (127 != 127), F1: 42, F2: 0}
                                _ = v30
                            case 0:
                                var v31 float64 = -2.25
                                _ = v31
                            default:
                                v32 := GenS5{F0: "hello", F1: (!(65535 < 965)), F2: true}
                                _ = v32
                            }
                            var v33 int16 = 771
                            _ = v33
                        }
                    case "\n\t":
                        v27 = (^(v27))
                    case "foo bar":
                        var v34 string = "\n\t"
                        _ = v34
                    }
                    v35 := v2[v27]
                    var v36 int = 2147483647
                    _ = v35
                    _ = v36
                } else {
                    v37 := GenS2{F0: -114, F1: 42, F2: -1}
                    v38 := GenS6{F0: 2, F1: 43}
                    _ = v37
                    _ = v38
                }
                _ = fi26
                _ = v27
            }
        default:
            var v39 uint32 = 541
            _ = v39
        }
        for fi40 := 0; fi40 < 20; fi40++ {
            v41 := GenS2{F0: -611, F1: 2, F2: 42}
            v4 = 1.5
            _ = len(v5)
            _ = fi40
            _ = v41
        }
        _ = fi16
    }
    fmt.Println(v4, v4)
    v1 = append(v1, 85)
    if true {
        var v42 uint16 = 65535
        v6 += -2147483648
        switch "foo bar" {
        case "foo bar":
            if (-415 <= 42) {
                switch 614 {
                case 42:
                    var v43 int64 = -154
                    _ = v43
                case 85:
                    v6 += (v6 % (((v6 + v6)) | 1))
                case -76:
                    delete(v2, 153)
                }
            }
        }
        _ = v42
    } else {
        fmt.Println(v4, v4)
        var v44 int32 = 828
        var v45 int64 = -144
        _ = v44
        _ = v45
    }
    switch "hello" {
    case "a":
        var v46 int = (v6 * 4294967295)
        _ = v46
    case "foo bar":
        v47 := GenS4{F0: (!(46 > 100)), F1: 2147483647, F2: 1}
        _ = v47
    case "\n\t":
        v6 = 42
    }
    if (genFn0(-1, 2, -32768) >= 556) {
        for fi48 := 0; fi48 < 1; fi48++ {
            fmt.Println(fi48)
            var v49 int64 = -932
            var v50 uint32 = 59
            _ = fi48
            _ = v49
            _ = v50
        }
        delete(v2, 2147483648)
    } else {
        var v51 string = ("a" + "\n\t")
        for fi52 := 0; fi52 < 15; fi52++ {
            switch v51 {
            case "\n\t":
                for fi53 := 0; fi53 < 6; fi53++ {
                    var v54 int = (fi53 << ((0) & 63))
                    fmt.Println(v4)
                    v54 += (fi52 << ((860) & 63))
                    _ = fi53
                    _ = v54
                }
            case "":
                fmt.Println(fi52)
            default:
                switch "\n\t" {
                case "foo bar":
                    v51 = ""
                case "\n\t":
                    var v55 float32 = v4
                    _ = v55
                default:
                    v56 := GenS3{F0: 957}
                    _ = v56
                }
            }
            _ = len(v5)
            _ = fi52
        }
        var v57 bool = (false && true)
        _ = v51
        _ = v57
    }
    _ = len(v2)
    var v58 int32 = -2147483648
    v4 += v4
    var v59 float32 = float32((0.1 * float64((-2.25 - v4))))
    for fi60 := 0; fi60 < 13; fi60++ {
        var v61 float64 = -2.25
        _ = fi60
        _ = v61
    }
    switch "foo bar" {
    case "hello":
        var v62 uint16 = 985
        _ = v62
    default:
        switch v6 {
        case -1:
            for fi63 := 0; fi63 < 15; fi63++ {
                var v64 string = ("\n\t" + ("hello" + (("a" + "\n\t") + "\n\t")))
                delete(v2, 2147483648)
                _ = fi63
                _ = v64
            }
        case -50:
            var v65 uint8 = 101
            _ = v65
        }
    }
    for fi66 := 0; fi66 < 15; fi66++ {
        fmt.Println(fi66, v58)
        if true {
            switch (-(v6)) {
            case 0:
                var v67 int32 = v58
                _ = v67
            case 4294967295:
                v4 -= ((v4 / (v4 / ((3.14159 + v59) * v4))) * float32(0.0))
            }
            v68 := GenS7{F0: 32767, F1: 18446744073709551615}
            _ = v68
        } else {
            switch "hello" {
            case "\n\t":
                var v69 int16 = 32767
                _ = v69
            }
            if ((0.1 / (1e10 + 1e10)) > (1e10 * 0.1)) {
                if len(v3) > 0 {
                    v3[0] = (("foo bar" + "\n\t") + ("hello" + ("hello" + "foo bar")))
                }
                v59 = v59
                var v70 uint32 = 2147483648
                _ = v70
            } else {
                switch "\n\t" {
                case "hello":
                    if true {
                        var v71 float32 = v59
                        v6 += 588
                        _ = v71
                    }
                case "a":
                    for fi72 := 0; fi72 < 15; fi72++ {
                        fmt.Println(fi66, v58)
                        if (!(float32(float64(1e10)) != (v59 / v4))) {
                            v6 -= fi66
                            v59 += ((1e10 - float32((0.0 - float64(v59)))) - ((v59 + (float32(-2.25) - v59)) / -2.25))
                            var v73 uint8 = 161
                            _ = v73
                        } else {
                            v58 = 251
                            v74 := GenS6{F0: 18446744073709551615, F1: -120}
                            switch ("a" + "hello") {
                            case "\n\t":
                                v58 += (v58 % (((v58 - (v58 / (((v58 ^ (v58 + (v58 % ((-2147483648) | 1))))) | 1)))) | 1))
                            default:
                                v58 = (v58 * v58)
                            }
                            _ = v74
                        }
                        var v75 int = (fi72 ^ fi72)
                        _ = fi72
                        _ = v75
                    }
                case "":
                    v76 := GenS1{F0: 4294967295, F1: -128}
                    _ = v76
                default:
                    if false {
                        if false {
                            for fi77 := 0; fi77 < 9; fi77++ {
                                fmt.Println(v4, fi66, v4)
                                var v78 float64 = 0.1
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                var v79 float32 = (v59 + 1e10)
                                if true {
                                    v80 := GenS2{F0: -1, F1: 255, F2: -667}
                                    v81 := GenS8{F0: 100}
                                    if false {
                                        var v82 uint8 = 255
                                        switch fi77 {
                                        case 4294967295:
                                            if (v78 > (v78 * (v78 * v78))) {
                                                if false {
                                                    var v83 int = 198
                                                    _ = v83
                                                } else {
                                                    v84 := v2[1]
                                                    _ = v84
                                                }
                                                if true {
                                                    if false {
                                                        switch "" {
                                                        case "hello":
                                                            if len(v3) >= 2 {
                                                                v3 = v3[1:len(v3)]
                                                            }
                                                        case "":
                                                            v79 += (3.14159 + float32((3.14159 + 1.5)))
                                                        }
                                                        v79 = v59
                                                        v1 = append(v1, (fi66 ^ fi77))
                                                    } else {
                                                        v85 := GenS7{F0: -32768, F1: 0}
                                                        var v86 float32 = (v4 * (v79 / v4))
                                                        _ = v85
                                                        _ = v86
                                                    }
                                                    var v87 int16 = -1
                                                    v79 = 3.14159
                                                    _ = v87
                                                } else {
                                                    if false {
                                                        var v88 int8 = -128
                                                        v89 := GenS8{F0: (v88 >> ((668) & 63))}
                                                        if (!true) {
                                                            var v90 uint8 = v82
                                                            _ = v90
                                                        }
                                                        _ = v88
                                                        _ = v89
                                                    }
                                                    v82 = v82
                                                }
                                                var v91 int8 = -101
                                                _ = v91
                                            } else {
                                                v4 += (v4 - 1.5)
                                                v92 := GenS9{F0: 2147483648, F1: v6}
                                                _ = v92
                                            }
                                        }
                                        _ = v82
                                    } else {
                                        if true {
                                            if (!(true && true)) {
                                                var v93 uint8 = 174
                                                v58 += (v58 - -577)
                                                fmt.Println(v78, v93)
                                                _ = v93
                                            }
                                            if true {
                                                v94 := GenS5{F0: ("\n\t" + "foo bar"), F1: (false || true), F2: (!true)}
                                                v95 := GenS10{F0: genFn0(127, 42, -226), F1: -590}
                                                _ = len(v3)
                                                fmt.Println(fi66, v79)
                                                if false {
                                                    v96 := GenS11{F0: (("foo bar" + "foo bar") + "\n\t"), F1: 363, F2: 65535}
                                                    var v97 int16 = -406
                                                    v59 -= v4
                                                    _ = v96
                                                    _ = v97
                                                } else {
                                                    fmt.Println(v58)
                                                }
                                                _ = v94
                                                _ = v95
                                            }
                                            if (!(-1 == -1)) {
                                                if ((65535 != 0) || false) {
                                                    var v98 uint64 = 4294967295
                                                    var v99 float64 = v78
                                                    _ = v98
                                                    _ = v99
                                                } else {
                                                    if ("foo bar" >= "hello") {
                                                        switch ("" + "\n\t") {
                                                        case "":
                                                            fmt.Println(v4, v6)
                                                        default:
                                                            v79 = (0.1 * 1.5)
                                                        }
                                                        v59 = float32(v78)
                                                        v100 := GenS4{F0: (42 > 32767), F1: v58, F2: 128}
                                                        _ = v100
                                                    } else {
                                                        switch "\n\t" {
                                                        case "hello":
                                                            fmt.Println(v78, v78, v79)
                                                        case "":
                                                            switch (v6 >> ((4294967295) & 63)) {
                                                            case 128:
                                                                switch "foo bar" {
                                                                case "hello":
                                                                    v101 := GenS12{F0: "\n\t"}
                                                                    _ = v101
                                                                case "\n\t":
                                                                    v2[2147483648] = fi77
                                                                    v102 := GenS11{F0: ("foo bar" + ("" + ("" + ""))), F1: -713, F2: 42}
                                                                    _ = v102
                                                                }
                                                            default:
                                                                var v103 int32 = (v58 % (((-(v58))) | 1))
                                                                _ = v103
                                                            }
                                                        case "foo bar":
                                                            var v104 uint64 = 4294967295
                                                            _ = v104
                                                        }
                                                    }
                                                    delete(v2, 757)
                                                    v1 = append(v1, (fi66 >> ((255) & 63)))
                                                }
                                            }
                                        }
                                        _ = len(v2)
                                        var v105 uint16 = 42
                                        _ = v105
                                    }
                                    _ = v80
                                    _ = v81
                                }
                                v78 = v78
                                _ = fi77
                                _ = v78
                                _ = v79
                            }
                        }
                    } else {
                        v106 := GenS9{F0: 2, F1: (v6 >> ((2) & 63))}
                        _ = v106
                    }
                }
                if (false && ("" > "\n\t")) {
                    if true {
                        if (1 == -73) {
                            v107 := GenS8{F0: 77}
                            var v108 uint8 = 119
                            _ = v107
                            _ = v108
                        } else {
                            if (2 < 2147483647) {
                                v109 := GenS13{F0: 1}
                                _ = v109
                            } else {
                                for fi110 := 0; fi110 < 4; fi110++ {
                                    v111 := GenS9{F0: 227, F1: 32768}
                                    _ = fi110
                                    _ = v111
                                }
                                fmt.Println(v58, v6, v58)
                                switch 256 {
                                case 1:
                                    for fi112 := 0; fi112 < 1; fi112++ {
                                        v113 := GenS2{F0: 1000, F1: 125, F2: 0}
                                        _ = fi112
                                        _ = v113
                                    }
                                case -9223372036854775808:
                                    v4 = 1.5
                                default:
                                    var v114 int = fi66
                                    _ = v114
                                }
                            }
                            fmt.Println(v4)
                        }
                    } else {
                        switch (fi66 >> ((18446744073709551615) & 63)) {
                        case 128:
                            v2[826] = fi66
                        default:
                            if (!(32767 < 0)) {
                                var v115 int16 = 0
                                _ = v115
                            } else {
                                v3 = append(v3, "")
                                var v116 int8 = 0
                                _ = v116
                            }
                        }
                        delete(v2, 2)
                        for fi117 := 0; fi117 < 19; fi117++ {
                            fmt.Println(v4, v4)
                            _ = fi117
                        }
                    }
                    fmt.Println(v58)
                } else {
                    switch "" {
                    case "foo bar":
                        fmt.Println(v4, v6, v4)
                    case "hello":
                        for fi118 := 0; fi118 < 12; fi118++ {
                            if len(v3) > 0 {
                                v3[0] = ""
                            }
                            _ = fi118
                        }
                    default:
                        v119 := GenS11{F0: ("" + (("foo bar" + "a") + "\n\t")), F1: 1, F2: 1}
                        _ = v119
                    }
                }
            }
        }
        v59 -= float32((3.14159 * (0.1 - 3.14159)))
        _ = fi66
    }
    v58 -= (v58 / ((v58) | 1))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v58:", v58)
    fmt.Println("v59:", v59)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v58
    _ = v59
    _ = v6
    _ = v9
}
