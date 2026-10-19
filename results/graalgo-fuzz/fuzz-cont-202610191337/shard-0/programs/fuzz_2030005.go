package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16; F2 int64 }
type GenS2 struct { F0 uint32; F1 int16; F2 uint32 }

func main() {
    var v1 uint = 42
    var v2 int64 = 9223372036854775807
    var v3 int32 = -1
    var v4 uint64 = 670
    var v5 int64 = v2
    for fi6 := 0; fi6 < 2; fi6++ {
        v7 := GenS1{F0: 1, F1: 0, F2: v5}
        fmt.Println(v3, v4)
        v8 := []int{}
        _ = fi6
        _ = v7
        _ = v8
    }
    for fi9 := 0; fi9 < 17; fi9++ {
        if ((^(v1)) == v1) {
            fmt.Println(v1, v3, v3)
            v10 := []int{}
            _ = v10
        }
        fmt.Println(v3, fi9)
        _ = fi9
    }
    fmt.Println(v2, v3, v4)
    var v11 uint64 = v4
    var v12 uint64 = (v4 & (v11 % ((468) | 1)))
    var v13 uint8 = 127
    var v14 int64 = -1
    var v15 int = 174
    var v16 int = (v15 & v15)
    v17 := GenS1{F0: 42, F1: -32768, F2: (^(v14))}
    v12 = (v12 + v4)
    v18 := GenS2{F0: 2147483647, F1: -157, F2: 65535}
    v19 := GenS2{F0: 4294967295, F1: -32768, F2: 4294967295}
    fmt.Println(v1, v16, v16)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
