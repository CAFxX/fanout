package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint8; F1 uint16; F2 int16 }

func main() {
    v1 := make(map[string]int)
    v2 := make(map[uint32]bool)
    var v3 int64 = 9223372036854775807
    v4 := make(map[uint32]string)
    fmt.Println(v3)
    delete(v2, 5)
    var v5 int16 = 0
    var v6 int32 = int32(v5)
    v7 := GenS1{F0: 2}
    v8 := GenS1{F0: 1}
    switch 1 {
    case -2147483648:
        if true {
            v9 := GenS2{F0: 0, F1: 568, F2: v5}
            _ = v9
        } else {
            var v10 bool = (-302 > 421)
            _ = v10
        }
    case 0:
        v11 := GenS1{F0: 42}
        _ = v11
    default:
        v1["a"] = 32767
    }
    var v12 float64 = float64(0.1)
    v13 := v1["hello"]
    v5 = -258
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
