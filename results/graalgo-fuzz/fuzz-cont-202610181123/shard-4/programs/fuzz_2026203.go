package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int }
type GenS2 struct { F0 uint8; F1 uint16; F2 uint32 }
type GenS3 struct { F0 int; F1 uint16; F2 bool }
type GenS4 struct { F0 uint; F1 int16 }
type GenS5 struct { F0 uint; F1 int16; F2 uint8 }
type GenS6 struct { F0 string; F1 int8 }
type GenS7 struct { F0 int; F1 string }
type GenS8 struct { F0 uint64; F1 int16 }
type GenS9 struct { F0 int64 }
type GenS10 struct { F0 string; F1 int16; F2 uint64 }
type GenS11 struct { F0 string }
type GenS12 struct { F0 uint }
type GenS13 struct { F0 int64; F1 int32; F2 int32 }
func genFn0(p0 uint8, p1 int8, p2 int8) int64 {
    if ((p0 * p0) != (p0 >> ((710) & 63))) {
        if (42 <= -22) {
            v7 := GenS1{F0: 65535, F1: -2147483648}
            for fi8 := 0; fi8 < 6; fi8++ {
                for fi9 := 0; fi9 < 6; fi9++ {
                    if (("\n\t" > (("\n\t" + "hello") + ("hello" + "hello"))) || false) {
                        fmt.Println(fi9)
                        v10 := []int{-620, fi9}
                        _ = v10
                    } else {
                        p2 = (p2 * p2)
                        fmt.Println(fi9)
                    }
                    var v11 int32 = 497
                    _ = fi9
                    _ = v11
                }
                _ = fi8
            }
            _ = v7
        }
        var v12 string = "a"
        _ = v12
    }
    if false {
        v13 := []int{-1, 4294967295, -2147483648}
        _ = v13
    }
    p0 -= 42
    for fi14 := 0; fi14 < 10; fi14++ {
        for fi15 := 0; fi15 < 1; fi15++ {
            p0 -= p0
            _ = fi15
        }
        v16 := []int{fi14, fi14}
        switch ("foo bar" + "\n\t") {
        case "":
            if (1 == p0) {
                switch ("" + "hello") {
                case "foo bar":
                    var v17 bool = true
                    _ = v17
                default:
                    if (!(true && (44 != (p0 + (p0 + 242))))) {
                        v18 := GenS1{F0: 431, F1: fi14}
                        _ = v18
                    }
                }
            } else {
                p2 = p2
                fmt.Println(p1, p2, p2)
                v19 := GenS2{F0: (p0 * p0), F1: 42, F2: 266}
                _ = v19
            }
        case "hello":
            for fi20 := 0; fi20 < 6; fi20++ {
                var v21 uint64 = 132
                _ = fi20
                _ = v21
            }
        }
        _ = fi14
        _ = v16
    }
    var v22 uint16 = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v22
    return 0
}
func genFn1(p0 uint32) uint16 {
    var v23 uint64 = 593
    fmt.Println(p0, v23)
    for fi24 := 0; fi24 < 13; fi24++ {
        p0 = (p0 - (p0 / ((42) | 1)))
        _ = fi24
    }
    var v25 int = -2147483648
    switch "hello" {
    case "":
        v23 = (^(v23))
    }
    _ = p0
    _ = v23
    _ = v25
    return 32767
}

