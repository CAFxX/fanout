package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8 }
type GenS2 struct { F0 uint8; F1 uint64 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint64; F1 int }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 string; F1 uint64 }
type GenS7 struct { F0 uint16; F1 uint64; F2 int16 }
type GenS8 struct { F0 bool; F1 string; F2 uint8 }
type GenS9 struct { F0 uint16; F1 bool; F2 uint32 }
type GenS10 struct { F0 int16 }
type GenS11 struct { F0 int64 }
type GenS12 struct { F0 uint32; F1 int8; F2 int16 }
type GenS13 struct { F0 string }
type GenS14 struct { F0 int8; F1 uint64 }
type GenS15 struct { F0 int16; F1 uint16; F2 uint16 }
type GenS16 struct { F0 int64; F1 int32 }
type GenS17 struct { F0 uint8; F1 bool }
type GenS18 struct { F0 uint; F1 string }
type GenS19 struct { F0 bool; F1 uint32 }
type GenS20 struct { F0 int64; F1 bool }
type GenS21 struct { F0 string }
type GenS22 struct { F0 string }
type GenS23 struct { F0 uint16; F1 int32; F2 uint }
type GenS24 struct { F0 uint64; F1 int32 }
type GenS25 struct { F0 string; F1 uint64 }
type GenS26 struct { F0 uint16; F1 int64; F2 uint }
type GenS27 struct { F0 uint }
func genFn0() bool {
    var v6 string = ("hello" + "hello")
    var v7 string = "foo bar"
    _ = v6
    _ = v7
    return true
}

