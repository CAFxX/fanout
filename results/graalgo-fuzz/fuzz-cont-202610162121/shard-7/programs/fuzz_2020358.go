package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64; F2 uint }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int8; F1 uint16; F2 uint16 }

func main() {
    var v1 uint8 = 2
    v2 := make(map[string]bool)
    v3 := []uint32{1, 2}
    v4 := GenS1{F0: "foo bar", F1: 967, F2: 244}
    switch "\n\t" {
    case "hello":
        if false {
            _ = len(v3)
        } else {
            v1 -= 0
        }
    case "\n\t":
        var v5 int16 = 42
        _ = v5
    case "":
        for fi6 := 0; fi6 < 16; fi6++ {
            var v7 uint64 = 604
            for fi8 := 0; fi8 < 2; fi8++ {
                var v9 uint8 = 193
                v1 += (^(v9))
                if true {
                    var v10 int = 1
                    var v11 float64 = ((1.5 / 0.1) + 3.14159)
                    _ = v10
                    _ = v11
                }
                _ = fi8
                _ = v9
            }
            _ = fi6
            _ = v7
        }
    default:
        _ = len(v3)
    }
    v12 := GenS2{F0: true}
    _ = len(v3)
    if true {
        v13 := GenS3{F0: -11, F1: 32767, F2: uint16(v1)}
        v14 := GenS3{F0: 0, F1: 0, F2: 65535}
        _ = v13
        _ = v14
    }
    var v15 uint16 = 0
    var v16 float64 = 0.0
    if true {
        if (v15 != (^(v15))) {
            if (!false) {
                v17 := GenS2{F0: (false && (!false))}
                _ = v17
            } else {
                var v18 float64 = v16
                v18 = ((float64(float32(-2.25)) + v18) - (-2.25 - 0.1))
                _ = v18
            }
            if len(v3) > 0 {
                v3[0] = 2
            }
        } else {
            v19 := GenS3{F0: 100, F1: (v15 - (v15 / ((v15) | 1))), F2: v15}
            v15 -= v15
            _ = v19
        }
        switch -9223372036854775808 {
        case 4294967295:
            switch ("" + "foo bar") {
            case "foo bar":
                var v20 float32 = (3.14159 - 0.0)
                _ = v20
            case "\n\t":
                if false {
                    v15 = 451
                    for fi21 := 0; fi21 < 7; fi21++ {
                        var v22 float32 = float32(v16)
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        _ = fi21
                        _ = v22
                    }
                } else {
                    var v23 uint64 = 2
                    _ = v23
                }
            }
        case 0:
            _ = len(v3)
        case -2147483648:
            v1 = v1
        }
        v16 = v16
    }
    switch (("" + "hello") + "\n\t") {
    case "\n\t":
        fmt.Println(v16, v15, v15)
    case "a":
        fmt.Println(v16, v15, v1)
    case "":
        for fi24 := 0; fi24 < 19; fi24++ {
            switch (("hello" + "a") + "a") {
            case "a":
                fmt.Println(v16)
            case "":
                switch "a" {
                case "foo bar":
                    v15 = v15
                case "hello":
                    for fi25 := 0; fi25 < 13; fi25++ {
                        v15 -= v15
                        v15 += (v15 % (((v15 | (v15 + v15))) | 1))
                        var v26 uint = 826
                        _ = fi25
                        _ = v26
                    }
                case "":
                    if len(v3) > 0 {
                        v3[0] = 328
                    }
                default:
                    fmt.Println(v16)
                }
            }
            for fi27 := 0; fi27 < 17; fi27++ {
                var v28 float32 = 1e10
                _ = fi27
                _ = v28
            }
            for fi29 := 0; fi29 < 8; fi29++ {
                var v30 int8 = 118
                _ = fi29
                _ = v30
            }
            _ = fi24
        }
    }
    v31 := GenS3{F0: 0, F1: v15, F2: (v15 - uint16(v1))}
    v16 = (1e10 - ((1.5 / v16) * v16))
    _ = len(v2)
    v15 = v15
    v32 := GenS3{F0: -128, F1: v15, F2: (v15 & 545)}
    _ = len(v3)
    v1 -= 0
    if (4294967295 > 1) {
        v15 = (v15 << ((18446744073709551615) & 63))
    } else {
        v1 += v1
        var v33 float32 = (-2.25 * 1.5)
        _ = v33
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v12
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v4
}
