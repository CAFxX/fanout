package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint32; F1 int16; F2 uint64 }
type GenS3 struct { F0 string; F1 uint64; F2 uint }
type GenS4 struct { F0 int32; F1 string }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 string; F1 int16 }
type GenS7 struct { F0 string; F1 int; F2 int16 }
type GenS8 struct { F0 int32 }
type GenS9 struct { F0 int64; F1 uint64; F2 int64 }
type GenS10 struct { F0 string; F1 uint }
type GenS11 struct { F0 uint64; F1 uint64; F2 string }
type GenS12 struct { F0 int32; F1 uint; F2 uint8 }
type GenS13 struct { F0 bool }
type GenS14 struct { F0 uint64 }
type GenS15 struct { F0 string; F1 uint64 }
type GenS16 struct { F0 int16 }
type GenS17 struct { F0 uint16 }
type GenS18 struct { F0 int8 }
type GenS19 struct { F0 string }
type GenS20 struct { F0 int8; F1 int64; F2 uint8 }
type GenS21 struct { F0 string }
type GenS22 struct { F0 uint64; F1 int }
type GenS23 struct { F0 int16; F1 uint16 }
type GenS24 struct { F0 int64; F1 uint64; F2 uint8 }
type GenS25 struct { F0 uint16; F1 int64 }
func genFn0(p0 int32, p1 int64) int8 {
    v7 := []int{472}
    p0 = -1
    _ = p0
    _ = p1
    _ = v7
    return 1
}
func genFn1(p0 uint16) uint8 {
    var v8 uint16 = p0
    v9 := GenS1{F0: 72}
    var v10 float32 = 0.0
    var v11 uint64 = 4294967295
    v12 := []int{}
    _ = p0
    _ = v10
    _ = v11
    _ = v12
    _ = v8
    _ = v9
    return 127
}

