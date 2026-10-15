package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
func genFn0(p0 uint, p1 uint32, p2 int32) uint32 {
    p2 += (p2 % ((p2) | 1))
    p1 += 42
    p0 = p0
    _ = p0
    _ = p1
    _ = p2
    return 594
}

func main() {
    v1 := []string{"", "foo bar", ""}
    v2 := make(map[uint32]int)
    var v3 int16 = 0
    v4 := make(map[int]bool)
    v5 := make(map[uint32]bool)
    var v6 int = 1
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v7 string = (("hello" + ("a" + ("\n\t" + "a"))) + "a")
    var v8 uint8 = 255
    v7 = ((v7 + "foo bar") + v7)
    for fi9 := 0; fi9 < 13; fi9++ {
        v10 := GenS1{F0: 1}
        _ = fi9
        _ = v10
    }
    var v11 float32 = -2.25
    var v12 uint64 = 61
    var v13 int8 = int8(v3)
    for fi14 := 0; fi14 < 20; fi14++ {
        v7 = v7
        _ = fi14
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
