package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
func genFn0() string {
    var v5 float32 = 0.1
    var v6 float32 = v5
    v5 = (v5 - ((v5 + 1.5) / (-2.25 - v6)))
    _ = v5
    _ = v6
    return "\n\t"
}

func main() {
    var v1 int8 = 0
    var v2 uint32 = 371
    var v3 int = 0
    var v4 uint16 = 0
    for fi7 := 0; fi7 < 20; fi7++ {
        v8 := []int{137}
        if false {
            var v9 int64 = -1
            if len(v8) > 0 {
                v8[0] = (^(fi7))
            }
            _ = v9
        }
        _ = fi7
        _ = v8
    }
    switch v3 {
    case 9223372036854775807:
        v4 = v4
    }
    v10 := GenS1{F0: v3}
    v2 = v2
    switch "\n\t" {
    case "a":
        var v11 int = v3
        _ = v11
    case "\n\t":
        v12 := []int{v3, v3, v3, v3}
        _ = v12
    }
    for fi13 := 0; fi13 < 5; fi13++ {
        var v14 uint = 2
        _ = fi13
        _ = v14
    }
    v4 = 2
    var v15 int = v3
    v2 = v2
    v16 := []int{(v3 | (v15 | (v15 % (((-(v3))) | 1))))}
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
