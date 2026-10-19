package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint; F1 int }

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[string]string)
    var v3 int32 = -2147483648
    v4 := []string{"foo bar", "\n\t"}
    var v5 uint = 577
    var v6 uint64 = 1
    var v7 uint64 = 4294967295
    var v8 uint64 = (v6 << (((v5 >> ((v5) & 63))) & 63))
    v4 = append(v4, (("hello" + "") + ((("hello" + "foo bar") + "foo bar") + "")))
    var v9 int = 2147483647
    fmt.Println(v8, v3, v5)
    delete(v1, 4294967295)
    if false {
        v6 += v6
    } else {
        var v10 bool = true
        var v11 uint32 = uint32(v5)
        _ = v10
        _ = v11
    }
    if len(v4) > 0 {
        v4[0] = "hello"
    }
    v7 -= v8
    v3 += v3
    v12 := GenS1{F0: v3}
    v13 := GenS2{F0: (v5 >> ((42) & 63)), F1: v9}
    v6 -= 18446744073709551615
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
