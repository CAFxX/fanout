package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8; F2 uint }

func main() {
    var v1 string = (("hello" + "hello") + "hello")
    var v2 int64 = 0
    v3 := make(map[uint32]bool)
    fmt.Println(v1, v1)
    var v4 int32 = 1
    var v5 int8 = 1
    v5 = (v5 & (v5 | (v5 & v5)))
    var v6 uint = 1
    var v7 int8 = (^(v5))
    var v8 uint32 = 65535
    if (127 >= 157) {
        v9 := []int{-414, 127}
        delete(v3, 65535)
        _ = v9
    }
    var v10 uint32 = (v8 / ((42) | 1))
    var v11 int8 = -49
    fmt.Println(v11, v8, v1)
    v12 := GenS1{F0: (v6 % ((v6) | 1)), F1: (^(v7)), F2: (v6 | (v6 << (((v6 + v6)) & 63)))}
    v4 -= v4
    var v13 bool = (85 != 127)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
