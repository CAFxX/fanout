package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int16; F1 bool; F2 uint }
type GenS3 struct { F0 uint32; F1 uint; F2 int16 }
type GenS4 struct { F0 int; F1 int32; F2 int32 }
type GenS5 struct { F0 uint8; F1 uint }
type GenS6 struct { F0 int32; F1 int64 }
type GenS7 struct { F0 int; F1 int32 }
type GenS8 struct { F0 int64; F1 int8; F2 int16 }
type GenS9 struct { F0 int64 }
type GenS10 struct { F0 uint64; F1 uint }
func genFn0(p0 int64) int8 {
    var v4 int64 = p0
    v5 := []int{-681, 682, -794}
    if true {
        for fi6 := 0; fi6 < 19; fi6++ {
            fmt.Println(p0)
            _ = fi6
        }
        switch 128 {
        case 32768:
            v5 = append(v5, 9223372036854775807)
        case 0:
            if ((!true) && false) {
                v7 := GenS1{F0: 231}
                _ = v7
            } else {
                var v8 int16 = 0
                v9 := GenS2{F0: v8, F1: false, F2: 398}
                _ = v8
                _ = v9
            }
        default:
            if ((false && (42 > -599)) || true) {
                var v10 float64 = (-2.25 / 3.14159)
                _ = v10
            }
        }
    }
    if false {
        if (348 <= 32767) {
            var v11 int8 = 1
            if (0 != -1) {
                v11 -= v11
                v12 := GenS3{F0: 4294967295, F1: 65535, F2: 272}
                var v13 string = "hello"
                _ = v12
                _ = v13
            }
            for fi14 := 0; fi14 < 19; fi14++ {
                for fi15 := 0; fi15 < 12; fi15++ {
                    fmt.Println(v11, v11)
                    if len(v5) > 0 {
                        v5[0] = fi15
                    }
                    _ = fi15
                }
                fmt.Println(p0)
                v16 := GenS4{F0: 980, F1: -648, F2: 2147483647}
                _ = fi14
                _ = v16
            }
            _ = v11
        }
        for fi17 := 0; fi17 < 2; fi17++ {
            for fi18 := 0; fi18 < 3; fi18++ {
                if len(v5) > 0 {
                    v5[0] = fi17
                }
                _ = fi18
            }
            _ = fi17
        }
        var v19 uint = 2
        _ = v19
    }
    _ = p0
    _ = v4
    _ = v5
    return 1
}

