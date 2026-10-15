package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int8; F1 bool }
type GenS3 struct { F0 string; F1 int; F2 int16 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint32; F1 int8; F2 int8 }
type GenS6 struct { F0 int32; F1 bool; F2 uint64 }
type GenS7 struct { F0 uint32; F1 uint8 }
type GenS8 struct { F0 int }
type GenS9 struct { F0 int16 }
type GenS10 struct { F0 int8; F1 uint64; F2 string }
type GenS11 struct { F0 int64; F1 int; F2 uint16 }
type GenS12 struct { F0 int8; F1 int8; F2 int }
type GenS13 struct { F0 int8 }
type GenS14 struct { F0 uint8; F1 uint16 }
func genFn0(p0 int8) int16 {
    fmt.Println(p0)
    var v7 uint32 = 65535
    var v8 int8 = 0
    switch int(p0) {
    case 483:
        var v9 int8 = (p0 & (-(v8)))
        _ = v9
    case -1:
        v8 += (p0 * -1)
    case 32767:
        if (!true) {
            v8 -= (v8 >> ((4294967295) & 63))
            for fi10 := 0; fi10 < 13; fi10++ {
                v11 := GenS1{F0: 2}
                _ = fi10
                _ = v11
            }
        } else {
            p0 = p0
            for fi12 := 0; fi12 < 2; fi12++ {
                v8 += 127
                v13 := []int{}
                _ = fi12
                _ = v13
            }
            for fi14 := 0; fi14 < 4; fi14++ {
                v15 := []int{(-(fi14))}
                _ = fi14
                _ = v15
            }
        }
    default:
        v16 := GenS1{F0: 32767}
        _ = v16
    }
    _ = p0
    _ = v7
    _ = v8
    return 1000
}

