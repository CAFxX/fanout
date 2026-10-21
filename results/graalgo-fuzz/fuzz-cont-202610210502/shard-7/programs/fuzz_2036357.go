package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint; F2 int16 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint; F1 int; F2 int }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 string }
type GenS6 struct { F0 int; F1 int32 }
type GenS7 struct { F0 int; F1 int8; F2 int16 }
type GenS8 struct { F0 int32 }
type GenS9 struct { F0 uint32; F1 int32 }
type GenS10 struct { F0 uint16; F1 uint8; F2 int32 }
type GenS11 struct { F0 uint }
type GenS12 struct { F0 int32; F1 int64; F2 int8 }
type GenS13 struct { F0 int16; F1 string; F2 int8 }
type GenS14 struct { F0 uint64; F1 int64 }
type GenS15 struct { F0 int; F1 int32; F2 uint }
type GenS16 struct { F0 string; F1 uint; F2 int8 }
type GenS17 struct { F0 int; F1 int8; F2 uint8 }
type GenS18 struct { F0 bool }

func main() {
    var v1 string = (("foo bar" + ("hello" + ("\n\t" + "foo bar"))) + "a")
    var v2 uint64 = 18446744073709551615
    var v3 int8 = 59
    var v4 int = 128
    var v5 int = (v4 | (v4 >> ((0) & 63)))
    if ((("" < v1) || (0 <= 2147483647)) && true) {
        var v6 uint = 255
        _ = v6
    } else {
        if false {
            v7 := GenS1{F0: (-(v3)), F1: 6, F2: -32768}
            _ = v7
        } else {
            v1 = v1
            v8 := []int{}
            var v9 uint64 = (^(v2))
            _ = v8
            _ = v9
        }
        if true {
            var v10 int8 = v3
            _ = v10
        } else {
            fmt.Println(v5)
            v2 -= (v2 << ((18446744073709551615) & 63))
            var v11 int8 = v3
            _ = v11
        }
        v12 := []int{(v5 - v5), (v5 / (((-(v4))) | 1)), v4, v5}
        _ = v12
    }
    v3 = -128
    v4 = v5
    var v13 uint8 = 128
    fmt.Println(v13, v2)
    var v14 float32 = 1e10
    if (!(284 > -1)) {
        if (32767 > 968) {
            for fi15 := 0; fi15 < 17; fi15++ {
                v16 := []int{(-(v4))}
                if (-2147483648 <= -1) {
                    if len(v16) > 0 {
                        v16[0] = v4
                    }
                } else {
                    if true {
                        fmt.Println(v4, v13)
                    }
                    v2 = (v2 / ((v2) | 1))
                    for fi17 := 0; fi17 < 6; fi17++ {
                        fmt.Println(v4, v13, v2)
                        v18 := GenS2{F0: 2}
                        _ = fi17
                        _ = v18
                    }
                }
                for fi19 := 0; fi19 < 7; fi19++ {
                    v20 := GenS1{F0: (v3 % (((^(v3))) | 1)), F1: 670, F2: -1}
                    var v21 int8 = (v3 / ((v3) | 1))
                    var v22 bool = (v14 != float32(float64((v14 - v14))))
                    var v23 uint8 = 1
                    var v24 float64 = 1e10
                    switch v1 {
                    case "hello":
                        var v25 float32 = v14
                        _ = v25
                    }
                    _ = fi19
                    _ = v20
                    _ = v21
                    _ = v22
                    _ = v23
                    _ = v24
                }
                _ = fi15
                _ = v16
            }
            v4 -= v5
        }
    }
    v5 = v4
    var v26 string = "a"
    if (v13 > v13) {
        v27 := GenS3{F0: 65535, F1: v4, F2: (v4 + -299)}
        for fi28 := 0; fi28 < 18; fi28++ {
            var v29 int64 = -796
            if true {
                for fi30 := 0; fi30 < 5; fi30++ {
                    v3 = 0
                    v31 := GenS1{F0: -1, F1: 4294967295, F2: 728}
                    v3 -= (v3 * v3)
                    if true {
                        v2 = v2
                        var v32 int16 = 1000
                        _ = v32
                    } else {
                        var v33 uint16 = 255
                        if (v14 < v14) {
                            v34 := GenS2{F0: uint16(v13)}
                            if false {
                                v35 := []int{-1, 32767}
                                fmt.Println(v1, v2, fi30)
                                var v36 uint16 = v33
                                _ = v35
                                _ = v36
                            } else {
                                v37 := GenS4{F0: (v3 <= v3)}
                                v1 = ""
                                _ = v37
                            }
                            v33 = (v33 / (((v33 - (v33 >> ((65535) & 63)))) | 1))
                            var v38 uint32 = uint32(v2)
                            fmt.Println(v1, v14)
                            v39 := GenS5{F0: v1}
                            switch (v5 ^ fi28) {
                            case 0:
                                v5 += v4
                            default:
                                if (false && true) {
                                    if (true || false) {
                                        v40 := GenS6{F0: v5, F1: 42}
                                        _ = v40
                                    } else {
                                        v13 = (v13 / (((v13 + uint8(v2))) | 1))
                                    }
                                    switch fi30 {
                                    case 9223372036854775807:
                                        v38 -= 2147483647
                                    default:
                                        v41 := GenS7{F0: (fi30 % ((32767) | 1)), F1: int8(v2), F2: int16(v29)}
                                        _ = v41
                                    }
                                }
                            }
                            fmt.Println(v26)
                            if true {
                                var v42 int32 = 92
                                v2 = v2
                                v43 := []int{}
                                _ = v42
                                _ = v43
                            }
                            if (-622 >= 0) {
                                v44 := []int{v5, 535, v4}
                                _ = v44
                            } else {
                                v45 := GenS5{F0: v26}
                                _ = v45
                            }
                            _ = v34
                            _ = v38
                            _ = v39
                        }
                        if true {
                            v46 := GenS8{F0: -3}
                            fmt.Println(v3)
                            v47 := []int{(v4 | (v4 + v5)), (v5 ^ v4), 0, 42}
                            _ = v46
                            _ = v47
                        }
                        if (!true) {
                            if false {
                                var v48 int64 = -280
                                _ = v48
                            }
                        }
                        _ = v33
                    }
                    _ = fi30
                    _ = v31
                }
            } else {
                var v49 int = (v4 >> ((18446744073709551615) & 63))
                v50 := []int{-2147483648, v4}
                for fi51 := 0; fi51 < 5; fi51++ {
                    if len(v50) >= 2 {
                        v50 = v50[1:len(v50)]
                    }
                    v52 := GenS8{F0: 0}
                    _ = fi51
                    _ = v52
                }
                _ = v49
                _ = v50
            }
            v53 := GenS5{F0: ("foo bar" + ("a" + ""))}
            _ = fi28
            _ = v29
            _ = v53
        }
        _ = v27
    } else {
        switch v1 {
        case "foo bar":
            if false {
                v54 := GenS9{F0: 878, F1: -2147483648}
                for fi55 := 0; fi55 < 12; fi55++ {
                    v4 += (v4 << ((4294967295) & 63))
                    _ = fi55
                }
                for fi56 := 0; fi56 < 12; fi56++ {
                    switch v1 {
                    case "a":
                        if false {
                            fmt.Println(v4, v1, v1)
                        }
                    case "hello":
                        var v57 string = "a"
                        _ = v57
                    }
                    if ((v2 & v2) >= 18446744073709551615) {
                        v13 = (v13 * v13)
                        switch v26 {
                        case "":
                            v58 := GenS9{F0: 2147483647, F1: -155}
                            _ = v58
                        case "hello":
                            v59 := []int{v4, (fi56 << ((65535) & 63)), (v5 / ((-2147483648) | 1)), 42}
                            _ = v59
                        }
                        v60 := GenS10{F0: 865, F1: (v13 >> ((255) & 63)), F2: 871}
                        _ = v60
                    } else {
                        v61 := []int{v5, (fi56 | v4), (v4 + int(v2))}
                        for fi62 := 0; fi62 < 7; fi62++ {
                            v63 := GenS2{F0: 367}
                            var v64 uint16 = 919
                            v1 = v1
                            _ = fi62
                            _ = v63
                            _ = v64
                        }
                        _ = v61
                    }
                    v65 := []int{v4}
                    _ = fi56
                    _ = v65
                }
                _ = v54
            } else {
                var v66 int64 = -1
                v4 = (v5 ^ -486)
                if true {
                    switch v5 {
                    case -9223372036854775808:
                        v1 = (v1 + v26)
                    case 0:
                        fmt.Println(v1)
                    }
                    var v67 string = ((v1 + "a") + "hello")
                    _ = v67
                } else {
                    switch "" {
                    case "":
                        v68 := GenS6{F0: 1, F1: 1}
                        _ = v68
                    case "\n\t":
                        fmt.Println(v26, v4, v1)
                    }
                    v13 += (v13 ^ v13)
                }
                _ = v66
            }
        default:
            var v69 int16 = -32768
            _ = v69
        }
        var v70 string = "hello"
        var v71 bool = false
        _ = v70
        _ = v71
    }
    v72 := GenS4{F0: true}
    switch "a" {
    case "foo bar":
        var v73 int16 = 42
        _ = v73
    case "hello":
        fmt.Println(v14)
    case "a":
        fmt.Println(v3, v26, v4)
    default:
        switch v4 {
        case 2147483647:
            if true {
                v26 = v26
                v3 += 100
            }
        default:
            v74 := []int{v5, (^(v5)), (v4 % ((v4) | 1))}
            _ = v74
        }
    }
    var v75 uint = 65535
    if ((uint16(v75) < 65535) || (("foo bar" + (v1 + "\n\t")) < "\n\t")) {
        if ((v75 << ((uint(v13)) & 63)) >= 412) {
            v1 = "a"
        } else {
            switch (v4 - 4294967295) {
            case -371:
                fmt.Println(v26)
            case 32768:
                v76 := GenS9{F0: 2147483647, F1: -1}
                _ = v76
            default:
                switch v4 {
                case 42:
                    v77 := []int{(v4 - (v5 - (v5 ^ (v4 * 128)))), (v4 ^ v4), (^(v4))}
                    _ = v77
                case 445:
                    switch (v26 + (v1 + ((v1 + ("\n\t" + "a")) + ((v26 + ("\n\t" + (v26 + "a"))) + v1)))) {
                    case "\n\t":
                        v14 = (((v14 + (1.5 - v14)) - v14) * 0.1)
                    case "hello":
                        v1 = "foo bar"
                    case "a":
                        if (false || true) {
                            v14 = 1e10
                            if true {
                                for fi78 := 0; fi78 < 12; fi78++ {
                                    v14 = v14
                                    v4 = v4
                                    if false {
                                        if (!true) {
                                            v75 = (v75 << ((81) & 63))
                                            var v79 int = 128
                                            v80 := GenS2{F0: 1}
                                            _ = v79
                                            _ = v80
                                        } else {
                                            v75 = (^(v75))
                                        }
                                    } else {
                                        switch 9223372036854775807 {
                                        case 399:
                                            for fi81 := 0; fi81 < 14; fi81++ {
                                                v4 -= fi78
                                                v2 = 4294967295
                                                _ = fi81
                                            }
                                        case 4294967295:
                                            for fi82 := 0; fi82 < 9; fi82++ {
                                                v26 = v1
                                                switch -506 {
                                                case 32768:
                                                    switch fi82 {
                                                    case 32767:
                                                        var v83 int16 = -1
                                                        _ = v83
                                                    case 155:
                                                        switch v5 {
                                                        case 128:
                                                            v84 := []int{v5, 32767, v5, (v5 << (((v75 - (v75 - 2))) & 63))}
                                                            _ = v84
                                                        case -940:
                                                            var v85 uint64 = v2
                                                            _ = v85
                                                        case 696:
                                                            v26 = "foo bar"
                                                        }
                                                    case -2147483648:
                                                        fmt.Println(v2)
                                                    default:
                                                        if (65535 <= 2) {
                                                            v5 += (v5 % ((v5) | 1))
                                                            v3 -= -128
                                                        } else {
                                                            if (!(v13 != v13)) {
                                                                if (v14 == -2.25) {
                                                                    v86 := GenS11{F0: 255}
                                                                    switch v5 {
                                                                    case 9223372036854775807:
                                                                        switch 940 {
                                                                        case 1:
                                                                            v87 := []int{}
                                                                            _ = v87
                                                                        case 4294967295:
                                                                            v5 -= int(v75)
                                                                        case -2147483648:
                                                                            switch ("" + "") {
                                                                            case "":
                                                                                var v88 uint = (v75 * (v75 % (((v75 >> (((^(v75))) & 63))) | 1)))
                                                                                _ = v88
                                                                            case "a":
                                                                                switch v26 {
                                                                                case "hello":
                                                                                    var v89 uint = (v75 & v75)
                                                                                    _ = v89
                                                                                case "foo bar":
                                                                                    switch v5 {
                                                                                    case -783:
                                                                                        fmt.Println(v1, v1)
                                                                                    case -746:
                                                                                        var v90 uint32 = 718
                                                                                        _ = v90
                                                                                    case 9223372036854775807:
                                                                                        if (!true) {
                                                                                            fmt.Println(fi82, fi82)
                                                                                            if ((-(v3)) <= v3) {
                                                                                                v4 -= fi78
                                                                                                fmt.Println(fi82, v13)
                                                                                            } else {
                                                                                                switch v26 {
                                                                                                case "a":
                                                                                                    if false {
                                                                                                        var v91 int8 = int8(v13)
                                                                                                        _ = v91
                                                                                                    } else {
                                                                                                        fmt.Println(v3)
                                                                                                    }
                                                                                                case "\n\t":
                                                                                                    v92 := []int{}
                                                                                                    _ = v92
                                                                                                }
                                                                                                if true {
                                                                                                    var v93 int16 = 1000
                                                                                                    v94 := []int{1, -794}
                                                                                                    var v95 uint64 = (v2 >> ((18446744073709551615) & 63))
                                                                                                    _ = v93
                                                                                                    _ = v94
                                                                                                    _ = v95
                                                                                                }
                                                                                            }
                                                                                            v96 := GenS12{F0: -177, F1: 666, F2: v3}
                                                                                            v75 = 42
                                                                                            _ = v96
                                                                                        }
                                                                                    default:
                                                                                        if false {
                                                                                            v97 := GenS13{F0: 32767, F1: (v26 + (v26 + ("a" + "\n\t"))), F2: 127}
                                                                                            switch v1 {
                                                                                            case "\n\t":
                                                                                                v2 -= 0
                                                                                            case "":
                                                                                                v26 = ("a" + "")
                                                                                            }
                                                                                            _ = v97
                                                                                        }
                                                                                    }
                                                                                default:
                                                                                    v98 := GenS6{F0: (fi78 << ((18446744073709551615) & 63)), F1: -1}
                                                                                    _ = v98
                                                                                }
                                                                            case "hello":
                                                                                if false {
                                                                                    v2 -= (v2 / ((v2) | 1))
                                                                                }
                                                                            default:
                                                                                v3 = 0
                                                                            }
                                                                        default:
                                                                            switch (v26 + v26) {
                                                                            case "\n\t":
                                                                                v2 += (v2 % (((v2 << (((v75 | 65535)) & 63))) | 1))
                                                                            default:
                                                                                v26 = v1
                                                                                v99 := []int{(-(v5)), (-(fi82)), v4, (fi82 & v5)}
                                                                                _ = v99
                                                                            }
                                                                        }
                                                                    case -918:
                                                                        var v100 uint16 = 42
                                                                        _ = v100
                                                                    case -915:
                                                                        v101 := []int{(fi78 | -1)}
                                                                        _ = v101
                                                                    }
                                                                    _ = v86
                                                                }
                                                                var v102 int64 = -1
                                                                _ = v102
                                                            } else {
                                                                v103 := GenS14{F0: (v2 % ((v2) | 1)), F1: 67}
                                                                fmt.Println(v13, v3, v5)
                                                                fmt.Println(fi82, v4, v5)
                                                                _ = v103
                                                            }
                                                            var v104 int8 = (v3 - -128)
                                                            v105 := GenS15{F0: v4, F1: 42, F2: (^(v75))}
                                                            switch (v1 + (v1 + (v26 + ((((v26 + "foo bar") + (v1 + v26)) + v1) + ((v26 + v26) + "hello"))))) {
                                                            case "\n\t":
                                                                switch (v4 >> ((65535) & 63)) {
                                                                case -2147483648:
                                                                    fmt.Println(v75)
                                                                case 4294967295:
                                                                    var v106 int = (v4 / ((42) | 1))
                                                                    _ = v106
                                                                }
                                                            case "a":
                                                                v107 := GenS16{F0: v1, F1: v75, F2: 127}
                                                                _ = v107
                                                            case "":
                                                                v5 = -847
                                                            }
                                                            _ = v104
                                                            _ = v105
                                                        }
                                                    }
                                                case 4294967295:
                                                    v108 := GenS12{F0: -342, F1: int64(v75), F2: -83}
                                                    _ = v108
                                                case 829:
                                                    var v109 uint16 = 353
                                                    v109 += (^(v109))
                                                    v110 := []int{(v4 - fi78), (^(v4)), (v5 / ((int(v13)) | 1)), 127}
                                                    v111 := GenS1{F0: 0, F1: 1, F2: -1}
                                                    _ = v109
                                                    _ = v110
                                                    _ = v111
                                                default:
                                                    var v112 int8 = (v3 ^ v3)
                                                    _ = v112
                                                }
                                                _ = fi82
                                            }
                                        }
                                        v1 = "a"
                                    }
                                    _ = fi78
                                }
                            } else {
                                v26 = v26
                                v1 = "\n\t"
                            }
                            v4 = (v4 % ((int(v75)) | 1))
                        } else {
                            fmt.Println(v5, v26, v4)
                            v2 = v2
                            var v113 int32 = -685
                            _ = v113
                        }
                    default:
                        fmt.Println(v26, v26, v1)
                    }
                }
            }
            v114 := GenS5{F0: v1}
            fmt.Println(v26, v3)
            _ = v114
        }
        v26 = v1
    }
    for fi115 := 0; fi115 < 10; fi115++ {
        var v116 int8 = v3
        for fi117 := 0; fi117 < 14; fi117++ {
            if true {
                switch "" {
                case "":
                    if (0 >= -1) {
                        if (42 == 255) {
                            v118 := GenS6{F0: v5, F1: 2147483647}
                            if (v14 == v14) {
                                fmt.Println(v75, v116)
                            }
                            _ = v118
                        } else {
                            var v119 int64 = 988
                            v26 = v1
                            _ = v119
                        }
                        switch v26 {
                        case "\n\t":
                            fmt.Println(v5)
                        case "a":
                            if false {
                                v120 := []int{(v4 * -1), v5, v5, (v4 >> (((v75 * (v75 & 255))) & 63))}
                                fmt.Println(fi117, fi115)
                                _ = v120
                            } else {
                                var v121 string = (((("\n\t" + "a") + "\n\t") + (("foo bar" + v1) + "\n\t")) + v26)
                                v116 -= (-(v3))
                                _ = v121
                            }
                            v13 = uint8(v75)
                        }
                    }
                default:
                    if true {
                        var v122 int8 = v3
                        _ = v122
                    } else {
                        v123 := []int{v5}
                        if true {
                            var v124 int32 = -1
                            _ = v124
                        }
                        if (!(v13 == v13)) {
                            var v125 int = (v4 | (fi115 - (fi115 % ((42) | 1))))
                            v5 = v125
                            _ = v125
                        } else {
                            if len(v123) > 0 {
                                v123[0] = -29
                            }
                            _ = len(v123)
                            v14 += v14
                        }
                        var v126 float64 = 0.0
                        _ = v123
                        _ = v126
                    }
                    var v127 int64 = int64(v75)
                    if true {
                        v128 := []int{v4, fi115, (fi115 ^ (-(v5)))}
                        _ = len(v128)
                        var v129 string = v26
                        switch ("\n\t" + v1) {
                        case "foo bar":
                            if true {
                                if len(v128) >= 2 {
                                    v128 = v128[1:len(v128)]
                                }
                            } else {
                                if (0 >= 18446744073709551615) {
                                    v128 = append(v128, (v5 * (v4 * fi117)))
                                    v128 = append(v128, v4)
                                    var v130 int8 = v116
                                    var v131 uint32 = 4294967295
                                    if (!true) {
                                        fmt.Println(v5)
                                        fmt.Println(v75, v2, v75)
                                        if false {
                                            var v132 int16 = 1000
                                            var v133 float64 = (float64(float32(0.1)) / 1e10)
                                            _ = v132
                                            _ = v133
                                        }
                                    }
                                    if true {
                                        if (true && (false || ("a" >= (v129 + v129)))) {
                                            v128 = append(v128, int(v75))
                                            v128 = append(v128, v5)
                                            v4 -= fi117
                                        } else {
                                            var v134 uint16 = uint16(v131)
                                            switch fi117 {
                                            case 32768:
                                                var v135 string = v1
                                                _ = v135
                                            case 1:
                                                if len(v128) > 0 {
                                                    v128[0] = (fi117 << ((v75) & 63))
                                                }
                                            case -9223372036854775808:
                                                _ = len(v128)
                                            }
                                            v128 = append(v128, (v5 & -434))
                                            _ = v134
                                        }
                                        v14 = 1e10
                                    } else {
                                        var v136 int8 = 42
                                        fmt.Println(v14, v13)
                                        if true {
                                            switch (fi117 ^ (fi117 ^ v5)) {
                                            case 4294967295:
                                                var v137 uint32 = (v131 ^ (v131 + v131))
                                                _ = v137
                                            case 0:
                                                var v138 uint = (v75 << ((2) & 63))
                                                _ = v138
                                            case -2147483648:
                                                v136 = v116
                                            }
                                            var v139 int8 = -1
                                            var v140 int16 = -957
                                            _ = v139
                                            _ = v140
                                        }
                                        _ = v136
                                    }
                                    _ = v130
                                    _ = v131
                                } else {
                                    if len(v128) >= 2 {
                                        v128 = v128[1:len(v128)]
                                    }
                                    if (-2147483648 < 2147483647) {
                                        switch v4 {
                                        case 187:
                                            v26 = "foo bar"
                                        case 128:
                                            switch (v26 + (((v26 + "a") + v26) + v129)) {
                                            case "\n\t":
                                                var v141 string = "a"
                                                _ = v141
                                            }
                                        default:
                                            v1 = v26
                                        }
                                        var v142 int16 = -32768
                                        _ = v142
                                    }
                                    v5 += v5
                                }
                                v128 = append(v128, fi115)
                                v128 = append(v128, fi117)
                            }
                        case "a":
                            switch v1 {
                            case "":
                                var v143 uint16 = 350
                                _ = v143
                            case "foo bar":
                                var v144 int32 = 2147483647
                                _ = v144
                            }
                        case "":
                            if ((!(!(0 != 32767))) || true) {
                                switch (v4 >> ((v75) & 63)) {
                                case 42:
                                    if (922 >= (v75 | (v75 & v75))) {
                                        v3 = (v116 + (^(v3)))
                                        v75 = 4294967295
                                        v5 += (fi115 / ((-949) | 1))
                                    } else {
                                        if (3.14159 >= 1e10) {
                                            var v145 float32 = 0.1
                                            v26 = ((("foo bar" + "") + (v129 + "a")) + v26)
                                            _ = v145
                                        } else {
                                            if true {
                                                var v146 uint8 = v13
                                                var v147 float32 = float32(0.0)
                                                if len(v128) >= 2 {
                                                    v128 = v128[1:len(v128)]
                                                }
                                                _ = v146
                                                _ = v147
                                            }
                                            v26 = "\n\t"
                                        }
                                        v26 = "\n\t"
                                        if len(v128) >= 2 {
                                            v128 = v128[1:len(v128)]
                                        }
                                    }
                                case -2147483648:
                                    var v148 uint16 = 210
                                    _ = v148
                                }
                                fmt.Println(v75)
                                if len(v128) >= 2 {
                                    v128 = v128[1:len(v128)]
                                }
                                var v149 int16 = -1
                                v13 -= v13
                                _ = v149
                            }
                        }
                        var v150 int32 = 42
                        _ = v128
                        _ = v129
                        _ = v150
                    }
                    _ = v127
                }
                v151 := GenS17{F0: v4, F1: -27, F2: (v13 << ((255) & 63))}
                v1 = "a"
                _ = v151
            }
            if (-876 == 1) {
                v152 := GenS2{F0: 803}
                v153 := GenS18{F0: false}
                _ = v152
                _ = v153
            }
            v154 := []int{}
            _ = fi117
            _ = v154
        }
        _ = fi115
        _ = v116
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v75:", v75)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v72
    _ = v75
}
