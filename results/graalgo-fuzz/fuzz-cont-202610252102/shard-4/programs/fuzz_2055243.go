package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string; F2 int64 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint32; F1 bool; F2 int32 }
type GenS5 struct { F0 uint; F1 uint32; F2 int8 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 int64; F1 bool; F2 string }
type GenS9 struct { F0 int32 }
type GenS10 struct { F0 int32 }
type GenS11 struct { F0 int64; F1 uint64 }
type GenS12 struct { F0 int16 }
type GenS13 struct { F0 uint64 }
type GenS14 struct { F0 uint32 }
type GenS15 struct { F0 uint8; F1 int64; F2 int }
type GenS16 struct { F0 bool; F1 int8; F2 int64 }
type GenS17 struct { F0 int32; F1 uint32; F2 bool }
func genFn0(p0 uint32, p1 int8) bool {
    switch (("foo bar" + "hello") + "\n\t") {
    case "hello":
        if true {
            v4 := []int{}
            _ = v4
        } else {
            v5 := GenS1{F0: 1, F1: "", F2: 9223372036854775807}
            _ = v5
        }
    case "a":
        if false {
            v6 := []int{4294967295, 127, 9223372036854775807}
            _ = v6
        }
    case "foo bar":
        p0 = p0
    default:
        for fi7 := 0; fi7 < 1; fi7++ {
            var v8 int8 = (-(p1))
            v8 = v8
            _ = fi7
            _ = v8
        }
    }
    var v9 uint = 18446744073709551615
    switch "a" {
    case "\n\t":
        var v10 uint32 = (p0 - (p0 - p0))
        _ = v10
    case "foo bar":
        if (p1 != (p1 + (-(p1)))) {
            var v11 uint64 = 939
            _ = v11
        } else {
            fmt.Println(p1, p0)
            v12 := []int{32767, 539, -37, 4294967295}
            v9 -= (v9 << (((v9 - (v9 & v9))) & 63))
            _ = v12
        }
    case "a":
        for fi13 := 0; fi13 < 11; fi13++ {
            p1 = (p1 / ((0) | 1))
            if ("a" < "foo bar") {
                if (true || ((p0 % ((621) | 1)) >= p0)) {
                    if (false || (true || (!true))) {
                        v14 := GenS2{F0: p1}
                        v15 := GenS1{F0: 0, F1: "\n\t", F2: -9223372036854775808}
                        var v16 uint16 = 0
                        _ = v14
                        _ = v15
                        _ = v16
                    } else {
                        v17 := GenS3{F0: int(v9)}
                        for fi18 := 0; fi18 < 9; fi18++ {
                            var v19 float32 = 3.14159
                            var v20 float32 = v19
                            var v21 string = "a"
                            _ = fi18
                            _ = v19
                            _ = v20
                            _ = v21
                        }
                        for fi22 := 0; fi22 < 13; fi22++ {
                            if true {
                                var v23 uint16 = 0
                                v9 = v9
                                _ = v23
                            } else {
                                switch "\n\t" {
                                case "foo bar":
                                    p0 += (p0 % ((p0) | 1))
                                case "\n\t":
                                    p0 = 65535
                                default:
                                    p1 -= -10
                                }
                                v24 := []int{}
                                if (uint64(p0) == 898) {
                                    fmt.Println(fi13, p0)
                                } else {
                                    var v25 int64 = 9223372036854775807
                                    fmt.Println(p1, fi22, p0)
                                    v26 := GenS3{F0: fi13}
                                    _ = v25
                                    _ = v26
                                }
                                _ = v24
                            }
                            switch (fi13 & -2147483648) {
                            case 32767:
                                var v27 int = 32767
                                _ = v27
                            default:
                                v28 := []int{}
                                _ = v28
                            }
                            var v29 string = ("foo bar" + ("hello" + ("\n\t" + "")))
                            _ = fi22
                            _ = v29
                        }
                        _ = v17
                    }
                }
                v9 += v9
            }
            _ = fi13
        }
    default:
        if (255 < 38) {
            v9 -= (v9 | 467)
            switch -9223372036854775808 {
            case -365:
                p1 += p1
            case 808:
                switch "hello" {
                case "":
                    switch ("" + "hello") {
                    case "\n\t":
                        var v30 uint16 = 377
                        _ = v30
                    case "hello":
                        v31 := GenS4{F0: 508, F1: true, F2: 2147483647}
                        _ = v31
                    case "":
                        switch ("\n\t" + "foo bar") {
                        case "":
                            p0 -= p0
                        case "foo bar":
                            var v32 uint8 = 42
                            _ = v32
                        case "hello":
                            var v33 uint = (^(v9))
                            _ = v33
                        default:
                            var v34 int8 = (-(p1))
                            _ = v34
                        }
                    default:
                        switch "" {
                        case "\n\t":
                            var v35 uint16 = 510
                            _ = v35
                        case "a":
                            v36 := []int{-753, -353, -704}
                            _ = v36
                        case "foo bar":
                            var v37 string = "a"
                            _ = v37
                        }
                    }
                }
            default:
                p1 = (p1 ^ (p1 * p1))
            }
            for fi38 := 0; fi38 < 18; fi38++ {
                v39 := []int{int(p1), (^(fi38))}
                _ = fi38
                _ = v39
            }
        }
    }
    var v40 uint8 = 255
    _ = p0
    _ = p1
    _ = v40
    _ = v9
    return true
}
func genFn1(p0 uint8) bool {
    for fi41 := 0; fi41 < 7; fi41++ {
        v42 := GenS4{F0: 872, F1: false, F2: 767}
        _ = fi41
        _ = v42
    }
    var v43 bool = (65535 != 710)
    _ = p0
    _ = v43
    return v43
}

