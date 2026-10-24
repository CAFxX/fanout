package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint8; F1 int; F2 uint64 }
type GenS3 struct { F0 string; F1 int16 }
type GenS4 struct { F0 uint8; F1 string }
type GenS5 struct { F0 int32; F1 bool; F2 string }
type GenS6 struct { F0 string; F1 uint64 }
type GenS7 struct { F0 bool; F1 int32; F2 uint32 }
type GenS8 struct { F0 uint; F1 uint; F2 string }
type GenS9 struct { F0 uint; F1 int; F2 int8 }
type GenS10 struct { F0 int8 }
type GenS11 struct { F0 string }
type GenS12 struct { F0 uint32 }
type GenS13 struct { F0 uint64; F1 int16 }
func genFn0() uint8 {
    for fi7 := 0; fi7 < 9; fi7++ {
        if false {
            if true {
                v8 := GenS1{F0: 255}
                switch (fi7 & fi7) {
                case -2147483648:
                    v9 := GenS1{F0: 128}
                    _ = v9
                case -972:
                    v10 := GenS1{F0: 196}
                    _ = v10
                case 0:
                    v11 := []int{(fi7 * (fi7 % (((fi7 | (fi7 ^ 9223372036854775807))) | 1))), (fi7 | (fi7 % ((fi7) | 1)))}
                    _ = v11
                default:
                    var v12 string = "a"
                    _ = v12
                }
                _ = v8
            }
        } else {
            var v13 float64 = 0.1
            _ = v13
        }
        _ = fi7
    }
    for fi14 := 0; fi14 < 11; fi14++ {
        var v15 uint = 65535
        _ = fi14
        _ = v15
    }
    switch ("foo bar" + "foo bar") {
    case "a":
        switch 42 {
        case 4294967295:
            var v16 uint32 = 4294967295
            _ = v16
        case 1:
            v17 := GenS2{F0: 216, F1: -499, F2: 795}
            _ = v17
        default:
            if true {
                if (true || (481 == -2147483648)) {
                    var v18 uint8 = 127
                    var v19 uint = 255
                    v18 += (v18 / ((2) | 1))
                    _ = v18
                    _ = v19
                } else {
                    if true {
                        v20 := []int{397, -2147483648, -1}
                        var v21 int16 = 0
                        _ = v20
                        _ = v21
                    }
                }
            }
        }
    case "foo bar":
        for fi22 := 0; fi22 < 1; fi22++ {
            v23 := GenS3{F0: ("" + ("" + "a")), F1: -32768}
            var v24 int = -207
            var v25 uint64 = 4294967295
            _ = fi22
            _ = v23
            _ = v24
            _ = v25
        }
    }
    return 10
}
func genFn1() uint {
    switch "foo bar" {
    case "foo bar":
        var v26 string = "\n\t"
        _ = v26
    default:
        var v27 int32 = 0
        _ = v27
    }
    var v28 int8 = 100
    v28 += (v28 ^ (v28 << ((729) & 63)))
    _ = v28
    return 255
}

