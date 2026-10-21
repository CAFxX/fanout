package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint16; F1 uint16 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int8; F1 uint8 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 uint; F1 string; F2 int64 }
type GenS7 struct { F0 int32; F1 uint8; F2 uint8 }
type GenS8 struct { F0 string; F1 uint16 }

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[uint32]bool)
    v3 := make(map[string]uint32)
    v4 := make(map[uint32]int)
    v5 := []uint32{2147483647, 472, 1, 2147483648}
    v6 := []int{128, -357, -159}
    var v7 int32 = 42
    var v8 float32 = 0.0
    for fi9 := 0; fi9 < 16; fi9++ {
        if ((0.0 - 1e10) != 1.5) {
            v10 := GenS1{F0: (917 > 0)}
            _ = v10
        }
        _ = fi9
    }
    if (v8 < v8) {
        if true {
            switch -1 {
            case 1:
                v7 += 2147483647
            }
            v8 -= 0.1
            if (!false) {
                switch "a" {
                case "":
                    fmt.Println(v8)
                case "\n\t":
                    v11 := GenS1{F0: false}
                    _ = v11
                case "foo bar":
                    var v12 float32 = 1e10
                    _ = v12
                }
            }
        } else {
            _ = len(v6)
            v8 = v8
        }
        v7 -= v7
        v7 = 42
    } else {
        if (1e10 == -2.25) {
            v7 += v7
        } else {
            v13 := GenS2{F0: 255, F1: 387}
            var v14 uint16 = 928
            _ = v13
            _ = v14
        }
    }
    if (1 <= 32768) {
        var v15 bool = (((("foo bar" + "") + "hello") != "foo bar") || (true || ("a" != "foo bar")))
        _ = v15
    } else {
        var v16 int32 = (v7 >> ((535) & 63))
        var v17 int32 = v16
        _ = v16
        _ = v17
    }
    for fi18 := 0; fi18 < 7; fi18++ {
        if (true && ((!(42 >= 846)) || (-1 == 127))) {
            var v19 uint32 = 4294967295
            for fi20 := 0; fi20 < 7; fi20++ {
                if true {
                    v21 := GenS2{F0: 32767, F1: 1}
                    fmt.Println(v7, v19, fi18)
                    _ = v21
                }
                var v22 int64 = 0
                v23 := GenS2{F0: 220, F1: 0}
                _ = fi20
                _ = v22
                _ = v23
            }
            switch (fi18 & int(v7)) {
            case -803:
                for fi24 := 0; fi24 < 17; fi24++ {
                    v25 := GenS3{F0: 1000}
                    v7 -= v7
                    switch "hello" {
                    case "":
                        fmt.Println(v7, fi18, fi24)
                        if ((42 >= 255) || false) {
                            v26 := GenS4{F0: -128, F1: 2}
                            var v27 float32 = (3.14159 - (1e10 - 1.5))
                            var v28 float64 = 3.14159
                            _ = v26
                            _ = v27
                            _ = v28
                        } else {
                            fmt.Println(fi24)
                            switch (fi24 & fi18) {
                            case -580:
                                if ((v19 / ((0) | 1)) <= (v19 ^ 2147483647)) {
                                    var v29 bool = true
                                    v7 = v7
                                    fmt.Println(v29, v8)
                                    _ = v29
                                }
                            case 32768:
                                var v30 uint = 2
                                _ = v30
                            default:
                                v8 = (v8 - float32(((float64(1e10) - -2.25) * (3.14159 - 1e10))))
                            }
                            fmt.Println(fi24)
                        }
                    case "foo bar":
                        fmt.Println(v7)
                    case "a":
                        switch fi24 {
                        case 32767:
                            var v31 uint64 = 42
                            _ = v31
                        case 42:
                            var v32 int64 = -9223372036854775808
                            _ = v32
                        }
                    }
                    _ = fi24
                    _ = v25
                }
            case 127:
                v33 := GenS4{F0: 76, F1: 128}
                _ = v33
            case 328:
                v34 := GenS2{F0: uint16(v19), F1: 32767}
                _ = v34
            default:
                var v35 uint32 = (v19 & v19)
                _ = v35
            }
            _ = v19
        } else {
            switch ("" + ("foo bar" + "a")) {
            case "":
                v36 := GenS5{F0: 2147483647}
                _ = v36
            default:
                var v37 uint16 = 42
                _ = v37
            }
        }
        v38 := v4[357]
        _ = fi18
        _ = v38
    }
    v39 := GenS6{F0: 4294967295, F1: "\n\t", F2: 42}
    switch 32768 {
    case 42:
        for fi40 := 0; fi40 < 6; fi40++ {
            v8 = ((v8 * v8) / v8)
            var v41 int = (fi40 + (fi40 + (fi40 % (((fi40 ^ 32768)) | 1))))
            var v42 uint = 367
            _ = fi40
            _ = v41
            _ = v42
        }
    case -1:
        var v43 uint64 = 4294967295
        _ = v43
    default:
        if (!false) {
            for fi44 := 0; fi44 < 15; fi44++ {
                switch "foo bar" {
                case "\n\t":
                    if (true && (!(-525 >= -677))) {
                        v7 = v7
                        if true {
                            v45 := GenS7{F0: (v7 - v7), F1: 128, F2: 1}
                            _ = v45
                        }
                    } else {
                        if ((("a" > "") || false) && (false && true)) {
                            var v46 int64 = -424
                            _ = v46
                        }
                    }
                case "foo bar":
                    v7 -= v7
                case "hello":
                    fmt.Println(v7)
                }
                _ = fi44
            }
            _ = len(v5)
        }
    }
    var v47 uint8 = 42
    var v48 uint = 2
    if false {
        v49 := GenS6{F0: v48, F1: "", F2: 523}
        fmt.Println(v47, v7, v7)
        _ = v49
    } else {
        for fi50 := 0; fi50 < 20; fi50++ {
            if ((-1 != 9223372036854775807) || ((2147483647 <= 1) || (-2147483648 >= (-(v7))))) {
                for fi51 := 0; fi51 < 1; fi51++ {
                    v47 += (^(v47))
                    switch (fi50 * fi51) {
                    case -9223372036854775808:
                        fmt.Println(fi51, v8)
                    case 32767:
                        if (!(true || false)) {
                            fmt.Println(v8, v47, v7)
                            v52 := GenS1{F0: false}
                            v48 = v48
                            if (0 >= 0) {
                                if (0 < 65535) {
                                    var v53 uint = v48
                                    v47 = (v47 % ((v47) | 1))
                                    _ = v53
                                }
                                if false {
                                    v7 += v7
                                    var v54 int8 = 100
                                    if false {
                                        v55 := GenS3{F0: -1}
                                        v48 += 65535
                                        v56 := GenS8{F0: "foo bar", F1: 0}
                                        fmt.Println(v8)
                                        v57 := GenS4{F0: (v54 - int8(v47)), F1: (v47 % ((v47) | 1))}
                                        _ = v55
                                        _ = v56
                                        _ = v57
                                    }
                                    fmt.Println(v47, v47, v48)
                                    fmt.Println(fi50, fi50)
                                    _ = v54
                                }
                                v47 += (^(v47))
                            }
                            _ = v52
                        }
                    case 42:
                        _ = len(v3)
                    }
                    _ = fi51
                }
            } else {
                v58 := GenS2{F0: 255, F1: 42}
                var v59 uint8 = (^(v47))
                _ = v58
                _ = v59
            }
            if false {
                v8 -= float32((((1e10 * -2.25) / (0.1 * (float64((((0.1 + float32(((float64(v8) - 0.1) + -2.25))) + ((v8 + ((float32(0.0) + v8) / v8)) - 3.14159)) * v8)) - ((1.5 / -2.25) / 1.5)))) + -2.25))
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                var v60 int = (-(fi50))
                _ = v60
            }
            fmt.Println(v7, v48)
            _ = fi50
        }
        for fi61 := 0; fi61 < 5; fi61++ {
            for fi62 := 0; fi62 < 10; fi62++ {
                v7 = (v7 | (^(v7)))
                _ = fi62
            }
            v8 = v8
            _ = fi61
        }
    }
    if (18446744073709551615 > 326) {
        var v63 uint8 = (v47 & (v47 * (v47 | (v47 << (((v48 % (((v48 ^ 18446744073709551615)) | 1))) & 63)))))
        v47 = (v47 | (^(v47)))
        _ = v63
    } else {
        v48 = v48
        var v64 int32 = (v7 * 886)
        fmt.Println(v48, v64)
        _ = v64
    }
    v48 = (v48 | v48)
    _ = len(v5)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v47
    _ = v48
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
