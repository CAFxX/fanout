package main

import (
    "fmt"
)


func main() {
    v1 := []uint32{2147483647, 2, 486, 4294967295}
    var v2 uint32 = 4294967295
    var v3 int32 = -2147483648
    if len(v1) > 0 {
        v1[0] = (v2 % ((v2) | 1))
    }
    for fi4 := 0; fi4 < 19; fi4++ {
        var v5 int32 = (v3 >> ((0) & 63))
        _ = fi4
        _ = v5
    }
    var v6 uint32 = 2147483648
    var v7 bool = false
    v6 = (^(v6))
    var v8 int16 = -923
    v6 = 65535
    var v9 bool = ((v2 >> ((1) & 63)) > (v2 >> ((1) & 63)))
    fmt.Println(v3, v6)
    if len(v1) > 0 {
        v1[0] = (^(v2))
    }
    var v10 int32 = int32(v2)
    _ = len(v1)
    v7 = v9
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
