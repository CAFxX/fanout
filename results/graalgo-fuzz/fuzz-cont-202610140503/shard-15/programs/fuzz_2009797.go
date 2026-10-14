package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8 }
type GenS2 struct { F0 uint8; F1 string }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int8; F1 int64 }

func main() {
    var v1 float32 = 1.5
    v2 := make(map[int]string)
    v3 := make(map[string]int)
    var v4 int16 = -32768
    var v5 int32 = 137
    switch 380 {
    case 138:
        if ((32768 < -1) && true) {
            v6 := GenS1{F0: 127, F1: 127}
            var v7 uint32 = 2147483648
            if true {
                var v8 int16 = v4
                if (!(v7 != 844)) {
                    var v9 uint = 18446744073709551615
                    v4 = v4
                    if (v8 <= (v4 >> ((0) & 63))) {
                        v10 := GenS1{F0: 165, F1: 182}
                        _ = v10
                    }
                    _ = v9
                } else {
                    if true {
                        _ = len(v3)
                        v11 := []int{128, 0, -2147483648, 42}
                        var v12 uint32 = v7
                        _ = v11
                        _ = v12
                    } else {
                        for fi13 := 0; fi13 < 6; fi13++ {
                            switch ("a" + "") {
                            case "a":
                                v1 = 1e10
                            case "":
                                v14 := GenS2{F0: 0, F1: ("" + ("hello" + "a"))}
                                _ = v14
                            default:
                                if true {
                                    v15 := GenS1{F0: 255, F1: 252}
                                    var v16 uint = 74
                                    v17 := v2[(^(fi13))]
                                    _ = v15
                                    _ = v16
                                    _ = v17
                                }
                            }
                            var v18 uint8 = uint8(v7)
                            delete(v3, "\n\t")
                            _ = fi13
                            _ = v18
                        }
                    }
                    v1 -= (v1 * 1e10)
                }
                _ = v8
            }
            _ = v6
            _ = v7
        } else {
            v1 = (v1 * v1)
        }
    case 669:
        fmt.Println(v4, v1, v1)
    case 32767:
        v1 -= 3.14159
    }
    var v19 uint32 = 42
    _ = len(v2)
    var v20 uint16 = 65535
    fmt.Println(v20)
    if false {
        v5 = (v5 / (((v5 % ((int32(v19)) | 1))) | 1))
        switch ("hello" + "\n\t") {
        case "a":
            v5 += (v5 | (^(v5)))
        case "hello":
            for fi21 := 0; fi21 < 12; fi21++ {
                var v22 int32 = v5
                v20 += v20
                var v23 uint32 = (v19 >> ((65535) & 63))
                _ = fi21
                _ = v22
                _ = v23
            }
        case "foo bar":
            switch ((("hello" + ("a" + "")) + "hello") + "foo bar") {
            case "a":
                if (!false) {
                    v5 += -252
                    var v24 uint = 18446744073709551615
                    switch "" {
                    case "a":
                        for fi25 := 0; fi25 < 13; fi25++ {
                            if false {
                                if true {
                                    v26 := []int{(fi25 / (((fi25 + (fi25 * fi25))) | 1))}
                                    for fi27 := 0; fi27 < 3; fi27++ {
                                        if len(v26) >= 2 {
                                            v26 = v26[1:len(v26)]
                                        }
                                        if (100 != -1) {
                                            if true {
                                                switch "hello" {
                                                case "":
                                                    v28 := GenS2{F0: 128, F1: ("a" + "")}
                                                    _ = v28
                                                case "foo bar":
                                                    if len(v26) > 0 {
                                                        v26[0] = (fi27 & (fi27 / ((-24) | 1)))
                                                    }
                                                case "hello":
                                                    switch fi25 {
                                                    case 32768:
                                                        _ = len(v3)
                                                    case -766:
                                                        v1 = (v1 / v1)
                                                    case 127:
                                                        v29 := GenS1{F0: 0, F1: 44}
                                                        _ = v29
                                                    default:
                                                        v30 := GenS2{F0: 99, F1: ""}
                                                        _ = v30
                                                    }
                                                }
                                            } else {
                                                v1 = (-2.25 / (v1 - 0.0))
                                                var v31 int8 = -11
                                                if len(v26) > 0 {
                                                    v26[0] = int(v4)
                                                }
                                                _ = v31
                                            }
                                        }
                                        _ = fi27
                                    }
                                    v20 -= uint16(v19)
                                    _ = v26
                                } else {
                                    v32 := []int{int(v4), 32767, (fi25 >> (((v24 * (v24 << (((v24 / (((v24 & v24)) | 1))) & 63)))) & 63))}
                                    _ = v32
                                }
                                for fi33 := 0; fi33 < 1; fi33++ {
                                    v20 += (v20 - (^(v20)))
                                    v34 := GenS1{F0: 92, F1: 251}
                                    _ = fi33
                                    _ = v34
                                }
                            }
                            v1 = (float32((0.0 + (-2.25 * 1e10))) * v1)
                            _ = fi25
                        }
                    case "foo bar":
                        v35 := []int{-9223372036854775808, 0}
                        _ = v35
                    default:
                        if true {
                            for fi36 := 0; fi36 < 2; fi36++ {
                                v24 += (^(v24))
                                fmt.Println(v19, v20)
                                _ = fi36
                            }
                            v37 := GenS3{F0: v24}
                            var v38 float64 = 1e10
                            _ = v37
                            _ = v38
                        }
                    }
                    _ = v24
                } else {
                    delete(v3, "a")
                }
            case "foo bar":
                v39 := []int{-554}
                _ = v39
            default:
                for fi40 := 0; fi40 < 15; fi40++ {
                    for fi41 := 0; fi41 < 8; fi41++ {
                        v42 := GenS2{F0: 250, F1: "\n\t"}
                        _ = fi41
                        _ = v42
                    }
                    var v43 uint = 65535
                    for fi44 := 0; fi44 < 18; fi44++ {
                        switch "" {
                        case "foo bar":
                            v20 -= v20
                        case "":
                            var v45 bool = false
                            _ = v45
                        default:
                            if ("foo bar" > ("a" + "a")) {
                                v46 := []int{(fi40 % (((fi44 >> (((v43 + v43)) & 63))) | 1)), (-(fi40))}
                                var v47 uint16 = 65535
                                _ = v46
                                _ = v47
                            } else {
                                if (v43 > v43) {
                                    if ((!("hello" != "\n\t")) || (!(18446744073709551615 == 4294967295))) {
                                        var v48 uint = 730
                                        v49 := []int{127, fi44, fi40, (-(fi44))}
                                        _ = v48
                                        _ = v49
                                    } else {
                                        switch "a" {
                                        case "hello":
                                            var v50 uint64 = 2
                                            _ = v50
                                        }
                                        v51 := GenS3{F0: v43}
                                        v52 := GenS1{F0: 1, F1: 2}
                                        _ = v51
                                        _ = v52
                                    }
                                } else {
                                    if (-128 == -128) {
                                        var v53 uint = v43
                                        _ = v53
                                    } else {
                                        fmt.Println(v43, v43)
                                        var v54 int8 = -101
                                        _ = v54
                                    }
                                    v55 := []int{32768}
                                    _ = v55
                                }
                                var v56 bool = (!true)
                                switch "foo bar" {
                                case "hello":
                                    if (v56 || (v19 > (v19 - v19))) {
                                        v57 := []int{fi44}
                                        if v56 {
                                            var v58 uint = (v43 - v43)
                                            fmt.Println(v1, v5)
                                            fmt.Println(v1, v58)
                                            _ = v58
                                        }
                                        _ = v57
                                    }
                                default:
                                    switch 0 {
                                    case -9223372036854775808:
                                        v5 -= -1
                                    default:
                                        if ((fi40 << (((v43 / ((4294967295) | 1))) & 63)) == fi40) {
                                            switch "foo bar" {
                                            case "\n\t":
                                                var v59 uint32 = v19
                                                _ = v59
                                            default:
                                                var v60 bool = v56
                                                _ = v60
                                            }
                                        }
                                        v20 = v20
                                        fmt.Println(v19, v4, v4)
                                    }
                                }
                                _ = v56
                            }
                        }
                        v1 = (v1 / v1)
                        _ = fi44
                    }
                    _ = fi40
                    _ = v43
                }
            }
        }
        v19 = 723
    } else {
        for fi61 := 0; fi61 < 7; fi61++ {
            v20 = v20
            _ = fi61
        }
        fmt.Println(v5, v19, v20)
    }
    v62 := []int{224, 898, -1}
    for fi63 := 0; fi63 < 18; fi63++ {
        v64 := GenS4{F0: 18446744073709551615}
        var v65 int8 = 1
        _ = fi63
        _ = v64
        _ = v65
    }
    for fi66 := 0; fi66 < 14; fi66++ {
        if (v5 >= (v5 >> ((1) & 63))) {
            if ((208 != 0) && false) {
                if (fi66 >= (fi66 ^ -265)) {
                    if false {
                        for fi67 := 0; fi67 < 1; fi67++ {
                            delete(v2, (^(fi66)))
                            switch ("a" + "a") {
                            case "foo bar":
                                switch int(v20) {
                                case 2147483647:
                                    v19 = v19
                                default:
                                    fmt.Println(v4)
                                }
                            case "hello":
                                v20 = 828
                            }
                            _ = fi67
                        }
                    } else {
                        var v68 bool = (!false)
                        if (v19 > v19) {
                            for fi69 := 0; fi69 < 3; fi69++ {
                                if true {
                                    _ = len(v62)
                                } else {
                                    if v68 {
                                        var v70 uint = 65535
                                        v19 = 15
                                        _ = v70
                                    }
                                    _ = len(v3)
                                }
                                v19 = v19
                                v71 := GenS5{F0: -128, F1: 1}
                                _ = fi69
                                _ = v71
                            }
                        }
                        _ = v68
                    }
                    v19 = v19
                }
            } else {
                var v72 bool = true
                v4 += v4
                var v73 uint = 0
                _ = v72
                _ = v73
            }
        }
        v19 = 2147483647
        _ = fi66
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v62:", len(v62), v62)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v62
}
