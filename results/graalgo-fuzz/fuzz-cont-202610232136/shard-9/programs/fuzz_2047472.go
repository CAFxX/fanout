package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32 }
type GenS2 struct { F0 uint16; F1 int8 }
type GenS3 struct { F0 uint64; F1 uint; F2 bool }
type GenS4 struct { F0 int16; F1 uint }
type GenS5 struct { F0 int8; F1 int; F2 uint64 }
type GenS6 struct { F0 string; F1 int16; F2 uint64 }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 int16; F1 uint32; F2 int64 }
type GenS9 struct { F0 uint32 }
type GenS10 struct { F0 int16; F1 uint64 }
type GenS11 struct { F0 int16; F1 int64; F2 uint }
type GenS12 struct { F0 string; F1 int16 }
type GenS13 struct { F0 int32; F1 string; F2 int32 }
type GenS14 struct { F0 int }
type GenS15 struct { F0 bool }
type GenS16 struct { F0 int8; F1 uint8 }
type GenS17 struct { F0 bool; F1 bool }
type GenS18 struct { F0 int64; F1 uint16; F2 int8 }

func main() {
    var v1 int32 = 42
    var v2 int16 = 1000
    var v3 float32 = -2.25
    v4 := make(map[uint32]int)
    var v5 bool = (!(65535 != 623))
    var v6 uint8 = 42
    if false {
        for fi7 := 0; fi7 < 17; fi7++ {
            v8 := GenS1{F0: 2, F1: (v1 >> ((4294967295) & 63))}
            var v9 uint64 = 336
            _ = fi7
            _ = v8
            _ = v9
        }
        v6 += v6
    }
    for fi10 := 0; fi10 < 8; fi10++ {
        if (81 == 0) {
            var v11 int16 = (-(v2))
            _ = v11
        } else {
            switch 1 {
            case -1:
                for fi12 := 0; fi12 < 16; fi12++ {
                    v13 := GenS2{F0: 42, F1: 82}
                    v14 := GenS1{F0: 18446744073709551615, F1: (v1 & (v1 >> ((493) & 63)))}
                    switch ("a" + "a") {
                    case "hello":
                        var v15 float64 = 3.14159
                        _ = v15
                    default:
                        if ((true && false) && ((v5 && ((0 < 1) || v5)) && v5)) {
                            var v16 uint32 = 819
                            v17 := v4[65535]
                            var v18 uint16 = 1
                            _ = v16
                            _ = v17
                            _ = v18
                        } else {
                            v19 := GenS2{F0: 2, F1: -1}
                            _ = v19
                        }
                    }
                    fmt.Println(v2)
                    _ = fi12
                    _ = v13
                    _ = v14
                }
            case -423:
                fmt.Println(v2)
            case 1:
                var v20 uint16 = 505
                _ = v20
            }
        }
        _ = fi10
    }
    v21 := GenS3{F0: 0, F1: 609, F2: (!false)}
    if true {
        v22 := GenS4{F0: (v2 << ((18446744073709551615) & 63)), F1: 4294967295}
        _ = v22
    } else {
        if (4294967295 < 386) {
            v23 := GenS2{F0: 42, F1: 127}
            _ = v23
        }
        if false {
            switch "foo bar" {
            case "foo bar":
                if (v5 && (v5 || false)) {
                    v24 := GenS5{F0: 42, F1: 1, F2: 1}
                    _ = v24
                } else {
                    if (!v5) {
                        var v25 uint = 564
                        _ = v25
                    }
                    if (((v2 * -32768) <= (v2 << ((750) & 63))) && (9223372036854775807 != -956)) {
                        switch -945 {
                        case 2147483647:
                            switch "a" {
                            case "\n\t":
                                var v26 string = "\n\t"
                                _ = v26
                            case "a":
                                _ = len(v4)
                            case "hello":
                                v27 := GenS2{F0: 1, F1: -1}
                                _ = v27
                            default:
                                v28 := v4[799]
                                _ = v28
                            }
                        case -885:
                            if (0 == 574) {
                                var v29 uint32 = 2147483648
                                v30 := []int{}
                                if (-308 == -891) {
                                    v2 += v2
                                    for fi31 := 0; fi31 < 1; fi31++ {
                                        switch 533 {
                                        case 32767:
                                            v3 = v3
                                        }
                                        _ = fi31
                                    }
                                } else {
                                    _ = len(v4)
                                }
                                _ = v29
                                _ = v30
                            }
                        default:
                            v1 -= 0
                        }
                        if v5 {
                            v32 := []int{0, 4294967295, -9223372036854775808, -307}
                            if ((v5 || (4294967295 >= 42)) || (v5 && true)) {
                                var v33 uint16 = 0
                                fmt.Println(v2, v6, v33)
                                v2 = 42
                                _ = v33
                            }
                            _ = v32
                        } else {
                            var v34 float32 = v3
                            _ = v34
                        }
                        for fi35 := 0; fi35 < 4; fi35++ {
                            v36 := v4[53]
                            switch (v36 << ((uint(v6)) & 63)) {
                            case 2147483647:
                                switch (-(fi35)) {
                                case 127:
                                    v3 -= v3
                                case 4294967295:
                                    v4[4294967295] = -9223372036854775808
                                case 128:
                                    var v37 int32 = 420
                                    _ = v37
                                }
                            case -2147483648:
                                v38 := GenS2{F0: 2, F1: 127}
                                _ = v38
                            case 4294967295:
                                if ((((v1 % ((0) | 1)) >= (v1 & v1)) || false) && ((false && (!(128 != 42))) || v5)) {
                                    switch "\n\t" {
                                    case "hello":
                                        switch (fi35 - v36) {
                                        case 127:
                                            v6 = (v6 - v6)
                                        case 368:
                                            if v5 {
                                                for fi39 := 0; fi39 < 16; fi39++ {
                                                    var v40 uint8 = (v6 >> ((65535) & 63))
                                                    fmt.Println(fi35, v2, fi39)
                                                    fmt.Println(v3, v40)
                                                    _ = fi39
                                                    _ = v40
                                                }
                                            }
                                        case -9223372036854775808:
                                            var v41 uint = 272
                                            _ = v41
                                        }
                                    case "\n\t":
                                        var v42 uint16 = 756
                                        _ = v42
                                    }
                                    delete(v4, 2147483647)
                                } else {
                                    var v43 int64 = -9223372036854775808
                                    _ = v43
                                }
                            default:
                                var v44 uint8 = v6
                                _ = v44
                            }
                            v5 = (v3 > (v3 * (3.14159 / float32(float64(((v3 / v3) - v3))))))
                            _ = fi35
                            _ = v36
                        }
                    } else {
                        if (1 > 255) {
                            if v5 {
                                v3 = v3
                                var v45 float64 = float64(-2.25)
                                for fi46 := 0; fi46 < 10; fi46++ {
                                    v47 := []int{(fi46 / ((9223372036854775807) | 1)), -624}
                                    _ = fi46
                                    _ = v47
                                }
                                _ = v45
                            }
                            for fi48 := 0; fi48 < 3; fi48++ {
                                if ("foo bar" == ("" + "\n\t")) {
                                    v49 := GenS6{F0: "hello", F1: v2, F2: 42}
                                    fmt.Println(v1, v2, v3)
                                    for fi50 := 0; fi50 < 11; fi50++ {
                                        v6 = v6
                                        _ = fi50
                                    }
                                    _ = v49
                                } else {
                                    v2 += 42
                                }
                                _ = fi48
                            }
                            if (((float64(((v3 / v3) * 3.14159)) + -2.25) > (float64((v3 - 3.14159)) / 1.0)) || (65535 < 255)) {
                                fmt.Println(v3)
                            } else {
                                fmt.Println(v2, v6)
                                var v51 uint64 = 4294967295
                                var v52 int8 = -128
                                _ = v51
                                _ = v52
                            }
                        } else {
                            v53 := GenS7{F0: v2}
                            if v5 {
                                var v54 uint32 = 65535
                                if v5 {
                                    v55 := GenS8{F0: 1000, F1: (v54 ^ (^(v54))), F2: 9223372036854775807}
                                    if false {
                                        for fi56 := 0; fi56 < 15; fi56++ {
                                            v57 := GenS5{F0: -1, F1: fi56, F2: 4294967295}
                                            v58 := GenS8{F0: (v2 / ((42) | 1)), F1: v54, F2: int64(v54)}
                                            v2 += -32768
                                            _ = fi56
                                            _ = v57
                                            _ = v58
                                        }
                                        fmt.Println(v2, v1)
                                    } else {
                                        if (!false) {
                                            for fi59 := 0; fi59 < 17; fi59++ {
                                                for fi60 := 0; fi60 < 12; fi60++ {
                                                    if ((true || false) || false) {
                                                        if ("a" != ("hello" + ("foo bar" + "foo bar"))) {
                                                            v61 := GenS6{F0: (("foo bar" + "a") + ""), F1: (v2 * v2), F2: 2}
                                                            var v62 uint16 = 32767
                                                            _ = v61
                                                            _ = v62
                                                        } else {
                                                            var v63 int8 = 0
                                                            var v64 uint = 820
                                                            if true {
                                                                var v65 uint = v64
                                                                _ = v65
                                                            }
                                                            _ = v63
                                                            _ = v64
                                                        }
                                                        fmt.Println(v5)
                                                    } else {
                                                        var v66 uint8 = 1
                                                        _ = v66
                                                    }
                                                    v54 = v54
                                                    fmt.Println(v6, v6, v2)
                                                    _ = fi60
                                                }
                                                switch -961 {
                                                case 9223372036854775807:
                                                    v6 -= (v6 | (v6 + v6))
                                                case -800:
                                                    v67 := GenS9{F0: v54}
                                                    _ = v67
                                                case 32767:
                                                    v68 := GenS3{F0: 4294967295, F1: 987, F2: true}
                                                    _ = v68
                                                }
                                                _ = fi59
                                            }
                                        } else {
                                            var v69 float64 = (((3.14159 + float64(v3)) * -2.25) + ((3.14159 - 1.5) + 0.1))
                                            _ = v69
                                        }
                                        var v70 int = -1
                                        _ = v70
                                    }
                                    _ = v55
                                } else {
                                    if true {
                                        v71 := []int{int(v54), -457}
                                        _ = v71
                                    }
                                    switch "hello" {
                                    case "":
                                        v2 = v2
                                    case "a":
                                        v72 := []int{}
                                        _ = v72
                                    default:
                                        var v73 float32 = ((v3 * (float32((1e10 * ((0.0 / 0.1) * -2.25))) * float32(float64(0.1)))) + ((float32(0.0) / 1e10) / (0.0 + v3)))
                                        _ = v73
                                    }
                                }
                                _ = v54
                            }
                            var v74 uint16 = 65535
                            _ = v53
                            _ = v74
                        }
                        switch "foo bar" {
                        case "\n\t":
                            var v75 int16 = 1000
                            _ = v75
                        default:
                            if (v3 >= (float32((((0.1 + (1e10 / float64(v3))) + float64((v3 + ((v3 - v3) / v3)))) - 0.1)) / v3)) {
                                for fi76 := 0; fi76 < 11; fi76++ {
                                    v3 = 3.14159
                                    var v77 uint64 = 18446744073709551615
                                    _ = fi76
                                    _ = v77
                                }
                                var v78 string = ("\n\t" + "")
                                var v79 bool = (65535 != 2147483647)
                                _ = v78
                                _ = v79
                            }
                        }
                    }
                    v80 := GenS10{F0: 0, F1: 18446744073709551615}
                    _ = v80
                }
            }
            v4[2147483648] = 0
        }
    }
    v81 := GenS5{F0: -12, F1: -1, F2: 0}
    switch "hello" {
    case "":
        var v82 uint8 = v6
        _ = v82
    case "hello":
        var v83 string = "foo bar"
        _ = v83
    case "foo bar":
        if v5 {
            switch "hello" {
            case "\n\t":
                v6 = (^(v6))
            case "hello":
                if v5 {
                    var v84 float32 = 1.5
                    v85 := GenS2{F0: 368, F1: 100}
                    v2 -= v2
                    _ = v84
                    _ = v85
                } else {
                    for fi86 := 0; fi86 < 3; fi86++ {
                        for fi87 := 0; fi87 < 11; fi87++ {
                            if (v5 || (255 > 396)) {
                                v88 := GenS1{F0: 0, F1: 1}
                                var v89 uint32 = uint32(v6)
                                var v90 uint = 0
                                var v91 int32 = (-(v1))
                                v92 := GenS9{F0: v89}
                                v6 = 128
                                _ = v88
                                _ = v89
                                _ = v90
                                _ = v91
                                _ = v92
                            } else {
                                v93 := GenS11{F0: (v2 | 42), F1: 310, F2: 18446744073709551615}
                                v3 -= v3
                                _ = v93
                            }
                            _ = fi87
                        }
                        v94 := GenS12{F0: ("a" + "foo bar"), F1: (v2 + (v2 % (((^(v2))) | 1)))}
                        var v95 float64 = ((0.0 - 1e10) + 0.0)
                        _ = fi86
                        _ = v94
                        _ = v95
                    }
                }
            case "foo bar":
                fmt.Println(v5)
            default:
                v96 := GenS13{F0: v1, F1: ("" + "hello"), F2: 2147483647}
                _ = v96
            }
            v6 -= (v6 % ((128) | 1))
            var v97 string = ("" + (("hello" + ("a" + ("a" + ("" + "foo bar")))) + "hello"))
            _ = v97
        }
    default:
        v98 := GenS14{F0: 128}
        _ = v98
    }
    var v99 uint16 = 667
    v100 := GenS1{F0: 1, F1: -467}
    var v101 int16 = 609
    v5 = (!v5)
    v102 := v4[675]
    for fi103 := 0; fi103 < 20; fi103++ {
        switch (-(v102)) {
        case 32768:
            v6 += v6
        case 9223372036854775807:
            for fi104 := 0; fi104 < 8; fi104++ {
                if false {
                    fmt.Println(v2, fi104)
                    if ("" < "\n\t") {
                        var v105 string = (("\n\t" + "hello") + (("foo bar" + "a") + "hello"))
                        var v106 int32 = -2147483648
                        _ = v105
                        _ = v106
                    } else {
                        if ((0.0 * 3.14159) < float32(-2.25)) {
                            switch "hello" {
                            case "foo bar":
                                _ = len(v4)
                            case "\n\t":
                                var v107 uint16 = (v99 % ((255) | 1))
                                _ = v107
                            }
                            var v108 uint32 = 4294967295
                            var v109 float32 = v3
                            _ = v108
                            _ = v109
                        }
                        var v110 bool = v5
                        switch "" {
                        case "a":
                            v111 := v4[42]
                            _ = v111
                        case "":
                            v112 := GenS8{F0: v101, F1: 1, F2: 1}
                            _ = v112
                        case "\n\t":
                            if false {
                                v101 += int16(v99)
                            } else {
                                v113 := []int{v102, -472, 127, (v102 - fi104)}
                                _ = v113
                            }
                        }
                        _ = v110
                    }
                } else {
                    v114 := GenS7{F0: v101}
                    var v115 int = (v102 | (fi103 ^ (fi103 & fi104)))
                    if (-120 > -1) {
                        switch "\n\t" {
                        case "a":
                            v3 = 1.5
                        case "\n\t":
                            v115 = -248
                        case "":
                            if (false && v5) {
                                v116 := []int{}
                                v101 -= 0
                                v117 := GenS7{F0: 1}
                                delete(v4, 2)
                                var v118 float32 = float32(float64(0.1))
                                _ = v116
                                _ = v117
                                _ = v118
                            } else {
                                switch v102 {
                                case 127:
                                    v119 := GenS15{F0: (!(!(!(!v5))))}
                                    _ = v119
                                case 32767:
                                    v120 := GenS14{F0: -2147483648}
                                    _ = v120
                                case 0:
                                    var v121 int8 = 0
                                    _ = v121
                                }
                                _ = len(v4)
                            }
                        }
                        if (0 < v99) {
                            v122 := []int{v115, (v115 * fi103)}
                            _ = v122
                        }
                        var v123 uint8 = v6
                        _ = v123
                    } else {
                        if v5 {
                            v99 -= v99
                        } else {
                            fmt.Println(v1, v99, fi103)
                            v124 := GenS16{F0: 0, F1: 151}
                            fmt.Println(v101)
                            _ = v124
                        }
                    }
                    v125 := GenS17{F0: v5, F1: false}
                    v102 -= int(v99)
                    _ = v114
                    _ = v115
                    _ = v125
                }
                _ = fi104
            }
        case 32767:
            if v5 {
                var v126 uint8 = 127
                if (!(1 != 0)) {
                    v127 := []int{0, (v102 | (^(fi103)))}
                    fmt.Println(v101)
                    switch v102 {
                    case -9223372036854775808:
                        switch "\n\t" {
                        case "hello":
                            switch (fi103 ^ 32767) {
                            case 955:
                                v128 := GenS8{F0: (v101 + v101), F1: 631, F2: 9223372036854775807}
                                _ = v128
                            }
                        default:
                            for fi129 := 0; fi129 < 6; fi129++ {
                                var v130 int32 = v1
                                _ = fi129
                                _ = v130
                            }
                        }
                    }
                    _ = v127
                } else {
                    switch fi103 {
                    case -2147483648:
                        fmt.Println(v5)
                    default:
                        v2 = v101
                    }
                    if (3.14159 >= 0.0) {
                        v1 = -295
                    }
                    v99 = v99
                }
                var v131 uint8 = v6
                _ = v126
                _ = v131
            } else {
                v132 := GenS12{F0: "a", F1: (v2 + (v2 / ((1) | 1)))}
                _ = v132
            }
        }
        v5 = (v2 < v2)
        _ = fi103
    }
    var v133 float32 = float32(0.1)
    for fi134 := 0; fi134 < 18; fi134++ {
        var v135 int = v102
        _ = fi134
        _ = v135
    }
    var v136 int32 = -1
    v137 := GenS18{F0: 9223372036854775807, F1: v99, F2: -128}
    fmt.Println("v1:", v1)
    fmt.Println("v101:", v101)
    fmt.Println("v102:", v102)
    fmt.Println("v133:", v133)
    fmt.Println("v136:", v136)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v99:", v99)
    _ = v1
    _ = v100
    _ = v101
    _ = v102
    _ = v133
    _ = v136
    _ = v137
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v81
    _ = v99
}
