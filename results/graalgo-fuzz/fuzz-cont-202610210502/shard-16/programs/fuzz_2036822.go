package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool; F2 int16 }
type GenS2 struct { F0 int16; F1 uint16 }
type GenS3 struct { F0 int32 }

func main() {
    v1 := []uint32{}
    var v2 int = -9223372036854775808
    var v3 uint16 = 32767
    if true {
        var v4 int8 = 1
        _ = v4
    } else {
        switch "foo bar" {
        case "foo bar":
            v1 = append(v1, 72)
        case "a":
            for fi5 := 0; fi5 < 6; fi5++ {
                var v6 float64 = (-2.25 / 1.0)
                _ = fi5
                _ = v6
            }
        }
    }
    v2 -= (v2 ^ (v2 + (v2 % (((v2 ^ (v2 & (v2 / (((v2 << ((921) & 63))) | 1))))) | 1))))
    if ((^(v3)) > v3) {
        if (!false) {
            _ = len(v1)
            v3 = (v3 - 1)
        } else {
            fmt.Println(v3, v2)
        }
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        var v7 uint = 183
        _ = v7
    } else {
        var v8 uint8 = 32
        if (!false) {
            v3 = (v3 / ((v3) | 1))
        } else {
            fmt.Println(v8, v8)
            if false {
                if (!false) {
                    switch "" {
                    case "":
                        var v9 uint16 = (v3 ^ 213)
                        _ = v9
                    case "a":
                        v8 = v8
                    }
                    v1 = append(v1, 0)
                }
                v10 := GenS1{F0: v3, F1: (1 <= 1), F2: 32767}
                fmt.Println(v8)
                _ = v10
            }
        }
        switch "foo bar" {
        case "":
            switch v2 {
            case 181:
                if (0.1 < 1e10) {
                    if (!true) {
                        v3 += (v3 / (((v3 << ((4294967295) & 63))) | 1))
                        var v11 string = "hello"
                        _ = v11
                    }
                }
            case -1:
                if (!true) {
                    if len(v1) > 0 {
                        v1[0] = 2
                    }
                } else {
                    for fi12 := 0; fi12 < 10; fi12++ {
                        var v13 int8 = -64
                        switch fi12 {
                        case 1:
                            v3 = (^(v3))
                        case 32768:
                            switch "\n\t" {
                            case "\n\t":
                                v8 += 255
                            }
                        default:
                            fmt.Println(v3, v13)
                        }
                        _ = fi12
                        _ = v13
                    }
                }
            }
        case "a":
            var v14 int16 = 42
            _ = v14
        case "hello":
            for fi15 := 0; fi15 < 18; fi15++ {
                v16 := GenS2{F0: 0, F1: v3}
                v1 = append(v1, 1)
                var v17 int64 = 0
                _ = fi15
                _ = v16
                _ = v17
            }
        default:
            v3 -= 345
        }
        _ = v8
    }
    var v18 uint8 = 2
    fmt.Println(v18, v2, v18)
    for fi19 := 0; fi19 < 1; fi19++ {
        var v20 int32 = 1
        _ = fi19
        _ = v20
    }
    var v21 int16 = -757
    _ = len(v1)
    var v22 uint64 = 2
    v18 += (v18 | 128)
    fmt.Println(v18)
    var v23 float32 = ((0.0 * 1.5) * (0.0 + (0.1 / 0.1)))
    v24 := GenS3{F0: -201}
    v2 += (v2 + v2)
    for fi25 := 0; fi25 < 4; fi25++ {
        for fi26 := 0; fi26 < 1; fi26++ {
            v21 = v21
            _ = fi26
        }
        if (!((v22 | (v22 % (((v22 / (((v22 ^ (v22 | v22))) | 1))) | 1))) <= v22)) {
            switch "a" {
            case "foo bar":
                if ("" <= "foo bar") {
                    if ((!true) && false) {
                        v3 -= 42
                    } else {
                        v27 := GenS2{F0: v21, F1: v3}
                        var v28 int = (fi25 & (v2 * int(v18)))
                        _ = v27
                        _ = v28
                    }
                    v22 += (v22 * (v22 - (v22 + v22)))
                }
            case "a":
                if len(v1) > 0 {
                    v1[0] = 206
                }
            }
        }
        v23 = 1e10
        _ = fi25
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v18
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
}
