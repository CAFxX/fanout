package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16 }
type GenS2 struct { F0 uint64; F1 int }

func main() {
    var v1 uint = 255
    v2 := []string{}
    var v3 uint = (v1 ^ 0)
    v4 := GenS1{F0: 42, F1: -32768}
    var v5 uint64 = 142
    if len(v2) > 0 {
        v2[0] = "a"
    }
    v6 := GenS1{F0: -1, F1: 1000}
    var v7 int16 = 999
    switch "foo bar" {
    case "\n\t":
        v1 = (v3 | 65535)
    case "foo bar":
        v5 -= (v5 + v5)
    case "a":
        if true {
            if (v3 == v1) {
                v8 := GenS1{F0: v7, F1: 1}
                var v9 uint32 = 654
                if true {
                    v3 = 957
                }
                _ = v8
                _ = v9
            } else {
                if (-28 <= -1) {
                    v7 = -532
                    if false {
                        v5 = v5
                    } else {
                        var v10 uint16 = 551
                        _ = v10
                    }
                    if (-32768 >= (v7 << ((v3) & 63))) {
                        v3 = 720
                        if len(v2) > 0 {
                            v2[0] = ("" + "hello")
                        }
                    }
                }
                v5 = 2
                v5 -= (v5 | (^(v5)))
            }
            v1 += v1
            if false {
                switch int(v7) {
                case -2147483648:
                    if true {
                        var v11 uint16 = 42
                        v1 = (^(v3))
                        switch "foo bar" {
                        case "":
                            if (35 < 1) {
                                var v12 int64 = -676
                                _ = v12
                            } else {
                                var v13 int = 9223372036854775807
                                _ = v13
                            }
                        default:
                            if true {
                                var v14 uint8 = 255
                                var v15 uint8 = 190
                                _ = v14
                                _ = v15
                            } else {
                                switch "\n\t" {
                                case "foo bar":
                                    var v16 float32 = float32((float64(1e10) + (3.14159 + 0.1)))
                                    _ = v16
                                case "hello":
                                    v5 = v5
                                default:
                                    v1 -= v1
                                }
                            }
                        }
                        _ = v11
                    }
                case 42:
                    for fi17 := 0; fi17 < 9; fi17++ {
                        var v18 int = fi17
                        _ = fi17
                        _ = v18
                    }
                }
                var v19 float64 = float64(3.14159)
                _ = v19
            } else {
                var v20 uint8 = 214
                _ = v20
            }
        }
    default:
        v21 := GenS2{F0: v5, F1: -2147483648}
        _ = v21
    }
    v5 = v5
    var v22 uint8 = 9
    var v23 float64 = 0.1
    v7 += (v7 | 1000)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
