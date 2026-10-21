package main

import (
    "fmt"
)


func main() {
    v1 := make(map[int]bool)
    var v2 int64 = -208
    var v3 uint16 = 2
    for fi4 := 0; fi4 < 3; fi4++ {
        v3 = (v3 >> ((42) & 63))
        if (int64(fi4) != 1) {
            v5 := []int{-10}
            fmt.Println(v2, fi4)
            v3 += v3
            _ = v5
        } else {
            var v6 bool = (127 < 127)
            _ = v6
        }
        _ = fi4
    }
    v2 = (v2 / ((v2) | 1))
    var v7 int = -837
    var v8 int = 367
    var v9 uint32 = 2147483647
    fmt.Println(v2)
    v3 = (v3 >> ((484) & 63))
    v10 := []int{(v8 ^ -2147483648)}
    var v11 uint8 = 1
    if len(v10) > 0 {
        v10[0] = v8
    }
    var v12 int32 = 1
    if false {
        for fi13 := 0; fi13 < 19; fi13++ {
            var v14 int64 = v2
            _ = fi13
            _ = v14
        }
    } else {
        fmt.Println(v9, v2)
        v15 := v1[-2147483648]
        v2 += (v2 ^ 0)
        _ = v15
    }
    var v16 int32 = (v12 + 1)
    v16 -= v16
    switch 4294967295 {
    case 2147483647:
        _ = len(v10)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v16
    _ = v2
    _ = v3
    _ = v7
    _ = v8
    _ = v9
}
