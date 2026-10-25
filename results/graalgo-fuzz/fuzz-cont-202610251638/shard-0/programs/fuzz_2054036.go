package main

import (
    "fmt"
)


func main() {
    var v1 uint16 = 255
    var v2 uint = 42
    v3 := make(map[uint32]uint32)
    var v4 int64 = 42
    var v5 int8 = 100
    var v6 bool = true
    for fi7 := 0; fi7 < 16; fi7++ {
        v8 := []int{-588, (fi7 + fi7)}
        if len(v8) >= 2 {
            v8 = v8[1:len(v8)]
        }
        _ = fi7
        _ = v8
    }
    v1 += (v1 & 656)
    var v9 int16 = 1000
    v10 := []int{128, -152, -212}
    if true {
        var v11 uint = v2
        var v12 uint8 = 1
        _ = v11
        _ = v12
    } else {
        v6 = true
    }
    fmt.Println(v9, v4, v1)
    _ = len(v3)
    fmt.Println(v9)
    v4 += 42
    v6 = v6
    fmt.Println(v1, v4)
    fmt.Println(v5, v5, v6)
    v4 += v4
    if len(v10) >= 2 {
        v10 = v10[1:len(v10)]
    }
    if v6 {
        v4 -= 690
    }
    var v13 int32 = 2147483647
    v4 -= (v4 | (v4 << ((v2) & 63)))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
