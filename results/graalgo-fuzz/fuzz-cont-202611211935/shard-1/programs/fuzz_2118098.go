package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
func genFn0(p0 int16) int32 {
    p0 = p0
    var v6 float64 = 1e10
    fmt.Println(p0, v6)
    _ = p0
    _ = v6
    return 2147483647
}

func main() {
    var v1 uint32 = 517
    var v2 uint64 = 42
    var v3 int32 = -368
    var v4 float64 = 1.5
    v5 := []string{}
    var v7 uint32 = v1
    if len(v5) > 0 {
        v5[0] = "\n\t"
    }
    if false {
        var v8 string = "\n\t"
        v9 := GenS1{F0: uint64(v1)}
        _ = v8
        _ = v9
    } else {
        v5 = append(v5, ((("a" + "a") + "hello") + (("foo bar" + "\n\t") + "hello")))
    }
    v7 = (v7 / (((v1 >> ((0) & 63))) | 1))
    var v10 int64 = -349
    v11 := GenS1{F0: v2}
    fmt.Println(v3)
    v12 := GenS1{F0: v2}
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v3 += 2147483647
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
