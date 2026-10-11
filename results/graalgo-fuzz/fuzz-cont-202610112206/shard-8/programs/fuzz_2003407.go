package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32; F2 int8 }

func main() {
    var v1 bool = (1000 >= -150)
    var v2 int = -194
    v3 := make(map[string]uint32)
    v4 := []uint32{}
    v5 := []uint32{42, 655}
    var v6 uint8 = 255
    fmt.Println(v2, v6)
    var v7 string = "foo bar"
    var v8 int64 = 9223372036854775807
    var v9 bool = (!v1)
    v10 := GenS1{F0: 1, F1: 0, F2: -128}
    if len(v4) > 0 {
        v4[0] = 425
    }
    if v1 {
        fmt.Println(v6, v7)
    }
    var v11 int8 = 100
    var v12 bool = (!false)
    for fi13 := 0; fi13 < 4; fi13++ {
        if v9 {
            var v14 int = fi13
            fmt.Println(v1, v1, v2)
            v14 = 0
            _ = v14
        } else {
            v15 := v3["a"]
            _ = len(v5)
            var v16 uint32 = 4294967295
            _ = v15
            _ = v16
        }
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        _ = fi13
    }
    var v17 uint = 18446744073709551615
    if true {
        if v12 {
            var v18 bool = (885 >= 296)
            v2 = 32767
            var v19 float64 = 3.14159
            _ = v18
            _ = v19
        }
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    }
    fmt.Println(v2)
    v12 = v12
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
