package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int16; F1 string }
type GenS5 struct { F0 int16; F1 bool; F2 uint64 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 int }
type GenS8 struct { F0 string; F1 int }
type GenS9 struct { F0 string; F1 int16 }
type GenS10 struct { F0 int8; F1 string }
type GenS11 struct { F0 uint16; F1 int16; F2 int64 }

func main() {
    v1 := make(map[int]string)
    var v2 uint8 = 255
    v3 := make(map[string]string)
    var v4 uint = 2
    var v5 uint = (v4 & 615)
    v6 := make(map[string]string)
    var v7 float32 = float32((0.0 * 1e10))
    for fi8 := 0; fi8 < 7; fi8++ {
        v4 += v5
        if (-1 <= int32(fi8)) {
            for fi9 := 0; fi9 < 15; fi9++ {
                v10 := []int{(^(fi8)), (fi9 ^ (fi8 << ((2) & 63))), fi9, fi9}
                var v11 int32 = 1
                _ = fi9
                _ = v10
                _ = v11
            }
            v12 := GenS1{F0: 42, F1: false}
            _ = v12
        }
        var v13 int16 = 1000
        _ = fi8
        _ = v13
    }
    var v14 int32 = 0
    v7 -= (v7 * ((0.1 - (float32(float64(float32(-2.25))) / (v7 + (v7 / v7)))) / (v7 * float32(float64((v7 * ((1.5 / v7) + (float32((0.1 * (0.0 / -2.25))) / ((v7 - 3.14159) + (v7 * v7))))))))))
    if (!(9223372036854775807 > -9223372036854775808)) {
        if (v7 == v7) {
            if (!(32767 == 1)) {
                for fi15 := 0; fi15 < 1; fi15++ {
                    var v16 string = "foo bar"
                    switch "hello" {
                    case "a":
                        switch (^(fi15)) {
                        case 4294967295:
                            var v17 int64 = 1
                            _ = v17
                        case 0:
                            if false {
                                v4 = v5
                                v18 := GenS1{F0: 2147483648, F1: false}
                                if false {
                                    v19 := []int{(fi15 << ((v5) & 63)), int(v4), fi15}
                                    if true {
                                        var v20 int16 = 962
                                        for fi21 := 0; fi21 < 6; fi21++ {
                                            if (!(42 <= 799)) {
                                                fmt.Println(fi15)
                                                v7 += v7
                                            } else {
                                                v20 = v20
                                            }
                                            _ = fi21
                                        }
                                        _ = v20
                                    } else {
                                        if false {
                                            for fi22 := 0; fi22 < 15; fi22++ {
                                                if (!(127 != int8(fi15))) {
                                                    var v23 uint16 = 0
                                                    v7 -= (((1e10 * v7) - float32(1.5)) * 0.1)
                                                    if ((v5 % ((v5) | 1)) <= (v4 / ((v5) | 1))) {
                                                        if len(v19) >= 2 {
                                                            v19 = v19[1:len(v19)]
                                                        }
                                                        v24 := GenS1{F0: 308, F1: true}
                                                        v25 := GenS1{F0: 767, F1: (!true)}
                                                        _ = v24
                                                        _ = v25
                                                    }
                                                    var v26 uint8 = v2
                                                    var v27 uint16 = v23
                                                    _ = v23
                                                    _ = v26
                                                    _ = v27
                                                }
                                                if true {
                                                    v28 := GenS2{F0: 261}
                                                    var v29 uint16 = 548
                                                    v30 := GenS2{F0: 227}
                                                    _ = v28
                                                    _ = v29
                                                    _ = v30
                                                }
                                                fmt.Println(fi22)
                                                switch 0 {
                                                case 355:
                                                    v31 := GenS1{F0: 826, F1: false}
                                                    switch v16 {
                                                    case "a":
                                                        v4 += v5
                                                    }
                                                    var v32 int = fi15
                                                    _ = v31
                                                    _ = v32
                                                case 128:
                                                    v33 := GenS1{F0: 2, F1: true}
                                                    _ = v33
                                                default:
                                                    if len(v19) >= 2 {
                                                        v19 = v19[1:len(v19)]
                                                    }
                                                }
                                                _ = fi22
                                            }
                                            delete(v6, v16)
                                            for fi34 := 0; fi34 < 17; fi34++ {
                                                if false {
                                                    v16 = ""
                                                } else {
                                                    if (v4 != (^(v4))) {
                                                        switch v16 {
                                                        case "hello":
                                                            v7 += float32(0.0)
                                                        }
                                                        fmt.Println(v16, fi15, fi34)
                                                        var v35 uint8 = 127
                                                        v7 = v7
                                                        if true {
                                                            var v36 uint32 = 0
                                                            v2 -= 1
                                                            v37 := v6[v16]
                                                            _ = v36
                                                            _ = v37
                                                        }
                                                        v35 = 127
                                                        var v38 int64 = 0
                                                        if (18446744073709551615 < 2) {
                                                            fmt.Println(v7, fi34, v14)
                                                            v39 := GenS1{F0: 743, F1: true}
                                                            _ = v39
                                                        }
                                                        _ = v35
                                                        _ = v38
                                                    }
                                                    if len(v19) > 0 {
                                                        v19[0] = fi15
                                                    }
                                                    v40 := GenS3{F0: 0}
                                                    _ = v40
                                                }
                                                _ = len(v19)
                                                var v41 int16 = -1
                                                _ = fi34
                                                _ = v41
                                            }
                                        } else {
                                            v14 = v14
                                            var v42 float64 = 1.5
                                            if false {
                                                var v43 uint64 = 2
                                                v4 = v4
                                                var v44 bool = false
                                                _ = v43
                                                _ = v44
                                            }
                                            _ = v42
                                        }
                                        v5 = v5
                                        var v45 int64 = -343
                                        _ = v45
                                    }
                                    for fi46 := 0; fi46 < 4; fi46++ {
                                        if false {
                                            var v47 int = int(v14)
                                            var v48 uint64 = 176
                                            if ((^(v14)) > v14) {
                                                if true {
                                                    if len(v19) > 0 {
                                                        v19[0] = v47
                                                    }
                                                    if true {
                                                        v19 = append(v19, (v47 | (fi46 * v47)))
                                                        if (true && false) {
                                                            var v49 int64 = -9223372036854775808
                                                            v50 := GenS4{F0: -1, F1: v16}
                                                            _ = v49
                                                            _ = v50
                                                        } else {
                                                            v4 -= 42
                                                            var v51 int = fi15
                                                            v52 := GenS1{F0: 127, F1: (1 == 748)}
                                                            _ = v51
                                                            _ = v52
                                                        }
                                                        switch v16 {
                                                        case "":
                                                            v53 := GenS2{F0: uint64(v4)}
                                                            _ = v53
                                                        case "a":
                                                            if false {
                                                                fmt.Println(v4)
                                                                switch (v47 - fi46) {
                                                                case 574:
                                                                    var v54 string = ((v16 + v16) + v16)
                                                                    _ = v54
                                                                }
                                                            } else {
                                                                v47 = (fi46 ^ (fi46 * fi46))
                                                                if (!(v4 < v4)) {
                                                                    if (!false) {
                                                                        var v55 float32 = v7
                                                                        v19 = append(v19, fi15)
                                                                        fmt.Println(v14, fi46)
                                                                        var v56 uint16 = 0
                                                                        var v57 string = ("hello" + "hello")
                                                                        _ = v55
                                                                        _ = v56
                                                                        _ = v57
                                                                    } else {
                                                                        v48 -= v48
                                                                        if true {
                                                                            var v58 string = "hello"
                                                                            var v59 int64 = 9223372036854775807
                                                                            var v60 float32 = (v7 / v7)
                                                                            _ = v58
                                                                            _ = v59
                                                                            _ = v60
                                                                        }
                                                                    }
                                                                    switch "\n\t" {
                                                                    case "hello":
                                                                        if (v7 != (3.14159 + ((v7 - -2.25) + (v7 + 1.5)))) {
                                                                            _ = len(v19)
                                                                            fmt.Println(v2, v47)
                                                                        }
                                                                    case "a":
                                                                        if true {
                                                                            if len(v19) >= 2 {
                                                                                v19 = v19[1:len(v19)]
                                                                            }
                                                                        } else {
                                                                            v2 = v2
                                                                            fmt.Println(v48, v14)
                                                                            if true {
                                                                                var v61 uint = 2
                                                                                var v62 float64 = (1e10 - 1e10)
                                                                                fmt.Println(v2, v14, v7)
                                                                                v7 = float32(v62)
                                                                                _ = v61
                                                                                _ = v62
                                                                            } else {
                                                                                if true {
                                                                                    fmt.Println(v4)
                                                                                } else {
                                                                                    var v63 uint8 = uint8(v5)
                                                                                    _ = v63
                                                                                }
                                                                                var v64 int32 = int32(v48)
                                                                                _ = v64
                                                                            }
                                                                            v14 = -2147483648
                                                                            var v65 int64 = 42
                                                                            var v66 int16 = -32768
                                                                            _ = v65
                                                                            _ = v66
                                                                        }
                                                                    default:
                                                                        switch (v16 + v16) {
                                                                        case "\n\t":
                                                                            v16 = ("" + ("" + "a"))
                                                                        case "":
                                                                            fmt.Println(v16, fi46)
                                                                        }
                                                                    }
                                                                    v67 := GenS4{F0: 414, F1: v16}
                                                                    _ = v67
                                                                }
                                                                if (v48 != v48) {
                                                                    var v68 int32 = (v14 << (((v5 % ((v5) | 1))) & 63))
                                                                    _ = v68
                                                                }
                                                            }
                                                            v7 = (v7 * 3.14159)
                                                            if false {
                                                                v69 := GenS2{F0: (v48 - (v48 + 90))}
                                                                var v70 bool = (v2 <= (v2 | 2))
                                                                _ = v69
                                                                _ = v70
                                                            }
                                                            var v71 uint32 = 640
                                                            _ = v71
                                                        default:
                                                            v19 = append(v19, -569)
                                                        }
                                                        v47 = (v47 & (-(v47)))
                                                        fmt.Println(v16, fi46)
                                                    } else {
                                                        var v72 int32 = -1
                                                        v7 -= -2.25
                                                        fmt.Println(v7)
                                                        var v73 float64 = 3.14159
                                                        v48 += (v48 ^ (v48 + v48))
                                                        _ = v72
                                                        _ = v73
                                                    }
                                                    var v74 int64 = -1
                                                    _ = v74
                                                } else {
                                                    v47 -= fi15
                                                    if false {
                                                        if true {
                                                            if true {
                                                                switch (("" + v16) + v16) {
                                                                case "foo bar":
                                                                    var v75 uint16 = 65535
                                                                    _ = v75
                                                                case "":
                                                                    if true {
                                                                        v76 := v6[v16]
                                                                        _ = v76
                                                                    }
                                                                default:
                                                                    switch (v47 / (((fi46 << (((v5 % ((v5) | 1))) & 63))) | 1)) {
                                                                    case 2147483647:
                                                                        v77 := GenS5{F0: -1, F1: false, F2: (v48 ^ 2)}
                                                                        _ = v77
                                                                    }
                                                                }
                                                                fmt.Println(v2, v5)
                                                                switch fi15 {
                                                                case -2147483648:
                                                                    v78 := GenS2{F0: (v48 | v48)}
                                                                    _ = v78
                                                                case -9223372036854775808:
                                                                    v79 := GenS6{F0: 0}
                                                                    _ = v79
                                                                case 152:
                                                                    v5 = v4
                                                                }
                                                            } else {
                                                                if true {
                                                                    switch ((v16 + "foo bar") + v16) {
                                                                    case "hello":
                                                                        switch (v16 + v16) {
                                                                        case "hello":
                                                                            var v80 float64 = 3.14159
                                                                            _ = v80
                                                                        }
                                                                    case "\n\t":
                                                                        if true {
                                                                            var v81 int = 2147483647
                                                                            _ = v81
                                                                        }
                                                                    default:
                                                                        fmt.Println(v4)
                                                                    }
                                                                    _ = len(v6)
                                                                    if true {
                                                                        var v82 int = (fi46 - (-(v47)))
                                                                        _ = v82
                                                                    }
                                                                }
                                                                var v83 uint32 = 0
                                                                _ = v83
                                                            }
                                                            v16 = ((v16 + "\n\t") + v16)
                                                            if (false || false) {
                                                                fmt.Println(v5, v4)
                                                                v47 -= 583
                                                                delete(v3, v16)
                                                            } else {
                                                                fmt.Println(v7, v16, v4)
                                                            }
                                                            fmt.Println(v48, v14, v4)
                                                            v2 = v2
                                                        }
                                                        var v84 uint64 = 655
                                                        _ = v84
                                                    }
                                                    v48 = (v48 + (^(v48)))
                                                }
                                                if true {
                                                    v48 = (v48 | 0)
                                                } else {
                                                    var v85 string = (v16 + "a")
                                                    var v86 string = "foo bar"
                                                    _ = v85
                                                    _ = v86
                                                }
                                            }
                                            _ = v47
                                            _ = v48
                                        } else {
                                            if len(v19) >= 2 {
                                                v19 = v19[1:len(v19)]
                                            }
                                            v5 = 2
                                        }
                                        if true {
                                            fmt.Println(v7)
                                            delete(v1, (fi46 | fi15))
                                            _ = len(v19)
                                            if (false || ((((float32(0.0) * v7) * v7) * 0.1) == v7)) {
                                                v4 = (v4 + 65535)
                                            } else {
                                                var v87 float32 = (((v7 * v7) * v7) - (float32(((((3.14159 / 3.14159) / (((3.14159 + -2.25) - 1.5) / 1.5)) * 1e10) * 0.0)) / 1.5))
                                                fmt.Println(v4)
                                                _ = v87
                                            }
                                            v88 := v1[(fi46 - (fi46 / ((-9223372036854775808) | 1)))]
                                            _ = v88
                                        } else {
                                            v5 -= (v5 | 865)
                                            v5 = (^(v4))
                                        }
                                        v14 += v14
                                        if true {
                                            v89 := v3[v16]
                                            _ = v89
                                        }
                                        _ = fi46
                                    }
                                    _ = v19
                                } else {
                                    var v90 bool = (v16 <= v16)
                                    _ = v90
                                }
                                _ = v18
                            } else {
                                v5 = (v5 / (((v5 >> ((42) & 63))) | 1))
                                v91 := GenS7{F0: (fi15 - fi15)}
                                for fi92 := 0; fi92 < 3; fi92++ {
                                    if (false && true) {
                                        v93 := []int{128}
                                        fmt.Println(v2, v5, fi15)
                                        _ = v93
                                    } else {
                                        v94 := []int{}
                                        _ = len(v6)
                                        _ = v94
                                    }
                                    var v95 float32 = ((((1e10 / v7) * v7) * float32(0.1)) / 1e10)
                                    _ = fi92
                                    _ = v95
                                }
                                _ = v91
                            }
                        default:
                            for fi96 := 0; fi96 < 14; fi96++ {
                                var v97 string = (v16 + v16)
                                v98 := GenS8{F0: "", F1: -27}
                                v99 := GenS8{F0: v16, F1: fi15}
                                _ = fi96
                                _ = v97
                                _ = v98
                                _ = v99
                            }
                        }
                    }
                    _ = fi15
                    _ = v16
                }
                var v100 float32 = float32(1e10)
                if (230 <= 2147483647) {
                    fmt.Println(v14, v5, v2)
                    v5 -= v4
                    if true {
                        var v101 uint64 = uint64(v2)
                        _ = v101
                    }
                } else {
                    fmt.Println(v7)
                    fmt.Println(v14, v7)
                    for fi102 := 0; fi102 < 7; fi102++ {
                        for fi103 := 0; fi103 < 16; fi103++ {
                            v104 := GenS9{F0: ("foo bar" + "a"), F1: -1}
                            v7 += float32((0.1 - (((((0.1 - float64(float32((1e10 - -2.25)))) - float64(0.0)) * 1.5) - 3.14159) * ((1.5 * -2.25) * 1.5))))
                            var v105 int = fi102
                            _ = fi103
                            _ = v104
                            _ = v105
                        }
                        v106 := GenS7{F0: fi102}
                        _ = fi102
                        _ = v106
                    }
                }
                _ = v100
            }
        }
    } else {
        if (((!false) || ((!(v2 >= v2)) && (true || (true && true)))) || true) {
            for fi107 := 0; fi107 < 18; fi107++ {
                v108 := GenS10{F0: 1, F1: "a"}
                _ = fi107
                _ = v108
            }
            fmt.Println(v7, v5)
        } else {
            v2 -= 2
            v7 = 0.0
            fmt.Println(v7, v7)
        }
    }
    switch -2147483648 {
    case 127:
        fmt.Println(v4, v5)
    case -9223372036854775808:
        v5 -= 65535
    case -5:
        v7 = 0.1
    }
    fmt.Println(v7)
    switch 128 {
    case -15:
        var v109 int8 = -128
        _ = v109
    case -1:
        delete(v6, "a")
    case 32768:
        for fi110 := 0; fi110 < 18; fi110++ {
            var v111 float64 = (float64((float32(1.5) + (v7 + float32(0.0)))) - (0.1 / (1e10 * 3.14159)))
            _ = fi110
            _ = v111
        }
    default:
        if (v14 == v14) {
            var v112 uint = (v4 + (v5 >> ((255) & 63)))
            switch -2147483648 {
            case -1:
                v1[128] = (("" + (("\n\t" + "\n\t") + "")) + "")
            case -950:
                fmt.Println(v4, v14, v2)
            }
            _ = v112
        } else {
            switch 351 {
            case 975:
                v1[4294967295] = ("" + "foo bar")
            case 128:
                v113 := GenS6{F0: 42}
                _ = v113
            default:
                var v114 uint64 = 577
                _ = v114
            }
            v4 = (v4 + v5)
            for fi115 := 0; fi115 < 5; fi115++ {
                v116 := GenS11{F0: 2, F1: 42, F2: 631}
                _ = fi115
                _ = v116
            }
        }
    }
    v4 = 255
    var v117 uint64 = 1
    v14 = (v14 & 1)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v117:", v117)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v117
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
