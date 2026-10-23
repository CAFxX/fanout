package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint16; F1 uint8; F2 uint64 }
type GenS5 struct { F0 int8; F1 uint; F2 int64 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 uint64; F1 uint8 }
type GenS8 struct { F0 uint16; F1 int; F2 int32 }
type GenS9 struct { F0 string }
type GenS10 struct { F0 bool; F1 int; F2 int32 }
type GenS11 struct { F0 int; F1 int; F2 uint8 }
type GenS12 struct { F0 bool }
type GenS13 struct { F0 int8; F1 uint }
type GenS14 struct { F0 string; F1 uint32; F2 uint16 }
type GenS15 struct { F0 string; F1 bool }
type GenS16 struct { F0 uint32; F1 int64; F2 int16 }
func genFn0(p0 bool, p1 uint, p2 uint32) uint {
    p0 = true
    fmt.Println(p2)
    var v4 uint64 = 18446744073709551615
    var v5 uint16 = 21
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return p1
}
func genFn1(p0 string, p1 string) uint16 {
    for fi6 := 0; fi6 < 17; fi6++ {
        var v7 int64 = 42
        _ = fi6
        _ = v7
    }
    v8 := []int{-2147483648, 32767}
    var v9 int = 86
    p1 = p1
    _ = p0
    _ = p1
    _ = v8
    _ = v9
    return 0
}

