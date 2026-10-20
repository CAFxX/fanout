package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint; F2 string }
type GenS2 struct { F0 int16; F1 uint8 }
type GenS3 struct { F0 uint32 }
func genFn0() int16 {
    var v4 int = 127
    var v5 int64 = 1
    var v6 int8 = 0
    var v7 int8 = 100
    v8 := GenS1{F0: 419, F1: 2, F2: ""}
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return 42
}

func main() {
    v1 := []int{127, 9223372036854775807, -1}
    v2 := []string{"hello", ("hello" + ("" + "hello"))}
    var v3 uint32 = 440
    v9 := GenS2{F0: -934, F1: 2}
    var v10 int8 = -128
    var v11 uint = 18446744073709551615
    v11 -= (v11 ^ v11)
    var v12 float64 = 1e10
    _ = len(v1)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if len(v1) > 0 {
        v1[0] = 0
    }
    for fi13 := 0; fi13 < 4; fi13++ {
        var v14 int16 = 117
        v15 := GenS3{F0: v3}
        _ = fi13
        _ = v14
        _ = v15
    }
    for fi16 := 0; fi16 < 10; fi16++ {
        fmt.Println(v10)
        _ = fi16
    }
    if false {
        _ = len(v1)
        var v17 int8 = (v10 | v10)
        if len(v1) > 0 {
            v1[0] = -1
        }
        _ = v17
    }
    v3 = v3
    fmt.Println(v12, v12, v3)
    fmt.Println(v3, v3)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v9
}
