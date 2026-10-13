package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint16; F1 int64; F2 int32 }
type GenS6 struct { F0 int32; F1 int16 }
type GenS7 struct { F0 uint32; F1 uint16 }
type GenS8 struct { F0 uint8 }
type GenS9 struct { F0 int; F1 uint8 }
type GenS10 struct { F0 int8; F1 uint16 }
type GenS11 struct { F0 uint32; F1 string }
type GenS12 struct { F0 int }
type GenS13 struct { F0 string; F1 uint64; F2 int32 }
type GenS14 struct { F0 int16 }
type GenS15 struct { F0 uint; F1 uint64 }
type GenS16 struct { F0 uint32; F1 int8 }
type GenS17 struct { F0 int64 }
func genFn0(p0 int16, p1 uint32, p2 int) uint32 {
    v7 := []int{p2, 42, (p2 + (p2 << ((65535) & 63)))}
    fmt.Println(p0, p1)
    if len(v7) > 0 {
        v7[0] = 32767
    }
    for fi8 := 0; fi8 < 15; fi8++ {
        var v9 uint64 = 2
        var v10 uint = 2
        v7 = append(v7, int(p0))
        _ = fi8
        _ = v10
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return (p1 | p1)
}
func genFn1(p0 uint8, p1 int16, p2 int) int8 {
    var v11 bool = (1e10 == 0.0)
    var v12 int32 = 42
    for fi13 := 0; fi13 < 4; fi13++ {
        var v14 float64 = 1.5
        v15 := GenS1{F0: (v12 << ((255) & 63)), F1: 100}
        _ = fi13
        _ = v14
        _ = v15
    }
    switch 127 {
    case 2147483647:
        var v16 uint8 = (p0 + p0)
        _ = v16
    case -9223372036854775808:
        var v17 uint16 = 432
        _ = v17
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    return 0
}

func main() {
    v1 := []uint32{4294967295}
    var v2 uint = 2
    var v3 uint64 = 2
    v4 := make(map[int]uint32)
    v5 := make(map[int]string)
    v6 := make(map[int]string)
    switch "a" {
    case "hello":
        var v18 int8 = 1
        _ = v18
    case "a":
        for fi19 := 0; fi19 < 1; fi19++ {
            for fi20 := 0; fi20 < 20; fi20++ {
                if (false || (94 != 100)) {
                    delete(v4, fi20)
                    v3 += (v3 & v3)
                } else {
                    v2 += uint(v3)
                    v1 = append(v1, 2147483647)
                    v3 += v3
                    v21 := v6[int(v2)]
                    _ = v21
                }
                _ = fi20
            }
            _ = fi19
        }
    }
    var v22 uint32 = 1
    for fi23 := 0; fi23 < 2; fi23++ {
        switch "" {
        case "foo bar":
            v24 := GenS2{F0: -2147483648}
            _ = v24
        case "\n\t":
            _ = len(v1)
        }
        v25 := GenS3{F0: false}
        _ = len(v4)
        _ = fi23
        _ = v25
    }
    if false {
        v2 = 18446744073709551615
        v5[905] = "foo bar"
        v2 += v2
    }
    if (1.5 <= 1e10) {
        v22 = v22
        if (!(true || true)) {
            for fi26 := 0; fi26 < 18; fi26++ {
                if (((0.1 / -2.25) - (3.14159 * (-2.25 - 1e10))) != (1e10 + 1.5)) {
                    switch "hello" {
                    case "\n\t":
                        if ((0 != (v22 ^ v22)) && ((^(v2)) <= (v2 & (v2 | 802)))) {
                            for fi27 := 0; fi27 < 6; fi27++ {
                                v2 -= (v2 * v2)
                                fmt.Println(v2, v2)
                                var v28 uint32 = (v22 / ((v22) | 1))
                                _ = fi27
                                _ = v28
                            }
                            switch (-(fi26)) {
                            case 4294967295:
                                for fi29 := 0; fi29 < 12; fi29++ {
                                    switch 722 {
                                    case 32767:
                                        v2 = (^(v2))
                                    case -469:
                                        var v30 float32 = float32((0.1 / 3.14159))
                                        _ = v30
                                    case 1:
                                        v3 += v3
                                    default:
                                        if true {
                                            if (!(false && ((false && (414 < 774)) || (1 != 2147483647)))) {
                                                if (((fi29 << ((v2) & 63)) == (fi29 ^ fi29)) && ((true && (2 >= 0)) && ("\n\t" != "foo bar"))) {
                                                    if (-2.25 < -2.25) {
                                                        if true {
                                                            if false {
                                                                if true {
                                                                    switch 4294967295 {
                                                                    case 0:
                                                                        if false {
                                                                            var v31 int32 = 42
                                                                            v2 += v2
                                                                            var v32 int32 = (v31 ^ v31)
                                                                            _ = v31
                                                                            _ = v32
                                                                        }
                                                                    }
                                                                }
                                                                v33 := GenS4{F0: (fi26 >> ((18446744073709551615) & 63))}
                                                                _ = v33
                                                            } else {
                                                                fmt.Println(fi29, fi29)
                                                                if len(v1) >= 2 {
                                                                    v1 = v1[1:len(v1)]
                                                                }
                                                                _ = len(v5)
                                                            }
                                                            switch fi26 {
                                                            case 0:
                                                                var v34 int16 = 768
                                                                _ = v34
                                                            }
                                                            fmt.Println(v2, v22, fi26)
                                                        }
                                                    } else {
                                                        var v35 bool = false
                                                        v22 -= (v22 + (v22 % (((v22 ^ (^(v22)))) | 1)))
                                                        var v36 uint32 = (^(v22))
                                                        _ = v35
                                                        _ = v36
                                                    }
                                                }
                                                var v37 float64 = 0.0
                                                delete(v6, fi26)
                                                _ = v37
                                            }
                                            v22 -= (v22 % ((v22) | 1))
                                            fmt.Println(v2, v22, v2)
                                        }
                                    }
                                    _ = fi29
                                }
                            case 127:
                                var v38 int = int(v3)
                                _ = v38
                            case 32768:
                                if ((v3 | (v3 - 0)) < v3) {
                                    if (-821 != 1) {
                                        fmt.Println(v3, v22, v3)
                                        v2 = (^(v2))
                                        v39 := GenS5{F0: 65535, F1: 9223372036854775807, F2: 42}
                                        _ = v39
                                    }
                                } else {
                                    v22 = genFn0(1, v22, (^(fi26)))
                                    for fi40 := 0; fi40 < 13; fi40++ {
                                        _ = len(v5)
                                        v41 := GenS6{F0: 42, F1: -408}
                                        v42 := GenS1{F0: 115, F1: -121}
                                        _ = fi40
                                        _ = v41
                                        _ = v42
                                    }
                                }
                            default:
                                _ = len(v1)
                            }
                            v3 = v3
                        }
                    case "foo bar":
                        if false {
                            if (!false) {
                                var v43 uint = v2
                                _ = v43
                            }
                        } else {
                            v44 := GenS7{F0: (v22 / (((v22 & v22)) | 1)), F1: 404}
                            _ = v44
                        }
                    case "a":
                        v45 := GenS8{F0: 2}
                        _ = v45
                    default:
                        for fi46 := 0; fi46 < 13; fi46++ {
                            v2 = v2
                            _ = fi46
                        }
                    }
                    v47 := GenS9{F0: 127, F1: 89}
                    _ = v47
                } else {
                    switch fi26 {
                    case 9223372036854775807:
                        v3 -= v3
                    }
                }
                var v48 uint64 = (^(v3))
                switch "hello" {
                case "hello":
                    if true {
                        for fi49 := 0; fi49 < 20; fi49++ {
                            var v50 int = -1
                            fmt.Println(v50, v50, v22)
                            _ = fi49
                            _ = v50
                        }
                        _ = len(v6)
                    }
                case "\n\t":
                    var v51 float64 = 3.14159
                    _ = v51
                case "":
                    var v52 int64 = 9223372036854775807
                    _ = v52
                }
                _ = fi26
                _ = v48
            }
            v22 = genFn0(32767, (^(v22)), 1)
            v53 := GenS8{F0: 91}
            _ = v53
        } else {
            v2 += (v2 | 65535)
            v22 -= genFn0(-32768, 2, 9223372036854775807)
            switch "" {
            case "":
                switch "\n\t" {
                case "\n\t":
                    if ((!(true && (-215 > 0))) && true) {
                        if (42 <= -128) {
                            v54 := GenS2{F0: 0}
                            _ = v54
                        }
                    }
                case "a":
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                }
            case "foo bar":
                v55 := GenS4{F0: 32767}
                _ = v55
            case "a":
                v22 = v22
            }
        }
    } else {
        fmt.Println(v22)
        for fi56 := 0; fi56 < 14; fi56++ {
            var v57 int = 127
            if false {
                if false {
                    v3 += v3
                    v2 -= (v2 & v2)
                    v2 -= v2
                } else {
                    for fi58 := 0; fi58 < 11; fi58++ {
                        v22 = v22
                        if false {
                            if true {
                                v59 := GenS10{F0: 127, F1: 0}
                                v3 = (v3 ^ v3)
                                v2 -= v2
                                _ = v59
                            }
                            v6[(fi58 * (^(fi56)))] = ("foo bar" + ("a" + "foo bar"))
                            fmt.Println(v57, v3)
                            if (!false) {
                                fmt.Println(v3)
                                if true {
                                    if (false || (false || false)) {
                                        switch (fi56 * (fi58 / (((fi58 - -2147483648)) | 1))) {
                                        case 32767:
                                            var v60 uint = 450
                                            _ = v60
                                        case 32768:
                                            v57 = (-(fi58))
                                        case 444:
                                            var v61 uint64 = 4294967295
                                            _ = v61
                                        }
                                    } else {
                                        if false {
                                            v57 -= fi56
                                        } else {
                                            if true {
                                                if (9223372036854775807 != -113) {
                                                    switch 9223372036854775807 {
                                                    case 511:
                                                        var v62 uint = (v2 / ((v2) | 1))
                                                        _ = v62
                                                    case 1:
                                                        fmt.Println(v22)
                                                    }
                                                } else {
                                                    v1 = append(v1, 42)
                                                    v63 := GenS11{F0: 2, F1: "a"}
                                                    if (!(false || (!((100 >= -1) && false)))) {
                                                        v3 -= (^(v3))
                                                        var v64 uint8 = 167
                                                        var v65 uint = (v2 - (v2 >> (((v2 % ((42) | 1))) & 63)))
                                                        _ = v64
                                                        _ = v65
                                                    } else {
                                                        v57 = 1
                                                        switch ("" + "\n\t") {
                                                        case "hello":
                                                            var v66 uint = (^(v2))
                                                            _ = v66
                                                        case "\n\t":
                                                            v3 = v3
                                                        case "foo bar":
                                                            v3 = (v3 / ((4294967295) | 1))
                                                        default:
                                                            fmt.Println(v3, fi58, fi58)
                                                            if true {
                                                                v67 := v4[v57]
                                                                v68 := GenS12{F0: (-(fi56))}
                                                                fmt.Println(v2, v67, v22)
                                                                _ = v67
                                                                _ = v68
                                                            }
                                                            _ = len(v4)
                                                            var v69 string = "\n\t"
                                                            _ = v69
                                                        }
                                                        v70 := GenS3{F0: ((true || false) && (100 < -1))}
                                                        _ = v70
                                                    }
                                                    var v71 int32 = -689
                                                    if (128 > 127) {
                                                        v57 = (fi56 >> ((uint(v22)) & 63))
                                                    } else {
                                                        if false {
                                                            var v72 string = ("foo bar" + ("a" + ("\n\t" + "hello")))
                                                            switch (v72 + "hello") {
                                                            case "a":
                                                                if ((v3 % ((v3) | 1)) <= v3) {
                                                                    var v73 bool = false
                                                                    v74 := GenS9{F0: (fi56 / ((fi58) | 1)), F1: 127}
                                                                    v2 -= 18446744073709551615
                                                                    _ = v73
                                                                    _ = v74
                                                                } else {
                                                                    v75 := GenS13{F0: "\n\t", F1: 2, F2: int32(v2)}
                                                                    _ = v75
                                                                }
                                                            }
                                                            _ = v72
                                                        }
                                                        var v76 int = (^(v57))
                                                        var v77 uint64 = 2
                                                        var v78 uint32 = v22
                                                        if ((42 != 1) || false) {
                                                            switch v57 {
                                                            case -420:
                                                                var v79 string = (("hello" + "foo bar") + ("\n\t" + "a"))
                                                                _ = v79
                                                            case -9223372036854775808:
                                                                var v80 int16 = 0
                                                                _ = v80
                                                            default:
                                                                var v81 uint = (v2 << ((uint(v77)) & 63))
                                                                _ = v81
                                                            }
                                                            var v82 uint16 = 255
                                                            switch "foo bar" {
                                                            case "foo bar":
                                                                var v83 int8 = 127
                                                                _ = v83
                                                            case "":
                                                                v82 -= (v82 ^ 2)
                                                            default:
                                                                switch (^(v57)) {
                                                                case 0:
                                                                    v22 += v22
                                                                case 775:
                                                                    v84 := GenS8{F0: 42}
                                                                    _ = v84
                                                                default:
                                                                    v76 = (v76 + (v57 * (fi56 - 127)))
                                                                }
                                                            }
                                                            _ = v82
                                                        } else {
                                                            v2 = v2
                                                        }
                                                        var v85 int64 = 579
                                                        _ = v76
                                                        _ = v77
                                                        _ = v78
                                                        _ = v85
                                                    }
                                                    _ = v63
                                                    _ = v71
                                                }
                                                var v86 int32 = 42
                                                var v87 int32 = (v86 | (v86 & v86))
                                                _ = v86
                                                _ = v87
                                            } else {
                                                v3 -= v3
                                                var v88 int32 = 2147483647
                                                v89 := GenS14{F0: 1000}
                                                _ = v88
                                                _ = v89
                                            }
                                            v1 = append(v1, genFn0(1, (v22 >> (((v2 / (((v2 - v2)) | 1))) & 63)), 32768))
                                            switch ("" + ("foo bar" + ("a" + ""))) {
                                            case "a":
                                                switch (fi58 | v57) {
                                                case 128:
                                                    switch ("hello" + "foo bar") {
                                                    case "hello":
                                                        if true {
                                                            v1 = append(v1, (v22 >> (((v2 >> (((^(v2))) & 63))) & 63)))
                                                        } else {
                                                            delete(v5, v57)
                                                            var v90 float64 = float64(0.1)
                                                            fmt.Println(v57, v57)
                                                            _ = v90
                                                        }
                                                    case "foo bar":
                                                        switch fi58 {
                                                        case 2147483647:
                                                            if true {
                                                                if len(v1) >= 2 {
                                                                    v1 = v1[1:len(v1)]
                                                                }
                                                                v22 = genFn0(726, 560, 42)
                                                            } else {
                                                                var v91 uint8 = 195
                                                                v92 := GenS6{F0: 0, F1: 32767}
                                                                _ = v91
                                                                _ = v92
                                                            }
                                                        default:
                                                            if (2147483647 != -1) {
                                                                fmt.Println(v22, fi58, v2)
                                                                if (!false) {
                                                                    if (1 < 175) {
                                                                        v2 += v2
                                                                        delete(v6, (fi56 / (((fi58 & -1)) | 1)))
                                                                    }
                                                                } else {
                                                                    v22 += (^(v22))
                                                                    v1 = append(v1, v22)
                                                                }
                                                                var v93 int64 = -9223372036854775808
                                                                var v94 int = v57
                                                                _ = v93
                                                                _ = v94
                                                            }
                                                        }
                                                    default:
                                                        v95 := GenS10{F0: 100, F1: 32767}
                                                        _ = v95
                                                    }
                                                case 2147483647:
                                                    v96 := GenS15{F0: (^(v2)), F1: (v3 % (((v3 << (((v2 - v2)) & 63))) | 1))}
                                                    _ = v96
                                                case 1:
                                                    var v97 uint32 = 65535
                                                    _ = v97
                                                }
                                            case "":
                                                if ((true || (!false)) && false) {
                                                    v2 -= v2
                                                    fmt.Println(v22)
                                                    switch "hello" {
                                                    case "":
                                                        var v98 int64 = 1
                                                        _ = v98
                                                    case "a":
                                                        fmt.Println(v22, fi58)
                                                    case "\n\t":
                                                        if (!((v22 << ((v2) & 63)) > v22)) {
                                                            var v99 string = "hello"
                                                            _ = v99
                                                        } else {
                                                            var v100 string = ("a" + "")
                                                            v101 := GenS7{F0: (v22 % (((v22 - v22)) | 1)), F1: 0}
                                                            _ = v100
                                                            _ = v101
                                                        }
                                                    }
                                                } else {
                                                    switch 0 {
                                                    case 128:
                                                        v57 = (v57 / (((fi56 >> (((v2 >> ((v2) & 63))) & 63))) | 1))
                                                    case 0:
                                                        var v102 int64 = 155
                                                        _ = v102
                                                    }
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                    var v103 uint8 = 42
                                                    _ = v103
                                                }
                                            case "foo bar":
                                                if true {
                                                    v1 = append(v1, (^(v22)))
                                                    _ = len(v5)
                                                    var v104 uint8 = 127
                                                    _ = v104
                                                } else {
                                                    var v105 int8 = -1
                                                    _ = v105
                                                }
                                            default:
                                                fmt.Println(v2, fi58, v2)
                                            }
                                            var v106 uint16 = 65535
                                            if true {
                                                fmt.Println(fi56, v22, v3)
                                                var v107 float32 = 0.0
                                                v108 := GenS11{F0: v22, F1: (("" + "\n\t") + "foo bar")}
                                                fmt.Println(fi56)
                                                var v109 int64 = -68
                                                _ = v107
                                                _ = v108
                                                _ = v109
                                            } else {
                                                v2 = v2
                                                fmt.Println(fi56)
                                                fmt.Println(v22, v22, fi56)
                                                v3 = v3
                                            }
                                            v22 += (v22 + v22)
                                            var v110 float32 = 0.1
                                            v2 = 18446744073709551615
                                            _ = v106
                                            _ = v110
                                        }
                                    }
                                }
                                v57 -= fi56
                                v111 := GenS12{F0: -719}
                                var v112 string = "\n\t"
                                v2 = v2
                                _ = v111
                                _ = v112
                            } else {
                                _ = len(v1)
                                var v113 uint64 = (v3 + v3)
                                _ = v113
                            }
                        }
                        fmt.Println(v2, v2, v22)
                        if true {
                            var v114 float64 = 0.1
                            _ = v114
                        }
                        delete(v4, (fi58 / ((fi56) | 1)))
                        _ = fi58
                    }
                }
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                v3 += 1
            }
            for fi115 := 0; fi115 < 14; fi115++ {
                v6[fi115] = "\n\t"
                _ = fi115
            }
            _ = fi56
            _ = v57
        }
    }
    v22 = 1
    for fi116 := 0; fi116 < 16; fi116++ {
        v117 := GenS16{F0: v22, F1: 64}
        v2 = (^(v2))
        _ = fi116
        _ = v117
    }
    if len(v1) > 0 {
        v1[0] = (^(v22))
    }
    v22 = (v22 >> ((255) & 63))
    switch 9223372036854775807 {
    case 0:
        v22 += genFn0(708, (v22 + 65535), 128)
    case 32767:
        v22 -= (v22 << (((v2 ^ uint(v3))) & 63))
    }
    v118 := GenS17{F0: -9223372036854775808}
    var v119 string = ""
    var v120 float64 = 0.0
    v3 += (^(v3))
    v22 = 408
    v2 += v2
    delete(v4, -9223372036854775808)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v119:", v119)
    fmt.Println("v120:", v120)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v118
    _ = v119
    _ = v120
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
