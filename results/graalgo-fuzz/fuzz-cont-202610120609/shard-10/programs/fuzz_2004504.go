package main

import (
    "fmt"
)


func main() {
    v1 := make(map[int]bool)
    v2 := []uint32{}
    v3 := make(map[uint32]int)
    v4 := []int{9223372036854775807, -9223372036854775808, 529}
    v5 := make(map[int]int)
    v6 := []int{0, 2147483647, -1, 32768}
    if (false && true) {
        var v7 int16 = 0
        var v8 int8 = 0
        _ = v7
        _ = v8
    } else {
        v9 := v1[18]
        _ = v9
    }
    switch 2147483647 {
    case 32767:
        var v10 uint8 = 128
        _ = v10
    }
    var v11 float32 = (-2.25 - -2.25)
    var v12 string = ""
    delete(v5, 473)
    _ = len(v5)
    if true {
        _ = len(v3)
    }
    v12 = "foo bar"
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
