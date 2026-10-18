package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64; F2 int16 }

func main() {
    var v1 int16 = 1000
    v2 := make(map[int]string)
    v3 := make(map[int]bool)
    v4 := []string{"a", "", ("" + (("\n\t" + ("\n\t" + "")) + ("" + "\n\t")))}
    fmt.Println(v1)
    v5 := GenS1{F0: 128, F1: 2, F2: v1}
    fmt.Println(v1)
    for fi6 := 0; fi6 < 7; fi6++ {
        for fi7 := 0; fi7 < 10; fi7++ {
            v1 = v1
            var v8 int16 = (v1 % ((v1) | 1))
            v3[(fi6 | fi6)] = true
            _ = fi7
            _ = v8
        }
        _ = fi6
    }
    v1 = (-(v1))
    if false {
        v1 -= v1
        var v9 int = -47
        switch "\n\t" {
        case "hello":
            var v10 uint32 = 469
            _ = v10
        case "\n\t":
            var v11 uint64 = 0
            _ = v11
        case "a":
            var v12 bool = true
            _ = v12
        default:
            if len(v4) > 0 {
                v4[0] = "\n\t"
            }
        }
        _ = v9
    }
    switch 0 {
    case 4294967295:
        v13 := GenS1{F0: 1, F1: 1, F2: -32768}
        _ = v13
    default:
        v1 = (v1 + (v1 * 899))
    }
    v1 = (v1 << ((18446744073709551615) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
