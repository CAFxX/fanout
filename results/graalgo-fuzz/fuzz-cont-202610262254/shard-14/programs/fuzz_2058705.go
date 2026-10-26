package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int64 }

func main() {
    v1 := []uint32{2147483648, 4294967295, 883, 150}
    v2 := make(map[string]string)
    var v3 int16 = 1000
    fmt.Println(v3)
    for fi4 := 0; fi4 < 8; fi4++ {
        _ = len(v2)
        _ = fi4
    }
    switch 32768 {
    case 9223372036854775807:
        v3 += 1
    case 0:
        var v5 uint8 = 0
        _ = v5
    case -1:
        v6 := GenS1{F0: 18446744073709551615, F1: -9223372036854775808}
        _ = v6
    default:
        var v7 int = -1
        _ = v7
    }
    switch 1 {
    case 812:
        v3 += v3
    case 9223372036854775807:
        var v8 bool = false
        _ = v8
    default:
        var v9 uint16 = 209
        _ = v9
    }
    switch 128 {
    case 127:
        if (-2147483648 <= 1) {
            v10 := GenS1{F0: 65535, F1: 9223372036854775807}
            fmt.Println(v3)
            _ = v10
        } else {
            switch (("" + "") + "a") {
            case "a":
                v11 := GenS1{F0: 44, F1: 1}
                _ = v11
            }
            _ = len(v2)
        }
    case 128:
        for fi12 := 0; fi12 < 19; fi12++ {
            switch 874 {
            case 2147483647:
                if ((!(true && (false || true))) && true) {
                    var v13 int64 = -1
                    _ = v13
                } else {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    fmt.Println(v3, v3)
                    v14 := GenS1{F0: 85, F1: 744}
                    _ = v14
                }
            default:
                v3 = (v3 / ((1) | 1))
            }
            _ = fi12
        }
    default:
        for fi15 := 0; fi15 < 4; fi15++ {
            if (!false) {
                switch "hello" {
                case "foo bar":
                    if ((1e10 - 1e10) >= 1.5) {
                        v3 = v3
                    }
                case "\n\t":
                    for fi16 := 0; fi16 < 16; fi16++ {
                        fmt.Println(v3, fi15, fi15)
                        v17 := GenS2{F0: 606}
                        if false {
                            var v18 int = fi15
                            var v19 uint64 = 614
                            v20 := GenS3{F0: 688}
                            if true {
                                v3 -= v3
                                switch "hello" {
                                case "foo bar":
                                    var v21 uint = 42
                                    _ = v21
                                case "":
                                    v3 += v3
                                case "a":
                                    var v22 int16 = (v3 | v3)
                                    _ = v22
                                }
                            }
                            _ = v18
                            _ = v19
                            _ = v20
                        }
                        _ = fi16
                        _ = v17
                    }
                default:
                    v3 = v3
                }
            }
            var v23 uint16 = 3
            _ = fi15
            _ = v23
        }
    }
    if false {
        for fi24 := 0; fi24 < 4; fi24++ {
            if true {
                var v25 string = "a"
                _ = v25
            } else {
                v3 = (v3 << ((2) & 63))
                v3 = -655
                fmt.Println(fi24, v3)
            }
            _ = len(v2)
            v3 -= v3
            _ = fi24
        }
        var v26 int16 = v3
        _ = v26
    } else {
        var v27 float64 = ((3.14159 / 3.14159) + (0.0 * 1e10))
        v27 = ((v27 / 1.0) * (v27 * (v27 / ((-2.25 + (float64((0.1 + 3.14159)) + v27)) + v27))))
        v28 := v2[("hello" + "a")]
        _ = v27
        _ = v28
    }
    fmt.Println(v3)
    v3 = v3
    fmt.Println(v3)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v3
}
