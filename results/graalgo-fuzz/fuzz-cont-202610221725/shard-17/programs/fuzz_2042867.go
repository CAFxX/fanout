package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int; F2 int }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int32; F1 uint8; F2 uint8 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 uint8; F1 uint64; F2 uint16 }
type GenS7 struct { F0 uint32; F1 uint32; F2 bool }
func genFn0(p0 int, p1 string) int32 {
    for fi7 := 0; fi7 < 9; fi7++ {
        p1 = (p1 + p1)
        p0 += (fi7 + (^(p0)))
        p0 += (p0 << ((4294967295) & 63))
        _ = fi7
    }
    var v8 string = p1
    _ = p0
    _ = p1
    _ = v8
    return 763
}

func main() {
    var v1 float32 = 1e10
    var v2 string = ("a" + ((("" + ("\n\t" + ("\n\t" + "hello"))) + "\n\t") + "a"))
    var v3 uint32 = 747
    v4 := make(map[int]string)
    var v5 int16 = -32768
    v6 := []string{"", v2, v2}
    for fi9 := 0; fi9 < 5; fi9++ {
        v10 := GenS1{F0: 42, F1: (fi9 | (^(fi9))), F2: 2147483647}
        v1 = v1
        _ = fi9
        _ = v10
    }
    v2 = v2
    fmt.Println(v2)
    for fi11 := 0; fi11 < 7; fi11++ {
        var v12 int16 = v5
        _ = fi11
        _ = v12
    }
    delete(v4, 4294967295)
    v13 := GenS1{F0: 184, F1: 127, F2: 9223372036854775807}
    for fi14 := 0; fi14 < 10; fi14++ {
        v15 := GenS1{F0: 2147483647, F1: (fi14 + -2147483648), F2: (-(fi14))}
        v16 := GenS2{F0: 0}
        if (2 < 42) {
            if (392 != (v3 ^ (^(v3)))) {
                var v17 float32 = v1
                var v18 int16 = (v5 >> ((2) & 63))
                var v19 int8 = 42
                _ = v17
                _ = v18
                _ = v19
            } else {
                v20 := GenS3{F0: 0, F1: 255, F2: 154}
                if true {
                    v1 += v1
                    var v21 float32 = v1
                    switch ((v2 + ("" + "a")) + ("foo bar" + v2)) {
                    case "hello":
                        delete(v4, (fi14 / ((-83) | 1)))
                    case "\n\t":
                        for fi22 := 0; fi22 < 20; fi22++ {
                            if (4294967295 <= 255) {
                                v3 -= (v3 % (((^(v3))) | 1))
                            } else {
                                v2 = v2
                                if true {
                                    v23 := GenS3{F0: 42, F1: 1, F2: 0}
                                    if len(v6) >= 2 {
                                        v6 = v6[1:len(v6)]
                                    }
                                    v1 -= float32((0.1 / (1.5 - 0.0)))
                                    _ = v23
                                } else {
                                    v6 = append(v6, "\n\t")
                                    var v24 string = (v2 + (v2 + v2))
                                    _ = v24
                                }
                                var v25 bool = (((v5 >> ((0) & 63)) >= 32767) && true)
                                _ = v25
                            }
                            var v26 uint = 42
                            _ = fi22
                            _ = v26
                        }
                    default:
                        var v27 int = fi14
                        _ = v27
                    }
                    _ = v21
                }
                _ = v20
            }
            if true {
                v3 = 1
            } else {
                var v28 int16 = -32768
                v29 := GenS4{F0: 65535}
                var v30 uint = 387
                _ = v28
                _ = v29
                _ = v30
            }
            var v31 int64 = -135
            _ = v31
        } else {
            v5 -= v5
            if (true || true) {
                var v32 uint8 = 127
                if (!true) {
                    switch 128 {
                    case -9223372036854775808:
                        fmt.Println(fi14, v3, fi14)
                    case 128:
                        v32 = 255
                    case 127:
                        v32 = (v32 ^ v32)
                    }
                } else {
                    fmt.Println(fi14, v1, v5)
                    var v33 uint16 = 1
                    for fi34 := 0; fi34 < 6; fi34++ {
                        var v35 float64 = float64(((v1 + -2.25) + v1))
                        _ = fi34
                        _ = v35
                    }
                    _ = v33
                }
                _ = v32
            }
        }
        _ = fi14
        _ = v15
        _ = v16
    }
    switch "a" {
    case "hello":
        if (1 == -1) {
            if ((false || (2 == 241)) || (1 >= 1)) {
                v1 -= 1e10
                v5 += (-(v5))
            }
        } else {
            v5 = v5
        }
    default:
        for fi36 := 0; fi36 < 3; fi36++ {
            if len(v6) > 0 {
                v6[0] = v2
            }
            for fi37 := 0; fi37 < 2; fi37++ {
                v38 := GenS5{F0: (v5 << ((1) & 63))}
                _ = fi37
                _ = v38
            }
            _ = fi36
        }
    }
    v39 := GenS4{F0: 42}
    for fi40 := 0; fi40 < 12; fi40++ {
        switch fi40 {
        case 32768:
            var v41 uint16 = 2
            _ = v41
        case 128:
            for fi42 := 0; fi42 < 17; fi42++ {
                v43 := GenS4{F0: 18446744073709551615}
                if ((18446744073709551615 == 4294967295) || (!(false || (127 != 195)))) {
                    if false {
                        switch (fi40 >> ((18446744073709551615) & 63)) {
                        case 715:
                            if (2147483647 < 137) {
                                if false {
                                    delete(v4, fi40)
                                    if true {
                                        v1 = 0.0
                                        var v44 uint32 = (^(v3))
                                        _ = v44
                                    } else {
                                        v5 = 0
                                    }
                                }
                                switch fi42 {
                                case -9:
                                    if (float64(float32(float64((3.14159 - (((((float32(1e10) + 3.14159) - 1.5) * (1.5 - (((v1 / 3.14159) - (-2.25 * (1e10 + 0.0))) * 3.14159))) + (v1 + float32(float64(1e10)))) * 1e10))))) != -2.25) {
                                        v3 -= v3
                                    }
                                    fmt.Println(v3, v1)
                                    fmt.Println(fi40, v5, v2)
                                case 0:
                                    if len(v6) > 0 {
                                        v6[0] = ("hello" + ((v2 + "foo bar") + ((v2 + ("a" + ((v2 + ("" + v2)) + ("foo bar" + v2)))) + "a")))
                                    }
                                }
                                var v45 int64 = 9223372036854775807
                                _ = v45
                            } else {
                                v1 = (((float32((1e10 + ((1.5 / 1.0) * 1.5))) + v1) * 1e10) / (1e10 + 1.5))
                            }
                        case 592:
                            v3 = v3
                        default:
                            _ = len(v4)
                        }
                    }
                }
                v46 := GenS6{F0: 128, F1: 0, F2: 65535}
                v2 = "a"
                if false {
                    var v47 uint32 = v3
                    var v48 uint8 = 255
                    switch (fi42 + int(v5)) {
                    case 837:
                        var v49 string = (v2 + "hello")
                        _ = v49
                    case 0:
                        if false {
                            if true {
                                fmt.Println(v47, v5)
                                var v50 uint64 = 72
                                v4[(fi40 / ((fi40) | 1))] = ("" + ("foo bar" + "foo bar"))
                                var v51 int16 = (v5 ^ v5)
                                fmt.Println(v1)
                                var v52 bool = false
                                _ = v50
                                _ = v51
                                _ = v52
                            } else {
                                var v53 uint16 = 1
                                _ = v53
                            }
                            var v54 int64 = 9223372036854775807
                            _ = v54
                        }
                    default:
                        var v55 int16 = 0
                        _ = v55
                    }
                    _ = v47
                    _ = v48
                } else {
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                    var v56 float64 = (0.0 * -2.25)
                    var v57 int64 = 9223372036854775807
                    fmt.Println(v56, v1, v57)
                    v58 := v4[(fi40 >> ((255) & 63))]
                    v4[(fi40 * (^(fi40)))] = v58
                    _ = v56
                    _ = v57
                    _ = v58
                }
                _ = fi42
                _ = v43
                _ = v46
            }
        case 32767:
            for fi59 := 0; fi59 < 18; fi59++ {
                switch fi40 {
                case -2147483648:
                    v60 := v4[fi40]
                    _ = v60
                }
                fmt.Println(v3, v5, v5)
                fmt.Println(v5)
                v61 := GenS2{F0: 0}
                _ = fi59
                _ = v61
            }
        }
        switch ("hello" + "\n\t") {
        case "a":
            for fi62 := 0; fi62 < 12; fi62++ {
                v63 := GenS3{F0: -2147483648, F1: 42, F2: 2}
                _ = fi62
                _ = v63
            }
        }
        _ = fi40
    }
    v2 = v2
    var v64 int = 127
    v64 = (^(v64))
    v65 := GenS7{F0: (^(v3)), F1: 591, F2: true}
    v4[v64] = (v2 + "hello")
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v64:", v64)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v5
    _ = v6
    _ = v64
    _ = v65
}
