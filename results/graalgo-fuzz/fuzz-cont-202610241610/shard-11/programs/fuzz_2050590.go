package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8 }

func main() {
    var v1 uint32 = 2147483648
    v2 := make(map[int]string)
    v3 := []int{32768, 872}
    v4 := make(map[uint32]string)
    var v5 int64 = 279
    for fi6 := 0; fi6 < 7; fi6++ {
        fmt.Println(fi6, v1, v5)
        var v7 uint16 = 254
        for fi8 := 0; fi8 < 3; fi8++ {
            delete(v4, (v1 << ((437) & 63)))
            var v9 int8 = -128
            var v10 int32 = 1
            v10 += 962
            _ = fi8
            _ = v10
            _ = v9
        }
        _ = fi6
        _ = v7
    }
    var v11 uint16 = 1
    v11 += (v11 / ((v11) | 1))
    var v12 float32 = (float32(3.14159) - 0.1)
    switch 9223372036854775807 {
    case -120:
        v5 = -1
    case -558:
        if false {
            if false {
                v13 := GenS1{F0: 0, F1: -32}
                switch -844 {
                case -873:
                    var v14 uint32 = v1
                    _ = v14
                default:
                    for fi15 := 0; fi15 < 2; fi15++ {
                        for fi16 := 0; fi16 < 14; fi16++ {
                            v11 += (v11 + (v11 + (v11 | v11)))
                            v3 = append(v3, -983)
                            _ = fi16
                        }
                        _ = fi15
                    }
                }
                _ = v13
            }
        } else {
            fmt.Println(v5, v11)
            if true {
                _ = len(v3)
            }
            _ = len(v3)
        }
    default:
        fmt.Println(v1)
    }
    switch "hello" {
    case "foo bar":
        v1 = (v1 >> ((2) & 63))
    default:
        var v17 int32 = -776
        _ = v17
    }
    var v18 uint8 = 255
    switch "" {
    case "":
        var v19 string = ("foo bar" + "")
        _ = v19
    case "\n\t":
        fmt.Println(v18, v12)
    case "hello":
        var v20 uint8 = 127
        _ = v20
    }
    v21 := GenS1{F0: v18, F1: -128}
    v5 += v5
    switch 1 {
    case -245:
        v22 := GenS1{F0: v18, F1: 100}
        _ = v22
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v18
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
