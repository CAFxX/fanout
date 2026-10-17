package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16; F2 int32 }
type GenS2 struct { F0 uint16; F1 uint8 }
type GenS3 struct { F0 int; F1 int16 }
type GenS4 struct { F0 string; F1 uint }
type GenS5 struct { F0 uint; F1 uint16; F2 int32 }
type GenS6 struct { F0 int; F1 int8; F2 uint8 }

func main() {
    v1 := make(map[uint32]string)
    v2 := []int{-9223372036854775808, 351}
    var v3 uint = 223
    v4 := []int{32767, -2147483648, 1}
    var v5 int16 = 622
    var v6 int32 = int32(v5)
    var v7 uint = (v3 + v3)
    for fi8 := 0; fi8 < 5; fi8++ {
        var v9 uint64 = 4294967295
        var v10 int = (fi8 / (((fi8 | fi8)) | 1))
        _ = fi8
        _ = v10
        _ = v9
    }
    v7 = (v3 ^ v7)
    v11 := GenS1{F0: (v5 / (((v5 << (((v3 | 1)) & 63))) | 1)), F1: 65535, F2: (v6 & v6)}
    v2 = append(v2, 128)
    var v12 int64 = -9223372036854775808
    v13 := GenS1{F0: v5, F1: 0, F2: (^(v6))}
    var v14 int16 = v5
    if false {
        if ((v7 * (v3 / ((v3) | 1))) >= (^(v3))) {
            _ = len(v1)
            var v15 int8 = 1
            _ = v15
        } else {
            v6 += v6
            var v16 float64 = (3.14159 * 1.5)
            v12 = (v12 | v12)
            _ = v16
        }
    } else {
        v6 = v6
        if (100 == 100) {
            if (741 <= 18446744073709551615) {
                v12 -= 1
            }
            if (!(307 > 125)) {
                v12 += -9223372036854775808
                switch "\n\t" {
                case "hello":
                    var v17 float64 = (1e10 / 1.0)
                    _ = v17
                case "foo bar":
                    if (-1 != 42) {
                        switch "\n\t" {
                        case "hello":
                            for fi18 := 0; fi18 < 9; fi18++ {
                                v6 = 2147483647
                                v14 += v5
                                _ = fi18
                            }
                        case "foo bar":
                            if (!((true || false) && (549 != 769))) {
                                v14 = -235
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                            }
                        case "a":
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        default:
                            v19 := GenS2{F0: 554, F1: 128}
                            _ = v19
                        }
                        v20 := GenS3{F0: 4294967295, F1: v14}
                        _ = v20
                    } else {
                        for fi21 := 0; fi21 < 1; fi21++ {
                            v5 += v14
                            v6 += (v6 >> (((v7 - v7)) & 63))
                            v6 -= (v6 >> (((v7 | (v7 | v7))) & 63))
                            _ = fi21
                        }
                    }
                case "\n\t":
                    var v22 uint32 = 2147483647
                    _ = v22
                }
                for fi23 := 0; fi23 < 16; fi23++ {
                    var v24 uint16 = 65535
                    _ = fi23
                    _ = v24
                }
            } else {
                v25 := GenS4{F0: "hello", F1: v7}
                v3 = v3
                for fi26 := 0; fi26 < 2; fi26++ {
                    v3 += 0
                    for fi27 := 0; fi27 < 4; fi27++ {
                        v28 := GenS5{F0: (v7 % (((^(v3))) | 1)), F1: 42, F2: (v6 - (v6 ^ v6))}
                        _ = fi27
                        _ = v28
                    }
                    v6 = v6
                    _ = fi26
                }
                _ = v25
            }
        } else {
            v3 -= (v3 | 255)
            switch (("" + "\n\t") + ("\n\t" + "a")) {
            case "hello":
                var v29 uint32 = 233
                _ = v29
            case "foo bar":
                var v30 uint32 = 42
                _ = v30
            }
        }
    }
    for fi31 := 0; fi31 < 14; fi31++ {
        var v32 int8 = 100
        _ = fi31
        _ = v32
    }
    fmt.Println(v12, v6)
    if false {
        fmt.Println(v5, v6, v12)
    }
    fmt.Println(v3)
    if (0.0 != 0.1) {
        if false {
            v5 += (v14 / ((664) | 1))
            var v33 int = -2147483648
            _ = v33
        } else {
            var v34 int = 0
            _ = v34
        }
        v14 -= v14
    }
    switch (("foo bar" + "foo bar") + "") {
    case "":
        v1[254] = "a"
    default:
        var v35 float64 = 3.14159
        _ = v35
    }
    for fi36 := 0; fi36 < 20; fi36++ {
        v37 := GenS6{F0: (-(fi36)), F1: 42, F2: 1}
        for fi38 := 0; fi38 < 20; fi38++ {
            switch (("foo bar" + ("" + "foo bar")) + ("" + "a")) {
            case "hello":
                v1[2] = ""
                v3 = 255
            case "\n\t":
                var v39 uint = 65535
                _ = v39
            case "a":
                v7 -= (v7 % ((v7) | 1))
            }
            var v40 int8 = 42
            _ = fi38
            _ = v40
        }
        _ = fi36
        _ = v37
    }
    v3 += v7
    switch "a" {
    case "":
        var v41 int16 = (v5 - v14)
        _ = v41
    case "hello":
        if true {
            v42 := GenS3{F0: 127, F1: v14}
            v7 -= (v3 * v3)
            v43 := GenS5{F0: (v3 | v7), F1: 65535, F2: 380}
            _ = v42
            _ = v43
        } else {
            for fi44 := 0; fi44 < 1; fi44++ {
                switch (fi44 % ((4294967295) | 1)) {
                case 128:
                    fmt.Println(v5)
                case 32767:
                    v6 = (v6 - (v6 << (((v3 * v3)) & 63)))
                case 42:
                    v6 = v6
                default:
                    if (105 >= 42) {
                        if len(v4) > 0 {
                            v4[0] = 4294967295
                        }
                        for fi45 := 0; fi45 < 4; fi45++ {
                            var v46 int32 = v6
                            if true {
                                fmt.Println(v12, v7, v46)
                                v5 -= (v5 << (((v7 + v7)) & 63))
                            }
                            v5 = (v14 - (v5 ^ 42))
                            var v47 uint32 = 885
                            _ = fi45
                            _ = v46
                            _ = v47
                        }
                    }
                }
                switch "a" {
                case "":
                    fmt.Println(v6, v3)
                }
                _ = fi44
            }
            if false {
                v12 = v12
            }
            var v48 int16 = v14
            _ = v48
        }
    case "foo bar":
        var v49 bool = true
        _ = v49
    default:
        var v50 uint32 = 275
        _ = v50
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
