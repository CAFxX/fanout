package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int; F1 int }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int64; F1 uint }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 string; F1 int32; F2 uint64 }
type GenS8 struct { F0 uint64; F1 uint64; F2 int8 }
type GenS9 struct { F0 uint64; F1 int32; F2 int8 }
type GenS10 struct { F0 int32; F1 int32; F2 uint16 }
type GenS11 struct { F0 string; F1 int8; F2 int }
type GenS12 struct { F0 uint32; F1 string; F2 bool }
type GenS13 struct { F0 uint16; F1 int8; F2 string }
type GenS14 struct { F0 uint32; F1 int; F2 int32 }
func genFn0(p0 bool) uint16 {
    p0 = (200 >= 255)
    if (756 <= -613) {
        for fi5 := 0; fi5 < 8; fi5++ {
            v6 := GenS1{F0: 0, F1: "hello"}
            v7 := []int{fi5}
            _ = fi5
            _ = v6
            _ = v7
        }
        for fi8 := 0; fi8 < 6; fi8++ {
            if false {
                var v9 uint8 = 246
                var v10 int64 = -1
                var v11 uint32 = 850
                _ = v10
                _ = v11
                _ = v9
            } else {
                switch 780 {
                case 0:
                    for fi12 := 0; fi12 < 13; fi12++ {
                        if p0 {
                            v13 := []int{fi8, fi12, fi12, fi8}
                            var v14 uint16 = 2
                            var v15 float32 = float32(0.0)
                            _ = v13
                            _ = v14
                            _ = v15
                        } else {
                            v16 := GenS2{F0: -233}
                            p0 = false
                            p0 = (p0 || (0 == 127))
                            _ = v16
                        }
                        _ = fi12
                    }
                case 792:
                    fmt.Println(p0, fi8)
                case 32767:
                    var v17 float64 = (1.5 + float64(0.0))
                    _ = v17
                default:
                    p0 = p0
                }
                var v18 int16 = 1000
                _ = v18
            }
            switch "foo bar" {
            case "":
                v19 := []int{(fi8 >> ((255) & 63)), fi8}
                _ = v19
            case "\n\t":
                fmt.Println(fi8)
            case "foo bar":
                for fi20 := 0; fi20 < 2; fi20++ {
                    var v21 uint8 = 255
                    v22 := GenS3{F0: (^(fi8)), F1: int(v21)}
                    v21 = (v21 >> ((2) & 63))
                    _ = fi20
                    _ = v21
                    _ = v22
                }
            default:
                var v23 uint32 = 2147483647
                _ = v23
            }
            switch "foo bar" {
            case "a":
                switch "foo bar" {
                case "hello":
                    v24 := []int{(fi8 - 2147483647), 4}
                    _ = v24
                case "foo bar":
                    if p0 {
                        switch fi8 {
                        case -243:
                            v25 := GenS2{F0: 42}
                            _ = v25
                        case -1:
                            v26 := GenS2{F0: -1}
                            _ = v26
                        }
                        p0 = p0
                    }
                case "\n\t":
                    switch 2147483647 {
                    case 128:
                        if p0 {
                            for fi27 := 0; fi27 < 10; fi27++ {
                                switch (fi8 | fi8) {
                                case 2147483647:
                                    switch "a" {
                                    case "hello":
                                        p0 = (("" + ("a" + "foo bar")) < "foo bar")
                                    case "\n\t":
                                        if true {
                                            var v28 int32 = -2147483648
                                            _ = v28
                                        } else {
                                            v29 := []int{fi8}
                                            var v30 bool = p0
                                            if len(v29) >= 2 {
                                                v29 = v29[1:len(v29)]
                                            }
                                            _ = v29
                                            _ = v30
                                        }
                                    default:
                                        v31 := []int{fi8}
                                        _ = v31
                                    }
                                case 42:
                                    var v32 uint64 = 60
                                    switch (^(fi27)) {
                                    case 886:
                                        v32 -= v32
                                    }
                                    switch "" {
                                    case "":
                                        v32 += v32
                                    default:
                                        if (true && p0) {
                                            if (!(p0 || false)) {
                                                fmt.Println(fi27, fi27, fi8)
                                                if (false || (!(0 <= 997))) {
                                                    switch ("" + "") {
                                                    case "":
                                                        v32 += (v32 * v32)
                                                    default:
                                                        if ((fi8 - fi8) >= (fi27 + fi8)) {
                                                            var v33 uint = 42
                                                            var v34 int8 = -68
                                                            if true {
                                                                if (!(!true)) {
                                                                    if (1 > 127) {
                                                                        var v35 int16 = -635
                                                                        _ = v35
                                                                    } else {
                                                                        v36 := GenS4{F0: (!p0)}
                                                                        _ = v36
                                                                    }
                                                                    v37 := GenS5{F0: 823, F1: v33}
                                                                    v34 = (v34 >> ((v33) & 63))
                                                                    v38 := GenS2{F0: -1}
                                                                    _ = v37
                                                                    _ = v38
                                                                }
                                                            } else {
                                                                switch (("foo bar" + "") + (("" + ("" + ("hello" + ""))) + "a")) {
                                                                case "a":
                                                                    v32 = 0
                                                                case "\n\t":
                                                                    var v39 int32 = -1
                                                                    _ = v39
                                                                default:
                                                                    v40 := []int{(fi27 >> (((v33 - v33)) & 63))}
                                                                    _ = v40
                                                                }
                                                                v34 = 0
                                                                v41 := GenS3{F0: -2147483648, F1: (fi8 + (^(fi27)))}
                                                                _ = v41
                                                            }
                                                            if (1e10 < float64(float32((-2.25 - 1.5)))) {
                                                                v42 := []int{}
                                                                var v43 uint64 = v32
                                                                _ = v42
                                                                _ = v43
                                                            }
                                                            v44 := GenS4{F0: p0}
                                                            var v45 int = 454
                                                            v45 -= -9223372036854775808
                                                            _ = v33
                                                            _ = v34
                                                            _ = v44
                                                            _ = v45
                                                        }
                                                    }
                                                    v46 := []int{}
                                                    _ = v46
                                                }
                                            } else {
                                                v47 := GenS3{F0: (-(fi8)), F1: (fi8 >> ((141) & 63))}
                                                _ = v47
                                            }
                                            p0 = (-1 == -32768)
                                            v48 := []int{}
                                            var v49 float64 = (1.5 + (0.0 + ((-2.25 / (3.14159 / 1.5)) + 3.14159)))
                                            p0 = p0
                                            _ = v48
                                            _ = v49
                                        }
                                        fmt.Println(p0, fi8, fi8)
                                        v50 := []int{}
                                        v32 += (v32 - (v32 / ((v32) | 1)))
                                        _ = v50
                                    }
                                    _ = v32
                                case -9223372036854775808:
                                    var v51 int8 = 127
                                    _ = v51
                                default:
                                    fmt.Println(fi27, fi8)
                                    switch (fi27 & fi8) {
                                    case 32767:
                                        v52 := []int{}
                                        _ = v52
                                    case -872:
                                        v53 := []int{fi8, fi27, fi8}
                                        _ = v53
                                    case 653:
                                        switch "a" {
                                        case "":
                                            switch fi27 {
                                            case -262:
                                                var v54 uint8 = 128
                                                _ = v54
                                            case 753:
                                                var v55 int = -2147483648
                                                _ = v55
                                            }
                                        default:
                                            if false {
                                                switch fi27 {
                                                case -67:
                                                    fmt.Println(fi27, fi27, p0)
                                                    var v56 float32 = 1.5
                                                    v57 := []int{}
                                                    _ = v56
                                                    _ = v57
                                                case 4294967295:
                                                    var v58 uint64 = 42
                                                    _ = v58
                                                case -192:
                                                    switch fi8 {
                                                    case 76:
                                                        var v59 int16 = 786
                                                        _ = v59
                                                    case 9223372036854775807:
                                                        fmt.Println(fi8)
                                                    case 0:
                                                        p0 = p0
                                                    default:
                                                        var v60 int = fi27
                                                        _ = v60
                                                    }
                                                }
                                            } else {
                                                var v61 uint8 = 42
                                                _ = v61
                                            }
                                        }
                                    }
                                }
                                _ = fi27
                            }
                            fmt.Println(p0, fi8)
                            for fi62 := 0; fi62 < 17; fi62++ {
                                v63 := []int{(fi62 / (((fi8 / ((0) | 1))) | 1)), (fi62 * fi62), (fi62 ^ -1)}
                                p0 = false
                                p0 = p0
                                var v64 int8 = -1
                                switch (fi62 - (fi62 ^ fi8)) {
                                case 910:
                                    if p0 {
                                        if false {
                                            fmt.Println(v64, fi8)
                                            p0 = p0
                                            p0 = (false || (0 < 0))
                                        } else {
                                            var v65 uint = 451
                                            _ = v65
                                        }
                                        if (!((-2.25 / -2.25) == 0.0)) {
                                            var v66 bool = (true || p0)
                                            _ = v66
                                        } else {
                                            var v67 int = fi62
                                            if p0 {
                                                var v68 float64 = 3.14159
                                                switch v67 {
                                                case 9223372036854775807:
                                                    fmt.Println(v68, p0, v68)
                                                case 42:
                                                    v69 := GenS6{F0: 58}
                                                    _ = v69
                                                case -2147483648:
                                                    if (!(18446744073709551615 <= 664)) {
                                                        fmt.Println(v68)
                                                        p0 = (((868 < 2) && (42 > 763)) && p0)
                                                        var v70 string = "a"
                                                        _ = v70
                                                    } else {
                                                        if len(v63) > 0 {
                                                            v63[0] = 128
                                                        }
                                                        if p0 {
                                                            v71 := GenS2{F0: int32(v64)}
                                                            var v72 int8 = (v64 % ((v64) | 1))
                                                            if true {
                                                                v73 := GenS1{F0: 42, F1: "foo bar"}
                                                                switch (^(v67)) {
                                                                case 1:
                                                                    v74 := GenS7{F0: "", F1: 42, F2: 1000}
                                                                    _ = v74
                                                                }
                                                                _ = v73
                                                            } else {
                                                                v75 := GenS7{F0: "", F1: 0, F2: 465}
                                                                _ = v75
                                                            }
                                                            _ = v71
                                                            _ = v72
                                                        }
                                                    }
                                                default:
                                                    if p0 {
                                                        var v76 int = (-(v67))
                                                        fmt.Println(p0, v68)
                                                        _ = v76
                                                    }
                                                }
                                                v77 := GenS8{F0: 1, F1: 0, F2: (v64 | -55)}
                                                _ = v68
                                                _ = v77
                                            } else {
                                                v63 = append(v63, fi8)
                                                v78 := GenS8{F0: 1, F1: 2, F2: (v64 / ((v64) | 1))}
                                                _ = v78
                                            }
                                            if len(v63) > 0 {
                                                v63[0] = 127
                                            }
                                            switch (fi8 | (v67 - 4294967295)) {
                                            case 1:
                                                if false {
                                                    switch v67 {
                                                    case -9223372036854775808:
                                                        _ = len(v63)
                                                    case 4294967295:
                                                        v79 := GenS9{F0: 190, F1: 42, F2: v64}
                                                        _ = v79
                                                    case -2147483648:
                                                        v63 = append(v63, fi62)
                                                    default:
                                                        var v80 string = ("" + "a")
                                                        _ = v80
                                                    }
                                                    if len(v63) >= 2 {
                                                        v63 = v63[1:len(v63)]
                                                    }
                                                    if p0 {
                                                        v67 = (-(fi62))
                                                    }
                                                } else {
                                                    if ((!(true || p0)) && ("" >= "foo bar")) {
                                                        fmt.Println(p0, v67)
                                                    } else {
                                                        var v81 bool = (!p0)
                                                        _ = v81
                                                    }
                                                }
                                            case -414:
                                                if true {
                                                    var v82 uint8 = 0
                                                    _ = v82
                                                } else {
                                                    if false {
                                                        p0 = false
                                                        if p0 {
                                                            var v83 uint = 65535
                                                            if p0 {
                                                                var v84 int32 = 42
                                                                var v85 string = "a"
                                                                v86 := GenS4{F0: (!p0)}
                                                                _ = v84
                                                                _ = v85
                                                                _ = v86
                                                            }
                                                            _ = v83
                                                        }
                                                        var v87 int32 = 2147483647
                                                        _ = v87
                                                    }
                                                    v88 := GenS10{F0: -596, F1: 42, F2: 459}
                                                    v89 := GenS8{F0: 4294967295, F1: 0, F2: -1}
                                                    _ = v88
                                                    _ = v89
                                                }
                                            }
                                            var v90 uint = 255
                                            _ = v67
                                            _ = v90
                                        }
                                        v91 := GenS10{F0: 0, F1: -942, F2: 2}
                                        _ = v91
                                    } else {
                                        switch "hello" {
                                        case "hello":
                                            fmt.Println(v64, v64, v64)
                                            if (4294967295 == 42) {
                                                if len(v63) > 0 {
                                                    v63[0] = fi62
                                                }
                                                if (!(0 != 42)) {
                                                    if p0 {
                                                        var v92 uint16 = 255
                                                        _ = v92
                                                    }
                                                } else {
                                                    p0 = false
                                                }
                                            } else {
                                                switch (fi8 / ((fi62) | 1)) {
                                                case 2147483647:
                                                    p0 = (65535 != 0)
                                                }
                                                _ = len(v63)
                                            }
                                        case "a":
                                            var v93 uint16 = 255
                                            _ = v93
                                        case "foo bar":
                                            v94 := GenS1{F0: 1, F1: (("a" + "") + ("hello" + "foo bar"))}
                                            _ = v94
                                        }
                                        v95 := GenS8{F0: 0, F1: 0, F2: -128}
                                        v63 = append(v63, fi62)
                                        _ = v95
                                    }
                                    var v96 string = "foo bar"
                                    v64 = v64
                                    if true {
                                        if (((^(v64)) >= -128) && p0) {
                                            v97 := GenS2{F0: -1}
                                            v98 := GenS11{F0: v96, F1: v64, F2: -233}
                                            var v99 string = v96
                                            _ = v97
                                            _ = v98
                                            _ = v99
                                        }
                                        var v100 uint16 = 65535
                                        _ = v100
                                    }
                                    var v101 int8 = v64
                                    _ = v101
                                    _ = v96
                                case -2147483648:
                                    fmt.Println(fi62)
                                default:
                                    v64 += -128
                                }
                                _ = fi62
                                _ = v63
                                _ = v64
                            }
                        } else {
                            if p0 {
                                fmt.Println(fi8, p0)
                                if p0 {
                                    v102 := GenS12{F0: 2147483648, F1: ("foo bar" + ("a" + "hello")), F2: p0}
                                    if (818 != 802) {
                                        p0 = (p0 && (!p0))
                                        v103 := []int{fi8, -1}
                                        _ = v103
                                    }
                                    var v104 float64 = (1.5 * (1.5 - float64(float32(0.0))))
                                    _ = v102
                                    _ = v104
                                }
                                var v105 bool = false
                                _ = v105
                            }
                            p0 = (p0 && true)
                            if true {
                                if p0 {
                                    for fi106 := 0; fi106 < 2; fi106++ {
                                        var v107 int16 = 1000
                                        _ = fi106
                                        _ = v107
                                    }
                                    p0 = false
                                    for fi108 := 0; fi108 < 3; fi108++ {
                                        v109 := GenS10{F0: -416, F1: 42, F2: 42}
                                        v110 := GenS2{F0: 0}
                                        _ = fi108
                                        _ = v109
                                        _ = v110
                                    }
                                }
                                v111 := []int{-584, 635, fi8, (^(fi8))}
                                _ = v111
                            }
                        }
                    case 32768:
                        fmt.Println(p0)
                    }
                }
            case "\n\t":
                if (!p0) {
                    var v112 uint = 18446744073709551615
                    var v113 float32 = float32(1.5)
                    var v114 int32 = -2147483648
                    _ = v112
                    _ = v113
                    _ = v114
                } else {
                    for fi115 := 0; fi115 < 12; fi115++ {
                        p0 = p0
                        _ = fi115
                    }
                    v116 := GenS13{F0: 255, F1: 1, F2: "foo bar"}
                    _ = v116
                }
            }
            _ = fi8
        }
    } else {
        var v117 uint64 = 1
        v117 -= (v117 + 1000)
        _ = v117
    }
    v118 := []int{9223372036854775807}
    _ = len(v118)
    p0 = p0
    _ = p0
    _ = v118
    return 740
}

