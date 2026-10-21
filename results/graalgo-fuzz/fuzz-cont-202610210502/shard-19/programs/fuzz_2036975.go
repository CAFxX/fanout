package main

import (
    "fmt"
)


func main() {
    v1 := []int{32768, 42, -2147483648}
    var v2 uint64 = 4294967295
    v3 := make(map[string]uint32)
    var v4 bool = false
    if v4 {
        if len(v1) > 0 {
            v1[0] = -1
        }
        switch -779 {
        case 127:
            v4 = v4
        }
    }
    var v5 int32 = 42
    v2 -= v2
    delete(v3, "foo bar")
    var v6 int16 = int16(v2)
    var v7 uint8 = 255
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v7 = v7
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
