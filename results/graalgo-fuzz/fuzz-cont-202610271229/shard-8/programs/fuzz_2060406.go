package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint }
type GenS2 struct { F0 uint16; F1 int32; F2 uint }

func main() {
    var v1 int16 = 42
    v2 := []uint32{}
    v3 := make(map[uint32]bool)
    var v4 uint32 = 2147483648
    v5 := make(map[uint32]string)
    var v6 uint32 = (^(v4))
    v6 += (v4 | (v4 % ((0) | 1)))
    v7 := v3[v6]
    if len(v2) > 0 {
        v2[0] = v6
    }
    v1 = (v1 | (v1 / (((v1 * v1)) | 1)))
    fmt.Println(v4, v6)
    for fi8 := 0; fi8 < 20; fi8++ {
        v9 := GenS1{F0: 32767, F1: 222}
        _ = fi8
        _ = v9
    }
    switch -2147483648 {
    case 42:
        switch 2147483647 {
        case 128:
            for fi10 := 0; fi10 < 2; fi10++ {
                _ = len(v3)
                _ = fi10
            }
        case 42:
            for fi11 := 0; fi11 < 19; fi11++ {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                var v12 int64 = 42
                _ = len(v2)
                _ = fi11
                _ = v12
            }
        }
    default:
        v13 := GenS1{F0: -1, F1: 2}
        _ = v13
    }
    var v14 int32 = -1
    v15 := GenS2{F0: 255, F1: v14, F2: 247}
    v16 := GenS2{F0: 0, F1: -2147483648, F2: 1}
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
