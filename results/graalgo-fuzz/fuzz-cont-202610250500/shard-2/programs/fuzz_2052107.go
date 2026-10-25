package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8; F2 int64 }
type GenS2 struct { F0 int64; F1 uint64; F2 uint64 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint; F1 int64; F2 bool }
type GenS5 struct { F0 int8; F1 uint }
type GenS6 struct { F0 int64; F1 uint32; F2 uint }
type GenS7 struct { F0 uint16; F1 bool; F2 uint32 }
type GenS8 struct { F0 bool; F1 uint }
type GenS9 struct { F0 uint }
type GenS10 struct { F0 bool; F1 bool }
func genFn0(p0 uint32) uint {
    var v5 int64 = 42
    for fi6 := 0; fi6 < 1; fi6++ {
        if (!(!false)) {
            p0 -= p0
            var v7 uint32 = (p0 ^ 42)
            _ = v7
        } else {
            p0 = p0
        }
        if ((v5 & -9223372036854775808) > (v5 / ((9223372036854775807) | 1))) {
            for fi8 := 0; fi8 < 19; fi8++ {
                var v9 uint16 = 32767
                if (723 < 777) {
                    v9 = v9
                }
                p0 -= (p0 ^ p0)
                _ = fi8
                _ = v9
            }
        }
        for fi10 := 0; fi10 < 9; fi10++ {
            var v11 uint16 = 1
            _ = fi10
            _ = v11
        }
        _ = fi6
    }
    _ = p0
    _ = v5
    return 18446744073709551615
}
func genFn1(p0 int32, p1 uint32, p2 string) int8 {
    v12 := GenS1{F0: 1000, F1: 1, F2: 42}
    var v13 int16 = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v13
    return 42
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int32 = -182
    var v3 string = ((("foo bar" + ("foo bar" + "foo bar")) + ("hello" + (((("a" + "\n\t") + "\n\t") + ("a" + "")) + "foo bar"))) + "a")
    v4 := make(map[uint32]uint32)
    var v14 uint16 = 255
    switch 127 {
    case 32768:
        delete(v1, 141)
    }
    if false {
        fmt.Println(v3)
        v15 := v1[512]
        var v16 uint8 = 2
        _ = v15
        _ = v16
    }
    for fi17 := 0; fi17 < 17; fi17++ {
        for fi18 := 0; fi18 < 17; fi18++ {
            _ = len(v1)
            switch (fi18 + fi17) {
            case 0:
                if ((float32(0.1) >= -2.25) && false) {
                    var v19 bool = (42 != 1000)
                    v20 := GenS1{F0: 42, F1: 42, F2: -1}
                    v1[0] = 2147483648
                    _ = v19
                    _ = v20
                } else {
                    fmt.Println(fi18, v14, v2)
                    v14 += v14
                    fmt.Println(fi17, v14, v3)
                }
            case 9223372036854775807:
                switch 4294967295 {
                case 0:
                    if false {
                        if (v14 == 2) {
                            var v21 bool = ("\n\t" <= "")
                            if (4294967295 >= 620) {
                                var v22 float64 = 3.14159
                                var v23 uint8 = 2
                                v23 -= v23
                                v24 := GenS2{F0: 1, F1: 42, F2: 42}
                                _ = v22
                                _ = v23
                                _ = v24
                            }
                            var v25 bool = v21
                            _ = v21
                            _ = v25
                        }
                        var v26 string = (v3 + (v3 + v3))
                        if true {
                            v26 = v3
                            switch fi18 {
                            case 42:
                                switch fi18 {
                                case 1:
                                    v27 := v4[42]
                                    _ = v27
                                case -943:
                                    fmt.Println(v14)
                                }
                            case 1:
                                v2 += (v2 % (((v2 & (^(v2)))) | 1))
                            case -673:
                                fmt.Println(v14, fi18, v2)
                            }
                        } else {
                            if (((-411 > 42) && true) && (18446744073709551615 <= 4294967295)) {
                                fmt.Println(fi17, v2)
                                var v28 string = "\n\t"
                                _ = v28
                            }
                        }
                        _ = v26
                    }
                case 9223372036854775807:
                    var v29 int8 = 0
                    _ = v29
                default:
                    v14 = 589
                }
            case 128:
                if (!(3.14159 <= (1e10 / (3.14159 + ((((0.1 + 1.5) / float64(1.5)) - ((0.1 - float64(float32(-2.25))) + 3.14159)) + -2.25))))) {
                    var v30 float64 = float64(-2.25)
                    v31 := GenS3{F0: 638}
                    if (42 == genFn0(2147483648)) {
                        if (!false) {
                            v32 := GenS2{F0: 42, F1: 177, F2: 18446744073709551615}
                            _ = v32
                        }
                        fmt.Println(v14, v30)
                        v3 = v3
                    }
                    _ = v30
                    _ = v31
                } else {
                    if false {
                        if (0.1 != 1.5) {
                            v33 := GenS2{F0: -9223372036854775808, F1: 0, F2: 18446744073709551615}
                            if false {
                                var v34 uint = genFn0(2147483647)
                                v35 := GenS4{F0: v34, F1: 0, F2: (fi17 <= (fi18 * fi17))}
                                v2 -= v2
                                _ = len(v1)
                                _ = v34
                                _ = v35
                            }
                            fmt.Println(v2, v14, fi17)
                            var v36 string = "\n\t"
                            var v37 uint16 = (^(v14))
                            _ = v33
                            _ = v36
                            _ = v37
                        } else {
                            switch (v3 + "\n\t") {
                            case "a":
                                fmt.Println(v14)
                            }
                        }
                        v3 = v3
                    } else {
                        var v38 int64 = -565
                        v39 := GenS1{F0: 32767, F1: -1, F2: 556}
                        if false {
                            v40 := GenS5{F0: -1, F1: genFn0(0)}
                            _ = v40
                        } else {
                            v41 := GenS1{F0: 42, F1: genFn1(42, 64, v3), F2: 1}
                            delete(v1, 244)
                            if true {
                                v2 -= v2
                                var v42 uint64 = 902
                                _ = v42
                            }
                            _ = v41
                        }
                        if false {
                            var v43 int64 = (-(v38))
                            switch "foo bar" {
                            case "foo bar":
                                v44 := GenS1{F0: 0, F1: 42, F2: (v38 % ((v43) | 1))}
                                _ = v44
                            case "":
                                if (0.0 != 0.0) {
                                    v43 = int64(fi18)
                                    fmt.Println(v43, fi18, fi17)
                                    var v45 int = 27
                                    _ = v45
                                } else {
                                    var v46 string = (("foo bar" + "\n\t") + v3)
                                    var v47 string = "a"
                                    switch (fi17 / (((fi17 | fi17)) | 1)) {
                                    case 2147483647:
                                        var v48 int16 = 99
                                        _ = v48
                                    default:
                                        switch ("hello" + (v46 + v3)) {
                                        case "\n\t":
                                            fmt.Println(v47, fi18)
                                            fmt.Println(v3)
                                        case "a":
                                            fmt.Println(fi17)
                                        default:
                                            var v49 uint32 = 269
                                            _ = v49
                                        }
                                    }
                                    _ = v46
                                    _ = v47
                                }
                                var v50 float64 = (0.0 * float64(1e10))
                                _ = v50
                            default:
                                var v51 int = fi18
                                _ = v51
                            }
                            v52 := v4[2147483647]
                            _ = v43
                            _ = v52
                        } else {
                            switch fi18 {
                            case 577:
                                v1[42] = 937
                            case 716:
                                v53 := []int{-1, 0, fi17, (fi18 << ((1) & 63))}
                                _ = v53
                            default:
                                var v54 uint64 = 18446744073709551615
                                _ = v54
                            }
                        }
                        _ = v38
                        _ = v39
                    }
                    v2 = v2
                    var v55 int8 = 107
                    _ = v55
                }
            default:
                switch v3 {
                case "a":
                    fmt.Println(v3, v14, fi17)
                case "hello":
                    var v56 int32 = -1
                    _ = v56
                case "foo bar":
                    fmt.Println(v2)
                }
            }
            _ = fi18
        }
        fmt.Println(v2)
        _ = fi17
    }
    v57 := GenS6{F0: 42, F1: 4294967295, F2: 834}
    var v58 int8 = 42
    switch 128 {
    case 1:
        v59 := []int{4294967295, 157, 42, 4294967295}
        _ = v59
    case -1:
        switch (v3 + v3) {
        case "":
            if true {
                var v60 int32 = v2
                fmt.Println(v2, v2, v14)
                if true {
                    for fi61 := 0; fi61 < 18; fi61++ {
                        v62 := GenS2{F0: 9223372036854775807, F1: 4294967295, F2: uint64(v14)}
                        _ = len(v1)
                        if (965 > 42) {
                            v63 := GenS1{F0: -31, F1: -128, F2: 0}
                            var v64 string = v3
                            _ = v63
                            _ = v64
                        } else {
                            if (3.14159 == 0.1) {
                                if false {
                                    var v65 string = v3
                                    _ = v65
                                }
                                var v66 bool = (!((1.5 / 1.0) != 3.14159))
                                _ = v66
                            } else {
                                switch (fi61 + (fi61 + fi61)) {
                                case -367:
                                    switch (-(fi61)) {
                                    case 9223372036854775807:
                                        v67 := []int{-1, 1, 143}
                                        _ = v67
                                    case 128:
                                        if false {
                                            var v68 float64 = 3.14159
                                            for fi69 := 0; fi69 < 14; fi69++ {
                                                var v70 uint16 = v14
                                                v68 += -2.25
                                                switch fi61 {
                                                case -9223372036854775808:
                                                    var v71 int8 = genFn1((v60 * -2147483648), 2147483647, ("foo bar" + (v3 + v3)))
                                                    _ = v71
                                                default:
                                                    if false {
                                                        if (!true) {
                                                            var v72 uint32 = 42
                                                            switch (fi61 | fi61) {
                                                            case -1:
                                                                v70 += v14
                                                            default:
                                                                v73 := []int{fi69, (fi61 % (((fi69 >> ((2) & 63))) | 1)), fi61, fi61}
                                                                _ = v73
                                                            }
                                                            if (1.5 < float64(1.5)) {
                                                                fmt.Println(v60, v72, v72)
                                                            } else {
                                                                switch fi69 {
                                                                case 9223372036854775807:
                                                                    if true {
                                                                        v58 = 1
                                                                        v14 = v70
                                                                    }
                                                                default:
                                                                    var v74 int = -459
                                                                    _ = v74
                                                                }
                                                                switch fi69 {
                                                                case -9223372036854775808:
                                                                    switch (v3 + v3) {
                                                                    case "":
                                                                        var v75 uint8 = 42
                                                                        _ = v75
                                                                    case "hello":
                                                                        v58 = -109
                                                                    case "\n\t":
                                                                        v76 := GenS3{F0: 42}
                                                                        _ = v76
                                                                    }
                                                                default:
                                                                    if false {
                                                                        switch ("a" + v3) {
                                                                        case "\n\t":
                                                                            if (0 == 255) {
                                                                                var v77 float32 = (((float32((((((((v68 * v68) / (float64(1.5) * (v68 + ((v68 * (float64(0.1) + v68)) + 3.14159)))) * ((v68 - ((0.0 - ((v68 + (((float64(((1e10 - (float32(float64(1e10)) + 0.0)) + ((3.14159 / 3.14159) * (1e10 - 1.5)))) + (float64(float32(v68)) - 1.5)) / (v68 / 3.14159)) - -2.25)) - ((0.0 * v68) * v68))) + v68)) / (((((0.0 - -2.25) / (v68 * -2.25)) / float64(((((0.1 * 0.1) + 1e10) + 3.14159) + 1.5))) + 1.5) - v68))) - (v68 / 1.5)) - v68) + -2.25) * float64(0.1))) / -2.25) + 0.0) - ((3.14159 + 3.14159) * (float32((0.0 - (((0.1 / 1e10) * v68) * 1e10))) / 0.1)))
                                                                                v68 -= (v68 / v68)
                                                                                _ = v77
                                                                            }
                                                                        case "":
                                                                            v78 := []int{fi61}
                                                                            _ = v78
                                                                        }
                                                                    }
                                                                }
                                                            }
                                                            _ = v72
                                                        }
                                                    } else {
                                                        fmt.Println(v60)
                                                        var v79 bool = (!true)
                                                        _ = v79
                                                    }
                                                    v60 -= v60
                                                }
                                                _ = fi69
                                                _ = v70
                                            }
                                            _ = v68
                                        } else {
                                            switch (fi61 << ((4294967295) & 63)) {
                                            case 0:
                                                _ = len(v1)
                                            }
                                        }
                                    }
                                case 32768:
                                    switch (v3 + "\n\t") {
                                    case "\n\t":
                                        var v80 int32 = (v60 % ((v2) | 1))
                                        _ = v80
                                    default:
                                        v3 = "\n\t"
                                    }
                                case 32767:
                                    v81 := GenS7{F0: v14, F1: true, F2: 2147483648}
                                    _ = v81
                                }
                            }
                            v82 := GenS6{F0: 278, F1: 2, F2: 4294967295}
                            _ = v82
                        }
                        _ = fi61
                        _ = v62
                    }
                } else {
                    v83 := []int{}
                    v84 := GenS7{F0: 2, F1: ("a" != "foo bar"), F2: 4294967295}
                    _ = v83
                    _ = v84
                }
                _ = v60
            }
        case "\n\t":
            var v85 float32 = 0.0
            _ = v85
        }
    }
    v58 -= v58
    v86 := v4[0]
    if false {
        for fi87 := 0; fi87 < 7; fi87++ {
            fmt.Println(v58, v14)
            for fi88 := 0; fi88 < 11; fi88++ {
                v2 -= (-(v2))
                _ = fi88
            }
            _ = fi87
        }
        var v89 uint32 = (v86 << ((2) & 63))
        fmt.Println(v89, v14)
        _ = v89
    }
    _ = len(v1)
    if (v58 == -128) {
        var v90 uint8 = 93
        v2 -= int32(v90)
        fmt.Println(v90)
        _ = v90
    } else {
        delete(v4, 4294967295)
        v91 := GenS8{F0: (!(-843 != 1)), F1: genFn0((^(v86)))}
        fmt.Println(v2)
        _ = v91
    }
    v92 := GenS9{F0: 624}
    var v93 string = v3
    if (true && (float32(((-2.25 * ((1.5 + 0.1) - 3.14159)) * 1e10)) == 3.14159)) {
        v94 := GenS2{F0: -1, F1: 42, F2: 0}
        for fi95 := 0; fi95 < 9; fi95++ {
            v96 := GenS10{F0: ((v14 | v14) >= (^(v14))), F1: (true || false)}
            for fi97 := 0; fi97 < 16; fi97++ {
                v98 := []int{fi97, fi97, (fi95 * fi95), 128}
                v99 := GenS4{F0: 337, F1: 0, F2: true}
                fmt.Println(v2)
                _ = fi97
                _ = v98
                _ = v99
            }
            _ = fi95
            _ = v96
        }
        _ = v94
    } else {
        v2 += v2
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v58:", v58)
    fmt.Println("v86:", v86)
    fmt.Println("v93:", v93)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v57
    _ = v58
    _ = v86
    _ = v92
    _ = v93
}
