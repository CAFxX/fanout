package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int8; F1 int16; F2 uint32 }

func main() {
    v1 := make(map[string]int)
    var v2 int8 = 127
    var v3 uint16 = 32767
    v4 := make(map[uint32]int)
    v3 += (^(v3))
    v5 := []int{0}
    var v6 int32 = 42
    if (true || true) {
        var v7 int32 = 1
        var v8 float32 = 1e10
        _ = v7
        _ = v8
    }
    var v9 uint8 = 188
    _ = len(v1)
    if true {
        v3 -= (v3 << ((0) & 63))
        var v10 string = "\n\t"
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        _ = v10
    } else {
        v11 := GenS1{F0: 18446744073709551615}
        var v12 int8 = 0
        _ = v11
        _ = v12
    }
    v13 := GenS2{F0: v2, F1: 918, F2: 2147483647}
    v14 := GenS2{F0: (v2 & v2), F1: -997, F2: 2147483648}
    v15 := GenS1{F0: 42}
    var v16 bool = (4294967295 <= 521)
    v17 := GenS1{F0: 607}
    var v18 int32 = (v6 - v6)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
