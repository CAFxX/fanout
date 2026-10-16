package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32; F2 int32 }
type GenS2 struct { F0 uint64; F1 int32 }
type GenS3 struct { F0 string; F1 uint }
type GenS4 struct { F0 uint32; F1 int }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 int32; F1 bool }
type GenS7 struct { F0 int8; F1 int16 }
type GenS8 struct { F0 uint; F1 uint32; F2 int64 }
func genFn0() int16 {
    var v6 float32 = 0.1
    v6 -= 1.5
    var v7 float64 = 1.5
    v8 := GenS1{F0: 42, F1: 123, F2: -785}
    if (127 >= 0) {
        v9 := []int{42}
        _ = v9
    }
    _ = v6
    _ = v7
    _ = v8
    return 1000
}
func genFn1() int16 {
    v10 := []int{0, 32768}
    v11 := GenS1{F0: 1, F1: -1, F2: -2147483648}
    v10 = append(v10, 128)
    var v12 bool = (-9223372036854775808 != -1)
    _ = v10
    _ = v11
    _ = v12
    return -1
}

func main() {
    var v1 float64 = -2.25
    var v2 int = 275
    var v3 float64 = v1
    var v4 uint = 467
    var v5 bool = true
    for fi13 := 0; fi13 < 4; fi13++ {
        var v14 uint = v4
        if (false && false) {
            fmt.Println(fi13, fi13, v2)
        } else {
            for fi15 := 0; fi15 < 16; fi15++ {
                v5 = true
                fmt.Println(v5)
                _ = fi15
            }
            if ((v4 << ((54) & 63)) == 4294967295) {
                for fi16 := 0; fi16 < 3; fi16++ {
                    var v17 int8 = -31
                    var v18 int16 = genFn1()
                    _ = fi16
                    _ = v17
                    _ = v18
                }
                for fi19 := 0; fi19 < 7; fi19++ {
                    fmt.Println(v14, v5)
                    switch (fi13 & (v2 % ((32768) | 1))) {
                    case -662:
                        v20 := []int{1, (fi13 >> ((0) & 63))}
                        _ = v20
                    default:
                        v21 := []int{(v2 << ((793) & 63)), fi19}
                        _ = v21
                    }
                    v3 -= 0.0
                    _ = fi19
                }
                v3 += (((1e10 + (float64(3.14159) - (v3 - (-2.25 * 3.14159)))) - 1.5) - (v1 - float64((0.1 + 1e10))))
            }
            switch v2 {
            case -669:
                switch v2 {
                case -9223372036854775808:
                    if (0 >= 4294967295) {
                        for fi22 := 0; fi22 < 20; fi22++ {
                            switch (v2 | 9223372036854775807) {
                            case 128:
                                v23 := GenS1{F0: v4, F1: 2147483647, F2: 2147483647}
                                _ = v23
                            }
                            fmt.Println(v3)
                            _ = fi22
                        }
                        for fi24 := 0; fi24 < 3; fi24++ {
                            v2 -= fi24
                            _ = fi24
                        }
                        fmt.Println(v4, v5)
                    }
                case 44:
                    switch "a" {
                    case "foo bar":
                        var v25 int8 = 1
                        _ = v25
                    }
                default:
                    for fi26 := 0; fi26 < 15; fi26++ {
                        var v27 int8 = -128
                        var v28 uint16 = 0
                        switch -819 {
                        case 2147483647:
                            switch v2 {
                            case 32768:
                                v29 := []int{(fi13 >> (((v4 * (v4 >> (((v4 & (v4 ^ (v4 | 42)))) & 63)))) & 63)), (v2 - fi13), (fi13 ^ 2147483647), fi26}
                                _ = v29
                            case -2147483648:
                                v2 += (fi26 << ((v4) & 63))
                            default:
                                var v30 int64 = -23
                                _ = v30
                            }
                        case 9223372036854775807:
                            var v31 uint32 = 65535
                            _ = v31
                        }
                        _ = fi26
                        _ = v27
                        _ = v28
                    }
                }
            case 0:
                switch 42 {
                case 42:
                    for fi32 := 0; fi32 < 14; fi32++ {
                        if (v5 && v5) {
                            v4 -= (v4 | (v4 ^ 42))
                        } else {
                            v33 := []int{}
                            v34 := GenS1{F0: v14, F1: 2147483647, F2: 1}
                            var v35 int16 = 32767
                            _ = v33
                            _ = v34
                            _ = v35
                        }
                        if true {
                            v36 := GenS1{F0: (v14 ^ v14), F1: 2147483647, F2: -743}
                            v3 = ((float64(0.1) / v1) - (v1 / v3))
                            var v37 uint = 65535
                            _ = v36
                            _ = v37
                        } else {
                            v38 := []int{2147483647}
                            _ = v38
                        }
                        v39 := GenS1{F0: (v4 + (v4 << ((v4) & 63))), F1: 1, F2: 2147483647}
                        _ = fi32
                        _ = v39
                    }
                case 0:
                    var v40 uint = v4
                    _ = v40
                case -1:
                    var v41 int64 = 116
                    _ = v41
                }
            case -9223372036854775808:
                v14 += v14
            default:
                for fi42 := 0; fi42 < 2; fi42++ {
                    v43 := []int{}
                    v44 := GenS2{F0: 0, F1: 42}
                    switch ("\n\t" + "\n\t") {
                    case "hello":
                        switch v2 {
                        case -2147483648:
                            v14 += 65535
                            v45 := GenS3{F0: "", F1: (v4 + (v4 & (v4 | 2)))}
                            _ = v45
                        default:
                            v4 = (^(v4))
                        }
                    case "a":
                        v3 += 1.5
                    case "":
                        v46 := GenS4{F0: 2147483647, F1: (^(v2))}
                        _ = v46
                    }
                    _ = fi42
                    _ = v43
                    _ = v44
                }
            }
        }
        var v47 uint64 = 4294967295
        _ = fi13
        _ = v14
        _ = v47
    }
    switch "a" {
    case "a":
        var v48 uint64 = 4294967295
        _ = v48
    case "hello":
        v5 = true
    default:
        var v49 uint64 = 4294967295
        _ = v49
    }
    var v50 uint64 = 0
    fmt.Println(v1, v2)
    v51 := GenS5{F0: 779}
    switch 0 {
    case -795:
        fmt.Println(v5, v4, v2)
    case 9223372036854775807:
        var v52 float32 = 1.5
        _ = v52
    default:
        v53 := GenS6{F0: -421, F1: v5}
        _ = v53
    }
    fmt.Println(v50, v3, v3)
    var v54 bool = (v1 <= 3.14159)
    v4 = 1
    switch (-(v2)) {
    case 2147483647:
        for fi55 := 0; fi55 < 12; fi55++ {
            for fi56 := 0; fi56 < 16; fi56++ {
                var v57 int32 = -709
                _ = fi56
                _ = v57
            }
            v50 = (v50 + 684)
            _ = fi55
        }
    case -2147483648:
        v58 := []int{(v2 | 42)}
        _ = v58
    }
    for fi59 := 0; fi59 < 12; fi59++ {
        v60 := []int{(v2 << (((v4 >> (((v4 >> (((^(v4))) & 63))) & 63))) & 63)), (v2 / (((fi59 ^ (v2 + v2))) | 1)), (fi59 + (fi59 % (((-(v2))) | 1)))}
        switch v2 {
        case 32768:
            var v61 int = v2
            _ = v61
        case 128:
            v62 := GenS7{F0: -71, F1: -99}
            _ = v62
        case 1:
            fmt.Println(v50, fi59, v54)
        default:
            for fi63 := 0; fi63 < 15; fi63++ {
                if v5 {
                    v5 = (v54 && false)
                    if (v5 && (!(!(false && v54)))) {
                        switch v2 {
                        case -9223372036854775808:
                            switch (fi63 % (((fi63 ^ fi59)) | 1)) {
                            case 1:
                                v2 -= (fi63 * fi63)
                            }
                        }
                        var v64 uint = 0
                        if false {
                            v54 = ((93 == 2147483648) || true)
                            v1 += (((((v3 + -2.25) + 0.1) + 3.14159) + (3.14159 - 3.14159)) * float64((0.1 - (3.14159 * (1.5 + (float32(v3) * 0.1))))))
                        } else {
                            switch "hello" {
                            case "\n\t":
                                v54 = ((v50 >> ((229) & 63)) == v50)
                                v65 := GenS6{F0: -2147483648, F1: (v4 != 0)}
                                var v66 float32 = 3.14159
                                fmt.Println(fi63, fi59)
                                _ = v65
                                _ = v66
                            default:
                                if v5 {
                                    v4 += v64
                                }
                            }
                            var v67 int = fi59
                            v54 = v5
                            _ = v67
                        }
                        _ = v64
                    } else {
                        switch (fi63 << (((v4 ^ v4)) & 63)) {
                        case -2147483648:
                            var v68 string = "foo bar"
                            _ = v68
                        case -1:
                            v5 = v5
                        case 1:
                            var v69 string = "foo bar"
                            _ = v69
                        default:
                            v1 -= 1e10
                        }
                        if true {
                            if v54 {
                                switch 1 {
                                case 0:
                                    v54 = v54
                                case -2147483648:
                                    var v70 int8 = 33
                                    _ = v70
                                case 128:
                                    if len(v60) >= 2 {
                                        v60 = v60[1:len(v60)]
                                    }
                                default:
                                    switch ("foo bar" + "\n\t") {
                                    case "":
                                        v54 = false
                                    case "a":
                                        var v71 uint = v4
                                        _ = v71
                                    case "\n\t":
                                        if v5 {
                                            v72 := GenS2{F0: (v50 + (v50 ^ v50)), F1: 42}
                                            if false {
                                                v60 = append(v60, -1)
                                                switch "a" {
                                                case "\n\t":
                                                    v4 += v4
                                                case "a":
                                                    var v73 uint8 = 255
                                                    _ = v73
                                                default:
                                                    if false {
                                                        switch 127 {
                                                        case 0:
                                                            if ((!v54) || false) {
                                                                v74 := GenS8{F0: v4, F1: 65535, F2: 1}
                                                                fmt.Println(v2)
                                                                _ = v74
                                                            } else {
                                                                fmt.Println(v3, v50)
                                                                switch "foo bar" {
                                                                case "foo bar":
                                                                    switch ("" + "a") {
                                                                    case "hello":
                                                                        var v75 uint8 = 0
                                                                        _ = v75
                                                                    default:
                                                                        fmt.Println(v50, v3)
                                                                    }
                                                                case "a":
                                                                    var v76 int16 = 32767
                                                                    _ = v76
                                                                case "hello":
                                                                    v77 := GenS4{F0: 2147483648, F1: (v2 - fi59)}
                                                                    _ = v77
                                                                default:
                                                                    _ = len(v60)
                                                                }
                                                                var v78 int64 = 0
                                                                _ = v78
                                                            }
                                                        case 32768:
                                                            switch "" {
                                                            case "foo bar":
                                                                v79 := GenS7{F0: 42, F1: 32767}
                                                                _ = v79
                                                            case "hello":
                                                                var v80 uint32 = 65535
                                                                _ = v80
                                                            }
                                                        }
                                                    }
                                                }
                                            } else {
                                                v50 = 149
                                            }
                                            fmt.Println(v50)
                                            _ = v72
                                        }
                                    default:
                                        var v81 uint = (v4 & (v4 / (((v4 % ((65535) | 1))) | 1)))
                                        _ = v81
                                    }
                                }
                                var v82 uint8 = 72
                                v82 = 2
                                _ = v82
                            }
                        }
                        var v83 int32 = -287
                        _ = v83
                    }
                    if (1 == 534) {
                        v1 -= (v3 + float64(-2.25))
                        v50 = (v50 & v50)
                    }
                    fmt.Println(fi63)
                } else {
                    if v54 {
                        fmt.Println(v54)
                    }
                    v84 := GenS2{F0: v50, F1: 419}
                    _ = v84
                }
                _ = fi63
            }
        }
        _ = fi59
        _ = v60
    }
    for fi85 := 0; fi85 < 6; fi85++ {
        if (!false) {
            v3 -= 1.5
        } else {
            for fi86 := 0; fi86 < 2; fi86++ {
                if v5 {
                    var v87 int16 = 32767
                    var v88 uint = (v4 & v4)
                    var v89 uint64 = (v50 & v50)
                    _ = v87
                    _ = v88
                    _ = v89
                } else {
                    v50 = v50
                }
                var v90 int8 = 0
                v50 = v50
                _ = fi86
                _ = v90
            }
        }
        _ = fi85
    }
    if (3.14159 >= ((-2.25 * 0.0) - 1e10)) {
        var v91 float32 = 0.0
        switch "\n\t" {
        case "":
            v3 = v3
        default:
            var v92 uint = 0
            _ = v92
        }
        fmt.Println(v2, v3)
        _ = v91
    }
    for fi93 := 0; fi93 < 1; fi93++ {
        var v94 int64 = 1
        v95 := []int{v2}
        _ = fi93
        _ = v94
        _ = v95
    }
    fmt.Println(v54)
    v96 := GenS1{F0: 177, F1: -1, F2: 2147483647}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("v54:", v54)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v50
    _ = v51
    _ = v54
    _ = v96
}
