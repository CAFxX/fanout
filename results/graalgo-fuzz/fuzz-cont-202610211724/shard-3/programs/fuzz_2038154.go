package main

import (
    "fmt"
)

func genFn0(p0 string) int16 {
    var v7 uint64 = 2
    var v8 int64 = int64(v7)
    _ = p0
    _ = v7
    _ = v8
    return -540
}

func main() {
    var v1 string = "foo bar"
    var v2 int32 = 0
    var v3 int32 = (-(v2))
    var v4 float64 = 3.14159
    var v5 uint8 = 1
    v6 := make(map[int]string)
    var v9 string = v1
    _ = len(v6)
    var v10 bool = true
    fmt.Println(v10, v4, v1)
    var v11 int64 = 9223372036854775807
    var v12 int = 127
    var v13 int16 = -57
    v12 += (v12 + v12)
    var v14 int32 = v3
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
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
    _ = v9
}
