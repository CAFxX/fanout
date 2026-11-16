package main

import (
    "fmt"
)


func main() {
    var v1 int32 = 0
    var v2 string = "hello"
    var v3 uint64 = 4294967295
    v4 := make(map[uint32]bool)
    var v5 uint64 = (v3 >> ((255) & 63))
    switch ((v2 + v2) + (("\n\t" + v2) + "a")) {
    case "":
        v3 = 18446744073709551615
    case "hello":
        _ = len(v4)
    case "a":
        var v6 uint64 = (v3 << ((18446744073709551615) & 63))
        _ = v6
    }
    v7 := []int{127, 0, 960, -9223372036854775808}
    fmt.Println(v2, v3)
    fmt.Println(v1, v1, v1)
    var v8 int64 = int64(v3)
    v1 += (v1 / ((v1) | 1))
    _ = len(v4)
    if true {
        fmt.Println(v8)
        var v9 float64 = 3.14159
        var v10 uint = 4294967295
        _ = v10
        _ = v9
    } else {
        v2 = "foo bar"
        fmt.Println(v2, v2, v8)
        v3 += v3
    }
    var v11 uint64 = v3
    var v12 string = (v2 + v2)
    var v13 int32 = -2147483648
    _ = len(v4)
    v2 = "hello"
    fmt.Println(v2, v1, v1)
    v2 = v12
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
}
