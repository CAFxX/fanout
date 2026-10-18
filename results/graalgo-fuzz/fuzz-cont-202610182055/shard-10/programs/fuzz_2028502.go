package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }

func main() {
    var v1 uint8 = 55
    var v2 uint8 = (v1 | (v1 ^ v1))
    v3 := make(map[uint32]string)
    var v4 uint64 = 4294967295
    var v5 string = ("" + "foo bar")
    v5 = "a"
    var v6 uint8 = 0
    var v7 int8 = -128
    var v8 int16 = 42
    v9 := []int{32767, 127, -1}
    if ((v4 >> ((2) & 63)) <= 4294967295) {
        var v10 int = 0
        var v11 int = (v10 | v10)
        for fi12 := 0; fi12 < 5; fi12++ {
            v5 = "\n\t"
            _ = fi12
        }
        _ = v10
        _ = v11
    }
    v13 := GenS1{F0: 9223372036854775807}
    v14 := GenS1{F0: 136}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
