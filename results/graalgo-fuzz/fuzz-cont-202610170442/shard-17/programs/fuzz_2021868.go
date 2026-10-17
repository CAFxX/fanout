package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint; F1 bool }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint; F1 uint8 }
type GenS6 struct { F0 uint64; F1 uint8; F2 uint64 }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 uint8; F1 int32 }
type GenS9 struct { F0 uint8; F1 string; F2 int64 }
type GenS10 struct { F0 string }
type GenS11 struct { F0 int16; F1 uint16 }
func genFn0(p0 uint, p1 int32) int {
    var v7 uint16 = 1
    if (582 < 1) {
        switch 658 {
        case -299:
            for fi8 := 0; fi8 < 17; fi8++ {
                var v9 int32 = p1
                _ = fi8
                _ = v9
            }
        case -2147483648:
            for fi10 := 0; fi10 < 9; fi10++ {
                var v11 uint = 1
                switch "hello" {
                case "":
                    v12 := GenS1{F0: 529, F1: (^(p1))}
                    _ = v12
                case "foo bar":
                    if (-1 > -128) {
                        var v13 int16 = -1
                        var v14 uint16 = v7
                        var v15 int16 = 0
                        _ = v13
                        _ = v14
                        _ = v15
                    }
                case "a":
                    var v16 string = ""
                    _ = v16
                default:
                    var v17 uint16 = v7
                    _ = v17
                }
                _ = fi10
                _ = v11
            }
        case 1:
            var v18 int16 = -32768
            _ = v18
        }
        v19 := GenS1{F0: 2, F1: p1}
        _ = v19
    } else {
        for fi20 := 0; fi20 < 20; fi20++ {
            if false {
                v21 := []int{}
                for fi22 := 0; fi22 < 8; fi22++ {
                    if len(v21) >= 2 {
                        v21 = v21[1:len(v21)]
                    }
                    _ = fi22
                }
                for fi23 := 0; fi23 < 5; fi23++ {
                    var v24 int8 = 1
                    _ = fi23
                    _ = v24
                }
                _ = v21
            } else {
                v25 := []int{(^(fi20)), (fi20 << ((2) & 63)), fi20, 4294967295}
                v25 = append(v25, (fi20 ^ (fi20 - fi20)))
                _ = v25
            }
            _ = fi20
        }
        fmt.Println(p0, p1)
    }
    var v26 int32 = p1
    _ = p0
    _ = p1
    _ = v26
    _ = v7
    return -864
}
func genFn1() int8 {
    v27 := GenS1{F0: 4294967295, F1: 0}
    var v28 float32 = (1e10 - (1.5 + 3.14159))
    for fi29 := 0; fi29 < 5; fi29++ {
        switch "a" {
        case "":
            v30 := GenS1{F0: 42, F1: 0}
            _ = v30
        default:
            for fi31 := 0; fi31 < 5; fi31++ {
                var v32 bool = (333 >= 0)
                if (2 < 1) {
                    fmt.Println(fi31)
                    switch "foo bar" {
                    case "\n\t":
                        switch "" {
                        case "":
                            v33 := []int{fi31, (fi31 * (fi29 * 0))}
                            _ = v33
                        case "foo bar":
                            v32 = v32
                        }
                    case "":
                        v34 := GenS1{F0: 1, F1: -637}
                        _ = v34
                    }
                    if (919 > 2) {
                        v32 = v32
                        v35 := []int{(fi29 - (fi31 >> ((42) & 63)))}
                        _ = v35
                    } else {
                        var v36 bool = v32
                        if (42 == 42) {
                            var v37 int64 = -1
                            _ = v37
                        }
                        _ = v36
                    }
                    var v38 uint8 = 242
                    _ = v38
                } else {
                    v39 := GenS1{F0: 42, F1: 2147483647}
                    v32 = false
                    _ = v39
                }
                _ = fi31
                _ = v32
            }
        }
        switch "\n\t" {
        case "a":
            for fi40 := 0; fi40 < 6; fi40++ {
                var v41 float32 = 1.5
                v42 := []int{fi40, (fi40 << ((498) & 63)), fi29}
                _ = fi40
                _ = v41
                _ = v42
            }
        case "foo bar":
            if ((127 != 0) || false) {
                for fi43 := 0; fi43 < 8; fi43++ {
                    if (2 > 338) {
                        var v44 uint16 = 255
                        v44 -= (v44 - v44)
                        switch 2147483647 {
                        case 32768:
                            v44 = (v44 - v44)
                        case 127:
                            v45 := GenS1{F0: 225, F1: -1}
                            _ = v45
                        case 4294967295:
                            var v46 int64 = 0
                            _ = v46
                        }
                        var v47 int32 = 1
                        v47 -= v47
                        v48 := []int{-904, (fi43 - 127), fi29, fi43}
                        _ = v44
                        _ = v47
                        _ = v48
                    } else {
                        v28 = ((((v28 * ((-2.25 + v28) - 0.0)) + (float32(-2.25) * v28)) + -2.25) * v28)
                        var v49 int64 = -1
                        v28 = v28
                        _ = v49
                    }
                    v50 := GenS1{F0: 18446744073709551615, F1: 0}
                    _ = fi43
                    _ = v50
                }
            }
        default:
            v51 := []int{(fi29 & (fi29 / ((genFn0(65535, 0)) | 1))), fi29, 418}
            _ = v51
        }
        _ = fi29
    }
    v28 += (-2.25 / 1.5)
    _ = v27
    _ = v28
    return -12
}

