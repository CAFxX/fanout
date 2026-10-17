package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint; F2 int16 }
func genFn0(p0 uint32, p1 uint) int64 {
    v6 := GenS1{F0: 790, F1: 65535, F2: -1}
    var v7 uint = 18446744073709551615
    p1 = v7
    fmt.Println(p0, p0, p0)
    if false {
        fmt.Println(p0)
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 1
}

func main() {
    var v1 int16 = 32767
    v2 := []uint32{42, 0, 2147483648, 702}
    var v3 uint64 = 832
    var v4 uint32 = 4294967295
    v5 := []string{}
    v1 = v1
    var v8 int16 = int16(v4)
    v4 = v4
    v2 = append(v2, v4)
    _ = len(v5)
    if true {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    }
    v8 = v1
    var v9 int32 = 391
    fmt.Println(v9)
    var v10 bool = (127 <= 128)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
