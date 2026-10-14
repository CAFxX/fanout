package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8; F2 string }
type GenS2 struct { F0 int32; F1 uint64 }

func main() {
    v1 := []int{1, 311, 247}
    v2 := make(map[uint32]uint32)
    v3 := make(map[string]bool)
    var v4 int64 = -1
    var v5 string = ""
    v6 := make(map[int]bool)
    fmt.Println(v4, v4)
    v7 := GenS1{F0: 42, F1: 42, F2: v5}
    for fi8 := 0; fi8 < 14; fi8++ {
        if true {
            v9 := GenS1{F0: 0, F1: 127, F2: v5}
            _ = v9
        }
        _ = fi8
    }
    for fi10 := 0; fi10 < 15; fi10++ {
        var v11 string = "foo bar"
        v12 := GenS1{F0: 573, F1: -1, F2: "foo bar"}
        fmt.Println(v5, v11, v11)
        _ = fi10
        _ = v11
        _ = v12
    }
    if true {
        fmt.Println(v5)
    } else {
        v13 := GenS2{F0: 42, F1: 4294967295}
        var v14 uint64 = 18446744073709551615
        fmt.Println(v4, v4)
        _ = v13
        _ = v14
    }
    v15 := GenS1{F0: 1, F1: 42, F2: "a"}
    v5 = "a"
    var v16 uint64 = 18446744073709551615
    if true {
        for fi17 := 0; fi17 < 5; fi17++ {
            for fi18 := 0; fi18 < 10; fi18++ {
                var v19 float32 = 1.5
                var v20 int32 = -1
                _ = fi18
                _ = v19
                _ = v20
            }
            _ = fi17
        }
    }
    v6[42] = false
    v21 := GenS2{F0: int32(v4), F1: 2}
    v16 = (v16 >> ((255) & 63))
    var v22 string = v5
    for fi23 := 0; fi23 < 18; fi23++ {
        v5 = v5
        var v24 int32 = 2147483647
        _ = fi23
        _ = v24
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
