package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16 }
type GenS2 struct { F0 uint; F1 uint64; F2 uint16 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int; F1 uint64; F2 int64 }
type GenS6 struct { F0 int8; F1 int; F2 uint64 }
type GenS7 struct { F0 uint32 }

func main() {
    var v1 int = 128
    v2 := []int{}
    var v3 uint = 691
    var v4 int64 = 0
    var v5 string = "foo bar"
    v6 := make(map[int]int)
    if false {
        fmt.Println(v1, v4)
        for fi7 := 0; fi7 < 20; fi7++ {
            var v8 uint = v3
            _ = fi7
            _ = v8
        }
        v9 := GenS1{F0: -32768, F1: 32767}
        _ = v9
    } else {
        var v10 uint32 = 2147483647
        _ = v10
    }
    v11 := GenS1{F0: -285, F1: 331}
    v12 := GenS2{F0: (v3 | (v3 ^ v3)), F1: 4294967295, F2: 32767}
    var v13 int16 = 761
    v14 := GenS1{F0: (v13 / (((^(v13))) | 1)), F1: (v13 ^ 1000)}
    v15 := GenS2{F0: v3, F1: 18446744073709551615, F2: 32767}
    v16 := GenS1{F0: -533, F1: (v13 << (((v3 & v3)) & 63))}
    if (v13 >= (v13 << (((v3 << ((v3) & 63))) & 63))) {
        var v17 uint8 = 56
        switch (v1 << (((^(v3))) & 63)) {
        case -9223372036854775808:
            v18 := GenS1{F0: 32767, F1: v13}
            _ = v18
        }
        if true {
            var v19 float32 = 1e10
            for fi20 := 0; fi20 < 5; fi20++ {
                switch (v5 + v5) {
                case "hello":
                    var v21 int64 = v4
                    _ = v21
                case "":
                    if true {
                        v22 := GenS1{F0: v13, F1: (v13 - (v13 & (v13 >> ((v3) & 63))))}
                        _ = v22
                    } else {
                        var v23 uint32 = 2
                        _ = v23
                    }
                default:
                    v4 -= (v4 / (((v4 / ((v4) | 1))) | 1))
                }
                fmt.Println(v17)
                switch "foo bar" {
                case "\n\t":
                    v24 := GenS2{F0: v3, F1: 0, F2: 42}
                    _ = v24
                case "a":
                    var v25 float64 = 3.14159
                    _ = v25
                case "foo bar":
                    v17 += (v17 & (v17 - v17))
                }
                _ = fi20
            }
            _ = v19
        }
        _ = v17
    }
    v13 += 481
    var v26 uint = 4294967295
    v27 := GenS1{F0: 32767, F1: (v13 / (((v13 / ((v13) | 1))) | 1))}
    for fi28 := 0; fi28 < 12; fi28++ {
        if (false && (!(((v4 % (((v4 % ((1) | 1))) | 1)) == (-(v4))) || false))) {
            switch (v5 + v5) {
            case "\n\t":
                v3 = 833
            default:
                switch (v1 / (((-(v1))) | 1)) {
                case 2147483647:
                    v29 := GenS3{F0: 65535}
                    _ = v29
                case 0:
                    var v30 int32 = -101
                    _ = v30
                case 32767:
                    var v31 uint16 = 1
                    _ = v31
                }
            }
            var v32 int = fi28
            if (v13 == -575) {
                fmt.Println(fi28, fi28)
                _ = len(v6)
                switch (-(fi28)) {
                case -1:
                    v4 = v4
                }
            } else {
                if (!(2 <= 4294967295)) {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    fmt.Println(fi28, v26)
                } else {
                    switch "" {
                    case "":
                        if false {
                            for fi33 := 0; fi33 < 1; fi33++ {
                                var v34 float32 = (float32((1.5 - float64((-2.25 / -2.25)))) + 0.0)
                                switch v5 {
                                case "a":
                                    v34 -= v34
                                case "hello":
                                    if ((444 != 42) && (uint32(v3) > 2147483647)) {
                                        var v35 uint = (v26 / (((v3 ^ 4294967295)) | 1))
                                        _ = v35
                                    } else {
                                        v5 = "a"
                                        v1 -= -647
                                        var v36 int16 = v13
                                        _ = v36
                                    }
                                }
                                var v37 string = ((v5 + (v5 + (v5 + v5))) + v5)
                                _ = fi33
                                _ = v34
                                _ = v37
                            }
                            for fi38 := 0; fi38 < 12; fi38++ {
                                var v39 uint = 370
                                _ = fi38
                                _ = v39
                            }
                            for fi40 := 0; fi40 < 18; fi40++ {
                                var v41 float32 = 0.1
                                v32 = (-(v32))
                                v2 = append(v2, v1)
                                _ = fi40
                                _ = v41
                            }
                        }
                    case "foo bar":
                        if (65535 > 2) {
                            v42 := GenS4{F0: -694}
                            v13 = (v13 >> ((v3) & 63))
                            _ = v42
                        }
                    case "hello":
                        for fi43 := 0; fi43 < 13; fi43++ {
                            if true {
                                v4 = 9223372036854775807
                            } else {
                                v5 = v5
                            }
                            if (0 > 143) {
                                v4 = (^(v4))
                            } else {
                                var v44 int = v32
                                switch (v5 + "foo bar") {
                                case "":
                                    v45 := GenS5{F0: v44, F1: 4, F2: (v4 + 0)}
                                    _ = v45
                                default:
                                    v3 -= (^(v26))
                                }
                                _ = v44
                            }
                            _ = fi43
                        }
                    default:
                        for fi46 := 0; fi46 < 6; fi46++ {
                            switch 32768 {
                            case 32767:
                                if len(v2) > 0 {
                                    v2[0] = int(v13)
                                }
                                v47 := v6[-463]
                                v4 = v4
                                _ = v47
                            }
                            fmt.Println(v32, fi46, v4)
                            v26 -= (v3 ^ (v3 << ((18446744073709551615) & 63)))
                            _ = fi46
                        }
                    }
                    v48 := GenS6{F0: -128, F1: (v32 << ((v3) & 63)), F2: 254}
                    for fi49 := 0; fi49 < 10; fi49++ {
                        v3 = 4294967295
                        var v50 int = v1
                        _ = fi49
                        _ = v50
                    }
                    _ = v48
                }
                for fi51 := 0; fi51 < 9; fi51++ {
                    if true {
                        v52 := GenS7{F0: 2}
                        _ = v52
                    }
                    var v53 float32 = (0.0 + -2.25)
                    _ = fi51
                    _ = v53
                }
            }
            _ = v32
        }
        fmt.Println(v13, v3, v5)
        _ = len(v6)
        _ = fi28
    }
    v54 := GenS6{F0: 42, F1: (v1 % (((v1 >> ((v26) & 63))) | 1)), F2: 118}
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v54
    _ = v6
}
