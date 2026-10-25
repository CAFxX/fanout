package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }

func main() {
    v1 := make(map[string]bool)
    var v2 uint8 = 42
    var v3 string = ("foo bar" + ("" + "a"))
    v4 := GenS1{F0: (v2 / (((v2 % ((255) | 1))) | 1))}
    for fi5 := 0; fi5 < 11; fi5++ {
        v6 := []int{0, (fi5 ^ fi5), fi5}
        var v7 int8 = 127
        v7 = v7
        _ = fi5
        _ = v6
        _ = v7
    }
    if true {
        var v8 uint = 0
        v3 = v3
        _ = v8
    } else {
        v3 = v3
        v2 += (v2 << ((42) & 63))
    }
    delete(v1, v3)
    if true {
        var v9 uint32 = 2
        var v10 string = v3
        _ = v10
        _ = v9
    }
    v11 := []int{436, -454, 32768}
    var v12 uint = 115
    v13 := v1[v3]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
}
