package main

import (
    "fmt"
)

func genFn0() bool {
    var v4 uint8 = 54
    var v5 int32 = 42
    _ = v4
    _ = v5
    return (true || false)
}

func main() {
    var v1 uint8 = 0
    var v2 string = ""
    v3 := make(map[string]string)
    v2 = ""
    v1 = v1
    var v6 string = v2
    v2 = ("hello" + v2)
    var v7 uint32 = 2
    fmt.Println(v6, v6)
    var v8 int16 = 1000
    var v9 int16 = (v8 * v8)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
