package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16; F2 int }
func genFn0(p0 int32, p1 uint64) uint16 {
    fmt.Println(p0, p1)
    var v6 uint64 = 722
    _ = p0
    _ = p1
    _ = v6
    return 2
}

func main() {
    var v1 int64 = 42
    var v2 int = -2147483648
    v3 := make(map[int]int)
    var v4 uint8 = 128
    var v5 int16 = 42
    v1 -= -1
    for fi7 := 0; fi7 < 13; fi7++ {
        v8 := GenS1{F0: 255, F1: v5, F2: v2}
        var v9 uint32 = 1
        var v10 int32 = 1
        _ = fi7
        _ = v10
        _ = v8
        _ = v9
    }
    v4 += v4
    var v11 int16 = (^(v5))
    var v12 uint32 = 142
    v12 -= v12
    v2 -= 0
    v3[(v2 * (v2 ^ v2))] = v2
    var v13 int = (v2 | 32768)
    var v14 int = (v13 % (((v13 * v13)) | 1))
    var v15 uint64 = 835
    v16 := GenS1{F0: v4, F1: v11, F2: v13}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
