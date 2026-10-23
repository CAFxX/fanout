package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32; F2 string }
type GenS2 struct { F0 int32; F1 int16 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint64; F1 int32 }
type GenS5 struct { F0 uint; F1 string; F2 int64 }
type GenS6 struct { F0 uint32; F1 string; F2 bool }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 uint16; F1 bool; F2 bool }
type GenS9 struct { F0 int64; F1 int8 }
type GenS10 struct { F0 int8 }
type GenS11 struct { F0 uint8 }
type GenS12 struct { F0 bool; F1 uint32 }
type GenS13 struct { F0 int8 }
type GenS14 struct { F0 int8; F1 int64; F2 int64 }
type GenS15 struct { F0 uint }
type GenS16 struct { F0 uint; F1 int64; F2 int16 }
type GenS17 struct { F0 int32; F1 int }
type GenS18 struct { F0 uint16; F1 int32 }
type GenS19 struct { F0 int32; F1 int8 }
type GenS20 struct { F0 bool; F1 uint32 }
type GenS21 struct { F0 uint64; F1 int16; F2 uint8 }
type GenS22 struct { F0 int; F1 uint }
type GenS23 struct { F0 uint16; F1 uint32; F2 uint16 }
type GenS24 struct { F0 int32; F1 int; F2 uint8 }
type GenS25 struct { F0 uint; F1 uint16; F2 uint16 }
type GenS26 struct { F0 uint32; F1 uint8; F2 int8 }
type GenS27 struct { F0 uint }
type GenS28 struct { F0 uint32; F1 int8; F2 uint8 }
type GenS29 struct { F0 int8; F1 uint16 }
type GenS30 struct { F0 int16; F1 uint64; F2 int64 }
func genFn0() uint32 {
    var v7 bool = true
    if (645 != 65535) {
        var v8 int8 = -24
        _ = v8
    }
    if (v7 && (!(!(-9223372036854775808 < -9223372036854775808)))) {
        v7 = true
        var v9 float64 = (1.5 + (1e10 + -2.25))
        for fi10 := 0; fi10 < 14; fi10++ {
            v11 := GenS1{F0: (fi10 & (-(fi10))), F1: 538, F2: "a"}
            for fi12 := 0; fi12 < 17; fi12++ {
                v13 := GenS1{F0: (-(fi10)), F1: 1, F2: ""}
                _ = fi12
                _ = v13
            }
            _ = fi10
            _ = v11
        }
        _ = v9
    }
    _ = v7
    return 4294967295
}

