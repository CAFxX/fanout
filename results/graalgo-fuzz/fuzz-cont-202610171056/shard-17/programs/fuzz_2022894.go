package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64; F2 int8 }

func main() {
    var v1 int8 = 110
    v2 := []uint32{4294967295, 42}
    v3 := []string{(("hello" + (("" + "a") + "hello")) + "hello"), "\n\t", "", "a"}
    var v4 int64 = -913
    var v5 uint8 = 255
    v6 := GenS1{F0: false, F1: 1, F2: (v1 >> ((324) & 63))}
    v1 = (v1 * 100)
    if len(v2) > 0 {
        v2[0] = 2147483648
    }
    var v7 int32 = -1
    for fi8 := 0; fi8 < 5; fi8++ {
        var v9 float32 = 1e10
        v5 = (v5 + (v5 | 42))
        v7 -= 2147483647
        _ = fi8
        _ = v9
    }
    v3 = append(v3, (((("foo bar" + ("foo bar" + "hello")) + "") + ("hello" + (("" + ("a" + "foo bar")) + "\n\t"))) + "a"))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
