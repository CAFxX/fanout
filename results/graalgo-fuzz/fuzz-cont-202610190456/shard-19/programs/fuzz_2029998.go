package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8; F2 string }

func main() {
    var v1 int64 = -540
    var v2 int = -926
    var v3 int64 = v1
    var v4 bool = true
    var v5 int64 = (v1 & -449)
    var v6 uint16 = 42
    v7 := []int{(v2 + 42), (v2 + v2), (^(v2)), (v2 % ((9223372036854775807) | 1))}
    v3 = v1
    switch v2 {
    case 1:
        if true {
            var v8 uint32 = 65535
            v9 := GenS1{F0: 2147483647, F1: -128, F2: "a"}
            var v10 int8 = 42
            _ = v10
            _ = v8
            _ = v9
        }
    default:
        v11 := GenS1{F0: v2, F1: 96, F2: "hello"}
        _ = v11
    }
    for fi12 := 0; fi12 < 15; fi12++ {
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
        _ = fi12
    }
    for fi13 := 0; fi13 < 19; fi13++ {
        fmt.Println(v3)
        _ = fi13
    }
    fmt.Println(v3)
    var v14 float32 = 0.0
    fmt.Println(v3, v1)
    v15 := GenS1{F0: 428, F1: -81, F2: (("\n\t" + "") + "foo bar")}
    if (!v4) {
        for fi16 := 0; fi16 < 8; fi16++ {
            if v4 {
                v14 += v14
            } else {
                var v17 float64 = 1e10
                _ = v17
            }
            _ = fi16
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
