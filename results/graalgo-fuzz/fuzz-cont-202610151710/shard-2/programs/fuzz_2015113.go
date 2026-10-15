package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16; F2 int64 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int32; F1 uint16; F2 uint64 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint32; F1 uint64; F2 uint64 }
type GenS6 struct { F0 int; F1 int; F2 int16 }
type GenS7 struct { F0 int16; F1 int8 }
type GenS8 struct { F0 uint64 }
func genFn0(p0 uint16, p1 uint32, p2 int16) int64 {
    p0 -= p0
    for fi4 := 0; fi4 < 18; fi4++ {
        fmt.Println(p1, p0, fi4)
        p2 = 42
        _ = fi4
    }
    fmt.Println(p0, p2, p2)
    var v5 uint16 = (^(p0))
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 42
}
func genFn1(p0 int64, p1 int64) uint32 {
    if (true || false) {
        var v6 uint64 = 0
        switch 127 {
        case 1:
            if (false && true) {
                v7 := GenS1{F0: genFn0(255, 645, 32767), F1: 1, F2: (p0 - genFn0(1, 2147483647, 1))}
                v8 := GenS2{F0: 149}
                switch "a" {
                case "\n\t":
                    v9 := GenS1{F0: (p1 + (p1 % ((-9223372036854775808) | 1))), F1: 65535, F2: p0}
                    _ = v9
                case "foo bar":
                    var v10 uint16 = 2
                    _ = v10
                case "":
                    p1 += genFn0(706, 893, -32768)
                default:
                    if true {
                        switch 32768 {
                        case 9223372036854775807:
                            for fi11 := 0; fi11 < 14; fi11++ {
                                var v12 int8 = -59
                                var v13 int16 = 1
                                var v14 string = ((("" + (("\n\t" + "a") + ((("hello" + "foo bar") + ("foo bar" + "a")) + (("\n\t" + "hello") + "a")))) + ("hello" + "foo bar")) + "hello")
                                _ = fi11
                                _ = v12
                                _ = v13
                                _ = v14
                            }
                        case 42:
                            switch 32767 {
                            case -9223372036854775808:
                                for fi15 := 0; fi15 < 15; fi15++ {
                                    var v16 uint64 = (v6 - (v6 | 42))
                                    if false {
                                        for fi17 := 0; fi17 < 16; fi17++ {
                                            v16 = (^(v6))
                                            _ = fi17
                                        }
                                    }
                                    _ = fi15
                                    _ = v16
                                }
                            default:
                                v18 := []int{32767, 4294967295, 4294967295}
                                _ = v18
                            }
                        case -9223372036854775808:
                            v6 -= (v6 ^ (v6 % ((v6) | 1)))
                        default:
                            fmt.Println(v6)
                        }
                    } else {
                        v19 := []int{}
                        _ = v19
                    }
                }
                _ = v7
                _ = v8
            } else {
                p0 = (p0 << ((255) & 63))
                if true {
                    for fi20 := 0; fi20 < 14; fi20++ {
                        switch fi20 {
                        case 32767:
                            var v21 uint = 255
                            _ = v21
                        case 2147483647:
                            var v22 float32 = (3.14159 + 0.0)
                            _ = v22
                        case 0:
                            v23 := GenS1{F0: p0, F1: 65535, F2: p1}
                            _ = v23
                        default:
                            if (275 < v6) {
                                switch "foo bar" {
                                case "foo bar":
                                    fmt.Println(fi20, fi20)
                                }
                                for fi24 := 0; fi24 < 8; fi24++ {
                                    p0 = p0
                                    _ = fi24
                                }
                                switch (fi20 | fi20) {
                                case -596:
                                    v25 := GenS2{F0: 1}
                                    _ = v25
                                case 32767:
                                    p1 = p1
                                case 1:
                                    if (true && true) {
                                        for fi26 := 0; fi26 < 11; fi26++ {
                                            p0 -= (p0 << ((0) & 63))
                                            v6 -= v6
                                            _ = fi26
                                        }
                                        if (!((58 != 0) || (0 == 2147483647))) {
                                            if true {
                                                fmt.Println(v6, v6)
                                                if false {
                                                    p0 += 42
                                                    var v27 int32 = -1
                                                    _ = v27
                                                }
                                                v28 := []int{127, fi20, fi20, 8}
                                                _ = v28
                                            }
                                            var v29 int = -595
                                            _ = v29
                                        }
                                        switch (fi20 * fi20) {
                                        case -2147483648:
                                            var v30 uint8 = 127
                                            _ = v30
                                        }
                                    } else {
                                        v31 := []int{fi20}
                                        for fi32 := 0; fi32 < 4; fi32++ {
                                            p0 += (-(p0))
                                            var v33 float32 = 1.5
                                            if (0.0 == 0.1) {
                                                p1 += (^(p1))
                                            }
                                            _ = fi32
                                            _ = v33
                                        }
                                        var v34 uint = 2
                                        _ = v31
                                        _ = v34
                                    }
                                }
                            }
                        }
                        for fi35 := 0; fi35 < 12; fi35++ {
                            var v36 int16 = 0
                            _ = fi35
                            _ = v36
                        }
                        p0 += (p1 * (p1 >> ((409) & 63)))
                        _ = fi20
                    }
                    if false {
                        var v37 int8 = 1
                        _ = v37
                    }
                    if true {
                        switch -2147483648 {
                        case 127:
                            v38 := GenS3{F0: 1, F1: 1, F2: 24}
                            _ = v38
                        case 42:
                            switch 480 {
                            case 128:
                                if false {
                                    if (true && true) {
                                        p1 = genFn0(1, 4294967295, 1)
                                    } else {
                                        for fi39 := 0; fi39 < 18; fi39++ {
                                            if (!false) {
                                                var v40 uint32 = 65535
                                                var v41 string = "\n\t"
                                                _ = v40
                                                _ = v41
                                            } else {
                                                switch "foo bar" {
                                                case "a":
                                                    v6 += (v6 << ((929) & 63))
                                                }
                                                if (false || (32767 >= 1000)) {
                                                    for fi42 := 0; fi42 < 5; fi42++ {
                                                        p1 = p0
                                                        _ = fi42
                                                    }
                                                    v43 := GenS4{F0: 3}
                                                    v44 := GenS5{F0: 893, F1: (v6 + (v6 % ((v6) | 1))), F2: (v6 & v6)}
                                                    _ = v43
                                                    _ = v44
                                                } else {
                                                    switch "a" {
                                                    case "":
                                                        v45 := GenS6{F0: fi39, F1: 128, F2: 42}
                                                        _ = v45
                                                    case "foo bar":
                                                        v46 := []int{}
                                                        _ = v46
                                                    }
                                                    v47 := []int{(^(fi39)), 32768}
                                                    _ = v47
                                                }
                                            }
                                            p0 -= -9223372036854775808
                                            switch (((("" + "\n\t") + "a") + ("hello" + ((("a" + ("" + (("hello" + "a") + ("foo bar" + ((("a" + "\n\t") + ("hello" + "\n\t")) + "foo bar"))))) + (((("" + "\n\t") + "") + "hello") + "hello")) + "a"))) + ("foo bar" + ("a" + "a"))) {
                                            case "a":
                                                var v48 int32 = -129
                                                _ = v48
                                            }
                                            _ = fi39
                                        }
                                    }
                                } else {
                                    var v49 int64 = (p0 * genFn0(42, 0, -1))
                                    v50 := GenS3{F0: -587, F1: 913, F2: (v6 & 1)}
                                    p0 = (v49 ^ p1)
                                    _ = v49
                                    _ = v50
                                }
                            case 0:
                                var v51 uint16 = 65535
                                _ = v51
                            }
                        default:
                            if true {
                                p0 += p1
                                for fi52 := 0; fi52 < 9; fi52++ {
                                    for fi53 := 0; fi53 < 6; fi53++ {
                                        v54 := []int{}
                                        if (!((fi53 >> ((65535) & 63)) < (fi53 >> ((2) & 63)))) {
                                            if true {
                                                p0 = 42
                                            }
                                        } else {
                                            if (false || ((!(("hello" <= ("hello" + "")) && (0 < 1))) && true)) {
                                                if false {
                                                    v6 = (v6 & v6)
                                                    v55 := GenS7{F0: -429, F1: 0}
                                                    _ = v55
                                                }
                                            }
                                            fmt.Println(fi52, fi53, fi53)
                                            v56 := GenS4{F0: -128}
                                            _ = v56
                                        }
                                        v57 := GenS7{F0: -666, F1: 100}
                                        if len(v54) > 0 {
                                            v54[0] = fi53
                                        }
                                        fmt.Println(p0, fi52)
                                        _ = fi53
                                        _ = v54
                                        _ = v57
                                    }
                                    v6 = 1
                                    p0 = p0
                                    _ = fi52
                                }
                            }
                        }
                        v6 += (v6 / ((v6) | 1))
                    } else {
                        var v58 bool = (!(0 != 1000))
                        _ = v58
                    }
                }
            }
        }
        fmt.Println(p1)
        _ = v6
    } else {
        if (251 < 9) {
            p0 = (p0 - genFn0(32767, 1, 42))
            p0 = p1
            v59 := GenS1{F0: genFn0(32767, 137, -32768), F1: 0, F2: p1}
            _ = v59
        } else {
            var v60 float64 = 3.14159
            _ = v60
        }
        var v61 uint = 18446744073709551615
        _ = v61
    }
    if (255 >= 0) {
        for fi62 := 0; fi62 < 15; fi62++ {
            for fi63 := 0; fi63 < 11; fi63++ {
                var v64 bool = (81 != fi62)
                _ = fi63
                _ = v64
            }
            for fi65 := 0; fi65 < 3; fi65++ {
                switch fi65 {
                case 297:
                    var v66 int64 = p1
                    _ = v66
                }
                _ = fi65
            }
            for fi67 := 0; fi67 < 10; fi67++ {
                var v68 int32 = -127
                var v69 uint8 = 255
                fmt.Println(fi62, fi67)
                _ = fi67
                _ = v68
                _ = v69
            }
            _ = fi62
        }
    }
    v70 := []int{-1, 5, 2147483647}
    p1 += genFn0(32767, 2147483648, 722)
    v71 := GenS2{F0: 128}
    _ = p0
    _ = p1
    _ = v70
    _ = v71
    return 42
}

