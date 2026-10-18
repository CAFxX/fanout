package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int8; F1 int }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint64; F1 int64; F2 int }
type GenS5 struct { F0 uint32; F1 uint64; F2 int }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 int32; F1 int64; F2 uint64 }
type GenS8 struct { F0 uint16; F1 uint; F2 string }
type GenS9 struct { F0 uint8 }
type GenS10 struct { F0 string }
type GenS11 struct { F0 int16 }
type GenS12 struct { F0 int16; F1 uint16; F2 int64 }
type GenS13 struct { F0 uint64 }
type GenS14 struct { F0 uint32; F1 int32 }
type GenS15 struct { F0 int }
type GenS16 struct { F0 uint8; F1 int16 }
type GenS17 struct { F0 uint16; F1 int32; F2 int8 }
type GenS18 struct { F0 bool; F1 string }
type GenS19 struct { F0 int64; F1 int; F2 int }
type GenS20 struct { F0 uint32; F1 int32 }
type GenS21 struct { F0 uint16; F1 uint64 }
type GenS22 struct { F0 int8; F1 uint16 }
func genFn0(p0 int16, p1 int32) int {
    var v7 uint16 = 32767
    var v8 int = -9223372036854775808
    if true {
        v7 += v7
        v9 := GenS1{F0: 747}
        v10 := GenS1{F0: 2}
        _ = v10
        _ = v9
    }
    var v11 uint32 = 2147483648
    _ = p0
    _ = p1
    _ = v11
    _ = v7
    _ = v8
    return v8
}

