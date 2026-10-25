package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint16; F1 int16; F2 uint }
type GenS3 struct { F0 int64; F1 int32; F2 uint32 }
type GenS4 struct { F0 int32; F1 int16; F2 int16 }
type GenS5 struct { F0 uint; F1 uint32 }
type GenS6 struct { F0 int64; F1 bool; F2 int8 }
type GenS7 struct { F0 string; F1 int8 }
func genFn0(p0 int64, p1 int16) uint16 {
    switch -756 {
    case -1:
        for fi6 := 0; fi6 < 12; fi6++ {
            var v7 uint64 = 160
            _ = fi6
            _ = v7
        }
    case 0:
        v8 := []int{-751, 329, -2147483648}
        _ = v8
    default:
        for fi9 := 0; fi9 < 16; fi9++ {
            p1 += (p1 << ((2) & 63))
            _ = fi9
        }
    }
    p0 += (p0 - (p0 >> ((18446744073709551615) & 63)))
    if (!false) {
        switch 42 {
        case 604:
            v10 := GenS1{F0: -1}
            _ = v10
        case -2147483648:
            switch -2147483648 {
            case -2147483648:
                var v11 int16 = -410
                _ = v11
            case 2147483647:
                var v12 bool = (false || (p0 != (^(p0))))
                _ = v12
            }
        case 128:
            fmt.Println(p0)
        default:
            p0 = p0
        }
        if true {
            var v13 int16 = p1
            _ = v13
        } else {
            p0 = p0
        }
    } else {
        v14 := []int{}
        _ = v14
    }
    _ = p0
    _ = p1
    return 65535
}

