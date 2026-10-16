package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int }
type GenS2 struct { F0 bool; F1 int }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int8; F1 uint16; F2 int32 }
type GenS5 struct { F0 uint16; F1 int8 }
type GenS6 struct { F0 uint64; F1 uint16; F2 bool }
type GenS7 struct { F0 uint16; F1 bool; F2 int32 }
type GenS8 struct { F0 string; F1 uint64 }
type GenS9 struct { F0 int8 }
type GenS10 struct { F0 int; F1 int16; F2 int64 }
type GenS11 struct { F0 uint16 }
type GenS12 struct { F0 string; F1 bool; F2 int64 }
type GenS13 struct { F0 string; F1 int8 }
type GenS14 struct { F0 bool; F1 int64 }
type GenS15 struct { F0 int32 }
func genFn0(p0 uint, p1 int16) uint {
    p0 = (p0 & (^(p0)))
    var v7 bool = (955 < 0)
    for fi8 := 0; fi8 < 14; fi8++ {
        v9 := []int{(fi8 * fi8), fi8, (fi8 * (-(fi8))), fi8}
        v10 := GenS1{F0: "hello", F1: fi8}
        switch fi8 {
        case 1:
            var v11 string = ("hello" + "hello")
            _ = v11
        }
        _ = fi8
        _ = v10
        _ = v9
    }
    fmt.Println(p0, p0)
    v12 := []int{1, 4294967295, -743, 565}
    _ = p0
    _ = p1
    _ = v12
    _ = v7
    return (p0 - p0)
}
func genFn1(p0 bool) uint8 {
    var v13 int16 = 662
    for fi14 := 0; fi14 < 20; fi14++ {
        v15 := []int{fi14, (fi14 | fi14), -1, fi14}
        if p0 {
            var v16 uint16 = 2
            for fi17 := 0; fi17 < 1; fi17++ {
                var v18 uint8 = 0
                _ = fi17
                _ = v18
            }
            for fi19 := 0; fi19 < 16; fi19++ {
                switch "hello" {
                case "hello":
                    p0 = (!(127 > -122))
                    v13 = (v13 << ((42) & 63))
                    switch "" {
                    case "foo bar":
                        if (!true) {
                            var v20 uint16 = 1
                            var v21 uint64 = 947
                            _ = v20
                            _ = v21
                        }
                    case "hello":
                        v13 += (v13 | (v13 >> ((557) & 63)))
                    }
                    switch "\n\t" {
                    case "\n\t":
                        if p0 {
                            var v22 uint64 = 352
                            v16 += v16
                            _ = v22
                        }
                        v23 := GenS1{F0: "", F1: fi14}
                        fmt.Println(v16, p0)
                        if (!((^(v13)) == (v13 << ((65535) & 63)))) {
                            v24 := GenS2{F0: p0, F1: (fi19 | (fi19 / (((fi14 | (-(fi14)))) | 1)))}
                            _ = v24
                        }
                        switch fi19 {
                        case -1:
                            fmt.Println(v13, v13, p0)
                        case 18:
                            if p0 {
                                if len(v15) >= 2 {
                                    v15 = v15[1:len(v15)]
                                }
                                var v25 uint64 = 42
                                v16 -= v16
                                _ = v25
                            }
                        case 128:
                            fmt.Println(fi19, fi14)
                        default:
                            var v26 uint = 4294967295
                            _ = v26
                        }
                        _ = v23
                    case "hello":
                        var v27 float64 = 1e10
                        _ = v27
                    case "":
                        p0 = (true || p0)
                        switch (-(fi14)) {
                        case 32768:
                            v13 = (v13 + (v13 >> ((700) & 63)))
                        case 827:
                            var v28 int32 = 729
                            _ = v28
                        default:
                            if (!(0.1 != 0.0)) {
                                fmt.Println(p0, v13)
                                var v29 int = (^(fi14))
                                if p0 {
                                    if (!p0) {
                                        v30 := GenS1{F0: (("" + "foo bar") + "hello"), F1: (fi19 << ((1) & 63))}
                                        var v31 uint64 = 18446744073709551615
                                        p0 = (false || false)
                                        _ = v30
                                        _ = v31
                                    } else {
                                        v29 = fi14
                                        if len(v15) > 0 {
                                            v15[0] = (v29 + (fi19 - (fi14 * -782)))
                                        }
                                        if len(v15) > 0 {
                                            v15[0] = fi19
                                        }
                                    }
                                    var v32 int = fi14
                                    fmt.Println(v32, fi14, fi14)
                                    fmt.Println(v29, fi19, v16)
                                    _ = v32
                                }
                                fmt.Println(v16, p0, p0)
                                var v33 int64 = 42
                                _ = v29
                                _ = v33
                            } else {
                                var v34 int8 = 127
                                _ = v34
                            }
                            p0 = (((("\n\t" + "hello") + "") < (("\n\t" + "") + "a")) && (2147483647 >= 42))
                            v13 = -575
                            if len(v15) > 0 {
                                v15[0] = -9223372036854775808
                            }
                        }
                        v16 -= 0
                    default:
                        if p0 {
                            fmt.Println(v16)
                        }
                    }
                case "foo bar":
                    var v35 uint64 = 737
                    _ = v35
                case "":
                    if ((0.0 != 1e10) && (!(2 < 255))) {
                        var v36 float64 = (float64(3.14159) / (0.1 / 3.14159))
                        _ = len(v15)
                        switch fi14 {
                        case 127:
                            if (genFn0(609, (v13 + -32768)) == 65535) {
                                v36 = (v36 / (0.1 / v36))
                            } else {
                                fmt.Println(p0, v16)
                                if ((-(v13)) == (^(v13))) {
                                    if (v16 <= (v16 - (v16 & (v16 + v16)))) {
                                        if (v13 <= v13) {
                                            if (0.0 < (1e10 + v36)) {
                                                if true {
                                                    v16 = v16
                                                } else {
                                                    p0 = p0
                                                    var v37 uint64 = 34
                                                    _ = v37
                                                }
                                            } else {
                                                if (-32768 != (^(v13))) {
                                                    if true {
                                                        v38 := GenS3{F0: 0}
                                                        v39 := GenS3{F0: 1}
                                                        if (!(p0 && true)) {
                                                            v40 := GenS3{F0: 4294967295}
                                                            var v41 uint = 2
                                                            fmt.Println(v41)
                                                            _ = v40
                                                            _ = v41
                                                        } else {
                                                            fmt.Println(fi14)
                                                            fmt.Println(fi19)
                                                            var v42 uint16 = v16
                                                            v43 := GenS4{F0: 42, F1: (v42 & v42), F2: 42}
                                                            _ = v42
                                                            _ = v43
                                                        }
                                                        var v44 int32 = 1
                                                        fmt.Println(fi19, v44)
                                                        _ = v38
                                                        _ = v39
                                                        _ = v44
                                                    } else {
                                                        v16 = (v16 ^ (v16 & 255))
                                                        v45 := GenS5{F0: v16, F1: -128}
                                                        if p0 {
                                                            fmt.Println(v16, fi14, v13)
                                                            if len(v15) > 0 {
                                                                v15[0] = 1
                                                            }
                                                            v36 = (v36 - -2.25)
                                                            v16 = 32767
                                                            var v46 float32 = 1.5
                                                            p0 = false
                                                            _ = v46
                                                        } else {
                                                            switch fi19 {
                                                            case -9223372036854775808:
                                                                if len(v15) >= 2 {
                                                                    v15 = v15[1:len(v15)]
                                                                }
                                                            case 32767:
                                                                var v47 uint8 = 42
                                                                _ = v47
                                                            default:
                                                                if (float64(float32((v36 + -2.25))) >= (v36 * v36)) {
                                                                    v48 := GenS6{F0: 591, F1: (v16 << ((18446744073709551615) & 63)), F2: false}
                                                                    v36 -= v36
                                                                    v49 := GenS7{F0: 2, F1: (-88 != 100), F2: 42}
                                                                    _ = v48
                                                                    _ = v49
                                                                }
                                                            }
                                                        }
                                                        _ = v45
                                                    }
                                                } else {
                                                    switch fi14 {
                                                    case -958:
                                                        v50 := GenS8{F0: ("" + "foo bar"), F1: 2}
                                                        _ = v50
                                                    case 9223372036854775807:
                                                        fmt.Println(fi14)
                                                        v16 += v16
                                                    }
                                                    if p0 {
                                                        if len(v15) >= 2 {
                                                            v15 = v15[1:len(v15)]
                                                        }
                                                        if len(v15) > 0 {
                                                            v15[0] = fi14
                                                        }
                                                        var v51 uint16 = v16
                                                        _ = v51
                                                    }
                                                }
                                            }
                                            if false {
                                                fmt.Println(v36, p0)
                                            } else {
                                                p0 = (!(-1 != -84))
                                            }
                                            v52 := GenS7{F0: v16, F1: p0, F2: 0}
                                            var v53 uint16 = (v16 | 870)
                                            _ = v52
                                            _ = v53
                                        }
                                    } else {
                                        var v54 int16 = v13
                                        var v55 uint64 = 0
                                        var v56 uint8 = 127
                                        _ = v54
                                        _ = v55
                                        _ = v56
                                    }
                                    var v57 int = 128
                                    var v58 uint8 = 42
                                    _ = v57
                                    _ = v58
                                } else {
                                    v16 = v16
                                    if (fi14 == 127) {
                                        v13 -= 662
                                        v59 := GenS4{F0: -117, F1: 0, F2: 407}
                                        v60 := GenS9{F0: -93}
                                        _ = v59
                                        _ = v60
                                    }
                                }
                                v61 := GenS10{F0: fi14, F1: (-(v13)), F2: -9223372036854775808}
                                _ = v61
                            }
                        case 32767:
                            if false {
                                v36 = (v36 * v36)
                                p0 = (1 != 9223372036854775807)
                                v62 := GenS11{F0: v16}
                                _ = v62
                            }
                        case 269:
                            var v63 uint = 18446744073709551615
                            _ = v63
                        }
                        _ = v36
                    }
                }
                _ = fi19
            }
            _ = v16
        }
        for fi64 := 0; fi64 < 15; fi64++ {
            fmt.Println(fi64, fi64)
            v15 = append(v15, (fi64 + 42))
            _ = fi64
        }
        _ = fi14
        _ = v15
    }
    _ = p0
    _ = v13
    return 0
}

