package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }

func main() {
    var v1 bool = false
    var v2 int8 = 38
    var v3 int8 = -128
    var v4 uint32 = 0
    var v5 string = "hello"
    for fi6 := 0; fi6 < 4; fi6++ {
        fmt.Println(v1, v3)
        v3 = v2
        _ = fi6
    }
    var v7 uint = 2
    v8 := []int{1, 636}
    var v9 int = 32768
    v5 = (v5 + (v5 + ("a" + "hello")))
    _ = len(v8)
    fmt.Println(v1, v1)
    for fi10 := 0; fi10 < 18; fi10++ {
        if (0 != 0) {
            var v11 int64 = 1
            _ = v11
        }
        var v12 uint = 76
        fmt.Println(v3, v7)
        _ = fi10
        _ = v12
    }
    if v1 {
        v3 = (^(v2))
    }
    if true {
        var v13 int64 = -148
        var v14 int64 = 42
        switch v9 {
        case 32767:
            v15 := GenS1{F0: (v14 | v13)}
            _ = v15
        }
        _ = v13
        _ = v14
    }
    var v16 int8 = v2
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
