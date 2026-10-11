package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }

func main() {
    var v1 int64 = 42
    v2 := []int{}
    v3 := make(map[uint32]uint32)
    var v4 int32 = 0
    var v5 int64 = (v1 >> ((0) & 63))
    var v6 float32 = ((float32(0.0) + -2.25) + -2.25)
    if true {
        v6 -= 1e10
    }
    fmt.Println(v1)
    v4 = (v4 + (v4 << ((449) & 63)))
    v3[1] = 6
    for fi7 := 0; fi7 < 14; fi7++ {
        v8 := GenS1{F0: 942}
        _ = fi7
        _ = v8
    }
    v9 := GenS1{F0: 78}
    if true {
        var v10 string = "hello"
        _ = v10
    }
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
    _ = v9
}
