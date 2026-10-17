package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint8; F1 uint64 }
type GenS3 struct { F0 uint16; F1 int32 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int }
type GenS6 struct { F0 uint64; F1 uint }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 uint32; F1 int }
type GenS9 struct { F0 int32; F1 uint32; F2 int32 }
type GenS10 struct { F0 uint8; F1 uint32; F2 uint32 }
type GenS11 struct { F0 string; F1 bool }
type GenS12 struct { F0 int }
type GenS13 struct { F0 int; F1 uint; F2 uint64 }
type GenS14 struct { F0 int8; F1 uint32 }
func genFn0() string {
    if (false || (1 >= 2147483647)) {
        var v7 int16 = 414
        for fi8 := 0; fi8 < 12; fi8++ {
            v9 := GenS1{F0: v7}
            switch "a" {
            case "foo bar":
                v10 := GenS1{F0: -32768}
                _ = v10
            case "hello":
                var v11 uint8 = 136
                _ = v11
            default:
                var v12 float64 = 0.1
                _ = v12
            }
            _ = fi8
            _ = v9
        }
        _ = v7
    } else {
        var v13 string = "\n\t"
        switch 9223372036854775807 {
        case 42:
            switch v13 {
            case "hello":
                if false {
                    var v14 float32 = 3.14159
                    if (!true) {
                        var v15 float32 = v14
                        _ = v15
                    } else {
                        var v16 uint = 42
                        _ = v16
                    }
                    var v17 int32 = 0
                    _ = v14
                    _ = v17
                }
            case "\n\t":
                v18 := GenS2{F0: 0, F1: 1}
                _ = v18
            case "foo bar":
                v13 = (v13 + (v13 + ("a" + v13)))
            default:
                if (255 > 4294967295) {
                    var v19 int32 = 959
                    var v20 int32 = (v19 | v19)
                    _ = v19
                    _ = v20
                }
            }
        case -2147483648:
            fmt.Println(v13)
        case -1:
            if false {
                if false {
                    if false {
                        v21 := GenS3{F0: 1, F1: 197}
                        _ = v21
                    } else {
                        v22 := GenS2{F0: 244, F1: 18446744073709551615}
                        _ = v22
                    }
                    v23 := []int{42}
                    switch 9223372036854775807 {
                    case 127:
                        fmt.Println(v13)
                    case 193:
                        v13 = v13
                    case 2147483647:
                        for fi24 := 0; fi24 < 18; fi24++ {
                            v13 = v13
                            _ = fi24
                        }
                    default:
                        if (0 >= 42) {
                            if len(v23) >= 2 {
                                v23 = v23[1:len(v23)]
                            }
                        }
                    }
                    _ = v23
                } else {
                    v25 := []int{}
                    _ = v25
                }
                v26 := GenS4{F0: false}
                _ = v26
            }
        default:
            if (!true) {
                if true {
                    for fi27 := 0; fi27 < 4; fi27++ {
                        var v28 int8 = 0
                        v28 += v28
                        if (!true) {
                            v28 += 78
                            switch "" {
                            case "":
                                v29 := []int{fi27, 0, fi27, (fi27 / ((128) | 1))}
                                _ = v29
                            default:
                                v30 := GenS3{F0: 255, F1: -2147483648}
                                _ = v30
                            }
                            var v31 uint32 = 42
                            _ = v31
                        } else {
                            if false {
                                v13 = (v13 + ("\n\t" + v13))
                                for fi32 := 0; fi32 < 6; fi32++ {
                                    v13 = "hello"
                                    switch v13 {
                                    case "foo bar":
                                        if false {
                                            var v33 int32 = -808
                                            _ = v33
                                        } else {
                                            v13 = "foo bar"
                                        }
                                    case "":
                                        fmt.Println(v13, v13, v13)
                                    }
                                    v34 := []int{}
                                    _ = fi32
                                    _ = v34
                                }
                                v35 := []int{(fi27 & fi27)}
                                _ = v35
                            } else {
                                v36 := GenS5{F0: -1}
                                for fi37 := 0; fi37 < 6; fi37++ {
                                    v28 = v28
                                    _ = fi37
                                }
                                _ = v36
                            }
                            fmt.Println(v28)
                            if (true || (32767 != -399)) {
                                v28 += v28
                            } else {
                                for fi38 := 0; fi38 < 13; fi38++ {
                                    v28 = (-(v28))
                                    _ = fi38
                                }
                            }
                        }
                        _ = fi27
                        _ = v28
                    }
                }
            } else {
                fmt.Println(v13)
                v13 = ((("foo bar" + "a") + ("a" + ((v13 + v13) + ("a" + v13)))) + v13)
                v13 = v13
            }
        }
        for fi39 := 0; fi39 < 5; fi39++ {
            v40 := []int{fi39, (-(fi39)), fi39}
            var v41 uint = 65535
            _ = fi39
            _ = v40
            _ = v41
        }
        _ = v13
    }
    var v42 float32 = 1e10
    v42 = v42
    v43 := []int{4294967295, -9223372036854775808, 32767, 42}
    _ = v42
    _ = v43
    return ("hello" + "hello")
}
func genFn1() int16 {
    var v44 bool = (3.14159 > 0.0)
    v44 = false
    _ = v44
    return -32768
}

