package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64 }
type GenS2 struct { F0 int; F1 uint; F2 uint8 }
type GenS3 struct { F0 uint; F1 int64; F2 int16 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int8; F1 int; F2 int8 }
type GenS6 struct { F0 uint32; F1 int32; F2 uint64 }
func genFn0() uint16 {
    if false {
        switch 502 {
        case -660:
            var v4 string = "hello"
            _ = v4
        default:
            v5 := GenS1{F0: 4294967295, F1: -147}
            _ = v5
        }
        var v6 int64 = 9223372036854775807
        _ = v6
    }
    var v7 string = "hello"
    for fi8 := 0; fi8 < 5; fi8++ {
        v9 := GenS1{F0: (-(fi8)), F1: -185}
        _ = fi8
        _ = v9
    }
    if (612 < 4294967295) {
        for fi10 := 0; fi10 < 9; fi10++ {
            v11 := []int{fi10, (^(fi10)), (fi10 & (^(fi10)))}
            if (3.14159 >= (0.1 / 3.14159)) {
                for fi12 := 0; fi12 < 19; fi12++ {
                    switch (fi10 | (fi10 % ((fi10) | 1))) {
                    case 1:
                        v7 = "\n\t"
                    default:
                        if (2 != 127) {
                            v7 = v7
                            var v13 string = v7
                            _ = v13
                        }
                    }
                    _ = fi12
                }
                var v14 uint = 216
                if len(v11) > 0 {
                    v11[0] = fi10
                }
                _ = v14
            } else {
                var v15 int32 = 42
                _ = v15
            }
            _ = fi10
            _ = v11
        }
        fmt.Println(v7)
        if (201 != 0) {
            v7 = v7
        }
    } else {
        v7 = v7
        v7 = v7
    }
    _ = v7
    return 541
}