func main() {
    var v1 uint8 = 1
    var v2 int = 93
    v3 := make(map[uint32]int)
    v4 := make(map[string]uint32)
    v1 -= v1
    v119 := []int{(v2 | -581), v2, v2, (v2 - 32767)}
    var v120 uint16 = 255
    v1 = 128
    v2 += 32768
    if false {
        v1 -= (v1 & 2)
        v2 = (v2 ^ int(v1))
        for fi121 := 0; fi121 < 11; fi121++ {
            if (9223372036854775807 > 42) {
                var v122 string = "a"
                v120 = (^(v120))
                var v123 int16 = 1000
                _ = v122
                _ = v123
            } else {
                if (42 > 2147483647) {
                    var v124 string = ""
                    v120 = 87
                    _ = v124
                }
                v125 := GenS14{F0: 94, F1: -566, F2: 943}
                _ = v125
            }
            for fi126 := 0; fi126 < 12; fi126++ {
                v120 += genFn0((255 == 993))
                fmt.Println(v120, fi121)
                var v127 float32 = (0.1 - (3.14159 * (float32((0.0 - 1e10)) + 0.1)))
                _ = fi126
                _ = v127
            }
            _ = fi121
        }
    } else {
        v4["a"] = uint32(v120)
    }
    if (!(-1 != 100)) {
        for fi128 := 0; fi128 < 1; fi128++ {
            var v129 float64 = 1e10
            v130 := GenS4{F0: true}
            var v131 uint16 = (v120 & genFn0(false))
            _ = fi128
            _ = v129
            _ = v130
            _ = v131
        }
    }
    fmt.Println(v120, v2)
    v120 -= v120
    var v132 int = 9223372036854775807
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v119:", len(v119), v119)
    fmt.Println("v120:", v120)
    fmt.Println("v132:", v132)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v119
    _ = v120
    _ = v132
    _ = v2
    _ = v3
    _ = v4
}
