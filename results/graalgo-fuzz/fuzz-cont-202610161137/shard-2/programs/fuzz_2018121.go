package main

import (
    "fmt"
)


func main() {
    v1 := make(map[uint32]string)
    var v2 uint8 = 1
    var v3 int64 = 1
    v2 = v2
    var v4 float32 = 1.5
    var v5 float32 = 0.1
    for fi6 := 0; fi6 < 3; fi6++ {
        v1[65535] = ""
        _ = fi6
    }
    var v7 uint64 = 42
    var v8 uint = 255
    v8 -= (v8 >> (((^(v8))) & 63))
    v5 -= (v4 / v5)
    fmt.Println(v2)
    var v9 float32 = v4
    var v10 uint16 = 903
    v7 += v7
    var v11 int = 937
    v9 = v5
    v12 := []int{-902, v11}
    for fi13 := 0; fi13 < 9; fi13++ {
        v1[4294967295] = ("a" + "foo bar")
        delete(v1, 882)
        fmt.Println(v7)
        _ = fi13
    }
    var v14 int = v11
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
