package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64 }
type GenS2 struct { F0 uint64; F1 int8; F2 uint }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint32; F1 bool }

func main() {
    v1 := []int{32767, 42}
    var v2 uint = 218
    v3 := make(map[int]bool)
    v4 := make(map[uint32]string)
    var v5 float32 = (float32((0.1 - 1.5)) * 3.14159)
    v6 := make(map[string]bool)
    v4[724] = "foo bar"
    var v7 uint32 = 320
    for fi8 := 0; fi8 < 15; fi8++ {
        v5 += (1e10 + (-2.25 * (v5 + ((-2.25 + (v5 + v5)) / (0.0 - v5)))))
        if (true && (-628 >= -189)) {
            var v9 float32 = float32((0.0 + 0.0))
            if ("foo bar" != "\n\t") {
                switch "" {
                case "\n\t":
                    fmt.Println(v7)
                }
                v10 := GenS1{F0: (v7 + v7), F1: 1}
                _ = v10
            } else {
                switch "\n\t" {
                case "foo bar":
                    v4[(v7 | (v7 & (v7 | (v7 >> (((v2 / ((v2) | 1))) & 63)))))] = ("a" + "")
                case "a":
                    var v11 int32 = -1
                    _ = v11
                }
                for fi12 := 0; fi12 < 12; fi12++ {
                    fmt.Println(fi8)
                    _ = fi12
                }
                v2 -= 350
            }
            _ = v9
        } else {
            v13 := GenS2{F0: 4294967295, F1: -1, F2: (v2 + (v2 << (((v2 | v2)) & 63)))}
            v2 += v2
            _ = v13
        }
        _ = fi8
    }
    if (!((false && false) || true)) {
        switch ((("\n\t" + (("hello" + ("\n\t" + "")) + ("hello" + ("a" + "\n\t")))) + "hello") + ("" + "\n\t")) {
        case "\n\t":
            var v14 int = 4294967295
            _ = v14
        }
    }
    for fi15 := 0; fi15 < 19; fi15++ {
        if true {
            v5 = (v5 + v5)
            _ = len(v6)
            if (!false) {
                for fi16 := 0; fi16 < 1; fi16++ {
                    v7 += 2147483648
                    fmt.Println(v5)
                    v7 += (^(v7))
                    _ = fi16
                }
                v7 = (v7 << (((v2 * (v2 & (v2 * v2)))) & 63))
                v7 += v7
            } else {
                if true {
                    if len(v1) > 0 {
                        v1[0] = (fi15 % ((fi15) | 1))
                    }
                    v17 := GenS1{F0: (v7 + uint32(v2)), F1: 42}
                    _ = v17
                }
                v7 += uint32(v2)
            }
        } else {
            switch ("foo bar" + "hello") {
            case "":
                var v18 int8 = -104
                _ = v18
            case "a":
                switch 127 {
                case 32768:
                    var v19 uint = (v2 * 2)
                    _ = v19
                case 0:
                    v2 = 42
                case -1:
                    if false {
                        fmt.Println(v2)
                    } else {
                        var v20 uint64 = 2
                        switch 127 {
                        case 1:
                            if (-2147483648 < 467) {
                                v7 = (v7 >> (((v2 | (v2 ^ (^(v2))))) & 63))
                                for fi21 := 0; fi21 < 7; fi21++ {
                                    v2 = (v2 * (v2 ^ v2))
                                    _ = fi21
                                }
                                _ = len(v1)
                            }
                        case -731:
                            switch (("hello" + ("a" + ("hello" + "\n\t"))) + "foo bar") {
                            case "a":
                                var v22 float64 = (1.5 * (float64(v5) * -2.25))
                                _ = v22
                            default:
                                for fi23 := 0; fi23 < 7; fi23++ {
                                    v2 = 18446744073709551615
                                    _ = fi23
                                }
                            }
                        case 127:
                            for fi24 := 0; fi24 < 12; fi24++ {
                                v25 := GenS2{F0: 18446744073709551615, F1: 0, F2: 0}
                                _ = fi24
                                _ = v25
                            }
                        default:
                            v26 := GenS3{F0: 405}
                            _ = v26
                        }
                        _ = v20
                    }
                default:
                    v7 += (v7 | 342)
                }
            default:
                fmt.Println(v2, v7)
            }
            v27 := GenS4{F0: 461, F1: (869 != 65)}
            _ = v27
        }
        var v28 uint64 = 0
        _ = fi15
        _ = v28
    }
    v29 := GenS1{F0: (v7 + (v7 << (((v2 / ((v2) | 1))) & 63))), F1: -367}
    var v30 uint8 = 255
    switch "\n\t" {
    case "a":
        var v31 string = ""
        _ = v31
    }
    var v32 uint16 = 42
    var v33 int = 2147483647
    var v34 int = v33
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
