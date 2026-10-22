package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8 }
type GenS2 struct { F0 int64; F1 uint8; F2 int16 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint64; F1 uint32; F2 uint16 }
type GenS5 struct { F0 uint; F1 bool }
type GenS6 struct { F0 uint; F1 uint32 }
type GenS7 struct { F0 uint }
func genFn0() bool {
    v5 := []int{880}
    var v6 int64 = 0
    switch 204 {
    case -55:
        switch "a" {
        case "\n\t":
            switch 1 {
            case 4294967295:
                v6 = v6
            case 127:
                var v7 int8 = -41
                _ = v7
            case 42:
                var v8 bool = true
                _ = v8
            }
        case "foo bar":
            var v9 uint8 = 45
            _ = v9
        default:
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        }
    default:
        switch ("" + "a") {
        case "\n\t":
            var v10 int16 = 1000
            _ = v10
        case "foo bar":
            var v11 uint = 18446744073709551615
            _ = v11
        case "":
            var v12 float32 = (float32(((1e10 * 0.0) + float64((-2.25 * 0.1)))) + (0.1 + (1e10 - (float32(float64((float32(0.1) + ((1.5 + 3.14159) - -2.25)))) / (((0.0 * 0.1) + ((1e10 / 1e10) - (1e10 + (0.1 / (3.14159 - (0.0 * ((0.1 - 3.14159) * 0.1))))))) * -2.25)))))
            _ = v12
        }
    }
    _ = v5
    _ = v6
    return true
}
func genFn1(p0 uint) uint32 {
    var v13 float64 = 1.5
    if false {
        switch ("hello" + "foo bar") {
        case "hello":
            for fi14 := 0; fi14 < 3; fi14++ {
                var v15 uint8 = 205
                switch "hello" {
                case "hello":
                    var v16 uint16 = 0
                    _ = v16
                }
                var v17 uint16 = 992
                _ = fi14
                _ = v15
                _ = v17
            }
        case "a":
            for fi18 := 0; fi18 < 10; fi18++ {
                fmt.Println(p0, v13)
                _ = fi18
            }
        }
    } else {
        if true {
            var v19 int = -2147483648
            if (!false) {
                v20 := []int{v19, v19}
                v13 += ((v13 - 0.0) / (float64((1e10 * (0.0 + 3.14159))) + 3.14159))
                v21 := GenS1{F0: 0, F1: 0}
                _ = v20
                _ = v21
            }
            _ = v19
        } else {
            v22 := []int{32767, int(p0)}
            v23 := GenS1{F0: 127, F1: -85}
            if (42 == 1) {
                if false {
                    _ = len(v22)
                    v13 -= (v13 / (1e10 + 0.1))
                } else {
                    v24 := GenS1{F0: -128, F1: -111}
                    var v25 uint = p0
                    _ = v24
                    _ = v25
                }
            }
            _ = v22
            _ = v23
        }
        if (true && false) {
            var v26 int64 = -1
            fmt.Println(v13, p0, p0)
            _ = v26
        }
        v27 := []int{42, -790, -39, -9223372036854775808}
        _ = v27
    }
    for fi28 := 0; fi28 < 17; fi28++ {
        for fi29 := 0; fi29 < 4; fi29++ {
            var v30 int8 = -90
            fmt.Println(v13, fi28, fi28)
            var v31 int32 = 1
            _ = fi29
            _ = v30
            _ = v31
        }
        _ = fi28
    }
    _ = p0
    _ = v13
    return 1
}

