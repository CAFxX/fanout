package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32; F2 int16 }
type GenS2 struct { F0 uint32; F1 uint8 }
type GenS3 struct { F0 int16; F1 bool }
type GenS4 struct { F0 int; F1 uint; F2 uint }
type GenS5 struct { F0 int32; F1 string }
type GenS6 struct { F0 uint32; F1 uint64 }
type GenS7 struct { F0 int16; F1 uint }
type GenS8 struct { F0 int16; F1 int8 }
type GenS9 struct { F0 uint16 }
type GenS10 struct { F0 int64; F1 uint16; F2 bool }
type GenS11 struct { F0 uint; F1 uint64; F2 int16 }
func genFn0(p0 int8, p1 uint32, p2 uint) uint16 {
    p1 += (p1 ^ (p1 << ((4294967295) & 63)))
    v6 := []int{-189, -333, -354, 32768}
    if (-1 >= -872) {
        fmt.Println(p2)
        p2 += (p2 & p2)
        fmt.Println(p2)
    } else {
        var v7 uint8 = 32
        v8 := GenS1{F0: 42, F1: 2147483647, F2: 1000}
        fmt.Println(p0, p1, p2)
        _ = v7
        _ = v8
    }
    var v9 int = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v9
    return 255
}
func genFn1(p0 uint64) uint64 {
    var v10 bool = (128 != 201)
    v11 := GenS2{F0: 957, F1: 1}
    v12 := GenS3{F0: 1, F1: ((false || ((p0 + p0) > (p0 >> ((1) & 63)))) && v10)}
    var v13 int64 = -9223372036854775808
    _ = p0
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    return 1
}

