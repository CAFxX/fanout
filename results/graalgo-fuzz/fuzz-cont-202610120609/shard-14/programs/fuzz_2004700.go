package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int; F1 uint32; F2 int16 }

func main() {
    v1 := make(map[int]string)
    var v2 uint64 = 18446744073709551615
    var v3 uint32 = 42
    var v4 int8 = 16
    v5 := make(map[string]uint32)
    v6 := GenS1{F0: 65535}
    v7 := GenS2{F0: -9223372036854775808, F1: 162, F2: 1}
    v8 := GenS2{F0: 32768, F1: (^(v3)), F2: 42}
    switch (("hello" + "a") + "a") {
    case "foo bar":
        v9 := GenS1{F0: 1}
        _ = v9
    }
    v5[("a" + ("" + (("" + "") + "a")))] = 255
    if (((v4 >> ((2) & 63)) >= 0) && true) {
        var v10 uint32 = (v3 >> ((65535) & 63))
        v11 := []int{}
        v5["hello"] = (v10 * v10)
        _ = v10
        _ = v11
    } else {
        var v12 uint64 = (v2 ^ (v2 | v2))
        _ = v12
    }
    var v13 int8 = (v4 & (v4 * 1))
    var v14 float32 = 1e10
    v1[42] = (("foo bar" + "hello") + ("" + (("" + ("\n\t" + ("\n\t" + "foo bar"))) + "")))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