func main() {
    v1 := make(map[string]string)
    var v2 float32 = 1.5
    v3 := make(map[string]bool)
    var v4 int64 = 42
    var v5 int8 = -1
    v6 := make(map[uint32]string)
    for fi45 := 0; fi45 < 4; fi45++ {
        var v46 string = "foo bar"
        _ = fi45
        _ = v46
    }
    v5 += -45
    v47 := []int{32767, 127, -193}
    for fi48 := 0; fi48 < 3; fi48++ {
        var v49 int16 = genFn1()
        v4 += -529
        switch ("" + "a") {
        case "":
            var v50 int32 = 1
            _ = v50
        case "\n\t":
            if (!true) {
                switch (fi48 << ((65535) & 63)) {
                case 153:
                    if ((!false) || (!(!(42 == v5)))) {
                        fmt.Println(v4)
                        v51 := GenS4{F0: (!((false && false) || true))}
                        v5 -= (v5 << ((599) & 63))
                        _ = v51
                    }
                case -989:
                    v47 = append(v47, (fi48 % (((^(fi48))) | 1)))
                case 1:
                    for fi52 := 0; fi52 < 3; fi52++ {
                        v1["hello"] = (("a" + "hello") + ("foo bar" + ""))
                        var v53 float64 = -2.25
                        _ = fi52
                        _ = v53
                    }
                default:
                    v54 := GenS6{F0: 0, F1: 1}
                    _ = v54
                }
                v55 := GenS7{F0: (v5 ^ (v5 & 96))}
                var v56 float32 = 0.1
                _ = v55
                _ = v56
            } else {
                v2 -= -2.25
                if (!true) {
                    v2 += -2.25
                    if (false || true) {
                        v49 = v49
                        fmt.Println(v5, fi48)
                    } else {
                        if len(v47) >= 2 {
                            v47 = v47[1:len(v47)]
                        }
                        if (v49 >= 32767) {
                            if false {
                                for fi57 := 0; fi57 < 10; fi57++ {
                                    v49 = (-(v49))
                                    if (0 == 1) {
                                        fmt.Println(fi48, fi57, v5)
                                        if ((v2 * (((float32(((1.5 * (-2.25 * float64(v2))) * float64((((v2 * v2) + (v2 + (v2 + v2))) + v2)))) + (-2.25 * v2)) - v2) / v2)) == (((v2 + 0.0) + 3.14159) - (v2 * v2))) {
                                            switch "foo bar" {
                                            case "foo bar":
                                                if len(v47) > 0 {
                                                    v47[0] = (^(fi48))
                                                }
                                            case "":
                                                v2 += -2.25
                                            case "a":
                                                if len(v47) >= 2 {
                                                    v47 = v47[1:len(v47)]
                                                }
                                            }
                                        }
                                        var v58 float32 = v2
                                        _ = v58
                                    }
                                    _ = fi57
                                }
                            }
                        } else {
                            v4 -= 0
                        }
                    }
                } else {
                    for fi59 := 0; fi59 < 3; fi59++ {
                        if false {
                            if false {
                                fmt.Println(v2, fi59)
                                v60 := GenS4{F0: (v5 < 42)}
                                v61 := GenS6{F0: 1, F1: 447}
                                _ = v60
                                _ = v61
                            } else {
                                switch "\n\t" {
                                case "":
                                    v1[("a" + "\n\t")] = ("foo bar" + ("foo bar" + ""))
                                case "\n\t":
                                    if (-2147483648 == -95) {
                                        fmt.Println(fi59)
                                        v62 := GenS5{F0: fi59}
                                        switch (fi48 - fi59) {
                                        case 127:
                                            var v63 int8 = (v5 ^ (v5 | (v5 & 36)))
                                            _ = v63
                                        default:
                                            var v64 int32 = -925
                                            _ = v64
                                        }
                                        _ = v62
                                    } else {
                                        switch "hello" {
                                        case "a":
                                            switch (^(fi59)) {
                                            case 438:
                                                _ = len(v47)
                                                switch fi48 {
                                                case 4294967295:
                                                    switch (fi59 / ((fi48) | 1)) {
                                                    case -9223372036854775808:
                                                        var v65 int = 32768
                                                        _ = v65
                                                    case 32768:
                                                        v47 = append(v47, -671)
                                                    default:
                                                        if true {
                                                            v49 = v49
                                                            v49 = v49
                                                            if len(v47) >= 2 {
                                                                v47 = v47[1:len(v47)]
                                                            }
                                                            if false {
                                                                var v66 int8 = (v5 / ((v5) | 1))
                                                                _ = v66
                                                            }
                                                        } else {
                                                            if (1 <= 127) {
                                                                if len(v47) > 0 {
                                                                    v47[0] = -391
                                                                }
                                                            } else {
                                                                v5 = (v5 << ((2) & 63))
                                                            }
                                                            var v67 uint8 = 164
                                                            if (uint16(v67) > 255) {
                                                                fmt.Println(v4, v2)
                                                            } else {
                                                                var v68 int16 = v49
                                                                if true {
                                                                    fmt.Println(v67)
                                                                }
                                                                var v69 bool = ((("a" + ("hello" + ("hello" + ("\n\t" + "a")))) != "\n\t") && true)
                                                                if v69 {
                                                                    v49 -= int16(v4)
                                                                    var v70 int64 = (v4 & v4)
                                                                    var v71 string = ("foo bar" + "foo bar")
                                                                    _ = v70
                                                                    _ = v71
                                                                }
                                                                if (v69 || (18446744073709551615 >= 607)) {
                                                                    v72 := GenS2{F0: 18, F1: 18446744073709551615}
                                                                    v5 += 33
                                                                    var v73 uint32 = 65535
                                                                    _ = v72
                                                                    _ = v73
                                                                } else {
                                                                    switch fi48 {
                                                                    case 42:
                                                                        switch "" {
                                                                        case "":
                                                                            if len(v47) >= 2 {
                                                                                v47 = v47[1:len(v47)]
                                                                            }
                                                                        case "a":
                                                                            if len(v47) > 0 {
                                                                                v47[0] = (fi48 ^ (fi48 & fi59))
                                                                            }
                                                                        case "foo bar":
                                                                            fmt.Println(v69)
                                                                        }
                                                                    case 4294967295:
                                                                        var v74 bool = true
                                                                        _ = v74
                                                                    case 128:
                                                                        v2 += (v2 * 1e10)
                                                                    }
                                                                    var v75 int8 = -74
                                                                    if (1 != 568) {
                                                                        v76 := GenS2{F0: v67, F1: 1}
                                                                        var v77 int32 = -1
                                                                        _ = v76
                                                                        _ = v77
                                                                    }
                                                                    var v78 uint32 = 42
                                                                    _ = v75
                                                                    _ = v78
                                                                }
                                                                _ = v68
                                                                _ = v69
                                                            }
                                                            _ = v67
                                                        }
                                                    }
                                                    var v79 string = (("\n\t" + "") + "hello")
                                                    if (false && (1 >= 42)) {
                                                        v80 := GenS8{F0: 0, F1: (fi48 - fi59)}
                                                        switch 1 {
                                                        case -424:
                                                            if (false || true) {
                                                                if ((fi59 >> ((0) & 63)) != 32767) {
                                                                    v4 = -9223372036854775808
                                                                    var v81 uint16 = 0
                                                                    var v82 uint8 = 128
                                                                    _ = v81
                                                                    _ = v82
                                                                } else {
                                                                    v83 := GenS9{F0: 2147483647, F1: 2147483648, F2: -2147483648}
                                                                    _ = v83
                                                                }
                                                                var v84 uint8 = 127
                                                                v79 = (v79 + (v79 + v79))
                                                                v85 := GenS6{F0: 2, F1: 4294967295}
                                                                _ = v84
                                                                _ = v85
                                                            } else {
                                                                v86 := GenS10{F0: 179, F1: 4294967295, F2: 973}
                                                                if false {
                                                                    if false {
                                                                        if len(v47) >= 2 {
                                                                            v47 = v47[1:len(v47)]
                                                                        }
                                                                    }
                                                                    _ = len(v47)
                                                                    switch (^(fi59)) {
                                                                    case 845:
                                                                        if true {
                                                                            var v87 bool = false
                                                                            v49 = (v49 ^ (v49 << ((4294967295) & 63)))
                                                                            _ = v87
                                                                        }
                                                                    case 9223372036854775807:
                                                                        var v88 int8 = (v5 / ((-2) | 1))
                                                                        _ = v88
                                                                    case 901:
                                                                        switch v79 {
                                                                        case "a":
                                                                            v5 += (v5 + -107)
                                                                        case "":
                                                                            var v89 int = -322
                                                                            _ = v89
                                                                        default:
                                                                            if len(v47) >= 2 {
                                                                                v47 = v47[1:len(v47)]
                                                                            }
                                                                        }
                                                                    default:
                                                                        switch fi59 {
                                                                        case -9223372036854775808:
                                                                            fmt.Println(fi59)
                                                                        case 386:
                                                                            var v90 int16 = v49
                                                                            _ = v90
                                                                        }
                                                                    }
                                                                }
                                                                _ = v86
                                                            }
                                                        }
                                                        var v91 int32 = -141
                                                        _ = v80
                                                        _ = v91
                                                    } else {
                                                        v2 += -2.25
                                                    }
                                                    var v92 uint16 = 2
                                                    _ = v79
                                                    _ = v92
                                                case 0:
                                                    fmt.Println(v2, v4, v49)
                                                case 22:
                                                    switch "foo bar" {
                                                    case "a":
                                                        v49 += 0
                                                    }
                                                default:
                                                    var v93 uint32 = 2147483648
                                                    _ = v93
                                                }
                                                fmt.Println(v4, v2)
                                            case -1:
                                                var v94 float64 = ((((0.0 - 3.14159) / (3.14159 - 0.0)) * -2.25) - (-2.25 - 3.14159))
                                                _ = v94
                                            case -55:
                                                v95 := GenS11{F0: ("hello" + ("a" + "a")), F1: ((!((fi48 / ((0) | 1)) > fi48)) && (1 != 1))}
                                                _ = v95
                                            default:
                                                fmt.Println(fi48, v4, v4)
                                            }
                                        case "":
                                            v4 = (v4 - -395)
                                        case "hello":
                                            var v96 int32 = 0
                                            switch (fi59 / (((fi59 ^ 42)) | 1)) {
                                            case 9223372036854775807:
                                                v97 := GenS5{F0: -487}
                                                _ = v97
                                            case 128:
                                                v98 := GenS7{F0: v5}
                                                _ = v98
                                            case -886:
                                                fmt.Println(v2, v96)
                                            }
                                            if false {
                                                switch (fi59 * (fi48 * int(v4))) {
                                                case 0:
                                                    if (v96 < 42) {
                                                        v5 += 0
                                                        v99 := GenS1{F0: int16(v4)}
                                                        _ = v99
                                                    }
                                                case -35:
                                                    fmt.Println(fi48, v96, fi59)
                                                }
                                            } else {
                                                fmt.Println(v2, fi59)
                                            }
                                            _ = v96
                                        }
                                        v49 = (v49 % (((v49 >> ((1) & 63))) | 1))
                                        var v100 int8 = -128
                                        _ = v100
                                    }
                                case "foo bar":
                                    switch ("hello" + "hello") {
                                    case "hello":
                                        v5 = 127
                                    case "\n\t":
                                        v101 := GenS12{F0: (fi59 + (fi59 ^ -635))}
                                        _ = v101
                                    default:
                                        if true {
                                            _ = len(v6)
                                        } else {
                                            v5 -= 1
                                            var v102 int32 = 42
                                            v102 = v102
                                            _ = v102
                                        }
                                    }
                                }
                                _ = len(v1)
                            }
                            v47 = append(v47, (fi48 ^ fi48))
                        } else {
                            switch "a" {
                            case "":
                                var v103 int64 = v4
                                _ = v103
                            case "hello":
                                if (2147483648 <= 425) {
                                    if len(v47) > 0 {
                                        v47[0] = -593
                                    }
                                    var v104 int16 = genFn1()
                                    _ = v104
                                }
                            case "foo bar":
                                v105 := GenS13{F0: fi48, F1: 444, F2: 716}
                                _ = len(v47)
                                _ = v105
                            default:
                                if true {
                                    var v106 uint = 65535
                                    switch "foo bar" {
                                    case "":
                                        v107 := GenS4{F0: ((1e10 - 0.0) < 1e10)}
                                        _ = v107
                                    }
                                    _ = v106
                                }
                            }
                            var v108 uint64 = 42
                            fmt.Println(v2, fi59)
                            _ = v108
                        }
                        v49 -= (v49 - (v49 / (((-(v49))) | 1)))
                        fmt.Println(v49, v5, v49)
                        _ = fi59
                    }
                }
                var v109 int8 = v5
                _ = v109
            }
        }
        _ = fi48
        _ = v49
    }
    v2 = v2
    switch 9223372036854775807 {
    case 127:
        v110 := GenS2{F0: 115, F1: 399}
        _ = v110
    default:
        v4 -= v4
    }
    fmt.Println(v2, v4, v5)
    fmt.Println(v2, v4)
    v111 := GenS14{F0: v5, F1: 65535}
    if false {
        var v112 int64 = (v4 * -347)
        _ = v112
    }
    var v113 int64 = v4
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v113:", v113)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v47:", len(v47), v47)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v111
    _ = v113
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v5
    _ = v6
}
