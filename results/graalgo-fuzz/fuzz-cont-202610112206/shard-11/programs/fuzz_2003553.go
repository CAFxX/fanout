package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint }
func genFn0(p0 uint32, p1 int) uint64 {
    p0 -= p0
    fmt.Println(p0)
    v6 := GenS1{F0: 100, F1: 65535}
    if true {
        if (false || ((!false) || (int32(p0) == 42))) {
            v7 := []int{p1, 127}
            _ = v7
        }
        p1 += 4294967295
    }
    _ = p0
    _ = p1
    _ = v6
    return 715
}

func main() {
    v1 := make(map[string]int)
    var v2 uint16 = 721
    var v3 int32 = 0
    var v4 int32 = (v3 | v3)
    var v5 bool = false
    switch ("hello" + (("hello" + ("\n\t" + "hello")) + "foo bar")) {
    case "hello":
        v5 = ((v3 ^ -1) <= v3)
    default:
        var v8 uint32 = 501
        _ = v8
    }
    var v9 uint16 = v2
    var v10 uint = 4294967295
    var v11 int64 = -422
    var v12 uint16 = 0
    v12 += v2
    var v13 int8 = 127
    for fi14 := 0; fi14 < 10; fi14++ {
        v4 = 42
        var v15 uint16 = (v12 * v2)
        _ = fi14
        _ = v15
    }
    v16 := []int{}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
