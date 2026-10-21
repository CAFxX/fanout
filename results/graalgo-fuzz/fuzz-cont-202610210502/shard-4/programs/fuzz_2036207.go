package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }

func main() {
    var v1 int8 = 1
    var v2 uint = 255
    var v3 uint32 = 42
    var v4 int32 = 1
    var v5 uint16 = 255
    v3 += 550
    var v6 uint8 = 128
    var v7 int64 = 1
    for fi8 := 0; fi8 < 19; fi8++ {
        v3 += (v3 << ((v2) & 63))
        _ = fi8
    }
    v9 := []int{42}
    if len(v9) > 0 {
        v9[0] = 9223372036854775807
    }
    v6 -= (^(v6))
    v5 = v5
    fmt.Println(v1, v4)
    var v10 uint8 = v6
    v6 = v6
    v11 := GenS1{F0: v4}
    fmt.Println(v2, v3, v4)
    var v12 int = -448
    if ((v5 - v5) > uint16(v2)) {
        v2 = (v2 | (v2 & v2))
        for fi13 := 0; fi13 < 5; fi13++ {
            v9 = append(v9, (-(fi13)))
            _ = fi13
        }
    }
    v3 = (v3 | 261)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v9:", len(v9), v9)
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
    _ = v9
}
