package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint32; F1 int8; F2 int8 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 int64; F1 bool; F2 uint16 }
type GenS7 struct { F0 string }
type GenS8 struct { F0 string; F1 uint16; F2 uint32 }
type GenS9 struct { F0 int16; F1 int64; F2 int16 }
type GenS10 struct { F0 int; F1 string; F2 uint }
type GenS11 struct { F0 uint32; F1 int64; F2 uint }
func genFn0(p0 string) uint8 {
    p0 = (p0 + (p0 + "a"))
    fmt.Println(p0)
    v6 := []int{1, 4294967295}
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    for fi7 := 0; fi7 < 9; fi7++ {
        for fi8 := 0; fi8 < 7; fi8++ {
            switch ("" + p0) {
            case "":
                p0 = "a"
            case "hello":
                switch (fi7 & fi8) {
                case 787:
                    var v9 uint8 = 248
                    _ = v9
                case 4294967295:
                    if (p0 > (((p0 + p0) + ((p0 + ((("\n\t" + p0) + p0) + "")) + p0)) + ("hello" + (("\n\t" + ((("foo bar" + p0) + ((p0 + p0) + p0)) + (p0 + ("hello" + p0)))) + p0)))) {
                        v10 := GenS1{F0: false, F1: -9223372036854775808}
                        var v11 uint = 65535
                        var v12 uint16 = 42
                        _ = v10
                        _ = v11
                        _ = v12
                    }
                default:
                    fmt.Println(fi7, p0)
                }
            case "a":
                v13 := GenS2{F0: 42}
                _ = v13
            }
            _ = fi8
        }
        for fi14 := 0; fi14 < 16; fi14++ {
            fmt.Println(fi7)
            _ = len(v6)
            v15 := GenS1{F0: false, F1: -860}
            fmt.Println(p0)
            var v16 uint = 18446744073709551615
            _ = fi14
            _ = v15
            _ = v16
        }
        _ = fi7
    }
    _ = p0
    _ = v6
    return 2
}
func genFn1(p0 uint8, p1 int8, p2 int) uint16 {
    fmt.Println(p1)
    switch "" {
    case "hello":
        switch "a" {
        case "\n\t":
            v17 := []int{(-(p2)), -178, p2, p2}
            _ = v17
        }
    case "foo bar":
        for fi18 := 0; fi18 < 20; fi18++ {
            var v19 bool = (3.14159 != (((0.0 / (float64(0.1) - (1.5 * 3.14159))) * float64(0.0)) * (3.14159 + 1.5)))
            _ = fi18
            _ = v19
        }
    case "":
        v20 := GenS3{F0: 2147483648, F1: 127, F2: -128}
        _ = v20
    default:
        var v21 float64 = (-2.25 - 1.5)
        _ = v21
    }
    if (!false) {
        var v22 float64 = float64((1e10 - 1.5))
        _ = v22
    } else {
        switch 1 {
        case 127:
            if (p2 != p2) {
                v23 := []int{}
                v24 := GenS3{F0: 42, F1: p1, F2: -1}
                var v25 float64 = (float64(0.0) / 0.1)
                _ = v23
                _ = v24
                _ = v25
            }
        case -616:
            if (2 <= 42) {
                if true {
                    for fi26 := 0; fi26 < 7; fi26++ {
                        var v27 int16 = -32768
                        _ = fi26
                        _ = v27
                    }
                } else {
                    var v28 int = -892
                    for fi29 := 0; fi29 < 12; fi29++ {
                        var v30 int64 = 0
                        for fi31 := 0; fi31 < 7; fi31++ {
                            v32 := GenS3{F0: 2, F1: (p1 << ((182) & 63)), F2: 127}
                            var v33 int8 = int8(p0)
                            _ = fi31
                            _ = v32
                            _ = v33
                        }
                        _ = fi29
                        _ = v30
                    }
                    _ = v28
                }
                var v34 int = p2
                var v35 int16 = -32768
                _ = v34
                _ = v35
            }
        case -9223372036854775808:
            p2 -= (p2 | p2)
        default:
            v36 := []int{(p2 % ((132) | 1)), (^(p2))}
            _ = v36
        }
        v37 := []int{}
        if (false || (!false)) {
            if len(v37) > 0 {
                v37[0] = p2
            }
            var v38 uint32 = 1
            _ = v38
        } else {
            var v39 int32 = 790
            p1 -= (p1 - 1)
            for fi40 := 0; fi40 < 8; fi40++ {
                switch (("\n\t" + "hello") + ("" + (("hello" + "hello") + ("foo bar" + ("a" + "\n\t"))))) {
                case "":
                    v41 := GenS3{F0: 2147483647, F1: p1, F2: p1}
                    _ = v41
                case "foo bar":
                    v42 := GenS4{F0: (p0 + p0)}
                    _ = v42
                }
                _ = fi40
            }
            _ = v39
        }
        _ = v37
    }
    _ = p0
    _ = p1
    _ = p2
    return 941
}

