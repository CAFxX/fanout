package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint16 }

func main() {
    var v1 int16 = 42
    v2 := make(map[string]int)
    var v3 int16 = (v1 << ((42) & 63))
    var v4 int64 = 322
    _ = len(v2)
    for fi5 := 0; fi5 < 8; fi5++ {
        v1 = (v3 ^ v3)
        v6 := []int{-192, fi5, (fi5 ^ (fi5 % ((fi5) | 1)))}
        _ = fi5
        _ = v6
    }
    for fi7 := 0; fi7 < 6; fi7++ {
        if true {
            var v8 uint64 = 1
            v9 := GenS1{F0: 18446744073709551615}
            v4 += 0
            _ = v8
            _ = v9
        }
        v1 -= 0
        var v10 int8 = 1
        _ = fi7
        _ = v10
    }
    v1 = v3
    v11 := []int{32768, -1, 89}
    var v12 int32 = 2147483647
    var v13 uint16 = 0
    var v14 float64 = 3.14159
    fmt.Println(v4, v14, v1)
    v15 := GenS2{F0: 42}
    var v16 int32 = -2147483648
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
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
}
