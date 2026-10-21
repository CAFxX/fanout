package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16; F2 int64 }
type GenS2 struct { F0 int16; F1 int; F2 int64 }
type GenS3 struct { F0 bool; F1 int32 }
type GenS4 struct { F0 int64; F1 uint64; F2 uint32 }
type GenS5 struct { F0 uint64; F1 uint64; F2 int32 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 int16 }
type GenS9 struct { F0 string }
type GenS10 struct { F0 int8; F1 string }
type GenS11 struct { F0 int32 }

func main() {
    var v1 uint32 = 0
    var v2 uint64 = 18446744073709551615
    var v3 bool = ((float32(-2.25) - ((1e10 / (float32((0.1 * (0.1 + 1.5))) / 0.1)) + 1.5)) > -2.25)
    var v4 uint16 = 1
    var v5 int16 = 32767
    var v6 string = "foo bar"
    switch 127 {
    case 32768:
        v7 := []int{4294967295, 632, 127}
        _ = v7
    case -351:
        v4 -= 593
    case -47:
        v5 = (v5 & 341)
    default:
        fmt.Println(v1)
    }
    switch -2147483648 {
    case 828:
        fmt.Println(v6)
    case -59:
        switch -9223372036854775808 {
        case 4294967295:
            v8 := GenS1{F0: 42, F1: (v4 * (v4 / ((32767) | 1))), F2: 724}
            _ = v8
        default:
            v9 := GenS2{F0: v5, F1: 0, F2: -261}
            _ = v9
        }
    default:
        v1 = (v1 | 65535)
    }
    v10 := []int{-1, 9223372036854775807, 32768, 829}
    if (0.1 == 1.5) {
        v11 := GenS1{F0: 255, F1: 266, F2: 1}
        var v12 float32 = 1.5
        if len(v10) > 0 {
            v10[0] = 1
        }
        _ = v11
        _ = v12
    } else {
        if false {
            switch -1 {
            case 0:
                v10 = append(v10, 458)
            case -1:
                switch ("foo bar" + ((("hello" + "a") + v6) + "foo bar")) {
                case "foo bar":
                    _ = len(v10)
                case "\n\t":
                    for fi13 := 0; fi13 < 18; fi13++ {
                        _ = len(v10)
                        v14 := GenS3{F0: (-128 <= 0), F1: 0}
                        if (!false) {
                            if (-9223372036854775808 > -1) {
                                if ((v2 ^ 777) == v2) {
                                    for fi15 := 0; fi15 < 13; fi15++ {
                                        var v16 int32 = -2147483648
                                        var v17 uint32 = 43
                                        if false {
                                            var v18 int64 = -211
                                            if (v3 && true) {
                                                var v19 uint8 = 127
                                                if (!((v19 / (((v19 * 196)) | 1)) > v19)) {
                                                    if (v3 && (!v3)) {
                                                        var v20 uint = 395
                                                        if (float64((1e10 + 3.14159)) == (3.14159 + (float64(3.14159) + (1e10 - (((0.0 + float64(((-2.25 - 0.0) + ((1.5 + 1.5) * 1.5)))) + (3.14159 + 1e10)) - (1.5 - (1.5 * -2.25))))))) {
                                                            if len(v10) >= 2 {
                                                                v10 = v10[1:len(v10)]
                                                            }
                                                            v21 := GenS4{F0: (v18 - 1), F1: 1, F2: (v17 >> (((v20 / ((uint(v4)) | 1))) & 63))}
                                                            _ = v21
                                                        }
                                                        switch "" {
                                                        case "foo bar":
                                                            if false {
                                                                v5 = v5
                                                            } else {
                                                                if true {
                                                                    if len(v10) > 0 {
                                                                        v10[0] = (fi13 >> (((v20 + v20)) & 63))
                                                                    }
                                                                } else {
                                                                    var v22 uint64 = (^(v2))
                                                                    _ = v22
                                                                }
                                                                var v23 string = ("a" + (("hello" + (v6 + v6)) + ("\n\t" + ("a" + "foo bar"))))
                                                                if ("\n\t" > v6) {
                                                                    v2 = v2
                                                                } else {
                                                                    if len(v10) >= 2 {
                                                                        v10 = v10[1:len(v10)]
                                                                    }
                                                                    switch "a" {
                                                                    case "\n\t":
                                                                        v3 = false
                                                                    default:
                                                                        v2 -= (v2 - uint64(v20))
                                                                    }
                                                                    var v24 int64 = v18
                                                                    _ = v24
                                                                }
                                                                _ = v23
                                                            }
                                                        case "a":
                                                            var v25 float32 = float32(-2.25)
                                                            _ = v25
                                                        case "hello":
                                                            if (v4 > 0) {
                                                                var v26 uint = (^(v20))
                                                                var v27 float32 = (3.14159 / (((((float32((1.5 / float64(1.5))) * 0.0) / 3.14159) / 0.1) - 0.0) + 1.5))
                                                                _ = v26
                                                                _ = v27
                                                            }
                                                        }
                                                        _ = v20
                                                    } else {
                                                        v4 = (v4 | (v4 % (((v4 & v4)) | 1)))
                                                        var v28 string = (v6 + (v6 + v6))
                                                        _ = v28
                                                    }
                                                }
                                                switch v6 {
                                                case "foo bar":
                                                    v16 = v16
                                                case "hello":
                                                    if v3 {
                                                        if (((v3 && (false || true)) || (v16 != v16)) || v3) {
                                                            var v29 int8 = 1
                                                            _ = v29
                                                        } else {
                                                            v30 := GenS4{F0: (v18 ^ v18), F1: (v2 % (((v2 - (v2 ^ (v2 / ((v2) | 1))))) | 1)), F2: (v1 ^ (v1 | (v1 - v17)))}
                                                            if ((^(v4)) < v4) {
                                                                var v31 int64 = v18
                                                                v31 -= v18
                                                                v32 := GenS4{F0: v18, F1: v2, F2: v17}
                                                                _ = v31
                                                                _ = v32
                                                            }
                                                            _ = v30
                                                        }
                                                        switch v6 {
                                                        case "":
                                                            var v33 uint16 = v4
                                                            _ = v33
                                                        }
                                                    }
                                                }
                                                _ = v19
                                            }
                                            var v34 uint = 1
                                            _ = v18
                                            _ = v34
                                        }
                                        _ = fi15
                                        _ = v16
                                        _ = v17
                                    }
                                }
                                if ((!(fi13 < (-(fi13)))) || false) {
                                    if (v3 || (!(!v3))) {
                                        for fi35 := 0; fi35 < 5; fi35++ {
                                            fmt.Println(v6)
                                            var v36 int32 = 2147483647
                                            if (v2 <= v2) {
                                                if len(v10) >= 2 {
                                                    v10 = v10[1:len(v10)]
                                                }
                                                v10 = append(v10, (fi13 >> ((440) & 63)))
                                            }
                                            switch v6 {
                                            case "\n\t":
                                                var v37 float64 = (1.5 * 3.14159)
                                                _ = v37
                                            case "hello":
                                                if (v1 >= (v1 + (v1 * uint32(v4)))) {
                                                    var v38 int16 = (v5 % (((v5 ^ v5)) | 1))
                                                    _ = v38
                                                }
                                            case "foo bar":
                                                v2 -= (v2 << ((75) & 63))
                                            default:
                                                v5 = (^(v5))
                                            }
                                            var v39 uint64 = v2
                                            if false {
                                                if len(v10) > 0 {
                                                    v10[0] = 914
                                                }
                                                v2 = (v2 | (v39 - (v2 ^ 1)))
                                            }
                                            _ = fi35
                                            _ = v36
                                            _ = v39
                                        }
                                        for fi40 := 0; fi40 < 20; fi40++ {
                                            v2 = (v2 << ((42) & 63))
                                            switch (-(fi40)) {
                                            case 1:
                                                if false {
                                                    if (907 > uint16(v1)) {
                                                        var v41 uint64 = (v2 % ((v2) | 1))
                                                        var v42 int = fi40
                                                        var v43 int8 = 66
                                                        _ = v41
                                                        _ = v42
                                                        _ = v43
                                                    }
                                                    var v44 uint64 = v2
                                                    switch (fi13 & (fi40 | 128)) {
                                                    case -9223372036854775808:
                                                        v10 = append(v10, fi13)
                                                        _ = len(v10)
                                                        if false {
                                                            var v45 uint32 = v1
                                                            var v46 uint = 18446744073709551615
                                                            _ = v45
                                                            _ = v46
                                                        }
                                                    default:
                                                        v4 = v4
                                                    }
                                                    _ = v44
                                                }
                                                if len(v10) >= 2 {
                                                    v10 = v10[1:len(v10)]
                                                }
                                            }
                                            if v3 {
                                                v1 = v1
                                                var v47 int32 = 2147483647
                                                var v48 int64 = int64(v47)
                                                _ = v47
                                                _ = v48
                                            } else {
                                                v49 := GenS5{F0: v2, F1: v2, F2: -486}
                                                if (35 <= -19) {
                                                    var v50 int = (fi13 | fi40)
                                                    v3 = true
                                                    switch (v6 + "\n\t") {
                                                    case "hello":
                                                        v5 = v5
                                                    case "":
                                                        var v51 int64 = -861
                                                        _ = v51
                                                    default:
                                                        var v52 int64 = -800
                                                        _ = v52
                                                    }
                                                    _ = v50
                                                }
                                                var v53 string = v6
                                                fmt.Println(v53, v1)
                                                _ = v49
                                                _ = v53
                                            }
                                            _ = fi40
                                        }
                                    }
                                }
                                var v54 int64 = 9223372036854775807
                                _ = v54
                            }
                            var v55 int32 = -1
                            _ = v55
                        } else {
                            v1 = v1
                            v2 -= (v2 - v2)
                            if v3 {
                                if true {
                                    switch (fi13 / ((824) | 1)) {
                                    case 32768:
                                        v56 := GenS5{F0: uint64(v1), F1: v2, F2: 1}
                                        _ = v56
                                    default:
                                        v10 = append(v10, (fi13 << ((65535) & 63)))
                                    }
                                    if false {
                                        if len(v10) > 0 {
                                            v10[0] = (fi13 & (^(fi13)))
                                        }
                                    } else {
                                        var v57 int = 617
                                        _ = v57
                                    }
                                    fmt.Println(v5, v2, fi13)
                                } else {
                                    _ = len(v10)
                                    if v3 {
                                        if len(v10) > 0 {
                                            v10[0] = fi13
                                        }
                                    } else {
                                        v58 := GenS6{F0: (true || v3)}
                                        _ = v58
                                    }
                                    var v59 bool = v3
                                    _ = v59
                                }
                                v60 := GenS7{F0: (v2 * v2)}
                                _ = v60
                            } else {
                                for fi61 := 0; fi61 < 4; fi61++ {
                                    if (441 == v4) {
                                        v4 = (v4 >> ((0) & 63))
                                        fmt.Println(v2, v6)
                                    } else {
                                        var v62 bool = (127 == 32)
                                        fmt.Println(v4)
                                        v5 -= v5
                                        v1 += (v1 % ((uint32(v2)) | 1))
                                        _ = v62
                                    }
                                    var v63 string = v6
                                    if true {
                                        v2 -= (v2 >> ((0) & 63))
                                    } else {
                                        fmt.Println(v63)
                                        v64 := GenS3{F0: (4294967295 != 2), F1: -124}
                                        v6 = v63
                                        _ = v64
                                    }
                                    _ = fi61
                                    _ = v63
                                }
                                if len(v10) > 0 {
                                    v10[0] = 1
                                }
                            }
                        }
                        _ = fi13
                        _ = v14
                    }
                default:
                    v65 := GenS8{F0: (v5 << ((4294967295) & 63))}
                    _ = v65
                }
            case 32768:
                v5 -= (v5 >> ((42) & 63))
            }
            v66 := GenS9{F0: ((((v6 + "foo bar") + v6) + v6) + ((("\n\t" + ("\n\t" + (v6 + v6))) + v6) + v6))}
            v67 := GenS10{F0: 1, F1: ((("\n\t" + v6) + ("\n\t" + v6)) + "\n\t")}
            _ = v66
            _ = v67
        } else {
            if (v3 || v3) {
                for fi68 := 0; fi68 < 20; fi68++ {
                    if (v3 && (!false)) {
                        var v69 int16 = v5
                        _ = v69
                    } else {
                        switch ("\n\t" + ((v6 + "foo bar") + v6)) {
                        case "foo bar":
                            fmt.Println(fi68, v1)
                        case "hello":
                            var v70 bool = true
                            _ = v70
                        }
                        for fi71 := 0; fi71 < 9; fi71++ {
                            switch (fi68 * (fi71 >> ((819) & 63))) {
                            case 2147483647:
                                var v72 int16 = v5
                                _ = v72
                            case 128:
                                if v3 {
                                    switch fi68 {
                                    case 32767:
                                        fmt.Println(v1, v1)
                                    case 42:
                                        v73 := GenS4{F0: 42, F1: (^(v2)), F2: 181}
                                        _ = v73
                                    default:
                                        _ = len(v10)
                                    }
                                    v5 = 1000
                                } else {
                                    var v74 int8 = -79
                                    if (56 > 524) {
                                        fmt.Println(v2)
                                        var v75 bool = v3
                                        _ = v75
                                    }
                                    _ = v74
                                }
                            default:
                                var v76 float64 = 0.1
                                _ = v76
                            }
                            _ = fi71
                        }
                        v77 := GenS9{F0: v6}
                        _ = v77
                    }
                    var v78 int64 = 0
                    _ = fi68
                    _ = v78
                }
                v79 := GenS2{F0: -1, F1: 79, F2: -228}
                _ = v79
            }
        }
    }
    v2 = (^(v2))
    var v80 uint32 = (v1 / ((4294967295) | 1))
    if len(v10) >= 2 {
        v10 = v10[1:len(v10)]
    }
    fmt.Println(v4)
    v81 := GenS10{F0: 54, F1: v6}
    v5 += v5
    var v82 int = 4294967295
    for fi83 := 0; fi83 < 17; fi83++ {
        if (!true) {
            v6 = v6
            v84 := GenS7{F0: v2}
            _ = v84
        }
        _ = fi83
    }
    var v85 bool = v3
    v1 = 65535
    switch "" {
    case "":
        for fi86 := 0; fi86 < 7; fi86++ {
            v2 += v2
            fmt.Println(v82, v82)
            for fi87 := 0; fi87 < 9; fi87++ {
                var v88 float64 = 1e10
                _ = fi87
                _ = v88
            }
            _ = fi86
        }
    case "\n\t":
        var v89 string = v6
        _ = v89
    default:
        v10 = append(v10, (v82 ^ (v82 | v82)))
    }
    v6 = ""
    v90 := GenS11{F0: 42}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v80:", v80)
    fmt.Println("v82:", v82)
    fmt.Println("v85:", v85)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v80
    _ = v81
    _ = v82
    _ = v85
    _ = v90
}
