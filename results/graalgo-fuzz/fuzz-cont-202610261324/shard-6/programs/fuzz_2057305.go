package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint; F1 int32 }
type GenS6 struct { F0 int64; F1 string }
type GenS7 struct { F0 string; F1 int }
func genFn0(p0 uint16) uint16 {
    for fi7 := 0; fi7 < 15; fi7++ {
        for fi8 := 0; fi8 < 9; fi8++ {
            p0 -= p0
            v9 := GenS1{F0: 42, F1: fi7}
            _ = fi8
            _ = v9
        }
        _ = fi7
    }
    v10 := GenS1{F0: 100, F1: 127}
    var v11 int16 = -1
    _ = p0
    _ = v10
    _ = v11
    return p0
}
func genFn1(p0 uint16) int32 {
    for fi12 := 0; fi12 < 9; fi12++ {
        for fi13 := 0; fi13 < 6; fi13++ {
            if true {
                var v14 float64 = (1.5 * (float64(((1.5 / 0.1) * 0.1)) - 1e10))
                v15 := GenS2{F0: 128}
                _ = v14
                _ = v15
            }
            fmt.Println(p0, fi13, fi12)
            v16 := GenS2{F0: 128}
            fmt.Println(fi13)
            p0 += (p0 + p0)
            _ = fi13
            _ = v16
        }
        var v17 int = fi12
        _ = fi12
        _ = v17
    }
    if true {
        if (32767 == 127) {
            if true {
                v18 := []int{}
                var v19 int16 = 42
                switch "foo bar" {
                case "":
                    for fi20 := 0; fi20 < 20; fi20++ {
                        v19 = v19
                        if false {
                            v19 += int16(p0)
                        } else {
                            var v21 int = fi20
                            _ = v21
                        }
                        _ = fi20
                    }
                case "hello":
                    p0 = p0
                case "a":
                    if false {
                        switch 32767 {
                        case 128:
                            for fi22 := 0; fi22 < 9; fi22++ {
                                var v23 uint16 = 0
                                _ = fi22
                                _ = v23
                            }
                        }
                    } else {
                        p0 += p0
                        if (v19 < (^(v19))) {
                            if len(v18) >= 2 {
                                v18 = v18[1:len(v18)]
                            }
                            var v24 int16 = (v19 & (v19 / ((v19) | 1)))
                            _ = v24
                        }
                    }
                default:
                    v19 = v19
                }
                _ = v18
                _ = v19
            }
        }
        v25 := GenS3{F0: -32768}
        var v26 int64 = -1
        _ = v25
        _ = v26
    } else {
        fmt.Println(p0)
        if (!(!(1e10 != 3.14159))) {
            if (-128 <= 0) {
                for fi27 := 0; fi27 < 4; fi27++ {
                    if (-32768 != 594) {
                        var v28 float64 = 3.14159
                        var v29 int32 = 1
                        p0 = p0
                        _ = v28
                        _ = v29
                    } else {
                        var v30 string = ("a" + ((("\n\t" + "foo bar") + ("foo bar" + ("foo bar" + "foo bar"))) + "a"))
                        for fi31 := 0; fi31 < 11; fi31++ {
                            var v32 int = -2147483648
                            if false {
                                v30 = (v30 + v30)
                                v32 = fi31
                                var v33 float64 = (3.14159 / 1e10)
                                _ = v33
                            } else {
                                v34 := []int{(v32 >> ((18446744073709551615) & 63)), -9223372036854775808, 823}
                                v35 := GenS2{F0: 107}
                                _ = v34
                                _ = v35
                            }
                            p0 = p0
                            _ = fi31
                            _ = v32
                        }
                        _ = v30
                    }
                    for fi36 := 0; fi36 < 14; fi36++ {
                        v37 := GenS2{F0: 79}
                        _ = fi36
                        _ = v37
                    }
                    _ = fi27
                }
                var v38 int16 = -32768
                _ = v38
            } else {
                v39 := GenS3{F0: 42}
                _ = v39
            }
            var v40 int = -1
            _ = v40
        }
    }
    _ = p0
    return -436
}

