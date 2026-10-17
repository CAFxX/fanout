package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint64; F1 uint64; F2 uint8 }
type GenS3 struct { F0 string; F1 uint8; F2 uint }
type GenS4 struct { F0 int; F1 string; F2 string }
type GenS5 struct { F0 uint32; F1 int8 }
type GenS6 struct { F0 string; F1 int }
type GenS7 struct { F0 int8; F1 int64 }
type GenS8 struct { F0 uint16; F1 bool; F2 uint64 }
type GenS9 struct { F0 int16 }
type GenS10 struct { F0 string; F1 int32; F2 uint32 }
type GenS11 struct { F0 uint32 }
type GenS12 struct { F0 bool; F1 uint16 }
type GenS13 struct { F0 string; F1 uint }
type GenS14 struct { F0 int; F1 int64; F2 bool }
type GenS15 struct { F0 int64; F1 int8 }
type GenS16 struct { F0 uint16 }
func genFn0(p0 uint8) uint16 {
    p0 -= (^(p0))
    v7 := []int{773}
    _ = p0
    _ = v7
    return 465
}
func genFn1() int {
    var v8 int8 = 56
    fmt.Println(v8)
    _ = v8
    return 127
}

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[int]bool)
    var v3 uint8 = 255
    var v4 string = ""
    var v5 uint16 = 2
    var v6 int8 = 127
    fmt.Println(v4, v6, v5)
    v5 = 32767
    switch v4 {
    case "a":
        var v9 uint8 = v3
        _ = v9
    default:
        switch 32768 {
        case -674:
            fmt.Println(v6, v4)
        default:
            for fi10 := 0; fi10 < 13; fi10++ {
                var v11 uint32 = 2147483648
                var v12 int = fi10
                switch (fi10 + v12) {
                case 1:
                    v13 := []int{genFn1(), (^(fi10)), -1, -624}
                    _ = v13
                }
                _ = fi10
                _ = v11
                _ = v12
            }
        }
    }
    if true {
        v14 := GenS1{F0: genFn0((v3 / (((v3 & (v3 * (^(v3))))) | 1)))}
        _ = v14
    }
    _ = len(v2)
    for fi15 := 0; fi15 < 2; fi15++ {
        var v16 int64 = 42
        switch "a" {
        case "foo bar":
            v17 := GenS1{F0: (v5 >> ((uint(v3)) & 63))}
            _ = v17
        case "":
            v4 = v4
        default:
            var v18 int32 = -1
            _ = v18
        }
        switch (^(fi15)) {
        case 517:
            fmt.Println(v16, fi15)
        case -804:
            for fi19 := 0; fi19 < 9; fi19++ {
                v16 = v16
                _ = fi19
            }
        }
        _ = fi15
        _ = v16
    }
    if (65535 >= 939) {
        if (2 > 321) {
            for fi20 := 0; fi20 < 19; fi20++ {
                var v21 int8 = 100
                for fi22 := 0; fi22 < 9; fi22++ {
                    if (!((!false) || (!((fi20 & 289) != (fi22 & genFn1()))))) {
                        if true {
                            v23 := []int{fi22, 128, (fi20 & genFn1())}
                            _ = v23
                        }
                        var v24 uint = 813
                        v5 += (v5 * (v5 % ((v5) | 1)))
                        _ = v24
                    } else {
                        v25 := v1["a"]
                        _ = v25
                    }
                    if (1.5 > float64((float32(1e10) * -2.25))) {
                        if (1 <= uint32(v5)) {
                            switch (fi22 + fi22) {
                            case 127:
                                v26 := GenS1{F0: (v5 & genFn0(v3))}
                                _ = v26
                            case 42:
                                if false {
                                    if (1 > 4294967295) {
                                        if (!true) {
                                            delete(v1, ("a" + v4))
                                            v27 := GenS1{F0: (v5 % (((v5 << ((0) & 63))) | 1))}
                                            v5 = genFn0(v3)
                                            var v28 string = v4
                                            if false {
                                                var v29 int = fi22
                                                _ = v29
                                            } else {
                                                if false {
                                                    fmt.Println(v3, v6)
                                                    v30 := GenS2{F0: 1, F1: 42, F2: (v3 + (v3 | v3))}
                                                    v6 += (v6 / ((v21) | 1))
                                                    _ = v30
                                                }
                                            }
                                            _ = v27
                                            _ = v28
                                        } else {
                                            v6 -= -1
                                        }
                                    } else {
                                        switch -1 {
                                        case -248:
                                            v4 = v4
                                        case 0:
                                            var v31 uint16 = genFn0((v3 / (((v3 | (v3 ^ (^(v3))))) | 1)))
                                            _ = v31
                                        case 128:
                                            var v32 int16 = 32767
                                            var v33 uint = uint(v5)
                                            v6 -= v6
                                            _ = v32
                                            _ = v33
                                        default:
                                            v34 := []int{(-(fi20)), fi20}
                                            _ = v34
                                        }
                                        v3 = v3
                                        var v35 uint16 = genFn0(227)
                                        _ = v35
                                    }
                                    var v36 uint16 = (v5 - v5)
                                    var v37 float32 = ((0.1 * 0.1) / (((0.1 / 1.0) - 1.5) / 1.0))
                                    var v38 int8 = v6
                                    _ = v36
                                    _ = v37
                                    _ = v38
                                }
                            case -2147483648:
                                fmt.Println(fi22, fi20)
                            }
                            v39 := []int{-2147483648, fi22, genFn1()}
                            _ = v39
                        }
                        fmt.Println(fi22)
                    } else {
                        fmt.Println(fi22, v6, fi20)
                        var v40 uint32 = 4294967295
                        _ = v40
                    }
                    _ = fi22
                }
                for fi41 := 0; fi41 < 11; fi41++ {
                    if false {
                        v42 := GenS2{F0: 2, F1: 702, F2: v3}
                        v43 := GenS2{F0: 0, F1: 89, F2: (v3 % ((85) | 1))}
                        _ = v42
                        _ = v43
                    }
                    var v44 int8 = v21
                    fmt.Println(fi20, fi41, v6)
                    _ = fi41
                    _ = v44
                }
                _ = fi20
                _ = v21
            }
            var v45 int8 = v6
            v1[v4] = 0
            _ = v45
        }
        if (genFn0((v3 ^ 128)) < v5) {
            v46 := GenS1{F0: v5}
            if (-1 == -607) {
                delete(v1, (v4 + (v4 + "hello")))
                if (255 <= 65535) {
                    v47 := []int{}
                    _ = v47
                } else {
                    if false {
                        v4 = v4
                    } else {
                        var v48 uint64 = 0
                        v49 := []int{-412, genFn1(), 4294967295, -547}
                        _ = v48
                        _ = v49
                    }
                }
            } else {
                if (("a" + ("\n\t" + "\n\t")) == v4) {
                    var v50 uint32 = 1
                    for fi51 := 0; fi51 < 18; fi51++ {
                        v52 := GenS3{F0: "a", F1: 128, F2: 42}
                        v53 := []int{}
                        _ = fi51
                        _ = v52
                        _ = v53
                    }
                    _ = v50
                }
            }
            v54 := GenS1{F0: genFn0(v3)}
            _ = v46
            _ = v54
        } else {
            for fi55 := 0; fi55 < 12; fi55++ {
                v56 := GenS4{F0: (fi55 / ((genFn1()) | 1)), F1: "a", F2: v4}
                for fi57 := 0; fi57 < 4; fi57++ {
                    if (-1 >= 9223372036854775807) {
                        v5 = v5
                        v58 := []int{fi55}
                        v59 := GenS5{F0: 380, F1: 127}
                        if (!(!false)) {
                            switch (v4 + v4) {
                            case "hello":
                                if (!(!(((1e10 - float32(((1.5 * 1.5) + (float64(1.5) + 1.5)))) < -2.25) || false))) {
                                    if true {
                                        v4 = v4
                                        if (false || (1000 == -667)) {
                                            if true {
                                                switch "\n\t" {
                                                case "\n\t":
                                                    switch 32768 {
                                                    case -709:
                                                        var v60 bool = (0.0 >= -2.25)
                                                        _ = v60
                                                    case 127:
                                                        if false {
                                                            v61 := GenS6{F0: "a", F1: (fi57 * (fi57 ^ genFn1()))}
                                                            var v62 uint8 = v3
                                                            v3 = (v62 % ((v3) | 1))
                                                            _ = v61
                                                            _ = v62
                                                        } else {
                                                            var v63 uint = 936
                                                            _ = v63
                                                        }
                                                        v4 = "a"
                                                        var v64 uint = 18446744073709551615
                                                        fmt.Println(v64, v6)
                                                        v65 := GenS7{F0: v6, F1: 1}
                                                        if (-36 >= -1) {
                                                            if (!false) {
                                                                v4 = v4
                                                            } else {
                                                                v3 -= (v3 * 42)
                                                                v66 := GenS8{F0: 2, F1: (((-1 != 42) || ((false && true) || (!((!true) || true)))) || false), F2: 18446744073709551615}
                                                                _ = v66
                                                            }
                                                            var v67 float64 = (float64(0.1) * (0.1 / 1e10))
                                                            v68 := v2[607]
                                                            _ = v67
                                                            _ = v68
                                                        }
                                                        switch -298 {
                                                        case -2147483648:
                                                            v4 = v4
                                                        case 32767:
                                                            if true {
                                                                v4 = (((v4 + (v4 + v4)) + (v4 + v4)) + "")
                                                                if true {
                                                                    switch "hello" {
                                                                    case "\n\t":
                                                                        v6 = v6
                                                                    case "":
                                                                        switch int(v5) {
                                                                        case 250:
                                                                            if false {
                                                                                var v69 bool = (251 != (v5 - v5))
                                                                                fmt.Println(v6, v6, v6)
                                                                                _ = v69
                                                                            } else {
                                                                                if true {
                                                                                    delete(v1, "\n\t")
                                                                                    if len(v58) >= 2 {
                                                                                        v58 = v58[1:len(v58)]
                                                                                    }
                                                                                    v70 := GenS9{F0: -32768}
                                                                                    _ = v70
                                                                                } else {
                                                                                    v64 += (v64 % (((v64 * (^(v64)))) | 1))
                                                                                    var v71 int16 = -412
                                                                                    _ = v71
                                                                                }
                                                                                fmt.Println(fi57)
                                                                                var v72 uint16 = genFn0((^(v3)))
                                                                                _ = v72
                                                                            }
                                                                            v4 = ""
                                                                            v73 := GenS2{F0: 4294967295, F1: 777, F2: v3}
                                                                            _ = v73
                                                                        case 4294967295:
                                                                            var v74 int8 = v6
                                                                            _ = v74
                                                                        }
                                                                    default:
                                                                        if false {
                                                                            var v75 uint32 = 65535
                                                                            if len(v58) >= 2 {
                                                                                v58 = v58[1:len(v58)]
                                                                            }
                                                                            _ = v75
                                                                        }
                                                                    }
                                                                }
                                                                v64 -= 18446744073709551615
                                                            }
                                                        }
                                                        _ = v64
                                                        _ = v65
                                                    default:
                                                        v4 = v4
                                                    }
                                                }
                                                switch -795 {
                                                case -607:
                                                    var v76 int16 = -242
                                                    _ = v76
                                                default:
                                                    v5 -= 255
                                                }
                                            }
                                        }
                                    }
                                } else {
                                    v4 = (v4 + ((v4 + v4) + v4))
                                    v4 = v4
                                    v6 -= (-(v6))
                                }
                                var v77 uint32 = 349
                                v77 -= (v77 * (^(v77)))
                                _ = v77
                            case "a":
                                if false {
                                    var v78 int64 = 42
                                    _ = v78
                                } else {
                                    var v79 uint16 = (^(v5))
                                    _ = v79
                                }
                            case "foo bar":
                                var v80 int = fi57
                                _ = v80
                            }
                            if true {
                                fmt.Println(v3, fi55)
                                var v81 uint = 616
                                _ = v81
                            }
                            switch genFn1() {
                            case 9223372036854775807:
                                v82 := GenS10{F0: v4, F1: -594, F2: 0}
                                _ = v82
                            default:
                                v3 = v3
                            }
                        } else {
                            switch int(v3) {
                            case 42:
                                var v83 uint16 = genFn0(v3)
                                _ = v83
                            case 4294967295:
                                v3 += (v3 ^ 255)
                            case 256:
                                v5 = (v5 * v5)
                            default:
                                if ((uint64(v5) != 684) && (37 > -185)) {
                                    switch fi55 {
                                    case -9223372036854775808:
                                        if true {
                                            var v84 int64 = -250
                                            fmt.Println(v6, v3, v5)
                                            _ = v84
                                        } else {
                                            v6 = v6
                                            v1[((v4 + "foo bar") + "a")] = 2
                                        }
                                    case 42:
                                        v5 = (v5 ^ (v5 / (((v5 ^ (v5 << ((1) & 63)))) | 1)))
                                    }
                                }
                            }
                            var v85 uint8 = (^(v3))
                            _ = v85
                        }
                        v86 := GenS11{F0: 42}
                        _ = v58
                        _ = v59
                        _ = v86
                    } else {
                        if (!(1 >= -600)) {
                            if (32767 <= -1) {
                                v87 := GenS12{F0: ((v5 % ((v5) | 1)) >= (v5 * (v5 << ((42) & 63)))), F1: v5}
                                v88 := GenS1{F0: genFn0((v3 + (^(v3))))}
                                if (18446744073709551615 != uint64(v3)) {
                                    fmt.Println(fi55)
                                    var v89 float64 = (0.1 * float64(((-2.25 - 1.5) - 1e10)))
                                    if true {
                                        v90 := GenS13{F0: "\n\t", F1: 19}
                                        _ = v90
                                    }
                                    v91 := []int{-9223372036854775808, (fi55 / ((genFn1()) | 1)), 1}
                                    _ = v89
                                    _ = v91
                                } else {
                                    var v92 uint32 = 653
                                    _ = len(v2)
                                    _ = v92
                                }
                                _ = v87
                                _ = v88
                            }
                            var v93 int8 = (-(v6))
                            var v94 uint16 = uint16(v3)
                            _ = v93
                            _ = v94
                        }
                    }
                    v5 = v5
                    v6 = (v6 & v6)
                    if true {
                        fmt.Println(fi57, fi55, v3)
                    }
                    _ = fi57
                }
                _ = fi55
                _ = v56
            }
            fmt.Println(v4, v3)
        }
    } else {
        if false {
            for fi95 := 0; fi95 < 5; fi95++ {
                v96 := GenS8{F0: genFn0((v3 * v3)), F1: (1000 == 1), F2: 382}
                var v97 string = ("a" + "hello")
                fmt.Println(fi95, v4)
                _ = fi95
                _ = v96
                _ = v97
            }
        }
        var v98 int16 = -32768
        _ = v98
    }
    for fi99 := 0; fi99 < 13; fi99++ {
        v5 = genFn0(v3)
        v1[(v4 + ("a" + (("a" + ((("hello" + ((("" + (v4 + v4)) + v4) + ((((("a" + v4) + "\n\t") + (v4 + (v4 + v4))) + (v4 + ("foo bar" + v4))) + (v4 + v4)))) + v4) + v4)) + ("" + v4))))] = 65535
        switch fi99 {
        case 1:
            var v100 bool = (!true)
            _ = v100
        case 0:
            switch (fi99 | fi99) {
            case 0:
                for fi101 := 0; fi101 < 12; fi101++ {
                    v102 := v2[genFn1()]
                    var v103 int = fi99
                    _ = fi101
                    _ = v102
                    _ = v103
                }
            case 2147483647:
                for fi104 := 0; fi104 < 4; fi104++ {
                    v105 := []int{genFn1(), (fi99 - 653), genFn1(), 32768}
                    if false {
                        if (false && (999 != 1000)) {
                            switch (fi99 * fi99) {
                            case 42:
                                v4 = v4
                            case 540:
                                v5 -= 255
                            default:
                                if len(v105) >= 2 {
                                    v105 = v105[1:len(v105)]
                                }
                            }
                            v6 = -128
                            if false {
                                var v106 float32 = 1e10
                                var v107 int32 = -1
                                _ = v106
                                _ = v107
                            } else {
                                if len(v105) > 0 {
                                    v105[0] = (fi104 >> ((1) & 63))
                                }
                            }
                        } else {
                            _ = len(v105)
                        }
                        switch v4 {
                        case "\n\t":
                            if true {
                                v108 := GenS14{F0: 9223372036854775807, F1: -9223372036854775808, F2: (!(!(-435 > 1)))}
                                var v109 uint = 2
                                var v110 float32 = 1e10
                                fmt.Println(v110, v109)
                                if (-760 == -1) {
                                    fmt.Println(fi99, v109)
                                } else {
                                    switch -1 {
                                    case 4294967295:
                                        switch genFn1() {
                                        case 9223372036854775807:
                                            switch fi99 {
                                            case 4294967295:
                                                fmt.Println(v6, v6, v109)
                                            case -9223372036854775808:
                                                var v111 int64 = 9223372036854775807
                                                fmt.Println(v3, v4)
                                                v110 -= float32(1e10)
                                                _ = v111
                                            }
                                        default:
                                            var v112 int = fi104
                                            _ = v112
                                        }
                                    case 143:
                                        switch (^(fi99)) {
                                        case 127:
                                            v113 := GenS15{F0: int64(v109), F1: (^(v6))}
                                            _ = v113
                                        }
                                    default:
                                        switch v4 {
                                        case "hello":
                                            if ((false || (706 <= 42)) && false) {
                                                v109 = (^(v109))
                                                v110 -= float32((1.5 * float64(float32(1.5))))
                                            } else {
                                                if (true && false) {
                                                    v105 = append(v105, fi99)
                                                }
                                                v114 := GenS16{F0: (^(v5))}
                                                _ = v114
                                            }
                                        }
                                    }
                                    if (0 != 4294967295) {
                                        if false {
                                            v115 := v2[genFn1()]
                                            fmt.Println(fi99, fi104)
                                            switch fi99 {
                                            case 32767:
                                                var v116 uint = v109
                                                _ = v116
                                            default:
                                                _ = len(v105)
                                            }
                                            _ = v115
                                        } else {
                                            v6 = 100
                                            switch "foo bar" {
                                            case "":
                                                delete(v1, v4)
                                            case "hello":
                                                fmt.Println(fi104, v5, fi104)
                                            case "\n\t":
                                                var v117 int8 = (^(v6))
                                                _ = v117
                                            default:
                                                fmt.Println(fi104, fi104, v3)
                                            }
                                            v3 -= v3
                                        }
                                        if ((!(1.5 <= ((1.5 / 3.14159) / 0.1))) && false) {
                                            v6 -= 127
                                        } else {
                                            if (-2.25 != -2.25) {
                                                fmt.Println(v110, v110, v109)
                                                var v118 int32 = 0
                                                _ = v118
                                            }
                                        }
                                        if len(v105) >= 2 {
                                            v105 = v105[1:len(v105)]
                                        }
                                    }
                                    fmt.Println(v3, v6, fi99)
                                    var v119 uint32 = 2
                                    _ = len(v105)
                                    _ = v119
                                }
                                _ = v108
                                _ = v109
                                _ = v110
                            }
                        default:
                            switch "hello" {
                            case "":
                                var v120 int64 = 971
                                _ = v120
                            case "\n\t":
                                switch (v4 + v4) {
                                case "\n\t":
                                    v121 := v2[(fi99 * fi99)]
                                    _ = v121
                                }
                            }
                        }
                    } else {
                        fmt.Println(fi104)
                    }
                    var v122 uint32 = 549
                    _ = fi104
                    _ = v105
                    _ = v122
                }
            case -796:
                var v123 int64 = -1
                _ = v123
            }
        case 136:
            switch ((v4 + "") + v4) {
            case "\n\t":
                if (-895 <= -162) {
                    v124 := []int{-1, (-(fi99)), (fi99 - (-(fi99))), fi99}
                    _ = v124
                }
            case "hello":
                var v125 uint64 = 908
                _ = v125
            default:
                switch v4 {
                case "a":
                    fmt.Println(fi99)
                case "":
                    switch (fi99 & (fi99 + (fi99 << ((1) & 63)))) {
                    case -954:
                        v6 = v6
                    case 4294967295:
                        for fi126 := 0; fi126 < 2; fi126++ {
                            var v127 int8 = (v6 << ((580) & 63))
                            v128 := GenS4{F0: fi99, F1: v4, F2: ((v4 + "a") + ("a" + "hello"))}
                            v5 = (v5 % (((v5 * (v5 % (((v5 >> ((0) & 63))) | 1)))) | 1))
                            _ = fi126
                            _ = v127
                            _ = v128
                        }
                    case 1:
                        for fi129 := 0; fi129 < 13; fi129++ {
                            var v130 float64 = 0.0
                            v131 := []int{fi129, genFn1(), fi99, fi129}
                            _ = fi129
                            _ = v130
                            _ = v131
                        }
                    default:
                        v5 = (v5 ^ genFn0((v3 & (v3 & (v3 ^ v3)))))
                    }
                default:
                    v132 := GenS4{F0: genFn1(), F1: (v4 + v4), F2: v4}
                    _ = v132
                }
            }
        default:
            for fi133 := 0; fi133 < 9; fi133++ {
                v3 = (v3 % ((v3) | 1))
                var v134 int8 = (v6 - v6)
                _ = fi133
                _ = v134
            }
        }
        _ = fi99
    }
    v3 = v3
    fmt.Println(v5, v5)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
