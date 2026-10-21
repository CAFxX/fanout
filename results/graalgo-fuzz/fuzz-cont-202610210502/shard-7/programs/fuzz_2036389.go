package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32; F2 bool }
type GenS2 struct { F0 string; F1 int16; F2 uint64 }
type GenS3 struct { F0 int; F1 uint64; F2 uint16 }
type GenS4 struct { F0 bool; F1 string; F2 uint8 }
type GenS5 struct { F0 int; F1 uint8 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 int8; F1 int8; F2 int }
func genFn0(p0 int16, p1 bool) int32 {
    p1 = ((-433 < 2147483647) && p1)
    for fi6 := 0; fi6 < 12; fi6++ {
        if p1 {
            v7 := []int{}
            _ = v7
        } else {
            var v8 int32 = -2147483648
            var v9 uint16 = 42
            _ = v8
            _ = v9
        }
        _ = fi6
    }
    var v10 float32 = 0.1
    fmt.Println(p0, v10)
    for fi11 := 0; fi11 < 14; fi11++ {
        fmt.Println(p0, p0)
        for fi12 := 0; fi12 < 19; fi12++ {
            p1 = (false && ((fi11 + (-(fi11))) > fi11))
            _ = fi12
        }
        _ = fi11
    }
    _ = p0
    _ = p1
    _ = v10
    return -952
}
func genFn1(p0 bool, p1 uint32, p2 uint16) uint64 {
    var v13 int16 = 42
    var v14 uint16 = p2
    v15 := []int{9223372036854775807, 4294967295}
    var v16 float32 = 3.14159
    v15 = append(v15, 9223372036854775807)
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    return 604
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := []string{("\n\t" + ("" + "\n\t")), "\n\t", "foo bar", "hello"}
    v3 := make(map[int]int)
    var v4 int32 = 861
    var v5 uint8 = 255
    v4 -= 0
    if true {
        if true {
            var v17 uint16 = 1
            _ = v17
        } else {
            switch -2147483648 {
            case 2147483647:
                var v18 uint8 = v5
                _ = v18
            case 32767:
                if (!(-568 > -32768)) {
                    if true {
                        var v19 uint16 = 0
                        v20 := GenS1{F0: 18446744073709551615, F1: 2147483647, F2: ((255 <= 0) && (100 >= -75))}
                        _ = v19
                        _ = v20
                    } else {
                        v5 = v5
                        if true {
                            var v21 string = ("hello" + "")
                            var v22 int64 = 42
                            v23 := GenS2{F0: v21, F1: 1000, F2: 4294967295}
                            _ = v21
                            _ = v22
                            _ = v23
                        } else {
                            fmt.Println(v5)
                        }
                    }
                }
            case 42:
                if ((3.14159 / 1.5) < 0.0) {
                    var v24 float32 = 3.14159
                    if true {
                        switch -2147483648 {
                        case -817:
                            fmt.Println(v4, v5, v24)
                        case -2147483648:
                            if (true || true) {
                                var v25 uint32 = 2
                                _ = len(v3)
                                v5 += 42
                                _ = v25
                            }
                        case 121:
                            _ = len(v2)
                        }
                        fmt.Println(v5, v5, v4)
                        var v26 string = "\n\t"
                        _ = v26
                    }
                    fmt.Println(v4, v5)
                    _ = v24
                }
            }
            var v27 int32 = 179
            _ = v27
        }
        switch "" {
        case "foo bar":
            switch "\n\t" {
            case "hello":
                v5 += v5
            default:
                if (287 > 198) {
                    v28 := GenS3{F0: 0, F1: 565, F2: 0}
                    if true {
                        var v29 uint = 0
                        _ = v29
                    }
                    if false {
                        v30 := v3[4294967295]
                        if true {
                            var v31 uint64 = genFn1(false, 2147483648, 255)
                            _ = v31
                        } else {
                            if (!((32767 >= 255) || true)) {
                                var v32 uint = 1
                                _ = v32
                            } else {
                                if true {
                                    var v33 uint8 = v5
                                    v33 = (v33 | v5)
                                    _ = v33
                                }
                            }
                        }
                        _ = v30
                    }
                    _ = v28
                }
            }
        case "\n\t":
            if false {
                v2 = append(v2, "hello")
                v4 = v4
            } else {
                switch "\n\t" {
                case "":
                    var v34 bool = true
                    _ = v34
                case "hello":
                    v5 = (v5 | v5)
                }
            }
        default:
            for fi35 := 0; fi35 < 19; fi35++ {
                if (42 >= 0) {
                    var v36 uint16 = 2
                    v36 += v36
                    _ = len(v2)
                    _ = v36
                }
                _ = fi35
            }
        }
        v37 := GenS2{F0: ("hello" + "foo bar"), F1: -830, F2: 644}
        _ = v37
    } else {
        v38 := GenS4{F0: false, F1: "a", F2: v5}
        v2 = append(v2, ("" + "\n\t"))
        _ = v38
    }
    for fi39 := 0; fi39 < 20; fi39++ {
        if len(v2) > 0 {
            v2[0] = (("hello" + ((("\n\t" + "hello") + "") + "")) + "foo bar")
        }
        if false {
            if true {
                if false {
                    v40 := GenS5{F0: fi39, F1: v5}
                    _ = v40
                }
                v41 := GenS5{F0: 1, F1: 127}
                _ = v41
            }
            for fi42 := 0; fi42 < 7; fi42++ {
                switch "\n\t" {
                case "a":
                    v5 -= v5
                case "\n\t":
                    fmt.Println(v5, fi42)
                default:
                    if (!(true && (2 < 496))) {
                        fmt.Println(fi39, v5, v5)
                        fmt.Println(v5, fi42, fi42)
                    }
                }
                v4 += (v4 | (v4 % ((v4) | 1)))
                v4 += v4
                _ = fi42
            }
        } else {
            if true {
                v5 = 128
            } else {
                var v43 uint16 = 42
                v4 = genFn0(-530, true)
                if true {
                    v43 -= v43
                }
                _ = v43
            }
            for fi44 := 0; fi44 < 16; fi44++ {
                var v45 bool = true
                v4 = (^(v4))
                _ = fi44
                _ = v45
            }
            v2 = append(v2, "")
        }
        _ = fi39
    }
    var v46 uint32 = 4294967295
    v47 := GenS5{F0: -708, F1: 70}
    v46 += v46
    if true {
        var v48 uint64 = 458
        var v49 bool = (-29 < 127)
        v49 = v49
        _ = v48
        _ = v49
    } else {
        var v50 int8 = 42
        v51 := GenS6{F0: 32767}
        _ = v50
        _ = v51
    }
    var v52 int8 = 100
    var v53 float64 = (1e10 * (1e10 / 1.0))
    if (v53 > ((1.5 + v53) + (v53 * float64(3.14159)))) {
        _ = len(v3)
        v46 = v46
    } else {
        fmt.Println(v46, v53, v5)
    }
    if (!(false && true)) {
        if ((335 > -616) || true) {
            var v54 int8 = v52
            _ = v54
        } else {
            for fi55 := 0; fi55 < 3; fi55++ {
                v52 = v52
                _ = fi55
            }
            for fi56 := 0; fi56 < 19; fi56++ {
                var v57 uint8 = v5
                if (1 <= v46) {
                    v52 -= v52
                    var v58 uint64 = 714
                    switch fi56 {
                    case 42:
                        fmt.Println(v46, v52)
                    default:
                        v59 := GenS7{F0: v58}
                        _ = v59
                    }
                    _ = v58
                } else {
                    var v60 int64 = 0
                    _ = v60
                }
                _ = fi56
                _ = v57
            }
        }
        for fi61 := 0; fi61 < 15; fi61++ {
            var v62 bool = (42 == 0)
            switch "hello" {
            case "hello":
                if (56 > v46) {
                    v46 = v46
                } else {
                    switch (fi61 % ((4294967295) | 1)) {
                    case -201:
                        var v63 uint16 = 65535
                        _ = v63
                    }
                    fmt.Println(v4, v52)
                }
            default:
                if (!v62) {
                    var v64 float32 = float32(v53)
                    if len(v2) > 0 {
                        v2[0] = "foo bar"
                    }
                    _ = v64
                }
            }
            _ = fi61
            _ = v62
        }
        switch 127 {
        case 0:
            for fi65 := 0; fi65 < 19; fi65++ {
                fmt.Println(v46, fi65, v4)
                var v66 int32 = v4
                _ = fi65
                _ = v66
            }
        case -13:
            v5 = (v5 + (v5 - (v5 * (v5 % (((v5 << ((681) & 63))) | 1)))))
        }
    } else {
        if false {
            v46 += v46
            v53 = v53
            fmt.Println(v46, v52, v46)
        } else {
            v2 = append(v2, "\n\t")
        }
        if (v5 == (v5 << ((uint(v5)) & 63))) {
            for fi67 := 0; fi67 < 8; fi67++ {
                var v68 bool = false
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                for fi69 := 0; fi69 < 15; fi69++ {
                    v70 := GenS5{F0: (fi67 >> ((4294967295) & 63)), F1: (^(v5))}
                    var v71 uint32 = v46
                    _ = fi69
                    _ = v70
                    _ = v71
                }
                _ = fi67
                _ = v68
            }
            switch "hello" {
            case "a":
                for fi72 := 0; fi72 < 2; fi72++ {
                    for fi73 := 0; fi73 < 6; fi73++ {
                        v74 := GenS6{F0: 0}
                        _ = fi73
                        _ = v74
                    }
                    v1[(v46 * (^(v46)))] = false
                    _ = fi72
                }
            default:
                if ((!("hello" == (("foo bar" + "a") + ""))) || true) {
                    v75 := GenS4{F0: (42 != (-(v52))), F1: "a", F2: (v5 & v5)}
                    v53 = v53
                    v4 = -1
                    _ = v75
                }
            }
            for fi76 := 0; fi76 < 7; fi76++ {
                for fi77 := 0; fi77 < 10; fi77++ {
                    v52 = 100
                    _ = fi77
                }
                _ = fi76
            }
        } else {
            if false {
                var v78 int16 = 0
                _ = v78
            }
        }
        v52 -= -1
    }
    var v79 uint16 = 65535
    v80 := GenS5{F0: 128, F1: (v5 & v5)}
    var v81 uint = 2
    var v82 uint = (v81 | v81)
    var v83 int64 = -78
    if (false || (62 < 32768)) {
        var v84 uint8 = 42
        v85 := GenS8{F0: (v52 | v52), F1: (v52 % ((v52) | 1)), F2: -962}
        _ = v84
        _ = v85
    } else {
        v83 += v83
    }
    var v86 uint = 255
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    fmt.Println("v52:", v52)
    fmt.Println("v53:", v53)
    fmt.Println("v79:", v79)
    fmt.Println("v81:", v81)
    fmt.Println("v82:", v82)
    fmt.Println("v83:", v83)
    fmt.Println("v86:", v86)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v47
    _ = v5
    _ = v52
    _ = v53
    _ = v79
    _ = v80
    _ = v81
    _ = v82
    _ = v83
    _ = v86
}
