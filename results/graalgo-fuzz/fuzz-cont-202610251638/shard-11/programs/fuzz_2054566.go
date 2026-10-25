package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32; F2 int }

func main() {
    var v1 string = ""
    v2 := make(map[uint32]bool)
    var v3 int64 = -221
    var v4 string = v1
    var v5 uint = 0
    for fi6 := 0; fi6 < 2; fi6++ {
        if true {
            var v7 int = (fi6 | 138)
            _ = v7
        } else {
            v5 = 576
        }
        _ = fi6
    }
    v4 = (((v1 + (v4 + "hello")) + (((v4 + v4) + v1) + v4)) + v4)
    var v8 uint8 = 2
    for fi9 := 0; fi9 < 10; fi9++ {
        fmt.Println(v5)
        fmt.Println(v4, v8, v4)
        v5 += (^(v5))
        _ = fi9
    }
    v2[65535] = false
    v4 = v1
    v4 = v4
    v8 -= (v8 / ((107) | 1))
    var v10 uint8 = (v8 & v8)
    var v11 uint = 2
    v10 = (v10 >> ((v5) & 63))
    if false {
        v12 := GenS1{F0: v1, F1: 42, F2: 128}
        v4 = ("foo bar" + v4)
        _ = v12
    }
    var v13 float32 = (1e10 * 0.1)
    fmt.Println(v8, v4)
    fmt.Println(v11, v11)
    var v14 uint16 = 65535
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
