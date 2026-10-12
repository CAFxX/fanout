package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32 }
func genFn0(p0 uint16) int {
    fmt.Println(p0)
    var v6 int64 = -722
    _ = p0
    _ = v6
    return 4294967295
}
func genFn1() int32 {
    v7 := GenS1{F0: 100, F1: 2147483648}
    var v8 uint64 = 18446744073709551615
    var v9 int = 414
    v8 = v8
    _ = v7
    _ = v8
    _ = v9
    return 42
}

func main() {
    var v1 int16 = -350
    var v2 bool = ("a" <= ("foo bar" + (("hello" + "hello") + (("hello" + "") + ("a" + "a")))))
    v3 := []string{("hello" + ""), ("\n\t" + "hello"), "\n\t", (((((("hello" + "\n\t") + "foo bar") + "hello") + ("a" + "foo bar")) + ("a" + ("" + (("a" + "foo bar") + "")))) + ("hello" + "\n\t"))}
    v4 := make(map[uint32]uint32)
    var v5 int = 32768
    fmt.Println(v2, v2, v5)
    v4[65535] = 42
    var v10 float64 = float64(float32(((3.14159 * 1e10) - ((1.5 * (3.14159 / float64((0.1 * 0.1)))) / 3.14159))))
    if false {
        fmt.Println(v10, v10)
        fmt.Println(v10, v5, v2)
    }
    v5 -= (v5 << ((42) & 63))
    var v11 int64 = -1
    fmt.Println(v11)
    v10 = 0.0
    var v12 float64 = v10
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