func main() {
    var v1 float64 = 1.5
    v2 := []string{}
    var v3 uint = 127
    v1 = 1e10
    if (!false) {
        v3 += v3
    }
    if (((!((2147483648 > 2) || (-182 >= -40))) && true) || (((!false) && (v3 != (v3 * 2))) && false)) {
        var v72 uint = (v3 - 1)
        _ = v72
    }
    for fi73 := 0; fi73 < 15; fi73++ {
        v74 := GenS8{F0: 17}
        if true {
            var v75 uint = (^(v3))
            _ = v75
        } else {
            v3 -= (v3 / ((327) | 1))
            v76 := GenS5{F0: 0, F1: 0, F2: 43}
            _ = v76
        }
        for fi77 := 0; fi77 < 8; fi77++ {
            var v78 int8 = 1
            v79 := GenS6{F0: (-(fi77)), F1: (fi73 << ((2) & 63)), F2: -1}
            fmt.Println(v3)
            _ = fi77
            _ = v78
            _ = v79
        }
        _ = fi73
        _ = v74
    }
    v1 = -2.25
    v80 := GenS4{F0: -128}
    v81 := GenS4{F0: -1}
    switch -2147483648 {
    case 0:
        _ = len(v2)
    case -2147483648:
        v3 += v3
    }
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v3
    _ = v80
    _ = v81
}
