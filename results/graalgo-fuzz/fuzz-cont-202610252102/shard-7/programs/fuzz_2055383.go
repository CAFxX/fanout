package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int8; F1 int }
type GenS3 struct { F0 int16; F1 string; F2 uint8 }
type GenS4 struct { F0 int16; F1 int64; F2 uint }
type GenS5 struct { F0 int32; F1 bool }
type GenS6 struct { F0 int32; F1 uint16 }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 int64; F1 uint64 }
type GenS9 struct { F0 bool }
type GenS10 struct { F0 uint64; F1 bool; F2 int8 }
type GenS11 struct { F0 int16; F1 bool; F2 uint32 }
type GenS12 struct { F0 int32; F1 uint16; F2 int16 }
type GenS13 struct { F0 uint; F1 bool; F2 uint }
type GenS14 struct { F0 int8; F1 int32; F2 int }
type GenS15 struct { F0 string; F1 string }
type GenS16 struct { F0 int }
func genFn0(p0 int, p1 int16, p2 uint32) uint32 {
    if false {
        switch (("foo bar" + ("\n\t" + "\n\t")) + ("foo bar" + "\n\t")) {
        case "a":
            v5 := []int{p0, p0, p0}
            _ = v5
        case "foo bar":
            p0 = p0
        }
    }
    p0 -= (p0 << ((0) & 63))
    p1 = (p1 - p1)
    v6 := GenS1{F0: 2147483647}
    var v7 uint8 = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return p2
}

