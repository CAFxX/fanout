package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16; F2 uint8 }

func main() {
    var v1 int64 = -9223372036854775808
    v2 := make(map[string]string)
    var v3 int = 128
    var v4 int16 = -891
    v5 := []uint32{794}
    var v6 int32 = -2147483648
    v7 := GenS1{F0: 25, F1: 2, F2: 255}
    for fi8 := 0; fi8 < 7; fi8++ {
        var v9 int = 32767
        v9 = (fi8 - (v9 >> ((0) & 63)))
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        _ = fi8
        _ = v9
    }
    v10 := GenS1{F0: 233, F1: 42, F2: 127}
    _ = len(v5)
    v11 := GenS1{F0: 42, F1: 264, F2: 127}
    v6 = (v6 ^ -989)
    var v12 uint64 = 243
    v4 -= v4
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
