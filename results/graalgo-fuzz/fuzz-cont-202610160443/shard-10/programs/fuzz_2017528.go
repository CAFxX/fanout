package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64; F2 uint16 }
type GenS2 struct { F0 uint; F1 int32 }
type GenS3 struct { F0 int8; F1 uint16; F2 uint16 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint16; F1 int; F2 bool }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 string }
type GenS8 struct { F0 int32; F1 int64 }
type GenS9 struct { F0 bool }
type GenS10 struct { F0 int32; F1 uint8; F2 uint64 }
func genFn0(p0 int16) int32 {
    for fi7 := 0; fi7 < 11; fi7++ {
        v8 := GenS1{F0: 53, F1: 42, F2: 32767}
        p0 += (p0 / ((p0) | 1))
        for fi9 := 0; fi9 < 13; fi9++ {
            var v10 uint = 4294967295
            v10 = (^(v10))
            _ = fi9
            _ = v10
        }
        _ = fi7
        _ = v8
    }
    v11 := GenS1{F0: 42, F1: 423, F2: 255}
    v12 := GenS1{F0: 42, F1: 9223372036854775807, F2: 2}
    p0 -= (p0 - 42)
    p0 = p0
    _ = p0
    _ = v11
    _ = v12
    return 0
}
func genFn1(p0 int16, p1 int8, p2 bool) uint8 {
    switch 1 {
    case -1:
        var v13 uint32 = 42
        _ = v13
    default:
        for fi14 := 0; fi14 < 17; fi14++ {
            var v15 uint32 = 2
            var v16 uint8 = 2
            var v17 int32 = genFn0(p0)
            _ = fi14
            _ = v15
            _ = v16
            _ = v17
        }
    }
    p0 -= 32767
    if (p0 > 1) {
        p0 += p0
        var v18 uint8 = 128
        _ = v18
    }
    for fi19 := 0; fi19 < 5; fi19++ {
        switch fi19 {
        case 1:
            if ((((-2.25 - ((0.1 + (-2.25 + -2.25)) / (3.14159 * 0.1))) / 1.0) + 3.14159) < -2.25) {
                fmt.Println(fi19, fi19)
            } else {
                p0 += (p0 / ((1000) | 1))
                if false {
                    for fi20 := 0; fi20 < 13; fi20++ {
                        p1 = 127
                        v21 := GenS1{F0: 127, F1: -162, F2: 1}
                        _ = fi20
                        _ = v21
                    }
                }
                switch fi19 {
                case -9223372036854775808:
                    v22 := GenS1{F0: 146, F1: 0, F2: 255}
                    _ = v22
                case 128:
                    v23 := []int{(fi19 >> ((255) & 63))}
                    _ = v23
                }
            }
        case 128:
            for fi24 := 0; fi24 < 10; fi24++ {
                fmt.Println(fi19, p2, p2)
                _ = fi24
            }
        }
        _ = fi19
    }
    _ = p0
    _ = p1
    _ = p2
    return 204
}

