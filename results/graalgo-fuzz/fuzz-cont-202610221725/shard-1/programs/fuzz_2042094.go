package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[uint32]uint32)
    v3 := []string{("foo bar" + "foo bar")}
    var v4 int64 = -931
    v5 := []uint32{}
    fmt.Println(v4)
    v6 := GenS1{F0: 255}
    v4 += (v4 << ((150) & 63))
    if false {
        v4 -= v4
    }
    switch 9223372036854775807 {
    case 127:
        var v7 int16 = 590
        _ = v7
    case 32767:
        v4 -= -1
    case 0:
        v8 := GenS1{F0: 42}
        _ = v8
    }
    _ = len(v3)
    v4 -= v4
    switch 188 {
    case -1:
        v9 := GenS1{F0: 255}
        _ = v9
    case 758:
        var v10 int8 = 42
        _ = v10
    default:
        v4 = v4
    }
    _ = len(v2)
    if len(v5) > 0 {
        v5[0] = 2147483647
    }
    var v11 string = "\n\t"
    fmt.Println(v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
