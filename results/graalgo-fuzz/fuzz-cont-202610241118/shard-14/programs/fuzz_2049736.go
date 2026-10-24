package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint; F2 uint }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int32; F1 string; F2 uint32 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint; F1 uint64 }
type GenS6 struct { F0 uint8; F1 uint; F2 int32 }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 int64; F1 uint }
type GenS9 struct { F0 bool; F1 int32 }
type GenS10 struct { F0 uint16; F1 uint8; F2 int16 }
type GenS11 struct { F0 uint64; F1 uint32 }
func genFn0() string {
    var v4 float32 = 0.0
    v4 = 0.1
    v4 = (v4 * 1.5)
    for fi5 := 0; fi5 < 19; fi5++ {
        v6 := GenS1{F0: 1, F1: 609, F2: 4294967295}
        var v7 uint16 = 32767
        v8 := GenS2{F0: "foo bar"}
        _ = fi5
        _ = v6
        _ = v7
        _ = v8
    }
    v4 += (v4 * v4)
    _ = v4
    return "hello"
}
func genFn1(p0 int32, p1 int32) int32 {
    if false {
        if (1 < 112) {
            p0 = p1
            for fi9 := 0; fi9 < 12; fi9++ {
                if (false || false) {
                    p1 = p0
                } else {
                    v10 := GenS3{F0: p0, F1: (("a" + "") + "\n\t"), F2: 2}
                    _ = v10
                }
                _ = fi9
            }
            var v11 int = -2147483648
            _ = v11
        } else {
            var v12 int16 = -432
            _ = v12
        }
        p1 -= -776
    } else {
        v13 := []int{}
        fmt.Println(p0, p0)
        if (!true) {
            for fi14 := 0; fi14 < 9; fi14++ {
                for fi15 := 0; fi15 < 16; fi15++ {
                    p1 = -25
                    v16 := GenS4{F0: 18446744073709551615}
                    _ = fi15
                    _ = v16
                }
                _ = len(v13)
                _ = fi14
            }
            _ = len(v13)
        } else {
            if (0 > 1) {
                _ = len(v13)
                var v17 uint64 = 2
                if len(v13) > 0 {
                    v13[0] = 127
                }
                _ = v17
            } else {
                if false {
                    var v18 float32 = 0.0
                    _ = v18
                }
                fmt.Println(p0)
            }
            p1 += 42
            p1 = (p0 - p0)
        }
        _ = v13
    }
    var v19 float32 = -2.25
    var v20 int32 = 29
    v21 := []int{128, 316, 32767, 235}
    _ = p0
    _ = p1
    _ = v19
    _ = v20
    _ = v21
    return p0
}

