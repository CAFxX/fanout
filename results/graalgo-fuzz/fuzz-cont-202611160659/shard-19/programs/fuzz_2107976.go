package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32 }
func genFn0(p0 uint64) uint8 {
    p0 = p0
    var v7 int = 32768
    for fi8 := 0; fi8 < 20; fi8++ {
        var v9 uint8 = 128
        v7 = v7
        _ = fi8
        _ = v9
    }
    fmt.Println(p0)
    _ = p0
    _ = v7
    return 127
}

func main() {
    v1 := []int{0}
    var v2 uint64 = 0
    v3 := make(map[uint32]int)
    var v4 int = -1
    var v5 uint16 = 610
    var v6 bool = (true || (32767 >= (v5 | 65535)))
    delete(v3, 65535)
    var v10 uint16 = v5
    fmt.Println(v4, v5, v2)
    var v11 int8 = 42
    for fi12 := 0; fi12 < 13; fi12++ {
        if true {
            v13 := v3[65535]
            _ = v13
        }
        var v14 int64 = 0
        v1 = append(v1, fi12)
        _ = fi12
        _ = v14
    }
    if len(v1) > 0 {
        v1[0] = 127
    }
    var v15 float32 = 1e10
    var v16 int8 = (v11 % ((v11) | 1))
    var v17 float32 = float32(3.14159)
    var v18 uint64 = v2
    if v6 {
        v19 := GenS1{F0: -880, F1: 4294967295}
        _ = v19
    } else {
        v16 = (v11 ^ v11)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
