package main

import (
    "fmt"
)


func main() {
    var v1 uint16 = 697
    v2 := make(map[uint32]string)
    var v3 string = "hello"
    v3 = v3
    var v4 float32 = 0.0
    v4 -= v4
    fmt.Println(v4, v1)
    if (v3 < v3) {
        if (-1 == 0) {
            var v5 uint = 1
            _ = v5
        }
        v4 -= float32(0.1)
    } else {
        var v6 int16 = 990
        _ = v6
    }
    var v7 int16 = -32768
    v7 = (v7 * v7)
    for fi8 := 0; fi8 < 20; fi8++ {
        var v9 string = "hello"
        var v10 uint64 = 682
        var v11 bool = false
        _ = fi8
        _ = v10
        _ = v11
        _ = v9
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v7
}