func main() {
    var v1 uint32 = 2147483647
    var v2 float64 = 1e10
    var v3 uint = 849
    v4 := []uint32{911, v1}
    var v5 int16 = 1000
    var v6 int = 32767
    v41 := GenS4{F0: 9}
    for fi42 := 0; fi42 < 17; fi42++ {
        var v43 bool = (127 == 192)
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = fi42
        _ = v43
    }
    if true {
        var v44 int32 = 1
        for fi45 := 0; fi45 < 6; fi45++ {
            if len(v4) > 0 {
                v4[0] = v1
            }
            v46 := GenS2{F0: 1}
            _ = fi45
            _ = v46
        }
        v3 = (v3 | (v3 >> (((v3 & 654)) & 63)))
        _ = v44
    }
    v2 = 0.1
    fmt.Println(v5, v3)
    var v47 int32 = 42
    if false {
        var v48 uint32 = v1
        _ = v48
    } else {
        var v49 int8 = 100
        for fi50 := 0; fi50 < 9; fi50++ {
            v4 = append(v4, (v1 % ((v1) | 1)))
            fmt.Println(v1)
            _ = fi50
        }
        fmt.Println(v5)
        _ = v49
    }
    v1 += (^(v1))
    for fi51 := 0; fi51 < 12; fi51++ {
        v2 += v2
        var v52 float32 = (0.0 - (float32(3.14159) + 0.0))
        _ = fi51
        _ = v52
    }
    var v53 int16 = -32768
    if (false && (!true)) {
        v47 += v47
        for fi54 := 0; fi54 < 14; fi54++ {
            v55 := GenS3{F0: (v53 / ((0) | 1))}
            v5 += v5
            _ = fi54
            _ = v55
        }
    }
    v53 -= 1000
    v56 := GenS5{F0: (v3 << ((747) & 63)), F1: (v47 * (v47 | 1))}
    if (!(0.1 >= (float32(((v2 * v2) - 0.0)) * (-2.25 / 1.0)))) {
        for fi57 := 0; fi57 < 13; fi57++ {
            if len(v4) > 0 {
                v4[0] = 2147483647
            }
            for fi58 := 0; fi58 < 20; fi58++ {
                var v59 uint8 = 128
                v1 = (v1 % ((v1) | 1))
                _ = fi58
                _ = v59
            }
            _ = fi57
        }
        switch v6 {
        case 614:
            fmt.Println(v47, v3, v1)
        case 32768:
            if (true || false) {
                v60 := GenS5{F0: v3, F1: (v47 - 0)}
                v61 := GenS2{F0: 56}
                _ = v60
                _ = v61
            } else {
                var v62 float32 = ((((1.5 / 1.0) * (3.14159 + 3.14159)) + 1e10) * 1.5)
                if (false || true) {
                    v63 := GenS6{F0: 0, F1: "hello"}
                    var v64 int64 = 333
                    _ = len(v4)
                    _ = v63
                    _ = v64
                } else {
                    if (v3 == (v3 % ((v3) | 1))) {
                        switch "hello" {
                        case "\n\t":
                            switch "\n\t" {
                            case "\n\t":
                                switch 32768 {
                                case 1:
                                    var v65 int8 = -38
                                    _ = v65
                                default:
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                }
                            case "":
                                fmt.Println(v6, v53)
                            case "hello":
                                var v66 uint8 = 0
                                _ = v66
                            }
                        case "hello":
                            v2 = ((v2 + v2) / float64(((v62 * v62) * 0.0)))
                        default:
                            v6 = v6
                        }
                        if (!true) {
                            v6 = v6
                            var v67 uint32 = 171
                            _ = v67
                        }
                        var v68 int = 1
                        _ = v68
                    } else {
                        var v69 int16 = v5
                        v70 := GenS1{F0: 1, F1: v6}
                        _ = v69
                        _ = v70
                    }
                    if (false || ((v53 + v5) == -835)) {
                        var v71 int16 = v53
                        var v72 string = "foo bar"
                        fmt.Println(v72)
                        _ = v71
                        _ = v72
                    }
                }
                fmt.Println(v5, v2, v3)
                _ = v62
            }
        case 9223372036854775807:
            switch "" {
            case "hello":
                switch 32768 {
                case 9223372036854775807:
                    v73 := GenS1{F0: -1, F1: 446}
                    _ = v73
                case 127:
                    for fi74 := 0; fi74 < 19; fi74++ {
                        for fi75 := 0; fi75 < 20; fi75++ {
                            if true {
                                _ = len(v4)
                            }
                            if (!false) {
                                switch fi74 {
                                case 2147483647:
                                    v4 = append(v4, (v1 << (((^(v3))) & 63)))
                                default:
                                    if false {
                                        var v76 uint16 = 0
                                        _ = v76
                                    }
                                }
                                v2 += v2
                            }
                            switch v6 {
                            case 32767:
                                v47 -= v47
                            case -298:
                                var v77 string = ("hello" + "a")
                                _ = v77
                            default:
                                var v78 float64 = v2
                                _ = v78
                            }
                            _ = fi75
                        }
                        _ = fi74
                    }
                }
            default:
                fmt.Println(v6, v47)
            }
        default:
            if false {
                v47 = 121
                switch ("a" + "") {
                case "":
                    var v79 int32 = v47
                    _ = v79
                case "\n\t":
                    var v80 float64 = v2
                    _ = v80
                case "hello":
                    v81 := GenS7{F0: "a", F1: (v6 | (v6 * (v6 - 0)))}
                    _ = v81
                }
            }
        }
        var v82 bool = (-9223372036854775808 >= 652)
        _ = v82
    }
    for fi83 := 0; fi83 < 19; fi83++ {
        for fi84 := 0; fi84 < 6; fi84++ {
            v6 += (fi84 << (((v3 >> (((v3 ^ 2)) & 63))) & 63))
            _ = fi84
        }
        _ = fi83
    }
    var v85 int32 = (-(v47))
    var v86 uint32 = (v1 | (v1 % (((v1 ^ (v1 + v1))) | 1)))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    fmt.Println("v53:", v53)
    fmt.Println("v6:", v6)
    fmt.Println("v85:", v85)
    fmt.Println("v86:", v86)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v47
    _ = v5
    _ = v53
    _ = v56
    _ = v6
    _ = v85
    _ = v86
}
