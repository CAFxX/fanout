package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 bool; F1 string }
type GenS3 struct { F0 int16; F1 uint }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int64; F1 uint16; F2 string }
type GenS6 struct { F0 uint; F1 uint8; F2 int32 }
type GenS7 struct { F0 uint64; F1 int64 }
type GenS8 struct { F0 int32; F1 int64 }
func genFn0(p0 uint, p1 uint) uint8 {
    var v6 bool = ((497 == 1) && true)
    p0 += 2
    var v7 int = -356
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 138
}
func genFn1(p0 int, p1 uint) int16 {
    var v8 float64 = 1e10
    v9 := []int{(p0 << ((p1) & 63)), p0, p0}
    if true {
        for fi10 := 0; fi10 < 13; fi10++ {
            p0 = 9223372036854775807
            v11 := GenS1{F0: false}
            if (true && false) {
                switch ("foo bar" + "hello") {
                case "hello":
                    var v12 uint32 = 2147483647
                    _ = v12
                case "\n\t":
                    switch 4294967295 {
                    case 127:
                        var v13 int16 = -95
                        _ = v13
                    default:
                        var v14 bool = true
                        _ = v14
                    }
                case "":
                    var v15 uint64 = 111
                    _ = v15
                }
            }
            _ = fi10
            _ = v11
        }
    } else {
        var v16 int64 = 0
        for fi17 := 0; fi17 < 7; fi17++ {
            var v18 uint64 = 4294967295
            v16 += (^(v16))
            _ = fi17
            _ = v18
        }
        switch (-(p0)) {
        case 0:
            if false {
                v16 = 9223372036854775807
                if false {
                    var v19 uint32 = 4294967295
                    if (("\n\t" == "foo bar") || false) {
                        fmt.Println(v19, v16)
                        v20 := GenS1{F0: false}
                        if len(v9) >= 2 {
                            v9 = v9[1:len(v9)]
                        }
                        _ = v20
                    } else {
                        switch ("foo bar" + "foo bar") {
                        case "foo bar":
                            var v21 int8 = -1
                            _ = v21
                        case "a":
                            var v22 int64 = v16
                            _ = v22
                        default:
                            var v23 uint = (p1 | p1)
                            _ = v23
                        }
                    }
                    _ = v19
                }
            } else {
                switch (("hello" + "hello") + ("foo bar" + ("\n\t" + "hello"))) {
                case "\n\t":
                    p1 += p1
                case "hello":
                    v16 += (v16 * v16)
                }
                v24 := GenS1{F0: (v16 >= (v16 << ((130) & 63)))}
                switch "" {
                case "\n\t":
                    var v25 string = (("a" + "\n\t") + "a")
                    _ = v25
                }
                _ = v24
            }
        default:
            for fi26 := 0; fi26 < 1; fi26++ {
                for fi27 := 0; fi27 < 19; fi27++ {
                    switch p0 {
                    case 890:
                        p0 = fi26
                    default:
                        switch p0 {
                        case -1:
                            v28 := GenS1{F0: (1 == 0)}
                            _ = v28
                        case -2147483648:
                            p0 -= fi26
                        default:
                            v16 -= (v16 - v16)
                        }
                    }
                    if len(v9) >= 2 {
                        v9 = v9[1:len(v9)]
                    }
                    _ = fi27
                }
                var v29 int16 = 0
                _ = fi26
                _ = v29
            }
        }
        _ = v16
    }
    var v30 float64 = v8
    _ = p0
    _ = p1
    _ = v30
    _ = v8
    _ = v9
    return -844
}

