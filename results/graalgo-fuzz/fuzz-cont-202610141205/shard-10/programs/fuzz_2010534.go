package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint8; F1 uint64 }
type GenS4 struct { F0 int32; F1 string }
type GenS5 struct { F0 uint; F1 uint64; F2 uint }
type GenS6 struct { F0 int16; F1 bool }
type GenS7 struct { F0 int32; F1 uint8 }
type GenS8 struct { F0 uint16; F1 int8; F2 int }
type GenS9 struct { F0 int64; F1 int16; F2 int }
type GenS10 struct { F0 uint64; F1 uint; F2 int8 }
type GenS11 struct { F0 uint16; F1 int; F2 uint }
func genFn0(p0 uint, p1 uint64) uint32 {
    for fi7 := 0; fi7 < 11; fi7++ {
        v8 := GenS1{F0: p1}
        p1 = 0
        _ = fi7
        _ = v8
    }
    for fi9 := 0; fi9 < 18; fi9++ {
        switch "\n\t" {
        case "a":
            for fi10 := 0; fi10 < 7; fi10++ {
                p1 = (p1 - p1)
                switch "hello" {
                case "foo bar":
                    var v11 uint64 = p1
                    _ = v11
                case "\n\t":
                    if false {
                        p1 += 901
                        var v12 string = "\n\t"
                        p1 = (p1 << ((p0) & 63))
                        _ = v12
                    } else {
                        p0 -= p0
                        if true {
                            var v13 int16 = -32768
                            v14 := []int{}
                            p1 -= p1
                            _ = v13
                            _ = v14
                        } else {
                            switch fi9 {
                            case -2147483648:
                                var v15 float32 = float32(1.5)
                                _ = v15
                            default:
                                var v16 float32 = 1.5
                                _ = v16
                            }
                            v17 := GenS1{F0: (p1 % (((^(p1))) | 1))}
                            _ = v17
                        }
                        fmt.Println(p1, p1, fi9)
                    }
                case "a":
                    v18 := GenS1{F0: (^(p1))}
                    _ = v18
                }
                _ = fi10
            }
        case "":
            for fi19 := 0; fi19 < 11; fi19++ {
                p1 = (p1 - (^(p1)))
                v20 := GenS2{F0: 255}
                var v21 int16 = -972
                if false {
                    var v22 uint64 = p1
                    _ = v22
                } else {
                    var v23 int8 = 42
                    v24 := GenS3{F0: 42, F1: uint64(p0)}
                    _ = v23
                    _ = v24
                }
                _ = fi19
                _ = v20
                _ = v21
            }
        case "\n\t":
            if ("foo bar" <= "\n\t") {
                if (32767 <= 705) {
                    var v25 uint = (p0 ^ (p0 * p0))
                    p0 += 42
                    _ = v25
                }
                p0 = 255
            } else {
                v26 := GenS1{F0: 42}
                _ = v26
            }
        }
        _ = fi9
    }
    p0 -= (p0 | 2)
    var v27 int8 = -32
    _ = p0
    _ = p1
    _ = v27
    return 286
}
func genFn1() int8 {
    var v28 float64 = 3.14159
    var v29 int8 = -27
    switch 954 {
    case 62:
        for fi30 := 0; fi30 < 20; fi30++ {
            v31 := GenS2{F0: 65535}
            _ = fi30
            _ = v31
        }
    case 2147483647:
        for fi32 := 0; fi32 < 6; fi32++ {
            var v33 int16 = 0
            _ = fi32
            _ = v33
        }
    default:
        var v34 string = ""
        _ = v34
    }
    for fi35 := 0; fi35 < 18; fi35++ {
        var v36 int8 = v29
        if (42 >= -2147483648) {
            for fi37 := 0; fi37 < 20; fi37++ {
                v38 := []int{(fi35 * 946), 32768, (fi37 & fi35), (fi37 % (((fi35 - (fi37 & fi37))) | 1))}
                v39 := GenS1{F0: 817}
                _ = fi37
                _ = v38
                _ = v39
            }
        } else {
            if (true && (0 > 255)) {
                fmt.Println(v28, v28)
            }
            var v40 uint16 = 0
            _ = v40
        }
        _ = fi35
        _ = v36
    }
    for fi41 := 0; fi41 < 11; fi41++ {
        if true {
            var v42 uint16 = 32767
            fmt.Println(v29)
            _ = v42
        }
        switch "\n\t" {
        case "\n\t":
            v43 := []int{-2147483648}
            _ = v43
        case "hello":
            v44 := GenS4{F0: 0, F1: "foo bar"}
            _ = v44
        default:
            switch "\n\t" {
            case "a":
                for fi45 := 0; fi45 < 11; fi45++ {
                    var v46 uint16 = 734
                    switch -65 {
                    case 32767:
                        v29 = v29
                    case 127:
                        var v47 bool = false
                        _ = v47
                    }
                    v28 = 0.1
                    _ = fi45
                    _ = v46
                }
            case "hello":
                var v48 int32 = 2147483647
                _ = v48
            }
        }
        v28 = v28
        _ = fi41
    }
    _ = v28
    _ = v29
    return v29
}

