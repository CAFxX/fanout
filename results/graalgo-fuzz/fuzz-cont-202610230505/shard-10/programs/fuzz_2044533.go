package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }

func main() {
    var v1 uint16 = 65535
    var v2 uint64 = 4294967295
    var v3 int32 = -840
    var v4 float32 = 1e10
    var v5 int8 = -45
    var v6 int32 = -2147483648
    var v7 uint32 = 65535
    fmt.Println(v4)
    var v8 uint32 = uint32(v1)
    for fi9 := 0; fi9 < 3; fi9++ {
        v6 = -2147483648
        v10 := []int{fi9, (fi9 - fi9)}
        fmt.Println(v4)
        _ = fi9
        _ = v10
    }
    v11 := []int{int(v2), 0, -2147483648}
    v4 += (v4 / v4)
    v12 := GenS1{F0: 90}
    var v13 uint32 = 1
    if len(v11) > 0 {
        v11[0] = 912
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
