package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32 }

func main() {
    v1 := []int{32767, 9223372036854775807, 32768, 1}
    v2 := make(map[uint32]uint32)
    var v3 uint8 = 255
    if (!false) {
        if ((v3 >= 73) && true) {
            v4 := GenS1{F0: 2, F1: 505}
            _ = v4
        }
        v3 -= v3
        for fi5 := 0; fi5 < 15; fi5++ {
            v6 := GenS1{F0: uint(v3), F1: 42}
            v7 := GenS1{F0: 1, F1: 2147483648}
            if ((!false) || true) {
                v3 -= (v3 / (((^(v3))) | 1))
                for fi8 := 0; fi8 < 9; fi8++ {
                    var v9 bool = ((fi5 - fi8) != (fi8 ^ fi5))
                    if (-369 > -1) {
                        _ = len(v2)
                        fmt.Println(v9)
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    }
                    _ = fi8
                    _ = v9
                }
                v10 := GenS1{F0: 0, F1: 65535}
                _ = v10
            } else {
                var v11 uint64 = 427
                _ = v11
            }
            _ = fi5
            _ = v6
            _ = v7
        }
    }
    var v12 int8 = 42
    _ = len(v2)
    v13 := v2[42]
    var v14 uint8 = v3
    v3 = 1
    v15 := GenS1{F0: 495, F1: (^(v13))}
    delete(v2, v13)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
}