func main() {
    var v1 uint64 = 1
    var v2 float64 = 1e10
    var v3 int = 127
    v4 := []uint32{524, 42}
    v5 := make(map[int]string)
    var v6 int16 = 0
    v3 -= v3
    v52 := GenS1{F0: v1, F1: 988}
    switch "hello" {
    case "foo bar":
        fmt.Println(v3, v1)
    case "":
        for fi53 := 0; fi53 < 19; fi53++ {
            if (false || true) {
                fmt.Println(v1, v3, v3)
                fmt.Println(fi53)
                var v54 int16 = v6
                _ = v54
            }
            v55 := GenS1{F0: 1, F1: 139}
            _ = fi53
            _ = v55
        }
    }
    for fi56 := 0; fi56 < 12; fi56++ {
        switch "" {
        case "":
            if true {
                v2 += ((v2 / (v2 - 1.5)) * float64(float32((v2 / 1.0))))
                v57 := GenS1{F0: v1, F1: 0}
                v58 := GenS1{F0: v1, F1: -477}
                _ = v57
                _ = v58
            }
        default:
            v59 := GenS1{F0: (v1 ^ (^(v1))), F1: 1}
            _ = v59
        }
        _ = fi56
    }
    v60 := GenS1{F0: v1, F1: 0}
    for fi61 := 0; fi61 < 2; fi61++ {
        fmt.Println(v2, v6)
        for fi62 := 0; fi62 < 10; fi62++ {
            v63 := GenS1{F0: (v1 & v1), F1: -1}
            _ = fi62
            _ = v63
        }
        _ = fi61
    }
    fmt.Println(v3, v2, v1)
    if (false && (1 >= -69)) {
        v2 = 3.14159
        v64 := v5[(v3 / ((v3) | 1))]
        _ = v64
    } else {
        var v65 int32 = -625
        if len(v4) > 0 {
            v4[0] = 4294967295
        }
        v66 := v5[(-(v3))]
        _ = v65
        _ = v66
    }
    v3 = 4294967295
    if false {
        var v67 uint = 4294967295
        switch v3 {
        case 4294967295:
            fmt.Println(v6, v3, v1)
        }
        var v68 bool = true
        _ = v67
        _ = v68
    }
    if true {
        fmt.Println(v2, v3, v2)
        v3 = v3
        for fi69 := 0; fi69 < 6; fi69++ {
            v2 = float64(1.5)
            v70 := GenS2{F0: 2}
            fmt.Println(v2)
            _ = fi69
            _ = v70
        }
    }
    fmt.Println(v3, v3)
    if true {
        for fi71 := 0; fi71 < 7; fi71++ {
            if true {
                v72 := GenS2{F0: 448}
                var v73 int8 = 127
                _ = v72
                _ = v73
            } else {
                v1 += v1
            }
            switch "\n\t" {
            case "\n\t":
                fmt.Println(v3, v6)
            case "a":
                if true {
                    if (968 <= 437) {
                        v2 = -2.25
                        for fi74 := 0; fi74 < 15; fi74++ {
                            v1 += v1
                            v75 := GenS3{F0: 339, F1: false}
                            fmt.Println(v2, fi71)
                            _ = fi74
                            _ = v75
                        }
                    }
                }
            case "hello":
                v76 := GenS4{F0: "foo bar"}
                _ = v76
            }
            _ = fi71
        }
        v77 := GenS1{F0: (^(v1)), F1: 42}
        for fi78 := 0; fi78 < 9; fi78++ {
            if len(v4) > 0 {
                v4[0] = 527
            }
            v6 = v6
            _ = fi78
        }
        _ = v77
    } else {
        v79 := GenS4{F0: "a"}
        for fi80 := 0; fi80 < 12; fi80++ {
            v81 := GenS1{F0: (v1 / ((2) | 1)), F1: 2147483647}
            var v82 float32 = 3.14159
            v83 := v5[v3]
            _ = fi80
            _ = v81
            _ = v82
            _ = v83
        }
        v84 := GenS4{F0: (("a" + ("\n\t" + ("a" + ("a" + ("a" + "foo bar"))))) + (("foo bar" + "\n\t") + ("\n\t" + "")))}
        _ = v79
        _ = v84
    }
    var v85 int16 = (v6 << ((42) & 63))
    for fi86 := 0; fi86 < 1; fi86++ {
        delete(v5, 730)
        if false {
            if len(v4) > 0 {
                v4[0] = 0
            }
            var v87 float32 = ((((0.0 + ((1.5 - 0.1) * ((0.0 * (3.14159 * (3.14159 - (float32(((1.5 / (v2 / ((v2 + v2) * float64(float32((-2.25 + 0.1)))))) - -2.25)) / 1.5)))) + (0.0 / (((0.0 + 0.0) - (-2.25 / 1.0)) - (float32(3.14159) - float32((v2 / (v2 * v2))))))))) + 1e10) - ((1e10 + (0.0 - (1e10 / 1.5))) / 0.1)) + (-2.25 + 0.1))
            v4 = append(v4, 590)
            _ = v87
        } else {
            if ((-(v3)) <= (v3 % (((-(fi86))) | 1))) {
                fmt.Println(v85, v85, v85)
                if false {
                    if false {
                        v88 := GenS5{F0: 393, F1: 1}
                        _ = v88
                    } else {
                        v2 += (v2 / v2)
                    }
                    if true {
                        switch (("\n\t" + "") + "foo bar") {
                        case "":
                            if false {
                                for fi89 := 0; fi89 < 14; fi89++ {
                                    v1 += v1
                                    var v90 string = ("a" + "\n\t")
                                    if (false && (255 == 481)) {
                                        var v91 float64 = v2
                                        var v92 string = "foo bar"
                                        _ = v91
                                        _ = v92
                                    }
                                    v1 = v1
                                    fmt.Println(fi89, v6, fi86)
                                    _ = fi89
                                    _ = v90
                                }
                            } else {
                                var v93 bool = false
                                var v94 int32 = -1
                                _ = v93
                                _ = v94
                            }
                        case "\n\t":
                            var v95 int8 = 100
                            _ = v95
                        case "a":
                            if len(v4) > 0 {
                                v4[0] = uint32(v1)
                            }
                        }
                        v3 = -9223372036854775808
                    }
                    if false {
                        switch v3 {
                        case -9223372036854775808:
                            v96 := GenS6{F0: v1, F1: 2, F2: 968}
                            _ = v96
                        }
                        fmt.Println(v6)
                    }
                }
                if false {
                    var v97 int32 = 42
                    _ = v97
                }
            } else {
                for fi98 := 0; fi98 < 3; fi98++ {
                    v99 := GenS7{F0: 1}
                    _ = fi98
                    _ = v99
                }
            }
            for fi100 := 0; fi100 < 16; fi100++ {
                v3 = (v3 * -9223372036854775808)
                _ = len(v4)
                switch ("a" + "foo bar") {
                case "hello":
                    v1 = v1
                case "":
                    if (v6 < v85) {
                        var v101 uint8 = 128
                        _ = v101
                    } else {
                        v102 := GenS7{F0: 42}
                        v103 := GenS3{F0: 2, F1: (genFn0(361, -1) >= -190)}
                        v85 -= (v85 & 1)
                        v104 := GenS4{F0: (("a" + "foo bar") + "hello")}
                        _ = v102
                        _ = v103
                        _ = v104
                    }
                case "\n\t":
                    fmt.Println(fi100, fi100)
                default:
                    if true {
                        if (v85 >= (v6 - 0)) {
                            if true {
                                if false {
                                    delete(v5, (v3 & (-(v3))))
                                    if (-7 == v3) {
                                        var v105 int16 = 0
                                        v106 := GenS3{F0: 255, F1: (true && true)}
                                        var v107 string = "hello"
                                        _ = v105
                                        _ = v106
                                        _ = v107
                                    }
                                } else {
                                    switch "" {
                                    case "foo bar":
                                        fmt.Println(fi86, fi100)
                                    case "hello":
                                        v6 -= 1
                                    default:
                                        var v108 string = "foo bar"
                                        switch fi100 {
                                        case -2147483648:
                                            var v109 uint32 = 1
                                            _ = v109
                                        case -561:
                                            if (2 >= 1) {
                                                var v110 int = fi100
                                                v1 -= (^(v1))
                                                v108 = (v108 + v108)
                                                _ = v110
                                            } else {
                                                v111 := GenS6{F0: v1, F1: 9, F2: v1}
                                                v112 := GenS8{F0: 1, F1: 477}
                                                if true {
                                                    if len(v4) > 0 {
                                                        v4[0] = 0
                                                    }
                                                    _ = len(v5)
                                                    if (2 < (v1 >> ((0) & 63))) {
                                                        var v113 uint32 = 2
                                                        v108 = v108
                                                        switch 32768 {
                                                        case 26:
                                                            if (977 != 4294967295) {
                                                                v113 = (v113 / (((v113 ^ v113)) | 1))
                                                            }
                                                        case -1:
                                                            v6 += -32768
                                                        default:
                                                            var v114 uint32 = (v113 | 845)
                                                            _ = v114
                                                        }
                                                        _ = v113
                                                    }
                                                }
                                                _ = v111
                                                _ = v112
                                            }
                                        case -614:
                                            if ((128 >= 1) && (true && false)) {
                                                var v115 int16 = (^(v85))
                                                _ = v115
                                            } else {
                                                var v116 uint8 = 2
                                                _ = v116
                                            }
                                        }
                                        _ = v108
                                    }
                                    _ = len(v4)
                                }
                                if (-2.25 <= 1.5) {
                                    v117 := GenS5{F0: 332, F1: 42}
                                    if (((0.0 < 1.5) || (!(2 >= 128))) && true) {
                                        if len(v4) > 0 {
                                            v4[0] = 65535
                                        }
                                        var v118 int = v3
                                        var v119 int32 = 0
                                        _ = v118
                                        _ = v119
                                    } else {
                                        var v120 uint = 1
                                        _ = v120
                                    }
                                    fmt.Println(v85)
                                    _ = v117
                                }
                                v2 = (v2 / v2)
                                v4 = append(v4, 2)
                                v121 := GenS6{F0: v1, F1: 106, F2: (v1 + 4294967295)}
                                v122 := GenS9{F0: 127, F1: "\n\t", F2: 245}
                                _ = v121
                                _ = v122
                            }
                            v2 = 3.14159
                            v123 := GenS2{F0: 42}
                            v124 := v5[fi86]
                            _ = v123
                            _ = v124
                        }
                    }
                }
                _ = fi100
            }
        }
        v1 = (v1 * v1)
        _ = fi86
    }
    for fi125 := 0; fi125 < 12; fi125++ {
        if ((v3 % (((^(v3))) | 1)) < v3) {
            if true {
                var v126 float64 = ((v2 * (v2 / v2)) - (1.5 / ((3.14159 + v2) - v2)))
                if false {
                    v126 = v2
                    if (!false) {
                        for fi127 := 0; fi127 < 17; fi127++ {
                            v3 += v3
                            v128 := GenS10{F0: ("" + ("\n\t" + (("" + "\n\t") + (("foo bar" + "a") + ("\n\t" + "hello")))))}
                            _ = fi127
                            _ = v128
                        }
                        v2 = (v2 * v126)
                        for fi129 := 0; fi129 < 18; fi129++ {
                            var v130 bool = false
                            _ = fi129
                            _ = v130
                        }
                    }
                }
                switch "hello" {
                case "":
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                case "a":
                    v131 := GenS1{F0: 1, F1: -93}
                    _ = v131
                case "\n\t":
                    v3 = fi125
                default:
                    v126 -= v126
                }
                _ = v126
            }
            for fi132 := 0; fi132 < 4; fi132++ {
                v133 := GenS11{F0: v6, F1: 42}
                _ = fi132
                _ = v133
            }
            v5[(v3 * (fi125 | (fi125 * v3)))] = ""
        }
        _ = fi125
    }
    switch ("\n\t" + "foo bar") {
    case "foo bar":
        v134 := GenS8{F0: 127, F1: -2147483648}
        _ = v134
    case "\n\t":
        var v135 float32 = 0.0
        _ = v135
    case "":
        v2 -= v2
    default:
        v136 := GenS7{F0: 42}
        _ = v136
    }
    switch 0 {
    case 2147483647:
        var v137 int = (v3 & int(v1))
        _ = v137
    case 32767:
        var v138 int16 = 1
        _ = v138
    case 128:
        v2 -= (v2 - ((((1e10 - v2) / v2) / ((((v2 / v2) * float64(1.5)) - (-2.25 / (v2 - 1e10))) - float64(0.0))) + ((v2 + (1e10 / (v2 / v2))) * ((0.0 / (1.5 * (v2 * v2))) / v2))))
    default:
        fmt.Println(v6)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v85:", v85)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v52
    _ = v6
    _ = v60
    _ = v85
}
