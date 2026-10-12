package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint16 }

func main() {
    var v1 string = ("a" + "")
    var v2 float64 = (3.14159 - -2.25)
    var v3 int = 687
    v4 := make(map[int]string)
    v5 := []uint32{65535, 151, 65535}
    fmt.Println(v1, v1)
    var v6 int64 = 0
    var v7 float64 = ((1e10 + ((v2 - v2) * 3.14159)) + v2)
    fmt.Println(v6)
    var v8 uint64 = 1
    var v9 int16 = 1000
    v10 := GenS1{F0: 42}
    v11 := GenS2{F0: 1}
    var v12 uint8 = 177
    var v13 float64 = (((1e10 * (v7 * float64(3.14159))) / v7) * v2)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
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
    _ = v9
}
