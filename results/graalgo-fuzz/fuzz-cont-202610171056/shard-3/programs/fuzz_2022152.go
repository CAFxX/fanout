package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint; F2 string }
type GenS2 struct { F0 int; F1 uint64; F2 bool }
type GenS3 struct { F0 uint64; F1 int16; F2 bool }
type GenS4 struct { F0 int8 }

func main() {
    v1 := make(map[int]bool)
    var v2 int32 = 737
    v3 := make(map[uint32]bool)
    switch 32767 {
    case 694:
        var v4 uint = 238
        _ = v4
    case -678:
        v2 = (v2 / (((v2 << ((140) & 63))) | 1))
    case 2147483647:
        var v5 int16 = 42
        _ = v5
    }
    v6 := GenS1{F0: 1, F1: 830, F2: (("foo bar" + ("a" + "\n\t")) + ("" + ("\n\t" + "a")))}
    var v7 int16 = 870
    for fi8 := 0; fi8 < 14; fi8++ {
        for fi9 := 0; fi9 < 8; fi9++ {
            var v10 bool = true
            v11 := GenS1{F0: 42, F1: 547, F2: ("foo bar" + "hello")}
            _ = fi9
            _ = v10
            _ = v11
        }
        v12 := v1[(fi8 - (-(fi8)))]
        var v13 uint = 54
        _ = fi8
        _ = v12
        _ = v13
    }
    if true {
        fmt.Println(v7, v7)
    } else {
        var v14 int16 = v7
        _ = v14
    }
    var v15 float32 = 0.0
    var v16 int = -9223372036854775808
    switch ("hello" + "\n\t") {
    case "\n\t":
        var v17 uint8 = 127
        _ = v17
    default:
        v18 := GenS2{F0: v16, F1: 0, F2: true}
        _ = v18
    }
    switch v16 {
    case 128:
        v19 := GenS1{F0: 255, F1: 2, F2: "foo bar"}
        _ = v19
    default:
        var v20 uint16 = 2
        _ = v20
    }
    var v21 uint = 42
    if (0 != 673) {
        if false {
            fmt.Println(v2)
            if (970 >= 1) {
                switch (^(v16)) {
                case 4294967295:
                    var v22 int16 = v7
                    _ = v22
                default:
                    fmt.Println(v16, v16, v16)
                }
            }
        } else {
            switch "foo bar" {
            case "hello":
                for fi23 := 0; fi23 < 11; fi23++ {
                    for fi24 := 0; fi24 < 16; fi24++ {
                        v25 := GenS3{F0: 954, F1: v7, F2: (!(true && (411 > 42)))}
                        v26 := GenS3{F0: 119, F1: (v7 + -419), F2: ("\n\t" <= ("foo bar" + "\n\t"))}
                        _ = fi24
                        _ = v25
                        _ = v26
                    }
                    for fi27 := 0; fi27 < 19; fi27++ {
                        if true {
                            switch ("foo bar" + "a") {
                            case "hello":
                                var v28 float64 = 0.0
                                _ = v28
                            }
                            var v29 string = "\n\t"
                            _ = v29
                        }
                        v16 = (^(fi27))
                        v15 = (float32(((0.1 + (1e10 - (((1.5 + float64((3.14159 / v15))) + float64(v15)) + -2.25))) * 0.1)) * 0.0)
                        _ = fi27
                    }
                    _ = fi23
                }
            }
            v7 -= v7
            v15 = (v15 / v15)
        }
        for fi30 := 0; fi30 < 7; fi30++ {
            v16 -= (v16 ^ fi30)
            var v31 float32 = (v15 * (0.0 * ((((1e10 / float32(((((0.1 / 1.0) - 1e10) * 1e10) + 3.14159))) - (v15 - (float32(float64((v15 / float32(float64((float32((float64(v15) / 1e10)) / 3.14159)))))) + v15))) - 0.0) * ((0.0 - (v15 + 1.5)) + 3.14159))))
            _ = fi30
            _ = v31
        }
    } else {
        var v32 int16 = v7
        var v33 int32 = (v2 | -800)
        for fi34 := 0; fi34 < 4; fi34++ {
            var v35 uint16 = 777
            _ = fi34
            _ = v35
        }
        _ = v32
        _ = v33
    }
    for fi36 := 0; fi36 < 16; fi36++ {
        for fi37 := 0; fi37 < 12; fi37++ {
            v21 = (v21 & 65535)
            _ = fi37
        }
        v38 := GenS4{F0: 0}
        _ = fi36
        _ = v38
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v21
    _ = v3
    _ = v6
    _ = v7
}
