package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint16; F1 uint; F2 int32 }
type GenS5 struct { F0 int8; F1 uint16 }
type GenS6 struct { F0 uint; F1 int64; F2 int64 }
type GenS7 struct { F0 int8; F1 uint16 }
type GenS8 struct { F0 int32; F1 int8; F2 int32 }
type GenS9 struct { F0 int8; F1 string; F2 uint }
type GenS10 struct { F0 int64; F1 uint }
func genFn0(p0 int64) uint {
    fmt.Println(p0)
    p0 = p0
    for fi5 := 0; fi5 < 12; fi5++ {
        var v6 uint16 = 255
        p0 = -50
        _ = fi5
        _ = v6
    }
    var v7 uint = 65535
    v8 := []int{122, 610, 339, 177}
    _ = p0
    _ = v7
    _ = v8
    return 2
}
func genFn1(p0 int64) int {
    var v9 int8 = 1
    v10 := []int{42, -872, 128}
    _ = p0
    _ = v10
    _ = v9
    return 127
}

func main() {
    var v1 uint = 935
    var v2 int16 = 1000
    var v3 int16 = v2
    var v4 uint16 = 42
    for fi11 := 0; fi11 < 12; fi11++ {
        if ((v1 % ((genFn0(-9223372036854775808)) | 1)) != v1) {
            var v12 uint32 = 2147483648
            fmt.Println(v1, fi11, v12)
            fmt.Println(v3)
            _ = v12
        }
        _ = fi11
    }
    switch (("foo bar" + "") + "a") {
    case "a":
        v13 := []int{127, -1, 9223372036854775807, genFn1(0)}
        _ = v13
    case "":
        for fi14 := 0; fi14 < 2; fi14++ {
            if (!true) {
                v15 := []int{(^(fi14)), 0}
                _ = v15
            } else {
                if (!(0 == 2147483647)) {
                    fmt.Println(v4, v3)
                    v16 := []int{fi14}
                    v3 = v2
                    _ = v16
                }
            }
            _ = fi14
        }
    }
    if true {
        v4 -= 363
        v4 += (v4 / (((v4 ^ v4)) | 1))
    } else {
        if (!true) {
            v3 = (v3 % ((-724) | 1))
            v17 := GenS1{F0: 424}
            v2 = v2
            _ = v17
        } else {
            v18 := GenS2{F0: ("" + "a")}
            _ = v18
        }
    }
    for fi19 := 0; fi19 < 1; fi19++ {
        var v20 uint = v1
        _ = fi19
        _ = v20
    }
    if true {
        for fi21 := 0; fi21 < 14; fi21++ {
            var v22 bool = (!(-1 > 504))
            _ = fi21
            _ = v22
        }
    } else {
        v23 := []int{9223372036854775807, 1}
        var v24 int = genFn1(272)
        fmt.Println(v3, v1)
        _ = v23
        _ = v24
    }
    v2 = v3
    if (!false) {
        for fi25 := 0; fi25 < 15; fi25++ {
            if (true || (!false)) {
                for fi26 := 0; fi26 < 17; fi26++ {
                    var v27 uint16 = v4
                    var v28 uint64 = 2
                    var v29 int16 = v2
                    _ = fi26
                    _ = v27
                    _ = v28
                    _ = v29
                }
            } else {
                v30 := []int{}
                _ = v30
            }
            for fi31 := 0; fi31 < 5; fi31++ {
                v1 -= (v1 / ((v1) | 1))
                if ((!(1 < 2)) || true) {
                    var v32 float64 = float64(-2.25)
                    _ = v32
                } else {
                    fmt.Println(v4, v3)
                    switch fi31 {
                    case 4294967295:
                        fmt.Println(v3, v4, fi25)
                    case 525:
                        if (-176 < 0) {
                            fmt.Println(v1, fi25, fi25)
                            var v33 bool = false
                            _ = v33
                        } else {
                            switch genFn1(0) {
                            case 0:
                                fmt.Println(v2)
                                if (1 >= -432) {
                                    v34 := GenS1{F0: 18446744073709551615}
                                    if (2147483647 < -1) {
                                        var v35 int16 = v2
                                        v36 := []int{(-(fi31))}
                                        var v37 int64 = 0
                                        _ = v35
                                        _ = v36
                                        _ = v37
                                    } else {
                                        var v38 float32 = ((0.1 / -2.25) - 3.14159)
                                        _ = v38
                                    }
                                    _ = v34
                                }
                                v39 := GenS3{F0: 42}
                                v40 := GenS4{F0: v4, F1: v1, F2: -1}
                                v41 := GenS5{F0: 42, F1: (v4 - 661)}
                                if true {
                                    var v42 bool = (true && (-2.25 != -2.25))
                                    if v42 {
                                        switch (("hello" + ("\n\t" + "")) + "a") {
                                        case "hello":
                                            var v43 uint16 = 429
                                            _ = v43
                                        }
                                    } else {
                                        var v44 uint = genFn0(0)
                                        var v45 float32 = -2.25
                                        v46 := GenS2{F0: ("" + "foo bar")}
                                        switch ("\n\t" + "a") {
                                        case "foo bar":
                                            v47 := []int{int(v2), fi25}
                                            _ = v47
                                        case "\n\t":
                                            if v42 {
                                                v45 += ((v45 - v45) / (v45 - 3.14159))
                                                v48 := GenS5{F0: -84, F1: v4}
                                                if true {
                                                    var v49 int = 32767
                                                    _ = v49
                                                }
                                                _ = v48
                                            } else {
                                                v2 = v2
                                                var v50 int32 = 1
                                                if (v42 && ((2 != 1) && v42)) {
                                                    var v51 float32 = 3.14159
                                                    _ = v51
                                                } else {
                                                    var v52 float32 = ((3.14159 - (((v45 - float32(0.0)) * 3.14159) * 1e10)) / v45)
                                                    _ = v52
                                                }
                                                _ = v50
                                            }
                                        default:
                                            v53 := GenS1{F0: 0}
                                            _ = v53
                                        }
                                        _ = v44
                                        _ = v45
                                        _ = v46
                                    }
                                    switch fi25 {
                                    case 81:
                                        if ((642 < 2147483648) && v42) {
                                            fmt.Println(v2)
                                            var v54 int16 = (v3 * v3)
                                            v55 := []int{752, (fi25 << ((704) & 63)), int(v4)}
                                            _ = v54
                                            _ = v55
                                        } else {
                                            var v56 int32 = 1
                                            fmt.Println(v2)
                                            v4 -= v4
                                            _ = v56
                                        }
                                    default:
                                        v4 = 255
                                    }
                                    _ = v42
                                }
                                var v57 uint64 = 1
                                _ = v39
                                _ = v40
                                _ = v41
                                _ = v57
                            case 184:
                                fmt.Println(fi31, v3)
                            case 32767:
                                var v58 uint64 = 2
                                _ = v58
                            }
                            if true {
                                v3 -= 42
                                if false {
                                    if false {
                                        var v59 uint8 = 0
                                        var v60 uint = (v1 + v1)
                                        v61 := []int{fi31}
                                        _ = v59
                                        _ = v60
                                        _ = v61
                                    }
                                }
                                v3 -= (v2 % ((0) | 1))
                            } else {
                                var v62 string = ""
                                var v63 float32 = (1.5 * (3.14159 / 1.0))
                                if false {
                                    v63 -= ((float32(0.0) - (v63 * (float32(-2.25) / v63))) * v63)
                                    if (-1 == 127) {
                                        var v64 int = fi25
                                        _ = v64
                                    } else {
                                        v4 = v4
                                        switch (fi25 * fi31) {
                                        case 32767:
                                            v1 = v1
                                            switch "hello" {
                                            case "a":
                                                v65 := GenS6{F0: v1, F1: -472, F2: 440}
                                                _ = v65
                                            }
                                        default:
                                            v66 := GenS6{F0: v1, F1: 42, F2: 0}
                                            _ = v66
                                        }
                                    }
                                } else {
                                    v67 := GenS6{F0: 18446744073709551615, F1: -751, F2: 9223372036854775807}
                                    switch (v62 + v62) {
                                    case "\n\t":
                                        fmt.Println(fi31)
                                    case "a":
                                        v4 -= v4
                                    default:
                                        v68 := []int{fi31, genFn1(-222), (fi31 ^ -781)}
                                        _ = v68
                                    }
                                    v62 = "foo bar"
                                    _ = v67
                                }
                                _ = v62
                                _ = v63
                            }
                            v4 = (v4 << ((genFn0(42)) & 63))
                            if false {
                                fmt.Println(v4, v3, fi25)
                                v69 := GenS7{F0: -1, F1: 0}
                                _ = v69
                            }
                            v2 += (v2 & (-(v2)))
                        }
                    case 34:
                        fmt.Println(v1)
                    }
                    var v70 uint8 = 241
                    _ = v70
                }
                _ = fi31
            }
            for fi71 := 0; fi71 < 9; fi71++ {
                switch int(v3) {
                case 176:
                    v2 -= -400
                    if (2147483647 >= 698) {
                        v72 := GenS1{F0: 2}
                        if false {
                            switch fi71 {
                            case -9223372036854775808:
                                v4 -= (^(v4))
                            default:
                                v73 := []int{}
                                _ = v73
                            }
                        }
                        if true {
                            v3 -= (v3 | 32767)
                        }
                        _ = v72
                    }
                case 2147483647:
                    if true {
                        v3 += 0
                        var v74 int32 = 1
                        v75 := []int{genFn1(0)}
                        _ = v74
                        _ = v75
                    } else {
                        var v76 int16 = int16(v4)
                        _ = v76
                    }
                }
                if (128 == 0) {
                    v77 := []int{4294967295, (fi71 & fi71), -373, (fi25 << ((v1) & 63))}
                    switch ("a" + "\n\t") {
                    case "\n\t":
                        if (true || false) {
                            var v78 int64 = 1
                            _ = v78
                        } else {
                            v79 := GenS8{F0: 2147483647, F1: -128, F2: -443}
                            v2 -= (v2 - v2)
                            fmt.Println(v3, v2, fi71)
                            _ = v79
                        }
                    case "":
                        var v80 int64 = 1
                        _ = v80
                    case "foo bar":
                        v81 := GenS3{F0: 60}
                        fmt.Println(v4, v1)
                        _ = v81
                    }
                    v82 := GenS7{F0: int8(v4), F1: v4}
                    var v83 int64 = 9223372036854775807
                    _ = v77
                    _ = v82
                    _ = v83
                }
                _ = fi71
            }
            _ = fi25
        }
    } else {
        v84 := []int{-1, 703}
        switch 32767 {
        case -1:
            var v85 uint16 = uint16(v1)
            _ = v85
        default:
            fmt.Println(v4)
        }
        _ = v84
    }
    switch 1 {
    case -9223372036854775808:
        for fi86 := 0; fi86 < 4; fi86++ {
            var v87 int32 = 2147483647
            v88 := GenS6{F0: (v1 & genFn0(-1)), F1: -9223372036854775808, F2: 275}
            _ = fi86
            _ = v87
            _ = v88
        }
    case -2147483648:
        for fi89 := 0; fi89 < 15; fi89++ {
            for fi90 := 0; fi90 < 15; fi90++ {
                v91 := GenS5{F0: 100, F1: v4}
                switch ("a" + "\n\t") {
                case "foo bar":
                    var v92 float64 = 1.5
                    _ = v92
                default:
                    if (!(180 <= 193)) {
                        v3 += -655
                    } else {
                        v93 := GenS3{F0: 64}
                        v94 := GenS9{F0: 0, F1: "hello", F2: (v1 ^ genFn0(-185))}
                        _ = v93
                        _ = v94
                    }
                }
                var v95 uint16 = v4
                if (-128 != -1) {
                    v96 := []int{}
                    switch (^(fi90)) {
                    case -852:
                        if true {
                            switch "" {
                            case "\n\t":
                                var v97 uint = genFn0(1)
                                _ = v97
                            default:
                                v3 = (-(v2))
                            }
                        } else {
                            if len(v96) >= 2 {
                                v96 = v96[1:len(v96)]
                            }
                            var v98 float32 = (-2.25 - 3.14159)
                            _ = v98
                        }
                    case 32767:
                        var v99 int32 = 928
                        _ = v99
                    }
                    _ = v96
                } else {
                    var v100 bool = false
                    var v101 uint32 = 894
                    if v100 {
                        v102 := GenS10{F0: 158, F1: genFn0(9223372036854775807)}
                        fmt.Println(v100, v4, v4)
                        switch (^(fi89)) {
                        case 42:
                            var v103 int8 = -128
                            _ = v103
                        case 4294967295:
                            fmt.Println(fi89, v4)
                        default:
                            switch fi89 {
                            case 2147483647:
                                v95 = v4
                            case 548:
                                var v104 uint = v1
                                _ = v104
                            case 9223372036854775807:
                                var v105 int32 = 2147483647
                                v106 := GenS2{F0: "\n\t"}
                                v100 = v100
                                _ = v105
                                _ = v106
                            }
                        }
                        _ = v102
                    } else {
                        v107 := []int{}
                        if len(v107) > 0 {
                            v107[0] = (fi89 * genFn1(1))
                        }
                        v95 += v95
                        _ = v107
                    }
                    _ = v100
                    _ = v101
                }
                var v108 uint32 = 0
                _ = fi90
                _ = v108
                _ = v91
                _ = v95
            }
            _ = fi89
        }
    case -963:
        v1 = (v1 + v1)
    }
    fmt.Println(v2, v3)
    switch "\n\t" {
    case "\n\t":
        v3 -= (v2 & 92)
    default:
        fmt.Println(v4)
    }
    for fi109 := 0; fi109 < 9; fi109++ {
        v110 := GenS2{F0: (((((("foo bar" + "\n\t") + ("hello" + "")) + (("foo bar" + ("foo bar" + "a")) + ("hello" + "\n\t"))) + (("hello" + "\n\t") + "\n\t")) + "hello") + "\n\t")}
        _ = fi109
        _ = v110
    }
    if (false && false) {
        v1 = v1
        var v111 int16 = -1
        _ = v111
    }
    v112 := []int{genFn1(1), 228}
    var v113 uint16 = 78
    v1 = genFn0(517)
    if (0.1 < -2.25) {
        var v114 int8 = 0
        _ = v114
    } else {
        for fi115 := 0; fi115 < 9; fi115++ {
            v4 = v4
            if (!(2 != 2147483647)) {
                for fi116 := 0; fi116 < 3; fi116++ {
                    switch 593 {
                    case 4294967295:
                        var v117 int16 = (v3 >> (((v1 >> ((v1) & 63))) & 63))
                        if (v113 >= 854) {
                            var v118 int64 = 1
                            v117 = v117
                            _ = v118
                        } else {
                            v3 = (^(v3))
                        }
                        _ = v117
                    case -2147483648:
                        fmt.Println(v4, fi115, v113)
                    case 32768:
                        var v119 string = "a"
                        _ = v119
                    }
                    v120 := GenS7{F0: 0, F1: 1}
                    _ = fi116
                    _ = v120
                }
            }
            _ = fi115
        }
        for fi121 := 0; fi121 < 10; fi121++ {
            if true {
                var v122 int8 = 100
                _ = v122
            }
            _ = fi121
        }
        var v123 float32 = 3.14159
        _ = v123
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v112:", len(v112), v112)
    fmt.Println("v113:", v113)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v112
    _ = v113
    _ = v2
    _ = v3
    _ = v4
}
