package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int8; F1 int64; F2 int64 }
type GenS4 struct { F0 uint8; F1 int16 }
type GenS5 struct { F0 uint; F1 int8 }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 int; F1 int8 }
type GenS8 struct { F0 uint16; F1 uint16; F2 int8 }
type GenS9 struct { F0 uint64; F1 int32 }
type GenS10 struct { F0 int32; F1 uint; F2 int16 }
type GenS11 struct { F0 int8; F1 int16 }
type GenS12 struct { F0 int64 }
type GenS13 struct { F0 int }
type GenS14 struct { F0 bool; F1 int64 }
type GenS15 struct { F0 string; F1 int32 }
type GenS16 struct { F0 uint32; F1 int }

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[int]string)
    var v3 int = 127
    var v4 int8 = 34
    v5 := make(map[int]int)
    v6 := []int{v3, v3, v3}
    fmt.Println(v3, v3)
    var v7 int8 = v4
    v3 = (v3 / ((v3) | 1))
    v7 = -128
    v3 += (-(v3))
    if (true && false) {
        var v8 uint32 = 4294967295
        v4 -= (v4 >> ((18446744073709551615) & 63))
        _ = v8
    } else {
        switch v3 {
        case 32767:
            var v9 bool = (255 <= 42)
            _ = v9
        default:
            v10 := GenS1{F0: 72, F1: v4}
            _ = v10
        }
        if (!(-1 == -1)) {
            fmt.Println(v4, v4, v3)
        } else {
            v4 = (v7 * v4)
            v11 := GenS1{F0: 1, F1: 127}
            for fi12 := 0; fi12 < 17; fi12++ {
                var v13 uint = 828
                _ = fi12
                _ = v13
            }
            _ = v11
        }
        if false {
            v7 = v4
            for fi14 := 0; fi14 < 9; fi14++ {
                v15 := GenS1{F0: 9223372036854775807, F1: -128}
                switch 128 {
                case -385:
                    v4 += 100
                case -469:
                    for fi16 := 0; fi16 < 9; fi16++ {
                        if (true && (false && ("" != ((("\n\t" + "a") + "\n\t") + "hello")))) {
                            if (v7 != (v7 - 100)) {
                                var v17 string = "\n\t"
                                if ((231 >= 128) || (42 < 1000)) {
                                    fmt.Println(fi14)
                                    v17 = v17
                                }
                                fmt.Println(fi14, v7, v7)
                                _ = v17
                            } else {
                                v7 += 114
                            }
                            if len(v6) > 0 {
                                v6[0] = (fi16 & -9223372036854775808)
                            }
                        } else {
                            switch "hello" {
                            case "hello":
                                switch (fi16 - -9223372036854775808) {
                                case 42:
                                    fmt.Println(fi14, v3)
                                case 0:
                                    var v18 uint = 42
                                    _ = v18
                                default:
                                    fmt.Println(v4, v3, fi16)
                                }
                            case "a":
                                fmt.Println(v3, v3, fi16)
                            default:
                                var v19 int8 = -28
                                var v20 bool = (!true)
                                var v21 float32 = 3.14159
                                _ = v19
                                _ = v20
                                _ = v21
                            }
                            if (!false) {
                                v22 := GenS2{F0: fi16}
                                var v23 int8 = v4
                                v7 = v7
                                _ = len(v6)
                                var v24 int8 = v4
                                v2[fi16] = ""
                                switch "\n\t" {
                                case "hello":
                                    _ = len(v6)
                                }
                                _ = v22
                                _ = v23
                                _ = v24
                            } else {
                                v6 = append(v6, -2147483648)
                                fmt.Println(v7, v3, v4)
                                var v25 uint8 = 143
                                _ = len(v6)
                                _ = v25
                            }
                        }
                        _ = fi16
                    }
                default:
                    if false {
                        if len(v6) > 0 {
                            v6[0] = (fi14 << ((0) & 63))
                        }
                    }
                }
                for fi26 := 0; fi26 < 4; fi26++ {
                    if ((^(fi26)) <= fi26) {
                        var v27 int64 = -1
                        _ = v27
                    }
                    _ = fi26
                }
                _ = fi14
                _ = v15
            }
        } else {
            for fi28 := 0; fi28 < 9; fi28++ {
                v29 := GenS2{F0: (^(fi28))}
                var v30 uint32 = 168
                var v31 string = "a"
                _ = fi28
                _ = v29
                _ = v30
                _ = v31
            }
            var v32 uint8 = 0
            _ = v32
        }
    }
    switch (("a" + ("a" + "foo bar")) + "") {
    case "a":
        var v33 int32 = 42
        _ = v33
    default:
        v34 := GenS3{F0: (v7 ^ v4), F1: 0, F2: 42}
        _ = v34
    }
    switch ("foo bar" + "foo bar") {
    case "hello":
        var v35 float32 = (-2.25 * 1.5)
        _ = v35
    case "a":
        v3 = v3
    }
    if true {
        v36 := GenS3{F0: v7, F1: 1, F2: 42}
        v7 += 42
        if len(v6) > 0 {
            v6[0] = (^(v3))
        }
        _ = v36
    } else {
        if (74 > 42) {
            if (!(true || (255 >= 1))) {
                v37 := GenS4{F0: 113, F1: 42}
                _ = v37
            } else {
                v7 -= (v4 + (-(v4)))
                v4 -= (v4 / ((v4) | 1))
                switch "hello" {
                case "hello":
                    _ = len(v6)
                case "foo bar":
                    for fi38 := 0; fi38 < 8; fi38++ {
                        if false {
                            var v39 int8 = v4
                            for fi40 := 0; fi40 < 18; fi40++ {
                                delete(v2, (-(v3)))
                                if (false || false) {
                                    var v41 int = (-(v3))
                                    var v42 float32 = 1e10
                                    if (!true) {
                                        if (!(176 != 18446744073709551615)) {
                                            if (85 <= 0) {
                                                var v43 int8 = -60
                                                _ = v43
                                            }
                                            var v44 int16 = 32767
                                            _ = v44
                                        }
                                        v45 := GenS5{F0: 1, F1: int8(v41)}
                                        _ = v45
                                    } else {
                                        v41 += fi40
                                        if (((1 <= 2) || true) || true) {
                                            var v46 string = "foo bar"
                                            if ((v46 <= v46) || false) {
                                                v1[("hello" + "")] = false
                                                fmt.Println(fi38, fi38)
                                                if (!false) {
                                                    if len(v6) > 0 {
                                                        v6[0] = v41
                                                    }
                                                } else {
                                                    var v47 uint32 = 42
                                                    fmt.Println(fi40, fi40)
                                                    _ = v47
                                                }
                                            } else {
                                                v42 = (v42 + (1e10 - v42))
                                                v48 := GenS6{F0: -32768}
                                                var v49 int16 = -32768
                                                v50 := v2[127]
                                                fmt.Println(fi38, v4)
                                                _ = v48
                                                _ = v49
                                                _ = v50
                                            }
                                            fmt.Println(v7)
                                            _ = v46
                                        }
                                        switch ("" + "") {
                                        case "a":
                                            v51 := GenS7{F0: v3, F1: v39}
                                            _ = v51
                                        case "\n\t":
                                            if (true && true) {
                                                v4 = v39
                                            }
                                        default:
                                            if (0 != -266) {
                                                v2[965] = "a"
                                                v7 = v39
                                                switch "" {
                                                case "":
                                                    v52 := GenS8{F0: 126, F1: 423, F2: (v39 & v4)}
                                                    _ = v52
                                                case "a":
                                                    switch "hello" {
                                                    case "\n\t":
                                                        v53 := GenS7{F0: fi40, F1: v39}
                                                        _ = v53
                                                    case "hello":
                                                        if true {
                                                            var v54 int64 = -842
                                                            _ = v54
                                                        }
                                                    }
                                                case "hello":
                                                    switch fi40 {
                                                    case 2147483647:
                                                        if (v41 == (^(fi40))) {
                                                            var v55 uint8 = 34
                                                            v56 := GenS9{F0: 42, F1: -2147483648}
                                                            _ = v55
                                                            _ = v56
                                                        }
                                                    case 128:
                                                        var v57 uint32 = 126
                                                        _ = v57
                                                    default:
                                                        fmt.Println(v41, fi40, v4)
                                                    }
                                                }
                                                if (-188 >= 1000) {
                                                    fmt.Println(v4, v39)
                                                    var v58 int64 = 524
                                                    if true {
                                                        var v59 int32 = -1
                                                        fmt.Println(v59, fi38, fi40)
                                                        var v60 uint32 = 42
                                                        if (((0.0 * (((float64(1.5) + (float64(float32((-2.25 - ((((1.5 / 1.0) + (0.0 * 0.1)) * float64((float32(3.14159) + 1e10))) * float64(v42))))) + 1.5)) * -2.25) - -2.25)) <= 1e10) || true) {
                                                            if false {
                                                                if ("foo bar" >= "hello") {
                                                                    fmt.Println(v7)
                                                                    var v61 string = "a"
                                                                    v62 := v5[(fi40 + -2147483648)]
                                                                    _ = v61
                                                                    _ = v62
                                                                } else {
                                                                    if len(v6) >= 2 {
                                                                        v6 = v6[1:len(v6)]
                                                                    }
                                                                    fmt.Println(v3, v41)
                                                                }
                                                                if (1 >= 42) {
                                                                    var v63 int64 = (v58 - (^(v58)))
                                                                    var v64 bool = ((v59 + (v59 >> ((500) & 63))) > (v59 ^ v59))
                                                                    _ = v63
                                                                    _ = v64
                                                                } else {
                                                                    v6 = append(v6, 233)
                                                                    v59 -= (v59 * 42)
                                                                }
                                                            } else {
                                                                v59 = 1
                                                            }
                                                            var v65 int64 = (v58 | v58)
                                                            switch (v41 % ((32768) | 1)) {
                                                            case 9223372036854775807:
                                                                v66 := GenS10{F0: (v59 << ((675) & 63)), F1: 65535, F2: 1000}
                                                                _ = v66
                                                            default:
                                                                if ("" >= ("hello" + ((("a" + "a") + ("\n\t" + "")) + ""))) {
                                                                    var v67 int64 = (-(v58))
                                                                    _ = v67
                                                                } else {
                                                                    v65 += int64(v59)
                                                                }
                                                            }
                                                            _ = v65
                                                        }
                                                        _ = v59
                                                        _ = v60
                                                    }
                                                    _ = v58
                                                } else {
                                                    if false {
                                                        v6 = append(v6, 1)
                                                        switch v41 {
                                                        case 321:
                                                            if true {
                                                                var v68 uint16 = 824
                                                                _ = v68
                                                            } else {
                                                                if (42 > 42) {
                                                                    v7 = (v7 + -128)
                                                                    var v69 int = (v41 - (-(v41)))
                                                                    var v70 bool = (-299 >= v41)
                                                                    var v71 uint = 781
                                                                    _ = v69
                                                                    _ = v70
                                                                    _ = v71
                                                                } else {
                                                                    fmt.Println(v7, v42)
                                                                }
                                                                if ((-2.25 - ((1.5 + (1e10 - 0.1)) * 0.0)) == ((0.1 + 0.1) + 3.14159)) {
                                                                    var v72 float32 = ((0.1 / (3.14159 / 1.0)) / v42)
                                                                    _ = v72
                                                                } else {
                                                                    if len(v6) > 0 {
                                                                        v6[0] = (v41 | (v3 - (-(fi38))))
                                                                    }
                                                                    var v73 uint32 = 2147483648
                                                                    _ = v73
                                                                }
                                                                var v74 int = 327
                                                                _ = v74
                                                            }
                                                        default:
                                                            v75 := GenS11{F0: v4, F1: 305}
                                                            _ = v75
                                                        }
                                                        fmt.Println(v39, v42, v7)
                                                    } else {
                                                        if true {
                                                            v76 := GenS10{F0: -2147483648, F1: 18446744073709551615, F2: -32768}
                                                            if false {
                                                                var v77 int8 = (^(v39))
                                                                _ = v77
                                                            }
                                                            _ = v76
                                                        }
                                                        var v78 uint = 65535
                                                        switch (fi38 * (v41 << (((v78 * v78)) & 63))) {
                                                        case 32768:
                                                            var v79 uint16 = 255
                                                            _ = v79
                                                        case 32767:
                                                            if len(v6) > 0 {
                                                                v6[0] = (v3 << ((v78) & 63))
                                                            }
                                                        }
                                                        _ = v78
                                                    }
                                                    v80 := GenS12{F0: -1}
                                                    _ = v80
                                                }
                                            }
                                        }
                                        _ = len(v6)
                                    }
                                    v5[v3] = -9223372036854775808
                                    if (true || (745 > 42)) {
                                        _ = len(v1)
                                        var v81 uint = 65535
                                        _ = v81
                                    }
                                    _ = v41
                                    _ = v42
                                }
                                fmt.Println(fi38)
                                var v82 bool = (false || (v4 > 0))
                                _ = fi40
                                _ = v82
                            }
                            if true {
                                for fi83 := 0; fi83 < 7; fi83++ {
                                    var v84 bool = ((1e10 / 1.0) != ((0.1 / -2.25) / ((float32(0.0) + (0.1 + (1e10 - 1e10))) * (((0.1 / 0.1) + ((float32(1.5) + 0.1) * -2.25)) + 3.14159))))
                                    var v85 float64 = (1.5 + (3.14159 - (-2.25 + 0.0)))
                                    _ = fi83
                                    _ = v84
                                    _ = v85
                                }
                                v3 += -9223372036854775808
                            }
                            _ = v39
                        }
                        switch 127 {
                        case -1:
                            if (!true) {
                                if (!(false || (-900 > 2147483647))) {
                                    v86 := GenS5{F0: 2, F1: (v7 % ((v7) | 1))}
                                    switch "a" {
                                    case "a":
                                        var v87 uint16 = 1
                                        _ = v87
                                    }
                                    var v88 string = "\n\t"
                                    _ = v86
                                    _ = v88
                                }
                            } else {
                                fmt.Println(fi38, v4)
                            }
                        case -2147483648:
                            v89 := GenS13{F0: (v3 * (fi38 / (((v3 * 929)) | 1)))}
                            _ = v89
                        case 127:
                            v90 := GenS14{F0: true, F1: 42}
                            _ = v90
                        default:
                            v91 := GenS15{F0: ("\n\t" + "\n\t"), F1: -320}
                            _ = v91
                        }
                        _ = fi38
                    }
                default:
                    var v92 int = (-(v3))
                    _ = v92
                }
            }
        } else {
            switch ("foo bar" + "") {
            case "foo bar":
                v7 += v7
            case "a":
                fmt.Println(v7)
            }
        }
        _ = len(v6)
        v3 = 32768
    }
    switch 127 {
    case 7:
        _ = len(v6)
    case 127:
        _ = len(v2)
    case 2147483647:
        fmt.Println(v7, v3, v7)
    default:
        v93 := GenS16{F0: 349, F1: v3}
        _ = v93
    }
    for fi94 := 0; fi94 < 10; fi94++ {
        fmt.Println(v7, v4)
        fmt.Println(v3)
        _ = len(v2)
        _ = fi94
    }
    if len(v6) > 0 {
        v6[0] = v3
    }
    v7 = v7
    for fi95 := 0; fi95 < 11; fi95++ {
        v3 = v3
        _ = fi95
    }
    v96 := v2[v3]
    _ = len(v6)
    switch (v96 + (v96 + "hello")) {
    case "\n\t":
        v96 = v96
    }
    v3 = (v3 & (v3 << ((75) & 63)))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v96:", v96)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v96
}
