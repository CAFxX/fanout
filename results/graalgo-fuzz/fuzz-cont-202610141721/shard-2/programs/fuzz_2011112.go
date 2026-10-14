package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }

func main() {
    var v1 uint8 = 1
    v2 := make(map[uint32]bool)
    var v3 uint32 = 65535
    var v4 uint = 0
    var v5 uint8 = (^(v1))
    v6 := []string{"foo bar"}
    var v7 bool = false
    var v8 uint64 = 576
    for fi9 := 0; fi9 < 16; fi9++ {
        fmt.Println(v3)
        var v10 float32 = 0.1
        _ = fi9
        _ = v10
    }
    v4 += (v4 << (((v4 | (v4 >> ((4294967295) & 63)))) & 63))
    switch "a" {
    case "hello":
        delete(v2, 2147483648)
    }
    var v11 uint = v4
    v12 := GenS1{F0: uint64(v11)}
    var v13 bool = ((!(v3 >= uint32(v8))) || (9223372036854775807 <= -1))
    for fi14 := 0; fi14 < 19; fi14++ {
        var v15 int = (fi14 & 337)
        var v16 bool = true
        _ = fi14
        _ = v15
        _ = v16
    }
    var v17 uint64 = v8
    if (uint16(v3) < uint16(v8)) {
        v5 -= 128
    } else {
        v4 -= v4
        v4 = v11
        _ = len(v2)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
