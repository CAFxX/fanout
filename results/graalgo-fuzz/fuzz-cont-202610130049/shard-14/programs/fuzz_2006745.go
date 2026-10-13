package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32 }
type GenS2 struct { F0 uint; F1 uint; F2 bool }
type GenS3 struct { F0 uint8; F1 uint32; F2 uint8 }
type GenS4 struct { F0 int8; F1 uint16 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 int8; F1 uint16; F2 uint32 }
type GenS7 struct { F0 int64; F1 int8; F2 uint64 }
type GenS8 struct { F0 int; F1 uint8; F2 uint8 }
type GenS9 struct { F0 uint16; F1 bool; F2 int }
type GenS10 struct { F0 int32; F1 int32; F2 int32 }
type GenS11 struct { F0 uint32; F1 string; F2 int32 }
type GenS12 struct { F0 int32; F1 int8; F2 bool }
type GenS13 struct { F0 uint8; F1 bool }
type GenS14 struct { F0 int8; F1 int64; F2 uint }
type GenS15 struct { F0 uint8; F1 bool }
func genFn0(p0 uint32, p1 uint8, p2 uint16) uint32 {
    var v5 string = "hello"
    v6 := []int{}
    p0 += (p0 ^ (p0 / ((p0) | 1)))
    for fi7 := 0; fi7 < 17; fi7++ {
        var v8 int8 = 0
        _ = fi7
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return (p0 >> ((4294967295) & 63))
}
func genFn1() bool {
    var v9 float64 = (float64((0.1 * 0.1)) * 1e10)
    fmt.Println(v9)
    _ = v9
    return true
}

func main() {
    var v1 int16 = 0
    v2 := make(map[uint32]bool)
    v3 := make(map[string]string)
    v4 := []string{"\n\t", (("a" + "") + "hello")}
    if (!(((-2.25 + 1.5) + float64(1e10)) > 0.1)) {
        for fi10 := 0; fi10 < 18; fi10++ {
            switch 1 {
            case -808:
                var v11 bool = (!((2147483647 <= 2147483647) && ((false || false) || false)))
                _ = v11
            case 797:
                fmt.Println(fi10, v1)
            default:
                var v12 int = fi10
                _ = v12
            }
            switch (^(fi10)) {
            case 32768:
                var v13 uint = 296
                _ = v13
            case -742:
                _ = len(v4)
            }
            _ = fi10
        }
        var v14 string = "\n\t"
        _ = v14
    } else {
        if (1 <= 198) {
            if (!(797 > 557)) {
                var v15 float64 = 3.14159
                v15 = -2.25
                _ = v15
            }
            v16 := GenS1{F0: 100, F1: -342}
            v1 = (v1 % (((v1 % (((v1 * (v1 << ((65535) & 63)))) | 1))) | 1))
            _ = v16
        } else {
            v17 := GenS1{F0: 68, F1: 0}
            if (42 >= -82) {
                var v18 int32 = 2147483647
                v19 := GenS2{F0: 4294967295, F1: 1, F2: true}
                v20 := GenS3{F0: 1, F1: 4294967295, F2: 1}
                _ = v18
                _ = v19
                _ = v20
            } else {
                v21 := GenS3{F0: 2, F1: 42, F2: 0}
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                if (int(v1) <= 2147483647) {
                    for fi22 := 0; fi22 < 6; fi22++ {
                        v23 := GenS2{F0: 741, F1: 4294967295, F2: false}
                        if true {
                            if ((fi22 & fi22) != fi22) {
                                var v24 int16 = (v1 - v1)
                                var v25 float32 = 3.14159
                                var v26 int = fi22
                                _ = v24
                                _ = v25
                                _ = v26
                            } else {
                                v27 := GenS4{F0: 0, F1: 2}
                                v3[""] = ""
                                delete(v3, "foo bar")
                                _ = v27
                            }
                            var v28 string = "hello"
                            _ = v28
                        } else {
                            var v29 uint64 = 1
                            v30 := GenS2{F0: 0, F1: 46, F2: (!(979 < 1))}
                            _ = v29
                            _ = v30
                        }
                        _ = fi22
                        _ = v23
                    }
                    switch 9223372036854775807 {
                    case -2147483648:
                        delete(v2, 2147483647)
                    case 676:
                        v1 = 32767
                    }
                    v1 -= (v1 / ((v1) | 1))
                } else {
                    v31 := GenS3{F0: 127, F1: 185, F2: 2}
                    switch 51 {
                    case 127:
                        var v32 uint64 = 0
                        _ = v32
                    case 128:
                        var v33 float32 = 1.5
                        _ = v33
                    case -369:
                        var v34 int16 = v1
                        _ = v34
                    }
                    v1 = (v1 % ((v1) | 1))
                    _ = v31
                }
                _ = v21
            }
            for fi35 := 0; fi35 < 13; fi35++ {
                var v36 float64 = 3.14159
                if false {
                    var v37 bool = false
                    switch "\n\t" {
                    case "hello":
                        for fi38 := 0; fi38 < 2; fi38++ {
                            v39 := v2[42]
                            if v39 {
                                var v40 int8 = 42
                                switch (fi35 & 127) {
                                case -9223372036854775808:
                                    var v41 bool = (v37 && (!(v37 || (4294967295 < 4294967295))))
                                    _ = v41
                                case 508:
                                    v42 := GenS4{F0: (v40 & (v40 << ((0) & 63))), F1: 32767}
                                    _ = v42
                                }
                                var v43 uint16 = 32767
                                v44 := v2[42]
                                _ = v40
                                _ = v43
                                _ = v44
                            }
                            var v45 int64 = 9223372036854775807
                            _ = fi38
                            _ = v39
                            _ = v45
                        }
                    case "a":
                        switch ("a" + ("\n\t" + "\n\t")) {
                        case "\n\t":
                            v46 := GenS5{F0: -863}
                            _ = v46
                        case "foo bar":
                            if v37 {
                                switch "" {
                                case "\n\t":
                                    v1 += 190
                                }
                            } else {
                                switch fi35 {
                                case 1:
                                    var v47 int64 = -1
                                    _ = v47
                                case -2147483648:
                                    for fi48 := 0; fi48 < 14; fi48++ {
                                        var v49 uint32 = 4294967295
                                        v50 := GenS6{F0: -102, F1: 765, F2: 2}
                                        v36 = ((float64(1e10) + 1.5) - v36)
                                        _ = fi48
                                        _ = v49
                                        _ = v50
                                    }
                                case 865:
                                    v51 := GenS7{F0: 0, F1: -128, F2: 42}
                                    _ = v51
                                default:
                                    switch "\n\t" {
                                    case "a":
                                        var v52 int32 = -645
                                        _ = v52
                                    default:
                                        if (971 != 1) {
                                            v37 = v37
                                        } else {
                                            v53 := GenS6{F0: 1, F1: 42, F2: genFn0(852, 127, 2)}
                                            var v54 int = 32767
                                            v55 := GenS3{F0: 157, F1: 2, F2: 128}
                                            _ = v53
                                            _ = v54
                                            _ = v55
                                        }
                                    }
                                }
                                v56 := GenS8{F0: fi35, F1: 0, F2: 42}
                                _ = v56
                            }
                        case "a":
                            v57 := GenS9{F0: 236, F1: v37, F2: (fi35 ^ (fi35 | fi35))}
                            _ = v57
                        }
                    case "foo bar":
                        var v58 int64 = -858
                        _ = v58
                    default:
                        var v59 int32 = 42
                        _ = v59
                    }
                    if false {
                        switch fi35 {
                        case 9223372036854775807:
                            var v60 uint8 = 128
                            _ = v60
                        case 2147483647:
                            if v37 {
                                v61 := GenS9{F0: 431, F1: true, F2: (fi35 / ((1) | 1))}
                                _ = v61
                            } else {
                                if len(v4) > 0 {
                                    v4[0] = "a"
                                }
                            }
                        }
                        if len(v4) > 0 {
                            v4[0] = "hello"
                        }
                    } else {
                        fmt.Println(v36, v36, v1)
                        for fi62 := 0; fi62 < 17; fi62++ {
                            v37 = ((((float32(0.0) * (1.5 + 1.5)) / 1e10) / 1.5) < 1.5)
                            if true {
                                _ = len(v4)
                                v63 := GenS10{F0: 0, F1: 2147483647, F2: 0}
                                switch (-(fi35)) {
                                case -1:
                                    v64 := GenS7{F0: -412, F1: -100, F2: 0}
                                    _ = v64
                                case -2147483648:
                                    if false {
                                        v37 = v37
                                        v1 = v1
                                        fmt.Println(v1)
                                        v1 -= (v1 | v1)
                                        v65 := GenS11{F0: 1, F1: "", F2: -1}
                                        _ = v65
                                    }
                                case 4294967295:
                                    if v37 {
                                        var v66 bool = v37
                                        switch "\n\t" {
                                        case "":
                                            v1 -= (v1 / (((v1 * (v1 * (v1 & v1)))) | 1))
                                        default:
                                            v37 = v66
                                        }
                                        var v67 int16 = v1
                                        var v68 int8 = -128
                                        _ = v66
                                        _ = v67
                                        _ = v68
                                    }
                                }
                                _ = v63
                            } else {
                                v69 := GenS12{F0: -2147483648, F1: 127, F2: true}
                                v70 := GenS11{F0: 536, F1: ("foo bar" + ("" + "\n\t")), F2: 2147483647}
                                switch "a" {
                                case "foo bar":
                                    fmt.Println(v36, v37, fi35)
                                case "a":
                                    if false {
                                        switch "" {
                                        case "":
                                            switch ("a" + ("foo bar" + "\n\t")) {
                                            case "":
                                                v36 = ((0.0 + ((float64((1e10 - 1e10)) / (v36 - 1.5)) * float64((1e10 + float32(float64((float32(v36) / 1.5))))))) / 1.0)
                                            case "hello":
                                                if false {
                                                    switch (fi62 ^ (^(fi35))) {
                                                    case 4294967295:
                                                        if v37 {
                                                            fmt.Println(v1, v36, v1)
                                                            if ("hello" <= "") {
                                                                v71 := v2[65535]
                                                                v37 = (-75 < -1)
                                                                _ = v71
                                                            }
                                                        } else {
                                                            var v72 float64 = (v36 + v36)
                                                            v73 := GenS10{F0: -2147483648, F1: 42, F2: -2147483648}
                                                            v1 += v1
                                                            _ = v72
                                                            _ = v73
                                                        }
                                                    case -2147483648:
                                                        fmt.Println(v37)
                                                    case 885:
                                                        var v74 uint32 = 787
                                                        _ = v74
                                                    }
                                                }
                                            }
                                        case "foo bar":
                                            if v37 {
                                                fmt.Println(fi35, v37, v36)
                                            } else {
                                                v75 := GenS4{F0: -113, F1: 956}
                                                _ = v75
                                            }
                                        case "a":
                                            switch 9223372036854775807 {
                                            case 117:
                                                v76 := GenS10{F0: 1, F1: 0, F2: -2147483648}
                                                _ = v76
                                            case -2147483648:
                                                v36 = -2.25
                                            }
                                        default:
                                            fmt.Println(v1, v37, v37)
                                        }
                                        fmt.Println(v1, v37)
                                    }
                                default:
                                    var v77 int64 = 42
                                    _ = v77
                                }
                                _ = v69
                                _ = v70
                            }
                            _ = fi62
                        }
                    }
                    _ = v37
                } else {
                    v36 = float64(0.1)
                    for fi78 := 0; fi78 < 7; fi78++ {
                        var v79 bool = true
                        v1 -= (v1 ^ 0)
                        switch ("hello" + "a") {
                        case "\n\t":
                            switch "hello" {
                            case "\n\t":
                                if (!v79) {
                                    if v79 {
                                        if false {
                                            v36 += float64(((-2.25 * (float32(float64(0.1)) - (((1e10 / ((0.1 - -2.25) + 0.0)) + 1.5) + 0.1))) * (1.5 / (float32((v36 + ((v36 - (float64((((-2.25 / 1.5) - 1.5) / (1e10 / 1.5))) * v36)) * (v36 - v36)))) - 1e10))))
                                            _ = len(v4)
                                        } else {
                                            switch 4294967295 {
                                            case 333:
                                                if len(v4) >= 2 {
                                                    v4 = v4[1:len(v4)]
                                                }
                                            case 32767:
                                                if v79 {
                                                    delete(v2, genFn0(genFn0(1, 127, 968), 185, 0))
                                                    var v80 int = (fi78 - (^(fi78)))
                                                    _ = v80
                                                } else {
                                                    if (v79 && (127 > 1)) {
                                                        delete(v3, "foo bar")
                                                        v81 := GenS1{F0: 1, F1: 1}
                                                        if false {
                                                            var v82 float64 = v36
                                                            var v83 float64 = (v82 / v36)
                                                            var v84 int64 = 1
                                                            _ = v82
                                                            _ = v83
                                                            _ = v84
                                                        } else {
                                                            if len(v4) >= 2 {
                                                                v4 = v4[1:len(v4)]
                                                            }
                                                        }
                                                        if v79 {
                                                            if false {
                                                                v36 -= v36
                                                                if false {
                                                                    v85 := v3["a"]
                                                                    if v79 {
                                                                        fmt.Println(v36)
                                                                        fmt.Println(v1, v79, v85)
                                                                        v79 = (2 < 128)
                                                                        delete(v3, (v85 + "hello"))
                                                                        v86 := GenS13{F0: 128, F1: v79}
                                                                        _ = v86
                                                                    }
                                                                    fmt.Println(v85, v79, fi35)
                                                                    _ = v85
                                                                }
                                                            } else {
                                                                var v87 int32 = 764
                                                                if len(v4) > 0 {
                                                                    v4[0] = "foo bar"
                                                                }
                                                                fmt.Println(v87, v1, fi35)
                                                                _ = v87
                                                            }
                                                            var v88 string = "hello"
                                                            _ = v88
                                                        } else {
                                                            if v79 {
                                                                var v89 float64 = 0.0
                                                                if ((!v79) && (fi35 != fi35)) {
                                                                    v90 := GenS14{F0: -56, F1: -9223372036854775808, F2: 4294967295}
                                                                    v79 = (1 == 209)
                                                                    v91 := GenS12{F0: -2147483648, F1: -1, F2: (!v79)}
                                                                    _ = v90
                                                                    _ = v91
                                                                }
                                                                v36 += v36
                                                                _ = v89
                                                            }
                                                            v92 := GenS12{F0: -1, F1: -1, F2: v79}
                                                            if (0 >= 2147483648) {
                                                                var v93 int16 = v1
                                                                v93 = 936
                                                                v94 := GenS11{F0: genFn0(2147483648, 34, 32767), F1: "a", F2: -1}
                                                                _ = v93
                                                                _ = v94
                                                            }
                                                            _ = v92
                                                        }
                                                        if (296 != 667) {
                                                            var v95 int8 = 46
                                                            var v96 uint32 = 164
                                                            _ = v95
                                                            _ = v96
                                                        } else {
                                                            var v97 uint16 = 32767
                                                            v4 = append(v4, "foo bar")
                                                            _ = v97
                                                        }
                                                        _ = v81
                                                    }
                                                }
                                            default:
                                                _ = len(v3)
                                            }
                                            fmt.Println(fi78, v36)
                                        }
                                    }
                                    v79 = (65535 == 32767)
                                    var v98 float64 = v36
                                    _ = v98
                                } else {
                                    var v99 uint16 = 599
                                    _ = v99
                                }
                            default:
                                delete(v3, "\n\t")
                            }
                            v36 -= (0.0 + 0.0)
                        case "a":
                            v79 = ((fi35 << ((18446744073709551615) & 63)) <= fi78)
                        case "":
                            v1 = -872
                        default:
                            fmt.Println(fi35, fi78, fi78)
                        }
                        _ = fi78
                        _ = v79
                    }
                    v1 = (v1 ^ v1)
                }
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                _ = fi35
                _ = v36
            }
            _ = v17
        }
    }
    switch "a" {
    case "foo bar":
        var v100 uint64 = 18446744073709551615
        _ = v100
    default:
        v101 := GenS3{F0: 161, F1: 4294967295, F2: 182}
        _ = v101
    }
    if true {
        v102 := GenS4{F0: 100, F1: 0}
        fmt.Println(v1)
        _ = v102
    } else {
        var v103 uint64 = 976
        _ = v103
    }
    for fi104 := 0; fi104 < 11; fi104++ {
        switch fi104 {
        case 0:
            v1 = (-(v1))
        case 42:
            if true {
                for fi105 := 0; fi105 < 16; fi105++ {
                    v106 := GenS8{F0: 128, F1: 42, F2: 2}
                    fmt.Println(fi104, fi104)
                    if true {
                        fmt.Println(fi105)
                        var v107 string = "\n\t"
                        var v108 int8 = int8(v1)
                        var v109 int8 = (v108 | (v108 * (-(v108))))
                        switch fi104 {
                        case 32768:
                            var v110 uint = 18446744073709551615
                            _ = v110
                        case 127:
                            var v111 uint = 18446744073709551615
                            _ = v111
                        case 128:
                            v112 := GenS3{F0: 229, F1: 2147483647, F2: 42}
                            _ = v112
                        default:
                            switch (-(fi105)) {
                            case 127:
                                if true {
                                    v109 += v108
                                    if (0.1 <= 1.5) {
                                        v113 := GenS15{F0: 255, F1: true}
                                        var v114 float64 = (float64((((-2.25 + 1.5) + 1.5) - float32(1e10))) / 1.5)
                                        _ = v113
                                        _ = v114
                                    } else {
                                        if true {
                                            var v115 int32 = 660
                                            _ = v115
                                        } else {
                                            var v116 int64 = 294
                                            var v117 uint8 = 127
                                            _ = v116
                                            _ = v117
                                        }
                                        var v118 uint32 = 0
                                        switch (fi104 * fi105) {
                                        case -9223372036854775808:
                                            v119 := GenS7{F0: 219, F1: -1, F2: 0}
                                            _ = v119
                                        case 206:
                                            v120 := GenS1{F0: 100, F1: 42}
                                            _ = v120
                                        default:
                                            if true {
                                                v121 := GenS6{F0: v108, F1: 0, F2: (v118 & (v118 % ((v118) | 1)))}
                                                _ = v121
                                            }
                                        }
                                        _ = v118
                                    }
                                }
                            case 0:
                                var v122 float32 = float32((1.5 + -2.25))
                                _ = v122
                            case -557:
                                v107 = "hello"
                            }
                        }
                        _ = v107
                        _ = v108
                        _ = v109
                    } else {
                        var v123 uint32 = 673
                        _ = v123
                    }
                    _ = fi105
                    _ = v106
                }
                v124 := GenS9{F0: 497, F1: true, F2: fi104}
                _ = v124
            }
        default:
            v1 += v1
        }
        var v125 int16 = v1
        _ = fi104
        _ = v125
    }
    for fi126 := 0; fi126 < 7; fi126++ {
        var v127 int16 = v1
        switch "a" {
        case "a":
            if true {
                v127 = (-(v1))
                if false {
                    for fi128 := 0; fi128 < 1; fi128++ {
                        v129 := v2[0]
                        _ = fi128
                        _ = v129
                    }
                    fmt.Println(v127)
                } else {
                    v1 -= 1
                    var v130 uint32 = 2
                    _ = v130
                }
            }
        }
        var v131 uint16 = 467
        _ = fi126
        _ = v127
        _ = v131
    }
    v1 -= (v1 % ((v1) | 1))
    fmt.Println(v1)
    for fi132 := 0; fi132 < 5; fi132++ {
        v1 = (v1 >> ((0) & 63))
        _ = fi132
    }
    v1 = (v1 << ((1) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
}
