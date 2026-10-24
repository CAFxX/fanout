package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8 }

func main() {
    var v1 uint32 = 0
    var v2 int64 = 42
    var v3 int64 = (v2 % (((-(v2))) | 1))
    v4 := make(map[int]string)
    v5 := []uint32{v1, (^(v1))}
    v6 := make(map[uint32]string)
    v7 := v6[(^(v1))]
    var v8 int8 = 0
    v9 := GenS1{F0: 986, F1: 42}
    if false {
        fmt.Println(v3)
        v10 := v6[(v1 * 65535)]
        _ = v10
    }
    for fi11 := 0; fi11 < 4; fi11++ {
        v7 = v7
        _ = fi11
    }
    v3 -= (^(v2))
    var v12 int8 = (v8 - 1)
    v13 := GenS1{F0: 255, F1: 128}
    v8 = -1
    for fi14 := 0; fi14 < 8; fi14++ {
        fmt.Println(fi14)
        _ = fi14
    }
    var v15 int16 = 32767
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v13
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
