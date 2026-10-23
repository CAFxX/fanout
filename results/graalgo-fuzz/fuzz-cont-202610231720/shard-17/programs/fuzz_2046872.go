package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint8; F1 uint64 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int8; F1 int8; F2 int32 }
type GenS6 struct { F0 uint16; F1 uint64; F2 uint16 }
type GenS7 struct { F0 int; F1 uint; F2 uint16 }
type GenS8 struct { F0 int32 }
func genFn0() string {
    v4 := GenS1{F0: 148, F1: 2}
    var v5 int = 2147483647
    _ = v4
    _ = v5
    return "hello"
}
func genFn1() uint32 {
    var v6 uint = 0
    v6 -= 522
    if false {
        fmt.Println(v6)
        var v7 bool = true
        fmt.Println(v6, v7)
        _ = v7
    } else {
        v6 -= 4294967295
    }
    v8 := []int{}
    _ = v6
    _ = v8
    return 65535
}

func main() {
    v1 := []uint32{42, 719, 180}
    v2 := []uint32{2147483648, 1, 4294967295}
    var v3 float64 = 1e10
    for fi9 := 0; fi9 < 13; fi9++ {
        v10 := GenS2{F0: (fi9 / ((fi9) | 1))}
        for fi11 := 0; fi11 < 3; fi11++ {
            if false {
                v3 = (0.0 + (1.5 * 1.5))
                var v12 string = (("\n\t" + "") + "\n\t")
                _ = v12
            } else {
                v3 = (((v3 - (3.14159 / v3)) + ((v3 / 1e10) - (3.14159 + 0.1))) * ((float64(1.5) + v3) + (v3 * (0.0 + (((v3 * (0.1 - v3)) * (v3 + v3)) - float64(-2.25))))))
                switch "a" {
                case "hello":
                    var v13 int64 = -9223372036854775808
                    _ = v13
                case "a":
                    _ = len(v2)
                case "foo bar":
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                }
                var v14 uint16 = 2
                _ = v14
            }
            _ = fi11
        }
        v15 := GenS3{F0: 127, F1: 42}
        _ = fi9
        _ = v10
        _ = v15
    }
    for fi16 := 0; fi16 < 17; fi16++ {
        for fi17 := 0; fi17 < 9; fi17++ {
            if (4 >= -32768) {
                if (("foo bar" + ("foo bar" + "")) > "a") {
                    v18 := GenS4{F0: 0}
                    var v19 float32 = (3.14159 - 0.0)
                    _ = v18
                    _ = v19
                } else {
                    switch fi17 {
                    case 127:
                        var v20 int8 = 0
                        _ = v20
                    case -9223372036854775808:
                        v3 = v3
                    }
                    v2 = append(v2, 2147483647)
                    v3 += float64((float32(v3) * (1.5 + 0.1)))
                }
                var v21 uint = 255
                _ = v21
            } else {
                var v22 uint8 = 227
                if (32767 == 1000) {
                    if true {
                        var v23 uint8 = (v22 + v22)
                        if true {
                            v22 = (^(v23))
                            v22 += (^(v23))
                            var v24 int16 = 0
                            _ = v24
                        } else {
                            v25 := GenS4{F0: v22}
                            _ = v25
                        }
                        _ = v23
                    } else {
                        v3 += v3
                    }
                }
                v26 := GenS5{F0: 127, F1: 127, F2: -1}
                _ = v22
                _ = v26
            }
            if len(v2) > 0 {
                v2[0] = 2147483647
            }
            _ = fi17
        }
        var v27 uint16 = 2
        _ = fi16
        _ = v27
    }
    var v28 int8 = 106
    if len(v2) > 0 {
        v2[0] = 2
    }
    fmt.Println(v3, v3)
    v29 := GenS4{F0: 2}
    if false {
        var v30 bool = (42 >= 2)
        fmt.Println(v3, v30)
        _ = v30
    } else {
        fmt.Println(v3)
    }
    var v31 float32 = 1.5
    switch 128 {
    case 127:
        var v32 int = -71
        _ = v32
    case -9223372036854775808:
        var v33 float32 = v31
        _ = v33
    case -559:
        if (v3 <= ((v3 - 3.14159) + ((1.5 + v3) - float64(v31)))) {
            var v34 int32 = 600
            if true {
                switch "hello" {
                case "foo bar":
                    var v35 float64 = (3.14159 + (float64((v31 + (v31 / float32((float64(v31) / v3))))) / ((0.1 + (v3 / 1e10)) + 0.1)))
                    _ = v35
                case "\n\t":
                    var v36 string = ("hello" + (("a" + "hello") + ("" + "\n\t")))
                    _ = v36
                case "hello":
                    v37 := GenS3{F0: 127, F1: 131}
                    _ = v37
                default:
                    if (((v34 + v34) != v34) && true) {
                        v28 -= v28
                        if (1 != 65535) {
                            if true {
                                v38 := GenS1{F0: 255, F1: 2}
                                v3 += (v3 - ((0.1 / v3) / (float64((v31 + float32(float64(((v31 + (v31 * (1.5 + v31))) * v31))))) * (float64(v31) * v3))))
                                _ = v38
                            }
                            var v39 float64 = v3
                            _ = v39
                        }
                        switch 2147483647 {
                        case -886:
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        case 1:
                            v40 := GenS6{F0: 574, F1: 0, F2: 32767}
                            _ = v40
                        case -9223372036854775808:
                            v3 = v3
                        default:
                            if true {
                                var v41 int32 = v34
                                if (v28 == (v28 | 80)) {
                                    switch "foo bar" {
                                    case "\n\t":
                                        var v42 int32 = -506
                                        _ = v42
                                    case "a":
                                        if true {
                                            if len(v1) > 0 {
                                                v1[0] = genFn1()
                                            }
                                            if false {
                                                fmt.Println(v3, v3, v28)
                                            } else {
                                                var v43 uint32 = 0
                                                _ = v43
                                            }
                                        }
                                    default:
                                        var v44 uint16 = 65535
                                        _ = v44
                                    }
                                    v31 = v31
                                    if true {
                                        for fi45 := 0; fi45 < 3; fi45++ {
                                            fmt.Println(v41, v31, v34)
                                            var v46 float64 = -2.25
                                            _ = fi45
                                            _ = v46
                                        }
                                        switch "foo bar" {
                                        case "\n\t":
                                            if false {
                                                fmt.Println(v34)
                                                var v47 uint16 = 1
                                                for fi48 := 0; fi48 < 13; fi48++ {
                                                    var v49 float64 = v3
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                    if ((v28 >= (v28 + int8(fi48))) && (2 < 4294967295)) {
                                                        var v50 int64 = 9223372036854775807
                                                        if (255 == 128) {
                                                            if len(v1) >= 2 {
                                                                v1 = v1[1:len(v1)]
                                                            }
                                                            if true {
                                                                var v51 uint64 = 910
                                                                switch fi48 {
                                                                case -9223372036854775808:
                                                                    _ = len(v1)
                                                                }
                                                                _ = v51
                                                            } else {
                                                                if true {
                                                                    v49 = 1.5
                                                                    if len(v2) >= 2 {
                                                                        v2 = v2[1:len(v2)]
                                                                    }
                                                                } else {
                                                                    v41 += 0
                                                                    for fi52 := 0; fi52 < 1; fi52++ {
                                                                        v28 -= v28
                                                                        _ = fi52
                                                                    }
                                                                    if false {
                                                                        var v53 int64 = 42
                                                                        _ = v53
                                                                    }
                                                                }
                                                                v31 -= 0.0
                                                                fmt.Println(v49)
                                                            }
                                                            _ = len(v2)
                                                        } else {
                                                            switch -189 {
                                                            case 128:
                                                                for fi54 := 0; fi54 < 15; fi54++ {
                                                                    if (!false) {
                                                                        if (!(92 >= 21)) {
                                                                            _ = len(v2)
                                                                            var v55 uint8 = 127
                                                                            v34 = v41
                                                                            _ = v55
                                                                        } else {
                                                                            var v56 uint32 = 42
                                                                            var v57 uint32 = 341
                                                                            v31 -= 1e10
                                                                            _ = v56
                                                                            _ = v57
                                                                        }
                                                                        var v58 int8 = (v28 + v28)
                                                                        var v59 float64 = 0.1
                                                                        _ = v58
                                                                        _ = v59
                                                                    } else {
                                                                        fmt.Println(v41, fi54)
                                                                        v47 += (^(v47))
                                                                    }
                                                                    _ = fi54
                                                                }
                                                            case -2147483648:
                                                                var v60 uint8 = 1
                                                                _ = v60
                                                            case -1:
                                                                v61 := GenS5{F0: (-(v28)), F1: int8(v50), F2: (v34 - v34)}
                                                                _ = v61
                                                            }
                                                            if (255 <= v47) {
                                                                if true {
                                                                    for fi62 := 0; fi62 < 17; fi62++ {
                                                                        v63 := GenS7{F0: -9223372036854775808, F1: 65535, F2: 65535}
                                                                        if (32767 < -399) {
                                                                            v31 = ((v31 + v31) + 1e10)
                                                                        }
                                                                        switch "\n\t" {
                                                                        case "a":
                                                                            var v64 int16 = 1
                                                                            var v65 uint8 = 127
                                                                            _ = v64
                                                                            _ = v65
                                                                        default:
                                                                            v66 := GenS7{F0: fi62, F1: 18446744073709551615, F2: (^(v47))}
                                                                            _ = v66
                                                                        }
                                                                        _ = fi62
                                                                        _ = v63
                                                                    }
                                                                    if len(v2) > 0 {
                                                                        v2[0] = 563
                                                                    }
                                                                }
                                                                v28 -= (v28 - v28)
                                                            } else {
                                                                for fi67 := 0; fi67 < 9; fi67++ {
                                                                    fmt.Println(fi48, v49, v50)
                                                                    _ = fi67
                                                                }
                                                                if (42 >= 335) {
                                                                    if ((v34 - v41) != v34) {
                                                                        switch fi48 {
                                                                        case -1:
                                                                            fmt.Println(v41, fi48, fi48)
                                                                        }
                                                                    } else {
                                                                        v68 := GenS3{F0: 0, F1: 0}
                                                                        v69 := GenS4{F0: 1}
                                                                        for fi70 := 0; fi70 < 12; fi70++ {
                                                                            var v71 int = fi48
                                                                            var v72 string = "foo bar"
                                                                            _ = fi70
                                                                            _ = v71
                                                                            _ = v72
                                                                        }
                                                                        _ = v68
                                                                        _ = v69
                                                                    }
                                                                }
                                                            }
                                                        }
                                                        _ = v50
                                                    }
                                                    _ = fi48
                                                    _ = v49
                                                }
                                                _ = v47
                                            }
                                        default:
                                            if false {
                                                v73 := GenS7{F0: 0, F1: 311, F2: 42}
                                                if true {
                                                    for fi74 := 0; fi74 < 11; fi74++ {
                                                        _ = len(v2)
                                                        v41 += (-(v34))
                                                        if (1.5 <= v31) {
                                                            fmt.Println(v3, v28, v31)
                                                        }
                                                        _ = fi74
                                                    }
                                                }
                                                if false {
                                                    var v75 int = -9223372036854775808
                                                    if true {
                                                        fmt.Println(v31, v75, v28)
                                                        var v76 uint16 = 65535
                                                        switch v75 {
                                                        case -1:
                                                            var v77 int16 = 32767
                                                            _ = v77
                                                        case 9223372036854775807:
                                                            for fi78 := 0; fi78 < 12; fi78++ {
                                                                for fi79 := 0; fi79 < 10; fi79++ {
                                                                    if len(v2) > 0 {
                                                                        v2[0] = 832
                                                                    }
                                                                    switch (fi79 | (v75 >> ((241) & 63))) {
                                                                    case 410:
                                                                        var v80 int64 = 1
                                                                        _ = v80
                                                                    case 42:
                                                                        if false {
                                                                            var v81 uint32 = 1
                                                                            if false {
                                                                                v82 := GenS8{F0: v41}
                                                                                v31 -= 1.5
                                                                                fmt.Println(v41)
                                                                                _ = v82
                                                                            }
                                                                            _ = v81
                                                                        }
                                                                    }
                                                                    _ = fi79
                                                                }
                                                                _ = fi78
                                                            }
                                                        case 32768:
                                                            var v83 float64 = v3
                                                            _ = v83
                                                        default:
                                                            v84 := GenS1{F0: 2, F1: 127}
                                                            _ = v84
                                                        }
                                                        _ = v76
                                                    }
                                                    var v85 string = "hello"
                                                    _ = v75
                                                    _ = v85
                                                } else {
                                                    v31 += float32(v3)
                                                    if (v3 <= v3) {
                                                        v1 = append(v1, 0)
                                                        fmt.Println(v31)
                                                        v34 -= v41
                                                    } else {
                                                        switch 580 {
                                                        case -896:
                                                            for fi86 := 0; fi86 < 18; fi86++ {
                                                                v41 = v34
                                                                fmt.Println(v28)
                                                                for fi87 := 0; fi87 < 12; fi87++ {
                                                                    switch 4294967295 {
                                                                    case 42:
                                                                        if false {
                                                                            v1 = append(v1, 65535)
                                                                        }
                                                                    case -2147483648:
                                                                        switch ("" + "\n\t") {
                                                                        case "\n\t":
                                                                            switch "hello" {
                                                                            case "\n\t":
                                                                                var v88 float32 = (1.5 - v31)
                                                                                _ = v88
                                                                            case "a":
                                                                                var v89 string = "a"
                                                                                _ = v89
                                                                            }
                                                                        case "hello":
                                                                            if (!false) {
                                                                                if true {
                                                                                    if len(v2) > 0 {
                                                                                        v2[0] = genFn1()
                                                                                    }
                                                                                    v3 = v3
                                                                                }
                                                                            } else {
                                                                                v41 = v34
                                                                                if (18446744073709551615 != 42) {
                                                                                    fmt.Println(v31, fi86, fi86)
                                                                                    v90 := GenS1{F0: 255, F1: 127}
                                                                                    _ = v90
                                                                                }
                                                                                if len(v1) > 0 {
                                                                                    v1[0] = 2147483647
                                                                                }
                                                                            }
                                                                        case "foo bar":
                                                                            if (true || (((!true) && (42 == 55)) || false)) {
                                                                                v31 = -2.25
                                                                                var v91 int16 = 1000
                                                                                v34 = int32(fi86)
                                                                                _ = v91
                                                                            }
                                                                        }
                                                                    case 882:
                                                                        v41 += v34
                                                                    default:
                                                                        fmt.Println(fi86, v3, v31)
                                                                    }
                                                                    _ = fi87
                                                                }
                                                                _ = fi86
                                                            }
                                                        case -9223372036854775808:
                                                            var v92 uint32 = 60
                                                            _ = v92
                                                        }
                                                    }
                                                    var v93 string = ("hello" + "foo bar")
                                                    _ = v93
                                                }
                                                _ = v73
                                            }
                                        }
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                    }
                                }
                                v94 := GenS2{F0: -272}
                                _ = v41
                                _ = v94
                            }
                        }
                    }
                }
                var v95 bool = (128 == -9223372036854775808)
                var v96 int64 = 382
                _ = v95
                _ = v96
            } else {
                v3 = 0.1
                var v97 int16 = -32768
                _ = v97
            }
            var v98 int = 554
            _ = v34
            _ = v98
        }
    }
    for fi99 := 0; fi99 < 12; fi99++ {
        v28 = (v28 & (v28 << ((4294967295) & 63)))
        v100 := GenS3{F0: 1, F1: 934}
        fmt.Println(fi99)
        _ = fi99
        _ = v100
    }
    v101 := GenS5{F0: (v28 >> ((18446744073709551615) & 63)), F1: 42, F2: 42}
    if len(v1) > 0 {
        v1[0] = genFn1()
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    for fi102 := 0; fi102 < 1; fi102++ {
        fmt.Println(fi102, fi102, v3)
        switch 715 {
        case 291:
            var v103 int64 = 473
            _ = v103
        case -835:
            v28 -= (v28 / (((v28 * 100)) | 1))
        default:
            switch "" {
            case "":
                var v104 uint = 18446744073709551615
                _ = v104
            }
        }
        v28 = 1
        _ = fi102
    }
    if false {
        if (65535 < 1) {
            var v105 uint32 = 318
            _ = v105
        }
        var v106 int32 = -921
        _ = v106
    } else {
        v107 := GenS2{F0: -916}
        _ = v107
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    _ = v1
    _ = v101
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v31
}
