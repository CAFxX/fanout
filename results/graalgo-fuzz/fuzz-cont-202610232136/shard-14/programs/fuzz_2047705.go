package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint }

func main() {
    v1 := make(map[string]bool)
    var v2 uint8 = 42
    v3 := make(map[uint32]bool)
    v4 := make(map[uint32]uint32)
    v5 := GenS1{F0: 42, F1: 18446744073709551615}
    v2 = v2
    v6 := GenS1{F0: 127, F1: 2}
    var v7 int16 = 32767
    v8 := []int{32768, 4294967295, 1}
    for fi9 := 0; fi9 < 6; fi9++ {
        for fi10 := 0; fi10 < 8; fi10++ {
            var v11 int8 = 81
            v2 = v2
            v7 -= v7
            v7 = v7
            _ = fi10
            _ = v11
        }
        fmt.Println(v7, fi9)
        v7 -= 1
        _ = fi9
    }
    switch "hello" {
    case "":
        v12 := GenS1{F0: -731, F1: 2}
        _ = v12
    case "hello":
        for fi13 := 0; fi13 < 1; fi13++ {
            var v14 uint = 4294967295
            _ = fi13
            _ = v14
        }
    case "a":
        v2 = (v2 >> ((4294967295) & 63))
    }
    v2 += 0
    var v15 int32 = 2147483647
    v7 = 1000
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
