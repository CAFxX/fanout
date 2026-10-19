package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int8; F1 string }

func main() {
    v1 := make(map[uint32]string)
    var v2 int64 = 9223372036854775807
    var v3 uint16 = 32767
    v4 := []uint32{4294967295, 596, 814}
    v5 := make(map[uint32]int)
    v6 := make(map[string]uint32)
    v7 := GenS1{F0: 65535}
    if true {
        fmt.Println(v3, v3)
        var v8 int64 = v2
        _ = v8
    } else {
        var v9 uint64 = 18446744073709551615
        var v10 string = "a"
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = v10
        _ = v9
    }
    var v11 int32 = 0
    v12 := GenS2{F0: 94, F1: "a"}
    fmt.Println(v3, v3)
    _ = len(v4)
    var v13 int8 = 70
    v14 := GenS2{F0: v13, F1: (("hello" + "foo bar") + "")}
    fmt.Println(v2, v2, v11)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
