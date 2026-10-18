package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32 }
type GenS2 struct { F0 uint32; F1 uint8; F2 uint32 }
type GenS3 struct { F0 uint; F1 int8 }

func main() {
    var v1 uint16 = 255
    var v2 float64 = 0.1
    var v3 int = 9223372036854775807
    var v4 int = (^(v3))
    v3 += (v3 - v4)
    v5 := GenS1{F0: 65535, F1: 944}
    v6 := GenS2{F0: 845, F1: 20, F2: 4294967295}
    var v7 bool = ((273 > 2147483647) && false)
    v8 := []int{4294967295, (v4 | (v3 & (v3 << ((129) & 63)))), v4}
    for fi9 := 0; fi9 < 6; fi9++ {
        fmt.Println(v7, v7, fi9)
        _ = fi9
    }
    v10 := GenS3{F0: 4294967295, F1: -44}
    v11 := GenS2{F0: 65535, F1: 63, F2: 398}
    for fi12 := 0; fi12 < 4; fi12++ {
        var v13 uint = 2
        _ = fi12
        _ = v13
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
