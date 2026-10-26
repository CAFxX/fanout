package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint; F1 uint16; F2 uint16 }
type GenS3 struct { F0 uint; F1 bool; F2 uint16 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint16; F1 uint16 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 int16; F1 uint8 }
type GenS8 struct { F0 string; F1 int8 }
type GenS9 struct { F0 uint64; F1 int16; F2 int16 }
type GenS10 struct { F0 int8; F1 uint8; F2 uint64 }
type GenS11 struct { F0 int64 }
type GenS12 struct { F0 int64; F1 int64 }
type GenS13 struct { F0 int64; F1 uint64; F2 uint64 }
type GenS14 struct { F0 uint; F1 int32 }
type GenS15 struct { F0 uint64; F1 uint8; F2 int8 }
type GenS16 struct { F0 int32 }
type GenS17 struct { F0 bool; F1 int64; F2 uint32 }
type GenS18 struct { F0 int }
type GenS19 struct { F0 uint8; F1 int32 }
func genFn0(p0 uint32, p1 bool, p2 uint16) uint16 {
    if p1 {
        if (0 > 1) {
            var v4 uint8 = 39
            v5 := []int{32768, 127}
            _ = v4
            _ = v5
        }
        p0 -= p0
        for fi6 := 0; fi6 < 14; fi6++ {
            v7 := GenS1{F0: 255}
            _ = fi6
            _ = v7
        }
    }
    for fi8 := 0; fi8 < 12; fi8++ {
        if (p2 != 42) {
            var v9 float64 = 1e10
            var v10 float32 = 1.5
            var v11 uint = 42
            _ = v10
            _ = v11
            _ = v9
        } else {
            if true {
                p0 = (p0 + 1)
            }
        }
        _ = fi8
    }
    v12 := []int{128}
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    return p2
}
func genFn1(p0 int64, p1 uint, p2 uint16) int32 {
    p2 -= (p2 - (p2 / ((65535) | 1)))
    v13 := []int{}
    var v14 uint8 = 42
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v14
    return 0
}

