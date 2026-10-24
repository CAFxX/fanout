package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int; F1 uint32; F2 uint16 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 int32; F1 uint64; F2 string }
type GenS7 struct { F0 uint32; F1 int16 }
type GenS8 struct { F0 uint16; F1 int16 }
type GenS9 struct { F0 uint; F1 int32 }
type GenS10 struct { F0 int64; F1 string; F2 int64 }
type GenS11 struct { F0 int; F1 bool }
type GenS12 struct { F0 uint; F1 int64; F2 int64 }
type GenS13 struct { F0 bool; F1 uint16 }
type GenS14 struct { F0 uint8; F1 uint32; F2 uint64 }
type GenS15 struct { F0 int16 }
type GenS16 struct { F0 uint8 }
type GenS17 struct { F0 uint32; F1 uint8 }
type GenS18 struct { F0 int; F1 uint8; F2 int8 }
type GenS19 struct { F0 int64; F1 int32 }
func genFn0() uint8 {
    var v6 uint = 860
    fmt.Println(v6)
    _ = v6
    return 1
}
func genFn1() uint {
    var v7 int = 32767
    for fi8 := 0; fi8 < 6; fi8++ {
        fmt.Println(v7, v7)
        _ = fi8
    }
    switch 1 {
    case -114:
        var v9 uint16 = 32767
        _ = v9
    case -670:
        var v10 uint32 = 2147483647
        _ = v10
    }
    _ = v7
    return 42
}

