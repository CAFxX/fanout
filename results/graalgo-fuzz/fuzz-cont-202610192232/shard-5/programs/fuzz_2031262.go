package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int32; F1 int64 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int64; F1 uint32 }
type GenS6 struct { F0 uint16; F1 int8 }
type GenS7 struct { F0 int32; F1 uint; F2 int8 }
type GenS8 struct { F0 uint; F1 uint16; F2 int32 }
type GenS9 struct { F0 int; F1 uint64 }
type GenS10 struct { F0 string; F1 int8; F2 uint16 }
type GenS11 struct { F0 int8; F1 uint8; F2 int32 }
type GenS12 struct { F0 uint64; F1 int32 }
type GenS13 struct { F0 bool }
type GenS14 struct { F0 uint16 }
func genFn0() int16 {
    for fi6 := 0; fi6 < 20; fi6++ {
        var v7 int8 = 100
        v7 -= v7
        v8 := GenS1{F0: (("\n\t" + "a") + "a")}
        _ = fi6
        _ = v7
        _ = v8
    }
    var v9 int = 624
    v10 := GenS1{F0: ""}
    _ = v10
    _ = v9
    return 32767
}

func main() {
    var v1 int8 = 42
    var v2 int32 = 250
    v3 := []int{-194, -2147483648}
    var v4 int32 = (v2 / ((v2) | 1))
    v5 := []uint32{365}
    if false {
        if true {
            if ("a" <= "hello") {
                v4 += (v2 << ((1) & 63))
            }
            var v11 uint8 = 0
            var v12 float64 = (-2.25 - (0.1 / float64(float32(1.5))))
            _ = v11
            _ = v12
        }
        v13 := GenS2{F0: 42, F1: 1}
        _ = v13
    } else {
        var v14 float32 = 3.14159
        var v15 int8 = (v1 / ((-46) | 1))
        v16 := GenS3{F0: ((v14 + ((v14 + 0.1) - (v14 / v14))) <= v14)}
        _ = v14
        _ = v15
        _ = v16
    }
    var v17 uint = 2
    _ = len(v5)
    fmt.Println(v1)
    fmt.Println(v4)
    switch 601 {
    case -1:
        v18 := GenS4{F0: v4}
        _ = v18
    case 1:
        switch "" {
        case "hello":
            v4 = (v4 * v4)
        case "\n\t":
            v1 = v1
        }
    case 0:
        for fi19 := 0; fi19 < 8; fi19++ {
            v20 := GenS1{F0: ""}
            var v21 string = (("hello" + "\n\t") + ("" + ("a" + "hello")))
            _ = fi19
            _ = v20
            _ = v21
        }
    }
    if false {
        v22 := GenS2{F0: v2, F1: -9223372036854775808}
        for fi23 := 0; fi23 < 11; fi23++ {
            var v24 int64 = -9223372036854775808
            var v25 int32 = (v2 ^ (v4 + v2))
            _ = fi23
            _ = v24
            _ = v25
        }
        _ = v22
    } else {
        fmt.Println(v4, v4)
        var v26 uint32 = 0
        _ = v26
    }
    var v27 uint16 = 224
    for fi28 := 0; fi28 < 10; fi28++ {
        var v29 uint16 = (^(v27))
        if len(v5) > 0 {
            v5[0] = 822
        }
        switch (fi28 | fi28) {
        case 344:
            var v30 int = -913
            _ = v30
        case 4294967295:
            fmt.Println(v29)
        }
        _ = fi28
        _ = v29
    }
    if (genFn0() <= 0) {
        for fi31 := 0; fi31 < 19; fi31++ {
            switch (fi31 * (fi31 | (fi31 + fi31))) {
            case 127:
                if false {
                    switch "" {
                    case "hello":
                        v2 = (^(v2))
                    case "":
                        for fi32 := 0; fi32 < 1; fi32++ {
                            v2 += v4
                            _ = fi32
                        }
                    case "\n\t":
                        if false {
                            v17 = 42
                        } else {
                            fmt.Println(v2, v4, v4)
                            fmt.Println(v27, v27, v27)
                            v33 := GenS5{F0: -1, F1: 875}
                            _ = v33
                        }
                    default:
                        var v34 int32 = int32(v27)
                        _ = v34
                    }
                    if ((169 != 128) && (!false)) {
                        if false {
                            v1 -= (-(v1))
                            var v35 int16 = 42
                            _ = v35
                        } else {
                            if len(v5) > 0 {
                                v5[0] = 1
                            }
                            if false {
                                if true {
                                    switch fi31 {
                                    case 4294967295:
                                        if len(v5) >= 2 {
                                            v5 = v5[1:len(v5)]
                                        }
                                    case -122:
                                        if false {
                                            v1 -= 90
                                            switch 1 {
                                            case -50:
                                                var v36 float64 = (-2.25 * float64(-2.25))
                                                _ = v36
                                            }
                                            v1 += (-(v1))
                                        } else {
                                            _ = len(v3)
                                        }
                                    default:
                                        v5 = append(v5, 0)
                                    }
                                    v17 = v17
                                    for fi37 := 0; fi37 < 16; fi37++ {
                                        if false {
                                            v38 := GenS4{F0: (v2 >> (((v17 / ((469) | 1))) & 63))}
                                            var v39 uint16 = 42
                                            v5 = append(v5, 514)
                                            _ = v38
                                            _ = v39
                                        } else {
                                            if (!(!(3.14159 != 3.14159))) {
                                                fmt.Println(fi31)
                                                var v40 uint16 = (v27 >> (((^(v17))) & 63))
                                                var v41 int = int(v40)
                                                _ = v40
                                                _ = v41
                                            }
                                            var v42 int64 = -9223372036854775808
                                            if (!(!false)) {
                                                if (!(32767 > -549)) {
                                                    v43 := GenS3{F0: true}
                                                    _ = v43
                                                }
                                            }
                                            var v44 int64 = (-(v42))
                                            _ = v42
                                            _ = v44
                                        }
                                        _ = fi37
                                    }
                                } else {
                                    v27 += v27
                                }
                                var v45 float64 = 1e10
                                _ = v45
                            } else {
                                var v46 uint32 = 1
                                if true {
                                    if (!(fi31 == fi31)) {
                                        fmt.Println(v1, fi31, v2)
                                        v47 := GenS3{F0: ((true || true) && true)}
                                        var v48 int16 = 1
                                        _ = v47
                                        _ = v48
                                    } else {
                                        if true {
                                            v27 -= (v27 * (v27 * (v27 + 2)))
                                            fmt.Println(v17, fi31, v4)
                                            switch (-(fi31)) {
                                            case -352:
                                                _ = len(v5)
                                            }
                                        } else {
                                            fmt.Println(v2, v1, v46)
                                            v17 = v17
                                        }
                                        for fi49 := 0; fi49 < 2; fi49++ {
                                            if false {
                                                switch fi49 {
                                                case -2147483648:
                                                    if false {
                                                        fmt.Println(fi31, v1)
                                                        var v50 uint32 = v46
                                                        var v51 int32 = (v4 / (((v4 << ((v17) & 63))) | 1))
                                                        switch "hello" {
                                                        case "foo bar":
                                                            if len(v5) >= 2 {
                                                                v5 = v5[1:len(v5)]
                                                            }
                                                        case "\n\t":
                                                            var v52 int64 = int64(v46)
                                                            _ = v52
                                                        case "hello":
                                                            if false {
                                                                v1 -= (-(v1))
                                                            }
                                                        }
                                                        _ = v50
                                                        _ = v51
                                                    } else {
                                                        v46 = v46
                                                        if ("\n\t" >= ("\n\t" + "")) {
                                                            v46 += (v46 & (v46 / ((599) | 1)))
                                                            switch "" {
                                                            case "a":
                                                                v53 := GenS5{F0: 569, F1: (v46 * v46)}
                                                                var v54 uint = v17
                                                                _ = v53
                                                                _ = v54
                                                            case "\n\t":
                                                                fmt.Println(v27, v46)
                                                            case "hello":
                                                                switch "\n\t" {
                                                                case "a":
                                                                    fmt.Println(v2)
                                                                case "foo bar":
                                                                    _ = len(v5)
                                                                case "":
                                                                    var v55 int64 = 489
                                                                    _ = v55
                                                                }
                                                            }
                                                            if len(v3) >= 2 {
                                                                v3 = v3[1:len(v3)]
                                                            }
                                                        } else {
                                                            if ((157 >= 762) || true) {
                                                                var v56 uint64 = 169
                                                                _ = v56
                                                            }
                                                        }
                                                        switch "" {
                                                        case "hello":
                                                            v27 += (v27 - 2)
                                                            var v57 int32 = (v4 + (^(v2)))
                                                            v58 := GenS5{F0: -186, F1: uint32(v17)}
                                                            v59 := GenS4{F0: v4}
                                                            fmt.Println(fi31, fi31, v57)
                                                            _ = v57
                                                            _ = v58
                                                            _ = v59
                                                        case "a":
                                                            _ = len(v3)
                                                        case "":
                                                            var v60 uint32 = v46
                                                            _ = v60
                                                        }
                                                    }
                                                default:
                                                    _ = len(v5)
                                                }
                                                v61 := GenS6{F0: uint16(v17), F1: (v1 % (((^(v1))) | 1))}
                                                v62 := GenS3{F0: (!(v1 <= (-(v1))))}
                                                _ = v61
                                                _ = v62
                                            }
                                            var v63 bool = true
                                            var v64 int8 = v1
                                            _ = fi49
                                            _ = v63
                                            _ = v64
                                        }
                                    }
                                    var v65 string = ("foo bar" + ("hello" + ("foo bar" + "hello")))
                                    v66 := GenS6{F0: (v27 & (v27 >> ((v17) & 63))), F1: v1}
                                    _ = v65
                                    _ = v66
                                } else {
                                    var v67 int = (-(fi31))
                                    var v68 float64 = (0.0 + 0.0)
                                    if (1 <= 140) {
                                        v69 := GenS3{F0: (v46 == v46)}
                                        var v70 uint16 = v27
                                        v71 := GenS3{F0: false}
                                        _ = v69
                                        _ = v70
                                        _ = v71
                                    }
                                    _ = v67
                                    _ = v68
                                }
                                fmt.Println(fi31, fi31)
                                _ = v46
                            }
                            switch fi31 {
                            case -560:
                                switch -9223372036854775808 {
                                case 658:
                                    v72 := GenS6{F0: 2, F1: 1}
                                    _ = v72
                                case 1:
                                    var v73 int16 = int16(v1)
                                    _ = v73
                                case 94:
                                    if ((fi31 - fi31) >= 127) {
                                        _ = len(v3)
                                        for fi74 := 0; fi74 < 14; fi74++ {
                                            fmt.Println(fi74, v17, v4)
                                            v1 = v1
                                            v4 = -67
                                            _ = fi74
                                        }
                                    } else {
                                        v1 = v1
                                        if (!(0.0 == 0.1)) {
                                            for fi75 := 0; fi75 < 16; fi75++ {
                                                var v76 int8 = 100
                                                switch (("a" + ("hello" + "a")) + "") {
                                                case "\n\t":
                                                    if false {
                                                        switch fi31 {
                                                        case 32768:
                                                            fmt.Println(v17, v76)
                                                        }
                                                        v77 := GenS7{F0: v4, F1: (^(v17)), F2: -128}
                                                        _ = v77
                                                    }
                                                    v78 := GenS8{F0: (^(v17)), F1: (v27 | v27), F2: v4}
                                                    if len(v5) > 0 {
                                                        v5[0] = 2147483648
                                                    }
                                                    _ = v78
                                                default:
                                                    var v79 int32 = 42
                                                    _ = v79
                                                }
                                                _ = len(v5)
                                                _ = fi75
                                                _ = v76
                                            }
                                        } else {
                                            v1 += 100
                                            var v80 uint = (v17 >> ((2) & 63))
                                            _ = v80
                                        }
                                        v17 -= (v17 >> ((v17) & 63))
                                    }
                                default:
                                    for fi81 := 0; fi81 < 2; fi81++ {
                                        var v82 int32 = -1
                                        _ = fi81
                                        _ = v82
                                    }
                                }
                            default:
                                v2 = (v4 << ((v17) & 63))
                            }
                        }
                        var v83 int32 = v2
                        var v84 int8 = (^(v1))
                        _ = v83
                        _ = v84
                    } else {
                        if ((0.1 + ((1e10 / 1e10) / 1.0)) > (3.14159 / 1.5)) {
                            var v85 uint16 = 65535
                            v86 := GenS5{F0: int64(fi31), F1: 748}
                            v87 := GenS5{F0: -1, F1: 271}
                            _ = v85
                            _ = v86
                            _ = v87
                        } else {
                            v88 := GenS9{F0: -2147483648, F1: uint64(v27)}
                            _ = v88
                        }
                        _ = len(v3)
                        v89 := GenS10{F0: "a", F1: v1, F2: 118}
                        _ = v89
                    }
                    for fi90 := 0; fi90 < 11; fi90++ {
                        if false {
                            var v91 float32 = (1.5 - float32(1.5))
                            if false {
                                v92 := GenS11{F0: 1, F1: uint8(v27), F2: v2}
                                fmt.Println(fi90, v17, v17)
                                switch "\n\t" {
                                case "":
                                    switch "a" {
                                    case "a":
                                        var v93 float32 = 0.1
                                        var v94 float32 = float32(0.0)
                                        var v95 uint = (v17 & (v17 >> (((v17 >> ((v17) & 63))) & 63)))
                                        _ = v93
                                        _ = v94
                                        _ = v95
                                    case "foo bar":
                                        var v96 int8 = 0
                                        _ = v96
                                    case "\n\t":
                                        v27 -= v27
                                    }
                                case "foo bar":
                                    fmt.Println(fi31)
                                }
                                _ = v92
                            } else {
                                var v97 uint64 = 39
                                v97 = 0
                                _ = v97
                            }
                            if false {
                                v98 := GenS11{F0: (v1 & v1), F1: 124, F2: (v2 % (((v2 | v4)) | 1))}
                                fmt.Println(fi90)
                                if (("\n\t" + "hello") <= "") {
                                    var v99 uint16 = 255
                                    if true {
                                        if true {
                                            var v100 uint64 = uint64(v17)
                                            var v101 uint32 = 2147483648
                                            _ = v100
                                            _ = v101
                                        } else {
                                            v102 := GenS11{F0: -1, F1: 14, F2: (v2 * (^(v2)))}
                                            var v103 float32 = (0.1 + v91)
                                            v5 = append(v5, 377)
                                            _ = v102
                                            _ = v103
                                        }
                                        if (fi90 == (-(fi31))) {
                                            var v104 int16 = 32767
                                            if true {
                                                v104 -= 1
                                                v105 := GenS12{F0: 0, F1: v2}
                                                v106 := GenS5{F0: 0, F1: 0}
                                                _ = v105
                                                _ = v106
                                            }
                                            fmt.Println(v104, v104)
                                            _ = v104
                                        } else {
                                            v99 += 0
                                            switch fi31 {
                                            case 42:
                                                v91 = float32(float64(float32(float64(float32(float64(v91))))))
                                                var v107 string = ""
                                                _ = v107
                                            case -2147483648:
                                                v91 -= (0.1 / (1e10 * v91))
                                            default:
                                                var v108 int = 2147483647
                                                _ = v108
                                            }
                                        }
                                        switch (((("foo bar" + "foo bar") + "\n\t") + ("\n\t" + ("foo bar" + "hello"))) + "\n\t") {
                                        case "":
                                            switch fi31 {
                                            case -2147483648:
                                                if false {
                                                    v3 = append(v3, (^(fi31)))
                                                    var v109 string = "foo bar"
                                                    _ = v109
                                                } else {
                                                    v110 := GenS2{F0: (v2 ^ (v2 & int32(fi31))), F1: -514}
                                                    _ = v110
                                                }
                                            case 4294967295:
                                                var v111 int64 = 1
                                                _ = v111
                                            case 882:
                                                v91 += v91
                                            }
                                            v1 = (v1 % ((-75) | 1))
                                            var v112 int16 = int16(fi31)
                                            v5 = append(v5, 918)
                                            _ = v112
                                        }
                                    }
                                    switch ("a" + "a") {
                                    case "foo bar":
                                        switch 4294967295 {
                                        case 9223372036854775807:
                                            var v113 uint = v17
                                            if len(v5) > 0 {
                                                v5[0] = 4294967295
                                            }
                                            _ = v113
                                        default:
                                            v114 := GenS13{F0: false}
                                            _ = v114
                                        }
                                    case "":
                                        var v115 uint64 = 131
                                        _ = v115
                                    }
                                    _ = v99
                                }
                                fmt.Println(v17)
                                _ = v98
                            } else {
                                v1 += -1
                                switch "" {
                                case "hello":
                                    var v116 uint8 = 0
                                    _ = v116
                                case "":
                                    var v117 string = ("foo bar" + "\n\t")
                                    _ = v117
                                }
                                var v118 float32 = v91
                                _ = v118
                            }
                            _ = v91
                        }
                        v119 := GenS14{F0: 32767}
                        var v120 int8 = (v1 << ((v17) & 63))
                        var v121 int32 = (v2 | v4)
                        _ = fi90
                        _ = v119
                        _ = v120
                        _ = v121
                    }
                } else {
                    v2 = v4
                    var v122 bool = (float32(3.14159) != -2.25)
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    _ = v122
                }
            }
            var v123 uint32 = 4294967295
            var v124 uint64 = 1
            _ = fi31
            _ = v123
            _ = v124
        }
        var v125 uint64 = 18446744073709551615
        _ = v125
    }
    var v126 float64 = float64((3.14159 + float32(0.0)))
    if (!((3.14159 / (1e10 / 1.0)) <= ((3.14159 - 1e10) - 1e10))) {
        switch "foo bar" {
        case "\n\t":
            v127 := GenS14{F0: 0}
            _ = v127
        case "foo bar":
            var v128 uint32 = 0
            _ = v128
        case "":
            v129 := GenS1{F0: "a"}
            _ = v129
        }
        v4 = (v2 & (v4 + 42))
        if ("foo bar" != "foo bar") {
            if (v17 < (v17 << (((v17 * (v17 / (((v17 << ((v17) & 63))) | 1)))) & 63))) {
                if false {
                    v1 = (v1 | (v1 ^ (v1 + (v1 & v1))))
                } else {
                    v17 -= v17
                    if (!(!false)) {
                        var v130 bool = true
                        _ = v130
                    }
                    fmt.Println(v2, v2)
                }
                v131 := GenS4{F0: 0}
                var v132 uint32 = 2147483647
                _ = v131
                _ = v132
            }
            var v133 float32 = 3.14159
            _ = v133
        } else {
            v17 = (v17 & (^(v17)))
            fmt.Println(v4, v126)
            var v134 string = "hello"
            _ = v134
        }
    }
    for fi135 := 0; fi135 < 6; fi135++ {
        if (("hello" + "hello") == "\n\t") {
            v4 = (v4 ^ v2)
        } else {
            for fi136 := 0; fi136 < 14; fi136++ {
                v17 += (^(v17))
                if ((90 < 0) && true) {
                    if len(v3) > 0 {
                        v3[0] = fi136
                    }
                } else {
                    var v137 int64 = -1
                    v138 := GenS10{F0: "foo bar", F1: (v1 | (-(v1))), F2: 0}
                    var v139 bool = ((!(!((false && (!true)) || (!true)))) || (v1 <= 0))
                    _ = v137
                    _ = v138
                    _ = v139
                }
                _ = fi136
            }
            fmt.Println(v126)
        }
        _ = fi135
    }
    v126 -= v126
    v126 = v126
    fmt.Println("v1:", v1)
    fmt.Println("v126:", v126)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v126
    _ = v17
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
