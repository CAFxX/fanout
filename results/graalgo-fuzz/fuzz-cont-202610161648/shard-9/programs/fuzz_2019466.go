package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64 }

func main() {
    var v1 float32 = 1.5
    var v2 uint8 = 1
    v3 := make(map[int]uint32)
    var v4 int16 = 0
    var v5 int16 = (^(v4))
    for fi6 := 0; fi6 < 1; fi6++ {
        var v7 int16 = v5
        switch ("" + "\n\t") {
        case "foo bar":
            switch "a" {
            case "hello":
                if (!false) {
                    v4 += (v4 - int16(v2))
                }
            case "":
                v2 = (^(v2))
            default:
                v8 := GenS1{F0: 253, F1: 0}
                _ = v8
            }
        case "a":
            for fi9 := 0; fi9 < 17; fi9++ {
                v7 += (v5 * v7)
                _ = fi9
            }
        case "hello":
            v4 -= -843
        }
        for fi10 := 0; fi10 < 7; fi10++ {
            v4 = (v7 << ((171) & 63))
            _ = fi10
        }
        _ = fi6
        _ = v7
    }
    var v11 uint32 = 721
    v5 += v4
    fmt.Println(v11)
    v1 = v1
    for fi12 := 0; fi12 < 18; fi12++ {
        v13 := v3[(-(fi12))]
        v13 = (v13 + (v13 & (v11 - 4294967295)))
        v11 = (v13 >> ((65535) & 63))
        _ = fi12
        _ = v13
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
