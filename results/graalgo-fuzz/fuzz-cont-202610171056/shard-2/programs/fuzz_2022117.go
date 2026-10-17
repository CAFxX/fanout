package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint; F2 uint }
type GenS2 struct { F0 int32; F1 string }
type GenS3 struct { F0 int64; F1 string }
type GenS4 struct { F0 uint8; F1 int32 }
type GenS5 struct { F0 int64; F1 int; F2 uint8 }
type GenS6 struct { F0 uint32; F1 int8 }
type GenS7 struct { F0 string; F1 uint16; F2 uint8 }
type GenS8 struct { F0 bool; F1 uint8; F2 uint8 }
func genFn0(p0 uint64) int {
    p0 -= (^(p0))
    if (2 > 1) {
        v4 := GenS1{F0: 1, F1: 42, F2: 31}
        p0 -= 474
        _ = v4
    } else {
        p0 = p0
    }
    for fi5 := 0; fi5 < 18; fi5++ {
        fmt.Println(p0)
        var v6 uint16 = 42
        _ = fi5
        _ = v6
    }
    v7 := []int{}
    var v8 uint8 = 2
    _ = p0
    _ = v7
    _ = v8
    return 664
}
func genFn1(p0 int16) uint8 {
    v9 := GenS1{F0: p0, F1: 18446744073709551615, F2: 947}
    var v10 uint8 = 15
    switch "" {
    case "foo bar":
        var v11 uint16 = 65535
        _ = v11
    case "":
        var v12 bool = (!true)
        _ = v12
    default:
        var v13 float64 = float64(0.1)
        _ = v13
    }
    var v14 int = 0
    if ((-2.25 > 0.1) || (false && true)) {
        if ((-1 > 359) && (((("a" + "hello") + "foo bar") == ("hello" + "")) || false)) {
            for fi15 := 0; fi15 < 18; fi15++ {
                if ((p0 >> ((371) & 63)) != (p0 & (p0 + (p0 & p0)))) {
                    v10 = 42
                    var v16 bool = false
                    _ = v16
                } else {
                    v10 += (v10 ^ v10)
                    var v17 bool = true
                    _ = v17
                }
                _ = fi15
            }
            v18 := GenS2{F0: 2147483647, F1: "\n\t"}
            _ = v18
        }
        var v19 uint16 = 390
        if (((p0 & 42) < (p0 << ((255) & 63))) && true) {
            var v20 int8 = -128
            var v21 uint8 = v10
            _ = v20
            _ = v21
        } else {
            fmt.Println(v14, v19)
        }
        _ = v19
    } else {
        switch "\n\t" {
        case "hello":
            v22 := GenS3{F0: 0, F1: (("foo bar" + "a") + "")}
            _ = v22
        case "a":
            var v23 uint16 = 42
            _ = v23
        case "foo bar":
            for fi24 := 0; fi24 < 6; fi24++ {
                if (624 > -2147483648) {
                    fmt.Println(v14)
                    p0 = (^(p0))
                    v14 = (fi24 / ((genFn0(283)) | 1))
                } else {
                    v25 := []int{(fi24 * 42), (v14 | fi24), 9223372036854775807, -775}
                    v26 := GenS2{F0: 0, F1: ((((("\n\t" + "\n\t") + (("a" + "\n\t") + "a")) + "") + ("" + ("a" + "hello"))) + ("\n\t" + ""))}
                    v10 += 1
                    _ = v25
                    _ = v26
                }
                var v27 int8 = 42
                var v28 uint = 255
                _ = fi24
                _ = v27
                _ = v28
            }
        }
        for fi29 := 0; fi29 < 1; fi29++ {
            var v30 uint64 = 4294967295
            var v31 int32 = 0
            _ = fi29
            _ = v30
            _ = v31
        }
        if (!((^(p0)) == 32767)) {
            if (false || false) {
                v14 = 128
                fmt.Println(v10, v14, v14)
                p0 -= -1
            } else {
                switch genFn0(18446744073709551615) {
                case -153:
                    v10 += 127
                case 127:
                    p0 -= 1000
                case -1:
                    var v32 int64 = -1
                    _ = v32
                default:
                    v10 = v10
                }
            }
            v14 = v14
        } else {
            v33 := []int{(v14 >> ((0) & 63))}
            switch ((("foo bar" + "") + "\n\t") + "hello") {
            case "foo bar":
                var v34 uint16 = 213
                _ = v34
            case "a":
                _ = len(v33)
            case "":
                if (-32768 == p0) {
                    var v35 uint16 = 32767
                    _ = v35
                }
            }
            _ = v33
        }
    }
    _ = p0
    _ = v10
    _ = v14
    _ = v9
    return v10
}

