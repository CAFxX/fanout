package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint16; F1 uint32; F2 int }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 string; F1 int64 }
type GenS5 struct { F0 uint8; F1 bool }
type GenS6 struct { F0 bool; F1 bool }
type GenS7 struct { F0 int8; F1 string; F2 uint16 }
type GenS8 struct { F0 bool; F1 uint }
type GenS9 struct { F0 int; F1 int; F2 int8 }
func genFn0(p0 int, p1 int8) uint32 {
    p1 += p1
    if (0 > -400) {
        p0 -= p0
        v7 := GenS1{F0: 42}
        fmt.Println(p0)
        _ = v7
    } else {
        if (true && false) {
            for fi8 := 0; fi8 < 15; fi8++ {
                switch (((((((("hello" + ("hello" + "foo bar")) + "hello") + "a") + "hello") + ("a" + "")) + "") + ("a" + ("\n\t" + "foo bar"))) + ((("hello" + "foo bar") + "foo bar") + ("\n\t" + ("\n\t" + ("\n\t" + "foo bar"))))) {
                case "":
                    v9 := GenS1{F0: 65535}
                    _ = v9
                case "\n\t":
                    switch ("a" + ((("" + "hello") + ("foo bar" + "")) + ("foo bar" + ""))) {
                    case "a":
                        switch (^(fi8)) {
                        case 621:
                            v10 := []int{(^(p0)), p0}
                            _ = v10
                        }
                    case "":
                        switch (p0 << ((10) & 63)) {
                        case 42:
                            switch 992 {
                            case -745:
                                for fi11 := 0; fi11 < 2; fi11++ {
                                    fmt.Println(fi11, fi8)
                                    _ = fi11
                                }
                            }
                        case -1:
                            var v12 uint = 617
                            _ = v12
                        case 9223372036854775807:
                            for fi13 := 0; fi13 < 13; fi13++ {
                                v14 := []int{(-(p0))}
                                p1 = p1
                                _ = fi13
                                _ = v14
                            }
                        default:
                            var v15 uint = 4294967295
                            _ = v15
                        }
                    case "foo bar":
                        for fi16 := 0; fi16 < 8; fi16++ {
                            v17 := GenS1{F0: 0}
                            p0 += (fi8 ^ p0)
                            _ = fi16
                            _ = v17
                        }
                    default:
                        for fi18 := 0; fi18 < 6; fi18++ {
                            var v19 int8 = p1
                            p0 -= 32767
                            _ = fi18
                            _ = v19
                        }
                    }
                }
                _ = fi8
            }
        }
        v20 := GenS2{F0: 224, F1: 2147483647, F2: 2147483647}
        switch p0 {
        case -1:
            p1 -= p1
        case -160:
            if false {
                p1 = -32
            } else {
                var v21 string = ("hello" + "a")
                v22 := GenS1{F0: 486}
                _ = v21
                _ = v22
            }
        case 2147483647:
            switch ("hello" + "") {
            case "":
                v23 := GenS3{F0: -1}
                _ = v23
            case "\n\t":
                if (18446744073709551615 >= 292) {
                    if true {
                        v24 := []int{}
                        fmt.Println(p0, p0)
                        _ = v24
                    } else {
                        if false {
                            switch "" {
                            case "hello":
                                var v25 string = ("a" + ("hello" + ("\n\t" + "a")))
                                _ = v25
                            }
                            var v26 int32 = 42
                            _ = v26
                        }
                        var v27 int16 = -748
                        _ = v27
                    }
                    switch (^(p0)) {
                    case 34:
                        fmt.Println(p0, p0)
                    case 380:
                        fmt.Println(p0)
                    default:
                        fmt.Println(p0)
                    }
                    p1 = p1
                } else {
                    switch (p0 - p0) {
                    case 4294967295:
                        fmt.Println(p1)
                    case 128:
                        fmt.Println(p0, p0)
                    default:
                        v28 := []int{}
                        _ = v28
                    }
                    var v29 uint32 = 2147483648
                    _ = v29
                }
            case "hello":
                var v30 uint = 4294967295
                _ = v30
            }
        }
        _ = v20
    }
    p0 = (p0 + 4294967295)
    var v31 int32 = -268
    _ = p0
    _ = p1
    _ = v31
    return 1
}
func genFn1(p0 string) int64 {
    var v32 int32 = 0
    var v33 uint64 = 538
    switch 257 {
    case 32768:
        for fi34 := 0; fi34 < 14; fi34++ {
            for fi35 := 0; fi35 < 20; fi35++ {
                var v36 float32 = (float32(((((1.5 - 0.0) + (((1e10 * 0.0) * (float64((1e10 * float32(0.0))) - (-2.25 * 0.1))) + 1.5)) * 1.5) / 0.1)) / 1.0)
                _ = fi35
                _ = v36
            }
            var v37 uint32 = 2147483647
            _ = fi34
            _ = v37
        }
    }
    for fi38 := 0; fi38 < 8; fi38++ {
        v39 := GenS1{F0: 1}
        v40 := []int{fi38, (fi38 * fi38)}
        var v41 uint = 69
        _ = fi38
        _ = v39
        _ = v40
        _ = v41
    }
    _ = p0
    _ = v32
    _ = v33
    return 42
}