func main() {
    v1 := make(map[uint32]int)
    v2 := []uint32{0, 65535, 2147483648, 1}
    v3 := make(map[string]int)
    var v4 float32 = 0.0
    var v5 uint32 = 543
    v1[v5] = 32767
    v31 := GenS2{F0: (-118 < 127), F1: "a"}
    if ((32767 > 0) || false) {
        if len(v2) > 0 {
            v2[0] = 0
        }
        if true {
            for fi32 := 0; fi32 < 8; fi32++ {
                for fi33 := 0; fi33 < 16; fi33++ {
                    _ = len(v2)
                    var v34 int64 = 0
                    v4 -= v4
                    fmt.Println(v5, v34, v34)
                    v5 = 0
                    _ = fi33
                    _ = v34
                }
                for fi35 := 0; fi35 < 15; fi35++ {
                    var v36 string = "hello"
                    var v37 uint16 = 16
                    v4 = 1e10
                    _ = fi35
                    _ = v36
                    _ = v37
                }
                for fi38 := 0; fi38 < 12; fi38++ {
                    var v39 uint32 = 4294967295
                    v40 := GenS3{F0: -32768, F1: 502}
                    if (65535 <= 809) {
                        delete(v3, "a")
                        v5 = v39
                        switch -9223372036854775808 {
                        case 740:
                            switch "a" {
                            case "":
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                            case "a":
                                if (false && true) {
                                    if false {
                                        var v41 float64 = -2.25
                                        v39 = (^(v39))
                                        var v42 uint = 4294967295
                                        _ = v41
                                        _ = v42
                                    }
                                }
                            case "foo bar":
                                if true {
                                    v5 = v5
                                }
                            default:
                                if (0 <= 1) {
                                    v5 = v5
                                    v4 -= (1e10 - ((v4 - v4) + (float32(-2.25) / 0.1)))
                                    switch (fi32 % ((32767) | 1)) {
                                    case 43:
                                        var v43 uint64 = 96
                                        _ = v43
                                    case 0:
                                        var v44 int8 = -128
                                        _ = v44
                                    case 798:
                                        switch (fi38 - fi32) {
                                        case -9223372036854775808:
                                            var v45 uint8 = 255
                                            _ = v45
                                        case -1:
                                            switch (fi32 << ((65535) & 63)) {
                                            case 128:
                                                fmt.Println(v4, fi38)
                                            case -708:
                                                var v46 int32 = 1
                                                switch ("a" + "\n\t") {
                                                case "foo bar":
                                                    v4 = float32(1e10)
                                                    var v47 uint32 = 0
                                                    _ = v47
                                                default:
                                                    fmt.Println(fi38, v4, fi32)
                                                }
                                                v46 -= v46
                                                _ = v46
                                            default:
                                                if (0.1 >= 0.0) {
                                                    v4 = v4
                                                    delete(v3, "hello")
                                                }
                                            }
                                        }
                                    }
                                }
                            }
                        case 9223372036854775807:
                            if true {
                                _ = len(v2)
                                if len(v2) > 0 {
                                    v2[0] = 636
                                }
                            }
                        case 133:
                            switch 1 {
                            case 1:
                                v39 += (v5 / (((v39 * v39)) | 1))
                                var v48 uint32 = (v39 - v5)
                                _ = v48
                            case 32767:
                                fmt.Println(fi32, fi38)
                            case 4294967295:
                                var v49 string = "a"
                                _ = v49
                            default:
                                v50 := GenS4{F0: -75}
                                _ = v50
                            }
                        default:
                            fmt.Println(v39, v5)
                        }
                        var v51 int32 = 1
                        v52 := GenS1{F0: (-518 <= -912)}
                        _ = v51
                        _ = v52
                    } else {
                        var v53 int8 = 1
                        v54 := GenS1{F0: (!(false || true))}
                        fmt.Println(v4, v5)
                        _ = v53
                        _ = v54
                    }
                    _ = fi38
                    _ = v39
                    _ = v40
                }
                _ = fi32
            }
            var v55 uint8 = 255
            _ = v55
        }
        var v56 float32 = ((((v4 - (v4 * v4)) * 0.1) - float32((0.1 - (1.5 * 0.1)))) / v4)
        _ = v56
    } else {
        v5 += v5
        v3["hello"] = 2147483647
        v4 += v4
    }
    for fi57 := 0; fi57 < 8; fi57++ {
        var v58 uint = 18446744073709551615
        fmt.Println(fi57, v5)
        _ = fi57
        _ = v58
    }
    var v59 uint = 0
    v60 := GenS5{F0: 163, F1: 430, F2: ("\n\t" + "\n\t")}
    for fi61 := 0; fi61 < 2; fi61++ {
        v5 -= v5
        for fi62 := 0; fi62 < 20; fi62++ {
            var v63 uint = v59
            v3[""] = fi62
            _ = fi62
            _ = v63
        }
        v2 = append(v2, uint32(v59))
        _ = fi61
    }
    var v64 int32 = 963
    for fi65 := 0; fi65 < 13; fi65++ {
        for fi66 := 0; fi66 < 13; fi66++ {
            v67 := GenS1{F0: (0 != 32767)}
            v5 = (v5 & v5)
            _ = fi66
            _ = v67
        }
        v4 += 3.14159
        _ = fi65
    }
    var v68 uint64 = 209
    fmt.Println(v59, v5)
    if true {
        v64 = 42
    } else {
        v69 := GenS6{F0: v59, F1: 0, F2: -2147483648}
        switch 9223372036854775807 {
        case 9223372036854775807:
            v70 := GenS1{F0: ((^(v68)) >= 1)}
            _ = v70
        default:
            v4 += (v4 * v4)
        }
        _ = v69
    }
    if true {
        v2 = append(v2, (^(v5)))
        switch (("a" + (("\n\t" + "\n\t") + "foo bar")) + "") {
        case "a":
            if (((0.1 + 1.5) * float64(0.0)) == (-2.25 / (1e10 + 0.0))) {
                for fi71 := 0; fi71 < 13; fi71++ {
                    v72 := v3[(("a" + ("\n\t" + "hello")) + "a")]
                    for fi73 := 0; fi73 < 12; fi73++ {
                        var v74 uint16 = 33
                        if (false || (v5 <= (v5 ^ uint32(v74)))) {
                            switch int(v68) {
                            case -348:
                                v75 := GenS7{F0: v68, F1: 1}
                                _ = v75
                            case 4294967295:
                                if (int16(v59) >= -1) {
                                    var v76 uint8 = 2
                                    _ = v76
                                } else {
                                    delete(v1, (v5 - v5))
                                    if false {
                                        v74 -= v74
                                    }
                                }
                            default:
                                var v77 uint16 = (v74 & (^(v74)))
                                _ = v77
                            }
                        }
                        _ = fi73
                        _ = v74
                    }
                    _ = fi71
                    _ = v72
                }
                if true {
                    for fi78 := 0; fi78 < 6; fi78++ {
                        var v79 uint = (^(v59))
                        _ = fi78
                        _ = v79
                    }
                    var v80 uint8 = 42
                    _ = v80
                } else {
                    v64 -= (v64 ^ v64)
                }
            } else {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                v64 -= -2147483648
            }
        }
    } else {
        if (!(!false)) {
            v81 := GenS8{F0: v64, F1: 9223372036854775807}
            _ = v81
        } else {
            v1[42] = 42
            for fi82 := 0; fi82 < 10; fi82++ {
                v64 = v64
                _ = fi82
            }
        }
    }
    v4 = (float32((3.14159 / 1.5)) - (((1.5 - v4) * v4) * (1e10 + v4)))
    var v83 uint16 = 282
    var v84 int64 = 765
    for fi85 := 0; fi85 < 8; fi85++ {
        var v86 uint8 = 1
        var v87 int32 = (-(v64))
        v5 -= (v5 / ((42) | 1))
        _ = fi85
        _ = v86
        _ = v87
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v59:", v59)
    fmt.Println("v64:", v64)
    fmt.Println("v68:", v68)
    fmt.Println("v83:", v83)
    fmt.Println("v84:", v84)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v59
    _ = v60
    _ = v64
    _ = v68
    _ = v83
    _ = v84
}
