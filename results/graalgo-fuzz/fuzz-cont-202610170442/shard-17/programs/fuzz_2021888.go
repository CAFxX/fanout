package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32 }
type GenS2 struct { F0 int16; F1 int8 }

func main() {
    var v1 int32 = 42
    v2 := []uint32{}
    var v3 uint16 = 453
    var v4 uint8 = 247
    v5 := make(map[int]string)
    var v6 int32 = (v1 - -201)
    v4 = 2
    var v7 int = 9223372036854775807
    v1 += 0
    v8 := GenS1{F0: 42, F1: 968}
    if (v7 != -342) {
        v9 := GenS2{F0: int16(v3), F1: -42}
        _ = v9
    }
    fmt.Println(v7)
    v2 = append(v2, 65535)
    if false {
        for fi10 := 0; fi10 < 3; fi10++ {
            if len(v2) > 0 {
                v2[0] = 0
            }
            _ = fi10
        }
    } else {
        switch "hello" {
        case "\n\t":
            var v11 int16 = 1000
            _ = v11
        case "":
            v4 = v4
        default:
            for fi12 := 0; fi12 < 14; fi12++ {
                var v13 int = fi12
                _ = len(v5)
                _ = fi12
                _ = v13
            }
        }
        var v14 int = (v7 << ((251) & 63))
        _ = v14
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
