package main

import (
    "fmt"
)

func genFn0(p0 string) uint16 {
    var v5 uint8 = 127
    v6 := []int{}
    switch "" {
    case "hello":
        var v7 int16 = 32767
        _ = v7
    default:
        p0 = p0
    }
    var v8 int32 = 1
    p0 = p0
    _ = p0
    _ = v5
    _ = v6
    _ = v8
    return 255
}

func main() {
    var v1 uint16 = 65535
    var v2 bool = true
    var v3 int8 = 116
    v4 := make(map[string]int)
    fmt.Println(v1, v1)
    fmt.Println(v3)
    _ = len(v4)
    v1 += 42
    v2 = (-936 <= -1)
    if v2 {
        v9 := []int{2147483647}
        _ = v9
    } else {
        var v10 uint64 = 2
        _ = v10
    }
    var v11 int8 = v3
    fmt.Println(v11, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
}
