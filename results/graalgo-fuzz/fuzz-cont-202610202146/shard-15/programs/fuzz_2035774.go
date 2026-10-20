package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint8; F1 int64 }
type GenS3 struct { F0 uint8; F1 uint32; F2 int8 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int32; F1 uint64 }
type GenS6 struct { F0 uint; F1 uint }
type GenS7 struct { F0 int8; F1 uint64; F2 uint16 }
type GenS8 struct { F0 uint32 }

func main() {
    var v1 bool = (!true)
    var v2 bool = false
    var v3 bool = (252 <= -9223372036854775808)
    v4 := make(map[string]int)
    var v5 int64 = 9223372036854775807
    var v6 int32 = 2147483647
    if v2 {
        var v7 uint64 = 255
        v2 = v3
        _ = v7
    }
    var v8 uint = 4294967295
    v9 := []int{-2147483648, 4294967295, 32768}
    v5 = v5
    for fi10 := 0; fi10 < 11; fi10++ {
        if len(v9) > 0 {
            v9[0] = fi10
        }
        v6 = v6
        var v11 uint16 = 255
        _ = fi10
        _ = v11
    }
    var v12 float64 = ((((1e10 - 0.1) * (1e10 / 1.5)) / 1.5) + 1e10)
    v12 = float64(3.14159)
    if (v8 >= v8) {
        for fi13 := 0; fi13 < 1; fi13++ {
            switch "\n\t" {
            case "a":
                if (v2 && (((v5 * (v5 / (((v5 % ((-668) | 1))) | 1))) >= v5) && ((v5 % ((v5) | 1)) < v5))) {
                    v5 = (v5 * (v5 ^ v5))
                    if v3 {
                        for fi14 := 0; fi14 < 4; fi14++ {
                            v15 := v4["foo bar"]
                            var v16 float32 = float32((v12 / v12))
                            fmt.Println(v12, v15)
                            if ((v6 % (((v6 ^ (-(v6)))) | 1)) == 1) {
                                _ = len(v9)
                                if (4294967295 > 0) {
                                    v2 = v1
                                    v15 -= -9223372036854775808
                                    if (((!v2) && (true || (2 < 788))) || (!((^(v8)) != v8))) {
                                        var v17 int = 9223372036854775807
                                        var v18 uint16 = 2
                                        _ = v17
                                        _ = v18
                                    } else {
                                        var v19 int = (^(fi13))
                                        v20 := GenS1{F0: (^(v8))}
                                        _ = v19
                                        _ = v20
                                    }
                                }
                                var v21 float32 = (-2.25 / (v16 / float32((v12 / (float64(v16) / (float64(((3.14159 + v16) - ((-2.25 - float32(v12)) - v16))) + float64(-2.25)))))))
                                _ = v21
                            } else {
                                v22 := GenS1{F0: v8}
                                v16 -= v16
                                var v23 int64 = (v5 % (((v5 | v5)) | 1))
                                _ = v22
                                _ = v23
                            }
                            _ = fi14
                            _ = v15
                            _ = v16
                        }
                    }
                }
            case "":
                if ((v6 != (v6 ^ (v6 & v6))) || (699 == 42)) {
                    var v24 bool = ((2 < 2) || (false && v2))
                    _ = v24
                }
            case "foo bar":
                v9 = append(v9, fi13)
            }
            delete(v4, ("a" + "\n\t"))
            if v1 {
                switch -406 {
                case 32767:
                    if v3 {
                        if (v1 || (!v1)) {
                            delete(v4, (("\n\t" + "foo bar") + ("hello" + (("\n\t" + "foo bar") + "\n\t"))))
                            fmt.Println(v3)
                        }
                        v6 = (v6 >> ((v8) & 63))
                    }
                default:
                    for fi25 := 0; fi25 < 16; fi25++ {
                        switch (fi13 & (fi13 / (((fi25 ^ fi13)) | 1))) {
                        case -9223372036854775808:
                            var v26 uint64 = 288
                            _ = v26
                        case 32767:
                            v8 += v8
                        case 2147483647:
                            v6 = 1
                        default:
                            var v27 int16 = -32768
                            _ = v27
                        }
                        if (true || v2) {
                            v4["foo bar"] = fi25
                            switch 0 {
                            case -578:
                                var v28 int8 = -61
                                _ = v28
                            case 127:
                                if true {
                                    fmt.Println(v6, fi13, v8)
                                } else {
                                    fmt.Println(v6)
                                    v1 = (42 < (v6 / ((v6) | 1)))
                                    v12 = float64((1e10 - 1.5))
                                    v9 = append(v9, fi13)
                                    var v29 uint32 = 42
                                    fmt.Println(v6, v5, v12)
                                    if (-128 != 0) {
                                        if ((-128 > int8(v5)) || (((1.5 * float32((v12 * float64(-2.25)))) / 1.0) >= 1.5)) {
                                            fmt.Println(v12)
                                        }
                                        v30 := v4[("" + ((((("hello" + ("" + "hello")) + ("" + ("hello" + "foo bar"))) + (((("\n\t" + "a") + "hello") + "\n\t") + "foo bar")) + "foo bar") + "\n\t"))]
                                        _ = v30
                                    } else {
                                        var v31 int8 = 1
                                        if len(v9) >= 2 {
                                            v9 = v9[1:len(v9)]
                                        }
                                        _ = v31
                                    }
                                    var v32 string = ("foo bar" + "")
                                    _ = v29
                                    _ = v32
                                }
                            case 42:
                                v8 = (v8 - 255)
                            }
                            if len(v9) > 0 {
                                v9[0] = fi13
                            }
                        } else {
                            var v33 string = "hello"
                            if (fi25 >= -9223372036854775808) {
                                if v1 {
                                    var v34 uint16 = 728
                                    switch (-(fi25)) {
                                    case 1:
                                        switch (^(fi25)) {
                                        case -9223372036854775808:
                                            v8 = v8
                                        case -836:
                                            _ = len(v9)
                                        case 32767:
                                            var v35 uint16 = (v34 & v34)
                                            _ = v35
                                        }
                                    case -9223372036854775808:
                                        fmt.Println(v33, fi13)
                                    }
                                    var v36 int64 = v5
                                    _ = v34
                                    _ = v36
                                } else {
                                    v37 := GenS1{F0: v8}
                                    v5 = v5
                                    v5 -= v5
                                    v8 = v8
                                    if v1 {
                                        if len(v9) >= 2 {
                                            v9 = v9[1:len(v9)]
                                        }
                                        v3 = (!(fi13 <= (fi13 >> (((v8 * (v8 >> ((255) & 63)))) & 63))))
                                    } else {
                                        v2 = v2
                                        switch fi13 {
                                        case 127:
                                            if (v3 && v1) {
                                                fmt.Println(v1, v6, fi13)
                                            }
                                        case 2147483647:
                                            v38 := GenS1{F0: v8}
                                            _ = v38
                                        default:
                                            v8 -= (^(v8))
                                        }
                                    }
                                    _ = v37
                                }
                                v3 = v3
                                switch fi13 {
                                case 215:
                                    v39 := GenS1{F0: v8}
                                    fmt.Println(v1, v6, v1)
                                    _ = v39
                                case 1:
                                    v40 := GenS1{F0: (v8 >> (((v8 - v8)) & 63))}
                                    _ = v40
                                default:
                                    if true {
                                        v8 -= (v8 - 18446744073709551615)
                                        fmt.Println(v8, v3, fi25)
                                        if v3 {
                                            v3 = (uint8(v8) == 1)
                                            if v3 {
                                                if v3 {
                                                    v33 = (v33 + v33)
                                                    v6 += v6
                                                    if true {
                                                        v33 = v33
                                                        fmt.Println(v5, v6, v12)
                                                        v5 = v5
                                                        fmt.Println(fi13, v12, v2)
                                                    }
                                                    if ((3.14159 == v12) && (fi13 > (^(fi13)))) {
                                                        var v41 int8 = -1
                                                        var v42 float32 = 0.0
                                                        if len(v9) > 0 {
                                                            v9[0] = -359
                                                        }
                                                        _ = v41
                                                        _ = v42
                                                    }
                                                } else {
                                                    switch "foo bar" {
                                                    case "foo bar":
                                                        v5 = (v5 & v5)
                                                    case "hello":
                                                        v6 += v6
                                                    default:
                                                        v43 := GenS1{F0: (v8 << ((v8) & 63))}
                                                        _ = v43
                                                    }
                                                    switch fi25 {
                                                    case -474:
                                                        switch "a" {
                                                        case "hello":
                                                            if (-500 <= 1) {
                                                                fmt.Println(v3, v6)
                                                                fmt.Println(v12, v8, v12)
                                                                if len(v9) > 0 {
                                                                    v9[0] = (fi25 >> (((v8 + 0)) & 63))
                                                                }
                                                            }
                                                        case "\n\t":
                                                            v1 = v1
                                                        case "a":
                                                            if len(v9) >= 2 {
                                                                v9 = v9[1:len(v9)]
                                                            }
                                                        default:
                                                            switch fi13 {
                                                            case 2147483647:
                                                                v33 = v33
                                                                v6 -= int32(v8)
                                                                var v44 uint16 = 32767
                                                                _ = v44
                                                            case 401:
                                                                v45 := GenS1{F0: (v8 | 18446744073709551615)}
                                                                v46 := GenS1{F0: (^(v8))}
                                                                _ = v45
                                                                _ = v46
                                                            default:
                                                                _ = len(v9)
                                                            }
                                                        }
                                                    case -1:
                                                        var v47 uint8 = 128
                                                        _ = v47
                                                    default:
                                                        switch v33 {
                                                        case "\n\t":
                                                            if v1 {
                                                                var v48 bool = true
                                                                v49 := v4[((v33 + (v33 + v33)) + "foo bar")]
                                                                _ = len(v9)
                                                                if (50 < 65535) {
                                                                    fmt.Println(v1)
                                                                    var v50 int8 = -17
                                                                    if len(v9) >= 2 {
                                                                        v9 = v9[1:len(v9)]
                                                                    }
                                                                    v51 := GenS1{F0: (^(v8))}
                                                                    var v52 uint = (^(v8))
                                                                    _ = v50
                                                                    _ = v51
                                                                    _ = v52
                                                                }
                                                                v2 = (!v48)
                                                                _ = v48
                                                                _ = v49
                                                            } else {
                                                                if (v2 || true) {
                                                                    var v53 uint8 = 2
                                                                    if v3 {
                                                                        var v54 int64 = (v5 << (((^(v8))) & 63))
                                                                        var v55 uint16 = 832
                                                                        _ = v54
                                                                        _ = v55
                                                                    }
                                                                    v53 -= v53
                                                                    _ = v53
                                                                }
                                                                fmt.Println(v33)
                                                                v56 := GenS1{F0: v8}
                                                                var v57 bool = (v12 < 1.5)
                                                                v9 = append(v9, (fi13 & 381))
                                                                _ = v56
                                                                _ = v57
                                                            }
                                                        case "":
                                                            var v58 int16 = 672
                                                            _ = v58
                                                        case "foo bar":
                                                            v59 := GenS2{F0: 42, F1: v5}
                                                            var v60 int16 = 592
                                                            if v2 {
                                                                if ("\n\t" < (v33 + (v33 + (v33 + ("\n\t" + v33))))) {
                                                                    if (1 != 4294967295) {
                                                                        v8 = 255
                                                                        v61 := GenS3{F0: 0, F1: 65535, F2: 127}
                                                                        var v62 bool = (((3.14159 - float64(1.5)) + v12) > float64(3.14159))
                                                                        fmt.Println(v62, v3, v60)
                                                                        fmt.Println(v62, v2, fi25)
                                                                        _ = v61
                                                                        _ = v62
                                                                    } else {
                                                                        _ = len(v9)
                                                                        var v63 bool = ((!v2) && ((!true) || (v6 == (^(v6)))))
                                                                        _ = v63
                                                                    }
                                                                    var v64 int32 = (v6 ^ int32(v8))
                                                                    _ = v64
                                                                } else {
                                                                    delete(v4, v33)
                                                                    fmt.Println(v6, v33, fi13)
                                                                    v65 := GenS3{F0: 211, F1: 4294967295, F2: 1}
                                                                    _ = v65
                                                                }
                                                            } else {
                                                                if (42 < 1) {
                                                                    v8 = v8
                                                                    if len(v9) > 0 {
                                                                        v9[0] = (fi25 | (fi13 - (fi25 - fi13)))
                                                                    }
                                                                }
                                                            }
                                                            _ = v59
                                                            _ = v60
                                                        default:
                                                            if (2 != 2) {
                                                                var v66 uint64 = 159
                                                                fmt.Println(v5, v5, v12)
                                                                v9 = append(v9, int(v5))
                                                                var v67 int64 = v5
                                                                _ = v66
                                                                _ = v67
                                                            } else {
                                                                switch 4294967295 {
                                                                case 2147483647:
                                                                    if false {
                                                                        var v68 string = (((("foo bar" + "a") + "foo bar") + v33) + v33)
                                                                        _ = v68
                                                                    } else {
                                                                        v69 := GenS1{F0: 42}
                                                                        _ = len(v9)
                                                                        if (255 < 255) {
                                                                            var v70 int64 = v5
                                                                            v5 = (v70 * (v5 ^ -9223372036854775808))
                                                                            _ = v70
                                                                        } else {
                                                                            if (true || v3) {
                                                                                v5 += (v5 + (v5 ^ (^(v5))))
                                                                                fmt.Println(v2, v1)
                                                                            }
                                                                        }
                                                                        _ = v69
                                                                    }
                                                                    var v71 uint32 = 389
                                                                    _ = v71
                                                                case 487:
                                                                    v72 := GenS4{F0: (^(fi25))}
                                                                    _ = v72
                                                                }
                                                            }
                                                        }
                                                    }
                                                    switch 1 {
                                                    case 1:
                                                        switch fi13 {
                                                        case 669:
                                                            var v73 int32 = v6
                                                            _ = v73
                                                        case -9223372036854775808:
                                                            if len(v9) >= 2 {
                                                                v9 = v9[1:len(v9)]
                                                            }
                                                        default:
                                                            v5 += v5
                                                        }
                                                        delete(v4, v33)
                                                        fmt.Println(v33)
                                                    case 0:
                                                        switch "" {
                                                        case "":
                                                            switch fi25 {
                                                            case 127:
                                                                fmt.Println(v8)
                                                            case -220:
                                                                v74 := GenS5{F0: (v6 - 1), F1: 1}
                                                                _ = v74
                                                            default:
                                                                var v75 float64 = v12
                                                                _ = v75
                                                            }
                                                        case "a":
                                                            if (201 >= 255) {
                                                                var v76 int32 = (v6 >> ((114) & 63))
                                                                _ = v76
                                                            } else {
                                                                _ = len(v4)
                                                                v77 := GenS4{F0: fi13}
                                                                v33 = "\n\t"
                                                                _ = v77
                                                            }
                                                        }
                                                    }
                                                }
                                                var v78 int64 = (v5 ^ (v5 >> ((v8) & 63)))
                                                _ = v78
                                            } else {
                                                _ = len(v4)
                                            }
                                            if len(v9) >= 2 {
                                                v9 = v9[1:len(v9)]
                                            }
                                        } else {
                                            var v79 int = fi25
                                            switch -6 {
                                            case -838:
                                                v80 := GenS4{F0: -729}
                                                _ = v80
                                            case 293:
                                                v12 = v12
                                            case 127:
                                                var v81 uint64 = 18446744073709551615
                                                _ = v81
                                            default:
                                                v9 = append(v9, (fi13 % ((4294967295) | 1)))
                                            }
                                            if len(v9) >= 2 {
                                                v9 = v9[1:len(v9)]
                                            }
                                            _ = v79
                                        }
                                    }
                                }
                                v82 := GenS6{F0: (v8 | (v8 * v8)), F1: 2}
                                _ = v82
                            }
                            fmt.Println(v1)
                            _ = v33
                        }
                        var v83 int32 = v6
                        _ = fi25
                        _ = v83
                    }
                }
                var v84 int8 = 98
                for fi85 := 0; fi85 < 6; fi85++ {
                    switch "hello" {
                    case "hello":
                        v86 := GenS5{F0: (v6 | v6), F1: 18446744073709551615}
                        _ = v86
                    case "\n\t":
                        var v87 uint = 636
                        _ = v87
                    default:
                        v5 = -734
                    }
                    _ = fi85
                }
                _ = v84
            } else {
                _ = len(v9)
                for fi88 := 0; fi88 < 15; fi88++ {
                    var v89 int16 = -1
                    v3 = ((v89 + (v89 % ((v89) | 1))) >= (^(v89)))
                    v90 := GenS7{F0: -38, F1: 2, F2: 1}
                    switch (-(fi13)) {
                    case -943:
                        delete(v4, "hello")
                    case 128:
                        if (!v3) {
                            v91 := GenS1{F0: 255}
                            _ = v91
                        } else {
                            if (((v6 * v6) != -2147483648) && v2) {
                                switch (fi88 + (fi88 & 127)) {
                                case 287:
                                    switch ("\n\t" + "\n\t") {
                                    case "foo bar":
                                        var v92 int16 = (^(v89))
                                        _ = v92
                                    default:
                                        v9 = append(v9, fi88)
                                    }
                                }
                            }
                            var v93 uint = v8
                            v6 += v6
                            _ = v93
                        }
                    default:
                        var v94 uint16 = 0
                        _ = v94
                    }
                    _ = fi88
                    _ = v89
                    _ = v90
                }
                v12 += 0.0
            }
            _ = fi13
        }
        if (!(v12 < ((v12 / v12) / v12))) {
            if (v1 || false) {
                v95 := v4[("" + "a")]
                _ = v95
            } else {
                switch (("\n\t" + ("a" + ("a" + "\n\t"))) + ("" + "a")) {
                case "\n\t":
                    var v96 uint64 = 42
                    _ = v96
                case "":
                    if v2 {
                        switch -2147483648 {
                        case 889:
                            v5 = 0
                        case -5:
                            v2 = v3
                        }
                        var v97 uint16 = 558
                        v12 -= v12
                        _ = v97
                    }
                default:
                    var v98 uint = (v8 ^ v8)
                    _ = v98
                }
                v99 := GenS1{F0: v8}
                for fi100 := 0; fi100 < 16; fi100++ {
                    if len(v9) > 0 {
                        v9[0] = (fi100 % (((fi100 << ((0) & 63))) | 1))
                    }
                    _ = fi100
                }
                _ = v99
            }
            _ = len(v9)
        }
    }
    v101 := GenS8{F0: 2}
    var v102 uint32 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v102:", v102)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v101
    _ = v102
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
