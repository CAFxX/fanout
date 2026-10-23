package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int8; F1 bool; F2 uint8 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int64; F1 uint8 }
type GenS5 struct { F0 uint8; F1 string; F2 uint64 }
type GenS6 struct { F0 string }
type GenS7 struct { F0 uint32; F1 uint; F2 uint16 }
type GenS8 struct { F0 int32; F1 int }
type GenS9 struct { F0 int16 }
type GenS10 struct { F0 string; F1 string; F2 int8 }
func genFn0(p0 uint16, p1 uint8) uint64 {
    for fi5 := 0; fi5 < 17; fi5++ {
        v6 := []int{}
        fmt.Println(p0, p0)
        var v7 uint32 = 2147483648
        _ = fi5
        _ = v6
        _ = v7
    }
    v8 := []int{128}
    v9 := GenS1{F0: -2147483648}
    _ = p0
    _ = p1
    _ = v8
    _ = v9
    return 1
}
func genFn1() uint32 {
    for fi10 := 0; fi10 < 10; fi10++ {
        for fi11 := 0; fi11 < 4; fi11++ {
            fmt.Println(fi10, fi10)
            _ = fi11
        }
        var v12 uint16 = 1
        var v13 string = ("a" + ("hello" + "a"))
        _ = fi10
        _ = v12
        _ = v13
    }
    for fi14 := 0; fi14 < 8; fi14++ {
        for fi15 := 0; fi15 < 4; fi15++ {
            switch ("\n\t" + "foo bar") {
            case "a":
                if false {
                    switch "foo bar" {
                    case "\n\t":
                        var v16 int64 = 9223372036854775807
                        _ = v16
                    default:
                        v17 := GenS2{F0: 36, F1: true, F2: 2}
                        _ = v17
                    }
                    if (1e10 > 3.14159) {
                        if (-2.25 > (0.1 * 3.14159)) {
                            var v18 uint16 = 0
                            var v19 uint8 = 0
                            fmt.Println(v19)
                            _ = v18
                            _ = v19
                        }
                        if false {
                            v20 := []int{582, fi15, (fi15 % ((fi14) | 1))}
                            switch (fi15 + fi15) {
                            case 2147483647:
                                switch (fi14 + fi15) {
                                case 0:
                                    fmt.Println(fi14)
                                    var v21 bool = true
                                    _ = v21
                                case 1:
                                    var v22 uint64 = genFn0(106, 255)
                                    _ = v22
                                default:
                                    var v23 uint = 1
                                    _ = v23
                                }
                            case -2147483648:
                                fmt.Println(fi14, fi15)
                            default:
                                var v24 float64 = 3.14159
                                _ = v24
                            }
                            _ = v20
                        } else {
                            if (true && (-1 >= -274)) {
                                fmt.Println(fi14, fi15)
                            } else {
                                if false {
                                    if (255 >= 255) {
                                        if ("hello" != "\n\t") {
                                            var v25 uint32 = 89
                                            var v26 int = 32768
                                            var v27 uint64 = 1
                                            _ = v25
                                            _ = v26
                                            _ = v27
                                        } else {
                                            if (-20 <= 95) {
                                                v28 := []int{(^(fi15)), 32768, fi14, fi14}
                                                _ = v28
                                            } else {
                                                switch (fi15 + (fi14 & fi14)) {
                                                case 32768:
                                                    v29 := []int{}
                                                    _ = v29
                                                case -1:
                                                    switch "hello" {
                                                    case "foo bar":
                                                        v30 := GenS1{F0: 467}
                                                        var v31 bool = true
                                                        _ = v30
                                                        _ = v31
                                                    case "\n\t":
                                                        if ((1.5 >= float32((((1e10 + -2.25) + (0.0 * ((1e10 + (float64(1.5) - (0.0 * ((0.0 - -2.25) / 3.14159)))) - 3.14159))) + (1e10 / float64(1.5))))) || false) {
                                                            v32 := []int{fi15, fi15}
                                                            if false {
                                                                var v33 uint16 = 42
                                                                var v34 uint = 42
                                                                fmt.Println(v33)
                                                                _ = v33
                                                                _ = v34
                                                            }
                                                            if len(v32) > 0 {
                                                                v32[0] = (^(fi15))
                                                            }
                                                            var v35 uint8 = 42
                                                            switch "foo bar" {
                                                            case "hello":
                                                                v35 = (v35 << ((42) & 63))
                                                            default:
                                                                var v36 uint8 = v35
                                                                if false {
                                                                    v37 := GenS3{F0: -594}
                                                                    if len(v32) >= 2 {
                                                                        v32 = v32[1:len(v32)]
                                                                    }
                                                                    switch "a" {
                                                                    case "\n\t":
                                                                        fmt.Println(fi15, fi14)
                                                                    case "a":
                                                                        v32 = append(v32, fi14)
                                                                        var v38 float64 = 3.14159
                                                                        if (!false) {
                                                                            v36 += 13
                                                                        }
                                                                        _ = v38
                                                                    case "":
                                                                        if (false || false) {
                                                                            if len(v32) > 0 {
                                                                                v32[0] = -1
                                                                            }
                                                                            v39 := GenS2{F0: 127, F1: true, F2: v35}
                                                                            var v40 uint = 1
                                                                            _ = v39
                                                                            _ = v40
                                                                        }
                                                                    default:
                                                                        fmt.Println(fi14, v35, fi15)
                                                                    }
                                                                    var v41 uint32 = 4294967295
                                                                    _ = v37
                                                                    _ = v41
                                                                } else {
                                                                    var v42 uint32 = 42
                                                                    var v43 int32 = 0
                                                                    v42 -= (^(v42))
                                                                    _ = v42
                                                                    _ = v43
                                                                }
                                                                _ = len(v32)
                                                                _ = v36
                                                            }
                                                            _ = v32
                                                            _ = v35
                                                        } else {
                                                            switch "hello" {
                                                            case "":
                                                                var v44 int = fi14
                                                                _ = v44
                                                            default:
                                                                var v45 bool = false
                                                                _ = v45
                                                            }
                                                            var v46 float64 = 1.5
                                                            _ = v46
                                                        }
                                                    case "hello":
                                                        var v47 int64 = 9223372036854775807
                                                        _ = v47
                                                    }
                                                    var v48 uint64 = 4294967295
                                                    fmt.Println(v48)
                                                    _ = v48
                                                case -2147483648:
                                                    var v49 uint64 = 1
                                                    _ = v49
                                                default:
                                                    switch "a" {
                                                    case "hello":
                                                        var v50 int8 = -128
                                                        _ = v50
                                                    }
                                                }
                                                var v51 uint = 576
                                                if false {
                                                    switch "a" {
                                                    case "a":
                                                        var v52 uint64 = 1
                                                        _ = v52
                                                    default:
                                                        v53 := []int{}
                                                        _ = v53
                                                    }
                                                    if (841 == 0) {
                                                        v54 := GenS2{F0: 26, F1: true, F2: 255}
                                                        _ = v54
                                                    }
                                                    if true {
                                                        switch "foo bar" {
                                                        case "":
                                                            var v55 uint = v51
                                                            _ = v55
                                                        case "\n\t":
                                                            v51 = v51
                                                        case "a":
                                                            v51 -= (v51 - v51)
                                                        }
                                                        if (false || true) {
                                                            var v56 float32 = (1e10 * 1.5)
                                                            if (0 <= 2) {
                                                                v57 := GenS1{F0: -1}
                                                                _ = v57
                                                            }
                                                            v58 := GenS4{F0: -1, F1: 1}
                                                            _ = v56
                                                            _ = v58
                                                        } else {
                                                            v51 = 0
                                                            var v59 int16 = -1
                                                            _ = v59
                                                        }
                                                    }
                                                }
                                                if (152 == 197) {
                                                    var v60 uint16 = 275
                                                    switch ("foo bar" + "\n\t") {
                                                    case "hello":
                                                        v61 := []int{(fi15 % ((fi14) | 1))}
                                                        _ = v61
                                                    case "\n\t":
                                                        var v62 float32 = ((1.5 / ((((0.1 * 0.1) / (3.14159 + 0.0)) + (0.0 * 1.5)) / float32((float64(0.1) * (float64((1e10 / 1.0)) - -2.25))))) * (0.1 + ((1.5 - (float32(float64((3.14159 / ((0.1 * -2.25) - float32(1.5))))) * 1.5)) + (-2.25 / (0.0 + float32(((-2.25 / (0.0 + 1e10)) - 0.1)))))))
                                                        _ = v62
                                                    case "a":
                                                        v60 -= v60
                                                    default:
                                                        v60 = 255
                                                    }
                                                    var v63 int = (fi15 & (-(fi15)))
                                                    _ = v60
                                                    _ = v63
                                                } else {
                                                    fmt.Println(v51, fi14, fi15)
                                                    if false {
                                                        if (81 != 128) {
                                                            v64 := GenS2{F0: 42, F1: (uint8(v51) >= 255), F2: 0}
                                                            switch -373 {
                                                            case 32767:
                                                                var v65 uint64 = genFn0(0, 128)
                                                                _ = v65
                                                            case 32768:
                                                                if false {
                                                                    var v66 float32 = 0.1
                                                                    _ = v66
                                                                } else {
                                                                    switch ("\n\t" + "hello") {
                                                                    case "hello":
                                                                        fmt.Println(fi15, fi15, fi15)
                                                                    case "foo bar":
                                                                        v51 = v51
                                                                    case "\n\t":
                                                                        v51 -= v51
                                                                    default:
                                                                        fmt.Println(v51, fi15, fi14)
                                                                    }
                                                                }
                                                            default:
                                                                var v67 uint = v51
                                                                _ = v67
                                                            }
                                                            _ = v64
                                                        } else {
                                                            switch "a" {
                                                            case "a":
                                                                fmt.Println(fi14, fi14)
                                                                switch "\n\t" {
                                                                case "a":
                                                                    v68 := []int{(fi14 >> ((1) & 63)), (-(fi14)), (fi15 + (fi15 * -397)), (fi14 * (fi14 * fi15))}
                                                                    _ = v68
                                                                case "foo bar":
                                                                    v69 := []int{fi14, fi14}
                                                                    _ = v69
                                                                default:
                                                                    var v70 int = fi15
                                                                    _ = v70
                                                                }
                                                                var v71 int32 = -317
                                                                v72 := GenS5{F0: 128, F1: "\n\t", F2: 4294967295}
                                                                var v73 uint16 = 214
                                                                _ = v71
                                                                _ = v72
                                                                _ = v73
                                                            case "":
                                                                if false {
                                                                    v51 -= 1
                                                                    var v74 uint32 = 4294967295
                                                                    _ = v74
                                                                } else {
                                                                    v51 += (^(v51))
                                                                    v75 := GenS1{F0: 42}
                                                                    _ = v75
                                                                }
                                                                v51 += (v51 + 501)
                                                            case "foo bar":
                                                                if true {
                                                                    v51 += v51
                                                                    var v76 bool = false
                                                                    var v77 int32 = -415
                                                                    if ((16 <= 611) || (true || ("foo bar" >= "hello"))) {
                                                                        var v78 int64 = 797
                                                                        var v79 float64 = ((3.14159 * ((1.5 + float64((0.0 + -2.25))) + 0.1)) - 1.5)
                                                                        v77 = (v77 >> (((v51 * v51)) & 63))
                                                                        _ = v78
                                                                        _ = v79
                                                                    } else {
                                                                        var v80 string = ("hello" + "foo bar")
                                                                        _ = v80
                                                                    }
                                                                    v81 := []int{fi15, (fi14 & (fi14 & (^(fi14))))}
                                                                    var v82 uint16 = 2
                                                                    var v83 uint32 = 65535
                                                                    _ = v76
                                                                    _ = v77
                                                                    _ = v81
                                                                    _ = v82
                                                                    _ = v83
                                                                }
                                                            default:
                                                                v51 = 2
                                                            }
                                                            v84 := GenS5{F0: 0, F1: "hello", F2: 80}
                                                            var v85 int16 = 32767
                                                            _ = v84
                                                            _ = v85
                                                        }
                                                    } else {
                                                        var v86 uint16 = 621
                                                        v87 := GenS2{F0: 0, F1: ((true || true) || ((("" + (("hello" + ("" + "")) + ("hello" + ("\n\t" + ("hello" + "a"))))) > "hello") && true)), F2: 4}
                                                        v88 := []int{(fi14 & (fi15 + fi14)), 1, fi15, fi15}
                                                        _ = v86
                                                        _ = v87
                                                        _ = v88
                                                    }
                                                }
                                                _ = v51
                                            }
                                            var v89 int = (fi14 | fi14)
                                            _ = v89
                                        }
                                        if false {
                                            var v90 uint8 = 17
                                            v90 += (v90 & v90)
                                            _ = v90
                                        }
                                    }
                                    v91 := []int{fi14, fi15, 32767, 723}
                                    _ = v91
                                }
                                var v92 uint16 = 32767
                                _ = v92
                            }
                            v93 := []int{fi15, fi14, (fi14 * -9223372036854775808), (fi15 ^ -519)}
                            if false {
                                if len(v93) > 0 {
                                    v93[0] = (fi14 % ((fi15) | 1))
                                }
                                v94 := GenS5{F0: 0, F1: "a", F2: 4294967295}
                                _ = v94
                            }
                            switch "a" {
                            case "hello":
                                if len(v93) >= 2 {
                                    v93 = v93[1:len(v93)]
                                }
                            case "a":
                                v95 := GenS6{F0: (((("hello" + "\n\t") + "a") + "\n\t") + (("" + (("\n\t" + "hello") + ("a" + "foo bar"))) + "\n\t"))}
                                _ = v95
                            case "":
                                var v96 uint16 = 2
                                _ = v96
                            }
                            switch fi15 {
                            case 42:
                                switch (fi14 & -2147483648) {
                                case 36:
                                    var v97 int64 = 996
                                    _ = v97
                                }
                            case 127:
                                v98 := GenS6{F0: ""}
                                _ = v98
                            }
                            var v99 int32 = 42
                            _ = v93
                            _ = v99
                        }
                        v100 := GenS7{F0: 42, F1: 908, F2: 951}
                        _ = v100
                    }
                } else {
                    v101 := GenS7{F0: 2, F1: 255, F2: 789}
                    if false {
                        var v102 bool = false
                        _ = v102
                    } else {
                        var v103 float64 = 0.1
                        _ = v103
                    }
                    switch "" {
                    case "\n\t":
                        switch fi15 {
                        case 128:
                            var v104 uint32 = 648
                            _ = v104
                        case 4294967295:
                            if true {
                                var v105 int32 = -1
                                var v106 uint16 = 65535
                                if (!true) {
                                    v107 := []int{32767, (fi15 ^ (fi14 | fi14))}
                                    v108 := GenS1{F0: (v105 % (((v105 << ((65535) & 63))) | 1))}
                                    _ = v107
                                    _ = v108
                                }
                                switch "\n\t" {
                                case "a":
                                    fmt.Println(v106, v105, fi15)
                                    fmt.Println(v106, fi15, fi14)
                                case "foo bar":
                                    v109 := GenS8{F0: (v105 & (v105 >> ((0) & 63))), F1: 2147483647}
                                    _ = v109
                                default:
                                    fmt.Println(v106)
                                }
                                _ = v105
                                _ = v106
                            } else {
                                switch (-(fi14)) {
                                case -578:
                                    var v110 uint32 = 42
                                    _ = v110
                                case -1:
                                    var v111 int = (fi14 - 2147483647)
                                    _ = v111
                                case 9223372036854775807:
                                    v112 := []int{fi14, fi14, (fi14 + fi15)}
                                    _ = v112
                                default:
                                    var v113 uint = 65535
                                    _ = v113
                                }
                            }
                        }
                    }
                    _ = v101
                }
            }
            _ = fi15
        }
        var v114 float32 = (1.5 + ((-2.25 * (1.5 * 1.5)) - ((1.5 / 1.5) - ((float32(0.1) / -2.25) + ((1e10 - 1e10) / -2.25)))))
        v115 := []int{}
        _ = fi14
        _ = v114
        _ = v115
    }
    return 559
}

