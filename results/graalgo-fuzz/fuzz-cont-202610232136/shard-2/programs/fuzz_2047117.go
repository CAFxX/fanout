package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int; F2 int64 }
type GenS2 struct { F0 bool; F1 int8; F2 string }

func main() {
    var v1 int16 = -1
    var v2 int16 = (v1 << ((710) & 63))
    var v3 int32 = 248
    var v4 int16 = 0
    v5 := GenS1{F0: 4294967295, F1: 32768, F2: 0}
    v2 = -598
    v6 := GenS1{F0: 445, F1: int(v2), F2: -1}
    v2 = v2
    v3 += -1
    v7 := GenS2{F0: (585 <= 2), F1: -128, F2: "foo bar"}
    v8 := GenS2{F0: ((3.14159 < 0.1) && false), F1: 78, F2: "a"}
    var v9 float32 = 3.14159
    var v10 uint64 = 362
    var v11 uint16 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
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
