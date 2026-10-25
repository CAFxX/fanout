package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int8; F1 bool; F2 uint16 }
type GenS3 struct { F0 uint16; F1 uint8 }

func main() {
    var v1 float64 = 3.14159
    var v2 int32 = -671
    var v3 int = 9223372036854775807
    v4 := make(map[uint32]uint32)
    var v5 int8 = -1
    v6 := make(map[uint32]bool)
    v1 = 1.5
    v7 := GenS1{F0: 204}
    _ = len(v4)
    for fi8 := 0; fi8 < 2; fi8++ {
        var v9 float64 = -2.25
        _ = fi8
        _ = v9
    }
    var v10 uint32 = 2147483648
    v11 := GenS2{F0: 127, F1: ((-9223372036854775808 == -803) && (!true)), F2: 255}
    var v12 string = "\n\t"
    v13 := GenS3{F0: 537, F1: 2}
    v14 := GenS1{F0: v3}
    switch v12 {
    case "hello":
        for fi15 := 0; fi15 < 8; fi15++ {
            fmt.Println(v1, v2, v12)
            var v16 uint8 = 42
            _ = fi15
            _ = v16
        }
    case "\n\t":
        var v17 uint64 = 880
        _ = v17
    case "":
        switch v12 {
        case "hello":
            v2 = v2
        default:
            delete(v4, (v10 << ((0) & 63)))
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
