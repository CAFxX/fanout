package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8 }
type GenS2 struct { F0 bool; F1 uint32; F2 string }
type GenS3 struct { F0 uint8; F1 uint32 }
type GenS4 struct { F0 int8; F1 int32; F2 int16 }
type GenS5 struct { F0 uint64; F1 uint32 }
type GenS6 struct { F0 int64 }
func genFn0(p0 int32, p1 uint, p2 int8) string {
    p2 += p2
    v7 := GenS1{F0: p1, F1: -1}
    p1 -= p1
    fmt.Println(p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return "\n\t"
}
func genFn1(p0 bool, p1 uint32, p2 bool) int16 {
    v8 := GenS1{F0: 2, F1: -122}
    p1 += p1
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 42
}

func main() {
    var v1 float64 = (float64(-2.25) / float64(1.5))
    var v2 string = (("hello" + ("a" + ("a" + ""))) + ("" + ""))
    var v3 uint16 = 65535
    var v4 string = "foo bar"
    var v5 uint64 = 0
    v6 := []int{-36}
    v9 := GenS1{F0: 42, F1: 127}
    var v10 float64 = (1e10 / (float64(0.0) * v1))
    if (1.5 < (3.14159 + (0.1 * 0.0))) {
        fmt.Println(v2, v5, v4)
        for fi11 := 0; fi11 < 1; fi11++ {
            if ((fi11 / (((fi11 / ((4294967295) | 1))) | 1)) <= 127) {
                switch ("\n\t" + v4) {
                case "\n\t":
                    if len(v6) > 0 {
                        v6[0] = fi11
                    }
                case "hello":
                    fmt.Println(v5, v1)
                case "a":
                    for fi12 := 0; fi12 < 5; fi12++ {
                        if true {
                            if len(v6) > 0 {
                                v6[0] = (fi11 / ((566) | 1))
                            }
                            var v13 uint16 = (v3 / (((^(v3))) | 1))
                            _ = v13
                        } else {
                            if (false && (-977 > 9223372036854775807)) {
                                var v14 int16 = genFn1(false, 4294967295, false)
                                if ((v14 << ((201) & 63)) >= (v14 << ((0) & 63))) {
                                    v5 -= (v5 + v5)
                                } else {
                                    v15 := GenS1{F0: 18446744073709551615, F1: 42}
                                    var v16 uint64 = v5
                                    switch 32768 {
                                    case 9223372036854775807:
                                        if true {
                                            v17 := GenS1{F0: 580, F1: 42}
                                            v6 = append(v6, (^(fi12)))
                                            if (189 != uint8(v3)) {
                                                fmt.Println(v16, v5)
                                                if false {
                                                    v16 -= v16
                                                    var v18 uint32 = 2147483648
                                                    switch 32768 {
                                                    case 261:
                                                        v4 = ((v4 + v2) + v4)
                                                    case -2147483648:
                                                        fmt.Println(v5, v3, v2)
                                                    }
                                                    v18 += 8
                                                    if true {
                                                        switch (v2 + v2) {
                                                        case "a":
                                                            if ((v16 << ((670) & 63)) != (^(v16))) {
                                                                v16 = (v16 - 42)
                                                                var v19 uint16 = v3
                                                                var v20 uint16 = v3
                                                                _ = v19
                                                                _ = v20
                                                            }
                                                            v18 = v18
                                                            var v21 uint16 = v3
                                                            _ = v21
                                                        case "":
                                                            v3 = v3
                                                        case "hello":
                                                            var v22 uint16 = (v3 * v3)
                                                            var v23 uint = 1
                                                            if len(v6) > 0 {
                                                                v6[0] = fi12
                                                            }
                                                            _ = v22
                                                            _ = v23
                                                        default:
                                                            fmt.Println(v4, v14)
                                                        }
                                                        var v24 int16 = (v14 - -1)
                                                        _ = v24
                                                    } else {
                                                        fmt.Println(fi12, v3, v16)
                                                        v25 := GenS1{F0: 890, F1: -1}
                                                        _ = v25
                                                    }
                                                    var v26 uint64 = v5
                                                    v27 := GenS2{F0: false, F1: v18, F2: v4}
                                                    _ = v18
                                                    _ = v26
                                                    _ = v27
                                                }
                                            }
                                            var v28 uint8 = 231
                                            _ = v17
                                            _ = v28
                                        }
                                    case 199:
                                        if ((1 == 42) || (v16 < v5)) {
                                            var v29 uint = 2
                                            v29 -= 0
                                            v30 := GenS3{F0: 1, F1: 1}
                                            var v31 int8 = -42
                                            v4 = ("a" + (v4 + "\n\t"))
                                            v2 = v4
                                            fmt.Println(v31)
                                            _ = v29
                                            _ = v30
                                            _ = v31
                                        } else {
                                            switch v4 {
                                            case "":
                                                var v32 string = v2
                                                _ = v32
                                            case "hello":
                                                if len(v6) >= 2 {
                                                    v6 = v6[1:len(v6)]
                                                }
                                            default:
                                                switch "a" {
                                                case "\n\t":
                                                    var v33 bool = true
                                                    _ = v33
                                                default:
                                                    fmt.Println(v3)
                                                }
                                            }
                                            v14 = (v14 + (v14 & 1000))
                                            switch fi11 {
                                            case -2147483648:
                                                switch ("hello" + ((v2 + (v2 + (v2 + v4))) + "hello")) {
                                                case "\n\t":
                                                    v4 = v2
                                                default:
                                                    if false {
                                                        fmt.Println(v16, v2, v2)
                                                        v34 := GenS3{F0: 128, F1: 287}
                                                        _ = v34
                                                    }
                                                    var v35 int16 = (v14 ^ (v14 + genFn1(((v14 | 318) <= 42), 65535, true)))
                                                    v1 = v10
                                                    _ = v35
                                                }
                                            }
                                        }
                                    }
                                    _ = v15
                                    _ = v16
                                }
                                switch v4 {
                                case "foo bar":
                                    var v36 int64 = -503
                                    _ = v36
                                case "hello":
                                    var v37 int32 = -1
                                    _ = v37
                                }
                                _ = v14
                            } else {
                                var v38 float32 = 3.14159
                                v10 += v1
                                v39 := GenS3{F0: 146, F1: 42}
                                _ = len(v6)
                                _ = v38
                                _ = v39
                            }
                        }
                        v4 = v2
                        fmt.Println(v2, v10, v1)
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                        _ = fi12
                    }
                }
            } else {
                if (!true) {
                    v40 := GenS4{F0: -128, F1: 0, F2: -144}
                    if (127 == (fi11 * (fi11 ^ 127))) {
                        var v41 uint64 = v5
                        for fi42 := 0; fi42 < 8; fi42++ {
                            var v43 int32 = -86
                            v44 := GenS4{F0: 1, F1: v43, F2: -610}
                            v2 = ""
                            _ = fi42
                            _ = v43
                            _ = v44
                        }
                        v1 = v10
                        _ = v41
                    } else {
                        var v45 uint = 0
                        _ = v45
                    }
                    for fi46 := 0; fi46 < 14; fi46++ {
                        var v47 uint64 = (v5 - v5)
                        var v48 int16 = -1
                        _ = fi46
                        _ = v47
                        _ = v48
                    }
                    _ = v40
                } else {
                    for fi49 := 0; fi49 < 12; fi49++ {
                        var v50 uint64 = (v5 * v5)
                        v51 := GenS5{F0: v5, F1: 853}
                        if len(v6) > 0 {
                            v6[0] = (fi11 << ((uint(v50)) & 63))
                        }
                        fmt.Println(v5, v2, v50)
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                        var v52 float32 = -2.25
                        if true {
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                            v52 -= v52
                            if true {
                                var v53 int32 = -1
                                switch fi49 {
                                case 558:
                                    var v54 int16 = 32767
                                    _ = v54
                                case -2147483648:
                                    var v55 bool = true
                                    _ = v55
                                case 42:
                                    v53 += 2147483647
                                default:
                                    var v56 float64 = (float64((v52 * (3.14159 - v52))) + float64(((v52 * 3.14159) + v52)))
                                    _ = v56
                                }
                                _ = v53
                            } else {
                                fmt.Println(v10, v52)
                                var v57 uint64 = v5
                                _ = v57
                            }
                        }
                        _ = fi49
                        _ = v50
                        _ = v51
                        _ = v52
                    }
                    switch (fi11 >> ((4294967295) & 63)) {
                    case 127:
                        for fi58 := 0; fi58 < 2; fi58++ {
                            v10 -= (v1 + ((v1 - (v10 + ((0.1 / 3.14159) + (-2.25 * v1)))) / (((float64(0.0) + v1) + float64(1e10)) - 3.14159)))
                            switch fi58 {
                            case 42:
                                var v59 uint16 = v3
                                _ = v59
                            case 9223372036854775807:
                                _ = len(v6)
                            case -609:
                                fmt.Println(v1, v2)
                                if true {
                                    var v60 float64 = v1
                                    _ = v60
                                } else {
                                    if len(v6) > 0 {
                                        v6[0] = fi11
                                    }
                                    fmt.Println(v1, fi11, v5)
                                }
                                v4 = v4
                            }
                            var v61 uint64 = 42
                            switch 32768 {
                            case 2147483647:
                                var v62 int16 = 42
                                _ = v62
                            case 42:
                                var v63 uint = 1
                                _ = v63
                            case 101:
                                fmt.Println(v1, v1)
                            }
                            v1 -= -2.25
                            _ = fi58
                            _ = v61
                        }
                    case -211:
                        var v64 int64 = -457
                        _ = v64
                    }
                    for fi65 := 0; fi65 < 14; fi65++ {
                        v1 = 0.1
                        _ = len(v6)
                        _ = fi65
                    }
                }
                var v66 uint32 = 0
                if true {
                    fmt.Println(v5, v2)
                    switch int(v3) {
                    case 128:
                        var v67 bool = (1 < 694)
                        _ = v67
                    case -350:
                        if len(v6) > 0 {
                            v6[0] = -196
                        }
                    default:
                        v3 -= 32767
                    }
                    var v68 int8 = -1
                    _ = v68
                } else {
                    fmt.Println(v1)
                }
                _ = v66
            }
            v6 = append(v6, 32768)
            _ = fi11
        }
    }
    for fi69 := 0; fi69 < 14; fi69++ {
        fmt.Println(v10)
        _ = fi69
    }
    fmt.Println(v1)
    for fi70 := 0; fi70 < 14; fi70++ {
        v4 = v2
        _ = fi70
    }
    switch -273 {
    case -2147483648:
        var v71 uint8 = 37
        _ = v71
    case 836:
        var v72 int32 = -426
        _ = v72
    case 939:
        v73 := GenS3{F0: 127, F1: 65535}
        _ = v73
    default:
        v74 := GenS6{F0: 1}
        _ = v74
    }
    if (!(-35 != 39)) {
        var v75 uint32 = 0
        var v76 uint = 42
        _ = v75
        _ = v76
    }
    if len(v6) > 0 {
        v6[0] = 69
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