func main() {
    var v1 uint32 = 2147483647
    v2 := make(map[string]uint32)
    var v3 uint = 846
    v3 = (v3 * v3)
    for fi36 := 0; fi36 < 19; fi36++ {
        var v37 uint8 = 211
        _ = fi36
        _ = v37
    }
    v38 := GenS4{F0: 2, F1: 42}
    v1 = 555
    var v39 string = "a"
    if (-26 > -128) {
        v39 = (v39 + v39)
        switch ((v39 + v39) + "\n\t") {
        case "a":
            switch v39 {
            case "\n\t":
                if (true || true) {
                    for fi40 := 0; fi40 < 3; fi40++ {
                        var v41 uint64 = 1
                        var v42 int = (fi40 / ((-758) | 1))
                        var v43 int16 = -169
                        _ = fi40
                        _ = v41
                        _ = v42
                        _ = v43
                    }
                }
            case "hello":
                v44 := GenS4{F0: genFn1(1), F1: 1}
                _ = v44
            }
        }
        switch 1 {
        case 165:
            v1 = (v1 - v1)
        case 2147483647:
            for fi45 := 0; fi45 < 1; fi45++ {
                v46 := GenS5{F0: 9223372036854775807, F1: (fi45 - genFn0(4294967295)), F2: 255}
                if ((750 == 9223372036854775807) || (22 >= 0)) {
                    v47 := v2[v39]
                    _ = v47
                } else {
                    if (!true) {
                        if (int32(v3) != -1) {
                            var v48 uint16 = 33
                            if true {
                                v49 := v2[v39]
                                var v50 uint64 = 4294967295
                                switch v39 {
                                case "":
                                    v50 = 18446744073709551615
                                case "hello":
                                    for fi51 := 0; fi51 < 18; fi51++ {
                                        v52 := GenS6{F0: (v49 % ((v1) | 1)), F1: -1}
                                        _ = fi51
                                        _ = v52
                                    }
                                }
                                _ = v49
                                _ = v50
                            }
                            v53 := []int{fi45, fi45}
                            _ = v48
                            _ = v53
                        }
                    }
                }
                _ = fi45
                _ = v46
            }
        default:
            var v54 uint = 42
            _ = v54
        }
    } else {
        v1 = v1
    }
    if true {
        for fi55 := 0; fi55 < 19; fi55++ {
            if ((fi55 <= 128) || true) {
                var v56 uint64 = 2
                v1 = 1
                for fi57 := 0; fi57 < 6; fi57++ {
                    _ = len(v2)
                    _ = len(v2)
                    fmt.Println(v56, fi57)
                    _ = fi57
                }
                _ = v56
            } else {
                fmt.Println(v3, v1, fi55)
                fmt.Println(v3, v3)
                v58 := v2[v39]
                _ = v58
            }
            v3 = v3
            _ = fi55
        }
    }
    for fi59 := 0; fi59 < 5; fi59++ {
        var v60 int = fi59
        var v61 uint8 = 128
        _ = fi59
        _ = v60
        _ = v61
    }
    v39 = v39
    switch "foo bar" {
    case "hello":
        fmt.Println(v39)
    case "\n\t":
        if (42 != -60) {
            for fi62 := 0; fi62 < 8; fi62++ {
                v63 := []int{genFn0(4294967295), fi62, fi62, 128}
                var v64 float32 = 1e10
                _ = fi62
                _ = v63
                _ = v64
            }
            if true {
                delete(v2, (("" + v39) + ((v39 + "foo bar") + v39)))
            }
            v65 := []int{32768, -2147483648, -2147483648, 0}
            _ = v65
        } else {
            var v66 bool = (false && (784 != 1))
            for fi67 := 0; fi67 < 20; fi67++ {
                var v68 uint8 = genFn1(1000)
                for fi69 := 0; fi69 < 14; fi69++ {
                    switch genFn0(42) {
                    case -996:
                        var v70 float64 = 0.1
                        _ = v70
                    case -247:
                        var v71 float32 = 0.0
                        _ = v71
                    default:
                        if false {
                            if v66 {
                                v39 = v39
                                v2[v39] = 42
                            }
                            v39 = "\n\t"
                            if v66 {
                                var v72 uint = (v3 + (v3 >> ((93) & 63)))
                                var v73 int = fi69
                                _ = v72
                                _ = v73
                            }
                        }
                    }
                    if (((((v39 + v39) + "") + v39) != v39) && ((!true) || v66)) {
                        v68 = (v68 << ((18446744073709551615) & 63))
                        var v74 float32 = 3.14159
                        _ = v74
                    } else {
                        if ((!(((-2.25 - 3.14159) + 0.0) < 0.0)) && (!v66)) {
                            if ((v3 ^ 255) < (v3 << ((v3) & 63))) {
                                switch v39 {
                                case "foo bar":
                                    var v75 float32 = (((1.5 * (1e10 / 1.5)) - ((1.5 + 0.0) * 0.1)) + (-2.25 - -2.25))
                                    _ = v75
                                case "a":
                                    fmt.Println(fi69, v39, v39)
                                default:
                                    delete(v2, v39)
                                }
                                var v76 uint = 437
                                if (v66 && v66) {
                                    fmt.Println(v66, fi67, v3)
                                    var v77 int16 = 0
                                    v3 -= v76
                                    _ = v77
                                }
                                if v66 {
                                    delete(v2, (((("a" + v39) + v39) + v39) + "\n\t"))
                                    v78 := []int{(fi69 - int(v1))}
                                    _ = v78
                                } else {
                                    switch 0 {
                                    case -2147483648:
                                        if (862 <= v1) {
                                            if v66 {
                                                v68 = v68
                                                v66 = (!v66)
                                            } else {
                                                if v66 {
                                                    var v79 int16 = -750
                                                    _ = v79
                                                } else {
                                                    v80 := GenS5{F0: 0, F1: (fi69 - (fi67 ^ 0)), F2: (v68 >> (((v76 | v3)) & 63))}
                                                    var v81 int16 = 1000
                                                    v3 = 4294967295
                                                    _ = v80
                                                    _ = v81
                                                }
                                                v82 := GenS6{F0: uint32(v68), F1: 42}
                                                if v66 {
                                                    fmt.Println(v76, v76)
                                                } else {
                                                    v83 := GenS7{F0: v39, F1: 2, F2: 128}
                                                    _ = v83
                                                }
                                                _ = v82
                                            }
                                            v66 = (95 == 1)
                                        }
                                    }
                                    v1 += (^(v1))
                                }
                                v84 := []int{}
                                _ = v76
                                _ = v84
                            } else {
                                var v85 uint = v3
                                switch fi67 {
                                case 32767:
                                    v2[(((v39 + v39) + ((v39 + ((v39 + v39) + "a")) + v39)) + (v39 + (("hello" + ("\n\t" + v39)) + v39)))] = v1
                                case -1:
                                    fmt.Println(v39, v68, v3)
                                case 9223372036854775807:
                                    var v86 int16 = -293
                                    _ = v86
                                default:
                                    v3 = (v3 | (v85 + (v85 >> ((v3) & 63))))
                                }
                                switch v39 {
                                case "a":
                                    if false {
                                        v1 += v1
                                    } else {
                                        v85 -= v85
                                        fmt.Println(fi67)
                                    }
                                }
                                _ = v85
                            }
                            v87 := []int{(fi67 / ((679) | 1)), 32768, (fi69 ^ fi67), 4294967295}
                            _ = v87
                        } else {
                            var v88 uint32 = v1
                            if false {
                                v66 = (1 <= 127)
                                var v89 uint16 = 2
                                switch ("foo bar" + v39) {
                                case "":
                                    var v90 uint8 = (^(v68))
                                    _ = v90
                                case "foo bar":
                                    switch (-(fi67)) {
                                    case 32768:
                                        v91 := GenS6{F0: (v88 + v1), F1: -1}
                                        _ = v91
                                    }
                                case "\n\t":
                                    var v92 uint = uint(v89)
                                    _ = v92
                                default:
                                    if true {
                                        v93 := GenS2{F0: 1, F1: v39}
                                        fmt.Println(v39, fi69)
                                        _ = v93
                                    }
                                    if false {
                                        v68 = v68
                                        switch 177 {
                                        case -379:
                                            v39 = "\n\t"
                                        }
                                    } else {
                                        switch (v39 + v39) {
                                        case "":
                                            v88 += v1
                                        default:
                                            v39 = (("a" + v39) + v39)
                                        }
                                    }
                                    v1 = v1
                                }
                                _ = v89
                            }
                            _ = v88
                        }
                        if ((!(!v66)) || (v66 || (genFn0(558) >= 0))) {
                            delete(v2, "\n\t")
                            v94 := []int{-893}
                            _ = v94
                        } else {
                            if (0 > 42) {
                                if v66 {
                                    fmt.Println(fi69, fi67, v68)
                                    delete(v2, "")
                                } else {
                                    v1 = v1
                                }
                                v68 += v68
                                v95 := GenS8{F0: ((v3 >> ((219) & 63)) == 42), F1: v68, F2: 127}
                                _ = v95
                            }
                        }
                        var v96 int = fi67
                        _ = v96
                    }
                    _ = fi69
                }
                if v66 {
                    v66 = (-2.25 != 1e10)
                } else {
                    var v97 int32 = 1
                    var v98 int32 = 1
                    _ = v97
                    _ = v98
                }
                _ = fi67
                _ = v68
            }
            _ = v66
        }
    default:
        var v99 uint64 = 984
        _ = v99
    }
    var v100 int64 = 9223372036854775807
    fmt.Println("v1:", v1)
    fmt.Println("v100:", v100)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    _ = v1
    _ = v100
    _ = v2
    _ = v3
    _ = v38
    _ = v39
}