func main() {
    v1 := make(map[string]uint32)
    var v2 int8 = 31
    var v3 uint8 = 0
    var v22 bool = false
    if (!v22) {
        v23 := GenS3{F0: 1, F1: "\n\t", F2: 65535}
        _ = v23
    } else {
        v24 := GenS1{F0: 0, F1: 42, F2: 65535}
        _ = v24
    }
    v3 = 1
    v22 = false
    var v25 uint32 = 42
    if (1 >= 42) {
        for fi26 := 0; fi26 < 14; fi26++ {
            var v27 uint32 = v25
            v28 := GenS3{F0: 42, F1: "foo bar", F2: v27}
            var v29 int = 456
            _ = fi26
            _ = v27
            _ = v28
            _ = v29
        }
        if v22 {
            if false {
                v3 -= (v3 ^ 128)
                if (!(1.5 < (1e10 * -2.25))) {
                    if v22 {
                        v30 := GenS1{F0: 2147483647, F1: 412, F2: 1}
                        v31 := []int{0, int(v25)}
                        v22 = true
                        _ = v30
                        _ = v31
                    } else {
                        var v32 int = -2147483648
                        _ = v32
                    }
                    for fi33 := 0; fi33 < 9; fi33++ {
                        var v34 uint = uint(v25)
                        _ = fi33
                        _ = v34
                    }
                } else {
                    var v35 int = 32768
                    _ = v35
                }
            }
            v2 += v2
            delete(v1, ("" + "\n\t"))
        }
        var v36 uint8 = v3
        _ = v36
    }
    _ = len(v1)
    if (-1 >= 42) {
        for fi37 := 0; fi37 < 13; fi37++ {
            var v38 bool = true
            switch "" {
            case "foo bar":
                v39 := v1[(("hello" + "") + ("" + "\n\t"))]
                _ = v39
            case "":
                v3 += v3
            case "\n\t":
                v3 = v3
            }
            _ = fi37
            _ = v38
        }
        var v40 int = 35
        switch "" {
        case "hello":
            if (!false) {
                v40 -= (v40 + (v40 / ((9223372036854775807) | 1)))
            } else {
                v41 := GenS1{F0: (v40 % ((v40) | 1)), F1: 1, F2: 2}
                var v42 string = "foo bar"
                var v43 int64 = -9223372036854775808
                _ = v41
                _ = v42
                _ = v43
            }
        case "\n\t":
            for fi44 := 0; fi44 < 1; fi44++ {
                switch 127 {
                case -9223372036854775808:
                    for fi45 := 0; fi45 < 19; fi45++ {
                        v46 := []int{(v40 | fi44)}
                        switch v40 {
                        case -826:
                            var v47 uint64 = 2
                            _ = v47
                        case -1:
                            var v48 int8 = (v2 / ((v2) | 1))
                            _ = v48
                        case 128:
                            if (!((!v22) && (!(v22 && (!v22))))) {
                                fmt.Println(v40, v25)
                            }
                        }
                        var v49 int8 = (v2 & 77)
                        var v50 uint32 = v25
                        var v51 int8 = (^(v2))
                        _ = fi45
                        _ = v46
                        _ = v49
                        _ = v50
                        _ = v51
                    }
                case 128:
                    v52 := GenS2{F0: "foo bar"}
                    _ = v52
                }
                v53 := GenS5{F0: 255, F1: 0}
                _ = fi44
                _ = v53
            }
        case "a":
            v3 = (v3 - (^(v3)))
        }
        _ = v40
    } else {
        if (42 <= genFn1(-2147483648, 42)) {
            v22 = false
            var v54 uint64 = 800
            _ = v54
        } else {
            var v55 uint = 4294967295
            v55 += 539
            var v56 bool = true
            _ = v55
            _ = v56
        }
        var v57 uint64 = 0
        var v58 int64 = 9223372036854775807
        _ = v57
        _ = v58
    }
    if (!(32767 >= -52)) {
        fmt.Println(v3)
        var v59 float32 = float32(float64(0.0))
        _ = v59
    }
    for fi60 := 0; fi60 < 11; fi60++ {
        if (v22 || (!v22)) {
            v61 := GenS6{F0: (v3 & v3), F1: 892, F2: 689}
            fmt.Println(v2, v3)
            _ = v61
        } else {
            v62 := GenS5{F0: 13, F1: 558}
            _ = v62
        }
        var v63 float64 = -2.25
        _ = fi60
        _ = v63
    }
    v64 := []int{}
    for fi65 := 0; fi65 < 16; fi65++ {
        var v66 uint64 = 667
        for fi67 := 0; fi67 < 14; fi67++ {
            if (((2 != 4294967295) || (4294967295 != (fi65 >> ((1) & 63)))) || (!v22)) {
                v1[(("\n\t" + "hello") + ("a" + "a"))] = (v25 ^ v25)
            }
            if v22 {
                switch "foo bar" {
                case "":
                    switch "a" {
                    case "foo bar":
                        var v68 int64 = -1
                        _ = v68
                    case "\n\t":
                        var v69 int = -100
                        _ = v69
                    default:
                        var v70 uint8 = v3
                        _ = v70
                    }
                }
                var v71 int64 = 42
                _ = v71
            }
            _ = fi67
        }
        var v72 int64 = 42
        _ = fi65
        _ = v66
        _ = v72
    }
    var v73 float64 = 3.14159
    v25 -= (v25 ^ 42)
    switch "hello" {
    case "":
        if (((-2.25 == 1e10) || true) || (0.1 >= 0.1)) {
            for fi74 := 0; fi74 < 16; fi74++ {
                for fi75 := 0; fi75 < 14; fi75++ {
                    var v76 float32 = -2.25
                    var v77 int32 = -2147483648
                    _ = len(v1)
                    _ = fi75
                    _ = v76
                    _ = v77
                }
                fmt.Println(v3, fi74)
                _ = fi74
            }
            if (4294967295 >= 4294967295) {
                if len(v64) >= 2 {
                    v64 = v64[1:len(v64)]
                }
            }
            v2 = (v2 / (((v2 | (v2 / ((-128) | 1)))) | 1))
        }
    case "\n\t":
        switch "\n\t" {
        case "hello":
            v78 := GenS7{F0: 0}
            _ = v78
        case "":
            if len(v64) >= 2 {
                v64 = v64[1:len(v64)]
            }
        }
    default:
        if v22 {
            v79 := GenS8{F0: -336, F1: 18446744073709551615}
            _ = v79
        } else {
            if true {
                switch 0 {
                case 433:
                    fmt.Println(v22, v25, v3)
                case 807:
                    v80 := GenS9{F0: false, F1: -432}
                    _ = v80
                case 4294967295:
                    fmt.Println(v2, v22, v3)
                default:
                    var v81 int = int(v2)
                    _ = v81
                }
                if v22 {
                    if len(v64) > 0 {
                        v64[0] = 32767
                    }
                    for fi82 := 0; fi82 < 18; fi82++ {
                        v83 := GenS7{F0: -128}
                        switch (((("" + "") + "a") + "hello") + "a") {
                        case "foo bar":
                            for fi84 := 0; fi84 < 7; fi84++ {
                                v73 = (v73 - (v73 / (3.14159 - 1e10)))
                                v85 := GenS8{F0: -414, F1: 42}
                                var v86 uint8 = v3
                                _ = fi84
                                _ = v85
                                _ = v86
                            }
                        case "a":
                            v87 := GenS10{F0: 255, F1: 255, F2: 504}
                            _ = v87
                        case "\n\t":
                            if (!v22) {
                                _ = len(v1)
                            }
                        }
                        switch 127 {
                        case -1:
                            var v88 int64 = -9223372036854775808
                            _ = v88
                        }
                        _ = fi82
                        _ = v83
                    }
                } else {
                    v2 = v2
                    if v22 {
                        for fi89 := 0; fi89 < 6; fi89++ {
                            var v90 float64 = (v73 / 1.5)
                            v91 := GenS6{F0: 0, F1: 255, F2: -246}
                            fmt.Println(fi89)
                            _ = fi89
                            _ = v90
                            _ = v91
                        }
                    }
                    v3 += v3
                }
                if false {
                    var v92 bool = false
                    v93 := GenS3{F0: 2147483647, F1: (((("a" + ("\n\t" + "foo bar")) + "") + ("hello" + "a")) + "a"), F2: v25}
                    var v94 bool = (-926 < -1)
                    _ = v92
                    _ = v93
                    _ = v94
                } else {
                    fmt.Println(v2)
                    if (!v22) {
                        switch "\n\t" {
                        case "foo bar":
                            fmt.Println(v25)
                        case "a":
                            if v22 {
                                v95 := v1[(("a" + "hello") + "foo bar")]
                                _ = v95
                            }
                        }
                        v22 = (((v73 + v73) - float64((0.1 / float32(float64((((0.1 + (float32(float64((float32(v73) + 0.1))) * (1e10 + 3.14159))) * 0.0) * 0.0)))))) > -2.25)
                    }
                    if (!(!true)) {
                        v2 = (v2 % (((-(v2))) | 1))
                        v3 += v3
                        v22 = true
                    }
                }
            } else {
                fmt.Println(v22)
                v96 := GenS11{F0: 4294967295, F1: 118}
                _ = v96
            }
            v73 += -2.25
            v97 := GenS5{F0: 1, F1: 18446744073709551615}
            _ = v97
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v64:", len(v64), v64)
    fmt.Println("v73:", v73)
    _ = v1
    _ = v2
    _ = v22
    _ = v25
    _ = v3
    _ = v64
    _ = v73
}
