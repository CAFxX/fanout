package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
func genFn0(p0 uint16, p1 int16) uint64 {
    p1 += 297
    for fi6 := 0; fi6 < 2; fi6++ {
        if (true || false) {
            v7 := GenS1{F0: 0}
            v8 := GenS1{F0: 2147483648}
            p1 = (p1 - p1)
            _ = v7
            _ = v8
        }
        _ = fi6
    }
    var v9 uint16 = (^(p0))
    _ = p0
    _ = p1
    _ = v9
    return uint64(v9)
}
func genFn1(p0 int64, p1 uint8, p2 int) uint32 {
    switch (p2 & p2) {
    case 42:
        v10 := GenS1{F0: 470}
        _ = v10
    case 4294967295:
        v11 := GenS1{F0: 2147483648}
        _ = v11
    }
    var v12 int16 = 1000
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    return 681
}

func main() {
    var v1 bool = true
    var v2 uint8 = 218
    v3 := make(map[uint32]string)
    var v4 uint8 = v2
    v5 := []int{4294967295, 42}
    if v1 {
        var v13 uint8 = v4
        _ = v13
    }
    var v14 int32 = -655
    if len(v5) > 0 {
        v5[0] = 128
    }
    var v15 int64 = 9223372036854775807
    fmt.Println(v1)
    v2 -= v2
    v4 += 183
    fmt.Println(v14)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
