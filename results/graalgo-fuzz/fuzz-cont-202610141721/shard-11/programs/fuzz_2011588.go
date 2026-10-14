package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32; F2 string }
type GenS2 struct { F0 string; F1 uint; F2 uint8 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint; F1 int }
type GenS5 struct { F0 uint; F1 uint64; F2 int }
type GenS6 struct { F0 uint16; F1 int64 }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 bool; F1 uint64 }
type GenS9 struct { F0 uint64 }
type GenS10 struct { F0 int16; F1 uint16; F2 uint }
type GenS11 struct { F0 uint64 }
type GenS12 struct { F0 int64; F1 uint64; F2 int16 }
type GenS13 struct { F0 int32; F1 int8 }
type GenS14 struct { F0 uint8; F1 int8; F2 int64 }
type GenS15 struct { F0 uint16 }
type GenS16 struct { F0 uint32 }
type GenS17 struct { F0 uint32; F1 uint16; F2 int16 }
type GenS18 struct { F0 int64 }
type GenS19 struct { F0 int32; F1 uint16; F2 int32 }
func genFn0() string {
    v6 := GenS1{F0: 1, F1: 2, F2: "hello"}
    var v7 uint8 = 1
    var v8 uint8 = v7
    for fi9 := 0; fi9 < 15; fi9++ {
        v7 -= v8
        switch fi9 {
        case 128:
            var v10 float64 = ((0.0 * -2.25) * float64(float32(float64((-2.25 * -2.25)))))
            _ = v10
        default:
            v11 := []int{fi9}
            _ = v11
        }
        var v12 int8 = 1
        _ = fi9
        _ = v12
    }
    _ = v6
    _ = v7
    _ = v8
    return ""
}

