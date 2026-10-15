package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool; F2 uint64 }
type GenS2 struct { F0 string; F1 bool }
type GenS3 struct { F0 int32; F1 int64; F2 string }
type GenS4 struct { F0 string; F1 int; F2 int64 }
type GenS5 struct { F0 int64; F1 int8; F2 uint }
type GenS6 struct { F0 int }
type GenS7 struct { F0 int32; F1 int32 }
type GenS8 struct { F0 int16 }
type GenS9 struct { F0 uint; F1 string }
type GenS10 struct { F0 uint8 }
type GenS11 struct { F0 int8 }
type GenS12 struct { F0 int8 }
type GenS13 struct { F0 int32 }
func genFn0(p0 uint32, p1 int64, p2 int32) int {
    v7 := GenS1{F0: (p0 ^ p0), F1: (!false), F2: 575}
    v8 := GenS1{F0: p0, F1: ("hello" != "a"), F2: 1}
    for fi9 := 0; fi9 < 19; fi9++ {
        p1 = p1
        _ = fi9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return -75
}
func genFn1(p0 int16, p1 int8) int16 {
    var v10 float32 = 3.14159
    v11 := GenS1{F0: 2147483647, F1: ("" < ""), F2: 42}
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    return (p0 * 42)
}

func main() {
    var v1 int64 = 276
    var v2 int64 = v1
    var v3 int8 = 0
    v4 := []int{2147483647, 32768}
    var v5 uint32 = 2147483647
    var v6 int8 = v3
    switch 4294967295 {
    case 0:
        switch ("\n\t" + "foo bar") {
        case "hello":
            for fi12 := 0; fi12 < 15; fi12++ {
                var v13 uint8 = 42
                _ = fi12
                _ = v13
            }
        }
    }
    if false {
        if (!false) {
            switch "foo bar" {
            case "foo bar":
                var v14 int32 = int32(v6)
                _ = v14
            case "":
                v2 = -1
            case "a":
                switch "hello" {
                case "hello":
                    v3 += v3
                case "a":
                    v15 := GenS1{F0: 4294967295, F1: (128 <= 96), F2: 0}
                    _ = v15
                default:
                    v3 = v6
                }
            default:
                for fi16 := 0; fi16 < 12; fi16++ {
                    for fi17 := 0; fi17 < 4; fi17++ {
                        fmt.Println(v1)
                        _ = fi17
                    }
                    var v18 string = ("a" + "hello")
                    _ = fi16
                    _ = v18
                }
            }
        } else {
            var v19 uint32 = (^(v5))
            for fi20 := 0; fi20 < 3; fi20++ {
                var v21 int8 = v6
                v22 := GenS1{F0: v19, F1: false, F2: 2}
                if (1.5 <= 1.5) {
                    v3 = (v6 - (-(v21)))
                    v19 = 2147483648
                    _ = len(v4)
                }
                _ = fi20
                _ = v21
                _ = v22
            }
            if (v3 == v3) {
                _ = len(v4)
                var v23 uint32 = (^(v5))
                if false {
                    v2 -= v2
                    var v24 uint = 85
                    _ = v24
                } else {
                    fmt.Println(v23)
                    v19 += (^(v5))
                }
                _ = v23
            } else {
                var v25 uint16 = 0
                if (!true) {
                    fmt.Println(v19)
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                } else {
                    fmt.Println(v3, v25, v2)
                }
                _ = v25
            }
            _ = v19
        }
    } else {
        v26 := GenS1{F0: (v5 << ((0) & 63)), F1: (1.5 != (1e10 * float32(((-2.25 - ((0.0 + 0.0) * -2.25)) * (0.1 * 1.5))))), F2: 951}
        var v27 uint = 0
        v27 -= v27
        _ = v26
        _ = v27
    }
    if (1.5 >= float32(0.1)) {
        v28 := GenS2{F0: "", F1: false}
        var v29 int = 42
        _ = v28
        _ = v29
    }
    var v30 float64 = 1.5
    var v31 float32 = (0.1 + 1.5)
    switch ("foo bar" + ("hello" + "foo bar")) {
    case "foo bar":
        for fi32 := 0; fi32 < 3; fi32++ {
            v30 -= 1e10
            var v33 string = "foo bar"
            if true {
                if false {
                    var v34 float64 = (v30 + ((0.1 - 1.5) / (v30 - -2.25)))
                    _ = v34
                } else {
                    var v35 uint16 = 0
                    _ = v35
                }
                v1 -= (-(v2))
                if (127 == 65) {
                    fmt.Println(v30, fi32)
                    switch (fi32 & fi32) {
                    case -9223372036854775808:
                        for fi36 := 0; fi36 < 2; fi36++ {
                            var v37 string = (("" + ("hello" + ((v33 + (v33 + ("a" + ("hello" + (((v33 + ("a" + v33)) + "hello") + v33))))) + ""))) + v33)
                            switch fi32 {
                            case -984:
                                v38 := GenS3{F0: -1, F1: int64(v5), F2: v33}
                                var v39 int = genFn0(v5, (-(v2)), 1)
                                if ((v3 <= -1) || ((1.5 / v31) != ((v31 - v31) + (v31 * 1e10)))) {
                                    v30 = 3.14159
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                    var v40 uint64 = 18446744073709551615
                                    _ = v40
                                }
                                v41 := GenS4{F0: v37, F1: (v39 >> ((42) & 63)), F2: (v1 & -157)}
                                _ = v38
                                _ = v39
                                _ = v41
                            case 427:
                                var v42 int16 = 42
                                _ = v42
                            case 859:
                                v2 = v1
                            default:
                                var v43 int16 = 758
                                _ = v43
                            }
                            v30 += (1.5 + (float64(0.0) / ((1.5 * 1e10) * 1e10)))
                            _ = fi36
                            _ = v37
                        }
                    default:
                        if true {
                            var v44 uint = 835
                            _ = len(v4)
                            _ = v44
                        } else {
                            if false {
                                for fi45 := 0; fi45 < 17; fi45++ {
                                    if len(v4) > 0 {
                                        v4[0] = (fi45 * -2147483648)
                                    }
                                    if len(v4) > 0 {
                                        v4[0] = fi45
                                    }
                                    _ = fi45
                                }
                                v46 := GenS5{F0: (-(v1)), F1: v6, F2: 18446744073709551615}
                                _ = v46
                            } else {
                                switch fi32 {
                                case 127:
                                    v47 := GenS6{F0: genFn0(435, -826, 78)}
                                    _ = v47
                                case 1:
                                    var v48 bool = true
                                    _ = v48
                                case 42:
                                    fmt.Println(v5, v31)
                                default:
                                    if (!(!(v33 > v33))) {
                                        if len(v4) > 0 {
                                            v4[0] = (fi32 >> ((953) & 63))
                                        }
                                    }
                                }
                                fmt.Println(v1)
                            }
                            v49 := GenS7{F0: 867, F1: 0}
                            v50 := GenS1{F0: (^(v5)), F1: false, F2: 970}
                            _ = v49
                            _ = v50
                        }
                    }
                    if (false || (1e10 <= (1.5 / 0.1))) {
                        for fi51 := 0; fi51 < 1; fi51++ {
                            var v52 uint64 = 1
                            var v53 bool = false
                            switch genFn0((^(v5)), (^(v1)), 2147483647) {
                            case -9223372036854775808:
                                var v54 int16 = 32767
                                _ = v54
                            case 42:
                                v55 := GenS8{F0: 379}
                                _ = v55
                            default:
                                switch v33 {
                                case "\n\t":
                                    switch (fi51 ^ (fi32 | fi51)) {
                                    case 2147483647:
                                        if false {
                                            v6 += 127
                                            switch int(v1) {
                                            case 32767:
                                                if len(v4) > 0 {
                                                    v4[0] = (fi51 >> ((874) & 63))
                                                }
                                            case -948:
                                                v4 = append(v4, genFn0(v5, -1, -1))
                                            default:
                                                if v53 {
                                                    v2 = 9223372036854775807
                                                    v2 = 316
                                                } else {
                                                    var v56 int = (fi32 ^ fi32)
                                                    v30 = v30
                                                    _ = v56
                                                }
                                            }
                                            var v57 int = (fi51 - (^(fi32)))
                                            _ = v57
                                        }
                                    }
                                case "hello":
                                    var v58 float32 = float32(v30)
                                    _ = v58
                                case "":
                                    var v59 string = ("hello" + ("\n\t" + ""))
                                    _ = v59
                                default:
                                    _ = len(v4)
                                    v2 = v1
                                    if (-32768 != -646) {
                                        if (42 <= (-(v3))) {
                                            v60 := GenS9{F0: 900, F1: v33}
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                            v6 = 1
                                            _ = v60
                                        }
                                        v3 -= 0
                                        if false {
                                            fmt.Println(v53)
                                            v5 += v5
                                        } else {
                                            v53 = (346 <= -2147483648)
                                            v31 -= v31
                                            var v61 float32 = -2.25
                                            _ = v61
                                        }
                                    }
                                    var v62 float64 = v30
                                    _ = v62
                                }
                            }
                            _ = fi51
                            _ = v52
                            _ = v53
                        }
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                    } else {
                        fmt.Println(v6)
                        v33 = ((v33 + (v33 + v33)) + "")
                    }
                } else {
                    v3 += v6
                    var v63 int = (-(fi32))
                    _ = v63
                }
            }
            _ = fi32
            _ = v33
        }
    }
    v30 += v30
    v1 = 42
    var v64 int64 = v2
    v5 -= 743
    v2 -= (v1 & v64)
    for fi65 := 0; fi65 < 5; fi65++ {
        if (true || false) {
            var v66 bool = (false && true)
            if v66 {
                if (fi65 != (fi65 | (fi65 >> ((4294967295) & 63)))) {
                    v4 = append(v4, genFn0(v5, 1, 0))
                    for fi67 := 0; fi67 < 18; fi67++ {
                        v5 -= (v5 << ((65535) & 63))
                        _ = fi67
                    }
                } else {
                    for fi68 := 0; fi68 < 16; fi68++ {
                        if (!v66) {
                            var v69 int = (fi65 / (((fi65 / (((fi68 ^ (fi68 & (fi65 & fi68)))) | 1))) | 1))
                            v5 += (v5 << ((0) & 63))
                            v64 = -9223372036854775808
                            _ = v69
                        }
                        _ = fi68
                    }
                    v70 := GenS1{F0: v5, F1: (v66 || v66), F2: 1}
                    v30 = v30
                    _ = v70
                }
            }
            _ = v66
        } else {
            v71 := GenS3{F0: 0, F1: v64, F2: ("" + "foo bar")}
            v30 -= ((0.0 - (((v30 / (v30 / float64(v31))) - 3.14159) * v30)) - (float64((((3.14159 / 1.0) - (float32(v30) - v31)) * v31)) - float64(v31)))
            switch "foo bar" {
            case "a":
                var v72 int64 = (v2 >> ((255) & 63))
                _ = v72
            case "\n\t":
                switch (fi65 + fi65) {
                case -53:
                    v1 += (-(v2))
                case 844:
                    for fi73 := 0; fi73 < 11; fi73++ {
                        if true {
                            switch fi65 {
                            case -1:
                                if len(v4) > 0 {
                                    v4[0] = fi65
                                }
                            }
                            switch ("\n\t" + ("foo bar" + (("\n\t" + (("foo bar" + "a") + "\n\t")) + "hello"))) {
                            case "hello":
                                v74 := GenS10{F0: 42}
                                v75 := GenS11{F0: (v6 % ((v6) | 1))}
                                fmt.Println(v3, v2, v31)
                                _ = v74
                                _ = v75
                            case "\n\t":
                                if true {
                                    switch "hello" {
                                    case "\n\t":
                                        var v76 int8 = 51
                                        _ = v76
                                    case "foo bar":
                                        var v77 bool = (!false)
                                        var v78 string = ("hello" + "")
                                        v79 := GenS7{F0: -214, F1: int32(v5)}
                                        _ = v77
                                        _ = v78
                                        _ = v79
                                    default:
                                        v30 += 1.5
                                    }
                                    fmt.Println(fi65)
                                } else {
                                    fmt.Println(v1)
                                    if (true && false) {
                                        v80 := GenS1{F0: 0, F1: false, F2: 4294967295}
                                        switch fi73 {
                                        case 690:
                                            var v81 uint16 = 178
                                            if ((v5 - (v5 | v5)) < 2147483647) {
                                                if true {
                                                    var v82 string = ((("\n\t" + "a") + ("foo bar" + "")) + ("hello" + ""))
                                                    _ = v82
                                                } else {
                                                    var v83 uint8 = 127
                                                    v2 += v64
                                                    v84 := GenS12{F0: v3}
                                                    _ = v83
                                                    _ = v84
                                                }
                                                if true {
                                                    v31 = (-2.25 / (v31 + v31))
                                                    _ = len(v4)
                                                    var v85 bool = false
                                                    switch (("hello" + "hello") + "\n\t") {
                                                    case "foo bar":
                                                        var v86 bool = (2147483647 >= 1)
                                                        _ = v86
                                                    case "hello":
                                                        var v87 uint64 = 2
                                                        _ = v87
                                                    default:
                                                        fmt.Println(v5)
                                                        v31 -= ((v31 * ((float32((float64(0.0) * v30)) + v31) / v31)) - v31)
                                                        if ((v30 <= float64(-2.25)) && ((2 != uint64(v5)) || true)) {
                                                            if (v6 < (v6 >> ((0) & 63))) {
                                                                v64 = -1
                                                                if v85 {
                                                                    var v88 uint64 = 2
                                                                    var v89 uint32 = (^(v5))
                                                                    var v90 uint = 0
                                                                    _ = v88
                                                                    _ = v89
                                                                    _ = v90
                                                                }
                                                                fmt.Println(v5)
                                                                if len(v4) > 0 {
                                                                    v4[0] = 4294967295
                                                                }
                                                                v91 := GenS5{F0: (v1 | (v1 ^ v2)), F1: 127, F2: uint(v5)}
                                                                _ = v91
                                                            } else {
                                                                var v92 int8 = 80
                                                                var v93 uint8 = 0
                                                                _ = v92
                                                                _ = v93
                                                            }
                                                            var v94 uint8 = 2
                                                            v3 -= (v6 - (v3 >> ((42) & 63)))
                                                            _ = v94
                                                        } else {
                                                            v3 = v3
                                                            v95 := GenS11{F0: -1}
                                                            _ = v95
                                                        }
                                                        var v96 bool = v85
                                                        var v97 uint64 = 872
                                                        _ = v96
                                                        _ = v97
                                                    }
                                                    v64 += (^(v64))
                                                    _ = v85
                                                } else {
                                                    switch "\n\t" {
                                                    case "foo bar":
                                                        var v98 bool = false
                                                        _ = v98
                                                    case "hello":
                                                        var v99 string = "hello"
                                                        _ = v99
                                                    }
                                                }
                                            } else {
                                                var v100 bool = (1 <= 1)
                                                v31 += v31
                                                _ = v100
                                            }
                                            v2 = v64
                                            _ = v81
                                        default:
                                            var v101 uint32 = v5
                                            _ = v101
                                        }
                                        if len(v4) > 0 {
                                            v4[0] = fi73
                                        }
                                        _ = v80
                                    } else {
                                        v102 := GenS10{F0: 1}
                                        switch "foo bar" {
                                        case "a":
                                            if len(v4) > 0 {
                                                v4[0] = fi65
                                            }
                                        case "":
                                            var v103 uint32 = v5
                                            _ = v103
                                        }
                                        _ = v102
                                    }
                                }
                            }
                        } else {
                            switch "a" {
                            case "\n\t":
                                var v104 uint64 = 18446744073709551615
                                _ = v104
                            default:
                                v31 = float32(((1e10 / v30) + (float64((v31 + (((float32(v30) + v31) * (v31 / 1.0)) / v31))) + 0.1)))
                                var v105 float32 = v31
                                if false {
                                    var v106 float32 = float32(v30)
                                    _ = v106
                                }
                                v107 := GenS13{F0: 1}
                                v4 = append(v4, genFn0(v5, v2, 0))
                                if (2 > 42) {
                                    v108 := GenS8{F0: 42}
                                    var v109 uint32 = 0
                                    if len(v4) > 0 {
                                        v4[0] = fi73
                                    }
                                    _ = v108
                                    _ = v109
                                } else {
                                    v5 += (^(v5))
                                    switch "" {
                                    case "hello":
                                        var v110 int64 = v2
                                        _ = v110
                                    default:
                                        var v111 uint = 42
                                        _ = v111
                                    }
                                    var v112 float64 = v30
                                    _ = v112
                                }
                                _ = v105
                                _ = v107
                            }
                        }
                        v113 := GenS10{F0: 1}
                        v114 := GenS11{F0: v6}
                        _ = fi73
                        _ = v113
                        _ = v114
                    }
                case 790:
                    var v115 int = int(v3)
                    _ = v115
                }
            default:
                for fi116 := 0; fi116 < 13; fi116++ {
                    var v117 int64 = (v1 * v1)
                    _ = fi116
                    _ = v117
                }
            }
            _ = v71
        }
        if (1000 >= 1000) {
            if false {
                v118 := GenS8{F0: 772}
                _ = v118
            }
            v1 = 9223372036854775807
            fmt.Println(v3)
        } else {
            if len(v4) > 0 {
                v4[0] = (-(fi65))
            }
            v3 += v3
        }
        var v119 int8 = v6
        _ = fi65
        _ = v119
    }
    fmt.Println(v3, v5, v5)
    v120 := GenS1{F0: 964, F1: true, F2: 4294967295}
    if len(v4) > 0 {
        v4[0] = -596
    }
    v2 -= v2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v64:", v64)
    _ = v1
    _ = v120
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
    _ = v64
}
