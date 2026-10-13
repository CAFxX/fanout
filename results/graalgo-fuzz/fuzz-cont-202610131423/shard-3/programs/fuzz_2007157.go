package main

import (
    "fmt"
)


func main() {
    var v1 float64 = ((0.0 - 3.14159) - 1e10)
    v2 := []uint32{2147483647, 653}
    v3 := make(map[uint32]bool)
    fmt.Println(v1)
    _ = len(v2)
    fmt.Println(v1)
    var v4 uint32 = 65535
    fmt.Println(v4, v1)
    if (!false) {
        v2 = append(v2, v4)
        fmt.Println(v4, v4)
    } else {
        fmt.Println(v1, v1)
        delete(v3, (v4 << ((42) & 63)))
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    var v5 uint16 = 2
    var v6 string = ("" + "a")
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
