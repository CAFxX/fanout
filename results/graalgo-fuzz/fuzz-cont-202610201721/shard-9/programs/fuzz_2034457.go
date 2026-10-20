package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int8; F1 int64; F2 int }
type GenS3 struct { F0 uint; F1 uint64; F2 uint64 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int; F1 int8; F2 uint32 }
type GenS6 struct { F0 int16; F1 uint }
func genFn0(p0 bool) uint {
    v4 := []int{745}
    switch 0 {
    case -443:
        v5 := GenS1{F0: 41}
        _ = v5
    case -308:
        if (736 >= 42) {
            if (!(false && true)) {
                for fi6 := 0; fi6 < 5; fi6++ {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    if false {
                        for fi7 := 0; fi7 < 3; fi7++ {
                            _ = len(v4)
                            _ = fi7
                        }
                        var v8 float64 = float64(0.0)
                        switch (-(fi6)) {
                        case -2147483648:
                            switch "a" {
                            case "foo bar":
                                v9 := GenS1{F0: 42}
                                _ = v9
                            case "a":
                                _ = len(v4)
                            }
                        case 2147483647:
                            if len(v4) > 0 {
                                v4[0] = fi6
                            }
                        case -9223372036854775808:
                            if (!((2147483647 <= 1) && true)) {
                                switch "hello" {
                                case "foo bar":
                                    v4 = append(v4, -9223372036854775808)
                                case "hello":
                                    _ = len(v4)
                                case "a":
                                    fmt.Println(fi6, v8)
                                }
                                if true {
                                    switch ("foo bar" + ((("" + "") + "foo bar") + "")) {
                                    case "hello":
                                        switch "a" {
                                        case "hello":
                                            v8 = v8
                                        }
                                    }
                                    v8 -= (v8 * v8)
                                } else {
                                    v10 := GenS1{F0: 7}
                                    _ = v10
                                }
                                var v11 bool = ((!(465 != 65535)) && true)
                                _ = v11
                            }
                        }
                        _ = v8
                    }
                    var v12 int = (fi6 & fi6)
                    _ = fi6
                    _ = v12
                }
                v13 := GenS2{F0: -1, F1: -9223372036854775808, F2: 2147483647}
                _ = v13
            } else {
                v14 := GenS3{F0: 870, F1: 169, F2: 2}
                p0 = false
                _ = v14
            }
        }
    case 2147483647:
        if true {
            if (4294967295 != 835) {
                var v15 int32 = -2147483648
                var v16 string = "hello"
                _ = v15
                _ = v16
            } else {
                v4 = append(v4, -2147483648)
            }
            v4 = append(v4, 32767)
            var v17 uint64 = 389
            _ = v17
        }
    }
    for fi18 := 0; fi18 < 8; fi18++ {
        p0 = (42 > 65535)
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        for fi19 := 0; fi19 < 4; fi19++ {
            var v20 float64 = 3.14159
            fmt.Println(fi19)
            _ = fi19
            _ = v20
        }
        _ = fi18
    }
    _ = p0
    _ = v4
    return 682
}
func genFn1(p0 int, p1 uint16) uint {
    var v21 uint64 = 42
    v21 += 941
    v22 := []int{(p0 / ((p0) | 1))}
    var v23 bool = false
    if (p0 > p0) {
        fmt.Println(p0, p1, p1)
        switch p0 {
        case 2147483647:
            switch (p0 * (^(p0))) {
            case 9223372036854775807:
                if len(v22) >= 2 {
                    v22 = v22[1:len(v22)]
                }
            case 1:
                v24 := GenS3{F0: uint(v21), F1: v21, F2: 4294967295}
                _ = v24
            case 32767:
                var v25 string = "foo bar"
                _ = v25
            default:
                p1 += (p1 * 32767)
            }
        case 32767:
            v26 := GenS4{F0: p0}
            _ = v26
        default:
            v22 = append(v22, -2147483648)
        }
        var v27 uint64 = v21
        _ = v27
    } else {
        switch "foo bar" {
        case "\n\t":
            switch ("\n\t" + "foo bar") {
            case "":
                var v28 int64 = 0
                _ = v28
            case "hello":
                if v23 {
                    if len(v22) > 0 {
                        v22[0] = (-(p0))
                    }
                    for fi29 := 0; fi29 < 19; fi29++ {
                        fmt.Println(p1, fi29)
                        var v30 int = (fi29 & (fi29 / (((fi29 & p0)) | 1)))
                        _ = fi29
                        _ = v30
                    }
                    v31 := GenS5{F0: (p0 * p0), F1: -128, F2: 42}
                    _ = v31
                }
            default:
                var v32 int32 = 2147483647
                _ = v32
            }
        default:
            p0 -= p0
        }
        for fi33 := 0; fi33 < 6; fi33++ {
            v34 := GenS4{F0: fi33}
            var v35 float32 = 1.5
            _ = fi33
            _ = v34
            _ = v35
        }
    }
    _ = p0
    _ = p1
    _ = v21
    _ = v22
    _ = v23
    return genFn0(true)
}

