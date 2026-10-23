package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }

func main() {
    v1 := []uint32{1, 1, 0, 42}
    var v2 int = 9223372036854775807
    var v3 int8 = 36
    v4 := make(map[int]uint32)
    var v5 bool = (false || false)
    v6 := make(map[int]uint32)
    var v7 string = ("\n\t" + ("hello" + ""))
    _ = len(v4)
    var v8 uint64 = 4294967295
    var v9 uint = 0
    for fi10 := 0; fi10 < 8; fi10++ {
        var v11 uint8 = 58
        _ = fi10
        _ = v11
    }
    var v12 int8 = v3
    v8 += (v8 | 844)
    var v13 uint16 = 1
    v2 = 42
    for fi14 := 0; fi14 < 4; fi14++ {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = fi14
    }
    fmt.Println(v8, v7, v13)
    v9 += (v9 << ((v9) & 63))
    var v15 int64 = -1
    fmt.Println(v9, v7)
    v5 = v5
    for fi16 := 0; fi16 < 14; fi16++ {
        var v17 uint64 = 2
        _ = fi16
        _ = v17
    }
    v18 := GenS1{F0: int8(v8)}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v13
    _ = v15
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
