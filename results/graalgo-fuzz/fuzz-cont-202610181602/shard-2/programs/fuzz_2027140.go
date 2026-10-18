package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int64; F1 uint64; F2 string }
type GenS3 struct { F0 uint8 }

func main() {
    var v1 uint8 = 153
    v2 := make(map[uint32]int)
    var v3 string = ("hello" + "hello")
    var v4 int64 = 270
    v5 := make(map[int]int)
    var v6 float64 = (0.1 * -2.25)
    var v7 uint64 = 1
    v7 = v7
    var v8 int16 = -32768
    v9 := GenS1{F0: uint32(v1)}
    fmt.Println(v3, v1, v7)
    if true {
        for fi10 := 0; fi10 < 8; fi10++ {
            var v11 int8 = 127
            var v12 int64 = (v4 / ((v4) | 1))
            _ = fi10
            _ = v11
            _ = v12
        }
        var v13 bool = (v1 > 166)
        _ = v13
    }
    var v14 int8 = 107
    fmt.Println(v1, v8, v4)
    v5[9223372036854775807] = 2147483647
    v15 := GenS2{F0: v4, F1: v7, F2: "\n\t"}
    if false {
        if ((-(v8)) != -32768) {
            v16 := GenS3{F0: v1}
            _ = v16
        } else {
            var v17 uint16 = 787
            _ = v17
        }
        v18 := []int{32767, 2147483647}
        delete(v5, -9223372036854775808)
        _ = v18
    }
    var v19 float32 = float32(v6)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v14
    _ = v15
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
