package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64; F2 int32 }
type GenS2 struct { F0 int; F1 uint32 }
type GenS3 struct { F0 uint64; F1 int64; F2 int32 }
type GenS4 struct { F0 int; F1 uint }
type GenS5 struct { F0 int64; F1 string }
type GenS6 struct { F0 int8; F1 int32; F2 int64 }
type GenS7 struct { F0 uint; F1 int8; F2 int64 }
type GenS8 struct { F0 uint; F1 bool; F2 uint32 }
type GenS9 struct { F0 uint8 }
type GenS10 struct { F0 uint32; F1 int8; F2 uint64 }
type GenS11 struct { F0 int16; F1 uint64; F2 int16 }
type GenS12 struct { F0 uint32; F1 int64 }
type GenS13 struct { F0 int32 }
type GenS14 struct { F0 uint16; F1 uint16; F2 int }
type GenS15 struct { F0 uint32; F1 int8 }
func genFn0(p0 uint8, p1 bool) uint {
    p0 = p0
    p1 = p1
    var v5 uint8 = (p0 & (^(p0)))
    switch 419 {
    case 1:
        p0 -= v5
    }
    p1 = ((0 != -1) && (((-1 >= -581) || (!p1)) && (!p1)))
    _ = p0
    _ = p1
    _ = v5
    return 2
}
func genFn1() int64 {
    var v6 uint64 = 204
    if (!(float32(float64((1.5 / 0.1))) == float32(3.14159))) {
        v7 := GenS1{F0: -1, F1: -9223372036854775808, F2: 0}
        var v8 uint32 = 0
        _ = v7
        _ = v8
    } else {
        var v9 uint8 = 255
        v10 := []int{9223372036854775807, 42, 9223372036854775807, 790}
        v6 = 1
        _ = v10
        _ = v9
    }
    for fi11 := 0; fi11 < 9; fi11++ {
        for fi12 := 0; fi12 < 11; fi12++ {
            v6 -= v6
            v6 -= 2
            _ = fi12
        }
        _ = fi11
    }
    v6 -= (v6 + (v6 * v6))
    _ = v6
    return -1
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[uint32]uint32)
    v3 := []uint32{2, 693, 2147483647}
    v4 := make(map[uint32]uint32)
    v13 := GenS1{F0: 2147483647, F1: 0, F2: 0}
    switch 128 {
    case -9223372036854775808:
        v14 := GenS1{F0: 2147483647, F1: 42, F2: 42}
        _ = v14
    default:
        if false {
            var v15 int64 = genFn1()
            var v16 int32 = 1
            _ = v15
            _ = v16
        } else {
            var v17 float64 = 0.0
            switch "\n\t" {
            case "a":
                var v18 int8 = 100
                _ = v18
            case "hello":
                delete(v1, 1)
            }
            var v19 float32 = 0.1
            _ = v17
            _ = v19
        }
    }
    var v20 string = (("\n\t" + ("foo bar" + "a")) + "hello")
    if false {
        if len(v3) > 0 {
            v3[0] = 536
        }
        var v21 uint8 = 2
        _ = v21
    } else {
        if false {
            fmt.Println(v20)
            v20 = ("a" + v20)
            if false {
                v22 := GenS1{F0: -1, F1: 9223372036854775807, F2: 1}
                v20 = v20
                _ = len(v3)
                _ = v22
            }
        } else {
            switch v20 {
            case "foo bar":
                v23 := GenS2{F0: 204, F1: 2147483647}
                _ = v23
            }
            v20 = "hello"
        }
    }
    fmt.Println(v20)
    fmt.Println(v20)
    v20 = (v20 + v20)
    _ = len(v2)
    v24 := GenS3{F0: 18446744073709551615, F1: genFn1(), F2: 1}
    v25 := GenS3{F0: 2, F1: 921, F2: 497}
    for fi26 := 0; fi26 < 4; fi26++ {
        if (v20 > (((v20 + (v20 + v20)) + "") + ("" + v20))) {
            var v27 uint16 = 774
            _ = v27
        } else {
            fmt.Println(fi26, fi26)
            switch ("foo bar" + ("foo bar" + "a")) {
            case "foo bar":
                for fi28 := 0; fi28 < 11; fi28++ {
                    var v29 uint = 450
                    var v30 int32 = -477
                    _ = len(v4)
                    _ = fi28
                    _ = v29
                    _ = v30
                }
            case "a":
                fmt.Println(fi26, v20)
            case "":
                if false {
                    v20 = v20
                } else {
                    if false {
                        var v31 int8 = 1
                        _ = v31
                    }
                    delete(v1, fi26)
                    var v32 int8 = int8(fi26)
                    _ = v32
                }
            }
        }
        for fi33 := 0; fi33 < 8; fi33++ {
            if false {
                v34 := GenS4{F0: (fi26 ^ fi33), F1: 255}
                _ = v34
            }
            var v35 int = (fi26 & 491)
            v20 = ("" + "foo bar")
            v36 := GenS3{F0: 1, F1: -1, F2: 367}
            if (!(!false)) {
                v35 -= fi26
                v37 := GenS1{F0: -336, F1: 362, F2: -2147483648}
                if true {
                    v38 := GenS5{F0: 0, F1: (("\n\t" + "hello") + (v20 + "a"))}
                    if false {
                        if (!false) {
                            if false {
                                v39 := v4[65535]
                                _ = v39
                            }
                            v40 := GenS6{F0: -112, F1: -1, F2: -1}
                            fmt.Println(fi33, v35, v35)
                            _ = v40
                        }
                    }
                    if true {
                        if false {
                            var v41 int32 = 1
                            v3 = append(v3, 2147483648)
                            _ = v41
                        }
                        switch "foo bar" {
                        case "\n\t":
                            v1[v35] = v20
                        }
                        _ = len(v4)
                    }
                    v42 := v4[493]
                    if false {
                        var v43 string = v20
                        fmt.Println(v35)
                        _ = v43
                    } else {
                        var v44 float32 = 1.5
                        if (!(false && (!false))) {
                            switch (v20 + v20) {
                            case "foo bar":
                                var v45 int32 = -1
                                _ = v45
                            case "\n\t":
                                if false {
                                    var v46 float32 = v44
                                    _ = v46
                                }
                            default:
                                var v47 bool = false
                                _ = v47
                            }
                        } else {
                            v4[(^(v42))] = v42
                            if (false || ((100 < -102) || true)) {
                                var v48 uint8 = 2
                                switch (v35 - fi26) {
                                case 9223372036854775807:
                                    var v49 int64 = genFn1()
                                    _ = v49
                                case 4294967295:
                                    var v50 int8 = 42
                                    _ = v50
                                case -1:
                                    if (!false) {
                                        fmt.Println(v42, v35, v48)
                                    }
                                default:
                                    if true {
                                        v48 += (v48 + v48)
                                        var v51 bool = false
                                        _ = v51
                                    }
                                }
                                if true {
                                    switch "a" {
                                    case "":
                                        switch fi26 {
                                        case 9223372036854775807:
                                            if (false || false) {
                                                var v52 bool = true
                                                if v52 {
                                                    var v53 int = -2147483648
                                                    v54 := GenS4{F0: (v53 << ((42) & 63)), F1: 18446744073709551615}
                                                    v55 := v4[v42]
                                                    _ = v53
                                                    _ = v54
                                                    _ = v55
                                                } else {
                                                    v1[(fi26 & 4294967295)] = (v20 + v20)
                                                    fmt.Println(v42)
                                                    var v56 int64 = genFn1()
                                                    _ = v56
                                                }
                                                delete(v1, fi26)
                                                _ = v52
                                            }
                                        case 776:
                                            if (!false) {
                                                if (!(false || false)) {
                                                    v44 = float32(1e10)
                                                } else {
                                                    fmt.Println(v48, fi33)
                                                    v57 := GenS7{F0: 2, F1: 0, F2: 666}
                                                    v58 := GenS2{F0: fi26, F1: (v42 | v42)}
                                                    _ = v57
                                                    _ = v58
                                                }
                                                var v59 uint32 = v42
                                                _ = v59
                                            }
                                        default:
                                            var v60 int8 = 98
                                            _ = v60
                                        }
                                        v61 := GenS5{F0: -1, F1: ((v20 + v20) + v20)}
                                        _ = v61
                                    case "a":
                                        v44 += v44
                                    case "hello":
                                        var v62 bool = (false && (-909 == 308))
                                        var v63 uint = 65535
                                        v64 := GenS6{F0: -1, F1: -2147483648, F2: genFn1()}
                                        _ = v62
                                        _ = v63
                                        _ = v64
                                    default:
                                        v3 = append(v3, v42)
                                    }
                                }
                                _ = v48
                            } else {
                                v20 = (("\n\t" + "a") + v20)
                            }
                            v42 -= v42
                            if true {
                                v20 = ("hello" + (v20 + ("hello" + "foo bar")))
                            } else {
                                _ = len(v1)
                                _ = len(v3)
                                var v65 int64 = 0
                                if len(v3) > 0 {
                                    v3[0] = (v42 & v42)
                                }
                                switch -1 {
                                case -338:
                                    v35 = (fi26 / ((fi33) | 1))
                                }
                                _ = v65
                            }
                        }
                        if false {
                            var v66 float64 = float64(float32((0.1 * 0.1)))
                            _ = v66
                        } else {
                            v20 = (((v20 + v20) + v20) + v20)
                            if len(v3) > 0 {
                                v3[0] = v42
                            }
                            _ = len(v3)
                        }
                        var v67 uint8 = 128
                        v68 := GenS3{F0: 832, F1: 75, F2: -681}
                        _ = v44
                        _ = v67
                        _ = v68
                    }
                    var v69 int64 = 9223372036854775807
                    if (0.1 >= 0.0) {
                        var v70 uint64 = 1
                        v71 := GenS2{F0: 128, F1: v42}
                        switch ((((v20 + v20) + ("foo bar" + "\n\t")) + v20) + ((v20 + ((("foo bar" + v20) + "a") + ((v20 + ("\n\t" + v20)) + v20))) + v20)) {
                        case "a":
                            if false {
                                v72 := GenS6{F0: 0, F1: 0, F2: (v69 & v69)}
                                v69 += genFn1()
                                var v73 uint64 = (^(v70))
                                _ = v72
                                _ = v73
                            }
                        case "hello":
                            var v74 uint64 = 100
                            _ = v74
                        default:
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        }
                        v20 = v20
                        if true {
                            v75 := GenS8{F0: 0, F1: ((v69 <= genFn1()) && (!true)), F2: v42}
                            _ = v75
                        }
                        v35 = v35
                        _ = v70
                        _ = v71
                    }
                    if true {
                        switch -999 {
                        case -1:
                            if false {
                                v42 = 1
                            } else {
                                v76 := GenS4{F0: -578, F1: 18446744073709551615}
                                if len(v3) > 0 {
                                    v3[0] = 2147483648
                                }
                                v77 := GenS6{F0: -86, F1: 256, F2: 9223372036854775807}
                                _ = v76
                                _ = v77
                            }
                            switch ((v20 + v20) + v20) {
                            case "":
                                v69 += genFn1()
                            case "\n\t":
                                fmt.Println(v42, fi26)
                            case "a":
                                switch fi26 {
                                case 9223372036854775807:
                                    v35 -= 378
                                case 575:
                                    switch 32767 {
                                    case -445:
                                        if false {
                                            var v78 int16 = -32768
                                            var v79 int8 = 1
                                            _ = v78
                                            _ = v79
                                        }
                                    case 32767:
                                        if false {
                                            delete(v1, -1)
                                            if len(v3) > 0 {
                                                v3[0] = 2147483648
                                            }
                                            v42 -= 4294967295
                                        } else {
                                            v69 = genFn1()
                                            var v80 float32 = 1.5
                                            var v81 string = v20
                                            _ = v80
                                            _ = v81
                                        }
                                    default:
                                        if ((!(530 >= -985)) || (121 > 1)) {
                                            var v82 string = v20
                                            if (genFn1() != (v69 & v69)) {
                                                if false {
                                                    v82 = v20
                                                    var v83 int32 = -709
                                                    if true {
                                                        _ = len(v3)
                                                        fmt.Println(v82, fi33)
                                                        var v84 float32 = 1.5
                                                        var v85 uint = genFn0(128, (0 != 42))
                                                        if (!(42 < (v42 >> ((v85) & 63)))) {
                                                            switch fi33 {
                                                            case 9223372036854775807:
                                                                v42 += (v42 % (((v42 ^ uint32(v85))) | 1))
                                                            case 1:
                                                                var v86 uint8 = 127
                                                                _ = v86
                                                            default:
                                                                var v87 uint16 = 32767
                                                                _ = v87
                                                            }
                                                        }
                                                        _ = v84
                                                        _ = v85
                                                    }
                                                    _ = v83
                                                }
                                                var v88 float32 = 1.5
                                                v89 := GenS9{F0: 211}
                                                _ = v88
                                                _ = v89
                                            } else {
                                                v35 = fi26
                                                v90 := GenS10{F0: (v42 % (((v42 / (((^(v42))) | 1))) | 1)), F1: 28, F2: 18446744073709551615}
                                                _ = v90
                                            }
                                            _ = v82
                                        }
                                    }
                                }
                            }
                            v3 = append(v3, (v42 % (((v42 ^ (v42 * 755))) | 1)))
                        case -9223372036854775808:
                            var v91 string = v20
                            _ = v91
                        case 9223372036854775807:
                            v20 = v20
                        }
                        v3 = append(v3, v42)
                    } else {
                        var v92 float32 = 1e10
                        var v93 uint64 = 4294967295
                        _ = v92
                        _ = v93
                    }
                    v20 = v20
                    _ = v38
                    _ = v42
                    _ = v69
                } else {
                    var v94 int16 = 32767
                    _ = v94
                }
                _ = v37
            }
            v35 = fi26
            var v95 int16 = -32768
            if true {
                v35 -= (^(fi33))
                v20 = (v20 + ("\n\t" + v20))
                switch (fi33 / ((fi26) | 1)) {
                case 9223372036854775807:
                    var v96 int32 = -2147483648
                    _ = v96
                case 32767:
                    if (1.5 <= 1e10) {
                        delete(v1, (fi33 & fi26))
                        if (3.14159 != 3.14159) {
                            v20 = "a"
                        }
                        switch (fi33 << ((952) & 63)) {
                        case 42:
                            v95 -= (v95 - 32767)
                        case 32768:
                            v97 := GenS8{F0: 2, F1: (452 == 42), F2: 42}
                            _ = v97
                        case 852:
                            fmt.Println(v95, v20, v35)
                        default:
                            if (42 < 42) {
                                if ((!true) || false) {
                                    v35 = (-(fi33))
                                    var v98 int = fi33
                                    _ = v98
                                }
                                v99 := GenS5{F0: 762, F1: v20}
                                _ = v99
                            } else {
                                switch ((v20 + v20) + v20) {
                                case "foo bar":
                                    if (61 == 42) {
                                        v100 := GenS7{F0: 2, F1: 0, F2: -1}
                                        var v101 float64 = 1.5
                                        v35 -= v35
                                        switch "foo bar" {
                                        case "hello":
                                            var v102 uint16 = 65535
                                            _ = v102
                                        case "a":
                                            if (100 == 100) {
                                                var v103 uint32 = 279
                                                _ = v103
                                            }
                                            v104 := GenS11{F0: (-(v95)), F1: 0, F2: 42}
                                            var v105 int = fi33
                                            v105 = (v35 & (fi33 << ((65535) & 63)))
                                            if (32767 == 65535) {
                                                v106 := GenS12{F0: 2, F1: 9223372036854775807}
                                                _ = v106
                                            }
                                            _ = v104
                                            _ = v105
                                        case "":
                                            if len(v3) > 0 {
                                                v3[0] = 2
                                            }
                                        }
                                        _ = v100
                                        _ = v101
                                    } else {
                                        switch "hello" {
                                        case "":
                                            v20 = v20
                                        case "\n\t":
                                            if ((176 == 0) || ((fi33 & -1) < 128)) {
                                                v107 := GenS13{F0: 2147483647}
                                                fmt.Println(fi33)
                                                v108 := GenS14{F0: 1, F1: 0, F2: (fi33 * fi33)}
                                                _ = v107
                                                _ = v108
                                            }
                                        default:
                                            if len(v3) > 0 {
                                                v3[0] = 2147483647
                                            }
                                        }
                                        v109 := GenS6{F0: -80, F1: -924, F2: 939}
                                        v20 = v20
                                        v35 -= (fi33 % (((fi33 * (^(fi26)))) | 1))
                                        var v110 uint8 = 2
                                        _ = v109
                                        _ = v110
                                    }
                                case "":
                                    var v111 int64 = -767
                                    _ = v111
                                case "\n\t":
                                    fmt.Println(fi33, v95)
                                default:
                                    v20 = v20
                                }
                            }
                        }
                    } else {
                        var v112 bool = (!(v20 < v20))
                        _ = v112
                    }
                default:
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                }
            }
            _ = fi33
            _ = v35
            _ = v36
            _ = v95
        }
        var v113 uint = 0
        _ = fi26
        _ = v113
    }
    switch "\n\t" {
    case "foo bar":
        if (genFn0(119, (false && false)) < 1) {
            v114 := GenS15{F0: 0, F1: 127}
            _ = v114
        } else {
            v20 = v20
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
    case "":
        for fi115 := 0; fi115 < 20; fi115++ {
            var v116 string = ("\n\t" + v20)
            _ = fi115
            _ = v116
        }
    }
    switch ("foo bar" + ((((("" + ("foo bar" + v20)) + v20) + v20) + "\n\t") + "hello")) {
    case "\n\t":
        for fi117 := 0; fi117 < 4; fi117++ {
            if (65535 != 42) {
                v20 = v20
            } else {
                fmt.Println(v20, v20)
                v20 = v20
            }
            _ = fi117
        }
    }
    var v118 uint = 65535
    v20 = (v20 + "a")
    var v119 uint32 = 2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v118:", v118)
    fmt.Println("v119:", v119)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v118
    _ = v119
    _ = v13
    _ = v2
    _ = v20
    _ = v24
    _ = v25
    _ = v3
    _ = v4
}
