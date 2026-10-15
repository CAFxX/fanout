package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool; F2 uint64 }

func main() {
    var v1 bool = false
    v2 := make(map[int]uint32)
    var v3 uint16 = 2
    v4 := make(map[uint32]string)
    var v5 bool = v1
    var v6 int32 = 1
    var v7 uint32 = 631
    v5 = (0 < 263)
    var v8 bool = true
    var v9 int = 127
    var v10 int16 = 1000
    v11 := GenS1{F0: 1000, F1: (127 >= (^(v9))), F2: 82}
    v10 = v10
    v1 = true
    fmt.Println(v10)
    var v12 int = (v9 - 999)
    var v13 int32 = v6
    for fi14 := 0; fi14 < 13; fi14++ {
        var v15 int8 = 62
        var v16 int64 = 9223372036854775807
        _ = fi14
        _ = v15
        _ = v16
    }
    var v17 int8 = -61
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
