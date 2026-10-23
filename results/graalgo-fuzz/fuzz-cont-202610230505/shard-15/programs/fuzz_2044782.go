package main

import (
    "fmt"
)


func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[string]string)
    var v3 uint64 = 4294967295
    var v4 float32 = 0.1
    v5 := []uint32{42}
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v4 += 1.5
    for fi6 := 0; fi6 < 10; fi6++ {
        if true {
            var v7 float32 = 3.14159
            _ = v7
        }
        var v8 int64 = 42
        var v9 uint64 = 2
        _ = fi6
        _ = v8
        _ = v9
    }
    var v10 int16 = 1
    v10 += v10
    var v11 uint64 = v3
    v4 = (3.14159 - v4)
    var v12 uint8 = 2
    if true {
        var v13 string = "hello"
        switch 32767 {
        case -1:
            v4 -= 1e10
        default:
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        }
        _ = v13
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
