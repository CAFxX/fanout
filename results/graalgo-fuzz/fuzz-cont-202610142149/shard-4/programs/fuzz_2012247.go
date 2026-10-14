package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint32; F1 int32; F2 uint32 }
type GenS4 struct { F0 uint8; F1 uint; F2 uint8 }
type GenS5 struct { F0 uint16; F1 int64; F2 int64 }
type GenS6 struct { F0 uint32; F1 int64 }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 uint16 }

func main() {
    var v1 int32 = -296
    var v2 int64 = 9223372036854775807
    var v3 uint32 = 2147483647
    v4 := []string{"\n\t"}
    v5 := []uint32{v3, (v3 + (v3 >> ((4294967295) & 63))), v3, 2147483648}
    var v6 int = 1
    if false {
        var v7 int64 = v2
        _ = v7
    } else {
        if len(v4) > 0 {
            v4[0] = "a"
        }
        switch "" {
        case "hello":
            if (4294967295 == 422) {
                switch ("\n\t" + "") {
                case "a":
                    var v8 bool = (288 < 312)
                    _ = v8
                case "":
                    if true {
                        var v9 int32 = (-(v1))
                        v6 = v6
                        v10 := GenS1{F0: 18446744073709551615}
                        _ = v10
                        _ = v9
                    } else {
                        v6 += 9223372036854775807
                        v6 += 1
                    }
                case "hello":
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                default:
                    v6 += (v6 - v6)
                }
            }
        default:
            fmt.Println(v6, v3, v6)
        }
        fmt.Println(v1, v3, v2)
    }
    if true {
        var v11 float64 = 1e10
        v12 := GenS1{F0: 2}
        _ = v11
        _ = v12
    } else {
        v13 := GenS2{F0: 0}
        _ = v13
    }
    v5 = append(v5, 1)
    fmt.Println(v6, v3, v6)
    v14 := GenS2{F0: 1}
    var v15 float64 = (1.5 - 3.14159)
    if len(v5) > 0 {
        v5[0] = v3
    }
    if false {
        if (838 >= 416) {
            v15 = v15
            if true {
                switch "hello" {
                case "hello":
                    if true {
                        for fi16 := 0; fi16 < 6; fi16++ {
                            var v17 uint64 = 691
                            if false {
                                switch (fi16 << ((316) & 63)) {
                                case 9223372036854775807:
                                    v18 := GenS1{F0: 42}
                                    _ = v18
                                case -2147483648:
                                    for fi19 := 0; fi19 < 20; fi19++ {
                                        switch fi19 {
                                        case 32767:
                                            v20 := GenS3{F0: (v3 - 665), F1: v1, F2: 42}
                                            _ = v20
                                        case -1:
                                            var v21 uint = 255
                                            _ = v21
                                        case 2147483647:
                                            if ((3.14159 <= 1e10) && ((v3 << ((65535) & 63)) >= 1)) {
                                                if ((v3 % (((v3 * (v3 << ((255) & 63)))) | 1)) == (^(v3))) {
                                                    v22 := GenS2{F0: 18446744073709551615}
                                                    _ = v22
                                                }
                                                if len(v4) > 0 {
                                                    v4[0] = "hello"
                                                }
                                            } else {
                                                if (0 <= v2) {
                                                    v4 = append(v4, "")
                                                    v2 = int64(fi19)
                                                    v23 := GenS3{F0: (^(v3)), F1: v1, F2: (^(v3))}
                                                    v15 = ((v15 * v15) / (3.14159 - v15))
                                                    var v24 int = (^(fi19))
                                                    _ = v23
                                                    _ = v24
                                                } else {
                                                    if (true && true) {
                                                        v15 = (-2.25 + (v15 / ((v15 * 0.1) * v15)))
                                                        v1 = (v1 & (v1 % ((v1) | 1)))
                                                    } else {
                                                        switch (v6 - -727) {
                                                        case 127:
                                                            v2 -= -168
                                                        }
                                                        if false {
                                                            switch "foo bar" {
                                                            case "\n\t":
                                                                if (565 < 358) {
                                                                    v1 += -2147483648
                                                                    var v25 uint64 = (v17 & v17)
                                                                    v6 = (fi19 << ((0) & 63))
                                                                    _ = v25
                                                                } else {
                                                                    v1 = 1
                                                                    v26 := GenS3{F0: (v3 ^ 2), F1: (v1 + int32(fi19)), F2: (^(v3))}
                                                                    _ = v26
                                                                }
                                                            case "hello":
                                                                if (v1 > -1) {
                                                                    switch v6 {
                                                                    case 9223372036854775807:
                                                                        var v27 int = (v6 % (((fi19 ^ 127)) | 1))
                                                                        _ = v27
                                                                    default:
                                                                        var v28 int8 = 0
                                                                        _ = v28
                                                                    }
                                                                }
                                                            }
                                                            if true {
                                                                if ((^(v3)) < v3) {
                                                                    var v29 int32 = v1
                                                                    if (!(-58 >= 68)) {
                                                                        switch ("\n\t" + "\n\t") {
                                                                        case "":
                                                                            switch ("foo bar" + "\n\t") {
                                                                            case "hello":
                                                                                var v30 int8 = 42
                                                                                _ = v30
                                                                            case "\n\t":
                                                                                v6 -= (fi16 + (^(v6)))
                                                                            default:
                                                                                if (!((false && (55 != 59)) && false)) {
                                                                                    var v31 int8 = 127
                                                                                    _ = len(v5)
                                                                                    v3 = (^(v3))
                                                                                    _ = v31
                                                                                } else {
                                                                                    fmt.Println(v3, v15)
                                                                                    v2 += v2
                                                                                }
                                                                            }
                                                                            fmt.Println(v15, v2)
                                                                            v32 := GenS1{F0: 255}
                                                                            _ = v32
                                                                        case "hello":
                                                                            switch v6 {
                                                                            case 32767:
                                                                                v33 := GenS1{F0: 4294967295}
                                                                                _ = v33
                                                                            case -2147483648:
                                                                                v34 := GenS3{F0: 411, F1: 2147483647, F2: 421}
                                                                                _ = v34
                                                                            default:
                                                                                fmt.Println(v6)
                                                                            }
                                                                            v2 -= 42
                                                                            switch "a" {
                                                                            case "":
                                                                                var v35 uint16 = 42
                                                                                _ = v35
                                                                            case "foo bar":
                                                                                var v36 uint32 = v3
                                                                                _ = v36
                                                                            case "\n\t":
                                                                                _ = len(v4)
                                                                            }
                                                                        case "\n\t":
                                                                            v37 := GenS1{F0: 636}
                                                                            _ = v37
                                                                        }
                                                                        v17 -= v17
                                                                        var v38 uint8 = 127
                                                                        v17 = 1
                                                                        var v39 uint16 = 1
                                                                        var v40 int64 = v2
                                                                        fmt.Println(v38)
                                                                        _ = v38
                                                                        _ = v39
                                                                        _ = v40
                                                                    }
                                                                    v4 = append(v4, "\n\t")
                                                                    _ = v29
                                                                } else {
                                                                    v15 = v15
                                                                }
                                                                v3 -= 65535
                                                                switch "foo bar" {
                                                                case "\n\t":
                                                                    switch v6 {
                                                                    case -9223372036854775808:
                                                                        fmt.Println(v6, v17, v1)
                                                                    case 42:
                                                                        var v41 int64 = (v2 ^ v2)
                                                                        _ = v41
                                                                    default:
                                                                        v17 -= 18446744073709551615
                                                                        v42 := GenS3{F0: (v3 + (v3 * (v3 * 2147483647))), F1: (-(v1)), F2: v3}
                                                                        _ = v42
                                                                    }
                                                                    v43 := GenS3{F0: v3, F1: (v1 << ((73) & 63)), F2: (v3 | (v3 / (((^(v3))) | 1)))}
                                                                    v1 = v1
                                                                    _ = v43
                                                                case "hello":
                                                                    if (1 >= 65535) {
                                                                        switch int(v3) {
                                                                        case 128:
                                                                            v4 = append(v4, ("\n\t" + "\n\t"))
                                                                        case 127:
                                                                            var v44 int64 = -1
                                                                            _ = v44
                                                                        case 1:
                                                                            var v45 uint = 255
                                                                            if len(v5) > 0 {
                                                                                v5[0] = v3
                                                                            }
                                                                            if true {
                                                                                var v46 uint16 = 42
                                                                                v47 := GenS1{F0: (v45 | uint(v46))}
                                                                                _ = v46
                                                                                _ = v47
                                                                            }
                                                                            v48 := GenS2{F0: 18446744073709551615}
                                                                            _ = v45
                                                                            _ = v48
                                                                        }
                                                                        if true {
                                                                            var v49 bool = (false && true)
                                                                            _ = v49
                                                                        } else {
                                                                            if len(v4) >= 2 {
                                                                                v4 = v4[1:len(v4)]
                                                                            }
                                                                        }
                                                                        v1 += v1
                                                                    }
                                                                case "foo bar":
                                                                    var v50 uint64 = v17
                                                                    _ = v50
                                                                default:
                                                                    fmt.Println(v6, v6)
                                                                }
                                                            }
                                                            fmt.Println(v3, v15, fi19)
                                                        } else {
                                                            var v51 float32 = 0.0
                                                            v52 := GenS3{F0: v3, F1: v1, F2: v3}
                                                            _ = v51
                                                            _ = v52
                                                        }
                                                    }
                                                    if len(v5) > 0 {
                                                        v5[0] = 1
                                                    }
                                                }
                                                v1 = 1
                                            }
                                        }
                                        v3 = 445
                                        switch 2147483647 {
                                        case -2147483648:
                                            v2 = (v2 ^ 1)
                                        case 1:
                                            var v53 uint = 2
                                            _ = v53
                                        case 42:
                                            v54 := GenS4{F0: 42, F1: 18446744073709551615, F2: 118}
                                            _ = v54
                                        }
                                        var v55 string = "a"
                                        _ = fi19
                                        _ = v55
                                    }
                                }
                            }
                            var v56 int16 = 1
                            _ = fi16
                            _ = v17
                            _ = v56
                        }
                    } else {
                        v6 = (v6 - (^(v6)))
                        for fi57 := 0; fi57 < 5; fi57++ {
                            var v58 float32 = (1.5 * (0.0 / 0.1))
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                            fmt.Println(fi57)
                            _ = fi57
                            _ = v58
                        }
                    }
                default:
                    switch (^(v6)) {
                    case 670:
                        if true {
                            var v59 uint64 = 681
                            _ = v59
                        }
                    default:
                        _ = len(v4)
                    }
                }
                fmt.Println(v2, v2)
            } else {
                var v60 uint = 65535
                _ = v60
            }
        } else {
            fmt.Println(v3, v2)
        }
    } else {
        if false {
            _ = len(v5)
            fmt.Println(v15, v1)
            if true {
                v61 := GenS5{F0: 124, F1: v2, F2: 42}
                var v62 float32 = -2.25
                if (-579 > v1) {
                    for fi63 := 0; fi63 < 7; fi63++ {
                        v3 = v3
                        _ = fi63
                    }
                }
                _ = v61
                _ = v62
            }
        } else {
            var v64 int64 = v2
            for fi65 := 0; fi65 < 20; fi65++ {
                for fi66 := 0; fi66 < 9; fi66++ {
                    v3 -= v3
                    switch (fi66 % ((int(v3)) | 1)) {
                    case 42:
                        fmt.Println(v2)
                    default:
                        var v67 uint16 = 897
                        _ = v67
                    }
                    v1 = (^(v1))
                    var v68 int64 = (v64 + v64)
                    _ = fi66
                    _ = v68
                }
                for fi69 := 0; fi69 < 12; fi69++ {
                    if false {
                        fmt.Println(v15, v64, v15)
                        if false {
                            var v70 uint = 255
                            v2 = -34
                            _ = v70
                        } else {
                            var v71 float32 = (0.0 * float32(v15))
                            _ = v71
                        }
                    } else {
                        if ((1e10 >= float32(1e10)) && (v1 >= -434)) {
                            fmt.Println(v15)
                        } else {
                            var v72 int8 = 127
                            _ = v72
                        }
                    }
                    var v73 uint8 = 2
                    var v74 float32 = (float32(v15) - 3.14159)
                    _ = fi69
                    _ = v73
                    _ = v74
                }
                if false {
                    for fi75 := 0; fi75 < 7; fi75++ {
                        switch (fi65 + (fi65 << ((806) & 63))) {
                        case -973:
                            var v76 int = (fi65 >> ((18446744073709551615) & 63))
                            _ = v76
                        default:
                            var v77 bool = true
                            _ = v77
                        }
                        var v78 int32 = v1
                        var v79 bool = true
                        _ = fi75
                        _ = v78
                        _ = v79
                    }
                    fmt.Println(v6, v2, v3)
                    if false {
                        v80 := GenS5{F0: 299, F1: (v64 >> ((1) & 63)), F2: -1}
                        if (-9223372036854775808 < 128) {
                            if ((v3 % ((332) | 1)) < v3) {
                                var v81 uint8 = uint8(v3)
                                if (v3 > v3) {
                                    var v82 int8 = -1
                                    _ = v82
                                } else {
                                    if false {
                                        var v83 string = "a"
                                        if true {
                                            v4 = append(v4, ((("hello" + "\n\t") + "foo bar") + v83))
                                            _ = len(v4)
                                            v84 := GenS6{F0: 4294967295, F1: (v64 % ((v2) | 1))}
                                            _ = v84
                                        }
                                        _ = v83
                                    } else {
                                        var v85 float64 = v15
                                        switch (v6 & 233) {
                                        case 2147483647:
                                            for fi86 := 0; fi86 < 20; fi86++ {
                                                switch (-(v6)) {
                                                case 32768:
                                                    if len(v4) >= 2 {
                                                        v4 = v4[1:len(v4)]
                                                    }
                                                case -2147483648:
                                                    v87 := GenS7{F0: int32(v3)}
                                                    _ = v87
                                                default:
                                                    v15 = v85
                                                }
                                                _ = fi86
                                            }
                                        }
                                        _ = v85
                                    }
                                    fmt.Println(v2)
                                    var v88 float32 = 1e10
                                    _ = v88
                                }
                                for fi89 := 0; fi89 < 8; fi89++ {
                                    if (false || (713 >= 1)) {
                                        v81 = v81
                                    }
                                    _ = fi89
                                }
                                _ = v81
                            } else {
                                _ = len(v4)
                                if (!(!false)) {
                                    fmt.Println(fi65, v1)
                                    var v90 uint8 = 1
                                    fmt.Println(v6, v1)
                                    _ = v90
                                }
                                fmt.Println(v1)
                            }
                        } else {
                            v5 = append(v5, v3)
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                            fmt.Println(v3, v1)
                        }
                        v6 += 128
                        _ = v80
                    }
                } else {
                    if false {
                        var v91 bool = (-1 > int8(v3))
                        if v91 {
                            if v91 {
                                if (v91 || v91) {
                                    fmt.Println(v64)
                                } else {
                                    switch ("hello" + ("foo bar" + "hello")) {
                                    case "foo bar":
                                        v4 = append(v4, "\n\t")
                                    case "\n\t":
                                        if false {
                                            var v92 uint32 = (^(v3))
                                            if (!v91) {
                                                v4 = append(v4, "hello")
                                                for fi93 := 0; fi93 < 18; fi93++ {
                                                    v3 += 65535
                                                    if len(v4) >= 2 {
                                                        v4 = v4[1:len(v4)]
                                                    }
                                                    fmt.Println(fi65, v15)
                                                    v94 := GenS8{F0: 2}
                                                    switch 32767 {
                                                    case -2147483648:
                                                        var v95 float32 = 0.1
                                                        _ = v95
                                                    case -294:
                                                        var v96 int = (v6 - int(v64))
                                                        _ = v96
                                                    case -655:
                                                        var v97 string = ("foo bar" + ("" + "\n\t"))
                                                        _ = v97
                                                    default:
                                                        v1 = v1
                                                    }
                                                    _ = fi93
                                                    _ = v94
                                                }
                                            }
                                            _ = v92
                                        }
                                    }
                                }
                                v98 := GenS8{F0: 563}
                                v99 := GenS2{F0: 885}
                                _ = v98
                                _ = v99
                            }
                            var v100 bool = false
                            _ = v100
                        } else {
                            _ = len(v4)
                            var v101 int16 = 1
                            _ = v101
                        }
                        _ = v91
                    }
                }
                _ = fi65
            }
            if (!((("a" + "") <= ("hello" + "a")) && true)) {
                var v102 int16 = 42
                if len(v5) > 0 {
                    v5[0] = 4294967295
                }
                v3 -= v3
                _ = v102
            }
            _ = v64
        }
        var v103 bool = ((v6 << ((633) & 63)) != (-(v6)))
        _ = v103
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
