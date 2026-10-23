package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int; F1 uint16; F2 uint8 }
type GenS3 struct { F0 int64; F1 uint64; F2 uint }
type GenS4 struct { F0 uint16; F1 int32 }
type GenS5 struct { F0 int8; F1 uint; F2 uint8 }
type GenS6 struct { F0 uint16; F1 int64; F2 string }
type GenS7 struct { F0 uint16; F1 int; F2 string }
type GenS8 struct { F0 uint }
type GenS9 struct { F0 int64; F1 int; F2 uint32 }
func genFn0(p0 int16, p1 int8, p2 string) string {
    v7 := []int{-252, -1}
    v8 := GenS1{F0: -563}
    var v9 bool = true
    if v9 {
        for fi10 := 0; fi10 < 18; fi10++ {
            p0 += 1000
            _ = fi10
        }
        if len(v7) > 0 {
            v7[0] = 4294967295
        }
        if v9 {
            for fi11 := 0; fi11 < 17; fi11++ {
                p2 = "hello"
                _ = fi11
            }
            if (686 > 0) {
                var v12 string = ((p2 + p2) + p2)
                for fi13 := 0; fi13 < 8; fi13++ {
                    p1 += (^(p1))
                    _ = fi13
                }
                v9 = (!(false && v9))
                _ = v12
            } else {
                v9 = v9
                if v9 {
                    var v14 int64 = 0
                    _ = v14
                } else {
                    var v15 int16 = p0
                    var v16 string = p2
                    _ = v15
                    _ = v16
                }
            }
            var v17 int8 = 53
            _ = v17
        }
    } else {
        if (2 < 18446744073709551615) {
            p1 = (p1 / (((p1 * (p1 / ((0) | 1)))) | 1))
            var v18 string = (p2 + ("" + p2))
            for fi19 := 0; fi19 < 20; fi19++ {
                fmt.Println(p2, v18, p2)
                v20 := GenS2{F0: 32768, F1: 255, F2: 22}
                _ = fi19
                _ = v20
            }
            _ = v18
        } else {
            var v21 uint32 = 4294967295
            _ = len(v7)
            _ = v21
        }
        var v22 int = 1
        _ = v22
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return p2
}
func genFn1(p0 uint64, p1 int64) uint {
    if false {
        p1 = p1
        var v23 uint32 = 198
        _ = v23
    }
    for fi24 := 0; fi24 < 18; fi24++ {
        fmt.Println(fi24)
        _ = fi24
    }
    var v25 string = "a"
    for fi26 := 0; fi26 < 4; fi26++ {
        v27 := GenS3{F0: p1, F1: p0, F2: 65535}
        _ = fi26
        _ = v27
    }
    _ = p0
    _ = p1
    _ = v25
    return 90
}

func main() {
    var v1 int16 = 42
    var v2 uint8 = 2
    v3 := make(map[string]uint32)
    v4 := make(map[int]int)
    var v5 uint16 = 0
    var v6 uint32 = 771
    switch 1 {
    case 32768:
        for fi28 := 0; fi28 < 13; fi28++ {
            v29 := []int{32767, fi28, (fi28 + int(v6))}
            _ = fi28
            _ = v29
        }
    case 622:
        var v30 string = "a"
        _ = v30
    case 847:
        var v31 uint32 = v6
        _ = v31
    default:
        var v32 float32 = 0.1
        _ = v32
    }
    v2 -= (v2 * (v2 / ((uint8(v5)) | 1)))
    v33 := []int{42, 42, 2147483647}
    if false {
        delete(v3, ("a" + "hello"))
        switch "\n\t" {
        case "a":
            _ = len(v3)
        case "":
            for fi34 := 0; fi34 < 14; fi34++ {
                var v35 int = fi34
                var v36 uint8 = (^(v2))
                _ = fi34
                _ = v35
                _ = v36
            }
        default:
            switch "" {
            case "\n\t":
                switch ("a" + "\n\t") {
                case "a":
                    switch 1 {
                    case -1:
                        if ((v6 ^ (v6 >> ((2) & 63))) < (v6 | 2)) {
                            v1 -= v1
                            var v37 int16 = (v1 & (^(v1)))
                            _ = v37
                        }
                    case 128:
                        var v38 string = ""
                        _ = v38
                    case -2147483648:
                        switch -9223372036854775808 {
                        case 127:
                            var v39 int16 = (v1 % ((-394) | 1))
                            _ = v39
                        default:
                            v40 := v4[-852]
                            _ = v40
                        }
                    }
                case "":
                    v41 := GenS1{F0: 1}
                    _ = v41
                case "\n\t":
                    var v42 float64 = 1.5
                    _ = v42
                }
            case "":
                var v43 int64 = 396
                _ = v43
            default:
                switch 599 {
                case 127:
                    v5 -= (v5 ^ (v5 & 0))
                case 42:
                    v5 -= v5
                default:
                    v1 = 1
                }
            }
        }
    }
    var v44 int = 128
    if true {
        v6 += 4294967295
        var v45 uint64 = 42
        fmt.Println(v2, v45)
        _ = v45
    }
    fmt.Println(v5, v1, v2)
    for fi46 := 0; fi46 < 2; fi46++ {
        var v47 string = "hello"
        v44 -= -9223372036854775808
        if (((v5 < 2) || (float32((-2.25 / ((float64(1.5) - (0.0 / float64(((((1e10 + 3.14159) - (1e10 - 3.14159)) + ((1.5 + (((((0.1 * 3.14159) / 1.5) / 3.14159) - (0.0 / 0.1)) / (1.5 - ((1e10 / (1e10 / 0.1)) / (1.5 * 1e10))))) + (1e10 / 0.1))) - 0.1)))) - ((float64(((1e10 + 3.14159) - float32(((float64((0.0 * 0.0)) + 1.5) * 0.0)))) * (((-2.25 + (float64(1.5) + 0.1)) - (-2.25 / 1.0)) * (0.0 * 0.0))) * 0.0)))) > 0.0)) || (0 == 2147483647)) {
            var v48 bool = (int64(v6) < 42)
            v49 := GenS4{F0: (v5 - (v5 % (((^(v5))) | 1))), F1: -1}
            var v50 bool = (42 >= 2147483647)
            _ = v48
            _ = v49
            _ = v50
        }
        _ = fi46
        _ = v47
    }
    var v51 uint64 = 2
    fmt.Println(v5, v51)
    var v52 int64 = 9223372036854775807
    for fi53 := 0; fi53 < 15; fi53++ {
        _ = len(v33)
        _ = fi53
    }
    switch ("hello" + ("\n\t" + "")) {
    case "foo bar":
        if false {
            var v54 bool = (349 == (v5 + v5))
            _ = v54
        } else {
            switch ("foo bar" + ("" + "foo bar")) {
            case "":
                if false {
                    fmt.Println(v6, v6, v2)
                    var v55 float32 = (0.1 - ((float32(1e10) / float32(float64((float32(1e10) / 1.0)))) * (1e10 / 1.0)))
                    var v56 int64 = v52
                    _ = v55
                    _ = v56
                }
            case "\n\t":
                var v57 uint = 1
                _ = v57
            default:
                if ((^(v5)) < v5) {
                    for fi58 := 0; fi58 < 11; fi58++ {
                        var v59 uint16 = v5
                        for fi60 := 0; fi60 < 13; fi60++ {
                            v61 := GenS2{F0: -9223372036854775808, F1: v59, F2: v2}
                            _ = fi60
                            _ = v61
                        }
                        _ = fi58
                        _ = v59
                    }
                }
            }
            var v62 int8 = 117
            for fi63 := 0; fi63 < 10; fi63++ {
                v62 -= v62
                var v64 int = fi63
                if (((1.5 + 1e10) <= (((3.14159 - 0.0) + -2.25) * 0.0)) || (0.0 == (1e10 + 1e10))) {
                    var v65 bool = false
                    if v65 {
                        for fi66 := 0; fi66 < 14; fi66++ {
                            var v67 uint32 = (v6 / ((v6) | 1))
                            switch "hello" {
                            case "\n\t":
                                var v68 uint8 = v2
                                if (false && (0.0 <= 1.5)) {
                                    v6 -= v67
                                } else {
                                    switch (^(fi63)) {
                                    case -1:
                                        if v65 {
                                            if (v51 < v51) {
                                                fmt.Println(v51, v68, v64)
                                                if (2 >= v67) {
                                                    v69 := GenS5{F0: (v62 << ((65535) & 63)), F1: 65535, F2: uint8(v6)}
                                                    if len(v33) > 0 {
                                                        v33[0] = fi66
                                                    }
                                                    _ = len(v33)
                                                    _ = v69
                                                }
                                                var v70 float32 = 3.14159
                                                _ = v70
                                            }
                                            _ = len(v4)
                                        } else {
                                            if (1000 > (-(v1))) {
                                                if (true && (!v65)) {
                                                    v4[(-(v64))] = fi66
                                                    _ = len(v33)
                                                    var v71 float32 = 1.5
                                                    v67 = (^(v67))
                                                    _ = v71
                                                } else {
                                                    var v72 int32 = 1
                                                    fmt.Println(v5)
                                                    v73 := GenS2{F0: v44, F1: v5, F2: (v2 * (v2 | 127))}
                                                    _ = v72
                                                    _ = v73
                                                }
                                                var v74 string = ""
                                                v52 = v52
                                                _ = v74
                                            }
                                        }
                                    case -2147483648:
                                        v75 := GenS6{F0: v5, F1: (-(v52)), F2: ("foo bar" + "\n\t")}
                                        _ = v75
                                    }
                                    v76 := GenS7{F0: (^(v5)), F1: (-(v44)), F2: "\n\t"}
                                    v6 -= v6
                                    _ = v76
                                }
                                _ = len(v33)
                                _ = v68
                            default:
                                switch fi63 {
                                case -26:
                                    v52 -= v52
                                case 434:
                                    v1 += (v1 % ((-577) | 1))
                                case 0:
                                    v5 = 32767
                                    if ((1 > 2147483647) || true) {
                                        var v77 int16 = v1
                                        switch "" {
                                        case "a":
                                            v78 := GenS8{F0: genFn1(uint64(v2), -9223372036854775808)}
                                            var v79 uint16 = v5
                                            v80 := GenS9{F0: v52, F1: (fi66 * (v64 * (v64 | fi66))), F2: 2147483648}
                                            _ = v78
                                            _ = v79
                                            _ = v80
                                        case "":
                                            fmt.Println(v77)
                                        }
                                        v67 += (v6 >> ((65535) & 63))
                                        _ = v77
                                    }
                                default:
                                    v52 -= v52
                                }
                                if (false && v65) {
                                    switch (fi63 >> ((18446744073709551615) & 63)) {
                                    case 127:
                                        var v81 bool = ((1.5 * (-2.25 + 1.5)) >= float32((3.14159 * float64(-2.25))))
                                        v82 := GenS9{F0: v52, F1: fi66, F2: (v67 & v6)}
                                        _ = len(v3)
                                        _ = v81
                                        _ = v82
                                    default:
                                        var v83 int8 = v62
                                        _ = v83
                                    }
                                    if len(v33) > 0 {
                                        v33[0] = fi63
                                    }
                                    var v84 uint16 = v5
                                    v65 = v65
                                    fmt.Println(v67, fi66, v65)
                                    var v85 int = 42
                                    var v86 uint8 = v2
                                    _ = v84
                                    _ = v85
                                    _ = v86
                                }
                                var v87 bool = (0.1 >= (1e10 + 0.0))
                                _ = v87
                            }
                            _ = fi66
                            _ = v67
                        }
                        fmt.Println(v62, v64, v64)
                    }
                    _ = v65
                }
                _ = fi63
                _ = v64
            }
            _ = v62
        }
    case "hello":
        v3["hello"] = v6
    default:
        var v88 bool = ((v52 == v52) || ((v44 % (((-(v44))) | 1)) != v44))
        _ = v88
    }
    for fi89 := 0; fi89 < 12; fi89++ {
        v52 = 1
        v5 = (v5 - v5)
        switch ("" + ("\n\t" + "a")) {
        case "a":
            v90 := GenS8{F0: 114}
            _ = v90
        case "hello":
            var v91 int32 = int32(v6)
            _ = v91
        case "\n\t":
            if ((true && true) && ((v2 ^ v2) <= 127)) {
                for fi92 := 0; fi92 < 3; fi92++ {
                    var v93 uint = genFn1(uint64(v6), (-(v52)))
                    _ = fi92
                    _ = v93
                }
                var v94 float32 = ((1e10 - -2.25) - 0.0)
                _ = v94
            } else {
                for fi95 := 0; fi95 < 18; fi95++ {
                    var v96 uint16 = v5
                    v51 += v51
                    if len(v33) > 0 {
                        v33[0] = fi89
                    }
                    _ = fi95
                    _ = v96
                }
                var v97 uint16 = (v5 - v5)
                for fi98 := 0; fi98 < 15; fi98++ {
                    v1 = (v1 + v1)
                    var v99 int32 = -314
                    _ = fi98
                    _ = v99
                }
                _ = v97
            }
        }
        _ = fi89
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v44
    _ = v5
    _ = v51
    _ = v52
    _ = v6
}