func main() {
    v1 := make(map[string]bool)
    var v2 int64 = 42
    var v3 uint32 = 2147483647
    v4 := make(map[uint32]uint32)
    var v5 uint = 18446744073709551615
    var v6 uint32 = (^(v3))
    var v12 int16 = 1
    var v13 int64 = v2
    v2 = v2
    for fi14 := 0; fi14 < 13; fi14++ {
        if (false || true) {
            v15 := []int{fi14, (fi14 * genFn0(v12, -702)), (fi14 * (fi14 >> ((1) & 63))), 9223372036854775807}
            switch "\n\t" {
            case "a":
                if (false || true) {
                    v1[("\n\t" + "")] = true
                } else {
                    for fi16 := 0; fi16 < 9; fi16++ {
                        delete(v4, v3)
                        v5 = v5
                        var v17 string = ""
                        if ((true || false) || (42 <= (fi16 >> ((170) & 63)))) {
                            if (42 == 814) {
                                switch fi14 {
                                case 127:
                                    fmt.Println(v3)
                                default:
                                    if ("a" >= "\n\t") {
                                        if (false || true) {
                                            switch genFn0(v12, -2147483648) {
                                            case 578:
                                                v5 = v5
                                            }
                                            _ = len(v1)
                                        }
                                        var v18 uint = (v5 << ((v5) & 63))
                                        _ = v18
                                    } else {
                                        v19 := GenS1{F0: v5}
                                        _ = v19
                                    }
                                }
                                v17 = v17
                                v17 = ("" + v17)
                            }
                        } else {
                            v15 = append(v15, 258)
                        }
                        if (42 >= -82) {
                            var v20 float64 = 1e10
                            if (!true) {
                                v3 += 711
                                v21 := GenS2{F0: 51, F1: fi16}
                                switch ((v17 + v17) + ("a" + ("a" + v17))) {
                                case "":
                                    _ = len(v4)
                                }
                                _ = v21
                            } else {
                                _ = len(v15)
                                v20 -= -2.25
                                if (v17 < v17) {
                                    if (127 <= -128) {
                                        v3 = uint32(v5)
                                        if ((!(v20 < (float64(float32((v20 + v20))) * 0.1))) && (true && false)) {
                                            var v22 uint16 = 255
                                            _ = v22
                                        }
                                    }
                                    if ((v12 - (v12 * v12)) < (v12 / ((v12) | 1))) {
                                        v13 = v13
                                        v23 := v1[((v17 + "\n\t") + "hello")]
                                        var v24 uint64 = 766
                                        _ = v23
                                        _ = v24
                                    }
                                    if (87 != 42) {
                                        if (("a" + (v17 + v17)) != v17) {
                                            if true {
                                                switch (fi14 & -1) {
                                                case 856:
                                                    var v25 int64 = -9223372036854775808
                                                    _ = v25
                                                case -171:
                                                    fmt.Println(fi14, v13)
                                                case 363:
                                                    _ = len(v15)
                                                    if false {
                                                        if len(v15) >= 2 {
                                                            v15 = v15[1:len(v15)]
                                                        }
                                                    }
                                                default:
                                                    fmt.Println(fi14, v6)
                                                    v26 := GenS3{F0: true}
                                                    _ = v26
                                                }
                                            } else {
                                                switch ("foo bar" + ("a" + ((v17 + v17) + v17))) {
                                                case "":
                                                    v27 := v1[("foo bar" + ((((v17 + "a") + v17) + "hello") + v17))]
                                                    _ = v27
                                                case "a":
                                                    switch "a" {
                                                    case "a":
                                                        v2 += v13
                                                    case "hello":
                                                        var v28 uint32 = 2147483648
                                                        if true {
                                                            if ((!(!(v17 >= v17))) && (127 >= 249)) {
                                                                v15 = append(v15, 986)
                                                                v6 += 903
                                                                v13 += int64(fi16)
                                                                v29 := GenS1{F0: 534}
                                                                _ = v29
                                                            } else {
                                                                switch (fi16 % (((fi14 / (((fi14 * fi16)) | 1))) | 1)) {
                                                                case -149:
                                                                    var v30 float64 = v20
                                                                    _ = v30
                                                                case -2147483648:
                                                                    fmt.Println(v20, v5, fi14)
                                                                case 2147483647:
                                                                    switch v17 {
                                                                    case "\n\t":
                                                                        switch ((((("hello" + "foo bar") + "foo bar") + ((v17 + ((v17 + "") + v17)) + v17)) + v17) + (v17 + ("\n\t" + "foo bar"))) {
                                                                        case "hello":
                                                                            v31 := GenS3{F0: false}
                                                                            _ = v31
                                                                        case "":
                                                                            v32 := GenS1{F0: v5}
                                                                            _ = v32
                                                                        default:
                                                                            var v33 string = v17
                                                                            _ = v33
                                                                        }
                                                                    case "foo bar":
                                                                        v3 += v3
                                                                    case "hello":
                                                                        switch ("hello" + "a") {
                                                                        case "foo bar":
                                                                            if ((3.14159 >= (float64((3.14159 - 3.14159)) - 3.14159)) && (true && ((v20 * v20) <= (v20 + v20)))) {
                                                                                var v34 int = (^(fi14))
                                                                                _ = v34
                                                                            }
                                                                        case "\n\t":
                                                                            switch -2147483648 {
                                                                            case 4294967295:
                                                                                switch "" {
                                                                                case "a":
                                                                                    var v35 int32 = 2147483647
                                                                                    _ = v35
                                                                                case "":
                                                                                    var v36 uint64 = 1
                                                                                    _ = v36
                                                                                case "foo bar":
                                                                                    v37 := GenS4{F0: 42, F1: (v13 % (((v13 & v2)) | 1)), F2: (fi14 << ((v5) & 63))}
                                                                                    _ = v37
                                                                                }
                                                                            case 127:
                                                                                if (2147483647 > 0) {
                                                                                    v2 = (v13 >> ((v5) & 63))
                                                                                    if false {
                                                                                        var v38 int16 = (^(v12))
                                                                                        _ = v38
                                                                                    } else {
                                                                                        var v39 uint = (v5 - (v5 & 2))
                                                                                        var v40 float64 = ((((((((v20 * v20) + v20) + ((1.5 * (v20 - ((3.14159 + v20) * v20))) + v20)) + (v20 / v20)) / (v20 - 1e10)) - (v20 + (float64(1.5) * v20))) + (v20 * 0.0)) / ((1e10 - float64((float32(1e10) - -2.25))) - v20))
                                                                                        _ = v39
                                                                                        _ = v40
                                                                                    }
                                                                                }
                                                                            }
                                                                        default:
                                                                            var v41 int32 = 42
                                                                            _ = v41
                                                                        }
                                                                    default:
                                                                        switch v17 {
                                                                        case "":
                                                                            switch (fi16 + (^(fi16))) {
                                                                            case -2147483648:
                                                                                v17 = ((("hello" + v17) + v17) + ((v17 + (v17 + ("" + v17))) + ""))
                                                                            }
                                                                        }
                                                                    }
                                                                }
                                                                delete(v4, (v3 / ((329) | 1)))
                                                            }
                                                            v17 = "hello"
                                                        }
                                                        if (-2147483648 > -535) {
                                                            v3 = uint32(v5)
                                                            if (!(("hello" + "foo bar") == v17)) {
                                                                if true {
                                                                    var v42 float64 = float64(float32((((v20 - 0.0) + (1.5 / (0.0 - ((0.1 + (((v20 + v20) * v20) * (v20 * v20))) + 1.5)))) - float64(1.5))))
                                                                    _ = v42
                                                                }
                                                                v43 := GenS5{F0: (v3 ^ 65535), F1: 2, F2: 127}
                                                                if true {
                                                                    switch "foo bar" {
                                                                    case "hello":
                                                                        switch 4294967295 {
                                                                        case 522:
                                                                            _ = len(v4)
                                                                        case -9223372036854775808:
                                                                            if false {
                                                                                switch v17 {
                                                                                case "a":
                                                                                    switch fi16 {
                                                                                    case -1:
                                                                                        var v44 uint32 = v28
                                                                                        _ = v44
                                                                                    case 1:
                                                                                        v12 += v12
                                                                                    default:
                                                                                        var v45 int = genFn0((^(v12)), -2147483648)
                                                                                        _ = v45
                                                                                    }
                                                                                default:
                                                                                    var v46 int64 = (v2 << (((v5 | 0)) & 63))
                                                                                    _ = v46
                                                                                }
                                                                                var v47 int8 = -1
                                                                                _ = v47
                                                                            } else {
                                                                                _ = len(v4)
                                                                            }
                                                                        default:
                                                                            v15 = append(v15, fi14)
                                                                            v4[v6] = (v3 >> ((v5) & 63))
                                                                            switch v17 {
                                                                            case "a":
                                                                                v48 := GenS4{F0: 2, F1: (v13 * 0), F2: (^(fi14))}
                                                                                v20 = v20
                                                                                var v49 float64 = 3.14159
                                                                                _ = v48
                                                                                _ = v49
                                                                            }
                                                                        }
                                                                    case "":
                                                                        var v50 int64 = -1
                                                                        _ = v50
                                                                    }
                                                                } else {
                                                                    if (255 == (v5 * v5)) {
                                                                        if ((!((!false) && (!(true && false)))) || (true && false)) {
                                                                            var v51 int16 = v12
                                                                            _ = v51
                                                                        } else {
                                                                            v20 -= -2.25
                                                                        }
                                                                        if (!false) {
                                                                            v52 := GenS5{F0: v3, F1: 4294967295, F2: genFn0(v12, 1)}
                                                                            _ = v52
                                                                        } else {
                                                                            v53 := GenS3{F0: false}
                                                                            switch (^(fi16)) {
                                                                            case 128:
                                                                                v5 = v5
                                                                            case -1:
                                                                                if (279 <= 255) {
                                                                                    var v54 int8 = 1
                                                                                    fmt.Println(v20, v28, v17)
                                                                                    v2 = (-(v13))
                                                                                    _ = v54
                                                                                } else {
                                                                                    v2 += v2
                                                                                    var v55 string = v17
                                                                                    _ = v55
                                                                                }
                                                                            case -2147483648:
                                                                                v28 = (v6 * (^(v28)))
                                                                            default:
                                                                                v15 = append(v15, (fi16 | (-(fi14))))
                                                                            }
                                                                            v56 := GenS6{F0: 135}
                                                                            _ = v53
                                                                            _ = v56
                                                                        }
                                                                        fmt.Println(v28, v2)
                                                                    }
                                                                }
                                                                switch v17 {
                                                                case "hello":
                                                                    switch genFn0(v12, -110) {
                                                                    case 462:
                                                                        if true {
                                                                            v12 -= v12
                                                                            var v57 float32 = (1.5 / (3.14159 + (0.0 / 3.14159)))
                                                                            v5 -= (^(v5))
                                                                            _ = v57
                                                                        }
                                                                    case 127:
                                                                        v15 = append(v15, (fi16 | fi16))
                                                                    default:
                                                                        var v58 string = v17
                                                                        var v59 uint = 838
                                                                        _ = v58
                                                                        _ = v59
                                                                    }
                                                                case "a":
                                                                    var v60 int16 = (v12 | (^(v12)))
                                                                    _ = v60
                                                                case "foo bar":
                                                                    if len(v15) >= 2 {
                                                                        v15 = v15[1:len(v15)]
                                                                    }
                                                                default:
                                                                    v61 := GenS7{F0: 0, F1: v13, F2: 1}
                                                                    _ = v61
                                                                }
                                                                _ = v43
                                                            } else {
                                                                v62 := GenS7{F0: -1, F1: (-(v13)), F2: 2}
                                                                _ = v62
                                                            }
                                                            var v63 string = ""
                                                            _ = v63
                                                        }
                                                        _ = v28
                                                    case "":
                                                        if true {
                                                            v15 = append(v15, fi14)
                                                            v20 -= (v20 / v20)
                                                            v64 := GenS6{F0: 78}
                                                            _ = v64
                                                        } else {
                                                            fmt.Println(v12, v2, v12)
                                                            fmt.Println(v17)
                                                            switch (fi16 | genFn0((^(v12)), 2147483647)) {
                                                            case 2147483647:
                                                                fmt.Println(v5)
                                                            case 0:
                                                                switch 128 {
                                                                case -9223372036854775808:
                                                                    if true {
                                                                        v65 := GenS6{F0: 32767}
                                                                        _ = v65
                                                                    }
                                                                default:
                                                                    v66 := GenS2{F0: -128, F1: fi14}
                                                                    _ = v66
                                                                }
                                                            case 889:
                                                                fmt.Println(fi16, v12, v17)
                                                            default:
                                                                var v67 uint16 = uint16(v5)
                                                                _ = v67
                                                            }
                                                        }
                                                    default:
                                                        var v68 int64 = 597
                                                        _ = v68
                                                    }
                                                    v20 = float64(0.1)
                                                    switch "foo bar" {
                                                    case "":
                                                        v3 += (^(v6))
                                                    case "\n\t":
                                                        if (v20 != v20) {
                                                            v69 := GenS2{F0: 1, F1: int(v5)}
                                                            switch (-(fi14)) {
                                                            case 42:
                                                                v70 := GenS8{F0: 42, F1: v5, F2: v17}
                                                                _ = v70
                                                            default:
                                                                var v71 uint8 = 1
                                                                switch v17 {
                                                                case "hello":
                                                                    v2 += v2
                                                                case "":
                                                                    if true {
                                                                        switch fi14 {
                                                                        case 996:
                                                                            var v72 uint = v5
                                                                            _ = v72
                                                                        case -9223372036854775808:
                                                                            var v73 int64 = int64(v6)
                                                                            _ = v73
                                                                        case 9223372036854775807:
                                                                            if false {
                                                                                fmt.Println(v12, v20)
                                                                            } else {
                                                                                switch "foo bar" {
                                                                                case "":
                                                                                    v74 := v4[(^(v6))]
                                                                                    _ = v74
                                                                                case "foo bar":
                                                                                    if (!(1.5 <= (3.14159 * v20))) {
                                                                                        if true {
                                                                                            v20 = v20
                                                                                            var v75 bool = true
                                                                                            var v76 float64 = v20
                                                                                            _ = v75
                                                                                            _ = v76
                                                                                        }
                                                                                        v71 -= v71
                                                                                        var v77 uint8 = (^(v71))
                                                                                        _ = v77
                                                                                    }
                                                                                }
                                                                                if true {
                                                                                    switch v17 {
                                                                                    case "hello":
                                                                                        var v78 int = fi16
                                                                                        _ = v78
                                                                                    }
                                                                                    switch v17 {
                                                                                    case "a":
                                                                                        var v79 int16 = 32767
                                                                                        _ = v79
                                                                                    default:
                                                                                        v80 := GenS1{F0: (v5 - 42)}
                                                                                        _ = v80
                                                                                    }
                                                                                    v81 := GenS9{F0: v71}
                                                                                    _ = v81
                                                                                } else {
                                                                                    switch v17 {
                                                                                    case "foo bar":
                                                                                        fmt.Println(v6, v71, v5)
                                                                                    default:
                                                                                        var v82 int16 = 412
                                                                                        _ = v82
                                                                                    }
                                                                                    if true {
                                                                                        fmt.Println(v2, v71)
                                                                                        if false {
                                                                                            if ((0.0 <= ((0.0 * (float32(1.5) - (3.14159 / (-2.25 / 1e10)))) + (float32(v20) * 1.5))) && false) {
                                                                                                if false {
                                                                                                    delete(v1, "")
                                                                                                    var v83 uint8 = (v71 << ((v5) & 63))
                                                                                                    _ = v83
                                                                                                } else {
                                                                                                    v6 += 4294967295
                                                                                                    var v84 float64 = (3.14159 - v20)
                                                                                                    fmt.Println(v84, v17, v17)
                                                                                                    _ = v84
                                                                                                }
                                                                                                var v85 uint = 224
                                                                                                _ = v85
                                                                                            }
                                                                                            var v86 uint64 = 4294967295
                                                                                            v2 = (v13 - v13)
                                                                                            _ = v86
                                                                                        }
                                                                                        fmt.Println(fi14)
                                                                                    } else {
                                                                                        if false {
                                                                                            if (int8(v6) < 100) {
                                                                                                var v87 uint32 = (v6 % (((^(v6))) | 1))
                                                                                                var v88 int = (-(fi14))
                                                                                                delete(v1, ("\n\t" + "foo bar"))
                                                                                                _ = v87
                                                                                                _ = v88
                                                                                            } else {
                                                                                                v89 := v4[51]
                                                                                                v90 := GenS1{F0: v5}
                                                                                                fmt.Println(v2, v12, fi14)
                                                                                                _ = v89
                                                                                                _ = v90
                                                                                            }
                                                                                            if len(v15) > 0 {
                                                                                                v15[0] = (fi14 & 128)
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                    if (true || (v20 < ((v20 * float64(0.0)) + float64(((float32(0.1) - ((float32(1.5) + 1.5) + -2.25)) - 0.1))))) {
                                                                                        v91 := GenS10{F0: "hello"}
                                                                                        var v92 int8 = -16
                                                                                        _ = v91
                                                                                        _ = v92
                                                                                    } else {
                                                                                        v93 := GenS11{F0: (v12 & v12)}
                                                                                        _ = v93
                                                                                    }
                                                                                    v5 = (^(v5))
                                                                                }
                                                                                var v94 int8 = 127
                                                                                _ = v94
                                                                            }
                                                                        }
                                                                        _ = len(v4)
                                                                    } else {
                                                                        fmt.Println(v20, fi14, v3)
                                                                    }
                                                                default:
                                                                    if false {
                                                                        if false {
                                                                            fmt.Println(v13, v3)
                                                                            var v95 int8 = 0
                                                                            v71 -= 205
                                                                            _ = v95
                                                                        } else {
                                                                            if true {
                                                                                var v96 uint32 = (^(v6))
                                                                                _ = v96
                                                                            }
                                                                            v97 := GenS2{F0: 42, F1: (fi16 / ((fi16) | 1))}
                                                                            _ = v97
                                                                        }
                                                                        fmt.Println(v71)
                                                                        var v98 uint = v5
                                                                        _ = v98
                                                                    }
                                                                }
                                                                if false {
                                                                    var v99 float32 = 3.14159
                                                                    v17 = (v17 + "hello")
                                                                    v1[v17] = (0.0 <= float64(v99))
                                                                    _ = v99
                                                                } else {
                                                                    if true {
                                                                        v6 = v6
                                                                        v5 = (v5 ^ (v5 + 994))
                                                                        var v100 int16 = v12
                                                                        _ = v100
                                                                    } else {
                                                                        v6 += v3
                                                                        if len(v15) >= 2 {
                                                                            v15 = v15[1:len(v15)]
                                                                        }
                                                                    }
                                                                    fmt.Println(v71, v5, v6)
                                                                }
                                                                _ = v71
                                                            }
                                                            _ = v69
                                                        } else {
                                                            v5 = v5
                                                            var v101 string = (v17 + "\n\t")
                                                            v102 := GenS1{F0: (v5 << (((v5 >> ((65535) & 63))) & 63))}
                                                            _ = v101
                                                            _ = v102
                                                        }
                                                    default:
                                                        var v103 int8 = 1
                                                        _ = v103
                                                    }
                                                    fmt.Println(fi14, v13)
                                                    var v104 uint16 = 89
                                                    v105 := GenS1{F0: v5}
                                                    v2 = v13
                                                    if true {
                                                        var v106 int8 = -8
                                                        _ = v106
                                                    }
                                                    var v107 float32 = ((((float32(v20) - 1e10) - (0.0 + float32(v20))) * (-2.25 * -2.25)) * 0.1)
                                                    _ = v104
                                                    _ = v105
                                                    _ = v107
                                                }
                                            }
                                            if (!((!(true || false)) || (true && (228 <= 42)))) {
                                                switch (-(fi14)) {
                                                case 42:
                                                    v13 -= -707
                                                case 9223372036854775807:
                                                    var v108 int = fi16
                                                    _ = v108
                                                }
                                                v20 = (v20 - 3.14159)
                                            } else {
                                                v12 -= int16(fi16)
                                                v20 = 1.5
                                                var v109 int64 = int64(v3)
                                                _ = v109
                                            }
                                        }
                                        v110 := GenS8{F0: 2, F1: 189, F2: ((((("" + "\n\t") + v17) + v17) + (("\n\t" + v17) + (((v17 + v17) + ("\n\t" + v17)) + (((v17 + (v17 + v17)) + (v17 + "hello")) + v17)))) + "a")}
                                        delete(v4, (^(v6)))
                                        _ = v110
                                    }
                                }
                            }
                            _ = v20
                        } else {
                            if ((^(v13)) > (v13 | v2)) {
                                _ = len(v15)
                                var v111 int8 = -1
                                _ = v111
                            }
                            v112 := v4[(v3 ^ (v3 >> (((v5 * 92)) & 63)))]
                            var v113 uint64 = 18446744073709551615
                            _ = v112
                            _ = v113
                        }
                        _ = fi16
                        _ = v17
                    }
                }
            default:
                for fi114 := 0; fi114 < 20; fi114++ {
                    if false {
                        fmt.Println(v6)
                        fmt.Println(v5, v6, fi114)
                        switch fi114 {
                        case -2147483648:
                            v115 := GenS12{F0: int16(fi14), F1: 199, F2: v13}
                            _ = v115
                        }
                    } else {
                        v13 -= (v13 | 0)
                        if false {
                            v116 := GenS13{F0: 18446744073709551615}
                            switch (^(fi14)) {
                            case 32767:
                                var v117 int = genFn0((v12 - v12), -305)
                                _ = v117
                            }
                            _ = v116
                        } else {
                            v118 := GenS10{F0: "hello"}
                            _ = v118
                        }
                        fmt.Println(v13, v3)
                        if false {
                            v119 := GenS3{F0: false}
                            if (fi114 >= fi14) {
                                v5 -= v5
                                if false {
                                    var v120 int32 = -2147483648
                                    if (!((v5 / (((v5 ^ v5)) | 1)) > v5)) {
                                        var v121 int16 = v12
                                        _ = v121
                                    }
                                    if true {
                                        v4[1] = (v6 * 93)
                                        v120 -= v120
                                        v12 = 1000
                                    } else {
                                        var v122 uint64 = 0
                                        fmt.Println(v2, v13, fi14)
                                        if (false || false) {
                                            v123 := GenS14{F0: (v6 / ((2147483648) | 1)), F1: (v120 ^ 0)}
                                            v124 := GenS15{F0: (fi114 - fi14)}
                                            _ = v123
                                            _ = v124
                                        } else {
                                            if false {
                                                var v125 int16 = (v12 - (v12 >> (((v5 & (v5 >> ((v5) & 63)))) & 63)))
                                                if true {
                                                    v126 := GenS1{F0: 422}
                                                    var v127 uint64 = (^(v122))
                                                    if (-2147483648 <= v120) {
                                                        var v128 uint = (v5 | v5)
                                                        fmt.Println(v120, v125, v2)
                                                        var v129 int = (fi114 / ((int(v13)) | 1))
                                                        _ = v128
                                                        _ = v129
                                                    } else {
                                                        delete(v1, "foo bar")
                                                        switch fi14 {
                                                        case 4294967295:
                                                            var v130 uint16 = 1
                                                            _ = v130
                                                        case -9223372036854775808:
                                                            v1[("hello" + "foo bar")] = false
                                                        }
                                                        v131 := GenS2{F0: -33, F1: 127}
                                                        delete(v4, v3)
                                                        _ = v131
                                                    }
                                                    _ = v126
                                                    _ = v127
                                                } else {
                                                    var v132 int = fi114
                                                    _ = v132
                                                }
                                                _ = v125
                                            }
                                            var v133 int8 = 1
                                            delete(v4, (v6 - 320))
                                            switch int(v133) {
                                            case 375:
                                                switch fi14 {
                                                case 128:
                                                    var v134 string = ("hello" + ("" + "foo bar"))
                                                    _ = v134
                                                case -600:
                                                    v3 += v3
                                                }
                                            case 302:
                                                v120 -= v120
                                            }
                                            _ = v133
                                        }
                                        _ = v122
                                    }
                                    var v135 int = (fi114 & 128)
                                    if true {
                                        v12 -= 0
                                        v136 := v1["a"]
                                        _ = v136
                                    }
                                    v2 += v13
                                    switch ("hello" + (("" + "a") + (("a" + ("\n\t" + ("" + "foo bar"))) + "foo bar"))) {
                                    case "foo bar":
                                        v137 := GenS16{F0: 42, F1: (-(v12))}
                                        _ = v137
                                    case "a":
                                        if true {
                                            var v138 bool = (1 == 255)
                                            _ = v138
                                        } else {
                                            if (120 <= 2) {
                                                fmt.Println(v135)
                                                if true {
                                                    if len(v15) >= 2 {
                                                        v15 = v15[1:len(v15)]
                                                    }
                                                    v12 -= v12
                                                } else {
                                                    var v139 int16 = (v12 ^ v12)
                                                    v139 = v139
                                                    if (0.0 < ((float32(-2.25) * 1e10) + 1.5)) {
                                                        switch ("" + "foo bar") {
                                                        case "":
                                                            v15 = append(v15, 0)
                                                        case "\n\t":
                                                            switch ("hello" + "foo bar") {
                                                            case "a":
                                                                if false {
                                                                    if true {
                                                                        var v140 int16 = (v139 + (v139 % ((v12) | 1)))
                                                                        v140 = (v12 >> ((v5) & 63))
                                                                        _ = v140
                                                                    }
                                                                    var v141 int16 = (v139 ^ -32768)
                                                                    fmt.Println(v12, v3)
                                                                    _ = v141
                                                                } else {
                                                                    if len(v15) > 0 {
                                                                        v15[0] = genFn0(32767, 0)
                                                                    }
                                                                }
                                                            case "foo bar":
                                                                var v142 int16 = -32768
                                                                _ = v142
                                                            }
                                                        default:
                                                            v120 = (v120 | v120)
                                                        }
                                                    } else {
                                                        var v143 uint = v5
                                                        _ = v143
                                                    }
                                                    if len(v15) >= 2 {
                                                        v15 = v15[1:len(v15)]
                                                    }
                                                    fmt.Println(v6, v13, v5)
                                                    _ = v139
                                                }
                                            } else {
                                                v144 := GenS4{F0: 785, F1: -957, F2: (fi114 & (^(v135)))}
                                                _ = v144
                                            }
                                            var v145 float64 = 0.0
                                            switch ("" + "a") {
                                            case "hello":
                                                fmt.Println(fi114, v135)
                                            default:
                                                if len(v15) > 0 {
                                                    v15[0] = genFn0(0, (-(v120)))
                                                }
                                            }
                                            _ = v145
                                        }
                                    default:
                                        if false {
                                            if len(v15) >= 2 {
                                                v15 = v15[1:len(v15)]
                                            }
                                            var v146 uint32 = 812
                                            _ = v146
                                        }
                                    }
                                    _ = v120
                                    _ = v135
                                } else {
                                    v12 -= -373
                                    if true {
                                        fmt.Println(fi114, v3, fi114)
                                        v12 += v12
                                        var v147 string = (((("" + "\n\t") + "a") + (("" + "\n\t") + "\n\t")) + ("hello" + "a"))
                                        v12 -= v12
                                        _ = v147
                                    } else {
                                        v2 -= 1
                                        var v148 int16 = (v12 << (((v5 * (v5 | 679))) & 63))
                                        v149 := GenS9{F0: 48}
                                        v150 := GenS9{F0: 1}
                                        _ = v148
                                        _ = v149
                                        _ = v150
                                    }
                                    switch ("hello" + ((("\n\t" + ("\n\t" + "")) + ((("\n\t" + "") + ("a" + "hello")) + "hello")) + "a")) {
                                    case "":
                                        v151 := GenS6{F0: 472}
                                        _ = v151
                                    case "a":
                                        switch fi14 {
                                        case 42:
                                            if (!true) {
                                                var v152 uint16 = 32767
                                                _ = v152
                                            } else {
                                                v1["\n\t"] = (994 < (fi14 / (((^(fi14))) | 1)))
                                                var v153 float64 = 1.5
                                                if (0 > v5) {
                                                    v3 += v3
                                                    v154 := GenS17{F0: uint16(v5), F1: 0, F2: 42}
                                                    _ = v154
                                                }
                                                _ = v153
                                            }
                                        }
                                    }
                                    v155 := GenS14{F0: (v3 | 4294967295), F1: 0}
                                    fmt.Println(v5)
                                    _ = v155
                                }
                                fmt.Println(v12, v12, v2)
                            }
                            _ = v119
                        } else {
                            if false {
                                switch "hello" {
                                case "\n\t":
                                    switch "a" {
                                    case "\n\t":
                                        var v156 uint = v5
                                        _ = v156
                                    case "foo bar":
                                        var v157 int8 = 127
                                        _ = v157
                                    default:
                                        v15 = append(v15, fi14)
                                        var v158 uint = (v5 + v5)
                                        var v159 uint = v158
                                        var v160 uint16 = 1
                                        v161 := GenS18{F0: false, F1: "hello"}
                                        _ = v158
                                        _ = v159
                                        _ = v160
                                        _ = v161
                                    }
                                default:
                                    _ = len(v15)
                                }
                                var v162 int16 = 1
                                _ = v162
                            }
                            switch (fi114 - genFn0((v12 & (v12 | v12)), -371)) {
                            case 50:
                                if (false && (42 > -764)) {
                                    v163 := GenS5{F0: (v3 ^ 2147483648), F1: 18446744073709551615, F2: fi114}
                                    switch "\n\t" {
                                    case "foo bar":
                                        v164 := GenS19{F0: 0, F1: fi114, F2: genFn0(562, -403)}
                                        _ = v164
                                    case "":
                                        v6 = 2
                                    case "hello":
                                        if len(v15) > 0 {
                                            v15[0] = fi14
                                        }
                                    }
                                    _ = v163
                                }
                            case -1:
                                var v165 float32 = 0.1
                                _ = v165
                            }
                            v166 := GenS12{F0: (v12 / ((int16(v6)) | 1)), F1: 42, F2: (v2 & -1)}
                            var v167 int16 = (v12 << ((v5) & 63))
                            v168 := GenS18{F0: true, F1: (("" + "\n\t") + ("hello" + ("" + "\n\t")))}
                            _ = v166
                            _ = v167
                            _ = v168
                        }
                    }
                    _ = fi114
                }
            }
            _ = v15
        } else {
            if true {
                v169 := GenS20{F0: 42, F1: -2147483648}
                v3 = 985
                var v170 float32 = 1.5
                _ = v169
                _ = v170
            } else {
                v12 -= v12
                v12 += (v12 / ((95) | 1))
            }
            var v171 float32 = -2.25
            _ = v171
        }
        _ = fi14
    }
    switch "\n\t" {
    case "hello":
        v172 := GenS21{F0: 255, F1: 1}
        _ = v172
    case "":
        if (!(v2 >= v13)) {
            delete(v1, "hello")
            fmt.Println(v3, v13, v3)
            var v173 int64 = v13
            _ = v173
        }
    case "foo bar":
        v1["\n\t"] = (!false)
    }
    if (!false) {
        var v174 uint16 = 1
        var v175 float64 = 1.5
        if (!true) {
            var v176 uint32 = v6
            var v177 int16 = v12
            _ = v176
            _ = v177
        }
        _ = v174
        _ = v175
    }
    if (-128 > -80) {
        v178 := GenS8{F0: 1000, F1: 18446744073709551615, F2: "foo bar"}
        var v179 uint16 = 65535
        _ = v178
        _ = v179
    }
    for fi180 := 0; fi180 < 14; fi180++ {
        var v181 uint64 = 4294967295
        _ = fi180
        _ = v181
    }
    v182 := GenS20{F0: v3, F1: -1}
    fmt.Println(v3, v2)
    if true {
        for fi183 := 0; fi183 < 15; fi183++ {
            if ((100 >= 42) || true) {
                v13 += (^(v13))
                switch genFn0(v12, 1) {
                case -9223372036854775808:
                    v184 := GenS22{F0: 127, F1: 771}
                    _ = v184
                default:
                    v6 -= v3
                }
            } else {
                for fi185 := 0; fi185 < 10; fi185++ {
                    var v186 uint64 = 0
                    v186 = v186
                    _ = fi185
                    _ = v186
                }
                if (0.1 > 3.14159) {
                    var v187 uint16 = 1
                    _ = v187
                }
            }
            v188 := []int{}
            for fi189 := 0; fi189 < 8; fi189++ {
                var v190 uint8 = 2
                _ = fi189
                _ = v190
            }
            _ = fi183
            _ = v188
        }
        v191 := []int{256, -611, 546}
        _ = v191
    }
    switch ("" + ("a" + ((("a" + ("a" + "hello")) + ((("\n\t" + "foo bar") + ("foo bar" + "hello")) + ("" + "foo bar"))) + ("\n\t" + "a")))) {
    case "hello":
        var v192 int32 = -2147483648
        _ = v192
    }
    v193 := GenS16{F0: 2, F1: v12}
    var v194 int64 = 9223372036854775807
    var v195 int16 = (v12 ^ -1)
    v194 += v13
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v194:", v194)
    fmt.Println("v195:", v195)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v182
    _ = v193
    _ = v194
    _ = v195
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
