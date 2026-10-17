package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int16; F1 int64; F2 int8 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint64; F1 int32 }
type GenS6 struct { F0 bool }
func genFn0(p0 string, p1 uint8) string {
    fmt.Println(p1, p1)
    if true {
        var v7 uint8 = p1
        _ = v7
    } else {
        p1 += p1
        v8 := []int{42, 9223372036854775807, -736}
        _ = v8
    }
    var v9 int8 = 127
    if (false || false) {
        p0 = (("foo bar" + p0) + p0)
    }
    v10 := GenS1{F0: 723}
    _ = p0
    _ = p1
    _ = v10
    _ = v9
    return "hello"
}

func main() {
    var v1 float64 = -2.25
    var v2 uint = 4294967295
    var v3 uint32 = 4294967295
    var v4 float64 = (0.1 * (0.0 - v1))
    var v5 bool = (("" >= "") || true)
    var v6 uint8 = 127
    v4 -= v4
    switch "" {
    case "\n\t":
        fmt.Println(v6)
    }
    v11 := []int{1, -265, 9223372036854775807}
    v12 := GenS1{F0: -9223372036854775808}
    var v13 int8 = 1
    var v14 uint32 = (v3 | (^(v3)))
    for fi15 := 0; fi15 < 11; fi15++ {
        v3 -= 1
        _ = fi15
    }
    var v16 uint64 = 1
    if (v4 >= -2.25) {
        switch 32767 {
        case 0:
            if v5 {
                fmt.Println(v13)
                v13 -= (v13 ^ v13)
                var v17 uint16 = 65535
                _ = v17
            }
        case -9223372036854775808:
            var v18 bool = v5
            _ = v18
        case -384:
            if v5 {
                var v19 int = 0
                _ = v19
            } else {
                v14 = v14
            }
        default:
            if len(v11) > 0 {
                v11[0] = 279
            }
        }
        v14 = v14
    } else {
        v6 = 128
        switch -271 {
        case 0:
            if ((v3 < uint32(v16)) || (!((1.5 * 3.14159) == v1))) {
                switch "\n\t" {
                case "":
                    for fi20 := 0; fi20 < 13; fi20++ {
                        _ = len(v11)
                        if len(v11) > 0 {
                            v11[0] = fi20
                        }
                        fmt.Println(v16, v2, v13)
                        _ = fi20
                    }
                case "a":
                    if len(v11) >= 2 {
                        v11 = v11[1:len(v11)]
                    }
                case "foo bar":
                    if len(v11) > 0 {
                        v11[0] = 42
                    }
                default:
                    for fi21 := 0; fi21 < 18; fi21++ {
                        fmt.Println(v2, v4, v16)
                        _ = fi21
                    }
                }
                for fi22 := 0; fi22 < 1; fi22++ {
                    var v23 int16 = 32767
                    if v5 {
                        var v24 uint8 = (v6 - 2)
                        _ = v24
                    }
                    var v25 uint8 = (v6 & v6)
                    _ = fi22
                    _ = v23
                    _ = v25
                }
                switch 9223372036854775807 {
                case 0:
                    var v26 int16 = 609
                    _ = v26
                }
            }
        case 32768:
            for fi27 := 0; fi27 < 1; fi27++ {
                v11 = append(v11, 32768)
                if len(v11) > 0 {
                    v11[0] = (fi27 / (((fi27 % ((123) | 1))) | 1))
                }
                if v5 {
                    v28 := GenS1{F0: -1}
                    if (!((v2 & v2) > (^(v2)))) {
                        for fi29 := 0; fi29 < 11; fi29++ {
                            if v5 {
                                v3 = (v14 | uint32(v16))
                                if len(v11) > 0 {
                                    v11[0] = fi29
                                }
                            } else {
                                var v30 int64 = -1
                                _ = v30
                            }
                            if len(v11) > 0 {
                                v11[0] = fi27
                            }
                            _ = fi29
                        }
                    } else {
                        for fi31 := 0; fi31 < 8; fi31++ {
                            if ("\n\t" < (("" + "\n\t") + "a")) {
                                if (v5 && v5) {
                                    v4 -= 0.0
                                } else {
                                    var v32 string = "hello"
                                    _ = v32
                                }
                                if v5 {
                                    switch fi31 {
                                    case 32767:
                                        var v33 bool = (((!false) && (2147483647 > 2147483647)) || (v5 || (v5 && v5)))
                                        _ = v33
                                    case -671:
                                        v34 := GenS2{F0: v2}
                                        _ = v34
                                    case 4294967295:
                                        v35 := GenS2{F0: 754}
                                        _ = v35
                                    }
                                    var v36 float32 = 1e10
                                    _ = v36
                                }
                            } else {
                                v5 = v5
                                if (v16 <= (v16 % ((v16) | 1))) {
                                    if v5 {
                                        if (!(!false)) {
                                            var v37 int16 = 1000
                                            v38 := GenS3{F0: (-(v37)), F1: 964, F2: -91}
                                            v39 := GenS2{F0: v2}
                                            _ = v37
                                            _ = v38
                                            _ = v39
                                        }
                                    }
                                    v40 := GenS3{F0: -717, F1: 1, F2: int8(v3)}
                                    if ("foo bar" <= "\n\t") {
                                        v5 = ((v16 % (((v16 * 42)) | 1)) < (v16 - v16))
                                        v41 := GenS3{F0: 1000, F1: -818, F2: (v13 & (-(v13)))}
                                        v4 += 0.0
                                        _ = v41
                                    }
                                    _ = v40
                                } else {
                                    if (v5 && v5) {
                                        if (fi31 != (fi27 ^ fi31)) {
                                            v11 = append(v11, (fi31 / (((fi27 * fi31)) | 1)))
                                        } else {
                                            switch (("\n\t" + ("a" + "")) + ("\n\t" + "foo bar")) {
                                            case "foo bar":
                                                v42 := GenS2{F0: (v2 % ((750) | 1))}
                                                _ = v42
                                            case "\n\t":
                                                v13 += (v13 % (((v13 >> (((v2 & 2)) & 63))) | 1))
                                            case "a":
                                                v43 := GenS1{F0: -1}
                                                _ = v43
                                            default:
                                                v16 = v16
                                            }
                                            switch "" {
                                            case "\n\t":
                                                switch "\n\t" {
                                                case "":
                                                    v2 += 776
                                                case "a":
                                                    var v44 string = "\n\t"
                                                    _ = v44
                                                case "foo bar":
                                                    v11 = append(v11, -2147483648)
                                                }
                                            case "a":
                                                if len(v11) > 0 {
                                                    v11[0] = fi31
                                                }
                                            case "foo bar":
                                                if (false || v5) {
                                                    if (v5 && (!(!((v6 / ((42) | 1)) != (^(v6)))))) {
                                                        var v45 uint = 726
                                                        v2 = 606
                                                        v11 = append(v11, (fi27 ^ (fi31 ^ (fi31 & fi27))))
                                                        _ = v45
                                                    } else {
                                                        if v5 {
                                                            v46 := GenS4{F0: 0}
                                                            var v47 int64 = 0
                                                            var v48 uint8 = 255
                                                            _ = v46
                                                            _ = v47
                                                            _ = v48
                                                        }
                                                        v16 -= (v16 << (((^(v2))) & 63))
                                                        if (3.14159 == 1.5) {
                                                            fmt.Println(fi31)
                                                            v49 := GenS4{F0: uint16(v16)}
                                                            _ = v49
                                                        } else {
                                                            v5 = (1.5 < 1.5)
                                                            var v50 uint16 = 42
                                                            var v51 uint = v2
                                                            _ = v50
                                                            _ = v51
                                                        }
                                                    }
                                                    var v52 uint64 = (v16 - v16)
                                                    if (false || v5) {
                                                        fmt.Println(v3, v6, fi31)
                                                    }
                                                    _ = v52
                                                } else {
                                                    v4 -= float64(-2.25)
                                                }
                                            default:
                                                var v53 uint16 = 148
                                                _ = v53
                                            }
                                            var v54 int8 = 100
                                            _ = v54
                                        }
                                        fmt.Println(v16, fi31)
                                        v55 := GenS5{F0: v16, F1: -522}
                                        _ = v55
                                    }
                                    var v56 int32 = -2147483648
                                    v2 -= v2
                                    _ = v56
                                }
                            }
                            var v57 bool = (-63 < -2147483648)
                            _ = fi31
                            _ = v57
                        }
                    }
                    _ = v28
                } else {
                    for fi58 := 0; fi58 < 17; fi58++ {
                        var v59 float32 = 1e10
                        var v60 float64 = (0.0 + (float64((1e10 - float32(3.14159))) + v4))
                        var v61 uint8 = (v6 >> ((65535) & 63))
                        _ = fi58
                        _ = v59
                        _ = v60
                        _ = v61
                    }
                    fmt.Println(v13, v1, v3)
                    switch 603 {
                    case 32768:
                        v62 := GenS6{F0: (1 > 9223372036854775807)}
                        _ = v62
                    case -633:
                        var v63 float32 = float32((v4 * v4))
                        _ = v63
                    case 760:
                        if len(v11) > 0 {
                            v11[0] = (fi27 & fi27)
                        }
                    default:
                        v11 = append(v11, 2147483647)
                    }
                }
                _ = fi27
            }
        case 127:
            var v64 int8 = v13
            _ = v64
        }
    }
    if (-1 > -32768) {
        if len(v11) > 0 {
            v11[0] = 1
        }
        var v65 uint8 = (v6 + v6)
        _ = v65
    } else {
        v6 = 127
        switch ("foo bar" + (("a" + "\n\t") + "hello")) {
        case "\n\t":
            for fi66 := 0; fi66 < 6; fi66++ {
                v67 := GenS5{F0: (^(v16)), F1: 1}
                _ = fi66
                _ = v67
            }
        case "hello":
            var v68 bool = (v5 || v5)
            _ = v68
        }
    }
    v5 = (!((v2 % (((v2 & uint(v3))) | 1)) <= v2))
    for fi69 := 0; fi69 < 15; fi69++ {
        v70 := GenS5{F0: (v16 / (((^(v16))) | 1)), F1: 0}
        _ = fi69
        _ = v70
    }
    v11 = append(v11, 32768)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
