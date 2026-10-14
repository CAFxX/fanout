package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int; F2 int }
type GenS2 struct { F0 uint8; F1 uint16 }
type GenS3 struct { F0 bool; F1 uint8; F2 string }
type GenS4 struct { F0 string; F1 string; F2 uint32 }
func genFn0() int {
    if (true || (942 != 743)) {
        var v6 int32 = -1
        var v7 int = -229
        var v8 float32 = 0.1
        _ = v6
        _ = v7
        _ = v8
    } else {
        var v9 string = "a"
        _ = v9
    }
    var v10 bool = (-1 < 100)
    v10 = (("" + ((((("hello" + "\n\t") + "a") + ("foo bar" + "\n\t")) + (("a" + "hello") + ("foo bar" + "\n\t"))) + "")) != "foo bar")
    fmt.Println(v10)
    _ = v10
    return 42
}
func genFn1(p0 int64, p1 uint, p2 bool) int {
    var v11 float32 = float32(0.0)
    v12 := []int{128, -732, genFn0(), 32767}
    _ = len(v12)
    p1 += p1
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    return -296
}

func main() {
    var v1 string = ("a" + "")
    var v2 string = v1
    v3 := make(map[uint32]uint32)
    v4 := []int{}
    v5 := make(map[int]bool)
    if (true || true) {
        v1 = (v2 + v1)
    }
    switch "" {
    case "hello":
        v1 = v1
    case "foo bar":
        if false {
            if len(v4) > 0 {
                v4[0] = 42
            }
        }
    default:
        v1 = v1
    }
    if true {
        var v13 int32 = -642
        v1 = "hello"
        _ = v13
    } else {
        var v14 float64 = (0.0 / float64(0.1))
        switch 397 {
        case 127:
            v1 = ("\n\t" + (v2 + (v2 + v1)))
        case 9223372036854775807:
            switch genFn0() {
            case 4294967295:
                var v15 bool = (42 > 49)
                _ = v15
            case 0:
                v14 -= (v14 - v14)
            }
        case 1:
            v14 = -2.25
        }
        fmt.Println(v2, v2, v14)
        _ = v14
    }
    if false {
        for fi16 := 0; fi16 < 10; fi16++ {
            if false {
                v1 = v2
                v1 = ((v1 + (v2 + v2)) + v2)
                for fi17 := 0; fi17 < 18; fi17++ {
                    v18 := GenS1{F0: 255, F1: (fi16 >> ((330) & 63)), F2: 42}
                    fmt.Println(v2)
                    var v19 uint32 = 314
                    _ = fi17
                    _ = v18
                    _ = v19
                }
            }
            for fi20 := 0; fi20 < 2; fi20++ {
                fmt.Println(v2, v1)
                _ = fi20
            }
            var v21 int32 = 2147483647
            _ = fi16
            _ = v21
        }
        for fi22 := 0; fi22 < 16; fi22++ {
            if true {
                if (32767 >= -32768) {
                    for fi23 := 0; fi23 < 12; fi23++ {
                        var v24 int64 = 42
                        var v25 uint32 = 0
                        _ = fi23
                        _ = v24
                        _ = v25
                    }
                    v1 = v2
                } else {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                }
                v1 = (v1 + v2)
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            } else {
                var v26 int64 = 9223372036854775807
                _ = v26
            }
            _ = fi22
        }
    }
    for fi27 := 0; fi27 < 13; fi27++ {
        v1 = "hello"
        v2 = "foo bar"
        _ = fi27
    }
    if (true && ((-2.25 * 1.5) > (-2.25 - (-2.25 - (1.5 + ((0.1 / 1e10) + -2.25)))))) {
        var v28 int32 = -2147483648
        if (11 > 1) {
            v29 := GenS2{F0: 0, F1: 1}
            if (18446744073709551615 >= 984) {
                if false {
                    v30 := GenS2{F0: 2, F1: 32767}
                    var v31 float32 = 1e10
                    _ = v30
                    _ = v31
                }
                var v32 int64 = -617
                _ = v32
            } else {
                v33 := GenS1{F0: 489, F1: 2147483647, F2: 0}
                v34 := GenS3{F0: (v2 < (("\n\t" + (v2 + "")) + v1)), F1: 130, F2: "a"}
                if true {
                    v35 := GenS1{F0: 758, F1: 1, F2: 32768}
                    fmt.Println(v1)
                    _ = v35
                } else {
                    v36 := GenS1{F0: 255, F1: 2147483647, F2: 9223372036854775807}
                    switch 188 {
                    case 4294967295:
                        fmt.Println(v2, v2, v2)
                    case -9223372036854775808:
                        switch -898 {
                        case 935:
                            v1 = (v2 + (("" + v2) + (("a" + v1) + "hello")))
                        case -9223372036854775808:
                            for fi37 := 0; fi37 < 6; fi37++ {
                                if false {
                                    for fi38 := 0; fi38 < 9; fi38++ {
                                        v2 = v2
                                        if true {
                                            v3[406] = 2147483647
                                        }
                                        var v39 float32 = (((0.0 + 0.0) / 3.14159) / 1.5)
                                        _ = fi38
                                        _ = v39
                                    }
                                    var v40 int32 = (v28 * (v28 << ((18446744073709551615) & 63)))
                                    _ = v40
                                }
                                v41 := v3[201]
                                _ = fi37
                                _ = v41
                            }
                        default:
                            v3[700] = 2147483647
                        }
                    }
                    _ = len(v3)
                    _ = v36
                }
                _ = v33
                _ = v34
            }
            for fi42 := 0; fi42 < 16; fi42++ {
                for fi43 := 0; fi43 < 1; fi43++ {
                    v1 = v2
                    var v44 string = v2
                    var v45 string = v2
                    _ = fi43
                    _ = v44
                    _ = v45
                }
                for fi46 := 0; fi46 < 8; fi46++ {
                    v47 := GenS1{F0: 2, F1: fi46, F2: (fi46 / ((-663) | 1))}
                    if (265 < 65535) {
                        v28 -= -1
                        if true {
                            if false {
                                v48 := v5[961]
                                var v49 bool = v48
                                var v50 bool = true
                                _ = v48
                                _ = v49
                                _ = v50
                            } else {
                                v1 = ((("" + (v1 + v1)) + "a") + "hello")
                                v1 = v2
                                if len(v4) > 0 {
                                    v4[0] = fi46
                                }
                            }
                        }
                        if ((("" + "hello") >= "a") || true) {
                            var v51 int64 = -9223372036854775808
                            v52 := v3[191]
                            _ = v51
                            _ = v52
                        }
                        if ((!false) && false) {
                            v4 = append(v4, (fi42 % (((fi46 / (((-(fi42))) | 1))) | 1)))
                            var v53 int8 = -1
                            if len(v4) > 0 {
                                v4[0] = (fi42 ^ (fi46 / ((fi46) | 1)))
                            }
                            _ = v53
                        }
                        if false {
                            var v54 float32 = float32(((0.1 - -2.25) / 1e10))
                            v2 = v2
                            if false {
                                v3[2147483648] = 42
                            } else {
                                if (32767 < 255) {
                                    if true {
                                        v55 := GenS2{F0: 13, F1: 255}
                                        var v56 bool = (float32(0.0) >= 1e10)
                                        var v57 uint64 = 4294967295
                                        _ = v55
                                        _ = v56
                                        _ = v57
                                    }
                                    v58 := GenS2{F0: 127, F1: 0}
                                    _ = v58
                                } else {
                                    if false {
                                        fmt.Println(v2, v2, fi42)
                                        var v59 int = (fi46 | (^(fi42)))
                                        _ = v59
                                    } else {
                                        if (!true) {
                                            v54 += ((((v54 / (v54 * ((((v54 * ((3.14159 + v54) * v54)) - v54) - v54) - 0.1))) / v54) + v54) - 3.14159)
                                            _ = len(v5)
                                            if (-414 > 179) {
                                                fmt.Println(v2, fi46)
                                                fmt.Println(v28, v1)
                                                v54 = v54
                                                var v60 int32 = 435
                                                switch fi46 {
                                                case 996:
                                                    fmt.Println(fi42, v28, v1)
                                                case -2147483648:
                                                    var v61 uint16 = 155
                                                    _ = v61
                                                }
                                                var v62 int = (fi42 >> ((4294967295) & 63))
                                                fmt.Println(fi42, v2, fi42)
                                                _ = v60
                                                _ = v62
                                            } else {
                                                var v63 uint = 470
                                                var v64 int8 = 0
                                                _ = v63
                                                _ = v64
                                            }
                                            v65 := GenS3{F0: false, F1: 255, F2: v2}
                                            v2 = ""
                                            _ = v65
                                        } else {
                                            if true {
                                                var v66 bool = (true || (!(!(255 < 42))))
                                                _ = v66
                                            } else {
                                                if (((0.0 - float32((0.0 + (-2.25 * -2.25)))) - v54) >= (v54 / (1e10 + v54))) {
                                                    if (v1 <= v2) {
                                                        v67 := GenS2{F0: 212, F1: 255}
                                                        _ = v67
                                                    } else {
                                                        v1 = v2
                                                        fmt.Println(v54, v28, v28)
                                                        v54 = 0.0
                                                    }
                                                } else {
                                                    v54 = v54
                                                    switch v1 {
                                                    case "foo bar":
                                                        if (true && false) {
                                                            var v68 uint32 = 664
                                                            fmt.Println(fi42)
                                                            v69 := GenS1{F0: 65535, F1: (fi46 | (fi42 & (fi42 % (((^(fi42))) | 1)))), F2: fi42}
                                                            _ = v68
                                                            _ = v69
                                                        }
                                                    case "hello":
                                                        var v70 string = (v1 + v2)
                                                        _ = v70
                                                    case "\n\t":
                                                        v2 = (v2 + "")
                                                    default:
                                                        if (0 > 972) {
                                                            var v71 uint64 = 1
                                                            switch genFn1(-345, 42, (v28 != -2147483648)) {
                                                            case 2147483647:
                                                                if (!false) {
                                                                    v1 = ((v2 + "") + v1)
                                                                }
                                                            case 302:
                                                                var v72 int8 = 0
                                                                _ = v72
                                                            default:
                                                                v5[fi42] = (((255 > 1) || ((uint16(v71) >= 89) || false)) || false)
                                                            }
                                                            if true {
                                                                var v73 int8 = -53
                                                                var v74 int32 = v28
                                                                _ = v73
                                                                _ = v74
                                                            } else {
                                                                v28 -= v28
                                                                if len(v4) >= 2 {
                                                                    v4 = v4[1:len(v4)]
                                                                }
                                                            }
                                                            _ = v71
                                                        }
                                                    }
                                                    switch v2 {
                                                    case "a":
                                                        switch (v2 + v1) {
                                                        case "a":
                                                            var v75 uint64 = 1
                                                            _ = v75
                                                        case "":
                                                            var v76 int32 = (v28 << ((900) & 63))
                                                            _ = v76
                                                        case "\n\t":
                                                            fmt.Println(v54)
                                                        }
                                                    }
                                                }
                                                v54 -= float32(1e10)
                                                v77 := GenS2{F0: 131, F1: 1}
                                                _ = v77
                                            }
                                            v78 := GenS1{F0: 899, F1: (fi46 ^ 0), F2: (fi42 >> ((0) & 63))}
                                            _ = v78
                                        }
                                        switch -361 {
                                        case 2147483647:
                                            v54 = ((float32((((1.5 / -2.25) - 1e10) / (0.1 / -2.25))) / (v54 * v54)) * (1e10 - ((float32(-2.25) - (-2.25 - (((float32(1.5) * (v54 / (v54 - 1.5))) + (0.0 * v54)) + (((v54 + -2.25) - ((v54 * v54) - (v54 + 1e10))) * (v54 * (0.0 / v54)))))) - 0.0)))
                                        }
                                        v28 += (-(v28))
                                    }
                                    v1 = (v1 + ("\n\t" + "a"))
                                    v79 := GenS4{F0: v1, F1: "", F2: 409}
                                    _ = v79
                                }
                                if (true && false) {
                                    if true {
                                        v2 = ("" + v1)
                                        var v80 uint = 4294967295
                                        _ = v80
                                    } else {
                                        if true {
                                            v54 = (3.14159 / ((v54 * v54) / v54))
                                        } else {
                                            v2 = ((v1 + ("" + (v1 + v2))) + v1)
                                            fmt.Println(fi42)
                                            if (231 == 65535) {
                                                fmt.Println(fi46, v2, v1)
                                                var v81 float32 = 1.5
                                                _ = v81
                                            } else {
                                                v2 = v1
                                                v4 = append(v4, fi46)
                                            }
                                        }
                                        delete(v3, 765)
                                    }
                                    var v82 uint8 = 147
                                    v3[392] = 2147483648
                                    _ = v82
                                }
                            }
                            var v83 uint8 = 125
                            _ = v54
                            _ = v83
                        }
                    }
                    var v84 int = (fi46 & 429)
                    _ = fi46
                    _ = v47
                    _ = v84
                }
                var v85 int16 = 199
                _ = fi42
                _ = v85
            }
            _ = v29
        }
        delete(v3, 4294967295)
        _ = v28
    }
    v5[-2147483648] = ((("foo bar" + "") > "hello") || false)
    v1 = v1
    v1 = v1
    var v86 bool = true
    var v87 int32 = 42
    for fi88 := 0; fi88 < 13; fi88++ {
        for fi89 := 0; fi89 < 18; fi89++ {
            v2 = "foo bar"
            v2 = v2
            _ = fi89
        }
        v90 := GenS4{F0: (v2 + ""), F1: v1, F2: 2147483648}
        _ = fi88
        _ = v90
    }
    v86 = (v86 || true)
    switch ("\n\t" + ("hello" + ("a" + ("hello" + v2)))) {
    case "foo bar":
        _ = len(v4)
    default:
        if (v87 > (^(v87))) {
            switch (((v2 + v1) + "\n\t") + (v1 + v1)) {
            case "a":
                for fi91 := 0; fi91 < 18; fi91++ {
                    switch "" {
                    case "":
                        delete(v3, 42)
                    }
                    _ = fi91
                }
            case "\n\t":
                if ((((float64(3.14159) * (1.5 / 0.1)) / 1.5) + 1e10) != 1e10) {
                    if false {
                        if (-2.25 > 1e10) {
                            delete(v5, 4294967295)
                        }
                        v2 = v2
                        v92 := v3[2147483648]
                        _ = v92
                    }
                    if (1 >= 58) {
                        v93 := GenS2{F0: 102, F1: 0}
                        var v94 uint64 = 42
                        _ = v93
                        _ = v94
                    }
                }
            default:
                fmt.Println(v1, v1)
            }
            var v95 float64 = 1.5
            v95 = float64((0.0 - 1e10))
            _ = v95
        }
    }
    for fi96 := 0; fi96 < 4; fi96++ {
        var v97 uint16 = 32767
        if (!(v86 && v86)) {
            _ = len(v4)
            v86 = ((true || (v86 && (100 == -128))) || false)
        }
        _ = fi96
        _ = v97
    }
    var v98 float32 = float32(-2.25)
    _ = len(v4)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v86:", v86)
    fmt.Println("v87:", v87)
    fmt.Println("v98:", v98)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v86
    _ = v87
    _ = v98
}