func main() {
    var v1 uint64 = 42
    var v2 int8 = 42
    var v3 uint8 = 42
    v4 := make(map[uint32]string)
    var v5 uint32 = 535
    v6 := []uint32{65535}
    var v17 int64 = int64(v5)
    fmt.Println(v5, v1)
    var v18 bool = (false || true)
    if (v18 || v18) {
        switch -9223372036854775808 {
        case 789:
            fmt.Println(v5, v5)
        case -2147483648:
            for fi19 := 0; fi19 < 13; fi19++ {
                v20 := GenS1{F0: 42}
                v2 = -128
                var v21 int16 = 32767
                _ = fi19
                _ = v20
                _ = v21
            }
        case 4294967295:
            switch "\n\t" {
            case "a":
                for fi22 := 0; fi22 < 19; fi22++ {
                    v23 := GenS1{F0: 255}
                    _ = fi22
                    _ = v23
                }
            case "hello":
                if ((v5 / (((^(v5))) | 1)) != (v5 >> ((749) & 63))) {
                    var v24 uint8 = v3
                    for fi25 := 0; fi25 < 8; fi25++ {
                        fmt.Println(v2, v24)
                        if (0 >= 1000) {
                            v5 -= (v5 | (v5 & 772))
                            var v26 uint8 = (v3 - (v3 - (v3 / (((v24 / ((v3) | 1))) | 1))))
                            _ = v26
                        }
                        _ = fi25
                    }
                    var v27 float32 = 3.14159
                    _ = v24
                    _ = v27
                } else {
                    var v28 int64 = v17
                    for fi29 := 0; fi29 < 5; fi29++ {
                        if (!((2 <= v1) || (!true))) {
                            fmt.Println(v18, v5, v2)
                        } else {
                            var v30 int64 = 1
                            fmt.Println(v28, v5, v1)
                            _ = v30
                        }
                        _ = fi29
                    }
                    for fi31 := 0; fi31 < 5; fi31++ {
                        v3 = v3
                        _ = fi31
                    }
                    _ = v28
                }
            case "foo bar":
                if (1 < (v1 + (v1 - v1))) {
                    var v32 uint32 = v5
                    switch "a" {
                    case "\n\t":
                        var v33 uint32 = (v32 + v32)
                        _ = v33
                    case "hello":
                        switch "\n\t" {
                        case "hello":
                            switch "\n\t" {
                            case "\n\t":
                                var v34 string = ""
                                _ = v34
                            default:
                                v35 := GenS1{F0: 42}
                                _ = v35
                            }
                        case "\n\t":
                            if (9223372036854775807 < 98) {
                                if v18 {
                                    v1 = (v1 / (((v1 ^ 873)) | 1))
                                    delete(v4, (v32 - 4294967295))
                                } else {
                                    for fi36 := 0; fi36 < 3; fi36++ {
                                        for fi37 := 0; fi37 < 13; fi37++ {
                                            v38 := GenS2{F0: v2, F1: v18}
                                            v39 := GenS3{F0: "a", F1: fi36, F2: 638}
                                            if v18 {
                                                v6 = append(v6, 848)
                                                v40 := GenS4{F0: ((v5 & v5) > (^(v5)))}
                                                v41 := GenS2{F0: (v2 / (((v2 - (v2 & v2))) | 1)), F1: v18}
                                                var v42 int64 = v17
                                                _ = v40
                                                _ = v41
                                                _ = v42
                                            } else {
                                                v6 = append(v6, v32)
                                                var v43 uint32 = uint32(v3)
                                                fmt.Println(v32)
                                                _ = v43
                                            }
                                            _ = fi37
                                            _ = v38
                                            _ = v39
                                        }
                                        var v44 int16 = -1
                                        for fi45 := 0; fi45 < 18; fi45++ {
                                            v2 = (v2 - v2)
                                            var v46 int32 = int32(fi36)
                                            _ = fi45
                                            _ = v46
                                        }
                                        _ = fi36
                                        _ = v44
                                    }
                                }
                                v47 := GenS2{F0: (v2 | v2), F1: true}
                                _ = v47
                            }
                        default:
                            fmt.Println(v1, v5)
                        }
                    }
                    _ = v32
                }
            default:
                for fi48 := 0; fi48 < 6; fi48++ {
                    for fi49 := 0; fi49 < 2; fi49++ {
                        v50 := GenS2{F0: v2, F1: ((v2 % (((v2 >> ((65535) & 63))) | 1)) != 42)}
                        if v18 {
                            var v51 bool = false
                            var v52 float64 = ((float64(-2.25) * float64(float32(3.14159))) - 0.1)
                            v2 -= v2
                            _ = v51
                            _ = v52
                        } else {
                            fmt.Println(v2, v17)
                        }
                        _ = fi49
                        _ = v50
                    }
                    _ = fi48
                }
            }
        default:
            switch "hello" {
            case "hello":
                for fi53 := 0; fi53 < 11; fi53++ {
                    switch "" {
                    case "hello":
                        v54 := GenS5{F0: v5, F1: v2, F2: -5}
                        _ = v54
                    case "a":
                        var v55 float32 = -2.25
                        _ = v55
                    case "foo bar":
                        for fi56 := 0; fi56 < 9; fi56++ {
                            v57 := GenS6{F0: 0, F1: v18, F2: (v1 / ((v1) | 1))}
                            var v58 uint = 561
                            _ = fi56
                            _ = v57
                            _ = v58
                        }
                    default:
                        if v18 {
                            var v59 uint8 = 128
                            _ = v59
                        }
                    }
                    v2 = v2
                    _ = fi53
                }
            case "foo bar":
                var v60 bool = (0 < 32767)
                _ = v60
            }
        }
        fmt.Println(v1, v3)
    } else {
        if (32767 <= uint16(v5)) {
            for fi61 := 0; fi61 < 11; fi61++ {
                for fi62 := 0; fi62 < 16; fi62++ {
                    var v63 int64 = (-(v17))
                    v1 += 473
                    _ = fi62
                    _ = v63
                }
                _ = fi61
            }
            switch -1 {
            case 1:
                switch ("foo bar" + "hello") {
                case "":
                    v5 = (v5 - 1)
                }
            case 127:
                var v64 uint16 = 42
                _ = v64
            case 42:
                var v65 int = -944
                _ = v65
            }
        } else {
            v66 := GenS6{F0: -2147483648, F1: ((v3 / ((2) | 1)) < 127), F2: v1}
            _ = v66
        }
        var v67 uint16 = 2
        var v68 string = ""
        _ = v67
        _ = v68
    }
    var v69 uint64 = v1
    fmt.Println(v18, v3)
    if (1 == v3) {
        if v18 {
            for fi70 := 0; fi70 < 19; fi70++ {
                var v71 uint32 = (v5 - 227)
                _ = fi70
                _ = v71
            }
            fmt.Println(v17)
        } else {
            v72 := v4[v5]
            var v73 int8 = (v2 + v2)
            if v18 {
                v74 := GenS7{F0: v5, F1: v3}
                _ = v74
            } else {
                switch v72 {
                case "":
                    var v75 string = v72
                    _ = v75
                }
                v69 = v1
                if v18 {
                    if (2147483648 == (v5 << ((0) & 63))) {
                        if v18 {
                            var v76 string = (v72 + v72)
                            v77 := GenS8{F0: 1}
                            for fi78 := 0; fi78 < 18; fi78++ {
                                var v79 int16 = 32767
                                _ = fi78
                                _ = v79
                            }
                            _ = v76
                            _ = v77
                        } else {
                            var v80 int64 = int64(v1)
                            _ = v80
                        }
                    } else {
                        switch 2147483647 {
                        case -9223372036854775808:
                            v81 := GenS1{F0: 2}
                            _ = v81
                        }
                        switch -9223372036854775808 {
                        case 32767:
                            switch 32767 {
                            case 42:
                                switch (v72 + v72) {
                                case "\n\t":
                                    var v82 int64 = (v17 & v17)
                                    _ = v82
                                case "hello":
                                    v83 := GenS6{F0: -1, F1: false, F2: v1}
                                    _ = v83
                                default:
                                    for fi84 := 0; fi84 < 20; fi84++ {
                                        fmt.Println(v3)
                                        if v18 {
                                            fmt.Println(fi84, v17, v2)
                                            if len(v6) > 0 {
                                                v6[0] = 2147483648
                                            }
                                        }
                                        _ = fi84
                                    }
                                }
                            case 32767:
                                _ = len(v4)
                            }
                        case 104:
                            if (-32768 > -1) {
                                var v85 int8 = (v2 * (v2 % ((v73) | 1)))
                                _ = v85
                            }
                        case -1:
                            if v18 {
                                for fi86 := 0; fi86 < 16; fi86++ {
                                    for fi87 := 0; fi87 < 19; fi87++ {
                                        var v88 uint64 = 1
                                        switch (v72 + (v72 + v72)) {
                                        case "a":
                                            v5 = v5
                                        case "\n\t":
                                            fmt.Println(v73, v1, v69)
                                        case "":
                                            fmt.Println(fi87, v72, v88)
                                        }
                                        fmt.Println(v1, v88, v17)
                                        v18 = v18
                                        var v89 uint16 = 2
                                        _ = fi87
                                        _ = v88
                                        _ = v89
                                    }
                                    for fi90 := 0; fi90 < 20; fi90++ {
                                        if false {
                                            var v91 float32 = ((((-2.25 * ((-2.25 * 0.0) / 3.14159)) + float32(0.0)) + 1.5) + 0.1)
                                            if true {
                                                v92 := GenS9{F0: -992}
                                                _ = v92
                                            } else {
                                                if (v1 > (v1 % (((v1 >> ((255) & 63))) | 1))) {
                                                    var v93 float32 = v91
                                                    _ = v93
                                                } else {
                                                    if v18 {
                                                        delete(v4, v5)
                                                        v18 = false
                                                        v94 := GenS4{F0: true}
                                                        _ = v94
                                                    }
                                                    switch ("\n\t" + (v72 + v72)) {
                                                    case "a":
                                                        fmt.Println(v69)
                                                    case "":
                                                        if (2 >= (v69 - 2)) {
                                                            v72 = v72
                                                            v95 := GenS3{F0: ("\n\t" + (((v72 + (v72 + ((v72 + v72) + ((((((v72 + v72) + (v72 + ("a" + (v72 + v72)))) + "a") + "hello") + ("hello" + ((v72 + "a") + "\n\t"))) + ("" + (v72 + v72)))))) + "foo bar") + v72)), F1: fi90, F2: 32767}
                                                            _ = v95
                                                        }
                                                    default:
                                                        fmt.Println(v1)
                                                    }
                                                    v1 = 437
                                                }
                                                v18 = ((!(-32768 > 0)) || true)
                                            }
                                            _ = v91
                                        } else {
                                            v96 := GenS4{F0: ((^(v5)) == v5)}
                                            v4[v5] = v72
                                            if (0 < 42) {
                                                if (false || v18) {
                                                    v3 = (v3 - v3)
                                                    v97 := GenS5{F0: v5, F1: (v2 % ((v2) | 1)), F2: v2}
                                                    switch (^(fi90)) {
                                                    case 42:
                                                        switch (fi90 - 0) {
                                                        case 42:
                                                            if (!(21 <= 2147483647)) {
                                                                v98 := GenS2{F0: v73, F1: (0 == (v3 - 2))}
                                                                v99 := GenS3{F0: (("a" + "hello") + "\n\t"), F1: fi86, F2: 42}
                                                                if false {
                                                                    switch "" {
                                                                    case "a":
                                                                        var v100 int = fi86
                                                                        _ = v100
                                                                    case "\n\t":
                                                                        switch fi86 {
                                                                        case -1:
                                                                            v4[(^(v5))] = ""
                                                                        }
                                                                    default:
                                                                        if v18 {
                                                                            v69 -= v69
                                                                            var v101 string = ("a" + v72)
                                                                            _ = v101
                                                                        } else {
                                                                            var v102 float64 = -2.25
                                                                            v103 := GenS10{F0: (^(v2)), F1: v69, F2: v72}
                                                                            fmt.Println(v5)
                                                                            fmt.Println(v73, v3)
                                                                            _ = v102
                                                                            _ = v103
                                                                        }
                                                                    }
                                                                }
                                                                var v104 uint32 = 2147483648
                                                                _ = v104
                                                                _ = v98
                                                                _ = v99
                                                            } else {
                                                                v105 := v4[42]
                                                                _ = v105
                                                            }
                                                        case 2147483647:
                                                            switch ((v72 + "foo bar") + (v72 + "\n\t")) {
                                                            case "hello":
                                                                fmt.Println(v72, v17, v72)
                                                                var v106 uint16 = 42
                                                                _ = v106
                                                            case "\n\t":
                                                                fmt.Println(v18)
                                                            case "foo bar":
                                                                v72 = v72
                                                            default:
                                                                fmt.Println(v18, v17)
                                                            }
                                                        }
                                                    case 127:
                                                        if len(v6) >= 2 {
                                                            v6 = v6[1:len(v6)]
                                                        }
                                                    case 32767:
                                                        fmt.Println(v5, v72, fi90)
                                                    default:
                                                        v17 = -9223372036854775808
                                                    }
                                                    if v18 {
                                                        fmt.Println(v73)
                                                        fmt.Println(v18, v18, v2)
                                                    } else {
                                                        if true {
                                                            var v107 int64 = 159
                                                            var v108 uint64 = 366
                                                            _ = v107
                                                            _ = v108
                                                        } else {
                                                            v109 := GenS11{F0: (-(v17)), F1: -513, F2: 151}
                                                            _ = v109
                                                        }
                                                        v110 := GenS10{F0: v2, F1: (v69 & v1), F2: "hello"}
                                                        _ = v110
                                                    }
                                                    if (65535 >= uint(v5)) {
                                                        v72 = "hello"
                                                        fmt.Println(v18, v1, v17)
                                                    } else {
                                                        switch fi86 {
                                                        case -598:
                                                            var v111 int16 = int16(v17)
                                                            _ = v111
                                                        }
                                                    }
                                                    _ = v97
                                                }
                                                _ = len(v6)
                                                var v112 int8 = (v73 << ((431) & 63))
                                                _ = v112
                                            } else {
                                                v113 := GenS2{F0: v2, F1: (!v18)}
                                                v3 -= (v3 & v3)
                                                _ = v113
                                            }
                                            _ = v96
                                        }
                                        var v114 uint8 = v3
                                        v115 := GenS8{F0: (-(fi86))}
                                        _ = fi90
                                        _ = v114
                                        _ = v115
                                    }
                                    _ = fi86
                                }
                                if ((15 >= (v3 << ((4294967295) & 63))) && (!(v5 != (v5 >> ((651) & 63))))) {
                                    v116 := GenS12{F0: v73, F1: (v2 ^ (v73 | (^(v2)))), F2: 458}
                                    if (0 == 32767) {
                                        for fi117 := 0; fi117 < 12; fi117++ {
                                            var v118 bool = (!(!(true || v18)))
                                            _ = fi117
                                            _ = v118
                                        }
                                        var v119 uint8 = 1
                                        _ = v119
                                    } else {
                                        if (650 >= (v1 | (^(v69)))) {
                                            fmt.Println(v3, v17)
                                            var v120 int8 = -1
                                            _ = v120
                                        }
                                        v121 := GenS5{F0: (^(v5)), F1: v73, F2: 64}
                                        v122 := GenS13{F0: 127}
                                        _ = v121
                                        _ = v122
                                    }
                                    var v123 int = 42
                                    _ = v116
                                    _ = v123
                                }
                                v17 = (v17 | (v17 / ((0) | 1)))
                            } else {
                                if true {
                                    if v18 {
                                        if v18 {
                                            if len(v6) > 0 {
                                                v6[0] = v5
                                            }
                                            if len(v6) >= 2 {
                                                v6 = v6[1:len(v6)]
                                            }
                                            v69 -= 789
                                        }
                                        v124 := GenS9{F0: 0}
                                        _ = v124
                                    }
                                } else {
                                    v125 := GenS1{F0: 255}
                                    if v18 {
                                        var v126 int64 = (-(v17))
                                        delete(v4, v5)
                                        _ = v126
                                    }
                                    v127 := GenS10{F0: (v73 / (((v73 % (((v73 << ((0) & 63))) | 1))) | 1)), F1: v69, F2: v72}
                                    _ = v125
                                    _ = v127
                                }
                            }
                        default:
                            for fi128 := 0; fi128 < 10; fi128++ {
                                var v129 string = (v72 + v72)
                                var v130 int64 = (v17 & -1)
                                fmt.Println(v130, v17, v72)
                                _ = fi128
                                _ = v129
                                _ = v130
                            }
                        }
                    }
                    var v131 bool = false
                    v132 := GenS13{F0: v2}
                    _ = v131
                    _ = v132
                }
            }
            _ = v72
            _ = v73
        }
    } else {
        switch 2147483647 {
        case 892:
            var v133 float64 = (-2.25 * (0.0 * (float64(0.1) - (0.1 - 0.1))))
            _ = v133
        case 2147483647:
            v2 -= 100
        }
        v134 := GenS1{F0: uint16(v3)}
        _ = v134
    }
    var v135 int8 = v2
    if (int64(v5) > v17) {
        v136 := GenS14{F0: v3, F1: 403}
        fmt.Println(v17)
        var v137 bool = (int(v3) == 301)
        _ = v136
        _ = v137
    } else {
        var v138 uint64 = (v69 | v1)
        var v139 uint = 42
        _ = v138
        _ = v139
    }
    fmt.Println("v1:", v1)
    fmt.Println("v135:", v135)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v69:", v69)
    _ = v1
    _ = v135
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v69
}
