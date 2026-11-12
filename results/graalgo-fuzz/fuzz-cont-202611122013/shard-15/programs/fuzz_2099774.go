package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32 }
type GenS2 struct { F0 uint; F1 uint; F2 int64 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint; F1 int8; F2 string }
type GenS5 struct { F0 uint8; F1 int32; F2 int8 }
type GenS6 struct { F0 uint8; F1 uint8 }
type GenS7 struct { F0 int16; F1 uint64 }
type GenS8 struct { F0 string; F1 uint32; F2 uint8 }
type GenS9 struct { F0 uint16 }
type GenS10 struct { F0 int16; F1 uint8 }
type GenS11 struct { F0 uint8; F1 int64 }
type GenS12 struct { F0 int8 }
func genFn0(p0 uint) int64 {
    p0 -= 302
    var v7 uint64 = 46
    if true {
        var v8 string = "foo bar"
        for fi9 := 0; fi9 < 2; fi9++ {
            v10 := GenS1{F0: -9223372036854775808, F1: 1}
            for fi11 := 0; fi11 < 9; fi11++ {
                p0 = (p0 >> ((2) & 63))
                var v12 int32 = -1
                var v13 int8 = -31
                _ = fi11
                _ = v12
                _ = v13
            }
            v14 := GenS2{F0: 673, F1: 692, F2: 158}
            _ = fi9
            _ = v10
            _ = v14
        }
        _ = v8
    } else {
        for fi15 := 0; fi15 < 6; fi15++ {
            if (0 > -44) {
                v16 := []int{0, (fi15 | (fi15 >> (((p0 * (^(p0)))) & 63))), fi15, (fi15 + fi15)}
                var v17 int64 = 785
                v17 = (v17 ^ int64(v7))
                _ = v16
                _ = v17
            }
            _ = fi15
        }
        var v18 uint16 = 42
        p0 = 679
        _ = v18
    }
    _ = p0
    _ = v7
    return -9223372036854775808
}
func genFn1(p0 uint, p1 uint16, p2 int8) int {
    v19 := []int{661}
    for fi20 := 0; fi20 < 16; fi20++ {
        v19 = append(v19, fi20)
        p1 += (p1 ^ p1)
        _ = fi20
    }
    p2 = 52
    v19 = append(v19, -2147483648)
    p1 = (p1 * p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v19
    return 214
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[uint32]string)
    var v3 uint32 = 290
    var v4 float64 = (3.14159 - float64((0.1 + 0.1)))
    var v5 uint = 4294967295
    var v6 int16 = -1
    if false {
        if ((v3 * (v3 ^ v3)) == 4294967295) {
            if (v6 >= v6) {
                var v21 int32 = 42
                v2[4294967295] = "hello"
                _ = v21
            }
            v22 := []int{}
            for fi23 := 0; fi23 < 13; fi23++ {
                v24 := GenS1{F0: -1, F1: (v3 / (((v3 * v3)) | 1))}
                if (true && (42 == 0)) {
                    var v25 int16 = -32768
                    _ = v25
                } else {
                    v26 := GenS2{F0: (v5 >> ((v5) & 63)), F1: v5, F2: 0}
                    _ = v26
                }
                v6 += (v6 * (v6 ^ v6))
                _ = fi23
                _ = v24
            }
            _ = v22
        }
    } else {
        v4 += 0.1
        for fi27 := 0; fi27 < 11; fi27++ {
            v2[(v3 + (^(v3)))] = "hello"
            for fi28 := 0; fi28 < 11; fi28++ {
                fmt.Println(v6, fi28)
                var v29 int32 = -1
                _ = fi28
                _ = v29
            }
            _ = fi27
        }
        var v30 uint8 = 127
        _ = v30
    }
    fmt.Println(v3)
    if (false && ((1.5 != 1.5) && false)) {
        delete(v2, v3)
        var v31 float32 = float32(v4)
        _ = v31
    }
    fmt.Println(v3)
    if false {
        fmt.Println(v6, v4, v3)
        v32 := GenS3{F0: 0}
        _ = v32
    } else {
        for fi33 := 0; fi33 < 6; fi33++ {
            fmt.Println(v6, v6, v4)
            for fi34 := 0; fi34 < 16; fi34++ {
                switch int(v3) {
                case -737:
                    var v35 uint16 = 65535
                    _ = v35
                default:
                    if (42 > -2147483648) {
                        v4 -= 0.1
                        v36 := GenS2{F0: 42, F1: (v5 * (v5 + v5)), F2: 1}
                        if ((0.0 - (1.5 * ((1.5 * 0.1) / float32(0.1)))) == 1.5) {
                            v6 += (v6 << ((v5) & 63))
                            var v37 int8 = 42
                            v38 := GenS2{F0: (v5 - v5), F1: v5, F2: 272}
                            _ = v37
                            _ = v38
                        }
                        _ = v36
                    } else {
                        v5 -= (v5 + v5)
                        v39 := GenS3{F0: 50}
                        _ = v39
                    }
                }
                _ = fi34
            }
            switch fi33 {
            case -9223372036854775808:
                v5 = (v5 % ((v5) | 1))
            }
            _ = fi33
        }
        switch "foo bar" {
        case "\n\t":
            switch int(v3) {
            case 42:
                v5 = (^(v5))
            case 32768:
                if (true || false) {
                    var v40 uint8 = 128
                    _ = v40
                } else {
                    v4 -= 1e10
                    v2[2147483648] = "hello"
                }
            case 9223372036854775807:
                if ((!true) && (true && ((0 > 127) || false))) {
                    switch genFn1((v5 >> (((^(v5))) & 63)), 42, -16) {
                    case 0:
                        for fi41 := 0; fi41 < 10; fi41++ {
                            v42 := GenS4{F0: 255, F1: -36, F2: "\n\t"}
                            for fi43 := 0; fi43 < 9; fi43++ {
                                fmt.Println(fi43, v6, v4)
                                if false {
                                    if true {
                                        if (!((2 > 42) && false)) {
                                            var v44 uint16 = 32767
                                            if false {
                                                if ("\n\t" <= "a") {
                                                    switch "foo bar" {
                                                    case "\n\t":
                                                        switch ("" + "foo bar") {
                                                        case "hello":
                                                            delete(v2, v3)
                                                        case "foo bar":
                                                            switch ("a" + "\n\t") {
                                                            case "\n\t":
                                                                v4 += ((float64(float32(v4)) / v4) - float64(1.5))
                                                            default:
                                                                if false {
                                                                    fmt.Println(fi43, v44)
                                                                    if false {
                                                                        var v45 int32 = 1
                                                                        var v46 uint64 = 18446744073709551615
                                                                        var v47 uint64 = (^(v46))
                                                                        _ = v45
                                                                        _ = v46
                                                                        _ = v47
                                                                    } else {
                                                                        var v48 string = "hello"
                                                                        var v49 int64 = genFn0((^(v5)))
                                                                        v5 = 42
                                                                        _ = v48
                                                                        _ = v49
                                                                    }
                                                                    var v50 int8 = 127
                                                                    v3 = v3
                                                                    switch (-(fi43)) {
                                                                    case 42:
                                                                        v51 := GenS2{F0: v5, F1: v5, F2: genFn0(v5)}
                                                                        v6 = (^(v6))
                                                                        fmt.Println(v3, fi43, v5)
                                                                        _ = v51
                                                                    case 128:
                                                                        if true {
                                                                            var v52 uint16 = v44
                                                                            v53 := GenS3{F0: 18446744073709551615}
                                                                            _ = v52
                                                                            _ = v53
                                                                        }
                                                                    case -9223372036854775808:
                                                                        fmt.Println(v50, fi43, fi41)
                                                                    }
                                                                    _ = v50
                                                                } else {
                                                                    if false {
                                                                        v5 += v5
                                                                        switch 42 {
                                                                        case 400:
                                                                            _ = len(v2)
                                                                        case -1:
                                                                            fmt.Println(v44)
                                                                        default:
                                                                            var v54 int32 = 2147483647
                                                                            _ = v54
                                                                        }
                                                                    }
                                                                    var v55 uint8 = 244
                                                                    _ = v55
                                                                }
                                                            }
                                                        case "a":
                                                            v56 := GenS4{F0: v5, F1: 1, F2: (("a" + "hello") + "")}
                                                            v57 := GenS4{F0: 18446744073709551615, F1: -128, F2: "\n\t"}
                                                            var v58 int = 32767
                                                            _ = v56
                                                            _ = v57
                                                            _ = v58
                                                        default:
                                                            v59 := GenS4{F0: 65535, F1: 1, F2: ("a" + ("" + "foo bar"))}
                                                            _ = v59
                                                        }
                                                        if false {
                                                            if false {
                                                                v60 := []int{}
                                                                if len(v60) > 0 {
                                                                    v60[0] = (fi43 ^ 128)
                                                                }
                                                                var v61 float64 = v4
                                                                var v62 int = (fi43 % ((-303) | 1))
                                                                _ = v60
                                                                _ = v61
                                                                _ = v62
                                                            }
                                                            fmt.Println(v44)
                                                            var v63 string = ""
                                                            _ = v63
                                                        } else {
                                                            var v64 uint32 = v3
                                                            fmt.Println(v3)
                                                            _ = v64
                                                        }
                                                    }
                                                    if false {
                                                        switch "foo bar" {
                                                        case "\n\t":
                                                            v5 = 65535
                                                        case "a":
                                                            if (false && true) {
                                                                v5 = (v5 << ((v5) & 63))
                                                            } else {
                                                                var v65 uint32 = (v3 & (v3 + 2147483647))
                                                                switch ("\n\t" + "a") {
                                                                case "":
                                                                    var v66 int16 = v6
                                                                    _ = v66
                                                                case "\n\t":
                                                                    var v67 int16 = v6
                                                                    _ = v67
                                                                case "hello":
                                                                    if ((v3 >> (((^(v5))) & 63)) > v3) {
                                                                        if (!false) {
                                                                            v65 += v65
                                                                        }
                                                                        if true {
                                                                            if false {
                                                                                v2[2147483648] = ("\n\t" + (("a" + "a") + (("foo bar" + "foo bar") + ("foo bar" + ((("a" + "foo bar") + "") + "a")))))
                                                                            }
                                                                            fmt.Println(v6, v4, v6)
                                                                            v44 = (v44 >> ((v5) & 63))
                                                                        }
                                                                    }
                                                                default:
                                                                    if false {
                                                                        v68 := GenS2{F0: (v5 ^ v5), F1: v5, F2: 658}
                                                                        v6 = (v6 << (((^(v5))) & 63))
                                                                        v65 = v3
                                                                        _ = v68
                                                                    }
                                                                }
                                                                if true {
                                                                    v69 := []int{-241, (fi41 - int(v5)), (fi43 & 127), (fi43 & (fi43 | 71))}
                                                                    _ = v69
                                                                } else {
                                                                    v70 := GenS3{F0: 21}
                                                                    if (1e10 <= -2.25) {
                                                                        if false {
                                                                            v4 = -2.25
                                                                        }
                                                                    }
                                                                    _ = v70
                                                                }
                                                                v71 := []int{(fi41 / (((-(fi43))) | 1)), (fi41 ^ (fi41 / ((fi41) | 1))), fi43}
                                                                _ = v65
                                                                _ = v71
                                                            }
                                                        case "":
                                                            var v72 uint32 = (v3 >> ((1) & 63))
                                                            _ = v72
                                                        }
                                                        var v73 bool = true
                                                        _ = v73
                                                    } else {
                                                        var v74 int32 = -938
                                                        v2[v3] = "\n\t"
                                                        var v75 int64 = 9223372036854775807
                                                        v5 = (v5 & (v5 / ((v5) | 1)))
                                                        var v76 uint64 = 253
                                                        fmt.Println(v6)
                                                        _ = v74
                                                        _ = v75
                                                        _ = v76
                                                    }
                                                }
                                            } else {
                                                switch fi43 {
                                                case 135:
                                                    fmt.Println(v5)
                                                case 757:
                                                    var v77 int32 = 1
                                                    _ = v77
                                                case 1:
                                                    var v78 int32 = -1
                                                    _ = v78
                                                }
                                                v79 := []int{genFn1(v5, v44, 127), genFn1(uint(v44), v44, 42), (fi41 % (((fi41 | fi43)) | 1)), 4294967295}
                                                _ = v79
                                            }
                                            _ = v44
                                        }
                                        if (((1.5 + 3.14159) == -2.25) && (-9223372036854775808 == 0)) {
                                            var v80 string = ("a" + (("" + ("\n\t" + "a")) + "hello"))
                                            v81 := []int{(fi41 & (fi43 << (((^(v5))) & 63))), int(v5)}
                                            _ = v80
                                            _ = v81
                                        }
                                    } else {
                                        switch "" {
                                        case "a":
                                            switch genFn1(42, 65535, -128) {
                                            case 128:
                                                var v82 float64 = v4
                                                _ = v82
                                            case -2147483648:
                                                v83 := GenS5{F0: 127, F1: -2147483648, F2: 42}
                                                _ = v83
                                            case -499:
                                                v4 += ((v4 + v4) - 3.14159)
                                                var v84 float64 = ((v4 + v4) / float64(1e10))
                                                _ = v84
                                            default:
                                                v85 := GenS6{F0: 0, F1: 1}
                                                _ = v85
                                            }
                                        case "hello":
                                            if ((1 == 2147483647) || (!(-128 > 42))) {
                                                v86 := GenS7{F0: 1, F1: 466}
                                                v87 := GenS8{F0: "\n\t", F1: (v3 ^ v3), F2: 255}
                                                _ = v86
                                                _ = v87
                                            }
                                        default:
                                            delete(v1, (v3 / ((2) | 1)))
                                        }
                                        var v88 float32 = 1e10
                                        _ = v88
                                    }
                                    fmt.Println(v4, fi43, fi41)
                                    if (true || (128 < 128)) {
                                        if (11 != v3) {
                                            switch ("a" + "hello") {
                                            case "foo bar":
                                                var v89 float32 = 1.5
                                                _ = v89
                                            case "\n\t":
                                                var v90 int8 = 127
                                                _ = v90
                                            default:
                                                switch (("" + "") + ("\n\t" + (("foo bar" + "a") + ""))) {
                                                case "foo bar":
                                                    v5 = (v5 - (v5 * (v5 / ((v5) | 1))))
                                                default:
                                                    v91 := GenS2{F0: 388, F1: (v5 & v5), F2: genFn0(18446744073709551615)}
                                                    _ = v91
                                                }
                                            }
                                            v3 -= 765
                                            var v92 int32 = 2147483647
                                            var v93 uint = v5
                                            _ = v92
                                            _ = v93
                                        } else {
                                            var v94 int16 = (^(v6))
                                            fmt.Println(fi41)
                                            _ = v94
                                        }
                                        fmt.Println(fi41, v4, v4)
                                        var v95 int16 = (v6 ^ v6)
                                        if (v5 >= 0) {
                                            delete(v2, 4294967295)
                                            var v96 int64 = 42
                                            v4 += ((3.14159 * (v4 + -2.25)) + 3.14159)
                                            _ = v96
                                        } else {
                                            switch ((("foo bar" + "a") + ((("foo bar" + "") + ("" + "a")) + "foo bar")) + "a") {
                                            case "hello":
                                                if false {
                                                    v5 -= v5
                                                    fmt.Println(v6)
                                                    v97 := GenS7{F0: (^(v95)), F1: 1}
                                                    _ = v97
                                                } else {
                                                    v3 += 0
                                                    v98 := GenS9{F0: 32767}
                                                    v99 := GenS10{F0: 1, F1: 185}
                                                    _ = v98
                                                    _ = v99
                                                }
                                            case "foo bar":
                                                var v100 uint64 = 541
                                                _ = v100
                                            }
                                        }
                                        v3 = v3
                                        _ = v95
                                    } else {
                                        var v101 int32 = 449
                                        _ = v101
                                    }
                                    var v102 uint64 = 0
                                    fmt.Println(fi41, v102, v6)
                                    _ = v102
                                }
                                if ((v3 != (v3 / (((v3 << (((v5 ^ (v5 << (((v5 - (v5 * v5))) & 63)))) & 63))) | 1))) && (false && false)) {
                                    v3 -= (v3 >> (((v5 / ((2) | 1))) & 63))
                                } else {
                                    v4 += (3.14159 / v4)
                                }
                                _ = fi43
                            }
                            if false {
                                v103 := GenS11{F0: 42, F1: 552}
                                for fi104 := 0; fi104 < 13; fi104++ {
                                    var v105 int = (-(fi104))
                                    if (!true) {
                                        v5 = 255
                                    } else {
                                        v5 = v5
                                        v106 := GenS12{F0: 127}
                                        v3 -= 990
                                        _ = v106
                                    }
                                    var v107 int32 = 0
                                    _ = fi104
                                    _ = v105
                                    _ = v107
                                }
                                _ = v103
                            }
                            _ = fi41
                            _ = v42
                        }
                    case 1:
                        var v108 string = (("" + "hello") + ("hello" + ""))
                        _ = v108
                    }
                    v109 := []int{42, 128}
                    v109 = append(v109, -2147483648)
                    _ = v109
                }
            default:
                v110 := GenS5{F0: 0, F1: 42, F2: 127}
                _ = v110
            }
        case "foo bar":
            var v111 float32 = (1e10 + float32((0.1 - (float64((float32((v4 - v4)) / (0.1 / 1e10))) + v4))))
            _ = v111
        case "":
            if false {
                var v112 float32 = (3.14159 - 1.5)
                v3 = 2147483648
                _ = v112
            }
        default:
            v113 := GenS4{F0: 876, F1: 100, F2: "foo bar"}
            _ = v113
        }
        v6 = v6
    }
    var v114 uint = (v5 * 281)
    for fi115 := 0; fi115 < 6; fi115++ {
        var v116 uint8 = 128
        var v117 float64 = v4
        _ = fi115
        _ = v116
        _ = v117
    }
    fmt.Println(v114, v3)
    if true {
        var v118 uint32 = v3
        v3 = v3
        fmt.Println(v4, v118)
        _ = v118
    } else {
        fmt.Println(v114, v5, v6)
        if ((-2147483648 == 42) && true) {
            switch "\n\t" {
            case "a":
                if true {
                    v4 -= ((v4 / 0.1) - ((v4 * -2.25) + v4))
                    v4 -= 1.5
                }
            }
            switch "\n\t" {
            case "foo bar":
                var v119 int16 = (^(v6))
                _ = v119
            default:
                var v120 bool = (-2147483648 < 0)
                _ = v120
            }
        } else {
            if (2147483647 < -2147483648) {
                var v121 int = 7
                switch (v121 & v121) {
                case 4294967295:
                    if ((42 == 127) && (!false)) {
                        switch (("" + "foo bar") + ("foo bar" + "")) {
                        case "hello":
                            var v122 uint64 = 2
                            _ = v122
                        case "\n\t":
                            v123 := GenS3{F0: 1}
                            _ = v123
                        case "":
                            v6 -= v6
                        default:
                            var v124 int64 = 42
                            _ = v124
                        }
                    }
                case 127:
                    _ = len(v2)
                case -1:
                    v6 = -32768
                default:
                    v125 := []int{-984, 9223372036854775807, 600, (v121 - (v121 >> (((v114 ^ (^(v114)))) & 63)))}
                    _ = v125
                }
                v6 -= (v6 + v6)
                _ = v121
            }
        }
        for fi126 := 0; fi126 < 1; fi126++ {
            var v127 bool = (!(v4 != (((-2.25 / float64((1e10 - float32(((((v4 * ((v4 / v4) + (v4 - v4))) + float64(-2.25)) + 1e10) - v4))))) * v4) / (0.1 + ((v4 * v4) - (1e10 / (0.0 / (float64(float32(-2.25)) - v4))))))))
            v3 = 123
            _ = fi126
            _ = v127
        }
    }
    v128 := []int{128}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v114:", v114)
    fmt.Println("slicelen v128:", len(v128), v128)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v114
    _ = v128
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
