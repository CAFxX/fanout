package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64; F2 int64 }
type GenS2 struct { F0 int32; F1 int64 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint8 }

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[int]uint32)
    var v3 int8 = 45
    var v4 int = -640
    var v5 int32 = -963
    v6 := []int{v4, v4, v4}
    for fi7 := 0; fi7 < 12; fi7++ {
        v4 = (v4 & (v4 ^ fi7))
        var v8 uint64 = 4294967295
        for fi9 := 0; fi9 < 9; fi9++ {
            v3 -= (v3 ^ 42)
            v5 = 31
            _ = fi9
        }
        _ = fi7
        _ = v8
    }
    fmt.Println(v4, v5, v3)
    v4 = (v4 ^ v4)
    for fi10 := 0; fi10 < 13; fi10++ {
        fmt.Println(fi10)
        _ = fi10
    }
    v1[("foo bar" + "foo bar")] = (!false)
    if true {
        var v11 uint16 = 1
        v12 := GenS1{F0: (v4 ^ (v4 >> ((87) & 63))), F1: 9223372036854775807, F2: -1}
        _ = v11
        _ = v12
    } else {
        if (4294967295 < 65535) {
            if false {
                v13 := GenS2{F0: (^(v5)), F1: 42}
                _ = v13
            } else {
                fmt.Println(v5, v3, v5)
            }
            if len(v6) > 0 {
                v6[0] = v4
            }
        } else {
            v4 -= (-(v4))
            v4 = v4
            fmt.Println(v3, v4)
        }
        var v14 int16 = -704
        if (false || false) {
            if true {
                switch (v4 ^ -2147483648) {
                case -1:
                    v15 := GenS2{F0: (v5 / ((v5) | 1)), F1: 369}
                    _ = v15
                default:
                    v4 = (v4 << ((4294967295) & 63))
                }
            }
            fmt.Println(v14)
            for fi16 := 0; fi16 < 5; fi16++ {
                if (18446744073709551615 != 65535) {
                    fmt.Println(v3)
                    if true {
                        v17 := GenS3{F0: 2}
                        var v18 float32 = -2.25
                        if true {
                            var v19 int64 = -751
                            switch v4 {
                            case 4294967295:
                                v20 := GenS2{F0: 42, F1: -224}
                                _ = v20
                            case -2147483648:
                                var v21 float32 = (0.1 * v18)
                                _ = v21
                            case 0:
                                fmt.Println(v5)
                            default:
                                v3 -= v3
                            }
                            for fi22 := 0; fi22 < 6; fi22++ {
                                if len(v6) > 0 {
                                    v6[0] = fi16
                                }
                                _ = fi22
                            }
                            _ = v19
                        }
                        _ = v17
                        _ = v18
                    } else {
                        var v23 uint = 587
                        _ = v23
                    }
                    var v24 int16 = v14
                    _ = v24
                }
                var v25 int8 = (v3 & v3)
                var v26 int8 = 42
                _ = fi16
                _ = v25
                _ = v26
            }
        }
        _ = v14
    }
    v27 := GenS2{F0: (v5 ^ (v5 / ((1) | 1))), F1: 86}
    delete(v2, v4)
    v3 += (v3 * v3)
    switch v4 {
    case 4294967295:
        fmt.Println(v3)
    case -205:
        var v28 uint = 65535
        _ = v28
    case 32767:
        v5 = 1
    default:
        if (v5 < -2147483648) {
            switch (v4 - 127) {
            case 533:
                for fi29 := 0; fi29 < 15; fi29++ {
                    switch "" {
                    case "a":
                        var v30 int16 = 42
                        _ = v30
                    case "":
                        v31 := GenS3{F0: 1}
                        _ = v31
                    }
                    fmt.Println(fi29, v5, v3)
                    _ = fi29
                }
            }
            var v32 int64 = 9223372036854775807
            _ = v32
        }
    }
    if true {
        _ = len(v6)
        v4 -= (v4 * (v4 + v4))
        for fi33 := 0; fi33 < 6; fi33++ {
            switch "\n\t" {
            case "":
                if true {
                    v5 = -1
                } else {
                    fmt.Println(v4, v4)
                    v4 = (v4 | (fi33 / ((fi33) | 1)))
                }
            case "a":
                var v34 float32 = 1e10
                _ = v34
            default:
                if (902 >= -1) {
                    switch "a" {
                    case "\n\t":
                        v35 := GenS4{F0: 255}
                        _ = v35
                    }
                }
            }
            var v36 float64 = 1.5
            var v37 string = ("a" + "")
            _ = fi33
            _ = v36
            _ = v37
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
