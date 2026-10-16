package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint32; F1 uint16 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 uint16; F1 int16; F2 int8 }
type GenS7 struct { F0 uint; F1 int64 }
type GenS8 struct { F0 int32; F1 uint; F2 uint }
type GenS9 struct { F0 uint64; F1 uint }
type GenS10 struct { F0 uint; F1 int8 }
type GenS11 struct { F0 uint; F1 int32; F2 uint64 }
type GenS12 struct { F0 int16; F1 int32; F2 int32 }
type GenS13 struct { F0 int8 }
type GenS14 struct { F0 uint64 }
type GenS15 struct { F0 uint; F1 int16 }
type GenS16 struct { F0 int32; F1 int; F2 int16 }
type GenS17 struct { F0 int32; F1 int; F2 int8 }
type GenS18 struct { F0 uint; F1 uint32; F2 int16 }
type GenS19 struct { F0 uint32 }
type GenS20 struct { F0 int32; F1 uint16 }
type GenS21 struct { F0 int8 }
type GenS22 struct { F0 int16; F1 int64; F2 int8 }
type GenS23 struct { F0 int32 }
type GenS24 struct { F0 uint64; F1 int64 }
type GenS25 struct { F0 uint16; F1 int16; F2 int32 }
type GenS26 struct { F0 int8; F1 uint64 }
type GenS27 struct { F0 int16 }
type GenS28 struct { F0 int16; F1 uint32 }
func genFn0(p0 int8, p1 bool, p2 uint) int {
    if p1 {
        if p1 {
            var v7 uint16 = 32767
            var v8 float32 = float32((1e10 - float64((1e10 / 1.0))))
            p2 = (p2 + 1)
            _ = v7
            _ = v8
        } else {
            if p1 {
                fmt.Println(p1, p2)
                p0 = (p0 | p0)
                var v9 uint64 = 2
                _ = v9
            }
            if ((false && p1) || p1) {
                p0 = (p0 / (((p0 >> ((p2) & 63))) | 1))
            }
            p1 = (!(!false))
        }
    } else {
        v10 := []int{}
        var v11 bool = (0.1 <= 1e10)
        _ = v10
        _ = v11
    }
    v12 := []int{9223372036854775807, -9223372036854775808, 42}
    p0 += p0
    var v13 int8 = (p0 % ((p0) | 1))
    if len(v12) > 0 {
        v12[0] = 9223372036854775807
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v13
    return 32768
}
func genFn1(p0 int8, p1 int, p2 uint64) int32 {
    v14 := GenS1{F0: 753, F1: true}
    if true {
        p2 -= (p2 | 4294967295)
        var v15 uint32 = 2147483648
        var v16 uint = 65535
        _ = v15
        _ = v16
    } else {
        v17 := []int{-568, 1, (p1 >> ((0) & 63)), (p1 * p1)}
        for fi18 := 0; fi18 < 17; fi18++ {
            v19 := GenS2{F0: (p2 << ((0) & 63))}
            var v20 int16 = 32767
            p2 += p2
            _ = fi18
            _ = v19
            _ = v20
        }
        p2 = 1
        _ = v17
    }
    var v21 uint = 18446744073709551615
    v21 = 4294967295
    for fi22 := 0; fi22 < 7; fi22++ {
        var v23 int16 = -997
        _ = fi22
        _ = v23
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    _ = v21
    return int32(p0)
}

func main() {
    var v1 uint32 = 373
    var v2 int8 = 1
    var v3 string = (((("\n\t" + "") + ("\n\t" + "foo bar")) + ("\n\t" + "")) + (("" + "a") + ("a" + "hello")))
    var v4 uint32 = (^(v1))
    v5 := make(map[uint32]int)
    var v6 uint8 = 127
    fmt.Println(v6)
    for fi24 := 0; fi24 < 9; fi24++ {
        for fi25 := 0; fi25 < 10; fi25++ {
            v6 += 148
            switch genFn0(1, false, 2) {
            case 128:
                v26 := GenS3{F0: v1, F1: uint16(v1)}
                _ = v26
            case -429:
                v1 = 938
            }
            fmt.Println(v1, v4)
            _ = fi25
        }
        switch v3 {
        case "":
            v27 := GenS2{F0: 0}
            _ = v27
        default:
            var v28 int = fi24
            _ = v28
        }
        for fi29 := 0; fi29 < 5; fi29++ {
            if true {
                v1 = v4
                switch "a" {
                case "\n\t":
                    var v30 int = genFn0(v2, ((v4 % (((v1 + v4)) | 1)) > (v1 << ((18446744073709551615) & 63))), 2)
                    _ = v30
                case "hello":
                    v2 -= 0
                default:
                    v3 = "foo bar"
                }
                if false {
                    v31 := GenS3{F0: v4, F1: 124}
                    switch "" {
                    case "a":
                        if (240 <= 42) {
                            v32 := GenS3{F0: (v1 ^ v1), F1: 0}
                            v33 := GenS2{F0: 42}
                            _ = v32
                            _ = v33
                        } else {
                            if ((((-2.25 / 1.5) / (1e10 + (3.14159 * 1e10))) + 0.1) < 1e10) {
                                if (255 == 18446744073709551615) {
                                    fmt.Println(fi24, v3, v3)
                                    v4 = v4
                                    v4 -= (^(v1))
                                    switch "a" {
                                    case "hello":
                                        v34 := GenS3{F0: v1, F1: 255}
                                        _ = v34
                                    case "foo bar":
                                        if false {
                                            var v35 float32 = 0.1
                                            v35 += (((float32(-2.25) / (-2.25 + v35)) + v35) / 3.14159)
                                            _ = v35
                                        }
                                    default:
                                        var v36 bool = (2147483647 <= -2147483648)
                                        _ = v36
                                    }
                                    if false {
                                        v2 += (v2 + v2)
                                        if (v3 <= (v3 + ((v3 + v3) + (v3 + "\n\t")))) {
                                            v37 := GenS4{F0: 65535}
                                            var v38 int = fi24
                                            var v39 int16 = 32767
                                            _ = v37
                                            _ = v38
                                            _ = v39
                                        } else {
                                            v40 := v5[(^(v4))]
                                            _ = v40
                                        }
                                    } else {
                                        var v41 float32 = 1e10
                                        v6 -= (v6 << ((439) & 63))
                                        _ = v41
                                    }
                                    if false {
                                        v42 := GenS3{F0: (v4 >> ((0) & 63)), F1: 1}
                                        switch (v3 + v3) {
                                        case "a":
                                            v43 := GenS5{F0: v2}
                                            _ = v43
                                        case "foo bar":
                                            switch fi29 {
                                            case 32768:
                                                v3 = "a"
                                            case 0:
                                                v44 := GenS2{F0: 18446744073709551615}
                                                _ = v44
                                            default:
                                                v1 = v1
                                            }
                                        }
                                        v2 += (v2 | (v2 & (v2 + 9)))
                                        _ = v42
                                    }
                                    switch 32767 {
                                    case 128:
                                        var v45 int8 = int8(fi24)
                                        _ = v45
                                    case 9223372036854775807:
                                        v2 = v2
                                    default:
                                        if (42 != 485) {
                                            v46 := GenS6{F0: 42, F1: 1000, F2: 42}
                                            var v47 uint8 = v6
                                            v48 := GenS7{F0: 656, F1: -1}
                                            v49 := GenS8{F0: 1, F1: 18446744073709551615, F2: 77}
                                            switch ((("\n\t" + ("\n\t" + v3)) + v3) + (v3 + "a")) {
                                            case "hello":
                                                v1 += v4
                                            case "foo bar":
                                                v47 = v6
                                            case "":
                                                v50 := []int{9223372036854775807, 0}
                                                v1 = v1
                                                _ = v50
                                            default:
                                                var v51 int16 = -653
                                                _ = v51
                                            }
                                            if true {
                                                var v52 int8 = (v2 | v2)
                                                var v53 int8 = (^(v2))
                                                switch (^(fi24)) {
                                                case -65:
                                                    v54 := []int{fi24, (fi24 / (((fi24 >> ((18446744073709551615) & 63))) | 1))}
                                                    _ = v54
                                                case -616:
                                                    var v55 float64 = ((1e10 + 0.1) + 1e10)
                                                    _ = v55
                                                }
                                                _ = v52
                                                _ = v53
                                            }
                                            _ = v46
                                            _ = v47
                                            _ = v48
                                            _ = v49
                                        } else {
                                            fmt.Println(fi24, v2)
                                            v56 := GenS9{F0: 4294967295, F1: 255}
                                            _ = v56
                                        }
                                    }
                                }
                            } else {
                                if true {
                                    v3 = v3
                                } else {
                                    v57 := GenS10{F0: uint(v6), F1: int8(v1)}
                                    fmt.Println(v4)
                                    _ = v57
                                }
                            }
                        }
                    case "\n\t":
                        v2 -= 24
                        v3 = (v3 + "foo bar")
                        v58 := GenS9{F0: 2, F1: 267}
                        _ = v58
                    default:
                        if false {
                            v2 = v2
                            delete(v5, 100)
                            v6 = (^(v6))
                        } else {
                            fmt.Println(v1)
                            switch ("hello" + (("foo bar" + "\n\t") + v3)) {
                            case "hello":
                                v6 = (v6 % ((1) | 1))
                            case "foo bar":
                                fmt.Println(v4)
                            case "\n\t":
                                v59 := GenS4{F0: 32767}
                                _ = v59
                            default:
                                if true {
                                    switch (fi29 * fi29) {
                                    case 9223372036854775807:
                                        var v60 float64 = 3.14159
                                        _ = v60
                                    case 128:
                                        var v61 float64 = 1.5
                                        switch v3 {
                                        case "":
                                            var v62 int16 = 1000
                                            fmt.Println(v6)
                                            if false {
                                                v63 := GenS11{F0: 65535, F1: 0, F2: 233}
                                                var v64 float32 = (1e10 - -2.25)
                                                switch v3 {
                                                case "a":
                                                    var v65 int32 = -1
                                                    _ = v65
                                                case "\n\t":
                                                    fmt.Println(v2)
                                                }
                                                v66 := GenS10{F0: 2, F1: v2}
                                                _ = v63
                                                _ = v64
                                                _ = v66
                                            }
                                            _ = v62
                                        case "a":
                                            fmt.Println(v2)
                                            v67 := GenS5{F0: 100}
                                            if (!false) {
                                                fmt.Println(fi24, v1)
                                            }
                                            _ = v67
                                        }
                                        switch (fi29 ^ fi29) {
                                        case 128:
                                            fmt.Println(v1)
                                        case -941:
                                            v1 = 2147483648
                                        case -859:
                                            fmt.Println(fi29)
                                        }
                                        _ = v61
                                    case -9223372036854775808:
                                        var v68 int32 = -2147483648
                                        _ = v68
                                    }
                                } else {
                                    v69 := GenS10{F0: 2, F1: v2}
                                    _ = v69
                                }
                            }
                            v70 := GenS10{F0: 690, F1: v2}
                            v6 += 255
                            _ = v70
                        }
                    }
                    if (false && ((false || true) && (1e10 > 3.14159))) {
                        v5[v1] = -103
                        fmt.Println(v1, v3)
                        var v71 int16 = 715
                        _ = v71
                    }
                    _ = v31
                } else {
                    if false {
                        v4 -= 1
                        var v72 float32 = -2.25
                        _ = v72
                    } else {
                        var v73 int = fi24
                        v3 = v3
                        var v74 string = v3
                        _ = v73
                        _ = v74
                    }
                }
                switch v3 {
                case "foo bar":
                    var v75 int = fi29
                    _ = v75
                case "\n\t":
                    var v76 int16 = -1
                    _ = v76
                }
            }
            switch v3 {
            case "\n\t":
                switch v3 {
                case "":
                    v1 += (v1 & (v4 >> ((1) & 63)))
                default:
                    if (!(false && ((!true) || false))) {
                        v4 -= 328
                    }
                }
            case "a":
                var v77 uint16 = 1
                _ = v77
            case "":
                fmt.Println(v3)
            }
            if ((true && ((!(42 == -984)) && false)) && (!((-2147483648 < -950) || (1e10 < 3.14159)))) {
                v78 := []int{genFn0(v2, false, 594)}
                if len(v78) > 0 {
                    v78[0] = (-(fi29))
                }
                var v79 int16 = 1000
                _ = v78
                _ = v79
            } else {
                _ = len(v5)
            }
            _ = fi29
        }
        _ = fi24
    }
    if true {
        var v80 uint8 = (v6 % ((2) | 1))
        if ("" <= v3) {
            for fi81 := 0; fi81 < 10; fi81++ {
                switch ("" + v3) {
                case "":
                    var v82 string = ""
                    _ = v82
                default:
                    switch (v3 + v3) {
                    case "":
                        for fi83 := 0; fi83 < 11; fi83++ {
                            switch (("" + (v3 + v3)) + (v3 + "foo bar")) {
                            case "foo bar":
                                v6 += v6
                                v84 := GenS12{F0: 129, F1: 42, F2: 2147483647}
                                fmt.Println(v6, v2, v4)
                                if false {
                                    if false {
                                        v2 = (v2 ^ v2)
                                        switch fi83 {
                                        case 42:
                                            switch (fi83 & 42) {
                                            case -9223372036854775808:
                                                v85 := GenS8{F0: -544, F1: uint(v80), F2: 65535}
                                                _ = v85
                                            case 9223372036854775807:
                                                v86 := GenS10{F0: 65535, F1: v2}
                                                _ = v86
                                            }
                                        case -1:
                                            if (-1 >= 42) {
                                                var v87 int16 = 32767
                                                var v88 int = fi83
                                                if (971 >= 850) {
                                                    _ = len(v5)
                                                    var v89 int = 32767
                                                    if true {
                                                        switch v3 {
                                                        case "hello":
                                                            v3 = v3
                                                        case "a":
                                                            if (1.5 >= 0.1) {
                                                                v5[v4] = (-(v89))
                                                                if (2 != 2) {
                                                                    var v90 bool = (v4 == v1)
                                                                    if true {
                                                                        var v91 uint8 = (v80 & (v6 / ((128) | 1)))
                                                                        v92 := GenS1{F0: v87, F1: v90}
                                                                        if (1.5 < float32(1.5)) {
                                                                            var v93 uint16 = 65535
                                                                            _ = v93
                                                                        } else {
                                                                            v94 := GenS13{F0: v2}
                                                                            _ = v94
                                                                        }
                                                                        v95 := []int{110}
                                                                        _ = v91
                                                                        _ = v92
                                                                        _ = v95
                                                                    } else {
                                                                        var v96 bool = (18446744073709551615 != 18446744073709551615)
                                                                        v96 = (1e10 == (float64(0.0) / 3.14159))
                                                                        var v97 int32 = -2147483648
                                                                        _ = v96
                                                                        _ = v97
                                                                    }
                                                                    _ = v90
                                                                }
                                                            }
                                                        }
                                                        if (false && true) {
                                                            if (0.1 < (-2.25 * (1.5 * float32(0.1)))) {
                                                                v88 = (^(fi83))
                                                            }
                                                            v98 := []int{genFn0(-49, (!(v4 >= v4)), 431), genFn0(v2, (-9223372036854775808 > -648), 473), fi83, 42}
                                                            switch (fi81 & 127) {
                                                            case 32768:
                                                                v6 -= (v80 | 128)
                                                            case -9223372036854775808:
                                                                if false {
                                                                    if true {
                                                                        _ = len(v98)
                                                                        if ((float64(0.0) * -2.25) > 1.5) {
                                                                            _ = len(v5)
                                                                            delete(v5, 2147483647)
                                                                            var v99 int8 = (^(v2))
                                                                            var v100 string = "hello"
                                                                            var v101 int32 = 281
                                                                            _ = v100
                                                                            _ = v101
                                                                            _ = v99
                                                                        } else {
                                                                            v6 += 255
                                                                            v102 := GenS12{F0: (v87 | -32768), F1: 2147483647, F2: -829}
                                                                            v89 = v89
                                                                            _ = v102
                                                                        }
                                                                        fmt.Println(v3)
                                                                    }
                                                                }
                                                                v89 -= (v88 & fi81)
                                                                var v103 int64 = -377
                                                                _ = v103
                                                            case 9223372036854775807:
                                                                switch 1 {
                                                                case 0:
                                                                    var v104 int32 = 1
                                                                    _ = v104
                                                                case 4294967295:
                                                                    switch v3 {
                                                                    case "\n\t":
                                                                        var v105 bool = false
                                                                        _ = v105
                                                                    case "hello":
                                                                        if ((-(v89)) >= v89) {
                                                                            v2 += (v2 | v2)
                                                                        } else {
                                                                            var v106 float64 = (1e10 - 1e10)
                                                                            var v107 int8 = (v2 ^ v2)
                                                                            _ = v106
                                                                            _ = v107
                                                                        }
                                                                    case "a":
                                                                        fmt.Println(v87)
                                                                    default:
                                                                        v80 = uint8(v4)
                                                                    }
                                                                default:
                                                                    switch v3 {
                                                                    case "foo bar":
                                                                        v108 := GenS6{F0: 76, F1: (v87 + -162), F2: v2}
                                                                        _ = v108
                                                                    case "hello":
                                                                        var v109 float64 = 0.1
                                                                        _ = v109
                                                                    case "a":
                                                                        if false {
                                                                            delete(v5, v4)
                                                                        }
                                                                    default:
                                                                        v87 += (v87 * v87)
                                                                    }
                                                                    v88 = v88
                                                                }
                                                            }
                                                            _ = v98
                                                        }
                                                        fmt.Println(fi81)
                                                    } else {
                                                        v110 := []int{9223372036854775807, (^(fi81)), int(v2)}
                                                        _ = v110
                                                    }
                                                    _ = v89
                                                } else {
                                                    var v111 int64 = -1
                                                    _ = v111
                                                }
                                                _ = v87
                                                _ = v88
                                            } else {
                                                switch (fi81 / ((genFn0(127, false, 0)) | 1)) {
                                                case 128:
                                                    v112 := []int{genFn0((v2 << ((0) & 63)), true, 276), fi81, genFn0(v2, ((((v3 + "hello") <= v3) || (false || false)) && (false && (42 != 1))), 0)}
                                                    _ = v112
                                                case -860:
                                                    switch 42 {
                                                    case -1:
                                                        v113 := []int{genFn0((v2 << ((392) & 63)), ((fi81 >= genFn0(v2, ((fi83 != (fi83 >> ((1) & 63))) || (!true)), 42)) && false), uint(v80))}
                                                        var v114 int32 = 0
                                                        if (!true) {
                                                            v115 := GenS14{F0: 18446744073709551615}
                                                            switch (v3 + "\n\t") {
                                                            case "hello":
                                                                switch ((v3 + v3) + v3) {
                                                                case "":
                                                                    v114 += (v114 - (-(v114)))
                                                                default:
                                                                    if (408 != 2) {
                                                                        v116 := GenS15{F0: 18446744073709551615, F1: 1000}
                                                                        _ = v116
                                                                    }
                                                                    if (-2147483648 == -1) {
                                                                        switch fi83 {
                                                                        case -9223372036854775808:
                                                                            v117 := GenS7{F0: 0, F1: -511}
                                                                            _ = v117
                                                                        default:
                                                                            v118 := GenS8{F0: (^(v114)), F1: 255, F2: 4294967295}
                                                                            _ = v118
                                                                        }
                                                                        v2 = -1
                                                                        v4 -= v1
                                                                    }
                                                                    var v119 int64 = 42
                                                                    v120 := GenS9{F0: 1, F1: 172}
                                                                    v121 := v5[2147483648]
                                                                    _ = v119
                                                                    _ = v120
                                                                    _ = v121
                                                                }
                                                            case "":
                                                                switch 375 {
                                                                case 2147483647:
                                                                    var v122 string = v3
                                                                    _ = v122
                                                                case 1:
                                                                    v123 := GenS11{F0: 65535, F1: genFn1(v2, genFn0(34, false, 42), 18446744073709551615), F2: 4294967295}
                                                                    v80 -= (v6 + v80)
                                                                    _ = v123
                                                                }
                                                            }
                                                            if false {
                                                                fmt.Println(fi81, v4)
                                                                if ((((v3 + ((((v3 + (v3 + (v3 + ("\n\t" + "")))) + (v3 + "\n\t")) + "\n\t") + ((("hello" + v3) + v3) + "hello"))) + v3) + (v3 + v3)) <= v3) {
                                                                    v4 += (v1 - 117)
                                                                    fmt.Println(v4, v3)
                                                                }
                                                                v124 := GenS3{F0: 2147483648, F1: 255}
                                                                _ = v124
                                                            }
                                                            _ = v115
                                                        }
                                                        _ = v113
                                                        _ = v114
                                                    }
                                                case 0:
                                                    if false {
                                                        v125 := v5[(v4 + (v4 << ((65535) & 63)))]
                                                        _ = v125
                                                    } else {
                                                        if false {
                                                            if (2 > 1) {
                                                                v126 := GenS16{F0: 1, F1: fi81, F2: 0}
                                                                switch (fi81 - (fi83 ^ (-(fi83)))) {
                                                                case 0:
                                                                    v127 := v5[v4]
                                                                    _ = v127
                                                                default:
                                                                    if (1 <= 424) {
                                                                        delete(v5, (v4 - v1))
                                                                    }
                                                                    if (1 == -699) {
                                                                        if true {
                                                                            var v128 uint32 = (v4 + 1)
                                                                            var v129 float64 = ((3.14159 + 1e10) + 1e10)
                                                                            _ = v128
                                                                            _ = v129
                                                                        }
                                                                        if true {
                                                                            v130 := GenS17{F0: -2147483648, F1: fi83, F2: v2}
                                                                            var v131 uint8 = (v6 << ((658) & 63))
                                                                            _ = v130
                                                                            _ = v131
                                                                        }
                                                                        var v132 string = ("\n\t" + v3)
                                                                        _ = v132
                                                                    }
                                                                    fmt.Println(v6, v2)
                                                                    var v133 string = "\n\t"
                                                                    _ = v133
                                                                }
                                                                v3 = "foo bar"
                                                                v134 := GenS2{F0: 4294967295}
                                                                _ = v126
                                                                _ = v134
                                                            } else {
                                                                v3 = v3
                                                                v2 -= (^(v2))
                                                                v3 = v3
                                                                var v135 int8 = 122
                                                                _ = v135
                                                            }
                                                            var v136 int = fi83
                                                            _ = v136
                                                        } else {
                                                            if (true || false) {
                                                                v137 := GenS18{F0: 0, F1: 2147483648, F2: 42}
                                                                var v138 float32 = (1.5 + 1e10)
                                                                _ = v137
                                                                _ = v138
                                                            } else {
                                                                v139 := GenS11{F0: 464, F1: 42, F2: 2}
                                                                if false {
                                                                    if (!true) {
                                                                        switch -256 {
                                                                        case -415:
                                                                            v2 = (v2 % ((-110) | 1))
                                                                        default:
                                                                            var v140 uint16 = 42
                                                                            _ = v140
                                                                        }
                                                                        var v141 int32 = int32(fi81)
                                                                        v142 := GenS19{F0: (v4 >> ((652) & 63))}
                                                                        _ = v141
                                                                        _ = v142
                                                                    }
                                                                    v143 := GenS11{F0: 18446744073709551615, F1: -1, F2: uint64(v6)}
                                                                    v144 := GenS6{F0: 2, F1: -124, F2: (v2 % ((-47) | 1))}
                                                                    _ = v143
                                                                    _ = v144
                                                                }
                                                                _ = v139
                                                            }
                                                            if ((v1 - v4) == (^(v1))) {
                                                                v145 := []int{(^(fi81))}
                                                                switch "a" {
                                                                case "\n\t":
                                                                    switch v3 {
                                                                    case "":
                                                                        var v146 int64 = 0
                                                                        _ = v146
                                                                    case "hello":
                                                                        var v147 uint64 = 4294967295
                                                                        _ = v147
                                                                    case "\n\t":
                                                                        if len(v145) > 0 {
                                                                            v145[0] = 127
                                                                        }
                                                                        v3 = (v3 + "")
                                                                        if true {
                                                                            switch "" {
                                                                            case "\n\t":
                                                                                switch fi81 {
                                                                                case 571:
                                                                                    if len(v145) >= 2 {
                                                                                        v145 = v145[1:len(v145)]
                                                                                    }
                                                                                default:
                                                                                    fmt.Println(v3, v1)
                                                                                }
                                                                            case "a":
                                                                                if true {
                                                                                    v4 -= (^(v4))
                                                                                } else {
                                                                                    var v148 uint32 = 647
                                                                                    _ = v148
                                                                                }
                                                                            default:
                                                                                v1 += (v4 / ((v4) | 1))
                                                                            }
                                                                            if (-2.25 != 0.0) {
                                                                                var v149 int32 = 42
                                                                                v3 = ((v3 + (("a" + (v3 + v3)) + "hello")) + "\n\t")
                                                                                var v150 uint16 = 42
                                                                                _ = v149
                                                                                _ = v150
                                                                            } else {
                                                                                switch (v3 + (v3 + v3)) {
                                                                                case "\n\t":
                                                                                    v2 = (v2 + 6)
                                                                                case "":
                                                                                    v151 := GenS20{F0: 0, F1: 1}
                                                                                    _ = v151
                                                                                case "a":
                                                                                    if (v4 < 583) {
                                                                                        v6 = (v6 | (v80 ^ v6))
                                                                                    } else {
                                                                                        v3 = "hello"
                                                                                        v5[2147483648] = 127
                                                                                        v1 = v1
                                                                                        v152 := GenS21{F0: (v2 % (((v2 >> ((4294967295) & 63))) | 1))}
                                                                                        var v153 string = "foo bar"
                                                                                        _ = v152
                                                                                        _ = v153
                                                                                    }
                                                                                default:
                                                                                    v154 := GenS2{F0: 2}
                                                                                    _ = v154
                                                                                }
                                                                                v155 := GenS17{F0: -814, F1: (fi83 & fi81), F2: v2}
                                                                                if true {
                                                                                    v145 = append(v145, (fi83 * 9223372036854775807))
                                                                                    if len(v145) >= 2 {
                                                                                        v145 = v145[1:len(v145)]
                                                                                    }
                                                                                    fmt.Println(v1, v1)
                                                                                } else {
                                                                                    if ((2 > 42) || true) {
                                                                                        _ = len(v5)
                                                                                        v1 = (v4 << ((4294967295) & 63))
                                                                                        var v156 uint16 = 42
                                                                                        _ = v156
                                                                                    } else {
                                                                                        var v157 int64 = -9223372036854775808
                                                                                        _ = v157
                                                                                    }
                                                                                    var v158 int8 = int8(v4)
                                                                                    _ = v158
                                                                                }
                                                                                _ = v155
                                                                            }
                                                                            var v159 float32 = -2.25
                                                                            _ = v159
                                                                        } else {
                                                                            v4 = v4
                                                                            switch (v3 + v3) {
                                                                            case "":
                                                                                if (!false) {
                                                                                    v6 += v80
                                                                                    switch (fi83 & (fi83 / ((fi81) | 1))) {
                                                                                    case 2147483647:
                                                                                        var v160 uint32 = (v1 * 65535)
                                                                                        _ = v160
                                                                                    }
                                                                                }
                                                                            default:
                                                                                var v161 int64 = 9223372036854775807
                                                                                var v162 uint64 = 2
                                                                                _ = v161
                                                                                _ = v162
                                                                            }
                                                                        }
                                                                    }
                                                                case "a":
                                                                    var v163 uint16 = 959
                                                                    _ = v163
                                                                }
                                                                var v164 string = v3
                                                                v1 += 1
                                                                _ = v145
                                                                _ = v164
                                                            }
                                                            switch v3 {
                                                            case "\n\t":
                                                                var v165 uint64 = 277
                                                                _ = v165
                                                            case "hello":
                                                                fmt.Println(v1, v4, v4)
                                                            case "foo bar":
                                                                switch (fi81 % ((fi83) | 1)) {
                                                                case 32767:
                                                                    fmt.Println(v6, v3, v80)
                                                                case -254:
                                                                    if ((false || (1 < int64(fi83))) || (v2 > (v2 & 42))) {
                                                                        if true {
                                                                            fmt.Println(v1, v1)
                                                                            var v166 int16 = -324
                                                                            fmt.Println(v4, fi81)
                                                                            _ = v166
                                                                        }
                                                                        v4 = v1
                                                                        var v167 int16 = 1
                                                                        _ = v167
                                                                    } else {
                                                                        v3 = (v3 + ("foo bar" + "a"))
                                                                        switch (fi81 | -9223372036854775808) {
                                                                        case 9223372036854775807:
                                                                            v2 += int8(fi81)
                                                                        case 2147483647:
                                                                            if (32767 != 65535) {
                                                                                switch fi83 {
                                                                                case 425:
                                                                                    v1 = 4294967295
                                                                                case -194:
                                                                                    v168 := v5[v4]
                                                                                    _ = v168
                                                                                case 32767:
                                                                                    if true {
                                                                                        if true {
                                                                                            delete(v5, (v1 - (^(v1))))
                                                                                        } else {
                                                                                            v4 -= 65535
                                                                                            v169 := []int{genFn0(-101, (!(156 == 223)), 108)}
                                                                                            switch fi83 {
                                                                                            case -865:
                                                                                                var v170 int8 = (v2 & v2)
                                                                                                _ = v170
                                                                                            }
                                                                                            switch fi81 {
                                                                                            case 42:
                                                                                                v171 := GenS1{F0: -425, F1: true}
                                                                                                _ = v171
                                                                                            case 128:
                                                                                                v172 := GenS4{F0: 993}
                                                                                                v80 -= v6
                                                                                                v3 = (v3 + "\n\t")
                                                                                                v6 -= v6
                                                                                                _ = v172
                                                                                            }
                                                                                            if (!false) {
                                                                                                v173 := GenS14{F0: 18446744073709551615}
                                                                                                var v174 int64 = 0
                                                                                                _ = v173
                                                                                                _ = v174
                                                                                            } else {
                                                                                                v175 := GenS5{F0: v2}
                                                                                                v176 := v5[v1]
                                                                                                _ = len(v169)
                                                                                                fmt.Println(v2, v1, v2)
                                                                                                var v177 float64 = (1.5 / 3.14159)
                                                                                                v1 = 0
                                                                                                switch "a" {
                                                                                                case "\n\t":
                                                                                                    v178 := GenS22{F0: -301, F1: 42, F2: -113}
                                                                                                    _ = v178
                                                                                                case "foo bar":
                                                                                                    v5[v4] = fi81
                                                                                                default:
                                                                                                    v80 -= (v80 | (v80 ^ 255))
                                                                                                }
                                                                                                _ = v175
                                                                                                _ = v176
                                                                                                _ = v177
                                                                                            }
                                                                                            _ = v169
                                                                                        }
                                                                                        v179 := GenS21{F0: 0}
                                                                                        var v180 int64 = 9223372036854775807
                                                                                        _ = v179
                                                                                        _ = v180
                                                                                    }
                                                                                }
                                                                                switch ((v3 + v3) + "hello") {
                                                                                case "a":
                                                                                    switch "hello" {
                                                                                    case "foo bar":
                                                                                        if true {
                                                                                            switch v3 {
                                                                                            case "hello":
                                                                                                var v181 uint = 42
                                                                                                _ = v181
                                                                                            case "\n\t":
                                                                                                switch "foo bar" {
                                                                                                case "foo bar":
                                                                                                    v3 = v3
                                                                                                    var v182 int64 = 42
                                                                                                    switch "" {
                                                                                                    case "":
                                                                                                        v3 = v3
                                                                                                    case "a":
                                                                                                        if true {
                                                                                                            v5[v1] = fi81
                                                                                                        } else {
                                                                                                            v182 = (v182 / (((v182 / ((v182) | 1))) | 1))
                                                                                                            v5[4294967295] = int(v80)
                                                                                                            var v183 uint64 = 0
                                                                                                            switch (("hello" + v3) + ("hello" + v3)) {
                                                                                                            case "":
                                                                                                                delete(v5, (v1 << ((0) & 63)))
                                                                                                            case "\n\t":
                                                                                                                var v184 bool = true
                                                                                                                _ = v184
                                                                                                            default:
                                                                                                                switch fi83 {
                                                                                                                case -961:
                                                                                                                    switch 42 {
                                                                                                                    case -9223372036854775808:
                                                                                                                        var v185 float32 = (1e10 + 1.5)
                                                                                                                        _ = v185
                                                                                                                    case 9223372036854775807:
                                                                                                                        if false {
                                                                                                                            v6 -= (v6 * 128)
                                                                                                                            var v186 uint32 = (v4 | 2147483648)
                                                                                                                            _ = v186
                                                                                                                        }
                                                                                                                    case 169:
                                                                                                                        v187 := []int{-634, fi83, 153, (fi81 + fi81)}
                                                                                                                        _ = v187
                                                                                                                    default:
                                                                                                                        v5[0] = (fi83 ^ 9223372036854775807)
                                                                                                                    }
                                                                                                                case 779:
                                                                                                                    var v188 int64 = 699
                                                                                                                    _ = v188
                                                                                                                }
                                                                                                            }
                                                                                                            _ = len(v5)
                                                                                                            var v189 uint16 = 255
                                                                                                            v2 = v2
                                                                                                            _ = v183
                                                                                                            _ = v189
                                                                                                        }
                                                                                                    case "hello":
                                                                                                        fmt.Println(v6)
                                                                                                    default:
                                                                                                        var v190 int16 = 1000
                                                                                                        _ = v190
                                                                                                    }
                                                                                                    _ = v182
                                                                                                case "":
                                                                                                    v191 := []int{(fi83 / ((42) | 1)), (^(fi81))}
                                                                                                    _ = v191
                                                                                                case "\n\t":
                                                                                                    if true {
                                                                                                        v80 -= 0
                                                                                                        v192 := []int{}
                                                                                                        _ = v192
                                                                                                    } else {
                                                                                                        fmt.Println(v6)
                                                                                                        v6 = 1
                                                                                                        v1 += 2147483648
                                                                                                    }
                                                                                                default:
                                                                                                    v193 := []int{4294967295, fi83, fi83, 146}
                                                                                                    _ = v193
                                                                                                }
                                                                                            case "foo bar":
                                                                                                switch (("\n\t" + (v3 + v3)) + ("a" + v3)) {
                                                                                                case "":
                                                                                                    fmt.Println(fi83, v6)
                                                                                                case "a":
                                                                                                    v194 := []int{fi81, (fi83 << ((413) & 63)), (fi83 + (fi83 ^ (fi83 + (-(fi83))))), (fi83 ^ fi81)}
                                                                                                    _ = v194
                                                                                                case "foo bar":
                                                                                                    v4 -= (v4 * v1)
                                                                                                }
                                                                                            default:
                                                                                                var v195 int64 = -9223372036854775808
                                                                                                _ = v195
                                                                                            }
                                                                                        } else {
                                                                                            var v196 string = ("" + v3)
                                                                                            _ = v196
                                                                                        }
                                                                                        var v197 float64 = (((3.14159 * 1e10) / 1e10) / 0.1)
                                                                                        _ = v197
                                                                                    case "hello":
                                                                                        switch "a" {
                                                                                        case "foo bar":
                                                                                            v80 = 172
                                                                                        default:
                                                                                            v198 := GenS16{F0: 0, F1: (fi83 ^ (^(fi83))), F2: 1}
                                                                                            _ = v198
                                                                                        }
                                                                                    default:
                                                                                        v199 := GenS5{F0: (v2 - -128)}
                                                                                        _ = v199
                                                                                    }
                                                                                }
                                                                                fmt.Println(v6, v80)
                                                                            } else {
                                                                                if (true && (42 != 4294967295)) {
                                                                                    var v200 string = ("foo bar" + v3)
                                                                                    _ = v200
                                                                                }
                                                                                v201 := GenS17{F0: -1, F1: fi81, F2: v2}
                                                                                _ = v201
                                                                            }
                                                                        }
                                                                    }
                                                                default:
                                                                    v202 := GenS14{F0: 133}
                                                                    _ = v202
                                                                }
                                                            default:
                                                                if (-1 == -487) {
                                                                    switch "foo bar" {
                                                                    case "foo bar":
                                                                        v203 := GenS23{F0: genFn1(-36, (fi81 & genFn0((v2 - 127), (32767 < uint16(v80)), 0)), 1)}
                                                                        _ = v203
                                                                    }
                                                                    v204 := GenS8{F0: 2147483647, F1: 255, F2: 2}
                                                                    _ = v204
                                                                }
                                                            }
                                                            v1 -= (v4 & 204)
                                                            fmt.Println(v6, v4)
                                                            v80 -= (v80 - (v80 / (((v6 * v6)) | 1)))
                                                        }
                                                        switch (fi83 % ((-1) | 1)) {
                                                        case -870:
                                                            v205 := []int{}
                                                            var v206 int64 = 42
                                                            var v207 string = v3
                                                            _ = v205
                                                            _ = v206
                                                            _ = v207
                                                        case -645:
                                                            var v208 string = v3
                                                            _ = v208
                                                        case -612:
                                                            v1 -= v4
                                                        }
                                                    }
                                                default:
                                                    if (((true && true) || (-863 < -1)) && false) {
                                                        fmt.Println(v80, v1)
                                                        var v209 uint64 = 1
                                                        if false {
                                                            _ = len(v5)
                                                            v209 -= 625
                                                            v210 := GenS1{F0: 32767, F1: (!true)}
                                                            _ = v210
                                                        } else {
                                                            v211 := v5[(v1 / ((112) | 1))]
                                                            v3 = v3
                                                            _ = v211
                                                        }
                                                        _ = v209
                                                    }
                                                }
                                                if false {
                                                    fmt.Println(v3, fi81, v4)
                                                }
                                                if true {
                                                    var v212 uint16 = 338
                                                    _ = v212
                                                } else {
                                                    v213 := []int{genFn0(-1, (533 < 855), 1), fi81}
                                                    v1 = v4
                                                    _ = v213
                                                }
                                                fmt.Println(v3, fi81, fi81)
                                                if false {
                                                    v1 = (^(v4))
                                                    if true {
                                                        var v214 int16 = 463
                                                        var v215 string = v3
                                                        _ = v214
                                                        _ = v215
                                                    } else {
                                                        var v216 int8 = (v2 - (v2 % ((int8(v6)) | 1)))
                                                        if (!false) {
                                                            v217 := []int{fi83, fi81}
                                                            _ = v217
                                                        }
                                                        var v218 uint16 = 0
                                                        _ = v216
                                                        _ = v218
                                                    }
                                                }
                                            }
                                            var v219 bool = (!true)
                                            _ = v219
                                        default:
                                            var v220 int8 = (v2 >> ((42) & 63))
                                            _ = v220
                                        }
                                        var v221 bool = (-1 == 9223372036854775807)
                                        _ = v221
                                    }
                                }
                                var v222 int64 = 42
                                var v223 bool = true
                                var v224 int32 = 42
                                _ = v222
                                _ = v223
                                _ = v224
                                _ = v84
                            case "hello":
                                v3 = ("a" + (v3 + (v3 + v3)))
                            case "\n\t":
                                v225 := []int{42, genFn0(0, false, 42)}
                                _ = v225
                            }
                            fmt.Println(fi83, v80)
                            v80 = (v6 | v6)
                            fmt.Println(v3, v4, fi81)
                            var v226 int16 = 42
                            _ = fi83
                            _ = v226
                        }
                    case "foo bar":
                        v227 := GenS18{F0: 42, F1: v4, F2: -99}
                        _ = v227
                    case "\n\t":
                        var v228 int8 = v2
                        _ = v228
                    default:
                        if true {
                            if true {
                                var v229 int64 = -9223372036854775808
                                v230 := GenS24{F0: 4294967295, F1: (v229 << ((65535) & 63))}
                                _ = v229
                                _ = v230
                            } else {
                                v4 += 511
                                for fi231 := 0; fi231 < 11; fi231++ {
                                    v80 = 2
                                    delete(v5, 2147483648)
                                    v80 = v6
                                    v80 = v80
                                    if (true || true) {
                                        v1 -= (v1 * (^(v4)))
                                        _ = len(v5)
                                        v1 += (v1 * (v1 + v1))
                                    } else {
                                        switch v3 {
                                        case "\n\t":
                                            fmt.Println(v4, v4)
                                        }
                                    }
                                    _ = fi231
                                }
                                var v232 float64 = float64(1e10)
                                _ = v232
                            }
                            var v233 bool = ((((-2.25 + -2.25) + (-2.25 - (1e10 * 0.0))) >= 3.14159) && false)
                            if (!((int32(v80) == 336) || v233)) {
                                v80 = 0
                            } else {
                                for fi234 := 0; fi234 < 6; fi234++ {
                                    var v235 uint32 = (^(v4))
                                    _ = fi234
                                    _ = v235
                                }
                                var v236 string = (v3 + (((v3 + (v3 + (v3 + v3))) + v3) + v3))
                                _ = v236
                            }
                            _ = v233
                        }
                    }
                }
                var v237 bool = ((true && false) && (-1 == 0))
                v1 = (v1 << ((27) & 63))
                _ = fi81
                _ = v237
            }
            v238 := GenS25{F0: 2, F1: 1000, F2: 2147483647}
            v5[2] = 413
            _ = v238
        }
        if false {
            fmt.Println(v80)
        }
        _ = v80
    } else {
        var v239 uint = 4294967295
        var v240 int64 = 0
        switch 2147483647 {
        case -202:
            v241 := GenS26{F0: v2, F1: 452}
            _ = v241
        }
        _ = v239
        _ = v240
    }
    v242 := GenS22{F0: -434, F1: -221, F2: (v2 - 127)}
    switch 458 {
    case 317:
        v2 += (v2 % ((127) | 1))
    case 1:
        for fi243 := 0; fi243 < 17; fi243++ {
            if ((fi243 != -9223372036854775808) && (!(!(true || (((250 == 18446744073709551615) || false) && true))))) {
                v5[v4] = (fi243 * 4294967295)
            }
            v244 := []int{fi243, fi243, (fi243 + (fi243 / ((fi243) | 1)))}
            _ = fi243
            _ = v244
        }
    case 32768:
        switch v3 {
        case "hello":
            fmt.Println(v2, v3)
        case "foo bar":
            var v245 int64 = 1
            _ = v245
        case "\n\t":
            var v246 bool = false
            _ = v246
        }
    default:
        v247 := GenS27{F0: 91}
        _ = v247
    }
    v1 += (^(v1))
    if (!(!(1e10 == (1.5 + 0.1)))) {
        var v248 bool = true
        switch -2147483648 {
        case -2147483648:
            delete(v5, v1)
        case 4294967295:
            v249 := v5[0]
            _ = v249
        }
        if (false && true) {
            v3 = v3
        }
        _ = v248
    } else {
        for fi250 := 0; fi250 < 9; fi250++ {
            v251 := GenS10{F0: 918, F1: v2}
            _ = fi250
            _ = v251
        }
    }
    v4 += v1
    v3 = "foo bar"
    switch 9223372036854775807 {
    case 844:
        v4 = v4
    case 9223372036854775807:
        v252 := GenS10{F0: 997, F1: -1}
        _ = v252
    case 0:
        var v253 bool = (!(false || false))
        _ = v253
    default:
        v254 := GenS28{F0: -249, F1: 2147483648}
        _ = v254
    }
    v1 -= (^(v1))
    v5[(v1 * v1)] = 32767
    if false {
        v255 := GenS20{F0: -2147483648, F1: 42}
        _ = v255
    }
    for fi256 := 0; fi256 < 6; fi256++ {
        var v257 int16 = -32768
        v258 := GenS26{F0: (v2 << ((1) & 63)), F1: 18446744073709551615}
        _ = fi256
        _ = v257
        _ = v258
    }
    v259 := GenS26{F0: (^(v2)), F1: 4294967295}
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v242
    _ = v259
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
