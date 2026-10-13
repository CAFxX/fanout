package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int }
type GenS2 struct { F0 uint16; F1 uint; F2 uint }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 bool; F1 int32 }
type GenS5 struct { F0 uint64; F1 int; F2 int }
type GenS6 struct { F0 int; F1 int8 }
type GenS7 struct { F0 uint32; F1 uint64 }
type GenS8 struct { F0 int; F1 string }
type GenS9 struct { F0 int16 }
type GenS10 struct { F0 int64 }
type GenS11 struct { F0 uint64 }
type GenS12 struct { F0 bool }
type GenS13 struct { F0 string; F1 int16 }
type GenS14 struct { F0 uint32; F1 uint32 }
type GenS15 struct { F0 string; F1 int32; F2 int }
func genFn0(p0 int, p1 int64) string {
    v6 := GenS1{F0: true, F1: 32767}
    v7 := GenS2{F0: 0, F1: 504, F2: 42}
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return ((("hello" + (((("\n\t" + "") + (("foo bar" + "hello") + (("\n\t" + "hello") + "\n\t"))) + "") + "a")) + "hello") + ("hello" + "\n\t"))
}
func genFn1(p0 int, p1 bool, p2 int64) uint32 {
    p0 = (p0 << ((4294967295) & 63))
    v8 := GenS3{F0: -433}
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 1
}

