package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64; F2 uint64 }
func genFn0(p0 bool, p1 int32) bool {
    if (!((2 > 2) && (true && p0))) {
        for fi5 := 0; fi5 < 13; fi5++ {
            v6 := GenS1{F0: 378, F1: 229, F2: 1}
            v7 := []int{}
            _ = len(v7)
            _ = fi5
            _ = v6
            _ = v7
        }
        if true {
            var v8 int16 = 42
            fmt.Println(p0)
            v9 := []int{}
            _ = v8
            _ = v9
        }
    }
    p0 = (p1 >= 0)
    p0 = (!true)
    _ = p0
    _ = p1
    return (true || p0)
}

func main() {
    var v1 uint16 = 65535
    var v2 int = 32767
    var v3 uint8 = 2
    v4 := make(map[string]bool)
    fmt.Println(v2, v3, v2)
    v4["a"] = (true && false)
    v2 = (v2 | v2)
    v10 := v4["foo bar"]
    var v11 bool = (!(42 >= 1))
    v12 := []int{(v2 * (v2 | v2)), v2, v2}
    var v13 uint8 = (v3 & (v3 & v3))
    if v10 {
        var v14 int32 = 0
        var v15 int = (v2 / ((-915) | 1))
        _ = v14
        _ = v15
    }
    v16 := GenS1{F0: v2, F1: 18446744073709551615, F2: 18446744073709551615}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
