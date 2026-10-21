package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16 }
type GenS2 struct { F0 bool; F1 int8 }
type GenS3 struct { F0 uint16; F1 int }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint32; F1 int16; F2 int }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 int16 }
func genFn0() uint16 {
    var v7 float32 = (float32(float64(0.1)) * float32(-2.25))
    if false {
        for fi8 := 0; fi8 < 8; fi8++ {
            fmt.Println(fi8, v7)
            for fi9 := 0; fi9 < 7; fi9++ {
                if false {
                    var v10 uint16 = 527
                    _ = v10
                } else {
                    v11 := []int{fi9, 127, -908}
                    v12 := GenS1{F0: 1, F1: 42}
                    _ = v11
                    _ = v12
                }
                _ = fi9
            }
            _ = fi8
        }
        v7 = v7
    }
    v7 += (v7 / float32(0.1))
    if (!false) {
        v7 += v7
    } else {
        v7 += v7
        var v13 uint16 = 255
        v14 := []int{32768, -2147483648, 420, -122}
        _ = v13
        _ = v14
    }
    v7 = 1e10
    _ = v7
    return 1
}
func genFn1(p0 uint, p1 uint) uint32 {
    if (!false) {
        for fi15 := 0; fi15 < 11; fi15++ {
            switch "foo bar" {
            case "hello":
                if (!(!true)) {
                    v16 := GenS2{F0: false, F1: -29}
                    var v17 int32 = 443
                    _ = v16
                    _ = v17
                } else {
                    for fi18 := 0; fi18 < 17; fi18++ {
                        switch "a" {
                        case "foo bar":
                            var v19 int16 = 192
                            _ = v19
                        }
                        if false {
                            switch ("foo bar" + ((("foo bar" + "\n\t") + "hello") + "foo bar")) {
                            case "\n\t":
                                if true {
                                    fmt.Println(fi18, p1, fi15)
                                    v20 := GenS3{F0: 472, F1: fi18}
                                    _ = v20
                                }
                            case "a":
                                switch ("\n\t" + "foo bar") {
                                case "foo bar":
                                    v21 := []int{fi18, (fi18 >> (((^(p0))) & 63)), -1, fi18}
                                    _ = v21
                                case "a":
                                    p0 -= (p1 % ((p1) | 1))
                                }
                            }
                            fmt.Println(p0)
                            p1 = (p1 & p0)
                        }
                        _ = fi18
                    }
                }
            case "a":
                switch (("" + "\n\t") + ("a" + ("foo bar" + "hello"))) {
                case "":
                    p1 += (p1 >> ((1) & 63))
                case "foo bar":
                    var v22 int8 = 42
                    _ = v22
                }
            case "\n\t":
                fmt.Println(fi15)
            default:
                fmt.Println(p0)
            }
            if true {
                if (true || true) {
                    p0 -= 2
                    v23 := GenS1{F0: 383, F1: 1}
                    _ = v23
                }
                var v24 int8 = 1
                _ = v24
            }
            var v25 int8 = 127
            _ = fi15
            _ = v25
        }
        switch "hello" {
        case "":
            var v26 int = 128
            _ = v26
        case "hello":
            v27 := GenS3{F0: uint16(p1), F1: 9223372036854775807}
            _ = v27
        default:
            v28 := GenS3{F0: 550, F1: -9223372036854775808}
            _ = v28
        }
    } else {
        p1 = 42
        for fi29 := 0; fi29 < 11; fi29++ {
            fmt.Println(p1, fi29)
            _ = fi29
        }
    }
    switch -850 {
    case 128:
        if false {
            v30 := []int{2147483647, 32768, 0}
            v30 = append(v30, 0)
            var v31 uint32 = 0
            _ = v30
            _ = v31
        }
    case -9223372036854775808:
        fmt.Println(p1)
    case -647:
        if (false || ((((^(p0)) <= 0) && true) || (true && (false && false)))) {
            p0 = p1
        }
    default:
        switch 9223372036854775807 {
        case -1:
            p1 -= 255
        }
    }
    v32 := []int{32768}
    if len(v32) >= 2 {
        v32 = v32[1:len(v32)]
    }
    if false {
        if true {
            for fi33 := 0; fi33 < 8; fi33++ {
                switch "hello" {
                case "hello":
                    v34 := GenS2{F0: false, F1: 1}
                    _ = v34
                default:
                    p1 = (p1 << ((18446744073709551615) & 63))
                }
                for fi35 := 0; fi35 < 9; fi35++ {
                    p0 = (p1 ^ (p0 & p1))
                    v36 := GenS3{F0: 299, F1: int(p1)}
                    _ = fi35
                    _ = v36
                }
                var v37 uint64 = 75
                _ = fi33
                _ = v37
            }
            p0 = p1
        } else {
            p1 += (p1 << (((p0 % ((p0) | 1))) & 63))
            v38 := GenS4{F0: 32768}
            _ = v38
        }
        fmt.Println(p1, p0)
        var v39 int64 = 533
        _ = v39
    }
    _ = p0
    _ = p1
    _ = v32
    return 0
}

