package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }

func main() {
    var v1 float32 = (3.14159 * 0.0)
    v2 := make(map[int]uint32)
    v3 := make(map[string]int)
    v4 := make(map[uint32]string)
    var v5 uint16 = 433
    var v6 int16 = 42
    if false {
        var v7 int8 = 42
        _ = v7
    } else {
        v6 = v6
    }
    for fi8 := 0; fi8 < 8; fi8++ {
        for fi9 := 0; fi9 < 4; fi9++ {
            v1 -= float32(0.1)
            var v10 uint8 = 1
            _ = fi9
            _ = v10
        }
        v1 -= 3.14159
        v11 := []int{fi8}
        _ = fi8
        _ = v11
    }
    if ("hello" != "") {
        v1 -= 0.0
    }
    v12 := GenS1{F0: 32767}
    fmt.Println(v6, v1, v6)
    v5 = (^(v5))
    for fi13 := 0; fi13 < 11; fi13++ {
        fmt.Println(v6)
        var v14 int64 = 0
        _ = fi13
        _ = v14
    }
    var v15 int8 = 127
    var v16 uint64 = 302
    var v17 int8 = -1
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
