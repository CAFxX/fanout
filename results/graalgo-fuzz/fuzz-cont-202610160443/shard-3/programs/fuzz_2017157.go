package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint32; F1 int64; F2 int32 }

func main() {
    var v1 uint32 = 2
    v2 := []uint32{v1, (^(v1))}
    var v3 int64 = 1
    v4 := make(map[uint32]bool)
    var v5 int = 128
    v6 := make(map[uint32]string)
    v3 = -9223372036854775808
    var v7 bool = true
    if (0 >= 18446744073709551615) {
        if ((!true) && true) {
            v7 = (255 >= 210)
        }
        var v8 uint = 0
        _ = v8
    }
    v3 -= (v3 - (v3 % ((v3) | 1)))
    v3 = v3
    v9 := GenS1{F0: 888}
    v7 = (0 >= 4294967295)
    for fi10 := 0; fi10 < 15; fi10++ {
        for fi11 := 0; fi11 < 16; fi11++ {
            v1 = v1
            _ = fi11
        }
        v12 := GenS2{F0: (^(v1)), F1: 0, F2: 329}
        _ = fi10
        _ = v12
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
