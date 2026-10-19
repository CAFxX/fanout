package main

import (
    "fmt"
)

func genFn0() uint16 {
    var v6 float64 = -2.25
    var v7 int32 = 42
    _ = v6
    _ = v7
    return 255
}

func main() {
    v1 := make(map[uint32]int)
    var v2 uint16 = 385
    var v3 uint16 = 65535
    v4 := []uint32{}
    var v5 uint8 = 85
    var v8 uint64 = 18446744073709551615
    v2 = (v2 / (((v2 << ((18446744073709551615) & 63))) | 1))
    var v9 uint32 = 0
    fmt.Println(v8, v5)
    v9 -= (v9 + v9)
    _ = len(v1)
    if false {
        fmt.Println(v8, v9, v9)
    } else {
        for fi10 := 0; fi10 < 10; fi10++ {
            var v11 int8 = -1
            _ = fi10
            _ = v11
        }
    }
    fmt.Println(v3)
    switch 0 {
    case 32768:
        v5 = (v5 & v5)
    case 127:
        v5 = 127
    case 1:
        for fi12 := 0; fi12 < 1; fi12++ {
            var v13 int8 = -1
            _ = fi12
            _ = v13
        }
    default:
        fmt.Println(v5)
    }
    var v14 int8 = -1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
