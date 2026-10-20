package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16; F2 string }
type GenS2 struct { F0 int32; F1 uint32; F2 string }
type GenS3 struct { F0 string; F1 int8 }
type GenS4 struct { F0 string; F1 int8 }
type GenS5 struct { F0 uint32; F1 int16; F2 int8 }
type GenS6 struct { F0 uint64; F1 int }
type GenS7 struct { F0 uint16; F1 int32 }
type GenS8 struct { F0 int; F1 bool; F2 uint32 }
type GenS9 struct { F0 int; F1 uint }
type GenS10 struct { F0 int }
type GenS11 struct { F0 int16 }
func genFn0() uint16 {
    for fi4 := 0; fi4 < 20; fi4++ {
        v5 := []int{fi4, fi4, (fi4 ^ (fi4 & 0))}
        var v6 uint8 = 2
        v6 = 2
        _ = fi4
        _ = v5
        _ = v6
    }
    var v7 int32 = 253
    for fi8 := 0; fi8 < 5; fi8++ {
        v7 = (v7 * (v7 | -2147483648))
        _ = fi8
    }
    v9 := GenS1{F0: "a", F1: 255, F2: "hello"}
    _ = v7
    _ = v9
    return 0
}
func genFn1(p0 uint16, p1 uint64) uint {
    v10 := []int{32767, 496, 128, 1}
    if len(v10) > 0 {
        v10[0] = -9223372036854775808
    }
    if false {
        var v11 bool = true
        _ = v11
    } else {
        for fi12 := 0; fi12 < 2; fi12++ {
            p0 = (p0 & 8)
            fmt.Println(fi12, p0)
            _ = fi12
        }
        _ = len(v10)
        if false {
            var v13 uint = 366
            var v14 uint64 = 2
            v10 = append(v10, 4294967295)
            _ = v13
            _ = v14
        }
    }
    if true {
        switch 502 {
        case 4294967295:
            if false {
                switch -9223372036854775808 {
                case -9223372036854775808:
                    p0 -= p0
                case -2147483648:
                    v10 = append(v10, -9223372036854775808)
                case 4294967295:
                    var v15 uint = 684
                    _ = v15
                }
                var v16 bool = true
                _ = v16
            } else {
                v10 = append(v10, 2147483647)
            }
        case -9223372036854775808:
            v17 := GenS2{F0: 1, F1: 370, F2: "foo bar"}
            _ = v17
        default:
            fmt.Println(p0, p0)
        }
    }
    var v18 float32 = 3.14159
    _ = p0
    _ = p1
    _ = v10
    _ = v18
    return 1
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 int8 = 113
    var v3 int32 = 42
    v19 := GenS2{F0: (v3 >> ((genFn1(255, 18446744073709551615)) & 63)), F1: 1, F2: ("a" + ("foo bar" + "hello"))}
    switch "\n\t" {
    case "a":
        _ = len(v1)
    }
    var v20 bool = (2 != 117)
    var v21 uint32 = 65535
    v22 := []int{}
    if false {
        v23 := GenS3{F0: "foo bar", F1: 127}
        _ = v23
    }
    v2 = (^(v2))
    var v24 uint16 = 42
    for fi25 := 0; fi25 < 19; fi25++ {
        v3 = (-(v3))
        if false {
            if (("\n\t" + "") <= ("a" + ("hello" + "\n\t"))) {
                var v26 int8 = 127
                for fi27 := 0; fi27 < 7; fi27++ {
                    v28 := GenS2{F0: (-(v3)), F1: (v21 % (((v21 >> ((42) & 63))) | 1)), F2: "\n\t"}
                    if v20 {
                        v29 := GenS3{F0: "hello", F1: (v2 % ((v26) | 1))}
                        _ = v29
                    } else {
                        switch (-(fi25)) {
                        case -87:
                            var v30 float64 = 0.0
                            _ = v30
                        case 32768:
                            var v31 float32 = -2.25
                            _ = v31
                        case 127:
                            fmt.Println(fi25, v2, v21)
                        default:
                            if v20 {
                                if v20 {
                                    v24 -= v24
                                }
                                v32 := GenS2{F0: 2147483647, F1: v21, F2: ("a" + ("hello" + (("a" + "hello") + (("hello" + ("foo bar" + ("\n\t" + ""))) + "\n\t"))))}
                                if false {
                                    if (v3 == (v3 + (-(v3)))) {
                                        switch fi25 {
                                        case 0:
                                            v3 -= (v3 * (v3 | v3))
                                        case 425:
                                            var v33 int64 = -1
                                            _ = v33
                                        }
                                        v2 += (v2 / ((100) | 1))
                                        v24 = (v24 - genFn0())
                                    } else {
                                        v24 = 0
                                    }
                                }
                                _ = v32
                            }
                        }
                    }
                    _ = fi27
                    _ = v28
                }
                if len(v22) > 0 {
                    v22[0] = 127
                }
                _ = v26
            } else {
                var v34 uint8 = 42
                switch "foo bar" {
                case "":
                    for fi35 := 0; fi35 < 16; fi35++ {
                        switch "" {
                        case "foo bar":
                            var v36 int64 = -438
                            _ = v36
                        case "a":
                            var v37 uint64 = 1
                            _ = v37
                        case "\n\t":
                            v1[v21] = v20
                        default:
                            v24 += genFn0()
                        }
                        v24 = genFn0()
                        v20 = (false && false)
                        if v20 {
                            var v38 int8 = 1
                            v2 -= v38
                            _ = v38
                        }
                        v20 = false
                        if (((0.1 + 1.5) * -2.25) >= 3.14159) {
                            fmt.Println(v24)
                            var v39 string = ((("" + ("hello" + "a")) + "") + "a")
                            v20 = (!false)
                            _ = v39
                        } else {
                            _ = len(v1)
                            if false {
                                var v40 string = (("foo bar" + "foo bar") + "\n\t")
                                switch (fi35 | -763) {
                                case 9223372036854775807:
                                    var v41 bool = ((false || v20) || (-1 <= -1))
                                    _ = v41
                                default:
                                    var v42 int16 = 0
                                    _ = v42
                                }
                                _ = v40
                            } else {
                                fmt.Println(v20)
                            }
                            var v43 int64 = 42
                            _ = v43
                        }
                        _ = fi35
                    }
                default:
                    v2 = int8(fi25)
                }
                _ = v34
            }
            var v44 int32 = -958
            _ = v44
        }
        if (!v20) {
            for fi45 := 0; fi45 < 16; fi45++ {
                v46 := GenS4{F0: "\n\t", F1: -112}
                v47 := GenS2{F0: v3, F1: v21, F2: (("\n\t" + ("a" + "")) + (((("\n\t" + "") + "foo bar") + ("foo bar" + ("hello" + ""))) + "a"))}
                v48 := v1[(v21 + 4294967295)]
                fmt.Println(fi25, fi25)
                fmt.Println(fi45, v3)
                _ = fi45
                _ = v46
                _ = v47
                _ = v48
            }
            if (false || v20) {
                _ = len(v22)
                v49 := GenS3{F0: "", F1: (v2 / ((v2) | 1))}
                var v50 int32 = (v3 ^ -1)
                _ = v49
                _ = v50
            } else {
                if v20 {
                    for fi51 := 0; fi51 < 14; fi51++ {
                        v24 = (v24 | (^(v24)))
                        v2 += v2
                        var v52 uint64 = 1
                        switch "a" {
                        case "hello":
                            v53 := GenS5{F0: 2, F1: -1, F2: v2}
                            _ = v53
                        default:
                            v54 := GenS6{F0: v52, F1: 4294967295}
                            _ = v54
                        }
                        var v55 bool = v20
                        _ = fi51
                        _ = v52
                        _ = v55
                    }
                    if (true || v20) {
                        var v56 int64 = -9223372036854775808
                        _ = v56
                    } else {
                        for fi57 := 0; fi57 < 8; fi57++ {
                            var v58 bool = false
                            v2 = (v2 % ((v2) | 1))
                            _ = fi57
                            _ = v58
                        }
                        if ((((0 >= v2) && true) && (fi25 >= fi25)) && v20) {
                            v20 = v20
                        }
                        for fi59 := 0; fi59 < 3; fi59++ {
                            fmt.Println(fi25, v20)
                            v2 -= v2
                            _ = fi59
                        }
                    }
                }
            }
            var v60 uint16 = v24
            _ = v60
        } else {
            if false {
                fmt.Println(v20, v3)
                if ((1e10 / ((3.14159 * float64(0.1)) / (float64(float32((float64(float32(0.1)) + 0.1))) + float64(((1.5 / 3.14159) - ((-2.25 + 0.1) - (1.5 / 1.5))))))) != (0.0 * 1e10)) {
                    if false {
                        var v61 int8 = 42
                        if ((true || (((^(v3)) <= (v3 - v3)) && true)) || false) {
                            v61 = v61
                        }
                        fmt.Println(v20)
                        _ = v61
                    }
                    v62 := GenS3{F0: ("hello" + ("foo bar" + (("" + "\n\t") + "foo bar"))), F1: -114}
                    _ = v62
                } else {
                    if false {
                        var v63 uint32 = (v21 + (v21 & (v21 | (v21 % (((v21 << ((866) & 63))) | 1)))))
                        fmt.Println(v21, v24)
                        _ = v63
                    } else {
                        v64 := GenS7{F0: (v24 << ((255) & 63)), F1: int32(v21)}
                        _ = v64
                    }
                    v20 = ("" > "hello")
                }
                v3 = (v3 >> ((genFn1(v24, 12)) & 63))
            } else {
                for fi65 := 0; fi65 < 14; fi65++ {
                    var v66 uint32 = 2147483648
                    _ = fi65
                    _ = v66
                }
                for fi67 := 0; fi67 < 16; fi67++ {
                    switch "hello" {
                    case "\n\t":
                        v68 := GenS8{F0: 0, F1: true, F2: v21}
                        _ = v68
                    default:
                        if (42 >= 1) {
                            v21 = 747
                            switch 987 {
                            case -292:
                                if len(v22) >= 2 {
                                    v22 = v22[1:len(v22)]
                                }
                            case 32768:
                                var v69 int16 = 42
                                _ = v69
                            case 9223372036854775807:
                                v70 := GenS9{F0: (fi25 / ((fi67) | 1)), F1: 0}
                                _ = v70
                            default:
                                var v71 int = (^(fi25))
                                _ = v71
                            }
                            if v20 {
                                _ = len(v1)
                                v72 := GenS10{F0: fi25}
                                fmt.Println(v24, fi67, fi67)
                                _ = v72
                            } else {
                                v2 -= 104
                            }
                        } else {
                            fmt.Println(v3, v2)
                            switch (fi25 * fi67) {
                            case 9223372036854775807:
                                v24 -= genFn0()
                            case 32767:
                                if v20 {
                                    switch ("" + "a") {
                                    case "foo bar":
                                        var v73 string = "\n\t"
                                        _ = v73
                                    case "a":
                                        if len(v22) > 0 {
                                            v22[0] = (fi67 - (fi25 - fi25))
                                        }
                                    }
                                }
                            }
                            v3 = -238
                            if (!true) {
                                if len(v22) > 0 {
                                    v22[0] = fi67
                                }
                                fmt.Println(v24, v2, v2)
                                v74 := GenS11{F0: 1000}
                                var v75 int8 = v2
                                _ = v74
                                _ = v75
                            } else {
                                var v76 float64 = 0.0
                                var v77 float64 = ((v76 / 1.0) - (3.14159 + (((float64((3.14159 + ((0.1 / -2.25) + 1e10))) * v76) - v76) / float64(1e10))))
                                v78 := GenS3{F0: ("hello" + "hello"), F1: v2}
                                _ = v76
                                _ = v77
                                _ = v78
                            }
                        }
                    }
                    _ = fi67
                }
            }
            v3 = v3
        }
        _ = fi25
    }
    var v79 uint = 1
    if len(v22) > 0 {
        v22[0] = 32767
    }
    if v20 {
        var v80 int32 = -1
        var v81 int8 = (v2 | (v2 % ((0) | 1)))
        _ = v80
        _ = v81
    }
    _ = len(v22)
    v79 = v79
    var v82 float32 = float32(1e10)
    v79 = (^(v79))
    if (!(42 <= 65535)) {
        v83 := GenS4{F0: "foo bar", F1: (-(v2))}
        v84 := GenS5{F0: v21, F1: 1000, F2: (v2 << ((v79) & 63))}
        for fi85 := 0; fi85 < 11; fi85++ {
            switch (fi85 % ((fi85) | 1)) {
            case 2147483647:
                for fi86 := 0; fi86 < 3; fi86++ {
                    v21 = v21
                    v3 = v3
                    _ = fi86
                }
            case 932:
                var v87 uint8 = 1
                _ = v87
            case 128:
                v24 = 115
            }
            fmt.Println(v3, v3, v82)
            v2 = (v2 % ((127) | 1))
            _ = fi85
        }
        _ = v83
        _ = v84
    } else {
        var v88 int = 369
        switch (v88 | -9223372036854775808) {
        case 779:
            v1[v21] = (!((!(!v20)) || false))
        case 2147483647:
            var v89 int32 = (v3 % ((-292) | 1))
            _ = v89
        case 9223372036854775807:
            v20 = false
        default:
            var v90 int32 = v3
            _ = v90
        }
        v82 -= float32(((0.1 * -2.25) + 1.5))
        _ = v88
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v79:", v79)
    fmt.Println("v82:", v82)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v24
    _ = v3
    _ = v79
    _ = v82
}
