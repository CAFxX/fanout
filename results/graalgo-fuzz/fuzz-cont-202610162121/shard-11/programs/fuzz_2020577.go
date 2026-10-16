package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32; F2 int16 }
type GenS2 struct { F0 uint64; F1 string; F2 uint32 }

func main() {
    var v1 int64 = 1
    var v2 int32 = 157
    var v3 uint64 = 2
    v4 := []int{2147483647, 0}
    var v5 int32 = (v2 + (v2 + (^(v2))))
    var v6 int16 = -839
    fmt.Println(v3, v6)
    var v7 uint8 = 1
    v8 := GenS1{F0: 42, F1: 2147483647, F2: (v6 | v6)}
    for fi9 := 0; fi9 < 9; fi9++ {
        v1 -= 42
        var v10 string = ("" + ("" + ""))
        _ = len(v4)
        _ = fi9
        _ = v10
    }
    v11 := GenS2{F0: v3, F1: "", F2: 2147483648}
    v3 -= 18446744073709551615
    var v12 int32 = v5
    if false {
        fmt.Println(v6)
    } else {
        fmt.Println(v2, v3, v2)
        v1 += int64(v12)
        v12 = v5
    }
    if true {
        v4 = append(v4, -627)
        if false {
            var v13 uint8 = 127
            fmt.Println(v12, v13, v12)
            var v14 uint = 709
            _ = v13
            _ = v14
        }
    }
    var v15 int32 = (v12 + (v12 * 2147483647))
    v16 := GenS1{F0: 4294967295, F1: 1, F2: (v6 << ((557) & 63))}
    var v17 int16 = v6
    v6 += v17
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v12
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
