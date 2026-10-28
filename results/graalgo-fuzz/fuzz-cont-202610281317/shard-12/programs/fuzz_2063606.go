package main

import (
    "fmt"
)

type GenS1 struct { F0 string }

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint16 = 841
    var v3 float64 = 1e10
    var v4 int8 = 1
    var v5 bool = ((-(v4)) != (v4 | (v4 ^ v4)))
    v5 = ((32767 != (^(v2))) && v5)
    v6 := []int{}
    if (4294967295 < uint64(v2)) {
        fmt.Println(v4)
    }
    fmt.Println(v2, v4)
    var v7 int32 = -2147483648
    v7 += v7
    v8 := GenS1{F0: "hello"}
    fmt.Println(v2, v7)
    for fi9 := 0; fi9 < 17; fi9++ {
        fmt.Println(fi9)
        _ = fi9
    }
    for fi10 := 0; fi10 < 2; fi10++ {
        v11 := GenS1{F0: "\n\t"}
        _ = fi10
        _ = v11
    }
    v12 := GenS1{F0: (("hello" + ("" + "hello")) + "hello")}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
