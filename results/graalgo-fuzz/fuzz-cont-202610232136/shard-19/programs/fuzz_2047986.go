package main

import (
    "fmt"
)


func main() {
    v1 := make(map[string]string)
    v2 := make(map[int]bool)
    v3 := make(map[uint32]string)
    var v4 uint = 18446744073709551615
    fmt.Println(v4)
    fmt.Println(v4)
    v4 -= 65535
    var v5 int8 = 1
    v6 := v3[1]
    v7 := []int{9223372036854775807, 32768, 305, -248}
    v4 = v4
    var v8 int16 = 1000
    v2[935] = true
    fmt.Println(v5, v5, v4)
    switch 32767 {
    case 4294967295:
        v5 = v5
    case -9223372036854775808:
        v3[745] = "foo bar"
    }
    if false {
        delete(v2, 127)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
