package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int; F2 int8 }

func main() {
    var v1 uint8 = 2
    v2 := make(map[string]bool)
    var v3 int8 = 5
    var v4 uint8 = (^(v1))
    v5 := []uint32{65535, 1, 413, 2147483648}
    var v6 float32 = 3.14159
    v7 := GenS1{F0: 66, F1: -2147483648, F2: v3}
    v4 = v4
    v6 = v6
    v8 := v2["\n\t"]
    v2[("\n\t" + ((("\n\t" + "a") + "foo bar") + ""))] = v8
    if false {
        v3 = 127
        v6 += (v6 - v6)
        var v9 uint16 = 42
        _ = v9
    } else {
        fmt.Println(v4, v1)
    }
    fmt.Println(v4)
    var v10 int = 127
    fmt.Println(v1)
    if (1 != 746) {
        for fi11 := 0; fi11 < 9; fi11++ {
            var v12 uint32 = 65535
            _ = fi11
            _ = v12
        }
        if ((("" + "a") + ("foo bar" + "a")) <= ("" + "foo bar")) {
            var v13 uint32 = 2147483647
            fmt.Println(v6)
            _ = v13
        }
    } else {
        var v14 uint64 = 0
        _ = v14
    }
    var v15 uint64 = 4294967295
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
