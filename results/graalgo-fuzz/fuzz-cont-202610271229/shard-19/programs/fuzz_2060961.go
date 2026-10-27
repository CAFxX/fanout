package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64 }
func genFn0(p0 int16) int64 {
    p0 = -32768
    var v5 uint8 = 0
    v6 := GenS1{F0: 4294967295, F1: int64(p0)}
    v7 := GenS1{F0: 251, F1: 1}
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    return 1
}

func main() {
    v1 := []string{"foo bar", "hello", ("foo bar" + "hello")}
    var v2 int64 = 1
    var v3 uint16 = 32767
    v4 := make(map[uint32]int)
    var v8 uint = 2
    var v9 int8 = 1
    var v10 int16 = 42
    fmt.Println(v9)
    var v11 int32 = 870
    fmt.Println(v10, v3)
    switch ("" + "\n\t") {
    case "\n\t":
        v12 := v4[2147483648]
        _ = v12
    default:
        v10 += (v10 % ((int16(v9)) | 1))
    }
    fmt.Println(v2, v3, v10)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