func main() {
    v1 := make(map[int]uint32)
    var v2 int32 = 231
    var v3 uint32 = 2147483647
    var v4 uint64 = 4294967295
    v5 := make(map[uint32]string)
    var v6 int32 = v2
    v6 -= v2
    v6 += v2
    for fi26 := 0; fi26 < 10; fi26++ {
        v3 += 518
        v2 = (v6 ^ -2147483648)
        var v27 int16 = 32767
        _ = fi26
        _ = v27
    }
    for fi28 := 0; fi28 < 11; fi28++ {
        v29 := GenS3{F0: (fi28 ^ fi28), F1: 2, F2: ((^(v4)) <= v4)}
        if (!(42 == 0)) {
            v30 := []int{}
            var v31 uint8 = uint8(v4)
            switch int(v2) {
            case -351:
                v32 := GenS1{F0: 18446744073709551615, F1: (fi28 / ((fi28) | 1))}
                _ = v32
            default:
                if (!((v31 <= v31) || (1 > -971))) {
                    v4 += 42
                    _ = len(v5)
                }
            }
            _ = v30
            _ = v31
        }
        if false {
            var v33 uint64 = v4
            _ = v33
        }
        _ = fi28
        _ = v29
    }
    v34 := GenS3{F0: -143, F1: 1, F2: (1 < 42)}
    if (217 != 127) {
        switch -9223372036854775808 {
        case -9223372036854775808:
            var v35 int32 = (v2 - v2)
            _ = v35
        default:
            var v36 float32 = (1.5 * -2.25)
            _ = v36
        }
    }
    switch "hello" {
    case "hello":
        for fi37 := 0; fi37 < 18; fi37++ {
            if true {
                v3 -= (v3 % ((4294967295) | 1))
            }
            _ = fi37
        }
    default:
        _ = len(v1)
    }
    var v38 uint16 = 181
    if (float64(1e10) < 3.14159) {
        if (genFn1((v3 * (v3 * 42))) != (v38 + genFn1(v3))) {
            v3 = (v3 % (((^(v3))) | 1))
            var v39 int8 = -1
            _ = v39
        } else {
            for fi40 := 0; fi40 < 12; fi40++ {
                v2 += 1
                v3 = (^(v3))
                for fi41 := 0; fi41 < 20; fi41++ {
                    fmt.Println(v3)
                    switch fi41 {
                    case 128:
                        if ((v4 / ((1) | 1)) < (v4 ^ (v4 - v4))) {
                            v42 := v5[(^(v3))]
                            v43 := GenS1{F0: 4294967295, F1: 9223372036854775807}
                            fmt.Println(v4, v4)
                            _ = v42
                            _ = v43
                        } else {
                            var v44 string = (("foo bar" + ("a" + ("hello" + ""))) + ("" + "\n\t"))
                            var v45 uint64 = v4
                            _ = v44
                            _ = v45
                        }
                    case -2147483648:
                        var v46 int32 = -1
                        _ = v46
                    case 4294967295:
                        if ((^(v4)) != (v4 & 232)) {
                            fmt.Println(v3, v6)
                            var v47 float64 = (0.1 - 1.5)
                            v6 = 886
                            _ = v47
                        }
                        var v48 int16 = -32768
                        v6 += v6
                        _ = v48
                    default:
                        var v49 uint = 65535
                        _ = v49
                    }
                    v50 := GenS4{F0: 67, F1: -752}
                    _ = fi41
                    _ = v50
                }
                _ = fi40
            }
        }
        delete(v1, 4294967295)
        if true {
            for fi51 := 0; fi51 < 20; fi51++ {
                if (float64(1.5) <= -2.25) {
                    v52 := []int{(fi51 - fi51)}
                    v53 := GenS1{F0: 65535, F1: fi51}
                    _ = v52
                    _ = v53
                }
                if (false && (false && false)) {
                    if (true && ((v38 << ((42) & 63)) < (v38 << ((4294967295) & 63)))) {
                        v4 = v4
                        switch "" {
                        case "\n\t":
                            if ((100 != 100) || false) {
                                v2 = (v2 << ((1) & 63))
                                for fi54 := 0; fi54 < 9; fi54++ {
                                    v38 = (v38 | genFn1(943))
                                    _ = fi54
                                }
                            } else {
                                switch (fi51 - fi51) {
                                case 270:
                                    v55 := GenS5{F0: uint(v4), F1: 32767, F2: 2}
                                    _ = v55
                                case 42:
                                    var v56 bool = false
                                    _ = v56
                                }
                            }
                        default:
                            for fi57 := 0; fi57 < 12; fi57++ {
                                var v58 uint8 = 1
                                _ = fi57
                                _ = v58
                            }
                        }
                    } else {
                        switch 2147483647 {
                        case 32768:
                            if true {
                                switch ("" + "hello") {
                                case "hello":
                                    if true {
                                        var v59 bool = false
                                        v4 += (v4 ^ 251)
                                        _ = v59
                                    } else {
                                        for fi60 := 0; fi60 < 5; fi60++ {
                                            switch fi60 {
                                            case 1:
                                                if (v38 > 42) {
                                                    v61 := []int{-1, (-(fi51)), (^(fi51))}
                                                    var v62 float64 = 0.0
                                                    _ = v61
                                                    _ = v62
                                                } else {
                                                    v2 += 0
                                                    v3 = 65535
                                                    fmt.Println(v4, v3)
                                                }
                                            case -49:
                                                if (1 >= 2) {
                                                    if true {
                                                        v63 := GenS4{F0: 2, F1: 938}
                                                        v64 := []int{}
                                                        v65 := GenS6{F0: "a", F1: 42}
                                                        _ = v63
                                                        _ = v64
                                                        _ = v65
                                                    } else {
                                                        if true {
                                                            v6 -= 120
                                                        } else {
                                                            fmt.Println(v38)
                                                            v2 -= (v2 + v2)
                                                        }
                                                        switch -2147483648 {
                                                        case 233:
                                                            var v66 int64 = 383
                                                            _ = v66
                                                        case 722:
                                                            switch "a" {
                                                            case "\n\t":
                                                                fmt.Println(v4)
                                                            case "foo bar":
                                                                v67 := GenS7{F0: fi60, F1: ""}
                                                                _ = v67
                                                            default:
                                                                v38 = v38
                                                            }
                                                        default:
                                                            var v68 uint = 255
                                                            _ = v68
                                                        }
                                                        v4 -= v4
                                                    }
                                                    v3 -= v3
                                                    _ = len(v5)
                                                }
                                            case 128:
                                                var v69 bool = true
                                                _ = v69
                                            }
                                            v3 = 526
                                            switch int(v2) {
                                            case 9223372036854775807:
                                                v70 := GenS8{F0: (v4 * 255), F1: -32768}
                                                _ = v70
                                            default:
                                                v71 := []int{}
                                                v72 := GenS6{F0: "hello", F1: -116}
                                                var v73 int32 = v2
                                                _ = v71
                                                _ = v72
                                                _ = v73
                                            }
                                            _ = fi60
                                        }
                                        fmt.Println(v38, fi51, v38)
                                    }
                                default:
                                    v74 := GenS9{F0: -9223372036854775808}
                                    _ = v74
                                }
                                switch fi51 {
                                case -1:
                                    v6 = (-(v6))
                                }
                            } else {
                                if false {
                                    for fi75 := 0; fi75 < 14; fi75++ {
                                        switch fi51 {
                                        case -682:
                                            if (!false) {
                                                v6 = v6
                                            } else {
                                                var v76 uint32 = (^(v3))
                                                fmt.Println(fi75, v3, v76)
                                                var v77 uint = 4294967295
                                                _ = v76
                                                _ = v77
                                            }
                                        case 127:
                                            _ = len(v1)
                                        case -1:
                                            v3 = v3
                                        default:
                                            var v78 bool = (613 == 255)
                                            _ = v78
                                        }
                                        if (float32((-2.25 - 0.0)) == 0.0) {
                                            if false {
                                                v38 += (^(v38))
                                                switch (fi75 % ((128) | 1)) {
                                                case 178:
                                                    switch 379 {
                                                    case -2147483648:
                                                        v4 = v4
                                                    }
                                                case 9223372036854775807:
                                                    if true {
                                                        v79 := GenS8{F0: 459, F1: 32767}
                                                        v4 = 882
                                                        fmt.Println(v38, v38, v6)
                                                        v6 += (v6 << ((uint(v4)) & 63))
                                                        _ = v79
                                                    } else {
                                                        v4 += 2
                                                    }
                                                case 32768:
                                                    v80 := []int{fi75, fi51, (fi51 | 1)}
                                                    _ = v80
                                                }
                                                var v81 int8 = -128
                                                _ = v81
                                            }
                                            v2 = (v6 % ((int32(fi75)) | 1))
                                            var v82 uint32 = 4294967295
                                            _ = v82
                                        } else {
                                            if ((3.14159 * (-2.25 * (1.5 - 1e10))) != 3.14159) {
                                                delete(v1, fi51)
                                                var v83 int64 = -1
                                                if (v6 == 2147483647) {
                                                    if (((fi51 ^ -9223372036854775808) <= (fi75 | (-(fi75)))) && (484 == (v38 + 0))) {
                                                        switch "foo bar" {
                                                        case "hello":
                                                            if true {
                                                                if false {
                                                                    v84 := []int{(fi75 / ((127) | 1)), 127}
                                                                    v3 -= uint32(v38)
                                                                    _ = v84
                                                                } else {
                                                                    v85 := GenS5{F0: 224, F1: 613, F2: 30}
                                                                    _ = v85
                                                                }
                                                                switch (fi75 & fi75) {
                                                                case 42:
                                                                    if (0 >= 65535) {
                                                                        if (!false) {
                                                                            v4 += (^(v4))
                                                                            fmt.Println(v38, v38)
                                                                            switch "" {
                                                                            case "hello":
                                                                                switch fi51 {
                                                                                case 9223372036854775807:
                                                                                    var v86 int32 = 42
                                                                                    _ = v86
                                                                                default:
                                                                                    var v87 uint16 = v38
                                                                                    _ = v87
                                                                                }
                                                                            }
                                                                        } else {
                                                                            v3 += v3
                                                                            v2 -= (v2 + v2)
                                                                            v88 := GenS10{F0: "a", F1: -32768, F2: (v4 + (v4 & 42))}
                                                                            _ = v88
                                                                        }
                                                                        fmt.Println(v38, v38)
                                                                        v6 += int32(v3)
                                                                    }
                                                                case -661:
                                                                    var v89 uint16 = 121
                                                                    _ = v89
                                                                case 32767:
                                                                    v90 := GenS11{F0: ""}
                                                                    _ = v90
                                                                }
                                                                var v91 uint64 = (^(v4))
                                                                _ = v91
                                                            }
                                                        case "\n\t":
                                                            v92 := v1[(-(fi75))]
                                                            fmt.Println(v3)
                                                            fmt.Println(v3)
                                                            var v93 uint8 = 187
                                                            _ = v92
                                                            _ = v93
                                                        case "a":
                                                            v4 = (v4 + (v4 - 450))
                                                        default:
                                                            v6 += (-(v2))
                                                        }
                                                        var v94 bool = true
                                                        _ = v94
                                                    } else {
                                                        v6 = -2147483648
                                                        v95 := GenS12{F0: 18446744073709551615}
                                                        _ = v95
                                                    }
                                                    switch fi51 {
                                                    case 127:
                                                        var v96 uint32 = (v3 * 1)
                                                        _ = v96
                                                    case -962:
                                                        v38 = 65535
                                                    case 32768:
                                                        var v97 int64 = (v83 / ((v83) | 1))
                                                        _ = v97
                                                    }
                                                    if false {
                                                        fmt.Println(v2, v38)
                                                        if false {
                                                            v98 := v5[(v3 << ((uint(v3)) & 63))]
                                                            v2 = v2
                                                            _ = v98
                                                        } else {
                                                            var v99 uint16 = (v38 << ((676) & 63))
                                                            v2 = (v2 << ((42) & 63))
                                                            _ = v99
                                                        }
                                                    }
                                                }
                                                _ = v83
                                            } else {
                                                fmt.Println(v3, fi51)
                                            }
                                            if (0 != (v6 | (v6 % (((-(v6))) | 1)))) {
                                                var v100 uint16 = (v38 - v38)
                                                var v101 int32 = (v6 - 2147483647)
                                                v102 := v5[1]
                                                _ = v100
                                                _ = v101
                                                _ = v102
                                            }
                                            var v103 uint16 = v38
                                            _ = v103
                                        }
                                        var v104 uint8 = 248
                                        _ = fi75
                                        _ = v104
                                    }
                                } else {
                                    if true {
                                        fmt.Println(v6)
                                    } else {
                                        v1[(^(fi51))] = (v3 | 0)
                                        var v105 uint32 = (v3 & (v3 / ((v3) | 1)))
                                        _ = v105
                                    }
                                    switch (-(fi51)) {
                                    case 2147483647:
                                        v4 -= 607
                                    case 128:
                                        var v106 int16 = 42
                                        _ = v106
                                    case 677:
                                        v38 -= v38
                                    default:
                                        if ((0.0 == 3.14159) || false) {
                                            switch "" {
                                            case "":
                                                var v107 int = fi51
                                                _ = v107
                                            default:
                                                v3 = (^(v3))
                                            }
                                            v108 := GenS6{F0: "", F1: 42}
                                            _ = v108
                                        } else {
                                            v109 := GenS10{F0: "foo bar", F1: -1, F2: v4}
                                            fmt.Println(v6)
                                            _ = v109
                                        }
                                    }
                                    for fi110 := 0; fi110 < 10; fi110++ {
                                        v3 += 0
                                        var v111 uint64 = (^(v4))
                                        _ = fi110
                                        _ = v111
                                    }
                                }
                                v112 := []int{(fi51 ^ -23), fi51, 128, fi51}
                                v113 := GenS7{F0: fi51, F1: "hello"}
                                _ = v112
                                _ = v113
                            }
                        case -2147483648:
                            v114 := []int{fi51, -9223372036854775808, (fi51 * fi51), (fi51 ^ fi51)}
                            _ = v114
                        case 2147483647:
                            var v115 int64 = genFn0(0, -100, 100)
                            _ = v115
                        }
                        v116 := GenS10{F0: ("\n\t" + "a"), F1: -265, F2: v4}
                        v117 := GenS3{F0: 283, F1: 42, F2: (-760 <= 0)}
                        _ = v116
                        _ = v117
                    }
                    if (-458 < 1) {
                        v118 := []int{fi51, fi51}
                        v119 := GenS13{F0: 42, F1: (v6 + v2), F2: int32(v4)}
                        _ = v118
                        _ = v119
                    } else {
                        var v120 bool = false
                        var v121 int16 = 827
                        _ = v120
                        _ = v121
                    }
                }
                for fi122 := 0; fi122 < 11; fi122++ {
                    fmt.Println(fi51, fi122)
                    v4 -= v4
                    var v123 int32 = v6
                    v2 = (v123 ^ v6)
                    v124 := GenS12{F0: 65535}
                    _ = fi122
                    _ = v123
                    _ = v124
                }
                _ = fi51
            }
            v6 += v2
        }
    } else {
        v125 := GenS2{F0: 128, F1: (v38 & (v38 ^ (v38 | genFn1(2147483648)))), F2: v3}
        _ = v125
    }
    var v126 int8 = -39
    if true {
        v127 := GenS9{F0: 42}
        _ = v127
    }
    v128 := GenS10{F0: "", F1: -32768, F2: (v4 - (v4 << ((42) & 63)))}
    var v129 uint8 = 241
    fmt.Println(v38)
    var v130 float64 = 1e10
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v126:", v126)
    fmt.Println("v129:", v129)
    fmt.Println("v130:", v130)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v126
    _ = v128
    _ = v129
    _ = v130
    _ = v2
    _ = v3
    _ = v34
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
