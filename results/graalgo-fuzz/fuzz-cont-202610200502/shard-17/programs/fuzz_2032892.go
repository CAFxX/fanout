package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool; F2 int32 }
type GenS2 struct { F0 bool; F1 int8; F2 uint }
type GenS3 struct { F0 uint; F1 uint64 }
type GenS4 struct { F0 uint16; F1 int8; F2 int8 }
type GenS5 struct { F0 uint16; F1 uint8; F2 int8 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 bool; F1 string }
type GenS8 struct { F0 int64; F1 int; F2 uint64 }
type GenS9 struct { F0 uint64 }
type GenS10 struct { F0 uint8 }
type GenS11 struct { F0 uint8; F1 bool }
type GenS12 struct { F0 uint64; F1 uint16 }
type GenS13 struct { F0 uint64; F1 string }
type GenS14 struct { F0 bool; F1 uint; F2 int16 }
type GenS15 struct { F0 uint8; F1 uint32 }
type GenS16 struct { F0 string; F1 uint8 }
type GenS17 struct { F0 bool; F1 bool }
type GenS18 struct { F0 int32; F1 uint32 }
type GenS19 struct { F0 uint32; F1 uint16 }
type GenS20 struct { F0 uint8; F1 int32; F2 uint16 }

func main() {
    var v1 int32 = -1
    v2 := make(map[int]uint32)
    var v3 uint16 = 143
    v4 := make(map[int]int)
    var v5 uint64 = 42
    var v6 uint64 = v5
    v4[1] = -286
    v6 += v6
    if (374 == -709) {
        v3 -= v3
        v7 := v4[-2147483648]
        _ = v7
    } else {
        v6 -= v5
        v1 = v1
        var v8 int8 = 127
        _ = v8
    }
    v9 := GenS1{F0: v5, F1: (2147483647 != 128), F2: int32(v3)}
    var v10 uint8 = 255
    v11 := GenS1{F0: (v5 % (((^(v6))) | 1)), F1: true, F2: (-(v1))}
    var v12 int64 = 9223372036854775807
    if (!true) {
        if false {
            switch 557 {
            case 9223372036854775807:
                if true {
                    switch "foo bar" {
                    case "hello":
                        if (2147483647 > 2147483647) {
                            v13 := []int{-9223372036854775808, 128, -2147483648}
                            _ = v13
                        } else {
                            var v14 float32 = 1e10
                            _ = v14
                        }
                    case "a":
                        if (((1 >= 65535) || (((v5 / (((v5 + v6)) | 1)) < v6) || (float32(1e10) == (1e10 * -2.25)))) || ((((^(v1)) != v1) && ((!(403 >= 42)) || (1 == (v3 - v3)))) || ((false || false) || (191 < 325)))) {
                            if false {
                                v5 += v5
                                if ("a" < ("foo bar" + ("foo bar" + ((("\n\t" + "\n\t") + ("hello" + "a")) + "foo bar")))) {
                                    v1 = v1
                                    for fi15 := 0; fi15 < 16; fi15++ {
                                        v16 := GenS1{F0: v5, F1: (1.5 == float64(3.14159)), F2: (v1 * v1)}
                                        if false {
                                            v12 += v12
                                            v6 = uint64(v3)
                                        }
                                        _ = fi15
                                        _ = v16
                                    }
                                    v17 := GenS1{F0: v6, F1: (false && true), F2: 42}
                                    _ = v17
                                }
                            } else {
                                for fi18 := 0; fi18 < 4; fi18++ {
                                    for fi19 := 0; fi19 < 10; fi19++ {
                                        switch (("foo bar" + "foo bar") + "foo bar") {
                                        case "a":
                                            switch "a" {
                                            case "foo bar":
                                                if true {
                                                    v5 -= v6
                                                    v10 = v10
                                                } else {
                                                    var v20 bool = false
                                                    _ = v20
                                                }
                                            case "\n\t":
                                                var v21 string = "foo bar"
                                                if (false || false) {
                                                    if true {
                                                        if (v6 == (v5 ^ v5)) {
                                                            var v22 float32 = float32((-2.25 / 1.0))
                                                            _ = v22
                                                        } else {
                                                            v6 += (v5 ^ (^(v5)))
                                                            switch int(v1) {
                                                            case 4294967295:
                                                                v23 := GenS1{F0: v6, F1: ((v10 | uint8(v5)) > v10), F2: v1}
                                                                _ = v23
                                                            case 2147483647:
                                                                switch (fi18 | (fi18 | fi19)) {
                                                                case 9223372036854775807:
                                                                    v24 := GenS2{F0: (false || (int8(v6) > 1)), F1: -1, F2: 4294967295}
                                                                    _ = v24
                                                                case 42:
                                                                    v21 = "a"
                                                                }
                                                            }
                                                            v10 = uint8(v6)
                                                        }
                                                        _ = len(v4)
                                                        var v25 uint8 = (v10 / ((2) | 1))
                                                        switch 2147483647 {
                                                        case 4294967295:
                                                            switch "a" {
                                                            case "\n\t":
                                                                if (!(!(v3 <= (^(v3))))) {
                                                                    if (((true || (!true)) && ("a" >= v21)) || (18446744073709551615 < 65535)) {
                                                                        fmt.Println(fi19, v12)
                                                                        if false {
                                                                            v1 = (v1 & (v1 * v1))
                                                                            var v26 uint64 = v6
                                                                            _ = v26
                                                                        } else {
                                                                            v27 := GenS2{F0: true, F1: 59, F2: 0}
                                                                            var v28 uint64 = v5
                                                                            fmt.Println(v12, fi19, v21)
                                                                            _ = v27
                                                                            _ = v28
                                                                        }
                                                                        var v29 float32 = 1e10
                                                                        _ = v29
                                                                    } else {
                                                                        v12 = v12
                                                                        fmt.Println(v1, v25)
                                                                    }
                                                                    fmt.Println(v1, v12)
                                                                    var v30 uint8 = (^(v10))
                                                                    _ = v30
                                                                } else {
                                                                    if ((v1 >> ((192) & 63)) >= (v1 * int32(v12))) {
                                                                        if true {
                                                                            fmt.Println(v5)
                                                                            v31 := GenS3{F0: uint(v3), F1: v5}
                                                                            _ = len(v4)
                                                                            _ = v31
                                                                        }
                                                                        v6 = v5
                                                                    } else {
                                                                        fmt.Println(fi18, v21, v6)
                                                                        v32 := GenS4{F0: v3, F1: int8(fi18), F2: -20}
                                                                        v33 := GenS2{F0: false, F1: -128, F2: 2}
                                                                        _ = v32
                                                                        _ = v33
                                                                    }
                                                                    v34 := []int{}
                                                                    _ = len(v4)
                                                                    _ = v34
                                                                }
                                                                var v35 float64 = (-2.25 * 0.1)
                                                                _ = v35
                                                            case "hello":
                                                                switch 32768 {
                                                                case 0:
                                                                    var v36 uint = 4294967295
                                                                    _ = v36
                                                                case -9223372036854775808:
                                                                    v21 = (v21 + "foo bar")
                                                                case -2147483648:
                                                                    v37 := []int{fi19}
                                                                    _ = v37
                                                                default:
                                                                    switch v21 {
                                                                    case "hello":
                                                                        var v38 string = (("a" + ("\n\t" + v21)) + v21)
                                                                        v10 = (^(v10))
                                                                        fmt.Println(v5, v38, v3)
                                                                        v39 := GenS5{F0: uint16(v10), F1: (v25 % ((128) | 1)), F2: 11}
                                                                        v3 -= (v3 - v3)
                                                                        _ = v38
                                                                        _ = v39
                                                                    case "":
                                                                        if true {
                                                                            v40 := GenS5{F0: v3, F1: v10, F2: 46}
                                                                            var v41 int32 = (v1 * (v1 >> ((0) & 63)))
                                                                            fmt.Println(v6, v12)
                                                                            _ = v40
                                                                            _ = v41
                                                                        }
                                                                    }
                                                                }
                                                            }
                                                        case 128:
                                                            _ = len(v4)
                                                        case 0:
                                                            var v42 uint32 = 1
                                                            _ = v42
                                                        }
                                                        var v43 uint = 18446744073709551615
                                                        v6 -= 42
                                                        _ = v25
                                                        _ = v43
                                                    } else {
                                                        v44 := []int{fi18}
                                                        var v45 float32 = -2.25
                                                        _ = v44
                                                        _ = v45
                                                    }
                                                    v12 = (v12 / ((v12) | 1))
                                                    switch (fi19 << ((840) & 63)) {
                                                    case 127:
                                                        switch (fi19 * int(v10)) {
                                                        case 739:
                                                            if false {
                                                                v46 := GenS6{F0: (v10 / (((v10 % ((v10) | 1))) | 1))}
                                                                v47 := GenS7{F0: (!true), F1: "foo bar"}
                                                                v48 := []int{97, 0, (fi18 & fi18)}
                                                                _ = v46
                                                                _ = v47
                                                                _ = v48
                                                            } else {
                                                                var v49 bool = (((fi18 << ((496) & 63)) != 42) || false)
                                                                v6 = (^(v5))
                                                                var v50 float64 = 1.5
                                                                _ = v49
                                                                _ = v50
                                                            }
                                                        }
                                                    case -9223372036854775808:
                                                        var v51 int8 = 127
                                                        _ = v51
                                                    case -725:
                                                        var v52 string = (v21 + "a")
                                                        _ = v52
                                                    default:
                                                        _ = len(v4)
                                                    }
                                                } else {
                                                    var v53 int = fi19
                                                    var v54 int8 = 42
                                                    if false {
                                                        var v55 int64 = int64(v54)
                                                        var v56 uint = 536
                                                        _ = v55
                                                        _ = v56
                                                    } else {
                                                        fmt.Println(v5)
                                                        v57 := []int{(v53 ^ (fi18 ^ 128)), v53, v53}
                                                        _ = v57
                                                    }
                                                    _ = v53
                                                    _ = v54
                                                }
                                                _ = v21
                                            }
                                        case "foo bar":
                                            var v58 float64 = -2.25
                                            v59 := GenS8{F0: v12, F1: -9223372036854775808, F2: (v5 - (v5 >> ((uint(v3)) & 63)))}
                                            _ = v58
                                            _ = v59
                                        case "hello":
                                            v60 := []int{}
                                            _ = v60
                                        }
                                        _ = fi19
                                    }
                                    _ = fi18
                                }
                                for fi61 := 0; fi61 < 20; fi61++ {
                                    v62 := GenS9{F0: (v5 << ((18446744073709551615) & 63))}
                                    v12 = (v12 ^ (-(v12)))
                                    var v63 string = "\n\t"
                                    _ = fi61
                                    _ = v62
                                    _ = v63
                                }
                            }
                            var v64 float32 = (float32(1.5) * 0.0)
                            _ = v64
                        }
                    default:
                        var v65 uint64 = uint64(v10)
                        _ = v65
                    }
                    var v66 int8 = -127
                    _ = v66
                }
            case -2147483648:
                if (!((false || false) && (1 < v5))) {
                    if false {
                        for fi67 := 0; fi67 < 9; fi67++ {
                            v68 := GenS6{F0: (v10 | 255)}
                            _ = fi67
                            _ = v68
                        }
                        if true {
                            if true {
                                delete(v4, 128)
                                var v69 uint = 735
                                if false {
                                    v3 += (v3 + (^(v3)))
                                    if false {
                                        if true {
                                            for fi70 := 0; fi70 < 4; fi70++ {
                                                var v71 int32 = v1
                                                v71 -= (-(v1))
                                                var v72 int64 = v12
                                                _ = fi70
                                                _ = v71
                                                _ = v72
                                            }
                                            var v73 int32 = v1
                                            _ = v73
                                        } else {
                                            v74 := v2[240]
                                            _ = v74
                                        }
                                        var v75 uint8 = 1
                                        _ = v75
                                    } else {
                                        if (1 >= v10) {
                                            var v76 bool = true
                                            _ = v76
                                        }
                                        var v77 float32 = 1.5
                                        _ = v77
                                    }
                                } else {
                                    switch 0 {
                                    case 128:
                                        v10 = v10
                                    case -199:
                                        switch "" {
                                        case "foo bar":
                                            _ = len(v2)
                                        case "":
                                            switch ("a" + "foo bar") {
                                            case "a":
                                                fmt.Println(v3, v1)
                                            case "hello":
                                                v78 := GenS1{F0: 0, F1: ((v12 + int64(v5)) == (v12 - (-(v12)))), F2: (v1 | v1)}
                                                _ = v78
                                            case "":
                                                var v79 bool = (3.14159 < (0.0 / 1.0))
                                                _ = v79
                                            default:
                                                if true {
                                                    v6 = v6
                                                    delete(v2, -1)
                                                    if ("hello" < (("\n\t" + "foo bar") + ("" + "\n\t"))) {
                                                        switch -9223372036854775808 {
                                                        case -2147483648:
                                                            if (2147483647 < 65535) {
                                                                v80 := GenS10{F0: v10}
                                                                _ = v80
                                                            } else {
                                                                var v81 uint32 = 0
                                                                var v82 bool = (uint(v10) != v69)
                                                                v83 := []int{127, -279}
                                                                _ = v81
                                                                _ = v82
                                                                _ = v83
                                                            }
                                                        case 1:
                                                            v84 := []int{}
                                                            _ = v84
                                                        default:
                                                            v69 -= 196
                                                        }
                                                        var v85 uint = 662
                                                        if (!(!((-(v1)) == v1))) {
                                                            for fi86 := 0; fi86 < 4; fi86++ {
                                                                v87 := []int{630, (^(fi86)), (fi86 >> ((uint(v3)) & 63))}
                                                                if true {
                                                                    var v88 int64 = -9223372036854775808
                                                                    var v89 float64 = float64(1e10)
                                                                    _ = v88
                                                                    _ = v89
                                                                } else {
                                                                    if (false || true) {
                                                                        v6 = v6
                                                                        v90 := GenS5{F0: v3, F1: (^(v10)), F2: 11}
                                                                        _ = v90
                                                                    } else {
                                                                        for fi91 := 0; fi91 < 14; fi91++ {
                                                                            var v92 bool = (float64(0.0) <= (3.14159 + (3.14159 + 1e10)))
                                                                            var v93 uint64 = v6
                                                                            if v92 {
                                                                                delete(v4, fi86)
                                                                            }
                                                                            _ = fi91
                                                                            _ = v92
                                                                            _ = v93
                                                                        }
                                                                        v6 = (v5 / (((v5 >> ((v69) & 63))) | 1))
                                                                        v94 := v4[0]
                                                                        _ = v94
                                                                    }
                                                                    if len(v87) > 0 {
                                                                        v87[0] = fi86
                                                                    }
                                                                    if (!true) {
                                                                        var v95 int16 = 1
                                                                        if ((1000 >= v95) && false) {
                                                                            if (!(9223372036854775807 <= (-(v12)))) {
                                                                                if false {
                                                                                    if (false && ((uint8(v6) < uint8(v3)) || false)) {
                                                                                        v95 -= -145
                                                                                        fmt.Println(v12)
                                                                                        for fi96 := 0; fi96 < 3; fi96++ {
                                                                                            v85 = 0
                                                                                            _ = fi96
                                                                                        }
                                                                                    }
                                                                                } else {
                                                                                    v97 := v2[(-(fi86))]
                                                                                    _ = v97
                                                                                }
                                                                            }
                                                                        } else {
                                                                            var v98 uint32 = 42
                                                                            _ = v98
                                                                        }
                                                                        fmt.Println(v6)
                                                                        _ = v95
                                                                    } else {
                                                                        v99 := GenS11{F0: (v10 * (v10 << (((v69 | (^(v69)))) & 63))), F1: false}
                                                                        switch fi86 {
                                                                        case 9223372036854775807:
                                                                            for fi100 := 0; fi100 < 2; fi100++ {
                                                                                _ = len(v87)
                                                                                var v101 uint16 = 901
                                                                                var v102 int = (-(fi86))
                                                                                var v103 int8 = 0
                                                                                if ((!true) || (v69 <= (v69 | v85))) {
                                                                                    v104 := v4[127]
                                                                                    if false {
                                                                                        if (!(0.0 == ((1e10 - 0.0) / 1e10))) {
                                                                                            var v105 float32 = 1e10
                                                                                            v106 := GenS2{F0: (!true), F1: (v103 ^ 42), F2: v85}
                                                                                            _ = v105
                                                                                            _ = v106
                                                                                        } else {
                                                                                            switch v102 {
                                                                                            case 1:
                                                                                                fmt.Println(v10, v104, v85)
                                                                                            case 4294967295:
                                                                                                delete(v2, (fi86 | (v102 / ((v102) | 1))))
                                                                                            default:
                                                                                                if len(v87) >= 2 {
                                                                                                    v87 = v87[1:len(v87)]
                                                                                                }
                                                                                                switch "hello" {
                                                                                                case "a":
                                                                                                    var v107 uint32 = 65535
                                                                                                    _ = v107
                                                                                                default:
                                                                                                    v102 = fi100
                                                                                                }
                                                                                                var v108 int64 = 9223372036854775807
                                                                                                _ = v108
                                                                                            }
                                                                                            fmt.Println(v10)
                                                                                            switch v104 {
                                                                                            case 32768:
                                                                                                if false {
                                                                                                    if (!false) {
                                                                                                        if true {
                                                                                                            v85 -= (v85 + (v69 ^ (v85 % ((v85) | 1))))
                                                                                                            v1 -= v1
                                                                                                            var v109 uint16 = uint16(v85)
                                                                                                            _ = v109
                                                                                                        } else {
                                                                                                            var v110 bool = false
                                                                                                            _ = v110
                                                                                                        }
                                                                                                    } else {
                                                                                                        var v111 uint64 = 4294967295
                                                                                                        v111 += 269
                                                                                                        _ = v111
                                                                                                    }
                                                                                                }
                                                                                                var v112 string = "foo bar"
                                                                                                if false {
                                                                                                    v104 = (fi86 * 32768)
                                                                                                    fmt.Println(v112, v6)
                                                                                                    var v113 int = (fi100 * (v102 + (fi86 & (-(fi86)))))
                                                                                                    _ = v113
                                                                                                }
                                                                                                switch (v102 | fi100) {
                                                                                                case 464:
                                                                                                    if false {
                                                                                                        switch "hello" {
                                                                                                        case "hello":
                                                                                                            fmt.Println(v69, v1, fi86)
                                                                                                        case "\n\t":
                                                                                                            switch v104 {
                                                                                                            case 4294967295:
                                                                                                                var v114 uint64 = (v5 | (v5 >> (((v85 << ((v85) & 63))) & 63)))
                                                                                                                _ = v114
                                                                                                            default:
                                                                                                                switch (^(v102)) {
                                                                                                                case -1:
                                                                                                                    switch ("a" + "\n\t") {
                                                                                                                    case "hello":
                                                                                                                        v85 -= (v85 ^ v85)
                                                                                                                    case "foo bar":
                                                                                                                        v115 := GenS9{F0: (v5 ^ v5)}
                                                                                                                        _ = v115
                                                                                                                    case "a":
                                                                                                                        v116 := GenS2{F0: (!true), F1: (-(v103)), F2: (v69 + v85)}
                                                                                                                        v6 -= (v6 << ((uint(v3)) & 63))
                                                                                                                        _ = v116
                                                                                                                    }
                                                                                                                case 42:
                                                                                                                    var v117 uint32 = 42
                                                                                                                    v5 += v6
                                                                                                                    _ = v117
                                                                                                                }
                                                                                                            }
                                                                                                        case "a":
                                                                                                            if (false || true) {
                                                                                                                var v118 int8 = -128
                                                                                                                if ((v3 >> (((v69 << (((^(v85))) & 63))) & 63)) > uint16(v5)) {
                                                                                                                    v87 = append(v87, -2147483648)
                                                                                                                } else {
                                                                                                                    v119 := v2[fi86]
                                                                                                                    var v120 int32 = v1
                                                                                                                    v121 := GenS3{F0: 318, F1: (^(v6))}
                                                                                                                    _ = v119
                                                                                                                    _ = v120
                                                                                                                    _ = v121
                                                                                                                }
                                                                                                                var v122 bool = true
                                                                                                                fmt.Println(v102)
                                                                                                                if v122 {
                                                                                                                    if ((4 > 313) || true) {
                                                                                                                        v123 := GenS3{F0: v69, F1: uint64(v69)}
                                                                                                                        var v124 uint16 = (v101 % ((255) | 1))
                                                                                                                        if len(v87) >= 2 {
                                                                                                                            v87 = v87[1:len(v87)]
                                                                                                                        }
                                                                                                                        _ = v123
                                                                                                                        _ = v124
                                                                                                                    }
                                                                                                                    _ = len(v87)
                                                                                                                }
                                                                                                                v125 := GenS9{F0: (v6 + (v6 ^ v5))}
                                                                                                                var v126 int = fi100
                                                                                                                _ = v118
                                                                                                                _ = v122
                                                                                                                _ = v125
                                                                                                                _ = v126
                                                                                                            }
                                                                                                        default:
                                                                                                            if (false && false) {
                                                                                                                v85 = 18446744073709551615
                                                                                                            } else {
                                                                                                                v101 -= (^(v101))
                                                                                                            }
                                                                                                        }
                                                                                                        fmt.Println(v3)
                                                                                                    } else {
                                                                                                        if false {
                                                                                                            var v127 string = v112
                                                                                                            switch v127 {
                                                                                                            case "hello":
                                                                                                                v85 = v85
                                                                                                            default:
                                                                                                                var v128 bool = false
                                                                                                                _ = v128
                                                                                                            }
                                                                                                            _ = v127
                                                                                                        }
                                                                                                    }
                                                                                                case 32768:
                                                                                                    v12 = v12
                                                                                                default:
                                                                                                    fmt.Println(v10, fi86)
                                                                                                }
                                                                                                var v129 uint = (v85 << (((v69 + v85)) & 63))
                                                                                                v87 = append(v87, (^(fi86)))
                                                                                                _ = v112
                                                                                                _ = v129
                                                                                            case 128:
                                                                                                v87 = append(v87, 0)
                                                                                            case 1:
                                                                                                v130 := GenS8{F0: (v12 / (((v12 - v12)) | 1)), F1: (v104 / (((-(v104))) | 1)), F2: v5}
                                                                                                _ = v130
                                                                                            default:
                                                                                                if true {
                                                                                                    v101 += v101
                                                                                                    var v131 int32 = 0
                                                                                                    v131 = -1
                                                                                                    fmt.Println(v101, v5)
                                                                                                    _ = v131
                                                                                                } else {
                                                                                                    var v132 uint32 = 694
                                                                                                    _ = v132
                                                                                                }
                                                                                            }
                                                                                        }
                                                                                        v103 = (v103 >> ((v85) & 63))
                                                                                    }
                                                                                    _ = v104
                                                                                }
                                                                                _ = fi100
                                                                                _ = v101
                                                                                _ = v102
                                                                                _ = v103
                                                                            }
                                                                        }
                                                                        _ = v99
                                                                    }
                                                                }
                                                                _ = fi86
                                                                _ = v87
                                                            }
                                                            v133 := []int{-516, -9223372036854775808}
                                                            _ = v133
                                                        } else {
                                                            v2[9223372036854775807] = 2147483647
                                                            var v134 string = "foo bar"
                                                            v135 := GenS6{F0: (v10 & (v10 ^ v10))}
                                                            _ = v134
                                                            _ = v135
                                                        }
                                                        _ = v85
                                                    } else {
                                                        for fi136 := 0; fi136 < 5; fi136++ {
                                                            v12 = (v12 | (^(v12)))
                                                            _ = fi136
                                                        }
                                                    }
                                                }
                                            }
                                        default:
                                            for fi137 := 0; fi137 < 14; fi137++ {
                                                switch "\n\t" {
                                                case "a":
                                                    if (v69 <= (v69 * 65535)) {
                                                        v12 = v12
                                                        for fi138 := 0; fi138 < 1; fi138++ {
                                                            switch (fi137 % ((fi137) | 1)) {
                                                            case 801:
                                                                if (v3 > (^(v3))) {
                                                                    if false {
                                                                        v12 -= v12
                                                                        var v139 int8 = 100
                                                                        var v140 int8 = (v139 ^ v139)
                                                                        v5 = uint64(v3)
                                                                        if false {
                                                                            if true {
                                                                                var v141 int32 = (v1 / ((v1) | 1))
                                                                                var v142 int64 = (v12 / ((9223372036854775807) | 1))
                                                                                switch fi138 {
                                                                                case 166:
                                                                                    switch fi137 {
                                                                                    case 127:
                                                                                        v3 = 2
                                                                                    case 2147483647:
                                                                                        v143 := GenS1{F0: (v6 - v5), F1: (true || ((v140 ^ 104) == 49)), F2: -1}
                                                                                        _ = v143
                                                                                    }
                                                                                case 9223372036854775807:
                                                                                    delete(v4, 9223372036854775807)
                                                                                }
                                                                                _ = v141
                                                                                _ = v142
                                                                            } else {
                                                                                v144 := []int{4294967295}
                                                                                _ = v144
                                                                            }
                                                                            v140 = (v140 / ((int8(v3)) | 1))
                                                                            if (fi138 <= fi138) {
                                                                                var v145 int64 = v12
                                                                                var v146 int8 = v139
                                                                                var v147 bool = (v1 > v1)
                                                                                _ = v145
                                                                                _ = v146
                                                                                _ = v147
                                                                            } else {
                                                                                var v148 float32 = 0.1
                                                                                var v149 bool = true
                                                                                var v150 uint64 = v6
                                                                                switch (("hello" + ((("hello" + "a") + ("" + "")) + "foo bar")) + ("" + "hello")) {
                                                                                case "a":
                                                                                    v151 := []int{fi138, int(v69)}
                                                                                    _ = v151
                                                                                case "hello":
                                                                                    var v152 int = (fi138 | 127)
                                                                                    _ = v152
                                                                                case "\n\t":
                                                                                    switch fi138 {
                                                                                    case -1:
                                                                                        v153 := v4[(fi137 * (^(fi137)))]
                                                                                        _ = v153
                                                                                    case -442:
                                                                                        v154 := GenS2{F0: (-264 >= 1), F1: (v139 ^ v140), F2: 18446744073709551615}
                                                                                        _ = v154
                                                                                    default:
                                                                                        v3 = (^(v3))
                                                                                    }
                                                                                default:
                                                                                    if ((v149 || ((v10 * (v10 << ((165) & 63))) > uint8(v3))) && (!((v12 - int64(v69)) != (v12 - (v12 | v12))))) {
                                                                                        v148 = 1e10
                                                                                        v155 := []int{fi137, fi138}
                                                                                        _ = v155
                                                                                    }
                                                                                }
                                                                                _ = v148
                                                                                _ = v149
                                                                                _ = v150
                                                                            }
                                                                            if false {
                                                                                v69 += (v69 * v69)
                                                                                var v156 int64 = v12
                                                                                if ((fi138 + fi138) > 42) {
                                                                                    fmt.Println(fi137)
                                                                                } else {
                                                                                    var v157 uint16 = v3
                                                                                    v158 := GenS4{F0: uint16(v69), F1: (v139 % ((v140) | 1)), F2: v139}
                                                                                    var v159 int64 = (-(v156))
                                                                                    _ = v157
                                                                                    _ = v158
                                                                                    _ = v159
                                                                                }
                                                                                _ = v156
                                                                            } else {
                                                                                v160 := GenS9{F0: v6}
                                                                                var v161 uint8 = (v10 % ((v10) | 1))
                                                                                _ = v160
                                                                                _ = v161
                                                                            }
                                                                        } else {
                                                                            var v162 int8 = (v140 | (^(v139)))
                                                                            fmt.Println(v12, v162)
                                                                            _ = v162
                                                                        }
                                                                        _ = v139
                                                                        _ = v140
                                                                    }
                                                                    var v163 int8 = 100
                                                                    v164 := GenS12{F0: v6, F1: 255}
                                                                    v1 -= 42
                                                                    if (v10 == 1) {
                                                                        var v165 uint = 65535
                                                                        _ = v165
                                                                    }
                                                                    _ = v163
                                                                    _ = v164
                                                                } else {
                                                                    v12 += v12
                                                                    if true {
                                                                        v3 += (v3 | 2)
                                                                        v6 -= (v5 * v5)
                                                                    } else {
                                                                        if false {
                                                                            fmt.Println(v5, v5, v3)
                                                                            if true {
                                                                                var v166 uint8 = 202
                                                                                v69 = 4294967295
                                                                                _ = v166
                                                                            } else {
                                                                                var v167 uint16 = 1
                                                                                delete(v4, fi138)
                                                                                v168 := []int{(fi137 - fi137), fi138, fi137, -551}
                                                                                _ = v167
                                                                                _ = v168
                                                                            }
                                                                            var v169 int8 = int8(v5)
                                                                            _ = v169
                                                                        }
                                                                        v170 := GenS13{F0: (v6 * 0), F1: "hello"}
                                                                        var v171 uint8 = v10
                                                                        _ = v170
                                                                        _ = v171
                                                                    }
                                                                    switch ((("foo bar" + ("\n\t" + "\n\t")) + ((("a" + ("hello" + "foo bar")) + (("hello" + ("" + ("a" + "hello"))) + "a")) + (("foo bar" + "hello") + "\n\t"))) + "foo bar") {
                                                                    case "a":
                                                                        if (true || (-316 != 32767)) {
                                                                            if true {
                                                                                switch "foo bar" {
                                                                                case "\n\t":
                                                                                    v3 += 255
                                                                                default:
                                                                                    fmt.Println(v69, fi138)
                                                                                }
                                                                            } else {
                                                                                switch fi138 {
                                                                                case 294:
                                                                                    if true {
                                                                                        var v172 uint16 = (v3 * v3)
                                                                                        _ = v172
                                                                                    } else {
                                                                                        v173 := GenS4{F0: (v3 | v3), F1: -128, F2: -84}
                                                                                        _ = v173
                                                                                    }
                                                                                }
                                                                            }
                                                                            var v174 int = fi137
                                                                            var v175 uint = (^(v69))
                                                                            _ = v174
                                                                            _ = v175
                                                                        }
                                                                    case "":
                                                                        v1 -= 0
                                                                    case "hello":
                                                                        delete(v4, (fi138 % ((fi138) | 1)))
                                                                    }
                                                                }
                                                            case 2147483647:
                                                                fmt.Println(v3, v12)
                                                            case -979:
                                                                v176 := []int{(fi138 & (^(fi138))), 127, fi138, fi137}
                                                                _ = v176
                                                            }
                                                            v4[4294967295] = (fi137 >> (((v69 % ((uint(v3)) | 1))) & 63))
                                                            _ = fi138
                                                        }
                                                        switch ("a" + ("foo bar" + "\n\t")) {
                                                        case "":
                                                            var v177 uint = (v69 - (v69 + (v69 / (((v69 >> ((uint(v3)) & 63))) | 1))))
                                                            _ = v177
                                                        case "hello":
                                                            if true {
                                                                v178 := v4[fi137]
                                                                if false {
                                                                    var v179 float32 = 1.5
                                                                    var v180 float64 = (3.14159 * 1.5)
                                                                    _ = v179
                                                                    _ = v180
                                                                } else {
                                                                    fmt.Println(v178, v12, fi137)
                                                                }
                                                                fmt.Println(v6, v178, v178)
                                                                _ = v178
                                                            } else {
                                                                v1 = (v1 * 211)
                                                                fmt.Println(v12, v1, v10)
                                                            }
                                                        }
                                                    }
                                                case "foo bar":
                                                    v181 := GenS12{F0: (v6 / (((v6 ^ 589)) | 1)), F1: v3}
                                                    _ = v181
                                                case "hello":
                                                    v1 -= -2147483648
                                                }
                                                v182 := v2[fi137]
                                                v6 += v6
                                                _ = fi137
                                                _ = v182
                                            }
                                        }
                                    }
                                    if (9223372036854775807 <= -87) {
                                        switch -2147483648 {
                                        case 9223372036854775807:
                                            v6 += (^(v6))
                                        case 454:
                                            v183 := []int{-481, 4294967295}
                                            _ = v183
                                        }
                                        v184 := GenS11{F0: (v10 % ((v10) | 1)), F1: true}
                                        _ = v184
                                    } else {
                                        v185 := v4[int(v69)]
                                        v186 := []int{v185, v185}
                                        var v187 int32 = 647
                                        _ = v185
                                        _ = v186
                                        _ = v187
                                    }
                                    var v188 int64 = 0
                                    _ = v188
                                }
                                _ = v69
                            } else {
                                if true {
                                    v12 += (v12 << ((42) & 63))
                                    v189 := v4[2147483647]
                                    for fi190 := 0; fi190 < 4; fi190++ {
                                        if (false || false) {
                                            if true {
                                                var v191 float32 = 1e10
                                                _ = v191
                                            } else {
                                                fmt.Println(v189, fi190, v189)
                                            }
                                            v5 = v5
                                        } else {
                                            switch (v189 ^ 32768) {
                                            case -829:
                                                for fi192 := 0; fi192 < 20; fi192++ {
                                                    var v193 uint = 825
                                                    if (0 > -1) {
                                                        if true {
                                                            if false {
                                                                v194 := v2[-9223372036854775808]
                                                                _ = v194
                                                            } else {
                                                                v195 := GenS4{F0: (v3 - v3), F1: 100, F2: 100}
                                                                v196 := GenS14{F0: true, F1: v193, F2: 42}
                                                                v10 = 2
                                                                _ = v195
                                                                _ = v196
                                                            }
                                                            fmt.Println(v6, v189, v5)
                                                        }
                                                    }
                                                    switch v189 {
                                                    case 0:
                                                        v3 -= uint16(v10)
                                                    }
                                                    var v197 uint8 = (v10 / ((v10) | 1))
                                                    _ = fi192
                                                    _ = v193
                                                    _ = v197
                                                }
                                            case 127:
                                                v12 = v12
                                            case -142:
                                                v5 -= 0
                                            }
                                            var v198 int16 = 32767
                                            _ = v198
                                        }
                                        v199 := []int{}
                                        var v200 int = int(v1)
                                        _ = fi190
                                        _ = v199
                                        _ = v200
                                    }
                                    _ = v189
                                }
                                v12 = 0
                            }
                            var v201 int16 = -1
                            var v202 float32 = (1e10 * 0.1)
                            _ = v201
                            _ = v202
                        }
                        if false {
                            v203 := GenS15{F0: 231, F1: 2}
                            v10 += 183
                            _ = v203
                        }
                    } else {
                        var v204 int = -9223372036854775808
                        _ = v204
                    }
                    v205 := GenS16{F0: "hello", F1: (v10 % (((v10 & (^(v10)))) | 1))}
                    _ = v205
                } else {
                    v5 = (v5 & 18446744073709551615)
                    for fi206 := 0; fi206 < 15; fi206++ {
                        switch fi206 {
                        case 127:
                            for fi207 := 0; fi207 < 15; fi207++ {
                                v208 := v4[9223372036854775807]
                                _ = fi207
                                _ = v208
                            }
                        case -928:
                            var v209 uint64 = v6
                            _ = v209
                        case 0:
                            for fi210 := 0; fi210 < 10; fi210++ {
                                if false {
                                    v211 := GenS8{F0: v12, F1: (-(fi206)), F2: uint64(v3)}
                                    _ = v211
                                } else {
                                    var v212 int64 = v12
                                    fmt.Println(v5, fi210, v5)
                                    if false {
                                        v3 = (v3 + 0)
                                        var v213 uint8 = v10
                                        _ = v213
                                    }
                                    v214 := []int{fi206}
                                    _ = v212
                                    _ = v214
                                }
                                v10 += 79
                                _ = fi210
                            }
                        }
                        if (int(v10) > (fi206 * -314)) {
                            var v215 uint8 = (v10 | (v10 ^ 42))
                            _ = v215
                        } else {
                            var v216 int32 = v1
                            _ = v216
                        }
                        _ = fi206
                    }
                }
            }
            v217 := GenS10{F0: v10}
            _ = v217
        }
        if ((0.0 - 1e10) != 0.1) {
            fmt.Println(v1, v3)
            v12 = 292
            if (874 == -158) {
                if false {
                    for fi218 := 0; fi218 < 20; fi218++ {
                        v219 := GenS9{F0: 160}
                        var v220 int32 = v1
                        fmt.Println(v220)
                        _ = fi218
                        _ = v219
                        _ = v220
                    }
                    v221 := GenS1{F0: v5, F1: true, F2: (v1 / ((1) | 1))}
                    _ = v221
                }
                switch 4294967295 {
                case -1:
                    switch -1 {
                    case -9223372036854775808:
                        for fi222 := 0; fi222 < 9; fi222++ {
                            for fi223 := 0; fi223 < 3; fi223++ {
                                switch fi223 {
                                case 9223372036854775807:
                                    if (v3 != (v3 << ((805) & 63))) {
                                        var v224 uint16 = (v3 << ((uint(v10)) & 63))
                                        _ = v224
                                    } else {
                                        var v225 int64 = (^(v12))
                                        _ = v225
                                    }
                                case -1:
                                    var v226 float32 = 0.0
                                    var v227 uint64 = 0
                                    _ = v226
                                    _ = v227
                                default:
                                    v3 = v3
                                }
                                _ = fi223
                            }
                            switch (((("" + "foo bar") + "foo bar") + "a") + (("" + (("hello" + ("foo bar" + "a")) + "")) + "")) {
                            case "\n\t":
                                if true {
                                    v12 = (v12 >> ((0) & 63))
                                    for fi228 := 0; fi228 < 8; fi228++ {
                                        v229 := []int{int(v5), fi222, (fi222 & fi228), (fi222 ^ 377)}
                                        switch ("\n\t" + (("a" + "foo bar") + "hello")) {
                                        case "hello":
                                            delete(v4, (-(fi228)))
                                        case "foo bar":
                                            _ = len(v229)
                                        case "\n\t":
                                            v229 = append(v229, (fi222 / (((fi228 % (((fi228 ^ fi222)) | 1))) | 1)))
                                        }
                                        _ = fi228
                                        _ = v229
                                    }
                                    delete(v2, int(v12))
                                } else {
                                    v3 = 731
                                    var v230 float32 = 1.5
                                    var v231 int64 = (v12 | (-(v12)))
                                    _ = v230
                                    _ = v231
                                }
                            default:
                                v12 -= -993
                            }
                            _ = fi222
                        }
                    case -131:
                        var v232 int8 = 23
                        _ = v232
                    case -372:
                        var v233 float32 = (3.14159 + 1.5)
                        _ = v233
                    }
                case -212:
                    for fi234 := 0; fi234 < 2; fi234++ {
                        v6 = v5
                        _ = fi234
                    }
                }
            } else {
                if false {
                    var v235 float32 = 1e10
                    var v236 string = (((("" + "foo bar") + "a") + "foo bar") + "a")
                    _ = v235
                    _ = v236
                }
                v237 := []int{}
                for fi238 := 0; fi238 < 10; fi238++ {
                    v10 = (^(v10))
                    _ = fi238
                }
                _ = v237
            }
        } else {
            for fi239 := 0; fi239 < 10; fi239++ {
                v240 := v4[fi239]
                _ = fi239
                _ = v240
            }
        }
        v5 -= (v5 / ((uint64(v10)) | 1))
    }
    if (42 <= 320) {
        switch -9223372036854775808 {
        case 32768:
            _ = len(v4)
        default:
            var v241 float64 = 0.1
            _ = v241
        }
        if false {
            v12 = (v12 + v12)
        } else {
            var v242 int32 = v1
            v243 := GenS12{F0: (v5 * (v5 >> ((65535) & 63))), F1: uint16(v6)}
            var v244 float64 = (1e10 - 0.0)
            _ = v242
            _ = v243
            _ = v244
        }
    } else {
        if false {
            v245 := GenS17{F0: (true || (!(true || ((1 == 32767) || (0 < 4294967295))))), F1: true}
            v246 := GenS7{F0: true, F1: "foo bar"}
            _ = v245
            _ = v246
        } else {
            if false {
                v10 -= 2
                switch "" {
                case "hello":
                    var v247 string = ((("\n\t" + "\n\t") + "a") + "\n\t")
                    _ = v247
                case "a":
                    v248 := []int{-2147483648, 27, 854, 1}
                    _ = v248
                case "foo bar":
                    delete(v2, 42)
                default:
                    v12 = 442
                }
            } else {
                var v249 string = "\n\t"
                if false {
                    var v250 int8 = 127
                    _ = v250
                } else {
                    var v251 uint64 = v5
                    _ = v251
                }
                if true {
                    v252 := GenS1{F0: v5, F1: (v3 <= (v3 & v3)), F2: 0}
                    _ = v252
                }
                _ = v249
            }
            var v253 float32 = 1e10
            _ = v253
        }
        v254 := []int{}
        _ = v254
    }
    v255 := []int{-1}
    v256 := GenS5{F0: (v3 | v3), F1: v10, F2: 1}
    if len(v255) >= 2 {
        v255 = v255[1:len(v255)]
    }
    for fi257 := 0; fi257 < 8; fi257++ {
        if true {
            fmt.Println(v1)
            if true {
                var v258 int16 = 32767
                var v259 uint64 = v6
                _ = v258
                _ = v259
            } else {
                v1 = v1
                for fi260 := 0; fi260 < 4; fi260++ {
                    if (0.1 == 3.14159) {
                        switch (fi257 - fi257) {
                        case 42:
                            v261 := GenS18{F0: (v1 % ((-759) | 1)), F1: 65535}
                            _ = v261
                        default:
                            if false {
                                fmt.Println(v3)
                                switch "foo bar" {
                                case "":
                                    v10 = (^(v10))
                                }
                                if false {
                                    v262 := GenS1{F0: (v6 + 791), F1: false, F2: (-(v1))}
                                    var v263 int16 = 0
                                    _ = v262
                                    _ = v263
                                } else {
                                    v10 += (v10 % ((2) | 1))
                                }
                            }
                        }
                    }
                    _ = fi260
                }
                var v264 float32 = (((-2.25 - float32((1.5 - -2.25))) + (0.1 * (-2.25 * (float32((1e10 * -2.25)) + ((0.1 + (((-2.25 / 1e10) - 1e10) * (1.5 * (((1e10 / float32((3.14159 - -2.25))) + ((0.1 / (0.1 * float32(float64(0.1)))) - 1.5)) + (1e10 - -2.25))))) / 1.5))))) - 1.5)
                _ = v264
            }
        }
        if ((!false) || true) {
            v10 -= 199
            v3 -= (v3 * v3)
            var v265 int16 = 53
            _ = v265
        } else {
            switch "hello" {
            case "":
                v266 := GenS19{F0: 0, F1: v3}
                _ = v266
            default:
                v267 := GenS20{F0: (v10 | (v10 + v10)), F1: v1, F2: 1}
                _ = v267
            }
            if true {
                v268 := GenS17{F0: (v10 < v10), F1: true}
                _ = v268
            }
        }
        v6 -= 4294967295
        _ = fi257
    }
    if (!false) {
        var v269 uint16 = (v3 + (v3 << ((22) & 63)))
        _ = v269
    } else {
        delete(v2, -2147483648)
    }
    for fi270 := 0; fi270 < 5; fi270++ {
        var v271 float64 = 0.1
        _ = fi270
        _ = v271
    }
    for fi272 := 0; fi272 < 1; fi272++ {
        var v273 uint8 = 42
        v5 = (^(v6))
        _ = fi272
        _ = v273
    }
    fmt.Println(v5, v3, v5)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v255:", len(v255), v255)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v255
    _ = v256
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
