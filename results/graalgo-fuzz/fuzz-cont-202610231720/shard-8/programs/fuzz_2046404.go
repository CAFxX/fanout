package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint64; F1 uint; F2 uint32 }

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[string]uint32)
    var v3 uint8 = 131
    var v4 uint64 = 2
    v5 := make(map[string]string)
    v6 := GenS1{F0: 100}
    var v7 int16 = 32767
    var v8 float32 = 1e10
    v9 := GenS1{F0: -128}
    var v10 int16 = (v7 >> ((65535) & 63))
    _ = len(v1)
    v11 := GenS1{F0: 113}
    var v12 string = "foo bar"
    v8 += ((float32(1.5) / 1.0) / (v8 + ((v8 - v8) / 1.0)))
    fmt.Println(v4)
    var v13 int16 = (v7 & 32767)
    switch "a" {
    case "\n\t":
        fmt.Println(v8, v13, v13)
    case "":
        v14 := GenS2{F0: (^(v4)), F1: 1, F2: 2147483647}
        _ = v14
    }
    var v15 uint16 = uint16(v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
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