func main() {
    var v1 string = "a"
    v2 := []uint32{}
    var v3 float64 = float64(-2.25)
    v4 := []int{1, 128}
    var v5 float32 = -2.25
    if (!true) {
        var v8 uint16 = 1
        for fi9 := 0; fi9 < 11; fi9++ {
            v4 = append(v4, (fi9 >> ((uint(v8)) & 63)))
            var v10 float32 = 1.5
            for fi11 := 0; fi11 < 15; fi11++ {
                v2 = append(v2, 42)
                if true {
                    v12 := GenS1{F0: 42, F1: 100}
                    v13 := GenS2{F0: 196, F1: 302}
                    if false {
                        v14 := GenS1{F0: 43, F1: int8(v8)}
                        _ = v14
                    } else {
                        var v15 int16 = 1000
                        v4 = append(v4, 42)
                        _ = v15
                    }
                    _ = v12
                    _ = v13
                } else {
                    if true {
                        v10 = v5
                    } else {
                        var v16 uint64 = 78
                        var v17 string = v1
                        v5 -= (v5 - (((float32((v3 + v3)) - float32(0.1)) + ((v5 + v5) - (1.5 - (0.0 * v5)))) * v5))
                        _ = v16
                        _ = v17
                    }
                    if ((^(v8)) != (v8 >> ((18446744073709551615) & 63))) {
                        if len(v4) > 0 {
                            v4[0] = 127
                        }
                    } else {
                        if false {
                            switch fi9 {
                            case -189:
                                if (!true) {
                                    v10 = v10
                                    var v18 int32 = -1
                                    _ = v18
                                } else {
                                    v19 := GenS2{F0: 42, F1: 2}
                                    var v20 uint16 = v8
                                    switch (^(fi9)) {
                                    case 0:
                                        fmt.Println(v1, fi9)
                                    case 9223372036854775807:
                                        var v21 bool = (uint32(v20) >= 2147483647)
                                        _ = v21
                                    default:
                                        v10 = (float32(v3) - (1e10 * v5))
                                    }
                                    _ = v19
                                    _ = v20
                                }
                                fmt.Println(v3, v5, v8)
                            case 4294967295:
                                if (!true) {
                                    if true {
                                        if (127 >= 255) {
                                            _ = len(v4)
                                            v22 := GenS2{F0: 2, F1: 352}
                                            _ = v22
                                        } else {
                                            if true {
                                                v23 := GenS1{F0: 128, F1: 100}
                                                var v24 int64 = 668
                                                if len(v4) >= 2 {
                                                    v4 = v4[1:len(v4)]
                                                }
                                                _ = v23
                                                _ = v24
                                            } else {
                                                v25 := GenS3{F0: (^(v8))}
                                                var v26 uint64 = 2
                                                _ = v25
                                                _ = v26
                                            }
                                            if ((!("foo bar" > ("hello" + v1))) || (8 > 0)) {
                                                v2 = append(v2, 1)
                                            } else {
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                            }
                                            v27 := GenS4{F0: 2, F1: int(v8)}
                                            _ = v27
                                        }
                                    }
                                }
                            }
                            var v28 int64 = 0
                            v10 = -2.25
                            var v29 uint32 = 329
                            switch -48 {
                            case 4294967295:
                                v30 := GenS1{F0: uint8(v29), F1: 56}
                                _ = v30
                            case -2147483648:
                                if ((!false) && true) {
                                    switch (-(fi9)) {
                                    case 9223372036854775807:
                                        v29 += 2147483648
                                    case 32768:
                                        switch (fi9 * (fi9 % ((32768) | 1))) {
                                        case -2147483648:
                                            if len(v2) > 0 {
                                                v2[0] = (^(v29))
                                            }
                                        case -659:
                                            if len(v2) > 0 {
                                                v2[0] = 403
                                            }
                                        case -451:
                                            if true {
                                                var v31 int16 = 1
                                                v10 = v10
                                                _ = v31
                                            } else {
                                                if false {
                                                    var v32 uint64 = 2
                                                    if false {
                                                        if true {
                                                            v8 -= 834
                                                            if true {
                                                                switch v1 {
                                                                case "foo bar":
                                                                    v33 := GenS4{F0: (v32 - uint64(v8)), F1: fi9}
                                                                    _ = v33
                                                                case "hello":
                                                                    var v34 int8 = int8(v8)
                                                                    _ = v34
                                                                }
                                                                v3 -= 1.5
                                                                v32 += (v32 << ((18446744073709551615) & 63))
                                                            }
                                                            switch fi11 {
                                                            case -22:
                                                                v5 += 3.14159
                                                            }
                                                        } else {
                                                            fmt.Println(v28, v28)
                                                            var v35 uint32 = v29
                                                            if true {
                                                                if len(v2) > 0 {
                                                                    v2[0] = v29
                                                                }
                                                            } else {
                                                                switch fi11 {
                                                                case -1:
                                                                    if (false || false) {
                                                                        var v36 uint32 = (v35 / ((0) | 1))
                                                                        v37 := GenS5{F0: v29}
                                                                        _ = v36
                                                                        _ = v37
                                                                    } else {
                                                                        v38 := GenS6{F0: ("a" + v1), F1: 427}
                                                                        v8 = 924
                                                                        if (-685 < -1) {
                                                                            if len(v4) > 0 {
                                                                                v4[0] = fi11
                                                                            }
                                                                            var v39 string = (("a" + v1) + "foo bar")
                                                                            v35 -= v35
                                                                            var v40 uint16 = v8
                                                                            _ = v39
                                                                            _ = v40
                                                                        } else {
                                                                            if true {
                                                                                if ("hello" >= ((v1 + ((v1 + v1) + (v1 + "foo bar"))) + v1)) {
                                                                                    fmt.Println(v5, fi11)
                                                                                } else {
                                                                                    var v41 int64 = (v28 >> ((4294967295) & 63))
                                                                                    _ = v41
                                                                                }
                                                                            }
                                                                            fmt.Println(fi9)
                                                                            var v42 uint = 85
                                                                            _ = v42
                                                                        }
                                                                        _ = v38
                                                                    }
                                                                case -28:
                                                                    fmt.Println(v29)
                                                                case 2147483647:
                                                                    v43 := GenS2{F0: 128, F1: (v32 % (((v32 | (v32 - v32))) | 1))}
                                                                    _ = v43
                                                                }
                                                                var v44 uint8 = 127
                                                                v45 := GenS6{F0: "foo bar", F1: (v32 >> ((623) & 63))}
                                                                _ = v44
                                                                _ = v45
                                                            }
                                                            _ = v35
                                                        }
                                                        v28 = -605
                                                    } else {
                                                        var v46 uint64 = (v32 % ((v32) | 1))
                                                        v28 = v28
                                                        _ = v46
                                                    }
                                                    _ = v32
                                                } else {
                                                    if false {
                                                        var v47 float32 = v5
                                                        fmt.Println(v28, v10)
                                                        _ = v47
                                                    } else {
                                                        v48 := GenS5{F0: uint32(v8)}
                                                        v2 = append(v2, (v29 >> ((0) & 63)))
                                                        _ = v48
                                                    }
                                                    switch (fi9 >> ((350) & 63)) {
                                                    case 1:
                                                        var v49 float32 = v10
                                                        _ = v49
                                                    case 42:
                                                        var v50 bool = ((!(127 > 255)) || (10 > 127))
                                                        _ = v50
                                                    default:
                                                        v51 := GenS1{F0: 42, F1: 127}
                                                        _ = v51
                                                    }
                                                    if (-128 != 1) {
                                                        switch (-(fi9)) {
                                                        case -909:
                                                            fmt.Println(v29, v28, v1)
                                                        default:
                                                            fmt.Println(v28, v5, fi9)
                                                        }
                                                        fmt.Println(v3, v1, v28)
                                                    } else {
                                                        var v52 uint = 18446744073709551615
                                                        v53 := GenS5{F0: v29}
                                                        _ = len(v2)
                                                        v4 = append(v4, 2147483647)
                                                        _ = v52
                                                        _ = v53
                                                    }
                                                }
                                            }
                                            switch (v1 + "hello") {
                                            case "a":
                                                if (false && false) {
                                                    if false {
                                                        if (1 > 18446744073709551615) {
                                                            var v54 uint32 = 2147483647
                                                            if (true && true) {
                                                                var v55 float32 = 3.14159
                                                                switch "foo bar" {
                                                                case "a":
                                                                    v56 := GenS7{F0: (v8 - 459), F1: 4294967295, F2: -545}
                                                                    _ = v56
                                                                case "foo bar":
                                                                    var v57 uint8 = 42
                                                                    v55 += v55
                                                                    v58 := GenS8{F0: false, F1: v1, F2: v57}
                                                                    _ = v57
                                                                    _ = v58
                                                                }
                                                                _ = v55
                                                            }
                                                            if true {
                                                                if false {
                                                                    var v59 uint64 = 18446744073709551615
                                                                    _ = v59
                                                                } else {
                                                                    if (!true) {
                                                                        v2 = append(v2, v29)
                                                                        var v60 int32 = -337
                                                                        _ = v60
                                                                    } else {
                                                                        switch (fi9 - fi11) {
                                                                        case 32768:
                                                                            var v61 uint64 = 2
                                                                            _ = v61
                                                                        case 9223372036854775807:
                                                                            var v62 int16 = 761
                                                                            _ = v62
                                                                        case -9223372036854775808:
                                                                            v54 -= v29
                                                                        }
                                                                        if true {
                                                                            _ = len(v2)
                                                                            var v63 uint64 = 42
                                                                            _ = v63
                                                                        } else {
                                                                            var v64 int64 = (-(v28))
                                                                            v29 += 4294967295
                                                                            _ = len(v4)
                                                                            _ = v64
                                                                        }
                                                                    }
                                                                    v54 -= (v29 % ((944) | 1))
                                                                    if true {
                                                                        fmt.Println(v28, v5, v54)
                                                                        v65 := GenS4{F0: uint64(v8), F1: (fi9 + -473)}
                                                                        v10 = float32(v3)
                                                                        _ = v65
                                                                    }
                                                                }
                                                                v10 += 0.1
                                                                fmt.Println(fi11)
                                                                if true {
                                                                    if false {
                                                                        var v66 int = (fi11 << ((271) & 63))
                                                                        _ = v66
                                                                    }
                                                                    v1 = ("a" + v1)
                                                                }
                                                                fmt.Println(v10)
                                                            }
                                                            _ = v54
                                                        }
                                                        if true {
                                                            fmt.Println(fi11, v3, v28)
                                                            switch ("" + (v1 + v1)) {
                                                            case "hello":
                                                                v3 = (float64((v10 + 3.14159)) * float64(v5))
                                                                fmt.Println(v29, v3, v10)
                                                                var v67 int32 = int32(v29)
                                                                fmt.Println(v1)
                                                                _ = v67
                                                            case "foo bar":
                                                                var v68 uint16 = v8
                                                                _ = v68
                                                            default:
                                                                var v69 float32 = (v5 - v10)
                                                                _ = v69
                                                            }
                                                            switch v1 {
                                                            case "\n\t":
                                                                v8 += v8
                                                            case "hello":
                                                                var v70 float64 = v3
                                                                _ = v70
                                                            default:
                                                                fmt.Println(v5)
                                                            }
                                                        } else {
                                                            fmt.Println(v29, v29, v8)
                                                            if (98 < -128) {
                                                                v71 := GenS8{F0: true, F1: (v1 + "\n\t"), F2: 0}
                                                                v1 = "a"
                                                                v1 = v1
                                                                switch (("foo bar" + v1) + "\n\t") {
                                                                case "hello":
                                                                    switch (fi9 ^ fi11) {
                                                                    case -2147483648:
                                                                        if true {
                                                                            var v72 int16 = 1
                                                                            v73 := GenS9{F0: v8, F1: false, F2: 1}
                                                                            _ = v72
                                                                            _ = v73
                                                                        }
                                                                    case 9223372036854775807:
                                                                        var v74 uint32 = v29
                                                                        _ = v74
                                                                    case -563:
                                                                        _ = len(v4)
                                                                    default:
                                                                        if ((!true) && false) {
                                                                            v75 := GenS8{F0: false, F1: (("foo bar" + ((v1 + v1) + "hello")) + ""), F2: 1}
                                                                            var v76 uint8 = 2
                                                                            if true {
                                                                                v77 := GenS10{F0: -262}
                                                                                if true {
                                                                                    if true {
                                                                                        v2 = append(v2, v29)
                                                                                        v78 := GenS2{F0: 128, F1: 1}
                                                                                        var v79 uint16 = 191
                                                                                        _ = v78
                                                                                        _ = v79
                                                                                    }
                                                                                    var v80 float64 = (((1.5 - v3) / ((-2.25 - v3) - 3.14159)) - -2.25)
                                                                                    _ = v80
                                                                                } else {
                                                                                    fmt.Println(v3, v76)
                                                                                    var v81 float64 = v3
                                                                                    var v82 float64 = ((0.1 + (3.14159 - (v81 + (((v3 / v3) * 1e10) - float64(float32(3.14159)))))) - (v3 - 1.5))
                                                                                    _ = v81
                                                                                    _ = v82
                                                                                }
                                                                                switch v1 {
                                                                                case "foo bar":
                                                                                    fmt.Println(v29)
                                                                                case "hello":
                                                                                    _ = len(v2)
                                                                                default:
                                                                                    fmt.Println(v1)
                                                                                }
                                                                                _ = v77
                                                                            } else {
                                                                                if (fi9 != fi11) {
                                                                                    if len(v2) >= 2 {
                                                                                        v2 = v2[1:len(v2)]
                                                                                    }
                                                                                    if (!false) {
                                                                                        if len(v2) > 0 {
                                                                                            v2[0] = 42
                                                                                        }
                                                                                    }
                                                                                    var v83 int64 = (^(v28))
                                                                                    _ = v83
                                                                                } else {
                                                                                    fmt.Println(v3, v1, v28)
                                                                                }
                                                                                v3 = -2.25
                                                                                var v84 int32 = 2147483647
                                                                                _ = v84
                                                                            }
                                                                            _ = v75
                                                                            _ = v76
                                                                        }
                                                                        if false {
                                                                            var v85 uint8 = 128
                                                                            _ = v85
                                                                        } else {
                                                                            v4 = append(v4, fi9)
                                                                        }
                                                                        var v86 int = fi11
                                                                        _ = v86
                                                                    }
                                                                case "":
                                                                    if (v29 < 411) {
                                                                        fmt.Println(v8, v3)
                                                                    } else {
                                                                        v10 -= (1.5 + (v10 - 1e10))
                                                                        if (true && (float32(((1e10 / (v3 + -2.25)) / (float64(1.5) + v3))) != v5)) {
                                                                            if ((fi11 / (((fi11 + fi9)) | 1)) < (fi11 ^ -446)) {
                                                                                fmt.Println(v8, fi11, v28)
                                                                                fmt.Println(fi11)
                                                                                v87 := GenS11{F0: v28}
                                                                                _ = v87
                                                                            } else {
                                                                                v5 = v5
                                                                                v88 := GenS9{F0: 42, F1: (!false), F2: v29}
                                                                                var v89 string = (v1 + v1)
                                                                                _ = v88
                                                                                _ = v89
                                                                            }
                                                                            if (2 <= 42) {
                                                                                var v90 int16 = -853
                                                                                _ = len(v4)
                                                                                _ = v90
                                                                            }
                                                                        }
                                                                    }
                                                                case "\n\t":
                                                                    v4 = append(v4, (fi9 + (fi9 / ((fi11) | 1))))
                                                                default:
                                                                    if (!false) {
                                                                        if len(v2) > 0 {
                                                                            v2[0] = 2147483647
                                                                        }
                                                                    }
                                                                }
                                                                if true {
                                                                    fmt.Println(fi11)
                                                                } else {
                                                                    v91 := GenS12{F0: 647, F1: -128, F2: -32768}
                                                                    switch ("foo bar" + (v1 + (v1 + ("hello" + (v1 + v1))))) {
                                                                    case "foo bar":
                                                                        var v92 string = "a"
                                                                        _ = v92
                                                                    case "":
                                                                        v28 = 1
                                                                    case "a":
                                                                        fmt.Println(v1)
                                                                    default:
                                                                        v5 += v10
                                                                    }
                                                                    _ = v91
                                                                }
                                                                _ = v71
                                                            } else {
                                                                if true {
                                                                    var v93 uint = 0
                                                                    if true {
                                                                        v94 := GenS4{F0: 417, F1: (fi9 * fi9)}
                                                                        if (-32768 < -1) {
                                                                            var v95 uint8 = 127
                                                                            _ = v95
                                                                        } else {
                                                                            v96 := GenS1{F0: 128, F1: -84}
                                                                            var v97 uint32 = (^(v29))
                                                                            _ = v96
                                                                            _ = v97
                                                                        }
                                                                        _ = v94
                                                                    } else {
                                                                        var v98 bool = (v8 == v8)
                                                                        fmt.Println(fi9, v10, v5)
                                                                        _ = v98
                                                                    }
                                                                    var v99 bool = (v8 >= (v8 >> ((1) & 63)))
                                                                    _ = v93
                                                                    _ = v99
                                                                } else {
                                                                    if len(v2) > 0 {
                                                                        v2[0] = 591
                                                                    }
                                                                    _ = len(v4)
                                                                    v5 -= 0.1
                                                                }
                                                            }
                                                            fmt.Println(fi9, v29)
                                                        }
                                                        fmt.Println(fi9)
                                                    } else {
                                                        if len(v4) >= 2 {
                                                            v4 = v4[1:len(v4)]
                                                        }
                                                        var v100 int8 = -20
                                                        if (false || (-660 <= int32(fi9))) {
                                                            _ = len(v4)
                                                        }
                                                        _ = v100
                                                    }
                                                }
                                            case "hello":
                                                v29 += v29
                                            case "foo bar":
                                                v101 := GenS10{F0: 426}
                                                _ = v101
                                            }
                                            var v102 float32 = v10
                                            _ = v102
                                        default:
                                            if (((v29 - (v29 << ((255) & 63))) <= (v29 / (((v29 >> ((172) & 63))) | 1))) || true) {
                                                switch (^(fi9)) {
                                                case -691:
                                                    if ((v1 + ((v1 + v1) + v1)) <= "foo bar") {
                                                        var v103 float32 = 3.14159
                                                        _ = v103
                                                    }
                                                case 32768:
                                                    if true {
                                                        var v104 float64 = ((((((((((v3 / ((v3 / v3) - -2.25)) + float64(-2.25)) - float64(3.14159)) - (v3 - ((((((0.0 * (v3 / v3)) + ((v3 * (v3 / v3)) + ((v3 + ((0.1 - (1.5 + ((v3 - v3) * float64(v10)))) - ((v3 + -2.25) + v3))) * float64(-2.25)))) / 0.1) + float64((-2.25 + v5))) - v3) / -2.25))) * v3) / v3) + -2.25) * float64(v10)) + ((((v3 + v3) + v3) / 1.0) * (-2.25 * (v3 * 1.5)))) + (float64(3.14159) - (((((float64(float32(v3)) / (v3 + v3)) + v3) + v3) / (float64(v10) / v3)) - (0.1 + v3))))
                                                        switch (v1 + (v1 + (v1 + "foo bar"))) {
                                                        case "foo bar":
                                                            v28 = (v28 & (v28 | (v28 / ((v28) | 1))))
                                                        default:
                                                            if len(v4) > 0 {
                                                                v4[0] = (fi9 << ((2) & 63))
                                                            }
                                                        }
                                                        _ = v104
                                                    } else {
                                                        var v105 int = 9223372036854775807
                                                        v28 = -9223372036854775808
                                                        v28 = 0
                                                        fmt.Println(v5, v5, v3)
                                                        v3 = v3
                                                        _ = v105
                                                    }
                                                case 4294967295:
                                                    var v106 uint64 = 173
                                                    _ = v106
                                                }
                                            } else {
                                                var v107 uint64 = 407
                                                if len(v2) > 0 {
                                                    v2[0] = (v29 ^ (v29 - 42))
                                                }
                                                if false {
                                                    var v108 int32 = 1
                                                    v5 -= v5
                                                    _ = v108
                                                } else {
                                                    if (false || false) {
                                                        if (!(!(886 != 4294967295))) {
                                                            if false {
                                                                var v109 int = fi9
                                                                _ = v109
                                                            } else {
                                                                var v110 int64 = (v28 ^ v28)
                                                                var v111 float32 = v10
                                                                var v112 int64 = -9223372036854775808
                                                                _ = v110
                                                                _ = v111
                                                                _ = v112
                                                            }
                                                            _ = len(v4)
                                                            v28 = v28
                                                        }
                                                    }
                                                    var v113 uint64 = (v107 * 2)
                                                    if ((!((!(!(42 < -361))) || true)) || (!((42 == -43) || true))) {
                                                        fmt.Println(v29)
                                                        v113 += (v107 % ((0) | 1))
                                                        switch fi11 {
                                                        case 4294967295:
                                                            v29 += (^(v29))
                                                        default:
                                                            v3 -= v3
                                                            v29 = (^(v29))
                                                            var v114 int8 = -128
                                                            _ = v114
                                                        }
                                                    }
                                                    v107 += (v113 * (^(v113)))
                                                    if (false || ((v1 + v1) > v1)) {
                                                        v2 = append(v2, v29)
                                                    }
                                                    _ = v113
                                                }
                                                _ = v107
                                            }
                                        }
                                    case 127:
                                        v4 = append(v4, fi11)
                                    default:
                                        v28 += int64(v29)
                                        switch (fi11 ^ fi9) {
                                        case 0:
                                            switch (fi11 << ((255) & 63)) {
                                            case 128:
                                                var v115 string = v1
                                                _ = v115
                                            case 1:
                                                v29 = (^(v29))
                                            case 9223372036854775807:
                                                switch 1 {
                                                case 127:
                                                    if (false || ((v8 % (((v8 % (((v8 % ((0) | 1))) | 1))) | 1)) < 42)) {
                                                        if false {
                                                            var v116 int32 = 636
                                                            v5 += (((v10 * v5) / float32(float64(v10))) + v5)
                                                            var v117 int8 = 5
                                                            if (2 <= 180) {
                                                                v4 = append(v4, (-(fi11)))
                                                                fmt.Println(v8, fi11, v117)
                                                            } else {
                                                                var v118 float64 = 3.14159
                                                                v119 := GenS13{F0: (v1 + v1)}
                                                                if true {
                                                                    v120 := GenS3{F0: (v8 + (v8 % ((v8) | 1)))}
                                                                    v121 := GenS4{F0: 18446744073709551615, F1: (fi9 << ((uint(v29)) & 63))}
                                                                    _ = v120
                                                                    _ = v121
                                                                } else {
                                                                    if true {
                                                                        switch fi9 {
                                                                        case 0:
                                                                            var v122 int8 = 42
                                                                            _ = v122
                                                                        default:
                                                                            switch 128 {
                                                                            case 127:
                                                                                var v123 uint = 785
                                                                                _ = v123
                                                                            case -1:
                                                                                var v124 int64 = (-(v28))
                                                                                _ = v124
                                                                            case -484:
                                                                                if (!false) {
                                                                                    var v125 bool = (!(false && true))
                                                                                    _ = v125
                                                                                } else {
                                                                                    var v126 int16 = 258
                                                                                    _ = v126
                                                                                }
                                                                            }
                                                                        }
                                                                    }
                                                                }
                                                                _ = v118
                                                                _ = v119
                                                            }
                                                            _ = v116
                                                            _ = v117
                                                        } else {
                                                            switch (v1 + "foo bar") {
                                                            case "hello":
                                                                if (1 != 0) {
                                                                    if ((-493 <= -769) || true) {
                                                                        v127 := GenS14{F0: 1, F1: 796}
                                                                        var v128 int = 0
                                                                        _ = v127
                                                                        _ = v128
                                                                    } else {
                                                                        fmt.Println(v29, v28, fi9)
                                                                        var v129 uint64 = 4294967295
                                                                        v130 := GenS15{F0: 1000, F1: (^(v8)), F2: (v8 - v8)}
                                                                        _ = v129
                                                                        _ = v130
                                                                    }
                                                                    v28 -= v28
                                                                    v10 = v10
                                                                }
                                                            }
                                                            switch (v1 + (v1 + v1)) {
                                                            case "a":
                                                                if len(v2) >= 2 {
                                                                    v2 = v2[1:len(v2)]
                                                                }
                                                            case "foo bar":
                                                                var v131 float64 = 1.5
                                                                _ = v131
                                                            case "":
                                                                v10 += 0.1
                                                            }
                                                        }
                                                    } else {
                                                        if (!false) {
                                                            var v132 uint16 = v8
                                                            var v133 bool = (((fi9 - (^(fi11))) > (fi11 & -567)) && false)
                                                            switch fi11 {
                                                            case 9223372036854775807:
                                                                var v134 uint8 = 42
                                                                _ = v134
                                                            default:
                                                                var v135 uint32 = 2147483647
                                                                _ = v135
                                                            }
                                                            switch (^(fi9)) {
                                                            case 1:
                                                                if false {
                                                                    if (2147483647 == -127) {
                                                                        if (!(!((478 >= uint64(v8)) && v133))) {
                                                                            if (v29 == v29) {
                                                                                var v136 string = "\n\t"
                                                                                _ = v136
                                                                            } else {
                                                                                if (v8 >= v8) {
                                                                                    if (v133 || (int16(fi11) >= -1)) {
                                                                                        if (-457 >= 42) {
                                                                                            var v137 float32 = (v5 - 1e10)
                                                                                            v138 := GenS16{F0: (v28 / ((v28) | 1)), F1: 1}
                                                                                            _ = v137
                                                                                            _ = v138
                                                                                        }
                                                                                        if (-2147483648 <= 1) {
                                                                                            fmt.Println(v1, v29)
                                                                                            v132 = v8
                                                                                        } else {
                                                                                            fmt.Println(v28)
                                                                                            switch v1 {
                                                                                            case "\n\t":
                                                                                                if (!v133) {
                                                                                                    switch v1 {
                                                                                                    case "foo bar":
                                                                                                        v139 := GenS16{F0: (^(v28)), F1: 0}
                                                                                                        _ = v139
                                                                                                    default:
                                                                                                        v4 = append(v4, (fi11 & (fi11 % ((fi11) | 1))))
                                                                                                    }
                                                                                                    fmt.Println(fi11, v10)
                                                                                                    var v140 int64 = (v28 ^ v28)
                                                                                                    var v141 uint8 = 2
                                                                                                    _ = v140
                                                                                                    _ = v141
                                                                                                } else {
                                                                                                    var v142 int8 = 100
                                                                                                    if v133 {
                                                                                                        if true {
                                                                                                            var v143 uint32 = (v29 % ((2147483647) | 1))
                                                                                                            _ = len(v2)
                                                                                                            _ = v143
                                                                                                        }
                                                                                                        switch "a" {
                                                                                                        case "hello":
                                                                                                            fmt.Println(v133, v29)
                                                                                                        case "":
                                                                                                            if len(v4) > 0 {
                                                                                                                v4[0] = fi11
                                                                                                            }
                                                                                                            switch "" {
                                                                                                            case "\n\t":
                                                                                                                var v144 uint16 = v8
                                                                                                                _ = v144
                                                                                                            case "foo bar":
                                                                                                                if (((v1 + "hello") + "\n\t") >= "") {
                                                                                                                    if (v132 != (v8 | (^(v8)))) {
                                                                                                                        fmt.Println(v1, v10)
                                                                                                                    }
                                                                                                                    v145 := GenS2{F0: 1, F1: 1}
                                                                                                                    if len(v2) >= 2 {
                                                                                                                        v2 = v2[1:len(v2)]
                                                                                                                    }
                                                                                                                    _ = v145
                                                                                                                } else {
                                                                                                                    v146 := GenS17{F0: 228, F1: true}
                                                                                                                    v1 = ((v1 + v1) + ((v1 + "") + v1))
                                                                                                                    _ = v146
                                                                                                                }
                                                                                                                if len(v4) > 0 {
                                                                                                                    v4[0] = (fi11 & fi11)
                                                                                                                }
                                                                                                                var v147 float32 = (v5 * v10)
                                                                                                                _ = v147
                                                                                                            case "":
                                                                                                                v10 -= (0.1 * ((v5 - v5) + (v5 / 0.1)))
                                                                                                            }
                                                                                                            v3 -= 1.5
                                                                                                        case "\n\t":
                                                                                                            if v133 {
                                                                                                                var v148 uint = 716
                                                                                                                v5 = 3.14159
                                                                                                                switch (((((((v1 + v1) + "") + ("hello" + ("" + ((v1 + "\n\t") + "")))) + ((v1 + v1) + "foo bar")) + (v1 + "a")) + (v1 + ("hello" + v1))) + v1) {
                                                                                                                case "foo bar":
                                                                                                                    var v149 uint64 = 1
                                                                                                                    _ = v149
                                                                                                                default:
                                                                                                                    v2 = append(v2, v29)
                                                                                                                }
                                                                                                                var v150 bool = (v148 == v148)
                                                                                                                switch 127 {
                                                                                                                case 127:
                                                                                                                    v151 := GenS11{F0: v28}
                                                                                                                    _ = v151
                                                                                                                case 4294967295:
                                                                                                                    v3 += 3.14159
                                                                                                                default:
                                                                                                                    v3 -= (float64((0.0 * 1e10)) + v3)
                                                                                                                }
                                                                                                                _ = v148
                                                                                                                _ = v150
                                                                                                            }
                                                                                                        default:
                                                                                                            var v152 uint = 438
                                                                                                            _ = v152
                                                                                                        }
                                                                                                        v153 := GenS18{F0: 65535, F1: ""}
                                                                                                        _ = v153
                                                                                                    } else {
                                                                                                        if true {
                                                                                                            if (!v133) {
                                                                                                                if v133 {
                                                                                                                    if v133 {
                                                                                                                        if len(v2) >= 2 {
                                                                                                                            v2 = v2[1:len(v2)]
                                                                                                                        }
                                                                                                                    } else {
                                                                                                                        var v154 string = "\n\t"
                                                                                                                        _ = v154
                                                                                                                    }
                                                                                                                    if true {
                                                                                                                        v28 = v28
                                                                                                                        _ = len(v4)
                                                                                                                        var v155 float32 = 0.0
                                                                                                                        _ = v155
                                                                                                                    } else {
                                                                                                                        var v156 float64 = v3
                                                                                                                        _ = v156
                                                                                                                    }
                                                                                                                }
                                                                                                                v28 += (v28 << ((4294967295) & 63))
                                                                                                                fmt.Println(v142, fi9, v29)
                                                                                                            }
                                                                                                            var v157 int = fi9
                                                                                                            v133 = ((^(v142)) != v142)
                                                                                                            if v133 {
                                                                                                                v3 = (0.0 - v3)
                                                                                                                v1 = (v1 + (v1 + v1))
                                                                                                            } else {
                                                                                                                fmt.Println(v3, v10)
                                                                                                                if v133 {
                                                                                                                    switch (-(v157)) {
                                                                                                                    case 42:
                                                                                                                        switch fi9 {
                                                                                                                        case 128:
                                                                                                                            v158 := GenS19{F0: v133, F1: 4294967295}
                                                                                                                            _ = v158
                                                                                                                        case 42:
                                                                                                                            _ = len(v2)
                                                                                                                        }
                                                                                                                        v132 += 255
                                                                                                                    }
                                                                                                                    if (47 < v132) {
                                                                                                                        v10 -= v10
                                                                                                                        if len(v2) >= 2 {
                                                                                                                            v2 = v2[1:len(v2)]
                                                                                                                        }
                                                                                                                        fmt.Println(v29)
                                                                                                                        v29 = 0
                                                                                                                    }
                                                                                                                }
                                                                                                            }
                                                                                                            var v159 float64 = v3
                                                                                                            _ = v157
                                                                                                            _ = v159
                                                                                                        } else {
                                                                                                            if true {
                                                                                                                fmt.Println(v132, v133)
                                                                                                                var v160 uint64 = 4294967295
                                                                                                                v161 := GenS4{F0: 1, F1: fi9}
                                                                                                                _ = v160
                                                                                                                _ = v161
                                                                                                            } else {
                                                                                                                v162 := GenS2{F0: 64, F1: 42}
                                                                                                                _ = v162
                                                                                                            }
                                                                                                        }
                                                                                                    }
                                                                                                    _ = len(v4)
                                                                                                    _ = v142
                                                                                                }
                                                                                            case "":
                                                                                                v163 := GenS5{F0: v29}
                                                                                                _ = v163
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                    _ = len(v2)
                                                                                    var v164 float64 = v3
                                                                                    _ = v164
                                                                                }
                                                                            }
                                                                            if v133 {
                                                                                v29 = v29
                                                                                v133 = (v3 <= v3)
                                                                            }
                                                                            switch fi11 {
                                                                            case 4294967295:
                                                                                v3 -= v3
                                                                            default:
                                                                                var v165 int = fi11
                                                                                _ = v165
                                                                            }
                                                                        }
                                                                        switch fi11 {
                                                                        case -9223372036854775808:
                                                                            if (v133 || v133) {
                                                                                var v166 int8 = 127
                                                                                _ = v166
                                                                            }
                                                                        default:
                                                                            v167 := GenS13{F0: ""}
                                                                            _ = v167
                                                                        }
                                                                    }
                                                                }
                                                            case 2147483647:
                                                                if v133 {
                                                                    v168 := GenS13{F0: v1}
                                                                    _ = v168
                                                                } else {
                                                                    var v169 bool = false
                                                                    fmt.Println(v8, v3, fi9)
                                                                    switch "" {
                                                                    case "foo bar":
                                                                        if v133 {
                                                                            var v170 uint64 = 0
                                                                            var v171 string = v1
                                                                            v132 = (v132 & (v8 ^ (^(v8))))
                                                                            _ = v170
                                                                            _ = v171
                                                                        }
                                                                        if (42 < -32768) {
                                                                            v172 := GenS20{F0: v28, F1: false}
                                                                            var v173 float64 = v3
                                                                            _ = v172
                                                                            _ = v173
                                                                        }
                                                                        fmt.Println(v132, v8, v28)
                                                                    case "a":
                                                                        v3 = (0.1 - 1.5)
                                                                    case "\n\t":
                                                                        var v174 int = 1
                                                                        _ = v174
                                                                    default:
                                                                        v28 += v28
                                                                    }
                                                                    _ = v169
                                                                }
                                                            }
                                                            switch (v1 + "") {
                                                            case "foo bar":
                                                                var v175 int16 = -276
                                                                fmt.Println(v3, v8)
                                                                switch v1 {
                                                                case "foo bar":
                                                                    var v176 float64 = (1e10 * (((((3.14159 / (0.1 - float64(0.0))) / (v3 / (v3 + 3.14159))) - (float64(v10) * 1e10)) - (0.1 - v3)) - float64((0.1 / (0.1 - ((float32((1.5 - -2.25)) * v5) / (v10 - ((v5 + 1.5) + (v10 - v5)))))))))
                                                                    _ = v176
                                                                default:
                                                                    if true {
                                                                        if true {
                                                                            v177 := GenS19{F0: false, F1: v29}
                                                                            _ = v177
                                                                        } else {
                                                                            v1 = (v1 + "foo bar")
                                                                            v178 := GenS21{F0: v1}
                                                                            _ = v178
                                                                        }
                                                                        switch "\n\t" {
                                                                        case "foo bar":
                                                                            if false {
                                                                                _ = len(v4)
                                                                                var v179 float64 = (0.0 * v3)
                                                                                if false {
                                                                                    fmt.Println(v5, v175, v10)
                                                                                    v3 -= v179
                                                                                    if false {
                                                                                        var v180 int64 = (^(v28))
                                                                                        var v181 int = (fi11 ^ (fi9 & 42))
                                                                                        _ = len(v2)
                                                                                        if len(v4) >= 2 {
                                                                                            v4 = v4[1:len(v4)]
                                                                                        }
                                                                                        _ = v180
                                                                                        _ = v181
                                                                                    } else {
                                                                                        v10 = v10
                                                                                        fmt.Println(v10, v3)
                                                                                        if v133 {
                                                                                            v2 = append(v2, 50)
                                                                                            var v182 float32 = v10
                                                                                            _ = len(v4)
                                                                                            _ = v182
                                                                                        }
                                                                                        if v133 {
                                                                                            v175 = v175
                                                                                            var v183 int16 = (v175 % ((-32768) | 1))
                                                                                            _ = v183
                                                                                        }
                                                                                    }
                                                                                } else {
                                                                                    fmt.Println(v5, fi11)
                                                                                }
                                                                                _ = v179
                                                                            } else {
                                                                                if ((fi11 << ((4294967295) & 63)) >= fi11) {
                                                                                    var v184 string = v1
                                                                                    var v185 float64 = (v3 * v3)
                                                                                    var v186 int = fi9
                                                                                    _ = v184
                                                                                    _ = v185
                                                                                    _ = v186
                                                                                }
                                                                                fmt.Println(v1, v5, v8)
                                                                            }
                                                                        case "hello":
                                                                            var v187 int8 = 127
                                                                            _ = v187
                                                                        case "a":
                                                                            switch (fi11 ^ (fi9 << ((255) & 63))) {
                                                                            case -1:
                                                                                v188 := GenS12{F0: (^(v29)), F1: 13, F2: (^(v175))}
                                                                                _ = v188
                                                                            default:
                                                                                v189 := GenS7{F0: v8, F1: 4294967295, F2: v175}
                                                                                _ = v189
                                                                            }
                                                                        }
                                                                        var v190 float64 = 1.5
                                                                        _ = v190
                                                                    }
                                                                }
                                                                _ = v175
                                                            }
                                                            _ = v132
                                                            _ = v133
                                                        }
                                                        var v191 uint8 = 42
                                                        _ = v191
                                                    }
                                                case -1:
                                                    v192 := GenS22{F0: v1}
                                                    _ = v192
                                                case -485:
                                                    if (v28 != (v28 * v28)) {
                                                        v10 = float32((float64(0.0) + (-2.25 + v3)))
                                                        _ = len(v2)
                                                    }
                                                }
                                            default:
                                                if false {
                                                    var v193 int32 = 1
                                                    v194 := GenS23{F0: (v8 / ((v8) | 1)), F1: (v193 ^ v193), F2: 1}
                                                    _ = v193
                                                    _ = v194
                                                } else {
                                                    if ((!true) || false) {
                                                        v5 += (3.14159 + (-2.25 + v10))
                                                        fmt.Println(v29, v5, v5)
                                                    }
                                                }
                                            }
                                        case 127:
                                            fmt.Println(v29)
                                            v195 := GenS24{F0: 614, F1: 42}
                                            _ = v195
                                        case -1:
                                            fmt.Println(fi11, v3)
                                        default:
                                            if (!(!(false && (0 < -795)))) {
                                                v196 := GenS21{F0: ((((v1 + (v1 + v1)) + v1) + ((v1 + v1) + (v1 + v1))) + "\n\t")}
                                                switch -1 {
                                                case -31:
                                                    var v197 int8 = 100
                                                    _ = v197
                                                case 0:
                                                    _ = len(v4)
                                                default:
                                                    v8 = 2
                                                }
                                                if (71 == -128) {
                                                    var v198 int = -447
                                                    _ = v198
                                                } else {
                                                    var v199 int64 = 9223372036854775807
                                                    _ = v199
                                                }
                                                _ = v196
                                            }
                                        }
                                    }
                                }
                            case 53:
                                v8 = (v8 << ((227) & 63))
                            default:
                                v29 += (^(v29))
                            }
                            if false {
                                v1 = "foo bar"
                                if (((2 > 2) || ((!(!true)) && (-128 < 1))) || true) {
                                    if false {
                                        var v200 int32 = 762
                                        _ = v200
                                    } else {
                                        switch (v1 + v1) {
                                        case "\n\t":
                                            if false {
                                                if (!false) {
                                                    v201 := GenS3{F0: v8}
                                                    _ = len(v2)
                                                    _ = v201
                                                }
                                            }
                                        default:
                                            v29 = 579
                                        }
                                    }
                                }
                                var v202 int8 = int8(v28)
                                _ = v202
                            }
                            switch v1 {
                            case "hello":
                                switch 726 {
                                case 0:
                                    if (!false) {
                                        if ((3.14159 == v3) || ((float64(v5) / v3) != v3)) {
                                            if true {
                                                v3 = (0.0 - float64(float32((v3 - 0.0))))
                                            } else {
                                                if (!false) {
                                                    if (false || (-2147483648 != 1)) {
                                                        fmt.Println(fi11)
                                                        if (-425 != 727) {
                                                            _ = len(v4)
                                                            if (0.0 == (0.0 * v3)) {
                                                                fmt.Println(v28, v5)
                                                            }
                                                            var v203 int16 = 32767
                                                            switch ("" + "foo bar") {
                                                            case "\n\t":
                                                                var v204 uint = 2
                                                                _ = v204
                                                            case "":
                                                                v205 := GenS1{F0: 255, F1: 23}
                                                                _ = v205
                                                            default:
                                                                var v206 uint32 = v29
                                                                v1 = ((v1 + "a") + "a")
                                                                _ = len(v4)
                                                                _ = v206
                                                            }
                                                            _ = v203
                                                        }
                                                        _ = len(v4)
                                                    } else {
                                                        v29 = (v29 / (((^(v29))) | 1))
                                                        if ((!false) && true) {
                                                            v207 := GenS12{F0: uint32(v8), F1: 127, F2: 0}
                                                            v3 = (0.1 + v3)
                                                            switch v1 {
                                                            case "\n\t":
                                                                v208 := GenS17{F0: 0, F1: (0 >= 1)}
                                                                _ = v208
                                                            case "":
                                                                var v209 uint32 = 4294967295
                                                                _ = v209
                                                            default:
                                                                v210 := GenS23{F0: (v8 & v8), F1: 868, F2: 2}
                                                                _ = v210
                                                            }
                                                            _ = v207
                                                        } else {
                                                            var v211 bool = (!true)
                                                            _ = v211
                                                        }
                                                        switch (fi9 << ((2) & 63)) {
                                                        case -1:
                                                            v212 := GenS2{F0: uint8(v8), F1: 102}
                                                            _ = v212
                                                        }
                                                    }
                                                    v5 += v10
                                                } else {
                                                    var v213 int32 = -1
                                                    var v214 string = "hello"
                                                    v4 = append(v4, -1)
                                                    _ = v213
                                                    _ = v214
                                                }
                                                var v215 int64 = 800
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                                _ = v215
                                            }
                                            switch 1 {
                                            case 32768:
                                                v4 = append(v4, 507)
                                            case -2147483648:
                                                var v216 uint32 = (v29 % ((v29) | 1))
                                                _ = v216
                                            }
                                        }
                                        if len(v4) > 0 {
                                            v4[0] = (fi9 & fi11)
                                        }
                                        var v217 uint8 = 35
                                        _ = v217
                                    } else {
                                        v4 = append(v4, (fi9 >> ((65535) & 63)))
                                        var v218 uint16 = v8
                                        if (!(v1 >= v1)) {
                                            switch ((("foo bar" + "\n\t") + "hello") + v1) {
                                            case "a":
                                                if false {
                                                    var v219 uint16 = v8
                                                    var v220 int16 = 101
                                                    _ = len(v2)
                                                    _ = v219
                                                    _ = v220
                                                }
                                            case "":
                                                v218 = (v218 | v218)
                                            default:
                                                v28 = v28
                                                if (!(((3.14159 / 0.1) + (0.0 / float64((v5 * v5)))) != v3)) {
                                                    v3 = (v3 * v3)
                                                    if (404 > 544) {
                                                        if (1 < 661) {
                                                            var v221 uint32 = v29
                                                            _ = v221
                                                        } else {
                                                            fmt.Println(fi11, v29, v8)
                                                        }
                                                    }
                                                } else {
                                                    v222 := GenS25{F0: ("a" + (v1 + ((v1 + v1) + "foo bar"))), F1: 739}
                                                    var v223 uint16 = (v218 / ((158) | 1))
                                                    _ = v222
                                                    _ = v223
                                                }
                                                fmt.Println(v29, fi11)
                                            }
                                        } else {
                                            var v224 int = fi9
                                            _ = v224
                                        }
                                        v8 = v8
                                        _ = v218
                                    }
                                case 1:
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                case 2147483647:
                                    if (false || (2 > 2)) {
                                        var v225 float64 = ((v3 * (float64((1.5 - (1.5 + (1e10 - v5)))) - v3)) + float64(1.5))
                                        _ = v225
                                    }
                                default:
                                    v5 = v10
                                }
                            case "a":
                                v226 := GenS8{F0: true, F1: v1, F2: 42}
                                _ = v226
                            case "":
                                fmt.Println(fi11, v29)
                            }
                            v227 := GenS13{F0: "a"}
                            var v228 uint8 = 255
                            _ = v227
                            _ = v228
                            _ = v28
                            _ = v29
                        }
                        fmt.Println(v8, v8)
                        switch fi9 {
                        case 9223372036854775807:
                            v8 -= (v8 & v8)
                        case -1:
                            v10 -= float32(v3)
                        case 258:
                            var v229 int = 32767
                            _ = v229
                        default:
                            var v230 uint16 = v8
                            _ = v230
                        }
                        v231 := GenS9{F0: 32767, F1: false, F2: 2147483648}
                        _ = v231
                    }
                    v232 := GenS16{F0: 1, F1: 2147483647}
                    _ = len(v2)
                    _ = v232
                }
                _ = fi11
            }
            _ = fi9
            _ = v10
        }
        if false {
            fmt.Println(v8, v1, v5)
            if true {
                var v233 int8 = -37
                v233 = v233
                _ = v233
            } else {
                fmt.Println(v3, v8, v1)
                v4 = append(v4, -1)
                var v234 uint8 = 1
                _ = v234
            }
        } else {
            v235 := GenS13{F0: v1}
            _ = v235
        }
        _ = v8
    } else {
        for fi236 := 0; fi236 < 14; fi236++ {
            switch (fi236 << ((569) & 63)) {
            case 227:
                v3 -= v3
            }
            _ = fi236
        }
        v237 := GenS26{F0: 333, F1: -9223372036854775808, F2: 1}
        if true {
            v238 := GenS3{F0: 979}
            if len(v2) > 0 {
                v2[0] = 2147483647
            }
            _ = len(v4)
            _ = v238
        }
        _ = v237
    }
    var v239 uint8 = 127
    fmt.Println(v239)
    fmt.Println(v5)
    v1 = "a"
    for fi240 := 0; fi240 < 17; fi240++ {
        var v241 int16 = -32768
        _ = fi240
        _ = v241
    }
    v2 = append(v2, 2)
    switch 862 {
    case 32767:
        var v242 bool = true
        _ = v242
    case -9223372036854775808:
        for fi243 := 0; fi243 < 10; fi243++ {
            v244 := GenS27{F0: 2}
            var v245 uint64 = 42
            var v246 int16 = -1
            _ = fi243
            _ = v244
            _ = v245
            _ = v246
        }
    case 127:
        for fi247 := 0; fi247 < 15; fi247++ {
            for fi248 := 0; fi248 < 1; fi248++ {
                var v249 uint8 = v239
                v250 := GenS27{F0: 418}
                _ = fi248
                _ = v249
                _ = v250
            }
            switch 9223372036854775807 {
            case -308:
                var v251 string = (v1 + "hello")
                _ = v251
            case -406:
                v3 = ((v3 + 0.1) / (v3 - (((v3 / (1e10 - ((v3 * -2.25) - -2.25))) * float64(float32((v3 * v3)))) - v3)))
            default:
                var v252 float32 = ((v5 + 1.5) * v5)
                _ = v252
            }
            _ = fi247
        }
    default:
        if false {
            for fi253 := 0; fi253 < 2; fi253++ {
                for fi254 := 0; fi254 < 16; fi254++ {
                    fmt.Println(v3)
                    var v255 string = "hello"
                    _ = fi254
                    _ = v255
                }
                v239 = (^(v239))
                var v256 string = ("foo bar" + (v1 + v1))
                _ = fi253
                _ = v256
            }
            var v257 float32 = v5
            var v258 int8 = 1
            _ = v257
            _ = v258
        }
    }
    for fi259 := 0; fi259 < 19; fi259++ {
        for fi260 := 0; fi260 < 14; fi260++ {
            var v261 int = (fi260 % ((int(v239)) | 1))
            _ = fi260
            _ = v261
        }
        var v262 uint32 = 65535
        fmt.Println(v5, v239)
        _ = fi259
        _ = v262
    }
    if (80 < 77) {
        var v263 uint8 = (^(v239))
        _ = v263
    } else {
        if (false && (!false)) {
            var v264 uint8 = (v239 - v239)
            _ = v264
        } else {
            var v265 int = -2147483648
            v1 = ""
            _ = v265
        }
        fmt.Println(v3)
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v239:", v239)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v239
    _ = v3
    _ = v4
    _ = v5
}