func main() {
    var v1 int = -700
    v2 := []uint32{65535, 4294967295}
    v3 := []uint32{}
    if len(v3) > 0 {
        v3[0] = 2147483648
    }
    switch "foo bar" {
    case "foo bar":
        v16 := GenS2{F0: v1, F1: 65535, F2: 42}
        _ = v16
    default:
        switch v1 {
        case 4294967295:
            var v17 float32 = 0.1
            _ = v17
        case -9223372036854775808:
            v1 = v1
        case 235:
            v18 := GenS2{F0: v1, F1: 584, F2: 1}
            _ = v18
        }
    }
    switch (v1 - -1) {
    case -275:
        v1 -= v1
    case 1:
        if false {
            if true {
                var v19 float32 = (((1.5 * 0.1) / -2.25) + ((3.14159 / (1.5 * 3.14159)) * 3.14159))
                for fi20 := 0; fi20 < 11; fi20++ {
                    var v21 int32 = -2147483648
                    var v22 uint8 = 52
                    var v23 int = fi20
                    _ = fi20
                    _ = v21
                    _ = v22
                    _ = v23
                }
                _ = v19
            }
            fmt.Println(v1)
        }
    }
    var v24 bool = ("foo bar" == "a")
    for fi25 := 0; fi25 < 12; fi25++ {
        _ = len(v3)
        if (!(false && v24)) {
            v24 = v24
        }
        v1 += (fi25 | v1)
        _ = fi25
    }
    fmt.Println(v24, v24)
    var v26 string = ("a" + "")
    for fi27 := 0; fi27 < 15; fi27++ {
        var v28 int = (v1 % (((fi27 - fi27)) | 1))
        _ = fi27
        _ = v28
    }
    for fi29 := 0; fi29 < 2; fi29++ {
        v24 = (2147483647 > -1)
        _ = fi29
    }
    var v30 int64 = 9223372036854775807
    if v24 {
        v30 += int64(v1)
    } else {
        for fi31 := 0; fi31 < 9; fi31++ {
            v1 -= fi31
            v26 = (v26 + v26)
            _ = fi31
        }
        if (!v24) {
            var v32 int = (^(v1))
            if ((42 > 1) && (-68 >= 1)) {
                if len(v3) > 0 {
                    v3[0] = 4294967295
                }
            } else {
                v33 := GenS3{F0: 4294967295, F1: -1, F2: 1}
                _ = v33
            }
            _ = v32
        } else {
            v34 := GenS4{F0: v30}
            _ = v34
        }
    }
    if true {
        for fi35 := 0; fi35 < 2; fi35++ {
            fmt.Println(fi35)
            if v24 {
                for fi36 := 0; fi36 < 1; fi36++ {
                    var v37 int64 = v30
                    var v38 uint64 = 18446744073709551615
                    _ = fi36
                    _ = v37
                    _ = v38
                }
                v30 += v30
                v39 := GenS2{F0: (v1 * fi35), F1: 1, F2: 127}
                _ = v39
            }
            if (false || ((float32((-2.25 + (float64(0.0) / (1.5 + -2.25)))) < (0.0 + 1.5)) && v24)) {
                var v40 int = fi35
                if (128 > fi35) {
                    var v41 bool = false
                    _ = v41
                }
                _ = v40
            }
            _ = fi35
        }
        v1 -= 1
        switch v26 {
        case "foo bar":
            for fi42 := 0; fi42 < 7; fi42++ {
                v43 := GenS4{F0: (v30 - 1)}
                _ = fi42
                _ = v43
            }
        case "\n\t":
            var v44 float64 = 3.14159
            _ = v44
        }
    } else {
        v30 = (v30 % ((v30) | 1))
        v45 := GenS1{F0: 2147483647, F1: v30}
        _ = v45
    }
    fmt.Println(v30, v1, v26)
    v46 := GenS5{F0: -128, F1: v1, F2: 62}
    if v24 {
        _ = len(v3)
        var v47 int8 = 1
        _ = v47
    }
    switch v1 {
    case 127:
        v48 := GenS3{F0: 255, F1: v30, F2: -842}
        _ = v48
    }
    if true {
        if ((705 >= genFn0()) && (1 < 909)) {
            var v49 int16 = -403
            v26 = (v26 + v26)
            _ = v49
        } else {
            var v50 int64 = v30
            _ = v50
        }
        switch 94 {
        case -475:
            if len(v3) > 0 {
                v3[0] = 65535
            }
        case 4294967295:
            fmt.Println(v24)
        case 32767:
            if false {
                v3 = append(v3, 2147483647)
                for fi51 := 0; fi51 < 7; fi51++ {
                    if true {
                        var v52 float64 = float64(float32(0.1))
                        _ = v52
                    } else {
                        v53 := GenS5{F0: 0, F1: fi51, F2: -128}
                        v2 = append(v2, 2147483648)
                        if v24 {
                            if ((v24 && (v24 && ((v24 && v24) && (true && (v1 >= (fi51 | (fi51 % ((fi51) | 1)))))))) || (false || ((!(1 <= 65535)) || v24))) {
                                v30 = (v30 * v30)
                                for fi54 := 0; fi54 < 13; fi54++ {
                                    v30 = (v30 - -1)
                                    v1 = fi54
                                    _ = fi54
                                }
                                var v55 uint16 = 960
                                _ = v55
                            } else {
                                for fi56 := 0; fi56 < 1; fi56++ {
                                    switch 461 {
                                    case -591:
                                        switch fi51 {
                                        case 4294967295:
                                            v57 := GenS1{F0: (fi51 & fi56), F1: (v30 % ((0) | 1))}
                                            _ = v57
                                        }
                                    case 586:
                                        v58 := GenS1{F0: (v1 | 1), F1: 338}
                                        _ = v58
                                    }
                                    fmt.Println(fi51, v1)
                                    var v59 int16 = 0
                                    _ = fi56
                                    _ = v59
                                }
                            }
                            for fi60 := 0; fi60 < 7; fi60++ {
                                v3 = append(v3, 4294967295)
                                v61 := GenS4{F0: v30}
                                if v24 {
                                    v24 = (29 >= 128)
                                } else {
                                    if (695 != 1) {
                                        var v62 float64 = (1e10 * (0.0 + 0.0))
                                        _ = v62
                                    } else {
                                        v63 := GenS6{F0: 2147483648, F1: 2147483647, F2: 1}
                                        _ = v63
                                    }
                                }
                                _ = fi60
                                _ = v61
                            }
                        }
                        _ = v53
                    }
                    v1 += (v1 % ((v1) | 1))
                    var v64 uint64 = 2
                    _ = fi51
                    _ = v64
                }
            } else {
                switch "a" {
                case "\n\t":
                    var v65 uint32 = 963
                    _ = v65
                }
                switch "hello" {
                case "\n\t":
                    if len(v3) > 0 {
                        v3[0] = 2147483648
                    }
                case "foo bar":
                    fmt.Println(v30)
                case "hello":
                    for fi66 := 0; fi66 < 1; fi66++ {
                        v67 := GenS6{F0: 712, F1: -1, F2: 18446744073709551615}
                        v2 = append(v2, 865)
                        v1 = -488
                        _ = fi66
                        _ = v67
                    }
                }
                v1 -= (^(v1))
            }
        }
        for fi68 := 0; fi68 < 12; fi68++ {
            var v69 float64 = 1.5
            _ = fi68
            _ = v69
        }
    } else {
        for fi70 := 0; fi70 < 2; fi70++ {
            for fi71 := 0; fi71 < 20; fi71++ {
                fmt.Println(v1, fi71)
                v72 := GenS3{F0: 0, F1: 42, F2: 1}
                _ = fi71
                _ = v72
            }
            _ = fi70
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    _ = v1
    _ = v2
    _ = v24
    _ = v26
    _ = v3
    _ = v30
    _ = v46
}
