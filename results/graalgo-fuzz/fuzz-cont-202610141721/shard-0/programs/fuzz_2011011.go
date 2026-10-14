package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32; F2 int }
type GenS2 struct { F0 string; F1 uint8 }
type GenS3 struct { F0 int32; F1 uint16; F2 uint16 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint16; F1 bool }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 int16; F1 uint8 }
type GenS8 struct { F0 int; F1 uint32; F2 uint }
type GenS9 struct { F0 uint64 }
type GenS10 struct { F0 uint64; F1 int16; F2 uint64 }
type GenS11 struct { F0 int64; F1 int16; F2 uint64 }
type GenS12 struct { F0 uint32; F1 int16 }
type GenS13 struct { F0 int64; F1 int32 }
type GenS14 struct { F0 uint64 }
type GenS15 struct { F0 uint; F1 uint }
type GenS16 struct { F0 bool }
type GenS17 struct { F0 bool }
func genFn0(p0 uint8) int32 {
    v7 := []int{1, 128}
    switch -9223372036854775808 {
    case 2147483647:
        var v8 int32 = 0
        _ = v8
    default:
        p0 = p0
    }
    _ = p0
    _ = v7
    return 0
}
func genFn1(p0 uint) int16 {
    if true {
        p0 += 4
        switch 42 {
        case 641:
            if ((((3.14159 * 3.14159) + (1e10 + ((((0.1 - 0.1) / 1.5) - float32(3.14159)) / 1.0))) == 3.14159) || false) {
                for fi9 := 0; fi9 < 17; fi9++ {
                    v10 := GenS1{F0: true, F1: 913, F2: (fi9 & fi9)}
                    switch "foo bar" {
                    case "hello":
                        if false {
                            if false {
                                fmt.Println(fi9, fi9)
                                fmt.Println(fi9, fi9)
                                switch ("hello" + "foo bar") {
                                case "hello":
                                    var v11 uint = p0
                                    _ = v11
                                case "foo bar":
                                    v12 := []int{fi9, fi9}
                                    _ = v12
                                default:
                                    v13 := []int{}
                                    _ = v13
                                }
                            } else {
                                var v14 uint8 = 0
                                if ((false || false) && true) {
                                    var v15 int8 = 0
                                    v16 := []int{(fi9 / ((-2147483648) | 1)), fi9, (fi9 | (fi9 & int(p0))), 841}
                                    var v17 int64 = 565
                                    _ = v15
                                    _ = v16
                                    _ = v17
                                }
                                _ = v14
                            }
                        } else {
                            var v18 uint64 = 896
                            v18 -= 0
                            _ = v18
                        }
                    case "foo bar":
                        p0 -= p0
                    case "\n\t":
                        v19 := GenS1{F0: false, F1: 72, F2: fi9}
                        _ = v19
                    default:
                        p0 -= p0
                    }
                    _ = fi9
                    _ = v10
                }
                if true {
                    fmt.Println(p0)
                }
            }
        default:
            fmt.Println(p0)
        }
    } else {
        if false {
            fmt.Println(p0)
            v20 := GenS1{F0: false, F1: 873, F2: 9223372036854775807}
            _ = v20
        }
        if (!false) {
            switch -2147483648 {
            case 128:
                v21 := GenS1{F0: false, F1: 333, F2: 9223372036854775807}
                _ = v21
            }
            var v22 int16 = 718
            _ = v22
        } else {
            p0 = p0
            fmt.Println(p0)
        }
    }
    v23 := GenS1{F0: false, F1: 848, F2: 128}
    _ = p0
    _ = v23
    return 464
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint = 300
    var v3 float64 = (-2.25 + (-2.25 + 1.5))
    var v4 uint32 = uint32(v2)
    var v5 int8 = 54
    var v6 int8 = v5
    var v24 uint8 = 2
    var v25 uint16 = 2
    var v26 uint64 = uint64(v4)
    v2 -= uint(v24)
    var v27 int64 = -548
    switch ("" + "hello") {
    case "\n\t":
        for fi28 := 0; fi28 < 8; fi28++ {
            v2 = 1
            fmt.Println(v24, v5)
            switch fi28 {
            case 42:
                for fi29 := 0; fi29 < 15; fi29++ {
                    if false {
                        v25 -= 65535
                    }
                    _ = fi29
                }
            case 0:
                switch fi28 {
                case 127:
                    fmt.Println(v25)
                case -437:
                    v4 = v4
                case 287:
                    var v30 float32 = 3.14159
                    _ = v30
                }
            default:
                switch 2147483647 {
                case 1:
                    for fi31 := 0; fi31 < 2; fi31++ {
                        fmt.Println(v26, v27, v6)
                        if false {
                            v32 := []int{int(v25)}
                            _ = v32
                        } else {
                            v33 := GenS2{F0: "", F1: (v24 * 127)}
                            v1[2] = v4
                            _ = v33
                        }
                        _ = fi31
                    }
                case 32768:
                    if true {
                        var v34 uint = (v2 ^ (^(v2)))
                        _ = len(v1)
                        _ = v34
                    } else {
                        v6 += (v5 >> ((v2) & 63))
                        delete(v1, v4)
                        fmt.Println(v3)
                    }
                case 4294967295:
                    var v35 int32 = -1
                    _ = v35
                default:
                    v36 := GenS3{F0: genFn0(117), F1: (v25 << (((^(v2))) & 63)), F2: (^(v25))}
                    _ = v36
                }
            }
            _ = fi28
        }
    case "hello":
        v37 := []int{32767, 32768, 42, 9223372036854775807}
        _ = v37
    case "a":
        v4 = v4
    }
    var v38 int8 = 127
    if false {
        var v39 uint64 = 532
        for fi40 := 0; fi40 < 18; fi40++ {
            for fi41 := 0; fi41 < 3; fi41++ {
                if true {
                    var v42 bool = true
                    fmt.Println(v25, v39)
                    v27 -= v27
                    _ = v42
                } else {
                    v43 := GenS1{F0: (((2 == v24) && (1 < (v39 + v39))) && false), F1: v4, F2: (fi40 << (((v2 | (v2 ^ 18446744073709551615))) & 63))}
                    fmt.Println(v38, fi41, v3)
                    var v44 string = "\n\t"
                    v45 := []int{(fi41 >> ((v2) & 63)), (fi40 / ((fi40) | 1)), (^(fi41))}
                    fmt.Println(v4, v25, v38)
                    v25 = uint16(v24)
                    v26 = 688
                    _ = v43
                    _ = v44
                    _ = v45
                }
                if (uint8(v2) != v24) {
                    v46 := []int{(fi40 / ((9223372036854775807) | 1))}
                    fmt.Println(v38, fi40)
                    v27 = 42
                    _ = v46
                } else {
                    v47 := GenS4{F0: 42}
                    if false {
                        var v48 float32 = (0.0 * 1.5)
                        fmt.Println(v27, v2)
                        var v49 bool = (1 > 128)
                        fmt.Println(v3, v39, v49)
                        _ = v48
                        _ = v49
                    }
                    _ = v47
                }
                _ = fi41
            }
            var v50 uint16 = (v25 | v25)
            _ = fi40
            _ = v50
        }
        v51 := GenS2{F0: "", F1: v24}
        _ = v39
        _ = v51
    }
    delete(v1, (v4 - 2147483647))
    if ("a" != "hello") {
        v52 := []int{4294967295, 42}
        fmt.Println(v24)
        var v53 uint64 = (v26 / ((v26) | 1))
        _ = v52
        _ = v53
    } else {
        if (int(v6) >= 9223372036854775807) {
            var v54 string = "foo bar"
            var v55 bool = false
            for fi56 := 0; fi56 < 10; fi56++ {
                switch "hello" {
                case "a":
                    fmt.Println(v24, fi56, v2)
                default:
                    if v55 {
                        fmt.Println(v27)
                        if v55 {
                            var v57 uint32 = v4
                            switch fi56 {
                            case 128:
                                if v55 {
                                    v5 = (v5 & (v6 / (((v6 % ((42) | 1))) | 1)))
                                    v2 = v2
                                    switch (fi56 >> (((v2 / (((v2 & 901)) | 1))) & 63)) {
                                    case 0:
                                        fmt.Println(v57, v55, v5)
                                    case 127:
                                        v58 := GenS3{F0: 1, F1: v25, F2: 32767}
                                        _ = v58
                                    case 42:
                                        v59 := GenS5{F0: 65535, F1: v55}
                                        _ = v59
                                    }
                                } else {
                                    v27 += (v27 & v27)
                                }
                            default:
                                v60 := GenS6{F0: (v25 << ((v2) & 63))}
                                _ = v60
                            }
                            v61 := []int{(fi56 | int(v2)), (fi56 >> ((uint(v25)) & 63))}
                            _ = v57
                            _ = v61
                        } else {
                            for fi62 := 0; fi62 < 11; fi62++ {
                                switch "foo bar" {
                                case "foo bar":
                                    if v55 {
                                        var v63 bool = v55
                                        if (false && (v55 && ((v26 == v26) || (128 != uint8(v26))))) {
                                            v1[v4] = v4
                                            if v55 {
                                                var v64 float64 = (v3 * (v3 * v3))
                                                _ = v64
                                            } else {
                                                v65 := v1[v4]
                                                switch fi62 {
                                                case -183:
                                                    v66 := GenS6{F0: (v25 ^ 255)}
                                                    _ = v66
                                                case 42:
                                                    switch "hello" {
                                                    case "hello":
                                                        switch v54 {
                                                        case "\n\t":
                                                            var v67 uint8 = uint8(v4)
                                                            _ = v67
                                                        case "foo bar":
                                                            var v68 int8 = v5
                                                            _ = v68
                                                        }
                                                    case "":
                                                        if (!((!((!true) && v63)) || (3.14159 < (3.14159 * 3.14159)))) {
                                                            if (false || true) {
                                                                if ((v65 ^ (v65 / (((v65 / ((v4) | 1))) | 1))) >= (v65 / ((v4) | 1))) {
                                                                    if true {
                                                                        switch int(v24) {
                                                                        case 380:
                                                                            v63 = (float32(v3) <= (-2.25 + 1e10))
                                                                        }
                                                                        v54 = (v54 + ("foo bar" + "a"))
                                                                        switch fi56 {
                                                                        case 4294967295:
                                                                            v69 := GenS7{F0: -945, F1: 42}
                                                                            _ = v69
                                                                        case 32768:
                                                                            v65 += v4
                                                                        }
                                                                    } else {
                                                                        var v70 uint32 = v65
                                                                        var v71 int64 = v27
                                                                        _ = v70
                                                                        _ = v71
                                                                    }
                                                                    switch -2147483648 {
                                                                    case 128:
                                                                        if (fi62 != (^(fi62))) {
                                                                            v65 += 65535
                                                                        }
                                                                    case 42:
                                                                        v24 = v24
                                                                    case 809:
                                                                        v63 = (-2147483648 >= -1)
                                                                    }
                                                                } else {
                                                                    v24 += v24
                                                                    v4 -= (v65 - v4)
                                                                    var v72 float32 = 0.0
                                                                    _ = v72
                                                                }
                                                                v73 := GenS3{F0: 1, F1: (v25 % ((v25) | 1)), F2: v25}
                                                                _ = v73
                                                            }
                                                            v74 := GenS4{F0: -103}
                                                            _ = v74
                                                        }
                                                        v26 = (v26 ^ (v26 / ((uint64(v25)) | 1)))
                                                    case "\n\t":
                                                        fmt.Println(v4)
                                                    }
                                                }
                                                v3 += ((v3 - v3) - (v3 - 0.1))
                                                _ = v65
                                            }
                                            fmt.Println(v5)
                                            var v75 int32 = -2147483648
                                            v76 := GenS8{F0: fi62, F1: v4, F2: v2}
                                            if v63 {
                                                switch 32767 {
                                                case 0:
                                                    v77 := GenS3{F0: (v75 << (((v2 / ((v2) | 1))) & 63)), F1: (v25 - (v25 | (v25 >> ((uint(v24)) & 63)))), F2: uint16(v4)}
                                                    v55 = true
                                                    v75 = v75
                                                    var v78 string = "hello"
                                                    v3 -= v3
                                                    _ = v77
                                                    _ = v78
                                                case 32768:
                                                    v79 := GenS9{F0: (^(v26))}
                                                    _ = v79
                                                case -1:
                                                    var v80 float32 = float32(3.14159)
                                                    _ = v80
                                                }
                                                if (v6 < (v5 ^ v6)) {
                                                    fmt.Println(v4, fi56)
                                                } else {
                                                    switch (^(fi62)) {
                                                    case 127:
                                                        if ("\n\t" <= (v54 + ("a" + v54))) {
                                                            switch "hello" {
                                                            case "hello":
                                                                fmt.Println(v75, v55)
                                                            default:
                                                                var v81 int8 = -46
                                                                _ = v81
                                                            }
                                                            var v82 string = (v54 + v54)
                                                            v75 += int32(v24)
                                                            v24 = 0
                                                            _ = v82
                                                        }
                                                        v83 := []int{fi62, fi62, fi56}
                                                        _ = v83
                                                    case 32767:
                                                        var v84 int64 = (v27 ^ (v27 + v27))
                                                        _ = v84
                                                    case 0:
                                                        fmt.Println(v38, v4)
                                                    }
                                                    var v85 float64 = ((v3 - (((float64(1e10) + v3) * v3) - (((v3 - v3) + v3) - v3))) - (v3 / 1.5))
                                                    _ = v85
                                                }
                                                if v63 {
                                                    var v86 bool = ((v24 ^ (^(v24))) != (v24 / ((255) | 1)))
                                                    v6 = (v6 << (((v2 << (((v2 ^ v2)) & 63))) & 63))
                                                    v87 := GenS9{F0: (v26 & 987)}
                                                    _ = v86
                                                    _ = v87
                                                }
                                                var v88 int16 = 0
                                                switch fi56 {
                                                case -223:
                                                    v89 := []int{}
                                                    _ = v89
                                                case 0:
                                                    if (((1e10 / -2.25) / 3.14159) >= -2.25) {
                                                        if v55 {
                                                            var v90 float64 = 0.0
                                                            _ = v90
                                                        }
                                                        if (((v5 % ((v38) | 1)) >= v6) || (0 < (^(v27)))) {
                                                            fmt.Println(v88)
                                                            var v91 int8 = int8(v25)
                                                            switch fi56 {
                                                            case 2147483647:
                                                                if v63 {
                                                                    var v92 uint = v2
                                                                    if v63 {
                                                                        fmt.Println(v88, v5)
                                                                        v93 := v1[v4]
                                                                        _ = v93
                                                                    } else {
                                                                        v6 = -8
                                                                        switch 9223372036854775807 {
                                                                        case -14:
                                                                            var v94 float32 = (3.14159 * 1e10)
                                                                            if v55 {
                                                                                v55 = true
                                                                                switch (v54 + v54) {
                                                                                case "a":
                                                                                    var v95 int32 = v75
                                                                                    _ = v95
                                                                                case "\n\t":
                                                                                    var v96 bool = (v54 < v54)
                                                                                    _ = v96
                                                                                case "foo bar":
                                                                                    v97 := []int{(fi56 >> (((v92 << ((255) & 63))) & 63)), (fi56 * 127), (fi56 >> (((v2 + (v92 % (((^(v2))) | 1)))) & 63))}
                                                                                    _ = v97
                                                                                default:
                                                                                    fmt.Println(v88)
                                                                                }
                                                                            }
                                                                            _ = v94
                                                                        case 9223372036854775807:
                                                                            v92 += (v92 << ((1) & 63))
                                                                        }
                                                                    }
                                                                    _ = v92
                                                                } else {
                                                                    v3 -= 0.0
                                                                    if v55 {
                                                                        fmt.Println(fi56, v24, v54)
                                                                        v26 -= v26
                                                                    }
                                                                }
                                                            case -1:
                                                                v98 := GenS10{F0: v26, F1: (^(v88)), F2: v26}
                                                                _ = v98
                                                            default:
                                                                fmt.Println(v24, v75, v4)
                                                            }
                                                            _ = v91
                                                        } else {
                                                            v6 = (v6 % ((-128) | 1))
                                                        }
                                                        if v55 {
                                                            var v99 int16 = (v88 % (((v88 % ((v88) | 1))) | 1))
                                                            _ = len(v1)
                                                            switch (fi56 + fi56) {
                                                            case 2147483647:
                                                                v100 := GenS6{F0: (^(v25))}
                                                                _ = v100
                                                            case 32768:
                                                                v54 = v54
                                                            case 32767:
                                                                if (v25 == 498) {
                                                                    fmt.Println(v38)
                                                                    v101 := GenS11{F0: 9223372036854775807, F1: genFn1((v2 / ((2) | 1))), F2: 2}
                                                                    _ = v101
                                                                }
                                                            }
                                                            _ = v99
                                                        }
                                                    }
                                                default:
                                                    switch 0 {
                                                    case -2147483648:
                                                        var v102 float64 = ((v3 + v3) / 0.1)
                                                        _ = v102
                                                    case 128:
                                                        v103 := v1[2147483648]
                                                        _ = v103
                                                    }
                                                }
                                                _ = v88
                                            }
                                            if false {
                                                var v104 int = fi62
                                                if (((false && false) && ((3.14159 <= 3.14159) || false)) && v55) {
                                                    v105 := GenS12{F0: v4, F1: -1}
                                                    v75 += 0
                                                    var v106 float64 = 1e10
                                                    _ = v105
                                                    _ = v106
                                                }
                                                _ = v104
                                            }
                                            if ((v5 * (v38 >> (((v2 + 0)) & 63))) != v5) {
                                                switch v54 {
                                                case "hello":
                                                    var v107 string = "hello"
                                                    _ = v107
                                                case "":
                                                    v5 = (v38 & v6)
                                                }
                                                if v63 {
                                                    var v108 int = fi62
                                                    v108 -= fi56
                                                    switch "foo bar" {
                                                    case "hello":
                                                        switch v54 {
                                                        case "hello":
                                                            v109 := GenS10{F0: uint64(v25), F1: -566, F2: (^(v26))}
                                                            _ = v109
                                                        default:
                                                            if ((v4 >> ((v2) & 63)) >= (v4 << ((42) & 63))) {
                                                                switch v54 {
                                                                case "hello":
                                                                    v110 := GenS13{F0: (v27 - (v27 & v27)), F1: (^(v75))}
                                                                    v25 += uint16(v2)
                                                                    _ = v110
                                                                case "foo bar":
                                                                    v111 := GenS14{F0: (^(v26))}
                                                                    _ = v111
                                                                default:
                                                                    v112 := GenS12{F0: v4, F1: 911}
                                                                    _ = v112
                                                                }
                                                            } else {
                                                                v27 -= int64(v24)
                                                                var v113 float64 = v3
                                                                v25 = (v25 * 2)
                                                                _ = v113
                                                            }
                                                        }
                                                    }
                                                    _ = v108
                                                }
                                            } else {
                                                v114 := GenS6{F0: 550}
                                                v75 = -1
                                                _ = v114
                                            }
                                            _ = v75
                                            _ = v76
                                        }
                                        _ = v63
                                    }
                                case "":
                                    var v115 string = (v54 + (v54 + "a"))
                                    _ = v115
                                default:
                                    var v116 string = v54
                                    _ = v116
                                }
                                switch v54 {
                                case "a":
                                    if v55 {
                                        if true {
                                            var v117 int32 = int32(v27)
                                            if (v55 || (v27 == 692)) {
                                                v118 := []int{fi62, 42, (fi62 * (fi56 - 4294967295)), fi56}
                                                _ = v118
                                            }
                                            _ = v117
                                        }
                                    } else {
                                        if true {
                                            if (1.5 >= 0.1) {
                                                v119 := GenS15{F0: v2, F1: (v2 & (v2 | v2))}
                                                var v120 int64 = v27
                                                switch fi56 {
                                                case 699:
                                                    v121 := GenS15{F0: v2, F1: v2}
                                                    _ = v121
                                                default:
                                                    fmt.Println(v55, v6, v24)
                                                }
                                                fmt.Println(v6, v25)
                                                _ = v119
                                                _ = v120
                                            }
                                        }
                                        var v122 int8 = v38
                                        _ = v122
                                    }
                                case "foo bar":
                                    var v123 uint = (^(v2))
                                    _ = v123
                                default:
                                    fmt.Println(v55, v25)
                                }
                                if v55 {
                                    switch (("hello" + v54) + v54) {
                                    case "foo bar":
                                        var v124 bool = (v55 && v55)
                                        _ = v124
                                    }
                                    if v55 {
                                        v5 = int8(fi62)
                                        if (v4 != v4) {
                                            v24 = 13
                                        } else {
                                            var v125 uint64 = v26
                                            _ = v125
                                        }
                                    }
                                } else {
                                    var v126 uint8 = (^(v24))
                                    _ = v126
                                }
                                _ = fi62
                            }
                            v127 := v1[(v4 >> ((0) & 63))]
                            if (135 > 128) {
                                v128 := GenS7{F0: -1, F1: uint8(v127)}
                                if true {
                                    v1[(v4 << ((uint(v25)) & 63))] = 1
                                } else {
                                    v129 := GenS13{F0: (v27 << (((v2 / ((v2) | 1))) & 63)), F1: genFn0(uint8(v25))}
                                    switch ("\n\t" + v54) {
                                    case "a":
                                        if v55 {
                                            v55 = (!false)
                                        }
                                    case "hello":
                                        v26 += uint64(v2)
                                    default:
                                        if v55 {
                                            fmt.Println(v27, v5)
                                        } else {
                                            fmt.Println(fi56, v25)
                                        }
                                    }
                                    _ = v129
                                }
                                _ = v128
                            }
                            _ = v127
                        }
                    }
                }
                if (v55 && ((2147483647 <= v4) || v55)) {
                    if v55 {
                        if ((0 <= 1) || (v25 >= uint16(v2))) {
                            var v130 uint16 = v25
                            switch ((("hello" + v54) + v54) + ("hello" + v54)) {
                            case "a":
                                if v55 {
                                    var v131 uint8 = (^(v24))
                                    var v132 uint32 = v4
                                    for fi133 := 0; fi133 < 5; fi133++ {
                                        v4 = v4
                                        _ = fi133
                                    }
                                    _ = v131
                                    _ = v132
                                } else {
                                    v38 -= v38
                                }
                            }
                            v27 = 9223372036854775807
                            _ = v130
                        }
                        v24 = 2
                    } else {
                        var v134 float64 = (float64(1.5) - ((v3 / (v3 * ((v3 * v3) + (1.5 / -2.25)))) + v3))
                        v135 := []int{(fi56 % ((int(v27)) | 1))}
                        v136 := GenS11{F0: int64(v25), F1: 0, F2: (v26 + (v26 * (v26 ^ (^(v26)))))}
                        _ = v134
                        _ = v135
                        _ = v136
                    }
                    for fi137 := 0; fi137 < 16; fi137++ {
                        v138 := GenS2{F0: v54, F1: 157}
                        _ = fi137
                        _ = v138
                    }
                    var v139 uint16 = v25
                    _ = v139
                } else {
                    switch (fi56 << ((v2) & 63)) {
                    case 2147483647:
                        switch (("hello" + ("foo bar" + ("a" + v54))) + v54) {
                        case "":
                            switch "" {
                            case "hello":
                                fmt.Println(fi56, v5, v25)
                            case "a":
                                if (0.0 > 1.5) {
                                    for fi140 := 0; fi140 < 3; fi140++ {
                                        if true {
                                            var v141 int16 = -1
                                            var v142 uint = uint(v24)
                                            _ = v141
                                            _ = v142
                                        } else {
                                            var v143 uint64 = 2
                                            v144 := []int{-9223372036854775808, (fi140 + (^(fi56))), (fi56 * (fi140 * (-(fi140))))}
                                            _ = v143
                                            _ = v144
                                        }
                                        if true {
                                            v145 := GenS15{F0: (v2 & (v2 * v2)), F1: (v2 ^ v2)}
                                            _ = v145
                                        } else {
                                            var v146 uint = 159
                                            v147 := GenS15{F0: uint(v24), F1: v146}
                                            v148 := GenS9{F0: 28}
                                            _ = v146
                                            _ = v147
                                            _ = v148
                                        }
                                        _ = fi140
                                    }
                                } else {
                                    v4 = (v4 % (((v4 * v4)) | 1))
                                }
                            case "":
                                var v149 int32 = int32(v5)
                                _ = v149
                            }
                        }
                    case 127:
                        v26 += 1
                    case -9223372036854775808:
                        if false {
                            for fi150 := 0; fi150 < 2; fi150++ {
                                v24 -= uint8(v25)
                                switch v54 {
                                case "":
                                    fmt.Println(fi150, v5, v24)
                                }
                                _ = fi150
                            }
                            if v55 {
                                for fi151 := 0; fi151 < 14; fi151++ {
                                    v152 := GenS6{F0: 664}
                                    _ = fi151
                                    _ = v152
                                }
                                for fi153 := 0; fi153 < 14; fi153++ {
                                    v5 -= (v6 * (v5 << ((856) & 63)))
                                    v154 := GenS14{F0: (v26 - 2)}
                                    _ = fi153
                                    _ = v154
                                }
                                v6 -= v38
                            } else {
                                if ((v55 || (true || (!v55))) || v55) {
                                    v2 += uint(v4)
                                    v155 := v1[uint32(v2)]
                                    _ = v155
                                } else {
                                    v156 := []int{}
                                    v55 = v55
                                    switch v54 {
                                    case "hello":
                                        if false {
                                            v157 := GenS3{F0: -250, F1: (v25 * 1), F2: (^(v25))}
                                            v1[(^(v4))] = uint32(v25)
                                            _ = v157
                                        }
                                    case "\n\t":
                                        if false {
                                            var v158 uint16 = v25
                                            v159 := GenS16{F0: false}
                                            _ = v158
                                            _ = v159
                                        } else {
                                            var v160 int8 = -128
                                            _ = v160
                                        }
                                    }
                                    _ = v156
                                }
                                v161 := GenS17{F0: v55}
                                v27 = v27
                                _ = v161
                            }
                        }
                    default:
                        var v162 uint32 = v4
                        _ = v162
                    }
                    v1[v4] = (^(v4))
                }
                v163 := v1[v4]
                _ = fi56
                _ = v163
            }
            _ = v54
            _ = v55
        } else {
            if ((v27 / ((v27) | 1)) > v27) {
                var v164 uint = v2
                var v165 int16 = -684
                _ = v164
                _ = v165
            }
            var v166 int = 288
            _ = v166
        }
        var v167 uint = (v2 << ((65535) & 63))
        _ = v167
    }
    var v168 uint = (v2 | 42)
    for fi169 := 0; fi169 < 14; fi169++ {
        fmt.Println(v27, v2)
        fmt.Println(v38)
        _ = fi169
    }
    var v170 float64 = v3
    var v171 uint16 = 376
    v172 := v1[2147483647]
    v173 := []int{4294967295, 941, 386, 32767}
    if (3.14159 > (0.0 - 0.1)) {
        v27 += (v27 << (((^(v168))) & 63))
        fmt.Println(v168, v168)
        var v174 uint64 = (^(v26))
        _ = v174
    }
    var v175 uint = v168
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v168:", v168)
    fmt.Println("v170:", v170)
    fmt.Println("v171:", v171)
    fmt.Println("v172:", v172)
    fmt.Println("slicelen v173:", len(v173), v173)
    fmt.Println("v175:", v175)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v168
    _ = v170
    _ = v171
    _ = v172
    _ = v173
    _ = v175
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
