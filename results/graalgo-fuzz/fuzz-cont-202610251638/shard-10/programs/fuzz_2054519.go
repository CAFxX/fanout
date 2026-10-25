package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32 }
type GenS2 struct { F0 string; F1 uint64 }
func genFn0() uint32 {
    var v5 uint16 = 42
    switch 32767 {
    case 213:
        var v6 uint64 = 0
        _ = v6
    }
    _ = v5
    return 502
}

func main() {
    var v1 int64 = -1
    v2 := make(map[uint32]uint32)
    var v3 int64 = v1
    v4 := make(map[int]uint32)
    delete(v2, 4294967295)
    fmt.Println(v3, v3)
    v3 += (v1 - v1)
    var v7 int64 = v3
    fmt.Println(v7, v3)
    for fi8 := 0; fi8 < 9; fi8++ {
        v7 = (v3 % ((9223372036854775807) | 1))
        for fi9 := 0; fi9 < 16; fi9++ {
            var v10 int32 = -124
            v11 := GenS1{F0: 505, F1: 65535}
            _ = fi9
            _ = v10
            _ = v11
        }
        var v12 string = ("hello" + "foo bar")
        _ = fi8
        _ = v12
    }
    fmt.Println(v1)
    var v13 int32 = 2147483647
    var v14 float32 = 0.1
    v3 = v1
    v14 = 1.5
    var v15 int = -775
    v16 := GenS2{F0: ("" + ("hello" + ("a" + "hello"))), F1: 901}
    fmt.Println(v14, v1, v13)
    var v17 bool = false
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v7
}
