package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64; F2 uint8 }
type GenS2 struct { F0 uint; F1 int8 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int16; F1 uint16 }
type GenS5 struct { F0 uint; F1 string }
type GenS6 struct { F0 uint; F1 int8 }
type GenS7 struct { F0 int; F1 string; F2 int32 }
type GenS8 struct { F0 int64; F1 uint32 }
type GenS9 struct { F0 int64; F1 int32; F2 int }
func genFn0() uint16 {
    if (1 <= -128) {
        var v5 int = 4294967295
        if (((true && ("hello" <= ("a" + "foo bar"))) || (!(628 > 122))) && true) {
            var v6 int8 = 42
            var v7 float32 = -2.25
            v8 := []int{v5, v5, 514, (v5 ^ v5)}
            _ = v6
            _ = v7
            _ = v8
        } else {
            v9 := GenS1{F0: 0, F1: 9223372036854775807, F2: 0}
            v10 := GenS1{F0: 1, F1: -282, F2: 2}
            v11 := GenS1{F0: 255, F1: 9223372036854775807, F2: 1}
            _ = v10
            _ = v11
            _ = v9
        }
        var v12 int8 = -128
        _ = v12
        _ = v5
    }
    if true {
        var v13 uint16 = 65535
        fmt.Println(v13)
        if true {
            var v14 int64 = 90
            var v15 string = ("" + "a")
            _ = v14
            _ = v15
        }
        _ = v13
    } else {
        var v16 bool = false
        v17 := []int{1, 0, 32768}
        for fi18 := 0; fi18 < 13; fi18++ {
            var v19 bool = v16
            if (v19 || ((!(1.5 < (3.14159 * (0.0 / (0.1 + 1.5))))) && v19)) {
                for fi20 := 0; fi20 < 19; fi20++ {
                    v19 = v16
                    _ = fi20
                }
            } else {
                fmt.Println(v16, v19, v19)
                v21 := GenS2{F0: 18446744073709551615, F1: 100}
                v19 = (((float64(-2.25) - (1.5 + (1.5 + 1e10))) <= float64(0.0)) && (-32768 == -469))
                _ = v21
            }
            _ = fi18
            _ = v19
        }
        _ = v16
        _ = v17
    }
    v22 := []int{633, 4294967295}
    for fi23 := 0; fi23 < 17; fi23++ {
        if ("hello" < ("foo bar" + "hello")) {
            switch "foo bar" {
            case "":
                for fi24 := 0; fi24 < 4; fi24++ {
                    var v25 float32 = 0.0
                    if (true || false) {
                        var v26 uint = 42
                        v26 -= 65535
                        switch "foo bar" {
                        case "a":
                            v27 := GenS2{F0: (v26 / (((v26 / ((v26) | 1))) | 1)), F1: 17}
                            _ = v27
                        default:
                            v25 = v25
                        }
                        _ = v26
                    } else {
                        var v28 uint8 = 245
                        var v29 int32 = -960
                        v30 := GenS3{F0: true}
                        v31 := GenS2{F0: 65535, F1: 127}
                        var v32 uint32 = 2147483648
                        _ = v28
                        _ = v29
                        _ = v30
                        _ = v31
                        _ = v32
                    }
                    v25 = (3.14159 * (-2.25 / ((((0.0 * v25) + (((float32(1e10) - v25) * ((v25 * (3.14159 / v25)) * v25)) * (-2.25 * 0.0))) * ((v25 + ((v25 + v25) / 3.14159)) + 0.0)) - v25)))
                    _ = fi24
                    _ = v25
                }
            case "foo bar":
                if len(v22) > 0 {
                    v22[0] = (fi23 - fi23)
                }
            }
            switch 97 {
            case 32768:
                _ = len(v22)
            case -2147483648:
                var v33 int = fi23
                _ = v33
            case 379:
                var v34 uint64 = 169
                _ = v34
            }
        } else {
            switch fi23 {
            case -692:
                var v35 uint = 65535
                _ = v35
            case 4294967295:
                var v36 int32 = 2147483647
                _ = v36
            case 127:
                if (!true) {
                    switch fi23 {
                    case 32767:
                        switch (fi23 << ((42) & 63)) {
                        case 127:
                            for fi37 := 0; fi37 < 1; fi37++ {
                                var v38 uint64 = 200
                                _ = fi37
                                _ = v38
                            }
                        }
                    case 506:
                        if len(v22) >= 2 {
                            v22 = v22[1:len(v22)]
                        }
                    case 0:
                        v39 := GenS3{F0: false}
                        _ = v39
                    }
                }
            default:
                if true {
                    var v40 uint32 = 4
                    if false {
                        var v41 float32 = 3.14159
                        fmt.Println(v40, v40, v41)
                        var v42 int16 = -32768
                        _ = v41
                        _ = v42
                    } else {
                        var v43 uint16 = 42
                        _ = v43
                    }
                    _ = v40
                } else {
                    if (((-2.25 * 3.14159) * (1.5 * 3.14159)) >= (1.5 * ((float32(0.1) - 1.5) / 1.0))) {
                        for fi44 := 0; fi44 < 18; fi44++ {
                            if false {
                                switch "a" {
                                case "\n\t":
                                    var v45 int64 = -9223372036854775808
                                    _ = v45
                                case "hello":
                                    v46 := GenS3{F0: (!false)}
                                    _ = v46
                                case "a":
                                    var v47 uint8 = 255
                                    _ = v47
                                }
                                var v48 float32 = 1e10
                                switch ("" + (("hello" + "\n\t") + ("hello" + "a"))) {
                                case "foo bar":
                                    v48 = 3.14159
                                default:
                                    if (2 >= 0) {
                                        v48 = float32((-2.25 * 3.14159))
                                        var v49 bool = (548 > 1)
                                        _ = v49
                                    }
                                }
                                var v50 uint64 = 360
                                _ = v48
                                _ = v50
                            }
                            var v51 uint32 = 687
                            var v52 int32 = 63
                            _ = len(v22)
                            var v53 bool = false
                            switch (fi44 >> ((65535) & 63)) {
                            case 128:
                                var v54 float64 = (3.14159 / 3.14159)
                                _ = v54
                            case -1:
                                fmt.Println(v51)
                            case -2147483648:
                                v22 = append(v22, fi23)
                            }
                            _ = fi44
                            _ = v51
                            _ = v52
                            _ = v53
                        }
                        if (0 > 18446744073709551615) {
                            _ = len(v22)
                            if (1 < 1) {
                                var v55 uint32 = 2
                                var v56 uint = 920
                                _ = v55
                                _ = v56
                            } else {
                                if (!true) {
                                    for fi57 := 0; fi57 < 16; fi57++ {
                                        var v58 string = ""
                                        v59 := GenS4{F0: -606, F1: 603}
                                        _ = fi57
                                        _ = v58
                                        _ = v59
                                    }
                                    var v60 string = "foo bar"
                                    _ = v60
                                }
                            }
                        }
                    } else {
                        if (!(!true)) {
                            var v61 bool = ("\n\t" == "hello")
                            fmt.Println(fi23)
                            _ = v61
                        }
                        for fi62 := 0; fi62 < 16; fi62++ {
                            if false {
                                v63 := GenS5{F0: 65535, F1: "hello"}
                                var v64 int = fi23
                                _ = v63
                                _ = v64
                            } else {
                                _ = len(v22)
                                fmt.Println(fi23, fi62)
                                if (477 > 4294967295) {
                                    var v65 string = "a"
                                    switch ((v65 + v65) + v65) {
                                    case "a":
                                        v65 = v65
                                    case "\n\t":
                                        fmt.Println(fi23, fi62, fi23)
                                    default:
                                        var v66 uint32 = 969
                                        _ = v66
                                    }
                                    _ = v65
                                } else {
                                    var v67 uint64 = 565
                                    var v68 int8 = 100
                                    _ = v67
                                    _ = v68
                                }
                                var v69 uint8 = 255
                                _ = v69
                            }
                            var v70 uint = 255
                            _ = fi62
                            _ = v70
                        }
                        if true {
                            if false {
                                if true {
                                    var v71 float64 = 0.0
                                    fmt.Println(v71, v71)
                                    _ = v71
                                } else {
                                    var v72 float32 = 1e10
                                    var v73 int64 = -9223372036854775808
                                    _ = v72
                                    _ = v73
                                }
                            }
                            switch "a" {
                            case "\n\t":
                                for fi74 := 0; fi74 < 18; fi74++ {
                                    var v75 int8 = 100
                                    if false {
                                        if (732 != 101) {
                                            v75 = v75
                                            if true {
                                                if (1 != 1) {
                                                    fmt.Println(fi74)
                                                } else {
                                                    v76 := GenS6{F0: 0, F1: -128}
                                                    var v77 int32 = -2147483648
                                                    var v78 uint32 = 2
                                                    _ = v76
                                                    _ = v77
                                                    _ = v78
                                                }
                                            }
                                            v75 = (^(v75))
                                        } else {
                                            var v79 int32 = -536
                                            if (!(!(!((-(v79)) == v79)))) {
                                                switch "hello" {
                                                case "\n\t":
                                                    v75 += 100
                                                case "":
                                                    _ = len(v22)
                                                }
                                                v75 = v75
                                            } else {
                                                if true {
                                                    switch fi23 {
                                                    case 2147483647:
                                                        var v80 int16 = 1
                                                        _ = v80
                                                    }
                                                    v75 -= -1
                                                } else {
                                                    v75 = v75
                                                    v81 := GenS7{F0: 32767, F1: ("" + "\n\t"), F2: v79}
                                                    if false {
                                                        if false {
                                                            var v82 int32 = v79
                                                            var v83 float64 = float64(((0.0 / 3.14159) * -2.25))
                                                            var v84 int8 = v75
                                                            _ = v82
                                                            _ = v83
                                                            _ = v84
                                                        } else {
                                                            var v85 int8 = (v75 + 42)
                                                            _ = v85
                                                        }
                                                        var v86 uint64 = 0
                                                        _ = v86
                                                    } else {
                                                        v87 := GenS8{F0: 9223372036854775807, F1: 639}
                                                        _ = v87
                                                    }
                                                    _ = v81
                                                }
                                                var v88 int32 = 2147483647
                                                fmt.Println(v88, fi74, fi23)
                                                _ = v88
                                            }
                                            fmt.Println(v79)
                                            fmt.Println(v79)
                                            var v89 string = "hello"
                                            var v90 int32 = (^(v79))
                                            _ = v79
                                            _ = v89
                                            _ = v90
                                        }
                                        fmt.Println(fi74, fi74, v75)
                                        v91 := GenS2{F0: 4294967295, F1: v75}
                                        _ = v91
                                    }
                                    var v92 float64 = (-2.25 * 1.5)
                                    _ = fi74
                                    _ = v75
                                    _ = v92
                                }
                            case "":
                                var v93 uint64 = 42
                                _ = v93
                            case "a":
                                v94 := GenS3{F0: (true || true)}
                                _ = v94
                            }
                        }
                    }
                    var v95 uint8 = 0
                    _ = v95
                }
            }
        }
        var v96 float64 = -2.25
        _ = fi23
        _ = v96
    }
    _ = v22
    return 0
}
func genFn1() int32 {
    var v97 uint = 4294967295
    for fi98 := 0; fi98 < 9; fi98++ {
        var v99 int32 = 702
        switch (fi98 * 0) {
        case -1:
            v100 := GenS6{F0: (v97 % ((42) | 1)), F1: -1}
            _ = v100
        case 231:
            if false {
                v97 -= 0
            } else {
                v97 += v97
                v101 := []int{(-(fi98)), (fi98 & (fi98 ^ -1)), -9223372036854775808}
                _ = v101
            }
        default:
            for fi102 := 0; fi102 < 2; fi102++ {
                v97 = v97
                v103 := []int{2147483647, (fi98 >> (((v97 | 20)) & 63))}
                v104 := GenS9{F0: 802, F1: (v99 & v99), F2: fi102}
                var v105 int16 = 32767
                _ = fi102
                _ = v103
                _ = v104
                _ = v105
            }
        }
        switch -24 {
        case 258:
            switch 128 {
            case -527:
                for fi106 := 0; fi106 < 7; fi106++ {
                    v97 -= (v97 - v97)
                    _ = fi106
                }
            case 1:
                v97 -= v97
            }
        default:
            for fi107 := 0; fi107 < 1; fi107++ {
                var v108 string = "foo bar"
                var v109 bool = true
                v110 := []int{(fi107 * fi98), fi98, (fi98 >> ((v97) & 63))}
                var v111 uint = (v97 / (((v97 * 18446744073709551615)) | 1))
                _ = fi107
                _ = v108
                _ = v109
                _ = v110
                _ = v111
            }
        }
        _ = fi98
        _ = v99
    }
    _ = v97
    return 42
}

