package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64; F2 uint8 }
type GenS2 struct { F0 int64; F1 int32 }
type GenS3 struct { F0 uint8; F1 uint; F2 uint16 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int32; F1 int16; F2 int }
type GenS6 struct { F0 uint16; F1 int64; F2 int16 }

func main() {
    var v1 uint64 = 4294967295
    v2 := make(map[int]string)
    v3 := []string{(("\n\t" + "foo bar") + "a"), "a", "\n\t", ("" + "")}
    var v4 uint = 837
    var v5 float64 = 0.0
    if false {
        fmt.Println(v4, v1, v4)
        v6 := GenS1{F0: -2147483648, F1: 131, F2: 42}
        _ = v6
    } else {
        var v7 string = "a"
        var v8 uint32 = 631
        _ = v7
        _ = v8
    }
    if true {
        if len(v3) > 0 {
            v3[0] = "\n\t"
        }
    } else {
        var v9 uint32 = 2147483647
        _ = v9
    }
    var v10 float32 = 3.14159
    if len(v3) > 0 {
        v3[0] = ""
    }
    var v11 int = 547
    v11 += (v11 & v11)
    if (-107 < -1) {
        if false {
            v12 := GenS1{F0: 0, F1: 0, F2: uint8(v4)}
            var v13 int8 = -128
            var v14 string = "\n\t"
            _ = v12
            _ = v13
            _ = v14
        } else {
            var v15 uint8 = 0
            _ = v15
        }
        v2[128] = "hello"
    } else {
        var v16 int = (v11 + (v11 * (-(v11))))
        var v17 bool = true
        var v18 int = 9223372036854775807
        _ = v16
        _ = v17
        _ = v18
    }
    if true {
        v19 := GenS2{F0: -9223372036854775808, F1: -908}
        if (!(255 >= 65535)) {
            v4 += uint(v1)
        } else {
            var v20 string = ("a" + ("a" + (("a" + ("" + ("" + ("" + "")))) + "")))
            _ = v20
        }
        if true {
            if (!(true || true)) {
                v11 += -1
            }
        } else {
            if true {
                for fi21 := 0; fi21 < 4; fi21++ {
                    _ = len(v2)
                    for fi22 := 0; fi22 < 18; fi22++ {
                        switch (((("foo bar" + "") + "hello") + "hello") + (("" + "\n\t") + ("" + ("" + "hello")))) {
                        case "hello":
                            v23 := GenS2{F0: int64(v1), F1: 912}
                            _ = v23
                        case "foo bar":
                            fmt.Println(fi21, fi21)
                        }
                        _ = fi22
                    }
                    if true {
                        switch (v11 - 314) {
                        case 0:
                            v24 := GenS3{F0: 61, F1: v4, F2: 32767}
                            _ = v24
                        case 128:
                            v10 -= v10
                        case 2147483647:
                            for fi25 := 0; fi25 < 6; fi25++ {
                                v26 := GenS2{F0: int64(v4), F1: 1}
                                var v27 int32 = 401
                                _ = fi25
                                _ = v26
                                _ = v27
                            }
                        default:
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        }
                    }
                    _ = fi21
                }
            } else {
                switch "hello" {
                case "a":
                    v28 := GenS3{F0: 118, F1: v4, F2: 65535}
                    _ = v28
                }
                v10 -= v10
            }
            if (((!true) && true) || false) {
                if (2 < 0) {
                    for fi29 := 0; fi29 < 7; fi29++ {
                        v5 -= v5
                        for fi30 := 0; fi30 < 19; fi30++ {
                            fmt.Println(v1)
                            _ = fi30
                        }
                        _ = fi29
                    }
                } else {
                    fmt.Println(v1)
                    v5 += ((v5 / v5) - v5)
                }
                var v31 int16 = -449
                fmt.Println(v5)
                _ = v31
            } else {
                v3 = append(v3, ("" + "foo bar"))
                if (32767 <= 947) {
                    v10 = ((0.0 * v10) - v10)
                }
                v1 = v1
            }
        }
        _ = v19
    } else {
        if (42 > 32767) {
            if (!(1 < 857)) {
                v10 += (0.1 * (v10 - v10))
                if (!false) {
                    delete(v2, (v11 * v11))
                    v11 = v11
                } else {
                    v32 := GenS4{F0: 9223372036854775807}
                    _ = v32
                }
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            }
        } else {
            if (true && false) {
                switch "\n\t" {
                case "hello":
                    v10 += v10
                default:
                    if (true || true) {
                        v1 = 42
                    } else {
                        if (true && false) {
                            v10 += v10
                            var v33 uint64 = (v1 % (((v1 >> ((2) & 63))) | 1))
                            fmt.Println(v10, v11)
                            _ = v33
                        }
                    }
                }
            }
            fmt.Println(v11)
            switch ("" + "") {
            case "foo bar":
                v11 = v11
            case "\n\t":
                switch (v11 / ((-435) | 1)) {
                case 437:
                    switch 128 {
                    case -2147483648:
                        v3 = append(v3, "a")
                    case 317:
                        var v34 int64 = -1
                        _ = v34
                    }
                default:
                    v35 := GenS1{F0: 2147483647, F1: 1, F2: 203}
                    _ = v35
                }
            case "hello":
                v2[v11] = ("a" + "")
            default:
                if (((v1 + 42) <= (v1 % ((4294967295) | 1))) && (v4 < v4)) {
                    v1 = 18446744073709551615
                    if true {
                        switch -9223372036854775808 {
                        case 344:
                            var v36 int8 = 0
                            _ = v36
                        default:
                            v10 = 1e10
                        }
                    }
                    var v37 uint = (v4 / ((v4) | 1))
                    _ = v37
                } else {
                    switch (v11 + (v11 * (v11 >> (((v4 * (^(v4)))) & 63)))) {
                    case 9223372036854775807:
                        for fi38 := 0; fi38 < 15; fi38++ {
                            if true {
                                for fi39 := 0; fi39 < 18; fi39++ {
                                    var v40 float64 = v5
                                    if (42 >= -1) {
                                        _ = len(v3)
                                        if true {
                                            var v41 string = ("a" + (("a" + "") + ("\n\t" + "hello")))
                                            v42 := GenS2{F0: 9223372036854775807, F1: 42}
                                            fmt.Println(v5, v4, fi39)
                                            _ = v41
                                            _ = v42
                                        }
                                    } else {
                                        fmt.Println(v10, v40, v11)
                                        if ((false || (1 > -2147483648)) || true) {
                                            var v43 int = 9223372036854775807
                                            v44 := GenS5{F0: -1, F1: -61, F2: (-(fi39))}
                                            _ = v43
                                            _ = v44
                                        }
                                    }
                                    var v45 uint32 = 2
                                    _ = fi39
                                    _ = v40
                                    _ = v45
                                }
                                if (float64(float32(v5)) >= (float64(float32(3.14159)) - v5)) {
                                    for fi46 := 0; fi46 < 19; fi46++ {
                                        fmt.Println(v4, v1)
                                        v1 = 0
                                        _ = fi46
                                    }
                                    if false {
                                        v47 := GenS6{F0: 1, F1: -907, F2: -799}
                                        v5 = (v5 + 1e10)
                                        fmt.Println(v5)
                                        _ = v47
                                    }
                                }
                            } else {
                                v2[(fi38 | (fi38 | fi38))] = ""
                                if false {
                                    v10 += 0.0
                                    var v48 float64 = v5
                                    for fi49 := 0; fi49 < 11; fi49++ {
                                        var v50 uint8 = 255
                                        _ = fi49
                                        _ = v50
                                    }
                                    _ = v48
                                }
                                if len(v3) > 0 {
                                    v3[0] = (((("" + "a") + "") + (((("hello" + "foo bar") + "a") + "foo bar") + "foo bar")) + "a")
                                }
                            }
                            _ = fi38
                        }
                    case -1:
                        for fi51 := 0; fi51 < 8; fi51++ {
                            switch "a" {
                            case "a":
                                v11 -= (v11 | (v11 ^ (v11 + (v11 & v11))))
                            }
                            var v52 uint8 = 0
                            for fi53 := 0; fi53 < 9; fi53++ {
                                v1 = uint64(v52)
                                _ = fi53
                            }
                            _ = fi51
                            _ = v52
                        }
                    case 0:
                        switch "\n\t" {
                        case "foo bar":
                            v5 += -2.25
                        case "a":
                            if (452 == 65535) {
                                v54 := GenS6{F0: 42, F1: 198, F2: 1}
                                switch 32767 {
                                case 128:
                                    var v55 int64 = -673
                                    _ = v55
                                }
                                _ = v54
                            } else {
                                var v56 int16 = 42
                                v56 = v56
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                _ = v56
                            }
                        default:
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        }
                    }
                }
            }
        }
        for fi57 := 0; fi57 < 20; fi57++ {
            var v58 string = "a"
            v59 := GenS4{F0: 9223372036854775807}
            _ = fi57
            _ = v58
            _ = v59
        }
        switch ("foo bar" + "") {
        case "":
            for fi60 := 0; fi60 < 19; fi60++ {
                v61 := GenS2{F0: -1, F1: 42}
                v5 -= (float64((v10 + (v10 / float32(v5)))) / v5)
                _ = fi60
                _ = v61
            }
        case "\n\t":
            fmt.Println(v11, v1, v11)
        }
    }
    fmt.Println(v5, v11)
    var v62 uint32 = 210
    var v63 bool = (true || true)
    var v64 float32 = float32(v5)
    var v65 uint32 = (v62 + v62)
    for fi66 := 0; fi66 < 5; fi66++ {
        var v67 bool = (!(true && (false && (9223372036854775807 != 1))))
        _ = fi66
        _ = v67
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v62:", v62)
    fmt.Println("v63:", v63)
    fmt.Println("v64:", v64)
    fmt.Println("v65:", v65)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v62
    _ = v63
    _ = v64
    _ = v65
}
