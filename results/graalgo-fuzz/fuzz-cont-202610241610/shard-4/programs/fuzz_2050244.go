package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool }

func main() {
    v1 := []string{}
    var v2 string = "hello"
    var v3 string = ((((("foo bar" + v2) + v2) + (v2 + v2)) + (v2 + "hello")) + v2)
    v4 := []uint32{65535, 153, 4294967295}
    var v5 string = ((v3 + v2) + "\n\t")
    v6 := make(map[uint32]bool)
    if true {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        v7 := GenS1{F0: 1, F1: true}
        for fi8 := 0; fi8 < 2; fi8++ {
            fmt.Println(v3, v2)
            _ = fi8
        }
        _ = v7
    } else {
        var v9 int8 = -128
        v5 = "hello"
        _ = v9
    }
    v2 = v5
    var v10 uint32 = 4294967295
    if true {
        if len(v1) > 0 {
            v1[0] = v5
        }
        v11 := GenS1{F0: (v10 << ((225) & 63)), F1: false}
        _ = v11
    }
    for fi12 := 0; fi12 < 14; fi12++ {
        var v13 float32 = 1.5
        v13 = -2.25
        _ = fi12
        _ = v13
    }
    var v14 int = -256
    v15 := GenS1{F0: (v10 ^ 2147483647), F1: false}
    var v16 int = 0
    var v17 int32 = -1
    v18 := GenS1{F0: v10, F1: (v17 <= 2147483647)}
    v19 := GenS1{F0: (v10 % (((v10 / ((v10) | 1))) | 1)), F1: (true && (!(42 >= 362)))}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
