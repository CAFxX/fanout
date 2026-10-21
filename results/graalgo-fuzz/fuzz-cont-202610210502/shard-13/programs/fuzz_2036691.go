package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8; F2 int8 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int16; F1 bool }
type GenS4 struct { F0 int; F1 uint }

func main() {
    var v1 uint64 = 366
    v2 := make(map[string]bool)
    var v3 uint16 = 1
    v4 := make(map[uint32]int)
    for fi5 := 0; fi5 < 8; fi5++ {
        switch (fi5 & (fi5 % ((fi5) | 1))) {
        case 128:
            v3 -= (v3 / ((v3) | 1))
        default:
            v2["a"] = (!true)
        }
        v1 += (v1 + 1)
        _ = fi5
    }
    v3 = 1
    delete(v2, "\n\t")
    v6 := GenS1{F0: 127, F1: 2, F2: 42}
    v7 := GenS2{F0: (v1 - 4294967295)}
    if (18446744073709551615 <= v1) {
        var v8 int64 = -1
        v8 -= (v8 + v8)
        v3 -= (v3 ^ (v3 >> ((131) & 63)))
        _ = v8
    }
    var v9 float32 = 3.14159
    switch "foo bar" {
    case "foo bar":
        v10 := GenS3{F0: 0, F1: (!(true && true))}
        _ = v10
    case "\n\t":
        v11 := GenS4{F0: 127, F1: 348}
        _ = v11
    default:
        v3 = (v3 >> ((18446744073709551615) & 63))
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
    _ = v9
}
