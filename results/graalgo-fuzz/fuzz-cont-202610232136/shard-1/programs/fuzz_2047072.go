package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64 }
type GenS2 struct { F0 int32; F1 string; F2 uint8 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int64; F1 uint }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 int; F1 uint16; F2 uint8 }
type GenS7 struct { F0 uint16; F1 uint8 }
type GenS8 struct { F0 uint8; F1 uint16; F2 int }
type GenS9 struct { F0 int32; F1 string }
type GenS10 struct { F0 int16; F1 int32; F2 uint16 }
type GenS11 struct { F0 uint16; F1 int64 }
type GenS12 struct { F0 uint64 }
type GenS13 struct { F0 int; F1 int; F2 uint8 }
type GenS14 struct { F0 int64; F1 int16 }
type GenS15 struct { F0 int8; F1 int; F2 uint64 }
type GenS16 struct { F0 int16; F1 uint64 }
type GenS17 struct { F0 uint32; F1 int16; F2 int8 }
type GenS18 struct { F0 int16; F1 int32; F2 uint }
type GenS19 struct { F0 uint; F1 uint; F2 uint8 }
type GenS20 struct { F0 int }
type GenS21 struct { F0 int8 }

func main() {
    var v1 float64 = -2.25
    var v2 bool = false
    v3 := make(map[int]bool)
    v4 := make(map[uint32]uint32)
    fmt.Println(v2, v1)
    fmt.Println(v1, v1)
    var v5 int16 = 177
    if v2 {
        var v6 int64 = 0
        _ = v6
    } else {
        v7 := GenS1{F0: 42, F1: -1}
        _ = v7
    }
    v5 -= (^(v5))
    for fi8 := 0; fi8 < 10; fi8++ {
        var v9 bool = (!false)
        var v10 int8 = 1
        _ = fi8
        _ = v10
        _ = v9
    }
    var v11 string = ("" + "")
    fmt.Println(v2, v5, v5)
    var v12 int64 = 9223372036854775807
    if (!true) {
        if (v12 == v12) {
            v2 = (1e10 < 1e10)
        }
        if true {
            v13 := GenS1{F0: 0, F1: v12}
            _ = v13
        }
    } else {
        for fi14 := 0; fi14 < 15; fi14++ {
            switch 9223372036854775807 {
            case -208:
                if (v2 && (v11 == v11)) {
                    var v15 int = (fi14 << ((4294967295) & 63))
                    v12 = (v12 % ((v12) | 1))
                    _ = v15
                } else {
                    fmt.Println(v12, v1, v11)
                    v11 = v11
                }
            case 128:
                if (v5 != v5) {
                    for fi16 := 0; fi16 < 18; fi16++ {
                        var v17 float64 = v1
                        var v18 string = ("\n\t" + "")
                        v19 := []int{(fi16 | fi14), 316}
                        _ = fi16
                        _ = v17
                        _ = v18
                        _ = v19
                    }
                    v20 := GenS2{F0: 42, F1: ((("foo bar" + v11) + "foo bar") + "foo bar"), F2: 74}
                    _ = v20
                } else {
                    for fi21 := 0; fi21 < 9; fi21++ {
                        v22 := GenS3{F0: (v5 & v5)}
                        var v23 float32 = float32(-2.25)
                        _ = fi21
                        _ = v22
                        _ = v23
                    }
                }
            }
            var v24 uint16 = 623
            _ = fi14
            _ = v24
        }
        v1 += v1
        var v25 uint8 = 2
        _ = v25
    }
    for fi26 := 0; fi26 < 7; fi26++ {
        for fi27 := 0; fi27 < 3; fi27++ {
            v1 += 0.1
            _ = fi27
        }
        _ = len(v4)
        _ = fi26
    }
    v2 = (2147483648 < 1)
    for fi28 := 0; fi28 < 18; fi28++ {
        var v29 uint8 = 128
        _ = fi28
        _ = v29
    }
    if (!(!("" >= v11))) {
        if v2 {
            fmt.Println(v2, v5)
            v1 = v1
            for fi30 := 0; fi30 < 2; fi30++ {
                v11 = v11
                for fi31 := 0; fi31 < 7; fi31++ {
                    v32 := GenS4{F0: v12, F1: 0}
                    _ = fi31
                    _ = v32
                }
                v2 = ((!v2) && ((v2 && v2) && false))
                _ = fi30
            }
        }
        for fi33 := 0; fi33 < 12; fi33++ {
            v1 -= (0.0 * v1)
            var v34 uint32 = 42
            if ((v2 && true) && v2) {
                var v35 uint = 255
                _ = v35
            }
            _ = fi33
            _ = v34
        }
        for fi36 := 0; fi36 < 17; fi36++ {
            for fi37 := 0; fi37 < 12; fi37++ {
                fmt.Println(v1, v12)
                var v38 float32 = ((1e10 / (1e10 - float32(float64(0.1)))) * float32(v1))
                if ((!v2) || ((!(0.1 >= -2.25)) && ((v38 != v38) || v2))) {
                    v39 := GenS5{F0: 136}
                    if (4294967295 >= 1) {
                        v12 = 671
                    } else {
                        v5 += -32768
                        v2 = ((v12 * -382) == 1)
                        var v40 int = fi36
                        _ = v40
                    }
                    var v41 int8 = 1
                    v5 -= (v5 - 1000)
                    fmt.Println(v5, v12, fi37)
                    var v42 int8 = 42
                    _ = v39
                    _ = v41
                    _ = v42
                }
                if v2 {
                    switch "" {
                    case "hello":
                        fmt.Println(fi37, v2, v38)
                    case "":
                        switch (fi37 | 128) {
                        case 2147483647:
                            if v2 {
                                v12 -= v12
                                v43 := GenS6{F0: fi36, F1: 32767, F2: 42}
                                _ = v43
                            } else {
                                v5 -= v5
                                fmt.Println(v38)
                                v44 := GenS7{F0: 2, F1: 249}
                                _ = v44
                            }
                        }
                    case "a":
                        fmt.Println(fi36, v5)
                    default:
                        var v45 int = fi36
                        _ = v45
                    }
                }
                _ = fi37
                _ = v38
            }
            for fi46 := 0; fi46 < 18; fi46++ {
                var v47 float64 = v1
                v48 := GenS8{F0: 119, F1: 255, F2: (fi36 - 127)}
                _ = fi46
                _ = v47
                _ = v48
            }
            _ = fi36
        }
    } else {
        switch -9223372036854775808 {
        case -9223372036854775808:
            if ((!false) || v2) {
                v2 = ((-337 > 32768) && ("\n\t" != "a"))
            }
        default:
            v12 += (v12 << ((2) & 63))
        }
        v5 -= 319
        fmt.Println(v2)
    }
    if true {
        switch v11 {
        case "hello":
            v5 = (v5 / (((v5 * v5)) | 1))
        case "a":
            var v49 uint8 = 1
            _ = v49
        case "foo bar":
            v50 := []int{}
            _ = v50
        default:
            if (false || (65535 <= 65535)) {
                var v51 uint8 = 1
                if v2 {
                    v52 := GenS7{F0: 255, F1: v51}
                    _ = v52
                }
                if true {
                    fmt.Println(v12)
                    if (v2 && v2) {
                        fmt.Println(v5)
                    }
                    if (!v2) {
                        switch "" {
                        case "a":
                            v53 := []int{-2147483648}
                            _ = v53
                        default:
                            v5 = v5
                        }
                    } else {
                        for fi54 := 0; fi54 < 14; fi54++ {
                            if true {
                                for fi55 := 0; fi55 < 9; fi55++ {
                                    var v56 float64 = 0.0
                                    v56 = (v56 / float64(float32(0.1)))
                                    _ = fi55
                                    _ = v56
                                }
                                for fi57 := 0; fi57 < 7; fi57++ {
                                    fmt.Println(v2, v51, v11)
                                    var v58 int64 = v12
                                    if v2 {
                                        switch fi57 {
                                        case 1:
                                            v59 := GenS8{F0: v51, F1: 65535, F2: fi57}
                                            _ = v59
                                        case 128:
                                            var v60 uint32 = 2147483647
                                            _ = v60
                                        default:
                                            v61 := []int{}
                                            _ = v61
                                        }
                                        v62 := []int{fi54, (fi54 % ((-79) | 1)), fi54, (fi54 - -1)}
                                        _ = v62
                                    }
                                    var v63 bool = ((((float32((v1 * float64(float32(((float64((1e10 / (((float32(((v1 - v1) - 3.14159)) * (3.14159 * -2.25)) / (float32(v1) * 0.1)) + float32(1.5)))) + 1e10) + (v1 - (((v1 - v1) - (v1 - (v1 + v1))) / 1.0))))))) * ((1e10 / ((-2.25 - -2.25) + 1e10)) * (1.5 + float32(1.5)))) * 0.0) - 0.1) < 0.1)
                                    v2 = ((0.0 / 1.0) > ((float32(float64(0.1)) + (3.14159 + 0.1)) / float32((((1e10 * v1) - v1) / float64(-2.25)))))
                                    var v64 float32 = (0.0 + -2.25)
                                    var v65 int32 = -1
                                    _ = fi57
                                    _ = v58
                                    _ = v63
                                    _ = v64
                                    _ = v65
                                }
                            }
                            for fi66 := 0; fi66 < 11; fi66++ {
                                var v67 string = v11
                                v68 := GenS9{F0: -1, F1: ("foo bar" + v11)}
                                if false {
                                    switch (fi54 << ((484) & 63)) {
                                    case -2147483648:
                                        switch (v67 + "foo bar") {
                                        case "":
                                            if v2 {
                                                v12 += (v12 / ((42) | 1))
                                                fmt.Println(v5)
                                                fmt.Println(v12, v11, v67)
                                            }
                                        case "a":
                                            v69 := GenS2{F0: 0, F1: v11, F2: v51}
                                            _ = v69
                                        default:
                                            v70 := GenS10{F0: (v5 | v5), F1: -691, F2: 2}
                                            _ = v70
                                        }
                                    case -178:
                                        if v2 {
                                            var v71 uint8 = (v51 / (((v51 & v51)) | 1))
                                            v72 := GenS11{F0: 65535, F1: (v12 << ((103) & 63))}
                                            v11 = "a"
                                            _ = v71
                                            _ = v72
                                        }
                                    }
                                    fmt.Println(fi54, fi66)
                                    delete(v4, 2147483648)
                                }
                                switch 2147483647 {
                                case -1:
                                    var v73 int64 = v12
                                    fmt.Println(v1)
                                    if (true && ((-21 < 1) && (18446744073709551615 < 4294967295))) {
                                        v12 = v73
                                        var v74 int8 = 100
                                        _ = v74
                                    } else {
                                        v75 := GenS10{F0: 1, F1: 848, F2: 42}
                                        if (v2 && v2) {
                                            v1 = v1
                                        } else {
                                            v67 = "a"
                                            if v2 {
                                                v73 = (v12 & -809)
                                                v5 = (v5 / (((v5 | (v5 - -592))) | 1))
                                            } else {
                                                delete(v3, (fi66 + fi66))
                                                v73 -= (v73 * (v73 / ((v12) | 1)))
                                                if v2 {
                                                    v73 = v73
                                                    v1 += v1
                                                }
                                            }
                                            var v76 int8 = 28
                                            var v77 string = v67
                                            _ = v76
                                            _ = v77
                                        }
                                        v78 := GenS11{F0: 963, F1: (^(v12))}
                                        _ = v75
                                        _ = v78
                                    }
                                    switch ("a" + "\n\t") {
                                    case "\n\t":
                                        fmt.Println(v51, v5, v2)
                                    case "":
                                        switch v11 {
                                        case "a":
                                            v79 := GenS7{F0: 2, F1: v51}
                                            _ = v79
                                        default:
                                            fmt.Println(v12)
                                        }
                                    }
                                    _ = v73
                                default:
                                    if false {
                                        var v80 int = fi66
                                        _ = v80
                                    }
                                }
                                _ = fi66
                                _ = v67
                                _ = v68
                            }
                            v1 += 3.14159
                            _ = fi54
                        }
                        var v81 float64 = v1
                        _ = v81
                    }
                } else {
                    switch (v11 + v11) {
                    case "hello":
                        v82 := v3[586]
                        _ = v82
                    default:
                        v83 := GenS12{F0: 42}
                        _ = v83
                    }
                }
                _ = v51
            } else {
                var v84 uint8 = 1
                var v85 int = 368
                _ = v84
                _ = v85
            }
        }
        switch -222 {
        case 32768:
            var v86 float64 = v1
            _ = v86
        case 127:
            var v87 float64 = v1
            _ = v87
        default:
            for fi88 := 0; fi88 < 1; fi88++ {
                v12 += v12
                _ = fi88
            }
        }
        if (-2.25 <= (float64(3.14159) / (1.5 + (v1 * float64(((-2.25 * -2.25) - 0.0)))))) {
            _ = len(v4)
            v89 := GenS2{F0: -1, F1: v11, F2: 2}
            var v90 int32 = 1
            _ = v89
            _ = v90
        }
    } else {
        switch -1 {
        case 333:
            var v91 uint8 = 2
            _ = v91
        default:
            if v2 {
                for fi92 := 0; fi92 < 1; fi92++ {
                    v93 := []int{}
                    for fi94 := 0; fi94 < 20; fi94++ {
                        _ = len(v3)
                        var v95 uint16 = 2
                        delete(v4, 1)
                        var v96 uint8 = 156
                        var v97 int = (^(fi92))
                        fmt.Println(v12, v95)
                        _ = fi94
                        _ = v95
                        _ = v96
                        _ = v97
                    }
                    if (!((true && v2) && true)) {
                        v1 = 1e10
                        fmt.Println(v11, v12)
                    } else {
                        switch "foo bar" {
                        case "foo bar":
                            v2 = (!v2)
                        case "\n\t":
                            v98 := GenS6{F0: fi92, F1: 32767, F2: 127}
                            _ = v98
                        case "hello":
                            if v2 {
                                var v99 float64 = float64((1.5 - ((float32((((v1 - -2.25) * v1) + v1)) / float32((((-2.25 + (v1 / v1)) - float64((0.0 / 3.14159))) + v1))) / float32((v1 + (v1 - ((float64(0.1) + ((float64(1e10) / 3.14159) / float64(1.5))) / 1.0)))))))
                                switch (fi92 ^ fi92) {
                                case 47:
                                    if v2 {
                                        v11 = ("hello" + ((((v11 + (v11 + (v11 + v11))) + "hello") + v11) + (v11 + (v11 + "\n\t"))))
                                        for fi100 := 0; fi100 < 12; fi100++ {
                                            v99 += v1
                                            if true {
                                                var v101 float32 = 0.0
                                                v102 := GenS13{F0: fi100, F1: (-(fi100)), F2: 127}
                                                switch "" {
                                                case "foo bar":
                                                    if false {
                                                        switch ((v11 + ((v11 + ("foo bar" + v11)) + v11)) + ((((v11 + v11) + (((v11 + v11) + v11) + ("\n\t" + v11))) + "hello") + v11)) {
                                                        case "foo bar":
                                                            v101 = (0.1 / (v101 + 1.5))
                                                        case "hello":
                                                            var v103 float64 = (3.14159 * (v99 * 1.5))
                                                            _ = v103
                                                        }
                                                    }
                                                case "a":
                                                    v104 := GenS12{F0: 2}
                                                    _ = v104
                                                default:
                                                    if (true && v2) {
                                                        v3[(^(fi92))] = v2
                                                        v105 := GenS6{F0: -903, F1: 255, F2: 1}
                                                        _ = v105
                                                    }
                                                }
                                                v99 -= v1
                                                fmt.Println(v101)
                                                _ = v101
                                                _ = v102
                                            } else {
                                                fmt.Println(v99, fi92, v2)
                                                var v106 uint8 = 229
                                                v12 = (-(v12))
                                                _ = v106
                                            }
                                            fmt.Println(v11)
                                            _ = fi100
                                        }
                                        if false {
                                            if (4294967295 > 65535) {
                                                v1 += v1
                                                if true {
                                                    var v107 float64 = (0.1 + (0.0 * 1e10))
                                                    _ = v107
                                                } else {
                                                    switch v11 {
                                                    case "a":
                                                        var v108 bool = (!(float64(1e10) == ((v99 * (0.1 - v1)) + ((((float64(-2.25) / v99) - (v99 * v99)) + ((1e10 - v99) / 1.5)) - (v99 * -2.25)))))
                                                        _ = v108
                                                    default:
                                                        v5 = v5
                                                    }
                                                    for fi109 := 0; fi109 < 13; fi109++ {
                                                        var v110 int8 = 0
                                                        _ = len(v93)
                                                        if (1e10 < 0.1) {
                                                            v1 -= float64(1.5)
                                                        } else {
                                                            if (v2 || (!("foo bar" > v11))) {
                                                                if ((v12 == 861) && (v2 || true)) {
                                                                    v99 += -2.25
                                                                }
                                                                if true {
                                                                    switch 32768 {
                                                                    case 1:
                                                                        if (((!v2) && (!v2)) && v2) {
                                                                            var v111 int64 = (v12 | (v12 + v12))
                                                                            if false {
                                                                                fmt.Println(fi109)
                                                                            } else {
                                                                                switch v11 {
                                                                                case "\n\t":
                                                                                    if len(v93) > 0 {
                                                                                        v93[0] = (fi92 / ((fi92) | 1))
                                                                                    }
                                                                                }
                                                                                var v112 bool = (1 >= 42)
                                                                                _ = v112
                                                                            }
                                                                            v111 -= v111
                                                                            _ = v111
                                                                        }
                                                                    case -2147483648:
                                                                        _ = len(v93)
                                                                    default:
                                                                        v113 := GenS14{F0: (v12 / (((v12 % ((v12) | 1))) | 1)), F1: v5}
                                                                        _ = v113
                                                                    }
                                                                }
                                                                v3[(fi92 << ((4294967295) & 63))] = (!((!v2) || (!false)))
                                                            }
                                                        }
                                                        _ = fi109
                                                        _ = v110
                                                    }
                                                    switch v11 {
                                                    case "hello":
                                                        var v114 string = "foo bar"
                                                        _ = v114
                                                    case "\n\t":
                                                        if len(v93) >= 2 {
                                                            v93 = v93[1:len(v93)]
                                                        }
                                                    case "foo bar":
                                                        var v115 float32 = (3.14159 / 3.14159)
                                                        _ = v115
                                                    }
                                                }
                                                v116 := GenS1{F0: 1, F1: (-(v12))}
                                                _ = v116
                                            } else {
                                                v117 := GenS12{F0: 18446744073709551615}
                                                for fi118 := 0; fi118 < 1; fi118++ {
                                                    delete(v3, (fi92 + fi118))
                                                    if v2 {
                                                        var v119 string = v11
                                                        v120 := GenS7{F0: 0, F1: 1}
                                                        v2 = true
                                                        _ = v119
                                                        _ = v120
                                                    } else {
                                                        var v121 uint64 = 2
                                                        if (0 > -819) {
                                                            v11 = "a"
                                                        } else {
                                                            switch v11 {
                                                            case "\n\t":
                                                                if true {
                                                                    var v122 int64 = (v12 + (v12 - int64(fi92)))
                                                                    fmt.Println(v122, v5)
                                                                    v122 = (-(v122))
                                                                    _ = v122
                                                                } else {
                                                                    if ("hello" != v11) {
                                                                        v123 := GenS15{F0: -20, F1: (fi92 ^ fi118), F2: v121}
                                                                        v12 = (-(v12))
                                                                        _ = v123
                                                                    } else {
                                                                        var v124 uint = 4294967295
                                                                        _ = len(v4)
                                                                        v12 += int64(fi118)
                                                                        if (fi92 >= fi118) {
                                                                            switch 32767 {
                                                                            case 42:
                                                                                if ((v5 / ((v5) | 1)) == -1) {
                                                                                    if (false && (752 < 1)) {
                                                                                        if (1e10 >= (v1 / 3.14159)) {
                                                                                            v11 = v11
                                                                                        }
                                                                                    }
                                                                                    switch fi118 {
                                                                                    case -9223372036854775808:
                                                                                        delete(v4, 2147483647)
                                                                                    case 0:
                                                                                        fmt.Println(v99, v5, v124)
                                                                                    case 32768:
                                                                                        var v125 uint8 = 144
                                                                                        _ = v125
                                                                                    default:
                                                                                        switch fi92 {
                                                                                        case 42:
                                                                                            v126 := GenS16{F0: (^(v5)), F1: (^(v121))}
                                                                                            _ = v126
                                                                                        case -2147483648:
                                                                                            var v127 int16 = 187
                                                                                            _ = v127
                                                                                        default:
                                                                                            var v128 uint32 = 2147483647
                                                                                            _ = v128
                                                                                        }
                                                                                    }
                                                                                }
                                                                            case 32768:
                                                                                if false {
                                                                                    switch -1 {
                                                                                    case 261:
                                                                                        var v129 int32 = 42
                                                                                        _ = v129
                                                                                    case 42:
                                                                                        v93 = append(v93, 4294967295)
                                                                                    case 171:
                                                                                        if v2 {
                                                                                            var v130 float64 = 0.0
                                                                                            v131 := v3[fi92]
                                                                                            _ = v130
                                                                                            _ = v131
                                                                                        }
                                                                                    }
                                                                                    v5 = v5
                                                                                }
                                                                            case 32767:
                                                                                v132 := GenS17{F0: 2147483648, F1: -32768, F2: -128}
                                                                                _ = v132
                                                                            }
                                                                            v133 := v4[0]
                                                                            _ = v133
                                                                        }
                                                                        var v134 uint8 = 127
                                                                        _ = v124
                                                                        _ = v134
                                                                    }
                                                                    if (1 > 0) {
                                                                        if (913 > (v12 & -9223372036854775808)) {
                                                                            if v2 {
                                                                                v121 += v121
                                                                                v135 := GenS6{F0: (fi92 + 32768), F1: 1, F2: 128}
                                                                                _ = len(v93)
                                                                                _ = v135
                                                                            } else {
                                                                                if v2 {
                                                                                    if v2 {
                                                                                        switch v11 {
                                                                                        case "a":
                                                                                            _ = len(v93)
                                                                                        case "":
                                                                                            if false {
                                                                                                v12 = v12
                                                                                            }
                                                                                            v5 += 1000
                                                                                            var v136 uint16 = 255
                                                                                            _ = v136
                                                                                        case "hello":
                                                                                            _ = len(v93)
                                                                                        default:
                                                                                            if v2 {
                                                                                                v93 = append(v93, (fi118 * -499))
                                                                                                var v137 int8 = 100
                                                                                                var v138 int16 = 0
                                                                                                _ = v137
                                                                                                _ = v138
                                                                                            } else {
                                                                                                delete(v4, 2147483648)
                                                                                            }
                                                                                        }
                                                                                        v139 := v4[2147483647]
                                                                                        v140 := GenS11{F0: 1, F1: v12}
                                                                                        fmt.Println(fi118)
                                                                                        _ = v139
                                                                                        _ = v140
                                                                                    } else {
                                                                                        v141 := GenS15{F0: 84, F1: 42, F2: (v121 / (((^(v121))) | 1))}
                                                                                        var v142 string = v11
                                                                                        _ = v141
                                                                                        _ = v142
                                                                                    }
                                                                                    v2 = (!true)
                                                                                } else {
                                                                                    fmt.Println(v5, v12, v1)
                                                                                }
                                                                                var v143 float64 = v99
                                                                                _ = v143
                                                                            }
                                                                        }
                                                                        fmt.Println(v1, v12, fi118)
                                                                        _ = len(v3)
                                                                    }
                                                                    v12 -= v12
                                                                    switch fi92 {
                                                                    case 9223372036854775807:
                                                                        if (128 >= 127) {
                                                                            v2 = ((v5 / ((int16(v12)) | 1)) < v5)
                                                                        } else {
                                                                            switch v11 {
                                                                            case "":
                                                                                if false {
                                                                                    switch v11 {
                                                                                    case "":
                                                                                        v5 -= (v5 | v5)
                                                                                    }
                                                                                } else {
                                                                                    v121 += 816
                                                                                    fmt.Println(v121)
                                                                                    if (!v2) {
                                                                                        v2 = v2
                                                                                        switch "\n\t" {
                                                                                        case "a":
                                                                                            fmt.Println(fi118)
                                                                                        case "\n\t":
                                                                                            switch fi92 {
                                                                                            case -2147483648:
                                                                                                if v2 {
                                                                                                    var v144 bool = v2
                                                                                                    _ = v144
                                                                                                }
                                                                                            }
                                                                                        case "":
                                                                                            var v145 int = int(v121)
                                                                                            _ = v145
                                                                                        }
                                                                                        v12 = v12
                                                                                    } else {
                                                                                        var v146 bool = (v2 || (v1 >= (1e10 / -2.25)))
                                                                                        var v147 uint32 = 2
                                                                                        _ = v146
                                                                                        _ = v147
                                                                                    }
                                                                                    if len(v93) >= 2 {
                                                                                        v93 = v93[1:len(v93)]
                                                                                    }
                                                                                }
                                                                            case "\n\t":
                                                                                var v148 float64 = ((1e10 * float64(3.14159)) - (v99 + v99))
                                                                                _ = v148
                                                                            default:
                                                                                v149 := GenS18{F0: (v5 | (v5 ^ (v5 / (((v5 / (((v5 << ((967) & 63))) | 1))) | 1)))), F1: -2147483648, F2: 2}
                                                                                _ = v149
                                                                            }
                                                                            if (v2 && (198 <= 2)) {
                                                                                v2 = (fi118 >= (fi92 >> ((339) & 63)))
                                                                                v150 := GenS5{F0: 255}
                                                                                _ = v150
                                                                            }
                                                                            switch 484 {
                                                                            case -1:
                                                                                v2 = (!(1 < 867))
                                                                            default:
                                                                                var v151 int32 = 42
                                                                                _ = v151
                                                                            }
                                                                        }
                                                                    case 128:
                                                                        var v152 float64 = v1
                                                                        _ = v152
                                                                    case 32767:
                                                                        v121 -= (v121 % (((v121 / (((v121 << ((18446744073709551615) & 63))) | 1))) | 1))
                                                                    default:
                                                                        v11 = (v11 + v11)
                                                                    }
                                                                    v153 := GenS10{F0: v5, F1: 2147483647, F2: 2}
                                                                    v12 = (v12 << ((4294967295) & 63))
                                                                    var v154 float32 = ((1.5 / 1.5) * 1.5)
                                                                    v155 := v3[fi92]
                                                                    _ = v153
                                                                    _ = v154
                                                                    _ = v155
                                                                }
                                                            case "":
                                                                var v156 uint16 = 42
                                                                _ = v156
                                                            default:
                                                                v157 := v3[int(v12)]
                                                                _ = v157
                                                            }
                                                            if (!(!(!v2))) {
                                                                if (v5 >= -1) {
                                                                    v158 := GenS12{F0: (v121 % ((v121) | 1))}
                                                                    var v159 int16 = v5
                                                                    var v160 string = (v11 + ("hello" + "a"))
                                                                    _ = v158
                                                                    _ = v159
                                                                    _ = v160
                                                                } else {
                                                                    v2 = v2
                                                                    var v161 float64 = float64(1e10)
                                                                    if len(v93) >= 2 {
                                                                        v93 = v93[1:len(v93)]
                                                                    }
                                                                    _ = v161
                                                                }
                                                                var v162 float64 = 1e10
                                                                v3[32768] = true
                                                                _ = v162
                                                            }
                                                            var v163 uint32 = 792
                                                            _ = v163
                                                        }
                                                        _ = v121
                                                    }
                                                    v5 -= 0
                                                    v1 += v1
                                                    _ = len(v3)
                                                    _ = fi118
                                                }
                                                switch v11 {
                                                case "\n\t":
                                                    v99 = float64(((-2.25 / 1.5) * 0.1))
                                                }
                                                _ = v117
                                            }
                                            v99 = v1
                                        } else {
                                            var v164 string = ""
                                            v164 = "foo bar"
                                            _ = v164
                                        }
                                    }
                                default:
                                    var v165 int16 = v5
                                    _ = v165
                                }
                                _ = v99
                            } else {
                                if v2 {
                                    var v166 bool = (v11 < ((v11 + ("" + ("" + ""))) + ("foo bar" + "")))
                                    fmt.Println(v2)
                                    _ = v166
                                }
                                switch ("hello" + "a") {
                                case "a":
                                    delete(v3, fi92)
                                default:
                                    var v167 bool = true
                                    _ = v167
                                }
                            }
                        default:
                            var v168 uint32 = 4294967295
                            _ = v168
                        }
                        v11 = "\n\t"
                        if (0 == -1) {
                            if (47 == -3) {
                                var v169 bool = true
                                _ = v169
                            } else {
                                switch ("" + ("foo bar" + v11)) {
                                case "":
                                    v2 = ((v12 & (v12 * v12)) <= (v12 - (^(v12))))
                                case "\n\t":
                                    var v170 uint32 = 601
                                    _ = v170
                                default:
                                    v171 := GenS19{F0: 710, F1: 924, F2: 23}
                                    _ = v171
                                }
                                if ((32767 != 32767) && v2) {
                                    if true {
                                        if ((!v2) || (2 <= 18446744073709551615)) {
                                            var v172 uint16 = 1
                                            _ = v172
                                        } else {
                                            if (18446744073709551615 < 302) {
                                                switch v11 {
                                                case "\n\t":
                                                    fmt.Println(v1)
                                                case "hello":
                                                    if true {
                                                        var v173 float32 = ((0.0 * (((1.5 * -2.25) - 0.1) / (((0.1 - float32(1e10)) * 1e10) / (((-2.25 / 1e10) / 1.0) * (1.5 / (1.5 + 0.1)))))) / 1.0)
                                                        v5 = v5
                                                        v174 := GenS20{F0: fi92}
                                                        _ = v173
                                                        _ = v174
                                                    }
                                                }
                                                var v175 float64 = (v1 / ((0.1 / v1) + ((0.1 + v1) + (float64(-2.25) + (v1 - 0.0)))))
                                                switch (fi92 | (fi92 * 9223372036854775807)) {
                                                case 128:
                                                    for fi176 := 0; fi176 < 5; fi176++ {
                                                        if v2 {
                                                            switch ("foo bar" + v11) {
                                                            case "\n\t":
                                                                var v177 uint8 = 42
                                                                _ = v177
                                                            case "foo bar":
                                                                v178 := GenS15{F0: 127, F1: 447, F2: 129}
                                                                _ = v178
                                                            default:
                                                                v1 -= -2.25
                                                                v179 := GenS2{F0: 0, F1: (("\n\t" + v11) + (("hello" + v11) + ("hello" + v11))), F2: 2}
                                                                var v180 int16 = (v5 & v5)
                                                                v12 = v12
                                                                if ((433 > 331) && true) {
                                                                    fmt.Println(v12)
                                                                    switch v11 {
                                                                    case "":
                                                                        v181 := GenS1{F0: -547, F1: v12}
                                                                        _ = v181
                                                                    case "a":
                                                                        v12 += (-(v12))
                                                                    case "\n\t":
                                                                        var v182 bool = v2
                                                                        switch v11 {
                                                                        case "foo bar":
                                                                            v12 -= v12
                                                                        case "":
                                                                            var v183 uint8 = 102
                                                                            _ = v183
                                                                        case "\n\t":
                                                                            var v184 bool = v182
                                                                            _ = v184
                                                                        default:
                                                                            var v185 float32 = 0.1
                                                                            _ = v185
                                                                        }
                                                                        _ = v182
                                                                    default:
                                                                        v186 := GenS11{F0: 0, F1: v12}
                                                                        _ = v186
                                                                    }
                                                                    if (-29 >= 100) {
                                                                        if v2 {
                                                                            v93 = append(v93, (fi176 << ((4294967295) & 63)))
                                                                            v175 = v1
                                                                        }
                                                                    } else {
                                                                        v187 := v4[65535]
                                                                        v188 := GenS18{F0: v180, F1: -26, F2: 655}
                                                                        _ = v187
                                                                        _ = v188
                                                                    }
                                                                }
                                                                v5 = 0
                                                                _ = v179
                                                                _ = v180
                                                            }
                                                            v189 := GenS8{F0: 139, F1: 0, F2: (^(fi92))}
                                                            switch "hello" {
                                                            case "a":
                                                                var v190 float32 = ((3.14159 - (-2.25 * 1e10)) * 3.14159)
                                                                _ = v190
                                                            case "hello":
                                                                v191 := GenS11{F0: 32767, F1: v12}
                                                                _ = v191
                                                            case "foo bar":
                                                                var v192 uint64 = 42
                                                                _ = v192
                                                            default:
                                                                v2 = v2
                                                            }
                                                            _ = v189
                                                        }
                                                        _ = fi176
                                                    }
                                                case -702:
                                                    v1 -= float64((-2.25 - -2.25))
                                                default:
                                                    var v193 int8 = -1
                                                    _ = v193
                                                }
                                                _ = v175
                                            }
                                            if (!(!false)) {
                                                v1 += v1
                                                v194 := GenS6{F0: (fi92 / (((fi92 % ((663) | 1))) | 1)), F1: 1, F2: 42}
                                                _ = v194
                                            }
                                        }
                                        var v195 uint64 = 106
                                        _ = v195
                                    }
                                    var v196 float32 = 3.14159
                                    _ = v196
                                } else {
                                    v2 = (-1 == 127)
                                    fmt.Println(v2, v5)
                                }
                                v197 := GenS9{F0: 42, F1: (("a" + "\n\t") + v11)}
                                _ = v197
                            }
                            if v2 {
                                if ((!v2) || false) {
                                    fmt.Println(v12, v12)
                                    if true {
                                        fmt.Println(v1, fi92, v11)
                                    } else {
                                        var v198 int = 471
                                        var v199 int64 = v12
                                        fmt.Println(v11, v5)
                                        _ = v198
                                        _ = v199
                                    }
                                    v12 -= v12
                                } else {
                                    var v200 float64 = 1e10
                                    _ = v200
                                }
                                if v2 {
                                    fmt.Println(fi92, v1)
                                } else {
                                    var v201 int8 = -105
                                    var v202 int16 = v5
                                    v203 := GenS9{F0: -2147483648, F1: v11}
                                    _ = v201
                                    _ = v202
                                    _ = v203
                                }
                            }
                        }
                    }
                    _ = fi92
                    _ = v93
                }
                for fi204 := 0; fi204 < 13; fi204++ {
                    v2 = true
                    if ((v11 + "hello") > "a") {
                        for fi205 := 0; fi205 < 6; fi205++ {
                            v1 = (float64((((1e10 * 0.1) * 1.5) / 1.5)) * v1)
                            v206 := GenS8{F0: 0, F1: 86, F2: (fi205 & 127)}
                            var v207 uint32 = 1
                            _ = fi205
                            _ = v206
                            _ = v207
                        }
                        v5 -= v5
                    }
                    _ = fi204
                }
                v208 := GenS21{F0: -50}
                _ = v208
            }
        }
    }
    v209 := []int{-372, 32767}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v209:", len(v209), v209)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v209
    _ = v3
    _ = v4
    _ = v5
}
