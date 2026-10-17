package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64; F2 uint32 }
type GenS2 struct { F0 uint8; F1 int16; F2 int16 }

func main() {
    var v1 uint64 = 18446744073709551615
    v2 := make(map[uint32]uint32)
    var v3 uint32 = 4294967295
    var v4 string = "foo bar"
    var v5 int8 = -78
    v1 -= v1
    var v6 int8 = v5
    fmt.Println(v3, v3)
    v4 = "\n\t"
    var v7 uint64 = v1
    if true {
        var v8 uint8 = 128
        _ = v8
    }
    var v9 int32 = -1
    v10 := GenS1{F0: 42, F1: -9223372036854775808, F2: (v3 - v3)}
    v11 := GenS1{F0: v3, F1: 1, F2: (v3 >> ((143) & 63))}
    var v12 bool = true
    for fi13 := 0; fi13 < 19; fi13++ {
        v14 := GenS1{F0: v3, F1: 42, F2: v3}
        _ = fi13
        _ = v14
    }
    for fi15 := 0; fi15 < 6; fi15++ {
        var v16 int64 = -9223372036854775808
        v16 -= (^(v16))
        v17 := GenS1{F0: 65535, F1: v16, F2: (v3 >> ((255) & 63))}
        _ = fi15
        _ = v16
        _ = v17
    }
    var v18 float32 = 1e10
    var v19 int16 = 42
    var v20 uint16 = 32767
    v12 = ((v9 << ((255) & 63)) > v9)
    fmt.Println(v5, v1, v20)
    v21 := GenS2{F0: 128, F1: v19, F2: (v19 >> ((0) & 63))}
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
