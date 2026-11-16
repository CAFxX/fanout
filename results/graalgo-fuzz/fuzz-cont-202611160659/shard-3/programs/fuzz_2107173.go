package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint; F1 uint16; F2 int16 }
type GenS3 struct { F0 uint8; F1 uint64 }
type GenS4 struct { F0 uint32; F1 string; F2 int }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 uint; F1 uint32 }
type GenS7 struct { F0 uint; F1 int8 }
type GenS8 struct { F0 int64; F1 uint32; F2 int64 }
type GenS9 struct { F0 uint16 }
type GenS10 struct { F0 int8; F1 uint16 }
type GenS11 struct { F0 int8; F1 uint8; F2 int32 }
type GenS12 struct { F0 string }
type GenS13 struct { F0 int8; F1 bool }
func genFn0(p0 int8) int32 {
    fmt.Println(p0)
    v5 := GenS1{F0: 2}
    _ = p0
    _ = v5
    return 0
}
func genFn1(p0 string, p1 uint64) int32 {
    var v6 uint8 = 42
    for fi7 := 0; fi7 < 4; fi7++ {
        v8 := []int{}
        _ = fi7
        _ = v8
    }
    var v9 uint16 = 65535
    for fi10 := 0; fi10 < 5; fi10++ {
        if false {
            v11 := []int{}
            switch fi10 {
            case 127:
                var v12 uint8 = (v6 << ((0) & 63))
                _ = v12
            case 32768:
                if len(v11) >= 2 {
                    v11 = v11[1:len(v11)]
                }
            case 128:
                for fi13 := 0; fi13 < 20; fi13++ {
                    var v14 uint8 = (v6 % (((v6 << ((0) & 63))) | 1))
                    var v15 int16 = 755
                    _ = fi13
                    _ = v14
                    _ = v15
                }
            default:
                switch fi10 {
                case -656:
                    fmt.Println(v9, v9, v6)
                case 127:
                    switch "\n\t" {
                    case "a":
                        v11 = append(v11, 35)
                    case "hello":
                        p1 = p1
                    case "foo bar":
                        var v16 int16 = -32768
                        _ = v16
                    }
                }
            }
            _ = v11
        }
        _ = fi10
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v9
    return int32(v9)
}

func main() {
    var v1 uint64 = 216
    v2 := make(map[int]string)
    v3 := []string{"hello"}
    var v4 uint16 = uint16(v1)
    for fi17 := 0; fi17 < 4; fi17++ {
        if (!false) {
            var v18 float32 = ((0.1 / -2.25) + (((((3.14159 + 3.14159) + 1.5) + float32(3.14159)) + 3.14159) / 3.14159))
            _ = v18
        }
        _ = fi17
    }
    for fi19 := 0; fi19 < 11; fi19++ {
        for fi20 := 0; fi20 < 1; fi20++ {
            delete(v2, 2147483647)
            v21 := GenS2{F0: 1, F1: 2, F2: -1}
            _ = fi20
            _ = v21
        }
        _ = fi19
    }
    fmt.Println(v4, v4)
    if (!false) {
        var v22 string = (("a" + "hello") + "\n\t")
        fmt.Println(v22)
        _ = v22
    }
    v23 := GenS1{F0: 128}
    var v24 int32 = -1
    if false {
        v24 = -753
        v25 := GenS1{F0: 9}
        for fi26 := 0; fi26 < 16; fi26++ {
            fmt.Println(fi26, v4, v24)
            for fi27 := 0; fi27 < 8; fi27++ {
                switch fi27 {
                case 128:
                    if (47 > -808) {
                        var v28 int32 = -2147483648
                        var v29 float64 = (0.0 * (((3.14159 * 1e10) - (float64(-2.25) - (((float64(0.1) * 1e10) * 1e10) - 1e10))) / 1e10))
                        v24 = v28
                        _ = v28
                        _ = v29
                    }
                case 9223372036854775807:
                    if (false && false) {
                        var v30 float32 = 1.5
                        v31 := GenS2{F0: 0, F1: 0, F2: int16(v1)}
                        switch "hello" {
                        case "":
                            if (v24 != -556) {
                                v24 = genFn1("hello", 2)
                                var v32 int32 = v24
                                var v33 string = "foo bar"
                                _ = v32
                                _ = v33
                            }
                        case "a":
                            delete(v2, int(v1))
                        case "hello":
                            v4 = v4
                        }
                        v4 = v4
                        _ = v30
                        _ = v31
                    }
                case 2147483647:
                    var v34 bool = (!true)
                    _ = v34
                default:
                    var v35 uint64 = v1
                    _ = v35
                }
                _ = fi27
            }
            _ = fi26
        }
        _ = v25
    } else {
        v24 -= (v24 + (v24 | v24))
        if true {
            switch "a" {
            case "foo bar":
                for fi36 := 0; fi36 < 16; fi36++ {
                    v37 := GenS3{F0: 42, F1: v1}
                    _ = fi36
                    _ = v37
                }
            case "":
                var v38 uint32 = 0
                _ = v38
            case "\n\t":
                switch 788 {
                case 32767:
                    for fi39 := 0; fi39 < 19; fi39++ {
                        if false {
                            var v40 uint16 = v4
                            _ = v40
                        } else {
                            var v41 int = (fi39 & (fi39 * fi39))
                            for fi42 := 0; fi42 < 6; fi42++ {
                                var v43 float32 = 3.14159
                                if (true || true) {
                                    fmt.Println(v4, v41, v43)
                                    v44 := GenS1{F0: 1}
                                    if (!(0 != 32767)) {
                                        if len(v3) > 0 {
                                            v3[0] = ("\n\t" + "hello")
                                        }
                                        v45 := GenS2{F0: 352, F1: v4, F2: 0}
                                        var v46 uint32 = 42
                                        _ = v45
                                        _ = v46
                                    } else {
                                        if true {
                                            v41 += v41
                                        } else {
                                            v3 = append(v3, ((("hello" + ("a" + (("foo bar" + ("" + "hello")) + "hello"))) + "hello") + ("\n\t" + "\n\t")))
                                            fmt.Println(v41, fi39, v4)
                                        }
                                        v47 := GenS2{F0: 0, F1: (v4 % (((v4 % ((v4) | 1))) | 1)), F2: -659}
                                        switch int(v4) {
                                        case 1:
                                            v48 := GenS2{F0: 1, F1: v4, F2: -631}
                                            _ = v48
                                        case -541:
                                            switch 392 {
                                            case 128:
                                                var v49 float64 = float64(v43)
                                                _ = v49
                                            default:
                                                if false {
                                                    fmt.Println(fi39, fi42)
                                                    fmt.Println(v24, v24, v4)
                                                    v1 += (v1 & (v1 & v1))
                                                }
                                            }
                                        case 0:
                                            switch "\n\t" {
                                            case "a":
                                                if (true && false) {
                                                    var v50 int = (fi39 - fi39)
                                                    var v51 float64 = 0.1
                                                    v43 = ((v43 / v43) - 1e10)
                                                    if (v1 >= 2) {
                                                        var v52 bool = false
                                                        switch "\n\t" {
                                                        case "\n\t":
                                                            v53 := GenS4{F0: 2147483648, F1: "hello", F2: (fi42 + (fi42 * fi42))}
                                                            _ = v53
                                                        }
                                                        var v54 uint32 = 42
                                                        _ = v52
                                                        _ = v54
                                                    }
                                                    var v55 uint16 = (v4 | (v4 + (v4 % ((v4) | 1))))
                                                    _ = v50
                                                    _ = v51
                                                    _ = v55
                                                } else {
                                                    var v56 int = 128
                                                    _ = v56
                                                }
                                            default:
                                                var v57 string = "foo bar"
                                                _ = v57
                                            }
                                        default:
                                            switch v41 {
                                            case 128:
                                                fmt.Println(v4)
                                            }
                                            var v58 bool = (!true)
                                            _ = len(v3)
                                            _ = v58
                                        }
                                        _ = v47
                                    }
                                    if false {
                                        if false {
                                            v59 := GenS1{F0: 2}
                                            _ = v59
                                        } else {
                                            switch (("" + ("" + "\n\t")) + "\n\t") {
                                            case "a":
                                                var v60 uint32 = 65535
                                                _ = v60
                                            case "hello":
                                                if (42 < 42) {
                                                    v61 := GenS4{F0: 2, F1: "", F2: fi39}
                                                    _ = v61
                                                } else {
                                                    v62 := GenS5{F0: 924}
                                                    _ = v62
                                                }
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                                v63 := GenS6{F0: 1, F1: 2}
                                                _ = v63
                                            }
                                        }
                                    } else {
                                        _ = len(v2)
                                        switch ("" + "foo bar") {
                                        case "\n\t":
                                            v41 = (fi42 % ((v41) | 1))
                                        case "foo bar":
                                            var v64 int32 = genFn1(("\n\t" + "foo bar"), v1)
                                            _ = v64
                                        }
                                        if false {
                                            fmt.Println(v41, fi42, v4)
                                            v4 += v4
                                            v65 := GenS7{F0: 65535, F1: -128}
                                            _ = v65
                                        }
                                        v66 := GenS3{F0: 2, F1: v1}
                                        _ = v66
                                    }
                                    v67 := GenS1{F0: 2}
                                    var v68 int8 = -1
                                    _ = v44
                                    _ = v67
                                    _ = v68
                                }
                                _ = fi42
                                _ = v43
                            }
                            for fi69 := 0; fi69 < 9; fi69++ {
                                switch (fi39 >> ((136) & 63)) {
                                case -794:
                                    v24 -= 172
                                case -468:
                                    v70 := GenS2{F0: 0, F1: 255, F2: 980}
                                    _ = v70
                                case 823:
                                    switch (-(fi39)) {
                                    case -382:
                                        v71 := GenS1{F0: 2}
                                        _ = v71
                                    case -2147483648:
                                        if ((fi69 / (((v41 << ((603) & 63))) | 1)) == -259) {
                                            v72 := GenS1{F0: 141}
                                            var v73 float32 = 3.14159
                                            var v74 int64 = 9223372036854775807
                                            _ = v72
                                            _ = v73
                                            _ = v74
                                        } else {
                                            switch ("\n\t" + "foo bar") {
                                            case "\n\t":
                                                v75 := v2[v41]
                                                _ = v75
                                            case "":
                                                v4 = 42
                                            case "a":
                                                var v76 int16 = 1000
                                                _ = v76
                                            default:
                                                if true {
                                                    v77 := GenS5{F0: 65535}
                                                    var v78 int64 = 80
                                                    switch (("hello" + "foo bar") + "a") {
                                                    case "":
                                                        var v79 uint64 = 18446744073709551615
                                                        switch "hello" {
                                                        case "":
                                                            var v80 int64 = (v78 / (((v78 / ((-1) | 1))) | 1))
                                                            _ = v80
                                                        case "foo bar":
                                                            fmt.Println(v24, v4, v41)
                                                        case "a":
                                                            var v81 int8 = -128
                                                            _ = v81
                                                        }
                                                        v3 = append(v3, "hello")
                                                        _ = v79
                                                    default:
                                                        v41 = fi69
                                                    }
                                                    if true {
                                                        if (0 >= 127) {
                                                            if true {
                                                                var v82 int32 = (v24 - v24)
                                                                v83 := GenS3{F0: 42, F1: (v1 | (^(v1)))}
                                                                _ = len(v3)
                                                                _ = v82
                                                                _ = v83
                                                            }
                                                            if ((true || (true && (!true))) || (1 == 0)) {
                                                                v1 -= (v1 >> ((528) & 63))
                                                            } else {
                                                                _ = len(v2)
                                                                v78 = (v78 ^ v78)
                                                                fmt.Println(v78, fi69, v78)
                                                            }
                                                        }
                                                    } else {
                                                        v4 = (^(v4))
                                                        var v84 int16 = -997
                                                        var v85 int = fi69
                                                        _ = v84
                                                        _ = v85
                                                    }
                                                    var v86 uint8 = 128
                                                    _ = v77
                                                    _ = v78
                                                    _ = v86
                                                } else {
                                                    if false {
                                                        var v87 int = (fi39 - fi69)
                                                        var v88 uint = 1
                                                        v89 := GenS7{F0: v88, F1: 100}
                                                        _ = v87
                                                        _ = v88
                                                        _ = v89
                                                    }
                                                }
                                            }
                                            if true {
                                                v41 = fi69
                                                v1 = (v1 << ((42) & 63))
                                                v90 := GenS8{F0: 258, F1: 42, F2: 616}
                                                _ = v90
                                            }
                                            var v91 int16 = 0
                                            _ = v91
                                        }
                                    case -874:
                                        v24 -= genFn1("hello", 223)
                                    default:
                                        if true {
                                            if true {
                                                v4 = (v4 / ((v4) | 1))
                                                v3 = append(v3, (("a" + ((("hello" + "") + "hello") + ("\n\t" + "hello"))) + "hello"))
                                            } else {
                                                fmt.Println(fi69, v24, v1)
                                                if false {
                                                    var v92 uint = 65535
                                                    v41 = (fi69 % ((-585) | 1))
                                                    _ = v92
                                                }
                                                v41 = 586
                                            }
                                            var v93 bool = true
                                            v94 := GenS9{F0: v4}
                                            _ = v93
                                            _ = v94
                                        }
                                    }
                                default:
                                    v95 := v2[fi69]
                                    _ = v95
                                }
                                if (2 >= 1) {
                                    switch "foo bar" {
                                    case "foo bar":
                                        v96 := GenS5{F0: 42}
                                        _ = v96
                                    case "hello":
                                        v24 = v24
                                    case "":
                                        var v97 uint = 255
                                        _ = v97
                                    }
                                } else {
                                    fmt.Println(v41, fi39)
                                    if false {
                                        v41 += fi39
                                        var v98 uint8 = 219
                                        v24 += -1
                                        _ = v98
                                    }
                                    if (((-2.25 + (1e10 - 1.5)) * ((-2.25 - 3.14159) - float32(0.0))) != -2.25) {
                                        var v99 uint64 = 282
                                        switch ("a" + "\n\t") {
                                        case "\n\t":
                                            fmt.Println(v4, v99)
                                        case "foo bar":
                                            v2[(fi39 / (((v41 ^ (^(v41)))) | 1))] = ""
                                        case "hello":
                                            var v100 int16 = 1000
                                            _ = v100
                                        default:
                                            var v101 uint = 255
                                            _ = v101
                                        }
                                        _ = v99
                                    }
                                }
                                if false {
                                    v102 := GenS10{F0: 1, F1: (v4 & (^(v4)))}
                                    switch 2147483647 {
                                    case 1:
                                        delete(v2, fi39)
                                    case 127:
                                        fmt.Println(fi69, fi39)
                                    case -1:
                                        switch (("" + "a") + "hello") {
                                        case "foo bar":
                                            v103 := GenS10{F0: -46, F1: v4}
                                            _ = v103
                                        case "":
                                            fmt.Println(fi39)
                                            v104 := GenS2{F0: 2, F1: 65535, F2: 0}
                                            _ = v104
                                        }
                                    default:
                                        if (1 > 676) {
                                            var v105 uint64 = v1
                                            if false {
                                                var v106 string = "hello"
                                                if (v4 >= (^(v4))) {
                                                    fmt.Println(fi69, v1, v1)
                                                }
                                                _ = v106
                                            }
                                            _ = v105
                                        } else {
                                            v2[int(v4)] = "hello"
                                            delete(v2, 127)
                                        }
                                    }
                                    if false {
                                        fmt.Println(v1)
                                        v107 := GenS5{F0: 255}
                                        _ = v107
                                    }
                                    var v108 int16 = -488
                                    _ = v102
                                    _ = v108
                                }
                                _ = fi69
                            }
                            _ = v41
                        }
                        var v109 uint8 = 2
                        if (0 == 127) {
                            v110 := GenS8{F0: 1, F1: 1, F2: 9223372036854775807}
                            var v111 int64 = 388
                            v112 := GenS5{F0: 0}
                            _ = v110
                            _ = v111
                            _ = v112
                        }
                        _ = fi39
                        _ = v109
                    }
                case -9:
                    var v113 uint = 0
                    _ = v113
                case -9223372036854775808:
                    v4 = (v4 + (v4 & 65535))
                }
            }
            v114 := GenS7{F0: 42, F1: 0}
            _ = v114
        } else {
            for fi115 := 0; fi115 < 11; fi115++ {
                switch (fi115 << ((325) & 63)) {
                case 127:
                    switch 127 {
                    case 127:
                        _ = len(v2)
                    case 0:
                        fmt.Println(v4)
                    case -2147483648:
                        var v116 int = (^(fi115))
                        _ = v116
                    default:
                        if true {
                            var v117 uint16 = 0
                            _ = v117
                        } else {
                            var v118 float64 = 0.0
                            switch fi115 {
                            case 128:
                                v2[128] = "foo bar"
                            case 32767:
                                v118 -= (1e10 - v118)
                            default:
                                v119 := GenS2{F0: 4294967295, F1: 32767, F2: 1000}
                                _ = v119
                            }
                            var v120 uint = 282
                            _ = v118
                            _ = v120
                        }
                    }
                case 32768:
                    v1 = (v1 ^ v1)
                case 128:
                    if true {
                        var v121 int64 = 899
                        var v122 int64 = v121
                        v1 -= v1
                        _ = v121
                        _ = v122
                    } else {
                        if (74 == 0) {
                            var v123 int = (fi115 << ((2) & 63))
                            if (-9223372036854775808 > (v123 % ((v123) | 1))) {
                                if false {
                                    v4 -= 639
                                    _ = len(v2)
                                    fmt.Println(v24, v24, v4)
                                }
                                v1 = v1
                                var v124 uint8 = 0
                                _ = v124
                            }
                            _ = v123
                        } else {
                            v4 = v4
                            v125 := GenS11{F0: 127, F1: 128, F2: (v24 + (v24 % (((v24 * (v24 >> ((689) & 63)))) | 1)))}
                            _ = v125
                        }
                        for fi126 := 0; fi126 < 4; fi126++ {
                            v24 -= v24
                            var v127 bool = (v4 > v4)
                            var v128 uint64 = v1
                            _ = fi126
                            _ = v127
                            _ = v128
                        }
                        var v129 int64 = 1
                        _ = v129
                    }
                default:
                    if true {
                        if ((v1 >> ((886) & 63)) <= 42) {
                            var v130 bool = true
                            v4 += v4
                            v131 := GenS12{F0: "\n\t"}
                            _ = v130
                            _ = v131
                        }
                        var v132 string = ("foo bar" + "hello")
                        _ = len(v3)
                        _ = v132
                    } else {
                        v133 := GenS13{F0: -30, F1: false}
                        var v134 string = ("" + "hello")
                        _ = v133
                        _ = v134
                    }
                }
                _ = fi115
            }
        }
        fmt.Println(v4)
    }
    var v135 int64 = -9223372036854775808
    fmt.Println("v1:", v1)
    fmt.Println("v135:", v135)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v135
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
}