func main() {
    var v1 bool = false
    v2 := []uint32{}
    var v3 float32 = 0.0
    var v4 int = 42
    var v5 string = ("" + "a")
    v43 := GenS2{F0: 128}
    v44 := GenS5{F0: (v4 << ((42) & 63))}
    switch v5 {
    case "a":
        v45 := GenS6{F0: 0, F1: v1, F2: 42}
        _ = v45
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    switch v4 {
    case -330:
        v46 := GenS7{F0: (v5 + v5)}
        _ = v46
    case 4294967295:
        v1 = v1
    }
    _ = len(v2)
    if v1 {
        v3 = 0.1
        if ((v4 * (-(v4))) >= v4) {
            v2 = append(v2, 0)
            v5 = (("foo bar" + v5) + "hello")
            switch "foo bar" {
            case "hello":
                var v47 uint = 65535
                _ = v47
            default:
                v3 = v3
            }
        }
        _ = len(v2)
    }
    v4 = 4294967295
    for fi48 := 0; fi48 < 1; fi48++ {
        v2 = append(v2, 65535)
        _ = fi48
    }
    for fi49 := 0; fi49 < 13; fi49++ {
        var v50 int32 = 1
        _ = fi49
        _ = v50
    }
    for fi51 := 0; fi51 < 18; fi51++ {
        for fi52 := 0; fi52 < 12; fi52++ {
            if v1 {
                switch "hello" {
                case "":
                    if len(v2) > 0 {
                        v2[0] = 1
                    }
                case "a":
                    if (0 == 42) {
                        v4 += 32767
                        var v53 int = (fi51 / ((v4) | 1))
                        switch "foo bar" {
                        case "foo bar":
                            v1 = (!(false && v1))
                        }
                        _ = v53
                    } else {
                        var v54 int8 = 1
                        v2 = append(v2, 2147483647)
                        _ = v54
                    }
                default:
                    var v55 int64 = 9223372036854775807
                    _ = v55
                }
                var v56 int64 = 443
                var v57 int16 = 967
                _ = v56
                _ = v57
            }
            v1 = v1
            v58 := GenS8{F0: ("foo bar" + v5), F1: 899, F2: 0}
            _ = fi52
            _ = v58
        }
        _ = fi51
    }
    if (0 == -677) {
        var v59 int64 = -9223372036854775808
        var v60 int = (^(v4))
        if v1 {
            if (v1 && v1) {
                v60 -= v4
            }
            v61 := GenS2{F0: 61}
            fmt.Println(v59)
            _ = v61
        }
        _ = v59
        _ = v60
    } else {
        fmt.Println(v4)
        _ = len(v2)
        var v62 uint64 = 18446744073709551615
        _ = v62
    }
    if ((v4 ^ -153) == (v4 >> ((371) & 63))) {
        if len(v2) > 0 {
            v2[0] = 4294967295
        }
        _ = len(v2)
        for fi63 := 0; fi63 < 6; fi63++ {
            switch (v5 + (v5 + v5)) {
            case "foo bar":
                v64 := GenS8{F0: (v5 + v5), F1: 65535, F2: 4294967295}
                _ = v64
            case "":
                var v65 int8 = 84
                _ = v65
            case "hello":
                v66 := GenS9{F0: 1, F1: -990, F2: 1000}
                _ = v66
            default:
                _ = len(v2)
            }
            switch -955 {
            case 2147483647:
                switch v5 {
                case "foo bar":
                    v67 := GenS10{F0: fi63, F1: (v5 + v5), F2: 191}
                    _ = v67
                default:
                    for fi68 := 0; fi68 < 19; fi68++ {
                        v4 -= fi68
                        var v69 uint16 = 298
                        _ = fi68
                        _ = v69
                    }
                }
            case -163:
                switch -1 {
                case 1:
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                case -710:
                    if len(v2) > 0 {
                        v2[0] = 65535
                    }
                default:
                    v1 = (((false && false) && false) || (v3 <= v3))
                }
            case 42:
                var v70 uint64 = 398
                _ = v70
            }
            _ = fi63
        }
    } else {
        var v71 string = v5
        if (!true) {
            v72 := GenS8{F0: "\n\t", F1: 2, F2: 365}
            _ = v72
        } else {
            v5 = (("\n\t" + (("" + (("hello" + ("\n\t" + v5)) + v71)) + (("" + "foo bar") + v71))) + v71)
            v3 = (v3 * v3)
            v73 := GenS11{F0: 4294967295, F1: -9223372036854775808, F2: 4294967295}
            _ = v73
        }
        for fi74 := 0; fi74 < 17; fi74++ {
            for fi75 := 0; fi75 < 19; fi75++ {
                fmt.Println(v4)
                v71 = v71
                _ = fi75
            }
            var v76 int = (-(v4))
            _ = fi74
            _ = v76
        }
        _ = v71
    }
    switch 4294967295 {
    case 42:
        for fi77 := 0; fi77 < 19; fi77++ {
            v78 := GenS11{F0: 1, F1: 0, F2: 2}
            var v79 uint32 = 2147483647
            v3 -= (v3 * (float32((-2.25 / (3.14159 / (1.5 + (1.5 - (-2.25 / (0.1 * 1.5))))))) * v3))
            _ = fi77
            _ = v78
            _ = v79
        }
    }
    v5 = ("hello" + "")
    if len(v2) > 0 {
        v2[0] = 2
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v5
}
