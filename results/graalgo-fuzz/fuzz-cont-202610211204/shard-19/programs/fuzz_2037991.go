package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8 }

func main() {
    v1 := make(map[int]string)
    v2 := make(map[uint32]string)
    v3 := []uint32{709}
    var v4 uint32 = 4294967295
    for fi5 := 0; fi5 < 10; fi5++ {
        var v6 string = ("\n\t" + "a")
        _ = fi5
        _ = v6
    }
    var v7 int8 = 100
    v2[v4] = "hello"
    v7 = -6
    fmt.Println(v7, v4)
    if (-257 > 338) {
        v8 := GenS1{F0: false, F1: v7}
        v7 -= v7
        var v9 int32 = 223
        _ = v8
        _ = v9
    } else {
        for fi10 := 0; fi10 < 13; fi10++ {
            fmt.Println(v7, v4)
            fmt.Println(v4)
            _ = fi10
        }
    }
    fmt.Println(v7)
    fmt.Println(v7)
    v11 := GenS1{F0: (v7 >= v7), F1: v7}
    for fi12 := 0; fi12 < 15; fi12++ {
        v7 -= v7
        var v13 string = "hello"
        _ = fi12
        _ = v13
    }
    if len(v3) > 0 {
        v3[0] = (^(v4))
    }
    if ("" == "\n\t") {
        var v14 int = -1
        _ = v14
    } else {
        var v15 int16 = 1000
        _ = v15
    }
    v4 = v4
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v7
}
