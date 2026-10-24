package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32 }

func main() {
    var v1 int8 = 127
    var v2 uint32 = 2
    var v3 int8 = (^(v1))
    v4 := []uint32{}
    var v5 int16 = -1
    for fi6 := 0; fi6 < 2; fi6++ {
        fmt.Println(v2, v5)
        fmt.Println(v2, v5, v5)
        v1 = (v1 / ((v1) | 1))
        _ = fi6
    }
    var v7 uint8 = 2
    if true {
        for fi8 := 0; fi8 < 18; fi8++ {
            for fi9 := 0; fi9 < 18; fi9++ {
                if len(v4) > 0 {
                    v4[0] = v2
                }
                switch "a" {
                case "hello":
                    var v10 int32 = -1
                    _ = v10
                case "foo bar":
                    if ((4294967295 < 1) && (!(255 > 18446744073709551615))) {
                        v11 := GenS1{F0: -353, F1: 2147483648}
                        var v12 int32 = 477
                        _ = v11
                        _ = v12
                    } else {
                        if (false && true) {
                            v3 += -111
                        }
                        if len(v4) > 0 {
                            v4[0] = 65535
                        }
                    }
                case "\n\t":
                    if ((127 < 126) || (false && false)) {
                        fmt.Println(fi8)
                    }
                default:
                    fmt.Println(fi9, fi9, v7)
                }
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                _ = fi9
            }
            _ = fi8
        }
    }
    v13 := GenS1{F0: 1, F1: v2}
    for fi14 := 0; fi14 < 9; fi14++ {
        v3 = (^(v3))
        _ = fi14
    }
    for fi15 := 0; fi15 < 12; fi15++ {
        var v16 int = fi15
        if false {
            if (!(false && true)) {
                _ = len(v4)
                var v17 int = (v16 >> ((uint(v7)) & 63))
                fmt.Println(v2)
                _ = v17
            } else {
                var v18 uint64 = 18446744073709551615
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                _ = v18
            }
        }
        _ = fi15
        _ = v16
    }
    fmt.Println(v3, v1)
    if len(v4) > 0 {
        v4[0] = 551
    }
    if (int8(v2) == (v1 ^ 7)) {
        var v19 uint16 = 643
        switch "" {
        case "foo bar":
            v1 -= (v1 * 0)
        }
        var v20 int32 = 0
        _ = v19
        _ = v20
    } else {
        fmt.Println(v1, v3)
        v3 = 100
    }
    v1 = v1
    var v21 int16 = 0
    v22 := GenS1{F0: 1, F1: v2}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
