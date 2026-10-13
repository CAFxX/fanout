package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint64; F1 string }
type GenS5 struct { F0 int8; F1 uint16 }
type GenS6 struct { F0 bool; F1 uint8 }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 int64 }
type GenS9 struct { F0 int; F1 bool; F2 uint8 }
type GenS10 struct { F0 int64; F1 uint; F2 int16 }
type GenS11 struct { F0 int }
type GenS12 struct { F0 int8; F1 uint32 }
type GenS13 struct { F0 bool }
type GenS14 struct { F0 int; F1 int; F2 int64 }
type GenS15 struct { F0 int8 }
type GenS16 struct { F0 string; F1 int8 }
type GenS17 struct { F0 bool }
type GenS18 struct { F0 int16; F1 int64 }
type GenS19 struct { F0 int32; F1 uint16 }
type GenS20 struct { F0 uint; F1 int64 }
type GenS21 struct { F0 int64; F1 uint; F2 bool }
type GenS22 struct { F0 uint64 }
type GenS23 struct { F0 string }
type GenS24 struct { F0 string; F1 uint }
type GenS25 struct { F0 uint32; F1 int64 }
type GenS26 struct { F0 int }
type GenS27 struct { F0 bool; F1 bool }
type GenS28 struct { F0 int8 }
type GenS29 struct { F0 uint; F1 int64; F2 uint }
type GenS30 struct { F0 int; F1 uint16 }
type GenS31 struct { F0 uint }
func genFn0() int8 {
    v6 := []int{42, 537, 127, -9223372036854775808}
    if true {
        if true {
            var v7 uint16 = 42
            var v8 uint8 = 68
            var v9 int32 = 42
            _ = v7
            _ = v8
            _ = v9
        } else {
            _ = len(v6)
            var v10 uint16 = 0
            v11 := GenS1{F0: (v10 ^ 1)}
            _ = v10
            _ = v11
        }
        v12 := GenS1{F0: 0}
        _ = v12
    }
    _ = v6
    return 21
}

