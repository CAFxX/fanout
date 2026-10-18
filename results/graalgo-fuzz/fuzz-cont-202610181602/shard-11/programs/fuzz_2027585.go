package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32 }
type GenS2 struct { F0 int16; F1 string; F2 int64 }
type GenS3 struct { F0 uint8; F1 int32 }
type GenS4 struct { F0 uint8; F1 int8 }
type GenS5 struct { F0 int8; F1 uint; F2 int8 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 uint32; F1 int64; F2 int32 }
type GenS9 struct { F0 uint16; F1 int64 }
type GenS10 struct { F0 int32 }
type GenS11 struct { F0 uint32; F1 int64; F2 string }
type GenS12 struct { F0 uint32; F1 int }

func main() {
    var v1 string = "hello"
    var v2 uint = 4294967295
    var v3 int64 = 0
    v4 := make(map[string]bool)
    var v5 uint64 = 0
    v3 -= 9223372036854775807
    v5 = (^(v5))
    var v6 string = v1
    var v7 uint = (v2 | (v2 & (v2 - (v2 | (v2 | v2)))))
    var v8 uint64 = v5
    if true {
        var v9 int64 = (-(v3))
        v10 := GenS1{F0: (v6 + (v1 + "hello")), F1: 42}
        var v11 uint32 = 294
        _ = v10
        _ = v11
        _ = v9
    } else {
        if true {
            var v12 bool = false
            _ = v12
        } else {
            if (0 >= 65535) {
                v13 := GenS2{F0: -874, F1: v6, F2: (v3 ^ v3)}
                if ((!false) || true) {
                    for fi14 := 0; fi14 < 8; fi14++ {
                        if false {
                            v2 = (^(v2))
                            for fi15 := 0; fi15 < 15; fi15++ {
                                switch v1 {
                                case "\n\t":
                                    v6 = v6
                                case "a":
                                    v5 = v5
                                default:
                                    switch (^(fi14)) {
                                    case 516:
                                        var v16 int8 = 1
                                        _ = v16
                                    case 32767:
                                        v17 := GenS3{F0: 133, F1: 42}
                                        fmt.Println(v7)
                                        _ = v17
                                    }
                                }
                                switch (-(fi14)) {
                                case 127:
                                    fmt.Println(v5)
                                case -2147483648:
                                    if true {
                                        var v18 string = v6
                                        _ = v18
                                    } else {
                                        var v19 float64 = 0.0
                                        fmt.Println(v7, v5)
                                        _ = v19
                                    }
                                case 0:
                                    switch v6 {
                                    case "":
                                        fmt.Println(v7, fi14, v8)
                                    case "\n\t":
                                        if (!true) {
                                            var v20 string = v6
                                            _ = v20
                                        } else {
                                            switch int(v8) {
                                            case 911:
                                                v3 = (v3 ^ (v3 << (((v2 ^ v7)) & 63)))
                                            default:
                                                if ((1.5 > 0.1) || ((!false) && (v3 == v3))) {
                                                    v5 += (v5 & (v5 % ((v5) | 1)))
                                                    var v21 uint = 839
                                                    _ = v21
                                                } else {
                                                    if true {
                                                        var v22 uint16 = 1
                                                        var v23 float64 = ((1.5 / -2.25) + 0.0)
                                                        _ = v22
                                                        _ = v23
                                                    } else {
                                                        var v24 uint16 = 1
                                                        v24 += 2
                                                        if false {
                                                            switch (v6 + v1) {
                                                            case "foo bar":
                                                                if (fi15 == -1) {
                                                                    fmt.Println(v5, v24, v1)
                                                                    var v25 int8 = 100
                                                                    v4[("foo bar" + v1)] = (12 == 42)
                                                                    _ = v25
                                                                }
                                                            }
                                                        } else {
                                                            v2 += v2
                                                        }
                                                        _ = v24
                                                    }
                                                }
                                            }
                                            var v26 bool = false
                                            v6 = "foo bar"
                                            _ = v26
                                        }
                                    case "hello":
                                        var v27 int64 = (v3 & (^(v3)))
                                        _ = v27
                                    }
                                }
                                _ = fi15
                            }
                            if false {
                                v7 -= 1
                                fmt.Println(v5, v8, v7)
                            } else {
                                var v28 uint32 = 48
                                v6 = "a"
                                if (true || ((!true) && false)) {
                                    v3 = (v3 >> ((674) & 63))
                                    var v29 int16 = 0
                                    for fi30 := 0; fi30 < 11; fi30++ {
                                        var v31 int16 = v29
                                        v1 = v6
                                        _ = fi30
                                        _ = v31
                                    }
                                    _ = v29
                                }
                                _ = v28
                            }
                        } else {
                            var v32 uint8 = 2
                            _ = v32
                        }
                        switch 128 {
                        case -704:
                            switch v1 {
                            case "a":
                                if true {
                                    v3 -= (^(v3))
                                    v33 := GenS4{F0: 2, F1: 127}
                                    v8 -= (v5 * (v5 / ((18446744073709551615) | 1)))
                                    _ = v33
                                } else {
                                    v34 := GenS5{F0: 100, F1: v7, F2: 1}
                                    for fi35 := 0; fi35 < 18; fi35++ {
                                        fmt.Println(fi14, fi35, v6)
                                        var v36 float32 = (-2.25 - (1.5 * (((0.0 - 1e10) * -2.25) - ((0.1 - (0.0 - 0.0)) * 1.5))))
                                        _ = fi35
                                        _ = v36
                                    }
                                    _ = v34
                                }
                            case "\n\t":
                                var v37 string = (v6 + "")
                                _ = v37
                            default:
                                fmt.Println(v1, fi14, v7)
                            }
                        case -44:
                            v38 := GenS6{F0: 762}
                            _ = v38
                        case 128:
                            v39 := GenS3{F0: 12, F1: 42}
                            _ = v39
                        }
                        _ = fi14
                    }
                    var v40 uint64 = (v5 << (((v7 & v7)) & 63))
                    v2 += v2
                    _ = v40
                } else {
                    switch 0 {
                    case 128:
                        v41 := []int{-469, 128, 42, -439}
                        _ = v41
                    case 2147483647:
                        for fi42 := 0; fi42 < 11; fi42++ {
                            v5 -= 42
                            var v43 int32 = -1
                            var v44 uint = v7
                            _ = fi42
                            _ = v43
                            _ = v44
                        }
                    case 0:
                        for fi45 := 0; fi45 < 10; fi45++ {
                            var v46 int64 = int64(v7)
                            var v47 int8 = -128
                            _ = fi45
                            _ = v46
                            _ = v47
                        }
                    }
                }
                var v48 uint8 = 128
                _ = v13
                _ = v48
            } else {
                var v49 string = "a"
                if true {
                    if (true && true) {
                        v50 := []int{32768, 32767, -1}
                        v51 := GenS7{F0: v7}
                        for fi52 := 0; fi52 < 12; fi52++ {
                            if true {
                                var v53 int8 = -47
                                v8 = v8
                                v54 := GenS8{F0: 824, F1: (v3 / (((v3 - v3)) | 1)), F2: -2147483648}
                                _ = v53
                                _ = v54
                            }
                            _ = fi52
                        }
                        _ = v50
                        _ = v51
                    } else {
                        switch 0 {
                        case 4294967295:
                            var v55 float32 = 0.0
                            _ = v55
                        case 42:
                            v56 := GenS3{F0: 152, F1: 42}
                            _ = v56
                        default:
                            if false {
                                v57 := GenS9{F0: 740, F1: v3}
                                fmt.Println(v1, v6)
                                v58 := []int{42, -431}
                                _ = v57
                                _ = v58
                            }
                        }
                    }
                    switch 121 {
                    case 1:
                        for fi59 := 0; fi59 < 2; fi59++ {
                            if (1 >= 1) {
                                v1 = ((((v49 + ("" + v6)) + "") + v49) + (("a" + "hello") + (v1 + (("foo bar" + "a") + "\n\t"))))
                                v60 := GenS10{F0: -205}
                                var v61 float32 = (1e10 * (float32(0.1) * 1.5))
                                _ = v60
                                _ = v61
                            }
                            _ = fi59
                        }
                    default:
                        v62 := GenS5{F0: -63, F1: 399, F2: 117}
                        _ = v62
                    }
                } else {
                    v63 := GenS11{F0: 0, F1: v3, F2: ((v1 + "\n\t") + v1)}
                    if true {
                        v64 := GenS12{F0: uint32(v5), F1: 124}
                        _ = v64
                    }
                    _ = v63
                }
                _ = v49
            }
            for fi65 := 0; fi65 < 13; fi65++ {
                var v66 int = fi65
                v67 := []int{(fi65 + fi65), (fi65 % (((fi65 * fi65)) | 1)), fi65}
                for fi68 := 0; fi68 < 7; fi68++ {
                    if (65535 < 552) {
                        fmt.Println(v2, v3, fi68)
                    } else {
                        if (false || true) {
                            var v69 uint16 = 255
                            if true {
                                v70 := GenS9{F0: (v69 << ((v2) & 63)), F1: v3}
                                _ = v70
                            } else {
                                v2 = (v2 << ((v2) & 63))
                            }
                            _ = v69
                        } else {
                            var v71 int16 = -659
                            if (false || (!false)) {
                                v6 = v1
                                v72 := GenS11{F0: 2147483648, F1: v3, F2: ((v1 + (v6 + ("a" + (((v6 + (("hello" + (("foo bar" + v6) + (v1 + ((v6 + v6) + (("foo bar" + ("a" + v1)) + ""))))) + "hello")) + v1) + ((v6 + "") + v1))))) + ((v6 + v6) + v1))}
                                _ = v72
                            }
                            _ = v71
                        }
                    }
                    _ = fi68
                }
                _ = fi65
                _ = v66
                _ = v67
            }
        }
        v73 := []int{1}
        if (!false) {
            _ = len(v73)
        }
        _ = v73
    }
    switch (v1 + "hello") {
    case "":
        var v74 string = "\n\t"
        _ = v74
    case "hello":
        v75 := []int{128, 42, 1}
        _ = v75
    case "foo bar":
        var v76 string = (("\n\t" + ((v6 + (("\n\t" + ("" + v1)) + v1)) + (v6 + ""))) + v1)
        _ = v76
    default:
        v1 = v1
    }
    for fi77 := 0; fi77 < 12; fi77++ {
        delete(v4, ((v6 + ("\n\t" + "hello")) + ("" + v6)))
        _ = fi77
    }
    fmt.Println(v7)
    fmt.Println(v3, v7)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
