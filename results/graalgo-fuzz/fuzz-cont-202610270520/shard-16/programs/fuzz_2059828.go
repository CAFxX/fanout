package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int8; F1 uint32; F2 int16 }
type GenS4 struct { F0 int32; F1 uint }
type GenS5 struct { F0 string }
type GenS6 struct { F0 string; F1 uint32; F2 uint16 }
type GenS7 struct { F0 int8 }
func genFn0(p0 uint16, p1 int8, p2 uint8) uint8 {
    if true {
        fmt.Println(p1, p0)
    } else {
        for fi6 := 0; fi6 < 14; fi6++ {
            p1 += (p1 + (p1 ^ (p1 << ((255) & 63))))
            _ = fi6
        }
        if (true || false) {
            p1 = (p1 << ((564) & 63))
            v7 := []int{}
            for fi8 := 0; fi8 < 11; fi8++ {
                p2 = p2
                switch "\n\t" {
                case "":
                    for fi9 := 0; fi9 < 10; fi9++ {
                        var v10 int32 = -2147483648
                        _ = fi9
                        _ = v10
                    }
                }
                v11 := GenS1{F0: p1, F1: p0}
                _ = fi8
                _ = v11
            }
            _ = v7
        }
    }
    fmt.Println(p2)
    var v12 bool = (true && (((173 >= 301) && (!(737 < 0))) && ((p0 ^ p0) >= uint16(p2))))
    p2 += p2
    if (p1 != (p1 << ((65535) & 63))) {
        var v13 uint32 = 468
        v12 = false
        if (0.0 >= 0.0) {
            v14 := []int{42, 0}
            v14 = append(v14, 32767)
            var v15 int64 = 329
            _ = v14
            _ = v15
        }
        _ = v13
    } else {
        if v12 {
            var v16 float64 = -2.25
            for fi17 := 0; fi17 < 11; fi17++ {
                var v18 int8 = -127
                v19 := []int{(-(fi17)), (fi17 >> ((65535) & 63)), (fi17 - (fi17 | (fi17 * fi17))), 9223372036854775807}
                v16 = 1e10
                _ = fi17
                _ = v18
                _ = v19
            }
            switch 421 {
            case 0:
                v16 = ((3.14159 - (1.5 * 1.5)) * ((v16 - v16) / -2.25))
            case -9223372036854775808:
                p2 = p2
            }
            _ = v16
        }
        var v20 uint32 = uint32(p2)
        var v21 int64 = -500
        _ = v20
        _ = v21
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    return p2
}
func genFn1() int64 {
    v22 := GenS1{F0: 42, F1: 1}
    for fi23 := 0; fi23 < 12; fi23++ {
        v24 := GenS1{F0: -128, F1: 32767}
        _ = fi23
        _ = v24
    }
    var v25 int32 = 2147483647
    _ = v22
    _ = v25
    return -1
}

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[int]bool)
    v3 := []uint32{42, 224}
    var v4 uint64 = 18446744073709551615
    var v5 int16 = -43
    _ = len(v1)
    var v26 int = 127
    _ = len(v3)
    v5 += (v5 * int16(v4))
    switch 32767 {
    case 1:
        if true {
            v5 += (v5 << ((255) & 63))
        } else {
            switch v26 {
            case 32768:
                switch "\n\t" {
                case "a":
                    for fi27 := 0; fi27 < 7; fi27++ {
                        switch ("hello" + "foo bar") {
                        case "hello":
                            if false {
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                if (65535 == 1) {
                                    var v28 uint8 = 128
                                    if ((v4 & v4) == (v4 / (((v4 + (v4 / (((v4 << ((2) & 63))) | 1)))) | 1))) {
                                        for fi29 := 0; fi29 < 8; fi29++ {
                                            var v30 uint64 = (v4 * (v4 % (((v4 % ((v4) | 1))) | 1)))
                                            switch (^(v26)) {
                                            case -632:
                                                switch (fi27 << ((255) & 63)) {
                                                case -2147483648:
                                                    var v31 int8 = -1
                                                    _ = v31
                                                case -647:
                                                    var v32 float64 = 1.5
                                                    _ = v32
                                                case -988:
                                                    fmt.Println(v30)
                                                    v33 := GenS2{F0: -232}
                                                    fmt.Println(v30, v26)
                                                    _ = v33
                                                default:
                                                    if true {
                                                        switch 0 {
                                                        case 32768:
                                                            v5 += v5
                                                        }
                                                    } else {
                                                        var v34 bool = (((3.14159 * 1e10) != -2.25) && ((1 >= -128) && (false || ((1 < -312) && (0.1 >= -2.25)))))
                                                        var v35 int = fi27
                                                        _ = v34
                                                        _ = v35
                                                    }
                                                }
                                            case -9223372036854775808:
                                                v36 := GenS1{F0: -112, F1: 833}
                                                _ = v36
                                            case 2147483647:
                                                if (("\n\t" + "foo bar") >= "foo bar") {
                                                    if len(v3) > 0 {
                                                        v3[0] = 88
                                                    }
                                                    v4 = 0
                                                }
                                                v26 = fi27
                                            }
                                            if true {
                                                var v37 int64 = 51
                                                var v38 uint32 = 2
                                                fmt.Println(v5, v5)
                                                _ = v37
                                                _ = v38
                                            }
                                            _ = fi29
                                            _ = v30
                                        }
                                        var v39 uint32 = 4294967295
                                        if (-300 < -1) {
                                            v40 := GenS2{F0: (fi27 / ((v26) | 1))}
                                            v26 -= (fi27 & (fi27 % ((v26) | 1)))
                                            var v41 bool = (((v5 / (((v5 | (-(v5)))) | 1)) != (v5 & v5)) || true)
                                            _ = v40
                                            _ = v41
                                        }
                                        _ = v39
                                    }
                                    _ = v28
                                }
                                for fi42 := 0; fi42 < 1; fi42++ {
                                    v4 = (v4 & 2)
                                    v43 := GenS2{F0: v26}
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    _ = fi42
                                    _ = v43
                                }
                            } else {
                                for fi44 := 0; fi44 < 20; fi44++ {
                                    if true {
                                        v26 = fi27
                                        fmt.Println(fi27)
                                        v4 -= v4
                                    }
                                    var v45 int32 = 42
                                    var v46 string = ("foo bar" + "a")
                                    _ = fi44
                                    _ = v45
                                    _ = v46
                                }
                                var v47 int8 = 1
                                v47 = v47
                                _ = v47
                            }
                        }
                        v48 := GenS3{F0: 100, F1: 497, F2: v5}
                        fmt.Println(v4, v5, v4)
                        _ = fi27
                        _ = v48
                    }
                case "\n\t":
                    v5 -= (v5 & (v5 + v5))
                case "foo bar":
                    for fi49 := 0; fi49 < 5; fi49++ {
                        for fi50 := 0; fi50 < 17; fi50++ {
                            fmt.Println(fi49)
                            switch (("a" + "foo bar") + "hello") {
                            case "":
                                v26 -= (v26 | 32767)
                            default:
                                v3 = append(v3, 2)
                            }
                            v4 = v4
                            _ = fi50
                        }
                        _ = fi49
                    }
                }
            case -2147483648:
                if true {
                    switch ("\n\t" + ("foo bar" + ("foo bar" + "hello"))) {
                    case "a":
                        v4 += 18446744073709551615
                    }
                }
            case 344:
                var v51 uint16 = 255
                _ = v51
            default:
                if true {
                    var v52 int64 = 9223372036854775807
                    for fi53 := 0; fi53 < 14; fi53++ {
                        delete(v2, v26)
                        var v54 uint16 = 596
                        for fi55 := 0; fi55 < 6; fi55++ {
                            switch "foo bar" {
                            case "":
                                v26 += -425
                                switch (-(fi53)) {
                                case -413:
                                    fmt.Println(v5)
                                    switch fi55 {
                                    case 4294967295:
                                        v3 = append(v3, 42)
                                    default:
                                        v52 = (v52 >> ((0) & 63))
                                    }
                                    var v56 int16 = 42
                                    _ = v56
                                case 4294967295:
                                    v57 := GenS3{F0: 42, F1: 0, F2: v5}
                                    _ = v57
                                case 32768:
                                    v54 = v54
                                }
                            }
                            _ = fi55
                        }
                        _ = fi53
                        _ = v54
                    }
                    _ = v52
                }
            }
            _ = len(v1)
            var v58 float64 = 0.0
            _ = v58
        }
    default:
        v59 := GenS4{F0: 197, F1: 411}
        _ = v59
    }
    for fi60 := 0; fi60 < 1; fi60++ {
        v61 := GenS2{F0: (fi60 >> ((1) & 63))}
        var v62 int64 = 1
        _ = fi60
        _ = v61
        _ = v62
    }
    fmt.Println(v5)
    if false {
        var v63 uint64 = 0
        var v64 float32 = 0.0
        _ = v63
        _ = v64
    } else {
        v65 := v2[32767]
        _ = v65
    }
    v66 := GenS5{F0: "\n\t"}
    if false {
        var v67 uint16 = 0
        _ = v67
    } else {
        var v68 float32 = -2.25
        for fi69 := 0; fi69 < 18; fi69++ {
            _ = len(v2)
            _ = fi69
        }
        _ = v68
    }
    var v70 float64 = (1e10 / float64(-2.25))
    for fi71 := 0; fi71 < 4; fi71++ {
        if false {
            switch (((((("hello" + "\n\t") + "a") + "\n\t") + "hello") + ("\n\t" + "a")) + ("" + "\n\t")) {
            case "\n\t":
                switch v26 {
                case 0:
                    for fi72 := 0; fi72 < 6; fi72++ {
                        v73 := GenS2{F0: (^(v26))}
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        _ = fi72
                        _ = v73
                    }
                case 32767:
                    v74 := GenS6{F0: "", F1: 1, F2: 42}
                    _ = v74
                case 127:
                    var v75 float64 = ((((v70 * (v70 - (0.1 - ((v70 / v70) / v70)))) * 0.1) * 1e10) * v70)
                    _ = v75
                default:
                    for fi76 := 0; fi76 < 18; fi76++ {
                        v77 := v1[4294967295]
                        _ = fi76
                        _ = v77
                    }
                }
            }
        } else {
            var v78 int32 = 0
            v79 := GenS7{F0: int8(v26)}
            _ = v78
            _ = v79
        }
        fmt.Println(v70)
        v80 := GenS2{F0: fi71}
        _ = fi71
        _ = v80
    }
    var v81 string = "\n\t"
    fmt.Println(v5, v81)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v70:", v70)
    fmt.Println("v81:", v81)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v66
    _ = v70
    _ = v81
}
