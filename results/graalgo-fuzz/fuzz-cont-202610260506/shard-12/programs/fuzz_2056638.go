package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16; F2 int8 }
type GenS2 struct { F0 uint32; F1 int; F2 int8 }
type GenS3 struct { F0 int64; F1 int64 }
type GenS4 struct { F0 string; F1 string; F2 uint8 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 int32; F1 uint; F2 uint16 }
type GenS7 struct { F0 uint32; F1 int8; F2 int16 }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 int64; F1 uint8 }
func genFn0(p0 int8, p1 uint, p2 int16) uint32 {
    var v4 uint32 = 65535
    v5 := []int{32767}
    if (int8(p2) == (p0 | (p0 % (((p0 / (((p0 & (p0 & 1))) | 1))) | 1)))) {
        if (("foo bar" + ("a" + "")) <= "hello") {
            p1 = p1
            v4 = v4
        } else {
            fmt.Println(p1)
            v6 := GenS1{F0: (p0 | (^(p0))), F1: 826, F2: -4}
            _ = v6
        }
    } else {
        if (-2147483648 >= -70) {
            switch -830 {
            case 127:
                _ = len(v5)
            case 790:
                fmt.Println(p0, v4, p1)
            case -272:
                if (42 == 4294967295) {
                    v4 = v4
                }
            }
            if (0.0 >= 0.1) {
                switch -1 {
                case 42:
                    var v7 uint16 = 0
                    _ = v7
                default:
                    for fi8 := 0; fi8 < 5; fi8++ {
                        fmt.Println(p2, v4)
                        var v9 int16 = (p2 * -38)
                        _ = fi8
                        _ = v9
                    }
                }
                switch ("\n\t" + ((("foo bar" + "foo bar") + "\n\t") + (("foo bar" + "") + "foo bar"))) {
                case "a":
                    v10 := GenS2{F0: (v4 + v4), F1: 127, F2: p0}
                    _ = v10
                case "\n\t":
                    var v11 uint8 = 2
                    _ = v11
                case "foo bar":
                    var v12 string = "a"
                    _ = v12
                default:
                    var v13 int32 = 436
                    _ = v13
                }
            }
            p0 = 0
        } else {
            var v14 int8 = -128
            v15 := GenS2{F0: v4, F1: 197, F2: (v14 % ((-128) | 1))}
            _ = v14
            _ = v15
        }
        switch "foo bar" {
        case "a":
            switch "foo bar" {
            case "a":
                var v16 int8 = (p0 ^ 0)
                _ = v16
            case "\n\t":
                v4 = (v4 ^ v4)
            default:
                v4 = (^(v4))
            }
        case "hello":
            var v17 float64 = 0.1
            _ = v17
        default:
            if (!false) {
                v5 = append(v5, 32768)
                for fi18 := 0; fi18 < 9; fi18++ {
                    v5 = append(v5, fi18)
                    switch 0 {
                    case 9223372036854775807:
                        for fi19 := 0; fi19 < 7; fi19++ {
                            fmt.Println(v4)
                            var v20 uint64 = 685
                            _ = fi19
                            _ = v20
                        }
                    case 0:
                        switch fi18 {
                        case -2147483648:
                            var v21 int32 = 42
                            _ = v21
                        case 42:
                            switch "" {
                            case "":
                                if false {
                                    switch fi18 {
                                    case 2147483647:
                                        v4 += v4
                                    case 989:
                                        fmt.Println(p1, p2)
                                    }
                                    fmt.Println(p0, p2, p0)
                                }
                            case "\n\t":
                                v22 := GenS2{F0: 2147483648, F1: (fi18 & fi18), F2: 127}
                                _ = v22
                            }
                        case 127:
                            if (false && ((1.5 / 1.0) > 1e10)) {
                                fmt.Println(p0, p0, p1)
                                for fi23 := 0; fi23 < 15; fi23++ {
                                    v24 := GenS3{F0: -1, F1: -1}
                                    fmt.Println(fi23, v4, p0)
                                    if true {
                                        var v25 int8 = p0
                                        fmt.Println(p2, p2, p2)
                                        v4 -= (v4 * v4)
                                        var v26 float32 = float32((1.5 - 1.5))
                                        if len(v5) >= 2 {
                                            v5 = v5[1:len(v5)]
                                        }
                                        _ = v25
                                        _ = v26
                                    }
                                    v4 = v4
                                    p2 += (p2 >> (((p1 & (p1 * p1))) & 63))
                                    p1 = p1
                                    _ = fi23
                                    _ = v24
                                }
                                var v27 bool = false
                                _ = v27
                            }
                        default:
                            if (!(42 >= 2)) {
                                p0 = (-(p0))
                                var v28 uint8 = uint8(p1)
                                v29 := GenS2{F0: 819, F1: fi18, F2: 127}
                                _ = v28
                                _ = v29
                            }
                        }
                    default:
                        fmt.Println(p0)
                    }
                    _ = len(v5)
                    _ = fi18
                }
                if len(v5) > 0 {
                    v5[0] = 2147483647
                }
            } else {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                v30 := GenS4{F0: ("\n\t" + ("hello" + "hello")), F1: "a", F2: 42}
                v31 := GenS4{F0: "hello", F1: "", F2: 2}
                _ = v30
                _ = v31
            }
        }
    }
    switch "foo bar" {
    case "foo bar":
        for fi32 := 0; fi32 < 7; fi32++ {
            var v33 uint16 = 65535
            if ((((p0 < (p0 - p0)) || ((!(fi32 < (^(fi32)))) && (!true))) || false) && true) {
                if (-1 > 1) {
                    p0 = p0
                    var v34 float64 = ((float64(float32(0.1)) * 3.14159) - float64(0.1))
                    _ = v34
                }
                p1 -= 185
            }
            fmt.Println(v33, p0)
            _ = fi32
            _ = v33
        }
    case "":
        var v35 int = 9223372036854775807
        _ = v35
    }
    if len(v5) > 0 {
        v5[0] = -42
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return (v4 ^ v4)
}

func main() {
    var v1 int = 659
    var v2 string = "a"
    var v3 float32 = (3.14159 + (1e10 - 1e10))
    v1 = (v1 << ((42) & 63))
    for fi36 := 0; fi36 < 14; fi36++ {
        switch ("hello" + v2) {
        case "":
            v37 := []int{-9223372036854775808, -991}
            _ = v37
        case "a":
            if false {
                if false {
                    var v38 int16 = -232
                    v39 := GenS2{F0: genFn0(42, 42, -997), F1: (-(fi36)), F2: 0}
                    _ = v38
                    _ = v39
                }
                var v40 int16 = -701
                _ = v40
            } else {
                v41 := GenS5{F0: 18446744073709551615}
                v42 := GenS6{F0: 188, F1: 3, F2: 2}
                _ = v41
                _ = v42
            }
        default:
            if false {
                fmt.Println(v2, v1)
            } else {
                switch (v1 * (fi36 - (-(fi36)))) {
                case 127:
                    switch "hello" {
                    case "hello":
                        fmt.Println(v1)
                    }
                default:
                    fmt.Println(v3)
                }
                v3 += (v3 + v3)
            }
        }
        var v43 uint16 = 1
        _ = fi36
        _ = v43
    }
    if false {
        fmt.Println(v1, v3)
        switch "hello" {
        case "":
            var v44 uint = 42
            _ = v44
        case "\n\t":
            v1 = (v1 & v1)
        }
    } else {
        fmt.Println(v3)
    }
    var v45 uint64 = 2
    switch (^(v1)) {
    case 2147483647:
        v1 = (v1 % ((-2147483648) | 1))
    case 32767:
        v46 := GenS4{F0: v2, F1: (v2 + "a"), F2: 0}
        _ = v46
    }
    v3 = (1e10 + 3.14159)
    v1 = (v1 / ((v1) | 1))
    v2 = ""
    var v47 uint = 18446744073709551615
    switch v1 {
    case 42:
        v47 -= (v47 / (((v47 ^ 1)) | 1))
    case -336:
        for fi48 := 0; fi48 < 5; fi48++ {
            if false {
                if (911 <= 2147483647) {
                    var v49 int16 = 755
                    _ = v49
                } else {
                    var v50 int = (v1 + -2147483648)
                    var v51 uint = (v47 * v47)
                    switch (v2 + "\n\t") {
                    case "hello":
                        var v52 uint = 19
                        _ = v52
                    case "\n\t":
                        v53 := GenS7{F0: 1, F1: 0, F2: -1}
                        _ = v53
                    default:
                        var v54 float64 = 3.14159
                        _ = v54
                    }
                    _ = v50
                    _ = v51
                }
            } else {
                v45 += v45
                var v55 int32 = 152
                if false {
                    switch "a" {
                    case "foo bar":
                        var v56 uint16 = 499
                        _ = v56
                    }
                    switch -9223372036854775808 {
                    case -2147483648:
                        v47 = (v47 ^ (^(v47)))
                    case -1:
                        for fi57 := 0; fi57 < 11; fi57++ {
                            v2 = ""
                            var v58 int64 = -1
                            if (-1 <= -1) {
                                v59 := GenS2{F0: 65535, F1: fi48, F2: int8(v58)}
                                var v60 int32 = v55
                                v2 = "foo bar"
                                _ = v59
                                _ = v60
                            }
                            _ = fi57
                            _ = v58
                        }
                    case 32767:
                        for fi61 := 0; fi61 < 20; fi61++ {
                            fmt.Println(v2, v45)
                            var v62 int = fi61
                            _ = fi61
                            _ = v62
                        }
                    }
                    for fi63 := 0; fi63 < 12; fi63++ {
                        fmt.Println(v3, v55, fi48)
                        if false {
                            v1 = 42
                        } else {
                            fmt.Println(v2, fi48, v47)
                            fmt.Println(v47)
                        }
                        v64 := GenS8{F0: (!false)}
                        v65 := []int{}
                        fmt.Println(v2, fi48, v47)
                        _ = fi63
                        _ = v64
                        _ = v65
                    }
                }
                _ = v55
            }
            for fi66 := 0; fi66 < 2; fi66++ {
                v1 = -24
                v1 += v1
                var v67 int16 = 1000
                _ = fi66
                _ = v67
            }
            _ = fi48
        }
    case 707:
        if true {
            v2 = v2
            switch v1 {
            case 32768:
                v68 := GenS7{F0: 980, F1: 0, F2: -1}
                _ = v68
            default:
                if (-2.25 >= 0.0) {
                    v69 := GenS8{F0: true}
                    _ = v69
                } else {
                    var v70 uint16 = 65535
                    _ = v70
                }
            }
        }
    default:
        if (v3 != 1e10) {
            for fi71 := 0; fi71 < 8; fi71++ {
                v72 := GenS9{F0: -9223372036854775808, F1: 100}
                v73 := []int{v1, -2147483648, -802}
                _ = fi71
                _ = v72
                _ = v73
            }
        }
    }
    v74 := []int{-462, v1, 557, v1}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v45:", v45)
    fmt.Println("v47:", v47)
    fmt.Println("slicelen v74:", len(v74), v74)
    _ = v1
    _ = v2
    _ = v3
    _ = v45
    _ = v47
    _ = v74
}
