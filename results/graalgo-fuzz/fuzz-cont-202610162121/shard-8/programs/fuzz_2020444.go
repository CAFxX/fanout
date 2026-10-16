package main

import (
    "fmt"
)


func main() {
    var v1 uint16 = 1
    v2 := []string{""}
    v3 := make(map[uint32]int)
    v4 := []string{"\n\t"}
    var v5 uint = 480
    v5 -= v5
    var v6 string = "foo bar"
    fmt.Println(v6)
    var v7 int = 92
    var v8 uint16 = (v1 + 763)
    fmt.Println(v5)
    for fi9 := 0; fi9 < 5; fi9++ {
        var v10 bool = (-32768 >= 0)
        _ = fi9
        _ = v10
    }
    var v11 int = 0
    var v12 uint = v5
    for fi13 := 0; fi13 < 13; fi13++ {
        v7 = 0
        _ = len(v3)
        _ = fi13
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