func main() {
    var v1 float64 = 0.1
    var v2 int16 = -32768
    var v3 float32 = float32((v1 * v1))
    v4 := make(map[uint32]uint32)
    var v5 bool = (18446744073709551615 >= 0)
    var v9 uint32 = 1
    _ = len(v4)
    v10 := GenS4{F0: (423 <= 54), F1: 1}
    var v11 uint32 = v9
    var v12 uint32 = (v9 - (^(v11)))
    switch "" {
    case "\n\t":
        v2 = v2
    case "a":
        v3 += v3
    case "hello":
        if v5 {
            v2 = (v2 * v2)
        } else {
            var v13 int32 = 242
            _ = v13
        }
    }
    v1 = -2.25
    if v5 {
        for fi14 := 0; fi14 < 7; fi14++ {
            fmt.Println(v5, v11)
            _ = fi14
        }
        switch int(v2) {
        case 2147483647:
            var v15 float32 = (1.5 + v3)
            _ = v15
        case 128:
            if (128 >= 255) {
                v16 := []int{0}
                _ = len(v4)
                for fi17 := 0; fi17 < 17; fi17++ {
                    for fi18 := 0; fi18 < 13; fi18++ {
                        v5 = v5
                        v1 += v1
                        _ = fi18
                    }
                    switch "\n\t" {
                    case "a":
                        var v19 uint = 1
                        _ = v19
                    case "":
                        v5 = v5
                    default:
                        fmt.Println(v2, v3, v9)
                    }
                    if len(v16) >= 2 {
                        v16 = v16[1:len(v16)]
                    }
                    _ = fi17
                }
                _ = v16
            }
        default:
            v2 = v2
        }
    }
    for fi20 := 0; fi20 < 18; fi20++ {
        switch fi20 {
        case -9223372036854775808:
            var v21 int32 = -754
            _ = v21
        case 32768:
            for fi22 := 0; fi22 < 16; fi22++ {
                v9 += (^(v11))
                _ = fi22
            }
        default:
            if (-9223372036854775808 > -9223372036854775808) {
                for fi23 := 0; fi23 < 3; fi23++ {
                    switch "hello" {
                    case "a":
                        var v24 uint = 931
                        _ = v24
                    case "hello":
                        if false {
                            if v5 {
                                v9 = (v12 << ((2) & 63))
                                if true {
                                    if ((!true) && (-711 <= -1)) {
                                        if v5 {
                                            v25 := GenS5{F0: 572, F1: fi23, F2: (^(fi23))}
                                            v9 = (v9 & genFn1(fi20, true, 0))
                                            var v26 int8 = 0
                                            _ = v25
                                            _ = v26
                                        }
                                        var v27 int8 = -73
                                        _ = v27
                                    }
                                }
                                fmt.Println(fi23, fi23)
                                v3 = 1.5
                            } else {
                                switch fi23 {
                                case 32768:
                                    fmt.Println(v1, fi23, v5)
                                default:
                                    v11 -= (v9 << ((2) & 63))
                                }
                                switch (("foo bar" + "foo bar") + ("a" + ("foo bar" + "hello"))) {
                                case "hello":
                                    var v28 uint32 = 2147483647
                                    _ = v28
                                case "\n\t":
                                    switch (fi20 / ((32768) | 1)) {
                                    case 32767:
                                        v29 := GenS5{F0: 0, F1: 127, F2: (fi20 << ((65535) & 63))}
                                        _ = v29
                                    case 0:
                                        v30 := GenS6{F0: (fi23 >> ((255) & 63)), F1: -26}
                                        _ = v30
                                    }
                                default:
                                    _ = len(v4)
                                }
                                v31 := GenS2{F0: 255, F1: 42, F2: 674}
                                _ = v31
                            }
                        } else {
                            if (!v5) {
                                fmt.Println(v1)
                                var v32 int16 = (v2 % ((1) | 1))
                                _ = v32
                            } else {
                                v33 := GenS7{F0: (v11 / ((v9) | 1)), F1: uint64(v9)}
                                _ = v33
                            }
                            v1 += ((v1 / 1.5) + (1.5 - ((v1 + -2.25) + v1)))
                        }
                    case "foo bar":
                        v34 := GenS8{F0: (fi20 << ((2) & 63)), F1: ""}
                        _ = v34
                    default:
                        if (v5 || (2 != 1)) {
                            v5 = true
                            var v35 uint32 = v9
                            switch "hello" {
                            case "foo bar":
                                v4[v12] = v12
                            default:
                                if (764 <= 32767) {
                                    switch ("" + ("hello" + ("\n\t" + ("hello" + "\n\t")))) {
                                    case "hello":
                                        v4[1] = (v11 ^ 42)
                                        v3 = (v3 + v3)
                                        v3 -= float32(v1)
                                    }
                                } else {
                                    if v5 {
                                        if v5 {
                                            var v36 uint32 = v12
                                            _ = v36
                                        }
                                    }
                                }
                            }
                            v37 := []int{425, -12, 127, fi20}
                            var v38 int64 = 9223372036854775807
                            switch fi20 {
                            case 568:
                                var v39 float32 = v3
                                _ = v39
                            case 779:
                                v11 += 1
                            default:
                                var v40 float64 = (v1 + float64(float32(3.14159)))
                                _ = v40
                            }
                            v41 := GenS4{F0: false, F1: -1}
                            _ = v35
                            _ = v37
                            _ = v38
                            _ = v41
                        }
                    }
                    var v42 uint16 = 65535
                    fmt.Println(v12)
                    _ = fi23
                    _ = v42
                }
                v11 = 145
                var v43 int64 = 9223372036854775807
                _ = v43
            } else {
                v1 -= 1e10
                v1 += (1.5 / (3.14159 + v1))
            }
        }
        v44 := GenS9{F0: v2}
        _ = fi20
        _ = v44
    }
    switch 32767 {
    case 2147483647:
        for fi45 := 0; fi45 < 9; fi45++ {
            v46 := GenS10{F0: -1}
            switch 127 {
            case -1:
                var v47 uint16 = 291
                _ = v47
            case 4294967295:
                for fi48 := 0; fi48 < 5; fi48++ {
                    var v49 float64 = v1
                    v3 = (-2.25 * v3)
                    switch int(v9) {
                    case -593:
                        v50 := GenS3{F0: v2}
                        _ = v50
                    case 670:
                        switch (fi45 % (((fi48 / ((32767) | 1))) | 1)) {
                        case 32768:
                            switch "foo bar" {
                            case "":
                                v9 -= (v11 | genFn1(fi48, true, -1))
                            case "foo bar":
                                v51 := GenS1{F0: v5, F1: fi48}
                                _ = v51
                            case "\n\t":
                                v2 = (v2 & (v2 & v2))
                            }
                            v52 := GenS11{F0: 2}
                            _ = v52
                        case 128:
                            var v53 int32 = 294
                            _ = v53
                        case 4294967295:
                            v54 := GenS12{F0: (!(!(true || (false && v5))))}
                            _ = v54
                        default:
                            var v55 uint16 = 65535
                            _ = v55
                        }
                    }
                    var v56 uint = 18446744073709551615
                    fmt.Println(v3, v11)
                    if v5 {
                        fmt.Println(v11, v5)
                    } else {
                        if (v5 && v5) {
                            v57 := v4[v11]
                            v58 := []int{(-(fi48)), (fi45 + (fi48 & (fi45 + (fi48 + 338)))), 32767, fi45}
                            v11 += (v57 * genFn1((fi48 & (fi48 ^ (fi48 + fi48))), v5, -9223372036854775808))
                            _ = v57
                            _ = v58
                        }
                        var v59 int8 = -1
                        _ = v59
                    }
                    switch "a" {
                    case "foo bar":
                        if (100 > -128) {
                            var v60 uint8 = 2
                            v61 := GenS6{F0: (fi45 | 32767), F1: -1}
                            v9 += (^(v12))
                            _ = v60
                            _ = v61
                        }
                    case "":
                        switch ("a" + ("" + "\n\t")) {
                        case "foo bar":
                            var v62 bool = v5
                            _ = v62
                        default:
                            var v63 float32 = 0.1
                            _ = v63
                        }
                    }
                    _ = fi48
                    _ = v49
                    _ = v56
                }
            case -2147483648:
                v5 = (741 != 2147483647)
            }
            var v64 int = -1
            _ = fi45
            _ = v46
            _ = v64
        }
    default:
        for fi65 := 0; fi65 < 14; fi65++ {
            v66 := GenS13{F0: "foo bar", F1: v2}
            var v67 float64 = 0.0
            if (-1 < -128) {
                for fi68 := 0; fi68 < 2; fi68++ {
                    v69 := GenS11{F0: 112}
                    if (fi68 <= (-(fi68))) {
                        v5 = (42 <= 0)
                        var v70 uint = 18446744073709551615
                        v12 += v12
                        _ = v70
                    } else {
                        fmt.Println(v5, v67)
                    }
                    _ = fi68
                    _ = v69
                }
                for fi71 := 0; fi71 < 12; fi71++ {
                    v12 = (v9 / ((65535) | 1))
                    _ = fi71
                }
                v12 = (v9 / ((v9) | 1))
            } else {
                v67 = 0.1
                for fi72 := 0; fi72 < 19; fi72++ {
                    v73 := GenS14{F0: 2147483647, F1: 2}
                    v5 = (false || (255 >= 65535))
                    if (!((^(v12)) == genFn1(127, (0 >= 0), -9223372036854775808))) {
                        v11 -= v9
                        fmt.Println(v2, v5)
                    } else {
                        fmt.Println(v5, v2)
                    }
                    _ = fi72
                    _ = v73
                }
                var v74 float32 = v3
                _ = v74
            }
            _ = fi65
            _ = v66
            _ = v67
        }
    }
    v75 := GenS15{F0: ("foo bar" + "a"), F1: -2147483648, F2: -78}
    if v5 {
        v76 := v4[v11]
        _ = v76
    }
    v11 = genFn1(-836, (!v5), -9223372036854775808)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v75
    _ = v9
}