func main() {
    var v1 uint64 = 528
    var v2 int16 = 32767
    v3 := []int{2147483647, 128, 1, -9223372036854775808}
    v4 := make(map[int]int)
    var v5 int64 = 105
    if false {
        for fi15 := 0; fi15 < 13; fi15++ {
            switch "a" {
            case "a":
                v1 = (v1 % ((v1) | 1))
            case "foo bar":
                _ = len(v4)
            case "\n\t":
                v16 := GenS2{F0: genFn0(v5, (^(v2))), F1: (^(v2)), F2: 631}
                _ = v16
            }
            var v17 uint32 = 65535
            switch "hello" {
            case "":
                delete(v4, fi15)
            case "hello":
                if false {
                    fmt.Println(v2, v17)
                    for fi18 := 0; fi18 < 13; fi18++ {
                        fmt.Println(v5)
                        if ((v17 << ((42) & 63)) >= v17) {
                            var v19 uint64 = uint64(v17)
                            var v20 int64 = v5
                            _ = v19
                            _ = v20
                        } else {
                            v2 = int16(fi15)
                        }
                        v2 -= v2
                        _ = fi18
                    }
                }
            case "a":
                _ = len(v4)
            default:
                var v21 uint16 = 42
                _ = v21
            }
            _ = fi15
            _ = v17
        }
    }
    for fi22 := 0; fi22 < 17; fi22++ {
        if ((v2 << ((uint(v1)) & 63)) != v2) {
            var v23 uint8 = 128
            for fi24 := 0; fi24 < 5; fi24++ {
                if (-92 < 111) {
                    switch "hello" {
                    case "a":
                        if ((42 >= -1) || ((v5 >> ((42) & 63)) < v5)) {
                            v1 += v1
                            v25 := GenS3{F0: v5, F1: 0, F2: 2147483647}
                            _ = v25
                        }
                    }
                    fmt.Println(v23, v1)
                    v5 += v5
                    v5 += (v5 & 42)
                } else {
                    var v26 int64 = (-(v5))
                    switch fi24 {
                    case 204:
                        var v27 int8 = 42
                        _ = v27
                    }
                    _ = v26
                }
                var v28 uint8 = uint8(v1)
                v28 -= v23
                _ = fi24
                _ = v28
            }
            if true {
                var v29 int64 = (v5 * int64(v1))
                v30 := GenS4{F0: -663, F1: v2, F2: v2}
                v5 -= v29
                _ = v29
                _ = v30
            }
            _ = v23
        }
        for fi31 := 0; fi31 < 6; fi31++ {
            var v32 int = (fi31 / ((fi31) | 1))
            v4[1] = (fi22 >> ((531) & 63))
            var v33 uint16 = 255
            _ = fi31
            _ = v32
            _ = v33
        }
        _ = fi22
    }
    v2 = (v2 - -32768)
    var v34 int = 32767
    v35 := GenS5{F0: 219, F1: 0}
    switch (v34 | 127) {
    case 4294967295:
        for fi36 := 0; fi36 < 14; fi36++ {
            v5 = v5
            if (0 == 32767) {
                var v37 int32 = 42
                if ((!true) && true) {
                    v37 = (v37 ^ (v37 / (((^(v37))) | 1)))
                    delete(v4, 0)
                } else {
                    v5 -= v5
                }
                var v38 int = (-(v34))
                _ = v37
                _ = v38
            }
            _ = len(v4)
            _ = fi36
        }
    }
    if false {
        switch "\n\t" {
        case "hello":
            if (!("hello" == "foo bar")) {
                fmt.Println(v5)
                var v39 int16 = 1
                _ = v39
            }
        default:
            var v40 bool = (true || (!true))
            _ = v40
        }
        fmt.Println(v1, v34, v1)
    } else {
        var v41 int32 = 0
        v42 := GenS1{F0: 0}
        _ = v41
        _ = v42
    }
    for fi43 := 0; fi43 < 9; fi43++ {
        for fi44 := 0; fi44 < 14; fi44++ {
            if ((v1 >> ((18446744073709551615) & 63)) >= v1) {
                _ = len(v3)
                if (0 <= 0) {
                    if false {
                        var v45 int32 = 0
                        v1 += (v1 & (v1 + (v1 / ((v1) | 1))))
                        var v46 int = (v34 / ((fi43) | 1))
                        if (!(!(!(2 == 42)))) {
                            fmt.Println(fi43)
                            delete(v4, fi43)
                            var v47 uint64 = v1
                            _ = v47
                        } else {
                            if false {
                                var v48 bool = false
                                v49 := GenS1{F0: 54}
                                _ = v48
                                _ = v49
                            }
                            var v50 float64 = (0.0 - -2.25)
                            _ = v50
                        }
                        var v51 uint = 618
                        v52 := GenS2{F0: 65535, F1: 32767, F2: v51}
                        _ = v45
                        _ = v46
                        _ = v51
                        _ = v52
                    }
                    switch ("foo bar" + (("" + "foo bar") + "\n\t")) {
                    case "hello":
                        if false {
                            if (!true) {
                                _ = len(v3)
                                v34 += int(v2)
                            } else {
                                v53 := v4[fi43]
                                switch fi43 {
                                case 4294967295:
                                    if len(v3) > 0 {
                                        v3[0] = v34
                                    }
                                case 127:
                                    if (!((4294967295 != 2147483648) && ((-(v5)) != (v5 * v5)))) {
                                        v5 += (v5 + (v5 & 0))
                                        v54 := GenS6{F0: v5, F1: (!((!true) && ((((v2 ^ v2) <= (^(v2))) || true) || (v5 < (v5 & v5))))), F2: -1}
                                        _ = v54
                                    } else {
                                        v55 := GenS5{F0: 18446744073709551615, F1: 2147483647}
                                        if ((!true) && true) {
                                            var v56 uint16 = 1
                                            _ = v56
                                        }
                                        _ = v55
                                    }
                                case 51:
                                    v5 = v5
                                }
                                _ = v53
                            }
                        }
                    default:
                        fmt.Println(fi44, v34)
                    }
                } else {
                    v5 += (v5 << ((65535) & 63))
                    switch "hello" {
                    case "foo bar":
                        if len(v3) > 0 {
                            v3[0] = 0
                        }
                    case "a":
                        switch 42 {
                        case 966:
                            if (2147483647 != -2147483648) {
                                delete(v4, -414)
                                var v57 int8 = 43
                                var v58 int32 = 42
                                _ = v57
                                _ = v58
                            } else {
                                switch fi44 {
                                case 32768:
                                    v5 = v5
                                case 636:
                                    v59 := v4[(^(fi43))]
                                    _ = v59
                                }
                                fmt.Println(v34, fi44)
                            }
                        default:
                            if ((("\n\t" + ("\n\t" + ("foo bar" + ""))) <= "\n\t") || (fi44 != (v34 / ((fi44) | 1)))) {
                                v60 := GenS7{F0: ("\n\t" + "hello"), F1: -128}
                                _ = v60
                            }
                        }
                    default:
                        var v61 string = "\n\t"
                        _ = v61
                    }
                    switch 32767 {
                    case 840:
                        switch "foo bar" {
                        case "\n\t":
                            var v62 uint64 = (^(v1))
                            _ = v62
                        }
                    case 32768:
                        fmt.Println(fi44, v34, v2)
                    case -367:
                        fmt.Println(v5, v5, v5)
                    default:
                        v63 := GenS6{F0: v5, F1: (!(591 >= 2)), F2: -119}
                        _ = v63
                    }
                    var v64 bool = (v1 < (v1 << ((42) & 63)))
                    if false {
                        switch fi43 {
                        case 128:
                            v1 = v1
                        }
                        var v65 int64 = v5
                        _ = v65
                    }
                    _ = v64
                }
                v1 = (v1 + (v1 | (v1 << ((124) & 63))))
                v34 -= fi43
            }
            _ = fi44
        }
        switch ("hello" + ("a" + "\n\t")) {
        case "":
            delete(v4, 9223372036854775807)
        case "hello":
            for fi66 := 0; fi66 < 6; fi66++ {
                if (!(true && false)) {
                    switch (("a" + ("foo bar" + "")) + ("" + "hello")) {
                    case "\n\t":
                        fmt.Println(v34, v5, v34)
                    case "foo bar":
                        fmt.Println(v34, fi66, v1)
                    case "":
                        v67 := GenS6{F0: (v5 - v5), F1: true, F2: -78}
                        _ = v67
                    }
                    v68 := v4[-2147483648]
                    _ = v68
                }
                v34 -= (-(fi66))
                fmt.Println(fi66, v5, fi43)
                _ = fi66
            }
        }
        _ = fi43
    }
    for fi69 := 0; fi69 < 3; fi69++ {
        fmt.Println(v2)
        v70 := v4[(fi69 << ((0) & 63))]
        v71 := GenS5{F0: 390, F1: 42}
        _ = fi69
        _ = v70
        _ = v71
    }
    if (0.0 == (3.14159 - (0.1 + 1e10))) {
        _ = len(v4)
    }
    if true {
        fmt.Println(v5, v5)
        v3 = append(v3, (v34 >> ((18446744073709551615) & 63)))
        switch "hello" {
        case "":
            var v72 uint16 = 490
            _ = v72
        case "hello":
            var v73 int64 = 9223372036854775807
            _ = v73
        default:
            v74 := v4[(v34 * v34)]
            _ = v74
        }
    }
    v75 := GenS7{F0: "\n\t", F1: 95}
    var v76 uint64 = v1
    fmt.Println(v2)
    var v77 uint = 4294967295
    for fi78 := 0; fi78 < 1; fi78++ {
        if (1.5 == -2.25) {
            switch fi78 {
            case 0:
                switch v34 {
                case 127:
                    v34 = (^(v34))
                case 42:
                    var v79 bool = ((0.1 * (-2.25 / -2.25)) <= -2.25)
                    _ = v79
                default:
                    v80 := GenS3{F0: v5, F1: 0, F2: 65535}
                    _ = v80
                }
            default:
                v4[(fi78 << (((v77 & (v77 << ((1) & 63)))) & 63))] = 42
            }
        }
        var v81 int = (v34 ^ v34)
        _ = fi78
        _ = v81
    }
    var v82 uint64 = 4294967295
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v76:", v76)
    fmt.Println("v77:", v77)
    fmt.Println("v82:", v82)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v75
    _ = v76
    _ = v77
    _ = v82
}
