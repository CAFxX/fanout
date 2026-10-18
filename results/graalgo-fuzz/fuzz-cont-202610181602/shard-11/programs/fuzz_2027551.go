package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32; F2 uint32 }
type GenS2 struct { F0 int16; F1 int; F2 int16 }

func main() {
    v1 := []uint32{934, 2, 2147483647, 2}
    var v2 string = "hello"
    var v3 uint16 = 42
    var v4 string = (v2 + v2)
    var v5 string = "foo bar"
    var v6 float32 = 0.1
    fmt.Println(v4, v4, v2)
    v3 -= v3
    v3 = v3
    var v7 uint32 = 0
    var v8 int64 = 9223372036854775807
    if ((79 <= -52) || (v6 > v6)) {
        v9 := GenS1{F0: 1, F1: v7, F2: v7}
        if len(v1) > 0 {
            v1[0] = 42
        }
        _ = v9
    }
    for fi10 := 0; fi10 < 18; fi10++ {
        v11 := GenS2{F0: 42, F1: fi10, F2: 1}
        var v12 float64 = (0.0 / (1.5 - 3.14159))
        _ = fi10
        _ = v11
        _ = v12
    }
    switch 203 {
    case 4294967295:
        for fi13 := 0; fi13 < 15; fi13++ {
            fmt.Println(v5)
            v14 := GenS1{F0: 4294967295, F1: v7, F2: v7}
            _ = fi13
            _ = v14
        }
    }
    v8 += v8
    _ = len(v1)
    for fi15 := 0; fi15 < 16; fi15++ {
        if false {
            var v16 float32 = ((float32(0.1) + v6) - v6)
            if true {
                var v17 uint8 = 255
                _ = v17
            } else {
                var v18 uint8 = 42
                _ = v18
            }
            v2 = v2
            _ = v16
        }
        _ = fi15
    }
    var v19 int32 = 2147483647
    fmt.Println(v8, v8)
    fmt.Println(v19)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
