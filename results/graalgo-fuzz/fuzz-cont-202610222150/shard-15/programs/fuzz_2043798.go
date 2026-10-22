package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint; F2 uint }
type GenS2 struct { F0 string; F1 int8; F2 uint64 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint; F1 uint64; F2 string }
type GenS5 struct { F0 int32; F1 bool; F2 uint16 }
type GenS6 struct { F0 int32; F1 int; F2 uint8 }
type GenS7 struct { F0 int16; F1 string }
func genFn0() uint {
    for fi4 := 0; fi4 < 10; fi4++ {
        var v5 bool = (!false)
        var v6 string = "\n\t"
        var v7 string = ("hello" + v6)
        _ = fi4
        _ = v5
        _ = v6
        _ = v7
    }
    if true {
        var v8 int64 = 42
        if true {
            for fi9 := 0; fi9 < 8; fi9++ {
                v10 := GenS1{F0: 42, F1: 313, F2: 65535}
                _ = fi9
                _ = v10
            }
            var v11 int = 478
            _ = v11
        }
        switch "a" {
        case "\n\t":
            if true {
                v12 := GenS1{F0: 100, F1: 65535, F2: 42}
                v8 -= 0
                for fi13 := 0; fi13 < 19; fi13++ {
                    v14 := []int{(fi13 + fi13), (fi13 | fi13)}
                    v8 = (v8 / (((-(v8))) | 1))
                    _ = fi13
                    _ = v14
                }
                _ = v12
            }
        case "hello":
            v8 -= v8
        default:
            switch -717 {
            case 32768:
                v8 -= (^(v8))
            case 76:
                v15 := []int{9223372036854775807, 1, 0}
                _ = v15
            case 9223372036854775807:
                v8 -= (^(v8))
            default:
                v16 := GenS1{F0: 100, F1: 815, F2: 1}
                _ = v16
            }
        }
        _ = v8
    } else {
        for fi17 := 0; fi17 < 11; fi17++ {
            if true {
                switch fi17 {
                case -9223372036854775808:
                    var v18 string = "hello"
                    _ = v18
                case -822:
                    if false {
                        v19 := GenS1{F0: -128, F1: 4294967295, F2: 777}
                        _ = v19
                    }
                default:
                    v20 := []int{796}
                    _ = v20
                }
                fmt.Println(fi17)
                v21 := []int{(-(fi17))}
                _ = v21
            }
            _ = fi17
        }
        if false {
            v22 := []int{32768, -9223372036854775808, 990}
            v23 := GenS2{F0: "hello", F1: 15, F2: 42}
            var v24 uint = 2
            _ = v22
            _ = v23
            _ = v24
        }
    }
    return 659
}

