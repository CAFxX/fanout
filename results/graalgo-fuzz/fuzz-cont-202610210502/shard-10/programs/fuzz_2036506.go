package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8 }
func genFn0() bool {
    var v4 string = "foo bar"
    fmt.Println(v4)
    _ = v4
    return (42 <= 2)
}
func genFn1(p0 uint, p1 string, p2 int32) uint32 {
    var v5 int8 = 127
    v6 := []int{}
    v5 = 42
    var v7 uint64 = 0
    p1 = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    return 42
}

func main() {
    var v1 int8 = -1
    v2 := []uint32{27, 621}
    var v3 uint8 = 1
    v8 := GenS1{F0: -107, F1: v1}
    v9 := GenS1{F0: -1, F1: 127}
    v3 += (v3 * v3)
    if (!false) {
        fmt.Println(v1, v3)
        if (!false) {
            var v10 string = ("foo bar" + ("hello" + "a"))
            for fi11 := 0; fi11 < 2; fi11++ {
                v3 -= 26
                _ = fi11
            }
            _ = v10
        }
        v12 := GenS1{F0: -832, F1: -68}
        _ = v12
    }
    var v13 uint64 = 2
    v13 = v13
    switch (("a" + "foo bar") + "\n\t") {
    case "a":
        v1 = v1
    }
    if true {
        v13 = 18446744073709551615
        v3 = v3
        var v14 bool = (false && false)
        _ = v14
    }
    for fi15 := 0; fi15 < 10; fi15++ {
        _ = len(v2)
        v1 = (v1 - int8(v3))
        _ = fi15
    }
    var v16 uint64 = (v13 >> ((4294967295) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v13
    _ = v16
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
