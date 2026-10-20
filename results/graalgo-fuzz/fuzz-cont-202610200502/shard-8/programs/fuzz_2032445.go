package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16; F2 uint8 }
type GenS2 struct { F0 int8; F1 int }
type GenS3 struct { F0 int; F1 uint64 }
type GenS4 struct { F0 int32; F1 uint64 }
type GenS5 struct { F0 int64; F1 int8 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 uint32; F1 uint; F2 uint8 }
type GenS8 struct { F0 uint32 }
type GenS9 struct { F0 int; F1 uint32; F2 uint64 }
type GenS10 struct { F0 bool }
type GenS11 struct { F0 bool }
type GenS12 struct { F0 int }
type GenS13 struct { F0 uint; F1 uint8; F2 uint32 }
type GenS14 struct { F0 int16 }
type GenS15 struct { F0 bool }
type GenS16 struct { F0 int8; F1 uint; F2 uint32 }
type GenS17 struct { F0 uint32; F1 uint64 }
type GenS18 struct { F0 uint64 }
func genFn0() uint {
    v7 := GenS1{F0: 65535, F1: 556, F2: 2}
    for fi8 := 0; fi8 < 8; fi8++ {
        var v9 uint16 = 0
        v9 = (v9 << ((1) & 63))
        _ = fi8
        _ = v9
    }
    if (255 < 42) {
        if false {
            v10 := GenS2{F0: 42, F1: -260}
            switch 775 {
            case 0:
                switch "\n\t" {
                case "hello":
                    var v11 int8 = 43
                    _ = v11
                default:
                    var v12 float32 = -2.25
                    _ = v12
                }
            case 1:
                var v13 uint = 46
                _ = v13
            case -601:
                var v14 string = "a"
                _ = v14
            default:
                var v15 float64 = -2.25
                _ = v15
            }
            switch "a" {
            case "foo bar":
                switch ("foo bar" + "") {
                case "foo bar":
                    v16 := []int{-723}
                    _ = v16
                case "":
                    if (4294967295 >= 65535) {
                        for fi17 := 0; fi17 < 5; fi17++ {
                            var v18 uint32 = 2147483648
                            if true {
                                var v19 string = "foo bar"
                                if false {
                                    v20 := GenS2{F0: -1, F1: (fi17 ^ (fi17 * fi17))}
                                    v21 := GenS3{F0: fi17, F1: 4294967295}
                                    v22 := GenS2{F0: -1, F1: (fi17 >> ((21) & 63))}
                                    _ = v20
                                    _ = v21
                                    _ = v22
                                } else {
                                    v23 := []int{(-(fi17)), (^(fi17)), fi17, (^(fi17))}
                                    v24 := GenS1{F0: 1, F1: 65535, F2: 0}
                                    if (-32768 <= 0) {
                                        v19 = "a"
                                        if true {
                                            var v25 uint = 1
                                            _ = v25
                                        }
                                    }
                                    _ = v23
                                    _ = v24
                                }
                                v26 := GenS1{F0: 32767, F1: 42, F2: 255}
                                _ = v19
                                _ = v26
                            } else {
                                v27 := []int{fi17}
                                _ = v27
                            }
                            _ = fi17
                            _ = v18
                        }
                    } else {
                        v28 := []int{-501}
                        var v29 uint8 = 2
                        _ = v28
                        _ = v29
                    }
                default:
                    var v30 bool = true
                    _ = v30
                }
            case "\n\t":
                if false {
                    v31 := GenS1{F0: 42, F1: 1, F2: 0}
                    v32 := GenS4{F0: 690, F1: 0}
                    _ = v31
                    _ = v32
                } else {
                    var v33 uint8 = 189
                    v34 := GenS5{F0: 1, F1: -128}
                    _ = v33
                    _ = v34
                }
            case "a":
                var v35 float32 = 1.5
                _ = v35
            }
            _ = v10
        }
    } else {
        v36 := []int{-9223372036854775808, -9223372036854775808}
        _ = v36
    }
    for fi37 := 0; fi37 < 7; fi37++ {
        fmt.Println(fi37)
        v38 := []int{(fi37 % ((fi37) | 1))}
        _ = fi37
        _ = v38
    }
    var v39 int64 = 42
    _ = v39
    _ = v7
    return 2
}
func genFn1(p0 int32, p1 string, p2 int32) uint32 {
    p1 = p1
    fmt.Println(p2)
    _ = p0
    _ = p1
    _ = p2
    return 2147483648
}

