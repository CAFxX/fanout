package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint8 }

func main() {
    var v1 uint64 = 4294967295
    var v2 int64 = 0
    v3 := make(map[int]string)
    var v4 uint64 = v1
    var v5 string = "\n\t"
    v6 := GenS1{F0: 2147483647}
    v5 = v5
    var v7 int8 = 84
    v8 := GenS2{F0: 0}
    fmt.Println(v1, v7, v7)
    for fi9 := 0; fi9 < 16; fi9++ {
        v7 += v7
        _ = fi9
    }
    v10 := []int{127, 127, 4294967295, 42}
    v11 := v3[0]
    switch (v5 + ((v11 + (("" + v5) + "foo bar")) + v11)) {
    case "foo bar":
        v7 += 42
    default:
        v2 = 9223372036854775807
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
