package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }

func main() {
    v1 := make(map[uint32]int)
    var v2 float32 = 1.5
    v3 := make(map[int]bool)
    var v4 int16 = 729
    v5 := []uint32{339, 317, 606}
    var v6 uint16 = 0
    v5 = append(v5, 2)
    if len(v5) > 0 {
        v5[0] = 0
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v5 = append(v5, 1)
    v7 := GenS1{F0: (v6 * (^(v6)))}
    for fi8 := 0; fi8 < 18; fi8++ {
        switch "a" {
        case "":
            _ = len(v5)
        }
        _ = fi8
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
