package main

import (
    "fmt"
)


func main() {
    v1 := []int{529, -628}
    v2 := []int{-1, -9223372036854775808}
    v3 := []string{"", ("hello" + ("\n\t" + "\n\t")), "hello"}
    var v4 int8 = -53
    v5 := make(map[uint32]bool)
    var v6 uint32 = 4294967295
    for fi7 := 0; fi7 < 12; fi7++ {
        fmt.Println(v4, fi7)
        _ = fi7
    }
    v4 += -42
    fmt.Println(v4)
    var v8 int8 = v4
    var v9 uint32 = (^(v6))
    var v10 int = 0
    fmt.Println(v8)
    var v11 bool = true
    var v12 int = int(v4)
    for fi13 := 0; fi13 < 13; fi13++ {
        fmt.Println(v4, v4, v12)
        var v14 int32 = 42
        _ = fi13
        _ = v14
    }
    var v15 uint16 = 436
    if (252 != 1) {
        var v16 int32 = 1
        _ = v16
    }
    v4 = v8
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
