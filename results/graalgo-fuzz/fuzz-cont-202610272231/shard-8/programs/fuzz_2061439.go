package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool; F2 uint32 }
type GenS2 struct { F0 int16; F1 int; F2 int8 }
type GenS3 struct { F0 bool; F1 uint64; F2 int }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int }
type GenS6 struct { F0 uint64; F1 uint64; F2 uint8 }
type GenS7 struct { F0 int8; F1 int8; F2 int8 }
type GenS8 struct { F0 int8 }
type GenS9 struct { F0 int16; F1 string; F2 uint }
type GenS10 struct { F0 uint16; F1 uint16; F2 int64 }
type GenS11 struct { F0 uint64; F1 int32; F2 bool }
func genFn0() int64 {
    var v7 int = 42
    v8 := GenS1{F0: 425, F1: true, F2: 2}
    v9 := GenS2{F0: -1, F1: v7, F2: 1}
    var v10 string = ("\n\t" + "a")
    switch (v7 / ((v7) | 1)) {
    case 4294967295:
        fmt.Println(v10, v10)
    case -2147483648:
        if ((90 < 128) || ((!(1 == -2147483648)) || ((847 >= 9223372036854775807) || (1 >= 255)))) {
            if (283 > 42) {
                for fi11 := 0; fi11 < 7; fi11++ {
                    v12 := []int{}
                    _ = fi11
                    _ = v12
                }
            } else {
                var v13 uint32 = 1
                _ = v13
            }
            for fi14 := 0; fi14 < 12; fi14++ {
                fmt.Println(v7)
                _ = fi14
            }
        } else {
            v15 := GenS1{F0: 1, F1: (1 <= 0), F2: 1}
            if false {
                var v16 float64 = 0.0
                _ = v16
            } else {
                fmt.Println(v7, v7)
                if (4294967295 != 616) {
                    switch -366 {
                    case 1:
                        fmt.Println(v7, v10)
                    case -865:
                        for fi17 := 0; fi17 < 16; fi17++ {
                            v7 = fi17
                            _ = fi17
                        }
                    default:
                        v18 := GenS3{F0: (1 >= 778), F1: 18446744073709551615, F2: (v7 << ((255) & 63))}
                        _ = v18
                    }
                }
                v19 := []int{v7}
                _ = v19
            }
            _ = v15
        }
    default:
        var v20 uint8 = 168
        _ = v20
    }
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return -536
}
func genFn1(p0 int64, p1 string) int16 {
    for fi21 := 0; fi21 < 17; fi21++ {
        p0 -= (-(p0))
        v22 := GenS2{F0: 0, F1: (fi21 / ((fi21) | 1)), F2: 0}
        _ = fi21
        _ = v22
    }
    for fi23 := 0; fi23 < 3; fi23++ {
        fmt.Println(p0)
        _ = fi23
    }
    p0 -= genFn0()
    var v24 float64 = (1e10 / -2.25)
    _ = p0
    _ = p1
    _ = v24
    return 627
}

