package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint }
type GenS2 struct { F0 int8; F1 int64; F2 int16 }
func genFn0(p0 uint8, p1 int64, p2 bool) uint64 {
    p2 = p2
    p1 += p1
    _ = p0
    _ = p1
    _ = p2
    return 620
}
func genFn1(p0 uint64) int16 {
    switch "\n\t" {
    case "hello":
        var v4 int = -2147483648
        _ = v4
    }
    v5 := GenS1{F0: 1, F1: 65535}
    if ((p0 - p0) <= 42) {
        v6 := []int{9223372036854775807}
        for fi7 := 0; fi7 < 8; fi7++ {
            var v8 int8 = 1
            var v9 int8 = v8
            _ = fi7
            _ = v8
            _ = v9
        }
        var v10 int16 = 1
        _ = v10
        _ = v6
    } else {
        var v11 int = 0
        _ = v11
    }
    var v12 int8 = -108
    _ = p0
    _ = v12
    _ = v5
    return -845
}

func main() {
    v1 := make(map[int]bool)
    var v2 int32 = 0
    v3 := make(map[int]int)
    var v13 int16 = 1
    var v14 float32 = (0.0 * (1e10 * 0.1))
    v3[32767] = 42
    v15 := GenS2{F0: -1, F1: -1, F2: (v13 - (v13 * (v13 ^ (^(v13)))))}
    fmt.Println(v2, v2, v2)
    v1[42] = (42 <= 9223372036854775807)
    var v16 int8 = 0
    switch "hello" {
    case "a":
        v17 := GenS1{F0: 0, F1: 65535}
        _ = v17
    default:
        fmt.Println(v16, v2, v16)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
}
