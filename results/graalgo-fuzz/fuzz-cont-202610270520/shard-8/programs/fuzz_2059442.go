package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32; F2 uint }
type GenS2 struct { F0 int8; F1 int; F2 int8 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 uint64; F1 uint32 }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 uint64; F1 bool; F2 uint8 }
type GenS9 struct { F0 int16 }
type GenS10 struct { F0 int64 }
type GenS11 struct { F0 uint64; F1 int; F2 uint }
type GenS12 struct { F0 bool }
type GenS13 struct { F0 string; F1 int8; F2 uint }
func genFn0() int32 {
    var v6 uint8 = 223
    if (3.14159 < 1e10) {
        fmt.Println(v6)
    } else {
        switch 42 {
        case 1:
            for fi7 := 0; fi7 < 16; fi7++ {
                var v8 uint = 227
                _ = fi7
                _ = v8
            }
        }
        v6 -= v6
    }
    _ = v6
    return 42
}

func main() {
    var v1 string = "hello"
    v2 := []int{1, 0, -2147483648, 42}
    v3 := make(map[string]bool)
    v4 := make(map[uint32]uint32)
    v5 := []string{"\n\t", v1, "hello", v1}
    switch -2147483648 {
    case -266:
        if true {
            var v9 int = 323
            _ = v9
        }
    default:
        for fi10 := 0; fi10 < 15; fi10++ {
            v1 = "a"
            _ = fi10
        }
    }
    switch 2147483647 {
    case 128:
        if (0.1 >= (1.5 * 1.5)) {
            if (180 == 128) {
                v11 := GenS1{F0: "hello", F1: 813, F2: 18446744073709551615}
                _ = v11
            }
            var v12 int16 = 117
            v13 := GenS1{F0: "", F1: 4294967295, F2: 0}
            _ = v12
            _ = v13
        } else {
            v14 := GenS2{F0: 0, F1: 32767, F2: -63}
            var v15 int = 211
            _ = v14
            _ = v15
        }
    default:
        _ = len(v4)
    }
    for fi16 := 0; fi16 < 6; fi16++ {
        fmt.Println(v1, fi16)
        _ = fi16
    }
    _ = len(v3)
    if (false && (92 <= 1)) {
        var v17 uint = 679
        switch v1 {
        case "foo bar":
            v17 = v17
        case "\n\t":
            switch 127 {
            case 32767:
                v17 = (v17 / ((v17) | 1))
            case -662:
                v17 = 255
            case 128:
                if (-1 == -1) {
                    for fi18 := 0; fi18 < 17; fi18++ {
                        fmt.Println(v1)
                        for fi19 := 0; fi19 < 1; fi19++ {
                            _ = len(v2)
                            v17 -= 1
                            var v20 bool = false
                            _ = fi19
                            _ = v20
                        }
                        _ = fi18
                    }
                } else {
                    v1 = (v1 + v1)
                    v1 = v1
                }
            }
        case "hello":
            for fi21 := 0; fi21 < 17; fi21++ {
                if (65535 <= 32767) {
                    v22 := GenS2{F0: 127, F1: (fi21 % (((-(fi21))) | 1)), F2: 42}
                    switch fi21 {
                    case -9223372036854775808:
                        for fi23 := 0; fi23 < 2; fi23++ {
                            var v24 int = (fi21 | fi23)
                            _ = fi23
                            _ = v24
                        }
                    default:
                        var v25 string = (v1 + (("foo bar" + ((v1 + "hello") + (v1 + ("\n\t" + v1)))) + (v1 + v1)))
                        _ = v25
                    }
                    v2 = append(v2, fi21)
                    _ = v22
                } else {
                    if (false || (-615 >= 0)) {
                        v26 := GenS1{F0: "foo bar", F1: 2147483647, F2: v17}
                        v27 := GenS3{F0: -128}
                        v1 = ("\n\t" + ("" + v1))
                        _ = v26
                        _ = v27
                    } else {
                        v17 += v17
                        v28 := GenS4{F0: 100}
                        _ = v28
                    }
                    v17 -= (v17 % ((738) | 1))
                }
                if false {
                    v29 := GenS2{F0: -76, F1: (fi21 & (fi21 + 9223372036854775807)), F2: 127}
                    if len(v5) > 0 {
                        v5[0] = "hello"
                    }
                    _ = v29
                } else {
                    if len(v2) > 0 {
                        v2[0] = 2147483647
                    }
                }
                _ = fi21
            }
        default:
            v17 -= 574
        }
        _ = v17
    } else {
        delete(v4, 2147483648)
        if false {
            v1 = v1
            switch v1 {
            case "\n\t":
                v1 = v1
            default:
                var v30 int8 = 42
                _ = v30
            }
            var v31 uint64 = 18446744073709551615
            _ = v31
        }
    }
    v32 := GenS3{F0: -98}
    fmt.Println(v1)
    v33 := v4[2147483648]
    if ((!(-1 != -1)) && true) {
        var v34 int64 = 42
        v34 = (v34 ^ v34)
        for fi35 := 0; fi35 < 19; fi35++ {
            v36 := GenS2{F0: -1, F1: (fi35 * -751), F2: 42}
            for fi37 := 0; fi37 < 19; fi37++ {
                if ((0.1 * -2.25) >= ((0.0 + 1.5) * 1.5)) {
                    _ = len(v3)
                    var v38 int16 = 32767
                    var v39 uint32 = (v33 >> ((18446744073709551615) & 63))
                    _ = v38
                    _ = v39
                }
                _ = fi37
            }
            _ = fi35
            _ = v36
        }
        _ = v34
    } else {
        var v40 int64 = -1
        _ = v40
    }
    for fi41 := 0; fi41 < 18; fi41++ {
        v1 = v1
        if false {
            fmt.Println(v33)
            _ = len(v5)
        } else {
            v33 = 716
            if (1.5 == (1.5 * (0.0 + (-2.25 + 1.5)))) {
                v42 := GenS5{F0: 2147483647}
                fmt.Println(fi41, v33)
                _ = v42
            } else {
                if true {
                    v3[(v1 + "")] = (v1 <= ((v1 + v1) + ((v1 + (v1 + "hello")) + v1)))
                    var v43 int = 32767
                    if false {
                        if (3.14159 > -2.25) {
                            for fi44 := 0; fi44 < 16; fi44++ {
                                if ((((32767 < -380) && (0 != (v33 & (v33 + (v33 << ((42) & 63)))))) && false) || false) {
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    v43 = (fi44 / (((v43 / ((127) | 1))) | 1))
                                    v43 = (v43 ^ fi41)
                                    if false {
                                        v45 := GenS5{F0: 2147483647}
                                        _ = v45
                                    } else {
                                        if true {
                                            switch "\n\t" {
                                            case "hello":
                                                var v46 uint = 4294967295
                                                _ = v46
                                            case "foo bar":
                                                v1 = "hello"
                                            case "":
                                                var v47 bool = false
                                                _ = v47
                                            default:
                                                if len(v2) > 0 {
                                                    v2[0] = fi44
                                                }
                                            }
                                        }
                                    }
                                    switch (v43 & 4294967295) {
                                    case 32767:
                                        fmt.Println(v1)
                                    case 686:
                                        switch v1 {
                                        case "foo bar":
                                            if (!(-1 == -272)) {
                                                v48 := GenS6{F0: 42, F1: (^(v33))}
                                                _ = v48
                                            } else {
                                                if true {
                                                    if true {
                                                        var v49 int8 = 127
                                                        v50 := GenS7{F0: 2147483647}
                                                        var v51 uint16 = 475
                                                        _ = v49
                                                        _ = v50
                                                        _ = v51
                                                    } else {
                                                        v52 := GenS4{F0: 1}
                                                        _ = v52
                                                    }
                                                    if (4294967295 < 42) {
                                                        if true {
                                                            if (false || false) {
                                                                v53 := v3[((v1 + ("foo bar" + v1)) + "hello")]
                                                                if (42 < 42) {
                                                                    v54 := GenS8{F0: 18446744073709551615, F1: v53, F2: 128}
                                                                    _ = v54
                                                                }
                                                                delete(v3, "\n\t")
                                                                _ = v53
                                                            }
                                                        } else {
                                                            var v55 float64 = -2.25
                                                            _ = v55
                                                        }
                                                        v56 := GenS2{F0: 0, F1: 9223372036854775807, F2: -1}
                                                        if true {
                                                            v3[("a" + ("foo bar" + ""))] = (1 < 1)
                                                            v43 += (fi44 | 127)
                                                        }
                                                        _ = v56
                                                    }
                                                }
                                                v57 := GenS9{F0: 0}
                                                v43 -= (v43 - 9223372036854775807)
                                                _ = v57
                                            }
                                        default:
                                            var v58 bool = true
                                            _ = v58
                                        }
                                    default:
                                        fmt.Println(fi41)
                                    }
                                    var v59 int64 = 0
                                    v60 := GenS10{F0: int64(v33)}
                                    _ = v59
                                    _ = v60
                                }
                                v61 := v4[v33]
                                _ = fi44
                                _ = v61
                            }
                            v1 = v1
                        }
                        v33 -= 148
                    } else {
                        var v62 string = v1
                        _ = v62
                    }
                    _ = v43
                } else {
                    v33 = v33
                    for fi63 := 0; fi63 < 2; fi63++ {
                        switch v1 {
                        case "":
                            v33 = (v33 ^ 4294967295)
                        case "a":
                            switch ("foo bar" + (v1 + ("a" + (("hello" + ((("hello" + v1) + (v1 + "foo bar")) + (v1 + (v1 + "foo bar")))) + v1)))) {
                            case "hello":
                                v5 = append(v5, (v1 + ((v1 + v1) + "a")))
                            case "a":
                                v33 = (v33 & v33)
                            case "foo bar":
                                var v64 float32 = 1.5
                                _ = v64
                            }
                        default:
                            v33 += 2
                        }
                        switch fi63 {
                        case 32768:
                            v2 = append(v2, (fi63 >> ((42) & 63)))
                        case -1:
                            v33 = 4294967295
                        case -592:
                            v33 = 2
                        }
                        _ = fi63
                    }
                }
                for fi65 := 0; fi65 < 5; fi65++ {
                    switch fi65 {
                    case 9223372036854775807:
                        fmt.Println(fi65, v1)
                    case 32768:
                        switch (fi65 << ((2) & 63)) {
                        case 750:
                            var v66 float32 = float32(1e10)
                            _ = v66
                        case 4294967295:
                            if len(v5) > 0 {
                                v5[0] = ("" + ((v1 + "foo bar") + v1))
                            }
                        default:
                            var v67 int16 = -540
                            v68 := GenS6{F0: 1, F1: (v33 ^ (^(v33)))}
                            var v69 string = (v1 + v1)
                            _ = v67
                            _ = v68
                            _ = v69
                        }
                    default:
                        switch v1 {
                        case "hello":
                            var v70 int = fi65
                            _ = v70
                        case "foo bar":
                            v71 := GenS11{F0: 4294967295, F1: (fi41 * -9223372036854775808), F2: 989}
                            _ = v71
                        default:
                            if true {
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                v1 = (((("a" + "a") + "hello") + v1) + v1)
                                fmt.Println(v33)
                                var v72 int16 = 42
                                var v73 string = v1
                                _ = v72
                                _ = v73
                            } else {
                                var v74 int64 = 0
                                _ = v74
                            }
                            fmt.Println(v33)
                        }
                    }
                    switch "" {
                    case "hello":
                        var v75 uint64 = 1
                        _ = v75
                    case "a":
                        var v76 uint = 255
                        v77 := GenS5{F0: genFn0()}
                        v78 := GenS5{F0: -221}
                        _ = v76
                        _ = v77
                        _ = v78
                    default:
                        var v79 uint16 = 65535
                        _ = v79
                    }
                    _ = fi65
                }
                if ((105 != 0) && false) {
                    if false {
                        var v80 int16 = 638
                        _ = v80
                    } else {
                        fmt.Println(v33)
                        v1 = "hello"
                    }
                    switch (fi41 << ((165) & 63)) {
                    case 117:
                        var v81 string = (("\n\t" + ("" + v1)) + "")
                        _ = v81
                    case -829:
                        v33 -= 248
                    }
                } else {
                    var v82 uint32 = (v33 / (((v33 * (v33 >> ((65535) & 63)))) | 1))
                    switch (fi41 / (((fi41 % (((fi41 & fi41)) | 1))) | 1)) {
                    case 344:
                        v83 := GenS2{F0: 1, F1: fi41, F2: 42}
                        _ = v83
                    case 217:
                        v82 = (v82 + v33)
                    case -1:
                        var v84 string = v1
                        _ = v84
                    default:
                        if true {
                            fmt.Println(v82, v82)
                            v82 -= v82
                        } else {
                            v1 = "a"
                            var v85 uint32 = 1
                            _ = v85
                        }
                    }
                    _ = v82
                }
            }
        }
        _ = fi41
    }
    if (1 < -1) {
        v2 = append(v2, 670)
        v1 = "a"
    }
    v2 = append(v2, 9223372036854775807)
    if ((v1 > ((v1 + (v1 + v1)) + "hello")) || (2147483648 <= v33)) {
        for fi86 := 0; fi86 < 9; fi86++ {
            v33 -= (v33 << ((42) & 63))
            v87 := GenS12{F0: (42 == 42)}
            _ = fi86
            _ = v87
        }
        var v88 string = v1
        _ = v88
    } else {
        v33 = (v33 / (((v33 | 1)) | 1))
    }
    v89 := GenS10{F0: 9223372036854775807}
    var v90 int16 = 0
    var v91 uint = 255
    v92 := GenS3{F0: 42}
    v93 := GenS13{F0: "hello", F1: 76, F2: v91}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v90:", v90)
    fmt.Println("v91:", v91)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v89
    _ = v90
    _ = v91
    _ = v92
    _ = v93
}
