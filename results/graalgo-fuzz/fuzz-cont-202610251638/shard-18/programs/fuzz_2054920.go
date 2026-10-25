package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32; F2 uint8 }
func genFn0(p0 int8, p1 uint64, p2 uint) int {
    var v6 uint32 = 139
    p1 = p1
    v7 := []int{}
    p0 = p0
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 58
}

func main() {
    v1 := make(map[int]uint32)
    v2 := []string{"a", (("" + ("foo bar" + "a")) + "foo bar"), "a", "foo bar"}
    v3 := []string{"hello", ("\n\t" + "a"), "a"}
    var v4 int32 = 741
    var v5 uint32 = 503
    var v8 int64 = 744
    if (v8 >= v8) {
        v9 := GenS1{F0: v4, F1: v5, F2: 232}
        var v10 int = -2147483648
        var v11 int = v10
        _ = v10
        _ = v11
        _ = v9
    } else {
        var v12 string = ""
        _ = v12
    }
    v5 -= 42
    v5 -= v5
    switch ((("" + ("a" + "a")) + "hello") + "foo bar") {
    case "hello":
        v8 = v8
    case "":
        if true {
            fmt.Println(v5)
            var v13 int32 = v4
            _ = v13
        } else {
            fmt.Println(v5)
        }
    case "foo bar":
        v8 += v8
    }
    var v14 uint64 = 18446744073709551615
    v5 = 2
    v8 = v8
    var v15 int8 = 100
    v14 -= v14
    fmt.Println(v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
