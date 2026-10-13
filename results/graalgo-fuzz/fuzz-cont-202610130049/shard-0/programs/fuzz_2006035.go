package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
func genFn0(p0 int, p1 int64) int32 {
    p0 += p0
    v7 := []int{128, (p0 + p0)}
    p1 = p1
    switch "a" {
    case "a":
        var v8 uint16 = 449
        _ = v8
    }
    p1 = 543
    _ = p0
    _ = p1
    _ = v7
    return 2147483647
}

func main() {
    var v1 uint8 = 127
    v2 := make(map[string]string)
    var v3 int32 = 42
    v4 := []string{("hello" + "a"), "foo bar"}
    v5 := make(map[int]uint32)
    var v6 string = "hello"
    switch 127 {
    case 128:
        var v9 int16 = 449
        _ = v9
    case -536:
        var v10 uint32 = 65535
        _ = v10
    case 127:
        v1 = (v1 % ((v1) | 1))
    }
    if ((((0.1 / (1.5 + (0.1 / (-2.25 * 3.14159)))) + (-2.25 + 1.5)) == (((((0.1 + ((0.1 / -2.25) - ((1e10 * (3.14159 * 0.1)) - (1e10 / 1.0)))) + (((-2.25 / ((3.14159 / float32(((1e10 - (0.1 / 1.0)) / 0.1))) * (0.0 + 3.14159))) + 1e10) - 1e10)) / ((1e10 / float32(((0.1 * 1.5) - 0.1))) - 3.14159)) + 1.5) - 1e10)) && (!(!false))) {
        var v11 uint8 = v1
        _ = v11
    } else {
        delete(v5, 840)
    }
    fmt.Println(v1, v1, v6)
    fmt.Println(v1, v1, v6)
    switch 4294967295 {
    case -715:
        v1 = 93
    }
    switch v6 {
    case "hello":
        v3 = (v3 ^ v3)
    }
    switch 2147483647 {
    case -456:
        v4 = append(v4, "hello")
    default:
        v12 := GenS1{F0: 42}
        _ = v12
    }
    _ = len(v4)
    fmt.Println(v6)
    v1 += v1
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
