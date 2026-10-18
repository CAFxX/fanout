package main

import (
    "fmt"
)


func main() {
    var v1 float32 = 0.0
    var v2 int64 = -1
    var v3 float32 = 0.0
    v4 := []string{("a" + "\n\t"), ("foo bar" + ""), "\n\t", "\n\t"}
    var v5 uint32 = 4294967295
    fmt.Println(v2, v1)
    fmt.Println(v1, v2)
    var v6 string = "a"
    var v7 int32 = 42
    fmt.Println(v6, v7, v7)
    for fi8 := 0; fi8 < 6; fi8++ {
        fmt.Println(v7)
        v2 += v2
        _ = fi8
    }
    var v9 uint = 803
    fmt.Println(v3, v6)
    var v10 uint = (v9 ^ 18446744073709551615)
    if true {
        _ = len(v4)
        var v11 uint8 = 127
        _ = v11
    }
    fmt.Println(v2, v10)
    var v12 int32 = v7
    v7 = (^(v7))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
