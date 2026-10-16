package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }

func main() {
    var v1 int8 = 42
    var v2 float32 = 1e10
    v3 := make(map[int]int)
    var v4 int32 = 807
    v2 = (v2 - 0.0)
    switch -1 {
    case 1:
        fmt.Println(v2, v1, v4)
    case -1:
        v5 := GenS1{F0: 4294967295}
        _ = v5
    case 42:
        switch "foo bar" {
        case "":
            fmt.Println(v2)
        case "a":
            v6 := GenS1{F0: 987}
            _ = v6
        }
    }
    var v7 uint32 = 882
    if ((1 <= (-(v1))) || (v4 == -809)) {
        v8 := GenS1{F0: 438}
        var v9 uint64 = 557
        fmt.Println(v1, v7)
        _ = v8
        _ = v9
    }
    v7 = (v7 * (v7 * (v7 - 42)))
    fmt.Println(v1, v1, v7)
    var v10 uint32 = 0
    v3[1] = -2147483648
    var v11 uint = 7
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v7
}
