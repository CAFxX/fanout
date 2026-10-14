package main

import (
    "fmt"
)


func main() {
    var v1 string = "hello"
    v2 := make(map[string]uint32)
    var v3 uint8 = 0
    v4 := []uint32{0}
    var v5 uint64 = 1
    v6 := make(map[int]bool)
    for fi7 := 0; fi7 < 2; fi7++ {
        var v8 bool = (1.5 <= (1e10 / 1.5))
        v3 -= 211
        _ = fi7
        _ = v8
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v9 uint8 = (v3 >> ((18446744073709551615) & 63))
    fmt.Println(v3, v9)
    for fi10 := 0; fi10 < 4; fi10++ {
        var v11 int32 = 0
        v9 += (^(v3))
        _ = fi10
        _ = v11
    }
    fmt.Println(v5)
    v3 += (v9 - (v3 + (v3 ^ (v9 << ((0) & 63)))))
    var v12 int8 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
