package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32 }

func main() {
    var v1 string = "\n\t"
    var v2 uint64 = 0
    var v3 string = v1
    var v4 int16 = -32768
    v5 := make(map[string]uint32)
    var v6 string = (("a" + v3) + ((v3 + v1) + "a"))
    v7 := []int{9223372036854775807, -1, -2147483648}
    fmt.Println(v3)
    v1 = "a"
    for fi8 := 0; fi8 < 17; fi8++ {
        v9 := GenS1{F0: 361, F1: 65535}
        _ = fi8
        _ = v9
    }
    fmt.Println(v2)
    if (-2.25 == 1.5) {
        v10 := GenS1{F0: 65535, F1: 2147483647}
        _ = v10
    } else {
        v5[("hello" + ((((("a" + ("foo bar" + (v6 + "foo bar"))) + v1) + "hello") + v3) + v6))] = 42
    }
    if false {
        v6 = v3
        var v11 string = "\n\t"
        _ = v11
    } else {
        v6 = (v3 + v3)
        fmt.Println(v1, v2)
    }
    v12 := GenS1{F0: 811, F1: 4294967295}
    var v13 string = v6
    var v14 int64 = -1
    fmt.Println(v4, v4, v3)
    var v15 uint16 = 51
    var v16 float32 = 1.5
    var v17 uint32 = 985
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
