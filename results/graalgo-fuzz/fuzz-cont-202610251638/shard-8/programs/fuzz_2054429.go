package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
func genFn0(p0 string) int64 {
    p0 = p0
    v6 := []int{}
    switch -382 {
    case 1:
        _ = len(v6)
    default:
        var v7 bool = false
        _ = v7
    }
    _ = p0
    _ = v6
    return 9223372036854775807
}

func main() {
    v1 := []uint32{1, 1, 2147483647, 2147483648}
    v2 := make(map[string]uint32)
    v3 := []string{"hello", "", "foo bar"}
    v4 := make(map[int]uint32)
    v5 := make(map[int]uint32)
    var v8 int16 = -1
    delete(v4, 0)
    v9 := v2[("" + "\n\t")]
    v10 := GenS1{F0: 805}
    fmt.Println(v8, v8)
    _ = len(v5)
    for fi11 := 0; fi11 < 14; fi11++ {
        var v12 uint32 = v9
        _ = fi11
        _ = v12
    }
    v9 = 549
    v8 = v8
    if true {
        v13 := GenS1{F0: 1}
        _ = v13
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
