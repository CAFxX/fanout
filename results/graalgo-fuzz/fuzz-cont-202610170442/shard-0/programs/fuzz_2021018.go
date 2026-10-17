package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int32; F1 int64 }

func main() {
    var v1 uint = 912
    v2 := []int{963, 1, 32768}
    v3 := make(map[int]bool)
    v4 := make(map[uint32]bool)
    fmt.Println(v1)
    var v5 uint64 = 4294967295
    var v6 int8 = int8(v1)
    v7 := GenS1{F0: 656}
    if (1 != 1000) {
        v8 := v4[792]
        _ = v8
    }
    v9 := v3[32767]
    v6 = v6
    v2 = append(v2, 994)
    for fi10 := 0; fi10 < 11; fi10++ {
        var v11 uint = v1
        _ = fi10
        _ = v11
    }
    _ = len(v2)
    fmt.Println(v9, v6)
    var v12 int32 = 42
    v5 -= 4294967295
    if v9 {
        var v13 uint32 = 2147483647
        var v14 int16 = 104
        v15 := GenS2{F0: v12, F1: 7}
        _ = v13
        _ = v14
        _ = v15
    } else {
        var v16 uint64 = (v5 / (((v5 - 2)) | 1))
        _ = v16
    }
    var v17 string = "a"
    delete(v4, 285)
    for fi18 := 0; fi18 < 20; fi18++ {
        var v19 uint = v1
        v12 = int32(v19)
        var v20 int = fi18
        _ = fi18
        _ = v19
        _ = v20
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
