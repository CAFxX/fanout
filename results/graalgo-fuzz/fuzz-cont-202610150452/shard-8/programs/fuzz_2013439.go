package main

import (
    "fmt"
)


func main() {
    var v1 int16 = 105
    v2 := []int{32767}
    v3 := []string{"a"}
    v4 := []int{9223372036854775807}
    var v5 int = 32767
    v6 := []uint32{4294967295, 42, 2147483647}
    v5 -= (v5 >> ((75) & 63))
    var v7 uint8 = 192
    v5 = v5
    v7 -= 255
    fmt.Println(v5, v5)
    v3 = append(v3, ("" + ""))
    fmt.Println(v5, v5, v1)
    if (v5 > (v5 * (v5 | (v5 & (v5 | 128))))) {
        v2 = append(v2, (^(v5)))
    } else {
        var v8 uint32 = 515
        _ = v8
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
