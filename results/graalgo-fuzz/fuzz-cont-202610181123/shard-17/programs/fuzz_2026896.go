package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16; F2 int }
func genFn0() uint {
    var v7 uint = 65535
    v7 = v7
    v7 = (v7 % (((^(v7))) | 1))
    var v8 uint64 = 42
    _ = v7
    _ = v8
    return 42
}

func main() {
    v1 := make(map[string]bool)
    var v2 uint = 1
    v3 := []int{}
    var v4 int32 = 1
    v5 := make(map[int]string)
    var v6 string = ""
    _ = len(v3)
    switch 9223372036854775807 {
    case 4294967295:
        fmt.Println(v2)
    case 2147483647:
        v9 := GenS1{F0: 32767, F1: 2, F2: 128}
        _ = v9
    case 128:
        var v10 uint16 = 32767
        _ = v10
    }
    var v11 string = (v6 + v6)
    var v12 uint64 = 42
    v4 -= (v4 & v4)
    switch v6 {
    case "foo bar":
        v4 -= 814
    }
    v2 = (v2 << ((v2) & 63))
    var v13 int8 = 42
    fmt.Println(v11)
    v14 := GenS1{F0: 0, F1: 242, F2: 128}
    var v15 int32 = (-(v4))
    var v16 int = 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