func main() {
    v1 := make(map[string]string)
    var v2 uint = 0
    var v3 bool = (true && false)
    delete(v1, ("hello" + ""))
    v2 -= (^(v2))
    switch "foo bar" {
    case "hello":
        if ((42 <= 139) && (1 > 128)) {
            v1["foo bar"] = ("" + "\n\t")
        } else {
            for fi44 := 0; fi44 < 19; fi44++ {
                if (-83 <= -32768) {
                    switch "a" {
                    case "\n\t":
                        for fi45 := 0; fi45 < 8; fi45++ {
                            delete(v1, "foo bar")
                            _ = fi45
                        }
                    case "foo bar":
                        v2 -= (v2 & 65535)
                    default:
                        var v46 uint32 = 0
                        _ = v46
                    }
                } else {
                    v47 := GenS5{F0: 2, F1: 150, F2: 23}
                    v48 := GenS4{F0: 2147483648, F1: v3, F2: 2147483647}
                    var v49 int8 = -128
                    _ = v47
                    _ = v48
                    _ = v49
                }
                if (!false) {
                    var v50 uint8 = 42
                    _ = v50
                }
                _ = len(v1)
                _ = fi44
            }
            v2 -= v2
            if (float32(1e10) > (float32(1e10) + 1e10)) {
                v3 = v3
                fmt.Println(v3, v3)
                for fi51 := 0; fi51 < 19; fi51++ {
                    v3 = v3
                    for fi52 := 0; fi52 < 20; fi52++ {
                        var v53 bool = v3
                        var v54 uint32 = 2
                        v55 := GenS6{F0: v2}
                        _ = fi52
                        _ = v53
                        _ = v54
                        _ = v55
                    }
                    _ = fi51
                }
            } else {
                switch "a" {
                case "":
                    if (!false) {
                        v56 := []int{}
                        fmt.Println(v2)
                        if v3 {
                            var v57 string = ("\n\t" + "foo bar")
                            var v58 string = "hello"
                            _ = v57
                            _ = v58
                        }
                        _ = v56
                    }
                default:
                    var v59 string = "hello"
                    _ = v59
                }
                fmt.Println(v2, v2)
                for fi60 := 0; fi60 < 15; fi60++ {
                    fmt.Println(fi60, v3, fi60)
                    var v61 int8 = -15
                    _ = fi60
                    _ = v61
                }
            }
        }
    }
    switch "foo bar" {
    case "a":
        fmt.Println(v2)
    default:
        for fi62 := 0; fi62 < 20; fi62++ {
            v2 = v2
            switch (fi62 * fi62) {
            case 32767:
                var v63 int16 = -694
                _ = v63
            case 127:
                var v64 string = ("a" + (("hello" + "foo bar") + "a"))
                _ = v64
            }
            v2 = 1
            _ = fi62
        }
    }
    var v65 uint32 = 2147483647
    v65 += v65
    v3 = ((v2 > (v2 * v2)) || (v3 || v3))
    switch -108 {
    case 9223372036854775807:
        fmt.Println(v2)
    case 127:
        v66 := GenS7{F0: 255}
        _ = v66
    default:
        v2 += v2
    }
    if true {
        if (934 >= 1) {
            v1["\n\t"] = "\n\t"
            for fi67 := 0; fi67 < 1; fi67++ {
                var v68 uint16 = 995
                _ = fi67
                _ = v68
            }
        }
    } else {
        v3 = (int32(v65) == -346)
        v69 := GenS8{F0: 42, F1: v3, F2: ("" + ("foo bar" + ("foo bar" + "")))}
        _ = v69
    }
    switch -354 {
    case 1:
        v65 = (v65 * v65)
    }
    switch -9223372036854775808 {
    case 2147483647:
        v70 := GenS7{F0: 61}
        _ = v70
    case -305:
        v65 += (v65 << ((v2) & 63))
    default:
        _ = len(v1)
    }
    for fi71 := 0; fi71 < 7; fi71++ {
        v72 := GenS9{F0: -2147483648}
        _ = fi71
        _ = v72
    }
    if v3 {
        for fi73 := 0; fi73 < 19; fi73++ {
            for fi74 := 0; fi74 < 13; fi74++ {
                v65 += (v65 & uint32(v2))
                v75 := v1[("hello" + "hello")]
                fmt.Println(v75)
                if v3 {
                    if v3 {
                        switch (-(fi74)) {
                        case 9223372036854775807:
                            var v76 float32 = 1.5
                            _ = v76
                        case 465:
                            if true {
                                v65 = (v65 << (((v2 % (((^(v2))) | 1))) & 63))
                                v75 = (v75 + (v75 + (v75 + "")))
                                v77 := GenS10{F0: 509}
                                _ = v77
                            }
                        case 4294967295:
                            v78 := GenS11{F0: 0, F1: 982}
                            _ = v78
                        default:
                            switch (((v75 + v75) + v75) + "") {
                            case "foo bar":
                                var v79 float64 = (1.5 - 3.14159)
                                _ = v79
                            case "\n\t":
                                v80 := GenS12{F0: -56}
                                _ = v80
                            case "":
                                v81 := []int{32767}
                                _ = v81
                            }
                        }
                        fmt.Println(v3, v2)
                        v65 = 965
                        v65 = 65
                    } else {
                        switch fi74 {
                        case 1:
                            switch (fi73 + fi74) {
                            case 0:
                                var v82 float32 = 3.14159
                                _ = v82
                            case 32767:
                                v83 := v1[(v75 + "hello")]
                                _ = v83
                            case 714:
                                fmt.Println(fi74)
                            default:
                                switch (fi73 % (((fi74 - 0)) | 1)) {
                                case 739:
                                    v84 := v1[v75]
                                    _ = v84
                                case -794:
                                    switch "\n\t" {
                                    case "foo bar":
                                        if v3 {
                                            var v85 float64 = 0.1
                                            var v86 uint16 = 2
                                            v1[v75] = ((v75 + ("a" + "hello")) + v75)
                                            _ = v85
                                            _ = v86
                                        }
                                    default:
                                        fmt.Println(v75)
                                    }
                                default:
                                    v87 := GenS9{F0: -1}
                                    _ = v87
                                }
                            }
                        default:
                            if true {
                                fmt.Println(v65, v2, v75)
                                switch v75 {
                                case "foo bar":
                                    if true {
                                        v88 := GenS13{F0: 4294967295}
                                        _ = v88
                                    }
                                case "":
                                    v89 := GenS5{F0: v2, F1: v65, F2: 44}
                                    _ = v89
                                case "hello":
                                    v90 := GenS5{F0: (v2 | v2), F1: (^(v65)), F2: int8(v65)}
                                    _ = v90
                                }
                                var v91 float32 = 0.1
                                _ = v91
                            } else {
                                v75 = v75
                            }
                        }
                    }
                    v92 := []int{(fi74 * (fi74 & fi73))}
                    var v93 uint64 = 0
                    _ = v92
                    _ = v93
                }
                _ = fi74
                _ = v75
            }
            _ = fi73
        }
    } else {
        fmt.Println(v65)
        switch "\n\t" {
        case "\n\t":
            switch 1 {
            case -2147483648:
                for fi94 := 0; fi94 < 6; fi94++ {
                    v3 = (-529 > -32768)
                    v2 = 137
                    for fi95 := 0; fi95 < 6; fi95++ {
                        switch "" {
                        case "foo bar":
                            v96 := GenS14{F0: (v65 + (^(v65)))}
                            _ = v96
                        case "hello":
                            v97 := GenS15{F0: 127, F1: 0, F2: (-(fi95))}
                            if true {
                                v98 := GenS16{F0: (((680 > 0) || v3) || (0.1 >= (float64(-2.25) * 1.5))), F1: -117, F2: 485}
                                _ = v98
                            } else {
                                v99 := []int{(fi95 << ((v2) & 63)), 0, fi95}
                                v100 := GenS17{F0: 517, F1: v65, F2: (((v65 | (v65 / ((65535) | 1))) != v65) || (-1 != 1))}
                                _ = v100
                                _ = v99
                            }
                            _ = v97
                        }
                        if (!(128 != 131)) {
                            v101 := GenS15{F0: 0, F1: -9223372036854775808, F2: (fi94 * (fi95 & fi94))}
                            var v102 uint = v2
                            var v103 float64 = -2.25
                            _ = v101
                            _ = v102
                            _ = v103
                        }
                        v3 = (v3 || ("foo bar" != (("hello" + "foo bar") + "\n\t")))
                        _ = fi95
                    }
                    _ = fi94
                }
            }
        case "":
            if (!v3) {
                v104 := []int{0, 127, 0}
                var v105 float32 = 1.5
                _ = v104
                _ = v105
            }
        default:
            v65 -= (^(v65))
        }
    }
    v3 = (1 != 1)
    v2 = (^(v2))
    v3 = false
    for fi106 := 0; fi106 < 2; fi106++ {
        for fi107 := 0; fi107 < 3; fi107++ {
            if false {
                if (0 != 32767) {
                    v3 = v3
                    fmt.Println(fi107, v2)
                    fmt.Println(v3, v3)
                }
            } else {
                v108 := GenS16{F0: v3, F1: -113, F2: 0}
                fmt.Println(fi107, fi107)
                fmt.Println(v2, fi106)
                _ = v108
            }
            fmt.Println(fi106, fi107, fi106)
            _ = fi107
        }
        v65 = (v65 << ((v2) & 63))
        if true {
            var v109 uint = (v2 % ((v2) | 1))
            _ = v109
        } else {
            v65 = 2147483647
        }
        _ = fi106
    }
    var v110 int = 127
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v110:", v110)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v65:", v65)
    _ = v1
    _ = v110
    _ = v2
    _ = v3
    _ = v65
}
