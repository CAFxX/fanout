package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }

func main() {
    v1 := []int{-9223372036854775808, 0}
    v2 := make(map[int]bool)
    var v3 int = 9223372036854775807
    v4 := []string{("\n\t" + "a"), "", "hello", "hello"}
    var v5 uint = 18446744073709551615
    v6 := []string{}
    var v7 uint16 = 2
    v3 -= (v3 | (v3 / ((491) | 1)))
    v8 := GenS1{F0: 230}
    v3 -= (v3 & 127)
    switch "\n\t" {
    case "foo bar":
        var v9 int16 = 42
        _ = v9
    }
    v7 = (v7 - 5)
    v5 = 255
    v5 += (v5 >> (((v5 | v5)) & 63))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
