package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint16 }

func main() {
    var v1 int64 = -942
    v2 := make(map[int]int)
    var v3 bool = (!false)
    var v4 int8 = -1
    v5 := []uint32{}
    var v6 uint32 = 42
    var v7 int32 = 2147483647
    _ = len(v5)
    var v8 uint = 672
    if v3 {
        if v3 {
            delete(v2, int(v4))
        } else {
            _ = len(v5)
            v7 = v7
            v9 := GenS1{F0: 2}
            _ = v9
        }
    }
    var v10 uint8 = uint8(v6)
    var v11 int = 207
    v7 = int32(v6)
    v5 = append(v5, 240)
    v12 := GenS1{F0: 255}
    _ = len(v2)
    for fi13 := 0; fi13 < 17; fi13++ {
        for fi14 := 0; fi14 < 8; fi14++ {
            if false {
                _ = len(v5)
                var v15 uint64 = 0
                _ = v15
            }
            v16 := GenS2{F0: v8}
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            v17 := GenS1{F0: 102}
            v18 := GenS3{F0: 1}
            _ = fi14
            _ = v16
            _ = v17
            _ = v18
        }
        v1 = 42
        v5 = append(v5, (v6 * (v6 ^ 962)))
        _ = fi13
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
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
    _ = v6
    _ = v7
    _ = v8
}
