package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8; F2 string }

func main() {
    var v1 bool = (662 <= 65535)
    v2 := make(map[uint32]bool)
    v3 := make(map[int]uint32)
    var v4 uint16 = 536
    var v5 bool = v1
    var v6 uint16 = (v4 >> ((1) & 63))
    if (!(94 == 127)) {
        v7 := GenS1{F0: 0, F1: uint8(v6), F2: ""}
        _ = v7
    }
    v5 = true
    if (-614 <= 9223372036854775807) {
        v1 = (0 <= 2)
    }
    if (65535 >= (v6 * 0)) {
        fmt.Println(v5, v1)
        var v8 uint8 = 177
        _ = v8
    } else {
        v5 = (!v1)
        v9 := []int{-401, 4294967295}
        _ = v9
    }
    for fi10 := 0; fi10 < 6; fi10++ {
        v11 := []int{}
        v6 += 591
        _ = fi10
        _ = v11
    }
    v4 = (v6 & v6)
    _ = len(v2)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