func main() {
    var v1 int16 = -570
    v2 := make(map[uint32]string)
    v3 := make(map[uint32]bool)
    var v4 uint = 255
    var v5 string = "\n\t"
    var v6 uint64 = 42
    v1 = v1
    var v14 int = 688
    if (false && (255 != 514)) {
        v6 -= uint64(v4)
        for fi15 := 0; fi15 < 13; fi15++ {
            var v16 string = (v5 + v5)
            if false {
                if false {
                    v1 = int16(v4)
                    _ = len(v3)
                } else {
                    if false {
                        delete(v3, 119)
                        for fi17 := 0; fi17 < 11; fi17++ {
                            if true {
                                var v18 uint32 = 1
                                v14 = (v14 >> (((v4 | uint(v18))) & 63))
                                if (true || (!true)) {
                                    var v19 int8 = 0
                                    var v20 bool = true
                                    fmt.Println(v18, fi17, v19)
                                    _ = v19
                                    _ = v20
                                }
                                _ = v18
                            } else {
                                switch ("foo bar" + (v5 + "foo bar")) {
                                case "":
                                    if true {
                                        fmt.Println(v14, fi15, v6)
                                        if (((1e10 + float64((0.1 / ((1.5 - -2.25) + -2.25)))) / (float64((-2.25 - 0.0)) - float64((0.0 * ((float32(0.1) + (((((1e10 + 1.5) * (float32(float64(0.0)) - (float32((0.0 * 0.0)) * (float32(3.14159) / (1.5 + 1.5))))) + (-2.25 / 1.0)) - 1.5) * 0.1)) / 1.0))))) <= (3.14159 * float64(((((0.0 / (3.14159 + 1.5)) - (1.5 / 0.1)) / float32(((1e10 - (((0.1 * float64(1.5)) - 1.5) * 0.1)) / float64(-2.25)))) * float32(3.14159))))) {
                                            if (v1 > v1) {
                                                fmt.Println(v6, fi17, v4)
                                                v21 := GenS1{F0: fi15, F1: genFn0(), F2: v5}
                                                v22 := v2[1]
                                                _ = v21
                                                _ = v22
                                            }
                                            v1 = v1
                                            var v23 int8 = 93
                                            _ = v23
                                        }
                                        fmt.Println(fi15, v1)
                                        var v24 float32 = 3.14159
                                        switch "\n\t" {
                                        case "a":
                                            var v25 uint16 = 32767
                                            _ = v25
                                        case "\n\t":
                                            v14 = fi15
                                        case "":
                                            switch fi15 {
                                            case -1:
                                                v26 := GenS2{F0: 42, F1: (v1 << ((v4) & 63))}
                                                _ = v26
                                            case -156:
                                                v4 = (v4 - 2)
                                            default:
                                                var v27 uint16 = 484
                                                _ = v27
                                            }
                                            if (2 > 163) {
                                                switch (v16 + v16) {
                                                case "a":
                                                    var v28 uint64 = 42
                                                    _ = v28
                                                }
                                            }
                                            var v29 int = fi17
                                            _ = v29
                                        }
                                        _ = v24
                                    }
                                case "hello":
                                    if (-1 != -9223372036854775808) {
                                        fmt.Println(v5, v1, v16)
                                        if (2 <= 226) {
                                            switch v14 {
                                            case 0:
                                                var v30 uint16 = 0
                                                _ = v30
                                            case 9223372036854775807:
                                                switch (fi15 << ((v4) & 63)) {
                                                case -2147483648:
                                                    v3[65535] = false
                                                default:
                                                    var v31 string = ("\n\t" + v16)
                                                    _ = v31
                                                }
                                                var v32 string = (v5 + v5)
                                                _ = v32
                                            case 1:
                                                switch v14 {
                                                case -2147483648:
                                                    fmt.Println(v16)
                                                case -9223372036854775808:
                                                    var v33 uint32 = 2147483647
                                                    _ = v33
                                                case 771:
                                                    v16 = "a"
                                                }
                                            default:
                                                v34 := GenS1{F0: (fi15 ^ fi17), F1: 2147483647, F2: v5}
                                                _ = v34
                                            }
                                            if (1 > 146) {
                                                switch -2147483648 {
                                                case 0:
                                                    if false {
                                                        if (v1 == 42) {
                                                            v35 := GenS3{F0: (v6 + v6)}
                                                            v36 := GenS4{F0: v6, F1: 0}
                                                            var v37 uint8 = 101
                                                            _ = v35
                                                            _ = v36
                                                            _ = v37
                                                        } else {
                                                            v38 := GenS3{F0: 4294967295}
                                                            v2[65535] = v16
                                                            _ = len(v3)
                                                            _ = v38
                                                        }
                                                        fmt.Println(v6, v5)
                                                        v39 := GenS2{F0: 510, F1: v1}
                                                        _ = v39
                                                    }
                                                case 4294967295:
                                                    v14 = (fi15 - (-(v14)))
                                                    if false {
                                                        if (4294967295 >= 4294967295) {
                                                            var v40 uint16 = 707
                                                            var v41 float32 = ((3.14159 - -2.25) / (float32(-2.25) * float32(1.5)))
                                                            _ = v40
                                                            _ = v41
                                                        }
                                                        if true {
                                                            var v42 int32 = 1
                                                            _ = v42
                                                        }
                                                        if false {
                                                            _ = len(v3)
                                                            v43 := []int{v14, fi17}
                                                            fmt.Println(v5, v5)
                                                            _ = v43
                                                        }
                                                        if false {
                                                            fmt.Println(v16, v16, v5)
                                                            var v44 uint8 = 1
                                                            var v45 int32 = 0
                                                            _ = v44
                                                            _ = v45
                                                        } else {
                                                            v46 := GenS2{F0: -1, F1: (v1 ^ 812)}
                                                            v16 = v16
                                                            if ((("foo bar" + v5) + v16) > (v16 + (v5 + ""))) {
                                                                if false {
                                                                    switch fi15 {
                                                                    case 0:
                                                                        fmt.Println(fi15)
                                                                    case 42:
                                                                        v5 = (((v16 + "") + "a") + (("hello" + "foo bar") + "hello"))
                                                                    case -9223372036854775808:
                                                                        v6 = (v6 + uint64(v4))
                                                                    }
                                                                    v47 := GenS4{F0: 0, F1: 42}
                                                                    _ = v47
                                                                }
                                                                var v48 int32 = 42
                                                                _ = v48
                                                            }
                                                            _ = v46
                                                        }
                                                        switch v5 {
                                                        case "hello":
                                                            v2[65535] = ((v16 + v16) + v16)
                                                        }
                                                    } else {
                                                        fmt.Println(v14, v1, v6)
                                                        v2[731] = "a"
                                                        fmt.Println(v1, v14)
                                                        if false {
                                                            v14 -= 2147483647
                                                            var v49 int = 32767
                                                            var v50 int16 = v1
                                                            _ = v49
                                                            _ = v50
                                                        } else {
                                                            var v51 bool = (128 <= 242)
                                                            v5 = v16
                                                            _ = v51
                                                        }
                                                    }
                                                case 32768:
                                                    if ((true && true) && ((v1 * 32767) == (-(v1)))) {
                                                        if true {
                                                            v52 := GenS5{F0: (^(v4)), F1: "foo bar", F2: -9223372036854775808}
                                                            var v53 bool = false
                                                            if (v16 < v5) {
                                                                fmt.Println(fi15, v4, v4)
                                                                v54 := []int{fi15, 127, (v14 ^ 32768)}
                                                                _ = v54
                                                            }
                                                            _ = v52
                                                            _ = v53
                                                        } else {
                                                            v5 = "a"
                                                            if false {
                                                                if (false && true) {
                                                                    switch 32767 {
                                                                    case -471:
                                                                        v55 := GenS6{F0: 4294967295, F1: v5, F2: (true || (false || false))}
                                                                        _ = v55
                                                                    case 682:
                                                                        var v56 string = v5
                                                                        _ = v56
                                                                    default:
                                                                        var v57 uint64 = (v6 % ((18446744073709551615) | 1))
                                                                        _ = v57
                                                                    }
                                                                    v58 := GenS7{F0: 1}
                                                                    var v59 string = v5
                                                                    _ = v58
                                                                    _ = v59
                                                                } else {
                                                                    var v60 uint32 = 2147483648
                                                                    var v61 int = v14
                                                                    _ = v60
                                                                    _ = v61
                                                                }
                                                                var v62 float32 = ((float32(3.14159) / 1.0) * (0.0 - ((1.5 + 1e10) + 3.14159)))
                                                                _ = v62
                                                            }
                                                        }
                                                        v5 = "hello"
                                                    } else {
                                                        var v63 uint64 = 395
                                                        if (-2147483648 > -1) {
                                                            if true {
                                                                switch "a" {
                                                                case "":
                                                                    if (false || true) {
                                                                        var v64 float64 = 3.14159
                                                                        var v65 int8 = 1
                                                                        v66 := GenS2{F0: 0, F1: -515}
                                                                        _ = v64
                                                                        _ = v65
                                                                        _ = v66
                                                                    } else {
                                                                        switch fi17 {
                                                                        case 32767:
                                                                            v4 += v4
                                                                        default:
                                                                            switch "foo bar" {
                                                                            case "hello":
                                                                                v67 := GenS4{F0: (^(v63)), F1: 500}
                                                                                _ = v67
                                                                            }
                                                                        }
                                                                    }
                                                                    if (!(!(2147483647 < 211))) {
                                                                        if false {
                                                                            if (1 != -1) {
                                                                                var v68 uint16 = 0
                                                                                v3[820] = (v4 <= (v4 + (v4 << ((18446744073709551615) & 63))))
                                                                                _ = v68
                                                                            } else {
                                                                                var v69 string = "\n\t"
                                                                                var v70 float64 = 0.0
                                                                                _ = v69
                                                                                _ = v70
                                                                            }
                                                                            v16 = "a"
                                                                        } else {
                                                                            var v71 int32 = 1
                                                                            _ = v71
                                                                        }
                                                                        if false {
                                                                            if (-1 <= 1) {
                                                                                v72 := GenS1{F0: fi15, F1: 2, F2: ("foo bar" + v5)}
                                                                                _ = v72
                                                                            }
                                                                            fmt.Println(v6, v6)
                                                                        }
                                                                        if false {
                                                                            _ = len(v3)
                                                                            if false {
                                                                                v73 := []int{}
                                                                                _ = v73
                                                                            } else {
                                                                                switch v14 {
                                                                                case -712:
                                                                                    _ = len(v3)
                                                                                case -9223372036854775808:
                                                                                    switch (v14 & v14) {
                                                                                    case 1:
                                                                                        var v74 int8 = -1
                                                                                        _ = v74
                                                                                    case 127:
                                                                                        v16 = ""
                                                                                    case -2147483648:
                                                                                        v4 -= 4294967295
                                                                                    }
                                                                                }
                                                                                _ = len(v3)
                                                                            }
                                                                        }
                                                                    }
                                                                    v75 := GenS8{F0: 2, F1: (!true), F2: (42 != 42)}
                                                                    _ = v75
                                                                case "a":
                                                                    var v76 uint16 = 804
                                                                    _ = v76
                                                                default:
                                                                    var v77 int32 = 42
                                                                    _ = v77
                                                                }
                                                            }
                                                        }
                                                        _ = v63
                                                    }
                                                default:
                                                    var v78 bool = (1.5 < 1.5)
                                                    _ = v78
                                                }
                                                v1 += -32768
                                            }
                                            switch (v5 + "hello") {
                                            case "a":
                                                v79 := GenS6{F0: 65535, F1: "a", F2: false}
                                                _ = v79
                                            case "foo bar":
                                                v80 := GenS8{F0: 968, F1: ((!((v4 | v4) < (v4 % ((38) | 1)))) && (!true)), F2: true}
                                                _ = v80
                                            default:
                                                if true {
                                                    fmt.Println(v6, v4)
                                                }
                                            }
                                        } else {
                                            if true {
                                                var v81 int16 = 854
                                                var v82 float64 = (-2.25 + (1e10 * ((((3.14159 * (1e10 * 0.0)) + (3.14159 / 1.5)) / 0.1) - 1.5)))
                                                _ = v81
                                                _ = v82
                                            }
                                        }
                                    } else {
                                        v83 := GenS9{F0: 42, F1: -104}
                                        v84 := GenS4{F0: v6, F1: 1}
                                        _ = v83
                                        _ = v84
                                    }
                                case "a":
                                    var v85 float64 = (0.0 + float64(0.1))
                                    _ = v85
                                default:
                                    var v86 int16 = v1
                                    _ = v86
                                }
                                if false {
                                    v87 := GenS10{F0: 1}
                                    _ = v87
                                }
                                switch ((((v16 + (v16 + v5)) + v16) + "") + (v16 + "a")) {
                                case "hello":
                                    switch ((v5 + ("\n\t" + v5)) + "foo bar") {
                                    case "foo bar":
                                        if true {
                                            v88 := GenS3{F0: (v6 >> (((v4 * (v4 % ((248) | 1)))) & 63))}
                                            fmt.Println(v16, v4)
                                            if true {
                                                if true {
                                                    v14 -= 1
                                                    var v89 int16 = int16(v6)
                                                    _ = v89
                                                } else {
                                                    switch v5 {
                                                    case "hello":
                                                        switch fi15 {
                                                        case 9223372036854775807:
                                                            v90 := GenS11{F0: 127}
                                                            _ = v90
                                                        default:
                                                            switch (-(v14)) {
                                                            case 1:
                                                                v5 = (v5 + (("hello" + v16) + v5))
                                                            case 0:
                                                                fmt.Println(v16)
                                                                v4 = v4
                                                                if (v16 < "hello") {
                                                                    fmt.Println(v4)
                                                                    if false {
                                                                        if true {
                                                                            switch 127 {
                                                                            case 127:
                                                                                if true {
                                                                                    switch (v5 + (v5 + (v16 + v16))) {
                                                                                    case "hello":
                                                                                        fmt.Println(v5)
                                                                                    case "":
                                                                                        var v91 float64 = float64(0.1)
                                                                                        _ = v91
                                                                                    default:
                                                                                        var v92 uint32 = genFn0()
                                                                                        _ = v92
                                                                                    }
                                                                                    v93 := GenS7{F0: -128}
                                                                                    _ = v93
                                                                                } else {
                                                                                    if (1 != 9223372036854775807) {
                                                                                        if true {
                                                                                            if (1 < 2) {
                                                                                                if (1 <= (v1 ^ (v1 + 911))) {
                                                                                                    v94 := GenS4{F0: v6, F1: -2147483648}
                                                                                                    v95 := GenS12{F0: false, F1: 65535}
                                                                                                    if ((^(v4)) != v4) {
                                                                                                        var v96 int = (-(v14))
                                                                                                        v97 := []int{}
                                                                                                        var v98 uint8 = 0
                                                                                                        var v99 int32 = 0
                                                                                                        var v100 bool = false
                                                                                                        _ = v100
                                                                                                        _ = v96
                                                                                                        _ = v97
                                                                                                        _ = v98
                                                                                                        _ = v99
                                                                                                    }
                                                                                                    _ = v94
                                                                                                    _ = v95
                                                                                                } else {
                                                                                                    _ = len(v3)
                                                                                                    v101 := v3[876]
                                                                                                    switch -182 {
                                                                                                    case -2147483648:
                                                                                                        switch 339 {
                                                                                                        case -456:
                                                                                                            v102 := GenS10{F0: 1}
                                                                                                            _ = v102
                                                                                                        case 1:
                                                                                                            var v103 uint8 = 128
                                                                                                            _ = v103
                                                                                                        default:
                                                                                                            v101 = (false && true)
                                                                                                        }
                                                                                                    case 4294967295:
                                                                                                        fmt.Println(v6)
                                                                                                    }
                                                                                                    _ = v101
                                                                                                }
                                                                                                if false {
                                                                                                    v16 = v16
                                                                                                } else {
                                                                                                    if true {
                                                                                                        v4 = (^(v4))
                                                                                                        v104 := GenS1{F0: (fi15 | fi15), F1: 2, F2: v16}
                                                                                                        _ = v104
                                                                                                    }
                                                                                                    if (2147483648 <= 1) {
                                                                                                        v14 += (^(fi17))
                                                                                                    }
                                                                                                }
                                                                                                if true {
                                                                                                    v5 = v5
                                                                                                    v105 := v3[2147483647]
                                                                                                    _ = v105
                                                                                                }
                                                                                            } else {
                                                                                                if false {
                                                                                                    v106 := []int{(fi15 | 32767), fi17, (fi15 | fi17), -861}
                                                                                                    _ = v106
                                                                                                }
                                                                                            }
                                                                                            switch (fi15 << ((18446744073709551615) & 63)) {
                                                                                            case -831:
                                                                                                var v107 int64 = 9223372036854775807
                                                                                                _ = v107
                                                                                            default:
                                                                                                v1 += (v1 << (((v4 ^ 18446744073709551615)) & 63))
                                                                                            }
                                                                                            if true {
                                                                                                switch (((((v16 + "\n\t") + v5) + v5) + (v16 + v5)) + (v16 + "a")) {
                                                                                                case "foo bar":
                                                                                                    v108 := GenS12{F0: (false && false), F1: 2147483647}
                                                                                                    _ = v108
                                                                                                case "hello":
                                                                                                    v109 := []int{fi17, v14}
                                                                                                    _ = v109
                                                                                                }
                                                                                                v6 += v6
                                                                                                v110 := GenS13{F0: 1}
                                                                                                fmt.Println(v16, v16, v5)
                                                                                                v4 = (v4 - v4)
                                                                                                _ = v110
                                                                                            }
                                                                                        }
                                                                                        v5 = "hello"
                                                                                        v6 = 1
                                                                                        fmt.Println(fi15, v1)
                                                                                    } else {
                                                                                        if false {
                                                                                            var v111 uint8 = 216
                                                                                            var v112 uint8 = 127
                                                                                            _ = v111
                                                                                            _ = v112
                                                                                        }
                                                                                        if (1e10 < 0.1) {
                                                                                            v113 := GenS10{F0: 123}
                                                                                            _ = v113
                                                                                        }
                                                                                        v4 = v4
                                                                                    }
                                                                                    var v114 float32 = 3.14159
                                                                                    _ = v114
                                                                                }
                                                                            case 4294967295:
                                                                                v14 = (fi15 | 128)
                                                                                v4 = v4
                                                                            case -9223372036854775808:
                                                                                v1 = -1
                                                                            default:
                                                                                if (!(((((float32(0.1) / -2.25) * 0.0) <= (float32(0.1) * (-2.25 - float32(0.0)))) && false) && ((false && false) && true))) {
                                                                                    fmt.Println(v16)
                                                                                    var v115 int = fi15
                                                                                    _ = v115
                                                                                }
                                                                            }
                                                                        } else {
                                                                            v5 = v16
                                                                            var v116 uint16 = 0
                                                                            v117 := []int{}
                                                                            v16 = "foo bar"
                                                                            if len(v117) >= 2 {
                                                                                v117 = v117[1:len(v117)]
                                                                            }
                                                                            var v118 float32 = (0.1 - 3.14159)
                                                                            v119 := v3[623]
                                                                            v118 = (float32(1.5) / -2.25)
                                                                            _ = v116
                                                                            _ = v117
                                                                            _ = v118
                                                                            _ = v119
                                                                        }
                                                                        var v120 bool = false
                                                                        fmt.Println(v120)
                                                                        v6 = v6
                                                                        v14 += fi17
                                                                        _ = v120
                                                                    }
                                                                    var v121 uint = (v4 ^ v4)
                                                                    _ = v121
                                                                }
                                                                var v122 int16 = (v1 * v1)
                                                                delete(v3, 0)
                                                                _ = v122
                                                            case 42:
                                                                var v123 float64 = (-2.25 + 0.1)
                                                                _ = v123
                                                            }
                                                        }
                                                    case "":
                                                        v124 := GenS14{F0: -111, F1: -9223372036854775808, F2: int64(v6)}
                                                        _ = v124
                                                    case "\n\t":
                                                        fmt.Println(v6, v5)
                                                    default:
                                                        if false {
                                                            v125 := []int{v14}
                                                            if (!false) {
                                                                if ((v1 << (((^(v4))) & 63)) <= -680) {
                                                                    v4 -= v4
                                                                    v5 = (v16 + v16)
                                                                } else {
                                                                    v126 := GenS15{F0: v4}
                                                                    fmt.Println(v4, v4)
                                                                    _ = v126
                                                                }
                                                                switch v16 {
                                                                case "a":
                                                                    v4 -= 817
                                                                case "\n\t":
                                                                    v127 := GenS8{F0: 65535, F1: false, F2: true}
                                                                    _ = v127
                                                                }
                                                            }
                                                            var v128 string = v5
                                                            _ = v125
                                                            _ = v128
                                                        }
                                                    }
                                                }
                                                v129 := GenS13{F0: 42}
                                                var v130 int = fi17
                                                _ = v129
                                                _ = v130
                                            }
                                            switch "hello" {
                                            case "":
                                                if (float32(1.5) <= 0.1) {
                                                    v5 = v16
                                                    if (!false) {
                                                        var v131 int64 = -1
                                                        var v132 uint16 = 1
                                                        v14 -= 127
                                                        _ = v131
                                                        _ = v132
                                                    } else {
                                                        if (!true) {
                                                            v133 := GenS14{F0: 1, F1: 320, F2: -709}
                                                            v16 = "hello"
                                                            _ = v133
                                                        }
                                                        v5 = (v16 + (("foo bar" + (((v16 + "hello") + "") + "\n\t")) + "\n\t"))
                                                        v4 -= (v4 / (((v4 & (v4 + v4))) | 1))
                                                        if (true || false) {
                                                            v6 = v6
                                                        } else {
                                                            v1 = 1000
                                                            var v134 float64 = 3.14159
                                                            _ = v134
                                                        }
                                                        if false {
                                                            fmt.Println(v1, v14, v14)
                                                            v135 := []int{}
                                                            _ = v135
                                                        }
                                                        if false {
                                                            var v136 int64 = -930
                                                            _ = v136
                                                        } else {
                                                            fmt.Println(v6)
                                                            v137 := GenS16{F0: v4, F1: 9223372036854775807, F2: v1}
                                                            switch v5 {
                                                            case "a":
                                                                var v138 uint32 = 2147483647
                                                                _ = v138
                                                            default:
                                                                switch (^(fi15)) {
                                                                case -628:
                                                                    var v139 uint = (v4 ^ 1)
                                                                    _ = v139
                                                                case 9223372036854775807:
                                                                    _ = len(v2)
                                                                case -515:
                                                                    fmt.Println(v6, v16, fi17)
                                                                }
                                                            }
                                                            _ = v137
                                                        }
                                                        if false {
                                                            v140 := []int{fi17, v14, (fi15 % ((fi15) | 1))}
                                                            _ = v140
                                                        } else {
                                                            if (!(1 == (v6 << ((v4) & 63)))) {
                                                                var v141 float64 = 3.14159
                                                                _ = v141
                                                            } else {
                                                                var v142 uint64 = 42
                                                                if ((128 < 128) && false) {
                                                                    v143 := GenS3{F0: v6}
                                                                    var v144 uint32 = 160
                                                                    v145 := v2[v144]
                                                                    _ = v143
                                                                    _ = v144
                                                                    _ = v145
                                                                } else {
                                                                    var v146 int16 = (v1 + v1)
                                                                    v6 += v142
                                                                    var v147 string = "\n\t"
                                                                    _ = v146
                                                                    _ = v147
                                                                }
                                                                _ = v142
                                                            }
                                                            switch -9223372036854775808 {
                                                            case -1:
                                                                if true {
                                                                    var v148 int64 = 1
                                                                    var v149 float32 = 1.5
                                                                    if (false || (fi17 < -9223372036854775808)) {
                                                                        var v150 uint8 = 0
                                                                        switch fi15 {
                                                                        case 0:
                                                                            fmt.Println(v14, v4, v14)
                                                                        case 9223372036854775807:
                                                                            v151 := GenS17{F0: -1, F1: (-(v14))}
                                                                            _ = v151
                                                                        case -1:
                                                                            v4 += (v4 / ((65535) | 1))
                                                                        default:
                                                                            if (true || (((v6 >> (((v4 * v4)) & 63)) <= (v6 | (v6 << (((v4 % ((v4) | 1))) & 63)))) && ((0.1 - float64((0.1 * 3.14159))) != 3.14159))) {
                                                                                var v152 string = (v16 + v5)
                                                                                v153 := GenS18{F0: uint16(v4), F1: -2147483648}
                                                                                _ = v152
                                                                                _ = v153
                                                                            }
                                                                        }
                                                                        if true {
                                                                            v2[531] = "a"
                                                                            v150 += (v150 - 183)
                                                                            fmt.Println(v14, v149)
                                                                        }
                                                                        switch v14 {
                                                                        case 2147483647:
                                                                            var v154 uint8 = v150
                                                                            _ = v154
                                                                        case 408:
                                                                            var v155 uint32 = 1
                                                                            _ = v155
                                                                        case 32767:
                                                                            switch "hello" {
                                                                            case "\n\t":
                                                                                var v156 float32 = v149
                                                                                _ = v156
                                                                            case "":
                                                                                v6 -= v6
                                                                            case "foo bar":
                                                                                v4 = v4
                                                                            default:
                                                                                v157 := []int{(fi15 % ((42) | 1)), fi17, 128, 9223372036854775807}
                                                                                _ = v157
                                                                            }
                                                                        default:
                                                                            v158 := GenS19{F0: 42, F1: int8(fi17)}
                                                                            delete(v3, 4294967295)
                                                                            var v159 uint = (v4 / (((v4 >> (((v4 | v4)) & 63))) | 1))
                                                                            _ = v158
                                                                            _ = v159
                                                                        }
                                                                        var v160 float32 = (1.5 - v149)
                                                                        _ = v150
                                                                        _ = v160
                                                                    } else {
                                                                        if false {
                                                                            var v161 uint16 = 65535
                                                                            v4 -= v4
                                                                            _ = v161
                                                                        }
                                                                        var v162 string = v5
                                                                        _ = v162
                                                                    }
                                                                    _ = v148
                                                                    _ = v149
                                                                }
                                                            case 42:
                                                                switch v14 {
                                                                case 569:
                                                                    v16 = "foo bar"
                                                                case 1:
                                                                    v163 := []int{-571, (fi17 >> ((v4) & 63)), 654, (fi17 << ((255) & 63))}
                                                                    _ = v163
                                                                case 32767:
                                                                    var v164 string = v16
                                                                    _ = v164
                                                                default:
                                                                    var v165 int = -1
                                                                    _ = v165
                                                                }
                                                            case 2147483647:
                                                                v166 := GenS11{F0: 212}
                                                                _ = v166
                                                            }
                                                            var v167 float64 = (((1e10 - (3.14159 / 1.5)) + 1.5) * 1.5)
                                                            fmt.Println(fi17, v1, v167)
                                                            v168 := v3[2147483647]
                                                            v167 = v167
                                                            v169 := GenS10{F0: 100}
                                                            var v170 float32 = ((-2.25 + (float32(1.5) - 1.5)) * float32(0.1))
                                                            _ = v167
                                                            _ = v168
                                                            _ = v169
                                                            _ = v170
                                                        }
                                                        var v171 float32 = 1.5
                                                        _ = v171
                                                    }
                                                    v172 := []int{fi17, fi17, fi17, v14}
                                                    _ = v172
                                                } else {
                                                    v5 = v16
                                                    var v173 string = (v16 + ("\n\t" + (v16 + v16)))
                                                    _ = v173
                                                }
                                            case "hello":
                                                var v174 uint32 = 2
                                                _ = v174
                                            }
                                            _ = v88
                                        } else {
                                            var v175 uint16 = 42
                                            v3[463] = true
                                            v176 := []int{}
                                            if true {
                                                fmt.Println(v1)
                                                v177 := GenS7{F0: 0}
                                                if (v1 <= v1) {
                                                    v178 := GenS20{F0: (false && (false || false)), F1: 277}
                                                    _ = v178
                                                } else {
                                                    v16 = (v5 + v5)
                                                    var v179 int32 = 786
                                                    switch "hello" {
                                                    case "foo bar":
                                                        switch v16 {
                                                        case "\n\t":
                                                            v180 := v3[2]
                                                            _ = v180
                                                        case "hello":
                                                            v181 := GenS6{F0: 1, F1: v16, F2: true}
                                                            _ = v181
                                                        default:
                                                            fmt.Println(v5)
                                                        }
                                                    case "\n\t":
                                                        v16 = v5
                                                    }
                                                    _ = v179
                                                }
                                                if (1 > 1) {
                                                    v182 := GenS10{F0: 53}
                                                    v183 := GenS21{F0: (v6 ^ (v6 << ((v4) & 63))), F1: v1, F2: 82}
                                                    switch fi17 {
                                                    case 42:
                                                        v16 = (v16 + "hello")
                                                        v176 = append(v176, fi15)
                                                        v1 = v1
                                                    case 32768:
                                                        v184 := GenS20{F0: false, F1: 2147483648}
                                                        _ = v184
                                                    case -481:
                                                        var v185 uint16 = v175
                                                        _ = v185
                                                    }
                                                    _ = v182
                                                    _ = v183
                                                }
                                                _ = v177
                                            } else {
                                                var v186 float32 = 0.1
                                                if len(v176) >= 2 {
                                                    v176 = v176[1:len(v176)]
                                                }
                                                if len(v176) >= 2 {
                                                    v176 = v176[1:len(v176)]
                                                }
                                                switch (v16 + "a") {
                                                case "foo bar":
                                                    if (42 == v175) {
                                                        v187 := GenS15{F0: v4}
                                                        if ((false && (!false)) || (718 > -533)) {
                                                            if (9223372036854775807 <= 1) {
                                                                v188 := GenS6{F0: 373, F1: (v5 + v5), F2: (true && (false || false))}
                                                                v3[94] = false
                                                                _ = v188
                                                            } else {
                                                                var v189 int8 = 1
                                                                _ = v189
                                                            }
                                                            switch (("a" + (((v16 + "") + v5) + v16)) + "hello") {
                                                            case "foo bar":
                                                                v190 := GenS11{F0: 185}
                                                                _ = v190
                                                            case "hello":
                                                                fmt.Println(v5)
                                                            case "":
                                                                v1 -= (-(v1))
                                                                v191 := v2[2]
                                                                _ = v191
                                                            }
                                                        }
                                                        _ = v187
                                                    } else {
                                                        var v192 uint16 = (^(v175))
                                                        _ = v192
                                                    }
                                                case "\n\t":
                                                    if (true || false) {
                                                        var v193 uint8 = 127
                                                        if (42 > 0) {
                                                            v194 := GenS22{F0: (fi17 * int(v4)), F1: 604}
                                                            v195 := GenS23{F0: v175, F1: 4294967295, F2: (^(v175))}
                                                            _ = v194
                                                            _ = v195
                                                        }
                                                        v186 -= (v186 / ((v186 * v186) + (0.1 / v186)))
                                                        _ = v193
                                                    } else {
                                                        v196 := GenS9{F0: int64(v6), F1: 57}
                                                        if ((genFn0() <= genFn0()) || false) {
                                                            fmt.Println(v1, v175, fi15)
                                                            var v197 uint16 = 1
                                                            _ = v197
                                                        }
                                                        _ = v196
                                                    }
                                                case "a":
                                                    fmt.Println(fi15)
                                                }
                                                _ = v186
                                            }
                                            if true {
                                                fmt.Println(v16)
                                                v14 = fi17
                                                v5 = v5
                                            } else {
                                                v198 := GenS10{F0: 42}
                                                var v199 bool = (!false)
                                                _ = v198
                                                _ = v199
                                            }
                                            if (442 < (v6 * (v6 + v6))) {
                                                v200 := GenS24{F0: 1, F1: v14, F2: 2}
                                                if (!((fi15 + (fi17 | 127)) < 128)) {
                                                    fmt.Println(v5)
                                                    v6 = 1
                                                } else {
                                                    var v201 bool = (1 > 2147483647)
                                                    v202 := v2[636]
                                                    var v203 uint16 = (v175 / ((65535) | 1))
                                                    _ = v201
                                                    _ = v202
                                                    _ = v203
                                                }
                                                v175 -= v175
                                                v14 = fi17
                                                if (v6 >= (v6 << ((42) & 63))) {
                                                    fmt.Println(v4, v5)
                                                    v6 = v6
                                                    if len(v176) > 0 {
                                                        v176[0] = (v14 * fi15)
                                                    }
                                                }
                                                var v204 float64 = 0.1
                                                if false {
                                                    fmt.Println(v14)
                                                    var v205 string = v5
                                                    _ = v205
                                                }
                                                _ = v200
                                                _ = v204
                                            }
                                            _ = v175
                                            _ = v176
                                        }
                                    case "\n\t":
                                        var v206 uint16 = 329
                                        _ = v206
                                    default:
                                        if false {
                                            var v207 uint8 = 100
                                            var v208 uint16 = 956
                                            _ = v207
                                            _ = v208
                                        }
                                    }
                                case "foo bar":
                                    var v209 int8 = 127
                                    _ = v209
                                case "a":
                                    if (v6 <= (v6 ^ 18446744073709551615)) {
                                        if (!(1 < (v1 - -1))) {
                                            v6 = v6
                                        }
                                        v16 = ""
                                    }
                                default:
                                    if false {
                                        var v210 uint8 = 243
                                        v14 = (v14 ^ (fi17 ^ 127))
                                        fmt.Println(fi17, v1)
                                        _ = v210
                                    }
                                }
                                v1 -= v1
                                var v211 bool = (true && false)
                                _ = v211
                            }
                            fmt.Println(fi17)
                            if true {
                                fmt.Println(v16)
                                switch "foo bar" {
                                case "hello":
                                    if (!((!(((1.5 / 1e10) / ((float64(3.14159) - -2.25) - ((float64(-2.25) - (((0.1 * -2.25) - -2.25) - (0.0 + (3.14159 + 3.14159)))) + 0.1))) <= 1.5)) || ((!((v14 | 127) >= 0)) || (!(2 != 2))))) {
                                        switch (v14 / (((fi15 >> ((v4) & 63))) | 1)) {
                                        case -1:
                                            v212 := []int{(v14 % (((^(v14))) | 1))}
                                            _ = v212
                                        default:
                                            v6 -= v6
                                        }
                                    }
                                case "\n\t":
                                    v4 = (v4 >> (((v4 << (((v4 % ((42) | 1))) & 63))) & 63))
                                }
                                v213 := v2[0]
                                _ = v213
                            } else {
                                v4 += (v4 + v4)
                                switch -1 {
                                case 4294967295:
                                    v214 := []int{fi17, fi15, (fi17 / (((fi17 << (((v4 / ((v4) | 1))) & 63))) | 1)), 2147483647}
                                    _ = v214
                                case 253:
                                    if (float64(1e10) == (3.14159 * 0.1)) {
                                        if false {
                                            var v215 uint16 = 0
                                            switch (v14 % (((fi15 - fi15)) | 1)) {
                                            case 32767:
                                                var v216 float64 = float64((-2.25 + 0.1))
                                                _ = v216
                                            default:
                                                v14 -= fi15
                                            }
                                            v1 += (-(v1))
                                            var v217 float32 = 0.0
                                            v218 := GenS18{F0: (v215 % ((169) | 1)), F1: -1}
                                            _ = v215
                                            _ = v217
                                            _ = v218
                                        }
                                        fmt.Println(fi15, v1)
                                        delete(v3, genFn0())
                                    } else {
                                        if true {
                                            var v219 uint64 = 196
                                            _ = v219
                                        } else {
                                            fmt.Println(v6)
                                            v220 := GenS9{F0: 0, F1: -50}
                                            fmt.Println(fi17, v5)
                                            _ = v220
                                        }
                                        v221 := GenS25{F0: v4, F1: 89, F2: 255}
                                        fmt.Println(v4)
                                        v222 := GenS26{F0: 42, F1: 1, F2: -1}
                                        _ = v221
                                        _ = v222
                                    }
                                }
                            }
                            _ = fi17
                        }
                        var v223 float32 = 1.5
                        _ = v223
                    }
                    if (fi15 > (fi15 ^ fi15)) {
                        var v224 int64 = 1
                        _ = v224
                    } else {
                        fmt.Println(fi15, v5)
                    }
                    for fi225 := 0; fi225 < 16; fi225++ {
                        switch 9223372036854775807 {
                        case 256:
                            fmt.Println(v5)
                        }
                        var v226 int8 = -112
                        _ = fi225
                        _ = v226
                    }
                }
                var v227 int = fi15
                if (!false) {
                    var v228 int = int(v1)
                    _ = v228
                } else {
                    v6 = (^(v6))
                    if false {
                        v5 = v5
                        var v229 uint64 = v6
                        v230 := GenS15{F0: v4}
                        _ = v229
                        _ = v230
                    }
                }
                _ = v227
            }
            var v231 int64 = -1
            _ = fi15
            _ = v16
            _ = v231
        }
    }
    v5 = (v5 + ("hello" + v5))
    v5 = "a"
    for fi232 := 0; fi232 < 4; fi232++ {
        switch (v5 + "a") {
        case "a":
            var v233 uint32 = 2147483647
            _ = v233
        case "":
            var v234 uint8 = 42
            _ = v234
        case "foo bar":
            var v235 uint64 = v6
            _ = v235
        default:
            for fi236 := 0; fi236 < 15; fi236++ {
                var v237 int16 = (v1 & (v1 << (((v4 + (^(v4)))) & 63)))
                _ = fi236
                _ = v237
            }
        }
        v238 := GenS27{F0: (^(v4))}
        switch ((((("" + "hello") + v5) + "foo bar") + ((v5 + (("hello" + (((v5 + v5) + v5) + (v5 + v5))) + v5)) + v5)) + v5) {
        case "hello":
            v239 := []int{}
            _ = v239
        case "\n\t":
            for fi240 := 0; fi240 < 3; fi240++ {
                if false {
                    v6 += (^(v6))
                }
                _ = fi240
            }
        case "a":
            switch "foo bar" {
            case "\n\t":
                v241 := []int{(fi232 % (((fi232 | (v14 / ((fi232) | 1)))) | 1))}
                _ = v241
            case "":
                v14 += (v14 / (((fi232 - (v14 | -1))) | 1))
            case "hello":
                v242 := GenS17{F0: 0, F1: (fi232 / (((v14 % (((fi232 + v14)) | 1))) | 1))}
                _ = v242
            }
        default:
            if (1 != -1) {
                v6 = uint64(v4)
                if ((!((0.1 + (1.5 - 1e10)) == 0.0)) || false) {
                    if false {
                        var v243 int32 = 399
                        _ = v243
                    }
                    v244 := GenS28{F0: 984, F1: 1, F2: 217}
                    if ((-1 == -1) || (v5 == v5)) {
                        var v245 int8 = -1
                        _ = v245
                    } else {
                        var v246 uint32 = 812
                        switch (v5 + v5) {
                        case "":
                            switch ((v5 + (v5 + v5)) + ((v5 + "") + v5)) {
                            case "":
                                v247 := GenS14{F0: 1, F1: -212, F2: -19}
                                _ = v247
                            default:
                                v6 = 42
                            }
                        case "\n\t":
                            v248 := []int{fi232}
                            _ = v248
                        default:
                            var v249 int64 = 771
                            _ = v249
                        }
                        _ = v246
                    }
                    _ = v244
                }
            } else {
                var v250 uint8 = 42
                if (v5 >= v5) {
                    var v251 uint32 = 2147483648
                    if true {
                        v4 += v4
                        var v252 int64 = 0
                        var v253 int32 = 42
                        _ = v252
                        _ = v253
                    }
                    _ = v251
                }
                _ = v250
            }
        }
        _ = fi232
        _ = v238
    }
    var v254 int8 = -48
    var v255 int64 = -1
    v256 := GenS5{F0: 149, F1: ((v5 + "hello") + v5), F2: v255}
    v257 := GenS15{F0: 540}
    switch v5 {
    case "foo bar":
        var v258 int16 = (^(v1))
        _ = v258
    }
    switch 32767 {
    case -684:
        _ = len(v2)
    default:
        v255 -= v255
    }
    v1 = v1
    switch "" {
    case "a":
        v4 = v4
    case "foo bar":
        for fi259 := 0; fi259 < 10; fi259++ {
            v260 := GenS29{F0: (v254 - v254), F1: 32767}
            var v261 uint64 = 286
            v1 = (v1 << ((v4) & 63))
            _ = fi259
            _ = v260
            _ = v261
        }
    default:
        if (!(((v1 + -135) < (v1 + (v1 >> ((v4) & 63)))) || false)) {
            v262 := GenS30{F0: v1, F1: v6, F2: (v255 * 9223372036854775807)}
            for fi263 := 0; fi263 < 2; fi263++ {
                var v264 float64 = 0.1
                var v265 uint8 = 1
                _ = fi263
                _ = v264
                _ = v265
            }
            _ = v262
        } else {
            var v266 int32 = 0
            v267 := GenS30{F0: -878, F1: (v6 % (((^(v6))) | 1)), F2: (v255 | v255)}
            _ = v266
            _ = v267
        }
    }
    var v268 uint8 = 128
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v254:", v254)
    fmt.Println("v255:", v255)
    fmt.Println("v268:", v268)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v2
    _ = v254
    _ = v255
    _ = v256
    _ = v257
    _ = v268
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
