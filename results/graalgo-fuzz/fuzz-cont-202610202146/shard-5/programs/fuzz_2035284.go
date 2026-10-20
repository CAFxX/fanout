package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int16; F1 uint32 }
type GenS3 struct { F0 uint64; F1 int64; F2 uint }
type GenS4 struct { F0 uint32; F1 int64 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 string; F1 uint }
type GenS8 struct { F0 int; F1 uint64 }
type GenS9 struct { F0 int32; F1 int8 }
type GenS10 struct { F0 int32 }
type GenS11 struct { F0 string }
type GenS12 struct { F0 int16; F1 uint8; F2 int64 }
type GenS13 struct { F0 uint64 }
type GenS14 struct { F0 uint8 }
type GenS15 struct { F0 int; F1 bool; F2 int32 }
type GenS16 struct { F0 bool; F1 uint64; F2 uint32 }
type GenS17 struct { F0 int8; F1 string; F2 bool }
type GenS18 struct { F0 string; F1 uint64 }
type GenS19 struct { F0 bool }
type GenS20 struct { F0 string }
type GenS21 struct { F0 uint16 }

func main() {
    v1 := make(map[int]bool)
    var v2 string = ("a" + "\n\t")
    var v3 int32 = 42
    var v4 int64 = 0
    var v5 int64 = (v4 + (v4 - 42))
    v6 := []int{811}
    var v7 uint16 = 1
    var v8 int32 = (^(v3))
    v5 += (-(v5))
    v8 += (v8 * (v8 & v8))
    v8 -= v8
    if (!true) {
        for fi9 := 0; fi9 < 13; fi9++ {
            if true {
                if false {
                    if ((true || ((!false) || false)) || false) {
                        switch (v2 + (v2 + (v2 + v2))) {
                        case "a":
                            v1[(fi9 / ((fi9) | 1))] = (true || (true && (!false)))
                        case "":
                            v10 := GenS1{F0: uint64(v7)}
                            _ = v10
                        case "foo bar":
                            switch fi9 {
                            case -9223372036854775808:
                                fmt.Println(v2, fi9)
                            case -752:
                                fmt.Println(v4, v4)
                            }
                        }
                        switch v2 {
                        case "":
                            var v11 int64 = (^(v5))
                            _ = v11
                        case "\n\t":
                            v2 = v2
                        case "a":
                            if (-2.25 > 1.5) {
                                for fi12 := 0; fi12 < 9; fi12++ {
                                    fmt.Println(v3, v2)
                                    fmt.Println(v4, v7, v2)
                                    v6 = append(v6, -1)
                                    _ = fi12
                                }
                                var v13 uint32 = 2147483647
                                if ((v13 == v13) || true) {
                                    for fi14 := 0; fi14 < 15; fi14++ {
                                        if true {
                                            v15 := GenS2{F0: 189, F1: 0}
                                            _ = len(v6)
                                            if (!false) {
                                                if false {
                                                    if (2 >= 128) {
                                                        v16 := v1[int(v3)]
                                                        v16 = false
                                                        var v17 int16 = int16(v7)
                                                        _ = v16
                                                        _ = v17
                                                    } else {
                                                        v18 := GenS3{F0: 0, F1: v5, F2: 42}
                                                        v19 := GenS4{F0: 192, F1: 42}
                                                        _ = v18
                                                        _ = v19
                                                    }
                                                    v7 += 724
                                                }
                                                if false {
                                                    v20 := GenS5{F0: v5}
                                                    var v21 string = "hello"
                                                    _ = v20
                                                    _ = v21
                                                } else {
                                                    v13 = (v13 & (v13 & v13))
                                                    var v22 uint64 = 1
                                                    v2 = (v2 + "")
                                                    _ = v22
                                                }
                                                if false {
                                                    v4 = 9223372036854775807
                                                    v23 := v1[4294967295]
                                                    v24 := GenS6{F0: 1}
                                                    _ = v23
                                                    _ = v24
                                                }
                                            }
                                            if (91 < 208) {
                                                _ = len(v6)
                                            }
                                            _ = v15
                                        }
                                        _ = fi14
                                    }
                                    var v25 string = v2
                                    _ = v25
                                }
                                _ = v13
                            }
                        }
                    } else {
                        v7 -= v7
                    }
                    if false {
                        if false {
                            switch "hello" {
                            case "":
                                if len(v6) > 0 {
                                    v6[0] = fi9
                                }
                            case "hello":
                                for fi26 := 0; fi26 < 11; fi26++ {
                                    v8 = 252
                                    if false {
                                        _ = len(v6)
                                        var v27 bool = false
                                        _ = v27
                                    }
                                    _ = fi26
                                }
                            case "a":
                                v1[(fi9 / (((fi9 >> ((1) & 63))) | 1))] = (("\n\t" + "\n\t") <= v2)
                            default:
                                var v28 int8 = -1
                                _ = v28
                            }
                        }
                        var v29 float32 = -2.25
                        _ = v29
                    }
                    var v30 uint = 42
                    _ = v30
                }
                var v31 int = fi9
                if false {
                    v32 := GenS2{F0: 0, F1: 1}
                    _ = v32
                } else {
                    for fi33 := 0; fi33 < 4; fi33++ {
                        switch (v2 + (v2 + (v2 + ("" + "hello")))) {
                        case "hello":
                            var v34 uint64 = 18446744073709551615
                            _ = v34
                        case "a":
                            v31 = 32768
                        case "":
                            if (true || true) {
                                v35 := GenS7{F0: ((v2 + (v2 + v2)) + ""), F1: 255}
                                if (false && true) {
                                    var v36 int16 = 1
                                    _ = v36
                                }
                                _ = v35
                            } else {
                                if true {
                                    v37 := v1[(v31 - fi33)]
                                    if (42 >= 100) {
                                        if v37 {
                                            v5 -= v4
                                            if (!(!v37)) {
                                                if (21 > 255) {
                                                    fmt.Println(v5, v3, v5)
                                                } else {
                                                    v38 := GenS1{F0: 2}
                                                    v39 := GenS8{F0: 32767, F1: 0}
                                                    _ = v38
                                                    _ = v39
                                                }
                                                var v40 uint16 = (v7 / ((1) | 1))
                                                _ = v40
                                            } else {
                                                if len(v6) >= 2 {
                                                    v6 = v6[1:len(v6)]
                                                }
                                                fmt.Println(fi9)
                                                v7 = (^(v7))
                                            }
                                            var v41 float32 = (0.0 + float32(0.1))
                                            _ = v41
                                        }
                                        switch (^(fi33)) {
                                        case 1:
                                            v42 := GenS6{F0: (v5 | v5)}
                                            _ = v42
                                        default:
                                            if true {
                                                v43 := GenS9{F0: (v3 >> ((255) & 63)), F1: 116}
                                                v44 := GenS7{F0: v2, F1: 841}
                                                _ = v43
                                                _ = v44
                                            }
                                        }
                                        var v45 uint32 = 928
                                        _ = v45
                                    } else {
                                        if len(v6) >= 2 {
                                            v6 = v6[1:len(v6)]
                                        }
                                        v6 = append(v6, v31)
                                    }
                                    _ = v37
                                }
                            }
                        default:
                            if ((v5 % (((v5 - (-(v4)))) | 1)) <= (^(v4))) {
                                if true {
                                    var v46 int64 = (v4 << ((191) & 63))
                                    fmt.Println(v4, v46)
                                    v46 += (v46 << ((969) & 63))
                                    _ = v46
                                }
                                fmt.Println(v31, v4, v3)
                                v5 += v4
                            } else {
                                var v47 int8 = -1
                                v3 += (v3 | -749)
                                _ = v47
                            }
                        }
                        var v48 int = fi9
                        var v49 float32 = ((((((0.1 / 1.0) / 1.0) - 1e10) + float32(0.1)) + 0.1) * (-2.25 + (0.1 - (1e10 / 1e10))))
                        _ = fi33
                        _ = v48
                        _ = v49
                    }
                }
                _ = v31
            } else {
                fmt.Println(v5, v4)
                for fi50 := 0; fi50 < 11; fi50++ {
                    switch v2 {
                    case "a":
                        v8 -= (v8 - -2147483648)
                        v51 := GenS8{F0: fi50, F1: 517}
                        var v52 int = -394
                        _ = v51
                        _ = v52
                    }
                    fmt.Println(fi9)
                    v2 = (v2 + v2)
                    _ = fi50
                }
                for fi53 := 0; fi53 < 14; fi53++ {
                    if ((v4 & v4) != v4) {
                        v1[(fi53 % ((-538) | 1))] = false
                        if (!(!((!(18446744073709551615 <= 0)) && false))) {
                            if false {
                                var v54 int32 = 0
                                _ = v54
                            }
                            var v55 bool = (0 <= 42)
                            _ = v55
                        } else {
                            delete(v1, (fi9 | (-(fi53))))
                            v5 += v5
                        }
                    } else {
                        var v56 int8 = -128
                        _ = v56
                    }
                    if false {
                        v57 := GenS3{F0: 153, F1: v5, F2: 108}
                        if (2 < 18446744073709551615) {
                            if true {
                                if ((-1 < v4) && (873 >= 0)) {
                                    _ = len(v1)
                                    var v58 bool = (int8(v4) <= -95)
                                    var v59 uint8 = 0
                                    var v60 float32 = 3.14159
                                    _ = v58
                                    _ = v59
                                    _ = v60
                                }
                                v61 := GenS10{F0: 0}
                                fmt.Println(v2, v3)
                                if false {
                                    _ = len(v1)
                                    v62 := GenS6{F0: -610}
                                    var v63 float64 = 0.1
                                    fmt.Println(v8, v63, v63)
                                    var v64 int16 = -1
                                    if (!(!true)) {
                                        var v65 int8 = 127
                                        switch v2 {
                                        case "hello":
                                            fmt.Println(v4, v2)
                                        case "\n\t":
                                            var v66 uint8 = 128
                                            _ = v66
                                        }
                                        var v67 int16 = -1
                                        if len(v6) > 0 {
                                            v6[0] = fi53
                                        }
                                        switch ("a" + "a") {
                                        case "\n\t":
                                            switch (v2 + (v2 + (("\n\t" + v2) + v2))) {
                                            case "\n\t":
                                                v1[-288] = true
                                            case "foo bar":
                                                v68 := v1[int(v5)]
                                                _ = v68
                                            default:
                                                var v69 int = (fi53 ^ int(v3))
                                                _ = v69
                                            }
                                        default:
                                            v70 := GenS9{F0: 42, F1: (v65 & v65)}
                                            _ = v70
                                        }
                                        _ = v65
                                        _ = v67
                                    } else {
                                        var v71 float32 = 0.0
                                        if len(v6) >= 2 {
                                            v6 = v6[1:len(v6)]
                                        }
                                        if len(v6) >= 2 {
                                            v6 = v6[1:len(v6)]
                                        }
                                        _ = v71
                                    }
                                    _ = v62
                                    _ = v63
                                    _ = v64
                                } else {
                                    v6 = append(v6, 827)
                                }
                                _ = v61
                            } else {
                                var v72 int32 = v8
                                v3 = v3
                                v72 = (v8 - -2147483648)
                                _ = v72
                            }
                            switch (^(fi9)) {
                            case -194:
                                fmt.Println(v7, v4, v7)
                            case 127:
                                v5 = 796
                            case -1:
                                v73 := GenS11{F0: ("" + v2)}
                                _ = v73
                            default:
                                switch fi9 {
                                case 4294967295:
                                    var v74 uint16 = v7
                                    _ = v74
                                case 679:
                                    if (0.0 != float64(3.14159)) {
                                        switch fi53 {
                                        case 32768:
                                            switch 452 {
                                            case 656:
                                                v75 := GenS12{F0: 42, F1: 42, F2: (v5 | (v4 + (v5 * v5)))}
                                                _ = v75
                                            case 516:
                                                fmt.Println(v2)
                                            case 128:
                                                v7 += (v7 | (v7 >> ((4294967295) & 63)))
                                            }
                                        default:
                                            fmt.Println(v3)
                                        }
                                        var v76 bool = false
                                        var v77 int32 = v3
                                        var v78 uint64 = 778
                                        v79 := v1[fi9]
                                        _ = v76
                                        _ = v77
                                        _ = v78
                                        _ = v79
                                    } else {
                                        v5 += (v5 - (v5 | -417))
                                    }
                                default:
                                    if (false || (1 >= 316)) {
                                        v80 := GenS13{F0: 1}
                                        var v81 float64 = 1.5
                                        fmt.Println(v81)
                                        _ = v80
                                        _ = v81
                                    } else {
                                        var v82 uint8 = 148
                                        fmt.Println(fi53, v2)
                                        if true {
                                            switch (fi53 >> ((18446744073709551615) & 63)) {
                                            case 550:
                                                v5 = (v5 >> ((359) & 63))
                                            case -897:
                                                v6 = append(v6, (fi9 >> ((4294967295) & 63)))
                                            case 9223372036854775807:
                                                if false {
                                                    if true {
                                                        if (false || true) {
                                                            fmt.Println(v5, v3)
                                                        }
                                                        if false {
                                                            v8 -= v3
                                                            if ((188 <= 2) && false) {
                                                                fmt.Println(fi9, fi53)
                                                                if false {
                                                                    v83 := GenS14{F0: v82}
                                                                    v7 -= 32767
                                                                    switch "a" {
                                                                    case "":
                                                                        v84 := GenS15{F0: (fi53 | 0), F1: false, F2: (v3 >> ((12) & 63))}
                                                                        _ = v84
                                                                    case "foo bar":
                                                                        switch "\n\t" {
                                                                        case "a":
                                                                            fmt.Println(v5, v82, v8)
                                                                        case "hello":
                                                                            switch ((v2 + "hello") + "hello") {
                                                                            case "foo bar":
                                                                                v3 += v8
                                                                            case "a":
                                                                                v82 = v82
                                                                            default:
                                                                                v6 = append(v6, fi9)
                                                                            }
                                                                        case "\n\t":
                                                                            v3 += v8
                                                                        default:
                                                                            v8 -= (v3 / ((v3) | 1))
                                                                        }
                                                                    }
                                                                    v3 = (v8 % (((v3 | -2147483648)) | 1))
                                                                    if false {
                                                                        var v85 bool = true
                                                                        _ = v85
                                                                    } else {
                                                                        v86 := GenS16{F0: ((v8 / ((v8) | 1)) != v3), F1: 528, F2: 2}
                                                                        v87 := GenS17{F0: 100, F1: v2, F2: false}
                                                                        if false {
                                                                            var v88 uint16 = v7
                                                                            v3 -= 42
                                                                            _ = v88
                                                                        } else {
                                                                            if (!true) {
                                                                                v3 += v3
                                                                                v7 = v7
                                                                            }
                                                                        }
                                                                        _ = v86
                                                                        _ = v87
                                                                    }
                                                                    switch fi53 {
                                                                    case -1:
                                                                        switch 2147483647 {
                                                                        case 32767:
                                                                            fmt.Println(v8, fi9)
                                                                        }
                                                                    case -367:
                                                                        if len(v6) > 0 {
                                                                            v6[0] = fi9
                                                                        }
                                                                    default:
                                                                        if (65535 == 990) {
                                                                            v82 -= (v82 & v82)
                                                                        }
                                                                    }
                                                                    _ = v83
                                                                } else {
                                                                    var v89 int64 = v5
                                                                    _ = v89
                                                                }
                                                            } else {
                                                                delete(v1, fi53)
                                                            }
                                                        }
                                                        var v90 float32 = 0.0
                                                        _ = v90
                                                    } else {
                                                        v5 = 9223372036854775807
                                                        v91 := GenS18{F0: ("hello" + (v2 + v2)), F1: uint64(v82)}
                                                        var v92 float32 = ((1e10 - 0.0) - 0.1)
                                                        _ = v91
                                                        _ = v92
                                                    }
                                                    v7 = v7
                                                    v93 := GenS7{F0: "", F1: 652}
                                                    _ = len(v6)
                                                    fmt.Println(v4, fi9)
                                                    _ = v93
                                                }
                                            }
                                            switch "hello" {
                                            case "foo bar":
                                                _ = len(v1)
                                            case "hello":
                                                var v94 uint64 = 2
                                                _ = v94
                                            case "\n\t":
                                                if false {
                                                    var v95 int = (^(fi9))
                                                    v2 = v2
                                                    _ = v95
                                                } else {
                                                    if true {
                                                        switch 0 {
                                                        case 1:
                                                            if false {
                                                                if ((0.1 * -2.25) > 1.5) {
                                                                    if (65535 >= 969) {
                                                                        v8 -= (^(v3))
                                                                        switch fi53 {
                                                                        case 42:
                                                                            if true {
                                                                                var v96 float32 = 1.5
                                                                                v5 = (v4 % (((v4 | v5)) | 1))
                                                                                _ = v96
                                                                            }
                                                                        }
                                                                    }
                                                                } else {
                                                                    v97 := GenS19{F0: (!false)}
                                                                    _ = v97
                                                                }
                                                            } else {
                                                                v98 := GenS20{F0: "\n\t"}
                                                                _ = v98
                                                            }
                                                        case 350:
                                                            var v99 uint64 = 2
                                                            _ = v99
                                                        }
                                                        var v100 bool = (int16(v7) > -1)
                                                        var v101 int16 = -414
                                                        var v102 int8 = 127
                                                        var v103 int = fi9
                                                        _ = v100
                                                        _ = v101
                                                        _ = v102
                                                        _ = v103
                                                    }
                                                    var v104 uint8 = (v82 & v82)
                                                    _ = v104
                                                }
                                            }
                                            switch (fi53 / ((fi9) | 1)) {
                                            case 306:
                                                v8 -= v8
                                            case 826:
                                                v7 = (v7 + 2)
                                            default:
                                                switch "" {
                                                case "a":
                                                    fmt.Println(v8, fi53, v8)
                                                case "":
                                                    v105 := GenS21{F0: v7}
                                                    _ = v105
                                                }
                                            }
                                            var v106 uint64 = 0
                                            v8 = (v3 & (v3 >> ((4294967295) & 63)))
                                            _ = v106
                                        }
                                        _ = v82
                                    }
                                }
                            }
                        }
                        if (!(false && (true || (true || (true || false))))) {
                            v8 = (v8 ^ v8)
                            var v107 float32 = (1.5 + -2.25)
                            switch (fi53 << ((255) & 63)) {
                            case 9223372036854775807:
                                v5 = v5
                            }
                            var v108 float32 = v107
                            fmt.Println(v5)
                            _ = v107
                            _ = v108
                        }
                        var v109 int32 = (v3 / ((int32(v7)) | 1))
                        _ = v109
                        _ = v57
                    }
                    v110 := GenS6{F0: v4}
                    v2 = ("" + "hello")
                    var v111 int32 = v3
                    _ = fi53
                    _ = v110
                    _ = v111
                }
            }
            _ = len(v6)
            _ = fi9
        }
        fmt.Println(v2, v5)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
