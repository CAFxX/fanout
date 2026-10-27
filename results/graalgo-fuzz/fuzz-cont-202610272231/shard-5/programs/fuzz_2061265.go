package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64; F2 int16 }
type GenS2 struct { F0 int64; F1 uint32; F2 int64 }

func main() {
    var v1 int64 = 9223372036854775807
    var v2 uint8 = 1
    v3 := make(map[int]uint32)
    var v4 uint = 4294967295
    v5 := make(map[uint32]bool)
    if false {
        for fi6 := 0; fi6 < 15; fi6++ {
            v1 += (v1 ^ (v1 ^ (v1 / ((-9223372036854775808) | 1))))
            v7 := GenS1{F0: 1000, F1: 42, F2: 0}
            _ = fi6
            _ = v7
        }
        v1 = 42
    }
    fmt.Println(v1)
    if true {
        switch -1 {
        case 128:
            var v8 uint = v4
            _ = v8
        case 0:
            for fi9 := 0; fi9 < 9; fi9++ {
                v4 += uint(v2)
                for fi10 := 0; fi10 < 4; fi10++ {
                    v1 -= v1
                    _ = fi10
                }
                _ = fi9
            }
        }
        v1 = (v1 - 9223372036854775807)
        v11 := []int{32767, 32767, -1}
        _ = v11
    } else {
        if false {
            fmt.Println(v1)
            v2 = v2
            if true {
                fmt.Println(v1, v4)
            }
        } else {
            v2 = (v2 * (v2 >> (((v4 >> ((255) & 63))) & 63)))
        }
    }
    if false {
        var v12 int = 9223372036854775807
        for fi13 := 0; fi13 < 13; fi13++ {
            v14 := GenS1{F0: 42, F1: 4294967295, F2: -993}
            _ = fi13
            _ = v14
        }
        _ = v12
    }
    for fi15 := 0; fi15 < 4; fi15++ {
        v5[uint32(v2)] = false
        fmt.Println(fi15)
        _ = fi15
    }
    var v16 uint64 = 18446744073709551615
    v17 := GenS2{F0: (v1 - (v1 % (((-(v1))) | 1))), F1: 2147483648, F2: v1}
    for fi18 := 0; fi18 < 10; fi18++ {
        var v19 uint16 = 1
        _ = fi18
        _ = v19
    }
    v3[127] = 100
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