func main() {
    v1 := make(map[string]int)
    v2 := make(map[string]bool)
    v3 := make(map[string]uint32)
    var v4 uint = 0
    v116 := GenS9{F0: 32767}
    v117 := GenS5{F0: 21, F1: "foo bar", F2: genFn0(42, 42)}
    if ((false && (!(!false))) || false) {
        v4 -= v4
        v118 := GenS7{F0: 26, F1: (v4 % ((358) | 1)), F2: 61}
        _ = v118
    }
    fmt.Println(v4)
    fmt.Println(v4)
    v4 -= 42
    var v119 float32 = float32((1.5 - 3.14159))
    v120 := []int{}
    if ((v119 != float32(0.0)) && false) {
        fmt.Println(v4, v119)
        switch 32768 {
        case 315:
            var v121 uint64 = 18446744073709551615
            _ = v121
        case 32767:
            v122 := GenS9{F0: -392}
            _ = v122
        }
    }
    switch "a" {
    case "foo bar":
        switch (("hello" + (("\n\t" + ("hello" + "a")) + "foo bar")) + "foo bar") {
        case "hello":
            v4 -= v4
        }
    case "hello":
        v119 += (((v119 / 0.1) + v119) - v119)
    case "\n\t":
        v123 := GenS7{F0: genFn1(), F1: (v4 ^ (^(v4))), F2: 764}
        _ = v123
    default:
        v124 := GenS10{F0: "a", F1: "hello", F2: -1}
        _ = v124
    }
    v119 -= 3.14159
    if true {
        switch ((("hello" + "hello") + ("a" + "a")) + "foo bar") {
        case "":
            if false {
                var v125 int64 = 1
                fmt.Println(v125, v125)
                _ = v125
            }
        case "hello":
            if true {
                for fi126 := 0; fi126 < 5; fi126++ {
                    var v127 float32 = float32(-2.25)
                    _ = fi126
                    _ = v127
                }
                v128 := v1["foo bar"]
                fmt.Println(v119, v119, v4)
                _ = v128
            }
        }
        _ = len(v2)
        var v129 float32 = v119
        _ = v129
    } else {
        for fi130 := 0; fi130 < 11; fi130++ {
            v120 = append(v120, 42)
            _ = fi130
        }
        v4 += (v4 >> ((661) & 63))
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v119:", v119)
    fmt.Println("slicelen v120:", len(v120), v120)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v116
    _ = v117
    _ = v119
    _ = v120
    _ = v2
    _ = v3
    _ = v4
}
