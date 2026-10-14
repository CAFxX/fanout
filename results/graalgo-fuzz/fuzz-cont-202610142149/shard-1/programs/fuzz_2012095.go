package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64; F2 uint64 }
func genFn0(p0 int32) int {
    p0 = 35
    p0 += p0
    _ = p0
    return 761
}

func main() {
    v1 := []string{("foo bar" + "foo bar"), "hello", "a"}
    var v2 int8 = 101
    v3 := make(map[uint32]string)
    v4 := make(map[uint32]int)
    var v5 uint32 = 1
    v5 += (^(v5))
    for fi6 := 0; fi6 < 1; fi6++ {
        v2 += -1
        _ = fi6
    }
    if (false && (1 <= 958)) {
        delete(v3, v5)
        v7 := GenS1{F0: v2, F1: 42, F2: 1}
        _ = v7
    }
    v2 -= v2
    v5 += v5
    switch ("\n\t" + "") {
    case "":
        var v8 string = ""
        _ = v8
    default:
        delete(v3, 661)
    }
    v5 = v5
    fmt.Println(v2, v2)
    v9 := GenS1{F0: (v2 * v2), F1: 715, F2: 18446744073709551615}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
