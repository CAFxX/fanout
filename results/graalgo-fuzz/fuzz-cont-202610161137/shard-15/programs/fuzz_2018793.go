package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int; F2 string }
func genFn0() uint16 {
    var v4 float32 = 0.1
    if (false || (-31 >= 20)) {
        v4 = float32(0.0)
    }
    fmt.Println(v4)
    v4 = 1.5
    _ = v4
    return 32767
}
func genFn1(p0 uint16, p1 uint, p2 uint) int {
    fmt.Println(p2, p2, p1)
    if true {
        v5 := []int{}
        var v6 int64 = -9223372036854775808
        if true {
            v6 = (-(v6))
            v7 := GenS1{F0: (v6 | v6), F1: 32768, F2: ""}
            fmt.Println(p0)
            _ = v7
        } else {
            p1 -= 255
        }
        _ = v5
        _ = v6
    } else {
        p2 = (p1 >> ((p1) & 63))
    }
    if true {
        p2 = (p1 * (^(p1)))
    }
    var v8 uint32 = 266
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 472
}

func main() {
    var v1 uint = 1
    v2 := make(map[int]int)
    v3 := make(map[string]string)
    var v9 string = "\n\t"
    if (42 > 1) {
        v9 = ("\n\t" + "foo bar")
    }
    v1 = (v1 ^ 4294967295)
    for fi10 := 0; fi10 < 14; fi10++ {
        v11 := GenS1{F0: 0, F1: fi10, F2: v9}
        _ = fi10
        _ = v11
    }
    v9 = (v9 + ("hello" + v9))
    switch "a" {
    case "foo bar":
        delete(v2, 4294967295)
    case "hello":
        v3[""] = v9
    case "\n\t":
        v12 := []int{}
        _ = v12
    default:
        var v13 uint = 998
        _ = v13
    }
    fmt.Println(v1, v1)
    v14 := []int{genFn1(32767, 4294967295, v1), -1, -2147483648}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v9
}
