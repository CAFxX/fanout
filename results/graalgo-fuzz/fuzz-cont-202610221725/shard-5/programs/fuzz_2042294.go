package main

import (
    "fmt"
)


func main() {
    v1 := make(map[uint32]string)
    var v2 uint32 = 2147483648
    v3 := make(map[uint32]string)
    var v4 int32 = -2147483648
    var v5 bool = (!true)
    var v6 int8 = -128
    var v7 string = ("hello" + "\n\t")
    v5 = v5
    var v8 int = 2147483647
    var v9 int32 = (v4 >> ((255) & 63))
    fmt.Println(v2)
    _ = len(v3)
    var v10 int16 = 1000
    var v11 float32 = (1e10 * 3.14159)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
