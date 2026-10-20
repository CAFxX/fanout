package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16 }
func genFn0(p0 uint, p1 int16) uint8 {
    v5 := GenS1{F0: p0, F1: p1}
    if true {
        var v6 uint16 = 65535
        _ = v6
    }
    _ = p0
    _ = p1
    _ = v5
    return 228
}

func main() {
    v1 := make(map[string]uint32)
    v2 := []uint32{2147483647, 2147483648, 65535, 2147483647}
    var v3 bool = (2 < 32767)
    var v4 uint64 = 18446744073709551615
    delete(v1, (("foo bar" + "") + ("\n\t" + "a")))
    fmt.Println(v4, v3)
    v2 = append(v2, 364)
    fmt.Println(v4, v4)
    var v7 bool = ("a" != ("hello" + ("\n\t" + ("" + ""))))
    fmt.Println(v3, v4)
    fmt.Println(v7, v3, v3)
    _ = len(v2)
    var v8 float32 = (3.14159 - (-2.25 / 0.1))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
