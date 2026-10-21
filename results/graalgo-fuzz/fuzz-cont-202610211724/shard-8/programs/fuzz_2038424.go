package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64; F2 int16 }
type GenS2 struct { F0 int32; F1 int32 }
type GenS3 struct { F0 int16; F1 uint16 }
type GenS4 struct { F0 int64; F1 uint; F2 uint32 }

func main() {
    v1 := []int{}
    v2 := []string{("\n\t" + ("foo bar" + ("a" + "foo bar"))), "hello", ("\n\t" + ("a" + "foo bar")), ((((("" + (((("\n\t" + (("a" + ("foo bar" + "hello")) + "hello")) + "foo bar") + "") + ("foo bar" + "foo bar"))) + (("foo bar" + ("hello" + ("a" + "\n\t"))) + "\n\t")) + "hello") + "") + "hello")}
    var v3 int16 = 1000
    v4 := []uint32{1, 2147483647}
    if (18446744073709551615 <= 1) {
        v3 += 293
    } else {
        if (!(0 > 118)) {
            v3 += -522
            fmt.Println(v3)
        }
        for fi5 := 0; fi5 < 6; fi5++ {
            for fi6 := 0; fi6 < 8; fi6++ {
                v3 = v3
                fmt.Println(v3, v3)
                v3 = v3
                _ = fi6
            }
            _ = fi5
        }
    }
    switch "" {
    case "foo bar":
        var v7 uint16 = 65535
        _ = v7
    case "":
        v3 += (v3 ^ -87)
    case "a":
        var v8 int16 = (v3 % ((1000) | 1))
        _ = v8
    default:
        if false {
            var v9 uint16 = 42
            v3 += 42
            _ = v9
        }
    }
    if (true || (false && (42 == 0))) {
        v3 = -32768
        _ = len(v1)
        v3 += 32767
    }
    v10 := GenS1{F0: 607, F1: -774, F2: (v3 >> ((1) & 63))}
    for fi11 := 0; fi11 < 9; fi11++ {
        if true {
            var v12 uint64 = 0
            _ = v12
        } else {
            switch "a" {
            case "":
                v13 := GenS2{F0: -970, F1: -1}
                _ = v13
            case "foo bar":
                var v14 uint = 2
                _ = v14
            }
            for fi15 := 0; fi15 < 13; fi15++ {
                v3 += v3
                switch fi11 {
                case 2147483647:
                    var v16 float32 = ((1.5 * 1e10) - 1.5)
                    _ = v16
                case 598:
                    v17 := GenS1{F0: 393, F1: 9223372036854775807, F2: (v3 % (((v3 % ((575) | 1))) | 1))}
                    _ = v17
                default:
                    _ = len(v4)
                }
                _ = fi15
            }
            if (42 > 1) {
                v3 += (v3 | 1)
                switch "foo bar" {
                case "hello":
                    v3 = (v3 - (v3 | (v3 % ((1) | 1))))
                }
                v3 = (^(v3))
            } else {
                for fi18 := 0; fi18 < 12; fi18++ {
                    v3 += (v3 | v3)
                    if (-85 == -128) {
                        if true {
                            v1 = append(v1, (fi18 ^ fi11))
                            fmt.Println(fi18, v3, v3)
                            if (!true) {
                                var v19 uint = 187
                                _ = v19
                            } else {
                                v3 += (v3 + (v3 | (v3 % (((v3 << ((596) & 63))) | 1))))
                                fmt.Println(fi18)
                            }
                            v20 := GenS3{F0: (v3 - 679), F1: 494}
                            _ = v20
                        }
                    }
                    fmt.Println(fi18)
                    var v21 string = "\n\t"
                    _ = fi18
                    _ = v21
                }
            }
        }
        var v22 uint64 = 2
        if false {
            for fi23 := 0; fi23 < 20; fi23++ {
                if true {
                    v22 -= (v22 + 18446744073709551615)
                    v3 = 1000
                    var v24 int8 = 100
                    _ = v24
                }
                _ = fi23
            }
        } else {
            _ = len(v4)
            v22 += 244
        }
        _ = fi11
        _ = v22
    }
    v3 -= (v3 | v3)
    if (0 != 466) {
        v25 := GenS4{F0: -84, F1: 0, F2: 42}
        v3 = (^(v3))
        var v26 int64 = -9223372036854775808
        _ = v25
        _ = v26
    } else {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    var v27 int64 = -313
    v1 = append(v1, -2147483648)
    _ = len(v4)
    if (0 >= 18446744073709551615) {
        var v28 uint8 = 216
        _ = v28
    }
    v27 = int64(v3)
    v3 = -1
    switch -9223372036854775808 {
    case 2147483647:
        v3 -= v3
    default:
        v27 = (v27 % ((v27) | 1))
    }
    for fi29 := 0; fi29 < 8; fi29++ {
        v30 := GenS2{F0: 2147483647, F1: 0}
        if false {
            var v31 int32 = -872
            fmt.Println(v27, v31)
            fmt.Println(v31, v3, v27)
            _ = v31
        }
        _ = fi29
        _ = v30
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v10
    _ = v2
    _ = v27
    _ = v3
    _ = v4
}
