package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
func genFn0() uint64 {
    var v6 int64 = 0
    fmt.Println(v6)
    var v7 uint32 = 65535
    v7 -= v7
    v7 -= v7
    _ = v6
    _ = v7
    return 533
}

func main() {
    var v1 int16 = 209
    var v2 int16 = (^(v1))
    v3 := make(map[int]string)
    v4 := []string{}
    v5 := make(map[uint32]bool)
    v2 = v2
    var v8 uint8 = 0
    v8 += 72
    var v9 uint32 = 42
    for fi10 := 0; fi10 < 15; fi10++ {
        v9 += v9
        _ = fi10
    }
    v1 = v2
    var v11 int64 = 0
    v12 := v3[-9223372036854775808]
    delete(v3, -464)
    fmt.Println(v12, v12, v1)
    v13 := GenS1{F0: 2}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
