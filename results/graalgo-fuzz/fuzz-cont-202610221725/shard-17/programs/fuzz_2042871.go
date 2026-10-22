package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16; F2 string }
func genFn0(p0 int8, p1 int) int32 {
    var v7 string = "\n\t"
    p0 = (p0 * p0)
    fmt.Println(v7, p1)
    var v8 uint8 = 67
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 42
}

func main() {
    v1 := []uint32{42}
    var v2 int16 = 91
    var v3 int8 = 127
    var v4 int = -809
    var v5 float64 = -2.25
    v6 := make(map[int]int)
    v9 := GenS1{F0: 4294967295, F1: (v2 + 0), F2: (("foo bar" + ((((("hello" + "\n\t") + "foo bar") + "") + (("" + "") + ("hello" + "a"))) + (("hello" + "\n\t") + "\n\t"))) + "a")}
    v6[32768] = v4
    switch "\n\t" {
    case "\n\t":
        var v10 int64 = 9223372036854775807
        _ = v10
    case "":
        v3 += v3
    case "hello":
        v5 = v5
    }
    v2 = v2
    if len(v1) > 0 {
        v1[0] = 0
    }
    fmt.Println(v5, v4)
    v11 := GenS1{F0: 214, F1: v2, F2: "\n\t"}
    var v12 int8 = 100
    if (0.0 >= 0.0) {
        var v13 int64 = -798
        delete(v6, int(v13))
        _ = v13
    }
    var v14 uint64 = 42
    var v15 uint16 = 1
    var v16 float32 = 1.5
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v11
    _ = v12
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
