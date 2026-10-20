package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }

func main() {
    v1 := []int{}
    var v2 int64 = -770
    v3 := []uint32{303}
    var v4 string = ("hello" + ("hello" + ("a" + ((("foo bar" + "") + "hello") + "a"))))
    var v5 uint64 = 18446744073709551615
    for fi6 := 0; fi6 < 6; fi6++ {
        v7 := GenS1{F0: 0}
        var v8 uint64 = v5
        v9 := GenS1{F0: -1}
        _ = fi6
        _ = v7
        _ = v8
        _ = v9
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v10 float64 = 1e10
    fmt.Println(v4, v2, v4)
    var v11 uint64 = v5
    fmt.Println(v10, v11)
    var v12 uint8 = 42
    v13 := GenS1{F0: 301}
    v5 = v11
    v10 = float64((0.0 + 0.1))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
