package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8 }
func genFn0(p0 int, p1 int64, p2 uint32) uint64 {
    var v4 int64 = (p1 ^ p1)
    v5 := []int{p0, 42}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 4294967295
}

func main() {
    v1 := make(map[string]string)
    var v2 int8 = 0
    var v3 uint64 = 787
    v6 := GenS1{F0: 396, F1: v2}
    v7 := []int{}
    v2 += -93
    for fi8 := 0; fi8 < 3; fi8++ {
        var v9 uint8 = 180
        for fi10 := 0; fi10 < 16; fi10++ {
            var v11 int = fi10
            delete(v1, ("hello" + ("a" + ("hello" + "a"))))
            fmt.Println(v11)
            _ = fi10
            _ = v11
        }
        v7 = append(v7, fi8)
        _ = fi8
        _ = v9
    }
    var v12 string = "foo bar"
    fmt.Println(v2)
    fmt.Println(v12, v3)
    var v13 uint64 = v3
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v6
    _ = v7
}
