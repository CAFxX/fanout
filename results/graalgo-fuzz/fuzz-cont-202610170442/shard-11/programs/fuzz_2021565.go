package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16; F2 int32 }
func genFn0(p0 uint64, p1 int8, p2 uint32) string {
    var v6 uint32 = 2147483648
    v7 := GenS1{F0: 2, F1: -694, F2: 0}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return (("hello" + "foo bar") + ("foo bar" + (("" + "foo bar") + (((("foo bar" + "") + "hello") + "\n\t") + ("hello" + ("foo bar" + "foo bar"))))))
}

func main() {
    var v1 uint64 = 346
    var v2 int = -2147483648
    v3 := make(map[string]string)
    v4 := []string{("" + ""), "hello"}
    v5 := make(map[int]bool)
    v1 += (v1 - (v1 | (v1 + (^(v1)))))
    v5[(v2 ^ (v2 >> ((607) & 63)))] = true
    v4 = append(v4, "hello")
    var v8 uint8 = 2
    v5[(v2 / ((42) | 1))] = false
    if (true && true) {
        v4 = append(v4, "foo bar")
    }
    var v9 float32 = 1e10
    _ = len(v4)
    fmt.Println(v8, v8, v9)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