func main() {
    var v1 string = ("\n\t" + (((("foo bar" + (("foo bar" + ("" + "foo bar")) + ("foo bar" + "a"))) + ("\n\t" + ("foo bar" + "a"))) + "\n\t") + "foo bar"))
    var v2 uint = 18446744073709551615
    var v3 string = v1
    v4 := []uint32{832, 517}
    v5 := make(map[string]uint32)
    var v6 int8 = 0
    var v13 uint64 = 996
    v3 = "hello"
    fmt.Println(v2, v2, v13)
    v14 := GenS1{F0: (^(v2))}
    var v15 int64 = -9223372036854775808
    var v16 int32 = 42
    v15 = v15
    for fi17 := 0; fi17 < 15; fi17++ {
        v3 = "a"
        if false {
            var v18 int64 = -87
            var v19 uint8 = 0
            switch (fi17 / ((fi17) | 1)) {
            case -446:
                fmt.Println(v6, v15)
            default:
                v20 := GenS2{F0: 0, F1: -423, F2: (^(v13))}
                _ = v20
            }
            _ = v18
            _ = v19
        }
        var v21 uint = (v2 * 938)
        _ = fi17
        _ = v21
    }
    if false {
        if len(v4) > 0 {
            v4[0] = 42
        }
        for fi22 := 0; fi22 < 1; fi22++ {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            for fi23 := 0; fi23 < 9; fi23++ {
                fmt.Println(v13)
                v24 := GenS1{F0: (v2 >> ((v2) & 63))}
                _ = fi23
                _ = v24
            }
            var v25 uint = v2
            _ = fi22
            _ = v25
        }
        if (!false) {
            for fi26 := 0; fi26 < 3; fi26++ {
                switch v3 {
                case "":
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                default:
                    switch v1 {
                    case "a":
                        if (!(v1 == v1)) {
                            for fi27 := 0; fi27 < 15; fi27++ {
                                var v28 uint16 = 42
                                fmt.Println(v15)
                                switch "\n\t" {
                                case "a":
                                    switch fi27 {
                                    case -865:
                                        _ = len(v4)
                                    default:
                                        v13 = 18446744073709551615
                                    }
                                default:
                                    _ = len(v5)
                                }
                                var v29 int32 = (v16 & int32(v13))
                                _ = fi27
                                _ = v28
                                _ = v29
                            }
                        } else {
                            fmt.Println(v13, fi26, fi26)
                        }
                    case "":
                        fmt.Println(v6, v6)
                    }
                }
                _ = fi26
            }
        } else {
            for fi30 := 0; fi30 < 4; fi30++ {
                switch fi30 {
                case 32767:
                    for fi31 := 0; fi31 < 17; fi31++ {
                        var v32 uint16 = 65535
                        if false {
                            fmt.Println(v13, v1)
                            switch (-(fi31)) {
                            case 727:
                                var v33 uint16 = v32
                                _ = v33
                            case 32768:
                                v34 := GenS3{F0: ("" + ("foo bar" + (v3 + ("\n\t" + v3)))), F1: v13, F2: (v2 % (((v2 - 58)) | 1))}
                                _ = v34
                            default:
                                fmt.Println(v2, v15)
                            }
                        } else {
                            _ = len(v5)
                            delete(v5, "a")
                            v5[(v3 + v1)] = 186
                        }
                        switch fi30 {
                        case -1:
                            v35 := GenS2{F0: 42, F1: 605, F2: v13}
                            _ = v35
                        case 0:
                            v5[("foo bar" + v3)] = 0
                        default:
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                        }
                        if (1.5 == ((1e10 + float64((3.14159 * ((1e10 / (3.14159 / 1e10)) + float32(1.5))))) + (1.5 * 0.0))) {
                            if true {
                                var v36 uint64 = (v13 * v13)
                                _ = v36
                            } else {
                                if false {
                                    var v37 float64 = 3.14159
                                    _ = v37
                                } else {
                                    switch v3 {
                                    case "\n\t":
                                        v38 := GenS3{F0: (v3 + "foo bar"), F1: 0, F2: 181}
                                        _ = v38
                                    case "a":
                                        switch fi30 {
                                        case 2147483647:
                                            v15 += 9223372036854775807
                                        case -623:
                                            var v39 uint32 = 2147483648
                                            v3 = v1
                                            fmt.Println(v3, v2, v1)
                                            _ = v39
                                        case 430:
                                            switch fi31 {
                                            case 128:
                                                switch int(v2) {
                                                case -2147483648:
                                                    v13 -= (v13 >> (((v2 & (v2 + (v2 ^ (v2 % (((^(v2))) | 1)))))) & 63))
                                                }
                                            case 858:
                                                if false {
                                                    switch "" {
                                                    case "\n\t":
                                                        var v40 uint = (v2 & (^(v2)))
                                                        _ = v40
                                                    case "a":
                                                        v32 -= uint16(v13)
                                                    }
                                                    switch (fi31 + 42) {
                                                    case 2147483647:
                                                        if (true || false) {
                                                            var v41 uint64 = v13
                                                            fmt.Println(v15)
                                                            switch v1 {
                                                            case "foo bar":
                                                                v32 = (v32 / ((v32) | 1))
                                                            case "a":
                                                                var v42 int = (fi30 % (((^(fi31))) | 1))
                                                                _ = v42
                                                            default:
                                                                var v43 uint16 = uint16(v2)
                                                                _ = v43
                                                            }
                                                            _ = v41
                                                        } else {
                                                            v1 = "hello"
                                                            if (!false) {
                                                                switch v1 {
                                                                case "":
                                                                    fmt.Println(v16, fi31)
                                                                case "hello":
                                                                    if true {
                                                                        if true {
                                                                            fmt.Println(v2, fi30, v1)
                                                                            var v44 string = v3
                                                                            var v45 string = ("hello" + v1)
                                                                            _ = v44
                                                                            _ = v45
                                                                        } else {
                                                                            switch ((v1 + "") + v3) {
                                                                            case "\n\t":
                                                                                v2 = 18446744073709551615
                                                                            default:
                                                                                var v46 float64 = float64(0.1)
                                                                                _ = v46
                                                                            }
                                                                            fmt.Println(fi30)
                                                                            var v47 int64 = v15
                                                                            _ = v47
                                                                        }
                                                                        if (v13 > 42) {
                                                                            v16 += v16
                                                                            var v48 float32 = (-2.25 * 0.0)
                                                                            v3 = v3
                                                                            _ = v48
                                                                        } else {
                                                                            var v49 int16 = 466
                                                                            fmt.Println(v16, fi30)
                                                                            if true {
                                                                                fmt.Println(v13, v3)
                                                                                v3 = ("a" + v3)
                                                                            }
                                                                            _ = v49
                                                                        }
                                                                        if len(v4) > 0 {
                                                                            v4[0] = 4294967295
                                                                        }
                                                                        switch ((v1 + "") + ("foo bar" + v1)) {
                                                                        case "foo bar":
                                                                            switch "foo bar" {
                                                                            case "foo bar":
                                                                                v1 = ""
                                                                            case "":
                                                                                fmt.Println(v2)
                                                                                switch v3 {
                                                                                case "\n\t":
                                                                                    switch (fi31 | 42) {
                                                                                    case 850:
                                                                                        v1 = v1
                                                                                    case 32768:
                                                                                        if (!false) {
                                                                                            v1 = (v3 + ((v3 + ("a" + ((v1 + ("" + (("hello" + ("\n\t" + v1)) + ""))) + "\n\t"))) + ((v1 + "hello") + "a")))
                                                                                            v32 = v32
                                                                                            if (!true) {
                                                                                                switch ("\n\t" + ((v3 + ((v1 + v1) + v3)) + "a")) {
                                                                                                case "foo bar":
                                                                                                    var v50 uint = (v2 & 0)
                                                                                                    _ = v50
                                                                                                default:
                                                                                                    fmt.Println(v6, v1)
                                                                                                }
                                                                                            } else {
                                                                                                var v51 uint8 = uint8(v2)
                                                                                                switch fi30 {
                                                                                                case -1:
                                                                                                    switch v3 {
                                                                                                    case "a":
                                                                                                        _ = len(v5)
                                                                                                    default:
                                                                                                        var v52 uint = v2
                                                                                                        _ = v52
                                                                                                    }
                                                                                                case 1:
                                                                                                    if (false && (3.14159 != 3.14159)) {
                                                                                                        if (false || true) {
                                                                                                            if true {
                                                                                                                switch "\n\t" {
                                                                                                                case "hello":
                                                                                                                    fmt.Println(v3, v51, v1)
                                                                                                                case "a":
                                                                                                                    v5[v1] = 158
                                                                                                                case "":
                                                                                                                    v16 = v16
                                                                                                                default:
                                                                                                                    var v53 int8 = -20
                                                                                                                    _ = v53
                                                                                                                }
                                                                                                                var v54 string = v3
                                                                                                                _ = v54
                                                                                                            }
                                                                                                            v55 := GenS2{F0: 477, F1: 32767, F2: uint64(v51)}
                                                                                                            delete(v5, v3)
                                                                                                            _ = v55
                                                                                                        }
                                                                                                        var v56 uint8 = v51
                                                                                                        fmt.Println(v2, fi31, fi31)
                                                                                                        _ = v56
                                                                                                    } else {
                                                                                                        var v57 int32 = (-(v16))
                                                                                                        v57 = -2147483648
                                                                                                        _ = v57
                                                                                                    }
                                                                                                }
                                                                                                v16 += 747
                                                                                                _ = v51
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                case "a":
                                                                                    fmt.Println(v32, v2)
                                                                                case "hello":
                                                                                    var v58 int = (fi30 | fi30)
                                                                                    if true {
                                                                                        var v59 int32 = v16
                                                                                        switch (^(fi31)) {
                                                                                        case 886:
                                                                                            fmt.Println(v13, v16, v15)
                                                                                        case 838:
                                                                                            switch v3 {
                                                                                            case "":
                                                                                                v60 := GenS2{F0: 0, F1: 0, F2: uint64(v2)}
                                                                                                _ = v60
                                                                                            case "\n\t":
                                                                                                if false {
                                                                                                    var v61 int = (-(v58))
                                                                                                    var v62 int16 = 42
                                                                                                    v16 -= (v59 << (((v2 ^ 2)) & 63))
                                                                                                    _ = v61
                                                                                                    _ = v62
                                                                                                } else {
                                                                                                    v4 = append(v4, 2147483647)
                                                                                                    v63 := GenS1{F0: 4294967295}
                                                                                                    v59 = (v59 | v16)
                                                                                                    var v64 uint = (v2 % ((313) | 1))
                                                                                                    fmt.Println(fi30, v3, v32)
                                                                                                    _ = v63
                                                                                                    _ = v64
                                                                                                }
                                                                                            }
                                                                                        default:
                                                                                            v1 = v1
                                                                                        }
                                                                                        var v65 float32 = 1.5
                                                                                        _ = len(v4)
                                                                                        if false {
                                                                                            var v66 bool = ("" < ((("" + (v3 + v1)) + "") + v3))
                                                                                            _ = v66
                                                                                        }
                                                                                        _ = v59
                                                                                        _ = v65
                                                                                    }
                                                                                    _ = v58
                                                                                }
                                                                                if len(v4) >= 2 {
                                                                                    v4 = v4[1:len(v4)]
                                                                                }
                                                                            case "hello":
                                                                                v4 = append(v4, 2147483647)
                                                                            }
                                                                        case "a":
                                                                            _ = len(v5)
                                                                        default:
                                                                            v2 = (v2 - (v2 ^ uint(v32)))
                                                                        }
                                                                        if len(v4) > 0 {
                                                                            v4[0] = uint32(v32)
                                                                        }
                                                                        v2 = 790
                                                                    }
                                                                default:
                                                                    v67 := v5[v1]
                                                                    _ = v67
                                                                }
                                                                v1 = ""
                                                                var v68 int8 = v6
                                                                _ = v68
                                                            } else {
                                                                v69 := GenS3{F0: (("a" + v1) + "a"), F1: (v13 * 2), F2: uint(v32)}
                                                                switch v1 {
                                                                case "\n\t":
                                                                    var v70 float64 = (1.5 / 1.0)
                                                                    _ = v70
                                                                }
                                                                _ = v69
                                                            }
                                                            switch fi30 {
                                                            case -1:
                                                                v1 = ""
                                                            default:
                                                                if (358 > 4294967295) {
                                                                    switch int(v6) {
                                                                    case 42:
                                                                        v16 = (v16 | (v16 << (((v2 / ((v2) | 1))) & 63)))
                                                                    case -780:
                                                                        v71 := GenS4{F0: 735, F1: v1}
                                                                        _ = v71
                                                                    case 9223372036854775807:
                                                                        if (!(!true)) {
                                                                            v72 := GenS2{F0: 42, F1: 32767, F2: 656}
                                                                            var v73 int64 = -9223372036854775808
                                                                            fmt.Println(v2, v3, v3)
                                                                            _ = v72
                                                                            _ = v73
                                                                        } else {
                                                                            v13 -= 734
                                                                            if (!true) {
                                                                                if false {
                                                                                    if len(v4) >= 2 {
                                                                                        v4 = v4[1:len(v4)]
                                                                                    }
                                                                                    fmt.Println(v2)
                                                                                    v2 = uint(v13)
                                                                                } else {
                                                                                    if (65535 <= 2) {
                                                                                        v74 := GenS3{F0: v3, F1: uint64(v2), F2: 884}
                                                                                        v75 := GenS2{F0: 65535, F1: 42, F2: v13}
                                                                                        _ = v74
                                                                                        _ = v75
                                                                                    }
                                                                                    v16 = (v16 | 1)
                                                                                }
                                                                                if len(v4) >= 2 {
                                                                                    v4 = v4[1:len(v4)]
                                                                                }
                                                                                fmt.Println(v15)
                                                                            } else {
                                                                                if (genFn0((-(v16)), v15) == -82) {
                                                                                    if ((!true) && true) {
                                                                                        var v76 uint64 = 660
                                                                                        _ = v76
                                                                                    }
                                                                                    v77 := GenS2{F0: 318, F1: -1, F2: v13}
                                                                                    _ = v77
                                                                                } else {
                                                                                    var v78 uint32 = 1
                                                                                    v79 := GenS5{F0: v78}
                                                                                    v16 = int32(v15)
                                                                                    v13 += 1
                                                                                    v1 = v1
                                                                                    var v80 float32 = 0.0
                                                                                    var v81 uint8 = 42
                                                                                    var v82 int8 = v6
                                                                                    _ = v78
                                                                                    _ = v79
                                                                                    _ = v80
                                                                                    _ = v81
                                                                                    _ = v82
                                                                                }
                                                                            }
                                                                            v83 := GenS6{F0: "", F1: 32767}
                                                                            _ = v83
                                                                        }
                                                                    }
                                                                }
                                                            }
                                                        }
                                                    default:
                                                        var v84 string = "hello"
                                                        _ = v84
                                                    }
                                                    if (v15 >= v15) {
                                                        var v85 uint = v2
                                                        _ = v85
                                                    } else {
                                                        var v86 int16 = int16(v16)
                                                        if ((v32 | (v32 << ((18446744073709551615) & 63))) < v32) {
                                                            switch fi31 {
                                                            case 0:
                                                                v13 = (v13 << (((^(v2))) & 63))
                                                            case 9223372036854775807:
                                                                if false {
                                                                    if len(v4) >= 2 {
                                                                        v4 = v4[1:len(v4)]
                                                                    }
                                                                    fmt.Println(v3, v1, v16)
                                                                } else {
                                                                    var v87 int = (fi31 ^ (fi30 & fi30))
                                                                    v86 += (v86 << (((^(v2))) & 63))
                                                                    _ = v87
                                                                }
                                                            }
                                                            switch v3 {
                                                            case "hello":
                                                                if true {
                                                                    var v88 float32 = 3.14159
                                                                    var v89 string = ""
                                                                    _ = v88
                                                                    _ = v89
                                                                } else {
                                                                    var v90 int32 = (v16 / (((^(v16))) | 1))
                                                                    switch v3 {
                                                                    case "hello":
                                                                        v91 := GenS5{F0: 354}
                                                                        _ = v91
                                                                    case "\n\t":
                                                                        v16 += (v16 - (v90 << ((v2) & 63)))
                                                                    }
                                                                    var v92 uint8 = 127
                                                                    _ = v90
                                                                    _ = v92
                                                                }
                                                            case "foo bar":
                                                                if (0.0 <= 0.0) {
                                                                    v13 = (v13 + v13)
                                                                    var v93 uint = (v2 - (^(v2)))
                                                                    _ = v93
                                                                }
                                                            case "\n\t":
                                                                var v94 int64 = v15
                                                                _ = v94
                                                            }
                                                        } else {
                                                            v32 -= 65535
                                                        }
                                                        _ = v86
                                                    }
                                                } else {
                                                    var v95 string = (v1 + ("" + v3))
                                                    v16 -= (v16 / ((v16) | 1))
                                                    _ = v95
                                                }
                                            case 1:
                                                switch "\n\t" {
                                                case "\n\t":
                                                    if (false || (!true)) {
                                                        v96 := GenS7{F0: "foo bar", F1: int(v32), F2: -1}
                                                        var v97 int = (-(fi30))
                                                        _ = v96
                                                        _ = v97
                                                    } else {
                                                        v16 = (v16 | v16)
                                                        fmt.Println(fi30, v6)
                                                        v98 := GenS8{F0: (^(v16))}
                                                        var v99 uint16 = (v32 << ((632) & 63))
                                                        v100 := GenS9{F0: v15, F1: (v13 & 175), F2: -1}
                                                        _ = len(v5)
                                                        _ = v100
                                                        _ = v98
                                                        _ = v99
                                                    }
                                                case "a":
                                                    v101 := GenS10{F0: v1, F1: v2}
                                                    _ = v101
                                                case "foo bar":
                                                    switch "foo bar" {
                                                    case "a":
                                                        switch (fi31 + (fi31 - (fi30 % ((int(v6)) | 1)))) {
                                                        case 32768:
                                                            v102 := v5[v3]
                                                            _ = v102
                                                        default:
                                                            v13 -= v13
                                                        }
                                                    case "":
                                                        var v103 int64 = v15
                                                        _ = v103
                                                    case "hello":
                                                        switch (v1 + (("hello" + ("hello" + v3)) + "hello")) {
                                                        case "":
                                                            v32 -= (v32 % ((v32) | 1))
                                                        case "foo bar":
                                                            switch "\n\t" {
                                                            case "":
                                                                v104 := v5[v1]
                                                                _ = v104
                                                            case "hello":
                                                                v16 = 42
                                                            case "foo bar":
                                                                var v105 string = v3
                                                                _ = v105
                                                            default:
                                                                var v106 int16 = 1
                                                                _ = v106
                                                            }
                                                        case "a":
                                                            v32 = 472
                                                        }
                                                    default:
                                                        var v107 uint32 = 2147483647
                                                        _ = v107
                                                    }
                                                default:
                                                    var v108 uint64 = v13
                                                    _ = v108
                                                }
                                            }
                                        }
                                    }
                                }
                                if (false && true) {
                                    v109 := v5[(v1 + v1)]
                                    v6 -= 100
                                    if false {
                                        v2 += (^(v2))
                                        v110 := GenS4{F0: 2147483647, F1: "hello"}
                                        _ = v110
                                    }
                                    var v111 int8 = (^(v6))
                                    if (!true) {
                                        if false {
                                            v13 = v13
                                            var v112 uint64 = (v13 | uint64(v2))
                                            v5[((v1 + v1) + (v3 + v1))] = 1
                                            _ = v112
                                        } else {
                                            if true {
                                                var v113 int16 = 1
                                                _ = v113
                                            }
                                        }
                                    }
                                    fmt.Println(v111, v6, v111)
                                    _ = v109
                                    _ = v111
                                } else {
                                    if false {
                                        if true {
                                            var v114 int8 = 127
                                            _ = v114
                                        }
                                        if false {
                                            switch 127 {
                                            case 42:
                                                if false {
                                                    v115 := GenS8{F0: (v16 % (((v16 + (-(v16)))) | 1))}
                                                    if (false && (!(uint(v32) != (v2 | (v2 % ((v2) | 1)))))) {
                                                        _ = len(v5)
                                                        var v116 uint8 = 0
                                                        _ = v116
                                                    } else {
                                                        v117 := GenS11{F0: v13, F1: (v13 << (((^(v2))) & 63)), F2: "foo bar"}
                                                        v1 = ((("foo bar" + "") + (v3 + "hello")) + v1)
                                                        var v118 uint16 = v32
                                                        if len(v4) > 0 {
                                                            v4[0] = 65
                                                        }
                                                        _ = v117
                                                        _ = v118
                                                    }
                                                    _ = v115
                                                }
                                            case -296:
                                                var v119 int32 = (-(v16))
                                                _ = v119
                                            case 128:
                                                _ = len(v4)
                                            default:
                                                var v120 uint = (v2 * (v2 & (v2 ^ (^(v2)))))
                                                _ = v120
                                            }
                                            v4 = append(v4, 821)
                                        }
                                    } else {
                                        v32 = v32
                                        v121 := GenS10{F0: v3, F1: 255}
                                        var v122 uint16 = (v32 + v32)
                                        switch ((v3 + (v1 + ((((v1 + v3) + v1) + (((("" + v1) + "") + "") + ((v3 + v3) + "\n\t"))) + ""))) + v1) {
                                        case "foo bar":
                                            _ = len(v4)
                                        case "a":
                                            var v123 int32 = (v16 ^ (v16 / (((v16 << ((42) & 63))) | 1)))
                                            _ = v123
                                        }
                                        v3 = (v1 + (v3 + "hello"))
                                        v6 = (v6 >> (((v2 % ((921) | 1))) & 63))
                                        fmt.Println(v13, v2, v13)
                                        _ = v121
                                        _ = v122
                                    }
                                    v13 += v13
                                }
                                fmt.Println(v13, v2)
                            }
                            _ = len(v4)
                        } else {
                            var v124 int64 = (v15 + v15)
                            v125 := GenS3{F0: v1, F1: v13, F2: (v2 % ((v2) | 1))}
                            fmt.Println(v3, v2, v15)
                            v126 := GenS12{F0: (v16 | v16), F1: (v2 ^ v2), F2: 255}
                            switch -9223372036854775808 {
                            case 0:
                                if (false || (((v2 & (v2 / ((599) | 1))) >= v2) && true)) {
                                    var v127 int8 = (v6 << ((v2) & 63))
                                    var v128 bool = false
                                    v5[v3] = 806
                                    _ = v127
                                    _ = v128
                                } else {
                                    switch fi30 {
                                    case 7:
                                        v129 := GenS8{F0: (v16 << (((^(v2))) & 63))}
                                        _ = v129
                                    case 9223372036854775807:
                                        _ = len(v5)
                                    }
                                    var v130 int64 = v15
                                    if len(v4) > 0 {
                                        v4[0] = 275
                                    }
                                    _ = v130
                                }
                            default:
                                v32 += (v32 ^ (v32 * 255))
                                var v131 int16 = 1000
                                _ = v131
                            }
                            _ = v124
                            _ = v125
                            _ = v126
                        }
                        _ = fi31
                        _ = v32
                    }
                default:
                    if (!true) {
                        var v132 int64 = (v15 ^ v15)
                        _ = v132
                    } else {
                        v16 -= 2147483647
                        v133 := v5[""]
                        var v134 float32 = (0.0 * (-2.25 - float32((-2.25 - (1.5 * 0.1)))))
                        _ = v133
                        _ = v134
                    }
                }
                _ = fi30
            }
        }
    } else {
        for fi135 := 0; fi135 < 14; fi135++ {
            _ = len(v4)
            _ = fi135
        }
        if (((1.5 - (1.5 / 3.14159)) * float64(float32(0.1))) < 0.0) {
            v136 := v5[v1]
            for fi137 := 0; fi137 < 4; fi137++ {
                var v138 int16 = 42
                switch (fi137 & (fi137 - fi137)) {
                case -276:
                    v139 := GenS13{F0: false}
                    _ = v139
                case 2147483647:
                    switch (v1 + v1) {
                    case "a":
                        var v140 uint32 = 0
                        _ = v140
                    case "\n\t":
                        for fi141 := 0; fi141 < 8; fi141++ {
                            var v142 int32 = (v16 + int32(fi141))
                            var v143 uint32 = v136
                            switch fi137 {
                            case 1:
                                if true {
                                    var v144 float64 = (1.5 - 0.0)
                                    _ = v144
                                } else {
                                    var v145 bool = true
                                    fmt.Println(v142)
                                    switch v3 {
                                    case "a":
                                        var v146 int16 = 419
                                        _ = v146
                                    case "":
                                        if true {
                                            if ((v13 << (((v2 << ((v2) & 63))) & 63)) > v13) {
                                                v6 = (v6 - v6)
                                            }
                                            var v147 float32 = ((1.5 / 0.1) / 0.1)
                                            switch fi141 {
                                            case -561:
                                                v148 := GenS14{F0: v13}
                                                _ = v148
                                            }
                                            var v149 uint32 = (v136 << (((v2 / ((v2) | 1))) & 63))
                                            _ = v147
                                            _ = v149
                                        } else {
                                            switch v3 {
                                            case "hello":
                                                _ = len(v4)
                                            case "a":
                                                switch ("" + v1) {
                                                case "hello":
                                                    v150 := GenS15{F0: "\n\t", F1: 220}
                                                    _ = v150
                                                }
                                            case "foo bar":
                                                fmt.Println(v13, v6, v1)
                                            default:
                                                switch 128 {
                                                case 42:
                                                    if len(v4) > 0 {
                                                        v4[0] = v143
                                                    }
                                                case 212:
                                                    var v151 int = (fi137 + 32767)
                                                    _ = v151
                                                case 749:
                                                    switch v3 {
                                                    case "foo bar":
                                                        var v152 int = (-(fi141))
                                                        _ = v152
                                                    case "hello":
                                                        v142 -= 2147483647
                                                    default:
                                                        v153 := GenS16{F0: v138}
                                                        _ = v153
                                                    }
                                                    if (1 == 19) {
                                                        if (((v6 | (^(v6))) >= (v6 & genFn0((v142 - 2147483647), (v15 - (v15 | (v15 >> ((v2) & 63))))))) && v145) {
                                                            var v154 uint = (v2 | 18446744073709551615)
                                                            if false {
                                                                if len(v4) >= 2 {
                                                                    v4 = v4[1:len(v4)]
                                                                }
                                                                switch v1 {
                                                                case "hello":
                                                                    fmt.Println(v136, fi141, v142)
                                                                case "foo bar":
                                                                    if v145 {
                                                                        v142 -= int32(v138)
                                                                    } else {
                                                                        if v145 {
                                                                            delete(v5, "foo bar")
                                                                        }
                                                                        v142 += int32(v6)
                                                                        _ = len(v5)
                                                                        if ((-(v142)) == v142) {
                                                                            v1 = v3
                                                                        }
                                                                        if (((v16 * 42) >= (v16 ^ 0)) || (!v145)) {
                                                                            var v155 float64 = 1e10
                                                                            v145 = true
                                                                            v156 := GenS7{F0: ((v1 + v3) + v3), F1: (fi137 - fi141), F2: -933}
                                                                            _ = v155
                                                                            _ = v156
                                                                        } else {
                                                                            if (true && ((((v3 + v1) + (v1 + "foo bar")) + v1) == "hello")) {
                                                                                _ = len(v5)
                                                                                fmt.Println(v142, v15)
                                                                            }
                                                                            if len(v4) >= 2 {
                                                                                v4 = v4[1:len(v4)]
                                                                            }
                                                                            v157 := v5[v3]
                                                                            var v158 uint32 = v136
                                                                            _ = v157
                                                                            _ = v158
                                                                        }
                                                                    }
                                                                case "a":
                                                                    if v145 {
                                                                        var v159 uint64 = 18446744073709551615
                                                                        var v160 uint16 = 1
                                                                        var v161 uint = (v154 * (v154 ^ 65535))
                                                                        if v145 {
                                                                            v162 := GenS8{F0: (v16 + (v142 - v16))}
                                                                            switch 127 {
                                                                            case 2147483647:
                                                                                v161 = v2
                                                                            default:
                                                                                v163 := v5["a"]
                                                                                switch v3 {
                                                                                case "a":
                                                                                    var v164 int8 = genFn0(v16, v15)
                                                                                    _ = v164
                                                                                case "":
                                                                                    switch fi141 {
                                                                                    case 1:
                                                                                        var v165 int = (-(fi137))
                                                                                        fmt.Println(v1, v16)
                                                                                        _ = v165
                                                                                    case 331:
                                                                                        v160 = v160
                                                                                    default:
                                                                                        var v166 int8 = int8(v15)
                                                                                        _ = v166
                                                                                    }
                                                                                    if true {
                                                                                        delete(v5, ("a" + v3))
                                                                                        v167 := GenS10{F0: (((v1 + v3) + "") + v1), F1: (v154 >> ((v161) & 63))}
                                                                                        fmt.Println(fi137)
                                                                                        v168 := GenS17{F0: 42}
                                                                                        v13 += v13
                                                                                        if len(v4) > 0 {
                                                                                            v4[0] = v136
                                                                                        }
                                                                                        switch ("hello" + v1) {
                                                                                        case "":
                                                                                            switch v3 {
                                                                                            case "foo bar":
                                                                                                if true {
                                                                                                    fmt.Println(fi137, v154)
                                                                                                    if (v145 || (255 != 186)) {
                                                                                                        switch fi141 {
                                                                                                        case 930:
                                                                                                            v143 = (v136 ^ v163)
                                                                                                        default:
                                                                                                            var v169 string = v1
                                                                                                            _ = v169
                                                                                                        }
                                                                                                        var v170 float32 = 0.1
                                                                                                        var v171 float32 = v170
                                                                                                        _ = v170
                                                                                                        _ = v171
                                                                                                    } else {
                                                                                                        fmt.Println(v161, v1, v161)
                                                                                                        v138 -= 1000
                                                                                                    }
                                                                                                    var v172 string = (v1 + ("hello" + "foo bar"))
                                                                                                    _ = v172
                                                                                                } else {
                                                                                                    fmt.Println(v6, v3)
                                                                                                    switch fi137 {
                                                                                                    case 32767:
                                                                                                        _ = len(v4)
                                                                                                    case 2147483647:
                                                                                                        var v173 uint8 = 64
                                                                                                        _ = v173
                                                                                                    default:
                                                                                                        v142 += v142
                                                                                                    }
                                                                                                    v174 := GenS18{F0: (-(v6))}
                                                                                                    delete(v5, "foo bar")
                                                                                                    _ = len(v4)
                                                                                                    _ = v174
                                                                                                }
                                                                                            case "":
                                                                                                var v175 float32 = 0.1
                                                                                                _ = v175
                                                                                            case "\n\t":
                                                                                                v176 := GenS19{F0: "hello"}
                                                                                                _ = v176
                                                                                            }
                                                                                        case "hello":
                                                                                            v177 := GenS15{F0: v3, F1: (v13 % ((v13) | 1))}
                                                                                            _ = v177
                                                                                        case "foo bar":
                                                                                            switch 4294967295 {
                                                                                            case 0:
                                                                                                if v145 {
                                                                                                    switch (v3 + v3) {
                                                                                                    case "a":
                                                                                                        v16 += v142
                                                                                                    case "":
                                                                                                        fmt.Println(v6)
                                                                                                    }
                                                                                                    switch 32767 {
                                                                                                    case 32767:
                                                                                                        var v178 int64 = 42
                                                                                                        _ = v178
                                                                                                    case -2147483648:
                                                                                                        var v179 int32 = v16
                                                                                                        _ = v179
                                                                                                    }
                                                                                                    v4 = append(v4, 0)
                                                                                                } else {
                                                                                                    fmt.Println(v160)
                                                                                                    var v180 uint32 = (v143 ^ 42)
                                                                                                    _ = v180
                                                                                                }
                                                                                            default:
                                                                                                if ((0.1 > -2.25) || v145) {
                                                                                                    if len(v4) > 0 {
                                                                                                        v4[0] = 4294967295
                                                                                                    }
                                                                                                    var v181 uint8 = 128
                                                                                                    if (!((!false) || (127 <= v181))) {
                                                                                                        v13 = v13
                                                                                                        fmt.Println(fi137, v163, v143)
                                                                                                        switch v3 {
                                                                                                        case "hello":
                                                                                                            var v182 uint32 = v163
                                                                                                            _ = v182
                                                                                                        case "a":
                                                                                                            if ((v138 / (((v138 << ((512) & 63))) | 1)) != int16(v15)) {
                                                                                                                switch v3 {
                                                                                                                case "\n\t":
                                                                                                                    if ((v142 & v142) < v16) {
                                                                                                                        v138 = v138
                                                                                                                        var v183 int = (-(fi141))
                                                                                                                        _ = v183
                                                                                                                    } else {
                                                                                                                        fmt.Println(v16, v6, v136)
                                                                                                                        fmt.Println(v138, fi141, v142)
                                                                                                                        delete(v5, "a")
                                                                                                                        v163 += (v143 << (((^(v154))) & 63))
                                                                                                                    }
                                                                                                                case "hello":
                                                                                                                    var v184 int64 = -9223372036854775808
                                                                                                                    _ = v184
                                                                                                                default:
                                                                                                                    var v185 float32 = 1e10
                                                                                                                    var v186 int8 = genFn0((-(v16)), v15)
                                                                                                                    _ = v185
                                                                                                                    _ = v186
                                                                                                                }
                                                                                                            } else {
                                                                                                                var v187 string = v3
                                                                                                                var v188 string = v3
                                                                                                                v159 = v13
                                                                                                                switch ((((v1 + "hello") + ("a" + v187)) + (v1 + ((v3 + ("\n\t" + v187)) + ((v3 + v187) + v187)))) + "\n\t") {
                                                                                                                case "":
                                                                                                                    v154 = v2
                                                                                                                }
                                                                                                                var v189 uint = uint(v181)
                                                                                                                _ = v187
                                                                                                                _ = v188
                                                                                                                _ = v189
                                                                                                            }
                                                                                                        case "foo bar":
                                                                                                            fmt.Println(v138, v136, v163)
                                                                                                        }
                                                                                                        v15 = (v15 / ((v15) | 1))
                                                                                                    }
                                                                                                    _ = v181
                                                                                                }
                                                                                            }
                                                                                        }
                                                                                        _ = v167
                                                                                        _ = v168
                                                                                    } else {
                                                                                        switch (((v1 + v3) + v3) + v1) {
                                                                                        case "":
                                                                                            switch "a" {
                                                                                            case "\n\t":
                                                                                                fmt.Println(v163)
                                                                                            case "hello":
                                                                                                var v190 uint8 = 127
                                                                                                _ = v190
                                                                                            }
                                                                                        }
                                                                                        var v191 string = (v3 + "foo bar")
                                                                                        fmt.Println(v191)
                                                                                        if len(v4) >= 2 {
                                                                                            v4 = v4[1:len(v4)]
                                                                                        }
                                                                                        if (("" <= v1) && ((fi137 + (fi141 + int(v159))) < fi137)) {
                                                                                            v138 += int16(v15)
                                                                                            v192 := GenS20{F0: v6, F1: 1, F2: 49}
                                                                                            _ = v192
                                                                                        } else {
                                                                                            var v193 int = int(v143)
                                                                                            var v194 int16 = (v138 | -354)
                                                                                            _ = v193
                                                                                            _ = v194
                                                                                        }
                                                                                        _ = v191
                                                                                    }
                                                                                    var v195 bool = (v160 < 65535)
                                                                                    _ = v195
                                                                                case "foo bar":
                                                                                    v3 = ""
                                                                                }
                                                                                var v196 string = ((v3 + v3) + v1)
                                                                                _ = v163
                                                                                _ = v196
                                                                            }
                                                                            _ = v162
                                                                        }
                                                                        if true {
                                                                            var v197 int64 = v15
                                                                            v198 := GenS17{F0: v160}
                                                                            v13 += (^(v159))
                                                                            _ = v197
                                                                            _ = v198
                                                                        }
                                                                        _ = v159
                                                                        _ = v160
                                                                        _ = v161
                                                                    }
                                                                }
                                                            }
                                                            _ = v154
                                                        }
                                                    }
                                                    v143 = 4294967295
                                                    if (142 <= 43) {
                                                        v199 := GenS13{F0: (true || (0.1 >= 3.14159))}
                                                        var v200 string = (v3 + ("foo bar" + v1))
                                                        v143 += 2
                                                        var v201 int = fi137
                                                        _ = v199
                                                        _ = v200
                                                        _ = v201
                                                    } else {
                                                        v136 = v143
                                                        if (-128 > (v6 + -128)) {
                                                            v2 -= v2
                                                            _ = len(v5)
                                                        }
                                                        v202 := GenS21{F0: v1}
                                                        _ = v202
                                                    }
                                                    if (false || (float64(0.0) <= 0.1)) {
                                                        var v203 uint64 = 4294967295
                                                        v204 := GenS22{F0: v13, F1: -1}
                                                        _ = v203
                                                        _ = v204
                                                    }
                                                }
                                            }
                                            var v205 uint64 = v13
                                            _ = v205
                                        }
                                    case "\n\t":
                                        v206 := GenS23{F0: (v138 & (^(v138))), F1: 991}
                                        _ = v206
                                    }
                                    _ = v145
                                }
                            case -58:
                                switch "a" {
                                case "\n\t":
                                    switch ((("\n\t" + "a") + "") + (v3 + v1)) {
                                    case "\n\t":
                                        v4 = append(v4, (v143 & v136))
                                    case "a":
                                        switch v3 {
                                        case "\n\t":
                                            switch v1 {
                                            case "foo bar":
                                                v3 = "\n\t"
                                            case "\n\t":
                                                switch fi137 {
                                                case -419:
                                                    v4 = append(v4, v143)
                                                case 128:
                                                    if (0.1 >= 1.5) {
                                                        switch (fi137 - int(v136)) {
                                                        case 457:
                                                            var v207 float32 = -2.25
                                                            _ = v207
                                                        case 0:
                                                            v4 = append(v4, v143)
                                                        case 829:
                                                            var v208 string = ((v3 + ((v1 + (v3 + "foo bar")) + v1)) + (v3 + v3))
                                                            fmt.Println(v143, v13)
                                                            v209 := GenS9{F0: (v15 | 9223372036854775807), F1: (v13 | 4294967295), F2: v15}
                                                            _ = v208
                                                            _ = v209
                                                        default:
                                                            fmt.Println(v136, v6)
                                                            switch (fi137 * fi137) {
                                                            case 128:
                                                                switch "hello" {
                                                                case "hello":
                                                                    switch fi141 {
                                                                    case 4294967295:
                                                                        fmt.Println(v16, v16)
                                                                    }
                                                                    var v210 uint32 = v136
                                                                    v211 := GenS4{F0: (v142 >> ((2) & 63)), F1: ("a" + ("" + ((v3 + (("" + "hello") + (v1 + v1))) + ("a" + v1))))}
                                                                    fmt.Println(v3)
                                                                    _ = v210
                                                                    _ = v211
                                                                default:
                                                                    var v212 int32 = int32(v15)
                                                                    _ = v212
                                                                }
                                                            default:
                                                                var v213 string = ""
                                                                _ = v213
                                                            }
                                                        }
                                                    } else {
                                                        v142 = v142
                                                        v143 += v143
                                                        v138 -= (v138 >> (((v2 & v2)) & 63))
                                                    }
                                                case 127:
                                                    switch v1 {
                                                    case "\n\t":
                                                        fmt.Println(v2)
                                                    }
                                                default:
                                                    v1 = v1
                                                }
                                            default:
                                                if false {
                                                    if len(v4) >= 2 {
                                                        v4 = v4[1:len(v4)]
                                                    }
                                                    v214 := GenS24{F0: (^(v15)), F1: v13, F2: 42}
                                                    var v215 uint16 = 2
                                                    _ = v214
                                                    _ = v215
                                                } else {
                                                    v216 := v5[("foo bar" + "")]
                                                    fmt.Println(v143, v216, v6)
                                                    if (18446744073709551615 <= v13) {
                                                        var v217 float64 = (0.1 + (float64(3.14159) * 0.1))
                                                        _ = v217
                                                    }
                                                    _ = v216
                                                }
                                                var v218 int32 = (v142 * (v142 * v142))
                                                _ = v218
                                            }
                                        case "hello":
                                            v219 := v5["a"]
                                            _ = v219
                                        case "a":
                                            _ = len(v5)
                                        default:
                                            v3 = v1
                                        }
                                    }
                                    var v220 int8 = genFn0(v16, 145)
                                    _ = len(v5)
                                    _ = v220
                                case "":
                                    v1 = v1
                                    fmt.Println(fi141, v15)
                                }
                            case 32767:
                                var v221 bool = false
                                _ = v221
                            default:
                                v3 = ("" + v1)
                                v5[(v3 + v1)] = v143
                                var v222 uint64 = v13
                                _ = v222
                            }
                            _ = fi141
                            _ = v142
                            _ = v143
                        }
                    case "hello":
                        v223 := GenS25{F0: 65535, F1: -9223372036854775808}
                        _ = v223
                    default:
                        var v224 uint32 = 4294967295
                        _ = v224
                    }
                default:
                    switch fi137 {
                    case 128:
                        v225 := GenS23{F0: v138, F1: 0}
                        _ = v225
                    }
                }
                _ = fi137
                _ = v138
            }
            _ = v136
        }
        var v226 uint32 = 1
        _ = v226
    }
    v13 -= (v13 | 3)
    v227 := GenS15{F0: v3, F1: (v13 % (((v13 & (v13 - 2))) | 1))}
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v227
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
