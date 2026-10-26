package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8; F2 int64 }
func genFn0(p0 uint8, p1 int64) int64 {
    if (true || (true && false)) {
        var v5 string = ("" + "\n\t")
        p1 += -336
        var v6 int8 = 127
        _ = v5
        _ = v6
    }
    var v7 uint64 = 997
    _ = p0
    _ = p1
    _ = v7
    return (p1 << ((255) & 63))
}

func main() {
    var v1 int64 = -1
    v2 := []string{}
    v3 := []int{}
    v4 := []string{(("foo bar" + (("a" + "hello") + "a")) + "hello")}
    var v8 uint64 = 18
    v1 = (v1 >> ((311) & 63))
    var v9 uint32 = 4294967295
    v8 = v8
    fmt.Println(v9, v1, v1)
    v9 -= (v9 / ((v9) | 1))
    fmt.Println(v8)
    v10 := GenS1{F0: (v9 << ((38) & 63)), F1: -1, F2: (-(v1))}
    for fi11 := 0; fi11 < 3; fi11++ {
        v9 = v9
        for fi12 := 0; fi12 < 14; fi12++ {
            _ = len(v2)
            v1 += v1
            _ = fi12
        }
        _ = fi11
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