func main() {
    var v1 int = 0
    v2 := make(map[uint32]string)
    var v3 uint8 = 207
    var v4 uint8 = (v3 % ((v3) | 1))
    var v5 uint = 0
    var v6 uint16 = 2
    var v25 int32 = genFn0(-1)
    for fi26 := 0; fi26 < 19; fi26++ {
        v6 = (v6 >> ((v5) & 63))
        _ = fi26
    }
    if (32767 >= 1000) {
        v4 += genFn1(720, -14, false)
        v25 = (v25 * -367)
        if (float32(0.0) != 3.14159) {
            var v27 bool = (true || (!true))
            var v28 uint = 2
            if (v27 && false) {
                switch -9223372036854775808 {
                case 2147483647:
                    for fi29 := 0; fi29 < 4; fi29++ {
                        switch v1 {
                        case 127:
                            _ = len(v2)
                        case 42:
                            if false {
                                v30 := GenS2{F0: (v5 & v5), F1: 910}
                                if v27 {
                                    for fi31 := 0; fi31 < 11; fi31++ {
                                        var v32 float32 = 1.5
                                        _ = fi31
                                        _ = v32
                                    }
                                    v33 := GenS2{F0: v28, F1: v25}
                                    v34 := GenS2{F0: v5, F1: (v25 & (v25 + genFn0(-686)))}
                                    _ = v33
                                    _ = v34
                                } else {
                                    v35 := []int{}
                                    switch ((("a" + "\n\t") + ("\n\t" + "\n\t")) + "\n\t") {
                                    case "a":
                                        if (!(v27 && ((v27 || ((fi29 / (((v1 ^ v1)) | 1)) >= fi29)) || (v6 > (v6 % ((v6) | 1)))))) {
                                            _ = len(v35)
                                            for fi36 := 0; fi36 < 7; fi36++ {
                                                v3 = uint8(v5)
                                                v37 := GenS3{F0: -3, F1: v6, F2: v6}
                                                _ = fi36
                                                _ = v37
                                            }
                                            for fi38 := 0; fi38 < 5; fi38++ {
                                                v27 = v27
                                                _ = fi38
                                            }
                                        }
                                    case "foo bar":
                                        v39 := GenS4{F0: ("hello" + "foo bar")}
                                        _ = v39
                                    default:
                                        var v40 float64 = ((1e10 - (1e10 / (float64(3.14159) + 3.14159))) - ((3.14159 - 1e10) * 1e10))
                                        _ = v40
                                    }
                                    fmt.Println(v27, v5, v27)
                                    _ = v35
                                }
                                var v41 int8 = 100
                                _ = v30
                                _ = v41
                            } else {
                                var v42 float64 = 3.14159
                                var v43 uint32 = 42
                                _ = v42
                                _ = v43
                            }
                        case -1:
                            if (!(((false || (!(0.0 <= 3.14159))) && ((^(v4)) != v4)) && ((v25 ^ (v25 / ((v25) | 1))) != 2147483647))) {
                                switch "" {
                                case "\n\t":
                                    delete(v2, 1)
                                case "a":
                                    if (!v27) {
                                        for fi44 := 0; fi44 < 9; fi44++ {
                                            var v45 uint = 0
                                            _ = fi44
                                            _ = v45
                                        }
                                        var v46 uint64 = 218
                                        if (!(!(v25 == int32(v46)))) {
                                            fmt.Println(v25)
                                            if ((!(-33 != 42)) || (genFn1(-32768, 127, v27) < v3)) {
                                                v47 := GenS5{F0: (v6 / ((v6) | 1)), F1: 524, F2: ((v28 != (v5 | v28)) && ((1e10 * 0.1) >= 3.14159))}
                                                delete(v2, 0)
                                                _ = v47
                                            } else {
                                                var v48 uint = (v28 * (v28 % ((v28) | 1)))
                                                _ = v48
                                            }
                                            for fi49 := 0; fi49 < 6; fi49++ {
                                                _ = len(v2)
                                                _ = fi49
                                            }
                                        }
                                        _ = v46
                                    } else {
                                        if v27 {
                                            var v50 uint64 = 1
                                            _ = v50
                                        }
                                    }
                                case "foo bar":
                                    var v51 int16 = 1000
                                    _ = v51
                                }
                            } else {
                                switch ("" + (((("a" + "\n\t") + ("hello" + ("a" + "foo bar"))) + "\n\t") + ("a" + "a"))) {
                                case "foo bar":
                                    v52 := GenS6{F0: (v6 >> ((4294967295) & 63))}
                                    _ = v52
                                case "hello":
                                    var v53 int64 = -1
                                    _ = v53
                                case "a":
                                    if false {
                                        var v54 float32 = 1e10
                                        _ = v54
                                    } else {
                                        for fi55 := 0; fi55 < 13; fi55++ {
                                            switch "" {
                                            case "a":
                                                if (0.0 < -2.25) {
                                                    v56 := []int{(fi55 + v1)}
                                                    _ = v56
                                                } else {
                                                    fmt.Println(v5)
                                                    v3 += genFn1(1000, -128, v27)
                                                    switch 42 {
                                                    case 2147483647:
                                                        if true {
                                                            v28 += 42
                                                        } else {
                                                            v57 := GenS7{F0: "hello"}
                                                            _ = v57
                                                        }
                                                    case 32767:
                                                        var v58 bool = true
                                                        _ = v58
                                                    case 127:
                                                        if v27 {
                                                            v28 = 962
                                                            v3 -= (v4 << (((v5 / ((v28) | 1))) & 63))
                                                            if ((!false) && (248 == 2147483647)) {
                                                                v27 = v27
                                                                v1 = (v1 % (((fi29 >> ((v5) & 63))) | 1))
                                                            } else {
                                                                fmt.Println(fi55)
                                                                var v59 uint16 = 32767
                                                                _ = v59
                                                            }
                                                        }
                                                    }
                                                }
                                            case "\n\t":
                                                var v60 uint = (^(v5))
                                                _ = v60
                                            default:
                                                fmt.Println(v28, fi55)
                                            }
                                            if v27 {
                                                v61 := GenS7{F0: "\n\t"}
                                                var v62 float64 = 3.14159
                                                _ = v61
                                                _ = v62
                                            } else {
                                                var v63 float32 = float32(1.5)
                                                v64 := []int{fi55, -1, 4294967295, fi29}
                                                switch fi55 {
                                                case 9223372036854775807:
                                                    v4 = genFn1(42, 100, true)
                                                case -2147483648:
                                                    var v65 string = ""
                                                    _ = v65
                                                case 2147483647:
                                                    if (4294967295 <= 507) {
                                                        v25 = 0
                                                        fmt.Println(v63, v28, v5)
                                                    }
                                                }
                                                _ = v63
                                                _ = v64
                                            }
                                            _ = fi55
                                        }
                                    }
                                }
                            }
                        }
                        for fi66 := 0; fi66 < 11; fi66++ {
                            switch 128 {
                            case 796:
                                switch "foo bar" {
                                case "hello":
                                    var v67 bool = v27
                                    _ = v67
                                default:
                                    var v68 int32 = v25
                                    fmt.Println(v3, v5, v5)
                                    delete(v2, 546)
                                    _ = v68
                                }
                            }
                            v5 -= v28
                            _ = fi66
                        }
                        var v69 int8 = 100
                        _ = fi29
                        _ = v69
                    }
                }
                v25 = (v25 & v25)
            }
            _ = v27
            _ = v28
        } else {
            var v70 int32 = genFn0(-1)
            _ = v70
        }
    }
    v2[65535] = ("hello" + "")
    if false {
        var v71 uint64 = 2
        v72 := GenS8{F0: (-(v25)), F1: -433}
        switch "hello" {
        case "hello":
            switch "a" {
            case "a":
                v71 = v71
            default:
                v6 = v6
            }
        }
        _ = v71
        _ = v72
    } else {
        if (!true) {
            var v73 uint32 = 42
            _ = v73
        } else {
            if false {
                v74 := GenS9{F0: false}
                var v75 uint16 = (v6 >> ((v5) & 63))
                _ = v74
                _ = v75
            }
        }
    }
    switch "" {
    case "\n\t":
        v5 = (v5 / ((2) | 1))
    case "":
        var v76 int32 = v25
        _ = v76
    }
    v4 += v3
    for fi77 := 0; fi77 < 12; fi77++ {
        var v78 string = "foo bar"
        for fi79 := 0; fi79 < 17; fi79++ {
            switch 4294967295 {
            case -436:
                v80 := GenS10{F0: (v25 ^ v25), F1: v3, F2: 18446744073709551615}
                _ = v80
            case -279:
                fmt.Println(v1, fi79)
            default:
                v81 := GenS1{F0: uint8(v6), F1: -9223372036854775808, F2: v6}
                _ = v81
            }
            fmt.Println(v25)
            _ = fi79
        }
        var v82 uint = v5
        _ = fi77
        _ = v78
        _ = v82
    }
    var v83 uint16 = v6
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v83:", v83)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v83
}