func main() {
    var v1 int16 = 32767
    var v2 int16 = (v1 - (v1 & 32767))
    v3 := make(map[uint32]int)
    var v4 float64 = 3.14159
    fmt.Println(v2, v4)
    var v112 int8 = 1
    v1 = (v1 & 1)
    if ((!false) || (4294967295 >= 2)) {
        fmt.Println(v1, v1)
    } else {
        v112 = (v112 + (^(v112)))
        if false {
            var v113 int16 = v2
            var v114 int8 = v112
            _ = v113
            _ = v114
        } else {
            for fi115 := 0; fi115 < 11; fi115++ {
                v3[290] = fi115
                _ = fi115
            }
        }
        v4 = 1e10
    }
    v4 += v4
    v1 -= -1
    for fi116 := 0; fi116 < 3; fi116++ {
        var v117 float64 = v4
        _ = fi116
        _ = v117
    }
    v1 = (v1 + v1)
    if (false && true) {
        v118 := GenS9{F0: 42, F1: genFn1(), F2: 128}
        _ = v118
    } else {
        v119 := v3[470]
        _ = v119
    }
    var v120 int64 = 930
    for fi121 := 0; fi121 < 17; fi121++ {
        fmt.Println(fi121, v112)
        v120 = v120
        _ = fi121
    }
    if false {
        var v122 uint64 = 2
        _ = v122
    }
    v4 -= 0.0
    v123 := []int{9223372036854775807}
    v1 += -679
    v4 -= v4
    for fi124 := 0; fi124 < 12; fi124++ {
        var v125 int32 = 1
        _ = fi124
        _ = v125
    }
    fmt.Println("v1:", v1)
    fmt.Println("v112:", v112)
    fmt.Println("v120:", v120)
    fmt.Println("slicelen v123:", len(v123), v123)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v112
    _ = v120
    _ = v123
    _ = v2
    _ = v3
    _ = v4
}
