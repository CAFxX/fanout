package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8; F2 int }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int16; F1 int8 }
type GenS4 struct { F0 int8; F1 bool }
type GenS5 struct { F0 uint32; F1 string }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 int; F1 uint16; F2 int32 }
type GenS9 struct { F0 int16; F1 uint8; F2 int }
func genFn0() uint8 {
    switch "hello" {
    case "hello":
        if true {
            var v6 uint8 = 1
            if true {
                var v7 int8 = 1
                _ = v7
            }
            _ = v6
        }
    case "foo bar":
        if (true || false) {
            for fi8 := 0; fi8 < 6; fi8++ {
                v9 := []int{(fi8 % (((fi8 ^ fi8)) | 1)), fi8, -91}
                if (0 <= 0) {
                    var v10 uint = 2
                    _ = v10
                }
                var v11 string = "hello"
                _ = fi8
                _ = v11
                _ = v9
            }
            var v12 uint = 2
            var v13 bool = (true && true)
            _ = v12
            _ = v13
        } else {
            v14 := GenS1{F0: 32768, F1: -1, F2: 127}
            v15 := []int{}
            if true {
                if true {
                    for fi16 := 0; fi16 < 19; fi16++ {
                        for fi17 := 0; fi17 < 1; fi17++ {
                            if (!(true || (true && ((0 <= 104) && (4294967295 != 65535))))) {
                                switch (^(fi17)) {
                                case 1:
                                    if (3.14159 < (1e10 * 3.14159)) {
                                        _ = len(v15)
                                        var v18 float64 = 3.14159
                                        _ = v18
                                    } else {
                                        _ = len(v15)
                                    }
                                }
                            } else {
                                switch "foo bar" {
                                case "a":
                                    switch fi17 {
                                    case -1:
                                        var v19 uint64 = 2
                                        v20 := GenS1{F0: fi17, F1: -128, F2: fi17}
                                        _ = v19
                                        _ = v20
                                    case -955:
                                        var v21 int = fi17
                                        _ = v21
                                    }
                                }
                                v22 := GenS2{F0: ("hello" + ("foo bar" + ("hello" + "a")))}
                                var v23 uint32 = 2
                                var v24 string = "foo bar"
                                v24 = v24
                                _ = len(v15)
                                v24 = v24
                                _ = v22
                                _ = v23
                                _ = v24
                            }
                            if (0 <= 40) {
                                var v25 uint64 = 18446744073709551615
                                v25 += v25
                                _ = len(v15)
                                _ = v25
                            }
                            var v26 int16 = -1
                            _ = fi17
                            _ = v26
                        }
                        v27 := GenS3{F0: 1000, F1: -1}
                        var v28 int = fi16
                        _ = fi16
                        _ = v27
                        _ = v28
                    }
                    if len(v15) > 0 {
                        v15[0] = 4294967295
                    }
                } else {
                    switch "foo bar" {
                    case "hello":
                        for fi29 := 0; fi29 < 8; fi29++ {
                            _ = len(v15)
                            _ = fi29
                        }
                    case "\n\t":
                        var v30 int32 = -440
                        _ = v30
                    default:
                        switch -1 {
                        case -2147483648:
                            if len(v15) >= 2 {
                                v15 = v15[1:len(v15)]
                            }
                        }
                    }
                }
            }
            _ = v14
            _ = v15
        }
    default:
        switch -322 {
        case 128:
            var v31 float64 = 0.0
            _ = v31
        case 42:
            var v32 int = 2147483647
            _ = v32
        }
    }
    v33 := []int{42}
    if true {
        v33 = append(v33, 9223372036854775807)
        v34 := GenS4{F0: -128, F1: false}
        _ = v34
    }
    for fi35 := 0; fi35 < 12; fi35++ {
        var v36 int = fi35
        v36 += fi35
        _ = fi35
        _ = v36
    }
    _ = v33
    return 255
}

