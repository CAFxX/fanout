package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int16; F1 uint64 }
func genFn0(p0 uint16) int32 {
    var v6 uint32 = 65535
    if false {
        v7 := []int{1}
        _ = v7
    }
    _ = p0
    _ = v6
    return 1
}

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[uint32]int)
    v3 := []string{"foo bar", "foo bar", "hello", ("a" + ("foo bar" + "foo bar"))}
    var v4 float32 = 1e10
    var v5 uint = 0
    v8 := GenS1{F0: 128}
    v9 := GenS1{F0: 245}
    v4 -= ((((0.0 + (((v4 + -2.25) + 1e10) * v4)) - v4) - ((3.14159 / 1.0) - v4)) * (v4 / v4))
    var v10 uint16 = 42
    v11 := GenS2{F0: 1, F1: 4294967295}
    if false {
        var v12 int64 = 9223372036854775807
        _ = v12
    } else {
        var v13 string = (("hello" + "foo bar") + ((("foo bar" + ("\n\t" + "a")) + "hello") + ""))
        _ = v13
    }
    fmt.Println(v5)
    var v14 int64 = 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
