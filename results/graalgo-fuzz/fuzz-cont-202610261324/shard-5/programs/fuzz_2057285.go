package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32 }
type GenS2 struct { F0 int; F1 string }
type GenS3 struct { F0 uint32; F1 int8 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint32; F1 uint32; F2 uint16 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 uint; F1 uint }
type GenS8 struct { F0 int8; F1 int32 }
type GenS9 struct { F0 int8; F1 bool; F2 string }
func genFn0() int32 {
    if false {
        v5 := []int{}
        _ = v5
    }
    if (!true) {
        var v6 int16 = 32767
        _ = v6
    }
    if false {
        for fi7 := 0; fi7 < 2; fi7++ {
            var v8 int = 128
            for fi9 := 0; fi9 < 9; fi9++ {
                v8 = v8
                var v10 int = fi7
                v8 += -2147483648
                _ = fi9
                _ = v10
            }
            if (42 != 9223372036854775807) {
                for fi11 := 0; fi11 < 15; fi11++ {
                    var v12 float32 = 1e10
                    v12 = v12
                    v12 = 3.14159
                    _ = fi11
                    _ = v12
                }
                v13 := GenS1{F0: 9223372036854775807, F1: 4294967295}
                _ = v13
            } else {
                v8 = (fi7 * (fi7 % ((fi7) | 1)))
                v14 := GenS2{F0: (v8 | -1), F1: ("\n\t" + ("hello" + ("\n\t" + "a")))}
                _ = v14
            }
            _ = fi7
            _ = v8
        }
    }
    var v15 string = "foo bar"
    _ = v15
    return 929
}

