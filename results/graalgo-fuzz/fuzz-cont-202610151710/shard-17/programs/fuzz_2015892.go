package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 string; F1 int32 }
type GenS3 struct { F0 uint; F1 int; F2 uint }
type GenS4 struct { F0 int32; F1 int16; F2 bool }
type GenS5 struct { F0 bool; F1 int8 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 uint; F1 uint; F2 bool }
type GenS9 struct { F0 bool }
type GenS10 struct { F0 uint }
type GenS11 struct { F0 uint16; F1 uint64 }
type GenS12 struct { F0 uint }
type GenS13 struct { F0 uint64; F1 uint64 }
type GenS14 struct { F0 int32; F1 int8; F2 int16 }
type GenS15 struct { F0 int16 }
type GenS16 struct { F0 uint }
type GenS17 struct { F0 int32; F1 uint32 }
type GenS18 struct { F0 uint64; F1 uint16 }
type GenS19 struct { F0 uint64 }
type GenS20 struct { F0 uint8; F1 int16 }
type GenS21 struct { F0 string }
type GenS22 struct { F0 bool }
func genFn0(p0 int16) int32 {
    v7 := []int{32768, -9223372036854775808, -2147483648, 2147483647}
    var v8 int16 = 0
    switch ("hello" + "") {
    case "foo bar":
        var v9 bool = ("" > "hello")
        _ = v9
    case "hello":
        var v10 int64 = -1
        _ = v10
    case "\n\t":
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
    }
    v8 -= -1
    var v11 uint32 = 0
    _ = p0
    _ = v11
    _ = v7
    _ = v8
    return 42
}
func genFn1(p0 uint32, p1 int8, p2 uint8) bool {
    for fi12 := 0; fi12 < 10; fi12++ {
        if false {
            for fi13 := 0; fi13 < 14; fi13++ {
                fmt.Println(fi13, p0)
                var v14 float64 = -2.25
                if (-1 > p1) {
                    fmt.Println(p2, p1)
                }
                _ = fi13
                _ = v14
            }
            p2 += p2
            if true {
                fmt.Println(p0)
                if (!true) {
                    var v15 uint8 = p2
                    _ = v15
                } else {
                    var v16 uint = 1
                    _ = v16
                }
                v17 := GenS1{F0: 1000}
                _ = v17
            } else {
                v18 := GenS1{F0: -32768}
                switch 127 {
                case -267:
                    p1 -= -8
                case 32768:
                    v19 := GenS1{F0: 1}
                    _ = v19
                }
                if ((0.0 * -2.25) > 0.0) {
                    var v20 bool = (18446744073709551615 < 2)
                    _ = v20
                } else {
                    v21 := []int{fi12}
                    switch (fi12 / ((fi12) | 1)) {
                    case 642:
                        p0 -= p0
                    case 219:
                        p0 += p0
                    default:
                        for fi22 := 0; fi22 < 17; fi22++ {
                            var v23 string = "foo bar"
                            var v24 uint16 = 821
                            if (false || (1000 >= 0)) {
                                if true {
                                    fmt.Println(p2, v24, v23)
                                    if true {
                                        if ((p0 >> ((18446744073709551615) & 63)) >= (p0 + p0)) {
                                            v25 := GenS1{F0: 32767}
                                            if false {
                                                if true {
                                                    if true {
                                                        var v26 string = v23
                                                        v24 -= 812
                                                        var v27 int32 = 1
                                                        _ = v26
                                                        _ = v27
                                                    }
                                                    p0 += p0
                                                    switch -2147483648 {
                                                    case 9223372036854775807:
                                                        p2 = 128
                                                    case -9223372036854775808:
                                                        if false {
                                                            if ((v24 * (^(v24))) > (v24 << ((0) & 63))) {
                                                                p0 = p0
                                                                var v28 int16 = -1
                                                                _ = v28
                                                            }
                                                            if (true && (!false)) {
                                                                var v29 uint = 2
                                                                if true {
                                                                    v23 = (v23 + (("a" + ("" + v23)) + ((("foo bar" + "hello") + (v23 + (v23 + ((v23 + v23) + v23)))) + v23)))
                                                                } else {
                                                                    v30 := GenS1{F0: 292}
                                                                    var v31 int8 = (p1 - p1)
                                                                    _ = v30
                                                                    _ = v31
                                                                }
                                                                v29 = v29
                                                                _ = v29
                                                            } else {
                                                                v24 = (v24 & v24)
                                                                switch v23 {
                                                                case "":
                                                                    v24 = 498
                                                                case "\n\t":
                                                                    fmt.Println(p1)
                                                                case "foo bar":
                                                                    var v32 uint = 871
                                                                    _ = v32
                                                                default:
                                                                    v21 = append(v21, int(p1))
                                                                }
                                                                fmt.Println(p0, p1, p1)
                                                            }
                                                            p0 -= (^(p0))
                                                            if len(v21) >= 2 {
                                                                v21 = v21[1:len(v21)]
                                                            }
                                                            v33 := GenS1{F0: 42}
                                                            p0 -= 4294967295
                                                            var v34 string = "a"
                                                            if true {
                                                                v24 = (v24 << ((1) & 63))
                                                                fmt.Println(v34, v23, p2)
                                                            }
                                                            if false {
                                                                v35 := GenS2{F0: v23, F1: -2147483648}
                                                                _ = v35
                                                            } else {
                                                                if (!(((true && (((^(p2)) != 2) && true)) && false) && (156 != p0))) {
                                                                    var v36 uint32 = p0
                                                                    _ = v36
                                                                }
                                                            }
                                                            v37 := GenS1{F0: -32768}
                                                            var v38 bool = true
                                                            _ = v33
                                                            _ = v34
                                                            _ = v37
                                                            _ = v38
                                                        } else {
                                                            v39 := GenS2{F0: "", F1: 1}
                                                            v40 := GenS3{F0: 0, F1: 32768, F2: 42}
                                                            var v41 string = v23
                                                            _ = v39
                                                            _ = v40
                                                            _ = v41
                                                        }
                                                    case 4294967295:
                                                        _ = len(v21)
                                                    }
                                                    fmt.Println(p1, fi22, fi22)
                                                    v21 = append(v21, (fi12 & fi12))
                                                    switch v23 {
                                                    case "a":
                                                        if (!(!false)) {
                                                            fmt.Println(p2, v24)
                                                            var v42 string = "hello"
                                                            _ = v42
                                                        }
                                                    case "\n\t":
                                                        fmt.Println(p0, v23, p2)
                                                    default:
                                                        v23 = v23
                                                    }
                                                    if false {
                                                        var v43 string = (v23 + v23)
                                                        v44 := GenS1{F0: 0}
                                                        var v45 int64 = int64(p2)
                                                        switch (fi12 * (fi22 >> ((4294967295) & 63))) {
                                                        case 32768:
                                                            v43 = v43
                                                        case -1:
                                                            if ((0 < 32767) && (481 > uint(v24))) {
                                                                if true {
                                                                    var v46 float32 = -2.25
                                                                    switch ("a" + v43) {
                                                                    case "a":
                                                                        p1 = -128
                                                                    }
                                                                    switch (v23 + (v43 + v23)) {
                                                                    case "a":
                                                                        v47 := GenS4{F0: 42, F1: 285, F2: false}
                                                                        _ = v47
                                                                    }
                                                                    _ = v46
                                                                } else {
                                                                    switch "\n\t" {
                                                                    case "a":
                                                                        var v48 uint = 2
                                                                        _ = v48
                                                                    case "foo bar":
                                                                        var v49 int16 = -1
                                                                        _ = v49
                                                                    case "hello":
                                                                        var v50 int64 = (-(v45))
                                                                        _ = v50
                                                                    }
                                                                    fmt.Println(fi22, p0)
                                                                    if false {
                                                                        if (true && (((420 <= 30) || (255 != p2)) && (32767 >= 1))) {
                                                                            v51 := GenS3{F0: 2, F1: fi22, F2: 4294967295}
                                                                            fmt.Println(v23, p1, v24)
                                                                            p2 -= 193
                                                                            _ = v51
                                                                        }
                                                                        var v52 int8 = int8(v24)
                                                                        _ = v52
                                                                    } else {
                                                                        if len(v21) >= 2 {
                                                                            v21 = v21[1:len(v21)]
                                                                        }
                                                                    }
                                                                }
                                                                fmt.Println(p2, v24)
                                                                v53 := GenS5{F0: false, F1: 112}
                                                                _ = v53
                                                            }
                                                        case 42:
                                                            var v54 uint32 = 439
                                                            _ = v54
                                                        default:
                                                            if len(v21) >= 2 {
                                                                v21 = v21[1:len(v21)]
                                                            }
                                                        }
                                                        _ = v43
                                                        _ = v44
                                                        _ = v45
                                                    } else {
                                                        var v55 string = "\n\t"
                                                        if false {
                                                            v56 := GenS6{F0: 682}
                                                            switch (v55 + v23) {
                                                            case "foo bar":
                                                                fmt.Println(p2, fi12, v55)
                                                            case "":
                                                                fmt.Println(fi22, p2, p2)
                                                            case "\n\t":
                                                                if false {
                                                                    if (1.5 != -2.25) {
                                                                        var v57 int16 = 1
                                                                        switch ("hello" + v23) {
                                                                        case "a":
                                                                            v58 := GenS1{F0: 42}
                                                                            _ = v58
                                                                        case "foo bar":
                                                                            v59 := GenS6{F0: 785}
                                                                            _ = v59
                                                                        case "hello":
                                                                            if (!false) {
                                                                                switch v55 {
                                                                                case "foo bar":
                                                                                    p0 -= (p0 * p0)
                                                                                    if false {
                                                                                        var v60 int32 = 279
                                                                                        _ = v60
                                                                                    } else {
                                                                                        p2 = p2
                                                                                        var v61 int32 = 46
                                                                                        _ = v61
                                                                                    }
                                                                                    fmt.Println(v57)
                                                                                default:
                                                                                    var v62 float64 = 1e10
                                                                                    _ = v62
                                                                                }
                                                                            }
                                                                        }
                                                                        p1 = p1
                                                                        v24 = v24
                                                                        _ = v57
                                                                    }
                                                                    v63 := GenS5{F0: false, F1: (p1 % ((p1) | 1))}
                                                                    var v64 bool = true
                                                                    _ = len(v21)
                                                                    var v65 float32 = 0.1
                                                                    _ = v63
                                                                    _ = v64
                                                                    _ = v65
                                                                } else {
                                                                    _ = len(v21)
                                                                    fmt.Println(v24, v23)
                                                                    switch (-(fi12)) {
                                                                    case -1:
                                                                        switch "\n\t" {
                                                                        case "a":
                                                                            fmt.Println(fi12)
                                                                        }
                                                                    default:
                                                                        if (!("a" != v55)) {
                                                                            var v66 int32 = 985
                                                                            _ = v66
                                                                        } else {
                                                                            v24 = 657
                                                                        }
                                                                    }
                                                                    v24 -= v24
                                                                    if true {
                                                                        if ((p2 * p2) >= 91) {
                                                                            var v67 uint32 = 2147483648
                                                                            fmt.Println(v55, v67)
                                                                            var v68 uint8 = 78
                                                                            v68 = 1
                                                                            fmt.Println(p1, fi22, fi12)
                                                                            _ = v67
                                                                            _ = v68
                                                                        } else {
                                                                            var v69 uint16 = v24
                                                                            if (1.5 > 0.0) {
                                                                                v24 += v24
                                                                                var v70 uint64 = 328
                                                                                _ = v70
                                                                            }
                                                                            switch v55 {
                                                                            case "":
                                                                                var v71 uint64 = uint64(p0)
                                                                                _ = v71
                                                                            case "\n\t":
                                                                                v23 = ("a" + "hello")
                                                                            case "hello":
                                                                                var v72 int = 127
                                                                                var v73 int16 = 1000
                                                                                var v74 bool = (true || ((v73 >> ((931) & 63)) >= 1))
                                                                                _ = v72
                                                                                _ = v73
                                                                                _ = v74
                                                                            }
                                                                            _ = v69
                                                                        }
                                                                        fmt.Println(p1)
                                                                        p2 = p2
                                                                    }
                                                                }
                                                            default:
                                                                var v75 float64 = 3.14159
                                                                _ = v75
                                                            }
                                                            if true {
                                                                if (!(42 > 597)) {
                                                                    if len(v21) > 0 {
                                                                        v21[0] = (fi22 / ((9223372036854775807) | 1))
                                                                    }
                                                                    var v76 uint32 = 387
                                                                    v24 -= v24
                                                                    _ = v76
                                                                }
                                                            }
                                                            _ = v56
                                                        } else {
                                                            v77 := GenS2{F0: "\n\t", F1: -1}
                                                            p2 = uint8(p0)
                                                            v78 := GenS7{F0: (0.0 < (-2.25 * (0.1 * 1.5)))}
                                                            _ = v77
                                                            _ = v78
                                                        }
                                                        _ = v55
                                                    }
                                                } else {
                                                    if true {
                                                        v23 = v23
                                                        var v79 int64 = -9223372036854775808
                                                        var v80 float32 = -2.25
                                                        var v81 int = (fi12 % ((fi22) | 1))
                                                        if (true || true) {
                                                            var v82 int16 = int16(p1)
                                                            v24 = (v24 * v24)
                                                            _ = v82
                                                        } else {
                                                            v83 := GenS8{F0: 1, F1: 18446744073709551615, F2: false}
                                                            _ = v83
                                                        }
                                                        _ = v79
                                                        _ = v80
                                                        _ = v81
                                                    } else {
                                                        p1 += 39
                                                        _ = len(v21)
                                                    }
                                                }
                                                var v84 int8 = p1
                                                v21 = append(v21, fi22)
                                                _ = v84
                                            }
                                            var v85 int64 = -9223372036854775808
                                            _ = v25
                                            _ = v85
                                        }
                                        switch fi12 {
                                        case 990:
                                            var v86 uint8 = 1
                                            _ = v86
                                        case 4294967295:
                                            switch fi22 {
                                            case -797:
                                                switch ("hello" + v23) {
                                                case "hello":
                                                    v21 = append(v21, fi22)
                                                default:
                                                    fmt.Println(fi12)
                                                }
                                            }
                                        default:
                                            var v87 int8 = (p1 - -1)
                                            _ = v87
                                        }
                                        if false {
                                            v88 := GenS9{F0: false}
                                            _ = len(v21)
                                            p2 -= (p2 & uint8(p0))
                                            _ = v88
                                        } else {
                                            v21 = append(v21, 9223372036854775807)
                                            var v89 int16 = -1
                                            _ = v89
                                        }
                                    }
                                    var v90 uint64 = 4294967295
                                    var v91 uint64 = v90
                                    _ = v90
                                    _ = v91
                                } else {
                                    p2 += 49
                                    var v92 bool = false
                                    _ = len(v21)
                                    if (42 == 0) {
                                        v21 = append(v21, (-(fi22)))
                                        v93 := GenS7{F0: (v92 || (79 >= 218))}
                                        p0 += uint32(p2)
                                        _ = v93
                                    }
                                    _ = v92
                                }
                                var v94 int16 = int16(v24)
                                _ = v94
                            }
                            _ = fi22
                            _ = v23
                            _ = v24
                        }
                    }
                    if false {
                        if true {
                            if len(v21) > 0 {
                                v21[0] = fi12
                            }
                        } else {
                            var v95 uint = 620
                            p1 += p1
                            _ = v95
                        }
                    }
                    _ = v21
                }
                _ = v18
            }
        }
        var v96 uint = 42
        _ = fi12
        _ = v96
    }
    p0 = 666
    var v97 float64 = 0.0
    if true {
        for fi98 := 0; fi98 < 16; fi98++ {
            fmt.Println(p1)
            var v99 float32 = 0.0
            fmt.Println(fi98, p0)
            _ = fi98
            _ = v99
        }
        p0 = 42
    } else {
        for fi100 := 0; fi100 < 3; fi100++ {
            v101 := []int{127, int(p0), 128}
            _ = fi100
            _ = v101
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v97
    return (1 != 647)
}

func main() {
    v1 := []uint32{}
    v2 := make(map[uint32]string)
    var v3 float32 = float32(0.1)
    v4 := make(map[int]uint32)
    var v5 string = "\n\t"
    v6 := make(map[int]string)
    if (0 < 32767) {
        fmt.Println(v3)
        fmt.Println(v5, v5)
        if (0 >= 815) {
            fmt.Println(v3, v3)
        }
    } else {
        v3 += v3
        for fi102 := 0; fi102 < 15; fi102++ {
            v103 := GenS10{F0: 71}
            var v104 uint8 = 255
            _ = fi102
            _ = v103
            _ = v104
        }
    }
    if false {
        var v105 uint = 656
        _ = v105
    }
    var v106 uint16 = 533
    switch 269 {
    case 128:
        if len(v1) > 0 {
            v1[0] = 0
        }
    case 850:
        _ = len(v1)
    case 949:
        fmt.Println(v3)
    default:
        if (!false) {
            for fi107 := 0; fi107 < 18; fi107++ {
                if ((3.14159 == float64(float32((0.0 + 1.5)))) && (v5 != ("" + "hello"))) {
                    if (true || false) {
                        fmt.Println(v5)
                        switch v5 {
                        case "":
                            fmt.Println(fi107, v3, v106)
                        }
                    } else {
                        v108 := GenS11{F0: (v106 ^ (v106 - 32767)), F1: 18446744073709551615}
                        _ = v108
                    }
                    switch -637 {
                    case 0:
                        for fi109 := 0; fi109 < 19; fi109++ {
                            switch ("a" + ("\n\t" + (v5 + "hello"))) {
                            case "foo bar":
                                var v110 int16 = 157
                                _ = v110
                            case "":
                                v6[fi109] = v5
                            case "hello":
                                switch fi107 {
                                case 942:
                                    v3 += (v3 + v3)
                                case -9223372036854775808:
                                    v106 += v106
                                    v5 = v5
                                case 42:
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                }
                            default:
                                v106 = 65535
                            }
                            _ = fi109
                        }
                    case 32768:
                        var v111 float64 = float64((((1e10 + 1e10) - v3) * v3))
                        _ = v111
                    case 128:
                        if (!false) {
                            _ = len(v6)
                            fmt.Println(v3, v5, fi107)
                            if true {
                                _ = len(v1)
                                var v112 uint8 = 2
                                delete(v6, (fi107 & (-(fi107))))
                                _ = v112
                            }
                        } else {
                            switch 2147483647 {
                            case -1:
                                for fi113 := 0; fi113 < 10; fi113++ {
                                    v3 = (v3 + v3)
                                    var v114 int64 = 42
                                    v115 := v4[(-(fi107))]
                                    _ = fi113
                                    _ = v114
                                    _ = v115
                                }
                            case 9223372036854775807:
                                for fi116 := 0; fi116 < 5; fi116++ {
                                    switch 42 {
                                    case 42:
                                        var v117 int64 = -9223372036854775808
                                        _ = v117
                                    case 129:
                                        if (false && false) {
                                            v118 := GenS10{F0: 65535}
                                            _ = v118
                                        } else {
                                            var v119 float64 = 0.0
                                            _ = v119
                                        }
                                    case 1:
                                        if (4294967295 > 42) {
                                            switch fi116 {
                                            case 945:
                                                switch v5 {
                                                case "":
                                                    var v120 string = v5
                                                    _ = v120
                                                default:
                                                    _ = len(v2)
                                                }
                                            case -9223372036854775808:
                                                _ = len(v1)
                                            case -1:
                                                if true {
                                                    switch fi116 {
                                                    case -929:
                                                        if true {
                                                            if false {
                                                                var v121 float64 = (0.1 * 0.0)
                                                                var v122 string = (v5 + v5)
                                                                var v123 uint16 = (v106 | v106)
                                                                switch fi107 {
                                                                case 2147483647:
                                                                    v121 -= v121
                                                                case 9223372036854775807:
                                                                    if ((!((true || (-1 <= 228)) && true)) && false) {
                                                                        switch fi107 {
                                                                        case 32768:
                                                                            v106 += (v106 * v106)
                                                                        case -702:
                                                                            var v124 float64 = v121
                                                                            _ = v124
                                                                        case 32767:
                                                                            var v125 int16 = -32768
                                                                            _ = v125
                                                                        }
                                                                    } else {
                                                                        switch (((v5 + v122) + "\n\t") + v5) {
                                                                        case "":
                                                                            var v126 string = v122
                                                                            _ = v126
                                                                        case "\n\t":
                                                                            v127 := GenS10{F0: 2}
                                                                            _ = v127
                                                                        case "a":
                                                                            v121 = (v121 * -2.25)
                                                                        }
                                                                        v5 = (v122 + "a")
                                                                    }
                                                                }
                                                                fmt.Println(v3)
                                                                v1 = append(v1, 2147483648)
                                                                _ = v121
                                                                _ = v122
                                                                _ = v123
                                                            }
                                                            if false {
                                                                if false {
                                                                    var v128 int16 = 1000
                                                                    _ = v128
                                                                }
                                                            }
                                                            v106 = (^(v106))
                                                        }
                                                    case -2147483648:
                                                        if false {
                                                            var v129 int16 = -1
                                                            fmt.Println(v129)
                                                            if false {
                                                                var v130 bool = (fi107 < (fi116 >> ((259) & 63)))
                                                                if (42 >= 131) {
                                                                    v131 := GenS9{F0: ((^(v129)) >= v129)}
                                                                    _ = v131
                                                                } else {
                                                                    switch (v5 + "foo bar") {
                                                                    case "foo bar":
                                                                        var v132 uint = 808
                                                                        _ = v132
                                                                    case "a":
                                                                        _ = len(v1)
                                                                    case "\n\t":
                                                                        if ("a" <= ("hello" + v5)) {
                                                                            var v133 string = (v5 + v5)
                                                                            v130 = true
                                                                            var v134 int32 = genFn0(-130)
                                                                            _ = v133
                                                                            _ = v134
                                                                        }
                                                                    default:
                                                                        v135 := GenS11{F0: v106, F1: 0}
                                                                        _ = v135
                                                                    }
                                                                    var v136 uint8 = 1
                                                                    _ = v136
                                                                }
                                                                _ = v130
                                                            }
                                                            _ = v129
                                                        } else {
                                                            if true {
                                                                if len(v1) > 0 {
                                                                    v1[0] = 236
                                                                }
                                                                v1 = append(v1, 52)
                                                            }
                                                            v106 += (v106 / (((^(v106))) | 1))
                                                            var v137 uint32 = 42
                                                            v137 -= 0
                                                            _ = v137
                                                        }
                                                        v138 := GenS8{F0: 4294967295, F1: 2, F2: false}
                                                        if len(v1) > 0 {
                                                            v1[0] = 2147483648
                                                        }
                                                        _ = v138
                                                    case 4294967295:
                                                        _ = len(v1)
                                                    }
                                                    if false {
                                                        v106 += v106
                                                    }
                                                    fmt.Println(fi107, fi107)
                                                } else {
                                                    switch (fi116 * (fi116 / (((fi116 + fi107)) | 1))) {
                                                    case 128:
                                                        v106 += v106
                                                    case -2147483648:
                                                        if (9223372036854775807 != 9223372036854775807) {
                                                            v3 -= (v3 + v3)
                                                        } else {
                                                            _ = len(v6)
                                                            var v139 int = fi116
                                                            v3 -= ((0.0 + 0.1) / (v3 + (v3 + (1e10 / 0.1))))
                                                            _ = v139
                                                        }
                                                    case -774:
                                                        if len(v1) > 0 {
                                                            v1[0] = 4294967295
                                                        }
                                                    default:
                                                        v140 := GenS11{F0: v106, F1: 0}
                                                        _ = v140
                                                    }
                                                    var v141 float32 = (v3 - v3)
                                                    if len(v1) > 0 {
                                                        v1[0] = 337
                                                    }
                                                    switch ((v5 + (v5 + v5)) + v5) {
                                                    case "":
                                                        switch "foo bar" {
                                                        case "hello":
                                                            var v142 bool = (-291 < 0)
                                                            _ = v142
                                                        case "a":
                                                            v143 := GenS6{F0: 990}
                                                            _ = v143
                                                        case "\n\t":
                                                            v144 := GenS12{F0: 506}
                                                            _ = v144
                                                        default:
                                                            v3 += 0.0
                                                        }
                                                    case "a":
                                                        fmt.Println(v106, v141, fi116)
                                                    case "hello":
                                                        if false {
                                                            var v145 uint8 = uint8(v106)
                                                            _ = v145
                                                        }
                                                    default:
                                                        switch "" {
                                                        case "hello":
                                                            var v146 uint8 = 216
                                                            _ = v146
                                                        default:
                                                            if (127 < 0) {
                                                                fmt.Println(v5, fi107, v3)
                                                                if false {
                                                                    if true {
                                                                        if (true && true) {
                                                                            v5 = (v5 + v5)
                                                                            var v147 int = (^(fi107))
                                                                            v141 = 1.5
                                                                            _ = v147
                                                                        } else {
                                                                            var v148 int = fi116
                                                                            _ = v148
                                                                        }
                                                                        var v149 uint8 = 1
                                                                        v150 := GenS13{F0: 4294967295, F1: 0}
                                                                        _ = v149
                                                                        _ = v150
                                                                    }
                                                                }
                                                            }
                                                        }
                                                    }
                                                    var v151 string = v5
                                                    var v152 int = (fi107 << ((18446744073709551615) & 63))
                                                    if ((1e10 / (((3.14159 / -2.25) + 3.14159) * float64(v3))) > 1.5) {
                                                        v3 = v141
                                                        v3 += v141
                                                    }
                                                    _ = v141
                                                    _ = v151
                                                    _ = v152
                                                }
                                            }
                                            var v153 string = v5
                                            _ = v153
                                        }
                                    default:
                                        fmt.Println(v5, v106, v106)
                                    }
                                    v154 := GenS14{F0: -2147483648, F1: 106, F2: -32768}
                                    var v155 int = fi116
                                    _ = fi116
                                    _ = v154
                                    _ = v155
                                }
                            default:
                                var v156 float32 = v3
                                _ = v156
                            }
                            if (-1 <= -606) {
                                v5 = v5
                            } else {
                                v3 += v3
                            }
                            v106 += (v106 % (((v106 % (((v106 ^ v106)) | 1))) | 1))
                        }
                    default:
                        fmt.Println(v106, v3, v106)
                    }
                    for fi157 := 0; fi157 < 11; fi157++ {
                        if (!false) {
                            if (!true) {
                                v158 := GenS2{F0: (v5 + "foo bar"), F1: 437}
                                var v159 int8 = 42
                                if true {
                                    v106 = (v106 << ((uint(v106)) & 63))
                                    v5 = (((v5 + v5) + (v5 + (((("a" + v5) + v5) + "\n\t") + v5))) + (v5 + v5))
                                } else {
                                    var v160 uint = 255
                                    _ = v160
                                }
                                _ = v158
                                _ = v159
                            }
                            fmt.Println(v106, v106)
                        }
                        if len(v1) > 0 {
                            v1[0] = 2147483648
                        }
                        _ = fi157
                    }
                }
                if (!(genFn0(42) >= genFn0(-32768))) {
                    v3 -= v3
                    if false {
                        v161 := GenS15{F0: -1}
                        for fi162 := 0; fi162 < 10; fi162++ {
                            if len(v1) > 0 {
                                v1[0] = 122
                            }
                            if true {
                                fmt.Println(fi162)
                                v163 := GenS13{F0: 300, F1: 571}
                                switch (v5 + ((v5 + (v5 + v5)) + (v5 + ((v5 + "a") + "a")))) {
                                case "hello":
                                    v164 := GenS1{F0: -1}
                                    if false {
                                        var v165 bool = (!false)
                                        v166 := GenS16{F0: 4294967295}
                                        var v167 float32 = float32(0.0)
                                        var v168 uint8 = 0
                                        _ = v165
                                        _ = v166
                                        _ = v167
                                        _ = v168
                                    } else {
                                        var v169 int8 = -1
                                        _ = v169
                                    }
                                    fmt.Println(fi162, fi107)
                                    _ = v164
                                case "":
                                    var v170 uint16 = v106
                                    _ = v170
                                case "\n\t":
                                    var v171 float32 = ((v3 - ((0.1 - (v3 * v3)) / 1e10)) * v3)
                                    _ = v171
                                }
                                _ = v163
                            }
                            fmt.Println(fi162, v3, fi107)
                            v5 = "hello"
                            _ = fi162
                        }
                        _ = v161
                    } else {
                        switch v5 {
                        case "foo bar":
                            switch fi107 {
                            case -669:
                                v2[2147483648] = v5
                            }
                        }
                    }
                }
                _ = fi107
            }
            v172 := GenS17{F0: -1, F1: 2}
            var v173 uint8 = 255
            _ = v172
            _ = v173
        } else {
            fmt.Println(v106, v106, v5)
            switch v5 {
            case "hello":
                v174 := GenS18{F0: 4294967295, F1: (v106 & 1)}
                _ = v174
            case "":
                for fi175 := 0; fi175 < 12; fi175++ {
                    for fi176 := 0; fi176 < 10; fi176++ {
                        var v177 int = -2147483648
                        _ = fi176
                        _ = v177
                    }
                    _ = fi175
                }
            default:
                fmt.Println(v5, v3, v106)
            }
            v3 = 0.1
        }
    }
    for fi178 := 0; fi178 < 2; fi178++ {
        switch (("foo bar" + v5) + v5) {
        case "hello":
            switch fi178 {
            case 0:
                var v179 int32 = 368
                _ = v179
            case 42:
                fmt.Println(v5)
            case 959:
                fmt.Println(v5, fi178)
            }
        }
        _ = fi178
    }
    v180 := GenS6{F0: 42}
    switch 127 {
    case 1:
        if (!(!((1 < 0) || (!false)))) {
            v106 = v106
        }
    }
    var v181 uint = 255
    switch -86 {
    case 42:
        fmt.Println(v106, v106, v5)
    case -545:
        v106 += 255
    }
    var v182 uint16 = (v106 ^ (^(v106)))
    v183 := GenS19{F0: 1}
    var v184 uint64 = 18446744073709551615
    switch 32768 {
    case 0:
        for fi185 := 0; fi185 < 10; fi185++ {
            switch (v5 + "\n\t") {
            case "hello":
                var v186 string = ""
                _ = v186
            case "":
                var v187 int = 0
                _ = v187
            case "\n\t":
                v184 = 103
            default:
                var v188 int = (fi185 + (fi185 % ((fi185) | 1)))
                _ = v188
            }
            v3 += v3
            switch (v5 + v5) {
            case "":
                v182 += 1
            case "\n\t":
                v5 = v5
            case "foo bar":
                v189 := GenS20{F0: 42, F1: 42}
                _ = v189
            default:
                if ("a" <= v5) {
                    delete(v2, 2147483648)
                } else {
                    v190 := GenS15{F0: -1}
                    _ = v190
                }
            }
            _ = fi185
        }
    case 1:
        for fi191 := 0; fi191 < 9; fi191++ {
            v5 = v5
            switch v5 {
            case "foo bar":
                if false {
                    var v192 uint16 = (v182 >> (((v181 | (v181 % ((v181) | 1)))) & 63))
                    _ = len(v1)
                    v1 = append(v1, 1)
                    _ = v192
                } else {
                    switch (fi191 * -2147483648) {
                    case 127:
                        fmt.Println(v3, v182, v184)
                    case 42:
                        var v193 uint64 = v184
                        _ = v193
                    }
                    for fi194 := 0; fi194 < 11; fi194++ {
                        var v195 int8 = 0
                        if ((v181 * v181) == v181) {
                            var v196 int = fi191
                            var v197 int64 = -9223372036854775808
                            _ = v196
                            _ = v197
                        }
                        var v198 int32 = 1
                        _ = fi194
                        _ = v195
                        _ = v198
                    }
                    for fi199 := 0; fi199 < 4; fi199++ {
                        if false {
                            fmt.Println(v184, v182)
                        } else {
                            if ((((v5 + v5) + v5) + v5) != (((v5 + "") + v5) + v5)) {
                                v4[fi199] = 912
                                var v200 int16 = 233
                                fmt.Println(fi191)
                                if (100 <= int8(v200)) {
                                    if (true || true) {
                                        v3 = (-2.25 * 1e10)
                                        fmt.Println(v106, v200)
                                    } else {
                                        v182 -= uint16(v181)
                                        v200 -= 32767
                                    }
                                    delete(v2, 42)
                                    if true {
                                        if ((float32(1.5) * ((1e10 * v3) - (0.1 + v3))) == v3) {
                                            if true {
                                                if len(v1) > 0 {
                                                    v1[0] = 908
                                                }
                                                if true {
                                                    v201 := GenS21{F0: (v5 + "foo bar")}
                                                    v184 = 712
                                                    _ = v201
                                                }
                                            }
                                        } else {
                                            v184 -= (^(v184))
                                        }
                                    }
                                }
                                switch 1 {
                                case -386:
                                    v5 = v5
                                default:
                                    switch (fi199 >> ((v181) & 63)) {
                                    case 0:
                                        var v202 int = (fi191 * (fi191 / ((128) | 1)))
                                        _ = v202
                                    }
                                }
                                switch (fi191 ^ fi191) {
                                case 32768:
                                    var v203 float32 = 0.0
                                    _ = v203
                                }
                                var v204 int16 = v200
                                v205 := GenS22{F0: (!(false && true))}
                                if (!((!true) && (v106 >= (^(v106))))) {
                                    v206 := GenS14{F0: genFn0(v200), F1: 100, F2: v200}
                                    fmt.Println(v106)
                                    _ = v206
                                }
                                _ = v200
                                _ = v204
                                _ = v205
                            }
                        }
                        _ = fi199
                    }
                }
            case "":
                for fi207 := 0; fi207 < 4; fi207++ {
                    v3 += -2.25
                    v1 = append(v1, 2147483647)
                    _ = fi207
                }
            case "a":
                v184 -= (v184 >> (((v181 * (v181 >> (((v181 >> ((v181) & 63))) & 63)))) & 63))
            default:
                v181 = uint(v184)
            }
            v5 = ""
            _ = fi191
        }
    case 2147483647:
        switch "hello" {
        case "\n\t":
            var v208 int32 = 1
            _ = v208
        }
    }
    v5 = (v5 + v5)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v106:", v106)
    fmt.Println("v181:", v181)
    fmt.Println("v182:", v182)
    fmt.Println("v184:", v184)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v106
    _ = v180
    _ = v181
    _ = v182
    _ = v183
    _ = v184
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
