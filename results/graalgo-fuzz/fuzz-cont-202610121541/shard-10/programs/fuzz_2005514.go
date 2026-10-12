package main

import (
    "fmt"
)

func genFn0(p0 int64, p1 int8, p2 string) string {
    if (false && (!true)) {
        p2 = "hello"
    } else {
        p0 = p0
    }
    fmt.Println(p2, p2, p1)
    _ = p0
    _ = p1
    _ = p2
    return p2
}

func main() {
    var v1 uint32 = 0
    var v2 int32 = -1
    var v3 uint16 = 42
    var v4 string = "a"
    v2 = int32(v1)
    v3 -= 1
    var v5 float32 = (3.14159 / (1e10 - (3.14159 - 0.1)))
    fmt.Println(v3, v1)
    fmt.Println(v1, v4)
    v6 := []int{32767}
    var v7 int8 = 42
    v4 = v4
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
