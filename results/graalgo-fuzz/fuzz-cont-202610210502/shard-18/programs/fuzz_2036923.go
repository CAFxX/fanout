package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint; F1 string; F2 int }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int32; F1 string }
type GenS6 struct { F0 int8; F1 int16; F2 string }
type GenS7 struct { F0 string; F1 uint32; F2 int16 }
type GenS8 struct { F0 int; F1 int16; F2 uint16 }
type GenS9 struct { F0 int16; F1 uint8 }
func genFn0(p0 int) int16 {
    if (760 == 65535) {
        v5 := []int{-267, 128, p0, p0}
        for fi6 := 0; fi6 < 16; fi6++ {
            v7 := GenS1{F0: -32768}
            var v8 uint8 = 128
            v9 := GenS1{F0: -1}
            _ = fi6
            _ = v7
            _ = v8
            _ = v9
        }
        _ = v5
    } else {
        v10 := []int{p0, (-(p0)), (p0 | (^(p0)))}
        p0 = 4294967295
        p0 -= p0
        _ = v10
    }
    var v11 int32 = -1
    fmt.Println(p0, p0)
    if true {
        v12 := GenS1{F0: 1000}
        _ = v12
    }
    v11 = v11
    _ = p0
    _ = v11
    return -1
}

func main() {
    v1 := []int{-9223372036854775808, 56, 42}
    var v2 uint32 = 2147483648
    var v3 uint64 = 0
    var v4 uint = 2
    if len(v1) > 0 {
        v1[0] = -2147483648
    }
    v13 := GenS2{F0: -2147483648}
    var v14 uint16 = 1
    switch "hello" {
    case "a":
        v1 = append(v1, 42)
    case "hello":
        var v15 string = ("foo bar" + "\n\t")
        _ = v15
    }
    for fi16 := 0; fi16 < 4; fi16++ {
        for fi17 := 0; fi17 < 1; fi17++ {
            switch (("a" + ("" + "")) + "a") {
            case "hello":
                v18 := GenS2{F0: 2147483647}
                _ = v18
            default:
                var v19 int = (fi17 * fi17)
                _ = v19
            }
            fmt.Println(v14, fi17)
            _ = fi17
        }
        for fi20 := 0; fi20 < 2; fi20++ {
            var v21 uint = uint(v3)
            switch "foo bar" {
            case "hello":
                var v22 uint32 = (v2 ^ (v2 + (v2 % ((2147483647) | 1))))
                _ = v22
            case "foo bar":
                v14 = v14
            case "":
                v3 -= uint64(v2)
            }
            _ = fi20
            _ = v21
        }
        var v23 uint8 = 1
        _ = fi16
        _ = v23
    }
    var v24 uint32 = 4294967295
    for fi25 := 0; fi25 < 18; fi25++ {
        var v26 int64 = 268
        if len(v1) > 0 {
            v1[0] = fi25
        }
        if false {
            var v27 bool = true
            _ = v27
        }
        _ = fi25
        _ = v26
    }
    v28 := GenS3{F0: (^(v4)), F1: "\n\t", F2: 127}
    for fi29 := 0; fi29 < 13; fi29++ {
        var v30 float32 = 1e10
        switch (fi29 + fi29) {
        case -243:
            var v31 uint64 = (^(v3))
            _ = v31
        case -763:
            if (true && true) {
                for fi32 := 0; fi32 < 12; fi32++ {
                    switch ("" + (("" + "\n\t") + "hello")) {
                    case "":
                        v33 := GenS4{F0: 32767}
                        var v34 int8 = 127
                        switch ("foo bar" + "\n\t") {
                        case "foo bar":
                            if (int8(v3) < (v34 & (v34 & v34))) {
                                switch "" {
                                case "\n\t":
                                    switch (("a" + "foo bar") + (("" + "foo bar") + "hello")) {
                                    case "":
                                        fmt.Println(v34, v14)
                                    }
                                case "":
                                    v2 += 713
                                }
                                var v35 uint64 = v3
                                if len(v1) > 0 {
                                    v1[0] = -940
                                }
                                _ = v35
                            } else {
                                var v36 int16 = int16(v4)
                                if (!(true && true)) {
                                    v4 += (^(v4))
                                    v36 += genFn0((fi29 | 32767))
                                    if len(v1) > 0 {
                                        v1[0] = fi29
                                    }
                                }
                                v37 := GenS5{F0: 441, F1: ("" + "")}
                                _ = v36
                                _ = v37
                            }
                        case "\n\t":
                            fmt.Println(v30, v4)
                            var v38 int8 = v34
                            var v39 int16 = -32768
                            _ = v38
                            _ = v39
                        default:
                            if false {
                                _ = len(v1)
                            } else {
                                v1 = append(v1, (-(fi32)))
                                if true {
                                    if true {
                                        var v40 uint8 = 241
                                        _ = v40
                                    }
                                }
                                if (((^(v4)) != v4) && (((("\n\t" + "hello") + "foo bar") + "a") > "hello")) {
                                    v41 := GenS4{F0: 1000}
                                    _ = v41
                                }
                            }
                        }
                        v42 := GenS2{F0: 0}
                        _ = v33
                        _ = v34
                        _ = v42
                    default:
                        if true {
                            fmt.Println(v14, v4, v14)
                        } else {
                            v3 = 714
                            if len(v1) > 0 {
                                v1[0] = 63
                            }
                            if false {
                                switch fi32 {
                                case 128:
                                    fmt.Println(fi32)
                                case -9223372036854775808:
                                    if len(v1) > 0 {
                                        v1[0] = fi32
                                    }
                                case 42:
                                    var v43 string = "hello"
                                    _ = v43
                                default:
                                    switch 0 {
                                    case 9223372036854775807:
                                        if true {
                                            v3 = v3
                                            if true {
                                                switch "hello" {
                                                case "":
                                                    v2 -= 65535
                                                case "hello":
                                                    if len(v1) > 0 {
                                                        v1[0] = fi32
                                                    }
                                                }
                                                var v44 float32 = v30
                                                _ = v44
                                            } else {
                                                v14 += 65535
                                                v1 = append(v1, (-(fi29)))
                                                _ = len(v1)
                                            }
                                            if (!true) {
                                                var v45 uint16 = uint16(v3)
                                                var v46 uint64 = v3
                                                _ = v45
                                                _ = v46
                                            }
                                        } else {
                                            v24 = 42
                                        }
                                    case -9223372036854775808:
                                        var v47 uint8 = 246
                                        _ = v47
                                    case 4294967295:
                                        fmt.Println(fi32, v4, v2)
                                    }
                                }
                                var v48 uint16 = (v14 - v14)
                                v24 = (v24 << ((v4) & 63))
                                _ = v48
                            }
                        }
                        if (((1 != 42) || true) || true) {
                            v14 = 0
                            var v49 uint32 = v24
                            if false {
                                var v50 float64 = 1.5
                                switch "foo bar" {
                                case "hello":
                                    v24 += v2
                                }
                                var v51 float32 = 1e10
                                _ = v50
                                _ = v51
                            }
                            _ = v49
                        }
                        v52 := GenS5{F0: -2147483648, F1: (("" + ("a" + "")) + "hello")}
                        _ = v52
                    }
                    _ = fi32
                }
                var v53 uint8 = 255
                v54 := GenS6{F0: 42, F1: 1, F2: "\n\t"}
                _ = v53
                _ = v54
            }
        default:
            var v55 int32 = 667
            _ = v55
        }
        _ = fi29
        _ = v30
    }
    switch 276 {
    case 128:
        var v56 int16 = 311
        _ = v56
    case 32768:
        var v57 string = ""
        _ = v57
    }
    v4 -= (v4 << (((v4 >> ((v4) & 63))) & 63))
    if (0.0 > 1e10) {
        v2 = 914
        var v58 int32 = 2147483647
        switch "foo bar" {
        case "a":
            if true {
                for fi59 := 0; fi59 < 19; fi59++ {
                    v60 := GenS4{F0: genFn0((-(fi59)))}
                    v61 := GenS6{F0: 127, F1: int16(v2), F2: ("hello" + "a")}
                    _ = fi59
                    _ = v60
                    _ = v61
                }
                v24 = v24
                if ((v24 + v2) <= v2) {
                    var v62 uint64 = v3
                    fmt.Println(v4, v14, v14)
                    for fi63 := 0; fi63 < 15; fi63++ {
                        if true {
                            if true {
                                _ = len(v1)
                                var v64 uint8 = 128
                                v65 := GenS7{F0: ("" + "\n\t"), F1: uint32(v4), F2: 32767}
                                _ = v64
                                _ = v65
                            }
                            var v66 uint32 = (v2 & v24)
                            if true {
                                var v67 int32 = (v58 / ((v58) | 1))
                                if (v14 <= (v14 | (^(v14)))) {
                                    if (v14 != v14) {
                                        var v68 uint32 = v24
                                        _ = v68
                                    }
                                    var v69 int = (fi63 | (fi63 % (((fi63 * 0)) | 1)))
                                    _ = v69
                                } else {
                                    v70 := GenS3{F0: (v4 / ((v4) | 1)), F1: "foo bar", F2: 32767}
                                    _ = v70
                                }
                                _ = v67
                            } else {
                                v58 += -2147483648
                                fmt.Println(v62)
                                v4 = (^(v4))
                            }
                            _ = v66
                        }
                        _ = fi63
                    }
                    _ = v62
                } else {
                    v71 := GenS8{F0: 0, F1: 42, F2: v14}
                    switch 42 {
                    case 4294967295:
                        if ("" != ("hello" + "")) {
                            var v72 int8 = -128
                            var v73 uint64 = (v3 * v3)
                            switch 2147483647 {
                            case -9223372036854775808:
                                if false {
                                    v74 := GenS5{F0: 1, F1: ("a" + "a")}
                                    if false {
                                        v4 += 65535
                                    }
                                    _ = v74
                                }
                            case 0:
                                v58 += v58
                            default:
                                var v75 int16 = 295
                                _ = v75
                            }
                            _ = v72
                            _ = v73
                        }
                    }
                    v76 := GenS9{F0: 42, F1: 255}
                    _ = v71
                    _ = v76
                }
            }
        case "":
            if false {
                fmt.Println(v24, v3, v14)
            } else {
                var v77 float64 = 3.14159
                _ = v77
            }
        case "foo bar":
            v78 := GenS7{F0: "\n\t", F1: 0, F2: 0}
            _ = v78
        default:
            fmt.Println(v3, v24)
        }
        _ = v58
    } else {
        v2 -= uint32(v14)
    }
    v79 := GenS1{F0: 42}
    v2 += uint32(v3)
    v80 := GenS1{F0: -504}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v24
    _ = v28
    _ = v3
    _ = v4
    _ = v79
    _ = v80
}
