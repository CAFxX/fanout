package main

import (
    "fmt"
)

type GenS1 struct { F0 string }

func main() {
    var v1 int16 = 0
    var v2 uint = 65535
    var v3 int = 9223372036854775807
    var v4 int32 = -327
    v5 := []string{("foo bar" + (("foo bar" + "foo bar") + "foo bar")), "foo bar"}
    v6 := make(map[uint32]uint32)
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    var v7 int64 = 0
    var v8 int8 = -1
    var v9 uint64 = 0
    v10 := GenS1{F0: ""}
    v5 = append(v5, ("foo bar" + ("" + "hello")))
    var v11 int32 = v4
    var v12 int64 = v7
    delete(v6, 40)
    for fi13 := 0; fi13 < 2; fi13++ {
        v1 = v1
        v12 += -9223372036854775808
        fmt.Println(v3)
        _ = fi13
    }
    if true {
        fmt.Println(v4, v2)
        _ = len(v5)
        var v14 bool = true
        _ = v14
    }
    v4 += (v4 + (v11 | (-(v11))))
    v3 = (^(v3))
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