func main() {
    var v1 uint = 278
    v2 := []string{}
    var v3 string = ""
    switch 526 {
    case 32767:
        fmt.Println(v3, v3)
    case 2147483647:
        var v20 uint = v1
        _ = v20
    case 560:
        var v21 int16 = 42
        _ = v21
    default:
        v1 += v1
    }
    var v22 string = v3
    for fi23 := 0; fi23 < 10; fi23++ {
        v24 := GenS5{F0: 31, F1: (v1 ^ v1)}
        _ = fi23
        _ = v24
    }
    v25 := GenS6{F0: 2147483647, F1: 661}
    if (-1 >= -668) {
        if (true || false) {
            var v26 int16 = 1
            if (3.14159 < 3.14159) {
                v3 = "\n\t"
                v26 = v26
            } else {
                if true {
                    v3 = v3
                } else {
                    switch v3 {
                    case "\n\t":
                        var v27 int32 = -1
                        _ = v27
                    default:
                        var v28 int32 = 0
                        _ = v28
                    }
                    v26 = v26
                }
                switch v3 {
                case "":
                    if (0 != 294) {
                        var v29 int16 = 258
                        _ = v29
                    } else {
                        v26 += v26
                        v1 += 65535
                        if true {
                            v26 += v26
                            if (true || (!(83 >= 0))) {
                                var v30 float32 = 0.1
                                for fi31 := 0; fi31 < 19; fi31++ {
                                    v1 = v1
                                    var v32 int64 = 42
                                    var v33 uint8 = 2
                                    _ = fi31
                                    _ = v32
                                    _ = v33
                                }
                                _ = v30
                            } else {
                                var v34 int32 = 1
                                _ = len(v2)
                                _ = v34
                            }
                            var v35 uint8 = 1
                            _ = v35
                        } else {
                            var v36 int = -2147483648
                            var v37 float64 = 1.5
                            _ = v36
                            _ = v37
                        }
                    }
                case "hello":
                    if (42 >= 824) {
                        var v38 uint32 = 154
                        switch v3 {
                        case "\n\t":
                            var v39 int8 = -128
                            _ = v39
                        case "a":
                            v40 := GenS7{F0: -9223372036854775808, F1: -1}
                            _ = v40
                        case "foo bar":
                            var v41 uint64 = 42
                            _ = v41
                        default:
                            var v42 uint = v1
                            _ = v42
                        }
                        _ = v38
                    }
                case "\n\t":
                    if (!false) {
                        switch v22 {
                        case "hello":
                            v2 = append(v2, (((v22 + "hello") + "") + "a"))
                        case "foo bar":
                            var v43 bool = true
                            _ = v43
                        default:
                            v44 := GenS4{F0: -2147483648, F1: 0, F2: -1}
                            _ = v44
                        }
                        v3 = "foo bar"
                        if false {
                            _ = len(v2)
                        } else {
                            if (972 == -2147483648) {
                                v3 = v3
                                v26 = v26
                            } else {
                                v2 = append(v2, "foo bar")
                                var v45 int32 = 0
                                if (!false) {
                                    var v46 int32 = 0
                                    _ = v46
                                }
                                _ = v45
                            }
                        }
                    } else {
                        if false {
                            if len(v2) > 0 {
                                v2[0] = v22
                            }
                            if false {
                                fmt.Println(v1, v26, v22)
                            }
                            v47 := GenS8{F0: -1, F1: -80, F2: (v26 * (v26 / ((1) | 1)))}
                            _ = v47
                        }
                    }
                }
            }
            _ = v26
        }
    } else {
        var v48 uint32 = 4294967295
        v3 = "foo bar"
        _ = len(v2)
        _ = v48
    }
    _ = len(v2)
    v2 = append(v2, ("\n\t" + (v22 + ("\n\t" + v22))))
    for fi49 := 0; fi49 < 6; fi49++ {
        switch v3 {
        case "":
            var v50 int32 = -2147483648
            _ = v50
        case "hello":
            switch v22 {
            case "foo bar":
                fmt.Println(v22, v22)
            }
        case "\n\t":
            if ((v1 + v1) != (v1 << (((v1 >> (((v1 % (((^(v1))) | 1))) & 63))) & 63))) {
                for fi51 := 0; fi51 < 16; fi51++ {
                    var v52 bool = (903 == 0)
                    _ = fi51
                    _ = v52
                }
                fmt.Println(v1, v22, v1)
                if (0.1 >= ((0.0 + 3.14159) - -2.25)) {
                    if false {
                        for fi53 := 0; fi53 < 15; fi53++ {
                            v3 = (("" + "\n\t") + v22)
                            v1 = v1
                            _ = fi53
                        }
                    }
                    for fi54 := 0; fi54 < 10; fi54++ {
                        var v55 uint8 = 1
                        switch (fi54 | fi54) {
                        case 4294967295:
                            var v56 int8 = -1
                            _ = v56
                        }
                        v3 = ("a" + (v22 + v3))
                        _ = fi54
                        _ = v55
                    }
                    switch v22 {
                    case "":
                        var v57 bool = (100 == -122)
                        _ = v57
                    case "hello":
                        v58 := GenS9{F0: 1}
                        _ = v58
                    case "foo bar":
                        v59 := GenS3{F0: 351, F1: (v1 ^ v1), F2: 1000}
                        _ = v59
                    default:
                        fmt.Println(v1, fi49, v1)
                    }
                }
            } else {
                var v60 uint8 = 255
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                v3 = "a"
                _ = v60
            }
        default:
            for fi61 := 0; fi61 < 12; fi61++ {
                v3 = ""
                fmt.Println(v3)
                var v62 uint64 = 18446744073709551615
                _ = fi61
                _ = v62
            }
        }
        var v63 int64 = 1
        _ = fi49
        _ = v63
    }
    v3 = v22
    for fi64 := 0; fi64 < 1; fi64++ {
        if ((((^(v1)) == (v1 * 18446744073709551615)) || (9223372036854775807 < -870)) && (!true)) {
            var v65 float64 = -2.25
            switch fi64 {
            case 128:
                v1 -= v1
            case 1:
                v65 -= v65
            case 686:
                if (!true) {
                    var v66 uint8 = 0
                    var v67 int32 = 1
                    v65 = (3.14159 / float64(1e10))
                    _ = v66
                    _ = v67
                }
            }
            _ = v65
        }
        v68 := GenS10{F0: 42, F1: (v1 & v1)}
        v1 = v1
        _ = fi64
        _ = v68
    }
    fmt.Println(v3)
    for fi69 := 0; fi69 < 6; fi69++ {
        if (-2.25 != float32((1.5 * 0.0))) {
            if len(v2) > 0 {
                v2[0] = v3
            }
            v70 := GenS7{F0: (fi69 * fi69), F1: 884}
            _ = v70
        }
        if (float32(((1e10 - (-2.25 + (float64(float32(((float64(1.5) + 0.1) + -2.25))) / (1e10 + -2.25)))) - (3.14159 / float64(1.5)))) == (-2.25 / 1.0)) {
            fmt.Println(v22, v22, v3)
        }
        var v71 bool = false
        _ = fi69
        _ = v71
    }
    for fi72 := 0; fi72 < 15; fi72++ {
        var v73 int8 = 0
        _ = fi72
        _ = v73
    }
    fmt.Println(v1, v3)
    v1 = v1
    var v74 uint32 = 42
    var v75 int64 = 228
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v74:", v74)
    fmt.Println("v75:", v75)
    _ = v1
    _ = v2
    _ = v22
    _ = v25
    _ = v3
    _ = v74
    _ = v75
}
