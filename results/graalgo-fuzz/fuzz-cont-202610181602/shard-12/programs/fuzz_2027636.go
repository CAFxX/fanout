package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string; F2 bool }
type GenS2 struct { F0 uint16; F1 int64; F2 int64 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 string; F1 int64 }
type GenS5 struct { F0 uint8; F1 int8 }
type GenS6 struct { F0 string; F1 uint }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 uint64 }
type GenS9 struct { F0 uint; F1 int64 }

func main() {
    var v1 uint16 = 32767
    var v2 int16 = 32767
    var v3 uint32 = 4294967295
    v4 := []uint32{}
    var v5 int32 = 2147483647
    _ = len(v4)
    var v6 bool = (!((1 == -128) || false))
    switch ("hello" + "a") {
    case "foo bar":
        for fi7 := 0; fi7 < 14; fi7++ {
            var v8 float32 = 3.14159
            var v9 uint8 = 170
            _ = fi7
            _ = v8
            _ = v9
        }
    default:
        var v10 int64 = 0
        _ = v10
    }
    for fi11 := 0; fi11 < 12; fi11++ {
        if true {
            v1 = (v1 >> ((4294967295) & 63))
        }
        var v12 uint16 = (v1 >> ((661) & 63))
        _ = fi11
        _ = v12
    }
    switch "a" {
    case "\n\t":
        switch "a" {
        case "hello":
            switch ("\n\t" + ("\n\t" + "a")) {
            case "a":
                v13 := GenS1{F0: (v5 - v5), F1: (("foo bar" + ((("\n\t" + "a") + "foo bar") + ((("\n\t" + "a") + (((((("foo bar" + "foo bar") + "") + "") + "foo bar") + "hello") + ((("hello" + ("a" + "a")) + "hello") + ("foo bar" + ("" + "hello"))))) + ("foo bar" + "\n\t")))) + ("" + ("a" + ((("hello" + "hello") + "hello") + "foo bar")))), F2: v6}
                _ = v13
            case "":
                for fi14 := 0; fi14 < 18; fi14++ {
                    var v15 float32 = ((1.5 * 1e10) * -2.25)
                    _ = fi14
                    _ = v15
                }
            default:
                v16 := GenS1{F0: v5, F1: "\n\t", F2: (3.14159 <= 0.1)}
                _ = v16
            }
        case "a":
            switch "" {
            case "hello":
                switch "\n\t" {
                case "":
                    var v17 int32 = (-(v5))
                    _ = v17
                case "a":
                    switch "foo bar" {
                    case "\n\t":
                        for fi18 := 0; fi18 < 20; fi18++ {
                            var v19 int16 = v2
                            switch (("a" + "\n\t") + "\n\t") {
                            case "\n\t":
                                for fi20 := 0; fi20 < 18; fi20++ {
                                    v21 := GenS1{F0: (v5 & (v5 ^ v5)), F1: "a", F2: ((-2.25 / 1e10) == (3.14159 * 1.5))}
                                    _ = fi20
                                    _ = v21
                                }
                            case "a":
                                switch "" {
                                case "\n\t":
                                    _ = len(v4)
                                case "a":
                                    v1 = (v1 << ((uint(v3)) & 63))
                                default:
                                    v19 += (v2 % ((v2) | 1))
                                }
                            case "foo bar":
                                var v22 uint64 = 0
                                _ = v22
                            }
                            _ = fi18
                            _ = v19
                        }
                    case "a":
                        if (32767 != v1) {
                            if (!(9223372036854775807 == 1)) {
                                fmt.Println(v1, v6, v6)
                                v23 := GenS1{F0: (v5 * (v5 & 42)), F1: "a", F2: true}
                                if false {
                                    for fi24 := 0; fi24 < 16; fi24++ {
                                        fmt.Println(v6)
                                        if (v6 || (!("hello" <= ("a" + "\n\t")))) {
                                            var v25 string = "\n\t"
                                            _ = v25
                                        } else {
                                            for fi26 := 0; fi26 < 20; fi26++ {
                                                v27 := GenS1{F0: (v5 - v5), F1: ("foo bar" + "foo bar"), F2: (true || v6)}
                                                var v28 uint = 843
                                                _ = fi26
                                                _ = v27
                                                _ = v28
                                            }
                                        }
                                        v29 := GenS2{F0: (v1 | 661), F1: -416, F2: -188}
                                        _ = fi24
                                        _ = v29
                                    }
                                    switch -2147483648 {
                                    case 42:
                                        if false {
                                            switch 2147483647 {
                                            case 167:
                                                fmt.Println(v6, v1, v3)
                                            case 428:
                                                for fi30 := 0; fi30 < 5; fi30++ {
                                                    var v31 uint = 2
                                                    switch (("foo bar" + "foo bar") + (("foo bar" + ("hello" + ("foo bar" + ("hello" + (("\n\t" + ("hello" + "\n\t")) + "hello"))))) + ((((((("" + "foo bar") + "") + ("\n\t" + "foo bar")) + "a") + "\n\t") + "\n\t") + ""))) {
                                                    case "\n\t":
                                                        v1 += 65535
                                                    case "hello":
                                                        v5 = (v5 & v5)
                                                    default:
                                                        v32 := GenS3{F0: fi30}
                                                        _ = v32
                                                    }
                                                    _ = fi30
                                                    _ = v31
                                                }
                                            }
                                        }
                                    }
                                }
                                _ = v23
                            }
                        }
                    }
                default:
                    for fi33 := 0; fi33 < 2; fi33++ {
                        var v34 float32 = (((-2.25 / (1.5 - (0.1 - (1.5 * float32(-2.25))))) / (-2.25 * ((-2.25 / 1e10) / -2.25))) + float32(1.5))
                        _ = fi33
                        _ = v34
                    }
                }
            }
        case "\n\t":
            switch -1 {
            case -2147483648:
                v35 := GenS4{F0: "foo bar", F1: 9223372036854775807}
                _ = v35
            case 1:
                if v6 {
                    _ = len(v4)
                }
            default:
                var v36 int16 = (v2 / (((-(v2))) | 1))
                _ = v36
            }
        }
    case "a":
        for fi37 := 0; fi37 < 19; fi37++ {
            if v6 {
                if ((v2 ^ v2) > (-(v2))) {
                    if len(v4) > 0 {
                        v4[0] = 2
                    }
                    v6 = (!((4294967295 >= 92) && v6))
                } else {
                    var v38 float64 = float64(((-2.25 / -2.25) * -2.25))
                    _ = v38
                }
                if v6 {
                    var v39 int8 = 64
                    for fi40 := 0; fi40 < 20; fi40++ {
                        var v41 int32 = -2147483648
                        fmt.Println(v6, v41)
                        switch "a" {
                        case "hello":
                            v39 -= (v39 / (((v39 << ((18446744073709551615) & 63))) | 1))
                        default:
                            var v42 uint8 = 1
                            _ = v42
                        }
                        _ = fi40
                        _ = v41
                    }
                    v2 = v2
                    _ = v39
                }
            }
            for fi43 := 0; fi43 < 11; fi43++ {
                v44 := GenS5{F0: 255, F1: 65}
                v6 = (42 > 0)
                var v45 int32 = (v5 / ((0) | 1))
                v6 = (!(2 > v1))
                _ = fi43
                _ = v44
                _ = v45
            }
            _ = len(v4)
            _ = fi37
        }
    case "hello":
        switch 128 {
        case 9223372036854775807:
            switch 2147483647 {
            case -9223372036854775808:
                v3 += (v3 << ((509) & 63))
            case 32768:
                if ((1e10 / 1e10) <= ((3.14159 + (-2.25 / (1e10 - (float64((0.0 * 3.14159)) * float64(1e10))))) - (0.0 * 3.14159))) {
                    switch 128 {
                    case 9223372036854775807:
                        var v46 bool = (!(4294967295 <= 42))
                        _ = v46
                    case -9223372036854775808:
                        _ = len(v4)
                    default:
                        v4 = append(v4, 0)
                    }
                    var v47 uint8 = 127
                    if true {
                        var v48 uint8 = (v47 >> ((18446744073709551615) & 63))
                        v49 := GenS6{F0: "hello", F1: 65535}
                        _ = v48
                        _ = v49
                    }
                    _ = v47
                } else {
                    v50 := GenS7{F0: (v5 % (((v5 | (v5 << ((2) & 63)))) | 1))}
                    v51 := GenS8{F0: 1}
                    _ = v50
                    _ = v51
                }
            }
        case -9223372036854775808:
            v2 += v2
        default:
            for fi52 := 0; fi52 < 3; fi52++ {
                var v53 uint = 2
                v5 = (v5 - v5)
                var v54 uint8 = 255
                _ = fi52
                _ = v53
                _ = v54
            }
        }
    }
    if (-1 > -128) {
        for fi55 := 0; fi55 < 1; fi55++ {
            var v56 int64 = 1
            _ = fi55
            _ = v56
        }
    }
    var v57 float32 = (float32((1.5 * float64(-2.25))) + 1.5)
    if v6 {
        if true {
            var v58 float32 = v57
            v3 = (v3 / ((uint32(v1)) | 1))
            _ = v58
        }
        fmt.Println(v6, v3)
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    switch 127 {
    case 32767:
        fmt.Println(v6, v5)
    case 127:
        v2 = v2
    case 191:
        for fi59 := 0; fi59 < 12; fi59++ {
            if ((v6 || (!v6)) || (0 >= 100)) {
                var v60 uint16 = v1
                _ = v60
            }
            for fi61 := 0; fi61 < 11; fi61++ {
                var v62 int8 = 0
                fmt.Println(v6, v57)
                var v63 int64 = 9223372036854775807
                _ = fi61
                _ = v62
                _ = v63
            }
            fmt.Println(v6, v2, fi59)
            _ = fi59
        }
    }
    var v64 float64 = 1.5
    for fi65 := 0; fi65 < 14; fi65++ {
        v66 := GenS3{F0: 0}
        _ = fi65
        _ = v66
    }
    v67 := GenS9{F0: 42, F1: -1}
    _ = len(v4)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v57:", v57)
    fmt.Println("v6:", v6)
    fmt.Println("v64:", v64)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v57
    _ = v6
    _ = v64
    _ = v67
}