func main() {
    var v1 float32 = 0.1
    var v2 string = "foo bar"
    var v3 uint64 = 712
    var v4 uint = 508
    var v5 int32 = 0
    if true {
        fmt.Println(v4, v1)
    } else {
        v3 = 1
        v37 := GenS5{F0: 65535, F1: (v2 + (v2 + (((v2 + v2) + v2) + "a")))}
        var v38 string = "a"
        _ = v37
        _ = v38
    }
    fmt.Println(v2)
    if false {
        v3 += 634
        var v39 uint32 = 2
        var v40 int = -1
        _ = v39
        _ = v40
    }
    v5 = -936
    v41 := []int{-668}
    v1 = v1
    for fi42 := 0; fi42 < 4; fi42++ {
        for fi43 := 0; fi43 < 4; fi43++ {
            switch fi42 {
            case 418:
                if (false || true) {
                    v1 -= (v1 / v1)
                    switch (v2 + v2) {
                    case "":
                        fmt.Println(v3, v2, v5)
                    case "hello":
                        fmt.Println(v1)
                    case "foo bar":
                        var v44 int8 = -1
                        _ = v44
                    default:
                        switch (v2 + v2) {
                        case "a":
                            fmt.Println(fi43, v2, v4)
                        case "hello":
                            if false {
                                fmt.Println(fi43)
                                fmt.Println(fi42, fi43, v5)
                            }
                        case "foo bar":
                            var v45 uint = 342
                            _ = v45
                        default:
                            v3 -= v3
                        }
                    }
                }
                v5 -= -2147483648
                _ = len(v41)
            default:
                var v46 uint = v4
                _ = v46
            }
            fmt.Println(v4, v2, fi43)
            v3 = (^(v3))
            if len(v41) >= 2 {
                v41 = v41[1:len(v41)]
            }
            if (false || false) {
                if false {
                    v1 = float32(1.5)
                    v5 += 1
                    v47 := GenS6{F0: v5}
                    fmt.Println(fi43)
                    _ = v47
                }
                if len(v41) >= 2 {
                    v41 = v41[1:len(v41)]
                }
                v48 := GenS7{F0: (^(v4))}
                switch (fi42 ^ fi42) {
                case 1:
                    switch fi43 {
                    case 9223372036854775807:
                        var v49 uint8 = 128
                        _ = v49
                    case -9223372036854775808:
                        if false {
                            switch 127 {
                            case -178:
                                if true {
                                    v41 = append(v41, fi42)
                                } else {
                                    if (v5 >= 843) {
                                        v4 += v4
                                        var v50 int16 = 32767
                                        _ = v50
                                    } else {
                                        v51 := GenS7{F0: v4}
                                        _ = v51
                                    }
                                    switch 128 {
                                    case 387:
                                        v41 = append(v41, (^(fi42)))
                                    default:
                                        if ((-14 > 42) && (float32(0.1) == ((v1 / v1) * ((1.5 / v1) / ((v1 - float32(float64(((v1 + ((v1 / (0.0 / (v1 * v1))) * v1)) / float32(3.14159))))) / v1))))) {
                                            if (!false) {
                                                v1 = 0.1
                                            }
                                            switch fi42 {
                                            case 891:
                                                var v52 float32 = 1.5
                                                _ = v52
                                            case 432:
                                                v2 = v2
                                            default:
                                                var v53 int64 = int64(fi42)
                                                _ = v53
                                            }
                                        }
                                    }
                                    v1 = (v1 / 0.1)
                                }
                            }
                            var v54 bool = false
                            _ = v54
                        }
                    case 42:
                        if len(v41) > 0 {
                            v41[0] = (fi43 << ((v4) & 63))
                        }
                    }
                case 2147483647:
                    if len(v41) >= 2 {
                        v41 = v41[1:len(v41)]
                    }
                default:
                    v55 := GenS5{F0: 2147483647, F1: ((v2 + v2) + v2)}
                    _ = v55
                }
                fmt.Println(v4, v2)
                _ = v48
            } else {
                v5 = 0
                v3 = (v3 ^ v3)
                v4 = v4
            }
            _ = fi43
        }
        v3 = 0
        _ = fi42
    }
    for fi56 := 0; fi56 < 8; fi56++ {
        var v57 uint8 = 128
        if true {
            v4 += (v4 * (v4 % ((922) | 1)))
            var v58 bool = false
            var v59 int32 = v5
            _ = v58
            _ = v59
        } else {
            if (-268 >= 32767) {
                for fi60 := 0; fi60 < 16; fi60++ {
                    var v61 uint64 = v3
                    _ = fi60
                    _ = v61
                }
            }
        }
        v5 = (v5 >> ((255) & 63))
        _ = fi56
        _ = v57
    }
    v41 = append(v41, 60)
    v62 := GenS8{F0: -2147483648, F1: 701, F2: 433}
    var v63 uint64 = 4294967295
    v64 := GenS2{F0: v2}
    v63 = (v63 << (((v4 | v4)) & 63))
    v5 = (v5 ^ (v5 / ((420) | 1)))
    fmt.Println(v5, v5, v63)
    for fi65 := 0; fi65 < 1; fi65++ {
        v41 = append(v41, (fi65 * fi65))
        for fi66 := 0; fi66 < 19; fi66++ {
            v1 += v1
            v67 := GenS7{F0: v4}
            var v68 uint16 = 32767
            _ = fi66
            _ = v67
            _ = v68
        }
        switch v2 {
        case "hello":
            var v69 int16 = -32768
            _ = v69
        }
        _ = fi65
    }
    v70 := GenS9{F0: -798, F1: 255, F2: 526}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v41:", len(v41), v41)
    fmt.Println("v5:", v5)
    fmt.Println("v63:", v63)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v5
    _ = v62
    _ = v63
    _ = v64
    _ = v70
}