func main() {
    var v1 int8 = 127
    var v2 float64 = 0.1
    var v3 int = 2147483647
    if (genFn0(true) == 4294967295) {
        for fi36 := 0; fi36 < 14; fi36++ {
            if false {
                fmt.Println(fi36)
                var v37 float32 = 0.1
                fmt.Println(v1)
                _ = v37
            } else {
                if false {
                    v38 := GenS2{F0: 2, F1: -9223372036854775808, F2: v3}
                    _ = v38
                }
                if true {
                    for fi39 := 0; fi39 < 6; fi39++ {
                        switch fi36 {
                        case -2147483648:
                            v40 := []int{(fi39 % ((fi39) | 1)), (v3 % (((v3 >> ((255) & 63))) | 1)), 1}
                            _ = v40
                        case 9223372036854775807:
                            var v41 int32 = 1
                            _ = v41
                        default:
                            v3 = 4294967295
                        }
                        _ = fi39
                    }
                    if false {
                        v42 := []int{(fi36 << ((4294967295) & 63)), 9223372036854775807}
                        v42 = append(v42, (fi36 % ((v3) | 1)))
                        _ = v42
                    } else {
                        var v43 float32 = (0.0 * 0.0)
                        v44 := []int{(v3 << ((2) & 63)), (v3 - 942), (v3 & 9223372036854775807), v3}
                        _ = v43
                        _ = v44
                    }
                } else {
                    v1 = (v1 >> ((65535) & 63))
                    v2 = float64(0.1)
                }
                v45 := []int{0, v3}
                _ = v45
            }
            switch "a" {
            case "\n\t":
                var v46 uint32 = 42
                _ = v46
            case "a":
                if ((true || (!false)) || (411 == 985)) {
                    v3 -= fi36
                    for fi47 := 0; fi47 < 1; fi47++ {
                        var v48 string = ((("" + (((("" + (("\n\t" + ("" + ("a" + "hello"))) + ("foo bar" + "a"))) + "hello") + "\n\t") + "\n\t")) + ("a" + "\n\t")) + ("" + ""))
                        var v49 int32 = 42
                        v50 := []int{(^(v3)), fi47, (fi36 | fi47)}
                        _ = fi47
                        _ = v48
                        _ = v49
                        _ = v50
                    }
                    for fi51 := 0; fi51 < 1; fi51++ {
                        if true {
                            v52 := []int{(v3 / ((fi51) | 1)), fi51, fi36, (v3 >> ((0) & 63))}
                            _ = v52
                        } else {
                            v53 := GenS6{F0: 42, F1: 255}
                            if (3.14159 >= ((3.14159 + 1e10) / (((0.0 / 1e10) - float32(v2)) / -2.25))) {
                                fmt.Println(v3, fi36, v1)
                                v54 := GenS4{F0: fi51}
                                v55 := []int{fi51}
                                _ = v54
                                _ = v55
                            } else {
                                if (309 > 78) {
                                    fmt.Println(v1)
                                    v2 = v2
                                } else {
                                    v56 := []int{}
                                    fmt.Println(fi36, fi51)
                                    if len(v56) >= 2 {
                                        v56 = v56[1:len(v56)]
                                    }
                                    _ = v56
                                }
                            }
                            _ = v53
                        }
                        _ = fi51
                    }
                } else {
                    var v57 float64 = ((((((float64((float32((1.5 * v2)) * 0.0)) * ((float64(1e10) * float64((1e10 + -2.25))) / v2)) + 0.1) * v2) / v2) / (v2 - ((((0.0 / (v2 + v2)) + (v2 * (v2 / 1e10))) + ((1e10 + v2) * (1e10 / v2))) * (v2 * v2)))) - float64(float32(v2)))
                    _ = v57
                }
            }
            _ = fi36
        }
    } else {
        var v58 string = ("" + "\n\t")
        if true {
            for fi59 := 0; fi59 < 19; fi59++ {
                if (true || false) {
                    var v60 uint16 = 112
                    v3 += fi59
                    var v61 int32 = 42
                    _ = v60
                    _ = v61
                } else {
                    if false {
                        v1 -= (v1 << ((65535) & 63))
                        v58 = (((v58 + "") + v58) + "\n\t")
                    }
                    switch v58 {
                    case "a":
                        var v62 int = -1
                        _ = v62
                    }
                }
                var v63 int64 = -194
                v64 := []int{128}
                _ = fi59
                _ = v63
                _ = v64
            }
        }
        v2 = (3.14159 - (-2.25 / v2))
        _ = v58
    }
    for fi65 := 0; fi65 < 19; fi65++ {
        var v66 uint16 = 336
        _ = fi65
        _ = v66
    }
    var v67 uint = 4294967295
    v1 = (^(v1))
    if (!(-2.25 >= 1.5)) {
        v1 = (v1 & (v1 % ((v1) | 1)))
    }
    for fi68 := 0; fi68 < 1; fi68++ {
        v69 := GenS3{F0: v67, F1: 0, F2: 1}
        for fi70 := 0; fi70 < 11; fi70++ {
            v1 -= 1
            var v71 uint8 = 188
            var v72 int64 = -32
            _ = fi70
            _ = v71
            _ = v72
        }
        for fi73 := 0; fi73 < 20; fi73++ {
            if ((v3 & -738) >= v3) {
                var v74 float64 = (float64(0.1) * v2)
                _ = v74
            }
            var v75 uint32 = 0
            _ = fi73
            _ = v75
        }
        _ = fi68
        _ = v69
    }
    v76 := GenS2{F0: v1, F1: -1, F2: (^(v3))}
    if true {
        fmt.Println(v3)
        v2 = 1.5
    } else {
        var v77 string = ("foo bar" + (("" + (("hello" + "\n\t") + "a")) + ("" + "hello")))
        _ = v77
    }
    v3 = -524
    v78 := []int{v3, -63, 42, int(v67)}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v67:", v67)
    fmt.Println("slicelen v78:", len(v78), v78)
    _ = v1
    _ = v2
    _ = v3
    _ = v67
    _ = v76
    _ = v78
}