func main() {
    v1 := []int{128, 616, -863, -628}
    var v2 uint = 636
    var v3 int16 = 0
    var v4 uint = v2
    v5 := make(map[int]string)
    var v6 float64 = 3.14159
    var v40 float64 = (float64((((float32((-2.25 + (((((-2.25 * 1.5) * 1.5) + float64(1e10)) - v6) + ((1.5 / v6) * ((v6 * v6) - (1.5 / v6)))))) * (float32(0.1) / (0.0 + ((-2.25 - ((-2.25 / 1.5) / 1.0)) * 3.14159)))) + 1e10) * 3.14159)) * v6)
    v4 -= (v2 & (^(v4)))
    for fi41 := 0; fi41 < 15; fi41++ {
        v42 := GenS5{F0: 1, F1: v3, F2: (fi41 % ((fi41) | 1))}
        v4 = v4
        v6 = (((v40 / (-2.25 - v6)) + float64(0.1)) + v6)
        _ = fi41
        _ = v42
    }
    if true {
        if (!false) {
            if false {
                for fi43 := 0; fi43 < 10; fi43++ {
                    switch "hello" {
                    case "":
                        v44 := GenS3{F0: 42, F1: (fi43 & (fi43 & (fi43 | -426)))}
                        _ = v44
                    default:
                        v4 -= (v2 + (v4 & (v2 % ((v2) | 1))))
                    }
                    _ = fi43
                }
            } else {
                for fi45 := 0; fi45 < 7; fi45++ {
                    if false {
                        v3 = 712
                    } else {
                        v3 = v3
                    }
                    v40 += v40
                    if (true || false) {
                        for fi46 := 0; fi46 < 1; fi46++ {
                            switch (fi46 + fi46) {
                            case -419:
                                var v47 uint64 = uint64(v2)
                                _ = v47
                            default:
                                switch "foo bar" {
                                case "\n\t":
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    v4 -= v4
                                    var v48 int8 = -128
                                    _ = v48
                                case "":
                                    v49 := GenS3{F0: 2, F1: (fi45 - fi46)}
                                    _ = v49
                                }
                            }
                            v50 := v5[(fi45 << (((^(v4))) & 63))]
                            if (661 != 18446744073709551615) {
                                _ = len(v5)
                                var v51 string = ""
                                var v52 string = v50
                                v1 = append(v1, 2147483647)
                                if (236 <= 128) {
                                    var v53 uint16 = 255
                                    _ = v53
                                }
                                _ = v51
                                _ = v52
                            } else {
                                _ = len(v5)
                            }
                            _ = fi46
                            _ = v50
                        }
                        v6 -= float64((0.1 * -2.25))
                        v3 += v3
                    }
                    _ = fi45
                }
                v2 = v2
            }
        } else {
            v4 = (v4 / ((461) | 1))
        }
        v5[-9223372036854775808] = ("" + "a")
        for fi54 := 0; fi54 < 20; fi54++ {
            delete(v5, 32767)
            for fi55 := 0; fi55 < 10; fi55++ {
                var v56 string = ("" + ("foo bar" + "hello"))
                var v57 uint64 = 42
                var v58 float32 = float32((v40 / v6))
                fmt.Println(v40)
                delete(v5, -1)
                _ = fi55
                _ = v56
                _ = v57
                _ = v58
            }
            _ = fi54
        }
    }
    var v59 int16 = v3
    switch (("foo bar" + ("foo bar" + "hello")) + ("hello" + "")) {
    case "hello":
        var v60 int64 = 564
        _ = v60
    case "a":
        fmt.Println(v2)
    default:
        v40 += 0.1
    }
    v1 = append(v1, 1)
    if ((4294967295 == 907) || true) {
        v61 := v5[-191]
        for fi62 := 0; fi62 < 15; fi62++ {
            for fi63 := 0; fi63 < 19; fi63++ {
                if false {
                    if (true && true) {
                        if false {
                            switch 44 {
                            case 693:
                                v2 -= v4
                            case -837:
                                v64 := GenS3{F0: 680, F1: (fi62 / (((fi62 << ((v4) & 63))) | 1))}
                                _ = v64
                            case 0:
                                v65 := GenS5{F0: 941, F1: (v59 >> ((65535) & 63)), F2: (fi62 + 32768)}
                                _ = v65
                            }
                            if false {
                                var v66 uint16 = 65535
                                v5[(fi63 | 1)] = (v61 + ("a" + ("foo bar" + (v61 + ((v61 + v61) + v61)))))
                                if true {
                                    var v67 int8 = 0
                                    _ = v67
                                }
                                _ = v66
                            } else {
                                fmt.Println(v3, fi62, v4)
                                if ((fi62 % (((fi63 + (fi62 * fi62))) | 1)) != (^(fi63))) {
                                    var v68 float64 = v40
                                    var v69 float32 = 1e10
                                    _ = v68
                                    _ = v69
                                } else {
                                    v40 = v40
                                }
                                v1 = append(v1, -877)
                            }
                            v70 := GenS4{F0: (fi63 % ((-187) | 1))}
                            v71 := GenS6{F0: 0}
                            if true {
                                v72 := GenS7{F0: (v59 ^ (v3 % (((v59 | (v59 & (^(v59))))) | 1)))}
                                switch (v61 + (v61 + ("foo bar" + ("\n\t" + v61)))) {
                                case "hello":
                                    var v73 float64 = ((v6 * -2.25) - v6)
                                    _ = v73
                                }
                                fmt.Println(v61, v6, v59)
                                _ = v72
                            }
                            _ = v70
                            _ = v71
                        } else {
                            v40 -= (v40 / 1e10)
                        }
                        if true {
                            v40 = -2.25
                        }
                        var v74 bool = (((3.14159 / 3.14159) - 0.1) != 1e10)
                        var v75 uint64 = 264
                        _ = v74
                        _ = v75
                    } else {
                        v3 += v3
                        v4 -= 4294967295
                    }
                    v76 := v5[128]
                    _ = v76
                } else {
                    if false {
                        v77 := GenS2{F0: true, F1: -32}
                        _ = v77
                    }
                }
                _ = fi63
            }
            _ = fi62
        }
        _ = v61
    }
    var v78 bool = false
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v59:", v59)
    fmt.Println("v6:", v6)
    fmt.Println("v78:", v78)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v5
    _ = v59
    _ = v6
    _ = v78
}
