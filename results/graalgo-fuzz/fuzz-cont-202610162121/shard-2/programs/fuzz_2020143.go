package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint16; F1 int16 }
type GenS4 struct { F0 uint16; F1 int16; F2 int64 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 uint32; F1 bool }
type GenS7 struct { F0 int64; F1 int }
type GenS8 struct { F0 uint16; F1 bool; F2 int }
type GenS9 struct { F0 string }
type GenS10 struct { F0 int16; F1 uint64; F2 uint32 }
type GenS11 struct { F0 uint16 }
type GenS12 struct { F0 bool; F1 uint }
type GenS13 struct { F0 bool }
func genFn0() uint32 {
    for fi6 := 0; fi6 < 6; fi6++ {
        v7 := []int{(-(fi6)), (fi6 & (fi6 / ((1) | 1)))}
        _ = fi6
        _ = v7
    }
    var v8 int64 = -761
    _ = v8
    return 0
}
func genFn1() int32 {
    for fi9 := 0; fi9 < 14; fi9++ {
        switch ("a" + "") {
        case "foo bar":
            if (true && false) {
                switch fi9 {
                case 148:
                    for fi10 := 0; fi10 < 6; fi10++ {
                        var v11 int8 = 9
                        _ = fi10
                        _ = v11
                    }
                case -1:
                    for fi12 := 0; fi12 < 9; fi12++ {
                        var v13 int16 = 1000
                        switch "\n\t" {
                        case "foo bar":
                            fmt.Println(v13)
                        case "hello":
                            v13 = v13
                        case "a":
                            var v14 int = (fi9 - (fi9 % ((-530) | 1)))
                            _ = v14
                        }
                        fmt.Println(fi12)
                        var v15 uint8 = 42
                        _ = fi12
                        _ = v13
                        _ = v15
                    }
                case -9223372036854775808:
                    for fi16 := 0; fi16 < 20; fi16++ {
                        var v17 string = ("foo bar" + "\n\t")
                        _ = fi16
                        _ = v17
                    }
                }
            } else {
                var v18 uint = 42
                for fi19 := 0; fi19 < 14; fi19++ {
                    var v20 int8 = 42
                    var v21 uint = v18
                    var v22 string = ("foo bar" + ("foo bar" + "a"))
                    _ = fi19
                    _ = v20
                    _ = v21
                    _ = v22
                }
                _ = v18
            }
        case "a":
            switch "a" {
            case "\n\t":
                var v23 bool = true
                _ = v23
            }
        case "hello":
            var v24 string = (("" + ("hello" + (("\n\t" + "a") + "foo bar"))) + "a")
            _ = v24
        default:
            v25 := GenS1{F0: fi9, F1: "\n\t"}
            _ = v25
        }
        for fi26 := 0; fi26 < 1; fi26++ {
            if (0.0 != 1.5) {
                if ((fi9 << ((42) & 63)) <= (fi9 >> ((564) & 63))) {
                    if true {
                        if true {
                            var v27 uint = 255
                            v27 -= v27
                            v28 := GenS1{F0: 32767, F1: ""}
                            _ = v27
                            _ = v28
                        }
                        if true {
                            var v29 string = ("hello" + "\n\t")
                            _ = v29
                        }
                        var v30 uint8 = 128
                        _ = v30
                    }
                    if false {
                        v31 := []int{fi9, (fi26 ^ 1), fi9}
                        _ = v31
                    } else {
                        v32 := GenS2{F0: 42}
                        _ = v32
                    }
                    var v33 int64 = -1
                    _ = v33
                }
            }
            if (475 != 0) {
                var v34 bool = (!(42 < 958))
                fmt.Println(v34, v34)
                _ = v34
            } else {
                switch fi9 {
                case 1:
                    v35 := []int{fi9, 42, -360}
                    _ = v35
                case 14:
                    var v36 uint32 = 988
                    _ = v36
                case 4294967295:
                    var v37 uint64 = 554
                    _ = v37
                default:
                    v38 := GenS3{F0: 32767, F1: 1}
                    _ = v38
                }
                var v39 int16 = 754
                _ = v39
            }
            _ = fi26
        }
        _ = fi9
    }
    v40 := []int{}
    switch ("hello" + ("hello" + "")) {
    case "a":
        if true {
            for fi41 := 0; fi41 < 12; fi41++ {
                fmt.Println(fi41)
                _ = fi41
            }
            for fi42 := 0; fi42 < 12; fi42++ {
                v40 = append(v40, (fi42 << ((1) & 63)))
                var v43 uint8 = 0
                if false {
                    v43 -= (v43 + v43)
                    for fi44 := 0; fi44 < 3; fi44++ {
                        if (genFn0() >= 1) {
                            if (1e10 > (float32(1.5) + 1e10)) {
                                var v45 uint = 2
                                _ = v45
                            }
                            v43 -= 6
                            if len(v40) > 0 {
                                v40[0] = (fi44 & -2147483648)
                            }
                        } else {
                            fmt.Println(v43)
                        }
                        _ = fi44
                    }
                } else {
                    for fi46 := 0; fi46 < 14; fi46++ {
                        switch "" {
                        case "a":
                            var v47 bool = (false && (2 < 820))
                            _ = v47
                        case "\n\t":
                            var v48 float32 = 3.14159
                            v49 := GenS4{F0: 32767, F1: 758, F2: -1}
                            _ = v48
                            _ = v49
                        case "hello":
                            fmt.Println(fi42)
                        }
                        switch fi46 {
                        case 0:
                            v43 = v43
                        case 32767:
                            var v50 float32 = float32(1e10)
                            _ = v50
                        default:
                            if len(v40) >= 2 {
                                v40 = v40[1:len(v40)]
                            }
                        }
                        if true {
                            if ((float32((-2.25 + (1.5 * 1e10))) >= float32((((1.5 - 0.1) / 1.5) - -2.25))) && (2 == (v43 - v43))) {
                                if (!(0 > 1)) {
                                    var v51 uint8 = 1
                                    v52 := GenS5{F0: 9223372036854775807}
                                    v53 := GenS5{F0: -668}
                                    _ = v51
                                    _ = v52
                                    _ = v53
                                }
                                v54 := GenS4{F0: 1, F1: 1000, F2: 9223372036854775807}
                                var v55 uint64 = 988
                                _ = v54
                                _ = v55
                            } else {
                                var v56 string = "foo bar"
                                _ = v56
                            }
                            fmt.Println(fi46, fi46, v43)
                            fmt.Println(fi42)
                        } else {
                            var v57 int = (fi42 ^ fi46)
                            _ = len(v40)
                            var v58 string = "foo bar"
                            if ((("\n\t" + ((v58 + v58) + "hello")) + "hello") <= (v58 + v58)) {
                                v59 := GenS1{F0: (v57 - 4294967295), F1: v58}
                                _ = v59
                            }
                            fmt.Println(v43)
                            _ = v57
                            _ = v58
                        }
                        if true {
                            if false {
                                switch fi42 {
                                case 9223372036854775807:
                                    fmt.Println(v43)
                                }
                            } else {
                                var v60 string = "\n\t"
                                _ = v60
                            }
                        }
                        _ = fi46
                    }
                }
                _ = fi42
                _ = v43
            }
            if (42 <= 255) {
                if len(v40) >= 2 {
                    v40 = v40[1:len(v40)]
                }
                var v61 uint = 1
                _ = v61
            }
        }
    case "foo bar":
        _ = len(v40)
    }
    _ = v40
    return -1
}