func main() {
    var v1 string = "hello"
    v2 := make(map[uint32]string)
    var v3 uint = 349
    for fi10 := 0; fi10 < 7; fi10++ {
        if ((v1 != "hello") || (749 <= 42)) {
            v3 = (^(v3))
            if false {
                switch (-(fi10)) {
                case 1:
                    v1 = v1
                case 128:
                    var v11 int = -2147483648
                    _ = v11
                case -1:
                    var v12 uint = (v3 + 65535)
                    _ = v12
                default:
                    if (!(1 != 1)) {
                        v1 = "\n\t"
                        switch "a" {
                        case "a":
                            v13 := GenS1{F0: 42}
                            _ = v13
                        }
                        v14 := GenS2{F0: 1}
                        _ = v14
                    } else {
                        if (128 != uint8(v3)) {
                            for fi15 := 0; fi15 < 9; fi15++ {
                                v16 := GenS2{F0: 121}
                                var v17 int = fi10
                                fmt.Println(v17, v3, v3)
                                _ = fi15
                                _ = v16
                                _ = v17
                            }
                            v3 = genFn0((true || (false && true)), (^(v3)), 2147483648)
                            var v18 int32 = 2147483647
                            _ = v18
                        }
                        v19 := GenS3{F0: 32767}
                        _ = v19
                    }
                }
            } else {
                var v20 float32 = -2.25
                _ = v20
            }
            var v21 uint16 = 664
            _ = v21
        }
        _ = fi10
    }
    for fi22 := 0; fi22 < 19; fi22++ {
        var v23 string = (v1 + "hello")
        v24 := GenS4{F0: 1, F1: 42, F2: 18446744073709551615}
        var v25 int16 = -730
        _ = fi22
        _ = v23
        _ = v24
        _ = v25
    }
    switch -961 {
    case 32767:
        v3 += 1
    case 0:
        var v26 uint32 = 2
        _ = v26
    default:
        if (true && (!(-2.25 == -2.25))) {
            if false {
                switch (("foo bar" + v1) + "foo bar") {
                case "foo bar":
                    v1 = "hello"
                case "":
                    for fi27 := 0; fi27 < 15; fi27++ {
                        if ((float64(-2.25) >= -2.25) && false) {
                            _ = len(v2)
                            for fi28 := 0; fi28 < 5; fi28++ {
                                v29 := []int{(fi27 + fi28)}
                                _ = fi28
                                _ = v29
                            }
                            var v30 float64 = -2.25
                            _ = v30
                        }
                        _ = fi27
                    }
                }
                for fi31 := 0; fi31 < 18; fi31++ {
                    var v32 bool = (2 != 255)
                    switch (fi31 & (-(fi31))) {
                    case 42:
                        for fi33 := 0; fi33 < 7; fi33++ {
                            v1 = "hello"
                            _ = fi33
                        }
                    }
                    var v34 bool = v32
                    _ = fi31
                    _ = v32
                    _ = v34
                }
            } else {
                for fi35 := 0; fi35 < 17; fi35++ {
                    var v36 int8 = 127
                    for fi37 := 0; fi37 < 15; fi37++ {
                        v1 = (v1 + v1)
                        var v38 string = "foo bar"
                        v39 := []int{-2147483648, fi37, -388, fi35}
                        v40 := GenS4{F0: 42, F1: 0, F2: 661}
                        if (v1 < "foo bar") {
                            switch v1 {
                            case "\n\t":
                                switch (fi35 ^ (fi35 - -547)) {
                                case 2147483647:
                                    var v41 uint = (v3 << ((v3) & 63))
                                    _ = v41
                                case 4294967295:
                                    v42 := GenS2{F0: 42}
                                    _ = v42
                                default:
                                    fmt.Println(v3, v3, fi37)
                                }
                            default:
                                var v43 bool = true
                                _ = v43
                            }
                            switch (((v1 + v38) + "hello") + v1) {
                            case "a":
                                v44 := GenS5{F0: v36, F1: 397, F2: -219}
                                _ = v44
                            case "":
                                if len(v39) > 0 {
                                    v39[0] = 1
                                }
                            default:
                                v1 = "a"
                            }
                            var v45 bool = (!true)
                            _ = v45
                        } else {
                            fmt.Println(fi37, fi37)
                            var v46 uint64 = 2
                            _ = v46
                        }
                        _ = fi37
                        _ = v38
                        _ = v39
                        _ = v40
                    }
                    _ = fi35
                    _ = v36
                }
                if false {
                    fmt.Println(v3, v1)
                    v3 = v3
                    for fi47 := 0; fi47 < 14; fi47++ {
                        for fi48 := 0; fi48 < 19; fi48++ {
                            v1 = (v1 + v1)
                            _ = fi48
                        }
                        var v49 uint64 = 412
                        if false {
                            var v50 bool = true
                            switch (v1 + "foo bar") {
                            case "foo bar":
                                v50 = (v49 == 2)
                            case "\n\t":
                                switch fi47 {
                                case 435:
                                    var v51 uint8 = 127
                                    _ = v51
                                case 0:
                                    if false {
                                        var v52 int32 = 42
                                        _ = v52
                                    }
                                case -2147483648:
                                    if ((fi47 << (((v3 | v3)) & 63)) > (fi47 - 32767)) {
                                        if (!v50) {
                                            if false {
                                                if v50 {
                                                    switch -83 {
                                                    case 32768:
                                                        if false {
                                                            if v50 {
                                                                v49 += (v49 * 0)
                                                            } else {
                                                                v50 = (((!v50) || true) && false)
                                                            }
                                                            for fi53 := 0; fi53 < 12; fi53++ {
                                                                fmt.Println(v49, v49)
                                                                v3 += 0
                                                                fmt.Println(v49)
                                                                _ = fi53
                                                            }
                                                        } else {
                                                            var v54 float32 = 0.1
                                                            v50 = (!true)
                                                            _ = v54
                                                        }
                                                    case 2147483647:
                                                        v50 = false
                                                    case 1:
                                                        var v55 uint = v3
                                                        _ = v55
                                                    default:
                                                        fmt.Println(v1, v1, v1)
                                                    }
                                                    v49 = v49
                                                    v56 := GenS6{F0: -2147483648}
                                                    _ = v56
                                                } else {
                                                    v57 := GenS7{F0: (v49 * v49), F1: 128}
                                                    var v58 uint32 = 2147483648
                                                    fmt.Println(v50, v3, v50)
                                                    _ = v57
                                                    _ = v58
                                                }
                                                switch fi47 {
                                                case 622:
                                                    var v59 int32 = -1
                                                    _ = v59
                                                default:
                                                    for fi60 := 0; fi60 < 8; fi60++ {
                                                        v2[2147483648] = ""
                                                        _ = fi60
                                                    }
                                                }
                                                v61 := v2[511]
                                                _ = v61
                                            }
                                            switch v1 {
                                            case "":
                                                var v62 bool = true
                                                _ = v62
                                            case "hello":
                                                var v63 uint = (v3 >> ((v3) & 63))
                                                _ = v63
                                            case "\n\t":
                                                v1 = v1
                                            }
                                        }
                                    } else {
                                        v64 := GenS4{F0: uint16(v3), F1: 255, F2: uint64(v3)}
                                        delete(v2, 42)
                                        _ = v64
                                    }
                                default:
                                    v65 := GenS8{F0: 42, F1: 32768, F2: -2147483648}
                                    _ = v65
                                }
                            }
                            _ = v50
                        }
                        _ = fi47
                        _ = v49
                    }
                } else {
                    var v66 uint64 = 4294967295
                    v67 := GenS1{F0: 77}
                    for fi68 := 0; fi68 < 14; fi68++ {
                        if false {
                            for fi69 := 0; fi69 < 10; fi69++ {
                                if (1 >= 42) {
                                    if false {
                                        if (1 != 249) {
                                            switch v1 {
                                            case "":
                                                var v70 int32 = 2147483647
                                                _ = v70
                                            case "\n\t":
                                                if true {
                                                    if (42 == 128) {
                                                        v71 := GenS7{F0: 418, F1: 1}
                                                        v72 := GenS3{F0: 0}
                                                        v66 = 1
                                                        _ = v71
                                                        _ = v72
                                                    } else {
                                                        switch "hello" {
                                                        case "":
                                                            fmt.Println(fi68)
                                                        case "\n\t":
                                                            if (((!(!(3.14159 < 0.0))) || (!(false && false))) && (!false)) {
                                                                var v73 int16 = 32767
                                                                if (1e10 != (0.1 * (0.1 / ((0.1 + -2.25) / 3.14159)))) {
                                                                    v66 = 4294967295
                                                                } else {
                                                                    var v74 float64 = 0.0
                                                                    v73 = (-(v73))
                                                                    _ = v74
                                                                }
                                                                _ = v73
                                                            }
                                                            fmt.Println(fi69, v1)
                                                            v75 := GenS9{F0: (v1 + ("\n\t" + "\n\t"))}
                                                            _ = v75
                                                        }
                                                        if (1 > v66) {
                                                            var v76 uint64 = v66
                                                            var v77 uint16 = 2
                                                            delete(v2, 4294967295)
                                                            _ = v76
                                                            _ = v77
                                                        } else {
                                                            var v78 bool = false
                                                            var v79 uint8 = 2
                                                            var v80 int16 = int16(v66)
                                                            _ = v78
                                                            _ = v79
                                                            _ = v80
                                                        }
                                                        var v81 uint = (v3 / ((v3) | 1))
                                                        _ = v81
                                                    }
                                                    if true {
                                                        var v82 string = ""
                                                        _ = v82
                                                    } else {
                                                        delete(v2, 1)
                                                    }
                                                } else {
                                                    if (783 > v66) {
                                                        if (32767 != 65535) {
                                                            switch (fi68 ^ fi69) {
                                                            case 2147483647:
                                                                if true {
                                                                    switch v1 {
                                                                    case "foo bar":
                                                                        var v83 int32 = 2147483647
                                                                        _ = v83
                                                                    }
                                                                    if (2 != 65535) {
                                                                        v3 = v3
                                                                    } else {
                                                                        v84 := GenS10{F0: ((v3 & 971) != genFn0((255 != 2), v3, 4294967295)), F1: -240, F2: 1}
                                                                        v85 := v2[2147483647]
                                                                        _ = v84
                                                                        _ = v85
                                                                    }
                                                                } else {
                                                                    fmt.Println(v3, v66)
                                                                    v86 := GenS1{F0: 4294967295}
                                                                    _ = v86
                                                                }
                                                            case -385:
                                                                v66 = (v66 ^ v66)
                                                            }
                                                        } else {
                                                            fmt.Println(v3, fi69, fi69)
                                                        }
                                                    } else {
                                                        var v87 bool = (((v3 - 18446744073709551615) != genFn0((false || (!(0 > -1))), (v3 >> ((v3) & 63)), 1)) || (false || false))
                                                        _ = v87
                                                    }
                                                    v88 := GenS1{F0: 726}
                                                    fmt.Println(fi69)
                                                    _ = v88
                                                }
                                            case "foo bar":
                                                v2[791] = v1
                                                if false {
                                                    var v89 uint64 = v66
                                                    v89 = 591
                                                    _ = v89
                                                } else {
                                                    v3 += v3
                                                    v90 := GenS11{F0: fi69, F1: int(v3), F2: 225}
                                                    _ = v90
                                                }
                                            }
                                            if true {
                                                var v91 uint32 = 0
                                                v92 := []int{}
                                                v2[uint32(v66)] = v1
                                                if true {
                                                    switch "\n\t" {
                                                    case "hello":
                                                        var v93 uint8 = 0
                                                        _ = v93
                                                    }
                                                    fmt.Println(fi68, v1, v66)
                                                    v3 += v3
                                                } else {
                                                    var v94 string = (v1 + "\n\t")
                                                    switch (^(fi69)) {
                                                    case -9223372036854775808:
                                                        var v95 int16 = 1
                                                        _ = v95
                                                    case 256:
                                                        var v96 int16 = -896
                                                        fmt.Println(v3, fi68, v94)
                                                        var v97 int8 = 42
                                                        _ = v96
                                                        _ = v97
                                                    }
                                                    var v98 float32 = 0.1
                                                    if len(v92) >= 2 {
                                                        v92 = v92[1:len(v92)]
                                                    }
                                                    v1 = ""
                                                    _ = v94
                                                    _ = v98
                                                }
                                                _ = v91
                                                _ = v92
                                            }
                                            v99 := []int{-302, 4294967295, fi69}
                                            _ = v99
                                        } else {
                                            fmt.Println(fi68, fi68, v66)
                                        }
                                        v100 := GenS6{F0: 2147483647}
                                        switch v1 {
                                        case "hello":
                                            if true {
                                                if ((^(v3)) != genFn0(false, v3, 4294967295)) {
                                                    delete(v2, 4294967295)
                                                    v2[963] = v1
                                                } else {
                                                    var v101 int8 = int8(v3)
                                                    _ = v101
                                                }
                                            } else {
                                                v102 := GenS7{F0: v66, F1: 214}
                                                v66 = v66
                                                _ = v102
                                            }
                                            var v103 int = fi68
                                            _ = v103
                                        }
                                        if (1e10 != 0.0) {
                                            switch fi68 {
                                            case -364:
                                                var v104 uint32 = 190
                                                _ = v104
                                            default:
                                                var v105 bool = (true || false)
                                                _ = v105
                                            }
                                        }
                                        _ = v100
                                    } else {
                                        switch fi69 {
                                        case -2147483648:
                                            if true {
                                                var v106 string = v1
                                                var v107 uint64 = (v66 | v66)
                                                v3 -= uint(v66)
                                                _ = v106
                                                _ = v107
                                            }
                                        case 2147483647:
                                            fmt.Println(v1, fi68, v3)
                                        }
                                        var v108 int64 = -9223372036854775808
                                        fmt.Println(v66)
                                        _ = v108
                                    }
                                    switch ((v1 + "hello") + v1) {
                                    case "hello":
                                        fmt.Println(v66, v66)
                                    default:
                                        v109 := GenS8{F0: 997, F1: int(v66), F2: -2147483648}
                                        _ = v109
                                    }
                                    v1 = (v1 + v1)
                                } else {
                                    var v110 int8 = 100
                                    fmt.Println(v3, v110)
                                    if (("" + v1) > v1) {
                                        v111 := GenS12{F0: ((^(v3)) != genFn0((927 == 4294967295), 477, 65535))}
                                        _ = v111
                                    } else {
                                        var v112 int16 = 1
                                        v113 := GenS5{F0: (v110 >> (((^(v3))) & 63)), F1: v3, F2: 9223372036854775807}
                                        var v114 uint16 = 2
                                        _ = v112
                                        _ = v113
                                        _ = v114
                                    }
                                    _ = v110
                                }
                                switch (fi69 ^ (fi69 ^ fi69)) {
                                case -2147483648:
                                    v115 := GenS1{F0: 231}
                                    switch (^(fi68)) {
                                    case 128:
                                        var v116 int64 = -159
                                        _ = v116
                                    default:
                                        var v117 int32 = 1
                                        _ = v117
                                    }
                                    if false {
                                        var v118 uint = v3
                                        if (!(true || true)) {
                                            if true {
                                                v3 += 0
                                                v3 -= (v118 + genFn0((!true), (v118 | 44), 0))
                                                v66 += (v66 - v66)
                                            }
                                            v66 += (v66 >> ((v118) & 63))
                                            v3 = v3
                                        } else {
                                            var v119 bool = (true && false)
                                            _ = v119
                                        }
                                        _ = v118
                                    } else {
                                        delete(v2, 42)
                                        v66 -= v66
                                    }
                                    _ = v115
                                case 2147483647:
                                    fmt.Println(v66)
                                case 1:
                                    v1 = v1
                                default:
                                    if true {
                                        v120 := GenS13{F0: -1, F1: (v3 + (^(v3)))}
                                        v121 := GenS8{F0: 42, F1: int(v66), F2: -241}
                                        if true {
                                            if (0 != -9223372036854775808) {
                                                switch (v1 + "foo bar") {
                                                case "hello":
                                                    v1 = v1
                                                    v3 -= (v3 << ((65535) & 63))
                                                case "\n\t":
                                                    if (!true) {
                                                        v3 = 65535
                                                        fmt.Println(fi69, v3)
                                                        v122 := GenS14{F0: "", F1: 4294967295, F2: 2}
                                                        _ = v122
                                                    } else {
                                                        v3 = v3
                                                    }
                                                case "a":
                                                    v123 := GenS7{F0: v66, F1: 2}
                                                    _ = v123
                                                }
                                                switch fi68 {
                                                case 0:
                                                    v124 := []int{fi68, 2147483647}
                                                    _ = v124
                                                case 128:
                                                    var v125 uint8 = 127
                                                    _ = v125
                                                case -9223372036854775808:
                                                    fmt.Println(v1, v3)
                                                default:
                                                    if (32767 >= 797) {
                                                        var v126 int = -870
                                                        _ = v126
                                                    } else {
                                                        var v127 int = fi69
                                                        delete(v2, 2147483647)
                                                        v66 -= v66
                                                        _ = v127
                                                    }
                                                }
                                            }
                                            var v128 int = 379
                                            _ = v128
                                        }
                                        _ = v120
                                        _ = v121
                                    } else {
                                        v3 = genFn0((255 <= 0), v3, 2)
                                        var v129 int = int(v3)
                                        v3 += 18446744073709551615
                                        _ = v129
                                    }
                                }
                                _ = fi69
                            }
                            var v130 uint32 = 0
                            if (false && false) {
                                v1 = (v1 + v1)
                                for fi131 := 0; fi131 < 11; fi131++ {
                                    var v132 int64 = 928
                                    if true {
                                        var v133 uint = v3
                                        _ = v133
                                    } else {
                                        _ = len(v2)
                                        switch "" {
                                        case "hello":
                                            fmt.Println(v132, v130, v3)
                                        case "foo bar":
                                            v134 := GenS13{F0: 4, F1: genFn0(true, (v3 ^ v3), uint32(v66))}
                                            _ = v134
                                        }
                                    }
                                    _ = fi131
                                    _ = v132
                                }
                            } else {
                                var v135 bool = (((!true) && false) && true)
                                _ = v135
                            }
                            _ = v130
                        }
                        _ = fi68
                    }
                    _ = v66
                    _ = v67
                }
            }
            var v136 int = 240
            for fi137 := 0; fi137 < 14; fi137++ {
                for fi138 := 0; fi138 < 12; fi138++ {
                    v139 := GenS9{F0: ""}
                    var v140 uint = (v3 + v3)
                    if (v1 < v1) {
                        fmt.Println(v1, fi138, fi137)
                        var v141 uint64 = 2
                        v3 = v3
                        var v142 float32 = 3.14159
                        _ = v141
                        _ = v142
                    }
                    _ = len(v2)
                    _ = fi138
                    _ = v139
                    _ = v140
                }
                v136 -= 2147483647
                _ = fi137
            }
            _ = v136
        } else {
            fmt.Println(v1, v1)
            var v143 uint8 = 2
            v1 = ("foo bar" + ((v1 + ((("\n\t" + v1) + "\n\t") + v1)) + (v1 + (v1 + "hello"))))
            _ = v143
        }
    }
    v144 := GenS15{F0: "a", F1: (100 == 69)}
    v3 = v3
    for fi145 := 0; fi145 < 3; fi145++ {
        var v146 float32 = (-2.25 / (float32(1.5) * ((float32(-2.25) / 1e10) / 1e10)))
        v146 = v146
        _ = fi145
        _ = v146
    }
    v147 := GenS15{F0: v1, F1: false}
    v148 := GenS16{F0: 65535, F1: -9223372036854775808, F2: 42}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v144
    _ = v147
    _ = v148
    _ = v2
    _ = v3
}
