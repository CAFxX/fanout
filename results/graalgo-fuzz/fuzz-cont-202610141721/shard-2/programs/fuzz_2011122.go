package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64 }
type GenS2 struct { F0 int8; F1 int8 }
type GenS3 struct { F0 int64; F1 uint }
type GenS4 struct { F0 int16; F1 uint32; F2 int16 }
type GenS5 struct { F0 int16; F1 uint8 }
type GenS6 struct { F0 uint8; F1 uint64; F2 uint }
type GenS7 struct { F0 uint16; F1 uint64; F2 int }
type GenS8 struct { F0 int32 }
type GenS9 struct { F0 int; F1 string }
type GenS10 struct { F0 uint8 }
type GenS11 struct { F0 int16; F1 uint8 }
type GenS12 struct { F0 int8; F1 int; F2 uint16 }

func main() {
    v1 := []uint32{42, 284, 995, 917}
    v2 := make(map[uint32]uint32)
    var v3 int32 = 42
    v4 := []string{"a"}
    v5 := make(map[uint32]string)
    var v6 uint = 65535
    var v7 int32 = (v3 * (v3 % (((v3 >> ((2) & 63))) | 1)))
    if (-9223372036854775808 >= 128) {
        var v8 uint16 = 42
        v9 := GenS1{F0: 1, F1: 2}
        v10 := GenS2{F0: -1, F1: 100}
        _ = v10
        _ = v8
        _ = v9
    }
    for fi11 := 0; fi11 < 13; fi11++ {
        if (float64(0.1) != 0.0) {
            if true {
                for fi12 := 0; fi12 < 14; fi12++ {
                    v13 := GenS1{F0: -25, F1: 1}
                    _ = fi12
                    _ = v13
                }
                v14 := GenS3{F0: 543, F1: 65535}
                for fi15 := 0; fi15 < 18; fi15++ {
                    switch "a" {
                    case "hello":
                        v3 -= (v7 >> ((4294967295) & 63))
                    default:
                        if false {
                            v6 = (v6 - v6)
                            v16 := GenS3{F0: -1, F1: (v6 / (((v6 - (v6 << (((v6 * 18446744073709551615)) & 63)))) | 1))}
                            var v17 int = (fi15 & (fi15 - (fi15 << (((v6 << ((v6) & 63))) & 63))))
                            _ = v16
                            _ = v17
                        } else {
                            if false {
                                v6 = v6
                                _ = len(v2)
                            } else {
                                fmt.Println(v7, v6, v7)
                                if (!(((v3 / ((0) | 1)) != v7) || true)) {
                                    var v18 string = "a"
                                    v18 = ("hello" + "foo bar")
                                    v19 := GenS3{F0: -653, F1: v6}
                                    _ = v18
                                    _ = v19
                                } else {
                                    var v20 int16 = -32768
                                    _ = v20
                                }
                                if (!(false || true)) {
                                    v21 := GenS1{F0: 127, F1: 155}
                                    _ = v21
                                } else {
                                    v22 := GenS3{F0: 0, F1: 42}
                                    v23 := GenS2{F0: 100, F1: 100}
                                    if true {
                                        if ((float64(((0.0 - 0.1) / (((1.5 / -2.25) / (float32(float64((3.14159 + 1e10))) / ((0.1 * ((((-2.25 + 1.5) / (float32(-2.25) * ((3.14159 / 1.0) * -2.25))) - -2.25) - 1e10)) * ((0.0 - float32(-2.25)) + 1.5)))) - (0.1 / 3.14159)))) + (1e10 - 1e10)) != 0.1) {
                                            var v24 uint8 = 2
                                            var v25 int = 32767
                                            _ = v24
                                            _ = v25
                                        } else {
                                            var v26 int8 = -1
                                            switch "hello" {
                                            case "":
                                                var v27 int64 = -1
                                                _ = v27
                                            }
                                            _ = v26
                                        }
                                        v28 := GenS2{F0: 11, F1: -43}
                                        v6 += (v6 - (v6 >> ((2) & 63)))
                                        fmt.Println(v7, v6, v3)
                                        _ = v28
                                    }
                                    _ = v22
                                    _ = v23
                                }
                                fmt.Println(v6, v6)
                                v3 = v3
                            }
                        }
                    }
                    switch ("a" + "a") {
                    case "a":
                        v29 := GenS4{F0: 833, F1: uint32(v6), F2: 207}
                        _ = v29
                    case "\n\t":
                        var v30 int = -782
                        _ = v30
                    default:
                        if ((^(v6)) >= 18446744073709551615) {
                            var v31 uint8 = 1
                            fmt.Println(fi15)
                            if true {
                                var v32 uint32 = 2147483647
                                switch (-(fi15)) {
                                case 127:
                                    var v33 int64 = 0
                                    v34 := GenS4{F0: -476, F1: v32, F2: 629}
                                    fmt.Println(fi11)
                                    _ = v33
                                    _ = v34
                                case 0:
                                    delete(v5, 710)
                                case 32768:
                                    var v35 int = fi15
                                    _ = v35
                                }
                                _ = v32
                            }
                            _ = v31
                        }
                    }
                    _ = fi15
                }
                _ = v14
            }
            fmt.Println(v6)
        }
        _ = fi11
    }
    var v36 int8 = -55
    fmt.Println(v7, v36, v3)
    var v37 int32 = v7
    v38 := GenS5{F0: -474, F1: 255}
    v39 := GenS6{F0: 167, F1: 4294967295, F2: (v6 ^ (v6 & 42))}
    v40 := GenS7{F0: 65535, F1: 2, F2: 4294967295}
    v41 := GenS8{F0: v37}
    if false {
        switch 0 {
        case 1:
            fmt.Println(v6, v3)
        case -2147483648:
            _ = len(v1)
        case 729:
            if len(v1) > 0 {
                v1[0] = 722
            }
        default:
            v42 := v2[427]
            _ = v42
        }
    }
    switch "foo bar" {
    case "foo bar":
        v43 := GenS9{F0: 142, F1: "a"}
        _ = v43
    case "":
        var v44 uint64 = 0
        _ = v44
    default:
        if (-1 >= 9223372036854775807) {
            var v45 int8 = (-(v36))
            _ = v45
        }
    }
    v46 := GenS10{F0: 0}
    if (42 >= 108) {
        for fi47 := 0; fi47 < 20; fi47++ {
            if (v6 <= v6) {
                var v48 uint8 = 1
                if len(v1) > 0 {
                    v1[0] = 1
                }
                _ = v48
            } else {
                for fi49 := 0; fi49 < 12; fi49++ {
                    switch -666 {
                    case 32768:
                        v50 := GenS11{F0: 0, F1: 127}
                        _ = v50
                    case -371:
                        fmt.Println(v37)
                    default:
                        switch "\n\t" {
                        case "foo bar":
                            var v51 float64 = float64((float32(1e10) - 1.5))
                            _ = v51
                        case "hello":
                            var v52 float32 = 0.1
                            _ = v52
                        case "":
                            _ = len(v4)
                        }
                        v6 -= v6
                        v36 += 42
                    }
                    switch fi49 {
                    case -9:
                        _ = len(v4)
                    case -528:
                        v3 = (v3 % ((42) | 1))
                    case 0:
                        var v53 float32 = 0.1
                        _ = v53
                    default:
                        v36 += v36
                    }
                    v54 := GenS5{F0: int16(v6), F1: 2}
                    _ = fi49
                    _ = v54
                }
                v55 := GenS12{F0: 63, F1: (fi47 + fi47), F2: 65535}
                _ = v55
            }
            if (false && false) {
                var v56 int = (fi47 ^ (fi47 | fi47))
                _ = v56
            } else {
                v37 = v37
            }
            _ = fi47
        }
        switch 128 {
        case 1:
            v57 := GenS4{F0: 346, F1: 2147483647, F2: 32767}
            _ = v57
        case -1:
            var v58 int16 = 188
            _ = v58
        default:
            v59 := GenS4{F0: int16(v37), F1: 2147483647, F2: -499}
            _ = v59
        }
    } else {
        v6 = v6
        fmt.Println(v37, v7)
        v36 += -1
    }
    switch 1 {
    case 220:
        v36 = v36
    case 32768:
        v37 -= (^(v7))
    case 4294967295:
        switch -661 {
        case 32768:
            var v60 uint32 = 65535
            _ = v60
        case -9223372036854775808:
            v61 := v2[65535]
            _ = v61
        case 4294967295:
            for fi62 := 0; fi62 < 3; fi62++ {
                if (true && true) {
                    v4 = append(v4, "hello")
                    var v63 int = (-(fi62))
                    if false {
                        v63 = 128
                        var v64 string = "hello"
                        var v65 uint = (v6 ^ 24)
                        _ = v64
                        _ = v65
                    } else {
                        for fi66 := 0; fi66 < 19; fi66++ {
                            v2[4294967295] = 935
                            var v67 float32 = ((0.0 + (1e10 + 1e10)) / 1e10)
                            _ = fi66
                            _ = v67
                        }
                        var v68 string = (("hello" + ("foo bar" + (("foo bar" + "") + ("hello" + ("hello" + "a"))))) + "hello")
                        _ = v68
                    }
                    _ = v63
                } else {
                    switch "" {
                    case "":
                        v6 += 255
                    }
                    var v69 int8 = -128
                    v37 = (v37 | (v3 ^ 1))
                    _ = v69
                }
                _ = fi62
            }
        default:
            for fi70 := 0; fi70 < 7; fi70++ {
                var v71 int16 = -32768
                var v72 uint = (^(v6))
                fmt.Println(fi70)
                _ = fi70
                _ = v71
                _ = v72
            }
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v46
    _ = v5
    _ = v6
    _ = v7
}
