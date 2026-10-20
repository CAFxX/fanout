package main

import (
    "fmt"
)

type GenS1 struct { F0 int }

func main() {
    var v1 uint32 = 0
    v2 := make(map[uint32]bool)
    v3 := make(map[string]int)
    var v4 uint32 = 171
    var v5 int = -9223372036854775808
    var v6 string = ("a" + "")
    switch v5 {
    case 9223372036854775807:
        fmt.Println(v5)
    }
    var v7 uint64 = 1
    switch "\n\t" {
    case "a":
        var v8 int8 = -37
        _ = v8
    case "hello":
        if true {
            var v9 int8 = 0
            _ = v9
        } else {
            var v10 int64 = 1
            _ = v10
        }
    case "":
        var v11 int8 = 100
        _ = v11
    }
    v12 := GenS1{F0: (v5 ^ (v5 | 9223372036854775807))}
    switch -891 {
    case 4294967295:
        fmt.Println(v7)
    }
    fmt.Println(v4, v7)
    var v13 int8 = 27
    fmt.Println(v1, v4)
    var v14 bool = false
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
