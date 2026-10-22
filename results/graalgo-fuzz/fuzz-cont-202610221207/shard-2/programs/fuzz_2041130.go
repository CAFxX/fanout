package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 string; F1 uint32 }
type GenS3 struct { F0 int64; F1 int16 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint16; F1 int32 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 bool; F1 uint8; F2 uint }
type GenS8 struct { F0 int64; F1 uint8; F2 int8 }
func genFn0(p0 int64) string {
    p0 = p0
    var v6 uint64 = 0
    if false {
        var v7 bool = (false && true)
        _ = v7
    } else {
        if (v6 <= (^(v6))) {
            for fi8 := 0; fi8 < 16; fi8++ {
                p0 = 42
                _ = fi8
            }
            v9 := GenS1{F0: (p0 % (((p0 & (p0 + (p0 % ((p0) | 1))))) | 1))}
            p0 = (p0 + (p0 & (p0 | 0)))
            _ = v9
        }
        if (42 == 505) {
            var v10 int8 = 1
            for fi11 := 0; fi11 < 1; fi11++ {
                p0 += (p0 + p0)
                v12 := GenS1{F0: p0}
                _ = fi11
                _ = v12
            }
            _ = v10
        }
        v13 := GenS1{F0: (^(p0))}
        _ = v13
    }
    v14 := []int{628, 2147483647}
    v15 := GenS1{F0: 1}
    _ = p0
    _ = v14
    _ = v15
    _ = v6
    return ""
}
func genFn1(p0 uint, p1 uint32) int64 {
    v16 := GenS1{F0: -499}
    var v17 int64 = 0
    var v18 uint8 = 0
    var v19 uint16 = 577
    switch "foo bar" {
    case "\n\t":
        p1 = p1
    case "foo bar":
        for fi20 := 0; fi20 < 12; fi20++ {
            v18 = (v18 >> (((p0 / (((p0 / ((uint(p1)) | 1))) | 1))) & 63))
            for fi21 := 0; fi21 < 18; fi21++ {
                if (!(18446744073709551615 >= 17)) {
                    if true {
                        v22 := []int{(fi20 * int(v19)), 2147483647}
                        v22 = append(v22, (fi21 ^ fi20))
                        if false {
                            switch "a" {
                            case "a":
                                var v23 uint = (p0 % ((136) | 1))
                                _ = v23
                            }
                            if (795 > 4294967295) {
                                if false {
                                    switch "hello" {
                                    case "hello":
                                        if len(v22) > 0 {
                                            v22[0] = fi20
                                        }
                                    case "foo bar":
                                        var v24 int16 = int16(v18)
                                        _ = v24
                                    }
                                    var v25 uint = p0
                                    _ = v25
                                } else {
                                    if false {
                                        v18 += (v18 >> ((1) & 63))
                                    }
                                    v26 := GenS2{F0: ("hello" + "a"), F1: (p1 % ((p1) | 1))}
                                    var v27 int8 = -55
                                    _ = v26
                                    _ = v27
                                }
                                v28 := GenS3{F0: v17, F1: 1}
                                _ = v28
                            }
                        } else {
                            switch "\n\t" {
                            case "foo bar":
                                switch 42 {
                                case 128:
                                    fmt.Println(p0)
                                case 1:
                                    v29 := GenS4{F0: (false || ("\n\t" > "a"))}
                                    _ = v29
                                case 32767:
                                    if (p0 <= (p0 | (p0 % ((p0) | 1)))) {
                                        p1 = 2
                                        if len(v22) > 0 {
                                            v22[0] = fi20
                                        }
                                        v22 = append(v22, int(p0))
                                        if false {
                                            var v30 int8 = 127
                                            switch fi21 {
                                            case 0:
                                                if len(v22) >= 2 {
                                                    v22 = v22[1:len(v22)]
                                                }
                                            case -974:
                                                p1 = (p1 >> ((p0) & 63))
                                            }
                                            v17 = -227
                                            _ = v30
                                        } else {
                                            fmt.Println(p0, v18, p0)
                                        }
                                    }
                                default:
                                    v19 = (v19 - (v19 | 32767))
                                }
                            case "":
                                fmt.Println(v18)
                            case "hello":
                                if false {
                                    _ = len(v22)
                                    if false {
                                        var v31 uint64 = 243
                                        fmt.Println(fi21, v18, fi20)
                                        _ = v31
                                    } else {
                                        switch 1 {
                                        case 127:
                                            if len(v22) > 0 {
                                                v22[0] = fi21
                                            }
                                        case 9223372036854775807:
                                            if len(v22) > 0 {
                                                v22[0] = (fi20 / ((fi21) | 1))
                                            }
                                        default:
                                            v32 := GenS1{F0: v17}
                                            _ = v32
                                        }
                                        v18 = 0
                                    }
                                }
                                v33 := GenS3{F0: v17, F1: 170}
                                _ = v33
                            }
                            var v34 int16 = 42
                            p0 += (p0 + p0)
                            _ = v34
                        }
                        _ = v22
                    } else {
                        switch "\n\t" {
                        case "\n\t":
                            p1 = (p1 | p1)
                        case "hello":
                            var v35 uint = p0
                            _ = v35
                        case "foo bar":
                            fmt.Println(v19, fi21)
                        }
                    }
                } else {
                    var v36 uint = (p0 * p0)
                    _ = v36
                }
                v37 := GenS3{F0: -231, F1: 802}
                _ = fi21
                _ = v37
            }
            _ = fi20
        }
    }
    _ = p0
    _ = p1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    return v17
}

