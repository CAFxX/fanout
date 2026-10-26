package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool; F2 int64 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint16; F1 uint8 }
type GenS4 struct { F0 int32; F1 uint16; F2 int64 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 int16 }
type GenS9 struct { F0 string; F1 uint16; F2 uint64 }
func genFn0() uint {
    switch 128 {
    case -9223372036854775808:
        var v7 bool = false
        _ = v7
    default:
        v8 := GenS1{F0: 1, F1: false, F2: -526}
        _ = v8
    }
    var v9 uint64 = 624
    var v10 uint = uint(v9)
    _ = v10
    _ = v9
    return (v10 | 4294967295)
}
func genFn1(p0 int64, p1 uint64) int8 {
    v11 := []int{42, -682}
    for fi12 := 0; fi12 < 19; fi12++ {
        var v13 int64 = p0
        for fi14 := 0; fi14 < 18; fi14++ {
            if true {
                var v15 string = ("a" + ("\n\t" + "foo bar"))
                fmt.Println(v15, v15, p1)
                switch v15 {
                case "hello":
                    v16 := GenS2{F0: 805}
                    if len(v11) >= 2 {
                        v11 = v11[1:len(v11)]
                    }
                    _ = v16
                }
                _ = v15
            } else {
                fmt.Println(p0)
                v17 := GenS2{F0: -2147483648}
                _ = v17
            }
            _ = fi14
        }
        _ = fi12
        _ = v13
    }
    _ = len(v11)
    if len(v11) > 0 {
        v11[0] = 2147483647
    }
    var v18 uint = 128
    _ = p0
    _ = p1
    _ = v11
    _ = v18
    return 32
}