func main() {
    v1 := []uint32{}
    var v2 int32 = -2147483648
    v3 := make(map[string]int)
    var v4 int32 = v2
    var v32 uint32 = 65535
    v32 = (^(v32))
    var v33 int = 42
    for fi34 := 0; fi34 < 15; fi34++ {
        if (((-2.25 * 0.1) / -2.25) <= float64(-2.25)) {
            var v35 int = v33
            var v36 int32 = (v2 >> ((65535) & 63))
            if ("hello" > "a") {
                v37 := GenS2{F0: 0, F1: 98, F2: 32767}
                _ = v37
            } else {
                var v38 int16 = 709
                if false {
                    v39 := GenS2{F0: 42, F1: 128, F2: v38}
                    _ = v39
                } else {
                    v1 = append(v1, v32)
                    v3["foo bar"] = 432
                }
                fmt.Println(v35, v33)
                _ = v38
            }
            _ = v35
            _ = v36
        }
        if false {
            fmt.Println(v33, v33)
        } else {
            switch -2147483648 {
            case 9223372036854775807:
                var v40 bool = (-268 < 851)
                _ = v40
            }
            var v41 uint = 42
            for fi42 := 0; fi42 < 10; fi42++ {
                v2 -= (^(v4))
                if true {
                    var v43 int64 = 42
                    if false {
                        switch fi42 {
                        case 128:
                            if true {
                                v41 = v41
                            }
                        case 0:
                            var v44 int32 = v4
                            _ = v44
                        case 32767:
                            var v45 bool = true
                            _ = v45
                        default:
                            v3[((("" + (("a" + "hello") + ("\n\t" + "hello"))) + ("\n\t" + (("\n\t" + "foo bar") + "hello"))) + "\n\t")] = v33
                        }
                        var v46 float32 = 3.14159
                        switch "" {
                        case "a":
                            _ = len(v3)
                        case "foo bar":
                            if (!(((v43 / (((v43 + v43)) | 1)) < int64(v4)) && false)) {
                                switch "hello" {
                                case "hello":
                                    _ = len(v1)
                                    if (!true) {
                                        var v47 float64 = -2.25
                                        v48 := GenS3{F0: 0}
                                        _ = v47
                                        _ = v48
                                    }
                                    var v49 int32 = v2
                                    switch (v33 + (fi42 % ((v33) | 1))) {
                                    case 127:
                                        v4 = 0
                                    case 4294967295:
                                        var v50 bool = (255 > 808)
                                        _ = v50
                                    case 893:
                                        v51 := GenS4{F0: 0, F1: (v32 ^ v32), F2: 255}
                                        _ = v51
                                    default:
                                        var v52 int64 = (v43 << ((2) & 63))
                                        _ = v52
                                    }
                                    v53 := GenS1{F0: -43, F1: 42}
                                    v4 += v49
                                    if true {
                                        v43 -= int64(v4)
                                        fmt.Println(v41, v4)
                                        if len(v1) > 0 {
                                            v1[0] = genFn1(v41)
                                        }
                                        v32 += (v32 >> (((v41 % ((244) | 1))) & 63))
                                        var v54 uint = 4294967295
                                        _ = v54
                                    }
                                    switch "hello" {
                                    case "":
                                        switch (fi34 % ((fi42) | 1)) {
                                        case 2147483647:
                                            v41 += v41
                                        case -9223372036854775808:
                                            v1 = append(v1, genFn1(v41))
                                            v46 = 3.14159
                                            if (65535 <= 255) {
                                                if false {
                                                    if true {
                                                        v46 -= float32((1.5 + 3.14159))
                                                        if false {
                                                            switch (fi42 << (((v41 - (v41 - 255))) & 63)) {
                                                            case -2147483648:
                                                                var v55 int16 = -540
                                                                _ = v55
                                                            }
                                                            var v56 int16 = 1
                                                            _ = v56
                                                        } else {
                                                            var v57 string = "foo bar"
                                                            var v58 uint16 = 330
                                                            _ = v57
                                                            _ = v58
                                                        }
                                                    }
                                                }
                                                switch v33 {
                                                case -46:
                                                    var v59 float32 = v46
                                                    _ = v59
                                                default:
                                                    switch 32768 {
                                                    case 32767:
                                                        var v60 int16 = 0
                                                        _ = v60
                                                    case 9223372036854775807:
                                                        v61 := GenS5{F0: (v41 << (((v41 - v41)) & 63)), F1: false}
                                                        _ = v61
                                                    case -9223372036854775808:
                                                        fmt.Println(v41)
                                                    default:
                                                        var v62 uint8 = 120
                                                        _ = v62
                                                    }
                                                }
                                                var v63 int = fi34
                                                _ = v63
                                            }
                                        }
                                    case "hello":
                                        var v64 uint = (v41 - (v41 ^ (^(v41))))
                                        _ = v64
                                    case "a":
                                        var v65 int = -962
                                        _ = v65
                                    default:
                                        v2 -= (v49 - v4)
                                    }
                                    _ = v49
                                    _ = v53
                                default:
                                    v2 += v4
                                }
                            }
                        case "":
                            var v66 int = fi34
                            _ = v66
                        default:
                            v43 += (v43 % ((584) | 1))
                        }
                        _ = v46
                    } else {
                        var v67 int8 = 100
                        _ = len(v3)
                        if true {
                            v67 += 127
                            fmt.Println(fi42, v4)
                            var v68 string = "\n\t"
                            v69 := GenS3{F0: 953}
                            var v70 uint32 = (v32 / ((v32) | 1))
                            _ = v68
                            _ = v69
                            _ = v70
                        }
                        _ = v67
                    }
                    var v71 uint16 = 779
                    v1 = append(v1, v32)
                    var v72 float64 = 3.14159
                    _ = v43
                    _ = v71
                    _ = v72
                } else {
                    if (v32 <= uint32(v41)) {
                        switch fi42 {
                        case 128:
                            v73 := GenS5{F0: v41, F1: false}
                            _ = v73
                        case -996:
                            var v74 uint32 = genFn1((^(v41)))
                            _ = v74
                        case -523:
                            var v75 float32 = (0.0 - (-2.25 - ((-2.25 + ((0.0 + 3.14159) - 0.1)) - 0.0)))
                            _ = v75
                        }
                        if false {
                            if (true && (true && true)) {
                                var v76 int64 = 1
                                var v77 uint = 18446744073709551615
                                _ = v76
                                _ = v77
                            } else {
                                fmt.Println(v33, v2)
                                v32 = (v32 << (((v41 * (v41 + (v41 * v41)))) & 63))
                            }
                        } else {
                            v78 := GenS5{F0: 65535, F1: true}
                            var v79 int32 = 42
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            var v80 int32 = -483
                            _ = v78
                            _ = v79
                            _ = v80
                        }
                        if len(v1) > 0 {
                            v1[0] = uint32(v41)
                        }
                    }
                    v4 -= v2
                }
                _ = fi42
            }
            _ = v41
        }
        _ = fi34
    }
    switch 42 {
    case -9223372036854775808:
        v81 := GenS6{F0: 328, F1: v32}
        _ = v81
    case -1:
        v82 := GenS7{F0: 42}
        _ = v82
    }
    v4 = 2147483647
    fmt.Println(v33)
    for fi83 := 0; fi83 < 17; fi83++ {
        v32 = (v32 << ((65535) & 63))
        v4 -= 2147483647
        _ = fi83
    }
    fmt.Println(v33, v2, v33)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v4
}
