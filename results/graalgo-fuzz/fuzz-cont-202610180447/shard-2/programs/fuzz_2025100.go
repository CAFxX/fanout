package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int64; F1 int64; F2 uint32 }

func main() {
    v1 := make(map[uint32]int)
    var v2 int64 = 423
    v3 := make(map[int]bool)
    var v4 bool = false
    v5 := []int{32768, 42}
    var v6 string = (("a" + "a") + ("a" + "hello"))
    for fi7 := 0; fi7 < 13; fi7++ {
        v4 = (0 != 1)
        v8 := GenS1{F0: (((1.5 <= -2.25) || (fi7 != fi7)) || (!(2147483647 != 1)))}
        v6 = "hello"
        _ = fi7
        _ = v8
    }
    var v9 string = (v6 + v6)
    var v10 uint = 65535
    v9 = v6
    var v11 int64 = (v2 << (((v10 << (((v10 + (^(v10)))) & 63))) & 63))
    var v12 int64 = (v2 * v11)
    v5 = append(v5, -1)
    v13 := GenS2{F0: 9223372036854775807, F1: (v12 & (v2 ^ v12)), F2: 4294967295}
    _ = len(v5)
    if (4294967295 == 935) {
        if v4 {
            var v14 float64 = float64(3.14159)
            _ = v14
        } else {
            v3[-785] = (42 > int8(v12))
        }
        var v15 uint64 = 289
        v10 += (^(v10))
        _ = v15
    } else {
        fmt.Println(v9, v2, v11)
        var v16 uint32 = 42
        fmt.Println(v4, v11)
        _ = v16
    }
    v11 = 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
