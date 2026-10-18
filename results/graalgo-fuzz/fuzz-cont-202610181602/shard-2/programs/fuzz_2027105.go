package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64 }
type GenS2 struct { F0 uint16 }

func main() {
    var v1 string = "hello"
    var v2 float32 = 1.5
    var v3 uint32 = 0
    v4 := make(map[uint32]int)
    var v5 uint8 = 42
    v6 := make(map[string]string)
    v7 := GenS1{F0: 65535, F1: 605}
    var v8 int16 = -32768
    v9 := GenS1{F0: 173, F1: 115}
    v4[(v3 % ((v3) | 1))] = 128
    var v10 uint16 = 0
    if false {
        var v11 uint = 269
        delete(v6, v1)
        var v12 uint32 = (v3 | 797)
        _ = v11
        _ = v12
    }
    var v13 int32 = -2147483648
    v14 := GenS2{F0: (v10 & 0)}
    switch 688 {
    case -773:
        v1 = (v1 + "")
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
