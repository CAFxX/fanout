package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8 }
type GenS2 struct { F0 int64; F1 int }

func main() {
    v1 := make(map[string]string)
    v2 := make(map[int]uint32)
    v3 := make(map[int]string)
    v4 := make(map[string]int)
    var v5 float32 = 3.14159
    v6 := make(map[uint32]bool)
    fmt.Println(v5)
    v7 := GenS1{F0: 135, F1: 209}
    var v8 int64 = 1
    var v9 int8 = -1
    v9 -= (v9 * (v9 & v9))
    var v10 uint = 881
    fmt.Println(v9)
    v11 := GenS2{F0: (v8 | (v8 % ((v8) | 1))), F1: 1}
    v12 := []int{263, 546, 9223372036854775807}
    switch 2147483647 {
    case -2147483648:
        var v13 int16 = 32767
        _ = v13
    case 32768:
        var v14 uint64 = 392
        _ = v14
    default:
        var v15 uint8 = 255
        _ = v15
    }
    switch 32768 {
    case 983:
        for fi16 := 0; fi16 < 14; fi16++ {
            fmt.Println(v10, v8)
            v9 += (v9 >> ((2) & 63))
            var v17 uint = (v10 ^ v10)
            _ = fi16
            _ = v17
        }
    case 4294967295:
        for fi18 := 0; fi18 < 5; fi18++ {
            v19 := GenS1{F0: 42, F1: 8}
            var v20 int16 = -497
            var v21 int32 = -2147483648
            _ = fi18
            _ = v19
            _ = v20
            _ = v21
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
