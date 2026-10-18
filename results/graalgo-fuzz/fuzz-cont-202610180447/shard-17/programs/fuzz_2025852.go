package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64; F2 int }
type GenS2 struct { F0 uint32; F1 int32 }
type GenS3 struct { F0 uint8; F1 int16; F2 uint16 }
type GenS4 struct { F0 int64; F1 int8; F2 int }
type GenS5 struct { F0 bool; F1 int64 }
type GenS6 struct { F0 string; F1 int16; F2 int32 }
type GenS7 struct { F0 int; F1 int64; F2 uint8 }
type GenS8 struct { F0 uint32; F1 uint; F2 int16 }
type GenS9 struct { F0 int8 }
type GenS10 struct { F0 int64; F1 uint8 }

func main() {
    var v1 int16 = 986
    v2 := make(map[uint32]uint32)
    var v3 string = (("a" + "a") + (("foo bar" + ("\n\t" + (("hello" + "hello") + "a"))) + ""))
    v4 := make(map[int]int)
    v1 = 549
    switch v3 {
    case "\n\t":
        v3 = v3
    case "":
        if false {
            var v5 float32 = (((1e10 - (1.5 - 3.14159)) * 3.14159) * 1.5)
            _ = v5
        } else {
            if false {
                var v6 string = v3
                v6 = ((v3 + v6) + "hello")
                _ = v6
            } else {
                fmt.Println(v3)
                v1 = (v1 - 245)
            }
            if (2 > 42) {
                for fi7 := 0; fi7 < 7; fi7++ {
                    switch fi7 {
                    case 128:
                        for fi8 := 0; fi8 < 18; fi8++ {
                            var v9 uint = 758
                            _ = fi8
                            _ = v9
                        }
                    case 4294967295:
                        var v10 bool = (0 > 1)
                        _ = v10
                    default:
                        var v11 int = (fi7 << ((2) & 63))
                        _ = v11
                    }
                    v1 = v1
                    fmt.Println(fi7, v1, fi7)
                    _ = fi7
                }
            }
            v1 -= v1
        }
    case "foo bar":
        var v12 int32 = 42
        _ = v12
    }
    switch (((v3 + v3) + v3) + "a") {
    case "foo bar":
        delete(v4, -9223372036854775808)
    case "a":
        if true {
            for fi13 := 0; fi13 < 11; fi13++ {
                for fi14 := 0; fi14 < 4; fi14++ {
                    v15 := []int{fi14, fi13, fi13}
                    v3 = ((v3 + ("" + "")) + v3)
                    _ = fi14
                    _ = v15
                }
                _ = fi13
            }
            var v16 bool = false
            _ = v16
        } else {
            var v17 int64 = 0
            _ = v17
        }
    case "":
        v18 := GenS1{F0: 68, F1: 669, F2: 352}
        _ = v18
    }
    switch ("a" + (v3 + v3)) {
    case "\n\t":
        if (0 >= 1) {
            for fi19 := 0; fi19 < 9; fi19++ {
                v1 -= v1
                v1 -= -737
                for fi20 := 0; fi20 < 16; fi20++ {
                    switch v3 {
                    case "foo bar":
                        var v21 uint32 = 789
                        _ = v21
                    case "":
                        fmt.Println(v3)
                    case "a":
                        if false {
                            if true {
                                var v22 bool = true
                                v23 := GenS2{F0: 2147483648, F1: 42}
                                v24 := GenS1{F0: 100, F1: 42, F2: (-(fi20))}
                                switch "\n\t" {
                                case "a":
                                    v1 += (v1 + int16(fi19))
                                }
                                v22 = true
                                var v25 int8 = -1
                                var v26 uint16 = 250
                                if (0.1 == ((3.14159 / -2.25) / 1.5)) {
                                    if v22 {
                                        v27 := []int{}
                                        v22 = v22
                                        if (!(42 > 1)) {
                                            var v28 bool = v22
                                            var v29 int32 = 0
                                            v1 = v1
                                            _ = v28
                                            _ = v29
                                        } else {
                                            if v22 {
                                                v30 := v4[fi20]
                                                v31 := v4[(-(fi20))]
                                                _ = v30
                                                _ = v31
                                            } else {
                                                var v32 uint = 696
                                                var v33 float64 = (1.5 + 1.5)
                                                _ = v32
                                                _ = v33
                                            }
                                            v26 += 32767
                                        }
                                        _ = v27
                                    }
                                    if (!v22) {
                                        v34 := GenS2{F0: 4294967295, F1: 2147483647}
                                        switch v3 {
                                        case "foo bar":
                                            v3 = (v3 + "foo bar")
                                        case "":
                                            var v35 float64 = ((0.1 * 1.5) * 1.5)
                                            _ = v35
                                        }
                                        _ = v34
                                    } else {
                                        if (1 < 652) {
                                            v36 := GenS1{F0: 1, F1: 0, F2: 9223372036854775807}
                                            fmt.Println(v25, v1)
                                            v1 = (v1 ^ (v1 % (((v1 % ((32767) | 1))) | 1)))
                                            _ = v36
                                        }
                                        var v37 uint = 1
                                        _ = v37
                                    }
                                } else {
                                    switch fi20 {
                                    case 32767:
                                        switch -2147483648 {
                                        case 42:
                                            v38 := GenS3{F0: 255, F1: (v1 >> ((298) & 63)), F2: v26}
                                            _ = v38
                                        case 356:
                                            if false {
                                                delete(v2, 42)
                                                v39 := GenS2{F0: 65535, F1: 42}
                                                v40 := []int{fi20}
                                                _ = v39
                                                _ = v40
                                            } else {
                                                if (8 != (v26 - (v26 << ((1) & 63)))) {
                                                    if v22 {
                                                        switch fi19 {
                                                        case 143:
                                                            v41 := GenS1{F0: 127, F1: 18446744073709551615, F2: fi19}
                                                            _ = v41
                                                        case -9223372036854775808:
                                                            fmt.Println(v22, v1, v26)
                                                        case -2147483648:
                                                            var v42 int32 = 42
                                                            _ = v42
                                                        default:
                                                            v22 = false
                                                            v26 = v26
                                                        }
                                                        v26 = v26
                                                    } else {
                                                        v22 = v22
                                                    }
                                                }
                                                var v43 int32 = int32(fi20)
                                                _ = v43
                                            }
                                        }
                                    default:
                                        var v44 int16 = v1
                                        _ = v44
                                    }
                                    fmt.Println(v3)
                                }
                                v1 += v1
                                _ = v22
                                _ = v23
                                _ = v24
                                _ = v25
                                _ = v26
                            } else {
                                if false {
                                    v1 = (v1 + v1)
                                    v45 := GenS4{F0: -1, F1: 100, F2: (fi20 + (fi20 % (((fi20 & 445)) | 1)))}
                                    v3 = "a"
                                    _ = v45
                                } else {
                                    v1 = v1
                                }
                                v46 := []int{(fi20 ^ (fi19 / (((fi20 & (fi19 ^ fi20))) | 1))), 865}
                                _ = v46
                            }
                            v47 := GenS2{F0: 841, F1: 42}
                            _ = v47
                        } else {
                            var v48 bool = true
                            if v48 {
                                v1 += v1
                            } else {
                                v49 := GenS5{F0: v48, F1: -363}
                                _ = v49
                            }
                            _ = v48
                        }
                    default:
                        var v50 uint8 = 127
                        var v51 float32 = ((float32((-2.25 * 1.5)) * 1.5) * (((1.5 * 0.0) + (-2.25 / 3.14159)) / (float32(3.14159) + -2.25)))
                        _ = v50
                        _ = v51
                    }
                    var v52 uint16 = 1
                    _ = fi20
                    _ = v52
                }
                _ = fi19
            }
            v1 = (v1 | 32767)
            v53 := []int{-9223372036854775808, -109}
            _ = v53
        } else {
            if true {
                var v54 string = ""
                var v55 uint8 = 91
                _ = v54
                _ = v55
            }
            if false {
                switch 9223372036854775807 {
                case -705:
                    v56 := GenS1{F0: -47, F1: 309, F2: 9223372036854775807}
                    _ = v56
                case 405:
                    fmt.Println(v3, v1)
                }
                for fi57 := 0; fi57 < 2; fi57++ {
                    for fi58 := 0; fi58 < 8; fi58++ {
                        var v59 string = v3
                        v3 = v3
                        _ = fi58
                        _ = v59
                    }
                    _ = fi57
                }
            }
        }
    case "hello":
        for fi60 := 0; fi60 < 2; fi60++ {
            if false {
                if false {
                    v1 = 1
                }
                if (0 < 1000) {
                    v61 := []int{(fi60 * (fi60 % ((-9223372036854775808) | 1))), (fi60 / (((fi60 << ((65535) & 63))) | 1)), fi60}
                    var v62 string = "a"
                    _ = v61
                    _ = v62
                }
                v3 = v3
            }
            v1 -= 720
            _ = fi60
        }
    case "a":
        var v63 uint64 = 42
        _ = v63
    default:
        v3 = v3
    }
    var v64 uint32 = 0
    for fi65 := 0; fi65 < 1; fi65++ {
        switch int(v64) {
        case -907:
            v66 := []int{}
            _ = v66
        case 2147483647:
            v64 = (v64 >> ((1) & 63))
        case 42:
            if (true || true) {
                v3 = (v3 + "hello")
                for fi67 := 0; fi67 < 7; fi67++ {
                    v3 = v3
                    v1 -= (v1 / ((32767) | 1))
                    _ = fi67
                }
            }
        }
        v1 -= v1
        v3 = v3
        _ = fi65
    }
    v68 := []int{1}
    var v69 int = -1
    var v70 int32 = -2147483648
    v71 := GenS3{F0: 245, F1: (v1 / (((v1 >> ((854) & 63))) | 1)), F2: 2}
    switch (v69 * int(v64)) {
    case -524:
        switch ("" + "a") {
        case "foo bar":
            v70 = v70
        case "a":
            if true {
                v72 := GenS6{F0: v3, F1: v1, F2: v70}
                _ = v72
            }
        case "":
            delete(v2, 1)
        default:
            v1 = (v1 * v1)
        }
    case 9223372036854775807:
        fmt.Println(v69, v69, v69)
    case 42:
        for fi73 := 0; fi73 < 6; fi73++ {
            for fi74 := 0; fi74 < 2; fi74++ {
                if false {
                    switch -9223372036854775808 {
                    case -1:
                        v75 := GenS7{F0: 148, F1: 1, F2: 255}
                        _ = v75
                    case 2147483647:
                        if true {
                            v76 := GenS8{F0: (v64 - (v64 % (((v64 >> ((0) & 63))) | 1))), F1: 18446744073709551615, F2: (v1 + -32768)}
                            if true {
                                var v77 string = ("a" + v3)
                                switch fi73 {
                                case 42:
                                    if (false || true) {
                                        var v78 uint32 = v64
                                        v79 := v4[fi73]
                                        var v80 float32 = 1.5
                                        _ = v78
                                        _ = v79
                                        _ = v80
                                    }
                                case 0:
                                    if false {
                                        switch "" {
                                        case "hello":
                                            switch v77 {
                                            case "":
                                                if true {
                                                    v77 = v3
                                                }
                                            case "hello":
                                                v81 := GenS4{F0: 265, F1: 1, F2: (^(v69))}
                                                _ = v81
                                            }
                                        case "a":
                                            if len(v68) >= 2 {
                                                v68 = v68[1:len(v68)]
                                            }
                                        }
                                        v64 = (v64 - (v64 * (v64 << ((18446744073709551615) & 63))))
                                        fmt.Println(v1, v69)
                                    } else {
                                        if len(v68) > 0 {
                                            v68[0] = (v69 * (v69 + (^(v69))))
                                        }
                                        v82 := GenS3{F0: 127, F1: (v1 + (v1 % ((0) | 1))), F2: 408}
                                        _ = v82
                                    }
                                case 4294967295:
                                    v83 := GenS9{F0: -1}
                                    _ = v83
                                default:
                                    if (false && (v77 <= "")) {
                                        if false {
                                            v77 = ("a" + "a")
                                            var v84 int16 = (v1 >> ((42) & 63))
                                            v69 += v69
                                            _ = v84
                                        }
                                        switch v3 {
                                        case "a":
                                            fmt.Println(v69, fi74, v64)
                                        case "\n\t":
                                            var v85 int32 = v70
                                            _ = v85
                                        case "foo bar":
                                            var v86 uint = 69
                                            _ = v86
                                        }
                                    }
                                }
                                var v87 bool = true
                                _ = v77
                                _ = v87
                            }
                            if (0.1 >= (((1.5 - (0.1 + float32(-2.25))) - (0.0 / -2.25)) / (-2.25 * (3.14159 / (1e10 + 0.0))))) {
                                _ = len(v68)
                                v69 = fi74
                            }
                            _ = v76
                        } else {
                            switch "hello" {
                            case "a":
                                v68 = append(v68, (fi74 / (((v69 * -9223372036854775808)) | 1)))
                            case "\n\t":
                                switch (v3 + "foo bar") {
                                case "hello":
                                    v88 := GenS2{F0: 578, F1: (v70 / (((v70 % ((1) | 1))) | 1))}
                                    _ = v88
                                }
                            case "":
                                var v89 uint32 = 2
                                _ = v89
                            default:
                                v1 = v1
                            }
                        }
                    }
                    var v90 uint32 = (v64 - 1)
                    _ = v90
                } else {
                    var v91 uint = 0
                    v1 = (v1 / (((v1 >> ((v91) & 63))) | 1))
                    var v92 uint16 = 0
                    _ = v91
                    _ = v92
                }
                v1 = v1
                switch fi74 {
                case -9223372036854775808:
                    v4[fi74] = 715
                case 1:
                    var v93 int16 = v1
                    var v94 int = (fi73 | (^(fi74)))
                    v93 = v93
                    _ = v93
                    _ = v94
                }
                if len(v68) > 0 {
                    v68[0] = fi74
                }
                if (61 <= 42) {
                    switch (fi74 % (((^(v69))) | 1)) {
                    case 418:
                        v64 -= v64
                    case 9223372036854775807:
                        _ = len(v2)
                    default:
                        v1 -= v1
                    }
                    if (506 != 2) {
                        v64 = v64
                        var v95 uint8 = 69
                        v70 = (v70 ^ (v70 % (((v70 >> ((956) & 63))) | 1)))
                        _ = v95
                    } else {
                        fmt.Println(fi74)
                        delete(v2, (^(v64)))
                    }
                    var v96 uint32 = 2147483647
                    v97 := GenS10{F0: -9223372036854775808, F1: 128}
                    _ = v96
                    _ = v97
                }
                _ = fi74
            }
            v69 = (v69 - 1)
            _ = fi73
        }
    default:
        v3 = "\n\t"
    }
    var v98 uint64 = 42
    v70 = v70
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v64:", v64)
    fmt.Println("slicelen v68:", len(v68), v68)
    fmt.Println("v69:", v69)
    fmt.Println("v70:", v70)
    fmt.Println("v98:", v98)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v64
    _ = v68
    _ = v69
    _ = v70
    _ = v71
    _ = v98
}
