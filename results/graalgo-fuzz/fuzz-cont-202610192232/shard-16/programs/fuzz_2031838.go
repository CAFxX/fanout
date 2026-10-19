package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32 }
func genFn0(p0 uint32) int32 {
    var v4 uint32 = p0
    var v5 string = "hello"
    _ = p0
    _ = v4
    _ = v5
    return -2147483648
}

func main() {
    var v1 int64 = 9223372036854775807
    var v2 uint = 236
    v3 := make(map[string]int)
    if false {
        v6 := []int{9223372036854775807}
        v1 = 42
        v7 := GenS1{F0: v2, F1: 0}
        _ = v6
        _ = v7
    }
    fmt.Println(v1, v2)
    var v8 string = ""
    var v9 int8 = 42
    switch "\n\t" {
    case "\n\t":
        switch (v8 + (("a" + "\n\t") + v8)) {
        case "foo bar":
            for fi10 := 0; fi10 < 20; fi10++ {
                v11 := GenS1{F0: (v2 + (v2 ^ (^(v2)))), F1: genFn0(548)}
                _ = fi10
                _ = v11
            }
        }
    case "a":
        fmt.Println(v9, v2)
    default:
        var v12 uint64 = 1
        _ = v12
    }
    var v13 uint16 = 32767
    v14 := GenS1{F0: v2, F1: -2147483648}
    v15 := GenS1{F0: 1, F1: -570}
    fmt.Println(v2, v8)
    v16 := []int{127, -9223372036854775808, -249}
    v17 := GenS1{F0: v2, F1: 1}
    v2 += (v2 ^ v2)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
