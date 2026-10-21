package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool; F2 int32 }

func main() {
    v1 := make(map[string]string)
    var v2 uint8 = 0
    v3 := make(map[int]bool)
    var v4 uint32 = 373
    v5 := make(map[uint32]uint32)
    v2 += 165
    v4 += (^(v4))
    fmt.Println(v4)
    v6 := GenS1{F0: 829, F1: (-9223372036854775808 != -459), F2: -1}
    v7 := GenS1{F0: 1, F1: true, F2: -2147483648}
    if true {
        v1[("\n\t" + "\n\t")] = ((((("hello" + ("\n\t" + "a")) + "a") + "") + "") + (("foo bar" + ("\n\t" + ("a" + "foo bar"))) + ""))
    } else {
        fmt.Println(v2)
    }
    if (((3.14159 - 0.0) + 1.5) >= 3.14159) {
        v8 := []int{1, 324}
        var v9 bool = false
        _ = v8
        _ = v9
    } else {
        v2 = v2
        v4 -= v4
        v10 := []int{128}
        _ = v10
    }
    fmt.Println(v4, v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
