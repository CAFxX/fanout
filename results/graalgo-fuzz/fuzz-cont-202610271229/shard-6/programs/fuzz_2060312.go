package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint64; F1 int8; F2 uint16 }
type GenS3 struct { F0 bool; F1 int32; F2 uint }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 string; F1 uint }
type GenS6 struct { F0 int8; F1 int }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 uint16; F1 int8; F2 int16 }
func genFn0(p0 string, p1 uint8) uint64 {
    for fi5 := 0; fi5 < 9; fi5++ {
        for fi6 := 0; fi6 < 16; fi6++ {
            var v7 uint8 = 205
            _ = fi6
            _ = v7
        }
        _ = fi5
    }
    v8 := GenS1{F0: false}
    fmt.Println(p1, p0)
    _ = p0
    _ = p1
    _ = v8
    return 1
}

func main() {
    var v1 uint32 = 42
    v2 := []int{-9223372036854775808, -371}
    v3 := []uint32{v1, v1}
    var v4 int32 = 0
    fmt.Println(v1, v1)
    for fi9 := 0; fi9 < 9; fi9++ {
        v4 = v4
        _ = fi9
    }
    v10 := GenS1{F0: (!false)}
    fmt.Println(v1, v4)
    var v11 int64 = 216
    for fi12 := 0; fi12 < 4; fi12++ {
        if len(v2) > 0 {
            v2[0] = fi12
        }
        if (127 == 0) {
            v11 = (v11 % ((9223372036854775807) | 1))
            v3 = append(v3, v1)
        }
        _ = fi12
    }
    v13 := GenS2{F0: 983, F1: 127, F2: 146}
    v14 := GenS2{F0: 18446744073709551615, F1: -124, F2: 255}
    v4 = v4
    switch ("" + ("foo bar" + "foo bar")) {
    case "foo bar":
        fmt.Println(v11)
    case "hello":
        v15 := GenS2{F0: 18446744073709551615, F1: 1, F2: 42}
        _ = v15
    default:
        var v16 float32 = (((-2.25 / (0.1 + (((float32(((3.14159 - (float64((0.1 - 1e10)) - float64((-2.25 / (((0.1 / float32(-2.25)) - (float32(3.14159) / 1.5)) / 1.0))))) - 1.5)) / 1.5) * -2.25) * 0.0))) * 1.5) / (0.0 - 3.14159))
        _ = v16
    }
    for fi17 := 0; fi17 < 13; fi17++ {
        if len(v3) > 0 {
            v3[0] = 709
        }
        _ = fi17
    }
    v18 := GenS3{F0: true, F1: (v4 ^ (^(v4))), F2: 65535}
    var v19 uint32 = (^(v1))
    if false {
        var v20 int8 = 42
        var v21 int8 = (v20 >> ((255) & 63))
        for fi22 := 0; fi22 < 15; fi22++ {
            switch (fi22 + int(v1)) {
            case 899:
                if (v20 != v21) {
                    var v23 uint16 = 65535
                    switch "" {
                    case "a":
                        var v24 string = ("foo bar" + "hello")
                        _ = v24
                    case "hello":
                        var v25 int8 = (v20 - v20)
                        _ = v25
                    }
                    v20 += v21
                    _ = v23
                } else {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                }
            case 32767:
                _ = len(v3)
            case 128:
                v4 = int32(fi22)
            }
            _ = fi22
        }
        _ = v20
        _ = v21
    } else {
        switch "foo bar" {
        case "foo bar":
            if false {
                fmt.Println(v19, v11, v1)
                fmt.Println(v4, v19, v1)
                v26 := GenS3{F0: (uint16(v1) <= 255), F1: -1, F2: 943}
                _ = v26
            }
        case "\n\t":
            if false {
                v27 := GenS2{F0: 261, F1: 42, F2: 358}
                fmt.Println(v4, v1, v19)
                v4 -= 1
                _ = v27
            }
        default:
            for fi28 := 0; fi28 < 7; fi28++ {
                v29 := GenS3{F0: (("hello" + (("\n\t" + (((("" + ("\n\t" + "hello")) + "\n\t") + "hello") + (("foo bar" + (("\n\t" + "a") + "hello")) + ""))) + "foo bar")) < "foo bar"), F1: (v4 % (((v4 ^ v4)) | 1)), F2: 42}
                v30 := GenS4{F0: (v19 ^ 2147483647)}
                var v31 int = 9223372036854775807
                _ = fi28
                _ = v29
                _ = v30
                _ = v31
            }
        }
        v11 += -9223372036854775808
    }
    var v32 int16 = 32767
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v33 := GenS4{F0: (v19 >> ((65535) & 63))}
    switch ("hello" + (("a" + "foo bar") + "a")) {
    case "hello":
        if len(v3) > 0 {
            v3[0] = v1
        }
    case "foo bar":
        if false {
            for fi34 := 0; fi34 < 17; fi34++ {
                v35 := GenS5{F0: "foo bar", F1: 1}
                _ = fi34
                _ = v35
            }
            fmt.Println(v4, v11)
            for fi36 := 0; fi36 < 4; fi36++ {
                if (160 < 2) {
                    var v37 uint32 = (v19 ^ (v1 / ((2147483647) | 1)))
                    if len(v3) > 0 {
                        v3[0] = v19
                    }
                    for fi38 := 0; fi38 < 14; fi38++ {
                        v32 += -1
                        v2 = append(v2, 1)
                        if (209 == 1) {
                            if true {
                                if len(v3) > 0 {
                                    v3[0] = (v1 | (v1 & 919))
                                }
                                fmt.Println(v32, v19, v11)
                                v39 := GenS6{F0: 127, F1: (fi36 >> ((898) & 63))}
                                _ = v39
                            } else {
                                var v40 uint16 = 255
                                var v41 uint16 = v40
                                _ = v40
                                _ = v41
                            }
                            v42 := GenS3{F0: true, F1: (v4 + v4), F2: 2}
                            v43 := GenS7{F0: v4}
                            var v44 int64 = 42
                            _ = v42
                            _ = v43
                            _ = v44
                        }
                        _ = fi38
                    }
                    _ = v37
                } else {
                    var v45 float32 = 1e10
                    var v46 uint8 = 127
                    var v47 float64 = ((1e10 / (0.0 - (1e10 - 3.14159))) * (3.14159 / 3.14159))
                    _ = v45
                    _ = v46
                    _ = v47
                }
                v3 = append(v3, (v19 % ((v1) | 1)))
                _ = fi36
            }
        } else {
            for fi48 := 0; fi48 < 13; fi48++ {
                switch "foo bar" {
                case "a":
                    if true {
                        var v49 int8 = -128
                        _ = v49
                    }
                case "foo bar":
                    v50 := GenS6{F0: 100, F1: (fi48 * (fi48 + fi48))}
                    _ = v50
                case "hello":
                    for fi51 := 0; fi51 < 12; fi51++ {
                        v1 = (^(v1))
                        _ = len(v3)
                        switch "" {
                        case "hello":
                            v52 := GenS8{F0: 623, F1: 127, F2: v32}
                            _ = v52
                        case "":
                            switch (fi51 + -1) {
                            case 2147483647:
                                v1 -= v19
                            default:
                                if false {
                                    var v53 uint16 = 255
                                    fmt.Println(v11, v53, v4)
                                    _ = v53
                                } else {
                                    var v54 int = -9223372036854775808
                                    if false {
                                        v19 = 4294967295
                                    }
                                    _ = v54
                                }
                            }
                        case "a":
                            var v55 uint16 = 0
                            _ = v55
                        default:
                            if true {
                                var v56 int16 = (v32 & v32)
                                _ = v56
                            } else {
                                var v57 uint32 = (^(v1))
                                fmt.Println(fi48, v4)
                                v4 = (v4 ^ (v4 * v4))
                                _ = v57
                            }
                        }
                        _ = fi51
                    }
                }
                for fi58 := 0; fi58 < 3; fi58++ {
                    var v59 uint16 = 65535
                    _ = fi58
                    _ = v59
                }
                _ = fi48
            }
            var v60 int8 = -38
            _ = v60
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v13
    _ = v14
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v4
}
