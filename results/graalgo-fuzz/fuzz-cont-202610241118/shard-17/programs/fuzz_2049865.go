package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64; F2 uint8 }
type GenS2 struct { F0 string; F1 uint32; F2 bool }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int32; F1 int8; F2 uint16 }
type GenS6 struct { F0 uint32; F1 bool }
type GenS7 struct { F0 uint32; F1 bool }
type GenS8 struct { F0 uint16; F1 int32; F2 int8 }
func genFn0(p0 uint, p1 uint64, p2 bool) int64 {
    p0 += 4294967295
    p1 = p1
    v5 := GenS1{F0: 128, F1: 4294967295, F2: 128}
    v6 := GenS2{F0: "foo bar", F1: 0, F2: true}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return -539
}
func genFn1() int32 {
    var v7 int16 = 32767
    v8 := []int{128, -1}
    switch 32768 {
    case -286:
        fmt.Println(v7)
    default:
        for fi9 := 0; fi9 < 13; fi9++ {
            for fi10 := 0; fi10 < 12; fi10++ {
                fmt.Println(fi10, v7)
                var v11 int32 = 2147483647
                _ = fi10
                _ = v11
            }
            var v12 uint8 = 8
            _ = fi9
            _ = v12
        }
    }
    switch -1 {
    case -1:
        switch "" {
        case "":
            v13 := GenS2{F0: "a", F1: 0, F2: true}
            _ = v13
        default:
            var v14 int8 = 100
            _ = v14
        }
    }
    for fi15 := 0; fi15 < 6; fi15++ {
        v8 = append(v8, fi15)
        v7 += 201
        _ = fi15
    }
    _ = v7
    _ = v8
    return 2147483647
}

