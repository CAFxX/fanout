package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool; F2 uint }

func main() {
    v1 := make(map[string]uint32)
    v2 := []uint32{0, 437, 2147483648, 2147483648}
    var v3 string = "foo bar"
    var v4 float32 = 1.5
    v5 := []uint32{2147483647, 1}
    var v6 int32 = -1
    v7 := GenS1{F0: 632, F1: (!true), F2: 523}
    var v8 int64 = 9223372036854775807
    _ = len(v5)
    if false {
        v3 = v3
        fmt.Println(v6)
        fmt.Println(v4)
    } else {
        for fi9 := 0; fi9 < 19; fi9++ {
            v10 := GenS1{F0: v8, F1: true, F2: 65535}
            var v11 uint32 = 646
            var v12 int = 127
            _ = fi9
            _ = v10
            _ = v11
            _ = v12
        }
        for fi13 := 0; fi13 < 7; fi13++ {
            v6 -= v6
            v14 := GenS1{F0: v8, F1: true, F2: 18446744073709551615}
            _ = fi13
            _ = v14
        }
    }
    var v15 float32 = 1.5
    fmt.Println(v15, v4)
    v8 = v8
    fmt.Println(v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
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
