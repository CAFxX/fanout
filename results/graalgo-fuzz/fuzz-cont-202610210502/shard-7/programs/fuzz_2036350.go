package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint32; F1 int64; F2 uint64 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint8; F1 int8 }
type GenS6 struct { F0 string; F1 int }

func main() {
    var v1 string = "foo bar"
    v2 := []string{}
    v3 := make(map[uint32]int)
    v4 := make(map[int]int)
    var v5 int8 = -1
    var v6 float32 = 1e10
    var v7 int = 127
    for fi8 := 0; fi8 < 8; fi8++ {
        if true {
            v9 := GenS1{F0: 755}
            if true {
                var v10 uint8 = 42
                var v11 float64 = ((-2.25 * -2.25) - 1.5)
                if false {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                }
                _ = v10
                _ = v11
            }
            if false {
                switch (-(fi8)) {
                case -723:
                    for fi12 := 0; fi12 < 20; fi12++ {
                        if false {
                            v7 = (fi8 % ((v7) | 1))
                            v7 -= fi8
                            if false {
                                var v13 int8 = 100
                                _ = v13
                            } else {
                                var v14 float64 = 0.0
                                _ = v14
                            }
                        }
                        switch "foo bar" {
                        case "hello":
                            v6 = 1e10
                        case "":
                            var v15 string = ("hello" + ("" + "foo bar"))
                            _ = v15
                        case "a":
                            if true {
                                v1 = (v1 + ("\n\t" + "hello"))
                                var v16 int = fi12
                                var v17 int = (fi12 & v7)
                                _ = v16
                                _ = v17
                            }
                        }
                        switch v1 {
                        case "foo bar":
                            switch (fi12 >> ((2) & 63)) {
                            case -2147483648:
                                if len(v2) > 0 {
                                    v2[0] = "\n\t"
                                }
                            }
                        case "hello":
                            if (!true) {
                                if len(v2) > 0 {
                                    v2[0] = "a"
                                }
                            }
                        case "":
                            switch fi8 {
                            case 32767:
                                var v18 bool = (1 < -290)
                                _ = v18
                            }
                        default:
                            switch (v1 + "\n\t") {
                            case "":
                                var v19 uint32 = 0
                                _ = v19
                            case "hello":
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                            default:
                                v5 -= (^(v5))
                            }
                        }
                        v6 += 3.14159
                        v6 = v6
                        switch 32768 {
                        case 58:
                            var v20 bool = (true || true)
                            _ = v20
                        case 979:
                            switch "foo bar" {
                            case "foo bar":
                                v21 := GenS2{F0: 0}
                                _ = v21
                            case "hello":
                                v1 = "\n\t"
                            default:
                                switch fi8 {
                                case -2147483648:
                                    switch fi8 {
                                    case 32768:
                                        if len(v2) > 0 {
                                            v2[0] = v1
                                        }
                                    }
                                case 281:
                                    fmt.Println(v6)
                                }
                            }
                        case 32767:
                            v1 = v1
                        default:
                            if len(v2) > 0 {
                                v2[0] = v1
                            }
                        }
                        v22 := v4[fi12]
                        _ = len(v3)
                        v1 = (("\n\t" + (v1 + (v1 + ((v1 + "a") + v1)))) + (((v1 + (v1 + "foo bar")) + v1) + (v1 + v1)))
                        if (((42 > 2) && false) || true) {
                            var v23 uint32 = 1
                            if (true && true) {
                                v24 := GenS3{F0: v23, F1: 9223372036854775807, F2: 18446744073709551615}
                                var v25 uint32 = v23
                                _ = v24
                                _ = v25
                            }
                            v26 := GenS2{F0: 123}
                            _ = v23
                            _ = v26
                        }
                        _ = fi12
                        _ = v22
                    }
                default:
                    switch v1 {
                    case "":
                        for fi27 := 0; fi27 < 2; fi27++ {
                            v28 := GenS1{F0: -47}
                            v4[(fi8 & v7)] = (v7 >> ((18446744073709551615) & 63))
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            _ = fi27
                            _ = v28
                        }
                    case "a":
                        for fi29 := 0; fi29 < 14; fi29++ {
                            v30 := GenS1{F0: 1}
                            v7 = 128
                            _ = fi29
                            _ = v30
                        }
                    }
                }
                v6 = ((v6 / 1e10) / (float32(0.0) - (v6 / 1.5)))
                var v31 uint16 = 358
                _ = v31
            }
            _ = v9
        } else {
            v32 := GenS1{F0: 42}
            for fi33 := 0; fi33 < 3; fi33++ {
                var v34 float32 = 1.5
                switch (v7 % ((fi33) | 1)) {
                case 42:
                    var v35 int = v7
                    _ = v35
                case -9223372036854775808:
                    fmt.Println(v6)
                }
                v36 := GenS1{F0: -374}
                _ = fi33
                _ = v34
                _ = v36
            }
            _ = v32
        }
        for fi37 := 0; fi37 < 11; fi37++ {
            v38 := v4[32768]
            v39 := GenS2{F0: 42}
            var v40 int32 = 2147483647
            _ = fi37
            _ = v38
            _ = v39
            _ = v40
        }
        switch (v7 | (v7 / ((fi8) | 1))) {
        case 0:
            var v41 int64 = 9223372036854775807
            _ = v41
        case 127:
            v5 += 52
        default:
            fmt.Println(fi8, v6)
        }
        _ = fi8
    }
    v6 = (float32(1e10) - v6)
    if true {
        delete(v3, 998)
        switch v7 {
        case 4294967295:
            for fi42 := 0; fi42 < 11; fi42++ {
                _ = len(v4)
                v5 = v5
                _ = fi42
            }
        case 2147483647:
            var v43 bool = false
            _ = v43
        case -117:
            if (!false) {
                v44 := GenS4{F0: (v1 + "foo bar")}
                _ = v44
            }
        }
        switch v1 {
        case "\n\t":
            switch v7 {
            case -2147483648:
                var v45 int32 = 0
                _ = v45
            case -185:
                v5 += 1
            }
        case "a":
            v1 = v1
        case "hello":
            fmt.Println(v5, v7, v6)
        default:
            v7 = (v7 ^ v7)
        }
    }
    v6 = ((v6 * (v6 + (v6 - v6))) + ((float32(1.5) * (v6 - ((v6 - v6) + (v6 / 3.14159)))) * (float32(1e10) + 3.14159)))
    var v46 int = (v7 + v7)
    switch "hello" {
    case "hello":
        if true {
            v47 := GenS3{F0: 746, F1: -1, F2: 130}
            _ = v47
        } else {
            for fi48 := 0; fi48 < 3; fi48++ {
                v5 = 100
                if (((v1 + ("a" + v1)) + (v1 + ("hello" + (v1 + v1)))) <= v1) {
                    for fi49 := 0; fi49 < 9; fi49++ {
                        v2 = append(v2, v1)
                        fmt.Println(v6, fi49, fi48)
                        _ = fi49
                    }
                    var v50 int16 = -32768
                    if true {
                        v46 = -1
                        for fi51 := 0; fi51 < 4; fi51++ {
                            if false {
                                var v52 int32 = 1
                                fmt.Println(fi48, fi48, v46)
                                v5 = (v5 | (v5 % (((v5 % ((int8(v7)) | 1))) | 1)))
                                _ = v52
                            }
                            _ = fi51
                        }
                    } else {
                        var v53 int8 = (v5 >> ((65535) & 63))
                        _ = v53
                    }
                    _ = v50
                }
                _ = fi48
            }
            v46 += (-(v46))
        }
    case "a":
        fmt.Println(v5)
    case "\n\t":
        var v54 uint8 = 0
        _ = v54
    default:
        fmt.Println(v7)
    }
    v55 := GenS4{F0: v1}
    v6 += 0.0
    for fi56 := 0; fi56 < 5; fi56++ {
        if ((false || true) || false) {
            v7 += (v46 + v7)
        }
        if false {
            if true {
                v57 := GenS5{F0: 255, F1: (v5 >> ((607) & 63))}
                switch (fi56 & v46) {
                case 42:
                    fmt.Println(v46, v6)
                case -9223372036854775808:
                    v3[441] = v46
                case 32768:
                    v1 = "a"
                }
                v7 += v7
                _ = v57
            }
            v58 := GenS6{F0: (v1 + v1), F1: (v46 * 9223372036854775807)}
            _ = v58
        }
        if false {
            v59 := GenS5{F0: 2, F1: v5}
            if true {
                fmt.Println(fi56)
                if true {
                    _ = len(v4)
                    var v60 uint = 65535
                    for fi61 := 0; fi61 < 13; fi61++ {
                        var v62 uint64 = 198
                        _ = fi61
                        _ = v62
                    }
                    _ = v60
                }
            } else {
                if (42 != 459) {
                    fmt.Println(fi56)
                } else {
                    v3[557] = fi56
                }
                fmt.Println(v7, v46)
                var v63 string = "foo bar"
                _ = v63
            }
            _ = v59
        } else {
            if len(v2) > 0 {
                v2[0] = v1
            }
            fmt.Println(fi56)
        }
        _ = fi56
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v5
    _ = v55
    _ = v6
    _ = v7
}
