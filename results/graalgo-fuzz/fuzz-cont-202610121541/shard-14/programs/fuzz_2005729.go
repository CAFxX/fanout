package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint8; F1 int8 }
type GenS5 struct { F0 int8; F1 int }
type GenS6 struct { F0 uint32; F1 int64 }
type GenS7 struct { F0 uint64; F1 int16 }
type GenS8 struct { F0 int64; F1 uint8; F2 uint16 }
type GenS9 struct { F0 int16; F1 int64 }
func genFn0(p0 uint32, p1 uint32) string {
    var v4 uint = 244
    for fi5 := 0; fi5 < 17; fi5++ {
        p0 += p0
        _ = fi5
    }
    v6 := GenS1{F0: 1000}
    v7 := GenS2{F0: (552 > 0)}
    fmt.Println(v4, p1)
    _ = p0
    _ = p1
    _ = v4
    _ = v6
    _ = v7
    return "foo bar"
}

func main() {
    var v1 uint = 255
    v2 := []string{}
    v3 := []uint32{}
    for fi8 := 0; fi8 < 18; fi8++ {
        v1 = 948
        switch ((("hello" + ("foo bar" + (((("" + "foo bar") + "foo bar") + "") + "\n\t"))) + ((("" + "\n\t") + "a") + ("hello" + ("foo bar" + (("hello" + "\n\t") + ""))))) + ("hello" + "foo bar")) {
        case "a":
            var v9 uint64 = 1
            _ = v9
        }
        _ = fi8
    }
    fmt.Println(v1)
    fmt.Println(v1)
    v1 = (v1 & v1)
    var v10 uint = v1
    if false {
        if (!false) {
            switch "foo bar" {
            case "hello":
                if (true || (!true)) {
                    v10 = 0
                } else {
                    v11 := GenS3{F0: 42}
                    if (2 <= 42) {
                        if ((uint64(v1) != 2) || true) {
                            var v12 uint8 = 42
                            _ = v12
                        } else {
                            switch "hello" {
                            case "hello":
                                v10 = v1
                            case "\n\t":
                                var v13 uint32 = 308
                                _ = v13
                            case "":
                                for fi14 := 0; fi14 < 14; fi14++ {
                                    for fi15 := 0; fi15 < 12; fi15++ {
                                        _ = len(v3)
                                        v16 := GenS2{F0: (!(true && false))}
                                        _ = fi15
                                        _ = v16
                                    }
                                    _ = fi14
                                }
                            }
                        }
                        for fi17 := 0; fi17 < 4; fi17++ {
                            var v18 uint = 65535
                            v1 += 2
                            switch (fi17 % ((-659) | 1)) {
                            case -9223372036854775808:
                                v19 := GenS2{F0: true}
                                _ = v19
                            }
                            _ = fi17
                            _ = v18
                        }
                        if (251 != 769) {
                            v1 += v1
                            var v20 uint8 = 128
                            for fi21 := 0; fi21 < 7; fi21++ {
                                v1 += 13
                                _ = fi21
                            }
                            _ = v20
                        }
                    } else {
                        v1 -= v10
                    }
                    _ = v11
                }
            case "a":
                v22 := GenS4{F0: 128, F1: 98}
                _ = v22
            }
            fmt.Println(v10, v1)
        } else {
            switch 49 {
            case 33:
                if (!(1 >= -1)) {
                    for fi23 := 0; fi23 < 17; fi23++ {
                        if (-32768 < 1) {
                            var v24 bool = false
                            v24 = (1000 > 481)
                            for fi25 := 0; fi25 < 9; fi25++ {
                                var v26 int32 = int32(v10)
                                _ = fi25
                                _ = v26
                            }
                            _ = v24
                        } else {
                            if true {
                                if (float32(3.14159) <= 0.1) {
                                    switch fi23 {
                                    case -573:
                                        for fi27 := 0; fi27 < 10; fi27++ {
                                            v1 = (v1 | (v10 & v1))
                                            _ = fi27
                                        }
                                    case 1:
                                        v10 -= v10
                                    case 127:
                                        for fi28 := 0; fi28 < 14; fi28++ {
                                            var v29 uint16 = 65535
                                            v10 -= (v10 / (((v1 << (((v1 / (((v10 * (v1 % ((v10) | 1)))) | 1))) & 63))) | 1))
                                            _ = fi28
                                            _ = v29
                                        }
                                    default:
                                        for fi30 := 0; fi30 < 19; fi30++ {
                                            _ = len(v2)
                                            _ = fi30
                                        }
                                    }
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    fmt.Println(v1, v10, v1)
                                } else {
                                    v1 += v1
                                    fmt.Println(fi23, v1)
                                }
                                var v31 float32 = float32(0.0)
                                _ = len(v3)
                                _ = v31
                            }
                            switch -40 {
                            case 4294967295:
                                fmt.Println(fi23, v1, fi23)
                            case 273:
                                v32 := GenS5{F0: 0, F1: -1}
                                _ = v32
                            }
                        }
                        v3 = append(v3, 959)
                        var v33 uint32 = 2147483647
                        _ = fi23
                        _ = v33
                    }
                    fmt.Println(v10)
                }
            case 1:
                var v34 int16 = 0
                _ = v34
            }
            v1 += 43
        }
        v1 = 1
    } else {
        v35 := GenS2{F0: (false && false)}
        if ((false && ((float64(1e10) + 3.14159) <= 0.1)) || (true && true)) {
            _ = len(v2)
        }
        switch "hello" {
        case "a":
            v1 = (v1 ^ (v1 + v1))
        case "foo bar":
            for fi36 := 0; fi36 < 10; fi36++ {
                v37 := GenS6{F0: 163, F1: -9223372036854775808}
                _ = fi36
                _ = v37
            }
        case "":
            for fi38 := 0; fi38 < 13; fi38++ {
                _ = len(v2)
                var v39 int16 = 0
                v40 := GenS6{F0: 4294967295, F1: 42}
                _ = fi38
                _ = v39
                _ = v40
            }
        default:
            var v41 uint64 = 312
            _ = v41
        }
        _ = v35
    }
    v1 = v1
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if (240 > 460) {
        v10 = v1
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if (32767 <= 255) {
        var v42 uint64 = 1
        if false {
            var v43 bool = (142 == 699)
            v44 := GenS6{F0: 931, F1: -934}
            var v45 int = 9223372036854775807
            _ = v43
            _ = v44
            _ = v45
        } else {
            var v46 int16 = -1
            if (true && false) {
                fmt.Println(v10, v42)
                if true {
                    fmt.Println(v46, v1, v46)
                    fmt.Println(v46)
                    switch "hello" {
                    case "foo bar":
                        fmt.Println(v10, v42, v46)
                    case "":
                        switch 1 {
                        case 564:
                            v1 -= v10
                        case 42:
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        case -46:
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        default:
                            fmt.Println(v42, v46, v42)
                        }
                    }
                } else {
                    for fi47 := 0; fi47 < 5; fi47++ {
                        v48 := GenS7{F0: (v42 - 2), F1: (v46 - v46)}
                        if len(v2) > 0 {
                            v2[0] = ("hello" + "hello")
                        }
                        var v49 uint8 = 42
                        _ = fi47
                        _ = v48
                        _ = v49
                    }
                    if len(v3) > 0 {
                        v3[0] = 1
                    }
                    for fi50 := 0; fi50 < 14; fi50++ {
                        v10 = (v10 * (v10 ^ (v10 % (((v10 >> (((v10 & v10)) & 63))) | 1))))
                        if len(v3) > 0 {
                            v3[0] = 875
                        }
                        _ = fi50
                    }
                }
            }
            _ = v46
        }
        if false {
            var v51 uint = v10
            if (1000 > 670) {
                v52 := GenS6{F0: 811, F1: 1}
                _ = v52
            } else {
                switch 42 {
                case 32767:
                    fmt.Println(v1, v51)
                case 157:
                    switch ("foo bar" + "\n\t") {
                    case "a":
                        if false {
                            var v53 float64 = 0.1
                            if false {
                                _ = len(v2)
                                var v54 float32 = 0.0
                                _ = v54
                            } else {
                                v55 := GenS5{F0: -1, F1: -239}
                                fmt.Println(v53, v42)
                                var v56 float64 = float64((0.0 - 1.5))
                                _ = v55
                                _ = v56
                            }
                            _ = v53
                        } else {
                            switch "\n\t" {
                            case "":
                                _ = len(v3)
                            case "hello":
                                v57 := GenS7{F0: v42, F1: -68}
                                _ = v57
                            }
                            v1 += (v51 / ((v51) | 1))
                        }
                    case "\n\t":
                        v1 = v1
                    case "foo bar":
                        if false {
                            for fi58 := 0; fi58 < 14; fi58++ {
                                v51 = (v51 & 65535)
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                _ = fi58
                            }
                        } else {
                            for fi59 := 0; fi59 < 3; fi59++ {
                                v42 = v42
                                _ = fi59
                            }
                        }
                    default:
                        if true {
                            if ((v10 * 450) == v10) {
                                var v60 int = 32767
                                _ = v60
                            }
                            v51 = 209
                        }
                    }
                case 127:
                    if len(v3) > 0 {
                        v3[0] = uint32(v10)
                    }
                }
                _ = len(v3)
                if false {
                    for fi61 := 0; fi61 < 3; fi61++ {
                        fmt.Println(v10, v42)
                        for fi62 := 0; fi62 < 5; fi62++ {
                            _ = len(v3)
                            _ = fi62
                        }
                        _ = fi61
                    }
                    v63 := GenS8{F0: 501, F1: 255, F2: 474}
                    if (65535 == 652) {
                        var v64 uint8 = 255
                        v65 := GenS9{F0: -1, F1: int64(v1)}
                        _ = v64
                        _ = v65
                    }
                    _ = v63
                }
            }
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = v51
        } else {
            if (!false) {
                fmt.Println(v1)
            } else {
                v42 += 18446744073709551615
                var v66 int16 = 831
                _ = v66
            }
        }
        _ = v42
    }
    v10 = 958
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
}
