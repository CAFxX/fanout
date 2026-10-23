package main

import (
    "fmt"
)


func main() {
    var v1 int16 = -879
    var v2 uint64 = 345
    var v3 string = ""
    var v4 int = int(v2)
    var v5 float32 = 1e10
    v6 := []int{}
    v4 -= v4
    var v7 int64 = -9223372036854775808
    fmt.Println(v5)
    v1 = -23
    var v8 bool = true
    for fi9 := 0; fi9 < 14; fi9++ {
        v6 = append(v6, fi9)
        if false {
            v7 = (v7 & v7)
        }
        fmt.Println(v5)
        _ = fi9
    }
    if len(v6) > 0 {
        v6[0] = (v4 ^ 1)
    }
    for fi10 := 0; fi10 < 2; fi10++ {
        _ = len(v6)
        _ = fi10
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
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
