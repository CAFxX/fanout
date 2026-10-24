package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8 }
type GenS2 struct { F0 int32; F1 uint8; F2 uint32 }
type GenS3 struct { F0 bool; F1 uint }
type GenS4 struct { F0 uint; F1 uint8 }
type GenS5 struct { F0 bool; F1 uint64 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 bool; F1 uint16 }
type GenS8 struct { F0 string; F1 int64; F2 string }
type GenS9 struct { F0 uint; F1 uint }
type GenS10 struct { F0 uint8; F1 string; F2 int32 }
type GenS11 struct { F0 int64 }
type GenS12 struct { F0 int16 }
type GenS13 struct { F0 string; F1 uint8; F2 uint32 }
type GenS14 struct { F0 uint32 }
type GenS15 struct { F0 uint32; F1 uint64; F2 uint8 }
type GenS16 struct { F0 uint16; F1 uint; F2 uint64 }
type GenS17 struct { F0 uint8; F1 int8; F2 int }
type GenS18 struct { F0 int16; F1 int64 }
type GenS19 struct { F0 uint16; F1 int }
type GenS20 struct { F0 int8; F1 int16; F2 int8 }
type GenS21 struct { F0 bool; F1 uint64; F2 uint16 }
type GenS22 struct { F0 uint64; F1 uint16; F2 uint }
type GenS23 struct { F0 uint8; F1 int32; F2 int64 }
type GenS24 struct { F0 int32; F1 int16 }
type GenS25 struct { F0 int16; F1 uint; F2 int8 }
type GenS26 struct { F0 int8; F1 int16; F2 int16 }
type GenS27 struct { F0 uint8 }
type GenS28 struct { F0 uint16 }
type GenS29 struct { F0 int8 }
type GenS30 struct { F0 int; F1 string }
type GenS31 struct { F0 uint; F1 int64; F2 uint16 }
type GenS32 struct { F0 uint8; F1 bool }
func genFn0(p0 uint, p1 int16) int {
    p0 = p0
    v5 := []int{831}
    _ = p0
    _ = p1
    _ = v5
    return 32768
}

