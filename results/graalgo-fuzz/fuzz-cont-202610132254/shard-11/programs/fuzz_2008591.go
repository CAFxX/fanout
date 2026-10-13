package main

import (
    "fmt"
)


func main() {
    v1 := make(map[uint32]bool)
    var v2 uint16 = 32767
    var v3 uint64 = 278
    var v4 int = 9223372036854775807
    v5 := make(map[string]string)
    var v6 bool = (-58 <= 162)
    var v7 float32 = -2.25
    for fi8 := 0; fi8 < 20; fi8++ {
        delete(v5, "foo bar")
        _ = fi8
    }
    v6 = (1 != 2)
    var v9 float64 = 1e10
    fmt.Println(v6, v7, v9)
    var v10 int64 = -1
    if (v7 != 0.0) {
        if (((!(!v6)) || (-1 < 0)) && false) {
            var v11 uint64 = (v3 % ((v3) | 1))
            fmt.Println(v10, v7)
            _ = v11
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
