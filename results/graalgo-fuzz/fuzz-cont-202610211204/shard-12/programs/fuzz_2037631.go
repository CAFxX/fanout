package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64; F2 bool }
type GenS2 struct { F0 uint16 }

func main() {
    var v1 int8 = -128
    var v2 uint8 = 110
    var v3 string = "a"
    v4 := []uint32{42, 685}
    if false {
        for fi5 := 0; fi5 < 2; fi5++ {
            var v6 int16 = 1
            v6 = -1
            _ = fi5
            _ = v6
        }
        var v7 int = 4294967295
        _ = v7
    }
    v2 = (^(v2))
    var v8 uint64 = 0
    v9 := GenS1{F0: -32768, F1: 42, F2: false}
    var v10 string = (v3 + (v3 + "foo bar"))
    v8 = (v8 | (v8 << ((65535) & 63)))
    v11 := GenS1{F0: -32768, F1: -896, F2: (0 > -1)}
    var v12 uint32 = 2147483647
    if false {
        for fi13 := 0; fi13 < 12; fi13++ {
            fmt.Println(v1, v2, v10)
            v1 = -128
            _ = fi13
        }
        for fi14 := 0; fi14 < 12; fi14++ {
            v4 = append(v4, 955)
            _ = fi14
        }
    }
    if ((1000 < 42) || false) {
        var v15 float32 = 3.14159
        _ = v15
    }
    v16 := GenS2{F0: 698}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
