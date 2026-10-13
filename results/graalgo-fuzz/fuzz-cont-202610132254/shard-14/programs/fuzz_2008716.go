package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8 }
type GenS2 struct { F0 uint64; F1 uint }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint8; F1 uint8 }
type GenS5 struct { F0 int8; F1 bool; F2 uint }
type GenS6 struct { F0 int32; F1 int16; F2 int8 }
type GenS7 struct { F0 int64; F1 string }
type GenS8 struct { F0 uint; F1 int32; F2 int8 }
type GenS9 struct { F0 int64; F1 uint16 }
type GenS10 struct { F0 uint }
func genFn0(p0 uint32, p1 int32, p2 int) uint8 {
    var v6 int = (-(p2))
    var v7 int = p2
    v8 := GenS1{F0: 665, F1: 2}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return 42
}
func genFn1(p0 int64, p1 uint32) uint {
    p1 = p1
    fmt.Println(p0)
    var v9 int = 968
    for fi10 := 0; fi10 < 13; fi10++ {
        v11 := GenS1{F0: 2, F1: 26}
        switch "a" {
        case "hello":
            p0 -= p0
        case "":
            for fi12 := 0; fi12 < 6; fi12++ {
                var v13 int8 = 1
                var v14 int8 = 42
                v9 += fi10
                _ = fi12
                _ = v13
                _ = v14
            }
        }
        _ = fi10
        _ = v11
    }
    _ = p0
    _ = p1
    _ = v9
    return 2
}

