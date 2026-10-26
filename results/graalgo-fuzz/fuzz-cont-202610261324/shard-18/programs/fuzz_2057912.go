package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32; F2 int8 }
type GenS2 struct { F0 int16; F1 string }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 int8; F1 int64 }
type GenS8 struct { F0 uint16; F1 uint16; F2 int32 }
type GenS9 struct { F0 int8; F1 int; F2 uint16 }
type GenS10 struct { F0 uint16; F1 int8; F2 uint64 }
type GenS11 struct { F0 uint32; F1 uint32; F2 uint8 }
type GenS12 struct { F0 uint16 }
type GenS13 struct { F0 int; F1 int8; F2 int16 }
type GenS14 struct { F0 string; F1 uint16; F2 int64 }
type GenS15 struct { F0 int64; F1 bool }
type GenS16 struct { F0 int64; F1 int8 }
type GenS17 struct { F0 string; F1 int32 }
type GenS18 struct { F0 int8; F1 uint64 }
type GenS19 struct { F0 uint; F1 int64; F2 int8 }
type GenS20 struct { F0 uint32; F1 int; F2 string }
type GenS21 struct { F0 uint16; F1 int8; F2 uint8 }
type GenS22 struct { F0 uint16; F1 int32 }
type GenS23 struct { F0 int }
type GenS24 struct { F0 uint; F1 int16; F2 uint16 }
type GenS25 struct { F0 uint; F1 int64; F2 int8 }
type GenS26 struct { F0 int16; F1 int16; F2 bool }
type GenS27 struct { F0 int32; F1 uint64 }
type GenS28 struct { F0 uint64 }
type GenS29 struct { F0 uint; F1 uint }
type GenS30 struct { F0 int64; F1 uint8; F2 uint }
type GenS31 struct { F0 uint8 }
type GenS32 struct { F0 int; F1 uint64; F2 uint16 }
type GenS33 struct { F0 uint8; F1 uint64; F2 int64 }
type GenS34 struct { F0 uint64; F1 string; F2 uint64 }
type GenS35 struct { F0 uint32; F1 uint8; F2 string }
type GenS36 struct { F0 uint32; F1 uint64 }
type GenS37 struct { F0 uint32 }
type GenS38 struct { F0 uint; F1 int }
type GenS39 struct { F0 uint16 }
type GenS40 struct { F0 int16; F1 int16 }
type GenS41 struct { F0 int16; F1 int64 }
type GenS42 struct { F0 int; F1 uint32 }
type GenS43 struct { F0 bool }
type GenS44 struct { F0 int32; F1 bool }
type GenS45 struct { F0 uint32; F1 int64 }
type GenS46 struct { F0 uint16; F1 bool; F2 uint32 }
type GenS47 struct { F0 uint8 }
type GenS48 struct { F0 uint }