func main() {
    var v1 float32 = 1.5
    var v2 string = "foo bar"
    var v3 uint16 = 2
    var v4 int = 42
    v5 := []string{v2, "foo bar"}
    v6 := []int{}
    v25 := GenS4{F0: 4294967295}
    v3 += (v3 * (v3 - 65535))
    fmt.Println(v3, v3, v3)
    v26 := GenS1{F0: (v3 << ((714) & 63)), F1: false, F2: 65535}
    if len(v6) > 0 {
        v6[0] = v4
    }
    v6 = append(v6, (^(v4)))
    fmt.Println(v3, v1, v3)
    if (true || true) {
        v1 = 0.0
        if true {
            switch v2 {
            case "foo bar":
                switch "\n\t" {
                case "a":
                    if (v1 >= float32(-2.25)) {
                        v27 := GenS5{F0: (-(v4))}
                        _ = v27
                    }
                case "":
                    var v28 uint16 = (v3 - (v3 + v3))
                    _ = v28
                case "\n\t":
                    v1 -= 1.5
                }
            case "":
                var v29 float32 = v1
                _ = v29
            case "a":
                switch (("hello" + v2) + v2) {
                case "foo bar":
                    var v30 uint = 401
                    _ = v30
                case "a":
                    var v31 int64 = -503
                    _ = v31
                case "":
                    var v32 uint = 18446744073709551615
                    _ = v32
                default:
                    v6 = append(v6, v4)
                }
            default:
                _ = len(v5)
            }
        } else {
            v3 = v3
            v33 := GenS5{F0: (v4 ^ 9223372036854775807)}
            _ = v33
        }
        var v34 uint = 0
        _ = v34
    }
    for fi35 := 0; fi35 < 12; fi35++ {
        for fi36 := 0; fi36 < 2; fi36++ {
            var v37 int = fi36
            _ = fi36
            _ = v37
        }
        _ = fi35
    }
    fmt.Println(v2, v3, v3)
    switch v4 {
    case -631:
        v3 = 878
    }
    if (!true) {
        for fi38 := 0; fi38 < 8; fi38++ {
            switch v4 {
            case 530:
                v39 := GenS5{F0: fi38}
                _ = v39
            case 32768:
                if len(v5) > 0 {
                    v5[0] = v2
                }
            }
            if false {
                switch v2 {
                case "":
                    v40 := GenS2{F0: genFn1(-9223372036854775808, v2), F1: -1, F2: -1}
                    _ = v40
                case "hello":
                    switch fi38 {
                    case -2147483648:
                        v41 := GenS6{F0: 1, F1: 0, F2: 2}
                        _ = v41
                    case 128:
                        if false {
                            for fi42 := 0; fi42 < 3; fi42++ {
                                if true {
                                    v4 = (fi42 * v4)
                                }
                                _ = fi42
                            }
                        } else {
                            switch -2147483648 {
                            case 229:
                                var v43 float64 = (-2.25 * 0.1)
                                _ = v43
                            case 32768:
                                if true {
                                    switch 42 {
                                    case -160:
                                        var v44 uint64 = 2
                                        _ = v44
                                    case 817:
                                        v2 = v2
                                    case 1:
                                        v45 := GenS7{F0: 27, F1: 100, F2: -1}
                                        _ = v45
                                    }
                                    var v46 uint16 = v3
                                    _ = v46
                                }
                            case 4294967295:
                                switch "foo bar" {
                                case "a":
                                    v47 := GenS8{F0: 22}
                                    _ = v47
                                default:
                                    v2 = "a"
                                }
                            default:
                                var v48 uint64 = 1
                                _ = v48
                            }
                        }
                    default:
                        v6 = append(v6, (v4 * fi38))
                    }
                case "foo bar":
                    switch fi38 {
                    case 908:
                        switch "a" {
                        case "\n\t":
                            var v49 int8 = 127
                            _ = v49
                        case "foo bar":
                            fmt.Println(v3)
                        default:
                            var v50 bool = (false || ((1.5 / 0.1) >= 0.0))
                            _ = v50
                        }
                    case -250:
                        switch int(v3) {
                        case -752:
                            v51 := GenS9{F0: -32768, F1: (v2 + "a"), F2: 15}
                            _ = v51
                        case -151:
                            var v52 int64 = 0
                            _ = v52
                        default:
                            if true {
                                for fi53 := 0; fi53 < 18; fi53++ {
                                    if false {
                                        var v54 int64 = 0
                                        if false {
                                            var v55 int8 = 0
                                            v56 := GenS6{F0: 18446744073709551615, F1: 2, F2: 0}
                                            switch (v2 + v2) {
                                            case "\n\t":
                                                v57 := GenS4{F0: 42}
                                                _ = v57
                                            case "":
                                                var v58 int8 = (v55 * (v55 / (((v55 << ((4294967295) & 63))) | 1)))
                                                _ = v58
                                            default:
                                                var v59 uint64 = 4294967295
                                                _ = v59
                                            }
                                            _ = v55
                                            _ = v56
                                        } else {
                                            v1 = (3.14159 + float32(((1.5 - 0.0) + 3.14159)))
                                            var v60 uint8 = 39
                                            v1 = v1
                                            _ = v60
                                        }
                                        v61 := GenS10{F0: v3, F1: v3, F2: v54}
                                        var v62 int8 = 127
                                        _ = v54
                                        _ = v61
                                        _ = v62
                                    } else {
                                        v3 -= v3
                                        v63 := GenS1{F0: 952, F1: (61 > 70), F2: 42}
                                        v64 := GenS1{F0: (v3 + v3), F1: false, F2: 2147483647}
                                        _ = v63
                                        _ = v64
                                    }
                                    if (0 < uint64(v3)) {
                                        if true {
                                            fmt.Println(v2, fi53, fi53)
                                            v65 := GenS5{F0: fi53}
                                            _ = len(v5)
                                            _ = v65
                                        }
                                        fmt.Println(v3, v2)
                                    }
                                    var v66 uint = 18446744073709551615
                                    _ = fi53
                                    _ = v66
                                }
                                var v67 string = "a"
                                _ = v67
                            } else {
                                var v68 uint8 = 1
                                if true {
                                    if len(v5) > 0 {
                                        v5[0] = v2
                                    }
                                    if false {
                                        v1 += 1.5
                                        v68 = (v68 & (^(v68)))
                                        switch 187 {
                                        case -2147483648:
                                            var v69 string = ((v2 + (v2 + "foo bar")) + (v2 + ("foo bar" + v2)))
                                            _ = v69
                                        }
                                    }
                                }
                                _ = v68
                            }
                        }
                    default:
                        switch fi38 {
                        case 1:
                            _ = len(v6)
                        default:
                            switch (v4 | (-(v4))) {
                            case 29:
                                if ((fi38 / ((-9223372036854775808) | 1)) <= v4) {
                                    _ = len(v6)
                                } else {
                                    v2 = (v2 + (v2 + "\n\t"))
                                    v70 := GenS5{F0: fi38}
                                    _ = v70
                                }
                            case 32768:
                                for fi71 := 0; fi71 < 6; fi71++ {
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                    _ = fi71
                                }
                            case -338:
                                for fi72 := 0; fi72 < 2; fi72++ {
                                    var v73 int8 = 100
                                    _ = fi72
                                    _ = v73
                                }
                            default:
                                v3 = v3
                            }
                        }
                    }
                }
            } else {
                switch ("foo bar" + v2) {
                case "foo bar":
                    v74 := GenS11{F0: 42, F1: 1, F2: (1 > 577)}
                    _ = v74
                case "hello":
                    v3 = 2
                default:
                    v3 += v3
                }
                var v75 uint8 = 1
                switch 9223372036854775807 {
                case 32768:
                    for fi76 := 0; fi76 < 9; fi76++ {
                        var v77 uint32 = 65535
                        var v78 uint64 = 2
                        var v79 int16 = 1000
                        _ = fi76
                        _ = v77
                        _ = v78
                        _ = v79
                    }
                case 2147483647:
                    v80 := GenS6{F0: 0, F1: 1, F2: 2}
                    _ = v80
                }
                _ = v75
            }
            _ = fi38
        }
    } else {
        for fi81 := 0; fi81 < 14; fi81++ {
            if len(v6) > 0 {
                v6[0] = (v4 + (-(v4)))
            }
            _ = fi81
        }
    }
    v4 += v4
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
