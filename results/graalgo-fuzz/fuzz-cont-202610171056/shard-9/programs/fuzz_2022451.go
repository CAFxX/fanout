package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64; F2 int }
func genFn0(p0 uint64, p1 uint, p2 int16) uint32 {
    v6 := []int{1}
    p2 -= (p2 - (^(p2)))
    v6 = append(v6, -529)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 1
}

func main() {
    var v1 int32 = -2147483648
    v2 := make(map[string]uint32)
    var v3 uint8 = 2
    var v4 bool = false
    var v5 uint = 18446744073709551615
    v7 := GenS1{F0: ("a" + ("a" + "hello")), F1: 232, F2: -191}
    if (!false) {
        var v8 bool = v4
        _ = v8
    }
    var v9 float32 = 1.5
    v10 := []int{250, -253}
    v3 -= (v3 + 128)
    v11 := GenS1{F0: "\n\t", F1: -1, F2: 32767}
    v12 := GenS1{F0: "hello", F1: 0, F2: 32767}
    v13 := GenS1{F0: (("hello" + ("\n\t" + (("a" + "hello") + "foo bar"))) + "foo bar"), F1: 42, F2: 128}
    v14 := GenS1{F0: "foo bar", F1: 9223372036854775807, F2: 42}
    if false {
        var v15 int16 = -32768
        _ = v15
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v9
}
