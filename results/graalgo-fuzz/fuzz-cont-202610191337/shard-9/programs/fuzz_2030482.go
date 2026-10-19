package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 bool; F1 uint32; F2 string }

func main() {
    v1 := make(map[string]string)
    var v2 int64 = 0
    var v3 uint = 65535
    var v4 uint = v3
    fmt.Println(v4, v2, v3)
    var v5 uint16 = 65535
    v6 := GenS1{F0: "\n\t"}
    v3 += (v3 >> ((v4) & 63))
    var v7 uint = v4
    switch "\n\t" {
    case "":
        fmt.Println(v4, v5, v5)
    }
    var v8 uint = (v4 & (v7 & 0))
    v7 += (v3 / (((v7 % (((v3 + v3)) | 1))) | 1))
    for fi9 := 0; fi9 < 12; fi9++ {
        var v10 int64 = (v2 - v2)
        v11 := GenS2{F0: true, F1: 0, F2: ("a" + ((("hello" + "a") + "hello") + "foo bar"))}
        _ = fi9
        _ = v10
        _ = v11
    }
    var v12 uint16 = (v5 | 1)
    for fi13 := 0; fi13 < 16; fi13++ {
        delete(v1, "a")
        var v14 uint16 = 294
        v15 := []int{}
        _ = fi13
        _ = v14
        _ = v15
    }
    var v16 int32 = 0
    fmt.Println(v5, v8, v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
