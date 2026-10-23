package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8 }
type GenS2 struct { F0 bool; F1 uint32; F2 int32 }

func main() {
    v1 := make(map[int]uint32)
    var v2 int16 = -32768
    v3 := []uint32{2147483648, 2147483648, 2147483648}
    var v4 int16 = v2
    var v5 uint = 65535
    v5 = 2
    var v6 uint16 = 42
    var v7 string = "foo bar"
    switch 128 {
    case 128:
        var v8 string = ("" + v7)
        _ = v8
    default:
        v5 = (v5 + v5)
    }
    switch 128 {
    case 128:
        fmt.Println(v6, v7)
    case -2147483648:
        switch 42 {
        case 42:
            var v9 int64 = -793
            _ = v9
        }
    default:
        for fi10 := 0; fi10 < 17; fi10++ {
            var v11 bool = ((fi10 / ((fi10) | 1)) <= (^(fi10)))
            switch ("foo bar" + v7) {
            case "":
                var v12 bool = true
                _ = v12
            }
            fmt.Println(v11, v2, v11)
            _ = fi10
            _ = v11
        }
    }
    v2 -= v4
    v13 := GenS1{F0: -128, F1: 155}
    switch -1 {
    case -1:
        v5 = (v5 >> ((255) & 63))
    case 32768:
        v14 := GenS2{F0: false, F1: 42, F2: -1}
        _ = v14
    case 4294967295:
        v2 -= v2
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
