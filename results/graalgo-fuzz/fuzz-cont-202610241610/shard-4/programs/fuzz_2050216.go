package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint16; F1 uint; F2 uint16 }
type GenS3 struct { F0 int16; F1 uint }
type GenS4 struct { F0 int16; F1 int8; F2 uint }
type GenS5 struct { F0 uint64; F1 uint32; F2 uint16 }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 int32; F1 uint }
func genFn0() uint {
    var v6 int8 = 120
    v7 := []int{1, 2147483647}
    for fi8 := 0; fi8 < 17; fi8++ {
        switch fi8 {
        case -734:
            v7 = append(v7, (fi8 % ((fi8) | 1)))
        case -1:
            for fi9 := 0; fi9 < 3; fi9++ {
                _ = len(v7)
                switch fi9 {
                case -2147483648:
                    v10 := GenS1{F0: 18446744073709551615}
                    _ = v10
                case 289:
                    fmt.Println(fi9, fi8)
                case 32767:
                    var v11 uint = 932
                    _ = v11
                }
                var v12 uint16 = 2
                _ = fi9
                _ = v12
            }
        }
        _ = fi8
    }
    switch "" {
    case "":
        var v13 uint8 = 128
        _ = v13
    case "a":
        switch "" {
        case "":
            v6 += (v6 - v6)
        case "hello":
            for fi14 := 0; fi14 < 3; fi14++ {
                for fi15 := 0; fi15 < 13; fi15++ {
                    var v16 uint = 332
                    v7 = append(v7, fi15)
                    v7 = append(v7, fi15)
                    if ((!true) || true) {
                        var v17 uint64 = 0
                        switch "\n\t" {
                        case "foo bar":
                            var v18 bool = false
                            _ = v18
                        case "a":
                            if (1 < 0) {
                                v19 := GenS2{F0: 255, F1: 124, F2: 95}
                                _ = len(v7)
                                _ = v19
                            }
                            if (v6 == -35) {
                                v16 = 65535
                                v16 = (v16 & (v16 + uint(v17)))
                            } else {
                                var v20 uint8 = 42
                                _ = v20
                            }
                        case "":
                            v16 += 1
                        }
                        _ = v17
                    } else {
                        var v21 float32 = (0.1 - (1e10 + 3.14159))
                        var v22 uint8 = 127
                        var v23 int = (fi15 / (((fi15 << ((v16) & 63))) | 1))
                        switch "hello" {
                        case "":
                            var v24 uint64 = 0
                            _ = v24
                        case "foo bar":
                            v25 := GenS2{F0: 65535, F1: v16, F2: 0}
                            _ = v25
                        }
                        _ = v21
                        _ = v22
                        _ = v23
                    }
                    _ = fi15
                    _ = v16
                }
                for fi26 := 0; fi26 < 5; fi26++ {
                    var v27 uint = 42
                    if (!(((42 == v27) && true) || (!true))) {
                        if (!(0 == 1)) {
                            if (float64(3.14159) < 1e10) {
                                var v28 int16 = -32768
                                _ = v28
                            } else {
                                v29 := GenS2{F0: 32767, F1: v27, F2: 461}
                                v6 += (v6 & (v6 % ((v6) | 1)))
                                _ = v29
                            }
                        }
                        fmt.Println(v6, v27)
                    } else {
                        v6 += (v6 + (v6 % (((v6 ^ (v6 | int8(v27)))) | 1)))
                        v6 += (-(v6))
                        var v30 bool = (0 < 128)
                        _ = v30
                    }
                    _ = fi26
                    _ = v27
                }
                fmt.Println(fi14, fi14)
                _ = fi14
            }
        }
    case "\n\t":
        switch 32767 {
        case 32768:
            switch 32768 {
            case 127:
                if true {
                    if ((723 == -32768) && true) {
                        var v31 uint8 = 1
                        v7 = append(v7, 497)
                        _ = v31
                    }
                    _ = len(v7)
                    switch "\n\t" {
                    case "a":
                        var v32 bool = ((!true) && ((!(true || false)) && true))
                        _ = v32
                    case "":
                        if len(v7) > 0 {
                            v7[0] = -9223372036854775808
                        }
                    case "foo bar":
                        var v33 int = 4294967295
                        _ = v33
                    }
                }
            case -918:
                for fi34 := 0; fi34 < 1; fi34++ {
                    var v35 int8 = (v6 & 42)
                    var v36 int8 = 1
                    v37 := GenS1{F0: 0}
                    _ = fi34
                    _ = v35
                    _ = v36
                    _ = v37
                }
            case -411:
                switch "a" {
                case "a":
                    if len(v7) >= 2 {
                        v7 = v7[1:len(v7)]
                    }
                }
            }
        case 2147483647:
            switch -46 {
            case -9223372036854775808:
                for fi38 := 0; fi38 < 1; fi38++ {
                    v7 = append(v7, fi38)
                    switch (fi38 + 0) {
                    case -2147483648:
                        v7 = append(v7, (fi38 & (fi38 - fi38)))
                    }
                    _ = fi38
                }
            case 4294967295:
                var v39 bool = true
                _ = v39
            }
        case 127:
            fmt.Println(v6)
        default:
            var v40 int8 = v6
            _ = v40
        }
    }
    _ = v6
    _ = v7
    return 728
}

