package main

import (
    "fmt"
)


func main() {
    var v1 uint16 = 561
    var v2 uint8 = 251
    v3 := []string{""}
    v4 := make(map[int]string)
    v5 := make(map[int]uint32)
    v1 += (v1 * v1)
    fmt.Println(v1, v1)
    _ = len(v4)
    var v6 uint = 0
    _ = len(v3)
    var v7 int8 = 42
    v1 -= (v1 / (((v1 * v1)) | 1))
    var v8 uint8 = (v2 % ((0) | 1))
    var v9 string = "foo bar"
    var v10 uint16 = (v1 % ((v1) | 1))
    var v11 uint = v6
    fmt.Println(v2, v7, v1)
    var v12 float32 = ((((((3.14159 / (float32(1.5) + 1e10)) - (1e10 / float32(0.1))) / -2.25) + 1e10) / -2.25) + 1e10)
    var v13 int16 = int16(v7)
    var v14 uint16 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
