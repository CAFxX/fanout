package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[uint32]string)
    v3 := []int{32767, 4294967295, 882, 405}
    for fi4 := 0; fi4 < 8; fi4++ {
        var v5 int64 = 1
        _ = fi4
        _ = v5
    }
    v3 = append(v3, 128)
    var v6 int16 = -846
    var v7 int64 = -9223372036854775808
    v7 += -1
    for fi8 := 0; fi8 < 2; fi8++ {
        v9 := GenS1{F0: 920, F1: fi8}
        if true {
            fmt.Println(v7)
            v2[138] = "\n\t"
        }
        for fi10 := 0; fi10 < 12; fi10++ {
            v6 += v6
            _ = fi10
        }
        _ = fi8
        _ = v9
    }
    for fi11 := 0; fi11 < 17; fi11++ {
        v12 := GenS1{F0: 0, F1: fi11}
        _ = fi11
        _ = v12
    }
    if false {
        for fi13 := 0; fi13 < 6; fi13++ {
            v7 += (v7 - v7)
            _ = fi13
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v6
    _ = v7
}
