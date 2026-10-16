package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64; F2 string }
type GenS2 struct { F0 int64; F1 int8; F2 uint64 }
type GenS3 struct { F0 int; F1 int32 }
type GenS4 struct { F0 bool; F1 int }
type GenS5 struct { F0 uint; F1 int }
type GenS6 struct { F0 uint32; F1 uint16 }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 uint64; F1 bool; F2 bool }
type GenS9 struct { F0 uint32; F1 bool }
type GenS10 struct { F0 uint64 }
type GenS11 struct { F0 int64; F1 int32; F2 uint16 }
type GenS12 struct { F0 int8; F1 bool }
type GenS13 struct { F0 int64; F1 int; F2 uint16 }
type GenS14 struct { F0 int16; F1 uint64 }
type GenS15 struct { F0 uint64; F1 uint8 }
type GenS16 struct { F0 int8; F1 int16 }
type GenS17 struct { F0 bool; F1 uint16 }
type GenS18 struct { F0 uint16; F1 bool }
type GenS19 struct { F0 int64 }
type GenS20 struct { F0 uint32 }
type GenS21 struct { F0 bool }
type GenS22 struct { F0 bool }
type GenS23 struct { F0 uint32; F1 int8 }
type GenS24 struct { F0 int64 }
type GenS25 struct { F0 int8; F1 uint64; F2 int32 }
type GenS26 struct { F0 uint32 }
func genFn0(p0 uint) int8 {
    var v6 int64 = 9223372036854775807
    var v7 int32 = 1
    for fi8 := 0; fi8 < 11; fi8++ {
        if true {
            v9 := []int{(fi8 / ((fi8) | 1))}
            _ = v9
        } else {
            if ((p0 >> ((243) & 63)) == p0) {
                switch fi8 {
                case 128:
                    v10 := []int{(fi8 | (fi8 % (((^(fi8))) | 1))), (fi8 | fi8), fi8, (fi8 & (fi8 >> (((p0 + (p0 >> ((65535) & 63)))) & 63)))}
                    _ = v10
                case 0:
                    for fi11 := 0; fi11 < 12; fi11++ {
                        v12 := GenS1{F0: 65535, F1: 4294967295, F2: "\n\t"}
                        v6 -= (v6 % ((v6) | 1))
                        if true {
                            if true {
                                var v13 uint = (p0 & (p0 >> ((527) & 63)))
                                p0 = v13
                                v6 = v6
                                v14 := GenS1{F0: p0, F1: 18446744073709551615, F2: "foo bar"}
                                if ((!(0 <= 42)) && ((148 < 255) && (1 >= 32767))) {
                                    switch 128 {
                                    case 1:
                                        v13 = 916
                                    case 4294967295:
                                        v15 := []int{fi8, fi11, fi8, 232}
                                        _ = v15
                                    default:
                                        v6 -= (^(v6))
                                    }
                                    v16 := []int{(^(fi8)), (-(fi8)), fi11}
                                    if (("hello" + ("hello" + "\n\t")) == "foo bar") {
                                        if true {
                                            if false {
                                                v6 -= v6
                                                fmt.Println(fi11, fi11, v13)
                                                if true {
                                                    var v17 bool = (160 >= 231)
                                                    v18 := GenS2{F0: -664, F1: 68, F2: 1}
                                                    p0 += (v13 ^ (v13 % ((p0) | 1)))
                                                    switch ("\n\t" + (((("\n\t" + ("\n\t" + "hello")) + "hello") + "\n\t") + "hello")) {
                                                    case "foo bar":
                                                        v13 = (p0 / ((v13) | 1))
                                                    case "":
                                                        v13 = 1
                                                    case "a":
                                                        var v19 bool = false
                                                        _ = v19
                                                    default:
                                                        var v20 uint = (v13 * (v13 << (((^(v13))) & 63)))
                                                        _ = v20
                                                    }
                                                    _ = v17
                                                    _ = v18
                                                }
                                            } else {
                                                v16 = append(v16, (fi11 / ((42) | 1)))
                                                v7 = -447
                                                _ = len(v16)
                                            }
                                            v16 = append(v16, (fi11 ^ 380))
                                            fmt.Println(v6, fi8)
                                            var v21 float32 = -2.25
                                            if len(v16) > 0 {
                                                v16[0] = fi8
                                            }
                                            _ = v21
                                        }
                                        switch (fi11 >> ((p0) & 63)) {
                                        case -361:
                                            if (1 < 0) {
                                                v22 := GenS1{F0: p0, F1: 18446744073709551615, F2: ("a" + "foo bar")}
                                                _ = v22
                                            } else {
                                                var v23 string = "\n\t"
                                                if (false || false) {
                                                    if len(v16) >= 2 {
                                                        v16 = v16[1:len(v16)]
                                                    }
                                                } else {
                                                    if true {
                                                        if ((-(v6)) == int64(v7)) {
                                                            var v24 int8 = -128
                                                            _ = v24
                                                        } else {
                                                            v25 := GenS1{F0: v13, F1: 2, F2: ((("\n\t" + "hello") + v23) + ("a" + v23))}
                                                            v6 -= 42
                                                            var v26 float32 = 1e10
                                                            _ = v25
                                                            _ = v26
                                                        }
                                                    }
                                                }
                                                fmt.Println(p0, v6)
                                                _ = v23
                                            }
                                        default:
                                            fmt.Println(v13, v7, v7)
                                        }
                                        var v27 float32 = 1.5
                                        _ = v27
                                    } else {
                                        var v28 int32 = 0
                                        _ = v28
                                    }
                                    _ = v16
                                } else {
                                    v6 = (v6 ^ (-(v6)))
                                    var v29 uint = v13
                                    _ = v29
                                }
                                _ = v13
                                _ = v14
                            }
                        }
                        _ = fi11
                        _ = v12
                    }
                default:
                    p0 = p0
                }
            }
        }
        var v30 float64 = (0.0 - 1e10)
        v31 := []int{fi8}
        _ = fi8
        _ = v30
        _ = v31
    }
    _ = p0
    _ = v6
    _ = v7
    return 0
}
func genFn1(p0 uint32) string {
    if (0 != -9223372036854775808) {
        var v32 int8 = -48
        if false {
            v33 := GenS2{F0: 9223372036854775807, F1: (v32 % ((v32) | 1)), F2: 1}
            for fi34 := 0; fi34 < 7; fi34++ {
                switch "" {
                case "hello":
                    for fi35 := 0; fi35 < 13; fi35++ {
                        var v36 uint8 = 0
                        if true {
                            p0 = p0
                            p0 += p0
                            var v37 int32 = -1
                            _ = v37
                        }
                        p0 = p0
                        _ = fi35
                        _ = v36
                    }
                case "":
                    fmt.Println(fi34, v32)
                case "a":
                    p0 = p0
                }
                _ = fi34
            }
            _ = v33
        }
        v38 := []int{-1, 127}
        _ = v32
        _ = v38
    } else {
        if (905 != -9223372036854775808) {
            v39 := []int{}
            v40 := GenS3{F0: 1, F1: 872}
            if false {
                fmt.Println(p0)
                for fi41 := 0; fi41 < 1; fi41++ {
                    if false {
                        if (2 >= 32767) {
                            if ((-70 <= -128) || (!(uint8(p0) == 128))) {
                                v42 := GenS4{F0: true, F1: (fi41 - fi41)}
                                _ = v42
                            }
                            if len(v39) >= 2 {
                                v39 = v39[1:len(v39)]
                            }
                            if len(v39) > 0 {
                                v39[0] = (-(fi41))
                            }
                        } else {
                            fmt.Println(p0)
                            v43 := GenS2{F0: -576, F1: -128, F2: uint64(p0)}
                            p0 = p0
                            _ = v43
                        }
                        var v44 int32 = 1
                        _ = v44
                    } else {
                        if true {
                            if (!((458 >= 1) || (false && (!true)))) {
                                for fi45 := 0; fi45 < 3; fi45++ {
                                    p0 = p0
                                    if true {
                                        if false {
                                            if len(v39) >= 2 {
                                                v39 = v39[1:len(v39)]
                                            }
                                            p0 += p0
                                            v46 := GenS2{F0: 9223372036854775807, F1: -128, F2: 42}
                                            _ = v46
                                        }
                                    } else {
                                        fmt.Println(fi41, fi41, p0)
                                        if (1 >= -114) {
                                            if (((((-2.25 / ((0.1 * 0.1) + 1.5)) / -2.25) + 1.5) >= float64(((3.14159 - 0.0) - ((-2.25 - (0.0 + 1.5)) * 0.1)))) && (!(695 < 2))) {
                                                if false {
                                                    var v47 float64 = 0.1
                                                    _ = v47
                                                } else {
                                                    var v48 int8 = genFn0(505)
                                                    _ = v48
                                                }
                                                p0 = (p0 + p0)
                                            } else {
                                                p0 = (p0 & p0)
                                            }
                                            p0 = p0
                                            if (18446744073709551615 <= 42) {
                                                p0 = (p0 ^ p0)
                                                if true {
                                                    v49 := GenS1{F0: 42, F1: 1, F2: ("hello" + "")}
                                                    _ = v49
                                                }
                                                var v50 float64 = 0.1
                                                p0 += (^(p0))
                                                var v51 int8 = -19
                                                fmt.Println(v51, v50, fi41)
                                                var v52 uint64 = 2
                                                _ = v50
                                                _ = v51
                                                _ = v52
                                            } else {
                                                p0 += p0
                                            }
                                        }
                                        p0 = p0
                                    }
                                    _ = fi45
                                }
                                _ = len(v39)
                                switch fi41 {
                                case -2147483648:
                                    for fi53 := 0; fi53 < 11; fi53++ {
                                        var v54 string = "hello"
                                        _ = fi53
                                        _ = v54
                                    }
                                default:
                                    v55 := GenS5{F0: 1, F1: (fi41 ^ fi41)}
                                    _ = v55
                                }
                            } else {
                                var v56 uint8 = 255
                                switch "" {
                                case "hello":
                                    v56 = (v56 % ((v56) | 1))
                                }
                                var v57 float64 = (((0.1 + float64(0.1)) - 0.0) - 3.14159)
                                _ = v56
                                _ = v57
                            }
                            if false {
                                p0 += (p0 ^ p0)
                                v39 = append(v39, (fi41 % ((fi41) | 1)))
                                var v58 bool = ((p0 | (p0 * 42)) != (p0 / ((386) | 1)))
                                _ = v58
                            }
                        } else {
                            if (genFn0(65535) > genFn0(65535)) {
                                switch (fi41 - 32768) {
                                case -1:
                                    var v59 int8 = genFn0(498)
                                    _ = v59
                                default:
                                    if true {
                                        p0 -= (^(p0))
                                        if false {
                                            for fi60 := 0; fi60 < 1; fi60++ {
                                                v61 := GenS2{F0: -17, F1: 100, F2: 379}
                                                v39 = append(v39, (fi60 + fi41))
                                                var v62 string = "\n\t"
                                                _ = fi60
                                                _ = v61
                                                _ = v62
                                            }
                                            if len(v39) >= 2 {
                                                v39 = v39[1:len(v39)]
                                            }
                                            if len(v39) >= 2 {
                                                v39 = v39[1:len(v39)]
                                            }
                                        } else {
                                            if len(v39) > 0 {
                                                v39[0] = (fi41 % (((fi41 + (fi41 / (((fi41 * (fi41 & fi41))) | 1)))) | 1))
                                            }
                                            p0 = p0
                                        }
                                        var v63 string = "foo bar"
                                        _ = v63
                                    }
                                }
                                p0 -= p0
                                p0 -= (p0 & p0)
                            } else {
                                p0 -= (p0 >> ((1) & 63))
                                var v64 uint32 = 2147483648
                                _ = v64
                            }
                            _ = len(v39)
                        }
                        fmt.Println(p0, fi41)
                        p0 = p0
                    }
                    p0 = 1
                    fmt.Println(p0, p0)
                    _ = fi41
                }
                if true {
                    p0 -= p0
                } else {
                    for fi65 := 0; fi65 < 14; fi65++ {
                        v66 := GenS5{F0: 74, F1: fi65}
                        _ = fi65
                        _ = v66
                    }
                    var v67 uint32 = (p0 / ((42) | 1))
                    _ = v67
                }
            }
            _ = v39
            _ = v40
        }
    }
    if (true || true) {
        if false {
            var v68 int64 = 1
            v69 := []int{}
            _ = v68
            _ = v69
        } else {
            p0 = 65535
            if true {
                var v70 bool = true
                v71 := GenS6{F0: 897, F1: 338}
                _ = v70
                _ = v71
            }
        }
        for fi72 := 0; fi72 < 20; fi72++ {
            fmt.Println(p0, fi72)
            _ = fi72
        }
        var v73 int16 = 1
        _ = v73
    } else {
        p0 = p0
        var v74 int32 = -2147483648
        _ = v74
    }
    _ = p0
    return "\n\t"
}