func main() {
    v1 := []string{("" + ("foo bar" + "")), "", "a", "foo bar"}
    v2 := []uint32{}
    var v3 int = 2147483647
    v4 := make(map[uint32]string)
    var v5 int = v3
    fmt.Println(v3, v3)
    if false {
        switch v5 {
        case 9223372036854775807:
            var v62 uint16 = 1
            _ = v62
        }
    } else {
        if ((true && true) || (1 > 581)) {
            if false {
                if true {
                    switch (v5 / (((-(v3))) | 1)) {
                    case 9223372036854775807:
                        var v63 uint16 = 32767
                        _ = v63
                    default:
                        if ((v5 & -294) > (v3 >> ((18446744073709551615) & 63))) {
                            fmt.Println(v3)
                            for fi64 := 0; fi64 < 2; fi64++ {
                                if false {
                                    _ = len(v4)
                                }
                                _ = fi64
                            }
                            v4[genFn0()] = ("\n\t" + "a")
                        } else {
                            v5 -= -124
                        }
                    }
                    v5 += v3
                }
                if (!(false || false)) {
                    v65 := GenS2{F0: 346}
                    switch "a" {
                    case "":
                        var v66 int8 = 75
                        _ = v66
                    default:
                        fmt.Println(v3, v5)
                    }
                    _ = v65
                }
            } else {
                if (!(false && false)) {
                    fmt.Println(v5)
                    var v67 uint16 = 255
                    _ = v67
                }
            }
            for fi68 := 0; fi68 < 8; fi68++ {
                if (817 < genFn0()) {
                    v5 += fi68
                    v69 := GenS6{F0: 998, F1: (!(1 == 2))}
                    if (((1.5 / 0.1) - 0.1) <= float32(((0.0 + 1e10) * 1e10))) {
                        var v70 uint8 = 42
                        fmt.Println(v70, v5)
                        _ = v70
                    }
                    _ = v69
                } else {
                    v71 := GenS7{F0: 33, F1: 1}
                    v5 = fi68
                    _ = v71
                }
                v5 = (v3 | 0)
                v5 += -9223372036854775808
                _ = fi68
            }
        }
        for fi72 := 0; fi72 < 17; fi72++ {
            for fi73 := 0; fi73 < 3; fi73++ {
                v5 -= (fi73 - v3)
                var v74 uint64 = 1
                _ = len(v4)
                _ = fi73
                _ = v74
            }
            _ = fi72
        }
        v4[0] = ("" + "foo bar")
    }
    fmt.Println(v3)
    v75 := GenS8{F0: 81, F1: (!("foo bar" >= "foo bar")), F2: -711}
    var v76 uint16 = 940
    if (((v5 * (v5 | v5)) >= v3) || true) {
        for fi77 := 0; fi77 < 6; fi77++ {
            switch "\n\t" {
            case "a":
                switch (v5 >> ((255) & 63)) {
                case -2147483648:
                    for fi78 := 0; fi78 < 18; fi78++ {
                        v3 = v3
                        v79 := v4[genFn0()]
                        v80 := v4[genFn0()]
                        if false {
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            fmt.Println(v79, v3)
                            if false {
                                var v81 string = "\n\t"
                                _ = v81
                            } else {
                                v79 = ""
                                v2 = append(v2, 0)
                                if true {
                                    var v82 uint8 = 255
                                    var v83 uint64 = 709
                                    if true {
                                        if ((-32768 != -1) && true) {
                                            _ = len(v4)
                                        }
                                        var v84 uint8 = (v82 | (v82 | v82))
                                        _ = v84
                                    } else {
                                        var v85 float64 = (1.5 - 3.14159)
                                        _ = v85
                                    }
                                    _ = v82
                                    _ = v83
                                } else {
                                    v5 += (v3 >> ((499) & 63))
                                    v76 = v76
                                }
                            }
                        }
                        _ = fi78
                        _ = v79
                        _ = v80
                    }
                case 1:
                    for fi86 := 0; fi86 < 10; fi86++ {
                        var v87 bool = (true || false)
                        v3 += (^(fi77))
                        v3 = v3
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        fmt.Println(v3, fi77)
                        var v88 string = ("" + ("\n\t" + "hello"))
                        if (0 > 100) {
                            var v89 uint8 = 128
                            v3 = (-(fi77))
                            _ = v89
                        } else {
                            v90 := GenS9{F0: v88}
                            v5 = 32768
                            _ = v90
                        }
                        _ = fi86
                        _ = v87
                        _ = v88
                    }
                case 128:
                    switch (v3 * (v5 & (v5 * (v5 % ((v5) | 1))))) {
                    case -839:
                        for fi91 := 0; fi91 < 16; fi91++ {
                            v92 := GenS5{F0: -1}
                            v76 = (v76 >> ((715) & 63))
                            var v93 uint8 = 42
                            _ = fi91
                            _ = v92
                            _ = v93
                        }
                    case 2147483647:
                        var v94 int64 = 9223372036854775807
                        _ = v94
                    default:
                        v5 -= 128
                    }
                default:
                    v4[4294967295] = "a"
                }
            case "":
                var v95 uint32 = genFn0()
                _ = v95
            default:
                for fi96 := 0; fi96 < 8; fi96++ {
                    v97 := GenS10{F0: -32768, F1: 1, F2: 2147483648}
                    _ = fi96
                    _ = v97
                }
            }
            switch 232 {
            case 1:
                v5 = -2147483648
            case 127:
                v98 := GenS11{F0: v76}
                _ = v98
            }
            _ = fi77
        }
        var v99 int = -231
        _ = v99
    } else {
        for fi100 := 0; fi100 < 11; fi100++ {
            if (!(0.0 >= 1.5)) {
                if (18446744073709551615 >= 4294967295) {
                    if ((true || false) || (false && (true && (v76 <= (v76 >> ((0) & 63)))))) {
                        delete(v4, 0)
                    } else {
                        for fi101 := 0; fi101 < 13; fi101++ {
                            v3 -= (fi100 * fi101)
                            var v102 int16 = -1
                            _ = fi101
                            _ = v102
                        }
                        if (1.5 < 1.5) {
                            switch 143 {
                            case 2147483647:
                                delete(v4, 2147483647)
                            case -146:
                                var v103 uint8 = 1
                                _ = v103
                            case -932:
                                v5 = fi100
                            default:
                                v104 := GenS11{F0: (^(v76))}
                                _ = v104
                            }
                            v105 := GenS1{F0: (v5 % ((v3) | 1)), F1: "foo bar"}
                            var v106 string = "a"
                            _ = v105
                            _ = v106
                        } else {
                            var v107 int16 = 1
                            _ = v107
                        }
                    }
                    if (!(-1 != int32(v76))) {
                        switch -570 {
                        case 9223372036854775807:
                            switch fi100 {
                            case -2147483648:
                                v108 := GenS4{F0: v76, F1: 581, F2: 0}
                                _ = v108
                            }
                        case -425:
                            _ = len(v2)
                        }
                    }
                }
            } else {
                fmt.Println(v3, v3)
                v5 -= 32767
                v76 += 255
            }
            _ = fi100
        }
        v109 := GenS12{F0: false, F1: 4294967295}
        _ = v109
    }
    switch (v5 << ((1) & 63)) {
    case 304:
        switch "" {
        case "":
            for fi110 := 0; fi110 < 3; fi110++ {
                for fi111 := 0; fi111 < 3; fi111++ {
                    _ = len(v4)
                    _ = len(v2)
                    v76 -= v76
                    _ = fi111
                }
                for fi112 := 0; fi112 < 2; fi112++ {
                    var v113 uint32 = 0
                    if false {
                        if false {
                            if (!(1 > 2147483647)) {
                                v3 = (fi112 / ((4294967295) | 1))
                            } else {
                                if len(v2) > 0 {
                                    v2[0] = (^(v113))
                                }
                                v114 := GenS13{F0: (-360 > 203)}
                                _ = v114
                            }
                            v4[genFn0()] = ""
                            switch (fi112 * -9223372036854775808) {
                            case 170:
                                v1 = append(v1, "a")
                            default:
                                _ = len(v4)
                            }
                        }
                    }
                    delete(v4, (v113 * v113))
                    _ = fi112
                    _ = v113
                }
                _ = fi110
            }
        }
    case 2147483647:
        _ = len(v1)
    }
    switch "hello" {
    case "hello":
        var v115 int64 = -1
        _ = v115
    case "\n\t":
        switch (v3 | v5) {
        case 1:
            switch v5 {
            case 753:
                var v116 uint32 = 0
                _ = v116
            case 1:
                v3 = 42
            default:
                v117 := v4[1]
                _ = v117
            }
        case -1:
            var v118 float64 = 0.1
            _ = v118
        case -9223372036854775808:
            v5 = 128
        default:
            v3 += (v5 & (^(v3)))
        }
    default:
        var v119 int32 = 1
        _ = v119
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v76:", v76)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v75
    _ = v76
}