func main() {
    v1 := make(map[string]string)
    var v2 uint32 = 0
    var v3 int8 = -3
    v4 := make(map[uint32]string)
    var v5 uint16 = 0
    v41 := GenS3{F0: 350, F1: 4294967295}
    if false {
        switch 571 {
        case -183:
            fmt.Println(v2)
        case 32767:
            fmt.Println(v2, v2)
        }
    }
    v42 := GenS1{F0: 0}
    for fi43 := 0; fi43 < 1; fi43++ {
        for fi44 := 0; fi44 < 3; fi44++ {
            var v45 uint64 = 4294967295
            v46 := []int{-1}
            v5 -= v5
            var v47 int = (^(fi43))
            _ = fi44
            _ = v45
            _ = v46
            _ = v47
        }
        _ = fi43
    }
    v48 := GenS4{F0: 42, F1: v3, F2: 4294967295}
    _ = len(v1)
    if false {
        var v49 int = 127
        fmt.Println(v3, v2, v5)
        _ = v49
    } else {
        for fi50 := 0; fi50 < 17; fi50++ {
            switch (fi50 / ((fi50) | 1)) {
            case -193:
                for fi51 := 0; fi51 < 15; fi51++ {
                    fmt.Println(fi50, fi51, v3)
                    if (!true) {
                        v52 := GenS5{F0: 18446744073709551615, F1: v2, F2: v5}
                        v2 = (v2 | (v2 ^ (^(v2))))
                        v3 -= (v3 / (((v3 << ((4294967295) & 63))) | 1))
                        var v53 bool = false
                        fmt.Println(v3, fi50, v53)
                        if ((0.0 != ((float32((((1.5 + (-2.25 / (float64(1.5) / (1e10 + 0.0)))) + (1.5 / 1.0)) / 1.0)) * (1e10 - (3.14159 * (1e10 - float32((1.5 + -2.25)))))) * (1e10 / 3.14159))) || v53) {
                            var v54 bool = v53
                            if ("" != ("foo bar" + "hello")) {
                                var v55 float32 = 1e10
                                if v54 {
                                    v56 := []int{fi51}
                                    var v57 uint32 = (v2 - (v2 + (^(v2))))
                                    var v58 uint = 0
                                    _ = v56
                                    _ = v57
                                    _ = v58
                                } else {
                                    v59 := GenS2{F0: (v5 * 42), F1: 0, F2: 0}
                                    if (false && (-72 < 1000)) {
                                        fmt.Println(fi50, v55, fi50)
                                        switch "foo bar" {
                                        case "hello":
                                            var v60 int8 = v3
                                            _ = v60
                                        case "foo bar":
                                            v55 = v55
                                        case "a":
                                            v61 := v4[v2]
                                            _ = v61
                                        default:
                                            v62 := GenS4{F0: -431, F1: (v3 ^ (v3 & v3)), F2: 517}
                                            _ = v62
                                        }
                                    } else {
                                        fmt.Println(v54)
                                        switch (^(fi50)) {
                                        case 127:
                                            var v63 uint = genFn0()
                                            _ = v63
                                        case 1:
                                            if v54 {
                                                switch 128 {
                                                case 4294967295:
                                                    v3 = 127
                                                }
                                            }
                                        case 4294967295:
                                            if true {
                                                v64 := []int{}
                                                v2 += v2
                                                _ = v64
                                            } else {
                                                switch 127 {
                                                case 127:
                                                    var v65 int = (^(fi51))
                                                    _ = v65
                                                }
                                                v55 -= v55
                                            }
                                        }
                                    }
                                    _ = v59
                                }
                                _ = v55
                            }
                            _ = v54
                        } else {
                            var v66 int8 = v3
                            _ = v66
                        }
                        _ = v52
                        _ = v53
                    }
                    _ = fi51
                }
            default:
                v3 -= v3
            }
            for fi67 := 0; fi67 < 3; fi67++ {
                v68 := GenS6{F0: -257}
                _ = fi67
                _ = v68
            }
            _ = fi50
        }
    }
    for fi69 := 0; fi69 < 9; fi69++ {
        v5 = 255
        var v70 int8 = (v3 % ((v3) | 1))
        _ = fi69
        _ = v70
    }
    var v71 int64 = 0
    for fi72 := 0; fi72 < 6; fi72++ {
        var v73 int64 = v71
        _ = fi72
        _ = v73
    }
    if true {
        for fi74 := 0; fi74 < 7; fi74++ {
            v71 -= (v71 + v71)
            for fi75 := 0; fi75 < 5; fi75++ {
                v76 := GenS1{F0: 2}
                _ = fi75
                _ = v76
            }
            _ = fi74
        }
        for fi77 := 0; fi77 < 1; fi77++ {
            switch fi77 {
            case 2147483647:
                for fi78 := 0; fi78 < 6; fi78++ {
                    var v79 bool = ((0.1 < 0.1) || true)
                    _ = fi78
                    _ = v79
                }
            case 32768:
                v3 += v3
            case -651:
                fmt.Println(v2, v2)
            default:
                if (!true) {
                    v80 := GenS7{F0: 0}
                    for fi81 := 0; fi81 < 19; fi81++ {
                        v82 := v4[v2]
                        v3 = (v3 / ((100) | 1))
                        var v83 int = 0
                        _ = fi81
                        _ = v82
                        _ = v83
                    }
                    _ = v80
                }
            }
            v2 = v2
            v84 := GenS8{F0: int32(fi77), F1: 255}
            _ = fi77
            _ = v84
        }
        var v85 string = ("" + ("foo bar" + ""))
        _ = v85
    }
    v86 := v4[42]
    fmt.Println(v5, v5)
    v87 := GenS5{F0: 631, F1: 1, F2: v5}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v71:", v71)
    fmt.Println("v86:", v86)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v48
    _ = v5
    _ = v71
    _ = v86
    _ = v87
}
