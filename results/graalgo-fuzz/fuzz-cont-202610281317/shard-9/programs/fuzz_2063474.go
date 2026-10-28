package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int64 }
func genFn0(p0 int) int {
    fmt.Println(p0)
    p0 -= (p0 << ((421) & 63))
    _ = p0
    return (p0 / ((p0) | 1))
}

func main() {
    v1 := make(map[string]string)
    var v2 int32 = 42
    var v3 uint64 = 0
    var v4 uint = 18446744073709551615
    var v5 uint = (v4 % (((v4 & v4)) | 1))
    var v6 string = "hello"
    v7 := GenS1{F0: 0}
    v8 := GenS1{F0: -32768}
    fmt.Println(v4, v4)
    v9 := GenS2{F0: -1}
    var v10 int8 = 100
    fmt.Println(v4, v5)
    v4 = v4
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
