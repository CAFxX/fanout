package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64 }
type GenS2 struct { F0 int64; F1 int64; F2 uint64 }
type GenS3 struct { F0 int32; F1 bool; F2 uint }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int; F1 bool }
type GenS6 struct { F0 uint; F1 uint; F2 uint64 }
type GenS7 struct { F0 int; F1 int64 }
type GenS8 struct { F0 int64 }
type GenS9 struct { F0 int }
type GenS10 struct { F0 int }
type GenS11 struct { F0 int8 }
type GenS12 struct { F0 uint8 }
func genFn0(p0 uint) int8 {
    var v7 int = 4294967295
    var v8 int = -2147483648
    var v9 string = ((("hello" + "foo bar") + "hello") + (("\n\t" + "") + "foo bar"))
    v10 := GenS1{F0: (42 < 4294967295), F1: 1}
    _ = p0
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return 1
}
func genFn1(p0 uint16) int32 {
    if true {
        if (!((true && false) && (0 > 2147483648))) {
            v11 := GenS1{F0: true, F1: 0}
            var v12 bool = (!false)
            _ = v11
            _ = v12
        }
        p0 += (p0 + 32767)
        if false {
            v13 := GenS2{F0: 0, F1: 9223372036854775807, F2: 994}
            _ = v13
        }
    } else {
        fmt.Println(p0)
        v14 := GenS2{F0: 9223372036854775807, F1: 0, F2: 1}
        v15 := []int{1, 32767}
        _ = v14
        _ = v15
    }
    if (-794 >= -1) {
        for fi16 := 0; fi16 < 11; fi16++ {
            if false {
                for fi17 := 0; fi17 < 8; fi17++ {
                    switch (^(fi17)) {
                    case -2147483648:
                        if false {
                            var v18 int32 = -76
                            _ = v18
                        } else {
                            var v19 uint64 = 4294967295
                            _ = v19
                        }
                    case 2147483647:
                        v20 := []int{(fi17 - fi17), fi17, (fi17 & (fi16 >> ((65535) & 63))), (fi17 * 1)}
                        _ = v20
                    default:
                        p0 = (p0 - (p0 << ((992) & 63)))
                    }
                    var v21 int16 = 1
                    _ = fi17
                    _ = v21
                }
                switch fi16 {
                case 2147483647:
                    var v22 bool = (2147483647 == -1)
                    _ = v22
                case -961:
                    fmt.Println(p0)
                case 1:
                    fmt.Println(fi16, fi16)
                }
            } else {
                if false {
                    switch fi16 {
                    case 864:
                        if (4294967295 < 270) {
                            p0 -= 42
                        } else {
                            var v23 float32 = 1.5
                            if true {
                                p0 -= p0
                            } else {
                                v24 := []int{fi16, (fi16 * fi16)}
                                if (147 < 2) {
                                    _ = len(v24)
                                } else {
                                    v24 = append(v24, 4294967295)
                                }
                                _ = v24
                            }
                            v25 := []int{(fi16 % ((fi16) | 1)), (fi16 | fi16)}
                            _ = v23
                            _ = v25
                        }
                    case 4294967295:
                        p0 = (p0 | p0)
                    default:
                        if false {
                            if (679 <= 0) {
                                for fi26 := 0; fi26 < 11; fi26++ {
                                    var v27 string = ((("foo bar" + ("" + "a")) + ("foo bar" + "foo bar")) + ("hello" + "foo bar"))
                                    v27 = v27
                                    _ = fi26
                                    _ = v27
                                }
                            } else {
                                fmt.Println(fi16, fi16)
                            }
                        }
                    }
                } else {
                    v28 := GenS2{F0: -1, F1: 42, F2: 436}
                    p0 -= (p0 % ((p0) | 1))
                    p0 -= p0
                    _ = v28
                }
                var v29 int32 = 2147483647
                if false {
                    switch (-(fi16)) {
                    case 42:
                        v30 := GenS2{F0: 0, F1: -9223372036854775808, F2: 18446744073709551615}
                        _ = v30
                    case -486:
                        v31 := []int{}
                        _ = v31
                    default:
                        fmt.Println(v29, fi16, p0)
                    }
                    var v32 int32 = (v29 >> ((42) & 63))
                    _ = v32
                }
                _ = v29
            }
            p0 -= (p0 << ((2) & 63))
            v33 := GenS2{F0: 0, F1: 1, F2: 4294967295}
            _ = fi16
            _ = v33
        }
        var v34 bool = (false && false)
        v35 := []int{-164, 9223372036854775807, 1}
        _ = v34
        _ = v35
    }
    for fi36 := 0; fi36 < 20; fi36++ {
        for fi37 := 0; fi37 < 7; fi37++ {
            if false {
                if (!false) {
                    v38 := []int{fi36}
                    var v39 float64 = 1.5
                    _ = v38
                    _ = v39
                } else {
                    v40 := GenS2{F0: 9223372036854775807, F1: 42, F2: 4294967295}
                    var v41 int16 = -588
                    _ = v40
                    _ = v41
                }
            }
            if true {
                p0 += 255
            } else {
                p0 += (^(p0))
                p0 = p0
            }
            _ = fi37
        }
        var v42 int8 = 1
        p0 -= p0
        _ = fi36
        _ = v42
    }
    fmt.Println(p0)
    _ = p0
    return 1
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := []string{""}
    v3 := make(map[int]uint32)
    var v4 int = 32768
    var v5 int8 = 81
    var v6 bool = true
    if v6 {
        var v43 string = ("hello" + "\n\t")
        if (((!v6) && (42 >= 42)) && (v6 && v6)) {
            var v44 int64 = -5
            v6 = v6
            v2 = append(v2, "\n\t")
            _ = v44
        }
        if false {
            var v45 bool = (true || false)
            if v6 {
                for fi46 := 0; fi46 < 5; fi46++ {
                    var v47 uint8 = 42
                    v2 = append(v2, ("foo bar" + v43))
                    if (-32768 < 1000) {
                        var v48 string = v43
                        _ = v48
                    }
                    _ = fi46
                    _ = v47
                }
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            } else {
                var v49 float64 = 1e10
                for fi50 := 0; fi50 < 7; fi50++ {
                    for fi51 := 0; fi51 < 5; fi51++ {
                        switch v43 {
                        case "foo bar":
                            if v45 {
                                var v52 int16 = -32768
                                _ = v52
                            } else {
                                var v53 float64 = float64(-2.25)
                                if (4294967295 != 282) {
                                    var v54 int = 32767
                                    v6 = v45
                                    v43 = v43
                                    _ = v54
                                } else {
                                    var v55 int = fi51
                                    _ = v55
                                }
                                v2 = append(v2, ("a" + "hello"))
                                v56 := GenS2{F0: -1, F1: 42, F2: 42}
                                v57 := GenS2{F0: 655, F1: -9223372036854775808, F2: 4294967295}
                                _ = v53
                                _ = v56
                                _ = v57
                            }
                        case "hello":
                            if v45 {
                                switch (v43 + "foo bar") {
                                case "\n\t":
                                    v58 := GenS1{F0: (-1 != 1000), F1: 18446744073709551615}
                                    _ = v58
                                case "hello":
                                    if true {
                                        v59 := v1[2147483647]
                                        v49 -= v49
                                        _ = v59
                                    }
                                case "":
                                    if (2 < 1) {
                                        v49 = v49
                                        delete(v1, 0)
                                    }
                                    v43 = "a"
                                default:
                                    if ((1 < -1) || (3.14159 >= ((3.14159 * (3.14159 - (3.14159 * float32((v49 - (((float64(float32(v49)) * ((v49 * (v49 * v49)) + v49)) / 1.0) - v49)))))) / float32(v49)))) {
                                        v60 := v1[2147483648]
                                        v43 = ((v43 + ((((((v43 + "a") + ("a" + (((v43 + ("a" + v43)) + "hello") + v43))) + v43) + v43) + (v43 + ("hello" + v43))) + (("a" + "hello") + "\n\t"))) + v43)
                                        var v61 uint32 = 433
                                        _ = v60
                                        _ = v61
                                    } else {
                                        if (42 == 1) {
                                            fmt.Println(v5)
                                            var v62 int8 = (v5 * -11)
                                            if (173 >= 1) {
                                                var v63 string = v43
                                                var v64 int32 = 42
                                                var v65 uint32 = 791
                                                v62 -= v62
                                                if v45 {
                                                    var v66 int16 = 32767
                                                    _ = v66
                                                }
                                                _ = v63
                                                _ = v64
                                                _ = v65
                                            }
                                            _ = v62
                                        } else {
                                            v67 := GenS1{F0: v45, F1: 2}
                                            _ = v67
                                        }
                                        switch (-(fi50)) {
                                        case 1:
                                            switch ("foo bar" + "a") {
                                            case "\n\t":
                                                v68 := GenS3{F0: genFn1(1), F1: (false || (v5 > v5)), F2: 42}
                                                _ = v68
                                            case "":
                                                switch 127 {
                                                case -379:
                                                    var v69 uint64 = 366
                                                    _ = v69
                                                default:
                                                    if v6 {
                                                        var v70 int8 = (v5 + v5)
                                                        var v71 uint = 255
                                                        _ = v70
                                                        _ = v71
                                                    }
                                                }
                                                v6 = (v6 || ((3.14159 * v49) >= v49))
                                            case "a":
                                                if (914 < 1) {
                                                    v72 := v3[(fi50 % ((0) | 1))]
                                                    switch v43 {
                                                    case "foo bar":
                                                        v43 = v43
                                                    }
                                                    _ = v72
                                                } else {
                                                    v1[4294967295] = v45
                                                    var v73 uint = 2
                                                    _ = v73
                                                }
                                            default:
                                                v74 := GenS3{F0: -2147483648, F1: true, F2: 43}
                                                _ = v74
                                            }
                                        }
                                    }
                                }
                            }
                            v6 = (true || false)
                        default:
                            var v75 bool = (!(v45 || v6))
                            _ = v75
                        }
                        v76 := GenS1{F0: ((18446744073709551615 > 445) && (1 != 1)), F1: 1}
                        v77 := GenS3{F0: 1, F1: (!(37 < genFn1(255))), F2: 401}
                        _ = fi51
                        _ = v76
                        _ = v77
                    }
                    _ = fi50
                }
                if v6 {
                    v78 := GenS1{F0: v6, F1: 350}
                    for fi79 := 0; fi79 < 16; fi79++ {
                        fmt.Println(v49, v4, v4)
                        fmt.Println(v49)
                        _ = fi79
                    }
                    v4 += (v4 & 32767)
                    _ = v78
                }
                _ = v49
            }
            var v80 uint64 = 2
            _ = v45
            _ = v80
        } else {
            v81 := GenS2{F0: 1, F1: -929, F2: 508}
            switch (v4 / ((v4) | 1)) {
            case -2147483648:
                for fi82 := 0; fi82 < 19; fi82++ {
                    for fi83 := 0; fi83 < 2; fi83++ {
                        var v84 uint8 = 127
                        v85 := GenS1{F0: v6, F1: 4294967295}
                        switch v43 {
                        case "\n\t":
                            fmt.Println(v4)
                        default:
                            var v86 uint32 = 2147483648
                            _ = v86
                        }
                        _ = fi83
                        _ = v84
                        _ = v85
                    }
                    _ = fi82
                }
            default:
                if ((4294967295 > 1) && v6) {
                    var v87 uint8 = 255
                    _ = v87
                }
            }
            _ = v81
        }
        _ = v43
    } else {
        if (1e10 <= ((0.1 - 1e10) * (float64(1e10) + (0.1 + float64(1e10))))) {
            v88 := GenS4{F0: 4294967295}
            _ = v88
        }
        for fi89 := 0; fi89 < 20; fi89++ {
            for fi90 := 0; fi90 < 18; fi90++ {
                v5 += (v5 | 1)
                v91 := v1[42]
                _ = fi90
                _ = v91
            }
            _ = fi89
        }
        v5 = v5
    }
    var v92 uint32 = 406
    switch (v4 + v4) {
    case 128:
        v93 := GenS5{F0: v4, F1: v6}
        _ = v93
    default:
        v94 := v1[v92]
        _ = v94
    }
    v4 -= v4
    for fi95 := 0; fi95 < 10; fi95++ {
        v96 := GenS2{F0: 1, F1: 9223372036854775807, F2: 559}
        _ = fi95
        _ = v96
    }
    var v97 int8 = v5
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v6 = ((("a" + "foo bar") + ("hello" + ("\n\t" + "a"))) > "\n\t")
    v6 = true
    _ = len(v3)
    var v98 uint = 65535
    var v99 string = "hello"
    if (0.0 == (0.0 + float32(3.14159))) {
        var v100 uint = v98
        _ = v100
    } else {
        switch -102 {
        case -9223372036854775808:
            v101 := GenS6{F0: 2, F1: v98, F2: 1}
            _ = v101
        case -154:
            if v6 {
                switch 2147483647 {
                case 4294967295:
                    for fi102 := 0; fi102 < 14; fi102++ {
                        v6 = v6
                        _ = fi102
                    }
                default:
                    if true {
                        if (v6 && v6) {
                            var v103 string = v99
                            _ = v103
                        }
                    }
                }
                v3[v4] = (v92 % ((v92) | 1))
            } else {
                if (((((1.5 * ((0.1 - (3.14159 - 1e10)) * (0.0 - 3.14159))) / (1e10 * 1e10)) / (float64((1.5 / 1e10)) * -2.25)) / 1.5) == 1.5) {
                    for fi104 := 0; fi104 < 13; fi104++ {
                        v99 = (v99 + v99)
                        _ = fi104
                    }
                    var v105 uint64 = 849
                    _ = v105
                }
                fmt.Println(v4, v98, v97)
                if v6 {
                    switch v99 {
                    case "":
                        for fi106 := 0; fi106 < 14; fi106++ {
                            fmt.Println(v6, v98, v6)
                            fmt.Println(v92, fi106)
                            switch fi106 {
                            case 582:
                                for fi107 := 0; fi107 < 8; fi107++ {
                                    if false {
                                        switch v4 {
                                        case 42:
                                            switch "a" {
                                            case "foo bar":
                                                v99 = "foo bar"
                                            case "a":
                                                if v6 {
                                                    v99 = (v99 + v99)
                                                }
                                            case "hello":
                                                v5 = (v97 + 0)
                                                v108 := GenS7{F0: fi107, F1: -9223372036854775808}
                                                var v109 int = (v4 | v4)
                                                v92 = 65535
                                                fmt.Println(v109, v4)
                                                _ = v108
                                                _ = v109
                                            }
                                            v5 = int8(v92)
                                            v110 := GenS7{F0: (v4 ^ (fi107 << (((^(v98))) & 63))), F1: -1}
                                            _ = v110
                                        case 0:
                                            v5 -= (^(v97))
                                            v99 = "foo bar"
                                            switch v99 {
                                            case "foo bar":
                                                fmt.Println(v6, v6)
                                                _ = len(v1)
                                                if true {
                                                    v6 = v6
                                                    fmt.Println(fi107, v4)
                                                    v99 = "a"
                                                } else {
                                                    var v111 uint16 = 557
                                                    v112 := GenS8{F0: -9223372036854775808}
                                                    switch "" {
                                                    case "\n\t":
                                                        var v113 int8 = 0
                                                        _ = v113
                                                    }
                                                    _ = v111
                                                    _ = v112
                                                }
                                            default:
                                                var v114 uint8 = 237
                                                _ = v114
                                            }
                                        default:
                                            if (!v6) {
                                                if false {
                                                    var v115 int16 = 1
                                                    _ = v115
                                                }
                                                v116 := v3[(v4 ^ 426)]
                                                var v117 uint64 = 474
                                                _ = len(v2)
                                                _ = v116
                                                _ = v117
                                            } else {
                                                var v118 bool = ((v4 & 2147483647) >= (fi107 << ((255) & 63)))
                                                fmt.Println(fi106, v92, v92)
                                                _ = v118
                                            }
                                        }
                                        v119 := GenS9{F0: fi106}
                                        var v120 int8 = 17
                                        _ = v119
                                        _ = v120
                                    } else {
                                        v121 := GenS9{F0: 9223372036854775807}
                                        v122 := GenS10{F0: (fi107 | fi107)}
                                        _ = v121
                                        _ = v122
                                    }
                                    _ = fi107
                                }
                            default:
                                for fi123 := 0; fi123 < 14; fi123++ {
                                    v6 = (-1 == -1)
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    _ = fi123
                                }
                            }
                            _ = fi106
                        }
                    case "hello":
                        var v124 uint8 = 42
                        _ = v124
                    case "foo bar":
                        var v125 bool = true
                        _ = v125
                    default:
                        v126 := GenS11{F0: genFn0(v98)}
                        _ = v126
                    }
                    for fi127 := 0; fi127 < 4; fi127++ {
                        delete(v3, fi127)
                        _ = fi127
                    }
                } else {
                    if true {
                        switch -2147483648 {
                        case -471:
                            if len(v2) > 0 {
                                v2[0] = "a"
                            }
                        case -2147483648:
                            switch (("hello" + v99) + (("" + ("" + "")) + ((v99 + "\n\t") + v99))) {
                            case "foo bar":
                                for fi128 := 0; fi128 < 15; fi128++ {
                                    switch "hello" {
                                    case "\n\t":
                                        var v129 float32 = (((1.5 + 0.0) * 0.0) - (3.14159 * 0.0))
                                        _ = v129
                                    case "foo bar":
                                        if false {
                                            if (v98 < 65535) {
                                                _ = len(v1)
                                                if v6 {
                                                    for fi130 := 0; fi130 < 17; fi130++ {
                                                        v98 = (v98 >> ((v98) & 63))
                                                        var v131 uint16 = 65535
                                                        _ = fi130
                                                        _ = v131
                                                    }
                                                    var v132 float64 = 3.14159
                                                    _ = v132
                                                } else {
                                                    v133 := v3[(fi128 / ((v4) | 1))]
                                                    _ = v133
                                                }
                                            } else {
                                                v134 := v3[-2147483648]
                                                v2 = append(v2, v99)
                                                _ = v134
                                            }
                                            v5 = (v97 & -26)
                                        }
                                    }
                                    var v135 bool = v6
                                    v98 -= (v98 % (((v98 - v98)) | 1))
                                    _ = fi128
                                    _ = v135
                                }
                            case "hello":
                                if true {
                                    v92 = v92
                                }
                            case "\n\t":
                                if false {
                                    v98 = (^(v98))
                                }
                            default:
                                fmt.Println(v4, v99)
                            }
                        }
                        v136 := GenS12{F0: 123}
                        var v137 int16 = -954
                        _ = v136
                        _ = v137
                    } else {
                        var v138 uint8 = 42
                        v92 -= 388
                        var v139 string = ("a" + v99)
                        _ = v138
                        _ = v139
                    }
                }
            }
        default:
            if len(v2) > 0 {
                v2[0] = ("" + (v99 + v99))
            }
        }
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v92:", v92)
    fmt.Println("v97:", v97)
    fmt.Println("v98:", v98)
    fmt.Println("v99:", v99)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v92
    _ = v97
    _ = v98
    _ = v99
}
