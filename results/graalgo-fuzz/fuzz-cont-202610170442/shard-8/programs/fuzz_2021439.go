package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16; F2 uint }
type GenS2 struct { F0 uint32; F1 int8 }
type GenS3 struct { F0 uint64; F1 bool; F2 bool }
type GenS4 struct { F0 int32; F1 int8; F2 uint }
type GenS5 struct { F0 int8; F1 uint }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int = 42
    v3 := make(map[string]uint32)
    v4 := make(map[uint32]string)
    var v5 int8 = 72
    var v6 int = (v2 & (v2 * 32768))
    v7 := GenS1{F0: 2, F1: 0, F2: 42}
    v8 := GenS1{F0: 42, F1: -928, F2: 1}
    switch "hello" {
    case "hello":
        v9 := GenS1{F0: 65535, F1: 42, F2: 408}
        _ = v9
    case "foo bar":
        v10 := GenS2{F0: 246, F1: v5}
        _ = v10
    case "a":
        v6 += (v2 & v2)
    }
    for fi11 := 0; fi11 < 19; fi11++ {
        v6 = fi11
        v12 := GenS1{F0: 338, F1: -32768, F2: 208}
        v4[2147483647] = "\n\t"
        _ = fi11
        _ = v12
    }
    v13 := GenS3{F0: 2, F1: true, F2: true}
    for fi14 := 0; fi14 < 19; fi14++ {
        fmt.Println(v2, v6)
        _ = fi14
    }
    v15 := GenS4{F0: 676, F1: int8(v6), F2: 42}
    for fi16 := 0; fi16 < 1; fi16++ {
        v2 = (fi16 & v6)
        _ = fi16
    }
    v17 := GenS5{F0: v5, F1: 627}
    if ((^(v5)) > (v5 / ((v5) | 1))) {
        var v18 int16 = 0
        v19 := GenS2{F0: 957, F1: (v5 << ((1) & 63))}
        delete(v1, 2)
        _ = v18
        _ = v19
    }
    var v20 float32 = 3.14159
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v15
    _ = v17
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
