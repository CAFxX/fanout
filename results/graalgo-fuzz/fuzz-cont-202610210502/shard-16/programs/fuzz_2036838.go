package main

import (
    "fmt"
)


func main() {
    v1 := []uint32{65535, 858, 4294967295, 220}
    var v2 int = 32768
    var v3 uint16 = 497
    switch "hello" {
    case "foo bar":
        fmt.Println(v3, v3)
    }
    _ = len(v1)
    if true {
        var v4 uint8 = 113
        v1 = append(v1, 685)
        v2 += (^(v2))
        _ = v4
    }
    fmt.Println(v2, v2)
    fmt.Println(v3)
    v1 = append(v1, 992)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v5 uint = 453
    if len(v1) > 0 {
        v1[0] = 265
    }
    v5 -= 486
    fmt.Println(v3, v2, v3)
    var v6 uint64 = 2
    var v7 float32 = ((1e10 - 0.0) * ((-2.25 / 1.0) - 1e10))
    fmt.Println(v3, v7, v2)
    fmt.Println(v7)
    if false {
        var v8 float64 = (1.5 * 1e10)
        var v9 uint = 2
        var v10 uint16 = v3
        _ = v10
        _ = v8
        _ = v9
    } else {
        v7 = v7
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v5
    _ = v6
    _ = v7
}
