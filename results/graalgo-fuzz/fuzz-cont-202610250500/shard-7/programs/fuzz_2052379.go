package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64; F2 uint8 }
type GenS2 struct { F0 bool; F1 int8; F2 uint32 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 bool; F1 int32 }

func main() {
    var v1 uint32 = 738
    var v2 float32 = float32(1e10)
    var v3 uint32 = v1
    v4 := []int{-48}
    fmt.Println(v3, v2)
    if len(v4) > 0 {
        v4[0] = 9223372036854775807
    }
    var v5 uint32 = v1
    var v6 uint32 = (^(v1))
    switch "" {
    case "a":
        if (!false) {
            v6 = (v1 | v1)
            for fi7 := 0; fi7 < 2; fi7++ {
                switch "a" {
                case "\n\t":
                    var v8 bool = (-559 < 1)
                    _ = v8
                }
                _ = fi7
            }
            fmt.Println(v2)
        }
    case "hello":
        switch 127 {
        case -822:
            var v9 float32 = ((float32(3.14159) + v2) + 0.1)
            _ = v9
        default:
            if ((217 > 2) && true) {
                var v10 int32 = 1
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                _ = v10
            } else {
                fmt.Println(v2, v6, v6)
                var v11 int = 9223372036854775807
                fmt.Println(v1, v5, v3)
                _ = v11
            }
        }
    case "\n\t":
        switch "\n\t" {
        case "\n\t":
            v12 := GenS1{F0: 391, F1: 682, F2: 235}
            _ = v12
        case "a":
            v3 -= (v5 - v3)
        }
    default:
        v6 -= (^(v3))
    }
    var v13 uint = 295
    for fi14 := 0; fi14 < 16; fi14++ {
        fmt.Println(v1, v5)
        _ = fi14
    }
    if true {
        var v15 string = ((("foo bar" + "") + "foo bar") + "foo bar")
        switch ((((("a" + "\n\t") + (v15 + "\n\t")) + ((((v15 + v15) + v15) + "\n\t") + v15)) + v15) + (v15 + (("\n\t" + "a") + "hello"))) {
        case "hello":
            fmt.Println(v5, v15, v6)
        case "foo bar":
            if false {
                if (65535 >= 665) {
                    v16 := GenS2{F0: (true && false), F1: 42, F2: v5}
                    for fi17 := 0; fi17 < 15; fi17++ {
                        var v18 int64 = 0
                        for fi19 := 0; fi19 < 14; fi19++ {
                            var v20 float64 = ((-2.25 / 1.0) / 1.5)
                            _ = fi19
                            _ = v20
                        }
                        for fi21 := 0; fi21 < 18; fi21++ {
                            var v22 uint = (v13 / ((v13) | 1))
                            var v23 uint8 = 68
                            _ = fi21
                            _ = v22
                            _ = v23
                        }
                        _ = fi17
                        _ = v18
                    }
                    _ = v16
                } else {
                    if false {
                        for fi24 := 0; fi24 < 11; fi24++ {
                            v25 := GenS1{F0: 42, F1: -166, F2: 127}
                            var v26 uint64 = 120
                            fmt.Println(v5)
                            _ = fi24
                            _ = v25
                            _ = v26
                        }
                        var v27 int8 = 42
                        _ = v27
                    }
                }
            }
        case "\n\t":
            v3 -= v6
        }
        _ = v15
    }
    var v28 float64 = 3.14159
    v6 = (v6 << (((v13 % ((195) | 1))) & 63))
    if len(v4) > 0 {
        v4[0] = 9223372036854775807
    }
    v2 = -2.25
    for fi29 := 0; fi29 < 6; fi29++ {
        v28 += 0.1
        _ = fi29
    }
    v13 = uint(v3)
    if true {
        v30 := GenS3{F0: true}
        v31 := GenS4{F0: false, F1: -136}
        _ = v30
        _ = v31
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