func main() {
    var v1 int = 0
    v2 := make(map[int]string)
    var v3 string = "hello"
    var v4 uint32 = 377
    if ((true || true) || (true || true)) {
        v3 = v3
        v6 := GenS1{F0: 65535, F1: -125}
        var v7 int = (v1 % ((v1) | 1))
        _ = v6
        _ = v7
    } else {
        if false {
            for fi8 := 0; fi8 < 13; fi8++ {
                v2[genFn0(4294967295, 397)] = v3
                v9 := GenS1{F0: 85, F1: 1}
                var v10 int32 = 0
                _ = fi8
                _ = v10
                _ = v9
            }
            fmt.Println(v4, v1)
            switch v1 {
            case -9223372036854775808:
                if (v4 == v4) {
                    v1 += genFn0(2, 32767)
                } else {
                    v4 = v4
                    var v11 bool = (true && true)
                    _ = v11
                }
            case 2147483647:
                var v12 uint64 = 708
                _ = v12
            default:
                for fi13 := 0; fi13 < 17; fi13++ {
                    v3 = (v3 + (v3 + "\n\t"))
                    _ = fi13
                }
            }
        }
    }
    v14 := []int{v1, (^(v1)), 9223372036854775807, genFn0(255, 32767)}
    v15 := v2[(v1 << ((42) & 63))]
    _ = len(v2)
    var v16 int32 = 0
    v17 := GenS1{F0: 32767, F1: -54}
    if (203 > 127) {
        _ = len(v2)
    } else {
        switch (((v15 + "foo bar") + v3) + v15) {
        case "a":
            var v18 uint8 = 255
            _ = v18
        default:
            if true {
                var v19 float32 = ((-2.25 * -2.25) / 3.14159)
                fmt.Println(v19)
                _ = v19
            } else {
                v20 := GenS2{F0: v16, F1: 0, F2: v4}
                if true {
                    v21 := GenS3{F0: (float32((1e10 * 1.5)) < 1.5), F1: 1}
                    _ = v21
                } else {
                    if (!false) {
                        switch v1 {
                        case 692:
                            var v22 float32 = 1.5
                            _ = v22
                        }
                        if true {
                            if len(v14) > 0 {
                                v14[0] = v1
                            }
                            delete(v2, 267)
                            fmt.Println(v1)
                        }
                    } else {
                        fmt.Println(v3)
                        v14 = append(v14, 272)
                    }
                }
                _ = v20
            }
        }
    }
    if (!false) {
        for fi23 := 0; fi23 < 9; fi23++ {
            v14 = append(v14, v1)
            for fi24 := 0; fi24 < 2; fi24++ {
                switch fi24 {
                case 530:
                    v15 = v3
                }
                fmt.Println(fi23, fi24)
                if len(v14) > 0 {
                    v14[0] = v1
                }
                _ = fi24
            }
            _ = fi23
        }
        for fi25 := 0; fi25 < 16; fi25++ {
            v26 := GenS4{F0: 0, F1: 228}
            if len(v14) >= 2 {
                v14 = v14[1:len(v14)]
            }
            v15 = (((v3 + v15) + ("foo bar" + "\n\t")) + (v3 + "a"))
            _ = fi25
            _ = v26
        }
    } else {
        fmt.Println(v15, v15)
        v16 += -2147483648
    }
    for fi27 := 0; fi27 < 18; fi27++ {
        for fi28 := 0; fi28 < 10; fi28++ {
            if true {
                if len(v14) > 0 {
                    v14[0] = (fi28 + fi28)
                }
                var v29 uint32 = (v4 / ((4294967295) | 1))
                switch "\n\t" {
                case "\n\t":
                    fmt.Println(v16, v29, v4)
                case "":
                    switch ((v15 + v3) + v3) {
                    case "\n\t":
                        v2[genFn0(18446744073709551615, 32767)] = ("a" + v15)
                    case "hello":
                        switch fi28 {
                        case 25:
                            v30 := v2[fi27]
                            _ = v30
                        case -896:
                            switch ("hello" + v3) {
                            case "hello":
                                var v31 int8 = 75
                                _ = v31
                            case "":
                                var v32 string = v3
                                _ = v32
                            }
                            fmt.Println(v3, v16)
                        case 32767:
                            var v33 uint8 = 14
                            _ = v33
                        }
                    default:
                        if true {
                            var v34 uint64 = 18446744073709551615
                            switch (fi28 | fi28) {
                            case 32768:
                                var v35 uint8 = 158
                                _ = v35
                            case 127:
                                v36 := GenS5{F0: false, F1: (v34 % (((v34 >> ((42) & 63))) | 1))}
                                _ = v36
                            case 2147483647:
                                var v37 bool = ((true && ((0 >= 81) || ((v4 < 0) || true))) || false)
                                _ = v37
                            default:
                                v4 -= (v29 ^ (^(v4)))
                            }
                            _ = v34
                        } else {
                            if (true && true) {
                                _ = len(v14)
                            }
                        }
                    }
                default:
                    fmt.Println(v29, v4)
                }
                _ = v29
            } else {
                var v38 int16 = -32768
                fmt.Println(v4, fi27)
                _ = v38
            }
            if (v4 == (^(v4))) {
                var v39 int64 = -298
                if true {
                    fmt.Println(v39, v16)
                    fmt.Println(fi27, v4, v1)
                    switch (v15 + v3) {
                    case "\n\t":
                        fmt.Println(v1)
                    case "foo bar":
                        v1 += 4294967295
                        if false {
                            var v40 int64 = v39
                            _ = v40
                        } else {
                            var v41 uint32 = 2147483648
                            var v42 uint8 = 1
                            var v43 int64 = (v39 - v39)
                            _ = v41
                            _ = v42
                            _ = v43
                        }
                        var v44 float64 = 0.1
                        _ = v44
                    }
                    switch "hello" {
                    case "hello":
                        if false {
                            v45 := v2[9223372036854775807]
                            switch int(v16) {
                            case 9223372036854775807:
                                switch (v1 & genFn0(18446744073709551615, -723)) {
                                case -510:
                                    if true {
                                        v39 = 143
                                        _ = len(v2)
                                    }
                                case 922:
                                    var v46 uint32 = 0
                                    _ = v46
                                case 127:
                                    switch -1 {
                                    case 9223372036854775807:
                                        fmt.Println(v4, fi28, v1)
                                    }
                                }
                                v45 = v15
                                v39 -= (v39 << ((1) & 63))
                                var v47 float64 = ((((1e10 - (1e10 * -2.25)) / (((0.0 - 3.14159) + -2.25) / 1e10)) * 1.5) / 0.1)
                                v1 -= genFn0(1, -1)
                                if false {
                                    switch v15 {
                                    case "a":
                                        switch "\n\t" {
                                        case "\n\t":
                                            if true {
                                                if true {
                                                    var v48 uint = 298
                                                    v2[v1] = "a"
                                                    fmt.Println(v15, fi27)
                                                    var v49 uint16 = 42
                                                    _ = v48
                                                    _ = v49
                                                } else {
                                                    fmt.Println(v3, v3, v45)
                                                    v3 = v15
                                                    v2[genFn0(250, -11)] = v3
                                                }
                                                if len(v14) >= 2 {
                                                    v14 = v14[1:len(v14)]
                                                }
                                            } else {
                                                fmt.Println(v47, v3, fi28)
                                                if (1 < 4294967295) {
                                                    v16 = v16
                                                    v1 = (fi27 - genFn0(4294967295, 1000))
                                                } else {
                                                    var v50 uint32 = (v4 ^ v4)
                                                    _ = v50
                                                }
                                                switch 4294967295 {
                                                case 32768:
                                                    v47 -= 0.0
                                                }
                                                if (42 < 544) {
                                                    var v51 uint16 = 0
                                                    _ = len(v2)
                                                    fmt.Println(v15, fi28)
                                                    v52 := GenS4{F0: 1, F1: 128}
                                                    _ = v51
                                                    _ = v52
                                                } else {
                                                    v39 = 9223372036854775807
                                                    v53 := GenS4{F0: 42, F1: 42}
                                                    _ = v53
                                                }
                                                fmt.Println(v15, fi27, v45)
                                            }
                                        case "foo bar":
                                            var v54 float32 = (float32(v47) / ((1.5 / 1e10) - (1e10 / 3.14159)))
                                            _ = v54
                                        default:
                                            var v55 int64 = (^(v39))
                                            _ = v55
                                        }
                                        if false {
                                            delete(v2, v1)
                                        }
                                        v56 := GenS1{F0: 178, F1: 127}
                                        _ = v56
                                    default:
                                        switch 4294967295 {
                                        case 2147483647:
                                            var v57 uint64 = 4294967295
                                            _ = v57
                                        }
                                    }
                                    v58 := GenS1{F0: 32767, F1: 127}
                                    _ = v58
                                }
                                v1 += fi27
                                var v59 int16 = -847
                                v39 += 42
                                _ = v47
                                _ = v59
                            case 585:
                                var v60 int16 = 42
                                _ = v60
                            case -130:
                                var v61 uint32 = (v4 & (v4 % (((v4 << ((540) & 63))) | 1)))
                                _ = v61
                            }
                            if (1000 == 1000) {
                                switch (v15 + "a") {
                                case "":
                                    v1 = (^(fi27))
                                default:
                                    if ((0.1 / 1.0) > 0.0) {
                                        switch ("hello" + v3) {
                                        case "":
                                            v62 := GenS2{F0: (v16 / ((v16) | 1)), F1: 128, F2: 42}
                                            _ = v62
                                        case "hello":
                                            v16 = (^(v16))
                                        case "foo bar":
                                            v63 := GenS4{F0: 4294967295, F1: 128}
                                            _ = v63
                                        }
                                    } else {
                                        v4 = 0
                                        switch v45 {
                                        case "\n\t":
                                            v64 := GenS6{F0: -9}
                                            _ = v64
                                        case "foo bar":
                                            fmt.Println(v39, v4, v15)
                                        case "a":
                                            switch (((v15 + (v15 + "foo bar")) + "\n\t") + v3) {
                                            case "a":
                                                v45 = v45
                                                switch (v3 + ("a" + v45)) {
                                                case "a":
                                                    _ = len(v2)
                                                case "foo bar":
                                                    if true {
                                                        v4 = 889
                                                    } else {
                                                        v4 -= (^(v4))
                                                    }
                                                case "":
                                                    v3 = ("a" + v15)
                                                }
                                                switch v1 {
                                                case 699:
                                                    v65 := GenS7{F0: ("a" == "a"), F1: 65535}
                                                    _ = v65
                                                case -2147483648:
                                                    if true {
                                                        fmt.Println(v15, v3, v39)
                                                    }
                                                case 32768:
                                                    var v66 uint64 = 1
                                                    _ = v66
                                                default:
                                                    v3 = v3
                                                }
                                                v4 = 1
                                            case "hello":
                                                v4 = 42
                                            default:
                                                switch (fi28 % ((fi27) | 1)) {
                                                case 32768:
                                                    var v67 string = v3
                                                    _ = v67
                                                case -1:
                                                    var v68 uint8 = 127
                                                    _ = v68
                                                case 4294967295:
                                                    switch (v45 + v45) {
                                                    case "foo bar":
                                                        fmt.Println(v39)
                                                    case "":
                                                        if len(v14) > 0 {
                                                            v14[0] = 2147483647
                                                        }
                                                    }
                                                default:
                                                    if false {
                                                        fmt.Println(v45, v3)
                                                        v69 := GenS2{F0: v16, F1: 0, F2: 42}
                                                        _ = v69
                                                    } else {
                                                        var v70 int32 = v16
                                                        var v71 int16 = 511
                                                        var v72 int8 = 127
                                                        _ = v70
                                                        _ = v71
                                                        _ = v72
                                                    }
                                                }
                                            }
                                        }
                                        v45 = v3
                                    }
                                }
                                v4 -= (v4 % ((v4) | 1))
                            }
                            _ = v45
                        } else {
                            v16 = v16
                            v73 := GenS4{F0: 296, F1: 127}
                            v74 := GenS2{F0: v16, F1: 0, F2: 2147483647}
                            if (18446744073709551615 < 586) {
                                v75 := GenS8{F0: "hello", F1: -1, F2: "a"}
                                v76 := GenS9{F0: 2, F1: 4294967295}
                                _ = v75
                                _ = v76
                            } else {
                                delete(v2, fi27)
                            }
                            _ = v73
                            _ = v74
                        }
                    default:
                        if (!false) {
                            v77 := GenS8{F0: (v3 + v15), F1: v39, F2: ("hello" + v3)}
                            _ = v77
                        }
                        switch (v3 + (("" + v3) + v3)) {
                        case "hello":
                            switch v3 {
                            case "a":
                                if len(v14) >= 2 {
                                    v14 = v14[1:len(v14)]
                                }
                            case "hello":
                                fmt.Println(v39)
                                v3 = (v3 + "hello")
                                v78 := GenS10{F0: 235, F1: v15, F2: v16}
                                if true {
                                    if true {
                                        v1 += (fi27 % (((v1 * fi27)) | 1))
                                        v16 += v16
                                        v79 := GenS5{F0: false, F1: 2}
                                        _ = v79
                                    }
                                    v80 := GenS11{F0: v39}
                                    _ = len(v2)
                                    v81 := GenS7{F0: false, F1: 2}
                                    v3 = v3
                                    _ = v80
                                    _ = v81
                                }
                                _ = v78
                            }
                        case "a":
                            var v82 uint64 = 1
                            _ = v82
                        case "foo bar":
                            switch "a" {
                            case "hello":
                                if (!(!(!(!true)))) {
                                    var v83 bool = ((true || (!(0 >= -11))) && false)
                                    _ = v83
                                }
                            }
                        default:
                            if (v3 == ("foo bar" + v15)) {
                                if true {
                                    v84 := GenS12{F0: int16(v39)}
                                    switch fi27 {
                                    case 127:
                                        v85 := GenS10{F0: 4, F1: (v15 + "foo bar"), F2: v16}
                                        _ = v85
                                    case 32768:
                                        v39 -= 250
                                    case -112:
                                        if (!(v3 != v3)) {
                                            var v86 int64 = (v39 ^ -9223372036854775808)
                                            _ = v86
                                        } else {
                                            v14 = append(v14, int(v39))
                                            switch 2147483647 {
                                            case 128:
                                                v4 = v4
                                            case 32767:
                                                fmt.Println(v15, v16)
                                            case -158:
                                                var v87 string = (v3 + v3)
                                                _ = v87
                                            default:
                                                switch "\n\t" {
                                                case "":
                                                    v88 := GenS4{F0: 92, F1: 2}
                                                    v3 = ""
                                                    _ = v88
                                                case "hello":
                                                    var v89 string = v15
                                                    _ = v89
                                                default:
                                                    v39 = (v39 >> ((85) & 63))
                                                }
                                            }
                                        }
                                    default:
                                        if len(v14) >= 2 {
                                            v14 = v14[1:len(v14)]
                                        }
                                    }
                                    var v90 uint32 = (v4 | v4)
                                    _ = v84
                                    _ = v90
                                } else {
                                    v91 := GenS6{F0: v16}
                                    var v92 int64 = v39
                                    switch (v3 + "a") {
                                    case "foo bar":
                                        v93 := GenS13{F0: (v3 + (v3 + ("foo bar" + v3))), F1: 2, F2: v4}
                                        _ = v93
                                    case "":
                                        v94 := GenS13{F0: v15, F1: 255, F2: (v4 & v4)}
                                        _ = v94
                                    case "a":
                                        v95 := GenS10{F0: 130, F1: v3, F2: v16}
                                        _ = v95
                                    default:
                                        var v96 uint8 = 127
                                        _ = v96
                                    }
                                    if true {
                                        var v97 float32 = 0.0
                                        _ = v97
                                    } else {
                                        if ((!((fi28 ^ 32768) != int(v16))) && false) {
                                            v4 = v4
                                            v3 = v3
                                            v98 := GenS3{F0: false, F1: 65535}
                                            _ = v98
                                        } else {
                                            v99 := GenS2{F0: 1, F1: 2, F2: 2147483648}
                                            var v100 uint16 = 203
                                            delete(v2, fi28)
                                            _ = v100
                                            _ = v99
                                        }
                                    }
                                    _ = v91
                                    _ = v92
                                }
                                fmt.Println(v1)
                            } else {
                                if true {
                                    if true {
                                        var v101 int32 = v16
                                        _ = v101
                                    }
                                    v102 := GenS14{F0: v4}
                                    if len(v14) >= 2 {
                                        v14 = v14[1:len(v14)]
                                    }
                                    _ = v102
                                }
                                v103 := GenS15{F0: v4, F1: 18446744073709551615, F2: 2}
                                v104 := GenS16{F0: 65535, F1: 451, F2: 42}
                                switch (("" + ("\n\t" + v3)) + ("foo bar" + "\n\t")) {
                                case "a":
                                    fmt.Println(fi27)
                                }
                                switch v3 {
                                case "a":
                                    var v105 uint = 42
                                    _ = v105
                                case "":
                                    v14 = append(v14, (v1 & genFn0(4294967295, -32768)))
                                case "\n\t":
                                    v39 = (v39 & (v39 ^ v39))
                                    if ((4294967295 < 972) || (false && true)) {
                                        if false {
                                            v1 += -407
                                            var v106 int64 = (v39 + 751)
                                            v2[fi27] = v15
                                            var v107 float32 = 3.14159
                                            _ = v106
                                            _ = v107
                                        } else {
                                            switch 1 {
                                            case 1:
                                                var v108 string = "\n\t"
                                                _ = v108
                                            case 297:
                                                switch (("" + "") + v3) {
                                                case "hello":
                                                    if (255 < 1) {
                                                        fmt.Println(v3)
                                                        switch 0 {
                                                        case 343:
                                                            var v109 int = (^(fi28))
                                                            _ = v109
                                                        default:
                                                            switch (v15 + ((v3 + "") + (v3 + v3))) {
                                                            case "foo bar":
                                                                v39 += -648
                                                            case "\n\t":
                                                                switch (fi27 + fi27) {
                                                                case -2147483648:
                                                                    fmt.Println(v4, v4, v15)
                                                                default:
                                                                    v39 = int64(fi27)
                                                                }
                                                            default:
                                                                v3 = ((v15 + v3) + (v3 + "a"))
                                                            }
                                                        }
                                                        var v110 int = 9223372036854775807
                                                        var v111 float32 = ((((((1e10 / -2.25) * 0.1) + 0.0) / -2.25) + ((0.0 / (float32(-2.25) * float32(0.1))) / 1e10)) - 0.1)
                                                        _ = v110
                                                        _ = v111
                                                    } else {
                                                        var v112 uint16 = 1
                                                        if false {
                                                            if false {
                                                                var v113 int64 = (v39 >> ((2) & 63))
                                                                v114 := GenS7{F0: (2 == (^(v4))), F1: (^(v112))}
                                                                if len(v14) > 0 {
                                                                    v14[0] = (fi27 & fi28)
                                                                }
                                                                _ = v113
                                                                _ = v114
                                                            } else {
                                                                _ = len(v14)
                                                                v115 := GenS11{F0: (v39 % ((int64(v4)) | 1))}
                                                                v16 += 0
                                                                _ = v115
                                                            }
                                                        } else {
                                                            v39 -= (v39 % ((v39) | 1))
                                                            var v116 uint32 = v4
                                                            _ = v116
                                                        }
                                                        _ = v112
                                                    }
                                                    v3 = v15
                                                default:
                                                    if true {
                                                        if false {
                                                            switch 2147483647 {
                                                            case 505:
                                                                v117 := GenS17{F0: 255, F1: 127, F2: (fi27 >> ((255) & 63))}
                                                                _ = v117
                                                            case 42:
                                                                v4 -= 0
                                                            case 989:
                                                                if true {
                                                                    fmt.Println(v4)
                                                                    switch (v3 + v15) {
                                                                    case "hello":
                                                                        var v118 int16 = 848
                                                                        _ = v118
                                                                    case "":
                                                                        switch int(v4) {
                                                                        case -1:
                                                                            v15 = "a"
                                                                            v16 -= 1
                                                                        case -2147483648:
                                                                            if (4294967295 >= 27) {
                                                                                var v119 uint64 = 155
                                                                                v39 -= v39
                                                                                _ = v119
                                                                            }
                                                                        default:
                                                                            var v120 uint64 = 113
                                                                            _ = v120
                                                                        }
                                                                    case "foo bar":
                                                                        var v121 float64 = (1.5 - float64((float32(-2.25) - 3.14159)))
                                                                        _ = v121
                                                                    default:
                                                                        v15 = v3
                                                                    }
                                                                    v122 := GenS18{F0: 1000, F1: (v39 & v39)}
                                                                    _ = v122
                                                                }
                                                            }
                                                        } else {
                                                            switch fi27 {
                                                            case -603:
                                                                v16 = (v16 & 361)
                                                            case -188:
                                                                var v123 uint64 = 0
                                                                _ = v123
                                                            case 32768:
                                                                switch v15 {
                                                                case "foo bar":
                                                                    var v124 int16 = -32768
                                                                    _ = v124
                                                                case "hello":
                                                                    v125 := GenS19{F0: 65535, F1: v1}
                                                                    switch v15 {
                                                                    case "\n\t":
                                                                        if (v39 != int64(v4)) {
                                                                            if ((0.1 != (1e10 * 1.5)) || (float64(3.14159) > float64((1e10 / float32(-2.25))))) {
                                                                                if (((float64(float32(float64(0.0))) * -2.25) / -2.25) == 1.5) {
                                                                                    v126 := GenS1{F0: 42, F1: 0}
                                                                                    v2[42] = (v3 + "")
                                                                                    _ = v126
                                                                                } else {
                                                                                    switch (fi27 << ((65535) & 63)) {
                                                                                    case 545:
                                                                                        v1 += genFn0(2, 42)
                                                                                    default:
                                                                                        fmt.Println(fi28, v3, v1)
                                                                                        switch v3 {
                                                                                        case "":
                                                                                            if false {
                                                                                                if false {
                                                                                                    if (0 != 2) {
                                                                                                        v127 := GenS20{F0: 127, F1: 42, F2: int8(v4)}
                                                                                                        _ = v127
                                                                                                    } else {
                                                                                                        var v128 int32 = (v16 << ((920) & 63))
                                                                                                        if (!(!(42 == 2))) {
                                                                                                            fmt.Println(v3)
                                                                                                        }
                                                                                                        if (((float32(1e10) - (0.1 + 3.14159)) / 1.5) == 0.0) {
                                                                                                            if (!(!false)) {
                                                                                                                if true {
                                                                                                                    var v129 int8 = 101
                                                                                                                    fmt.Println(v39, fi27)
                                                                                                                    var v130 uint = 65535
                                                                                                                    _ = v129
                                                                                                                    _ = v130
                                                                                                                }
                                                                                                                _ = len(v2)
                                                                                                                var v131 uint16 = 32767
                                                                                                                _ = v131
                                                                                                            } else {
                                                                                                                v132 := GenS1{F0: 0, F1: 127}
                                                                                                                _ = v132
                                                                                                            }
                                                                                                            _ = len(v14)
                                                                                                        } else {
                                                                                                            var v133 int32 = (^(v16))
                                                                                                            v128 += v128
                                                                                                            _ = v133
                                                                                                        }
                                                                                                        _ = v128
                                                                                                    }
                                                                                                    v4 = 835
                                                                                                }
                                                                                                if len(v14) > 0 {
                                                                                                    v14[0] = (fi27 - (v1 << ((2) & 63)))
                                                                                                }
                                                                                                if true {
                                                                                                    v2[fi27] = ("a" + ((v3 + v15) + v15))
                                                                                                    fmt.Println(v3)
                                                                                                }
                                                                                            }
                                                                                        case "a":
                                                                                            switch v15 {
                                                                                            case "hello":
                                                                                                if (!true) {
                                                                                                    v3 = v15
                                                                                                    fmt.Println(v16)
                                                                                                }
                                                                                            case "\n\t":
                                                                                                if (442 <= 205) {
                                                                                                    fmt.Println(v3)
                                                                                                }
                                                                                            case "foo bar":
                                                                                                v16 -= 42
                                                                                            }
                                                                                        default:
                                                                                            switch 0 {
                                                                                            case 42:
                                                                                                switch "a" {
                                                                                                case "hello":
                                                                                                    v16 += (v16 << ((1) & 63))
                                                                                                case "foo bar":
                                                                                                    switch v3 {
                                                                                                    case "foo bar":
                                                                                                        v16 += (v16 | v16)
                                                                                                    case "\n\t":
                                                                                                        v2[genFn0(22, 42)] = ("\n\t" + (v15 + v3))
                                                                                                    default:
                                                                                                        var v134 int64 = (v39 | 0)
                                                                                                        _ = v134
                                                                                                    }
                                                                                                default:
                                                                                                    switch genFn0(406, int16(v39)) {
                                                                                                    case 128:
                                                                                                        _ = len(v14)
                                                                                                    case -2147483648:
                                                                                                        v15 = (("\n\t" + v15) + v15)
                                                                                                    case -179:
                                                                                                        var v135 uint64 = 2
                                                                                                        _ = v135
                                                                                                    default:
                                                                                                        v4 = (v4 - (v4 ^ (v4 * v4)))
                                                                                                    }
                                                                                                }
                                                                                            case -485:
                                                                                                v15 = (v15 + ("\n\t" + v3))
                                                                                            case 4294967295:
                                                                                                if false {
                                                                                                    if (7 == 42) {
                                                                                                        var v136 int = v1
                                                                                                        v2[v136] = v15
                                                                                                        fmt.Println(v16, v4, fi28)
                                                                                                        v1 = v1
                                                                                                        var v137 float64 = (1.5 - (3.14159 * 1.5))
                                                                                                        var v138 uint16 = 1
                                                                                                        if ((127 < -107) || (!false)) {
                                                                                                            var v139 uint16 = v138
                                                                                                            if true {
                                                                                                                v140 := GenS21{F0: (v39 == v39), F1: 941, F2: 2}
                                                                                                                v141 := GenS15{F0: v4, F1: 2, F2: 47}
                                                                                                                v15 = (v3 + "hello")
                                                                                                                _ = v140
                                                                                                                _ = v141
                                                                                                            } else {
                                                                                                                if (v39 != v39) {
                                                                                                                    var v142 string = "foo bar"
                                                                                                                    var v143 string = ""
                                                                                                                    var v144 uint32 = 1
                                                                                                                    _ = v142
                                                                                                                    _ = v143
                                                                                                                    _ = v144
                                                                                                                } else {
                                                                                                                    _ = len(v2)
                                                                                                                }
                                                                                                                var v145 int8 = 100
                                                                                                                _ = v145
                                                                                                            }
                                                                                                            if false {
                                                                                                                var v146 uint = 0
                                                                                                                fmt.Println(v39, v139)
                                                                                                                _ = v146
                                                                                                            }
                                                                                                            _ = v139
                                                                                                        } else {
                                                                                                            v3 = v3
                                                                                                        }
                                                                                                        var v147 float32 = 0.1
                                                                                                        v136 += (fi27 * (fi27 >> ((2) & 63)))
                                                                                                        _ = v136
                                                                                                        _ = v137
                                                                                                        _ = v138
                                                                                                        _ = v147
                                                                                                    } else {
                                                                                                        v148 := v2[v1]
                                                                                                        if true {
                                                                                                            if (v16 < 0) {
                                                                                                                v39 += 42
                                                                                                                var v149 string = (v15 + "hello")
                                                                                                                if ((false || true) && true) {
                                                                                                                    var v150 bool = true
                                                                                                                    if len(v14) > 0 {
                                                                                                                        v14[0] = int(v39)
                                                                                                                    }
                                                                                                                    _ = v150
                                                                                                                }
                                                                                                                _ = v149
                                                                                                            } else {
                                                                                                                v151 := GenS22{F0: 18446744073709551615, F1: 255, F2: 42}
                                                                                                                var v152 int64 = (v39 % ((-9223372036854775808) | 1))
                                                                                                                _ = v151
                                                                                                                _ = v152
                                                                                                            }
                                                                                                        } else {
                                                                                                            v153 := v2[(fi27 - int(v4))]
                                                                                                            fmt.Println(v1, fi28, fi27)
                                                                                                            switch (((v3 + v153) + "\n\t") + ((v3 + ("hello" + v148)) + v3)) {
                                                                                                            case "foo bar":
                                                                                                                v148 = "a"
                                                                                                            case "hello":
                                                                                                                var v154 uint8 = 42
                                                                                                                _ = v154
                                                                                                            case "\n\t":
                                                                                                                v155 := v2[(-(v1))]
                                                                                                                _ = v155
                                                                                                            default:
                                                                                                                var v156 int32 = v16
                                                                                                                if len(v14) >= 2 {
                                                                                                                    v14 = v14[1:len(v14)]
                                                                                                                }
                                                                                                                _ = v156
                                                                                                            }
                                                                                                            _ = v153
                                                                                                        }
                                                                                                        _ = v148
                                                                                                    }
                                                                                                } else {
                                                                                                    v14 = append(v14, (fi28 - v1))
                                                                                                }
                                                                                            default:
                                                                                                switch v3 {
                                                                                                case "foo bar":
                                                                                                    var v157 float64 = (-2.25 / 1.0)
                                                                                                    fmt.Println(v157, v15)
                                                                                                    _ = v157
                                                                                                case "a":
                                                                                                    var v158 int16 = 1
                                                                                                    _ = v158
                                                                                                }
                                                                                            }
                                                                                        }
                                                                                        if true {
                                                                                            if ((genFn0(2, 1) >= fi28) && false) {
                                                                                                switch ("\n\t" + v15) {
                                                                                                case "\n\t":
                                                                                                    if false {
                                                                                                        var v159 uint = 255
                                                                                                        if false {
                                                                                                            if ((false && ((v39 > v39) && false)) && (!true)) {
                                                                                                                fmt.Println(v15)
                                                                                                                v3 = "a"
                                                                                                                v160 := GenS15{F0: v4, F1: 42, F2: 42}
                                                                                                                if (true && true) {
                                                                                                                    switch (v3 + (v15 + "")) {
                                                                                                                    case "\n\t":
                                                                                                                        v161 := GenS23{F0: 42, F1: v16, F2: v39}
                                                                                                                        _ = v161
                                                                                                                    case "":
                                                                                                                        switch "a" {
                                                                                                                        case "":
                                                                                                                            var v162 int64 = v39
                                                                                                                            _ = v162
                                                                                                                        }
                                                                                                                    }
                                                                                                                    var v163 string = v15
                                                                                                                    v163 = v163
                                                                                                                    _ = v163
                                                                                                                }
                                                                                                                _ = v160
                                                                                                            } else {
                                                                                                                var v164 uint64 = 0
                                                                                                                _ = len(v14)
                                                                                                                v165 := GenS5{F0: true, F1: (v164 / ((v164) | 1))}
                                                                                                                _ = v164
                                                                                                                _ = v165
                                                                                                            }
                                                                                                        } else {
                                                                                                            switch fi27 {
                                                                                                            case 9223372036854775807:
                                                                                                                if false {
                                                                                                                    fmt.Println(v1)
                                                                                                                } else {
                                                                                                                    v1 += genFn0(uint(v4), 1)
                                                                                                                    v159 = 255
                                                                                                                    fmt.Println(v1)
                                                                                                                    if (-2.25 > float32(((-2.25 - 0.1) * (0.0 + (0.1 / (-2.25 * (3.14159 - 1.5))))))) {
                                                                                                                        var v166 uint8 = 211
                                                                                                                        v167 := GenS23{F0: (v166 >> ((v159) & 63)), F1: (v16 >> ((196) & 63)), F2: v39}
                                                                                                                        _ = v166
                                                                                                                        _ = v167
                                                                                                                    } else {
                                                                                                                        _ = len(v14)
                                                                                                                    }
                                                                                                                }
                                                                                                            case 1:
                                                                                                                switch v3 {
                                                                                                                case "":
                                                                                                                    var v168 float32 = 0.0
                                                                                                                    _ = v168
                                                                                                                case "a":
                                                                                                                    var v169 int8 = -128
                                                                                                                    _ = v169
                                                                                                                default:
                                                                                                                    v170 := v2[int(v39)]
                                                                                                                    _ = v170
                                                                                                                }
                                                                                                            }
                                                                                                        }
                                                                                                        _ = v159
                                                                                                    } else {
                                                                                                        var v171 uint = 479
                                                                                                        v172 := GenS17{F0: 1, F1: -70, F2: genFn0(v171, 32767)}
                                                                                                        _ = v171
                                                                                                        _ = v172
                                                                                                    }
                                                                                                case "a":
                                                                                                    v14 = append(v14, (fi27 ^ genFn0(42, 32767)))
                                                                                                case "":
                                                                                                    if (80 <= 1) {
                                                                                                        var v173 uint8 = 128
                                                                                                        var v174 int16 = 42
                                                                                                        _ = v173
                                                                                                        _ = v174
                                                                                                    }
                                                                                                default:
                                                                                                    var v175 uint16 = 2
                                                                                                    _ = v175
                                                                                                }
                                                                                            } else {
                                                                                                if true {
                                                                                                    if len(v14) > 0 {
                                                                                                        v14[0] = int(v16)
                                                                                                    }
                                                                                                    v176 := GenS22{F0: 42, F1: 32767, F2: 1}
                                                                                                    delete(v2, fi28)
                                                                                                    _ = v176
                                                                                                } else {
                                                                                                    var v177 bool = true
                                                                                                    v178 := v2[genFn0(421, 726)]
                                                                                                    var v179 int = (fi27 / ((fi27) | 1))
                                                                                                    _ = v177
                                                                                                    _ = v178
                                                                                                    _ = v179
                                                                                                }
                                                                                                v14 = append(v14, int(v16))
                                                                                                switch fi27 {
                                                                                                case 128:
                                                                                                    v180 := GenS24{F0: v16, F1: -1}
                                                                                                    _ = v180
                                                                                                case 127:
                                                                                                    switch fi28 {
                                                                                                    case 2147483647:
                                                                                                        switch v15 {
                                                                                                        case "hello":
                                                                                                            switch fi28 {
                                                                                                            case -9223372036854775808:
                                                                                                                v4 = v4
                                                                                                            case 127:
                                                                                                                fmt.Println(fi27, fi28)
                                                                                                            case -43:
                                                                                                                fmt.Println(v4)
                                                                                                            default:
                                                                                                                var v181 uint32 = (v4 * 1)
                                                                                                                _ = v181
                                                                                                            }
                                                                                                        }
                                                                                                    case 127:
                                                                                                        var v182 uint8 = 255
                                                                                                        _ = v182
                                                                                                    case -1:
                                                                                                        fmt.Println(v1, v16)
                                                                                                    }
                                                                                                }
                                                                                            }
                                                                                            v183 := GenS25{F0: 1, F1: 198, F2: 1}
                                                                                            _ = v183
                                                                                        }
                                                                                    }
                                                                                    v15 = v15
                                                                                }
                                                                            }
                                                                        }
                                                                    case "hello":
                                                                        var v184 string = v15
                                                                        var v185 bool = false
                                                                        switch v3 {
                                                                        case "hello":
                                                                            v4 -= v4
                                                                        case "a":
                                                                            switch "" {
                                                                            case "":
                                                                                var v186 int8 = 22
                                                                                _ = v186
                                                                            default:
                                                                                if v185 {
                                                                                    var v187 string = (v3 + v184)
                                                                                    var v188 float64 = -2.25
                                                                                    switch fi28 {
                                                                                    case 128:
                                                                                        switch -2147483648 {
                                                                                        case -247:
                                                                                            switch (v3 + "") {
                                                                                            case "a":
                                                                                                switch "hello" {
                                                                                                case "":
                                                                                                    v189 := GenS11{F0: (v39 ^ (v39 & v39))}
                                                                                                    _ = v189
                                                                                                default:
                                                                                                    if (42 <= -42) {
                                                                                                        v39 = (v39 >> ((1) & 63))
                                                                                                        switch "\n\t" {
                                                                                                        case "hello":
                                                                                                            v190 := GenS26{F0: 127, F1: -809, F2: 42}
                                                                                                            _ = v190
                                                                                                        default:
                                                                                                            fmt.Println(fi28, v15, v184)
                                                                                                        }
                                                                                                    } else {
                                                                                                        var v191 uint = 0
                                                                                                        _ = v191
                                                                                                    }
                                                                                                }
                                                                                            case "foo bar":
                                                                                                if len(v14) >= 2 {
                                                                                                    v14 = v14[1:len(v14)]
                                                                                                }
                                                                                            default:
                                                                                                var v192 int16 = -32768
                                                                                                _ = v192
                                                                                            }
                                                                                        case 9223372036854775807:
                                                                                            v193 := GenS27{F0: 2}
                                                                                            _ = v193
                                                                                        default:
                                                                                            if len(v14) > 0 {
                                                                                                v14[0] = (v1 % (((fi27 * (v1 % ((fi28) | 1)))) | 1))
                                                                                            }
                                                                                        }
                                                                                    case 9223372036854775807:
                                                                                        var v194 bool = (v4 == (^(v4)))
                                                                                        _ = v194
                                                                                    default:
                                                                                        var v195 float64 = v188
                                                                                        _ = v195
                                                                                    }
                                                                                    _ = v187
                                                                                    _ = v188
                                                                                }
                                                                            }
                                                                        case "foo bar":
                                                                            v16 -= 0
                                                                        }
                                                                        _ = v184
                                                                        _ = v185
                                                                    }
                                                                    _ = v125
                                                                }
                                                            }
                                                            var v196 uint = 826
                                                            var v197 uint16 = 2
                                                            _ = v196
                                                            _ = v197
                                                        }
                                                        var v198 float64 = float64(3.14159)
                                                        v15 = (v3 + v3)
                                                        _ = v198
                                                    } else {
                                                        v4 += 321
                                                        delete(v2, -9223372036854775808)
                                                    }
                                                }
                                            }
                                            v16 += v16
                                            v199 := GenS23{F0: 1, F1: (^(v16)), F2: v39}
                                            _ = v199
                                        }
                                        v200 := GenS19{F0: 1, F1: 127}
                                        v201 := GenS24{F0: -792, F1: -32768}
                                        _ = v200
                                        _ = v201
                                    }
                                    v202 := GenS4{F0: 255, F1: 156}
                                    v203 := v2[914]
                                    var v204 uint16 = 42
                                    if false {
                                        var v205 uint16 = 32767
                                        var v206 int64 = (v39 * (^(v39)))
                                        var v207 uint = 0
                                        _ = v205
                                        _ = v206
                                        _ = v207
                                    }
                                    switch v3 {
                                    case "":
                                        if (true && true) {
                                            if true {
                                                if len(v14) > 0 {
                                                    v14[0] = v1
                                                }
                                                switch (fi28 << ((629) & 63)) {
                                                case 32767:
                                                    v208 := GenS28{F0: v204}
                                                    _ = v208
                                                }
                                                switch (((v3 + v203) + v203) + v203) {
                                                case "hello":
                                                    if false {
                                                        switch v203 {
                                                        case "":
                                                            v209 := GenS29{F0: 0}
                                                            _ = v209
                                                        default:
                                                            switch genFn0(329, 42) {
                                                            case -693:
                                                                v15 = ("hello" + v15)
                                                            case -326:
                                                                if false {
                                                                    var v210 uint16 = v204
                                                                    fmt.Println(fi28, v210, v39)
                                                                    _ = v210
                                                                }
                                                            }
                                                        }
                                                        v211 := GenS24{F0: 31, F1: 1000}
                                                        v212 := GenS30{F0: 42, F1: "a"}
                                                        _ = v211
                                                        _ = v212
                                                    }
                                                default:
                                                    v213 := GenS30{F0: v1, F1: v15}
                                                    _ = v213
                                                }
                                            } else {
                                                fmt.Println(v39)
                                                v1 += int(v39)
                                            }
                                            if ((true && (42 != 956)) && (((!((float64(1.5) * (float64(1e10) / 1.0)) >= float64(float32((0.1 + (3.14159 - 0.0)))))) && ((((0.1 + 1.5) + 0.0) - 0.0) <= 0.1)) && true)) {
                                                switch "foo bar" {
                                                case "hello":
                                                    v203 = v15
                                                default:
                                                    v204 -= v204
                                                }
                                                if len(v14) >= 2 {
                                                    v14 = v14[1:len(v14)]
                                                }
                                                v15 = (v3 + v15)
                                            } else {
                                                fmt.Println(v204, v15, fi27)
                                                _ = len(v14)
                                                if false {
                                                    var v214 int32 = (v16 ^ 0)
                                                    fmt.Println(fi28, v203, v203)
                                                    var v215 int32 = (v16 ^ int32(v39))
                                                    _ = v214
                                                    _ = v215
                                                }
                                                v14 = append(v14, (fi28 * (^(fi27))))
                                                var v216 uint16 = v204
                                                v217 := GenS4{F0: 42, F1: 127}
                                                var v218 uint64 = 659
                                                _ = v216
                                                _ = v217
                                                _ = v218
                                            }
                                        } else {
                                            delete(v2, v1)
                                        }
                                    case "\n\t":
                                        var v219 float32 = -2.25
                                        _ = v219
                                    default:
                                        v220 := GenS27{F0: 0}
                                        _ = v220
                                    }
                                    fmt.Println(v4, v16, fi27)
                                    v3 = v203
                                    _ = v202
                                    _ = v203
                                    _ = v204
                                default:
                                    delete(v2, (fi28 << ((752) & 63)))
                                }
                                switch ((("foo bar" + v3) + ("hello" + v15)) + (((v3 + "foo bar") + v3) + (((v15 + "a") + "a") + ("hello" + "a")))) {
                                case "hello":
                                    v16 = -2147483648
                                case "":
                                    var v221 int32 = (v16 % ((-185) | 1))
                                    _ = v221
                                case "\n\t":
                                    v1 = (fi28 << ((255) & 63))
                                default:
                                    var v222 int8 = 1
                                    _ = v222
                                }
                                if (-102 >= -87) {
                                    fmt.Println(v1)
                                    v2[(-(v1))] = ("" + v3)
                                    delete(v2, (fi28 << ((255) & 63)))
                                }
                                _ = v103
                                _ = v104
                            }
                        }
                        if true {
                            if true {
                                if ((v16 + int32(v4)) != (v16 << ((2) & 63))) {
                                    var v223 string = (((("" + ("\n\t" + "a")) + "\n\t") + (((v15 + v3) + v3) + v3)) + (v15 + v15))
                                    var v224 int64 = v39
                                    _ = v223
                                    _ = v224
                                }
                                v225 := GenS31{F0: 65535, F1: v39, F2: 255}
                                switch ("" + "a") {
                                case "":
                                    var v226 string = v15
                                    _ = v226
                                default:
                                    var v227 int8 = 27
                                    _ = v227
                                }
                                var v228 uint8 = 2
                                fmt.Println(v39, v1)
                                _ = v225
                                _ = v228
                            }
                        }
                    }
                }
                if true {
                    var v229 uint16 = 2
                    switch v3 {
                    case "hello":
                        if true {
                            if (42 == 100) {
                                fmt.Println(v15, v16, v3)
                                v230 := GenS7{F0: (v39 < (v39 / (((v39 / (((^(v39))) | 1))) | 1))), F1: uint16(v4)}
                                v15 = v3
                                _ = v230
                            }
                            if len(v14) >= 2 {
                                v14 = v14[1:len(v14)]
                            }
                            v229 = 0
                        }
                    case "":
                        v231 := GenS3{F0: true, F1: 2}
                        _ = v231
                    }
                    delete(v2, v1)
                    fmt.Println(v4)
                    _ = v229
                }
                _ = v39
            }
            v3 = ("hello" + v3)
            _ = fi28
        }
        if false {
            if (-32768 <= 32767) {
                _ = len(v2)
            } else {
                v232 := GenS32{F0: 131, F1: ((v16 << ((2) & 63)) >= (v16 + 1))}
                for fi233 := 0; fi233 < 19; fi233++ {
                    if false {
                        var v234 float32 = ((0.1 - 0.1) * 1e10)
                        switch fi233 {
                        case 0:
                            var v235 bool = true
                            _ = v235
                        case 1:
                            var v236 int64 = 1
                            _ = v236
                        case 2147483647:
                            v16 += (v16 & 42)
                        }
                        _ = v234
                    } else {
                        var v237 int = (fi233 / (((^(fi27))) | 1))
                        _ = v237
                    }
                    var v238 int64 = 0
                    _ = fi233
                    _ = v238
                }
                _ = v232
            }
        }
        _ = fi27
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}