func main() {
    v1 := make(map[int]string)
    var v2 uint = 270
    v3 := make(map[int]int)
    v4 := []int{127, 32768}
    v16 := GenS2{F0: "foo bar", F1: 933, F2: (((61 <= 0) || false) || true)}
    var v17 int = 32768
    for fi18 := 0; fi18 < 9; fi18++ {
        switch 42 {
        case 2147483647:
            var v19 uint = (^(v2))
            _ = v19
        case -9223372036854775808:
            for fi20 := 0; fi20 < 13; fi20++ {
                var v21 bool = false
                _ = fi20
                _ = v21
            }
        case -591:
            var v22 uint8 = 1
            _ = v22
        }
        for fi23 := 0; fi23 < 12; fi23++ {
            switch (fi23 | -2147483648) {
            case 0:
                fmt.Println(v2, fi23, v2)
            case -39:
                switch "foo bar" {
                case "\n\t":
                    v24 := GenS1{F0: 840, F1: 273, F2: 1}
                    _ = v24
                default:
                    v25 := GenS1{F0: (fi23 & (^(v17))), F1: 2, F2: 127}
                    _ = v25
                }
            default:
                var v26 string = "\n\t"
                _ = v26
            }
            if (true || false) {
                var v27 int64 = -9223372036854775808
                v28 := GenS2{F0: "\n\t", F1: 467, F2: (-2147483648 <= genFn1())}
                _ = v27
                _ = v28
            } else {
                var v29 int8 = 1
                v17 -= 328
                _ = v29
            }
            _ = fi23
        }
        fmt.Println(v2, v2)
        _ = fi18
    }
    var v30 float64 = -2.25
    var v31 int8 = -73
    if (!((-2.25 != (3.14159 + (1.5 + (-2.25 / 1.0)))) || true)) {
        for fi32 := 0; fi32 < 7; fi32++ {
            v17 = v17
            _ = fi32
        }
        if false {
            v30 = v30
            for fi33 := 0; fi33 < 10; fi33++ {
                var v34 int64 = 42
                v35 := v3[v17]
                v36 := GenS2{F0: "", F1: 65535, F2: true}
                _ = fi33
                _ = v34
                _ = v35
                _ = v36
            }
            v37 := GenS3{F0: 1}
            _ = v37
        } else {
            var v38 float32 = 0.0
            var v39 float32 = float32(v30)
            var v40 int = 2147483647
            _ = v38
            _ = v39
            _ = v40
        }
    } else {
        switch v17 {
        case 9223372036854775807:
            v17 -= (v17 >> (((v2 << (((v2 / ((4294967295) | 1))) & 63))) & 63))
        default:
            if false {
                if true {
                    for fi41 := 0; fi41 < 2; fi41++ {
                        for fi42 := 0; fi42 < 14; fi42++ {
                            v43 := GenS4{F0: -9223372036854775808}
                            fmt.Println(v31, fi42)
                            var v44 uint = v2
                            _ = fi42
                            _ = v43
                            _ = v44
                        }
                        fmt.Println(v30)
                        if (0 >= -1) {
                            fmt.Println(fi41, fi41)
                            delete(v3, (fi41 + 474))
                        }
                        _ = fi41
                    }
                    var v45 uint32 = 0
                    _ = len(v4)
                    _ = v45
                } else {
                    v4 = append(v4, int(v31))
                    v2 = 42
                    if (1.5 < (3.14159 - 0.0)) {
                        var v46 int64 = 0
                        _ = v46
                    }
                }
                fmt.Println(v2)
            } else {
                fmt.Println(v31, v17)
                switch (v17 >> (((v2 >> (((v2 ^ v2)) & 63))) & 63)) {
                case 128:
                    var v47 uint = v2
                    _ = v47
                default:
                    var v48 int64 = 858
                    _ = v48
                }
                if len(v4) > 0 {
                    v4[0] = (v17 >> (((v2 + v2)) & 63))
                }
            }
        }
    }
    v17 = v17
    for fi49 := 0; fi49 < 10; fi49++ {
        for fi50 := 0; fi50 < 12; fi50++ {
            v17 = (-(fi50))
            var v51 int64 = -1
            v2 -= (^(v2))
            _ = fi50
            _ = v51
        }
        v52 := GenS1{F0: 0, F1: 56, F2: 0}
        var v53 string = ("" + "")
        _ = fi49
        _ = v52
        _ = v53
    }
    switch (v17 >> ((1) & 63)) {
    case 127:
        v17 = -9223372036854775808
    case -763:
        v54 := GenS2{F0: (("foo bar" + (((("a" + "foo bar") + ("hello" + (("foo bar" + "a") + ("" + ("" + "foo bar"))))) + "hello") + ("foo bar" + ((("" + ("foo bar" + "a")) + ((("a" + ("foo bar" + ("hello" + "hello"))) + "\n\t") + ("\n\t" + ""))) + "hello")))) + "foo bar"), F1: 2, F2: true}
        _ = v54
    case 0:
        for fi55 := 0; fi55 < 20; fi55++ {
            if (!true) {
                switch (^(fi55)) {
                case 10:
                    var v56 uint8 = 0
                    _ = v56
                case 397:
                    v17 = fi55
                case 1:
                    for fi57 := 0; fi57 < 4; fi57++ {
                        v58 := GenS5{F0: 0, F1: v31, F2: 342}
                        switch -1 {
                        case 642:
                            _ = len(v4)
                        case -2147483648:
                            v59 := GenS4{F0: 42}
                            _ = v59
                        case 4294967295:
                            if len(v4) > 0 {
                                v4[0] = (fi55 % ((fi55) | 1))
                            }
                        }
                        v30 -= (float64(float32((float64(-2.25) - (((v30 + v30) * 3.14159) / (v30 - (v30 * 3.14159)))))) / (v30 * v30))
                        _ = fi57
                        _ = v58
                    }
                }
                v4 = append(v4, (^(v17)))
                var v60 int32 = int32(v2)
                _ = v60
            }
            v61 := GenS6{F0: 1, F1: (!(0 <= 2))}
            var v62 uint16 = 42
            _ = fi55
            _ = v61
            _ = v62
        }
    default:
        if (-666 > 0) {
            v63 := GenS4{F0: genFn0(v2, 42, true)}
            _ = v63
        }
    }
    for fi64 := 0; fi64 < 16; fi64++ {
        v30 -= v30
        switch "" {
        case "hello":
            var v65 int32 = genFn1()
            _ = v65
        case "\n\t":
            v17 = fi64
        default:
            var v66 int8 = -117
            _ = v66
        }
        var v67 int32 = -47
        _ = fi64
        _ = v67
    }
    if true {
        if (0 > -32768) {
            v30 = 1.5
        }
        for fi68 := 0; fi68 < 1; fi68++ {
            if ((^(v2)) >= 415) {
                switch (v17 + v17) {
                case 1:
                    fmt.Println(fi68, v2, v2)
                case 4294967295:
                    v69 := GenS6{F0: 2147483647, F1: true}
                    _ = v69
                case 42:
                    v70 := GenS5{F0: 0, F1: (v31 + int8(fi68)), F2: 255}
                    _ = v70
                default:
                    if false {
                        if false {
                            switch -135 {
                            case 9223372036854775807:
                                v17 += v17
                            }
                            fmt.Println(v31, v17, v2)
                            if true {
                                switch "foo bar" {
                                case "a":
                                    var v71 bool = (false || true)
                                    _ = v71
                                case "hello":
                                    if false {
                                        switch v17 {
                                        case 32768:
                                            if ((2147483647 < 2147483647) && true) {
                                                switch v17 {
                                                case 32768:
                                                    fmt.Println(v2, fi68)
                                                case 4294967295:
                                                    v2 -= 1
                                                default:
                                                    var v72 int16 = 1
                                                    _ = v72
                                                }
                                                v73 := GenS7{F0: 0, F1: false}
                                                _ = v73
                                            }
                                        }
                                    }
                                }
                            }
                        } else {
                            if (42 == -357) {
                                var v74 float32 = ((0.1 + (1.5 + (0.1 - 0.0))) + -2.25)
                                _ = len(v4)
                                _ = v74
                            }
                            var v75 uint32 = 42
                            _ = v75
                        }
                    }
                }
            } else {
                if (int64(v17) <= 283) {
                    _ = len(v4)
                    v4 = append(v4, v17)
                } else {
                    v2 += 0
                    v2 = (v2 ^ (v2 - v2))
                }
            }
            v76 := GenS8{F0: 32767, F1: 0, F2: (^(v31))}
            _ = fi68
            _ = v76
        }
    }
    v30 = ((v30 * 1e10) + (v30 + 1.5))
    fmt.Println(v31, v17)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v4
}
