package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32 }
type GenS2 struct { F0 int16; F1 uint; F2 uint16 }
type GenS3 struct { F0 uint64; F1 uint8; F2 int16 }
type GenS4 struct { F0 uint8; F1 int }
type GenS5 struct { F0 string }

func main() {
    var v1 uint = 4294967295
    v2 := make(map[uint32]int)
    v3 := []string{("\n\t" + "\n\t")}
    var v4 uint = v1
    if false {
        v5 := GenS1{F0: 2, F1: 4294967295}
        var v6 int = -2147483648
        _ = v5
        _ = v6
    }
    var v7 float32 = -2.25
    var v8 int32 = -2147483648
    if false {
        fmt.Println(v4)
        var v9 string = ""
        _ = v9
    }
    fmt.Println(v8, v4)
    v1 = (v1 >> (((^(v1))) & 63))
    if ((0.1 + v7) != v7) {
        v10 := GenS1{F0: 18446744073709551615, F1: 4294967295}
        _ = v10
    }
    switch "\n\t" {
    case "\n\t":
        var v11 int64 = -529
        _ = v11
    case "a":
        if (!(!false)) {
            v12 := v2[567]
            _ = v12
        }
    default:
        fmt.Println(v1, v1)
    }
    v13 := v2[42]
    for fi14 := 0; fi14 < 19; fi14++ {
        for fi15 := 0; fi15 < 19; fi15++ {
            if false {
                v4 = (v1 % ((v4) | 1))
            } else {
                v1 += v4
                _ = len(v2)
                var v16 uint16 = 0
                _ = v16
            }
            v1 -= 0
            v7 -= v7
            _ = fi15
        }
        v17 := GenS2{F0: -1, F1: v1, F2: 0}
        v7 = 1e10
        _ = fi14
        _ = v17
    }
    if true {
        var v18 bool = ((!false) || (!(127 == 127)))
        fmt.Println(v1, v18)
        _ = v18
    } else {
        fmt.Println(v1)
        if false {
            switch ("foo bar" + ("" + "hello")) {
            case "foo bar":
                v1 -= (^(v1))
            }
            var v19 float64 = 1.5
            v20 := GenS3{F0: 365, F1: 42, F2: 341}
            _ = v19
            _ = v20
        }
    }
    for fi21 := 0; fi21 < 5; fi21++ {
        var v22 uint32 = 728
        if false {
            for fi23 := 0; fi23 < 12; fi23++ {
                v13 = int(v8)
                _ = fi23
            }
            for fi24 := 0; fi24 < 1; fi24++ {
                if false {
                    v25 := GenS4{F0: 0, F1: (v13 ^ (fi21 - 0))}
                    v26 := GenS1{F0: v4, F1: v22}
                    _ = v25
                    _ = v26
                }
                v2[(v22 / ((1) | 1))] = fi21
                _ = fi24
            }
            var v27 int64 = 77
            _ = v27
        } else {
            v22 -= (v22 / (((^(v22))) | 1))
            var v28 uint64 = 2
            _ = v28
        }
        _ = fi21
        _ = v22
    }
    if ((!true) || true) {
        v29 := GenS3{F0: 0, F1: 2, F2: -135}
        if false {
            v1 = (v4 & 1)
            fmt.Println(v13, v1)
        } else {
            for fi30 := 0; fi30 < 15; fi30++ {
                var v31 string = "foo bar"
                v4 += (v1 | (^(v4)))
                _ = fi30
                _ = v31
            }
            switch "a" {
            case "hello":
                v32 := GenS5{F0: ("foo bar" + "foo bar")}
                _ = v32
            default:
                for fi33 := 0; fi33 < 18; fi33++ {
                    if ((128 >= 127) || false) {
                        v4 = 4294967295
                        if false {
                            v8 -= v8
                        }
                        fmt.Println(v1, v1, v4)
                    }
                    var v34 uint64 = 380
                    _ = fi33
                    _ = v34
                }
            }
            if false {
                switch "hello" {
                case "foo bar":
                    if (-33 != 42) {
                        if false {
                            v35 := GenS5{F0: ""}
                            _ = v35
                        } else {
                            v8 -= (v8 * (v8 * v8))
                            v13 -= v13
                        }
                    }
                }
                for fi36 := 0; fi36 < 7; fi36++ {
                    v8 = (v8 >> (((v1 - 4294967295)) & 63))
                    fmt.Println(v1, fi36, v7)
                    switch -9223372036854775808 {
                    case 4294967295:
                        v37 := GenS3{F0: 326, F1: 127, F2: 1000}
                        _ = v37
                    }
                    _ = fi36
                }
            }
        }
        if (42 < 0) {
            v4 += v4
        }
        _ = v29
    }
    var v38 uint = (v1 * 18446744073709551615)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v7
    _ = v8
}
