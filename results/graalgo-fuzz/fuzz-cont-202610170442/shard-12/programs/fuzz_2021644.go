package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8 }

func main() {
    var v1 uint = 255
    var v2 uint64 = 56
    v3 := make(map[uint32]string)
    v4 := make(map[string]bool)
    v5 := []int{9223372036854775807, 32768, 127, 9223372036854775807}
    for fi6 := 0; fi6 < 10; fi6++ {
        v1 -= uint(v2)
        if true {
            v7 := GenS1{F0: "hello", F1: uint8(v2)}
            _ = v7
        }
        for fi8 := 0; fi8 < 10; fi8++ {
            var v9 float64 = 0.0
            var v10 uint = v1
            _ = fi8
            _ = v10
            _ = v9
        }
        _ = fi6
    }
    v2 = (^(v2))
    if ((!true) && (true || false)) {
        fmt.Println(v2, v1)
        if true {
            if len(v5) > 0 {
                v5[0] = 42
            }
            v11 := GenS1{F0: "hello", F1: 42}
            _ = v11
        }
    } else {
        var v12 int16 = 1000
        _ = v12
    }
    for fi13 := 0; fi13 < 11; fi13++ {
        v1 = v1
        v4[(("hello" + "\n\t") + "hello")] = false
        _ = fi13
    }
    var v14 uint32 = 65535
    v2 -= (^(v2))
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    var v15 uint64 = v2
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
