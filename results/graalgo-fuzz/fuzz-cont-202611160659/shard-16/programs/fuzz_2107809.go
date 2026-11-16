package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int; F1 string; F2 uint64 }
type GenS5 struct { F0 int32; F1 int; F2 int32 }
type GenS6 struct { F0 uint32; F1 uint16; F2 int32 }
type GenS7 struct { F0 int64; F1 uint32; F2 uint64 }
type GenS8 struct { F0 string; F1 int64 }
type GenS9 struct { F0 uint32; F1 uint64 }
type GenS10 struct { F0 string; F1 int16 }
type GenS11 struct { F0 uint; F1 int32; F2 uint32 }
type GenS12 struct { F0 uint32 }
type GenS13 struct { F0 string; F1 uint16; F2 int }
type GenS14 struct { F0 int8 }
type GenS15 struct { F0 int64 }
type GenS16 struct { F0 uint64; F1 uint64 }
type GenS17 struct { F0 int8 }
type GenS18 struct { F0 int; F1 bool }
type GenS19 struct { F0 int8 }
type GenS20 struct { F0 string; F1 int }
type GenS21 struct { F0 int16; F1 uint }
type GenS22 struct { F0 bool }
type GenS23 struct { F0 int64; F1 int32; F2 int32 }
type GenS24 struct { F0 int8 }
type GenS25 struct { F0 int8; F1 uint }
type GenS26 struct { F0 int; F1 int8 }
type GenS27 struct { F0 uint; F1 int16 }
func genFn0() uint16 {
    var v7 int16 = 184
    for fi8 := 0; fi8 < 12; fi8++ {
        for fi9 := 0; fi9 < 10; fi9++ {
            v10 := GenS1{F0: 0}
            var v11 float32 = -2.25
            _ = fi9
            _ = v10
            _ = v11
        }
        var v12 float64 = 0.1
        _ = fi8
        _ = v12
    }
    fmt.Println(v7)
    v13 := GenS2{F0: 65535}
    switch "" {
    case "\n\t":
        var v14 uint32 = 65535
        _ = v14
    case "a":
        switch "a" {
        case "foo bar":
            fmt.Println(v7)
        case "\n\t":
            v15 := GenS1{F0: -262}
            _ = v15
        }
    case "foo bar":
        var v16 int64 = 42
        _ = v16
    default:
        v17 := GenS1{F0: 0}
        _ = v17
    }
    _ = v13
    _ = v7
    return 32767
}
func genFn1(p0 bool, p1 int8, p2 int64) uint64 {
    v18 := GenS3{F0: p1}
    switch "hello" {
    case "hello":
        fmt.Println(p2, p0)
    case "a":
        var v19 int8 = p1
        _ = v19
    default:
        fmt.Println(p1, p0, p1)
    }
    if false {
        for fi20 := 0; fi20 < 6; fi20++ {
            if (143 <= 435) {
                v21 := GenS4{F0: (fi20 >> ((1) & 63)), F1: "a", F2: 2}
                v22 := GenS5{F0: 42, F1: fi20, F2: 42}
                _ = v21
                _ = v22
            }
            switch "foo bar" {
            case "hello":
                p0 = p0
            default:
                v23 := []int{fi20, 42}
                _ = v23
            }
            _ = fi20
        }
    }
    fmt.Println(p1, p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v18
    return 18446744073709551615
}

func main() {
    var v1 bool = (471 != -1)
    var v2 int8 = -114
    var v3 uint8 = 1
    v4 := []string{"hello"}
    v5 := make(map[uint32]uint32)
    var v6 string = "foo bar"
    v3 += v3
    if (!false) {
        v5[42] = 328
        switch v6 {
        case "foo bar":
            delete(v5, 65535)
        case "\n\t":
            if false {
                for fi24 := 0; fi24 < 9; fi24++ {
                    if len(v4) > 0 {
                        v4[0] = ("foo bar" + "\n\t")
                    }
                    var v25 int8 = (v2 * 127)
                    for fi26 := 0; fi26 < 11; fi26++ {
                        var v27 uint = 2
                        _ = fi26
                        _ = v27
                    }
                    _ = fi24
                    _ = v25
                }
            }
        case "hello":
            if v1 {
                fmt.Println(v2, v3)
                for fi28 := 0; fi28 < 12; fi28++ {
                    v2 = v2
                    _ = fi28
                }
                if v1 {
                    if (v1 && true) {
                        if ((((401 <= -1) && (v1 && true)) && (v1 && (42 >= -1))) && false) {
                            if len(v4) > 0 {
                                v4[0] = v6
                            }
                        } else {
                            for fi29 := 0; fi29 < 19; fi29++ {
                                if (!(!v1)) {
                                    switch "foo bar" {
                                    case "hello":
                                        for fi30 := 0; fi30 < 9; fi30++ {
                                            _ = len(v4)
                                            v31 := GenS6{F0: 4294967295, F1: 856, F2: 42}
                                            var v32 int = (fi30 % (((fi30 & (fi29 & (-(fi29))))) | 1))
                                            switch v6 {
                                            case "":
                                                v33 := GenS6{F0: 0, F1: genFn0(), F2: 1}
                                                _ = v33
                                            }
                                            var v34 uint8 = v3
                                            v6 = "a"
                                            fmt.Println(fi30)
                                            _ = fi30
                                            _ = v31
                                            _ = v32
                                            _ = v34
                                        }
                                    }
                                    for fi35 := 0; fi35 < 18; fi35++ {
                                        v1 = v1
                                        if false {
                                            var v36 uint8 = (v3 >> ((36) & 63))
                                            if v1 {
                                                v37 := GenS7{F0: 0, F1: 607, F2: 4294967295}
                                                v38 := GenS2{F0: 44}
                                                var v39 uint64 = 0
                                                v3 -= (^(v3))
                                                fmt.Println(fi29, fi35, v2)
                                                _ = v37
                                                _ = v38
                                                _ = v39
                                            } else {
                                                if ((^(v2)) >= 1) {
                                                    v2 -= (^(v2))
                                                    var v40 int8 = v2
                                                    _ = len(v5)
                                                    _ = v40
                                                }
                                                if (((float32((-2.25 + 1e10)) + ((1.5 * float32(float64(3.14159))) + -2.25)) * 0.0) > (-2.25 * 0.1)) {
                                                    if (-128 <= (-(v2))) {
                                                        if (1e10 < 1e10) {
                                                            if v1 {
                                                                var v41 int64 = -9223372036854775808
                                                                _ = v41
                                                            }
                                                            var v42 float64 = float64(-2.25)
                                                            fmt.Println(v6)
                                                            _ = v42
                                                        } else {
                                                            v43 := GenS8{F0: "\n\t", F1: -9223372036854775808}
                                                            if ((v6 + (("" + "\n\t") + "a")) <= ("foo bar" + (v6 + ((v6 + v6) + (("\n\t" + "") + ""))))) {
                                                                if v1 {
                                                                    fmt.Println(fi35, fi35)
                                                                }
                                                                _ = len(v4)
                                                            }
                                                            _ = v43
                                                        }
                                                        v3 += (v36 | 174)
                                                    }
                                                    v36 = v3
                                                    var v44 int = (fi29 * fi35)
                                                    _ = v44
                                                } else {
                                                    if ("hello" > v6) {
                                                        delete(v5, 846)
                                                        fmt.Println(v3)
                                                        if false {
                                                            if v1 {
                                                                v1 = (0 != 100)
                                                                v3 = 122
                                                                switch (fi35 / (((fi29 + int(v2))) | 1)) {
                                                                case -707:
                                                                    v2 += (v2 + 127)
                                                                    if ((fi35 + (fi35 % (((fi35 | (-(fi29)))) | 1))) > 0) {
                                                                        var v45 int16 = 1
                                                                        var v46 uint = 65535
                                                                        _ = v45
                                                                        _ = v46
                                                                    }
                                                                    v6 = v6
                                                                    fmt.Println(fi29, fi35)
                                                                    var v47 uint64 = 0
                                                                    _ = v47
                                                                }
                                                                v6 = v6
                                                                v1 = (1 >= 255)
                                                            } else {
                                                                switch 32768 {
                                                                case -473:
                                                                    v3 -= (^(v36))
                                                                case -768:
                                                                    var v48 int16 = 42
                                                                    _ = v48
                                                                case 32768:
                                                                    var v49 string = (((v6 + (v6 + "hello")) + ((v6 + (v6 + "")) + v6)) + v6)
                                                                    _ = v49
                                                                }
                                                                var v50 float64 = (0.0 - (((1.5 * (float64(((float32(0.1) * (0.0 / -2.25)) / (-2.25 + -2.25))) * (1e10 / 1e10))) - (1.5 + 0.0)) / (1e10 / 3.14159)))
                                                                _ = v50
                                                            }
                                                            var v51 uint16 = 42
                                                            _ = v51
                                                        }
                                                    } else {
                                                        v6 = (v6 + ("" + v6))
                                                        v6 = v6
                                                        var v52 uint = 65535
                                                        _ = v52
                                                    }
                                                }
                                                if v1 {
                                                    v1 = false
                                                } else {
                                                    v1 = false
                                                    v2 = -1
                                                    v3 += v3
                                                }
                                                fmt.Println(fi35, v6, v1)
                                                if len(v4) >= 2 {
                                                    v4 = v4[1:len(v4)]
                                                }
                                                if v1 {
                                                    v53 := GenS4{F0: 0, F1: "hello", F2: 2}
                                                    if (286 < 800) {
                                                        var v54 int32 = -2147483648
                                                        v3 = 222
                                                        var v55 int64 = 1
                                                        var v56 int8 = 1
                                                        v6 = (v6 + "a")
                                                        switch "foo bar" {
                                                        case "hello":
                                                            v55 = (v55 % ((42) | 1))
                                                        default:
                                                            v57 := GenS2{F0: genFn0()}
                                                            _ = v57
                                                        }
                                                        var v58 int16 = -32768
                                                        _ = v54
                                                        _ = v55
                                                        _ = v56
                                                        _ = v58
                                                    }
                                                    fmt.Println(fi29)
                                                    v2 = (v2 / ((v2) | 1))
                                                    _ = v53
                                                } else {
                                                    fmt.Println(v2, v36, v1)
                                                    fmt.Println(v6, v2, fi35)
                                                    v6 = (v6 + "hello")
                                                }
                                                v1 = true
                                            }
                                            if true {
                                                if (v1 && ((42 >= 1) && (((v36 | v3) >= v36) && (((v6 + (("hello" + ("a" + ((v6 + "a") + v6))) + v6)) < ("a" + v6)) && v1)))) {
                                                    switch fi35 {
                                                    case -9223372036854775808:
                                                        v59 := GenS8{F0: v6, F1: -47}
                                                        _ = v59
                                                    default:
                                                        switch fi35 {
                                                        case 1:
                                                            v2 -= v2
                                                        default:
                                                            v3 += v3
                                                        }
                                                    }
                                                    v60 := GenS9{F0: 2, F1: 18446744073709551615}
                                                    _ = v60
                                                }
                                                if v1 {
                                                    v61 := GenS5{F0: -2147483648, F1: -1, F2: 42}
                                                    _ = v61
                                                }
                                            } else {
                                                var v62 int64 = 1
                                                _ = v62
                                            }
                                            switch (fi35 + 4294967295) {
                                            case 128:
                                                v5[4294967295] = 4294967295
                                            case 32768:
                                                var v63 uint16 = 255
                                                _ = v63
                                            default:
                                                switch (fi35 << ((65535) & 63)) {
                                                case 4294967295:
                                                    fmt.Println(v36)
                                                }
                                            }
                                            _ = v36
                                        } else {
                                            var v64 bool = (102 < -113)
                                            v6 = v6
                                            _ = v64
                                        }
                                        var v65 uint = 4294967295
                                        _ = fi35
                                        _ = v65
                                    }
                                    var v66 bool = (408 == 18446744073709551615)
                                    _ = v66
                                } else {
                                    for fi67 := 0; fi67 < 5; fi67++ {
                                        var v68 int32 = 1
                                        if ((v1 || true) || (v2 < (v2 << ((1) & 63)))) {
                                            v69 := GenS10{F0: "foo bar", F1: 1}
                                            v68 += (v68 | (v68 ^ (-(v68))))
                                            _ = v69
                                        } else {
                                            v70 := GenS11{F0: 435, F1: (v68 << ((4294967295) & 63)), F2: 42}
                                            _ = v70
                                        }
                                        _ = fi67
                                        _ = v68
                                    }
                                    switch ("hello" + v6) {
                                    case "\n\t":
                                        v3 = v3
                                    case "hello":
                                        if v1 {
                                            var v71 uint = 0
                                            _ = v71
                                        } else {
                                            v3 = v3
                                            var v72 int32 = -2147483648
                                            _ = v72
                                        }
                                    }
                                    v3 -= v3
                                }
                                for fi73 := 0; fi73 < 14; fi73++ {
                                    switch fi73 {
                                    case 0:
                                        var v74 float32 = ((0.0 - 1e10) + float32(0.1))
                                        _ = v74
                                    default:
                                        v4 = append(v4, "\n\t")
                                    }
                                    if (32767 == 1) {
                                        v4 = append(v4, "")
                                        var v75 string = "\n\t"
                                        _ = v75
                                    } else {
                                        if (!v1) {
                                            if v1 {
                                                if len(v4) >= 2 {
                                                    v4 = v4[1:len(v4)]
                                                }
                                                v1 = ((((!v1) && (v3 <= 1)) || false) && (!((567 < 0) && (v1 && v1))))
                                                switch ((v6 + v6) + v6) {
                                                case "\n\t":
                                                    v3 -= v3
                                                    if v1 {
                                                        var v76 float64 = (3.14159 / 1.5)
                                                        switch fi29 {
                                                        case -881:
                                                            if (-595 != 9223372036854775807) {
                                                                if v1 {
                                                                    var v77 bool = false
                                                                    switch -2147483648 {
                                                                    case 127:
                                                                        v78 := GenS1{F0: 1}
                                                                        _ = v78
                                                                    case 666:
                                                                        if v1 {
                                                                            switch (v6 + v6) {
                                                                            case "foo bar":
                                                                                if (v6 >= v6) {
                                                                                    switch v6 {
                                                                                    case "a":
                                                                                        _ = len(v4)
                                                                                    default:
                                                                                        if (42 > 0) {
                                                                                            var v79 float32 = 0.1
                                                                                            var v80 uint = 1
                                                                                            v81 := GenS12{F0: 65535}
                                                                                            _ = v79
                                                                                            _ = v80
                                                                                            _ = v81
                                                                                        } else {
                                                                                            var v82 uint16 = 2
                                                                                            fmt.Println(v6, fi29)
                                                                                            _ = v82
                                                                                        }
                                                                                    }
                                                                                    switch v6 {
                                                                                    case "\n\t":
                                                                                        v83 := GenS13{F0: (("hello" + v6) + v6), F1: 367, F2: 1}
                                                                                        _ = v83
                                                                                    case "foo bar":
                                                                                        var v84 int8 = (v2 << ((255) & 63))
                                                                                        _ = v84
                                                                                    case "":
                                                                                        v6 = v6
                                                                                    }
                                                                                } else {
                                                                                    switch v6 {
                                                                                    case "hello":
                                                                                        fmt.Println(v77, fi73)
                                                                                    default:
                                                                                        var v85 int16 = 1
                                                                                        _ = v85
                                                                                    }
                                                                                    fmt.Println(v76, v3)
                                                                                }
                                                                            case "":
                                                                                v2 -= v2
                                                                                v4 = append(v4, v6)
                                                                                if len(v4) >= 2 {
                                                                                    v4 = v4[1:len(v4)]
                                                                                }
                                                                            case "\n\t":
                                                                                var v86 int16 = 42
                                                                                _ = v86
                                                                            }
                                                                        } else {
                                                                            var v87 int32 = 48
                                                                            fmt.Println(v2)
                                                                            _ = v87
                                                                        }
                                                                    case 0:
                                                                        if v1 {
                                                                            var v88 int32 = 0
                                                                            _ = v88
                                                                        } else {
                                                                            v3 += v3
                                                                        }
                                                                    }
                                                                    switch (v6 + ("hello" + ("a" + v6))) {
                                                                    case "\n\t":
                                                                        var v89 float64 = float64(1.5)
                                                                        _ = v89
                                                                    case "foo bar":
                                                                        var v90 uint64 = genFn1(false, (v2 - (v2 << ((0) & 63))), 9223372036854775807)
                                                                        _ = v90
                                                                    }
                                                                    if false {
                                                                        switch (v6 + ("a" + (v6 + "hello"))) {
                                                                        case "foo bar":
                                                                            var v91 uint16 = 255
                                                                            _ = v91
                                                                        default:
                                                                            if ((v1 || v77) || true) {
                                                                                v92 := v5[1]
                                                                                if (v77 || v77) {
                                                                                    var v93 int8 = (v2 / (((^(v2))) | 1))
                                                                                    var v94 int8 = (^(v2))
                                                                                    if ((v94 & -1) == (v93 >> ((255) & 63))) {
                                                                                        var v95 uint16 = 902
                                                                                        _ = v95
                                                                                    }
                                                                                    _ = v93
                                                                                    _ = v94
                                                                                } else {
                                                                                    var v96 uint16 = genFn0()
                                                                                    fmt.Println(v3)
                                                                                    _ = v96
                                                                                }
                                                                                v1 = (genFn0() > 557)
                                                                                _ = v92
                                                                            }
                                                                        }
                                                                        if (!(!true)) {
                                                                            v97 := GenS14{F0: v2}
                                                                            fmt.Println(v2)
                                                                            _ = v97
                                                                        } else {
                                                                            v5[65535] = 0
                                                                            if ((true && (v3 <= (v3 - v3))) || false) {
                                                                                var v98 int8 = 42
                                                                                v2 += (v2 + 106)
                                                                                v1 = (65535 == 65535)
                                                                                _ = v98
                                                                            }
                                                                            if v1 {
                                                                                var v99 int32 = 2147483647
                                                                                switch v6 {
                                                                                case "a":
                                                                                    var v100 uint32 = 1
                                                                                    _ = v100
                                                                                case "hello":
                                                                                    fmt.Println(v3, fi73)
                                                                                    if (!v1) {
                                                                                        var v101 float32 = 1.5
                                                                                        _ = v101
                                                                                    } else {
                                                                                        v102 := GenS15{F0: 1}
                                                                                        v2 = v2
                                                                                        _ = v102
                                                                                    }
                                                                                }
                                                                                _ = v99
                                                                            }
                                                                            if v77 {
                                                                                switch 707 {
                                                                                case 32767:
                                                                                    switch (v6 + v6) {
                                                                                    case "\n\t":
                                                                                        var v103 float64 = 3.14159
                                                                                        _ = v103
                                                                                    case "hello":
                                                                                        v77 = ((42 > 700) && true)
                                                                                    }
                                                                                }
                                                                            } else {
                                                                                if len(v4) >= 2 {
                                                                                    v4 = v4[1:len(v4)]
                                                                                }
                                                                                var v104 int64 = 9223372036854775807
                                                                                _ = v104
                                                                            }
                                                                            v6 = v6
                                                                        }
                                                                        var v105 string = ("hello" + v6)
                                                                        _ = v105
                                                                    } else {
                                                                        var v106 int32 = -198
                                                                        v107 := GenS16{F0: 527, F1: 2}
                                                                        if v1 {
                                                                            fmt.Println(v77, fi73, fi73)
                                                                            var v108 int32 = (v106 * v106)
                                                                            fmt.Println(v2, v2, v106)
                                                                            _ = v108
                                                                        } else {
                                                                            switch fi73 {
                                                                            case -2147483648:
                                                                                _ = len(v4)
                                                                            default:
                                                                                v77 = false
                                                                            }
                                                                            fmt.Println(v3)
                                                                            var v109 int8 = v2
                                                                            v77 = (!false)
                                                                            _ = v109
                                                                        }
                                                                        _ = v106
                                                                        _ = v107
                                                                    }
                                                                    var v110 int = (fi73 - (-(fi73)))
                                                                    _ = v110
                                                                    _ = v77
                                                                }
                                                            }
                                                        }
                                                        v111 := GenS6{F0: 1, F1: genFn0(), F2: -1}
                                                        var v112 int64 = 9223372036854775807
                                                        var v113 string = (v6 + ((v6 + v6) + (("\n\t" + v6) + (("\n\t" + v6) + ""))))
                                                        _ = v111
                                                        _ = v112
                                                        _ = v113
                                                        _ = v76
                                                    } else {
                                                        v114 := GenS17{F0: (v2 & v2)}
                                                        if v1 {
                                                            v3 -= v3
                                                            var v115 float64 = 0.0
                                                            fmt.Println(v3, fi73, v115)
                                                            v115 = 3.14159
                                                            _ = v115
                                                        }
                                                        _ = v114
                                                    }
                                                case "a":
                                                    if v1 {
                                                        if (!((true || v1) || v1)) {
                                                            var v116 int = (fi73 << ((4294967295) & 63))
                                                            v4 = append(v4, ((((v6 + v6) + "hello") + ("hello" + ("foo bar" + ("foo bar" + v6)))) + (((v6 + "foo bar") + v6) + v6)))
                                                            _ = v116
                                                        } else {
                                                            v3 = v3
                                                            fmt.Println(v2, v2)
                                                            var v117 uint32 = 2147483647
                                                            _ = v117
                                                        }
                                                        if (42 == 18446744073709551615) {
                                                            var v118 float64 = 1e10
                                                            v4 = append(v4, (v6 + v6))
                                                            _ = v118
                                                        } else {
                                                            var v119 int8 = 0
                                                            fmt.Println(v1)
                                                            _ = v119
                                                        }
                                                        if (-196 >= 875) {
                                                            if false {
                                                                var v120 int8 = v2
                                                                if len(v4) > 0 {
                                                                    v4[0] = (((("hello" + "") + v6) + v6) + ((v6 + "hello") + v6))
                                                                }
                                                                switch (fi73 * fi29) {
                                                                case 181:
                                                                    if false {
                                                                        fmt.Println(v1, v120)
                                                                    } else {
                                                                        if (65535 != 65535) {
                                                                            var v121 int32 = -1
                                                                            v4 = append(v4, ("" + ""))
                                                                            if v1 {
                                                                                var v122 uint8 = (v3 << ((255) & 63))
                                                                                if (4294967295 != 486) {
                                                                                    v120 = v120
                                                                                    if v1 {
                                                                                        var v123 float32 = 0.1
                                                                                        _ = v123
                                                                                    }
                                                                                    if v1 {
                                                                                        fmt.Println(v2, v122)
                                                                                        if (!((!v1) && v1)) {
                                                                                            var v124 uint32 = 426
                                                                                            var v125 uint8 = v3
                                                                                            _ = v124
                                                                                            _ = v125
                                                                                        }
                                                                                    }
                                                                                }
                                                                                fmt.Println(v6)
                                                                                _ = v122
                                                                            } else {
                                                                                delete(v5, 2)
                                                                                var v126 string = v6
                                                                                var v127 uint = 0
                                                                                _ = v126
                                                                                _ = v127
                                                                            }
                                                                            _ = v121
                                                                        }
                                                                        v128 := GenS1{F0: 0}
                                                                        switch fi73 {
                                                                        case 32768:
                                                                            fmt.Println(v3, fi29)
                                                                        case 700:
                                                                            if v1 {
                                                                                if (genFn0() != 65535) {
                                                                                    _ = len(v5)
                                                                                } else {
                                                                                    if true {
                                                                                        if (v1 || ((!((!(42 < 42)) && v1)) && (float64((((((float32((0.1 - -2.25)) - 0.0) / 3.14159) + 0.1) / ((float32(1.5) - float32(1.5)) - 3.14159)) + (-2.25 / 1e10))) != -2.25))) {
                                                                                            var v129 int32 = 2147483647
                                                                                            fmt.Println(v3)
                                                                                            var v130 int64 = -829
                                                                                            _ = v129
                                                                                            _ = v130
                                                                                        }
                                                                                    } else {
                                                                                        delete(v5, 2)
                                                                                        v6 = v6
                                                                                    }
                                                                                }
                                                                                fmt.Println(v120)
                                                                            }
                                                                        default:
                                                                            switch v6 {
                                                                            case "hello":
                                                                                if (!true) {
                                                                                    fmt.Println(v120)
                                                                                }
                                                                            case "foo bar":
                                                                                switch ((v6 + (v6 + v6)) + v6) {
                                                                                case "\n\t":
                                                                                    v6 = v6
                                                                                default:
                                                                                    var v131 uint64 = 0
                                                                                    _ = v131
                                                                                }
                                                                            }
                                                                        }
                                                                        _ = v128
                                                                    }
                                                                default:
                                                                    v3 = (v3 - v3)
                                                                }
                                                                _ = v120
                                                            }
                                                        } else {
                                                            switch v6 {
                                                            case "\n\t":
                                                                v132 := GenS1{F0: 42}
                                                                _ = v132
                                                            case "a":
                                                                v133 := v5[1]
                                                                _ = v133
                                                            case "":
                                                                v2 += v2
                                                            default:
                                                                if v1 {
                                                                    v134 := GenS1{F0: -1}
                                                                    switch (((v6 + v6) + v6) + ("foo bar" + v6)) {
                                                                    case "\n\t":
                                                                        v3 = 42
                                                                    default:
                                                                        v135 := GenS12{F0: 963}
                                                                        _ = v135
                                                                    }
                                                                    _ = v134
                                                                } else {
                                                                    v1 = (((v1 && v1) && v1) || (1 <= genFn0()))
                                                                    fmt.Println(v2, v6)
                                                                }
                                                            }
                                                        }
                                                        if len(v4) >= 2 {
                                                            v4 = v4[1:len(v4)]
                                                        }
                                                        var v136 float32 = (((3.14159 + 0.0) / (1.5 / 0.1)) / (1e10 + -2.25))
                                                        var v137 float64 = (1.5 - 0.1)
                                                        _ = v136
                                                        _ = v137
                                                    } else {
                                                        switch v6 {
                                                        case "a":
                                                            fmt.Println(v1)
                                                        }
                                                        v138 := GenS18{F0: (^(fi73)), F1: (-2147483648 != -1)}
                                                        _ = v138
                                                    }
                                                }
                                            }
                                        } else {
                                            if (((v3 == (v3 + (^(v3)))) && v1) || v1) {
                                                var v139 float32 = (((0.0 / (0.0 + -2.25)) - (((1.5 - (float32(((-2.25 / 1.5) + ((3.14159 + (0.0 + float64(0.0))) - 3.14159))) * 1e10)) - 1e10) - 0.0)) / 1.0)
                                                _ = v139
                                            } else {
                                                if (((("a" + v6) + "\n\t") < "") || (74 >= -1)) {
                                                    switch v6 {
                                                    case "":
                                                        v1 = (42 <= -2147483648)
                                                    case "foo bar":
                                                        var v140 uint8 = (v3 << ((42) & 63))
                                                        _ = v140
                                                    case "hello":
                                                        var v141 int8 = v2
                                                        _ = v141
                                                    }
                                                    var v142 int16 = 42
                                                    var v143 int64 = 0
                                                    v1 = ((v143 > (v143 | 9223372036854775807)) || (0.0 <= (float32(0.1) * (3.14159 / float32(1.5)))))
                                                    _ = v142
                                                    _ = v143
                                                }
                                                v144 := GenS12{F0: 4294967295}
                                                if ("" < v6) {
                                                    var v145 bool = v1
                                                    var v146 float64 = 0.0
                                                    _ = v145
                                                    _ = v146
                                                }
                                                _ = v144
                                            }
                                        }
                                        var v147 int64 = int64(fi29)
                                        _ = v147
                                    }
                                    var v148 int32 = int32(v3)
                                    _ = fi73
                                    _ = v148
                                }
                                _ = fi29
                            }
                            v149 := GenS19{F0: -128}
                            var v150 uint16 = 65535
                            _ = v149
                            _ = v150
                        }
                    } else {
                        if ((v3 - 103) > 1) {
                            for fi151 := 0; fi151 < 9; fi151++ {
                                v6 = v6
                                fmt.Println(fi151, v1)
                                var v152 float64 = -2.25
                                _ = fi151
                                _ = v152
                            }
                            if v1 {
                                if (((1e10 > 0.0) || v1) && ((0.1 != ((1e10 + 0.1) * 0.1)) && v1)) {
                                    if false {
                                        var v153 int64 = 655
                                        _ = v153
                                    } else {
                                        var v154 int8 = v2
                                        var v155 bool = (-2147483648 <= 2147483647)
                                        v154 -= -1
                                        _ = v154
                                        _ = v155
                                    }
                                } else {
                                    if (v6 >= ("hello" + v6)) {
                                        v156 := GenS8{F0: v6, F1: -342}
                                        switch -9223372036854775808 {
                                        case -9223372036854775808:
                                            v157 := v5[831]
                                            _ = v157
                                        case 32768:
                                            var v158 int = 128
                                            _ = v158
                                        }
                                        _ = v156
                                    } else {
                                        if (!((!v1) && v1)) {
                                            fmt.Println(v3, v1, v3)
                                            for fi159 := 0; fi159 < 6; fi159++ {
                                                var v160 uint = 42
                                                fmt.Println(v2, v6, v6)
                                                _ = fi159
                                                _ = v160
                                            }
                                            if (-2147483648 != 1) {
                                                if false {
                                                    for fi161 := 0; fi161 < 4; fi161++ {
                                                        var v162 int = int(v3)
                                                        v163 := GenS20{F0: (("hello" + (("\n\t" + (("a" + "foo bar") + v6)) + v6)) + "\n\t"), F1: fi161}
                                                        _ = fi161
                                                        _ = v162
                                                        _ = v163
                                                    }
                                                    v2 -= (v2 * v2)
                                                } else {
                                                    switch v6 {
                                                    case "foo bar":
                                                        fmt.Println(v2)
                                                    default:
                                                        for fi164 := 0; fi164 < 5; fi164++ {
                                                            var v165 uint8 = (v3 + v3)
                                                            delete(v5, 4294967295)
                                                            delete(v5, 602)
                                                            _ = fi164
                                                            _ = v165
                                                        }
                                                    }
                                                    if (v2 < v2) {
                                                        var v166 uint64 = 4294967295
                                                        _ = v166
                                                    }
                                                }
                                                var v167 string = (v6 + ("a" + v6))
                                                var v168 uint = 255
                                                _ = v167
                                                _ = v168
                                            } else {
                                                for fi169 := 0; fi169 < 3; fi169++ {
                                                    for fi170 := 0; fi170 < 1; fi170++ {
                                                        var v171 string = (("foo bar" + ((v6 + "hello") + v6)) + ("" + v6))
                                                        if (18446744073709551615 <= 182) {
                                                            v1 = (2 >= 726)
                                                            if v1 {
                                                                v171 = (v171 + "foo bar")
                                                                v172 := GenS21{F0: -329, F1: 568}
                                                                _ = v172
                                                            } else {
                                                                v2 += 0
                                                                if v1 {
                                                                    v173 := GenS22{F0: (v1 && v1)}
                                                                    var v174 int64 = 9223372036854775807
                                                                    var v175 int32 = 1
                                                                    _ = v173
                                                                    _ = v174
                                                                    _ = v175
                                                                } else {
                                                                    var v176 float32 = 1e10
                                                                    v177 := GenS16{F0: 395, F1: 108}
                                                                    _ = v176
                                                                    _ = v177
                                                                }
                                                            }
                                                        }
                                                        _ = fi170
                                                        _ = v171
                                                    }
                                                    var v178 int16 = -1
                                                    _ = fi169
                                                    _ = v178
                                                }
                                                fmt.Println(v1, v6)
                                                var v179 int64 = 42
                                                _ = v179
                                            }
                                        } else {
                                            v180 := GenS23{F0: 0, F1: 0, F2: 1}
                                            var v181 int16 = 1
                                            _ = v180
                                            _ = v181
                                        }
                                    }
                                    switch v6 {
                                    case "a":
                                        switch v6 {
                                        case "\n\t":
                                            if v1 {
                                                var v182 uint64 = 227
                                                v5[0] = 42
                                                var v183 uint8 = 100
                                                _ = v182
                                                _ = v183
                                            } else {
                                                v2 -= (v2 - v2)
                                                if (v1 && true) {
                                                    if (2147483647 < -9223372036854775808) {
                                                        v1 = v1
                                                        for fi184 := 0; fi184 < 10; fi184++ {
                                                            fmt.Println(v3)
                                                            for fi185 := 0; fi185 < 13; fi185++ {
                                                                v3 = v3
                                                                _ = fi185
                                                            }
                                                            _ = fi184
                                                        }
                                                    } else {
                                                        switch v6 {
                                                        case "a":
                                                            v186 := GenS18{F0: -9223372036854775808, F1: ((801 < uint(v3)) || v1)}
                                                            _ = v186
                                                        case "foo bar":
                                                            v187 := GenS14{F0: v2}
                                                            _ = v187
                                                        }
                                                        for fi188 := 0; fi188 < 11; fi188++ {
                                                            v189 := GenS22{F0: (0.0 == ((1e10 + (0.1 - 0.0)) * 1e10))}
                                                            var v190 int16 = 1000
                                                            _ = fi188
                                                            _ = v189
                                                            _ = v190
                                                        }
                                                        v191 := v5[4294967295]
                                                        _ = v191
                                                    }
                                                } else {
                                                    for fi192 := 0; fi192 < 2; fi192++ {
                                                        var v193 float64 = 0.0
                                                        _ = fi192
                                                        _ = v193
                                                    }
                                                    v1 = (!((false && (v1 && v1)) && (!(0 <= -376))))
                                                    if len(v4) >= 2 {
                                                        v4 = v4[1:len(v4)]
                                                    }
                                                }
                                            }
                                        default:
                                            if (-2147483648 > 163) {
                                                switch 128 {
                                                case 2147483647:
                                                    var v194 int64 = 0
                                                    _ = v194
                                                default:
                                                    var v195 int16 = -32768
                                                    _ = v195
                                                }
                                                if v1 {
                                                    v196 := GenS7{F0: 0, F1: 2147483647, F2: 880}
                                                    if ((0.0 + 1.5) < 1.5) {
                                                        var v197 int16 = 32767
                                                        _ = v197
                                                    } else {
                                                        switch v6 {
                                                        case "":
                                                            switch -322 {
                                                            case -2147483648:
                                                                var v198 float32 = 0.1
                                                                _ = v198
                                                            case 1:
                                                                for fi199 := 0; fi199 < 14; fi199++ {
                                                                    var v200 int8 = (^(v2))
                                                                    var v201 uint8 = (v3 / (((v3 + (v3 & (v3 * (^(v3)))))) | 1))
                                                                    v202 := GenS6{F0: 923, F1: 2, F2: -41}
                                                                    _ = fi199
                                                                    _ = v200
                                                                    _ = v201
                                                                    _ = v202
                                                                }
                                                            case 743:
                                                                var v203 uint8 = v3
                                                                _ = v203
                                                            }
                                                        }
                                                    }
                                                    if len(v4) > 0 {
                                                        v4[0] = ("" + (v6 + "hello"))
                                                    }
                                                    _ = v196
                                                }
                                            }
                                        }
                                    default:
                                        for fi204 := 0; fi204 < 19; fi204++ {
                                            if v1 {
                                                var v205 uint32 = uint32(v3)
                                                _ = v205
                                            }
                                            _ = fi204
                                        }
                                    }
                                    v206 := GenS24{F0: (v2 * (v2 & v2))}
                                    _ = v206
                                }
                                if (v3 == v3) {
                                    if (!(false || (-128 <= v2))) {
                                        for fi207 := 0; fi207 < 16; fi207++ {
                                            v208 := GenS17{F0: -88}
                                            _ = fi207
                                            _ = v208
                                        }
                                        switch 4294967295 {
                                        case -9223372036854775808:
                                            v2 = (v2 << ((2) & 63))
                                        case 32767:
                                            fmt.Println(v2, v3)
                                        case 4294967295:
                                            if (205 <= 0) {
                                                switch v6 {
                                                case "foo bar":
                                                    switch 215 {
                                                    case 70:
                                                        v209 := GenS25{F0: 127, F1: 255}
                                                        _ = v209
                                                    }
                                                default:
                                                    var v210 int = -9223372036854775808
                                                    _ = v210
                                                }
                                            }
                                        default:
                                            v2 = v2
                                        }
                                        if v1 {
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                            if len(v4) > 0 {
                                                v4[0] = ("" + ("hello" + "a"))
                                            }
                                            switch v6 {
                                            case "a":
                                                v2 = v2
                                            default:
                                                fmt.Println(v6, v3)
                                            }
                                        }
                                    }
                                } else {
                                    if v1 {
                                        var v211 int64 = 715
                                        var v212 uint8 = (v3 % ((v3) | 1))
                                        _ = v211
                                        _ = v212
                                    }
                                    if ((!(!v1)) || v1) {
                                        var v213 int32 = 2147483647
                                        for fi214 := 0; fi214 < 15; fi214++ {
                                            var v215 uint8 = 42
                                            if false {
                                                for fi216 := 0; fi216 < 2; fi216++ {
                                                    if len(v4) >= 2 {
                                                        v4 = v4[1:len(v4)]
                                                    }
                                                    _ = fi216
                                                }
                                            } else {
                                                for fi217 := 0; fi217 < 7; fi217++ {
                                                    v213 = v213
                                                    _ = fi217
                                                }
                                                var v218 int64 = 42
                                                _ = v218
                                            }
                                            var v219 uint8 = (v215 % (((^(v3))) | 1))
                                            _ = fi214
                                            _ = v215
                                            _ = v219
                                        }
                                        _ = v213
                                    } else {
                                        for fi220 := 0; fi220 < 1; fi220++ {
                                            fmt.Println(v6, v1, v6)
                                            fmt.Println(v1, v2)
                                            v4 = append(v4, v6)
                                            _ = fi220
                                        }
                                        v4 = append(v4, "")
                                        if (-716 != 0) {
                                            switch ((v6 + ("\n\t" + v6)) + (v6 + "")) {
                                            case "":
                                                v221 := GenS17{F0: v2}
                                                _ = v221
                                            case "\n\t":
                                                v2 = v2
                                            default:
                                                if false {
                                                    var v222 int16 = 1000
                                                    for fi223 := 0; fi223 < 1; fi223++ {
                                                        var v224 float32 = float32(0.1)
                                                        _ = fi223
                                                        _ = v224
                                                    }
                                                    _ = v222
                                                }
                                            }
                                            var v225 uint8 = v3
                                            _ = v225
                                        }
                                    }
                                }
                            } else {
                                if ((v3 > (v3 | (^(v3)))) || ("hello" == ("foo bar" + ((v6 + "") + ("hello" + v6))))) {
                                    var v226 int64 = -1
                                    _ = v226
                                } else {
                                    var v227 uint32 = 2147483648
                                    if (((-(v2)) < (v2 / ((0) | 1))) || ((-123 < 0) || v1)) {
                                        var v228 uint32 = 0
                                        v6 = v6
                                        _ = v228
                                    } else {
                                        v2 = (v2 * v2)
                                        if (511 > 18446744073709551615) {
                                            fmt.Println(v3, v1)
                                            if v1 {
                                                if false {
                                                    if len(v4) >= 2 {
                                                        v4 = v4[1:len(v4)]
                                                    }
                                                    if (true || false) {
                                                        var v229 float64 = (((0.0 / 1e10) + 1.5) - 1e10)
                                                        if v1 {
                                                            var v230 int = -1
                                                            _ = v230
                                                        } else {
                                                            v5[v227] = v227
                                                        }
                                                        v6 = (v6 + ("\n\t" + v6))
                                                        _ = v229
                                                    }
                                                }
                                                for fi231 := 0; fi231 < 6; fi231++ {
                                                    switch v6 {
                                                    case "a":
                                                        v227 += v227
                                                    case "\n\t":
                                                        fmt.Println(fi231, v3, v3)
                                                    case "":
                                                        var v232 bool = (v1 || (967 == 65535))
                                                        _ = v232
                                                    default:
                                                        v3 = v3
                                                    }
                                                    _ = fi231
                                                }
                                                var v233 uint16 = 335
                                                _ = v233
                                            }
                                            switch -2147483648 {
                                            case 1:
                                                fmt.Println(v227, v3, v1)
                                            case 32768:
                                                if ((0.0 - 1.5) >= 1.5) {
                                                    for fi234 := 0; fi234 < 1; fi234++ {
                                                        switch -790 {
                                                        case -2147483648:
                                                            v4 = append(v4, v6)
                                                        case 32767:
                                                            var v235 uint = 42
                                                            _ = v235
                                                        case 991:
                                                            v236 := GenS14{F0: (v2 >> ((526) & 63))}
                                                            _ = v236
                                                        default:
                                                            var v237 uint16 = 255
                                                            _ = v237
                                                        }
                                                        for fi238 := 0; fi238 < 6; fi238++ {
                                                            var v239 float32 = (0.1 / 1.0)
                                                            _ = fi238
                                                            _ = v239
                                                        }
                                                        var v240 uint32 = (v227 >> ((1) & 63))
                                                        _ = fi234
                                                        _ = v240
                                                    }
                                                    var v241 int32 = -2147483648
                                                    _ = v241
                                                } else {
                                                    fmt.Println(v6)
                                                    var v242 float64 = 0.1
                                                    _ = v242
                                                }
                                            default:
                                                switch 42 {
                                                case -9223372036854775808:
                                                    var v243 uint8 = v3
                                                    _ = v243
                                                default:
                                                    switch 247 {
                                                    case 127:
                                                        v244 := GenS14{F0: 58}
                                                        _ = v244
                                                    default:
                                                        if (255 <= 845) {
                                                            if (false || v1) {
                                                                var v245 uint8 = (v3 - v3)
                                                                _ = v245
                                                            }
                                                            v2 = (v2 * v2)
                                                            var v246 bool = v1
                                                            _ = v246
                                                        } else {
                                                            var v247 int64 = -927
                                                            v227 += v227
                                                            for fi248 := 0; fi248 < 17; fi248++ {
                                                                v2 = (v2 % (((v2 * v2)) | 1))
                                                                if false {
                                                                    if v1 {
                                                                        if true {
                                                                            fmt.Println(v227, fi248)
                                                                            var v249 float64 = 0.0
                                                                            v1 = (!(1.5 != 0.1))
                                                                            _ = v249
                                                                        }
                                                                    }
                                                                    var v250 float32 = 1.5
                                                                    _ = v250
                                                                } else {
                                                                    var v251 uint32 = (v227 + v227)
                                                                    v252 := GenS26{F0: (fi248 * fi248), F1: v2}
                                                                    _ = v251
                                                                    _ = v252
                                                                }
                                                                var v253 uint = 18446744073709551615
                                                                _ = fi248
                                                                _ = v253
                                                            }
                                                            _ = v247
                                                        }
                                                    }
                                                }
                                            }
                                        } else {
                                            if (v1 && ((2147483647 != v227) && v1)) {
                                                v254 := v5[2147483647]
                                                _ = v254
                                            } else {
                                                if (!(v1 && v1)) {
                                                    delete(v5, v227)
                                                } else {
                                                    v255 := GenS27{F0: 1, F1: int16(v2)}
                                                    if true {
                                                        v2 -= v2
                                                    }
                                                    _ = v255
                                                }
                                            }
                                        }
                                    }
                                    v3 -= 155
                                    _ = v227
                                }
                                v256 := GenS8{F0: (v6 + v6), F1: -9223372036854775808}
                                fmt.Println(v6, v3)
                                _ = v256
                            }
                            switch 42 {
                            case 32767:
                                v1 = false
                            case -837:
                                fmt.Println(v2, v1, v1)
                            default:
                                for fi257 := 0; fi257 < 12; fi257++ {
                                    if (v2 < v2) {
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                        var v258 uint64 = 42
                                        _ = v258
                                    } else {
                                        v1 = (v1 || ((("a" + "a") < "a") || (42 != 2)))
                                        if (!(true && v1)) {
                                            v3 = v3
                                            for fi259 := 0; fi259 < 13; fi259++ {
                                                if v1 {
                                                    switch v6 {
                                                    case "a":
                                                        v3 -= v3
                                                    case "hello":
                                                        var v260 float32 = -2.25
                                                        _ = v260
                                                    case "foo bar":
                                                        v261 := GenS7{F0: -9223372036854775808, F1: 42, F2: 18446744073709551615}
                                                        _ = v261
                                                    default:
                                                        switch 9223372036854775807 {
                                                        case 9223372036854775807:
                                                            var v262 bool = (!v1)
                                                            _ = v262
                                                        case 128:
                                                            v3 -= v3
                                                        }
                                                    }
                                                }
                                                if v1 {
                                                    if v1 {
                                                        fmt.Println(v1)
                                                        var v263 int32 = -787
                                                        if (float32(((0.0 * 1e10) + (1.5 / 1.5))) != -2.25) {
                                                            v263 -= -916
                                                            var v264 uint64 = genFn1(true, v2, int64(fi259))
                                                            _ = v264
                                                        } else {
                                                            fmt.Println(v1, fi259)
                                                        }
                                                        _ = v263
                                                    } else {
                                                        v265 := GenS24{F0: v2}
                                                        var v266 int32 = 161
                                                        _ = v265
                                                        _ = v266
                                                    }
                                                }
                                                if len(v4) >= 2 {
                                                    v4 = v4[1:len(v4)]
                                                }
                                                _ = fi259
                                            }
                                        }
                                    }
                                    switch 4294967295 {
                                    case -948:
                                        fmt.Println(fi257)
                                    case 128:
                                        v267 := GenS22{F0: v1}
                                        _ = v267
                                    case 0:
                                        if (float32(1e10) == -2.25) {
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                            var v268 int64 = 228
                                            _ = v268
                                        }
                                    default:
                                        v5[2] = 346
                                    }
                                    _ = fi257
                                }
                            }
                        }
                        if v1 {
                            v5[2147483648] = 65535
                            if v1 {
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                                delete(v5, 0)
                                switch -9223372036854775808 {
                                case 32767:
                                    var v269 int8 = v2
                                    _ = v269
                                case 127:
                                    fmt.Println(v2)
                                default:
                                    v1 = true
                                }
                            }
                        }
                    }
                    var v270 bool = false
                    _ = v270
                }
            } else {
                for fi271 := 0; fi271 < 5; fi271++ {
                    v6 = v6
                    v1 = v1
                    _ = fi271
                }
            }
        default:
            if true {
                var v272 int = 4294967295
                fmt.Println(v2, v272)
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                _ = v272
            }
        }
    }
    if false {
        if true {
            v1 = true
            v2 = 1
        }
        switch 4294967295 {
        case -1:
            var v273 uint32 = 42
            _ = v273
        case 4294967295:
            v5[2147483648] = 65535
        default:
            for fi274 := 0; fi274 < 11; fi274++ {
                for fi275 := 0; fi275 < 15; fi275++ {
                    v6 = (v6 + v6)
                    var v276 uint32 = 2147483648
                    _ = fi275
                    _ = v276
                }
                v3 -= v3
                if (v2 == v2) {
                    v1 = ((v2 ^ (v2 | v2)) == (v2 & 127))
                    for fi277 := 0; fi277 < 7; fi277++ {
                        fmt.Println(v6, v1)
                        _ = fi277
                    }
                    var v278 int = 2147483647
                    _ = v278
                }
                _ = fi274
            }
        }
        var v279 float64 = -2.25
        _ = v279
    }
    v280 := v5[2]
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v6 = "\n\t"
    fmt.Println(v280)
    var v281 int32 = 2147483647
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v280:", v280)
    fmt.Println("v281:", v281)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v280
    _ = v281
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
