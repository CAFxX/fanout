package main

import (
    "fmt"
)

type GenS1 struct { F0 int }

func main() {
    var v1 uint32 = 1
    var v2 uint16 = 2
    var v3 bool = true
    var v4 float32 = 0.0
    v5 := []int{-536, 128, 128}
    var v6 uint = 255
    var v7 int64 = 1
    v3 = (486 >= v6)
    for fi8 := 0; fi8 < 19; fi8++ {
        v9 := GenS1{F0: (fi8 ^ 128)}
        _ = fi8
        _ = v9
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    for fi10 := 0; fi10 < 6; fi10++ {
        v11 := GenS1{F0: 445}
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        _ = fi10
        _ = v11
    }
    v6 -= (v6 << (((v6 * v6)) & 63))
    v1 = 2
    v6 = v6
    var v12 float64 = ((float64(v4) * (1.5 * 3.14159)) + 3.14159)
    if ((^(v7)) >= v7) {
        fmt.Println(v2)
        v13 := GenS1{F0: 801}
        _ = v13
    }
    if (true || v3) {
        var v14 bool = (42 == -643)
        _ = v14
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
