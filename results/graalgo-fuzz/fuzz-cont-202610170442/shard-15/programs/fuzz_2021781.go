package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32 }

func main() {
    v1 := make(map[string]string)
    v2 := []uint32{}
    var v3 int8 = 100
    var v4 float32 = 0.1
    v5 := make(map[uint32]uint32)
    switch "" {
    case "":
        var v6 int = 4294967295
        _ = v6
    default:
        if ((-2.25 <= 1e10) && (2 >= 485)) {
            v4 = 3.14159
            var v7 uint32 = 2147483647
            var v8 int = -355
            _ = v7
            _ = v8
        }
    }
    switch 42 {
    case -9223372036854775808:
        delete(v5, 65535)
    case 9223372036854775807:
        switch -2147483648 {
        case 42:
            var v9 uint8 = 0
            _ = v9
        case -673:
            var v10 int16 = -941
            _ = v10
        }
    case -621:
        v4 = (v4 + v4)
    }
    v3 = (v3 ^ v3)
    for fi11 := 0; fi11 < 20; fi11++ {
        if true {
            v12 := GenS1{F0: 42, F1: -289}
            for fi13 := 0; fi13 < 11; fi13++ {
                var v14 uint16 = 113
                _ = fi13
                _ = v14
            }
            _ = v12
        }
        v5[2147483648] = 4294967295
        _ = len(v1)
        _ = fi11
    }
    switch 0 {
    case 485:
        v4 += 0.1
    default:
        switch "a" {
        case "\n\t":
            if len(v2) > 0 {
                v2[0] = 78
            }
        case "hello":
            var v15 uint8 = 1
            _ = v15
        }
    }
    var v16 int64 = -423
    for fi17 := 0; fi17 < 20; fi17++ {
        _ = len(v2)
        switch "" {
        case "foo bar":
            var v18 int32 = -1
            _ = v18
        case "\n\t":
            var v19 int16 = -671
            _ = v19
        }
        _ = fi17
    }
    var v20 uint16 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v16
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
