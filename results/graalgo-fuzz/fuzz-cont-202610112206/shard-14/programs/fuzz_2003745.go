package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int64; F1 int }
type GenS3 struct { F0 uint8; F1 bool }
type GenS4 struct { F0 int64; F1 uint8 }
type GenS5 struct { F0 uint8; F1 int32; F2 uint }
func genFn0() uint8 {
    var v6 string = "a"
    v7 := []int{0}
    _ = v6
    _ = v7
    return 219
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 float32 = ((1e10 / 3.14159) + 0.0)
    v3 := []int{127, 4294967295, 9223372036854775807}
    v4 := make(map[string]string)
    var v5 uint8 = 101
    fmt.Println(v5)
    var v8 uint = 4294967295
    v2 = (((1.5 + 0.0) * 1.5) + v2)
    fmt.Println(v2, v8, v2)
    v9 := v4[("a" + ("hello" + "\n\t"))]
    for fi10 := 0; fi10 < 8; fi10++ {
        delete(v4, v9)
        v8 += 4294967295
        if false {
            v8 = (v8 * v8)
            for fi11 := 0; fi11 < 19; fi11++ {
                switch v9 {
                case "hello":
                    v12 := GenS1{F0: (v5 % ((0) | 1))}
                    _ = v12
                case "foo bar":
                    fmt.Println(fi11, v8)
                default:
                    var v13 int8 = 127
                    _ = v13
                }
                v14 := GenS1{F0: v5}
                var v15 int = (-(fi10))
                _ = fi11
                _ = v14
                _ = v15
            }
        } else {
            switch (fi10 << ((v8) & 63)) {
            case -2147483648:
                if false {
                    var v16 int64 = -1
                    if (129 != (v5 & v5)) {
                        var v17 int8 = -128
                        var v18 string = (v9 + v9)
                        var v19 bool = (true || true)
                        _ = v17
                        _ = v18
                        _ = v19
                    }
                    _ = v16
                } else {
                    v20 := GenS1{F0: v5}
                    switch -243 {
                    case 1:
                        if false {
                            v4[v9] = ("hello" + (v9 + ((((v9 + (v9 + "a")) + "") + (v9 + (v9 + ((v9 + "") + (v9 + v9))))) + v9)))
                        } else {
                            switch (("a" + ("a" + v9)) + v9) {
                            case "a":
                                fmt.Println(v8)
                            default:
                                var v21 int64 = -9223372036854775808
                                _ = v21
                            }
                            fmt.Println(v9, v9)
                        }
                    case 32768:
                        for fi22 := 0; fi22 < 9; fi22++ {
                            if false {
                                var v23 uint8 = 127
                                _ = v23
                            } else {
                                var v24 bool = (0 >= 0)
                                v24 = true
                                _ = v24
                            }
                            _ = fi22
                        }
                    case 9223372036854775807:
                        for fi25 := 0; fi25 < 5; fi25++ {
                            var v26 int8 = 42
                            var v27 uint32 = 2147483647
                            _ = fi25
                            _ = v26
                            _ = v27
                        }
                    }
                    if (!false) {
                        v28 := v1[1]
                        _ = v28
                    } else {
                        switch (-(fi10)) {
                        case -9223372036854775808:
                            if (("a" + (v9 + v9)) > ("\n\t" + v9)) {
                                for fi29 := 0; fi29 < 4; fi29++ {
                                    fmt.Println(v8, v2)
                                    v5 += v5
                                    _ = fi29
                                }
                                var v30 int32 = -455
                                _ = v30
                            } else {
                                switch (fi10 / (((fi10 + fi10)) | 1)) {
                                case 127:
                                    var v31 uint64 = 42
                                    _ = v31
                                case -9223372036854775808:
                                    var v32 float32 = v2
                                    _ = v32
                                case 32767:
                                    v9 = v9
                                }
                            }
                        case 459:
                            fmt.Println(v2, v9, v5)
                        case -383:
                            v33 := GenS2{F0: -988, F1: fi10}
                            _ = v33
                        default:
                            v34 := GenS2{F0: 9223372036854775807, F1: 287}
                            _ = v34
                        }
                        v8 = (v8 | (v8 + uint(v5)))
                        v5 = v5
                    }
                    _ = v20
                }
            }
            if (65535 != 42) {
                if (!true) {
                    if (v5 >= 39) {
                        v2 = (float32(((3.14159 / 1e10) - 0.0)) / v2)
                        var v35 string = v9
                        var v36 float32 = v2
                        _ = v35
                        _ = v36
                    }
                    for fi37 := 0; fi37 < 2; fi37++ {
                        if true {
                            if (241 >= 42) {
                                v38 := GenS1{F0: (v5 >> (((v8 & 94)) & 63))}
                                _ = v38
                            } else {
                                var v39 string = ((v9 + v9) + "foo bar")
                                _ = v39
                            }
                            if len(v3) > 0 {
                                v3[0] = (^(fi37))
                            }
                            if false {
                                fmt.Println(v2, v8, fi10)
                                if false {
                                    var v40 uint = v8
                                    _ = v40
                                } else {
                                    var v41 int32 = 2147483647
                                    _ = v41
                                }
                            } else {
                                v9 = (("a" + (v9 + v9)) + (v9 + v9))
                            }
                        }
                        _ = fi37
                    }
                    var v42 int8 = -128
                    _ = v42
                }
                v43 := GenS3{F0: (v5 % (((v5 ^ uint8(v8))) | 1)), F1: ((false || false) && (v8 <= 508))}
                switch fi10 {
                case 32767:
                    fmt.Println(v2, fi10, v8)
                case -866:
                    for fi44 := 0; fi44 < 7; fi44++ {
                        var v45 bool = false
                        v2 = v2
                        _ = fi44
                        _ = v45
                    }
                case 1:
                    v46 := GenS1{F0: v5}
                    _ = v46
                }
                _ = v43
            }
        }
        _ = fi10
    }
    var v47 uint32 = 1
    v48 := GenS3{F0: v5, F1: false}
    for fi49 := 0; fi49 < 15; fi49++ {
        var v50 string = "foo bar"
        fmt.Println(v9)
        if len(v3) > 0 {
            v3[0] = fi49
        }
        _ = fi49
        _ = v50
    }
    var v51 int16 = -1
    if ((((v2 - 1e10) - float32(3.14159)) / v2) < ((((-2.25 - 1e10) / v2) + ((v2 / -2.25) - 0.1)) * 3.14159)) {
        var v52 float64 = 0.0
        if (65535 == 0) {
            v1[(v47 >> ((v8) & 63))] = 834
        } else {
            fmt.Println(v47)
        }
        if true {
            if false {
                v51 -= v51
            } else {
                if false {
                    v53 := GenS2{F0: -282, F1: int(v5)}
                    if false {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        _ = len(v3)
                    } else {
                        v54 := GenS4{F0: -1, F1: (v5 + genFn0())}
                        switch "" {
                        case "\n\t":
                            v52 = v52
                        case "":
                            var v55 bool = (false && (356 < 2))
                            _ = v55
                        case "a":
                            for fi56 := 0; fi56 < 3; fi56++ {
                                var v57 int16 = (v51 + (v51 * -32768))
                                _ = fi56
                                _ = v57
                            }
                        default:
                            v47 -= (v47 % (((^(v47))) | 1))
                        }
                        var v58 uint32 = 391
                        _ = v54
                        _ = v58
                    }
                    switch ((v9 + ((v9 + v9) + v9)) + (("" + v9) + v9)) {
                    case "\n\t":
                        var v59 uint64 = 2
                        _ = v59
                    case "a":
                        v60 := GenS3{F0: (v5 - (v5 | (v5 | genFn0()))), F1: (0.0 < 1.5)}
                        _ = v60
                    }
                    _ = v53
                } else {
                    v51 += v51
                }
                v51 = (-(v51))
                v8 += (^(v8))
            }
        }
        _ = v52
    } else {
        switch v9 {
        case "\n\t":
            if len(v3) > 0 {
                v3[0] = -915
            }
        default:
            _ = len(v1)
        }
    }
    if false {
        var v61 uint16 = 1
        _ = v61
    } else {
        switch 128 {
        case -75:
            var v62 float32 = (v2 - ((0.1 + 0.0) + (v2 / float32(-2.25))))
            _ = v62
        case 9223372036854775807:
            fmt.Println(v9, v47, v2)
        case 2147483647:
            var v63 int = 109
            _ = v63
        }
        for fi64 := 0; fi64 < 1; fi64++ {
            switch v9 {
            case "":
                var v65 int = (fi64 | -2147483648)
                _ = v65
            case "foo bar":
                v51 -= (v51 >> ((57) & 63))
            default:
                var v66 bool = true
                _ = v66
            }
            v5 += genFn0()
            for fi67 := 0; fi67 < 18; fi67++ {
                v68 := GenS3{F0: uint8(v8), F1: false}
                switch ("foo bar" + (v9 + v9)) {
                case "\n\t":
                    if false {
                        v8 = uint(v5)
                    } else {
                        v69 := GenS4{F0: 1, F1: (v5 * (v5 * v5))}
                        v3 = append(v3, -2147483648)
                        _ = v69
                    }
                default:
                    var v70 float32 = ((((float32(1e10) + float32(1.5)) / v2) + -2.25) / v2)
                    _ = v70
                }
                if (false || true) {
                    switch (fi64 ^ (fi64 ^ 4294967295)) {
                    case 127:
                        v71 := GenS5{F0: 0, F1: -1, F2: (v8 ^ uint(v47))}
                        _ = v71
                    case 174:
                        var v72 int = fi64
                        _ = v72
                    }
                } else {
                    switch v9 {
                    case "hello":
                        v5 -= (v5 + (v5 / ((genFn0()) | 1)))
                    case "a":
                        v47 -= (v47 | v47)
                    case "foo bar":
                        _ = len(v3)
                    default:
                        var v73 uint32 = (^(v47))
                        _ = v73
                    }
                    var v74 uint64 = 110
                    fmt.Println(v9, v5, v8)
                    switch (("\n\t" + v9) + "a") {
                    case "hello":
                        delete(v1, (v47 << (((^(v8))) & 63)))
                    default:
                        v75 := GenS1{F0: (v5 | (v5 << ((v8) & 63)))}
                        _ = v75
                    }
                    _ = v74
                }
                _ = fi67
                _ = v68
            }
            _ = fi64
        }
        v76 := GenS1{F0: (v5 * genFn0())}
        _ = v76
    }
    for fi77 := 0; fi77 < 14; fi77++ {
        if false {
            if false {
                v8 = 962
            } else {
                var v78 float64 = -2.25
                v79 := GenS4{F0: 501, F1: v5}
                _ = v78
                _ = v79
            }
            switch v9 {
            case "a":
                fmt.Println(v5, v47)
            case "foo bar":
                if (v47 <= 2) {
                    v9 = "foo bar"
                    fmt.Println(v2, v2)
                    v47 += v47
                }
            case "":
                switch (v9 + v9) {
                case "":
                    var v80 int32 = -1
                    _ = v80
                }
            }
        }
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi77
    }
    fmt.Println(v2, v8, v51)
    switch 128 {
    case 32768:
        for fi81 := 0; fi81 < 11; fi81++ {
            v82 := GenS5{F0: (v5 * genFn0()), F1: 0, F2: 65535}
            _ = fi81
            _ = v82
        }
    }
    var v83 float32 = (v2 / (1e10 * (v2 - 0.0)))
    if (940 >= 0) {
        switch 32767 {
        case 549:
            fmt.Println(v9, v9)
        }
    } else {
        switch v9 {
        case "hello":
            v5 = (v5 / (((v5 ^ (v5 * 130))) | 1))
        case "foo bar":
            var v84 uint32 = v47
            _ = v84
        default:
            fmt.Println(v47, v5)
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    fmt.Println("v51:", v51)
    fmt.Println("v8:", v8)
    fmt.Println("v83:", v83)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v48
    _ = v5
    _ = v51
    _ = v8
    _ = v83
    _ = v9
}
