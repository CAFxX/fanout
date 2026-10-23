package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int8; F1 uint }
type GenS4 struct { F0 uint16; F1 uint16 }
type GenS5 struct { F0 int8; F1 uint8; F2 string }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 uint64; F1 int16; F2 int32 }
type GenS8 struct { F0 int; F1 int64; F2 int32 }
type GenS9 struct { F0 uint16; F1 uint32; F2 uint32 }
type GenS10 struct { F0 bool }
type GenS11 struct { F0 uint }
type GenS12 struct { F0 string; F1 uint8; F2 bool }
type GenS13 struct { F0 uint32; F1 uint64; F2 int64 }
type GenS14 struct { F0 uint64; F1 int64; F2 uint32 }
func genFn0(p0 uint, p1 uint, p2 bool) uint32 {
    p0 -= 42
    v7 := GenS1{F0: 497}
    switch (("a" + "hello") + "hello") {
    case "\n\t":
        v8 := []int{128, -459}
        _ = v8
    case "a":
        for fi9 := 0; fi9 < 5; fi9++ {
            switch (("" + ("a" + ("hello" + ""))) + ("foo bar" + "")) {
            case "a":
                switch fi9 {
                case 32767:
                    var v10 int8 = -1
                    _ = v10
                case -9223372036854775808:
                    switch -2147483648 {
                    case 2147483647:
                        v11 := []int{128, 32768}
                        _ = v11
                    }
                default:
                    var v12 uint = 338
                    _ = v12
                }
            case "foo bar":
                var v13 int = fi9
                _ = v13
            case "hello":
                if (fi9 == fi9) {
                    v14 := []int{-9223372036854775808, (fi9 >> ((18446744073709551615) & 63)), fi9, (fi9 * (fi9 + (fi9 + (fi9 << ((p1) & 63)))))}
                    v15 := GenS1{F0: fi9}
                    _ = v14
                    _ = v15
                }
            }
            switch -692 {
            case 1:
                var v16 int = fi9
                _ = v16
            default:
                if (p2 || true) {
                    if (42 <= -1) {
                        v17 := []int{fi9, 127, fi9, fi9}
                        _ = v17
                    } else {
                        v18 := []int{fi9, fi9, (fi9 * 38)}
                        _ = v18
                    }
                    for fi19 := 0; fi19 < 8; fi19++ {
                        switch "foo bar" {
                        case "":
                            var v20 uint32 = 42
                            _ = v20
                        case "\n\t":
                            var v21 uint8 = 128
                            _ = v21
                        }
                        fmt.Println(fi19)
                        var v22 int16 = 42
                        _ = fi19
                        _ = v22
                    }
                } else {
                    switch int(p0) {
                    case 128:
                        fmt.Println(fi9)
                    case 622:
                        if p2 {
                            v23 := GenS2{F0: 42}
                            if ((!p2) || false) {
                                var v24 int8 = -1
                                v25 := []int{(^(fi9))}
                                v24 += (v24 ^ v24)
                                _ = v24
                                _ = v25
                            }
                            _ = v23
                        } else {
                            v26 := []int{fi9, 996, -894}
                            if (959 < 42) {
                                if len(v26) > 0 {
                                    v26[0] = (fi9 << ((p1) & 63))
                                }
                                var v27 int64 = -1
                                _ = v27
                            } else {
                                p1 += 881
                                v28 := GenS3{F0: 127, F1: p1}
                                v29 := GenS1{F0: -525}
                                _ = v28
                                _ = v29
                            }
                            switch (("\n\t" + "hello") + "") {
                            case "a":
                                p2 = (true || (float64((-2.25 / float32(1.5))) == 3.14159))
                            case "\n\t":
                                var v30 uint32 = 4294967295
                                _ = v30
                            case "foo bar":
                                var v31 uint64 = 4294967295
                                _ = v31
                            default:
                                var v32 uint32 = 0
                                _ = v32
                            }
                            _ = v26
                        }
                    case 32767:
                        fmt.Println(p0)
                    default:
                        p0 = 621
                    }
                }
            }
            _ = fi9
        }
    case "foo bar":
        if p2 {
            for fi33 := 0; fi33 < 10; fi33++ {
                for fi34 := 0; fi34 < 15; fi34++ {
                    fmt.Println(fi34)
                    v35 := []int{fi33, (fi33 / (((fi34 | fi33)) | 1))}
                    _ = fi34
                    _ = v35
                }
                _ = fi33
            }
            if false {
                switch 0 {
                case 127:
                    if (1 != 944) {
                        fmt.Println(p0)
                        fmt.Println(p0, p2)
                    }
                case -687:
                    switch 1 {
                    case 48:
                        p0 = 2
                    case 9223372036854775807:
                        p1 = 4294967295
                    }
                default:
                    v36 := GenS4{F0: 710, F1: 1}
                    _ = v36
                }
                v37 := GenS5{F0: 127, F1: 42, F2: (("\n\t" + (("" + (("" + "") + "hello")) + "")) + ("foo bar" + ("a" + "foo bar")))}
                p2 = (-9223372036854775808 != 1)
                _ = v37
            }
        }
    default:
        switch "hello" {
        case "\n\t":
            p2 = true
        case "":
            v38 := GenS5{F0: -1, F1: 1, F2: ("\n\t" + "foo bar")}
            _ = v38
        case "a":
            var v39 uint8 = 236
            _ = v39
        default:
            var v40 int32 = 1
            _ = v40
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 594
}
func genFn1(p0 int16, p1 int64, p2 int8) int32 {
    if (0.1 >= ((((((((3.14159 + 0.1) + 0.0) * float64(1e10)) * float64(3.14159)) - 0.0) + 3.14159) - float64(3.14159)) / 1.0)) {
        p2 = p2
        v41 := GenS3{F0: (p2 | (p2 | p2)), F1: 411}
        _ = v41
    } else {
        p2 -= p2
    }
    switch -473 {
    case 127:
        p0 = 32767
    case 9223372036854775807:
        switch ("" + ("foo bar" + "a")) {
        case "foo bar":
            p0 += (p0 & p0)
        case "\n\t":
            p2 = p2
        case "hello":
            fmt.Println(p2, p2, p2)
        default:
            p2 = 42
        }
    case -9223372036854775808:
        switch "" {
        case "\n\t":
            v42 := GenS3{F0: -83, F1: 2}
            _ = v42
        default:
            p0 -= 32767
        }
    default:
        p0 = (-(p0))
    }
    for fi43 := 0; fi43 < 6; fi43++ {
        var v44 uint = 65535
        var v45 uint8 = 42
        _ = fi43
        _ = v44
        _ = v45
    }
    _ = p0
    _ = p1
    _ = p2
    return -365
}

func main() {
    var v1 string = "foo bar"
    var v2 float32 = ((-2.25 * (0.1 - (-2.25 + 0.1))) + float32(float64(1.5)))
    var v3 bool = (float32(3.14159) > ((v2 * 0.0) * v2))
    v4 := make(map[uint32]string)
    var v5 uint16 = 1
    var v6 int8 = -128
    var v46 uint32 = genFn0(2, 42, (42 == 2))
    for fi47 := 0; fi47 < 10; fi47++ {
        var v48 int = (-(fi47))
        _ = fi47
        _ = v48
    }
    var v49 uint = 42
    v50 := GenS1{F0: -956}
    var v51 uint64 = 18446744073709551615
    for fi52 := 0; fi52 < 11; fi52++ {
        if true {
            if v3 {
                var v53 string = ("" + (v1 + "\n\t"))
                _ = v53
            } else {
                fmt.Println(v5)
                v5 += (^(v5))
                v46 = uint32(v49)
            }
            v1 = ((v1 + v1) + v1)
            v1 = (v1 + ("foo bar" + v1))
        }
        v46 -= (v46 & genFn0(v49, (v49 ^ (v49 & 571)), (!v3)))
        v54 := GenS1{F0: fi52}
        _ = fi52
        _ = v54
    }
    if (1.5 >= ((-2.25 - 1.5) / (v2 / 3.14159))) {
        var v55 uint = 42
        switch 9223372036854775807 {
        case -2147483648:
            var v56 string = (v1 + "foo bar")
            _ = v56
        case 42:
            for fi57 := 0; fi57 < 3; fi57++ {
                v58 := v4[485]
                _ = fi57
                _ = v58
            }
        }
        for fi59 := 0; fi59 < 15; fi59++ {
            v60 := GenS2{F0: v5}
            var v61 int = 661
            _ = fi59
            _ = v60
            _ = v61
        }
        _ = v55
    }
    v1 = v1
    for fi62 := 0; fi62 < 8; fi62++ {
        if true {
            if false {
                fmt.Println(v1)
                for fi63 := 0; fi63 < 10; fi63++ {
                    if (!false) {
                        v64 := []int{fi63, fi62, fi62}
                        _ = v64
                    } else {
                        var v65 uint32 = (v46 / ((v46) | 1))
                        if v3 {
                            v46 = v46
                            v51 = v51
                            var v66 int = (fi63 ^ -1)
                            _ = v66
                        } else {
                            v49 = (^(v49))
                            var v67 uint32 = (^(v46))
                            fmt.Println(v67, v2)
                            v68 := GenS6{F0: 201}
                            switch fi62 {
                            case 166:
                                switch fi63 {
                                case 128:
                                    var v69 uint16 = v5
                                    _ = v69
                                case 9223372036854775807:
                                    v70 := GenS4{F0: 667, F1: (^(v5))}
                                    fmt.Println(fi63, v65, v46)
                                    if v3 {
                                        if (v46 <= 582) {
                                            var v71 uint8 = 1
                                            v72 := []int{(fi63 + (-(fi62))), -9223372036854775808, (fi63 >> (((^(v49))) & 63)), 127}
                                            v73 := GenS3{F0: -1, F1: (v49 << ((v49) & 63))}
                                            _ = v71
                                            _ = v72
                                            _ = v73
                                        } else {
                                            var v74 uint8 = 1
                                            v75 := GenS4{F0: v5, F1: 0}
                                            var v76 uint = (v49 - (^(v49)))
                                            _ = v74
                                            _ = v75
                                            _ = v76
                                        }
                                        if v3 {
                                            fmt.Println(v1, v6, fi62)
                                        }
                                    }
                                    _ = v70
                                case 42:
                                    var v77 float64 = -2.25
                                    _ = v77
                                }
                            default:
                                v78 := GenS7{F0: uint64(v49), F1: int16(v49), F2: 859}
                                var v79 int = (fi63 >> ((v49) & 63))
                                if v3 {
                                    var v80 uint64 = v51
                                    _ = len(v4)
                                    _ = v80
                                }
                                v1 = (v1 + "")
                                v6 = v6
                                if v3 {
                                    if (-32768 == 0) {
                                        v81 := GenS8{F0: fi63, F1: 808, F2: -231}
                                        var v82 uint8 = 2
                                        var v83 bool = ((v51 % (((v51 | 2)) | 1)) <= 4294967295)
                                        if (v2 != v2) {
                                            var v84 uint64 = 2
                                            _ = v84
                                        }
                                        v85 := []int{fi63, (fi63 ^ fi62)}
                                        _ = v81
                                        _ = v82
                                        _ = v83
                                        _ = v85
                                    } else {
                                        if v3 {
                                            switch ("foo bar" + ("" + ("hello" + v1))) {
                                            case "hello":
                                                switch v1 {
                                                case "foo bar":
                                                    var v86 int64 = -1
                                                    _ = v86
                                                }
                                            case "\n\t":
                                                if (751 < (v5 & v5)) {
                                                    switch (v79 % ((int(v51)) | 1)) {
                                                    case 1:
                                                        v5 -= v5
                                                    case 411:
                                                        fmt.Println(v49, v51)
                                                    }
                                                }
                                            }
                                            var v87 uint32 = genFn0((^(v49)), (v49 / (((v49 & v49)) | 1)), false)
                                            switch v1 {
                                            case "hello":
                                                v88 := GenS9{F0: 65535, F1: v87, F2: (v87 - v67)}
                                                _ = v88
                                            case "\n\t":
                                                if (uint64(v5) >= 646) {
                                                    v89 := GenS5{F0: (-(v6)), F1: 2, F2: ""}
                                                    v90 := GenS10{F0: false}
                                                    switch -1 {
                                                    case 32767:
                                                        switch (v1 + v1) {
                                                        case "a":
                                                            var v91 uint16 = v5
                                                            _ = v91
                                                        case "\n\t":
                                                            if v3 {
                                                                v92 := GenS6{F0: -2147483648}
                                                                v93 := []int{fi63, 4294967295, fi62}
                                                                _ = v92
                                                                _ = v93
                                                            } else {
                                                                v1 = v1
                                                                if v3 {
                                                                    v94 := GenS10{F0: (!true)}
                                                                    switch v1 {
                                                                    case "\n\t":
                                                                        var v95 uint8 = 105
                                                                        _ = v95
                                                                    }
                                                                    _ = v94
                                                                } else {
                                                                    fmt.Println(v79, fi62)
                                                                    v2 = (v2 / -2.25)
                                                                    v96 := GenS10{F0: v3}
                                                                    _ = v96
                                                                }
                                                            }
                                                        case "":
                                                            var v97 int32 = 0
                                                            _ = v97
                                                        default:
                                                            v98 := GenS3{F0: -42, F1: 0}
                                                            _ = v98
                                                        }
                                                    case -9223372036854775808:
                                                        v6 = v6
                                                    case -365:
                                                        switch 2147483647 {
                                                        case 0:
                                                            if v3 {
                                                                v6 = -105
                                                                switch (fi63 - -330) {
                                                                case 32768:
                                                                    v99 := GenS7{F0: v51, F1: 118, F2: 2147483647}
                                                                    _ = v99
                                                                case -224:
                                                                    var v100 uint64 = v51
                                                                    _ = v100
                                                                }
                                                                v46 = uint32(v5)
                                                            } else {
                                                                switch "" {
                                                                case "a":
                                                                    switch (v1 + "a") {
                                                                    case "\n\t":
                                                                        var v101 uint16 = (v5 << ((v49) & 63))
                                                                        _ = v101
                                                                    case "a":
                                                                        var v102 string = v1
                                                                        _ = v102
                                                                    case "hello":
                                                                        var v103 int32 = -1
                                                                        _ = v103
                                                                    }
                                                                default:
                                                                    switch (v1 + (v1 + v1)) {
                                                                    case "a":
                                                                        if v3 {
                                                                            v104 := []int{}
                                                                            v105 := GenS9{F0: 861, F1: (v46 % ((genFn0(uint(v51), uint(v65), ((v49 % (((^(v49))) | 1)) >= (v49 >> (((v49 * 393)) & 63))))) | 1)), F2: genFn0((v49 ^ 65535), v49, v3)}
                                                                            if len(v104) >= 2 {
                                                                                v104 = v104[1:len(v104)]
                                                                            }
                                                                            _ = v104
                                                                            _ = v105
                                                                        }
                                                                    default:
                                                                        var v106 float32 = v2
                                                                        _ = v106
                                                                    }
                                                                }
                                                            }
                                                        case -1:
                                                            delete(v4, 42)
                                                        case 42:
                                                            switch v1 {
                                                            case "hello":
                                                                fmt.Println(v49, v2, v6)
                                                            }
                                                        default:
                                                            v107 := GenS11{F0: 2}
                                                            _ = v107
                                                        }
                                                    }
                                                    _ = v89
                                                    _ = v90
                                                }
                                            case "a":
                                                switch (v1 + v1) {
                                                case "a":
                                                    var v108 uint8 = 163
                                                    _ = v108
                                                case "\n\t":
                                                    var v109 int8 = (-(v6))
                                                    _ = v109
                                                default:
                                                    if ((3.14159 / float32((3.14159 + -2.25))) == v2) {
                                                        v110 := []int{fi62, (v79 % ((0) | 1)), v79, (fi62 + (-(fi63)))}
                                                        _ = v110
                                                    }
                                                }
                                            }
                                            _ = v87
                                        }
                                        v111 := GenS9{F0: uint16(v51), F1: (v65 % (((v65 & (v67 >> ((v49) & 63)))) | 1)), F2: v65}
                                        var v112 int8 = (v6 - 7)
                                        _ = v111
                                        _ = v112
                                    }
                                }
                                _ = v78
                                _ = v79
                            }
                            v113 := GenS12{F0: v1, F1: 2, F2: ((!v3) && v3)}
                            _ = v113
                            _ = v67
                            _ = v68
                        }
                        _ = v65
                    }
                    _ = fi63
                }
            }
            var v114 uint = v49
            _ = v114
        } else {
            var v115 int32 = genFn1(-398, -537, (^(v6)))
            if v3 {
                fmt.Println(v3, fi62)
                for fi116 := 0; fi116 < 1; fi116++ {
                    var v117 uint8 = 42
                    if ((9223372036854775807 >= fi62) || v3) {
                        var v118 uint64 = v51
                        _ = v118
                    } else {
                        v119 := GenS8{F0: fi116, F1: -124, F2: -108}
                        switch v1 {
                        case "foo bar":
                            v120 := v4[genFn0(v49, (v49 - 2), ((!v3) && true))]
                            _ = v120
                        case "a":
                            v51 -= 42
                        case "":
                            v115 -= genFn1(856, 292, v6)
                        default:
                            v121 := GenS13{F0: 1, F1: uint64(v117), F2: 9223372036854775807}
                            _ = v121
                        }
                        v122 := GenS14{F0: (v51 >> ((v49) & 63)), F1: -1, F2: v46}
                        _ = v119
                        _ = v122
                    }
                    _ = len(v4)
                    _ = fi116
                    _ = v117
                }
            }
            for fi123 := 0; fi123 < 11; fi123++ {
                switch "a" {
                case "a":
                    v124 := GenS1{F0: fi123}
                    _ = v124
                case "":
                    v125 := GenS14{F0: (v51 - v51), F1: 1, F2: 4294967295}
                    _ = v125
                }
                _ = fi123
            }
            _ = v115
        }
        var v126 int64 = 0
        var v127 int32 = 42
        _ = fi62
        _ = v126
        _ = v127
    }
    var v128 int64 = 0
    var v129 float32 = 0.1
    fmt.Println(v46)
    v51 += v51
    if (!v3) {
        var v130 int64 = -1
        if v3 {
            v4[(v46 * 2)] = (v1 + ((v1 + (v1 + ((v1 + v1) + (("" + ("\n\t" + v1)) + "a")))) + ((v1 + "a") + (((((((v1 + (v1 + v1)) + v1) + ("\n\t" + v1)) + ("foo bar" + v1)) + "a") + (v1 + v1)) + ("foo bar" + ((v1 + "") + "hello"))))))
            for fi131 := 0; fi131 < 1; fi131++ {
                if (v51 < 259) {
                    v132 := []int{(fi131 / (((fi131 | (-(fi131)))) | 1)), (fi131 / ((fi131) | 1))}
                    v1 = v1
                    var v133 int32 = 0
                    _ = v132
                    _ = v133
                }
                _ = fi131
            }
            var v134 uint8 = 42
            _ = v134
        }
        _ = v130
    }
    v128 += int64(v46)
    v135 := []int{}
    v46 = genFn0(255, 595, v3)
    fmt.Println("v1:", v1)
    fmt.Println("v128:", v128)
    fmt.Println("v129:", v129)
    fmt.Println("slicelen v135:", len(v135), v135)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v46:", v46)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v51:", v51)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v128
    _ = v129
    _ = v135
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v49
    _ = v5
    _ = v50
    _ = v51
    _ = v6
}
