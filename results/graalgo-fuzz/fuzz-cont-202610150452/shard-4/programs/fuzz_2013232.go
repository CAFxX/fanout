package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16; F2 uint32 }
type GenS2 struct { F0 string; F1 int64; F2 uint }

func main() {
    var v1 int64 = 42
    var v2 int = -1
    var v3 int = (v2 * 127)
    var v4 uint64 = 18446744073709551615
    var v5 uint32 = 1
    v6 := []uint32{2147483647, (^(v5)), 2, (v5 / ((v5) | 1))}
    if len(v6) > 0 {
        v6[0] = v5
    }
    v7 := GenS1{F0: 65535, F1: 1000, F2: (^(v5))}
    if len(v6) > 0 {
        v6[0] = 179
    }
    v4 += v4
    v2 = v2
    for fi8 := 0; fi8 < 11; fi8++ {
        if true {
            v1 -= v1
            v9 := GenS2{F0: "hello", F1: 9223372036854775807, F2: 2}
            _ = v9
        }
        switch fi8 {
        case 4294967295:
            for fi10 := 0; fi10 < 1; fi10++ {
                v11 := GenS2{F0: "hello", F1: -48, F2: 0}
                v12 := GenS1{F0: 164, F1: 32767, F2: (v5 % ((v5) | 1))}
                _ = fi10
                _ = v11
                _ = v12
            }
        case 0:
            var v13 string = "\n\t"
            _ = v13
        }
        _ = fi8
    }
    v5 = 2147483648
    var v14 uint8 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