func main() {
    v1 := make(map[uint32]int)
    var v2 int16 = 0
    v3 := make(map[int]int)
    v4 := make(map[uint32]uint32)
    v5 := []int{278, 4294967295, -604}
    if true {
        for fi11 := 0; fi11 < 13; fi11++ {
            var v12 int16 = (v2 << ((genFn1()) & 63))
            if len(v5) > 0 {
                v5[0] = fi11
            }
            for fi13 := 0; fi13 < 9; fi13++ {
                var v14 uint32 = 527
                _ = fi13
                _ = v14
            }
            _ = fi11
            _ = v12
        }
        var v15 bool = (-32768 < (v2 / (((-(v2))) | 1)))
        v15 = v15
        _ = v15
    } else {
        var v16 uint16 = 1
        switch ("hello" + "\n\t") {
        case "a":
            fmt.Println(v16)
        case "foo bar":
            delete(v1, 870)
        default:
            fmt.Println(v2, v16)
        }
        _ = v16
    }
    var v17 uint8 = 71
    v18 := GenS1{F0: 9223372036854775807}
    v5 = append(v5, -2147483648)
    var v19 int64 = 42
    if true {
        v19 = (v19 - v19)
        var v20 uint64 = 585
        _ = v20
    } else {
        var v21 uint16 = 0
        _ = v21
    }
    for fi22 := 0; fi22 < 11; fi22++ {
        v17 = 124
        v23 := GenS2{F0: v17}
        fmt.Println(v2)
        _ = fi22
        _ = v23
    }
    var v24 uint64 = 753
    v25 := GenS2{F0: (v17 ^ 127)}
    fmt.Println(v24, v17, v24)
    if true {
        if true {
            var v26 float32 = 1e10
            if len(v5) > 0 {
                v5[0] = -9223372036854775808
            }
            _ = v26
        }
        for fi27 := 0; fi27 < 18; fi27++ {
            var v28 uint16 = 559
            for fi29 := 0; fi29 < 2; fi29++ {
                if ("" != (("\n\t" + ("\n\t" + "hello")) + "a")) {
                    switch "foo bar" {
                    case "":
                        var v30 float32 = 0.1
                        _ = v30
                    case "foo bar":
                        v19 = v19
                    default:
                        if false {
                            v28 = v28
                            v17 += v17
                        } else {
                            fmt.Println(v28, v28, fi29)
                            var v31 int8 = -1
                            if (!(v17 == (v17 >> ((255) & 63)))) {
                                v28 = (^(v28))
                            } else {
                                if false {
                                    if false {
                                        var v32 float64 = (((1e10 / (3.14159 + -2.25)) + -2.25) * float64(-2.25))
                                        _ = v32
                                    }
                                    fmt.Println(v31, v31)
                                    if true {
                                        fmt.Println(fi29, v2)
                                        _ = len(v4)
                                        switch (fi27 >> ((255) & 63)) {
                                        case -9223372036854775808:
                                            v31 = v31
                                        default:
                                            var v33 uint8 = (v17 | (v17 % ((v17) | 1)))
                                            _ = v33
                                        }
                                    } else {
                                        if false {
                                            delete(v1, 0)
                                            var v34 uint16 = (v28 * 972)
                                            var v35 float64 = 0.0
                                            _ = v34
                                            _ = v35
                                        }
                                    }
                                }
                            }
                            switch fi27 {
                            case 0:
                                if false {
                                    if (genFn0() > (v17 << ((uint(v28)) & 63))) {
                                        v28 = 146
                                        _ = len(v3)
                                        switch fi27 {
                                        case 345:
                                            var v36 int64 = v19
                                            _ = v36
                                        case -2147483648:
                                            v17 -= genFn0()
                                        case 128:
                                            switch 281 {
                                            case 733:
                                                switch "hello" {
                                                case "a":
                                                    _ = len(v4)
                                                case "":
                                                    v37 := GenS2{F0: v17}
                                                    if len(v5) >= 2 {
                                                        v5 = v5[1:len(v5)]
                                                    }
                                                    _ = v37
                                                case "foo bar":
                                                    v24 = v24
                                                    v5 = append(v5, fi27)
                                                    if true {
                                                        switch "\n\t" {
                                                        case "":
                                                            if true {
                                                                fmt.Println(v31)
                                                                v28 = 553
                                                                if len(v5) >= 2 {
                                                                    v5 = v5[1:len(v5)]
                                                                }
                                                            }
                                                        default:
                                                            if true {
                                                                v38 := GenS3{F0: 42, F1: 861, F2: v28}
                                                                switch 0 {
                                                                case 32767:
                                                                    if len(v5) >= 2 {
                                                                        v5 = v5[1:len(v5)]
                                                                    }
                                                                default:
                                                                    fmt.Println(v28, v19)
                                                                    switch (fi27 & (fi27 * fi27)) {
                                                                    case -946:
                                                                        var v39 int16 = -301
                                                                        _ = v39
                                                                    case 42:
                                                                        fmt.Println(fi27)
                                                                        v19 += (v19 * v19)
                                                                    case 32768:
                                                                        switch (fi27 >> ((742) & 63)) {
                                                                        case -9223372036854775808:
                                                                            if (1 >= 2147483647) {
                                                                                switch "hello" {
                                                                                case "\n\t":
                                                                                    fmt.Println(v28, v17, v28)
                                                                                default:
                                                                                    var v40 float64 = (float64(1e10) / 1.0)
                                                                                    _ = v40
                                                                                }
                                                                                var v41 uint8 = (^(v17))
                                                                                _ = v41
                                                                            }
                                                                        case 4294967295:
                                                                            if true {
                                                                                if (531 <= 276) {
                                                                                    v42 := v1[997]
                                                                                    if true {
                                                                                        v42 = int(v24)
                                                                                        var v43 int = (fi29 >> ((702) & 63))
                                                                                        if true {
                                                                                            if true {
                                                                                                switch int(v17) {
                                                                                                case 2147483647:
                                                                                                    v42 = (v42 % (((fi29 % ((32768) | 1))) | 1))
                                                                                                case -799:
                                                                                                    v43 -= (v42 + (v43 << ((689) & 63)))
                                                                                                case 302:
                                                                                                    var v44 uint16 = (v28 * 782)
                                                                                                    _ = v44
                                                                                                }
                                                                                                v4[2] = 2147483647
                                                                                            }
                                                                                            if true {
                                                                                                var v45 uint8 = v17
                                                                                                fmt.Println(v31, v28, v17)
                                                                                                _ = v45
                                                                                            } else {
                                                                                                v46 := GenS1{F0: v19}
                                                                                                _ = len(v5)
                                                                                                _ = v46
                                                                                            }
                                                                                            var v47 float32 = (float32(1e10) - (3.14159 - 1e10))
                                                                                            _ = v47
                                                                                        } else {
                                                                                            var v48 uint32 = uint32(v24)
                                                                                            if (2 > v24) {
                                                                                                switch v43 {
                                                                                                case -99:
                                                                                                    if (0 != -712) {
                                                                                                        var v49 uint32 = v48
                                                                                                        var v50 float32 = float32((-2.25 + (1e10 / ((((-2.25 * (-2.25 - ((float64(3.14159) * (3.14159 + -2.25)) / 1.5))) + 1e10) + ((-2.25 + ((0.1 * -2.25) + 1.5)) / 1e10)) + 3.14159))))
                                                                                                        _ = v49
                                                                                                        _ = v50
                                                                                                    }
                                                                                                    var v51 uint8 = v17
                                                                                                    v28 = v28
                                                                                                    _ = v51
                                                                                                case -9223372036854775808:
                                                                                                    if len(v5) >= 2 {
                                                                                                        v5 = v5[1:len(v5)]
                                                                                                    }
                                                                                                case 128:
                                                                                                    _ = len(v4)
                                                                                                }
                                                                                                v31 = (v31 & v31)
                                                                                                switch ("hello" + ("a" + ("foo bar" + ""))) {
                                                                                                case "\n\t":
                                                                                                    v24 = 4294967295
                                                                                                case "":
                                                                                                    v52 := v3[fi29]
                                                                                                    _ = v52
                                                                                                case "hello":
                                                                                                    v53 := GenS4{F0: -850}
                                                                                                    _ = v53
                                                                                                }
                                                                                            }
                                                                                            var v54 uint64 = 2
                                                                                            _ = v48
                                                                                            _ = v54
                                                                                        }
                                                                                        _ = v43
                                                                                    } else {
                                                                                        var v55 float64 = (3.14159 * 0.1)
                                                                                        _ = v55
                                                                                    }
                                                                                    if false {
                                                                                        fmt.Println(v31, v28)
                                                                                        v28 = 0
                                                                                    } else {
                                                                                        v5 = append(v5, v42)
                                                                                        var v56 int64 = v19
                                                                                        v5 = append(v5, (v42 | (fi29 << ((2) & 63))))
                                                                                        v24 = uint64(v28)
                                                                                        v5 = append(v5, fi27)
                                                                                        var v57 uint = genFn1()
                                                                                        v2 = (v2 / ((503) | 1))
                                                                                        _ = v56
                                                                                        _ = v57
                                                                                    }
                                                                                    _ = v42
                                                                                }
                                                                                v19 -= 302
                                                                            } else {
                                                                                v31 = v31
                                                                            }
                                                                        default:
                                                                            if (3.14159 >= 0.0) {
                                                                                switch fi27 {
                                                                                case 1:
                                                                                    fmt.Println(fi27, fi29, v24)
                                                                                case 2147483647:
                                                                                    v31 = v31
                                                                                    if false {
                                                                                        if false {
                                                                                            var v58 int16 = (v2 ^ v2)
                                                                                            v59 := GenS2{F0: v17}
                                                                                            _ = v58
                                                                                            _ = v59
                                                                                        }
                                                                                    }
                                                                                    v5 = append(v5, (fi27 >> ((genFn1()) & 63)))
                                                                                case 4294967295:
                                                                                    switch "a" {
                                                                                    case "a":
                                                                                        var v60 float64 = (1.5 + 1.5)
                                                                                        var v61 int = 0
                                                                                        _ = v60
                                                                                        _ = v61
                                                                                    default:
                                                                                        if false {
                                                                                            if len(v5) > 0 {
                                                                                                v5[0] = fi27
                                                                                            }
                                                                                            v17 += v17
                                                                                            v5 = append(v5, (fi27 / (((fi27 - 624)) | 1)))
                                                                                            fmt.Println(v28, fi27)
                                                                                            var v62 uint64 = (v24 >> ((929) & 63))
                                                                                            _ = v62
                                                                                        }
                                                                                    }
                                                                                default:
                                                                                    var v63 uint16 = (v28 << ((1) & 63))
                                                                                    _ = v63
                                                                                }
                                                                            } else {
                                                                                v64 := GenS2{F0: v17}
                                                                                _ = v64
                                                                            }
                                                                        }
                                                                    default:
                                                                        var v65 int8 = (v31 * v31)
                                                                        _ = v65
                                                                    }
                                                                }
                                                                v66 := GenS4{F0: -1}
                                                                _ = v38
                                                                _ = v66
                                                            } else {
                                                                v24 -= v24
                                                                v67 := GenS5{F0: (v19 - v19)}
                                                                if (1e10 >= ((1.5 * -2.25) + -2.25)) {
                                                                    v68 := GenS3{F0: -2147483648, F1: 0, F2: v28}
                                                                    var v69 int = fi29
                                                                    _ = v68
                                                                    _ = v69
                                                                } else {
                                                                    var v70 string = ("\n\t" + "a")
                                                                    v28 -= v28
                                                                    switch v70 {
                                                                    case "hello":
                                                                        if (-453 != -1) {
                                                                            v70 = ((v70 + v70) + v70)
                                                                            var v71 uint16 = (v28 | v28)
                                                                            _ = v71
                                                                        }
                                                                    case "":
                                                                        v24 += v24
                                                                    case "foo bar":
                                                                        if false {
                                                                            var v72 float32 = ((((0.1 + (3.14159 + 0.0)) + (3.14159 + -2.25)) + (float32(3.14159) - 1.5)) * ((3.14159 * (-2.25 + 0.1)) / -2.25))
                                                                            var v73 uint = 18446744073709551615
                                                                            _ = v72
                                                                            _ = v73
                                                                        } else {
                                                                            v28 = uint16(v24)
                                                                        }
                                                                    }
                                                                    _ = v70
                                                                }
                                                                if (468 >= 1) {
                                                                    fmt.Println(v2, v31, fi27)
                                                                    v28 -= v28
                                                                    fmt.Println(fi29)
                                                                    var v74 bool = (255 >= (v28 ^ v28))
                                                                    _ = v74
                                                                } else {
                                                                    if (255 >= 1) {
                                                                        v19 = (v19 >> ((0) & 63))
                                                                    } else {
                                                                        v75 := GenS4{F0: -1}
                                                                        _ = v75
                                                                    }
                                                                }
                                                                v76 := GenS6{F0: int32(v28), F1: 1, F2: "\n\t"}
                                                                fmt.Println(v19, v31, fi27)
                                                                _ = v67
                                                                _ = v76
                                                            }
                                                        }
                                                        fmt.Println(v31, v28, v17)
                                                        v77 := GenS5{F0: 9223372036854775807}
                                                        switch (fi29 | -462) {
                                                        case 310:
                                                            if true {
                                                                _ = len(v3)
                                                                v19 = -556
                                                                v24 += (v24 << ((genFn1()) & 63))
                                                            } else {
                                                                v78 := GenS1{F0: int64(v24)}
                                                                _ = v78
                                                            }
                                                            v17 += v17
                                                            if ((v19 + (v19 | (v19 >> ((genFn1()) & 63)))) >= (v19 >> ((4294967295) & 63))) {
                                                                v79 := GenS7{F0: 4294967295, F1: v2}
                                                                var v80 int8 = v31
                                                                _ = v79
                                                                _ = v80
                                                            } else {
                                                                v17 = v17
                                                                if (v17 > (v17 | 1)) {
                                                                    if ((v28 * (v28 & 2)) >= (^(v28))) {
                                                                        v81 := v3[4294967295]
                                                                        var v82 uint8 = 42
                                                                        _ = v81
                                                                        _ = v82
                                                                    } else {
                                                                        fmt.Println(v24, v31)
                                                                    }
                                                                    var v83 string = "hello"
                                                                    if ((fi29 < 1) && true) {
                                                                        v84 := GenS2{F0: (v17 & genFn0())}
                                                                        _ = v84
                                                                    } else {
                                                                        if false {
                                                                            if len(v5) > 0 {
                                                                                v5[0] = 128
                                                                            }
                                                                            if (2147483647 != 0) {
                                                                                var v85 bool = true
                                                                                _ = v85
                                                                            }
                                                                            var v86 uint64 = 1
                                                                            _ = v86
                                                                        }
                                                                        switch fi27 {
                                                                        case 0:
                                                                            switch (v83 + "foo bar") {
                                                                            case "a":
                                                                                v87 := GenS3{F0: 127, F1: 2147483647, F2: uint16(v17)}
                                                                                _ = v87
                                                                            case "\n\t":
                                                                                if (((v31 + v31) >= (v31 % (((v31 & v31)) | 1))) || (false || false)) {
                                                                                    delete(v1, 2147483648)
                                                                                    v2 += v2
                                                                                }
                                                                                fmt.Println(v31)
                                                                                if len(v5) >= 2 {
                                                                                    v5 = v5[1:len(v5)]
                                                                                }
                                                                            case "hello":
                                                                                var v88 uint32 = 420
                                                                                _ = v88
                                                                            default:
                                                                                fmt.Println(v28, v2, v28)
                                                                                var v89 uint = 2
                                                                                _ = v89
                                                                            }
                                                                        default:
                                                                            var v90 int16 = (v2 - v2)
                                                                            _ = v90
                                                                        }
                                                                        switch ("a" + v83) {
                                                                        case "foo bar":
                                                                            if false {
                                                                                if ((v19 >> ((genFn1()) & 63)) != (v19 & v19)) {
                                                                                    var v91 float32 = -2.25
                                                                                    v19 -= -1
                                                                                    _ = len(v5)
                                                                                    _ = v91
                                                                                } else {
                                                                                    v17 += genFn0()
                                                                                    var v92 uint8 = 121
                                                                                    _ = v92
                                                                                }
                                                                                v24 += v24
                                                                            } else {
                                                                                v2 = 42
                                                                                if (435 > 0) {
                                                                                    if false {
                                                                                        v24 = (v24 >> ((294) & 63))
                                                                                        switch v83 {
                                                                                        case "\n\t":
                                                                                            var v93 uint16 = (v28 | (v28 * 983))
                                                                                            _ = v93
                                                                                        case "foo bar":
                                                                                            var v94 uint32 = 2147483648
                                                                                            var v95 uint32 = (v94 + v94)
                                                                                            v17 = v17
                                                                                            _ = v94
                                                                                            _ = v95
                                                                                        case "a":
                                                                                            v96 := GenS6{F0: 42, F1: (v24 << ((4294967295) & 63)), F2: "foo bar"}
                                                                                            _ = v96
                                                                                        default:
                                                                                            var v97 int16 = -1
                                                                                            _ = v97
                                                                                        }
                                                                                        var v98 uint8 = (v17 >> ((42) & 63))
                                                                                        var v99 int = (^(fi29))
                                                                                        _ = v98
                                                                                        _ = v99
                                                                                    }
                                                                                    switch (v83 + (v83 + ("a" + (v83 + "\n\t")))) {
                                                                                    case "\n\t":
                                                                                        if (!false) {
                                                                                            v100 := GenS8{F0: (^(v28)), F1: -177}
                                                                                            _ = v100
                                                                                        } else {
                                                                                            v28 += v28
                                                                                            var v101 float32 = 1e10
                                                                                            _ = v101
                                                                                        }
                                                                                    case "foo bar":
                                                                                        v2 = (v2 >> ((0) & 63))
                                                                                        var v102 int32 = 1
                                                                                        if ((-(v2)) > v2) {
                                                                                            var v103 int16 = -354
                                                                                            _ = v103
                                                                                        } else {
                                                                                            var v104 int64 = -9223372036854775808
                                                                                            _ = v104
                                                                                        }
                                                                                        _ = v102
                                                                                    case "a":
                                                                                        fmt.Println(v28, v17, v17)
                                                                                    }
                                                                                }
                                                                                switch (v83 + ((v83 + ("" + v83)) + (("a" + v83) + "\n\t"))) {
                                                                                case "foo bar":
                                                                                    v24 += 18446744073709551615
                                                                                    v31 = v31
                                                                                    var v105 uint = 0
                                                                                    _ = v105
                                                                                case "\n\t":
                                                                                    var v106 int64 = v19
                                                                                    _ = v106
                                                                                case "a":
                                                                                    v2 -= v2
                                                                                }
                                                                            }
                                                                        case "":
                                                                            v107 := GenS3{F0: (^(fi29)), F1: 42, F2: 1}
                                                                            _ = v107
                                                                        default:
                                                                            if ((!true) && true) {
                                                                                v17 = 127
                                                                                var v108 int64 = (v19 >> ((65535) & 63))
                                                                                _ = v108
                                                                            }
                                                                        }
                                                                        if (v19 == (v19 & (-(v19)))) {
                                                                            if (!((!(v28 < 255)) || ((^(v28)) == (v28 | 867)))) {
                                                                                if ((v2 ^ (v2 - int16(v31))) < v2) {
                                                                                    var v109 string = ((v83 + v83) + ((("a" + v83) + "hello") + v83))
                                                                                    _ = v109
                                                                                } else {
                                                                                    v5 = append(v5, fi27)
                                                                                }
                                                                                if (65535 <= 65535) {
                                                                                    if true {
                                                                                        fmt.Println(v19)
                                                                                        if (((1.5 + 1e10) * -2.25) <= 0.1) {
                                                                                            if true {
                                                                                                var v110 uint16 = (v28 | (v28 % (((v28 * (^(v28)))) | 1)))
                                                                                                _ = v110
                                                                                            } else {
                                                                                                if true {
                                                                                                    if ((true || true) && false) {
                                                                                                        var v111 uint = 2
                                                                                                        _ = v111
                                                                                                    } else {
                                                                                                        _ = len(v5)
                                                                                                        switch -234 {
                                                                                                        case -1:
                                                                                                            v19 = (v19 & (v19 ^ v19))
                                                                                                            if false {
                                                                                                                switch fi29 {
                                                                                                                case 426:
                                                                                                                    var v112 bool = true
                                                                                                                    _ = v112
                                                                                                                default:
                                                                                                                    if ((false || true) && false) {
                                                                                                                        fmt.Println(v31)
                                                                                                                    }
                                                                                                                }
                                                                                                                fmt.Println(v17, v17)
                                                                                                                if len(v5) > 0 {
                                                                                                                    v5[0] = fi27
                                                                                                                }
                                                                                                            }
                                                                                                            var v113 float32 = 1.5
                                                                                                            var v114 int32 = 143
                                                                                                            _ = len(v1)
                                                                                                            v28 += 73
                                                                                                            _ = v113
                                                                                                            _ = v114
                                                                                                        case -2147483648:
                                                                                                            var v115 int64 = v19
                                                                                                            _ = v115
                                                                                                        case 42:
                                                                                                            switch (fi29 | (fi29 / ((fi29) | 1))) {
                                                                                                            case 2147483647:
                                                                                                                fmt.Println(v17)
                                                                                                            }
                                                                                                        }
                                                                                                        switch int(v2) {
                                                                                                        case 79:
                                                                                                            var v116 uint32 = 0
                                                                                                            switch fi27 {
                                                                                                            case 1:
                                                                                                                v117 := GenS3{F0: 32767, F1: v116, F2: (v28 & (v28 >> ((804) & 63)))}
                                                                                                                _ = v117
                                                                                                            case -9223372036854775808:
                                                                                                                v2 = (v2 * int16(v24))
                                                                                                            case 32767:
                                                                                                                var v118 bool = (true && ((((fi27 >> ((2) & 63)) != fi29) || false) && false))
                                                                                                                _ = v118
                                                                                                            default:
                                                                                                                if ((557 != 1) || (0 != v17)) {
                                                                                                                    v1[v116] = (fi27 + fi27)
                                                                                                                }
                                                                                                            }
                                                                                                            v119 := GenS5{F0: 646}
                                                                                                            switch (v83 + "foo bar") {
                                                                                                            case "\n\t":
                                                                                                                v83 = (("foo bar" + (v83 + v83)) + "a")
                                                                                                            case "foo bar":
                                                                                                                if false {
                                                                                                                    fmt.Println(v17)
                                                                                                                    if false {
                                                                                                                        var v120 float64 = ((3.14159 / 3.14159) - (1e10 / ((-2.25 + (0.1 + 3.14159)) - 1.5)))
                                                                                                                        v120 -= 0.1
                                                                                                                        v121 := GenS3{F0: fi29, F1: v116, F2: (v28 << ((210) & 63))}
                                                                                                                        delete(v3, fi29)
                                                                                                                        switch (fi27 ^ (fi27 << ((126) & 63))) {
                                                                                                                        case 32767:
                                                                                                                            v2 = -35
                                                                                                                        }
                                                                                                                        _ = v120
                                                                                                                        _ = v121
                                                                                                                    } else {
                                                                                                                        var v122 int32 = 1
                                                                                                                        var v123 int8 = (-(v31))
                                                                                                                        var v124 int64 = 0
                                                                                                                        _ = v122
                                                                                                                        _ = v123
                                                                                                                        _ = v124
                                                                                                                    }
                                                                                                                    var v125 uint = uint(v116)
                                                                                                                    _ = v125
                                                                                                                }
                                                                                                            case "":
                                                                                                                var v126 uint = genFn1()
                                                                                                                _ = v126
                                                                                                            default:
                                                                                                                v127 := GenS2{F0: genFn0()}
                                                                                                                _ = v127
                                                                                                            }
                                                                                                            v17 += (^(v17))
                                                                                                            v128 := GenS9{F0: 42, F1: -2147483648}
                                                                                                            _ = v116
                                                                                                            _ = v119
                                                                                                            _ = v128
                                                                                                        case 4294967295:
                                                                                                            fmt.Println(fi29)
                                                                                                        default:
                                                                                                            v83 = (v83 + v83)
                                                                                                        }
                                                                                                    }
                                                                                                    if false {
                                                                                                        v2 = (v2 + -32768)
                                                                                                        v17 += (v17 - 46)
                                                                                                    }
                                                                                                    if true {
                                                                                                        v129 := GenS1{F0: (v19 % (((v19 % ((v19) | 1))) | 1))}
                                                                                                        if true {
                                                                                                            fmt.Println(v28, fi27)
                                                                                                            var v130 string = (v83 + "")
                                                                                                            var v131 int8 = v31
                                                                                                            v28 -= v28
                                                                                                            switch (fi29 >> ((1) & 63)) {
                                                                                                            case 42:
                                                                                                                switch (fi27 * fi27) {
                                                                                                                case 32768:
                                                                                                                    fmt.Println(v19, v17)
                                                                                                                case -9223372036854775808:
                                                                                                                    if ((v28 & v28) < (v28 * 0)) {
                                                                                                                        if (((v19 - (^(v19))) != (v19 & (v19 / (((v19 / ((-537) | 1))) | 1)))) && (true && true)) {
                                                                                                                            v19 = v19
                                                                                                                        }
                                                                                                                    }
                                                                                                                }
                                                                                                            case 32767:
                                                                                                                var v132 string = v83
                                                                                                                _ = v132
                                                                                                            case -522:
                                                                                                                v1[42] = (fi29 - -115)
                                                                                                            default:
                                                                                                                if true {
                                                                                                                    if (false || true) {
                                                                                                                        var v133 int64 = v19
                                                                                                                        switch v130 {
                                                                                                                        case "\n\t":
                                                                                                                            delete(v1, 54)
                                                                                                                        default:
                                                                                                                            var v134 int = -2147483648
                                                                                                                            _ = v134
                                                                                                                        }
                                                                                                                        _ = v133
                                                                                                                    } else {
                                                                                                                        switch -649 {
                                                                                                                        case 745:
                                                                                                                            if len(v5) >= 2 {
                                                                                                                                v5 = v5[1:len(v5)]
                                                                                                                            }
                                                                                                                        case 70:
                                                                                                                            var v135 uint32 = 2147483647
                                                                                                                            _ = v135
                                                                                                                        default:
                                                                                                                            v28 -= v28
                                                                                                                        }
                                                                                                                        v136 := GenS1{F0: (v19 >> ((1) & 63))}
                                                                                                                        _ = v136
                                                                                                                    }
                                                                                                                    if (!(v130 != v130)) {
                                                                                                                        switch (fi27 % ((32768) | 1)) {
                                                                                                                        case 127:
                                                                                                                            if len(v5) >= 2 {
                                                                                                                                v5 = v5[1:len(v5)]
                                                                                                                            }
                                                                                                                        }
                                                                                                                        var v137 float64 = 1.5
                                                                                                                        fmt.Println(v19, v24, v137)
                                                                                                                        _ = v137
                                                                                                                    } else {
                                                                                                                        v2 = (-(v2))
                                                                                                                    }
                                                                                                                }
                                                                                                            }
                                                                                                            _ = v130
                                                                                                            _ = v131
                                                                                                        } else {
                                                                                                            v138 := GenS10{F0: (v19 * v19), F1: ((v83 + v83) + (v83 + v83)), F2: v19}
                                                                                                            v31 += v31
                                                                                                            v24 = 18446744073709551615
                                                                                                            _ = v138
                                                                                                        }
                                                                                                        v139 := GenS11{F0: 524, F1: (1 != 0)}
                                                                                                        _ = v129
                                                                                                        _ = v139
                                                                                                    } else {
                                                                                                        fmt.Println(v31, fi29)
                                                                                                        switch (v83 + v83) {
                                                                                                        case "foo bar":
                                                                                                            v140 := GenS3{F0: (fi27 | fi29), F1: 0, F2: 65535}
                                                                                                            _ = v140
                                                                                                        case "\n\t":
                                                                                                            v31 -= v31
                                                                                                        case "a":
                                                                                                            _ = len(v1)
                                                                                                        }
                                                                                                        fmt.Println(fi27, v17, v2)
                                                                                                    }
                                                                                                } else {
                                                                                                    if (!true) {
                                                                                                        var v141 uint64 = (v24 % (((v24 * 821)) | 1))
                                                                                                        v142 := GenS9{F0: 4294967295, F1: -752}
                                                                                                        _ = v141
                                                                                                        _ = v142
                                                                                                    } else {
                                                                                                        var v143 float32 = 1e10
                                                                                                        v144 := GenS6{F0: -838, F1: 18446744073709551615, F2: v83}
                                                                                                        var v145 int8 = 1
                                                                                                        _ = v143
                                                                                                        _ = v144
                                                                                                        _ = v145
                                                                                                    }
                                                                                                    v19 -= v19
                                                                                                }
                                                                                                v146 := GenS9{F0: 538, F1: 1}
                                                                                                switch fi27 {
                                                                                                case -9223372036854775808:
                                                                                                    switch (fi29 / (((^(fi27))) | 1)) {
                                                                                                    case -1:
                                                                                                        v147 := GenS1{F0: v19}
                                                                                                        _ = v147
                                                                                                    case 32768:
                                                                                                        if ((v28 - (v28 << ((genFn1()) & 63))) > (v28 - 486)) {
                                                                                                            var v148 uint16 = v28
                                                                                                            _ = v148
                                                                                                        } else {
                                                                                                            _ = len(v1)
                                                                                                        }
                                                                                                    default:
                                                                                                        v24 -= v24
                                                                                                    }
                                                                                                case 9223372036854775807:
                                                                                                    var v149 int32 = 1
                                                                                                    _ = v149
                                                                                                default:
                                                                                                    fmt.Println(v2, v31, v28)
                                                                                                }
                                                                                                switch fi27 {
                                                                                                case 127:
                                                                                                    v31 = v31
                                                                                                case 128:
                                                                                                    if true {
                                                                                                        fmt.Println(v17)
                                                                                                    }
                                                                                                case 0:
                                                                                                    fmt.Println(v83)
                                                                                                }
                                                                                                _ = v146
                                                                                            }
                                                                                            v150 := GenS12{F0: 517, F1: (v19 ^ int64(v28)), F2: 911}
                                                                                            v151 := GenS12{F0: 65535, F1: (v19 + v19), F2: v19}
                                                                                            v24 += 4294967295
                                                                                            v2 += (v2 * v2)
                                                                                            switch v83 {
                                                                                            case "":
                                                                                                delete(v3, fi29)
                                                                                            default:
                                                                                                if (v83 > "hello") {
                                                                                                    switch fi27 {
                                                                                                    case -754:
                                                                                                        var v152 float32 = float32(((float64(1e10) + (((float64(3.14159) - (1.5 - ((-2.25 * -2.25) - (0.1 * 1.5)))) + 1.5) + (1.5 * ((-2.25 * 1e10) * (-2.25 / 1.0))))) * ((0.1 * 0.1) / 1e10)))
                                                                                                        _ = v152
                                                                                                    default:
                                                                                                        var v153 float32 = float32(3.14159)
                                                                                                        _ = v153
                                                                                                    }
                                                                                                    var v154 int8 = (v31 | (v31 & v31))
                                                                                                    _ = v154
                                                                                                }
                                                                                            }
                                                                                            _ = v150
                                                                                            _ = v151
                                                                                        }
                                                                                        v155 := GenS13{F0: false, F1: (v28 | v28)}
                                                                                        _ = v155
                                                                                    } else {
                                                                                        fmt.Println(fi29, v83, v31)
                                                                                    }
                                                                                } else {
                                                                                    v156 := GenS8{F0: uint16(v24), F1: (v2 << ((1) & 63))}
                                                                                    _ = v156
                                                                                }
                                                                                v157 := GenS7{F0: 124, F1: (^(v2))}
                                                                                v83 = v83
                                                                                v158 := GenS14{F0: (v17 ^ v17), F1: 2, F2: (^(v24))}
                                                                                if true {
                                                                                    var v159 string = ("hello" + v83)
                                                                                    var v160 bool = ((v159 + "a") == v159)
                                                                                    var v161 uint16 = v28
                                                                                    _ = v159
                                                                                    _ = v160
                                                                                    _ = v161
                                                                                }
                                                                                switch (-(fi29)) {
                                                                                case -497:
                                                                                    v31 += 100
                                                                                }
                                                                                if false {
                                                                                    var v162 float64 = 0.1
                                                                                    _ = v162
                                                                                } else {
                                                                                    var v163 uint32 = 299
                                                                                    var v164 int = fi29
                                                                                    if (v31 < v31) {
                                                                                        v5 = append(v5, (-(v164)))
                                                                                        switch (v164 / ((v164) | 1)) {
                                                                                        case -211:
                                                                                            switch v83 {
                                                                                            case "":
                                                                                                var v165 uint = 42
                                                                                                _ = v165
                                                                                            }
                                                                                        case -9223372036854775808:
                                                                                            v166 := GenS14{F0: (v17 * (v17 & (v17 / (((^(v17))) | 1)))), F1: 2147483648, F2: v24}
                                                                                            _ = v166
                                                                                        }
                                                                                        var v167 int16 = 974
                                                                                        _ = v167
                                                                                    }
                                                                                    _ = v163
                                                                                    _ = v164
                                                                                }
                                                                                _ = v157
                                                                                _ = v158
                                                                            }
                                                                            var v168 float64 = 0.0
                                                                            _ = v168
                                                                        } else {
                                                                            switch fi27 {
                                                                            case 32768:
                                                                                var v169 int32 = 42
                                                                                _ = v169
                                                                            default:
                                                                                v17 -= (^(v17))
                                                                            }
                                                                            fmt.Println(v2)
                                                                            v17 = genFn0()
                                                                        }
                                                                    }
                                                                    v24 = 0
                                                                    _ = v83
                                                                }
                                                            }
                                                        default:
                                                            if (389 <= -2147483648) {
                                                                var v170 uint32 = 496
                                                                var v171 int16 = (v2 << ((2) & 63))
                                                                _ = v170
                                                                _ = v171
                                                            } else {
                                                                v172 := GenS15{F0: 1}
                                                                v28 -= v28
                                                                _ = v172
                                                            }
                                                        }
                                                        _ = v77
                                                    } else {
                                                        var v173 float32 = (1e10 + ((0.1 - (0.0 / 1.0)) / -2.25))
                                                        fmt.Println(v19)
                                                        var v174 uint = 207
                                                        _ = v173
                                                        _ = v174
                                                    }
                                                }
                                            }
                                            if true {
                                                switch fi27 {
                                                case -787:
                                                    v2 -= (^(v2))
                                                case -899:
                                                    if true {
                                                        if false {
                                                            v175 := GenS16{F0: (v17 << ((genFn1()) & 63))}
                                                            var v176 int16 = v2
                                                            _ = v175
                                                            _ = v176
                                                        } else {
                                                            v177 := v3[fi27]
                                                            v1[992] = (v177 + 642)
                                                            fmt.Println(v2)
                                                            _ = v177
                                                        }
                                                    }
                                                default:
                                                    var v178 float64 = float64((0.1 * 0.0))
                                                    _ = v178
                                                }
                                                v179 := GenS10{F0: int64(v28), F1: "", F2: (v19 >> ((0) & 63))}
                                                v24 = v24
                                                _ = v179
                                            }
                                        default:
                                            switch ("\n\t" + "\n\t") {
                                            case "foo bar":
                                                var v180 string = "\n\t"
                                                _ = v180
                                            case "\n\t":
                                                v2 = int16(v28)
                                            }
                                        }
                                    } else {
                                        v2 = v2
                                    }
                                    v181 := v4[1]
                                    var v182 uint32 = 4294967295
                                    _ = v181
                                    _ = v182
                                } else {
                                    v183 := GenS17{F0: 0, F1: (v17 + genFn0())}
                                    v2 -= (^(v2))
                                    _ = v183
                                }
                            case 490:
                                var v184 int8 = v31
                                _ = v184
                            }
                            _ = v31
                        }
                    }
                    fmt.Println(fi29)
                }
                var v185 uint64 = v24
                _ = fi29
                _ = v185
            }
            if ((v24 > (v24 ^ v24)) && false) {
                v17 = (v17 - (^(v17)))
            }
            _ = fi27
            _ = v28
        }
        for fi186 := 0; fi186 < 7; fi186++ {
            fmt.Println(fi186, fi186, fi186)
            for fi187 := 0; fi187 < 13; fi187++ {
                v19 -= (v19 << ((95) & 63))
                var v188 int64 = v19
                var v189 string = "hello"
                v190 := GenS10{F0: (-(v19)), F1: v189, F2: v188}
                v191 := GenS14{F0: 128, F1: 668, F2: (v24 ^ v24)}
                if (!true) {
                    var v192 int = fi186
                    _ = v192
                }
                _ = fi187
                _ = v188
                _ = v189
                _ = v190
                _ = v191
            }
            if true {
                var v193 uint8 = genFn0()
                for fi194 := 0; fi194 < 11; fi194++ {
                    var v195 int16 = -338
                    var v196 uint = 1
                    _ = fi194
                    _ = v195
                    _ = v196
                }
                for fi197 := 0; fi197 < 2; fi197++ {
                    var v198 bool = false
                    v19 = v19
                    _ = fi197
                    _ = v198
                }
                _ = v193
            } else {
                var v199 float32 = (((0.1 + 1.5) * 0.1) - 0.1)
                _ = v199
            }
            _ = fi186
        }
    }
    for fi200 := 0; fi200 < 13; fi200++ {
        v201 := GenS18{F0: 4294967295, F1: 2, F2: -1}
        fmt.Println(fi200, fi200)
        _ = fi200
        _ = v201
    }
    v202 := GenS19{F0: (v19 | v19), F1: int32(v2)}
    var v203 uint = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v203:", v203)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v202
    _ = v203
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
