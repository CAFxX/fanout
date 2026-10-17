package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16 }
type GenS2 struct { F0 bool; F1 uint64; F2 int }

func main() {
    var v1 uint32 = 2147483647
    var v2 int64 = 9223372036854775807
    var v3 int32 = -685
    v4 := GenS1{F0: 1, F1: 255}
    v3 = v3
    v3 -= (-(v3))
    for fi5 := 0; fi5 < 16; fi5++ {
        if true {
            fmt.Println(fi5, v1, fi5)
        }
        _ = fi5
    }
    var v6 int32 = v3
    v6 = int32(v1)
    var v7 uint64 = 0
    var v8 uint8 = 255
    if (0.1 == ((-2.25 - 3.14159) / 0.1)) {
        v1 -= 0
        v2 += (v2 << ((18446744073709551615) & 63))
        v9 := GenS2{F0: (!false), F1: (v7 / ((0) | 1)), F2: 32768}
        _ = v9
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
    _ = v8
}
