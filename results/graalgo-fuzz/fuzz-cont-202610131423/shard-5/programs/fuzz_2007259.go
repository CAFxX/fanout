package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }

func main() {
    var v1 int16 = 0
    v2 := make(map[string]int)
    var v3 int16 = (^(v1))
    v4 := make(map[string]int)
    var v5 uint32 = 454
    var v6 uint64 = 42
    var v7 int32 = 2147483647
    var v8 uint8 = uint8(v6)
    fmt.Println(v5)
    v9 := GenS1{F0: v6}
    if (!((^(v8)) != (v8 >> ((18446744073709551615) & 63)))) {
        var v10 int = 4294967295
        v10 -= (v10 / (((v10 + v10)) | 1))
        v11 := GenS1{F0: v6}
        _ = v10
        _ = v11
    }
    v12 := GenS1{F0: v6}
    var v13 int8 = -39
    v3 -= (v3 | v1)
    switch -9223372036854775808 {
    case 32768:
        _ = len(v4)
    default:
        v4["hello"] = -9223372036854775808
    }
    switch "" {
    case "\n\t":
        var v14 uint64 = (v6 >> ((339) & 63))
        _ = v14
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
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
    _ = v9
}
