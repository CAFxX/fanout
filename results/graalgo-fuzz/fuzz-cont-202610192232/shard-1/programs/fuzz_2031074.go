package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }

func main() {
    var v1 int16 = 32767
    var v2 int8 = 83
    var v3 float64 = ((3.14159 + (1e10 + (0.1 - 0.1))) * 1.5)
    var v4 int32 = 2147483647
    var v5 uint32 = 65535
    if true {
        v6 := []int{}
        _ = v6
    }
    v4 += v4
    v5 -= v5
    v7 := []int{127, 9223372036854775807}
    _ = len(v7)
    var v8 int32 = -164
    v2 = -1
    var v9 float64 = -2.25
    for fi10 := 0; fi10 < 9; fi10++ {
        v11 := GenS1{F0: 0}
        _ = fi10
        _ = v11
    }
    v1 += (v1 & v1)
    switch -768 {
    case 127:
        var v12 int = 269
        _ = v12
    case 165:
        v13 := GenS1{F0: uint64(v5)}
        _ = v13
    default:
        v7 = append(v7, 32767)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