func main() {
    var v1 float64 = (3.14159 - 0.1)
    var v2 uint = 42
    v3 := []uint32{87, 2147483648}
    v4 := make(map[string]bool)
    var v5 uint16 = 32767
    switch "a" {
    case "":
        for fi15 := 0; fi15 < 7; fi15++ {
            fmt.Println(v2, v2, fi15)
            _ = fi15
        }
    case "a":
        if ("foo bar" > "\n\t") {
            switch (("a" + "a") + (("\n\t" + "") + "hello")) {
            case "foo bar":
                if true {
                    for fi16 := 0; fi16 < 13; fi16++ {
                        v2 = v2
                        _ = fi16
                    }
                } else {
                    switch "a" {
                    case "":
                        if true {
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            var v17 int8 = 0
                            v17 = (v17 ^ v17)
                            _ = v17
                        } else {
                            if false {
                                if (-9223372036854775808 == 9223372036854775807) {
                                    v18 := GenS2{F0: 42, F1: v2}
                                    _ = v18
                                } else {
                                    for fi19 := 0; fi19 < 7; fi19++ {
                                        v20 := v4["foo bar"]
                                        if ((v1 * float64((1e10 / (3.14159 - 1.5)))) <= float64((((1.5 + (0.0 - float32(v1))) + 0.1) * 3.14159))) {
                                            var v21 int = fi19
                                            _ = v21
                                        } else {
                                            v1 += v1
                                        }
                                        _ = fi19
                                        _ = v20
                                    }
                                    for fi22 := 0; fi22 < 14; fi22++ {
                                        var v23 int32 = -2147483648
                                        v1 = v1
                                        _ = fi22
                                        _ = v23
                                    }
                                    delete(v4, "\n\t")
                                }
                                fmt.Println(v5, v1)
                                var v24 uint16 = (v5 >> (((v2 - 470)) & 63))
                                _ = v24
                            }
                            if (389 <= 4294967295) {
                                if (true || (42 > 100)) {
                                    v25 := GenS3{F0: 929}
                                    for fi26 := 0; fi26 < 10; fi26++ {
                                        for fi27 := 0; fi27 < 2; fi27++ {
                                            var v28 int64 = -125
                                            if ((255 >= 127) && (!("\n\t" == "a"))) {
                                                v28 = v28
                                                switch "" {
                                                case "\n\t":
                                                    var v29 int8 = int8(v2)
                                                    _ = v29
                                                default:
                                                    if false {
                                                        switch "\n\t" {
                                                        case "a":
                                                            var v30 uint = 65535
                                                            _ = v30
                                                        case "foo bar":
                                                            v5 -= 255
                                                        }
                                                        var v31 float64 = (float64(0.1) + 3.14159)
                                                        _ = v31
                                                    }
                                                }
                                            }
                                            switch "foo bar" {
                                            case "foo bar":
                                                v5 += v5
                                            case "":
                                                v5 = (^(v5))
                                            case "a":
                                                v5 -= v5
                                            }
                                            v32 := GenS4{F0: 128, F1: 8}
                                            switch fi26 {
                                            case 9223372036854775807:
                                                if len(v3) > 0 {
                                                    v3[0] = 2147483648
                                                }
                                            default:
                                                var v33 float64 = -2.25
                                                _ = v33
                                            }
                                            _ = fi27
                                            _ = v28
                                            _ = v32
                                        }
                                        _ = fi26
                                    }
                                    var v34 uint32 = 2147483648
                                    _ = v25
                                    _ = v34
                                } else {
                                    v1 = v1
                                }
                            }
                            v5 = 42
                        }
                    }
                    if (1 != 1) {
                        switch "\n\t" {
                        case "":
                            switch 1 {
                            case 0:
                                delete(v4, ("hello" + "a"))
                            case 42:
                                if ((0.0 + -2.25) == 3.14159) {
                                    switch 1 {
                                    case 1:
                                        v35 := GenS3{F0: 128}
                                        _ = v35
                                    }
                                }
                            default:
                                var v36 int32 = -2147483648
                                _ = v36
                            }
                        case "a":
                            v5 = v5
                        case "\n\t":
                            var v37 float32 = float32(((((v1 / (v1 / v1)) + (v1 + (1e10 - v1))) - (v1 + v1)) - ((float64(float32(1.5)) + (v1 - v1)) + 0.1)))
                            _ = v37
                        default:
                            fmt.Println(v5)
                        }
                        var v38 bool = (127 != 0)
                        var v39 string = "a"
                        _ = v38
                        _ = v39
                    } else {
                        for fi40 := 0; fi40 < 2; fi40++ {
                            v41 := GenS3{F0: (fi40 | 907)}
                            delete(v4, "foo bar")
                            _ = fi40
                            _ = v41
                        }
                        for fi42 := 0; fi42 < 5; fi42++ {
                            v2 += v2
                            _ = fi42
                        }
                        var v43 uint = (v2 << (((v2 & (^(v2)))) & 63))
                        _ = v43
                    }
                }
            }
            for fi44 := 0; fi44 < 5; fi44++ {
                if len(v3) > 0 {
                    v3[0] = 2147483647
                }
                v5 = v5
                v2 -= (v2 & (^(v2)))
                _ = fi44
            }
            v45 := GenS5{F0: 6, F1: (2 != 42), F2: genFn1(717, 0)}
            _ = v45
        } else {
            v46 := GenS6{F0: 0, F1: 42, F2: -1}
            v47 := GenS7{F0: 9223372036854775807, F1: "hello"}
            var v48 float64 = (0.1 / ((v1 + 0.1) - (-2.25 + -2.25)))
            _ = v46
            _ = v47
            _ = v48
        }
    }
    if false {
        switch "a" {
        case "hello":
            var v49 uint = genFn1(int64(v2), 951)
            _ = v49
        case "\n\t":
            for fi50 := 0; fi50 < 7; fi50++ {
                v5 -= (v5 / (((^(v5))) | 1))
                v51 := GenS2{F0: 0, F1: v2}
                v2 += (^(v2))
                _ = fi50
                _ = v51
            }
        case "":
            var v52 int = 2147483647
            _ = v52
        }
        if false {
            for fi53 := 0; fi53 < 11; fi53++ {
                v54 := v4["hello"]
                v55 := v4[("a" + "foo bar")]
                _ = fi53
                _ = v54
                _ = v55
            }
        }
    } else {
        v1 += v1
        v56 := GenS8{F0: genFn1(0, 65535), F1: 789, F2: -1}
        v5 -= 298
        _ = v56
    }
    switch 32767 {
    case 128:
        var v57 uint16 = 2
        _ = v57
    case 42:
        switch -18 {
        case -9223372036854775808:
            v1 += v1
        case 966:
            switch "foo bar" {
            case "hello":
                var v58 uint8 = 108
                _ = v58
            case "foo bar":
                v3 = append(v3, 137)
            }
        case 855:
            delete(v4, "")
        default:
            if (2 >= 2) {
                if true {
                    if true {
                        if (!false) {
                            var v59 int16 = 858
                            var v60 int = 32768
                            v61 := GenS3{F0: v60}
                            _ = v59
                            _ = v60
                            _ = v61
                        }
                        var v62 uint64 = 170
                        _ = v62
                    }
                    v63 := GenS9{F0: -585, F1: (v5 - (v5 % (((v5 >> ((genFn1(130, 0)) & 63))) | 1)))}
                    _ = v63
                } else {
                    v1 += ((3.14159 - 1e10) * v1)
                }
                var v64 string = (("hello" + ("foo bar" + "a")) + "")
                _ = v64
            }
        }
    default:
        fmt.Println(v5, v5)
    }
    for fi65 := 0; fi65 < 4; fi65++ {
        v66 := GenS10{F0: (v2 * 65535)}
        if (0 != fi65) {
            _ = len(v4)
            fmt.Println(fi65, v5, v1)
            var v67 uint8 = 127
            _ = v67
        } else {
            if true {
                v68 := GenS5{F0: 42, F1: (0 <= 1), F2: (v2 * (v2 - 18446744073709551615))}
                var v69 int32 = -491
                _ = v68
                _ = v69
            } else {
                for fi70 := 0; fi70 < 14; fi70++ {
                    v71 := GenS6{F0: 42, F1: -649, F2: -128}
                    _ = fi70
                    _ = v71
                }
                for fi72 := 0; fi72 < 7; fi72++ {
                    v73 := GenS4{F0: 255, F1: 2}
                    switch "hello" {
                    case "a":
                        fmt.Println(v2, v2, v5)
                    case "\n\t":
                        v3 = append(v3, 414)
                    }
                    _ = fi72
                    _ = v73
                }
                v74 := GenS1{F0: genFn1(0, 118), F1: 127}
                _ = v74
            }
            var v75 bool = (32767 > v5)
            var v76 int64 = 1
            _ = v75
            _ = v76
        }
        _ = fi65
        _ = v66
    }
    fmt.Println(v1)
    fmt.Println(v1, v1)
    v77 := GenS4{F0: 42, F1: 0}
    if false {
        if true {
            var v78 int16 = 0
            var v79 bool = true
            var v80 uint16 = (v5 - v5)
            _ = v78
            _ = v79
            _ = v80
        } else {
            fmt.Println(v1, v1, v5)
            fmt.Println(v2, v1)
        }
        v81 := v4["a"]
        _ = v81
    } else {
        if true {
            if (65535 != 2) {
                if (-1 < -1) {
                    var v82 int32 = 0
                    fmt.Println(v82)
                    delete(v4, "foo bar")
                    _ = v82
                }
            } else {
                fmt.Println(v2, v2)
                _ = len(v3)
            }
            var v83 uint64 = 4294967295
            _ = v83
        } else {
            v5 = v5
            var v84 uint32 = 0
            _ = v84
        }
    }
    if (!true) {
        v2 -= (v2 - v2)
    }
    v2 += 4294967295
    v2 += (^(v2))
    fmt.Println(v1)
    var v85 uint16 = (v5 ^ v5)
    var v86 float64 = ((v1 + v1) - v1)
    if len(v3) > 0 {
        v3[0] = 2
    }
    switch 2147483647 {
    case 1:
        v2 = genFn1(341, 315)
    default:
        v3 = append(v3, 57)
    }
    v1 = 0.0
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v85:", v85)
    fmt.Println("v86:", v86)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v77
    _ = v85
    _ = v86
}
