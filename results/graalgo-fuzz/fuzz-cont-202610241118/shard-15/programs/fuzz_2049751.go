package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }

func main() {
    v1 := []uint32{42, 547, 65535}
    var v2 uint8 = 1
    var v3 string = ""
    v4 := make(map[int]bool)
    v3 = "hello"
    if false {
        v3 = ((v3 + "") + v3)
    }
    v2 = v2
    var v5 int32 = -295
    v5 = (v5 % ((42) | 1))
    for fi6 := 0; fi6 < 5; fi6++ {
        _ = len(v1)
        v2 += (v2 - (v2 | v2))
        v2 -= (^(v2))
        _ = fi6
    }
    for fi7 := 0; fi7 < 10; fi7++ {
        _ = len(v1)
        v2 -= 255
        _ = fi7
    }
    fmt.Println(v5, v3)
    fmt.Println(v2, v2)
    fmt.Println(v2, v3)
    v8 := GenS1{F0: 42}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
