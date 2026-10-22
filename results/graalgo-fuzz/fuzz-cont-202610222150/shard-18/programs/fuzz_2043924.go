package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64 }
type GenS2 struct { F0 string }

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[string]bool)
    var v3 uint = 4294967295
    var v4 int = -9223372036854775808
    var v5 string = ("a" + ("a" + ("foo bar" + ("a" + "foo bar"))))
    v6 := GenS1{F0: "", F1: 2}
    fmt.Println(v3, v5, v3)
    var v7 uint32 = 0
    v2[v5] = false
    var v8 uint32 = 451
    v8 += 430
    delete(v2, v5)
    v9 := GenS2{F0: v5}
    if true {
        var v10 uint16 = 255
        v4 -= v4
        switch v5 {
        case "hello":
            v4 = v4
        }
        _ = v10
    }
    var v11 int64 = -9223372036854775808
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
