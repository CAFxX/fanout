package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8; F2 uint }
func genFn0(p0 uint64, p1 uint32) uint32 {
    var v4 uint = 420
    v5 := GenS1{F0: 0, F1: 255, F2: v4}
    v6 := []int{0, 147, int(p0)}
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    return 0
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[int]string)
    var v3 int8 = 61
    v3 = (v3 - (v3 / ((v3) | 1)))
    fmt.Println(v3)
    v7 := v1[("hello" + ((("\n\t" + (("" + ((("hello" + "") + "foo bar") + "hello")) + "a")) + (("hello" + "hello") + ((("hello" + "a") + ("\n\t" + "")) + "hello"))) + ("a" + "")))]
    var v8 string = v7
    v7 = (("foo bar" + v8) + (v7 + v8))
    v1[v7] = ((("\n\t" + v8) + v8) + ((v7 + v8) + (("foo bar" + v8) + "foo bar")))
    v8 = "foo bar"
    v9 := GenS1{F0: 1, F1: 88, F2: 42}
    fmt.Println(v8)
    v10 := GenS1{F0: 0, F1: 255, F2: 481}
    switch ("foo bar" + v8) {
    case "\n\t":
        v11 := GenS1{F0: 1, F1: 1, F2: 255}
        _ = v11
    }
    var v12 int8 = v3
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v12
    _ = v2
    _ = v3
    _ = v7
    _ = v8
    _ = v9
}
