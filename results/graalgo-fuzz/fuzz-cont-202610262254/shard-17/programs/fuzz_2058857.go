package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64 }

func main() {
    var v1 uint64 = 2
    var v2 uint64 = v1
    var v3 int = 2147483647
    var v4 int = -1
    v5 := []uint32{1}
    v4 = -409
    fmt.Println(v4)
    var v6 bool = false
    v7 := GenS1{F0: 194, F1: 1}
    fmt.Println(v3)
    _ = len(v5)
    v1 -= (v1 << ((513) & 63))
    var v8 string = ""
    _ = len(v5)
    var v9 int8 = -1
    if ((uint16(v2) != 1) || false) {
        if ((v9 != v9) || (1 < -1)) {
            var v10 bool = v6
            _ = v10
        }
    }
    v8 = v8
    v8 = ((v8 + ("hello" + "foo bar")) + (v8 + ""))
    var v11 int32 = 0
    v2 += 365
    for fi12 := 0; fi12 < 9; fi12++ {
        for fi13 := 0; fi13 < 19; fi13++ {
            var v14 uint64 = (v2 / ((1) | 1))
            _ = fi13
            _ = v14
        }
        var v15 uint16 = 285
        _ = fi12
        _ = v15
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
