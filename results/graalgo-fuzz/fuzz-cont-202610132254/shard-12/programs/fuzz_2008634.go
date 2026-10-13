package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16; F2 string }
type GenS2 struct { F0 uint64; F1 int; F2 int16 }
type GenS3 struct { F0 uint32; F1 int16; F2 uint16 }
type GenS4 struct { F0 string; F1 uint16; F2 int64 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 int }
type GenS7 struct { F0 string; F1 uint; F2 int16 }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 uint32; F1 int8 }
func genFn0() int32 {
    var v4 int = 0
    fmt.Println(v4)
    var v5 uint64 = 484
    var v6 string = ""
    v6 = (v6 + ("foo bar" + v6))
    _ = v4
    _ = v5
    _ = v6
    return 2147483647
}
func genFn1(p0 uint32) uint64 {
    switch "hello" {
    case "\n\t":
        v7 := []int{-888, -9223372036854775808}
        _ = v7
    case "hello":
        var v8 float64 = 1.5
        _ = v8
    }
    for fi9 := 0; fi9 < 17; fi9++ {
        p0 -= p0
        if false {
            var v10 bool = true
            _ = v10
        }
        fmt.Println(fi9, fi9)
        _ = fi9
    }
    v11 := GenS1{F0: 42, F1: 1000, F2: (("\n\t" + "\n\t") + ("a" + "foo bar"))}
    var v12 bool = ((true && true) && ((p0 == p0) || (-1 != -2147483648)))
    _ = p0
    _ = v11
    _ = v12
    return 18446744073709551615
}