func main() {
    v1 := []uint32{2147483648, 2147483647, 42, 2}
    var v2 int16 = 1000
    v3 := make(map[uint32]uint32)
    var v4 float32 = 1.5
    if len(v1) > 0 {
        v1[0] = 2
    }
    for fi5 := 0; fi5 < 15; fi5++ {
        v3[4294967295] = 42
        v6 := GenS1{F0: (v2 * -32768), F1: 856, F2: 0}
        var v7 int8 = 100
        _ = fi5
        _ = v6
        _ = v7
    }
    fmt.Println(v2, v4)
    for fi8 := 0; fi8 < 19; fi8++ {
        switch (("a" + "hello") + ((("\n\t" + ("hello" + ("a" + "\n\t"))) + "a") + "")) {
        case "hello":
            if (((((1e10 / float64(v4)) + 1.5) * (1.5 - (1.5 / -2.25))) * -2.25) <= 1e10) {
                if (!true) {
                    if true {
                        switch ("foo bar" + "") {
                        case "\n\t":
                            delete(v3, 4294967295)
                        case "":
                            for fi9 := 0; fi9 < 2; fi9++ {
                                if true {
                                    var v10 int8 = 100
                                    _ = v10
                                } else {
                                    v11 := GenS1{F0: v2, F1: 4294967295, F2: -71}
                                    switch (^(fi8)) {
                                    case 359:
                                        if (-2147483648 <= -160) {
                                            if (v2 != v2) {
                                                var v12 uint = 4294967295
                                                _ = v12
                                            }
                                            v13 := GenS1{F0: (v2 / ((v2) | 1)), F1: 1, F2: -1}
                                            _ = v13
                                        } else {
                                            if (2 >= 0) {
                                                var v14 uint32 = 65535
                                                var v15 int8 = 1
                                                var v16 float64 = 0.0
                                                v16 -= (0.1 - (((v16 / ((((v16 + v16) + 0.0) - v16) + ((v16 - -2.25) / (1.5 - float64(v4))))) / (v16 - v16)) / ((1.5 * v16) + (float64((((3.14159 / v4) * (0.0 / v4)) / 0.1)) - float64(((1.5 / -2.25) - v4))))))
                                                _ = v14
                                                _ = v15
                                                _ = v16
                                            } else {
                                                v2 += v2
                                            }
                                            v17 := v3[756]
                                            if (4294967295 > 42) {
                                                if (true || (true || (false || true))) {
                                                    v17 += v17
                                                }
                                                v18 := GenS1{F0: 42, F1: 475, F2: 125}
                                                _ = v18
                                            }
                                            _ = v17
                                        }
                                    case -2147483648:
                                        var v19 float32 = 3.14159
                                        _ = v19
                                    case 810:
                                        v2 = v2
                                    }
                                    _ = v11
                                }
                                v4 -= v4
                                v20 := GenS1{F0: 1, F1: 65535, F2: 1}
                                _ = fi9
                                _ = v20
                            }
                        case "hello":
                            v4 = (float32(float64(0.1)) + 3.14159)
                        default:
                            if true {
                                var v21 float32 = float32(3.14159)
                                _ = v21
                            }
                        }
                    }
                } else {
                    v4 = v4
                }
                v4 = float32((0.0 - -2.25))
                var v22 bool = (false || (122 < -88))
                _ = v22
            }
        }
        _ = fi8
    }
    for fi23 := 0; fi23 < 15; fi23++ {
        v24 := GenS1{F0: -1, F1: 2147483647, F2: 100}
        switch 32768 {
        case -20:
            switch fi23 {
            case 127:
                if true {
                    if false {
                        _ = len(v3)
                        for fi25 := 0; fi25 < 20; fi25++ {
                            switch -2147483648 {
                            case 242:
                                var v26 uint64 = 553
                                _ = v26
                            case -2147483648:
                                var v27 uint64 = 25
                                _ = v27
                            }
                            var v28 int = (fi23 * -9223372036854775808)
                            _ = fi25
                            _ = v28
                        }
                    } else {
                        var v29 int = 4294967295
                        _ = v29
                    }
                    v1 = append(v1, 878)
                    v2 = v2
                }
            case 1:
                var v30 uint16 = 911
                _ = v30
            default:
                v31 := GenS1{F0: v2, F1: 65535, F2: 1}
                _ = v31
            }
        default:
            if true {
                for fi32 := 0; fi32 < 5; fi32++ {
                    var v33 float32 = v4
                    if false {
                        var v34 uint = 18446744073709551615
                        var v35 uint = v34
                        _ = v34
                        _ = v35
                    }
                    fmt.Println(fi23, v4)
                    _ = fi32
                    _ = v33
                }
                if true {
                    v4 += 0.1
                    if true {
                        v2 += (v2 | v2)
                    } else {
                        v36 := v3[65535]
                        _ = v36
                    }
                    for fi37 := 0; fi37 < 10; fi37++ {
                        switch 127 {
                        case -9223372036854775808:
                            if false {
                                if (!(!true)) {
                                    v38 := GenS2{F0: v2, F1: "foo bar"}
                                    switch "foo bar" {
                                    case "":
                                        var v39 uint32 = 1
                                        _ = v39
                                    case "hello":
                                        var v40 uint16 = 0
                                        _ = v40
                                    case "a":
                                        switch (-(fi23)) {
                                        case -9223372036854775808:
                                            var v41 uint16 = 152
                                            _ = v41
                                        case -335:
                                            _ = len(v3)
                                        }
                                    default:
                                        switch (-(fi37)) {
                                        case -874:
                                            fmt.Println(fi37)
                                        default:
                                            switch "hello" {
                                            case "a":
                                                v42 := GenS3{F0: 127}
                                                _ = v42
                                            }
                                        }
                                    }
                                    _ = v38
                                } else {
                                    v43 := GenS3{F0: 12}
                                    _ = v43
                                }
                                if (!false) {
                                    v44 := GenS3{F0: 255}
                                    v4 += -2.25
                                    var v45 uint8 = 123
                                    _ = v44
                                    _ = v45
                                }
                                switch "\n\t" {
                                case "hello":
                                    var v46 uint32 = 2
                                    _ = v46
                                case "a":
                                    switch -264 {
                                    case -862:
                                        fmt.Println(fi37)
                                    }
                                case "foo bar":
                                    if false {
                                        fmt.Println(fi23, fi37, fi37)
                                        if true {
                                            v2 += v2
                                        }
                                        v3[748] = 265
                                        v2 += -88
                                        switch "" {
                                        case "hello":
                                            switch fi37 {
                                            case 1:
                                                switch fi23 {
                                                case -9223372036854775808:
                                                    if false {
                                                        if (!false) {
                                                            fmt.Println(v2, v4, v2)
                                                            v47 := GenS3{F0: 42}
                                                            var v48 int = (fi37 | fi37)
                                                            _ = v47
                                                            _ = v48
                                                        } else {
                                                            fmt.Println(fi23)
                                                            v49 := GenS3{F0: 1}
                                                            _ = v49
                                                        }
                                                        switch "foo bar" {
                                                        case "hello":
                                                            if (77 >= 148) {
                                                                fmt.Println(fi37, v2, v4)
                                                                fmt.Println(v2, fi23)
                                                                v1 = append(v1, 2147483648)
                                                            }
                                                        case "foo bar":
                                                            v50 := GenS2{F0: v2, F1: "hello"}
                                                            _ = v50
                                                        case "":
                                                            v1 = append(v1, 2)
                                                        }
                                                    }
                                                case 32768:
                                                    if false {
                                                        fmt.Println(fi23, fi37, v2)
                                                        if (false && (-1 < 42)) {
                                                            delete(v3, 0)
                                                            _ = len(v3)
                                                            v4 += v4
                                                        } else {
                                                            var v51 int32 = 1
                                                            fmt.Println(fi37, v2, fi23)
                                                            if (fi23 == fi37) {
                                                                var v52 int64 = 0
                                                                fmt.Println(v51, v51, v52)
                                                                _ = v52
                                                            }
                                                            _ = v51
                                                        }
                                                        fmt.Println(fi23, fi23, fi37)
                                                    }
                                                default:
                                                    fmt.Println(v2, fi37, v4)
                                                }
                                            case 32767:
                                                var v53 int = 1
                                                if true {
                                                    var v54 uint8 = 2
                                                    v4 += 1e10
                                                    switch (fi37 >> ((698) & 63)) {
                                                    case -79:
                                                        var v55 float64 = 0.1
                                                        _ = v55
                                                    default:
                                                        switch fi37 {
                                                        case 42:
                                                            if false {
                                                                if (!((("\n\t" + ("\n\t" + ("a" + "hello"))) + ("foo bar" + "hello")) <= "")) {
                                                                    if true {
                                                                        _ = len(v3)
                                                                        if false {
                                                                            v54 += v54
                                                                            fmt.Println(fi37)
                                                                        } else {
                                                                            var v56 float32 = (v4 - float32(float64(-2.25)))
                                                                            if ((!(!(32768 != (fi23 << ((0) & 63))))) || false) {
                                                                                switch "" {
                                                                                case "\n\t":
                                                                                    switch 759 {
                                                                                    case -1:
                                                                                        var v57 bool = false
                                                                                        _ = v57
                                                                                    case 4294967295:
                                                                                        v58 := GenS4{F0: 42}
                                                                                        _ = v58
                                                                                    case 919:
                                                                                        var v59 int8 = 100
                                                                                        _ = v59
                                                                                    }
                                                                                    fmt.Println(v56, v2, v54)
                                                                                case "":
                                                                                    switch (v53 >> ((42) & 63)) {
                                                                                    case 9223372036854775807:
                                                                                        v1 = append(v1, 65535)
                                                                                    default:
                                                                                        v4 -= v56
                                                                                    }
                                                                                case "hello":
                                                                                    fmt.Println(fi23)
                                                                                }
                                                                                if (0 < 100) {
                                                                                    var v60 int8 = -1
                                                                                    if true {
                                                                                        fmt.Println(v4, fi37, fi23)
                                                                                        fmt.Println(fi37)
                                                                                    } else {
                                                                                        v61 := GenS4{F0: 1}
                                                                                        _ = v61
                                                                                    }
                                                                                    _ = v60
                                                                                }
                                                                            }
                                                                            v4 = v4
                                                                            _ = v56
                                                                        }
                                                                    }
                                                                } else {
                                                                    if true {
                                                                        v53 += fi37
                                                                        if (107 > 42) {
                                                                            var v62 uint64 = 2
                                                                            var v63 float64 = 1e10
                                                                            var v64 uint8 = (v54 >> ((255) & 63))
                                                                            _ = v62
                                                                            _ = v63
                                                                            _ = v64
                                                                        } else {
                                                                            var v65 string = "foo bar"
                                                                            fmt.Println(v53, v65, v65)
                                                                            v66 := GenS2{F0: 0, F1: v65}
                                                                            _ = v65
                                                                            _ = v66
                                                                        }
                                                                    }
                                                                    fmt.Println(v4, v53, v2)
                                                                    if true {
                                                                        var v67 float32 = 1e10
                                                                        _ = v67
                                                                    } else {
                                                                        v68 := GenS1{F0: v2, F1: 2, F2: 86}
                                                                        _ = v68
                                                                    }
                                                                    switch (fi37 >> ((0) & 63)) {
                                                                    case -837:
                                                                        switch v53 {
                                                                        case -49:
                                                                            var v69 int = v53
                                                                            _ = v69
                                                                        case -250:
                                                                            if ((-(v53)) == fi23) {
                                                                                switch "\n\t" {
                                                                                case "\n\t":
                                                                                    if (9223372036854775807 == 1) {
                                                                                        v4 += float32((0.0 - 1e10))
                                                                                    } else {
                                                                                        if true {
                                                                                            var v70 uint = 65535
                                                                                            v3[65535] = 2147483648
                                                                                            _ = v70
                                                                                        } else {
                                                                                            fmt.Println(v53, v53)
                                                                                            v71 := GenS1{F0: (v2 * (v2 >> ((1) & 63))), F1: 42, F2: 127}
                                                                                            v72 := GenS5{F0: true}
                                                                                            v73 := GenS2{F0: (v2 | (v2 >> ((uint(v54)) & 63))), F1: ("a" + ("" + (("" + (("" + ("\n\t" + "")) + ("hello" + ((("" + ("a" + "hello")) + ("\n\t" + (("" + "\n\t") + "a"))) + ("foo bar" + "\n\t"))))) + ("\n\t" + ("foo bar" + "")))))}
                                                                                            var v74 uint32 = 2147483647
                                                                                            _ = v71
                                                                                            _ = v72
                                                                                            _ = v73
                                                                                            _ = v74
                                                                                        }
                                                                                        if true {
                                                                                            var v75 uint32 = 0
                                                                                            fmt.Println(v75, v2, v4)
                                                                                            var v76 int8 = 1
                                                                                            _ = v75
                                                                                            _ = v76
                                                                                        } else {
                                                                                            var v77 uint = 669
                                                                                            v2 = (v2 & v2)
                                                                                            if false {
                                                                                                v78 := GenS5{F0: (-67 > 9223372036854775807)}
                                                                                                if true {
                                                                                                    if true {
                                                                                                        v79 := GenS4{F0: 42}
                                                                                                        _ = v79
                                                                                                    } else {
                                                                                                        fmt.Println(v77, fi23, fi23)
                                                                                                        switch fi37 {
                                                                                                        case 127:
                                                                                                            v2 -= v2
                                                                                                        case -604:
                                                                                                            _ = len(v1)
                                                                                                        case 128:
                                                                                                            v4 = float32(float64((-2.25 - v4)))
                                                                                                        }
                                                                                                    }
                                                                                                } else {
                                                                                                    fmt.Println(fi37)
                                                                                                    if true {
                                                                                                        switch "\n\t" {
                                                                                                        case "hello":
                                                                                                            var v80 uint8 = 49
                                                                                                            _ = v80
                                                                                                        }
                                                                                                    }
                                                                                                    switch fi23 {
                                                                                                    case 127:
                                                                                                        v2 = 1
                                                                                                    case 261:
                                                                                                        if (255 == 255) {
                                                                                                            v81 := GenS3{F0: (^(v54))}
                                                                                                            _ = v81
                                                                                                        } else {
                                                                                                            delete(v3, 816)
                                                                                                            v53 += (fi37 % ((1) | 1))
                                                                                                        }
                                                                                                    default:
                                                                                                        var v82 int16 = v2
                                                                                                        _ = v82
                                                                                                    }
                                                                                                }
                                                                                                _ = v78
                                                                                            }
                                                                                            _ = v77
                                                                                        }
                                                                                    }
                                                                                    var v83 uint64 = 1
                                                                                    _ = v83
                                                                                case "foo bar":
                                                                                    v84 := v3[0]
                                                                                    _ = v84
                                                                                default:
                                                                                    if true {
                                                                                        v54 += (^(v54))
                                                                                        v54 += v54
                                                                                        switch (fi23 | fi23) {
                                                                                        case 4294967295:
                                                                                            v85 := GenS5{F0: false}
                                                                                            var v86 int8 = 127
                                                                                            var v87 uint32 = 0
                                                                                            delete(v3, 0)
                                                                                            v87 -= 2147483648
                                                                                            _ = v85
                                                                                            _ = v86
                                                                                            _ = v87
                                                                                        case 680:
                                                                                            v88 := GenS6{F0: -12}
                                                                                            _ = v88
                                                                                        default:
                                                                                            if (0 > 0) {
                                                                                                v54 = (v54 - 42)
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                }
                                                                                var v89 bool = ((((50 < 0) || (-635 != 2147483647)) && (3.14159 > 1e10)) || false)
                                                                                _ = v89
                                                                            } else {
                                                                                if len(v1) > 0 {
                                                                                    v1[0] = 548
                                                                                }
                                                                                if ((1.5 * (float64(v4) * float64(v4))) <= 0.1) {
                                                                                    if true {
                                                                                        v3[65535] = 2147483648
                                                                                        v4 = v4
                                                                                    } else {
                                                                                        _ = len(v1)
                                                                                        if true {
                                                                                            v2 += (v2 + (v2 * v2))
                                                                                            var v90 int64 = -891
                                                                                            v1 = append(v1, 626)
                                                                                            _ = len(v3)
                                                                                            _ = v90
                                                                                        }
                                                                                        fmt.Println(v4, v54)
                                                                                        var v91 int8 = 117
                                                                                        v2 -= (v2 ^ (v2 + (-(v2))))
                                                                                        _ = v91
                                                                                    }
                                                                                    var v92 int = v53
                                                                                    _ = v92
                                                                                } else {
                                                                                    var v93 uint16 = 32767
                                                                                    var v94 uint64 = 4294967295
                                                                                    _ = v93
                                                                                    _ = v94
                                                                                }
                                                                            }
                                                                        case 570:
                                                                            if len(v1) > 0 {
                                                                                v1[0] = 527
                                                                            }
                                                                        default:
                                                                            v95 := GenS3{F0: v54}
                                                                            _ = v95
                                                                        }
                                                                    default:
                                                                        if (!(!((-(v2)) >= (v2 | v2)))) {
                                                                            var v96 float64 = -2.25
                                                                            v97 := GenS7{F0: 0, F1: 1}
                                                                            _ = v96
                                                                            _ = v97
                                                                        }
                                                                    }
                                                                    switch "hello" {
                                                                    case "hello":
                                                                        v98 := v3[2147483647]
                                                                        _ = v98
                                                                    }
                                                                    v54 = v54
                                                                }
                                                            } else {
                                                                var v99 uint8 = 38
                                                                if ((0 < 42) && (9223372036854775807 < 38)) {
                                                                    var v100 int8 = 100
                                                                    v101 := GenS7{F0: (v100 >> ((4294967295) & 63)), F1: 0}
                                                                    switch "\n\t" {
                                                                    case "hello":
                                                                        var v102 int64 = -9223372036854775808
                                                                        _ = v102
                                                                    case "a":
                                                                        fmt.Println(v99, v53)
                                                                        delete(v3, 2147483648)
                                                                        if (!(162 < 65535)) {
                                                                            var v103 int16 = (v2 % ((-51) | 1))
                                                                            _ = v103
                                                                        }
                                                                        v104 := GenS4{F0: 255}
                                                                        v2 = 42
                                                                        _ = v104
                                                                    case "foo bar":
                                                                        fmt.Println(fi37, fi23, v2)
                                                                        var v105 int8 = v100
                                                                        var v106 uint32 = 2147483647
                                                                        _ = len(v3)
                                                                        var v107 float32 = float32(0.1)
                                                                        _ = v105
                                                                        _ = v106
                                                                        _ = v107
                                                                    default:
                                                                        switch ("a" + ("a" + "foo bar")) {
                                                                        case "":
                                                                            if true {
                                                                                var v108 bool = true
                                                                                if v108 {
                                                                                    fmt.Println(v54, v108, v4)
                                                                                    switch (("" + ("foo bar" + "")) + "a") {
                                                                                    case "\n\t":
                                                                                        var v109 uint32 = 65535
                                                                                        _ = v109
                                                                                    case "":
                                                                                        v100 = v100
                                                                                    case "foo bar":
                                                                                        if len(v1) >= 2 {
                                                                                            v1 = v1[1:len(v1)]
                                                                                        }
                                                                                        var v110 uint32 = 2
                                                                                        _ = v110
                                                                                    default:
                                                                                        v3[2147483647] = 838
                                                                                    }
                                                                                    if false {
                                                                                        var v111 int = (^(fi37))
                                                                                        var v112 int8 = v100
                                                                                        switch "\n\t" {
                                                                                        case "\n\t":
                                                                                            switch ("" + "foo bar") {
                                                                                            case "":
                                                                                                switch "hello" {
                                                                                                case "\n\t":
                                                                                                    if v108 {
                                                                                                        switch "" {
                                                                                                        case "\n\t":
                                                                                                            switch (fi23 % ((600) | 1)) {
                                                                                                            case 128:
                                                                                                                var v113 int = (fi23 | fi23)
                                                                                                                _ = v113
                                                                                                            case 4294967295:
                                                                                                                v114 := GenS3{F0: 2}
                                                                                                                _ = v114
                                                                                                            case 32768:
                                                                                                                v100 -= 100
                                                                                                            }
                                                                                                        case "foo bar":
                                                                                                            var v115 uint64 = 2
                                                                                                            _ = v115
                                                                                                        default:
                                                                                                            if true {
                                                                                                                var v116 uint32 = 2147483648
                                                                                                                var v117 float32 = v4
                                                                                                                _ = v116
                                                                                                                _ = v117
                                                                                                            }
                                                                                                            _ = len(v1)
                                                                                                            _ = len(v3)
                                                                                                            delete(v3, 65535)
                                                                                                        }
                                                                                                        if (((v2 > v2) || false) && false) {
                                                                                                            if v108 {
                                                                                                                v1 = append(v1, 65535)
                                                                                                                fmt.Println(v53)
                                                                                                                var v118 uint64 = 4294967295
                                                                                                                _ = v118
                                                                                                            } else {
                                                                                                                var v119 uint64 = 51
                                                                                                                _ = v119
                                                                                                            }
                                                                                                            switch "foo bar" {
                                                                                                            case "foo bar":
                                                                                                                v111 += (fi37 - 42)
                                                                                                                v53 = fi37
                                                                                                                v4 = 1e10
                                                                                                            case "":
                                                                                                                if (v4 == v4) {
                                                                                                                    switch 1 {
                                                                                                                    case 9223372036854775807:
                                                                                                                        v120 := GenS6{F0: v100}
                                                                                                                        _ = v120
                                                                                                                    case 4294967295:
                                                                                                                        v100 -= v100
                                                                                                                    case -453:
                                                                                                                        v121 := GenS4{F0: 0}
                                                                                                                        _ = v121
                                                                                                                    }
                                                                                                                    fmt.Println(v54, v100, v53)
                                                                                                                    v112 = (v112 << ((65535) & 63))
                                                                                                                } else {
                                                                                                                    var v122 float64 = (float64(-2.25) / 1.0)
                                                                                                                    _ = v122
                                                                                                                }
                                                                                                            }
                                                                                                        } else {
                                                                                                            fmt.Println(v2, v4)
                                                                                                            if (v108 && v108) {
                                                                                                                var v123 uint32 = 1
                                                                                                                var v124 int64 = 0
                                                                                                                if v108 {
                                                                                                                    fmt.Println(v99, v53)
                                                                                                                    v1 = append(v1, 280)
                                                                                                                    fmt.Println(v124)
                                                                                                                    if (-880 > -2147483648) {
                                                                                                                        switch "hello" {
                                                                                                                        case "hello":
                                                                                                                            if true {
                                                                                                                                if v108 {
                                                                                                                                    v124 -= v124
                                                                                                                                    var v125 int8 = (v112 & v112)
                                                                                                                                    fmt.Println(v111, v124)
                                                                                                                                    _ = v125
                                                                                                                                }
                                                                                                                                var v126 int16 = 42
                                                                                                                                _ = v126
                                                                                                                            }
                                                                                                                        case "\n\t":
                                                                                                                            var v127 int8 = v112
                                                                                                                            _ = v127
                                                                                                                        case "":
                                                                                                                            if (1 <= 2) {
                                                                                                                                var v128 uint8 = v54
                                                                                                                                fmt.Println(v4)
                                                                                                                                _ = v128
                                                                                                                            }
                                                                                                                            v129 := GenS8{F0: 965, F1: 0, F2: 42}
                                                                                                                            v4 = 0.0
                                                                                                                            _ = v129
                                                                                                                        }
                                                                                                                        if true {
                                                                                                                            var v130 uint = 613
                                                                                                                            v112 = 0
                                                                                                                            fmt.Println(v4, v108)
                                                                                                                            _ = v130
                                                                                                                        }
                                                                                                                        var v131 int = v111
                                                                                                                        if (323 != 65535) {
                                                                                                                            var v132 string = (("foo bar" + "a") + (("" + "foo bar") + (("foo bar" + "") + "a")))
                                                                                                                            v53 -= (-(v53))
                                                                                                                            _ = v132
                                                                                                                        }
                                                                                                                        var v133 int = 61
                                                                                                                        _ = v131
                                                                                                                        _ = v133
                                                                                                                    }
                                                                                                                    v124 -= (v124 & v124)
                                                                                                                } else {
                                                                                                                    v124 -= (v124 | v124)
                                                                                                                }
                                                                                                                var v134 float64 = ((0.1 - 1.5) / ((3.14159 + 0.0) * (1.5 + (0.1 + 1.5))))
                                                                                                                _ = v123
                                                                                                                _ = v124
                                                                                                                _ = v134
                                                                                                            }
                                                                                                            if (!v108) {
                                                                                                                v135 := GenS9{F0: -128, F1: (fi23 * fi23), F2: 1}
                                                                                                                if len(v1) > 0 {
                                                                                                                    v1[0] = 2
                                                                                                                }
                                                                                                                var v136 int = -747
                                                                                                                _ = v135
                                                                                                                _ = v136
                                                                                                            } else {
                                                                                                                v137 := GenS10{F0: 1, F1: (v112 | 100), F2: 1}
                                                                                                                _ = v137
                                                                                                            }
                                                                                                            var v138 int16 = v2
                                                                                                            v54 = v99
                                                                                                            _ = v138
                                                                                                        }
                                                                                                    } else {
                                                                                                        fmt.Println(v99, fi37)
                                                                                                        v111 = v111
                                                                                                        if ("" != "foo bar") {
                                                                                                            var v139 uint = 2
                                                                                                            v4 = v4
                                                                                                            _ = v139
                                                                                                        } else {
                                                                                                            v112 = -1
                                                                                                        }
                                                                                                        v111 += v111
                                                                                                    }
                                                                                                case "foo bar":
                                                                                                    v108 = (!(true && (65535 <= 4294967295)))
                                                                                                case "":
                                                                                                    if v108 {
                                                                                                        switch v53 {
                                                                                                        case 4294967295:
                                                                                                            v111 += (fi37 ^ (-(fi37)))
                                                                                                        default:
                                                                                                            v140 := GenS11{F0: 42, F1: 555, F2: (v99 | (^(v99)))}
                                                                                                            _ = v140
                                                                                                        }
                                                                                                        v112 = v112
                                                                                                        v108 = v108
                                                                                                    }
                                                                                                default:
                                                                                                    var v141 uint16 = 925
                                                                                                    _ = v141
                                                                                                }
                                                                                            case "\n\t":
                                                                                                switch fi23 {
                                                                                                case 0:
                                                                                                    v142 := v3[65535]
                                                                                                    _ = v142
                                                                                                case 9223372036854775807:
                                                                                                    if ((v54 % ((v99) | 1)) < (v99 + v99)) {
                                                                                                        fmt.Println(v53)
                                                                                                    }
                                                                                                }
                                                                                            case "foo bar":
                                                                                                switch v111 {
                                                                                                case 0:
                                                                                                    if true {
                                                                                                        switch v111 {
                                                                                                        case 0:
                                                                                                            v99 += v99
                                                                                                            v53 -= fi23
                                                                                                        case -1:
                                                                                                            var v143 float32 = float32(-2.25)
                                                                                                            _ = v143
                                                                                                        case 1:
                                                                                                            var v144 int16 = 441
                                                                                                            _ = v144
                                                                                                        }
                                                                                                        v4 = (v4 * float32(1.5))
                                                                                                    }
                                                                                                    v145 := GenS9{F0: (v112 ^ v100), F1: (fi23 * (v111 % (((v53 % ((fi23) | 1))) | 1))), F2: 617}
                                                                                                    _ = v145
                                                                                                case -9223372036854775808:
                                                                                                    if ((v108 || v108) || false) {
                                                                                                        if (!((v108 && (!((v4 / 0.1) >= v4))) || false)) {
                                                                                                            var v146 float64 = 0.0
                                                                                                            _ = v146
                                                                                                        } else {
                                                                                                            var v147 int64 = 1
                                                                                                            _ = v147
                                                                                                        }
                                                                                                    } else {
                                                                                                        v99 = v54
                                                                                                        if len(v1) > 0 {
                                                                                                            v1[0] = 211
                                                                                                        }
                                                                                                    }
                                                                                                default:
                                                                                                    v148 := GenS10{F0: 0, F1: 122, F2: 195}
                                                                                                    _ = v148
                                                                                                }
                                                                                            default:
                                                                                                var v149 int64 = 1
                                                                                                _ = v149
                                                                                            }
                                                                                        case "hello":
                                                                                            var v150 int16 = (v2 % (((v2 ^ (^(v2)))) | 1))
                                                                                            _ = v150
                                                                                        case "":
                                                                                            v2 += v2
                                                                                        }
                                                                                        _ = v111
                                                                                        _ = v112
                                                                                    } else {
                                                                                        v108 = (0 >= uint64(v54))
                                                                                        fmt.Println(v54)
                                                                                    }
                                                                                }
                                                                                _ = v108
                                                                            } else {
                                                                                if false {
                                                                                    delete(v3, 65535)
                                                                                    v54 -= (v54 >> ((255) & 63))
                                                                                    if len(v1) >= 2 {
                                                                                        v1 = v1[1:len(v1)]
                                                                                    }
                                                                                }
                                                                            }
                                                                            v2 = (v2 * (-(v2)))
                                                                            if false {
                                                                                switch fi37 {
                                                                                case 9223372036854775807:
                                                                                    switch (fi37 / (((fi23 / ((fi23) | 1))) | 1)) {
                                                                                    case 768:
                                                                                        if (("\n\t" == "\n\t") || false) {
                                                                                            var v151 float32 = float32((-2.25 - ((1e10 - (-2.25 / 1e10)) + ((0.0 * 1.5) + (1e10 - (((float64(v4) + -2.25) - 0.0) - 0.0))))))
                                                                                            v100 -= (v100 - -69)
                                                                                            _ = v151
                                                                                        }
                                                                                    case 128:
                                                                                        switch "" {
                                                                                        case "\n\t":
                                                                                            var v152 uint8 = v99
                                                                                            v153 := GenS1{F0: v2, F1: 2147483647, F2: (v100 | v100)}
                                                                                            switch ("hello" + "hello") {
                                                                                            case "a":
                                                                                                v154 := GenS11{F0: 65535, F1: 678, F2: 127}
                                                                                                _ = v154
                                                                                            case "\n\t":
                                                                                                fmt.Println(v152, v54, v100)
                                                                                            default:
                                                                                                if false {
                                                                                                    fmt.Println(fi23, fi23, v2)
                                                                                                } else {
                                                                                                    v155 := GenS11{F0: 2147483648, F1: 42, F2: 0}
                                                                                                    _ = v155
                                                                                                }
                                                                                            }
                                                                                            v156 := GenS12{F0: 1}
                                                                                            v157 := GenS13{F0: 213, F1: v100, F2: 614}
                                                                                            _ = v152
                                                                                            _ = v153
                                                                                            _ = v156
                                                                                            _ = v157
                                                                                        case "":
                                                                                            v4 = -2.25
                                                                                        case "foo bar":
                                                                                            v4 -= v4
                                                                                        default:
                                                                                            v54 = v54
                                                                                        }
                                                                                        switch (v53 + fi37) {
                                                                                        case -1:
                                                                                            var v158 float64 = 0.0
                                                                                            _ = v158
                                                                                        case 42:
                                                                                            var v159 int8 = (v100 ^ 127)
                                                                                            _ = v159
                                                                                        case -912:
                                                                                            v4 += 1e10
                                                                                        default:
                                                                                            v2 += 1000
                                                                                        }
                                                                                    case 0:
                                                                                        var v160 uint8 = (v54 | 204)
                                                                                        _ = v160
                                                                                    }
                                                                                default:
                                                                                    if ((-9223372036854775808 != -920) || false) {
                                                                                        var v161 int = fi23
                                                                                        _ = v161
                                                                                    }
                                                                                    var v162 int8 = v100
                                                                                    var v163 uint64 = 385
                                                                                    _ = v162
                                                                                    _ = v163
                                                                                }
                                                                                v164 := GenS4{F0: 2}
                                                                                switch (fi23 + 42) {
                                                                                case 32768:
                                                                                    if (false && false) {
                                                                                        if (32767 > 32767) {
                                                                                            var v165 string = "hello"
                                                                                            if false {
                                                                                                switch (v165 + (v165 + v165)) {
                                                                                                case "":
                                                                                                    var v166 uint8 = v99
                                                                                                    _ = v166
                                                                                                case "hello":
                                                                                                    _ = len(v3)
                                                                                                case "a":
                                                                                                    if true {
                                                                                                        if false {
                                                                                                            if (0.0 > float64((v4 + 1e10))) {
                                                                                                                v100 = 0
                                                                                                                switch 0 {
                                                                                                                case -1:
                                                                                                                    fmt.Println(v165, v99, v54)
                                                                                                                case 127:
                                                                                                                    var v167 int32 = 2147483647
                                                                                                                    _ = v167
                                                                                                                }
                                                                                                            } else {
                                                                                                                v2 -= (v2 << ((4294967295) & 63))
                                                                                                            }
                                                                                                            fmt.Println(v54, v165, v53)
                                                                                                            v168 := GenS3{F0: v54}
                                                                                                            _ = v168
                                                                                                        } else {
                                                                                                            if (42 != 2) {
                                                                                                                var v169 int32 = 1
                                                                                                                switch (fi23 ^ (fi37 >> ((65535) & 63))) {
                                                                                                                case -1:
                                                                                                                    if (true && (v165 > v165)) {
                                                                                                                        if false {
                                                                                                                            var v170 float32 = v4
                                                                                                                            var v171 uint32 = 65535
                                                                                                                            v172 := GenS13{F0: fi37, F1: 42, F2: v2}
                                                                                                                            _ = v170
                                                                                                                            _ = v171
                                                                                                                            _ = v172
                                                                                                                        }
                                                                                                                        v4 = float32(1e10)
                                                                                                                        var v173 int8 = v100
                                                                                                                        _ = v173
                                                                                                                    } else {
                                                                                                                        _ = len(v1)
                                                                                                                        switch 32768 {
                                                                                                                        case -997:
                                                                                                                            var v174 int = v53
                                                                                                                            _ = v174
                                                                                                                        }
                                                                                                                        var v175 bool = (!false)
                                                                                                                        _ = v175
                                                                                                                    }
                                                                                                                case 32768:
                                                                                                                    if false {
                                                                                                                        var v176 uint64 = 1
                                                                                                                        fmt.Println(v53, fi23)
                                                                                                                        _ = v176
                                                                                                                    }
                                                                                                                case 1:
                                                                                                                    _ = len(v3)
                                                                                                                }
                                                                                                                v3[413] = 416
                                                                                                                _ = v169
                                                                                                            } else {
                                                                                                                _ = len(v1)
                                                                                                                v177 := GenS14{F0: (v165 + ((v165 + "hello") + "foo bar")), F1: 65535, F2: -1}
                                                                                                                if ((v100 - v100) >= v100) {
                                                                                                                    fmt.Println(v99, v99, v54)
                                                                                                                    v4 = 0.1
                                                                                                                } else {
                                                                                                                    delete(v3, 42)
                                                                                                                    v99 = 42
                                                                                                                    v178 := GenS8{F0: 35, F1: 255, F2: 42}
                                                                                                                    if len(v1) > 0 {
                                                                                                                        v1[0] = 561
                                                                                                                    }
                                                                                                                    switch v165 {
                                                                                                                    case "\n\t":
                                                                                                                        fmt.Println(v99)
                                                                                                                    }
                                                                                                                    _ = v178
                                                                                                                }
                                                                                                                _ = len(v3)
                                                                                                                switch v53 {
                                                                                                                case 42:
                                                                                                                    var v179 uint64 = 42
                                                                                                                    _ = v179
                                                                                                                case 2147483647:
                                                                                                                    switch "hello" {
                                                                                                                    case "":
                                                                                                                        _ = len(v1)
                                                                                                                    case "\n\t":
                                                                                                                        switch ("hello" + v165) {
                                                                                                                        case "hello":
                                                                                                                            v100 = v100
                                                                                                                        case "a":
                                                                                                                            var v180 float32 = 1.5
                                                                                                                            _ = v180
                                                                                                                        default:
                                                                                                                            if ((v54 % ((v99) | 1)) != 42) {
                                                                                                                                var v181 float64 = -2.25
                                                                                                                                _ = len(v1)
                                                                                                                                _ = v181
                                                                                                                            } else {
                                                                                                                                switch v165 {
                                                                                                                                case "foo bar":
                                                                                                                                    v182 := GenS15{F0: 1, F1: true}
                                                                                                                                    _ = v182
                                                                                                                                case "a":
                                                                                                                                    v2 += v2
                                                                                                                                case "\n\t":
                                                                                                                                    v183 := v3[2147483647]
                                                                                                                                    _ = v183
                                                                                                                                }
                                                                                                                                v99 += v99
                                                                                                                                if false {
                                                                                                                                    if true {
                                                                                                                                        v3[2] = 1
                                                                                                                                    } else {
                                                                                                                                        if true {
                                                                                                                                            if true {
                                                                                                                                                v100 -= v100
                                                                                                                                                switch 32767 {
                                                                                                                                                case -2147483648:
                                                                                                                                                    var v184 int8 = 127
                                                                                                                                                    _ = v184
                                                                                                                                                case -656:
                                                                                                                                                    if true {
                                                                                                                                                        if false {
                                                                                                                                                            var v185 float64 = 1.5
                                                                                                                                                            _ = v185
                                                                                                                                                        }
                                                                                                                                                        var v186 uint64 = 2
                                                                                                                                                        _ = v186
                                                                                                                                                    }
                                                                                                                                                case 826:
                                                                                                                                                    v187 := GenS9{F0: (v100 & v100), F1: fi23, F2: 255}
                                                                                                                                                    switch v165 {
                                                                                                                                                    case "foo bar":
                                                                                                                                                        var v188 uint8 = v54
                                                                                                                                                        _ = v188
                                                                                                                                                    case "a":
                                                                                                                                                        switch v165 {
                                                                                                                                                        case "a":
                                                                                                                                                            v165 = v165
                                                                                                                                                        case "foo bar":
                                                                                                                                                            var v189 int64 = -9223372036854775808
                                                                                                                                                            _ = v189
                                                                                                                                                        case "":
                                                                                                                                                            fmt.Println(v2, v99)
                                                                                                                                                        default:
                                                                                                                                                            _ = len(v3)
                                                                                                                                                        }
                                                                                                                                                    case "hello":
                                                                                                                                                        v190 := GenS10{F0: 2, F1: (v100 / (((^(v100))) | 1)), F2: 100}
                                                                                                                                                        _ = v190
                                                                                                                                                    }
                                                                                                                                                    fmt.Println(v53, v99, v99)
                                                                                                                                                    _ = v187
                                                                                                                                                default:
                                                                                                                                                    v191 := GenS9{F0: (v100 - (v100 | (v100 ^ (v100 >> ((1) & 63))))), F1: (fi23 / ((288) | 1)), F2: 138}
                                                                                                                                                    _ = v191
                                                                                                                                                }
                                                                                                                                            } else {
                                                                                                                                                if true {
                                                                                                                                                    var v192 bool = true
                                                                                                                                                    _ = v192
                                                                                                                                                } else {
                                                                                                                                                    if (977 >= 0) {
                                                                                                                                                        fmt.Println(v100, fi37, v165)
                                                                                                                                                        v99 = (v54 * v99)
                                                                                                                                                        v54 = v99
                                                                                                                                                    } else {
                                                                                                                                                        if false {
                                                                                                                                                            v193 := GenS4{F0: 2}
                                                                                                                                                            var v194 int64 = 42
                                                                                                                                                            if (v194 < v194) {
                                                                                                                                                                var v195 int8 = (^(v100))
                                                                                                                                                                _ = v195
                                                                                                                                                            }
                                                                                                                                                            _ = v193
                                                                                                                                                            _ = v194
                                                                                                                                                        } else {
                                                                                                                                                            var v196 uint32 = 245
                                                                                                                                                            var v197 uint64 = 847
                                                                                                                                                            v198 := GenS16{F0: -636, F1: (v100 << ((uint(v196)) & 63))}
                                                                                                                                                            _ = v196
                                                                                                                                                            _ = v197
                                                                                                                                                            _ = v198
                                                                                                                                                        }
                                                                                                                                                    }
                                                                                                                                                    fmt.Println(v165, fi37)
                                                                                                                                                    switch v165 {
                                                                                                                                                    case "\n\t":
                                                                                                                                                        v2 -= v2
                                                                                                                                                    default:
                                                                                                                                                        v54 = v54
                                                                                                                                                    }
                                                                                                                                                }
                                                                                                                                            }
                                                                                                                                            switch (((v165 + v165) + ((v165 + v165) + v165)) + v165) {
                                                                                                                                            case "":
                                                                                                                                                var v199 uint16 = 2
                                                                                                                                                var v200 bool = true
                                                                                                                                                delete(v3, 1)
                                                                                                                                                _ = v199
                                                                                                                                                _ = v200
                                                                                                                                            default:
                                                                                                                                                if (1 <= 18446744073709551615) {
                                                                                                                                                    switch (^(fi37)) {
                                                                                                                                                    case 32767:
                                                                                                                                                        v4 -= 0.1
                                                                                                                                                    case 473:
                                                                                                                                                        v201 := GenS7{F0: 42, F1: int64(v99)}
                                                                                                                                                        _ = v201
                                                                                                                                                    }
                                                                                                                                                    var v202 int16 = v2
                                                                                                                                                    _ = v202
                                                                                                                                                }
                                                                                                                                            }
                                                                                                                                            var v203 bool = false
                                                                                                                                            if v203 {
                                                                                                                                                var v204 string = "foo bar"
                                                                                                                                                v205 := GenS10{F0: 1, F1: (v100 << ((1) & 63)), F2: 0}
                                                                                                                                                if v203 {
                                                                                                                                                    switch v204 {
                                                                                                                                                    case "a":
                                                                                                                                                        var v206 string = v165
                                                                                                                                                        _ = v206
                                                                                                                                                    }
                                                                                                                                                } else {
                                                                                                                                                    var v207 uint64 = 1
                                                                                                                                                    _ = v207
                                                                                                                                                }
                                                                                                                                                _ = v204
                                                                                                                                                _ = v205
                                                                                                                                            } else {
                                                                                                                                                switch ("hello" + v165) {
                                                                                                                                                case "":
                                                                                                                                                    fmt.Println(v2)
                                                                                                                                                default:
                                                                                                                                                    switch (fi37 * 0) {
                                                                                                                                                    case 2147483647:
                                                                                                                                                        var v208 int = (fi23 - 9223372036854775807)
                                                                                                                                                        _ = v208
                                                                                                                                                    default:
                                                                                                                                                        if (false || v203) {
                                                                                                                                                            v54 += v99
                                                                                                                                                        } else {
                                                                                                                                                            var v209 float64 = 0.1
                                                                                                                                                            if (!v203) {
                                                                                                                                                                if v203 {
                                                                                                                                                                    if (!v203) {
                                                                                                                                                                        var v210 float64 = float64(v4)
                                                                                                                                                                        _ = v210
                                                                                                                                                                    } else {
                                                                                                                                                                        if (559 > 0) {
                                                                                                                                                                            var v211 int = (fi23 & (v53 % ((fi23) | 1)))
                                                                                                                                                                            if (v203 || v203) {
                                                                                                                                                                                var v212 float32 = 0.1
                                                                                                                                                                                v2 += 42
                                                                                                                                                                                _ = v212
                                                                                                                                                                            }
                                                                                                                                                                            _ = v211
                                                                                                                                                                        }
                                                                                                                                                                    }
                                                                                                                                                                    var v213 float64 = float64(float32(v209))
                                                                                                                                                                    _ = v213
                                                                                                                                                                } else {
                                                                                                                                                                    var v214 int32 = 2147483647
                                                                                                                                                                    v203 = (fi23 == v53)
                                                                                                                                                                    var v215 string = ""
                                                                                                                                                                    if v203 {
                                                                                                                                                                        delete(v3, 769)
                                                                                                                                                                    }
                                                                                                                                                                    switch (v53 % ((995) | 1)) {
                                                                                                                                                                    case -2147483648:
                                                                                                                                                                        fmt.Println(v53, v215, v2)
                                                                                                                                                                    case 128:
                                                                                                                                                                        if ((0.1 < v4) || (v100 == -84)) {
                                                                                                                                                                            fmt.Println(v2, v165)
                                                                                                                                                                        } else {
                                                                                                                                                                            v216 := GenS17{F0: "hello", F1: 0}
                                                                                                                                                                            v4 = (v4 * float32(v209))
                                                                                                                                                                            _ = v216
                                                                                                                                                                        }
                                                                                                                                                                    default:
                                                                                                                                                                        switch ((v165 + (v165 + "foo bar")) + v215) {
                                                                                                                                                                        case "\n\t":
                                                                                                                                                                            var v217 bool = ((!(355 != 0)) || (498 > 1))
                                                                                                                                                                            _ = v217
                                                                                                                                                                        case "hello":
                                                                                                                                                                            if (v203 && (65535 < 0)) {
                                                                                                                                                                                if ((!v203) || false) {
                                                                                                                                                                                    v99 += v99
                                                                                                                                                                                    var v218 bool = true
                                                                                                                                                                                    _ = v218
                                                                                                                                                                                } else {
                                                                                                                                                                                    var v219 int = fi23
                                                                                                                                                                                    fmt.Println(v215, v214)
                                                                                                                                                                                    _ = v219
                                                                                                                                                                                }
                                                                                                                                                                                var v220 uint64 = 640
                                                                                                                                                                                if len(v1) > 0 {
                                                                                                                                                                                    v1[0] = 328
                                                                                                                                                                                }
                                                                                                                                                                                _ = v220
                                                                                                                                                                            } else {
                                                                                                                                                                                v4 = (v4 + (1.5 - v4))
                                                                                                                                                                                fmt.Println(v214, v165, v215)
                                                                                                                                                                            }
                                                                                                                                                                        case "foo bar":
                                                                                                                                                                            v100 -= v100
                                                                                                                                                                            var v221 uint = 95
                                                                                                                                                                            v2 -= v2
                                                                                                                                                                            _ = v221
                                                                                                                                                                        default:
                                                                                                                                                                            fmt.Println(v54, v53)
                                                                                                                                                                        }
                                                                                                                                                                    }
                                                                                                                                                                    _ = v214
                                                                                                                                                                    _ = v215
                                                                                                                                                                }
                                                                                                                                                                v209 = v209
                                                                                                                                                                fmt.Println(v2, v100, v54)
                                                                                                                                                            }
                                                                                                                                                            if false {
                                                                                                                                                                v4 += 0.1
                                                                                                                                                                switch v165 {
                                                                                                                                                                case "a":
                                                                                                                                                                    if (!v203) {
                                                                                                                                                                        v209 += (3.14159 + -2.25)
                                                                                                                                                                        v222 := GenS16{F0: -9223372036854775808, F1: v100}
                                                                                                                                                                        _ = v222
                                                                                                                                                                    }
                                                                                                                                                                case "":
                                                                                                                                                                    v223 := GenS1{F0: v2, F1: 573, F2: (^(v100))}
                                                                                                                                                                    v224 := GenS10{F0: 1, F1: v100, F2: 18446744073709551615}
                                                                                                                                                                    _ = v223
                                                                                                                                                                    _ = v224
                                                                                                                                                                }
                                                                                                                                                            }
                                                                                                                                                            _ = v209
                                                                                                                                                        }
                                                                                                                                                    }
                                                                                                                                                }
                                                                                                                                                switch v165 {
                                                                                                                                                case "":
                                                                                                                                                    v53 = (fi23 + (fi23 ^ -2147483648))
                                                                                                                                                case "a":
                                                                                                                                                    var v225 float64 = -2.25
                                                                                                                                                    _ = v225
                                                                                                                                                case "\n\t":
                                                                                                                                                    if len(v1) >= 2 {
                                                                                                                                                        v1 = v1[1:len(v1)]
                                                                                                                                                    }
                                                                                                                                                default:
                                                                                                                                                    if ((v4 == v4) && (-2.25 == float32((float64(v4) + (1.5 - -2.25))))) {
                                                                                                                                                        var v226 float64 = 1.5
                                                                                                                                                        if true {
                                                                                                                                                            fmt.Println(v203, v54, v53)
                                                                                                                                                            fmt.Println(v4, v203)
                                                                                                                                                            if v203 {
                                                                                                                                                                _ = len(v1)
                                                                                                                                                                v2 = (v2 + -1)
                                                                                                                                                                var v227 uint64 = 42
                                                                                                                                                                _ = v227
                                                                                                                                                            } else {
                                                                                                                                                                var v228 bool = false
                                                                                                                                                                if v228 {
                                                                                                                                                                    v165 = (((v165 + v165) + v165) + (("foo bar" + v165) + v165))
                                                                                                                                                                    if ("foo bar" == "a") {
                                                                                                                                                                        if (v165 == v165) {
                                                                                                                                                                            var v229 uint8 = (^(v99))
                                                                                                                                                                            _ = v229
                                                                                                                                                                        } else {
                                                                                                                                                                            fmt.Println(v165)
                                                                                                                                                                            switch "foo bar" {
                                                                                                                                                                            case "a":
                                                                                                                                                                                if true {
                                                                                                                                                                                    v230 := v3[65535]
                                                                                                                                                                                    _ = v230
                                                                                                                                                                                } else {
                                                                                                                                                                                    if (v203 || v203) {
                                                                                                                                                                                        if (v54 == 244) {
                                                                                                                                                                                            fmt.Println(v203, v54)
                                                                                                                                                                                            if false {
                                                                                                                                                                                                fmt.Println(v4)
                                                                                                                                                                                                var v231 uint = 120
                                                                                                                                                                                                fmt.Println(v4, v54)
                                                                                                                                                                                                v2 += (v2 / ((v2) | 1))
                                                                                                                                                                                                _ = v231
                                                                                                                                                                                            } else {
                                                                                                                                                                                                if ((1e10 / v226) <= v226) {
                                                                                                                                                                                                    switch fi37 {
                                                                                                                                                                                                    case 9223372036854775807:
                                                                                                                                                                                                        v4 = 3.14159
                                                                                                                                                                                                    default:
                                                                                                                                                                                                        var v232 int8 = (^(v100))
                                                                                                                                                                                                        _ = v232
                                                                                                                                                                                                    }
                                                                                                                                                                                                } else {
                                                                                                                                                                                                    v233 := GenS18{F0: (v100 & (v100 % ((1) | 1))), F1: uint64(v99)}
                                                                                                                                                                                                    v1 = append(v1, 90)
                                                                                                                                                                                                    if len(v1) > 0 {
                                                                                                                                                                                                        v1[0] = 644
                                                                                                                                                                                                    }
                                                                                                                                                                                                    _ = v233
                                                                                                                                                                                                }
                                                                                                                                                                                            }
                                                                                                                                                                                            v234 := GenS19{F0: 65535, F1: 554, F2: (-(v100))}
                                                                                                                                                                                            if (!false) {
                                                                                                                                                                                                v4 += v4
                                                                                                                                                                                                switch fi23 {
                                                                                                                                                                                                case 32767:
                                                                                                                                                                                                    v235 := GenS20{F0: 0, F1: -9223372036854775808, F2: v165}
                                                                                                                                                                                                    _ = v235
                                                                                                                                                                                                case 128:
                                                                                                                                                                                                    if ((v2 + (v2 % (((-(v2))) | 1))) == -206) {
                                                                                                                                                                                                        var v236 float64 = v226
                                                                                                                                                                                                        _ = v236
                                                                                                                                                                                                    } else {
                                                                                                                                                                                                        if (true || v203) {
                                                                                                                                                                                                            v226 += 1e10
                                                                                                                                                                                                        } else {
                                                                                                                                                                                                            v226 = 0.1
                                                                                                                                                                                                            if len(v1) >= 2 {
                                                                                                                                                                                                                v1 = v1[1:len(v1)]
                                                                                                                                                                                                            }
                                                                                                                                                                                                            v4 -= v4
                                                                                                                                                                                                            if v203 {
                                                                                                                                                                                                                var v237 int16 = v2
                                                                                                                                                                                                                if v203 {
                                                                                                                                                                                                                    if (42 >= 20) {
                                                                                                                                                                                                                        var v238 string = v165
                                                                                                                                                                                                                        _ = v238
                                                                                                                                                                                                                    }
                                                                                                                                                                                                                    delete(v3, 65535)
                                                                                                                                                                                                                    v226 = v226
                                                                                                                                                                                                                    var v239 uint8 = v99
                                                                                                                                                                                                                    _ = v239
                                                                                                                                                                                                                }
                                                                                                                                                                                                                _ = v237
                                                                                                                                                                                                            } else {
                                                                                                                                                                                                                v240 := GenS8{F0: 65535, F1: 135, F2: 2147483647}
                                                                                                                                                                                                                if len(v1) >= 2 {
                                                                                                                                                                                                                    v1 = v1[1:len(v1)]
                                                                                                                                                                                                                }
                                                                                                                                                                                                                var v241 string = v165
                                                                                                                                                                                                                _ = v240
                                                                                                                                                                                                                _ = v241
                                                                                                                                                                                                            }
                                                                                                                                                                                                        }
                                                                                                                                                                                                        fmt.Println(v226)
                                                                                                                                                                                                        switch fi37 {
                                                                                                                                                                                                        case 0:
                                                                                                                                                                                                            switch 966 {
                                                                                                                                                                                                            case 4294967295:
                                                                                                                                                                                                                if (359 != 42) {
                                                                                                                                                                                                                    v99 = 1
                                                                                                                                                                                                                }
                                                                                                                                                                                                            default:
                                                                                                                                                                                                                v242 := GenS19{F0: 1, F1: 724, F2: -1}
                                                                                                                                                                                                                _ = v242
                                                                                                                                                                                                            }
                                                                                                                                                                                                        case 128:
                                                                                                                                                                                                            var v243 uint64 = 2
                                                                                                                                                                                                            _ = v243
                                                                                                                                                                                                        }
                                                                                                                                                                                                    }
                                                                                                                                                                                                case 9223372036854775807:
                                                                                                                                                                                                    var v244 int8 = v100
                                                                                                                                                                                                    _ = v244
                                                                                                                                                                                                default:
                                                                                                                                                                                                    fmt.Println(v100, v53, v228)
                                                                                                                                                                                                }
                                                                                                                                                                                                v53 = fi37
                                                                                                                                                                                                if v228 {
                                                                                                                                                                                                    switch fi37 {
                                                                                                                                                                                                    case 9223372036854775807:
                                                                                                                                                                                                        v3[814] = 2147483648
                                                                                                                                                                                                    case 2147483647:
                                                                                                                                                                                                        switch v165 {
                                                                                                                                                                                                        case "\n\t":
                                                                                                                                                                                                            v226 -= float64((float32(v226) + (v4 + v4)))
                                                                                                                                                                                                        default:
                                                                                                                                                                                                            delete(v3, 2147483648)
                                                                                                                                                                                                        }
                                                                                                                                                                                                    case 4294967295:
                                                                                                                                                                                                        v245 := v3[0]
                                                                                                                                                                                                        _ = v245
                                                                                                                                                                                                    }
                                                                                                                                                                                                    var v246 bool = v203
                                                                                                                                                                                                    if (v246 || v246) {
                                                                                                                                                                                                        v247 := v3[2]
                                                                                                                                                                                                        _ = len(v3)
                                                                                                                                                                                                        v165 = "\n\t"
                                                                                                                                                                                                        if true {
                                                                                                                                                                                                            switch (v165 + ((v165 + ((("" + "a") + v165) + v165)) + v165)) {
                                                                                                                                                                                                            case "\n\t":
                                                                                                                                                                                                                v228 = false
                                                                                                                                                                                                            }
                                                                                                                                                                                                            switch "\n\t" {
                                                                                                                                                                                                            case "hello":
                                                                                                                                                                                                                var v248 int8 = v100
                                                                                                                                                                                                                _ = v248
                                                                                                                                                                                                            default:
                                                                                                                                                                                                                var v249 int8 = -25
                                                                                                                                                                                                                _ = v249
                                                                                                                                                                                                            }
                                                                                                                                                                                                            v250 := GenS10{F0: 2, F1: -1, F2: 42}
                                                                                                                                                                                                            _ = v250
                                                                                                                                                                                                        } else {
                                                                                                                                                                                                            v226 += 0.1
                                                                                                                                                                                                            fmt.Println(fi23, v2, v53)
                                                                                                                                                                                                            fmt.Println(v54, fi23, v228)
                                                                                                                                                                                                        }
                                                                                                                                                                                                        _ = v247
                                                                                                                                                                                                    } else {
                                                                                                                                                                                                        switch fi23 {
                                                                                                                                                                                                        case 32768:
                                                                                                                                                                                                            switch v165 {
                                                                                                                                                                                                            case "a":
                                                                                                                                                                                                                v251 := GenS21{F0: 0, F1: v100, F2: 128}
                                                                                                                                                                                                                _ = v251
                                                                                                                                                                                                            case "hello":
                                                                                                                                                                                                                if (-1 != -1) {
                                                                                                                                                                                                                    fmt.Println(fi37, v203)
                                                                                                                                                                                                                } else {
                                                                                                                                                                                                                    fmt.Println(v2)
                                                                                                                                                                                                                }
                                                                                                                                                                                                            default:
                                                                                                                                                                                                                switch 671 {
                                                                                                                                                                                                                case 186:
                                                                                                                                                                                                                    fmt.Println(v165)
                                                                                                                                                                                                                }
                                                                                                                                                                                                                switch (fi23 & (fi23 - fi37)) {
                                                                                                                                                                                                                case -9223372036854775808:
                                                                                                                                                                                                                    switch 9223372036854775807 {
                                                                                                                                                                                                                    case -9223372036854775808:
                                                                                                                                                                                                                        var v252 uint64 = 1
                                                                                                                                                                                                                        _ = v252
                                                                                                                                                                                                                    case -915:
                                                                                                                                                                                                                        switch "" {
                                                                                                                                                                                                                        case "":
                                                                                                                                                                                                                            _ = len(v1)
                                                                                                                                                                                                                        case "a":
                                                                                                                                                                                                                            if false {
                                                                                                                                                                                                                                v99 = v99
                                                                                                                                                                                                                                fmt.Println(v99, v246)
                                                                                                                                                                                                                                delete(v3, 1)
                                                                                                                                                                                                                                delete(v3, 742)
                                                                                                                                                                                                                                if true {
                                                                                                                                                                                                                                    switch v165 {
                                                                                                                                                                                                                                    case "hello":
                                                                                                                                                                                                                                        v228 = ((v226 != -2.25) && false)
                                                                                                                                                                                                                                    case "a":
                                                                                                                                                                                                                                        if true {
                                                                                                                                                                                                                                            v4 = v4
                                                                                                                                                                                                                                            v3[2147483648] = 2147483648
                                                                                                                                                                                                                                        } else {
                                                                                                                                                                                                                                            var v253 bool = false
                                                                                                                                                                                                                                            v254 := GenS6{F0: (v100 ^ (v100 << ((255) & 63)))}
                                                                                                                                                                                                                                            switch fi23 {
                                                                                                                                                                                                                                            case 0:
                                                                                                                                                                                                                                                fmt.Println(v165, v253, v100)
                                                                                                                                                                                                                                            case -719:
                                                                                                                                                                                                                                                var v255 uint16 = 132
                                                                                                                                                                                                                                                _ = v255
                                                                                                                                                                                                                                            case 752:
                                                                                                                                                                                                                                                v4 = v4
                                                                                                                                                                                                                                            }
                                                                                                                                                                                                                                            fmt.Println(v99)
                                                                                                                                                                                                                                            var v256 uint8 = 255
                                                                                                                                                                                                                                            _ = v253
                                                                                                                                                                                                                                            _ = v254
                                                                                                                                                                                                                                            _ = v256
                                                                                                                                                                                                                                        }
                                                                                                                                                                                                                                    case "foo bar":
                                                                                                                                                                                                                                        var v257 float32 = ((((1e10 * (v4 * v4)) - (v4 + v4)) / float32((v226 * 0.0))) + 0.0)
                                                                                                                                                                                                                                        _ = v257
                                                                                                                                                                                                                                    }
                                                                                                                                                                                                                                }
                                                                                                                                                                                                                            } else {
                                                                                                                                                                                                                                v258 := GenS22{F0: 255, F1: 42}
                                                                                                                                                                                                                                if len(v1) > 0 {
                                                                                                                                                                                                                                    v1[0] = 768
                                                                                                                                                                                                                                }
                                                                                                                                                                                                                                _ = v258
                                                                                                                                                                                                                            }
                                                                                                                                                                                                                        default:
                                                                                                                                                                                                                            if false {
                                                                                                                                                                                                                                v259 := GenS7{F0: -8, F1: -1}
                                                                                                                                                                                                                                var v260 uint = 49
                                                                                                                                                                                                                                fmt.Println(v260, v226)
                                                                                                                                                                                                                                _ = v259
                                                                                                                                                                                                                                _ = v260
                                                                                                                                                                                                                            }
                                                                                                                                                                                                                            if false {
                                                                                                                                                                                                                                v1 = append(v1, 4294967295)
                                                                                                                                                                                                                                var v261 int64 = 895
                                                                                                                                                                                                                                _ = v261
                                                                                                                                                                                                                            }
                                                                                                                                                                                                                        }
                                                                                                                                                                                                                    case -568:
                                                                                                                                                                                                                        var v262 uint64 = 1
                                                                                                                                                                                                                        _ = v262
                                                                                                                                                                                                                    default:
                                                                                                                                                                                                                        delete(v3, 554)
                                                                                                                                                                                                                    }
                                                                                                                                                                                                                case -540:
                                                                                                                                                                                                                    var v263 float64 = ((v226 + (float64(1e10) * v226)) * v226)
                                                                                                                                                                                                                    if v203 {
                                                                                                                                                                                                                        fmt.Println(v263, v263)
                                                                                                                                                                                                                        v1 = append(v1, 892)
                                                                                                                                                                                                                        var v264 uint32 = uint32(v54)
                                                                                                                                                                                                                        _ = v264
                                                                                                                                                                                                                    } else {
                                                                                                                                                                                                                        fmt.Println(v99)
                                                                                                                                                                                                                    }
                                                                                                                                                                                                                    _ = v263
                                                                                                                                                                                                                case 140:
                                                                                                                                                                                                                    if false {
                                                                                                                                                                                                                        v265 := GenS11{F0: 65535, F1: 2147483648, F2: 201}
                                                                                                                                                                                                                        if (v4 > float32(v226)) {
                                                                                                                                                                                                                            var v266 int8 = v100
                                                                                                                                                                                                                            v267 := GenS11{F0: 65535, F1: 182, F2: 255}
                                                                                                                                                                                                                            var v268 uint8 = (v54 ^ v54)
                                                                                                                                                                                                                            _ = v266
                                                                                                                                                                                                                            _ = v267
                                                                                                                                                                                                                            _ = v268
                                                                                                                                                                                                                        } else {
                                                                                                                                                                                                                            v1 = append(v1, 4294967295)
                                                                                                                                                                                                                            var v269 bool = (false || v203)
                                                                                                                                                                                                                            _ = v269
                                                                                                                                                                                                                        }
                                                                                                                                                                                                                        fmt.Println(v53)
                                                                                                                                                                                                                        _ = v265
                                                                                                                                                                                                                    } else {
                                                                                                                                                                                                                        switch v165 {
                                                                                                                                                                                                                        case "a":
                                                                                                                                                                                                                            v270 := GenS23{F0: 128}
                                                                                                                                                                                                                            _ = v270
                                                                                                                                                                                                                        case "\n\t":
                                                                                                                                                                                                                            v271 := GenS5{F0: false}
                                                                                                                                                                                                                            _ = v271
                                                                                                                                                                                                                        default:
                                                                                                                                                                                                                            var v272 uint32 = 261
                                                                                                                                                                                                                            _ = v272
                                                                                                                                                                                                                        }
                                                                                                                                                                                                                        if v228 {
                                                                                                                                                                                                                            var v273 string = v165
                                                                                                                                                                                                                            if (!(true || true)) {
                                                                                                                                                                                                                                switch (("" + "\n\t") + "foo bar") {
                                                                                                                                                                                                                                case "foo bar":
                                                                                                                                                                                                                                    if (v246 && true) {
                                                                                                                                                                                                                                        if v228 {
                                                                                                                                                                                                                                            v1 = append(v1, 713)
                                                                                                                                                                                                                                            fmt.Println(v246)
                                                                                                                                                                                                                                            var v274 uint = 923
                                                                                                                                                                                                                                            v275 := GenS1{F0: (v2 << ((v274) & 63)), F1: 42, F2: -49}
                                                                                                                                                                                                                                            _ = v274
                                                                                                                                                                                                                                            _ = v275
                                                                                                                                                                                                                                        }
                                                                                                                                                                                                                                        fmt.Println(v203, v228, v2)
                                                                                                                                                                                                                                    }
                                                                                                                                                                                                                                    fmt.Println(fi23)
                                                                                                                                                                                                                                    if (!v228) {
                                                                                                                                                                                                                                        if v228 {
                                                                                                                                                                                                                                            if v246 {
                                                                                                                                                                                                                                                v276 := GenS12{F0: 32767}
                                                                                                                                                                                                                                                _ = v276
                                                                                                                                                                                                                                            }
                                                                                                                                                                                                                                            switch (fi37 << ((223) & 63)) {
                                                                                                                                                                                                                                            case -176:
                                                                                                                                                                                                                                                var v277 uint16 = 2
                                                                                                                                                                                                                                                _ = v277
                                                                                                                                                                                                                                            case 32768:
                                                                                                                                                                                                                                                if v203 {
                                                                                                                                                                                                                                                    v246 = v228
                                                                                                                                                                                                                                                    v278 := GenS4{F0: 2}
                                                                                                                                                                                                                                                    switch 42 {
                                                                                                                                                                                                                                                    case -1:
                                                                                                                                                                                                                                                        switch (v53 & (fi23 - (fi37 & 128))) {
                                                                                                                                                                                                                                                        case 32767:
                                                                                                                                                                                                                                                            v279 := GenS24{F0: 65535, F1: int16(v54), F2: 229}
                                                                                                                                                                                                                                                            _ = v279
                                                                                                                                                                                                                                                        }
                                                                                                                                                                                                                                                        var v280 string = "\n\t"
                                                                                                                                                                                                                                                        _ = v280
                                                                                                                                                                                                                                                    case 982:
                                                                                                                                                                                                                                                        if (2 == 2) {
                                                                                                                                                                                                                                                            v203 = v246
                                                                                                                                                                                                                                                        } else {
                                                                                                                                                                                                                                                            var v281 int32 = 0
                                                                                                                                                                                                                                                            fmt.Println(fi23, fi37)
                                                                                                                                                                                                                                                            delete(v3, 365)
                                                                                                                                                                                                                                                            _ = v281
                                                                                                                                                                                                                                                        }
                                                                                                                                                                                                                                                    }
                                                                                                                                                                                                                                                    _ = v278
                                                                                                                                                                                                                                                } else {
                                                                                                                                                                                                                                                    v1 = append(v1, 2147483648)
                                                                                                                                                                                                                                                    delete(v3, 2147483647)
                                                                                                                                                                                                                                                }
                                                                                                                                                                                                                                                v100 += 100
                                                                                                                                                                                                                                                var v282 int32 = 0
                                                                                                                                                                                                                                                _ = v282
                                                                                                                                                                                                                                            case -726:
                                                                                                                                                                                                                                                v228 = ((v2 % ((v2) | 1)) < (v2 ^ v2))
                                                                                                                                                                                                                                            default:
                                                                                                                                                                                                                                                delete(v3, 4294967295)
                                                                                                                                                                                                                                            }
                                                                                                                                                                                                                                        }
                                                                                                                                                                                                                                        switch (fi23 ^ (v53 / (((^(fi23))) | 1))) {
                                                                                                                                                                                                                                        case 128:
                                                                                                                                                                                                                                            if false {
                                                                                                                                                                                                                                                if v246 {
                                                                                                                                                                                                                                                    switch int(v100) {
                                                                                                                                                                                                                                                    case -1:
                                                                                                                                                                                                                                                        var v283 int32 = -1
                                                                                                                                                                                                                                                        _ = v283
                                                                                                                                                                                                                                                    case 503:
                                                                                                                                                                                                                                                        delete(v3, 579)
                                                                                                                                                                                                                                                    default:
                                                                                                                                                                                                                                                        if v228 {
                                                                                                                                                                                                                                                            v284 := GenS25{F0: 240, F1: -488, F2: -1}
                                                                                                                                                                                                                                                            _ = v284
                                                                                                                                                                                                                                                        }
                                                                                                                                                                                                                                                    }
                                                                                                                                                                                                                                                    var v285 bool = v228
                                                                                                                                                                                                                                                    v273 = "hello"
                                                                                                                                                                                                                                                    switch 1 {
                                                                                                                                                                                                                                                    case -357:
                                                                                                                                                                                                                                                        fmt.Println(v2, v226, v99)
                                                                                                                                                                                                                                                    default:
                                                                                                                                                                                                                                                        v285 = ((v100 * 42) >= v100)
                                                                                                                                                                                                                                                    }
                                                                                                                                                                                                                                                    _ = v285
                                                                                                                                                                                                                                                } else {
                                                                                                                                                                                                                                                    v3[65535] = 865
                                                                                                                                                                                                                                                    switch v165 {
                                                                                                                                                                                                                                                    case "a":
                                                                                                                                                                                                                                                        fmt.Println(fi23, v226)
                                                                                                                                                                                                                                                    }
                                                                                                                                                                                                                                                }
                                                                                                                                                                                                                                                v286 := v3[2147483648]
                                                                                                                                                                                                                                                v287 := GenS4{F0: 0}
                                                                                                                                                                                                                                                v4 = v4
                                                                                                                                                                                                                                                v203 = true
                                                                                                                                                                                                                                                _ = v286
                                                                                                                                                                                                                                                _ = v287
                                                                                                                                                                                                                                            }
                                                                                                                                                                                                                                        }
                                                                                                                                                                                                                                        v203 = (!((42 > 42) || ((((0.0 / (v4 + v4)) / 0.1) / (v4 - (float32(v226) + v4))) < v4)))
                                                                                                                                                                                                                                    } else {
                                                                                                                                                                                                                                        v288 := GenS11{F0: 42, F1: 266, F2: (v99 / (((v54 % ((v54) | 1))) | 1))}
                                                                                                                                                                                                                                        v246 = (v54 > (^(v54)))
                                                                                                                                                                                                                                        var v289 bool = (9223372036854775807 != -9223372036854775808)
                                                                                                                                                                                                                                        _ = v288
                                                                                                                                                                                                                                        _ = v289
                                                                                                                                                                                                                                    }
                                                                                                                                                                                                                                    if false {
                                                                                                                                                                                                                                        v54 = 0
                                                                                                                                                                                                                                        var v290 int64 = -1
                                                                                                                                                                                                                                        _ = v290
                                                                                                                                                                                                                                    } else {
                                                                                                                                                                                                                                        var v291 uint32 = 42
                                                                                                                                                                                                                                        _ = v291
                                                                                                                                                                                                                                    }
                                                                                                                                                                                                                                case "hello":
                                                                                                                                                                                                                                    if (18446744073709551615 == 581) {
                                                                                                                                                                                                                                        var v292 uint8 = 0
                                                                                                                                                                                                                                        _ = v292
                                                                                                                                                                                                                                    }
                                                                                                                                                                                                                                default:
                                                                                                                                                                                                                                    v293 := GenS4{F0: 42}
                                                                                                                                                                                                                                    _ = v293
                                                                                                                                                                                                                                }
                                                                                                                                                                                                                                switch 2147483647 {
                                                                                                                                                                                                                                case 128:
                                                                                                                                                                                                                                    v294 := v3[2147483648]
                                                                                                                                                                                                                                    _ = v294
                                                                                                                                                                                                                                case -1:
                                                                                                                                                                                                                                    if len(v1) >= 2 {
                                                                                                                                                                                                                                        v1 = v1[1:len(v1)]
                                                                                                                                                                                                                                    }
                                                                                                                                                                                                                                }
                                                                                                                                                                                                                                var v295 uint32 = 2
                                                                                                                                                                                                                                fmt.Println(v295)
                                                                                                                                                                                                                                _ = v295
                                                                                                                                                                                                                            } else {
                                                                                                                                                                                                                                var v296 int8 = 42
                                                                                                                                                                                                                                _ = v296
                                                                                                                                                                                                                            }
                                                                                                                                                                                                                            _ = v273
                                                                                                                                                                                                                        }
                                                                                                                                                                                                                        v297 := GenS11{F0: 970, F1: 0, F2: 2}
                                                                                                                                                                                                                        _ = v297
                                                                                                                                                                                                                    }
                                                                                                                                                                                                                }
                                                                                                                                                                                                            }
                                                                                                                                                                                                        default:
                                                                                                                                                                                                            if false {
                                                                                                                                                                                                                v298 := GenS9{F0: (^(v100)), F1: fi37, F2: 2}
                                                                                                                                                                                                                fmt.Println(v4, fi23, v165)
                                                                                                                                                                                                                var v299 uint32 = 763
                                                                                                                                                                                                                _ = v298
                                                                                                                                                                                                                _ = v299
                                                                                                                                                                                                            } else {
                                                                                                                                                                                                                v300 := GenS26{F0: v2, F1: v2, F2: v228}
                                                                                                                                                                                                                var v301 bool = v246
                                                                                                                                                                                                                _ = v300
                                                                                                                                                                                                                _ = v301
                                                                                                                                                                                                            }
                                                                                                                                                                                                        }
                                                                                                                                                                                                    }
                                                                                                                                                                                                    _ = v246
                                                                                                                                                                                                }
                                                                                                                                                                                                v53 += fi37
                                                                                                                                                                                            } else {
                                                                                                                                                                                                fmt.Println(v99, fi23)
                                                                                                                                                                                            }
                                                                                                                                                                                            v100 = -1
                                                                                                                                                                                            fmt.Println(v165)
                                                                                                                                                                                            if v228 {
                                                                                                                                                                                                if ((v165 + ("a" + "")) <= v165) {
                                                                                                                                                                                                    var v302 uint64 = 1
                                                                                                                                                                                                    v302 = v302
                                                                                                                                                                                                    _ = v302
                                                                                                                                                                                                }
                                                                                                                                                                                                fmt.Println(v203, v226)
                                                                                                                                                                                            }
                                                                                                                                                                                            _ = v234
                                                                                                                                                                                        }
                                                                                                                                                                                        v303 := GenS18{F0: (v100 ^ v100), F1: uint64(v54)}
                                                                                                                                                                                        fmt.Println(v99, v165, v100)
                                                                                                                                                                                        _ = v303
                                                                                                                                                                                    } else {
                                                                                                                                                                                        v228 = v203
                                                                                                                                                                                        v54 = 2
                                                                                                                                                                                    }
                                                                                                                                                                                }
                                                                                                                                                                            case "":
                                                                                                                                                                                var v304 int16 = (v2 ^ v2)
                                                                                                                                                                                if v228 {
                                                                                                                                                                                    v53 += fi37
                                                                                                                                                                                } else {
                                                                                                                                                                                    if len(v1) > 0 {
                                                                                                                                                                                        v1[0] = 2147483648
                                                                                                                                                                                    }
                                                                                                                                                                                }
                                                                                                                                                                                v226 = ((0.1 * (-2.25 - 0.0)) - 1.5)
                                                                                                                                                                                var v305 int8 = (v100 % ((v100) | 1))
                                                                                                                                                                                _ = v304
                                                                                                                                                                                _ = v305
                                                                                                                                                                            default:
                                                                                                                                                                                v306 := GenS27{F0: int32(fi23), F1: 4294967295}
                                                                                                                                                                                _ = v306
                                                                                                                                                                            }
                                                                                                                                                                        }
                                                                                                                                                                    } else {
                                                                                                                                                                        switch ("foo bar" + v165) {
                                                                                                                                                                        case "\n\t":
                                                                                                                                                                            var v307 int16 = v2
                                                                                                                                                                            _ = v307
                                                                                                                                                                        case "foo bar":
                                                                                                                                                                            if v203 {
                                                                                                                                                                                var v308 float64 = v226
                                                                                                                                                                                var v309 int64 = -500
                                                                                                                                                                                _ = v308
                                                                                                                                                                                _ = v309
                                                                                                                                                                            } else {
                                                                                                                                                                                v2 = (v2 & v2)
                                                                                                                                                                                var v310 int16 = v2
                                                                                                                                                                                v54 -= (v99 << ((4294967295) & 63))
                                                                                                                                                                                _ = v310
                                                                                                                                                                            }
                                                                                                                                                                        case "hello":
                                                                                                                                                                            var v311 float64 = float64(float32(1.5))
                                                                                                                                                                            _ = v311
                                                                                                                                                                        }
                                                                                                                                                                        var v312 int8 = (v100 % ((v100) | 1))
                                                                                                                                                                        _ = v312
                                                                                                                                                                    }
                                                                                                                                                                    if (uint32(v54) != 2147483648) {
                                                                                                                                                                        v165 = v165
                                                                                                                                                                        v3[65535] = 139
                                                                                                                                                                        v313 := v3[4294967295]
                                                                                                                                                                        _ = v313
                                                                                                                                                                    }
                                                                                                                                                                    if v228 {
                                                                                                                                                                        if (!v228) {
                                                                                                                                                                            fmt.Println(fi23, v2, v99)
                                                                                                                                                                        } else {
                                                                                                                                                                            _ = len(v3)
                                                                                                                                                                        }
                                                                                                                                                                        if false {
                                                                                                                                                                            var v314 int16 = v2
                                                                                                                                                                            fmt.Println(v226, v226, v228)
                                                                                                                                                                            switch v165 {
                                                                                                                                                                            case "hello":
                                                                                                                                                                                fmt.Println(v228)
                                                                                                                                                                            default:
                                                                                                                                                                                v315 := v3[65535]
                                                                                                                                                                                _ = v315
                                                                                                                                                                            }
                                                                                                                                                                            _ = v314
                                                                                                                                                                        }
                                                                                                                                                                        v228 = false
                                                                                                                                                                        v316 := GenS23{F0: v53}
                                                                                                                                                                        if true {
                                                                                                                                                                            v317 := GenS8{F0: 261, F1: 42, F2: 2147483647}
                                                                                                                                                                            v318 := GenS5{F0: (v228 && (false || true))}
                                                                                                                                                                            _ = v317
                                                                                                                                                                            _ = v318
                                                                                                                                                                        } else {
                                                                                                                                                                            v319 := GenS28{F0: 2}
                                                                                                                                                                            _ = v319
                                                                                                                                                                        }
                                                                                                                                                                        v320 := GenS29{F0: 2, F1: 922}
                                                                                                                                                                        v2 += -32768
                                                                                                                                                                        _ = v316
                                                                                                                                                                        _ = v320
                                                                                                                                                                    } else {
                                                                                                                                                                        v321 := GenS15{F0: -9223372036854775808, F1: (true && v203)}
                                                                                                                                                                        var v322 uint64 = 898
                                                                                                                                                                        v228 = (-776 >= 0)
                                                                                                                                                                        _ = v321
                                                                                                                                                                        _ = v322
                                                                                                                                                                    }
                                                                                                                                                                    v323 := GenS10{F0: 1, F1: 1, F2: 4294967295}
                                                                                                                                                                    _ = v323
                                                                                                                                                                }
                                                                                                                                                                v226 -= (((v226 / v226) / float64((0.0 + (0.1 * float32(((v226 - float64(v4)) * v226)))))) - v226)
                                                                                                                                                                var v324 int32 = 0
                                                                                                                                                                v100 += -1
                                                                                                                                                                _ = v228
                                                                                                                                                                _ = v324
                                                                                                                                                            }
                                                                                                                                                            if (-509 != 0) {
                                                                                                                                                                if (1e10 < float64(v4)) {
                                                                                                                                                                    var v325 uint32 = 42
                                                                                                                                                                    fmt.Println(v53, v226)
                                                                                                                                                                    _ = v325
                                                                                                                                                                }
                                                                                                                                                                v2 -= (v2 & v2)
                                                                                                                                                            }
                                                                                                                                                            v2 += 0
                                                                                                                                                        }
                                                                                                                                                        _ = v226
                                                                                                                                                    } else {
                                                                                                                                                        switch v165 {
                                                                                                                                                        case "\n\t":
                                                                                                                                                            if (!v203) {
                                                                                                                                                                var v326 uint64 = 4294967295
                                                                                                                                                                fmt.Println(v100)
                                                                                                                                                                v203 = v203
                                                                                                                                                                _ = v326
                                                                                                                                                            } else {
                                                                                                                                                                if len(v1) > 0 {
                                                                                                                                                                    v1[0] = 0
                                                                                                                                                                }
                                                                                                                                                                v327 := GenS2{F0: v2, F1: (("a" + (v165 + "foo bar")) + v165)}
                                                                                                                                                                _ = v327
                                                                                                                                                            }
                                                                                                                                                        case "foo bar":
                                                                                                                                                            v328 := GenS30{F0: 1, F1: (v54 * v54), F2: 255}
                                                                                                                                                            _ = v328
                                                                                                                                                        case "":
                                                                                                                                                            var v329 uint32 = 2
                                                                                                                                                            _ = v329
                                                                                                                                                        default:
                                                                                                                                                            fmt.Println(v2)
                                                                                                                                                        }
                                                                                                                                                        if true {
                                                                                                                                                            v3[65535] = 2
                                                                                                                                                        } else {
                                                                                                                                                            if (42 >= 0) {
                                                                                                                                                                v100 += (-(v100))
                                                                                                                                                                if len(v1) > 0 {
                                                                                                                                                                    v1[0] = 0
                                                                                                                                                                }
                                                                                                                                                                v203 = ((v99 < v99) || v203)
                                                                                                                                                            }
                                                                                                                                                            v2 += 1
                                                                                                                                                        }
                                                                                                                                                        switch (v165 + v165) {
                                                                                                                                                        case "foo bar":
                                                                                                                                                            if len(v1) > 0 {
                                                                                                                                                                v1[0] = 4294967295
                                                                                                                                                            }
                                                                                                                                                        case "\n\t":
                                                                                                                                                            v1 = append(v1, 2)
                                                                                                                                                        }
                                                                                                                                                    }
                                                                                                                                                }
                                                                                                                                            }
                                                                                                                                            if (-2147483648 > 42) {
                                                                                                                                                v99 += v54
                                                                                                                                                v4 += v4
                                                                                                                                            } else {
                                                                                                                                                v53 -= 306
                                                                                                                                                var v330 int8 = 100
                                                                                                                                                switch "\n\t" {
                                                                                                                                                case "foo bar":
                                                                                                                                                    if true {
                                                                                                                                                        v100 = (v330 + v100)
                                                                                                                                                    }
                                                                                                                                                case "\n\t":
                                                                                                                                                    fmt.Println(v100, v99)
                                                                                                                                                case "a":
                                                                                                                                                    var v331 bool = false
                                                                                                                                                    _ = v331
                                                                                                                                                }
                                                                                                                                                if (42 > 0) {
                                                                                                                                                    fmt.Println(v4, v203)
                                                                                                                                                    var v332 float64 = (0.1 / (0.0 + 1e10))
                                                                                                                                                    var v333 int8 = (^(v100))
                                                                                                                                                    _ = v332
                                                                                                                                                    _ = v333
                                                                                                                                                } else {
                                                                                                                                                    v99 += v54
                                                                                                                                                }
                                                                                                                                                _ = v330
                                                                                                                                            }
                                                                                                                                            _ = v203
                                                                                                                                        } else {
                                                                                                                                            var v334 int8 = (v100 * (^(v100)))
                                                                                                                                            var v335 int32 = 0
                                                                                                                                            v336 := v3[2147483647]
                                                                                                                                            _ = v334
                                                                                                                                            _ = v335
                                                                                                                                            _ = v336
                                                                                                                                        }
                                                                                                                                        _ = len(v1)
                                                                                                                                    }
                                                                                                                                }
                                                                                                                            }
                                                                                                                        }
                                                                                                                    default:
                                                                                                                        v100 = -128
                                                                                                                    }
                                                                                                                    var v337 int8 = (v100 << ((779) & 63))
                                                                                                                    _ = len(v3)
                                                                                                                    _ = v337
                                                                                                                case 412:
                                                                                                                    v338 := GenS31{F0: 128}
                                                                                                                    _ = v338
                                                                                                                default:
                                                                                                                    switch (-(fi37)) {
                                                                                                                    case -2147483648:
                                                                                                                        if (!true) {
                                                                                                                            switch v165 {
                                                                                                                            case "":
                                                                                                                                v99 -= v99
                                                                                                                            case "\n\t":
                                                                                                                                var v339 uint32 = 2147483647
                                                                                                                                _ = v339
                                                                                                                            default:
                                                                                                                                var v340 uint64 = 0
                                                                                                                                _ = v340
                                                                                                                            }
                                                                                                                            var v341 uint = 129
                                                                                                                            _ = v341
                                                                                                                        }
                                                                                                                    case -1:
                                                                                                                        var v342 int32 = -1
                                                                                                                        if false {
                                                                                                                            if (false || (4294967295 > 4294967295)) {
                                                                                                                                var v343 uint8 = 5
                                                                                                                                _ = v343
                                                                                                                            }
                                                                                                                            v344 := GenS32{F0: (-(fi37)), F1: 830, F2: uint16(v99)}
                                                                                                                            switch v165 {
                                                                                                                            case "hello":
                                                                                                                                var v345 string = v165
                                                                                                                                _ = v345
                                                                                                                            case "a":
                                                                                                                                if false {
                                                                                                                                    if true {
                                                                                                                                        v4 = v4
                                                                                                                                        if (1 >= 0) {
                                                                                                                                            if ((true || false) && true) {
                                                                                                                                                switch (^(fi23)) {
                                                                                                                                                case 32767:
                                                                                                                                                    v2 = (v2 ^ int16(v100))
                                                                                                                                                case 2147483647:
                                                                                                                                                    _ = len(v3)
                                                                                                                                                case 192:
                                                                                                                                                    v342 += (v342 * 1)
                                                                                                                                                }
                                                                                                                                                v53 = (fi23 / ((fi37) | 1))
                                                                                                                                                v3[598] = 2147483647
                                                                                                                                                v4 -= ((((v4 / v4) - 1.5) / ((float32(1e10) / -2.25) * 0.1)) * v4)
                                                                                                                                                v342 -= (v342 << ((0) & 63))
                                                                                                                                                v342 = (-(v342))
                                                                                                                                                if (-1 == -9223372036854775808) {
                                                                                                                                                    if (false && ((fi23 / ((-208) | 1)) <= (fi23 - fi23))) {
                                                                                                                                                        if true {
                                                                                                                                                            var v346 int8 = v100
                                                                                                                                                            v347 := GenS17{F0: v165, F1: (-(v342))}
                                                                                                                                                            _ = v346
                                                                                                                                                            _ = v347
                                                                                                                                                        } else {
                                                                                                                                                            v54 = 2
                                                                                                                                                            _ = len(v1)
                                                                                                                                                            var v348 uint32 = 2147483647
                                                                                                                                                            if len(v1) > 0 {
                                                                                                                                                                v1[0] = (v348 >> ((4294967295) & 63))
                                                                                                                                                            }
                                                                                                                                                            _ = v348
                                                                                                                                                        }
                                                                                                                                                        if false {
                                                                                                                                                            var v349 int32 = (^(v342))
                                                                                                                                                            _ = v349
                                                                                                                                                        }
                                                                                                                                                        v350 := GenS33{F0: (v54 >> ((65535) & 63)), F1: 42, F2: int64(v342)}
                                                                                                                                                        _ = v350
                                                                                                                                                    } else {
                                                                                                                                                        var v351 bool = true
                                                                                                                                                        v352 := GenS34{F0: 471, F1: v165, F2: 4294967295}
                                                                                                                                                        var v353 uint32 = 922
                                                                                                                                                        _ = v351
                                                                                                                                                        _ = v352
                                                                                                                                                        _ = v353
                                                                                                                                                    }
                                                                                                                                                    v354 := GenS23{F0: v53}
                                                                                                                                                    _ = v354
                                                                                                                                                } else {
                                                                                                                                                    if true {
                                                                                                                                                        if true {
                                                                                                                                                            _ = len(v1)
                                                                                                                                                        }
                                                                                                                                                        fmt.Println(v99)
                                                                                                                                                    } else {
                                                                                                                                                        if true {
                                                                                                                                                            fmt.Println(v342, v99, v54)
                                                                                                                                                        }
                                                                                                                                                        if false {
                                                                                                                                                            v355 := GenS35{F0: 288, F1: v54, F2: "\n\t"}
                                                                                                                                                            fmt.Println(v342, v99, v2)
                                                                                                                                                            var v356 uint64 = 208
                                                                                                                                                            _ = v355
                                                                                                                                                            _ = v356
                                                                                                                                                        } else {
                                                                                                                                                            if false {
                                                                                                                                                                if ("" < "hello") {
                                                                                                                                                                    fmt.Println(v54)
                                                                                                                                                                    var v357 uint8 = (v99 | (v54 % (((^(v54))) | 1)))
                                                                                                                                                                    v358 := GenS36{F0: 42, F1: 134}
                                                                                                                                                                    v342 = v342
                                                                                                                                                                    _ = v357
                                                                                                                                                                    _ = v358
                                                                                                                                                                }
                                                                                                                                                                _ = len(v1)
                                                                                                                                                                var v359 int32 = (v342 % (((v342 | (v342 % (((v342 - v342)) | 1)))) | 1))
                                                                                                                                                                _ = v359
                                                                                                                                                            }
                                                                                                                                                            if true {
                                                                                                                                                                if len(v1) > 0 {
                                                                                                                                                                    v1[0] = 65535
                                                                                                                                                                }
                                                                                                                                                                v360 := GenS1{F0: int16(v342), F1: 65535, F2: (v100 + v100)}
                                                                                                                                                                var v361 int32 = (^(v342))
                                                                                                                                                                _ = v360
                                                                                                                                                                _ = v361
                                                                                                                                                            } else {
                                                                                                                                                                var v362 float64 = -2.25
                                                                                                                                                                var v363 uint32 = 1
                                                                                                                                                                _ = v362
                                                                                                                                                                _ = v363
                                                                                                                                                            }
                                                                                                                                                        }
                                                                                                                                                        fmt.Println(v99, v165, fi37)
                                                                                                                                                    }
                                                                                                                                                    if len(v1) >= 2 {
                                                                                                                                                        v1 = v1[1:len(v1)]
                                                                                                                                                    }
                                                                                                                                                    var v364 bool = false
                                                                                                                                                    _ = v364
                                                                                                                                                }
                                                                                                                                            } else {
                                                                                                                                                var v365 uint = 180
                                                                                                                                                _ = v365
                                                                                                                                            }
                                                                                                                                        }
                                                                                                                                        var v366 uint32 = 4294967295
                                                                                                                                        _ = v366
                                                                                                                                    } else {
                                                                                                                                        var v367 int32 = v342
                                                                                                                                        var v368 int32 = (-(v342))
                                                                                                                                        if (!(1 < 2147483647)) {
                                                                                                                                            var v369 int8 = (v100 | 127)
                                                                                                                                            if (!(!(float64(v4) >= -2.25))) {
                                                                                                                                                if (!("hello" <= ((v165 + (v165 + "hello")) + (v165 + "foo bar")))) {
                                                                                                                                                    v367 += (^(v342))
                                                                                                                                                    delete(v3, 1)
                                                                                                                                                }
                                                                                                                                            }
                                                                                                                                            v370 := GenS29{F0: 18446744073709551615, F1: 4294967295}
                                                                                                                                            _ = v369
                                                                                                                                            _ = v370
                                                                                                                                        } else {
                                                                                                                                            fmt.Println(v4, v53, v2)
                                                                                                                                            switch v165 {
                                                                                                                                            case "hello":
                                                                                                                                                switch (v165 + (v165 + v165)) {
                                                                                                                                                case "a":
                                                                                                                                                    if len(v1) > 0 {
                                                                                                                                                        v1[0] = 65535
                                                                                                                                                    }
                                                                                                                                                case "\n\t":
                                                                                                                                                    fmt.Println(fi23, fi37)
                                                                                                                                                case "hello":
                                                                                                                                                    switch (v165 + (((((v165 + ("\n\t" + (v165 + ((v165 + "") + v165)))) + "foo bar") + v165) + "foo bar") + v165)) {
                                                                                                                                                    case "hello":
                                                                                                                                                        v371 := GenS25{F0: 212, F1: -345, F2: 1}
                                                                                                                                                        _ = v371
                                                                                                                                                    }
                                                                                                                                                default:
                                                                                                                                                    v54 += v99
                                                                                                                                                }
                                                                                                                                            case "foo bar":
                                                                                                                                                var v372 int16 = -1
                                                                                                                                                _ = v372
                                                                                                                                            default:
                                                                                                                                                var v373 uint = 103
                                                                                                                                                _ = v373
                                                                                                                                            }
                                                                                                                                            switch 127 {
                                                                                                                                            case 150:
                                                                                                                                                if (!(-2.25 == 0.1)) {
                                                                                                                                                    if (18446744073709551615 >= 18446744073709551615) {
                                                                                                                                                        if (65535 != 65535) {
                                                                                                                                                            var v374 uint8 = (v54 | v54)
                                                                                                                                                            _ = v374
                                                                                                                                                        } else {
                                                                                                                                                            var v375 int16 = v2
                                                                                                                                                            var v376 int16 = (-(v2))
                                                                                                                                                            var v377 string = v165
                                                                                                                                                            var v378 int8 = v100
                                                                                                                                                            _ = v375
                                                                                                                                                            _ = v376
                                                                                                                                                            _ = v377
                                                                                                                                                            _ = v378
                                                                                                                                                        }
                                                                                                                                                        var v379 int16 = v2
                                                                                                                                                        if len(v1) >= 2 {
                                                                                                                                                            v1 = v1[1:len(v1)]
                                                                                                                                                        }
                                                                                                                                                        _ = v379
                                                                                                                                                    }
                                                                                                                                                    if len(v1) > 0 {
                                                                                                                                                        v1[0] = 65535
                                                                                                                                                    }
                                                                                                                                                    v54 = v99
                                                                                                                                                }
                                                                                                                                            case -9223372036854775808:
                                                                                                                                                var v380 int32 = 53
                                                                                                                                                _ = v380
                                                                                                                                            case 32768:
                                                                                                                                                if (v54 >= v99) {
                                                                                                                                                    fmt.Println(v165, v165, v368)
                                                                                                                                                    var v381 uint = 1
                                                                                                                                                    _ = v381
                                                                                                                                                } else {
                                                                                                                                                    v382 := GenS37{F0: 2147483647}
                                                                                                                                                    _ = v382
                                                                                                                                                }
                                                                                                                                            default:
                                                                                                                                                var v383 float32 = v4
                                                                                                                                                _ = v383
                                                                                                                                            }
                                                                                                                                        }
                                                                                                                                        _ = v367
                                                                                                                                        _ = v368
                                                                                                                                    }
                                                                                                                                    var v384 uint16 = 0
                                                                                                                                    v385 := GenS29{F0: 260, F1: 34}
                                                                                                                                    _ = v384
                                                                                                                                    _ = v385
                                                                                                                                }
                                                                                                                            case "\n\t":
                                                                                                                                v54 = v99
                                                                                                                            }
                                                                                                                            _ = v344
                                                                                                                        } else {
                                                                                                                            if len(v1) >= 2 {
                                                                                                                                v1 = v1[1:len(v1)]
                                                                                                                            }
                                                                                                                            fmt.Println(v165, v99, fi37)
                                                                                                                        }
                                                                                                                        v99 -= (v54 & 1)
                                                                                                                        _ = v342
                                                                                                                    case 466:
                                                                                                                        var v386 string = "a"
                                                                                                                        _ = v386
                                                                                                                    }
                                                                                                                }
                                                                                                                _ = v177
                                                                                                            }
                                                                                                            switch ((v165 + v165) + v165) {
                                                                                                            case "\n\t":
                                                                                                                v4 += (float32(0.1) / v4)
                                                                                                            case "":
                                                                                                                v54 += v54
                                                                                                            case "foo bar":
                                                                                                                v99 = 127
                                                                                                            }
                                                                                                            if len(v1) >= 2 {
                                                                                                                v1 = v1[1:len(v1)]
                                                                                                            }
                                                                                                        }
                                                                                                    } else {
                                                                                                        var v387 string = ("foo bar" + "\n\t")
                                                                                                        v99 += v54
                                                                                                        switch "hello" {
                                                                                                        case "foo bar":
                                                                                                            v165 = ""
                                                                                                        case "hello":
                                                                                                            v388 := GenS18{F0: (v100 & v100), F1: 914}
                                                                                                            _ = v388
                                                                                                        case "\n\t":
                                                                                                            if (-2.25 <= -2.25) {
                                                                                                                var v389 int = 32768
                                                                                                                var v390 float64 = 0.1
                                                                                                                _ = v389
                                                                                                                _ = v390
                                                                                                            }
                                                                                                        default:
                                                                                                            var v391 float32 = (v4 + (-2.25 - float32(float64(v4))))
                                                                                                            _ = v391
                                                                                                        }
                                                                                                        _ = v387
                                                                                                    }
                                                                                                }
                                                                                                fmt.Println(v165, v4)
                                                                                                if false {
                                                                                                    var v392 bool = ((v4 > (3.14159 + v4)) && (float32((0.0 + 3.14159)) == 0.0))
                                                                                                    v393 := GenS30{F0: -1, F1: v54, F2: 65535}
                                                                                                    v100 = (-(v100))
                                                                                                    _ = v392
                                                                                                    _ = v393
                                                                                                } else {
                                                                                                    v2 += (v2 * 1000)
                                                                                                    v394 := GenS23{F0: (-(v53))}
                                                                                                    if (v4 > v4) {
                                                                                                        v395 := GenS31{F0: (v99 >> ((525) & 63))}
                                                                                                        if (((v2 << ((18446744073709551615) & 63)) < (v2 / ((v2) | 1))) && (!true)) {
                                                                                                            v396 := GenS38{F0: 4294967295, F1: fi37}
                                                                                                            var v397 int16 = (v2 << ((0) & 63))
                                                                                                            v99 = v54
                                                                                                            var v398 uint16 = 1
                                                                                                            _ = v396
                                                                                                            _ = v397
                                                                                                            _ = v398
                                                                                                        }
                                                                                                        v399 := GenS2{F0: (v2 ^ (v2 - (v2 >> ((2) & 63)))), F1: (v165 + (v165 + v165))}
                                                                                                        _ = v395
                                                                                                        _ = v399
                                                                                                    }
                                                                                                    _ = v394
                                                                                                }
                                                                                            } else {
                                                                                                var v400 int32 = 291
                                                                                                v401 := GenS15{F0: 9223372036854775807, F1: (!(2147483647 <= -49))}
                                                                                                if ((v400 - (v400 ^ 42)) > (^(v400))) {
                                                                                                    var v402 string = "\n\t"
                                                                                                    var v403 uint32 = 226
                                                                                                    v54 += (v54 << ((0) & 63))
                                                                                                    _ = v402
                                                                                                    _ = v403
                                                                                                } else {
                                                                                                    fmt.Println(v2, v54)
                                                                                                    switch (^(fi37)) {
                                                                                                    case 710:
                                                                                                        fmt.Println(v4, v4, v2)
                                                                                                    }
                                                                                                }
                                                                                                if true {
                                                                                                    if ((fi23 ^ 2147483647) < v53) {
                                                                                                        v99 = (v99 >> ((65535) & 63))
                                                                                                    }
                                                                                                    var v404 float32 = float32((0.0 + 1.5))
                                                                                                    v3[1] = 42
                                                                                                    _ = v404
                                                                                                }
                                                                                                var v405 int32 = v400
                                                                                                _ = v400
                                                                                                _ = v401
                                                                                                _ = v405
                                                                                            }
                                                                                            fmt.Println(fi37, v54)
                                                                                            _ = v165
                                                                                        } else {
                                                                                            v54 += (v99 * 1)
                                                                                            v4 = v4
                                                                                            var v406 float32 = -2.25
                                                                                            _ = v406
                                                                                        }
                                                                                    } else {
                                                                                        if (1 <= 80) {
                                                                                            var v407 float64 = 0.1
                                                                                            switch 9223372036854775807 {
                                                                                            case 128:
                                                                                                v53 = (v53 >> ((341) & 63))
                                                                                            case 348:
                                                                                                var v408 uint32 = 621
                                                                                                _ = v408
                                                                                            case -760:
                                                                                                v53 = -632
                                                                                            }
                                                                                            switch (((("a" + "foo bar") + "\n\t") + "\n\t") + ("" + ("a" + "\n\t"))) {
                                                                                            case "hello":
                                                                                                delete(v3, 1)
                                                                                                if true {
                                                                                                    var v409 uint16 = 65535
                                                                                                    var v410 int64 = 557
                                                                                                    _ = v409
                                                                                                    _ = v410
                                                                                                }
                                                                                            default:
                                                                                                var v411 float64 = v407
                                                                                                _ = v411
                                                                                            }
                                                                                            _ = v407
                                                                                        }
                                                                                    }
                                                                                case -429:
                                                                                    var v412 uint64 = 18446744073709551615
                                                                                    _ = v412
                                                                                case 0:
                                                                                    var v413 int16 = 1
                                                                                    _ = v413
                                                                                default:
                                                                                    v3[uint32(v99)] = 468
                                                                                }
                                                                                if (4294967295 <= fi23) {
                                                                                    v1 = append(v1, 1)
                                                                                    v99 = (^(v99))
                                                                                } else {
                                                                                    v414 := GenS37{F0: 4294967295}
                                                                                    switch -1 {
                                                                                    case 4294967295:
                                                                                        var v415 uint = 255
                                                                                        _ = v415
                                                                                    case 2147483647:
                                                                                        if true {
                                                                                            v99 += (v99 / (((v99 * v54)) | 1))
                                                                                            v4 -= (1.5 / 0.1)
                                                                                            switch "" {
                                                                                            case "hello":
                                                                                                if true {
                                                                                                    var v416 float32 = (-2.25 - v4)
                                                                                                    switch fi23 {
                                                                                                    case 32768:
                                                                                                        v3[364] = 65535
                                                                                                    case 1:
                                                                                                        v417 := GenS27{F0: -1, F1: 42}
                                                                                                        if (((("hello" + "foo bar") + ("\n\t" + "hello")) + "hello") == "a") {
                                                                                                            _ = len(v3)
                                                                                                            var v418 bool = ((v100 - (^(v100))) >= 1)
                                                                                                            if v418 {
                                                                                                                _ = len(v1)
                                                                                                                fmt.Println(v54)
                                                                                                                switch "\n\t" {
                                                                                                                case "a":
                                                                                                                    fmt.Println(v418, v99)
                                                                                                                case "hello":
                                                                                                                    v54 -= 127
                                                                                                                case "\n\t":
                                                                                                                    fmt.Println(v4, fi23)
                                                                                                                }
                                                                                                                if len(v1) > 0 {
                                                                                                                    v1[0] = 2147483648
                                                                                                                }
                                                                                                            } else {
                                                                                                                switch "a" {
                                                                                                                case "a":
                                                                                                                    var v419 bool = (fi37 < -768)
                                                                                                                    _ = v419
                                                                                                                case "":
                                                                                                                    var v420 float64 = 0.0
                                                                                                                    _ = v420
                                                                                                                case "\n\t":
                                                                                                                    switch fi23 {
                                                                                                                    case 193:
                                                                                                                        v421 := GenS25{F0: 255, F1: 9223372036854775807, F2: int8(v2)}
                                                                                                                        _ = v421
                                                                                                                    case 0:
                                                                                                                        if (v418 && (v418 || v418)) {
                                                                                                                            var v422 uint16 = 612
                                                                                                                            _ = v422
                                                                                                                        } else {
                                                                                                                            var v423 uint8 = (^(v54))
                                                                                                                            var v424 int8 = -128
                                                                                                                            _ = v423
                                                                                                                            _ = v424
                                                                                                                        }
                                                                                                                        var v425 uint8 = (v54 << ((266) & 63))
                                                                                                                        v2 = 109
                                                                                                                        var v426 float64 = (1.5 * (3.14159 + 3.14159))
                                                                                                                        _ = v425
                                                                                                                        _ = v426
                                                                                                                    case 42:
                                                                                                                        var v427 string = ("a" + "")
                                                                                                                        _ = v427
                                                                                                                    }
                                                                                                                }
                                                                                                                switch ("a" + "a") {
                                                                                                                case "foo bar":
                                                                                                                    delete(v3, 1)
                                                                                                                    v100 = (v100 * (v100 + v100))
                                                                                                                    switch (("a" + (("hello" + (("\n\t" + ("hello" + ("\n\t" + ("a" + ("a" + "hello"))))) + ("foo bar" + "hello"))) + ("hello" + ((("\n\t" + "foo bar") + "foo bar") + ("" + ("" + "")))))) + "foo bar") {
                                                                                                                    case "foo bar":
                                                                                                                        v53 -= 42
                                                                                                                    case "hello":
                                                                                                                        v428 := GenS39{F0: 0}
                                                                                                                        _ = v428
                                                                                                                    default:
                                                                                                                        _ = len(v1)
                                                                                                                    }
                                                                                                                case "":
                                                                                                                    if true {
                                                                                                                        v100 += -128
                                                                                                                        var v429 float32 = 0.0
                                                                                                                        fmt.Println(v99)
                                                                                                                        _ = v429
                                                                                                                    } else {
                                                                                                                        v1 = append(v1, 2147483648)
                                                                                                                    }
                                                                                                                case "\n\t":
                                                                                                                    var v430 int32 = -1
                                                                                                                    _ = v430
                                                                                                                default:
                                                                                                                    fmt.Println(v53)
                                                                                                                }
                                                                                                            }
                                                                                                            _ = v418
                                                                                                        }
                                                                                                        fmt.Println(fi23, v4)
                                                                                                        var v431 uint32 = 4294967295
                                                                                                        _ = v417
                                                                                                        _ = v431
                                                                                                    default:
                                                                                                        switch "foo bar" {
                                                                                                        case "foo bar":
                                                                                                            if true {
                                                                                                                v99 -= (v99 + (v99 - v54))
                                                                                                                var v432 int64 = 9223372036854775807
                                                                                                                _ = v432
                                                                                                            } else {
                                                                                                                var v433 uint = 65535
                                                                                                                _ = v433
                                                                                                            }
                                                                                                        case "":
                                                                                                            fmt.Println(v53, fi23, v416)
                                                                                                        default:
                                                                                                            if ((^(v99)) < (v99 ^ (v54 % ((v99) | 1)))) {
                                                                                                                v434 := GenS15{F0: -1, F1: (!true)}
                                                                                                                _ = v434
                                                                                                            } else {
                                                                                                                v435 := GenS19{F0: 489, F1: 42, F2: (v100 + v100)}
                                                                                                                if len(v1) > 0 {
                                                                                                                    v1[0] = 79
                                                                                                                }
                                                                                                                if len(v1) > 0 {
                                                                                                                    v1[0] = 2147483647
                                                                                                                }
                                                                                                                _ = v435
                                                                                                            }
                                                                                                        }
                                                                                                    }
                                                                                                    _ = v416
                                                                                                }
                                                                                            default:
                                                                                                if (163 != 4294967295) {
                                                                                                    if ((v4 * v4) == (-2.25 - float32(1.5))) {
                                                                                                        if true {
                                                                                                            v53 += 351
                                                                                                            var v436 float64 = 0.1
                                                                                                            _ = v436
                                                                                                        }
                                                                                                    }
                                                                                                    v53 += fi23
                                                                                                    v437 := GenS12{F0: 410}
                                                                                                    _ = v437
                                                                                                } else {
                                                                                                    fmt.Println(v99, v2, v100)
                                                                                                }
                                                                                            }
                                                                                            v2 += v2
                                                                                            v438 := GenS15{F0: -22, F1: (-2147483648 < 2147483647)}
                                                                                            _ = v438
                                                                                        }
                                                                                    case -432:
                                                                                        fmt.Println(fi37, v100, fi23)
                                                                                    }
                                                                                    _ = v414
                                                                                }
                                                                                var v439 string = "\n\t"
                                                                                _ = v164
                                                                                _ = v439
                                                                            } else {
                                                                                if false {
                                                                                    var v440 uint32 = 2
                                                                                    _ = v440
                                                                                } else {
                                                                                    v54 = (v54 | v54)
                                                                                    switch "a" {
                                                                                    case "\n\t":
                                                                                        fmt.Println(fi23, v100, fi37)
                                                                                    case "a":
                                                                                        var v441 uint = 42
                                                                                        v100 = v100
                                                                                        _ = v441
                                                                                    case "foo bar":
                                                                                        v442 := GenS40{F0: v2, F1: 1000}
                                                                                        _ = v442
                                                                                    }
                                                                                }
                                                                                v54 -= (v99 ^ v99)
                                                                            }
                                                                        default:
                                                                            v53 = (v53 % (((fi37 ^ 9223372036854775807)) | 1))
                                                                        }
                                                                    }
                                                                    _ = v100
                                                                    _ = v101
                                                                }
                                                                fmt.Println(v53)
                                                                _ = len(v1)
                                                                fmt.Println(fi37, v54, fi37)
                                                                fmt.Println(v4, v53, v53)
                                                                _ = v99
                                                            }
                                                        default:
                                                            fmt.Println(v4, v4, v53)
                                                        }
                                                    }
                                                    _ = v54
                                                }
                                                _ = v53
                                            case 496:
                                                v443 := GenS1{F0: v2, F1: 65535, F2: 104}
                                                _ = v443
                                            default:
                                                var v444 int16 = (v2 + 1)
                                                _ = v444
                                            }
                                        case "foo bar":
                                            v2 = (v2 % ((v2) | 1))
                                        case "":
                                            v2 += (v2 ^ (v2 * v2))
                                        }
                                    }
                                    v2 = v2
                                    if false {
                                        v445 := GenS27{F0: 0, F1: 4294967295}
                                        var v446 uint = 86
                                        v446 = v446
                                        fmt.Println(v446, v4, fi37)
                                        _ = v445
                                        _ = v446
                                    } else {
                                        v2 -= 0
                                        v4 -= 0.0
                                        var v447 string = "hello"
                                        _ = v447
                                    }
                                default:
                                    v448 := GenS9{F0: 1, F1: fi37, F2: 131}
                                    _ = v448
                                }
                                var v449 float32 = 0.1
                                v450 := GenS21{F0: 505, F1: 55, F2: 201}
                                _ = v449
                                _ = v450
                            }
                        case 127:
                            var v451 int64 = -922
                            if ((255 <= 1) && (1 > 128)) {
                                if false {
                                    switch (-(fi23)) {
                                    case -929:
                                        var v452 int8 = -128
                                        _ = v452
                                    case -653:
                                        switch "hello" {
                                        case "":
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                        case "foo bar":
                                            v4 -= 3.14159
                                        case "a":
                                            fmt.Println(fi23)
                                        }
                                    case 0:
                                        v453 := GenS24{F0: 1, F1: (v2 & 1000), F2: 269}
                                        _ = v453
                                    default:
                                        v2 = -22
                                    }
                                }
                            } else {
                                switch "a" {
                                case "\n\t":
                                    _ = len(v3)
                                case "":
                                    var v454 int8 = 0
                                    _ = v454
                                default:
                                    switch "foo bar" {
                                    case "a":
                                        var v455 uint16 = 0
                                        _ = v455
                                    default:
                                        delete(v3, 2)
                                    }
                                }
                            }
                            fmt.Println(v4)
                            v3[734] = 65535
                            _ = v451
                        default:
                            if false {
                                if (!((-750 == -2147483648) && false)) {
                                    v4 += (-2.25 - (v4 - 0.0))
                                } else {
                                    var v456 int16 = v2
                                    _ = v456
                                }
                                var v457 uint32 = 4294967295
                                _ = v457
                            } else {
                                v4 += v4
                                v1 = append(v1, 2147483648)
                                v458 := GenS41{F0: v2, F1: -1}
                                _ = v458
                            }
                        }
                        _ = fi37
                    }
                } else {
                    if false {
                        if (!(0 <= 0)) {
                            var v459 int8 = 127
                            v460 := GenS42{F0: fi23, F1: 2147483648}
                            _ = v459
                            _ = v460
                        }
                        v461 := GenS43{F0: (false || (-9223372036854775808 < -1))}
                        _ = v461
                    }
                }
            }
        }
        _ = fi23
        _ = v24
    }
    v2 = (v2 - 1000)
    fmt.Println(v2, v2)
    fmt.Println(v4, v4)
    switch (("foo bar" + "") + "") {
    case "hello":
        if false {
            v2 += v2
            v462 := GenS44{F0: 2147483647, F1: true}
            _ = v462
        } else {
            var v463 uint8 = 96
            v4 = 0.0
            delete(v3, 232)
            _ = v463
        }
    case "a":
        for fi464 := 0; fi464 < 7; fi464++ {
            for fi465 := 0; fi465 < 3; fi465++ {
                v466 := GenS13{F0: -2147483648, F1: 1, F2: v2}
                var v467 uint16 = 92
                _ = fi465
                _ = v466
                _ = v467
            }
            for fi468 := 0; fi468 < 16; fi468++ {
                if false {
                    v3[0] = 750
                    v4 = 0.0
                }
                var v469 int = -281
                v470 := GenS45{F0: 221, F1: -49}
                var v471 uint32 = 965
                v472 := GenS26{F0: -1, F1: (v2 | v2), F2: (false || (true || ("" <= "a")))}
                var v473 bool = true
                switch fi468 {
                case 128:
                    v2 -= v2
                default:
                    var v474 float64 = (float64(3.14159) * 0.0)
                    _ = v474
                }
                var v475 float32 = v4
                _ = fi468
                _ = v469
                _ = v470
                _ = v471
                _ = v472
                _ = v473
                _ = v475
            }
            v476 := GenS46{F0: 42, F1: (((!true) || true) || false), F2: 4294967295}
            _ = fi464
            _ = v476
        }
    case "\n\t":
        if (!false) {
            v1 = append(v1, 1)
            fmt.Println(v4, v4)
        } else {
            var v477 int16 = (^(v2))
            v477 += (v477 ^ v2)
            _ = v477
        }
    }
    switch 0 {
    case 128:
        if true {
            v478 := GenS47{F0: 182}
            v4 = v4
            v4 += float32(3.14159)
            _ = v478
        }
    case -1:
        v479 := GenS44{F0: 361, F1: false}
        _ = v479
    case 2147483647:
        v480 := GenS48{F0: 2}
        _ = v480
    }
    for fi481 := 0; fi481 < 15; fi481++ {
        if false {
            fmt.Println(v2, v4)
            v1 = append(v1, 42)
        }
        var v482 uint = 255
        var v483 string = ("a" + "foo bar")
        _ = fi481
        _ = v482
        _ = v483
    }
    v2 = (v2 / (((^(v2))) | 1))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
}