func main() {
    var v1 uint = 491
    var v2 uint32 = 2147483647
    var v3 uint8 = 130
    var v4 bool = true
    v5 := make(map[string]string)
    switch (("\n\t" + "a") + "\n\t") {
    case "":
        for fi13 := 0; fi13 < 3; fi13++ {
            switch (-(fi13)) {
            case -869:
                _ = len(v5)
            }
            if v4 {
                var v14 uint8 = v3
                _ = v14
            } else {
                v1 = (v1 % ((18446744073709551615) | 1))
                for fi15 := 0; fi15 < 15; fi15++ {
                    v2 -= uint32(v1)
                    var v16 int16 = -43
                    if (136 <= uint16(v1)) {
                        var v17 int32 = -13
                        if (100 >= 0) {
                            if ((^(v2)) != 2147483648) {
                                v4 = true
                            }
                            v3 -= 128
                            fmt.Println(v2)
                        } else {
                            delete(v5, "\n\t")
                            fmt.Println(v2)
                            var v18 float64 = 3.14159
                            _ = v18
                        }
                        v4 = v4
                        _ = v17
                    }
                    v19 := GenS1{F0: -472, F1: uint32(v1), F2: "hello"}
                    _ = fi15
                    _ = v16
                    _ = v19
                }
                delete(v5, "")
            }
            _ = fi13
        }
    case "foo bar":
        var v20 float64 = -2.25
        _ = v20
    case "hello":
        if (-128 != 1) {
            for fi21 := 0; fi21 < 15; fi21++ {
                fmt.Println(v3, v4)
                switch 227 {
                case -1:
                    if (((255 > 1) && v4) || ("\n\t" != "")) {
                        v3 += (v3 - 128)
                        if (!(v3 >= (v3 | v3))) {
                            if (!v4) {
                                var v22 uint64 = 610
                                for fi23 := 0; fi23 < 18; fi23++ {
                                    v1 = v1
                                    _ = fi23
                                }
                                for fi24 := 0; fi24 < 8; fi24++ {
                                    var v25 uint16 = 577
                                    fmt.Println(fi24, fi21, v1)
                                    if (0.0 >= -2.25) {
                                        v26 := GenS2{F0: ("\n\t" + "hello"), F1: (^(v1)), F2: v3}
                                        var v27 int32 = 212
                                        switch fi24 {
                                        case 2147483647:
                                            _ = len(v5)
                                        case 1:
                                            v5[("" + "hello")] = "a"
                                            v22 = (^(v22))
                                            v28 := v5["foo bar"]
                                            _ = v28
                                        default:
                                            v27 -= 42
                                        }
                                        _ = v26
                                        _ = v27
                                    } else {
                                        v2 += 407
                                        switch ("a" + ("\n\t" + ("a" + "\n\t"))) {
                                        case "":
                                            if true {
                                                v29 := GenS1{F0: int64(v3), F1: (v2 | (^(v2))), F2: ("foo bar" + "hello")}
                                                if ((^(v3)) == (v3 & (v3 << ((v1) & 63)))) {
                                                    if false {
                                                        v4 = ((^(fi24)) <= fi21)
                                                        if (!((!(128 <= fi24)) || (("\n\t" + ("foo bar" + (("\n\t" + ("\n\t" + "foo bar")) + "hello"))) == ("\n\t" + "foo bar")))) {
                                                            v30 := []int{(-(fi21)), 2147483647, fi21, (-(fi24))}
                                                            v4 = v4
                                                            v31 := GenS2{F0: "\n\t", F1: v1, F2: v3}
                                                            _ = v30
                                                            _ = v31
                                                        }
                                                        fmt.Println(v22)
                                                    } else {
                                                        var v32 int64 = 0
                                                        _ = len(v5)
                                                        _ = v32
                                                    }
                                                    _ = len(v5)
                                                    v33 := []int{fi24, 956, fi24, (fi21 % (((fi21 >> ((uint(v22)) & 63))) | 1))}
                                                    _ = v33
                                                }
                                                var v34 uint64 = (v22 & v22)
                                                var v35 int16 = -488
                                                _ = v29
                                                _ = v34
                                                _ = v35
                                            }
                                        case "foo bar":
                                            switch int(v2) {
                                            case 128:
                                                var v36 float32 = ((0.0 - 0.0) + float32(0.1))
                                                _ = v36
                                            default:
                                                v37 := []int{fi24}
                                                _ = v37
                                            }
                                        case "a":
                                            var v38 int16 = 32767
                                            _ = v38
                                        default:
                                            v39 := GenS3{F0: v25}
                                            _ = v39
                                        }
                                    }
                                    _ = fi24
                                    _ = v25
                                }
                                _ = v22
                            } else {
                                switch (("hello" + "a") + "hello") {
                                case "hello":
                                    for fi40 := 0; fi40 < 5; fi40++ {
                                        if v4 {
                                            v41 := []int{2147483647, (^(fi21))}
                                            if (!(true || v4)) {
                                                if (!true) {
                                                    v42 := GenS2{F0: "foo bar", F1: v1, F2: (v3 ^ 255)}
                                                    _ = v42
                                                }
                                            } else {
                                                if len(v41) > 0 {
                                                    v41[0] = int(v3)
                                                }
                                                switch (fi21 * (fi21 / (((fi21 * fi40)) | 1))) {
                                                case -249:
                                                    v3 = (v3 | (v3 - (v3 & 2)))
                                                case -9223372036854775808:
                                                    v43 := v5["a"]
                                                    _ = v43
                                                case 32768:
                                                    v5["foo bar"] = ""
                                                default:
                                                    switch fi40 {
                                                    case -1:
                                                        var v44 uint8 = v3
                                                        fmt.Println(v44, fi40)
                                                        v45 := GenS3{F0: 2}
                                                        _ = len(v41)
                                                        _ = v44
                                                        _ = v45
                                                    }
                                                }
                                                v1 = v1
                                            }
                                            _ = v41
                                        }
                                        var v46 float32 = (-2.25 * (float32(((-2.25 / (((float64(1e10) / (0.1 / 1.0)) - (1.5 / 1.5)) / (1.5 / (((0.0 + (float64(-2.25) / 1.0)) - float64((-2.25 - -2.25))) / 3.14159)))) + ((3.14159 - 1.5) + 1.5))) * 0.0))
                                        v3 -= v3
                                        _ = fi40
                                        _ = v46
                                    }
                                default:
                                    v3 = (^(v3))
                                }
                                for fi47 := 0; fi47 < 4; fi47++ {
                                    v48 := GenS2{F0: "\n\t", F1: 1, F2: 2}
                                    if v4 {
                                        switch 128 {
                                        case 2147483647:
                                            v4 = (("a" == "foo bar") || v4)
                                        case 316:
                                            if (4294967295 != 489) {
                                                var v49 int16 = 999
                                                v1 = 65535
                                                _ = v49
                                            } else {
                                                var v50 int32 = -645
                                                switch "" {
                                                case "hello":
                                                    if v4 {
                                                        v2 -= (v2 % (((v2 << ((v1) & 63))) | 1))
                                                        v51 := []int{}
                                                        _ = v51
                                                    } else {
                                                        var v52 int16 = 42
                                                        if true {
                                                            var v53 float64 = (1e10 + 0.0)
                                                            _ = v53
                                                        } else {
                                                            if v4 {
                                                                v54 := GenS4{F0: 2, F1: -817}
                                                                v52 -= (-(v52))
                                                                var v55 uint = (v1 - (v1 ^ v1))
                                                                v56 := GenS4{F0: (v55 ^ (v1 / (((^(v55))) | 1))), F1: fi21}
                                                                _ = v54
                                                                _ = v55
                                                                _ = v56
                                                            } else {
                                                                if (!(!v4)) {
                                                                    if v4 {
                                                                        var v57 int32 = v50
                                                                        _ = v57
                                                                    } else {
                                                                        if (!v4) {
                                                                            v1 -= v1
                                                                            v52 = (-(v52))
                                                                            switch "hello" {
                                                                            case "\n\t":
                                                                                switch "" {
                                                                                case "\n\t":
                                                                                    v58 := v5[""]
                                                                                    _ = v58
                                                                                case "a":
                                                                                    fmt.Println(v2, v50)
                                                                                default:
                                                                                    if v4 {
                                                                                        if (true && false) {
                                                                                            switch "\n\t" {
                                                                                            case "a":
                                                                                                v59 := []int{fi21, (fi47 & -1), fi47, fi21}
                                                                                                _ = v59
                                                                                            default:
                                                                                                if v4 {
                                                                                                    v60 := GenS3{F0: 32767}
                                                                                                    v61 := []int{(fi47 * fi47), 1}
                                                                                                    fmt.Println(v3, fi21)
                                                                                                    _ = v60
                                                                                                    _ = v61
                                                                                                }
                                                                                            }
                                                                                            if (1 == 9223372036854775807) {
                                                                                                var v62 bool = (v2 == v2)
                                                                                                fmt.Println(fi21, fi47, v62)
                                                                                                _ = v62
                                                                                            } else {
                                                                                                if (!v4) {
                                                                                                    if v4 {
                                                                                                        var v63 int8 = 79
                                                                                                        v2 = v2
                                                                                                        _ = v63
                                                                                                    } else {
                                                                                                        var v64 int32 = (v50 / (((v50 ^ 42)) | 1))
                                                                                                        var v65 float64 = 1e10
                                                                                                        v66 := GenS5{F0: v1, F1: 18446744073709551615, F2: (fi21 << (((v1 * v1)) & 63))}
                                                                                                        _ = v64
                                                                                                        _ = v65
                                                                                                        _ = v66
                                                                                                    }
                                                                                                    v67 := []int{fi21, fi21, fi47, (fi47 << ((866) & 63))}
                                                                                                    var v68 int64 = 981
                                                                                                    var v69 int = (fi21 * (fi21 / (((fi21 ^ (fi47 + (fi21 & fi47)))) | 1)))
                                                                                                    _ = v67
                                                                                                    _ = v68
                                                                                                    _ = v69
                                                                                                }
                                                                                                if true {
                                                                                                    if ((!v4) && false) {
                                                                                                        switch "foo bar" {
                                                                                                        case "foo bar":
                                                                                                            var v70 uint = 4294967295
                                                                                                            _ = v70
                                                                                                        case "\n\t":
                                                                                                            v4 = v4
                                                                                                        case "":
                                                                                                            v5["foo bar"] = ""
                                                                                                        }
                                                                                                    }
                                                                                                    var v71 uint64 = 18446744073709551615
                                                                                                    switch "a" {
                                                                                                    case "hello":
                                                                                                        fmt.Println(fi21)
                                                                                                    case "\n\t":
                                                                                                        var v72 int = (fi21 % (((fi21 & int(v1))) | 1))
                                                                                                        _ = v72
                                                                                                    case "foo bar":
                                                                                                        v3 += 42
                                                                                                        v3 += v3
                                                                                                        if ((v4 || (!(0 > 61))) || true) {
                                                                                                            var v73 int = 9223372036854775807
                                                                                                            v3 = (v3 >> ((v1) & 63))
                                                                                                            v74 := GenS4{F0: 0, F1: 4294967295}
                                                                                                            _ = v73
                                                                                                            _ = v74
                                                                                                        } else {
                                                                                                            if v4 {
                                                                                                                fmt.Println(v4)
                                                                                                                fmt.Println(v71, v50, fi47)
                                                                                                                v75 := GenS5{F0: (v1 ^ v1), F1: (v71 & v71), F2: fi47}
                                                                                                                fmt.Println(v71, v50)
                                                                                                                _ = v75
                                                                                                            } else {
                                                                                                                v76 := GenS2{F0: (("hello" + "foo bar") + "foo bar"), F1: 42, F2: v3}
                                                                                                                v1 -= v1
                                                                                                                switch ("" + "foo bar") {
                                                                                                                case "hello":
                                                                                                                    if (9223372036854775807 == -9223372036854775808) {
                                                                                                                        if true {
                                                                                                                            switch "foo bar" {
                                                                                                                            case "foo bar":
                                                                                                                                var v77 int = 32767
                                                                                                                                _ = v77
                                                                                                                            case "":
                                                                                                                                v50 += 791
                                                                                                                            }
                                                                                                                            if (!true) {
                                                                                                                                if false {
                                                                                                                                    switch "a" {
                                                                                                                                    case "a":
                                                                                                                                        v1 -= 1
                                                                                                                                    case "\n\t":
                                                                                                                                        if false {
                                                                                                                                            v3 -= v3
                                                                                                                                            if false {
                                                                                                                                                v50 -= (v50 | (-(v50)))
                                                                                                                                                delete(v5, "a")
                                                                                                                                                var v78 bool = v4
                                                                                                                                                switch fi21 {
                                                                                                                                                case 128:
                                                                                                                                                    v78 = (v52 == int16(v1))
                                                                                                                                                }
                                                                                                                                                v78 = v4
                                                                                                                                                var v79 int = 9223372036854775807
                                                                                                                                                fmt.Println(v50, v1, v52)
                                                                                                                                                _ = v78
                                                                                                                                                _ = v79
                                                                                                                                            } else {
                                                                                                                                                var v80 float32 = (1.5 / 0.1)
                                                                                                                                                _ = v80
                                                                                                                                            }
                                                                                                                                            if v4 {
                                                                                                                                                if (!v4) {
                                                                                                                                                    var v81 float64 = 0.1
                                                                                                                                                    v82 := GenS5{F0: (^(v1)), F1: uint64(v1), F2: (^(fi47))}
                                                                                                                                                    var v83 int64 = -9223372036854775808
                                                                                                                                                    _ = v81
                                                                                                                                                    _ = v82
                                                                                                                                                    _ = v83
                                                                                                                                                }
                                                                                                                                                var v84 bool = v4
                                                                                                                                                v84 = false
                                                                                                                                                _ = v84
                                                                                                                                            } else {
                                                                                                                                                var v85 int64 = 0
                                                                                                                                                if (v50 <= v50) {
                                                                                                                                                    v4 = ((^(v85)) <= (v85 * int64(v3)))
                                                                                                                                                    var v86 bool = v4
                                                                                                                                                    fmt.Println(v2)
                                                                                                                                                    _ = v86
                                                                                                                                                }
                                                                                                                                                var v87 int64 = -451
                                                                                                                                                _ = v85
                                                                                                                                                _ = v87
                                                                                                                                            }
                                                                                                                                        } else {
                                                                                                                                            fmt.Println(v1, v3, v4)
                                                                                                                                            v3 -= 98
                                                                                                                                            v88 := GenS5{F0: (^(v1)), F1: (^(v71)), F2: fi47}
                                                                                                                                            _ = v88
                                                                                                                                        }
                                                                                                                                    default:
                                                                                                                                        var v89 int8 = 33
                                                                                                                                        _ = v89
                                                                                                                                    }
                                                                                                                                    v90 := GenS2{F0: "foo bar", F1: v1, F2: (v3 - (v3 << (((v1 + (v1 ^ (v1 & v1)))) & 63)))}
                                                                                                                                    var v91 uint8 = v3
                                                                                                                                    _ = v90
                                                                                                                                    _ = v91
                                                                                                                                } else {
                                                                                                                                    v92 := GenS1{F0: -9223372036854775808, F1: (v2 ^ (v2 | uint32(v71))), F2: ((("\n\t" + "\n\t") + "hello") + "hello")}
                                                                                                                                    if (v2 == (^(v2))) {
                                                                                                                                        switch (fi21 | fi21) {
                                                                                                                                        case 0:
                                                                                                                                            v93 := GenS1{F0: -9223372036854775808, F1: v2, F2: "hello"}
                                                                                                                                            _ = v93
                                                                                                                                        default:
                                                                                                                                            fmt.Println(v4, v2, v3)
                                                                                                                                        }
                                                                                                                                        v94 := GenS5{F0: v1, F1: v71, F2: fi21}
                                                                                                                                        switch "\n\t" {
                                                                                                                                        case "":
                                                                                                                                            v3 = v3
                                                                                                                                        }
                                                                                                                                        _ = v94
                                                                                                                                    }
                                                                                                                                    if v4 {
                                                                                                                                        v50 = v50
                                                                                                                                        var v95 int32 = -863
                                                                                                                                        if ("\n\t" == (("\n\t" + "") + "foo bar")) {
                                                                                                                                            v1 = v1
                                                                                                                                            var v96 int64 = 9223372036854775807
                                                                                                                                            v52 -= (^(v52))
                                                                                                                                            _ = v96
                                                                                                                                        }
                                                                                                                                        _ = v95
                                                                                                                                    }
                                                                                                                                    _ = v92
                                                                                                                                }
                                                                                                                            }
                                                                                                                            v97 := []int{(fi47 >> ((v1) & 63)), (^(fi47)), int(v2), fi21}
                                                                                                                            var v98 uint16 = 42
                                                                                                                            _ = len(v97)
                                                                                                                            v99 := GenS5{F0: 18446744073709551615, F1: (v71 << ((241) & 63)), F2: fi47}
                                                                                                                            v2 = v2
                                                                                                                            _ = v97
                                                                                                                            _ = v98
                                                                                                                            _ = v99
                                                                                                                        }
                                                                                                                    }
                                                                                                                    v2 = v2
                                                                                                                case "":
                                                                                                                    var v100 uint64 = (v71 - 4294967295)
                                                                                                                    _ = v100
                                                                                                                case "\n\t":
                                                                                                                    var v101 int = fi47
                                                                                                                    _ = v101
                                                                                                                default:
                                                                                                                    v102 := GenS1{F0: 516, F1: v2, F2: "foo bar"}
                                                                                                                    _ = v102
                                                                                                                }
                                                                                                                v4 = true
                                                                                                                v103 := GenS5{F0: 4294967295, F1: (^(v71)), F2: (fi47 - 1)}
                                                                                                                _ = v103
                                                                                                                _ = v76
                                                                                                            }
                                                                                                        }
                                                                                                    }
                                                                                                    _ = v71
                                                                                                } else {
                                                                                                    v4 = v4
                                                                                                }
                                                                                                v2 = uint32(v3)
                                                                                                switch 0 {
                                                                                                case -2147483648:
                                                                                                    var v104 int = fi21
                                                                                                    var v105 uint32 = (v2 | (v2 >> ((v1) & 63)))
                                                                                                    if (!(2 <= 4294967295)) {
                                                                                                        if v4 {
                                                                                                            var v106 int32 = (v50 << ((v1) & 63))
                                                                                                            v5["hello"] = "a"
                                                                                                            delete(v5, "hello")
                                                                                                            _ = v106
                                                                                                        } else {
                                                                                                            var v107 bool = v4
                                                                                                            _ = v107
                                                                                                        }
                                                                                                        var v108 uint16 = 753
                                                                                                        _ = v108
                                                                                                    } else {
                                                                                                        if true {
                                                                                                            var v109 string = ("hello" + ("" + "hello"))
                                                                                                            _ = v109
                                                                                                        } else {
                                                                                                            _ = len(v5)
                                                                                                            var v110 uint64 = 545
                                                                                                            if ((!((!false) && ((-(fi21)) != fi21))) && ((v3 - v3) >= 255)) {
                                                                                                                if false {
                                                                                                                    v50 -= -780
                                                                                                                    fmt.Println(v4, v52)
                                                                                                                    v2 = v2
                                                                                                                }
                                                                                                            }
                                                                                                            _ = v110
                                                                                                        }
                                                                                                        v1 = (^(v1))
                                                                                                        var v111 uint16 = 42
                                                                                                        _ = v111
                                                                                                    }
                                                                                                    if (-1 == 9223372036854775807) {
                                                                                                        v104 -= v104
                                                                                                        var v112 string = ("" + "\n\t")
                                                                                                        _ = v112
                                                                                                    }
                                                                                                    var v113 int = (fi21 + -652)
                                                                                                    _ = v104
                                                                                                    _ = v105
                                                                                                    _ = v113
                                                                                                case 9223372036854775807:
                                                                                                    v1 += (v1 << (((v1 | 42)) & 63))
                                                                                                }
                                                                                                if ((((((v3 >= (^(v3))) || (!(!(("a" + ("hello" + "")) != ("\n\t" + ("\n\t" + "\n\t")))))) && (-2147483648 <= (fi47 >> ((v1) & 63)))) || (!(v4 && false))) && (uint32(v1) != 924)) && v4) {
                                                                                                    fmt.Println(v52, fi47)
                                                                                                    if true {
                                                                                                        v3 -= 2
                                                                                                    } else {
                                                                                                        v114 := GenS1{F0: -505, F1: (v2 + (v2 - (v2 + (^(v2))))), F2: "foo bar"}
                                                                                                        v115 := GenS1{F0: 9223372036854775807, F1: (^(v2)), F2: "\n\t"}
                                                                                                        var v116 float32 = 3.14159
                                                                                                        _ = v114
                                                                                                        _ = v115
                                                                                                        _ = v116
                                                                                                    }
                                                                                                    v117 := []int{921, 238, 2147483647, fi21}
                                                                                                    v3 = (v3 | 128)
                                                                                                    var v118 uint = (^(v1))
                                                                                                    _ = v117
                                                                                                    _ = v118
                                                                                                }
                                                                                            }
                                                                                        } else {
                                                                                            if v4 {
                                                                                                var v119 uint32 = 0
                                                                                                _ = v119
                                                                                            } else {
                                                                                                fmt.Println(v4, v2, v4)
                                                                                            }
                                                                                        }
                                                                                        if v4 {
                                                                                            _ = len(v5)
                                                                                            if false {
                                                                                                var v120 int64 = 525
                                                                                                fmt.Println(v52)
                                                                                                v50 += (v50 | 42)
                                                                                                _ = v120
                                                                                            }
                                                                                            var v121 uint = v1
                                                                                            var v122 int32 = 976
                                                                                            if v4 {
                                                                                                if false {
                                                                                                    var v123 int64 = 412
                                                                                                    _ = v123
                                                                                                }
                                                                                                v121 = (v1 ^ v1)
                                                                                            }
                                                                                            v50 = v50
                                                                                            _ = v121
                                                                                            _ = v122
                                                                                        } else {
                                                                                            var v124 float32 = -2.25
                                                                                            _ = v124
                                                                                        }
                                                                                        v125 := []int{fi21, (fi21 ^ fi47)}
                                                                                        var v126 uint64 = 1
                                                                                        delete(v5, "a")
                                                                                        _ = v125
                                                                                        _ = v126
                                                                                    }
                                                                                }
                                                                            }
                                                                        }
                                                                    }
                                                                    v127 := GenS2{F0: (("a" + "hello") + ("hello" + "")), F1: (v1 >> (((v1 << ((v1) & 63))) & 63)), F2: v3}
                                                                    fmt.Println(v50, v1)
                                                                    _ = v127
                                                                }
                                                                v128 := []int{fi21, fi21, fi47, (fi47 ^ (fi47 << ((2) & 63)))}
                                                                _ = v128
                                                            }
                                                            v5[("hello" + "foo bar")] = "\n\t"
                                                        }
                                                        var v129 bool = v4
                                                        if v4 {
                                                            fmt.Println(v129)
                                                        }
                                                        _ = v129
                                                        _ = v52
                                                    }
                                                }
                                                if (62 == 2) {
                                                    v130 := []int{fi47, fi21, fi21, (fi21 << ((1) & 63))}
                                                    v50 += 738
                                                    v50 += v50
                                                    _ = v130
                                                } else {
                                                    var v131 uint = v1
                                                    if (!v4) {
                                                        fmt.Println(v1)
                                                        v50 += (^(v50))
                                                        v4 = true
                                                    } else {
                                                        v132 := GenS1{F0: -536, F1: v2, F2: "foo bar"}
                                                        v133 := []int{-991, fi47}
                                                        var v134 bool = (18446744073709551615 != 2)
                                                        _ = v132
                                                        _ = v133
                                                        _ = v134
                                                    }
                                                    if (1e10 == 1e10) {
                                                        var v135 uint64 = 838
                                                        v136 := []int{fi47, fi47, (fi21 + fi21), fi47}
                                                        _ = len(v5)
                                                        _ = v135
                                                        _ = v136
                                                    } else {
                                                        if v4 {
                                                            if v4 {
                                                                fmt.Println(v1)
                                                                var v137 int16 = 279
                                                                v138 := GenS5{F0: v1, F1: 356, F2: -9223372036854775808}
                                                                _ = v137
                                                                _ = v138
                                                            }
                                                        }
                                                        if v4 {
                                                            switch "foo bar" {
                                                            case "":
                                                                switch (fi21 / (((fi21 % (((fi21 + (fi47 << (((^(v131))) & 63)))) | 1))) | 1)) {
                                                                case 9223372036854775807:
                                                                    switch "foo bar" {
                                                                    case "":
                                                                        v139 := GenS6{F0: 0, F1: 42}
                                                                        _ = v139
                                                                    }
                                                                case -1:
                                                                    var v140 int = (^(fi21))
                                                                    _ = v140
                                                                default:
                                                                    _ = len(v5)
                                                                }
                                                            }
                                                            switch "\n\t" {
                                                            case "foo bar":
                                                                if (101 <= 619) {
                                                                    var v141 float32 = (3.14159 - 1.5)
                                                                    _ = v141
                                                                } else {
                                                                    if (!v4) {
                                                                        switch "hello" {
                                                                        case "hello":
                                                                            var v142 uint16 = 314
                                                                            if v4 {
                                                                                switch (^(fi47)) {
                                                                                case 2147483647:
                                                                                    v4 = (v4 && true)
                                                                                case 42:
                                                                                    v143 := GenS7{F0: (v3 % (((v3 % ((128) | 1))) | 1))}
                                                                                    _ = v143
                                                                                case 32768:
                                                                                    v2 -= v2
                                                                                }
                                                                                if (v4 || ("foo bar" < "\n\t")) {
                                                                                    var v144 string = ""
                                                                                    if true {
                                                                                        v142 += v142
                                                                                    }
                                                                                    var v145 float32 = (1e10 * -2.25)
                                                                                    _ = v144
                                                                                    _ = v145
                                                                                } else {
                                                                                    if false {
                                                                                        var v146 uint64 = 130
                                                                                        if true {
                                                                                            if (v2 > v2) {
                                                                                                v147 := GenS8{F0: v4, F1: v146}
                                                                                                _ = v147
                                                                                            }
                                                                                            v3 = (v3 | (v3 % (((^(v3))) | 1)))
                                                                                        }
                                                                                        _ = v146
                                                                                    } else {
                                                                                        if (!v4) {
                                                                                            fmt.Println(fi47, v4)
                                                                                        }
                                                                                        if (42 >= v1) {
                                                                                            switch "\n\t" {
                                                                                            case "a":
                                                                                                v148 := []int{fi21, fi21}
                                                                                                _ = v148
                                                                                            }
                                                                                            v149 := []int{(fi47 & fi47), fi47}
                                                                                            v4 = false
                                                                                            _ = v149
                                                                                        } else {
                                                                                            v4 = (!(v50 < (v50 / ((-175) | 1))))
                                                                                        }
                                                                                    }
                                                                                    if (-9223372036854775808 != 0) {
                                                                                        var v150 uint32 = 620
                                                                                        var v151 int16 = 1000
                                                                                        _ = v150
                                                                                        _ = v151
                                                                                    } else {
                                                                                        var v152 uint32 = v2
                                                                                        var v153 int64 = 9223372036854775807
                                                                                        _ = v152
                                                                                        _ = v153
                                                                                    }
                                                                                    var v154 bool = (!v4)
                                                                                    if v154 {
                                                                                        v155 := GenS3{F0: (v142 % ((1) | 1))}
                                                                                        _ = v155
                                                                                    }
                                                                                    if true {
                                                                                        if v154 {
                                                                                            var v156 uint32 = (v2 << ((v1) & 63))
                                                                                            _ = v156
                                                                                        } else {
                                                                                            if ((false || (v4 && (!(!((42 < 666) && ((1e10 == 1e10) || false)))))) && (v4 && v154)) {
                                                                                                v2 = v2
                                                                                            }
                                                                                        }
                                                                                        fmt.Println(v1, v142)
                                                                                    } else {
                                                                                        _ = len(v5)
                                                                                        var v157 uint16 = v142
                                                                                        _ = v157
                                                                                    }
                                                                                    _ = v154
                                                                                }
                                                                                switch ("" + (("foo bar" + ("\n\t" + "foo bar")) + "\n\t")) {
                                                                                case "a":
                                                                                    if (v4 && (((!(v4 || false)) && v4) && ((^(v50)) == int32(fi47)))) {
                                                                                        v50 = 1
                                                                                        v158 := GenS4{F0: 42, F1: 127}
                                                                                        _ = v158
                                                                                    }
                                                                                case "hello":
                                                                                    var v159 uint32 = (^(v2))
                                                                                    _ = v159
                                                                                case "\n\t":
                                                                                    fmt.Println(fi21, v4, v1)
                                                                                default:
                                                                                    var v160 int = (^(fi47))
                                                                                    _ = v160
                                                                                }
                                                                            } else {
                                                                                v161 := GenS2{F0: "", F1: (v131 * (v1 | (v131 / (((^(v131))) | 1)))), F2: 0}
                                                                                _ = v161
                                                                            }
                                                                            _ = v142
                                                                        case "foo bar":
                                                                            fmt.Println(v3)
                                                                        case "\n\t":
                                                                            var v162 uint64 = 4294967295
                                                                            if (46 != v3) {
                                                                                v163 := GenS7{F0: 127}
                                                                                var v164 int8 = -10
                                                                                var v165 string = ""
                                                                                _ = v163
                                                                                _ = v164
                                                                                _ = v165
                                                                            } else {
                                                                                v131 = 0
                                                                                switch (("foo bar" + "") + "hello") {
                                                                                case "\n\t":
                                                                                    if v4 {
                                                                                        var v166 uint32 = (v2 ^ v2)
                                                                                        _ = v166
                                                                                    } else {
                                                                                        fmt.Println(fi47, v131, v50)
                                                                                        v167 := GenS9{F0: v162}
                                                                                        var v168 uint32 = (v2 >> ((uint(v2)) & 63))
                                                                                        _ = v167
                                                                                        _ = v168
                                                                                    }
                                                                                case "":
                                                                                    v169 := GenS8{F0: (v3 == (v3 % (((v3 - (v3 << (((v1 >> (((v131 - v1)) & 63))) & 63)))) | 1))), F1: uint64(v2)}
                                                                                    _ = v169
                                                                                case "a":
                                                                                    if v4 {
                                                                                        v162 = 537
                                                                                        if v4 {
                                                                                            v3 -= (v3 | (v3 + (v3 % ((45) | 1))))
                                                                                            switch "hello" {
                                                                                            case "":
                                                                                                if ((^(v1)) != 0) {
                                                                                                    if (!v4) {
                                                                                                        if v4 {
                                                                                                            switch ("\n\t" + "a") {
                                                                                                            case "\n\t":
                                                                                                                v5["\n\t"] = "foo bar"
                                                                                                            }
                                                                                                            v131 += 159
                                                                                                            var v170 float64 = 1e10
                                                                                                            _ = v170
                                                                                                        }
                                                                                                        if ((925 == 0) || v4) {
                                                                                                            v2 -= 906
                                                                                                            var v171 uint16 = 546
                                                                                                            _ = v171
                                                                                                        }
                                                                                                        v4 = v4
                                                                                                    }
                                                                                                    v2 += 1
                                                                                                    v172 := []int{fi21, fi47, fi47, (fi47 | 192)}
                                                                                                    fmt.Println(v4, v3, v2)
                                                                                                    v1 = (v131 << ((125) & 63))
                                                                                                    switch "\n\t" {
                                                                                                    case "foo bar":
                                                                                                        v1 = v131
                                                                                                    case "a":
                                                                                                        v173 := GenS10{F0: 0, F1: 65535, F2: 784}
                                                                                                        _ = v173
                                                                                                    }
                                                                                                    _ = v172
                                                                                                } else {
                                                                                                    if (!((605 < uint(v162)) || ((uint32(v3) == v2) || ((true && (!v4)) || v4)))) {
                                                                                                        if ((v1 + (v1 / ((v1) | 1))) > uint(v3)) {
                                                                                                            fmt.Println(v1)
                                                                                                            v174 := GenS11{F0: (v162 >> (((v1 + (v131 / ((v131) | 1)))) & 63))}
                                                                                                            if ((4294967295 == (^(v162))) || (!false)) {
                                                                                                                v175 := []int{fi47, fi47, 1, (fi21 & fi47)}
                                                                                                                switch "hello" {
                                                                                                                case "":
                                                                                                                    if v4 {
                                                                                                                        v176 := GenS1{F0: 42, F1: 561, F2: "foo bar"}
                                                                                                                        _ = v176
                                                                                                                    }
                                                                                                                }
                                                                                                                v177 := GenS12{F0: 9223372036854775807, F1: 202, F2: -343}
                                                                                                                _ = v175
                                                                                                                _ = v177
                                                                                                            } else {
                                                                                                                if (-1 != 127) {
                                                                                                                    var v178 bool = ((^(v162)) > 1)
                                                                                                                    if (false && v4) {
                                                                                                                        v179 := []int{fi21}
                                                                                                                        _ = v179
                                                                                                                    }
                                                                                                                    v178 = ((3.14159 - 1e10) > 0.1)
                                                                                                                    if ((v4 || false) || v4) {
                                                                                                                        v180 := GenS9{F0: v162}
                                                                                                                        _ = v180
                                                                                                                    } else {
                                                                                                                        switch ("\n\t" + "foo bar") {
                                                                                                                        case "foo bar":
                                                                                                                            v1 -= (v131 / ((uint(v3)) | 1))
                                                                                                                        default:
                                                                                                                            if true {
                                                                                                                                v181 := []int{(fi21 & (fi47 % (((fi47 >> (((v131 % ((v1) | 1))) & 63))) | 1))), (^(fi21)), fi21, (fi21 * -2147483648)}
                                                                                                                                delete(v5, ("\n\t" + ""))
                                                                                                                                if v178 {
                                                                                                                                    v2 -= (v2 + (v2 >> (((v131 / ((v131) | 1))) & 63)))
                                                                                                                                    var v182 uint64 = v162
                                                                                                                                    var v183 uint = (v131 - v1)
                                                                                                                                    v184 := GenS13{F0: (v50 / (((^(v50))) | 1)), F1: 53}
                                                                                                                                    _ = v182
                                                                                                                                    _ = v183
                                                                                                                                    _ = v184
                                                                                                                                } else {
                                                                                                                                    switch fi47 {
                                                                                                                                    case -9223372036854775808:
                                                                                                                                        if ("" <= ("a" + ("foo bar" + "hello"))) {
                                                                                                                                            var v185 uint16 = 144
                                                                                                                                            fmt.Println(v3)
                                                                                                                                            _ = v185
                                                                                                                                        }
                                                                                                                                    case 32768:
                                                                                                                                        var v186 float64 = float64(float32(((((-2.25 * 1.5) * float64(1.5)) + 0.0) * (-2.25 * 1.5))))
                                                                                                                                        _ = v186
                                                                                                                                    }
                                                                                                                                }
                                                                                                                                _ = v181
                                                                                                                            }
                                                                                                                        }
                                                                                                                        v178 = v4
                                                                                                                    }
                                                                                                                    if (!v178) {
                                                                                                                        v187 := GenS1{F0: 42, F1: (^(v2)), F2: ""}
                                                                                                                        if ("foo bar" <= ("hello" + "")) {
                                                                                                                            v5["foo bar"] = "hello"
                                                                                                                            fmt.Println(fi21, fi21, v4)
                                                                                                                        } else {
                                                                                                                            v188 := GenS7{F0: v3}
                                                                                                                            var v189 uint16 = 826
                                                                                                                            _ = v188
                                                                                                                            _ = v189
                                                                                                                        }
                                                                                                                        v190 := GenS14{F0: v3, F1: 42, F2: -9223372036854775808}
                                                                                                                        _ = v187
                                                                                                                        _ = v190
                                                                                                                    }
                                                                                                                    v2 = 42
                                                                                                                    var v191 int32 = v50
                                                                                                                    _ = v178
                                                                                                                    _ = v191
                                                                                                                }
                                                                                                                switch "a" {
                                                                                                                case "a":
                                                                                                                    if ((!true) && ((-31 > -128) && v4)) {
                                                                                                                        v192 := GenS7{F0: v3}
                                                                                                                        _ = v192
                                                                                                                    } else {
                                                                                                                        v131 -= 65535
                                                                                                                    }
                                                                                                                case "foo bar":
                                                                                                                    if (!true) {
                                                                                                                        v193 := []int{fi21, -9223372036854775808, (fi47 % (((-(fi21))) | 1)), fi21}
                                                                                                                        _ = v193
                                                                                                                    } else {
                                                                                                                        v50 = (-(v50))
                                                                                                                        fmt.Println(fi21, v162)
                                                                                                                        var v194 float32 = (-2.25 / 1.5)
                                                                                                                        v3 += v3
                                                                                                                        if (1000 < -188) {
                                                                                                                            var v195 int32 = -1
                                                                                                                            v196 := v5[("foo bar" + (("" + "\n\t") + "foo bar"))]
                                                                                                                            v1 = (^(v1))
                                                                                                                            _ = v195
                                                                                                                            _ = v196
                                                                                                                        }
                                                                                                                        switch ("a" + ("hello" + "hello")) {
                                                                                                                        case "a":
                                                                                                                            v162 = (v162 | (v162 - 42))
                                                                                                                        case "hello":
                                                                                                                            var v197 bool = (!false)
                                                                                                                            _ = v197
                                                                                                                        case "":
                                                                                                                            v198 := []int{(fi47 >> ((v1) & 63))}
                                                                                                                            _ = v198
                                                                                                                        default:
                                                                                                                            v199 := []int{int(v3), (-(fi47)), fi47, 1}
                                                                                                                            _ = v199
                                                                                                                        }
                                                                                                                        v194 = v194
                                                                                                                        var v200 int = 903
                                                                                                                        _ = v194
                                                                                                                        _ = v200
                                                                                                                    }
                                                                                                                default:
                                                                                                                    v50 -= -1
                                                                                                                }
                                                                                                            }
                                                                                                            _ = v174
                                                                                                        }
                                                                                                        if (v162 <= (v162 >> ((v1) & 63))) {
                                                                                                            v201 := []int{-2147483648, int(v50)}
                                                                                                            var v202 float32 = 0.1
                                                                                                            _ = v201
                                                                                                            _ = v202
                                                                                                        }
                                                                                                        v5[(("\n\t" + ("hello" + "")) + ("a" + "\n\t"))] = "a"
                                                                                                    }
                                                                                                }
                                                                                            case "foo bar":
                                                                                                v203 := []int{(fi47 ^ (fi47 << (((v1 >> ((1) & 63))) & 63))), (-(fi21)), fi47, (fi21 & (fi47 + (^(fi21))))}
                                                                                                _ = v203
                                                                                            case "\n\t":
                                                                                                if (!(v4 && (0 > 0))) {
                                                                                                    delete(v5, "a")
                                                                                                    if v4 {
                                                                                                        v204 := GenS10{F0: 1, F1: 2, F2: (v1 % (((^(v1))) | 1))}
                                                                                                        _ = v204
                                                                                                    }
                                                                                                }
                                                                                            default:
                                                                                                v4 = (!(!true))
                                                                                            }
                                                                                        } else {
                                                                                            v131 -= (v131 << ((v131) & 63))
                                                                                            var v205 string = "\n\t"
                                                                                            _ = v205
                                                                                        }
                                                                                    } else {
                                                                                        v3 -= (v3 | (v3 / ((v3) | 1)))
                                                                                    }
                                                                                default:
                                                                                    switch fi21 {
                                                                                    case 1:
                                                                                        v206 := []int{fi21}
                                                                                        _ = v206
                                                                                    case -812:
                                                                                        var v207 int32 = (v50 ^ v50)
                                                                                        _ = v207
                                                                                    default:
                                                                                        if (((float32((float64((0.0 - (3.14159 + float32(((float64(float32(0.0)) / 1.0) - (1.5 * ((0.0 * 0.0) - -2.25))))))) + -2.25)) - ((1.5 * (3.14159 + 1e10)) - float32(1.5))) / 1.5) >= 3.14159) {
                                                                                            fmt.Println(v2, v1)
                                                                                        } else {
                                                                                            v5["a"] = ("foo bar" + "foo bar")
                                                                                            if v4 {
                                                                                                var v208 bool = (1 > -64)
                                                                                                var v209 bool = false
                                                                                                var v210 uint = v131
                                                                                                _ = v208
                                                                                                _ = v209
                                                                                                _ = v210
                                                                                            } else {
                                                                                                v211 := GenS15{F0: 897}
                                                                                                if (v4 && v4) {
                                                                                                    if (!((^(v50)) >= v50)) {
                                                                                                        var v212 uint16 = 65535
                                                                                                        _ = v212
                                                                                                    } else {
                                                                                                        v4 = true
                                                                                                    }
                                                                                                    fmt.Println(v3)
                                                                                                }
                                                                                                v162 += v162
                                                                                                _ = v211
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                }
                                                                            }
                                                                            v213 := GenS15{F0: 485}
                                                                            _ = v162
                                                                            _ = v213
                                                                        }
                                                                        v50 = v50
                                                                        fmt.Println(v1)
                                                                    } else {
                                                                        v1 += v131
                                                                    }
                                                                    var v214 uint = 255
                                                                    _ = v214
                                                                }
                                                                _ = len(v5)
                                                                var v215 int64 = 961
                                                                fmt.Println(fi21)
                                                                _ = v215
                                                            default:
                                                                switch ("\n\t" + "a") {
                                                                case "foo bar":
                                                                    var v216 int8 = -95
                                                                    _ = v216
                                                                case "a":
                                                                    v5[("foo bar" + "foo bar")] = "a"
                                                                case "hello":
                                                                    var v217 uint64 = 990
                                                                    _ = v217
                                                                default:
                                                                    if (("a" + ("" + "")) >= "\n\t") {
                                                                        v218 := GenS1{F0: 85, F1: (v2 + v2), F2: "a"}
                                                                        _ = v218
                                                                    } else {
                                                                        if (910 < int16(v3)) {
                                                                            fmt.Println(fi21, fi21)
                                                                            fmt.Println(fi47)
                                                                        }
                                                                        v219 := GenS1{F0: 731, F1: (v2 << ((65535) & 63)), F2: "hello"}
                                                                        v220 := v5["\n\t"]
                                                                        switch "\n\t" {
                                                                        case "a":
                                                                            v50 = (v50 << ((0) & 63))
                                                                        case "foo bar":
                                                                            v221 := []int{(fi47 - fi21)}
                                                                            var v222 float64 = (0.1 / 0.1)
                                                                            if len(v221) >= 2 {
                                                                                v221 = v221[1:len(v221)]
                                                                            }
                                                                            var v223 bool = v4
                                                                            var v224 int64 = -776
                                                                            _ = v221
                                                                            _ = v222
                                                                            _ = v223
                                                                            _ = v224
                                                                        default:
                                                                            var v225 uint32 = v2
                                                                            _ = v225
                                                                        }
                                                                        _ = v219
                                                                        _ = v220
                                                                    }
                                                                }
                                                            }
                                                        }
                                                        v226 := []int{}
                                                        _ = v226
                                                    }
                                                    v131 = v131
                                                    _ = v131
                                                }
                                                _ = v50
                                            }
                                        }
                                        v2 = v2
                                        v227 := []int{(fi47 / ((fi21) | 1)), 32768}
                                        _ = v227
                                    }
                                    _ = len(v5)
                                    _ = fi47
                                    _ = v48
                                }
                                v4 = (0.0 == 1e10)
                            }
                        }
                        fmt.Println(v2, v2, fi21)
                    } else {
                        switch -773 {
                        case 781:
                            v228 := GenS16{F0: v2}
                            _ = v228
                        default:
                            var v229 uint64 = 615
                            _ = v229
                        }
                    }
                case -747:
                    var v230 bool = (v4 || v4)
                    _ = v230
                default:
                    switch "foo bar" {
                    case "foo bar":
                        switch "\n\t" {
                        case "":
                            for fi231 := 0; fi231 < 11; fi231++ {
                                var v232 float64 = -2.25
                                _ = fi231
                                _ = v232
                            }
                        case "foo bar":
                            switch ("" + "a") {
                            case "foo bar":
                                var v233 float64 = 0.1
                                _ = v233
                            case "a":
                                v234 := GenS17{F0: (v2 << ((4294967295) & 63)), F1: 0, F2: 42}
                                _ = v234
                            default:
                                var v235 string = "\n\t"
                                _ = v235
                            }
                        case "hello":
                            var v236 uint32 = (^(v2))
                            _ = v236
                        default:
                            var v237 int = 42
                            _ = v237
                        }
                    case "a":
                        fmt.Println(v3)
                    default:
                        v3 += 127
                    }
                }
                _ = fi21
            }
            v1 = v1
            v1 = v1
        } else {
            switch "foo bar" {
            case "foo bar":
                for fi238 := 0; fi238 < 15; fi238++ {
                    fmt.Println(v3)
                    var v239 uint64 = 1
                    _ = fi238
                    _ = v239
                }
            case "\n\t":
                var v240 uint64 = 1
                _ = v240
            case "":
                var v241 uint64 = 924
                _ = v241
            default:
                fmt.Println(v2, v3)
            }
        }
    default:
        var v242 int32 = 64
        _ = v242
    }
    if v4 {
        v4 = (!(v2 > 2147483648))
    } else {
        if (!(-1 != 66)) {
            var v243 int = 366
            _ = v243
        }
        v244 := GenS13{F0: -551, F1: 127}
        _ = v244
    }
    v245 := GenS18{F0: 9223372036854775807}
    for fi246 := 0; fi246 < 12; fi246++ {
        v1 = 255
        if true {
            var v247 uint16 = 0
            fmt.Println(v3)
            _ = len(v5)
            _ = v247
        }
        _ = fi246
    }
    v248 := GenS17{F0: (v2 / (((v2 * v2)) | 1)), F1: 32767, F2: -84}
    fmt.Println(v2)
    v3 -= (v3 + v3)
    var v249 uint16 = uint16(v3)
    v250 := GenS19{F0: 977, F1: v249, F2: 0}
    var v251 float32 = (-2.25 - 1e10)
    for fi252 := 0; fi252 < 8; fi252++ {
        v253 := []int{(fi252 / (((fi252 * 32767)) | 1))}
        var v254 int64 = -1
        _ = fi252
        _ = v253
        _ = v254
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v249:", v249)
    fmt.Println("v251:", v251)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v245
    _ = v248
    _ = v249
    _ = v250
    _ = v251
    _ = v3
    _ = v4
    _ = v5
}