func main() {
    var v1 uint = 2
    var v2 bool = (false && false)
    v3 := make(map[int]bool)
    v4 := make(map[string]int)
    var v8 uint32 = 2147483647
    fmt.Println(v2)
    switch 2147483647 {
    case -9223372036854775808:
        switch "hello" {
        case "":
            for fi9 := 0; fi9 < 2; fi9++ {
                for fi10 := 0; fi10 < 14; fi10++ {
                    fmt.Println(v8)
                    _ = fi10
                }
                switch "foo bar" {
                case "hello":
                    v11 := []int{fi9, -349, 9223372036854775807, (fi9 << (((v1 % (((^(v1))) | 1))) & 63))}
                    _ = v11
                case "":
                    if true {
                        v12 := []int{}
                        _ = v12
                    } else {
                        switch (^(fi9)) {
                        case 658:
                            var v13 string = "foo bar"
                            _ = v13
                        case -226:
                            v14 := GenS2{F0: -101, F1: int(v8)}
                            _ = v14
                        default:
                            v15 := []int{(-(fi9))}
                            _ = v15
                        }
                        v2 = (((!v2) || (4294967295 <= 1)) && true)
                        switch "a" {
                        case "\n\t":
                            v1 += v1
                        case "":
                            var v16 bool = ("foo bar" != "\n\t")
                            _ = v16
                        case "foo bar":
                            var v17 int32 = 2147483647
                            _ = v17
                        default:
                            switch fi9 {
                            case 243:
                                v8 = v8
                            case 32768:
                                v18 := GenS3{F0: 32767, F1: "\n\t", F2: 241}
                                _ = v18
                            default:
                                v1 += v1
                            }
                        }
                    }
                case "foo bar":
                    switch "hello" {
                    case "hello":
                        var v19 int32 = 341
                        _ = v19
                    case "":
                        if (1 != 1) {
                            v20 := GenS1{F0: (v8 + genFn0((fi9 << ((v1) & 63)), 355, v8))}
                            v21 := []int{801, (fi9 + (fi9 ^ (fi9 - (^(fi9))))), (fi9 ^ -2147483648)}
                            switch "\n\t" {
                            case "hello":
                                var v22 int64 = 473
                                _ = v22
                            }
                            _ = v20
                            _ = v21
                        }
                    default:
                        fmt.Println(v1, v2)
                    }
                default:
                    v23 := GenS4{F0: 42, F1: -9223372036854775808, F2: 55}
                    _ = v23
                }
                _ = fi9
            }
        case "foo bar":
            v8 = uint32(v1)
        default:
            fmt.Println(v2, v8, v8)
        }
    case -389:
        v24 := GenS4{F0: 33, F1: 52, F2: (^(v1))}
        _ = v24
    }
    switch 915 {
    case 0:
        fmt.Println(v1, v1)
    case 32768:
        if v2 {
            v1 -= 65535
            var v25 bool = (!v2)
            _ = v25
        }
    default:
        var v26 int64 = 42
        _ = v26
    }
    var v27 uint8 = 42
    for fi28 := 0; fi28 < 5; fi28++ {
        v29 := GenS3{F0: -32768, F1: "", F2: (^(v27))}
        _ = fi28
        _ = v29
    }
    if (float64(1e10) <= (-2.25 + (1e10 / 1.0))) {
        var v30 string = ("a" + ("\n\t" + "foo bar"))
        var v31 int8 = -1
        _ = v30
        _ = v31
    } else {
        v8 += 2147483647
        v2 = ((("\n\t" + "\n\t") + ("hello" + ("foo bar" + "foo bar"))) > "a")
        for fi32 := 0; fi32 < 13; fi32++ {
            fmt.Println(v2, fi32)
            fmt.Println(v8)
            _ = fi32
        }
    }
    var v33 int32 = 1
    v33 += 125
    switch ("hello" + "\n\t") {
    case "a":
        for fi34 := 0; fi34 < 7; fi34++ {
            v27 = (v27 * (^(v27)))
            fmt.Println(v8)
            _ = fi34
        }
    case "":
        if v2 {
            v2 = v2
            v35 := GenS4{F0: 1000, F1: -9223372036854775808, F2: v1}
            _ = v35
        }
    }
    if v2 {
        fmt.Println(v8, v8, v33)
        v1 -= 1
        v36 := []int{-617, 413}
        _ = v36
    }
    switch ("foo bar" + "hello") {
    case "":
        switch 127 {
        case 127:
            if v2 {
                fmt.Println(v1, v8)
            } else {
                var v37 int16 = -678
                _ = v37
            }
        }
    case "\n\t":
        var v38 int32 = (v33 % (((v33 ^ (^(v33)))) | 1))
        _ = v38
    default:
        for fi39 := 0; fi39 < 20; fi39++ {
            v40 := GenS2{F0: -92, F1: fi39}
            if (!true) {
                switch fi39 {
                case -750:
                    var v41 int = fi39
                    _ = v41
                case 0:
                    v42 := GenS4{F0: 1, F1: 9223372036854775807, F2: 804}
                    _ = v42
                case -1:
                    v2 = (1.5 <= 1e10)
                }
            } else {
                var v43 int16 = -119
                v44 := GenS3{F0: -266, F1: "foo bar", F2: (v27 - (v27 * 0))}
                _ = v43
                _ = v44
            }
            _ = fi39
            _ = v40
        }
    }
    v8 -= 2147483648
    if true {
        v45 := []int{-521}
        _ = v45
    } else {
        var v46 bool = (v2 && (!false))
        v2 = v46
        _ = v46
    }
    switch 42 {
    case 32767:
        switch "hello" {
        case "foo bar":
            switch 4294967295 {
            case -655:
                v2 = (!v2)
            case 32768:
                var v47 int32 = 42
                _ = v47
            default:
                fmt.Println(v2, v2)
            }
        case "":
            v48 := GenS4{F0: 0, F1: 42, F2: 649}
            _ = v48
        default:
            if v2 {
                if v2 {
                    v1 = 18446744073709551615
                    v8 = genFn0(128, 42, genFn0(-507, -456, v8))
                    switch 9223372036854775807 {
                    case 9223372036854775807:
                        v49 := []int{2147483647, -9223372036854775808, 445, -88}
                        _ = v49
                    }
                }
                var v50 uint8 = v27
                _ = v50
            }
        }
    case 42:
        v51 := []int{2147483647, 128}
        _ = v51
    }
    if v2 {
        v52 := GenS5{F0: 893, F1: (!(v1 > v1))}
        v53 := GenS6{F0: (v33 / ((0) | 1)), F1: 42}
        v54 := GenS7{F0: v1}
        _ = v52
        _ = v53
        _ = v54
    } else {
        _ = len(v4)
        for fi55 := 0; fi55 < 8; fi55++ {
            switch 9223372036854775807 {
            case 1:
                var v56 int = (fi55 & 1)
                _ = v56
            }
            if v2 {
                for fi57 := 0; fi57 < 2; fi57++ {
                    if false {
                        var v58 int64 = -380
                        _ = v58
                    } else {
                        if v2 {
                            if true {
                                if v2 {
                                    v59 := GenS8{F0: -9223372036854775808, F1: 2}
                                    _ = v59
                                } else {
                                    var v60 uint8 = uint8(v1)
                                    _ = v60
                                }
                                v2 = (-208 == -32768)
                                fmt.Println(fi55, v33, v2)
                            } else {
                                if true {
                                    v1 = 578
                                }
                                v61 := GenS6{F0: v33, F1: 293}
                                _ = v61
                            }
                            v33 = (v33 - (v33 * v33))
                            var v62 bool = (!false)
                            _ = v62
                        } else {
                            switch "a" {
                            case "hello":
                                switch (fi55 ^ (fi57 + fi55)) {
                                case 2147483647:
                                    delete(v3, 0)
                                case 0:
                                    if v2 {
                                        var v63 int8 = -1
                                        v1 = (v1 | v1)
                                        v64 := []int{fi55, fi55, fi55}
                                        _ = v63
                                        _ = v64
                                    }
                                }
                            default:
                                var v65 uint8 = (v27 - v27)
                                _ = v65
                            }
                            if false {
                                v66 := v4["foo bar"]
                                fmt.Println(v66)
                                _ = v66
                            }
                        }
                        if false {
                            if (!(-2147483648 > v33)) {
                                switch "\n\t" {
                                case "a":
                                    v8 = (v8 << ((v1) & 63))
                                case "":
                                    if (1 >= (v8 / (((v8 << (((v1 & (v1 | 355))) & 63))) | 1))) {
                                        var v67 uint64 = 18446744073709551615
                                        var v68 float32 = 3.14159
                                        v27 = 53
                                        _ = v67
                                        _ = v68
                                    } else {
                                        var v69 float32 = 0.1
                                        _ = v69
                                    }
                                }
                                v1 = 42
                                switch "\n\t" {
                                case "":
                                    v70 := GenS9{F0: v2}
                                    _ = v70
                                case "hello":
                                    var v71 int = fi57
                                    _ = v71
                                case "foo bar":
                                    var v72 int64 = 475
                                    _ = v72
                                }
                            }
                            switch (fi55 ^ (fi57 | fi57)) {
                            case -262:
                                v73 := GenS4{F0: int16(v33), F1: -1, F2: v1}
                                v74 := []int{fi57}
                                _ = v73
                                _ = v74
                            case 127:
                                v8 -= v8
                                var v75 int16 = 32767
                                _ = v75
                            case 603:
                                v76 := GenS10{F0: 4294967295, F1: v2, F2: -6}
                                _ = v76
                            }
                        } else {
                            switch "foo bar" {
                            case "a":
                                _ = len(v4)
                            default:
                                v8 = 369
                            }
                            var v77 int16 = 0
                            _ = v77
                        }
                        v2 = false
                        fmt.Println(v8, v2, v1)
                    }
                    if (!false) {
                        var v78 int8 = -92
                        fmt.Println(fi57, v78)
                        fmt.Println(v2, fi57)
                        var v79 uint8 = v27
                        _ = v78
                        _ = v79
                    } else {
                        v2 = v2
                        v8 -= 2147483648
                    }
                    switch "\n\t" {
                    case "":
                        v33 = (^(v33))
                    default:
                        if (!(0 < 42)) {
                            if v2 {
                                v27 = v27
                                v33 += v33
                                switch 9223372036854775807 {
                                case 1:
                                    if v2 {
                                        var v80 int = fi57
                                        if (v80 > (fi57 / ((v80) | 1))) {
                                            v81 := GenS11{F0: 42, F1: ((v2 || (4294967295 > 0)) || true), F2: v8}
                                            fmt.Println(v33, fi57, v27)
                                            _ = v81
                                        }
                                        v82 := GenS4{F0: 1, F1: 532, F2: v1}
                                        _ = v80
                                        _ = v82
                                    }
                                case 2147483647:
                                    switch (fi57 % ((-2147483648) | 1)) {
                                    case 4294967295:
                                        v33 -= (-(v33))
                                    case -659:
                                        fmt.Println(v8, fi55, v2)
                                    }
                                    v83 := GenS12{F0: 1, F1: 2, F2: 790}
                                    if v2 {
                                        v84 := []int{}
                                        switch "foo bar" {
                                        case "foo bar":
                                            var v85 int32 = (v33 / (((v33 + (v33 - (v33 * (v33 + v33))))) | 1))
                                            _ = v85
                                        case "\n\t":
                                            v86 := GenS13{F0: (v1 << ((v1) & 63)), F1: (!v2), F2: 148}
                                            _ = v86
                                        case "a":
                                            var v87 bool = true
                                            _ = v87
                                        }
                                        _ = v84
                                    } else {
                                        v88 := GenS7{F0: v1}
                                        var v89 int64 = 690
                                        _ = v88
                                        _ = v89
                                    }
                                    v90 := v4[("\n\t" + ("a" + ("\n\t" + "\n\t")))]
                                    _ = v83
                                    _ = v90
                                case -2147483648:
                                    switch "" {
                                    case "foo bar":
                                        v2 = (v2 && (-2.25 < float32(3.14159)))
                                    case "":
                                        if (81 > -53) {
                                            v91 := GenS3{F0: 32767, F1: "hello", F2: (v27 & 17)}
                                            var v92 string = ""
                                            var v93 bool = ((!(!false)) && ((994 <= 255) && true))
                                            _ = v91
                                            _ = v92
                                            _ = v93
                                        }
                                        if (-1 > 0) {
                                            var v94 int32 = (v33 + v33)
                                            v94 -= v94
                                            if v2 {
                                                v95 := GenS6{F0: v94, F1: 255}
                                                _ = v95
                                            }
                                            v33 = v94
                                            _ = v94
                                        } else {
                                            var v96 uint16 = 255
                                            _ = v96
                                        }
                                        if (2 != 9) {
                                            fmt.Println(v8, v8, v1)
                                            var v97 uint = (v1 / ((v1) | 1))
                                            _ = v97
                                        }
                                    case "a":
                                        v98 := GenS14{F0: 41, F1: 0, F2: -429}
                                        _ = v98
                                    default:
                                        if v2 {
                                            v99 := []int{}
                                            v100 := v3[(fi57 / ((fi57) | 1))]
                                            v1 -= (^(v1))
                                            _ = v100
                                            _ = v99
                                        } else {
                                            v101 := GenS15{F0: "foo bar", F1: "\n\t"}
                                            _ = v101
                                        }
                                    }
                                }
                            }
                            delete(v4, ("" + ("" + "a")))
                            v2 = false
                        }
                    }
                    _ = fi57
                }
            } else {
                fmt.Println(v33, v1)
            }
            _ = fi55
        }
    }
    if (127 <= 1) {
        v33 -= v33
        v33 -= (-(v33))
    } else {
        switch "foo bar" {
        case "hello":
            for fi102 := 0; fi102 < 14; fi102++ {
                switch fi102 {
                case 9223372036854775807:
                    for fi103 := 0; fi103 < 13; fi103++ {
                        var v104 uint32 = 1
                        v1 = (v1 >> ((18446744073709551615) & 63))
                        switch "foo bar" {
                        case "\n\t":
                            v105 := v3[fi103]
                            _ = v105
                        case "":
                            v106 := []int{}
                            _ = v106
                        default:
                            var v107 uint16 = 194
                            _ = v107
                        }
                        _ = fi103
                        _ = v104
                    }
                }
                v1 = v1
                _ = fi102
            }
        case "foo bar":
            switch -9223372036854775808 {
            case 32767:
                var v108 string = "hello"
                _ = v108
            case -239:
                for fi109 := 0; fi109 < 11; fi109++ {
                    var v110 int64 = 418
                    fmt.Println(v8)
                    v111 := GenS16{F0: fi109}
                    _ = fi109
                    _ = v110
                    _ = v111
                }
            }
        case "\n\t":
            v112 := GenS3{F0: 0, F1: "hello", F2: v27}
            _ = v112
        default:
            switch -2147483648 {
            case 128:
                var v113 float64 = (((3.14159 * (0.0 - 1.5)) * 3.14159) + -2.25)
                _ = v113
            case 68:
                v114 := []int{0, 2147483647, 2147483647, 129}
                _ = v114
            case -679:
                var v115 int = 9223372036854775807
                _ = v115
            default:
                var v116 int16 = 346
                _ = v116
            }
        }
        v8 = uint32(v27)
    }
    for fi117 := 0; fi117 < 11; fi117++ {
        v2 = ((v33 | v33) >= v33)
        v118 := GenS3{F0: 1, F1: "foo bar", F2: (v27 | v27)}
        _ = fi117
        _ = v118
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v33
    _ = v4
    _ = v8
}
