package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64; F2 string }
func genFn0(p0 string) int32 {
    fmt.Println(p0)
    var v6 int8 = 0
    p0 = p0
    _ = p0
    _ = v6
    return 0
}

func main() {
    v1 := make(map[string]int)
    v2 := []string{"foo bar", ("" + "")}
    var v3 int64 = 9223372036854775807
    v4 := []uint32{65535, 65535, 2}
    v5 := []string{}
    if len(v5) > 0 {
        v5[0] = ""
    }
    v3 = (-(v3))
    var v7 float32 = 1.5
    fmt.Println(v3, v3)
    var v8 int64 = 1
    v8 += 42
    var v9 float32 = v7
    for fi10 := 0; fi10 < 18; fi10++ {
        fmt.Println(v3)
        for fi11 := 0; fi11 < 15; fi11++ {
            var v12 string = ""
            _ = fi11
            _ = v12
        }
        var v13 int = fi10
        _ = fi10
        _ = v13
    }
    fmt.Println(v9, v3, v8)
    v14 := GenS1{F0: 453, F1: (v3 & (v3 * v3)), F2: "foo bar"}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