func main() {
    var v1 int32 = 0
    var v2 uint32 = 824
    var v3 uint = 255
    v4 := make(map[int]string)
    v5 := []uint32{(v2 >> ((v3) & 63)), v2, (v2 & v2), 2}
    var v6 uint64 = 2
    if true {
        v2 = (v2 - (v2 ^ 1))
    } else {
        for fi42 := 0; fi42 < 20; fi42++ {
            if (v6 >= 0) {
                v43 := GenS4{F0: "a", F1: -9223372036854775808}
                for fi44 := 0; fi44 < 1; fi44++ {
                    switch (-(fi44)) {
                    case 354:
                        var v45 int32 = (v1 >> (((^(v3))) & 63))
                        if (false || (!false)) {
                            v46 := GenS3{F0: 630}
                            v3 -= (v3 & (^(v3)))
                            _ = v46
                        }
                        var v47 int8 = 0
                        if false {
                            if false {
                                v1 -= v1
                                fmt.Println(v2)
                                v48 := GenS5{F0: 127, F1: (fi44 == (fi42 + 871))}
                                _ = v48
                            }
                        }
                        v49 := GenS6{F0: false, F1: false}
                        _ = v45
                        _ = v47
                        _ = v49
                    case 889:
                        if (255 >= 622) {
                            v50 := v4[128]
                            v51 := GenS7{F0: 0, F1: v50, F2: 0}
                            v52 := GenS8{F0: ((-2.25 + 1e10) == float64(float32(-2.25))), F1: v3}
                            _ = v50
                            _ = v51
                            _ = v52
                        }
                    }
                    fmt.Println(fi44, fi44, v2)
                    switch "\n\t" {
                    case "\n\t":
                        if true {
                            var v53 uint = v3
                            var v54 bool = (float64(0.1) >= 3.14159)
                            _ = v53
                            _ = v54
                        } else {
                            switch fi44 {
                            case 0:
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                            case -150:
                                if (-1 <= 0) {
                                    switch "\n\t" {
                                    case "a":
                                        var v55 uint8 = 144
                                        _ = v55
                                    default:
                                        var v56 uint8 = 255
                                        _ = v56
                                    }
                                    v6 = v6
                                } else {
                                    if (!(!((false || (9223372036854775807 < genFn1(""))) && (true && true)))) {
                                        var v57 int8 = 100
                                        _ = v57
                                    }
                                }
                            case 2147483647:
                                v3 -= 993
                            default:
                                _ = len(v4)
                            }
                            delete(v4, fi44)
                            v2 = 731
                        }
                    default:
                        v3 = (v3 / (((v3 % ((v3) | 1))) | 1))
                    }
                    _ = fi44
                }
                _ = v43
            }
            v3 += (v3 | (v3 - uint(v2)))
            fmt.Println(fi42)
            _ = fi42
        }
        v1 = v1
        if true {
            var v58 uint8 = 0
            fmt.Println(v58)
            if true {
                v5 = append(v5, v2)
            } else {
                v3 -= v3
            }
            _ = v58
        } else {
            for fi59 := 0; fi59 < 15; fi59++ {
                var v60 uint16 = 935
                _ = fi59
                _ = v60
            }
        }
    }
    if (true || (!true)) {
        var v61 int64 = 605
        _ = v61
    }
    var v62 string = "a"
    v63 := GenS1{F0: 357}
    for fi64 := 0; fi64 < 11; fi64++ {
        v65 := GenS8{F0: (-32768 < -32768), F1: 251}
        _ = fi64
        _ = v65
    }
    for fi66 := 0; fi66 < 1; fi66++ {
        v3 -= 4294967295
        _ = fi66
    }
    switch v62 {
    case "foo bar":
        v62 = v62
    case "a":
        v6 = v6
    case "\n\t":
        v5 = append(v5, (v2 & v2))
    default:
        if (!false) {
            var v67 uint = 113
            v68 := GenS9{F0: -2147483648, F1: 128, F2: 0}
            v6 += (v6 / ((v6) | 1))
            _ = v67
            _ = v68
        } else {
            v2 += v2
            for fi69 := 0; fi69 < 7; fi69++ {
                var v70 int = fi69
                var v71 int64 = 0
                _ = fi69
                _ = v70
                _ = v71
            }
        }
    }
    _ = len(v4)
    for fi72 := 0; fi72 < 17; fi72++ {
        v73 := GenS4{F0: "\n\t", F1: 42}
        for fi74 := 0; fi74 < 1; fi74++ {
            fmt.Println(v62, v1, v3)
            var v75 uint16 = 2
            _ = fi74
            _ = v75
        }
        v4[fi72] = ("hello" + v62)
        _ = fi72
        _ = v73
    }
    var v76 uint16 = 32767
    v4[412] = v62
    v2 = 50
    v3 = 997
    v1 += v1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v62:", v62)
    fmt.Println("v76:", v76)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v62
    _ = v63
    _ = v76
}
