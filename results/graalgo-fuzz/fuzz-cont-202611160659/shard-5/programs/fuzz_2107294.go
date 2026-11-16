package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int }
type GenS2 struct { F0 int64; F1 uint32 }
type GenS3 struct { F0 uint32; F1 int16 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 bool; F1 int32; F2 int16 }
type GenS7 struct { F0 int16; F1 string }
type GenS8 struct { F0 uint32 }
type GenS9 struct { F0 uint16; F1 uint32 }
type GenS10 struct { F0 int32 }
type GenS11 struct { F0 int64 }
type GenS12 struct { F0 uint; F1 bool; F2 uint8 }
type GenS13 struct { F0 int; F1 uint16; F2 int64 }
type GenS14 struct { F0 int64 }
func genFn0() int64 {
    v5 := GenS1{F0: 100, F1: -1}
    var v6 int64 = 0
    var v7 int = 2147483647
    _ = v5
    _ = v6
    _ = v7
    return v6
}
func genFn1(p0 uint64) int64 {
    p0 = 1
    var v8 bool = (813 > 699)
    v8 = (1.5 <= 1e10)
    _ = p0
    _ = v8
    return 260
}

func main() {
    v1 := make(map[string]string)
    var v2 uint = 42
    var v3 uint16 = 998
    var v4 uint32 = 427
    v4 = 1
    if (4294967295 != 0) {
        if ("hello" != ((("foo bar" + "foo bar") + ("foo bar" + ("" + "a"))) + ("a" + "hello"))) {
            switch "foo bar" {
            case "hello":
                var v9 float64 = ((0.0 - 3.14159) / 1.0)
                _ = v9
            case "foo bar":
                var v10 float32 = ((-2.25 + ((1e10 + 0.0) - 3.14159)) - 3.14159)
                _ = v10
            default:
                v1[""] = ""
            }
            v4 += v4
        } else {
            v11 := GenS2{F0: 1, F1: v4}
            for fi12 := 0; fi12 < 10; fi12++ {
                if (((("\n\t" + "hello") != "hello") && (-1 == -10)) || false) {
                    v2 = 18446744073709551615
                    v2 += (v2 + v2)
                    var v13 int8 = 1
                    _ = v13
                }
                v4 = v4
                var v14 int16 = 42
                _ = fi12
                _ = v14
            }
            var v15 bool = false
            _ = v11
            _ = v15
        }
        var v16 int64 = 1
        _ = v16
    }
    switch 602 {
    case 127:
        v17 := GenS1{F0: -128, F1: 32767}
        _ = v17
    case -951:
        v1["foo bar"] = "foo bar"
    default:
        if (-1 == -1) {
            switch "a" {
            case "hello":
                fmt.Println(v4)
            default:
                var v18 uint32 = (v4 << (((^(v2))) & 63))
                _ = v18
            }
            var v19 int64 = 42
            _ = v19
        } else {
            var v20 uint64 = 626
            switch "hello" {
            case "":
                if false {
                    var v21 int8 = -75
                    _ = v21
                } else {
                    v1[("" + (((("foo bar" + ("foo bar" + "hello")) + "\n\t") + "foo bar") + "hello"))] = ""
                    switch -194 {
                    case -191:
                        var v22 uint8 = 1
                        _ = v22
                    case 89:
                        var v23 float64 = float64(0.0)
                        _ = v23
                    case 546:
                        v3 -= (v3 / ((164) | 1))
                    default:
                        v3 = v3
                    }
                }
            case "a":
                var v24 int16 = 600
                _ = v24
            case "\n\t":
                if true {
                    v25 := GenS2{F0: int64(v3), F1: v4}
                    _ = v25
                } else {
                    for fi26 := 0; fi26 < 1; fi26++ {
                        if true {
                            v4 = 2147483648
                            var v27 int = 95
                            _ = v27
                        } else {
                            v2 = v2
                            v2 += v2
                            if (false && (164 > 1)) {
                                switch "foo bar" {
                                case "hello":
                                    if ((!true) || (!("a" > "a"))) {
                                        v20 = v20
                                        v20 = (v20 & (v20 | v20))
                                    }
                                }
                                var v28 uint64 = 0
                                _ = v28
                            } else {
                                if (1.5 == 1.5) {
                                    v29 := GenS1{F0: 103, F1: (fi26 + 1)}
                                    v4 = uint32(v3)
                                    switch "" {
                                    case "":
                                        for fi30 := 0; fi30 < 13; fi30++ {
                                            v31 := GenS1{F0: 100, F1: (fi26 - 2147483647)}
                                            fmt.Println(v3, v4, v20)
                                            v3 = v3
                                            v32 := GenS3{F0: v4, F1: 0}
                                            _ = fi30
                                            _ = v31
                                            _ = v32
                                        }
                                    case "hello":
                                        if true {
                                            fmt.Println(v4)
                                        }
                                    default:
                                        v33 := []int{}
                                        _ = v33
                                    }
                                    _ = v29
                                } else {
                                    switch ("\n\t" + (("" + "hello") + ("\n\t" + "foo bar"))) {
                                    case "foo bar":
                                        for fi34 := 0; fi34 < 9; fi34++ {
                                            var v35 uint32 = (v4 % ((2) | 1))
                                            if true {
                                                v36 := GenS3{F0: (v4 - (v35 & (v4 & 4294967295))), F1: 0}
                                                _ = v36
                                            }
                                            _ = fi34
                                            _ = v35
                                        }
                                    case "":
                                        for fi37 := 0; fi37 < 11; fi37++ {
                                            if (float32(3.14159) <= 1.5) {
                                                v38 := GenS1{F0: 100, F1: 32767}
                                                var v39 int32 = -958
                                                _ = v38
                                                _ = v39
                                            } else {
                                                var v40 float32 = 1.5
                                                v3 -= v3
                                                if (0 != -1) {
                                                    if (("\n\t" + "") <= ("\n\t" + "a")) {
                                                        if (true && true) {
                                                            var v41 int = (fi26 & 2147483647)
                                                            var v42 int32 = 42
                                                            delete(v1, (("a" + "hello") + ("a" + "a")))
                                                            switch (v41 ^ (-(fi26))) {
                                                            case 631:
                                                                var v43 float32 = (v40 + 0.1)
                                                                _ = v43
                                                            case 4294967295:
                                                                var v44 uint32 = 4294967295
                                                                _ = v44
                                                            }
                                                            _ = v41
                                                            _ = v42
                                                        } else {
                                                            v45 := []int{243, (fi26 % (((fi37 & (fi26 >> (((^(v2))) & 63)))) | 1)), fi37, fi26}
                                                            var v46 int32 = -1
                                                            if false {
                                                                var v47 uint16 = v3
                                                                _ = v47
                                                            } else {
                                                                v48 := GenS3{F0: (^(v4)), F1: 839}
                                                                if len(v45) > 0 {
                                                                    v45[0] = (fi26 >> ((172) & 63))
                                                                }
                                                                _ = v48
                                                            }
                                                            _ = v45
                                                            _ = v46
                                                        }
                                                    }
                                                    v49 := GenS2{F0: -299, F1: 2147483648}
                                                    _ = v49
                                                } else {
                                                    v40 -= 1.5
                                                    v50 := GenS4{F0: "\n\t"}
                                                    if true {
                                                        v3 -= 558
                                                    } else {
                                                        var v51 float64 = float64((v40 + v40))
                                                        _ = v51
                                                    }
                                                    var v52 float64 = (float64((v40 / (v40 / (v40 + v40)))) / 1e10)
                                                    _ = v50
                                                    _ = v52
                                                }
                                                _ = v40
                                            }
                                            _ = fi37
                                        }
                                    default:
                                        v53 := GenS2{F0: -79, F1: 2}
                                        _ = v53
                                    }
                                    fmt.Println(v3, fi26)
                                    var v54 uint8 = 1
                                    _ = v54
                                }
                                _ = len(v1)
                                if false {
                                    for fi55 := 0; fi55 < 7; fi55++ {
                                        var v56 int = 32768
                                        var v57 int8 = -1
                                        switch ("hello" + "foo bar") {
                                        case "":
                                            v58 := GenS2{F0: 9223372036854775807, F1: v4}
                                            _ = v58
                                        }
                                        _ = fi55
                                        _ = v56
                                        _ = v57
                                    }
                                    fmt.Println(v2)
                                    var v59 int32 = 0
                                    _ = v59
                                }
                            }
                        }
                        _ = fi26
                    }
                }
            }
            _ = v20
        }
    }
    if false {
        v60 := GenS5{F0: -1}
        v3 = v3
        v3 -= (^(v3))
        _ = v60
    } else {
        v61 := []int{}
        if ((v2 - (v2 / ((v2) | 1))) <= v2) {
            for fi62 := 0; fi62 < 1; fi62++ {
                var v63 bool = false
                _ = fi62
                _ = v63
            }
            v64 := GenS2{F0: genFn1(0), F1: v4}
            _ = v64
        } else {
            if (((1 == 695) && (true || false)) && true) {
                if (!true) {
                    if (4294967295 == -2147483648) {
                        switch "foo bar" {
                        case "\n\t":
                            v1["a"] = "a"
                        case "hello":
                            var v65 uint16 = uint16(v4)
                            _ = v65
                        }
                    } else {
                        switch 1 {
                        case 32767:
                            if true {
                                if true {
                                    switch "a" {
                                    case "a":
                                        var v66 int16 = -995
                                        _ = v66
                                    default:
                                        v2 += 867
                                    }
                                    for fi67 := 0; fi67 < 17; fi67++ {
                                        for fi68 := 0; fi68 < 1; fi68++ {
                                            v69 := GenS2{F0: 786, F1: (^(v4))}
                                            _ = fi68
                                            _ = v69
                                        }
                                        switch fi67 {
                                        case 42:
                                            v70 := GenS6{F0: (!(!false)), F1: -2147483648, F2: -835}
                                            _ = v70
                                        case 2147483647:
                                            for fi71 := 0; fi71 < 18; fi71++ {
                                                v72 := GenS7{F0: -933, F1: "a"}
                                                v73 := GenS5{F0: 127}
                                                _ = fi71
                                                _ = v72
                                                _ = v73
                                            }
                                        }
                                        _ = fi67
                                    }
                                }
                            } else {
                                var v74 int32 = 1
                                v2 = v2
                                _ = v74
                            }
                        }
                        var v75 float64 = -2.25
                        _ = v75
                    }
                    v2 -= 18446744073709551615
                }
                switch "\n\t" {
                case "":
                    delete(v1, ("\n\t" + ("hello" + "foo bar")))
                case "hello":
                    v61 = append(v61, int(v3))
                case "\n\t":
                    var v76 int16 = 32767
                    _ = v76
                default:
                    fmt.Println(v3)
                }
            } else {
                for fi77 := 0; fi77 < 13; fi77++ {
                    v4 = 2
                    var v78 uint8 = 253
                    for fi79 := 0; fi79 < 19; fi79++ {
                        if (v3 >= 328) {
                            if (true && false) {
                                if ((!false) && false) {
                                    var v80 int32 = 42
                                    var v81 string = "hello"
                                    v82 := GenS2{F0: genFn0(), F1: (v4 >> (((^(v2))) & 63))}
                                    _ = v80
                                    _ = v81
                                    _ = v82
                                }
                                v78 = 0
                                switch "" {
                                case "":
                                    v83 := GenS8{F0: (^(v4))}
                                    _ = v83
                                case "hello":
                                    if (1 == 370) {
                                        var v84 int16 = int16(v2)
                                        v78 -= (^(v78))
                                        _ = v84
                                    }
                                case "\n\t":
                                    v85 := GenS4{F0: "a"}
                                    _ = v85
                                }
                            } else {
                                v86 := GenS8{F0: (v4 ^ (v4 * (v4 % ((v4) | 1))))}
                                _ = v86
                            }
                        }
                        if false {
                            if false {
                                switch (-(fi77)) {
                                case 1:
                                    if false {
                                        if ((false && true) && (!(1 > 32767))) {
                                            if (-244 != -9223372036854775808) {
                                                if false {
                                                    if false {
                                                        fmt.Println(v78, fi79)
                                                        v2 = 922
                                                        v3 -= (v3 >> (((v2 + v2)) & 63))
                                                        var v87 uint8 = (^(v78))
                                                        v88 := GenS9{F0: (v3 / (((^(v3))) | 1)), F1: v4}
                                                        v2 = (^(v2))
                                                        v89 := GenS3{F0: v4, F1: 0}
                                                        _ = v87
                                                        _ = v88
                                                        _ = v89
                                                    } else {
                                                        var v90 bool = false
                                                        if len(v61) > 0 {
                                                            v61[0] = fi77
                                                        }
                                                        fmt.Println(v90, fi77, v3)
                                                        _ = v90
                                                    }
                                                } else {
                                                    var v91 int16 = 42
                                                    if (!false) {
                                                        v92 := GenS10{F0: 0}
                                                        _ = v92
                                                    } else {
                                                        switch "foo bar" {
                                                        case "\n\t":
                                                            var v93 uint32 = v4
                                                            _ = v93
                                                        case "hello":
                                                            v78 = v78
                                                        default:
                                                            if (506 >= (v4 + 42)) {
                                                                var v94 int16 = (v91 ^ (v91 << ((v2) & 63)))
                                                                delete(v1, "")
                                                                delete(v1, "foo bar")
                                                                var v95 uint8 = (v78 ^ v78)
                                                                _ = v94
                                                                _ = v95
                                                            }
                                                            if (false || true) {
                                                                var v96 int64 = 15
                                                                var v97 uint64 = 197
                                                                v91 += -32768
                                                                _ = v96
                                                                _ = v97
                                                            }
                                                            var v98 uint32 = v4
                                                            _ = v98
                                                        }
                                                    }
                                                    _ = v91
                                                }
                                                fmt.Println(v2)
                                                v4 += v4
                                                v2 = v2
                                            } else {
                                                if false {
                                                    v2 -= (v2 / ((1) | 1))
                                                }
                                                var v99 int64 = 0
                                                _ = v99
                                            }
                                            v1[("\n\t" + (("hello" + "\n\t") + "\n\t"))] = "foo bar"
                                        } else {
                                            switch "a" {
                                            case "hello":
                                                if len(v61) > 0 {
                                                    v61[0] = -9223372036854775808
                                                }
                                            default:
                                                var v100 int16 = 0
                                                _ = v100
                                            }
                                            _ = len(v1)
                                        }
                                        v3 = uint16(v2)
                                    }
                                }
                                var v101 string = ("" + "a")
                                if true {
                                    if (100 < -1) {
                                        var v102 uint8 = 84
                                        _ = v102
                                    } else {
                                        if false {
                                            if (!true) {
                                                if false {
                                                    v103 := v1[v101]
                                                    v101 = v101
                                                    v104 := GenS11{F0: -756}
                                                    var v105 string = v101
                                                    _ = v103
                                                    _ = v104
                                                    _ = v105
                                                } else {
                                                    if (76 < 1) {
                                                        delete(v1, v101)
                                                    }
                                                    v4 += uint32(v3)
                                                }
                                                switch (v101 + "") {
                                                case "\n\t":
                                                    v61 = append(v61, (fi79 - 290))
                                                case "foo bar":
                                                    v106 := GenS3{F0: v4, F1: 344}
                                                    _ = v106
                                                case "hello":
                                                    var v107 string = v101
                                                    _ = v107
                                                default:
                                                    if len(v61) > 0 {
                                                        v61[0] = (fi79 & 4294967295)
                                                    }
                                                }
                                                var v108 string = ((v101 + "a") + ("foo bar" + v101))
                                                fmt.Println(fi77, v101)
                                                v4 = v4
                                                v109 := GenS2{F0: int64(v3), F1: v4}
                                                if true {
                                                    if len(v61) > 0 {
                                                        v61[0] = fi79
                                                    }
                                                    _ = len(v61)
                                                    switch v108 {
                                                    case "":
                                                        v78 = 2
                                                    case "hello":
                                                        if true {
                                                            var v110 int16 = -1
                                                            v111 := GenS12{F0: v2, F1: (3.14159 != (float64(float32((0.1 * 3.14159))) - -2.25)), F2: 255}
                                                            _ = v110
                                                            _ = v111
                                                        }
                                                    case "foo bar":
                                                        fmt.Println(v108, fi77)
                                                    default:
                                                        var v112 float32 = 1e10
                                                        _ = v112
                                                    }
                                                    switch v108 {
                                                    case "hello":
                                                        v108 = (("\n\t" + (("a" + "a") + (v101 + ""))) + v101)
                                                    default:
                                                        v2 = 275
                                                    }
                                                    if (!true) {
                                                        v1[v108] = v108
                                                        if len(v61) >= 2 {
                                                            v61 = v61[1:len(v61)]
                                                        }
                                                    } else {
                                                        v101 = (((v101 + v108) + v101) + "")
                                                    }
                                                } else {
                                                    switch (fi79 & (^(fi79))) {
                                                    case 0:
                                                        if (!(fi79 == (fi77 ^ fi77))) {
                                                            v113 := GenS10{F0: 1}
                                                            switch (((v108 + "hello") + (v108 + "")) + ("a" + (v101 + v101))) {
                                                            case "foo bar":
                                                                var v114 string = (v108 + "foo bar")
                                                                _ = v114
                                                            }
                                                            _ = v113
                                                        }
                                                    case 42:
                                                        switch fi77 {
                                                        case 0:
                                                            fmt.Println(v2, v108, v3)
                                                        case 9223372036854775807:
                                                            var v115 float32 = 1.5
                                                            _ = v115
                                                        case 1:
                                                            if true {
                                                                var v116 string = ""
                                                                if len(v61) >= 2 {
                                                                    v61 = v61[1:len(v61)]
                                                                }
                                                                if false {
                                                                    v2 -= 148
                                                                    if false {
                                                                        var v117 int16 = 32767
                                                                        _ = v117
                                                                    }
                                                                    switch (v108 + v108) {
                                                                    case "hello":
                                                                        var v118 uint = v2
                                                                        _ = v118
                                                                    case "a":
                                                                        var v119 uint64 = 0
                                                                        _ = v119
                                                                    case "\n\t":
                                                                        var v120 int = fi77
                                                                        _ = v120
                                                                    default:
                                                                        v78 = v78
                                                                    }
                                                                } else {
                                                                    var v121 int16 = 1
                                                                    _ = v121
                                                                }
                                                                _ = v116
                                                            } else {
                                                                switch fi79 {
                                                                case 32767:
                                                                    v108 = v108
                                                                case -965:
                                                                    var v122 uint64 = 2
                                                                    var v123 int16 = -32768
                                                                    var v124 uint64 = v122
                                                                    if true {
                                                                        v78 = uint8(v3)
                                                                        var v125 int16 = v123
                                                                        v125 = (^(v123))
                                                                        fmt.Println(v123, v125)
                                                                        var v126 uint = v2
                                                                        v108 = (v108 + v101)
                                                                        _ = v125
                                                                        _ = v126
                                                                    }
                                                                    _ = v122
                                                                    _ = v123
                                                                    _ = v124
                                                                }
                                                                v101 = (v101 + v101)
                                                                v78 = v78
                                                            }
                                                        }
                                                    }
                                                    v127 := v1[v108]
                                                    _ = v127
                                                }
                                                if len(v61) >= 2 {
                                                    v61 = v61[1:len(v61)]
                                                }
                                                _ = v108
                                                _ = v109
                                            }
                                            var v128 string = (v101 + "a")
                                            _ = v128
                                        } else {
                                            switch fi79 {
                                            case 792:
                                                var v129 float32 = 0.1
                                                _ = v129
                                            case 26:
                                                if len(v61) > 0 {
                                                    v61[0] = fi79
                                                }
                                            }
                                            if (true && false) {
                                                v101 = v101
                                                switch fi79 {
                                                case -2147483648:
                                                    var v130 uint = (v2 - v2)
                                                    _ = v130
                                                case -932:
                                                    _ = len(v61)
                                                case 127:
                                                    var v131 uint64 = 0
                                                    fmt.Println(v131, v2, v4)
                                                    var v132 int32 = 2147483647
                                                    var v133 int = (-(fi77))
                                                    _ = v131
                                                    _ = v132
                                                    _ = v133
                                                default:
                                                    switch (fi79 ^ 128) {
                                                    case 0:
                                                        var v134 int = (^(fi79))
                                                        _ = v134
                                                    default:
                                                        if true {
                                                            fmt.Println(v4, v3, v101)
                                                        } else {
                                                            var v135 uint = v2
                                                            v1[v101] = (v101 + v101)
                                                            if (true && false) {
                                                                v1[v101] = "\n\t"
                                                            }
                                                            _ = v135
                                                        }
                                                    }
                                                }
                                            } else {
                                                var v136 int = fi77
                                                _ = v136
                                            }
                                        }
                                    }
                                    var v137 int = (fi77 % ((0) | 1))
                                    _ = len(v61)
                                    _ = v137
                                }
                                _ = v101
                            } else {
                                var v138 uint = (v2 + 375)
                                _ = v138
                            }
                            if (true || (true && true)) {
                                fmt.Println(fi79, v78)
                                if len(v61) > 0 {
                                    v61[0] = fi79
                                }
                            }
                        } else {
                            v139 := GenS13{F0: (fi77 + (fi77 & -768)), F1: (v3 - v3), F2: int64(v4)}
                            _ = v139
                        }
                        v61 = append(v61, fi79)
                        v78 = (v78 * v78)
                        var v140 int32 = 1
                        if len(v61) > 0 {
                            v61[0] = 128
                        }
                        v140 += v140
                        fmt.Println(v78)
                        var v141 float32 = (-2.25 / 1.5)
                        v140 -= (^(v140))
                        _ = fi79
                        _ = v140
                        _ = v141
                    }
                    _ = fi77
                    _ = v78
                }
                for fi142 := 0; fi142 < 17; fi142++ {
                    for fi143 := 0; fi143 < 19; fi143++ {
                        if (false && true) {
                            v3 -= 255
                            if (!false) {
                                v144 := GenS2{F0: int64(v4), F1: (v4 ^ 2)}
                                v145 := GenS4{F0: ("a" + "")}
                                _ = v144
                                _ = v145
                            } else {
                                fmt.Println(v4, v3)
                            }
                        } else {
                            v2 = (v2 % ((v2) | 1))
                            var v146 int32 = int32(fi142)
                            v147 := GenS14{F0: -1}
                            _ = len(v1)
                            _ = v146
                            _ = v147
                        }
                        v148 := GenS12{F0: v2, F1: true, F2: 255}
                        _ = fi143
                        _ = v148
                    }
                    _ = fi142
                }
                _ = len(v61)
            }
            if len(v61) >= 2 {
                v61 = v61[1:len(v61)]
            }
            var v149 uint8 = 2
            _ = v149
        }
        _ = v61
    }
    v150 := []int{-1, 9223372036854775807, 42, 845}
    v3 = v3
    fmt.Println(v4, v4)
    v2 -= (v2 / ((v2) | 1))
    var v151 uint8 = 1
    _ = len(v1)
    v4 = 4294967295
    var v152 uint8 = (^(v151))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v150:", len(v150), v150)
    fmt.Println("v151:", v151)
    fmt.Println("v152:", v152)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v150
    _ = v151
    _ = v152
    _ = v2
    _ = v3
    _ = v4
}
