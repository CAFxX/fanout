package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int; F1 bool; F2 int8 }
type GenS3 struct { F0 uint8; F1 bool; F2 int8 }
type GenS4 struct { F0 uint32; F1 uint }
type GenS5 struct { F0 uint32; F1 uint16; F2 uint16 }
type GenS6 struct { F0 bool; F1 int8 }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 string }
type GenS9 struct { F0 int }
type GenS10 struct { F0 int8; F1 uint }
type GenS11 struct { F0 uint16; F1 int; F2 uint }
type GenS12 struct { F0 int; F1 int; F2 int8 }
func genFn0() int {
    var v7 int32 = 0
    v8 := GenS1{F0: 352}
    v7 = (v7 + (v7 ^ (^(v7))))
    for fi9 := 0; fi9 < 16; fi9++ {
        v10 := GenS1{F0: 79}
        for fi11 := 0; fi11 < 14; fi11++ {
            fmt.Println(fi9)
            fmt.Println(v7, fi9)
            v7 = (v7 - v7)
            _ = fi11
        }
        _ = fi9
        _ = v10
    }
    _ = v7
    _ = v8
    return 128
}

func main() {
    v1 := []uint32{2, 422}
    var v2 int16 = -32768
    var v3 float64 = 0.0
    var v4 float32 = (0.1 + ((float32(v3) * ((0.1 * float32(((float64(1e10) * (-2.25 + 1.5)) / v3))) * 1e10)) / ((((float32((v3 * 1.5)) - (0.0 * float32((v3 / v3)))) / (3.14159 / 1.0)) - 1.5) + ((1.5 + (1e10 / ((((1e10 * 1e10) - -2.25) * -2.25) - -2.25))) + (0.1 * ((0.1 + 3.14159) / (1e10 * (-2.25 / 1.0))))))))
    v5 := make(map[int]string)
    v6 := make(map[int]uint32)
    var v12 string = ("hello" + (("\n\t" + "foo bar") + ("\n\t" + "\n\t")))
    var v13 uint8 = 33
    var v14 uint64 = 2
    v15 := GenS1{F0: 0}
    var v16 uint8 = 0
    v16 = 42
    v16 = (^(v16))
    switch -2147483648 {
    case 1:
        var v17 int64 = -1
        _ = v17
    case 956:
        var v18 uint64 = (v14 - (v14 - 1))
        _ = v18
    case 485:
        var v19 int = 4294967295
        _ = v19
    }
    v2 = v2
    if (!(((("" + v12) < "a") || false) || (!true))) {
        var v20 uint8 = v16
        _ = len(v1)
        var v21 float64 = 1.5
        _ = v20
        _ = v21
    }
    for fi22 := 0; fi22 < 13; fi22++ {
        var v23 int32 = 42
        _ = fi22
        _ = v23
    }
    var v24 string = v12
    v25 := GenS2{F0: 807, F1: (((float64(v4) * (-2.25 + -2.25)) * 0.1) == ((v3 * v3) * v3)), F2: 100}
    switch v12 {
    case "":
        if true {
            for fi26 := 0; fi26 < 4; fi26++ {
                var v27 uint16 = 2
                if ((^(v13)) < (v16 ^ (^(v13)))) {
                    fmt.Println(v16)
                    switch (fi26 - -503) {
                    case 42:
                        v28 := GenS2{F0: (^(fi26)), F1: (18446744073709551615 >= 287), F2: 1}
                        _ = v28
                    case 2147483647:
                        _ = len(v1)
                    case -9223372036854775808:
                        v4 += 1.5
                    }
                }
                _ = fi26
                _ = v27
            }
            v29 := GenS3{F0: 186, F1: true, F2: 100}
            v13 += v13
            _ = v29
        }
    }
    for fi30 := 0; fi30 < 17; fi30++ {
        if (900 != 2) {
            _ = len(v1)
        }
        for fi31 := 0; fi31 < 9; fi31++ {
            switch "\n\t" {
            case "\n\t":
                if false {
                    v12 = (("foo bar" + ("foo bar" + "\n\t")) + ("foo bar" + v12))
                    switch fi31 {
                    case 4294967295:
                        if (65535 > 255) {
                            delete(v6, (fi31 << ((99) & 63)))
                            v4 += float32(0.1)
                        } else {
                            if len(v1) > 0 {
                                v1[0] = 0
                            }
                            v32 := v6[fi31]
                            delete(v5, fi31)
                            if len(v1) > 0 {
                                v1[0] = v32
                            }
                            _ = v32
                        }
                    default:
                        var v33 int16 = 1000
                        _ = v33
                    }
                } else {
                    var v34 uint16 = 565
                    _ = v34
                }
            default:
                var v35 uint16 = 255
                _ = v35
            }
            v24 = (("hello" + v12) + v12)
            if true {
                if (float64((v4 - v4)) >= v3) {
                    if (2147483647 < 4294967295) {
                        switch v24 {
                        case "hello":
                            v24 = v12
                            var v36 uint = 18446744073709551615
                            _ = v36
                        }
                        var v37 int8 = 0
                        fmt.Println(fi30, v4, v13)
                        _ = v37
                    } else {
                        var v38 int = fi30
                        v14 += 0
                        v39 := GenS4{F0: 2, F1: 2}
                        _ = v38
                        _ = v39
                    }
                    v40 := GenS2{F0: genFn0(), F1: true, F2: 34}
                    v16 = v16
                    v14 -= 0
                    _ = v40
                } else {
                    v14 += v14
                    var v41 uint64 = v14
                    var v42 int64 = 490
                    _ = v41
                    _ = v42
                }
                v3 = (v3 * v3)
            } else {
                if false {
                    fmt.Println(v14, fi31)
                    v2 = (v2 << ((2) & 63))
                    v43 := GenS5{F0: 42, F1: uint16(v13), F2: 42}
                    _ = v43
                }
                delete(v6, fi30)
            }
            _ = fi31
        }
        _ = fi30
    }
    var v44 int16 = v2
    switch (("\n\t" + v24) + v24) {
    case "\n\t":
        if false {
            var v45 bool = (86 != 188)
            _ = v45
        } else {
            fmt.Println(v4, v16)
            for fi46 := 0; fi46 < 2; fi46++ {
                v13 = 42
                for fi47 := 0; fi47 < 2; fi47++ {
                    var v48 uint32 = 555
                    _ = fi47
                    _ = v48
                }
                var v49 uint = uint(v16)
                _ = fi46
                _ = v49
            }
        }
    case "a":
        var v50 uint = 70
        _ = v50
    default:
        switch v12 {
        case "":
            v14 = v14
        case "hello":
            if false {
                switch "" {
                case "foo bar":
                    v51 := GenS6{F0: false, F1: -1}
                    _ = v51
                }
                switch 34 {
                case -9223372036854775808:
                    var v52 string = v24
                    _ = v52
                case 32767:
                    if true {
                        var v53 uint16 = 415
                        v1 = append(v1, 311)
                        var v54 int8 = 48
                        _ = v53
                        _ = v54
                    } else {
                        switch v24 {
                        case "foo bar":
                            v14 = (^(v14))
                        case "\n\t":
                            if (float64(1.5) <= (v3 - (v3 / (1e10 * v3)))) {
                                for fi55 := 0; fi55 < 5; fi55++ {
                                    for fi56 := 0; fi56 < 12; fi56++ {
                                        v57 := v5[(fi56 * fi55)]
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        switch (fi55 - (fi55 << ((42) & 63))) {
                                        case -2147483648:
                                            if false {
                                                v58 := v5[(fi56 * fi55)]
                                                v59 := GenS1{F0: v14}
                                                v60 := GenS3{F0: v16, F1: false, F2: int8(fi55)}
                                                _ = v58
                                                _ = v59
                                                _ = v60
                                            }
                                        case 2147483647:
                                            var v61 int8 = 0
                                            _ = v61
                                        case -9223372036854775808:
                                            if true {
                                                switch fi55 {
                                                case 1:
                                                    v44 = (-(v2))
                                                default:
                                                    switch v24 {
                                                    case "":
                                                        var v62 uint = 200
                                                        _ = v62
                                                    default:
                                                        v63 := v5[genFn0()]
                                                        _ = v63
                                                    }
                                                }
                                                var v64 uint = 48
                                                if (!false) {
                                                    v44 = 974
                                                    var v65 bool = (-9223372036854775808 > -332)
                                                    _ = v65
                                                } else {
                                                    if (!true) {
                                                        fmt.Println(v4, v24)
                                                        if (65535 != 32767) {
                                                            if false {
                                                                if (42 >= 4294967295) {
                                                                    if (!((true && true) && true)) {
                                                                        v3 = v3
                                                                        var v66 uint32 = 473
                                                                        v67 := GenS3{F0: (v16 + (v13 << (((v64 - (v64 / ((uint(v16)) | 1)))) & 63))), F1: false, F2: -91}
                                                                        _ = v66
                                                                        _ = v67
                                                                    } else {
                                                                        if false {
                                                                            v68 := GenS5{F0: 2147483648, F1: 174, F2: uint16(v64)}
                                                                            v64 = (v64 & (v64 % ((v64) | 1)))
                                                                            _ = v68
                                                                        }
                                                                        if (((false && (!true)) || (0.1 > v3)) && false) {
                                                                            v16 -= v13
                                                                            v24 = "foo bar"
                                                                            v16 -= v16
                                                                            v69 := GenS7{F0: -128}
                                                                            v64 += v64
                                                                            _ = v69
                                                                        }
                                                                    }
                                                                    v70 := GenS7{F0: -128}
                                                                    _ = v70
                                                                } else {
                                                                    if (-128 == 0) {
                                                                        fmt.Println(fi55, v14, v3)
                                                                        v13 += v16
                                                                    } else {
                                                                        var v71 uint = (v64 + 1)
                                                                        if false {
                                                                            if false {
                                                                                fmt.Println(v24)
                                                                                if ((v4 == 0.1) || (!false)) {
                                                                                    v72 := GenS8{F0: "a"}
                                                                                    v73 := GenS1{F0: v14}
                                                                                    switch v57 {
                                                                                    case "foo bar":
                                                                                        if (("\n\t" >= "\n\t") && true) {
                                                                                            var v74 int64 = 22
                                                                                            v1 = append(v1, 864)
                                                                                            _ = v74
                                                                                        } else {
                                                                                            var v75 string = "\n\t"
                                                                                            _ = v75
                                                                                        }
                                                                                    case "":
                                                                                        var v76 string = v57
                                                                                        _ = v76
                                                                                    case "\n\t":
                                                                                        switch v57 {
                                                                                        case "hello":
                                                                                            var v77 float64 = float64(0.1)
                                                                                            _ = v77
                                                                                        default:
                                                                                            delete(v6, (fi55 >> (((v64 % ((v71) | 1))) & 63)))
                                                                                        }
                                                                                    }
                                                                                    _ = v72
                                                                                    _ = v73
                                                                                } else {
                                                                                    v6[fi56] = 866
                                                                                }
                                                                                switch (fi55 * (fi56 - genFn0())) {
                                                                                case -2147483648:
                                                                                    var v78 int = (fi56 & (fi56 % (((fi56 ^ int(v64))) | 1)))
                                                                                    v16 += 127
                                                                                    v79 := GenS2{F0: fi55, F1: false, F2: 1}
                                                                                    v80 := GenS3{F0: (v13 % (((v13 >> (((v64 + uint(v14))) & 63))) | 1)), F1: false, F2: 42}
                                                                                    if true {
                                                                                        v81 := GenS3{F0: (v13 - uint8(v64)), F1: (253 <= 915), F2: -117}
                                                                                        _ = v81
                                                                                    } else {
                                                                                        v64 += v64
                                                                                    }
                                                                                    var v82 bool = (((((0.1 * (v3 * ((0.1 * (v3 * 0.0)) + (v3 / (v3 - v3))))) - -2.25) + ((float64(1e10) - 0.0) * v3)) != float64((v4 * -2.25))) && (false && true))
                                                                                    v83 := v6[4294967295]
                                                                                    fmt.Println(v64, v2, v12)
                                                                                    fmt.Println(v83)
                                                                                    _ = v78
                                                                                    _ = v79
                                                                                    _ = v80
                                                                                    _ = v82
                                                                                    _ = v83
                                                                                case -662:
                                                                                    v71 = v71
                                                                                default:
                                                                                    v16 = uint8(v14)
                                                                                }
                                                                            } else {
                                                                                var v84 bool = (false || ((fi56 >= genFn0()) && (!(v4 != ((1e10 * v4) / 1.0)))))
                                                                                _ = v84
                                                                            }
                                                                            if len(v1) > 0 {
                                                                                v1[0] = 65535
                                                                            }
                                                                        } else {
                                                                            fmt.Println(v12, fi55, v71)
                                                                            var v85 bool = false
                                                                            fmt.Println(v64, fi55)
                                                                            _ = v85
                                                                        }
                                                                        _ = v71
                                                                    }
                                                                    var v86 string = v12
                                                                    _ = v86
                                                                }
                                                            }
                                                            switch (("hello" + "foo bar") + v12) {
                                                            case "":
                                                                if false {
                                                                    v44 -= (v2 ^ (-(v44)))
                                                                }
                                                            default:
                                                                if ((false && true) || false) {
                                                                    fmt.Println(v44, v64)
                                                                }
                                                            }
                                                            v87 := GenS1{F0: (v14 + uint64(v64))}
                                                            _ = v87
                                                        }
                                                        var v88 float32 = 1e10
                                                        v44 = v2
                                                        _ = v88
                                                    } else {
                                                        var v89 float64 = (((float64(1.5) + (v3 + 3.14159)) + (0.1 + v3)) / v3)
                                                        var v90 int8 = 123
                                                        v4 += float32(v89)
                                                        _ = v89
                                                        _ = v90
                                                    }
                                                }
                                                _ = v64
                                            } else {
                                                v91 := GenS4{F0: 2147483647, F1: 0}
                                                if true {
                                                    v92 := GenS9{F0: (^(fi56))}
                                                    if false {
                                                        var v93 int64 = -9223372036854775808
                                                        _ = v93
                                                    } else {
                                                        var v94 uint32 = 65535
                                                        var v95 int8 = 1
                                                        _ = v94
                                                        _ = v95
                                                    }
                                                    _ = v92
                                                }
                                                _ = v91
                                            }
                                        }
                                        _ = fi56
                                        _ = v57
                                    }
                                    var v96 uint32 = 2147483647
                                    _ = fi55
                                    _ = v96
                                }
                                if ((!true) || (((v14 - v14) == v14) || true)) {
                                    for fi97 := 0; fi97 < 3; fi97++ {
                                        var v98 uint64 = (v14 >> ((65535) & 63))
                                        _ = fi97
                                        _ = v98
                                    }
                                }
                            } else {
                                var v99 float64 = float64(v4)
                                for fi100 := 0; fi100 < 11; fi100++ {
                                    v2 = (-(v2))
                                    v101 := GenS8{F0: v12}
                                    _ = fi100
                                    _ = v101
                                }
                                delete(v6, 387)
                                _ = v99
                            }
                        default:
                            v16 -= v16
                        }
                        var v102 int16 = 1000
                        _ = v102
                    }
                case 32768:
                    switch "foo bar" {
                    case "foo bar":
                        v103 := GenS3{F0: v13, F1: true, F2: -64}
                        _ = v103
                    case "\n\t":
                        if true {
                            v4 += (0.0 * (v4 / 0.1))
                            v13 -= v16
                            v4 = (0.0 - v4)
                        }
                    case "hello":
                        v44 -= v2
                    }
                }
                v104 := GenS10{F0: 94, F1: 4294967295}
                _ = v104
            } else {
                v105 := GenS11{F0: 1, F1: genFn0(), F2: 18446744073709551615}
                _ = len(v1)
                for fi106 := 0; fi106 < 7; fi106++ {
                    if len(v1) > 0 {
                        v1[0] = 4294967295
                    }
                    if ((v16 / ((v13) | 1)) != v13) {
                        var v107 uint = 1
                        _ = v107
                    } else {
                        var v108 uint = 2
                        if (!(!false)) {
                            for fi109 := 0; fi109 < 16; fi109++ {
                                if true {
                                    v110 := GenS12{F0: 1, F1: fi109, F2: 100}
                                    _ = v110
                                }
                                v3 = v3
                                var v111 uint16 = 255
                                _ = fi109
                                _ = v111
                            }
                            v13 -= (v13 & (^(v16)))
                            fmt.Println(v12)
                        } else {
                            switch (fi106 >> ((v108) & 63)) {
                            case 42:
                                for fi112 := 0; fi112 < 8; fi112++ {
                                    fmt.Println(v24, v108, v44)
                                    switch fi112 {
                                    case 9223372036854775807:
                                        v24 = v12
                                    }
                                    v6[-2147483648] = 2147483647
                                    switch "" {
                                    case "\n\t":
                                        fmt.Println(v14, v2)
                                    case "a":
                                        var v113 float64 = -2.25
                                        _ = v113
                                    }
                                    if ((v12 > v12) && true) {
                                        v1 = append(v1, 2147483648)
                                    }
                                    var v114 int64 = -1
                                    _ = fi112
                                    _ = v114
                                }
                            case 539:
                                fmt.Println(v14, fi106)
                            case 32767:
                                v16 = v13
                            }
                            switch "foo bar" {
                            case "a":
                                for fi115 := 0; fi115 < 10; fi115++ {
                                    v116 := GenS6{F0: ((!(v4 > float32(v3))) && (!((v14 - 569) == 42))), F1: -128}
                                    switch (-(fi115)) {
                                    case 127:
                                        _ = len(v6)
                                    case 2147483647:
                                        fmt.Println(v24, v108, v14)
                                    default:
                                        var v117 int32 = -1
                                        _ = v117
                                    }
                                    var v118 float64 = (v3 - (v3 / 1.5))
                                    _ = fi115
                                    _ = v116
                                    _ = v118
                                }
                            case "hello":
                                fmt.Println(v2, v4, fi106)
                            case "":
                                var v119 int16 = (v44 >> ((897) & 63))
                                _ = v119
                            }
                        }
                        _ = v108
                    }
                    v14 -= v14
                    _ = fi106
                }
                _ = v105
            }
        default:
            v3 = float64(-2.25)
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v44
    _ = v5
    _ = v6
}
