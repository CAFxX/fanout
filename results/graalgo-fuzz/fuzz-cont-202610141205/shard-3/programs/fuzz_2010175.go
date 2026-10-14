package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8; F2 int16 }
type GenS2 struct { F0 int16; F1 uint64; F2 int64 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 bool; F1 uint8 }
type GenS6 struct { F0 int8; F1 uint32; F2 uint32 }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 string }
type GenS9 struct { F0 bool; F1 string; F2 int32 }
type GenS10 struct { F0 string; F1 uint32; F2 uint32 }
func genFn0(p0 uint32) uint8 {
    for fi7 := 0; fi7 < 11; fi7++ {
        v8 := GenS1{F0: (fi7 ^ -1), F1: 0, F2: int16(fi7)}
        for fi9 := 0; fi9 < 11; fi9++ {
            fmt.Println(p0, fi9, p0)
            switch (fi7 << ((886) & 63)) {
            case 42:
                v10 := GenS2{F0: 0, F1: 490, F2: 477}
                _ = v10
            case -9223372036854775808:
                var v11 uint32 = (p0 % ((p0) | 1))
                _ = v11
            case 595:
                if ((-735 > 0) || true) {
                    if false {
                        v12 := GenS3{F0: 0}
                        v13 := []int{(-(fi9)), -2147483648, fi9, (fi7 * fi9)}
                        _ = v12
                        _ = v13
                    } else {
                        v14 := []int{(fi9 + fi7)}
                        v14 = append(v14, fi9)
                        v15 := GenS4{F0: 1}
                        _ = v14
                        _ = v15
                    }
                    var v16 uint32 = p0
                    v17 := GenS5{F0: ((p0 != (v16 ^ (^(p0)))) && (false || (("" < "a") || true))), F1: 2}
                    _ = v16
                    _ = v17
                }
            default:
                v18 := GenS3{F0: -687}
                _ = v18
            }
            _ = fi9
        }
        var v19 uint16 = 65535
        _ = fi7
        _ = v19
        _ = v8
    }
    var v20 int16 = -32768
    if (127 >= 2147483647) {
        var v21 int64 = 0
        _ = v21
    } else {
        fmt.Println(p0, v20)
        v20 += (v20 & 1000)
    }
    _ = p0
    _ = v20
    return 42
}
func genFn1(p0 bool, p1 bool, p2 string) int {
    switch 0 {
    case 127:
        v22 := []int{1, -609}
        _ = v22
    }
    if ((-2147483648 < 2147483647) || (!p1)) {
        var v23 float64 = (1e10 - (1.5 + float64(1.5)))
        _ = v23
    } else {
        v24 := GenS6{F0: 100, F1: 1, F2: 2147483648}
        switch p2 {
        case "a":
            var v25 uint = 0
            _ = v25
        case "":
            var v26 uint64 = 42
            _ = v26
        case "\n\t":
            v27 := GenS7{F0: -1}
            _ = v27
        }
        var v28 int32 = 0
        _ = v24
        _ = v28
    }
    p1 = ((!(272 > 884)) && p0)
    _ = p0
    _ = p1
    _ = p2
    return 32768
}

