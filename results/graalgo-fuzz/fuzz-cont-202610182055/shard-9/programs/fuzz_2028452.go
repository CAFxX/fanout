package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint16; F1 int8 }
type GenS3 struct { F0 uint8 }

func main() {
    var v1 uint32 = 1
    v2 := []int{}
    v3 := []string{"a"}
    var v4 float32 = float32(0.0)
    var v5 int64 = -1
    switch 759 {
    case 1:
        for fi6 := 0; fi6 < 1; fi6++ {
            v7 := GenS1{F0: 243}
            _ = fi6
            _ = v7
        }
    case 128:
        switch "" {
        case "a":
            v5 = v5
        case "":
            v8 := GenS2{F0: 1, F1: 1}
            _ = v8
        case "\n\t":
            v3 = append(v3, "\n\t")
        default:
            v9 := GenS1{F0: 65}
            _ = v9
        }
    }
    _ = len(v3)
    var v10 uint8 = 118
    v4 -= v4
    v11 := GenS3{F0: 2}
    v4 = (((-2.25 * v4) - v4) - v4)
    v10 = (v10 << ((127) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
