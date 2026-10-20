package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint8 }
func genFn0(p0 int64, p1 string, p2 uint16) int64 {
    p2 += 65535
    p2 = p2
    var v4 int = -9223372036854775808
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return (p0 & -9223372036854775808)
}

func main() {
    v1 := make(map[int]string)
    var v2 bool = (false || (-282 == 1))
    var v3 int64 = -9223372036854775808
    v5 := GenS1{F0: -1}
    fmt.Println(v2, v2)
    var v6 uint32 = 2147483647
    for fi7 := 0; fi7 < 9; fi7++ {
        fmt.Println(v2, v6, v3)
        var v8 uint8 = 246
        var v9 int = (fi7 >> ((uint(v6)) & 63))
        _ = fi7
        _ = v8
        _ = v9
    }
    _ = len(v1)
    v10 := GenS1{F0: 0}
    var v11 uint16 = 2
    var v12 float64 = 1e10
    v12 -= 0.1
    for fi13 := 0; fi13 < 9; fi13++ {
        var v14 uint = 2
        v15 := GenS2{F0: 42}
        _ = fi13
        _ = v14
        _ = v15
    }
    if (42 != -32768) {
        fmt.Println(v6, v11)
    } else {
        v2 = (714 > 0)
        var v16 bool = false
        _ = v16
    }
    var v17 float32 = 1e10
    var v18 float32 = 0.1
    v6 = (v6 % ((2147483648) | 1))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v5
    _ = v6
}
