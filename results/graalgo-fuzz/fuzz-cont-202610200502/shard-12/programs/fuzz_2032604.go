package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }

func main() {
    var v1 uint = 4294967295
    v2 := make(map[int]bool)
    v3 := []string{("a" + "hello"), "foo bar", ""}
    fmt.Println(v1)
    var v4 float64 = 3.14159
    var v5 uint64 = uint64(v1)
    v4 -= v4
    var v6 int8 = 127
    v7 := GenS1{F0: ("hello" <= "a")}
    for fi8 := 0; fi8 < 19; fi8++ {
        var v9 int = fi8
        _ = fi8
        _ = v9
    }
    v1 -= 1
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
