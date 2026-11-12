package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16 }
type GenS2 struct { F0 uint64; F1 uint64; F2 bool }
type GenS3 struct { F0 uint16; F1 int32; F2 string }
type GenS4 struct { F0 int; F1 uint8 }
type GenS5 struct { F0 string; F1 uint8 }
type GenS6 struct { F0 uint16; F1 bool; F2 int64 }
type GenS7 struct { F0 string; F1 int16; F2 bool }
type GenS8 struct { F0 bool; F1 string; F2 uint8 }
type GenS9 struct { F0 bool }
type GenS10 struct { F0 string; F1 int64 }
type GenS11 struct { F0 int16; F1 uint16 }
func genFn0() int64 {
    var v5 int = 497
    v6 := GenS1{F0: 157, F1: 42}
    _ = v5
    _ = v6
    return 488
}

func main() {
    v1 := make(map[string]string)
    var v2 int8 = 42
    var v3 int32 = -1
    v4 := make(map[int]bool)
    var v7 int32 = -64
    if ((!(1 != 972)) || ((!(65535 != 1)) && (true && false))) {
        if (float64(float32(1.5)) >= ((0.1 * 1e10) / 1.0)) {
            for fi8 := 0; fi8 < 12; fi8++ {
                var v9 uint = 65535
                _ = fi8
                _ = v9
            }
        } else {
            var v10 int8 = (-(v2))
            _ = v10
        }
        switch (("\n\t" + (("" + "") + "hello")) + "foo bar") {
        case "\n\t":
            for fi11 := 0; fi11 < 2; fi11++ {
                v2 += v2
                var v12 int8 = (v2 + v2)
                var v13 uint16 = 692
                _ = fi11
                _ = v12
                _ = v13
            }
        case "hello":
            var v14 uint32 = 42
            _ = v14
        default:
            switch ("" + "a") {
            case "hello":
                v7 = v7
            case "\n\t":
                v2 = (v2 + (v2 | v2))
            case "a":
                v15 := GenS2{F0: 492, F1: 0, F2: ((v7 | -2147483648) < -919)}
                _ = v15
            default:
                var v16 uint = 4294967295
                _ = v16
            }
        }
        var v17 uint32 = 65535
        _ = v17
    }
    v18 := GenS3{F0: 133, F1: (v7 / ((v7) | 1)), F2: "foo bar"}
    v19 := []int{2147483647, 42}
    var v20 int16 = -917
    v20 = 832
    for fi21 := 0; fi21 < 16; fi21++ {
        var v22 uint64 = 1
        var v23 float32 = 1.5
        if true {
            if false {
                v20 -= 412
                v7 = (^(v3))
            }
        }
        _ = fi21
        _ = v22
        _ = v23
    }
    if len(v19) >= 2 {
        v19 = v19[1:len(v19)]
    }
    var v24 int16 = (v20 * -761)
    v25 := GenS1{F0: 42, F1: 1}
    if (!(1 > 704)) {
        if false {
            var v26 bool = ((^(v2)) == v2)
            _ = v26
        }
        var v27 float32 = (0.0 / ((1e10 / (-2.25 + (1.5 + ((-2.25 + (1.5 / 1e10)) - 0.0)))) / 1.5))
        _ = v27
    }
    fmt.Println(v3)
    if (4294967295 <= 18446744073709551615) {
        var v28 uint32 = 488
        v2 = (v2 + (v2 >> ((255) & 63)))
        _ = v28
    } else {
        if (!(9223372036854775807 != 2147483647)) {
            switch "hello" {
            case "hello":
                if ((-802 >= -77) && (!(true && true))) {
                    v29 := GenS4{F0: 128, F1: 127}
                    var v30 string = "hello"
                    var v31 bool = true
                    _ = v29
                    _ = v30
                    _ = v31
                }
            default:
                for fi32 := 0; fi32 < 18; fi32++ {
                    switch fi32 {
                    case 4294967295:
                        if (v3 >= (v3 & v7)) {
                            if (!false) {
                                var v33 float32 = float32((-2.25 / 1e10))
                                for fi34 := 0; fi34 < 11; fi34++ {
                                    v20 += v20
                                    _ = fi34
                                }
                                v19 = append(v19, (fi32 % ((fi32) | 1)))
                                _ = v33
                            } else {
                                var v35 uint8 = 42
                                _ = v35
                            }
                            v20 += v24
                        } else {
                            _ = len(v4)
                            for fi36 := 0; fi36 < 1; fi36++ {
                                switch "hello" {
                                case "foo bar":
                                    var v37 uint64 = 4294967295
                                    _ = v37
                                }
                                var v38 int16 = v24
                                var v39 int64 = -1
                                _ = fi36
                                _ = v38
                                _ = v39
                            }
                        }
                    case -216:
                        var v40 uint8 = 128
                        _ = v40
                    }
                    v3 -= -1
                    v24 -= int16(v3)
                    _ = fi32
                }
            }
            var v41 int32 = -80
            if (!((true || false) || false)) {
                if (!(1 > 4294967295)) {
                    v7 -= v41
                    var v42 int32 = v3
                    _ = v42
                }
            } else {
                v20 = (v20 >> ((65535) & 63))
            }
            _ = v41
        } else {
            var v43 uint8 = 127
            v2 += v2
            for fi44 := 0; fi44 < 8; fi44++ {
                if (!false) {
                    v45 := GenS5{F0: "foo bar", F1: v43}
                    if true {
                        fmt.Println(v20, v2, v20)
                    }
                    _ = v45
                }
                if false {
                    v43 -= v43
                    switch (fi44 ^ fi44) {
                    case -9223372036854775808:
                        v24 = 0
                    case -908:
                        var v46 float64 = 1e10
                        _ = v46
                    default:
                        for fi47 := 0; fi47 < 16; fi47++ {
                            v48 := GenS6{F0: 65535, F1: true, F2: -9223372036854775808}
                            v49 := GenS7{F0: ("hello" + "hello"), F1: v20, F2: true}
                            v2 += v2
                            _ = fi47
                            _ = v48
                            _ = v49
                        }
                    }
                } else {
                    v20 = v20
                }
                if true {
                    var v50 int16 = v24
                    for fi51 := 0; fi51 < 11; fi51++ {
                        if true {
                            v52 := GenS8{F0: ((^(v43)) == (v43 | (v43 - 1))), F1: "hello", F2: (v43 << ((4294967295) & 63))}
                            v3 -= (v7 | (v7 + (v3 << ((346) & 63))))
                            _ = v52
                        } else {
                            v53 := GenS6{F0: 255, F1: false, F2: -9223372036854775808}
                            var v54 uint8 = v43
                            _ = v53
                            _ = v54
                        }
                        v55 := GenS3{F0: 42, F1: (v3 / (((-(v7))) | 1)), F2: "foo bar"}
                        var v56 uint32 = 65535
                        _ = fi51
                        _ = v55
                        _ = v56
                    }
                    _ = v50
                } else {
                    for fi57 := 0; fi57 < 15; fi57++ {
                        v58 := v1["\n\t"]
                        var v59 int16 = 32767
                        v60 := v4[fi44]
                        _ = fi57
                        _ = v58
                        _ = v59
                        _ = v60
                    }
                    if ((v2 % ((42) | 1)) >= (v2 >> ((65535) & 63))) {
                        switch 4294967295 {
                        case 4294967295:
                            switch "foo bar" {
                            case "foo bar":
                                if true {
                                    if len(v19) >= 2 {
                                        v19 = v19[1:len(v19)]
                                    }
                                    v24 -= 32767
                                    switch "a" {
                                    case "foo bar":
                                        v7 += v7
                                    case "\n\t":
                                        v61 := GenS2{F0: 4294967295, F1: 18446744073709551615, F2: true}
                                        _ = v61
                                    case "":
                                        v3 += 2147483647
                                    default:
                                        if false {
                                            fmt.Println(fi44)
                                            switch (-(fi44)) {
                                            case 904:
                                                v43 = v43
                                            }
                                        }
                                    }
                                } else {
                                    v62 := GenS9{F0: true}
                                    if len(v19) > 0 {
                                        v19[0] = (^(fi44))
                                    }
                                    _ = v62
                                }
                            default:
                                switch (fi44 * fi44) {
                                case 727:
                                    for fi63 := 0; fi63 < 9; fi63++ {
                                        v24 += (v20 * (v20 * v20))
                                        _ = fi63
                                    }
                                case 32767:
                                    for fi64 := 0; fi64 < 7; fi64++ {
                                        if (745 > -263) {
                                            v43 = (v43 | 255)
                                            v24 = v20
                                        }
                                        if false {
                                            v43 += v43
                                            v65 := GenS10{F0: ("hello" + "a"), F1: -719}
                                            v66 := GenS11{F0: -1, F1: 255}
                                            if true {
                                                if len(v19) > 0 {
                                                    v19[0] = (fi64 ^ (fi64 + (fi44 >> ((961) & 63))))
                                                }
                                            }
                                            fmt.Println(v3, fi44)
                                            _ = v65
                                            _ = v66
                                        } else {
                                            var v67 int = (^(fi44))
                                            _ = v67
                                        }
                                        v68 := GenS3{F0: 65535, F1: (-(v3)), F2: "a"}
                                        var v69 uint = 800
                                        _ = fi64
                                        _ = v68
                                        _ = v69
                                    }
                                default:
                                    if len(v19) >= 2 {
                                        v19 = v19[1:len(v19)]
                                    }
                                }
                            }
                        case 1:
                            v43 = v43
                        default:
                            v70 := GenS5{F0: "foo bar", F1: (v43 >> ((18446744073709551615) & 63))}
                            _ = v70
                        }
                    }
                    switch "" {
                    case "\n\t":
                        fmt.Println(v2, v24, fi44)
                    }
                }
                _ = fi44
            }
            _ = v43
        }
        var v71 int32 = (v3 | v3)
        v72 := v4[543]
        _ = v71
        _ = v72
    }
    v20 -= -483
    var v73 int = 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v7:", v7)
    fmt.Println("v73:", v73)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v7
    _ = v73
}
