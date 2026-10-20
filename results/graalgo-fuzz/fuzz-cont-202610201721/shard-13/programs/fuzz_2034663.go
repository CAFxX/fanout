package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64 }
type GenS2 struct { F0 int8; F1 int8 }

func main() {
    var v1 uint16 = 32767
    v2 := []uint32{2147483648, 2147483648, 1}
    v3 := []string{}
    v4 := []uint32{0}
    fmt.Println(v1)
    fmt.Println(v1)
    var v5 float32 = float32(0.1)
    var v6 uint32 = 414
    if (((v1 | 2) > (^(v1))) || (!(1 >= 42))) {
        if true {
            fmt.Println(v6)
            v6 -= v6
            if true {
                v7 := GenS1{F0: v1, F1: 0}
                fmt.Println(v1, v6, v1)
                v6 -= v6
                _ = v7
            }
        } else {
            var v8 int32 = -1
            v6 -= 770
            _ = v8
        }
        v6 = uint32(v1)
    }
    fmt.Println(v1, v5, v5)
    v6 = v6
    if true {
        v2 = append(v2, (v6 << ((0) & 63)))
        v9 := GenS2{F0: 108, F1: 42}
        fmt.Println(v5)
        _ = v9
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
