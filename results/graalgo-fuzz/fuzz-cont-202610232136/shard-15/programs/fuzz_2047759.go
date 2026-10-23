package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[string]int)
    v3 := []uint32{}
    var v4 int = 127
    v5 := make(map[int]uint32)
    var v6 int16 = 793
    if len(v3) > 0 {
        v3[0] = 2
    }
    var v7 uint8 = 127
    v7 = 5
    if false {
        var v8 bool = (!(true || false))
        _ = v8
    }
    var v9 bool = false
    if false {
        v10 := GenS1{F0: -282, F1: "foo bar"}
        _ = v10
    }
    switch v4 {
    case 605:
        fmt.Println(v9, v6, v7)
    case 128:
        v11 := GenS1{F0: 675, F1: "a"}
        _ = v11
    }
    v12 := GenS1{F0: -1, F1: "hello"}
    var v13 int32 = 42
    v9 = false
    v14 := GenS1{F0: 9223372036854775807, F1: "hello"}
    for fi15 := 0; fi15 < 14; fi15++ {
        if v9 {
            var v16 uint16 = 958
            var v17 uint32 = 2147483647
            _ = v16
            _ = v17
        }
        _ = fi15
    }
    fmt.Println(v6, v7)
    v6 = 32767
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
