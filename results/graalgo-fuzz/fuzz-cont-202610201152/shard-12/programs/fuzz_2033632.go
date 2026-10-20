package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8; F2 string }

func main() {
    var v1 int32 = -2147483648
    v2 := make(map[string]bool)
    var v3 int16 = 1000
    var v4 string = (("hello" + "\n\t") + ("" + "\n\t"))
    v5 := make(map[int]uint32)
    var v6 int64 = 1
    var v7 uint16 = 729
    v3 = 32767
    fmt.Println(v4)
    fmt.Println(v6, v1)
    var v8 int8 = 127
    var v9 uint16 = (v7 << ((485) & 63))
    var v10 uint16 = v9
    var v11 uint = 4294967295
    v12 := GenS1{F0: (v11 + v11), F1: 255, F2: v4}
    v2[(v4 + v4)] = (3.14159 != 3.14159)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
