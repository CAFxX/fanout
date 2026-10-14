package main

import (
    "fmt"
)


func main() {
    var v1 uint32 = 2147483648
    var v2 float32 = (float32(1e10) + 1e10)
    var v3 int = 617
    for fi4 := 0; fi4 < 7; fi4++ {
        v3 -= v3
        v2 = v2
        _ = fi4
    }
    if (v2 <= (v2 / (v2 + v2))) {
        v3 = (v3 << ((4294967295) & 63))
        for fi5 := 0; fi5 < 4; fi5++ {
            fmt.Println(v1, v3)
            _ = fi5
        }
    }
    var v6 uint64 = 401
    v7 := []int{v3, v3}
    var v8 int32 = -1
    _ = len(v7)
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    v7 = append(v7, 0)
    var v9 int16 = 1
    v8 = v8
    v8 = v8
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
