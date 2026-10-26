package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string }
type GenS2 struct { F0 int16 }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[int]bool)
    var v3 uint16 = 42
    var v4 uint8 = 127
    v3 -= (v3 >> ((2) & 63))
    fmt.Println(v4, v3)
    v5 := []int{2147483647, 127, -878, 2147483647}
    v3 = 1
    switch "foo bar" {
    case "foo bar":
        v6 := GenS1{F0: 23, F1: "a"}
        _ = v6
    }
    v4 = (v4 ^ (v4 >> ((18446744073709551615) & 63)))
    v3 += (v3 % ((uint16(v4)) | 1))
    for fi7 := 0; fi7 < 18; fi7++ {
        var v8 int = 32767
        for fi9 := 0; fi9 < 7; fi9++ {
            var v10 uint8 = v4
            _ = fi9
            _ = v10
        }
        v2[(fi7 / ((-1) | 1))] = true
        _ = fi7
        _ = v8
    }
    for fi11 := 0; fi11 < 11; fi11++ {
        v4 -= (v4 | (v4 | (v4 - v4)))
        v12 := GenS2{F0: 42}
        fmt.Println(v3)
        _ = fi11
        _ = v12
    }
    for fi13 := 0; fi13 < 4; fi13++ {
        v4 = uint8(v3)
        for fi14 := 0; fi14 < 6; fi14++ {
            var v15 bool = true
            switch fi14 {
            case 4294967295:
                v16 := GenS2{F0: 32767}
                _ = v16
            }
            v17 := GenS2{F0: 0}
            _ = fi14
            _ = v15
            _ = v17
        }
        var v18 uint64 = 4294967295
        _ = fi13
        _ = v18
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
