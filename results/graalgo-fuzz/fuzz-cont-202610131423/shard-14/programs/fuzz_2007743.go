package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string }
type GenS2 struct { F0 int32; F1 uint8; F2 string }
type GenS3 struct { F0 int8; F1 int; F2 uint64 }
type GenS4 struct { F0 uint; F1 uint64 }
type GenS5 struct { F0 uint8; F1 uint }
type GenS6 struct { F0 bool; F1 int64; F2 uint32 }
type GenS7 struct { F0 uint32; F1 string; F2 uint8 }
type GenS8 struct { F0 uint32 }
type GenS9 struct { F0 int8; F1 int16 }
type GenS10 struct { F0 bool }
type GenS11 struct { F0 int64; F1 bool }
type GenS12 struct { F0 bool }
type GenS13 struct { F0 uint32 }
type GenS14 struct { F0 uint16; F1 bool }
type GenS15 struct { F0 int8; F1 uint16; F2 int32 }
type GenS16 struct { F0 int }
type GenS17 struct { F0 int16; F1 int64 }
type GenS18 struct { F0 uint64; F1 int8 }
type GenS19 struct { F0 int32 }

func main() {
    v1 := make(map[string]uint32)
    var v2 uint = 1
    var v3 uint8 = 97
    var v4 uint = v2
    var v5 int64 = -1
    var v6 uint64 = 878
    v5 = (^(v5))
    var v7 string = ""
    if false {
        v8 := GenS1{F0: 200, F1: ("" + (v7 + ((((("hello" + v7) + (v7 + "hello")) + "a") + v7) + v7)))}
        v2 = v4
        switch 667 {
        case 9223372036854775807:
            if false {
                for fi9 := 0; fi9 < 19; fi9++ {
                    var v10 int = 128
                    v11 := []int{-517, int(v5), (^(v10)), 128}
                    _ = fi9
                    _ = v10
                    _ = v11
                }
            } else {
                v12 := GenS1{F0: 42, F1: v7}
                if (!true) {
                    if (((255 >= v3) && true) && true) {
                        v2 += v4
                    }
                    v6 += 42
                    v5 += (^(v5))
                } else {
                    v13 := GenS2{F0: 0, F1: 127, F2: v7}
                    _ = v13
                }
                _ = v12
            }
        case 127:
            switch 128 {
            case 32767:
                for fi14 := 0; fi14 < 19; fi14++ {
                    for fi15 := 0; fi15 < 17; fi15++ {
                        v16 := GenS2{F0: int32(v3), F1: v3, F2: ("foo bar" + "\n\t")}
                        v17 := []int{(fi14 % (((fi14 * 4294967295)) | 1))}
                        if (((v3 >= (v3 % ((v3) | 1))) || (true && (2 >= uint32(v2)))) && ((42 > 1) && false)) {
                            _ = len(v17)
                            var v18 int8 = 1
                            v3 = (v3 << (((v4 << ((v2) & 63))) & 63))
                            var v19 uint16 = 255
                            _ = v18
                            _ = v19
                        } else {
                            delete(v1, (((v7 + (v7 + v7)) + v7) + (v7 + v7)))
                        }
                        switch "\n\t" {
                        case "hello":
                            fmt.Println(v7, v5, v4)
                        case "":
                            v4 += v4
                        default:
                            v5 = (v5 ^ -273)
                        }
                        _ = fi15
                        _ = v16
                        _ = v17
                    }
                    _ = fi14
                }
            case -2147483648:
                for fi20 := 0; fi20 < 10; fi20++ {
                    if true {
                        if false {
                            switch fi20 {
                            case 525:
                                v21 := []int{(fi20 | 42), (fi20 << ((4294967295) & 63)), (fi20 % ((fi20) | 1))}
                                _ = v21
                            case -352:
                                fmt.Println(v3, v2, fi20)
                            case -128:
                                for fi22 := 0; fi22 < 12; fi22++ {
                                    if (fi22 == 505) {
                                        var v23 string = (("\n\t" + v7) + ((v7 + "\n\t") + v7))
                                        _ = v23
                                    } else {
                                        if false {
                                            var v24 int = fi20
                                            _ = v24
                                        }
                                        v25 := GenS2{F0: -2147483648, F1: (^(v3)), F2: "\n\t"}
                                        _ = v25
                                    }
                                    v26 := GenS1{F0: v5, F1: v7}
                                    if true {
                                        if ((1.5 + 0.1) < float64(float32(1.5))) {
                                            v27 := GenS3{F0: 1, F1: (fi20 % ((fi22) | 1)), F2: (^(v6))}
                                            v7 = v7
                                            _ = v27
                                        } else {
                                            fmt.Println(fi22, v4, fi20)
                                            v28 := []int{fi20, fi22, (^(fi22))}
                                            _ = v28
                                        }
                                        v3 += uint8(v4)
                                        var v29 uint = (v4 & v2)
                                        v30 := GenS4{F0: (v29 % ((v4) | 1)), F1: 550}
                                        switch (fi20 & (fi20 - fi20)) {
                                        case -247:
                                            v29 = v29
                                        case 128:
                                            fmt.Println(fi22)
                                            if true {
                                                if (1 <= 1) {
                                                    v31 := []int{(fi22 | fi20), fi22}
                                                    if (v7 > "hello") {
                                                        var v32 bool = true
                                                        if len(v31) > 0 {
                                                            v31[0] = (fi20 * int(v2))
                                                        }
                                                        switch (v7 + v7) {
                                                        case "":
                                                            v1["\n\t"] = 2147483647
                                                        case "hello":
                                                            v7 = ("foo bar" + v7)
                                                        default:
                                                            var v33 int8 = -1
                                                            fmt.Println(v33, v2)
                                                            v2 -= v2
                                                            v5 = 1
                                                            _ = v33
                                                        }
                                                        _ = v32
                                                    } else {
                                                        fmt.Println(v5, v29)
                                                        if ((32768 > (fi22 - (fi20 & fi20))) || (56 == -128)) {
                                                            var v34 float32 = ((1.5 / -2.25) + 3.14159)
                                                            if (v6 <= v6) {
                                                                v35 := GenS1{F0: (v5 << (((v2 ^ v4)) & 63)), F1: "\n\t"}
                                                                fmt.Println(v7, v3)
                                                                var v36 float32 = -2.25
                                                                v36 -= ((1e10 * float32(0.1)) * (v36 - v36))
                                                                _ = v35
                                                                _ = v36
                                                            } else {
                                                                fmt.Println(v4, v2)
                                                                var v37 int64 = (-(v5))
                                                                if false {
                                                                    var v38 int = (^(fi22))
                                                                    var v39 int64 = 725
                                                                    switch "foo bar" {
                                                                    case "\n\t":
                                                                        if ((v4 / (((^(v29))) | 1)) == (^(v2))) {
                                                                            fmt.Println(v39)
                                                                        }
                                                                    case "":
                                                                        var v40 int = -1
                                                                        _ = v40
                                                                    default:
                                                                        v41 := GenS4{F0: v2, F1: (v6 << ((v4) & 63))}
                                                                        _ = v41
                                                                    }
                                                                    _ = v38
                                                                    _ = v39
                                                                }
                                                                _ = v37
                                                            }
                                                            if true {
                                                                switch (v7 + v7) {
                                                                case "a":
                                                                    switch "hello" {
                                                                    case "hello":
                                                                        v29 = 0
                                                                    case "":
                                                                        v42 := GenS4{F0: (v2 ^ v4), F1: (^(v6))}
                                                                        _ = v42
                                                                    case "\n\t":
                                                                        var v43 float32 = ((float32(1.5) - ((((((v34 - v34) / (v34 / float32((float64(1.5) - (0.1 - (0.1 / -2.25)))))) / 1.0) / 1e10) / 1.5) / 1e10)) * v34)
                                                                        var v44 uint = (v2 | (v2 & 354))
                                                                        switch v7 {
                                                                        case "hello":
                                                                            fmt.Println(v29)
                                                                            var v45 int64 = v5
                                                                            if true {
                                                                                v3 = (v3 ^ 2)
                                                                            }
                                                                            switch (v7 + v7) {
                                                                            case "hello":
                                                                                var v46 uint8 = v3
                                                                                _ = v46
                                                                            }
                                                                            fmt.Println(v5, v3)
                                                                            var v47 uint8 = (v3 - (v3 * 128))
                                                                            if false {
                                                                                v3 = v47
                                                                                if (-128 <= -91) {
                                                                                    if (v47 >= (v47 / (((^(v47))) | 1))) {
                                                                                        v48 := GenS2{F0: 1, F1: v47, F2: v7}
                                                                                        v29 += 491
                                                                                        switch fi20 {
                                                                                        case 977:
                                                                                            if true {
                                                                                                v49 := v1[((v7 + "hello") + v7)]
                                                                                                var v50 uint8 = (v47 << (((v29 + (v29 / (((v44 << ((2) & 63))) | 1)))) & 63))
                                                                                                _ = v49
                                                                                                _ = v50
                                                                                            } else {
                                                                                                fmt.Println(v3, v47)
                                                                                                v3 = (v3 >> ((v2) & 63))
                                                                                                var v51 bool = false
                                                                                                _ = v51
                                                                                            }
                                                                                        case -317:
                                                                                            fmt.Println(v5, v4)
                                                                                        default:
                                                                                            switch v7 {
                                                                                            case "a":
                                                                                                v44 += uint(v3)
                                                                                            case "foo bar":
                                                                                                if false {
                                                                                                    if (v7 < v7) {
                                                                                                        var v52 float64 = float64(((1e10 + float32(-2.25)) / (0.0 - (v43 * (0.0 * (v34 - v34))))))
                                                                                                        var v53 int16 = -231
                                                                                                        v44 = (^(v44))
                                                                                                        _ = v52
                                                                                                        _ = v53
                                                                                                    } else {
                                                                                                        fmt.Println(v4, v34)
                                                                                                        v54 := GenS3{F0: 8, F1: (-(fi20)), F2: v6}
                                                                                                        _ = v54
                                                                                                    }
                                                                                                    var v55 int64 = (v5 | 42)
                                                                                                    v56 := GenS5{F0: 1, F1: 2}
                                                                                                    _ = v55
                                                                                                    _ = v56
                                                                                                } else {
                                                                                                    if (360 > 2) {
                                                                                                        fmt.Println(fi20)
                                                                                                        v1[(v7 + v7)] = 4294967295
                                                                                                    }
                                                                                                    v6 += v6
                                                                                                }
                                                                                                v5 -= v45
                                                                                                v31 = append(v31, (fi22 << (((v2 | (v44 / ((v44) | 1)))) & 63)))
                                                                                            case "hello":
                                                                                                v4 = v2
                                                                                            }
                                                                                            v57 := GenS5{F0: uint8(v6), F1: (v29 & (v29 * v44))}
                                                                                            _ = v57
                                                                                        }
                                                                                        _ = v48
                                                                                    }
                                                                                    if true {
                                                                                        _ = len(v1)
                                                                                        var v58 bool = (true || (!false))
                                                                                        var v59 float32 = v34
                                                                                        _ = v58
                                                                                        _ = v59
                                                                                    } else {
                                                                                        if (32767 <= 65535) {
                                                                                            v7 = ""
                                                                                            var v60 int = 42
                                                                                            _ = v60
                                                                                        } else {
                                                                                            v31 = append(v31, 0)
                                                                                            delete(v1, "hello")
                                                                                        }
                                                                                        var v61 float64 = 0.0
                                                                                        _ = v61
                                                                                    }
                                                                                } else {
                                                                                    var v62 int8 = -43
                                                                                    if false {
                                                                                        switch -547 {
                                                                                        case -2147483648:
                                                                                            v7 = "foo bar"
                                                                                        case 2147483647:
                                                                                            switch fi22 {
                                                                                            case 600:
                                                                                                if ((!(!false)) || ((^(v47)) < (^(v47)))) {
                                                                                                    if (!(2147483648 > 2147483647)) {
                                                                                                        switch fi22 {
                                                                                                        case 32767:
                                                                                                            if true {
                                                                                                                var v63 int = (fi20 + 2147483647)
                                                                                                                v43 += v34
                                                                                                                v64 := GenS3{F0: -128, F1: (fi20 - 42), F2: 0}
                                                                                                                v31 = append(v31, (-(v63)))
                                                                                                                _ = v63
                                                                                                                _ = v64
                                                                                                            }
                                                                                                        case -9223372036854775808:
                                                                                                            v65 := GenS1{F0: (v5 * (v5 << (((^(v44))) & 63))), F1: v7}
                                                                                                            _ = v65
                                                                                                        }
                                                                                                        v43 = (v34 * -2.25)
                                                                                                        var v66 uint32 = 1
                                                                                                        if (((-2.25 - -2.25) + 0.0) > (1.5 + ((((-2.25 + (1.5 / 3.14159)) / 0.1) * 0.1) * (0.1 / float64(((v34 / 3.14159) * 3.14159)))))) {
                                                                                                            v31 = append(v31, (fi22 << (((v4 + 0)) & 63)))
                                                                                                        } else {
                                                                                                            fmt.Println(v6, v29, v3)
                                                                                                        }
                                                                                                        _ = v66
                                                                                                    } else {
                                                                                                        var v67 string = v7
                                                                                                        _ = v67
                                                                                                    }
                                                                                                    switch int(v62) {
                                                                                                    case 42:
                                                                                                        var v68 int8 = -94
                                                                                                        _ = v68
                                                                                                    case -1:
                                                                                                        if false {
                                                                                                            v6 -= 2
                                                                                                        } else {
                                                                                                            v69 := GenS3{F0: (v62 << ((1) & 63)), F1: fi22, F2: v6}
                                                                                                            fmt.Println(v6, v44, v3)
                                                                                                            switch (fi22 >> (((v44 * 0)) & 63)) {
                                                                                                            case -690:
                                                                                                                var v70 uint16 = 32767
                                                                                                                _ = v70
                                                                                                            default:
                                                                                                                switch (fi20 >> (((v2 | v2)) & 63)) {
                                                                                                                case -2147483648:
                                                                                                                    delete(v1, (v7 + ""))
                                                                                                                case 2147483647:
                                                                                                                    v71 := GenS1{F0: v45, F1: ((v7 + v7) + "hello")}
                                                                                                                    _ = v71
                                                                                                                case -417:
                                                                                                                    if true {
                                                                                                                        var v72 float64 = (0.1 + float64(3.14159))
                                                                                                                        v73 := GenS6{F0: false, F1: v5, F2: 0}
                                                                                                                        if false {
                                                                                                                            var v74 float32 = v43
                                                                                                                            if len(v31) > 0 {
                                                                                                                                v31[0] = (fi22 << ((uint(v47)) & 63))
                                                                                                                            }
                                                                                                                            fmt.Println(v5, v44, v45)
                                                                                                                            var v75 bool = (!(!(((v72 / 0.1) / (((((float64(0.0) + float64((v34 * (3.14159 / (float32(float64(v34)) + (1.5 + (1.5 - v74))))))) / v72) / v72) / float64(v43)) + v72)) == ((v72 - (float64(v74) * float64(1.5))) * ((((0.0 - float64(v34)) * -2.25) + float64(3.14159)) * float64(v34))))))
                                                                                                                            switch (fi20 - 9) {
                                                                                                                            case -9223372036854775808:
                                                                                                                                v29 = (v4 + v29)
                                                                                                                            default:
                                                                                                                                v7 = "hello"
                                                                                                                            }
                                                                                                                            _ = v74
                                                                                                                            _ = v75
                                                                                                                        } else {
                                                                                                                            fmt.Println(v2)
                                                                                                                            v47 = (v47 + 255)
                                                                                                                            v76 := GenS7{F0: 474, F1: v7, F2: (v3 / ((uint8(v6)) | 1))}
                                                                                                                            _ = v76
                                                                                                                        }
                                                                                                                        _ = v72
                                                                                                                        _ = v73
                                                                                                                    } else {
                                                                                                                        v77 := GenS8{F0: 2}
                                                                                                                        v47 = 127
                                                                                                                        _ = v77
                                                                                                                    }
                                                                                                                }
                                                                                                            }
                                                                                                            _ = v69
                                                                                                        }
                                                                                                    }
                                                                                                }
                                                                                            }
                                                                                        default:
                                                                                            var v78 uint16 = 32767
                                                                                            _ = v78
                                                                                        }
                                                                                        v79 := GenS6{F0: false, F1: v45, F2: 65535}
                                                                                        v80 := GenS6{F0: (true || true), F1: (v5 * v45), F2: 158}
                                                                                        fmt.Println(v5, v47)
                                                                                        _ = v79
                                                                                        _ = v80
                                                                                    }
                                                                                    var v81 int = fi20
                                                                                    v1["\n\t"] = 529
                                                                                    _ = v62
                                                                                    _ = v81
                                                                                }
                                                                                v7 = ""
                                                                            } else {
                                                                                v82 := GenS4{F0: 1, F1: (v6 / (((^(v6))) | 1))}
                                                                                _ = v82
                                                                            }
                                                                            if false {
                                                                                var v83 string = (("a" + (((v7 + (("foo bar" + "\n\t") + "a")) + v7) + v7)) + v7)
                                                                                v31 = append(v31, 9223372036854775807)
                                                                                _ = v83
                                                                            }
                                                                            if false {
                                                                                switch fi20 {
                                                                                case 42:
                                                                                    v84 := GenS7{F0: 1, F1: "\n\t", F2: 255}
                                                                                    _ = v84
                                                                                }
                                                                                if true {
                                                                                    v85 := GenS5{F0: v47, F1: v44}
                                                                                    switch (fi22 % (((fi22 % (((fi22 << (((v44 * v2)) & 63))) | 1))) | 1)) {
                                                                                    case 1:
                                                                                        if false {
                                                                                            var v86 int = fi20
                                                                                            var v87 uint32 = 2147483648
                                                                                            v29 -= v29
                                                                                            v88 := GenS9{F0: 42, F1: 1000}
                                                                                            _ = v86
                                                                                            _ = v87
                                                                                            _ = v88
                                                                                        }
                                                                                        v34 -= v34
                                                                                    }
                                                                                    v47 = v3
                                                                                    _ = v85
                                                                                }
                                                                                v89 := GenS10{F0: false}
                                                                                _ = v89
                                                                            }
                                                                            _ = v45
                                                                            _ = v47
                                                                        case "a":
                                                                            var v90 uint64 = (^(v6))
                                                                            _ = v90
                                                                        }
                                                                        _ = v43
                                                                        _ = v44
                                                                    }
                                                                case "\n\t":
                                                                    var v91 int64 = (v5 / ((v5) | 1))
                                                                    _ = v91
                                                                default:
                                                                    v34 = v34
                                                                }
                                                                var v92 float32 = v34
                                                                v3 = v3
                                                                v3 = v3
                                                                _ = v92
                                                            }
                                                            _ = v34
                                                        } else {
                                                            v1[v7] = 387
                                                            if false {
                                                                if ((!false) && (264 >= 0)) {
                                                                    fmt.Println(v3)
                                                                    _ = len(v31)
                                                                    var v93 int64 = 0
                                                                    _ = v93
                                                                } else {
                                                                    var v94 float64 = (float64(0.0) / float64(((3.14159 * (((float32(0.0) - 0.0) + 3.14159) + ((3.14159 * ((1.5 / 1.5) + (float32(0.0) / ((float32(1.5) * (((0.1 + ((3.14159 / ((-2.25 - -2.25) - float32(1.5))) - 3.14159)) * (-2.25 + (1.5 + ((float32(1e10) + -2.25) * 1e10)))) - -2.25)) - (1.5 * ((((1.5 * 1e10) + (float32(1.5) * 3.14159)) + 0.1) * 1.5)))))) - 1.5))) + float32((0.0 - -2.25)))))
                                                                    _ = v94
                                                                }
                                                                fmt.Println(v4, fi20, v4)
                                                                switch fi22 {
                                                                case 127:
                                                                    v95 := GenS8{F0: 2147483647}
                                                                    _ = v95
                                                                case 2147483647:
                                                                    if ((v4 / ((18446744073709551615) | 1)) < v4) {
                                                                        v6 -= (v6 / (((v6 % ((v6) | 1))) | 1))
                                                                        v3 += (v3 ^ v3)
                                                                        v96 := GenS11{F0: (v5 & 0), F1: true}
                                                                        v1[v7] = uint32(v3)
                                                                        _ = v96
                                                                    } else {
                                                                        fmt.Println(v2, v7, fi22)
                                                                    }
                                                                default:
                                                                    v1[v7] = 42
                                                                }
                                                            }
                                                        }
                                                    }
                                                    var v97 int16 = -248
                                                    _ = v31
                                                    _ = v97
                                                } else {
                                                    var v98 int = (fi22 >> ((v2) & 63))
                                                    switch (v7 + ("hello" + "foo bar")) {
                                                    case "":
                                                        v3 -= (^(v3))
                                                    case "\n\t":
                                                        v99 := GenS5{F0: v3, F1: v4}
                                                        _ = v99
                                                    case "foo bar":
                                                        v2 += (v2 | v2)
                                                    default:
                                                        var v100 float64 = 0.0
                                                        _ = v100
                                                    }
                                                    _ = v98
                                                }
                                            } else {
                                                if false {
                                                    if (((1e10 / 1.5) >= (0.1 + 1e10)) || true) {
                                                        v5 -= v5
                                                        if (((v2 == v4) || (!(!(float64((3.14159 / -2.25)) > 1.5)))) && (true || true)) {
                                                            v2 -= (v4 << ((v4) & 63))
                                                            var v101 int16 = 0
                                                            v2 += v2
                                                            var v102 int8 = int8(fi20)
                                                            _ = v101
                                                            _ = v102
                                                        } else {
                                                            v103 := []int{(fi22 % ((911) | 1)), fi20}
                                                            var v104 uint = v29
                                                            if (0 > 0) {
                                                                v104 += v104
                                                            } else {
                                                                var v105 uint16 = 1
                                                                _ = v105
                                                            }
                                                            _ = v103
                                                            _ = v104
                                                        }
                                                        var v106 float64 = 0.1
                                                        fmt.Println(v106)
                                                        switch (fi22 + int(v6)) {
                                                        case 42:
                                                            var v107 int32 = 313
                                                            _ = v107
                                                        }
                                                        _ = v106
                                                    } else {
                                                        v108 := GenS2{F0: 42, F1: 2, F2: v7}
                                                        switch fi20 {
                                                        case -162:
                                                            v5 -= (v5 + (^(v5)))
                                                        }
                                                        _ = v108
                                                    }
                                                    fmt.Println(v2, v2, v29)
                                                }
                                                var v109 int64 = 0
                                                var v110 bool = false
                                                fmt.Println(v3)
                                                _ = v109
                                                _ = v110
                                            }
                                            v2 += (v29 + 18446744073709551615)
                                        case -435:
                                            if false {
                                                var v111 uint = (^(v2))
                                                _ = v111
                                            } else {
                                                var v112 int32 = -239
                                                v113 := GenS12{F0: ((482 > 42) && false)}
                                                if (!((v112 + (v112 | (^(v112)))) <= (v112 ^ (v112 / (((v112 ^ 0)) | 1))))) {
                                                    var v114 int8 = -52
                                                    v112 = v112
                                                    var v115 int = (fi20 * (fi20 >> (((v29 * (v2 >> ((v29) & 63)))) & 63)))
                                                    _ = v114
                                                    _ = v115
                                                }
                                                _ = v112
                                                _ = v113
                                            }
                                        }
                                        var v116 int32 = -267
                                        _ = v116
                                        _ = v29
                                        _ = v30
                                    } else {
                                        var v117 float32 = (-2.25 - (-2.25 * 3.14159))
                                        var v118 int8 = -97
                                        _ = v117
                                        _ = v118
                                    }
                                    _ = fi22
                                    _ = v26
                                }
                            }
                        } else {
                            v119 := GenS13{F0: 225}
                            _ = v119
                        }
                        var v120 int16 = 32767
                        var v121 int16 = v120
                        _ = v120
                        _ = v121
                    } else {
                        var v122 int64 = v5
                        _ = v122
                    }
                    v6 = (v6 / ((v6) | 1))
                    _ = fi20
                }
            }
        case -540:
            for fi123 := 0; fi123 < 4; fi123++ {
                var v124 int32 = 2147483647
                v124 = v124
                _ = fi123
                _ = v124
            }
        }
        _ = v8
    } else {
        v125 := []int{0, -52}
        var v126 int16 = 242
        _ = v125
        _ = v126
    }
    v4 += 946
    v127 := v1[v7]
    if (-2147483648 <= -2147483648) {
        _ = len(v1)
        switch 613 {
        case 127:
            if false {
                v7 = (v7 + v7)
                fmt.Println(v5)
            } else {
                if false {
                    for fi128 := 0; fi128 < 8; fi128++ {
                        var v129 int32 = 2147483647
                        _ = fi128
                        _ = v129
                    }
                    delete(v1, ("" + (v7 + (v7 + v7))))
                    v130 := []int{1, 9223372036854775807, 322}
                    _ = v130
                }
                v7 = v7
            }
        case 1:
            v131 := GenS4{F0: 255, F1: 42}
            _ = v131
        }
        var v132 int64 = v5
        _ = v132
    } else {
        switch (v7 + ((v7 + v7) + "hello")) {
        case "foo bar":
            if true {
                switch v7 {
                case "\n\t":
                    if false {
                        var v133 int32 = 309
                        if false {
                            if false {
                                var v134 int32 = v133
                                v135 := GenS14{F0: 1, F1: (v5 <= 42)}
                                _ = v134
                                _ = v135
                            } else {
                                v133 = -42
                                for fi136 := 0; fi136 < 13; fi136++ {
                                    v7 = v7
                                    for fi137 := 0; fi137 < 16; fi137++ {
                                        if (-755 == -32768) {
                                            if ((3.14159 - 1e10) == 1.5) {
                                                var v138 int = int(v5)
                                                if true {
                                                    switch v7 {
                                                    case "foo bar":
                                                        if (true && ((v3 ^ (^(v3))) >= (v3 << ((v2) & 63)))) {
                                                            var v139 uint8 = v3
                                                            _ = v139
                                                        }
                                                    case "a":
                                                        fmt.Println(v4, v4)
                                                    case "hello":
                                                        var v140 uint32 = v127
                                                        _ = v140
                                                    default:
                                                        v1["\n\t"] = 82
                                                    }
                                                    if true {
                                                        v133 = (v133 ^ (v133 * (v133 >> (((v2 + v4)) & 63))))
                                                    } else {
                                                        v141 := []int{(fi136 ^ v138)}
                                                        var v142 bool = true
                                                        _ = v141
                                                        _ = v142
                                                    }
                                                } else {
                                                    if false {
                                                        var v143 int = int(v127)
                                                        v3 -= (v3 % (((v3 ^ (v3 % ((v3) | 1)))) | 1))
                                                        _ = v143
                                                    } else {
                                                        _ = len(v1)
                                                        var v144 int = (fi137 * fi137)
                                                        _ = v144
                                                    }
                                                    var v145 int = 32768
                                                    _ = v145
                                                }
                                                _ = v138
                                            } else {
                                                var v146 uint8 = (v3 + v3)
                                                v4 = v2
                                                _ = v146
                                            }
                                            if (((false && true) && (v7 >= "hello")) || (2 >= 0)) {
                                                switch (-(fi136)) {
                                                case 9223372036854775807:
                                                    v147 := GenS11{F0: 0, F1: (227 < (v5 >> (((v4 / ((255) | 1))) & 63)))}
                                                    _ = v147
                                                case -2147483648:
                                                    v3 = (v3 << ((0) & 63))
                                                case 0:
                                                    v148 := []int{(fi136 ^ -1)}
                                                    _ = v148
                                                default:
                                                    v4 = v4
                                                }
                                                if ((^(v3)) <= uint8(v4)) {
                                                    v2 = v4
                                                    v2 -= (v2 / (((v2 << ((v4) & 63))) | 1))
                                                    v127 -= 1
                                                }
                                            }
                                        } else {
                                            var v149 uint = v4
                                            if true {
                                                if (v3 > (v3 | v3)) {
                                                    if true {
                                                        var v150 int64 = (v5 % ((v5) | 1))
                                                        var v151 uint = v149
                                                        var v152 float64 = float64(1e10)
                                                        _ = v150
                                                        _ = v151
                                                        _ = v152
                                                    }
                                                } else {
                                                    fmt.Println(fi136, v6, v5)
                                                    if false {
                                                        var v153 int8 = -99
                                                        if (0 >= 559) {
                                                            v154 := v1[v7]
                                                            _ = v154
                                                        } else {
                                                            var v155 float32 = 3.14159
                                                            v156 := []int{fi137, (fi136 | (fi137 ^ fi137))}
                                                            var v157 uint64 = 0
                                                            fmt.Println(v153, v5)
                                                            _ = v155
                                                            _ = v156
                                                            _ = v157
                                                        }
                                                        _ = v153
                                                    } else {
                                                        v133 = (v133 | v133)
                                                        v158 := []int{}
                                                        v158 = append(v158, (-(fi136)))
                                                        _ = v158
                                                    }
                                                    v159 := []int{}
                                                    _ = v159
                                                }
                                            } else {
                                                var v160 uint64 = v6
                                                _ = v160
                                            }
                                            if (32767 != fi136) {
                                                var v161 int32 = v133
                                                fmt.Println(v127)
                                                var v162 int64 = v5
                                                _ = v161
                                                _ = v162
                                            } else {
                                                if false {
                                                    if false {
                                                        var v163 int32 = v133
                                                        _ = v163
                                                    } else {
                                                        v164 := GenS7{F0: 2, F1: "foo bar", F2: v3}
                                                        v4 = v149
                                                        v165 := []int{(fi136 - (fi137 / ((fi137) | 1))), fi137, fi137, (fi136 - -1)}
                                                        _ = v164
                                                        _ = v165
                                                    }
                                                    v4 = v2
                                                    fmt.Println(v4)
                                                } else {
                                                    var v166 int = fi137
                                                    _ = v166
                                                }
                                                v5 += v5
                                            }
                                            _ = v149
                                        }
                                        var v167 int32 = v133
                                        _ = fi137
                                        _ = v167
                                    }
                                    switch v7 {
                                    case "a":
                                        v3 -= (v3 % ((v3) | 1))
                                    case "foo bar":
                                        if (float32(0.0) < 0.0) {
                                            v168 := GenS3{F0: 100, F1: (fi136 % ((128) | 1)), F2: 30}
                                            v169 := GenS9{F0: -1, F1: -32768}
                                            var v170 uint32 = v127
                                            _ = v168
                                            _ = v169
                                            _ = v170
                                        } else {
                                            v171 := GenS7{F0: v127, F1: v7, F2: v3}
                                            _ = v171
                                        }
                                    case "hello":
                                        v5 += 897
                                    }
                                    _ = fi136
                                }
                                var v172 int = 337
                                _ = v172
                            }
                        }
                        _ = v133
                    }
                case "":
                    var v173 bool = (true && true)
                    _ = v173
                case "a":
                    if (!true) {
                        if (42 == int16(v6)) {
                            var v174 bool = (((-(v5)) == (^(v5))) && false)
                            var v175 bool = (v174 && v174)
                            v176 := GenS6{F0: true, F1: v5, F2: 2}
                            _ = v174
                            _ = v175
                            _ = v176
                        }
                        if false {
                            v177 := v1[v7]
                            v178 := []int{-838, 4294967295, -1, 626}
                            v179 := GenS15{F0: 127, F1: 765, F2: 369}
                            _ = v177
                            _ = v178
                            _ = v179
                        }
                        var v180 uint16 = 2
                        _ = v180
                    }
                default:
                    for fi181 := 0; fi181 < 13; fi181++ {
                        fmt.Println(v127, v5, v2)
                        if true {
                            v7 = v7
                            if true {
                                v6 = v6
                                var v182 int32 = 2147483647
                                _ = v182
                            }
                            for fi183 := 0; fi183 < 2; fi183++ {
                                if (true || true) {
                                    var v184 float64 = (float64(((0.1 * 1.5) + 1.5)) - -2.25)
                                    v185 := GenS12{F0: (0.0 > (0.0 - 1e10))}
                                    fmt.Println(v127, v5)
                                    _ = v184
                                    _ = v185
                                }
                                _ = fi183
                            }
                        } else {
                            for fi186 := 0; fi186 < 2; fi186++ {
                                var v187 uint8 = (v3 | v3)
                                v187 = 99
                                _ = fi186
                                _ = v187
                            }
                            if ((3.14159 + 3.14159) == (1e10 * -2.25)) {
                                for fi188 := 0; fi188 < 11; fi188++ {
                                    var v189 uint8 = uint8(v6)
                                    v190 := GenS6{F0: false, F1: (v5 >> ((v4) & 63)), F2: v127}
                                    v6 += 863
                                    _ = fi188
                                    _ = v189
                                    _ = v190
                                }
                                var v191 uint16 = uint16(v6)
                                _ = v191
                            }
                        }
                        _ = fi181
                    }
                }
                var v192 int32 = 2147483647
                var v193 bool = true
                _ = v192
                _ = v193
            } else {
                var v194 int = 32767
                if ((v6 / ((uint64(v3)) | 1)) >= (v6 / ((0) | 1))) {
                    v195 := []int{1, (v194 | (v194 % ((v194) | 1))), (v194 & v194), (v194 >> ((v4) & 63))}
                    _ = v195
                }
                _ = v194
            }
        case "":
            for fi196 := 0; fi196 < 1; fi196++ {
                switch (fi196 + (fi196 >> ((v4) & 63))) {
                case 3:
                    var v197 float64 = float64((3.14159 / (1.5 + float32((-2.25 * -2.25)))))
                    _ = v197
                case 870:
                    v198 := GenS3{F0: -54, F1: (fi196 ^ int(v6)), F2: (^(v6))}
                    _ = v198
                case 9223372036854775807:
                    fmt.Println(fi196, v127, v127)
                }
                for fi199 := 0; fi199 < 7; fi199++ {
                    var v200 uint16 = 386
                    v201 := GenS9{F0: -128, F1: 32767}
                    var v202 int64 = (-(v5))
                    _ = fi199
                    _ = v200
                    _ = v201
                    _ = v202
                }
                _ = fi196
            }
        default:
            switch (v7 + "\n\t") {
            case "hello":
                v203 := GenS4{F0: v4, F1: v6}
                _ = v203
            }
        }
        var v204 uint8 = v3
        fmt.Println(v5)
        _ = v204
    }
    v205 := GenS16{F0: 359}
    var v206 bool = (4294967295 != (^(v4)))
    v207 := GenS4{F0: 1, F1: 699}
    fmt.Println(v127, v206)
    v208 := GenS17{F0: 606, F1: v5}
    v209 := []int{127, 32767, 128, 0}
    for fi210 := 0; fi210 < 5; fi210++ {
        if v206 {
            v211 := GenS10{F0: (("\n\t" + v7) <= (v7 + "\n\t"))}
            var v212 int16 = 32767
            v209 = append(v209, (fi210 % (((^(fi210))) | 1)))
            _ = v211
            _ = v212
        } else {
            for fi213 := 0; fi213 < 4; fi213++ {
                v214 := GenS18{F0: v6, F1: 118}
                _ = fi213
                _ = v214
            }
        }
        _ = fi210
    }
    for fi215 := 0; fi215 < 5; fi215++ {
        v216 := GenS18{F0: v6, F1: 127}
        _ = fi215
        _ = v216
    }
    if v206 {
        v209 = append(v209, 2147483647)
        switch v7 {
        case "":
            switch v7 {
            case "a":
                var v217 float32 = 3.14159
                _ = v217
            }
        case "\n\t":
            switch "\n\t" {
            case "a":
                var v218 float32 = -2.25
                _ = v218
            case "hello":
                var v219 float64 = 3.14159
                _ = v219
            default:
                v220 := GenS19{F0: 0}
                _ = v220
            }
        default:
            fmt.Println(v2, v206, v206)
        }
    } else {
        v221 := GenS6{F0: v206, F1: v5, F2: v127}
        _ = v221
    }
    if (!v206) {
        for fi222 := 0; fi222 < 3; fi222++ {
            switch (fi222 + 1) {
            case 537:
                if v206 {
                    var v223 int = 42
                    _ = v223
                } else {
                    fmt.Println(fi222, v4, v7)
                }
            case 32767:
                if v206 {
                    var v224 int64 = 1
                    var v225 uint64 = v6
                    _ = v224
                    _ = v225
                }
            }
            _ = fi222
        }
        delete(v1, v7)
        for fi226 := 0; fi226 < 15; fi226++ {
            v7 = (v7 + v7)
            v4 = v4
            v209 = append(v209, (fi226 + fi226))
            _ = fi226
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v127:", v127)
    fmt.Println("v2:", v2)
    fmt.Println("v206:", v206)
    fmt.Println("slicelen v209:", len(v209), v209)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v127
    _ = v2
    _ = v205
    _ = v206
    _ = v207
    _ = v208
    _ = v209
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
