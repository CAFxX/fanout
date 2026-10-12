package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }

func main() {
    v1 := make(map[string]string)
    var v2 int32 = 332
    v3 := make(map[int]bool)
    v4 := make(map[uint32]int)
    v2 += v2
    v2 -= -2147483648
    switch ("hello" + (("foo bar" + "hello") + "")) {
    case "\n\t":
        if false {
            v2 = v2
            v5 := []int{}
            var v6 int = 42
            _ = v5
            _ = v6
        } else {
            v2 -= (^(v2))
        }
    case "a":
        v7 := []int{-9223372036854775808, 128}
        _ = v7
    }
    v2 -= v2
    if true {
        _ = len(v1)
        var v8 int8 = -128
        _ = v8
    }
    v9 := []int{42}
    if true {
        _ = len(v9)
    } else {
        v10 := GenS1{F0: 18446744073709551615}
        _ = v10
    }
    v2 += 0
    fmt.Println(v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
