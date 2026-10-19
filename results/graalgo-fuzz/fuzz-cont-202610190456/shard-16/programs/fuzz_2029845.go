package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int }
type GenS2 struct { F0 string; F1 int32 }
type GenS3 struct { F0 int8; F1 int32 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint32; F1 uint8; F2 bool }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 uint16; F1 uint; F2 string }
type GenS8 struct { F0 uint8 }
type GenS9 struct { F0 uint16 }
type GenS10 struct { F0 int8; F1 uint; F2 int32 }
type GenS11 struct { F0 uint }
type GenS12 struct { F0 uint64; F1 uint; F2 int }
type GenS13 struct { F0 bool; F1 bool; F2 uint }
type GenS14 struct { F0 int16; F1 uint8; F2 int16 }
type GenS15 struct { F0 uint8; F1 uint16 }
func genFn0(p0 bool, p1 uint32, p2 int64) uint8 {
    var v7 int16 = 1000
    switch 1 {
    case 32768:
        p0 = p0
    case 32767:
        fmt.Println(v7, p0)
    case 247:
        for fi8 := 0; fi8 < 15; fi8++ {
            var v9 int32 = -1
            v10 := GenS1{F0: "a", F1: -451}
            _ = fi8
            _ = v10
            _ = v9
        }
    }
    v11 := GenS1{F0: "hello", F1: 32768}
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v7
    return 42
}

