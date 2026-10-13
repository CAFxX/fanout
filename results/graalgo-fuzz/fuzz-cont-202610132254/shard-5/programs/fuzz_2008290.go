package main

import (
    "fmt"
)


func main() {
    var v1 uint32 = 4294967295
    v2 := make(map[string]bool)
    var v3 bool = (!false)
    var v4 string = "\n\t"
    v5 := make(map[string]uint32)
    var v6 uint16 = 2
    var v7 int64 = 0
    var v8 uint = 2
    v9 := []int{-9223372036854775808, -1}
    for fi10 := 0; fi10 < 18; fi10++ {
        v4 = v4
        _ = fi10
    }
    if v3 {
        v9 = append(v9, 42)
        var v11 int = 42
        _ = v11
    }
    v4 = v4
    if v3 {
        var v12 string = v4
        _ = v12
    } else {
        var v13 string = ""
        fmt.Println(v3, v1)
        _ = v13
    }
    v14 := v5[v4]
    switch -694 {
    case 32768:
        fmt.Println(v6)
    default:
        fmt.Println(v7)
    }
    var v15 uint64 = uint64(v8)
    _ = len(v9)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
