package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32 }

func main() {
    v1 := make(map[int]uint32)
    var v2 uint64 = 117
    var v3 uint16 = 2
    v4 := []int{583, 32768, 747, 0}
    v5 := make(map[string]bool)
    var v6 uint = 0
    var v7 uint32 = 114
    fmt.Println(v3, v2)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v8 := GenS1{F0: (v2 == (v2 % (((^(v2))) | 1))), F1: uint32(v3)}
    var v9 uint8 = 42
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v10 int = 32767
    for fi11 := 0; fi11 < 1; fi11++ {
        for fi12 := 0; fi12 < 20; fi12++ {
            v13 := GenS1{F0: (false && true), F1: v7}
            _ = fi12
            _ = v13
        }
        _ = fi11
    }
    v6 = (v6 << ((1) & 63))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
