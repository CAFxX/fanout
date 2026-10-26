package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int }
type GenS2 struct { F0 int8; F1 uint32 }
func genFn0() int64 {
    var v4 int8 = 42
    v5 := []int{42, 32768, 1, 128}
    switch "foo bar" {
    case "a":
        v6 := GenS1{F0: 2147483647, F1: 128}
        _ = v6
    default:
        v5 = append(v5, 127)
    }
    _ = v4
    _ = v5
    return -9223372036854775808
}

func main() {
    var v1 bool = (2147483648 == 4294967295)
    var v2 int = 32767
    v3 := []int{v2, 9223372036854775807, v2, v2}
    if true {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    v7 := GenS1{F0: 2147483647, F1: v2}
    var v8 uint32 = 151
    var v9 uint64 = 0
    v1 = (480 < 489)
    _ = len(v3)
    if len(v3) > 0 {
        v3[0] = (v2 / (((v2 - v2)) | 1))
    }
    switch v2 {
    case -655:
        if len(v3) > 0 {
            v3[0] = (v2 << ((255) & 63))
        }
    case -415:
        var v10 float32 = 1e10
        _ = v10
    default:
        for fi11 := 0; fi11 < 1; fi11++ {
            fmt.Println(v1, v1)
            _ = fi11
        }
    }
    var v12 uint8 = 81
    for fi13 := 0; fi13 < 11; fi13++ {
        var v14 uint = uint(v12)
        v3 = append(v3, int(v9))
        if len(v3) > 0 {
            v3[0] = (fi13 - v2)
        }
        _ = fi13
        _ = v14
    }
    switch v2 {
    case -814:
        v15 := GenS2{F0: -1, F1: 4294967295}
        _ = v15
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v7
    _ = v8
    _ = v9
}
