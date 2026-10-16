package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16 }
func genFn0() uint {
    var v7 uint64 = 576
    v8 := []int{}
    _ = v7
    _ = v8
    return 289
}

func main() {
    v1 := []uint32{65535}
    var v2 int = -2147483648
    v3 := []string{}
    var v4 uint16 = 0
    v5 := []string{}
    var v6 bool = false
    _ = len(v1)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    for fi9 := 0; fi9 < 12; fi9++ {
        for fi10 := 0; fi10 < 19; fi10++ {
            var v11 uint8 = 255
            fmt.Println(v4, v4, v2)
            _ = fi10
            _ = v11
        }
        _ = fi9
    }
    var v12 int = 1
    v12 = v2
    v13 := GenS1{F0: 0, F1: 781}
    v2 = v2
    _ = len(v5)
    var v14 uint8 = 255
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