func main() {
    var v1 uint8 = 42
    var v2 string = ("\n\t" + "\n\t")
    v3 := make(map[string]int)
    var v4 uint16 = 32767
    var v5 int16 = 19
    var v13 uint32 = 2147483647
    var v14 float64 = 3.14159
    v13 += (^(v13))
    v13 = 491
    v1 += 213
    var v15 uint16 = v4
    var v16 int8 = 65
    var v17 int32 = 744
    for fi18 := 0; fi18 < 16; fi18++ {
        if (v5 >= 409) {
            if true {
                fmt.Println(v4, v4)
            } else {
                var v19 uint16 = (v4 + (^(v4)))
                v17 -= v17
                _ = v19
            }
            var v20 uint64 = 0
            fmt.Println(v4, v16, v5)
            _ = v20
        } else {
            if (v15 != 1) {
                var v21 int8 = genFn0()
                _ = v21
            }
            var v22 uint16 = v4
            v4 = v15
            _ = v22
        }
        _ = fi18
    }
    v13 = 1
    if true {
        v23 := GenS2{F0: 32768}
        for fi24 := 0; fi24 < 17; fi24++ {
            v25 := GenS2{F0: -143}
            if false {
                v17 -= (v17 >> ((uint(v1)) & 63))
                v26 := []int{fi24}
                _ = v26
            } else {
                for fi27 := 0; fi27 < 18; fi27++ {
                    fmt.Println(v4)
                    _ = fi27
                }
            }
            v16 = v16
            _ = fi24
            _ = v25
        }
        _ = v23
    } else {
        v28 := []int{}
        if true {
            v29 := GenS3{F0: v17}
            _ = v29
        }
        fmt.Println(v4, v14, v16)
        _ = v28
    }
    if (!true) {
        var v30 int16 = v5
        _ = v30
    }
    for fi31 := 0; fi31 < 6; fi31++ {
        if (!(true || true)) {
            for fi32 := 0; fi32 < 12; fi32++ {
                if false {
                    v1 += v1
                    var v33 bool = true
                    var v34 int64 = 0
                    _ = v33
                    _ = v34
                } else {
                    var v35 string = (v2 + ("" + v2))
                    switch fi31 {
                    case -477:
                        if true {
                            switch (fi31 & (fi31 | fi31)) {
                            case 1:
                                if false {
                                    var v36 int = (^(fi32))
                                    if ((v5 - (v5 % ((v5) | 1))) >= (v5 & (v5 >> ((0) & 63)))) {
                                        v14 -= v14
                                        var v37 string = v2
                                        switch int(v16) {
                                        case 2147483647:
                                            v38 := GenS3{F0: -2147483648}
                                            _ = v38
                                        case -801:
                                            v39 := GenS4{F0: 42, F1: v37}
                                            _ = v39
                                        }
                                        v4 = v15
                                        _ = v37
                                    }
                                    _ = v36
                                }
                            default:
                                v13 = (v13 << ((0) & 63))
                            }
                            switch "\n\t" {
                            case "\n\t":
                                v1 = v1
                            case "hello":
                                v40 := []int{}
                                _ = v40
                            default:
                                switch (v35 + v35) {
                                case "hello":
                                    var v41 int8 = (v16 ^ v16)
                                    _ = v41
                                case "\n\t":
                                    v2 = ("foo bar" + "")
                                case "a":
                                    if true {
                                        v2 = ("a" + "hello")
                                        v42 := v3[((v2 + "a") + (("hello" + v2) + ((v35 + v2) + "hello")))]
                                        var v43 float32 = 3.14159
                                        _ = v42
                                        _ = v43
                                    }
                                default:
                                    var v44 float32 = 1e10
                                    _ = v44
                                }
                            }
                            if true {
                                fmt.Println(v15, fi31, v13)
                                if ((v1 >> ((88) & 63)) != v1) {
                                    var v45 bool = (3.14159 >= (((v14 * ((v14 / v14) / v14)) * v14) + v14))
                                    v46 := []int{}
                                    _ = v45
                                    _ = v46
                                } else {
                                    v14 = float64(0.0)
                                }
                                var v47 int = 32767
                                _ = v47
                            }
                        } else {
                            v48 := v3[v2]
                            v16 -= v16
                            _ = v48
                        }
                    case 42:
                        var v49 uint = 1
                        _ = v49
                    default:
                        v4 = v4
                        v15 = uint16(v13)
                        if true {
                            var v50 int32 = v17
                            v51 := []int{fi31, (^(fi31)), fi31}
                            if ((v1 ^ 127) < v1) {
                                v5 = (v5 | (v5 >> ((262) & 63)))
                                v52 := GenS5{F0: (v16 >> ((414) & 63)), F1: uint16(v13)}
                                _ = v52
                            } else {
                                var v53 uint64 = 891
                                v13 = (v13 ^ 65535)
                                _ = v53
                            }
                            switch v2 {
                            case "hello":
                                var v54 float64 = (-2.25 - v14)
                                _ = v54
                            case "":
                                var v55 int = int(v5)
                                _ = v55
                            case "a":
                                if ((v17 >> ((255) & 63)) <= (v17 % (((v17 + v50)) | 1))) {
                                    if (!false) {
                                        v56 := GenS6{F0: true, F1: (^(v1))}
                                        var v57 uint = 255
                                        _ = v56
                                        _ = v57
                                    }
                                    if false {
                                        v4 = v15
                                        v4 -= v4
                                        v35 = (((v35 + (("hello" + ("\n\t" + v2)) + "\n\t")) + (v2 + (v35 + ("foo bar" + v2)))) + ("foo bar" + ((((v35 + (("a" + ("a" + v2)) + v35)) + (v35 + ((v35 + (v35 + v2)) + v35))) + (v2 + v35)) + (v35 + "a"))))
                                    } else {
                                        v58 := GenS7{F0: v5}
                                        switch "foo bar" {
                                        case "foo bar":
                                            if (float32((v14 * v14)) == 0.1) {
                                                switch "\n\t" {
                                                case "a":
                                                    if ((!false) && ((v1 ^ 118) < (v1 ^ (v1 + 1)))) {
                                                        if ((v16 + genFn0()) < (v16 + (v16 * (v16 + -1)))) {
                                                            v59 := GenS8{F0: 70}
                                                            _ = v59
                                                        } else {
                                                            v50 += (v50 / ((42) | 1))
                                                            v60 := GenS9{F0: fi31, F1: true, F2: v1}
                                                            fmt.Println(fi31, v50)
                                                            _ = v60
                                                        }
                                                        var v61 int8 = (v16 ^ (v16 * int8(v50)))
                                                        _ = v61
                                                    } else {
                                                        switch (v35 + (v2 + v2)) {
                                                        case "hello":
                                                            var v62 uint8 = (v1 << ((273) & 63))
                                                            _ = v62
                                                        case "\n\t":
                                                            fmt.Println(v5)
                                                        case "a":
                                                            var v63 uint64 = uint64(v1)
                                                            _ = v63
                                                        }
                                                        if len(v51) >= 2 {
                                                            v51 = v51[1:len(v51)]
                                                        }
                                                        switch (fi31 - (^(fi31))) {
                                                        case 42:
                                                            v64 := GenS1{F0: v15}
                                                            v3[((("" + (v2 + "hello")) + ("foo bar" + (v2 + "\n\t"))) + v35)] = fi31
                                                            var v65 uint32 = 0
                                                            if len(v51) > 0 {
                                                                v51[0] = fi31
                                                            }
                                                            _ = v64
                                                            _ = v65
                                                        default:
                                                            var v66 int16 = v5
                                                            if len(v51) > 0 {
                                                                v51[0] = fi31
                                                            }
                                                            var v67 uint8 = (v1 & 2)
                                                            _ = v66
                                                            _ = v67
                                                        }
                                                        v68 := GenS10{F0: -858, F1: 0, F2: (^(v5))}
                                                        _ = v68
                                                    }
                                                case "\n\t":
                                                    switch v35 {
                                                    case "a":
                                                        if true {
                                                            v69 := GenS10{F0: -49, F1: 18446744073709551615, F2: v5}
                                                            fmt.Println(v4)
                                                            var v70 int32 = (v17 << ((2) & 63))
                                                            _ = v69
                                                            _ = v70
                                                        } else {
                                                            v71 := GenS11{F0: fi31}
                                                            var v72 string = v2
                                                            var v73 int16 = v5
                                                            _ = v71
                                                            _ = v72
                                                            _ = v73
                                                        }
                                                    }
                                                case "foo bar":
                                                    v74 := GenS3{F0: -2147483648}
                                                    _ = v74
                                                }
                                                if ((v50 % ((v50) | 1)) != -75) {
                                                    var v75 uint64 = 1
                                                    v76 := GenS12{F0: (v16 - (v16 % ((v16) | 1))), F1: 1}
                                                    switch fi32 {
                                                    case -720:
                                                        _ = len(v3)
                                                    case 32767:
                                                        fmt.Println(v2, v15)
                                                    case 271:
                                                        var v77 int8 = -1
                                                        _ = v77
                                                    }
                                                    _ = v75
                                                    _ = v76
                                                }
                                                v78 := GenS13{F0: false}
                                                _ = v78
                                            } else {
                                                if true {
                                                    if ((float64(((1.5 * 0.0) - float32((1e10 - (v14 * 0.1))))) + v14) != 1.5) {
                                                        if true {
                                                            fmt.Println(v4)
                                                            v79 := GenS3{F0: 194}
                                                            _ = v79
                                                        } else {
                                                            v4 = (v15 / (((v4 + v15)) | 1))
                                                        }
                                                        if true {
                                                            v14 = (v14 - 0.0)
                                                        }
                                                    } else {
                                                        switch (fi32 / ((0) | 1)) {
                                                        case 127:
                                                            switch "foo bar" {
                                                            case "hello":
                                                                if false {
                                                                    var v80 int32 = v17
                                                                    _ = v80
                                                                } else {
                                                                    if true {
                                                                        v81 := GenS14{F0: int(v5), F1: (^(fi31)), F2: -9223372036854775808}
                                                                        _ = v81
                                                                    } else {
                                                                        v2 = "a"
                                                                        fmt.Println(v35)
                                                                        var v82 float32 = ((1.5 / -2.25) * 1.5)
                                                                        v50 = int32(fi31)
                                                                        v5 -= v5
                                                                        v83 := GenS1{F0: 65535}
                                                                        _ = v82
                                                                        _ = v83
                                                                    }
                                                                    fmt.Println(v2, v50, fi31)
                                                                    v5 += v5
                                                                    var v84 uint16 = (^(v15))
                                                                    var v85 uint16 = v4
                                                                    _ = v84
                                                                    _ = v85
                                                                }
                                                            default:
                                                                if false {
                                                                    var v86 float32 = (3.14159 - (0.1 - float32(v14)))
                                                                    _ = len(v3)
                                                                    v50 = (v50 + v17)
                                                                    _ = v86
                                                                }
                                                            }
                                                        case 9223372036854775807:
                                                            switch (fi31 % (((^(fi32))) | 1)) {
                                                            case 972:
                                                                if true {
                                                                    if (!((false || ((4294967295 == 0) || (((v5 % (((v5 ^ -1)) | 1)) <= v5) || (true && false)))) || false)) {
                                                                        if ((0 <= 0) && (!false)) {
                                                                            var v87 uint64 = 221
                                                                            fmt.Println(v35, fi32, v15)
                                                                            _ = v87
                                                                        }
                                                                        if (1 >= -1) {
                                                                            var v88 int64 = 1
                                                                            _ = v88
                                                                        } else {
                                                                            v89 := GenS10{F0: 1, F1: 1, F2: -785}
                                                                            _ = v89
                                                                        }
                                                                    } else {
                                                                        v90 := GenS15{F0: v16}
                                                                        _ = v90
                                                                    }
                                                                    v91 := v3[v2]
                                                                    v17 -= (v50 / ((42) | 1))
                                                                    _ = v91
                                                                } else {
                                                                    if (uint64(v1) > 2) {
                                                                        var v92 uint32 = 1
                                                                        _ = v92
                                                                    } else {
                                                                        v93 := v3["foo bar"]
                                                                        _ = v93
                                                                    }
                                                                    v94 := GenS2{F0: 32768}
                                                                    v51 = append(v51, (fi31 - int(v5)))
                                                                    if false {
                                                                        var v95 int8 = (^(v16))
                                                                        _ = len(v51)
                                                                        v13 -= v13
                                                                        _ = v95
                                                                    } else {
                                                                        delete(v3, v2)
                                                                        if true {
                                                                            var v96 string = "foo bar"
                                                                            var v97 int64 = 9223372036854775807
                                                                            v97 = (-(v97))
                                                                            _ = v96
                                                                            _ = v97
                                                                        } else {
                                                                            var v98 uint = 0
                                                                            _ = v98
                                                                        }
                                                                        switch (fi31 ^ fi31) {
                                                                        case 32767:
                                                                            switch (((((("foo bar" + v35) + ("a" + (("hello" + (v2 + (v2 + "hello"))) + ("" + "foo bar")))) + "") + (v2 + "a")) + (v2 + ("\n\t" + "foo bar"))) + (v2 + (v2 + "foo bar"))) {
                                                                            case "\n\t":
                                                                                var v99 string = ("hello" + "")
                                                                                _ = v99
                                                                            }
                                                                        case 35:
                                                                            v1 = (v1 - (v1 >> ((42) & 63)))
                                                                        default:
                                                                            switch (fi31 << ((8) & 63)) {
                                                                            case -457:
                                                                                fmt.Println(v17)
                                                                            case -9223372036854775808:
                                                                                v100 := v3[("a" + "\n\t")]
                                                                                _ = v100
                                                                            default:
                                                                                var v101 int32 = v50
                                                                                var v102 string = v35
                                                                                _ = v101
                                                                                _ = v102
                                                                            }
                                                                            if (v16 > (v16 | 27)) {
                                                                                v103 := GenS16{F0: ("\n\t" + (v2 + (v35 + "foo bar"))), F1: int8(v5)}
                                                                                var v104 float32 = (0.1 / 3.14159)
                                                                                _ = v103
                                                                                _ = v104
                                                                            } else {
                                                                                v17 = int32(v13)
                                                                                switch v35 {
                                                                                case "":
                                                                                    var v105 uint = 2
                                                                                    _ = v105
                                                                                case "\n\t":
                                                                                    switch -9223372036854775808 {
                                                                                    case 4294967295:
                                                                                        if ((fi32 ^ (fi31 * (fi31 ^ (fi31 / ((fi31) | 1))))) == (fi32 % (((fi32 | (fi32 - (-(fi31))))) | 1))) {
                                                                                            if len(v51) >= 2 {
                                                                                                v51 = v51[1:len(v51)]
                                                                                            }
                                                                                            if (true || ((v13 / ((v13) | 1)) == v13)) {
                                                                                                var v106 int8 = -35
                                                                                                _ = v106
                                                                                            }
                                                                                            switch v35 {
                                                                                            case "foo bar":
                                                                                                fmt.Println(v17)
                                                                                            case "":
                                                                                                var v107 uint64 = 778
                                                                                                _ = v107
                                                                                            default:
                                                                                                fmt.Println(v50)
                                                                                            }
                                                                                        }
                                                                                    default:
                                                                                        switch fi32 {
                                                                                        case -2147483648:
                                                                                            if false {
                                                                                                var v108 string = v2
                                                                                                if len(v51) > 0 {
                                                                                                    v51[0] = 401
                                                                                                }
                                                                                                v5 += 285
                                                                                                _ = v108
                                                                                            }
                                                                                        case 980:
                                                                                            switch 61 {
                                                                                            case 1:
                                                                                                fmt.Println(fi32, fi31)
                                                                                            case 32767:
                                                                                                if false {
                                                                                                    if (v4 >= v15) {
                                                                                                        fmt.Println(fi32)
                                                                                                        v50 = -774
                                                                                                    }
                                                                                                    v13 += v13
                                                                                                    if false {
                                                                                                        var v109 int = -2147483648
                                                                                                        _ = v109
                                                                                                    } else {
                                                                                                        switch v2 {
                                                                                                        case "\n\t":
                                                                                                            if (float64(float32((v14 + v14))) > 0.1) {
                                                                                                                if false {
                                                                                                                    if true {
                                                                                                                        v5 = 42
                                                                                                                    }
                                                                                                                    fmt.Println(v35, fi32, v5)
                                                                                                                    v1 += v1
                                                                                                                    v13 = 4294967295
                                                                                                                    v50 = 42
                                                                                                                    switch (-(fi32)) {
                                                                                                                    case 128:
                                                                                                                        var v110 uint = 4294967295
                                                                                                                        _ = v110
                                                                                                                    }
                                                                                                                    v5 += v5
                                                                                                                }
                                                                                                                v50 = v17
                                                                                                                var v111 int16 = v5
                                                                                                                if (!true) {
                                                                                                                    v14 = float64(1.5)
                                                                                                                } else {
                                                                                                                    switch ("foo bar" + "foo bar") {
                                                                                                                    case "foo bar":
                                                                                                                        switch 0 {
                                                                                                                        case 32768:
                                                                                                                            fmt.Println(v16, v15, v5)
                                                                                                                        default:
                                                                                                                            v3[v35] = -1
                                                                                                                        }
                                                                                                                    case "\n\t":
                                                                                                                        switch fi31 {
                                                                                                                        case -2147483648:
                                                                                                                            var v112 float32 = (3.14159 * float32(-2.25))
                                                                                                                            _ = v112
                                                                                                                        case 4294967295:
                                                                                                                            v113 := v3[v35]
                                                                                                                            _ = v113
                                                                                                                        }
                                                                                                                    case "a":
                                                                                                                        switch fi32 {
                                                                                                                        case 32767:
                                                                                                                            v16 = v16
                                                                                                                        default:
                                                                                                                            var v114 int16 = v5
                                                                                                                            _ = v114
                                                                                                                        }
                                                                                                                        v2 = ""
                                                                                                                    default:
                                                                                                                        if false {
                                                                                                                            var v115 int32 = (v50 << ((194) & 63))
                                                                                                                            _ = v115
                                                                                                                        }
                                                                                                                    }
                                                                                                                    fmt.Println(v16, v50)
                                                                                                                    var v116 float32 = (0.1 * 1e10)
                                                                                                                    _ = v116
                                                                                                                }
                                                                                                                _ = v111
                                                                                                            }
                                                                                                        default:
                                                                                                            if true {
                                                                                                                switch v2 {
                                                                                                                case "":
                                                                                                                    if (true || false) {
                                                                                                                        var v117 uint8 = (v1 << ((18446744073709551615) & 63))
                                                                                                                        var v118 uint16 = (v15 & 0)
                                                                                                                        var v119 float32 = 1e10
                                                                                                                        if false {
                                                                                                                            var v120 uint32 = (v13 / (((v13 - uint32(v1))) | 1))
                                                                                                                            _ = v120
                                                                                                                        }
                                                                                                                        _ = v117
                                                                                                                        _ = v118
                                                                                                                        _ = v119
                                                                                                                    } else {
                                                                                                                        v3[v2] = fi32
                                                                                                                    }
                                                                                                                case "foo bar":
                                                                                                                    v51 = append(v51, 417)
                                                                                                                case "a":
                                                                                                                    v3[""] = (-(fi31))
                                                                                                                default:
                                                                                                                    fmt.Println(v1)
                                                                                                                }
                                                                                                                v121 := GenS5{F0: (v16 * v16), F1: 42}
                                                                                                                switch fi31 {
                                                                                                                case 32768:
                                                                                                                    fmt.Println(v4, v15, v15)
                                                                                                                case -132:
                                                                                                                    if (!true) {
                                                                                                                        switch (fi31 / ((fi32) | 1)) {
                                                                                                                        case 1:
                                                                                                                            v1 += 0
                                                                                                                        case 32767:
                                                                                                                            v35 = "foo bar"
                                                                                                                        default:
                                                                                                                            v122 := GenS11{F0: (fi31 | fi32)}
                                                                                                                            _ = v122
                                                                                                                        }
                                                                                                                        v2 = v35
                                                                                                                    } else {
                                                                                                                        if true {
                                                                                                                            delete(v3, v35)
                                                                                                                        }
                                                                                                                        v123 := GenS17{F0: false}
                                                                                                                        _ = v123
                                                                                                                    }
                                                                                                                case 256:
                                                                                                                    _ = len(v3)
                                                                                                                }
                                                                                                                _ = v121
                                                                                                            } else {
                                                                                                                if ("" != "a") {
                                                                                                                    if (v1 >= 42) {
                                                                                                                        if true {
                                                                                                                            fmt.Println(v13)
                                                                                                                        }
                                                                                                                        v13 -= 1
                                                                                                                        v14 += 3.14159
                                                                                                                    }
                                                                                                                    v124 := GenS18{F0: int16(v4), F1: 1}
                                                                                                                    fmt.Println(v15)
                                                                                                                    _ = v124
                                                                                                                }
                                                                                                                switch "" {
                                                                                                                case "a":
                                                                                                                    var v125 float64 = float64(1.5)
                                                                                                                    _ = v125
                                                                                                                default:
                                                                                                                    if true {
                                                                                                                        var v126 bool = (!true)
                                                                                                                        _ = v126
                                                                                                                    }
                                                                                                                }
                                                                                                                v17 = 1
                                                                                                            }
                                                                                                            v1 = v1
                                                                                                        }
                                                                                                        var v127 uint = 898
                                                                                                        if (-604 != 42) {
                                                                                                            v16 -= (v16 | v16)
                                                                                                            if (true || false) {
                                                                                                                v128 := GenS19{F0: -1, F1: (v4 | v15)}
                                                                                                                var v129 int = fi31
                                                                                                                v14 -= (v14 + 1e10)
                                                                                                                v16 += v16
                                                                                                                var v130 bool = false
                                                                                                                _ = v128
                                                                                                                _ = v129
                                                                                                                _ = v130
                                                                                                            } else {
                                                                                                                switch (v2 + v2) {
                                                                                                                case "":
                                                                                                                    if (true && ((0.1 * (1.5 * (1.5 + (1.5 / 1e10)))) == float32(0.1))) {
                                                                                                                        var v131 uint8 = uint8(v15)
                                                                                                                        v16 -= (-(v16))
                                                                                                                        _ = v131
                                                                                                                    }
                                                                                                                case "hello":
                                                                                                                    var v132 uint32 = uint32(v1)
                                                                                                                    _ = v132
                                                                                                                }
                                                                                                                v17 -= (v50 - 0)
                                                                                                                v1 -= (^(v1))
                                                                                                                v50 = int32(fi31)
                                                                                                            }
                                                                                                        } else {
                                                                                                            if (!(0 <= 83)) {
                                                                                                                var v133 int64 = 9223372036854775807
                                                                                                                _ = v133
                                                                                                            } else {
                                                                                                                v127 += (v127 >> ((v127) & 63))
                                                                                                                _ = len(v51)
                                                                                                                v134 := GenS11{F0: (fi32 << ((v127) & 63))}
                                                                                                                switch fi31 {
                                                                                                                case 32767:
                                                                                                                    v50 = (v50 * (v50 + v17))
                                                                                                                case 9223372036854775807:
                                                                                                                    v15 -= (v4 & uint16(v13))
                                                                                                                case 127:
                                                                                                                    v14 = v14
                                                                                                                }
                                                                                                                var v135 uint32 = v13
                                                                                                                _ = v134
                                                                                                                _ = v135
                                                                                                            }
                                                                                                        }
                                                                                                        var v136 bool = (true && (false && (!true)))
                                                                                                        _ = len(v3)
                                                                                                        v137 := GenS2{F0: fi31}
                                                                                                        if ((!true) || (fi32 <= (fi31 - (fi31 >> ((v127) & 63))))) {
                                                                                                            fmt.Println(fi31)
                                                                                                            var v138 float32 = 1e10
                                                                                                            v139 := GenS20{F0: 954, F1: -336}
                                                                                                            _ = v138
                                                                                                            _ = v139
                                                                                                        } else {
                                                                                                            if false {
                                                                                                                var v140 int8 = v16
                                                                                                                _ = v140
                                                                                                            } else {
                                                                                                                var v141 int = (fi32 + (-(fi32)))
                                                                                                                switch (fi32 * 8) {
                                                                                                                case 9223372036854775807:
                                                                                                                    v5 = -1
                                                                                                                case 203:
                                                                                                                    fmt.Println(v1, v13)
                                                                                                                case 209:
                                                                                                                    if (!(!((v127 % ((uint(v13)) | 1)) > uint(v1)))) {
                                                                                                                        if (248 < 0) {
                                                                                                                            if true {
                                                                                                                                if (0 != uint8(v13)) {
                                                                                                                                    v3[v2] = (fi32 << (((v127 & (v127 & (v127 * 4294967295)))) & 63))
                                                                                                                                    if (255 == v15) {
                                                                                                                                        v14 += (v14 / (3.14159 / 1.0))
                                                                                                                                        switch 0 {
                                                                                                                                        case -9223372036854775808:
                                                                                                                                            v127 = 1
                                                                                                                                        }
                                                                                                                                        v142 := GenS17{F0: true}
                                                                                                                                        _ = v142
                                                                                                                                    }
                                                                                                                                } else {
                                                                                                                                    if (false || (((float32(v14) + (((1e10 * (float32((v14 * ((v14 * (v14 + v14)) / v14))) / -2.25)) * -2.25) / 1e10)) + -2.25) <= (-2.25 + 1e10))) {
                                                                                                                                        var v143 uint8 = (^(v1))
                                                                                                                                        _ = v143
                                                                                                                                    } else {
                                                                                                                                        v144 := GenS4{F0: 1, F1: "foo bar"}
                                                                                                                                        _ = v144
                                                                                                                                    }
                                                                                                                                    var v145 uint = 255
                                                                                                                                    switch 127 {
                                                                                                                                    case 42:
                                                                                                                                        v3[(v2 + v35)] = v141
                                                                                                                                        v14 += 1e10
                                                                                                                                        v50 += v17
                                                                                                                                    case 4294967295:
                                                                                                                                        switch v35 {
                                                                                                                                        case "hello":
                                                                                                                                            switch v2 {
                                                                                                                                            case "a":
                                                                                                                                                v146 := v3[((v2 + v35) + (("hello" + (v2 + v2)) + v2))]
                                                                                                                                                _ = v146
                                                                                                                                            }
                                                                                                                                            if ((!v136) || (((-2.25 + 1.5) - 3.14159) > 0.0)) {
                                                                                                                                                switch (v35 + ("" + v2)) {
                                                                                                                                                case "a":
                                                                                                                                                    var v147 uint = (v127 / (((^(v127))) | 1))
                                                                                                                                                    _ = v147
                                                                                                                                                case "\n\t":
                                                                                                                                                    if (uint32(v1) <= uint32(v127)) {
                                                                                                                                                        if false {
                                                                                                                                                            if len(v51) >= 2 {
                                                                                                                                                                v51 = v51[1:len(v51)]
                                                                                                                                                            }
                                                                                                                                                            v148 := GenS10{F0: -241, F1: v145, F2: 42}
                                                                                                                                                            if true {
                                                                                                                                                                switch v2 {
                                                                                                                                                                case "foo bar":
                                                                                                                                                                    if len(v51) > 0 {
                                                                                                                                                                        v51[0] = (-(v141))
                                                                                                                                                                    }
                                                                                                                                                                default:
                                                                                                                                                                    v149 := GenS17{F0: v136}
                                                                                                                                                                    _ = v149
                                                                                                                                                                }
                                                                                                                                                                v150 := GenS21{F0: -9223372036854775808, F1: (v127 >> (((v145 * v127)) & 63)), F2: v136}
                                                                                                                                                                _ = v150
                                                                                                                                                            } else {
                                                                                                                                                                if (v16 == -128) {
                                                                                                                                                                    fmt.Println(v35)
                                                                                                                                                                    if v136 {
                                                                                                                                                                        v13 += (v13 + 381)
                                                                                                                                                                        v145 += 255
                                                                                                                                                                    }
                                                                                                                                                                    var v151 float32 = (0.1 * float32(0.1))
                                                                                                                                                                    _ = v151
                                                                                                                                                                } else {
                                                                                                                                                                    var v152 string = v2
                                                                                                                                                                    _ = v152
                                                                                                                                                                }
                                                                                                                                                                var v153 float64 = 1.5
                                                                                                                                                                fmt.Println(v14, v35, fi31)
                                                                                                                                                                _ = v153
                                                                                                                                                            }
                                                                                                                                                            _ = v148
                                                                                                                                                        } else {
                                                                                                                                                            v13 = 4294967295
                                                                                                                                                            v127 = 2
                                                                                                                                                        }
                                                                                                                                                    } else {
                                                                                                                                                        v145 = uint(v15)
                                                                                                                                                    }
                                                                                                                                                case "":
                                                                                                                                                    fmt.Println(v2, v145)
                                                                                                                                                default:
                                                                                                                                                    var v154 int32 = v50
                                                                                                                                                    _ = v154
                                                                                                                                                }
                                                                                                                                            } else {
                                                                                                                                                if false {
                                                                                                                                                    delete(v3, v2)
                                                                                                                                                    var v155 uint64 = 2
                                                                                                                                                    fmt.Println(v127, v16, fi31)
                                                                                                                                                    _ = v155
                                                                                                                                                }
                                                                                                                                            }
                                                                                                                                        case "":
                                                                                                                                            var v156 int8 = 100
                                                                                                                                            _ = v156
                                                                                                                                        case "\n\t":
                                                                                                                                            v141 += (^(fi31))
                                                                                                                                        default:
                                                                                                                                            _ = len(v51)
                                                                                                                                        }
                                                                                                                                    default:
                                                                                                                                        var v157 uint16 = 413
                                                                                                                                        _ = v157
                                                                                                                                    }
                                                                                                                                    _ = v145
                                                                                                                                }
                                                                                                                                v15 += (^(v4))
                                                                                                                                v4 += v15
                                                                                                                                if v136 {
                                                                                                                                    switch (^(fi31)) {
                                                                                                                                    case 997:
                                                                                                                                        if len(v51) > 0 {
                                                                                                                                            v51[0] = 1
                                                                                                                                        }
                                                                                                                                    case 32767:
                                                                                                                                        var v158 bool = ((fi32 - (^(v141))) > (fi31 * -392))
                                                                                                                                        _ = v158
                                                                                                                                    default:
                                                                                                                                        _ = len(v3)
                                                                                                                                    }
                                                                                                                                }
                                                                                                                                var v159 float64 = (0.0 - v14)
                                                                                                                                _ = v159
                                                                                                                            }
                                                                                                                            switch ("hello" + (("a" + v35) + ("hello" + v2))) {
                                                                                                                            case "\n\t":
                                                                                                                                v160 := v3[""]
                                                                                                                                _ = v160
                                                                                                                            case "foo bar":
                                                                                                                                var v161 bool = ((v127 % (((v127 & v127)) | 1)) > 18446744073709551615)
                                                                                                                                _ = v161
                                                                                                                            }
                                                                                                                            delete(v3, (("hello" + v35) + v35))
                                                                                                                            v162 := GenS22{F0: 247}
                                                                                                                            _ = v162
                                                                                                                        }
                                                                                                                    } else {
                                                                                                                        switch "hello" {
                                                                                                                        case "hello":
                                                                                                                            if v136 {
                                                                                                                                v15 += v15
                                                                                                                                v163 := GenS23{F0: v2}
                                                                                                                                v51 = append(v51, (fi32 - -9223372036854775808))
                                                                                                                                delete(v3, (v35 + v35))
                                                                                                                                var v164 float64 = (v14 + 0.0)
                                                                                                                                _ = v163
                                                                                                                                _ = v164
                                                                                                                            }
                                                                                                                        default:
                                                                                                                            if v136 {
                                                                                                                                v165 := GenS24{F0: "foo bar", F1: v127}
                                                                                                                                var v166 int64 = 764
                                                                                                                                _ = v165
                                                                                                                                _ = v166
                                                                                                                            }
                                                                                                                        }
                                                                                                                    }
                                                                                                                default:
                                                                                                                    if (1e10 == (0.1 / 1e10)) {
                                                                                                                        var v167 int16 = -1
                                                                                                                        if (!v136) {
                                                                                                                            if (-2.25 == 1.5) {
                                                                                                                                switch -1 {
                                                                                                                                case 128:
                                                                                                                                    switch (fi32 & (fi32 ^ v141)) {
                                                                                                                                    case 9223372036854775807:
                                                                                                                                        v168 := GenS11{F0: (v141 & (v141 / ((v141) | 1)))}
                                                                                                                                        _ = v168
                                                                                                                                    case 32767:
                                                                                                                                        var v169 uint32 = (v13 * (v13 ^ v13))
                                                                                                                                        _ = v169
                                                                                                                                    case 4294967295:
                                                                                                                                        switch (fi31 + (fi31 / ((fi32) | 1))) {
                                                                                                                                        case 2147483647:
                                                                                                                                            var v170 uint8 = (v1 % (((v1 & 150)) | 1))
                                                                                                                                            _ = v170
                                                                                                                                        case 1:
                                                                                                                                            var v171 bool = false
                                                                                                                                            _ = v171
                                                                                                                                        case 0:
                                                                                                                                            v172 := GenS25{F0: 2147483648, F1: 42}
                                                                                                                                            _ = v172
                                                                                                                                        default:
                                                                                                                                            if true {
                                                                                                                                                var v173 int16 = 0
                                                                                                                                                var v174 bool = (v136 || true)
                                                                                                                                                var v175 uint64 = 830
                                                                                                                                                _ = v173
                                                                                                                                                _ = v174
                                                                                                                                                _ = v175
                                                                                                                                            }
                                                                                                                                        }
                                                                                                                                    }
                                                                                                                                }
                                                                                                                            } else {
                                                                                                                                if v136 {
                                                                                                                                    if (1 == 42) {
                                                                                                                                        if len(v51) >= 2 {
                                                                                                                                            v51 = v51[1:len(v51)]
                                                                                                                                        }
                                                                                                                                        v15 += 65535
                                                                                                                                    }
                                                                                                                                    v15 -= v4
                                                                                                                                    var v176 int32 = -970
                                                                                                                                    v15 -= uint16(v127)
                                                                                                                                    if false {
                                                                                                                                        var v177 string = ("\n\t" + "a")
                                                                                                                                        var v178 uint8 = uint8(v13)
                                                                                                                                        _ = v177
                                                                                                                                        _ = v178
                                                                                                                                    }
                                                                                                                                    _ = v176
                                                                                                                                }
                                                                                                                                if v136 {
                                                                                                                                    v16 = v16
                                                                                                                                } else {
                                                                                                                                    v16 = v16
                                                                                                                                    v35 = v2
                                                                                                                                    switch v35 {
                                                                                                                                    case "a":
                                                                                                                                        var v179 int = v141
                                                                                                                                        _ = v179
                                                                                                                                    case "hello":
                                                                                                                                        v16 -= genFn0()
                                                                                                                                    }
                                                                                                                                }
                                                                                                                                v180 := GenS5{F0: (v16 & 30), F1: v4}
                                                                                                                                _ = v180
                                                                                                                            }
                                                                                                                            var v181 bool = v136
                                                                                                                            v182 := GenS26{F0: fi31}
                                                                                                                            _ = v181
                                                                                                                            _ = v182
                                                                                                                        } else {
                                                                                                                            switch "a" {
                                                                                                                            case "\n\t":
                                                                                                                                switch -985 {
                                                                                                                                case 2147483647:
                                                                                                                                    var v183 int64 = 1
                                                                                                                                    _ = v183
                                                                                                                                case 4294967295:
                                                                                                                                    var v184 string = "\n\t"
                                                                                                                                    _ = v184
                                                                                                                                case 811:
                                                                                                                                    v14 = 1e10
                                                                                                                                default:
                                                                                                                                    if v136 {
                                                                                                                                        var v185 float64 = v14
                                                                                                                                        var v186 int8 = v16
                                                                                                                                        v4 = v15
                                                                                                                                        _ = v185
                                                                                                                                        _ = v186
                                                                                                                                    } else {
                                                                                                                                        fmt.Println(v136)
                                                                                                                                        var v187 int8 = v16
                                                                                                                                        _ = v187
                                                                                                                                    }
                                                                                                                                    v167 = (v167 - (v167 & (v5 & v5)))
                                                                                                                                }
                                                                                                                            case "hello":
                                                                                                                                if (false || true) {
                                                                                                                                    v188 := GenS8{F0: 9223372036854775807}
                                                                                                                                    _ = v188
                                                                                                                                } else {
                                                                                                                                    v2 = "\n\t"
                                                                                                                                    var v189 uint64 = 2
                                                                                                                                    switch v2 {
                                                                                                                                    case "":
                                                                                                                                        if true {
                                                                                                                                            fmt.Println(v189, v167, v15)
                                                                                                                                        } else {
                                                                                                                                            v190 := GenS17{F0: false}
                                                                                                                                            if len(v51) > 0 {
                                                                                                                                                v51[0] = (-(v141))
                                                                                                                                            }
                                                                                                                                            v191 := GenS21{F0: 9223372036854775807, F1: uint(v4), F2: v136}
                                                                                                                                            _ = v190
                                                                                                                                            _ = v191
                                                                                                                                        }
                                                                                                                                    case "hello":
                                                                                                                                        v51 = append(v51, (fi31 - 42))
                                                                                                                                    default:
                                                                                                                                        if len(v51) >= 2 {
                                                                                                                                            v51 = v51[1:len(v51)]
                                                                                                                                        }
                                                                                                                                    }
                                                                                                                                    _ = v189
                                                                                                                                }
                                                                                                                            case "":
                                                                                                                                v141 -= (v141 % ((v141) | 1))
                                                                                                                            }
                                                                                                                        }
                                                                                                                        switch "foo bar" {
                                                                                                                        case "a":
                                                                                                                            var v192 bool = ((^(v16)) != (^(v16)))
                                                                                                                            _ = v192
                                                                                                                        }
                                                                                                                        _ = v167
                                                                                                                    }
                                                                                                                }
                                                                                                                _ = v141
                                                                                                            }
                                                                                                        }
                                                                                                        v193 := GenS27{F0: (v1 == v1), F1: (v1 < v1)}
                                                                                                        _ = v127
                                                                                                        _ = v136
                                                                                                        _ = v137
                                                                                                        _ = v193
                                                                                                    }
                                                                                                    v5 = v5
                                                                                                    if len(v51) >= 2 {
                                                                                                        v51 = v51[1:len(v51)]
                                                                                                    }
                                                                                                } else {
                                                                                                    if (9223372036854775807 != 1) {
                                                                                                        v5 = v5
                                                                                                        if (false || ((((0.0 * 3.14159) >= 1.5) || false) || true)) {
                                                                                                            var v194 bool = (-1 != 1)
                                                                                                            var v195 int32 = (v50 % ((v17) | 1))
                                                                                                            _ = v194
                                                                                                            _ = v195
                                                                                                        }
                                                                                                    } else {
                                                                                                        var v196 float64 = v14
                                                                                                        if (!false) {
                                                                                                            var v197 float32 = ((-2.25 * float32(v14)) + 1e10)
                                                                                                            _ = v197
                                                                                                        } else {
                                                                                                            v3[v35] = 32767
                                                                                                            _ = len(v3)
                                                                                                        }
                                                                                                        switch 55 {
                                                                                                        case 0:
                                                                                                            _ = len(v51)
                                                                                                        case 13:
                                                                                                            v13 -= v13
                                                                                                        }
                                                                                                        v13 = 2
                                                                                                        _ = v196
                                                                                                    }
                                                                                                    v198 := GenS9{F0: 32767, F1: false, F2: v1}
                                                                                                    _ = v198
                                                                                                }
                                                                                            }
                                                                                        default:
                                                                                            v199 := GenS23{F0: v2}
                                                                                            _ = v199
                                                                                        }
                                                                                    }
                                                                                case "foo bar":
                                                                                    v16 = (v16 + (-(v16)))
                                                                                }
                                                                            }
                                                                        }
                                                                        v2 = "foo bar"
                                                                    }
                                                                    if true {
                                                                        var v200 int = 1
                                                                        if (v13 == 4294967295) {
                                                                            fmt.Println(v35)
                                                                            if false {
                                                                                fmt.Println(fi31, v5)
                                                                                v201 := GenS4{F0: 1, F1: v35}
                                                                                fmt.Println(v5, v35, fi31)
                                                                                v1 += v1
                                                                                _ = v201
                                                                            }
                                                                        }
                                                                        _ = v200
                                                                    } else {
                                                                        v202 := GenS28{F0: 0}
                                                                        v203 := GenS25{F0: (v13 * v13), F1: -4}
                                                                        if true {
                                                                            switch fi31 {
                                                                            case 1:
                                                                                v51 = append(v51, fi32)
                                                                            default:
                                                                                v4 += 42
                                                                            }
                                                                            var v204 uint32 = v13
                                                                            _ = v204
                                                                        }
                                                                        fmt.Println(fi31, v14)
                                                                        var v205 int16 = v5
                                                                        fmt.Println(v14)
                                                                        var v206 float32 = 1.5
                                                                        v3[v2] = (fi32 - fi31)
                                                                        _ = v202
                                                                        _ = v203
                                                                        _ = v205
                                                                        _ = v206
                                                                    }
                                                                    _ = v94
                                                                }
                                                            default:
                                                                v13 += 2
                                                            }
                                                        case 128:
                                                            switch "\n\t" {
                                                            case "":
                                                                var v207 int32 = 429
                                                                _ = v207
                                                            case "a":
                                                                v50 -= -2147483648
                                                                var v208 float64 = 3.14159
                                                                _ = v208
                                                            default:
                                                                if ((!false) && (!(!true))) {
                                                                    v2 = v35
                                                                    var v209 uint16 = v15
                                                                    _ = v209
                                                                } else {
                                                                    var v210 int8 = genFn0()
                                                                    v13 += 2
                                                                    if true {
                                                                        v1 = (v1 >> ((711) & 63))
                                                                        v15 = 183
                                                                        delete(v3, v35)
                                                                    } else {
                                                                        if (-9223372036854775808 > int64(v5)) {
                                                                            v211 := GenS29{F0: 18446744073709551615, F1: -9223372036854775808, F2: 42}
                                                                            fmt.Println(v35, v4)
                                                                            _ = v211
                                                                        }
                                                                        var v212 string = v35
                                                                        v213 := GenS30{F0: -182, F1: v4}
                                                                        var v214 float32 = (3.14159 - (3.14159 - 1.5))
                                                                        _ = v212
                                                                        _ = v213
                                                                        _ = v214
                                                                    }
                                                                    _ = v210
                                                                }
                                                                v4 = v15
                                                                var v215 int8 = 1
                                                                v216 := v3[v35]
                                                                _ = v215
                                                                _ = v216
                                                            }
                                                        }
                                                        fmt.Println(v17, v4, fi31)
                                                    }
                                                    if true {
                                                        v51 = append(v51, 0)
                                                    } else {
                                                        var v217 uint32 = 0
                                                        v218 := GenS31{F0: 255}
                                                        _ = v217
                                                        _ = v218
                                                    }
                                                    v3[v35] = fi32
                                                } else {
                                                    v219 := GenS23{F0: ((v35 + ("foo bar" + v2)) + ("hello" + v35))}
                                                    v220 := GenS23{F0: (v2 + "\n\t")}
                                                    _ = v219
                                                    _ = v220
                                                }
                                                fmt.Println(v17, fi31, v17)
                                                v4 += (v4 << ((42) & 63))
                                            }
                                        case "\n\t":
                                            if ((v5 - 0) >= (v5 / ((32767) | 1))) {
                                                var v221 int = (fi31 << ((255) & 63))
                                                _ = v221
                                            }
                                        case "a":
                                            var v222 string = (((v35 + v35) + v2) + (v2 + "foo bar"))
                                            _ = v222
                                        default:
                                            v15 += (v4 << ((0) & 63))
                                        }
                                        v223 := v3[(((((v35 + v35) + ((v2 + "hello") + "hello")) + v2) + (((v2 + v35) + (v2 + (v35 + (("foo bar" + ((((v2 + v35) + (((("hello" + v35) + v2) + ("" + v2)) + ("" + v35))) + "") + ("hello" + v35))) + ((((v2 + v2) + (v35 + "a")) + v35) + "hello"))))) + v2)) + ("hello" + ("" + v35)))]
                                        _ = v223
                                        _ = v58
                                    }
                                    if false {
                                        var v224 float64 = v14
                                        var v225 float64 = (-2.25 / (v224 + v14))
                                        var v226 int32 = v17
                                        _ = v224
                                        _ = v225
                                        _ = v226
                                    }
                                    v17 += v17
                                } else {
                                    var v227 int64 = 1
                                    v1 = (v1 << ((2) & 63))
                                    if false {
                                        fmt.Println(v2, v15, v15)
                                    }
                                    _ = v227
                                }
                            }
                            v4 = v15
                            var v228 string = ((((("\n\t" + "") + v35) + v2) + v35) + v2)
                            var v229 uint32 = v13
                            _ = v228
                            _ = v229
                            _ = v50
                            _ = v51
                        }
                    }
                    _ = v35
                }
                v230 := GenS31{F0: 776}
                v14 -= v14
                switch "a" {
                case "":
                    v3[v2] = 32768
                    v16 += (v16 % ((1) | 1))
                    v15 = (v15 << ((42) & 63))
                case "hello":
                    if false {
                        v5 = (v5 ^ (v5 | 42))
                    }
                case "a":
                    if false {
                        fmt.Println(fi32, v2)
                        var v231 int16 = int16(v16)
                        v14 = v14
                        _ = v231
                    } else {
                        if (true || true) {
                            fmt.Println(fi31, v5, v4)
                        }
                        v2 = v2
                    }
                }
                delete(v3, v2)
                _ = fi32
                _ = v230
            }
            if (127 == v1) {
                var v232 bool = true
                for fi233 := 0; fi233 < 20; fi233++ {
                    v14 = -2.25
                    _ = fi233
                }
                var v234 uint16 = v4
                _ = v232
                _ = v234
            } else {
                fmt.Println(v17, v16, v1)
            }
            var v235 uint = 0
            _ = v235
        } else {
            v16 = 127
        }
        var v236 bool = false
        fmt.Println(v15)
        _ = fi31
        _ = v236
    }
    var v237 uint64 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v237:", v237)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v237
    _ = v3
    _ = v4
    _ = v5
}