func main() {
    var v1 float32 = ((3.14159 - 1.5) * 1e10)
    var v2 int16 = -1
    v3 := make(map[string]bool)
    switch "\n\t" {
    case "":
        v13 := GenS1{F0: 32767, F1: (v2 + v2), F2: ""}
        _ = v13
    case "\n\t":
        fmt.Println(v2, v2)
    case "hello":
        for fi14 := 0; fi14 < 11; fi14++ {
            _ = len(v3)
            var v15 uint8 = 2
            v16 := GenS2{F0: genFn1(2), F1: 9223372036854775807, F2: v2}
            _ = fi14
            _ = v15
            _ = v16
        }
    default:
        var v17 int8 = 29
        _ = v17
    }
    if (1 <= 2147483647) {
        switch 878 {
        case -2147483648:
            _ = len(v3)
        case 4294967295:
            var v18 int32 = genFn0()
            _ = v18
        case -57:
            if true {
                v1 = v1
                switch "" {
                case "a":
                    var v19 uint32 = 149
                    _ = v19
                case "":
                    delete(v3, ("" + "foo bar"))
                default:
                    v2 -= 42
                }
                switch 42 {
                case 0:
                    if false {
                        v1 = float32(0.0)
                        if false {
                            switch 42 {
                            case 9223372036854775807:
                                var v20 float32 = ((1e10 * v1) + (0.0 / 1.0))
                                _ = v20
                            case -9223372036854775808:
                                for fi21 := 0; fi21 < 12; fi21++ {
                                    var v22 uint = 678
                                    _ = fi21
                                    _ = v22
                                }
                            case 128:
                                v23 := v3[("" + "a")]
                                _ = v23
                            }
                            v24 := []int{128, 726}
                            _ = v24
                        } else {
                            _ = len(v3)
                            fmt.Println(v1, v2)
                            v25 := GenS3{F0: 2147483648, F1: v2, F2: 818}
                            _ = v25
                        }
                    }
                case 42:
                    if (2 != 1) {
                        v1 += (((3.14159 + 3.14159) - (1e10 / v1)) - 0.0)
                        var v26 float64 = float64(v1)
                        v2 -= -1
                        _ = v26
                    }
                default:
                    v27 := []int{32767, 448}
                    _ = v27
                }
            }
        default:
            if (!true) {
                v28 := GenS4{F0: ("hello" + ""), F1: 291, F2: -1}
                var v29 float64 = -2.25
                _ = v28
                _ = v29
            } else {
                var v30 uint = 18446744073709551615
                for fi31 := 0; fi31 < 16; fi31++ {
                    switch "a" {
                    case "hello":
                        var v32 uint8 = 124
                        _ = v32
                    }
                    _ = fi31
                }
                v33 := GenS3{F0: 4294967295, F1: 672, F2: 65535}
                _ = v30
                _ = v33
            }
        }
        if (((!false) && (true && true)) && (!(genFn1(2147483648) == genFn1(845)))) {
            v1 -= 1.5
            var v34 int16 = (v2 % ((v2) | 1))
            for fi35 := 0; fi35 < 7; fi35++ {
                if true {
                    v36 := []int{(fi35 ^ 38), 0, int(v2), (^(fi35))}
                    if (false || (v2 <= (v34 & v34))) {
                        for fi37 := 0; fi37 < 19; fi37++ {
                            var v38 string = "a"
                            v34 += (v2 >> ((0) & 63))
                            v2 = 0
                            _ = fi37
                            _ = v38
                        }
                        for fi39 := 0; fi39 < 11; fi39++ {
                            v34 = (v2 ^ v2)
                            var v40 float32 = 0.1
                            v41 := GenS5{F0: (-1 >= 42)}
                            _ = fi39
                            _ = v40
                            _ = v41
                        }
                    } else {
                        v42 := GenS6{F0: 339}
                        fmt.Println(fi35, fi35)
                        v34 = (^(v34))
                        _ = v42
                    }
                    _ = v36
                } else {
                    var v43 int8 = 32
                    _ = v43
                }
                _ = fi35
            }
            _ = v34
        }
        v2 -= v2
    } else {
        if ((-128 == 0) || true) {
            v1 = float32((1e10 + 0.0))
        }
    }
    if (!(-128 > 100)) {
        v44 := GenS1{F0: 273, F1: (v2 >> ((42) & 63)), F2: "hello"}
        fmt.Println(v2, v1)
        if true {
            var v45 uint64 = 42
            var v46 int64 = 42
            _ = v45
            _ = v46
        } else {
            v1 = v1
            v2 = (v2 & (^(v2)))
            for fi47 := 0; fi47 < 2; fi47++ {
                fmt.Println(v1)
                if true {
                    fmt.Println(v2, fi47)
                    v2 = (v2 ^ -18)
                    for fi48 := 0; fi48 < 18; fi48++ {
                        if (1 != 2) {
                            var v49 uint = 479
                            switch ("" + "") {
                            case "a":
                                v2 -= (v2 + v2)
                            default:
                                v50 := []int{(fi47 >> (((v49 * (^(v49)))) & 63)), (fi47 + fi48), fi48}
                                _ = v50
                            }
                            _ = v49
                        }
                        switch "hello" {
                        case "a":
                            var v51 int8 = -1
                            _ = v51
                        case "hello":
                            v52 := GenS7{F0: "a", F1: 42, F2: (-(v2))}
                            _ = v52
                        default:
                            if false {
                                v53 := []int{fi48, (fi47 << ((0) & 63)), (fi48 + fi48), (^(fi48))}
                                if true {
                                    v2 = v2
                                }
                                if (!(!false)) {
                                    var v54 int16 = v2
                                    if (1 < -172) {
                                        v2 = v54
                                        fmt.Println(v54, v54, v54)
                                        v55 := GenS2{F0: genFn1(42), F1: (fi47 | fi48), F2: (v54 | v2)}
                                        _ = v55
                                    }
                                    _ = v54
                                }
                                _ = v53
                            } else {
                                fmt.Println(fi48, v1, fi48)
                                var v56 int8 = 0
                                _ = v56
                            }
                        }
                        var v57 uint8 = 128
                        _ = fi48
                        _ = v57
                    }
                }
                _ = fi47
            }
        }
        _ = v44
    }
    if false {
        v58 := []int{9223372036854775807, -1, 0, 0}
        fmt.Println(v2)
        _ = v58
    } else {
        fmt.Println(v2)
        if true {
            var v59 int32 = 2147483647
            for fi60 := 0; fi60 < 5; fi60++ {
                var v61 int32 = (^(v59))
                if true {
                    var v62 int8 = 0
                    _ = v62
                }
                fmt.Println(fi60, fi60, fi60)
                _ = fi60
                _ = v61
            }
            switch 0 {
            case -9223372036854775808:
                if (true && false) {
                    v59 -= (v59 + (v59 >> ((0) & 63)))
                    for fi63 := 0; fi63 < 10; fi63++ {
                        v59 += (v59 << ((1) & 63))
                        _ = fi63
                    }
                    delete(v3, "foo bar")
                } else {
                    fmt.Println(v1, v59, v1)
                    var v64 uint = 672
                    v64 = v64
                    _ = v64
                }
            default:
                var v65 uint8 = 0
                _ = v65
            }
            _ = v59
        }
        var v66 string = ("" + "hello")
        _ = v66
    }
    if true {
        if false {
            if (127 <= 44) {
                v2 = (v2 & (v2 ^ -727))
                v2 -= v2
            }
        } else {
            v1 -= (v1 - (v1 * (0.0 - ((1.5 - ((-2.25 * v1) * (((float32(float64(float32(1.5))) + ((1.5 - float32(float64(((((v1 / 1.0) * v1) - v1) - (v1 + v1))))) - v1)) - -2.25) * v1))) * (v1 - (v1 / 1.0))))))
            if ((255 <= 18446744073709551615) && (true || true)) {
                for fi67 := 0; fi67 < 8; fi67++ {
                    v68 := []int{fi67, (fi67 >> ((4294967295) & 63))}
                    v1 -= 0.1
                    _ = fi67
                    _ = v68
                }
                switch 42 {
                case 4294967295:
                    v69 := []int{2147483647}
                    _ = v69
                case -9223372036854775808:
                    _ = len(v3)
                case -505:
                    fmt.Println(v2)
                default:
                    v2 += v2
                }
                fmt.Println(v2, v1)
            }
            v1 = ((v1 + v1) + float32(1.5))
        }
        for fi70 := 0; fi70 < 10; fi70++ {
            if (3.14159 > (3.14159 + v1)) {
                if true {
                    switch "a" {
                    case "hello":
                        if (492 <= 4294967295) {
                            switch "hello" {
                            case "a":
                                for fi71 := 0; fi71 < 14; fi71++ {
                                    var v72 uint32 = 465
                                    var v73 uint8 = 127
                                    _ = fi71
                                    _ = v72
                                    _ = v73
                                }
                            default:
                                v74 := GenS8{F0: false}
                                _ = v74
                            }
                        }
                    case "foo bar":
                        for fi75 := 0; fi75 < 10; fi75++ {
                            _ = len(v3)
                            _ = fi75
                        }
                    case "a":
                        v2 -= (v2 * (v2 & (v2 / (((v2 >> ((18446744073709551615) & 63))) | 1))))
                    default:
                        v76 := GenS9{F0: 441, F1: 0}
                        _ = v76
                    }
                    var v77 uint32 = 2147483648
                    fmt.Println(fi70)
                    _ = v77
                } else {
                    fmt.Println(fi70, fi70, v2)
                    for fi78 := 0; fi78 < 14; fi78++ {
                        switch -968 {
                        case -2147483648:
                            switch (fi70 % (((fi70 % ((32767) | 1))) | 1)) {
                            case 32768:
                                var v79 string = ("\n\t" + "hello")
                                _ = v79
                            case 128:
                                if false {
                                    v80 := GenS5{F0: (100 >= -32)}
                                    var v81 uint16 = 736
                                    fmt.Println(v81)
                                    _ = v80
                                    _ = v81
                                }
                            default:
                                v82 := []int{(fi78 | (-(fi70))), (^(fi78))}
                                _ = v82
                            }
                        case -9223372036854775808:
                            v83 := []int{(fi78 + fi70)}
                            _ = v83
                        default:
                            var v84 int32 = 42
                            _ = v84
                        }
                        _ = fi78
                    }
                }
                var v85 int32 = 42
                for fi86 := 0; fi86 < 18; fi86++ {
                    v87 := GenS4{F0: ("\n\t" + ""), F1: 1, F2: 9223372036854775807}
                    _ = fi86
                    _ = v87
                }
                _ = v85
            }
            if (255 <= 180) {
                for fi88 := 0; fi88 < 16; fi88++ {
                    v89 := []int{}
                    if true {
                        delete(v3, ("" + "foo bar"))
                        v1 -= (((0.0 / v1) - ((float32(1.5) + v1) / v1)) - (v1 / (float32(float64((0.0 + v1))) + 0.1)))
                    } else {
                        v1 += 0.0
                        v1 = v1
                        switch -159 {
                        case -9223372036854775808:
                            v1 += v1
                        }
                    }
                    var v90 uint64 = 42
                    var v91 uint32 = 294
                    switch fi88 {
                    case 4294967295:
                        v91 -= 2
                    }
                    _ = fi88
                    _ = v89
                    _ = v90
                    _ = v91
                }
            } else {
                v2 += v2
                v1 = 1.5
                var v92 int16 = (^(v2))
                _ = v92
            }
            v93 := []int{fi70, (-(fi70)), 42, fi70}
            _ = fi70
            _ = v93
        }
    }
    v2 = (^(v2))
    for fi94 := 0; fi94 < 18; fi94++ {
        switch ("\n\t" + (("hello" + "hello") + "\n\t")) {
        case "":
            v95 := []int{(fi94 % ((fi94) | 1))}
            _ = v95
        default:
            v96 := GenS8{F0: ((!(true || false)) && (-14 < -36))}
            _ = v96
        }
        _ = fi94
    }
    v1 += v1
    v2 -= (v2 % ((-32768) | 1))
    fmt.Println(v2, v2)
    delete(v3, "foo bar")
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v3
}
