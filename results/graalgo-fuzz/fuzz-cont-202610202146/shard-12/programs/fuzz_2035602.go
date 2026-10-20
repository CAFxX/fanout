package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint16; F1 uint16; F2 uint8 }

func main() {
    var v1 float64 = 0.0
    var v2 int64 = 0
    v3 := make(map[string]uint32)
    v4 := make(map[string]bool)
    var v5 uint16 = 583
    var v6 uint8 = 84
    var v7 bool = true
    v8 := GenS1{F0: v6}
    v9 := GenS1{F0: (v6 << ((uint(v5)) & 63))}
    v10 := GenS2{F0: uint32(v5)}
    var v11 uint8 = v6
    v6 -= 187
    var v12 int16 = -32768
    var v13 int8 = 42
    v14 := GenS3{F0: (v5 * (v5 | (v5 + v5))), F1: (v5 - 255), F2: 28}
    for fi15 := 0; fi15 < 10; fi15++ {
        var v16 uint16 = v5
        v11 -= (v11 - v11)
        var v17 uint32 = 2
        _ = fi15
        _ = v16
        _ = v17
    }
    var v18 string = "foo bar"
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