func main() {
    var v1 bool = true
    var v2 bool = (v1 && v1)
    v3 := make(map[uint32]uint32)
    var v4 int16 = 42
    v5 := make(map[uint32]uint32)
    v6 := make(map[uint32]string)
    for fi12 := 0; fi12 < 13; fi12++ {
        for fi13 := 0; fi13 < 15; fi13++ {
            switch (("" + ("" + ("\n\t" + "a"))) + "") {
            case "\n\t":
                v4 -= v4
            }
            v14 := GenS1{F0: "a", F1: 487}
            v2 = v1
            _ = fi13
            _ = v14
        }
        _ = fi12
    }
    v2 = (255 >= 211)
    v15 := GenS2{F0: ("" + "hello"), F1: 196}
    if v1 {
        switch "hello" {
        case "":
            for fi16 := 0; fi16 < 2; fi16++ {
                v2 = v1
                switch "" {
                case "a":
                    var v17 uint64 = 308
                    _ = v17
                case "\n\t":
                    var v18 int16 = (v4 + v4)
                    _ = v18
                case "":
                    if false {
                        if v2 {
                            v19 := []int{85, fi16}
                            if len(v19) > 0 {
                                v19[0] = (^(fi16))
                            }
                            _ = v19
                        } else {
                            var v20 uint = 421
                            var v21 uint16 = 42
                            _ = v20
                            _ = v21
                        }
                        switch ("a" + (("hello" + "\n\t") + "hello")) {
                        case "":
                            v1 = true
                        case "foo bar":
                            fmt.Println(v4)
                        default:
                            for fi22 := 0; fi22 < 3; fi22++ {
                                var v23 float32 = 1e10
                                v24 := GenS3{F0: 127, F1: 10}
                                switch 1 {
                                case 9223372036854775807:
                                    delete(v6, 2147483647)
                                case 32767:
                                    v2 = (!true)
                                case 127:
                                    var v25 int8 = -79
                                    _ = v25
                                }
                                _ = fi22
                                _ = v23
                                _ = v24
                            }
                        }
                        v4 = v4
                    } else {
                        for fi26 := 0; fi26 < 13; fi26++ {
                            switch "hello" {
                            case "":
                                v27 := GenS4{F0: (fi16 | 2147483647)}
                                _ = v27
                            case "foo bar":
                                v28 := GenS5{F0: 1, F1: 2, F2: (v1 || v2)}
                                _ = v28
                            default:
                                v1 = true
                            }
                            if true {
                                var v29 uint16 = 255
                                var v30 uint64 = 0
                                v29 = v29
                                if v1 {
                                    v1 = true
                                    fmt.Println(fi16, v2)
                                    v31 := GenS1{F0: "hello", F1: (^(fi26))}
                                    v32 := v6[1]
                                    var v33 float32 = 0.1
                                    switch 233 {
                                    case 229:
                                        if true {
                                            v34 := []int{364, fi26, fi16}
                                            var v35 string = "hello"
                                            _ = v34
                                            _ = v35
                                        } else {
                                            switch "" {
                                            case "foo bar":
                                                v36 := []int{(-(fi26)), -423, (fi16 << ((65535) & 63))}
                                                _ = v36
                                            default:
                                                fmt.Println(v30)
                                            }
                                        }
                                    case -518:
                                        fmt.Println(v2, v1, v29)
                                    case -151:
                                        delete(v6, 0)
                                        delete(v6, 661)
                                        v37 := v5[4294967295]
                                        _ = v37
                                    default:
                                        if (!v2) {
                                            var v38 int8 = -128
                                            switch (fi16 ^ fi16) {
                                            case 128:
                                                var v39 uint32 = 65535
                                                _ = v39
                                            case 32767:
                                                v40 := GenS1{F0: (("" + ("\n\t" + (v32 + (v32 + v32)))) + v32), F1: (-(fi16))}
                                                _ = v40
                                            default:
                                                v41 := GenS5{F0: 2147483647, F1: 40, F2: ((v38 | (v38 >> ((42) & 63))) <= v38)}
                                                _ = v41
                                            }
                                            _ = len(v6)
                                            _ = v38
                                        } else {
                                            if (1000 <= (v4 >> ((1) & 63))) {
                                                v42 := GenS4{F0: (fi26 - 9223372036854775807)}
                                                if (true || true) {
                                                    switch "hello" {
                                                    case "":
                                                        if true {
                                                            v43 := GenS3{F0: 0, F1: -2147483648}
                                                            _ = v43
                                                        } else {
                                                            v33 = (((v33 * 1e10) - v33) / (v33 - v33))
                                                        }
                                                    default:
                                                        switch v32 {
                                                        case "":
                                                            v44 := GenS5{F0: 4294967295, F1: 152, F2: (!v1)}
                                                            _ = v44
                                                        case "a":
                                                            v45 := GenS3{F0: 118, F1: 0}
                                                            _ = v45
                                                        case "\n\t":
                                                            var v46 bool = v1
                                                            _ = v46
                                                        }
                                                    }
                                                }
                                                var v47 int64 = -1
                                                _ = v42
                                                _ = v47
                                            }
                                        }
                                    }
                                    _ = v31
                                    _ = v32
                                    _ = v33
                                } else {
                                    fmt.Println(v4, v1, v30)
                                }
                                var v48 uint64 = 918
                                _ = v29
                                _ = v30
                                _ = v48
                            } else {
                                if ((-128 == 100) || ((false && v1) && (true || (("\n\t" <= "\n\t") || true)))) {
                                    var v49 int8 = 57
                                    var v50 uint32 = 2147483647
                                    v2 = v2
                                    _ = v49
                                    _ = v50
                                }
                                var v51 bool = true
                                _ = v51
                            }
                            _ = fi26
                        }
                    }
                default:
                    if (0 < 7) {
                        if false {
                            var v52 int16 = 25
                            v53 := GenS6{F0: 148}
                            _ = v52
                            _ = v53
                        }
                        switch fi16 {
                        case -9223372036854775808:
                            var v54 float32 = ((float32(3.14159) - (float32((3.14159 + -2.25)) + ((-2.25 - 1e10) - 0.1))) * (1.5 / 1.5))
                            _ = v54
                        case 42:
                            delete(v3, 65535)
                        case -2147483648:
                            v4 -= v4
                        }
                    }
                }
                v55 := []int{(fi16 * (fi16 & (fi16 % ((4294967295) | 1)))), 32767, (-(fi16)), (fi16 % (((fi16 | (fi16 - (fi16 << ((1) & 63))))) | 1))}
                _ = fi16
                _ = v55
            }
        case "a":
            var v56 uint8 = 0
            _ = v56
        case "hello":
            v57 := GenS4{F0: 42}
            _ = v57
        }
        switch "foo bar" {
        case "\n\t":
            var v58 uint8 = 128
            _ = v58
        case "a":
            if v1 {
                v1 = v2
                v2 = v2
            }
        case "foo bar":
            var v59 uint16 = 32767
            _ = v59
        }
    } else {
        var v60 int = 2147483647
        _ = v60
    }
    var v61 uint16 = 255
    var v62 uint16 = v61
    for fi63 := 0; fi63 < 15; fi63++ {
        var v64 int16 = (v4 / ((1000) | 1))
        v65 := GenS7{F0: (v61 >> ((245) & 63)), F1: 2, F2: ((("\n\t" + ("a" + "a")) + (((("foo bar" + ("" + "hello")) + ("" + ((("hello" + "hello") + "foo bar") + "\n\t"))) + "hello") + "foo bar")) + "\n\t")}
        _ = fi63
        _ = v64
        _ = v65
    }
    var v66 uint8 = genFn0(((-1 <= 1) && (true && v1)), 2147483648, 42)
    v67 := GenS8{F0: 127}
    if true {
        v62 -= (v61 % ((uint16(v66)) | 1))
        for fi68 := 0; fi68 < 12; fi68++ {
            v69 := v6[814]
            v70 := GenS9{F0: 1}
            _ = fi68
            _ = v69
            _ = v70
        }
        v62 -= 593
    }
    for fi71 := 0; fi71 < 11; fi71++ {
        v61 = (v61 ^ (v62 | v61))
        if (false || (100 == 100)) {
            var v72 uint16 = v61
            switch (fi71 << ((2) & 63)) {
            case 1:
                v73 := []int{fi71}
                _ = v73
            case 42:
                if true {
                    _ = len(v6)
                } else {
                    fmt.Println(v66, v2)
                }
            default:
                if false {
                    switch ("\n\t" + "hello") {
                    case "a":
                        v74 := GenS9{F0: (^(v62))}
                        _ = v74
                    case "foo bar":
                        v75 := GenS10{F0: -128, F1: 42, F2: 42}
                        _ = v75
                    default:
                        var v76 float32 = 3.14159
                        _ = v76
                    }
                    if ((((2 == 546) || (false || (!v2))) || v2) || (-1 == -118)) {
                        v72 = uint16(v66)
                    } else {
                        v77 := v6[2147483647]
                        if (4294967295 > 0) {
                            if (!(((1.5 + float64(1e10)) > (1e10 - 1e10)) && ((v61 << ((65535) & 63)) < (v62 ^ uint16(v66))))) {
                                v77 = v77
                                if v1 {
                                    if false {
                                        switch v77 {
                                        case "\n\t":
                                            v78 := GenS4{F0: (fi71 % ((fi71) | 1))}
                                            _ = v78
                                        case "":
                                            v79 := GenS6{F0: 623}
                                            _ = v79
                                        case "a":
                                            for fi80 := 0; fi80 < 12; fi80++ {
                                                v5[161] = 1
                                                _ = fi80
                                            }
                                        }
                                        delete(v3, 42)
                                        switch (-(fi71)) {
                                        case -44:
                                            switch 0 {
                                            case 1:
                                                var v81 uint8 = (v66 >> ((18446744073709551615) & 63))
                                                _ = v81
                                            }
                                        }
                                    } else {
                                        v82 := GenS11{F0: 1}
                                        _ = v82
                                    }
                                    for fi83 := 0; fi83 < 9; fi83++ {
                                        var v84 uint32 = 0
                                        v85 := GenS10{F0: -1, F1: 2, F2: 2147483647}
                                        var v86 uint8 = (v66 - (^(v66)))
                                        _ = fi83
                                        _ = v84
                                        _ = v85
                                        _ = v86
                                    }
                                    switch v77 {
                                    case "":
                                        fmt.Println(v66, fi71, fi71)
                                    case "\n\t":
                                        for fi87 := 0; fi87 < 13; fi87++ {
                                            var v88 bool = (float32(((float64(float32(0.1)) * 1e10) + 0.1)) < 1.5)
                                            v5[65535] = 65535
                                            _ = fi87
                                            _ = v88
                                        }
                                    case "foo bar":
                                        var v89 int64 = 174
                                        _ = v89
                                    }
                                }
                                v2 = true
                            } else {
                                for fi90 := 0; fi90 < 14; fi90++ {
                                    fmt.Println(v77, v1)
                                    v91 := GenS12{F0: 1, F1: 42, F2: (fi90 % (((fi90 / ((fi90) | 1))) | 1))}
                                    switch v77 {
                                    case "":
                                        if ((v4 << ((0) & 63)) > (v4 - v4)) {
                                            var v92 int16 = v4
                                            _ = v92
                                        }
                                        _ = len(v3)
                                    default:
                                        switch "a" {
                                        case "":
                                            v93 := GenS9{F0: v72}
                                            _ = v93
                                        }
                                    }
                                    _ = fi90
                                    _ = v91
                                }
                            }
                            v94 := v5[2]
                            _ = v94
                        } else {
                            v95 := GenS13{F0: v1, F1: v2, F2: 631}
                            v96 := v6[2147483647]
                            _ = v95
                            _ = v96
                        }
                        _ = v77
                    }
                } else {
                    for fi97 := 0; fi97 < 5; fi97++ {
                        v61 = 42
                        if ((((-2.25 + 1e10) / 1e10) + (((-2.25 + (((0.1 * (1.5 * (-2.25 + (3.14159 / (1.5 * ((3.14159 - float64(0.0)) * 1e10)))))) - 1e10) - ((-2.25 * (0.1 - 1.5)) - 0.1))) / -2.25) / 0.1)) != 0.0) {
                            v98 := []int{fi97, (fi97 | (fi97 >> ((2) & 63)))}
                            _ = v98
                        } else {
                            if (18446744073709551615 <= 489) {
                                v99 := []int{(fi97 | fi97), fi97, fi97}
                                if v1 {
                                    v4 = v4
                                    v2 = v2
                                }
                                _ = v99
                            } else {
                                var v100 uint32 = 4294967295
                                v101 := []int{128, int(v66), fi97}
                                var v102 uint64 = 1
                                v103 := GenS14{F0: (v4 / ((v4) | 1)), F1: (^(v66)), F2: 1000}
                                _ = v100
                                _ = v101
                                _ = v102
                                _ = v103
                            }
                            fmt.Println(v4, v2)
                            v66 += 42
                        }
                        var v104 int16 = v4
                        _ = fi97
                        _ = v104
                    }
                }
            }
            _ = v72
        } else {
            v2 = v2
        }
        _ = fi71
    }
    var v105 uint16 = v61
    for fi106 := 0; fi106 < 20; fi106++ {
        var v107 int16 = v4
        switch -688 {
        case -2147483648:
            for fi108 := 0; fi108 < 5; fi108++ {
                switch -9223372036854775808 {
                case 4294967295:
                    v109 := []int{fi108, -38, 32767, 57}
                    _ = v109
                case -9223372036854775808:
                    v110 := v5[144]
                    _ = v110
                }
                v61 += 933
                _ = fi108
            }
        case -1:
            switch -1 {
            case 277:
                if false {
                    v66 -= v66
                    v1 = v1
                    var v111 bool = (3.14159 < 3.14159)
                    _ = v111
                } else {
                    switch fi106 {
                    case -632:
                        var v112 int16 = 652
                        _ = v112
                    case -1:
                        fmt.Println(fi106)
                    case 32768:
                        if ((0 != 2147483647) && (v2 || false)) {
                            fmt.Println(v1, v2, v4)
                            var v113 string = ("a" + ("a" + ""))
                            v114 := GenS12{F0: 18446744073709551615, F1: 2, F2: 42}
                            _ = v113
                            _ = v114
                        }
                    }
                    fmt.Println(v4)
                }
            case -67:
                for fi115 := 0; fi115 < 4; fi115++ {
                    v116 := GenS15{F0: (v66 / ((genFn0((-128 < 0), 2, 326)) | 1)), F1: v62}
                    delete(v3, 168)
                    var v117 int32 = -127
                    if false {
                        v62 += v62
                        var v118 int8 = -61
                        _ = v118
                    } else {
                        fmt.Println(v61, v4, fi106)
                        var v119 int64 = -1
                        fmt.Println(v1)
                        _ = v119
                    }
                    _ = fi115
                    _ = v116
                    _ = v117
                }
            }
        case 2147483647:
            v2 = v2
        }
        _ = fi106
        _ = v107
    }
    if (((2147483647 >= int(v66)) && (true || ((v4 << ((20) & 63)) != v4))) && (18446744073709551615 != 4294967295)) {
        var v120 uint32 = 974
        var v121 uint8 = 255
        _ = v120
        _ = v121
    } else {
        v105 -= 1
        v122 := []int{127, 2147483647}
        var v123 int64 = 9223372036854775807
        _ = v122
        _ = v123
    }
    fmt.Println("v1:", v1)
    fmt.Println("v105:", v105)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v61:", v61)
    fmt.Println("v62:", v62)
    fmt.Println("v66:", v66)
    _ = v1
    _ = v105
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v61
    _ = v62
    _ = v66
    _ = v67
}
