package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8 }
type GenS2 struct { F0 bool; F1 int8; F2 uint16 }

func main() {
    var v1 uint32 = 4294967295
    v2 := make(map[uint32]bool)
    var v3 float64 = (((1e10 / (((float64(3.14159) + 0.1) * (float64(1.5) + (0.0 - 3.14159))) / 1e10)) - 1.5) / 1.0)
    v4 := make(map[uint32]uint32)
    var v5 uint8 = 42
    v6 := v4[(v1 << ((18446744073709551615) & 63))]
    v7 := []int{-1, 365, -66, 2147483647}
    v8 := GenS1{F0: 255, F1: -65}
    var v9 int = 1
    v9 = v9
    switch (("hello" + "foo bar") + "foo bar") {
    case "hello":
        var v10 int8 = -128
        _ = v10
    case "a":
        var v11 uint64 = 679
        _ = v11
    case "":
        v5 += 209
    default:
        v5 = (v5 >> ((0) & 63))
    }
    switch ("" + (("foo bar" + (("\n\t" + ((("a" + ("" + "")) + "a") + "foo bar")) + ("a" + ""))) + ("" + "a"))) {
    case "hello":
        if ("" <= "") {
            if false {
                v6 = v1
            }
            var v12 int32 = 2147483647
            _ = v12
        }
    default:
        for fi13 := 0; fi13 < 18; fi13++ {
            _ = len(v7)
            v14 := GenS1{F0: uint16(v5), F1: 42}
            _ = fi13
            _ = v14
        }
    }
    if (1 > -2147483648) {
        if false {
            for fi15 := 0; fi15 < 3; fi15++ {
                for fi16 := 0; fi16 < 12; fi16++ {
                    switch (v9 - -470) {
                    case -1:
                        if ((1000 < 42) && true) {
                            var v17 float32 = (0.1 + -2.25)
                            if len(v7) > 0 {
                                v7[0] = (^(v9))
                            }
                            var v18 float64 = v3
                            _ = v17
                            _ = v18
                        }
                    case -105:
                        switch "hello" {
                        case "foo bar":
                            v19 := GenS1{F0: 0, F1: 1}
                            _ = v19
                        }
                    }
                    v3 = (float64(1e10) + v3)
                    _ = fi16
                }
                _ = fi15
            }
        }
        fmt.Println(v5, v9, v9)
        fmt.Println(v6)
    } else {
        if false {
            v4[v1] = (^(v6))
            var v20 int16 = 32767
            for fi21 := 0; fi21 < 16; fi21++ {
                v22 := GenS1{F0: 2, F1: -1}
                _ = fi21
                _ = v22
            }
            _ = v20
        } else {
            fmt.Println(v3, v6)
            fmt.Println(v5, v6, v5)
        }
    }
    v23 := GenS2{F0: (false && true), F1: -64, F2: uint16(v6)}
    switch (^(v9)) {
    case -798:
        fmt.Println(v5, v9, v9)
    case -1:
        if false {
            var v24 int = (v9 + v9)
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
            var v25 float32 = (((1.5 - float32(0.1)) * 0.0) / 1.0)
            _ = v24
            _ = v25
        } else {
            for fi26 := 0; fi26 < 15; fi26++ {
                v5 = 128
                for fi27 := 0; fi27 < 1; fi27++ {
                    if true {
                        v3 -= float64((((1.5 / -2.25) / -2.25) / 1.0))
                    } else {
                        var v28 float64 = 1e10
                        _ = v28
                    }
                    if true {
                        v6 += v1
                        fmt.Println(v1)
                        var v29 uint32 = 464
                        _ = v29
                    } else {
                        var v30 int16 = -32768
                        _ = v30
                    }
                    _ = fi27
                }
                var v31 int64 = -581
                _ = fi26
                _ = v31
            }
        }
    case 127:
        var v32 uint8 = v5
        _ = v32
    }
    fmt.Println(v3)
    v6 += 4294967295
    var v33 bool = false
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
