package main

import (
    "fmt"
)

func genFn0(p0 int, p1 bool, p2 uint8) int32 {
    fmt.Println(p0, p0)
    p2 += p2
    var v5 int16 = -984
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return -388
}

func main() {
    v1 := []int{0, 32768}
    var v2 uint32 = 2
    var v3 string = ("foo bar" + "foo bar")
    var v4 uint32 = 42
    v3 = v3
    var v6 int64 = -9223372036854775808
    if len(v1) > 0 {
        v1[0] = 127
    }
    var v7 int16 = -32768
    var v8 int = 521
    var v9 int8 = -121
    var v10 bool = true
    v1 = append(v1, (v8 ^ (v8 * (v8 / (((v8 % (((v8 >> ((255) & 63))) | 1))) | 1)))))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
