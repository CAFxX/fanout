package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int64; F1 int16; F2 int32 }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[string]int)
    var v3 uint64 = 18446744073709551615
    var v4 bool = ((false && (0 >= -9223372036854775808)) || (127 >= 128))
    for fi5 := 0; fi5 < 15; fi5++ {
        v3 -= 2
        _ = fi5
    }
    v6 := GenS1{F0: (v3 << ((65535) & 63))}
    fmt.Println(v3, v4)
    v3 = 1
    var v7 int = -1
    fmt.Println(v4)
    if ((true && true) || v4) {
        fmt.Println(v7)
        for fi8 := 0; fi8 < 12; fi8++ {
            if v4 {
                v9 := GenS2{F0: 744, F1: 32767, F2: -1}
                fmt.Println(v4)
                var v10 uint64 = (v3 % ((2) | 1))
                _ = v10
                _ = v9
            }
            for fi11 := 0; fi11 < 18; fi11++ {
                v3 = v3
                v12 := GenS1{F0: v3}
                _ = fi11
                _ = v12
            }
            v13 := []int{-537}
            _ = fi8
            _ = v13
        }
    }
    var v14 uint32 = 42
    var v15 int8 = 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
}
