package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8; F2 uint8 }

func main() {
    v1 := []uint32{}
    v2 := make(map[uint32]bool)
    var v3 bool = true
    var v4 int8 = 100
    var v5 int32 = 948
    var v6 bool = (((v4 % (((v4 | 0)) | 1)) < v4) || false)
    var v7 uint16 = 32767
    v3 = false
    var v8 bool = v3
    switch "hello" {
    case "hello":
        v1 = append(v1, 65535)
    case "foo bar":
        v9 := GenS1{F0: 2147483647, F1: 127, F2: 42}
        _ = v9
    case "\n\t":
        if v6 {
            fmt.Println(v5, v5)
        }
    default:
        var v10 uint8 = 128
        _ = v10
    }
    v11 := GenS1{F0: uint32(v7), F1: 56, F2: 81}
    v12 := GenS1{F0: 4294967295, F1: 128, F2: 1}
    if v3 {
        var v13 string = ""
        for fi14 := 0; fi14 < 15; fi14++ {
            var v15 int64 = 9223372036854775807
            v1 = append(v1, 65535)
            _ = fi14
            _ = v15
        }
        v16 := GenS1{F0: 392, F1: 128, F2: 2}
        _ = v13
        _ = v16
    }
    switch 128 {
    case -2147483648:
        var v17 int16 = -1
        _ = v17
    }
    v5 -= int32(v4)
    v18 := GenS1{F0: 65535, F1: 128, F2: 80}
    v19 := GenS1{F0: 0, F1: 255, F2: 2}
    for fi20 := 0; fi20 < 14; fi20++ {
        v3 = (!v8)
        _ = fi20
    }
    var v21 int32 = (v5 | v5)
    delete(v2, 2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v18
    _ = v19
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
