package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint; F2 int32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint64; F1 int64 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int64; F1 int32 }
type GenS6 struct { F0 string; F1 string; F2 bool }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 uint16 }
type GenS9 struct { F0 uint }
type GenS10 struct { F0 uint }
func genFn0(p0 uint16, p1 bool) int64 {
    fmt.Println(p0, p0)
    var v5 uint8 = uint8(p0)
    _ = p0
    _ = p1
    _ = v5
    return 159
}
func genFn1(p0 uint64) int {
    if (100 != 1) {
        v6 := GenS1{F0: 1, F1: 25, F2: -117}
        fmt.Println(p0)
        var v7 float32 = 0.0
        _ = v6
        _ = v7
    }
    p0 -= p0
    for fi8 := 0; fi8 < 19; fi8++ {
        var v9 int8 = -128
        v10 := []int{fi8, (fi8 << ((888) & 63)), fi8}
        for fi11 := 0; fi11 < 6; fi11++ {
            switch 0 {
            case 2147483647:
                var v12 float64 = (((0.0 * (0.0 * 1.5)) / (1.5 + (1e10 + 0.1))) / 1.0)
                _ = v12
            case 42:
                if len(v10) >= 2 {
                    v10 = v10[1:len(v10)]
                }
            default:
                var v13 int64 = -9223372036854775808
                var v14 int = (fi8 & (fi11 & fi8))
                var v15 int32 = 42
                if false {
                    switch int(v9) {
                    case 32767:
                        v16 := GenS2{F0: (fi11 | (^(fi11)))}
                        _ = v16
                    case 707:
                        _ = len(v10)
                    case 4294967295:
                        v17 := GenS1{F0: v9, F1: 161, F2: (-(v15))}
                        _ = v17
                    }
                    if len(v10) >= 2 {
                        v10 = v10[1:len(v10)]
                    }
                }
                _ = v13
                _ = v14
                _ = v15
            }
            var v18 float64 = 0.1
            _ = len(v10)
            _ = fi11
            _ = v18
        }
        _ = fi8
        _ = v10
        _ = v9
    }
    var v19 int64 = genFn0(65535, true)
    _ = p0
    _ = v19
    return 9223372036854775807
}

