package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8; F2 int8 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int64; F1 bool; F2 int8 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint16; F1 uint16; F2 uint }
type GenS6 struct { F0 uint16; F1 uint8; F2 uint }
type GenS7 struct { F0 uint64; F1 int; F2 uint }
type GenS8 struct { F0 int8 }
type GenS9 struct { F0 uint; F1 int32 }
type GenS10 struct { F0 int16; F1 int32 }
type GenS11 struct { F0 int64; F1 int8 }
type GenS12 struct { F0 int64; F1 string; F2 int }
type GenS13 struct { F0 bool; F1 uint32; F2 uint }

func main() {
    var v1 uint = 0
    v2 := make(map[string]bool)
    var v3 string = ("" + ("\n\t" + "hello"))
    v4 := make(map[int]uint32)
    var v5 uint32 = 4294967295
    var v6 int64 = 1
    for fi7 := 0; fi7 < 13; fi7++ {
        v5 += 750
        fmt.Println(v5, v1, v6)
        _ = fi7
    }
    if ((246 <= 2) || true) {
        if (0 != 1) {
            var v8 float32 = 3.14159
            v9 := []int{}
            if ((((1.5 * ((1.5 / 1.0) - 0.1)) + 0.1) + float64((1.5 - float32(((float64(0.1) + ((1e10 + -2.25) + 0.0)) / 1.5))))) > (3.14159 * (float64(1.5) * -2.25))) {
                switch v3 {
                case "a":
                    for fi10 := 0; fi10 < 8; fi10++ {
                        for fi11 := 0; fi11 < 17; fi11++ {
                            var v12 uint64 = 149
                            v3 = (v3 + ((v3 + (((v3 + (v3 + "")) + ((v3 + v3) + v3)) + v3)) + v3))
                            _ = fi11
                            _ = v12
                        }
                        if false {
                            if true {
                                var v13 float64 = -2.25
                                v14 := GenS1{F0: 2, F1: 208, F2: -1}
                                _ = v13
                                _ = v14
                            }
                            for fi15 := 0; fi15 < 3; fi15++ {
                                if (false || false) {
                                    if (-1 != 926) {
                                        switch (^(fi15)) {
                                        case -1:
                                            v16 := GenS1{F0: 65535, F1: 2, F2: -128}
                                            _ = v16
                                        case 32768:
                                            v8 += v8
                                        case -9223372036854775808:
                                            var v17 uint = v1
                                            v3 = (v3 + ("foo bar" + v3))
                                            _ = v17
                                        default:
                                            fmt.Println(v8)
                                        }
                                    }
                                    v3 = (("\n\t" + "") + v3)
                                    fmt.Println(fi15, fi10, v5)
                                    var v18 uint32 = 2
                                    v1 += v1
                                    var v19 float64 = (3.14159 * 0.0)
                                    v4[fi15] = 264
                                    var v20 uint32 = v18
                                    _ = v18
                                    _ = v19
                                    _ = v20
                                }
                                var v21 uint8 = 42
                                var v22 int32 = -2147483648
                                var v23 int = (fi10 ^ fi10)
                                v5 = (v5 - (v5 % ((uint32(v1)) | 1)))
                                _ = fi15
                                _ = v21
                                _ = v22
                                _ = v23
                            }
                            v8 = v8
                        }
                        _ = fi10
                    }
                case "":
                    v24 := GenS2{F0: 0}
                    _ = v24
                }
                fmt.Println(v1)
                if true {
                    v25 := GenS2{F0: 902}
                    v8 -= (-2.25 / 0.1)
                    if (1 == 2) {
                        if ((18446744073709551615 > (v1 + v1)) && (!false)) {
                            delete(v4, 603)
                            for fi26 := 0; fi26 < 13; fi26++ {
                                switch (^(fi26)) {
                                case 318:
                                    for fi27 := 0; fi27 < 14; fi27++ {
                                        var v28 uint32 = (^(v5))
                                        v29 := GenS3{F0: 1, F1: (v8 != (v8 * v8)), F2: 100}
                                        _ = fi27
                                        _ = v28
                                        _ = v29
                                    }
                                case -9223372036854775808:
                                    v6 += (^(v6))
                                }
                                var v30 int32 = -1
                                for fi31 := 0; fi31 < 16; fi31++ {
                                    if len(v9) > 0 {
                                        v9[0] = -456
                                    }
                                    if (false || (v30 < v30)) {
                                        v32 := GenS4{F0: 1}
                                        _ = v32
                                    } else {
                                        v1 += (v1 ^ v1)
                                        if ((fi31 - (-(fi31))) < (fi26 % (((fi31 << ((v1) & 63))) | 1))) {
                                            if false {
                                                if true {
                                                    v33 := GenS5{F0: 2, F1: 255, F2: v1}
                                                    var v34 int64 = (^(v6))
                                                    v8 += ((3.14159 * ((1e10 - v8) - v8)) + (-2.25 / v8))
                                                    _ = v33
                                                    _ = v34
                                                }
                                                var v35 int16 = -860
                                                switch v3 {
                                                case "":
                                                    var v36 int32 = v30
                                                    _ = v36
                                                case "hello":
                                                    delete(v4, fi26)
                                                    var v37 uint = v1
                                                    v30 = v30
                                                    _ = v37
                                                }
                                                _ = v35
                                            }
                                            var v38 bool = false
                                            _ = v38
                                        }
                                    }
                                    _ = fi31
                                }
                                _ = fi26
                                _ = v30
                            }
                        } else {
                            v8 += 0.1
                            if false {
                                for fi39 := 0; fi39 < 20; fi39++ {
                                    for fi40 := 0; fi40 < 10; fi40++ {
                                        v41 := GenS6{F0: 42, F1: 111, F2: (v1 << (((v1 | v1)) & 63))}
                                        v6 += (v6 >> ((42) & 63))
                                        v9 = append(v9, (^(fi40)))
                                        _ = fi40
                                        _ = v41
                                    }
                                    var v42 uint32 = (v5 & 42)
                                    _ = fi39
                                    _ = v42
                                }
                                var v43 int32 = 42
                                _ = v43
                            }
                            for fi44 := 0; fi44 < 3; fi44++ {
                                delete(v4, 913)
                                v3 = ("foo bar" + v3)
                                fmt.Println(v8)
                                _ = fi44
                            }
                        }
                        if true {
                            v8 -= v8
                        } else {
                            if true {
                                var v45 bool = true
                                _ = v45
                            }
                        }
                        var v46 float64 = float64(v8)
                        _ = v46
                    } else {
                        for fi47 := 0; fi47 < 7; fi47++ {
                            var v48 uint = 42
                            if (-1 == 42) {
                                v49 := GenS2{F0: 588}
                                switch fi47 {
                                case 2147483647:
                                    v5 -= (v5 >> (((^(v48))) & 63))
                                case 4294967295:
                                    delete(v4, fi47)
                                default:
                                    v50 := GenS5{F0: 255, F1: 255, F2: v48}
                                    _ = v50
                                }
                                fmt.Println(v6, v3, v5)
                                _ = v49
                            }
                            if (42 != 587) {
                                v48 = (v48 ^ v48)
                                for fi51 := 0; fi51 < 4; fi51++ {
                                    var v52 uint32 = v5
                                    var v53 uint8 = 1
                                    _ = len(v4)
                                    _ = fi51
                                    _ = v52
                                    _ = v53
                                }
                            } else {
                                v3 = v3
                            }
                            _ = fi47
                            _ = v48
                        }
                        if true {
                            switch v3 {
                            case "hello":
                                switch -9223372036854775808 {
                                case 127:
                                    v54 := GenS7{F0: 42, F1: -1, F2: 2}
                                    _ = v54
                                case 379:
                                    if (0 < 1) {
                                        if true {
                                            if (v1 <= 2) {
                                                fmt.Println(v6, v8, v8)
                                                for fi55 := 0; fi55 < 12; fi55++ {
                                                    for fi56 := 0; fi56 < 9; fi56++ {
                                                        v3 = v3
                                                        switch fi55 {
                                                        case 42:
                                                            v57 := GenS5{F0: 65535, F1: 2, F2: (v1 ^ (v1 << (((^(v1))) & 63)))}
                                                            _ = v57
                                                        case -2147483648:
                                                            switch ("\n\t" + v3) {
                                                            case "foo bar":
                                                                if false {
                                                                    switch ((v3 + ("a" + ("foo bar" + v3))) + ((((v3 + (("\n\t" + "\n\t") + ("a" + (("\n\t" + v3) + (v3 + "hello"))))) + v3) + v3) + ((v3 + "a") + "foo bar"))) {
                                                                    case "\n\t":
                                                                        fmt.Println(v5, v6)
                                                                    default:
                                                                        delete(v2, v3)
                                                                    }
                                                                    fmt.Println(fi56, v1, fi55)
                                                                } else {
                                                                    switch fi56 {
                                                                    case 641:
                                                                        var v58 uint = (v1 & (v1 >> ((4294967295) & 63)))
                                                                        _ = v58
                                                                    case 41:
                                                                        v59 := v4[fi55]
                                                                        _ = v59
                                                                    case 2147483647:
                                                                        var v60 uint = (v1 ^ 868)
                                                                        _ = v60
                                                                    }
                                                                    v8 -= v8
                                                                    if (255 < 2) {
                                                                        fmt.Println(v6, v1)
                                                                        _ = len(v9)
                                                                    } else {
                                                                        var v61 float64 = 0.0
                                                                        var v62 uint32 = (v5 << ((v1) & 63))
                                                                        var v63 int16 = 32767
                                                                        _ = v61
                                                                        _ = v62
                                                                        _ = v63
                                                                    }
                                                                }
                                                            }
                                                        case -9223372036854775808:
                                                            switch 32767 {
                                                            case 0:
                                                                var v64 uint32 = (v5 << ((v1) & 63))
                                                                var v65 float32 = ((1e10 + v8) - (v8 / v8))
                                                                fmt.Println(v8, v5, v64)
                                                                _ = v64
                                                                _ = v65
                                                            case 4294967295:
                                                                v5 += 2147483647
                                                            case -1:
                                                                switch v3 {
                                                                case "foo bar":
                                                                    v6 -= (^(v6))
                                                                case "":
                                                                    var v66 uint64 = 502
                                                                    if ((true && true) || true) {
                                                                        fmt.Println(v3, fi55, v1)
                                                                        if (v66 > 1) {
                                                                            var v67 int64 = int64(fi55)
                                                                            _ = v67
                                                                        }
                                                                        var v68 uint64 = (v66 * v66)
                                                                        fmt.Println(v5)
                                                                        if true {
                                                                            if (!(0 > 903)) {
                                                                                var v69 uint8 = 244
                                                                                _ = v69
                                                                            }
                                                                            _ = len(v9)
                                                                            switch (fi56 - fi56) {
                                                                            case -9223372036854775808:
                                                                                v3 = v3
                                                                            case -1:
                                                                                var v70 uint16 = 824
                                                                                _ = v70
                                                                            case -700:
                                                                                var v71 int32 = -416
                                                                                _ = v71
                                                                            }
                                                                        } else {
                                                                            var v72 uint16 = 65535
                                                                            var v73 float64 = ((0.1 / (0.1 / 1.0)) * 0.1)
                                                                            _ = v72
                                                                            _ = v73
                                                                        }
                                                                        _ = v68
                                                                    }
                                                                    _ = v66
                                                                case "\n\t":
                                                                    if ((0.1 - (-2.25 * 1.5)) <= 3.14159) {
                                                                        var v74 int16 = -301
                                                                        var v75 float32 = 0.1
                                                                        _ = v74
                                                                        _ = v75
                                                                    }
                                                                default:
                                                                    var v76 uint64 = 522
                                                                    _ = v76
                                                                }
                                                            }
                                                        }
                                                        _ = fi56
                                                    }
                                                    var v77 uint32 = v5
                                                    _ = fi55
                                                    _ = v77
                                                }
                                                v8 += 1e10
                                            } else {
                                                v9 = append(v9, 931)
                                                var v78 bool = false
                                                _ = v78
                                            }
                                            v3 = ((v3 + "a") + "foo bar")
                                        } else {
                                            if len(v9) > 0 {
                                                v9[0] = 9223372036854775807
                                            }
                                            for fi79 := 0; fi79 < 16; fi79++ {
                                                v5 += 155
                                                if false {
                                                    if len(v9) >= 2 {
                                                        v9 = v9[1:len(v9)]
                                                    }
                                                } else {
                                                    if false {
                                                        var v80 uint16 = 2
                                                        _ = v80
                                                    }
                                                    var v81 int8 = 1
                                                    _ = v81
                                                }
                                                switch (^(fi79)) {
                                                case 4294967295:
                                                    switch fi79 {
                                                    case 4294967295:
                                                        v5 = (v5 | v5)
                                                    case 510:
                                                        if true {
                                                            v6 = (v6 * (v6 >> ((v1) & 63)))
                                                            if true {
                                                                for fi82 := 0; fi82 < 7; fi82++ {
                                                                    v6 = (v6 | (v6 & -458))
                                                                    v8 -= (v8 - 0.0)
                                                                    _ = fi82
                                                                }
                                                            } else {
                                                                var v83 uint8 = 128
                                                                _ = v83
                                                            }
                                                            v84 := GenS3{F0: 1, F1: (-2.25 <= (1.5 * 0.0)), F2: -1}
                                                            _ = v84
                                                        } else {
                                                            v8 += v8
                                                            v85 := GenS1{F0: 32767, F1: 127, F2: 122}
                                                            _ = v85
                                                        }
                                                    case 32767:
                                                        v3 = v3
                                                    default:
                                                        v86 := GenS8{F0: 99}
                                                        _ = v86
                                                    }
                                                case 181:
                                                    v87 := GenS9{F0: (v1 + uint(v5)), F1: 0}
                                                    _ = v87
                                                case -1:
                                                    var v88 int = fi79
                                                    _ = v88
                                                }
                                                _ = fi79
                                            }
                                        }
                                        if false {
                                            if (-1 != 100) {
                                                for fi89 := 0; fi89 < 6; fi89++ {
                                                    for fi90 := 0; fi90 < 16; fi90++ {
                                                        v6 = v6
                                                        if (42 == 32767) {
                                                            if len(v9) > 0 {
                                                                v9[0] = -323
                                                            }
                                                        }
                                                        _ = fi90
                                                    }
                                                    _ = fi89
                                                }
                                                var v91 uint32 = 678
                                                v6 = v6
                                                _ = v91
                                            } else {
                                                var v92 float64 = (0.1 + 3.14159)
                                                _ = v92
                                            }
                                        }
                                    } else {
                                        delete(v2, v3)
                                        for fi93 := 0; fi93 < 4; fi93++ {
                                            if false {
                                                if true {
                                                    for fi94 := 0; fi94 < 17; fi94++ {
                                                        fmt.Println(v1)
                                                        v1 -= (v1 * v1)
                                                        v4[fi93] = v5
                                                        _ = fi94
                                                    }
                                                }
                                            }
                                            v95 := v4[(fi93 - fi93)]
                                            _ = fi93
                                            _ = v95
                                        }
                                    }
                                default:
                                    if true {
                                        switch v3 {
                                        case "a":
                                            var v96 uint = v1
                                            _ = v96
                                        default:
                                            v2[v3] = false
                                        }
                                        v97 := GenS10{F0: -1, F1: -783}
                                        _ = v97
                                    } else {
                                        if false {
                                            v98 := v2[((v3 + v3) + v3)]
                                            fmt.Println(v3)
                                            _ = v98
                                        } else {
                                            delete(v2, v3)
                                            v5 = (v5 | (^(v5)))
                                        }
                                        v3 = (v3 + (v3 + v3))
                                    }
                                }
                            default:
                                var v99 uint16 = 1
                                _ = v99
                            }
                            if (18446744073709551615 >= 1) {
                                delete(v4, -2147483648)
                            } else {
                                if (39 < 54) {
                                    v3 = v3
                                } else {
                                    v8 = (v8 / 3.14159)
                                    var v100 string = (v3 + "hello")
                                    for fi101 := 0; fi101 < 18; fi101++ {
                                        for fi102 := 0; fi102 < 7; fi102++ {
                                            v8 = float32(1.5)
                                            v1 = 18446744073709551615
                                            _ = fi102
                                        }
                                        _ = fi101
                                    }
                                    _ = v100
                                }
                            }
                        } else {
                            v103 := GenS4{F0: 255}
                            v6 -= (v6 + v6)
                            _ = v103
                        }
                        fmt.Println(v5, v3)
                    }
                    _ = v25
                } else {
                    switch 32768 {
                    case 2147483647:
                        var v104 uint16 = 42
                        _ = v104
                    case 0:
                        var v105 float32 = v8
                        _ = v105
                    default:
                        var v106 float64 = float64(v8)
                        _ = v106
                    }
                }
            } else {
                for fi107 := 0; fi107 < 4; fi107++ {
                    v108 := v4[fi107]
                    _ = fi107
                    _ = v108
                }
                for fi109 := 0; fi109 < 19; fi109++ {
                    v5 += 65535
                    _ = fi109
                }
            }
            _ = v8
            _ = v9
        }
        v110 := GenS11{F0: (^(v6)), F1: 127}
        _ = v110
    } else {
        var v111 string = (v3 + (("foo bar" + (v3 + "hello")) + ("foo bar" + v3)))
        _ = v111
    }
    var v112 int16 = 0
    if true {
        var v113 string = v3
        switch -60 {
        case 40:
            v114 := GenS12{F0: (^(v6)), F1: v113, F2: 32768}
            _ = v114
        default:
            var v115 int8 = 127
            _ = v115
        }
        _ = v113
    } else {
        v116 := []int{}
        v1 -= (^(v1))
        _ = v116
    }
    fmt.Println(v1)
    v2[("a" + v3)] = (32767 >= 42)
    v1 -= (v1 & v1)
    v5 -= v5
    var v117 float32 = 3.14159
    fmt.Println(v5)
    fmt.Println(v6, v5, v5)
    if (false && (!(true || (0 >= 389)))) {
        var v118 int8 = 42
        delete(v4, -1)
        _ = v118
    }
    switch ("hello" + v3) {
    case "foo bar":
        if (255 > 2) {
            if ((false || (!(!(true && false)))) || true) {
                for fi119 := 0; fi119 < 4; fi119++ {
                    v3 = "foo bar"
                    if (142 > 159) {
                        v120 := GenS7{F0: 4294967295, F1: (fi119 + fi119), F2: 2}
                        _ = v120
                    } else {
                        if true {
                            v5 = (^(v5))
                        }
                        var v121 int8 = 48
                        var v122 int16 = (v112 << (((v1 >> ((0) & 63))) & 63))
                        _ = v121
                        _ = v122
                    }
                    _ = fi119
                }
            }
            var v123 int64 = v6
            v124 := []int{567, 400}
            _ = v123
            _ = v124
        } else {
            for fi125 := 0; fi125 < 8; fi125++ {
                v126 := []int{int(v112), fi125, 372}
                v5 += v5
                _ = fi125
                _ = v126
            }
            v127 := []int{}
            _ = v127
        }
    default:
        switch 32767 {
        case 556:
            var v128 int8 = 0
            _ = v128
        case 127:
            delete(v4, -976)
        }
    }
    v129 := GenS13{F0: true, F1: v5, F2: 255}
    if false {
        v112 += (v112 | (^(v112)))
        if (130 >= 128) {
            v6 += (^(v6))
        } else {
            v130 := GenS10{F0: (v112 / (((^(v112))) | 1)), F1: -854}
            var v131 uint = 1
            _ = v130
            _ = v131
        }
        v112 = -32768
    }
    fmt.Println("v1:", v1)
    fmt.Println("v112:", v112)
    fmt.Println("v117:", v117)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v112
    _ = v117
    _ = v129
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