func main() {
    var v1 float32 = (1e10 + 3.14159)
    var v2 string = "\n\t"
    var v3 uint8 = 187
    v4 := []int{1, -432}
    var v16 int64 = -9223372036854775808
    switch 1 {
    case 4294967295:
        fmt.Println(v1, v16, v2)
    case 128:
        switch v2 {
        case "a":
            v1 -= v1
        case "hello":
            v1 -= v1
        case "foo bar":
            v1 -= 1.5
        }
    default:
        v16 -= (v16 + v16)
    }
    v4 = append(v4, 4294967295)
    switch -809 {
    case 422:
        for fi17 := 0; fi17 < 9; fi17++ {
            if (883 >= 18446744073709551615) {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                switch "hello" {
                case "a":
                    v3 = 2
                case "hello":
                    v18 := GenS3{F0: 42, F1: 118}
                    _ = v18
                default:
                    v1 = (3.14159 - 1e10)
                }
                if (17 == -87) {
                    if len(v4) > 0 {
                        v4[0] = fi17
                    }
                    v16 += v16
                } else {
                    var v19 int64 = v16
                    v2 = v2
                    for fi20 := 0; fi20 < 11; fi20++ {
                        var v21 bool = ((3.14159 - 1e10) > (0.1 / -2.25))
                        _ = fi20
                        _ = v21
                    }
                    _ = v19
                }
            }
            fmt.Println(v1, v16, v2)
            v22 := GenS1{F0: (v16 | (v16 << ((2) & 63))), F1: uint32(v3)}
            _ = fi17
            _ = v22
        }
    default:
        fmt.Println(v2, v2, v16)
    }
    if false {
        for fi23 := 0; fi23 < 7; fi23++ {
            if (32767 > 42) {
                if false {
                    for fi24 := 0; fi24 < 15; fi24++ {
                        fmt.Println(v3, v2, v3)
                        fmt.Println(v3, v16)
                        _ = fi24
                    }
                    for fi25 := 0; fi25 < 1; fi25++ {
                        var v26 int8 = 127
                        _ = len(v4)
                        _ = fi25
                        _ = v26
                    }
                    if (!false) {
                        for fi27 := 0; fi27 < 5; fi27++ {
                            var v28 float64 = 0.0
                            v2 = ("" + v2)
                            switch (("foo bar" + "\n\t") + v2) {
                            case "\n\t":
                                fmt.Println(fi27, fi23, v16)
                            case "a":
                                if len(v4) > 0 {
                                    v4[0] = 42
                                }
                            case "":
                                v1 = v1
                            default:
                                if false {
                                    v1 = v1
                                    var v29 float32 = v1
                                    _ = v29
                                }
                            }
                            _ = fi27
                            _ = v28
                        }
                        var v30 uint64 = 18446744073709551615
                        _ = v30
                    } else {
                        v31 := GenS4{F0: 0}
                        _ = v31
                    }
                } else {
                    var v32 int64 = 410
                    var v33 uint32 = 2147483648
                    _ = v32
                    _ = v33
                }
                var v34 int32 = 0
                _ = v34
            } else {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                fmt.Println(v3, v3)
                for fi35 := 0; fi35 < 12; fi35++ {
                    fmt.Println(v2)
                    _ = fi35
                }
            }
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            v1 -= v1
            _ = fi23
        }
        for fi36 := 0; fi36 < 9; fi36++ {
            v16 += v16
            _ = fi36
        }
    } else {
        for fi37 := 0; fi37 < 8; fi37++ {
            if true {
                var v38 int = 9223372036854775807
                _ = v38
            }
            fmt.Println(v1, v1, v16)
            switch ((v2 + (v2 + "\n\t")) + "a") {
            case "\n\t":
                if len(v4) > 0 {
                    v4[0] = -9223372036854775808
                }
            case "":
                var v39 string = "foo bar"
                _ = v39
            default:
                if (1 > 747) {
                    v3 = (v3 ^ 0)
                    if (false && (65535 == 565)) {
                        for fi40 := 0; fi40 < 8; fi40++ {
                            v41 := GenS2{F0: 32768, F1: (v2 + v2)}
                            var v42 uint8 = v3
                            switch v2 {
                            case "foo bar":
                                v42 = v3
                            case "a":
                                _ = len(v4)
                            default:
                                var v43 string = v2
                                _ = v43
                            }
                            if (0 < 1000) {
                                var v44 float32 = v1
                                v16 += v16
                                v45 := GenS1{F0: -9223372036854775808, F1: 2147483647}
                                var v46 uint64 = 42
                                v4 = append(v4, fi40)
                                _ = v44
                                _ = v45
                                _ = v46
                            }
                            v2 = ("foo bar" + v2)
                            _ = fi40
                            _ = v41
                            _ = v42
                        }
                        _ = len(v4)
                        var v47 int8 = 100
                        _ = v47
                    } else {
                        for fi48 := 0; fi48 < 3; fi48++ {
                            v4 = append(v4, fi48)
                            _ = fi48
                        }
                        var v49 int64 = v16
                        v4 = append(v4, (^(fi37)))
                        _ = v49
                    }
                }
            }
            _ = fi37
        }
        if true {
            var v50 int32 = -1
            v3 = 128
            _ = v50
        } else {
            for fi51 := 0; fi51 < 6; fi51++ {
                for fi52 := 0; fi52 < 13; fi52++ {
                    if len(v4) > 0 {
                        v4[0] = 32767
                    }
                    v1 = (((v1 - 1.5) * v1) - v1)
                    _ = fi52
                }
                _ = fi51
            }
            var v53 int32 = -2147483648
            _ = len(v4)
            _ = v53
        }
    }
    v3 += (^(v3))
    v54 := GenS2{F0: 32768, F1: v2}
    v2 = ("hello" + "a")
    var v55 uint64 = 42
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v56 uint = 18446744073709551615
    if true {
        var v57 uint8 = v3
        _ = v57
    }
    if len(v4) > 0 {
        v4[0] = -2147483648
    }
    var v58 int = -9223372036854775808
    if len(v4) > 0 {
        v4[0] = v58
    }
    if false {
        var v59 uint16 = 255
        v60 := GenS5{F0: 451, F1: uint32(v3), F2: 404}
        _ = v59
        _ = v60
    }
    var v61 uint32 = 4294967295
    if true {
        v1 = (((v1 / 0.1) - v1) * (0.0 - float32(float64((v1 / 0.1)))))
        if (42 <= 416) {
            fmt.Println(v1)
        } else {
            for fi62 := 0; fi62 < 19; fi62++ {
                var v63 int8 = 38
                if false {
                    var v64 float64 = 0.0
                    if (!true) {
                        for fi65 := 0; fi65 < 17; fi65++ {
                            fmt.Println(v1, v3)
                            _ = fi65
                        }
                        v66 := GenS3{F0: 850, F1: (-(v63))}
                        fmt.Println(v1, v55, v16)
                        _ = v66
                    }
                    v67 := GenS6{F0: -1}
                    _ = v64
                    _ = v67
                } else {
                    v16 += -54
                    v61 += (v61 >> (((v56 << (((v56 & (v56 & (v56 ^ (v56 * 42))))) & 63))) & 63))
                    if ((!(true || ((!true) && false))) && (32767 < 0)) {
                        v3 += v3
                        v55 = v55
                    }
                }
                for fi68 := 0; fi68 < 17; fi68++ {
                    v2 = v2
                    if ((v2 <= (("foo bar" + "") + v2)) || true) {
                        if false {
                            if true {
                                switch fi62 {
                                case 127:
                                    if true {
                                        v4 = append(v4, fi62)
                                        if (0.0 >= ((((-2.25 * 3.14159) / (0.0 + 3.14159)) * 3.14159) * (1e10 - 3.14159))) {
                                            v61 = (v61 + v61)
                                        } else {
                                            switch "foo bar" {
                                            case "hello":
                                                v56 = uint(v3)
                                            }
                                        }
                                        v63 -= v63
                                    } else {
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                        switch v2 {
                                        case "":
                                            switch (v58 * fi68) {
                                            case 42:
                                                fmt.Println(v16, v2, v55)
                                            default:
                                                if true {
                                                    v16 = (-(v16))
                                                } else {
                                                    v69 := GenS1{F0: v16, F1: 787}
                                                    if (-503 >= v16) {
                                                        v61 = v61
                                                        fmt.Println(v2, v1, v1)
                                                        v61 = v61
                                                    }
                                                    v61 += 0
                                                    _ = v69
                                                }
                                            }
                                        default:
                                            _ = len(v4)
                                        }
                                    }
                                case 616:
                                    var v70 int16 = -719
                                    _ = v70
                                case 761:
                                    if true {
                                        var v71 int64 = v16
                                        fmt.Println(v58)
                                        _ = v71
                                    } else {
                                        v72 := GenS1{F0: v16, F1: (v61 * 4294967295)}
                                        var v73 string = v2
                                        v1 += (-2.25 / v1)
                                        _ = len(v4)
                                        _ = v72
                                        _ = v73
                                    }
                                }
                                v61 += 2147483647
                            } else {
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                                var v74 int16 = 42
                                _ = v74
                            }
                        } else {
                            var v75 string = v2
                            if false {
                                switch 0 {
                                case 32768:
                                    var v76 int = fi68
                                    _ = v76
                                default:
                                    v16 = 987
                                }
                                switch (fi62 + v58) {
                                case 2147483647:
                                    fmt.Println(v63, v63, v75)
                                case 453:
                                    if false {
                                        v56 -= (v56 % ((v56) | 1))
                                        v77 := GenS7{F0: (v56 + v56), F1: v56}
                                        if false {
                                            v61 = 42
                                            v63 += v63
                                            v78 := GenS2{F0: fi62, F1: (v2 + v75)}
                                            _ = v78
                                        } else {
                                            v1 += (0.0 * 0.1)
                                        }
                                        fmt.Println(v55, fi62, v3)
                                        var v79 uint16 = 65535
                                        _ = v77
                                        _ = v79
                                    }
                                default:
                                    var v80 bool = false
                                    _ = v80
                                }
                                if (42 <= 1000) {
                                    v1 = v1
                                }
                                _ = len(v4)
                            } else {
                                v58 = fi62
                            }
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            v1 = (v1 * (float32(0.1) * 1e10))
                            _ = v75
                        }
                    } else {
                        if true {
                            v81 := GenS3{F0: v61, F1: (v63 % ((v63) | 1))}
                            switch v2 {
                            case "a":
                                var v82 string = (("" + (v2 + "\n\t")) + ((v2 + v2) + v2))
                                _ = v82
                            case "":
                                if true {
                                    v83 := GenS8{F0: v63, F1: 1}
                                    v84 := GenS6{F0: -325}
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                    _ = v83
                                    _ = v84
                                }
                            }
                            _ = v81
                        } else {
                            switch (fi62 - v58) {
                            case 934:
                                var v85 uint64 = v55
                                _ = v85
                            }
                        }
                        if false {
                            var v86 bool = ((false && ((^(fi62)) < 1)) || (!false))
                            if len(v4) > 0 {
                                v4[0] = (fi68 % ((int(v3)) | 1))
                            }
                            v1 -= (v1 / (float32(1e10) * -2.25))
                            _ = v86
                        }
                    }
                    _ = fi68
                }
                _ = fi62
                _ = v63
            }
            var v87 bool = (false && true)
            _ = v87
        }
        for fi88 := 0; fi88 < 3; fi88++ {
            v89 := GenS9{F0: 42, F1: true, F2: "hello"}
            _ = fi88
            _ = v89
        }
    } else {
        _ = len(v4)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v55:", v55)
    fmt.Println("v56:", v56)
    fmt.Println("v58:", v58)
    fmt.Println("v61:", v61)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v54
    _ = v55
    _ = v56
    _ = v58
    _ = v61
}
