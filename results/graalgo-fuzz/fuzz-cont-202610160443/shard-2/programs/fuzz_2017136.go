package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8 }
type GenS2 struct { F0 uint64; F1 uint16; F2 uint64 }

func main() {
    var v1 uint8 = 96
    var v2 uint32 = 222
    v3 := make(map[uint32]string)
    var v4 bool = (true || (true && true))
    v5 := GenS1{F0: 0, F1: (v1 | (v1 + 211))}
    v6 := GenS2{F0: 1, F1: 0, F2: 42}
    v1 += (v1 << ((402) & 63))
    switch -2147483648 {
    case -2147483648:
        if false {
            v4 = v4
            if false {
                if true {
                    v7 := GenS1{F0: 1, F1: v1}
                    _ = v7
                }
                v2 = 4294967295
            } else {
                v8 := GenS2{F0: 42, F1: 32767, F2: 0}
                _ = v8
            }
        } else {
            v1 = 42
        }
    }
    var v9 uint64 = 975
    var v10 float64 = 0.0
    v10 += v10
    for fi11 := 0; fi11 < 10; fi11++ {
        v12 := GenS1{F0: (v9 ^ (^(v9))), F1: v1}
        _ = len(v3)
        _ = fi11
        _ = v12
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