func main() {
    v1 := make(map[uint32]int)
    var v2 float32 = -2.25
    v3 := make(map[string]int)
    switch 4294967295 {
    case -2147483648:
        for fi25 := 0; fi25 < 11; fi25++ {
            for fi26 := 0; fi26 < 1; fi26++ {
                var v27 uint64 = 2
                if (!false) {
                    if (!((fi25 ^ 753) < 1)) {
                        v28 := GenS3{F0: 2147483647}
                        v27 = 18446744073709551615
                        v29 := []int{-123}
                        _ = v28
                        _ = v29
                    }
                    switch "foo bar" {
                    case "\n\t":
                        if true {
                            switch fi25 {
                            case -1:
                                v27 += 361
                            case 897:
                                if true {
                                    v27 = v27
                                    var v30 uint16 = 638
                                    _ = v30
                                } else {
                                    v1[42] = fi26
                                }
                            default:
                                var v31 int32 = -2147483648
                                _ = v31
                            }
                        }
                    case "a":
                        v27 = (v27 << ((255) & 63))
                    case "":
                        var v32 bool = (592 > 0)
                        _ = v32
                    }
                }
                v2 += v2
                _ = fi26
                _ = v27
            }
            _ = fi25
        }
    case 300:
        switch "" {
        case "foo bar":
            v33 := GenS1{F0: -1, F1: genFn0(), F2: genFn0()}
            _ = v33
        case "":
            fmt.Println(v2)
        }
    case 128:
        v34 := GenS1{F0: -1, F1: 0, F2: 90}
        _ = v34
    default:
        v35 := GenS4{F0: 65535, F1: 1, F2: (("hello" + "a") + (("hello" + "a") + "hello"))}
        _ = v35
    }
    switch "foo bar" {
    case "\n\t":
        if (-9223372036854775808 > -1) {
            var v36 uint8 = 127
            _ = v36
        }
    case "":
        v37 := GenS5{F0: -1, F1: (18446744073709551615 >= 18446744073709551615), F2: 272}
        _ = v37
    }
    if (662 > 4294967295) {
        var v38 float32 = (v2 / 0.1)
        _ = v38
    }
    v2 += 1e10
    v2 = 1e10
    v39 := GenS5{F0: -2147483648, F1: (!true), F2: 0}
    v40 := GenS6{F0: 0, F1: -401, F2: 255}
    switch "" {
    case "a":
        for fi41 := 0; fi41 < 12; fi41++ {
            for fi42 := 0; fi42 < 10; fi42++ {
                fmt.Println(fi42)
                _ = fi42
            }
            v2 = float32((((((0.1 * (-2.25 * (0.0 - 0.0))) - 1e10) + 1e10) * (float64(v2) / float64(v2))) / 0.1))
            _ = fi41
        }
    default:
        v43 := []int{}
        _ = v43
    }
    var v44 int = 2147483647
    v44 = (^(v44))
    if false {
        v44 -= (v44 - 128)
        v45 := v1[42]
        var v46 int16 = 42
        _ = v45
        _ = v46
    } else {
        switch (^(v44)) {
        case -206:
            var v47 int = v44
            _ = v47
        case 0:
            var v48 string = ""
            _ = v48
        }
        switch "foo bar" {
        case "hello":
            v49 := GenS7{F0: 42, F1: ("a" + ("a" + ("" + "")))}
            _ = v49
        case "foo bar":
            switch v44 {
            case 588:
                v44 = v44
            case 32767:
                switch (v44 / ((v44) | 1)) {
                case 63:
                    for fi50 := 0; fi50 < 20; fi50++ {
                        fmt.Println(v2, v44, v2)
                        fmt.Println(v44)
                        v44 = v44
                        _ = fi50
                    }
                case -2147483648:
                    fmt.Println(v2, v44)
                case 4294967295:
                    var v51 int16 = 42
                    _ = v51
                default:
                    var v52 int16 = 924
                    _ = v52
                }
            case 127:
                if true {
                    switch (^(v44)) {
                    case -90:
                        for fi53 := 0; fi53 < 1; fi53++ {
                            for fi54 := 0; fi54 < 14; fi54++ {
                                var v55 int8 = 42
                                _ = fi54
                                _ = v55
                            }
                            v2 = v2
                            for fi56 := 0; fi56 < 11; fi56++ {
                                v44 -= 127
                                var v57 int8 = 42
                                v58 := []int{}
                                _ = fi56
                                _ = v57
                                _ = v58
                            }
                            _ = fi53
                        }
                    default:
                        switch ("hello" + "foo bar") {
                        case "a":
                            v59 := []int{v44}
                            _ = v59
                        case "\n\t":
                            if false {
                                switch 715 {
                                case -9223372036854775808:
                                    v60 := GenS4{F0: 170, F1: 906, F2: ("a" + "foo bar")}
                                    _ = v60
                                case -1:
                                    v44 -= v44
                                default:
                                    if true {
                                        v3[((("" + "\n\t") + "\n\t") + "foo bar")] = 128
                                    }
                                }
                                v61 := []int{}
                                if len(v61) > 0 {
                                    v61[0] = (v44 / ((v44) | 1))
                                }
                                _ = v61
                            }
                        case "hello":
                            v62 := []int{}
                            _ = v62
                        default:
                            v63 := v1[65535]
                            _ = v63
                        }
                    }
                    v64 := GenS1{F0: 127, F1: 2, F2: 42}
                    for fi65 := 0; fi65 < 2; fi65++ {
                        fmt.Println(v44)
                        _ = fi65
                    }
                    _ = v64
                }
            default:
                if false {
                    for fi66 := 0; fi66 < 3; fi66++ {
                        for fi67 := 0; fi67 < 20; fi67++ {
                            fmt.Println(fi67)
                            v44 = (^(fi66))
                            _ = fi67
                        }
                        _ = fi66
                    }
                    var v68 string = ("hello" + (("\n\t" + "") + "hello"))
                    fmt.Println(v44, v44, v44)
                    _ = v68
                }
            }
        }
        if (2 <= 2) {
            v2 = (v2 + (v2 - (float32((3.14159 + -2.25)) / 1e10)))
            var v69 int = v44
            if (!true) {
                v44 = 127
            } else {
                v69 = (v69 >> ((817) & 63))
                fmt.Println(v69, v44, v69)
            }
            _ = v69
        } else {
            var v70 float64 = 1e10
            _ = v70
        }
    }
    v71 := GenS7{F0: 1000, F1: ("a" + "")}
    fmt.Println(v2, v44)
    fmt.Println(v2, v2)
    switch "\n\t" {
    case "hello":
        var v72 uint64 = 0
        _ = v72
    default:
        v44 = v44
    }
    if (-1 >= -154) {
        v2 = v2
    } else {
        if true {
            if (76 != 137) {
                fmt.Println(v44, v2)
            } else {
                fmt.Println(v44, v2)
            }
            v3["\n\t"] = 42
            _ = len(v3)
        }
    }
    v44 = v44
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v44:", v44)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v40
    _ = v44
    _ = v71
}