func main() {
    var v1 bool = true
    v2 := make(map[uint32]bool)
    v3 := []int{9223372036854775807}
    var v4 float32 = 3.14159
    var v5 uint16 = 0
    var v14 int16 = 32767
    for fi15 := 0; fi15 < 19; fi15++ {
        var v16 string = ((("" + "") + "a") + ((("" + "foo bar") + ("hello" + "\n\t")) + "\n\t"))
        for fi17 := 0; fi17 < 16; fi17++ {
            v16 = v16
            _ = fi17
        }
        v18 := GenS4{F0: fi15, F1: 1, F2: 409}
        _ = fi15
        _ = v16
        _ = v18
    }
    if v1 {
        v19 := GenS3{F0: 72, F1: false}
        _ = v19
    } else {
        for fi20 := 0; fi20 < 2; fi20++ {
            var v21 bool = true
            fmt.Println(v5, v4)
            _ = fi20
            _ = v21
        }
        if (128 != 127) {
            switch "foo bar" {
            case "":
                switch -446 {
                case 32768:
                    v1 = (!false)
                case 127:
                    if false {
                        v3 = append(v3, 4294967295)
                        var v22 float64 = 1.5
                        if v1 {
                            var v23 bool = false
                            var v24 int16 = -32768
                            var v25 uint64 = 2
                            _ = v23
                            _ = v24
                            _ = v25
                        }
                        _ = v22
                    } else {
                        if len(v3) > 0 {
                            v3[0] = 1
                        }
                    }
                }
            case "hello":
                v26 := GenS5{F0: -2147483648, F1: "\n\t"}
                _ = v26
            }
            fmt.Println(v14, v1)
            if (v1 && (v5 != v5)) {
                if (!(v1 && true)) {
                    if false {
                        for fi27 := 0; fi27 < 18; fi27++ {
                            var v28 int64 = -35
                            v1 = v1
                            switch 334 {
                            case 4294967295:
                                fmt.Println(v4, v5, v1)
                            case 0:
                                for fi29 := 0; fi29 < 16; fi29++ {
                                    delete(v2, 1)
                                    v30 := GenS5{F0: -2147483648, F1: ("hello" + "a")}
                                    delete(v2, 42)
                                    _ = fi29
                                    _ = v30
                                }
                            case 128:
                                v5 -= genFn0(48, 1, 18446744073709551615)
                            default:
                                for fi31 := 0; fi31 < 20; fi31++ {
                                    if v1 {
                                        if true {
                                            var v32 float32 = v4
                                            v32 += v4
                                            v33 := v2[42]
                                            _ = v32
                                            _ = v33
                                        }
                                    } else {
                                        if (((((((float32(float64(1e10)) / (0.0 - v4)) - (v4 * ((0.0 / 3.14159) * v4))) * ((v4 - ((v4 + (v4 * 1.5)) - (v4 / (1.5 + v4)))) / (v4 + 1e10))) * 0.0) + v4) * v4) <= 0.0) {
                                            switch ("hello" + "foo bar") {
                                            case "a":
                                                if ((0.1 + (float64(0.0) * -2.25)) > (1e10 / 1e10)) {
                                                    var v34 int64 = v28
                                                    if (false || v1) {
                                                        switch "a" {
                                                        case "\n\t":
                                                            switch "\n\t" {
                                                            case "a":
                                                                switch fi27 {
                                                                case -9223372036854775808:
                                                                    switch "" {
                                                                    case "":
                                                                        v4 += ((v4 + v4) + float32((float64(v4) - -2.25)))
                                                                    case "\n\t":
                                                                        v5 = v5
                                                                    default:
                                                                        v35 := GenS6{F0: 925, F1: 42}
                                                                        v4 -= -2.25
                                                                        switch ("foo bar" + "a") {
                                                                        case "hello":
                                                                            if (true && (v1 && (v14 <= (v14 / ((v14) | 1))))) {
                                                                                var v36 uint32 = 2147483648
                                                                                _ = v36
                                                                            } else {
                                                                                switch "" {
                                                                                case "\n\t":
                                                                                    v37 := v2[598]
                                                                                    _ = v37
                                                                                case "":
                                                                                    if true {
                                                                                        if ((v14 << ((573) & 63)) < v14) {
                                                                                            var v38 int16 = (v14 ^ v14)
                                                                                            var v39 uint8 = 56
                                                                                            var v40 uint8 = (v39 * (v39 & (v39 & v39)))
                                                                                            _ = v38
                                                                                            _ = v39
                                                                                            _ = v40
                                                                                        } else {
                                                                                            if len(v3) > 0 {
                                                                                                v3[0] = (fi31 & fi31)
                                                                                            }
                                                                                            v34 -= (v28 + v34)
                                                                                            fmt.Println(fi31, fi31)
                                                                                        }
                                                                                        var v41 uint32 = 2147483648
                                                                                        _ = v41
                                                                                    } else {
                                                                                        v42 := v2[288]
                                                                                        if len(v3) > 0 {
                                                                                            v3[0] = fi27
                                                                                        }
                                                                                        _ = v42
                                                                                    }
                                                                                case "foo bar":
                                                                                    switch fi27 {
                                                                                    case 827:
                                                                                        delete(v2, 2)
                                                                                    case -2147483648:
                                                                                        _ = len(v3)
                                                                                    case -1:
                                                                                        var v43 bool = v1
                                                                                        _ = v43
                                                                                    }
                                                                                }
                                                                                var v44 int64 = int64(v14)
                                                                                _ = v44
                                                                            }
                                                                        case "foo bar":
                                                                            v45 := GenS7{F0: -1, F1: 449}
                                                                            _ = v45
                                                                        case "":
                                                                            v1 = v1
                                                                        default:
                                                                            v46 := GenS8{F0: (v14 ^ (v14 & v14)), F1: 42}
                                                                            _ = v46
                                                                        }
                                                                        _ = v35
                                                                    }
                                                                case 0:
                                                                    var v47 string = "\n\t"
                                                                    _ = v47
                                                                default:
                                                                    v34 = v28
                                                                }
                                                            case "foo bar":
                                                                v14 += v14
                                                            default:
                                                                var v48 uint16 = v5
                                                                _ = v48
                                                            }
                                                        default:
                                                            fmt.Println(v5, fi31, fi27)
                                                        }
                                                        v4 = 3.14159
                                                    } else {
                                                        var v49 int = (fi31 % ((fi31) | 1))
                                                        v1 = v1
                                                        v1 = ((!false) || v1)
                                                        _ = v49
                                                    }
                                                    if (!((v1 || (!v1)) || (((!(false && v1)) && (100 != 100)) && ("foo bar" == "a")))) {
                                                        v3 = append(v3, (fi31 - fi31))
                                                    }
                                                    _ = len(v2)
                                                    v34 -= (v28 ^ (-(v28)))
                                                    var v50 uint16 = (v5 % (((v5 | v5)) | 1))
                                                    v1 = false
                                                    _ = v34
                                                    _ = v50
                                                } else {
                                                    v28 -= -610
                                                    fmt.Println(v14, fi27, v4)
                                                }
                                            default:
                                                fmt.Println(v4, v14, fi27)
                                            }
                                        } else {
                                            v14 = (v14 | v14)
                                            fmt.Println(v28)
                                            fmt.Println(fi31, v5)
                                        }
                                        v5 = v5
                                        if (-662 < fi27) {
                                            var v51 int16 = (v14 | 1000)
                                            _ = v51
                                        }
                                    }
                                    var v52 uint64 = 42
                                    _ = fi31
                                    _ = v52
                                }
                            }
                            _ = fi27
                            _ = v28
                        }
                    } else {
                        var v53 uint = 579
                        v14 = v14
                        _ = v53
                    }
                    if (18446744073709551615 == 726) {
                        var v54 float32 = ((((0.1 * 1.5) * float32(1.5)) + v4) * v4)
                        _ = v54
                    } else {
                        var v55 int8 = 57
                        v56 := GenS8{F0: (v14 & (v14 * v14)), F1: 100}
                        _ = v55
                        _ = v56
                    }
                } else {
                    if (!(955 > 336)) {
                        v57 := GenS9{F0: v5}
                        v14 += 1
                        for fi58 := 0; fi58 < 9; fi58++ {
                            for fi59 := 0; fi59 < 10; fi59++ {
                                v4 += v4
                                if (true && (v14 < v14)) {
                                    v5 = (v5 + v5)
                                    var v60 uint16 = 2
                                    if v1 {
                                        v61 := GenS6{F0: 889, F1: 4294967295}
                                        var v62 uint16 = v60
                                        _ = v61
                                        _ = v62
                                    } else {
                                        var v63 uint16 = v60
                                        _ = v63
                                    }
                                    _ = v60
                                } else {
                                    if (fi59 < fi58) {
                                        v4 -= ((v4 - 0.1) + (v4 / float32(-2.25)))
                                        fmt.Println(fi58, fi58, v4)
                                        if v1 {
                                            v64 := GenS10{F0: 42, F1: (v5 + v5), F2: v1}
                                            if (!false) {
                                                var v65 uint64 = 674
                                                var v66 string = (("\n\t" + "a") + ("foo bar" + "hello"))
                                                _ = v65
                                                _ = v66
                                            } else {
                                                var v67 uint = 484
                                                _ = v67
                                            }
                                            v68 := v2[uint32(v5)]
                                            v69 := v2[747]
                                            if v69 {
                                                v68 = v1
                                                if v1 {
                                                    v14 = v14
                                                }
                                            }
                                            _ = v64
                                            _ = v68
                                            _ = v69
                                        }
                                    }
                                }
                                v3 = append(v3, (-(fi59)))
                                _ = fi59
                            }
                            _ = fi58
                        }
                        _ = v57
                    } else {
                        if true {
                            v70 := v2[2147483647]
                            switch "\n\t" {
                            case "":
                                switch 32767 {
                                case 0:
                                    fmt.Println(v4)
                                case 32768:
                                    _ = len(v2)
                                case 32767:
                                    v1 = (true && (!(((0.0 + (0.0 * (1e10 + 0.0))) - 3.14159) != -2.25)))
                                default:
                                    fmt.Println(v5, v4)
                                }
                            case "a":
                                v5 -= v5
                            default:
                                for fi71 := 0; fi71 < 9; fi71++ {
                                    var v72 int32 = 1
                                    switch 42 {
                                    case -822:
                                        v3 = append(v3, fi71)
                                    default:
                                        for fi73 := 0; fi73 < 20; fi73++ {
                                            v1 = (!(!v70))
                                            _ = fi73
                                        }
                                    }
                                    _ = fi71
                                    _ = v72
                                }
                            }
                            for fi74 := 0; fi74 < 7; fi74++ {
                                if v1 {
                                    var v75 uint = 596
                                    v1 = (!v70)
                                    _ = v75
                                }
                                _ = fi74
                            }
                            _ = v70
                        } else {
                            switch 9223372036854775807 {
                            case 9223372036854775807:
                                if true {
                                    if false {
                                        switch -857 {
                                        case 724:
                                            var v76 string = "foo bar"
                                            _ = v76
                                        case 32767:
                                            var v77 float32 = v4
                                            _ = v77
                                        }
                                    }
                                    v1 = v1
                                } else {
                                    v78 := GenS6{F0: 0, F1: 1}
                                    v4 += v4
                                    switch ((("a" + (("foo bar" + "\n\t") + "\n\t")) + "a") + "hello") {
                                    case "\n\t":
                                        fmt.Println(v14, v14, v1)
                                    case "hello":
                                        for fi79 := 0; fi79 < 10; fi79++ {
                                            switch (fi79 * fi79) {
                                            case 32767:
                                                if (255 >= (v5 << ((18446744073709551615) & 63))) {
                                                    for fi80 := 0; fi80 < 17; fi80++ {
                                                        v1 = (18446744073709551615 >= 4294967295)
                                                        _ = fi80
                                                    }
                                                    v81 := GenS3{F0: v14, F1: v1}
                                                    v1 = v1
                                                    _ = v81
                                                }
                                            case 64:
                                                var v82 bool = (2 >= 255)
                                                _ = v82
                                            default:
                                                var v83 int = fi79
                                                _ = v83
                                            }
                                            switch (fi79 ^ (^(fi79))) {
                                            case 1:
                                                if (!(((float64(float32((0.1 + 0.0))) + float64((v4 - (-2.25 + v4)))) / (1.5 / 1.0)) == (0.0 + (0.1 * 1e10)))) {
                                                    if (!v1) {
                                                        var v84 bool = ((0.1 - 0.0) <= 0.0)
                                                        switch "\n\t" {
                                                        case "foo bar":
                                                            for fi85 := 0; fi85 < 17; fi85++ {
                                                                var v86 int32 = 946
                                                                _ = fi85
                                                                _ = v86
                                                            }
                                                        case "hello":
                                                            v14 = v14
                                                        default:
                                                            delete(v2, 151)
                                                        }
                                                        v2[371] = v1
                                                        _ = v84
                                                    }
                                                }
                                            }
                                            if v1 {
                                                for fi87 := 0; fi87 < 1; fi87++ {
                                                    if (0 < -1) {
                                                        v5 += (v5 % (((v5 - v5)) | 1))
                                                        if (!(v1 || ((145 == 884) && v1))) {
                                                            switch -1 {
                                                            case 32767:
                                                                if v1 {
                                                                    var v88 uint16 = 65535
                                                                    v4 = v4
                                                                    _ = v88
                                                                }
                                                            }
                                                            fmt.Println(v5, fi87)
                                                            v89 := GenS8{F0: 1000, F1: 0}
                                                            var v90 string = ("foo bar" + "")
                                                            _ = v89
                                                            _ = v90
                                                        }
                                                        v91 := GenS9{F0: (v5 - (v5 & (v5 + v5)))}
                                                        _ = v91
                                                    } else {
                                                        v4 = v4
                                                        _ = len(v3)
                                                        if true {
                                                            v3 = append(v3, (fi87 % ((fi79) | 1)))
                                                            if len(v3) > 0 {
                                                                v3[0] = (fi87 & fi79)
                                                            }
                                                            var v92 uint8 = 30
                                                            _ = v92
                                                        }
                                                    }
                                                    var v93 uint64 = 78
                                                    var v94 int32 = 42
                                                    _ = fi87
                                                    _ = v93
                                                    _ = v94
                                                }
                                                for fi95 := 0; fi95 < 2; fi95++ {
                                                    v96 := GenS2{F0: 2147483648, F1: 151}
                                                    var v97 float64 = float64(v4)
                                                    _ = fi95
                                                    _ = v96
                                                    _ = v97
                                                }
                                            } else {
                                                switch "" {
                                                case "\n\t":
                                                    fmt.Println(v5, v4, fi79)
                                                }
                                            }
                                            _ = fi79
                                        }
                                    case "foo bar":
                                        var v98 uint64 = 267
                                        _ = v98
                                    }
                                    _ = v78
                                }
                            }
                            switch (("a" + "hello") + "a") {
                            case "foo bar":
                                v1 = (3.14159 <= (0.1 - (3.14159 / 1e10)))
                            case "hello":
                                var v99 int8 = 0
                                _ = v99
                            case "":
                                v1 = v1
                            default:
                                v4 -= v4
                            }
                            if (1 == 1) {
                                _ = len(v3)
                                v4 = float32((-2.25 / 1.0))
                            } else {
                                v1 = true
                            }
                        }
                        _ = len(v3)
                        v5 = v5
                    }
                    for fi100 := 0; fi100 < 6; fi100++ {
                        var v101 int32 = 878
                        for fi102 := 0; fi102 < 17; fi102++ {
                            var v103 int16 = v14
                            _ = fi102
                            _ = v103
                        }
                        _ = fi100
                        _ = v101
                    }
                }
                var v104 float32 = ((v4 / (v4 - (float32(1.5) / ((3.14159 + float32(1.5)) - 0.1)))) - v4)
                _ = v104
            } else {
                for fi105 := 0; fi105 < 3; fi105++ {
                    switch "a" {
                    case "\n\t":
                        v5 = v5
                    case "foo bar":
                        if v1 {
                            var v106 int8 = -16
                            v107 := v2[4294967295]
                            _ = v106
                            _ = v107
                        }
                    case "":
                        delete(v2, 0)
                    default:
                        var v108 uint = 42
                        _ = v108
                    }
                    fmt.Println(v5)
                    _ = fi105
                }
                for fi109 := 0; fi109 < 18; fi109++ {
                    v14 -= v14
                    _ = fi109
                }
            }
        } else {
            v5 = (v5 & 42)
            v3 = append(v3, -195)
            v5 = v5
        }
    }
    if false {
        v4 -= v4
        var v110 float32 = v4
        for fi111 := 0; fi111 < 14; fi111++ {
            var v112 uint32 = 442
            _ = fi111
            _ = v112
        }
        _ = v110
    }
    if v1 {
        v113 := GenS2{F0: 4294967295, F1: 42}
        _ = v113
    } else {
        v114 := GenS1{F0: 32767, F1: -870, F2: -1}
        for fi115 := 0; fi115 < 13; fi115++ {
            var v116 uint64 = 18446744073709551615
            fmt.Println(v4, v4)
            v14 -= (-(v14))
            _ = fi115
            _ = v116
        }
        v117 := GenS6{F0: 0, F1: genFn1(18446744073709551615)}
        _ = v114
        _ = v117
    }
    v118 := GenS8{F0: (v14 % ((v14) | 1)), F1: 42}
    var v119 int8 = 127
    var v120 float64 = float64(v4)
    v121 := GenS7{F0: v14, F1: 681}
    var v122 int32 = 42
    _ = len(v3)
    for fi123 := 0; fi123 < 20; fi123++ {
        if v1 {
            v1 = (673 < (^(v5)))
            v124 := GenS11{F0: 1, F1: 2, F2: int16(v119)}
            _ = v124
        }
        _ = fi123
    }
    fmt.Println("v1:", v1)
    fmt.Println("v119:", v119)
    fmt.Println("v120:", v120)
    fmt.Println("v122:", v122)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v118
    _ = v119
    _ = v120
    _ = v121
    _ = v122
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
