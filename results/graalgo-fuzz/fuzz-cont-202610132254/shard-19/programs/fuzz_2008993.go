package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32; F2 uint8 }
type GenS2 struct { F0 uint; F1 uint; F2 uint16 }
type GenS3 struct { F0 uint16; F1 int32; F2 int }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 int64; F1 bool; F2 int16 }
type GenS9 struct { F0 string; F1 uint64 }
type GenS10 struct { F0 int32; F1 uint8; F2 uint64 }
type GenS11 struct { F0 uint8; F1 uint16; F2 int }
type GenS12 struct { F0 uint8; F1 int32 }
type GenS13 struct { F0 bool }
type GenS14 struct { F0 uint64; F1 int64; F2 int16 }
type GenS15 struct { F0 int; F1 int8 }
type GenS16 struct { F0 uint; F1 uint64; F2 int32 }
type GenS17 struct { F0 uint64; F1 int8 }
func genFn0(p0 uint8, p1 uint8) uint16 {
    var v5 uint32 = 0
    v6 := GenS1{F0: 0, F1: v5, F2: 127}
    v5 = v5
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 65535
}
func genFn1() bool {
    var v7 uint8 = 2
    v7 = v7
    v7 = 127
    v7 -= v7
    _ = v7
    return true
}

func main() {
    var v1 uint8 = 49
    var v2 float64 = (0.0 / 3.14159)
    var v3 uint16 = 553
    var v4 int = 9223372036854775807
    v1 -= (v1 - (v1 * (v1 << ((0) & 63))))
    switch "hello" {
    case "":
        fmt.Println(v4, v1, v2)
    case "a":
        v8 := GenS2{F0: 1, F1: 913, F2: (v3 - v3)}
        _ = v8
    case "\n\t":
        switch v4 {
        case -2147483648:
            var v9 int16 = 0
            _ = v9
        default:
            if (true && true) {
                if (!((307 == 0) || (!("foo bar" != "hello")))) {
                    fmt.Println(v3, v3)
                    v10 := []int{(v4 / (((^(v4))) | 1))}
                    for fi11 := 0; fi11 < 5; fi11++ {
                        var v12 int64 = 763
                        _ = fi11
                        _ = v12
                    }
                    _ = v10
                } else {
                    var v13 uint = 962
                    switch v4 {
                    case 4294967295:
                        for fi14 := 0; fi14 < 14; fi14++ {
                            v15 := GenS3{F0: 2, F1: -1, F2: 32768}
                            switch -53 {
                            case 207:
                                v2 = (0.0 / v2)
                            case 1:
                                switch "" {
                                case "\n\t":
                                    for fi16 := 0; fi16 < 3; fi16++ {
                                        var v17 int16 = 42
                                        _ = fi16
                                        _ = v17
                                    }
                                default:
                                    v18 := GenS4{F0: 762}
                                    _ = v18
                                }
                            case 4294967295:
                                fmt.Println(v3, v1, v2)
                            }
                            _ = fi14
                            _ = v15
                        }
                    case 560:
                        switch v4 {
                        case -10:
                            var v19 int64 = 123
                            _ = v19
                        case -384:
                            var v20 int32 = -1
                            _ = v20
                        default:
                            v1 = v1
                        }
                    case 42:
                        var v21 string = ("hello" + ("\n\t" + "foo bar"))
                        _ = v21
                    default:
                        v22 := GenS5{F0: -1}
                        _ = v22
                    }
                    for fi23 := 0; fi23 < 20; fi23++ {
                        if (v13 < uint(v3)) {
                            v13 += (^(v13))
                            if false {
                                var v24 uint16 = (v3 * genFn0(v1, (^(v1))))
                                var v25 int = v4
                                if (!true) {
                                    v26 := GenS6{F0: v1}
                                    _ = v26
                                }
                                _ = v24
                                _ = v25
                            } else {
                                var v27 int64 = -286
                                v13 = v13
                                v1 -= 169
                                _ = v27
                            }
                            if (-987 < 165) {
                                for fi28 := 0; fi28 < 6; fi28++ {
                                    v29 := GenS7{F0: 42}
                                    v30 := []int{128, (-(fi28)), (v4 >> (((^(v13))) & 63)), fi28}
                                    switch "hello" {
                                    case "":
                                        if true {
                                            fmt.Println(v4)
                                        } else {
                                            if (42 <= 762) {
                                                switch (fi23 << (((^(v13))) & 63)) {
                                                case 4294967295:
                                                    var v31 int8 = 1
                                                    _ = v31
                                                default:
                                                    v32 := GenS4{F0: uint64(v1)}
                                                    _ = v32
                                                }
                                                v33 := GenS8{F0: 148, F1: true, F2: 347}
                                                _ = v33
                                            }
                                        }
                                    case "hello":
                                        var v34 string = ("\n\t" + "")
                                        _ = v34
                                    case "a":
                                        if ((1 <= -819) && false) {
                                            var v35 uint = 255
                                            _ = v35
                                        }
                                    default:
                                        if true {
                                            v4 -= 466
                                            switch ((((("a" + ("hello" + "a")) + "hello") + "hello") + "a") + "hello") {
                                            case "a":
                                                var v36 bool = true
                                                _ = v36
                                            case "foo bar":
                                                v37 := GenS9{F0: "foo bar", F1: 42}
                                                if true {
                                                    var v38 int64 = 42
                                                    fmt.Println(v38, fi28, v1)
                                                    if (!(((^(v13)) < (^(v13))) || (v3 > (v3 << (((v13 / ((v13) | 1))) & 63))))) {
                                                        fmt.Println(v1, v13)
                                                        if len(v30) > 0 {
                                                            v30[0] = fi23
                                                        }
                                                    } else {
                                                        if false {
                                                            v39 := GenS2{F0: 565, F1: (^(v13)), F2: 70}
                                                            _ = v39
                                                        } else {
                                                            if len(v30) > 0 {
                                                                v30[0] = fi28
                                                            }
                                                        }
                                                    }
                                                    _ = v38
                                                }
                                                v2 += float64(0.1)
                                                _ = v37
                                            case "hello":
                                                fmt.Println(v13, v3, v4)
                                            }
                                            var v40 uint32 = 593
                                            _ = v40
                                        }
                                    }
                                    _ = fi28
                                    _ = v29
                                    _ = v30
                                }
                                v41 := GenS10{F0: -2147483648, F1: v1, F2: 23}
                                v13 = uint(v3)
                                _ = v41
                            }
                        }
                        for fi42 := 0; fi42 < 4; fi42++ {
                            if false {
                                switch v4 {
                                case 128:
                                    v3 = genFn0(v1, v1)
                                case -1:
                                    switch (fi42 >> ((488) & 63)) {
                                    case 9223372036854775807:
                                        if (false || true) {
                                            v43 := GenS11{F0: (v1 * v1), F1: genFn0(v1, (^(v1))), F2: fi23}
                                            switch "foo bar" {
                                            case "hello":
                                                v2 = v2
                                            }
                                            var v44 int16 = -1
                                            _ = v43
                                            _ = v44
                                        }
                                    case 128:
                                        v45 := []int{fi23}
                                        _ = v45
                                    }
                                default:
                                    var v46 int = (fi42 << (((^(v13))) & 63))
                                    _ = v46
                                }
                                v3 = v3
                            } else {
                                v47 := []int{fi23}
                                var v48 int = fi42
                                _ = v47
                                _ = v48
                            }
                            v49 := GenS12{F0: v1, F1: -2147483648}
                            switch ("foo bar" + "\n\t") {
                            case "\n\t":
                                if true {
                                    var v50 int = (fi23 - 32767)
                                    _ = v50
                                } else {
                                    fmt.Println(fi42, fi42)
                                }
                            case "hello":
                                if (!(-25 > fi42)) {
                                    if (!false) {
                                        v1 = 146
                                        v51 := []int{}
                                        v1 += v1
                                        _ = v51
                                    } else {
                                        if false {
                                            v52 := GenS13{F0: (!false)}
                                            v53 := []int{(fi42 >> ((v13) & 63))}
                                            v1 += 2
                                            _ = v52
                                            _ = v53
                                        }
                                        switch (fi42 + (v4 - fi23)) {
                                        case 32768:
                                            if (("foo bar" > ("a" + "hello")) && false) {
                                                v54 := []int{}
                                                var v55 uint = (v13 ^ (^(v13)))
                                                _ = v54
                                                _ = v55
                                            } else {
                                                v56 := GenS6{F0: v1}
                                                var v57 int = v4
                                                v58 := []int{fi42}
                                                _ = v56
                                                _ = v57
                                                _ = v58
                                            }
                                        }
                                    }
                                    var v59 string = "foo bar"
                                    fmt.Println(v2, v13, fi23)
                                    _ = v59
                                } else {
                                    var v60 string = "foo bar"
                                    var v61 uint8 = (v1 % ((8) | 1))
                                    _ = v60
                                    _ = v61
                                }
                            }
                            _ = fi42
                            _ = v49
                        }
                        _ = fi23
                    }
                    _ = v13
                }
                v1 = (v1 | (v1 - (^(v1))))
                fmt.Println(v1)
            } else {
                var v62 int64 = -9223372036854775808
                switch "a" {
                case "hello":
                    fmt.Println(v4, v4)
                default:
                    fmt.Println(v62)
                }
                for fi63 := 0; fi63 < 12; fi63++ {
                    var v64 uint64 = 2
                    var v65 int64 = (v62 ^ v62)
                    switch (v4 * (-(fi63))) {
                    case 9223372036854775807:
                        v66 := GenS14{F0: v64, F1: 1000, F2: 1000}
                        _ = v66
                    case 42:
                        var v67 float64 = v2
                        _ = v67
                    }
                    _ = fi63
                    _ = v64
                    _ = v65
                }
                _ = v62
            }
        }
    }
    var v68 int8 = 0
    v4 += 32768
    var v69 int64 = 42
    v70 := []int{972, -1}
    v68 += v68
    v68 = (v68 | (v68 - (v68 << ((255) & 63))))
    switch (v4 ^ (v4 | int(v68))) {
    case 2147483647:
        var v71 uint = 612
        _ = v71
    case 128:
        for fi72 := 0; fi72 < 16; fi72++ {
            if (255 < v1) {
                v73 := GenS8{F0: (v69 + v69), F1: false, F2: 714}
                v68 -= v68
                for fi74 := 0; fi74 < 6; fi74++ {
                    var v75 int16 = -1
                    v2 = v2
                    _ = fi74
                    _ = v75
                }
                _ = v73
            } else {
                fmt.Println(v1, v68)
                for fi76 := 0; fi76 < 2; fi76++ {
                    v69 = v69
                    fmt.Println(v68)
                    _ = fi76
                }
                switch "hello" {
                case "":
                    if (!true) {
                        _ = len(v70)
                        v1 = 245
                        switch fi72 {
                        case 32767:
                            v4 = (fi72 & -2147483648)
                        case 32768:
                            v77 := GenS15{F0: (^(fi72)), F1: (v68 << ((531) & 63))}
                            _ = v77
                        case 68:
                            _ = len(v70)
                        default:
                            v78 := GenS2{F0: 995, F1: 301, F2: v3}
                            _ = v78
                        }
                    } else {
                        v79 := GenS16{F0: 950, F1: 2, F2: 42}
                        if len(v70) > 0 {
                            v70[0] = fi72
                        }
                        _ = v79
                    }
                case "\n\t":
                    v70 = append(v70, fi72)
                default:
                    switch v4 {
                    case 4294967295:
                        v70 = append(v70, 731)
                    case 826:
                        _ = len(v70)
                    default:
                        v70 = append(v70, (v4 % (((fi72 >> ((1) & 63))) | 1)))
                    }
                }
            }
            v80 := GenS17{F0: 0, F1: v68}
            v1 = 0
            _ = fi72
            _ = v80
        }
    case 0:
        if len(v70) >= 2 {
            v70 = v70[1:len(v70)]
        }
    default:
        v68 = v68
    }
    v2 = (0.0 - (1e10 * v2))
    if false {
        switch "a" {
        case "a":
            if len(v70) >= 2 {
                v70 = v70[1:len(v70)]
            }
        case "foo bar":
            if len(v70) >= 2 {
                v70 = v70[1:len(v70)]
            }
        default:
            var v81 uint64 = 50
            _ = v81
        }
        var v82 float32 = 0.1
        _ = v82
    } else {
        if true {
            v83 := GenS5{F0: 32767}
            _ = v83
        }
        var v84 uint32 = 2147483647
        _ = v84
    }
    v85 := GenS12{F0: v1, F1: 2147483647}
    if true {
        v86 := GenS5{F0: 1}
        fmt.Println(v1, v68, v2)
        v1 = v1
        _ = v86
    } else {
        v1 = 128
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v68:", v68)
    fmt.Println("v69:", v69)
    fmt.Println("slicelen v70:", len(v70), v70)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v68
    _ = v69
    _ = v70
    _ = v85
}