func main() {
    var v1 int8 = 100
    var v2 string = "a"
    v3 := make(map[uint32]uint32)
    var v4 string = (v2 + "hello")
    v5 := make(map[int]bool)
    for fi75 := 0; fi75 < 17; fi75++ {
        var v76 uint = 785
        var v77 uint8 = 42
        fmt.Println(v4, v2, fi75)
        _ = fi75
        _ = v76
        _ = v77
    }
    v4 = (((v2 + v2) + "") + "\n\t")
    fmt.Println(v4, v2)
    for fi78 := 0; fi78 < 3; fi78++ {
        for fi79 := 0; fi79 < 7; fi79++ {
            if (v1 < (v1 & -128)) {
                fmt.Println(v2)
                var v80 int8 = genFn0(535)
                _ = v80
            } else {
                if (!true) {
                    var v81 float64 = 1.5
                    switch ((v4 + (("" + (("" + (v4 + ((v4 + (v2 + ((("hello" + (((((v4 + (v4 + ("a" + v2))) + "a") + v2) + (v4 + ((v4 + ("" + "hello")) + v2))) + "foo bar")) + (("foo bar" + v4) + (("foo bar" + "\n\t") + (("\n\t" + "\n\t") + v2)))) + ("\n\t" + (v2 + v2))))) + ""))) + ("a" + (v2 + "a")))) + v4)) + "a") {
                    case "foo bar":
                        fmt.Println(v1, fi79)
                    case "a":
                        var v82 int8 = genFn0(2)
                        _ = v82
                    }
                    var v83 int8 = 127
                    _ = v81
                    _ = v83
                } else {
                    v4 = ""
                    if (254 > 2147483648) {
                        if false {
                            if true {
                                if ((("foo bar" + ((v4 + ("hello" + (v2 + ("hello" + v4)))) + "")) + (v2 + "foo bar")) == (v2 + v2)) {
                                    v84 := []int{fi79, fi79, fi78}
                                    switch fi79 {
                                    case 9223372036854775807:
                                        if false {
                                            v3[42] = 90
                                            var v85 int8 = v1
                                            if (0.1 == float64(-2.25)) {
                                                var v86 int = 1
                                                if (false || (413 <= 2147483647)) {
                                                    if true {
                                                        var v87 int32 = -2147483648
                                                        if len(v84) >= 2 {
                                                            v84 = v84[1:len(v84)]
                                                        }
                                                        _ = v87
                                                    }
                                                    v88 := GenS7{F0: 2147483647}
                                                    v89 := GenS4{F0: (383 > 2), F1: fi79}
                                                    _ = v88
                                                    _ = v89
                                                } else {
                                                    if true {
                                                        fmt.Println(v2)
                                                    }
                                                }
                                                fmt.Println(v2, v86, fi78)
                                                _ = v86
                                            } else {
                                                switch v2 {
                                                case "a":
                                                    var v90 string = v2
                                                    if ((!((!(-745 <= -1)) || (((true && true) && false) || true))) && false) {
                                                        if false {
                                                            v91 := v5[fi78]
                                                            _ = v91
                                                        }
                                                        v92 := GenS3{F0: fi79, F1: -1}
                                                        fmt.Println(v1, fi78)
                                                        _ = v92
                                                    } else {
                                                        var v93 int16 = 281
                                                        _ = v93
                                                    }
                                                    if ((2147483647 != -151) || false) {
                                                        var v94 float64 = 1.5
                                                        _ = v94
                                                    }
                                                    _ = v90
                                                case "hello":
                                                    v95 := GenS8{F0: 1, F1: true, F2: false}
                                                    _ = v95
                                                case "foo bar":
                                                    v96 := GenS9{F0: 2147483648, F1: false}
                                                    _ = v96
                                                default:
                                                    v97 := GenS10{F0: 415}
                                                    _ = v97
                                                }
                                                if false {
                                                    if false {
                                                        v98 := v3[372]
                                                        v1 = (v85 - 127)
                                                        var v99 uint8 = 255
                                                        v100 := GenS11{F0: 0, F1: -123, F2: 65535}
                                                        switch fi78 {
                                                        case 128:
                                                            fmt.Println(v1, v1, v85)
                                                            if false {
                                                                switch "hello" {
                                                                case "\n\t":
                                                                    switch (v4 + v2) {
                                                                    case "":
                                                                        switch fi78 {
                                                                        case 4294967295:
                                                                            switch (v4 + ("hello" + "hello")) {
                                                                            case "a":
                                                                                v101 := v5[fi78]
                                                                                switch ("" + v2) {
                                                                                case "hello":
                                                                                    var v102 uint16 = 255
                                                                                    _ = v102
                                                                                case "\n\t":
                                                                                    v4 = ((("" + "") + (v2 + "")) + "hello")
                                                                                case "a":
                                                                                    if v101 {
                                                                                        v98 = 42
                                                                                    } else {
                                                                                        switch fi78 {
                                                                                        case -1:
                                                                                            v98 -= v98
                                                                                        case -9223372036854775808:
                                                                                            v103 := GenS8{F0: 4294967295, F1: v101, F2: ((v98 / (((^(v98))) | 1)) > v98)}
                                                                                            _ = v103
                                                                                        default:
                                                                                            var v104 int32 = -1
                                                                                            _ = v104
                                                                                        }
                                                                                        var v105 int = 513
                                                                                        _ = len(v84)
                                                                                        _ = v105
                                                                                    }
                                                                                default:
                                                                                    var v106 float32 = 3.14159
                                                                                    _ = v106
                                                                                }
                                                                                v85 = v1
                                                                                _ = v101
                                                                            }
                                                                        }
                                                                    case "hello":
                                                                        v107 := GenS12{F0: genFn0(4294967295), F1: (18446744073709551615 >= 4294967295)}
                                                                        _ = v107
                                                                    }
                                                                case "foo bar":
                                                                    fmt.Println(v98)
                                                                case "":
                                                                    fmt.Println(v4, fi78)
                                                                default:
                                                                    v108 := GenS6{F0: (v98 - 42), F1: 1}
                                                                    _ = v108
                                                                }
                                                                if len(v84) >= 2 {
                                                                    v84 = v84[1:len(v84)]
                                                                }
                                                            }
                                                            var v109 uint32 = 2
                                                            fmt.Println(v98, fi79)
                                                            _ = v109
                                                        }
                                                        if true {
                                                            switch fi79 {
                                                            case 32768:
                                                                if false {
                                                                    v85 -= (-(v1))
                                                                } else {
                                                                    var v110 uint8 = v99
                                                                    var v111 float32 = 0.0
                                                                    v99 += (^(v99))
                                                                    _ = v110
                                                                    _ = v111
                                                                }
                                                            case 0:
                                                                var v112 int32 = 2147483647
                                                                _ = v112
                                                            default:
                                                                v113 := v3[v98]
                                                                v2 = v4
                                                                var v114 float64 = 1e10
                                                                _ = v113
                                                                _ = v114
                                                            }
                                                            var v115 float64 = (3.14159 / 0.1)
                                                            if (921 > -435) {
                                                                v116 := GenS8{F0: 42, F1: false, F2: (((1 > genFn0(1)) || false) || ((v99 << ((202) & 63)) >= v99))}
                                                                v117 := GenS10{F0: 810}
                                                                _ = v116
                                                                _ = v117
                                                            }
                                                            _ = v115
                                                        } else {
                                                            v99 += v99
                                                            switch -731 {
                                                            case 32768:
                                                                v4 = (v2 + v4)
                                                            }
                                                            v1 += (v85 * -10)
                                                            v84 = append(v84, (fi78 - (fi79 - (fi78 & (fi78 << ((418) & 63))))))
                                                            v118 := GenS13{F0: 9223372036854775807, F1: (fi79 >> ((42) & 63)), F2: 255}
                                                            switch fi79 {
                                                            case 127:
                                                                if true {
                                                                    v119 := GenS14{F0: 1000, F1: 1}
                                                                    v2 = v2
                                                                    fmt.Println(v98, v1)
                                                                    _ = v119
                                                                }
                                                            }
                                                            var v120 string = v4
                                                            fmt.Println(fi79, v1)
                                                            if (false || ((-9223372036854775808 >= 153) && false)) {
                                                                v98 = uint32(v99)
                                                            }
                                                            var v121 bool = (490 != 255)
                                                            v122 := GenS15{F0: 1, F1: (v99 - (^(v99)))}
                                                            fmt.Println(fi79)
                                                            var v123 float32 = ((0.1 * 1.5) * (0.1 + (float32((1.5 + 0.1)) * 3.14159)))
                                                            _ = v118
                                                            _ = v120
                                                            _ = v121
                                                            _ = v122
                                                            _ = v123
                                                        }
                                                        if (false && false) {
                                                            _ = len(v5)
                                                        } else {
                                                            v4 = v2
                                                            if (!true) {
                                                                var v124 float32 = (1.5 + 1e10)
                                                                if false {
                                                                    if true {
                                                                        switch (fi79 ^ 224) {
                                                                        case 42:
                                                                            v84 = append(v84, -1)
                                                                        }
                                                                        v99 = (v99 | 244)
                                                                    } else {
                                                                        _ = len(v3)
                                                                    }
                                                                    if false {
                                                                        var v125 uint32 = (v98 ^ v98)
                                                                        v2 = ("a" + "")
                                                                        if false {
                                                                            if len(v84) >= 2 {
                                                                                v84 = v84[1:len(v84)]
                                                                            }
                                                                            v84 = append(v84, (fi79 + (fi79 * (fi79 / ((fi79) | 1)))))
                                                                        }
                                                                        if (float32((3.14159 / 1.0)) < 1e10) {
                                                                            switch (-(fi79)) {
                                                                            case 32767:
                                                                                var v126 int = fi78
                                                                                _ = v126
                                                                            case 9223372036854775807:
                                                                                var v127 int32 = -2147483648
                                                                                _ = v127
                                                                            }
                                                                            var v128 int32 = -2147483648
                                                                            if (v1 < -43) {
                                                                                v129 := GenS6{F0: 485, F1: 225}
                                                                                _ = v129
                                                                            } else {
                                                                                var v130 uint32 = (v98 ^ (v98 * v125))
                                                                                v98 -= 2147483648
                                                                                if len(v84) >= 2 {
                                                                                    v84 = v84[1:len(v84)]
                                                                                }
                                                                                var v131 uint16 = 32767
                                                                                _ = v130
                                                                                _ = v131
                                                                            }
                                                                            _ = v128
                                                                        } else {
                                                                            fmt.Println(v98)
                                                                            v132 := GenS16{F0: genFn0(583), F1: 559}
                                                                            if (!(0 >= 18446744073709551615)) {
                                                                                var v133 uint = 335
                                                                                v134 := GenS10{F0: 18446744073709551615}
                                                                                v124 -= v124
                                                                                _ = v133
                                                                                _ = v134
                                                                            } else {
                                                                                if (v124 == v124) {
                                                                                    if false {
                                                                                        if (false || (genFn0(369) != (v85 | (v1 >> ((0) & 63))))) {
                                                                                            v1 += -1
                                                                                            var v135 int64 = 9223372036854775807
                                                                                            if (!((v2 < v4) || false)) {
                                                                                                var v136 int16 = 851
                                                                                                v135 -= v135
                                                                                                var v137 uint = 255
                                                                                                v138 := GenS8{F0: 18446744073709551615, F1: true, F2: (v124 >= -2.25)}
                                                                                                _ = v136
                                                                                                _ = v137
                                                                                                _ = v138
                                                                                            }
                                                                                            _ = v135
                                                                                        }
                                                                                        v4 = v2
                                                                                        v124 -= (((v124 * v124) / float32(0.1)) - (v124 / (0.0 + v124)))
                                                                                    }
                                                                                    v139 := GenS13{F0: -200, F1: -1, F2: 255}
                                                                                    v125 -= (v98 | 4294967295)
                                                                                    _ = v139
                                                                                } else {
                                                                                    v140 := GenS17{F0: false, F1: 42}
                                                                                    v141 := GenS18{F0: 541, F1: false}
                                                                                    _ = v140
                                                                                    _ = v141
                                                                                }
                                                                                switch 128 {
                                                                                case 32768:
                                                                                    var v142 int32 = -1
                                                                                    _ = v142
                                                                                default:
                                                                                    delete(v3, v125)
                                                                                }
                                                                                v125 = (v98 | v98)
                                                                            }
                                                                            _ = v132
                                                                        }
                                                                        v1 = v1
                                                                        _ = v125
                                                                    }
                                                                    v1 -= 42
                                                                }
                                                                _ = len(v3)
                                                                _ = v124
                                                            }
                                                            var v143 int = int(v98)
                                                            if len(v84) >= 2 {
                                                                v84 = v84[1:len(v84)]
                                                            }
                                                            if len(v84) >= 2 {
                                                                v84 = v84[1:len(v84)]
                                                            }
                                                            _ = v143
                                                        }
                                                        var v144 uint16 = 32767
                                                        var v145 int32 = 42
                                                        v146 := GenS11{F0: -1, F1: (v145 / ((1) | 1)), F2: 916}
                                                        fmt.Println(v85, fi79, v98)
                                                        v1 = (^(v85))
                                                        _ = v100
                                                        _ = v144
                                                        _ = v145
                                                        _ = v146
                                                        _ = v98
                                                        _ = v99
                                                    }
                                                    v147 := GenS19{F0: 482}
                                                    v148 := GenS20{F0: 0}
                                                    _ = v147
                                                    _ = v148
                                                }
                                            }
                                            v1 += v1
                                            var v149 float64 = 0.1
                                            _ = v149
                                            _ = v85
                                        } else {
                                            v1 = v1
                                        }
                                    case 277:
                                        fmt.Println(fi78)
                                    case 4294967295:
                                        v4 = ""
                                    default:
                                        if false {
                                            v2 = v2
                                            var v150 int32 = 1
                                            var v151 int16 = -32768
                                            _ = v150
                                            _ = v151
                                        } else {
                                            var v152 bool = (3.14159 == 0.1)
                                            switch v4 {
                                            case "":
                                                if len(v84) > 0 {
                                                    v84[0] = int(v1)
                                                }
                                            case "foo bar":
                                                v153 := GenS4{F0: true, F1: (fi78 & fi78)}
                                                _ = v153
                                            default:
                                                v154 := GenS21{F0: (!(1 == 2147483648))}
                                                _ = v154
                                            }
                                            _ = v152
                                        }
                                    }
                                    switch (fi79 % ((fi79) | 1)) {
                                    case -272:
                                        _ = len(v5)
                                    }
                                    switch (v4 + v4) {
                                    case "a":
                                        if (135 == 1) {
                                            if (false && false) {
                                                v2 = v2
                                            } else {
                                                v155 := GenS22{F0: (761 > 2147483647)}
                                                _ = len(v84)
                                                v4 = "a"
                                                _ = v155
                                            }
                                            fmt.Println(fi78, fi79, fi78)
                                            var v156 uint64 = 4294967295
                                            _ = v156
                                        }
                                    default:
                                        v2 = v2
                                    }
                                    _ = v84
                                }
                                switch (fi78 | 2147483647) {
                                case 127:
                                    v2 = ("" + ((((("foo bar" + ((v4 + "a") + v4)) + (v4 + "\n\t")) + (v2 + ("a" + (("\n\t" + ("hello" + v4)) + v2)))) + v2) + ((v4 + ((v2 + ("hello" + v2)) + v4)) + (v2 + v2))))
                                case 0:
                                    v2 = ("" + ("hello" + ("foo bar" + ("a" + v4))))
                                default:
                                    v157 := GenS23{F0: 2, F1: (v1 % (((v1 ^ genFn0(365))) | 1))}
                                    _ = v157
                                }
                            } else {
                                if false {
                                    if false {
                                        var v158 uint64 = 1
                                        var v159 int64 = int64(fi78)
                                        _ = v158
                                        _ = v159
                                    }
                                } else {
                                    v160 := GenS10{F0: 2}
                                    var v161 uint16 = 0
                                    v2 = ((v4 + ("\n\t" + (v2 + ("a" + v2)))) + "hello")
                                    _ = v160
                                    _ = v161
                                }
                                var v162 uint8 = 1
                                v162 += v162
                                if (0 < -1) {
                                    var v163 int16 = 0
                                    var v164 string = (v4 + v2)
                                    _ = v163
                                    _ = v164
                                }
                                fmt.Println(v4)
                                var v165 uint8 = (v162 << ((1) & 63))
                                _ = v162
                                _ = v165
                            }
                            v166 := []int{-2147483648, fi79, (fi78 ^ 127), -1}
                            v4 = ""
                            _ = v166
                        }
                    } else {
                        var v167 uint = 2
                        if true {
                            v1 -= v1
                            v2 = v4
                            v167 = (^(v167))
                        } else {
                            if false {
                                v1 = (v1 ^ v1)
                                delete(v5, (fi79 << (((^(v167))) & 63)))
                            }
                        }
                        if false {
                            if true {
                                var v168 uint16 = 1
                                _ = v168
                            } else {
                                var v169 bool = false
                                fmt.Println(fi79)
                                _ = v169
                            }
                            switch v2 {
                            case "":
                                fmt.Println(v2, v2)
                            }
                            _ = len(v3)
                        }
                        _ = v167
                    }
                    var v170 string = "foo bar"
                    _ = v170
                }
            }
            switch "a" {
            case "\n\t":
                delete(v5, 32767)
            case "foo bar":
                if true {
                    var v171 uint8 = 128
                    v2 = v4
                    _ = v171
                } else {
                    switch v2 {
                    case "\n\t":
                        v172 := GenS24{F0: -9}
                        _ = v172
                    default:
                        v173 := v5[1]
                        _ = v173
                    }
                    var v174 uint16 = 0
                    _ = v174
                }
            case "hello":
                switch ((v4 + "\n\t") + v4) {
                case "\n\t":
                    v1 = (v1 - genFn0(656))
                default:
                    var v175 uint8 = 128
                    _ = v175
                }
                fmt.Println(fi79, v4)
            default:
                if ("a" == "") {
                    var v176 int32 = 1
                    _ = v176
                }
            }
            _ = len(v5)
            _ = fi79
        }
        v177 := GenS2{F0: -670, F1: 127, F2: 0}
        _ = fi78
        _ = v177
    }
    v1 -= v1
    if (false || false) {
        v4 = v4
    } else {
        v178 := GenS16{F0: (v1 & v1), F1: 1}
        _ = v178
    }
    delete(v3, 42)
    v1 -= (v1 * genFn0(255))
    for fi179 := 0; fi179 < 8; fi179++ {
        v2 = (v2 + v2)
        _ = fi179
    }
    v180 := GenS20{F0: 1}
    var v181 float32 = -2.25
    v182 := []int{806, -138}
    if true {
        var v183 uint64 = 0
        _ = v183
    } else {
        v184 := GenS25{F0: (v1 << ((1) & 63)), F1: 1, F2: -1}
        if false {
            var v185 int8 = (v1 | (v1 * v1))
            switch 128 {
            case 0:
                v4 = (v4 + "a")
            case 4294967295:
                var v186 int16 = -1
                _ = v186
            default:
                var v187 uint64 = 549
                _ = v187
            }
            _ = v185
        } else {
            for fi188 := 0; fi188 < 16; fi188++ {
                for fi189 := 0; fi189 < 11; fi189++ {
                    v190 := GenS2{F0: -1, F1: v1, F2: 1}
                    _ = fi189
                    _ = v190
                }
                var v191 uint8 = 73
                v192 := GenS26{F0: 2147483647}
                _ = fi188
                _ = v191
                _ = v192
            }
            switch 128 {
            case -9223372036854775808:
                fmt.Println(v4)
            default:
                switch (((v2 + ((((v4 + v2) + "\n\t") + (v4 + v4)) + ("" + v2))) + "hello") + "\n\t") {
                case "hello":
                    v193 := GenS25{F0: (v1 + genFn0(18446744073709551615)), F1: 1, F2: 0}
                    _ = v193
                default:
                    v2 = (v2 + v2)
                }
            }
            v2 = v4
        }
        _ = v184
    }
    fmt.Println(v2, v2)
    var v194 uint = 65535
    v2 = "hello"
    var v195 int8 = genFn0((v194 & v194))
    fmt.Println("v1:", v1)
    fmt.Println("v181:", v181)
    fmt.Println("slicelen v182:", len(v182), v182)
    fmt.Println("v194:", v194)
    fmt.Println("v195:", v195)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v180
    _ = v181
    _ = v182
    _ = v194
    _ = v195
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
