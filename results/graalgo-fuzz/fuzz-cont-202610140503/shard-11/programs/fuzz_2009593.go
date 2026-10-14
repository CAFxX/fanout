package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }

func main() {
    var v1 uint64 = 1
    var v2 string = ""
    var v3 uint16 = 65535
    v4 := make(map[uint32]bool)
    var v5 float32 = 1.5
    v6 := make(map[int]string)
    var v7 int16 = 818
    for fi8 := 0; fi8 < 12; fi8++ {
        v9 := GenS1{F0: 38}
        _ = fi8
        _ = v9
    }
    v3 -= (v3 / ((v3) | 1))
    for fi10 := 0; fi10 < 6; fi10++ {
        fmt.Println(v7)
        _ = fi10
    }
    var v11 uint = 1
    v12 := GenS1{F0: 76}
    fmt.Println(v3, v3)
    delete(v6, -2147483648)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