func main() {
    var v1 int = 4294967295
    var v2 int = v1
    v3 := make(map[int]uint32)
    v4 := []int{(v1 >> ((65535) & 63)), (v1 >> ((42) & 63))}
    var v5 uint32 = 2147483647
    for fi38 := 0; fi38 < 11; fi38++ {
        for fi39 := 0; fi39 < 6; fi39++ {
            var v40 int16 = -842
            _ = fi39
            _ = v40
        }
        _ = fi38
    }
    fmt.Println(v1)
    if ((false || false) || (!true)) {
        var v41 int = v2
        v1 = (^(v41))
        _ = v41
    } else {
        if len(v4) > 0 {
            v4[0] = v1
        }
        v1 = v2
    }
    switch "a" {
    case "hello":
        if true {
            for fi42 := 0; fi42 < 10; fi42++ {
                v1 += (-(v1))
                for fi43 := 0; fi43 < 19; fi43++ {
                    var v44 float64 = -2.25
                    var v45 int = (fi42 | 4294967295)
                    v45 = 786
                    _ = fi43
                    _ = v44
                    _ = v45
                }
                var v46 string = "a"
                _ = fi42
                _ = v46
            }
            var v47 uint64 = 2
            v3[(-(v2))] = v5
            _ = v47
        } else {
            fmt.Println(v5)
            switch ((("hello" + "hello") + "a") + "hello") {
            case "hello":
                var v48 string = ("hello" + "")
                _ = v48
            case "a":
                var v49 int = v1
                _ = v49
            default:
                v2 -= v1
            }
            fmt.Println(v5, v1)
        }
    }
    for fi50 := 0; fi50 < 12; fi50++ {
        v51 := GenS5{F0: 1, F1: -61}
        var v52 int8 = -36
        if (2147483647 <= v5) {
            var v53 uint32 = 2147483647
            var v54 string = "a"
            v1 += v1
            _ = v53
            _ = v54
        }
        _ = fi50
        _ = v51
        _ = v52
    }
    v55 := GenS2{F0: (((("foo bar" + ((("hello" + ((("foo bar" + "a") + "a") + "\n\t")) + "\n\t") + "a")) + "hello") + "hello") + "foo bar"), F1: v5}
    for fi56 := 0; fi56 < 14; fi56++ {
        switch (("a" + "a") + (("a" + "\n\t") + ("a" + "foo bar"))) {
        case "":
            var v57 int = (fi56 >> ((616) & 63))
            _ = v57
        case "foo bar":
            v5 -= 352
        case "a":
            v58 := GenS1{F0: 1}
            _ = v58
        default:
            for fi59 := 0; fi59 < 11; fi59++ {
                v2 = v2
                v2 += (fi59 << ((255) & 63))
                v2 += v1
                _ = fi59
            }
        }
        if len(v4) > 0 {
            v4[0] = fi56
        }
        _ = fi56
    }
    if false {
        var v60 uint32 = 4294967295
        var v61 bool = (42 != 4294967295)
        v62 := GenS3{F0: 42, F1: -1}
        _ = v60
        _ = v61
        _ = v62
    }
    for fi63 := 0; fi63 < 10; fi63++ {
        var v64 float32 = 1.5
        v65 := GenS6{F0: v5}
        v66 := GenS2{F0: (("" + "") + "a"), F1: 65535}
        _ = fi63
        _ = v64
        _ = v65
        _ = v66
    }
    for fi67 := 0; fi67 < 14; fi67++ {
        var v68 int16 = 0
        fmt.Println(v1)
        v69 := GenS3{F0: genFn1(18446744073709551615, v5), F1: v68}
        _ = fi67
        _ = v68
        _ = v69
    }
    v1 += (^(v2))
    if len(v4) > 0 {
        v4[0] = (v2 << ((18446744073709551615) & 63))
    }
    v1 += (v1 % ((-1) | 1))
    var v70 float32 = 0.1
    var v71 bool = true
    for fi72 := 0; fi72 < 18; fi72++ {
        v2 += (^(v2))
        var v73 int8 = 100
        for fi74 := 0; fi74 < 9; fi74++ {
            var v75 string = ""
            if (v71 || (-1 > 74)) {
                v71 = (v70 >= 0.1)
                _ = len(v3)
                if v71 {
                    var v76 int = 9223372036854775807
                    _ = v76
                }
            }
            v73 += (v73 << ((816) & 63))
            _ = fi74
            _ = v75
        }
        _ = fi72
        _ = v73
    }
    if ((v71 && (-1 >= -9223372036854775808)) && v71) {
        for fi77 := 0; fi77 < 17; fi77++ {
            var v78 uint16 = 2
            _ = fi77
            _ = v78
        }
        fmt.Println(v71, v1, v71)
    } else {
        switch -2147483648 {
        case 385:
            if true {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            } else {
                for fi79 := 0; fi79 < 16; fi79++ {
                    switch -9223372036854775808 {
                    case 32768:
                        if v71 {
                            fmt.Println(fi79, v71, v2)
                            v1 -= 9223372036854775807
                            v5 += 42
                        } else {
                            var v80 uint8 = 0
                            var v81 float64 = float64(v70)
                            v1 += (v1 >> ((0) & 63))
                            _ = v80
                            _ = v81
                        }
                    default:
                        v82 := GenS7{F0: false, F1: 255, F2: 2}
                        _ = v82
                    }
                    switch fi79 {
                    case -1:
                        for fi83 := 0; fi83 < 17; fi83++ {
                            switch "hello" {
                            case "hello":
                                v71 = (((0.0 * 1e10) + 1.5) != 0.0)
                            case "foo bar":
                                switch (^(v1)) {
                                case 0:
                                    v70 -= v70
                                case 128:
                                    if (0 > -1) {
                                        v71 = false
                                        v5 = v5
                                        fmt.Println(v1, fi83, v1)
                                        var v84 int16 = 0
                                        _ = v84
                                    }
                                }
                            }
                            _ = fi83
                        }
                    case 4294967295:
                        v4 = append(v4, 32767)
                    case 2147483647:
                        v4 = append(v4, (v1 & (-(v1))))
                    }
                    fmt.Println(v1, v71)
                    _ = fi79
                }
                for fi85 := 0; fi85 < 12; fi85++ {
                    switch ("\n\t" + "foo bar") {
                    case "foo bar":
                        switch fi85 {
                        case 2147483647:
                            v2 += (fi85 + 152)
                        case -2147483648:
                            _ = len(v3)
                        case -1:
                            v86 := GenS3{F0: -1, F1: 0}
                            _ = v86
                        }
                    case "hello":
                        fmt.Println(v2)
                    }
                    v87 := GenS8{F0: 0, F1: 42, F2: 42}
                    _ = fi85
                    _ = v87
                }
                var v88 uint16 = 42
                _ = v88
            }
        }
        var v89 float32 = v70
        _ = v89
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v70:", v70)
    fmt.Println("v71:", v71)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v55
    _ = v70
    _ = v71
}
