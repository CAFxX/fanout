package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool }
type GenS2 struct { F0 int; F1 uint; F2 uint32 }
type GenS3 struct { F0 string; F1 int; F2 int }
type GenS4 struct { F0 uint8; F1 int64; F2 uint32 }
type GenS5 struct { F0 int8; F1 bool; F2 uint16 }
type GenS6 struct { F0 int64; F1 int64; F2 uint16 }
type GenS7 struct { F0 uint8; F1 uint8; F2 uint }
type GenS8 struct { F0 int8 }
type GenS9 struct { F0 uint; F1 int }
type GenS10 struct { F0 int16; F1 uint8; F2 string }
func genFn0() int {
    if false {
        v5 := GenS1{F0: 1, F1: (235 <= 2147483647)}
        _ = v5
    } else {
        if (true || (9223372036854775807 >= 32767)) {
            var v6 float64 = 0.0
            fmt.Println(v6)
            _ = v6
        } else {
            if ("" <= ("" + "foo bar")) {
                if (320 > 1) {
                    for fi7 := 0; fi7 < 20; fi7++ {
                        var v8 bool = (0 <= 42)
                        _ = fi7
                        _ = v8
                    }
                    var v9 int8 = -128
                    for fi10 := 0; fi10 < 7; fi10++ {
                        v11 := []int{fi10, -9223372036854775808, fi10, (fi10 >> ((255) & 63))}
                        var v12 uint32 = 2
                        _ = fi10
                        _ = v11
                        _ = v12
                    }
                    _ = v9
                } else {
                    switch 538 {
                    case 2147483647:
                        v13 := GenS1{F0: 0, F1: false}
                        _ = v13
                    }
                    v14 := []int{9223372036854775807}
                    v15 := GenS1{F0: 1, F1: true}
                    _ = v14
                    _ = v15
                }
            } else {
                var v16 float32 = 3.14159
                switch 42 {
                case 42:
                    v17 := []int{-2147483648, 128, 1, 127}
                    _ = v17
                case 374:
                    if (-32768 > 42) {
                        if true {
                            v18 := []int{127, 4294967295, -9223372036854775808}
                            switch (("" + ("hello" + "")) + "\n\t") {
                            case "\n\t":
                                v19 := GenS1{F0: 1, F1: false}
                                _ = v19
                            case "":
                                if false {
                                    v18 = append(v18, 128)
                                    if len(v18) >= 2 {
                                        v18 = v18[1:len(v18)]
                                    }
                                } else {
                                    for fi20 := 0; fi20 < 5; fi20++ {
                                        v21 := GenS1{F0: 1, F1: false}
                                        _ = fi20
                                        _ = v21
                                    }
                                    if false {
                                        var v22 int64 = 42
                                        _ = v22
                                    } else {
                                        v16 = -2.25
                                        var v23 int32 = 911
                                        for fi24 := 0; fi24 < 16; fi24++ {
                                            if (int8(v23) >= -128) {
                                                v16 += v16
                                            } else {
                                                var v25 float64 = -2.25
                                                _ = v25
                                            }
                                            fmt.Println(v23, fi24)
                                            v23 -= v23
                                            _ = fi24
                                        }
                                        _ = v23
                                    }
                                }
                            case "foo bar":
                                var v26 bool = (32767 == 1)
                                _ = v26
                            default:
                                fmt.Println(v16)
                            }
                            _ = v18
                        }
                        if ((1 > 1) && false) {
                            switch ("" + (("foo bar" + "") + "hello")) {
                            case "foo bar":
                                for fi27 := 0; fi27 < 9; fi27++ {
                                    v16 = 3.14159
                                    if false {
                                        for fi28 := 0; fi28 < 17; fi28++ {
                                            if true {
                                                v29 := []int{2147483647, fi27, fi28}
                                                _ = v29
                                            } else {
                                                if true {
                                                    var v30 int16 = 0
                                                    _ = v30
                                                }
                                                v31 := []int{836}
                                                v32 := GenS2{F0: fi28, F1: 234, F2: 2147483648}
                                                _ = len(v31)
                                                _ = v31
                                                _ = v32
                                            }
                                            if true {
                                                var v33 string = ("a" + (("foo bar" + ("hello" + "foo bar")) + "a"))
                                                v34 := GenS1{F0: 1, F1: (32767 >= 2)}
                                                _ = v33
                                                _ = v34
                                            }
                                            var v35 int64 = 691
                                            _ = fi28
                                            _ = v35
                                        }
                                        fmt.Println(v16, v16)
                                        var v36 uint16 = 32767
                                        _ = v36
                                    }
                                    switch (fi27 / (((fi27 * fi27)) | 1)) {
                                    case 2147483647:
                                        v37 := GenS2{F0: (-(fi27)), F1: 0, F2: 42}
                                        _ = v37
                                    default:
                                        v16 += v16
                                    }
                                    _ = fi27
                                }
                            case "":
                                v38 := []int{814, 875}
                                _ = v38
                            default:
                                var v39 int8 = 8
                                _ = v39
                            }
                            for fi40 := 0; fi40 < 7; fi40++ {
                                v41 := []int{(fi40 >> ((210) & 63)), -2147483648, (^(fi40)), (fi40 & fi40)}
                                _ = fi40
                                _ = v41
                            }
                            var v42 uint8 = 127
                            _ = v42
                        }
                    }
                default:
                    v43 := []int{128}
                    _ = v43
                }
                _ = v16
            }
        }
    }
    var v44 uint = 0
    _ = v44
    return 0
}

