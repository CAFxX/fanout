package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32 }
type GenS2 struct { F0 int32; F1 bool; F2 int32 }
type GenS3 struct { F0 int; F1 int64 }
type GenS4 struct { F0 uint16; F1 int64 }
type GenS5 struct { F0 int; F1 int64 }
type GenS6 struct { F0 uint32; F1 int16 }
type GenS7 struct { F0 uint8; F1 uint64; F2 uint16 }
type GenS8 struct { F0 bool; F1 int8 }
type GenS9 struct { F0 int16 }

func main() {
    var v1 string = "hello"
    var v2 bool = false
    v3 := make(map[uint32]bool)
    fmt.Println(v2)
    for fi4 := 0; fi4 < 4; fi4++ {
        v1 = ""
        var v5 uint64 = 735
        _ = fi4
        _ = v5
    }
    if (!((100 > 1) || v2)) {
        switch v1 {
        case "\n\t":
            var v6 uint32 = 2147483647
            _ = v6
        case "foo bar":
            var v7 uint16 = 703
            _ = v7
        case "a":
            v2 = (v2 || (((!v2) && v2) && (true || (0 < 1))))
        }
    }
    var v8 int16 = -475
    if (100 <= 38) {
        var v9 bool = (9 == 255)
        fmt.Println(v1, v1)
        _ = v9
    }
    v10 := GenS1{F0: 4294967295, F1: 0}
    fmt.Println(v2, v8, v1)
    var v11 uint64 = 238
    if (127 >= 127) {
        fmt.Println(v11, v2)
        switch v1 {
        case "hello":
            var v12 float32 = 3.14159
            _ = v12
        default:
            v13 := GenS1{F0: -1, F1: 2147483647}
            _ = v13
        }
    }
    if v2 {
        for fi14 := 0; fi14 < 11; fi14++ {
            v15 := GenS1{F0: (fi14 / ((fi14) | 1)), F1: 65535}
            _ = fi14
            _ = v15
        }
    } else {
        v2 = ((-94 <= 127) || v2)
        v16 := []int{-736}
        _ = v16
    }
    if (32768 >= -432) {
        switch -1 {
        case 9:
            var v17 uint = 253
            _ = v17
        case 323:
            var v18 int = -9223372036854775808
            _ = v18
        case 32768:
            var v19 uint32 = 65
            _ = v19
        default:
            if v2 {
                var v20 uint = 1
                _ = v20
            } else {
                fmt.Println(v1, v1, v8)
                delete(v3, 961)
            }
        }
        v21 := GenS2{F0: 42, F1: v2, F2: 1}
        _ = v21
    }
    var v22 int64 = 1
    var v23 int32 = -1
    for fi24 := 0; fi24 < 6; fi24++ {
        if (false || ((-(fi24)) <= fi24)) {
            v11 += (v11 & (v11 << ((0) & 63)))
            v25 := GenS2{F0: 1, F1: true, F2: (^(v23))}
            var v26 uint16 = 42
            _ = v25
            _ = v26
        } else {
            if (334 > 255) {
                if (v2 || true) {
                    v27 := GenS1{F0: fi24, F1: 1}
                    var v28 uint32 = 2147483647
                    _ = v27
                    _ = v28
                } else {
                    fmt.Println(v2)
                    v3[4294967295] = ("\n\t" <= "foo bar")
                }
                var v29 int8 = -128
                _ = v29
            }
        }
        var v30 int64 = 460
        switch "foo bar" {
        case "\n\t":
            if (((!true) && (!(42 != 193))) && v2) {
                var v31 uint8 = 0
                v32 := GenS2{F0: (v23 << ((65535) & 63)), F1: ((v2 && v2) || (343 >= v11)), F2: 1}
                _ = v31
                _ = v32
            }
        }
        _ = fi24
        _ = v30
    }
    var v33 uint = 457
    switch (v1 + "") {
    case "\n\t":
        if (((255 > 482) && (!false)) && (0 < 42)) {
            for fi34 := 0; fi34 < 3; fi34++ {
                v1 = ""
                _ = fi34
            }
        } else {
            fmt.Println(v23, v1, v1)
        }
    case "":
        var v35 uint32 = 2147483647
        _ = v35
    default:
        if (4294967295 != -1) {
            if true {
                v22 -= (-(v22))
                for fi36 := 0; fi36 < 13; fi36++ {
                    for fi37 := 0; fi37 < 12; fi37++ {
                        if true {
                            v38 := GenS3{F0: fi37, F1: (^(v22))}
                            _ = v38
                        } else {
                            if true {
                                var v39 uint = 8
                                v2 = (!(v22 < (v22 % (((^(v22))) | 1))))
                                _ = v39
                            }
                            if true {
                                switch (fi36 << ((0) & 63)) {
                                case 9223372036854775807:
                                    v40 := []int{fi36}
                                    v41 := GenS4{F0: 624, F1: -9223372036854775808}
                                    switch v1 {
                                    case "\n\t":
                                        v40 = append(v40, 9223372036854775807)
                                    default:
                                        v1 = v1
                                    }
                                    _ = v40
                                    _ = v41
                                case -1:
                                    v42 := GenS5{F0: -989, F1: v22}
                                    _ = v42
                                default:
                                    var v43 bool = v2
                                    _ = v43
                                }
                                v33 = 281
                            }
                            v44 := GenS6{F0: 42, F1: v8}
                            v45 := []int{}
                            if len(v45) > 0 {
                                v45[0] = fi37
                            }
                            _ = v44
                            _ = v45
                        }
                        var v46 uint = (^(v33))
                        fmt.Println(fi37)
                        _ = fi37
                        _ = v46
                    }
                    var v47 int64 = (v22 / (((v22 * v22)) | 1))
                    _ = fi36
                    _ = v47
                }
                switch 336 {
                case 4294967295:
                    for fi48 := 0; fi48 < 12; fi48++ {
                        v49 := GenS6{F0: 4294967295, F1: v8}
                        v23 = v23
                        _ = fi48
                        _ = v49
                    }
                case 9223372036854775807:
                    var v50 int32 = v23
                    _ = v50
                case 868:
                    fmt.Println(v11)
                default:
                    v23 += v23
                }
            }
            switch ("\n\t" + v1) {
            case "a":
                fmt.Println(v2)
            case "\n\t":
                v51 := GenS3{F0: 2147483647, F1: -1}
                _ = v51
            }
            v52 := GenS2{F0: (-(v23)), F1: (2 == 255), F2: (-(v23))}
            _ = v52
        } else {
            v11 = (v11 ^ (^(v11)))
            switch "hello" {
            case "":
                v53 := GenS7{F0: 194, F1: (v11 - 406), F2: 65535}
                _ = v53
            case "a":
                switch "" {
                case "a":
                    var v54 float64 = float64((0.1 + 1e10))
                    _ = v54
                case "hello":
                    var v55 float64 = 0.1
                    _ = v55
                }
            default:
                switch (v1 + v1) {
                case "hello":
                    var v56 string = "a"
                    _ = v56
                case "a":
                    for fi57 := 0; fi57 < 1; fi57++ {
                        if v2 {
                            v8 -= 1
                            switch "" {
                            case "a":
                                var v58 int8 = 0
                                _ = v58
                            case "hello":
                                if false {
                                    v33 -= 271
                                } else {
                                    for fi59 := 0; fi59 < 13; fi59++ {
                                        var v60 uint = 1
                                        var v61 int8 = 65
                                        _ = fi59
                                        _ = v60
                                        _ = v61
                                    }
                                    v3[4294967295] = (v1 == "hello")
                                    v62 := GenS4{F0: uint16(v11), F1: -1}
                                    _ = v62
                                }
                            case "foo bar":
                                if ((v11 / ((4294967295) | 1)) < (^(v11))) {
                                    var v63 uint = v33
                                    v8 = 0
                                    _ = v63
                                } else {
                                    if true {
                                        for fi64 := 0; fi64 < 16; fi64++ {
                                            if true {
                                                var v65 int = (fi64 & (fi57 | -9223372036854775808))
                                                _ = v65
                                            }
                                            switch (v1 + "hello") {
                                            case "hello":
                                                if (255 <= 42) {
                                                    if ((v23 & (v23 << (((^(v33))) & 63))) > v23) {
                                                        v66 := GenS7{F0: uint8(v33), F1: v11, F2: 0}
                                                        if (v23 > v23) {
                                                            var v67 float64 = 0.1
                                                            _ = v67
                                                        } else {
                                                            var v68 string = v1
                                                            switch v1 {
                                                            case "\n\t":
                                                                v2 = false
                                                            case "hello":
                                                                v69 := GenS2{F0: 2147483647, F1: (-7 >= 42), F2: -2147483648}
                                                                _ = v69
                                                            case "":
                                                                var v70 int = fi64
                                                                _ = v70
                                                            default:
                                                                v23 += int32(v11)
                                                            }
                                                            v8 -= v8
                                                            _ = v68
                                                        }
                                                        if false {
                                                            fmt.Println(v1, fi57, v1)
                                                            v71 := []int{(fi64 + fi64)}
                                                            _ = v71
                                                        } else {
                                                            var v72 uint8 = 204
                                                            fmt.Println(v11)
                                                            _ = v72
                                                        }
                                                        var v73 int32 = v23
                                                        fmt.Println(v2, v8, fi57)
                                                        _ = v66
                                                        _ = v73
                                                    } else {
                                                        var v74 string = "\n\t"
                                                        if ((v33 - 4294967295) <= v33) {
                                                            switch (fi57 / ((fi64) | 1)) {
                                                            case 4294967295:
                                                                switch (v74 + "foo bar") {
                                                                case "hello":
                                                                    if false {
                                                                        if v2 {
                                                                            v75 := v3[1]
                                                                            _ = v75
                                                                        }
                                                                    }
                                                                default:
                                                                    var v76 int16 = int16(v11)
                                                                    _ = v76
                                                                }
                                                            case 127:
                                                                var v77 float64 = 0.1
                                                                _ = v77
                                                            default:
                                                                if (((v11 ^ v11) <= 42) || false) {
                                                                    v33 += (v33 << ((v33) & 63))
                                                                    v22 = (v22 | int64(v23))
                                                                    v78 := []int{fi64, (fi57 >> ((65535) & 63)), fi64, (fi57 >> (((v33 >> ((v33) & 63))) & 63))}
                                                                    _ = v78
                                                                }
                                                            }
                                                        }
                                                        _ = v74
                                                    }
                                                    var v79 uint8 = 218
                                                    if (!v2) {
                                                        v3[506] = false
                                                        fmt.Println(v33, fi64)
                                                        v79 -= uint8(v11)
                                                    }
                                                    v2 = v2
                                                    _ = v79
                                                } else {
                                                    delete(v3, 420)
                                                    v80 := GenS5{F0: 4294967295, F1: int64(v11)}
                                                    switch (v1 + v1) {
                                                    case "\n\t":
                                                        if (((v8 << (((v33 - (v33 / (((v33 ^ v33)) | 1)))) & 63)) != v8) || (false || (true || (!(-533 <= 0))))) {
                                                            var v81 int8 = -128
                                                            var v82 uint = (v33 << ((18446744073709551615) & 63))
                                                            v22 += (v22 * (v22 | (v22 % (((v22 & v22)) | 1))))
                                                            v81 -= v81
                                                            _ = v81
                                                            _ = v82
                                                        }
                                                    case "hello":
                                                        v83 := []int{fi64, (fi64 & int(v8)), (fi57 / (((fi57 << ((v33) & 63))) | 1)), int(v23)}
                                                        _ = v83
                                                    case "foo bar":
                                                        v84 := GenS6{F0: 42, F1: (v8 / (((v8 >> ((v33) & 63))) | 1))}
                                                        _ = v84
                                                    default:
                                                        fmt.Println(v2)
                                                    }
                                                    fmt.Println(v33, v1, v33)
                                                    delete(v3, 453)
                                                    _ = v80
                                                }
                                            }
                                            _ = fi64
                                        }
                                        for fi85 := 0; fi85 < 3; fi85++ {
                                            v11 -= 42
                                            switch (^(fi57)) {
                                            case 716:
                                                v86 := []int{fi85, fi85}
                                                _ = v86
                                            case 1:
                                                delete(v3, 4294967295)
                                            case 9223372036854775807:
                                                switch (v1 + ("hello" + "a")) {
                                                case "\n\t":
                                                    v87 := GenS3{F0: fi57, F1: (-(v22))}
                                                    _ = v87
                                                case "hello":
                                                    v8 -= v8
                                                case "foo bar":
                                                    var v88 int8 = 1
                                                    _ = v88
                                                default:
                                                    fmt.Println(v8, v1, fi57)
                                                }
                                            }
                                            _ = fi85
                                        }
                                        delete(v3, 399)
                                    }
                                }
                            default:
                                var v89 int64 = v22
                                _ = v89
                            }
                        } else {
                            var v90 int = fi57
                            v2 = v2
                            _ = v90
                        }
                        v11 = v11
                        _ = fi57
                    }
                default:
                    if (true && v2) {
                        for fi91 := 0; fi91 < 4; fi91++ {
                            _ = len(v3)
                            v2 = v2
                            _ = fi91
                        }
                        if true {
                            fmt.Println(v33, v22, v33)
                            if (!(572 < 1)) {
                                v2 = v2
                                v3[2147483648] = (0 <= 116)
                            } else {
                                v11 += (^(v11))
                                var v92 float64 = 0.1
                                fmt.Println(v8)
                                _ = v92
                            }
                            var v93 float64 = -2.25
                            _ = v93
                        }
                        var v94 string = v1
                        _ = v94
                    } else {
                        for fi95 := 0; fi95 < 14; fi95++ {
                            fmt.Println(v8, v33, v1)
                            if (v22 >= int64(v11)) {
                                v96 := []int{(fi95 | fi95), (fi95 & -1)}
                                _ = v96
                            }
                            _ = fi95
                        }
                        if (2 > 2) {
                            if false {
                                switch ((v1 + (v1 + (v1 + (v1 + "")))) + ("\n\t" + ("" + (v1 + v1)))) {
                                case "a":
                                    for fi97 := 0; fi97 < 19; fi97++ {
                                        v98 := GenS3{F0: fi97, F1: int64(v11)}
                                        v2 = ((v1 == v1) || ((v11 + (^(v11))) == (v11 / ((4294967295) | 1))))
                                        _ = fi97
                                        _ = v98
                                    }
                                }
                                switch "\n\t" {
                                case "":
                                    for fi99 := 0; fi99 < 1; fi99++ {
                                        for fi100 := 0; fi100 < 4; fi100++ {
                                            v101 := []int{0, fi99, fi99, fi99}
                                            _ = fi100
                                            _ = v101
                                        }
                                        _ = fi99
                                    }
                                case "foo bar":
                                    switch 128 {
                                    case 580:
                                        v2 = (127 < 239)
                                    case 2147483647:
                                        v102 := []int{183, 339}
                                        _ = v102
                                    case 0:
                                        var v103 uint = v33
                                        _ = v103
                                    default:
                                        if v2 {
                                            switch v1 {
                                            case "foo bar":
                                                switch -2147483648 {
                                                case 128:
                                                    if true {
                                                        var v104 uint64 = 4294967295
                                                        switch ("" + (v1 + "foo bar")) {
                                                        case "\n\t":
                                                            v8 += 196
                                                        case "foo bar":
                                                            var v105 bool = v2
                                                            _ = v105
                                                        case "":
                                                            v8 = v8
                                                        default:
                                                            var v106 uint = (v33 + v33)
                                                            _ = v106
                                                        }
                                                        _ = v104
                                                    }
                                                case 187:
                                                    var v107 float64 = float64(1.5)
                                                    _ = v107
                                                default:
                                                    switch v1 {
                                                    case "a":
                                                        for fi108 := 0; fi108 < 19; fi108++ {
                                                            var v109 uint64 = (v11 * (v11 & v11))
                                                            _ = fi108
                                                            _ = v109
                                                        }
                                                    default:
                                                        v2 = v2
                                                    }
                                                }
                                            case "":
                                                for fi110 := 0; fi110 < 11; fi110++ {
                                                    var v111 uint32 = 4294967295
                                                    fmt.Println(v33, v33)
                                                    _ = fi110
                                                    _ = v111
                                                }
                                            }
                                            v11 += (v11 - v11)
                                            v112 := []int{2147483647, -370, -9223372036854775808}
                                            _ = v112
                                        } else {
                                            var v113 int64 = (v22 % ((v22) | 1))
                                            v113 = (v22 >> ((0) & 63))
                                            _ = v113
                                        }
                                    }
                                case "hello":
                                    v114 := []int{635, int(v33), 127}
                                    _ = v114
                                }
                                var v115 string = (((v1 + v1) + v1) + ((("a" + ((v1 + "") + (("a" + v1) + v1))) + ((((("foo bar" + "") + ((("foo bar" + (v1 + v1)) + (v1 + v1)) + v1)) + ("foo bar" + v1)) + v1) + v1)) + (v1 + "a")))
                                _ = v115
                            }
                            for fi116 := 0; fi116 < 16; fi116++ {
                                for fi117 := 0; fi117 < 11; fi117++ {
                                    v118 := GenS8{F0: ((!false) || v2), F1: 42}
                                    fmt.Println(v23, v23, v33)
                                    v119 := GenS4{F0: 1, F1: v22}
                                    _ = fi117
                                    _ = v118
                                    _ = v119
                                }
                                _ = fi116
                            }
                            v120 := GenS9{F0: v8}
                            _ = v120
                        }
                    }
                }
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v33
    _ = v8
}
