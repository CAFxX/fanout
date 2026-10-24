package main

import (
    "fmt"
)

type GenS1 struct { F0 int }

func main() {
    v1 := []uint32{}
    var v2 uint = 42
    v3 := make(map[string]string)
    v4 := []string{"hello"}
    v5 := []int{4294967295}
    var v6 int8 = 127
    v7 := GenS1{F0: -1}
    for fi8 := 0; fi8 < 5; fi8++ {
        v9 := GenS1{F0: fi8}
        _ = fi8
        _ = v9
    }
    v2 += v2
    if len(v5) > 0 {
        v5[0] = 42
    }
    fmt.Println(v2, v2)
    v10 := GenS1{F0: 784}
    if true {
        var v11 uint16 = 32767
        var v12 int64 = 1
        var v13 int16 = 42
        _ = v11
        _ = v12
        _ = v13
    }
    v14 := GenS1{F0: 399}
    fmt.Println(v2, v6)
    if true {
        var v15 int64 = -769
        v15 = 9223372036854775807
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = v15
    }
    switch 4294967295 {
    case 128:
        v16 := v3["hello"]
        _ = v16
    case 2147483647:
        v2 = 42
    default:
        var v17 uint32 = 1
        _ = v17
    }
    var v18 int = -1
    var v19 int32 = -2147483648
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v14
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