func main() {
    var v1 string = "hello"
    v2 := make(map[string]bool)
    var v3 uint32 = 0
    var v4 uint = 255
    var v5 uint8 = 0
    var v6 uint8 = (^(v5))
    v65 := []int{-9223372036854775808, -663}
    _ = len(v65)
    var v66 uint32 = v3
    switch 128 {
    case 9223372036854775807:
        var v67 float64 = 0.0
        _ = v67
    case 32768:
        for fi68 := 0; fi68 < 16; fi68++ {
            switch 671 {
            case 660:
                var v69 uint32 = (^(v66))
                _ = v69
            case 32768:
                v4 += (v4 ^ (v4 / (((v4 | (v4 ^ 255))) | 1)))
            }
            _ = fi68
        }
    default:
        v70 := GenS12{F0: ((v1 + v1) + (v1 + "a")), F1: (!true), F2: 42}
        _ = v70
    }
    var v71 int32 = 111
    v72 := GenS6{F0: 1, F1: 200, F2: true}
    fmt.Println(v1, v1, v66)
    for fi73 := 0; fi73 < 13; fi73++ {
        _ = len(v65)
        v6 += v5
        if false {
            var v74 float32 = (-2.25 / (float32(-2.25) + 1e10))
            for fi75 := 0; fi75 < 18; fi75++ {
                if (-1 >= -174) {
                    _ = len(v65)
                    var v76 uint16 = 1
                    _ = v76
                }
                v71 += (v71 - (v71 << ((v4) & 63)))
                v3 += (^(v66))
                fmt.Println(fi75, v5)
                switch v1 {
                case "\n\t":
                    v4 -= 595
                case "hello":
                    if len(v65) >= 2 {
                        v65 = v65[1:len(v65)]
                    }
                case "foo bar":
                    fmt.Println(v74)
                }
                _ = fi75
            }
            var v77 int16 = -32768
            _ = v74
            _ = v77
        }
        _ = fi73
    }
    var v78 int32 = v71
    if ((-(v71)) >= -495) {
        for fi79 := 0; fi79 < 10; fi79++ {
            fmt.Println(v3, v4, v66)
            _ = fi79
        }
        for fi80 := 0; fi80 < 5; fi80++ {
            v1 = ("foo bar" + "foo bar")
            v1 = "foo bar"
            if true {
                if (1 <= (v4 << (((v4 & 42)) & 63))) {
                    if (-87 > 42) {
                        v65 = append(v65, fi80)
                        _ = len(v65)
                    }
                    var v81 float64 = -2.25
                    v82 := GenS7{F0: 65535, F1: ((!true) && true), F2: (v71 ^ 42)}
                    _ = v81
                    _ = v82
                } else {
                    if ((fi80 << ((genFn0(v4, 42)) & 63)) != 601) {
                        if len(v65) >= 2 {
                            v65 = v65[1:len(v65)]
                        }
                    } else {
                        v65 = append(v65, (fi80 * 4294967295))
                        if (!(false && ((v78 + (v78 - v71)) <= 42))) {
                            v1 = (("" + "\n\t") + ("\n\t" + v1))
                            for fi83 := 0; fi83 < 4; fi83++ {
                                fmt.Println(v5, v6)
                                switch v1 {
                                case "":
                                    v4 -= genFn0(v4, 1)
                                case "foo bar":
                                    if (v4 <= genFn0(v4, 42)) {
                                        v84 := v2[v1]
                                        _ = v84
                                    }
                                case "a":
                                    var v85 uint = genFn0(uint(v6), -32768)
                                    _ = v85
                                default:
                                    v5 = 255
                                }
                                _ = fi83
                            }
                            v2["\n\t"] = (v3 < (v66 & v66))
                        } else {
                            if true {
                                v3 = v3
                                if (100 <= 0) {
                                    if false {
                                        var v86 float32 = 0.0
                                        switch (v1 + ((v1 + "hello") + (("hello" + v1) + v1))) {
                                        case "a":
                                            _ = len(v65)
                                        case "":
                                            v6 -= (^(v5))
                                        case "\n\t":
                                            var v87 int8 = int8(v6)
                                            _ = v87
                                        default:
                                            v71 += v78
                                        }
                                        _ = v86
                                    } else {
                                        var v88 uint8 = (^(v5))
                                        _ = v88
                                    }
                                }
                                fmt.Println(v4)
                            } else {
                                for fi89 := 0; fi89 < 1; fi89++ {
                                    if len(v65) > 0 {
                                        v65[0] = (fi89 ^ fi89)
                                    }
                                    _ = fi89
                                }
                                v66 = (v3 * v3)
                                for fi90 := 0; fi90 < 19; fi90++ {
                                    v4 = uint(v5)
                                    _ = fi90
                                }
                            }
                            if (!false) {
                                v2[v1] = false
                                var v91 uint = 0
                                v3 += 0
                                _ = v91
                            } else {
                                _ = len(v65)
                            }
                        }
                    }
                }
                switch (fi80 >> ((v4) & 63)) {
                case 128:
                    fmt.Println(v66, v1)
                default:
                    for fi92 := 0; fi92 < 18; fi92++ {
                        if (true || (!(2 >= 2))) {
                            v6 -= genFn1((v4 <= 435))
                            if true {
                                v71 -= -2147483648
                                v78 = 0
                                v71 += (v71 - v71)
                            } else {
                                v93 := GenS13{F0: v1, F1: 100}
                                _ = v93
                            }
                            var v94 uint8 = (v6 ^ v6)
                            _ = v94
                        } else {
                            var v95 int8 = int8(v66)
                            _ = v95
                        }
                        fmt.Println(fi80)
                        _ = fi92
                    }
                }
            } else {
                _ = len(v65)
            }
            _ = fi80
        }
    }
    delete(v2, ("a" + "foo bar"))
    if false {
        var v96 int8 = 64
        _ = v96
    }
    if false {
        var v97 float32 = -2.25
        var v98 bool = (!(v78 < (^(v78))))
        _ = v97
        _ = v98
    } else {
        for fi99 := 0; fi99 < 16; fi99++ {
            if (-1 <= 1) {
                for fi100 := 0; fi100 < 6; fi100++ {
                    fmt.Println(v4, v71, v6)
                    if (!(1 == -93)) {
                        if true {
                            v101 := GenS12{F0: "foo bar", F1: (int8(v71) == 4), F2: 828}
                            var v102 float64 = ((1e10 / 0.1) - float64(3.14159))
                            var v103 uint64 = 2
                            v104 := GenS14{F0: true, F1: 1}
                            v78 = (v78 + v71)
                            _ = v101
                            _ = v102
                            _ = v103
                            _ = v104
                        } else {
                            if false {
                                v105 := v2["foo bar"]
                                _ = v105
                            } else {
                                v71 = (v71 >> (((v4 & (v4 << ((0) & 63)))) & 63))
                                if true {
                                    var v106 int8 = 127
                                    v3 = (v3 - 667)
                                    var v107 string = v1
                                    _ = v106
                                    _ = v107
                                }
                                v108 := GenS7{F0: 32767, F1: true, F2: 0}
                                _ = v108
                            }
                        }
                        if true {
                            var v109 uint16 = 42
                            if false {
                                switch "" {
                                case "foo bar":
                                    switch 4294967295 {
                                    case 0:
                                        if true {
                                            v110 := v2["\n\t"]
                                            v110 = true
                                            if v110 {
                                                v111 := GenS10{F0: fi99, F1: 703, F2: int64(fi100)}
                                                _ = v111
                                            }
                                            _ = v110
                                        }
                                    case 32768:
                                        v3 = (v66 - (v3 - v3))
                                    default:
                                        v112 := GenS4{F0: 0, F1: v109, F2: v78}
                                        _ = v112
                                    }
                                }
                                _ = len(v2)
                                delete(v2, "hello")
                            }
                            v65 = append(v65, fi100)
                            _ = v109
                        } else {
                            var v113 float32 = (((1.5 + 1.5) / 1.0) / 3.14159)
                            _ = v113
                        }
                        v114 := GenS5{F0: 576, F1: int8(v3)}
                        _ = v114
                    } else {
                        if true {
                            fmt.Println(fi100, v78, v78)
                        }
                    }
                    v66 += v3
                    v1 = v1
                    if (false && (166 < 2)) {
                        if len(v65) >= 2 {
                            v65 = v65[1:len(v65)]
                        }
                    } else {
                        var v115 int64 = 0
                        fmt.Println(v115, v5, fi99)
                        if ((((v4 ^ v4) != v4) && (255 != 260)) || true) {
                            if true {
                                switch -17 {
                                case 32767:
                                    var v116 float64 = 1.5
                                    _ = v116
                                case 9223372036854775807:
                                    var v117 bool = true
                                    v65 = append(v65, (fi100 % (((fi99 | (fi100 - (fi100 + (fi100 << ((v4) & 63)))))) | 1)))
                                    v118 := GenS6{F0: 2, F1: 1, F2: v117}
                                    v119 := GenS14{F0: v117, F1: (^(v115))}
                                    switch v1 {
                                    case "\n\t":
                                        if len(v65) >= 2 {
                                            v65 = v65[1:len(v65)]
                                        }
                                    case "":
                                        if v117 {
                                            v65 = append(v65, fi99)
                                        } else {
                                            fmt.Println(v78, fi99, v6)
                                        }
                                    }
                                    _ = v117
                                    _ = v118
                                    _ = v119
                                case -1:
                                    var v120 uint32 = 2147483647
                                    _ = v120
                                }
                            } else {
                                v66 -= (^(v3))
                                v121 := v2[v1]
                                _ = v121
                            }
                            if (v115 == v115) {
                                v1 = v1
                            } else {
                                v122 := GenS15{F0: (v71 % ((-2147483648) | 1))}
                                _ = v122
                            }
                            var v123 int = fi100
                            _ = v123
                        }
                        _ = v115
                    }
                    _ = fi100
                }
                var v124 int8 = 108
                _ = v124
            }
            fmt.Println(v78, v3, v71)
            switch fi99 {
            case -9223372036854775808:
                var v125 float32 = 1.5
                _ = v125
            case 128:
                if false {
                    switch "a" {
                    case "\n\t":
                        var v126 uint32 = v3
                        _ = v126
                    case "a":
                        v71 = -100
                    }
                }
            default:
                v127 := GenS6{F0: 4294967295, F1: 2, F2: false}
                _ = v127
            }
            _ = fi99
        }
    }
    v128 := v2[("foo bar" + "hello")]
    fmt.Println("v1:", v1)
    fmt.Println("v128:", v128)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v65:", len(v65), v65)
    fmt.Println("v66:", v66)
    fmt.Println("v71:", v71)
    fmt.Println("v78:", v78)
    _ = v1
    _ = v128
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v65
    _ = v66
    _ = v71
    _ = v72
    _ = v78
}
