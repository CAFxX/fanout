package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int32; F1 uint64; F2 uint8 }
type GenS3 struct { F0 uint16; F1 int32 }

func main() {
    var v1 uint8 = 1
    v2 := make(map[int]int)
    var v3 int32 = 154
    v4 := make(map[string]int)
    fmt.Println(v3, v3)
    v1 = (v1 * v1)
    if ((!false) || (-297 == 0)) {
        var v5 int = -645
        switch (v5 | (v5 % ((1) | 1))) {
        case -641:
            v6 := GenS1{F0: 2}
            _ = v6
        case 128:
            v7 := GenS1{F0: 2}
            _ = v7
        default:
            for fi8 := 0; fi8 < 1; fi8++ {
                if (v3 > v3) {
                    v9 := GenS2{F0: (v3 / ((v3) | 1)), F1: 4294967295, F2: (v1 | v1)}
                    v10 := []int{}
                    _ = v10
                    _ = v9
                }
                _ = fi8
            }
        }
        _ = v5
    }
    v11 := GenS3{F0: 1, F1: v3}
    for fi12 := 0; fi12 < 2; fi12++ {
        v1 += (v1 << ((65535) & 63))
        _ = fi12
    }
    var v13 int8 = 113
    var v14 float32 = 3.14159
    v15 := GenS1{F0: 897}
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v11
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
}
