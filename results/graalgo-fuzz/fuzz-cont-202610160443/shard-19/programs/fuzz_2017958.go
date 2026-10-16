package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64 }
type GenS2 struct { F0 int8; F1 int32; F2 int32 }
type GenS3 struct { F0 uint; F1 uint; F2 int64 }
type GenS4 struct { F0 string; F1 uint16 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 int16; F1 int; F2 uint32 }
type GenS8 struct { F0 uint16; F1 uint8 }
type GenS9 struct { F0 int8; F1 int32 }
type GenS10 struct { F0 uint8; F1 int; F2 int8 }
type GenS11 struct { F0 int8; F1 uint64; F2 uint8 }
type GenS12 struct { F0 string }
type GenS13 struct { F0 uint16 }
func genFn0(p0 bool, p1 uint64, p2 bool) string {
    v7 := GenS1{F0: 0, F1: 42}
    if ((0.1 / 1.5) > 0.0) {
        v8 := GenS1{F0: 100, F1: 0}
        fmt.Println(p0)
        fmt.Println(p1)
        _ = v8
    } else {
        if p0 {
            p0 = true
            if p0 {
                p2 = (p2 || (255 == 32767))
                v9 := GenS1{F0: 127, F1: -9223372036854775808}
                _ = v9
            } else {
                v10 := []int{}
                _ = v10
            }
        } else {
            if ((p2 || (!true)) || true) {
                var v11 bool = (42 == 100)
                p2 = (!(855 > p1))
                var v12 uint16 = 1
                _ = v11
                _ = v12
            }
            v13 := []int{}
            _ = v13
        }
    }
    switch "a" {
    case "foo bar":
        var v14 float32 = 1e10
        _ = v14
    case "hello":
        var v15 float32 = ((-2.25 - 3.14159) + -2.25)
        _ = v15
    }
    fmt.Println(p0, p1, p1)
    for fi16 := 0; fi16 < 4; fi16++ {
        v17 := []int{fi16, -1}
        fmt.Println(p0)
        v18 := GenS1{F0: 42, F1: -153}
        _ = fi16
        _ = v17
        _ = v18
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return ("" + "\n\t")
}

func main() {
    var v1 uint8 = 128
    var v2 uint8 = (v1 << ((285) & 63))
    var v3 uint64 = uint64(v1)
    var v4 uint32 = 42
    v5 := []int{-2147483648}
    v6 := make(map[int]int)
    var v19 float32 = 1e10
    _ = len(v5)
    var v20 bool = (("a" < (("" + ((("\n\t" + "\n\t") + (((("" + "foo bar") + (("hello" + ("\n\t" + "foo bar")) + ("a" + "foo bar"))) + "") + (("foo bar" + "a") + ("\n\t" + ("a" + "\n\t"))))) + "hello")) + "a")) && true)
    if (127 >= 100) {
        for fi21 := 0; fi21 < 1; fi21++ {
            delete(v6, -512)
            if ((v4 == 42) && v20) {
                v22 := GenS1{F0: 0, F1: 992}
                var v23 int64 = 9223372036854775807
                _ = v22
                _ = v23
            } else {
                if v20 {
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    var v24 uint64 = (v3 / ((0) | 1))
                    _ = v24
                } else {
                    if (0 >= 4294967295) {
                        for fi25 := 0; fi25 < 19; fi25++ {
                            if (v20 && (((1e10 * -2.25) > -2.25) && v20)) {
                                var v26 float64 = 0.1
                                v27 := v6[4294967295]
                                if (int32(v2) > 2147483647) {
                                    v28 := GenS2{F0: -118, F1: 42, F2: 2147483647}
                                    fmt.Println(fi25)
                                    _ = v28
                                }
                                if false {
                                    v5 = append(v5, (-(v27)))
                                    if (-2.25 != (v19 + (v19 / 3.14159))) {
                                        switch (v27 / (((v27 ^ int(v2))) | 1)) {
                                        case 206:
                                            fmt.Println(v20, v19, v1)
                                        case 0:
                                            if len(v5) >= 2 {
                                                v5 = v5[1:len(v5)]
                                            }
                                        case -990:
                                            var v29 int64 = 42
                                            fmt.Println(v2)
                                            var v30 float32 = 0.1
                                            _ = v29
                                            _ = v30
                                        }
                                        var v31 float64 = v26
                                        fmt.Println(v27, v4)
                                        _ = v31
                                    } else {
                                        if false {
                                            switch (("" + "hello") + "") {
                                            case "a":
                                                switch "foo bar" {
                                                case "":
                                                    v20 = false
                                                case "hello":
                                                    fmt.Println(v19, v26, v27)
                                                case "a":
                                                    var v32 uint = 2
                                                    v1 = uint8(v3)
                                                    _ = v32
                                                }
                                            }
                                        } else {
                                            if len(v5) > 0 {
                                                v5[0] = 32767
                                            }
                                            v26 = 0.1
                                            v33 := GenS1{F0: 0, F1: 42}
                                            _ = v33
                                        }
                                    }
                                    if true {
                                        var v34 uint16 = 0
                                        var v35 bool = v20
                                        _ = v34
                                        _ = v35
                                    } else {
                                        v19 = v19
                                        v27 = -923
                                        if true {
                                            v36 := GenS1{F0: -30, F1: 42}
                                            v37 := GenS2{F0: 81, F1: -2147483648, F2: 42}
                                            if (true && (!(32767 >= -1))) {
                                                var v38 string = ""
                                                var v39 uint16 = 42
                                                var v40 int32 = -2147483648
                                                _ = v38
                                                _ = v39
                                                _ = v40
                                            } else {
                                                var v41 bool = ((v3 >> ((255) & 63)) >= (v3 | 2))
                                                _ = v41
                                            }
                                            _ = v36
                                            _ = v37
                                        } else {
                                            if (v20 && (-1 != -2147483648)) {
                                                _ = len(v6)
                                                if (255 != 0) {
                                                    var v42 uint16 = 796
                                                    v19 = ((v19 * (1.5 + 1e10)) + float32(float64(v19)))
                                                    _ = v42
                                                } else {
                                                    v20 = (!v20)
                                                }
                                            }
                                        }
                                    }
                                } else {
                                    fmt.Println(v4)
                                    var v43 int = fi21
                                    var v44 int32 = -634
                                    _ = v43
                                    _ = v44
                                }
                                fmt.Println(v4, v4)
                                _ = v26
                                _ = v27
                            } else {
                                var v45 int16 = 42
                                var v46 uint32 = 2147483647
                                fmt.Println(v19, v4, v2)
                                _ = v45
                                _ = v46
                            }
                            _ = fi25
                        }
                    } else {
                        v47 := GenS2{F0: 1, F1: 2147483647, F2: 42}
                        v2 = (v1 & v1)
                        _ = v47
                    }
                    v4 += 2147483648
                    v4 = v4
                }
                var v48 int8 = 51
                if (!(0 > 32767)) {
                    var v49 int16 = 42
                    v1 -= v1
                    v1 -= v1
                    _ = v49
                }
                _ = v48
            }
            v1 += (v2 - v1)
            _ = fi21
        }
        if (1e10 >= float64(((v19 / float32((1.5 * float64(1e10)))) * v19))) {
            if len(v5) > 0 {
                v5[0] = 453
            }
            fmt.Println(v1, v20)
            for fi50 := 0; fi50 < 1; fi50++ {
                v51 := GenS3{F0: 65535, F1: 700, F2: -1}
                for fi52 := 0; fi52 < 4; fi52++ {
                    if ((1.5 + (3.14159 + 1e10)) < float64(-2.25)) {
                        if (42 < 908) {
                            v2 -= (v2 - (v2 & (v2 + (v2 + 2))))
                        }
                        switch -9223372036854775808 {
                        case -9223372036854775808:
                            var v53 int64 = -9223372036854775808
                            _ = v53
                        case -674:
                            var v54 uint32 = (v4 % ((v4) | 1))
                            _ = v54
                        }
                        _ = len(v5)
                    }
                    _ = fi52
                }
                for fi55 := 0; fi55 < 16; fi55++ {
                    if ("hello" > "hello") {
                        v19 = 0.1
                        v56 := GenS1{F0: -1, F1: 860}
                        if ((false || ((^(v3)) >= (^(v3)))) || ("foo bar" <= "hello")) {
                            v4 = v4
                            v57 := GenS3{F0: 255, F1: 34, F2: 508}
                            v5 = append(v5, 296)
                            _ = v57
                        } else {
                            var v58 int32 = -266
                            _ = v58
                        }
                        _ = v56
                    }
                    _ = fi55
                }
                _ = fi50
                _ = v51
            }
        } else {
            fmt.Println(v1, v3, v2)
            if (18446744073709551615 > 0) {
                v3 += v3
            } else {
                if v20 {
                    if (-9223372036854775808 <= 0) {
                        fmt.Println(v1, v20)
                        for fi59 := 0; fi59 < 5; fi59++ {
                            switch (^(fi59)) {
                            case -1:
                                var v60 uint16 = 1
                                _ = v60
                            }
                            for fi61 := 0; fi61 < 8; fi61++ {
                                v6[0] = (fi59 | fi61)
                                v4 += (v4 | (v4 ^ uint32(v3)))
                                var v62 int16 = 1000
                                switch fi59 {
                                case 332:
                                    switch "hello" {
                                    case "\n\t":
                                        var v63 uint = 1
                                        _ = v63
                                    default:
                                        var v64 uint = 856
                                        _ = v64
                                    }
                                    v6[fi59] = (fi59 + fi59)
                                }
                                v65 := v6[fi61]
                                switch (-(fi59)) {
                                case 4294967295:
                                    var v66 uint32 = uint32(v3)
                                    _ = v66
                                default:
                                    v20 = (v19 != (1.5 + float32(0.1)))
                                }
                                if v20 {
                                    switch ("" + "foo bar") {
                                    case "a":
                                        if true {
                                            v1 = 185
                                            var v67 uint8 = 0
                                            _ = v67
                                        }
                                    case "":
                                        if ((false || ((false && v20) || (774 == 1))) || (v19 > v19)) {
                                            switch "" {
                                            case "":
                                                v68 := GenS4{F0: "hello", F1: 800}
                                                _ = v68
                                            case "foo bar":
                                                if v20 {
                                                    _ = len(v6)
                                                    delete(v6, (v65 << ((42) & 63)))
                                                } else {
                                                    if v20 {
                                                        var v69 float32 = 3.14159
                                                        _ = v69
                                                    }
                                                    switch ("a" + "hello") {
                                                    case "a":
                                                        switch v65 {
                                                        case 0:
                                                            fmt.Println(v19, fi59)
                                                        default:
                                                            v2 += v1
                                                        }
                                                    case "":
                                                        fmt.Println(v19, v1)
                                                    }
                                                    v2 -= 161
                                                    switch int(v4) {
                                                    case 9223372036854775807:
                                                        fmt.Println(fi59, v2)
                                                    case -1:
                                                        v20 = ((v20 || v20) || false)
                                                    case 32767:
                                                        if v20 {
                                                            _ = len(v6)
                                                        } else {
                                                            v65 = 267
                                                        }
                                                    }
                                                    fmt.Println(v3, v2, v62)
                                                    if ((v62 + v62) > v62) {
                                                        v70 := GenS5{F0: "hello"}
                                                        v20 = false
                                                        _ = v70
                                                    } else {
                                                        var v71 string = "hello"
                                                        _ = v71
                                                    }
                                                    v19 = (v19 - v19)
                                                }
                                            }
                                        } else {
                                            var v72 int32 = -1
                                            var v73 float32 = ((0.1 + 0.1) + float32(-2.25))
                                            fmt.Println(v62, v19)
                                            _ = v72
                                            _ = v73
                                        }
                                    }
                                    var v74 uint64 = v3
                                    if (false || false) {
                                        v75 := GenS6{F0: -897}
                                        _ = v75
                                    }
                                    var v76 bool = false
                                    _ = v74
                                    _ = v76
                                } else {
                                    v3 += 42
                                    v77 := GenS4{F0: "\n\t", F1: 42}
                                    v78 := GenS4{F0: "a", F1: 4}
                                    _ = v77
                                    _ = v78
                                }
                                if (v4 != 123) {
                                    var v79 int32 = 0
                                    _ = v79
                                } else {
                                    if true {
                                        _ = len(v6)
                                        switch v65 {
                                        case 42:
                                            var v80 int32 = 807
                                            _ = v80
                                        case -2147483648:
                                            fmt.Println(v2)
                                        default:
                                            var v81 bool = v20
                                            _ = v81
                                        }
                                        switch "hello" {
                                        case "foo bar":
                                            switch "foo bar" {
                                            case "a":
                                                if (0 < v62) {
                                                    v82 := GenS7{F0: v62, F1: (v65 - (-(v65))), F2: (v4 << ((607) & 63))}
                                                    if false {
                                                        if (v62 >= v62) {
                                                            v83 := v6[32768]
                                                            if true {
                                                                switch ("a" + ((("a" + "") + ("foo bar" + ("" + ""))) + "\n\t")) {
                                                                case "hello":
                                                                    var v84 uint = 42
                                                                    _ = v84
                                                                case "":
                                                                    v6[fi59] = (fi59 | (v83 - v65))
                                                                }
                                                            } else {
                                                                var v85 float64 = (3.14159 / (1e10 - 0.0))
                                                                var v86 int16 = (^(v62))
                                                                v87 := GenS8{F0: 744, F1: 1}
                                                                var v88 int16 = -1
                                                                _ = v85
                                                                _ = v86
                                                                _ = v87
                                                                _ = v88
                                                            }
                                                            _ = v83
                                                        } else {
                                                            v89 := GenS2{F0: 0, F1: 2147483647, F2: int32(v62)}
                                                            var v90 uint64 = uint64(v4)
                                                            var v91 float64 = (float64(0.0) + 1e10)
                                                            _ = v89
                                                            _ = v90
                                                            _ = v91
                                                        }
                                                        if false {
                                                            v4 -= 42
                                                        }
                                                        v92 := GenS9{F0: 42, F1: 1}
                                                        _ = v92
                                                    }
                                                    _ = v82
                                                } else {
                                                    v93 := GenS10{F0: v2, F1: (fi59 ^ -9223372036854775808), F2: 3}
                                                    _ = v93
                                                }
                                            case "\n\t":
                                                var v94 int64 = 42
                                                _ = v94
                                            default:
                                                fmt.Println(v20, fi59, fi61)
                                            }
                                        case "a":
                                            var v95 uint8 = (v1 * v2)
                                            _ = v95
                                        case "":
                                            if v20 {
                                                var v96 int32 = 2147483647
                                                switch ("\n\t" + ("\n\t" + "")) {
                                                case "\n\t":
                                                    var v97 bool = false
                                                    _ = v97
                                                case "hello":
                                                    fmt.Println(v19)
                                                }
                                                _ = v96
                                            } else {
                                                v98 := GenS8{F0: 0, F1: (v1 / ((v2) | 1))}
                                                var v99 int16 = -32768
                                                if v20 {
                                                    var v100 float64 = -2.25
                                                    _ = v100
                                                } else {
                                                    v101 := GenS2{F0: 127, F1: 275, F2: 942}
                                                    v62 += (^(v62))
                                                    if true {
                                                        fmt.Println(v19)
                                                        _ = len(v6)
                                                    } else {
                                                        v2 -= 72
                                                        _ = len(v6)
                                                        fmt.Println(v62, v2)
                                                    }
                                                    var v102 int16 = 42
                                                    if (-1 == -114) {
                                                        v102 = 0
                                                    }
                                                    _ = v101
                                                    _ = v102
                                                }
                                                _ = v98
                                                _ = v99
                                            }
                                        default:
                                            v103 := GenS8{F0: uint16(v4), F1: (v2 >> ((2) & 63))}
                                            _ = v103
                                        }
                                        var v104 int64 = 963
                                        _ = v104
                                    } else {
                                        var v105 int64 = 42
                                        _ = v105
                                    }
                                }
                                _ = fi61
                                _ = v62
                                _ = v65
                            }
                            if (!(!((-9223372036854775808 != 567) || true))) {
                                for fi106 := 0; fi106 < 5; fi106++ {
                                    switch 127 {
                                    case -9223372036854775808:
                                        var v107 uint8 = (^(v1))
                                        _ = v107
                                    case -1:
                                        fmt.Println(fi59)
                                    }
                                    v108 := v6[fi106]
                                    _ = fi106
                                    _ = v108
                                }
                            } else {
                                if v20 {
                                    v20 = v20
                                    var v109 bool = v20
                                    if v109 {
                                        v110 := GenS5{F0: ""}
                                        _ = v110
                                    }
                                    _ = v109
                                }
                                var v111 float64 = 0.0
                                for fi112 := 0; fi112 < 11; fi112++ {
                                    var v113 int64 = 1
                                    var v114 string = "a"
                                    var v115 uint8 = (v2 % ((200) | 1))
                                    fmt.Println(v111, v111, v2)
                                    _ = fi112
                                    _ = v113
                                    _ = v114
                                    _ = v115
                                }
                                _ = v111
                            }
                            _ = fi59
                        }
                        var v116 string = ("" + "hello")
                        _ = v116
                    } else {
                        _ = len(v6)
                        v3 = (v3 | (v3 ^ (^(v3))))
                        var v117 int = -574
                        _ = v117
                    }
                } else {
                    var v118 string = "\n\t"
                    _ = v118
                }
                for fi119 := 0; fi119 < 16; fi119++ {
                    for fi120 := 0; fi120 < 19; fi120++ {
                        var v121 uint = 2
                        var v122 int16 = -32768
                        _ = fi120
                        _ = v121
                        _ = v122
                    }
                    delete(v6, fi119)
                    for fi123 := 0; fi123 < 1; fi123++ {
                        v124 := GenS11{F0: 100, F1: (v3 >> ((379) & 63)), F2: 124}
                        fmt.Println(v2, v3, fi119)
                        _ = fi123
                        _ = v124
                    }
                    _ = fi119
                }
            }
        }
    } else {
        fmt.Println(v19, v19)
    }
    v2 -= v1
    fmt.Println(v4)
    var v125 bool = (!(468 >= -2147483648))
    var v126 int64 = -904
    for fi127 := 0; fi127 < 18; fi127++ {
        switch ("foo bar" + "a") {
        case "hello":
            v128 := GenS12{F0: "foo bar"}
            _ = v128
        case "a":
            var v129 int8 = 38
            _ = v129
        default:
            if (!v125) {
                v19 = ((v19 / v19) - v19)
                v3 = v3
            } else {
                delete(v6, (fi127 | 664))
            }
        }
        for fi130 := 0; fi130 < 11; fi130++ {
            v131 := GenS7{F0: -32768, F1: (fi127 + 32768), F2: (v4 << ((743) & 63))}
            _ = fi130
            _ = v131
        }
        for fi132 := 0; fi132 < 3; fi132++ {
            var v133 float64 = float64(v19)
            v134 := GenS13{F0: 0}
            _ = fi132
            _ = v133
            _ = v134
        }
        _ = fi127
    }
    fmt.Println("v1:", v1)
    fmt.Println("v125:", v125)
    fmt.Println("v126:", v126)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v125
    _ = v126
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
