package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }

func main() {
    v1 := make(map[string]int)
    v2 := make(map[int]uint32)
    v3 := make(map[string]uint32)
    v4 := make(map[uint32]string)
    v5 := []uint32{65535, 42, 2147483647, 4294967295}
    v6 := make(map[string]string)
    for fi7 := 0; fi7 < 4; fi7++ {
        var v8 int64 = 1
        var v9 int32 = 477
        fmt.Println(v9)
        _ = fi7
        _ = v8
        _ = v9
    }
    _ = len(v5)
    var v10 uint16 = 107
    var v11 float64 = -2.25
    v11 -= v11
    var v12 int64 = 9223372036854775807
    v12 -= v12
    var v13 int = 127
    v13 += v13
    var v14 int32 = 2147483647
    var v15 float64 = 3.14159
    for fi16 := 0; fi16 < 1; fi16++ {
        v17 := GenS1{F0: 0}
        _ = fi16
        _ = v17
    }
    var v18 int8 = 100
    var v19 uint16 = v10
    var v20 uint = 2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