func main() {
    var v1 uint32 = 2
    v2 := []int{288, 9223372036854775807, 2147483647}
    v3 := []int{}
    v4 := make(map[string]uint32)
    var v5 uint32 = 1
    v6 := []uint32{v1}
    v5 = (v5 | (v5 >> ((0) & 63)))
    fmt.Println(v5)
    switch -9223372036854775808 {
    case 2147483647:
        v5 -= v5
    case 128:
        var v49 float64 = (float64((float32(-2.25) * 1.5)) * 1e10)
        _ = v49
    case -2147483648:
        v1 += 0
    }
    for fi50 := 0; fi50 < 11; fi50++ {
        fmt.Println(v5)
        _ = fi50
    }
    var v51 int16 = 1000
    v52 := GenS2{F0: 1}
    for fi53 := 0; fi53 < 4; fi53++ {
        if true {
            for fi54 := 0; fi54 < 18; fi54++ {
                if (!false) {
                    v55 := GenS1{F0: 4294967295}
                    var v56 uint = 970
                    _ = v55
                    _ = v56
                }
                v1 += (v1 + (v5 | 2))
                _ = fi54
            }
        }
        v57 := GenS2{F0: 32767}
        v1 = v1
        _ = fi53
        _ = v57
    }
    switch 1 {
    case -724:
        v1 -= 2147483648
    case 68:
        if false {
            switch ("a" + "hello") {
            case "":
                v1 += (v5 - (v1 & v1))
            case "\n\t":
                var v58 uint = 255
                _ = v58
            }
        } else {
            var v59 int16 = v51
            var v60 int32 = -1
            _ = v59
            _ = v60
        }
    case 4294967295:
        for fi61 := 0; fi61 < 16; fi61++ {
            fmt.Println(v1, v51, v5)
            if true {
                if false {
                    v62 := GenS5{F0: 18446744073709551615, F1: 613, F2: 1}
                    fmt.Println(v51)
                    if (18446744073709551615 == 0) {
                        var v63 float32 = -2.25
                        delete(v4, "a")
                        v63 = (float32(0.1) / v63)
                        _ = v63
                    }
                    _ = v62
                }
                if (false || (42 != uint64(v5))) {
                    var v64 int = -2147483648
                    _ = len(v2)
                    var v65 int8 = genFn1()
                    _ = v64
                    _ = v65
                }
            } else {
                if true {
                    v1 -= 293
                    v66 := GenS6{F0: 42, F1: false}
                    for fi67 := 0; fi67 < 6; fi67++ {
                        if (true || true) {
                            _ = len(v4)
                            var v68 int64 = 947
                            _ = v68
                        }
                        var v69 uint64 = 874
                        _ = fi67
                        _ = v69
                    }
                    _ = v66
                } else {
                    if false {
                        for fi70 := 0; fi70 < 9; fi70++ {
                            fmt.Println(v5)
                            if (42 >= 42) {
                                if (573 != 4294967295) {
                                    var v71 int = 1
                                    switch fi61 {
                                    case 492:
                                        var v72 int8 = -128
                                        _ = v72
                                    }
                                    if true {
                                        if false {
                                            var v73 float32 = -2.25
                                            var v74 uint32 = v5
                                            if true {
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                            }
                                            _ = v73
                                            _ = v74
                                        } else {
                                            fmt.Println(v51, fi70)
                                        }
                                    } else {
                                        var v75 uint = 2
                                        _ = v75
                                    }
                                    _ = v71
                                } else {
                                    var v76 uint32 = genFn0(721, 4294967295)
                                    var v77 bool = (508 == 0)
                                    v77 = v77
                                    _ = v76
                                    _ = v77
                                }
                                if ((-304 < -454) && false) {
                                    v1 -= v1
                                    v1 = (v5 >> ((172) & 63))
                                    switch "\n\t" {
                                    case "foo bar":
                                        v1 = v5
                                        fmt.Println(v1, v1, v5)
                                    default:
                                        var v78 float64 = (1.5 + 0.0)
                                        _ = v78
                                    }
                                }
                            }
                            v79 := GenS5{F0: 2, F1: 151, F2: 938}
                            v80 := v4[""]
                            _ = fi70
                            _ = v79
                            _ = v80
                        }
                    } else {
                        v81 := GenS7{F0: 524, F1: 206}
                        var v82 uint64 = 1
                        if (42 != 1) {
                            var v83 int32 = 214
                            var v84 uint64 = 18446744073709551615
                            if false {
                                v1 = v5
                            }
                            _ = v83
                            _ = v84
                        } else {
                            switch fi61 {
                            case -9223372036854775808:
                                v85 := GenS7{F0: -1, F1: 2}
                                _ = v85
                            case 42:
                                if (1 <= uint16(v5)) {
                                    v86 := GenS8{F0: 2, F1: 55, F2: (fi61 / ((fi61) | 1))}
                                    fmt.Println(v5, v1, v51)
                                    var v87 int32 = -154
                                    _ = v86
                                    _ = v87
                                } else {
                                    var v88 float64 = 0.1
                                    fmt.Println(v82, v51, v51)
                                    switch "\n\t" {
                                    case "a":
                                        v1 = v1
                                    case "":
                                        v89 := GenS6{F0: -15, F1: (-1 != -190)}
                                        _ = v89
                                    }
                                    _ = v88
                                }
                            case 127:
                                if (!((46 > 202) && (1 >= 9223372036854775807))) {
                                    v51 = v51
                                }
                            default:
                                v90 := GenS4{F0: 42, F1: ("\n\t" + "")}
                                _ = v90
                            }
                            switch (fi61 % ((fi61) | 1)) {
                            case 0:
                                for fi91 := 0; fi91 < 6; fi91++ {
                                    _ = len(v3)
                                    var v92 int32 = -95
                                    _ = fi91
                                    _ = v92
                                }
                            case -9223372036854775808:
                                for fi93 := 0; fi93 < 13; fi93++ {
                                    switch "" {
                                    case "foo bar":
                                        var v94 string = ""
                                        switch (fi61 & (fi61 * -666)) {
                                        case -267:
                                            switch fi93 {
                                            case 42:
                                                v5 = v5
                                            case 764:
                                                var v95 int16 = 1000
                                                _ = v95
                                            default:
                                                v96 := GenS7{F0: 1, F1: 169}
                                                _ = v96
                                            }
                                        case 733:
                                            if true {
                                                fmt.Println(fi61, v51, v94)
                                                v4[v94] = 412
                                                fmt.Println(fi93)
                                            } else {
                                                v97 := v4["foo bar"]
                                                v97 = (v1 << ((uint(v82)) & 63))
                                                delete(v4, v94)
                                                _ = v97
                                            }
                                        case 734:
                                            v5 = v5
                                        default:
                                            v94 = ("foo bar" + v94)
                                        }
                                        var v98 string = "hello"
                                        v51 = v51
                                        _ = v94
                                        _ = v98
                                    }
                                    _ = fi93
                                }
                            case 1:
                                var v99 bool = (-313 > int64(v82))
                                _ = v99
                            }
                            v100 := GenS9{F0: 405, F1: (v51 % ((160) | 1)), F2: fi61}
                            _ = v100
                        }
                        _ = v81
                        _ = v82
                    }
                    fmt.Println(fi61, v1, fi61)
                }
            }
            _ = fi61
        }
    default:
        var v101 uint32 = genFn0(42, 42)
        _ = v101
    }
    _ = len(v4)
    v2 = append(v2, 924)
    if ("hello" == "\n\t") {
        if (false && (true || (0 >= 531))) {
            _ = len(v4)
            for fi102 := 0; fi102 < 20; fi102++ {
                var v103 int64 = 9223372036854775807
                var v104 uint32 = 1
                _ = fi102
                _ = v103
                _ = v104
            }
            v51 = (-(v51))
        }
        fmt.Println(v5, v1, v5)
        v105 := GenS3{F0: 128, F1: 275}
        _ = v105
    } else {
        v106 := GenS3{F0: 127, F1: 4294967295}
        _ = v106
    }
    if true {
        var v107 string = ""
        _ = v107
    }
    v6 = append(v6, (^(v5)))
    switch "\n\t" {
    case "":
        switch ("hello" + ("foo bar" + "a")) {
        case "":
            if true {
                fmt.Println(v51)
            } else {
                var v108 uint32 = v5
                var v109 float32 = -2.25
                _ = v108
                _ = v109
            }
        }
    case "foo bar":
        if true {
            var v110 string = ""
            _ = v110
        } else {
            fmt.Println(v1, v1, v51)
            if (!false) {
                v111 := GenS10{F0: 18446744073709551615, F1: 2, F2: -128}
                _ = v111
            }
        }
    case "hello":
        for fi112 := 0; fi112 < 18; fi112++ {
            if true {
                var v113 bool = (!(((1e10 * (1e10 + ((1e10 + (0.0 / 1.5)) - float64(float32(-2.25))))) - ((0.1 / 3.14159) - 1.5)) < -2.25))
                _ = v113
            } else {
                var v114 bool = true
                _ = v114
            }
            for fi115 := 0; fi115 < 9; fi115++ {
                if false {
                    v116 := GenS11{F0: 42, F1: (-(fi115)), F2: 2}
                    _ = v116
                } else {
                    _ = len(v4)
                    v2 = append(v2, (-(fi112)))
                }
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                _ = fi115
            }
            v117 := GenS9{F0: -824, F1: v51, F2: 4294967295}
            _ = fi112
            _ = v117
        }
    }
    for fi118 := 0; fi118 < 2; fi118++ {
        var v119 bool = false
        v5 -= v1
        _ = fi118
        _ = v119
    }
    if (v1 != v1) {
        fmt.Println(v51, v51, v1)
        v1 -= (^(v1))
    } else {
        if true {
            v51 = (v51 - v51)
            var v120 uint16 = 852
            v120 -= (v120 >> ((4294967295) & 63))
            _ = v120
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v51:", v51)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v51
    _ = v52
    _ = v6
}
