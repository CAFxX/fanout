package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8; F2 uint64 }

func main() {
    var v1 uint = 42
    var v2 float32 = 0.1
    v3 := []string{}
    var v4 int8 = -128
    switch "foo bar" {
    case "\n\t":
        for fi5 := 0; fi5 < 7; fi5++ {
            v2 -= float32(1e10)
            if true {
                v1 = (^(v1))
                switch 2147483647 {
                case -542:
                    _ = len(v3)
                case -2147483648:
                    if false {
                        if (!true) {
                            var v6 uint8 = 1
                            _ = v6
                        }
                        var v7 float64 = float64(v2)
                        _ = v7
                    }
                case -170:
                    var v8 uint64 = 64
                    _ = v8
                default:
                    var v9 string = "a"
                    _ = v9
                }
            } else {
                var v10 string = "a"
                _ = v10
            }
            for fi11 := 0; fi11 < 11; fi11++ {
                v12 := GenS1{F0: -1, F1: 100, F2: 0}
                v1 = (v1 / (((v1 / ((v1) | 1))) | 1))
                v2 = v2
                _ = fi11
                _ = v12
            }
            _ = fi5
        }
    case "foo bar":
        var v13 uint32 = 300
        _ = v13
    case "":
        if len(v3) > 0 {
            v3[0] = ""
        }
    default:
        v14 := GenS1{F0: -1, F1: v4, F2: 4294967295}
        _ = v14
    }
    if (v1 > 255) {
        for fi15 := 0; fi15 < 14; fi15++ {
            v4 = (v4 >> (((v1 >> ((v1) & 63))) & 63))
            _ = fi15
        }
        var v16 int32 = 42
        _ = v16
    }
    var v17 float32 = ((v2 - v2) / 1.0)
    var v18 uint32 = 2147483648
    v19 := GenS1{F0: 400, F1: v4, F2: 18446744073709551615}
    var v20 uint64 = 4294967295
    var v21 bool = (true || (!false))
    v22 := GenS1{F0: -798, F1: 106, F2: v20}
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v23 int32 = 2147483647
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
}
