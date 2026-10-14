package main

import (
    "fmt"
)


func main() {
    var v1 int8 = 41
    var v2 uint8 = 0
    v3 := []uint32{2}
    var v4 string = ""
    var v5 uint64 = 1
    for fi6 := 0; fi6 < 5; fi6++ {
        v4 = ((v4 + "a") + ((v4 + "hello") + v4))
        var v7 bool = false
        var v8 float32 = -2.25
        _ = fi6
        _ = v7
        _ = v8
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v9 uint = 1
    if false {
        for fi10 := 0; fi10 < 14; fi10++ {
            v3 = append(v3, 2)
            _ = fi10
        }
    }
    _ = len(v3)
    var v11 uint = 4294967295
    fmt.Println(v9, v5)
    var v12 int64 = 52
    v12 -= 42
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
