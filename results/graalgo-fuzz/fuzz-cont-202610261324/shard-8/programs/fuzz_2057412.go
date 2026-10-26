package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8 }
type GenS2 struct { F0 int; F1 uint32; F2 uint64 }
type GenS3 struct { F0 uint32; F1 uint64; F2 uint16 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 int32; F1 uint8; F2 uint32 }
type GenS8 struct { F0 string }
type GenS9 struct { F0 string; F1 int; F2 int64 }
type GenS10 struct { F0 uint16; F1 string }

func main() {
    var v1 int16 = -858
    var v2 float64 = (-2.25 / (1.5 + 1e10))
    var v3 int32 = -643
    var v4 bool = (0 == -1)
    switch "a" {
    case "hello":
        if v4 {
            v5 := GenS1{F0: 3, F1: 127}
            var v6 uint32 = 2
            fmt.Println(v2, v1)
            _ = v5
            _ = v6
        } else {
            for fi7 := 0; fi7 < 20; fi7++ {
                if v4 {
                    v1 = 1000
                }
                _ = fi7
            }
            var v8 bool = (!(!false))
            fmt.Println(v1, v4)
            _ = v8
        }
    case "a":
        for fi9 := 0; fi9 < 3; fi9++ {
            v10 := GenS1{F0: 127, F1: int8(v3)}
            v11 := []int{}
            v11 = append(v11, fi9)
            _ = fi9
            _ = v10
            _ = v11
        }
    case "foo bar":
        if (42 != 0) {
            switch "a" {
            case "a":
                fmt.Println(v1, v3)
            case "foo bar":
                var v12 uint64 = 0
                _ = v12
            case "\n\t":
                for fi13 := 0; fi13 < 4; fi13++ {
                    switch "\n\t" {
                    case "":
                        var v14 uint32 = 0
                        _ = v14
                    case "a":
                        fmt.Println(v4)
                    case "foo bar":
                        for fi15 := 0; fi15 < 3; fi15++ {
                            var v16 int = fi15
                            fmt.Println(v16, fi13)
                            _ = fi15
                            _ = v16
                        }
                    default:
                        var v17 int8 = -128
                        _ = v17
                    }
                    _ = fi13
                }
            }
            v1 = (v1 * v1)
            fmt.Println(v1, v1)
        }
    default:
        if (42 >= 2) {
            if (!false) {
                var v18 bool = v4
                _ = v18
            } else {
                var v19 uint = 255
                _ = v19
            }
            for fi20 := 0; fi20 < 20; fi20++ {
                fmt.Println(v3)
                switch (fi20 << ((2) & 63)) {
                case 276:
                    for fi21 := 0; fi21 < 7; fi21++ {
                        if v4 {
                            v3 = (-(v3))
                            if v4 {
                                v4 = ((v4 && ((v3 & (v3 * (v3 & (v3 % (((-(v3))) | 1))))) <= v3)) || (-128 < -128))
                                v22 := GenS1{F0: 64, F1: 0}
                                v23 := GenS2{F0: fi21, F1: 798, F2: 51}
                                _ = v22
                                _ = v23
                            }
                        }
                        v4 = v4
                        v4 = (!((255 <= 2) && (42 < 918)))
                        _ = fi21
                    }
                case -9223372036854775808:
                    v24 := GenS3{F0: 4294967295, F1: 18446744073709551615, F2: 255}
                    _ = v24
                case 127:
                    var v25 int8 = 127
                    _ = v25
                default:
                    for fi26 := 0; fi26 < 3; fi26++ {
                        var v27 uint16 = 452
                        _ = fi26
                        _ = v27
                    }
                }
                switch "" {
                case "\n\t":
                    v28 := GenS1{F0: 100, F1: 42}
                    _ = v28
                case "hello":
                    if false {
                        fmt.Println(v4)
                        var v29 float32 = (0.1 + (float32(((v2 + (v2 * float64(0.0))) * v2)) * float32(v2)))
                        v30 := GenS4{F0: 42}
                        _ = v29
                        _ = v30
                    }
                default:
                    v3 = (v3 ^ (v3 << ((805) & 63)))
                }
                _ = fi20
            }
            var v31 uint32 = 1
            _ = v31
        }
    }
    switch -1 {
    case 626:
        fmt.Println(v1, v2, v3)
    case 778:
        v32 := []int{-622}
        _ = v32
    case 32767:
        if (42 != (v3 - v3)) {
            fmt.Println(v2, v3, v1)
        } else {
            v33 := GenS3{F0: 4294967295, F1: 631, F2: 1}
            _ = v33
        }
    }
    for fi34 := 0; fi34 < 11; fi34++ {
        v35 := GenS2{F0: fi34, F1: 2147483647, F2: 604}
        for fi36 := 0; fi36 < 10; fi36++ {
            var v37 int32 = v3
            var v38 int64 = 9223372036854775807
            _ = fi36
            _ = v37
            _ = v38
        }
        switch "hello" {
        case "\n\t":
            v39 := GenS4{F0: 0}
            _ = v39
        case "":
            v40 := GenS2{F0: fi34, F1: 65535, F2: 18446744073709551615}
            _ = v40
        }
        _ = fi34
        _ = v35
    }
    if (v4 || v4) {
        var v41 uint16 = 1
        var v42 uint8 = 31
        v4 = (true && v4)
        _ = v41
        _ = v42
    } else {
        v43 := GenS5{F0: 127}
        v4 = true
        if (float32(v2) != 3.14159) {
            if true {
                switch "hello" {
                case "foo bar":
                    fmt.Println(v2, v2, v2)
                }
                v4 = (0 >= 2)
                if false {
                    v4 = (!v4)
                    v1 = v1
                    v1 += v1
                } else {
                    v44 := GenS2{F0: 9223372036854775807, F1: 42, F2: 0}
                    v45 := GenS4{F0: 4294967295}
                    for fi46 := 0; fi46 < 8; fi46++ {
                        switch "foo bar" {
                        case "":
                            for fi47 := 0; fi47 < 19; fi47++ {
                                v48 := []int{357, -2147483648, 128, fi47}
                                v49 := GenS6{F0: -166}
                                _ = fi47
                                _ = v48
                                _ = v49
                            }
                        case "hello":
                            v4 = v4
                        default:
                            v50 := []int{(fi46 / ((fi46) | 1)), fi46}
                            _ = v50
                        }
                        v51 := []int{fi46, (fi46 >> ((660) & 63)), 128, (fi46 / (((fi46 << ((4294967295) & 63))) | 1))}
                        _ = fi46
                        _ = v51
                    }
                    _ = v44
                    _ = v45
                }
            } else {
                var v52 int64 = 9223372036854775807
                _ = v52
            }
            var v53 int = 4294967295
            _ = v53
        } else {
            if v4 {
                if ((v1 + -323) <= v1) {
                    v54 := []int{}
                    if len(v54) > 0 {
                        v54[0] = -9223372036854775808
                    }
                    _ = v54
                } else {
                    v55 := GenS5{F0: 1}
                    if v4 {
                        switch -823 {
                        case 894:
                            v1 = (v1 - 559)
                        case 127:
                            switch 9223372036854775807 {
                            case 32767:
                                var v56 int64 = -158
                                _ = v56
                            case 1:
                                fmt.Println(v2, v4, v2)
                            }
                        }
                        v2 += v2
                        v57 := []int{-1, 127, 352, 230}
                        _ = v57
                    } else {
                        var v58 uint64 = 2
                        v59 := GenS2{F0: -53, F1: 1, F2: v58}
                        _ = v58
                        _ = v59
                    }
                    _ = v55
                }
                var v60 float32 = 1e10
                v61 := GenS7{F0: (v3 | 1), F1: 0, F2: 65535}
                _ = v60
                _ = v61
            } else {
                v62 := GenS8{F0: "hello"}
                _ = v62
            }
            var v63 uint64 = 42
            _ = v63
        }
        _ = v43
    }
    v64 := GenS4{F0: 18446744073709551615}
    fmt.Println(v2, v2)
    if v4 {
        for fi65 := 0; fi65 < 2; fi65++ {
            for fi66 := 0; fi66 < 16; fi66++ {
                var v67 uint32 = 0
                v68 := []int{(fi66 << ((42) & 63)), fi65}
                _ = len(v68)
                _ = fi66
                _ = v67
                _ = v68
            }
            fmt.Println(v1, v3)
            fmt.Println(fi65, v4)
            _ = fi65
        }
        for fi69 := 0; fi69 < 8; fi69++ {
            switch "" {
            case "\n\t":
                var v70 uint64 = 757
                _ = v70
            }
            _ = fi69
        }
        var v71 uint64 = 563
        _ = v71
    } else {
        v72 := GenS6{F0: -577}
        if (2147483648 < 4294967295) {
            if (1 < 0) {
                v1 -= v1
                var v73 int = -168
                for fi74 := 0; fi74 < 20; fi74++ {
                    var v75 uint64 = 18446744073709551615
                    v76 := GenS9{F0: "a", F1: -383, F2: 9223372036854775807}
                    _ = fi74
                    _ = v75
                    _ = v76
                }
                _ = v73
            } else {
                v77 := []int{4294967295, 1, -785, -9223372036854775808}
                _ = v77
            }
            fmt.Println(v3, v2, v3)
            if (9223372036854775807 != 42) {
                v2 -= v2
                if (!(!v4)) {
                    v78 := GenS10{F0: 255, F1: "hello"}
                    v79 := []int{4294967295, 127, -683}
                    _ = v78
                    _ = v79
                }
            } else {
                var v80 bool = true
                _ = v80
            }
        }
        fmt.Println(v2, v4)
        _ = v72
    }
    var v81 int64 = 0
    v82 := []int{127, 32768}
    var v83 uint64 = 42
    if true {
        for fi84 := 0; fi84 < 5; fi84++ {
            if true {
                switch (^(fi84)) {
                case -862:
                    _ = len(v82)
                }
            } else {
                v81 = (v81 % ((v81) | 1))
            }
            var v85 uint16 = 42
            fmt.Println(v1)
            _ = fi84
            _ = v85
        }
        _ = len(v82)
    } else {
        v4 = ((v1 | (v1 % ((v1) | 1))) < 32767)
    }
    for fi86 := 0; fi86 < 15; fi86++ {
        var v87 uint8 = 1
        var v88 int16 = (v1 << ((1) & 63))
        _ = fi86
        _ = v87
        _ = v88
    }
    switch "" {
    case "":
        for fi89 := 0; fi89 < 4; fi89++ {
            if v4 {
                v90 := GenS9{F0: "hello", F1: (fi89 + fi89), F2: v81}
                v91 := GenS7{F0: (v3 * (^(v3))), F1: 0, F2: 1}
                _ = v90
                _ = v91
            }
            var v92 bool = (!false)
            _ = fi89
            _ = v92
        }
    case "foo bar":
        v83 += 18446744073709551615
    case "a":
        v4 = true
    default:
        if len(v82) >= 2 {
            v82 = v82[1:len(v82)]
        }
    }
    if (int8(v3) > 0) {
        v81 = 0
        fmt.Println(v83)
        v4 = (!v4)
    } else {
        var v93 int32 = 42
        _ = v93
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v81:", v81)
    fmt.Println("slicelen v82:", len(v82), v82)
    fmt.Println("v83:", v83)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v64
    _ = v81
    _ = v82
    _ = v83
}
