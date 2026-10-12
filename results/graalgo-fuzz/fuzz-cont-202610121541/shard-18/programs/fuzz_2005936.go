package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64 }
type GenS2 struct { F0 uint64; F1 int64; F2 int16 }
type GenS3 struct { F0 int16; F1 int8; F2 uint64 }
type GenS4 struct { F0 int }

func main() {
    v1 := make(map[int]uint32)
    v2 := []int{42, 1}
    v3 := []uint32{4294967295}
    v4 := make(map[string]uint32)
    v5 := make(map[uint32]string)
    v6 := GenS1{F0: -10, F1: 1}
    var v7 uint32 = 65535
    fmt.Println(v7)
    v7 += 0
    fmt.Println(v7)
    var v8 float32 = -2.25
    if (32767 == 1) {
        _ = len(v1)
        v9 := GenS2{F0: 4294967295, F1: 1, F2: 872}
        _ = v9
    }
    var v10 float32 = float32(1e10)
    for fi11 := 0; fi11 < 3; fi11++ {
        if true {
            var v12 string = ("foo bar" + "a")
            _ = v12
        } else {
            v13 := GenS2{F0: 318, F1: 0, F2: -1}
            if false {
                var v14 uint16 = 65535
                for fi15 := 0; fi15 < 18; fi15++ {
                    v8 = float32((0.1 * (-2.25 / 3.14159)))
                    _ = fi15
                }
                _ = v14
            } else {
                v8 = -2.25
                for fi16 := 0; fi16 < 12; fi16++ {
                    if (255 >= 1) {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        var v17 int8 = 0
                        if ((v17 * (-(v17))) < v17) {
                            if (!true) {
                                v10 = (3.14159 - float32((1e10 + 1.5)))
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                v18 := GenS2{F0: 18446744073709551615, F1: -9223372036854775808, F2: -1}
                                switch "foo bar" {
                                case "hello":
                                    fmt.Println(v7)
                                case "":
                                    v19 := v1[127]
                                    _ = v19
                                case "\n\t":
                                    v17 -= v17
                                }
                                v20 := GenS3{F0: 694, F1: 50, F2: 42}
                                _ = v18
                                _ = v20
                            }
                            v21 := GenS1{F0: 889, F1: 42}
                            _ = v21
                        }
                        _ = v17
                    }
                    v7 -= (v7 >> ((255) & 63))
                    if ((true && false) && true) {
                        if true {
                            delete(v1, fi16)
                        }
                        v2 = append(v2, 2147483647)
                        delete(v5, v7)
                        var v22 int64 = 9223372036854775807
                        var v23 int64 = (v22 << ((18446744073709551615) & 63))
                        _ = v22
                        _ = v23
                    }
                    _ = fi16
                }
            }
            for fi24 := 0; fi24 < 16; fi24++ {
                if (32767 != 65535) {
                    v7 -= v7
                    v10 = ((0.1 * v10) / 1.0)
                } else {
                    if (625 < -591) {
                        switch "" {
                        case "\n\t":
                            var v25 int32 = 790
                            _ = v25
                        case "foo bar":
                            v4[("hello" + "")] = (v7 ^ 748)
                            fmt.Println(fi11, v8, fi24)
                            v7 -= v7
                        }
                    }
                    v10 = (v10 * v10)
                    fmt.Println(v8)
                    if ("" < "") {
                        v2 = append(v2, (fi11 | -892))
                        fmt.Println(fi11, fi11, fi24)
                        fmt.Println(fi24)
                    } else {
                        fmt.Println(fi11)
                        if (616 > 390) {
                            v3 = append(v3, 2147483647)
                            v7 -= v7
                        } else {
                            v3 = append(v3, (v7 * v7))
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            var v26 int = 321
                            _ = v26
                        }
                    }
                    var v27 uint16 = uint16(v7)
                    var v28 int = fi24
                    _ = v27
                    _ = v28
                }
                _ = fi24
            }
            _ = v13
        }
        _ = fi11
    }
    switch ("" + "\n\t") {
    case "foo bar":
        var v29 float32 = v8
        _ = v29
    case "hello":
        delete(v4, "hello")
    case "":
        var v30 int32 = 860
        _ = v30
    default:
        v8 += (float32(1e10) + v10)
    }
    switch -1 {
    case 1:
        v31 := GenS4{F0: 2147483647}
        _ = v31
    case -9223372036854775808:
        v7 = 65535
    default:
        var v32 int32 = 321
        _ = v32
    }
    var v33 int32 = 2147483647
    switch ("" + "hello") {
    case "hello":
        var v34 int = -59
        _ = v34
    case "a":
        if false {
            var v35 float32 = v8
            if (-9223372036854775808 == 4294967295) {
                v36 := v1[-68]
                var v37 int8 = 1
                _ = v36
                _ = v37
            } else {
                switch "\n\t" {
                case "\n\t":
                    if ((!((!(false || false)) && (-880 == -1))) && ((v33 / ((v33) | 1)) <= v33)) {
                        _ = len(v4)
                        _ = len(v5)
                        var v38 int = -963
                        _ = v38
                    }
                case "":
                    switch "foo bar" {
                    case "hello":
                        var v39 int16 = -32768
                        _ = v39
                    case "":
                        var v40 int32 = v33
                        _ = v40
                    case "\n\t":
                        v8 = v8
                    }
                case "foo bar":
                    v41 := GenS3{F0: 0, F1: 100, F2: 357}
                    _ = v41
                }
                var v42 uint32 = (^(v7))
                _ = v42
            }
            v10 += 0.0
            _ = v35
        }
    case "\n\t":
        v43 := GenS1{F0: 42, F1: 18446744073709551615}
        _ = v43
    }
    var v44 string = ("hello" + ("" + "a"))
    for fi45 := 0; fi45 < 4; fi45++ {
        switch "foo bar" {
        case "foo bar":
            fmt.Println(v7)
        case "\n\t":
            v3 = append(v3, v7)
        case "":
            v10 -= (v10 - 1.5)
        }
        if (!true) {
            var v46 uint32 = 2147483648
            _ = v46
        } else {
            var v47 bool = ((!true) && false)
            _ = v47
        }
        _ = fi45
    }
    fmt.Println(v7)
    var v48 uint8 = 2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v44:", v44)
    fmt.Println("v48:", v48)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v44
    _ = v48
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
