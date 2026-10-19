package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint8; F1 bool; F2 int32 }
type GenS3 struct { F0 int16 }
func genFn0(p0 int64, p1 int64) int32 {
    v7 := []int{-2147483648, 32768}
    p1 = p1
    switch 0 {
    case -158:
        v7 = append(v7, -1)
    case -164:
        var v8 int = -2147483648
        _ = v8
    case 127:
        p1 = p1
    }
    p0 -= -414
    p1 += p1
    _ = p0
    _ = p1
    _ = v7
    return 42
}

func main() {
    v1 := make(map[int]uint32)
    var v2 float32 = (0.1 - float32((1e10 / 1.0)))
    var v3 int32 = -327
    var v4 uint64 = 42
    var v5 bool = true
    v6 := make(map[string]int)
    delete(v1, 2147483647)
    fmt.Println(v3)
    v9 := GenS1{F0: 255}
    v10 := GenS2{F0: 165, F1: (2147483648 >= 2), F2: -1}
    var v11 float32 = (0.1 + (0.0 * v2))
    v12 := GenS3{F0: 67}
    for fi13 := 0; fi13 < 11; fi13++ {
        if v5 {
            v14 := []int{2147483647, -996}
            var v15 uint64 = v4
            _ = v14
            _ = v15
        } else {
            var v16 string = "foo bar"
            v17 := GenS2{F0: 42, F1: (-62 != int64(fi13)), F2: (v3 << ((65535) & 63))}
            _ = v16
            _ = v17
        }
        _ = fi13
    }
    delete(v1, 42)
    if (1 <= -1) {
        switch "hello" {
        case "foo bar":
            var v18 int16 = 679
            _ = v18
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
