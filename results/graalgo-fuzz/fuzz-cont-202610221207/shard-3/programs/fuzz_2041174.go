package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16; F2 int64 }
type GenS2 struct { F0 int64; F1 string }
type GenS3 struct { F0 uint; F1 uint16; F2 int }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 bool; F1 int32; F2 int }
type GenS6 struct { F0 int32; F1 int8; F2 uint32 }
func genFn0(p0 string, p1 int16) int16 {
    v4 := []int{int(p1), 32767}
    for fi5 := 0; fi5 < 2; fi5++ {
        v6 := GenS1{F0: 42, F1: (p1 | 32767), F2: 1}
        v4 = append(v4, fi5)
        if ((0.1 + 1.5) != (float64(-2.25) - float64((0.1 - float32((1.5 - -2.25)))))) {
            if len(v4) > 0 {
                v4[0] = 0
            }
            if len(v4) > 0 {
                v4[0] = fi5
            }
        }
        _ = fi5
        _ = v6
    }
    if (18446744073709551615 <= 4294967295) {
        switch p0 {
        case "":
            for fi7 := 0; fi7 < 15; fi7++ {
                fmt.Println(p0)
                var v8 uint16 = 2
                fmt.Println(p0, p0)
                _ = fi7
                _ = v8
            }
        }
        v9 := GenS2{F0: 9223372036854775807, F1: p0}
        var v10 uint64 = 2
        _ = v10
        _ = v9
    }
    v11 := GenS2{F0: -860, F1: p0}
    _ = p0
    _ = p1
    _ = v11
    _ = v4
    return p1
}
func genFn1() uint32 {
    var v12 uint = 296
    switch 9223372036854775807 {
    case 32767:
        v13 := []int{-9223372036854775808, -2147483648, -1}
        _ = v13
    case 1:
        v14 := []int{-9223372036854775808}
        _ = v14
    default:
        v12 -= v12
    }
    _ = v12
    return 212
}

