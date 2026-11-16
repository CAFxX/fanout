package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int16; F1 uint8 }
type GenS3 struct { F0 uint; F1 uint8 }

func main() {
    var v1 uint64 = 42
    v2 := make(map[uint32]uint32)
    var v3 uint16 = 298
    v4 := []string{"hello", "\n\t", "hello", ""}
    var v5 uint8 = 255
    var v6 uint32 = 42
    if (!true) {
        fmt.Println(v6, v3)
        var v7 uint16 = uint16(v5)
        _ = v7
    } else {
        for fi8 := 0; fi8 < 18; fi8++ {
            for fi9 := 0; fi9 < 7; fi9++ {
                var v10 float32 = 0.0
                var v11 uint32 = v6
                _ = fi9
                _ = v10
                _ = v11
            }
            _ = fi8
        }
    }
    v6 = (v6 & (v6 * (^(v6))))
    v12 := v2[v6]
    v13 := GenS1{F0: v5}
    for fi14 := 0; fi14 < 16; fi14++ {
        for fi15 := 0; fi15 < 3; fi15++ {
            v6 += v6
            if (330 >= 668) {
                switch "hello" {
                case "":
                    v3 -= v3
                case "foo bar":
                    switch "foo bar" {
                    case "":
                        var v16 int32 = 2147483647
                        _ = v16
                    case "\n\t":
                        var v17 uint64 = v1
                        _ = v17
                    default:
                        v5 -= v5
                    }
                case "hello":
                    var v18 int16 = 42
                    _ = v18
                default:
                    fmt.Println(v6, v6)
                }
                v19 := GenS1{F0: v5}
                switch "hello" {
                case "":
                    v5 = v5
                default:
                    switch -843 {
                    case 2147483647:
                        var v20 uint32 = v6
                        _ = v20
                    default:
                        v6 = 4294967295
                    }
                }
                _ = v19
            }
            var v21 float64 = 0.0
            _ = fi15
            _ = v21
        }
        _ = fi14
    }
    for fi22 := 0; fi22 < 7; fi22++ {
        var v23 uint16 = v3
        v23 -= (v3 >> ((4294967295) & 63))
        var v24 int8 = 1
        _ = fi22
        _ = v23
        _ = v24
    }
    if (102 == -128) {
        for fi25 := 0; fi25 < 14; fi25++ {
            for fi26 := 0; fi26 < 14; fi26++ {
                var v27 uint = 42
                v5 += v5
                _ = fi26
                _ = v27
            }
            v28 := GenS2{F0: 608, F1: (v5 ^ (v5 / ((42) | 1)))}
            _ = fi25
            _ = v28
        }
        _ = len(v4)
    }
    v29 := GenS3{F0: 2, F1: v5}
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
