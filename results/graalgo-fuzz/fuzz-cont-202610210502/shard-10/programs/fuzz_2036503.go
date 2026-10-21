package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32 }
type GenS2 struct { F0 uint8; F1 int16; F2 int }
type GenS3 struct { F0 uint8; F1 uint64 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int16; F1 int64; F2 int8 }
type GenS6 struct { F0 uint; F1 int32 }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 int8 }
type GenS9 struct { F0 uint32; F1 int32; F2 uint8 }

func main() {
    var v1 int32 = 2147483647
    var v2 int16 = 0
    var v3 int = 1
    var v4 int8 = 1
    var v5 int8 = v4
    for fi6 := 0; fi6 < 19; fi6++ {
        var v7 uint32 = 301
        _ = fi6
        _ = v7
    }
    var v8 string = "\n\t"
    var v9 int8 = (^(v4))
    switch v8 {
    case "\n\t":
        v10 := []int{4294967295, (v3 + v3)}
        _ = v10
    case "hello":
        for fi11 := 0; fi11 < 6; fi11++ {
            switch (v3 >> ((336) & 63)) {
            case 42:
                if false {
                    v3 += v3
                }
            case -1:
                var v12 uint32 = 640
                _ = v12
            case 173:
                if false {
                    if true {
                        v5 -= (v9 - (v4 ^ v9))
                    }
                    if (true && false) {
                        v4 = v9
                        for fi13 := 0; fi13 < 8; fi13++ {
                            v5 = (v9 / (((^(v5))) | 1))
                            if ((v3 ^ (fi13 + (v3 / ((v3) | 1)))) > 42) {
                                if (v8 <= v8) {
                                    v14 := GenS1{F0: 42, F1: 223}
                                    var v15 int = int(v1)
                                    var v16 uint32 = 427
                                    _ = v14
                                    _ = v15
                                    _ = v16
                                }
                                fmt.Println(v1, fi11, v4)
                            } else {
                                v17 := []int{(v3 & (^(fi13)))}
                                v17 = append(v17, v3)
                                _ = v17
                            }
                            fmt.Println(v4, v2, v8)
                            switch "\n\t" {
                            case "hello":
                                v18 := []int{(fi13 * (-(fi11)))}
                                _ = v18
                            default:
                                var v19 int16 = (v2 << ((2) & 63))
                                _ = v19
                            }
                            _ = fi13
                        }
                        for fi20 := 0; fi20 < 18; fi20++ {
                            v4 -= (v5 | 127)
                            _ = fi20
                        }
                    }
                } else {
                    v21 := GenS2{F0: 0, F1: v2, F2: (^(v3))}
                    switch fi11 {
                    case 1:
                        v22 := GenS1{F0: 0, F1: 122}
                        _ = v22
                    }
                    switch (("\n\t" + "a") + v8) {
                    case "foo bar":
                        if false {
                            var v23 uint16 = 2
                            v4 = v5
                            _ = v23
                        }
                    case "a":
                        v24 := GenS3{F0: 0, F1: 1}
                        _ = v24
                    }
                    _ = v21
                }
            }
            _ = fi11
        }
    case "foo bar":
        if false {
            var v25 bool = (true || false)
            _ = v25
        }
    default:
        v26 := GenS4{F0: 1}
        _ = v26
    }
    switch v8 {
    case "\n\t":
        v27 := []int{v3}
        _ = v27
    case "a":
        v8 = "hello"
    }
    v9 = v5
    fmt.Println(v8, v4)
    switch (v8 + v8) {
    case "a":
        v9 = v4
    case "":
        var v28 bool = (false && false)
        _ = v28
    }
    for fi29 := 0; fi29 < 2; fi29++ {
        if (0 != 9223372036854775807) {
            if (0 == 18446744073709551615) {
                v30 := []int{int(v1), -136, (v3 - 32767)}
                switch v8 {
                case "hello":
                    v5 = (^(v9))
                case "\n\t":
                    var v31 int64 = -9223372036854775808
                    _ = v31
                case "a":
                    for fi32 := 0; fi32 < 18; fi32++ {
                        if false {
                            v33 := GenS1{F0: 1, F1: -1}
                            _ = v33
                        }
                        switch 0 {
                        case 128:
                            var v34 int16 = v2
                            _ = v34
                        case 0:
                            if len(v30) > 0 {
                                v30[0] = -128
                            }
                            if (-1 >= -9223372036854775808) {
                                fmt.Println(v9, fi29)
                            } else {
                                if ((v9 + v4) != 100) {
                                    fmt.Println(v4, fi32)
                                    var v35 float32 = 1e10
                                    _ = v35
                                }
                                v36 := GenS5{F0: v2, F1: -638, F2: v4}
                                v2 -= 32767
                                _ = v36
                            }
                        }
                        _ = fi32
                    }
                default:
                    var v37 uint32 = 2147483647
                    _ = v37
                }
                if false {
                    if len(v30) > 0 {
                        v30[0] = (fi29 >> ((18446744073709551615) & 63))
                    }
                    var v38 uint32 = 0
                    _ = v38
                }
                _ = v30
            } else {
                v1 -= v1
            }
        } else {
            v5 = v9
        }
        for fi39 := 0; fi39 < 15; fi39++ {
            v8 = "foo bar"
            v2 = v2
            switch "\n\t" {
            case "foo bar":
                v5 -= 127
            case "hello":
                if (!true) {
                    v3 += 32768
                }
                var v40 uint16 = 1
                var v41 int8 = (v9 * (v5 >> ((623) & 63)))
                _ = v40
                _ = v41
            case "":
                v42 := []int{v3, fi29, (fi39 % ((32768) | 1)), 0}
                _ = v42
            default:
                v43 := GenS1{F0: 768, F1: -646}
                _ = v43
            }
            _ = fi39
        }
        if false {
            if true {
                v8 = v8
                for fi44 := 0; fi44 < 18; fi44++ {
                    switch ("hello" + v8) {
                    case "a":
                        switch v8 {
                        case "":
                            switch fi44 {
                            case 1:
                                fmt.Println(v8)
                            case 9223372036854775807:
                                if true {
                                    var v45 uint32 = 2147483648
                                    var v46 uint32 = 2
                                    var v47 int = (fi29 / (((v3 ^ (fi29 + fi29))) | 1))
                                    _ = v45
                                    _ = v46
                                    _ = v47
                                } else {
                                    v48 := GenS2{F0: 0, F1: v2, F2: (fi44 * v3)}
                                    if false {
                                        if (v2 >= v2) {
                                            fmt.Println(v8)
                                            var v49 float64 = (0.1 * float64(float32((0.1 + 1.5))))
                                            if ((2147483647 == 65535) || true) {
                                                v50 := GenS2{F0: 1, F1: 42, F2: (fi29 % (((^(fi44))) | 1))}
                                                v51 := GenS6{F0: 1, F1: v1}
                                                _ = v50
                                                _ = v51
                                            }
                                            _ = v49
                                        } else {
                                            v52 := []int{-2147483648, -9223372036854775808}
                                            fmt.Println(v4, fi29)
                                            switch (v3 | fi29) {
                                            case -2147483648:
                                                v5 -= 23
                                            case 791:
                                                _ = len(v52)
                                            }
                                            _ = v52
                                        }
                                        fmt.Println(v5)
                                        switch fi29 {
                                        case -9223372036854775808:
                                            v53 := []int{(^(fi44)), -2147483648, fi29}
                                            _ = v53
                                        case 2147483647:
                                            v9 += (v4 * (v4 * (v4 % (((v5 + (-(v9)))) | 1))))
                                        case 596:
                                            var v54 float64 = 3.14159
                                            _ = v54
                                        default:
                                            v9 += v5
                                        }
                                    }
                                    v55 := []int{(fi44 ^ fi29), (fi44 / (((fi29 * fi44)) | 1)), fi29}
                                    if len(v55) >= 2 {
                                        v55 = v55[1:len(v55)]
                                    }
                                    _ = v48
                                    _ = v55
                                }
                            case -716:
                                var v56 uint8 = 1
                                _ = v56
                            }
                        default:
                            fmt.Println(v8, v3)
                        }
                    case "hello":
                        v57 := GenS7{F0: 311}
                        _ = v57
                    default:
                        switch v8 {
                        case "":
                            v58 := GenS4{F0: 568}
                            _ = v58
                        case "foo bar":
                            fmt.Println(v4)
                        }
                    }
                    switch v8 {
                    case "a":
                        if (-2.25 < 0.1) {
                            v59 := GenS4{F0: 10}
                            switch (fi29 / ((fi29) | 1)) {
                            case 128:
                                v60 := []int{(fi29 + (-(v3))), fi29, v3, 9223372036854775807}
                                _ = v60
                            case 2147483647:
                                var v61 int64 = 1
                                if true {
                                    switch 0 {
                                    case 127:
                                        if true {
                                            v62 := []int{128}
                                            if true {
                                                var v63 int8 = (v9 ^ v5)
                                                _ = v63
                                            }
                                            _ = v62
                                        }
                                    case -950:
                                        if false {
                                            if (!false) {
                                                v5 -= v9
                                            }
                                        }
                                    case 128:
                                        switch v8 {
                                        case "hello":
                                            var v64 string = (((v8 + v8) + (v8 + "")) + (v8 + "foo bar"))
                                            _ = v64
                                        case "":
                                            switch v3 {
                                            case 991:
                                                v65 := []int{fi29, (v3 / (((fi44 | -924)) | 1)), fi29}
                                                _ = v65
                                            case -9223372036854775808:
                                                if true {
                                                    var v66 uint16 = 480
                                                    _ = v66
                                                } else {
                                                    v2 = -657
                                                    if false {
                                                        var v67 int32 = int32(v5)
                                                        var v68 string = (v8 + v8)
                                                        var v69 int = fi44
                                                        _ = v67
                                                        _ = v68
                                                        _ = v69
                                                    }
                                                    v8 = ("hello" + v8)
                                                }
                                            }
                                        case "a":
                                            switch 32768 {
                                            case 32767:
                                                if false {
                                                    v70 := []int{v3, (^(fi29))}
                                                    if (true && ((1.5 - (-2.25 + float64(1.5))) < 3.14159)) {
                                                        fmt.Println(fi29)
                                                        if true {
                                                            v9 += (v5 * v5)
                                                            v71 := GenS7{F0: v61}
                                                            var v72 int = 128
                                                            _ = v71
                                                            _ = v72
                                                        }
                                                        v73 := GenS5{F0: (v2 * 1), F1: (v61 & -546), F2: (v4 + -128)}
                                                        _ = v73
                                                    }
                                                    if (2 == 447) {
                                                        var v74 int = (-(fi29))
                                                        _ = v74
                                                    } else {
                                                        v2 = -29
                                                        switch v8 {
                                                        case "foo bar":
                                                            var v75 uint32 = 0
                                                            _ = v75
                                                        case "hello":
                                                            if (!((v8 + v8) != ("" + v8))) {
                                                                if len(v70) >= 2 {
                                                                    v70 = v70[1:len(v70)]
                                                                }
                                                                var v76 uint8 = 241
                                                                v61 -= 0
                                                                _ = v76
                                                            }
                                                        case "\n\t":
                                                            if len(v70) >= 2 {
                                                                v70 = v70[1:len(v70)]
                                                            }
                                                        default:
                                                            if (2147483647 > 296) {
                                                                switch fi29 {
                                                                case 32768:
                                                                    var v77 uint64 = 18446744073709551615
                                                                    _ = v77
                                                                }
                                                                var v78 uint32 = 441
                                                                v1 = v1
                                                                _ = v78
                                                            } else {
                                                                var v79 int16 = (v2 * -636)
                                                                v4 += v4
                                                                _ = v79
                                                            }
                                                        }
                                                    }
                                                    _ = v70
                                                } else {
                                                    v80 := []int{0, 42, (-(fi29))}
                                                    _ = v80
                                                }
                                            case -742:
                                                if (v1 != v1) {
                                                    v81 := []int{-1, (v3 << ((18446744073709551615) & 63))}
                                                    if false {
                                                        if len(v81) >= 2 {
                                                            v81 = v81[1:len(v81)]
                                                        }
                                                        var v82 int16 = (-(v2))
                                                        var v83 uint32 = 2
                                                        _ = v82
                                                        _ = v83
                                                    }
                                                    v5 += v9
                                                    var v84 bool = (v8 < ((v8 + v8) + v8))
                                                    fmt.Println(v5, v9)
                                                    _ = v81
                                                    _ = v84
                                                } else {
                                                    v1 = -677
                                                    v85 := GenS8{F0: v9}
                                                    v1 = v1
                                                    _ = v85
                                                }
                                            case 32768:
                                                v86 := GenS3{F0: 226, F1: 0}
                                                _ = v86
                                            default:
                                                if (float32(1.5) == (float32(0.0) / -2.25)) {
                                                    switch (("" + "hello") + (v8 + v8)) {
                                                    case "foo bar":
                                                        v1 -= v1
                                                    }
                                                } else {
                                                    var v87 int64 = v61
                                                    v88 := GenS3{F0: 0, F1: 154}
                                                    var v89 string = "hello"
                                                    _ = v87
                                                    _ = v88
                                                    _ = v89
                                                }
                                                v3 = 0
                                                switch v8 {
                                                case "foo bar":
                                                    var v90 float64 = 1e10
                                                    _ = v90
                                                }
                                            }
                                        default:
                                            switch "" {
                                            case "a":
                                                switch (fi29 << ((1) & 63)) {
                                                case 127:
                                                    fmt.Println(v9, v3, v2)
                                                case 1:
                                                    var v91 int32 = v1
                                                    _ = v91
                                                case -2147483648:
                                                    var v92 float64 = 0.0
                                                    _ = v92
                                                }
                                            }
                                        }
                                    }
                                    switch (fi44 << ((18446744073709551615) & 63)) {
                                    case -2147483648:
                                        v2 += 31
                                    }
                                    v9 -= v9
                                } else {
                                    fmt.Println(fi29, v9)
                                    var v93 uint64 = 18446744073709551615
                                    v94 := GenS6{F0: 42, F1: int32(v2)}
                                    _ = v93
                                    _ = v94
                                }
                                fmt.Println(v4)
                                _ = v61
                            case 4294967295:
                                v95 := GenS9{F0: 2147483648, F1: 2147483647, F2: 90}
                                _ = v95
                            }
                            var v96 float64 = float64(-2.25)
                            v97 := []int{-2147483648, (^(v3)), (fi44 - (fi44 - v3)), (fi44 % ((fi29) | 1))}
                            _ = v59
                            _ = v96
                            _ = v97
                        }
                    case "hello":
                        var v98 uint32 = 2147483647
                        _ = v98
                    }
                    v99 := []int{}
                    v4 -= v4
                    if len(v99) > 0 {
                        v99[0] = 4294967295
                    }
                    _ = fi44
                    _ = v99
                }
                if ((float32((1.5 * 3.14159)) >= (3.14159 * 0.0)) && ((-(v2)) >= (^(v2)))) {
                    if false {
                        v2 = (v2 ^ 32767)
                    }
                } else {
                    for fi100 := 0; fi100 < 3; fi100++ {
                        v101 := []int{(-(fi29)), v3, fi100}
                        _ = fi100
                        _ = v101
                    }
                    for fi102 := 0; fi102 < 18; fi102++ {
                        v103 := GenS6{F0: 1, F1: -459}
                        _ = fi102
                        _ = v103
                    }
                }
            } else {
                switch ((v8 + ((((v8 + (v8 + ("" + ("" + "")))) + "hello") + (v8 + v8)) + v8)) + ("" + v8)) {
                case "hello":
                    v104 := GenS8{F0: v4}
                    _ = v104
                case "a":
                    v9 -= 42
                }
            }
            for fi105 := 0; fi105 < 4; fi105++ {
                v1 -= int32(v2)
                _ = fi105
            }
        } else {
            v106 := GenS7{F0: 42}
            var v107 float64 = -2.25
            if (!true) {
                v9 += 100
                switch (^(fi29)) {
                case -414:
                    var v108 float64 = v107
                    _ = v108
                }
                var v109 string = v8
                _ = v109
            } else {
                v1 = (v1 | (v1 >> ((4294967295) & 63)))
            }
            _ = v106
            _ = v107
        }
        _ = fi29
    }
    for fi110 := 0; fi110 < 2; fi110++ {
        var v111 int32 = 783
        v112 := []int{(fi110 - (fi110 << ((2) & 63)))}
        _ = fi110
        _ = v111
        _ = v112
    }
    var v113 uint8 = 0
    for fi114 := 0; fi114 < 5; fi114++ {
        var v115 float32 = (3.14159 + (-2.25 + -2.25))
        v3 = 9223372036854775807
        _ = fi114
        _ = v115
    }
    switch 9223372036854775807 {
    case 127:
        v116 := []int{}
        _ = v116
    case 1:
        var v117 int32 = v1
        _ = v117
    }
    var v118 string = "hello"
    v1 += (v1 - (v1 << ((4294967295) & 63)))
    fmt.Println("v1:", v1)
    fmt.Println("v113:", v113)
    fmt.Println("v118:", v118)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v113
    _ = v118
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