func main() {
    v1 := []string{}
    v2 := make(map[uint32]bool)
    var v3 string = (("foo bar" + "\n\t") + "\n\t")
    var v4 uint16 = 2
    var v5 bool = true
    var v6 int = 520
    if len(v1) > 0 {
        v1[0] = v3
    }
    v3 = v3
    for fi19 := 0; fi19 < 4; fi19++ {
        var v20 float32 = 0.1
        v20 = v20
        v5 = v5
        _ = fi19
        _ = v20
    }
    v2[441] = false
    fmt.Println(v3, v3)
    switch v3 {
    case "\n\t":
        v5 = (697 >= 1)
    case "foo bar":
        var v21 int64 = 0
        _ = v21
    default:
        v4 = v4
    }
    v3 = v3
    switch (v6 % (((v6 & (v6 * -460))) | 1)) {
    case -738:
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    default:
        for fi22 := 0; fi22 < 6; fi22++ {
            if len(v1) > 0 {
                v1[0] = ("foo bar" + (v3 + v3))
            }
            _ = fi22
        }
    }
    if false {
        fmt.Println(v3, v6)
        if (((-2.25 + 1.5) + -2.25) < 3.14159) {
            _ = len(v2)
            switch v3 {
            case "":
                _ = len(v1)
            case "hello":
                for fi23 := 0; fi23 < 4; fi23++ {
                    _ = len(v2)
                    fmt.Println(v5, v6, v4)
                    v4 += (v4 >> ((65535) & 63))
                    _ = fi23
                }
            }
        }
    } else {
        v5 = ((v6 + v6) < v6)
        for fi24 := 0; fi24 < 16; fi24++ {
            switch 1 {
            case 127:
                if v5 {
                    var v25 uint8 = 128
                    var v26 float32 = float32(-2.25)
                    _ = v25
                    _ = v26
                }
            case -9223372036854775808:
                var v27 int64 = -847
                _ = v27
            case 2147483647:
                for fi28 := 0; fi28 < 3; fi28++ {
                    delete(v2, 65535)
                    var v29 float32 = 0.0
                    v2[21] = true
                    if false {
                        v30 := GenS2{F0: 2147483647}
                        _ = v30
                    }
                    var v31 uint = genFn0()
                    v31 = (v31 - (v31 + (v31 - (v31 << (((v31 - 65535)) & 63)))))
                    fmt.Println(v3)
                    v1 = append(v1, v3)
                    v1 = append(v1, "hello")
                    if false {
                        v5 = ((!("" != "foo bar")) && v5)
                    } else {
                        _ = len(v2)
                        if v5 {
                            v31 += v31
                            v29 = (3.14159 - ((0.0 - v29) + 1.5))
                            v3 = ("hello" + (v3 + ((("" + "hello") + (v3 + v3)) + v3)))
                        }
                        fmt.Println(v31, v31)
                        v32 := GenS3{F0: 32767, F1: 255}
                        _ = v32
                    }
                    fmt.Println(v6)
                    _ = fi28
                    _ = v29
                    _ = v31
                }
            }
            for fi33 := 0; fi33 < 13; fi33++ {
                if v5 {
                    switch (v3 + (v3 + v3)) {
                    case "a":
                        v3 = "hello"
                    default:
                        if (-2.25 <= ((3.14159 * -2.25) + (3.14159 * (3.14159 * 1e10)))) {
                            v4 += v4
                            fmt.Println(v3, fi24, v6)
                            var v34 int64 = -1
                            _ = v34
                        } else {
                            v3 = "hello"
                            var v35 int32 = 116
                            var v36 int16 = -316
                            _ = v35
                            _ = v36
                        }
                    }
                    fmt.Println(v6, v4)
                    v37 := GenS4{F0: 481, F1: v4, F2: -346}
                    _ = v37
                } else {
                    switch v3 {
                    case "":
                        _ = len(v2)
                    case "hello":
                        switch (("\n\t" + v3) + ((v3 + (v3 + v3)) + "a")) {
                        case "hello":
                            switch v3 {
                            case "hello":
                                v3 = v3
                            case "":
                                var v38 uint32 = 1
                                _ = v38
                            }
                        }
                    case "\n\t":
                        v4 = v4
                    }
                }
                _ = fi33
            }
            for fi39 := 0; fi39 < 11; fi39++ {
                v4 -= 42
                _ = len(v2)
                _ = fi39
            }
            _ = fi24
        }
        var v40 int8 = 127
        _ = v40
    }
    switch v3 {
    case "\n\t":
        var v41 float32 = (1e10 * (3.14159 + 3.14159))
        _ = v41
    case "foo bar":
        for fi42 := 0; fi42 < 3; fi42++ {
            v43 := GenS5{F0: -128}
            switch v3 {
            case "foo bar":
                _ = len(v2)
            case "":
                v5 = true
            case "\n\t":
                v5 = (!(65535 < 2))
            default:
                for fi44 := 0; fi44 < 14; fi44++ {
                    var v45 bool = (128 < 42)
                    switch "\n\t" {
                    case "a":
                        if v45 {
                            var v46 string = "a"
                            switch 0 {
                            case 128:
                                delete(v2, 469)
                            case 32767:
                                if false {
                                    v1 = append(v1, "\n\t")
                                } else {
                                    fmt.Println(fi44, v5, fi44)
                                    v3 = "hello"
                                    v47 := GenS6{F0: (^(v4))}
                                    _ = v47
                                }
                            }
                            _ = v46
                        } else {
                            v2[2147483647] = (!v5)
                            var v48 uint32 = 666
                            _ = v48
                        }
                    }
                    _ = fi44
                    _ = v45
                }
            }
            _ = fi42
            _ = v43
        }
    case "hello":
        var v49 uint = 0
        _ = v49
    default:
        var v50 uint16 = (v4 + v4)
        _ = v50
    }
    if (42 == 2) {
        fmt.Println(v5, v5)
        if false {
            v5 = (true || ((v4 % ((65535) | 1)) <= 1))
        } else {
            var v51 int32 = 535
            if false {
                for fi52 := 0; fi52 < 6; fi52++ {
                    fmt.Println(v4)
                    delete(v2, 414)
                    _ = fi52
                }
            }
            _ = v51
        }
    } else {
        switch -2147483648 {
        case -182:
            switch int(v4) {
            case 160:
                switch "foo bar" {
                case "foo bar":
                    v4 = v4
                default:
                    if v5 {
                        if len(v1) > 0 {
                            v1[0] = "foo bar"
                        }
                        for fi53 := 0; fi53 < 3; fi53++ {
                            v54 := GenS1{F0: -89, F1: false, F2: -662}
                            _ = fi53
                            _ = v54
                        }
                    } else {
                        v5 = ((573 > 42) && true)
                    }
                }
            default:
                if (!v5) {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    switch v3 {
                    case "\n\t":
                        v6 = (v6 << ((42) & 63))
                    case "foo bar":
                        var v55 float32 = 0.0
                        _ = v55
                    }
                } else {
                    if (!true) {
                        v4 = v4
                        fmt.Println(v4, v6)
                        v6 += v6
                    }
                    fmt.Println(v5)
                    if v5 {
                        if v5 {
                            v6 = 32768
                            switch 172 {
                            case 127:
                                var v56 uint16 = v4
                                _ = v56
                            case 555:
                                if (v4 >= (^(v4))) {
                                    v4 = v4
                                } else {
                                    v57 := GenS2{F0: 1}
                                    fmt.Println(v3, v6, v5)
                                    var v58 bool = v5
                                    _ = v57
                                    _ = v58
                                }
                            case 482:
                                for fi59 := 0; fi59 < 16; fi59++ {
                                    v60 := GenS7{F0: 100}
                                    for fi61 := 0; fi61 < 4; fi61++ {
                                        switch fi59 {
                                        case -1:
                                            switch v3 {
                                            case "hello":
                                                v2[2147483648] = (!(0 < -1))
                                            case "foo bar":
                                                v3 = (v3 + "a")
                                            default:
                                                var v62 int8 = genFn1(-517, 18446744073709551615)
                                                _ = v62
                                            }
                                            var v63 uint16 = (^(v4))
                                            v5 = (-1 != -128)
                                            _ = v63
                                        }
                                        _ = fi61
                                    }
                                    v4 -= (v4 * (v4 + v4))
                                    _ = fi59
                                    _ = v60
                                }
                            }
                        }
                        var v64 int8 = genFn1(-1, 696)
                        if (v4 > (v4 & 42)) {
                            for fi65 := 0; fi65 < 14; fi65++ {
                                v6 -= 0
                                if v5 {
                                    for fi66 := 0; fi66 < 8; fi66++ {
                                        fmt.Println(fi65)
                                        var v67 uint = 42
                                        _ = fi66
                                        _ = v67
                                    }
                                    switch fi65 {
                                    case 128:
                                        if false {
                                            if (((v5 && v5) || true) || v5) {
                                                switch (fi65 & 32768) {
                                                case 9223372036854775807:
                                                    v5 = v5
                                                default:
                                                    var v68 uint64 = 4294967295
                                                    _ = v68
                                                }
                                                v2[uint32(v4)] = (!v5)
                                                if (166 < 2147483647) {
                                                    switch (fi65 ^ 32767) {
                                                    case 128:
                                                        switch v3 {
                                                        case "":
                                                            fmt.Println(v64, v5, fi65)
                                                        default:
                                                            for fi69 := 0; fi69 < 9; fi69++ {
                                                                v4 = (v4 >> ((915) & 63))
                                                                _ = fi69
                                                            }
                                                        }
                                                    case 1:
                                                        v1 = append(v1, "foo bar")
                                                    }
                                                    v4 += v4
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                }
                                            }
                                            switch (fi65 >> ((2) & 63)) {
                                            case -696:
                                                v70 := GenS3{F0: 65535, F1: 255}
                                                _ = v70
                                            }
                                            var v71 bool = (v3 != v3)
                                            _ = v71
                                        } else {
                                            v64 = (v64 << ((2) & 63))
                                            v72 := GenS8{F0: 626}
                                            var v73 uint32 = 390
                                            _ = v72
                                            _ = v73
                                        }
                                    }
                                }
                                fmt.Println(v5, v5)
                                _ = fi65
                            }
                            if v5 {
                                for fi74 := 0; fi74 < 1; fi74++ {
                                    v6 = (fi74 - 127)
                                    v75 := GenS7{F0: (v64 - (v64 | v64))}
                                    if (!(v5 || ((v6 % (((v6 & (fi74 - 32768))) | 1)) >= (v6 | v6)))) {
                                        v6 = (v6 << ((0) & 63))
                                        fmt.Println(v4)
                                        var v76 int16 = 612
                                        _ = v76
                                    }
                                    _ = fi74
                                    _ = v75
                                }
                                for fi77 := 0; fi77 < 13; fi77++ {
                                    for fi78 := 0; fi78 < 7; fi78++ {
                                        var v79 int = (fi77 & (fi78 >> ((18446744073709551615) & 63)))
                                        v79 = -609
                                        if (!((^(v64)) < v64)) {
                                            var v80 uint64 = 67
                                            var v81 int = v79
                                            _ = v80
                                            _ = v81
                                        } else {
                                            var v82 uint32 = 2147483648
                                            v2[v82] = (((v82 / ((4294967295) | 1)) == (v82 % (((v82 | (^(v82)))) | 1))) || v5)
                                            _ = v82
                                        }
                                        _ = fi78
                                        _ = v79
                                    }
                                    _ = fi77
                                }
                                if true {
                                    v4 += 65535
                                    for fi83 := 0; fi83 < 16; fi83++ {
                                        if len(v1) > 0 {
                                            v1[0] = v3
                                        }
                                        v64 = 42
                                        fmt.Println(v4, fi83, fi83)
                                        _ = fi83
                                    }
                                }
                            }
                            if v5 {
                                var v84 uint8 = 0
                                for fi85 := 0; fi85 < 17; fi85++ {
                                    if v5 {
                                        v84 += v84
                                        for fi86 := 0; fi86 < 20; fi86++ {
                                            fmt.Println(v5, v3, fi86)
                                            fmt.Println(v5, v3, v84)
                                            _ = len(v1)
                                            _ = fi86
                                        }
                                        v5 = v5
                                    }
                                    v5 = v5
                                    if len(v1) > 0 {
                                        v1[0] = v3
                                    }
                                    _ = fi85
                                }
                                v6 -= v6
                                _ = v84
                            } else {
                                v6 = v6
                                if v5 {
                                    var v87 uint = 788
                                    if (v5 && false) {
                                        if false {
                                            v3 = v3
                                            v88 := GenS9{F0: v3, F1: 2, F2: 18446744073709551615}
                                            _ = v88
                                        }
                                    }
                                    v89 := GenS9{F0: ((("a" + "foo bar") + "foo bar") + v3), F1: (v4 & v4), F2: uint64(v4)}
                                    _ = v87
                                    _ = v89
                                }
                                var v90 int32 = 1
                                _ = v90
                            }
                        } else {
                            var v91 uint64 = 934
                            _ = len(v1)
                            _ = v91
                        }
                        _ = v64
                    } else {
                        fmt.Println(v4, v3, v4)
                        var v92 string = (v3 + v3)
                        _ = v92
                    }
                }
            }
        default:
            v3 = v3
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
