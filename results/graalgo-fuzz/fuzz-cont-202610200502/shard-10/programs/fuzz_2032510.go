package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int; F1 uint16; F2 uint }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 string; F1 int32 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 string; F1 int64; F2 int64 }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 int64 }
type GenS9 struct { F0 bool }
type GenS10 struct { F0 uint8 }
type GenS11 struct { F0 string }
type GenS12 struct { F0 uint32 }
type GenS13 struct { F0 uint64; F1 int16; F2 int64 }
type GenS14 struct { F0 uint; F1 uint8 }
func genFn0(p0 bool) int32 {
    for fi7 := 0; fi7 < 16; fi7++ {
        var v8 uint64 = 18446744073709551615
        v9 := GenS1{F0: 73}
        _ = fi7
        _ = v8
        _ = v9
    }
    var v10 bool = false
    _ = p0
    _ = v10
    return 1
}

func main() {
    var v1 uint8 = 0
    v2 := []uint32{2, 42, 1}
    var v3 float32 = 0.0
    v4 := make(map[int]string)
    var v5 float32 = v3
    var v6 int64 = -354
    for fi11 := 0; fi11 < 10; fi11++ {
        for fi12 := 0; fi12 < 1; fi12++ {
            if ((v5 - float32(1.5)) < v5) {
                v13 := GenS1{F0: v1}
                var v14 uint64 = 442
                _ = v13
                _ = v14
            } else {
                v6 += v6
                var v15 string = "hello"
                if len(v2) > 0 {
                    v2[0] = 0
                }
                v15 = (((v15 + v15) + ((v15 + "") + v15)) + v15)
                if (false && false) {
                    var v16 float64 = 0.1
                    _ = v16
                } else {
                    if len(v2) > 0 {
                        v2[0] = 4294967295
                    }
                    var v17 string = v15
                    _ = v17
                }
                _ = v15
            }
            if len(v2) > 0 {
                v2[0] = 2147483647
            }
            switch -464 {
            case -1:
                switch "" {
                case "":
                    var v18 float32 = v3
                    _ = v18
                case "foo bar":
                    var v19 float64 = 0.0
                    _ = v19
                case "a":
                    v2 = append(v2, 0)
                }
            case -970:
                if (!false) {
                    v2 = append(v2, 2147483648)
                } else {
                    var v20 uint16 = 2
                    v21 := v4[-9223372036854775808]
                    v1 = (^(v1))
                    _ = v20
                    _ = v21
                }
            case 32768:
                v3 = v5
            }
            v22 := GenS2{F0: 9223372036854775807, F1: 0, F2: 826}
            v23 := GenS3{F0: (v1 % (((v1 << ((18446744073709551615) & 63))) | 1))}
            v24 := GenS4{F0: "hello", F1: 11}
            if (32767 < -1) {
                v1 = (v1 + (v1 << ((831) & 63)))
                var v25 int8 = -128
                _ = v25
            } else {
                if false {
                    var v26 float32 = v3
                    var v27 bool = false
                    if (32767 < -1) {
                        fmt.Println(v1)
                        v28 := GenS3{F0: v1}
                        _ = v28
                    } else {
                        var v29 uint = 484
                        _ = v29
                    }
                    _ = v26
                    _ = v27
                } else {
                    v5 = 1.5
                    v5 = v3
                }
                if true {
                    v6 = (v6 << ((156) & 63))
                }
                switch "foo bar" {
                case "\n\t":
                    switch "foo bar" {
                    case "\n\t":
                        var v30 uint = 4294967295
                        _ = v30
                    case "hello":
                        v5 -= (1e10 - v5)
                    case "foo bar":
                        var v31 uint32 = 2147483647
                        _ = v31
                    default:
                        v3 = v5
                    }
                case "":
                    v32 := GenS5{F0: 42}
                    _ = v32
                default:
                    switch "foo bar" {
                    case "":
                        var v33 int = (-(fi11))
                        _ = v33
                    }
                }
                if false {
                    if (!false) {
                        var v34 int8 = -1
                        if true {
                            v34 = (v34 & 100)
                        } else {
                            var v35 float32 = v5
                            v2 = append(v2, 4294967295)
                            _ = v35
                        }
                        _ = v34
                    } else {
                        v1 += v1
                    }
                } else {
                    var v36 uint = 409
                    v37 := GenS6{F0: "", F1: (-(v6)), F2: v6}
                    _ = v36
                    _ = v37
                }
            }
            switch "" {
            case "\n\t":
                var v38 int = fi12
                _ = v38
            default:
                v39 := GenS7{F0: genFn0(((!(-2147483648 != 169)) || true))}
                _ = v39
            }
            _ = fi12
            _ = v22
            _ = v23
            _ = v24
        }
        for fi40 := 0; fi40 < 9; fi40++ {
            v41 := GenS3{F0: 2}
            if (float64(1e10) >= float64(v5)) {
                var v42 int32 = int32(fi40)
                _ = v42
            } else {
                if (!true) {
                    if (("\n\t" + "\n\t") <= (("foo bar" + ("" + "a")) + ("hello" + "foo bar"))) {
                        v3 -= 1.5
                        v6 = (v6 & 1)
                        if len(v2) > 0 {
                            v2[0] = 2
                        }
                    } else {
                        var v43 uint32 = 1
                        if (!true) {
                            var v44 int16 = -212
                            v4[-83] = "\n\t"
                            _ = v44
                        } else {
                            v1 = (v1 % ((128) | 1))
                            if ((float64(v3) > ((1.5 / 0.1) / -2.25)) && true) {
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                switch (("hello" + (("\n\t" + (("foo bar" + ("" + "hello")) + "a")) + ("hello" + ""))) + ((("\n\t" + ((("\n\t" + "hello") + ("foo bar" + ("\n\t" + "a"))) + "")) + "foo bar") + "\n\t")) {
                                case "hello":
                                    switch (fi40 - (fi40 & fi40)) {
                                    case 4294967295:
                                        fmt.Println(v43, v6)
                                    case 128:
                                        var v45 uint16 = 0
                                        _ = v45
                                    default:
                                        fmt.Println(v1, v43, v43)
                                        if (!false) {
                                            v6 = (v6 / (((^(v6))) | 1))
                                            switch "hello" {
                                            case "":
                                                v1 -= 0
                                            case "foo bar":
                                                var v46 int = 1
                                                _ = v46
                                            case "\n\t":
                                                var v47 float64 = ((1.5 + 0.0) - float64(3.14159))
                                                _ = v47
                                            default:
                                                delete(v4, (fi40 & (fi11 << ((uint(v1)) & 63))))
                                            }
                                            var v48 uint32 = 1
                                            _ = v48
                                        }
                                    }
                                    var v49 float32 = v3
                                    _ = v49
                                }
                            } else {
                                v43 = 487
                                var v50 uint8 = (v1 + v1)
                                v2 = append(v2, (^(v43)))
                                _ = v50
                            }
                        }
                        _ = v43
                    }
                    var v51 uint32 = 667
                    delete(v4, (fi11 | (fi40 ^ (fi40 << ((4294967295) & 63)))))
                    _ = v51
                }
                v52 := GenS5{F0: 865}
                switch (fi40 * fi11) {
                case 127:
                    delete(v4, 9223372036854775807)
                default:
                    v1 = 127
                    v5 = 1e10
                    if true {
                        fmt.Println(v1, fi11)
                        v3 -= (0.0 - v5)
                        fmt.Println(v3)
                    } else {
                        var v53 string = ("a" + (("hello" + "a") + ("foo bar" + "foo bar")))
                        if false {
                            var v54 float32 = ((v3 - 0.0) * (0.1 - (1.5 * 1e10)))
                            var v55 uint64 = 1
                            v5 = (1e10 + v5)
                            _ = v54
                            _ = v55
                        }
                        var v56 uint32 = 4294967295
                        _ = v53
                        _ = v56
                    }
                }
                _ = v52
            }
            _ = fi40
            _ = v41
        }
        v57 := GenS8{F0: (v6 - (v6 % ((v6) | 1)))}
        _ = fi11
        _ = v57
    }
    var v58 uint16 = 2
    v59 := GenS9{F0: true}
    v60 := GenS5{F0: 18446744073709551615}
    if (0 >= 1) {
        if false {
            v61 := GenS7{F0: int32(v6)}
            _ = v61
        }
        v62 := GenS10{F0: 255}
        delete(v4, 32768)
        _ = v62
    }
    var v63 int = 9223372036854775807
    var v64 int8 = 0
    switch -436 {
    case 127:
        switch (("\n\t" + "hello") + "") {
        case "":
            if false {
                switch v63 {
                case 127:
                    for fi65 := 0; fi65 < 8; fi65++ {
                        var v66 uint = 18446744073709551615
                        switch (fi65 % ((v63) | 1)) {
                        case -839:
                            if len(v2) > 0 {
                                v2[0] = 2147483648
                            }
                        case 46:
                            var v67 int64 = (v6 * (v6 | -43))
                            _ = v67
                        case -9223372036854775808:
                            switch v63 {
                            case 1:
                                for fi68 := 0; fi68 < 2; fi68++ {
                                    var v69 int16 = 351
                                    _ = fi68
                                    _ = v69
                                }
                            case 110:
                                _ = len(v2)
                            }
                        default:
                            if false {
                                if ("hello" == "hello") {
                                    fmt.Println(v63)
                                    if ((v6 + v6) > v6) {
                                        _ = len(v2)
                                    } else {
                                        if false {
                                            if (!((1 > -941) || false)) {
                                                fmt.Println(v3, v6)
                                                fmt.Println(v6, v63, v6)
                                            } else {
                                                var v70 int8 = v64
                                                v71 := GenS11{F0: ""}
                                                switch "foo bar" {
                                                case "hello":
                                                    var v72 uint64 = 0
                                                    _ = v72
                                                case "a":
                                                    var v73 uint8 = 87
                                                    _ = v73
                                                case "\n\t":
                                                    v74 := GenS8{F0: (v6 >> (((v66 ^ v66)) & 63))}
                                                    _ = v74
                                                default:
                                                    switch "hello" {
                                                    case "foo bar":
                                                        var v75 uint8 = (^(v1))
                                                        _ = v75
                                                    case "a":
                                                        fmt.Println(v70, v6, v58)
                                                    case "\n\t":
                                                        for fi76 := 0; fi76 < 8; fi76++ {
                                                            v6 -= -9223372036854775808
                                                            v77 := v4[92]
                                                            _ = fi76
                                                            _ = v77
                                                        }
                                                    }
                                                }
                                                _ = v70
                                                _ = v71
                                            }
                                            v1 -= 0
                                            for fi78 := 0; fi78 < 5; fi78++ {
                                                var v79 float64 = 1.5
                                                v80 := GenS12{F0: uint32(v58)}
                                                _ = fi78
                                                _ = v79
                                                _ = v80
                                            }
                                        } else {
                                            v6 += (^(v6))
                                            var v81 uint64 = 42
                                            _ = v81
                                        }
                                    }
                                    var v82 uint64 = 4294967295
                                    _ = v82
                                }
                                for fi83 := 0; fi83 < 13; fi83++ {
                                    fmt.Println(v6, fi65, v3)
                                    var v84 uint = (v66 % ((0) | 1))
                                    _ = fi83
                                    _ = v84
                                }
                                for fi85 := 0; fi85 < 20; fi85++ {
                                    v4[(v63 & (fi65 >> ((42) & 63)))] = "a"
                                    var v86 int8 = v64
                                    if (1e10 != float64(v3)) {
                                        v66 -= uint(v58)
                                        var v87 float32 = -2.25
                                        if (!true) {
                                            switch 410 {
                                            case -1:
                                                var v88 float32 = float32((((0.1 - ((0.0 / 3.14159) * 0.1)) * 1e10) + (-2.25 - 3.14159)))
                                                _ = v88
                                            case 992:
                                                v89 := v4[(-(fi85))]
                                                _ = v89
                                            default:
                                                v90 := GenS7{F0: 923}
                                                _ = v90
                                            }
                                            v66 += (v66 >> (((v66 ^ v66)) & 63))
                                        }
                                        _ = v87
                                    } else {
                                        if (v1 > v1) {
                                            if ((v1 >= v1) && true) {
                                                fmt.Println(fi65, v3)
                                                switch -1 {
                                                case -466:
                                                    fmt.Println(fi85)
                                                case 42:
                                                    switch "" {
                                                    case "hello":
                                                        var v91 int16 = 42
                                                        _ = v91
                                                    case "":
                                                        if len(v2) > 0 {
                                                            v2[0] = 4294967295
                                                        }
                                                    case "\n\t":
                                                        v58 = 256
                                                    }
                                                case -669:
                                                    switch ("a" + "a") {
                                                    case "a":
                                                        var v92 float32 = (1.5 - v5)
                                                        _ = v92
                                                    default:
                                                        v6 -= (v6 << ((v66) & 63))
                                                    }
                                                default:
                                                    if (1 < int32(v1)) {
                                                        fmt.Println(v3, v3, v5)
                                                    }
                                                }
                                                v3 += ((1.5 - (-2.25 / (v3 - v3))) / v5)
                                            } else {
                                                if false {
                                                    fmt.Println(v63, v66, v63)
                                                } else {
                                                    var v93 string = "\n\t"
                                                    fmt.Println(v63, v1)
                                                    _ = v93
                                                }
                                            }
                                            switch (-(fi85)) {
                                            case -418:
                                                v6 -= (-(v6))
                                            }
                                            _ = len(v4)
                                        }
                                    }
                                    _ = fi85
                                    _ = v86
                                }
                            }
                        }
                        _ = fi65
                        _ = v66
                    }
                case -9223372036854775808:
                    var v94 uint = 143
                    _ = v94
                case 717:
                    for fi95 := 0; fi95 < 9; fi95++ {
                        v58 += 1
                        _ = fi95
                    }
                }
            }
        case "hello":
            v5 = v5
        }
    default:
        v58 += (v58 * v58)
    }
    for fi96 := 0; fi96 < 15; fi96++ {
        fmt.Println(v3)
        fmt.Println(v64)
        v1 = 0
        _ = fi96
    }
    v97 := GenS11{F0: ((("" + "hello") + "foo bar") + ("foo bar" + ""))}
    fmt.Println(v64)
    var v98 int32 = -2147483648
    if false {
        if true {
            v99 := GenS13{F0: 522, F1: 1, F2: -9223372036854775808}
            for fi100 := 0; fi100 < 3; fi100++ {
                if false {
                    var v101 int16 = -1
                    var v102 uint64 = 1
                    _ = v101
                    _ = v102
                }
                _ = fi100
            }
            _ = v99
        }
    } else {
        v103 := GenS11{F0: ((("hello" + "\n\t") + ("a" + "hello")) + "")}
        v104 := GenS14{F0: 65535, F1: 255}
        _ = v103
        _ = v104
    }
    if ((v1 / ((v1) | 1)) != (v1 ^ v1)) {
        for fi105 := 0; fi105 < 19; fi105++ {
            v4[(fi105 / ((fi105) | 1))] = "a"
            _ = fi105
        }
        for fi106 := 0; fi106 < 15; fi106++ {
            switch "hello" {
            case "a":
                v4[int(v98)] = ""
            case "foo bar":
                v58 = (v58 << ((18446744073709551615) & 63))
            default:
                var v107 int16 = 1000
                _ = v107
            }
            _ = fi106
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v58:", v58)
    fmt.Println("v6:", v6)
    fmt.Println("v63:", v63)
    fmt.Println("v64:", v64)
    fmt.Println("v98:", v98)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v58
    _ = v59
    _ = v6
    _ = v60
    _ = v63
    _ = v64
    _ = v97
    _ = v98
}
