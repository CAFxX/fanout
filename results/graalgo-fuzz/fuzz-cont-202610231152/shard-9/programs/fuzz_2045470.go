package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 bool; F1 int8; F2 uint8 }
func genFn0(p0 int, p1 uint) int8 {
    v6 := []int{(p0 * 1), p0, (p0 | (p0 * (p0 >> (((p1 << ((65535) & 63))) & 63))))}
    v6 = append(v6, -539)
    v6 = append(v6, p0)
    _ = len(v6)
    _ = p0
    _ = p1
    _ = v6
    return 100
}

func main() {
    v1 := []uint32{1, 2147483648, 42, 65535}
    v2 := []string{}
    v3 := make(map[int]string)
    var v4 uint16 = 77
    v5 := make(map[int]int)
    var v7 int64 = 0
    var v8 int16 = 179
    if (-2147483648 <= 42) {
        v8 = 42
    } else {
        v4 += (v4 % ((v4) | 1))
        var v9 int = 127
        v9 -= (v9 * 42)
        _ = v9
    }
    v10 := GenS1{F0: 32767}
    v11 := GenS1{F0: 32767}
    var v12 int8 = 1
    for fi13 := 0; fi13 < 11; fi13++ {
        var v14 int16 = 1
        var v15 float64 = (float64(1e10) / 1e10)
        v16 := GenS2{F0: false, F1: (v12 % (((v12 * -55)) | 1)), F2: 0}
        _ = fi13
        _ = v14
        _ = v15
        _ = v16
    }
    v8 += (v8 / (((v8 % (((v8 >> ((2) & 63))) | 1))) | 1))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
}
