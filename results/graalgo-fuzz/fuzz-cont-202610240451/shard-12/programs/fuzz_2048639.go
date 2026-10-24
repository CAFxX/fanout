package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8 }
type GenS2 struct { F0 uint16; F1 uint8; F2 int16 }

func main() {
    var v1 uint16 = 1
    v2 := []uint32{65535, 1, 42}
    v3 := []string{"\n\t", ""}
    var v4 uint = 4294967295
    fmt.Println(v1)
    if (false || false) {
        v1 -= (v1 << (((^(v4))) & 63))
    }
    v1 += (v1 * 32767)
    v1 = v1
    for fi5 := 0; fi5 < 19; fi5++ {
        _ = len(v2)
        v6 := GenS1{F0: uint64(v4), F1: -1}
        _ = fi5
        _ = v6
    }
    v1 -= 255
    var v7 bool = false
    v8 := GenS1{F0: 806, F1: int8(v4)}
    var v9 int64 = 9223372036854775807
    var v10 int16 = 431
    v10 = (v10 << (((v4 ^ v4)) & 63))
    v11 := GenS2{F0: 32767, F1: 255, F2: (v10 - -1)}
    var v12 uint16 = v1
    for fi13 := 0; fi13 < 1; fi13++ {
        fmt.Println(v9, v4)
        v1 += 2
        _ = fi13
    }
    switch "a" {
    case "foo bar":
        var v14 uint8 = 2
        _ = v14
    case "":
        fmt.Println(v10, v12)
    case "hello":
        fmt.Println(v9)
    }
    var v15 uint32 = 2147483648
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