func main() {
    v1 := []int{-9223372036854775808, 2147483647, 606, 0}
    var v2 uint64 = 4294967295
    var v3 uint8 = 127
    var v4 bool = (false && true)
    var v45 int8 = -61
    var v46 int16 = 42
    var v47 float64 = 1.5
    if false {
        var v48 float32 = ((1.5 + (1e10 - float32(0.1))) - 1.5)
        var v49 string = "foo bar"
        var v50 int32 = -1
        _ = v48
        _ = v49
        _ = v50
    } else {
        var v51 uint8 = v3
        _ = v51
    }
    var v52 float64 = v47
    for fi53 := 0; fi53 < 18; fi53++ {
        var v54 bool = v4
        if len(v1) > 0 {
            v1[0] = genFn0()
        }
        var v55 uint8 = (v3 | (v3 << ((18446744073709551615) & 63)))
        _ = fi53
        _ = v54
        _ = v55
    }
    for fi56 := 0; fi56 < 15; fi56++ {
        v57 := GenS2{F0: (-(fi56)), F1: 18446744073709551615, F2: 2147483647}
        var v58 bool = false
        if v58 {
            var v59 uint32 = 955
            switch "" {
            case "foo bar":
                v45 = v45
            }
            var v60 uint = 1
            _ = v59
            _ = v60
        } else {
            v47 -= -2.25
            v45 = (v45 | -127)
            switch fi56 {
            case 128:
                if (v58 || ((^(v2)) > uint64(v3))) {
                    for fi61 := 0; fi61 < 17; fi61++ {
                        switch 128 {
                        case -1:
                            v2 = v2
                        case -466:
                            v46 -= (v46 / ((v46) | 1))
                        case 889:
                            switch "hello" {
                            case "\n\t":
                                v52 = v52
                            case "a":
                                if v4 {
                                    var v62 float64 = -2.25
                                    if (("\n\t" + "a") < "hello") {
                                        fmt.Println(v58, v58)
                                    } else {
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        v46 += v46
                                    }
                                    _ = v62
                                } else {
                                    var v63 bool = v58
                                    v64 := GenS1{F0: (v3 << ((65535) & 63)), F1: ((0 < 0) && (v3 >= v3))}
                                    _ = v63
                                    _ = v64
                                }
                            case "foo bar":
                                _ = len(v1)
                            }
                        default:
                            switch ("foo bar" + ("" + "")) {
                            case "":
                                switch "foo bar" {
                                case "foo bar":
                                    v65 := GenS3{F0: "\n\t", F1: fi56, F2: genFn0()}
                                    _ = v65
                                case "":
                                    fmt.Println(v2, v47, v4)
                                case "a":
                                    var v66 int8 = (v45 % ((127) | 1))
                                    _ = v66
                                default:
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                }
                            default:
                                if true {
                                    var v67 int16 = 23
                                    var v68 float32 = 3.14159
                                    v69 := GenS1{F0: (v3 & (v3 | (^(v3)))), F1: false}
                                    _ = v67
                                    _ = v68
                                    _ = v69
                                } else {
                                    switch ("" + "hello") {
                                    case "hello":
                                        var v70 int = fi56
                                        _ = v70
                                    }
                                    var v71 uint8 = (^(v3))
                                    v4 = (v2 != 478)
                                    v72 := GenS2{F0: fi61, F1: 42, F2: 208}
                                    var v73 uint = 4294967295
                                    fmt.Println(v46, v45)
                                    _ = v71
                                    _ = v72
                                    _ = v73
                                }
                            }
                        }
                        _ = len(v1)
                        var v74 int32 = -410
                        _ = fi61
                        _ = v74
                    }
                }
            case 2147483647:
                if (v2 > (v2 - (v2 - v2))) {
                    if v58 {
                        switch ("" + "foo bar") {
                        case "\n\t":
                            var v75 string = (("" + "a") + "hello")
                            _ = v75
                        case "a":
                            if (!false) {
                                v1 = append(v1, fi56)
                            }
                        case "foo bar":
                            v76 := GenS1{F0: (v3 >> ((42) & 63)), F1: (-2147483648 <= -164)}
                            _ = v76
                        default:
                            var v77 float32 = float32((((v47 / 1.5) / 1.0) * (v52 - 0.0)))
                            _ = v77
                        }
                        for fi78 := 0; fi78 < 8; fi78++ {
                            if ((v2 * v2) != v2) {
                                var v79 int = 127
                                _ = v79
                            } else {
                                var v80 int64 = -9223372036854775808
                                if false {
                                    v81 := GenS3{F0: ("\n\t" + "hello"), F1: 1, F2: 127}
                                    v47 += v47
                                    v82 := GenS4{F0: v3, F1: (v80 - v80), F2: 612}
                                    _ = v81
                                    _ = v82
                                } else {
                                    v58 = (-73 <= v45)
                                    v45 -= -65
                                    var v83 int = (fi78 >> ((255) & 63))
                                    _ = v83
                                }
                                _ = v80
                            }
                            _ = len(v1)
                            _ = fi78
                        }
                    }
                } else {
                    for fi84 := 0; fi84 < 5; fi84++ {
                        v1 = append(v1, genFn0())
                        var v85 string = "a"
                        _ = fi84
                        _ = v85
                    }
                    fmt.Println(v45, v3)
                    var v86 string = ""
                    _ = v86
                }
            case 4294967295:
                if true {
                    v3 += (v3 | 2)
                    var v87 int16 = v46
                    _ = v87
                } else {
                    v88 := GenS3{F0: "\n\t", F1: fi56, F2: (fi56 >> ((1) & 63))}
                    for fi89 := 0; fi89 < 20; fi89++ {
                        var v90 string = ""
                        if (v46 != v46) {
                            v47 += float64(3.14159)
                            if (!false) {
                                fmt.Println(v90, v3, v3)
                            }
                            var v91 float32 = -2.25
                            _ = v91
                        } else {
                            fmt.Println(fi89)
                            var v92 uint16 = 0
                            if (v58 || true) {
                                fmt.Println(v58)
                                if v4 {
                                    switch fi56 {
                                    case -9223372036854775808:
                                        switch v90 {
                                        case "foo bar":
                                            var v93 uint64 = 18446744073709551615
                                            _ = v93
                                        default:
                                            var v94 int16 = (v46 ^ v46)
                                            _ = v94
                                        }
                                    case -152:
                                        var v95 uint16 = v92
                                        _ = v95
                                    default:
                                        var v96 string = ("\n\t" + ((v90 + (v90 + v90)) + ""))
                                        _ = v96
                                    }
                                    switch fi56 {
                                    case -2147483648:
                                        if false {
                                            var v97 int32 = -674
                                            _ = v97
                                        } else {
                                            var v98 float64 = v47
                                            v2 = v2
                                            var v99 uint32 = 922
                                            _ = v98
                                            _ = v99
                                        }
                                    default:
                                        v90 = (v90 + v90)
                                    }
                                    if (!(-396 == genFn0())) {
                                        v100 := GenS3{F0: "hello", F1: (fi56 % (((fi56 << ((18446744073709551615) & 63))) | 1)), F2: 32}
                                        switch v90 {
                                        case "":
                                            v3 -= v3
                                        default:
                                            if (!((v90 + v90) <= "")) {
                                                var v101 float32 = (1.5 / 1.0)
                                                _ = v101
                                            }
                                        }
                                        v2 += (v2 - 42)
                                        _ = v100
                                    } else {
                                        var v102 int64 = 9223372036854775807
                                        _ = v102
                                    }
                                } else {
                                    var v103 uint = 0
                                    v104 := GenS5{F0: v45, F1: (v103 <= (v103 ^ 42)), F2: (v92 >> (((v103 / (((v103 - v103)) | 1))) & 63))}
                                    if ((v46 / (((v46 >> ((v103) & 63))) | 1)) == 42) {
                                        var v105 uint8 = (^(v3))
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        _ = v105
                                    }
                                    _ = v103
                                    _ = v104
                                }
                                fmt.Println(v3, v4)
                                v45 = v45
                                v106 := GenS5{F0: -106, F1: (((v2 <= v2) || false) && false), F2: uint16(v2)}
                                v107 := GenS6{F0: -1, F1: int64(v3), F2: (v92 >> ((229) & 63))}
                                _ = v106
                                _ = v107
                            } else {
                                switch (v90 + v90) {
                                case "hello":
                                    var v108 int = (fi56 ^ -705)
                                    _ = v108
                                }
                            }
                            var v109 uint32 = 4294967295
                            _ = v109
                            _ = v92
                        }
                        v3 -= (v3 & v3)
                        var v110 uint64 = (v2 % ((v2) | 1))
                        var v111 uint64 = 380
                        v112 := GenS7{F0: v3, F1: v3, F2: 4294967295}
                        v90 = (((v90 + ((v90 + v90) + v90)) + v90) + ("\n\t" + (v90 + v90)))
                        v4 = v58
                        v58 = (v4 && (-973 > 0))
                        var v113 int32 = 2147483647
                        _ = fi89
                        _ = v110
                        _ = v111
                        _ = v112
                        _ = v113
                        _ = v90
                    }
                    fmt.Println(fi56)
                    _ = v88
                }
            default:
                var v114 uint64 = v2
                _ = v114
            }
        }
        _ = fi56
        _ = v57
        _ = v58
    }
    v115 := GenS1{F0: v3, F1: true}
    if (v4 || v4) {
        v1 = append(v1, 253)
        for fi116 := 0; fi116 < 12; fi116++ {
            switch "" {
            case "":
                v117 := GenS6{F0: -1, F1: -879, F2: 42}
                _ = v117
            case "\n\t":
                if v4 {
                    v118 := GenS8{F0: -1}
                    v119 := GenS9{F0: 691, F1: fi116}
                    switch fi116 {
                    case 4294967295:
                        v120 := GenS10{F0: 32767, F1: 206, F2: "a"}
                        _ = v120
                    case 32767:
                        v4 = v4
                    case -652:
                        _ = len(v1)
                    default:
                        var v121 uint16 = 0
                        _ = v121
                    }
                    _ = v118
                    _ = v119
                } else {
                    v45 = v45
                }
            }
            _ = fi116
        }
    }
    v4 = true
    v3 = (v3 * 57)
    fmt.Println(v2)
    for fi122 := 0; fi122 < 17; fi122++ {
        var v123 float64 = v52
        v46 = 0
        _ = fi122
        _ = v123
    }
    v1 = append(v1, -2147483648)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("v52:", v52)
    _ = v1
    _ = v115
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v46
    _ = v47
    _ = v52
}
