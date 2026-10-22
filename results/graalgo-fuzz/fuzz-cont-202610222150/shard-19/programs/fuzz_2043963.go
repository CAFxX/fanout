package main

import (
    "fmt"
)

func genFn0(p0 uint16) int64 {
    var v6 uint64 = 42
    fmt.Println(v6, p0)
    p0 = 32767
    var v7 float64 = 1.5
    _ = p0
    _ = v6
    _ = v7
    return 42
}
func genFn1(p0 uint16, p1 uint8, p2 uint32) uint64 {
    v8 := []int{403}
    switch "hello" {
    case "hello":
        fmt.Println(p2)
    }
    var v9 int8 = 0
    var v10 int32 = 637
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v8
    _ = v9
    return 18446744073709551615
}

func main() {
    var v1 uint32 = 0
    v2 := make(map[uint32]int)
    var v3 int32 = -2147483648
    v4 := make(map[uint32]int)
    v5 := make(map[uint32]string)
    v3 -= (v3 + 1)
    fmt.Println(v3, v1)
    v3 -= (v3 + -422)
    var v11 bool = false
    v4[42] = 0
    switch (("\n\t" + ("foo bar" + "a")) + "a") {
    case "\n\t":
        var v12 int64 = -9223372036854775808
        _ = v12
    default:
        v1 += 761
    }
    v3 = 42
    fmt.Println(v1)
    switch "hello" {
    case "":
        var v13 int32 = (v3 + -865)
        _ = v13
    case "a":
        var v14 float64 = 0.0
        _ = v14
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