func main() {
    v1 := make(map[int]string)
    var v2 string = ""
    var v3 int = 0
    var v4 int32 = -2147483648
    v5 := make(map[string]bool)
    v6 := []uint32{571}
    if len(v6) > 0 {
        v6[0] = 65535
    }
    var v40 uint8 = 255
    var v41 float32 = 0.1
    delete(v1, -9223372036854775808)
    if ((!false) || false) {
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        var v42 int32 = v4
        _ = v42
    }
    v4 -= 0
    var v43 bool = true
    switch "foo bar" {
    case "\n\t":
        if v43 {
            v2 = v2
            switch v3 {
            case 1:
                if (-128 != 100) {
                    for fi44 := 0; fi44 < 2; fi44++ {
                        if true {
                            v45 := GenS4{F0: v4, F1: 688}
                            if true {
                                v46 := GenS6{F0: 65535}
                                _ = v46
                            }
                            switch "foo bar" {
                            case "foo bar":
                                fmt.Println(v4)
                            default:
                                var v47 uint64 = 18446744073709551615
                                _ = v47
                            }
                            _ = v45
                        } else {
                            switch v2 {
                            case "":
                                v6 = append(v6, 42)
                            case "foo bar":
                                if v43 {
                                    v48 := v5[""]
                                    v49 := GenS2{F0: -93, F1: 128}
                                    _ = v48
                                    _ = v49
                                }
                            case "\n\t":
                                v50 := GenS1{F0: 606, F1: 760, F2: (^(v40))}
                                _ = v50
                            default:
                                for fi51 := 0; fi51 < 8; fi51++ {
                                    fmt.Println(v2, v4, fi51)
                                    var v52 uint8 = 127
                                    _ = fi51
                                    _ = v52
                                }
                            }
                            if v43 {
                                if len(v6) >= 2 {
                                    v6 = v6[1:len(v6)]
                                }
                                switch (fi44 >> ((4294967295) & 63)) {
                                case 0:
                                    switch (fi44 >> ((genFn0()) & 63)) {
                                    case 127:
                                        fmt.Println(v40, v2)
                                    case -1:
                                        var v53 float64 = 1.5
                                        _ = v53
                                    default:
                                        if len(v6) > 0 {
                                            v6[0] = 588
                                        }
                                    }
                                case -801:
                                    for fi54 := 0; fi54 < 14; fi54++ {
                                        var v55 int8 = 100
                                        var v56 uint = 197
                                        v55 -= -1
                                        _ = fi54
                                        _ = v55
                                        _ = v56
                                    }
                                case 272:
                                    var v57 int16 = 1
                                    _ = v57
                                default:
                                    v58 := GenS1{F0: 65535, F1: 227, F2: 255}
                                    _ = v58
                                }
                                for fi59 := 0; fi59 < 17; fi59++ {
                                    var v60 int32 = v4
                                    _ = fi59
                                    _ = v60
                                }
                            }
                        }
                        v61 := GenS1{F0: 0, F1: 65535, F2: 136}
                        _ = fi44
                        _ = v61
                    }
                    for fi62 := 0; fi62 < 7; fi62++ {
                        v3 = fi62
                        var v63 uint16 = 565
                        v64 := v5["\n\t"]
                        _ = fi62
                        _ = v63
                        _ = v64
                    }
                } else {
                    v65 := GenS7{F0: 1, F1: 478, F2: 234}
                    if ((-32768 <= -32768) || v43) {
                        delete(v5, v2)
                    } else {
                        v40 -= v40
                        var v66 float64 = ((0.1 - (0.1 - -2.25)) + 1e10)
                        fmt.Println(v40, v66)
                        _ = v66
                    }
                    if (!false) {
                        for fi67 := 0; fi67 < 6; fi67++ {
                            var v68 uint64 = 4294967295
                            if len(v6) > 0 {
                                v6[0] = 1
                            }
                            if (776 > 42) {
                                switch ("\n\t" + "a") {
                                case "foo bar":
                                    v69 := GenS5{F0: 1, F1: 127}
                                    _ = v69
                                case "a":
                                    for fi70 := 0; fi70 < 1; fi70++ {
                                        _ = len(v6)
                                        _ = fi70
                                    }
                                case "":
                                    v71 := GenS8{F0: 0}
                                    _ = v71
                                default:
                                    var v72 int8 = 0
                                    _ = v72
                                }
                                v4 += v4
                                fmt.Println(fi67)
                            } else {
                                fmt.Println(v43, v3, v3)
                                var v73 int = 956
                                _ = v73
                            }
                            _ = fi67
                            _ = v68
                        }
                        v1[617] = ("\n\t" + "\n\t")
                        v74 := GenS9{F0: (v3 | int(v4)), F1: 324, F2: 4294967295}
                        _ = v74
                    }
                    _ = v65
                }
            case 0:
                for fi75 := 0; fi75 < 6; fi75++ {
                    if v43 {
                        var v76 string = "a"
                        _ = v76
                    }
                    var v77 uint32 = 770
                    _ = fi75
                    _ = v77
                }
            default:
                if len(v6) > 0 {
                    v6[0] = genFn1((v4 << ((18446744073709551615) & 63)), v2, v4)
                }
            }
            switch v3 {
            case 101:
                switch (v3 - v3) {
                case 2147483647:
                    v43 = false
                case 128:
                    switch "" {
                    case "":
                        var v78 float64 = float64(float32(0.1))
                        _ = v78
                    default:
                        switch v3 {
                        case 433:
                            v3 = 0
                        case 2147483647:
                            if v43 {
                                if (1 == -587) {
                                    var v79 int16 = 0
                                    var v80 int8 = -1
                                    _ = v79
                                    _ = v80
                                } else {
                                    var v81 int64 = -1
                                    v41 = 0.1
                                    _ = v81
                                }
                                switch 32768 {
                                case -732:
                                    if (1 < 1) {
                                        v2 = ("a" + "hello")
                                        for fi82 := 0; fi82 < 2; fi82++ {
                                            switch v2 {
                                            case "hello":
                                                v83 := GenS2{F0: 42, F1: (v3 | 32768)}
                                                _ = v83
                                            case "\n\t":
                                                var v84 bool = false
                                                _ = v84
                                            case "a":
                                                for fi85 := 0; fi85 < 6; fi85++ {
                                                    if v43 {
                                                        switch "" {
                                                        case "\n\t":
                                                            if len(v6) >= 2 {
                                                                v6 = v6[1:len(v6)]
                                                            }
                                                        case "foo bar":
                                                            var v86 int = -2147483648
                                                            _ = v86
                                                        case "a":
                                                            switch "hello" {
                                                            case "a":
                                                                switch v2 {
                                                                case "hello":
                                                                    var v87 uint32 = 2147483647
                                                                    _ = v87
                                                                case "":
                                                                    v88 := GenS10{F0: true}
                                                                    switch v3 {
                                                                    case 1:
                                                                        v40 = v40
                                                                    default:
                                                                        v5["hello"] = v43
                                                                    }
                                                                    _ = v88
                                                                case "\n\t":
                                                                    if true {
                                                                        v89 := GenS11{F0: true}
                                                                        v43 = false
                                                                        var v90 uint8 = v40
                                                                        _ = v89
                                                                        _ = v90
                                                                    } else {
                                                                        switch (v2 + (v2 + ("\n\t" + ""))) {
                                                                        case "\n\t":
                                                                            v3 = (fi82 | fi82)
                                                                        case "a":
                                                                            if v43 {
                                                                                fmt.Println(v3)
                                                                            }
                                                                        default:
                                                                            switch (("foo bar" + ("hello" + "a")) + v2) {
                                                                            case "hello":
                                                                                if true {
                                                                                    v6 = append(v6, 2)
                                                                                } else {
                                                                                    fmt.Println(v41)
                                                                                    v5["foo bar"] = (v43 || (((v2 + (("" + v2) + ("a" + "hello"))) + v2) != "\n\t"))
                                                                                    var v91 int64 = 0
                                                                                    _ = v91
                                                                                }
                                                                            }
                                                                            v92 := GenS12{F0: (-(v3))}
                                                                            var v93 int16 = -532
                                                                            _ = v92
                                                                            _ = v93
                                                                        }
                                                                        v2 = "\n\t"
                                                                        fmt.Println(v3, fi85, fi82)
                                                                    }
                                                                default:
                                                                    if v43 {
                                                                        v94 := GenS12{F0: 0}
                                                                        v1[768] = (((("foo bar" + v2) + v2) + "a") + (v2 + (((("\n\t" + v2) + ("hello" + ((((v2 + v2) + (v2 + v2)) + (((v2 + (v2 + ((v2 + v2) + v2))) + v2) + ("\n\t" + v2))) + "\n\t"))) + v2) + "foo bar")))
                                                                        _ = v94
                                                                    }
                                                                }
                                                            }
                                                        }
                                                    } else {
                                                        v95 := GenS13{F0: 557, F1: v40, F2: 2147483648}
                                                        v96 := GenS14{F0: 1}
                                                        v41 = v41
                                                        _ = v95
                                                        _ = v96
                                                    }
                                                    var v97 int64 = 22
                                                    if (1 == 18446744073709551615) {
                                                        switch v2 {
                                                        case "":
                                                            v98 := GenS15{F0: false}
                                                            _ = v98
                                                        case "a":
                                                            if ((1.5 - -2.25) >= float32(0.0)) {
                                                                var v99 bool = v43
                                                                v4 -= (v4 % ((-572) | 1))
                                                                if (65535 < 1) {
                                                                    _ = len(v5)
                                                                } else {
                                                                    v2 = ("hello" + v2)
                                                                    if v99 {
                                                                        if ((100 > 0) || (!v43)) {
                                                                            var v100 int16 = 42
                                                                            switch fi85 {
                                                                            case -9223372036854775808:
                                                                                v97 -= (v97 >> ((1) & 63))
                                                                            default:
                                                                                if (!(false && true)) {
                                                                                    var v101 float64 = float64((v41 - ((0.1 + -2.25) / ((v41 / -2.25) + float32((1e10 - (float64(float32(0.1)) * 3.14159)))))))
                                                                                    var v102 int32 = (v4 / (((v4 ^ (v4 >> ((uint(v40)) & 63)))) | 1))
                                                                                    fmt.Println(fi85)
                                                                                    var v103 uint64 = 42
                                                                                    v41 -= v41
                                                                                    _ = v101
                                                                                    _ = v102
                                                                                    _ = v103
                                                                                } else {
                                                                                    var v104 uint32 = 0
                                                                                    fmt.Println(v40, v100)
                                                                                    if false {
                                                                                        var v105 float64 = 0.1
                                                                                        var v106 bool = v99
                                                                                        _ = v105
                                                                                        _ = v106
                                                                                    }
                                                                                    _ = v104
                                                                                }
                                                                            }
                                                                            _ = v100
                                                                        } else {
                                                                            fmt.Println(fi82, v41, v97)
                                                                        }
                                                                        if len(v6) > 0 {
                                                                            v6[0] = 887
                                                                        }
                                                                    }
                                                                }
                                                                switch "hello" {
                                                                case "a":
                                                                    if v99 {
                                                                        v2 = "hello"
                                                                    } else {
                                                                        fmt.Println(v41, v40)
                                                                        if ((((v40 < v40) && (v43 || (!true))) || ((v2 == "foo bar") || v43)) || (v97 == (^(v97)))) {
                                                                            v107 := GenS16{F0: 0, F1: 244, F2: 100}
                                                                            _ = v107
                                                                        } else {
                                                                            if ((v97 << ((65535) & 63)) >= -9223372036854775808) {
                                                                                var v108 uint = 18446744073709551615
                                                                                if len(v6) >= 2 {
                                                                                    v6 = v6[1:len(v6)]
                                                                                }
                                                                                if true {
                                                                                    var v109 bool = v99
                                                                                    _ = v109
                                                                                } else {
                                                                                    delete(v1, int(v4))
                                                                                    v41 -= 3.14159
                                                                                }
                                                                                _ = v108
                                                                            }
                                                                        }
                                                                    }
                                                                default:
                                                                    v2 = ""
                                                                }
                                                                var v110 uint16 = 2
                                                                if true {
                                                                    v111 := GenS6{F0: 65535}
                                                                    if ((!true) || (v3 > v3)) {
                                                                        fmt.Println(v40, v43, v40)
                                                                        v3 = fi82
                                                                        switch (v2 + "foo bar") {
                                                                        case "a":
                                                                            var v112 int16 = 42
                                                                            _ = v112
                                                                        case "":
                                                                            var v113 uint64 = 18446744073709551615
                                                                            _ = v113
                                                                        case "\n\t":
                                                                            fmt.Println(v3, fi85)
                                                                        }
                                                                    }
                                                                    _ = v111
                                                                } else {
                                                                    v40 += v40
                                                                    v114 := v1[(v3 & (fi82 / (((fi85 / ((127) | 1))) | 1)))]
                                                                    _ = v114
                                                                }
                                                                var v115 uint64 = 958
                                                                _ = v110
                                                                _ = v115
                                                                _ = v99
                                                            }
                                                        case "foo bar":
                                                            if false {
                                                                if ((v4 ^ v4) < v4) {
                                                                    v3 -= (fi82 >> ((935) & 63))
                                                                    var v116 string = "hello"
                                                                    _ = v116
                                                                }
                                                            }
                                                        default:
                                                            if (!v43) {
                                                                v117 := GenS16{F0: 42, F1: 1, F2: 4294967295}
                                                                _ = v117
                                                            }
                                                        }
                                                        v118 := GenS17{F0: uint32(v40), F1: 0}
                                                        _ = v118
                                                    }
                                                    _ = fi85
                                                    _ = v97
                                                }
                                            default:
                                                var v119 float32 = 1.5
                                                _ = v119
                                            }
                                            v40 += v40
                                            _ = fi82
                                        }
                                        v43 = (v3 > (v3 | (-(v3))))
                                    } else {
                                        if len(v6) > 0 {
                                            v6[0] = 2
                                        }
                                        v120 := GenS16{F0: 1, F1: 762, F2: 42}
                                        var v121 uint8 = v40
                                        _ = v120
                                        _ = v121
                                    }
                                }
                            } else {
                                for fi122 := 0; fi122 < 4; fi122++ {
                                    v41 = (0.1 + (0.1 + -2.25))
                                    for fi123 := 0; fi123 < 20; fi123++ {
                                        var v124 int8 = 42
                                        switch 127 {
                                        case 2147483647:
                                            if v43 {
                                                fmt.Println(v124)
                                                v2 = "foo bar"
                                                v124 = (-(v124))
                                                v4 -= v4
                                                var v125 uint8 = 1
                                                _ = v125
                                            }
                                        case 128:
                                            fmt.Println(v4, v3)
                                        }
                                        v126 := GenS4{F0: 205, F1: 801}
                                        _ = fi123
                                        _ = v124
                                        _ = v126
                                    }
                                    var v127 uint = 255
                                    _ = fi122
                                    _ = v127
                                }
                                v128 := GenS16{F0: -128, F1: 42, F2: 4294967295}
                                _ = v128
                            }
                        default:
                            switch (("" + v2) + ("a" + ("hello" + "a"))) {
                            case "hello":
                                for fi129 := 0; fi129 < 7; fi129++ {
                                    var v130 bool = v43
                                    _ = fi129
                                    _ = v130
                                }
                            case "foo bar":
                                _ = len(v6)
                            }
                        }
                    }
                case 32768:
                    var v131 string = "foo bar"
                    _ = v131
                }
            case 32767:
                v40 -= (v40 - v40)
            }
        }
    case "a":
        v132 := GenS18{F0: 42}
        _ = v132
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v43:", v43)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v43
    _ = v5
    _ = v6
}