func main() {
    var v1 bool = (2 == 42)
    var v2 float32 = -2.25
    v3 := []uint32{}
    var v4 int32 = -1
    v1 = v1
    v4 = (v4 * (^(v4)))
    v20 := GenS3{F0: 1, F1: 857}
    switch "" {
    case "a":
        switch ((("foo bar" + ((("" + "foo bar") + "hello") + "\n\t")) + "a") + "\n\t") {
        case "a":
            switch "\n\t" {
            case "hello":
                fmt.Println(v4, v1, v2)
            case "a":
                fmt.Println(v2, v1)
            }
        case "foo bar":
            var v21 uint64 = 1
            _ = v21
        default:
            v1 = v1
        }
    case "foo bar":
        for fi22 := 0; fi22 < 2; fi22++ {
            for fi23 := 0; fi23 < 12; fi23++ {
                if v1 {
                    var v24 int8 = 127
                    _ = v24
                } else {
                    fmt.Println(v2)
                }
                v4 = (v4 + -1)
                var v25 int8 = 1
                _ = fi23
                _ = v25
            }
            var v26 int8 = 0
            fmt.Println(v4, fi22)
            _ = fi22
            _ = v26
        }
    }
    fmt.Println(v1, v2, v2)
    if (-525 == 42) {
        var v27 uint64 = 0
        if (v2 <= float32(0.0)) {
            if (((v2 + 1.5) / v2) == v2) {
                v28 := GenS2{F0: 4294967295}
                fmt.Println(v2, v1, v1)
                var v29 uint = 952
                _ = v28
                _ = v29
            } else {
                fmt.Println(v2, v1)
                if (!v1) {
                    v30 := GenS4{F0: v27}
                    if (((float64(3.14159) * ((0.0 * ((float64(0.0) + 3.14159) / 1e10)) / (3.14159 - 0.0))) / float64((float32(float64(0.1)) + v2))) < (3.14159 * 0.0)) {
                        switch 32767 {
                        case 32768:
                            for fi31 := 0; fi31 < 20; fi31++ {
                                var v32 uint16 = 37
                                fmt.Println(v4)
                                _ = fi31
                                _ = v32
                            }
                        case -2147483648:
                            if v1 {
                                var v33 int32 = (-(v4))
                                v27 = (v27 ^ (v27 % (((v27 & (v27 >> ((255) & 63)))) | 1)))
                                v33 = (v33 % (((^(v4))) | 1))
                                _ = v33
                            } else {
                                for fi34 := 0; fi34 < 15; fi34++ {
                                    v27 += (v27 % (((^(v27))) | 1))
                                    switch ("hello" + "a") {
                                    case "":
                                        var v35 uint16 = 65535
                                        _ = v35
                                    default:
                                        for fi36 := 0; fi36 < 4; fi36++ {
                                            switch "" {
                                            case "\n\t":
                                                v37 := GenS1{F0: -128, F1: 889, F2: 1}
                                                _ = v37
                                            case "hello":
                                                if len(v3) > 0 {
                                                    v3[0] = 65535
                                                }
                                            }
                                            var v38 uint64 = (v27 + v27)
                                            v39 := GenS5{F0: genFn0(642, ((-(fi34)) != -2147483648)), F1: v4}
                                            _ = fi36
                                            _ = v38
                                            _ = v39
                                        }
                                    }
                                    _ = fi34
                                }
                                if (0.1 < float32(-2.25)) {
                                    if (true && (2147483647 > 2147483647)) {
                                        if v1 {
                                            _ = len(v3)
                                        } else {
                                            var v40 int64 = -594
                                            _ = v40
                                        }
                                        var v41 float64 = 1e10
                                        _ = v41
                                    } else {
                                        var v42 int8 = 42
                                        _ = v42
                                    }
                                    for fi43 := 0; fi43 < 13; fi43++ {
                                        v44 := GenS6{F0: ("" + "hello"), F1: (("a" + "foo bar") + ("" + "\n\t")), F2: ((float64((v2 + v2)) - 3.14159) == 0.1)}
                                        _ = fi43
                                        _ = v44
                                    }
                                } else {
                                    v4 += (v4 / ((42) | 1))
                                    for fi45 := 0; fi45 < 1; fi45++ {
                                        v46 := GenS1{F0: 127, F1: 0, F2: v4}
                                        if v1 {
                                            v47 := GenS3{F0: v27, F1: -9223372036854775808}
                                            _ = len(v3)
                                            _ = v47
                                        } else {
                                            v2 = v2
                                        }
                                        _ = fi45
                                        _ = v46
                                    }
                                }
                                for fi48 := 0; fi48 < 6; fi48++ {
                                    for fi49 := 0; fi49 < 18; fi49++ {
                                        if len(v3) > 0 {
                                            v3[0] = 2
                                        }
                                        _ = fi49
                                    }
                                    switch 501 {
                                    case 32767:
                                        if ((!true) && (-128 >= -128)) {
                                            if ((!v1) || v1) {
                                                var v50 uint8 = 36
                                                _ = v50
                                            } else {
                                                switch (fi48 >> ((255) & 63)) {
                                                case 9223372036854775807:
                                                    var v51 uint16 = uint16(v27)
                                                    _ = v51
                                                default:
                                                    if len(v3) > 0 {
                                                        v3[0] = 481
                                                    }
                                                }
                                                v3 = append(v3, 858)
                                                v2 -= ((-2.25 * (float32((-2.25 - 0.0)) + v2)) * v2)
                                            }
                                            v27 += v27
                                        }
                                    }
                                    if (0.0 >= 1.5) {
                                        if (false && v1) {
                                            switch fi48 {
                                            case -1:
                                                for fi52 := 0; fi52 < 4; fi52++ {
                                                    if (1 < 2) {
                                                        var v53 int = fi52
                                                        if v1 {
                                                            v1 = (!v1)
                                                            var v54 float32 = v2
                                                            v53 = fi48
                                                            _ = v54
                                                        }
                                                        if len(v3) >= 2 {
                                                            v3 = v3[1:len(v3)]
                                                        }
                                                        var v55 int64 = 42
                                                        v55 = (v55 * v55)
                                                        _ = v53
                                                        _ = v55
                                                    } else {
                                                        v3 = append(v3, 2147483647)
                                                    }
                                                    if len(v3) >= 2 {
                                                        v3 = v3[1:len(v3)]
                                                    }
                                                    if (0 <= 236) {
                                                        var v56 int8 = -128
                                                        var v57 float32 = ((v2 * float32((float64(1e10) / -2.25))) * (v2 * v2))
                                                        v58 := GenS2{F0: -1}
                                                        _ = v56
                                                        _ = v57
                                                        _ = v58
                                                    } else {
                                                        var v59 bool = true
                                                        _ = v59
                                                    }
                                                    _ = fi52
                                                }
                                            }
                                            v4 -= (v4 | (v4 | (v4 * (v4 % ((-441) | 1)))))
                                        }
                                        v1 = (v2 <= 0.0)
                                    } else {
                                        var v60 int32 = (v4 | -2147483648)
                                        _ = v60
                                    }
                                    _ = fi48
                                }
                            }
                        default:
                            for fi61 := 0; fi61 < 19; fi61++ {
                                var v62 int64 = genFn0(32767, false)
                                for fi63 := 0; fi63 < 6; fi63++ {
                                    v64 := GenS4{F0: 2}
                                    if (128 < 138) {
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                        if false {
                                            fmt.Println(fi61, v27)
                                            v65 := GenS2{F0: -724}
                                            v4 = (v4 - (v4 >> ((590) & 63)))
                                            _ = v65
                                        } else {
                                            _ = len(v3)
                                            v66 := GenS6{F0: "hello", F1: "\n\t", F2: (892 <= 191)}
                                            _ = v66
                                        }
                                        v67 := GenS3{F0: (v27 / ((4294967295) | 1)), F1: (v62 % (((v62 * 42)) | 1))}
                                        var v68 uint8 = 6
                                        _ = v67
                                        _ = v68
                                    } else {
                                        if (v1 && false) {
                                            fmt.Println(v4, v2)
                                            switch "foo bar" {
                                            case "\n\t":
                                                var v69 float32 = 3.14159
                                                _ = v69
                                            case "hello":
                                                var v70 uint = 65535
                                                _ = v70
                                            default:
                                                if (602 >= 66) {
                                                    v4 -= int32(v62)
                                                } else {
                                                    if len(v3) >= 2 {
                                                        v3 = v3[1:len(v3)]
                                                    }
                                                }
                                            }
                                        }
                                    }
                                    _ = fi63
                                    _ = v64
                                }
                                switch 4294967295 {
                                case -601:
                                    switch "hello" {
                                    case "foo bar":
                                        if len(v3) > 0 {
                                            v3[0] = 2147483648
                                        }
                                    default:
                                        var v71 float64 = -2.25
                                        _ = v71
                                    }
                                case -2147483648:
                                    v72 := GenS7{F0: ((42 != 859) && v1)}
                                    _ = v72
                                case 4294967295:
                                    v27 += v27
                                }
                                _ = fi61
                                _ = v62
                            }
                        }
                    }
                    fmt.Println(v4, v4)
                    _ = v30
                }
            }
        } else {
            v4 = v4
            var v73 uint64 = v27
            _ = v73
        }
        _ = v27
    } else {
        if true {
            if (-132 < 32767) {
                switch "\n\t" {
                case "":
                    v4 = (v4 << ((621) & 63))
                case "\n\t":
                    v1 = true
                default:
                    v1 = v1
                }
            }
        } else {
            for fi74 := 0; fi74 < 9; fi74++ {
                switch "hello" {
                case "hello":
                    var v75 int8 = int8(fi74)
                    _ = v75
                case "":
                    fmt.Println(v4)
                case "\n\t":
                    v76 := GenS8{F0: 0}
                    _ = v76
                }
                if (true && (2 < 955)) {
                    switch (-(fi74)) {
                    case 959:
                        if v1 {
                            var v77 int8 = 52
                            if (v1 && (43 > 4294967295)) {
                                v2 -= (float32(float64(float32((1e10 / float64((((v2 - (((v2 + (float32(0.1) / ((0.1 - -2.25) / (float32(1e10) + v2)))) * float32(1.5)) - v2)) * (v2 + v2)) * 1.5)))))) + v2)
                                v77 += (v77 % (((v77 % ((-128) | 1))) | 1))
                                for fi78 := 0; fi78 < 18; fi78++ {
                                    var v79 int = fi78
                                    _ = fi78
                                    _ = v79
                                }
                            } else {
                                _ = len(v3)
                            }
                            _ = v77
                        }
                    case -2147483648:
                        v2 -= v2
                    default:
                        var v80 float32 = v2
                        _ = v80
                    }
                    var v81 uint = 65535
                    _ = v81
                } else {
                    switch -1 {
                    case 149:
                        for fi82 := 0; fi82 < 16; fi82++ {
                            _ = len(v3)
                            switch "hello" {
                            case "":
                                var v83 float64 = (1.5 + float64(float32((3.14159 - float64(v2)))))
                                _ = v83
                            case "foo bar":
                                var v84 int32 = 0
                                _ = v84
                            case "a":
                                v1 = (fi82 == fi74)
                            default:
                                v4 = 982
                            }
                            _ = fi82
                        }
                    case 4294967295:
                        if (4294967295 == 0) {
                            v85 := GenS9{F0: 255}
                            var v86 uint64 = 946
                            fmt.Println(v1)
                            _ = v85
                            _ = v86
                        }
                    }
                }
                _ = fi74
            }
        }
    }
    var v87 int16 = 123
    for fi88 := 0; fi88 < 11; fi88++ {
        v3 = append(v3, 2147483647)
        _ = fi88
    }
    if false {
        fmt.Println(v1)
        var v89 uint16 = 967
        if (!(-1 != 0)) {
            for fi90 := 0; fi90 < 18; fi90++ {
                var v91 uint64 = 707
                _ = fi90
                _ = v91
            }
        } else {
            for fi92 := 0; fi92 < 5; fi92++ {
                switch fi92 {
                case 128:
                    switch fi92 {
                    case 0:
                        if (0 > 255) {
                            if ((true || true) && (42 == 6)) {
                                v93 := GenS8{F0: v89}
                                v1 = (9223372036854775807 != 9223372036854775807)
                                v94 := GenS9{F0: 65535}
                                _ = v93
                                _ = v94
                            }
                        } else {
                            for fi95 := 0; fi95 < 8; fi95++ {
                                var v96 uint16 = v89
                                var v97 uint64 = 945
                                var v98 int = 171
                                _ = fi95
                                _ = v96
                                _ = v97
                                _ = v98
                            }
                            fmt.Println(v87, v87, v89)
                        }
                    default:
                        v1 = false
                    }
                case 0:
                    if len(v3) > 0 {
                        v3[0] = 1
                    }
                case 32768:
                    v1 = v1
                }
                _ = fi92
            }
            switch ("hello" + "") {
            case "foo bar":
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            case "a":
                var v99 int8 = -1
                _ = v99
            default:
                for fi100 := 0; fi100 < 6; fi100++ {
                    for fi101 := 0; fi101 < 9; fi101++ {
                        if (v1 && (0.1 >= 1e10)) {
                            fmt.Println(v2, fi100, fi100)
                            var v102 string = ""
                            _ = v102
                        }
                        switch "\n\t" {
                        case "a":
                            v3 = append(v3, 1)
                        case "":
                            var v103 uint16 = v89
                            _ = v103
                        }
                        v104 := GenS3{F0: 0, F1: 42}
                        var v105 int = 724
                        fmt.Println(v4)
                        _ = fi101
                        _ = v104
                        _ = v105
                    }
                    var v106 int32 = (v4 / (((v4 - -942)) | 1))
                    switch "a" {
                    case "foo bar":
                        for fi107 := 0; fi107 < 7; fi107++ {
                            v108 := GenS1{F0: 127, F1: 650, F2: -967}
                            v109 := GenS4{F0: 42}
                            v2 += v2
                            _ = fi107
                            _ = v108
                            _ = v109
                        }
                    case "":
                        if (!v1) {
                            v4 -= v106
                        } else {
                            var v110 uint8 = 42
                            _ = v110
                        }
                    case "a":
                        v3 = append(v3, 42)
                    default:
                        if (!false) {
                            v89 = (v89 & 206)
                            v111 := GenS7{F0: true}
                            for fi112 := 0; fi112 < 9; fi112++ {
                                fmt.Println(v89, v87, v4)
                                fmt.Println(fi112, v89)
                                _ = fi112
                            }
                            _ = v111
                        }
                    }
                    _ = fi100
                    _ = v106
                }
            }
        }
        _ = v89
    }
    _ = len(v3)
    if len(v3) > 0 {
        v3[0] = 472
    }
    for fi113 := 0; fi113 < 6; fi113++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        if (1 > 2) {
            switch "foo bar" {
            case "":
                v114 := GenS10{F0: 42}
                _ = v114
            default:
                if true {
                    v2 = (3.14159 + -2.25)
                    fmt.Println(v2, v1)
                } else {
                    v87 -= (v87 * v87)
                    var v115 uint8 = 255
                    _ = v115
                }
            }
            fmt.Println(v4)
            if len(v3) > 0 {
                v3[0] = 136
            }
        }
        fmt.Println(v87, v1, v1)
        _ = fi113
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v87:", v87)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v87
}
