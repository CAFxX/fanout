package main

import (
    "fmt"
)


func main() {
    var v1 uint8 = 128
    v2 := make(map[int]string)
    v3 := []string{"hello", "hello", "a"}
    var v4 uint = 494
    switch -2147483648 {
    case 127:
        delete(v2, 2147483647)
    }
    if true {
        var v5 bool = ((-1 >= -2) && ((false && ((!true) || true)) || (100 == 42)))
        _ = v5
    }
    _ = len(v2)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    _ = len(v3)
    for fi6 := 0; fi6 < 20; fi6++ {
        var v7 int64 = 1
        _ = fi6
        _ = v7
    }
    var v8 uint16 = 205
    v4 = (v4 * v4)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v8
}
