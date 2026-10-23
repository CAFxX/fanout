package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
func genFn0(p0 int) string {
    var v6 string = ("a" + "foo bar")
    fmt.Println(p0)
    p0 -= (^(p0))
    fmt.Println(v6, v6)
    var v7 int32 = -2147483648
    _ = p0
    _ = v6
    _ = v7
    return (("hello" + ((v6 + "a") + v6)) + (v6 + ("\n\t" + (((v6 + "\n\t") + (v6 + "hello")) + "foo bar"))))
}

func main() {
    v1 := make(map[string]uint32)
    var v2 uint64 = 18446744073709551615
    var v3 int32 = 1
    var v4 uint16 = 42
    v5 := make(map[int]bool)
    v8 := []int{42, -375, 128}
    fmt.Println(v3, v3)
    _ = len(v5)
    v9 := GenS1{F0: 255}
    _ = len(v8)
    var v10 uint64 = 18446744073709551615
    fmt.Println(v3, v4)
    var v11 int8 = -128
    v3 = int32(v4)
    v3 = v3
    v12 := GenS1{F0: 2}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
