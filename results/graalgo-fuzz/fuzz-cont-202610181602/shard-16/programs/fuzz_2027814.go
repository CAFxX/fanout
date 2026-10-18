package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }

func main() {
    var v1 uint8 = 1
    v2 := []int{128}
    v3 := make(map[int]bool)
    var v4 int8 = -128
    var v5 float32 = 3.14159
    v6 := []string{("\n\t" + "hello"), ("foo bar" + "hello"), "foo bar"}
    v4 = (v4 >> ((554) & 63))
    var v7 uint16 = 1
    if (false && ((v7 | (v7 - (v7 % ((v7) | 1)))) < v7)) {
        for fi8 := 0; fi8 < 19; fi8++ {
            v9 := GenS1{F0: 637}
            v4 = 115
            var v10 uint64 = 2
            _ = fi8
            _ = v10
            _ = v9
        }
        var v11 int8 = (v4 >> ((371) & 63))
        _ = v11
    }
    var v12 uint = 2
    v12 = (v12 ^ v12)
    var v13 string = ("foo bar" + "hello")
    fmt.Println(v5, v1)
    v13 = v13
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
