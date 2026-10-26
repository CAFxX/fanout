package main

import (
    "fmt"
)


func main() {
    var v1 uint8 = 128
    var v2 uint16 = 65535
    v3 := make(map[uint32]uint32)
    var v4 uint32 = 2
    var v5 float64 = float64(1.5)
    var v6 int16 = 0
    var v7 string = "a"
    v8 := []int{89}
    if false {
        _ = len(v8)
    }
    if false {
        v7 = v7
    } else {
        var v9 string = "hello"
        v9 = v7
        v10 := v3[2147483647]
        _ = v10
        _ = v9
    }
    switch v7 {
    case "\n\t":
        var v11 float64 = v5
        _ = v11
    case "":
        var v12 int = 42
        _ = v12
    case "a":
        delete(v3, (v4 % (((^(v4))) | 1)))
    }
    var v13 int = -9223372036854775808
    v13 -= v13
    var v14 uint16 = v2
    fmt.Println(v6)
    v13 += v13
    var v15 int8 = 127
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