func main() {
    v1 := []string{("hello" + "foo bar")}
    v2 := make(map[int]bool)
    v3 := make(map[string]uint32)
    if len(v1) > 0 {
        v1[0] = "hello"
    }
    if (true && (42 >= 159)) {
        for fi15 := 0; fi15 < 12; fi15++ {
            var v16 int = fi15
            v16 += (v16 + (v16 / (((v16 ^ v16)) | 1)))
            for fi17 := 0; fi17 < 18; fi17++ {
                v16 = -9223372036854775808
                v16 += fi15
                _ = fi17
            }
            _ = fi15
            _ = v16
        }
        var v18 float64 = 0.0
        v18 -= v18
        _ = v18
    }
    var v19 uint16 = 65535
    switch "a" {
    case "hello":
        var v20 float32 = (float32(float64(1e10)) / -2.25)
        _ = v20
    case "\n\t":
        v19 -= (v19 | v19)
    case "a":
        v3[("a" + "")] = 2
    default:
        switch 128 {
        case 493:
            fmt.Println(v19)
        case 32767:
            if true {
                v21 := GenS2{F0: 42, F1: "foo bar"}
                _ = v21
            } else {
                var v22 int16 = 1000
                fmt.Println(v22, v22)
                _ = v22
            }
        }
    }
    for fi23 := 0; fi23 < 7; fi23++ {
        v19 -= v19
        var v24 uint32 = 931
        _ = fi23
        _ = v24
    }
    for fi25 := 0; fi25 < 12; fi25++ {
        if (!false) {
            if (400 > 42) {
                if false {
                    if true {
                        v19 -= 0
                        switch "hello" {
                        case "\n\t":
                            v19 += v19
                        case "hello":
                            for fi26 := 0; fi26 < 3; fi26++ {
                                var v27 uint32 = 42
                                v27 -= (v27 * (v27 >> ((238) & 63)))
                                _ = fi26
                                _ = v27
                            }
                        default:
                            var v28 bool = (("foo bar" + "") >= "hello")
                            _ = v28
                        }
                    }
                    v19 = 42
                    var v29 int32 = 578
                    _ = v29
                } else {
                    v19 -= (v19 ^ (^(v19)))
                    for fi30 := 0; fi30 < 6; fi30++ {
                        v31 := GenS3{F0: 607, F1: 2, F2: fi25}
                        var v32 uint = 226
                        _ = fi30
                        _ = v31
                        _ = v32
                    }
                }
            }
            for fi33 := 0; fi33 < 16; fi33++ {
                switch fi33 {
                case 128:
                    if true {
                        var v34 float64 = 3.14159
                        _ = v34
                    }
                default:
                    if true {
                        if false {
                            var v35 uint16 = 345
                            v36 := GenS3{F0: 65535, F1: (^(v35)), F2: (^(fi33))}
                            v35 = (^(v19))
                            _ = v35
                            _ = v36
                        } else {
                            if false {
                                if true {
                                    var v37 bool = true
                                    var v38 float64 = 0.0
                                    v37 = ((^(v19)) >= 719)
                                    _ = v37
                                    _ = v38
                                }
                                _ = len(v1)
                                switch fi33 {
                                case -56:
                                    switch (fi33 & fi33) {
                                    case 1:
                                        if ((2 < 42) && false) {
                                            v19 += (v19 << ((2) & 63))
                                            var v39 int8 = 0
                                            var v40 int = (^(fi33))
                                            _ = v39
                                            _ = v40
                                        } else {
                                            switch fi25 {
                                            case 127:
                                                if true {
                                                    if false {
                                                        v19 -= (v19 + v19)
                                                        var v41 int32 = -2147483648
                                                        fmt.Println(fi25)
                                                        _ = v41
                                                    } else {
                                                        if (1.5 >= ((1.5 / -2.25) * 0.1)) {
                                                            var v42 uint = 65535
                                                            _ = v42
                                                        } else {
                                                            v43 := GenS3{F0: uint(v19), F1: 0, F2: 693}
                                                            if false {
                                                                switch "a" {
                                                                case "foo bar":
                                                                    fmt.Println(v19)
                                                                case "a":
                                                                    if (256 < 1) {
                                                                        var v44 uint64 = 4294967295
                                                                        var v45 uint16 = (v19 * (v19 & 0))
                                                                        var v46 string = ("foo bar" + (((("hello" + "") + "a") + "foo bar") + "a"))
                                                                        v46 = "\n\t"
                                                                        v46 = ((("" + v46) + (((v46 + (v46 + (v46 + "hello"))) + (v46 + v46)) + "")) + v46)
                                                                        _ = v44
                                                                        _ = v45
                                                                        _ = v46
                                                                    } else {
                                                                        if true {
                                                                            v47 := GenS4{F0: 127}
                                                                            v19 = v19
                                                                            _ = v47
                                                                        }
                                                                        switch ("hello" + "a") {
                                                                        case "foo bar":
                                                                            var v48 uint = 255
                                                                            _ = v48
                                                                        case "":
                                                                            var v49 int16 = 32767
                                                                            _ = v49
                                                                        case "a":
                                                                            v19 += v19
                                                                        }
                                                                        switch (fi33 << ((0) & 63)) {
                                                                        case 1:
                                                                            if ((v19 <= 2) && true) {
                                                                                v50 := GenS4{F0: -35}
                                                                                _ = v50
                                                                            } else {
                                                                                v51 := GenS3{F0: 42, F1: v19, F2: fi25}
                                                                                v52 := GenS3{F0: 0, F1: v19, F2: fi25}
                                                                                _ = v51
                                                                                _ = v52
                                                                            }
                                                                        default:
                                                                            fmt.Println(fi25, v19, fi33)
                                                                        }
                                                                    }
                                                                }
                                                                v53 := v2[fi25]
                                                                _ = v53
                                                            } else {
                                                                fmt.Println(fi33)
                                                            }
                                                            _ = v43
                                                        }
                                                        if len(v1) > 0 {
                                                            v1[0] = (("hello" + "a") + "foo bar")
                                                        }
                                                        var v54 int = -2147483648
                                                        _ = v54
                                                    }
                                                    if (uint32(v19) != 0) {
                                                        var v55 bool = ((fi25 < (fi33 * fi33)) && (uint8(v19) >= 178))
                                                        _ = v55
                                                    } else {
                                                        v56 := GenS1{F0: 2, F1: 1000, F2: 0}
                                                        _ = v56
                                                    }
                                                }
                                            }
                                            v57 := GenS2{F0: 42, F1: "hello"}
                                            _ = v57
                                        }
                                    }
                                case 32768:
                                    fmt.Println(fi25, v19)
                                case -436:
                                    if true {
                                        fmt.Println(fi33, fi25, fi33)
                                        fmt.Println(fi33, fi33, fi33)
                                        var v58 float32 = (((-2.25 - -2.25) - 1e10) + 3.14159)
                                        var v59 uint8 = 2
                                        v1 = append(v1, (("" + (("hello" + "\n\t") + (((("hello" + "") + ("" + ("\n\t" + "\n\t"))) + "a") + (("hello" + "\n\t") + "")))) + "foo bar"))
                                        _ = v58
                                        _ = v59
                                    }
                                default:
                                    if false {
                                        var v60 int64 = 9223372036854775807
                                        switch 0 {
                                        case 127:
                                            fmt.Println(v60, fi33)
                                        default:
                                            var v61 uint64 = 0
                                            _ = v61
                                        }
                                        _ = v60
                                    }
                                    var v62 int32 = 2147483647
                                    v63 := GenS3{F0: 255, F1: 42, F2: 128}
                                    _ = v62
                                    _ = v63
                                }
                            }
                        }
                    }
                    v19 += v19
                    v19 -= (v19 * 42)
                }
                _ = fi33
            }
        }
        if true {
            fmt.Println(fi25, fi25)
            for fi64 := 0; fi64 < 10; fi64++ {
                switch "foo bar" {
                case "\n\t":
                    v19 = (v19 + 954)
                }
                delete(v2, 32767)
                fmt.Println(v19, v19, fi64)
                _ = fi64
            }
            var v65 int16 = -32768
            _ = v65
        } else {
            for fi66 := 0; fi66 < 10; fi66++ {
                if (2147483647 >= uint32(v19)) {
                    var v67 float64 = 0.0
                    var v68 int16 = -429
                    _ = v67
                    _ = v68
                }
                var v69 float64 = 0.1
                var v70 uint = 0
                _ = fi66
                _ = v69
                _ = v70
            }
            if ((0.1 - 1e10) > -2.25) {
                v19 += v19
                v71 := GenS3{F0: 4294967295, F1: 255, F2: (^(fi25))}
                v19 += (v19 / ((v19) | 1))
                _ = v71
            } else {
                v19 -= (^(v19))
            }
            for fi72 := 0; fi72 < 4; fi72++ {
                v73 := GenS5{F0: (116 <= uint(v19)), F1: -2147483648, F2: (fi25 & 127)}
                v19 = v19
                _ = fi72
                _ = v73
            }
        }
        var v74 uint16 = (v19 - 928)
        _ = fi25
        _ = v74
    }
    fmt.Println(v19)
    v75 := GenS6{F0: -616, F1: 64, F2: 262}
    for fi76 := 0; fi76 < 15; fi76++ {
        fmt.Println(v19)
        _ = fi76
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v75
}
