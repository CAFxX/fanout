package main

import (
    "fmt"
)


func main() {
    v1 := []int{9223372036854775807, 32767, 42, 0}
    v2 := []uint32{42, 0, 510, 782}
    var v3 int8 = -128
    var v4 int32 = 2147483647
    for fi5 := 0; fi5 < 20; fi5++ {
        fmt.Println(fi5, fi5)
        _ = fi5
    }
    fmt.Println(v4)
    fmt.Println(v4)
    _ = len(v1)
    fmt.Println(v3)
    var v6 int = 42
    fmt.Println(v4, v3, v6)
    v1 = append(v1, 1)
    for fi7 := 0; fi7 < 15; fi7++ {
        var v8 uint8 = 42
        _ = fi7
        _ = v8
    }
    switch ("hello" + "") {
    case "hello":
        fmt.Println(v6, v6, v3)
    default:
        fmt.Println(v6)
    }
    v6 += v6
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v6
}
