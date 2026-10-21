package main

import (
    "fmt"
)


func main() {
    var v1 int64 = 1
    v2 := make(map[string]uint32)
    v3 := make(map[int]string)
    var v4 uint32 = 0
    v5 := make(map[int]int)
    v1 += v1
    v6 := []int{-947, 42}
    var v7 string = "hello"
    var v8 uint = 42
    v7 = v7
    for fi9 := 0; fi9 < 13; fi9++ {
        fmt.Println(v1, fi9, v7)
        var v10 uint64 = 0
        _ = fi9
        _ = v10
    }
    fmt.Println(v4)
    if (true || (!true)) {
        switch "" {
        case "":
            var v11 int = 42
            _ = v11
        case "hello":
            var v12 float32 = 3.14159
            _ = v12
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
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
