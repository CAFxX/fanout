package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string; F2 int16 }

func main() {
    var v1 uint32 = 4294967295
    v2 := make(map[uint32]bool)
    v3 := []int{2147483647, 9223372036854775807}
    v4 := make(map[int]bool)
    v5 := make(map[int]string)
    var v6 int = -662
    if len(v3) > 0 {
        v3[0] = v6
    }
    switch (v6 ^ (v6 << ((42) & 63))) {
    case 2147483647:
        if (false || (-128 == -128)) {
            var v7 int32 = 1
            _ = v7
        }
    case 127:
        v1 += v1
    case 42:
        v1 += (^(v1))
    default:
        fmt.Println(v6, v6)
    }
    v1 -= v1
    var v8 uint64 = 123
    v6 += -446
    v1 = v1
    v9 := GenS1{F0: ("hello" + ((("" + ("\n\t" + "\n\t")) + ("" + ("foo bar" + "\n\t"))) + "a")), F1: ((("hello" + ("\n\t" + ("foo bar" + ("" + "")))) + ("foo bar" + (("foo bar" + "a") + "foo bar"))) + ("hello" + "hello")), F2: 42}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
