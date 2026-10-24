package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32; F2 uint }
type GenS2 struct { F0 uint8; F1 bool; F2 int }
type GenS3 struct { F0 int; F1 uint16 }
type GenS4 struct { F0 int32; F1 bool }
type GenS5 struct { F0 int16; F1 int16; F2 int }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 int }
type GenS8 struct { F0 string }
type GenS9 struct { F0 int64; F1 uint64; F2 int }
type GenS10 struct { F0 uint16; F1 string; F2 uint8 }
type GenS11 struct { F0 int64 }
type GenS12 struct { F0 bool; F1 bool }
type GenS13 struct { F0 int16 }
type GenS14 struct { F0 int; F1 uint8 }
func genFn0(p0 uint16, p1 string) int {
    v7 := GenS1{F0: "hello", F1: 2, F2: 65535}
    v8 := GenS2{F0: 91, F1: (0.1 == (1.5 - (-2.25 + (0.0 / 1.0)))), F2: 128}
    var v9 float64 = 1e10
    v10 := GenS3{F0: 1, F1: p0}
    _ = p0
    _ = p1
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return -200
}

func main() {
    v1 := make(map[int]uint32)
    var v2 uint16 = 1
    var v3 int8 = 42
    v4 := make(map[uint32]uint32)
    v5 := make(map[string]bool)
    v6 := []int{}
    for fi11 := 0; fi11 < 19; fi11++ {
        v3 += v3
        v3 += 127
        _ = fi11
    }
    v2 -= v2
    if (491 < 42) {
        switch "\n\t" {
        case "foo bar":
            if (1 == 32767) {
                if true {
                    v3 = (^(v3))
                    var v12 int64 = -278
                    _ = v12
                } else {
                    for fi13 := 0; fi13 < 9; fi13++ {
                        switch ((("a" + "foo bar") + (("\n\t" + "hello") + "")) + "foo bar") {
                        case "":
                            v14 := GenS3{F0: (fi13 & 128), F1: (^(v2))}
                            _ = v14
                        case "a":
                            var v15 uint16 = v2
                            _ = v15
                        case "foo bar":
                            v16 := GenS4{F0: 972, F1: ((1.5 * 0.0) != 0.0)}
                            _ = v16
                        default:
                            if true {
                                var v17 uint64 = 18446744073709551615
                                _ = v17
                            } else {
                                v18 := GenS4{F0: -2147483648, F1: ((v3 - v3) <= v3)}
                                _ = v18
                            }
                        }
                        fmt.Println(v3, v3)
                        var v19 uint8 = 127
                        _ = fi13
                        _ = v19
                    }
                }
                var v20 uint16 = v2
                if (-637 < 327) {
                    var v21 int32 = 2147483647
                    fmt.Println(v21)
                    _ = v21
                }
                _ = v20
            }
        case "":
            v4[313] = 722
        default:
            if true {
                v22 := v1[42]
                v6 = append(v6, 32767)
                v3 = v3
                _ = v22
            }
        }
        fmt.Println(v2)
        v23 := GenS5{F0: 149, F1: -32768, F2: 128}
        _ = v23
    }
    if true {
        v3 = (v3 - (v3 - 0))
        for fi24 := 0; fi24 < 2; fi24++ {
            if (!(!false)) {
                var v25 int16 = -32768
                _ = v25
            } else {
                switch genFn0(758, ("" + "")) {
                case -2147483648:
                    fmt.Println(fi24)
                default:
                    switch "foo bar" {
                    case "foo bar":
                        switch fi24 {
                        case 32768:
                            var v26 int = 32767
                            _ = v26
                        case 42:
                            v2 -= 2
                        }
                    case "\n\t":
                        switch ("\n\t" + "a") {
                        case "hello":
                            if true {
                                v27 := GenS6{F0: v2}
                                if false {
                                    if (42 < 0) {
                                        fmt.Println(v2, fi24)
                                        for fi28 := 0; fi28 < 18; fi28++ {
                                            v5[("hello" + ((("foo bar" + "") + "\n\t") + ("\n\t" + ("foo bar" + ((("hello" + (((("a" + "foo bar") + ("a" + (("a" + "") + "hello"))) + "hello") + ("\n\t" + "a"))) + "\n\t") + "hello")))))] = false
                                            fmt.Println(v3, v2)
                                            fmt.Println(v3, v2, v2)
                                            var v29 int64 = -294
                                            switch "\n\t" {
                                            case "foo bar":
                                                if len(v6) >= 2 {
                                                    v6 = v6[1:len(v6)]
                                                }
                                            default:
                                                switch 128 {
                                                case -1:
                                                    if len(v6) >= 2 {
                                                        v6 = v6[1:len(v6)]
                                                    }
                                                case -804:
                                                    var v30 uint32 = 522
                                                    _ = v30
                                                }
                                            }
                                            if (-2.25 < (((float64((1.5 * ((0.0 + 1.5) - 1.5))) * 0.0) * 1.5) + -2.25)) {
                                                v2 -= (v2 & 42)
                                                if len(v6) > 0 {
                                                    v6[0] = (fi24 + -2147483648)
                                                }
                                                if true {
                                                    v3 = v3
                                                    v2 -= (^(v2))
                                                    v31 := v5["a"]
                                                    var v32 uint8 = 128
                                                    var v33 uint16 = v2
                                                    var v34 uint64 = 847
                                                    if v31 {
                                                        var v35 string = "foo bar"
                                                        _ = v35
                                                    } else {
                                                        var v36 uint = 745
                                                        _ = v36
                                                    }
                                                    _ = len(v5)
                                                    _ = v31
                                                    _ = v32
                                                    _ = v33
                                                    _ = v34
                                                } else {
                                                    if true {
                                                        if ((!true) || (-1 == v3)) {
                                                            v37 := GenS2{F0: 2, F1: true, F2: 108}
                                                            fmt.Println(fi24, v29, v2)
                                                            v29 -= 1
                                                            if false {
                                                                v38 := GenS7{F0: fi28}
                                                                _ = v38
                                                            }
                                                            _ = v37
                                                        } else {
                                                            var v39 uint64 = 0
                                                            var v40 int16 = 315
                                                            _ = v39
                                                            _ = v40
                                                        }
                                                        fmt.Println(v3)
                                                        fmt.Println(fi28, v3, v2)
                                                        if (true || true) {
                                                            v3 = (v3 % ((v3) | 1))
                                                            switch "hello" {
                                                            case "foo bar":
                                                                v29 = v29
                                                            }
                                                            fmt.Println(fi28, v3, v3)
                                                        }
                                                    } else {
                                                        var v41 uint64 = 4294967295
                                                        if len(v6) >= 2 {
                                                            v6 = v6[1:len(v6)]
                                                        }
                                                        var v42 float32 = 1e10
                                                        _ = v41
                                                        _ = v42
                                                    }
                                                    v4[2147483647] = 488
                                                    v3 -= v3
                                                }
                                            } else {
                                                v29 += v29
                                                var v43 int8 = (v3 + v3)
                                                _ = len(v5)
                                                _ = len(v6)
                                                _ = v43
                                            }
                                            if (1 > -614) {
                                                v29 += v29
                                                var v44 int = (fi24 % (((fi24 - (fi24 / ((42) | 1)))) | 1))
                                                _ = v44
                                            } else {
                                                v2 -= (v2 + v2)
                                                var v45 int8 = int8(v29)
                                                _ = v45
                                            }
                                            _ = fi28
                                            _ = v29
                                        }
                                        fmt.Println(fi24)
                                    }
                                } else {
                                    switch fi24 {
                                    case 278:
                                        if false {
                                            v3 += v3
                                            fmt.Println(fi24)
                                            v46 := GenS3{F0: (fi24 + (fi24 >> ((2) & 63))), F1: (^(v2))}
                                            _ = v46
                                        } else {
                                            var v47 uint8 = 128
                                            for fi48 := 0; fi48 < 17; fi48++ {
                                                v6 = append(v6, (-(fi24)))
                                                if (!false) {
                                                    if ((-857 > 32767) && true) {
                                                        v49 := GenS4{F0: -666, F1: false}
                                                        _ = v49
                                                    } else {
                                                        v2 = 32767
                                                    }
                                                    if false {
                                                        v2 += v2
                                                        v2 = 1
                                                        fmt.Println(v3, v3, v3)
                                                    } else {
                                                        var v50 int16 = 42
                                                        var v51 uint64 = 0
                                                        if (v3 <= 100) {
                                                            var v52 uint8 = 2
                                                            v53 := GenS8{F0: ("\n\t" + ("foo bar" + ""))}
                                                            v54 := GenS4{F0: -2147483648, F1: false}
                                                            _ = v52
                                                            _ = v53
                                                            _ = v54
                                                        }
                                                        _ = v50
                                                        _ = v51
                                                    }
                                                } else {
                                                    var v55 int16 = 32767
                                                    v3 = v3
                                                    _ = v55
                                                }
                                                if true {
                                                    v5["hello"] = false
                                                }
                                                _ = fi48
                                            }
                                            _ = v47
                                        }
                                    case -9223372036854775808:
                                        _ = len(v4)
                                    case 275:
                                        fmt.Println(v3, v2)
                                    }
                                }
                                _ = v27
                            }
                        case "foo bar":
                            for fi56 := 0; fi56 < 6; fi56++ {
                                if (!(947 >= 4294967295)) {
                                    v2 = v2
                                } else {
                                    v57 := GenS1{F0: (("hello" + "\n\t") + ""), F1: 42, F2: 0}
                                    fmt.Println(fi24, fi56)
                                    switch "a" {
                                    case "a":
                                        v2 = (v2 >> ((255) & 63))
                                    case "hello":
                                        if (("foo bar" + "") != "foo bar") {
                                            v58 := GenS9{F0: 865, F1: 1, F2: fi24}
                                            _ = v58
                                        }
                                    case "":
                                        v59 := GenS7{F0: 9223372036854775807}
                                        _ = v59
                                    default:
                                        if (!true) {
                                            delete(v4, 229)
                                            var v60 uint64 = 42
                                            v2 += v2
                                            var v61 uint32 = 555
                                            v3 += (v3 + 100)
                                            _ = v60
                                            _ = v61
                                        }
                                    }
                                    _ = v57
                                }
                                if true {
                                    delete(v4, 0)
                                    if true {
                                        var v62 uint16 = 583
                                        if true {
                                            if true {
                                                v62 = (v2 >> ((65535) & 63))
                                                var v63 int32 = -2147483648
                                                _ = v63
                                            } else {
                                                var v64 float32 = (-2.25 + -2.25)
                                                var v65 int32 = -98
                                                _ = v64
                                                _ = v65
                                            }
                                        }
                                        var v66 uint = 42
                                        _ = v62
                                        _ = v66
                                    }
                                    fmt.Println(fi24, fi56, fi24)
                                } else {
                                    var v67 int = fi24
                                    if len(v6) >= 2 {
                                        v6 = v6[1:len(v6)]
                                    }
                                    v67 += (v67 * 608)
                                    _ = v67
                                }
                                var v68 bool = (!(1 == -1))
                                _ = fi56
                                _ = v68
                            }
                        case "\n\t":
                            var v69 int8 = -128
                            _ = v69
                        default:
                            var v70 uint = 65535
                            _ = v70
                        }
                    default:
                        switch ("hello" + "foo bar") {
                        case "foo bar":
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                        case "\n\t":
                            var v71 int = fi24
                            _ = v71
                        default:
                            v3 -= (^(v3))
                        }
                    }
                }
            }
            v3 = (v3 / ((0) | 1))
            _ = fi24
        }
    } else {
        v3 = 0
        var v72 uint32 = 4294967295
        _ = v72
    }
    v73 := GenS10{F0: (v2 | 42), F1: "\n\t", F2: 128}
    var v74 int32 = -2147483648
    v4[306] = 2147483647
    if (-377 < (-(v74))) {
        for fi75 := 0; fi75 < 6; fi75++ {
            switch -2147483648 {
            case 127:
                if false {
                    for fi76 := 0; fi76 < 20; fi76++ {
                        var v77 int64 = 0
                        _ = fi76
                        _ = v77
                    }
                    v3 = 127
                    if false {
                        for fi78 := 0; fi78 < 13; fi78++ {
                            if (((fi75 + fi75) != (^(fi75))) && (-2.25 < 1.5)) {
                                v79 := GenS2{F0: uint8(v2), F1: false, F2: 853}
                                if (255 != 94) {
                                    v80 := GenS11{F0: 1}
                                    v81 := GenS5{F0: 1, F1: 1000, F2: fi78}
                                    _ = v80
                                    _ = v81
                                } else {
                                    fmt.Println(v74, v2)
                                    if (!true) {
                                        var v82 float64 = 1.5
                                        _ = v82
                                    } else {
                                        v4[91] = 65535
                                        v83 := GenS3{F0: fi78, F1: (^(v2))}
                                        _ = v83
                                    }
                                }
                                if true {
                                    switch 4294967295 {
                                    case 9223372036854775807:
                                        v84 := GenS12{F0: true, F1: true}
                                        _ = v84
                                    }
                                    switch ((("hello" + "\n\t") + "foo bar") + (("\n\t" + "foo bar") + "hello")) {
                                    case "\n\t":
                                        var v85 string = ""
                                        _ = v85
                                    case "a":
                                        fmt.Println(v3, v2, v2)
                                    case "foo bar":
                                        v1[127] = 4294967295
                                    default:
                                        _ = len(v6)
                                    }
                                    var v86 float32 = ((3.14159 / (-2.25 * 1e10)) / -2.25)
                                    _ = v86
                                } else {
                                    v87 := GenS11{F0: -9223372036854775808}
                                    var v88 int64 = 1
                                    if len(v6) > 0 {
                                        v6[0] = genFn0(1, "")
                                    }
                                    _ = v87
                                    _ = v88
                                }
                                if (2 > 42) {
                                    var v89 int32 = (v74 << ((1) & 63))
                                    fmt.Println(v2, fi78, fi78)
                                    _ = v89
                                } else {
                                    var v90 float32 = 0.0
                                    if false {
                                        v3 = 127
                                        var v91 bool = (-1 == -1)
                                        v91 = true
                                        _ = v91
                                    } else {
                                        var v92 int64 = 9223372036854775807
                                        v92 += (v92 % ((0) | 1))
                                        if len(v6) >= 2 {
                                            v6 = v6[1:len(v6)]
                                        }
                                        var v93 uint = 42
                                        v93 += (v93 >> ((v93) & 63))
                                        _ = v92
                                        _ = v93
                                    }
                                    var v94 int8 = v3
                                    _ = v90
                                    _ = v94
                                }
                                var v95 float64 = 0.1
                                _ = v79
                                _ = v95
                            } else {
                                v2 -= v2
                                if true {
                                    v74 -= v74
                                    if (756 < 2) {
                                        v2 -= v2
                                        fmt.Println(fi75, v74, v3)
                                        if (!true) {
                                            var v96 int32 = 2147483647
                                            _ = v96
                                        }
                                        switch "foo bar" {
                                        case "":
                                            var v97 uint = 730
                                            _ = v97
                                        case "hello":
                                            switch 9223372036854775807 {
                                            case -1:
                                                var v98 uint64 = 710
                                                _ = v98
                                            default:
                                                var v99 int8 = v3
                                                _ = v99
                                            }
                                        }
                                        v100 := GenS7{F0: (^(fi75))}
                                        _ = v100
                                    } else {
                                        var v101 uint16 = (v2 | (v2 << ((2) & 63)))
                                        _ = v101
                                    }
                                } else {
                                    v102 := GenS13{F0: -59}
                                    _ = v102
                                }
                                var v103 uint64 = 501
                                _ = v103
                            }
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                            if false {
                                var v104 float32 = (1e10 + (0.0 * 1e10))
                                v2 -= (v2 << ((255) & 63))
                                v105 := GenS14{F0: fi78, F1: 2}
                                _ = v104
                                _ = v105
                            } else {
                                if (v2 == v2) {
                                    if (!("hello" >= ("a" + ""))) {
                                        switch "" {
                                        case "hello":
                                            if (!false) {
                                                fmt.Println(v74)
                                                var v106 uint32 = 2147483647
                                                if false {
                                                    v3 += int8(fi75)
                                                    if false {
                                                        switch (^(fi75)) {
                                                        case 127:
                                                            if false {
                                                                v3 = v3
                                                            } else {
                                                                v3 += v3
                                                                if (128 >= 146) {
                                                                    if (("a" + (("\n\t" + "hello") + ("a" + "\n\t"))) >= "foo bar") {
                                                                        if ((3.14159 * (0.0 - 0.0)) == ((1.5 / 1.5) + -2.25)) {
                                                                            var v107 float32 = (-2.25 + 0.0)
                                                                            if true {
                                                                                v108 := GenS7{F0: genFn0((v2 << ((2) & 63)), "hello")}
                                                                                var v109 float32 = float32(float64(0.0))
                                                                                var v110 int16 = 1000
                                                                                v106 -= (v106 * v106)
                                                                                fmt.Println(v110, v74)
                                                                                _ = v108
                                                                                _ = v109
                                                                                _ = v110
                                                                            } else {
                                                                                var v111 float32 = ((float32(1.5) / 1.0) - float32((0.1 + ((1e10 + 0.1) / 1.5))))
                                                                                var v112 string = ("foo bar" + "hello")
                                                                                _ = v111
                                                                                _ = v112
                                                                            }
                                                                            v3 -= 0
                                                                            _ = v107
                                                                        }
                                                                    }
                                                                } else {
                                                                    v106 = v106
                                                                }
                                                                v113 := GenS8{F0: (("a" + "") + ("hello" + ""))}
                                                                _ = v113
                                                            }
                                                        case 1:
                                                            if (true || (!(-256 != 32767))) {
                                                                var v114 uint32 = v106
                                                                var v115 bool = (false && ((("" + ((("\n\t" + "a") + "") + "foo bar")) + ("foo bar" + ("hello" + ("\n\t" + "a")))) != ("" + "foo bar")))
                                                                _ = len(v1)
                                                                _ = v114
                                                                _ = v115
                                                            }
                                                        case -2147483648:
                                                            if len(v6) > 0 {
                                                                v6[0] = 786
                                                            }
                                                        }
                                                        var v116 float64 = 1e10
                                                        v106 = (v106 + 2)
                                                        _ = v116
                                                    } else {
                                                        fmt.Println(fi78, v106, v2)
                                                    }
                                                    if false {
                                                        fmt.Println(v106)
                                                        v117 := v1[fi78]
                                                        _ = v117
                                                    } else {
                                                        var v118 int16 = 0
                                                        _ = v118
                                                    }
                                                    if false {
                                                        var v119 string = ""
                                                        if false {
                                                            switch fi75 {
                                                            case -405:
                                                                v120 := v5[v119]
                                                                _ = v120
                                                            }
                                                        }
                                                        _ = v119
                                                    } else {
                                                        fmt.Println(v106, v2, v2)
                                                        var v121 uint8 = 176
                                                        _ = v121
                                                    }
                                                    switch "a" {
                                                    case "a":
                                                        var v122 bool = (2 != 42)
                                                        _ = v122
                                                    case "hello":
                                                        v74 = (v74 & 159)
                                                    case "":
                                                        var v123 uint64 = 4294967295
                                                        _ = v123
                                                    default:
                                                        var v124 string = ""
                                                        _ = v124
                                                    }
                                                }
                                                _ = v106
                                            } else {
                                                v125 := GenS11{F0: 0}
                                                v3 = v3
                                                _ = v125
                                            }
                                            var v126 uint64 = 1
                                            _ = v126
                                        case "\n\t":
                                            v127 := v1[(-(fi75))]
                                            _ = v127
                                        }
                                    }
                                }
                                switch "foo bar" {
                                case "foo bar":
                                    var v128 uint8 = 255
                                    _ = v128
                                default:
                                    delete(v5, "a")
                                }
                            }
                            _ = fi78
                        }
                        for fi129 := 0; fi129 < 2; fi129++ {
                            switch fi75 {
                            case 9223372036854775807:
                                v6 = append(v6, (fi75 + int(v3)))
                            }
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                            var v130 int64 = 1
                            _ = fi129
                            _ = v130
                        }
                    } else {
                        for fi131 := 0; fi131 < 12; fi131++ {
                            var v132 int64 = 42
                            _ = fi131
                            _ = v132
                        }
                        var v133 float32 = 0.0
                        _ = v133
                    }
                }
            case -1:
                v6 = append(v6, (fi75 * fi75))
            case 681:
                v134 := GenS11{F0: -1}
                _ = v134
            default:
                var v135 int = fi75
                _ = v135
            }
            v2 -= v2
            fmt.Println(v2, fi75, fi75)
            _ = fi75
        }
        fmt.Println(v74, v2, v74)
    } else {
        v2 -= v2
    }
    if (true || (!(!false))) {
        fmt.Println(v3, v2)
    }
    v136 := v4[2147483648]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v136:", v136)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v74:", v74)
    _ = v1
    _ = v136
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v73
    _ = v74
}
