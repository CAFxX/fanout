package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int; F1 int16 }
type GenS4 struct { F0 int32; F1 int64 }
type GenS5 struct { F0 uint; F1 string; F2 string }

func main() {
    v1 := make(map[uint32]int)
    var v2 uint = 255
    v3 := make(map[string]string)
    v4 := []uint32{2147483647, 42, 2147483647, 2}
    var v5 uint8 = 42
    var v6 bool = false
    var v7 float32 = 3.14159
    for fi8 := 0; fi8 < 7; fi8++ {
        for fi9 := 0; fi9 < 17; fi9++ {
            if false {
                v3["foo bar"] = "a"
                v2 -= (^(v2))
                v3[("\n\t" + ("foo bar" + "hello"))] = "hello"
            }
            if ((!true) && (!v6)) {
                v2 = 4294967295
            }
            var v10 float64 = 1.5
            _ = fi9
            _ = v10
        }
        fmt.Println(v5, v5, v7)
        _ = fi8
    }
    v2 = (v2 * 255)
    var v11 float32 = ((float32(float64(v7)) / 1.5) - v7)
    switch "" {
    case "a":
        for fi12 := 0; fi12 < 4; fi12++ {
            for fi13 := 0; fi13 < 2; fi13++ {
                v14 := GenS1{F0: int(v5)}
                if v6 {
                    switch "foo bar" {
                    case "\n\t":
                        var v15 float32 = v11
                        _ = v15
                    case "hello":
                        var v16 uint16 = 830
                        _ = v16
                    default:
                        switch (fi12 % (((^(fi12))) | 1)) {
                        case -619:
                            var v17 float64 = (0.1 * float64(((1.5 / (0.0 * (float32(((1.5 - ((1.5 - float64(v7)) - (1e10 * 0.1))) / (float64(v7) + -2.25))) + v11))) * v11)))
                            _ = v17
                        case 127:
                            v2 += (v2 * v2)
                        }
                    }
                    _ = len(v3)
                    var v18 uint = 42
                    if v6 {
                        fmt.Println(v2)
                    }
                    _ = v18
                } else {
                    fmt.Println(v2, v7)
                    v19 := GenS1{F0: 32767}
                    fmt.Println(fi12, v11)
                    _ = v19
                }
                _ = fi13
                _ = v14
            }
            var v20 int8 = 1
            _ = fi12
            _ = v20
        }
    case "\n\t":
        if true {
            if (0 <= 365) {
                switch 328 {
                case 128:
                    var v21 int64 = -314
                    _ = v21
                case 260:
                    v22 := GenS2{F0: 616}
                    _ = v22
                case 2147483647:
                    v2 += v2
                default:
                    fmt.Println(v6, v11)
                }
            } else {
                for fi23 := 0; fi23 < 19; fi23++ {
                    switch fi23 {
                    case 32767:
                        fmt.Println(v11)
                    }
                    fmt.Println(v5, v6)
                    var v24 float32 = (v7 * (v7 / v11))
                    _ = fi23
                    _ = v24
                }
                for fi25 := 0; fi25 < 15; fi25++ {
                    v6 = false
                    v5 = v5
                    var v26 float32 = (v7 / v7)
                    _ = fi25
                    _ = v26
                }
                switch 211 {
                case 452:
                    switch ("a" + "a") {
                    case "foo bar":
                        var v27 int32 = 0
                        _ = v27
                    default:
                        v2 = v2
                    }
                default:
                    for fi28 := 0; fi28 < 17; fi28++ {
                        fmt.Println(v5)
                        switch (^(fi28)) {
                        case -779:
                            var v29 int8 = 8
                            _ = v29
                        default:
                            if true {
                                v2 = (v2 + v2)
                            }
                        }
                        _ = fi28
                    }
                }
            }
            v11 += float32(1.5)
        }
    default:
        if len(v4) > 0 {
            v4[0] = 2147483647
        }
    }
    var v30 uint8 = (v5 * 255)
    v30 = v30
    if len(v4) > 0 {
        v4[0] = 1
    }
    fmt.Println(v7)
    if v6 {
        if (!(v2 > 42)) {
            v31 := GenS3{F0: -770, F1: -860}
            v32 := GenS4{F0: -1, F1: 9223372036854775807}
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = v31
            _ = v32
        }
        var v33 int8 = -128
        v34 := GenS5{F0: (v2 * v2), F1: ("a" + "a"), F2: "a"}
        _ = v33
        _ = v34
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
