package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8 }

func main() {
    v1 := make(map[uint32]bool)
    var v2 int16 = 56
    var v3 bool = (!(-9223372036854775808 > 0))
    v4 := []string{"hello", "", "foo bar", ("a" + "a")}
    switch 128 {
    case 9223372036854775807:
        var v5 int16 = v2
        _ = v5
    case 407:
        v3 = true
    case -978:
        var v6 uint16 = 42
        _ = v6
    }
    v7 := GenS1{F0: 53, F1: 127}
    v3 = (1 == 158)
    var v8 uint = 2
    _ = len(v1)
    v9 := GenS1{F0: 127, F1: 2}
    if (255 <= 2) {
        var v10 uint = (v8 << ((191) & 63))
        v3 = true
        _ = v10
    }
    v11 := GenS1{F0: 2, F1: 127}
    var v12 float32 = 0.0
    var v13 int8 = 42
    v3 = (false || v3)
    v13 = (v13 >> ((v8) & 63))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