func main() {
    var v1 uint16 = 0
    var v2 int64 = -532
    v3 := make(map[uint32]bool)
    v1 = v1
    v1 = 0
    for fi15 := 0; fi15 < 19; fi15++ {
        fmt.Println(v2, v1, fi15)
        _ = fi15
    }
    v1 += 773
    if (375 <= 42) {
        switch 32767 {
        case 9223372036854775807:
            v16 := GenS2{F0: 255, F1: genFn0(1, (!(true || false)), (v1 % (((v1 / ((65535) | 1))) | 1))), F2: v1}
            _ = v16
        case -763:
            for fi17 := 0; fi17 < 7; fi17++ {
                v18 := GenS1{F0: 703}
                _ = fi17
                _ = v18
            }
        default:
            v19 := v3[2147483648]
            _ = v19
        }
        v1 = genFn0(513, true, (v1 % (((v1 / (((v1 >> ((18446744073709551615) & 63))) | 1))) | 1)))
    } else {
        v20 := GenS2{F0: 4294967295, F1: (^(v1)), F2: genFn0(101, true, v1)}
        if true {
            v21 := GenS3{F0: 7, F1: (128 != 9223372036854775807), F2: genFn0(2, false, (v1 % ((32767) | 1)))}
            var v22 int64 = -1
            var v23 int32 = 847
            _ = v21
            _ = v22
            _ = v23
        } else {
            for fi24 := 0; fi24 < 17; fi24++ {
                var v25 int8 = -27
                v2 -= -96
                v26 := []int{fi24, (fi24 & (fi24 - 9223372036854775807)), fi24}
                _ = fi24
                _ = v25
                _ = v26
            }
            fmt.Println(v1)
            switch 42 {
            case 1:
                v1 = (v1 + v1)
            case 9223372036854775807:
                var v27 int64 = -1
                _ = v27
            case -2147483648:
                for fi28 := 0; fi28 < 15; fi28++ {
                    v3[439] = (!(4294967295 > 4294967295))
                    for fi29 := 0; fi29 < 7; fi29++ {
                        v30 := GenS4{F0: (212 <= 127)}
                        var v31 string = "foo bar"
                        _ = fi29
                        _ = v30
                        _ = v31
                    }
                    _ = fi28
                }
            default:
                if (0 < genFn1(v2, 0, (v1 & (v1 * (v1 / (((v1 / ((genFn0(646, (1 != 42), v1)) | 1))) | 1)))))) {
                    v3[2147483648] = (((0.1 + -2.25) <= float32(-2.25)) && true)
                    v1 = 2
                }
            }
        }
        v32 := GenS3{F0: 4294967295, F1: true, F2: (^(v1))}
        _ = v20
        _ = v32
    }
    v1 = 65535
    if true {
        delete(v3, 2147483647)
    }
    v33 := GenS3{F0: 255, F1: true, F2: v1}
    var v34 int = -2147483648
    var v35 uint = 4294967295
    if ((9223372036854775807 == v2) || true) {
        var v36 float32 = (((float32(float64(0.0)) * float32(1e10)) - 0.1) + 3.14159)
        if (!((-1 > 1) || false)) {
            if true {
                if (287 <= 364) {
                    v1 -= (v1 + (v1 ^ genFn0(uint32(v35), true, (v1 ^ v1))))
                    v2 = (v2 | (v2 >> (((v35 << ((1) & 63))) & 63)))
                }
                v36 = v36
            }
        }
        _ = v36
    }
    if (!(!true)) {
        var v37 int64 = v2
        _ = v37
    }
    switch "a" {
    case "":
        var v38 int = (^(v34))
        _ = v38
    case "\n\t":
        for fi39 := 0; fi39 < 12; fi39++ {
            for fi40 := 0; fi40 < 4; fi40++ {
                if (3.14159 < (0.1 + (-2.25 * 3.14159))) {
                    var v41 string = "a"
                    v42 := GenS3{F0: v35, F1: false, F2: genFn0(65535, (42 != -128), v1)}
                    _ = v41
                    _ = v42
                } else {
                    v43 := GenS3{F0: (^(v35)), F1: false, F2: 551}
                    _ = len(v3)
                    var v44 uint32 = 173
                    _ = v43
                    _ = v44
                }
                v45 := v3[1]
                _ = fi40
                _ = v45
            }
            _ = fi39
        }
    default:
        var v46 bool = (0 < 2147483648)
        _ = v46
    }
    v47 := GenS3{F0: 65535, F1: true, F2: genFn0(0, false, (v1 + genFn0(720, (genFn1((v2 ^ (v2 + 930)), v35, v1) > genFn1(v2, v35, v1)), v1)))}
    fmt.Println(v1, v35)
    v1 = v1
    var v48 uint16 = v1
    if true {
        switch (v34 | v34) {
        case 2147483647:
            if (!false) {
                if false {
                    _ = len(v3)
                    if true {
                        v34 -= -921
                    }
                } else {
                    var v49 int16 = 1000
                    v48 -= (v48 >> (((v35 ^ v35)) & 63))
                    _ = v49
                }
            } else {
                switch (v34 & 4294967295) {
                case -2147483648:
                    for fi50 := 0; fi50 < 11; fi50++ {
                        v1 += v48
                        _ = fi50
                    }
                default:
                    v51 := GenS3{F0: (v35 % ((18446744073709551615) | 1)), F1: (((-128 == 1) || true) || ((((0.1 * 3.14159) / 1e10) * 0.1) > 1e10)), F2: v1}
                    _ = v51
                }
            }
        case 127:
            for fi52 := 0; fi52 < 7; fi52++ {
                v34 -= 0
                v3[2147483647] = (-857 <= -1)
                _ = fi52
            }
        case -845:
            fmt.Println(v48, v2, v1)
        }
        if false {
            if true {
                if false {
                    fmt.Println(v34)
                    v35 = v35
                } else {
                    var v53 uint = v35
                    v48 = genFn0(218, (0 == 18446744073709551615), (v48 << ((1000) & 63)))
                    _ = v53
                }
                if true {
                    if true {
                        if false {
                            if (!(-40 >= 100)) {
                                var v54 int8 = 0
                                v55 := GenS3{F0: v35, F1: false, F2: (v1 / ((1) | 1))}
                                _ = v54
                                _ = v55
                            } else {
                                for fi56 := 0; fi56 < 20; fi56++ {
                                    for fi57 := 0; fi57 < 4; fi57++ {
                                        v58 := GenS4{F0: (4294967295 > 2)}
                                        v59 := []int{v34, fi56, v34}
                                        v2 -= v2
                                        _ = fi57
                                        _ = v58
                                        _ = v59
                                    }
                                    _ = fi56
                                }
                                v60 := v3[2147483648]
                                _ = v60
                            }
                            v61 := GenS5{F0: (v1 | v48), F1: genFn0(uint32(v35), (964 <= 42), (^(v48)))}
                            fmt.Println(v1, v35, v35)
                            _ = v61
                        } else {
                            var v62 uint32 = 65535
                            if false {
                                if true {
                                    var v63 int = 0
                                    _ = v63
                                } else {
                                    if (!(18446744073709551615 < 18446744073709551615)) {
                                        var v64 int8 = 27
                                        fmt.Println(v64)
                                        v35 = v35
                                        _ = v64
                                    } else {
                                        delete(v3, 65535)
                                        if ((v62 << ((v35) & 63)) > v62) {
                                            for fi65 := 0; fi65 < 17; fi65++ {
                                                fmt.Println(v34, v35, v35)
                                                _ = fi65
                                            }
                                        } else {
                                            if (14 <= 1) {
                                                _ = len(v3)
                                            }
                                            switch "foo bar" {
                                            case "\n\t":
                                                v66 := GenS6{F0: 126}
                                                _ = v66
                                            case "":
                                                if (!false) {
                                                    for fi67 := 0; fi67 < 17; fi67++ {
                                                        switch "hello" {
                                                        case "":
                                                            for fi68 := 0; fi68 < 4; fi68++ {
                                                                v69 := GenS7{F0: 683, F1: 81}
                                                                _ = fi68
                                                                _ = v69
                                                            }
                                                        case "foo bar":
                                                            if (false || (false || false)) {
                                                                var v70 float64 = 1e10
                                                                _ = v70
                                                            } else {
                                                                fmt.Println(v34, v1)
                                                                v34 = (fi67 % ((fi67) | 1))
                                                                switch "\n\t" {
                                                                case "foo bar":
                                                                    v71 := GenS8{F0: (("" + "\n\t") + "a"), F1: -4}
                                                                    _ = v71
                                                                case "\n\t":
                                                                    var v72 float32 = float32(float64(1e10))
                                                                    _ = v72
                                                                default:
                                                                    switch ("\n\t" + ("" + "")) {
                                                                    case "foo bar":
                                                                        var v73 bool = ((v34 >> (((^(v35))) & 63)) > (v34 >> (((^(v35))) & 63)))
                                                                        _ = v73
                                                                    case "\n\t":
                                                                        for fi74 := 0; fi74 < 2; fi74++ {
                                                                            v48 += genFn0(uint32(v48), (556 == v35), v1)
                                                                            if false {
                                                                                fmt.Println(v48, fi67)
                                                                                if (true && (128 != 127)) {
                                                                                    var v75 float64 = float64(3.14159)
                                                                                    v76 := []int{fi67, 0}
                                                                                    _ = v75
                                                                                    _ = v76
                                                                                }
                                                                                if true {
                                                                                    v1 += v48
                                                                                    switch (fi74 - (fi67 * -1)) {
                                                                                    case 9223372036854775807:
                                                                                        if (1.5 >= (-2.25 - 1e10)) {
                                                                                            if (((3.14159 - 0.0) < float32(-2.25)) || false) {
                                                                                                if (!false) {
                                                                                                    if false {
                                                                                                        var v77 uint = (v35 << ((v35) & 63))
                                                                                                        var v78 int16 = 5
                                                                                                        _ = v77
                                                                                                        _ = v78
                                                                                                    } else {
                                                                                                        switch "" {
                                                                                                        case "":
                                                                                                            v1 -= v48
                                                                                                            switch "a" {
                                                                                                            case "a":
                                                                                                                switch v34 {
                                                                                                                case 42:
                                                                                                                    v79 := GenS2{F0: 65535, F1: v1, F2: 1}
                                                                                                                    _ = v79
                                                                                                                case 789:
                                                                                                                    v48 += (v48 ^ (v1 << ((uint(v62)) & 63)))
                                                                                                                case 32768:
                                                                                                                    var v80 uint32 = v62
                                                                                                                    _ = v80
                                                                                                                }
                                                                                                            case "foo bar":
                                                                                                                if true {
                                                                                                                    v81 := GenS9{F0: 18446744073709551615, F1: 105, F2: -1}
                                                                                                                    v1 = genFn0(v62, false, v48)
                                                                                                                    v2 = -9223372036854775808
                                                                                                                    _ = v81
                                                                                                                } else {
                                                                                                                    v48 = genFn0((v62 - v62), true, (v48 << ((808) & 63)))
                                                                                                                }
                                                                                                            }
                                                                                                        case "hello":
                                                                                                            v34 += (fi74 | (-(fi67)))
                                                                                                        }
                                                                                                        v82 := []int{}
                                                                                                        var v83 int32 = 1
                                                                                                        var v84 int32 = 721
                                                                                                        var v85 int = (fi74 | (fi74 % ((9223372036854775807) | 1)))
                                                                                                        _ = v82
                                                                                                        _ = v83
                                                                                                        _ = v84
                                                                                                        _ = v85
                                                                                                    }
                                                                                                }
                                                                                                v48 += 780
                                                                                            } else {
                                                                                                v3[(v62 % ((4294967295) | 1))] = false
                                                                                            }
                                                                                            _ = len(v3)
                                                                                        }
                                                                                    case 346:
                                                                                        v35 += v35
                                                                                    case 0:
                                                                                        switch fi67 {
                                                                                        case 127:
                                                                                            v86 := GenS10{F0: 1, F1: 127, F2: 1}
                                                                                            _ = v86
                                                                                        case 4294967295:
                                                                                            v1 += (v1 & v48)
                                                                                        default:
                                                                                            v2 += (v2 + v2)
                                                                                        }
                                                                                    }
                                                                                }
                                                                            }
                                                                            _ = fi74
                                                                        }
                                                                    case "hello":
                                                                        v87 := GenS11{F0: (v2 | 275)}
                                                                        _ = v87
                                                                    }
                                                                }
                                                            }
                                                        }
                                                        _ = fi67
                                                    }
                                                }
                                            default:
                                                var v88 int64 = v2
                                                _ = v88
                                            }
                                            var v89 float64 = float64(-2.25)
                                            _ = v89
                                        }
                                    }
                                }
                                delete(v3, v62)
                                var v90 int16 = 0
                                _ = v90
                            }
                            v3[(v62 | v62)] = false
                            _ = v62
                        }
                        fmt.Println(v48, v48, v35)
                        _ = len(v3)
                    } else {
                        v35 = v35
                        for fi91 := 0; fi91 < 12; fi91++ {
                            switch 32768 {
                            case 332:
                                for fi92 := 0; fi92 < 7; fi92++ {
                                    v93 := GenS12{F0: v2, F1: v2}
                                    v94 := GenS8{F0: "foo bar", F1: 56}
                                    v34 = fi91
                                    _ = fi92
                                    _ = v93
                                    _ = v94
                                }
                            default:
                                switch (v34 + v34) {
                                case -395:
                                    switch ("" + ("a" + "")) {
                                    case "hello":
                                        if false {
                                            var v95 float64 = (1e10 + (((3.14159 + 1.5) - 0.0) + 0.1))
                                            var v96 uint16 = (v48 * (v1 << ((v35) & 63)))
                                            switch "foo bar" {
                                            case "a":
                                                v34 += (-(fi91))
                                            }
                                            _ = v95
                                            _ = v96
                                        }
                                    case "foo bar":
                                        for fi97 := 0; fi97 < 16; fi97++ {
                                            if (!false) {
                                                v1 = 0
                                                if false {
                                                    if (445 <= 32767) {
                                                        if (false && (false || true)) {
                                                            if true {
                                                                v1 += (v1 >> (((v35 | v35)) & 63))
                                                                v35 = (v35 | (v35 / ((v35) | 1)))
                                                                v2 = (v2 & v2)
                                                            } else {
                                                                v35 = 18446744073709551615
                                                                fmt.Println(fi91)
                                                                v98 := []int{(fi97 * v34), fi97, (fi97 % (((-(fi91))) | 1))}
                                                                _ = v98
                                                            }
                                                            v99 := []int{}
                                                            _ = v99
                                                        }
                                                        var v100 float32 = float32(float64(0.0))
                                                        switch "foo bar" {
                                                        case "hello":
                                                            v48 = 32767
                                                        case "a":
                                                            fmt.Println(v34)
                                                        case "":
                                                            v34 = 32768
                                                        default:
                                                            v3[848] = (1 >= 0)
                                                        }
                                                        v101 := GenS5{F0: (v1 + (^(v1))), F1: 1}
                                                        var v102 uint8 = 127
                                                        var v103 float64 = -2.25
                                                        if (fi91 != (fi97 * 42)) {
                                                            v104 := []int{(v34 % ((-303) | 1)), -1, fi97, fi97}
                                                            fmt.Println(v35)
                                                            _ = v104
                                                        }
                                                        _ = v100
                                                        _ = v101
                                                        _ = v102
                                                        _ = v103
                                                    }
                                                    fmt.Println(v35, v1)
                                                    v3[2147483648] = false
                                                }
                                            }
                                            _ = len(v3)
                                            var v105 int32 = -138
                                            v48 = (v48 | genFn0(2, true, 734))
                                            _ = fi97
                                            _ = v105
                                        }
                                    default:
                                        v106 := []int{}
                                        _ = v106
                                    }
                                default:
                                    switch 4294967295 {
                                    case 9223372036854775807:
                                        v34 -= fi91
                                    case 32767:
                                        for fi107 := 0; fi107 < 11; fi107++ {
                                            v108 := v3[834]
                                            if (100 <= -128) {
                                                switch v34 {
                                                case 127:
                                                    v109 := []int{fi91, fi107, (fi91 - fi91)}
                                                    _ = v109
                                                case 2147483647:
                                                    if v108 {
                                                        if (!false) {
                                                            fmt.Println(v1)
                                                            if (v1 > 197) {
                                                                var v110 int8 = 127
                                                                fmt.Println(v34, v2)
                                                                delete(v3, 65535)
                                                                fmt.Println(v108)
                                                                v34 = fi107
                                                                _ = v110
                                                            }
                                                            fmt.Println(v2)
                                                            v111 := []int{fi91, (^(v34)), fi91, (fi107 << ((uint(v1)) & 63))}
                                                            _ = v111
                                                        } else {
                                                            switch (fi91 * 32767) {
                                                            case 9223372036854775807:
                                                                var v112 float32 = 0.0
                                                                _ = v112
                                                            case -472:
                                                                v113 := []int{-382, -28, fi91}
                                                                _ = v113
                                                            case 2147483647:
                                                                var v114 uint8 = 46
                                                                _ = v114
                                                            default:
                                                                fmt.Println(v35, v108)
                                                            }
                                                        }
                                                    } else {
                                                        fmt.Println(v108, v108, v1)
                                                        fmt.Println(fi91, v48, v2)
                                                        v48 = (v1 ^ 42)
                                                    }
                                                case 42:
                                                    if v108 {
                                                        v115 := GenS13{F0: v2, F1: 4294967295, F2: 937}
                                                        fmt.Println(v1, fi107)
                                                        _ = v115
                                                    }
                                                }
                                                var v116 uint32 = uint32(v35)
                                                var v117 int16 = 32767
                                                _ = v116
                                                _ = v117
                                            }
                                            v118 := []int{fi91, (v34 % ((32767) | 1)), fi91, v34}
                                            _ = fi107
                                            _ = v108
                                            _ = v118
                                        }
                                    }
                                }
                            }
                            for fi119 := 0; fi119 < 8; fi119++ {
                                v120 := []int{}
                                fmt.Println(v35)
                                if false {
                                    switch 556 {
                                    case 32768:
                                        v121 := GenS14{F0: (v35 & (v35 >> ((65535) & 63))), F1: -2147483648}
                                        fmt.Println(v48, v2)
                                        v122 := GenS11{F0: -294}
                                        var v123 uint16 = genFn0(802, (v48 != (v1 | 0)), 0)
                                        _ = v121
                                        _ = v122
                                        _ = v123
                                    default:
                                        switch v34 {
                                        case 9223372036854775807:
                                            var v124 uint = v35
                                            _ = v124
                                        case 32768:
                                            v34 += fi91
                                        default:
                                            var v125 uint8 = 1
                                            _ = v125
                                        }
                                    }
                                }
                                _ = fi119
                                _ = v120
                            }
                            v35 -= 975
                            _ = fi91
                        }
                        v126 := GenS15{F0: 1, F1: 166, F2: 109}
                        _ = v126
                    }
                    var v127 float64 = (float64(1e10) / 3.14159)
                    _ = v127
                } else {
                    switch v34 {
                    case 2147483647:
                        fmt.Println(v48, v48)
                    case 42:
                        v128 := GenS16{F0: 543}
                        _ = v128
                    case 773:
                        for fi129 := 0; fi129 < 16; fi129++ {
                            for fi130 := 0; fi130 < 17; fi130++ {
                                if true {
                                    if true {
                                        v131 := GenS17{F0: true, F1: v2, F2: 1}
                                        v132 := GenS18{F0: (fi129 << ((v35) & 63))}
                                        if false {
                                            var v133 int32 = 2147483647
                                            if ((!(4294967295 >= 65535)) && true) {
                                                var v134 uint32 = 4294967295
                                                switch "\n\t" {
                                                case "\n\t":
                                                    if ((^(v35)) < (v35 >> ((2) & 63))) {
                                                        v48 = (v48 + (v48 << ((v35) & 63)))
                                                        v48 = uint16(v134)
                                                    }
                                                case "hello":
                                                    v135 := []int{(fi129 / (((v34 | (^(v34)))) | 1)), (fi129 / ((fi130) | 1)), fi129, (fi129 & (fi130 + fi130))}
                                                    _ = v135
                                                case "":
                                                    var v136 int64 = -9223372036854775808
                                                    _ = v136
                                                }
                                                _ = v134
                                            } else {
                                                if true {
                                                    v48 = 42
                                                } else {
                                                    switch (fi129 - fi130) {
                                                    case -1:
                                                        fmt.Println(v35)
                                                    }
                                                    var v137 bool = (-2.25 <= -2.25)
                                                    v138 := GenS19{F0: 255, F1: (v133 << (((v35 + 255)) & 63))}
                                                    _ = v137
                                                    _ = v138
                                                }
                                            }
                                            _ = v133
                                        }
                                        _ = v131
                                        _ = v132
                                    } else {
                                        v35 -= v35
                                        delete(v3, 887)
                                        var v139 uint = 81
                                        _ = v139
                                    }
                                }
                                v140 := GenS18{F0: (v34 - (v34 >> ((18446744073709551615) & 63)))}
                                switch (^(v34)) {
                                case 128:
                                    v141 := GenS19{F0: uint8(v35), F1: 811}
                                    _ = v141
                                case 42:
                                    var v142 uint64 = 18446744073709551615
                                    _ = v142
                                default:
                                    fmt.Println(fi129, v35, fi129)
                                }
                                _ = fi130
                                _ = v140
                            }
                            if (!false) {
                                v2 -= 1
                                v48 = v48
                            } else {
                                var v143 float32 = float32(float64(-2.25))
                                switch "\n\t" {
                                case "hello":
                                    for fi144 := 0; fi144 < 9; fi144++ {
                                        if false {
                                            switch fi144 {
                                            case 1:
                                                var v145 int32 = int32(v35)
                                                if ((v143 / 1.0) != v143) {
                                                    v143 = v143
                                                    switch (fi144 & fi129) {
                                                    case -9223372036854775808:
                                                        switch v34 {
                                                        case 32767:
                                                            v146 := []int{-974, (fi129 << (((v35 & (v35 * (v35 - (v35 * 255))))) & 63)), (v34 / (((fi144 + int(v1))) | 1)), v34}
                                                            _ = v146
                                                        case 42:
                                                            fmt.Println(v34, v145)
                                                        }
                                                    case 32768:
                                                        v1 = genFn0(265, (!true), 42)
                                                    }
                                                } else {
                                                    switch (fi144 + fi144) {
                                                    case 1:
                                                        v147 := []int{fi129, 1}
                                                        _ = v147
                                                    case -1:
                                                        delete(v3, 2147483648)
                                                    }
                                                    var v148 int64 = (v2 / ((-9223372036854775808) | 1))
                                                    v149 := v3[2]
                                                    _ = v148
                                                    _ = v149
                                                }
                                                _ = len(v3)
                                                _ = v145
                                            case -9223372036854775808:
                                                fmt.Println(v34)
                                            }
                                            var v150 uint64 = 18446744073709551615
                                            v151 := []int{}
                                            _ = v150
                                            _ = v151
                                        }
                                        v1 -= v48
                                        v3[65535] = true
                                        _ = fi144
                                    }
                                case "\n\t":
                                    v143 -= v143
                                }
                                switch ("foo bar" + "hello") {
                                case "\n\t":
                                    v152 := GenS6{F0: 2}
                                    _ = v152
                                case "":
                                    v34 += (fi129 & fi129)
                                case "a":
                                    for fi153 := 0; fi153 < 9; fi153++ {
                                        if false {
                                            delete(v3, uint32(v35))
                                            fmt.Println(fi129, fi129, v48)
                                        }
                                        _ = fi153
                                    }
                                default:
                                    v154 := []int{}
                                    _ = v154
                                }
                                _ = v143
                            }
                            _ = fi129
                        }
                    }
                    v155 := []int{(v34 & 32768), v34}
                    fmt.Println(v35)
                    _ = v155
                }
                for fi156 := 0; fi156 < 18; fi156++ {
                    var v157 int64 = 9223372036854775807
                    fmt.Println(fi156, v2, v34)
                    _ = fi156
                    _ = v157
                }
            }
            delete(v3, 4294967295)
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v48:", v48)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v47
    _ = v48
}
