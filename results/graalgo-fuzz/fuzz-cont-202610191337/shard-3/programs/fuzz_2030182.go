package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
func genFn0() int16 {
    for fi7 := 0; fi7 < 11; fi7++ {
        var v8 uint8 = 102
        var v9 int32 = -1
        _ = fi7
        _ = v8
        _ = v9
    }
    var v10 int8 = -1
    _ = v10
    return 42
}

func main() {
    v1 := make(map[int]bool)
    v2 := []string{("foo bar" + "foo bar"), "", ("a" + "foo bar"), ("a" + "\n\t")}
    v3 := make(map[string]bool)
    v4 := make(map[string]int)
    v5 := make(map[uint32]uint32)
    var v6 int8 = 100
    v6 = v6
    v6 -= v6
    v6 = (v6 + v6)
    if (!true) {
        v6 = (v6 - v6)
    }
    if len(v2) > 0 {
        v2[0] = "hello"
    }
    var v11 int64 = 852
    v6 = (-(v6))
    v5[2147483648] = 65535
    if true {
        v12 := GenS1{F0: (1 != -268)}
        _ = v12
    } else {
        v11 = (-(v11))
    }
    var v13 uint64 = 93
    _ = len(v2)
    _ = len(v3)
    var v14 int32 = 940
    var v15 uint32 = 2147483647
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
