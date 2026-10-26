package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int; F2 string }
type GenS2 struct { F0 int32; F1 string }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint32; F1 int64 }
type GenS5 struct { F0 string; F1 uint8; F2 bool }
type GenS6 struct { F0 uint16; F1 uint64; F2 int16 }
type GenS7 struct { F0 uint8 }

func main() {
    v1 := []string{"hello", "a", "", ("\n\t" + ("foo bar" + ("\n\t" + ("" + "\n\t"))))}
    v2 := []int{32768, 32768, 339, 9223372036854775807}
    var v3 uint = 65535
    var v4 uint8 = uint8(v3)
    v5 := []uint32{2147483648}
    switch "a" {
    case "a":
        for fi6 := 0; fi6 < 14; fi6++ {
            fmt.Println(fi6)
            if false {
                switch 42 {
                case -1:
                    v4 = v4
                default:
                    fmt.Println(v4)
                }
                v7 := GenS1{F0: 1000, F1: fi6, F2: "\n\t"}
                _ = v7
            } else {
                for fi8 := 0; fi8 < 20; fi8++ {
                    v3 = 255
                    _ = fi8
                }
            }
            v4 -= (v4 << ((832) & 63))
            _ = fi6
        }
    default:
        var v9 uint = (^(v3))
        _ = v9
    }
    for fi10 := 0; fi10 < 3; fi10++ {
        v3 = v3
        _ = fi10
    }
    if (!(1 < 65535)) {
        for fi11 := 0; fi11 < 4; fi11++ {
            var v12 uint16 = 340
            var v13 uint32 = 42
            if false {
                var v14 int64 = int64(v13)
                if false {
                    if (false && ((^(v14)) <= (^(v14)))) {
                        for fi15 := 0; fi15 < 6; fi15++ {
                            var v16 int16 = -32768
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            fmt.Println(v4, v4, v4)
                            _ = fi15
                            _ = v16
                        }
                    } else {
                        v17 := GenS1{F0: -32768, F1: -1, F2: ("" + "foo bar")}
                        switch "hello" {
                        case "hello":
                            var v18 uint32 = (v13 + v13)
                            _ = v18
                        default:
                            var v19 float64 = -2.25
                            _ = v19
                        }
                        _ = v17
                    }
                    fmt.Println(v4, v14, v3)
                    var v20 int = (fi11 << ((187) & 63))
                    _ = v20
                } else {
                    for fi21 := 0; fi21 < 20; fi21++ {
                        v2 = append(v2, fi21)
                        fmt.Println(v12, fi11)
                        _ = fi21
                    }
                    if (("a" >= "hello") && (327 < (v3 >> (((v3 + (v3 + 0))) & 63)))) {
                        var v22 int16 = 1000
                        v23 := GenS1{F0: v22, F1: 546, F2: "\n\t"}
                        _ = v22
                        _ = v23
                    }
                }
                _ = v14
            }
            _ = fi11
            _ = v12
            _ = v13
        }
    } else {
        var v24 int32 = -2147483648
        _ = v24
    }
    if (2 != 32767) {
        var v25 int8 = -20
        v26 := GenS1{F0: 1000, F1: -1, F2: ("\n\t" + "foo bar")}
        v27 := GenS2{F0: 2147483647, F1: "foo bar"}
        _ = v25
        _ = v26
        _ = v27
    } else {
        v28 := GenS3{F0: -997}
        _ = v28
    }
    for fi29 := 0; fi29 < 1; fi29++ {
        var v30 uint8 = (v4 >> ((18446744073709551615) & 63))
        _ = fi29
        _ = v30
    }
    if true {
        var v31 uint = v3
        fmt.Println(v4, v4)
        _ = v31
    } else {
        var v32 int8 = 42
        var v33 int8 = (v32 | (v32 * 1))
        if ("hello" < "\n\t") {
            if len(v5) > 0 {
                v5[0] = 0
            }
            var v34 int32 = -1
            fmt.Println(v33)
            _ = v34
        } else {
            v3 = v3
        }
        _ = v32
        _ = v33
    }
    switch 952 {
    case 0:
        for fi35 := 0; fi35 < 10; fi35++ {
            switch "a" {
            case "\n\t":
                v4 = uint8(v3)
            }
            if (!(!false)) {
                if (!(3.14159 < 0.0)) {
                    if true {
                        var v36 string = "a"
                        if true {
                            var v37 int8 = 0
                            _ = v37
                        } else {
                            var v38 uint8 = v4
                            _ = v38
                        }
                        if (!true) {
                            v36 = ("\n\t" + v36)
                        } else {
                            v3 = (v3 & 824)
                            v39 := GenS3{F0: (fi35 % ((fi35) | 1))}
                            _ = v39
                        }
                        _ = v36
                    }
                    v40 := GenS4{F0: 2147483647, F1: 1}
                    fmt.Println(fi35)
                    _ = v40
                }
                if (661 < 84) {
                    var v41 float64 = (float64(1.5) - 1e10)
                    for fi42 := 0; fi42 < 18; fi42++ {
                        v3 -= (^(v3))
                        v3 += (v3 / ((v3) | 1))
                        _ = fi42
                    }
                    for fi43 := 0; fi43 < 13; fi43++ {
                        var v44 float64 = 0.0
                        var v45 uint16 = 666
                        _ = fi43
                        _ = v44
                        _ = v45
                    }
                    _ = v41
                }
                v4 -= (v4 % ((v4) | 1))
            }
            var v46 uint = 136
            _ = fi35
            _ = v46
        }
    case 67:
        if false {
            var v47 uint = v3
            _ = v47
        } else {
            for fi48 := 0; fi48 < 12; fi48++ {
                fmt.Println(fi48, v4, fi48)
                if ((v3 <= 2) || false) {
                    if (-1 >= 32767) {
                        var v49 float64 = ((1.5 * (3.14159 + 0.1)) - 0.1)
                        var v50 int = (fi48 << (((v3 + 4294967295)) & 63))
                        _ = v49
                        _ = v50
                    }
                    switch 0 {
                    case -2147483648:
                        fmt.Println(fi48, fi48, v3)
                    case 731:
                        if true {
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                        }
                    case -9223372036854775808:
                        var v51 int = 1
                        _ = v51
                    }
                    v52 := GenS4{F0: 2147483647, F1: 546}
                    _ = v52
                }
                v1 = append(v1, "")
                _ = fi48
            }
        }
    case 32767:
        v3 = v3
    }
    var v53 uint32 = 468
    if false {
        var v54 uint = (v3 ^ v3)
        if false {
            var v55 uint64 = 1
            _ = len(v5)
            _ = v55
        } else {
            var v56 int32 = 2147483647
            _ = v56
        }
        _ = v54
    } else {
        if false {
            v57 := GenS5{F0: (((("\n\t" + "a") + (("hello" + "") + "\n\t")) + "foo bar") + "a"), F1: v4, F2: true}
            if true {
                var v58 int = 4294967295
                v2 = append(v2, v58)
                var v59 int32 = -1
                _ = v58
                _ = v59
            }
            for fi60 := 0; fi60 < 3; fi60++ {
                var v61 float32 = (1.5 + (((0.1 / 3.14159) - ((1.5 - (((0.0 + 3.14159) / 1.0) / 3.14159)) + float32(1.5))) - ((((-2.25 / (3.14159 + 1e10)) - 0.0) - (((-2.25 * 3.14159) - ((((float32((-2.25 * float64(-2.25))) - 0.0) - 1e10) * (((1e10 + ((0.1 / 0.1) + (1.5 - (1.5 / 1.0)))) / (3.14159 + (3.14159 - (-2.25 * 1e10)))) - 1.5)) / 1e10)) * -2.25)) / 0.1)))
                _ = fi60
                _ = v61
            }
            _ = v57
        }
        var v62 int8 = -1
        _ = v62
    }
    fmt.Println(v3, v53, v3)
    var v63 uint = v3
    if len(v2) > 0 {
        v2[0] = 128
    }
    var v64 float64 = (((1.5 - 1.5) * 0.1) - 1.5)
    switch "\n\t" {
    case "a":
        for fi65 := 0; fi65 < 10; fi65++ {
            if ((v63 < 255) && (100 < 1)) {
                for fi66 := 0; fi66 < 18; fi66++ {
                    var v67 int64 = -1
                    v68 := GenS6{F0: 255, F1: uint64(v3), F2: 42}
                    _ = fi66
                    _ = v67
                    _ = v68
                }
                switch "foo bar" {
                case "\n\t":
                    if len(v2) > 0 {
                        v2[0] = fi65
                    }
                }
            } else {
                v4 += v4
                var v69 uint16 = 65535
                _ = v69
            }
            v53 -= 42
            var v70 uint64 = 18446744073709551615
            _ = fi65
            _ = v70
        }
    case "foo bar":
        if ((false || (-367 >= 1000)) && ((1 <= 0) && (0 >= -126))) {
            var v71 int64 = 42
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = v71
        }
    case "hello":
        v5 = append(v5, v53)
    default:
        v72 := GenS5{F0: "foo bar", F1: 0, F2: true}
        _ = v72
    }
    switch 1 {
    case -2147483648:
        if len(v5) > 0 {
            v5[0] = 42
        }
    case -9223372036854775808:
        if true {
            for fi73 := 0; fi73 < 9; fi73++ {
                v53 += 65535
                for fi74 := 0; fi74 < 10; fi74++ {
                    var v75 uint64 = 4294967295
                    _ = fi74
                    _ = v75
                }
                _ = fi73
            }
            if (!((("hello" + "") > (("a" + ("" + "a")) + "\n\t")) || false)) {
                var v76 uint = 517
                v63 += v3
                _ = v76
            }
            v77 := GenS7{F0: 42}
            _ = v77
        }
    case 128:
        for fi78 := 0; fi78 < 13; fi78++ {
            var v79 float32 = (float32(3.14159) - 1.5)
            v80 := GenS2{F0: 2147483647, F1: (("" + "a") + "a")}
            _ = fi78
            _ = v79
            _ = v80
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v53:", v53)
    fmt.Println("v63:", v63)
    fmt.Println("v64:", v64)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v53
    _ = v63
    _ = v64
}