func main() {
    var v1 uint = 1
    var v2 int = 4294967295
    v3 := make(map[int]string)
    v4 := make(map[string]string)
    v5 := []uint32{4294967295, 42}
    v6 := []uint32{92, 501, 71}
    for fi29 := 0; fi29 < 15; fi29++ {
        fmt.Println(v2, v2, v2)
        v2 = (v2 >> (((v1 ^ (v1 % ((v1) | 1)))) & 63))
        _ = fi29
    }
    for fi30 := 0; fi30 < 8; fi30++ {
        switch "hello" {
        case "foo bar":
            v2 = (^(v2))
        default:
            if (-2147483648 == 66) {
                v6 = append(v6, 212)
                for fi31 := 0; fi31 < 20; fi31++ {
                    v1 += 18446744073709551615
                    _ = fi31
                }
                v32 := GenS4{F0: 79}
                _ = v32
            } else {
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                if true {
                    for fi33 := 0; fi33 < 15; fi33++ {
                        var v34 uint16 = 42
                        v2 = v2
                        var v35 float32 = 1e10
                        _ = fi33
                        _ = v34
                        _ = v35
                    }
                    for fi36 := 0; fi36 < 2; fi36++ {
                        var v37 float64 = float64(1.5)
                        if (65535 > 844) {
                            var v38 int32 = 0
                            v39 := v3[(fi30 >> ((220) & 63))]
                            _ = v38
                            _ = v39
                        } else {
                            v40 := GenS3{F0: -1}
                            v2 -= -2147483648
                            _ = v40
                        }
                        var v41 uint64 = 579
                        _ = fi36
                        _ = v37
                        _ = v41
                    }
                    fmt.Println(v1)
                } else {
                    if ((42 < 127) || true) {
                        v42 := GenS8{F0: "a"}
                        v1 -= 65535
                        _ = v42
                    }
                    var v43 int32 = 387
                    if true {
                        v2 = (v2 << (((v1 - (^(v1)))) & 63))
                        var v44 float32 = (((-2.25 * 1.5) - 0.0) + 1.5)
                        _ = v44
                    }
                    _ = v43
                }
                if (255 >= 296) {
                    for fi45 := 0; fi45 < 1; fi45++ {
                        if true {
                            if len(v5) > 0 {
                                v5[0] = 2147483647
                            }
                            if (((1.5 + 1.5) / 3.14159) < (1.5 - ((0.1 - 1e10) - 1.5))) {
                                var v46 int64 = 9223372036854775807
                                var v47 uint = v1
                                var v48 float64 = (1.5 - float64((((1e10 + 3.14159) + (float32((((1.5 / -2.25) / 3.14159) + 0.0)) - 0.1)) + -2.25)))
                                _ = v46
                                _ = v47
                                _ = v48
                            }
                            var v49 int = (fi30 >> ((255) & 63))
                            _ = v49
                        } else {
                            fmt.Println(v2, v2)
                            v2 = (fi45 + (v2 ^ -1))
                        }
                        var v50 string = "\n\t"
                        _ = fi45
                        _ = v50
                    }
                    for fi51 := 0; fi51 < 19; fi51++ {
                        _ = len(v3)
                        var v52 int64 = 9223372036854775807
                        if true {
                            v53 := GenS4{F0: 42}
                            var v54 int = v2
                            var v55 float64 = (-2.25 - (0.1 + -2.25))
                            _ = v53
                            _ = v54
                            _ = v55
                        }
                        _ = fi51
                        _ = v52
                    }
                }
            }
        }
        _ = fi30
    }
    v2 += (v2 << (((v1 << ((v1) & 63))) & 63))
    var v56 uint = v1
    var v57 int32 = -2147483648
    for fi58 := 0; fi58 < 15; fi58++ {
        for fi59 := 0; fi59 < 9; fi59++ {
            _ = len(v5)
            var v60 int16 = 731
            _ = fi59
            _ = v60
        }
        _ = fi58
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    for fi61 := 0; fi61 < 20; fi61++ {
        switch "\n\t" {
        case "hello":
            v2 -= v2
        case "a":
            fmt.Println(v57, v56)
        case "foo bar":
            switch "a" {
            case "foo bar":
                v1 += 18446744073709551615
            case "\n\t":
                if len(v6) > 0 {
                    v6[0] = 48
                }
            case "hello":
                var v62 float32 = -2.25
                _ = v62
            default:
                switch "" {
                case "":
                    if false {
                        fmt.Println(v56, v2)
                        if true {
                            var v63 uint16 = 255
                            _ = v63
                        }
                    } else {
                        switch (("a" + "a") + (("hello" + "a") + "foo bar")) {
                        case "foo bar":
                            v3[v2] = "a"
                        default:
                            var v64 float64 = 1e10
                            _ = v64
                        }
                        if (-2.25 < (1e10 * 3.14159)) {
                            v65 := GenS8{F0: "\n\t"}
                            for fi66 := 0; fi66 < 16; fi66++ {
                                v1 = 42
                                var v67 int64 = 1
                                _ = fi66
                                _ = v67
                            }
                            v56 = (v56 ^ 1)
                            _ = v65
                        } else {
                            switch "hello" {
                            case "a":
                                for fi68 := 0; fi68 < 20; fi68++ {
                                    v69 := GenS6{F0: 127, F1: 0, F2: 189}
                                    _ = fi68
                                    _ = v69
                                }
                            case "":
                                v70 := GenS7{F0: -681}
                                _ = v70
                            case "foo bar":
                                for fi71 := 0; fi71 < 2; fi71++ {
                                    var v72 bool = (128 > 89)
                                    v57 += (v57 - v57)
                                    _ = fi71
                                    _ = v72
                                }
                            }
                        }
                    }
                case "foo bar":
                    if false {
                        v73 := GenS9{F0: ((v2 & (fi61 >> (((v1 + (v56 ^ v56))) & 63))) < 127), F1: "foo bar", F2: v57}
                        _ = v73
                    } else {
                        var v74 bool = true
                        var v75 uint64 = 42
                        _ = v74
                        _ = v75
                    }
                case "a":
                    var v76 float32 = (3.14159 / 1e10)
                    _ = v76
                }
            }
        }
        var v77 uint = v1
        _ = fi61
        _ = v77
    }
    var v78 float64 = 1.5
    switch -9223372036854775808 {
    case 260:
        fmt.Println(v1, v1)
    case 128:
        v78 += (v78 / (v78 / 1.5))
    case 609:
        if false {
            v2 += v2
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
        } else {
            v2 += (-(v2))
        }
    }
    v57 += v57
    for fi79 := 0; fi79 < 2; fi79++ {
        v57 = -2147483648
        v56 = v56
        if (!false) {
            fmt.Println(v78, v57, v56)
            if (42 <= 243) {
                var v80 int = (fi79 * (fi79 | fi79))
                v81 := GenS10{F0: ("\n\t" + "foo bar"), F1: uint32(v1), F2: 1}
                _ = v80
                _ = v81
            }
            v56 = (v1 ^ (^(v56)))
        } else {
            var v82 float64 = v78
            _ = v82
        }
        _ = fi79
    }
    switch "a" {
    case "\n\t":
        v57 -= (v57 << ((v56) & 63))
    case "foo bar":
        _ = len(v4)
    default:
        fmt.Println(v1, v56, v78)
    }
    var v83 int16 = 1000
    var v84 uint32 = 4294967295
    var v85 int64 = 640
    var v86 int8 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v56:", v56)
    fmt.Println("v57:", v57)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v78:", v78)
    fmt.Println("v83:", v83)
    fmt.Println("v84:", v84)
    fmt.Println("v85:", v85)
    fmt.Println("v86:", v86)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v56
    _ = v57
    _ = v6
    _ = v78
    _ = v83
    _ = v84
    _ = v85
    _ = v86
}
