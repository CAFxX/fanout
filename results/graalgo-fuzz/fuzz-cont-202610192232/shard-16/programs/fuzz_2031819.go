package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64; F2 uint64 }

func main() {
    var v1 int32 = 864
    var v2 float64 = float64(1e10)
    var v3 uint16 = 576
    var v4 string = "hello"
    var v5 uint32 = 0
    fmt.Println(v4)
    v6 := GenS1{F0: 4294967295, F1: 2, F2: 1}
    v7 := []int{128}
    var v8 uint8 = 127
    for fi9 := 0; fi9 < 1; fi9++ {
        fmt.Println(v8)
        var v10 int64 = 1
        v8 += v8
        _ = fi9
        _ = v10
    }
    for fi11 := 0; fi11 < 6; fi11++ {
        switch "" {
        case "hello":
            if len(v7) > 0 {
                v7[0] = (fi11 % (((fi11 & fi11)) | 1))
            }
        }
        v1 = (v1 * 867)
        _ = fi11
    }
    if true {
        v3 += v3
        fmt.Println(v2)
        fmt.Println(v1, v5)
    }
    if len(v7) > 0 {
        v7[0] = -145
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
