package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32 }
type GenS2 struct { F0 uint32; F1 int64 }
type GenS3 struct { F0 int16; F1 uint }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint64; F1 int64 }
type GenS6 struct { F0 string }
type GenS7 struct { F0 string; F1 bool; F2 int16 }
type GenS8 struct { F0 uint8 }
type GenS9 struct { F0 int32; F1 uint; F2 uint8 }
type GenS10 struct { F0 int32; F1 int8; F2 uint16 }
type GenS11 struct { F0 uint8; F1 bool }
type GenS12 struct { F0 int32; F1 int32 }

func main() {
    var v1 uint16 = 422
    var v2 uint8 = 26
    v3 := make(map[string]uint32)
    var v4 uint16 = v1
    var v5 uint32 = 4294967295
    var v6 float32 = (1.5 + ((0.1 * 0.1) + 3.14159))
    v7 := []int{-455, -719, -811}
    v8 := GenS1{F0: 69, F1: v5}
    v9 := GenS1{F0: 320, F1: (^(v5))}
    if true {
        switch -9223372036854775808 {
        case 42:
            var v10 int64 = 969
            _ = v10
        default:
            var v11 float32 = (v6 * 1.5)
            _ = v11
        }
        v12 := v3["hello"]
        v4 += 744
        _ = v12
    }
    for fi13 := 0; fi13 < 14; fi13++ {
        for fi14 := 0; fi14 < 20; fi14++ {
            var v15 int64 = 42
            v16 := GenS2{F0: 209, F1: (v15 << ((42) & 63))}
            var v17 int64 = (v15 & (v15 * v15))
            _ = fi14
            _ = v15
            _ = v16
            _ = v17
        }
        _ = fi13
    }
    switch "a" {
    case "a":
        switch -2147483648 {
        case 127:
            var v18 float32 = v6
            _ = v18
        case -214:
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
        case 9223372036854775807:
            var v19 uint16 = v4
            _ = v19
        default:
            v3["a"] = v5
        }
    case "hello":
        v20 := GenS2{F0: v5, F1: 0}
        _ = v20
    default:
        if false {
            var v21 int16 = 103
            _ = v21
        }
    }
    var v22 int32 = -3
    v1 = (v4 * v4)
    v2 -= v2
    _ = len(v7)
    v2 = v2
    for fi23 := 0; fi23 < 1; fi23++ {
        if false {
            var v24 uint16 = v4
            var v25 int64 = 42
            _ = v24
            _ = v25
        }
        delete(v3, "hello")
        _ = fi23
    }
    for fi26 := 0; fi26 < 20; fi26++ {
        _ = len(v7)
        if true {
            var v27 int = (fi26 * fi26)
            if true {
                _ = len(v7)
            } else {
                for fi28 := 0; fi28 < 17; fi28++ {
                    var v29 uint = 161
                    _ = fi28
                    _ = v29
                }
                if ((!true) && true) {
                    v30 := GenS2{F0: (v5 >> ((867) & 63)), F1: 42}
                    if ((!false) || ("a" > ("" + ("\n\t" + "hello")))) {
                        var v31 int8 = int8(v27)
                        v3[("a" + "\n\t")] = v5
                        _ = v31
                    } else {
                        if (-9223372036854775808 == 1) {
                            var v32 int32 = v22
                            switch 271 {
                            case -636:
                                if (false || false) {
                                    v33 := GenS2{F0: 881, F1: -900}
                                    _ = v33
                                }
                            case 128:
                                v34 := GenS1{F0: 4294967295, F1: v5}
                                _ = v34
                            case 32767:
                                v3[(("\n\t" + ("foo bar" + ("a" + "a"))) + "")] = uint32(v2)
                            }
                            _ = v32
                        } else {
                            switch "a" {
                            case "\n\t":
                                v1 = v4
                            case "a":
                                for fi35 := 0; fi35 < 12; fi35++ {
                                    v36 := GenS1{F0: 0, F1: 2}
                                    _ = fi35
                                    _ = v36
                                }
                            }
                        }
                        v22 -= v22
                    }
                    _ = v30
                }
            }
            var v37 string = "hello"
            _ = v27
            _ = v37
        } else {
            if false {
                for fi38 := 0; fi38 < 3; fi38++ {
                    switch fi26 {
                    case 128:
                        if false {
                            switch "\n\t" {
                            case "\n\t":
                                v5 = (v5 ^ (v5 >> ((747) & 63)))
                            default:
                                v39 := GenS3{F0: -1, F1: 556}
                                _ = v39
                            }
                            v4 = (v4 >> ((417) & 63))
                        }
                    default:
                        switch "" {
                        case "\n\t":
                            var v40 uint32 = (v5 >> ((255) & 63))
                            if false {
                                if len(v7) > 0 {
                                    v7[0] = fi38
                                }
                                var v41 int8 = 1
                                _ = v41
                            } else {
                                if (3.14159 > (1.5 - v6)) {
                                    if false {
                                        _ = len(v7)
                                    }
                                    v40 = (^(v40))
                                    v7 = append(v7, (-(fi38)))
                                } else {
                                    _ = len(v7)
                                    switch fi38 {
                                    case -271:
                                        v3["\n\t"] = 4294967295
                                    default:
                                        var v42 uint8 = v2
                                        _ = v42
                                    }
                                }
                                if len(v7) > 0 {
                                    v7[0] = 881
                                }
                            }
                            var v43 uint8 = (^(v2))
                            _ = v40
                            _ = v43
                        case "foo bar":
                            v1 = 0
                        }
                    }
                    switch 4294967295 {
                    case 128:
                        var v44 uint = 2
                        _ = v44
                    case -79:
                        v45 := GenS2{F0: (v5 / ((2147483647) | 1)), F1: -609}
                        _ = v45
                    case 32767:
                        var v46 int16 = 840
                        _ = v46
                    default:
                        v22 += v22
                    }
                    _ = fi38
                }
                if false {
                    var v47 int64 = 1
                    var v48 float64 = 0.0
                    _ = v47
                    _ = v48
                } else {
                    fmt.Println(v6, v22, v2)
                }
                var v49 float32 = float32((0.0 * (float64(float32(((1e10 * ((1.5 * 1e10) * 0.0)) + 0.0))) + 0.1)))
                _ = v49
            } else {
                var v50 uint = 0
                switch "" {
                case "":
                    for fi51 := 0; fi51 < 9; fi51++ {
                        switch fi26 {
                        case 0:
                            v52 := GenS4{F0: (v2 * v2)}
                            _ = v52
                        case 1:
                            var v53 uint32 = (v5 + 2147483647)
                            _ = v53
                        case 127:
                            fmt.Println(v5, v5)
                        default:
                            v22 = v22
                        }
                        var v54 uint64 = 4294967295
                        delete(v3, ("\n\t" + ("foo bar" + "a")))
                        _ = fi51
                        _ = v54
                    }
                case "a":
                    if false {
                        v22 = v22
                        v55 := GenS2{F0: (^(v5)), F1: int64(v4)}
                        v56 := GenS5{F0: 1, F1: 42}
                        _ = v55
                        _ = v56
                    } else {
                        v5 = uint32(v1)
                    }
                default:
                    if false {
                        var v57 uint8 = (v2 | v2)
                        _ = v57
                    } else {
                        v58 := GenS6{F0: ""}
                        _ = v58
                    }
                }
                if true {
                    v6 = 0.0
                }
                _ = v50
            }
            for fi59 := 0; fi59 < 11; fi59++ {
                v60 := GenS7{F0: "", F1: (4294967295 >= 18446744073709551615), F2: 1}
                v4 -= (^(v4))
                v22 += v22
                _ = fi59
                _ = v60
            }
        }
        v1 += 356
        _ = fi26
    }
    switch 9223372036854775807 {
    case 2147483647:
        switch 9223372036854775807 {
        case 4294967295:
            var v61 uint8 = uint8(v4)
            _ = v61
        case 1:
            for fi62 := 0; fi62 < 5; fi62++ {
                if true {
                    for fi63 := 0; fi63 < 14; fi63++ {
                        var v64 int16 = 32767
                        v65 := GenS8{F0: (v2 << ((uint(v1)) & 63))}
                        v66 := GenS8{F0: 247}
                        if true {
                            if (!false) {
                                var v67 int = fi62
                                var v68 uint = 0
                                if false {
                                    v6 = ((1e10 / (v6 + ((1.5 / 3.14159) * (v6 + 1e10)))) * float32(3.14159))
                                    var v69 int8 = -128
                                    fmt.Println(v67)
                                    switch "\n\t" {
                                    case "foo bar":
                                        var v70 float64 = 1e10
                                        _ = v70
                                    case "a":
                                        var v71 uint64 = 611
                                        _ = v71
                                    case "":
                                        if (((-(v69)) < (v69 | -128)) || false) {
                                            if (1 != 0) {
                                                fmt.Println(v67, v69, v68)
                                                v7 = append(v7, (fi63 - -9223372036854775808))
                                                if false {
                                                    fmt.Println(v68, v22, v5)
                                                    var v72 int16 = v64
                                                    var v73 float64 = 0.1
                                                    _ = v72
                                                    _ = v73
                                                }
                                                var v74 string = "a"
                                                switch v74 {
                                                case "\n\t":
                                                    if len(v7) >= 2 {
                                                        v7 = v7[1:len(v7)]
                                                    }
                                                    v68 = (^(v68))
                                                    var v75 uint32 = (^(v5))
                                                    _ = v75
                                                case "":
                                                    var v76 uint64 = 4294967295
                                                    _ = v76
                                                default:
                                                    var v77 int16 = v64
                                                    _ = v77
                                                }
                                                _ = v74
                                            } else {
                                                var v78 int8 = v69
                                                switch "\n\t" {
                                                case "":
                                                    var v79 uint32 = uint32(v68)
                                                    _ = v79
                                                case "\n\t":
                                                    if ((v68 << ((4294967295) & 63)) >= (v68 << ((v68) & 63))) {
                                                        v80 := GenS1{F0: 2, F1: (v5 ^ (v5 / ((v5) | 1)))}
                                                        _ = v80
                                                    }
                                                }
                                                _ = v78
                                            }
                                            switch "foo bar" {
                                            case "\n\t":
                                                if len(v7) >= 2 {
                                                    v7 = v7[1:len(v7)]
                                                }
                                            case "":
                                                switch fi63 {
                                                case 363:
                                                    var v81 uint64 = 1
                                                    _ = v81
                                                }
                                            default:
                                                v82 := GenS5{F0: 558, F1: 685}
                                                _ = v82
                                            }
                                            var v83 int32 = v22
                                            fmt.Println(v83, fi62, fi63)
                                            var v84 string = "a"
                                            _ = v83
                                            _ = v84
                                        }
                                    }
                                    _ = v69
                                } else {
                                    var v85 float64 = 0.1
                                    if false {
                                        v86 := GenS4{F0: (v2 ^ v2)}
                                        var v87 int8 = -1
                                        _ = v86
                                        _ = v87
                                    } else {
                                        v85 = (float64(1e10) - -2.25)
                                        v88 := GenS3{F0: v64, F1: (v68 * v68)}
                                        _ = v88
                                    }
                                    switch v67 {
                                    case -1:
                                        fmt.Println(v68, v85, v5)
                                    case 32768:
                                        v68 = v68
                                    }
                                    delete(v3, ("hello" + "a"))
                                    _ = v85
                                }
                                v68 = v68
                                delete(v3, "a")
                                _ = v67
                                _ = v68
                            }
                            if false {
                                fmt.Println(v5)
                                if len(v7) > 0 {
                                    v7[0] = -2147483648
                                }
                                fmt.Println(v22, v1, fi62)
                                if false {
                                    delete(v3, "foo bar")
                                    v89 := GenS9{F0: v22, F1: 42, F2: (v2 - 127)}
                                    v5 += (v5 - (v5 ^ v5))
                                    _ = v89
                                } else {
                                    v64 += v64
                                }
                            } else {
                                v90 := v3["a"]
                                _ = v90
                            }
                            if false {
                                if false {
                                    if ("foo bar" == "hello") {
                                        fmt.Println(v6, v64, v2)
                                        var v91 uint = 65535
                                        v92 := GenS10{F0: (v22 % ((-2147483648) | 1)), F1: 1, F2: uint16(v5)}
                                        v5 -= (^(v5))
                                        v3["foo bar"] = (^(v5))
                                        _ = v91
                                        _ = v92
                                    }
                                    v93 := GenS10{F0: (v22 - v22), F1: -128, F2: v1}
                                    v5 += (v5 / ((v5) | 1))
                                    if (false && (!false)) {
                                        var v94 int8 = -1
                                        fmt.Println(v22, fi62, v64)
                                        _ = v94
                                    }
                                    if true {
                                        v95 := GenS11{F0: (v2 + v2), F1: false}
                                        v1 = 302
                                        _ = len(v3)
                                        _ = v95
                                    } else {
                                        v6 = 0.1
                                        v96 := GenS12{F0: 0, F1: v22}
                                        var v97 string = "a"
                                        v98 := GenS12{F0: (v22 * v22), F1: v22}
                                        _ = v96
                                        _ = v97
                                        _ = v98
                                    }
                                    _ = v93
                                }
                                var v99 int64 = -442
                                _ = v99
                            }
                        } else {
                            _ = len(v7)
                        }
                        _ = fi63
                        _ = v64
                        _ = v65
                        _ = v66
                    }
                } else {
                    for fi100 := 0; fi100 < 4; fi100++ {
                        var v101 uint16 = v1
                        fmt.Println(v5)
                        _ = fi100
                        _ = v101
                    }
                }
                for fi102 := 0; fi102 < 9; fi102++ {
                    fmt.Println(v6)
                    switch (("" + "hello") + "\n\t") {
                    case "a":
                        if (1000 != 1) {
                            v5 -= 2147483647
                            v103 := GenS5{F0: 982, F1: -183}
                            v104 := GenS1{F0: 789, F1: v5}
                            _ = v103
                            _ = v104
                        }
                        if (false && (-171 < 0)) {
                            fmt.Println(v5, v4, v6)
                        }
                    default:
                        v6 = (0.0 / 1.0)
                    }
                    _ = fi102
                }
                _ = fi62
            }
        case 42:
            fmt.Println(v1, v2)
        }
    }
    var v105 int8 = -1
    fmt.Println(v22, v1, v6)
    for fi106 := 0; fi106 < 16; fi106++ {
        switch (("foo bar" + "a") + "a") {
        case "a":
            v107 := GenS11{F0: v2, F1: true}
            _ = v107
        case "":
            v6 -= (((((1e10 - (((0.1 / v6) / 1e10) / v6)) * -2.25) / 1e10) - (v6 - v6)) + v6)
        }
        var v108 int16 = 0
        _ = len(v7)
        _ = fi106
        _ = v108
    }
    fmt.Println("v1:", v1)
    fmt.Println("v105:", v105)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v105
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
