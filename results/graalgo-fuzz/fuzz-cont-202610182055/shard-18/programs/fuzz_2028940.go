package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string; F2 int }
type GenS2 struct { F0 int64; F1 uint8 }
type GenS3 struct { F0 uint; F1 uint8; F2 int8 }
type GenS4 struct { F0 uint8; F1 int16 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 int16; F1 int32; F2 int }
type GenS7 struct { F0 uint8; F1 int16 }
type GenS8 struct { F0 uint32; F1 uint64; F2 string }
type GenS9 struct { F0 int64; F1 int32; F2 int8 }
type GenS10 struct { F0 uint8; F1 string }
type GenS11 struct { F0 uint16 }
type GenS12 struct { F0 uint8; F1 int16 }
type GenS13 struct { F0 int; F1 uint8; F2 uint8 }
type GenS14 struct { F0 int64 }
type GenS15 struct { F0 uint8; F1 int8; F2 int64 }
type GenS16 struct { F0 int8 }
type GenS17 struct { F0 string; F1 uint64; F2 bool }
func genFn0() int {
    v4 := []int{-1}
    var v5 float32 = (0.1 / 1e10)
    var v6 int8 = 44
    _ = v4
    _ = v5
    _ = v6
    return 32767
}

func main() {
    v1 := make(map[uint32]int)
    var v2 int64 = -1
    var v3 int8 = 127
    v7 := GenS1{F0: 255, F1: ((("a" + "\n\t") + "foo bar") + (((("a" + (("a" + ("" + (("foo bar" + ("hello" + "hello")) + "hello"))) + "")) + ("foo bar" + ("foo bar" + ("foo bar" + ("" + (("a" + "a") + "hello")))))) + ((((("foo bar" + (("" + (("hello" + "") + "")) + ((("" + "foo bar") + ("foo bar" + ((("hello" + "") + "") + "foo bar"))) + "foo bar"))) + (((("foo bar" + "hello") + "a") + "hello") + "a")) + "a") + "foo bar") + "foo bar")) + "a")), F2: 128}
    for fi8 := 0; fi8 < 11; fi8++ {
        for fi9 := 0; fi9 < 11; fi9++ {
            v10 := GenS2{F0: v2, F1: 2}
            if (223 <= 127) {
                v2 -= v2
                fmt.Println(v3, fi9, fi9)
                if false {
                    if true {
                        v11 := GenS1{F0: 255, F1: "\n\t", F2: (-(fi8))}
                        if ("foo bar" > "hello") {
                            if false {
                                v3 = (-(v3))
                                var v12 uint = 18446744073709551615
                                v13 := []int{fi9, (fi8 >> (((v12 % ((v12) | 1))) & 63))}
                                _ = v12
                                _ = v13
                            } else {
                                if true {
                                    if true {
                                        _ = len(v1)
                                        v3 -= (v3 >> ((2) & 63))
                                        v14 := GenS2{F0: 1, F1: 128}
                                        _ = v14
                                    }
                                    v1[747] = (fi9 >> ((42) & 63))
                                } else {
                                    var v15 int = fi8
                                    _ = v15
                                }
                            }
                        } else {
                            if ((-32768 > 1) || false) {
                                v2 = (v2 | (v2 * v2))
                            } else {
                                v16 := []int{(fi9 % ((genFn0()) | 1)), (fi9 % (((^(fi9))) | 1))}
                                _ = v16
                            }
                            v2 = v2
                        }
                        v17 := v1[1]
                        _ = v11
                        _ = v17
                    } else {
                        switch (fi9 % ((genFn0()) | 1)) {
                        case -9223372036854775808:
                            v3 = (v3 | v3)
                        }
                    }
                    v2 = (^(v2))
                    if ((420 <= 255) && (42 > 42)) {
                        switch -9223372036854775808 {
                        case 787:
                            v18 := []int{(fi9 ^ (fi8 - genFn0())), -2147483648, 4294967295}
                            _ = v18
                        default:
                            v2 = v2
                        }
                    }
                } else {
                    v19 := GenS2{F0: -9223372036854775808, F1: 42}
                    v3 += (v3 % (((v3 - 0)) | 1))
                    switch fi8 {
                    case 2147483647:
                        fmt.Println(fi8, v3, v3)
                    case 486:
                        var v20 uint8 = 32
                        _ = v20
                    case -174:
                        fmt.Println(fi8, v2, v2)
                    default:
                        fmt.Println(fi9, v3, fi8)
                    }
                    if true {
                        if false {
                            v3 = v3
                            v2 = v2
                            var v21 uint16 = 65535
                            _ = v21
                        } else {
                            switch fi9 {
                            case 2147483647:
                                switch "hello" {
                                case "foo bar":
                                    if false {
                                        switch "a" {
                                        case "a":
                                            if true {
                                                _ = len(v1)
                                                v2 = -9223372036854775808
                                                fmt.Println(v2, v3, fi9)
                                            }
                                        case "foo bar":
                                            if false {
                                                v2 += 278
                                                v3 += -126
                                            }
                                        }
                                        if false {
                                            v3 = v3
                                            delete(v1, 2147483648)
                                            fmt.Println(fi8, fi8, fi9)
                                        }
                                        switch "foo bar" {
                                        case "":
                                            switch (-(fi8)) {
                                            case 32768:
                                                if true {
                                                    delete(v1, 499)
                                                    var v22 int8 = v3
                                                    _ = v22
                                                } else {
                                                    v2 -= (-(v2))
                                                    fmt.Println(fi9, fi9)
                                                }
                                            case 9223372036854775807:
                                                _ = len(v1)
                                            }
                                        case "foo bar":
                                            var v23 uint64 = 1
                                            _ = v23
                                        case "a":
                                            var v24 int = (fi8 * fi9)
                                            _ = v24
                                        }
                                    } else {
                                        v25 := GenS1{F0: 261, F1: ("foo bar" + ("" + "")), F2: fi8}
                                        if ((!true) || ("a" != (("hello" + (((("foo bar" + "") + "") + "foo bar") + (("" + "") + "\n\t"))) + ("hello" + ("hello" + ("\n\t" + "\n\t")))))) {
                                            if false {
                                                var v26 uint8 = 42
                                                _ = v26
                                            }
                                            var v27 int32 = -1
                                            var v28 uint32 = 2147483647
                                            _ = v27
                                            _ = v28
                                        } else {
                                            fmt.Println(v3)
                                            fmt.Println(v2, v2)
                                        }
                                        fmt.Println(fi8)
                                        v29 := GenS3{F0: 2, F1: 127, F2: 127}
                                        v30 := GenS4{F0: 1, F1: 42}
                                        v3 = v3
                                        if false {
                                            v31 := []int{1, (fi8 >> ((65535) & 63))}
                                            if (1 >= (fi8 / ((1) | 1))) {
                                                v2 -= v2
                                                v2 = 1
                                                fmt.Println(fi9, fi8)
                                            } else {
                                                if ("hello" != "foo bar") {
                                                    var v32 int16 = -32768
                                                    v32 = (v32 % (((v32 ^ v32)) | 1))
                                                    if true {
                                                        if (!(0.1 <= 0.0)) {
                                                            var v33 int = int(v32)
                                                            v34 := GenS3{F0: 18446744073709551615, F1: 42, F2: v3}
                                                            var v35 bool = true
                                                            _ = v33
                                                            _ = v34
                                                            _ = v35
                                                        }
                                                        if false {
                                                            if (0.0 == 0.0) {
                                                                switch "" {
                                                                case "":
                                                                    v32 = (v32 / ((v32) | 1))
                                                                }
                                                                if true {
                                                                    if false {
                                                                        v2 -= v2
                                                                        v2 -= v2
                                                                        v3 = v3
                                                                    } else {
                                                                        v2 -= (v2 % ((v2) | 1))
                                                                    }
                                                                } else {
                                                                    fmt.Println(fi9, fi8)
                                                                }
                                                                v1[0] = (fi8 % ((fi8) | 1))
                                                            } else {
                                                                var v36 uint32 = 491
                                                                if (v2 != v2) {
                                                                    v37 := v1[(v36 - (v36 - (v36 / ((v36) | 1))))]
                                                                    _ = v37
                                                                } else {
                                                                    _ = len(v31)
                                                                    v32 = (v32 >> ((327) & 63))
                                                                }
                                                                switch 709 {
                                                                case 42:
                                                                    if len(v31) > 0 {
                                                                        v31[0] = fi8
                                                                    }
                                                                default:
                                                                    v38 := GenS2{F0: (v2 ^ -118), F1: 250}
                                                                    _ = v38
                                                                }
                                                                _ = v36
                                                            }
                                                            if len(v31) >= 2 {
                                                                v31 = v31[1:len(v31)]
                                                            }
                                                        }
                                                        var v39 int64 = v2
                                                        _ = v39
                                                    }
                                                    _ = v32
                                                } else {
                                                    var v40 float32 = float32(1e10)
                                                    _ = v40
                                                }
                                            }
                                            _ = v31
                                        } else {
                                            delete(v1, 42)
                                            var v41 int16 = -32768
                                            v42 := GenS3{F0: 65535, F1: 42, F2: (v3 % ((100) | 1))}
                                            _ = v41
                                            _ = v42
                                        }
                                        var v43 uint = 4294967295
                                        _ = v25
                                        _ = v29
                                        _ = v30
                                        _ = v43
                                    }
                                case "a":
                                    var v44 uint16 = 255
                                    _ = v44
                                default:
                                    if true {
                                        fmt.Println(v2, v3, v3)
                                        v2 += 42
                                        var v45 float64 = -2.25
                                        _ = v45
                                    } else {
                                        var v46 float32 = 0.0
                                        v46 += float32((float64(v46) / (((3.14159 - 0.1) * 3.14159) + (float64(v46) + float64(v46)))))
                                        v47 := []int{(^(fi9)), (fi8 % ((fi9) | 1)), fi9}
                                        _ = v46
                                        _ = v47
                                    }
                                }
                            }
                            fmt.Println(fi8, v2)
                            v48 := []int{fi9}
                            _ = v48
                        }
                    } else {
                        _ = len(v1)
                        v2 += (v2 & v2)
                        v3 = v3
                    }
                    _ = v19
                }
            }
            var v49 uint8 = 115
            _ = fi9
            _ = v10
            _ = v49
        }
        var v50 uint32 = 306
        if false {
            for fi51 := 0; fi51 < 11; fi51++ {
                var v52 int16 = -418
                if false {
                    if ((!false) || false) {
                        v52 += (v52 / ((-789) | 1))
                        v53 := GenS2{F0: v2, F1: 17}
                        switch int(v52) {
                        case 4294967295:
                            v3 -= (^(v3))
                        default:
                            v54 := GenS5{F0: ("" + "\n\t")}
                            _ = v54
                        }
                        _ = v53
                    } else {
                        v55 := GenS6{F0: v52, F1: 2147483647, F2: fi51}
                        fmt.Println(v2, v2)
                        v56 := GenS2{F0: v2, F1: 127}
                        v57 := GenS7{F0: 127, F1: (v52 - v52)}
                        _ = v55
                        _ = v56
                        _ = v57
                    }
                    var v58 uint16 = 576
                    v3 = v3
                    _ = v58
                }
                v59 := []int{fi51, (fi51 & fi8), (fi8 * fi8), genFn0()}
                if (0 > 2) {
                    v52 += (v52 >> ((18446744073709551615) & 63))
                    switch (fi51 >> ((2) & 63)) {
                    case -928:
                        switch "foo bar" {
                        case "a":
                            v52 -= -1
                        case "hello":
                            v2 = v2
                        case "":
                            v3 = v3
                        }
                    case 128:
                        v60 := GenS8{F0: 0, F1: 1, F2: ("hello" + "a")}
                        _ = v60
                    case 146:
                        fmt.Println(v52, v2, v50)
                    }
                }
                _ = fi51
                _ = v52
                _ = v59
            }
            var v61 int64 = v2
            v62 := []int{fi8, (-(fi8)), fi8, fi8}
            _ = v61
            _ = v62
        } else {
            for fi63 := 0; fi63 < 12; fi63++ {
                var v64 float32 = 1.5
                var v65 int = fi63
                _ = fi63
                _ = v64
                _ = v65
            }
            if (4294967295 < 366) {
                var v66 uint32 = (v50 - (v50 >> ((4294967295) & 63)))
                _ = v66
            }
            v50 = v50
        }
        _ = fi8
        _ = v50
    }
    for fi67 := 0; fi67 < 12; fi67++ {
        v68 := []int{-701, genFn0(), fi67}
        v69 := GenS9{F0: v2, F1: -1, F2: (v3 >> ((522) & 63))}
        _ = fi67
        _ = v68
        _ = v69
    }
    fmt.Println(v3)
    v3 -= 71
    v70 := GenS10{F0: 40, F1: "hello"}
    v2 -= (v2 << ((0) & 63))
    v3 += v3
    var v71 float32 = (1.5 * ((0.1 - ((1.5 - 0.1) / -2.25)) - (0.0 / 3.14159)))
    var v72 int64 = v2
    fmt.Println(v71, v71)
    v73 := GenS11{F0: 65535}
    v74 := []int{}
    fmt.Println(v2, v3, v3)
    switch 700 {
    case 0:
        if true {
            v75 := GenS11{F0: 515}
            _ = v75
        } else {
            switch -1 {
            case -1:
                if len(v74) > 0 {
                    v74[0] = 9223372036854775807
                }
            case -2147483648:
                if true {
                    switch 9223372036854775807 {
                    case 32768:
                        switch 128 {
                        case 993:
                            v3 = (v3 / ((v3) | 1))
                        case 1:
                            v76 := GenS12{F0: 239, F1: 32767}
                            _ = v76
                        case -215:
                            var v77 uint16 = 305
                            _ = v77
                        }
                    }
                    for fi78 := 0; fi78 < 1; fi78++ {
                        v71 += (v71 * v71)
                        switch "hello" {
                        case "":
                            fmt.Println(v71, v71, fi78)
                        case "a":
                            fmt.Println(v71, v3)
                        case "\n\t":
                            var v79 bool = false
                            _ = v79
                        }
                        if false {
                            v3 = v3
                        }
                        _ = fi78
                    }
                } else {
                    var v80 float32 = v71
                    switch -960 {
                    case 4294967295:
                        for fi81 := 0; fi81 < 8; fi81++ {
                            for fi82 := 0; fi82 < 14; fi82++ {
                                var v83 uint16 = 0
                                fmt.Println(fi81, v3)
                                _ = fi82
                                _ = v83
                            }
                            var v84 uint32 = 2147483648
                            var v85 uint64 = 0
                            _ = fi81
                            _ = v84
                            _ = v85
                        }
                    }
                    _ = v80
                }
            case 9223372036854775807:
                if true {
                    v72 -= (v72 ^ v72)
                    v72 += v72
                    if false {
                        if true {
                            var v86 int64 = 1
                            if (v71 != (1.5 - 3.14159)) {
                                v71 += (v71 * v71)
                                v3 += (v3 | (v3 >> ((255) & 63)))
                            }
                            _ = v86
                        } else {
                            fmt.Println(v71)
                            v3 = -1
                            switch genFn0() {
                            case 717:
                                v87 := GenS13{F0: -997, F1: 160, F2: 176}
                                _ = v87
                            }
                        }
                        fmt.Println(v72)
                        if (-153 <= 42) {
                            if true {
                                var v88 uint8 = 127
                                if len(v74) >= 2 {
                                    v74 = v74[1:len(v74)]
                                }
                                for fi89 := 0; fi89 < 6; fi89++ {
                                    var v90 int8 = v3
                                    v91 := GenS14{F0: (v2 / ((42) | 1))}
                                    _ = fi89
                                    _ = v90
                                    _ = v91
                                }
                                _ = v88
                            } else {
                                v72 = v2
                                for fi92 := 0; fi92 < 3; fi92++ {
                                    v71 = ((v71 / (-2.25 + v71)) * float32(0.1))
                                    _ = fi92
                                }
                                var v93 int = 32768
                                _ = v93
                            }
                            if (!true) {
                                v72 = 1
                                if true {
                                    v94 := GenS15{F0: 1, F1: v3, F2: v2}
                                    _ = v94
                                } else {
                                    switch (((("hello" + "\n\t") + "\n\t") + ("foo bar" + "\n\t")) + "a") {
                                    case "\n\t":
                                        if ((v3 < (-(v3))) || false) {
                                            if len(v74) > 0 {
                                                v74[0] = 32768
                                            }
                                            v95 := v1[2147483647]
                                            _ = v95
                                        }
                                    case "hello":
                                        v96 := GenS14{F0: v2}
                                        _ = v96
                                    default:
                                        if ((!true) || true) {
                                            for fi97 := 0; fi97 < 10; fi97++ {
                                                var v98 int16 = -1
                                                _ = fi97
                                                _ = v98
                                            }
                                        }
                                    }
                                }
                                if false {
                                    v99 := GenS16{F0: v3}
                                    fmt.Println(v71)
                                    var v100 uint16 = 129
                                    _ = v100
                                    _ = v99
                                }
                            }
                        }
                    }
                } else {
                    if false {
                        v71 += 3.14159
                    } else {
                        if (!false) {
                            v3 += v3
                            var v101 uint16 = 32767
                            _ = v101
                        } else {
                            v102 := GenS13{F0: -1, F1: 42, F2: 42}
                            if (false || false) {
                                v2 = -1
                            }
                            var v103 uint64 = 18446744073709551615
                            _ = v102
                            _ = v103
                        }
                        var v104 int8 = (-(v3))
                        _ = v104
                    }
                    for fi105 := 0; fi105 < 11; fi105++ {
                        var v106 uint64 = 18446744073709551615
                        if (!(127 < 255)) {
                            if true {
                                for fi107 := 0; fi107 < 4; fi107++ {
                                    fmt.Println(fi107)
                                    _ = fi107
                                }
                                if (((-(v2)) >= (v2 % (((v2 - 864)) | 1))) || (("" != "foo bar") && ((((v71 + v71) / ((v71 - ((v71 / (float32((float64(v71) - -2.25)) - ((-2.25 - (((0.0 / float32(-2.25)) * 0.1) * float32(float64((0.0 * float32(float64(((3.14159 - -2.25) / ((-2.25 + (float32(float64((float32((float64(3.14159) / ((0.1 * -2.25) - ((0.0 - float64(v71)) + 0.1)))) - (float32(-2.25) - v71)))) - (v71 + 3.14159))) + (v71 * (1.5 + 3.14159))))))))))) - float32(0.1)))) * v71)) + (v71 * v71))) / v71) > float32(1.5)))) {
                                    if (!(2 != 42)) {
                                        v72 -= v2
                                    }
                                    fmt.Println(fi105, v72)
                                }
                            }
                            v108 := GenS11{F0: 1}
                            for fi109 := 0; fi109 < 2; fi109++ {
                                if true {
                                    var v110 uint64 = (v106 & (v106 & (v106 * 545)))
                                    var v111 bool = (42 == 65535)
                                    switch genFn0() {
                                    case 32767:
                                        if true {
                                            if ("hello" >= ("foo bar" + (("" + "hello") + "hello"))) {
                                                v3 -= v3
                                                var v112 bool = (fi105 <= fi105)
                                                var v113 string = "a"
                                                _ = v112
                                                _ = v113
                                            }
                                        }
                                        v1[4294967295] = fi105
                                        var v114 bool = true
                                        _ = v114
                                    case 127:
                                        v72 = (v72 / (((v72 - (^(v72)))) | 1))
                                    case 9223372036854775807:
                                        switch "foo bar" {
                                        case "hello":
                                            v1[73] = (fi109 + 32767)
                                        case "":
                                            switch 42 {
                                            case -1:
                                                fmt.Println(v71, v3, v111)
                                            case 128:
                                                v106 += v110
                                            default:
                                                var v115 uint8 = 42
                                                _ = v115
                                            }
                                        default:
                                            var v116 string = "\n\t"
                                            _ = v116
                                        }
                                    }
                                    if (!(!(v111 || ((v106 - (v106 << ((1) & 63))) < 42)))) {
                                        if len(v74) > 0 {
                                            v74[0] = genFn0()
                                        }
                                        v117 := GenS17{F0: ("hello" + "foo bar"), F1: v110, F2: ((fi105 <= genFn0()) || (23 > 42))}
                                        _ = v117
                                    }
                                    if v111 {
                                        if v111 {
                                            if (4294967295 >= 1) {
                                                if (141 > 668) {
                                                    v106 = v110
                                                    v71 = (v71 + (3.14159 + 0.1))
                                                }
                                                var v118 int16 = 42
                                                var v119 int32 = 782
                                                v2 = (v2 | v72)
                                                v74 = append(v74, genFn0())
                                                _ = v118
                                                _ = v119
                                            }
                                        }
                                    }
                                    _ = v110
                                    _ = v111
                                }
                                if (-2147483648 >= 328) {
                                    delete(v1, 1)
                                    var v120 uint16 = 0
                                    _ = v120
                                }
                                fmt.Println(fi109)
                                _ = fi109
                            }
                            _ = v108
                        } else {
                            v74 = append(v74, (fi105 >> ((4294967295) & 63)))
                        }
                        _ = fi105
                        _ = v106
                    }
                    v121 := GenS4{F0: 42, F1: 1}
                    _ = v121
                }
            }
            var v122 uint = 822
            var v123 uint8 = 127
            _ = v122
            _ = v123
        }
    case -9223372036854775808:
        var v124 int16 = -364
        _ = v124
    default:
        var v125 int = 905
        _ = v125
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v71:", v71)
    fmt.Println("v72:", v72)
    fmt.Println("slicelen v74:", len(v74), v74)
    _ = v1
    _ = v2
    _ = v3
    _ = v7
    _ = v70
    _ = v71
    _ = v72
    _ = v73
    _ = v74
}
