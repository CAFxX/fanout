package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32 }
type GenS2 struct { F0 uint8; F1 uint64 }
type GenS3 struct { F0 uint16; F1 uint8; F2 int16 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint64 }

func main() {
    var v1 int64 = 1
    v2 := []uint32{4294967295, 65535, 4294967295}
    v3 := make(map[uint32]int)
    var v4 uint = 42
    v5 := make(map[string]uint32)
    var v6 int = -208
    if len(v2) > 0 {
        v2[0] = 2147483647
    }
    var v7 string = "foo bar"
    if false {
        v4 -= v4
        var v8 int8 = 1
        _ = v8
    }
    if false {
        for fi9 := 0; fi9 < 9; fi9++ {
            _ = len(v2)
            v10 := GenS1{F0: v1, F1: 2147483647}
            _ = fi9
            _ = v10
        }
        fmt.Println(v7, v4, v4)
    }
    if (1 <= (-(v6))) {
        v4 -= 255
    } else {
        if false {
            var v11 int16 = -146
            _ = v11
        } else {
            for fi12 := 0; fi12 < 16; fi12++ {
                v6 -= (-(v6))
                var v13 uint32 = 0
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = fi12
                _ = v13
            }
            var v14 float64 = ((((-2.25 + (1e10 + 1.5)) / float64((((float32((-2.25 + -2.25)) - 0.0) - float32(1e10)) + (float32(3.14159) - 1e10)))) + 0.1) / 3.14159)
            _ = v14
        }
        delete(v5, (((v7 + v7) + v7) + ""))
    }
    fmt.Println(v7, v4, v4)
    var v15 int32 = 446
    var v16 uint16 = 0
    var v17 int = (v6 % ((-15) | 1))
    if (((v17 <= (v6 % ((v6) | 1))) && (!(true || true))) && (127 > int8(v4))) {
        switch v7 {
        case "a":
            if len(v2) > 0 {
                v2[0] = 4294967295
            }
        case "foo bar":
            _ = len(v2)
        }
    } else {
        v7 = ("\n\t" + v7)
    }
    var v18 uint8 = 253
    delete(v3, 0)
    v19 := GenS1{F0: -9223372036854775808, F1: (v15 / ((v15) | 1))}
    if (!false) {
        if false {
            _ = len(v2)
            switch (v7 + ("hello" + "hello")) {
            case "\n\t":
                fmt.Println(v6, v4)
            case "foo bar":
                if len(v2) > 0 {
                    v2[0] = 2
                }
            case "hello":
                fmt.Println(v7, v16)
            }
        }
        switch v17 {
        case -1:
            switch v17 {
            case 42:
                if false {
                    v20 := GenS1{F0: v1, F1: -517}
                    _ = v20
                }
            default:
                v18 += (^(v18))
            }
        case -9223372036854775808:
            v21 := v5[("" + "")]
            _ = v21
        }
        if false {
            _ = len(v5)
        }
    } else {
        fmt.Println(v6)
        switch (v7 + ((v7 + v7) + v7)) {
        case "":
            v18 += 105
        case "hello":
            v6 = int(v18)
        case "\n\t":
            fmt.Println(v4, v1)
        default:
            switch ("hello" + v7) {
            case "\n\t":
                v22 := GenS1{F0: (v1 ^ 0), F1: (v15 + (v15 + 42))}
                _ = v22
            case "foo bar":
                var v23 float32 = (((0.1 + -2.25) * float32((3.14159 + float64(1e10)))) / float32((3.14159 / 1.5)))
                _ = v23
            }
        }
    }
    for fi24 := 0; fi24 < 18; fi24++ {
        for fi25 := 0; fi25 < 18; fi25++ {
            if false {
                v26 := GenS2{F0: v18, F1: 603}
                if (!false) {
                    var v27 int8 = 127
                    var v28 uint32 = 987
                    _ = v27
                    _ = v28
                }
                _ = v26
            }
            _ = fi25
        }
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = fi24
    }
    if (((671 < 612) || true) || false) {
        if false {
            switch "hello" {
            case "hello":
                for fi29 := 0; fi29 < 10; fi29++ {
                    var v30 uint16 = v16
                    _ = fi29
                    _ = v30
                }
            case "\n\t":
                for fi31 := 0; fi31 < 4; fi31++ {
                    fmt.Println(v1)
                    _ = fi31
                }
            case "":
                v32 := v3[858]
                _ = v32
            }
        } else {
            v5[v7] = 713
            if (0 >= -479) {
                for fi33 := 0; fi33 < 12; fi33++ {
                    for fi34 := 0; fi34 < 17; fi34++ {
                        var v35 int32 = 512
                        _ = fi34
                        _ = v35
                    }
                    _ = fi33
                }
                v36 := GenS1{F0: (-(v1)), F1: v15}
                var v37 int8 = 100
                _ = v36
                _ = v37
            } else {
                _ = len(v2)
                v38 := GenS1{F0: (v1 | v1), F1: v15}
                _ = v38
            }
            if false {
                v39 := GenS1{F0: v1, F1: v15}
                v2 = append(v2, 498)
                if false {
                    v40 := GenS3{F0: v16, F1: v18, F2: 32767}
                    _ = v40
                } else {
                    for fi41 := 0; fi41 < 6; fi41++ {
                        switch "hello" {
                        case "\n\t":
                            switch (((("" + (v7 + "\n\t")) + "hello") + "hello") + v7) {
                            case "hello":
                                for fi42 := 0; fi42 < 12; fi42++ {
                                    if false {
                                        if false {
                                            v18 += (v18 | v18)
                                            v43 := GenS3{F0: (^(v16)), F1: (v18 >> (((v4 >> (((^(v4))) & 63))) & 63)), F2: 356}
                                            _ = v43
                                        }
                                        switch v7 {
                                        case "":
                                            delete(v3, 65535)
                                        case "\n\t":
                                            if (false && (231 >= -992)) {
                                                v6 -= 9223372036854775807
                                                v4 = (v4 + v4)
                                            } else {
                                                if (1.5 <= (3.14159 + 0.1)) {
                                                    switch v17 {
                                                    case -244:
                                                        if true {
                                                            v15 = v15
                                                        }
                                                    case 762:
                                                        var v44 bool = false
                                                        _ = v44
                                                    case -677:
                                                        var v45 uint = (v4 ^ v4)
                                                        v46 := GenS3{F0: (^(v16)), F1: 2, F2: -32768}
                                                        _ = v45
                                                        _ = v46
                                                    }
                                                    fmt.Println(v6, v7)
                                                    var v47 int64 = 763
                                                    _ = v47
                                                } else {
                                                    v48 := v3[42]
                                                    _ = v48
                                                }
                                                v17 -= (v6 >> ((uint(v18)) & 63))
                                                switch 9223372036854775807 {
                                                case -9223372036854775808:
                                                    if true {
                                                        v15 = (v15 & v15)
                                                        if (2 == 65535) {
                                                            if (true || (v15 >= (v15 | -1))) {
                                                                if (false || (!("foo bar" != (((v7 + "a") + "foo bar") + ("hello" + v7))))) {
                                                                    v4 -= (v4 + (v4 >> ((932) & 63)))
                                                                } else {
                                                                    v49 := GenS1{F0: 98, F1: (-(v15))}
                                                                    fmt.Println(v1, v15, v7)
                                                                    v1 -= (v1 & v1)
                                                                    _ = v49
                                                                }
                                                            }
                                                            if true {
                                                                delete(v3, uint32(v4))
                                                                if true {
                                                                    v17 = fi42
                                                                    v50 := v5[((((v7 + "") + v7) + v7) + v7)]
                                                                    _ = v50
                                                                }
                                                            }
                                                            v51 := GenS2{F0: v18, F1: uint64(v18)}
                                                            _ = v51
                                                        }
                                                        v17 = v6
                                                    }
                                                case 619:
                                                    if (377 <= 1) {
                                                        v2 = append(v2, 0)
                                                        switch fi42 {
                                                        case 127:
                                                            v52 := GenS4{F0: (v1 | v1)}
                                                            _ = v52
                                                        }
                                                        v53 := GenS5{F0: 0}
                                                        fmt.Println(v7, v4, v6)
                                                        _ = v53
                                                    }
                                                case 1:
                                                    var v54 int32 = -595
                                                    _ = v54
                                                }
                                            }
                                        default:
                                            var v55 int32 = (v15 % (((v15 * (v15 / ((-2147483648) | 1)))) | 1))
                                            _ = v55
                                        }
                                    }
                                    v56 := v3[2147483647]
                                    if (("hello" + v7) == (v7 + ("\n\t" + ("foo bar" + "foo bar")))) {
                                        v56 = v56
                                    }
                                    _ = fi42
                                    _ = v56
                                }
                            case "":
                                var v57 uint8 = v18
                                _ = v57
                            }
                        }
                        _ = fi41
                    }
                    if len(v2) > 0 {
                        v2[0] = 2147483648
                    }
                    if (-58 == 127) {
                        if false {
                            var v58 int = v6
                            _ = v58
                        }
                        var v59 string = "hello"
                        delete(v5, "a")
                        _ = v59
                    }
                }
                _ = v39
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