func main() {
    var v1 uint8 = 1
    v2 := []uint32{969, 42, 2147483647}
    var v3 uint64 = 0
    var v4 int64 = 1
    var v5 float32 = (-2.25 * 1e10)
    v6 := []uint32{65535}
    v3 += v3
    v3 = v3
    v3 += v3
    switch 974 {
    case 1:
        if len(v6) > 0 {
            v6[0] = 0
        }
    case 920:
        for fi29 := 0; fi29 < 16; fi29++ {
            if (!(!(65535 <= 799))) {
                switch (fi29 | (fi29 - fi29)) {
                case 42:
                    switch fi29 {
                    case -9223372036854775808:
                        if (2147483647 >= 42) {
                            for fi30 := 0; fi30 < 14; fi30++ {
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                var v31 uint16 = 32767
                                var v32 uint8 = (v1 % (((v1 << ((uint(v3)) & 63))) | 1))
                                switch ("foo bar" + "foo bar") {
                                case "a":
                                    if false {
                                        var v33 string = ("foo bar" + ("" + ("foo bar" + "")))
                                        var v34 int16 = 32767
                                        v1 = uint8(v3)
                                        _ = v33
                                        _ = v34
                                    } else {
                                        v3 -= 700
                                        v1 -= v1
                                    }
                                    v35 := GenS3{F0: "\n\t", F1: -496}
                                    _ = v35
                                case "":
                                    v36 := GenS3{F0: "hello", F1: -1}
                                    _ = v36
                                default:
                                    var v37 uint32 = 89
                                    _ = v37
                                }
                                _ = len(v2)
                                v38 := GenS1{F0: v1}
                                if true {
                                    switch (("\n\t" + "foo bar") + ("hello" + (((("" + "hello") + "foo bar") + "\n\t") + "foo bar"))) {
                                    case "a":
                                        if true {
                                            var v39 float32 = 0.1
                                            v40 := GenS4{F0: genFn0(), F1: (("hello" + "") + "a")}
                                            v39 += (float32(3.14159) + 3.14159)
                                            v41 := GenS3{F0: ("hello" + ((("\n\t" + "foo bar") + ("\n\t" + ((("a" + "") + "") + "foo bar"))) + "a")), F1: 0}
                                            _ = v39
                                            _ = v40
                                            _ = v41
                                        }
                                    case "":
                                        v42 := GenS2{F0: 2, F1: (^(fi29)), F2: 727}
                                        _ = v42
                                    case "hello":
                                        fmt.Println(v31)
                                    default:
                                        v3 -= v3
                                    }
                                    if false {
                                        v43 := GenS5{F0: 371, F1: (-76 >= 67), F2: "a"}
                                        var v44 float32 = float32(0.1)
                                        _ = v43
                                        _ = v44
                                    } else {
                                        if ((^(v1)) != 128) {
                                            var v45 int8 = 42
                                            v46 := GenS5{F0: -1, F1: false, F2: ("a" + "foo bar")}
                                            _ = v45
                                            _ = v46
                                        } else {
                                            switch 127 {
                                            case -545:
                                                if (((v1 + (v1 << ((4294967295) & 63))) >= (v32 | (v32 | v1))) && true) {
                                                    if len(v2) > 0 {
                                                        v2[0] = 42
                                                    }
                                                    if (!true) {
                                                        v1 = v1
                                                        v47 := GenS1{F0: (v1 ^ (v32 & v1))}
                                                        _ = v47
                                                    }
                                                    if len(v2) >= 2 {
                                                        v2 = v2[1:len(v2)]
                                                    }
                                                } else {
                                                    v48 := GenS2{F0: (v1 + v32), F1: fi29, F2: v3}
                                                    if false {
                                                        v3 = v3
                                                        v1 -= (v32 + 79)
                                                    }
                                                    v49 := GenS6{F0: ("" + (((("\n\t" + "a") + ((("\n\t" + "foo bar") + ((("\n\t" + "foo bar") + "") + "\n\t")) + "foo bar")) + "foo bar") + "a")), F1: v3}
                                                    _ = v48
                                                    _ = v49
                                                }
                                            case 0:
                                                switch 9223372036854775807 {
                                                case 4294967295:
                                                    if false {
                                                        var v50 uint32 = 772
                                                        if ((v4 * 42) == (v4 % (((v4 / ((v4) | 1))) | 1))) {
                                                            if len(v2) >= 2 {
                                                                v2 = v2[1:len(v2)]
                                                            }
                                                        }
                                                        _ = v50
                                                    } else {
                                                        var v51 int16 = 957
                                                        v6 = append(v6, 177)
                                                        if true {
                                                            var v52 uint16 = 65535
                                                            _ = v52
                                                        }
                                                        _ = v51
                                                    }
                                                    var v53 float32 = (float32(0.1) - 0.1)
                                                    var v54 uint8 = (v32 & (v32 / (((v32 ^ (v32 + (v1 * v1)))) | 1)))
                                                    _ = v53
                                                    _ = v54
                                                case 365:
                                                    var v55 uint = genFn1()
                                                    _ = v55
                                                default:
                                                    v3 = (v3 / ((v3) | 1))
                                                }
                                            }
                                            var v56 int64 = v4
                                            var v57 uint16 = (v31 >> ((uint(v3)) & 63))
                                            _ = v56
                                            _ = v57
                                        }
                                        switch fi30 {
                                        case 32768:
                                            fmt.Println(v32, v4)
                                        case -2147483648:
                                            var v58 uint64 = (v3 % ((42) | 1))
                                            _ = v58
                                        case 924:
                                            v1 += v1
                                            var v59 float64 = 0.1
                                            _ = v59
                                        }
                                        switch (fi30 ^ fi29) {
                                        case -943:
                                            if ((!false) || (4294967295 != (v3 / (((v3 & v3)) | 1)))) {
                                                switch fi30 {
                                                case 32768:
                                                    var v60 string = "hello"
                                                    _ = v60
                                                case 9223372036854775807:
                                                    v1 += 42
                                                default:
                                                    var v61 uint64 = 42
                                                    v32 += (v1 + (v32 >> ((4294967295) & 63)))
                                                    _ = v61
                                                }
                                                v62 := GenS7{F0: (v3 < (v3 & 4294967295)), F1: -879, F2: 2}
                                                _ = v62
                                            } else {
                                                v63 := GenS8{F0: 793, F1: 42, F2: "\n\t"}
                                                fmt.Println(v5)
                                                v32 -= uint8(v3)
                                                _ = v63
                                            }
                                            v32 -= v1
                                        case 2147483647:
                                            if false {
                                                var v64 bool = true
                                                _ = v64
                                            }
                                        case 404:
                                            fmt.Println(v32)
                                            var v65 float64 = 3.14159
                                            _ = v65
                                        }
                                    }
                                }
                                v66 := GenS9{F0: 416, F1: (fi30 + 42), F2: 42}
                                v3 += v3
                                v3 -= 97
                                _ = fi30
                                _ = v31
                                _ = v32
                                _ = v38
                                _ = v66
                            }
                            if (-298 == 0) {
                                v5 = 0.0
                                if (32767 != 1000) {
                                    switch -1 {
                                    case 127:
                                        for fi67 := 0; fi67 < 15; fi67++ {
                                            if (true || (true && true)) {
                                                if (!false) {
                                                    v5 = ((v5 / (v5 * v5)) / 3.14159)
                                                    if false {
                                                        fmt.Println(v5, v4, v3)
                                                        var v68 float64 = 0.1
                                                        _ = v68
                                                    } else {
                                                        var v69 int16 = 360
                                                        var v70 uint8 = v1
                                                        _ = v69
                                                        _ = v70
                                                    }
                                                }
                                            }
                                            _ = fi67
                                        }
                                    case 782:
                                        if false {
                                            var v71 bool = (!(127 != (v1 ^ v1)))
                                            if false {
                                                _ = len(v2)
                                                v5 += (0.0 * (v5 - v5))
                                                if (v71 || true) {
                                                    var v72 uint32 = 2147483647
                                                    v72 = v72
                                                    _ = v72
                                                } else {
                                                    v5 = ((0.1 + (v5 * (v5 / (v5 - v5)))) * (v5 + (v5 + v5)))
                                                }
                                            }
                                            fmt.Println(v1, v5)
                                            _ = v71
                                        } else {
                                            v73 := GenS10{F0: 42}
                                            fmt.Println(v1, v3)
                                            _ = v73
                                        }
                                    case 99:
                                        v4 += (v4 >> ((587) & 63))
                                    }
                                    var v74 uint16 = 42
                                    var v75 float64 = (1e10 * 1e10)
                                    _ = v74
                                    _ = v75
                                }
                            }
                            fmt.Println(v5, v1)
                        } else {
                            var v76 int = fi29
                            _ = v76
                        }
                    default:
                        v4 = -984
                    }
                case 32767:
                    var v77 int = (fi29 ^ fi29)
                    _ = v77
                case 2147483647:
                    fmt.Println(v1)
                default:
                    var v78 float32 = v5
                    _ = v78
                }
            }
            fmt.Println(v1, v4, v3)
            _ = fi29
        }
    default:
        switch "a" {
        case "a":
            v1 = v1
        case "hello":
            if (false || false) {
                var v79 float32 = v5
                _ = v79
            } else {
                for fi80 := 0; fi80 < 16; fi80++ {
                    var v81 uint64 = (v3 << ((65535) & 63))
                    var v82 bool = false
                    fmt.Println(v82, v3)
                    _ = fi80
                    _ = v81
                    _ = v82
                }
            }
        case "\n\t":
            v1 = genFn0()
        }
    }
    for fi83 := 0; fi83 < 3; fi83++ {
        for fi84 := 0; fi84 < 14; fi84++ {
            if len(v6) > 0 {
                v6[0] = 122
            }
            var v85 uint = 191
            if true {
                if len(v6) > 0 {
                    v6[0] = 525
                }
                fmt.Println(v4)
            } else {
                if ((("" + "foo bar") + "a") != "foo bar") {
                    v4 -= (v4 << ((2) & 63))
                    var v86 uint8 = 128
                    if (true && ((((2 != 16) && (!true)) && (false && ((v4 * v4) == (-(v4))))) || false)) {
                        var v87 uint16 = 2
                        switch fi83 {
                        case 32767:
                            switch fi83 {
                            case 843:
                                var v88 int64 = v4
                                _ = v88
                            case -1:
                                v87 += v87
                            case 0:
                                var v89 int64 = (v4 >> (((v85 % ((0) | 1))) & 63))
                                _ = v89
                            default:
                                v90 := GenS7{F0: (2 > 1), F1: 223, F2: 2147483648}
                                _ = v90
                            }
                        }
                        v3 -= v3
                        _ = v87
                    }
                    _ = v86
                }
            }
            var v91 uint = genFn1()
            v92 := GenS5{F0: -2147483648, F1: (v3 < (v3 / (((v3 & (v3 | v3))) | 1))), F2: "a"}
            if true {
                v93 := GenS10{F0: -128}
                _ = v93
            }
            if len(v6) > 0 {
                v6[0] = 206
            }
            if (v3 > (v3 >> ((genFn1()) & 63))) {
                var v94 string = "foo bar"
                switch fi84 {
                case -2147483648:
                    fmt.Println(v5, fi83, v4)
                case 4294967295:
                    if false {
                        v94 = v94
                        v1 = v1
                    }
                default:
                    if false {
                        var v95 int64 = (^(v4))
                        v3 = (v3 + 4294967295)
                        _ = v95
                    } else {
                        fmt.Println(v94, fi83)
                        var v96 int32 = 0
                        _ = v96
                    }
                    fmt.Println(v94, v1, v5)
                    if false {
                        var v97 int64 = (-(v4))
                        fmt.Println(v4, v4, fi83)
                        var v98 uint64 = 2
                        _ = v97
                        _ = v98
                    }
                    if ((v4 <= int64(v3)) || false) {
                        var v99 int32 = 2147483647
                        var v100 int8 = 1
                        _ = v100
                        _ = v99
                    }
                    var v101 uint16 = 1
                    _ = v101
                }
                _ = v94
            }
            _ = fi84
            _ = v85
            _ = v91
            _ = v92
        }
        if len(v2) > 0 {
            v2[0] = 712
        }
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        _ = fi83
    }
    var v102 int8 = 127
    switch -1 {
    case -606:
        for fi103 := 0; fi103 < 6; fi103++ {
            if (!true) {
                fmt.Println(v3, v1)
                v104 := GenS5{F0: -2147483648, F1: false, F2: "hello"}
                _ = v104
            }
            var v105 uint8 = v1
            _ = fi103
            _ = v105
        }
    case -1:
        var v106 bool = (42 < 2147483647)
        _ = v106
    }
    v1 = (v1 * (v1 << ((2) & 63)))
    v107 := GenS11{F0: ""}
    for fi108 := 0; fi108 < 6; fi108++ {
        if false {
            var v109 float64 = 1.5
            _ = v109
        } else {
            if ("" >= "\n\t") {
                for fi110 := 0; fi110 < 18; fi110++ {
                    var v111 uint8 = (v1 % ((1) | 1))
                    _ = fi110
                    _ = v111
                }
                var v112 uint8 = v1
                _ = v112
            }
            fmt.Println(v3, v102)
            v1 -= v1
        }
        fmt.Println(fi108)
        if (2147483647 <= 42) {
            var v113 uint16 = 42
            if ((v5 * (3.14159 * float32(-2.25))) <= 1.5) {
                v1 = v1
                v114 := GenS12{F0: 266}
                _ = v114
            }
            var v115 int16 = -59
            _ = v113
            _ = v115
        } else {
            for fi116 := 0; fi116 < 12; fi116++ {
                v117 := GenS13{F0: (v3 * (^(v3))), F1: -1}
                v118 := GenS2{F0: v1, F1: fi116, F2: (v3 | 840)}
                fmt.Println(fi108, v1)
                var v119 int32 = -2147483648
                var v120 uint64 = 1
                _ = fi116
                _ = v117
                _ = v118
                _ = v119
                _ = v120
            }
            v5 = v5
            v3 = (v3 - uint64(v1))
        }
        _ = fi108
    }
    var v121 int32 = -2147483648
    for fi122 := 0; fi122 < 5; fi122++ {
        v102 = 127
        _ = fi122
    }
    fmt.Println("v1:", v1)
    fmt.Println("v102:", v102)
    fmt.Println("v121:", v121)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v102
    _ = v107
    _ = v121
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
