package main

import (
    "fmt"
)

type GenS1 struct { F0 int }

func main() {
    v1 := make(map[int]int)
    v2 := make(map[string]bool)
    var v3 uint32 = 279
    var v4 int32 = -361
    v5 := []uint32{}
    v3 += (^(v3))
    var v6 float32 = 1e10
    v7 := GenS1{F0: 0}
    fmt.Println(v4, v3)
    v8 := GenS1{F0: -1}
    for fi9 := 0; fi9 < 19; fi9++ {
        if (831 < 4294967295) {
            var v10 uint32 = v3
            v4 = 2147483647
            _ = v10
        }
        v2[(("a" + "") + "foo bar")] = false
        _ = fi9
    }
    var v11 string = "\n\t"
    v6 = v6
    fmt.Println(v11, v6, v3)
    fmt.Println(v6, v11, v6)
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    var v12 uint16 = 255
    _ = len(v1)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
