package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint32; F1 uint32 }
type GenS3 struct { F0 uint16; F1 uint64 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 uint8; F1 uint64 }
type GenS7 struct { F0 int8; F1 uint8 }

func main() {
    var v1 int = 1
    var v2 int32 = 2147483647
    v3 := []string{"", "a", ("hello" + ("foo bar" + "\n\t")), ("foo bar" + (("a" + ("" + "foo bar")) + ("foo bar" + (("" + "hello") + "\n\t"))))}
    for fi4 := 0; fi4 < 13; fi4++ {
        var v5 float32 = 1.5
        fmt.Println(v2, v2, v5)
        if (!true) {
            for fi6 := 0; fi6 < 10; fi6++ {
                fmt.Println(fi6, v5)
                _ = fi6
            }
            v7 := GenS1{F0: -93}
            if (1e10 >= 1.5) {
                var v8 float64 = (1e10 / (0.0 + 1e10))
                var v9 string = ("" + "a")
                _ = v8
                _ = v9
            } else {
                if len(v3) > 0 {
                    v3[0] = "hello"
                }
                if (0 <= 0) {
                    if (v5 <= (((1.5 - ((-2.25 + 3.14159) * -2.25)) - v5) + v5)) {
                        for fi10 := 0; fi10 < 15; fi10++ {
                            var v11 int32 = (v2 - (v2 - v2))
                            _ = fi10
                            _ = v11
                        }
                        v1 = v1
                    }
                }
                fmt.Println(v2, v1, v5)
            }
            _ = v7
        } else {
            var v12 int = -685
            if false {
                switch v1 {
                case 0:
                    var v13 uint16 = 65535
                    _ = v13
                default:
                    if false {
                        fmt.Println(v1)
                        v12 = fi4
                    } else {
                        v2 += v2
                    }
                }
                v3 = append(v3, (("a" + "a") + "a"))
            }
            if len(v3) > 0 {
                v3[0] = "hello"
            }
            _ = v12
        }
        _ = fi4
        _ = v5
    }
    var v14 string = ""
    v14 = "hello"
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    for fi15 := 0; fi15 < 13; fi15++ {
        v16 := GenS1{F0: -128}
        if (127 == 68) {
            v17 := GenS1{F0: 57}
            var v18 int32 = (v2 ^ (v2 + (v2 - (v2 * (v2 << ((1) & 63))))))
            for fi19 := 0; fi19 < 11; fi19++ {
                var v20 uint8 = 14
                _ = fi19
                _ = v20
            }
            _ = v17
            _ = v18
        } else {
            v3 = append(v3, v14)
        }
        _ = fi15
        _ = v16
    }
    if true {
        switch 42 {
        case 127:
            var v21 int64 = 42
            _ = v21
        case 32767:
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        default:
            v14 = (((v14 + (v14 + (("hello" + v14) + v14))) + ("a" + ("foo bar" + ""))) + v14)
        }
        _ = len(v3)
    }
    v22 := GenS1{F0: -1}
    v1 -= (v1 ^ v1)
    v23 := GenS2{F0: 853, F1: 0}
    var v24 uint = 0
    var v25 int64 = 1
    if (3.14159 >= ((1e10 - 0.1) - float32(3.14159))) {
        v1 = (v1 + v1)
        var v26 int16 = 1
        if false {
            if false {
                v1 += (^(v1))
                v27 := GenS3{F0: 0, F1: 18446744073709551615}
                _ = v27
            } else {
                switch (-(v1)) {
                case 2147483647:
                    v28 := GenS2{F0: 42, F1: 65535}
                    _ = v28
                case 128:
                    for fi29 := 0; fi29 < 15; fi29++ {
                        for fi30 := 0; fi30 < 12; fi30++ {
                            v31 := GenS4{F0: true}
                            _ = fi30
                            _ = v31
                        }
                        _ = fi29
                    }
                case 99:
                    _ = len(v3)
                default:
                    if false {
                        if true {
                            fmt.Println(v24)
                            v32 := GenS5{F0: uint8(v24)}
                            v33 := GenS3{F0: 32767, F1: 0}
                            _ = v32
                            _ = v33
                        } else {
                            if true {
                                if false {
                                    v14 = (v14 + (v14 + ("hello" + v14)))
                                    v1 -= (v1 & int(v25))
                                } else {
                                    if true {
                                        fmt.Println(v25)
                                        for fi34 := 0; fi34 < 12; fi34++ {
                                            switch ((((v14 + "foo bar") + "foo bar") + v14) + ((("a" + (v14 + v14)) + (("\n\t" + v14) + v14)) + v14)) {
                                            case "\n\t":
                                                if len(v3) > 0 {
                                                    v3[0] = (v14 + v14)
                                                }
                                            }
                                            var v35 uint64 = 42
                                            v1 -= fi34
                                            _ = fi34
                                            _ = v35
                                        }
                                    }
                                    var v36 int8 = -1
                                    _ = v36
                                }
                            }
                        }
                        v3 = append(v3, (("foo bar" + "\n\t") + "foo bar"))
                    }
                }
                fmt.Println(v1, v26)
            }
            if len(v3) > 0 {
                v3[0] = (v14 + "hello")
            }
            if (v24 != v24) {
                if true {
                    v25 -= (v25 - 1)
                    fmt.Println(v1)
                }
                switch (^(v1)) {
                case -9223372036854775808:
                    var v37 int = v1
                    _ = v37
                case -500:
                    v14 = v14
                case 32768:
                    var v38 float64 = float64(1e10)
                    _ = v38
                default:
                    if (true && false) {
                        fmt.Println(v1, v24)
                        switch v1 {
                        case -168:
                            fmt.Println(v26)
                        case 565:
                            _ = len(v3)
                        case 0:
                            v26 = int16(v25)
                        }
                    }
                }
            } else {
                var v39 uint8 = 128
                v3 = append(v3, v14)
                var v40 bool = (!(((v2 << (((v24 - v24)) & 63)) <= (v2 / ((v2) | 1))) && false))
                _ = v39
                _ = v40
            }
        } else {
            v24 += 951
            switch (v1 * v1) {
            case 996:
                v41 := GenS6{F0: 0, F1: 59}
                _ = v41
            case 9223372036854775807:
                var v42 int8 = 64
                _ = v42
            default:
                var v43 uint64 = 0
                _ = v43
            }
            var v44 uint64 = 1
            _ = v44
        }
        _ = v26
    } else {
        switch v1 {
        case 42:
            for fi45 := 0; fi45 < 6; fi45++ {
                v25 = (-(v25))
                _ = fi45
            }
        case 0:
            var v46 int = (v1 - 1)
            _ = v46
        }
        var v47 uint32 = 0
        _ = v47
    }
    for fi48 := 0; fi48 < 16; fi48++ {
        v49 := GenS5{F0: 0}
        fmt.Println(v14, v24, v1)
        _ = len(v3)
        _ = fi48
        _ = v49
    }
    switch 2147483647 {
    case -591:
        if (-115 > -105) {
            var v50 int32 = (v2 >> ((0) & 63))
            fmt.Println(v14, v24, v24)
            _ = v50
        } else {
            var v51 uint8 = 0
            v2 += v2
            _ = v51
        }
    case -534:
        switch "a" {
        case "foo bar":
            for fi52 := 0; fi52 < 9; fi52++ {
                fmt.Println(v24)
                fmt.Println(fi52, v25, v24)
                _ = fi52
            }
        case "a":
            var v53 uint64 = 1
            _ = v53
        case "":
            switch 766 {
            case 0:
                v25 = 1
            case 225:
                var v54 int32 = -2147483648
                _ = v54
            case 9223372036854775807:
                v1 = 42
            default:
                for fi55 := 0; fi55 < 1; fi55++ {
                    fmt.Println(v25, v2)
                    var v56 bool = (1 >= 201)
                    _ = fi55
                    _ = v56
                }
            }
        default:
            switch 32767 {
            case 901:
                v25 = 9223372036854775807
            }
        }
    default:
        switch (v1 & 32767) {
        case -1:
            if true {
                for fi57 := 0; fi57 < 7; fi57++ {
                    fmt.Println(v25)
                    var v58 bool = ((v1 - fi57) != (v1 * fi57))
                    _ = fi57
                    _ = v58
                }
                fmt.Println(v1, v24, v14)
            } else {
                switch "hello" {
                case "":
                    if len(v3) > 0 {
                        v3[0] = v14
                    }
                }
            }
        case 32768:
            for fi59 := 0; fi59 < 4; fi59++ {
                v60 := GenS7{F0: 100, F1: 8}
                v14 = v14
                _ = fi59
                _ = v60
            }
        default:
            if ((v1 + -641) <= v1) {
                fmt.Println(v1, v24, v25)
                for fi61 := 0; fi61 < 13; fi61++ {
                    var v62 float64 = 3.14159
                    switch (v14 + (("hello" + (v14 + v14)) + "a")) {
                    case "\n\t":
                        v25 -= (^(v25))
                    case "foo bar":
                        v1 = v1
                    }
                    fmt.Println(v2, v24, fi61)
                    _ = fi61
                    _ = v62
                }
            }
        }
    }
    v63 := GenS2{F0: 42, F1: 2147483648}
    for fi64 := 0; fi64 < 2; fi64++ {
        v3 = append(v3, (v14 + v14))
        v2 = 2147483647
        _ = fi64
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v14
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v63
}
